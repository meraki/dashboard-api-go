# GetOrganizationBrandingPolicies200ResponseInnerAdminSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppliesTo** | Pointer to **string** | Which kinds of admins this policy applies to. Can be one of &#39;All organization admins&#39;, &#39;All enterprise admins&#39;, &#39;All network admins&#39;, &#39;All admins of networks...&#39;, &#39;All admins of networks tagged...&#39;, &#39;Specific admins...&#39;, &#39;All admins&#39; or &#39;All SAML admins&#39;. | [optional] 
**Values** | Pointer to **[]string** |       If &#39;appliesTo&#39; is set to one of &#39;Specific admins...&#39;, &#39;All admins of networks...&#39; or &#39;All admins of networks tagged...&#39;, then you must specify this &#39;values&#39; property to provide the set of       entities to apply the branding policy to. For &#39;Specific admins...&#39;, specify an array of admin IDs. For &#39;All admins of       networks...&#39;, specify an array of network IDs and/or configuration template IDs. For &#39;All admins of networks tagged...&#39;,       specify an array of tag names.  | [optional] 

## Methods

### NewGetOrganizationBrandingPolicies200ResponseInnerAdminSettings

`func NewGetOrganizationBrandingPolicies200ResponseInnerAdminSettings() *GetOrganizationBrandingPolicies200ResponseInnerAdminSettings`

NewGetOrganizationBrandingPolicies200ResponseInnerAdminSettings instantiates a new GetOrganizationBrandingPolicies200ResponseInnerAdminSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationBrandingPolicies200ResponseInnerAdminSettingsWithDefaults

`func NewGetOrganizationBrandingPolicies200ResponseInnerAdminSettingsWithDefaults() *GetOrganizationBrandingPolicies200ResponseInnerAdminSettings`

NewGetOrganizationBrandingPolicies200ResponseInnerAdminSettingsWithDefaults instantiates a new GetOrganizationBrandingPolicies200ResponseInnerAdminSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppliesTo

`func (o *GetOrganizationBrandingPolicies200ResponseInnerAdminSettings) GetAppliesTo() string`

GetAppliesTo returns the AppliesTo field if non-nil, zero value otherwise.

### GetAppliesToOk

`func (o *GetOrganizationBrandingPolicies200ResponseInnerAdminSettings) GetAppliesToOk() (*string, bool)`

GetAppliesToOk returns a tuple with the AppliesTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliesTo

`func (o *GetOrganizationBrandingPolicies200ResponseInnerAdminSettings) SetAppliesTo(v string)`

SetAppliesTo sets AppliesTo field to given value.

### HasAppliesTo

`func (o *GetOrganizationBrandingPolicies200ResponseInnerAdminSettings) HasAppliesTo() bool`

HasAppliesTo returns a boolean if a field has been set.

### GetValues

`func (o *GetOrganizationBrandingPolicies200ResponseInnerAdminSettings) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetOrganizationBrandingPolicies200ResponseInnerAdminSettings) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetOrganizationBrandingPolicies200ResponseInnerAdminSettings) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *GetOrganizationBrandingPolicies200ResponseInnerAdminSettings) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


