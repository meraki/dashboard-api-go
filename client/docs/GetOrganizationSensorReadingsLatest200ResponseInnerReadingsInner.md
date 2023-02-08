# GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | Pointer to **string** | Time at which the reading occurred, in ISO8601 format. | [optional] 
**Metric** | Pointer to **string** | Type of sensor reading. | [optional] 
**Battery** | Pointer to [**GetOrganizationSensorReadingsHistory200ResponseInnerBattery**](GetOrganizationSensorReadingsHistory200ResponseInnerBattery.md) |  | [optional] 
**Button** | Pointer to [**GetOrganizationSensorReadingsHistory200ResponseInnerButton**](GetOrganizationSensorReadingsHistory200ResponseInnerButton.md) |  | [optional] 
**Door** | Pointer to [**GetOrganizationSensorReadingsHistory200ResponseInnerDoor**](GetOrganizationSensorReadingsHistory200ResponseInnerDoor.md) |  | [optional] 
**Humidity** | Pointer to [**GetOrganizationSensorReadingsHistory200ResponseInnerHumidity**](GetOrganizationSensorReadingsHistory200ResponseInnerHumidity.md) |  | [optional] 
**IndoorAirQuality** | Pointer to [**GetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality**](GetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality.md) |  | [optional] 
**Noise** | Pointer to [**GetOrganizationSensorReadingsHistory200ResponseInnerNoise**](GetOrganizationSensorReadingsHistory200ResponseInnerNoise.md) |  | [optional] 
**Pm25** | Pointer to [**GetOrganizationSensorReadingsHistory200ResponseInnerPm25**](GetOrganizationSensorReadingsHistory200ResponseInnerPm25.md) |  | [optional] 
**Temperature** | Pointer to [**GetOrganizationSensorReadingsHistory200ResponseInnerTemperature**](GetOrganizationSensorReadingsHistory200ResponseInnerTemperature.md) |  | [optional] 
**Tvoc** | Pointer to [**GetOrganizationSensorReadingsHistory200ResponseInnerTvoc**](GetOrganizationSensorReadingsHistory200ResponseInnerTvoc.md) |  | [optional] 
**Water** | Pointer to [**GetOrganizationSensorReadingsHistory200ResponseInnerWater**](GetOrganizationSensorReadingsHistory200ResponseInnerWater.md) |  | [optional] 

## Methods

### NewGetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner

`func NewGetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner() *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner`

NewGetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner instantiates a new GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationSensorReadingsLatest200ResponseInnerReadingsInnerWithDefaults

`func NewGetOrganizationSensorReadingsLatest200ResponseInnerReadingsInnerWithDefaults() *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner`

NewGetOrganizationSensorReadingsLatest200ResponseInnerReadingsInnerWithDefaults instantiates a new GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetMetric

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetBattery

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetBattery() GetOrganizationSensorReadingsHistory200ResponseInnerBattery`

GetBattery returns the Battery field if non-nil, zero value otherwise.

### GetBatteryOk

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetBatteryOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerBattery, bool)`

GetBatteryOk returns a tuple with the Battery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBattery

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) SetBattery(v GetOrganizationSensorReadingsHistory200ResponseInnerBattery)`

SetBattery sets Battery field to given value.

### HasBattery

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) HasBattery() bool`

HasBattery returns a boolean if a field has been set.

### GetButton

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetButton() GetOrganizationSensorReadingsHistory200ResponseInnerButton`

GetButton returns the Button field if non-nil, zero value otherwise.

### GetButtonOk

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetButtonOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerButton, bool)`

GetButtonOk returns a tuple with the Button field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButton

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) SetButton(v GetOrganizationSensorReadingsHistory200ResponseInnerButton)`

SetButton sets Button field to given value.

### HasButton

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) HasButton() bool`

HasButton returns a boolean if a field has been set.

### GetDoor

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetDoor() GetOrganizationSensorReadingsHistory200ResponseInnerDoor`

GetDoor returns the Door field if non-nil, zero value otherwise.

### GetDoorOk

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetDoorOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerDoor, bool)`

GetDoorOk returns a tuple with the Door field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoor

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) SetDoor(v GetOrganizationSensorReadingsHistory200ResponseInnerDoor)`

SetDoor sets Door field to given value.

### HasDoor

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) HasDoor() bool`

HasDoor returns a boolean if a field has been set.

### GetHumidity

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetHumidity() GetOrganizationSensorReadingsHistory200ResponseInnerHumidity`

GetHumidity returns the Humidity field if non-nil, zero value otherwise.

### GetHumidityOk

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetHumidityOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerHumidity, bool)`

GetHumidityOk returns a tuple with the Humidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumidity

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) SetHumidity(v GetOrganizationSensorReadingsHistory200ResponseInnerHumidity)`

SetHumidity sets Humidity field to given value.

### HasHumidity

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) HasHumidity() bool`

HasHumidity returns a boolean if a field has been set.

### GetIndoorAirQuality

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetIndoorAirQuality() GetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality`

GetIndoorAirQuality returns the IndoorAirQuality field if non-nil, zero value otherwise.

### GetIndoorAirQualityOk

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetIndoorAirQualityOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality, bool)`

GetIndoorAirQualityOk returns a tuple with the IndoorAirQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndoorAirQuality

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) SetIndoorAirQuality(v GetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality)`

SetIndoorAirQuality sets IndoorAirQuality field to given value.

### HasIndoorAirQuality

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) HasIndoorAirQuality() bool`

HasIndoorAirQuality returns a boolean if a field has been set.

### GetNoise

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetNoise() GetOrganizationSensorReadingsHistory200ResponseInnerNoise`

GetNoise returns the Noise field if non-nil, zero value otherwise.

### GetNoiseOk

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetNoiseOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerNoise, bool)`

GetNoiseOk returns a tuple with the Noise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoise

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) SetNoise(v GetOrganizationSensorReadingsHistory200ResponseInnerNoise)`

SetNoise sets Noise field to given value.

### HasNoise

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) HasNoise() bool`

HasNoise returns a boolean if a field has been set.

### GetPm25

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetPm25() GetOrganizationSensorReadingsHistory200ResponseInnerPm25`

GetPm25 returns the Pm25 field if non-nil, zero value otherwise.

### GetPm25Ok

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetPm25Ok() (*GetOrganizationSensorReadingsHistory200ResponseInnerPm25, bool)`

GetPm25Ok returns a tuple with the Pm25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPm25

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) SetPm25(v GetOrganizationSensorReadingsHistory200ResponseInnerPm25)`

SetPm25 sets Pm25 field to given value.

### HasPm25

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) HasPm25() bool`

HasPm25 returns a boolean if a field has been set.

### GetTemperature

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetTemperature() GetOrganizationSensorReadingsHistory200ResponseInnerTemperature`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetTemperatureOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerTemperature, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) SetTemperature(v GetOrganizationSensorReadingsHistory200ResponseInnerTemperature)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetTvoc

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetTvoc() GetOrganizationSensorReadingsHistory200ResponseInnerTvoc`

GetTvoc returns the Tvoc field if non-nil, zero value otherwise.

### GetTvocOk

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetTvocOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerTvoc, bool)`

GetTvocOk returns a tuple with the Tvoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTvoc

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) SetTvoc(v GetOrganizationSensorReadingsHistory200ResponseInnerTvoc)`

SetTvoc sets Tvoc field to given value.

### HasTvoc

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) HasTvoc() bool`

HasTvoc returns a boolean if a field has been set.

### GetWater

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetWater() GetOrganizationSensorReadingsHistory200ResponseInnerWater`

GetWater returns the Water field if non-nil, zero value otherwise.

### GetWaterOk

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetWaterOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerWater, bool)`

GetWaterOk returns a tuple with the Water field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWater

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) SetWater(v GetOrganizationSensorReadingsHistory200ResponseInnerWater)`

SetWater sets Water field to given value.

### HasWater

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) HasWater() bool`

HasWater returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


