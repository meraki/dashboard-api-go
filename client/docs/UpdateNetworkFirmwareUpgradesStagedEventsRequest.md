# UpdateNetworkFirmwareUpgradesStagedEventsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stages** | [**[]UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner**](UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner.md) | All firmware upgrade stages in the network with their start time. | 

## Methods

### NewUpdateNetworkFirmwareUpgradesStagedEventsRequest

`func NewUpdateNetworkFirmwareUpgradesStagedEventsRequest(stages []UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner, ) *UpdateNetworkFirmwareUpgradesStagedEventsRequest`

NewUpdateNetworkFirmwareUpgradesStagedEventsRequest instantiates a new UpdateNetworkFirmwareUpgradesStagedEventsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkFirmwareUpgradesStagedEventsRequestWithDefaults

`func NewUpdateNetworkFirmwareUpgradesStagedEventsRequestWithDefaults() *UpdateNetworkFirmwareUpgradesStagedEventsRequest`

NewUpdateNetworkFirmwareUpgradesStagedEventsRequestWithDefaults instantiates a new UpdateNetworkFirmwareUpgradesStagedEventsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStages

`func (o *UpdateNetworkFirmwareUpgradesStagedEventsRequest) GetStages() []UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *UpdateNetworkFirmwareUpgradesStagedEventsRequest) GetStagesOk() (*[]UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *UpdateNetworkFirmwareUpgradesStagedEventsRequest) SetStages(v []UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner)`

SetStages sets Stages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


