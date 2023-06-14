# GetOrganizationSensorReadingsHistory200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Serial number of the sensor that took the reading. | [optional] 
**Network** | Pointer to [**GetOrganizationSensorReadingsHistory200ResponseInnerNetwork**](GetOrganizationSensorReadingsHistory200ResponseInnerNetwork.md) |  | [optional] 
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

### NewGetOrganizationSensorReadingsHistory200ResponseInner

`func NewGetOrganizationSensorReadingsHistory200ResponseInner() *GetOrganizationSensorReadingsHistory200ResponseInner`

NewGetOrganizationSensorReadingsHistory200ResponseInner instantiates a new GetOrganizationSensorReadingsHistory200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationSensorReadingsHistory200ResponseInnerWithDefaults

`func NewGetOrganizationSensorReadingsHistory200ResponseInnerWithDefaults() *GetOrganizationSensorReadingsHistory200ResponseInner`

NewGetOrganizationSensorReadingsHistory200ResponseInnerWithDefaults instantiates a new GetOrganizationSensorReadingsHistory200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetNetwork

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetNetwork() GetOrganizationSensorReadingsHistory200ResponseInnerNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetNetworkOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) SetNetwork(v GetOrganizationSensorReadingsHistory200ResponseInnerNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetTs

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetMetric

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetBattery

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetBattery() GetOrganizationSensorReadingsHistory200ResponseInnerBattery`

GetBattery returns the Battery field if non-nil, zero value otherwise.

### GetBatteryOk

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetBatteryOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerBattery, bool)`

GetBatteryOk returns a tuple with the Battery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBattery

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) SetBattery(v GetOrganizationSensorReadingsHistory200ResponseInnerBattery)`

SetBattery sets Battery field to given value.

### HasBattery

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) HasBattery() bool`

HasBattery returns a boolean if a field has been set.

### GetButton

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetButton() GetOrganizationSensorReadingsHistory200ResponseInnerButton`

GetButton returns the Button field if non-nil, zero value otherwise.

### GetButtonOk

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetButtonOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerButton, bool)`

GetButtonOk returns a tuple with the Button field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButton

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) SetButton(v GetOrganizationSensorReadingsHistory200ResponseInnerButton)`

SetButton sets Button field to given value.

### HasButton

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) HasButton() bool`

HasButton returns a boolean if a field has been set.

### GetDoor

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetDoor() GetOrganizationSensorReadingsHistory200ResponseInnerDoor`

GetDoor returns the Door field if non-nil, zero value otherwise.

### GetDoorOk

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetDoorOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerDoor, bool)`

GetDoorOk returns a tuple with the Door field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoor

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) SetDoor(v GetOrganizationSensorReadingsHistory200ResponseInnerDoor)`

SetDoor sets Door field to given value.

### HasDoor

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) HasDoor() bool`

HasDoor returns a boolean if a field has been set.

### GetHumidity

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetHumidity() GetOrganizationSensorReadingsHistory200ResponseInnerHumidity`

GetHumidity returns the Humidity field if non-nil, zero value otherwise.

### GetHumidityOk

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetHumidityOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerHumidity, bool)`

GetHumidityOk returns a tuple with the Humidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumidity

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) SetHumidity(v GetOrganizationSensorReadingsHistory200ResponseInnerHumidity)`

SetHumidity sets Humidity field to given value.

### HasHumidity

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) HasHumidity() bool`

HasHumidity returns a boolean if a field has been set.

### GetIndoorAirQuality

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetIndoorAirQuality() GetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality`

GetIndoorAirQuality returns the IndoorAirQuality field if non-nil, zero value otherwise.

### GetIndoorAirQualityOk

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetIndoorAirQualityOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality, bool)`

GetIndoorAirQualityOk returns a tuple with the IndoorAirQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndoorAirQuality

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) SetIndoorAirQuality(v GetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality)`

SetIndoorAirQuality sets IndoorAirQuality field to given value.

### HasIndoorAirQuality

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) HasIndoorAirQuality() bool`

HasIndoorAirQuality returns a boolean if a field has been set.

### GetNoise

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetNoise() GetOrganizationSensorReadingsHistory200ResponseInnerNoise`

GetNoise returns the Noise field if non-nil, zero value otherwise.

### GetNoiseOk

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetNoiseOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerNoise, bool)`

GetNoiseOk returns a tuple with the Noise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoise

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) SetNoise(v GetOrganizationSensorReadingsHistory200ResponseInnerNoise)`

SetNoise sets Noise field to given value.

### HasNoise

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) HasNoise() bool`

HasNoise returns a boolean if a field has been set.

### GetPm25

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetPm25() GetOrganizationSensorReadingsHistory200ResponseInnerPm25`

GetPm25 returns the Pm25 field if non-nil, zero value otherwise.

### GetPm25Ok

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetPm25Ok() (*GetOrganizationSensorReadingsHistory200ResponseInnerPm25, bool)`

GetPm25Ok returns a tuple with the Pm25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPm25

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) SetPm25(v GetOrganizationSensorReadingsHistory200ResponseInnerPm25)`

SetPm25 sets Pm25 field to given value.

### HasPm25

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) HasPm25() bool`

HasPm25 returns a boolean if a field has been set.

### GetTemperature

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetTemperature() GetOrganizationSensorReadingsHistory200ResponseInnerTemperature`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetTemperatureOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerTemperature, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) SetTemperature(v GetOrganizationSensorReadingsHistory200ResponseInnerTemperature)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetTvoc

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetTvoc() GetOrganizationSensorReadingsHistory200ResponseInnerTvoc`

GetTvoc returns the Tvoc field if non-nil, zero value otherwise.

### GetTvocOk

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetTvocOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerTvoc, bool)`

GetTvocOk returns a tuple with the Tvoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTvoc

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) SetTvoc(v GetOrganizationSensorReadingsHistory200ResponseInnerTvoc)`

SetTvoc sets Tvoc field to given value.

### HasTvoc

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) HasTvoc() bool`

HasTvoc returns a boolean if a field has been set.

### GetWater

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetWater() GetOrganizationSensorReadingsHistory200ResponseInnerWater`

GetWater returns the Water field if non-nil, zero value otherwise.

### GetWaterOk

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetWaterOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerWater, bool)`

GetWaterOk returns a tuple with the Water field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWater

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) SetWater(v GetOrganizationSensorReadingsHistory200ResponseInnerWater)`

SetWater sets Water field to given value.

### HasWater

`func (o *GetOrganizationSensorReadingsHistory200ResponseInner) HasWater() bool`

HasWater returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


