# GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApparentPower** | Pointer to **int32** | Number of sensor alerts that occurred due to apparent power readings | [optional] 
**Co2** | Pointer to **int32** | Number of sensors that are currently alerting due to CO2 readings | [optional] 
**Current** | Pointer to **int32** | Number of sensor alerts that occurred due to electrical current readings | [optional] 
**Door** | Pointer to **int32** | Number of sensor alerts that occurred due to an open door | [optional] 
**Frequency** | Pointer to **int32** | Number of sensor alerts that occurred due to frequency readings | [optional] 
**Humidity** | Pointer to **int32** | Number of sensor alerts that occurred due to humidity readings | [optional] 
**IndoorAirQuality** | Pointer to **int32** | Number of sensor alerts that occurred due to indoor air quality readings | [optional] 
**Noise** | Pointer to [**GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise**](GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCountsNoise.md) |  | [optional] 
**Pm25** | Pointer to **int32** | Number of sensor alerts that occurred due to PM2.5 readings | [optional] 
**PowerFactor** | Pointer to **int32** | Number of sensor alerts that occurred due to power factor readings | [optional] 
**RealPower** | Pointer to **int32** | Number of sensor alerts that occurred due to real power readings | [optional] 
**Temperature** | Pointer to **int32** | Number of sensor alerts that occurred due to temperature readings | [optional] 
**Tvoc** | Pointer to **int32** | Number of sensor alerts that occurred due to TVOC readings | [optional] 
**UpstreamPower** | Pointer to **int32** | Number of sensor alerts that occurred due to upstream power outages | [optional] 
**Voltage** | Pointer to **int32** | Number of sensor alerts that occurred due to voltage readings | [optional] 
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

### GetApparentPower

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetApparentPower() int32`

GetApparentPower returns the ApparentPower field if non-nil, zero value otherwise.

### GetApparentPowerOk

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetApparentPowerOk() (*int32, bool)`

GetApparentPowerOk returns a tuple with the ApparentPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApparentPower

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) SetApparentPower(v int32)`

SetApparentPower sets ApparentPower field to given value.

### HasApparentPower

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) HasApparentPower() bool`

HasApparentPower returns a boolean if a field has been set.

### GetCo2

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetCo2() int32`

GetCo2 returns the Co2 field if non-nil, zero value otherwise.

### GetCo2Ok

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetCo2Ok() (*int32, bool)`

GetCo2Ok returns a tuple with the Co2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCo2

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) SetCo2(v int32)`

SetCo2 sets Co2 field to given value.

### HasCo2

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) HasCo2() bool`

HasCo2 returns a boolean if a field has been set.

### GetCurrent

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetCurrent() int32`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetCurrentOk() (*int32, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) SetCurrent(v int32)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

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

### GetFrequency

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetFrequency() int32`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetFrequencyOk() (*int32, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) SetFrequency(v int32)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

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

### GetPowerFactor

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetPowerFactor() int32`

GetPowerFactor returns the PowerFactor field if non-nil, zero value otherwise.

### GetPowerFactorOk

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetPowerFactorOk() (*int32, bool)`

GetPowerFactorOk returns a tuple with the PowerFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerFactor

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) SetPowerFactor(v int32)`

SetPowerFactor sets PowerFactor field to given value.

### HasPowerFactor

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) HasPowerFactor() bool`

HasPowerFactor returns a boolean if a field has been set.

### GetRealPower

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetRealPower() int32`

GetRealPower returns the RealPower field if non-nil, zero value otherwise.

### GetRealPowerOk

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetRealPowerOk() (*int32, bool)`

GetRealPowerOk returns a tuple with the RealPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealPower

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) SetRealPower(v int32)`

SetRealPower sets RealPower field to given value.

### HasRealPower

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) HasRealPower() bool`

HasRealPower returns a boolean if a field has been set.

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

### GetUpstreamPower

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetUpstreamPower() int32`

GetUpstreamPower returns the UpstreamPower field if non-nil, zero value otherwise.

### GetUpstreamPowerOk

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetUpstreamPowerOk() (*int32, bool)`

GetUpstreamPowerOk returns a tuple with the UpstreamPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamPower

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) SetUpstreamPower(v int32)`

SetUpstreamPower sets UpstreamPower field to given value.

### HasUpstreamPower

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) HasUpstreamPower() bool`

HasUpstreamPower returns a boolean if a field has been set.

### GetVoltage

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetVoltage() int32`

GetVoltage returns the Voltage field if non-nil, zero value otherwise.

### GetVoltageOk

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) GetVoltageOk() (*int32, bool)`

GetVoltageOk returns a tuple with the Voltage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoltage

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) SetVoltage(v int32)`

SetVoltage sets Voltage field to given value.

### HasVoltage

`func (o *GetNetworkSensorAlertsOverviewByMetric200ResponseInnerCounts) HasVoltage() bool`

HasVoltage returns a boolean if a field has been set.

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


