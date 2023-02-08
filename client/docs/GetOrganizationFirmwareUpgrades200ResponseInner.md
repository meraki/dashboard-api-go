# GetOrganizationFirmwareUpgrades200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpgradeId** | Pointer to **string** | The upgrade | [optional] 
**UpgradeBatchId** | Pointer to **string** | The upgrade batch | [optional] 
**Network** | Pointer to [**GetOrganizationFirmwareUpgrades200ResponseInnerNetwork**](GetOrganizationFirmwareUpgrades200ResponseInnerNetwork.md) |  | [optional] 
**Status** | Pointer to **string** | Status of upgrade event: [Cancelled, Completed] | [optional] 
**Time** | Pointer to **time.Time** | Scheduled start time | [optional] 
**CompletedAt** | Pointer to **string** | Timestamp when upgrade completed. Null if status pending. | [optional] 
**ProductType** | Pointer to **string** | product upgraded [wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor] | [optional] 
**ToVersion** | Pointer to [**GetOrganizationFirmwareUpgrades200ResponseInnerToVersion**](GetOrganizationFirmwareUpgrades200ResponseInnerToVersion.md) |  | [optional] 
**FromVersion** | Pointer to [**GetOrganizationFirmwareUpgrades200ResponseInnerFromVersion**](GetOrganizationFirmwareUpgrades200ResponseInnerFromVersion.md) |  | [optional] 

## Methods

### NewGetOrganizationFirmwareUpgrades200ResponseInner

`func NewGetOrganizationFirmwareUpgrades200ResponseInner() *GetOrganizationFirmwareUpgrades200ResponseInner`

NewGetOrganizationFirmwareUpgrades200ResponseInner instantiates a new GetOrganizationFirmwareUpgrades200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationFirmwareUpgrades200ResponseInnerWithDefaults

`func NewGetOrganizationFirmwareUpgrades200ResponseInnerWithDefaults() *GetOrganizationFirmwareUpgrades200ResponseInner`

NewGetOrganizationFirmwareUpgrades200ResponseInnerWithDefaults instantiates a new GetOrganizationFirmwareUpgrades200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpgradeId

`func (o *GetOrganizationFirmwareUpgrades200ResponseInner) GetUpgradeId() string`

GetUpgradeId returns the UpgradeId field if non-nil, zero value otherwise.

### GetUpgradeIdOk

`func (o *GetOrganizationFirmwareUpgrades200ResponseInner) GetUpgradeIdOk() (*string, bool)`

GetUpgradeIdOk returns a tuple with the UpgradeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeId

`func (o *GetOrganizationFirmwareUpgrades200ResponseInner) SetUpgradeId(v string)`

SetUpgradeId sets UpgradeId field to given value.

### HasUpgradeId

`func (o *GetOrganizationFirmwareUpgrades200ResponseInner) HasUpgradeId() bool`

HasUpgradeId returns a boolean if a field has been set.

### GetUpgradeBatchId

`func (o *GetOrganizationFirmwareUpgrades200ResponseInner) GetUpgradeBatchId() string`

GetUpgradeBatchId returns the UpgradeBatchId field if non-nil, zero value otherwise.

### GetUpgradeBatchIdOk

`func (o *GetOrganizationFirmwareUpgrades200ResponseInner) GetUpgradeBatchIdOk() (*string, bool)`

GetUpgradeBatchIdOk returns a tuple with the UpgradeBatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeBatchId

`func (o *GetOrganizationFirmwareUpgrades200ResponseInner) SetUpgradeBatchId(v string)`

SetUpgradeBatchId sets UpgradeBatchId field to given value.

### HasUpgradeBatchId

`func (o *GetOrganizationFirmwareUpgrades200ResponseInner) HasUpgradeBatchId() bool`

HasUpgradeBatchId returns a boolean if a field has been set.

### GetNetwork

`func (o *GetOrganizationFirmwareUpgrades200ResponseInner) GetNetwork() GetOrganizationFirmwareUpgrades200ResponseInnerNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetOrganizationFirmwareUpgrades200ResponseInner) GetNetworkOk() (*GetOrganizationFirmwareUpgrades200ResponseInnerNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetOrganizationFirmwareUpgrades200ResponseInner) SetNetwork(v GetOrganizationFirmwareUpgrades200ResponseInnerNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetOrganizationFirmwareUpgrades200ResponseInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetStatus

`func (o *GetOrganizationFirmwareUpgrades200ResponseInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetOrganizationFirmwareUpgrades200ResponseInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetOrganizationFirmwareUpgrades200ResponseInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetOrganizationFirmwareUpgrades200ResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTime

`func (o *GetOrganizationFirmwareUpgrades200ResponseInner) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *GetOrganizationFirmwareUpgrades200ResponseInner) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *GetOrganizationFirmwareUpgrades200ResponseInner) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *GetOrganizationFirmwareUpgrades200ResponseInner) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetCompletedAt

`func (o *GetOrganizationFirmwareUpgrades200ResponseInner) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *GetOrganizationFirmwareUpgrades200ResponseInner) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *GetOrganizationFirmwareUpgrades200ResponseInner) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *GetOrganizationFirmwareUpgrades200ResponseInner) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetProductType

`func (o *GetOrganizationFirmwareUpgrades200ResponseInner) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *GetOrganizationFirmwareUpgrades200ResponseInner) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *GetOrganizationFirmwareUpgrades200ResponseInner) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *GetOrganizationFirmwareUpgrades200ResponseInner) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetToVersion

`func (o *GetOrganizationFirmwareUpgrades200ResponseInner) GetToVersion() GetOrganizationFirmwareUpgrades200ResponseInnerToVersion`

GetToVersion returns the ToVersion field if non-nil, zero value otherwise.

### GetToVersionOk

`func (o *GetOrganizationFirmwareUpgrades200ResponseInner) GetToVersionOk() (*GetOrganizationFirmwareUpgrades200ResponseInnerToVersion, bool)`

GetToVersionOk returns a tuple with the ToVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToVersion

`func (o *GetOrganizationFirmwareUpgrades200ResponseInner) SetToVersion(v GetOrganizationFirmwareUpgrades200ResponseInnerToVersion)`

SetToVersion sets ToVersion field to given value.

### HasToVersion

`func (o *GetOrganizationFirmwareUpgrades200ResponseInner) HasToVersion() bool`

HasToVersion returns a boolean if a field has been set.

### GetFromVersion

`func (o *GetOrganizationFirmwareUpgrades200ResponseInner) GetFromVersion() GetOrganizationFirmwareUpgrades200ResponseInnerFromVersion`

GetFromVersion returns the FromVersion field if non-nil, zero value otherwise.

### GetFromVersionOk

`func (o *GetOrganizationFirmwareUpgrades200ResponseInner) GetFromVersionOk() (*GetOrganizationFirmwareUpgrades200ResponseInnerFromVersion, bool)`

GetFromVersionOk returns a tuple with the FromVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromVersion

`func (o *GetOrganizationFirmwareUpgrades200ResponseInner) SetFromVersion(v GetOrganizationFirmwareUpgrades200ResponseInnerFromVersion)`

SetFromVersion sets FromVersion field to given value.

### HasFromVersion

`func (o *GetOrganizationFirmwareUpgrades200ResponseInner) HasFromVersion() bool`

HasFromVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


