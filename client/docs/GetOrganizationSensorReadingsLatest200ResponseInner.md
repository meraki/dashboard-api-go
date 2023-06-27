# GetOrganizationSensorReadingsLatest200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Serial number of the sensor that took the readings. | [optional] 
**Network** | Pointer to [**GetOrganizationSensorReadingsHistory200ResponseInnerNetwork**](GetOrganizationSensorReadingsHistory200ResponseInnerNetwork.md) |  | [optional] 
**Readings** | Pointer to [**[]GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner**](GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner.md) | Array of latest readings from the sensor. Each object represents a single reading for a single metric. | [optional] 

## Methods

### NewGetOrganizationSensorReadingsLatest200ResponseInner

`func NewGetOrganizationSensorReadingsLatest200ResponseInner() *GetOrganizationSensorReadingsLatest200ResponseInner`

NewGetOrganizationSensorReadingsLatest200ResponseInner instantiates a new GetOrganizationSensorReadingsLatest200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationSensorReadingsLatest200ResponseInnerWithDefaults

`func NewGetOrganizationSensorReadingsLatest200ResponseInnerWithDefaults() *GetOrganizationSensorReadingsLatest200ResponseInner`

NewGetOrganizationSensorReadingsLatest200ResponseInnerWithDefaults instantiates a new GetOrganizationSensorReadingsLatest200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *GetOrganizationSensorReadingsLatest200ResponseInner) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *GetOrganizationSensorReadingsLatest200ResponseInner) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *GetOrganizationSensorReadingsLatest200ResponseInner) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *GetOrganizationSensorReadingsLatest200ResponseInner) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetNetwork

`func (o *GetOrganizationSensorReadingsLatest200ResponseInner) GetNetwork() GetOrganizationSensorReadingsHistory200ResponseInnerNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetOrganizationSensorReadingsLatest200ResponseInner) GetNetworkOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetOrganizationSensorReadingsLatest200ResponseInner) SetNetwork(v GetOrganizationSensorReadingsHistory200ResponseInnerNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetOrganizationSensorReadingsLatest200ResponseInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetReadings

`func (o *GetOrganizationSensorReadingsLatest200ResponseInner) GetReadings() []GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner`

GetReadings returns the Readings field if non-nil, zero value otherwise.

### GetReadingsOk

`func (o *GetOrganizationSensorReadingsLatest200ResponseInner) GetReadingsOk() (*[]GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner, bool)`

GetReadingsOk returns a tuple with the Readings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadings

`func (o *GetOrganizationSensorReadingsLatest200ResponseInner) SetReadings(v []GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner)`

SetReadings sets Readings field to given value.

### HasReadings

`func (o *GetOrganizationSensorReadingsLatest200ResponseInner) HasReadings() bool`

HasReadings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


