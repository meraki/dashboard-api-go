# GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | **string** | The type of sensor metric that will be monitored for changes. Available metrics are door, humidity, indoorAirQuality, noise, pm25, temperature, tvoc, and water. | 
**Threshold** | [**GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThreshold**](GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThreshold.md) |  | 
**Direction** | Pointer to **string** | If &#39;above&#39;, an alert will be sent when a sensor reads above the threshold. If &#39;below&#39;, an alert will be sent when a sensor reads below the threshold. Only applicable for temperature and humidity thresholds. | [optional] 
**Duration** | Pointer to **int32** | Length of time in seconds that the triggering state must persist before an alert is sent. Available options are 0 seconds, 1 minute, 2 minutes, 3 minutes, 4 minutes, 5 minutes, 10 minutes, 15 minutes, 30 minutes, and 1 hour. Default is 0. | [optional] 

## Methods

### NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner

`func NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner(metric string, threshold GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThreshold, ) *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner`

NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner instantiates a new GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerWithDefaults

`func NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerWithDefaults() *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner`

NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerWithDefaults instantiates a new GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) SetMetric(v string)`

SetMetric sets Metric field to given value.


### GetThreshold

`func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) GetThreshold() GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThreshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) GetThresholdOk() (*GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThreshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) SetThreshold(v GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThreshold)`

SetThreshold sets Threshold field to given value.


### GetDirection

`func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetDuration

`func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) HasDuration() bool`

HasDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


