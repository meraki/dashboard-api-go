# GetDeviceSwitchPortsStatuses200ResponseInnerLldp

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

### NewGetDeviceSwitchPortsStatuses200ResponseInnerLldp

`func NewGetDeviceSwitchPortsStatuses200ResponseInnerLldp() *GetDeviceSwitchPortsStatuses200ResponseInnerLldp`

NewGetDeviceSwitchPortsStatuses200ResponseInnerLldp instantiates a new GetDeviceSwitchPortsStatuses200ResponseInnerLldp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeviceSwitchPortsStatuses200ResponseInnerLldpWithDefaults

`func NewGetDeviceSwitchPortsStatuses200ResponseInnerLldpWithDefaults() *GetDeviceSwitchPortsStatuses200ResponseInnerLldp`

NewGetDeviceSwitchPortsStatuses200ResponseInnerLldpWithDefaults instantiates a new GetDeviceSwitchPortsStatuses200ResponseInnerLldp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSystemName

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.

### HasSystemName

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) HasSystemName() bool`

HasSystemName returns a boolean if a field has been set.

### GetSystemDescription

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) GetSystemDescription() string`

GetSystemDescription returns the SystemDescription field if non-nil, zero value otherwise.

### GetSystemDescriptionOk

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) GetSystemDescriptionOk() (*string, bool)`

GetSystemDescriptionOk returns a tuple with the SystemDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDescription

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) SetSystemDescription(v string)`

SetSystemDescription sets SystemDescription field to given value.

### HasSystemDescription

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) HasSystemDescription() bool`

HasSystemDescription returns a boolean if a field has been set.

### GetChassisId

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) GetChassisId() string`

GetChassisId returns the ChassisId field if non-nil, zero value otherwise.

### GetChassisIdOk

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) GetChassisIdOk() (*string, bool)`

GetChassisIdOk returns a tuple with the ChassisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisId

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) SetChassisId(v string)`

SetChassisId sets ChassisId field to given value.

### HasChassisId

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) HasChassisId() bool`

HasChassisId returns a boolean if a field has been set.

### GetPortId

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetManagementVlan

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) GetManagementVlan() int32`

GetManagementVlan returns the ManagementVlan field if non-nil, zero value otherwise.

### GetManagementVlanOk

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) GetManagementVlanOk() (*int32, bool)`

GetManagementVlanOk returns a tuple with the ManagementVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementVlan

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) SetManagementVlan(v int32)`

SetManagementVlan sets ManagementVlan field to given value.

### HasManagementVlan

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) HasManagementVlan() bool`

HasManagementVlan returns a boolean if a field has been set.

### GetPortVlan

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) GetPortVlan() int32`

GetPortVlan returns the PortVlan field if non-nil, zero value otherwise.

### GetPortVlanOk

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) GetPortVlanOk() (*int32, bool)`

GetPortVlanOk returns a tuple with the PortVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortVlan

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) SetPortVlan(v int32)`

SetPortVlan sets PortVlan field to given value.

### HasPortVlan

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) HasPortVlan() bool`

HasPortVlan returns a boolean if a field has been set.

### GetManagementAddress

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) GetManagementAddress() string`

GetManagementAddress returns the ManagementAddress field if non-nil, zero value otherwise.

### GetManagementAddressOk

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) GetManagementAddressOk() (*string, bool)`

GetManagementAddressOk returns a tuple with the ManagementAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAddress

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) SetManagementAddress(v string)`

SetManagementAddress sets ManagementAddress field to given value.

### HasManagementAddress

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) HasManagementAddress() bool`

HasManagementAddress returns a boolean if a field has been set.

### GetPortDescription

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) GetPortDescription() string`

GetPortDescription returns the PortDescription field if non-nil, zero value otherwise.

### GetPortDescriptionOk

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) GetPortDescriptionOk() (*string, bool)`

GetPortDescriptionOk returns a tuple with the PortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortDescription

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) SetPortDescription(v string)`

SetPortDescription sets PortDescription field to given value.

### HasPortDescription

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) HasPortDescription() bool`

HasPortDescription returns a boolean if a field has been set.

### GetSystemCapabilities

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) GetSystemCapabilities() string`

GetSystemCapabilities returns the SystemCapabilities field if non-nil, zero value otherwise.

### GetSystemCapabilitiesOk

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) GetSystemCapabilitiesOk() (*string, bool)`

GetSystemCapabilitiesOk returns a tuple with the SystemCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemCapabilities

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) SetSystemCapabilities(v string)`

SetSystemCapabilities sets SystemCapabilities field to given value.

### HasSystemCapabilities

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerLldp) HasSystemCapabilities() bool`

HasSystemCapabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


