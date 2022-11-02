# GetNetworkFirmwareUpgrades200ResponseProductsWireless

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentVersion** | Pointer to [**GetNetworkFirmwareUpgrades200ResponseProductsWirelessCurrentVersion**](GetNetworkFirmwareUpgrades200ResponseProductsWirelessCurrentVersion.md) |  | [optional] 
**LastUpgrade** | Pointer to [**GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade**](GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade.md) |  | [optional] 
**NextUpgrade** | Pointer to [**GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade**](GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade.md) |  | [optional] 
**AvailableVersions** | Pointer to [**[]GetNetworkFirmwareUpgrades200ResponseProductsWirelessAvailableVersionsInner**](GetNetworkFirmwareUpgrades200ResponseProductsWirelessAvailableVersionsInner.md) | Firmware versions available for upgrade | [optional] 
**ParticipateInNextBetaRelease** | Pointer to **bool** | Whether or not the network wants beta firmware | [optional] 

## Methods

### NewGetNetworkFirmwareUpgrades200ResponseProductsWireless

`func NewGetNetworkFirmwareUpgrades200ResponseProductsWireless() *GetNetworkFirmwareUpgrades200ResponseProductsWireless`

NewGetNetworkFirmwareUpgrades200ResponseProductsWireless instantiates a new GetNetworkFirmwareUpgrades200ResponseProductsWireless object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkFirmwareUpgrades200ResponseProductsWirelessWithDefaults

`func NewGetNetworkFirmwareUpgrades200ResponseProductsWirelessWithDefaults() *GetNetworkFirmwareUpgrades200ResponseProductsWireless`

NewGetNetworkFirmwareUpgrades200ResponseProductsWirelessWithDefaults instantiates a new GetNetworkFirmwareUpgrades200ResponseProductsWireless object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentVersion

`func (o *GetNetworkFirmwareUpgrades200ResponseProductsWireless) GetCurrentVersion() GetNetworkFirmwareUpgrades200ResponseProductsWirelessCurrentVersion`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *GetNetworkFirmwareUpgrades200ResponseProductsWireless) GetCurrentVersionOk() (*GetNetworkFirmwareUpgrades200ResponseProductsWirelessCurrentVersion, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *GetNetworkFirmwareUpgrades200ResponseProductsWireless) SetCurrentVersion(v GetNetworkFirmwareUpgrades200ResponseProductsWirelessCurrentVersion)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *GetNetworkFirmwareUpgrades200ResponseProductsWireless) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetLastUpgrade

`func (o *GetNetworkFirmwareUpgrades200ResponseProductsWireless) GetLastUpgrade() GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade`

GetLastUpgrade returns the LastUpgrade field if non-nil, zero value otherwise.

### GetLastUpgradeOk

`func (o *GetNetworkFirmwareUpgrades200ResponseProductsWireless) GetLastUpgradeOk() (*GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade, bool)`

GetLastUpgradeOk returns a tuple with the LastUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpgrade

`func (o *GetNetworkFirmwareUpgrades200ResponseProductsWireless) SetLastUpgrade(v GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade)`

SetLastUpgrade sets LastUpgrade field to given value.

### HasLastUpgrade

`func (o *GetNetworkFirmwareUpgrades200ResponseProductsWireless) HasLastUpgrade() bool`

HasLastUpgrade returns a boolean if a field has been set.

### GetNextUpgrade

`func (o *GetNetworkFirmwareUpgrades200ResponseProductsWireless) GetNextUpgrade() GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade`

GetNextUpgrade returns the NextUpgrade field if non-nil, zero value otherwise.

### GetNextUpgradeOk

`func (o *GetNetworkFirmwareUpgrades200ResponseProductsWireless) GetNextUpgradeOk() (*GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade, bool)`

GetNextUpgradeOk returns a tuple with the NextUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextUpgrade

`func (o *GetNetworkFirmwareUpgrades200ResponseProductsWireless) SetNextUpgrade(v GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade)`

SetNextUpgrade sets NextUpgrade field to given value.

### HasNextUpgrade

`func (o *GetNetworkFirmwareUpgrades200ResponseProductsWireless) HasNextUpgrade() bool`

HasNextUpgrade returns a boolean if a field has been set.

### GetAvailableVersions

`func (o *GetNetworkFirmwareUpgrades200ResponseProductsWireless) GetAvailableVersions() []GetNetworkFirmwareUpgrades200ResponseProductsWirelessAvailableVersionsInner`

GetAvailableVersions returns the AvailableVersions field if non-nil, zero value otherwise.

### GetAvailableVersionsOk

`func (o *GetNetworkFirmwareUpgrades200ResponseProductsWireless) GetAvailableVersionsOk() (*[]GetNetworkFirmwareUpgrades200ResponseProductsWirelessAvailableVersionsInner, bool)`

GetAvailableVersionsOk returns a tuple with the AvailableVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableVersions

`func (o *GetNetworkFirmwareUpgrades200ResponseProductsWireless) SetAvailableVersions(v []GetNetworkFirmwareUpgrades200ResponseProductsWirelessAvailableVersionsInner)`

SetAvailableVersions sets AvailableVersions field to given value.

### HasAvailableVersions

`func (o *GetNetworkFirmwareUpgrades200ResponseProductsWireless) HasAvailableVersions() bool`

HasAvailableVersions returns a boolean if a field has been set.

### GetParticipateInNextBetaRelease

`func (o *GetNetworkFirmwareUpgrades200ResponseProductsWireless) GetParticipateInNextBetaRelease() bool`

GetParticipateInNextBetaRelease returns the ParticipateInNextBetaRelease field if non-nil, zero value otherwise.

### GetParticipateInNextBetaReleaseOk

`func (o *GetNetworkFirmwareUpgrades200ResponseProductsWireless) GetParticipateInNextBetaReleaseOk() (*bool, bool)`

GetParticipateInNextBetaReleaseOk returns a tuple with the ParticipateInNextBetaRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipateInNextBetaRelease

`func (o *GetNetworkFirmwareUpgrades200ResponseProductsWireless) SetParticipateInNextBetaRelease(v bool)`

SetParticipateInNextBetaRelease sets ParticipateInNextBetaRelease field to given value.

### HasParticipateInNextBetaRelease

`func (o *GetNetworkFirmwareUpgrades200ResponseProductsWireless) HasParticipateInNextBetaRelease() bool`

HasParticipateInNextBetaRelease returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


