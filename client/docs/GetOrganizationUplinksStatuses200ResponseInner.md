# GetOrganizationUplinksStatuses200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | Pointer to **string** | Network identifier | [optional] 
**Serial** | Pointer to **string** | The uplink serial | [optional] 
**Model** | Pointer to **string** | The uplink model | [optional] 
**LastReportedAt** | Pointer to **time.Time** | Last reported time for the device | [optional] 
**Uplinks** | Pointer to [**[]GetOrganizationUplinksStatuses200ResponseInnerUplinksInner**](GetOrganizationUplinksStatuses200ResponseInnerUplinksInner.md) | Uplinks | [optional] 

## Methods

### NewGetOrganizationUplinksStatuses200ResponseInner

`func NewGetOrganizationUplinksStatuses200ResponseInner() *GetOrganizationUplinksStatuses200ResponseInner`

NewGetOrganizationUplinksStatuses200ResponseInner instantiates a new GetOrganizationUplinksStatuses200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationUplinksStatuses200ResponseInnerWithDefaults

`func NewGetOrganizationUplinksStatuses200ResponseInnerWithDefaults() *GetOrganizationUplinksStatuses200ResponseInner`

NewGetOrganizationUplinksStatuses200ResponseInnerWithDefaults instantiates a new GetOrganizationUplinksStatuses200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *GetOrganizationUplinksStatuses200ResponseInner) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *GetOrganizationUplinksStatuses200ResponseInner) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *GetOrganizationUplinksStatuses200ResponseInner) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *GetOrganizationUplinksStatuses200ResponseInner) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetSerial

`func (o *GetOrganizationUplinksStatuses200ResponseInner) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *GetOrganizationUplinksStatuses200ResponseInner) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *GetOrganizationUplinksStatuses200ResponseInner) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *GetOrganizationUplinksStatuses200ResponseInner) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetModel

`func (o *GetOrganizationUplinksStatuses200ResponseInner) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *GetOrganizationUplinksStatuses200ResponseInner) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *GetOrganizationUplinksStatuses200ResponseInner) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *GetOrganizationUplinksStatuses200ResponseInner) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetLastReportedAt

`func (o *GetOrganizationUplinksStatuses200ResponseInner) GetLastReportedAt() time.Time`

GetLastReportedAt returns the LastReportedAt field if non-nil, zero value otherwise.

### GetLastReportedAtOk

`func (o *GetOrganizationUplinksStatuses200ResponseInner) GetLastReportedAtOk() (*time.Time, bool)`

GetLastReportedAtOk returns a tuple with the LastReportedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReportedAt

`func (o *GetOrganizationUplinksStatuses200ResponseInner) SetLastReportedAt(v time.Time)`

SetLastReportedAt sets LastReportedAt field to given value.

### HasLastReportedAt

`func (o *GetOrganizationUplinksStatuses200ResponseInner) HasLastReportedAt() bool`

HasLastReportedAt returns a boolean if a field has been set.

### GetUplinks

`func (o *GetOrganizationUplinksStatuses200ResponseInner) GetUplinks() []GetOrganizationUplinksStatuses200ResponseInnerUplinksInner`

GetUplinks returns the Uplinks field if non-nil, zero value otherwise.

### GetUplinksOk

`func (o *GetOrganizationUplinksStatuses200ResponseInner) GetUplinksOk() (*[]GetOrganizationUplinksStatuses200ResponseInnerUplinksInner, bool)`

GetUplinksOk returns a tuple with the Uplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinks

`func (o *GetOrganizationUplinksStatuses200ResponseInner) SetUplinks(v []GetOrganizationUplinksStatuses200ResponseInnerUplinksInner)`

SetUplinks sets Uplinks field to given value.

### HasUplinks

`func (o *GetOrganizationUplinksStatuses200ResponseInner) HasUplinks() bool`

HasUplinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


