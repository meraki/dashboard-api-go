# InlineResponse20081Limits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomGroups** | Pointer to **int32** | Maximum number of user-created adaptive policy groups allowed in the organization. | [optional] 
**RulesInAnAcl** | Pointer to **int32** | Maximum number of rules allowed in an adaptive policy ACL in the organization. | [optional] 
**AclsInAPolicy** | Pointer to **int32** | Maximum number of adaptive policy ACLs that can be assigned to an adaptive policy in the organization. | [optional] 
**PolicyObjects** | Pointer to **int32** | Maximum number of policy objects (with the adaptive policy type) allowed in the organization. | [optional] 

## Methods

### NewInlineResponse20081Limits

`func NewInlineResponse20081Limits() *InlineResponse20081Limits`

NewInlineResponse20081Limits instantiates a new InlineResponse20081Limits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20081LimitsWithDefaults

`func NewInlineResponse20081LimitsWithDefaults() *InlineResponse20081Limits`

NewInlineResponse20081LimitsWithDefaults instantiates a new InlineResponse20081Limits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomGroups

`func (o *InlineResponse20081Limits) GetCustomGroups() int32`

GetCustomGroups returns the CustomGroups field if non-nil, zero value otherwise.

### GetCustomGroupsOk

`func (o *InlineResponse20081Limits) GetCustomGroupsOk() (*int32, bool)`

GetCustomGroupsOk returns a tuple with the CustomGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomGroups

`func (o *InlineResponse20081Limits) SetCustomGroups(v int32)`

SetCustomGroups sets CustomGroups field to given value.

### HasCustomGroups

`func (o *InlineResponse20081Limits) HasCustomGroups() bool`

HasCustomGroups returns a boolean if a field has been set.

### GetRulesInAnAcl

`func (o *InlineResponse20081Limits) GetRulesInAnAcl() int32`

GetRulesInAnAcl returns the RulesInAnAcl field if non-nil, zero value otherwise.

### GetRulesInAnAclOk

`func (o *InlineResponse20081Limits) GetRulesInAnAclOk() (*int32, bool)`

GetRulesInAnAclOk returns a tuple with the RulesInAnAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesInAnAcl

`func (o *InlineResponse20081Limits) SetRulesInAnAcl(v int32)`

SetRulesInAnAcl sets RulesInAnAcl field to given value.

### HasRulesInAnAcl

`func (o *InlineResponse20081Limits) HasRulesInAnAcl() bool`

HasRulesInAnAcl returns a boolean if a field has been set.

### GetAclsInAPolicy

`func (o *InlineResponse20081Limits) GetAclsInAPolicy() int32`

GetAclsInAPolicy returns the AclsInAPolicy field if non-nil, zero value otherwise.

### GetAclsInAPolicyOk

`func (o *InlineResponse20081Limits) GetAclsInAPolicyOk() (*int32, bool)`

GetAclsInAPolicyOk returns a tuple with the AclsInAPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclsInAPolicy

`func (o *InlineResponse20081Limits) SetAclsInAPolicy(v int32)`

SetAclsInAPolicy sets AclsInAPolicy field to given value.

### HasAclsInAPolicy

`func (o *InlineResponse20081Limits) HasAclsInAPolicy() bool`

HasAclsInAPolicy returns a boolean if a field has been set.

### GetPolicyObjects

`func (o *InlineResponse20081Limits) GetPolicyObjects() int32`

GetPolicyObjects returns the PolicyObjects field if non-nil, zero value otherwise.

### GetPolicyObjectsOk

`func (o *InlineResponse20081Limits) GetPolicyObjectsOk() (*int32, bool)`

GetPolicyObjectsOk returns a tuple with the PolicyObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyObjects

`func (o *InlineResponse20081Limits) SetPolicyObjects(v int32)`

SetPolicyObjects sets PolicyObjects field to given value.

### HasPolicyObjects

`func (o *InlineResponse20081Limits) HasPolicyObjects() bool`

HasPolicyObjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


