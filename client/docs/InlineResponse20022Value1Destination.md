# InlineResponse20022Value1Destination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **string** | E.g.: \&quot;any\&quot;, \&quot;0\&quot; (also means \&quot;any\&quot;), \&quot;8080\&quot;, \&quot;1-1024\&quot; | [optional] 
**Cidr** | Pointer to **string** | CIDR format address (e.g.\&quot;192.168.10.1\&quot;, which is the same as \&quot;192.168.10.1/32\&quot;), or \&quot;any\&quot;. Cannot be used in combination with the \&quot;vlan\&quot; or \&quot;fqdn\&quot; property | [optional] 
**Network** | Pointer to **string** | Meraki network ID. Currently only available under a template network, and the value should be ID of either same template network, or another template network currently. E.g.: \&quot;L_12345678\&quot;. | [optional] 
**Vlan** | Pointer to **int32** | VLAN ID of the configured VLAN in the Meraki network. Cannot be used in combination with the \&quot;cidr\&quot; or \&quot;fqdn\&quot; property and is currently only available under a template network. | [optional] 
**Host** | Pointer to **int32** | Host ID in the VLAN. Should not exceed the VLAN subnet capacity. Must be used along with the \&quot;vlan\&quot; property and is currently only available under a template network. | [optional] 
**Fqdn** | Pointer to **string** | FQDN format address. Cannot be used in combination with the \&quot;cidr\&quot; or \&quot;fqdn\&quot; property and is currently only available in the \&quot;destination\&quot; object of the \&quot;vpnTrafficUplinkPreference\&quot; object. E.g.: \&quot;www.google.com\&quot; | [optional] 

## Methods

### NewInlineResponse20022Value1Destination

`func NewInlineResponse20022Value1Destination() *InlineResponse20022Value1Destination`

NewInlineResponse20022Value1Destination instantiates a new InlineResponse20022Value1Destination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20022Value1DestinationWithDefaults

`func NewInlineResponse20022Value1DestinationWithDefaults() *InlineResponse20022Value1Destination`

NewInlineResponse20022Value1DestinationWithDefaults instantiates a new InlineResponse20022Value1Destination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *InlineResponse20022Value1Destination) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *InlineResponse20022Value1Destination) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *InlineResponse20022Value1Destination) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *InlineResponse20022Value1Destination) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetCidr

`func (o *InlineResponse20022Value1Destination) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *InlineResponse20022Value1Destination) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *InlineResponse20022Value1Destination) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *InlineResponse20022Value1Destination) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse20022Value1Destination) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse20022Value1Destination) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse20022Value1Destination) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse20022Value1Destination) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetVlan

`func (o *InlineResponse20022Value1Destination) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineResponse20022Value1Destination) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineResponse20022Value1Destination) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *InlineResponse20022Value1Destination) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetHost

`func (o *InlineResponse20022Value1Destination) GetHost() int32`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *InlineResponse20022Value1Destination) GetHostOk() (*int32, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *InlineResponse20022Value1Destination) SetHost(v int32)`

SetHost sets Host field to given value.

### HasHost

`func (o *InlineResponse20022Value1Destination) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetFqdn

`func (o *InlineResponse20022Value1Destination) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *InlineResponse20022Value1Destination) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *InlineResponse20022Value1Destination) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *InlineResponse20022Value1Destination) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


