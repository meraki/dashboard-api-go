# GetNetworkWirelessSsidTrafficShapingRules200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrafficShapingEnabled** | Pointer to **bool** | Whether traffic shaping rules are applied to clients on your SSID. | [optional] 
**DefaultRulesEnabled** | Pointer to **bool** | Whether default traffic shaping rules are enabled (true) or disabled (false). There are 4 default rules, which can be seen on your network&#39;s traffic shaping page. Note that default rules count against the rule limit of 8. | [optional] 
**Rules** | Pointer to [**[]GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner**](GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner.md) |     An array of traffic shaping rules. Rules are applied in the order that     they are specified in. An empty list (or null) means no rules. Note that     you are allowed a maximum of 8 rules.  | [optional] 

## Methods

### NewGetNetworkWirelessSsidTrafficShapingRules200Response

`func NewGetNetworkWirelessSsidTrafficShapingRules200Response() *GetNetworkWirelessSsidTrafficShapingRules200Response`

NewGetNetworkWirelessSsidTrafficShapingRules200Response instantiates a new GetNetworkWirelessSsidTrafficShapingRules200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkWirelessSsidTrafficShapingRules200ResponseWithDefaults

`func NewGetNetworkWirelessSsidTrafficShapingRules200ResponseWithDefaults() *GetNetworkWirelessSsidTrafficShapingRules200Response`

NewGetNetworkWirelessSsidTrafficShapingRules200ResponseWithDefaults instantiates a new GetNetworkWirelessSsidTrafficShapingRules200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrafficShapingEnabled

`func (o *GetNetworkWirelessSsidTrafficShapingRules200Response) GetTrafficShapingEnabled() bool`

GetTrafficShapingEnabled returns the TrafficShapingEnabled field if non-nil, zero value otherwise.

### GetTrafficShapingEnabledOk

`func (o *GetNetworkWirelessSsidTrafficShapingRules200Response) GetTrafficShapingEnabledOk() (*bool, bool)`

GetTrafficShapingEnabledOk returns a tuple with the TrafficShapingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficShapingEnabled

`func (o *GetNetworkWirelessSsidTrafficShapingRules200Response) SetTrafficShapingEnabled(v bool)`

SetTrafficShapingEnabled sets TrafficShapingEnabled field to given value.

### HasTrafficShapingEnabled

`func (o *GetNetworkWirelessSsidTrafficShapingRules200Response) HasTrafficShapingEnabled() bool`

HasTrafficShapingEnabled returns a boolean if a field has been set.

### GetDefaultRulesEnabled

`func (o *GetNetworkWirelessSsidTrafficShapingRules200Response) GetDefaultRulesEnabled() bool`

GetDefaultRulesEnabled returns the DefaultRulesEnabled field if non-nil, zero value otherwise.

### GetDefaultRulesEnabledOk

`func (o *GetNetworkWirelessSsidTrafficShapingRules200Response) GetDefaultRulesEnabledOk() (*bool, bool)`

GetDefaultRulesEnabledOk returns a tuple with the DefaultRulesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRulesEnabled

`func (o *GetNetworkWirelessSsidTrafficShapingRules200Response) SetDefaultRulesEnabled(v bool)`

SetDefaultRulesEnabled sets DefaultRulesEnabled field to given value.

### HasDefaultRulesEnabled

`func (o *GetNetworkWirelessSsidTrafficShapingRules200Response) HasDefaultRulesEnabled() bool`

HasDefaultRulesEnabled returns a boolean if a field has been set.

### GetRules

`func (o *GetNetworkWirelessSsidTrafficShapingRules200Response) GetRules() []GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *GetNetworkWirelessSsidTrafficShapingRules200Response) GetRulesOk() (*[]GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *GetNetworkWirelessSsidTrafficShapingRules200Response) SetRules(v []GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner)`

SetRules sets Rules field to given value.

### HasRules

`func (o *GetNetworkWirelessSsidTrafficShapingRules200Response) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


