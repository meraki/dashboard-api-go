# UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether or not to use EAP-TLS certificate-based authentication to validate wireless clients. | [optional] 
**UseLdap** | Pointer to **bool** | Whether or not to verify the certificate with LDAP. | [optional] 
**UseOcsp** | Pointer to **bool** | Whether or not to verify the certificate with OCSP. | [optional] 
**OcspResponderUrl** | Pointer to **string** | (Optional) The URL of the OCSP responder to verify client certificate status. | [optional] 
**ClientRootCaCertificate** | Pointer to [**UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthenticationClientRootCaCertificate**](UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthenticationClientRootCaCertificate.md) |  | [optional] 

## Methods

### NewUpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication

`func NewUpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication() *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication`

NewUpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication instantiates a new UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthenticationWithDefaults

`func NewUpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthenticationWithDefaults() *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication`

NewUpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthenticationWithDefaults instantiates a new UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetUseLdap

`func (o *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) GetUseLdap() bool`

GetUseLdap returns the UseLdap field if non-nil, zero value otherwise.

### GetUseLdapOk

`func (o *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) GetUseLdapOk() (*bool, bool)`

GetUseLdapOk returns a tuple with the UseLdap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLdap

`func (o *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) SetUseLdap(v bool)`

SetUseLdap sets UseLdap field to given value.

### HasUseLdap

`func (o *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) HasUseLdap() bool`

HasUseLdap returns a boolean if a field has been set.

### GetUseOcsp

`func (o *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) GetUseOcsp() bool`

GetUseOcsp returns the UseOcsp field if non-nil, zero value otherwise.

### GetUseOcspOk

`func (o *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) GetUseOcspOk() (*bool, bool)`

GetUseOcspOk returns a tuple with the UseOcsp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOcsp

`func (o *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) SetUseOcsp(v bool)`

SetUseOcsp sets UseOcsp field to given value.

### HasUseOcsp

`func (o *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) HasUseOcsp() bool`

HasUseOcsp returns a boolean if a field has been set.

### GetOcspResponderUrl

`func (o *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) GetOcspResponderUrl() string`

GetOcspResponderUrl returns the OcspResponderUrl field if non-nil, zero value otherwise.

### GetOcspResponderUrlOk

`func (o *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) GetOcspResponderUrlOk() (*string, bool)`

GetOcspResponderUrlOk returns a tuple with the OcspResponderUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspResponderUrl

`func (o *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) SetOcspResponderUrl(v string)`

SetOcspResponderUrl sets OcspResponderUrl field to given value.

### HasOcspResponderUrl

`func (o *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) HasOcspResponderUrl() bool`

HasOcspResponderUrl returns a boolean if a field has been set.

### GetClientRootCaCertificate

`func (o *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) GetClientRootCaCertificate() UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthenticationClientRootCaCertificate`

GetClientRootCaCertificate returns the ClientRootCaCertificate field if non-nil, zero value otherwise.

### GetClientRootCaCertificateOk

`func (o *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) GetClientRootCaCertificateOk() (*UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthenticationClientRootCaCertificate, bool)`

GetClientRootCaCertificateOk returns a tuple with the ClientRootCaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRootCaCertificate

`func (o *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) SetClientRootCaCertificate(v UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthenticationClientRootCaCertificate)`

SetClientRootCaCertificate sets ClientRootCaCertificate field to given value.

### HasClientRootCaCertificate

`func (o *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) HasClientRootCaCertificate() bool`

HasClientRootCaCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


