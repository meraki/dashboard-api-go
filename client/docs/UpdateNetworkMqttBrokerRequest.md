# UpdateNetworkMqttBrokerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the MQTT broker. | [optional] 
**Host** | Pointer to **string** | Host name/IP address where the MQTT broker runs. | [optional] 
**Port** | Pointer to **int32** | Host port though which the MQTT broker can be reached. | [optional] 
**Security** | Pointer to [**CreateNetworkMqttBrokerRequestSecurity**](CreateNetworkMqttBrokerRequestSecurity.md) |  | [optional] 
**Authentication** | Pointer to **map[string]interface{}** | Authentication settings of the MQTT broker | [optional] 

## Methods

### NewUpdateNetworkMqttBrokerRequest

`func NewUpdateNetworkMqttBrokerRequest() *UpdateNetworkMqttBrokerRequest`

NewUpdateNetworkMqttBrokerRequest instantiates a new UpdateNetworkMqttBrokerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkMqttBrokerRequestWithDefaults

`func NewUpdateNetworkMqttBrokerRequestWithDefaults() *UpdateNetworkMqttBrokerRequest`

NewUpdateNetworkMqttBrokerRequestWithDefaults instantiates a new UpdateNetworkMqttBrokerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateNetworkMqttBrokerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNetworkMqttBrokerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNetworkMqttBrokerRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateNetworkMqttBrokerRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHost

`func (o *UpdateNetworkMqttBrokerRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *UpdateNetworkMqttBrokerRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *UpdateNetworkMqttBrokerRequest) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *UpdateNetworkMqttBrokerRequest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *UpdateNetworkMqttBrokerRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UpdateNetworkMqttBrokerRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UpdateNetworkMqttBrokerRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *UpdateNetworkMqttBrokerRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSecurity

`func (o *UpdateNetworkMqttBrokerRequest) GetSecurity() CreateNetworkMqttBrokerRequestSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *UpdateNetworkMqttBrokerRequest) GetSecurityOk() (*CreateNetworkMqttBrokerRequestSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *UpdateNetworkMqttBrokerRequest) SetSecurity(v CreateNetworkMqttBrokerRequestSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *UpdateNetworkMqttBrokerRequest) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetAuthentication

`func (o *UpdateNetworkMqttBrokerRequest) GetAuthentication() map[string]interface{}`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *UpdateNetworkMqttBrokerRequest) GetAuthenticationOk() (*map[string]interface{}, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *UpdateNetworkMqttBrokerRequest) SetAuthentication(v map[string]interface{})`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *UpdateNetworkMqttBrokerRequest) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


