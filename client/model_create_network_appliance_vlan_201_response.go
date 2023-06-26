/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// CreateNetworkApplianceVlan201Response struct for CreateNetworkApplianceVlan201Response
type CreateNetworkApplianceVlan201Response struct {
	// The VLAN ID of the VLAN
	Id *string `json:"id,omitempty"`
	// The interface ID of the VLAN
	InterfaceId *string `json:"interfaceId,omitempty"`
	// The name of the VLAN
	Name *string `json:"name,omitempty"`
	// The subnet of the VLAN
	Subnet *string `json:"subnet,omitempty"`
	// The local IP of the appliance on the VLAN
	ApplianceIp *string `json:"applianceIp,omitempty"`
	// The id of the desired group policy to apply to the VLAN
	GroupPolicyId *string `json:"groupPolicyId,omitempty"`
	// Type of subnetting of the VLAN. Applicable only for template network.
	TemplateVlanType *string `json:"templateVlanType,omitempty"`
	// CIDR of the pool of subnets. Applicable only for template network. Each network bound to the template will automatically pick a subnet from this pool to build its own VLAN.
	Cidr *string `json:"cidr,omitempty"`
	// Mask used for the subnet of all bound to the template networks. Applicable only for template network.
	Mask *int32 `json:"mask,omitempty"`
	MandatoryDhcp *GetNetworkApplianceVlans200ResponseInnerMandatoryDhcp `json:"mandatoryDhcp,omitempty"`
	Ipv6 *GetNetworkApplianceVlans200ResponseInnerIpv6 `json:"ipv6,omitempty"`
}

// NewCreateNetworkApplianceVlan201Response instantiates a new CreateNetworkApplianceVlan201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkApplianceVlan201Response() *CreateNetworkApplianceVlan201Response {
	this := CreateNetworkApplianceVlan201Response{}
	var templateVlanType string = "same"
	this.TemplateVlanType = &templateVlanType
	return &this
}

// NewCreateNetworkApplianceVlan201ResponseWithDefaults instantiates a new CreateNetworkApplianceVlan201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkApplianceVlan201ResponseWithDefaults() *CreateNetworkApplianceVlan201Response {
	this := CreateNetworkApplianceVlan201Response{}
	var templateVlanType string = "same"
	this.TemplateVlanType = &templateVlanType
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateNetworkApplianceVlan201Response) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceVlan201Response) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateNetworkApplianceVlan201Response) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateNetworkApplianceVlan201Response) SetId(v string) {
	o.Id = &v
}

// GetInterfaceId returns the InterfaceId field value if set, zero value otherwise.
func (o *CreateNetworkApplianceVlan201Response) GetInterfaceId() string {
	if o == nil || isNil(o.InterfaceId) {
		var ret string
		return ret
	}
	return *o.InterfaceId
}

// GetInterfaceIdOk returns a tuple with the InterfaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceVlan201Response) GetInterfaceIdOk() (*string, bool) {
	if o == nil || isNil(o.InterfaceId) {
    return nil, false
	}
	return o.InterfaceId, true
}

// HasInterfaceId returns a boolean if a field has been set.
func (o *CreateNetworkApplianceVlan201Response) HasInterfaceId() bool {
	if o != nil && !isNil(o.InterfaceId) {
		return true
	}

	return false
}

// SetInterfaceId gets a reference to the given string and assigns it to the InterfaceId field.
func (o *CreateNetworkApplianceVlan201Response) SetInterfaceId(v string) {
	o.InterfaceId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateNetworkApplianceVlan201Response) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceVlan201Response) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateNetworkApplianceVlan201Response) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateNetworkApplianceVlan201Response) SetName(v string) {
	o.Name = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *CreateNetworkApplianceVlan201Response) GetSubnet() string {
	if o == nil || isNil(o.Subnet) {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceVlan201Response) GetSubnetOk() (*string, bool) {
	if o == nil || isNil(o.Subnet) {
    return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *CreateNetworkApplianceVlan201Response) HasSubnet() bool {
	if o != nil && !isNil(o.Subnet) {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *CreateNetworkApplianceVlan201Response) SetSubnet(v string) {
	o.Subnet = &v
}

// GetApplianceIp returns the ApplianceIp field value if set, zero value otherwise.
func (o *CreateNetworkApplianceVlan201Response) GetApplianceIp() string {
	if o == nil || isNil(o.ApplianceIp) {
		var ret string
		return ret
	}
	return *o.ApplianceIp
}

// GetApplianceIpOk returns a tuple with the ApplianceIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceVlan201Response) GetApplianceIpOk() (*string, bool) {
	if o == nil || isNil(o.ApplianceIp) {
    return nil, false
	}
	return o.ApplianceIp, true
}

// HasApplianceIp returns a boolean if a field has been set.
func (o *CreateNetworkApplianceVlan201Response) HasApplianceIp() bool {
	if o != nil && !isNil(o.ApplianceIp) {
		return true
	}

	return false
}

// SetApplianceIp gets a reference to the given string and assigns it to the ApplianceIp field.
func (o *CreateNetworkApplianceVlan201Response) SetApplianceIp(v string) {
	o.ApplianceIp = &v
}

// GetGroupPolicyId returns the GroupPolicyId field value if set, zero value otherwise.
func (o *CreateNetworkApplianceVlan201Response) GetGroupPolicyId() string {
	if o == nil || isNil(o.GroupPolicyId) {
		var ret string
		return ret
	}
	return *o.GroupPolicyId
}

// GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceVlan201Response) GetGroupPolicyIdOk() (*string, bool) {
	if o == nil || isNil(o.GroupPolicyId) {
    return nil, false
	}
	return o.GroupPolicyId, true
}

// HasGroupPolicyId returns a boolean if a field has been set.
func (o *CreateNetworkApplianceVlan201Response) HasGroupPolicyId() bool {
	if o != nil && !isNil(o.GroupPolicyId) {
		return true
	}

	return false
}

// SetGroupPolicyId gets a reference to the given string and assigns it to the GroupPolicyId field.
func (o *CreateNetworkApplianceVlan201Response) SetGroupPolicyId(v string) {
	o.GroupPolicyId = &v
}

// GetTemplateVlanType returns the TemplateVlanType field value if set, zero value otherwise.
func (o *CreateNetworkApplianceVlan201Response) GetTemplateVlanType() string {
	if o == nil || isNil(o.TemplateVlanType) {
		var ret string
		return ret
	}
	return *o.TemplateVlanType
}

// GetTemplateVlanTypeOk returns a tuple with the TemplateVlanType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceVlan201Response) GetTemplateVlanTypeOk() (*string, bool) {
	if o == nil || isNil(o.TemplateVlanType) {
    return nil, false
	}
	return o.TemplateVlanType, true
}

// HasTemplateVlanType returns a boolean if a field has been set.
func (o *CreateNetworkApplianceVlan201Response) HasTemplateVlanType() bool {
	if o != nil && !isNil(o.TemplateVlanType) {
		return true
	}

	return false
}

// SetTemplateVlanType gets a reference to the given string and assigns it to the TemplateVlanType field.
func (o *CreateNetworkApplianceVlan201Response) SetTemplateVlanType(v string) {
	o.TemplateVlanType = &v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *CreateNetworkApplianceVlan201Response) GetCidr() string {
	if o == nil || isNil(o.Cidr) {
		var ret string
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceVlan201Response) GetCidrOk() (*string, bool) {
	if o == nil || isNil(o.Cidr) {
    return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *CreateNetworkApplianceVlan201Response) HasCidr() bool {
	if o != nil && !isNil(o.Cidr) {
		return true
	}

	return false
}

// SetCidr gets a reference to the given string and assigns it to the Cidr field.
func (o *CreateNetworkApplianceVlan201Response) SetCidr(v string) {
	o.Cidr = &v
}

// GetMask returns the Mask field value if set, zero value otherwise.
func (o *CreateNetworkApplianceVlan201Response) GetMask() int32 {
	if o == nil || isNil(o.Mask) {
		var ret int32
		return ret
	}
	return *o.Mask
}

// GetMaskOk returns a tuple with the Mask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceVlan201Response) GetMaskOk() (*int32, bool) {
	if o == nil || isNil(o.Mask) {
    return nil, false
	}
	return o.Mask, true
}

// HasMask returns a boolean if a field has been set.
func (o *CreateNetworkApplianceVlan201Response) HasMask() bool {
	if o != nil && !isNil(o.Mask) {
		return true
	}

	return false
}

// SetMask gets a reference to the given int32 and assigns it to the Mask field.
func (o *CreateNetworkApplianceVlan201Response) SetMask(v int32) {
	o.Mask = &v
}

// GetMandatoryDhcp returns the MandatoryDhcp field value if set, zero value otherwise.
func (o *CreateNetworkApplianceVlan201Response) GetMandatoryDhcp() GetNetworkApplianceVlans200ResponseInnerMandatoryDhcp {
	if o == nil || isNil(o.MandatoryDhcp) {
		var ret GetNetworkApplianceVlans200ResponseInnerMandatoryDhcp
		return ret
	}
	return *o.MandatoryDhcp
}

// GetMandatoryDhcpOk returns a tuple with the MandatoryDhcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceVlan201Response) GetMandatoryDhcpOk() (*GetNetworkApplianceVlans200ResponseInnerMandatoryDhcp, bool) {
	if o == nil || isNil(o.MandatoryDhcp) {
    return nil, false
	}
	return o.MandatoryDhcp, true
}

// HasMandatoryDhcp returns a boolean if a field has been set.
func (o *CreateNetworkApplianceVlan201Response) HasMandatoryDhcp() bool {
	if o != nil && !isNil(o.MandatoryDhcp) {
		return true
	}

	return false
}

// SetMandatoryDhcp gets a reference to the given GetNetworkApplianceVlans200ResponseInnerMandatoryDhcp and assigns it to the MandatoryDhcp field.
func (o *CreateNetworkApplianceVlan201Response) SetMandatoryDhcp(v GetNetworkApplianceVlans200ResponseInnerMandatoryDhcp) {
	o.MandatoryDhcp = &v
}

// GetIpv6 returns the Ipv6 field value if set, zero value otherwise.
func (o *CreateNetworkApplianceVlan201Response) GetIpv6() GetNetworkApplianceVlans200ResponseInnerIpv6 {
	if o == nil || isNil(o.Ipv6) {
		var ret GetNetworkApplianceVlans200ResponseInnerIpv6
		return ret
	}
	return *o.Ipv6
}

// GetIpv6Ok returns a tuple with the Ipv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceVlan201Response) GetIpv6Ok() (*GetNetworkApplianceVlans200ResponseInnerIpv6, bool) {
	if o == nil || isNil(o.Ipv6) {
    return nil, false
	}
	return o.Ipv6, true
}

// HasIpv6 returns a boolean if a field has been set.
func (o *CreateNetworkApplianceVlan201Response) HasIpv6() bool {
	if o != nil && !isNil(o.Ipv6) {
		return true
	}

	return false
}

// SetIpv6 gets a reference to the given GetNetworkApplianceVlans200ResponseInnerIpv6 and assigns it to the Ipv6 field.
func (o *CreateNetworkApplianceVlan201Response) SetIpv6(v GetNetworkApplianceVlans200ResponseInnerIpv6) {
	o.Ipv6 = &v
}

func (o CreateNetworkApplianceVlan201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.InterfaceId) {
		toSerialize["interfaceId"] = o.InterfaceId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Subnet) {
		toSerialize["subnet"] = o.Subnet
	}
	if !isNil(o.ApplianceIp) {
		toSerialize["applianceIp"] = o.ApplianceIp
	}
	if !isNil(o.GroupPolicyId) {
		toSerialize["groupPolicyId"] = o.GroupPolicyId
	}
	if !isNil(o.TemplateVlanType) {
		toSerialize["templateVlanType"] = o.TemplateVlanType
	}
	if !isNil(o.Cidr) {
		toSerialize["cidr"] = o.Cidr
	}
	if !isNil(o.Mask) {
		toSerialize["mask"] = o.Mask
	}
	if !isNil(o.MandatoryDhcp) {
		toSerialize["mandatoryDhcp"] = o.MandatoryDhcp
	}
	if !isNil(o.Ipv6) {
		toSerialize["ipv6"] = o.Ipv6
	}
	return json.Marshal(toSerialize)
}

type NullableCreateNetworkApplianceVlan201Response struct {
	value *CreateNetworkApplianceVlan201Response
	isSet bool
}

func (v NullableCreateNetworkApplianceVlan201Response) Get() *CreateNetworkApplianceVlan201Response {
	return v.value
}

func (v *NullableCreateNetworkApplianceVlan201Response) Set(val *CreateNetworkApplianceVlan201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkApplianceVlan201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkApplianceVlan201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkApplianceVlan201Response(val *CreateNetworkApplianceVlan201Response) *NullableCreateNetworkApplianceVlan201Response {
	return &NullableCreateNetworkApplianceVlan201Response{value: val, isSet: true}
}

func (v NullableCreateNetworkApplianceVlan201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkApplianceVlan201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


