# DevicesSerialLiveToolsCableTestPostRequestMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CableTestId** | Pointer to **string** | Id of the cable test request. Used to check the status of the request. | [optional] 
**Url** | Pointer to **string** | GET this url to check the status of your cable test request. | [optional] 
**Request** | Pointer to [**CreateDeviceLiveToolsCableTest201ResponseRequest**](CreateDeviceLiveToolsCableTest201ResponseRequest.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the cable test request. | [optional] 
**Results** | Pointer to [**[]DevicesSerialLiveToolsCableTestPostRequestMessageResultsInner**](DevicesSerialLiveToolsCableTestPostRequestMessageResultsInner.md) | Results of the cable test request, one for each requested port. | [optional] 
**Error** | Pointer to **string** | An error message for a failed execution | [optional] 

## Methods

### NewDevicesSerialLiveToolsCableTestPostRequestMessage

`func NewDevicesSerialLiveToolsCableTestPostRequestMessage() *DevicesSerialLiveToolsCableTestPostRequestMessage`

NewDevicesSerialLiveToolsCableTestPostRequestMessage instantiates a new DevicesSerialLiveToolsCableTestPostRequestMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicesSerialLiveToolsCableTestPostRequestMessageWithDefaults

`func NewDevicesSerialLiveToolsCableTestPostRequestMessageWithDefaults() *DevicesSerialLiveToolsCableTestPostRequestMessage`

NewDevicesSerialLiveToolsCableTestPostRequestMessageWithDefaults instantiates a new DevicesSerialLiveToolsCableTestPostRequestMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCableTestId

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessage) GetCableTestId() string`

GetCableTestId returns the CableTestId field if non-nil, zero value otherwise.

### GetCableTestIdOk

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessage) GetCableTestIdOk() (*string, bool)`

GetCableTestIdOk returns a tuple with the CableTestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCableTestId

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessage) SetCableTestId(v string)`

SetCableTestId sets CableTestId field to given value.

### HasCableTestId

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessage) HasCableTestId() bool`

HasCableTestId returns a boolean if a field has been set.

### GetUrl

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessage) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessage) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessage) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessage) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRequest

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessage) GetRequest() CreateDeviceLiveToolsCableTest201ResponseRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessage) GetRequestOk() (*CreateDeviceLiveToolsCableTest201ResponseRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessage) SetRequest(v CreateDeviceLiveToolsCableTest201ResponseRequest)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessage) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetStatus

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessage) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessage) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessage) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessage) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResults

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessage) GetResults() []DevicesSerialLiveToolsCableTestPostRequestMessageResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessage) GetResultsOk() (*[]DevicesSerialLiveToolsCableTestPostRequestMessageResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessage) SetResults(v []DevicesSerialLiveToolsCableTestPostRequestMessageResultsInner)`

SetResults sets Results field to given value.

### HasResults

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessage) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetError

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessage) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessage) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessage) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessage) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


