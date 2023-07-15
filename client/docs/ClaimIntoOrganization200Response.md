# ClaimIntoOrganization200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orders** | Pointer to **[]string** | The numbers of the orders claimed | [optional] 
**Serials** | Pointer to **[]string** | The serials of the devices claimed | [optional] 
**Licenses** | Pointer to [**[]ClaimIntoOrganization200ResponseLicensesInner**](ClaimIntoOrganization200ResponseLicensesInner.md) | The licenses claimed | [optional] 

## Methods

### NewClaimIntoOrganization200Response

`func NewClaimIntoOrganization200Response() *ClaimIntoOrganization200Response`

NewClaimIntoOrganization200Response instantiates a new ClaimIntoOrganization200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClaimIntoOrganization200ResponseWithDefaults

`func NewClaimIntoOrganization200ResponseWithDefaults() *ClaimIntoOrganization200Response`

NewClaimIntoOrganization200ResponseWithDefaults instantiates a new ClaimIntoOrganization200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrders

`func (o *ClaimIntoOrganization200Response) GetOrders() []string`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *ClaimIntoOrganization200Response) GetOrdersOk() (*[]string, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *ClaimIntoOrganization200Response) SetOrders(v []string)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *ClaimIntoOrganization200Response) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetSerials

`func (o *ClaimIntoOrganization200Response) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *ClaimIntoOrganization200Response) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *ClaimIntoOrganization200Response) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *ClaimIntoOrganization200Response) HasSerials() bool`

HasSerials returns a boolean if a field has been set.

### GetLicenses

`func (o *ClaimIntoOrganization200Response) GetLicenses() []ClaimIntoOrganization200ResponseLicensesInner`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *ClaimIntoOrganization200Response) GetLicensesOk() (*[]ClaimIntoOrganization200ResponseLicensesInner, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *ClaimIntoOrganization200Response) SetLicenses(v []ClaimIntoOrganization200ResponseLicensesInner)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *ClaimIntoOrganization200Response) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


