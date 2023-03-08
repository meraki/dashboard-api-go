# NetworksNetworkIdApplianceFirewallL7FirewallRulesRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | Pointer to **string** | &#39;Deny&#39; traffic specified by this rule | [optional] 
**Type** | Pointer to **string** | Type of the L7 rule. One of: &#39;application&#39;, &#39;applicationCategory&#39;, &#39;host&#39;, &#39;port&#39;, &#39;ipRange&#39; | [optional] 
**Value** | Pointer to **string** | The &#39;value&#39; of what you want to block. Format of &#39;value&#39; varies depending on type of the rule. The application categories and application ids can be retrieved from the the &#39;MX L7 application categories&#39; endpoint. The countries follow the two-letter ISO 3166-1 alpha-2 format. | [optional] 

## Methods

### NewNetworksNetworkIdApplianceFirewallL7FirewallRulesRules

`func NewNetworksNetworkIdApplianceFirewallL7FirewallRulesRules() *NetworksNetworkIdApplianceFirewallL7FirewallRulesRules`

NewNetworksNetworkIdApplianceFirewallL7FirewallRulesRules instantiates a new NetworksNetworkIdApplianceFirewallL7FirewallRulesRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdApplianceFirewallL7FirewallRulesRulesWithDefaults

`func NewNetworksNetworkIdApplianceFirewallL7FirewallRulesRulesWithDefaults() *NetworksNetworkIdApplianceFirewallL7FirewallRulesRules`

NewNetworksNetworkIdApplianceFirewallL7FirewallRulesRulesWithDefaults instantiates a new NetworksNetworkIdApplianceFirewallL7FirewallRulesRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *NetworksNetworkIdApplianceFirewallL7FirewallRulesRules) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *NetworksNetworkIdApplianceFirewallL7FirewallRulesRules) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *NetworksNetworkIdApplianceFirewallL7FirewallRulesRules) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *NetworksNetworkIdApplianceFirewallL7FirewallRulesRules) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetType

`func (o *NetworksNetworkIdApplianceFirewallL7FirewallRulesRules) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworksNetworkIdApplianceFirewallL7FirewallRulesRules) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworksNetworkIdApplianceFirewallL7FirewallRulesRules) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NetworksNetworkIdApplianceFirewallL7FirewallRulesRules) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *NetworksNetworkIdApplianceFirewallL7FirewallRulesRules) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NetworksNetworkIdApplianceFirewallL7FirewallRulesRules) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NetworksNetworkIdApplianceFirewallL7FirewallRulesRules) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *NetworksNetworkIdApplianceFirewallL7FirewallRulesRules) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


