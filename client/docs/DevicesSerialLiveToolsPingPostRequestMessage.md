# DevicesSerialLiveToolsPingPostRequestMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PingId** | Pointer to **string** | Id to check the status of your ping request. | [optional] 
**Url** | Pointer to **string** | GET this url to check the status of your ping request. | [optional] 
**Request** | Pointer to [**CreateDeviceLiveToolsPing201ResponseRequest**](CreateDeviceLiveToolsPing201ResponseRequest.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the ping request. | [optional] 
**Results** | Pointer to [**DevicesSerialLiveToolsPingPostRequestMessageResults**](DevicesSerialLiveToolsPingPostRequestMessageResults.md) |  | [optional] 

## Methods

### NewDevicesSerialLiveToolsPingPostRequestMessage

`func NewDevicesSerialLiveToolsPingPostRequestMessage() *DevicesSerialLiveToolsPingPostRequestMessage`

NewDevicesSerialLiveToolsPingPostRequestMessage instantiates a new DevicesSerialLiveToolsPingPostRequestMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicesSerialLiveToolsPingPostRequestMessageWithDefaults

`func NewDevicesSerialLiveToolsPingPostRequestMessageWithDefaults() *DevicesSerialLiveToolsPingPostRequestMessage`

NewDevicesSerialLiveToolsPingPostRequestMessageWithDefaults instantiates a new DevicesSerialLiveToolsPingPostRequestMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPingId

`func (o *DevicesSerialLiveToolsPingPostRequestMessage) GetPingId() string`

GetPingId returns the PingId field if non-nil, zero value otherwise.

### GetPingIdOk

`func (o *DevicesSerialLiveToolsPingPostRequestMessage) GetPingIdOk() (*string, bool)`

GetPingIdOk returns a tuple with the PingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingId

`func (o *DevicesSerialLiveToolsPingPostRequestMessage) SetPingId(v string)`

SetPingId sets PingId field to given value.

### HasPingId

`func (o *DevicesSerialLiveToolsPingPostRequestMessage) HasPingId() bool`

HasPingId returns a boolean if a field has been set.

### GetUrl

`func (o *DevicesSerialLiveToolsPingPostRequestMessage) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DevicesSerialLiveToolsPingPostRequestMessage) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DevicesSerialLiveToolsPingPostRequestMessage) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *DevicesSerialLiveToolsPingPostRequestMessage) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRequest

`func (o *DevicesSerialLiveToolsPingPostRequestMessage) GetRequest() CreateDeviceLiveToolsPing201ResponseRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *DevicesSerialLiveToolsPingPostRequestMessage) GetRequestOk() (*CreateDeviceLiveToolsPing201ResponseRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *DevicesSerialLiveToolsPingPostRequestMessage) SetRequest(v CreateDeviceLiveToolsPing201ResponseRequest)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *DevicesSerialLiveToolsPingPostRequestMessage) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetStatus

`func (o *DevicesSerialLiveToolsPingPostRequestMessage) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DevicesSerialLiveToolsPingPostRequestMessage) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DevicesSerialLiveToolsPingPostRequestMessage) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DevicesSerialLiveToolsPingPostRequestMessage) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResults

`func (o *DevicesSerialLiveToolsPingPostRequestMessage) GetResults() DevicesSerialLiveToolsPingPostRequestMessageResults`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *DevicesSerialLiveToolsPingPostRequestMessage) GetResultsOk() (*DevicesSerialLiveToolsPingPostRequestMessageResults, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *DevicesSerialLiveToolsPingPostRequestMessage) SetResults(v DevicesSerialLiveToolsPingPostRequestMessageResults)`

SetResults sets Results field to given value.

### HasResults

`func (o *DevicesSerialLiveToolsPingPostRequestMessage) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


