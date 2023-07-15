# GetNetworkWirelessBilling200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | The currency code of this node group&#39;s billing plans | [optional] 
**Plans** | Pointer to [**[]GetNetworkWirelessBilling200ResponsePlansInner**](GetNetworkWirelessBilling200ResponsePlansInner.md) | Array of billing plans in the node group. (Can configure a maximum of 5) | [optional] 

## Methods

### NewGetNetworkWirelessBilling200Response

`func NewGetNetworkWirelessBilling200Response() *GetNetworkWirelessBilling200Response`

NewGetNetworkWirelessBilling200Response instantiates a new GetNetworkWirelessBilling200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkWirelessBilling200ResponseWithDefaults

`func NewGetNetworkWirelessBilling200ResponseWithDefaults() *GetNetworkWirelessBilling200Response`

NewGetNetworkWirelessBilling200ResponseWithDefaults instantiates a new GetNetworkWirelessBilling200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *GetNetworkWirelessBilling200Response) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetNetworkWirelessBilling200Response) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetNetworkWirelessBilling200Response) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GetNetworkWirelessBilling200Response) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPlans

`func (o *GetNetworkWirelessBilling200Response) GetPlans() []GetNetworkWirelessBilling200ResponsePlansInner`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *GetNetworkWirelessBilling200Response) GetPlansOk() (*[]GetNetworkWirelessBilling200ResponsePlansInner, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *GetNetworkWirelessBilling200Response) SetPlans(v []GetNetworkWirelessBilling200ResponsePlansInner)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *GetNetworkWirelessBilling200Response) HasPlans() bool`

HasPlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


