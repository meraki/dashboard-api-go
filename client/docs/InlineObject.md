# InlineObject

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

### NewInlineObject

`func NewInlineObject() *InlineObject`

NewInlineObject instantiates a new InlineObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObjectWithDefaults

`func NewInlineObjectWithDefaults() *InlineObject`

NewInlineObjectWithDefaults instantiates a new InlineObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *InlineObject) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineObject) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineObject) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineObject) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetLat

`func (o *InlineObject) GetLat() float32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *InlineObject) GetLatOk() (*float32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *InlineObject) SetLat(v float32)`

SetLat sets Lat field to given value.

### HasLat

`func (o *InlineObject) HasLat() bool`

HasLat returns a boolean if a field has been set.

### GetLng

`func (o *InlineObject) GetLng() float32`

GetLng returns the Lng field if non-nil, zero value otherwise.

### GetLngOk

`func (o *InlineObject) GetLngOk() (*float32, bool)`

GetLngOk returns a tuple with the Lng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLng

`func (o *InlineObject) SetLng(v float32)`

SetLng sets Lng field to given value.

### HasLng

`func (o *InlineObject) HasLng() bool`

HasLng returns a boolean if a field has been set.

### GetAddress

`func (o *InlineObject) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *InlineObject) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *InlineObject) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *InlineObject) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetNotes

`func (o *InlineObject) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *InlineObject) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *InlineObject) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *InlineObject) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetMoveMapMarker

`func (o *InlineObject) GetMoveMapMarker() bool`

GetMoveMapMarker returns the MoveMapMarker field if non-nil, zero value otherwise.

### GetMoveMapMarkerOk

`func (o *InlineObject) GetMoveMapMarkerOk() (*bool, bool)`

GetMoveMapMarkerOk returns a tuple with the MoveMapMarker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveMapMarker

`func (o *InlineObject) SetMoveMapMarker(v bool)`

SetMoveMapMarker sets MoveMapMarker field to given value.

### HasMoveMapMarker

`func (o *InlineObject) HasMoveMapMarker() bool`

HasMoveMapMarker returns a boolean if a field has been set.

### GetSwitchProfileId

`func (o *InlineObject) GetSwitchProfileId() string`

GetSwitchProfileId returns the SwitchProfileId field if non-nil, zero value otherwise.

### GetSwitchProfileIdOk

`func (o *InlineObject) GetSwitchProfileIdOk() (*string, bool)`

GetSwitchProfileIdOk returns a tuple with the SwitchProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchProfileId

`func (o *InlineObject) SetSwitchProfileId(v string)`

SetSwitchProfileId sets SwitchProfileId field to given value.

### HasSwitchProfileId

`func (o *InlineObject) HasSwitchProfileId() bool`

HasSwitchProfileId returns a boolean if a field has been set.

### GetFloorPlanId

`func (o *InlineObject) GetFloorPlanId() string`

GetFloorPlanId returns the FloorPlanId field if non-nil, zero value otherwise.

### GetFloorPlanIdOk

`func (o *InlineObject) GetFloorPlanIdOk() (*string, bool)`

GetFloorPlanIdOk returns a tuple with the FloorPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorPlanId

`func (o *InlineObject) SetFloorPlanId(v string)`

SetFloorPlanId sets FloorPlanId field to given value.

### HasFloorPlanId

`func (o *InlineObject) HasFloorPlanId() bool`

HasFloorPlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


