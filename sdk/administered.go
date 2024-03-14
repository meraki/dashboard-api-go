package meraki

import (
	"fmt"

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



 */
func (s *AdministeredService) GetAdministeredIDentitiesMe() (*ResponseAdministeredGetAdministeredIDentitiesMe, *resty.Response, error) {
	path := "/api/v1/administered/identities/me"
	s.rateLimiterBucket.Wait(1)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseAdministeredGetAdministeredIDentitiesMe{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetAdministeredIdentitiesMe")
	}

	result := response.Result().(*ResponseAdministeredGetAdministeredIDentitiesMe)
	return result, response, err

}
