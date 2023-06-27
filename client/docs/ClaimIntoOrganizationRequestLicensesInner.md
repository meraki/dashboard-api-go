# ClaimIntoOrganizationRequestLicensesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The key of the license | 
**Mode** | Pointer to **string** | Either &#39;renew&#39; or &#39;addDevices&#39;. &#39;addDevices&#39; will increase the license limit, while &#39;renew&#39; will extend the amount of time until expiration. Defaults to &#39;addDevices&#39;. All licenses must be claimed with the same mode, and at most one renewal can be claimed at a time. This parameter is legacy and does not apply to organizations with per-device licensing enabled. | [optional] 

## Methods

### NewClaimIntoOrganizationRequestLicensesInner

`func NewClaimIntoOrganizationRequestLicensesInner(key string, ) *ClaimIntoOrganizationRequestLicensesInner`

NewClaimIntoOrganizationRequestLicensesInner instantiates a new ClaimIntoOrganizationRequestLicensesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClaimIntoOrganizationRequestLicensesInnerWithDefaults

`func NewClaimIntoOrganizationRequestLicensesInnerWithDefaults() *ClaimIntoOrganizationRequestLicensesInner`

NewClaimIntoOrganizationRequestLicensesInnerWithDefaults instantiates a new ClaimIntoOrganizationRequestLicensesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ClaimIntoOrganizationRequestLicensesInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ClaimIntoOrganizationRequestLicensesInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ClaimIntoOrganizationRequestLicensesInner) SetKey(v string)`

SetKey sets Key field to given value.


### GetMode

`func (o *ClaimIntoOrganizationRequestLicensesInner) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ClaimIntoOrganizationRequestLicensesInner) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ClaimIntoOrganizationRequestLicensesInner) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ClaimIntoOrganizationRequestLicensesInner) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


