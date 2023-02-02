# InlineResponse20060

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | Mac address of the server. | [optional] 
**Vlan** | Pointer to **int32** | Vlan id of the server. | [optional] 
**ClientId** | Pointer to **string** | Client id of the server if available. | [optional] 
**IsAllowed** | Pointer to **bool** | Whether the server is allowed or blocked. Always true for configured servers. | [optional] 
**LastSeenAt** | Pointer to **time.Time** | Last time the server was seen. | [optional] 
**SeenBy** | Pointer to [**[]NetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy**](NetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy.md) | Devices that saw the server. | [optional] 
**Type** | Pointer to **string** | server type. Can be a &#39;device&#39;, &#39;stack&#39;, or &#39;discovered&#39; (i.e client). | [optional] 
**Device** | Pointer to [**NetworksNetworkIdSwitchDhcpV4ServersSeenDevice**](NetworksNetworkIdSwitchDhcpV4ServersSeenDevice.md) |  | [optional] 
**Ipv4** | Pointer to [**NetworksNetworkIdSwitchDhcpV4ServersSeenIpv4**](NetworksNetworkIdSwitchDhcpV4ServersSeenIpv4.md) |  | [optional] 
**IsConfigured** | Pointer to **bool** | Whether the server is configured. | [optional] 
**LastAck** | Pointer to [**NetworksNetworkIdSwitchDhcpV4ServersSeenLastAck**](NetworksNetworkIdSwitchDhcpV4ServersSeenLastAck.md) |  | [optional] 
**LastPacket** | Pointer to [**NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket**](NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket.md) |  | [optional] 

## Methods

### NewInlineResponse20060

`func NewInlineResponse20060() *InlineResponse20060`

NewInlineResponse20060 instantiates a new InlineResponse20060 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20060WithDefaults

`func NewInlineResponse20060WithDefaults() *InlineResponse20060`

NewInlineResponse20060WithDefaults instantiates a new InlineResponse20060 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *InlineResponse20060) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse20060) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse20060) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse20060) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetVlan

`func (o *InlineResponse20060) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineResponse20060) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineResponse20060) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *InlineResponse20060) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetClientId

`func (o *InlineResponse20060) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *InlineResponse20060) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *InlineResponse20060) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *InlineResponse20060) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetIsAllowed

`func (o *InlineResponse20060) GetIsAllowed() bool`

GetIsAllowed returns the IsAllowed field if non-nil, zero value otherwise.

### GetIsAllowedOk

`func (o *InlineResponse20060) GetIsAllowedOk() (*bool, bool)`

GetIsAllowedOk returns a tuple with the IsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllowed

`func (o *InlineResponse20060) SetIsAllowed(v bool)`

SetIsAllowed sets IsAllowed field to given value.

### HasIsAllowed

`func (o *InlineResponse20060) HasIsAllowed() bool`

HasIsAllowed returns a boolean if a field has been set.

### GetLastSeenAt

`func (o *InlineResponse20060) GetLastSeenAt() time.Time`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *InlineResponse20060) GetLastSeenAtOk() (*time.Time, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *InlineResponse20060) SetLastSeenAt(v time.Time)`

SetLastSeenAt sets LastSeenAt field to given value.

### HasLastSeenAt

`func (o *InlineResponse20060) HasLastSeenAt() bool`

HasLastSeenAt returns a boolean if a field has been set.

### GetSeenBy

`func (o *InlineResponse20060) GetSeenBy() []NetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy`

GetSeenBy returns the SeenBy field if non-nil, zero value otherwise.

### GetSeenByOk

`func (o *InlineResponse20060) GetSeenByOk() (*[]NetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy, bool)`

GetSeenByOk returns a tuple with the SeenBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeenBy

`func (o *InlineResponse20060) SetSeenBy(v []NetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy)`

SetSeenBy sets SeenBy field to given value.

### HasSeenBy

`func (o *InlineResponse20060) HasSeenBy() bool`

HasSeenBy returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse20060) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse20060) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse20060) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse20060) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDevice

`func (o *InlineResponse20060) GetDevice() NetworksNetworkIdSwitchDhcpV4ServersSeenDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *InlineResponse20060) GetDeviceOk() (*NetworksNetworkIdSwitchDhcpV4ServersSeenDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *InlineResponse20060) SetDevice(v NetworksNetworkIdSwitchDhcpV4ServersSeenDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *InlineResponse20060) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetIpv4

`func (o *InlineResponse20060) GetIpv4() NetworksNetworkIdSwitchDhcpV4ServersSeenIpv4`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *InlineResponse20060) GetIpv4Ok() (*NetworksNetworkIdSwitchDhcpV4ServersSeenIpv4, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *InlineResponse20060) SetIpv4(v NetworksNetworkIdSwitchDhcpV4ServersSeenIpv4)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *InlineResponse20060) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIsConfigured

`func (o *InlineResponse20060) GetIsConfigured() bool`

GetIsConfigured returns the IsConfigured field if non-nil, zero value otherwise.

### GetIsConfiguredOk

`func (o *InlineResponse20060) GetIsConfiguredOk() (*bool, bool)`

GetIsConfiguredOk returns a tuple with the IsConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConfigured

`func (o *InlineResponse20060) SetIsConfigured(v bool)`

SetIsConfigured sets IsConfigured field to given value.

### HasIsConfigured

`func (o *InlineResponse20060) HasIsConfigured() bool`

HasIsConfigured returns a boolean if a field has been set.

### GetLastAck

`func (o *InlineResponse20060) GetLastAck() NetworksNetworkIdSwitchDhcpV4ServersSeenLastAck`

GetLastAck returns the LastAck field if non-nil, zero value otherwise.

### GetLastAckOk

`func (o *InlineResponse20060) GetLastAckOk() (*NetworksNetworkIdSwitchDhcpV4ServersSeenLastAck, bool)`

GetLastAckOk returns a tuple with the LastAck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAck

`func (o *InlineResponse20060) SetLastAck(v NetworksNetworkIdSwitchDhcpV4ServersSeenLastAck)`

SetLastAck sets LastAck field to given value.

### HasLastAck

`func (o *InlineResponse20060) HasLastAck() bool`

HasLastAck returns a boolean if a field has been set.

### GetLastPacket

`func (o *InlineResponse20060) GetLastPacket() NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket`

GetLastPacket returns the LastPacket field if non-nil, zero value otherwise.

### GetLastPacketOk

`func (o *InlineResponse20060) GetLastPacketOk() (*NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket, bool)`

GetLastPacketOk returns a tuple with the LastPacket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPacket

`func (o *InlineResponse20060) SetLastPacket(v NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket)`

SetLastPacket sets LastPacket field to given value.

### HasLastPacket

`func (o *InlineResponse20060) HasLastPacket() bool`

HasLastPacket returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


