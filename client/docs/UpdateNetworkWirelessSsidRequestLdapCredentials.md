# UpdateNetworkWirelessSsidRequestLdapCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DistinguishedName** | Pointer to **string** | The distinguished name of the LDAP user account (example: cn&#x3D;user,dc&#x3D;meraki,dc&#x3D;com). | [optional] 
**Password** | Pointer to **string** | The password of the LDAP user account. | [optional] 

## Methods

### NewUpdateNetworkWirelessSsidRequestLdapCredentials

`func NewUpdateNetworkWirelessSsidRequestLdapCredentials() *UpdateNetworkWirelessSsidRequestLdapCredentials`

NewUpdateNetworkWirelessSsidRequestLdapCredentials instantiates a new UpdateNetworkWirelessSsidRequestLdapCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkWirelessSsidRequestLdapCredentialsWithDefaults

`func NewUpdateNetworkWirelessSsidRequestLdapCredentialsWithDefaults() *UpdateNetworkWirelessSsidRequestLdapCredentials`

NewUpdateNetworkWirelessSsidRequestLdapCredentialsWithDefaults instantiates a new UpdateNetworkWirelessSsidRequestLdapCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistinguishedName

`func (o *UpdateNetworkWirelessSsidRequestLdapCredentials) GetDistinguishedName() string`

GetDistinguishedName returns the DistinguishedName field if non-nil, zero value otherwise.

### GetDistinguishedNameOk

`func (o *UpdateNetworkWirelessSsidRequestLdapCredentials) GetDistinguishedNameOk() (*string, bool)`

GetDistinguishedNameOk returns a tuple with the DistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinguishedName

`func (o *UpdateNetworkWirelessSsidRequestLdapCredentials) SetDistinguishedName(v string)`

SetDistinguishedName sets DistinguishedName field to given value.

### HasDistinguishedName

`func (o *UpdateNetworkWirelessSsidRequestLdapCredentials) HasDistinguishedName() bool`

HasDistinguishedName returns a boolean if a field has been set.

### GetPassword

`func (o *UpdateNetworkWirelessSsidRequestLdapCredentials) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateNetworkWirelessSsidRequestLdapCredentials) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateNetworkWirelessSsidRequestLdapCredentials) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateNetworkWirelessSsidRequestLdapCredentials) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


