# InlineObject170

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Confirmed** | Pointer to **bool** | A boolean representing whether or not the batch has been confirmed. This property cannot be unset once it is true. | [optional] 
**Synchronous** | Pointer to **bool** | Set to true to force the batch to run synchronous. There can be at most 20 actions in synchronous batch. | [optional] 

## Methods

### NewInlineObject170

`func NewInlineObject170() *InlineObject170`

NewInlineObject170 instantiates a new InlineObject170 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject170WithDefaults

`func NewInlineObject170WithDefaults() *InlineObject170`

NewInlineObject170WithDefaults instantiates a new InlineObject170 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfirmed

`func (o *InlineObject170) GetConfirmed() bool`

GetConfirmed returns the Confirmed field if non-nil, zero value otherwise.

### GetConfirmedOk

`func (o *InlineObject170) GetConfirmedOk() (*bool, bool)`

GetConfirmedOk returns a tuple with the Confirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmed

`func (o *InlineObject170) SetConfirmed(v bool)`

SetConfirmed sets Confirmed field to given value.

### HasConfirmed

`func (o *InlineObject170) HasConfirmed() bool`

HasConfirmed returns a boolean if a field has been set.

### GetSynchronous

`func (o *InlineObject170) GetSynchronous() bool`

GetSynchronous returns the Synchronous field if non-nil, zero value otherwise.

### GetSynchronousOk

`func (o *InlineObject170) GetSynchronousOk() (*bool, bool)`

GetSynchronousOk returns a tuple with the Synchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchronous

`func (o *InlineObject170) SetSynchronous(v bool)`

SetSynchronous sets Synchronous field to given value.

### HasSynchronous

`func (o *InlineObject170) HasSynchronous() bool`

HasSynchronous returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


