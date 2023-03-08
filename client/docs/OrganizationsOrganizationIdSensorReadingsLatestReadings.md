# OrganizationsOrganizationIdSensorReadingsLatestReadings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | Pointer to **string** | Time at which the reading occurred, in ISO8601 format. | [optional] 
**Metric** | Pointer to **string** | Type of sensor reading. | [optional] 
**Battery** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryBattery**](OrganizationsOrganizationIdSensorReadingsHistoryBattery.md) |  | [optional] 
**Button** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryButton**](OrganizationsOrganizationIdSensorReadingsHistoryButton.md) |  | [optional] 
**Door** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryDoor**](OrganizationsOrganizationIdSensorReadingsHistoryDoor.md) |  | [optional] 
**Humidity** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryHumidity**](OrganizationsOrganizationIdSensorReadingsHistoryHumidity.md) |  | [optional] 
**IndoorAirQuality** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality**](OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality.md) |  | [optional] 
**Noise** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryNoise**](OrganizationsOrganizationIdSensorReadingsHistoryNoise.md) |  | [optional] 
**Pm25** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryPm25**](OrganizationsOrganizationIdSensorReadingsHistoryPm25.md) |  | [optional] 
**Temperature** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryTemperature**](OrganizationsOrganizationIdSensorReadingsHistoryTemperature.md) |  | [optional] 
**Tvoc** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryTvoc**](OrganizationsOrganizationIdSensorReadingsHistoryTvoc.md) |  | [optional] 
**Water** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryWater**](OrganizationsOrganizationIdSensorReadingsHistoryWater.md) |  | [optional] 

## Methods

### NewOrganizationsOrganizationIdSensorReadingsLatestReadings

`func NewOrganizationsOrganizationIdSensorReadingsLatestReadings() *OrganizationsOrganizationIdSensorReadingsLatestReadings`

NewOrganizationsOrganizationIdSensorReadingsLatestReadings instantiates a new OrganizationsOrganizationIdSensorReadingsLatestReadings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsOrganizationIdSensorReadingsLatestReadingsWithDefaults

`func NewOrganizationsOrganizationIdSensorReadingsLatestReadingsWithDefaults() *OrganizationsOrganizationIdSensorReadingsLatestReadings`

NewOrganizationsOrganizationIdSensorReadingsLatestReadingsWithDefaults instantiates a new OrganizationsOrganizationIdSensorReadingsLatestReadings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetMetric

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetBattery

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetBattery() OrganizationsOrganizationIdSensorReadingsHistoryBattery`

GetBattery returns the Battery field if non-nil, zero value otherwise.

### GetBatteryOk

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetBatteryOk() (*OrganizationsOrganizationIdSensorReadingsHistoryBattery, bool)`

GetBatteryOk returns a tuple with the Battery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBattery

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetBattery(v OrganizationsOrganizationIdSensorReadingsHistoryBattery)`

SetBattery sets Battery field to given value.

### HasBattery

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) HasBattery() bool`

HasBattery returns a boolean if a field has been set.

### GetButton

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetButton() OrganizationsOrganizationIdSensorReadingsHistoryButton`

GetButton returns the Button field if non-nil, zero value otherwise.

### GetButtonOk

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetButtonOk() (*OrganizationsOrganizationIdSensorReadingsHistoryButton, bool)`

GetButtonOk returns a tuple with the Button field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButton

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetButton(v OrganizationsOrganizationIdSensorReadingsHistoryButton)`

SetButton sets Button field to given value.

### HasButton

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) HasButton() bool`

HasButton returns a boolean if a field has been set.

### GetDoor

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetDoor() OrganizationsOrganizationIdSensorReadingsHistoryDoor`

GetDoor returns the Door field if non-nil, zero value otherwise.

### GetDoorOk

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetDoorOk() (*OrganizationsOrganizationIdSensorReadingsHistoryDoor, bool)`

GetDoorOk returns a tuple with the Door field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoor

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetDoor(v OrganizationsOrganizationIdSensorReadingsHistoryDoor)`

SetDoor sets Door field to given value.

### HasDoor

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) HasDoor() bool`

HasDoor returns a boolean if a field has been set.

### GetHumidity

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetHumidity() OrganizationsOrganizationIdSensorReadingsHistoryHumidity`

GetHumidity returns the Humidity field if non-nil, zero value otherwise.

### GetHumidityOk

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetHumidityOk() (*OrganizationsOrganizationIdSensorReadingsHistoryHumidity, bool)`

GetHumidityOk returns a tuple with the Humidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumidity

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetHumidity(v OrganizationsOrganizationIdSensorReadingsHistoryHumidity)`

SetHumidity sets Humidity field to given value.

### HasHumidity

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) HasHumidity() bool`

HasHumidity returns a boolean if a field has been set.

### GetIndoorAirQuality

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetIndoorAirQuality() OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality`

GetIndoorAirQuality returns the IndoorAirQuality field if non-nil, zero value otherwise.

### GetIndoorAirQualityOk

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetIndoorAirQualityOk() (*OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality, bool)`

GetIndoorAirQualityOk returns a tuple with the IndoorAirQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndoorAirQuality

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetIndoorAirQuality(v OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality)`

SetIndoorAirQuality sets IndoorAirQuality field to given value.

### HasIndoorAirQuality

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) HasIndoorAirQuality() bool`

HasIndoorAirQuality returns a boolean if a field has been set.

### GetNoise

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetNoise() OrganizationsOrganizationIdSensorReadingsHistoryNoise`

GetNoise returns the Noise field if non-nil, zero value otherwise.

### GetNoiseOk

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetNoiseOk() (*OrganizationsOrganizationIdSensorReadingsHistoryNoise, bool)`

GetNoiseOk returns a tuple with the Noise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoise

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetNoise(v OrganizationsOrganizationIdSensorReadingsHistoryNoise)`

SetNoise sets Noise field to given value.

### HasNoise

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) HasNoise() bool`

HasNoise returns a boolean if a field has been set.

### GetPm25

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetPm25() OrganizationsOrganizationIdSensorReadingsHistoryPm25`

GetPm25 returns the Pm25 field if non-nil, zero value otherwise.

### GetPm25Ok

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetPm25Ok() (*OrganizationsOrganizationIdSensorReadingsHistoryPm25, bool)`

GetPm25Ok returns a tuple with the Pm25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPm25

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetPm25(v OrganizationsOrganizationIdSensorReadingsHistoryPm25)`

SetPm25 sets Pm25 field to given value.

### HasPm25

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) HasPm25() bool`

HasPm25 returns a boolean if a field has been set.

### GetTemperature

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetTemperature() OrganizationsOrganizationIdSensorReadingsHistoryTemperature`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetTemperatureOk() (*OrganizationsOrganizationIdSensorReadingsHistoryTemperature, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetTemperature(v OrganizationsOrganizationIdSensorReadingsHistoryTemperature)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetTvoc

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetTvoc() OrganizationsOrganizationIdSensorReadingsHistoryTvoc`

GetTvoc returns the Tvoc field if non-nil, zero value otherwise.

### GetTvocOk

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetTvocOk() (*OrganizationsOrganizationIdSensorReadingsHistoryTvoc, bool)`

GetTvocOk returns a tuple with the Tvoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTvoc

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetTvoc(v OrganizationsOrganizationIdSensorReadingsHistoryTvoc)`

SetTvoc sets Tvoc field to given value.

### HasTvoc

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) HasTvoc() bool`

HasTvoc returns a boolean if a field has been set.

### GetWater

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetWater() OrganizationsOrganizationIdSensorReadingsHistoryWater`

GetWater returns the Water field if non-nil, zero value otherwise.

### GetWaterOk

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) GetWaterOk() (*OrganizationsOrganizationIdSensorReadingsHistoryWater, bool)`

GetWaterOk returns a tuple with the Water field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWater

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) SetWater(v OrganizationsOrganizationIdSensorReadingsHistoryWater)`

SetWater sets Water field to given value.

### HasWater

`func (o *OrganizationsOrganizationIdSensorReadingsLatestReadings) HasWater() bool`

HasWater returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


