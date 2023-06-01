# InlineResponse20030

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Id of staged upgrade group | [optional] 
**Name** | Pointer to **string** | Name of the Staged Upgrade Group | [optional] 
**Description** | Pointer to **string** | Description of the Staged Upgrade Group | [optional] 
**IsDefault** | Pointer to **bool** | Boolean indicating the default Group. Any device that does not have a group explicitly assigned will upgrade with this group | [optional] 
**AssignedDevices** | Pointer to [**NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices**](NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices.md) |  | [optional] 

## Methods

### NewInlineResponse20030

`func NewInlineResponse20030() *InlineResponse20030`

NewInlineResponse20030 instantiates a new InlineResponse20030 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20030WithDefaults

`func NewInlineResponse20030WithDefaults() *InlineResponse20030`

NewInlineResponse20030WithDefaults instantiates a new InlineResponse20030 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *InlineResponse20030) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *InlineResponse20030) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *InlineResponse20030) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *InlineResponse20030) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20030) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20030) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20030) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20030) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse20030) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse20030) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse20030) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse20030) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsDefault

`func (o *InlineResponse20030) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *InlineResponse20030) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *InlineResponse20030) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *InlineResponse20030) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetAssignedDevices

`func (o *InlineResponse20030) GetAssignedDevices() NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices`

GetAssignedDevices returns the AssignedDevices field if non-nil, zero value otherwise.

### GetAssignedDevicesOk

`func (o *InlineResponse20030) GetAssignedDevicesOk() (*NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices, bool)`

GetAssignedDevicesOk returns a tuple with the AssignedDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedDevices

`func (o *InlineResponse20030) SetAssignedDevices(v NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices)`

SetAssignedDevices sets AssignedDevices field to given value.

### HasAssignedDevices

`func (o *InlineResponse20030) HasAssignedDevices() bool`

HasAssignedDevices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


