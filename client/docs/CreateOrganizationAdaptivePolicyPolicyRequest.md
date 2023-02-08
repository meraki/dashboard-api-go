# CreateOrganizationAdaptivePolicyPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceGroup** | [**CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup**](CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup.md) |  | 
**DestinationGroup** | [**CreateOrganizationAdaptivePolicyPolicyRequestDestinationGroup**](CreateOrganizationAdaptivePolicyPolicyRequestDestinationGroup.md) |  | 
**Acls** | Pointer to [**[]CreateOrganizationAdaptivePolicyPolicyRequestAclsInner**](CreateOrganizationAdaptivePolicyPolicyRequestAclsInner.md) | An ordered array of adaptive policy ACLs (each requires one unique attribute) that apply to this policy (default: [])  | [optional] 
**LastEntryRule** | Pointer to **string** | The rule to apply if there is no matching ACL (default: \&quot;default\&quot;)  | [optional] 

## Methods

### NewCreateOrganizationAdaptivePolicyPolicyRequest

`func NewCreateOrganizationAdaptivePolicyPolicyRequest(sourceGroup CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup, destinationGroup CreateOrganizationAdaptivePolicyPolicyRequestDestinationGroup, ) *CreateOrganizationAdaptivePolicyPolicyRequest`

NewCreateOrganizationAdaptivePolicyPolicyRequest instantiates a new CreateOrganizationAdaptivePolicyPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationAdaptivePolicyPolicyRequestWithDefaults

`func NewCreateOrganizationAdaptivePolicyPolicyRequestWithDefaults() *CreateOrganizationAdaptivePolicyPolicyRequest`

NewCreateOrganizationAdaptivePolicyPolicyRequestWithDefaults instantiates a new CreateOrganizationAdaptivePolicyPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceGroup

`func (o *CreateOrganizationAdaptivePolicyPolicyRequest) GetSourceGroup() CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup`

GetSourceGroup returns the SourceGroup field if non-nil, zero value otherwise.

### GetSourceGroupOk

`func (o *CreateOrganizationAdaptivePolicyPolicyRequest) GetSourceGroupOk() (*CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup, bool)`

GetSourceGroupOk returns a tuple with the SourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGroup

`func (o *CreateOrganizationAdaptivePolicyPolicyRequest) SetSourceGroup(v CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup)`

SetSourceGroup sets SourceGroup field to given value.


### GetDestinationGroup

`func (o *CreateOrganizationAdaptivePolicyPolicyRequest) GetDestinationGroup() CreateOrganizationAdaptivePolicyPolicyRequestDestinationGroup`

GetDestinationGroup returns the DestinationGroup field if non-nil, zero value otherwise.

### GetDestinationGroupOk

`func (o *CreateOrganizationAdaptivePolicyPolicyRequest) GetDestinationGroupOk() (*CreateOrganizationAdaptivePolicyPolicyRequestDestinationGroup, bool)`

GetDestinationGroupOk returns a tuple with the DestinationGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationGroup

`func (o *CreateOrganizationAdaptivePolicyPolicyRequest) SetDestinationGroup(v CreateOrganizationAdaptivePolicyPolicyRequestDestinationGroup)`

SetDestinationGroup sets DestinationGroup field to given value.


### GetAcls

`func (o *CreateOrganizationAdaptivePolicyPolicyRequest) GetAcls() []CreateOrganizationAdaptivePolicyPolicyRequestAclsInner`

GetAcls returns the Acls field if non-nil, zero value otherwise.

### GetAclsOk

`func (o *CreateOrganizationAdaptivePolicyPolicyRequest) GetAclsOk() (*[]CreateOrganizationAdaptivePolicyPolicyRequestAclsInner, bool)`

GetAclsOk returns a tuple with the Acls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcls

`func (o *CreateOrganizationAdaptivePolicyPolicyRequest) SetAcls(v []CreateOrganizationAdaptivePolicyPolicyRequestAclsInner)`

SetAcls sets Acls field to given value.

### HasAcls

`func (o *CreateOrganizationAdaptivePolicyPolicyRequest) HasAcls() bool`

HasAcls returns a boolean if a field has been set.

### GetLastEntryRule

`func (o *CreateOrganizationAdaptivePolicyPolicyRequest) GetLastEntryRule() string`

GetLastEntryRule returns the LastEntryRule field if non-nil, zero value otherwise.

### GetLastEntryRuleOk

`func (o *CreateOrganizationAdaptivePolicyPolicyRequest) GetLastEntryRuleOk() (*string, bool)`

GetLastEntryRuleOk returns a tuple with the LastEntryRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEntryRule

`func (o *CreateOrganizationAdaptivePolicyPolicyRequest) SetLastEntryRule(v string)`

SetLastEntryRule sets LastEntryRule field to given value.

### HasLastEntryRule

`func (o *CreateOrganizationAdaptivePolicyPolicyRequest) HasLastEntryRule() bool`

HasLastEntryRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


