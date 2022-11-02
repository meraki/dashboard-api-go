# GetNetworkSmDeviceDeviceProfiles200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **string** | The Meraki managed device Id. | [optional] 
**Id** | Pointer to **string** | The numerical Meraki Id of the profile. | [optional] 
**IsEncrypted** | Pointer to **bool** | A boolean indicating if the profile is encrypted. | [optional] 
**IsManaged** | Pointer to **bool** | Whether or not the profile is managed by Meraki. | [optional] 
**ProfileData** | Pointer to **string** | A string containing a JSON object with the profile data. | [optional] 
**ProfileIdentifier** | Pointer to **string** | The identifier of the profile. | [optional] 
**Name** | Pointer to **string** | The name of the profile. | [optional] 
**Version** | Pointer to **string** | The verison of the profile. | [optional] 

## Methods

### NewGetNetworkSmDeviceDeviceProfiles200ResponseInner

`func NewGetNetworkSmDeviceDeviceProfiles200ResponseInner() *GetNetworkSmDeviceDeviceProfiles200ResponseInner`

NewGetNetworkSmDeviceDeviceProfiles200ResponseInner instantiates a new GetNetworkSmDeviceDeviceProfiles200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSmDeviceDeviceProfiles200ResponseInnerWithDefaults

`func NewGetNetworkSmDeviceDeviceProfiles200ResponseInnerWithDefaults() *GetNetworkSmDeviceDeviceProfiles200ResponseInner`

NewGetNetworkSmDeviceDeviceProfiles200ResponseInnerWithDefaults instantiates a new GetNetworkSmDeviceDeviceProfiles200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetId

`func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsEncrypted

`func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) GetIsEncrypted() bool`

GetIsEncrypted returns the IsEncrypted field if non-nil, zero value otherwise.

### GetIsEncryptedOk

`func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) GetIsEncryptedOk() (*bool, bool)`

GetIsEncryptedOk returns a tuple with the IsEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEncrypted

`func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) SetIsEncrypted(v bool)`

SetIsEncrypted sets IsEncrypted field to given value.

### HasIsEncrypted

`func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) HasIsEncrypted() bool`

HasIsEncrypted returns a boolean if a field has been set.

### GetIsManaged

`func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) GetIsManaged() bool`

GetIsManaged returns the IsManaged field if non-nil, zero value otherwise.

### GetIsManagedOk

`func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) GetIsManagedOk() (*bool, bool)`

GetIsManagedOk returns a tuple with the IsManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManaged

`func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) SetIsManaged(v bool)`

SetIsManaged sets IsManaged field to given value.

### HasIsManaged

`func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) HasIsManaged() bool`

HasIsManaged returns a boolean if a field has been set.

### GetProfileData

`func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) GetProfileData() string`

GetProfileData returns the ProfileData field if non-nil, zero value otherwise.

### GetProfileDataOk

`func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) GetProfileDataOk() (*string, bool)`

GetProfileDataOk returns a tuple with the ProfileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileData

`func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) SetProfileData(v string)`

SetProfileData sets ProfileData field to given value.

### HasProfileData

`func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) HasProfileData() bool`

HasProfileData returns a boolean if a field has been set.

### GetProfileIdentifier

`func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) GetProfileIdentifier() string`

GetProfileIdentifier returns the ProfileIdentifier field if non-nil, zero value otherwise.

### GetProfileIdentifierOk

`func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) GetProfileIdentifierOk() (*string, bool)`

GetProfileIdentifierOk returns a tuple with the ProfileIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIdentifier

`func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) SetProfileIdentifier(v string)`

SetProfileIdentifier sets ProfileIdentifier field to given value.

### HasProfileIdentifier

`func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) HasProfileIdentifier() bool`

HasProfileIdentifier returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


