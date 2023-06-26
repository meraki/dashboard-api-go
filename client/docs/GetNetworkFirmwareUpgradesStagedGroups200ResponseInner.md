# GetNetworkFirmwareUpgradesStagedGroups200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Id of staged upgrade group | [optional] 
**Name** | Pointer to **string** | Name of the Staged Upgrade Group | [optional] 
**Description** | Pointer to **string** | Description of the Staged Upgrade Group | [optional] 
**IsDefault** | Pointer to **bool** | Boolean indicating the default Group. Any device that does not have a group explicitly assigned will upgrade with this group | [optional] 
**AssignedDevices** | Pointer to [**GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices**](GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices.md) |  | [optional] 

## Methods

### NewGetNetworkFirmwareUpgradesStagedGroups200ResponseInner

`func NewGetNetworkFirmwareUpgradesStagedGroups200ResponseInner() *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner`

NewGetNetworkFirmwareUpgradesStagedGroups200ResponseInner instantiates a new GetNetworkFirmwareUpgradesStagedGroups200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerWithDefaults

`func NewGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerWithDefaults() *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner`

NewGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerWithDefaults instantiates a new GetNetworkFirmwareUpgradesStagedGroups200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsDefault

`func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetAssignedDevices

`func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) GetAssignedDevices() GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices`

GetAssignedDevices returns the AssignedDevices field if non-nil, zero value otherwise.

### GetAssignedDevicesOk

`func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) GetAssignedDevicesOk() (*GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices, bool)`

GetAssignedDevicesOk returns a tuple with the AssignedDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedDevices

`func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) SetAssignedDevices(v GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevices)`

SetAssignedDevices sets AssignedDevices field to given value.

### HasAssignedDevices

`func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInner) HasAssignedDevices() bool`

HasAssignedDevices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


