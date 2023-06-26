# NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A description of the rule | [optional] 
**Protocol** | Pointer to **string** | &#39;tcp&#39; or &#39;udp&#39; | [optional] 
**PublicPort** | Pointer to **string** | Destination port of the traffic that is arriving on the WAN | [optional] 
**LocalIp** | Pointer to **string** | Local IP address to which traffic will be forwarded | [optional] 
**LocalPort** | Pointer to **string** | Destination port of the forwarded traffic that will be sent from the MX to the specified host on the LAN. If you simply wish to forward the traffic without translating the port, this should be the same as the Public port | [optional] 
**AllowedIps** | Pointer to **[]string** | Remote IP addresses or ranges that are permitted to access the internal resource via this port forwarding rule, or &#39;any&#39; | [optional] 

## Methods

### NewNetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules

`func NewNetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules() *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules`

NewNetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules instantiates a new NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRulesWithDefaults

`func NewNetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRulesWithDefaults() *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules`

NewNetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRulesWithDefaults instantiates a new NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProtocol

`func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetPublicPort

`func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) GetPublicPort() string`

GetPublicPort returns the PublicPort field if non-nil, zero value otherwise.

### GetPublicPortOk

`func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) GetPublicPortOk() (*string, bool)`

GetPublicPortOk returns a tuple with the PublicPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicPort

`func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) SetPublicPort(v string)`

SetPublicPort sets PublicPort field to given value.

### HasPublicPort

`func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) HasPublicPort() bool`

HasPublicPort returns a boolean if a field has been set.

### GetLocalIp

`func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) GetLocalIp() string`

GetLocalIp returns the LocalIp field if non-nil, zero value otherwise.

### GetLocalIpOk

`func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) GetLocalIpOk() (*string, bool)`

GetLocalIpOk returns a tuple with the LocalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIp

`func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) SetLocalIp(v string)`

SetLocalIp sets LocalIp field to given value.

### HasLocalIp

`func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) HasLocalIp() bool`

HasLocalIp returns a boolean if a field has been set.

### GetLocalPort

`func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) GetLocalPort() string`

GetLocalPort returns the LocalPort field if non-nil, zero value otherwise.

### GetLocalPortOk

`func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) GetLocalPortOk() (*string, bool)`

GetLocalPortOk returns a tuple with the LocalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPort

`func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) SetLocalPort(v string)`

SetLocalPort sets LocalPort field to given value.

### HasLocalPort

`func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) HasLocalPort() bool`

HasLocalPort returns a boolean if a field has been set.

### GetAllowedIps

`func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) GetAllowedIps() []string`

GetAllowedIps returns the AllowedIps field if non-nil, zero value otherwise.

### GetAllowedIpsOk

`func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) GetAllowedIpsOk() (*[]string, bool)`

GetAllowedIpsOk returns a tuple with the AllowedIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedIps

`func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) SetAllowedIps(v []string)`

SetAllowedIps sets AllowedIps field to given value.

### HasAllowedIps

`func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules) HasAllowedIps() bool`

HasAllowedIps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


