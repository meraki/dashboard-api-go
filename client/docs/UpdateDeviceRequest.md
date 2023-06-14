# UpdateDeviceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of a device | [optional] 
**Tags** | Pointer to **[]string** | The list of tags of a device | [optional] 
**Lat** | Pointer to **float32** | The latitude of a device | [optional] 
**Lng** | Pointer to **float32** | The longitude of a device | [optional] 
**Address** | Pointer to **string** | The address of a device | [optional] 
**Notes** | Pointer to **string** | The notes for the device. String. Limited to 255 characters. | [optional] 
**MoveMapMarker** | Pointer to **bool** | Whether or not to set the latitude and longitude of a device based on the new address. Only applies when lat and lng are not specified. | [optional] 
**SwitchProfileId** | Pointer to **string** | The ID of a switch profile to bind to the device (for available switch profiles, see the &#39;Switch Profiles&#39; endpoint). Use null to unbind the switch device from the current profile. For a device to be bindable to a switch profile, it must (1) be a switch, and (2) belong to a network that is bound to a configuration template. | [optional] 
**FloorPlanId** | Pointer to **string** | The floor plan to associate to this device. null disassociates the device from the floorplan. | [optional] 

## Methods

### NewUpdateDeviceRequest

`func NewUpdateDeviceRequest() *UpdateDeviceRequest`

NewUpdateDeviceRequest instantiates a new UpdateDeviceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeviceRequestWithDefaults

`func NewUpdateDeviceRequestWithDefaults() *UpdateDeviceRequest`

NewUpdateDeviceRequestWithDefaults instantiates a new UpdateDeviceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateDeviceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateDeviceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateDeviceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateDeviceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *UpdateDeviceRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateDeviceRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateDeviceRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateDeviceRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetLat

`func (o *UpdateDeviceRequest) GetLat() float32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *UpdateDeviceRequest) GetLatOk() (*float32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *UpdateDeviceRequest) SetLat(v float32)`

SetLat sets Lat field to given value.

### HasLat

`func (o *UpdateDeviceRequest) HasLat() bool`

HasLat returns a boolean if a field has been set.

### GetLng

`func (o *UpdateDeviceRequest) GetLng() float32`

GetLng returns the Lng field if non-nil, zero value otherwise.

### GetLngOk

`func (o *UpdateDeviceRequest) GetLngOk() (*float32, bool)`

GetLngOk returns a tuple with the Lng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLng

`func (o *UpdateDeviceRequest) SetLng(v float32)`

SetLng sets Lng field to given value.

### HasLng

`func (o *UpdateDeviceRequest) HasLng() bool`

HasLng returns a boolean if a field has been set.

### GetAddress

`func (o *UpdateDeviceRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UpdateDeviceRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UpdateDeviceRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *UpdateDeviceRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetNotes

`func (o *UpdateDeviceRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *UpdateDeviceRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *UpdateDeviceRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *UpdateDeviceRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetMoveMapMarker

`func (o *UpdateDeviceRequest) GetMoveMapMarker() bool`

GetMoveMapMarker returns the MoveMapMarker field if non-nil, zero value otherwise.

### GetMoveMapMarkerOk

`func (o *UpdateDeviceRequest) GetMoveMapMarkerOk() (*bool, bool)`

GetMoveMapMarkerOk returns a tuple with the MoveMapMarker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveMapMarker

`func (o *UpdateDeviceRequest) SetMoveMapMarker(v bool)`

SetMoveMapMarker sets MoveMapMarker field to given value.

### HasMoveMapMarker

`func (o *UpdateDeviceRequest) HasMoveMapMarker() bool`

HasMoveMapMarker returns a boolean if a field has been set.

### GetSwitchProfileId

`func (o *UpdateDeviceRequest) GetSwitchProfileId() string`

GetSwitchProfileId returns the SwitchProfileId field if non-nil, zero value otherwise.

### GetSwitchProfileIdOk

`func (o *UpdateDeviceRequest) GetSwitchProfileIdOk() (*string, bool)`

GetSwitchProfileIdOk returns a tuple with the SwitchProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchProfileId

`func (o *UpdateDeviceRequest) SetSwitchProfileId(v string)`

SetSwitchProfileId sets SwitchProfileId field to given value.

### HasSwitchProfileId

`func (o *UpdateDeviceRequest) HasSwitchProfileId() bool`

HasSwitchProfileId returns a boolean if a field has been set.

### GetFloorPlanId

`func (o *UpdateDeviceRequest) GetFloorPlanId() string`

GetFloorPlanId returns the FloorPlanId field if non-nil, zero value otherwise.

### GetFloorPlanIdOk

`func (o *UpdateDeviceRequest) GetFloorPlanIdOk() (*string, bool)`

GetFloorPlanIdOk returns a tuple with the FloorPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorPlanId

`func (o *UpdateDeviceRequest) SetFloorPlanId(v string)`

SetFloorPlanId sets FloorPlanId field to given value.

### HasFloorPlanId

`func (o *UpdateDeviceRequest) HasFloorPlanId() bool`

HasFloorPlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


