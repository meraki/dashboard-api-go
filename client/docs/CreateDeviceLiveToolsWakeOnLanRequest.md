# CreateDeviceLiveToolsWakeOnLanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VlanId** | **int32** | The target&#39;s VLAN (1 to 4094) | 
**Mac** | **string** | The target&#39;s MAC address | 
**Callback** | Pointer to [**CreateDeviceLiveToolsArpTableRequestCallback**](CreateDeviceLiveToolsArpTableRequestCallback.md) |  | [optional] 

## Methods

### NewCreateDeviceLiveToolsWakeOnLanRequest

`func NewCreateDeviceLiveToolsWakeOnLanRequest(vlanId int32, mac string, ) *CreateDeviceLiveToolsWakeOnLanRequest`

NewCreateDeviceLiveToolsWakeOnLanRequest instantiates a new CreateDeviceLiveToolsWakeOnLanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDeviceLiveToolsWakeOnLanRequestWithDefaults

`func NewCreateDeviceLiveToolsWakeOnLanRequestWithDefaults() *CreateDeviceLiveToolsWakeOnLanRequest`

NewCreateDeviceLiveToolsWakeOnLanRequestWithDefaults instantiates a new CreateDeviceLiveToolsWakeOnLanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVlanId

`func (o *CreateDeviceLiveToolsWakeOnLanRequest) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *CreateDeviceLiveToolsWakeOnLanRequest) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *CreateDeviceLiveToolsWakeOnLanRequest) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.


### GetMac

`func (o *CreateDeviceLiveToolsWakeOnLanRequest) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *CreateDeviceLiveToolsWakeOnLanRequest) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *CreateDeviceLiveToolsWakeOnLanRequest) SetMac(v string)`

SetMac sets Mac field to given value.


### GetCallback

`func (o *CreateDeviceLiveToolsWakeOnLanRequest) GetCallback() CreateDeviceLiveToolsArpTableRequestCallback`

GetCallback returns the Callback field if non-nil, zero value otherwise.

### GetCallbackOk

`func (o *CreateDeviceLiveToolsWakeOnLanRequest) GetCallbackOk() (*CreateDeviceLiveToolsArpTableRequestCallback, bool)`

GetCallbackOk returns a tuple with the Callback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallback

`func (o *CreateDeviceLiveToolsWakeOnLanRequest) SetCallback(v CreateDeviceLiveToolsArpTableRequestCallback)`

SetCallback sets Callback field to given value.

### HasCallback

`func (o *CreateDeviceLiveToolsWakeOnLanRequest) HasCallback() bool`

HasCallback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


