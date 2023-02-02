# InlineResponse201

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** | The newly generated authentication token for the vMX instance | [optional] 
**ExpiresAt** | Pointer to **string** | The expiration time for the token, in ISO 8601 format | [optional] 

## Methods

### NewInlineResponse201

`func NewInlineResponse201() *InlineResponse201`

NewInlineResponse201 instantiates a new InlineResponse201 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse201WithDefaults

`func NewInlineResponse201WithDefaults() *InlineResponse201`

NewInlineResponse201WithDefaults instantiates a new InlineResponse201 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *InlineResponse201) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *InlineResponse201) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *InlineResponse201) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *InlineResponse201) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetExpiresAt

`func (o *InlineResponse201) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *InlineResponse201) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *InlineResponse201) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *InlineResponse201) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


