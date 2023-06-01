# InlineResponse20041ClientPrivacy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpireDataOlderThan** | Pointer to **int32** | The number of days, weeks, or months in Epoch time to expire the data before | [optional] 
**ExpireDataBefore** | Pointer to **time.Time** | The date to expire the data before | [optional] 

## Methods

### NewInlineResponse20041ClientPrivacy

`func NewInlineResponse20041ClientPrivacy() *InlineResponse20041ClientPrivacy`

NewInlineResponse20041ClientPrivacy instantiates a new InlineResponse20041ClientPrivacy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20041ClientPrivacyWithDefaults

`func NewInlineResponse20041ClientPrivacyWithDefaults() *InlineResponse20041ClientPrivacy`

NewInlineResponse20041ClientPrivacyWithDefaults instantiates a new InlineResponse20041ClientPrivacy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpireDataOlderThan

`func (o *InlineResponse20041ClientPrivacy) GetExpireDataOlderThan() int32`

GetExpireDataOlderThan returns the ExpireDataOlderThan field if non-nil, zero value otherwise.

### GetExpireDataOlderThanOk

`func (o *InlineResponse20041ClientPrivacy) GetExpireDataOlderThanOk() (*int32, bool)`

GetExpireDataOlderThanOk returns a tuple with the ExpireDataOlderThan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDataOlderThan

`func (o *InlineResponse20041ClientPrivacy) SetExpireDataOlderThan(v int32)`

SetExpireDataOlderThan sets ExpireDataOlderThan field to given value.

### HasExpireDataOlderThan

`func (o *InlineResponse20041ClientPrivacy) HasExpireDataOlderThan() bool`

HasExpireDataOlderThan returns a boolean if a field has been set.

### GetExpireDataBefore

`func (o *InlineResponse20041ClientPrivacy) GetExpireDataBefore() time.Time`

GetExpireDataBefore returns the ExpireDataBefore field if non-nil, zero value otherwise.

### GetExpireDataBeforeOk

`func (o *InlineResponse20041ClientPrivacy) GetExpireDataBeforeOk() (*time.Time, bool)`

GetExpireDataBeforeOk returns a tuple with the ExpireDataBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDataBefore

`func (o *InlineResponse20041ClientPrivacy) SetExpireDataBefore(v time.Time)`

SetExpireDataBefore sets ExpireDataBefore field to given value.

### HasExpireDataBefore

`func (o *InlineResponse20041ClientPrivacy) HasExpireDataBefore() bool`

HasExpireDataBefore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


