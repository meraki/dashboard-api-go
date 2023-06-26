# GetAdministeredIdentitiesMe200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Username | [optional] 
**Email** | Pointer to **string** | User email | [optional] 
**LastUsedDashboardAt** | Pointer to **time.Time** | Last seen active on Dashboard UI | [optional] 
**Authentication** | Pointer to [**GetAdministeredIdentitiesMe200ResponseAuthentication**](GetAdministeredIdentitiesMe200ResponseAuthentication.md) |  | [optional] 

## Methods

### NewGetAdministeredIdentitiesMe200Response

`func NewGetAdministeredIdentitiesMe200Response() *GetAdministeredIdentitiesMe200Response`

NewGetAdministeredIdentitiesMe200Response instantiates a new GetAdministeredIdentitiesMe200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAdministeredIdentitiesMe200ResponseWithDefaults

`func NewGetAdministeredIdentitiesMe200ResponseWithDefaults() *GetAdministeredIdentitiesMe200Response`

NewGetAdministeredIdentitiesMe200ResponseWithDefaults instantiates a new GetAdministeredIdentitiesMe200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetAdministeredIdentitiesMe200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAdministeredIdentitiesMe200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAdministeredIdentitiesMe200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetAdministeredIdentitiesMe200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *GetAdministeredIdentitiesMe200Response) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetAdministeredIdentitiesMe200Response) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetAdministeredIdentitiesMe200Response) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetAdministeredIdentitiesMe200Response) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetLastUsedDashboardAt

`func (o *GetAdministeredIdentitiesMe200Response) GetLastUsedDashboardAt() time.Time`

GetLastUsedDashboardAt returns the LastUsedDashboardAt field if non-nil, zero value otherwise.

### GetLastUsedDashboardAtOk

`func (o *GetAdministeredIdentitiesMe200Response) GetLastUsedDashboardAtOk() (*time.Time, bool)`

GetLastUsedDashboardAtOk returns a tuple with the LastUsedDashboardAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedDashboardAt

`func (o *GetAdministeredIdentitiesMe200Response) SetLastUsedDashboardAt(v time.Time)`

SetLastUsedDashboardAt sets LastUsedDashboardAt field to given value.

### HasLastUsedDashboardAt

`func (o *GetAdministeredIdentitiesMe200Response) HasLastUsedDashboardAt() bool`

HasLastUsedDashboardAt returns a boolean if a field has been set.

### GetAuthentication

`func (o *GetAdministeredIdentitiesMe200Response) GetAuthentication() GetAdministeredIdentitiesMe200ResponseAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *GetAdministeredIdentitiesMe200Response) GetAuthenticationOk() (*GetAdministeredIdentitiesMe200ResponseAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *GetAdministeredIdentitiesMe200Response) SetAuthentication(v GetAdministeredIdentitiesMe200ResponseAuthentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *GetAdministeredIdentitiesMe200Response) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


