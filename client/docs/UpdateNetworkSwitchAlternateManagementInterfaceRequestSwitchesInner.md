# UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | **string** | Switch serial number | 
**AlternateManagementIp** | **string** | Switch alternative management IP. To remove a prior IP setting, provide an empty string | 
**SubnetMask** | Pointer to **string** | Switch subnet mask must be in IP format. Only and must be specified for Polaris switches | [optional] 
**Gateway** | Pointer to **string** | Switch gateway must be in IP format. Only and must be specified for Polaris switches | [optional] 

## Methods

### NewUpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner

`func NewUpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner(serial string, alternateManagementIp string, ) *UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner`

NewUpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner instantiates a new UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInnerWithDefaults

`func NewUpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInnerWithDefaults() *UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner`

NewUpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInnerWithDefaults instantiates a new UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner) SetSerial(v string)`

SetSerial sets Serial field to given value.


### GetAlternateManagementIp

`func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner) GetAlternateManagementIp() string`

GetAlternateManagementIp returns the AlternateManagementIp field if non-nil, zero value otherwise.

### GetAlternateManagementIpOk

`func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner) GetAlternateManagementIpOk() (*string, bool)`

GetAlternateManagementIpOk returns a tuple with the AlternateManagementIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateManagementIp

`func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner) SetAlternateManagementIp(v string)`

SetAlternateManagementIp sets AlternateManagementIp field to given value.


### GetSubnetMask

`func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner) GetSubnetMask() string`

GetSubnetMask returns the SubnetMask field if non-nil, zero value otherwise.

### GetSubnetMaskOk

`func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner) GetSubnetMaskOk() (*string, bool)`

GetSubnetMaskOk returns a tuple with the SubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetMask

`func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner) SetSubnetMask(v string)`

SetSubnetMask sets SubnetMask field to given value.

### HasSubnetMask

`func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner) HasSubnetMask() bool`

HasSubnetMask returns a boolean if a field has been set.

### GetGateway

`func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner) HasGateway() bool`

HasGateway returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


