# GetOrganizationCellularGatewayUplinkStatuses200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | Pointer to **string** | Network Id | [optional] 
**Serial** | Pointer to **string** | Serial number of the device | [optional] 
**Model** | Pointer to **string** | Device model | [optional] 
**LastReportedAt** | Pointer to **time.Time** | Last reported time for the device | [optional] 
**Uplinks** | Pointer to [**[]GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner**](GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner.md) | Uplinks info | [optional] 

## Methods

### NewGetOrganizationCellularGatewayUplinkStatuses200ResponseInner

`func NewGetOrganizationCellularGatewayUplinkStatuses200ResponseInner() *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner`

NewGetOrganizationCellularGatewayUplinkStatuses200ResponseInner instantiates a new GetOrganizationCellularGatewayUplinkStatuses200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationCellularGatewayUplinkStatuses200ResponseInnerWithDefaults

`func NewGetOrganizationCellularGatewayUplinkStatuses200ResponseInnerWithDefaults() *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner`

NewGetOrganizationCellularGatewayUplinkStatuses200ResponseInnerWithDefaults instantiates a new GetOrganizationCellularGatewayUplinkStatuses200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetSerial

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetModel

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetLastReportedAt

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) GetLastReportedAt() time.Time`

GetLastReportedAt returns the LastReportedAt field if non-nil, zero value otherwise.

### GetLastReportedAtOk

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) GetLastReportedAtOk() (*time.Time, bool)`

GetLastReportedAtOk returns a tuple with the LastReportedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReportedAt

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) SetLastReportedAt(v time.Time)`

SetLastReportedAt sets LastReportedAt field to given value.

### HasLastReportedAt

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) HasLastReportedAt() bool`

HasLastReportedAt returns a boolean if a field has been set.

### GetUplinks

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) GetUplinks() []GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner`

GetUplinks returns the Uplinks field if non-nil, zero value otherwise.

### GetUplinksOk

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) GetUplinksOk() (*[]GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner, bool)`

GetUplinksOk returns a tuple with the Uplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinks

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) SetUplinks(v []GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner)`

SetUplinks sets Uplinks field to given value.

### HasUplinks

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInner) HasUplinks() bool`

HasUplinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


