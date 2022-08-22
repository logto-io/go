package client

import (
	"net/http"
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
	token     string
	scope     string
	expiresAt int64
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

func (logtoClient *LogtoClient) GetRefreshToken() string {
	return logtoClient.storage.GetItem("logto_refresh_token")
}

func (logtoClient *LogtoClient) SetRefreshToken(refreshToken string) {
	logtoClient.storage.SetItem("logto_refresh_token", refreshToken)
}

func (LogtoClient *LogtoClient) GetIdToken() string {
	return LogtoClient.storage.GetItem("logto_id_token")
}

func (logtoClient *LogtoClient) SetIdToken(idToken string) {
	logtoClient.storage.SetItem("logto_id_token", idToken)
}

func (logtoClient *LogtoClient) SaveAccessToken(key string, accessToken AccessToken) {
	logtoClient.accessTokenMap[key] = accessToken
	if logtoClient.logtoConfig.PersistAccessToken {
		logtoClient.persistAccessTokenMap()
	}
}

func (logtoClient *LogtoClient) GetAccessToken(key string) AccessToken {
	return logtoClient.accessTokenMap[key]
}
