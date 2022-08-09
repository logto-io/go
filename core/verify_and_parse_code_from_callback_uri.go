package core

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

func VerifyAndParseCodeFromCallbackUri(callbackUri, redirectUri, state string) (string, error) {
	if !strings.HasPrefix(callbackUri, redirectUri) {
		return "", ErrCallbackUriNotMatchRedirectUri
	}

	parsedUrl, urlParseError := url.Parse(callbackUri)
	if urlParseError != nil {
		return "", urlParseError
	}

	errorInUri := parsedUrl.Query().Get("error")

	if errorInUri != "" {
		errorDescriptionInUri := parsedUrl.Query().Get("error_description")
		if errorDescriptionInUri == "" {
			return "", errors.New(errorInUri)
		}
		return "", fmt.Errorf("%s: %s", errorInUri, errorDescriptionInUri)
	}

	stateInUri := parsedUrl.Query().Get("state")

	if stateInUri != state {
		return "", ErrStateNotMatch
	}

	code := parsedUrl.Query().Get("code")

	if code == "" {
		return "", ErrCodeNotFoundInCallbackUri
	}

	return code, nil
}
