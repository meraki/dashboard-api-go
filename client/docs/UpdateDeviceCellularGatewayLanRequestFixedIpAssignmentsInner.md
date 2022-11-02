# UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A descriptive name of the assignment | [optional] 
**Ip** | **string** | The IP address you want to assign to a specific server or device | 
**Mac** | **string** | The MAC address of the server or device that hosts the internal resource that you wish to receive the specified IP address | 

## Methods

### NewUpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner

`func NewUpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner(ip string, mac string, ) *UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner`

NewUpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner instantiates a new UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInnerWithDefaults

`func NewUpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInnerWithDefaults() *UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner`

NewUpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInnerWithDefaults instantiates a new UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIp

`func (o *UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner) SetIp(v string)`

SetIp sets Ip field to given value.


### GetMac

`func (o *UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner) SetMac(v string)`

SetMac sets Mac field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


