package client

import (
	"net/http"
	"time"

	"golang.org/x/exp/slices"

	"github.com/logto-io/go/core"
)

type AccessToken struct {
	Token     string `json:"token"`
	Scope     string `json:"scope"`
	ExpiresAt int64  `json:"expiresAt"`
}

type LogtoClient struct {
	httpClient     *http.Client
	logtoConfig    *LogtoConfig
	storage        Storage
	accessTokenMap map[string]AccessToken
}

func NewLogtoClient(config *LogtoConfig, storage Storage) *LogtoClient {
	config.normalized()
	logtoClient := LogtoClient{
		httpClient:     &http.Client{},
		logtoConfig:    config,
		storage:        storage,
		accessTokenMap: make(map[string]AccessToken),
	}

	logtoClient.loadAccessTokenMap()

	return &logtoClient
}

func (logtoClient *LogtoClient) IsAuthenticated() bool {
	return logtoClient.GetIdToken() != ""
}

func (logtoClient *LogtoClient) GetRefreshToken() string {
	return logtoClient.storage.GetItem(StorageKeyRefreshToken)
}

func (logtoClient *LogtoClient) SetRefreshToken(refreshToken string) {
	logtoClient.storage.SetItem(StorageKeyRefreshToken, refreshToken)
}

func (LogtoClient *LogtoClient) GetIdToken() string {
	return LogtoClient.storage.GetItem(StorageKeyIdToken)
}

func (logtoClient *LogtoClient) SetIdToken(idToken string) {
	logtoClient.storage.SetItem(StorageKeyIdToken, idToken)
}

func (logtoClient *LogtoClient) GetIdTokenClaims() (core.IdTokenClaims, error) {
	if !logtoClient.IsAuthenticated() {
		return core.IdTokenClaims{}, ErrNotAuthenticated
	}
	return core.DecodeIdToken(logtoClient.GetIdToken())
}

func (logtoClient *LogtoClient) GetOrganizationTokenClaims(organizationId string) (core.OrganizationAccessTokenClaims, error) {
	token, getTokenErr := logtoClient.GetOrganizationToken(organizationId)

	if getTokenErr != nil {
		return core.OrganizationAccessTokenClaims{}, getTokenErr
	}

	jwtObject, parseTokenErr := core.ParseSignedJwt(token.Token)

	if parseTokenErr != nil {
		return core.OrganizationAccessTokenClaims{}, parseTokenErr
	}

	var claims core.OrganizationAccessTokenClaims
	claimsErr := jwtObject.UnsafeClaimsWithoutVerification(&claims)

	if claimsErr != nil {
		return core.OrganizationAccessTokenClaims{}, claimsErr
	}

	return claims, claimsErr
}

func (logtoClient *LogtoClient) SaveAccessToken(key string, accessToken AccessToken) {
	logtoClient.accessTokenMap[key] = accessToken
	logtoClient.persistAccessTokenMap()
}

func (logtoClient *LogtoClient) getAccessToken(resource string, organizationId string) (AccessToken, error) {
	if !logtoClient.IsAuthenticated() {
		return AccessToken{}, ErrNotAuthenticated
	}

	if resource != "" {
		if !slices.Contains(logtoClient.logtoConfig.Resources, resource) {
			return AccessToken{}, ErrUnacknowledgedResourceFound
		}
	}

	accessTokenKey := buildAccessTokenKey([]string{}, resource, organizationId)
	if accessToken, ok := logtoClient.accessTokenMap[accessTokenKey]; ok {
		if accessToken.ExpiresAt > time.Now().Unix() {
			return accessToken, nil
		}
	}

	refreshToken := logtoClient.GetRefreshToken()

	if refreshToken == "" {
		return AccessToken{}, ErrNotAuthenticated
	}

	oidcConfig, fetchOidcConfigErr := logtoClient.fetchOidcConfig()

	if fetchOidcConfigErr != nil {
		return AccessToken{}, fetchOidcConfigErr
	}

	refreshedToken, refreshTokenErr := core.FetchTokenByRefreshToken(logtoClient.httpClient, &core.FetchTokenByRefreshTokenOptions{
		TokenEndpoint:  oidcConfig.TokenEndpoint,
		ClientId:       logtoClient.logtoConfig.AppId,
		ClientSecret:   logtoClient.logtoConfig.AppSecret,
		RefreshToken:   refreshToken,
		Resource:       resource,
		Scopes:         []string{},
		OrganizationId: organizationId,
	})

	if refreshTokenErr != nil {
		return AccessToken{}, refreshTokenErr
	}

	refreshedAccessToken := AccessToken{
		Token:     refreshedToken.AccessToken,
		Scope:     refreshedToken.Scope,
		ExpiresAt: time.Now().Unix() + int64(refreshedToken.ExpireIn),
	}

	verificationErr := logtoClient.verifyAndSaveTokenResponse(
		refreshedToken.IdToken,
		refreshedToken.RefreshToken,
		accessTokenKey,
		refreshedAccessToken,
		&oidcConfig,
	)

	if verificationErr != nil {
		return AccessToken{}, verificationErr
	}

	return refreshedAccessToken, nil
}

func (logtoClient *LogtoClient) GetAccessToken(resource string) (AccessToken, error) {
	return logtoClient.getAccessToken(resource, "")
}

func (logtoClient *LogtoClient) GetOrganizationToken(organizationId string) (AccessToken, error) {
	if !slices.Contains(logtoClient.logtoConfig.Scopes, core.UserScopeOrganizations) {
		return AccessToken{}, ErrMissingScopeOrganizations
	}

	return logtoClient.getAccessToken("", organizationId)
}

func (logtoClient *LogtoClient) FetchUserInfo() (core.UserInfoResponse, error) {
	if !logtoClient.IsAuthenticated() {
		return core.UserInfoResponse{}, ErrNotAuthenticated
	}

	oidcConfig, fetchOidcConfigErr := logtoClient.fetchOidcConfig()

	if fetchOidcConfigErr != nil {
		return core.UserInfoResponse{}, fetchOidcConfigErr
	}

	accessToken, getAccessTokenErr := logtoClient.GetAccessToken("")

	if getAccessTokenErr != nil {
		return core.UserInfoResponse{}, getAccessTokenErr
	}

	return core.FetchUserInfo(oidcConfig.UserinfoEndpoint, accessToken.Token)
}
