package client

import (
	"github.com/logto-io/go/v2/core"
)

func (logtoClient *LogtoClient) SignOut(postLogoutRedirectUri string) (string, error) {
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
		ClientId:              logtoClient.logtoConfig.AppId,
		PostLogoutRedirectUri: postLogoutRedirectUri,
	})

	if generateSignOutUriErr != nil {
		return "", generateSignOutUriErr
	}

	return signOutUri, nil
}
