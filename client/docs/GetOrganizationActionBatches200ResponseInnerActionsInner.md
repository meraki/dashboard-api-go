# GetOrganizationActionBatches200ResponseInnerActionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resource** | **string** | Unique identifier for the resource to be acted on | 
**Operation** | **string** | The operation to be used by this action | 
**Body** | Pointer to **map[string]interface{}** | Data provided in the body of the Action. Contents depend on the Action type | [optional] 

## Methods

### NewGetOrganizationActionBatches200ResponseInnerActionsInner

`func NewGetOrganizationActionBatches200ResponseInnerActionsInner(resource string, operation string, ) *GetOrganizationActionBatches200ResponseInnerActionsInner`

NewGetOrganizationActionBatches200ResponseInnerActionsInner instantiates a new GetOrganizationActionBatches200ResponseInnerActionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationActionBatches200ResponseInnerActionsInnerWithDefaults

`func NewGetOrganizationActionBatches200ResponseInnerActionsInnerWithDefaults() *GetOrganizationActionBatches200ResponseInnerActionsInner`

NewGetOrganizationActionBatches200ResponseInnerActionsInnerWithDefaults instantiates a new GetOrganizationActionBatches200ResponseInnerActionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResource

`func (o *GetOrganizationActionBatches200ResponseInnerActionsInner) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *GetOrganizationActionBatches200ResponseInnerActionsInner) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *GetOrganizationActionBatches200ResponseInnerActionsInner) SetResource(v string)`

SetResource sets Resource field to given value.


### GetOperation

`func (o *GetOrganizationActionBatches200ResponseInnerActionsInner) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *GetOrganizationActionBatches200ResponseInnerActionsInner) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *GetOrganizationActionBatches200ResponseInnerActionsInner) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetBody

`func (o *GetOrganizationActionBatches200ResponseInnerActionsInner) GetBody() map[string]interface{}`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *GetOrganizationActionBatches200ResponseInnerActionsInner) GetBodyOk() (*map[string]interface{}, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *GetOrganizationActionBatches200ResponseInnerActionsInner) SetBody(v map[string]interface{})`

SetBody sets Body field to given value.

### HasBody

`func (o *GetOrganizationActionBatches200ResponseInnerActionsInner) HasBody() bool`

HasBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


