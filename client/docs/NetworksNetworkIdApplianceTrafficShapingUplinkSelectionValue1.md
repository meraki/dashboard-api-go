# NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of this applicationCategory or application type traffic filter. E.g.: \&quot;meraki:layer7/category/1\&quot;, \&quot;meraki:layer7/application/4\&quot; | [optional] 
**Protocol** | Pointer to **string** | Protocol of this custom type traffic filter. Must be one of: &#39;tcp&#39;, &#39;udp&#39;, &#39;icmp&#39;, &#39;icmp6&#39; or &#39;any&#39; | [optional] 
**Source** | Pointer to [**NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Source**](NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Source.md) |  | [optional] 
**Destination** | Pointer to [**NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination**](NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination.md) |  | [optional] 

## Methods

### NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1

`func NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1() *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1`

NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1 instantiates a new NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1WithDefaults

`func NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1WithDefaults() *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1`

NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1WithDefaults instantiates a new NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProtocol

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSource

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) GetSource() NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Source`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) GetSourceOk() (*NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Source, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) SetSource(v NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Source)`

SetSource sets Source field to given value.

### HasSource

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetDestination

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) GetDestination() NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) GetDestinationOk() (*NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) SetDestination(v NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) HasDestination() bool`

HasDestination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


