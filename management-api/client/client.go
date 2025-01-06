package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/logto-io/go/management-api/logto"
	"golang.org/x/oauth2"
)

// LogtoTokenSource implements oauth2.TokenSource interface
type LogtoTokenSource struct {
	clientID     string
	clientSecret string
	resourceURL  string
	cacheStorage CacheStorage
	token        *oauth2.Token
	mutex        sync.RWMutex
	endpoint     string
}

type tokenCache struct {
	IssuedAt    time.Time `json:"issued_at"`
	AccessToken string    `json:"access_token"`
	ExpiresIn   int64     `json:"expires_in"`
	TokenType   string    `json:"token_type"`
	Scope       string    `json:"scope"`
}

// Token implements oauth2.TokenSource interface
func (l *LogtoTokenSource) Token() (*oauth2.Token, error) {
	if l.cacheStorage != nil {
		l.mutex.RLock()
		tokenCacheStr, err := l.cacheStorage.Get(l.clientID)
		if err != nil {
			l.mutex.RUnlock()
			return nil, fmt.Errorf("get access token from cache failed: %v", err)
		}

		if tokenCacheStr != "" {
			var tokenCache tokenCache
			if err := json.Unmarshal([]byte(tokenCacheStr), &tokenCache); err != nil {
				l.mutex.RUnlock()
				return nil, fmt.Errorf("decode access token from cache failed: %v", err)
			}

			if tokenCache.IssuedAt.Add(time.Duration(tokenCache.ExpiresIn) * time.Second).After(time.Now()) {
				l.token = &oauth2.Token{
					AccessToken: tokenCache.AccessToken,
					TokenType:   tokenCache.TokenType,
					Expiry:      tokenCache.IssuedAt.Add(time.Duration(tokenCache.ExpiresIn) * time.Second),
				}
				l.mutex.RUnlock()
				return l.token, nil
			}
		}
		l.mutex.RUnlock()
	} else {
		l.mutex.RLock()
		if l.token != nil && l.token.Valid() {
			defer l.mutex.RUnlock()
			return l.token, nil
		}
		l.mutex.RUnlock()
	}

	l.mutex.Lock()
	defer l.mutex.Unlock()

	// Request new token
	values := url.Values{
		"grant_type":    []string{"client_credentials"},
		"client_id":     []string{l.clientID},
		"client_secret": []string{l.clientSecret},
		"resource":      []string{l.resourceURL},
		"scope":         []string{"all"},
	}

	resp, err := http.PostForm(l.endpoint+"/oidc/token", values)
	if err != nil {
		return nil, fmt.Errorf("get access token failed: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read access token response failed: %v", err)
	}

	var tokenResp tokenCache
	if err := json.Unmarshal(respBody, &tokenResp); err != nil {
		return nil, fmt.Errorf("decode access token response failed: %v", err)
	}

	if tokenResp.AccessToken == "" {
		return nil, fmt.Errorf("get access token failed: %s", string(respBody))
	}

	if l.cacheStorage != nil {
		tokenResp.IssuedAt = time.Now()
		tokenCacheStr, err := json.Marshal(tokenResp)
		if err != nil {
			return nil, fmt.Errorf("encode access token to cache failed: %v", err)
		}
		l.cacheStorage.Setex(l.clientID, string(tokenCacheStr), int(tokenResp.ExpiresIn))
	}

	l.token = &oauth2.Token{
		AccessToken: tokenResp.AccessToken,
		TokenType:   tokenResp.TokenType,
		Expiry:      time.Now().Add(time.Duration(tokenResp.ExpiresIn) * time.Second),
	}

	return l.token, nil
}

type Client struct {
	*logto.APIClient
	openapiConfig *logto.Configuration
	Context       context.Context
}

func NewClient(config *Config, cache CacheStorage) *Client {
	urlInfo, err := url.Parse(config.Endpoint)
	if err != nil {
		panic(fmt.Sprintf("invalid endpoint: %v", err))
	}

	openapiConfig := logto.NewConfiguration()
	openapiConfig.Scheme = urlInfo.Scheme
	openapiConfig.Host = urlInfo.Host
	openapiConfig.Debug = config.Debug
	openapiConfig.Servers = []logto.ServerConfiguration{
		{
			URL: config.Endpoint,
		},
	}

	if config.ManagementApiResourceURL == "" {
		config.ManagementApiResourceURL = "https://default.logto.app/api"
	}

	tokenSource := &LogtoTokenSource{
		clientID:     config.AppId,
		clientSecret: config.AppSecret,
		cacheStorage: cache,
		endpoint:     config.Endpoint,
		resourceURL:  config.ManagementApiResourceURL,
	}

	client := &Client{
		openapiConfig: openapiConfig,
		APIClient:     logto.NewAPIClient(openapiConfig),
		Context:       context.WithValue(context.Background(), logto.ContextOAuth2, tokenSource),
	}

	return client
}
