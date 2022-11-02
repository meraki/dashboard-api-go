# CreateNetworkFirmwareUpgradesStagedGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the Staged Upgrade Group. Length must be 1 to 255 characters | 
**Description** | Pointer to **string** | Description of the Staged Upgrade Group. Length must be 1 to 255 characters | [optional] 
**IsDefault** | **bool** | Boolean indicating the default Group. Any device that does not have a group explicitly assigned will upgrade with this group | 
**AssignedDevices** | Pointer to [**CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices**](CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices.md) |  | [optional] 

## Methods

### NewCreateNetworkFirmwareUpgradesStagedGroupRequest

`func NewCreateNetworkFirmwareUpgradesStagedGroupRequest(name string, isDefault bool, ) *CreateNetworkFirmwareUpgradesStagedGroupRequest`

NewCreateNetworkFirmwareUpgradesStagedGroupRequest instantiates a new CreateNetworkFirmwareUpgradesStagedGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkFirmwareUpgradesStagedGroupRequestWithDefaults

`func NewCreateNetworkFirmwareUpgradesStagedGroupRequestWithDefaults() *CreateNetworkFirmwareUpgradesStagedGroupRequest`

NewCreateNetworkFirmwareUpgradesStagedGroupRequestWithDefaults instantiates a new CreateNetworkFirmwareUpgradesStagedGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNetworkFirmwareUpgradesStagedGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkFirmwareUpgradesStagedGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkFirmwareUpgradesStagedGroupRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateNetworkFirmwareUpgradesStagedGroupRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateNetworkFirmwareUpgradesStagedGroupRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateNetworkFirmwareUpgradesStagedGroupRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateNetworkFirmwareUpgradesStagedGroupRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsDefault

`func (o *CreateNetworkFirmwareUpgradesStagedGroupRequest) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *CreateNetworkFirmwareUpgradesStagedGroupRequest) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *CreateNetworkFirmwareUpgradesStagedGroupRequest) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetAssignedDevices

`func (o *CreateNetworkFirmwareUpgradesStagedGroupRequest) GetAssignedDevices() CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices`

GetAssignedDevices returns the AssignedDevices field if non-nil, zero value otherwise.

### GetAssignedDevicesOk

`func (o *CreateNetworkFirmwareUpgradesStagedGroupRequest) GetAssignedDevicesOk() (*CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices, bool)`

GetAssignedDevicesOk returns a tuple with the AssignedDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedDevices

`func (o *CreateNetworkFirmwareUpgradesStagedGroupRequest) SetAssignedDevices(v CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices)`

SetAssignedDevices sets AssignedDevices field to given value.

### HasAssignedDevices

`func (o *CreateNetworkFirmwareUpgradesStagedGroupRequest) HasAssignedDevices() bool`

HasAssignedDevices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


