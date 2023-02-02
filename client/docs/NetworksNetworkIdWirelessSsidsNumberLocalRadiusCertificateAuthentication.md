# NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether or not to use EAP-TLS certificate-based authentication to validate wireless clients. | [optional] 
**UseLdap** | Pointer to **bool** | Whether or not to verify the certificate with LDAP. | [optional] 
**UseOcsp** | Pointer to **bool** | Whether or not to verify the certificate with OCSP. | [optional] 
**OcspResponderUrl** | Pointer to **string** | (Optional) The URL of the OCSP responder to verify client certificate status. | [optional] 
**ClientRootCaCertificate** | Pointer to [**NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationClientRootCaCertificate**](NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationClientRootCaCertificate.md) |  | [optional] 

## Methods

### NewNetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication

`func NewNetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication() *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication`

NewNetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication instantiates a new NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationWithDefaults

`func NewNetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationWithDefaults() *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication`

NewNetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetUseLdap

`func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) GetUseLdap() bool`

GetUseLdap returns the UseLdap field if non-nil, zero value otherwise.

### GetUseLdapOk

`func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) GetUseLdapOk() (*bool, bool)`

GetUseLdapOk returns a tuple with the UseLdap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLdap

`func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) SetUseLdap(v bool)`

SetUseLdap sets UseLdap field to given value.

### HasUseLdap

`func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) HasUseLdap() bool`

HasUseLdap returns a boolean if a field has been set.

### GetUseOcsp

`func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) GetUseOcsp() bool`

GetUseOcsp returns the UseOcsp field if non-nil, zero value otherwise.

### GetUseOcspOk

`func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) GetUseOcspOk() (*bool, bool)`

GetUseOcspOk returns a tuple with the UseOcsp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOcsp

`func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) SetUseOcsp(v bool)`

SetUseOcsp sets UseOcsp field to given value.

### HasUseOcsp

`func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) HasUseOcsp() bool`

HasUseOcsp returns a boolean if a field has been set.

### GetOcspResponderUrl

`func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) GetOcspResponderUrl() string`

GetOcspResponderUrl returns the OcspResponderUrl field if non-nil, zero value otherwise.

### GetOcspResponderUrlOk

`func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) GetOcspResponderUrlOk() (*string, bool)`

GetOcspResponderUrlOk returns a tuple with the OcspResponderUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspResponderUrl

`func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) SetOcspResponderUrl(v string)`

SetOcspResponderUrl sets OcspResponderUrl field to given value.

### HasOcspResponderUrl

`func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) HasOcspResponderUrl() bool`

HasOcspResponderUrl returns a boolean if a field has been set.

### GetClientRootCaCertificate

`func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) GetClientRootCaCertificate() NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationClientRootCaCertificate`

GetClientRootCaCertificate returns the ClientRootCaCertificate field if non-nil, zero value otherwise.

### GetClientRootCaCertificateOk

`func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) GetClientRootCaCertificateOk() (*NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationClientRootCaCertificate, bool)`

GetClientRootCaCertificateOk returns a tuple with the ClientRootCaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRootCaCertificate

`func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) SetClientRootCaCertificate(v NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationClientRootCaCertificate)`

SetClientRootCaCertificate sets ClientRootCaCertificate field to given value.

### HasClientRootCaCertificate

`func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) HasClientRootCaCertificate() bool`

HasClientRootCaCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


