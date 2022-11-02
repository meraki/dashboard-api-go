# UpdateDeviceCellularSimsRequestSimsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slot** | Pointer to **string** | SIM slot being configured. Must be &#39;sim1&#39; on single-sim devices. | [optional] 
**IsPrimary** | Pointer to **bool** | If true, this SIM is used for boot. Must be true on single-sim devices. | [optional] [default to false]
**Apns** | Pointer to [**[]UpdateDeviceCellularSimsRequestSimsInnerApnsInner**](UpdateDeviceCellularSimsRequestSimsInnerApnsInner.md) | APN configurations. If empty, the default APN will be used. | [optional] 

## Methods

### NewUpdateDeviceCellularSimsRequestSimsInner

`func NewUpdateDeviceCellularSimsRequestSimsInner() *UpdateDeviceCellularSimsRequestSimsInner`

NewUpdateDeviceCellularSimsRequestSimsInner instantiates a new UpdateDeviceCellularSimsRequestSimsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeviceCellularSimsRequestSimsInnerWithDefaults

`func NewUpdateDeviceCellularSimsRequestSimsInnerWithDefaults() *UpdateDeviceCellularSimsRequestSimsInner`

NewUpdateDeviceCellularSimsRequestSimsInnerWithDefaults instantiates a new UpdateDeviceCellularSimsRequestSimsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlot

`func (o *UpdateDeviceCellularSimsRequestSimsInner) GetSlot() string`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *UpdateDeviceCellularSimsRequestSimsInner) GetSlotOk() (*string, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *UpdateDeviceCellularSimsRequestSimsInner) SetSlot(v string)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *UpdateDeviceCellularSimsRequestSimsInner) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### GetIsPrimary

`func (o *UpdateDeviceCellularSimsRequestSimsInner) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *UpdateDeviceCellularSimsRequestSimsInner) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *UpdateDeviceCellularSimsRequestSimsInner) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *UpdateDeviceCellularSimsRequestSimsInner) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### GetApns

`func (o *UpdateDeviceCellularSimsRequestSimsInner) GetApns() []UpdateDeviceCellularSimsRequestSimsInnerApnsInner`

GetApns returns the Apns field if non-nil, zero value otherwise.

### GetApnsOk

`func (o *UpdateDeviceCellularSimsRequestSimsInner) GetApnsOk() (*[]UpdateDeviceCellularSimsRequestSimsInnerApnsInner, bool)`

GetApnsOk returns a tuple with the Apns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApns

`func (o *UpdateDeviceCellularSimsRequestSimsInner) SetApns(v []UpdateDeviceCellularSimsRequestSimsInnerApnsInner)`

SetApns sets Apns field to given value.

### HasApns

`func (o *UpdateDeviceCellularSimsRequestSimsInner) HasApns() bool`

HasApns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


