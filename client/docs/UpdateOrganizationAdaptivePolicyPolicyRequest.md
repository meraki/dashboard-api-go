# UpdateOrganizationAdaptivePolicyPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceGroup** | Pointer to [**CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup**](CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup.md) |  | [optional] 
**DestinationGroup** | Pointer to [**CreateOrganizationAdaptivePolicyPolicyRequestDestinationGroup**](CreateOrganizationAdaptivePolicyPolicyRequestDestinationGroup.md) |  | [optional] 
**Acls** | Pointer to [**[]CreateOrganizationAdaptivePolicyPolicyRequestAclsInner**](CreateOrganizationAdaptivePolicyPolicyRequestAclsInner.md) | An ordered array of adaptive policy ACLs (each requires one unique attribute) that apply to this policy  | [optional] 
**LastEntryRule** | Pointer to **string** | The rule to apply if there is no matching ACL  | [optional] 

## Methods

### NewUpdateOrganizationAdaptivePolicyPolicyRequest

`func NewUpdateOrganizationAdaptivePolicyPolicyRequest() *UpdateOrganizationAdaptivePolicyPolicyRequest`

NewUpdateOrganizationAdaptivePolicyPolicyRequest instantiates a new UpdateOrganizationAdaptivePolicyPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrganizationAdaptivePolicyPolicyRequestWithDefaults

`func NewUpdateOrganizationAdaptivePolicyPolicyRequestWithDefaults() *UpdateOrganizationAdaptivePolicyPolicyRequest`

NewUpdateOrganizationAdaptivePolicyPolicyRequestWithDefaults instantiates a new UpdateOrganizationAdaptivePolicyPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceGroup

`func (o *UpdateOrganizationAdaptivePolicyPolicyRequest) GetSourceGroup() CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup`

GetSourceGroup returns the SourceGroup field if non-nil, zero value otherwise.

### GetSourceGroupOk

`func (o *UpdateOrganizationAdaptivePolicyPolicyRequest) GetSourceGroupOk() (*CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup, bool)`

GetSourceGroupOk returns a tuple with the SourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGroup

`func (o *UpdateOrganizationAdaptivePolicyPolicyRequest) SetSourceGroup(v CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup)`

SetSourceGroup sets SourceGroup field to given value.

### HasSourceGroup

`func (o *UpdateOrganizationAdaptivePolicyPolicyRequest) HasSourceGroup() bool`

HasSourceGroup returns a boolean if a field has been set.

### GetDestinationGroup

`func (o *UpdateOrganizationAdaptivePolicyPolicyRequest) GetDestinationGroup() CreateOrganizationAdaptivePolicyPolicyRequestDestinationGroup`

GetDestinationGroup returns the DestinationGroup field if non-nil, zero value otherwise.

### GetDestinationGroupOk

`func (o *UpdateOrganizationAdaptivePolicyPolicyRequest) GetDestinationGroupOk() (*CreateOrganizationAdaptivePolicyPolicyRequestDestinationGroup, bool)`

GetDestinationGroupOk returns a tuple with the DestinationGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationGroup

`func (o *UpdateOrganizationAdaptivePolicyPolicyRequest) SetDestinationGroup(v CreateOrganizationAdaptivePolicyPolicyRequestDestinationGroup)`

SetDestinationGroup sets DestinationGroup field to given value.

### HasDestinationGroup

`func (o *UpdateOrganizationAdaptivePolicyPolicyRequest) HasDestinationGroup() bool`

HasDestinationGroup returns a boolean if a field has been set.

### GetAcls

`func (o *UpdateOrganizationAdaptivePolicyPolicyRequest) GetAcls() []CreateOrganizationAdaptivePolicyPolicyRequestAclsInner`

GetAcls returns the Acls field if non-nil, zero value otherwise.

### GetAclsOk

`func (o *UpdateOrganizationAdaptivePolicyPolicyRequest) GetAclsOk() (*[]CreateOrganizationAdaptivePolicyPolicyRequestAclsInner, bool)`

GetAclsOk returns a tuple with the Acls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcls

`func (o *UpdateOrganizationAdaptivePolicyPolicyRequest) SetAcls(v []CreateOrganizationAdaptivePolicyPolicyRequestAclsInner)`

SetAcls sets Acls field to given value.

### HasAcls

`func (o *UpdateOrganizationAdaptivePolicyPolicyRequest) HasAcls() bool`

HasAcls returns a boolean if a field has been set.

### GetLastEntryRule

`func (o *UpdateOrganizationAdaptivePolicyPolicyRequest) GetLastEntryRule() string`

GetLastEntryRule returns the LastEntryRule field if non-nil, zero value otherwise.

### GetLastEntryRuleOk

`func (o *UpdateOrganizationAdaptivePolicyPolicyRequest) GetLastEntryRuleOk() (*string, bool)`

GetLastEntryRuleOk returns a tuple with the LastEntryRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEntryRule

`func (o *UpdateOrganizationAdaptivePolicyPolicyRequest) SetLastEntryRule(v string)`

SetLastEntryRule sets LastEntryRule field to given value.

### HasLastEntryRule

`func (o *UpdateOrganizationAdaptivePolicyPolicyRequest) HasLastEntryRule() bool`

HasLastEntryRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


