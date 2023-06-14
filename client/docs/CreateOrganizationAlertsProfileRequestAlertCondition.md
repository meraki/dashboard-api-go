# CreateOrganizationAlertsProfileRequestAlertCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **int32** | The total duration in seconds that the threshold should be crossed before alerting | [optional] 
**Window** | Pointer to **int32** | The look back period in seconds for sensing the alert | [optional] 
**BitRateBps** | Pointer to **int32** | The threshold the metric must cross to be valid for alerting. Used only for WAN Utilization alerts. | [optional] 
**LossRatio** | Pointer to **float32** | The threshold the metric must cross to be valid for alerting. Used only for Packet Loss alerts. | [optional] 
**LatencyMs** | Pointer to **int32** | The threshold the metric must cross to be valid for alerting. Used only for WAN Latency alerts. | [optional] 
**JitterMs** | Pointer to **int32** | The threshold the metric must cross to be valid for alerting. Used only for VoIP Jitter alerts. | [optional] 
**Mos** | Pointer to **float32** | The threshold the metric must drop below to be valid for alerting. Used only for VoIP MOS alerts. | [optional] 
**Interface** | Pointer to **string** | The uplink observed for the alert.  interface must be one of the following: wan1, wan2, cellular | [optional] 

## Methods

### NewCreateOrganizationAlertsProfileRequestAlertCondition

`func NewCreateOrganizationAlertsProfileRequestAlertCondition() *CreateOrganizationAlertsProfileRequestAlertCondition`

NewCreateOrganizationAlertsProfileRequestAlertCondition instantiates a new CreateOrganizationAlertsProfileRequestAlertCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationAlertsProfileRequestAlertConditionWithDefaults

`func NewCreateOrganizationAlertsProfileRequestAlertConditionWithDefaults() *CreateOrganizationAlertsProfileRequestAlertCondition`

NewCreateOrganizationAlertsProfileRequestAlertConditionWithDefaults instantiates a new CreateOrganizationAlertsProfileRequestAlertCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *CreateOrganizationAlertsProfileRequestAlertCondition) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CreateOrganizationAlertsProfileRequestAlertCondition) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CreateOrganizationAlertsProfileRequestAlertCondition) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *CreateOrganizationAlertsProfileRequestAlertCondition) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetWindow

`func (o *CreateOrganizationAlertsProfileRequestAlertCondition) GetWindow() int32`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *CreateOrganizationAlertsProfileRequestAlertCondition) GetWindowOk() (*int32, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *CreateOrganizationAlertsProfileRequestAlertCondition) SetWindow(v int32)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *CreateOrganizationAlertsProfileRequestAlertCondition) HasWindow() bool`

HasWindow returns a boolean if a field has been set.

### GetBitRateBps

`func (o *CreateOrganizationAlertsProfileRequestAlertCondition) GetBitRateBps() int32`

GetBitRateBps returns the BitRateBps field if non-nil, zero value otherwise.

### GetBitRateBpsOk

`func (o *CreateOrganizationAlertsProfileRequestAlertCondition) GetBitRateBpsOk() (*int32, bool)`

GetBitRateBpsOk returns a tuple with the BitRateBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitRateBps

`func (o *CreateOrganizationAlertsProfileRequestAlertCondition) SetBitRateBps(v int32)`

SetBitRateBps sets BitRateBps field to given value.

### HasBitRateBps

`func (o *CreateOrganizationAlertsProfileRequestAlertCondition) HasBitRateBps() bool`

HasBitRateBps returns a boolean if a field has been set.

### GetLossRatio

`func (o *CreateOrganizationAlertsProfileRequestAlertCondition) GetLossRatio() float32`

GetLossRatio returns the LossRatio field if non-nil, zero value otherwise.

### GetLossRatioOk

`func (o *CreateOrganizationAlertsProfileRequestAlertCondition) GetLossRatioOk() (*float32, bool)`

GetLossRatioOk returns a tuple with the LossRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLossRatio

`func (o *CreateOrganizationAlertsProfileRequestAlertCondition) SetLossRatio(v float32)`

SetLossRatio sets LossRatio field to given value.

### HasLossRatio

`func (o *CreateOrganizationAlertsProfileRequestAlertCondition) HasLossRatio() bool`

HasLossRatio returns a boolean if a field has been set.

### GetLatencyMs

`func (o *CreateOrganizationAlertsProfileRequestAlertCondition) GetLatencyMs() int32`

GetLatencyMs returns the LatencyMs field if non-nil, zero value otherwise.

### GetLatencyMsOk

`func (o *CreateOrganizationAlertsProfileRequestAlertCondition) GetLatencyMsOk() (*int32, bool)`

GetLatencyMsOk returns a tuple with the LatencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyMs

`func (o *CreateOrganizationAlertsProfileRequestAlertCondition) SetLatencyMs(v int32)`

SetLatencyMs sets LatencyMs field to given value.

### HasLatencyMs

`func (o *CreateOrganizationAlertsProfileRequestAlertCondition) HasLatencyMs() bool`

HasLatencyMs returns a boolean if a field has been set.

### GetJitterMs

`func (o *CreateOrganizationAlertsProfileRequestAlertCondition) GetJitterMs() int32`

GetJitterMs returns the JitterMs field if non-nil, zero value otherwise.

### GetJitterMsOk

`func (o *CreateOrganizationAlertsProfileRequestAlertCondition) GetJitterMsOk() (*int32, bool)`

GetJitterMsOk returns a tuple with the JitterMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitterMs

`func (o *CreateOrganizationAlertsProfileRequestAlertCondition) SetJitterMs(v int32)`

SetJitterMs sets JitterMs field to given value.

### HasJitterMs

`func (o *CreateOrganizationAlertsProfileRequestAlertCondition) HasJitterMs() bool`

HasJitterMs returns a boolean if a field has been set.

### GetMos

`func (o *CreateOrganizationAlertsProfileRequestAlertCondition) GetMos() float32`

GetMos returns the Mos field if non-nil, zero value otherwise.

### GetMosOk

`func (o *CreateOrganizationAlertsProfileRequestAlertCondition) GetMosOk() (*float32, bool)`

GetMosOk returns a tuple with the Mos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMos

`func (o *CreateOrganizationAlertsProfileRequestAlertCondition) SetMos(v float32)`

SetMos sets Mos field to given value.

### HasMos

`func (o *CreateOrganizationAlertsProfileRequestAlertCondition) HasMos() bool`

HasMos returns a boolean if a field has been set.

### GetInterface

`func (o *CreateOrganizationAlertsProfileRequestAlertCondition) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *CreateOrganizationAlertsProfileRequestAlertCondition) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *CreateOrganizationAlertsProfileRequestAlertCondition) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *CreateOrganizationAlertsProfileRequestAlertCondition) HasInterface() bool`

HasInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


