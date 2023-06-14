# CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | Pointer to **string** | The policy applied to matching traffic. Must be &#39;deny&#39;. | [optional] 
**Type** | Pointer to **string** | Type of the L7 Rule. Must be &#39;application&#39;, &#39;applicationCategory&#39;, &#39;host&#39;, &#39;port&#39; or &#39;ipRange&#39; | [optional] 
**Value** | Pointer to **string** | The &#39;value&#39; of what you want to block. If &#39;type&#39; is &#39;host&#39;, &#39;port&#39; or &#39;ipRange&#39;, &#39;value&#39; must be a string matching either a hostname (e.g. somewhere.com), a port (e.g. 8080), or an IP range (e.g. 192.1.0.0/16). If &#39;type&#39; is &#39;application&#39; or &#39;applicationCategory&#39;, then &#39;value&#39; must be an object with an ID for the application. | [optional] 

## Methods

### NewCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner

`func NewCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner() *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner`

NewCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner instantiates a new CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInnerWithDefaults

`func NewCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInnerWithDefaults() *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner`

NewCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInnerWithDefaults instantiates a new CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetType

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


