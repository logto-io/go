package client

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/logto-io/go/core"
)

func (logtoClient *LogtoClient) HandleSignInCallback(request *http.Request) error {
	signInSession := SignInSession{}
	parseSignInSessionErr := json.Unmarshal([]byte(logtoClient.storage.GetItem(StorageKeySignInSession)), &signInSession)

	if parseSignInSessionErr != nil {
		return parseSignInSessionErr
	}

	callbackUri := GetOriginRequestUrl(request)
	code, retrieveCodeErr := core.VerifyAndParseCodeFromCallbackUri(callbackUri, signInSession.RedirectUri, signInSession.State)
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
		CodeVerifier:  signInSession.CodeVerifier,
		ClientId:      logtoClient.logtoConfig.AppId,
		ClientSecret:  logtoClient.logtoConfig.AppSecret,
		RedirectUri:   signInSession.RedirectUri,
	})

	if fetchTokenErr != nil {
		return fetchTokenErr
	}

	logtoClient.storage.SetItem(StorageKeySignInSession, "")

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
