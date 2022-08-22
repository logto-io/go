package client

import (
	"net/http"
	"net/url"
	"sort"
	"strings"

	"gopkg.in/square/go-jose.v2/jwt"
)

func GetOriginRequestUrl(request *http.Request) string {
	return getRequestProtocol(request) + "://" + request.Host + request.RequestURI
}

func getRequestProtocol(request *http.Request) string {
	if request.TLS != nil {
		return "https"
	}
	proto := request.Header.Get("X-Forwarded-Proto")
	if proto != "" {
		extractedProto := strings.Split(proto, ",")[0]
		return strings.ToLower(strings.Trim(extractedProto, " "))
	}
	return "http"
}

func createHttpClient(logtoConfig *LogtoConfig) *http.Client {
	defaultTransport := http.DefaultTransport.(*http.Transport)
	customTransport := defaultTransport.Clone()
	customTransport.Proxy = func(req *http.Request) (*url.URL, error) {
		if logtoConfig.AppSecret != "" {
			req.SetBasicAuth(logtoConfig.AppId, logtoConfig.AppSecret)
		}
		return defaultTransport.Proxy(req)
	}

	return &http.Client{
		Transport: customTransport,
	}
}

func buildAccessTokenKey(scopes []string, resource string) string {
	sort.Strings(scopes)
	scopesPart := strings.Join(scopes, " ")
	return scopesPart + "@" + resource
}

func getResourceFromAccessToken(accessToken string) string {
	jwtObject, parseToJwtErr := jwt.ParseSigned(accessToken)
	if parseToJwtErr != nil {
		return ""
	}

	type audContainedClaims struct {
		Aud string `json:"aud"`
	}

	var audClaim audContainedClaims
	claimsErr := jwtObject.UnsafeClaimsWithoutVerification(&audClaim)
	if claimsErr != nil {
		return ""
	}

	return audClaim.Aud
}
