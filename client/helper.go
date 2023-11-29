package client

import (
	"encoding/json"
	"net/url"

	"github.com/logto-io/go/core"
	"gopkg.in/square/go-jose.v2"
)

func (logtoClient *LogtoClient) fetchOidcConfig() (core.OidcConfigResponse, error) {
	discoveryEndpoint, constructEndpointErr := url.JoinPath(logtoClient.logtoConfig.Endpoint, "/oidc/.well-known/openid-configuration")
	if constructEndpointErr != nil {
		return core.OidcConfigResponse{}, constructEndpointErr
	}
	return core.FetchOidcConfig(logtoClient.httpClient, discoveryEndpoint)
}

func (logtoClient *LogtoClient) loadAccessTokenMap() {
	accessTokenMap := make(map[string]AccessToken)
	accessTokenMapJsonString := logtoClient.storage.GetItem(StorageKeyAccessTokenMap)
	if accessTokenMapJsonString == "" {
		return
	}

	unmarshalErr := json.Unmarshal([]byte(accessTokenMapJsonString), &accessTokenMap)
	if unmarshalErr != nil {
		return
	}

	logtoClient.accessTokenMap = accessTokenMap
}

func (logtoClient *LogtoClient) persistAccessTokenMap() {
	accessTokenMapJsonString, err := json.Marshal(logtoClient.accessTokenMap)
	if err != nil {
		return
	}
	logtoClient.storage.SetItem(StorageKeyAccessTokenMap, string(accessTokenMapJsonString))
}

func (logtoClient *LogtoClient) createRemoteJwks(jwksUri string) (*jose.JSONWebKeySet, error) {
	jwksResponse, fetchJwksErr := core.FetchJwks(logtoClient.httpClient, jwksUri)
	if fetchJwksErr != nil {
		return nil, fetchJwksErr
	}

	jwks := jose.JSONWebKeySet{}
	for _, rawJsonWebKeyData := range jwksResponse.Keys {
		// Note: convert rawJsonWebKeyData to JSON string for we need to unmarshal it to JSONWebKey
		rawJsonWebKeyJsonString, parseToJsonWebKeyJsonErr := json.Marshal(rawJsonWebKeyData)
		if parseToJsonWebKeyJsonErr != nil {
			return nil, parseToJsonWebKeyJsonErr
		}

		jwk := jose.JSONWebKey{}
		// Note: Use rawJsonWebKeyJsonString to construct the JsonWebKey
		parseToJsonWebKeyErr := jwk.UnmarshalJSON(rawJsonWebKeyJsonString)
		if parseToJsonWebKeyErr != nil {
			return nil, parseToJsonWebKeyErr
		}

		jwks.Keys = append(jwks.Keys, jwk)
	}

	return &jwks, nil
}

func (logtoClient *LogtoClient) verifyAndSaveTokenResponse(
	idToken string,
	refreshToken string,
	accessTokenKey string,
	accessToken AccessToken,
	oidcConfig *core.OidcConfigResponse,
) error {
	if idToken != "" {
		jwks, createJwksErr := logtoClient.createRemoteJwks(oidcConfig.JwksUri)
		if createJwksErr != nil {
			return createJwksErr
		}

		verificationErr := core.VerifyIdToken(idToken, logtoClient.logtoConfig.AppId, oidcConfig.Issuer, jwks)
		if verificationErr != nil {
			return verificationErr
		}

		logtoClient.SetIdToken(idToken)
	}

	logtoClient.SetRefreshToken(refreshToken)

	logtoClient.SaveAccessToken(accessTokenKey, accessToken)

	return nil
}
