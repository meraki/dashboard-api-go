# GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dscp** | Pointer to **int32** | The Differentiated Services Code Point (DSCP) tag in the IP header that will be mapped to a particular Class-of-Service (CoS) queue. Value can be in the range of 0 to 63 inclusive. | [optional] 
**Cos** | Pointer to **int32** | The actual layer-2 CoS queue the DSCP value is mapped to. These are not bits set on outgoing frames. Value can be in the range of 0 to 5 inclusive. | [optional] 
**Title** | Pointer to **string** | Label for the mapping (optional). | [optional] 

## Methods

### NewGetNetworkSwitchDscpToCosMappings200ResponseMappingsInner

`func NewGetNetworkSwitchDscpToCosMappings200ResponseMappingsInner() *GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner`

NewGetNetworkSwitchDscpToCosMappings200ResponseMappingsInner instantiates a new GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSwitchDscpToCosMappings200ResponseMappingsInnerWithDefaults

`func NewGetNetworkSwitchDscpToCosMappings200ResponseMappingsInnerWithDefaults() *GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner`

NewGetNetworkSwitchDscpToCosMappings200ResponseMappingsInnerWithDefaults instantiates a new GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDscp

`func (o *GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner) GetDscp() int32`

GetDscp returns the Dscp field if non-nil, zero value otherwise.

### GetDscpOk

`func (o *GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner) GetDscpOk() (*int32, bool)`

GetDscpOk returns a tuple with the Dscp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscp

`func (o *GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner) SetDscp(v int32)`

SetDscp sets Dscp field to given value.

### HasDscp

`func (o *GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner) HasDscp() bool`

HasDscp returns a boolean if a field has been set.

### GetCos

`func (o *GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner) GetCos() int32`

GetCos returns the Cos field if non-nil, zero value otherwise.

### GetCosOk

`func (o *GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner) GetCosOk() (*int32, bool)`

GetCosOk returns a tuple with the Cos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCos

`func (o *GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner) SetCos(v int32)`

SetCos sets Cos field to given value.

### HasCos

`func (o *GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner) HasCos() bool`

HasCos returns a boolean if a field has been set.

### GetTitle

`func (o *GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GetNetworkSwitchDscpToCosMappings200ResponseMappingsInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


