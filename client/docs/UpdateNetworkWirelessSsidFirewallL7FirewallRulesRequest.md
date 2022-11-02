# UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner**](UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner.md) | An array of L7 firewall rules for this SSID. Rules will get applied in the same order user has specified in request. Empty array will clear the L7 firewall rule configuration. | [optional] 

## Methods

### NewUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest

`func NewUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest() *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest`

NewUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest instantiates a new UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestWithDefaults

`func NewUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestWithDefaults() *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest`

NewUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestWithDefaults instantiates a new UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest) GetRules() []UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest) GetRulesOk() (*[]UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest) SetRules(v []UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner)`

SetRules sets Rules field to given value.

### HasRules

`func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


