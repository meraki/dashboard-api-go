package meraki

import (
	"fmt"
	"strings"

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
type ResponseAdministeredGetAdministeredIDentitiesMeAPIKeys []ResponseItemAdministeredGetAdministeredIDentitiesMeAPIKeys // Array of ResponseAdministeredGetAdministeredIdentitiesMeApiKeys
type ResponseItemAdministeredGetAdministeredIDentitiesMeAPIKeys struct {
	CreatedAt string `json:"createdAt,omitempty"` // Date that the API key was created
	Suffix    string `json:"suffix,omitempty"`    // Last 4 characters of the API key
}
type ResponseAdministeredGenerateAdministeredIDentitiesMeAPIKeys struct {
	Key string `json:"key,omitempty"` // API key in plaintext. This value will not be accessible outside of key generation
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

//GetAdministeredIDentitiesMeAPIKeys List the non-sensitive metadata associated with the API keys that belong to the user
/* List the non-sensitive metadata associated with the API keys that belong to the user



 */
func (s *AdministeredService) GetAdministeredIDentitiesMeAPIKeys() (*ResponseAdministeredGetAdministeredIDentitiesMeAPIKeys, *resty.Response, error) {
	path := "/api/v1/administered/identities/me/api/keys"
	s.rateLimiterBucket.Wait(1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseAdministeredGetAdministeredIDentitiesMeAPIKeys{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetAdministeredIdentitiesMeApiKeys")
	}

	result := response.Result().(*ResponseAdministeredGetAdministeredIDentitiesMeAPIKeys)
	return result, response, err

}

//GenerateAdministeredIDentitiesMeAPIKeys Generates an API key for an identity
/* Generates an API key for an identity. For users who have access to more than one organization, the change will take up to five minutes to propagate. If one of the organizations is currently under maintenance, the change may not propagate fully until after the maintenance has been completed.



 */

func (s *AdministeredService) GenerateAdministeredIDentitiesMeAPIKeys() (*ResponseAdministeredGenerateAdministeredIDentitiesMeAPIKeys, *resty.Response, error) {
	path := "/api/v1/administered/identities/me/api/keys/generate"
	s.rateLimiterBucket.Wait(1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseAdministeredGenerateAdministeredIDentitiesMeAPIKeys{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GenerateAdministeredIdentitiesMeApiKeys")
	}

	result := response.Result().(*ResponseAdministeredGenerateAdministeredIDentitiesMeAPIKeys)
	return result, response, err

}

//RevokeAdministeredIDentitiesMeAPIKeys Revokes an identity's API key, using the last four characters of the key
/* Revokes an identity's API key, using the last four characters of the key. For users who have access to more than one organization, the change will take up to five minutes to propagate. If one of the organizations is currently under maintenance, the change may not propagate fully until after the maintenance has been completed.

@param suffix suffix path parameter.


*/

func (s *AdministeredService) RevokeAdministeredIDentitiesMeAPIKeys(suffix string) (*resty.Response, error) {
	path := "/api/v1/administered/identities/me/api/keys/{suffix}/revoke"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{suffix}", fmt.Sprintf("%v", suffix), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation RevokeAdministeredIdentitiesMeApiKeys")
	}

	return response, err

}
