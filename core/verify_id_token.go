package core

import (
	"time"

	"github.com/go-jose/go-jose/v4"
)

var ISSUED_AT_RESTRICTIONS int64 = 60 // in seconds

func VerifyIdToken(idToken, clientId, issuer string, jwks *jose.JSONWebKeySet) error {
	jws, err := ParseSignedJwt(idToken)
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
