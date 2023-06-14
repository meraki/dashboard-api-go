# CreateNetworkFirmwareUpgradesRollback200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to **string** | Product type to rollback (if the network is a combined network) | [optional] 
**Status** | Pointer to **string** | Status of the rollback | [optional] 
**UpgradeBatchId** | Pointer to **string** | Batch ID of the firmware rollback | [optional] 
**Time** | Pointer to **time.Time** | Scheduled time for the rollback | [optional] 
**ToVersion** | Pointer to [**CreateNetworkFirmwareUpgradesRollback200ResponseToVersion**](CreateNetworkFirmwareUpgradesRollback200ResponseToVersion.md) |  | [optional] 
**Reasons** | Pointer to [**[]CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner**](CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner.md) | Reasons for the rollback | [optional] 

## Methods

### NewCreateNetworkFirmwareUpgradesRollback200Response

`func NewCreateNetworkFirmwareUpgradesRollback200Response() *CreateNetworkFirmwareUpgradesRollback200Response`

NewCreateNetworkFirmwareUpgradesRollback200Response instantiates a new CreateNetworkFirmwareUpgradesRollback200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkFirmwareUpgradesRollback200ResponseWithDefaults

`func NewCreateNetworkFirmwareUpgradesRollback200ResponseWithDefaults() *CreateNetworkFirmwareUpgradesRollback200Response`

NewCreateNetworkFirmwareUpgradesRollback200ResponseWithDefaults instantiates a new CreateNetworkFirmwareUpgradesRollback200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *CreateNetworkFirmwareUpgradesRollback200Response) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *CreateNetworkFirmwareUpgradesRollback200Response) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *CreateNetworkFirmwareUpgradesRollback200Response) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *CreateNetworkFirmwareUpgradesRollback200Response) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetStatus

`func (o *CreateNetworkFirmwareUpgradesRollback200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateNetworkFirmwareUpgradesRollback200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateNetworkFirmwareUpgradesRollback200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateNetworkFirmwareUpgradesRollback200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpgradeBatchId

`func (o *CreateNetworkFirmwareUpgradesRollback200Response) GetUpgradeBatchId() string`

GetUpgradeBatchId returns the UpgradeBatchId field if non-nil, zero value otherwise.

### GetUpgradeBatchIdOk

`func (o *CreateNetworkFirmwareUpgradesRollback200Response) GetUpgradeBatchIdOk() (*string, bool)`

GetUpgradeBatchIdOk returns a tuple with the UpgradeBatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeBatchId

`func (o *CreateNetworkFirmwareUpgradesRollback200Response) SetUpgradeBatchId(v string)`

SetUpgradeBatchId sets UpgradeBatchId field to given value.

### HasUpgradeBatchId

`func (o *CreateNetworkFirmwareUpgradesRollback200Response) HasUpgradeBatchId() bool`

HasUpgradeBatchId returns a boolean if a field has been set.

### GetTime

`func (o *CreateNetworkFirmwareUpgradesRollback200Response) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *CreateNetworkFirmwareUpgradesRollback200Response) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *CreateNetworkFirmwareUpgradesRollback200Response) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *CreateNetworkFirmwareUpgradesRollback200Response) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetToVersion

`func (o *CreateNetworkFirmwareUpgradesRollback200Response) GetToVersion() CreateNetworkFirmwareUpgradesRollback200ResponseToVersion`

GetToVersion returns the ToVersion field if non-nil, zero value otherwise.

### GetToVersionOk

`func (o *CreateNetworkFirmwareUpgradesRollback200Response) GetToVersionOk() (*CreateNetworkFirmwareUpgradesRollback200ResponseToVersion, bool)`

GetToVersionOk returns a tuple with the ToVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToVersion

`func (o *CreateNetworkFirmwareUpgradesRollback200Response) SetToVersion(v CreateNetworkFirmwareUpgradesRollback200ResponseToVersion)`

SetToVersion sets ToVersion field to given value.

### HasToVersion

`func (o *CreateNetworkFirmwareUpgradesRollback200Response) HasToVersion() bool`

HasToVersion returns a boolean if a field has been set.

### GetReasons

`func (o *CreateNetworkFirmwareUpgradesRollback200Response) GetReasons() []CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *CreateNetworkFirmwareUpgradesRollback200Response) GetReasonsOk() (*[]CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *CreateNetworkFirmwareUpgradesRollback200Response) SetReasons(v []CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *CreateNetworkFirmwareUpgradesRollback200Response) HasReasons() bool`

HasReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


