package meraki

import (
	"github.com/go-resty/resty/v2"
)

type AdministeredService service

type ResponseAdministeredGetAdministeredIDentitiesMe struct {
	Authentication      *ResponseAdministeredGetAdministeredIDentitiesMeAuthentication `json:"authentication,omitempty"`      // Authentication info
	Email               string                                                         `json:"email,omitempty"`               // User email
	LastUsedDashboardAt string                                                         `json:"lastUsedDashboardAt,omitempty"` // Last seen active on Dashboard UI
	Name                string                                                         `json:"name,omitempty"`                // Username
}
type ResponseAdministeredGetAdministeredIDentitiesMeAuthentication struct {
	API       *ResponseAdministeredGetAdministeredIDentitiesMeAuthenticationAPI       `json:"api,omitempty"`       // API authentication
	Mode      string                                                                  `json:"mode,omitempty"`      // Authentication mode
	Saml      *ResponseAdministeredGetAdministeredIDentitiesMeAuthenticationSaml      `json:"saml,omitempty"`      // SAML authentication
	TwoFactor *ResponseAdministeredGetAdministeredIDentitiesMeAuthenticationTwoFactor `json:"twoFactor,omitempty"` // TwoFactor authentication
}
type ResponseAdministeredGetAdministeredIDentitiesMeAuthenticationAPI struct {
	Key *ResponseAdministeredGetAdministeredIDentitiesMeAuthenticationAPIKey `json:"key,omitempty"` // API key
}
type ResponseAdministeredGetAdministeredIDentitiesMeAuthenticationAPIKey struct {
	Created *bool `json:"created,omitempty"` // If API key is created for this user
}
type ResponseAdministeredGetAdministeredIDentitiesMeAuthenticationSaml struct {
	Enabled *bool `json:"enabled,omitempty"` // If SAML authentication is enabled for this user
}
type ResponseAdministeredGetAdministeredIDentitiesMeAuthenticationTwoFactor struct {
	Enabled *bool `json:"enabled,omitempty"` // If twoFactor authentication is enabled for this user
}

//GetAdministeredIDentitiesMe Returns the identity of the current user.
/* Returns the identity of the current user.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-administered-identities-me
*/
func (s *AdministeredService) GetAdministeredIDentitiesMe() (*ResponseAdministeredGetAdministeredIDentitiesMe, *resty.Response, error) {
	path := "/api/v1/administered/identities/me"

	return Get[ResponseAdministeredGetAdministeredIDentitiesMe](s.client, s.rateLimiterBucket, path, &ResponseAdministeredGetAdministeredIDentitiesMe{})
}
