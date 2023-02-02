# NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **string** | E.g.: \&quot;any\&quot;, \&quot;0\&quot; (also means \&quot;any\&quot;), \&quot;8080\&quot;, \&quot;1-1024\&quot; | [optional] 
**Cidr** | Pointer to **string** | CIDR format address, or \&quot;any\&quot;. E.g.: \&quot;192.168.10.0/24\&quot;,  \&quot;192.168.10.1\&quot; (same as \&quot;192.168.10.1/32\&quot;), \&quot;0.0.0.0/0\&quot; (same as \&quot;any\&quot;) | [optional] 
**Vlan** | Pointer to **int32** | VLAN ID of the configured VLAN in the Meraki network. Currently only available under a template network. | [optional] 
**Host** | Pointer to **int32** | Host ID in the VLAN, should be used along with &#39;vlan&#39;, and not exceed the vlan subnet capacity. Currently only available under a template network. | [optional] 

## Methods

### NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource

`func NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource() *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource`

NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource instantiates a new NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSourceWithDefaults

`func NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSourceWithDefaults() *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource`

NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSourceWithDefaults instantiates a new NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetCidr

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetVlan

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetHost

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource) GetHost() int32`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource) GetHostOk() (*int32, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource) SetHost(v int32)`

SetHost sets Host field to given value.

### HasHost

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource) HasHost() bool`

HasHost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


