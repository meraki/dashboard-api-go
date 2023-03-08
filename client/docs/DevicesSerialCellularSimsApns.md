# DevicesSerialCellularSimsApns

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | APN name. | 
**AllowedIpTypes** | **[]string** | IP versions to support (permitted values include &#39;ipv4&#39;, &#39;ipv6&#39;). | 
**Authentication** | Pointer to [**DevicesSerialCellularSimsAuthentication**](DevicesSerialCellularSimsAuthentication.md) |  | [optional] 

## Methods

### NewDevicesSerialCellularSimsApns

`func NewDevicesSerialCellularSimsApns(name string, allowedIpTypes []string, ) *DevicesSerialCellularSimsApns`

NewDevicesSerialCellularSimsApns instantiates a new DevicesSerialCellularSimsApns object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicesSerialCellularSimsApnsWithDefaults

`func NewDevicesSerialCellularSimsApnsWithDefaults() *DevicesSerialCellularSimsApns`

NewDevicesSerialCellularSimsApnsWithDefaults instantiates a new DevicesSerialCellularSimsApns object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DevicesSerialCellularSimsApns) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DevicesSerialCellularSimsApns) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DevicesSerialCellularSimsApns) SetName(v string)`

SetName sets Name field to given value.


### GetAllowedIpTypes

`func (o *DevicesSerialCellularSimsApns) GetAllowedIpTypes() []string`

GetAllowedIpTypes returns the AllowedIpTypes field if non-nil, zero value otherwise.

### GetAllowedIpTypesOk

`func (o *DevicesSerialCellularSimsApns) GetAllowedIpTypesOk() (*[]string, bool)`

GetAllowedIpTypesOk returns a tuple with the AllowedIpTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedIpTypes

`func (o *DevicesSerialCellularSimsApns) SetAllowedIpTypes(v []string)`

SetAllowedIpTypes sets AllowedIpTypes field to given value.


### GetAuthentication

`func (o *DevicesSerialCellularSimsApns) GetAuthentication() DevicesSerialCellularSimsAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *DevicesSerialCellularSimsApns) GetAuthenticationOk() (*DevicesSerialCellularSimsAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *DevicesSerialCellularSimsApns) SetAuthentication(v DevicesSerialCellularSimsAuthentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *DevicesSerialCellularSimsApns) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


