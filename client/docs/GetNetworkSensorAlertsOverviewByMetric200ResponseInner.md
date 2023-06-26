# GetNetworkSensorAlertsOverviewByMetric200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTs** | Pointer to **time.Time** | Start of the timespan over which sensor alerts are counted | [optional] 
**EndTs** | Pointer to **time.Time** | End of the timespan over which sensor alerts are counted | [optional] 
**Counts** | Pointer to [**GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts**](GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts.md) |  | [optional] 

## Methods

### NewGetNetworkSensorAlertsOverviewByMetric200ResponseInner

`func NewGetNetworkSensorAlertsOverviewByMetric200ResponseInner() *GetNetworkSensorAlertsOverviewByMetric200ResponseInner`

NewGetNetworkSensorAlertsOverviewByMetric200ResponseInner instantiates a new GetNetworkSensorAlertsOverviewByMetric200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSensorAlertsOverviewByMetric200ResponseInnerWithDefaults

`func NewGetNetworkSensorAlertsOverviewByMetric200ResponseInnerWithDefaults() *GetNetworkSensorAlertsOverviewByMetric200ResponseInner`

NewGetNetworkSensorAlertsOverviewByMetric200ResponseInnerWithDefaults instantiates a new GetNetworkSensorAlertsOverviewByMetric200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTs

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInner) GetStartTs() time.Time`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInner) GetStartTsOk() (*time.Time, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInner) SetStartTs(v time.Time)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInner) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetEndTs

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInner) GetEndTs() time.Time`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInner) GetEndTsOk() (*time.Time, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInner) SetEndTs(v time.Time)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInner) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetCounts

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInner) GetCounts() GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInner) GetCountsOk() (*GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInner) SetCounts(v GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInner) HasCounts() bool`

HasCounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


