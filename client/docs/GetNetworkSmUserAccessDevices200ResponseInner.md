# GetNetworkSmUserAccessDevices200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | device ID | [optional] 
**Name** | Pointer to **string** | device name | [optional] 
**SystemType** | Pointer to **string** | system type | [optional] 
**Mac** | Pointer to **string** | mac address | [optional] 
**Username** | Pointer to **string** | username | [optional] 
**Email** | Pointer to **string** | user email | [optional] 
**Tags** | Pointer to **[]string** | device tags | [optional] 
**TrustedAccessConnections** | Pointer to [**[]GetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner**](GetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner.md) | Array of trusted access configs | [optional] 

## Methods

### NewGetNetworkSmUserAccessDevices200ResponseInner

`func NewGetNetworkSmUserAccessDevices200ResponseInner() *GetNetworkSmUserAccessDevices200ResponseInner`

NewGetNetworkSmUserAccessDevices200ResponseInner instantiates a new GetNetworkSmUserAccessDevices200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSmUserAccessDevices200ResponseInnerWithDefaults

`func NewGetNetworkSmUserAccessDevices200ResponseInnerWithDefaults() *GetNetworkSmUserAccessDevices200ResponseInner`

NewGetNetworkSmUserAccessDevices200ResponseInnerWithDefaults instantiates a new GetNetworkSmUserAccessDevices200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkSmUserAccessDevices200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkSmUserAccessDevices200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkSmUserAccessDevices200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkSmUserAccessDevices200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkSmUserAccessDevices200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkSmUserAccessDevices200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkSmUserAccessDevices200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkSmUserAccessDevices200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSystemType

`func (o *GetNetworkSmUserAccessDevices200ResponseInner) GetSystemType() string`

GetSystemType returns the SystemType field if non-nil, zero value otherwise.

### GetSystemTypeOk

`func (o *GetNetworkSmUserAccessDevices200ResponseInner) GetSystemTypeOk() (*string, bool)`

GetSystemTypeOk returns a tuple with the SystemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemType

`func (o *GetNetworkSmUserAccessDevices200ResponseInner) SetSystemType(v string)`

SetSystemType sets SystemType field to given value.

### HasSystemType

`func (o *GetNetworkSmUserAccessDevices200ResponseInner) HasSystemType() bool`

HasSystemType returns a boolean if a field has been set.

### GetMac

`func (o *GetNetworkSmUserAccessDevices200ResponseInner) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *GetNetworkSmUserAccessDevices200ResponseInner) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *GetNetworkSmUserAccessDevices200ResponseInner) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *GetNetworkSmUserAccessDevices200ResponseInner) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetUsername

`func (o *GetNetworkSmUserAccessDevices200ResponseInner) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetNetworkSmUserAccessDevices200ResponseInner) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetNetworkSmUserAccessDevices200ResponseInner) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetNetworkSmUserAccessDevices200ResponseInner) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetEmail

`func (o *GetNetworkSmUserAccessDevices200ResponseInner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetNetworkSmUserAccessDevices200ResponseInner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetNetworkSmUserAccessDevices200ResponseInner) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetNetworkSmUserAccessDevices200ResponseInner) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetTags

`func (o *GetNetworkSmUserAccessDevices200ResponseInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetNetworkSmUserAccessDevices200ResponseInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetNetworkSmUserAccessDevices200ResponseInner) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetNetworkSmUserAccessDevices200ResponseInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTrustedAccessConnections

`func (o *GetNetworkSmUserAccessDevices200ResponseInner) GetTrustedAccessConnections() []GetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner`

GetTrustedAccessConnections returns the TrustedAccessConnections field if non-nil, zero value otherwise.

### GetTrustedAccessConnectionsOk

`func (o *GetNetworkSmUserAccessDevices200ResponseInner) GetTrustedAccessConnectionsOk() (*[]GetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner, bool)`

GetTrustedAccessConnectionsOk returns a tuple with the TrustedAccessConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedAccessConnections

`func (o *GetNetworkSmUserAccessDevices200ResponseInner) SetTrustedAccessConnections(v []GetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner)`

SetTrustedAccessConnections sets TrustedAccessConnections field to given value.

### HasTrustedAccessConnections

`func (o *GetNetworkSmUserAccessDevices200ResponseInner) HasTrustedAccessConnections() bool`

HasTrustedAccessConnections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


