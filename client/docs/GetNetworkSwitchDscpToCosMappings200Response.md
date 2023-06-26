# GetNetworkSwitchDscpToCosMappings200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mappings** | Pointer to [**[]GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner**](GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner.md) | An array of DSCP to CoS mappings. An empty array will reset the mappings to default. | [optional] 

## Methods

### NewGetNetworkSwitchDscpToCosMappings200Response

`func NewGetNetworkSwitchDscpToCosMappings200Response() *GetNetworkSwitchDscpToCosMappings200Response`

NewGetNetworkSwitchDscpToCosMappings200Response instantiates a new GetNetworkSwitchDscpToCosMappings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSwitchDscpToCosMappings200ResponseWithDefaults

`func NewGetNetworkSwitchDscpToCosMappings200ResponseWithDefaults() *GetNetworkSwitchDscpToCosMappings200Response`

NewGetNetworkSwitchDscpToCosMappings200ResponseWithDefaults instantiates a new GetNetworkSwitchDscpToCosMappings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMappings

`func (o *GetNetworkSwitchDscpToCosMappings200Response) GetMappings() []GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *GetNetworkSwitchDscpToCosMappings200Response) GetMappingsOk() (*[]GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *GetNetworkSwitchDscpToCosMappings200Response) SetMappings(v []GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *GetNetworkSwitchDscpToCosMappings200Response) HasMappings() bool`

HasMappings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


