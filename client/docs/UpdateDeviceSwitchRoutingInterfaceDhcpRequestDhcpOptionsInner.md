# UpdateDeviceSwitchRoutingInterfaceDhcpRequestDhcpOptionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | The code for DHCP option which should be from 2 to 254 | 
**Type** | **string** | The type of the DHCP option which should be one of (&#39;text&#39;, &#39;ip&#39;, &#39;integer&#39; or &#39;hex&#39;) | 
**Value** | **string** | The value of the DHCP option | 

## Methods

### NewUpdateDeviceSwitchRoutingInterfaceDhcpRequestDhcpOptionsInner

`func NewUpdateDeviceSwitchRoutingInterfaceDhcpRequestDhcpOptionsInner(code string, type_ string, value string, ) *UpdateDeviceSwitchRoutingInterfaceDhcpRequestDhcpOptionsInner`

NewUpdateDeviceSwitchRoutingInterfaceDhcpRequestDhcpOptionsInner instantiates a new UpdateDeviceSwitchRoutingInterfaceDhcpRequestDhcpOptionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeviceSwitchRoutingInterfaceDhcpRequestDhcpOptionsInnerWithDefaults

`func NewUpdateDeviceSwitchRoutingInterfaceDhcpRequestDhcpOptionsInnerWithDefaults() *UpdateDeviceSwitchRoutingInterfaceDhcpRequestDhcpOptionsInner`

NewUpdateDeviceSwitchRoutingInterfaceDhcpRequestDhcpOptionsInnerWithDefaults instantiates a new UpdateDeviceSwitchRoutingInterfaceDhcpRequestDhcpOptionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequestDhcpOptionsInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequestDhcpOptionsInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequestDhcpOptionsInner) SetCode(v string)`

SetCode sets Code field to given value.


### GetType

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequestDhcpOptionsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequestDhcpOptionsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequestDhcpOptionsInner) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequestDhcpOptionsInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequestDhcpOptionsInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequestDhcpOptionsInner) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


