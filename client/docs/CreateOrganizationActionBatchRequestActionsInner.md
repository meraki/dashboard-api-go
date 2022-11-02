# CreateOrganizationActionBatchRequestActionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resource** | **string** | Unique identifier for the resource to be acted on | 
**Operation** | **string** | The operation to be used | 
**Body** | Pointer to **map[string]interface{}** | The body of the action | [optional] 

## Methods

### NewCreateOrganizationActionBatchRequestActionsInner

`func NewCreateOrganizationActionBatchRequestActionsInner(resource string, operation string, ) *CreateOrganizationActionBatchRequestActionsInner`

NewCreateOrganizationActionBatchRequestActionsInner instantiates a new CreateOrganizationActionBatchRequestActionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationActionBatchRequestActionsInnerWithDefaults

`func NewCreateOrganizationActionBatchRequestActionsInnerWithDefaults() *CreateOrganizationActionBatchRequestActionsInner`

NewCreateOrganizationActionBatchRequestActionsInnerWithDefaults instantiates a new CreateOrganizationActionBatchRequestActionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResource

`func (o *CreateOrganizationActionBatchRequestActionsInner) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *CreateOrganizationActionBatchRequestActionsInner) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *CreateOrganizationActionBatchRequestActionsInner) SetResource(v string)`

SetResource sets Resource field to given value.


### GetOperation

`func (o *CreateOrganizationActionBatchRequestActionsInner) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *CreateOrganizationActionBatchRequestActionsInner) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *CreateOrganizationActionBatchRequestActionsInner) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetBody

`func (o *CreateOrganizationActionBatchRequestActionsInner) GetBody() map[string]interface{}`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CreateOrganizationActionBatchRequestActionsInner) GetBodyOk() (*map[string]interface{}, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CreateOrganizationActionBatchRequestActionsInner) SetBody(v map[string]interface{})`

SetBody sets Body field to given value.

### HasBody

`func (o *CreateOrganizationActionBatchRequestActionsInner) HasBody() bool`

HasBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


