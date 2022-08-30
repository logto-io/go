package client

import (
	"errors"
	"net/http"
	"time"

	"golang.org/x/exp/slices"

	"github.com/logto-io/go/core"
)

type LogtoConfig struct {
	Endpoint  string
	AppId     string
	AppSecret string
	Scopes    []string
	Resources []string
	Prompt    string
	// TODO: Find a way to make this value default to true
	PersistAccessToken bool
}

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
	logtoClient := LogtoClient{
		httpClient:     createHttpClient(config),
		logtoConfig:    config,
		storage:        storage,
		accessTokenMap: make(map[string]AccessToken),
	}

	if config.PersistAccessToken {
		logtoClient.loadAccessTokenMap()
	}

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
	return core.DecodeIdToken(logtoClient.GetIdToken())
}

func (logtoClient *LogtoClient) SaveAccessToken(key string, accessToken AccessToken) {
	logtoClient.accessTokenMap[key] = accessToken
	if logtoClient.logtoConfig.PersistAccessToken {
		logtoClient.persistAccessTokenMap()
	}
}

func (logtoClient *LogtoClient) GetAccessToken(resource string) (AccessToken, error) {
	if !logtoClient.IsAuthenticated() {
		return AccessToken{}, errors.New("not authenticated")
	}

	if !slices.Contains(logtoClient.logtoConfig.Resources, resource) {
		return AccessToken{}, errors.New("unacknowledged resource found")
	}

	accessTokenKey := buildAccessTokenKey([]string{}, resource)
	if accessToken, ok := logtoClient.accessTokenMap[accessTokenKey]; ok {
		if accessToken.ExpiresAt > time.Now().Unix() {
			return accessToken, nil
		}
	}

	refreshToken := logtoClient.GetRefreshToken()

	if refreshToken == "" {
		return AccessToken{}, errors.New("not authenticated")
	}

	oidcConfig, fetchOidcConfigErr := logtoClient.fetchOidcConfig()

	if fetchOidcConfigErr != nil {
		return AccessToken{}, fetchOidcConfigErr
	}

	var scopes []string
	if resource != "" {
		scopes = append(scopes, "offline_access")
	}

	refreshedToken, refreshTokenErr := core.FetchTokenByRefreshToken(logtoClient.httpClient, &core.FetchTokenByRefreshTokenOptions{
		TokenEndpoint: oidcConfig.TokenEndpoint,
		ClientId:      logtoClient.logtoConfig.AppId,
		RefreshToken:  refreshToken,
		Scopes:        scopes,
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
		refreshedAccessToken,
		&oidcConfig,
	)

	if verificationErr != nil {
		return AccessToken{}, verificationErr
	}

	return refreshedAccessToken, nil
}
