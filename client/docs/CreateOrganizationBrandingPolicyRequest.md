# CreateOrganizationBrandingPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the Dashboard branding policy. | 
**Enabled** | **bool** | Boolean indicating whether this policy is enabled. | 
**AdminSettings** | [**CreateOrganizationBrandingPolicyRequestAdminSettings**](CreateOrganizationBrandingPolicyRequestAdminSettings.md) |  | 
**HelpSettings** | Pointer to [**CreateOrganizationBrandingPolicyRequestHelpSettings**](CreateOrganizationBrandingPolicyRequestHelpSettings.md) |  | [optional] 
**CustomLogo** | Pointer to [**CreateOrganizationBrandingPolicyRequestCustomLogo**](CreateOrganizationBrandingPolicyRequestCustomLogo.md) |  | [optional] 

## Methods

### NewCreateOrganizationBrandingPolicyRequest

`func NewCreateOrganizationBrandingPolicyRequest(name string, enabled bool, adminSettings CreateOrganizationBrandingPolicyRequestAdminSettings, ) *CreateOrganizationBrandingPolicyRequest`

NewCreateOrganizationBrandingPolicyRequest instantiates a new CreateOrganizationBrandingPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationBrandingPolicyRequestWithDefaults

`func NewCreateOrganizationBrandingPolicyRequestWithDefaults() *CreateOrganizationBrandingPolicyRequest`

NewCreateOrganizationBrandingPolicyRequestWithDefaults instantiates a new CreateOrganizationBrandingPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateOrganizationBrandingPolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrganizationBrandingPolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrganizationBrandingPolicyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *CreateOrganizationBrandingPolicyRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateOrganizationBrandingPolicyRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateOrganizationBrandingPolicyRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAdminSettings

`func (o *CreateOrganizationBrandingPolicyRequest) GetAdminSettings() CreateOrganizationBrandingPolicyRequestAdminSettings`

GetAdminSettings returns the AdminSettings field if non-nil, zero value otherwise.

### GetAdminSettingsOk

`func (o *CreateOrganizationBrandingPolicyRequest) GetAdminSettingsOk() (*CreateOrganizationBrandingPolicyRequestAdminSettings, bool)`

GetAdminSettingsOk returns a tuple with the AdminSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminSettings

`func (o *CreateOrganizationBrandingPolicyRequest) SetAdminSettings(v CreateOrganizationBrandingPolicyRequestAdminSettings)`

SetAdminSettings sets AdminSettings field to given value.


### GetHelpSettings

`func (o *CreateOrganizationBrandingPolicyRequest) GetHelpSettings() CreateOrganizationBrandingPolicyRequestHelpSettings`

GetHelpSettings returns the HelpSettings field if non-nil, zero value otherwise.

### GetHelpSettingsOk

`func (o *CreateOrganizationBrandingPolicyRequest) GetHelpSettingsOk() (*CreateOrganizationBrandingPolicyRequestHelpSettings, bool)`

GetHelpSettingsOk returns a tuple with the HelpSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpSettings

`func (o *CreateOrganizationBrandingPolicyRequest) SetHelpSettings(v CreateOrganizationBrandingPolicyRequestHelpSettings)`

SetHelpSettings sets HelpSettings field to given value.

### HasHelpSettings

`func (o *CreateOrganizationBrandingPolicyRequest) HasHelpSettings() bool`

HasHelpSettings returns a boolean if a field has been set.

### GetCustomLogo

`func (o *CreateOrganizationBrandingPolicyRequest) GetCustomLogo() CreateOrganizationBrandingPolicyRequestCustomLogo`

GetCustomLogo returns the CustomLogo field if non-nil, zero value otherwise.

### GetCustomLogoOk

`func (o *CreateOrganizationBrandingPolicyRequest) GetCustomLogoOk() (*CreateOrganizationBrandingPolicyRequestCustomLogo, bool)`

GetCustomLogoOk returns a tuple with the CustomLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLogo

`func (o *CreateOrganizationBrandingPolicyRequest) SetCustomLogo(v CreateOrganizationBrandingPolicyRequestCustomLogo)`

SetCustomLogo sets CustomLogo field to given value.

### HasCustomLogo

`func (o *CreateOrganizationBrandingPolicyRequest) HasCustomLogo() bool`

HasCustomLogo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


