# InlineObject5

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **time.Time** | [optional] The snapshot will be taken from this time on the camera. The timestamp is expected to be in ISO 8601 format. If no timestamp is specified, we will assume current time. | [optional] 
**Fullframe** | Pointer to **bool** | [optional] If set to \&quot;true\&quot; the snapshot will be taken at full sensor resolution. This will error if used with timestamp. | [optional] 

## Methods

### NewInlineObject5

`func NewInlineObject5() *InlineObject5`

NewInlineObject5 instantiates a new InlineObject5 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject5WithDefaults

`func NewInlineObject5WithDefaults() *InlineObject5`

NewInlineObject5WithDefaults instantiates a new InlineObject5 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *InlineObject5) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *InlineObject5) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *InlineObject5) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *InlineObject5) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetFullframe

`func (o *InlineObject5) GetFullframe() bool`

GetFullframe returns the Fullframe field if non-nil, zero value otherwise.

### GetFullframeOk

`func (o *InlineObject5) GetFullframeOk() (*bool, bool)`

GetFullframeOk returns a tuple with the Fullframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullframe

`func (o *InlineObject5) SetFullframe(v bool)`

SetFullframe sets Fullframe field to given value.

### HasFullframe

`func (o *InlineObject5) HasFullframe() bool`

HasFullframe returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


