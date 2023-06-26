# UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enable or disable the interface. | [optional] 
**VlanTagging** | Pointer to [**GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging**](GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging.md) |  | [optional] 
**Svis** | Pointer to [**GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Svis**](GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Svis.md) |  | [optional] 
**Pppoe** | Pointer to [**UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe**](UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe.md) |  | [optional] 

## Methods

### NewUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2

`func NewUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2() *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2`

NewUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2 instantiates a new UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2WithDefaults

`func NewUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2WithDefaults() *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2`

NewUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2WithDefaults instantiates a new UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetVlanTagging

`func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) GetVlanTagging() GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging`

GetVlanTagging returns the VlanTagging field if non-nil, zero value otherwise.

### GetVlanTaggingOk

`func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) GetVlanTaggingOk() (*GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging, bool)`

GetVlanTaggingOk returns a tuple with the VlanTagging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanTagging

`func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) SetVlanTagging(v GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging)`

SetVlanTagging sets VlanTagging field to given value.

### HasVlanTagging

`func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) HasVlanTagging() bool`

HasVlanTagging returns a boolean if a field has been set.

### GetSvis

`func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) GetSvis() GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Svis`

GetSvis returns the Svis field if non-nil, zero value otherwise.

### GetSvisOk

`func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) GetSvisOk() (*GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Svis, bool)`

GetSvisOk returns a tuple with the Svis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvis

`func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) SetSvis(v GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Svis)`

SetSvis sets Svis field to given value.

### HasSvis

`func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) HasSvis() bool`

HasSvis returns a boolean if a field has been set.

### GetPppoe

`func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) GetPppoe() UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe`

GetPppoe returns the Pppoe field if non-nil, zero value otherwise.

### GetPppoeOk

`func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) GetPppoeOk() (*UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe, bool)`

GetPppoeOk returns a tuple with the Pppoe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPppoe

`func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) SetPppoe(v UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe)`

SetPppoe sets Pppoe field to given value.

### HasPppoe

`func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) HasPppoe() bool`

HasPppoe returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


