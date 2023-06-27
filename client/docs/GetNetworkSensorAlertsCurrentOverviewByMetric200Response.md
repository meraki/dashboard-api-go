# GetNetworkSensorAlertsCurrentOverviewByMetric200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportedMetrics** | Pointer to **[]string** | List of metrics that are supported for alerts, based on available sensor devices in the network | [optional] 
**Counts** | Pointer to [**GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts**](GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts.md) |  | [optional] 

## Methods

### NewGetNetworkSensorAlertsCurrentOverviewByMetric200Response

`func NewGetNetworkSensorAlertsCurrentOverviewByMetric200Response() *GetNetworkSensorAlertsCurrentOverviewByMetric200Response`

NewGetNetworkSensorAlertsCurrentOverviewByMetric200Response instantiates a new GetNetworkSensorAlertsCurrentOverviewByMetric200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSensorAlertsCurrentOverviewByMetric200ResponseWithDefaults

`func NewGetNetworkSensorAlertsCurrentOverviewByMetric200ResponseWithDefaults() *GetNetworkSensorAlertsCurrentOverviewByMetric200Response`

NewGetNetworkSensorAlertsCurrentOverviewByMetric200ResponseWithDefaults instantiates a new GetNetworkSensorAlertsCurrentOverviewByMetric200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportedMetrics

`func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200Response) GetSupportedMetrics() []string`

GetSupportedMetrics returns the SupportedMetrics field if non-nil, zero value otherwise.

### GetSupportedMetricsOk

`func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200Response) GetSupportedMetricsOk() (*[]string, bool)`

GetSupportedMetricsOk returns a tuple with the SupportedMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedMetrics

`func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200Response) SetSupportedMetrics(v []string)`

SetSupportedMetrics sets SupportedMetrics field to given value.

### HasSupportedMetrics

`func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200Response) HasSupportedMetrics() bool`

HasSupportedMetrics returns a boolean if a field has been set.

### GetCounts

`func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200Response) GetCounts() GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200Response) GetCountsOk() (*GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200Response) SetCounts(v GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200Response) HasCounts() bool`

HasCounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


