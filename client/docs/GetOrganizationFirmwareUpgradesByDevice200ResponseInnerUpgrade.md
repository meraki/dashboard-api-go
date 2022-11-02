# GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to **string** | Start time of the upgrade | [optional] 
**FromVersion** | Pointer to [**GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeFromVersion**](GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeFromVersion.md) |  | [optional] 
**ToVersion** | Pointer to [**GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion**](GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the upgrade | [optional] 
**Id** | Pointer to **string** | ID of the upgrade | [optional] 
**UpgradeBatchId** | Pointer to **string** | ID of the upgrade batch | [optional] 
**Staged** | Pointer to [**GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStaged**](GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStaged.md) |  | [optional] 

## Methods

### NewGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade

`func NewGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade() *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade`

NewGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade instantiates a new GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeWithDefaults

`func NewGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeWithDefaults() *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade`

NewGetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeWithDefaults instantiates a new GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetFromVersion

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) GetFromVersion() GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeFromVersion`

GetFromVersion returns the FromVersion field if non-nil, zero value otherwise.

### GetFromVersionOk

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) GetFromVersionOk() (*GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeFromVersion, bool)`

GetFromVersionOk returns a tuple with the FromVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromVersion

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) SetFromVersion(v GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeFromVersion)`

SetFromVersion sets FromVersion field to given value.

### HasFromVersion

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) HasFromVersion() bool`

HasFromVersion returns a boolean if a field has been set.

### GetToVersion

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) GetToVersion() GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion`

GetToVersion returns the ToVersion field if non-nil, zero value otherwise.

### GetToVersionOk

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) GetToVersionOk() (*GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion, bool)`

GetToVersionOk returns a tuple with the ToVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToVersion

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) SetToVersion(v GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeToVersion)`

SetToVersion sets ToVersion field to given value.

### HasToVersion

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) HasToVersion() bool`

HasToVersion returns a boolean if a field has been set.

### GetStatus

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetId

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUpgradeBatchId

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) GetUpgradeBatchId() string`

GetUpgradeBatchId returns the UpgradeBatchId field if non-nil, zero value otherwise.

### GetUpgradeBatchIdOk

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) GetUpgradeBatchIdOk() (*string, bool)`

GetUpgradeBatchIdOk returns a tuple with the UpgradeBatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeBatchId

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) SetUpgradeBatchId(v string)`

SetUpgradeBatchId sets UpgradeBatchId field to given value.

### HasUpgradeBatchId

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) HasUpgradeBatchId() bool`

HasUpgradeBatchId returns a boolean if a field has been set.

### GetStaged

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) GetStaged() GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStaged`

GetStaged returns the Staged field if non-nil, zero value otherwise.

### GetStagedOk

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) GetStagedOk() (*GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStaged, bool)`

GetStagedOk returns a tuple with the Staged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaged

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) SetStaged(v GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgradeStaged)`

SetStaged sets Staged field to given value.

### HasStaged

`func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) HasStaged() bool`

HasStaged returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


