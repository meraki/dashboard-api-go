# DevicesSerialLiveToolsWakeOnLanPostRequestMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WakeOnLanId** | Pointer to **string** | ID of the Wake-on-LAN job | [optional] 
**Url** | Pointer to **string** | GET this url to check the status of your ping request | [optional] 
**Status** | Pointer to **string** | Status of the Wake-on-LAN request | [optional] 
**Request** | Pointer to [**CreateDeviceLiveToolsWakeOnLan201ResponseRequest**](CreateDeviceLiveToolsWakeOnLan201ResponseRequest.md) |  | [optional] 
**Error** | Pointer to **string** | An error message for a failed execution | [optional] 

## Methods

### NewDevicesSerialLiveToolsWakeOnLanPostRequestMessage

`func NewDevicesSerialLiveToolsWakeOnLanPostRequestMessage() *DevicesSerialLiveToolsWakeOnLanPostRequestMessage`

NewDevicesSerialLiveToolsWakeOnLanPostRequestMessage instantiates a new DevicesSerialLiveToolsWakeOnLanPostRequestMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicesSerialLiveToolsWakeOnLanPostRequestMessageWithDefaults

`func NewDevicesSerialLiveToolsWakeOnLanPostRequestMessageWithDefaults() *DevicesSerialLiveToolsWakeOnLanPostRequestMessage`

NewDevicesSerialLiveToolsWakeOnLanPostRequestMessageWithDefaults instantiates a new DevicesSerialLiveToolsWakeOnLanPostRequestMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWakeOnLanId

`func (o *DevicesSerialLiveToolsWakeOnLanPostRequestMessage) GetWakeOnLanId() string`

GetWakeOnLanId returns the WakeOnLanId field if non-nil, zero value otherwise.

### GetWakeOnLanIdOk

`func (o *DevicesSerialLiveToolsWakeOnLanPostRequestMessage) GetWakeOnLanIdOk() (*string, bool)`

GetWakeOnLanIdOk returns a tuple with the WakeOnLanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWakeOnLanId

`func (o *DevicesSerialLiveToolsWakeOnLanPostRequestMessage) SetWakeOnLanId(v string)`

SetWakeOnLanId sets WakeOnLanId field to given value.

### HasWakeOnLanId

`func (o *DevicesSerialLiveToolsWakeOnLanPostRequestMessage) HasWakeOnLanId() bool`

HasWakeOnLanId returns a boolean if a field has been set.

### GetUrl

`func (o *DevicesSerialLiveToolsWakeOnLanPostRequestMessage) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DevicesSerialLiveToolsWakeOnLanPostRequestMessage) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DevicesSerialLiveToolsWakeOnLanPostRequestMessage) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *DevicesSerialLiveToolsWakeOnLanPostRequestMessage) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetStatus

`func (o *DevicesSerialLiveToolsWakeOnLanPostRequestMessage) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DevicesSerialLiveToolsWakeOnLanPostRequestMessage) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DevicesSerialLiveToolsWakeOnLanPostRequestMessage) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DevicesSerialLiveToolsWakeOnLanPostRequestMessage) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRequest

`func (o *DevicesSerialLiveToolsWakeOnLanPostRequestMessage) GetRequest() CreateDeviceLiveToolsWakeOnLan201ResponseRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *DevicesSerialLiveToolsWakeOnLanPostRequestMessage) GetRequestOk() (*CreateDeviceLiveToolsWakeOnLan201ResponseRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *DevicesSerialLiveToolsWakeOnLanPostRequestMessage) SetRequest(v CreateDeviceLiveToolsWakeOnLan201ResponseRequest)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *DevicesSerialLiveToolsWakeOnLanPostRequestMessage) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetError

`func (o *DevicesSerialLiveToolsWakeOnLanPostRequestMessage) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DevicesSerialLiveToolsWakeOnLanPostRequestMessage) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DevicesSerialLiveToolsWakeOnLanPostRequestMessage) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *DevicesSerialLiveToolsWakeOnLanPostRequestMessage) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


