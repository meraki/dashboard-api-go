# ClaimIntoOrganizationInventoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orders** | Pointer to **[]string** | The numbers of the orders that should be claimed | [optional] 
**Serials** | Pointer to **[]string** | The serials of the devices that should be claimed | [optional] 
**Licenses** | Pointer to [**[]ClaimIntoOrganizationInventoryRequestLicensesInner**](ClaimIntoOrganizationInventoryRequestLicensesInner.md) | The licenses that should be claimed | [optional] 

## Methods

### NewClaimIntoOrganizationInventoryRequest

`func NewClaimIntoOrganizationInventoryRequest() *ClaimIntoOrganizationInventoryRequest`

NewClaimIntoOrganizationInventoryRequest instantiates a new ClaimIntoOrganizationInventoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClaimIntoOrganizationInventoryRequestWithDefaults

`func NewClaimIntoOrganizationInventoryRequestWithDefaults() *ClaimIntoOrganizationInventoryRequest`

NewClaimIntoOrganizationInventoryRequestWithDefaults instantiates a new ClaimIntoOrganizationInventoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrders

`func (o *ClaimIntoOrganizationInventoryRequest) GetOrders() []string`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *ClaimIntoOrganizationInventoryRequest) GetOrdersOk() (*[]string, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *ClaimIntoOrganizationInventoryRequest) SetOrders(v []string)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *ClaimIntoOrganizationInventoryRequest) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetSerials

`func (o *ClaimIntoOrganizationInventoryRequest) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *ClaimIntoOrganizationInventoryRequest) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *ClaimIntoOrganizationInventoryRequest) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *ClaimIntoOrganizationInventoryRequest) HasSerials() bool`

HasSerials returns a boolean if a field has been set.

### GetLicenses

`func (o *ClaimIntoOrganizationInventoryRequest) GetLicenses() []ClaimIntoOrganizationInventoryRequestLicensesInner`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *ClaimIntoOrganizationInventoryRequest) GetLicensesOk() (*[]ClaimIntoOrganizationInventoryRequestLicensesInner, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *ClaimIntoOrganizationInventoryRequest) SetLicenses(v []ClaimIntoOrganizationInventoryRequestLicensesInner)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *ClaimIntoOrganizationInventoryRequest) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


