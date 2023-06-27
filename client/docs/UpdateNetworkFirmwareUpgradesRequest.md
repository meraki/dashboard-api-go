# UpdateNetworkFirmwareUpgradesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpgradeWindow** | Pointer to [**GetNetworkFirmwareUpgrades200ResponseUpgradeWindow**](GetNetworkFirmwareUpgrades200ResponseUpgradeWindow.md) |  | [optional] 
**Timezone** | Pointer to **string** | The timezone for the network | [optional] 
**Products** | Pointer to [**UpdateNetworkFirmwareUpgradesRequestProducts**](UpdateNetworkFirmwareUpgradesRequestProducts.md) |  | [optional] 

## Methods

### NewUpdateNetworkFirmwareUpgradesRequest

`func NewUpdateNetworkFirmwareUpgradesRequest() *UpdateNetworkFirmwareUpgradesRequest`

NewUpdateNetworkFirmwareUpgradesRequest instantiates a new UpdateNetworkFirmwareUpgradesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkFirmwareUpgradesRequestWithDefaults

`func NewUpdateNetworkFirmwareUpgradesRequestWithDefaults() *UpdateNetworkFirmwareUpgradesRequest`

NewUpdateNetworkFirmwareUpgradesRequestWithDefaults instantiates a new UpdateNetworkFirmwareUpgradesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpgradeWindow

`func (o *UpdateNetworkFirmwareUpgradesRequest) GetUpgradeWindow() GetNetworkFirmwareUpgrades200ResponseUpgradeWindow`

GetUpgradeWindow returns the UpgradeWindow field if non-nil, zero value otherwise.

### GetUpgradeWindowOk

`func (o *UpdateNetworkFirmwareUpgradesRequest) GetUpgradeWindowOk() (*GetNetworkFirmwareUpgrades200ResponseUpgradeWindow, bool)`

GetUpgradeWindowOk returns a tuple with the UpgradeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeWindow

`func (o *UpdateNetworkFirmwareUpgradesRequest) SetUpgradeWindow(v GetNetworkFirmwareUpgrades200ResponseUpgradeWindow)`

SetUpgradeWindow sets UpgradeWindow field to given value.

### HasUpgradeWindow

`func (o *UpdateNetworkFirmwareUpgradesRequest) HasUpgradeWindow() bool`

HasUpgradeWindow returns a boolean if a field has been set.

### GetTimezone

`func (o *UpdateNetworkFirmwareUpgradesRequest) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *UpdateNetworkFirmwareUpgradesRequest) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *UpdateNetworkFirmwareUpgradesRequest) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *UpdateNetworkFirmwareUpgradesRequest) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetProducts

`func (o *UpdateNetworkFirmwareUpgradesRequest) GetProducts() UpdateNetworkFirmwareUpgradesRequestProducts`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *UpdateNetworkFirmwareUpgradesRequest) GetProductsOk() (*UpdateNetworkFirmwareUpgradesRequestProducts, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *UpdateNetworkFirmwareUpgradesRequest) SetProducts(v UpdateNetworkFirmwareUpgradesRequestProducts)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *UpdateNetworkFirmwareUpgradesRequest) HasProducts() bool`

HasProducts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


