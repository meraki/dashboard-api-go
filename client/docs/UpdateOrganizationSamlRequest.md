# UpdateOrganizationSamlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Boolean for updating SAML SSO enabled settings. | [optional] 

## Methods

### NewUpdateOrganizationSamlRequest

`func NewUpdateOrganizationSamlRequest() *UpdateOrganizationSamlRequest`

NewUpdateOrganizationSamlRequest instantiates a new UpdateOrganizationSamlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrganizationSamlRequestWithDefaults

`func NewUpdateOrganizationSamlRequestWithDefaults() *UpdateOrganizationSamlRequest`

NewUpdateOrganizationSamlRequestWithDefaults instantiates a new UpdateOrganizationSamlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *UpdateOrganizationSamlRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateOrganizationSamlRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateOrganizationSamlRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateOrganizationSamlRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


