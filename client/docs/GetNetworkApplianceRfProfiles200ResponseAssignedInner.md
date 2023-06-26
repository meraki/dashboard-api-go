# GetNetworkApplianceRfProfiles200ResponseAssignedInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the RF Profile. | [optional] 
**NetworkId** | Pointer to **string** | ID of network this RF Profile belongs in. | [optional] 
**Name** | Pointer to **string** | The name of the profile. | [optional] 
**TwoFourGhzSettings** | Pointer to [**GetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings**](GetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings.md) |  | [optional] 
**FiveGhzSettings** | Pointer to [**GetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings**](GetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings.md) |  | [optional] 
**PerSsidSettings** | Pointer to [**GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings**](GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings.md) |  | [optional] 

## Methods

### NewGetNetworkApplianceRfProfiles200ResponseAssignedInner

`func NewGetNetworkApplianceRfProfiles200ResponseAssignedInner() *GetNetworkApplianceRfProfiles200ResponseAssignedInner`

NewGetNetworkApplianceRfProfiles200ResponseAssignedInner instantiates a new GetNetworkApplianceRfProfiles200ResponseAssignedInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkApplianceRfProfiles200ResponseAssignedInnerWithDefaults

`func NewGetNetworkApplianceRfProfiles200ResponseAssignedInnerWithDefaults() *GetNetworkApplianceRfProfiles200ResponseAssignedInner`

NewGetNetworkApplianceRfProfiles200ResponseAssignedInnerWithDefaults instantiates a new GetNetworkApplianceRfProfiles200ResponseAssignedInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNetworkId

`func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTwoFourGhzSettings

`func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) GetTwoFourGhzSettings() GetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings`

GetTwoFourGhzSettings returns the TwoFourGhzSettings field if non-nil, zero value otherwise.

### GetTwoFourGhzSettingsOk

`func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) GetTwoFourGhzSettingsOk() (*GetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings, bool)`

GetTwoFourGhzSettingsOk returns a tuple with the TwoFourGhzSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFourGhzSettings

`func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) SetTwoFourGhzSettings(v GetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings)`

SetTwoFourGhzSettings sets TwoFourGhzSettings field to given value.

### HasTwoFourGhzSettings

`func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) HasTwoFourGhzSettings() bool`

HasTwoFourGhzSettings returns a boolean if a field has been set.

### GetFiveGhzSettings

`func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) GetFiveGhzSettings() GetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings`

GetFiveGhzSettings returns the FiveGhzSettings field if non-nil, zero value otherwise.

### GetFiveGhzSettingsOk

`func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) GetFiveGhzSettingsOk() (*GetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings, bool)`

GetFiveGhzSettingsOk returns a tuple with the FiveGhzSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiveGhzSettings

`func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) SetFiveGhzSettings(v GetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings)`

SetFiveGhzSettings sets FiveGhzSettings field to given value.

### HasFiveGhzSettings

`func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) HasFiveGhzSettings() bool`

HasFiveGhzSettings returns a boolean if a field has been set.

### GetPerSsidSettings

`func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) GetPerSsidSettings() GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings`

GetPerSsidSettings returns the PerSsidSettings field if non-nil, zero value otherwise.

### GetPerSsidSettingsOk

`func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) GetPerSsidSettingsOk() (*GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings, bool)`

GetPerSsidSettingsOk returns a tuple with the PerSsidSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerSsidSettings

`func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) SetPerSsidSettings(v GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings)`

SetPerSsidSettings sets PerSsidSettings field to given value.

### HasPerSsidSettings

`func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) HasPerSsidSettings() bool`

HasPerSsidSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


