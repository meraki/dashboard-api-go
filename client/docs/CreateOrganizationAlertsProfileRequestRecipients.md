# CreateOrganizationAlertsProfileRequestRecipients

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emails** | Pointer to **[]string** | A list of emails that will receive information about the alert | [optional] 
**HttpServerIds** | Pointer to **[]string** | A list base64 encoded urls of webhook endpoints that will receive information about the alert | [optional] 

## Methods

### NewCreateOrganizationAlertsProfileRequestRecipients

`func NewCreateOrganizationAlertsProfileRequestRecipients() *CreateOrganizationAlertsProfileRequestRecipients`

NewCreateOrganizationAlertsProfileRequestRecipients instantiates a new CreateOrganizationAlertsProfileRequestRecipients object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationAlertsProfileRequestRecipientsWithDefaults

`func NewCreateOrganizationAlertsProfileRequestRecipientsWithDefaults() *CreateOrganizationAlertsProfileRequestRecipients`

NewCreateOrganizationAlertsProfileRequestRecipientsWithDefaults instantiates a new CreateOrganizationAlertsProfileRequestRecipients object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmails

`func (o *CreateOrganizationAlertsProfileRequestRecipients) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *CreateOrganizationAlertsProfileRequestRecipients) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *CreateOrganizationAlertsProfileRequestRecipients) SetEmails(v []string)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *CreateOrganizationAlertsProfileRequestRecipients) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetHttpServerIds

`func (o *CreateOrganizationAlertsProfileRequestRecipients) GetHttpServerIds() []string`

GetHttpServerIds returns the HttpServerIds field if non-nil, zero value otherwise.

### GetHttpServerIdsOk

`func (o *CreateOrganizationAlertsProfileRequestRecipients) GetHttpServerIdsOk() (*[]string, bool)`

GetHttpServerIdsOk returns a tuple with the HttpServerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpServerIds

`func (o *CreateOrganizationAlertsProfileRequestRecipients) SetHttpServerIds(v []string)`

SetHttpServerIds sets HttpServerIds field to given value.

### HasHttpServerIds

`func (o *CreateOrganizationAlertsProfileRequestRecipients) HasHttpServerIds() bool`

HasHttpServerIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


