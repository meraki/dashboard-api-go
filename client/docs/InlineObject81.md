# InlineObject81

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stages** | [**[]NetworksNetworkIdFirmwareUpgradesStagedEventsStages**](NetworksNetworkIdFirmwareUpgradesStagedEventsStages.md) | All completed or in-progress stages in the network with their new start times. All pending stages will be canceled | 
**Reasons** | Pointer to [**[]NetworksNetworkIdFirmwareUpgradesRollbacksReasons**](NetworksNetworkIdFirmwareUpgradesRollbacksReasons.md) | The reason for rolling back the staged upgrade | [optional] 

## Methods

### NewInlineObject81

`func NewInlineObject81(stages []NetworksNetworkIdFirmwareUpgradesStagedEventsStages, ) *InlineObject81`

NewInlineObject81 instantiates a new InlineObject81 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject81WithDefaults

`func NewInlineObject81WithDefaults() *InlineObject81`

NewInlineObject81WithDefaults instantiates a new InlineObject81 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStages

`func (o *InlineObject81) GetStages() []NetworksNetworkIdFirmwareUpgradesStagedEventsStages`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *InlineObject81) GetStagesOk() (*[]NetworksNetworkIdFirmwareUpgradesStagedEventsStages, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *InlineObject81) SetStages(v []NetworksNetworkIdFirmwareUpgradesStagedEventsStages)`

SetStages sets Stages field to given value.


### GetReasons

`func (o *InlineObject81) GetReasons() []NetworksNetworkIdFirmwareUpgradesRollbacksReasons`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *InlineObject81) GetReasonsOk() (*[]NetworksNetworkIdFirmwareUpgradesRollbacksReasons, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *InlineObject81) SetReasons(v []NetworksNetworkIdFirmwareUpgradesRollbacksReasons)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *InlineObject81) HasReasons() bool`

HasReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


