package client

import (
	"encoding/json"

	"github.com/logto-io/go/core"
)

type SignInSession struct {
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

	signInSession := SignInSession{
		RedirectUri:   redirectUri,
		CodeVerifier:  codeVerifier,
		CodeChallenge: codeChallenge,
		State:         state,
	}

	signInSessionJsonValue, marshalErr := json.Marshal(signInSession)
	if marshalErr != nil {
		return "", marshalErr
	}

	logtoClient.storage.SetItem(StorageKeySignInSession, string(signInSessionJsonValue))

	return signInUri, nil
}
