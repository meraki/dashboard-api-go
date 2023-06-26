# UpdateNetworkWirelessBillingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | The currency code of this node group&#39;s billing plans | [optional] 
**Plans** | Pointer to [**[]UpdateNetworkWirelessBillingRequestPlansInner**](UpdateNetworkWirelessBillingRequestPlansInner.md) | Array of billing plans in the node group. (Can configure a maximum of 5) | [optional] 

## Methods

### NewUpdateNetworkWirelessBillingRequest

`func NewUpdateNetworkWirelessBillingRequest() *UpdateNetworkWirelessBillingRequest`

NewUpdateNetworkWirelessBillingRequest instantiates a new UpdateNetworkWirelessBillingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkWirelessBillingRequestWithDefaults

`func NewUpdateNetworkWirelessBillingRequestWithDefaults() *UpdateNetworkWirelessBillingRequest`

NewUpdateNetworkWirelessBillingRequestWithDefaults instantiates a new UpdateNetworkWirelessBillingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *UpdateNetworkWirelessBillingRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UpdateNetworkWirelessBillingRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UpdateNetworkWirelessBillingRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UpdateNetworkWirelessBillingRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPlans

`func (o *UpdateNetworkWirelessBillingRequest) GetPlans() []UpdateNetworkWirelessBillingRequestPlansInner`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *UpdateNetworkWirelessBillingRequest) GetPlansOk() (*[]UpdateNetworkWirelessBillingRequestPlansInner, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *UpdateNetworkWirelessBillingRequest) SetPlans(v []UpdateNetworkWirelessBillingRequestPlansInner)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *UpdateNetworkWirelessBillingRequest) HasPlans() bool`

HasPlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


