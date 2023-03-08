# OrganizationsOrganizationIdInventoryClaimLicenses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The key of the license | 
**Mode** | Pointer to **string** | Co-term licensing only: either &#39;renew&#39; or &#39;addDevices&#39;. &#39;addDevices&#39; will increase the license limit, while &#39;renew&#39; will extend the amount of time until expiration. Defaults to &#39;addDevices&#39;. All licenses must be claimed with the same mode, and at most one renewal can be claimed at a time. Does not apply to organizations using per-device licensing model. | [optional] 

## Methods

### NewOrganizationsOrganizationIdInventoryClaimLicenses

`func NewOrganizationsOrganizationIdInventoryClaimLicenses(key string, ) *OrganizationsOrganizationIdInventoryClaimLicenses`

NewOrganizationsOrganizationIdInventoryClaimLicenses instantiates a new OrganizationsOrganizationIdInventoryClaimLicenses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsOrganizationIdInventoryClaimLicensesWithDefaults

`func NewOrganizationsOrganizationIdInventoryClaimLicensesWithDefaults() *OrganizationsOrganizationIdInventoryClaimLicenses`

NewOrganizationsOrganizationIdInventoryClaimLicensesWithDefaults instantiates a new OrganizationsOrganizationIdInventoryClaimLicenses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *OrganizationsOrganizationIdInventoryClaimLicenses) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *OrganizationsOrganizationIdInventoryClaimLicenses) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *OrganizationsOrganizationIdInventoryClaimLicenses) SetKey(v string)`

SetKey sets Key field to given value.


### GetMode

`func (o *OrganizationsOrganizationIdInventoryClaimLicenses) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *OrganizationsOrganizationIdInventoryClaimLicenses) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *OrganizationsOrganizationIdInventoryClaimLicenses) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *OrganizationsOrganizationIdInventoryClaimLicenses) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


