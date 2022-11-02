# GetOrganizationFirmwareUpgradesByDevice200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Serial of the device | [optional] 
**Name** | Pointer to **string** | Name assigned to the device | [optional] 
**DeviceStatus** | Pointer to **string** | Status of the device upgrade | [optional] 
**Upgrade** | Pointer to [**GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade**](GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade.md) |  | [optional] 

## Methods

### NewGetOrganizationFirmwareUpgradesByDevice200ResponseInner

`func NewGetOrganizationFirmwareUpgradesByDevice200ResponseInner() *GetOrganizationFirmwareUpgradesByDevice200ResponseInner`

NewGetOrganizationFirmwareUpgradesByDevice200ResponseInner instantiates a new GetOrganizationFirmwareUpgradesByDevice200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationFirmwareUpgradesByDevice200ResponseInnerWithDefaults

`func NewGetOrganizationFirmwareUpgradesByDevice200ResponseInnerWithDefaults() *GetOrganizationFirmwareUpgradesByDevice200ResponseInner`

NewGetOrganizationFirmwareUpgradesByDevice200ResponseInnerWithDefaults instantiates a new GetOrganizationFirmwareUpgradesByDevice200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInner) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInner) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInner) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInner) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetName

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDeviceStatus

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInner) GetDeviceStatus() string`

GetDeviceStatus returns the DeviceStatus field if non-nil, zero value otherwise.

### GetDeviceStatusOk

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInner) GetDeviceStatusOk() (*string, bool)`

GetDeviceStatusOk returns a tuple with the DeviceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceStatus

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInner) SetDeviceStatus(v string)`

SetDeviceStatus sets DeviceStatus field to given value.

### HasDeviceStatus

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInner) HasDeviceStatus() bool`

HasDeviceStatus returns a boolean if a field has been set.

### GetUpgrade

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInner) GetUpgrade() GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade`

GetUpgrade returns the Upgrade field if non-nil, zero value otherwise.

### GetUpgradeOk

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInner) GetUpgradeOk() (*GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade, bool)`

GetUpgradeOk returns a tuple with the Upgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgrade

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInner) SetUpgrade(v GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade)`

SetUpgrade sets Upgrade field to given value.

### HasUpgrade

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInner) HasUpgrade() bool`

HasUpgrade returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


