# GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Door** | Pointer to **int32** | Number of sensors that are currently alerting due to an open door | [optional] 
**Humidity** | Pointer to **int32** | Number of sensors that are currently alerting due to humidity readings | [optional] 
**IndoorAirQuality** | Pointer to **int32** | Number of sensors that are currently alerting due to indoor air quality readings | [optional] 
**Noise** | Pointer to [**GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCountsNoise**](GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCountsNoise.md) |  | [optional] 
**Pm25** | Pointer to **int32** | Number of sensors that are currently alerting due to PM2.5 readings | [optional] 
**Temperature** | Pointer to **int32** | Number of sensors that are currently alerting due to temperature readings | [optional] 
**Tvoc** | Pointer to **int32** | Number of sensors that are currently alerting due to TVOC readings | [optional] 
**Water** | Pointer to **int32** | Number of sensors that are currently alerting due to the presence of water | [optional] 

## Methods

### NewGetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts

`func NewGetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts() *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts`

NewGetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts instantiates a new GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCountsWithDefaults

`func NewGetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCountsWithDefaults() *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts`

NewGetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCountsWithDefaults instantiates a new GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDoor

`func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) GetDoor() int32`

GetDoor returns the Door field if non-nil, zero value otherwise.

### GetDoorOk

`func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) GetDoorOk() (*int32, bool)`

GetDoorOk returns a tuple with the Door field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoor

`func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) SetDoor(v int32)`

SetDoor sets Door field to given value.

### HasDoor

`func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) HasDoor() bool`

HasDoor returns a boolean if a field has been set.

### GetHumidity

`func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) GetHumidity() int32`

GetHumidity returns the Humidity field if non-nil, zero value otherwise.

### GetHumidityOk

`func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) GetHumidityOk() (*int32, bool)`

GetHumidityOk returns a tuple with the Humidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumidity

`func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) SetHumidity(v int32)`

SetHumidity sets Humidity field to given value.

### HasHumidity

`func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) HasHumidity() bool`

HasHumidity returns a boolean if a field has been set.

### GetIndoorAirQuality

`func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) GetIndoorAirQuality() int32`

GetIndoorAirQuality returns the IndoorAirQuality field if non-nil, zero value otherwise.

### GetIndoorAirQualityOk

`func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) GetIndoorAirQualityOk() (*int32, bool)`

GetIndoorAirQualityOk returns a tuple with the IndoorAirQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndoorAirQuality

`func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) SetIndoorAirQuality(v int32)`

SetIndoorAirQuality sets IndoorAirQuality field to given value.

### HasIndoorAirQuality

`func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) HasIndoorAirQuality() bool`

HasIndoorAirQuality returns a boolean if a field has been set.

### GetNoise

`func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) GetNoise() GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCountsNoise`

GetNoise returns the Noise field if non-nil, zero value otherwise.

### GetNoiseOk

`func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) GetNoiseOk() (*GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCountsNoise, bool)`

GetNoiseOk returns a tuple with the Noise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoise

`func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) SetNoise(v GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCountsNoise)`

SetNoise sets Noise field to given value.

### HasNoise

`func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) HasNoise() bool`

HasNoise returns a boolean if a field has been set.

### GetPm25

`func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) GetPm25() int32`

GetPm25 returns the Pm25 field if non-nil, zero value otherwise.

### GetPm25Ok

`func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) GetPm25Ok() (*int32, bool)`

GetPm25Ok returns a tuple with the Pm25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPm25

`func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) SetPm25(v int32)`

SetPm25 sets Pm25 field to given value.

### HasPm25

`func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) HasPm25() bool`

HasPm25 returns a boolean if a field has been set.

### GetTemperature

`func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) GetTemperature() int32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) GetTemperatureOk() (*int32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) SetTemperature(v int32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetTvoc

`func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) GetTvoc() int32`

GetTvoc returns the Tvoc field if non-nil, zero value otherwise.

### GetTvocOk

`func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) GetTvocOk() (*int32, bool)`

GetTvocOk returns a tuple with the Tvoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTvoc

`func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) SetTvoc(v int32)`

SetTvoc sets Tvoc field to given value.

### HasTvoc

`func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) HasTvoc() bool`

HasTvoc returns a boolean if a field has been set.

### GetWater

`func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) GetWater() int32`

GetWater returns the Water field if non-nil, zero value otherwise.

### GetWaterOk

`func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) GetWaterOk() (*int32, bool)`

GetWaterOk returns a tuple with the Water field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWater

`func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) SetWater(v int32)`

SetWater sets Water field to given value.

### HasWater

`func (o *GetNetworkSensorAlertsCurrentOverviewByMetric200ResponseCounts) HasWater() bool`

HasWater returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


