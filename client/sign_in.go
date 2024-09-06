package client

import (
	"encoding/json"

	"github.com/logto-io/go/core"
)

type SignInOptions struct {
	RedirectUri  string
	Prompt       string
	FirstScreen  string
	Identifiers  []string
	DirectSignIn *core.DirectSignInOptions
	LoginHint    string
	ExtraParams  map[string]string
}

type SignInSession struct {
	RedirectUri   string
	CodeVerifier  string
	CodeChallenge string
	State         string
}

func (logtoClient *LogtoClient) SignIn(options *SignInOptions) (string, error) {
	oidcConfig, fetchOidcConfigErr := logtoClient.fetchOidcConfig()

	if fetchOidcConfigErr != nil {
		return "", fetchOidcConfigErr
	}

	codeVerifier := core.GenerateCodeVerifier()
	codeChallenge := core.GenerateCodeChallenge(codeVerifier)
	state := core.GenerateState()

	prompt := options.Prompt
	if prompt == "" {
		prompt = logtoClient.logtoConfig.Prompt
	}

	signInUri, generateSignInUriErr := core.GenerateSignInUri(&core.SignInUriGenerationOptions{
		AuthorizationEndpoint: oidcConfig.AuthorizationEndpoint,
		ClientId:              logtoClient.logtoConfig.AppId,
		RedirectUri:           options.RedirectUri,
		CodeChallenge:         codeChallenge,
		State:                 state,
		Scopes:                logtoClient.logtoConfig.Scopes,
		Resources:             logtoClient.logtoConfig.Resources,
		Prompt:                prompt,
		FirstScreen:           options.FirstScreen,
		Identifiers:           options.Identifiers,
		DirectSignIn:          options.DirectSignIn,
		LoginHint:             options.LoginHint,
		ExtraParams:           options.ExtraParams,
		IncludeReservedScopes: logtoClient.logtoConfig.IncludeReservedScopes,
	})

	if generateSignInUriErr != nil {
		return "", generateSignInUriErr
	}

	signInSession := SignInSession{
		RedirectUri:   options.RedirectUri,
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

func (logtoClient *LogtoClient) SignInWithRedirectUri(redirectUri string) (string, error) {
	return logtoClient.SignIn(&SignInOptions{
		RedirectUri: redirectUri,
	})
}
