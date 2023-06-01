# InlineResponse20019ValueSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **string** | E.g.: \&quot;any\&quot;, \&quot;0\&quot; (also means \&quot;any\&quot;), \&quot;8080\&quot;, \&quot;1-1024\&quot; | [optional] 
**Cidr** | Pointer to **string** | CIDR format address (e.g.\&quot;192.168.10.1\&quot;, which is the same as \&quot;192.168.10.1/32\&quot;), or \&quot;any\&quot;. Cannot be used in combination with the \&quot;vlan\&quot; property | [optional] 
**Vlan** | Pointer to **int32** | VLAN ID of the configured VLAN in the Meraki network. Cannot be used in combination with the \&quot;cidr\&quot; property and is currently only available under a template network. | [optional] 
**Host** | Pointer to **int32** | Host ID in the VLAN. Should not exceed the VLAN subnet capacity. Must be used along with the \&quot;vlan\&quot; property and is currently only available under a template network. | [optional] 

## Methods

### NewInlineResponse20019ValueSource

`func NewInlineResponse20019ValueSource() *InlineResponse20019ValueSource`

NewInlineResponse20019ValueSource instantiates a new InlineResponse20019ValueSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20019ValueSourceWithDefaults

`func NewInlineResponse20019ValueSourceWithDefaults() *InlineResponse20019ValueSource`

NewInlineResponse20019ValueSourceWithDefaults instantiates a new InlineResponse20019ValueSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *InlineResponse20019ValueSource) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *InlineResponse20019ValueSource) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *InlineResponse20019ValueSource) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *InlineResponse20019ValueSource) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetCidr

`func (o *InlineResponse20019ValueSource) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *InlineResponse20019ValueSource) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *InlineResponse20019ValueSource) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *InlineResponse20019ValueSource) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetVlan

`func (o *InlineResponse20019ValueSource) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineResponse20019ValueSource) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineResponse20019ValueSource) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *InlineResponse20019ValueSource) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetHost

`func (o *InlineResponse20019ValueSource) GetHost() int32`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *InlineResponse20019ValueSource) GetHostOk() (*int32, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *InlineResponse20019ValueSource) SetHost(v int32)`

SetHost sets Host field to given value.

### HasHost

`func (o *InlineResponse20019ValueSource) HasHost() bool`

HasHost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


