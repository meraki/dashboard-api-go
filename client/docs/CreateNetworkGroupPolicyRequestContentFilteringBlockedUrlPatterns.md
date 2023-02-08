# CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlPatterns

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Settings** | Pointer to **string** | How URL patterns are applied. Can be &#39;network default&#39;, &#39;append&#39; or &#39;override&#39;. | [optional] 
**Patterns** | Pointer to **[]string** | A list of URL patterns that are blocked | [optional] 

## Methods

### NewCreateNetworkGroupPolicyRequestContentFilteringBlockedUrlPatterns

`func NewCreateNetworkGroupPolicyRequestContentFilteringBlockedUrlPatterns() *CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlPatterns`

NewCreateNetworkGroupPolicyRequestContentFilteringBlockedUrlPatterns instantiates a new CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlPatterns object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkGroupPolicyRequestContentFilteringBlockedUrlPatternsWithDefaults

`func NewCreateNetworkGroupPolicyRequestContentFilteringBlockedUrlPatternsWithDefaults() *CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlPatterns`

NewCreateNetworkGroupPolicyRequestContentFilteringBlockedUrlPatternsWithDefaults instantiates a new CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlPatterns object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettings

`func (o *CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlPatterns) GetSettings() string`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlPatterns) GetSettingsOk() (*string, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlPatterns) SetSettings(v string)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlPatterns) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetPatterns

`func (o *CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlPatterns) GetPatterns() []string`

GetPatterns returns the Patterns field if non-nil, zero value otherwise.

### GetPatternsOk

`func (o *CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlPatterns) GetPatternsOk() (*[]string, bool)`

GetPatternsOk returns a tuple with the Patterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatterns

`func (o *CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlPatterns) SetPatterns(v []string)`

SetPatterns sets Patterns field to given value.

### HasPatterns

`func (o *CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlPatterns) HasPatterns() bool`

HasPatterns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


