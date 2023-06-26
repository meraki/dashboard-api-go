# InlineResponse20069Switches

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Switch serial number | [optional] 
**AlternateManagementIp** | Pointer to **string** | Switch alternative management IP. To remove a prior IP setting, provide an empty string | [optional] 
**SubnetMask** | Pointer to **string** | Switch subnet mask must be in IP format. Only and must be specified for Polaris switches | [optional] 
**Gateway** | Pointer to **string** | Switch gateway must be in IP format. Only and must be specified for Polaris switches | [optional] 

## Methods

### NewInlineResponse20069Switches

`func NewInlineResponse20069Switches() *InlineResponse20069Switches`

NewInlineResponse20069Switches instantiates a new InlineResponse20069Switches object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20069SwitchesWithDefaults

`func NewInlineResponse20069SwitchesWithDefaults() *InlineResponse20069Switches`

NewInlineResponse20069SwitchesWithDefaults instantiates a new InlineResponse20069Switches object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse20069Switches) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse20069Switches) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse20069Switches) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse20069Switches) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetAlternateManagementIp

`func (o *InlineResponse20069Switches) GetAlternateManagementIp() string`

GetAlternateManagementIp returns the AlternateManagementIp field if non-nil, zero value otherwise.

### GetAlternateManagementIpOk

`func (o *InlineResponse20069Switches) GetAlternateManagementIpOk() (*string, bool)`

GetAlternateManagementIpOk returns a tuple with the AlternateManagementIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateManagementIp

`func (o *InlineResponse20069Switches) SetAlternateManagementIp(v string)`

SetAlternateManagementIp sets AlternateManagementIp field to given value.

### HasAlternateManagementIp

`func (o *InlineResponse20069Switches) HasAlternateManagementIp() bool`

HasAlternateManagementIp returns a boolean if a field has been set.

### GetSubnetMask

`func (o *InlineResponse20069Switches) GetSubnetMask() string`

GetSubnetMask returns the SubnetMask field if non-nil, zero value otherwise.

### GetSubnetMaskOk

`func (o *InlineResponse20069Switches) GetSubnetMaskOk() (*string, bool)`

GetSubnetMaskOk returns a tuple with the SubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetMask

`func (o *InlineResponse20069Switches) SetSubnetMask(v string)`

SetSubnetMask sets SubnetMask field to given value.

### HasSubnetMask

`func (o *InlineResponse20069Switches) HasSubnetMask() bool`

HasSubnetMask returns a boolean if a field has been set.

### GetGateway

`func (o *InlineResponse20069Switches) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *InlineResponse20069Switches) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *InlineResponse20069Switches) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *InlineResponse20069Switches) HasGateway() bool`

HasGateway returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


