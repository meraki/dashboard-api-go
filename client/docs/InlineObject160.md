# InlineObject160

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the Identity PSK | 
**Passphrase** | Pointer to **string** | The passphrase for client authentication. If left blank, one will be auto-generated. | [optional] 
**GroupPolicyId** | **string** | The group policy to be applied to clients | 

## Methods

### NewInlineObject160

`func NewInlineObject160(name string, groupPolicyId string, ) *InlineObject160`

NewInlineObject160 instantiates a new InlineObject160 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject160WithDefaults

`func NewInlineObject160WithDefaults() *InlineObject160`

NewInlineObject160WithDefaults instantiates a new InlineObject160 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject160) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject160) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject160) SetName(v string)`

SetName sets Name field to given value.


### GetPassphrase

`func (o *InlineObject160) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *InlineObject160) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *InlineObject160) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *InlineObject160) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### GetGroupPolicyId

`func (o *InlineObject160) GetGroupPolicyId() string`

GetGroupPolicyId returns the GroupPolicyId field if non-nil, zero value otherwise.

### GetGroupPolicyIdOk

`func (o *InlineObject160) GetGroupPolicyIdOk() (*string, bool)`

GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPolicyId

`func (o *InlineObject160) SetGroupPolicyId(v string)`

SetGroupPolicyId sets GroupPolicyId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


