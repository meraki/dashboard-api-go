# NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | **string** | Profile identifier. | 
**PortId** | **string** | Port identifier of switch port. For modules, the identifier is \&quot;SlotNumber_ModuleType_PortNumber\&quot; (Ex: \&quot;1_8X10G_1\&quot;), otherwise it is just the port number (Ex: \&quot;8\&quot;). | 

## Methods

### NewNetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts

`func NewNetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts(profile string, portId string, ) *NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts`

NewNetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts instantiates a new NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdSwitchLinkAggregationsSwitchProfilePortsWithDefaults

`func NewNetworksNetworkIdSwitchLinkAggregationsSwitchProfilePortsWithDefaults() *NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts`

NewNetworksNetworkIdSwitchLinkAggregationsSwitchProfilePortsWithDefaults instantiates a new NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfile

`func (o *NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetPortId

`func (o *NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts) SetPortId(v string)`

SetPortId sets PortId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


