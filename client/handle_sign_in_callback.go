package client

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/logto-io/go/core"
)

func (logtoClient *LogtoClient) HandleSignInCallback(request *http.Request) error {
	signInContext := SignInContext{}
	parseSignInContextErr := json.Unmarshal([]byte(logtoClient.storage.GetItem("logto_sign_in_context")), &signInContext)

	if parseSignInContextErr != nil {
		return parseSignInContextErr
	}

	callbackUri := GetOriginRequestUrl(request)
	code, retrieveCodeErr := core.VerifyAndParseCodeFromCallbackUri(callbackUri, signInContext.RedirectUri, signInContext.State)
	if retrieveCodeErr != nil {
		return retrieveCodeErr
	}

	oidcConfig, fetchOidcConfigErr := logtoClient.fetchOidcConfig()

	if fetchOidcConfigErr != nil {
		return fetchOidcConfigErr
	}

	codeTokenResponse, fetchTokenErr := core.FetchTokenByAuthorizationCode(logtoClient.httpClient, &core.FetchTokenByAuthorizationCodeOptions{
		TokenEndpoint: oidcConfig.TokenEndpoint,
		Code:          code,
		CodeVerifier:  signInContext.CodeVerifier,
		ClientId:      logtoClient.logtoConfig.AppId,
		RedirectUri:   signInContext.RedirectUri,
	})

	if fetchTokenErr != nil {
		return fetchTokenErr
	}

	logtoClient.storage.SetItem("logto_sign_in_context", "")

	accessToken := AccessToken{
		Token:     codeTokenResponse.AccessToken,
		Scope:     codeTokenResponse.Scope,
		ExpiresAt: time.Now().Unix() + int64(codeTokenResponse.ExpireIn),
	}

	verificationErr := logtoClient.verifyAndSaveTokenResponse(
		codeTokenResponse.IdToken,
		codeTokenResponse.RefreshToken,
		accessToken,
		&oidcConfig,
	)

	if verificationErr != nil {
		return verificationErr
	}

	return nil
}
