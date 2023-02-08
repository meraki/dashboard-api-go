# GetDeviceSwitchRoutingInterfaces200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterfaceId** | Pointer to **string** | The id | [optional] 
**Name** | Pointer to **string** | The name | [optional] 
**Subnet** | Pointer to **string** | IPv4 subnet | [optional] 
**InterfaceIp** | Pointer to **string** | IPv4 address | [optional] 
**MulticastRouting** | Pointer to **string** | Multicast routing status | [optional] 
**VlanId** | Pointer to **int32** | VLAN id | [optional] 
**DefaultGateway** | Pointer to **string** | IPv4 default gateway | [optional] 
**OspfSettings** | Pointer to [**GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings**](GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings.md) |  | [optional] 
**OspfV3** | Pointer to [**GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3**](GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3.md) |  | [optional] 
**Ipv6** | Pointer to [**GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6**](GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6.md) |  | [optional] 

## Methods

### NewGetDeviceSwitchRoutingInterfaces200ResponseInner

`func NewGetDeviceSwitchRoutingInterfaces200ResponseInner() *GetDeviceSwitchRoutingInterfaces200ResponseInner`

NewGetDeviceSwitchRoutingInterfaces200ResponseInner instantiates a new GetDeviceSwitchRoutingInterfaces200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeviceSwitchRoutingInterfaces200ResponseInnerWithDefaults

`func NewGetDeviceSwitchRoutingInterfaces200ResponseInnerWithDefaults() *GetDeviceSwitchRoutingInterfaces200ResponseInner`

NewGetDeviceSwitchRoutingInterfaces200ResponseInnerWithDefaults instantiates a new GetDeviceSwitchRoutingInterfaces200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterfaceId

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) GetInterfaceId() string`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) GetInterfaceIdOk() (*string, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) SetInterfaceId(v string)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### GetName

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubnet

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetInterfaceIp

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) GetInterfaceIp() string`

GetInterfaceIp returns the InterfaceIp field if non-nil, zero value otherwise.

### GetInterfaceIpOk

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) GetInterfaceIpOk() (*string, bool)`

GetInterfaceIpOk returns a tuple with the InterfaceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceIp

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) SetInterfaceIp(v string)`

SetInterfaceIp sets InterfaceIp field to given value.

### HasInterfaceIp

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) HasInterfaceIp() bool`

HasInterfaceIp returns a boolean if a field has been set.

### GetMulticastRouting

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) GetMulticastRouting() string`

GetMulticastRouting returns the MulticastRouting field if non-nil, zero value otherwise.

### GetMulticastRoutingOk

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) GetMulticastRoutingOk() (*string, bool)`

GetMulticastRoutingOk returns a tuple with the MulticastRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticastRouting

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) SetMulticastRouting(v string)`

SetMulticastRouting sets MulticastRouting field to given value.

### HasMulticastRouting

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) HasMulticastRouting() bool`

HasMulticastRouting returns a boolean if a field has been set.

### GetVlanId

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetDefaultGateway

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) GetDefaultGateway() string`

GetDefaultGateway returns the DefaultGateway field if non-nil, zero value otherwise.

### GetDefaultGatewayOk

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) GetDefaultGatewayOk() (*string, bool)`

GetDefaultGatewayOk returns a tuple with the DefaultGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGateway

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) SetDefaultGateway(v string)`

SetDefaultGateway sets DefaultGateway field to given value.

### HasDefaultGateway

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) HasDefaultGateway() bool`

HasDefaultGateway returns a boolean if a field has been set.

### GetOspfSettings

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) GetOspfSettings() GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings`

GetOspfSettings returns the OspfSettings field if non-nil, zero value otherwise.

### GetOspfSettingsOk

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) GetOspfSettingsOk() (*GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings, bool)`

GetOspfSettingsOk returns a tuple with the OspfSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfSettings

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) SetOspfSettings(v GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings)`

SetOspfSettings sets OspfSettings field to given value.

### HasOspfSettings

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) HasOspfSettings() bool`

HasOspfSettings returns a boolean if a field has been set.

### GetOspfV3

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) GetOspfV3() GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3`

GetOspfV3 returns the OspfV3 field if non-nil, zero value otherwise.

### GetOspfV3Ok

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) GetOspfV3Ok() (*GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3, bool)`

GetOspfV3Ok returns a tuple with the OspfV3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfV3

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) SetOspfV3(v GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3)`

SetOspfV3 sets OspfV3 field to given value.

### HasOspfV3

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) HasOspfV3() bool`

HasOspfV3 returns a boolean if a field has been set.

### GetIpv6

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) GetIpv6() GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) GetIpv6Ok() (*GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) SetIpv6(v GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


