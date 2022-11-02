# UpdateNetworkWirelessSsidRequestLdap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Servers** | Pointer to [**[]UpdateNetworkWirelessSsidRequestLdapServersInner**](UpdateNetworkWirelessSsidRequestLdapServersInner.md) | The LDAP servers to be used for authentication. | [optional] 
**Credentials** | Pointer to [**UpdateNetworkWirelessSsidRequestLdapCredentials**](UpdateNetworkWirelessSsidRequestLdapCredentials.md) |  | [optional] 
**BaseDistinguishedName** | Pointer to **string** | The base distinguished name of users on the LDAP server. | [optional] 
**ServerCaCertificate** | Pointer to [**UpdateNetworkWirelessSsidRequestLdapServerCaCertificate**](UpdateNetworkWirelessSsidRequestLdapServerCaCertificate.md) |  | [optional] 

## Methods

### NewUpdateNetworkWirelessSsidRequestLdap

`func NewUpdateNetworkWirelessSsidRequestLdap() *UpdateNetworkWirelessSsidRequestLdap`

NewUpdateNetworkWirelessSsidRequestLdap instantiates a new UpdateNetworkWirelessSsidRequestLdap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkWirelessSsidRequestLdapWithDefaults

`func NewUpdateNetworkWirelessSsidRequestLdapWithDefaults() *UpdateNetworkWirelessSsidRequestLdap`

NewUpdateNetworkWirelessSsidRequestLdapWithDefaults instantiates a new UpdateNetworkWirelessSsidRequestLdap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServers

`func (o *UpdateNetworkWirelessSsidRequestLdap) GetServers() []UpdateNetworkWirelessSsidRequestLdapServersInner`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *UpdateNetworkWirelessSsidRequestLdap) GetServersOk() (*[]UpdateNetworkWirelessSsidRequestLdapServersInner, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *UpdateNetworkWirelessSsidRequestLdap) SetServers(v []UpdateNetworkWirelessSsidRequestLdapServersInner)`

SetServers sets Servers field to given value.

### HasServers

`func (o *UpdateNetworkWirelessSsidRequestLdap) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetCredentials

`func (o *UpdateNetworkWirelessSsidRequestLdap) GetCredentials() UpdateNetworkWirelessSsidRequestLdapCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *UpdateNetworkWirelessSsidRequestLdap) GetCredentialsOk() (*UpdateNetworkWirelessSsidRequestLdapCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *UpdateNetworkWirelessSsidRequestLdap) SetCredentials(v UpdateNetworkWirelessSsidRequestLdapCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *UpdateNetworkWirelessSsidRequestLdap) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetBaseDistinguishedName

`func (o *UpdateNetworkWirelessSsidRequestLdap) GetBaseDistinguishedName() string`

GetBaseDistinguishedName returns the BaseDistinguishedName field if non-nil, zero value otherwise.

### GetBaseDistinguishedNameOk

`func (o *UpdateNetworkWirelessSsidRequestLdap) GetBaseDistinguishedNameOk() (*string, bool)`

GetBaseDistinguishedNameOk returns a tuple with the BaseDistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDistinguishedName

`func (o *UpdateNetworkWirelessSsidRequestLdap) SetBaseDistinguishedName(v string)`

SetBaseDistinguishedName sets BaseDistinguishedName field to given value.

### HasBaseDistinguishedName

`func (o *UpdateNetworkWirelessSsidRequestLdap) HasBaseDistinguishedName() bool`

HasBaseDistinguishedName returns a boolean if a field has been set.

### GetServerCaCertificate

`func (o *UpdateNetworkWirelessSsidRequestLdap) GetServerCaCertificate() UpdateNetworkWirelessSsidRequestLdapServerCaCertificate`

GetServerCaCertificate returns the ServerCaCertificate field if non-nil, zero value otherwise.

### GetServerCaCertificateOk

`func (o *UpdateNetworkWirelessSsidRequestLdap) GetServerCaCertificateOk() (*UpdateNetworkWirelessSsidRequestLdapServerCaCertificate, bool)`

GetServerCaCertificateOk returns a tuple with the ServerCaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCaCertificate

`func (o *UpdateNetworkWirelessSsidRequestLdap) SetServerCaCertificate(v UpdateNetworkWirelessSsidRequestLdapServerCaCertificate)`

SetServerCaCertificate sets ServerCaCertificate field to given value.

### HasServerCaCertificate

`func (o *UpdateNetworkWirelessSsidRequestLdap) HasServerCaCertificate() bool`

HasServerCaCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


