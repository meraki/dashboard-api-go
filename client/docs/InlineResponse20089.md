# InlineResponse20089

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTs** | Pointer to **time.Time** | The start time of the query range | [optional] 
**EndTs** | Pointer to **time.Time** | The end time of the query range | [optional] 
**AverageKbps** | Pointer to **int32** | Average data rate in kilobytes-per-second | [optional] 
**DownloadKbps** | Pointer to **int32** | Download rate in kilobytes-per-second | [optional] 
**UploadKbps** | Pointer to **int32** | Upload rate in kilobytes-per-second | [optional] 

## Methods

### NewInlineResponse20089

`func NewInlineResponse20089() *InlineResponse20089`

NewInlineResponse20089 instantiates a new InlineResponse20089 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20089WithDefaults

`func NewInlineResponse20089WithDefaults() *InlineResponse20089`

NewInlineResponse20089WithDefaults instantiates a new InlineResponse20089 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTs

`func (o *InlineResponse20089) GetStartTs() time.Time`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *InlineResponse20089) GetStartTsOk() (*time.Time, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *InlineResponse20089) SetStartTs(v time.Time)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *InlineResponse20089) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetEndTs

`func (o *InlineResponse20089) GetEndTs() time.Time`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *InlineResponse20089) GetEndTsOk() (*time.Time, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *InlineResponse20089) SetEndTs(v time.Time)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *InlineResponse20089) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetAverageKbps

`func (o *InlineResponse20089) GetAverageKbps() int32`

GetAverageKbps returns the AverageKbps field if non-nil, zero value otherwise.

### GetAverageKbpsOk

`func (o *InlineResponse20089) GetAverageKbpsOk() (*int32, bool)`

GetAverageKbpsOk returns a tuple with the AverageKbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageKbps

`func (o *InlineResponse20089) SetAverageKbps(v int32)`

SetAverageKbps sets AverageKbps field to given value.

### HasAverageKbps

`func (o *InlineResponse20089) HasAverageKbps() bool`

HasAverageKbps returns a boolean if a field has been set.

### GetDownloadKbps

`func (o *InlineResponse20089) GetDownloadKbps() int32`

GetDownloadKbps returns the DownloadKbps field if non-nil, zero value otherwise.

### GetDownloadKbpsOk

`func (o *InlineResponse20089) GetDownloadKbpsOk() (*int32, bool)`

GetDownloadKbpsOk returns a tuple with the DownloadKbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadKbps

`func (o *InlineResponse20089) SetDownloadKbps(v int32)`

SetDownloadKbps sets DownloadKbps field to given value.

### HasDownloadKbps

`func (o *InlineResponse20089) HasDownloadKbps() bool`

HasDownloadKbps returns a boolean if a field has been set.

### GetUploadKbps

`func (o *InlineResponse20089) GetUploadKbps() int32`

GetUploadKbps returns the UploadKbps field if non-nil, zero value otherwise.

### GetUploadKbpsOk

`func (o *InlineResponse20089) GetUploadKbpsOk() (*int32, bool)`

GetUploadKbpsOk returns a tuple with the UploadKbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadKbps

`func (o *InlineResponse20089) SetUploadKbps(v int32)`

SetUploadKbps sets UploadKbps field to given value.

### HasUploadKbps

`func (o *InlineResponse20089) HasUploadKbps() bool`

HasUploadKbps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


