# GetDeviceCameraVideoSettings200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalRtspEnabled** | Pointer to **bool** | Boolean indicating if external rtsp stream is exposed | [optional] 
**RtspUrl** | Pointer to **string** | External rstp url. Will only be returned if external rtsp stream is exposed | [optional] 

## Methods

### NewGetDeviceCameraVideoSettings200Response

`func NewGetDeviceCameraVideoSettings200Response() *GetDeviceCameraVideoSettings200Response`

NewGetDeviceCameraVideoSettings200Response instantiates a new GetDeviceCameraVideoSettings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeviceCameraVideoSettings200ResponseWithDefaults

`func NewGetDeviceCameraVideoSettings200ResponseWithDefaults() *GetDeviceCameraVideoSettings200Response`

NewGetDeviceCameraVideoSettings200ResponseWithDefaults instantiates a new GetDeviceCameraVideoSettings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalRtspEnabled

`func (o *GetDeviceCameraVideoSettings200Response) GetExternalRtspEnabled() bool`

GetExternalRtspEnabled returns the ExternalRtspEnabled field if non-nil, zero value otherwise.

### GetExternalRtspEnabledOk

`func (o *GetDeviceCameraVideoSettings200Response) GetExternalRtspEnabledOk() (*bool, bool)`

GetExternalRtspEnabledOk returns a tuple with the ExternalRtspEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRtspEnabled

`func (o *GetDeviceCameraVideoSettings200Response) SetExternalRtspEnabled(v bool)`

SetExternalRtspEnabled sets ExternalRtspEnabled field to given value.

### HasExternalRtspEnabled

`func (o *GetDeviceCameraVideoSettings200Response) HasExternalRtspEnabled() bool`

HasExternalRtspEnabled returns a boolean if a field has been set.

### GetRtspUrl

`func (o *GetDeviceCameraVideoSettings200Response) GetRtspUrl() string`

GetRtspUrl returns the RtspUrl field if non-nil, zero value otherwise.

### GetRtspUrlOk

`func (o *GetDeviceCameraVideoSettings200Response) GetRtspUrlOk() (*string, bool)`

GetRtspUrlOk returns a tuple with the RtspUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtspUrl

`func (o *GetDeviceCameraVideoSettings200Response) SetRtspUrl(v string)`

SetRtspUrl sets RtspUrl field to given value.

### HasRtspUrl

`func (o *GetDeviceCameraVideoSettings200Response) HasRtspUrl() bool`

HasRtspUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


