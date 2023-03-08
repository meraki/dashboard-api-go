# InlineObject164

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrafficShapingEnabled** | Pointer to **bool** | Whether traffic shaping rules are applied to clients on your SSID. | [optional] 
**DefaultRulesEnabled** | Pointer to **bool** | Whether default traffic shaping rules are enabled (true) or disabled (false). There are 4 default rules, which can be seen on your network&#39;s traffic shaping page. Note that default rules count against the rule limit of 8. | [optional] 
**Rules** | Pointer to [**[]NetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules**](NetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules.md) |     An array of traffic shaping rules. Rules are applied in the order that     they are specified in. An empty list (or null) means no rules. Note that     you are allowed a maximum of 8 rules.  | [optional] 

## Methods

### NewInlineObject164

`func NewInlineObject164() *InlineObject164`

NewInlineObject164 instantiates a new InlineObject164 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject164WithDefaults

`func NewInlineObject164WithDefaults() *InlineObject164`

NewInlineObject164WithDefaults instantiates a new InlineObject164 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrafficShapingEnabled

`func (o *InlineObject164) GetTrafficShapingEnabled() bool`

GetTrafficShapingEnabled returns the TrafficShapingEnabled field if non-nil, zero value otherwise.

### GetTrafficShapingEnabledOk

`func (o *InlineObject164) GetTrafficShapingEnabledOk() (*bool, bool)`

GetTrafficShapingEnabledOk returns a tuple with the TrafficShapingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficShapingEnabled

`func (o *InlineObject164) SetTrafficShapingEnabled(v bool)`

SetTrafficShapingEnabled sets TrafficShapingEnabled field to given value.

### HasTrafficShapingEnabled

`func (o *InlineObject164) HasTrafficShapingEnabled() bool`

HasTrafficShapingEnabled returns a boolean if a field has been set.

### GetDefaultRulesEnabled

`func (o *InlineObject164) GetDefaultRulesEnabled() bool`

GetDefaultRulesEnabled returns the DefaultRulesEnabled field if non-nil, zero value otherwise.

### GetDefaultRulesEnabledOk

`func (o *InlineObject164) GetDefaultRulesEnabledOk() (*bool, bool)`

GetDefaultRulesEnabledOk returns a tuple with the DefaultRulesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRulesEnabled

`func (o *InlineObject164) SetDefaultRulesEnabled(v bool)`

SetDefaultRulesEnabled sets DefaultRulesEnabled field to given value.

### HasDefaultRulesEnabled

`func (o *InlineObject164) HasDefaultRulesEnabled() bool`

HasDefaultRulesEnabled returns a boolean if a field has been set.

### GetRules

`func (o *InlineObject164) GetRules() []NetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineObject164) GetRulesOk() (*[]NetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineObject164) SetRules(v []NetworksNetworkIdWirelessSsidsNumberTrafficShapingRulesRules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InlineObject164) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


