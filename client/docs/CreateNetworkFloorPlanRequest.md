# CreateNetworkFloorPlanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of your floor plan. | 
**Center** | Pointer to [**CreateNetworkFloorPlanRequestCenter**](CreateNetworkFloorPlanRequestCenter.md) |  | [optional] 
**BottomLeftCorner** | Pointer to [**CreateNetworkFloorPlanRequestBottomLeftCorner**](CreateNetworkFloorPlanRequestBottomLeftCorner.md) |  | [optional] 
**BottomRightCorner** | Pointer to [**CreateNetworkFloorPlanRequestBottomRightCorner**](CreateNetworkFloorPlanRequestBottomRightCorner.md) |  | [optional] 
**TopLeftCorner** | Pointer to [**CreateNetworkFloorPlanRequestTopLeftCorner**](CreateNetworkFloorPlanRequestTopLeftCorner.md) |  | [optional] 
**TopRightCorner** | Pointer to [**CreateNetworkFloorPlanRequestTopRightCorner**](CreateNetworkFloorPlanRequestTopRightCorner.md) |  | [optional] 
**ImageContents** | **string** | The file contents (a base 64 encoded string) of your image. Supported formats are PNG, GIF, and JPG. Note that all images are saved as PNG files, regardless of the format they are uploaded in. | 

## Methods

### NewCreateNetworkFloorPlanRequest

`func NewCreateNetworkFloorPlanRequest(name string, imageContents string, ) *CreateNetworkFloorPlanRequest`

NewCreateNetworkFloorPlanRequest instantiates a new CreateNetworkFloorPlanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkFloorPlanRequestWithDefaults

`func NewCreateNetworkFloorPlanRequestWithDefaults() *CreateNetworkFloorPlanRequest`

NewCreateNetworkFloorPlanRequestWithDefaults instantiates a new CreateNetworkFloorPlanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNetworkFloorPlanRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkFloorPlanRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkFloorPlanRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCenter

`func (o *CreateNetworkFloorPlanRequest) GetCenter() CreateNetworkFloorPlanRequestCenter`

GetCenter returns the Center field if non-nil, zero value otherwise.

### GetCenterOk

`func (o *CreateNetworkFloorPlanRequest) GetCenterOk() (*CreateNetworkFloorPlanRequestCenter, bool)`

GetCenterOk returns a tuple with the Center field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCenter

`func (o *CreateNetworkFloorPlanRequest) SetCenter(v CreateNetworkFloorPlanRequestCenter)`

SetCenter sets Center field to given value.

### HasCenter

`func (o *CreateNetworkFloorPlanRequest) HasCenter() bool`

HasCenter returns a boolean if a field has been set.

### GetBottomLeftCorner

`func (o *CreateNetworkFloorPlanRequest) GetBottomLeftCorner() CreateNetworkFloorPlanRequestBottomLeftCorner`

GetBottomLeftCorner returns the BottomLeftCorner field if non-nil, zero value otherwise.

### GetBottomLeftCornerOk

`func (o *CreateNetworkFloorPlanRequest) GetBottomLeftCornerOk() (*CreateNetworkFloorPlanRequestBottomLeftCorner, bool)`

GetBottomLeftCornerOk returns a tuple with the BottomLeftCorner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBottomLeftCorner

`func (o *CreateNetworkFloorPlanRequest) SetBottomLeftCorner(v CreateNetworkFloorPlanRequestBottomLeftCorner)`

SetBottomLeftCorner sets BottomLeftCorner field to given value.

### HasBottomLeftCorner

`func (o *CreateNetworkFloorPlanRequest) HasBottomLeftCorner() bool`

HasBottomLeftCorner returns a boolean if a field has been set.

### GetBottomRightCorner

`func (o *CreateNetworkFloorPlanRequest) GetBottomRightCorner() CreateNetworkFloorPlanRequestBottomRightCorner`

GetBottomRightCorner returns the BottomRightCorner field if non-nil, zero value otherwise.

### GetBottomRightCornerOk

`func (o *CreateNetworkFloorPlanRequest) GetBottomRightCornerOk() (*CreateNetworkFloorPlanRequestBottomRightCorner, bool)`

GetBottomRightCornerOk returns a tuple with the BottomRightCorner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBottomRightCorner

`func (o *CreateNetworkFloorPlanRequest) SetBottomRightCorner(v CreateNetworkFloorPlanRequestBottomRightCorner)`

SetBottomRightCorner sets BottomRightCorner field to given value.

### HasBottomRightCorner

`func (o *CreateNetworkFloorPlanRequest) HasBottomRightCorner() bool`

HasBottomRightCorner returns a boolean if a field has been set.

### GetTopLeftCorner

`func (o *CreateNetworkFloorPlanRequest) GetTopLeftCorner() CreateNetworkFloorPlanRequestTopLeftCorner`

GetTopLeftCorner returns the TopLeftCorner field if non-nil, zero value otherwise.

### GetTopLeftCornerOk

`func (o *CreateNetworkFloorPlanRequest) GetTopLeftCornerOk() (*CreateNetworkFloorPlanRequestTopLeftCorner, bool)`

GetTopLeftCornerOk returns a tuple with the TopLeftCorner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLeftCorner

`func (o *CreateNetworkFloorPlanRequest) SetTopLeftCorner(v CreateNetworkFloorPlanRequestTopLeftCorner)`

SetTopLeftCorner sets TopLeftCorner field to given value.

### HasTopLeftCorner

`func (o *CreateNetworkFloorPlanRequest) HasTopLeftCorner() bool`

HasTopLeftCorner returns a boolean if a field has been set.

### GetTopRightCorner

`func (o *CreateNetworkFloorPlanRequest) GetTopRightCorner() CreateNetworkFloorPlanRequestTopRightCorner`

GetTopRightCorner returns the TopRightCorner field if non-nil, zero value otherwise.

### GetTopRightCornerOk

`func (o *CreateNetworkFloorPlanRequest) GetTopRightCornerOk() (*CreateNetworkFloorPlanRequestTopRightCorner, bool)`

GetTopRightCornerOk returns a tuple with the TopRightCorner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopRightCorner

`func (o *CreateNetworkFloorPlanRequest) SetTopRightCorner(v CreateNetworkFloorPlanRequestTopRightCorner)`

SetTopRightCorner sets TopRightCorner field to given value.

### HasTopRightCorner

`func (o *CreateNetworkFloorPlanRequest) HasTopRightCorner() bool`

HasTopRightCorner returns a boolean if a field has been set.

### GetImageContents

`func (o *CreateNetworkFloorPlanRequest) GetImageContents() string`

GetImageContents returns the ImageContents field if non-nil, zero value otherwise.

### GetImageContentsOk

`func (o *CreateNetworkFloorPlanRequest) GetImageContentsOk() (*string, bool)`

GetImageContentsOk returns a tuple with the ImageContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageContents

`func (o *CreateNetworkFloorPlanRequest) SetImageContents(v string)`

SetImageContents sets ImageContents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


