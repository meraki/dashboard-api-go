# CreateOrganizationBrandingPolicyRequestCustomLogo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether or not there is a custom logo enabled. | [optional] 
**Image** | Pointer to **map[string]interface{}** | Properties for setting the image. | [optional] 
**Contents** | Pointer to **string** | The file contents (a base 64 encoded string) of your new logo. | [optional] 
**Format** | Pointer to **string** | The format of the encoded contents.  Supported formats are &#39;png&#39;, &#39;gif&#39;, and jpg&#39;. Note that all images are saved as PNG files, regardless of the format they are uploaded in. | [optional] 

## Methods

### NewCreateOrganizationBrandingPolicyRequestCustomLogo

`func NewCreateOrganizationBrandingPolicyRequestCustomLogo() *CreateOrganizationBrandingPolicyRequestCustomLogo`

NewCreateOrganizationBrandingPolicyRequestCustomLogo instantiates a new CreateOrganizationBrandingPolicyRequestCustomLogo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationBrandingPolicyRequestCustomLogoWithDefaults

`func NewCreateOrganizationBrandingPolicyRequestCustomLogoWithDefaults() *CreateOrganizationBrandingPolicyRequestCustomLogo`

NewCreateOrganizationBrandingPolicyRequestCustomLogoWithDefaults instantiates a new CreateOrganizationBrandingPolicyRequestCustomLogo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *CreateOrganizationBrandingPolicyRequestCustomLogo) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateOrganizationBrandingPolicyRequestCustomLogo) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateOrganizationBrandingPolicyRequestCustomLogo) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateOrganizationBrandingPolicyRequestCustomLogo) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetImage

`func (o *CreateOrganizationBrandingPolicyRequestCustomLogo) GetImage() map[string]interface{}`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CreateOrganizationBrandingPolicyRequestCustomLogo) GetImageOk() (*map[string]interface{}, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CreateOrganizationBrandingPolicyRequestCustomLogo) SetImage(v map[string]interface{})`

SetImage sets Image field to given value.

### HasImage

`func (o *CreateOrganizationBrandingPolicyRequestCustomLogo) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetContents

`func (o *CreateOrganizationBrandingPolicyRequestCustomLogo) GetContents() string`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *CreateOrganizationBrandingPolicyRequestCustomLogo) GetContentsOk() (*string, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *CreateOrganizationBrandingPolicyRequestCustomLogo) SetContents(v string)`

SetContents sets Contents field to given value.

### HasContents

`func (o *CreateOrganizationBrandingPolicyRequestCustomLogo) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetFormat

`func (o *CreateOrganizationBrandingPolicyRequestCustomLogo) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CreateOrganizationBrandingPolicyRequestCustomLogo) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CreateOrganizationBrandingPolicyRequestCustomLogo) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CreateOrganizationBrandingPolicyRequestCustomLogo) HasFormat() bool`

HasFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


