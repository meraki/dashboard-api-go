# InlineObject97

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | One of \&quot;delete\&quot; or \&quot;restrict processing\&quot; | [optional] 
**Datasets** | Pointer to **[]string** | The datasets related to the provided key that should be deleted. Only applies to \&quot;delete\&quot; requests. The value \&quot;all\&quot; will be expanded to all datasets applicable to this type. The datasets by applicable to each type are: mac (usage, events, traffic), email (users, loginAttempts), username (users, loginAttempts), bluetoothMac (client, connectivity), smDeviceId (device), smUserId (user) | [optional] 
**Username** | Pointer to **string** | The username of a network log in. Only applies to \&quot;delete\&quot; requests. | [optional] 
**Email** | Pointer to **string** | The email of a network user account. Only applies to \&quot;delete\&quot; requests. | [optional] 
**Mac** | Pointer to **string** | The MAC of a network client device. Applies to both \&quot;restrict processing\&quot; and \&quot;delete\&quot; requests. | [optional] 
**SmDeviceId** | Pointer to **string** | The sm_device_id of a Systems Manager device. The only way to \&quot;restrict processing\&quot; or \&quot;delete\&quot; a Systems Manager device. Must include \&quot;device\&quot; in the dataset for a \&quot;delete\&quot; request to destroy the device. | [optional] 
**SmUserId** | Pointer to **string** | The sm_user_id of a Systems Manager user. The only way to \&quot;restrict processing\&quot; or \&quot;delete\&quot; a Systems Manager user. Must include \&quot;user\&quot; in the dataset for a \&quot;delete\&quot; request to destroy the user. | [optional] 

## Methods

### NewInlineObject97

`func NewInlineObject97() *InlineObject97`

NewInlineObject97 instantiates a new InlineObject97 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject97WithDefaults

`func NewInlineObject97WithDefaults() *InlineObject97`

NewInlineObject97WithDefaults instantiates a new InlineObject97 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InlineObject97) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineObject97) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineObject97) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineObject97) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDatasets

`func (o *InlineObject97) GetDatasets() []string`

GetDatasets returns the Datasets field if non-nil, zero value otherwise.

### GetDatasetsOk

`func (o *InlineObject97) GetDatasetsOk() (*[]string, bool)`

GetDatasetsOk returns a tuple with the Datasets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasets

`func (o *InlineObject97) SetDatasets(v []string)`

SetDatasets sets Datasets field to given value.

### HasDatasets

`func (o *InlineObject97) HasDatasets() bool`

HasDatasets returns a boolean if a field has been set.

### GetUsername

`func (o *InlineObject97) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *InlineObject97) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *InlineObject97) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *InlineObject97) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetEmail

`func (o *InlineObject97) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InlineObject97) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InlineObject97) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *InlineObject97) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetMac

`func (o *InlineObject97) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineObject97) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineObject97) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineObject97) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetSmDeviceId

`func (o *InlineObject97) GetSmDeviceId() string`

GetSmDeviceId returns the SmDeviceId field if non-nil, zero value otherwise.

### GetSmDeviceIdOk

`func (o *InlineObject97) GetSmDeviceIdOk() (*string, bool)`

GetSmDeviceIdOk returns a tuple with the SmDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmDeviceId

`func (o *InlineObject97) SetSmDeviceId(v string)`

SetSmDeviceId sets SmDeviceId field to given value.

### HasSmDeviceId

`func (o *InlineObject97) HasSmDeviceId() bool`

HasSmDeviceId returns a boolean if a field has been set.

### GetSmUserId

`func (o *InlineObject97) GetSmUserId() string`

GetSmUserId returns the SmUserId field if non-nil, zero value otherwise.

### GetSmUserIdOk

`func (o *InlineObject97) GetSmUserIdOk() (*string, bool)`

GetSmUserIdOk returns a tuple with the SmUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmUserId

`func (o *InlineObject97) SetSmUserId(v string)`

SetSmUserId sets SmUserId field to given value.

### HasSmUserId

`func (o *InlineObject97) HasSmUserId() bool`

HasSmUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


