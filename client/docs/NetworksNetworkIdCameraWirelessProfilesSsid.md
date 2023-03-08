# NetworksNetworkIdCameraWirelessProfilesSsid

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the SSID. | [optional] 
**AuthMode** | Pointer to **string** | The auth mode of the SSID. It can be set to (&#39;psk&#39;, &#39;8021x-radius&#39;). | [optional] 
**EncryptionMode** | Pointer to **string** | The encryption mode of the SSID. It can be set to (&#39;wpa&#39;, &#39;wpa-eap&#39;). With &#39;wpa&#39; mode, the authMode should be &#39;psk&#39; and with &#39;wpa-eap&#39; the authMode should be &#39;8021x-radius&#39; | [optional] 
**Psk** | Pointer to **string** | The pre-shared key of the SSID. | [optional] 

## Methods

### NewNetworksNetworkIdCameraWirelessProfilesSsid

`func NewNetworksNetworkIdCameraWirelessProfilesSsid() *NetworksNetworkIdCameraWirelessProfilesSsid`

NewNetworksNetworkIdCameraWirelessProfilesSsid instantiates a new NetworksNetworkIdCameraWirelessProfilesSsid object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdCameraWirelessProfilesSsidWithDefaults

`func NewNetworksNetworkIdCameraWirelessProfilesSsidWithDefaults() *NetworksNetworkIdCameraWirelessProfilesSsid`

NewNetworksNetworkIdCameraWirelessProfilesSsidWithDefaults instantiates a new NetworksNetworkIdCameraWirelessProfilesSsid object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NetworksNetworkIdCameraWirelessProfilesSsid) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworksNetworkIdCameraWirelessProfilesSsid) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworksNetworkIdCameraWirelessProfilesSsid) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworksNetworkIdCameraWirelessProfilesSsid) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAuthMode

`func (o *NetworksNetworkIdCameraWirelessProfilesSsid) GetAuthMode() string`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *NetworksNetworkIdCameraWirelessProfilesSsid) GetAuthModeOk() (*string, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *NetworksNetworkIdCameraWirelessProfilesSsid) SetAuthMode(v string)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *NetworksNetworkIdCameraWirelessProfilesSsid) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetEncryptionMode

`func (o *NetworksNetworkIdCameraWirelessProfilesSsid) GetEncryptionMode() string`

GetEncryptionMode returns the EncryptionMode field if non-nil, zero value otherwise.

### GetEncryptionModeOk

`func (o *NetworksNetworkIdCameraWirelessProfilesSsid) GetEncryptionModeOk() (*string, bool)`

GetEncryptionModeOk returns a tuple with the EncryptionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionMode

`func (o *NetworksNetworkIdCameraWirelessProfilesSsid) SetEncryptionMode(v string)`

SetEncryptionMode sets EncryptionMode field to given value.

### HasEncryptionMode

`func (o *NetworksNetworkIdCameraWirelessProfilesSsid) HasEncryptionMode() bool`

HasEncryptionMode returns a boolean if a field has been set.

### GetPsk

`func (o *NetworksNetworkIdCameraWirelessProfilesSsid) GetPsk() string`

GetPsk returns the Psk field if non-nil, zero value otherwise.

### GetPskOk

`func (o *NetworksNetworkIdCameraWirelessProfilesSsid) GetPskOk() (*string, bool)`

GetPskOk returns a tuple with the Psk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsk

`func (o *NetworksNetworkIdCameraWirelessProfilesSsid) SetPsk(v string)`

SetPsk sets Psk field to given value.

### HasPsk

`func (o *NetworksNetworkIdCameraWirelessProfilesSsid) HasPsk() bool`

HasPsk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


