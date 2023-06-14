# GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Door** | Pointer to **int32** | Number of sensor alerts that occurred due to an open door | [optional] 
**Humidity** | Pointer to **int32** | Number of sensor alerts that occurred due to humidity readings | [optional] 
**IndoorAirQuality** | Pointer to **int32** | Number of sensor alerts that occurred due to indoor air quality readings | [optional] 
**Noise** | Pointer to [**GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise**](GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise.md) |  | [optional] 
**Pm25** | Pointer to **int32** | Number of sensor alerts that occurred due to PM2.5 readings | [optional] 
**Temperature** | Pointer to **int32** | Number of sensor alerts that occurred due to temperature readings | [optional] 
**Tvoc** | Pointer to **int32** | Number of sensor alerts that occurred due to TVOC readings | [optional] 
**Water** | Pointer to **int32** | Number of sensor alerts that occurred due to the presence of water | [optional] 

## Methods

### NewGetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts

`func NewGetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts() *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts`

NewGetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts instantiates a new GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsWithDefaults

`func NewGetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsWithDefaults() *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts`

NewGetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsWithDefaults instantiates a new GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDoor

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetDoor() int32`

GetDoor returns the Door field if non-nil, zero value otherwise.

### GetDoorOk

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetDoorOk() (*int32, bool)`

GetDoorOk returns a tuple with the Door field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoor

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) SetDoor(v int32)`

SetDoor sets Door field to given value.

### HasDoor

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) HasDoor() bool`

HasDoor returns a boolean if a field has been set.

### GetHumidity

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetHumidity() int32`

GetHumidity returns the Humidity field if non-nil, zero value otherwise.

### GetHumidityOk

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetHumidityOk() (*int32, bool)`

GetHumidityOk returns a tuple with the Humidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumidity

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) SetHumidity(v int32)`

SetHumidity sets Humidity field to given value.

### HasHumidity

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) HasHumidity() bool`

HasHumidity returns a boolean if a field has been set.

### GetIndoorAirQuality

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetIndoorAirQuality() int32`

GetIndoorAirQuality returns the IndoorAirQuality field if non-nil, zero value otherwise.

### GetIndoorAirQualityOk

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetIndoorAirQualityOk() (*int32, bool)`

GetIndoorAirQualityOk returns a tuple with the IndoorAirQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndoorAirQuality

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) SetIndoorAirQuality(v int32)`

SetIndoorAirQuality sets IndoorAirQuality field to given value.

### HasIndoorAirQuality

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) HasIndoorAirQuality() bool`

HasIndoorAirQuality returns a boolean if a field has been set.

### GetNoise

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetNoise() GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise`

GetNoise returns the Noise field if non-nil, zero value otherwise.

### GetNoiseOk

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetNoiseOk() (*GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise, bool)`

GetNoiseOk returns a tuple with the Noise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoise

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) SetNoise(v GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise)`

SetNoise sets Noise field to given value.

### HasNoise

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) HasNoise() bool`

HasNoise returns a boolean if a field has been set.

### GetPm25

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetPm25() int32`

GetPm25 returns the Pm25 field if non-nil, zero value otherwise.

### GetPm25Ok

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetPm25Ok() (*int32, bool)`

GetPm25Ok returns a tuple with the Pm25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPm25

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) SetPm25(v int32)`

SetPm25 sets Pm25 field to given value.

### HasPm25

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) HasPm25() bool`

HasPm25 returns a boolean if a field has been set.

### GetTemperature

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetTemperature() int32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetTemperatureOk() (*int32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) SetTemperature(v int32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetTvoc

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetTvoc() int32`

GetTvoc returns the Tvoc field if non-nil, zero value otherwise.

### GetTvocOk

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetTvocOk() (*int32, bool)`

GetTvocOk returns a tuple with the Tvoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTvoc

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) SetTvoc(v int32)`

SetTvoc sets Tvoc field to given value.

### HasTvoc

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) HasTvoc() bool`

HasTvoc returns a boolean if a field has been set.

### GetWater

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetWater() int32`

GetWater returns the Water field if non-nil, zero value otherwise.

### GetWaterOk

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetWaterOk() (*int32, bool)`

GetWaterOk returns a tuple with the Water field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWater

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) SetWater(v int32)`

SetWater sets Water field to given value.

### HasWater

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) HasWater() bool`

HasWater returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


