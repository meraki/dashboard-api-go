# UpdateNetworkCameraWirelessProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the camera wireless profile. | [optional] 
**Ssid** | Pointer to [**CreateNetworkCameraWirelessProfileRequestSsid**](CreateNetworkCameraWirelessProfileRequestSsid.md) |  | [optional] 
**Identity** | Pointer to [**CreateNetworkCameraWirelessProfileRequestIdentity**](CreateNetworkCameraWirelessProfileRequestIdentity.md) |  | [optional] 

## Methods

### NewUpdateNetworkCameraWirelessProfileRequest

`func NewUpdateNetworkCameraWirelessProfileRequest() *UpdateNetworkCameraWirelessProfileRequest`

NewUpdateNetworkCameraWirelessProfileRequest instantiates a new UpdateNetworkCameraWirelessProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkCameraWirelessProfileRequestWithDefaults

`func NewUpdateNetworkCameraWirelessProfileRequestWithDefaults() *UpdateNetworkCameraWirelessProfileRequest`

NewUpdateNetworkCameraWirelessProfileRequestWithDefaults instantiates a new UpdateNetworkCameraWirelessProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateNetworkCameraWirelessProfileRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNetworkCameraWirelessProfileRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNetworkCameraWirelessProfileRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateNetworkCameraWirelessProfileRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSsid

`func (o *UpdateNetworkCameraWirelessProfileRequest) GetSsid() CreateNetworkCameraWirelessProfileRequestSsid`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *UpdateNetworkCameraWirelessProfileRequest) GetSsidOk() (*CreateNetworkCameraWirelessProfileRequestSsid, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *UpdateNetworkCameraWirelessProfileRequest) SetSsid(v CreateNetworkCameraWirelessProfileRequestSsid)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *UpdateNetworkCameraWirelessProfileRequest) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetIdentity

`func (o *UpdateNetworkCameraWirelessProfileRequest) GetIdentity() CreateNetworkCameraWirelessProfileRequestIdentity`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *UpdateNetworkCameraWirelessProfileRequest) GetIdentityOk() (*CreateNetworkCameraWirelessProfileRequestIdentity, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *UpdateNetworkCameraWirelessProfileRequest) SetIdentity(v CreateNetworkCameraWirelessProfileRequestIdentity)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *UpdateNetworkCameraWirelessProfileRequest) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


