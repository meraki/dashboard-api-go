# CreateNetworkFirmwareUpgradesRollbackRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to **string** | Product type to rollback (if the network is a combined network) | [optional] 
**Time** | Pointer to **time.Time** | Scheduled time for the rollback | [optional] 
**Reasons** | [**[]CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner**](CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner.md) | Reasons for the rollback | 
**ToVersion** | Pointer to [**CreateNetworkFirmwareUpgradesRollbackRequestToVersion**](CreateNetworkFirmwareUpgradesRollbackRequestToVersion.md) |  | [optional] 

## Methods

### NewCreateNetworkFirmwareUpgradesRollbackRequest

`func NewCreateNetworkFirmwareUpgradesRollbackRequest(reasons []CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner, ) *CreateNetworkFirmwareUpgradesRollbackRequest`

NewCreateNetworkFirmwareUpgradesRollbackRequest instantiates a new CreateNetworkFirmwareUpgradesRollbackRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkFirmwareUpgradesRollbackRequestWithDefaults

`func NewCreateNetworkFirmwareUpgradesRollbackRequestWithDefaults() *CreateNetworkFirmwareUpgradesRollbackRequest`

NewCreateNetworkFirmwareUpgradesRollbackRequestWithDefaults instantiates a new CreateNetworkFirmwareUpgradesRollbackRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *CreateNetworkFirmwareUpgradesRollbackRequest) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *CreateNetworkFirmwareUpgradesRollbackRequest) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *CreateNetworkFirmwareUpgradesRollbackRequest) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *CreateNetworkFirmwareUpgradesRollbackRequest) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetTime

`func (o *CreateNetworkFirmwareUpgradesRollbackRequest) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *CreateNetworkFirmwareUpgradesRollbackRequest) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *CreateNetworkFirmwareUpgradesRollbackRequest) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *CreateNetworkFirmwareUpgradesRollbackRequest) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetReasons

`func (o *CreateNetworkFirmwareUpgradesRollbackRequest) GetReasons() []CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *CreateNetworkFirmwareUpgradesRollbackRequest) GetReasonsOk() (*[]CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *CreateNetworkFirmwareUpgradesRollbackRequest) SetReasons(v []CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner)`

SetReasons sets Reasons field to given value.


### GetToVersion

`func (o *CreateNetworkFirmwareUpgradesRollbackRequest) GetToVersion() CreateNetworkFirmwareUpgradesRollbackRequestToVersion`

GetToVersion returns the ToVersion field if non-nil, zero value otherwise.

### GetToVersionOk

`func (o *CreateNetworkFirmwareUpgradesRollbackRequest) GetToVersionOk() (*CreateNetworkFirmwareUpgradesRollbackRequestToVersion, bool)`

GetToVersionOk returns a tuple with the ToVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToVersion

`func (o *CreateNetworkFirmwareUpgradesRollbackRequest) SetToVersion(v CreateNetworkFirmwareUpgradesRollbackRequestToVersion)`

SetToVersion sets ToVersion field to given value.

### HasToVersion

`func (o *CreateNetworkFirmwareUpgradesRollbackRequest) HasToVersion() bool`

HasToVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


