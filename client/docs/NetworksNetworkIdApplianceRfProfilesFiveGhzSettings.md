# NetworksNetworkIdApplianceRfProfilesFiveGhzSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinBitrate** | Pointer to **int32** | Sets min bitrate (Mbps) of 5Ghz band. Can be one of &#39;6&#39;, &#39;9&#39;, &#39;12&#39;, &#39;18&#39;, &#39;24&#39;, &#39;36&#39;, &#39;48&#39; or &#39;54&#39;. Defaults to 12. | [optional] 
**AxEnabled** | Pointer to **bool** | Determines whether ax radio on 5Ghz band is on or off. Can be either true or false. If false, we highly recommend disabling band steering. Defaults to true. | [optional] 

## Methods

### NewNetworksNetworkIdApplianceRfProfilesFiveGhzSettings

`func NewNetworksNetworkIdApplianceRfProfilesFiveGhzSettings() *NetworksNetworkIdApplianceRfProfilesFiveGhzSettings`

NewNetworksNetworkIdApplianceRfProfilesFiveGhzSettings instantiates a new NetworksNetworkIdApplianceRfProfilesFiveGhzSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdApplianceRfProfilesFiveGhzSettingsWithDefaults

`func NewNetworksNetworkIdApplianceRfProfilesFiveGhzSettingsWithDefaults() *NetworksNetworkIdApplianceRfProfilesFiveGhzSettings`

NewNetworksNetworkIdApplianceRfProfilesFiveGhzSettingsWithDefaults instantiates a new NetworksNetworkIdApplianceRfProfilesFiveGhzSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinBitrate

`func (o *NetworksNetworkIdApplianceRfProfilesFiveGhzSettings) GetMinBitrate() int32`

GetMinBitrate returns the MinBitrate field if non-nil, zero value otherwise.

### GetMinBitrateOk

`func (o *NetworksNetworkIdApplianceRfProfilesFiveGhzSettings) GetMinBitrateOk() (*int32, bool)`

GetMinBitrateOk returns a tuple with the MinBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBitrate

`func (o *NetworksNetworkIdApplianceRfProfilesFiveGhzSettings) SetMinBitrate(v int32)`

SetMinBitrate sets MinBitrate field to given value.

### HasMinBitrate

`func (o *NetworksNetworkIdApplianceRfProfilesFiveGhzSettings) HasMinBitrate() bool`

HasMinBitrate returns a boolean if a field has been set.

### GetAxEnabled

`func (o *NetworksNetworkIdApplianceRfProfilesFiveGhzSettings) GetAxEnabled() bool`

GetAxEnabled returns the AxEnabled field if non-nil, zero value otherwise.

### GetAxEnabledOk

`func (o *NetworksNetworkIdApplianceRfProfilesFiveGhzSettings) GetAxEnabledOk() (*bool, bool)`

GetAxEnabledOk returns a tuple with the AxEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAxEnabled

`func (o *NetworksNetworkIdApplianceRfProfilesFiveGhzSettings) SetAxEnabled(v bool)`

SetAxEnabled sets AxEnabled field to given value.

### HasAxEnabled

`func (o *NetworksNetworkIdApplianceRfProfilesFiveGhzSettings) HasAxEnabled() bool`

HasAxEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


