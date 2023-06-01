# InlineObject190

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orders** | Pointer to **[]string** | The numbers of the orders that should be claimed | [optional] 
**Serials** | Pointer to **[]string** | The serials of the devices that should be claimed | [optional] 
**Licenses** | Pointer to [**[]OrganizationsOrganizationIdClaimLicenses**](OrganizationsOrganizationIdClaimLicenses.md) | The licenses that should be claimed | [optional] 

## Methods

### NewInlineObject190

`func NewInlineObject190() *InlineObject190`

NewInlineObject190 instantiates a new InlineObject190 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject190WithDefaults

`func NewInlineObject190WithDefaults() *InlineObject190`

NewInlineObject190WithDefaults instantiates a new InlineObject190 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrders

`func (o *InlineObject190) GetOrders() []string`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *InlineObject190) GetOrdersOk() (*[]string, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *InlineObject190) SetOrders(v []string)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *InlineObject190) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetSerials

`func (o *InlineObject190) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *InlineObject190) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *InlineObject190) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *InlineObject190) HasSerials() bool`

HasSerials returns a boolean if a field has been set.

### GetLicenses

`func (o *InlineObject190) GetLicenses() []OrganizationsOrganizationIdClaimLicenses`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *InlineObject190) GetLicensesOk() (*[]OrganizationsOrganizationIdClaimLicenses, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *InlineObject190) SetLicenses(v []OrganizationsOrganizationIdClaimLicenses)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *InlineObject190) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


