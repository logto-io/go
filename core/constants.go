package core

var (
	ReservedScopeOpenId        = "openid"
	ReservedScopeOfflineAccess = "offline_access"
)

var (
	ReservedResourceOrganization = "urn:logto:resource:organizations"
)

var (
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
