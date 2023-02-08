# CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Settings** | Pointer to **string** | How firewall and traffic shaping rules are enforced. Can be &#39;network default&#39;, &#39;ignore&#39; or &#39;custom&#39;. | [optional] 
**TrafficShapingRules** | Pointer to [**[]CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingTrafficShapingRulesInner**](CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingTrafficShapingRulesInner.md) |     An array of traffic shaping rules. Rules are applied in the order that     they are specified in. An empty list (or null) means no rules. Note that     you are allowed a maximum of 8 rules.  | [optional] 
**L3FirewallRules** | Pointer to [**[]CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner**](CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner.md) | An ordered array of the L3 firewall rules | [optional] 
**L7FirewallRules** | Pointer to [**[]CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner**](CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner.md) | An ordered array of L7 firewall rules | [optional] 

## Methods

### NewCreateNetworkGroupPolicyRequestFirewallAndTrafficShaping

`func NewCreateNetworkGroupPolicyRequestFirewallAndTrafficShaping() *CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping`

NewCreateNetworkGroupPolicyRequestFirewallAndTrafficShaping instantiates a new CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingWithDefaults

`func NewCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingWithDefaults() *CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping`

NewCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingWithDefaults instantiates a new CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettings

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) GetSettings() string`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) GetSettingsOk() (*string, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) SetSettings(v string)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetTrafficShapingRules

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) GetTrafficShapingRules() []CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingTrafficShapingRulesInner`

GetTrafficShapingRules returns the TrafficShapingRules field if non-nil, zero value otherwise.

### GetTrafficShapingRulesOk

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) GetTrafficShapingRulesOk() (*[]CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingTrafficShapingRulesInner, bool)`

GetTrafficShapingRulesOk returns a tuple with the TrafficShapingRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficShapingRules

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) SetTrafficShapingRules(v []CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingTrafficShapingRulesInner)`

SetTrafficShapingRules sets TrafficShapingRules field to given value.

### HasTrafficShapingRules

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) HasTrafficShapingRules() bool`

HasTrafficShapingRules returns a boolean if a field has been set.

### GetL3FirewallRules

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) GetL3FirewallRules() []CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner`

GetL3FirewallRules returns the L3FirewallRules field if non-nil, zero value otherwise.

### GetL3FirewallRulesOk

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) GetL3FirewallRulesOk() (*[]CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner, bool)`

GetL3FirewallRulesOk returns a tuple with the L3FirewallRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL3FirewallRules

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) SetL3FirewallRules(v []CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner)`

SetL3FirewallRules sets L3FirewallRules field to given value.

### HasL3FirewallRules

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) HasL3FirewallRules() bool`

HasL3FirewallRules returns a boolean if a field has been set.

### GetL7FirewallRules

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) GetL7FirewallRules() []CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner`

GetL7FirewallRules returns the L7FirewallRules field if non-nil, zero value otherwise.

### GetL7FirewallRulesOk

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) GetL7FirewallRulesOk() (*[]CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner, bool)`

GetL7FirewallRulesOk returns a tuple with the L7FirewallRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL7FirewallRules

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) SetL7FirewallRules(v []CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner)`

SetL7FirewallRules sets L7FirewallRules field to given value.

### HasL7FirewallRules

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) HasL7FirewallRules() bool`

HasL7FirewallRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


