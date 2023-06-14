# CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Devices** | Pointer to [**[]CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner**](CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner.md) | Data Array of Devices containing the name and serial | [optional] 
**SwitchStacks** | Pointer to [**[]CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesSwitchStacksInner**](CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesSwitchStacksInner.md) | Data Array of Switch Stacks containing the name and id | [optional] 

## Methods

### NewCreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices

`func NewCreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices() *CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices`

NewCreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices instantiates a new CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesWithDefaults

`func NewCreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesWithDefaults() *CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices`

NewCreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesWithDefaults instantiates a new CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevices

`func (o *CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices) GetDevices() []CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices) GetDevicesOk() (*[]CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices) SetDevices(v []CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetSwitchStacks

`func (o *CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices) GetSwitchStacks() []CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesSwitchStacksInner`

GetSwitchStacks returns the SwitchStacks field if non-nil, zero value otherwise.

### GetSwitchStacksOk

`func (o *CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices) GetSwitchStacksOk() (*[]CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesSwitchStacksInner, bool)`

GetSwitchStacksOk returns a tuple with the SwitchStacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchStacks

`func (o *CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices) SetSwitchStacks(v []CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesSwitchStacksInner)`

SetSwitchStacks sets SwitchStacks field to given value.

### HasSwitchStacks

`func (o *CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices) HasSwitchStacks() bool`

HasSwitchStacks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


