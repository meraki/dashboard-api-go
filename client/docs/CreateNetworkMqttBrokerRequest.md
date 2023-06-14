# CreateNetworkMqttBrokerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the MQTT broker. | 
**Host** | **string** | Host name/IP address where the MQTT broker runs. | 
**Port** | **int32** | Host port though which the MQTT broker can be reached. | 
**Security** | Pointer to [**CreateNetworkMqttBrokerRequestSecurity**](CreateNetworkMqttBrokerRequestSecurity.md) |  | [optional] 
**Authentication** | Pointer to **map[string]interface{}** | Authentication settings of the MQTT broker | [optional] 

## Methods

### NewCreateNetworkMqttBrokerRequest

`func NewCreateNetworkMqttBrokerRequest(name string, host string, port int32, ) *CreateNetworkMqttBrokerRequest`

NewCreateNetworkMqttBrokerRequest instantiates a new CreateNetworkMqttBrokerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkMqttBrokerRequestWithDefaults

`func NewCreateNetworkMqttBrokerRequestWithDefaults() *CreateNetworkMqttBrokerRequest`

NewCreateNetworkMqttBrokerRequestWithDefaults instantiates a new CreateNetworkMqttBrokerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNetworkMqttBrokerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkMqttBrokerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkMqttBrokerRequest) SetName(v string)`

SetName sets Name field to given value.


### GetHost

`func (o *CreateNetworkMqttBrokerRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CreateNetworkMqttBrokerRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CreateNetworkMqttBrokerRequest) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *CreateNetworkMqttBrokerRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CreateNetworkMqttBrokerRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CreateNetworkMqttBrokerRequest) SetPort(v int32)`

SetPort sets Port field to given value.


### GetSecurity

`func (o *CreateNetworkMqttBrokerRequest) GetSecurity() CreateNetworkMqttBrokerRequestSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *CreateNetworkMqttBrokerRequest) GetSecurityOk() (*CreateNetworkMqttBrokerRequestSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *CreateNetworkMqttBrokerRequest) SetSecurity(v CreateNetworkMqttBrokerRequestSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *CreateNetworkMqttBrokerRequest) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetAuthentication

`func (o *CreateNetworkMqttBrokerRequest) GetAuthentication() map[string]interface{}`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *CreateNetworkMqttBrokerRequest) GetAuthenticationOk() (*map[string]interface{}, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *CreateNetworkMqttBrokerRequest) SetAuthentication(v map[string]interface{})`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *CreateNetworkMqttBrokerRequest) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


