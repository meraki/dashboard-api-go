# InlineResponse200123

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShortName** | Pointer to **string** | Short name of the early access feature | [optional] 
**Name** | Pointer to **string** | Name of the early access feature | [optional] 
**Descriptions** | Pointer to [**OrganizationsOrganizationIdEarlyAccessFeaturesDescriptions**](OrganizationsOrganizationIdEarlyAccessFeaturesDescriptions.md) |  | [optional] 
**Topic** | Pointer to **string** | Topic of the early access feature | [optional] 
**IsOrgScopedOnly** | Pointer to **bool** | If this early access feature can only be opted in for the entire organization | [optional] 
**DocumentationLink** | Pointer to **string** | Link to the documentation of this early access feature | [optional] 
**SupportLink** | Pointer to **string** | Link to get support for this early access feature | [optional] 

## Methods

### NewInlineResponse200123

`func NewInlineResponse200123() *InlineResponse200123`

NewInlineResponse200123 instantiates a new InlineResponse200123 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200123WithDefaults

`func NewInlineResponse200123WithDefaults() *InlineResponse200123`

NewInlineResponse200123WithDefaults instantiates a new InlineResponse200123 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShortName

`func (o *InlineResponse200123) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *InlineResponse200123) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *InlineResponse200123) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *InlineResponse200123) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200123) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200123) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200123) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200123) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescriptions

`func (o *InlineResponse200123) GetDescriptions() OrganizationsOrganizationIdEarlyAccessFeaturesDescriptions`

GetDescriptions returns the Descriptions field if non-nil, zero value otherwise.

### GetDescriptionsOk

`func (o *InlineResponse200123) GetDescriptionsOk() (*OrganizationsOrganizationIdEarlyAccessFeaturesDescriptions, bool)`

GetDescriptionsOk returns a tuple with the Descriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptions

`func (o *InlineResponse200123) SetDescriptions(v OrganizationsOrganizationIdEarlyAccessFeaturesDescriptions)`

SetDescriptions sets Descriptions field to given value.

### HasDescriptions

`func (o *InlineResponse200123) HasDescriptions() bool`

HasDescriptions returns a boolean if a field has been set.

### GetTopic

`func (o *InlineResponse200123) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *InlineResponse200123) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *InlineResponse200123) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *InlineResponse200123) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetIsOrgScopedOnly

`func (o *InlineResponse200123) GetIsOrgScopedOnly() bool`

GetIsOrgScopedOnly returns the IsOrgScopedOnly field if non-nil, zero value otherwise.

### GetIsOrgScopedOnlyOk

`func (o *InlineResponse200123) GetIsOrgScopedOnlyOk() (*bool, bool)`

GetIsOrgScopedOnlyOk returns a tuple with the IsOrgScopedOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrgScopedOnly

`func (o *InlineResponse200123) SetIsOrgScopedOnly(v bool)`

SetIsOrgScopedOnly sets IsOrgScopedOnly field to given value.

### HasIsOrgScopedOnly

`func (o *InlineResponse200123) HasIsOrgScopedOnly() bool`

HasIsOrgScopedOnly returns a boolean if a field has been set.

### GetDocumentationLink

`func (o *InlineResponse200123) GetDocumentationLink() string`

GetDocumentationLink returns the DocumentationLink field if non-nil, zero value otherwise.

### GetDocumentationLinkOk

`func (o *InlineResponse200123) GetDocumentationLinkOk() (*string, bool)`

GetDocumentationLinkOk returns a tuple with the DocumentationLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationLink

`func (o *InlineResponse200123) SetDocumentationLink(v string)`

SetDocumentationLink sets DocumentationLink field to given value.

### HasDocumentationLink

`func (o *InlineResponse200123) HasDocumentationLink() bool`

HasDocumentationLink returns a boolean if a field has been set.

### GetSupportLink

`func (o *InlineResponse200123) GetSupportLink() string`

GetSupportLink returns the SupportLink field if non-nil, zero value otherwise.

### GetSupportLinkOk

`func (o *InlineResponse200123) GetSupportLinkOk() (*string, bool)`

GetSupportLinkOk returns a tuple with the SupportLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLink

`func (o *InlineResponse200123) SetSupportLink(v string)`

SetSupportLink sets SupportLink field to given value.

### HasSupportLink

`func (o *InlineResponse200123) HasSupportLink() bool`

HasSupportLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


