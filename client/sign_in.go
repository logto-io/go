package client

import (
	"encoding/json"

	"logto.io/core"
)

type SignInContext struct {
	RedirectUri   string
	CodeVerifier  string
	CodeChallenge string
	State         string
}

func (logtoClient *LogtoClient) SignIn(redirectUri string) (string, error) {
	oidcConfig, fetchOidcConfigErr := logtoClient.fetchOidcConfig()

	if fetchOidcConfigErr != nil {
		return "", fetchOidcConfigErr
	}

	codeVerifier := core.GenerateCodeVerifier()
	codeChallenge := core.GenerateCodeChallenge(codeVerifier)
	state := core.GenerateState()

	signInUri, generateSignInUriErr := core.GenerateSignInUri(&core.SignInUriGenerationOptions{
		AuthorizationEndpoint: oidcConfig.AuthorizationEndpoint,
		ClientId:              logtoClient.logtoConfig.AppId,
		RedirectUri:           redirectUri,
		CodeChallenge:         codeChallenge,
		State:                 state,
		Scopes:                logtoClient.logtoConfig.Scopes,
		Resources:             logtoClient.logtoConfig.Resources,
		Prompt:                logtoClient.logtoConfig.Prompt,
	})

	if generateSignInUriErr != nil {
		return "", generateSignInUriErr
	}

	signInContext := SignInContext{
		RedirectUri:   redirectUri,
		CodeVerifier:  codeVerifier,
		CodeChallenge: codeChallenge,
		State:         state,
	}

	signInContextJsonValue, marshalErr := json.Marshal(signInContext)
	if marshalErr != nil {
		return "", marshalErr
	}

	logtoClient.storage.SetItem("logto_sign_in_context", string(signInContextJsonValue))

	return signInUri, nil
}
