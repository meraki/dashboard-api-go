# NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Settings** | Pointer to **string** | How firewall and traffic shaping rules are enforced. Can be &#39;network default&#39;, &#39;ignore&#39; or &#39;custom&#39;. | [optional] 
**TrafficShapingRules** | Pointer to [**[]NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules**](NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules.md) |     An array of traffic shaping rules. Rules are applied in the order that     they are specified in. An empty list (or null) means no rules. Note that     you are allowed a maximum of 8 rules.  | [optional] 
**L3FirewallRules** | Pointer to [**[]NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules**](NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules.md) | An ordered array of the L3 firewall rules | [optional] 
**L7FirewallRules** | Pointer to [**[]NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules**](NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules.md) | An ordered array of L7 firewall rules | [optional] 

## Methods

### NewNetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping

`func NewNetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping() *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping`

NewNetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping instantiates a new NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingWithDefaults

`func NewNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingWithDefaults() *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping`

NewNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingWithDefaults instantiates a new NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettings

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping) GetSettings() string`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping) GetSettingsOk() (*string, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping) SetSettings(v string)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetTrafficShapingRules

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping) GetTrafficShapingRules() []NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules`

GetTrafficShapingRules returns the TrafficShapingRules field if non-nil, zero value otherwise.

### GetTrafficShapingRulesOk

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping) GetTrafficShapingRulesOk() (*[]NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules, bool)`

GetTrafficShapingRulesOk returns a tuple with the TrafficShapingRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficShapingRules

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping) SetTrafficShapingRules(v []NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules)`

SetTrafficShapingRules sets TrafficShapingRules field to given value.

### HasTrafficShapingRules

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping) HasTrafficShapingRules() bool`

HasTrafficShapingRules returns a boolean if a field has been set.

### GetL3FirewallRules

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping) GetL3FirewallRules() []NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules`

GetL3FirewallRules returns the L3FirewallRules field if non-nil, zero value otherwise.

### GetL3FirewallRulesOk

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping) GetL3FirewallRulesOk() (*[]NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules, bool)`

GetL3FirewallRulesOk returns a tuple with the L3FirewallRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL3FirewallRules

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping) SetL3FirewallRules(v []NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules)`

SetL3FirewallRules sets L3FirewallRules field to given value.

### HasL3FirewallRules

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping) HasL3FirewallRules() bool`

HasL3FirewallRules returns a boolean if a field has been set.

### GetL7FirewallRules

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping) GetL7FirewallRules() []NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules`

GetL7FirewallRules returns the L7FirewallRules field if non-nil, zero value otherwise.

### GetL7FirewallRulesOk

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping) GetL7FirewallRulesOk() (*[]NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules, bool)`

GetL7FirewallRulesOk returns a tuple with the L7FirewallRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL7FirewallRules

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping) SetL7FirewallRules(v []NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules)`

SetL7FirewallRules sets L7FirewallRules field to given value.

### HasL7FirewallRules

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping) HasL7FirewallRules() bool`

HasL7FirewallRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


