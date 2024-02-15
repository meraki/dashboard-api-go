# DevicesSerialLiveToolsArpTablePostRequestMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArpTableId** | Pointer to **string** | Id of the ARP table request. Used to check the status of the request. | [optional] 
**Url** | Pointer to **string** | GET this url to check the status of your ARP table request. | [optional] 
**Request** | Pointer to [**CreateDeviceLiveToolsArpTable201ResponseRequest**](CreateDeviceLiveToolsArpTable201ResponseRequest.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the ARP table request. | [optional] 
**Entries** | Pointer to [**[]DevicesSerialLiveToolsArpTablePostRequestMessageEntriesInner**](DevicesSerialLiveToolsArpTablePostRequestMessageEntriesInner.md) | The ARP table entries | [optional] 
**Error** | Pointer to **string** | An error message for a failed execution | [optional] 

## Methods

### NewDevicesSerialLiveToolsArpTablePostRequestMessage

`func NewDevicesSerialLiveToolsArpTablePostRequestMessage() *DevicesSerialLiveToolsArpTablePostRequestMessage`

NewDevicesSerialLiveToolsArpTablePostRequestMessage instantiates a new DevicesSerialLiveToolsArpTablePostRequestMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicesSerialLiveToolsArpTablePostRequestMessageWithDefaults

`func NewDevicesSerialLiveToolsArpTablePostRequestMessageWithDefaults() *DevicesSerialLiveToolsArpTablePostRequestMessage`

NewDevicesSerialLiveToolsArpTablePostRequestMessageWithDefaults instantiates a new DevicesSerialLiveToolsArpTablePostRequestMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArpTableId

`func (o *DevicesSerialLiveToolsArpTablePostRequestMessage) GetArpTableId() string`

GetArpTableId returns the ArpTableId field if non-nil, zero value otherwise.

### GetArpTableIdOk

`func (o *DevicesSerialLiveToolsArpTablePostRequestMessage) GetArpTableIdOk() (*string, bool)`

GetArpTableIdOk returns a tuple with the ArpTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArpTableId

`func (o *DevicesSerialLiveToolsArpTablePostRequestMessage) SetArpTableId(v string)`

SetArpTableId sets ArpTableId field to given value.

### HasArpTableId

`func (o *DevicesSerialLiveToolsArpTablePostRequestMessage) HasArpTableId() bool`

HasArpTableId returns a boolean if a field has been set.

### GetUrl

`func (o *DevicesSerialLiveToolsArpTablePostRequestMessage) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DevicesSerialLiveToolsArpTablePostRequestMessage) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DevicesSerialLiveToolsArpTablePostRequestMessage) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *DevicesSerialLiveToolsArpTablePostRequestMessage) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRequest

`func (o *DevicesSerialLiveToolsArpTablePostRequestMessage) GetRequest() CreateDeviceLiveToolsArpTable201ResponseRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *DevicesSerialLiveToolsArpTablePostRequestMessage) GetRequestOk() (*CreateDeviceLiveToolsArpTable201ResponseRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *DevicesSerialLiveToolsArpTablePostRequestMessage) SetRequest(v CreateDeviceLiveToolsArpTable201ResponseRequest)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *DevicesSerialLiveToolsArpTablePostRequestMessage) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetStatus

`func (o *DevicesSerialLiveToolsArpTablePostRequestMessage) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DevicesSerialLiveToolsArpTablePostRequestMessage) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DevicesSerialLiveToolsArpTablePostRequestMessage) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DevicesSerialLiveToolsArpTablePostRequestMessage) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEntries

`func (o *DevicesSerialLiveToolsArpTablePostRequestMessage) GetEntries() []DevicesSerialLiveToolsArpTablePostRequestMessageEntriesInner`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *DevicesSerialLiveToolsArpTablePostRequestMessage) GetEntriesOk() (*[]DevicesSerialLiveToolsArpTablePostRequestMessageEntriesInner, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *DevicesSerialLiveToolsArpTablePostRequestMessage) SetEntries(v []DevicesSerialLiveToolsArpTablePostRequestMessageEntriesInner)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *DevicesSerialLiveToolsArpTablePostRequestMessage) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetError

`func (o *DevicesSerialLiveToolsArpTablePostRequestMessage) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DevicesSerialLiveToolsArpTablePostRequestMessage) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DevicesSerialLiveToolsArpTablePostRequestMessage) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *DevicesSerialLiveToolsArpTablePostRequestMessage) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


