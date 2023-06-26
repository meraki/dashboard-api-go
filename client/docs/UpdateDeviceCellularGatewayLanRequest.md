# UpdateDeviceCellularGatewayLanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReservedIpRanges** | Pointer to [**[]UpdateDeviceCellularGatewayLanRequestReservedIpRangesInner**](UpdateDeviceCellularGatewayLanRequestReservedIpRangesInner.md) | list of all reserved IP ranges for a single MG | [optional] 
**FixedIpAssignments** | Pointer to [**[]UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner**](UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner.md) | list of all fixed IP assignments for a single MG | [optional] 

## Methods

### NewUpdateDeviceCellularGatewayLanRequest

`func NewUpdateDeviceCellularGatewayLanRequest() *UpdateDeviceCellularGatewayLanRequest`

NewUpdateDeviceCellularGatewayLanRequest instantiates a new UpdateDeviceCellularGatewayLanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeviceCellularGatewayLanRequestWithDefaults

`func NewUpdateDeviceCellularGatewayLanRequestWithDefaults() *UpdateDeviceCellularGatewayLanRequest`

NewUpdateDeviceCellularGatewayLanRequestWithDefaults instantiates a new UpdateDeviceCellularGatewayLanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReservedIpRanges

`func (o *UpdateDeviceCellularGatewayLanRequest) GetReservedIpRanges() []UpdateDeviceCellularGatewayLanRequestReservedIpRangesInner`

GetReservedIpRanges returns the ReservedIpRanges field if non-nil, zero value otherwise.

### GetReservedIpRangesOk

`func (o *UpdateDeviceCellularGatewayLanRequest) GetReservedIpRangesOk() (*[]UpdateDeviceCellularGatewayLanRequestReservedIpRangesInner, bool)`

GetReservedIpRangesOk returns a tuple with the ReservedIpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedIpRanges

`func (o *UpdateDeviceCellularGatewayLanRequest) SetReservedIpRanges(v []UpdateDeviceCellularGatewayLanRequestReservedIpRangesInner)`

SetReservedIpRanges sets ReservedIpRanges field to given value.

### HasReservedIpRanges

`func (o *UpdateDeviceCellularGatewayLanRequest) HasReservedIpRanges() bool`

HasReservedIpRanges returns a boolean if a field has been set.

### GetFixedIpAssignments

`func (o *UpdateDeviceCellularGatewayLanRequest) GetFixedIpAssignments() []UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner`

GetFixedIpAssignments returns the FixedIpAssignments field if non-nil, zero value otherwise.

### GetFixedIpAssignmentsOk

`func (o *UpdateDeviceCellularGatewayLanRequest) GetFixedIpAssignmentsOk() (*[]UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner, bool)`

GetFixedIpAssignmentsOk returns a tuple with the FixedIpAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedIpAssignments

`func (o *UpdateDeviceCellularGatewayLanRequest) SetFixedIpAssignments(v []UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner)`

SetFixedIpAssignments sets FixedIpAssignments field to given value.

### HasFixedIpAssignments

`func (o *UpdateDeviceCellularGatewayLanRequest) HasFixedIpAssignments() bool`

HasFixedIpAssignments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


