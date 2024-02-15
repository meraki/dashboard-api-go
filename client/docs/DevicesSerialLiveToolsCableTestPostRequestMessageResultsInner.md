# DevicesSerialLiveToolsCableTestPostRequestMessageResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **string** | The port for which the test was performed. | [optional] 
**Status** | Pointer to **string** | The current status of the port. If the cable test is still being performed on the port, \&quot;in-progress\&quot; is used. If an error occurred during the cable test, \&quot;error\&quot; is used and the error property will be populated. | [optional] 
**SpeedMbps** | Pointer to **int32** | Speed in Mbps.  A speed of 0 indicates the port is down or the port speed is automatic. | [optional] 
**Error** | Pointer to **string** | If an error occurred during the cable test, the error message will be populated here. | [optional] 
**Pairs** | Pointer to [**[]DevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner**](DevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner.md) | Results for each twisted pair within the cable. | [optional] 

## Methods

### NewDevicesSerialLiveToolsCableTestPostRequestMessageResultsInner

`func NewDevicesSerialLiveToolsCableTestPostRequestMessageResultsInner() *DevicesSerialLiveToolsCableTestPostRequestMessageResultsInner`

NewDevicesSerialLiveToolsCableTestPostRequestMessageResultsInner instantiates a new DevicesSerialLiveToolsCableTestPostRequestMessageResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerWithDefaults

`func NewDevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerWithDefaults() *DevicesSerialLiveToolsCableTestPostRequestMessageResultsInner`

NewDevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerWithDefaults instantiates a new DevicesSerialLiveToolsCableTestPostRequestMessageResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessageResultsInner) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessageResultsInner) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessageResultsInner) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessageResultsInner) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetStatus

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessageResultsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessageResultsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessageResultsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessageResultsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpeedMbps

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessageResultsInner) GetSpeedMbps() int32`

GetSpeedMbps returns the SpeedMbps field if non-nil, zero value otherwise.

### GetSpeedMbpsOk

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessageResultsInner) GetSpeedMbpsOk() (*int32, bool)`

GetSpeedMbpsOk returns a tuple with the SpeedMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedMbps

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessageResultsInner) SetSpeedMbps(v int32)`

SetSpeedMbps sets SpeedMbps field to given value.

### HasSpeedMbps

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessageResultsInner) HasSpeedMbps() bool`

HasSpeedMbps returns a boolean if a field has been set.

### GetError

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessageResultsInner) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessageResultsInner) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessageResultsInner) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessageResultsInner) HasError() bool`

HasError returns a boolean if a field has been set.

### GetPairs

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessageResultsInner) GetPairs() []DevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner`

GetPairs returns the Pairs field if non-nil, zero value otherwise.

### GetPairsOk

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessageResultsInner) GetPairsOk() (*[]DevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner, bool)`

GetPairsOk returns a tuple with the Pairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairs

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessageResultsInner) SetPairs(v []DevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner)`

SetPairs sets Pairs field to given value.

### HasPairs

`func (o *DevicesSerialLiveToolsCableTestPostRequestMessageResultsInner) HasPairs() bool`

HasPairs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


