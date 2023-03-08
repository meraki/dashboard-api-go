# OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to **string** | Start time of the upgrade | [optional] 
**FromVersion** | Pointer to [**OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion**](OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion.md) |  | [optional] 
**ToVersion** | Pointer to [**OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeToVersion**](OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeToVersion.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the upgrade | [optional] 
**Id** | Pointer to **string** | ID of the upgrade | [optional] 
**UpgradeBatchId** | Pointer to **string** | ID of the upgrade batch | [optional] 
**Staged** | Pointer to [**OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStaged**](OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStaged.md) |  | [optional] 

## Methods

### NewOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade

`func NewOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade() *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade`

NewOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade instantiates a new OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeWithDefaults

`func NewOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeWithDefaults() *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade`

NewOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeWithDefaults instantiates a new OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetFromVersion

`func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) GetFromVersion() OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion`

GetFromVersion returns the FromVersion field if non-nil, zero value otherwise.

### GetFromVersionOk

`func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) GetFromVersionOk() (*OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion, bool)`

GetFromVersionOk returns a tuple with the FromVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromVersion

`func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) SetFromVersion(v OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion)`

SetFromVersion sets FromVersion field to given value.

### HasFromVersion

`func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) HasFromVersion() bool`

HasFromVersion returns a boolean if a field has been set.

### GetToVersion

`func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) GetToVersion() OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeToVersion`

GetToVersion returns the ToVersion field if non-nil, zero value otherwise.

### GetToVersionOk

`func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) GetToVersionOk() (*OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeToVersion, bool)`

GetToVersionOk returns a tuple with the ToVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToVersion

`func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) SetToVersion(v OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeToVersion)`

SetToVersion sets ToVersion field to given value.

### HasToVersion

`func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) HasToVersion() bool`

HasToVersion returns a boolean if a field has been set.

### GetStatus

`func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetId

`func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUpgradeBatchId

`func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) GetUpgradeBatchId() string`

GetUpgradeBatchId returns the UpgradeBatchId field if non-nil, zero value otherwise.

### GetUpgradeBatchIdOk

`func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) GetUpgradeBatchIdOk() (*string, bool)`

GetUpgradeBatchIdOk returns a tuple with the UpgradeBatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeBatchId

`func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) SetUpgradeBatchId(v string)`

SetUpgradeBatchId sets UpgradeBatchId field to given value.

### HasUpgradeBatchId

`func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) HasUpgradeBatchId() bool`

HasUpgradeBatchId returns a boolean if a field has been set.

### GetStaged

`func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) GetStaged() OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStaged`

GetStaged returns the Staged field if non-nil, zero value otherwise.

### GetStagedOk

`func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) GetStagedOk() (*OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStaged, bool)`

GetStagedOk returns a tuple with the Staged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaged

`func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) SetStaged(v OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStaged)`

SetStaged sets Staged field to given value.

### HasStaged

`func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) HasStaged() bool`

HasStaged returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


