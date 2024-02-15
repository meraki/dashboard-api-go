# UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | **string** | The Id of the Network | 
**Policies** | Pointer to [**[]UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInnerPoliciesInner**](UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInnerPoliciesInner.md) | Array of Sentry Group Policies for the Network | [optional] 

## Methods

### NewUpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner

`func NewUpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner(networkId string, ) *UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner`

NewUpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner instantiates a new UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInnerWithDefaults

`func NewUpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInnerWithDefaults() *UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner`

NewUpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInnerWithDefaults instantiates a new UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.


### GetPolicies

`func (o *UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner) GetPolicies() []UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInnerPoliciesInner`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner) GetPoliciesOk() (*[]UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInnerPoliciesInner, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner) SetPolicies(v []UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInnerPoliciesInner)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


