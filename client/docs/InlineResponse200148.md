# InlineResponse200148

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

### NewInlineResponse200148

`func NewInlineResponse200148() *InlineResponse200148`

NewInlineResponse200148 instantiates a new InlineResponse200148 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200148WithDefaults

`func NewInlineResponse200148WithDefaults() *InlineResponse200148`

NewInlineResponse200148WithDefaults instantiates a new InlineResponse200148 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200148) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200148) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200148) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200148) HasName() bool`

HasName returns a boolean if a field has been set.

### GetModel

`func (o *InlineResponse200148) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InlineResponse200148) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InlineResponse200148) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *InlineResponse200148) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetSerial

`func (o *InlineResponse200148) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200148) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200148) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200148) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetMac

`func (o *InlineResponse200148) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200148) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200148) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse200148) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetProductType

`func (o *InlineResponse200148) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *InlineResponse200148) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *InlineResponse200148) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *InlineResponse200148) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200148) GetNetwork() OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200148) GetNetworkOk() (*OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200148) SetNetwork(v OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200148) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetUsage

`func (o *InlineResponse200148) GetUsage() OrganizationsOrganizationIdSummaryTopDevicesByUsageUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *InlineResponse200148) GetUsageOk() (*OrganizationsOrganizationIdSummaryTopDevicesByUsageUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *InlineResponse200148) SetUsage(v OrganizationsOrganizationIdSummaryTopDevicesByUsageUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *InlineResponse200148) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetClients

`func (o *InlineResponse200148) GetClients() OrganizationsOrganizationIdSummaryTopDevicesByUsageClients`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *InlineResponse200148) GetClientsOk() (*OrganizationsOrganizationIdSummaryTopDevicesByUsageClients, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *InlineResponse200148) SetClients(v OrganizationsOrganizationIdSummaryTopDevicesByUsageClients)`

SetClients sets Clients field to given value.

### HasClients

`func (o *InlineResponse200148) HasClients() bool`

HasClients returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


