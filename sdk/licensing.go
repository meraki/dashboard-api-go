package meraki

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type LicensingService service

type GetAdministeredLicensingSubscriptionEntitlementsQueryParams struct {
	Skus []string `url:"skus[],omitempty"` //Filter to entitlements with the specified SKUs
}
type GetAdministeredLicensingSubscriptionSubscriptionsQueryParams struct {
	PerPage         int      `url:"perPage,omitempty"`           //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter   string   `url:"startingAfter,omitempty"`     //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore    string   `url:"endingBefore,omitempty"`      //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	SubscriptionIDs []string `url:"subscriptionIds[],omitempty"` //List of subscription ids to fetch
	OrganizationIDs []string `url:"organizationIds[],omitempty"` //Organizations to get associated subscriptions for
	Statuses        []string `url:"statuses[],omitempty"`        //List of statuses that returned subscriptions can have
	ProductTypes    []string `url:"productTypes[],omitempty"`    //List of product types that returned subscriptions need to have entitlements for.
	StartDate       string   `url:"startDate,omitempty"`         //Filter subscriptions by start date, ISO 8601 format. To filter with a range of dates, use 'startDate[<option>]=?' in the request. Accepted options include lt, gt, lte, gte.
	EndDate         string   `url:"endDate,omitempty"`           //Filter subscriptions by end date, ISO 8601 format. To filter with a range of dates, use 'endDate[<option>]=?' in the request. Accepted options include lt, gt, lte, gte.
}
type ClaimAdministeredLicensingSubscriptionSubscriptionsQueryParams struct {
	Validate bool `url:"validate,omitempty"` //Check if the provided claim key is valid and can be claimed into the organization.
}
type GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesQueryParams struct {
	OrganizationIDs []string `url:"organizationIds[],omitempty"` //Organizations to get subscription compliance information for
	SubscriptionIDs []string `url:"subscriptionIds[],omitempty"` //Subscription ids
}
type BindAdministeredLicensingSubscriptionSubscriptionQueryParams struct {
	Validate bool `url:"validate,omitempty"` //Check if the provided networks can be bound to the subscription. Returns any licensing problems and does not commit the results.
}
type GetOrganizationLicensingCotermLicensesQueryParams struct {
	PerPage       int    `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter string `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	Invalidated   bool   `url:"invalidated,omitempty"`   //Filter for licenses that are invalidated
	Expired       bool   `url:"expired,omitempty"`       //Filter for licenses that are expired
}

type ResponseLicensingGetAdministeredLicensingSubscriptionEntitlements struct {
	FeatureTier  string `json:"featureTier,omitempty"`  // The feature tier associated with the entitlement (null for add-ons)
	IsAddOn      *bool  `json:"isAddOn,omitempty"`      // Whether or not the entitlement is an add-on
	Name         string `json:"name,omitempty"`         // The user-facing name of the entitlement
	ProductClass string `json:"productClass,omitempty"` // The product class associated with the entitlement
	ProductType  string `json:"productType,omitempty"`  // The product type of the entitlement
	Sku          string `json:"sku,omitempty"`          // The SKU identifier of the entitlement
}
type ResponseLicensingGetAdministeredLicensingSubscriptionSubscriptions []ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptions // Array of ResponseLicensingGetAdministeredLicensingSubscriptionSubscriptions
type ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptions struct {
	Counts         *ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsCounts         `json:"counts,omitempty"`         // Numeric breakdown of network and entitlement counts
	Description    string                                                                                `json:"description,omitempty"`    // Subscription description
	EndDate        string                                                                                `json:"endDate,omitempty"`        // Subscription expiration date
	Entitlements   *[]ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsEntitlements `json:"entitlements,omitempty"`   // Entitlement info
	Name           string                                                                                `json:"name,omitempty"`           // Subscription name
	ProductTypes   []string                                                                              `json:"productTypes,omitempty"`   // Products the subscription has entitlements for
	StartDate      string                                                                                `json:"startDate,omitempty"`      // Subscription start date
	Status         string                                                                                `json:"status,omitempty"`         // Subscription status
	SubscriptionID string                                                                                `json:"subscriptionId,omitempty"` // Subscription's ID
	WebOrderID     string                                                                                `json:"webOrderId,omitempty"`     // Web order id
}
type ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsCounts struct {
	Networks *int                                                                               `json:"networks,omitempty"` // Number of networks bound to this subscription
	Seats    *ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsCountsSeats `json:"seats,omitempty"`    // Seat distribution
}
type ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsCountsSeats struct {
	Assigned  *int `json:"assigned,omitempty"`  // Number of seats in use
	Available *int `json:"available,omitempty"` // Number of seats available for use
	Limit     *int `json:"limit,omitempty"`     // Total number of seats provided by this subscription
}
type ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsEntitlements struct {
	Seats *ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsEntitlementsSeats `json:"seats,omitempty"` // Seat distribution
	Sku   string                                                                                   `json:"sku,omitempty"`   // SKU of the required product
}
type ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsEntitlementsSeats struct {
	Assigned  *int `json:"assigned,omitempty"`  // Number of seats in use
	Available *int `json:"available,omitempty"` // Number of seats available for use
	Limit     *int `json:"limit,omitempty"`     // Total number of seats provided by this subscription for this sku
}
type ResponseLicensingClaimAdministeredLicensingSubscriptionSubscriptions struct {
	Counts         *ResponseLicensingClaimAdministeredLicensingSubscriptionSubscriptionsCounts         `json:"counts,omitempty"`         // Numeric breakdown of network and entitlement counts
	Description    string                                                                              `json:"description,omitempty"`    // Subscription description
	EndDate        string                                                                              `json:"endDate,omitempty"`        // Subscription expiration date
	Entitlements   *[]ResponseLicensingClaimAdministeredLicensingSubscriptionSubscriptionsEntitlements `json:"entitlements,omitempty"`   // Entitlement info
	Name           string                                                                              `json:"name,omitempty"`           // Subscription name
	ProductTypes   []string                                                                            `json:"productTypes,omitempty"`   // Products the subscription has entitlements for
	StartDate      string                                                                              `json:"startDate,omitempty"`      // Subscription start date
	Status         string                                                                              `json:"status,omitempty"`         // Subscription status
	SubscriptionID string                                                                              `json:"subscriptionId,omitempty"` // Subscription's ID
	WebOrderID     string                                                                              `json:"webOrderId,omitempty"`     // Web order id
}
type ResponseLicensingClaimAdministeredLicensingSubscriptionSubscriptionsCounts struct {
	Networks *int                                                                             `json:"networks,omitempty"` // Number of networks bound to this subscription
	Seats    *ResponseLicensingClaimAdministeredLicensingSubscriptionSubscriptionsCountsSeats `json:"seats,omitempty"`    // Seat distribution
}
type ResponseLicensingClaimAdministeredLicensingSubscriptionSubscriptionsCountsSeats struct {
	Assigned  *int `json:"assigned,omitempty"`  // Number of seats in use
	Available *int `json:"available,omitempty"` // Number of seats available for use
	Limit     *int `json:"limit,omitempty"`     // Total number of seats provided by this subscription
}
type ResponseLicensingClaimAdministeredLicensingSubscriptionSubscriptionsEntitlements struct {
	Seats *ResponseLicensingClaimAdministeredLicensingSubscriptionSubscriptionsEntitlementsSeats `json:"seats,omitempty"` // Seat distribution
	Sku   string                                                                                 `json:"sku,omitempty"`   // SKU of the required product
}
type ResponseLicensingClaimAdministeredLicensingSubscriptionSubscriptionsEntitlementsSeats struct {
	Assigned  *int `json:"assigned,omitempty"`  // Number of seats in use
	Available *int `json:"available,omitempty"` // Number of seats available for use
	Limit     *int `json:"limit,omitempty"`     // Total number of seats provided by this subscription for this sku
}
type ResponseLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey struct {
	Counts         *ResponseLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyCounts         `json:"counts,omitempty"`         // Numeric breakdown of network and entitlement counts
	Description    string                                                                                         `json:"description,omitempty"`    // Subscription description
	EndDate        string                                                                                         `json:"endDate,omitempty"`        // Subscription expiration date
	Entitlements   *[]ResponseLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyEntitlements `json:"entitlements,omitempty"`   // Entitlement info
	Name           string                                                                                         `json:"name,omitempty"`           // Subscription name
	ProductTypes   []string                                                                                       `json:"productTypes,omitempty"`   // Products the subscription has entitlements for
	StartDate      string                                                                                         `json:"startDate,omitempty"`      // Subscription start date
	Status         string                                                                                         `json:"status,omitempty"`         // Subscription status
	SubscriptionID string                                                                                         `json:"subscriptionId,omitempty"` // Subscription's ID
	WebOrderID     string                                                                                         `json:"webOrderId,omitempty"`     // Web order id
}
type ResponseLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyCounts struct {
	Networks *int                                                                                        `json:"networks,omitempty"` // Number of networks bound to this subscription
	Seats    *ResponseLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyCountsSeats `json:"seats,omitempty"`    // Seat distribution
}
type ResponseLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyCountsSeats struct {
	Assigned  *int `json:"assigned,omitempty"`  // Number of seats in use
	Available *int `json:"available,omitempty"` // Number of seats available for use
	Limit     *int `json:"limit,omitempty"`     // Total number of seats provided by this subscription
}
type ResponseLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyEntitlements struct {
	Seats *ResponseLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyEntitlementsSeats `json:"seats,omitempty"` // Seat distribution
	Sku   string                                                                                            `json:"sku,omitempty"`   // SKU of the required product
}
type ResponseLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyEntitlementsSeats struct {
	Assigned  *int `json:"assigned,omitempty"`  // Number of seats in use
	Available *int `json:"available,omitempty"` // Number of seats available for use
	Limit     *int `json:"limit,omitempty"`     // Total number of seats provided by this subscription for this sku
}
type ResponseLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses []ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses // Array of ResponseLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses
type ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses struct {
	Subscription *ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription `json:"subscription,omitempty"` // Subscription details
	Violations   *ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolations   `json:"violations,omitempty"`   // Violations
}
type ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription struct {
	ID     string `json:"id,omitempty"`     // Subscription's ID
	Name   string `json:"name,omitempty"`   // Friendly name to identify the subscription
	Status string `json:"status,omitempty"` // One of the following: "inactive" | "active" | "out_of_compliance" | "expired" | "canceled"
}
type ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolations struct {
	ByProductClass *[]ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass `json:"byProductClass,omitempty"` // List of violations by product class that are not compliance
}
type ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass struct {
	GracePeriodEndsAt string                                                                                                                   `json:"gracePeriodEndsAt,omitempty"` // End date of the grace period in ISO 8601 format
	Missing           *ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClassMissing `json:"missing,omitempty"`           // Missing entitlements details
	ProductClass      string                                                                                                                   `json:"productClass,omitempty"`      // Name of the product class
}
type ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClassMissing struct {
	Entitlements *[]ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClassMissingEntitlements `json:"entitlements,omitempty"` // List of missing entitlements
}
type ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClassMissingEntitlements struct {
	Quantity *int   `json:"quantity,omitempty"` // Number required
	Sku      string `json:"sku,omitempty"`      // SKU of the required product
}
type ResponseLicensingBindAdministeredLicensingSubscriptionSubscription struct {
	Errors                   []string                                                                                      `json:"errors,omitempty"`                   // Array of errors if failed
	InsufficientEntitlements *[]ResponseLicensingBindAdministeredLicensingSubscriptionSubscriptionInsufficientEntitlements `json:"insufficientEntitlements,omitempty"` // A list of entitlements required to successfully bind the networks to the subscription
	Networks                 *[]ResponseLicensingBindAdministeredLicensingSubscriptionSubscriptionNetworks                 `json:"networks,omitempty"`                 // Unbound networks
	SubscriptionID           string                                                                                        `json:"subscriptionId,omitempty"`           // Subscription ID
}
type ResponseLicensingBindAdministeredLicensingSubscriptionSubscriptionInsufficientEntitlements struct {
	Quantity *int   `json:"quantity,omitempty"` // Number required
	Sku      string `json:"sku,omitempty"`      // SKU of the required product
}
type ResponseLicensingBindAdministeredLicensingSubscriptionSubscriptionNetworks struct {
	ID   string `json:"id,omitempty"`   // Network ID
	Name string `json:"name,omitempty"` // Network name
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
type RequestLicensingClaimAdministeredLicensingSubscriptionSubscriptions struct {
	ClaimKey       string `json:"claimKey,omitempty"`       // The subscription's claim key
	Description    string `json:"description,omitempty"`    // Extra details or notes about the subscription
	Name           string `json:"name,omitempty"`           // Friendly name to identify the subscription
	OrganizationID string `json:"organizationId,omitempty"` // The id of the organization claiming the subscription
}
type RequestLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey struct {
	ClaimKey string `json:"claimKey,omitempty"` // The subscription's claim key
}
type RequestLicensingBindAdministeredLicensingSubscriptionSubscription struct {
	NetworkIDs []string `json:"networkIds,omitempty"` // List of network ids to bind to the subscription
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

//GetAdministeredLicensingSubscriptionEntitlements Retrieve the list of purchasable entitlements
/* Retrieve the list of purchasable entitlements

@param getAdministeredLicensingSubscriptionEntitlementsQueryParams Filtering parameter


*/
func (s *LicensingService) GetAdministeredLicensingSubscriptionEntitlements(getAdministeredLicensingSubscriptionEntitlementsQueryParams *GetAdministeredLicensingSubscriptionEntitlementsQueryParams) (*ResponseLicensingGetAdministeredLicensingSubscriptionEntitlements, *resty.Response, error) {
	path := "/api/v1/administered/licensing/subscription/entitlements"
	s.rateLimiterBucket.Wait(1)

	queryString, _ := query.Values(getAdministeredLicensingSubscriptionEntitlementsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseLicensingGetAdministeredLicensingSubscriptionEntitlements{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetAdministeredLicensingSubscriptionEntitlements")
	}

	result := response.Result().(*ResponseLicensingGetAdministeredLicensingSubscriptionEntitlements)
	return result, response, err

}

//GetAdministeredLicensingSubscriptionSubscriptions List available subscriptions
/* List available subscriptions

@param getAdministeredLicensingSubscriptionSubscriptionsQueryParams Filtering parameter


*/
func (s *LicensingService) GetAdministeredLicensingSubscriptionSubscriptions(getAdministeredLicensingSubscriptionSubscriptionsQueryParams *GetAdministeredLicensingSubscriptionSubscriptionsQueryParams) (*ResponseLicensingGetAdministeredLicensingSubscriptionSubscriptions, *resty.Response, error) {
	path := "/api/v1/administered/licensing/subscription/subscriptions"
	s.rateLimiterBucket.Wait(1)

	queryString, _ := query.Values(getAdministeredLicensingSubscriptionSubscriptionsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseLicensingGetAdministeredLicensingSubscriptionSubscriptions{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetAdministeredLicensingSubscriptionSubscriptions")
	}

	result := response.Result().(*ResponseLicensingGetAdministeredLicensingSubscriptionSubscriptions)
	return result, response, err

}

//GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses Get compliance status for requested subscriptions
/* Get compliance status for requested subscriptions

@param getAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesQueryParams Filtering parameter


*/
func (s *LicensingService) GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses(getAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesQueryParams *GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesQueryParams) (*ResponseLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses, *resty.Response, error) {
	path := "/api/v1/administered/licensing/subscription/subscriptions/compliance/statuses"
	s.rateLimiterBucket.Wait(1)

	queryString, _ := query.Values(getAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses")
	}

	result := response.Result().(*ResponseLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses)
	return result, response, err

}

//GetOrganizationLicensingCotermLicenses List the licenses in a coterm organization
/* List the licenses in a coterm organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationLicensingCotermLicensesQueryParams Filtering parameter


*/
func (s *LicensingService) GetOrganizationLicensingCotermLicenses(organizationID string, getOrganizationLicensingCotermLicensesQueryParams *GetOrganizationLicensingCotermLicensesQueryParams) (*ResponseLicensingGetOrganizationLicensingCotermLicenses, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/licensing/coterm/licenses"
	s.rateLimiterBucket.Wait(1)
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

//ClaimAdministeredLicensingSubscriptionSubscriptions Claim a subscription into an organization.
/* Claim a subscription into an organization.

@param claimAdministeredLicensingSubscriptionSubscriptionsQueryParams Filtering parameter


*/

func (s *LicensingService) ClaimAdministeredLicensingSubscriptionSubscriptions(requestLicensingClaimAdministeredLicensingSubscriptionSubscriptions *RequestLicensingClaimAdministeredLicensingSubscriptionSubscriptions, claimAdministeredLicensingSubscriptionSubscriptionsQueryParams *ClaimAdministeredLicensingSubscriptionSubscriptionsQueryParams) (*ResponseLicensingClaimAdministeredLicensingSubscriptionSubscriptions, *resty.Response, error) {
	path := "/api/v1/administered/licensing/subscription/subscriptions/claim"
	s.rateLimiterBucket.Wait(1)

	queryString, _ := query.Values(claimAdministeredLicensingSubscriptionSubscriptionsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetBody(requestLicensingClaimAdministeredLicensingSubscriptionSubscriptions).
		SetResult(&ResponseLicensingClaimAdministeredLicensingSubscriptionSubscriptions{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ClaimAdministeredLicensingSubscriptionSubscriptions")
	}

	result := response.Result().(*ResponseLicensingClaimAdministeredLicensingSubscriptionSubscriptions)
	return result, response, err

}

//ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey Find a subscription by claim key
/* Find a subscription by claim key. Returns 400 if the key has already been claimed.



 */

func (s *LicensingService) ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey(requestLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey *RequestLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey) (*ResponseLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey, *resty.Response, error) {
	path := "/api/v1/administered/licensing/subscription/subscriptions/claimKey/validate"
	s.rateLimiterBucket.Wait(1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey).
		SetResult(&ResponseLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey")
	}

	result := response.Result().(*ResponseLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey)
	return result, response, err

}

//BindAdministeredLicensingSubscriptionSubscription Bind networks to a subscription
/* Bind networks to a subscription

@param subscriptionID subscriptionId path parameter. Subscription ID
@param bindAdministeredLicensingSubscriptionSubscriptionQueryParams Filtering parameter


*/

func (s *LicensingService) BindAdministeredLicensingSubscriptionSubscription(subscriptionID string, requestLicensingBindAdministeredLicensingSubscriptionSubscription *RequestLicensingBindAdministeredLicensingSubscriptionSubscription, bindAdministeredLicensingSubscriptionSubscriptionQueryParams *BindAdministeredLicensingSubscriptionSubscriptionQueryParams) (*ResponseLicensingBindAdministeredLicensingSubscriptionSubscription, *resty.Response, error) {
	path := "/api/v1/administered/licensing/subscription/subscriptions/{subscriptionId}/bind"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{subscriptionId}", fmt.Sprintf("%v", subscriptionID), -1)

	queryString, _ := query.Values(bindAdministeredLicensingSubscriptionSubscriptionQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetBody(requestLicensingBindAdministeredLicensingSubscriptionSubscription).
		SetResult(&ResponseLicensingBindAdministeredLicensingSubscriptionSubscription{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation BindAdministeredLicensingSubscriptionSubscription")
	}

	result := response.Result().(*ResponseLicensingBindAdministeredLicensingSubscriptionSubscription)
	return result, response, err

}

//MoveOrganizationLicensingCotermLicenses Moves a license to a different organization (coterm only)
/* Moves a license to a different organization (coterm only)

@param organizationID organizationId path parameter. Organization ID


*/

func (s *LicensingService) MoveOrganizationLicensingCotermLicenses(organizationID string, requestLicensingMoveOrganizationLicensingCotermLicenses *RequestLicensingMoveOrganizationLicensingCotermLicenses) (*ResponseLicensingMoveOrganizationLicensingCotermLicenses, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/licensing/coterm/licenses/move"
	s.rateLimiterBucket.Wait(1)
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
