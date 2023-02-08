# CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartLineNumber** | Pointer to **int32** | Starting line VTY number | [optional] 
**EndLineNumber** | Pointer to **int32** | Ending line VTY number | [optional] 
**Authentication** | Pointer to [**CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthentication**](CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthentication.md) |  | [optional] 
**Authorization** | Pointer to [**CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthorization**](CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthorization.md) |  | [optional] 
**AccessList** | Pointer to [**CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessList**](CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessList.md) |  | [optional] 
**RotaryNumber** | Pointer to **int32** | SSH rotary number | [optional] 

## Methods

### NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVty

`func NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVty() *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVty`

NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVty instantiates a new CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyWithDefaults

`func NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyWithDefaults() *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVty`

NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyWithDefaults instantiates a new CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartLineNumber

`func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVty) GetStartLineNumber() int32`

GetStartLineNumber returns the StartLineNumber field if non-nil, zero value otherwise.

### GetStartLineNumberOk

`func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVty) GetStartLineNumberOk() (*int32, bool)`

GetStartLineNumberOk returns a tuple with the StartLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartLineNumber

`func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVty) SetStartLineNumber(v int32)`

SetStartLineNumber sets StartLineNumber field to given value.

### HasStartLineNumber

`func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVty) HasStartLineNumber() bool`

HasStartLineNumber returns a boolean if a field has been set.

### GetEndLineNumber

`func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVty) GetEndLineNumber() int32`

GetEndLineNumber returns the EndLineNumber field if non-nil, zero value otherwise.

### GetEndLineNumberOk

`func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVty) GetEndLineNumberOk() (*int32, bool)`

GetEndLineNumberOk returns a tuple with the EndLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndLineNumber

`func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVty) SetEndLineNumber(v int32)`

SetEndLineNumber sets EndLineNumber field to given value.

### HasEndLineNumber

`func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVty) HasEndLineNumber() bool`

HasEndLineNumber returns a boolean if a field has been set.

### GetAuthentication

`func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVty) GetAuthentication() CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVty) GetAuthenticationOk() (*CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVty) SetAuthentication(v CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVty) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetAuthorization

`func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVty) GetAuthorization() CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthorization`

GetAuthorization returns the Authorization field if non-nil, zero value otherwise.

### GetAuthorizationOk

`func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVty) GetAuthorizationOk() (*CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthorization, bool)`

GetAuthorizationOk returns a tuple with the Authorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorization

`func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVty) SetAuthorization(v CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAuthorization)`

SetAuthorization sets Authorization field to given value.

### HasAuthorization

`func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVty) HasAuthorization() bool`

HasAuthorization returns a boolean if a field has been set.

### GetAccessList

`func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVty) GetAccessList() CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessList`

GetAccessList returns the AccessList field if non-nil, zero value otherwise.

### GetAccessListOk

`func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVty) GetAccessListOk() (*CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessList, bool)`

GetAccessListOk returns a tuple with the AccessList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessList

`func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVty) SetAccessList(v CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVtyAccessList)`

SetAccessList sets AccessList field to given value.

### HasAccessList

`func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVty) HasAccessList() bool`

HasAccessList returns a boolean if a field has been set.

### GetRotaryNumber

`func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVty) GetRotaryNumber() int32`

GetRotaryNumber returns the RotaryNumber field if non-nil, zero value otherwise.

### GetRotaryNumberOk

`func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVty) GetRotaryNumberOk() (*int32, bool)`

GetRotaryNumberOk returns a tuple with the RotaryNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotaryNumber

`func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVty) SetRotaryNumber(v int32)`

SetRotaryNumber sets RotaryNumber field to given value.

### HasRotaryNumber

`func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerVty) HasRotaryNumber() bool`

HasRotaryNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


