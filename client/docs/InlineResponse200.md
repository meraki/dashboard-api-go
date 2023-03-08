# InlineResponse200

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Username | [optional] 
**Email** | Pointer to **string** | User email | [optional] 
**LastUsedDashboardAt** | Pointer to **time.Time** | Last seen active on Dashboard UI | [optional] 
**Authentication** | Pointer to [**InlineResponse200Authentication**](InlineResponse200Authentication.md) |  | [optional] 

## Methods

### NewInlineResponse200

`func NewInlineResponse200() *InlineResponse200`

NewInlineResponse200 instantiates a new InlineResponse200 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200WithDefaults

`func NewInlineResponse200WithDefaults() *InlineResponse200`

NewInlineResponse200WithDefaults instantiates a new InlineResponse200 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *InlineResponse200) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InlineResponse200) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InlineResponse200) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *InlineResponse200) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetLastUsedDashboardAt

`func (o *InlineResponse200) GetLastUsedDashboardAt() time.Time`

GetLastUsedDashboardAt returns the LastUsedDashboardAt field if non-nil, zero value otherwise.

### GetLastUsedDashboardAtOk

`func (o *InlineResponse200) GetLastUsedDashboardAtOk() (*time.Time, bool)`

GetLastUsedDashboardAtOk returns a tuple with the LastUsedDashboardAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedDashboardAt

`func (o *InlineResponse200) SetLastUsedDashboardAt(v time.Time)`

SetLastUsedDashboardAt sets LastUsedDashboardAt field to given value.

### HasLastUsedDashboardAt

`func (o *InlineResponse200) HasLastUsedDashboardAt() bool`

HasLastUsedDashboardAt returns a boolean if a field has been set.

### GetAuthentication

`func (o *InlineResponse200) GetAuthentication() InlineResponse200Authentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *InlineResponse200) GetAuthenticationOk() (*InlineResponse200Authentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *InlineResponse200) SetAuthentication(v InlineResponse200Authentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *InlineResponse200) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


