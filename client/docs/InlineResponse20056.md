# InlineResponse20056

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

### NewInlineResponse20056

`func NewInlineResponse20056() *InlineResponse20056`

NewInlineResponse20056 instantiates a new InlineResponse20056 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20056WithDefaults

`func NewInlineResponse20056WithDefaults() *InlineResponse20056`

NewInlineResponse20056WithDefaults instantiates a new InlineResponse20056 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsRooted

`func (o *InlineResponse20056) GetIsRooted() bool`

GetIsRooted returns the IsRooted field if non-nil, zero value otherwise.

### GetIsRootedOk

`func (o *InlineResponse20056) GetIsRootedOk() (*bool, bool)`

GetIsRootedOk returns a tuple with the IsRooted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRooted

`func (o *InlineResponse20056) SetIsRooted(v bool)`

SetIsRooted sets IsRooted field to given value.

### HasIsRooted

`func (o *InlineResponse20056) HasIsRooted() bool`

HasIsRooted returns a boolean if a field has been set.

### GetHasAntiVirus

`func (o *InlineResponse20056) GetHasAntiVirus() bool`

GetHasAntiVirus returns the HasAntiVirus field if non-nil, zero value otherwise.

### GetHasAntiVirusOk

`func (o *InlineResponse20056) GetHasAntiVirusOk() (*bool, bool)`

GetHasAntiVirusOk returns a tuple with the HasAntiVirus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAntiVirus

`func (o *InlineResponse20056) SetHasAntiVirus(v bool)`

SetHasAntiVirus sets HasAntiVirus field to given value.

### HasHasAntiVirus

`func (o *InlineResponse20056) HasHasAntiVirus() bool`

HasHasAntiVirus returns a boolean if a field has been set.

### GetAntiVirusName

`func (o *InlineResponse20056) GetAntiVirusName() string`

GetAntiVirusName returns the AntiVirusName field if non-nil, zero value otherwise.

### GetAntiVirusNameOk

`func (o *InlineResponse20056) GetAntiVirusNameOk() (*string, bool)`

GetAntiVirusNameOk returns a tuple with the AntiVirusName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiVirusName

`func (o *InlineResponse20056) SetAntiVirusName(v string)`

SetAntiVirusName sets AntiVirusName field to given value.

### HasAntiVirusName

`func (o *InlineResponse20056) HasAntiVirusName() bool`

HasAntiVirusName returns a boolean if a field has been set.

### GetIsFireWallEnabled

`func (o *InlineResponse20056) GetIsFireWallEnabled() bool`

GetIsFireWallEnabled returns the IsFireWallEnabled field if non-nil, zero value otherwise.

### GetIsFireWallEnabledOk

`func (o *InlineResponse20056) GetIsFireWallEnabledOk() (*bool, bool)`

GetIsFireWallEnabledOk returns a tuple with the IsFireWallEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFireWallEnabled

`func (o *InlineResponse20056) SetIsFireWallEnabled(v bool)`

SetIsFireWallEnabled sets IsFireWallEnabled field to given value.

### HasIsFireWallEnabled

`func (o *InlineResponse20056) HasIsFireWallEnabled() bool`

HasIsFireWallEnabled returns a boolean if a field has been set.

### GetHasFireWallInstalled

`func (o *InlineResponse20056) GetHasFireWallInstalled() bool`

GetHasFireWallInstalled returns the HasFireWallInstalled field if non-nil, zero value otherwise.

### GetHasFireWallInstalledOk

`func (o *InlineResponse20056) GetHasFireWallInstalledOk() (*bool, bool)`

GetHasFireWallInstalledOk returns a tuple with the HasFireWallInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFireWallInstalled

`func (o *InlineResponse20056) SetHasFireWallInstalled(v bool)`

SetHasFireWallInstalled sets HasFireWallInstalled field to given value.

### HasHasFireWallInstalled

`func (o *InlineResponse20056) HasHasFireWallInstalled() bool`

HasHasFireWallInstalled returns a boolean if a field has been set.

### GetFireWallName

`func (o *InlineResponse20056) GetFireWallName() string`

GetFireWallName returns the FireWallName field if non-nil, zero value otherwise.

### GetFireWallNameOk

`func (o *InlineResponse20056) GetFireWallNameOk() (*string, bool)`

GetFireWallNameOk returns a tuple with the FireWallName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFireWallName

`func (o *InlineResponse20056) SetFireWallName(v string)`

SetFireWallName sets FireWallName field to given value.

### HasFireWallName

`func (o *InlineResponse20056) HasFireWallName() bool`

HasFireWallName returns a boolean if a field has been set.

### GetIsDiskEncrypted

`func (o *InlineResponse20056) GetIsDiskEncrypted() bool`

GetIsDiskEncrypted returns the IsDiskEncrypted field if non-nil, zero value otherwise.

### GetIsDiskEncryptedOk

`func (o *InlineResponse20056) GetIsDiskEncryptedOk() (*bool, bool)`

GetIsDiskEncryptedOk returns a tuple with the IsDiskEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDiskEncrypted

`func (o *InlineResponse20056) SetIsDiskEncrypted(v bool)`

SetIsDiskEncrypted sets IsDiskEncrypted field to given value.

### HasIsDiskEncrypted

`func (o *InlineResponse20056) HasIsDiskEncrypted() bool`

HasIsDiskEncrypted returns a boolean if a field has been set.

### GetIsAutoLoginDisabled

`func (o *InlineResponse20056) GetIsAutoLoginDisabled() bool`

GetIsAutoLoginDisabled returns the IsAutoLoginDisabled field if non-nil, zero value otherwise.

### GetIsAutoLoginDisabledOk

`func (o *InlineResponse20056) GetIsAutoLoginDisabledOk() (*bool, bool)`

GetIsAutoLoginDisabledOk returns a tuple with the IsAutoLoginDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoLoginDisabled

`func (o *InlineResponse20056) SetIsAutoLoginDisabled(v bool)`

SetIsAutoLoginDisabled sets IsAutoLoginDisabled field to given value.

### HasIsAutoLoginDisabled

`func (o *InlineResponse20056) HasIsAutoLoginDisabled() bool`

HasIsAutoLoginDisabled returns a boolean if a field has been set.

### GetId

`func (o *InlineResponse20056) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20056) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20056) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20056) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRunningProcs

`func (o *InlineResponse20056) GetRunningProcs() string`

GetRunningProcs returns the RunningProcs field if non-nil, zero value otherwise.

### GetRunningProcsOk

`func (o *InlineResponse20056) GetRunningProcsOk() (*string, bool)`

GetRunningProcsOk returns a tuple with the RunningProcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningProcs

`func (o *InlineResponse20056) SetRunningProcs(v string)`

SetRunningProcs sets RunningProcs field to given value.

### HasRunningProcs

`func (o *InlineResponse20056) HasRunningProcs() bool`

HasRunningProcs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


