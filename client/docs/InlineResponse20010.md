# InlineResponse20010

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | The serial number for the device | [optional] 
**ConnectionStats** | Pointer to [**InlineResponse20010ConnectionStats**](InlineResponse20010ConnectionStats.md) |  | [optional] 

## Methods

### NewInlineResponse20010

`func NewInlineResponse20010() *InlineResponse20010`

NewInlineResponse20010 instantiates a new InlineResponse20010 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20010WithDefaults

`func NewInlineResponse20010WithDefaults() *InlineResponse20010`

NewInlineResponse20010WithDefaults instantiates a new InlineResponse20010 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse20010) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse20010) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse20010) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse20010) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetConnectionStats

`func (o *InlineResponse20010) GetConnectionStats() InlineResponse20010ConnectionStats`

GetConnectionStats returns the ConnectionStats field if non-nil, zero value otherwise.

### GetConnectionStatsOk

`func (o *InlineResponse20010) GetConnectionStatsOk() (*InlineResponse20010ConnectionStats, bool)`

GetConnectionStatsOk returns a tuple with the ConnectionStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStats

`func (o *InlineResponse20010) SetConnectionStats(v InlineResponse20010ConnectionStats)`

SetConnectionStats sets ConnectionStats field to given value.

### HasConnectionStats

`func (o *InlineResponse20010) HasConnectionStats() bool`

HasConnectionStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


