# InlineResponse20022Events

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OccurredAt** | Pointer to **string** | An UTC ISO8601 string of the time the event occurred at. | [optional] 
**NetworkId** | Pointer to **string** | The ID of the network. | [optional] 
**Type** | Pointer to **string** | The type of event being listed. | [optional] 
**Description** | Pointer to **string** | A description of the event the happened. | [optional] 
**ClientId** | Pointer to **string** | A string identifying the client. This could be a client&#39;s MAC or IP address | [optional] 
**ClientDescription** | Pointer to **string** | A description of the client. This is usually the client&#39;s device name. | [optional] 
**ClientMac** | Pointer to **string** | The client&#39;s MAC address. | [optional] 
**DeviceSerial** | Pointer to **string** | The serial number of the device. Only shown if the device is an access point. | [optional] 
**DeviceName** | Pointer to **string** | The name of the device. Only shown if the device is an access point. | [optional] 
**SsidNumber** | Pointer to **int32** | The SSID number of the device. | [optional] 
**EventData** | Pointer to [**InlineResponse20022EventData**](InlineResponse20022EventData.md) |  | [optional] 

## Methods

### NewInlineResponse20022Events

`func NewInlineResponse20022Events() *InlineResponse20022Events`

NewInlineResponse20022Events instantiates a new InlineResponse20022Events object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20022EventsWithDefaults

`func NewInlineResponse20022EventsWithDefaults() *InlineResponse20022Events`

NewInlineResponse20022EventsWithDefaults instantiates a new InlineResponse20022Events object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOccurredAt

`func (o *InlineResponse20022Events) GetOccurredAt() string`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *InlineResponse20022Events) GetOccurredAtOk() (*string, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *InlineResponse20022Events) SetOccurredAt(v string)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *InlineResponse20022Events) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### GetNetworkId

`func (o *InlineResponse20022Events) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse20022Events) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse20022Events) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse20022Events) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse20022Events) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse20022Events) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse20022Events) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse20022Events) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse20022Events) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse20022Events) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse20022Events) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse20022Events) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetClientId

`func (o *InlineResponse20022Events) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *InlineResponse20022Events) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *InlineResponse20022Events) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *InlineResponse20022Events) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientDescription

`func (o *InlineResponse20022Events) GetClientDescription() string`

GetClientDescription returns the ClientDescription field if non-nil, zero value otherwise.

### GetClientDescriptionOk

`func (o *InlineResponse20022Events) GetClientDescriptionOk() (*string, bool)`

GetClientDescriptionOk returns a tuple with the ClientDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientDescription

`func (o *InlineResponse20022Events) SetClientDescription(v string)`

SetClientDescription sets ClientDescription field to given value.

### HasClientDescription

`func (o *InlineResponse20022Events) HasClientDescription() bool`

HasClientDescription returns a boolean if a field has been set.

### GetClientMac

`func (o *InlineResponse20022Events) GetClientMac() string`

GetClientMac returns the ClientMac field if non-nil, zero value otherwise.

### GetClientMacOk

`func (o *InlineResponse20022Events) GetClientMacOk() (*string, bool)`

GetClientMacOk returns a tuple with the ClientMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMac

`func (o *InlineResponse20022Events) SetClientMac(v string)`

SetClientMac sets ClientMac field to given value.

### HasClientMac

`func (o *InlineResponse20022Events) HasClientMac() bool`

HasClientMac returns a boolean if a field has been set.

### GetDeviceSerial

`func (o *InlineResponse20022Events) GetDeviceSerial() string`

GetDeviceSerial returns the DeviceSerial field if non-nil, zero value otherwise.

### GetDeviceSerialOk

`func (o *InlineResponse20022Events) GetDeviceSerialOk() (*string, bool)`

GetDeviceSerialOk returns a tuple with the DeviceSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSerial

`func (o *InlineResponse20022Events) SetDeviceSerial(v string)`

SetDeviceSerial sets DeviceSerial field to given value.

### HasDeviceSerial

`func (o *InlineResponse20022Events) HasDeviceSerial() bool`

HasDeviceSerial returns a boolean if a field has been set.

### GetDeviceName

`func (o *InlineResponse20022Events) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *InlineResponse20022Events) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *InlineResponse20022Events) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *InlineResponse20022Events) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetSsidNumber

`func (o *InlineResponse20022Events) GetSsidNumber() int32`

GetSsidNumber returns the SsidNumber field if non-nil, zero value otherwise.

### GetSsidNumberOk

`func (o *InlineResponse20022Events) GetSsidNumberOk() (*int32, bool)`

GetSsidNumberOk returns a tuple with the SsidNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidNumber

`func (o *InlineResponse20022Events) SetSsidNumber(v int32)`

SetSsidNumber sets SsidNumber field to given value.

### HasSsidNumber

`func (o *InlineResponse20022Events) HasSsidNumber() bool`

HasSsidNumber returns a boolean if a field has been set.

### GetEventData

`func (o *InlineResponse20022Events) GetEventData() InlineResponse20022EventData`

GetEventData returns the EventData field if non-nil, zero value otherwise.

### GetEventDataOk

`func (o *InlineResponse20022Events) GetEventDataOk() (*InlineResponse20022EventData, bool)`

GetEventDataOk returns a tuple with the EventData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventData

`func (o *InlineResponse20022Events) SetEventData(v InlineResponse20022EventData)`

SetEventData sets EventData field to given value.

### HasEventData

`func (o *InlineResponse20022Events) HasEventData() bool`

HasEventData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


