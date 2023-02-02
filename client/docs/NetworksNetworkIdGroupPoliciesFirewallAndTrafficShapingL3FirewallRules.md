# NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Description of the rule (optional) | [optional] 
**Policy** | **string** | &#39;allow&#39; or &#39;deny&#39; traffic specified by this rule | 
**Protocol** | **string** | The type of protocol (must be &#39;tcp&#39;, &#39;udp&#39;, &#39;icmp&#39;, &#39;icmp6&#39; or &#39;any&#39;) | 
**DestPort** | Pointer to **string** | Destination port (integer in the range 1-65535), a port range (e.g. 8080-9090), or &#39;any&#39; | [optional] 
**DestCidr** | **string** | Destination IP address (in IP or CIDR notation), a fully-qualified domain name (FQDN, if your network supports it) or &#39;any&#39;. | 

## Methods

### NewNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules

`func NewNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules(policy string, protocol string, destCidr string, ) *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules`

NewNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules instantiates a new NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRulesWithDefaults

`func NewNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRulesWithDefaults() *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules`

NewNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRulesWithDefaults instantiates a new NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetPolicy

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) SetPolicy(v string)`

SetPolicy sets Policy field to given value.


### GetProtocol

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetDestPort

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) GetDestPort() string`

GetDestPort returns the DestPort field if non-nil, zero value otherwise.

### GetDestPortOk

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) GetDestPortOk() (*string, bool)`

GetDestPortOk returns a tuple with the DestPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestPort

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) SetDestPort(v string)`

SetDestPort sets DestPort field to given value.

### HasDestPort

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) HasDestPort() bool`

HasDestPort returns a boolean if a field has been set.

### GetDestCidr

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) GetDestCidr() string`

GetDestCidr returns the DestCidr field if non-nil, zero value otherwise.

### GetDestCidrOk

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) GetDestCidrOk() (*string, bool)`

GetDestCidrOk returns a tuple with the DestCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestCidr

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingL3FirewallRules) SetDestCidr(v string)`

SetDestCidr sets DestCidr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


