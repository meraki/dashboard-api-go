# GetNetworkSwitchDhcpV4ServersSeen200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | Mac address of the server. | [optional] 
**Vlan** | Pointer to **int32** | Vlan id of the server. | [optional] 
**ClientId** | Pointer to **string** | Client id of the server if available. | [optional] 
**IsAllowed** | Pointer to **bool** | Whether the server is allowed or blocked. Always true for configured servers. | [optional] 
**LastSeenAt** | Pointer to **time.Time** | Last time the server was seen. | [optional] 
**SeenBy** | Pointer to [**[]GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerSeenByInner**](GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerSeenByInner.md) | Devices that saw the server. | [optional] 
**Type** | Pointer to **string** | server type. Can be a &#39;device&#39;, &#39;stack&#39;, or &#39;discovered&#39; (i.e client). | [optional] 
**Device** | Pointer to [**GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice**](GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice.md) |  | [optional] 
**Ipv4** | Pointer to [**GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4**](GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4.md) |  | [optional] 
**IsConfigured** | Pointer to **bool** | Whether the server is configured. | [optional] 
**LastAck** | Pointer to [**GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck**](GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck.md) |  | [optional] 
**LastPacket** | Pointer to [**GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket**](GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket.md) |  | [optional] 

## Methods

### NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInner

`func NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInner() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner`

NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInner instantiates a new GetNetworkSwitchDhcpV4ServersSeen200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerWithDefaults

`func NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerWithDefaults() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner`

NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerWithDefaults instantiates a new GetNetworkSwitchDhcpV4ServersSeen200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetVlan

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetClientId

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetIsAllowed

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetIsAllowed() bool`

GetIsAllowed returns the IsAllowed field if non-nil, zero value otherwise.

### GetIsAllowedOk

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetIsAllowedOk() (*bool, bool)`

GetIsAllowedOk returns a tuple with the IsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllowed

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) SetIsAllowed(v bool)`

SetIsAllowed sets IsAllowed field to given value.

### HasIsAllowed

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) HasIsAllowed() bool`

HasIsAllowed returns a boolean if a field has been set.

### GetLastSeenAt

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetLastSeenAt() time.Time`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetLastSeenAtOk() (*time.Time, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) SetLastSeenAt(v time.Time)`

SetLastSeenAt sets LastSeenAt field to given value.

### HasLastSeenAt

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) HasLastSeenAt() bool`

HasLastSeenAt returns a boolean if a field has been set.

### GetSeenBy

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetSeenBy() []GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerSeenByInner`

GetSeenBy returns the SeenBy field if non-nil, zero value otherwise.

### GetSeenByOk

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetSeenByOk() (*[]GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerSeenByInner, bool)`

GetSeenByOk returns a tuple with the SeenBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeenBy

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) SetSeenBy(v []GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerSeenByInner)`

SetSeenBy sets SeenBy field to given value.

### HasSeenBy

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) HasSeenBy() bool`

HasSeenBy returns a boolean if a field has been set.

### GetType

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDevice

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetDevice() GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetDeviceOk() (*GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) SetDevice(v GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetIpv4

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetIpv4() GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetIpv4Ok() (*GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) SetIpv4(v GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIsConfigured

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetIsConfigured() bool`

GetIsConfigured returns the IsConfigured field if non-nil, zero value otherwise.

### GetIsConfiguredOk

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetIsConfiguredOk() (*bool, bool)`

GetIsConfiguredOk returns a tuple with the IsConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConfigured

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) SetIsConfigured(v bool)`

SetIsConfigured sets IsConfigured field to given value.

### HasIsConfigured

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) HasIsConfigured() bool`

HasIsConfigured returns a boolean if a field has been set.

### GetLastAck

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetLastAck() GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck`

GetLastAck returns the LastAck field if non-nil, zero value otherwise.

### GetLastAckOk

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetLastAckOk() (*GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck, bool)`

GetLastAckOk returns a tuple with the LastAck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAck

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) SetLastAck(v GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck)`

SetLastAck sets LastAck field to given value.

### HasLastAck

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) HasLastAck() bool`

HasLastAck returns a boolean if a field has been set.

### GetLastPacket

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetLastPacket() GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket`

GetLastPacket returns the LastPacket field if non-nil, zero value otherwise.

### GetLastPacketOk

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetLastPacketOk() (*GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket, bool)`

GetLastPacketOk returns a tuple with the LastPacket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPacket

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) SetLastPacket(v GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket)`

SetLastPacket sets LastPacket field to given value.

### HasLastPacket

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) HasLastPacket() bool`

HasLastPacket returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


