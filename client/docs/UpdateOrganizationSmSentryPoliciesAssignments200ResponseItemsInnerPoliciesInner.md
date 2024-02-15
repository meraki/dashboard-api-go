# UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyId** | Pointer to **string** | The Id of the Sentry Policy | [optional] 
**NetworkId** | Pointer to **string** | The Id of the Network the Sentry Policy is associated with. In a locale, this should be the Wireless Group if present, otherwise the Wired Group. | [optional] 
**SmNetworkId** | Pointer to **string** | The Id of the Systems Manager Network the Sentry Policy is assigned to | [optional] 
**Tags** | Pointer to **[]string** | The tags of the Sentry Policy | [optional] 
**Scope** | Pointer to **string** | The scope of the Sentry Policy | [optional] 
**GroupNumber** | Pointer to **string** | The number of the Group Policy | [optional] 
**GroupPolicyId** | Pointer to **string** | The Id of the Group Policy. This is associated with the network specified by the networkId. | [optional] 
**Priority** | Pointer to **string** | The priority of the Sentry Policy | [optional] 
**CreatedAt** | Pointer to **string** | The creation time of the Sentry Policy | [optional] 
**LastUpdatedAt** | Pointer to **string** | The last update time of the Sentry Policy | [optional] 

## Methods

### NewUpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner

`func NewUpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner() *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner`

NewUpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner instantiates a new UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInnerWithDefaults

`func NewUpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInnerWithDefaults() *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner`

NewUpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInnerWithDefaults instantiates a new UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyId

`func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetNetworkId

`func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetSmNetworkId

`func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) GetSmNetworkId() string`

GetSmNetworkId returns the SmNetworkId field if non-nil, zero value otherwise.

### GetSmNetworkIdOk

`func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) GetSmNetworkIdOk() (*string, bool)`

GetSmNetworkIdOk returns a tuple with the SmNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmNetworkId

`func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) SetSmNetworkId(v string)`

SetSmNetworkId sets SmNetworkId field to given value.

### HasSmNetworkId

`func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) HasSmNetworkId() bool`

HasSmNetworkId returns a boolean if a field has been set.

### GetTags

`func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetScope

`func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetGroupNumber

`func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) GetGroupNumber() string`

GetGroupNumber returns the GroupNumber field if non-nil, zero value otherwise.

### GetGroupNumberOk

`func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) GetGroupNumberOk() (*string, bool)`

GetGroupNumberOk returns a tuple with the GroupNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupNumber

`func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) SetGroupNumber(v string)`

SetGroupNumber sets GroupNumber field to given value.

### HasGroupNumber

`func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) HasGroupNumber() bool`

HasGroupNumber returns a boolean if a field has been set.

### GetGroupPolicyId

`func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) GetGroupPolicyId() string`

GetGroupPolicyId returns the GroupPolicyId field if non-nil, zero value otherwise.

### GetGroupPolicyIdOk

`func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) GetGroupPolicyIdOk() (*string, bool)`

GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPolicyId

`func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) SetGroupPolicyId(v string)`

SetGroupPolicyId sets GroupPolicyId field to given value.

### HasGroupPolicyId

`func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) HasGroupPolicyId() bool`

HasGroupPolicyId returns a boolean if a field has been set.

### GetPriority

`func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) GetLastUpdatedAt() string`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) GetLastUpdatedAtOk() (*string, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) SetLastUpdatedAt(v string)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


