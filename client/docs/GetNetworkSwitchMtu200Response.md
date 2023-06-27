# GetNetworkSwitchMtu200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultMtuSize** | Pointer to **int32** | MTU size for the entire network. Default value is 9578. | [optional] 
**Overrides** | Pointer to [**[]GetNetworkSwitchMtu200ResponseOverridesInner**](GetNetworkSwitchMtu200ResponseOverridesInner.md) | Override MTU size for individual switches or switch profiles.       An empty array will clear overrides. | [optional] 

## Methods

### NewGetNetworkSwitchMtu200Response

`func NewGetNetworkSwitchMtu200Response() *GetNetworkSwitchMtu200Response`

NewGetNetworkSwitchMtu200Response instantiates a new GetNetworkSwitchMtu200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSwitchMtu200ResponseWithDefaults

`func NewGetNetworkSwitchMtu200ResponseWithDefaults() *GetNetworkSwitchMtu200Response`

NewGetNetworkSwitchMtu200ResponseWithDefaults instantiates a new GetNetworkSwitchMtu200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultMtuSize

`func (o *GetNetworkSwitchMtu200Response) GetDefaultMtuSize() int32`

GetDefaultMtuSize returns the DefaultMtuSize field if non-nil, zero value otherwise.

### GetDefaultMtuSizeOk

`func (o *GetNetworkSwitchMtu200Response) GetDefaultMtuSizeOk() (*int32, bool)`

GetDefaultMtuSizeOk returns a tuple with the DefaultMtuSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMtuSize

`func (o *GetNetworkSwitchMtu200Response) SetDefaultMtuSize(v int32)`

SetDefaultMtuSize sets DefaultMtuSize field to given value.

### HasDefaultMtuSize

`func (o *GetNetworkSwitchMtu200Response) HasDefaultMtuSize() bool`

HasDefaultMtuSize returns a boolean if a field has been set.

### GetOverrides

`func (o *GetNetworkSwitchMtu200Response) GetOverrides() []GetNetworkSwitchMtu200ResponseOverridesInner`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *GetNetworkSwitchMtu200Response) GetOverridesOk() (*[]GetNetworkSwitchMtu200ResponseOverridesInner, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *GetNetworkSwitchMtu200Response) SetOverrides(v []GetNetworkSwitchMtu200ResponseOverridesInner)`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *GetNetworkSwitchMtu200Response) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


