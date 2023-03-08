# InlineResponse20086Counts

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

### NewInlineResponse20086Counts

`func NewInlineResponse20086Counts() *InlineResponse20086Counts`

NewInlineResponse20086Counts instantiates a new InlineResponse20086Counts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20086CountsWithDefaults

`func NewInlineResponse20086CountsWithDefaults() *InlineResponse20086Counts`

NewInlineResponse20086CountsWithDefaults instantiates a new InlineResponse20086Counts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *InlineResponse20086Counts) GetGroups() int32`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *InlineResponse20086Counts) GetGroupsOk() (*int32, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *InlineResponse20086Counts) SetGroups(v int32)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *InlineResponse20086Counts) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetCustomGroups

`func (o *InlineResponse20086Counts) GetCustomGroups() int32`

GetCustomGroups returns the CustomGroups field if non-nil, zero value otherwise.

### GetCustomGroupsOk

`func (o *InlineResponse20086Counts) GetCustomGroupsOk() (*int32, bool)`

GetCustomGroupsOk returns a tuple with the CustomGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomGroups

`func (o *InlineResponse20086Counts) SetCustomGroups(v int32)`

SetCustomGroups sets CustomGroups field to given value.

### HasCustomGroups

`func (o *InlineResponse20086Counts) HasCustomGroups() bool`

HasCustomGroups returns a boolean if a field has been set.

### GetCustomAcls

`func (o *InlineResponse20086Counts) GetCustomAcls() int32`

GetCustomAcls returns the CustomAcls field if non-nil, zero value otherwise.

### GetCustomAclsOk

`func (o *InlineResponse20086Counts) GetCustomAclsOk() (*int32, bool)`

GetCustomAclsOk returns a tuple with the CustomAcls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAcls

`func (o *InlineResponse20086Counts) SetCustomAcls(v int32)`

SetCustomAcls sets CustomAcls field to given value.

### HasCustomAcls

`func (o *InlineResponse20086Counts) HasCustomAcls() bool`

HasCustomAcls returns a boolean if a field has been set.

### GetPolicies

`func (o *InlineResponse20086Counts) GetPolicies() int32`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *InlineResponse20086Counts) GetPoliciesOk() (*int32, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *InlineResponse20086Counts) SetPolicies(v int32)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *InlineResponse20086Counts) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetDenyPolicies

`func (o *InlineResponse20086Counts) GetDenyPolicies() int32`

GetDenyPolicies returns the DenyPolicies field if non-nil, zero value otherwise.

### GetDenyPoliciesOk

`func (o *InlineResponse20086Counts) GetDenyPoliciesOk() (*int32, bool)`

GetDenyPoliciesOk returns a tuple with the DenyPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyPolicies

`func (o *InlineResponse20086Counts) SetDenyPolicies(v int32)`

SetDenyPolicies sets DenyPolicies field to given value.

### HasDenyPolicies

`func (o *InlineResponse20086Counts) HasDenyPolicies() bool`

HasDenyPolicies returns a boolean if a field has been set.

### GetAllowPolicies

`func (o *InlineResponse20086Counts) GetAllowPolicies() int32`

GetAllowPolicies returns the AllowPolicies field if non-nil, zero value otherwise.

### GetAllowPoliciesOk

`func (o *InlineResponse20086Counts) GetAllowPoliciesOk() (*int32, bool)`

GetAllowPoliciesOk returns a tuple with the AllowPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPolicies

`func (o *InlineResponse20086Counts) SetAllowPolicies(v int32)`

SetAllowPolicies sets AllowPolicies field to given value.

### HasAllowPolicies

`func (o *InlineResponse20086Counts) HasAllowPolicies() bool`

HasAllowPolicies returns a boolean if a field has been set.

### GetPolicyObjects

`func (o *InlineResponse20086Counts) GetPolicyObjects() int32`

GetPolicyObjects returns the PolicyObjects field if non-nil, zero value otherwise.

### GetPolicyObjectsOk

`func (o *InlineResponse20086Counts) GetPolicyObjectsOk() (*int32, bool)`

GetPolicyObjectsOk returns a tuple with the PolicyObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyObjects

`func (o *InlineResponse20086Counts) SetPolicyObjects(v int32)`

SetPolicyObjects sets PolicyObjects field to given value.

### HasPolicyObjects

`func (o *InlineResponse20086Counts) HasPolicyObjects() bool`

HasPolicyObjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


