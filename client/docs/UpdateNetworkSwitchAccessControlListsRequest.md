# UpdateNetworkSwitchAccessControlListsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | [**[]UpdateNetworkSwitchAccessControlListsRequestRulesInner**](UpdateNetworkSwitchAccessControlListsRequestRulesInner.md) | An ordered array of the access control list rules (not including the default rule). An empty array will clear the rules. | 

## Methods

### NewUpdateNetworkSwitchAccessControlListsRequest

`func NewUpdateNetworkSwitchAccessControlListsRequest(rules []UpdateNetworkSwitchAccessControlListsRequestRulesInner, ) *UpdateNetworkSwitchAccessControlListsRequest`

NewUpdateNetworkSwitchAccessControlListsRequest instantiates a new UpdateNetworkSwitchAccessControlListsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkSwitchAccessControlListsRequestWithDefaults

`func NewUpdateNetworkSwitchAccessControlListsRequestWithDefaults() *UpdateNetworkSwitchAccessControlListsRequest`

NewUpdateNetworkSwitchAccessControlListsRequestWithDefaults instantiates a new UpdateNetworkSwitchAccessControlListsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *UpdateNetworkSwitchAccessControlListsRequest) GetRules() []UpdateNetworkSwitchAccessControlListsRequestRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *UpdateNetworkSwitchAccessControlListsRequest) GetRulesOk() (*[]UpdateNetworkSwitchAccessControlListsRequestRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *UpdateNetworkSwitchAccessControlListsRequest) SetRules(v []UpdateNetworkSwitchAccessControlListsRequestRulesInner)`

SetRules sets Rules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


