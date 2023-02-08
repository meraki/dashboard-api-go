# GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields

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
**Options** | Pointer to [**[]GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner**](GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner.md) | Additional DHCP options of the packet. | [optional] 

## Methods

### NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields

`func NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields`

NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields instantiates a new GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsWithDefaults

`func NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsWithDefaults() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields`

NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsWithDefaults instantiates a new GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetOp() int32`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetOpOk() (*int32, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) SetOp(v int32)`

SetOp sets Op field to given value.

### HasOp

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetHtype

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetHtype() int32`

GetHtype returns the Htype field if non-nil, zero value otherwise.

### GetHtypeOk

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetHtypeOk() (*int32, bool)`

GetHtypeOk returns a tuple with the Htype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtype

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) SetHtype(v int32)`

SetHtype sets Htype field to given value.

### HasHtype

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) HasHtype() bool`

HasHtype returns a boolean if a field has been set.

### GetHlen

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetHlen() int32`

GetHlen returns the Hlen field if non-nil, zero value otherwise.

### GetHlenOk

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetHlenOk() (*int32, bool)`

GetHlenOk returns a tuple with the Hlen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHlen

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) SetHlen(v int32)`

SetHlen sets Hlen field to given value.

### HasHlen

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) HasHlen() bool`

HasHlen returns a boolean if a field has been set.

### GetHops

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetHops() int32`

GetHops returns the Hops field if non-nil, zero value otherwise.

### GetHopsOk

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetHopsOk() (*int32, bool)`

GetHopsOk returns a tuple with the Hops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHops

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) SetHops(v int32)`

SetHops sets Hops field to given value.

### HasHops

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) HasHops() bool`

HasHops returns a boolean if a field has been set.

### GetXid

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetXid() string`

GetXid returns the Xid field if non-nil, zero value otherwise.

### GetXidOk

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetXidOk() (*string, bool)`

GetXidOk returns a tuple with the Xid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXid

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) SetXid(v string)`

SetXid sets Xid field to given value.

### HasXid

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) HasXid() bool`

HasXid returns a boolean if a field has been set.

### GetSecs

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetSecs() int32`

GetSecs returns the Secs field if non-nil, zero value otherwise.

### GetSecsOk

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetSecsOk() (*int32, bool)`

GetSecsOk returns a tuple with the Secs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecs

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) SetSecs(v int32)`

SetSecs sets Secs field to given value.

### HasSecs

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) HasSecs() bool`

HasSecs returns a boolean if a field has been set.

### GetFlags

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetFlags() string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetFlagsOk() (*string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) SetFlags(v string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetCiaddr

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetCiaddr() string`

GetCiaddr returns the Ciaddr field if non-nil, zero value otherwise.

### GetCiaddrOk

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetCiaddrOk() (*string, bool)`

GetCiaddrOk returns a tuple with the Ciaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiaddr

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) SetCiaddr(v string)`

SetCiaddr sets Ciaddr field to given value.

### HasCiaddr

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) HasCiaddr() bool`

HasCiaddr returns a boolean if a field has been set.

### GetYiaddr

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetYiaddr() string`

GetYiaddr returns the Yiaddr field if non-nil, zero value otherwise.

### GetYiaddrOk

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetYiaddrOk() (*string, bool)`

GetYiaddrOk returns a tuple with the Yiaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYiaddr

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) SetYiaddr(v string)`

SetYiaddr sets Yiaddr field to given value.

### HasYiaddr

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) HasYiaddr() bool`

HasYiaddr returns a boolean if a field has been set.

### GetSiaddr

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetSiaddr() string`

GetSiaddr returns the Siaddr field if non-nil, zero value otherwise.

### GetSiaddrOk

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetSiaddrOk() (*string, bool)`

GetSiaddrOk returns a tuple with the Siaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiaddr

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) SetSiaddr(v string)`

SetSiaddr sets Siaddr field to given value.

### HasSiaddr

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) HasSiaddr() bool`

HasSiaddr returns a boolean if a field has been set.

### GetGiaddr

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetGiaddr() string`

GetGiaddr returns the Giaddr field if non-nil, zero value otherwise.

### GetGiaddrOk

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetGiaddrOk() (*string, bool)`

GetGiaddrOk returns a tuple with the Giaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiaddr

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) SetGiaddr(v string)`

SetGiaddr sets Giaddr field to given value.

### HasGiaddr

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) HasGiaddr() bool`

HasGiaddr returns a boolean if a field has been set.

### GetChaddr

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetChaddr() string`

GetChaddr returns the Chaddr field if non-nil, zero value otherwise.

### GetChaddrOk

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetChaddrOk() (*string, bool)`

GetChaddrOk returns a tuple with the Chaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChaddr

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) SetChaddr(v string)`

SetChaddr sets Chaddr field to given value.

### HasChaddr

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) HasChaddr() bool`

HasChaddr returns a boolean if a field has been set.

### GetSname

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetSname() string`

GetSname returns the Sname field if non-nil, zero value otherwise.

### GetSnameOk

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetSnameOk() (*string, bool)`

GetSnameOk returns a tuple with the Sname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSname

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) SetSname(v string)`

SetSname sets Sname field to given value.

### HasSname

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) HasSname() bool`

HasSname returns a boolean if a field has been set.

### GetMagicCookie

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetMagicCookie() string`

GetMagicCookie returns the MagicCookie field if non-nil, zero value otherwise.

### GetMagicCookieOk

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetMagicCookieOk() (*string, bool)`

GetMagicCookieOk returns a tuple with the MagicCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMagicCookie

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) SetMagicCookie(v string)`

SetMagicCookie sets MagicCookie field to given value.

### HasMagicCookie

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) HasMagicCookie() bool`

HasMagicCookie returns a boolean if a field has been set.

### GetOptions

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetOptions() []GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetOptionsOk() (*[]GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) SetOptions(v []GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


