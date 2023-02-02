# NetworksNetworkIdWirelessBillingPlans

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the pricing plan to update. | [optional] 
**Price** | **float32** | The price of the billing plan. | 
**BandwidthLimits** | [**NetworksNetworkIdWirelessBillingBandwidthLimits**](NetworksNetworkIdWirelessBillingBandwidthLimits.md) |  | 
**TimeLimit** | **string** | The time limit of the pricing plan in minutes. Can be &#39;1 hour&#39;, &#39;1 day&#39;, &#39;1 week&#39;, or &#39;30 days&#39;. | 

## Methods

### NewNetworksNetworkIdWirelessBillingPlans

`func NewNetworksNetworkIdWirelessBillingPlans(price float32, bandwidthLimits NetworksNetworkIdWirelessBillingBandwidthLimits, timeLimit string, ) *NetworksNetworkIdWirelessBillingPlans`

NewNetworksNetworkIdWirelessBillingPlans instantiates a new NetworksNetworkIdWirelessBillingPlans object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdWirelessBillingPlansWithDefaults

`func NewNetworksNetworkIdWirelessBillingPlansWithDefaults() *NetworksNetworkIdWirelessBillingPlans`

NewNetworksNetworkIdWirelessBillingPlansWithDefaults instantiates a new NetworksNetworkIdWirelessBillingPlans object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworksNetworkIdWirelessBillingPlans) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworksNetworkIdWirelessBillingPlans) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworksNetworkIdWirelessBillingPlans) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NetworksNetworkIdWirelessBillingPlans) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPrice

`func (o *NetworksNetworkIdWirelessBillingPlans) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *NetworksNetworkIdWirelessBillingPlans) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *NetworksNetworkIdWirelessBillingPlans) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetBandwidthLimits

`func (o *NetworksNetworkIdWirelessBillingPlans) GetBandwidthLimits() NetworksNetworkIdWirelessBillingBandwidthLimits`

GetBandwidthLimits returns the BandwidthLimits field if non-nil, zero value otherwise.

### GetBandwidthLimitsOk

`func (o *NetworksNetworkIdWirelessBillingPlans) GetBandwidthLimitsOk() (*NetworksNetworkIdWirelessBillingBandwidthLimits, bool)`

GetBandwidthLimitsOk returns a tuple with the BandwidthLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthLimits

`func (o *NetworksNetworkIdWirelessBillingPlans) SetBandwidthLimits(v NetworksNetworkIdWirelessBillingBandwidthLimits)`

SetBandwidthLimits sets BandwidthLimits field to given value.


### GetTimeLimit

`func (o *NetworksNetworkIdWirelessBillingPlans) GetTimeLimit() string`

GetTimeLimit returns the TimeLimit field if non-nil, zero value otherwise.

### GetTimeLimitOk

`func (o *NetworksNetworkIdWirelessBillingPlans) GetTimeLimitOk() (*string, bool)`

GetTimeLimitOk returns a tuple with the TimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeLimit

`func (o *NetworksNetworkIdWirelessBillingPlans) SetTimeLimit(v string)`

SetTimeLimit sets TimeLimit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


