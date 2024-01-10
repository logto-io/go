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

type Organization struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UserInfoResponse struct {
	Sub                 string                 `json:"sub"`                   // The user's unique ID.
	Name                string                 `json:"name"`                  // The user's full name.
	Username            string                 `json:"username"`              // The user's username.
	Picture             string                 `json:"picture"`               // The user's profile picture URL.
	Email               string                 `json:"email"`                 // The user's email address.
	EmailVerified       bool                   `json:"email_verified"`        // Whether the user's email address is verified.
	PhoneNumber         string                 `json:"phone_number"`          // The user's phone number.
	PhoneNumberVerified bool                   `json:"phone_number_verified"` // Whether the user's phone number is verified.
	CustomData          map[string]interface{} `json:"custom_data"`           // The user's custom data
	Identities          map[string]interface{} `json:"identities"`            // The user's social identities information
	Roles               []string               `json:"roles"`                 // The role names of the current user.
	Organizations       []string               `json:"organizations"`         // The organization IDs that the user has membership.
	// The organization roles that the user has.
	// Each role is in the format of `<organization_id>:<role_name>`.
	// # Example #
	// The following array indicates that user is an admin of org1 and a member of org2:
	// ```go
	// {"org1:admin", "org2:member"}
	// ```
	OrganizationRoles []string       `json:"organization_roles"`
	OrganizationData  []Organization `json:"organization_data"` // The organization data that the user has membership.
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
