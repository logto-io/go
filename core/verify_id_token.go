package core

import (
	"time"

	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

var ISSUED_AT_RESTRICTIONS int64 = 60 // in seconds

func VerifyIdToken(idToken, clientId, issuer string, jwks jose.JSONWebKeySet) error {
	jws, err := jwt.ParseSigned(idToken)
	if err != nil {
		return err
	}

	// Verify signature
	idTokenClaims := IdTokenClaims{}
	verifySignatureError := jws.Claims(jwks, &idTokenClaims)

	if verifySignatureError != nil {
		return verifySignatureError
	}

	// Verify claims
	if idTokenClaims.Iss != issuer {
		return ErrTokenIssuerNotMatch
	}

	if idTokenClaims.Aud != clientId {
		return ErrTokenAudienceNotMatch
	}

	now := time.Now().Unix()

	if idTokenClaims.Exp < now {
		return ErrTokenExpired
	}

	if idTokenClaims.Iat > now+ISSUED_AT_RESTRICTIONS {
		return ErrTokenIssuedInTheFuture
	}

	if idTokenClaims.Iat < now-ISSUED_AT_RESTRICTIONS {
		return ErrTokenIssuedInThePast
	}

	return nil
}
