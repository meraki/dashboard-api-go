# UpdateNetworkWirelessSsidRequestLocalRadius

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheTimeout** | Pointer to **int32** | The duration (in seconds) for which LDAP and OCSP lookups are cached. | [optional] 
**PasswordAuthentication** | Pointer to [**UpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication**](UpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication.md) |  | [optional] 
**CertificateAuthentication** | Pointer to [**UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication**](UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication.md) |  | [optional] 

## Methods

### NewUpdateNetworkWirelessSsidRequestLocalRadius

`func NewUpdateNetworkWirelessSsidRequestLocalRadius() *UpdateNetworkWirelessSsidRequestLocalRadius`

NewUpdateNetworkWirelessSsidRequestLocalRadius instantiates a new UpdateNetworkWirelessSsidRequestLocalRadius object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkWirelessSsidRequestLocalRadiusWithDefaults

`func NewUpdateNetworkWirelessSsidRequestLocalRadiusWithDefaults() *UpdateNetworkWirelessSsidRequestLocalRadius`

NewUpdateNetworkWirelessSsidRequestLocalRadiusWithDefaults instantiates a new UpdateNetworkWirelessSsidRequestLocalRadius object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheTimeout

`func (o *UpdateNetworkWirelessSsidRequestLocalRadius) GetCacheTimeout() int32`

GetCacheTimeout returns the CacheTimeout field if non-nil, zero value otherwise.

### GetCacheTimeoutOk

`func (o *UpdateNetworkWirelessSsidRequestLocalRadius) GetCacheTimeoutOk() (*int32, bool)`

GetCacheTimeoutOk returns a tuple with the CacheTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheTimeout

`func (o *UpdateNetworkWirelessSsidRequestLocalRadius) SetCacheTimeout(v int32)`

SetCacheTimeout sets CacheTimeout field to given value.

### HasCacheTimeout

`func (o *UpdateNetworkWirelessSsidRequestLocalRadius) HasCacheTimeout() bool`

HasCacheTimeout returns a boolean if a field has been set.

### GetPasswordAuthentication

`func (o *UpdateNetworkWirelessSsidRequestLocalRadius) GetPasswordAuthentication() UpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication`

GetPasswordAuthentication returns the PasswordAuthentication field if non-nil, zero value otherwise.

### GetPasswordAuthenticationOk

`func (o *UpdateNetworkWirelessSsidRequestLocalRadius) GetPasswordAuthenticationOk() (*UpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication, bool)`

GetPasswordAuthenticationOk returns a tuple with the PasswordAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordAuthentication

`func (o *UpdateNetworkWirelessSsidRequestLocalRadius) SetPasswordAuthentication(v UpdateNetworkWirelessSsidRequestLocalRadiusPasswordAuthentication)`

SetPasswordAuthentication sets PasswordAuthentication field to given value.

### HasPasswordAuthentication

`func (o *UpdateNetworkWirelessSsidRequestLocalRadius) HasPasswordAuthentication() bool`

HasPasswordAuthentication returns a boolean if a field has been set.

### GetCertificateAuthentication

`func (o *UpdateNetworkWirelessSsidRequestLocalRadius) GetCertificateAuthentication() UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication`

GetCertificateAuthentication returns the CertificateAuthentication field if non-nil, zero value otherwise.

### GetCertificateAuthenticationOk

`func (o *UpdateNetworkWirelessSsidRequestLocalRadius) GetCertificateAuthenticationOk() (*UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication, bool)`

GetCertificateAuthenticationOk returns a tuple with the CertificateAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthentication

`func (o *UpdateNetworkWirelessSsidRequestLocalRadius) SetCertificateAuthentication(v UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication)`

SetCertificateAuthentication sets CertificateAuthentication field to given value.

### HasCertificateAuthentication

`func (o *UpdateNetworkWirelessSsidRequestLocalRadius) HasCertificateAuthentication() bool`

HasCertificateAuthentication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


