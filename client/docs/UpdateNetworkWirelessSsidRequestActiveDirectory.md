# UpdateNetworkWirelessSsidRequestActiveDirectory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Servers** | Pointer to [**[]UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner**](UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner.md) | The Active Directory servers to be used for authentication. | [optional] 
**Credentials** | Pointer to [**UpdateNetworkWirelessSsidRequestActiveDirectoryCredentials**](UpdateNetworkWirelessSsidRequestActiveDirectoryCredentials.md) |  | [optional] 

## Methods

### NewUpdateNetworkWirelessSsidRequestActiveDirectory

`func NewUpdateNetworkWirelessSsidRequestActiveDirectory() *UpdateNetworkWirelessSsidRequestActiveDirectory`

NewUpdateNetworkWirelessSsidRequestActiveDirectory instantiates a new UpdateNetworkWirelessSsidRequestActiveDirectory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkWirelessSsidRequestActiveDirectoryWithDefaults

`func NewUpdateNetworkWirelessSsidRequestActiveDirectoryWithDefaults() *UpdateNetworkWirelessSsidRequestActiveDirectory`

NewUpdateNetworkWirelessSsidRequestActiveDirectoryWithDefaults instantiates a new UpdateNetworkWirelessSsidRequestActiveDirectory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServers

`func (o *UpdateNetworkWirelessSsidRequestActiveDirectory) GetServers() []UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *UpdateNetworkWirelessSsidRequestActiveDirectory) GetServersOk() (*[]UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *UpdateNetworkWirelessSsidRequestActiveDirectory) SetServers(v []UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner)`

SetServers sets Servers field to given value.

### HasServers

`func (o *UpdateNetworkWirelessSsidRequestActiveDirectory) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetCredentials

`func (o *UpdateNetworkWirelessSsidRequestActiveDirectory) GetCredentials() UpdateNetworkWirelessSsidRequestActiveDirectoryCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *UpdateNetworkWirelessSsidRequestActiveDirectory) GetCredentialsOk() (*UpdateNetworkWirelessSsidRequestActiveDirectoryCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *UpdateNetworkWirelessSsidRequestActiveDirectory) SetCredentials(v UpdateNetworkWirelessSsidRequestActiveDirectoryCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *UpdateNetworkWirelessSsidRequestActiveDirectory) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


