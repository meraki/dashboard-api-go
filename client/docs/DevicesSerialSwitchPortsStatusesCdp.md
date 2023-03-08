# DevicesSerialSwitchPortsStatusesCdp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemName** | Pointer to **string** | The system name. | [optional] 
**Platform** | Pointer to **string** | Identifies the hardware platform of the device. | [optional] 
**DeviceId** | Pointer to **string** | Identifies the device name. | [optional] 
**PortId** | Pointer to **string** | Identifies the port from which the CDP packet was sent. | [optional] 
**NativeVlan** | Pointer to **int32** | Indicates, per interface, the assumed VLAN for untagged packets on the interface. | [optional] 
**Address** | Pointer to **string** | Contains network addresses of both receiving and sending devices. | [optional] 
**ManagementAddress** | Pointer to **string** | The device&#39;s management IP. | [optional] 
**Version** | Pointer to **string** | Contains the device software release information. | [optional] 
**VtpManagementDomain** | Pointer to **string** | Advertises the configured VLAN Trunking Protocl (VTP)-management-domain name of the system. | [optional] 
**Capabilities** | Pointer to **string** | Identifies the device type, which indicates the functional capabilities of the device. | [optional] 

## Methods

### NewDevicesSerialSwitchPortsStatusesCdp

`func NewDevicesSerialSwitchPortsStatusesCdp() *DevicesSerialSwitchPortsStatusesCdp`

NewDevicesSerialSwitchPortsStatusesCdp instantiates a new DevicesSerialSwitchPortsStatusesCdp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicesSerialSwitchPortsStatusesCdpWithDefaults

`func NewDevicesSerialSwitchPortsStatusesCdpWithDefaults() *DevicesSerialSwitchPortsStatusesCdp`

NewDevicesSerialSwitchPortsStatusesCdpWithDefaults instantiates a new DevicesSerialSwitchPortsStatusesCdp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSystemName

`func (o *DevicesSerialSwitchPortsStatusesCdp) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *DevicesSerialSwitchPortsStatusesCdp) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *DevicesSerialSwitchPortsStatusesCdp) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.

### HasSystemName

`func (o *DevicesSerialSwitchPortsStatusesCdp) HasSystemName() bool`

HasSystemName returns a boolean if a field has been set.

### GetPlatform

`func (o *DevicesSerialSwitchPortsStatusesCdp) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *DevicesSerialSwitchPortsStatusesCdp) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *DevicesSerialSwitchPortsStatusesCdp) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *DevicesSerialSwitchPortsStatusesCdp) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetDeviceId

`func (o *DevicesSerialSwitchPortsStatusesCdp) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *DevicesSerialSwitchPortsStatusesCdp) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *DevicesSerialSwitchPortsStatusesCdp) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *DevicesSerialSwitchPortsStatusesCdp) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetPortId

`func (o *DevicesSerialSwitchPortsStatusesCdp) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *DevicesSerialSwitchPortsStatusesCdp) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *DevicesSerialSwitchPortsStatusesCdp) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *DevicesSerialSwitchPortsStatusesCdp) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetNativeVlan

`func (o *DevicesSerialSwitchPortsStatusesCdp) GetNativeVlan() int32`

GetNativeVlan returns the NativeVlan field if non-nil, zero value otherwise.

### GetNativeVlanOk

`func (o *DevicesSerialSwitchPortsStatusesCdp) GetNativeVlanOk() (*int32, bool)`

GetNativeVlanOk returns a tuple with the NativeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeVlan

`func (o *DevicesSerialSwitchPortsStatusesCdp) SetNativeVlan(v int32)`

SetNativeVlan sets NativeVlan field to given value.

### HasNativeVlan

`func (o *DevicesSerialSwitchPortsStatusesCdp) HasNativeVlan() bool`

HasNativeVlan returns a boolean if a field has been set.

### GetAddress

`func (o *DevicesSerialSwitchPortsStatusesCdp) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DevicesSerialSwitchPortsStatusesCdp) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DevicesSerialSwitchPortsStatusesCdp) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DevicesSerialSwitchPortsStatusesCdp) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetManagementAddress

`func (o *DevicesSerialSwitchPortsStatusesCdp) GetManagementAddress() string`

GetManagementAddress returns the ManagementAddress field if non-nil, zero value otherwise.

### GetManagementAddressOk

`func (o *DevicesSerialSwitchPortsStatusesCdp) GetManagementAddressOk() (*string, bool)`

GetManagementAddressOk returns a tuple with the ManagementAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAddress

`func (o *DevicesSerialSwitchPortsStatusesCdp) SetManagementAddress(v string)`

SetManagementAddress sets ManagementAddress field to given value.

### HasManagementAddress

`func (o *DevicesSerialSwitchPortsStatusesCdp) HasManagementAddress() bool`

HasManagementAddress returns a boolean if a field has been set.

### GetVersion

`func (o *DevicesSerialSwitchPortsStatusesCdp) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DevicesSerialSwitchPortsStatusesCdp) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DevicesSerialSwitchPortsStatusesCdp) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DevicesSerialSwitchPortsStatusesCdp) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVtpManagementDomain

`func (o *DevicesSerialSwitchPortsStatusesCdp) GetVtpManagementDomain() string`

GetVtpManagementDomain returns the VtpManagementDomain field if non-nil, zero value otherwise.

### GetVtpManagementDomainOk

`func (o *DevicesSerialSwitchPortsStatusesCdp) GetVtpManagementDomainOk() (*string, bool)`

GetVtpManagementDomainOk returns a tuple with the VtpManagementDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVtpManagementDomain

`func (o *DevicesSerialSwitchPortsStatusesCdp) SetVtpManagementDomain(v string)`

SetVtpManagementDomain sets VtpManagementDomain field to given value.

### HasVtpManagementDomain

`func (o *DevicesSerialSwitchPortsStatusesCdp) HasVtpManagementDomain() bool`

HasVtpManagementDomain returns a boolean if a field has been set.

### GetCapabilities

`func (o *DevicesSerialSwitchPortsStatusesCdp) GetCapabilities() string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *DevicesSerialSwitchPortsStatusesCdp) GetCapabilitiesOk() (*string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *DevicesSerialSwitchPortsStatusesCdp) SetCapabilities(v string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *DevicesSerialSwitchPortsStatusesCdp) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


