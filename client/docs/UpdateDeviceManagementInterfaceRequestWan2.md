# UpdateDeviceManagementInterfaceRequestWan2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WanEnabled** | Pointer to **string** | Enable or disable the interface (only for MX devices). Valid values are &#39;enabled&#39;, &#39;disabled&#39;, and &#39;not configured&#39;. | [optional] 
**UsingStaticIp** | Pointer to **bool** | Configure the interface to have static IP settings or use DHCP. | [optional] 
**StaticIp** | Pointer to **string** | The IP the device should use on the WAN. | [optional] 
**StaticGatewayIp** | Pointer to **string** | The IP of the gateway on the WAN. | [optional] 
**StaticSubnetMask** | Pointer to **string** | The subnet mask for the WAN. | [optional] 
**StaticDns** | Pointer to **[]string** | Up to two DNS IPs. | [optional] 
**Vlan** | Pointer to **int32** | The VLAN that management traffic should be tagged with. Applies whether usingStaticIp is true or false. | [optional] 

## Methods

### NewUpdateDeviceManagementInterfaceRequestWan2

`func NewUpdateDeviceManagementInterfaceRequestWan2() *UpdateDeviceManagementInterfaceRequestWan2`

NewUpdateDeviceManagementInterfaceRequestWan2 instantiates a new UpdateDeviceManagementInterfaceRequestWan2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeviceManagementInterfaceRequestWan2WithDefaults

`func NewUpdateDeviceManagementInterfaceRequestWan2WithDefaults() *UpdateDeviceManagementInterfaceRequestWan2`

NewUpdateDeviceManagementInterfaceRequestWan2WithDefaults instantiates a new UpdateDeviceManagementInterfaceRequestWan2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWanEnabled

`func (o *UpdateDeviceManagementInterfaceRequestWan2) GetWanEnabled() string`

GetWanEnabled returns the WanEnabled field if non-nil, zero value otherwise.

### GetWanEnabledOk

`func (o *UpdateDeviceManagementInterfaceRequestWan2) GetWanEnabledOk() (*string, bool)`

GetWanEnabledOk returns a tuple with the WanEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanEnabled

`func (o *UpdateDeviceManagementInterfaceRequestWan2) SetWanEnabled(v string)`

SetWanEnabled sets WanEnabled field to given value.

### HasWanEnabled

`func (o *UpdateDeviceManagementInterfaceRequestWan2) HasWanEnabled() bool`

HasWanEnabled returns a boolean if a field has been set.

### GetUsingStaticIp

`func (o *UpdateDeviceManagementInterfaceRequestWan2) GetUsingStaticIp() bool`

GetUsingStaticIp returns the UsingStaticIp field if non-nil, zero value otherwise.

### GetUsingStaticIpOk

`func (o *UpdateDeviceManagementInterfaceRequestWan2) GetUsingStaticIpOk() (*bool, bool)`

GetUsingStaticIpOk returns a tuple with the UsingStaticIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsingStaticIp

`func (o *UpdateDeviceManagementInterfaceRequestWan2) SetUsingStaticIp(v bool)`

SetUsingStaticIp sets UsingStaticIp field to given value.

### HasUsingStaticIp

`func (o *UpdateDeviceManagementInterfaceRequestWan2) HasUsingStaticIp() bool`

HasUsingStaticIp returns a boolean if a field has been set.

### GetStaticIp

`func (o *UpdateDeviceManagementInterfaceRequestWan2) GetStaticIp() string`

GetStaticIp returns the StaticIp field if non-nil, zero value otherwise.

### GetStaticIpOk

`func (o *UpdateDeviceManagementInterfaceRequestWan2) GetStaticIpOk() (*string, bool)`

GetStaticIpOk returns a tuple with the StaticIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticIp

`func (o *UpdateDeviceManagementInterfaceRequestWan2) SetStaticIp(v string)`

SetStaticIp sets StaticIp field to given value.

### HasStaticIp

`func (o *UpdateDeviceManagementInterfaceRequestWan2) HasStaticIp() bool`

HasStaticIp returns a boolean if a field has been set.

### GetStaticGatewayIp

`func (o *UpdateDeviceManagementInterfaceRequestWan2) GetStaticGatewayIp() string`

GetStaticGatewayIp returns the StaticGatewayIp field if non-nil, zero value otherwise.

### GetStaticGatewayIpOk

`func (o *UpdateDeviceManagementInterfaceRequestWan2) GetStaticGatewayIpOk() (*string, bool)`

GetStaticGatewayIpOk returns a tuple with the StaticGatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticGatewayIp

`func (o *UpdateDeviceManagementInterfaceRequestWan2) SetStaticGatewayIp(v string)`

SetStaticGatewayIp sets StaticGatewayIp field to given value.

### HasStaticGatewayIp

`func (o *UpdateDeviceManagementInterfaceRequestWan2) HasStaticGatewayIp() bool`

HasStaticGatewayIp returns a boolean if a field has been set.

### GetStaticSubnetMask

`func (o *UpdateDeviceManagementInterfaceRequestWan2) GetStaticSubnetMask() string`

GetStaticSubnetMask returns the StaticSubnetMask field if non-nil, zero value otherwise.

### GetStaticSubnetMaskOk

`func (o *UpdateDeviceManagementInterfaceRequestWan2) GetStaticSubnetMaskOk() (*string, bool)`

GetStaticSubnetMaskOk returns a tuple with the StaticSubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticSubnetMask

`func (o *UpdateDeviceManagementInterfaceRequestWan2) SetStaticSubnetMask(v string)`

SetStaticSubnetMask sets StaticSubnetMask field to given value.

### HasStaticSubnetMask

`func (o *UpdateDeviceManagementInterfaceRequestWan2) HasStaticSubnetMask() bool`

HasStaticSubnetMask returns a boolean if a field has been set.

### GetStaticDns

`func (o *UpdateDeviceManagementInterfaceRequestWan2) GetStaticDns() []string`

GetStaticDns returns the StaticDns field if non-nil, zero value otherwise.

### GetStaticDnsOk

`func (o *UpdateDeviceManagementInterfaceRequestWan2) GetStaticDnsOk() (*[]string, bool)`

GetStaticDnsOk returns a tuple with the StaticDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticDns

`func (o *UpdateDeviceManagementInterfaceRequestWan2) SetStaticDns(v []string)`

SetStaticDns sets StaticDns field to given value.

### HasStaticDns

`func (o *UpdateDeviceManagementInterfaceRequestWan2) HasStaticDns() bool`

HasStaticDns returns a boolean if a field has been set.

### GetVlan

`func (o *UpdateDeviceManagementInterfaceRequestWan2) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *UpdateDeviceManagementInterfaceRequestWan2) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *UpdateDeviceManagementInterfaceRequestWan2) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *UpdateDeviceManagementInterfaceRequestWan2) HasVlan() bool`

HasVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


