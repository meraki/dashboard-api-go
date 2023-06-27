# GetOrganizationAdaptivePolicyOverview200ResponseCounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to **int32** | Number of adaptive policy groups currently in the organization. | [optional] 
**CustomGroups** | Pointer to **int32** | Number of user-created adaptive policy groups currently in the organization. | [optional] 
**CustomAcls** | Pointer to **int32** | Number of user-created adaptive policy ACLs currently in the organization. | [optional] 
**Policies** | Pointer to **int32** | Number of adaptive policies currently in the organization. | [optional] 
**DenyPolicies** | Pointer to **int32** | Number of adaptive policies currently in the organization that deny all traffic. | [optional] 
**AllowPolicies** | Pointer to **int32** | Number of adaptive policies currently in the organization that allow all traffic. | [optional] 
**PolicyObjects** | Pointer to **int32** | Number of policy objects (with the adaptive policy type) currently in the organization. | [optional] 

## Methods

### NewGetOrganizationAdaptivePolicyOverview200ResponseCounts

`func NewGetOrganizationAdaptivePolicyOverview200ResponseCounts() *GetOrganizationAdaptivePolicyOverview200ResponseCounts`

NewGetOrganizationAdaptivePolicyOverview200ResponseCounts instantiates a new GetOrganizationAdaptivePolicyOverview200ResponseCounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationAdaptivePolicyOverview200ResponseCountsWithDefaults

`func NewGetOrganizationAdaptivePolicyOverview200ResponseCountsWithDefaults() *GetOrganizationAdaptivePolicyOverview200ResponseCounts`

NewGetOrganizationAdaptivePolicyOverview200ResponseCountsWithDefaults instantiates a new GetOrganizationAdaptivePolicyOverview200ResponseCounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) GetGroups() int32`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) GetGroupsOk() (*int32, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) SetGroups(v int32)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetCustomGroups

`func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) GetCustomGroups() int32`

GetCustomGroups returns the CustomGroups field if non-nil, zero value otherwise.

### GetCustomGroupsOk

`func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) GetCustomGroupsOk() (*int32, bool)`

GetCustomGroupsOk returns a tuple with the CustomGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomGroups

`func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) SetCustomGroups(v int32)`

SetCustomGroups sets CustomGroups field to given value.

### HasCustomGroups

`func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) HasCustomGroups() bool`

HasCustomGroups returns a boolean if a field has been set.

### GetCustomAcls

`func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) GetCustomAcls() int32`

GetCustomAcls returns the CustomAcls field if non-nil, zero value otherwise.

### GetCustomAclsOk

`func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) GetCustomAclsOk() (*int32, bool)`

GetCustomAclsOk returns a tuple with the CustomAcls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAcls

`func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) SetCustomAcls(v int32)`

SetCustomAcls sets CustomAcls field to given value.

### HasCustomAcls

`func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) HasCustomAcls() bool`

HasCustomAcls returns a boolean if a field has been set.

### GetPolicies

`func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) GetPolicies() int32`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) GetPoliciesOk() (*int32, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) SetPolicies(v int32)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetDenyPolicies

`func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) GetDenyPolicies() int32`

GetDenyPolicies returns the DenyPolicies field if non-nil, zero value otherwise.

### GetDenyPoliciesOk

`func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) GetDenyPoliciesOk() (*int32, bool)`

GetDenyPoliciesOk returns a tuple with the DenyPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyPolicies

`func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) SetDenyPolicies(v int32)`

SetDenyPolicies sets DenyPolicies field to given value.

### HasDenyPolicies

`func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) HasDenyPolicies() bool`

HasDenyPolicies returns a boolean if a field has been set.

### GetAllowPolicies

`func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) GetAllowPolicies() int32`

GetAllowPolicies returns the AllowPolicies field if non-nil, zero value otherwise.

### GetAllowPoliciesOk

`func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) GetAllowPoliciesOk() (*int32, bool)`

GetAllowPoliciesOk returns a tuple with the AllowPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPolicies

`func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) SetAllowPolicies(v int32)`

SetAllowPolicies sets AllowPolicies field to given value.

### HasAllowPolicies

`func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) HasAllowPolicies() bool`

HasAllowPolicies returns a boolean if a field has been set.

### GetPolicyObjects

`func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) GetPolicyObjects() int32`

GetPolicyObjects returns the PolicyObjects field if non-nil, zero value otherwise.

### GetPolicyObjectsOk

`func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) GetPolicyObjectsOk() (*int32, bool)`

GetPolicyObjectsOk returns a tuple with the PolicyObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyObjects

`func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) SetPolicyObjects(v int32)`

SetPolicyObjects sets PolicyObjects field to given value.

### HasPolicyObjects

`func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) HasPolicyObjects() bool`

HasPolicyObjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


