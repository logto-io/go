package core

import "gopkg.in/square/go-jose.v2/jwt"

func DecodeIdToken(token string) (IdTokenClaims, error) {
	jwtObject, err := jwt.ParseSigned(token)

	if err != nil {
		return IdTokenClaims{}, err
	}

	var idTokenClaims IdTokenClaims
	claimsErr := jwtObject.UnsafeClaimsWithoutVerification(&idTokenClaims)

	if claimsErr != nil {
		return IdTokenClaims{}, claimsErr
	}

	return idTokenClaims, nil
}
