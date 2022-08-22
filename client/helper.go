package client

import (
	"encoding/json"
	"net/url"

	"gopkg.in/square/go-jose.v2"
	"logto.io/core"
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
	accessTokenMapJsonString := logtoClient.sessionStorage.GetItem("logto_access_token_map")
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
	logtoClient.sessionStorage.SetItem("logto_access_token_map", string(accessTokenMapJsonString))
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
