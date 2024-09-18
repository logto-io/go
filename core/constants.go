package core

const (
	ReservedScopeOpenId        = "openid"
	ReservedScopeOfflineAccess = "offline_access"
)

const (
	ReservedResourceOrganization = "urn:logto:resource:organizations"
)

const (
	UserScopeProfile           = "profile"
	UserScopeEmail             = "email"
	UserScopePhone             = "phone"
	UserScopeCustomData        = "custom_data"
	UserScopeIdentities        = "identities"
	UserScopeRoles             = "roles"
	UserScopeOrganizations     = "urn:logto:scope:organizations"
	UserScopeOrganizationRoles = "urn:logto:scope:organization_roles"
)

var (
	DefaultScopes = []string{
		ReservedScopeOpenId,
		ReservedScopeOfflineAccess,
		UserScopeProfile,
	}
)

const (
	QueryKeyClientID            = "client_id"
	QueryKeyRedirectURI         = "redirect_uri"
	QueryKeyCodeChallenge       = "code_challenge"
	QueryKeyCodeChallengeMethod = "code_challenge_method"
	QueryKeyState               = "state"
	QueryKeyScope               = "scope"
	QueryKeyResource            = "resource"
	QueryKeyResponseType        = "response_type"
	QueryKeyPrompt              = "prompt"
	QueryKeyLoginHint           = "login_hint"
	QueryKeyFirstScreen         = "first_screen"
	QueryKeyIdentifiers         = "identifiers"
	QueryKeyDirectSignIn        = "direct_sign_in"
)

const (
	IdentifierEmail    = "email"
	IdentifierPhone    = "phone"
	IdentifierUsername = "username"
)

const (
	DirectSignInMethodSocial = "social"
	DirectSignInMethodSso    = "sso"
)

const (
	PromptConsent = "consent"
	PromptLogin   = "login"
)

const (
	ResponseTypeCode = "code"
)

const (
	FirstScreenSignIn             = "sign_in"
	FirstScreenRegister           = "register"
	FirstScreenResetPassword      = "reset_password"
	FirstScreenSingleSignOn       = "single_sign_on"
	FirstScreenIdentifierSignIn   = "identifier:sign_in"
	FirstScreenIdentifierRegister = "identifier:register"
)
