# CreateNetworkSwitchLinkAggregationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SwitchPorts** | Pointer to [**[]CreateNetworkSwitchLinkAggregationRequestSwitchPortsInner**](CreateNetworkSwitchLinkAggregationRequestSwitchPortsInner.md) | Array of switch or stack ports for creating aggregation group. Minimum 2 and maximum 8 ports are supported. | [optional] 
**SwitchProfilePorts** | Pointer to [**[]CreateNetworkSwitchLinkAggregationRequestSwitchProfilePortsInner**](CreateNetworkSwitchLinkAggregationRequestSwitchProfilePortsInner.md) | Array of switch profile ports for creating aggregation group. Minimum 2 and maximum 8 ports are supported. | [optional] 

## Methods

### NewCreateNetworkSwitchLinkAggregationRequest

`func NewCreateNetworkSwitchLinkAggregationRequest() *CreateNetworkSwitchLinkAggregationRequest`

NewCreateNetworkSwitchLinkAggregationRequest instantiates a new CreateNetworkSwitchLinkAggregationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkSwitchLinkAggregationRequestWithDefaults

`func NewCreateNetworkSwitchLinkAggregationRequestWithDefaults() *CreateNetworkSwitchLinkAggregationRequest`

NewCreateNetworkSwitchLinkAggregationRequestWithDefaults instantiates a new CreateNetworkSwitchLinkAggregationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSwitchPorts

`func (o *CreateNetworkSwitchLinkAggregationRequest) GetSwitchPorts() []CreateNetworkSwitchLinkAggregationRequestSwitchPortsInner`

GetSwitchPorts returns the SwitchPorts field if non-nil, zero value otherwise.

### GetSwitchPortsOk

`func (o *CreateNetworkSwitchLinkAggregationRequest) GetSwitchPortsOk() (*[]CreateNetworkSwitchLinkAggregationRequestSwitchPortsInner, bool)`

GetSwitchPortsOk returns a tuple with the SwitchPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchPorts

`func (o *CreateNetworkSwitchLinkAggregationRequest) SetSwitchPorts(v []CreateNetworkSwitchLinkAggregationRequestSwitchPortsInner)`

SetSwitchPorts sets SwitchPorts field to given value.

### HasSwitchPorts

`func (o *CreateNetworkSwitchLinkAggregationRequest) HasSwitchPorts() bool`

HasSwitchPorts returns a boolean if a field has been set.

### GetSwitchProfilePorts

`func (o *CreateNetworkSwitchLinkAggregationRequest) GetSwitchProfilePorts() []CreateNetworkSwitchLinkAggregationRequestSwitchProfilePortsInner`

GetSwitchProfilePorts returns the SwitchProfilePorts field if non-nil, zero value otherwise.

### GetSwitchProfilePortsOk

`func (o *CreateNetworkSwitchLinkAggregationRequest) GetSwitchProfilePortsOk() (*[]CreateNetworkSwitchLinkAggregationRequestSwitchProfilePortsInner, bool)`

GetSwitchProfilePortsOk returns a tuple with the SwitchProfilePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchProfilePorts

`func (o *CreateNetworkSwitchLinkAggregationRequest) SetSwitchProfilePorts(v []CreateNetworkSwitchLinkAggregationRequestSwitchProfilePortsInner)`

SetSwitchProfilePorts sets SwitchProfilePorts field to given value.

### HasSwitchProfilePorts

`func (o *CreateNetworkSwitchLinkAggregationRequest) HasSwitchProfilePorts() bool`

HasSwitchProfilePorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


