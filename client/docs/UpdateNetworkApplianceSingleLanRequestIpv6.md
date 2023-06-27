# UpdateNetworkApplianceSingleLanRequestIpv6

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enable IPv6 on VLAN. | [optional] 
**PrefixAssignments** | Pointer to [**[]UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner**](UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner.md) | Prefix assignments on the VLAN | [optional] 

## Methods

### NewUpdateNetworkApplianceSingleLanRequestIpv6

`func NewUpdateNetworkApplianceSingleLanRequestIpv6() *UpdateNetworkApplianceSingleLanRequestIpv6`

NewUpdateNetworkApplianceSingleLanRequestIpv6 instantiates a new UpdateNetworkApplianceSingleLanRequestIpv6 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkApplianceSingleLanRequestIpv6WithDefaults

`func NewUpdateNetworkApplianceSingleLanRequestIpv6WithDefaults() *UpdateNetworkApplianceSingleLanRequestIpv6`

NewUpdateNetworkApplianceSingleLanRequestIpv6WithDefaults instantiates a new UpdateNetworkApplianceSingleLanRequestIpv6 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *UpdateNetworkApplianceSingleLanRequestIpv6) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateNetworkApplianceSingleLanRequestIpv6) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateNetworkApplianceSingleLanRequestIpv6) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateNetworkApplianceSingleLanRequestIpv6) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPrefixAssignments

`func (o *UpdateNetworkApplianceSingleLanRequestIpv6) GetPrefixAssignments() []UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner`

GetPrefixAssignments returns the PrefixAssignments field if non-nil, zero value otherwise.

### GetPrefixAssignmentsOk

`func (o *UpdateNetworkApplianceSingleLanRequestIpv6) GetPrefixAssignmentsOk() (*[]UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner, bool)`

GetPrefixAssignmentsOk returns a tuple with the PrefixAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixAssignments

`func (o *UpdateNetworkApplianceSingleLanRequestIpv6) SetPrefixAssignments(v []UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner)`

SetPrefixAssignments sets PrefixAssignments field to given value.

### HasPrefixAssignments

`func (o *UpdateNetworkApplianceSingleLanRequestIpv6) HasPrefixAssignments() bool`

HasPrefixAssignments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


