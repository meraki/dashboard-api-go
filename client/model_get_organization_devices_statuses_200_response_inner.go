/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 02 August, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetOrganizationDevicesStatuses200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationDevicesStatuses200ResponseInner{}

// GetOrganizationDevicesStatuses200ResponseInner struct for GetOrganizationDevicesStatuses200ResponseInner
type GetOrganizationDevicesStatuses200ResponseInner struct {
	// Device Name
	Name *string `json:"name,omitempty"`
	// Device Serial Number
	Serial *string `json:"serial,omitempty"`
	// MAC Address
	Mac *string `json:"mac,omitempty"`
	// Public IP Address
	PublicIp *string `json:"publicIp,omitempty"`
	// Network ID
	NetworkId *string `json:"networkId,omitempty"`
	// Device Status
	Status *string `json:"status,omitempty"`
	// Device Last Reported Location
	LastReportedAt *string `json:"lastReportedAt,omitempty"`
	// LAN IP Address
	LanIp *string `json:"lanIp,omitempty"`
	// IP Gateway
	Gateway *string `json:"gateway,omitempty"`
	// IP Type
	IpType *string `json:"ipType,omitempty"`
	// Primary DNS
	PrimaryDns *string `json:"primaryDns,omitempty"`
	// Secondary DNS
	SecondaryDns *string `json:"secondaryDns,omitempty"`
	// Product Type
	ProductType *string `json:"productType,omitempty"`
	Components *GetOrganizationDevicesStatuses200ResponseInnerComponents `json:"components,omitempty"`
	// Model
	Model *string `json:"model,omitempty"`
	// Tags
	Tags []string `json:"tags,omitempty"`
}

// NewGetOrganizationDevicesStatuses200ResponseInner instantiates a new GetOrganizationDevicesStatuses200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationDevicesStatuses200ResponseInner() *GetOrganizationDevicesStatuses200ResponseInner {
	this := GetOrganizationDevicesStatuses200ResponseInner{}
	return &this
}

// NewGetOrganizationDevicesStatuses200ResponseInnerWithDefaults instantiates a new GetOrganizationDevicesStatuses200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationDevicesStatuses200ResponseInnerWithDefaults() *GetOrganizationDevicesStatuses200ResponseInner {
	this := GetOrganizationDevicesStatuses200ResponseInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetOrganizationDevicesStatuses200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetOrganizationDevicesStatuses200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *GetOrganizationDevicesStatuses200ResponseInner) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInner) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInner) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *GetOrganizationDevicesStatuses200ResponseInner) SetSerial(v string) {
	o.Serial = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *GetOrganizationDevicesStatuses200ResponseInner) GetMac() string {
	if o == nil || IsNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInner) GetMacOk() (*string, bool) {
	if o == nil || IsNil(o.Mac) {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInner) HasMac() bool {
	if o != nil && !IsNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *GetOrganizationDevicesStatuses200ResponseInner) SetMac(v string) {
	o.Mac = &v
}

// GetPublicIp returns the PublicIp field value if set, zero value otherwise.
func (o *GetOrganizationDevicesStatuses200ResponseInner) GetPublicIp() string {
	if o == nil || IsNil(o.PublicIp) {
		var ret string
		return ret
	}
	return *o.PublicIp
}

// GetPublicIpOk returns a tuple with the PublicIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInner) GetPublicIpOk() (*string, bool) {
	if o == nil || IsNil(o.PublicIp) {
		return nil, false
	}
	return o.PublicIp, true
}

// HasPublicIp returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInner) HasPublicIp() bool {
	if o != nil && !IsNil(o.PublicIp) {
		return true
	}

	return false
}

// SetPublicIp gets a reference to the given string and assigns it to the PublicIp field.
func (o *GetOrganizationDevicesStatuses200ResponseInner) SetPublicIp(v string) {
	o.PublicIp = &v
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *GetOrganizationDevicesStatuses200ResponseInner) GetNetworkId() string {
	if o == nil || IsNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInner) GetNetworkIdOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkId) {
		return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInner) HasNetworkId() bool {
	if o != nil && !IsNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *GetOrganizationDevicesStatuses200ResponseInner) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetOrganizationDevicesStatuses200ResponseInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetOrganizationDevicesStatuses200ResponseInner) SetStatus(v string) {
	o.Status = &v
}

// GetLastReportedAt returns the LastReportedAt field value if set, zero value otherwise.
func (o *GetOrganizationDevicesStatuses200ResponseInner) GetLastReportedAt() string {
	if o == nil || IsNil(o.LastReportedAt) {
		var ret string
		return ret
	}
	return *o.LastReportedAt
}

// GetLastReportedAtOk returns a tuple with the LastReportedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInner) GetLastReportedAtOk() (*string, bool) {
	if o == nil || IsNil(o.LastReportedAt) {
		return nil, false
	}
	return o.LastReportedAt, true
}

// HasLastReportedAt returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInner) HasLastReportedAt() bool {
	if o != nil && !IsNil(o.LastReportedAt) {
		return true
	}

	return false
}

// SetLastReportedAt gets a reference to the given string and assigns it to the LastReportedAt field.
func (o *GetOrganizationDevicesStatuses200ResponseInner) SetLastReportedAt(v string) {
	o.LastReportedAt = &v
}

// GetLanIp returns the LanIp field value if set, zero value otherwise.
func (o *GetOrganizationDevicesStatuses200ResponseInner) GetLanIp() string {
	if o == nil || IsNil(o.LanIp) {
		var ret string
		return ret
	}
	return *o.LanIp
}

// GetLanIpOk returns a tuple with the LanIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInner) GetLanIpOk() (*string, bool) {
	if o == nil || IsNil(o.LanIp) {
		return nil, false
	}
	return o.LanIp, true
}

// HasLanIp returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInner) HasLanIp() bool {
	if o != nil && !IsNil(o.LanIp) {
		return true
	}

	return false
}

// SetLanIp gets a reference to the given string and assigns it to the LanIp field.
func (o *GetOrganizationDevicesStatuses200ResponseInner) SetLanIp(v string) {
	o.LanIp = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *GetOrganizationDevicesStatuses200ResponseInner) GetGateway() string {
	if o == nil || IsNil(o.Gateway) {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInner) GetGatewayOk() (*string, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInner) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *GetOrganizationDevicesStatuses200ResponseInner) SetGateway(v string) {
	o.Gateway = &v
}

// GetIpType returns the IpType field value if set, zero value otherwise.
func (o *GetOrganizationDevicesStatuses200ResponseInner) GetIpType() string {
	if o == nil || IsNil(o.IpType) {
		var ret string
		return ret
	}
	return *o.IpType
}

// GetIpTypeOk returns a tuple with the IpType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInner) GetIpTypeOk() (*string, bool) {
	if o == nil || IsNil(o.IpType) {
		return nil, false
	}
	return o.IpType, true
}

// HasIpType returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInner) HasIpType() bool {
	if o != nil && !IsNil(o.IpType) {
		return true
	}

	return false
}

// SetIpType gets a reference to the given string and assigns it to the IpType field.
func (o *GetOrganizationDevicesStatuses200ResponseInner) SetIpType(v string) {
	o.IpType = &v
}

// GetPrimaryDns returns the PrimaryDns field value if set, zero value otherwise.
func (o *GetOrganizationDevicesStatuses200ResponseInner) GetPrimaryDns() string {
	if o == nil || IsNil(o.PrimaryDns) {
		var ret string
		return ret
	}
	return *o.PrimaryDns
}

// GetPrimaryDnsOk returns a tuple with the PrimaryDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInner) GetPrimaryDnsOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryDns) {
		return nil, false
	}
	return o.PrimaryDns, true
}

// HasPrimaryDns returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInner) HasPrimaryDns() bool {
	if o != nil && !IsNil(o.PrimaryDns) {
		return true
	}

	return false
}

// SetPrimaryDns gets a reference to the given string and assigns it to the PrimaryDns field.
func (o *GetOrganizationDevicesStatuses200ResponseInner) SetPrimaryDns(v string) {
	o.PrimaryDns = &v
}

// GetSecondaryDns returns the SecondaryDns field value if set, zero value otherwise.
func (o *GetOrganizationDevicesStatuses200ResponseInner) GetSecondaryDns() string {
	if o == nil || IsNil(o.SecondaryDns) {
		var ret string
		return ret
	}
	return *o.SecondaryDns
}

// GetSecondaryDnsOk returns a tuple with the SecondaryDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInner) GetSecondaryDnsOk() (*string, bool) {
	if o == nil || IsNil(o.SecondaryDns) {
		return nil, false
	}
	return o.SecondaryDns, true
}

// HasSecondaryDns returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInner) HasSecondaryDns() bool {
	if o != nil && !IsNil(o.SecondaryDns) {
		return true
	}

	return false
}

// SetSecondaryDns gets a reference to the given string and assigns it to the SecondaryDns field.
func (o *GetOrganizationDevicesStatuses200ResponseInner) SetSecondaryDns(v string) {
	o.SecondaryDns = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *GetOrganizationDevicesStatuses200ResponseInner) GetProductType() string {
	if o == nil || IsNil(o.ProductType) {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInner) GetProductTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProductType) {
		return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInner) HasProductType() bool {
	if o != nil && !IsNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *GetOrganizationDevicesStatuses200ResponseInner) SetProductType(v string) {
	o.ProductType = &v
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *GetOrganizationDevicesStatuses200ResponseInner) GetComponents() GetOrganizationDevicesStatuses200ResponseInnerComponents {
	if o == nil || IsNil(o.Components) {
		var ret GetOrganizationDevicesStatuses200ResponseInnerComponents
		return ret
	}
	return *o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInner) GetComponentsOk() (*GetOrganizationDevicesStatuses200ResponseInnerComponents, bool) {
	if o == nil || IsNil(o.Components) {
		return nil, false
	}
	return o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInner) HasComponents() bool {
	if o != nil && !IsNil(o.Components) {
		return true
	}

	return false
}

// SetComponents gets a reference to the given GetOrganizationDevicesStatuses200ResponseInnerComponents and assigns it to the Components field.
func (o *GetOrganizationDevicesStatuses200ResponseInner) SetComponents(v GetOrganizationDevicesStatuses200ResponseInnerComponents) {
	o.Components = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *GetOrganizationDevicesStatuses200ResponseInner) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInner) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInner) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *GetOrganizationDevicesStatuses200ResponseInner) SetModel(v string) {
	o.Model = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GetOrganizationDevicesStatuses200ResponseInner) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInner) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInner) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GetOrganizationDevicesStatuses200ResponseInner) SetTags(v []string) {
	o.Tags = v
}

func (o GetOrganizationDevicesStatuses200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationDevicesStatuses200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !IsNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !IsNil(o.PublicIp) {
		toSerialize["publicIp"] = o.PublicIp
	}
	if !IsNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.LastReportedAt) {
		toSerialize["lastReportedAt"] = o.LastReportedAt
	}
	if !IsNil(o.LanIp) {
		toSerialize["lanIp"] = o.LanIp
	}
	if !IsNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	if !IsNil(o.IpType) {
		toSerialize["ipType"] = o.IpType
	}
	if !IsNil(o.PrimaryDns) {
		toSerialize["primaryDns"] = o.PrimaryDns
	}
	if !IsNil(o.SecondaryDns) {
		toSerialize["secondaryDns"] = o.SecondaryDns
	}
	if !IsNil(o.ProductType) {
		toSerialize["productType"] = o.ProductType
	}
	if !IsNil(o.Components) {
		toSerialize["components"] = o.Components
	}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableGetOrganizationDevicesStatuses200ResponseInner struct {
	value *GetOrganizationDevicesStatuses200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationDevicesStatuses200ResponseInner) Get() *GetOrganizationDevicesStatuses200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationDevicesStatuses200ResponseInner) Set(val *GetOrganizationDevicesStatuses200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationDevicesStatuses200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationDevicesStatuses200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationDevicesStatuses200ResponseInner(val *GetOrganizationDevicesStatuses200ResponseInner) *NullableGetOrganizationDevicesStatuses200ResponseInner {
	return &NullableGetOrganizationDevicesStatuses200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationDevicesStatuses200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationDevicesStatuses200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


