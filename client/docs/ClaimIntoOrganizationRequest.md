# ClaimIntoOrganizationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orders** | Pointer to **[]string** | The numbers of the orders that should be claimed | [optional] 
**Serials** | Pointer to **[]string** | The serials of the devices that should be claimed | [optional] 
**Licenses** | Pointer to [**[]ClaimIntoOrganizationRequestLicensesInner**](ClaimIntoOrganizationRequestLicensesInner.md) | The licenses that should be claimed | [optional] 

## Methods

### NewClaimIntoOrganizationRequest

`func NewClaimIntoOrganizationRequest() *ClaimIntoOrganizationRequest`

NewClaimIntoOrganizationRequest instantiates a new ClaimIntoOrganizationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClaimIntoOrganizationRequestWithDefaults

`func NewClaimIntoOrganizationRequestWithDefaults() *ClaimIntoOrganizationRequest`

NewClaimIntoOrganizationRequestWithDefaults instantiates a new ClaimIntoOrganizationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrders

`func (o *ClaimIntoOrganizationRequest) GetOrders() []string`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *ClaimIntoOrganizationRequest) GetOrdersOk() (*[]string, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *ClaimIntoOrganizationRequest) SetOrders(v []string)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *ClaimIntoOrganizationRequest) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetSerials

`func (o *ClaimIntoOrganizationRequest) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *ClaimIntoOrganizationRequest) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *ClaimIntoOrganizationRequest) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *ClaimIntoOrganizationRequest) HasSerials() bool`

HasSerials returns a boolean if a field has been set.

### GetLicenses

`func (o *ClaimIntoOrganizationRequest) GetLicenses() []ClaimIntoOrganizationRequestLicensesInner`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *ClaimIntoOrganizationRequest) GetLicensesOk() (*[]ClaimIntoOrganizationRequestLicensesInner, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *ClaimIntoOrganizationRequest) SetLicenses(v []ClaimIntoOrganizationRequestLicensesInner)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *ClaimIntoOrganizationRequest) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


