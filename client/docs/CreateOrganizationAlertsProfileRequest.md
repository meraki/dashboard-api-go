# CreateOrganizationAlertsProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The alert type | 
**AlertCondition** | [**CreateOrganizationAlertsProfileRequestAlertCondition**](CreateOrganizationAlertsProfileRequestAlertCondition.md) |  | 
**Recipients** | [**CreateOrganizationAlertsProfileRequestRecipients**](CreateOrganizationAlertsProfileRequestRecipients.md) |  | 
**NetworkTags** | **[]string** | Networks with these tags will be monitored for the alert | 
**Description** | Pointer to **string** | User supplied description of the alert | [optional] 

## Methods

### NewCreateOrganizationAlertsProfileRequest

`func NewCreateOrganizationAlertsProfileRequest(type_ string, alertCondition CreateOrganizationAlertsProfileRequestAlertCondition, recipients CreateOrganizationAlertsProfileRequestRecipients, networkTags []string, ) *CreateOrganizationAlertsProfileRequest`

NewCreateOrganizationAlertsProfileRequest instantiates a new CreateOrganizationAlertsProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationAlertsProfileRequestWithDefaults

`func NewCreateOrganizationAlertsProfileRequestWithDefaults() *CreateOrganizationAlertsProfileRequest`

NewCreateOrganizationAlertsProfileRequestWithDefaults instantiates a new CreateOrganizationAlertsProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateOrganizationAlertsProfileRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateOrganizationAlertsProfileRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateOrganizationAlertsProfileRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAlertCondition

`func (o *CreateOrganizationAlertsProfileRequest) GetAlertCondition() CreateOrganizationAlertsProfileRequestAlertCondition`

GetAlertCondition returns the AlertCondition field if non-nil, zero value otherwise.

### GetAlertConditionOk

`func (o *CreateOrganizationAlertsProfileRequest) GetAlertConditionOk() (*CreateOrganizationAlertsProfileRequestAlertCondition, bool)`

GetAlertConditionOk returns a tuple with the AlertCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertCondition

`func (o *CreateOrganizationAlertsProfileRequest) SetAlertCondition(v CreateOrganizationAlertsProfileRequestAlertCondition)`

SetAlertCondition sets AlertCondition field to given value.


### GetRecipients

`func (o *CreateOrganizationAlertsProfileRequest) GetRecipients() CreateOrganizationAlertsProfileRequestRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *CreateOrganizationAlertsProfileRequest) GetRecipientsOk() (*CreateOrganizationAlertsProfileRequestRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *CreateOrganizationAlertsProfileRequest) SetRecipients(v CreateOrganizationAlertsProfileRequestRecipients)`

SetRecipients sets Recipients field to given value.


### GetNetworkTags

`func (o *CreateOrganizationAlertsProfileRequest) GetNetworkTags() []string`

GetNetworkTags returns the NetworkTags field if non-nil, zero value otherwise.

### GetNetworkTagsOk

`func (o *CreateOrganizationAlertsProfileRequest) GetNetworkTagsOk() (*[]string, bool)`

GetNetworkTagsOk returns a tuple with the NetworkTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTags

`func (o *CreateOrganizationAlertsProfileRequest) SetNetworkTags(v []string)`

SetNetworkTags sets NetworkTags field to given value.


### GetDescription

`func (o *CreateOrganizationAlertsProfileRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateOrganizationAlertsProfileRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateOrganizationAlertsProfileRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateOrganizationAlertsProfileRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


