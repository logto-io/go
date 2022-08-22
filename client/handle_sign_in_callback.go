package client

import (
	"encoding/json"
	"time"

	"logto.io/core"
)

func (logtoClient *LogtoClient) HandleSignInCallback(callbackUri string) (core.CodeTokenResponse, error) {
	signInContext := SignInContext{}
	parseSignInContextErr := json.Unmarshal([]byte(logtoClient.sessionStorage.GetItem("logto_sign_in_context")), &signInContext)

	if parseSignInContextErr != nil {
		return core.CodeTokenResponse{}, parseSignInContextErr
	}

	code, retrieveCodeErr := core.VerifyAndParseCodeFromCallbackUri(callbackUri, signInContext.RedirectUri, signInContext.State)
	if retrieveCodeErr != nil {
		return core.CodeTokenResponse{}, retrieveCodeErr
	}

	oidcConfig, fetchOidcConfigErr := logtoClient.fetchOidcConfig()

	if fetchOidcConfigErr != nil {
		return core.CodeTokenResponse{}, fetchOidcConfigErr
	}

	codeTokenResponse, fetchTokenErr := core.FetchTokenByAuthorizationCode(logtoClient.httpClient, &core.FetchTokenByAuthorizationCodeOptions{
		TokenEndpoint: oidcConfig.TokenEndpoint,
		Code:          code,
		CodeVerifier:  signInContext.CodeVerifier,
		ClientId:      logtoClient.logtoConfig.AppId,
		RedirectUri:   signInContext.RedirectUri,
	})

	if fetchTokenErr != nil {
		return core.CodeTokenResponse{}, fetchTokenErr
	}

	jwks, createJwksErr := logtoClient.createRemoteJwks(oidcConfig.JwksUri)
	if createJwksErr != nil {
		return core.CodeTokenResponse{}, createJwksErr
	}

	verificationErr := core.VerifyIdToken(codeTokenResponse.IdToken, logtoClient.logtoConfig.AppId, oidcConfig.Issuer, jwks)

	if verificationErr != nil {
		return core.CodeTokenResponse{}, verificationErr
	}

	logtoClient.sessionStorage.SetItem("logto_sign_in_context", "")

	logtoClient.SetRefreshToken(codeTokenResponse.RefreshToken)
	logtoClient.SetIdToken(codeTokenResponse.IdToken)

	// TODO: extract resource from access token (audience)
	accessTokenKey := buildAccessTokenKey([]string{}, "")

	logtoClient.SaveAccessToken(accessTokenKey, AccessToken{
		token:     codeTokenResponse.AccessToken,
		scope:     codeTokenResponse.Scope,
		expiresAt: time.Now().Unix() + int64(codeTokenResponse.ExpireIn),
	})

	return codeTokenResponse, nil
}
