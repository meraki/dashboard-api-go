# InlineObject91

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the MQTT broker. | 
**Host** | **string** | Host name/IP address where the MQTT broker runs. | 
**Port** | **int32** | Host port though which the MQTT broker can be reached. | 
**Security** | Pointer to [**NetworksNetworkIdMqttBrokersSecurity**](NetworksNetworkIdMqttBrokersSecurity.md) |  | [optional] 
**Authentication** | Pointer to **map[string]interface{}** | Authentication settings of the MQTT broker | [optional] 

## Methods

### NewInlineObject91

`func NewInlineObject91(name string, host string, port int32, ) *InlineObject91`

NewInlineObject91 instantiates a new InlineObject91 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject91WithDefaults

`func NewInlineObject91WithDefaults() *InlineObject91`

NewInlineObject91WithDefaults instantiates a new InlineObject91 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject91) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject91) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject91) SetName(v string)`

SetName sets Name field to given value.


### GetHost

`func (o *InlineObject91) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *InlineObject91) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *InlineObject91) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *InlineObject91) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *InlineObject91) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *InlineObject91) SetPort(v int32)`

SetPort sets Port field to given value.


### GetSecurity

`func (o *InlineObject91) GetSecurity() NetworksNetworkIdMqttBrokersSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *InlineObject91) GetSecurityOk() (*NetworksNetworkIdMqttBrokersSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *InlineObject91) SetSecurity(v NetworksNetworkIdMqttBrokersSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *InlineObject91) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetAuthentication

`func (o *InlineObject91) GetAuthentication() map[string]interface{}`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *InlineObject91) GetAuthenticationOk() (*map[string]interface{}, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *InlineObject91) SetAuthentication(v map[string]interface{})`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *InlineObject91) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


