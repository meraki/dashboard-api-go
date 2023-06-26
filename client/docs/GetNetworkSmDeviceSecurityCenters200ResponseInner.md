# GetNetworkSmDeviceSecurityCenters200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsRooted** | Pointer to **bool** | Boolean indicating if the device is rooted. | [optional] 
**HasAntiVirus** | Pointer to **bool** | Boolean indicating if the device has Antivirus. | [optional] 
**AntiVirusName** | Pointer to **string** | The name of the Antivirus. | [optional] 
**IsFireWallEnabled** | Pointer to **bool** | Boolean indicating if the device has a Firewall enabled. | [optional] 
**HasFireWallInstalled** | Pointer to **bool** | Boolean indicating if the device has a Firewall installed. | [optional] 
**FireWallName** | Pointer to **string** | The name of the Firewall. | [optional] 
**IsDiskEncrypted** | Pointer to **bool** | Boolean indicating if the device has disk encryption. | [optional] 
**IsAutoLoginDisabled** | Pointer to **bool** | Boolean indicating if the device has auto login disabled. | [optional] 
**Id** | Pointer to **string** | The Meraki identifier for the security center record. | [optional] 
**RunningProcs** | Pointer to **string** | A comma seperated list of procs running on the device. | [optional] 

## Methods

### NewGetNetworkSmDeviceSecurityCenters200ResponseInner

`func NewGetNetworkSmDeviceSecurityCenters200ResponseInner() *GetNetworkSmDeviceSecurityCenters200ResponseInner`

NewGetNetworkSmDeviceSecurityCenters200ResponseInner instantiates a new GetNetworkSmDeviceSecurityCenters200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSmDeviceSecurityCenters200ResponseInnerWithDefaults

`func NewGetNetworkSmDeviceSecurityCenters200ResponseInnerWithDefaults() *GetNetworkSmDeviceSecurityCenters200ResponseInner`

NewGetNetworkSmDeviceSecurityCenters200ResponseInnerWithDefaults instantiates a new GetNetworkSmDeviceSecurityCenters200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsRooted

`func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) GetIsRooted() bool`

GetIsRooted returns the IsRooted field if non-nil, zero value otherwise.

### GetIsRootedOk

`func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) GetIsRootedOk() (*bool, bool)`

GetIsRootedOk returns a tuple with the IsRooted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRooted

`func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) SetIsRooted(v bool)`

SetIsRooted sets IsRooted field to given value.

### HasIsRooted

`func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) HasIsRooted() bool`

HasIsRooted returns a boolean if a field has been set.

### GetHasAntiVirus

`func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) GetHasAntiVirus() bool`

GetHasAntiVirus returns the HasAntiVirus field if non-nil, zero value otherwise.

### GetHasAntiVirusOk

`func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) GetHasAntiVirusOk() (*bool, bool)`

GetHasAntiVirusOk returns a tuple with the HasAntiVirus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAntiVirus

`func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) SetHasAntiVirus(v bool)`

SetHasAntiVirus sets HasAntiVirus field to given value.

### HasHasAntiVirus

`func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) HasHasAntiVirus() bool`

HasHasAntiVirus returns a boolean if a field has been set.

### GetAntiVirusName

`func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) GetAntiVirusName() string`

GetAntiVirusName returns the AntiVirusName field if non-nil, zero value otherwise.

### GetAntiVirusNameOk

`func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) GetAntiVirusNameOk() (*string, bool)`

GetAntiVirusNameOk returns a tuple with the AntiVirusName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiVirusName

`func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) SetAntiVirusName(v string)`

SetAntiVirusName sets AntiVirusName field to given value.

### HasAntiVirusName

`func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) HasAntiVirusName() bool`

HasAntiVirusName returns a boolean if a field has been set.

### GetIsFireWallEnabled

`func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) GetIsFireWallEnabled() bool`

GetIsFireWallEnabled returns the IsFireWallEnabled field if non-nil, zero value otherwise.

### GetIsFireWallEnabledOk

`func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) GetIsFireWallEnabledOk() (*bool, bool)`

GetIsFireWallEnabledOk returns a tuple with the IsFireWallEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFireWallEnabled

`func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) SetIsFireWallEnabled(v bool)`

SetIsFireWallEnabled sets IsFireWallEnabled field to given value.

### HasIsFireWallEnabled

`func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) HasIsFireWallEnabled() bool`

HasIsFireWallEnabled returns a boolean if a field has been set.

### GetHasFireWallInstalled

`func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) GetHasFireWallInstalled() bool`

GetHasFireWallInstalled returns the HasFireWallInstalled field if non-nil, zero value otherwise.

### GetHasFireWallInstalledOk

`func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) GetHasFireWallInstalledOk() (*bool, bool)`

GetHasFireWallInstalledOk returns a tuple with the HasFireWallInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFireWallInstalled

`func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) SetHasFireWallInstalled(v bool)`

SetHasFireWallInstalled sets HasFireWallInstalled field to given value.

### HasHasFireWallInstalled

`func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) HasHasFireWallInstalled() bool`

HasHasFireWallInstalled returns a boolean if a field has been set.

### GetFireWallName

`func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) GetFireWallName() string`

GetFireWallName returns the FireWallName field if non-nil, zero value otherwise.

### GetFireWallNameOk

`func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) GetFireWallNameOk() (*string, bool)`

GetFireWallNameOk returns a tuple with the FireWallName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFireWallName

`func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) SetFireWallName(v string)`

SetFireWallName sets FireWallName field to given value.

### HasFireWallName

`func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) HasFireWallName() bool`

HasFireWallName returns a boolean if a field has been set.

### GetIsDiskEncrypted

`func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) GetIsDiskEncrypted() bool`

GetIsDiskEncrypted returns the IsDiskEncrypted field if non-nil, zero value otherwise.

### GetIsDiskEncryptedOk

`func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) GetIsDiskEncryptedOk() (*bool, bool)`

GetIsDiskEncryptedOk returns a tuple with the IsDiskEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDiskEncrypted

`func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) SetIsDiskEncrypted(v bool)`

SetIsDiskEncrypted sets IsDiskEncrypted field to given value.

### HasIsDiskEncrypted

`func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) HasIsDiskEncrypted() bool`

HasIsDiskEncrypted returns a boolean if a field has been set.

### GetIsAutoLoginDisabled

`func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) GetIsAutoLoginDisabled() bool`

GetIsAutoLoginDisabled returns the IsAutoLoginDisabled field if non-nil, zero value otherwise.

### GetIsAutoLoginDisabledOk

`func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) GetIsAutoLoginDisabledOk() (*bool, bool)`

GetIsAutoLoginDisabledOk returns a tuple with the IsAutoLoginDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoLoginDisabled

`func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) SetIsAutoLoginDisabled(v bool)`

SetIsAutoLoginDisabled sets IsAutoLoginDisabled field to given value.

### HasIsAutoLoginDisabled

`func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) HasIsAutoLoginDisabled() bool`

HasIsAutoLoginDisabled returns a boolean if a field has been set.

### GetId

`func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRunningProcs

`func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) GetRunningProcs() string`

GetRunningProcs returns the RunningProcs field if non-nil, zero value otherwise.

### GetRunningProcsOk

`func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) GetRunningProcsOk() (*string, bool)`

GetRunningProcsOk returns a tuple with the RunningProcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningProcs

`func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) SetRunningProcs(v string)`

SetRunningProcs sets RunningProcs field to given value.

### HasRunningProcs

`func (o *GetNetworkSmDeviceSecurityCenters200ResponseInner) HasRunningProcs() bool`

HasRunningProcs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


