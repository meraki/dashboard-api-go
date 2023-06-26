# NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | Pointer to **int32** | Operation code of the packet. | [optional] 
**Htype** | Pointer to **int32** | Hardware type code of the packet. | [optional] 
**Hlen** | Pointer to **int32** | Hardware length of the packet. | [optional] 
**Hops** | Pointer to **int32** | Number of hops the packet took. | [optional] 
**Xid** | Pointer to **string** | Transaction id of the packet. | [optional] 
**Secs** | Pointer to **int32** | Number of seconds since receiving the packet. | [optional] 
**Flags** | Pointer to **string** | Packet flags. | [optional] 
**Ciaddr** | Pointer to **string** | Client IP address of the packet. | [optional] 
**Yiaddr** | Pointer to **string** | Assigned IP address of the packet. | [optional] 
**Siaddr** | Pointer to **string** | Server IP address of the packet. | [optional] 
**Giaddr** | Pointer to **string** | Gateway IP address of the packet. | [optional] 
**Chaddr** | Pointer to **string** | Client hardware address of the packet. | [optional] 
**Sname** | Pointer to **string** | Server identifier address of the packet. | [optional] 
**MagicCookie** | Pointer to **string** | Magic cookie of the packet. | [optional] 
**Options** | Pointer to [**[]NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions**](NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions.md) | Additional DHCP options of the packet. | [optional] 

## Methods

### NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields

`func NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields() *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields`

NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields instantiates a new NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsWithDefaults

`func NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsWithDefaults() *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields`

NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsWithDefaults instantiates a new NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetOp() int32`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetOpOk() (*int32, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) SetOp(v int32)`

SetOp sets Op field to given value.

### HasOp

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetHtype

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetHtype() int32`

GetHtype returns the Htype field if non-nil, zero value otherwise.

### GetHtypeOk

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetHtypeOk() (*int32, bool)`

GetHtypeOk returns a tuple with the Htype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtype

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) SetHtype(v int32)`

SetHtype sets Htype field to given value.

### HasHtype

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) HasHtype() bool`

HasHtype returns a boolean if a field has been set.

### GetHlen

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetHlen() int32`

GetHlen returns the Hlen field if non-nil, zero value otherwise.

### GetHlenOk

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetHlenOk() (*int32, bool)`

GetHlenOk returns a tuple with the Hlen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHlen

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) SetHlen(v int32)`

SetHlen sets Hlen field to given value.

### HasHlen

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) HasHlen() bool`

HasHlen returns a boolean if a field has been set.

### GetHops

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetHops() int32`

GetHops returns the Hops field if non-nil, zero value otherwise.

### GetHopsOk

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetHopsOk() (*int32, bool)`

GetHopsOk returns a tuple with the Hops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHops

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) SetHops(v int32)`

SetHops sets Hops field to given value.

### HasHops

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) HasHops() bool`

HasHops returns a boolean if a field has been set.

### GetXid

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetXid() string`

GetXid returns the Xid field if non-nil, zero value otherwise.

### GetXidOk

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetXidOk() (*string, bool)`

GetXidOk returns a tuple with the Xid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXid

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) SetXid(v string)`

SetXid sets Xid field to given value.

### HasXid

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) HasXid() bool`

HasXid returns a boolean if a field has been set.

### GetSecs

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetSecs() int32`

GetSecs returns the Secs field if non-nil, zero value otherwise.

### GetSecsOk

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetSecsOk() (*int32, bool)`

GetSecsOk returns a tuple with the Secs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecs

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) SetSecs(v int32)`

SetSecs sets Secs field to given value.

### HasSecs

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) HasSecs() bool`

HasSecs returns a boolean if a field has been set.

### GetFlags

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetFlags() string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetFlagsOk() (*string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) SetFlags(v string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetCiaddr

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetCiaddr() string`

GetCiaddr returns the Ciaddr field if non-nil, zero value otherwise.

### GetCiaddrOk

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetCiaddrOk() (*string, bool)`

GetCiaddrOk returns a tuple with the Ciaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiaddr

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) SetCiaddr(v string)`

SetCiaddr sets Ciaddr field to given value.

### HasCiaddr

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) HasCiaddr() bool`

HasCiaddr returns a boolean if a field has been set.

### GetYiaddr

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetYiaddr() string`

GetYiaddr returns the Yiaddr field if non-nil, zero value otherwise.

### GetYiaddrOk

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetYiaddrOk() (*string, bool)`

GetYiaddrOk returns a tuple with the Yiaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYiaddr

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) SetYiaddr(v string)`

SetYiaddr sets Yiaddr field to given value.

### HasYiaddr

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) HasYiaddr() bool`

HasYiaddr returns a boolean if a field has been set.

### GetSiaddr

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetSiaddr() string`

GetSiaddr returns the Siaddr field if non-nil, zero value otherwise.

### GetSiaddrOk

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetSiaddrOk() (*string, bool)`

GetSiaddrOk returns a tuple with the Siaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiaddr

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) SetSiaddr(v string)`

SetSiaddr sets Siaddr field to given value.

### HasSiaddr

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) HasSiaddr() bool`

HasSiaddr returns a boolean if a field has been set.

### GetGiaddr

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetGiaddr() string`

GetGiaddr returns the Giaddr field if non-nil, zero value otherwise.

### GetGiaddrOk

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetGiaddrOk() (*string, bool)`

GetGiaddrOk returns a tuple with the Giaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiaddr

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) SetGiaddr(v string)`

SetGiaddr sets Giaddr field to given value.

### HasGiaddr

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) HasGiaddr() bool`

HasGiaddr returns a boolean if a field has been set.

### GetChaddr

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetChaddr() string`

GetChaddr returns the Chaddr field if non-nil, zero value otherwise.

### GetChaddrOk

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetChaddrOk() (*string, bool)`

GetChaddrOk returns a tuple with the Chaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChaddr

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) SetChaddr(v string)`

SetChaddr sets Chaddr field to given value.

### HasChaddr

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) HasChaddr() bool`

HasChaddr returns a boolean if a field has been set.

### GetSname

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetSname() string`

GetSname returns the Sname field if non-nil, zero value otherwise.

### GetSnameOk

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetSnameOk() (*string, bool)`

GetSnameOk returns a tuple with the Sname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSname

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) SetSname(v string)`

SetSname sets Sname field to given value.

### HasSname

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) HasSname() bool`

HasSname returns a boolean if a field has been set.

### GetMagicCookie

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetMagicCookie() string`

GetMagicCookie returns the MagicCookie field if non-nil, zero value otherwise.

### GetMagicCookieOk

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetMagicCookieOk() (*string, bool)`

GetMagicCookieOk returns a tuple with the MagicCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMagicCookie

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) SetMagicCookie(v string)`

SetMagicCookie sets MagicCookie field to given value.

### HasMagicCookie

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) HasMagicCookie() bool`

HasMagicCookie returns a boolean if a field has been set.

### GetOptions

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetOptions() []NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) GetOptionsOk() (*[]NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) SetOptions(v []NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


