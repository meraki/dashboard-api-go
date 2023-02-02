# InlineObject92

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the MQTT broker. | [optional] 
**Host** | Pointer to **string** | Host name/IP address where the MQTT broker runs. | [optional] 
**Port** | Pointer to **int32** | Host port though which the MQTT broker can be reached. | [optional] 
**Security** | Pointer to [**NetworksNetworkIdMqttBrokersSecurity**](NetworksNetworkIdMqttBrokersSecurity.md) |  | [optional] 
**Authentication** | Pointer to **map[string]interface{}** | Authentication settings of the MQTT broker | [optional] 

## Methods

### NewInlineObject92

`func NewInlineObject92() *InlineObject92`

NewInlineObject92 instantiates a new InlineObject92 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject92WithDefaults

`func NewInlineObject92WithDefaults() *InlineObject92`

NewInlineObject92WithDefaults instantiates a new InlineObject92 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject92) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject92) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject92) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject92) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHost

`func (o *InlineObject92) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *InlineObject92) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *InlineObject92) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *InlineObject92) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *InlineObject92) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *InlineObject92) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *InlineObject92) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *InlineObject92) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSecurity

`func (o *InlineObject92) GetSecurity() NetworksNetworkIdMqttBrokersSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *InlineObject92) GetSecurityOk() (*NetworksNetworkIdMqttBrokersSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *InlineObject92) SetSecurity(v NetworksNetworkIdMqttBrokersSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *InlineObject92) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetAuthentication

`func (o *InlineObject92) GetAuthentication() map[string]interface{}`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *InlineObject92) GetAuthenticationOk() (*map[string]interface{}, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *InlineObject92) SetAuthentication(v map[string]interface{})`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *InlineObject92) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


