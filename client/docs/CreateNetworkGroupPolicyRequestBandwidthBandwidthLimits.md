# CreateNetworkGroupPolicyRequestBandwidthBandwidthLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LimitUp** | Pointer to **int32** | The maximum upload limit (integer, in Kbps). null indicates no limit | [optional] 
**LimitDown** | Pointer to **int32** | The maximum download limit (integer, in Kbps). null indicates no limit | [optional] 

## Methods

### NewCreateNetworkGroupPolicyRequestBandwidthBandwidthLimits

`func NewCreateNetworkGroupPolicyRequestBandwidthBandwidthLimits() *CreateNetworkGroupPolicyRequestBandwidthBandwidthLimits`

NewCreateNetworkGroupPolicyRequestBandwidthBandwidthLimits instantiates a new CreateNetworkGroupPolicyRequestBandwidthBandwidthLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkGroupPolicyRequestBandwidthBandwidthLimitsWithDefaults

`func NewCreateNetworkGroupPolicyRequestBandwidthBandwidthLimitsWithDefaults() *CreateNetworkGroupPolicyRequestBandwidthBandwidthLimits`

NewCreateNetworkGroupPolicyRequestBandwidthBandwidthLimitsWithDefaults instantiates a new CreateNetworkGroupPolicyRequestBandwidthBandwidthLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimitUp

`func (o *CreateNetworkGroupPolicyRequestBandwidthBandwidthLimits) GetLimitUp() int32`

GetLimitUp returns the LimitUp field if non-nil, zero value otherwise.

### GetLimitUpOk

`func (o *CreateNetworkGroupPolicyRequestBandwidthBandwidthLimits) GetLimitUpOk() (*int32, bool)`

GetLimitUpOk returns a tuple with the LimitUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitUp

`func (o *CreateNetworkGroupPolicyRequestBandwidthBandwidthLimits) SetLimitUp(v int32)`

SetLimitUp sets LimitUp field to given value.

### HasLimitUp

`func (o *CreateNetworkGroupPolicyRequestBandwidthBandwidthLimits) HasLimitUp() bool`

HasLimitUp returns a boolean if a field has been set.

### GetLimitDown

`func (o *CreateNetworkGroupPolicyRequestBandwidthBandwidthLimits) GetLimitDown() int32`

GetLimitDown returns the LimitDown field if non-nil, zero value otherwise.

### GetLimitDownOk

`func (o *CreateNetworkGroupPolicyRequestBandwidthBandwidthLimits) GetLimitDownOk() (*int32, bool)`

GetLimitDownOk returns a tuple with the LimitDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitDown

`func (o *CreateNetworkGroupPolicyRequestBandwidthBandwidthLimits) SetLimitDown(v int32)`

SetLimitDown sets LimitDown field to given value.

### HasLimitDown

`func (o *CreateNetworkGroupPolicyRequestBandwidthBandwidthLimits) HasLimitDown() bool`

HasLimitDown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


