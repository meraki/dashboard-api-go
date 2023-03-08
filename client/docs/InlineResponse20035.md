# InlineResponse20035

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTs** | Pointer to **time.Time** | Start of the timespan over which sensor alerts are counted | [optional] 
**EndTs** | Pointer to **time.Time** | End of the timespan over which sensor alerts are counted | [optional] 
**Counts** | Pointer to [**NetworksNetworkIdSensorAlertsOverviewByMetricCounts**](NetworksNetworkIdSensorAlertsOverviewByMetricCounts.md) |  | [optional] 

## Methods

### NewInlineResponse20035

`func NewInlineResponse20035() *InlineResponse20035`

NewInlineResponse20035 instantiates a new InlineResponse20035 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20035WithDefaults

`func NewInlineResponse20035WithDefaults() *InlineResponse20035`

NewInlineResponse20035WithDefaults instantiates a new InlineResponse20035 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTs

`func (o *InlineResponse20035) GetStartTs() time.Time`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *InlineResponse20035) GetStartTsOk() (*time.Time, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *InlineResponse20035) SetStartTs(v time.Time)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *InlineResponse20035) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetEndTs

`func (o *InlineResponse20035) GetEndTs() time.Time`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *InlineResponse20035) GetEndTsOk() (*time.Time, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *InlineResponse20035) SetEndTs(v time.Time)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *InlineResponse20035) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetCounts

`func (o *InlineResponse20035) GetCounts() NetworksNetworkIdSensorAlertsOverviewByMetricCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *InlineResponse20035) GetCountsOk() (*NetworksNetworkIdSensorAlertsOverviewByMetricCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *InlineResponse20035) SetCounts(v NetworksNetworkIdSensorAlertsOverviewByMetricCounts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *InlineResponse20035) HasCounts() bool`

HasCounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


