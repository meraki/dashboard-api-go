# DevicesSerialSwitchPortsStatusesLldp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemName** | Pointer to **string** | The device&#39;s system name. | [optional] 
**SystemDescription** | Pointer to **string** | The device&#39;s system description. | [optional] 
**ChassisId** | Pointer to **string** | The device&#39;s chassis ID. | [optional] 
**PortId** | Pointer to **string** | Identifies the port from which the LLDP packet was sent | [optional] 
**ManagementVlan** | Pointer to **int32** | The device&#39;s management VLAN. | [optional] 
**PortVlan** | Pointer to **int32** | The port&#39;s VLAN. | [optional] 
**ManagementAddress** | Pointer to **string** | The device&#39;s management IP. | [optional] 
**PortDescription** | Pointer to **string** | Description of the port from which the LLDP packet was sent. | [optional] 
**SystemCapabilities** | Pointer to **string** | Identifies the device type, which indicates the functional capabilities of the device. | [optional] 

## Methods

### NewDevicesSerialSwitchPortsStatusesLldp

`func NewDevicesSerialSwitchPortsStatusesLldp() *DevicesSerialSwitchPortsStatusesLldp`

NewDevicesSerialSwitchPortsStatusesLldp instantiates a new DevicesSerialSwitchPortsStatusesLldp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicesSerialSwitchPortsStatusesLldpWithDefaults

`func NewDevicesSerialSwitchPortsStatusesLldpWithDefaults() *DevicesSerialSwitchPortsStatusesLldp`

NewDevicesSerialSwitchPortsStatusesLldpWithDefaults instantiates a new DevicesSerialSwitchPortsStatusesLldp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSystemName

`func (o *DevicesSerialSwitchPortsStatusesLldp) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *DevicesSerialSwitchPortsStatusesLldp) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *DevicesSerialSwitchPortsStatusesLldp) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.

### HasSystemName

`func (o *DevicesSerialSwitchPortsStatusesLldp) HasSystemName() bool`

HasSystemName returns a boolean if a field has been set.

### GetSystemDescription

`func (o *DevicesSerialSwitchPortsStatusesLldp) GetSystemDescription() string`

GetSystemDescription returns the SystemDescription field if non-nil, zero value otherwise.

### GetSystemDescriptionOk

`func (o *DevicesSerialSwitchPortsStatusesLldp) GetSystemDescriptionOk() (*string, bool)`

GetSystemDescriptionOk returns a tuple with the SystemDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDescription

`func (o *DevicesSerialSwitchPortsStatusesLldp) SetSystemDescription(v string)`

SetSystemDescription sets SystemDescription field to given value.

### HasSystemDescription

`func (o *DevicesSerialSwitchPortsStatusesLldp) HasSystemDescription() bool`

HasSystemDescription returns a boolean if a field has been set.

### GetChassisId

`func (o *DevicesSerialSwitchPortsStatusesLldp) GetChassisId() string`

GetChassisId returns the ChassisId field if non-nil, zero value otherwise.

### GetChassisIdOk

`func (o *DevicesSerialSwitchPortsStatusesLldp) GetChassisIdOk() (*string, bool)`

GetChassisIdOk returns a tuple with the ChassisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisId

`func (o *DevicesSerialSwitchPortsStatusesLldp) SetChassisId(v string)`

SetChassisId sets ChassisId field to given value.

### HasChassisId

`func (o *DevicesSerialSwitchPortsStatusesLldp) HasChassisId() bool`

HasChassisId returns a boolean if a field has been set.

### GetPortId

`func (o *DevicesSerialSwitchPortsStatusesLldp) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *DevicesSerialSwitchPortsStatusesLldp) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *DevicesSerialSwitchPortsStatusesLldp) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *DevicesSerialSwitchPortsStatusesLldp) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetManagementVlan

`func (o *DevicesSerialSwitchPortsStatusesLldp) GetManagementVlan() int32`

GetManagementVlan returns the ManagementVlan field if non-nil, zero value otherwise.

### GetManagementVlanOk

`func (o *DevicesSerialSwitchPortsStatusesLldp) GetManagementVlanOk() (*int32, bool)`

GetManagementVlanOk returns a tuple with the ManagementVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementVlan

`func (o *DevicesSerialSwitchPortsStatusesLldp) SetManagementVlan(v int32)`

SetManagementVlan sets ManagementVlan field to given value.

### HasManagementVlan

`func (o *DevicesSerialSwitchPortsStatusesLldp) HasManagementVlan() bool`

HasManagementVlan returns a boolean if a field has been set.

### GetPortVlan

`func (o *DevicesSerialSwitchPortsStatusesLldp) GetPortVlan() int32`

GetPortVlan returns the PortVlan field if non-nil, zero value otherwise.

### GetPortVlanOk

`func (o *DevicesSerialSwitchPortsStatusesLldp) GetPortVlanOk() (*int32, bool)`

GetPortVlanOk returns a tuple with the PortVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortVlan

`func (o *DevicesSerialSwitchPortsStatusesLldp) SetPortVlan(v int32)`

SetPortVlan sets PortVlan field to given value.

### HasPortVlan

`func (o *DevicesSerialSwitchPortsStatusesLldp) HasPortVlan() bool`

HasPortVlan returns a boolean if a field has been set.

### GetManagementAddress

`func (o *DevicesSerialSwitchPortsStatusesLldp) GetManagementAddress() string`

GetManagementAddress returns the ManagementAddress field if non-nil, zero value otherwise.

### GetManagementAddressOk

`func (o *DevicesSerialSwitchPortsStatusesLldp) GetManagementAddressOk() (*string, bool)`

GetManagementAddressOk returns a tuple with the ManagementAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAddress

`func (o *DevicesSerialSwitchPortsStatusesLldp) SetManagementAddress(v string)`

SetManagementAddress sets ManagementAddress field to given value.

### HasManagementAddress

`func (o *DevicesSerialSwitchPortsStatusesLldp) HasManagementAddress() bool`

HasManagementAddress returns a boolean if a field has been set.

### GetPortDescription

`func (o *DevicesSerialSwitchPortsStatusesLldp) GetPortDescription() string`

GetPortDescription returns the PortDescription field if non-nil, zero value otherwise.

### GetPortDescriptionOk

`func (o *DevicesSerialSwitchPortsStatusesLldp) GetPortDescriptionOk() (*string, bool)`

GetPortDescriptionOk returns a tuple with the PortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortDescription

`func (o *DevicesSerialSwitchPortsStatusesLldp) SetPortDescription(v string)`

SetPortDescription sets PortDescription field to given value.

### HasPortDescription

`func (o *DevicesSerialSwitchPortsStatusesLldp) HasPortDescription() bool`

HasPortDescription returns a boolean if a field has been set.

### GetSystemCapabilities

`func (o *DevicesSerialSwitchPortsStatusesLldp) GetSystemCapabilities() string`

GetSystemCapabilities returns the SystemCapabilities field if non-nil, zero value otherwise.

### GetSystemCapabilitiesOk

`func (o *DevicesSerialSwitchPortsStatusesLldp) GetSystemCapabilitiesOk() (*string, bool)`

GetSystemCapabilitiesOk returns a tuple with the SystemCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemCapabilities

`func (o *DevicesSerialSwitchPortsStatusesLldp) SetSystemCapabilities(v string)`

SetSystemCapabilities sets SystemCapabilities field to given value.

### HasSystemCapabilities

`func (o *DevicesSerialSwitchPortsStatusesLldp) HasSystemCapabilities() bool`

HasSystemCapabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


