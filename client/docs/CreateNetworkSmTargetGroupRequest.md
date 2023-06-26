# CreateNetworkSmTargetGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of this target group | [optional] 
**Scope** | Pointer to **string** | The scope and tag options of the target group. Comma separated values beginning with one of withAny, withAll, withoutAny, withoutAll, all, none, followed by tags. Default to none if empty. | [optional] 

## Methods

### NewCreateNetworkSmTargetGroupRequest

`func NewCreateNetworkSmTargetGroupRequest() *CreateNetworkSmTargetGroupRequest`

NewCreateNetworkSmTargetGroupRequest instantiates a new CreateNetworkSmTargetGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkSmTargetGroupRequestWithDefaults

`func NewCreateNetworkSmTargetGroupRequestWithDefaults() *CreateNetworkSmTargetGroupRequest`

NewCreateNetworkSmTargetGroupRequestWithDefaults instantiates a new CreateNetworkSmTargetGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNetworkSmTargetGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkSmTargetGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkSmTargetGroupRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateNetworkSmTargetGroupRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScope

`func (o *CreateNetworkSmTargetGroupRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CreateNetworkSmTargetGroupRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CreateNetworkSmTargetGroupRequest) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *CreateNetworkSmTargetGroupRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


