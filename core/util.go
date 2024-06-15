package core

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jwt"
)

func AppendIfNotExisted(targets []string, item string) []string {
	for _, s := range targets {
		if s == item {
			return targets
		}
	}

	return append(targets, item)
}

func ParseSignedJwt(token string) (*jwt.JSONWebToken, error) {
	parts := strings.Split(token, ".")

	if len(parts) != 3 {
		return nil, fmt.Errorf("token format is invalid")
	}

	rawHeader, decodeHeaderErr := base64.RawURLEncoding.DecodeString(parts[0])

	if decodeHeaderErr != nil {
		return nil, decodeHeaderErr
	}

	var header struct {
		Alg string `json:"alg"`
	}

	unmarshalErr := json.Unmarshal(rawHeader, &header)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	// Check the `alg` field
	if header.Alg == "" {
		return nil, fmt.Errorf("alg field is missing in token header")
	}

	return jwt.ParseSigned(token, []jose.SignatureAlgorithm{jose.SignatureAlgorithm(header.Alg)})
}
