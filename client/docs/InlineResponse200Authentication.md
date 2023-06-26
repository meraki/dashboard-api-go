# InlineResponse200Authentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | Authentication mode | [optional] 
**Api** | Pointer to [**InlineResponse200AuthenticationApi**](InlineResponse200AuthenticationApi.md) |  | [optional] 
**TwoFactor** | Pointer to [**InlineResponse200AuthenticationTwoFactor**](InlineResponse200AuthenticationTwoFactor.md) |  | [optional] 
**Saml** | Pointer to [**InlineResponse200AuthenticationSaml**](InlineResponse200AuthenticationSaml.md) |  | [optional] 

## Methods

### NewInlineResponse200Authentication

`func NewInlineResponse200Authentication() *InlineResponse200Authentication`

NewInlineResponse200Authentication instantiates a new InlineResponse200Authentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200AuthenticationWithDefaults

`func NewInlineResponse200AuthenticationWithDefaults() *InlineResponse200Authentication`

NewInlineResponse200AuthenticationWithDefaults instantiates a new InlineResponse200Authentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineResponse200Authentication) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse200Authentication) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse200Authentication) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse200Authentication) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetApi

`func (o *InlineResponse200Authentication) GetApi() InlineResponse200AuthenticationApi`

GetApi returns the Api field if non-nil, zero value otherwise.

### GetApiOk

`func (o *InlineResponse200Authentication) GetApiOk() (*InlineResponse200AuthenticationApi, bool)`

GetApiOk returns a tuple with the Api field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApi

`func (o *InlineResponse200Authentication) SetApi(v InlineResponse200AuthenticationApi)`

SetApi sets Api field to given value.

### HasApi

`func (o *InlineResponse200Authentication) HasApi() bool`

HasApi returns a boolean if a field has been set.

### GetTwoFactor

`func (o *InlineResponse200Authentication) GetTwoFactor() InlineResponse200AuthenticationTwoFactor`

GetTwoFactor returns the TwoFactor field if non-nil, zero value otherwise.

### GetTwoFactorOk

`func (o *InlineResponse200Authentication) GetTwoFactorOk() (*InlineResponse200AuthenticationTwoFactor, bool)`

GetTwoFactorOk returns a tuple with the TwoFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactor

`func (o *InlineResponse200Authentication) SetTwoFactor(v InlineResponse200AuthenticationTwoFactor)`

SetTwoFactor sets TwoFactor field to given value.

### HasTwoFactor

`func (o *InlineResponse200Authentication) HasTwoFactor() bool`

HasTwoFactor returns a boolean if a field has been set.

### GetSaml

`func (o *InlineResponse200Authentication) GetSaml() InlineResponse200AuthenticationSaml`

GetSaml returns the Saml field if non-nil, zero value otherwise.

### GetSamlOk

`func (o *InlineResponse200Authentication) GetSamlOk() (*InlineResponse200AuthenticationSaml, bool)`

GetSamlOk returns a tuple with the Saml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml

`func (o *InlineResponse200Authentication) SetSaml(v InlineResponse200AuthenticationSaml)`

SetSaml sets Saml field to given value.

### HasSaml

`func (o *InlineResponse200Authentication) HasSaml() bool`

HasSaml returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


