# GetNetworkSwitchRoutingMulticast200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultSettings** | Pointer to [**GetNetworkSwitchRoutingMulticast200ResponseDefaultSettings**](GetNetworkSwitchRoutingMulticast200ResponseDefaultSettings.md) |  | [optional] 
**Overrides** | Pointer to [**[]GetNetworkSwitchRoutingMulticast200ResponseOverridesInner**](GetNetworkSwitchRoutingMulticast200ResponseOverridesInner.md) | Array of paired switches/stacks/profiles and corresponding multicast settings.       An empty array will clear the multicast settings. | [optional] 

## Methods

### NewGetNetworkSwitchRoutingMulticast200Response

`func NewGetNetworkSwitchRoutingMulticast200Response() *GetNetworkSwitchRoutingMulticast200Response`

NewGetNetworkSwitchRoutingMulticast200Response instantiates a new GetNetworkSwitchRoutingMulticast200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSwitchRoutingMulticast200ResponseWithDefaults

`func NewGetNetworkSwitchRoutingMulticast200ResponseWithDefaults() *GetNetworkSwitchRoutingMulticast200Response`

NewGetNetworkSwitchRoutingMulticast200ResponseWithDefaults instantiates a new GetNetworkSwitchRoutingMulticast200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultSettings

`func (o *GetNetworkSwitchRoutingMulticast200Response) GetDefaultSettings() GetNetworkSwitchRoutingMulticast200ResponseDefaultSettings`

GetDefaultSettings returns the DefaultSettings field if non-nil, zero value otherwise.

### GetDefaultSettingsOk

`func (o *GetNetworkSwitchRoutingMulticast200Response) GetDefaultSettingsOk() (*GetNetworkSwitchRoutingMulticast200ResponseDefaultSettings, bool)`

GetDefaultSettingsOk returns a tuple with the DefaultSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSettings

`func (o *GetNetworkSwitchRoutingMulticast200Response) SetDefaultSettings(v GetNetworkSwitchRoutingMulticast200ResponseDefaultSettings)`

SetDefaultSettings sets DefaultSettings field to given value.

### HasDefaultSettings

`func (o *GetNetworkSwitchRoutingMulticast200Response) HasDefaultSettings() bool`

HasDefaultSettings returns a boolean if a field has been set.

### GetOverrides

`func (o *GetNetworkSwitchRoutingMulticast200Response) GetOverrides() []GetNetworkSwitchRoutingMulticast200ResponseOverridesInner`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *GetNetworkSwitchRoutingMulticast200Response) GetOverridesOk() (*[]GetNetworkSwitchRoutingMulticast200ResponseOverridesInner, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *GetNetworkSwitchRoutingMulticast200Response) SetOverrides(v []GetNetworkSwitchRoutingMulticast200ResponseOverridesInner)`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *GetNetworkSwitchRoutingMulticast200Response) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


