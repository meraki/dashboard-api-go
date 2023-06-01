# InlineResponse20077

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTs** | Pointer to **time.Time** | The start time of the query range | [optional] 
**EndTs** | Pointer to **time.Time** | The end time of the query range | [optional] 
**UtilizationTotal** | Pointer to **float32** | Total channel utilization | [optional] 
**Utilization80211** | Pointer to **float32** | Average wifi utilization | [optional] 
**UtilizationNon80211** | Pointer to **float32** | Average signal interference | [optional] 

## Methods

### NewInlineResponse20077

`func NewInlineResponse20077() *InlineResponse20077`

NewInlineResponse20077 instantiates a new InlineResponse20077 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20077WithDefaults

`func NewInlineResponse20077WithDefaults() *InlineResponse20077`

NewInlineResponse20077WithDefaults instantiates a new InlineResponse20077 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTs

`func (o *InlineResponse20077) GetStartTs() time.Time`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *InlineResponse20077) GetStartTsOk() (*time.Time, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *InlineResponse20077) SetStartTs(v time.Time)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *InlineResponse20077) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetEndTs

`func (o *InlineResponse20077) GetEndTs() time.Time`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *InlineResponse20077) GetEndTsOk() (*time.Time, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *InlineResponse20077) SetEndTs(v time.Time)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *InlineResponse20077) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetUtilizationTotal

`func (o *InlineResponse20077) GetUtilizationTotal() float32`

GetUtilizationTotal returns the UtilizationTotal field if non-nil, zero value otherwise.

### GetUtilizationTotalOk

`func (o *InlineResponse20077) GetUtilizationTotalOk() (*float32, bool)`

GetUtilizationTotalOk returns a tuple with the UtilizationTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizationTotal

`func (o *InlineResponse20077) SetUtilizationTotal(v float32)`

SetUtilizationTotal sets UtilizationTotal field to given value.

### HasUtilizationTotal

`func (o *InlineResponse20077) HasUtilizationTotal() bool`

HasUtilizationTotal returns a boolean if a field has been set.

### GetUtilization80211

`func (o *InlineResponse20077) GetUtilization80211() float32`

GetUtilization80211 returns the Utilization80211 field if non-nil, zero value otherwise.

### GetUtilization80211Ok

`func (o *InlineResponse20077) GetUtilization80211Ok() (*float32, bool)`

GetUtilization80211Ok returns a tuple with the Utilization80211 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilization80211

`func (o *InlineResponse20077) SetUtilization80211(v float32)`

SetUtilization80211 sets Utilization80211 field to given value.

### HasUtilization80211

`func (o *InlineResponse20077) HasUtilization80211() bool`

HasUtilization80211 returns a boolean if a field has been set.

### GetUtilizationNon80211

`func (o *InlineResponse20077) GetUtilizationNon80211() float32`

GetUtilizationNon80211 returns the UtilizationNon80211 field if non-nil, zero value otherwise.

### GetUtilizationNon80211Ok

`func (o *InlineResponse20077) GetUtilizationNon80211Ok() (*float32, bool)`

GetUtilizationNon80211Ok returns a tuple with the UtilizationNon80211 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizationNon80211

`func (o *InlineResponse20077) SetUtilizationNon80211(v float32)`

SetUtilizationNon80211 sets UtilizationNon80211 field to given value.

### HasUtilizationNon80211

`func (o *InlineResponse20077) HasUtilizationNon80211() bool`

HasUtilizationNon80211 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


