# NetworksNetworkIdWirelessSsidsNumberLdap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Servers** | Pointer to [**[]NetworksNetworkIdWirelessSsidsNumberLdapServers**](NetworksNetworkIdWirelessSsidsNumberLdapServers.md) | The LDAP servers to be used for authentication. | [optional] 
**Credentials** | Pointer to [**NetworksNetworkIdWirelessSsidsNumberLdapCredentials**](NetworksNetworkIdWirelessSsidsNumberLdapCredentials.md) |  | [optional] 
**BaseDistinguishedName** | Pointer to **string** | The base distinguished name of users on the LDAP server. | [optional] 
**ServerCaCertificate** | Pointer to [**NetworksNetworkIdWirelessSsidsNumberLdapServerCaCertificate**](NetworksNetworkIdWirelessSsidsNumberLdapServerCaCertificate.md) |  | [optional] 

## Methods

### NewNetworksNetworkIdWirelessSsidsNumberLdap

`func NewNetworksNetworkIdWirelessSsidsNumberLdap() *NetworksNetworkIdWirelessSsidsNumberLdap`

NewNetworksNetworkIdWirelessSsidsNumberLdap instantiates a new NetworksNetworkIdWirelessSsidsNumberLdap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdWirelessSsidsNumberLdapWithDefaults

`func NewNetworksNetworkIdWirelessSsidsNumberLdapWithDefaults() *NetworksNetworkIdWirelessSsidsNumberLdap`

NewNetworksNetworkIdWirelessSsidsNumberLdapWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberLdap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServers

`func (o *NetworksNetworkIdWirelessSsidsNumberLdap) GetServers() []NetworksNetworkIdWirelessSsidsNumberLdapServers`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *NetworksNetworkIdWirelessSsidsNumberLdap) GetServersOk() (*[]NetworksNetworkIdWirelessSsidsNumberLdapServers, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *NetworksNetworkIdWirelessSsidsNumberLdap) SetServers(v []NetworksNetworkIdWirelessSsidsNumberLdapServers)`

SetServers sets Servers field to given value.

### HasServers

`func (o *NetworksNetworkIdWirelessSsidsNumberLdap) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetCredentials

`func (o *NetworksNetworkIdWirelessSsidsNumberLdap) GetCredentials() NetworksNetworkIdWirelessSsidsNumberLdapCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *NetworksNetworkIdWirelessSsidsNumberLdap) GetCredentialsOk() (*NetworksNetworkIdWirelessSsidsNumberLdapCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *NetworksNetworkIdWirelessSsidsNumberLdap) SetCredentials(v NetworksNetworkIdWirelessSsidsNumberLdapCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *NetworksNetworkIdWirelessSsidsNumberLdap) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetBaseDistinguishedName

`func (o *NetworksNetworkIdWirelessSsidsNumberLdap) GetBaseDistinguishedName() string`

GetBaseDistinguishedName returns the BaseDistinguishedName field if non-nil, zero value otherwise.

### GetBaseDistinguishedNameOk

`func (o *NetworksNetworkIdWirelessSsidsNumberLdap) GetBaseDistinguishedNameOk() (*string, bool)`

GetBaseDistinguishedNameOk returns a tuple with the BaseDistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDistinguishedName

`func (o *NetworksNetworkIdWirelessSsidsNumberLdap) SetBaseDistinguishedName(v string)`

SetBaseDistinguishedName sets BaseDistinguishedName field to given value.

### HasBaseDistinguishedName

`func (o *NetworksNetworkIdWirelessSsidsNumberLdap) HasBaseDistinguishedName() bool`

HasBaseDistinguishedName returns a boolean if a field has been set.

### GetServerCaCertificate

`func (o *NetworksNetworkIdWirelessSsidsNumberLdap) GetServerCaCertificate() NetworksNetworkIdWirelessSsidsNumberLdapServerCaCertificate`

GetServerCaCertificate returns the ServerCaCertificate field if non-nil, zero value otherwise.

### GetServerCaCertificateOk

`func (o *NetworksNetworkIdWirelessSsidsNumberLdap) GetServerCaCertificateOk() (*NetworksNetworkIdWirelessSsidsNumberLdapServerCaCertificate, bool)`

GetServerCaCertificateOk returns a tuple with the ServerCaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCaCertificate

`func (o *NetworksNetworkIdWirelessSsidsNumberLdap) SetServerCaCertificate(v NetworksNetworkIdWirelessSsidsNumberLdapServerCaCertificate)`

SetServerCaCertificate sets ServerCaCertificate field to given value.

### HasServerCaCertificate

`func (o *NetworksNetworkIdWirelessSsidsNumberLdap) HasServerCaCertificate() bool`

HasServerCaCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


