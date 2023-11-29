package client

import "errors"

var (
	ErrNotAuthenticated            = errors.New("not authenticated")
	ErrUnacknowledgedResourceFound = errors.New("unacknowledged resource found")
	ErrMissingScopeOrganizations   = errors.New("missing 'urn:logto:scope:organizations' scope")
)
