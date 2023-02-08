# UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A descriptive name for the rule | [optional] 
**LanIp** | **string** | The IP address of the server or device that hosts the internal resource that you wish to make available on the WAN | 
**Uplink** | Pointer to **string** | The physical WAN interface on which the traffic will arrive (&#39;internet1&#39; or, if available, &#39;internet2&#39; or &#39;both&#39;) | [optional] 
**PublicPort** | **string** | A port or port ranges that will be forwarded to the host on the LAN | 
**LocalPort** | **string** | A port or port ranges that will receive the forwarded traffic from the WAN | 
**AllowedIps** | **[]string** | An array of ranges of WAN IP addresses that are allowed to make inbound connections on the specified ports or port ranges (or any) | 
**Protocol** | **string** | TCP or UDP | 

## Methods

### NewUpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner

`func NewUpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner(lanIp string, publicPort string, localPort string, allowedIps []string, protocol string, ) *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner`

NewUpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner instantiates a new UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInnerWithDefaults

`func NewUpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInnerWithDefaults() *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner`

NewUpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInnerWithDefaults instantiates a new UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLanIp

`func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) GetLanIp() string`

GetLanIp returns the LanIp field if non-nil, zero value otherwise.

### GetLanIpOk

`func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) GetLanIpOk() (*string, bool)`

GetLanIpOk returns a tuple with the LanIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanIp

`func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) SetLanIp(v string)`

SetLanIp sets LanIp field to given value.


### GetUplink

`func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) GetUplink() string`

GetUplink returns the Uplink field if non-nil, zero value otherwise.

### GetUplinkOk

`func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) GetUplinkOk() (*string, bool)`

GetUplinkOk returns a tuple with the Uplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplink

`func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) SetUplink(v string)`

SetUplink sets Uplink field to given value.

### HasUplink

`func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) HasUplink() bool`

HasUplink returns a boolean if a field has been set.

### GetPublicPort

`func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) GetPublicPort() string`

GetPublicPort returns the PublicPort field if non-nil, zero value otherwise.

### GetPublicPortOk

`func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) GetPublicPortOk() (*string, bool)`

GetPublicPortOk returns a tuple with the PublicPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicPort

`func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) SetPublicPort(v string)`

SetPublicPort sets PublicPort field to given value.


### GetLocalPort

`func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) GetLocalPort() string`

GetLocalPort returns the LocalPort field if non-nil, zero value otherwise.

### GetLocalPortOk

`func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) GetLocalPortOk() (*string, bool)`

GetLocalPortOk returns a tuple with the LocalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPort

`func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) SetLocalPort(v string)`

SetLocalPort sets LocalPort field to given value.


### GetAllowedIps

`func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) GetAllowedIps() []string`

GetAllowedIps returns the AllowedIps field if non-nil, zero value otherwise.

### GetAllowedIpsOk

`func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) GetAllowedIpsOk() (*[]string, bool)`

GetAllowedIpsOk returns a tuple with the AllowedIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedIps

`func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) SetAllowedIps(v []string)`

SetAllowedIps sets AllowedIps field to given value.


### GetProtocol

`func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


