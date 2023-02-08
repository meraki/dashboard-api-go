# UpdateNetworkSwitchDscpToCosMappingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mappings** | [**[]UpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner**](UpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner.md) | An array of DSCP to CoS mappings. An empty array will reset the mappings to default. | 

## Methods

### NewUpdateNetworkSwitchDscpToCosMappingsRequest

`func NewUpdateNetworkSwitchDscpToCosMappingsRequest(mappings []UpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner, ) *UpdateNetworkSwitchDscpToCosMappingsRequest`

NewUpdateNetworkSwitchDscpToCosMappingsRequest instantiates a new UpdateNetworkSwitchDscpToCosMappingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkSwitchDscpToCosMappingsRequestWithDefaults

`func NewUpdateNetworkSwitchDscpToCosMappingsRequestWithDefaults() *UpdateNetworkSwitchDscpToCosMappingsRequest`

NewUpdateNetworkSwitchDscpToCosMappingsRequestWithDefaults instantiates a new UpdateNetworkSwitchDscpToCosMappingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMappings

`func (o *UpdateNetworkSwitchDscpToCosMappingsRequest) GetMappings() []UpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *UpdateNetworkSwitchDscpToCosMappingsRequest) GetMappingsOk() (*[]UpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *UpdateNetworkSwitchDscpToCosMappingsRequest) SetMappings(v []UpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner)`

SetMappings sets Mappings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


