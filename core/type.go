package core

type OidcConfigResponse struct {
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	TokenEndpoint         string `json:"token_endpoint"`
	UserinfoEndpoint      string `json:"userinfo_endpoint"`
	EndSessionEndpoint    string `json:"end_session_endpoint"`
	RevocationEndpoint    string `json:"revocation_endpoint"`
	JwksUri               string `json:"jwks_uri"`
	Issuer                string `json:"issuer"`
}

type CodeTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	IdToken      string `json:"id_token"`
	Scope        string `json:"scope"`
	ExpireIn     int    `json:"expires_in"`
}

type RefreshTokenResponse = CodeTokenResponse

type UserInfoResponse struct {
	Sub                 string                 `json:"sub"`
	Name                string                 `json:"name"`
	Username            string                 `json:"username"`
	Picture             string                 `json:"picture"`
	RoleNames           []string               `json:"role_names"`
	Email               string                 `json:"email"`
	EmailVerified       bool                   `json:"email_verified"`
	PhoneNumber         string                 `json:"phone_number"`
	PhoneNumberVerified bool                   `json:"phone_number_verified"`
	CustomData          map[string]interface{} `json:"custom_data"`
	Identities          map[string]interface{} `json:"identities"`
}

type IdTokenClaims struct {
	Sub       string   `json:"sub"`
	Aud       string   `json:"aud"`
	Exp       int64    `json:"exp"`
	Iat       int64    `json:"iat"`
	Iss       string   `json:"iss"`
	AtHash    string   `json:"at_hash"`
	Username  string   `json:"username"`
	Name      string   `json:"name"`
	Avatar    string   `json:"avatar"`
	RoleNames []string `json:"role_names"`
}
