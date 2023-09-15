# UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | **string** | Protocol. | 
**Destination** | Pointer to **string** | Destination address; hostname required for DNS, IPv4 otherwise. | [optional] 
**Port** | Pointer to **string** | Destination port. | [optional] 

## Methods

### NewUpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner

`func NewUpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner(protocol string, ) *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner`

NewUpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner instantiates a new UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInnerWithDefaults

`func NewUpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInnerWithDefaults() *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner`

NewUpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInnerWithDefaults instantiates a new UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetDestination

`func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetPort

`func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


