# CreateNetworkSwitchRoutingMulticastRendezvousPointRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterfaceIp** | **string** | The IP address of the interface where the RP needs to be created. | 
**MulticastGroup** | **string** | &#39;Any&#39;, or the IP address of a multicast group | 

## Methods

### NewCreateNetworkSwitchRoutingMulticastRendezvousPointRequest

`func NewCreateNetworkSwitchRoutingMulticastRendezvousPointRequest(interfaceIp string, multicastGroup string, ) *CreateNetworkSwitchRoutingMulticastRendezvousPointRequest`

NewCreateNetworkSwitchRoutingMulticastRendezvousPointRequest instantiates a new CreateNetworkSwitchRoutingMulticastRendezvousPointRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkSwitchRoutingMulticastRendezvousPointRequestWithDefaults

`func NewCreateNetworkSwitchRoutingMulticastRendezvousPointRequestWithDefaults() *CreateNetworkSwitchRoutingMulticastRendezvousPointRequest`

NewCreateNetworkSwitchRoutingMulticastRendezvousPointRequestWithDefaults instantiates a new CreateNetworkSwitchRoutingMulticastRendezvousPointRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterfaceIp

`func (o *CreateNetworkSwitchRoutingMulticastRendezvousPointRequest) GetInterfaceIp() string`

GetInterfaceIp returns the InterfaceIp field if non-nil, zero value otherwise.

### GetInterfaceIpOk

`func (o *CreateNetworkSwitchRoutingMulticastRendezvousPointRequest) GetInterfaceIpOk() (*string, bool)`

GetInterfaceIpOk returns a tuple with the InterfaceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceIp

`func (o *CreateNetworkSwitchRoutingMulticastRendezvousPointRequest) SetInterfaceIp(v string)`

SetInterfaceIp sets InterfaceIp field to given value.


### GetMulticastGroup

`func (o *CreateNetworkSwitchRoutingMulticastRendezvousPointRequest) GetMulticastGroup() string`

GetMulticastGroup returns the MulticastGroup field if non-nil, zero value otherwise.

### GetMulticastGroupOk

`func (o *CreateNetworkSwitchRoutingMulticastRendezvousPointRequest) GetMulticastGroupOk() (*string, bool)`

GetMulticastGroupOk returns a tuple with the MulticastGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticastGroup

`func (o *CreateNetworkSwitchRoutingMulticastRendezvousPointRequest) SetMulticastGroup(v string)`

SetMulticastGroup sets MulticastGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


