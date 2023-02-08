# CreateDeviceSwitchRoutingInterfaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A friendly name or description for the interface or VLAN. | [optional] 
**Subnet** | Pointer to **string** | The network that this routed interface is on, in CIDR notation (ex. 10.1.1.0/24). | [optional] 
**InterfaceIp** | Pointer to **string** | The IP address this switch will use for layer 3 routing on this VLAN or subnet. This cannot be the same         as the switch&#39;s management IP. | [optional] 
**MulticastRouting** | Pointer to **string** | Enable multicast support if, multicast routing between VLANs is required. Options are:         &#39;disabled&#39;, &#39;enabled&#39; or &#39;IGMP snooping querier&#39;. Default is &#39;disabled&#39;. | [optional] 
**VlanId** | Pointer to **int32** | The VLAN this routed interface is on. VLAN must be between 1 and 4094. | [optional] 
**DefaultGateway** | Pointer to **string** | The next hop for any traffic that isn&#39;t going to a directly connected subnet or over a static route.         This IP address must exist in a subnet with a routed interface. Required if this is the first IPv4 interface. | [optional] 
**OspfSettings** | Pointer to [**CreateDeviceSwitchRoutingInterfaceRequestOspfSettings**](CreateDeviceSwitchRoutingInterfaceRequestOspfSettings.md) |  | [optional] 
**OspfV3** | Pointer to [**CreateDeviceSwitchRoutingInterfaceRequestOspfV3**](CreateDeviceSwitchRoutingInterfaceRequestOspfV3.md) |  | [optional] 
**Ipv6** | Pointer to [**CreateDeviceSwitchRoutingInterfaceRequestIpv6**](CreateDeviceSwitchRoutingInterfaceRequestIpv6.md) |  | [optional] 

## Methods

### NewCreateDeviceSwitchRoutingInterfaceRequest

`func NewCreateDeviceSwitchRoutingInterfaceRequest() *CreateDeviceSwitchRoutingInterfaceRequest`

NewCreateDeviceSwitchRoutingInterfaceRequest instantiates a new CreateDeviceSwitchRoutingInterfaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDeviceSwitchRoutingInterfaceRequestWithDefaults

`func NewCreateDeviceSwitchRoutingInterfaceRequestWithDefaults() *CreateDeviceSwitchRoutingInterfaceRequest`

NewCreateDeviceSwitchRoutingInterfaceRequestWithDefaults instantiates a new CreateDeviceSwitchRoutingInterfaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateDeviceSwitchRoutingInterfaceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDeviceSwitchRoutingInterfaceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDeviceSwitchRoutingInterfaceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateDeviceSwitchRoutingInterfaceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubnet

`func (o *CreateDeviceSwitchRoutingInterfaceRequest) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *CreateDeviceSwitchRoutingInterfaceRequest) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *CreateDeviceSwitchRoutingInterfaceRequest) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *CreateDeviceSwitchRoutingInterfaceRequest) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetInterfaceIp

`func (o *CreateDeviceSwitchRoutingInterfaceRequest) GetInterfaceIp() string`

GetInterfaceIp returns the InterfaceIp field if non-nil, zero value otherwise.

### GetInterfaceIpOk

`func (o *CreateDeviceSwitchRoutingInterfaceRequest) GetInterfaceIpOk() (*string, bool)`

GetInterfaceIpOk returns a tuple with the InterfaceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceIp

`func (o *CreateDeviceSwitchRoutingInterfaceRequest) SetInterfaceIp(v string)`

SetInterfaceIp sets InterfaceIp field to given value.

### HasInterfaceIp

`func (o *CreateDeviceSwitchRoutingInterfaceRequest) HasInterfaceIp() bool`

HasInterfaceIp returns a boolean if a field has been set.

### GetMulticastRouting

`func (o *CreateDeviceSwitchRoutingInterfaceRequest) GetMulticastRouting() string`

GetMulticastRouting returns the MulticastRouting field if non-nil, zero value otherwise.

### GetMulticastRoutingOk

`func (o *CreateDeviceSwitchRoutingInterfaceRequest) GetMulticastRoutingOk() (*string, bool)`

GetMulticastRoutingOk returns a tuple with the MulticastRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticastRouting

`func (o *CreateDeviceSwitchRoutingInterfaceRequest) SetMulticastRouting(v string)`

SetMulticastRouting sets MulticastRouting field to given value.

### HasMulticastRouting

`func (o *CreateDeviceSwitchRoutingInterfaceRequest) HasMulticastRouting() bool`

HasMulticastRouting returns a boolean if a field has been set.

### GetVlanId

`func (o *CreateDeviceSwitchRoutingInterfaceRequest) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *CreateDeviceSwitchRoutingInterfaceRequest) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *CreateDeviceSwitchRoutingInterfaceRequest) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *CreateDeviceSwitchRoutingInterfaceRequest) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetDefaultGateway

`func (o *CreateDeviceSwitchRoutingInterfaceRequest) GetDefaultGateway() string`

GetDefaultGateway returns the DefaultGateway field if non-nil, zero value otherwise.

### GetDefaultGatewayOk

`func (o *CreateDeviceSwitchRoutingInterfaceRequest) GetDefaultGatewayOk() (*string, bool)`

GetDefaultGatewayOk returns a tuple with the DefaultGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGateway

`func (o *CreateDeviceSwitchRoutingInterfaceRequest) SetDefaultGateway(v string)`

SetDefaultGateway sets DefaultGateway field to given value.

### HasDefaultGateway

`func (o *CreateDeviceSwitchRoutingInterfaceRequest) HasDefaultGateway() bool`

HasDefaultGateway returns a boolean if a field has been set.

### GetOspfSettings

`func (o *CreateDeviceSwitchRoutingInterfaceRequest) GetOspfSettings() CreateDeviceSwitchRoutingInterfaceRequestOspfSettings`

GetOspfSettings returns the OspfSettings field if non-nil, zero value otherwise.

### GetOspfSettingsOk

`func (o *CreateDeviceSwitchRoutingInterfaceRequest) GetOspfSettingsOk() (*CreateDeviceSwitchRoutingInterfaceRequestOspfSettings, bool)`

GetOspfSettingsOk returns a tuple with the OspfSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfSettings

`func (o *CreateDeviceSwitchRoutingInterfaceRequest) SetOspfSettings(v CreateDeviceSwitchRoutingInterfaceRequestOspfSettings)`

SetOspfSettings sets OspfSettings field to given value.

### HasOspfSettings

`func (o *CreateDeviceSwitchRoutingInterfaceRequest) HasOspfSettings() bool`

HasOspfSettings returns a boolean if a field has been set.

### GetOspfV3

`func (o *CreateDeviceSwitchRoutingInterfaceRequest) GetOspfV3() CreateDeviceSwitchRoutingInterfaceRequestOspfV3`

GetOspfV3 returns the OspfV3 field if non-nil, zero value otherwise.

### GetOspfV3Ok

`func (o *CreateDeviceSwitchRoutingInterfaceRequest) GetOspfV3Ok() (*CreateDeviceSwitchRoutingInterfaceRequestOspfV3, bool)`

GetOspfV3Ok returns a tuple with the OspfV3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfV3

`func (o *CreateDeviceSwitchRoutingInterfaceRequest) SetOspfV3(v CreateDeviceSwitchRoutingInterfaceRequestOspfV3)`

SetOspfV3 sets OspfV3 field to given value.

### HasOspfV3

`func (o *CreateDeviceSwitchRoutingInterfaceRequest) HasOspfV3() bool`

HasOspfV3 returns a boolean if a field has been set.

### GetIpv6

`func (o *CreateDeviceSwitchRoutingInterfaceRequest) GetIpv6() CreateDeviceSwitchRoutingInterfaceRequestIpv6`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *CreateDeviceSwitchRoutingInterfaceRequest) GetIpv6Ok() (*CreateDeviceSwitchRoutingInterfaceRequestIpv6, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *CreateDeviceSwitchRoutingInterfaceRequest) SetIpv6(v CreateDeviceSwitchRoutingInterfaceRequestIpv6)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *CreateDeviceSwitchRoutingInterfaceRequest) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


