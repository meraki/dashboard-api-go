# GetOrganizationSummaryTopDevicesByUsage200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the device | [optional] 
**Model** | Pointer to **string** | Model of the device | [optional] 
**Serial** | Pointer to **string** | Serial number of the device | [optional] 
**Mac** | Pointer to **string** | Mac address of the device | [optional] 
**ProductType** | Pointer to **string** | Product type of the device | [optional] 
**Network** | Pointer to [**GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerNetwork**](GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerNetwork.md) |  | [optional] 
**Usage** | Pointer to [**GetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage**](GetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage.md) |  | [optional] 
**Clients** | Pointer to [**GetOrganizationSummaryTopDevicesByUsage200ResponseInnerClients**](GetOrganizationSummaryTopDevicesByUsage200ResponseInnerClients.md) |  | [optional] 

## Methods

### NewGetOrganizationSummaryTopDevicesByUsage200ResponseInner

`func NewGetOrganizationSummaryTopDevicesByUsage200ResponseInner() *GetOrganizationSummaryTopDevicesByUsage200ResponseInner`

NewGetOrganizationSummaryTopDevicesByUsage200ResponseInner instantiates a new GetOrganizationSummaryTopDevicesByUsage200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationSummaryTopDevicesByUsage200ResponseInnerWithDefaults

`func NewGetOrganizationSummaryTopDevicesByUsage200ResponseInnerWithDefaults() *GetOrganizationSummaryTopDevicesByUsage200ResponseInner`

NewGetOrganizationSummaryTopDevicesByUsage200ResponseInnerWithDefaults instantiates a new GetOrganizationSummaryTopDevicesByUsage200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetModel

`func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetSerial

`func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetMac

`func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetProductType

`func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetNetwork

`func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) GetNetwork() GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) GetNetworkOk() (*GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) SetNetwork(v GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetUsage

`func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) GetUsage() GetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) GetUsageOk() (*GetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) SetUsage(v GetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetClients

`func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) GetClients() GetOrganizationSummaryTopDevicesByUsage200ResponseInnerClients`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) GetClientsOk() (*GetOrganizationSummaryTopDevicesByUsage200ResponseInnerClients, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) SetClients(v GetOrganizationSummaryTopDevicesByUsage200ResponseInnerClients)`

SetClients sets Clients field to given value.

### HasClients

`func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) HasClients() bool`

HasClients returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


