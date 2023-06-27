# UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | IP address to which the APs will send RADIUS accounting messages | 
**Port** | Pointer to **int32** | Port on the RADIUS server that is listening for accounting messages | [optional] 
**Secret** | Pointer to **string** | Shared key used to authenticate messages between the APs and RADIUS server | [optional] 
**RadsecEnabled** | Pointer to **bool** | Use RADSEC (TLS over TCP) to connect to this RADIUS accounting server. Requires radiusProxyEnabled. | [optional] 
**CaCertificate** | Pointer to **string** | Certificate used for authorization for the RADSEC Server | [optional] 

## Methods

### NewUpdateNetworkWirelessSsidRequestRadiusAccountingServersInner

`func NewUpdateNetworkWirelessSsidRequestRadiusAccountingServersInner(host string, ) *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner`

NewUpdateNetworkWirelessSsidRequestRadiusAccountingServersInner instantiates a new UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkWirelessSsidRequestRadiusAccountingServersInnerWithDefaults

`func NewUpdateNetworkWirelessSsidRequestRadiusAccountingServersInnerWithDefaults() *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner`

NewUpdateNetworkWirelessSsidRequestRadiusAccountingServersInnerWithDefaults instantiates a new UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSecret

`func (o *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetRadsecEnabled

`func (o *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) GetRadsecEnabled() bool`

GetRadsecEnabled returns the RadsecEnabled field if non-nil, zero value otherwise.

### GetRadsecEnabledOk

`func (o *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) GetRadsecEnabledOk() (*bool, bool)`

GetRadsecEnabledOk returns a tuple with the RadsecEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadsecEnabled

`func (o *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) SetRadsecEnabled(v bool)`

SetRadsecEnabled sets RadsecEnabled field to given value.

### HasRadsecEnabled

`func (o *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) HasRadsecEnabled() bool`

HasRadsecEnabled returns a boolean if a field has been set.

### GetCaCertificate

`func (o *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) GetCaCertificate() string`

GetCaCertificate returns the CaCertificate field if non-nil, zero value otherwise.

### GetCaCertificateOk

`func (o *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) GetCaCertificateOk() (*string, bool)`

GetCaCertificateOk returns a tuple with the CaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificate

`func (o *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) SetCaCertificate(v string)`

SetCaCertificate sets CaCertificate field to given value.

### HasCaCertificate

`func (o *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) HasCaCertificate() bool`

HasCaCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


