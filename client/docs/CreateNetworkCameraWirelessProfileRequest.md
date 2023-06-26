# CreateNetworkCameraWirelessProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the camera wireless profile. This parameter is required. | 
**Ssid** | [**CreateNetworkCameraWirelessProfileRequestSsid**](CreateNetworkCameraWirelessProfileRequestSsid.md) |  | 
**Identity** | Pointer to [**CreateNetworkCameraWirelessProfileRequestIdentity**](CreateNetworkCameraWirelessProfileRequestIdentity.md) |  | [optional] 

## Methods

### NewCreateNetworkCameraWirelessProfileRequest

`func NewCreateNetworkCameraWirelessProfileRequest(name string, ssid CreateNetworkCameraWirelessProfileRequestSsid, ) *CreateNetworkCameraWirelessProfileRequest`

NewCreateNetworkCameraWirelessProfileRequest instantiates a new CreateNetworkCameraWirelessProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkCameraWirelessProfileRequestWithDefaults

`func NewCreateNetworkCameraWirelessProfileRequestWithDefaults() *CreateNetworkCameraWirelessProfileRequest`

NewCreateNetworkCameraWirelessProfileRequestWithDefaults instantiates a new CreateNetworkCameraWirelessProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNetworkCameraWirelessProfileRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkCameraWirelessProfileRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkCameraWirelessProfileRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSsid

`func (o *CreateNetworkCameraWirelessProfileRequest) GetSsid() CreateNetworkCameraWirelessProfileRequestSsid`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *CreateNetworkCameraWirelessProfileRequest) GetSsidOk() (*CreateNetworkCameraWirelessProfileRequestSsid, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *CreateNetworkCameraWirelessProfileRequest) SetSsid(v CreateNetworkCameraWirelessProfileRequestSsid)`

SetSsid sets Ssid field to given value.


### GetIdentity

`func (o *CreateNetworkCameraWirelessProfileRequest) GetIdentity() CreateNetworkCameraWirelessProfileRequestIdentity`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *CreateNetworkCameraWirelessProfileRequest) GetIdentityOk() (*CreateNetworkCameraWirelessProfileRequestIdentity, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *CreateNetworkCameraWirelessProfileRequest) SetIdentity(v CreateNetworkCameraWirelessProfileRequestIdentity)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *CreateNetworkCameraWirelessProfileRequest) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


