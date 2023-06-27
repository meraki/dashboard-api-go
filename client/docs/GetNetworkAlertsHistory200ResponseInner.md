# GetNetworkAlertsHistory200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OccurredAt** | Pointer to **string** | time when the event occurred | [optional] 
**AlertTypeId** | Pointer to **string** | type of alert | [optional] 
**AlertType** | Pointer to **string** | user friendly alert type | [optional] 
**Device** | Pointer to [**GetNetworkAlertsHistory200ResponseInnerDevice**](GetNetworkAlertsHistory200ResponseInnerDevice.md) |  | [optional] 
**Destinations** | Pointer to [**GetNetworkAlertsHistory200ResponseInnerDestinations**](GetNetworkAlertsHistory200ResponseInnerDestinations.md) |  | [optional] 
**AlertData** | Pointer to **map[string]interface{}** | relevant data about the event that caused the alert | [optional] 

## Methods

### NewGetNetworkAlertsHistory200ResponseInner

`func NewGetNetworkAlertsHistory200ResponseInner() *GetNetworkAlertsHistory200ResponseInner`

NewGetNetworkAlertsHistory200ResponseInner instantiates a new GetNetworkAlertsHistory200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkAlertsHistory200ResponseInnerWithDefaults

`func NewGetNetworkAlertsHistory200ResponseInnerWithDefaults() *GetNetworkAlertsHistory200ResponseInner`

NewGetNetworkAlertsHistory200ResponseInnerWithDefaults instantiates a new GetNetworkAlertsHistory200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOccurredAt

`func (o *GetNetworkAlertsHistory200ResponseInner) GetOccurredAt() string`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *GetNetworkAlertsHistory200ResponseInner) GetOccurredAtOk() (*string, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *GetNetworkAlertsHistory200ResponseInner) SetOccurredAt(v string)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *GetNetworkAlertsHistory200ResponseInner) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### GetAlertTypeId

`func (o *GetNetworkAlertsHistory200ResponseInner) GetAlertTypeId() string`

GetAlertTypeId returns the AlertTypeId field if non-nil, zero value otherwise.

### GetAlertTypeIdOk

`func (o *GetNetworkAlertsHistory200ResponseInner) GetAlertTypeIdOk() (*string, bool)`

GetAlertTypeIdOk returns a tuple with the AlertTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertTypeId

`func (o *GetNetworkAlertsHistory200ResponseInner) SetAlertTypeId(v string)`

SetAlertTypeId sets AlertTypeId field to given value.

### HasAlertTypeId

`func (o *GetNetworkAlertsHistory200ResponseInner) HasAlertTypeId() bool`

HasAlertTypeId returns a boolean if a field has been set.

### GetAlertType

`func (o *GetNetworkAlertsHistory200ResponseInner) GetAlertType() string`

GetAlertType returns the AlertType field if non-nil, zero value otherwise.

### GetAlertTypeOk

`func (o *GetNetworkAlertsHistory200ResponseInner) GetAlertTypeOk() (*string, bool)`

GetAlertTypeOk returns a tuple with the AlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertType

`func (o *GetNetworkAlertsHistory200ResponseInner) SetAlertType(v string)`

SetAlertType sets AlertType field to given value.

### HasAlertType

`func (o *GetNetworkAlertsHistory200ResponseInner) HasAlertType() bool`

HasAlertType returns a boolean if a field has been set.

### GetDevice

`func (o *GetNetworkAlertsHistory200ResponseInner) GetDevice() GetNetworkAlertsHistory200ResponseInnerDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *GetNetworkAlertsHistory200ResponseInner) GetDeviceOk() (*GetNetworkAlertsHistory200ResponseInnerDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *GetNetworkAlertsHistory200ResponseInner) SetDevice(v GetNetworkAlertsHistory200ResponseInnerDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *GetNetworkAlertsHistory200ResponseInner) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetDestinations

`func (o *GetNetworkAlertsHistory200ResponseInner) GetDestinations() GetNetworkAlertsHistory200ResponseInnerDestinations`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *GetNetworkAlertsHistory200ResponseInner) GetDestinationsOk() (*GetNetworkAlertsHistory200ResponseInnerDestinations, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *GetNetworkAlertsHistory200ResponseInner) SetDestinations(v GetNetworkAlertsHistory200ResponseInnerDestinations)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *GetNetworkAlertsHistory200ResponseInner) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.

### GetAlertData

`func (o *GetNetworkAlertsHistory200ResponseInner) GetAlertData() map[string]interface{}`

GetAlertData returns the AlertData field if non-nil, zero value otherwise.

### GetAlertDataOk

`func (o *GetNetworkAlertsHistory200ResponseInner) GetAlertDataOk() (*map[string]interface{}, bool)`

GetAlertDataOk returns a tuple with the AlertData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertData

`func (o *GetNetworkAlertsHistory200ResponseInner) SetAlertData(v map[string]interface{})`

SetAlertData sets AlertData field to given value.

### HasAlertData

`func (o *GetNetworkAlertsHistory200ResponseInner) HasAlertData() bool`

HasAlertData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


