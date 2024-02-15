# GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | Pointer to **string** | Time at which the reading occurred, in ISO8601 format. | [optional] 
**Metric** | Pointer to **string** | Type of sensor reading. | [optional] 
**ApparentPower** | Pointer to [**GetOrganizationSensorReadingsHistory200ResponseInnerApparentPower**](GetOrganizationSensorReadingsHistory200ResponseInnerApparentPower.md) |  | [optional] 
**Battery** | Pointer to [**GetOrganizationSensorReadingsHistory200ResponseInnerBattery**](GetOrganizationSensorReadingsHistory200ResponseInnerBattery.md) |  | [optional] 
**Button** | Pointer to [**GetOrganizationSensorReadingsHistory200ResponseInnerButton**](GetOrganizationSensorReadingsHistory200ResponseInnerButton.md) |  | [optional] 
**Co2** | Pointer to [**GetOrganizationSensorReadingsHistory200ResponseInnerCo2**](GetOrganizationSensorReadingsHistory200ResponseInnerCo2.md) |  | [optional] 
**Current** | Pointer to [**GetOrganizationSensorReadingsHistory200ResponseInnerCurrent**](GetOrganizationSensorReadingsHistory200ResponseInnerCurrent.md) |  | [optional] 
**Door** | Pointer to [**GetOrganizationSensorReadingsHistory200ResponseInnerDoor**](GetOrganizationSensorReadingsHistory200ResponseInnerDoor.md) |  | [optional] 
**DownstreamPower** | Pointer to [**GetOrganizationSensorReadingsHistory200ResponseInnerDownstreamPower**](GetOrganizationSensorReadingsHistory200ResponseInnerDownstreamPower.md) |  | [optional] 
**Frequency** | Pointer to [**GetOrganizationSensorReadingsHistory200ResponseInnerFrequency**](GetOrganizationSensorReadingsHistory200ResponseInnerFrequency.md) |  | [optional] 
**Humidity** | Pointer to [**GetOrganizationSensorReadingsHistory200ResponseInnerHumidity**](GetOrganizationSensorReadingsHistory200ResponseInnerHumidity.md) |  | [optional] 
**IndoorAirQuality** | Pointer to [**GetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality**](GetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality.md) |  | [optional] 
**Noise** | Pointer to [**GetOrganizationSensorReadingsHistory200ResponseInnerNoise**](GetOrganizationSensorReadingsHistory200ResponseInnerNoise.md) |  | [optional] 
**Pm25** | Pointer to [**GetOrganizationSensorReadingsHistory200ResponseInnerPm25**](GetOrganizationSensorReadingsHistory200ResponseInnerPm25.md) |  | [optional] 
**PowerFactor** | Pointer to [**GetOrganizationSensorReadingsHistory200ResponseInnerPowerFactor**](GetOrganizationSensorReadingsHistory200ResponseInnerPowerFactor.md) |  | [optional] 
**RealPower** | Pointer to [**GetOrganizationSensorReadingsHistory200ResponseInnerRealPower**](GetOrganizationSensorReadingsHistory200ResponseInnerRealPower.md) |  | [optional] 
**RemoteLockoutSwitch** | Pointer to [**GetOrganizationSensorReadingsHistory200ResponseInnerRemoteLockoutSwitch**](GetOrganizationSensorReadingsHistory200ResponseInnerRemoteLockoutSwitch.md) |  | [optional] 
**Temperature** | Pointer to [**GetOrganizationSensorReadingsHistory200ResponseInnerTemperature**](GetOrganizationSensorReadingsHistory200ResponseInnerTemperature.md) |  | [optional] 
**Tvoc** | Pointer to [**GetOrganizationSensorReadingsHistory200ResponseInnerTvoc**](GetOrganizationSensorReadingsHistory200ResponseInnerTvoc.md) |  | [optional] 
**Voltage** | Pointer to [**GetOrganizationSensorReadingsHistory200ResponseInnerVoltage**](GetOrganizationSensorReadingsHistory200ResponseInnerVoltage.md) |  | [optional] 
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

### GetApparentPower

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetApparentPower() GetOrganizationSensorReadingsHistory200ResponseInnerApparentPower`

GetApparentPower returns the ApparentPower field if non-nil, zero value otherwise.

### GetApparentPowerOk

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetApparentPowerOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerApparentPower, bool)`

GetApparentPowerOk returns a tuple with the ApparentPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApparentPower

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) SetApparentPower(v GetOrganizationSensorReadingsHistory200ResponseInnerApparentPower)`

SetApparentPower sets ApparentPower field to given value.

### HasApparentPower

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) HasApparentPower() bool`

HasApparentPower returns a boolean if a field has been set.

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

### GetCo2

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetCo2() GetOrganizationSensorReadingsHistory200ResponseInnerCo2`

GetCo2 returns the Co2 field if non-nil, zero value otherwise.

### GetCo2Ok

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetCo2Ok() (*GetOrganizationSensorReadingsHistory200ResponseInnerCo2, bool)`

GetCo2Ok returns a tuple with the Co2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCo2

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) SetCo2(v GetOrganizationSensorReadingsHistory200ResponseInnerCo2)`

SetCo2 sets Co2 field to given value.

### HasCo2

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) HasCo2() bool`

HasCo2 returns a boolean if a field has been set.

### GetCurrent

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetCurrent() GetOrganizationSensorReadingsHistory200ResponseInnerCurrent`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetCurrentOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerCurrent, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) SetCurrent(v GetOrganizationSensorReadingsHistory200ResponseInnerCurrent)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

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

### GetDownstreamPower

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetDownstreamPower() GetOrganizationSensorReadingsHistory200ResponseInnerDownstreamPower`

GetDownstreamPower returns the DownstreamPower field if non-nil, zero value otherwise.

### GetDownstreamPowerOk

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetDownstreamPowerOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerDownstreamPower, bool)`

GetDownstreamPowerOk returns a tuple with the DownstreamPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownstreamPower

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) SetDownstreamPower(v GetOrganizationSensorReadingsHistory200ResponseInnerDownstreamPower)`

SetDownstreamPower sets DownstreamPower field to given value.

### HasDownstreamPower

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) HasDownstreamPower() bool`

HasDownstreamPower returns a boolean if a field has been set.

### GetFrequency

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetFrequency() GetOrganizationSensorReadingsHistory200ResponseInnerFrequency`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetFrequencyOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerFrequency, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) SetFrequency(v GetOrganizationSensorReadingsHistory200ResponseInnerFrequency)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

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

### GetPowerFactor

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetPowerFactor() GetOrganizationSensorReadingsHistory200ResponseInnerPowerFactor`

GetPowerFactor returns the PowerFactor field if non-nil, zero value otherwise.

### GetPowerFactorOk

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetPowerFactorOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerPowerFactor, bool)`

GetPowerFactorOk returns a tuple with the PowerFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerFactor

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) SetPowerFactor(v GetOrganizationSensorReadingsHistory200ResponseInnerPowerFactor)`

SetPowerFactor sets PowerFactor field to given value.

### HasPowerFactor

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) HasPowerFactor() bool`

HasPowerFactor returns a boolean if a field has been set.

### GetRealPower

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetRealPower() GetOrganizationSensorReadingsHistory200ResponseInnerRealPower`

GetRealPower returns the RealPower field if non-nil, zero value otherwise.

### GetRealPowerOk

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetRealPowerOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerRealPower, bool)`

GetRealPowerOk returns a tuple with the RealPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealPower

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) SetRealPower(v GetOrganizationSensorReadingsHistory200ResponseInnerRealPower)`

SetRealPower sets RealPower field to given value.

### HasRealPower

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) HasRealPower() bool`

HasRealPower returns a boolean if a field has been set.

### GetRemoteLockoutSwitch

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetRemoteLockoutSwitch() GetOrganizationSensorReadingsHistory200ResponseInnerRemoteLockoutSwitch`

GetRemoteLockoutSwitch returns the RemoteLockoutSwitch field if non-nil, zero value otherwise.

### GetRemoteLockoutSwitchOk

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetRemoteLockoutSwitchOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerRemoteLockoutSwitch, bool)`

GetRemoteLockoutSwitchOk returns a tuple with the RemoteLockoutSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteLockoutSwitch

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) SetRemoteLockoutSwitch(v GetOrganizationSensorReadingsHistory200ResponseInnerRemoteLockoutSwitch)`

SetRemoteLockoutSwitch sets RemoteLockoutSwitch field to given value.

### HasRemoteLockoutSwitch

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) HasRemoteLockoutSwitch() bool`

HasRemoteLockoutSwitch returns a boolean if a field has been set.

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

### GetVoltage

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetVoltage() GetOrganizationSensorReadingsHistory200ResponseInnerVoltage`

GetVoltage returns the Voltage field if non-nil, zero value otherwise.

### GetVoltageOk

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) GetVoltageOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerVoltage, bool)`

GetVoltageOk returns a tuple with the Voltage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoltage

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) SetVoltage(v GetOrganizationSensorReadingsHistory200ResponseInnerVoltage)`

SetVoltage sets Voltage field to given value.

### HasVoltage

`func (o *GetOrganizationSensorReadingsLatest200ResponseInnerReadingsInner) HasVoltage() bool`

HasVoltage returns a boolean if a field has been set.

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


