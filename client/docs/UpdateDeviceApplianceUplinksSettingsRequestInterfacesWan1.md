# UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enable or disable the interface. | [optional] 
**VlanTagging** | Pointer to [**GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging**](GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging.md) |  | [optional] 
**Svis** | Pointer to [**GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Svis**](GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Svis.md) |  | [optional] 
**Pppoe** | Pointer to [**UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe**](UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe.md) |  | [optional] 

## Methods

### NewUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1

`func NewUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1() *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1`

NewUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1 instantiates a new UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1WithDefaults

`func NewUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1WithDefaults() *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1`

NewUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1WithDefaults instantiates a new UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetVlanTagging

`func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1) GetVlanTagging() GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging`

GetVlanTagging returns the VlanTagging field if non-nil, zero value otherwise.

### GetVlanTaggingOk

`func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1) GetVlanTaggingOk() (*GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging, bool)`

GetVlanTaggingOk returns a tuple with the VlanTagging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanTagging

`func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1) SetVlanTagging(v GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1VlanTagging)`

SetVlanTagging sets VlanTagging field to given value.

### HasVlanTagging

`func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1) HasVlanTagging() bool`

HasVlanTagging returns a boolean if a field has been set.

### GetSvis

`func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1) GetSvis() GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Svis`

GetSvis returns the Svis field if non-nil, zero value otherwise.

### GetSvisOk

`func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1) GetSvisOk() (*GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Svis, bool)`

GetSvisOk returns a tuple with the Svis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvis

`func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1) SetSvis(v GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Svis)`

SetSvis sets Svis field to given value.

### HasSvis

`func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1) HasSvis() bool`

HasSvis returns a boolean if a field has been set.

### GetPppoe

`func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1) GetPppoe() UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe`

GetPppoe returns the Pppoe field if non-nil, zero value otherwise.

### GetPppoeOk

`func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1) GetPppoeOk() (*UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe, bool)`

GetPppoeOk returns a tuple with the Pppoe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPppoe

`func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1) SetPppoe(v UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe)`

SetPppoe sets Pppoe field to given value.

### HasPppoe

`func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1) HasPppoe() bool`

HasPppoe returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


