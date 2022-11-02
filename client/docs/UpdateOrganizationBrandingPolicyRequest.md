# UpdateOrganizationBrandingPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the Dashboard branding policy. | [optional] 
**Enabled** | Pointer to **bool** | Boolean indicating whether this policy is enabled. | [optional] 
**AdminSettings** | Pointer to [**CreateOrganizationBrandingPolicyRequestAdminSettings**](CreateOrganizationBrandingPolicyRequestAdminSettings.md) |  | [optional] 
**HelpSettings** | Pointer to [**UpdateOrganizationBrandingPolicyRequestHelpSettings**](UpdateOrganizationBrandingPolicyRequestHelpSettings.md) |  | [optional] 
**CustomLogo** | Pointer to [**CreateOrganizationBrandingPolicyRequestCustomLogo**](CreateOrganizationBrandingPolicyRequestCustomLogo.md) |  | [optional] 

## Methods

### NewUpdateOrganizationBrandingPolicyRequest

`func NewUpdateOrganizationBrandingPolicyRequest() *UpdateOrganizationBrandingPolicyRequest`

NewUpdateOrganizationBrandingPolicyRequest instantiates a new UpdateOrganizationBrandingPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrganizationBrandingPolicyRequestWithDefaults

`func NewUpdateOrganizationBrandingPolicyRequestWithDefaults() *UpdateOrganizationBrandingPolicyRequest`

NewUpdateOrganizationBrandingPolicyRequestWithDefaults instantiates a new UpdateOrganizationBrandingPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateOrganizationBrandingPolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateOrganizationBrandingPolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateOrganizationBrandingPolicyRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateOrganizationBrandingPolicyRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateOrganizationBrandingPolicyRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateOrganizationBrandingPolicyRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateOrganizationBrandingPolicyRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateOrganizationBrandingPolicyRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAdminSettings

`func (o *UpdateOrganizationBrandingPolicyRequest) GetAdminSettings() CreateOrganizationBrandingPolicyRequestAdminSettings`

GetAdminSettings returns the AdminSettings field if non-nil, zero value otherwise.

### GetAdminSettingsOk

`func (o *UpdateOrganizationBrandingPolicyRequest) GetAdminSettingsOk() (*CreateOrganizationBrandingPolicyRequestAdminSettings, bool)`

GetAdminSettingsOk returns a tuple with the AdminSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminSettings

`func (o *UpdateOrganizationBrandingPolicyRequest) SetAdminSettings(v CreateOrganizationBrandingPolicyRequestAdminSettings)`

SetAdminSettings sets AdminSettings field to given value.

### HasAdminSettings

`func (o *UpdateOrganizationBrandingPolicyRequest) HasAdminSettings() bool`

HasAdminSettings returns a boolean if a field has been set.

### GetHelpSettings

`func (o *UpdateOrganizationBrandingPolicyRequest) GetHelpSettings() UpdateOrganizationBrandingPolicyRequestHelpSettings`

GetHelpSettings returns the HelpSettings field if non-nil, zero value otherwise.

### GetHelpSettingsOk

`func (o *UpdateOrganizationBrandingPolicyRequest) GetHelpSettingsOk() (*UpdateOrganizationBrandingPolicyRequestHelpSettings, bool)`

GetHelpSettingsOk returns a tuple with the HelpSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpSettings

`func (o *UpdateOrganizationBrandingPolicyRequest) SetHelpSettings(v UpdateOrganizationBrandingPolicyRequestHelpSettings)`

SetHelpSettings sets HelpSettings field to given value.

### HasHelpSettings

`func (o *UpdateOrganizationBrandingPolicyRequest) HasHelpSettings() bool`

HasHelpSettings returns a boolean if a field has been set.

### GetCustomLogo

`func (o *UpdateOrganizationBrandingPolicyRequest) GetCustomLogo() CreateOrganizationBrandingPolicyRequestCustomLogo`

GetCustomLogo returns the CustomLogo field if non-nil, zero value otherwise.

### GetCustomLogoOk

`func (o *UpdateOrganizationBrandingPolicyRequest) GetCustomLogoOk() (*CreateOrganizationBrandingPolicyRequestCustomLogo, bool)`

GetCustomLogoOk returns a tuple with the CustomLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLogo

`func (o *UpdateOrganizationBrandingPolicyRequest) SetCustomLogo(v CreateOrganizationBrandingPolicyRequestCustomLogo)`

SetCustomLogo sets CustomLogo field to given value.

### HasCustomLogo

`func (o *UpdateOrganizationBrandingPolicyRequest) HasCustomLogo() bool`

HasCustomLogo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


