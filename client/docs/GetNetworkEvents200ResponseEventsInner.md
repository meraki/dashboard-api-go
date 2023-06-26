# GetNetworkEvents200ResponseEventsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OccurredAt** | Pointer to **string** | An UTC ISO8601 string of the time the event occurred at. | [optional] 
**NetworkId** | Pointer to **string** | The ID of the network. | [optional] 
**Type** | Pointer to **string** | The type of event being listed. | [optional] 
**Description** | Pointer to **string** | A description of the event the happened. | [optional] 
**Category** | Pointer to **string** | The category that the event type belongs to | [optional] 
**ClientId** | Pointer to **string** | A string identifying the client. This could be a client&#39;s MAC or IP address | [optional] 
**ClientDescription** | Pointer to **string** | A description of the client. This is usually the client&#39;s device name. | [optional] 
**ClientMac** | Pointer to **string** | The client&#39;s MAC address. | [optional] 
**DeviceSerial** | Pointer to **string** | The serial number of the device. Only shown if the device is an access point. | [optional] 
**DeviceName** | Pointer to **string** | The name of the device. Only shown if the device is an access point. | [optional] 
**SsidNumber** | Pointer to **int32** | The SSID number of the device. | [optional] 
**EventData** | Pointer to [**GetNetworkEvents200ResponseEventsInnerEventData**](GetNetworkEvents200ResponseEventsInnerEventData.md) |  | [optional] 

## Methods

### NewGetNetworkEvents200ResponseEventsInner

`func NewGetNetworkEvents200ResponseEventsInner() *GetNetworkEvents200ResponseEventsInner`

NewGetNetworkEvents200ResponseEventsInner instantiates a new GetNetworkEvents200ResponseEventsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkEvents200ResponseEventsInnerWithDefaults

`func NewGetNetworkEvents200ResponseEventsInnerWithDefaults() *GetNetworkEvents200ResponseEventsInner`

NewGetNetworkEvents200ResponseEventsInnerWithDefaults instantiates a new GetNetworkEvents200ResponseEventsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOccurredAt

`func (o *GetNetworkEvents200ResponseEventsInner) GetOccurredAt() string`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *GetNetworkEvents200ResponseEventsInner) GetOccurredAtOk() (*string, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *GetNetworkEvents200ResponseEventsInner) SetOccurredAt(v string)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *GetNetworkEvents200ResponseEventsInner) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### GetNetworkId

`func (o *GetNetworkEvents200ResponseEventsInner) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *GetNetworkEvents200ResponseEventsInner) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *GetNetworkEvents200ResponseEventsInner) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *GetNetworkEvents200ResponseEventsInner) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetType

`func (o *GetNetworkEvents200ResponseEventsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetNetworkEvents200ResponseEventsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetNetworkEvents200ResponseEventsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetNetworkEvents200ResponseEventsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *GetNetworkEvents200ResponseEventsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetNetworkEvents200ResponseEventsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetNetworkEvents200ResponseEventsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetNetworkEvents200ResponseEventsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategory

`func (o *GetNetworkEvents200ResponseEventsInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetNetworkEvents200ResponseEventsInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetNetworkEvents200ResponseEventsInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetNetworkEvents200ResponseEventsInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetClientId

`func (o *GetNetworkEvents200ResponseEventsInner) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *GetNetworkEvents200ResponseEventsInner) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *GetNetworkEvents200ResponseEventsInner) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *GetNetworkEvents200ResponseEventsInner) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientDescription

`func (o *GetNetworkEvents200ResponseEventsInner) GetClientDescription() string`

GetClientDescription returns the ClientDescription field if non-nil, zero value otherwise.

### GetClientDescriptionOk

`func (o *GetNetworkEvents200ResponseEventsInner) GetClientDescriptionOk() (*string, bool)`

GetClientDescriptionOk returns a tuple with the ClientDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientDescription

`func (o *GetNetworkEvents200ResponseEventsInner) SetClientDescription(v string)`

SetClientDescription sets ClientDescription field to given value.

### HasClientDescription

`func (o *GetNetworkEvents200ResponseEventsInner) HasClientDescription() bool`

HasClientDescription returns a boolean if a field has been set.

### GetClientMac

`func (o *GetNetworkEvents200ResponseEventsInner) GetClientMac() string`

GetClientMac returns the ClientMac field if non-nil, zero value otherwise.

### GetClientMacOk

`func (o *GetNetworkEvents200ResponseEventsInner) GetClientMacOk() (*string, bool)`

GetClientMacOk returns a tuple with the ClientMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMac

`func (o *GetNetworkEvents200ResponseEventsInner) SetClientMac(v string)`

SetClientMac sets ClientMac field to given value.

### HasClientMac

`func (o *GetNetworkEvents200ResponseEventsInner) HasClientMac() bool`

HasClientMac returns a boolean if a field has been set.

### GetDeviceSerial

`func (o *GetNetworkEvents200ResponseEventsInner) GetDeviceSerial() string`

GetDeviceSerial returns the DeviceSerial field if non-nil, zero value otherwise.

### GetDeviceSerialOk

`func (o *GetNetworkEvents200ResponseEventsInner) GetDeviceSerialOk() (*string, bool)`

GetDeviceSerialOk returns a tuple with the DeviceSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSerial

`func (o *GetNetworkEvents200ResponseEventsInner) SetDeviceSerial(v string)`

SetDeviceSerial sets DeviceSerial field to given value.

### HasDeviceSerial

`func (o *GetNetworkEvents200ResponseEventsInner) HasDeviceSerial() bool`

HasDeviceSerial returns a boolean if a field has been set.

### GetDeviceName

`func (o *GetNetworkEvents200ResponseEventsInner) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *GetNetworkEvents200ResponseEventsInner) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *GetNetworkEvents200ResponseEventsInner) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *GetNetworkEvents200ResponseEventsInner) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetSsidNumber

`func (o *GetNetworkEvents200ResponseEventsInner) GetSsidNumber() int32`

GetSsidNumber returns the SsidNumber field if non-nil, zero value otherwise.

### GetSsidNumberOk

`func (o *GetNetworkEvents200ResponseEventsInner) GetSsidNumberOk() (*int32, bool)`

GetSsidNumberOk returns a tuple with the SsidNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidNumber

`func (o *GetNetworkEvents200ResponseEventsInner) SetSsidNumber(v int32)`

SetSsidNumber sets SsidNumber field to given value.

### HasSsidNumber

`func (o *GetNetworkEvents200ResponseEventsInner) HasSsidNumber() bool`

HasSsidNumber returns a boolean if a field has been set.

### GetEventData

`func (o *GetNetworkEvents200ResponseEventsInner) GetEventData() GetNetworkEvents200ResponseEventsInnerEventData`

GetEventData returns the EventData field if non-nil, zero value otherwise.

### GetEventDataOk

`func (o *GetNetworkEvents200ResponseEventsInner) GetEventDataOk() (*GetNetworkEvents200ResponseEventsInnerEventData, bool)`

GetEventDataOk returns a tuple with the EventData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventData

`func (o *GetNetworkEvents200ResponseEventsInner) SetEventData(v GetNetworkEvents200ResponseEventsInnerEventData)`

SetEventData sets EventData field to given value.

### HasEventData

`func (o *GetNetworkEvents200ResponseEventsInner) HasEventData() bool`

HasEventData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


