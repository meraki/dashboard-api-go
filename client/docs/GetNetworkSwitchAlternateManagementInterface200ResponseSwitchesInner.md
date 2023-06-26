# GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Switch serial number | [optional] 
**AlternateManagementIp** | Pointer to **string** | Switch alternative management IP. To remove a prior IP setting, provide an empty string | [optional] 
**SubnetMask** | Pointer to **string** | Switch subnet mask must be in IP format. Only and must be specified for Polaris switches | [optional] 
**Gateway** | Pointer to **string** | Switch gateway must be in IP format. Only and must be specified for Polaris switches | [optional] 

## Methods

### NewGetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner

`func NewGetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner() *GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner`

NewGetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner instantiates a new GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInnerWithDefaults

`func NewGetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInnerWithDefaults() *GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner`

NewGetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInnerWithDefaults instantiates a new GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetAlternateManagementIp

`func (o *GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) GetAlternateManagementIp() string`

GetAlternateManagementIp returns the AlternateManagementIp field if non-nil, zero value otherwise.

### GetAlternateManagementIpOk

`func (o *GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) GetAlternateManagementIpOk() (*string, bool)`

GetAlternateManagementIpOk returns a tuple with the AlternateManagementIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateManagementIp

`func (o *GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) SetAlternateManagementIp(v string)`

SetAlternateManagementIp sets AlternateManagementIp field to given value.

### HasAlternateManagementIp

`func (o *GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) HasAlternateManagementIp() bool`

HasAlternateManagementIp returns a boolean if a field has been set.

### GetSubnetMask

`func (o *GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) GetSubnetMask() string`

GetSubnetMask returns the SubnetMask field if non-nil, zero value otherwise.

### GetSubnetMaskOk

`func (o *GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) GetSubnetMaskOk() (*string, bool)`

GetSubnetMaskOk returns a tuple with the SubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetMask

`func (o *GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) SetSubnetMask(v string)`

SetSubnetMask sets SubnetMask field to given value.

### HasSubnetMask

`func (o *GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) HasSubnetMask() bool`

HasSubnetMask returns a boolean if a field has been set.

### GetGateway

`func (o *GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner) HasGateway() bool`

HasGateway returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


