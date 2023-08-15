/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 02 August, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetNetworkFloorPlans200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkFloorPlans200ResponseInner{}

// GetNetworkFloorPlans200ResponseInner struct for GetNetworkFloorPlans200ResponseInner
type GetNetworkFloorPlans200ResponseInner struct {
	// Floor plan ID
	FloorPlanId *string `json:"floorPlanId,omitempty"`
	// The url link for the floor plan image.
	ImageUrl *string `json:"imageUrl,omitempty"`
	// The time the image url link will expire.
	ImageUrlExpiresAt *string `json:"imageUrlExpiresAt,omitempty"`
	// The format type of the image.
	ImageExtension *string `json:"imageExtension,omitempty"`
	// The file contents (a base 64 encoded string) of your new image. Supported formats are PNG, GIF, and JPG. Note that all images are saved as PNG files, regardless of the format they are uploaded in. If you upload a new image, and you do NOT specify any new geolocation fields ('center, 'topLeftCorner', etc), the floor plan will be recentered with no rotation in order to maintain the aspect ratio of your new image.
	ImageMd5 *string `json:"imageMd5,omitempty"`
	// The name of your floor plan.
	Name *string `json:"name,omitempty"`
	// List of devices for the floorplan
	Devices []GetNetworkFloorPlans200ResponseInnerDevicesInner `json:"devices,omitempty"`
	// The width of your floor plan.
	Width *float32 `json:"width,omitempty"`
	// The height of your floor plan.
	Height *float32 `json:"height,omitempty"`
	Center *GetNetworkFloorPlans200ResponseInnerCenter `json:"center,omitempty"`
	BottomLeftCorner *GetNetworkFloorPlans200ResponseInnerBottomLeftCorner `json:"bottomLeftCorner,omitempty"`
	BottomRightCorner *GetNetworkFloorPlans200ResponseInnerBottomRightCorner `json:"bottomRightCorner,omitempty"`
	TopLeftCorner *GetNetworkFloorPlans200ResponseInnerTopLeftCorner `json:"topLeftCorner,omitempty"`
	TopRightCorner *GetNetworkFloorPlans200ResponseInnerTopRightCorner `json:"topRightCorner,omitempty"`
}

// NewGetNetworkFloorPlans200ResponseInner instantiates a new GetNetworkFloorPlans200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkFloorPlans200ResponseInner() *GetNetworkFloorPlans200ResponseInner {
	this := GetNetworkFloorPlans200ResponseInner{}
	return &this
}

// NewGetNetworkFloorPlans200ResponseInnerWithDefaults instantiates a new GetNetworkFloorPlans200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkFloorPlans200ResponseInnerWithDefaults() *GetNetworkFloorPlans200ResponseInner {
	this := GetNetworkFloorPlans200ResponseInner{}
	return &this
}

// GetFloorPlanId returns the FloorPlanId field value if set, zero value otherwise.
func (o *GetNetworkFloorPlans200ResponseInner) GetFloorPlanId() string {
	if o == nil || IsNil(o.FloorPlanId) {
		var ret string
		return ret
	}
	return *o.FloorPlanId
}

// GetFloorPlanIdOk returns a tuple with the FloorPlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFloorPlans200ResponseInner) GetFloorPlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.FloorPlanId) {
		return nil, false
	}
	return o.FloorPlanId, true
}

// HasFloorPlanId returns a boolean if a field has been set.
func (o *GetNetworkFloorPlans200ResponseInner) HasFloorPlanId() bool {
	if o != nil && !IsNil(o.FloorPlanId) {
		return true
	}

	return false
}

// SetFloorPlanId gets a reference to the given string and assigns it to the FloorPlanId field.
func (o *GetNetworkFloorPlans200ResponseInner) SetFloorPlanId(v string) {
	o.FloorPlanId = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *GetNetworkFloorPlans200ResponseInner) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFloorPlans200ResponseInner) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *GetNetworkFloorPlans200ResponseInner) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *GetNetworkFloorPlans200ResponseInner) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetImageUrlExpiresAt returns the ImageUrlExpiresAt field value if set, zero value otherwise.
func (o *GetNetworkFloorPlans200ResponseInner) GetImageUrlExpiresAt() string {
	if o == nil || IsNil(o.ImageUrlExpiresAt) {
		var ret string
		return ret
	}
	return *o.ImageUrlExpiresAt
}

// GetImageUrlExpiresAtOk returns a tuple with the ImageUrlExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFloorPlans200ResponseInner) GetImageUrlExpiresAtOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrlExpiresAt) {
		return nil, false
	}
	return o.ImageUrlExpiresAt, true
}

// HasImageUrlExpiresAt returns a boolean if a field has been set.
func (o *GetNetworkFloorPlans200ResponseInner) HasImageUrlExpiresAt() bool {
	if o != nil && !IsNil(o.ImageUrlExpiresAt) {
		return true
	}

	return false
}

// SetImageUrlExpiresAt gets a reference to the given string and assigns it to the ImageUrlExpiresAt field.
func (o *GetNetworkFloorPlans200ResponseInner) SetImageUrlExpiresAt(v string) {
	o.ImageUrlExpiresAt = &v
}

// GetImageExtension returns the ImageExtension field value if set, zero value otherwise.
func (o *GetNetworkFloorPlans200ResponseInner) GetImageExtension() string {
	if o == nil || IsNil(o.ImageExtension) {
		var ret string
		return ret
	}
	return *o.ImageExtension
}

// GetImageExtensionOk returns a tuple with the ImageExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFloorPlans200ResponseInner) GetImageExtensionOk() (*string, bool) {
	if o == nil || IsNil(o.ImageExtension) {
		return nil, false
	}
	return o.ImageExtension, true
}

// HasImageExtension returns a boolean if a field has been set.
func (o *GetNetworkFloorPlans200ResponseInner) HasImageExtension() bool {
	if o != nil && !IsNil(o.ImageExtension) {
		return true
	}

	return false
}

// SetImageExtension gets a reference to the given string and assigns it to the ImageExtension field.
func (o *GetNetworkFloorPlans200ResponseInner) SetImageExtension(v string) {
	o.ImageExtension = &v
}

// GetImageMd5 returns the ImageMd5 field value if set, zero value otherwise.
func (o *GetNetworkFloorPlans200ResponseInner) GetImageMd5() string {
	if o == nil || IsNil(o.ImageMd5) {
		var ret string
		return ret
	}
	return *o.ImageMd5
}

// GetImageMd5Ok returns a tuple with the ImageMd5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFloorPlans200ResponseInner) GetImageMd5Ok() (*string, bool) {
	if o == nil || IsNil(o.ImageMd5) {
		return nil, false
	}
	return o.ImageMd5, true
}

// HasImageMd5 returns a boolean if a field has been set.
func (o *GetNetworkFloorPlans200ResponseInner) HasImageMd5() bool {
	if o != nil && !IsNil(o.ImageMd5) {
		return true
	}

	return false
}

// SetImageMd5 gets a reference to the given string and assigns it to the ImageMd5 field.
func (o *GetNetworkFloorPlans200ResponseInner) SetImageMd5(v string) {
	o.ImageMd5 = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkFloorPlans200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFloorPlans200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkFloorPlans200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkFloorPlans200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetDevices returns the Devices field value if set, zero value otherwise.
func (o *GetNetworkFloorPlans200ResponseInner) GetDevices() []GetNetworkFloorPlans200ResponseInnerDevicesInner {
	if o == nil || IsNil(o.Devices) {
		var ret []GetNetworkFloorPlans200ResponseInnerDevicesInner
		return ret
	}
	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFloorPlans200ResponseInner) GetDevicesOk() ([]GetNetworkFloorPlans200ResponseInnerDevicesInner, bool) {
	if o == nil || IsNil(o.Devices) {
		return nil, false
	}
	return o.Devices, true
}

// HasDevices returns a boolean if a field has been set.
func (o *GetNetworkFloorPlans200ResponseInner) HasDevices() bool {
	if o != nil && !IsNil(o.Devices) {
		return true
	}

	return false
}

// SetDevices gets a reference to the given []GetNetworkFloorPlans200ResponseInnerDevicesInner and assigns it to the Devices field.
func (o *GetNetworkFloorPlans200ResponseInner) SetDevices(v []GetNetworkFloorPlans200ResponseInnerDevicesInner) {
	o.Devices = v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *GetNetworkFloorPlans200ResponseInner) GetWidth() float32 {
	if o == nil || IsNil(o.Width) {
		var ret float32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFloorPlans200ResponseInner) GetWidthOk() (*float32, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *GetNetworkFloorPlans200ResponseInner) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given float32 and assigns it to the Width field.
func (o *GetNetworkFloorPlans200ResponseInner) SetWidth(v float32) {
	o.Width = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *GetNetworkFloorPlans200ResponseInner) GetHeight() float32 {
	if o == nil || IsNil(o.Height) {
		var ret float32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFloorPlans200ResponseInner) GetHeightOk() (*float32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *GetNetworkFloorPlans200ResponseInner) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given float32 and assigns it to the Height field.
func (o *GetNetworkFloorPlans200ResponseInner) SetHeight(v float32) {
	o.Height = &v
}

// GetCenter returns the Center field value if set, zero value otherwise.
func (o *GetNetworkFloorPlans200ResponseInner) GetCenter() GetNetworkFloorPlans200ResponseInnerCenter {
	if o == nil || IsNil(o.Center) {
		var ret GetNetworkFloorPlans200ResponseInnerCenter
		return ret
	}
	return *o.Center
}

// GetCenterOk returns a tuple with the Center field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFloorPlans200ResponseInner) GetCenterOk() (*GetNetworkFloorPlans200ResponseInnerCenter, bool) {
	if o == nil || IsNil(o.Center) {
		return nil, false
	}
	return o.Center, true
}

// HasCenter returns a boolean if a field has been set.
func (o *GetNetworkFloorPlans200ResponseInner) HasCenter() bool {
	if o != nil && !IsNil(o.Center) {
		return true
	}

	return false
}

// SetCenter gets a reference to the given GetNetworkFloorPlans200ResponseInnerCenter and assigns it to the Center field.
func (o *GetNetworkFloorPlans200ResponseInner) SetCenter(v GetNetworkFloorPlans200ResponseInnerCenter) {
	o.Center = &v
}

// GetBottomLeftCorner returns the BottomLeftCorner field value if set, zero value otherwise.
func (o *GetNetworkFloorPlans200ResponseInner) GetBottomLeftCorner() GetNetworkFloorPlans200ResponseInnerBottomLeftCorner {
	if o == nil || IsNil(o.BottomLeftCorner) {
		var ret GetNetworkFloorPlans200ResponseInnerBottomLeftCorner
		return ret
	}
	return *o.BottomLeftCorner
}

// GetBottomLeftCornerOk returns a tuple with the BottomLeftCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFloorPlans200ResponseInner) GetBottomLeftCornerOk() (*GetNetworkFloorPlans200ResponseInnerBottomLeftCorner, bool) {
	if o == nil || IsNil(o.BottomLeftCorner) {
		return nil, false
	}
	return o.BottomLeftCorner, true
}

// HasBottomLeftCorner returns a boolean if a field has been set.
func (o *GetNetworkFloorPlans200ResponseInner) HasBottomLeftCorner() bool {
	if o != nil && !IsNil(o.BottomLeftCorner) {
		return true
	}

	return false
}

// SetBottomLeftCorner gets a reference to the given GetNetworkFloorPlans200ResponseInnerBottomLeftCorner and assigns it to the BottomLeftCorner field.
func (o *GetNetworkFloorPlans200ResponseInner) SetBottomLeftCorner(v GetNetworkFloorPlans200ResponseInnerBottomLeftCorner) {
	o.BottomLeftCorner = &v
}

// GetBottomRightCorner returns the BottomRightCorner field value if set, zero value otherwise.
func (o *GetNetworkFloorPlans200ResponseInner) GetBottomRightCorner() GetNetworkFloorPlans200ResponseInnerBottomRightCorner {
	if o == nil || IsNil(o.BottomRightCorner) {
		var ret GetNetworkFloorPlans200ResponseInnerBottomRightCorner
		return ret
	}
	return *o.BottomRightCorner
}

// GetBottomRightCornerOk returns a tuple with the BottomRightCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFloorPlans200ResponseInner) GetBottomRightCornerOk() (*GetNetworkFloorPlans200ResponseInnerBottomRightCorner, bool) {
	if o == nil || IsNil(o.BottomRightCorner) {
		return nil, false
	}
	return o.BottomRightCorner, true
}

// HasBottomRightCorner returns a boolean if a field has been set.
func (o *GetNetworkFloorPlans200ResponseInner) HasBottomRightCorner() bool {
	if o != nil && !IsNil(o.BottomRightCorner) {
		return true
	}

	return false
}

// SetBottomRightCorner gets a reference to the given GetNetworkFloorPlans200ResponseInnerBottomRightCorner and assigns it to the BottomRightCorner field.
func (o *GetNetworkFloorPlans200ResponseInner) SetBottomRightCorner(v GetNetworkFloorPlans200ResponseInnerBottomRightCorner) {
	o.BottomRightCorner = &v
}

// GetTopLeftCorner returns the TopLeftCorner field value if set, zero value otherwise.
func (o *GetNetworkFloorPlans200ResponseInner) GetTopLeftCorner() GetNetworkFloorPlans200ResponseInnerTopLeftCorner {
	if o == nil || IsNil(o.TopLeftCorner) {
		var ret GetNetworkFloorPlans200ResponseInnerTopLeftCorner
		return ret
	}
	return *o.TopLeftCorner
}

// GetTopLeftCornerOk returns a tuple with the TopLeftCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFloorPlans200ResponseInner) GetTopLeftCornerOk() (*GetNetworkFloorPlans200ResponseInnerTopLeftCorner, bool) {
	if o == nil || IsNil(o.TopLeftCorner) {
		return nil, false
	}
	return o.TopLeftCorner, true
}

// HasTopLeftCorner returns a boolean if a field has been set.
func (o *GetNetworkFloorPlans200ResponseInner) HasTopLeftCorner() bool {
	if o != nil && !IsNil(o.TopLeftCorner) {
		return true
	}

	return false
}

// SetTopLeftCorner gets a reference to the given GetNetworkFloorPlans200ResponseInnerTopLeftCorner and assigns it to the TopLeftCorner field.
func (o *GetNetworkFloorPlans200ResponseInner) SetTopLeftCorner(v GetNetworkFloorPlans200ResponseInnerTopLeftCorner) {
	o.TopLeftCorner = &v
}

// GetTopRightCorner returns the TopRightCorner field value if set, zero value otherwise.
func (o *GetNetworkFloorPlans200ResponseInner) GetTopRightCorner() GetNetworkFloorPlans200ResponseInnerTopRightCorner {
	if o == nil || IsNil(o.TopRightCorner) {
		var ret GetNetworkFloorPlans200ResponseInnerTopRightCorner
		return ret
	}
	return *o.TopRightCorner
}

// GetTopRightCornerOk returns a tuple with the TopRightCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFloorPlans200ResponseInner) GetTopRightCornerOk() (*GetNetworkFloorPlans200ResponseInnerTopRightCorner, bool) {
	if o == nil || IsNil(o.TopRightCorner) {
		return nil, false
	}
	return o.TopRightCorner, true
}

// HasTopRightCorner returns a boolean if a field has been set.
func (o *GetNetworkFloorPlans200ResponseInner) HasTopRightCorner() bool {
	if o != nil && !IsNil(o.TopRightCorner) {
		return true
	}

	return false
}

// SetTopRightCorner gets a reference to the given GetNetworkFloorPlans200ResponseInnerTopRightCorner and assigns it to the TopRightCorner field.
func (o *GetNetworkFloorPlans200ResponseInner) SetTopRightCorner(v GetNetworkFloorPlans200ResponseInnerTopRightCorner) {
	o.TopRightCorner = &v
}

func (o GetNetworkFloorPlans200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkFloorPlans200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FloorPlanId) {
		toSerialize["floorPlanId"] = o.FloorPlanId
	}
	if !IsNil(o.ImageUrl) {
		toSerialize["imageUrl"] = o.ImageUrl
	}
	if !IsNil(o.ImageUrlExpiresAt) {
		toSerialize["imageUrlExpiresAt"] = o.ImageUrlExpiresAt
	}
	if !IsNil(o.ImageExtension) {
		toSerialize["imageExtension"] = o.ImageExtension
	}
	if !IsNil(o.ImageMd5) {
		toSerialize["imageMd5"] = o.ImageMd5
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Devices) {
		toSerialize["devices"] = o.Devices
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if !IsNil(o.Center) {
		toSerialize["center"] = o.Center
	}
	if !IsNil(o.BottomLeftCorner) {
		toSerialize["bottomLeftCorner"] = o.BottomLeftCorner
	}
	if !IsNil(o.BottomRightCorner) {
		toSerialize["bottomRightCorner"] = o.BottomRightCorner
	}
	if !IsNil(o.TopLeftCorner) {
		toSerialize["topLeftCorner"] = o.TopLeftCorner
	}
	if !IsNil(o.TopRightCorner) {
		toSerialize["topRightCorner"] = o.TopRightCorner
	}
	return toSerialize, nil
}

type NullableGetNetworkFloorPlans200ResponseInner struct {
	value *GetNetworkFloorPlans200ResponseInner
	isSet bool
}

func (v NullableGetNetworkFloorPlans200ResponseInner) Get() *GetNetworkFloorPlans200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkFloorPlans200ResponseInner) Set(val *GetNetworkFloorPlans200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkFloorPlans200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkFloorPlans200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkFloorPlans200ResponseInner(val *GetNetworkFloorPlans200ResponseInner) *NullableGetNetworkFloorPlans200ResponseInner {
	return &NullableGetNetworkFloorPlans200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkFloorPlans200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkFloorPlans200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


