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
	Email               string                 `json:"email"`
	EmailVerified       bool                   `json:"email_verified"`
	PhoneNumber         string                 `json:"phone_number"`
	PhoneNumberVerified bool                   `json:"phone_number_verified"`
	CustomData          map[string]interface{} `json:"custom_data"`
	Identities          map[string]interface{} `json:"identities"`
}

type IdTokenClaims struct {
	Iss                 string   `json:"iss"`
	Sub                 string   `json:"sub"`
	Aud                 string   `json:"aud"`
	Exp                 int64    `json:"exp"`
	Iat                 int64    `json:"iat"`
	AtHash              string   `json:"at_hash"`
	Name                string   `json:"name"`
	Username            string   `json:"username"`
	Picture             string   `json:"picture"`
	Email               string   `json:"email"`
	EmailVerified       bool     `json:"email_verified"`
	PhoneNumber         string   `json:"phone_number"`
	PhoneNumberVerified bool     `json:"phone_number_verified"`
	Roles               []string `json:"roles"`
	Organizations       []string `json:"organizations"`
	OrganizationRoles   []string `json:"organization_roles"`
}

type OrganizationAccessTokenClaims struct {
	Iss      string `json:"iss"`
	Sub      string `json:"sub"`
	Aud      string `json:"aud"`
	Exp      int64  `json:"exp"`
	Iat      int64  `json:"iat"`
	ClientId string `json:"client_id"`
	Jti      string `json:"jti"`
	Scope    string `json:"scope"`
}
