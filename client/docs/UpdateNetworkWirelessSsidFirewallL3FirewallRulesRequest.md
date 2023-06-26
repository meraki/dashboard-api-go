# UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner**](UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner.md) | An ordered array of the firewall rules for this SSID (not including the local LAN access rule or the default rule) | [optional] 
**AllowLanAccess** | Pointer to **bool** | Allow wireless client access to local LAN (boolean value - true allows access and false denies access) (optional) | [optional] 

## Methods

### NewUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest

`func NewUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest() *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest`

NewUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest instantiates a new UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestWithDefaults

`func NewUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestWithDefaults() *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest`

NewUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestWithDefaults instantiates a new UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest) GetRules() []UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest) GetRulesOk() (*[]UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest) SetRules(v []UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner)`

SetRules sets Rules field to given value.

### HasRules

`func (o *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetAllowLanAccess

`func (o *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest) GetAllowLanAccess() bool`

GetAllowLanAccess returns the AllowLanAccess field if non-nil, zero value otherwise.

### GetAllowLanAccessOk

`func (o *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest) GetAllowLanAccessOk() (*bool, bool)`

GetAllowLanAccessOk returns a tuple with the AllowLanAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowLanAccess

`func (o *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest) SetAllowLanAccess(v bool)`

SetAllowLanAccess sets AllowLanAccess field to given value.

### HasAllowLanAccess

`func (o *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest) HasAllowLanAccess() bool`

HasAllowLanAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


