# GetNetworkFirmwareUpgrades200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpgradeWindow** | Pointer to [**GetNetworkFirmwareUpgrades200ResponseUpgradeWindow**](GetNetworkFirmwareUpgrades200ResponseUpgradeWindow.md) |  | [optional] 
**Timezone** | Pointer to **string** | The timezone for the network | [optional] 
**Products** | Pointer to [**GetNetworkFirmwareUpgrades200ResponseProducts**](GetNetworkFirmwareUpgrades200ResponseProducts.md) |  | [optional] 

## Methods

### NewGetNetworkFirmwareUpgrades200Response

`func NewGetNetworkFirmwareUpgrades200Response() *GetNetworkFirmwareUpgrades200Response`

NewGetNetworkFirmwareUpgrades200Response instantiates a new GetNetworkFirmwareUpgrades200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkFirmwareUpgrades200ResponseWithDefaults

`func NewGetNetworkFirmwareUpgrades200ResponseWithDefaults() *GetNetworkFirmwareUpgrades200Response`

NewGetNetworkFirmwareUpgrades200ResponseWithDefaults instantiates a new GetNetworkFirmwareUpgrades200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpgradeWindow

`func (o *GetNetworkFirmwareUpgrades200Response) GetUpgradeWindow() GetNetworkFirmwareUpgrades200ResponseUpgradeWindow`

GetUpgradeWindow returns the UpgradeWindow field if non-nil, zero value otherwise.

### GetUpgradeWindowOk

`func (o *GetNetworkFirmwareUpgrades200Response) GetUpgradeWindowOk() (*GetNetworkFirmwareUpgrades200ResponseUpgradeWindow, bool)`

GetUpgradeWindowOk returns a tuple with the UpgradeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeWindow

`func (o *GetNetworkFirmwareUpgrades200Response) SetUpgradeWindow(v GetNetworkFirmwareUpgrades200ResponseUpgradeWindow)`

SetUpgradeWindow sets UpgradeWindow field to given value.

### HasUpgradeWindow

`func (o *GetNetworkFirmwareUpgrades200Response) HasUpgradeWindow() bool`

HasUpgradeWindow returns a boolean if a field has been set.

### GetTimezone

`func (o *GetNetworkFirmwareUpgrades200Response) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *GetNetworkFirmwareUpgrades200Response) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *GetNetworkFirmwareUpgrades200Response) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *GetNetworkFirmwareUpgrades200Response) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetProducts

`func (o *GetNetworkFirmwareUpgrades200Response) GetProducts() GetNetworkFirmwareUpgrades200ResponseProducts`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *GetNetworkFirmwareUpgrades200Response) GetProductsOk() (*GetNetworkFirmwareUpgrades200ResponseProducts, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *GetNetworkFirmwareUpgrades200Response) SetProducts(v GetNetworkFirmwareUpgrades200ResponseProducts)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *GetNetworkFirmwareUpgrades200Response) HasProducts() bool`

HasProducts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


