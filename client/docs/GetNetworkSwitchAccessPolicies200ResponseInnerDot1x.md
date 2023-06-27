# GetNetworkSwitchAccessPolicies200ResponseInnerDot1x

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ControlDirection** | Pointer to **string** | Supports either &#39;both&#39; or &#39;inbound&#39;. Set to &#39;inbound&#39; to allow unauthorized egress on the switchport. Set to &#39;both&#39; to control both traffic directions with authorization. Defaults to &#39;both&#39; | [optional] 

## Methods

### NewGetNetworkSwitchAccessPolicies200ResponseInnerDot1x

`func NewGetNetworkSwitchAccessPolicies200ResponseInnerDot1x() *GetNetworkSwitchAccessPolicies200ResponseInnerDot1x`

NewGetNetworkSwitchAccessPolicies200ResponseInnerDot1x instantiates a new GetNetworkSwitchAccessPolicies200ResponseInnerDot1x object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSwitchAccessPolicies200ResponseInnerDot1xWithDefaults

`func NewGetNetworkSwitchAccessPolicies200ResponseInnerDot1xWithDefaults() *GetNetworkSwitchAccessPolicies200ResponseInnerDot1x`

NewGetNetworkSwitchAccessPolicies200ResponseInnerDot1xWithDefaults instantiates a new GetNetworkSwitchAccessPolicies200ResponseInnerDot1x object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControlDirection

`func (o *GetNetworkSwitchAccessPolicies200ResponseInnerDot1x) GetControlDirection() string`

GetControlDirection returns the ControlDirection field if non-nil, zero value otherwise.

### GetControlDirectionOk

`func (o *GetNetworkSwitchAccessPolicies200ResponseInnerDot1x) GetControlDirectionOk() (*string, bool)`

GetControlDirectionOk returns a tuple with the ControlDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlDirection

`func (o *GetNetworkSwitchAccessPolicies200ResponseInnerDot1x) SetControlDirection(v string)`

SetControlDirection sets ControlDirection field to given value.

### HasControlDirection

`func (o *GetNetworkSwitchAccessPolicies200ResponseInnerDot1x) HasControlDirection() bool`

HasControlDirection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


