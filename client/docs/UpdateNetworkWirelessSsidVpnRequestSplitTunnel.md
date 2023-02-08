# UpdateNetworkWirelessSsidVpnRequestSplitTunnel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | If true, VPN split tunnel is enabled. | [optional] 
**Rules** | Pointer to [**[]UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner**](UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner.md) | List of VPN split tunnel rules. | [optional] 

## Methods

### NewUpdateNetworkWirelessSsidVpnRequestSplitTunnel

`func NewUpdateNetworkWirelessSsidVpnRequestSplitTunnel() *UpdateNetworkWirelessSsidVpnRequestSplitTunnel`

NewUpdateNetworkWirelessSsidVpnRequestSplitTunnel instantiates a new UpdateNetworkWirelessSsidVpnRequestSplitTunnel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkWirelessSsidVpnRequestSplitTunnelWithDefaults

`func NewUpdateNetworkWirelessSsidVpnRequestSplitTunnelWithDefaults() *UpdateNetworkWirelessSsidVpnRequestSplitTunnel`

NewUpdateNetworkWirelessSsidVpnRequestSplitTunnelWithDefaults instantiates a new UpdateNetworkWirelessSsidVpnRequestSplitTunnel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnel) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRules

`func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnel) GetRules() []UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnel) GetRulesOk() (*[]UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnel) SetRules(v []UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner)`

SetRules sets Rules field to given value.

### HasRules

`func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnel) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


