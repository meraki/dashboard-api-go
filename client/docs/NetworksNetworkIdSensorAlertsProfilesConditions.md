# NetworksNetworkIdSensorAlertsProfilesConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | **string** | The type of sensor metric that will be monitored for changes. Available metrics are door, humidity, indoorAirQuality, noise, pm25, temperature, tvoc, and water. | 
**Threshold** | [**NetworksNetworkIdSensorAlertsProfilesThreshold**](NetworksNetworkIdSensorAlertsProfilesThreshold.md) |  | 
**Direction** | Pointer to **string** | If &#39;above&#39;, an alert will be sent when a sensor reads above the threshold. If &#39;below&#39;, an alert will be sent when a sensor reads below the threshold. Only applicable for temperature and humidity thresholds. | [optional] 
**Duration** | Pointer to **int32** | Length of time in seconds that the triggering state must persist before an alert is sent. Available options are 0 seconds, 1 minute, 2 minutes, 3 minutes, 4 minutes, 5 minutes, 10 minutes, 15 minutes, 30 minutes, and 1 hour. Default is 0. | [optional] 

## Methods

### NewNetworksNetworkIdSensorAlertsProfilesConditions

`func NewNetworksNetworkIdSensorAlertsProfilesConditions(metric string, threshold NetworksNetworkIdSensorAlertsProfilesThreshold, ) *NetworksNetworkIdSensorAlertsProfilesConditions`

NewNetworksNetworkIdSensorAlertsProfilesConditions instantiates a new NetworksNetworkIdSensorAlertsProfilesConditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdSensorAlertsProfilesConditionsWithDefaults

`func NewNetworksNetworkIdSensorAlertsProfilesConditionsWithDefaults() *NetworksNetworkIdSensorAlertsProfilesConditions`

NewNetworksNetworkIdSensorAlertsProfilesConditionsWithDefaults instantiates a new NetworksNetworkIdSensorAlertsProfilesConditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *NetworksNetworkIdSensorAlertsProfilesConditions) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *NetworksNetworkIdSensorAlertsProfilesConditions) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *NetworksNetworkIdSensorAlertsProfilesConditions) SetMetric(v string)`

SetMetric sets Metric field to given value.


### GetThreshold

`func (o *NetworksNetworkIdSensorAlertsProfilesConditions) GetThreshold() NetworksNetworkIdSensorAlertsProfilesThreshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *NetworksNetworkIdSensorAlertsProfilesConditions) GetThresholdOk() (*NetworksNetworkIdSensorAlertsProfilesThreshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *NetworksNetworkIdSensorAlertsProfilesConditions) SetThreshold(v NetworksNetworkIdSensorAlertsProfilesThreshold)`

SetThreshold sets Threshold field to given value.


### GetDirection

`func (o *NetworksNetworkIdSensorAlertsProfilesConditions) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *NetworksNetworkIdSensorAlertsProfilesConditions) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *NetworksNetworkIdSensorAlertsProfilesConditions) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *NetworksNetworkIdSensorAlertsProfilesConditions) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetDuration

`func (o *NetworksNetworkIdSensorAlertsProfilesConditions) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *NetworksNetworkIdSensorAlertsProfilesConditions) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *NetworksNetworkIdSensorAlertsProfilesConditions) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *NetworksNetworkIdSensorAlertsProfilesConditions) HasDuration() bool`

HasDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


