# InlineObject178

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceGroup** | [**OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup**](OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup.md) |  | 
**DestinationGroup** | [**OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup**](OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup.md) |  | 
**Acls** | Pointer to [**[]OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls**](OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls.md) | An ordered array of adaptive policy ACLs (each requires one unique attribute) that apply to this policy (default: []) | [optional] 
**LastEntryRule** | Pointer to **string** | The rule to apply if there is no matching ACL (default: \&quot;default\&quot;) | [optional] 

## Methods

### NewInlineObject178

`func NewInlineObject178(sourceGroup OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup, destinationGroup OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup, ) *InlineObject178`

NewInlineObject178 instantiates a new InlineObject178 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject178WithDefaults

`func NewInlineObject178WithDefaults() *InlineObject178`

NewInlineObject178WithDefaults instantiates a new InlineObject178 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceGroup

`func (o *InlineObject178) GetSourceGroup() OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup`

GetSourceGroup returns the SourceGroup field if non-nil, zero value otherwise.

### GetSourceGroupOk

`func (o *InlineObject178) GetSourceGroupOk() (*OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup, bool)`

GetSourceGroupOk returns a tuple with the SourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGroup

`func (o *InlineObject178) SetSourceGroup(v OrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup)`

SetSourceGroup sets SourceGroup field to given value.


### GetDestinationGroup

`func (o *InlineObject178) GetDestinationGroup() OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup`

GetDestinationGroup returns the DestinationGroup field if non-nil, zero value otherwise.

### GetDestinationGroupOk

`func (o *InlineObject178) GetDestinationGroupOk() (*OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup, bool)`

GetDestinationGroupOk returns a tuple with the DestinationGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationGroup

`func (o *InlineObject178) SetDestinationGroup(v OrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup)`

SetDestinationGroup sets DestinationGroup field to given value.


### GetAcls

`func (o *InlineObject178) GetAcls() []OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls`

GetAcls returns the Acls field if non-nil, zero value otherwise.

### GetAclsOk

`func (o *InlineObject178) GetAclsOk() (*[]OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls, bool)`

GetAclsOk returns a tuple with the Acls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcls

`func (o *InlineObject178) SetAcls(v []OrganizationsOrganizationIdAdaptivePolicyPoliciesAcls)`

SetAcls sets Acls field to given value.

### HasAcls

`func (o *InlineObject178) HasAcls() bool`

HasAcls returns a boolean if a field has been set.

### GetLastEntryRule

`func (o *InlineObject178) GetLastEntryRule() string`

GetLastEntryRule returns the LastEntryRule field if non-nil, zero value otherwise.

### GetLastEntryRuleOk

`func (o *InlineObject178) GetLastEntryRuleOk() (*string, bool)`

GetLastEntryRuleOk returns a tuple with the LastEntryRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEntryRule

`func (o *InlineObject178) SetLastEntryRule(v string)`

SetLastEntryRule sets LastEntryRule field to given value.

### HasLastEntryRule

`func (o *InlineObject178) HasLastEntryRule() bool`

HasLastEntryRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


