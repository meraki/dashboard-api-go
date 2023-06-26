# GetNetworkAppliancePorts200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **int32** | Number of the port | [optional] 
**Enabled** | Pointer to **bool** | The status of the port | [optional] 
**Type** | Pointer to **string** | The type of the port: &#39;access&#39; or &#39;trunk&#39;. | [optional] 
**DropUntaggedTraffic** | Pointer to **bool** | Whether the trunk port can drop all untagged traffic. | [optional] 
**Vlan** | Pointer to **int32** | Native VLAN when the port is in Trunk mode. Access VLAN when the port is in Access mode. | [optional] 
**AllowedVlans** | Pointer to **string** | Comma-delimited list of the VLAN ID&#39;s allowed on the port, or &#39;all&#39; to permit all VLAN&#39;s on the port. | [optional] 
**AccessPolicy** | Pointer to **string** | The name of the policy. Only applicable to Access ports. | [optional] 

## Methods

### NewGetNetworkAppliancePorts200ResponseInner

`func NewGetNetworkAppliancePorts200ResponseInner() *GetNetworkAppliancePorts200ResponseInner`

NewGetNetworkAppliancePorts200ResponseInner instantiates a new GetNetworkAppliancePorts200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkAppliancePorts200ResponseInnerWithDefaults

`func NewGetNetworkAppliancePorts200ResponseInnerWithDefaults() *GetNetworkAppliancePorts200ResponseInner`

NewGetNetworkAppliancePorts200ResponseInnerWithDefaults instantiates a new GetNetworkAppliancePorts200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *GetNetworkAppliancePorts200ResponseInner) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GetNetworkAppliancePorts200ResponseInner) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GetNetworkAppliancePorts200ResponseInner) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *GetNetworkAppliancePorts200ResponseInner) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetEnabled

`func (o *GetNetworkAppliancePorts200ResponseInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetNetworkAppliancePorts200ResponseInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetNetworkAppliancePorts200ResponseInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetNetworkAppliancePorts200ResponseInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetType

`func (o *GetNetworkAppliancePorts200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetNetworkAppliancePorts200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetNetworkAppliancePorts200ResponseInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetNetworkAppliancePorts200ResponseInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDropUntaggedTraffic

`func (o *GetNetworkAppliancePorts200ResponseInner) GetDropUntaggedTraffic() bool`

GetDropUntaggedTraffic returns the DropUntaggedTraffic field if non-nil, zero value otherwise.

### GetDropUntaggedTrafficOk

`func (o *GetNetworkAppliancePorts200ResponseInner) GetDropUntaggedTrafficOk() (*bool, bool)`

GetDropUntaggedTrafficOk returns a tuple with the DropUntaggedTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropUntaggedTraffic

`func (o *GetNetworkAppliancePorts200ResponseInner) SetDropUntaggedTraffic(v bool)`

SetDropUntaggedTraffic sets DropUntaggedTraffic field to given value.

### HasDropUntaggedTraffic

`func (o *GetNetworkAppliancePorts200ResponseInner) HasDropUntaggedTraffic() bool`

HasDropUntaggedTraffic returns a boolean if a field has been set.

### GetVlan

`func (o *GetNetworkAppliancePorts200ResponseInner) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *GetNetworkAppliancePorts200ResponseInner) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *GetNetworkAppliancePorts200ResponseInner) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *GetNetworkAppliancePorts200ResponseInner) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetAllowedVlans

`func (o *GetNetworkAppliancePorts200ResponseInner) GetAllowedVlans() string`

GetAllowedVlans returns the AllowedVlans field if non-nil, zero value otherwise.

### GetAllowedVlansOk

`func (o *GetNetworkAppliancePorts200ResponseInner) GetAllowedVlansOk() (*string, bool)`

GetAllowedVlansOk returns a tuple with the AllowedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedVlans

`func (o *GetNetworkAppliancePorts200ResponseInner) SetAllowedVlans(v string)`

SetAllowedVlans sets AllowedVlans field to given value.

### HasAllowedVlans

`func (o *GetNetworkAppliancePorts200ResponseInner) HasAllowedVlans() bool`

HasAllowedVlans returns a boolean if a field has been set.

### GetAccessPolicy

`func (o *GetNetworkAppliancePorts200ResponseInner) GetAccessPolicy() string`

GetAccessPolicy returns the AccessPolicy field if non-nil, zero value otherwise.

### GetAccessPolicyOk

`func (o *GetNetworkAppliancePorts200ResponseInner) GetAccessPolicyOk() (*string, bool)`

GetAccessPolicyOk returns a tuple with the AccessPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPolicy

`func (o *GetNetworkAppliancePorts200ResponseInner) SetAccessPolicy(v string)`

SetAccessPolicy sets AccessPolicy field to given value.

### HasAccessPolicy

`func (o *GetNetworkAppliancePorts200ResponseInner) HasAccessPolicy() bool`

HasAccessPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


