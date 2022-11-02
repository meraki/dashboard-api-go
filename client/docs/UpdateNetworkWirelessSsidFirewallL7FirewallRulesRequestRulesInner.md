# UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | Pointer to **string** | &#39;Deny&#39; traffic specified by this rule | [optional] 
**Type** | Pointer to **string** | Type of the L7 firewall rule. One of: &#39;application&#39;, &#39;applicationCategory&#39;, &#39;host&#39;, &#39;port&#39;, &#39;ipRange&#39; | [optional] 
**Value** | Pointer to **string** | The value of what needs to get blocked. Format of the value varies depending on type of the firewall rule selected. | [optional] 

## Methods

### NewUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner

`func NewUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner() *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner`

NewUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner instantiates a new UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInnerWithDefaults

`func NewUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInnerWithDefaults() *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner`

NewUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInnerWithDefaults instantiates a new UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetType

`func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


