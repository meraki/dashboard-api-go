# OrganizationsOrganizationIdBrandingPoliciesHelpSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HelpTab** | Pointer to **string** |       The Help tab, under which all support information resides. If this tab is hidden, no other &#39;Help&#39; branding       customizations will be visible. Can be one of &#39;default or inherit&#39;, &#39;hide&#39; or &#39;show&#39;.  | [optional] 
**GetHelpSubtab** | Pointer to **string** |       The &#39;Help -&gt; Get Help&#39; subtab on which Cisco Meraki KB, Product Manuals, and Support/Case Information are displayed. Note       that if this subtab is hidden, branding customizations for the KB on &#39;Get help&#39;, Cisco Meraki product documentation,       and support contact info will not be visible. Can be one of &#39;default or inherit&#39;, &#39;hide&#39; or &#39;show&#39;.  | [optional] 
**CommunitySubtab** | Pointer to **string** |       The &#39;Help -&gt; Community&#39; subtab which provides a link to Meraki Community. Can be one of &#39;default or inherit&#39;, &#39;hide&#39; or &#39;show&#39;.  | [optional] 
**CasesSubtab** | Pointer to **string** |       The &#39;Help -&gt; Cases&#39; Dashboard subtab on which Cisco Meraki support cases for this organization can be managed. Can be one       of &#39;default or inherit&#39;, &#39;hide&#39; or &#39;show&#39;.  | [optional] 
**DataProtectionRequestsSubtab** | Pointer to **string** |       The &#39;Help -&gt; Data protection requests&#39; Dashboard subtab on which requests to delete, restrict, or export end-user data can       be audited. Can be one of &#39;default or inherit&#39;, &#39;hide&#39; or &#39;show&#39;.  | [optional] 
**GetHelpSubtabKnowledgeBaseSearch** | Pointer to **string** |       The KB search box which appears on the Help page. Can be one of &#39;default or inherit&#39;, &#39;hide&#39;, &#39;show&#39;, or a replacement custom HTML string.  | [optional] 
**UniversalSearchKnowledgeBaseSearch** | Pointer to **string** |       The universal search box always visible on Dashboard will, by default, present results from the Meraki KB. This configures       whether these Meraki KB results should be returned. Can be one of &#39;default or inherit&#39;, &#39;hide&#39; or &#39;show&#39;.  | [optional] 
**CiscoMerakiProductDocumentation** | Pointer to **string** |       The &#39;Product Manuals&#39; section of the &#39;Help -&gt; Get Help&#39; subtab. Can be one of &#39;default or inherit&#39;, &#39;hide&#39;, &#39;show&#39;, or a replacement custom HTML string.  | [optional] 
**SupportContactInfo** | Pointer to **string** |       The &#39;Contact Meraki Support&#39; section of the &#39;Help -&gt; Get Help&#39; subtab. Can be one of &#39;default or inherit&#39;, &#39;hide&#39;, &#39;show&#39;, or a replacement custom HTML string.  | [optional] 
**NewFeaturesSubtab** | Pointer to **string** |       The &#39;Help -&gt; New features&#39; subtab where new Dashboard features are detailed. Can be one of &#39;default or inherit&#39;, &#39;hide&#39; or &#39;show&#39;.  | [optional] 
**FirewallInfoSubtab** | Pointer to **string** |       The &#39;Help -&gt; Firewall info&#39; subtab where necessary upstream firewall rules for communication to the Cisco Meraki cloud are       listed. Can be one of &#39;default or inherit&#39;, &#39;hide&#39; or &#39;show&#39;.  | [optional] 
**ApiDocsSubtab** | Pointer to **string** |       The &#39;Help -&gt; API docs&#39; subtab where a detailed description of the Dashboard API is listed. Can be one of       &#39;default or inherit&#39;, &#39;hide&#39; or &#39;show&#39;.  | [optional] 
**HardwareReplacementsSubtab** | Pointer to **string** |       The &#39;Help -&gt; Replacement info&#39; subtab where important information regarding device replacements is detailed. Can be one of       &#39;default or inherit&#39;, &#39;hide&#39; or &#39;show&#39;.  | [optional] 
**SmForums** | Pointer to **string** |       The &#39;SM Forums&#39; subtab which links to community-based support for Cisco Meraki Systems Manager. Only configurable for       organizations that contain Systems Manager networks. Can be one of &#39;default or inherit&#39;, &#39;hide&#39; or &#39;show&#39;.  | [optional] 
**HelpWidget** | Pointer to **string** |       The &#39;Help Widget&#39; is a support widget which provides access to live chat, documentation links, Sales contact info,       and other contact avenues to reach Meraki Support. Can be one of &#39;default or inherit&#39;, &#39;hide&#39; or &#39;show&#39;.  | [optional] 

## Methods

### NewOrganizationsOrganizationIdBrandingPoliciesHelpSettings

`func NewOrganizationsOrganizationIdBrandingPoliciesHelpSettings() *OrganizationsOrganizationIdBrandingPoliciesHelpSettings`

NewOrganizationsOrganizationIdBrandingPoliciesHelpSettings instantiates a new OrganizationsOrganizationIdBrandingPoliciesHelpSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsOrganizationIdBrandingPoliciesHelpSettingsWithDefaults

`func NewOrganizationsOrganizationIdBrandingPoliciesHelpSettingsWithDefaults() *OrganizationsOrganizationIdBrandingPoliciesHelpSettings`

NewOrganizationsOrganizationIdBrandingPoliciesHelpSettingsWithDefaults instantiates a new OrganizationsOrganizationIdBrandingPoliciesHelpSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHelpTab

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetHelpTab() string`

GetHelpTab returns the HelpTab field if non-nil, zero value otherwise.

### GetHelpTabOk

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetHelpTabOk() (*string, bool)`

GetHelpTabOk returns a tuple with the HelpTab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpTab

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) SetHelpTab(v string)`

SetHelpTab sets HelpTab field to given value.

### HasHelpTab

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasHelpTab() bool`

HasHelpTab returns a boolean if a field has been set.

### GetGetHelpSubtab

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetGetHelpSubtab() string`

GetGetHelpSubtab returns the GetHelpSubtab field if non-nil, zero value otherwise.

### GetGetHelpSubtabOk

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetGetHelpSubtabOk() (*string, bool)`

GetGetHelpSubtabOk returns a tuple with the GetHelpSubtab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetHelpSubtab

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) SetGetHelpSubtab(v string)`

SetGetHelpSubtab sets GetHelpSubtab field to given value.

### HasGetHelpSubtab

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasGetHelpSubtab() bool`

HasGetHelpSubtab returns a boolean if a field has been set.

### GetCommunitySubtab

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetCommunitySubtab() string`

GetCommunitySubtab returns the CommunitySubtab field if non-nil, zero value otherwise.

### GetCommunitySubtabOk

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetCommunitySubtabOk() (*string, bool)`

GetCommunitySubtabOk returns a tuple with the CommunitySubtab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunitySubtab

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) SetCommunitySubtab(v string)`

SetCommunitySubtab sets CommunitySubtab field to given value.

### HasCommunitySubtab

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasCommunitySubtab() bool`

HasCommunitySubtab returns a boolean if a field has been set.

### GetCasesSubtab

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetCasesSubtab() string`

GetCasesSubtab returns the CasesSubtab field if non-nil, zero value otherwise.

### GetCasesSubtabOk

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetCasesSubtabOk() (*string, bool)`

GetCasesSubtabOk returns a tuple with the CasesSubtab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCasesSubtab

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) SetCasesSubtab(v string)`

SetCasesSubtab sets CasesSubtab field to given value.

### HasCasesSubtab

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasCasesSubtab() bool`

HasCasesSubtab returns a boolean if a field has been set.

### GetDataProtectionRequestsSubtab

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetDataProtectionRequestsSubtab() string`

GetDataProtectionRequestsSubtab returns the DataProtectionRequestsSubtab field if non-nil, zero value otherwise.

### GetDataProtectionRequestsSubtabOk

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetDataProtectionRequestsSubtabOk() (*string, bool)`

GetDataProtectionRequestsSubtabOk returns a tuple with the DataProtectionRequestsSubtab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataProtectionRequestsSubtab

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) SetDataProtectionRequestsSubtab(v string)`

SetDataProtectionRequestsSubtab sets DataProtectionRequestsSubtab field to given value.

### HasDataProtectionRequestsSubtab

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasDataProtectionRequestsSubtab() bool`

HasDataProtectionRequestsSubtab returns a boolean if a field has been set.

### GetGetHelpSubtabKnowledgeBaseSearch

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetGetHelpSubtabKnowledgeBaseSearch() string`

GetGetHelpSubtabKnowledgeBaseSearch returns the GetHelpSubtabKnowledgeBaseSearch field if non-nil, zero value otherwise.

### GetGetHelpSubtabKnowledgeBaseSearchOk

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetGetHelpSubtabKnowledgeBaseSearchOk() (*string, bool)`

GetGetHelpSubtabKnowledgeBaseSearchOk returns a tuple with the GetHelpSubtabKnowledgeBaseSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetHelpSubtabKnowledgeBaseSearch

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) SetGetHelpSubtabKnowledgeBaseSearch(v string)`

SetGetHelpSubtabKnowledgeBaseSearch sets GetHelpSubtabKnowledgeBaseSearch field to given value.

### HasGetHelpSubtabKnowledgeBaseSearch

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasGetHelpSubtabKnowledgeBaseSearch() bool`

HasGetHelpSubtabKnowledgeBaseSearch returns a boolean if a field has been set.

### GetUniversalSearchKnowledgeBaseSearch

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetUniversalSearchKnowledgeBaseSearch() string`

GetUniversalSearchKnowledgeBaseSearch returns the UniversalSearchKnowledgeBaseSearch field if non-nil, zero value otherwise.

### GetUniversalSearchKnowledgeBaseSearchOk

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetUniversalSearchKnowledgeBaseSearchOk() (*string, bool)`

GetUniversalSearchKnowledgeBaseSearchOk returns a tuple with the UniversalSearchKnowledgeBaseSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniversalSearchKnowledgeBaseSearch

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) SetUniversalSearchKnowledgeBaseSearch(v string)`

SetUniversalSearchKnowledgeBaseSearch sets UniversalSearchKnowledgeBaseSearch field to given value.

### HasUniversalSearchKnowledgeBaseSearch

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasUniversalSearchKnowledgeBaseSearch() bool`

HasUniversalSearchKnowledgeBaseSearch returns a boolean if a field has been set.

### GetCiscoMerakiProductDocumentation

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetCiscoMerakiProductDocumentation() string`

GetCiscoMerakiProductDocumentation returns the CiscoMerakiProductDocumentation field if non-nil, zero value otherwise.

### GetCiscoMerakiProductDocumentationOk

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetCiscoMerakiProductDocumentationOk() (*string, bool)`

GetCiscoMerakiProductDocumentationOk returns a tuple with the CiscoMerakiProductDocumentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiscoMerakiProductDocumentation

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) SetCiscoMerakiProductDocumentation(v string)`

SetCiscoMerakiProductDocumentation sets CiscoMerakiProductDocumentation field to given value.

### HasCiscoMerakiProductDocumentation

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasCiscoMerakiProductDocumentation() bool`

HasCiscoMerakiProductDocumentation returns a boolean if a field has been set.

### GetSupportContactInfo

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetSupportContactInfo() string`

GetSupportContactInfo returns the SupportContactInfo field if non-nil, zero value otherwise.

### GetSupportContactInfoOk

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetSupportContactInfoOk() (*string, bool)`

GetSupportContactInfoOk returns a tuple with the SupportContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportContactInfo

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) SetSupportContactInfo(v string)`

SetSupportContactInfo sets SupportContactInfo field to given value.

### HasSupportContactInfo

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasSupportContactInfo() bool`

HasSupportContactInfo returns a boolean if a field has been set.

### GetNewFeaturesSubtab

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetNewFeaturesSubtab() string`

GetNewFeaturesSubtab returns the NewFeaturesSubtab field if non-nil, zero value otherwise.

### GetNewFeaturesSubtabOk

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetNewFeaturesSubtabOk() (*string, bool)`

GetNewFeaturesSubtabOk returns a tuple with the NewFeaturesSubtab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewFeaturesSubtab

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) SetNewFeaturesSubtab(v string)`

SetNewFeaturesSubtab sets NewFeaturesSubtab field to given value.

### HasNewFeaturesSubtab

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasNewFeaturesSubtab() bool`

HasNewFeaturesSubtab returns a boolean if a field has been set.

### GetFirewallInfoSubtab

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetFirewallInfoSubtab() string`

GetFirewallInfoSubtab returns the FirewallInfoSubtab field if non-nil, zero value otherwise.

### GetFirewallInfoSubtabOk

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetFirewallInfoSubtabOk() (*string, bool)`

GetFirewallInfoSubtabOk returns a tuple with the FirewallInfoSubtab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallInfoSubtab

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) SetFirewallInfoSubtab(v string)`

SetFirewallInfoSubtab sets FirewallInfoSubtab field to given value.

### HasFirewallInfoSubtab

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasFirewallInfoSubtab() bool`

HasFirewallInfoSubtab returns a boolean if a field has been set.

### GetApiDocsSubtab

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetApiDocsSubtab() string`

GetApiDocsSubtab returns the ApiDocsSubtab field if non-nil, zero value otherwise.

### GetApiDocsSubtabOk

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetApiDocsSubtabOk() (*string, bool)`

GetApiDocsSubtabOk returns a tuple with the ApiDocsSubtab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiDocsSubtab

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) SetApiDocsSubtab(v string)`

SetApiDocsSubtab sets ApiDocsSubtab field to given value.

### HasApiDocsSubtab

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasApiDocsSubtab() bool`

HasApiDocsSubtab returns a boolean if a field has been set.

### GetHardwareReplacementsSubtab

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetHardwareReplacementsSubtab() string`

GetHardwareReplacementsSubtab returns the HardwareReplacementsSubtab field if non-nil, zero value otherwise.

### GetHardwareReplacementsSubtabOk

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetHardwareReplacementsSubtabOk() (*string, bool)`

GetHardwareReplacementsSubtabOk returns a tuple with the HardwareReplacementsSubtab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareReplacementsSubtab

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) SetHardwareReplacementsSubtab(v string)`

SetHardwareReplacementsSubtab sets HardwareReplacementsSubtab field to given value.

### HasHardwareReplacementsSubtab

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasHardwareReplacementsSubtab() bool`

HasHardwareReplacementsSubtab returns a boolean if a field has been set.

### GetSmForums

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetSmForums() string`

GetSmForums returns the SmForums field if non-nil, zero value otherwise.

### GetSmForumsOk

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetSmForumsOk() (*string, bool)`

GetSmForumsOk returns a tuple with the SmForums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmForums

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) SetSmForums(v string)`

SetSmForums sets SmForums field to given value.

### HasSmForums

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasSmForums() bool`

HasSmForums returns a boolean if a field has been set.

### GetHelpWidget

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetHelpWidget() string`

GetHelpWidget returns the HelpWidget field if non-nil, zero value otherwise.

### GetHelpWidgetOk

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) GetHelpWidgetOk() (*string, bool)`

GetHelpWidgetOk returns a tuple with the HelpWidget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpWidget

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) SetHelpWidget(v string)`

SetHelpWidget sets HelpWidget field to given value.

### HasHelpWidget

`func (o *OrganizationsOrganizationIdBrandingPoliciesHelpSettings) HasHelpWidget() bool`

HasHelpWidget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


