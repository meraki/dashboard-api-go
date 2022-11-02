# CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Description of the rule (optional) | [optional] 
**Policy** | **string** | &#39;allow&#39; or &#39;deny&#39; traffic specified by this rule | 
**Protocol** | **string** | The type of protocol (must be &#39;tcp&#39;, &#39;udp&#39;, &#39;icmp&#39;, &#39;icmp6&#39; or &#39;any&#39;) | 
**DestPort** | Pointer to **string** | Destination port (integer in the range 1-65535), a port range (e.g. 8080-9090), or &#39;any&#39; | [optional] 
**DestCidr** | **string** | Destination IP address (in IP or CIDR notation), a fully-qualified domain name (FQDN, if your network supports it) or &#39;any&#39;. | 

## Methods

### NewCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner

`func NewCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner(policy string, protocol string, destCidr string, ) *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner`

NewCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner instantiates a new CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInnerWithDefaults

`func NewCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInnerWithDefaults() *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner`

NewCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInnerWithDefaults instantiates a new CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetPolicy

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) SetPolicy(v string)`

SetPolicy sets Policy field to given value.


### GetProtocol

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetDestPort

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) GetDestPort() string`

GetDestPort returns the DestPort field if non-nil, zero value otherwise.

### GetDestPortOk

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) GetDestPortOk() (*string, bool)`

GetDestPortOk returns a tuple with the DestPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestPort

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) SetDestPort(v string)`

SetDestPort sets DestPort field to given value.

### HasDestPort

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) HasDestPort() bool`

HasDestPort returns a boolean if a field has been set.

### GetDestCidr

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) GetDestCidr() string`

GetDestCidr returns the DestCidr field if non-nil, zero value otherwise.

### GetDestCidrOk

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) GetDestCidrOk() (*string, bool)`

GetDestCidrOk returns a tuple with the DestCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestCidr

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) SetDestCidr(v string)`

SetDestCidr sets DestCidr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


