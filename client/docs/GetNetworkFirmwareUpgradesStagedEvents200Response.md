# GetNetworkFirmwareUpgradesStagedEvents200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | Pointer to [**GetNetworkFirmwareUpgradesStagedEvents200ResponseProducts**](GetNetworkFirmwareUpgradesStagedEvents200ResponseProducts.md) |  | [optional] 
**Stages** | Pointer to [**[]GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner**](GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner.md) | The ordered stages in the network | [optional] 
**Reasons** | Pointer to [**[]CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner**](CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner.md) | Reasons for the rollback | [optional] 

## Methods

### NewGetNetworkFirmwareUpgradesStagedEvents200Response

`func NewGetNetworkFirmwareUpgradesStagedEvents200Response() *GetNetworkFirmwareUpgradesStagedEvents200Response`

NewGetNetworkFirmwareUpgradesStagedEvents200Response instantiates a new GetNetworkFirmwareUpgradesStagedEvents200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkFirmwareUpgradesStagedEvents200ResponseWithDefaults

`func NewGetNetworkFirmwareUpgradesStagedEvents200ResponseWithDefaults() *GetNetworkFirmwareUpgradesStagedEvents200Response`

NewGetNetworkFirmwareUpgradesStagedEvents200ResponseWithDefaults instantiates a new GetNetworkFirmwareUpgradesStagedEvents200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *GetNetworkFirmwareUpgradesStagedEvents200Response) GetProducts() GetNetworkFirmwareUpgradesStagedEvents200ResponseProducts`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *GetNetworkFirmwareUpgradesStagedEvents200Response) GetProductsOk() (*GetNetworkFirmwareUpgradesStagedEvents200ResponseProducts, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *GetNetworkFirmwareUpgradesStagedEvents200Response) SetProducts(v GetNetworkFirmwareUpgradesStagedEvents200ResponseProducts)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *GetNetworkFirmwareUpgradesStagedEvents200Response) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetStages

`func (o *GetNetworkFirmwareUpgradesStagedEvents200Response) GetStages() []GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *GetNetworkFirmwareUpgradesStagedEvents200Response) GetStagesOk() (*[]GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *GetNetworkFirmwareUpgradesStagedEvents200Response) SetStages(v []GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner)`

SetStages sets Stages field to given value.

### HasStages

`func (o *GetNetworkFirmwareUpgradesStagedEvents200Response) HasStages() bool`

HasStages returns a boolean if a field has been set.

### GetReasons

`func (o *GetNetworkFirmwareUpgradesStagedEvents200Response) GetReasons() []CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *GetNetworkFirmwareUpgradesStagedEvents200Response) GetReasonsOk() (*[]CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *GetNetworkFirmwareUpgradesStagedEvents200Response) SetReasons(v []CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *GetNetworkFirmwareUpgradesStagedEvents200Response) HasReasons() bool`

HasReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


