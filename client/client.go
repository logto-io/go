package client

import (
	"net/http"
	"time"

	"golang.org/x/exp/slices"

	"github.com/logto-io/go/v2/core"
)

type AccessToken struct {
	Token     string `json:"token"`
	Scope     string `json:"scope"`
	ExpiresAt int64  `json:"expiresAt"`
}

// GetAccessTokenOptions contains parameters for retrieving an access token.
// Use Resource to specify the API resource, and OrganizationId to specify the organization context.
// Both fields are optional; leave them empty if not needed.
type GetAccessTokenOptions struct {
	Resource       string
	OrganizationId string
}

// GetOrganizationTokenClaimsOptions contains parameters for retrieving organization access token claims.
// Use Resource to specify the API resource, and OrganizationId to specify the organization whose claims you want to retrieve.
// OrganizationId is required.
type GetOrganizationTokenClaimsOptions struct {
	Resource       string
	OrganizationId string
}

// LogtoClientOption is a functional option for configuring LogtoClient
type LogtoClientOption func(*LogtoClient)

// WithHttpClient sets a custom HTTP client for the LogtoClient
func WithHttpClient(client *http.Client) LogtoClientOption {
	return func(c *LogtoClient) {
		c.httpClient = client
	}
}

type LogtoClient struct {
	httpClient     *http.Client
	logtoConfig    *LogtoConfig
	storage        Storage
	accessTokenMap map[string]AccessToken
}

func NewLogtoClient(config *LogtoConfig, storage Storage, opts ...LogtoClientOption) *LogtoClient {
	config.normalized()
	logtoClient := LogtoClient{
		httpClient:     &http.Client{},
		logtoConfig:    config,
		storage:        storage,
		accessTokenMap: make(map[string]AccessToken),
	}

	// Apply options
	for _, opt := range opts {
		opt(&logtoClient)
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

func (logtoClient *LogtoClient) SaveAccessToken(key string, accessToken AccessToken) {
	logtoClient.accessTokenMap[key] = accessToken
	logtoClient.persistAccessTokenMap()
}

// GetAccessTokenWithOptions retrieves an access token for the specified resource and/or organization.
// Use GetAccessToken for resource-only tokens, or GetOrganizationToken for organization-only tokens.
// This method provides the most flexibility and is recommended for advanced scenarios.
func (logtoClient *LogtoClient) GetAccessTokenWithOptions(options GetAccessTokenOptions) (AccessToken, error) {
	if !logtoClient.IsAuthenticated() {
		return AccessToken{}, ErrNotAuthenticated
	}

	if options.Resource != "" {
		if !slices.Contains(logtoClient.logtoConfig.Resources, options.Resource) {
			return AccessToken{}, ErrUnacknowledgedResourceFound
		}
	}

	if options.OrganizationId != "" {
		if !slices.Contains(logtoClient.logtoConfig.Scopes, core.UserScopeOrganizations) {
			return AccessToken{}, ErrMissingScopeOrganizations
		}
	}

	accessTokenKey := buildAccessTokenKey([]string{}, options.Resource, options.OrganizationId)
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
		Resource:       options.Resource,
		Scopes:         []string{},
		OrganizationId: options.OrganizationId,
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

// GetOrganizationTokenClaimsWithOptions retrieves the claims from an organization access token
// for the specified resource and organization. OrganizationId is required.
// Use GetOrganizationTokenClaims for organization-only claims.
// This method is recommended for advanced scenarios where both resource and organization context are needed.
func (logtoClient *LogtoClient) GetOrganizationTokenClaimsWithOptions(options GetOrganizationTokenClaimsOptions) (core.OrganizationAccessTokenClaims, error) {
	if options.OrganizationId == "" {
		return core.OrganizationAccessTokenClaims{}, ErrMissingOrganizationId
	}

	token, getTokenErr := logtoClient.GetAccessTokenWithOptions(GetAccessTokenOptions{
		Resource:       options.Resource,
		OrganizationId: options.OrganizationId,
	})

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

// GetAccessToken retrieves an access token for the specified resource only.
// This method does not support organization-based access.
// If you need to specify an organization, use GetAccessTokenWithOptions instead.
func (logtoClient *LogtoClient) GetAccessToken(resource string) (AccessToken, error) {
	return logtoClient.GetAccessTokenWithOptions(GetAccessTokenOptions{
		Resource: resource,
	})
}

// GetOrganizationToken retrieves an access token for the specified organization only.
// This method does not support resource-based access.
// If you need to specify a resource, use GetAccessTokenWithOptions instead.
func (logtoClient *LogtoClient) GetOrganizationToken(organizationId string) (AccessToken, error) {
	return logtoClient.GetAccessTokenWithOptions(GetAccessTokenOptions{
		OrganizationId: organizationId,
	})
}

// GetOrganizationTokenClaims retrieves the claims from an organization access token for the specified organization only.
// This method does not support resource-based claims retrieval.
// If you need to specify a resource, use GetOrganizationTokenClaimsWithOptions instead.
func (logtoClient *LogtoClient) GetOrganizationTokenClaims(organizationId string) (core.OrganizationAccessTokenClaims, error) {
	return logtoClient.GetOrganizationTokenClaimsWithOptions(GetOrganizationTokenClaimsOptions{
		OrganizationId: organizationId,
	})
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

	return core.FetchUserInfoWithClient(logtoClient.httpClient, oidcConfig.UserinfoEndpoint, accessToken.Token)
}
