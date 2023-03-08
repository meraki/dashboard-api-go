# NetworksNetworkIdWirelessSsidsNumberRadiusServers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | IP address of your RADIUS server | 
**Port** | Pointer to **int32** | UDP port the RADIUS server listens on for Access-requests | [optional] 
**Secret** | Pointer to **string** | RADIUS client shared secret | [optional] 
**RadsecEnabled** | Pointer to **bool** | Use RADSEC (TLS over TCP) to connect to this RADIUS server. Requires radiusProxyEnabled. | [optional] 
**OpenRoamingCertificateId** | Pointer to **int32** | The ID of the Openroaming Certificate attached to radius server. | [optional] 
**CaCertificate** | Pointer to **string** | Certificate used for authorization for the RADSEC Server | [optional] 

## Methods

### NewNetworksNetworkIdWirelessSsidsNumberRadiusServers

`func NewNetworksNetworkIdWirelessSsidsNumberRadiusServers(host string, ) *NetworksNetworkIdWirelessSsidsNumberRadiusServers`

NewNetworksNetworkIdWirelessSsidsNumberRadiusServers instantiates a new NetworksNetworkIdWirelessSsidsNumberRadiusServers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdWirelessSsidsNumberRadiusServersWithDefaults

`func NewNetworksNetworkIdWirelessSsidsNumberRadiusServersWithDefaults() *NetworksNetworkIdWirelessSsidsNumberRadiusServers`

NewNetworksNetworkIdWirelessSsidsNumberRadiusServersWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberRadiusServers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSecret

`func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetRadsecEnabled

`func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) GetRadsecEnabled() bool`

GetRadsecEnabled returns the RadsecEnabled field if non-nil, zero value otherwise.

### GetRadsecEnabledOk

`func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) GetRadsecEnabledOk() (*bool, bool)`

GetRadsecEnabledOk returns a tuple with the RadsecEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadsecEnabled

`func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) SetRadsecEnabled(v bool)`

SetRadsecEnabled sets RadsecEnabled field to given value.

### HasRadsecEnabled

`func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) HasRadsecEnabled() bool`

HasRadsecEnabled returns a boolean if a field has been set.

### GetOpenRoamingCertificateId

`func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) GetOpenRoamingCertificateId() int32`

GetOpenRoamingCertificateId returns the OpenRoamingCertificateId field if non-nil, zero value otherwise.

### GetOpenRoamingCertificateIdOk

`func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) GetOpenRoamingCertificateIdOk() (*int32, bool)`

GetOpenRoamingCertificateIdOk returns a tuple with the OpenRoamingCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenRoamingCertificateId

`func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) SetOpenRoamingCertificateId(v int32)`

SetOpenRoamingCertificateId sets OpenRoamingCertificateId field to given value.

### HasOpenRoamingCertificateId

`func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) HasOpenRoamingCertificateId() bool`

HasOpenRoamingCertificateId returns a boolean if a field has been set.

### GetCaCertificate

`func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) GetCaCertificate() string`

GetCaCertificate returns the CaCertificate field if non-nil, zero value otherwise.

### GetCaCertificateOk

`func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) GetCaCertificateOk() (*string, bool)`

GetCaCertificateOk returns a tuple with the CaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificate

`func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) SetCaCertificate(v string)`

SetCaCertificate sets CaCertificate field to given value.

### HasCaCertificate

`func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) HasCaCertificate() bool`

HasCaCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


