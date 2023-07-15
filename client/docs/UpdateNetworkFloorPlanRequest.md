# UpdateNetworkFloorPlanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of your floor plan. | [optional] 
**Center** | Pointer to [**UpdateNetworkFloorPlanRequestCenter**](UpdateNetworkFloorPlanRequestCenter.md) |  | [optional] 
**BottomLeftCorner** | Pointer to [**GetNetworkFloorPlans200ResponseInnerBottomLeftCorner**](GetNetworkFloorPlans200ResponseInnerBottomLeftCorner.md) |  | [optional] 
**BottomRightCorner** | Pointer to [**GetNetworkFloorPlans200ResponseInnerBottomRightCorner**](GetNetworkFloorPlans200ResponseInnerBottomRightCorner.md) |  | [optional] 
**TopLeftCorner** | Pointer to [**GetNetworkFloorPlans200ResponseInnerTopLeftCorner**](GetNetworkFloorPlans200ResponseInnerTopLeftCorner.md) |  | [optional] 
**TopRightCorner** | Pointer to [**GetNetworkFloorPlans200ResponseInnerTopRightCorner**](GetNetworkFloorPlans200ResponseInnerTopRightCorner.md) |  | [optional] 
**ImageContents** | Pointer to **string** | The file contents (a base 64 encoded string) of your new image. Supported formats are PNG, GIF, and JPG. Note that all images are saved as PNG files, regardless of the format they are uploaded in. If you upload a new image, and you do NOT specify any new geolocation fields (&#39;center, &#39;topLeftCorner&#39;, etc), the floor plan will be recentered with no rotation in order to maintain the aspect ratio of your new image. | [optional] 

## Methods

### NewUpdateNetworkFloorPlanRequest

`func NewUpdateNetworkFloorPlanRequest() *UpdateNetworkFloorPlanRequest`

NewUpdateNetworkFloorPlanRequest instantiates a new UpdateNetworkFloorPlanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkFloorPlanRequestWithDefaults

`func NewUpdateNetworkFloorPlanRequestWithDefaults() *UpdateNetworkFloorPlanRequest`

NewUpdateNetworkFloorPlanRequestWithDefaults instantiates a new UpdateNetworkFloorPlanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateNetworkFloorPlanRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNetworkFloorPlanRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNetworkFloorPlanRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateNetworkFloorPlanRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCenter

`func (o *UpdateNetworkFloorPlanRequest) GetCenter() UpdateNetworkFloorPlanRequestCenter`

GetCenter returns the Center field if non-nil, zero value otherwise.

### GetCenterOk

`func (o *UpdateNetworkFloorPlanRequest) GetCenterOk() (*UpdateNetworkFloorPlanRequestCenter, bool)`

GetCenterOk returns a tuple with the Center field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCenter

`func (o *UpdateNetworkFloorPlanRequest) SetCenter(v UpdateNetworkFloorPlanRequestCenter)`

SetCenter sets Center field to given value.

### HasCenter

`func (o *UpdateNetworkFloorPlanRequest) HasCenter() bool`

HasCenter returns a boolean if a field has been set.

### GetBottomLeftCorner

`func (o *UpdateNetworkFloorPlanRequest) GetBottomLeftCorner() GetNetworkFloorPlans200ResponseInnerBottomLeftCorner`

GetBottomLeftCorner returns the BottomLeftCorner field if non-nil, zero value otherwise.

### GetBottomLeftCornerOk

`func (o *UpdateNetworkFloorPlanRequest) GetBottomLeftCornerOk() (*GetNetworkFloorPlans200ResponseInnerBottomLeftCorner, bool)`

GetBottomLeftCornerOk returns a tuple with the BottomLeftCorner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBottomLeftCorner

`func (o *UpdateNetworkFloorPlanRequest) SetBottomLeftCorner(v GetNetworkFloorPlans200ResponseInnerBottomLeftCorner)`

SetBottomLeftCorner sets BottomLeftCorner field to given value.

### HasBottomLeftCorner

`func (o *UpdateNetworkFloorPlanRequest) HasBottomLeftCorner() bool`

HasBottomLeftCorner returns a boolean if a field has been set.

### GetBottomRightCorner

`func (o *UpdateNetworkFloorPlanRequest) GetBottomRightCorner() GetNetworkFloorPlans200ResponseInnerBottomRightCorner`

GetBottomRightCorner returns the BottomRightCorner field if non-nil, zero value otherwise.

### GetBottomRightCornerOk

`func (o *UpdateNetworkFloorPlanRequest) GetBottomRightCornerOk() (*GetNetworkFloorPlans200ResponseInnerBottomRightCorner, bool)`

GetBottomRightCornerOk returns a tuple with the BottomRightCorner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBottomRightCorner

`func (o *UpdateNetworkFloorPlanRequest) SetBottomRightCorner(v GetNetworkFloorPlans200ResponseInnerBottomRightCorner)`

SetBottomRightCorner sets BottomRightCorner field to given value.

### HasBottomRightCorner

`func (o *UpdateNetworkFloorPlanRequest) HasBottomRightCorner() bool`

HasBottomRightCorner returns a boolean if a field has been set.

### GetTopLeftCorner

`func (o *UpdateNetworkFloorPlanRequest) GetTopLeftCorner() GetNetworkFloorPlans200ResponseInnerTopLeftCorner`

GetTopLeftCorner returns the TopLeftCorner field if non-nil, zero value otherwise.

### GetTopLeftCornerOk

`func (o *UpdateNetworkFloorPlanRequest) GetTopLeftCornerOk() (*GetNetworkFloorPlans200ResponseInnerTopLeftCorner, bool)`

GetTopLeftCornerOk returns a tuple with the TopLeftCorner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLeftCorner

`func (o *UpdateNetworkFloorPlanRequest) SetTopLeftCorner(v GetNetworkFloorPlans200ResponseInnerTopLeftCorner)`

SetTopLeftCorner sets TopLeftCorner field to given value.

### HasTopLeftCorner

`func (o *UpdateNetworkFloorPlanRequest) HasTopLeftCorner() bool`

HasTopLeftCorner returns a boolean if a field has been set.

### GetTopRightCorner

`func (o *UpdateNetworkFloorPlanRequest) GetTopRightCorner() GetNetworkFloorPlans200ResponseInnerTopRightCorner`

GetTopRightCorner returns the TopRightCorner field if non-nil, zero value otherwise.

### GetTopRightCornerOk

`func (o *UpdateNetworkFloorPlanRequest) GetTopRightCornerOk() (*GetNetworkFloorPlans200ResponseInnerTopRightCorner, bool)`

GetTopRightCornerOk returns a tuple with the TopRightCorner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopRightCorner

`func (o *UpdateNetworkFloorPlanRequest) SetTopRightCorner(v GetNetworkFloorPlans200ResponseInnerTopRightCorner)`

SetTopRightCorner sets TopRightCorner field to given value.

### HasTopRightCorner

`func (o *UpdateNetworkFloorPlanRequest) HasTopRightCorner() bool`

HasTopRightCorner returns a boolean if a field has been set.

### GetImageContents

`func (o *UpdateNetworkFloorPlanRequest) GetImageContents() string`

GetImageContents returns the ImageContents field if non-nil, zero value otherwise.

### GetImageContentsOk

`func (o *UpdateNetworkFloorPlanRequest) GetImageContentsOk() (*string, bool)`

GetImageContentsOk returns a tuple with the ImageContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageContents

`func (o *UpdateNetworkFloorPlanRequest) SetImageContents(v string)`

SetImageContents sets ImageContents field to given value.

### HasImageContents

`func (o *UpdateNetworkFloorPlanRequest) HasImageContents() bool`

HasImageContents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


