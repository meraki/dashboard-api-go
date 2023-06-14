# CreateNetworkFirmwareUpgradesStagedEventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | Pointer to [**CreateNetworkFirmwareUpgradesStagedEventRequestProducts**](CreateNetworkFirmwareUpgradesStagedEventRequestProducts.md) |  | [optional] 
**Stages** | [**[]UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner**](UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner.md) | All firmware upgrade stages in the network with their start time. | 

## Methods

### NewCreateNetworkFirmwareUpgradesStagedEventRequest

`func NewCreateNetworkFirmwareUpgradesStagedEventRequest(stages []UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner, ) *CreateNetworkFirmwareUpgradesStagedEventRequest`

NewCreateNetworkFirmwareUpgradesStagedEventRequest instantiates a new CreateNetworkFirmwareUpgradesStagedEventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkFirmwareUpgradesStagedEventRequestWithDefaults

`func NewCreateNetworkFirmwareUpgradesStagedEventRequestWithDefaults() *CreateNetworkFirmwareUpgradesStagedEventRequest`

NewCreateNetworkFirmwareUpgradesStagedEventRequestWithDefaults instantiates a new CreateNetworkFirmwareUpgradesStagedEventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *CreateNetworkFirmwareUpgradesStagedEventRequest) GetProducts() CreateNetworkFirmwareUpgradesStagedEventRequestProducts`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *CreateNetworkFirmwareUpgradesStagedEventRequest) GetProductsOk() (*CreateNetworkFirmwareUpgradesStagedEventRequestProducts, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *CreateNetworkFirmwareUpgradesStagedEventRequest) SetProducts(v CreateNetworkFirmwareUpgradesStagedEventRequestProducts)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *CreateNetworkFirmwareUpgradesStagedEventRequest) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetStages

`func (o *CreateNetworkFirmwareUpgradesStagedEventRequest) GetStages() []UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *CreateNetworkFirmwareUpgradesStagedEventRequest) GetStagesOk() (*[]UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *CreateNetworkFirmwareUpgradesStagedEventRequest) SetStages(v []UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner)`

SetStages sets Stages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


