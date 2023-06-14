# UpdateOrganizationAlertsProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Is the alert config enabled | [optional] 
**Type** | Pointer to **string** | The alert type | [optional] 
**AlertCondition** | Pointer to [**CreateOrganizationAlertsProfileRequestAlertCondition**](CreateOrganizationAlertsProfileRequestAlertCondition.md) |  | [optional] 
**Recipients** | Pointer to [**CreateOrganizationAlertsProfileRequestRecipients**](CreateOrganizationAlertsProfileRequestRecipients.md) |  | [optional] 
**NetworkTags** | Pointer to **[]string** | Networks with these tags will be monitored for the alert | [optional] 
**Description** | Pointer to **string** | User supplied description of the alert | [optional] 

## Methods

### NewUpdateOrganizationAlertsProfileRequest

`func NewUpdateOrganizationAlertsProfileRequest() *UpdateOrganizationAlertsProfileRequest`

NewUpdateOrganizationAlertsProfileRequest instantiates a new UpdateOrganizationAlertsProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrganizationAlertsProfileRequestWithDefaults

`func NewUpdateOrganizationAlertsProfileRequestWithDefaults() *UpdateOrganizationAlertsProfileRequest`

NewUpdateOrganizationAlertsProfileRequestWithDefaults instantiates a new UpdateOrganizationAlertsProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *UpdateOrganizationAlertsProfileRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateOrganizationAlertsProfileRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateOrganizationAlertsProfileRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateOrganizationAlertsProfileRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetType

`func (o *UpdateOrganizationAlertsProfileRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateOrganizationAlertsProfileRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateOrganizationAlertsProfileRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateOrganizationAlertsProfileRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAlertCondition

`func (o *UpdateOrganizationAlertsProfileRequest) GetAlertCondition() CreateOrganizationAlertsProfileRequestAlertCondition`

GetAlertCondition returns the AlertCondition field if non-nil, zero value otherwise.

### GetAlertConditionOk

`func (o *UpdateOrganizationAlertsProfileRequest) GetAlertConditionOk() (*CreateOrganizationAlertsProfileRequestAlertCondition, bool)`

GetAlertConditionOk returns a tuple with the AlertCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertCondition

`func (o *UpdateOrganizationAlertsProfileRequest) SetAlertCondition(v CreateOrganizationAlertsProfileRequestAlertCondition)`

SetAlertCondition sets AlertCondition field to given value.

### HasAlertCondition

`func (o *UpdateOrganizationAlertsProfileRequest) HasAlertCondition() bool`

HasAlertCondition returns a boolean if a field has been set.

### GetRecipients

`func (o *UpdateOrganizationAlertsProfileRequest) GetRecipients() CreateOrganizationAlertsProfileRequestRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *UpdateOrganizationAlertsProfileRequest) GetRecipientsOk() (*CreateOrganizationAlertsProfileRequestRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *UpdateOrganizationAlertsProfileRequest) SetRecipients(v CreateOrganizationAlertsProfileRequestRecipients)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *UpdateOrganizationAlertsProfileRequest) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetNetworkTags

`func (o *UpdateOrganizationAlertsProfileRequest) GetNetworkTags() []string`

GetNetworkTags returns the NetworkTags field if non-nil, zero value otherwise.

### GetNetworkTagsOk

`func (o *UpdateOrganizationAlertsProfileRequest) GetNetworkTagsOk() (*[]string, bool)`

GetNetworkTagsOk returns a tuple with the NetworkTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTags

`func (o *UpdateOrganizationAlertsProfileRequest) SetNetworkTags(v []string)`

SetNetworkTags sets NetworkTags field to given value.

### HasNetworkTags

`func (o *UpdateOrganizationAlertsProfileRequest) HasNetworkTags() bool`

HasNetworkTags returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateOrganizationAlertsProfileRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateOrganizationAlertsProfileRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateOrganizationAlertsProfileRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateOrganizationAlertsProfileRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


