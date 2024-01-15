/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetOrganizationDevicesProvisioningStatuses200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationDevicesProvisioningStatuses200ResponseInner{}

// GetOrganizationDevicesProvisioningStatuses200ResponseInner struct for GetOrganizationDevicesProvisioningStatuses200ResponseInner
type GetOrganizationDevicesProvisioningStatuses200ResponseInner struct {
	// The device MAC address.
	Mac *string `json:"mac,omitempty"`
	// The device name.
	Name *string `json:"name,omitempty"`
	Network *GetOrganizationDevicesAvailabilities200ResponseInnerNetwork `json:"network,omitempty"`
	// Device product type.
	ProductType *string `json:"productType,omitempty"`
	// The device serial number.
	Serial *string `json:"serial,omitempty"`
	// The device provisioning status. Possible statuses: unprovisioned, incomplete, complete.
	Status *string `json:"status,omitempty"`
	// List of custom tags for the device.
	Tags []string `json:"tags,omitempty"`
}

// NewGetOrganizationDevicesProvisioningStatuses200ResponseInner instantiates a new GetOrganizationDevicesProvisioningStatuses200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationDevicesProvisioningStatuses200ResponseInner() *GetOrganizationDevicesProvisioningStatuses200ResponseInner {
	this := GetOrganizationDevicesProvisioningStatuses200ResponseInner{}
	return &this
}

// NewGetOrganizationDevicesProvisioningStatuses200ResponseInnerWithDefaults instantiates a new GetOrganizationDevicesProvisioningStatuses200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationDevicesProvisioningStatuses200ResponseInnerWithDefaults() *GetOrganizationDevicesProvisioningStatuses200ResponseInner {
	this := GetOrganizationDevicesProvisioningStatuses200ResponseInner{}
	return &this
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) GetMac() string {
	if o == nil || IsNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) GetMacOk() (*string, bool) {
	if o == nil || IsNil(o.Mac) {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) HasMac() bool {
	if o != nil && !IsNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) SetMac(v string) {
	o.Mac = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) GetNetwork() GetOrganizationDevicesAvailabilities200ResponseInnerNetwork {
	if o == nil || IsNil(o.Network) {
		var ret GetOrganizationDevicesAvailabilities200ResponseInnerNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) GetNetworkOk() (*GetOrganizationDevicesAvailabilities200ResponseInnerNetwork, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given GetOrganizationDevicesAvailabilities200ResponseInnerNetwork and assigns it to the Network field.
func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) SetNetwork(v GetOrganizationDevicesAvailabilities200ResponseInnerNetwork) {
	o.Network = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) GetProductType() string {
	if o == nil || IsNil(o.ProductType) {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) GetProductTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProductType) {
		return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) HasProductType() bool {
	if o != nil && !IsNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) SetProductType(v string) {
	o.ProductType = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) SetSerial(v string) {
	o.Serial = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) SetStatus(v string) {
	o.Status = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GetOrganizationDevicesProvisioningStatuses200ResponseInner) SetTags(v []string) {
	o.Tags = v
}

func (o GetOrganizationDevicesProvisioningStatuses200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationDevicesProvisioningStatuses200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.ProductType) {
		toSerialize["productType"] = o.ProductType
	}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableGetOrganizationDevicesProvisioningStatuses200ResponseInner struct {
	value *GetOrganizationDevicesProvisioningStatuses200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationDevicesProvisioningStatuses200ResponseInner) Get() *GetOrganizationDevicesProvisioningStatuses200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationDevicesProvisioningStatuses200ResponseInner) Set(val *GetOrganizationDevicesProvisioningStatuses200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationDevicesProvisioningStatuses200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationDevicesProvisioningStatuses200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationDevicesProvisioningStatuses200ResponseInner(val *GetOrganizationDevicesProvisioningStatuses200ResponseInner) *NullableGetOrganizationDevicesProvisioningStatuses200ResponseInner {
	return &NullableGetOrganizationDevicesProvisioningStatuses200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationDevicesProvisioningStatuses200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationDevicesProvisioningStatuses200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


