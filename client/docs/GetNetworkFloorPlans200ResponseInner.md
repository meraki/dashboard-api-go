# GetNetworkFloorPlans200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FloorPlanId** | Pointer to **string** | Floor plan ID | [optional] 
**ImageUrl** | Pointer to **string** | The url link for the floor plan image. | [optional] 
**ImageUrlExpiresAt** | Pointer to **string** | The time the image url link will expire. | [optional] 
**ImageExtension** | Pointer to **string** | The format type of the image. | [optional] 
**ImageMd5** | Pointer to **string** | The file contents (a base 64 encoded string) of your new image. Supported formats are PNG, GIF, and JPG. Note that all images are saved as PNG files, regardless of the format they are uploaded in. If you upload a new image, and you do NOT specify any new geolocation fields (&#39;center, &#39;topLeftCorner&#39;, etc), the floor plan will be recentered with no rotation in order to maintain the aspect ratio of your new image. | [optional] 
**Name** | Pointer to **string** | The name of your floor plan. | [optional] 
**Devices** | Pointer to [**[]GetNetworkFloorPlans200ResponseInnerDevicesInner**](GetNetworkFloorPlans200ResponseInnerDevicesInner.md) | List of devices for the floorplan | [optional] 
**Width** | Pointer to **float32** | The width of your floor plan. | [optional] 
**Height** | Pointer to **float32** | The height of your floor plan. | [optional] 
**Center** | Pointer to [**GetNetworkFloorPlans200ResponseInnerCenter**](GetNetworkFloorPlans200ResponseInnerCenter.md) |  | [optional] 
**BottomLeftCorner** | Pointer to [**GetNetworkFloorPlans200ResponseInnerBottomLeftCorner**](GetNetworkFloorPlans200ResponseInnerBottomLeftCorner.md) |  | [optional] 
**BottomRightCorner** | Pointer to [**GetNetworkFloorPlans200ResponseInnerBottomRightCorner**](GetNetworkFloorPlans200ResponseInnerBottomRightCorner.md) |  | [optional] 
**TopLeftCorner** | Pointer to [**GetNetworkFloorPlans200ResponseInnerTopLeftCorner**](GetNetworkFloorPlans200ResponseInnerTopLeftCorner.md) |  | [optional] 
**TopRightCorner** | Pointer to [**GetNetworkFloorPlans200ResponseInnerTopRightCorner**](GetNetworkFloorPlans200ResponseInnerTopRightCorner.md) |  | [optional] 

## Methods

### NewGetNetworkFloorPlans200ResponseInner

`func NewGetNetworkFloorPlans200ResponseInner() *GetNetworkFloorPlans200ResponseInner`

NewGetNetworkFloorPlans200ResponseInner instantiates a new GetNetworkFloorPlans200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkFloorPlans200ResponseInnerWithDefaults

`func NewGetNetworkFloorPlans200ResponseInnerWithDefaults() *GetNetworkFloorPlans200ResponseInner`

NewGetNetworkFloorPlans200ResponseInnerWithDefaults instantiates a new GetNetworkFloorPlans200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFloorPlanId

`func (o *GetNetworkFloorPlans200ResponseInner) GetFloorPlanId() string`

GetFloorPlanId returns the FloorPlanId field if non-nil, zero value otherwise.

### GetFloorPlanIdOk

`func (o *GetNetworkFloorPlans200ResponseInner) GetFloorPlanIdOk() (*string, bool)`

GetFloorPlanIdOk returns a tuple with the FloorPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorPlanId

`func (o *GetNetworkFloorPlans200ResponseInner) SetFloorPlanId(v string)`

SetFloorPlanId sets FloorPlanId field to given value.

### HasFloorPlanId

`func (o *GetNetworkFloorPlans200ResponseInner) HasFloorPlanId() bool`

HasFloorPlanId returns a boolean if a field has been set.

### GetImageUrl

`func (o *GetNetworkFloorPlans200ResponseInner) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *GetNetworkFloorPlans200ResponseInner) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *GetNetworkFloorPlans200ResponseInner) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *GetNetworkFloorPlans200ResponseInner) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetImageUrlExpiresAt

`func (o *GetNetworkFloorPlans200ResponseInner) GetImageUrlExpiresAt() string`

GetImageUrlExpiresAt returns the ImageUrlExpiresAt field if non-nil, zero value otherwise.

### GetImageUrlExpiresAtOk

`func (o *GetNetworkFloorPlans200ResponseInner) GetImageUrlExpiresAtOk() (*string, bool)`

GetImageUrlExpiresAtOk returns a tuple with the ImageUrlExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrlExpiresAt

`func (o *GetNetworkFloorPlans200ResponseInner) SetImageUrlExpiresAt(v string)`

SetImageUrlExpiresAt sets ImageUrlExpiresAt field to given value.

### HasImageUrlExpiresAt

`func (o *GetNetworkFloorPlans200ResponseInner) HasImageUrlExpiresAt() bool`

HasImageUrlExpiresAt returns a boolean if a field has been set.

### GetImageExtension

`func (o *GetNetworkFloorPlans200ResponseInner) GetImageExtension() string`

GetImageExtension returns the ImageExtension field if non-nil, zero value otherwise.

### GetImageExtensionOk

`func (o *GetNetworkFloorPlans200ResponseInner) GetImageExtensionOk() (*string, bool)`

GetImageExtensionOk returns a tuple with the ImageExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageExtension

`func (o *GetNetworkFloorPlans200ResponseInner) SetImageExtension(v string)`

SetImageExtension sets ImageExtension field to given value.

### HasImageExtension

`func (o *GetNetworkFloorPlans200ResponseInner) HasImageExtension() bool`

HasImageExtension returns a boolean if a field has been set.

### GetImageMd5

`func (o *GetNetworkFloorPlans200ResponseInner) GetImageMd5() string`

GetImageMd5 returns the ImageMd5 field if non-nil, zero value otherwise.

### GetImageMd5Ok

`func (o *GetNetworkFloorPlans200ResponseInner) GetImageMd5Ok() (*string, bool)`

GetImageMd5Ok returns a tuple with the ImageMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageMd5

`func (o *GetNetworkFloorPlans200ResponseInner) SetImageMd5(v string)`

SetImageMd5 sets ImageMd5 field to given value.

### HasImageMd5

`func (o *GetNetworkFloorPlans200ResponseInner) HasImageMd5() bool`

HasImageMd5 returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkFloorPlans200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkFloorPlans200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkFloorPlans200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkFloorPlans200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDevices

`func (o *GetNetworkFloorPlans200ResponseInner) GetDevices() []GetNetworkFloorPlans200ResponseInnerDevicesInner`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *GetNetworkFloorPlans200ResponseInner) GetDevicesOk() (*[]GetNetworkFloorPlans200ResponseInnerDevicesInner, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *GetNetworkFloorPlans200ResponseInner) SetDevices(v []GetNetworkFloorPlans200ResponseInnerDevicesInner)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *GetNetworkFloorPlans200ResponseInner) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetWidth

`func (o *GetNetworkFloorPlans200ResponseInner) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *GetNetworkFloorPlans200ResponseInner) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *GetNetworkFloorPlans200ResponseInner) SetWidth(v float32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *GetNetworkFloorPlans200ResponseInner) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *GetNetworkFloorPlans200ResponseInner) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *GetNetworkFloorPlans200ResponseInner) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *GetNetworkFloorPlans200ResponseInner) SetHeight(v float32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *GetNetworkFloorPlans200ResponseInner) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetCenter

`func (o *GetNetworkFloorPlans200ResponseInner) GetCenter() GetNetworkFloorPlans200ResponseInnerCenter`

GetCenter returns the Center field if non-nil, zero value otherwise.

### GetCenterOk

`func (o *GetNetworkFloorPlans200ResponseInner) GetCenterOk() (*GetNetworkFloorPlans200ResponseInnerCenter, bool)`

GetCenterOk returns a tuple with the Center field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCenter

`func (o *GetNetworkFloorPlans200ResponseInner) SetCenter(v GetNetworkFloorPlans200ResponseInnerCenter)`

SetCenter sets Center field to given value.

### HasCenter

`func (o *GetNetworkFloorPlans200ResponseInner) HasCenter() bool`

HasCenter returns a boolean if a field has been set.

### GetBottomLeftCorner

`func (o *GetNetworkFloorPlans200ResponseInner) GetBottomLeftCorner() GetNetworkFloorPlans200ResponseInnerBottomLeftCorner`

GetBottomLeftCorner returns the BottomLeftCorner field if non-nil, zero value otherwise.

### GetBottomLeftCornerOk

`func (o *GetNetworkFloorPlans200ResponseInner) GetBottomLeftCornerOk() (*GetNetworkFloorPlans200ResponseInnerBottomLeftCorner, bool)`

GetBottomLeftCornerOk returns a tuple with the BottomLeftCorner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBottomLeftCorner

`func (o *GetNetworkFloorPlans200ResponseInner) SetBottomLeftCorner(v GetNetworkFloorPlans200ResponseInnerBottomLeftCorner)`

SetBottomLeftCorner sets BottomLeftCorner field to given value.

### HasBottomLeftCorner

`func (o *GetNetworkFloorPlans200ResponseInner) HasBottomLeftCorner() bool`

HasBottomLeftCorner returns a boolean if a field has been set.

### GetBottomRightCorner

`func (o *GetNetworkFloorPlans200ResponseInner) GetBottomRightCorner() GetNetworkFloorPlans200ResponseInnerBottomRightCorner`

GetBottomRightCorner returns the BottomRightCorner field if non-nil, zero value otherwise.

### GetBottomRightCornerOk

`func (o *GetNetworkFloorPlans200ResponseInner) GetBottomRightCornerOk() (*GetNetworkFloorPlans200ResponseInnerBottomRightCorner, bool)`

GetBottomRightCornerOk returns a tuple with the BottomRightCorner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBottomRightCorner

`func (o *GetNetworkFloorPlans200ResponseInner) SetBottomRightCorner(v GetNetworkFloorPlans200ResponseInnerBottomRightCorner)`

SetBottomRightCorner sets BottomRightCorner field to given value.

### HasBottomRightCorner

`func (o *GetNetworkFloorPlans200ResponseInner) HasBottomRightCorner() bool`

HasBottomRightCorner returns a boolean if a field has been set.

### GetTopLeftCorner

`func (o *GetNetworkFloorPlans200ResponseInner) GetTopLeftCorner() GetNetworkFloorPlans200ResponseInnerTopLeftCorner`

GetTopLeftCorner returns the TopLeftCorner field if non-nil, zero value otherwise.

### GetTopLeftCornerOk

`func (o *GetNetworkFloorPlans200ResponseInner) GetTopLeftCornerOk() (*GetNetworkFloorPlans200ResponseInnerTopLeftCorner, bool)`

GetTopLeftCornerOk returns a tuple with the TopLeftCorner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLeftCorner

`func (o *GetNetworkFloorPlans200ResponseInner) SetTopLeftCorner(v GetNetworkFloorPlans200ResponseInnerTopLeftCorner)`

SetTopLeftCorner sets TopLeftCorner field to given value.

### HasTopLeftCorner

`func (o *GetNetworkFloorPlans200ResponseInner) HasTopLeftCorner() bool`

HasTopLeftCorner returns a boolean if a field has been set.

### GetTopRightCorner

`func (o *GetNetworkFloorPlans200ResponseInner) GetTopRightCorner() GetNetworkFloorPlans200ResponseInnerTopRightCorner`

GetTopRightCorner returns the TopRightCorner field if non-nil, zero value otherwise.

### GetTopRightCornerOk

`func (o *GetNetworkFloorPlans200ResponseInner) GetTopRightCornerOk() (*GetNetworkFloorPlans200ResponseInnerTopRightCorner, bool)`

GetTopRightCornerOk returns a tuple with the TopRightCorner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopRightCorner

`func (o *GetNetworkFloorPlans200ResponseInner) SetTopRightCorner(v GetNetworkFloorPlans200ResponseInnerTopRightCorner)`

SetTopRightCorner sets TopRightCorner field to given value.

### HasTopRightCorner

`func (o *GetNetworkFloorPlans200ResponseInner) HasTopRightCorner() bool`

HasTopRightCorner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


