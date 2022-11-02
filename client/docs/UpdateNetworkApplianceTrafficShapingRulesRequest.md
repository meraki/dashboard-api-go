# UpdateNetworkApplianceTrafficShapingRulesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultRulesEnabled** | Pointer to **bool** | Whether default traffic shaping rules are enabled (true) or disabled (false). There are 4 default rules, which can be seen on your network&#39;s traffic shaping page. Note that default rules count against the rule limit of 8. | [optional] 
**Rules** | Pointer to [**[]UpdateNetworkApplianceTrafficShapingRulesRequestRulesInner**](UpdateNetworkApplianceTrafficShapingRulesRequestRulesInner.md) |     An array of traffic shaping rules. Rules are applied in the order that     they are specified in. An empty list (or null) means no rules. Note that     you are allowed a maximum of 8 rules.  | [optional] 

## Methods

### NewUpdateNetworkApplianceTrafficShapingRulesRequest

`func NewUpdateNetworkApplianceTrafficShapingRulesRequest() *UpdateNetworkApplianceTrafficShapingRulesRequest`

NewUpdateNetworkApplianceTrafficShapingRulesRequest instantiates a new UpdateNetworkApplianceTrafficShapingRulesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkApplianceTrafficShapingRulesRequestWithDefaults

`func NewUpdateNetworkApplianceTrafficShapingRulesRequestWithDefaults() *UpdateNetworkApplianceTrafficShapingRulesRequest`

NewUpdateNetworkApplianceTrafficShapingRulesRequestWithDefaults instantiates a new UpdateNetworkApplianceTrafficShapingRulesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultRulesEnabled

`func (o *UpdateNetworkApplianceTrafficShapingRulesRequest) GetDefaultRulesEnabled() bool`

GetDefaultRulesEnabled returns the DefaultRulesEnabled field if non-nil, zero value otherwise.

### GetDefaultRulesEnabledOk

`func (o *UpdateNetworkApplianceTrafficShapingRulesRequest) GetDefaultRulesEnabledOk() (*bool, bool)`

GetDefaultRulesEnabledOk returns a tuple with the DefaultRulesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRulesEnabled

`func (o *UpdateNetworkApplianceTrafficShapingRulesRequest) SetDefaultRulesEnabled(v bool)`

SetDefaultRulesEnabled sets DefaultRulesEnabled field to given value.

### HasDefaultRulesEnabled

`func (o *UpdateNetworkApplianceTrafficShapingRulesRequest) HasDefaultRulesEnabled() bool`

HasDefaultRulesEnabled returns a boolean if a field has been set.

### GetRules

`func (o *UpdateNetworkApplianceTrafficShapingRulesRequest) GetRules() []UpdateNetworkApplianceTrafficShapingRulesRequestRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *UpdateNetworkApplianceTrafficShapingRulesRequest) GetRulesOk() (*[]UpdateNetworkApplianceTrafficShapingRulesRequestRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *UpdateNetworkApplianceTrafficShapingRulesRequest) SetRules(v []UpdateNetworkApplianceTrafficShapingRulesRequestRulesInner)`

SetRules sets Rules field to given value.

### HasRules

`func (o *UpdateNetworkApplianceTrafficShapingRulesRequest) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


