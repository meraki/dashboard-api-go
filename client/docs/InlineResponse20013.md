# InlineResponse20013

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OccurredAt** | Pointer to **string** | time when the event occurred | [optional] 
**AlertTypeId** | Pointer to **string** | type of alert | [optional] 
**AlertType** | Pointer to **string** | user friendly alert type | [optional] 
**Device** | Pointer to [**NetworksNetworkIdAlertsHistoryDevice**](NetworksNetworkIdAlertsHistoryDevice.md) |  | [optional] 
**Destinations** | Pointer to [**NetworksNetworkIdAlertsHistoryDestinations**](NetworksNetworkIdAlertsHistoryDestinations.md) |  | [optional] 
**AlertData** | Pointer to **map[string]interface{}** | relevant data about the event that caused the alert | [optional] 

## Methods

### NewInlineResponse20013

`func NewInlineResponse20013() *InlineResponse20013`

NewInlineResponse20013 instantiates a new InlineResponse20013 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20013WithDefaults

`func NewInlineResponse20013WithDefaults() *InlineResponse20013`

NewInlineResponse20013WithDefaults instantiates a new InlineResponse20013 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOccurredAt

`func (o *InlineResponse20013) GetOccurredAt() string`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *InlineResponse20013) GetOccurredAtOk() (*string, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *InlineResponse20013) SetOccurredAt(v string)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *InlineResponse20013) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### GetAlertTypeId

`func (o *InlineResponse20013) GetAlertTypeId() string`

GetAlertTypeId returns the AlertTypeId field if non-nil, zero value otherwise.

### GetAlertTypeIdOk

`func (o *InlineResponse20013) GetAlertTypeIdOk() (*string, bool)`

GetAlertTypeIdOk returns a tuple with the AlertTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertTypeId

`func (o *InlineResponse20013) SetAlertTypeId(v string)`

SetAlertTypeId sets AlertTypeId field to given value.

### HasAlertTypeId

`func (o *InlineResponse20013) HasAlertTypeId() bool`

HasAlertTypeId returns a boolean if a field has been set.

### GetAlertType

`func (o *InlineResponse20013) GetAlertType() string`

GetAlertType returns the AlertType field if non-nil, zero value otherwise.

### GetAlertTypeOk

`func (o *InlineResponse20013) GetAlertTypeOk() (*string, bool)`

GetAlertTypeOk returns a tuple with the AlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertType

`func (o *InlineResponse20013) SetAlertType(v string)`

SetAlertType sets AlertType field to given value.

### HasAlertType

`func (o *InlineResponse20013) HasAlertType() bool`

HasAlertType returns a boolean if a field has been set.

### GetDevice

`func (o *InlineResponse20013) GetDevice() NetworksNetworkIdAlertsHistoryDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *InlineResponse20013) GetDeviceOk() (*NetworksNetworkIdAlertsHistoryDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *InlineResponse20013) SetDevice(v NetworksNetworkIdAlertsHistoryDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *InlineResponse20013) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetDestinations

`func (o *InlineResponse20013) GetDestinations() NetworksNetworkIdAlertsHistoryDestinations`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *InlineResponse20013) GetDestinationsOk() (*NetworksNetworkIdAlertsHistoryDestinations, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *InlineResponse20013) SetDestinations(v NetworksNetworkIdAlertsHistoryDestinations)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *InlineResponse20013) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.

### GetAlertData

`func (o *InlineResponse20013) GetAlertData() map[string]interface{}`

GetAlertData returns the AlertData field if non-nil, zero value otherwise.

### GetAlertDataOk

`func (o *InlineResponse20013) GetAlertDataOk() (*map[string]interface{}, bool)`

GetAlertDataOk returns a tuple with the AlertData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertData

`func (o *InlineResponse20013) SetAlertData(v map[string]interface{})`

SetAlertData sets AlertData field to given value.

### HasAlertData

`func (o *InlineResponse20013) HasAlertData() bool`

HasAlertData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


