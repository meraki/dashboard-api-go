# RollbacksNetworkFirmwareUpgradesStagedEventsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stages** | [**[]UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner**](UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner.md) | All completed or in-progress stages in the network with their new start times. All pending stages will be canceled | 
**Reasons** | Pointer to [**[]CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner**](CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner.md) | The reason for rolling back the staged upgrade | [optional] 

## Methods

### NewRollbacksNetworkFirmwareUpgradesStagedEventsRequest

`func NewRollbacksNetworkFirmwareUpgradesStagedEventsRequest(stages []UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner, ) *RollbacksNetworkFirmwareUpgradesStagedEventsRequest`

NewRollbacksNetworkFirmwareUpgradesStagedEventsRequest instantiates a new RollbacksNetworkFirmwareUpgradesStagedEventsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRollbacksNetworkFirmwareUpgradesStagedEventsRequestWithDefaults

`func NewRollbacksNetworkFirmwareUpgradesStagedEventsRequestWithDefaults() *RollbacksNetworkFirmwareUpgradesStagedEventsRequest`

NewRollbacksNetworkFirmwareUpgradesStagedEventsRequestWithDefaults instantiates a new RollbacksNetworkFirmwareUpgradesStagedEventsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStages

`func (o *RollbacksNetworkFirmwareUpgradesStagedEventsRequest) GetStages() []UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *RollbacksNetworkFirmwareUpgradesStagedEventsRequest) GetStagesOk() (*[]UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *RollbacksNetworkFirmwareUpgradesStagedEventsRequest) SetStages(v []UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner)`

SetStages sets Stages field to given value.


### GetReasons

`func (o *RollbacksNetworkFirmwareUpgradesStagedEventsRequest) GetReasons() []CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *RollbacksNetworkFirmwareUpgradesStagedEventsRequest) GetReasonsOk() (*[]CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *RollbacksNetworkFirmwareUpgradesStagedEventsRequest) SetReasons(v []CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *RollbacksNetworkFirmwareUpgradesStagedEventsRequest) HasReasons() bool`

HasReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


