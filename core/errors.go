package core

import "errors"

var (
	ErrTokenIssuerNotMatch            = errors.New("token issuer not match")
	ErrTokenAudienceNotMatch          = errors.New("token audience not match")
	ErrTokenExpired                   = errors.New("token expired")
	ErrTokenIssuedInTheFuture         = errors.New("token issued in the future")
	ErrTokenIssuedInThePast           = errors.New("token issued in the past")
	ErrCallbackUriNotMatchRedirectUri = errors.New("callback uri not match redirect uri")
	ErrStateNotMatch                  = errors.New("state not match")
	ErrCodeNotFoundInCallbackUri      = errors.New("code not found in callback uri")
)
