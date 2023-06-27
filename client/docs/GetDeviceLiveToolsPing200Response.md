# GetDeviceLiveToolsPing200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PingId** | Pointer to **string** | Id to check the status of your ping request. | [optional] 
**Url** | Pointer to **string** | GET this url to check the status of your ping request. | [optional] 
**Request** | Pointer to [**CreateDeviceLiveToolsPing201ResponseRequest**](CreateDeviceLiveToolsPing201ResponseRequest.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the ping request. | [optional] 
**Results** | Pointer to [**GetDeviceLiveToolsPing200ResponseResults**](GetDeviceLiveToolsPing200ResponseResults.md) |  | [optional] 

## Methods

### NewGetDeviceLiveToolsPing200Response

`func NewGetDeviceLiveToolsPing200Response() *GetDeviceLiveToolsPing200Response`

NewGetDeviceLiveToolsPing200Response instantiates a new GetDeviceLiveToolsPing200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeviceLiveToolsPing200ResponseWithDefaults

`func NewGetDeviceLiveToolsPing200ResponseWithDefaults() *GetDeviceLiveToolsPing200Response`

NewGetDeviceLiveToolsPing200ResponseWithDefaults instantiates a new GetDeviceLiveToolsPing200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPingId

`func (o *GetDeviceLiveToolsPing200Response) GetPingId() string`

GetPingId returns the PingId field if non-nil, zero value otherwise.

### GetPingIdOk

`func (o *GetDeviceLiveToolsPing200Response) GetPingIdOk() (*string, bool)`

GetPingIdOk returns a tuple with the PingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingId

`func (o *GetDeviceLiveToolsPing200Response) SetPingId(v string)`

SetPingId sets PingId field to given value.

### HasPingId

`func (o *GetDeviceLiveToolsPing200Response) HasPingId() bool`

HasPingId returns a boolean if a field has been set.

### GetUrl

`func (o *GetDeviceLiveToolsPing200Response) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetDeviceLiveToolsPing200Response) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetDeviceLiveToolsPing200Response) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GetDeviceLiveToolsPing200Response) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRequest

`func (o *GetDeviceLiveToolsPing200Response) GetRequest() CreateDeviceLiveToolsPing201ResponseRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *GetDeviceLiveToolsPing200Response) GetRequestOk() (*CreateDeviceLiveToolsPing201ResponseRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *GetDeviceLiveToolsPing200Response) SetRequest(v CreateDeviceLiveToolsPing201ResponseRequest)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *GetDeviceLiveToolsPing200Response) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetStatus

`func (o *GetDeviceLiveToolsPing200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetDeviceLiveToolsPing200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetDeviceLiveToolsPing200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetDeviceLiveToolsPing200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResults

`func (o *GetDeviceLiveToolsPing200Response) GetResults() GetDeviceLiveToolsPing200ResponseResults`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetDeviceLiveToolsPing200Response) GetResultsOk() (*GetDeviceLiveToolsPing200ResponseResults, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetDeviceLiveToolsPing200Response) SetResults(v GetDeviceLiveToolsPing200ResponseResults)`

SetResults sets Results field to given value.

### HasResults

`func (o *GetDeviceLiveToolsPing200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


