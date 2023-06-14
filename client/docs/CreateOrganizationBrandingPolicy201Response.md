# CreateOrganizationBrandingPolicy201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the Dashboard branding policy. | [optional] 
**Enabled** | Pointer to **bool** | Boolean indicating whether this policy is enabled. | [optional] 
**AdminSettings** | Pointer to [**GetOrganizationBrandingPolicies200ResponseInnerAdminSettings**](GetOrganizationBrandingPolicies200ResponseInnerAdminSettings.md) |  | [optional] 
**HelpSettings** | Pointer to [**CreateOrganizationBrandingPolicyRequestHelpSettings**](CreateOrganizationBrandingPolicyRequestHelpSettings.md) |  | [optional] 
**CustomLogo** | Pointer to [**GetOrganizationBrandingPolicies200ResponseInnerCustomLogo**](GetOrganizationBrandingPolicies200ResponseInnerCustomLogo.md) |  | [optional] 

## Methods

### NewCreateOrganizationBrandingPolicy201Response

`func NewCreateOrganizationBrandingPolicy201Response() *CreateOrganizationBrandingPolicy201Response`

NewCreateOrganizationBrandingPolicy201Response instantiates a new CreateOrganizationBrandingPolicy201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationBrandingPolicy201ResponseWithDefaults

`func NewCreateOrganizationBrandingPolicy201ResponseWithDefaults() *CreateOrganizationBrandingPolicy201Response`

NewCreateOrganizationBrandingPolicy201ResponseWithDefaults instantiates a new CreateOrganizationBrandingPolicy201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateOrganizationBrandingPolicy201Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrganizationBrandingPolicy201Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrganizationBrandingPolicy201Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateOrganizationBrandingPolicy201Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *CreateOrganizationBrandingPolicy201Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateOrganizationBrandingPolicy201Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateOrganizationBrandingPolicy201Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateOrganizationBrandingPolicy201Response) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAdminSettings

`func (o *CreateOrganizationBrandingPolicy201Response) GetAdminSettings() GetOrganizationBrandingPolicies200ResponseInnerAdminSettings`

GetAdminSettings returns the AdminSettings field if non-nil, zero value otherwise.

### GetAdminSettingsOk

`func (o *CreateOrganizationBrandingPolicy201Response) GetAdminSettingsOk() (*GetOrganizationBrandingPolicies200ResponseInnerAdminSettings, bool)`

GetAdminSettingsOk returns a tuple with the AdminSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminSettings

`func (o *CreateOrganizationBrandingPolicy201Response) SetAdminSettings(v GetOrganizationBrandingPolicies200ResponseInnerAdminSettings)`

SetAdminSettings sets AdminSettings field to given value.

### HasAdminSettings

`func (o *CreateOrganizationBrandingPolicy201Response) HasAdminSettings() bool`

HasAdminSettings returns a boolean if a field has been set.

### GetHelpSettings

`func (o *CreateOrganizationBrandingPolicy201Response) GetHelpSettings() CreateOrganizationBrandingPolicyRequestHelpSettings`

GetHelpSettings returns the HelpSettings field if non-nil, zero value otherwise.

### GetHelpSettingsOk

`func (o *CreateOrganizationBrandingPolicy201Response) GetHelpSettingsOk() (*CreateOrganizationBrandingPolicyRequestHelpSettings, bool)`

GetHelpSettingsOk returns a tuple with the HelpSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpSettings

`func (o *CreateOrganizationBrandingPolicy201Response) SetHelpSettings(v CreateOrganizationBrandingPolicyRequestHelpSettings)`

SetHelpSettings sets HelpSettings field to given value.

### HasHelpSettings

`func (o *CreateOrganizationBrandingPolicy201Response) HasHelpSettings() bool`

HasHelpSettings returns a boolean if a field has been set.

### GetCustomLogo

`func (o *CreateOrganizationBrandingPolicy201Response) GetCustomLogo() GetOrganizationBrandingPolicies200ResponseInnerCustomLogo`

GetCustomLogo returns the CustomLogo field if non-nil, zero value otherwise.

### GetCustomLogoOk

`func (o *CreateOrganizationBrandingPolicy201Response) GetCustomLogoOk() (*GetOrganizationBrandingPolicies200ResponseInnerCustomLogo, bool)`

GetCustomLogoOk returns a tuple with the CustomLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLogo

`func (o *CreateOrganizationBrandingPolicy201Response) SetCustomLogo(v GetOrganizationBrandingPolicies200ResponseInnerCustomLogo)`

SetCustomLogo sets CustomLogo field to given value.

### HasCustomLogo

`func (o *CreateOrganizationBrandingPolicy201Response) HasCustomLogo() bool`

HasCustomLogo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


