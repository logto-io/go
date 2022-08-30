package client

import (
	"errors"

	"github.com/logto-io/go/core"
)

func (logtoClient *LogtoClient) SignOut(postLogoutRedirectUri string) (string, error) {
	idToken := logtoClient.GetIdToken()
	if idToken == "" {
		return "", errors.New("not authenticated")
	}

	refreshToken := logtoClient.GetRefreshToken()

	logtoClient.accessTokenMap = make(map[string]AccessToken)
	logtoClient.storage.SetItem(StorageKeyAccessTokenMap, "")
	logtoClient.storage.SetItem(StorageKeyRefreshToken, "")
	logtoClient.storage.SetItem(StorageKeyIdToken, "")

	oidcConfig, fetchOidcConfigErr := logtoClient.fetchOidcConfig()

	if fetchOidcConfigErr != nil {
		return "", fetchOidcConfigErr
	}

	if refreshToken != "" {
		_ = core.Revoke(logtoClient.httpClient, &core.RevocationOptions{
			RevocationEndpoint: oidcConfig.RevocationEndpoint,
			ClientId:           logtoClient.logtoConfig.AppId,
			Token:              refreshToken,
		})
	}

	signOutUri, generateSignOutUriErr := core.GenerateSignOutUri(&core.SignOutUriGenerationOptions{
		EndSessionEndpoint:    oidcConfig.EndSessionEndpoint,
		IdToken:               idToken,
		PostLogoutRedirectUri: postLogoutRedirectUri,
	})

	if generateSignOutUriErr != nil {
		return "", generateSignOutUriErr
	}

	return signOutUri, nil
}
