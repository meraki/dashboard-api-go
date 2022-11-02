# CreateNetworkGroupPolicyRequestBonjourForwarding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Settings** | Pointer to **string** | How Bonjour rules are applied. Can be &#39;network default&#39;, &#39;ignore&#39; or &#39;custom&#39;. | [optional] 
**Rules** | Pointer to [**[]CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner**](CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner.md) | A list of the Bonjour forwarding rules for your group policy. If &#39;settings&#39; is set to &#39;custom&#39;, at least one rule must be specified. | [optional] 

## Methods

### NewCreateNetworkGroupPolicyRequestBonjourForwarding

`func NewCreateNetworkGroupPolicyRequestBonjourForwarding() *CreateNetworkGroupPolicyRequestBonjourForwarding`

NewCreateNetworkGroupPolicyRequestBonjourForwarding instantiates a new CreateNetworkGroupPolicyRequestBonjourForwarding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkGroupPolicyRequestBonjourForwardingWithDefaults

`func NewCreateNetworkGroupPolicyRequestBonjourForwardingWithDefaults() *CreateNetworkGroupPolicyRequestBonjourForwarding`

NewCreateNetworkGroupPolicyRequestBonjourForwardingWithDefaults instantiates a new CreateNetworkGroupPolicyRequestBonjourForwarding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettings

`func (o *CreateNetworkGroupPolicyRequestBonjourForwarding) GetSettings() string`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *CreateNetworkGroupPolicyRequestBonjourForwarding) GetSettingsOk() (*string, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *CreateNetworkGroupPolicyRequestBonjourForwarding) SetSettings(v string)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *CreateNetworkGroupPolicyRequestBonjourForwarding) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetRules

`func (o *CreateNetworkGroupPolicyRequestBonjourForwarding) GetRules() []CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *CreateNetworkGroupPolicyRequestBonjourForwarding) GetRulesOk() (*[]CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *CreateNetworkGroupPolicyRequestBonjourForwarding) SetRules(v []CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner)`

SetRules sets Rules field to given value.

### HasRules

`func (o *CreateNetworkGroupPolicyRequestBonjourForwarding) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


