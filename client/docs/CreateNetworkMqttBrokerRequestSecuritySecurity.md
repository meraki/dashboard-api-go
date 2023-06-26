# CreateNetworkMqttBrokerRequestSecuritySecurity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaCertificate** | Pointer to **string** | CA Certificate of the MQTT broker. | [optional] 
**VerifyHostnames** | Pointer to **bool** | Whether the TLS hostname verification is enabled for the MQTT broker. | [optional] 

## Methods

### NewCreateNetworkMqttBrokerRequestSecuritySecurity

`func NewCreateNetworkMqttBrokerRequestSecuritySecurity() *CreateNetworkMqttBrokerRequestSecuritySecurity`

NewCreateNetworkMqttBrokerRequestSecuritySecurity instantiates a new CreateNetworkMqttBrokerRequestSecuritySecurity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkMqttBrokerRequestSecuritySecurityWithDefaults

`func NewCreateNetworkMqttBrokerRequestSecuritySecurityWithDefaults() *CreateNetworkMqttBrokerRequestSecuritySecurity`

NewCreateNetworkMqttBrokerRequestSecuritySecurityWithDefaults instantiates a new CreateNetworkMqttBrokerRequestSecuritySecurity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaCertificate

`func (o *CreateNetworkMqttBrokerRequestSecuritySecurity) GetCaCertificate() string`

GetCaCertificate returns the CaCertificate field if non-nil, zero value otherwise.

### GetCaCertificateOk

`func (o *CreateNetworkMqttBrokerRequestSecuritySecurity) GetCaCertificateOk() (*string, bool)`

GetCaCertificateOk returns a tuple with the CaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificate

`func (o *CreateNetworkMqttBrokerRequestSecuritySecurity) SetCaCertificate(v string)`

SetCaCertificate sets CaCertificate field to given value.

### HasCaCertificate

`func (o *CreateNetworkMqttBrokerRequestSecuritySecurity) HasCaCertificate() bool`

HasCaCertificate returns a boolean if a field has been set.

### GetVerifyHostnames

`func (o *CreateNetworkMqttBrokerRequestSecuritySecurity) GetVerifyHostnames() bool`

GetVerifyHostnames returns the VerifyHostnames field if non-nil, zero value otherwise.

### GetVerifyHostnamesOk

`func (o *CreateNetworkMqttBrokerRequestSecuritySecurity) GetVerifyHostnamesOk() (*bool, bool)`

GetVerifyHostnamesOk returns a tuple with the VerifyHostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyHostnames

`func (o *CreateNetworkMqttBrokerRequestSecuritySecurity) SetVerifyHostnames(v bool)`

SetVerifyHostnames sets VerifyHostnames field to given value.

### HasVerifyHostnames

`func (o *CreateNetworkMqttBrokerRequestSecuritySecurity) HasVerifyHostnames() bool`

HasVerifyHostnames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


