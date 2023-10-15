# ReassignNetworkVlanProfilesAssignmentsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VlanProfile** | Pointer to [**ReassignNetworkVlanProfilesAssignmentsRequestVlanProfile**](ReassignNetworkVlanProfilesAssignmentsRequestVlanProfile.md) |  | [optional] 
**Serials** | **[]string** | Array of Device Serials | 
**StackIds** | **[]string** | Array of Switch Stack IDs | 

## Methods

### NewReassignNetworkVlanProfilesAssignmentsRequest

`func NewReassignNetworkVlanProfilesAssignmentsRequest(serials []string, stackIds []string, ) *ReassignNetworkVlanProfilesAssignmentsRequest`

NewReassignNetworkVlanProfilesAssignmentsRequest instantiates a new ReassignNetworkVlanProfilesAssignmentsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReassignNetworkVlanProfilesAssignmentsRequestWithDefaults

`func NewReassignNetworkVlanProfilesAssignmentsRequestWithDefaults() *ReassignNetworkVlanProfilesAssignmentsRequest`

NewReassignNetworkVlanProfilesAssignmentsRequestWithDefaults instantiates a new ReassignNetworkVlanProfilesAssignmentsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVlanProfile

`func (o *ReassignNetworkVlanProfilesAssignmentsRequest) GetVlanProfile() ReassignNetworkVlanProfilesAssignmentsRequestVlanProfile`

GetVlanProfile returns the VlanProfile field if non-nil, zero value otherwise.

### GetVlanProfileOk

`func (o *ReassignNetworkVlanProfilesAssignmentsRequest) GetVlanProfileOk() (*ReassignNetworkVlanProfilesAssignmentsRequestVlanProfile, bool)`

GetVlanProfileOk returns a tuple with the VlanProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanProfile

`func (o *ReassignNetworkVlanProfilesAssignmentsRequest) SetVlanProfile(v ReassignNetworkVlanProfilesAssignmentsRequestVlanProfile)`

SetVlanProfile sets VlanProfile field to given value.

### HasVlanProfile

`func (o *ReassignNetworkVlanProfilesAssignmentsRequest) HasVlanProfile() bool`

HasVlanProfile returns a boolean if a field has been set.

### GetSerials

`func (o *ReassignNetworkVlanProfilesAssignmentsRequest) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *ReassignNetworkVlanProfilesAssignmentsRequest) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *ReassignNetworkVlanProfilesAssignmentsRequest) SetSerials(v []string)`

SetSerials sets Serials field to given value.


### GetStackIds

`func (o *ReassignNetworkVlanProfilesAssignmentsRequest) GetStackIds() []string`

GetStackIds returns the StackIds field if non-nil, zero value otherwise.

### GetStackIdsOk

`func (o *ReassignNetworkVlanProfilesAssignmentsRequest) GetStackIdsOk() (*[]string, bool)`

GetStackIdsOk returns a tuple with the StackIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackIds

`func (o *ReassignNetworkVlanProfilesAssignmentsRequest) SetStackIds(v []string)`

SetStackIds sets StackIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


