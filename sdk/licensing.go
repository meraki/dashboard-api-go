package meraki

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type LicensingService service

type GetOrganizationLicensingCotermLicensesQueryParams struct {
	PerPage       int    `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter string `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	Invalidated   bool   `url:"invalidated,omitempty"`   //Filter for licenses that are invalidated
	Expired       bool   `url:"expired,omitempty"`       //Filter for licenses that are expired
}

type ResponseLicensingGetOrganizationLicensingCotermLicenses []ResponseItemLicensingGetOrganizationLicensingCotermLicenses // Array of ResponseLicensingGetOrganizationLicensingCotermLicenses
type ResponseItemLicensingGetOrganizationLicensingCotermLicenses struct {
	ClaimedAt      string                                                                 `json:"claimedAt,omitempty"`      // When the license was claimed into the organization
	Counts         *[]ResponseItemLicensingGetOrganizationLicensingCotermLicensesCounts   `json:"counts,omitempty"`         // The counts of the license by model type
	Duration       *int                                                                   `json:"duration,omitempty"`       // The duration (term length) of the license, measured in days
	Editions       *[]ResponseItemLicensingGetOrganizationLicensingCotermLicensesEditions `json:"editions,omitempty"`       // The editions of the license for each relevant product type
	Expired        *bool                                                                  `json:"expired,omitempty"`        // Flag to indicate if the license is expired
	Invalidated    *bool                                                                  `json:"invalidated,omitempty"`    // Flag to indicated that the license is invalidated
	InvalidatedAt  string                                                                 `json:"invalidatedAt,omitempty"`  // When the license was invalidated. Will be null for active licenses
	Key            string                                                                 `json:"key,omitempty"`            // The key of the license
	Mode           string                                                                 `json:"mode,omitempty"`           // The operation mode of the license when it was claimed
	OrganizationID string                                                                 `json:"organizationId,omitempty"` // The ID of the organization that the license is claimed in
	StartedAt      string                                                                 `json:"startedAt,omitempty"`      // When the license's term began (approximately the date when the license was created)
}
type ResponseItemLicensingGetOrganizationLicensingCotermLicensesCounts struct {
	Count *int   `json:"count,omitempty"` // The number of counts the license contains of this model
	Model string `json:"model,omitempty"` // The license model type
}
type ResponseItemLicensingGetOrganizationLicensingCotermLicensesEditions struct {
	Edition     string `json:"edition,omitempty"`     // The name of the license edition
	ProductType string `json:"productType,omitempty"` // The product type of the license edition
}
type ResponseLicensingMoveOrganizationLicensingCotermLicenses struct {
	MovedLicenses     *[]ResponseLicensingMoveOrganizationLicensingCotermLicensesMovedLicenses     `json:"movedLicenses,omitempty"`     // Newly moved licenses created in the destination organization of the license move operation
	RemainderLicenses *[]ResponseLicensingMoveOrganizationLicensingCotermLicensesRemainderLicenses `json:"remainderLicenses,omitempty"` // Remainder licenses created in the source organization as a result of moving a subset of the counts of a license
}
type ResponseLicensingMoveOrganizationLicensingCotermLicensesMovedLicenses struct {
	ClaimedAt      string                                                                           `json:"claimedAt,omitempty"`      // When the license was claimed into the organization
	Counts         *[]ResponseLicensingMoveOrganizationLicensingCotermLicensesMovedLicensesCounts   `json:"counts,omitempty"`         // The counts of the license by model type
	Duration       *int                                                                             `json:"duration,omitempty"`       // The duration (term length) of the license, measured in days
	Editions       *[]ResponseLicensingMoveOrganizationLicensingCotermLicensesMovedLicensesEditions `json:"editions,omitempty"`       // The editions of the license for each relevant product type
	Expired        *bool                                                                            `json:"expired,omitempty"`        // Flag to indicate if the license is expired
	Invalidated    *bool                                                                            `json:"invalidated,omitempty"`    // Flag to indicated that the license is invalidated
	InvalidatedAt  string                                                                           `json:"invalidatedAt,omitempty"`  // When the license was invalidated. Will be null for active licenses
	Key            string                                                                           `json:"key,omitempty"`            // The key of the license
	Mode           string                                                                           `json:"mode,omitempty"`           // The operation mode of the license when it was claimed
	OrganizationID string                                                                           `json:"organizationId,omitempty"` // The ID of the organization that the license is claimed in
	StartedAt      string                                                                           `json:"startedAt,omitempty"`      // When the license's term began (approximately the date when the license was created)
}
type ResponseLicensingMoveOrganizationLicensingCotermLicensesMovedLicensesCounts struct {
	Count *int   `json:"count,omitempty"` // The number of counts the license contains of this model
	Model string `json:"model,omitempty"` // The license model type
}
type ResponseLicensingMoveOrganizationLicensingCotermLicensesMovedLicensesEditions struct {
	Edition     string `json:"edition,omitempty"`     // The name of the license edition
	ProductType string `json:"productType,omitempty"` // The product type of the license edition
}
type ResponseLicensingMoveOrganizationLicensingCotermLicensesRemainderLicenses struct {
	ClaimedAt      string                                                                               `json:"claimedAt,omitempty"`      // When the license was claimed into the organization
	Counts         *[]ResponseLicensingMoveOrganizationLicensingCotermLicensesRemainderLicensesCounts   `json:"counts,omitempty"`         // The counts of the license by model type
	Duration       *int                                                                                 `json:"duration,omitempty"`       // The duration (term length) of the license, measured in days
	Editions       *[]ResponseLicensingMoveOrganizationLicensingCotermLicensesRemainderLicensesEditions `json:"editions,omitempty"`       // The editions of the license for each relevant product type
	Expired        *bool                                                                                `json:"expired,omitempty"`        // Flag to indicate if the license is expired
	Invalidated    *bool                                                                                `json:"invalidated,omitempty"`    // Flag to indicated that the license is invalidated
	InvalidatedAt  string                                                                               `json:"invalidatedAt,omitempty"`  // When the license was invalidated. Will be null for active licenses
	Key            string                                                                               `json:"key,omitempty"`            // The key of the license
	Mode           string                                                                               `json:"mode,omitempty"`           // The operation mode of the license when it was claimed
	OrganizationID string                                                                               `json:"organizationId,omitempty"` // The ID of the organization that the license is claimed in
	StartedAt      string                                                                               `json:"startedAt,omitempty"`      // When the license's term began (approximately the date when the license was created)
}
type ResponseLicensingMoveOrganizationLicensingCotermLicensesRemainderLicensesCounts struct {
	Count *int   `json:"count,omitempty"` // The number of counts the license contains of this model
	Model string `json:"model,omitempty"` // The license model type
}
type ResponseLicensingMoveOrganizationLicensingCotermLicensesRemainderLicensesEditions struct {
	Edition     string `json:"edition,omitempty"`     // The name of the license edition
	ProductType string `json:"productType,omitempty"` // The product type of the license edition
}
type RequestLicensingMoveOrganizationLicensingCotermLicenses struct {
	Destination *RequestLicensingMoveOrganizationLicensingCotermLicensesDestination `json:"destination,omitempty"` // Destination data for the license move
	Licenses    *[]RequestLicensingMoveOrganizationLicensingCotermLicensesLicenses  `json:"licenses,omitempty"`    // The list of licenses to move
}
type RequestLicensingMoveOrganizationLicensingCotermLicensesDestination struct {
	Mode           string `json:"mode,omitempty"`           // The claim mode of the moved license
	OrganizationID string `json:"organizationId,omitempty"` // The organization to move the license to
}
type RequestLicensingMoveOrganizationLicensingCotermLicensesLicenses struct {
	Counts *[]RequestLicensingMoveOrganizationLicensingCotermLicensesLicensesCounts `json:"counts,omitempty"` // The counts to move from the license by model type
	Key    string                                                                   `json:"key,omitempty"`    // The license key to move counts from
}
type RequestLicensingMoveOrganizationLicensingCotermLicensesLicensesCounts struct {
	Count *int   `json:"count,omitempty"` // The number of counts to move
	Model string `json:"model,omitempty"` // The license model type to move counts of
}

//GetOrganizationLicensingCotermLicenses List the licenses in a coterm organization
/* List the licenses in a coterm organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationLicensingCotermLicensesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-licensing-coterm-licenses
*/
func (s *LicensingService) GetOrganizationLicensingCotermLicenses(organizationID string, getOrganizationLicensingCotermLicensesQueryParams *GetOrganizationLicensingCotermLicensesQueryParams) (*ResponseLicensingGetOrganizationLicensingCotermLicenses, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/licensing/coterm/licenses"
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationLicensingCotermLicensesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseLicensingGetOrganizationLicensingCotermLicenses{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationLicensingCotermLicenses")
	}

	result := response.Result().(*ResponseLicensingGetOrganizationLicensingCotermLicenses)
	return result, response, err

}

//MoveOrganizationLicensingCotermLicenses Moves a license to a different organization (coterm only)
/* Moves a license to a different organization (coterm only)

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!move-organization-licensing-coterm-licenses
*/

func (s *LicensingService) MoveOrganizationLicensingCotermLicenses(organizationID string, requestLicensingMoveOrganizationLicensingCotermLicenses *RequestLicensingMoveOrganizationLicensingCotermLicenses) (*ResponseLicensingMoveOrganizationLicensingCotermLicenses, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/licensing/coterm/licenses/move"
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestLicensingMoveOrganizationLicensingCotermLicenses).
		SetResult(&ResponseLicensingMoveOrganizationLicensingCotermLicenses{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation MoveOrganizationLicensingCotermLicenses")
	}

	result := response.Result().(*ResponseLicensingMoveOrganizationLicensingCotermLicenses)
	return result, response, err

}
