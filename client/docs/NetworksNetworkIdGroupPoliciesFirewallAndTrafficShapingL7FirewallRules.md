# NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | Pointer to **string** | The policy applied to matching traffic. Must be &#39;deny&#39;. | [optional] 
**Type** | Pointer to **string** | Type of the L7 Rule. Must be &#39;application&#39;, &#39;applicationCategory&#39;, &#39;host&#39;, &#39;port&#39; or &#39;ipRange&#39; | [optional] 
**Value** | Pointer to **string** | The &#39;value&#39; of what you want to block. If &#39;type&#39; is &#39;host&#39;, &#39;port&#39; or &#39;ipRange&#39;, &#39;value&#39; must be a string matching either a hostname (e.g. somewhere.com), a port (e.g. 8080), or an IP range (e.g. 192.1.0.0/16). If &#39;type&#39; is &#39;application&#39; or &#39;applicationCategory&#39;, then &#39;value&#39; must be an object with an ID for the application. | [optional] 

## Methods

### NewNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules

`func NewNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules() *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules`

NewNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules instantiates a new NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRulesWithDefaults

`func NewNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRulesWithDefaults() *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules`

NewNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRulesWithDefaults instantiates a new NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetType

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL7FirewallRules) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


