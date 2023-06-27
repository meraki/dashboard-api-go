# CreateNetworkGroupPolicyRequestBandwidth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Settings** | Pointer to **string** | How bandwidth limits are enforced. Can be &#39;network default&#39;, &#39;ignore&#39; or &#39;custom&#39;. | [optional] 
**BandwidthLimits** | Pointer to [**CreateNetworkGroupPolicyRequestBandwidthBandwidthLimits**](CreateNetworkGroupPolicyRequestBandwidthBandwidthLimits.md) |  | [optional] 

## Methods

### NewCreateNetworkGroupPolicyRequestBandwidth

`func NewCreateNetworkGroupPolicyRequestBandwidth() *CreateNetworkGroupPolicyRequestBandwidth`

NewCreateNetworkGroupPolicyRequestBandwidth instantiates a new CreateNetworkGroupPolicyRequestBandwidth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkGroupPolicyRequestBandwidthWithDefaults

`func NewCreateNetworkGroupPolicyRequestBandwidthWithDefaults() *CreateNetworkGroupPolicyRequestBandwidth`

NewCreateNetworkGroupPolicyRequestBandwidthWithDefaults instantiates a new CreateNetworkGroupPolicyRequestBandwidth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettings

`func (o *CreateNetworkGroupPolicyRequestBandwidth) GetSettings() string`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *CreateNetworkGroupPolicyRequestBandwidth) GetSettingsOk() (*string, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *CreateNetworkGroupPolicyRequestBandwidth) SetSettings(v string)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *CreateNetworkGroupPolicyRequestBandwidth) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetBandwidthLimits

`func (o *CreateNetworkGroupPolicyRequestBandwidth) GetBandwidthLimits() CreateNetworkGroupPolicyRequestBandwidthBandwidthLimits`

GetBandwidthLimits returns the BandwidthLimits field if non-nil, zero value otherwise.

### GetBandwidthLimitsOk

`func (o *CreateNetworkGroupPolicyRequestBandwidth) GetBandwidthLimitsOk() (*CreateNetworkGroupPolicyRequestBandwidthBandwidthLimits, bool)`

GetBandwidthLimitsOk returns a tuple with the BandwidthLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthLimits

`func (o *CreateNetworkGroupPolicyRequestBandwidth) SetBandwidthLimits(v CreateNetworkGroupPolicyRequestBandwidthBandwidthLimits)`

SetBandwidthLimits sets BandwidthLimits field to given value.

### HasBandwidthLimits

`func (o *CreateNetworkGroupPolicyRequestBandwidth) HasBandwidthLimits() bool`

HasBandwidthLimits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


