# CreateNetworkApplianceVlan201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The VLAN ID of the VLAN | [optional] 
**InterfaceId** | Pointer to **string** | The interface ID of the VLAN | [optional] 
**Name** | Pointer to **string** | The name of the VLAN | [optional] 
**Subnet** | Pointer to **string** | The subnet of the VLAN | [optional] 
**ApplianceIp** | Pointer to **string** | The local IP of the appliance on the VLAN | [optional] 
**GroupPolicyId** | Pointer to **string** | The id of the desired group policy to apply to the VLAN | [optional] 
**TemplateVlanType** | Pointer to **string** | Type of subnetting of the VLAN. Applicable only for template network. | [optional] [default to "same"]
**Cidr** | Pointer to **string** | CIDR of the pool of subnets. Applicable only for template network. Each network bound to the template will automatically pick a subnet from this pool to build its own VLAN. | [optional] 
**Mask** | Pointer to **int32** | Mask used for the subnet of all bound to the template networks. Applicable only for template network. | [optional] 
**MandatoryDhcp** | Pointer to [**GetNetworkApplianceVlans200ResponseInnerMandatoryDhcp**](GetNetworkApplianceVlans200ResponseInnerMandatoryDhcp.md) |  | [optional] 
**Ipv6** | Pointer to [**GetNetworkApplianceVlans200ResponseInnerIpv6**](GetNetworkApplianceVlans200ResponseInnerIpv6.md) |  | [optional] 

## Methods

### NewCreateNetworkApplianceVlan201Response

`func NewCreateNetworkApplianceVlan201Response() *CreateNetworkApplianceVlan201Response`

NewCreateNetworkApplianceVlan201Response instantiates a new CreateNetworkApplianceVlan201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkApplianceVlan201ResponseWithDefaults

`func NewCreateNetworkApplianceVlan201ResponseWithDefaults() *CreateNetworkApplianceVlan201Response`

NewCreateNetworkApplianceVlan201ResponseWithDefaults instantiates a new CreateNetworkApplianceVlan201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateNetworkApplianceVlan201Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateNetworkApplianceVlan201Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateNetworkApplianceVlan201Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateNetworkApplianceVlan201Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterfaceId

`func (o *CreateNetworkApplianceVlan201Response) GetInterfaceId() string`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *CreateNetworkApplianceVlan201Response) GetInterfaceIdOk() (*string, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *CreateNetworkApplianceVlan201Response) SetInterfaceId(v string)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *CreateNetworkApplianceVlan201Response) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### GetName

`func (o *CreateNetworkApplianceVlan201Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkApplianceVlan201Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkApplianceVlan201Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateNetworkApplianceVlan201Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubnet

`func (o *CreateNetworkApplianceVlan201Response) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *CreateNetworkApplianceVlan201Response) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *CreateNetworkApplianceVlan201Response) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *CreateNetworkApplianceVlan201Response) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetApplianceIp

`func (o *CreateNetworkApplianceVlan201Response) GetApplianceIp() string`

GetApplianceIp returns the ApplianceIp field if non-nil, zero value otherwise.

### GetApplianceIpOk

`func (o *CreateNetworkApplianceVlan201Response) GetApplianceIpOk() (*string, bool)`

GetApplianceIpOk returns a tuple with the ApplianceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceIp

`func (o *CreateNetworkApplianceVlan201Response) SetApplianceIp(v string)`

SetApplianceIp sets ApplianceIp field to given value.

### HasApplianceIp

`func (o *CreateNetworkApplianceVlan201Response) HasApplianceIp() bool`

HasApplianceIp returns a boolean if a field has been set.

### GetGroupPolicyId

`func (o *CreateNetworkApplianceVlan201Response) GetGroupPolicyId() string`

GetGroupPolicyId returns the GroupPolicyId field if non-nil, zero value otherwise.

### GetGroupPolicyIdOk

`func (o *CreateNetworkApplianceVlan201Response) GetGroupPolicyIdOk() (*string, bool)`

GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPolicyId

`func (o *CreateNetworkApplianceVlan201Response) SetGroupPolicyId(v string)`

SetGroupPolicyId sets GroupPolicyId field to given value.

### HasGroupPolicyId

`func (o *CreateNetworkApplianceVlan201Response) HasGroupPolicyId() bool`

HasGroupPolicyId returns a boolean if a field has been set.

### GetTemplateVlanType

`func (o *CreateNetworkApplianceVlan201Response) GetTemplateVlanType() string`

GetTemplateVlanType returns the TemplateVlanType field if non-nil, zero value otherwise.

### GetTemplateVlanTypeOk

`func (o *CreateNetworkApplianceVlan201Response) GetTemplateVlanTypeOk() (*string, bool)`

GetTemplateVlanTypeOk returns a tuple with the TemplateVlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateVlanType

`func (o *CreateNetworkApplianceVlan201Response) SetTemplateVlanType(v string)`

SetTemplateVlanType sets TemplateVlanType field to given value.

### HasTemplateVlanType

`func (o *CreateNetworkApplianceVlan201Response) HasTemplateVlanType() bool`

HasTemplateVlanType returns a boolean if a field has been set.

### GetCidr

`func (o *CreateNetworkApplianceVlan201Response) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *CreateNetworkApplianceVlan201Response) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *CreateNetworkApplianceVlan201Response) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *CreateNetworkApplianceVlan201Response) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetMask

`func (o *CreateNetworkApplianceVlan201Response) GetMask() int32`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *CreateNetworkApplianceVlan201Response) GetMaskOk() (*int32, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *CreateNetworkApplianceVlan201Response) SetMask(v int32)`

SetMask sets Mask field to given value.

### HasMask

`func (o *CreateNetworkApplianceVlan201Response) HasMask() bool`

HasMask returns a boolean if a field has been set.

### GetMandatoryDhcp

`func (o *CreateNetworkApplianceVlan201Response) GetMandatoryDhcp() GetNetworkApplianceVlans200ResponseInnerMandatoryDhcp`

GetMandatoryDhcp returns the MandatoryDhcp field if non-nil, zero value otherwise.

### GetMandatoryDhcpOk

`func (o *CreateNetworkApplianceVlan201Response) GetMandatoryDhcpOk() (*GetNetworkApplianceVlans200ResponseInnerMandatoryDhcp, bool)`

GetMandatoryDhcpOk returns a tuple with the MandatoryDhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatoryDhcp

`func (o *CreateNetworkApplianceVlan201Response) SetMandatoryDhcp(v GetNetworkApplianceVlans200ResponseInnerMandatoryDhcp)`

SetMandatoryDhcp sets MandatoryDhcp field to given value.

### HasMandatoryDhcp

`func (o *CreateNetworkApplianceVlan201Response) HasMandatoryDhcp() bool`

HasMandatoryDhcp returns a boolean if a field has been set.

### GetIpv6

`func (o *CreateNetworkApplianceVlan201Response) GetIpv6() GetNetworkApplianceVlans200ResponseInnerIpv6`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *CreateNetworkApplianceVlan201Response) GetIpv6Ok() (*GetNetworkApplianceVlans200ResponseInnerIpv6, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *CreateNetworkApplianceVlan201Response) SetIpv6(v GetNetworkApplianceVlans200ResponseInnerIpv6)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *CreateNetworkApplianceVlan201Response) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


