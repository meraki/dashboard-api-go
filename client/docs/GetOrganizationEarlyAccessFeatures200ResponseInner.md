# GetOrganizationEarlyAccessFeatures200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShortName** | Pointer to **string** | Short name of the early access feature | [optional] 
**Name** | Pointer to **string** | Name of the early access feature | [optional] 
**Descriptions** | Pointer to [**GetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions**](GetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions.md) |  | [optional] 
**Topic** | Pointer to **string** | Topic of the early access feature | [optional] 
**IsOrgScopedOnly** | Pointer to **bool** | If this early access feature can only be opted in for the entire organization | [optional] 
**DocumentationLink** | Pointer to **string** | Link to the documentation of this early access feature | [optional] 
**SupportLink** | Pointer to **string** | Link to get support for this early access feature | [optional] 

## Methods

### NewGetOrganizationEarlyAccessFeatures200ResponseInner

`func NewGetOrganizationEarlyAccessFeatures200ResponseInner() *GetOrganizationEarlyAccessFeatures200ResponseInner`

NewGetOrganizationEarlyAccessFeatures200ResponseInner instantiates a new GetOrganizationEarlyAccessFeatures200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationEarlyAccessFeatures200ResponseInnerWithDefaults

`func NewGetOrganizationEarlyAccessFeatures200ResponseInnerWithDefaults() *GetOrganizationEarlyAccessFeatures200ResponseInner`

NewGetOrganizationEarlyAccessFeatures200ResponseInnerWithDefaults instantiates a new GetOrganizationEarlyAccessFeatures200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShortName

`func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetName

`func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescriptions

`func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) GetDescriptions() GetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions`

GetDescriptions returns the Descriptions field if non-nil, zero value otherwise.

### GetDescriptionsOk

`func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) GetDescriptionsOk() (*GetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions, bool)`

GetDescriptionsOk returns a tuple with the Descriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptions

`func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) SetDescriptions(v GetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions)`

SetDescriptions sets Descriptions field to given value.

### HasDescriptions

`func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) HasDescriptions() bool`

HasDescriptions returns a boolean if a field has been set.

### GetTopic

`func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetIsOrgScopedOnly

`func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) GetIsOrgScopedOnly() bool`

GetIsOrgScopedOnly returns the IsOrgScopedOnly field if non-nil, zero value otherwise.

### GetIsOrgScopedOnlyOk

`func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) GetIsOrgScopedOnlyOk() (*bool, bool)`

GetIsOrgScopedOnlyOk returns a tuple with the IsOrgScopedOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrgScopedOnly

`func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) SetIsOrgScopedOnly(v bool)`

SetIsOrgScopedOnly sets IsOrgScopedOnly field to given value.

### HasIsOrgScopedOnly

`func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) HasIsOrgScopedOnly() bool`

HasIsOrgScopedOnly returns a boolean if a field has been set.

### GetDocumentationLink

`func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) GetDocumentationLink() string`

GetDocumentationLink returns the DocumentationLink field if non-nil, zero value otherwise.

### GetDocumentationLinkOk

`func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) GetDocumentationLinkOk() (*string, bool)`

GetDocumentationLinkOk returns a tuple with the DocumentationLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationLink

`func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) SetDocumentationLink(v string)`

SetDocumentationLink sets DocumentationLink field to given value.

### HasDocumentationLink

`func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) HasDocumentationLink() bool`

HasDocumentationLink returns a boolean if a field has been set.

### GetSupportLink

`func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) GetSupportLink() string`

GetSupportLink returns the SupportLink field if non-nil, zero value otherwise.

### GetSupportLinkOk

`func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) GetSupportLinkOk() (*string, bool)`

GetSupportLinkOk returns a tuple with the SupportLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLink

`func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) SetSupportLink(v string)`

SetSupportLink sets SupportLink field to given value.

### HasSupportLink

`func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) HasSupportLink() bool`

HasSupportLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


