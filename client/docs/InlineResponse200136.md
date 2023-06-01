# InlineResponse200136

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the device | [optional] 
**Model** | Pointer to **string** | Model of the device | [optional] 
**Serial** | Pointer to **string** | Serial number of the device | [optional] 
**Mac** | Pointer to **string** | Mac address of the device | [optional] 
**ProductType** | Pointer to **string** | Product type of the device | [optional] 
**Network** | Pointer to [**OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork**](OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork.md) |  | [optional] 
**Usage** | Pointer to [**OrganizationsOrganizationIdSummaryTopDevicesByUsageUsage**](OrganizationsOrganizationIdSummaryTopDevicesByUsageUsage.md) |  | [optional] 
**Clients** | Pointer to [**OrganizationsOrganizationIdSummaryTopDevicesByUsageClients**](OrganizationsOrganizationIdSummaryTopDevicesByUsageClients.md) |  | [optional] 

## Methods

### NewInlineResponse200136

`func NewInlineResponse200136() *InlineResponse200136`

NewInlineResponse200136 instantiates a new InlineResponse200136 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200136WithDefaults

`func NewInlineResponse200136WithDefaults() *InlineResponse200136`

NewInlineResponse200136WithDefaults instantiates a new InlineResponse200136 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200136) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200136) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200136) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200136) HasName() bool`

HasName returns a boolean if a field has been set.

### GetModel

`func (o *InlineResponse200136) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InlineResponse200136) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InlineResponse200136) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *InlineResponse200136) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetSerial

`func (o *InlineResponse200136) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200136) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200136) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200136) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetMac

`func (o *InlineResponse200136) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200136) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200136) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse200136) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetProductType

`func (o *InlineResponse200136) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *InlineResponse200136) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *InlineResponse200136) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *InlineResponse200136) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200136) GetNetwork() OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200136) GetNetworkOk() (*OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200136) SetNetwork(v OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200136) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetUsage

`func (o *InlineResponse200136) GetUsage() OrganizationsOrganizationIdSummaryTopDevicesByUsageUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *InlineResponse200136) GetUsageOk() (*OrganizationsOrganizationIdSummaryTopDevicesByUsageUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *InlineResponse200136) SetUsage(v OrganizationsOrganizationIdSummaryTopDevicesByUsageUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *InlineResponse200136) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetClients

`func (o *InlineResponse200136) GetClients() OrganizationsOrganizationIdSummaryTopDevicesByUsageClients`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *InlineResponse200136) GetClientsOk() (*OrganizationsOrganizationIdSummaryTopDevicesByUsageClients, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *InlineResponse200136) SetClients(v OrganizationsOrganizationIdSummaryTopDevicesByUsageClients)`

SetClients sets Clients field to given value.

### HasClients

`func (o *InlineResponse200136) HasClients() bool`

HasClients returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


