# GetOrganizationSmSentryPoliciesAssignmentsByNetwork200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInner**](UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInner.md) | Sentry Group Policies for the Organization keyed by the Network or Locale Id the Policy belongs to | [optional] 
**Meta** | Pointer to [**GetOrganizationSmAdminsRoles200ResponseMeta**](GetOrganizationSmAdminsRoles200ResponseMeta.md) |  | [optional] 

## Methods

### NewGetOrganizationSmSentryPoliciesAssignmentsByNetwork200ResponseInner

`func NewGetOrganizationSmSentryPoliciesAssignmentsByNetwork200ResponseInner() *GetOrganizationSmSentryPoliciesAssignmentsByNetwork200ResponseInner`

NewGetOrganizationSmSentryPoliciesAssignmentsByNetwork200ResponseInner instantiates a new GetOrganizationSmSentryPoliciesAssignmentsByNetwork200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationSmSentryPoliciesAssignmentsByNetwork200ResponseInnerWithDefaults

`func NewGetOrganizationSmSentryPoliciesAssignmentsByNetwork200ResponseInnerWithDefaults() *GetOrganizationSmSentryPoliciesAssignmentsByNetwork200ResponseInner`

NewGetOrganizationSmSentryPoliciesAssignmentsByNetwork200ResponseInnerWithDefaults instantiates a new GetOrganizationSmSentryPoliciesAssignmentsByNetwork200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *GetOrganizationSmSentryPoliciesAssignmentsByNetwork200ResponseInner) GetItems() []UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GetOrganizationSmSentryPoliciesAssignmentsByNetwork200ResponseInner) GetItemsOk() (*[]UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GetOrganizationSmSentryPoliciesAssignmentsByNetwork200ResponseInner) SetItems(v []UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInner)`

SetItems sets Items field to given value.

### HasItems

`func (o *GetOrganizationSmSentryPoliciesAssignmentsByNetwork200ResponseInner) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMeta

`func (o *GetOrganizationSmSentryPoliciesAssignmentsByNetwork200ResponseInner) GetMeta() GetOrganizationSmAdminsRoles200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetOrganizationSmSentryPoliciesAssignmentsByNetwork200ResponseInner) GetMetaOk() (*GetOrganizationSmAdminsRoles200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetOrganizationSmSentryPoliciesAssignmentsByNetwork200ResponseInner) SetMeta(v GetOrganizationSmAdminsRoles200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetOrganizationSmSentryPoliciesAssignmentsByNetwork200ResponseInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


