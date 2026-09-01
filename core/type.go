package core

import "encoding/json"

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
	Sub                 string `json:"sub"`                   // The user's unique ID.
	Name                string `json:"name"`                  // The user's full name.
	Username            string `json:"username"`              // The user's username.
	Picture             string `json:"picture"`               // The user's profile picture URL.
	Email               string `json:"email"`                 // The user's email address.
	EmailVerified       bool   `json:"email_verified"`        // Whether the user's email address is verified.
	PhoneNumber         string `json:"phone_number"`          // The user's phone number.
	PhoneNumberVerified bool   `json:"phone_number_verified"` // Whether the user's phone number is verified.
	CreatedAt           int64  `json:"created_at"`            // Time when the user was created, in milliseconds since the Unix epoch.
	UpdatedAt           int64  `json:"updated_at"`            // Time when the user was last updated, in milliseconds since the Unix epoch.
	// The following claims are standard OIDC claims included in the `profile` scope.
	// Logto only returns them when their values are not empty, so absent claims are
	// left as zero values.
	FamilyName        string                 `json:"family_name"`        // The user's family name.
	GivenName         string                 `json:"given_name"`         // The user's given name.
	MiddleName        string                 `json:"middle_name"`        // The user's middle name.
	Nickname          string                 `json:"nickname"`           // The user's nickname.
	PreferredUsername string                 `json:"preferred_username"` // The username by which the user wishes to be referred to.
	Profile           string                 `json:"profile"`            // The URL of the user's profile page.
	Website           string                 `json:"website"`            // The URL of the user's website.
	Gender            string                 `json:"gender"`             // The user's gender.
	Birthdate         string                 `json:"birthdate"`          // The user's birthdate.
	Zoneinfo          string                 `json:"zoneinfo"`           // The user's time zone, e.g. `Europe/Paris`.
	Locale            string                 `json:"locale"`             // The user's locale, e.g. `en-US`.
	CustomData        map[string]interface{} `json:"custom_data"`        // The user's custom data
	Identities        map[string]interface{} `json:"identities"`         // The user's social identities information
	Roles             []string               `json:"roles"`              // The role names of the current user.
	Organizations     []string               `json:"organizations"`      // The organization IDs that the user has membership.
	// The organization roles that the user has.
	// Each role is in the format of `<organization_id>:<role_name>`.
	// # Example #
	// The following array indicates that user is an admin of org1 and a member of org2:
	// ```go
	// {"org1:admin", "org2:member"}
	// ```
	OrganizationRoles []string       `json:"organization_roles"`
	OrganizationData  []Organization `json:"organization_data"` // The organization data that the user has membership.

	// rawClaims holds all claims decoded from the raw JSON payload, including
	// claims that are not modeled as struct fields. Use GetClaim to access them.
	rawClaims map[string]any
}

// userInfoResponseAlias is an alias of UserInfoResponse used to avoid infinite
// recursion when unmarshaling.
type userInfoResponseAlias UserInfoResponse

// UnmarshalJSON implements the json.Unmarshaler interface.
// In addition to populating the modeled fields, it keeps all claims in the raw
// JSON payload so that unmodeled claims can be accessed via GetClaim.
func (userInfoResponse *UserInfoResponse) UnmarshalJSON(data []byte) error {
	var alias userInfoResponseAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}

	var rawClaims map[string]any
	if err := json.Unmarshal(data, &rawClaims); err != nil {
		return err
	}

	*userInfoResponse = UserInfoResponse(alias)
	userInfoResponse.rawClaims = rawClaims

	return nil
}

// GetClaim returns the raw value of the given claim in the userinfo response,
// including claims that are not modeled as struct fields, e.g. custom claims.
// The second return value indicates whether the claim is present.
func (userInfoResponse UserInfoResponse) GetClaim(name string) (any, bool) {
	value, ok := userInfoResponse.rawClaims[name]
	return value, ok
}

type IdTokenClaims struct {
	Iss                 string `json:"iss"`
	Sub                 string `json:"sub"`
	Aud                 string `json:"aud"`
	Exp                 int64  `json:"exp"`
	Iat                 int64  `json:"iat"`
	AtHash              string `json:"at_hash"`
	Name                string `json:"name"`
	Username            string `json:"username"`
	Picture             string `json:"picture"`
	Email               string `json:"email"`
	EmailVerified       bool   `json:"email_verified"`
	PhoneNumber         string `json:"phone_number"`
	PhoneNumberVerified bool   `json:"phone_number_verified"`
	CreatedAt           int64  `json:"created_at"` // Time when the user was created, in milliseconds since the Unix epoch.
	UpdatedAt           int64  `json:"updated_at"` // Time when the user was last updated, in milliseconds since the Unix epoch.
	// The following claims are standard OIDC claims included in the `profile` scope.
	// Logto only returns them when their values are not empty, so absent claims are
	// left as zero values.
	FamilyName        string   `json:"family_name"`        // The user's family name.
	GivenName         string   `json:"given_name"`         // The user's given name.
	MiddleName        string   `json:"middle_name"`        // The user's middle name.
	Nickname          string   `json:"nickname"`           // The user's nickname.
	PreferredUsername string   `json:"preferred_username"` // The username by which the user wishes to be referred to.
	Profile           string   `json:"profile"`            // The URL of the user's profile page.
	Website           string   `json:"website"`            // The URL of the user's website.
	Gender            string   `json:"gender"`             // The user's gender.
	Birthdate         string   `json:"birthdate"`          // The user's birthdate.
	Zoneinfo          string   `json:"zoneinfo"`           // The user's time zone, e.g. `Europe/Paris`.
	Locale            string   `json:"locale"`             // The user's locale, e.g. `en-US`.
	Roles             []string `json:"roles"`
	Organizations     []string `json:"organizations"`
	OrganizationRoles []string `json:"organization_roles"`

	// rawClaims holds all claims decoded from the raw JSON payload, including
	// claims that are not modeled as struct fields. Use GetClaim to access them.
	rawClaims map[string]any
}

// idTokenClaimsAlias is an alias of IdTokenClaims used to avoid infinite
// recursion when unmarshaling.
type idTokenClaimsAlias IdTokenClaims

// UnmarshalJSON implements the json.Unmarshaler interface.
// In addition to populating the modeled fields, it keeps all claims in the raw
// JSON payload so that unmodeled claims can be accessed via GetClaim.
func (idTokenClaims *IdTokenClaims) UnmarshalJSON(data []byte) error {
	var alias idTokenClaimsAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}

	var rawClaims map[string]any
	if err := json.Unmarshal(data, &rawClaims); err != nil {
		return err
	}

	*idTokenClaims = IdTokenClaims(alias)
	idTokenClaims.rawClaims = rawClaims

	return nil
}

// GetClaim returns the raw value of the given claim in the ID token, including
// claims that are not modeled as struct fields, e.g. custom claims.
// The second return value indicates whether the claim is present.
func (idTokenClaims IdTokenClaims) GetClaim(name string) (any, bool) {
	value, ok := idTokenClaims.rawClaims[name]
	return value, ok
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
