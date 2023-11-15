/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner{}

// GetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner struct for GetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner
type GetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner struct {
	// Uplink serial
	Serial *string `json:"serial,omitempty"`
	// Uplink name
	Interface *string `json:"interface,omitempty"`
	// Bytes sent
	Sent *int32 `json:"sent,omitempty"`
	// Bytes received
	Received *int32 `json:"received,omitempty"`
}

// NewGetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner instantiates a new GetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner() *GetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner {
	this := GetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner{}
	return &this
}

// NewGetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInnerWithDefaults instantiates a new GetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInnerWithDefaults() *GetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner {
	this := GetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *GetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *GetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *GetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner) SetSerial(v string) {
	o.Serial = &v
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *GetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner) GetInterface() string {
	if o == nil || IsNil(o.Interface) {
		var ret string
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner) GetInterfaceOk() (*string, bool) {
	if o == nil || IsNil(o.Interface) {
		return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *GetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner) HasInterface() bool {
	if o != nil && !IsNil(o.Interface) {
		return true
	}

	return false
}

// SetInterface gets a reference to the given string and assigns it to the Interface field.
func (o *GetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner) SetInterface(v string) {
	o.Interface = &v
}

// GetSent returns the Sent field value if set, zero value otherwise.
func (o *GetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner) GetSent() int32 {
	if o == nil || IsNil(o.Sent) {
		var ret int32
		return ret
	}
	return *o.Sent
}

// GetSentOk returns a tuple with the Sent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner) GetSentOk() (*int32, bool) {
	if o == nil || IsNil(o.Sent) {
		return nil, false
	}
	return o.Sent, true
}

// HasSent returns a boolean if a field has been set.
func (o *GetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner) HasSent() bool {
	if o != nil && !IsNil(o.Sent) {
		return true
	}

	return false
}

// SetSent gets a reference to the given int32 and assigns it to the Sent field.
func (o *GetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner) SetSent(v int32) {
	o.Sent = &v
}

// GetReceived returns the Received field value if set, zero value otherwise.
func (o *GetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner) GetReceived() int32 {
	if o == nil || IsNil(o.Received) {
		var ret int32
		return ret
	}
	return *o.Received
}

// GetReceivedOk returns a tuple with the Received field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner) GetReceivedOk() (*int32, bool) {
	if o == nil || IsNil(o.Received) {
		return nil, false
	}
	return o.Received, true
}

// HasReceived returns a boolean if a field has been set.
func (o *GetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner) HasReceived() bool {
	if o != nil && !IsNil(o.Received) {
		return true
	}

	return false
}

// SetReceived gets a reference to the given int32 and assigns it to the Received field.
func (o *GetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner) SetReceived(v int32) {
	o.Received = &v
}

func (o GetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !IsNil(o.Interface) {
		toSerialize["interface"] = o.Interface
	}
	if !IsNil(o.Sent) {
		toSerialize["sent"] = o.Sent
	}
	if !IsNil(o.Received) {
		toSerialize["received"] = o.Received
	}
	return toSerialize, nil
}

type NullableGetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner struct {
	value *GetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner
	isSet bool
}

func (v NullableGetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner) Get() *GetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner {
	return v.value
}

func (v *NullableGetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner) Set(val *GetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner(val *GetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner) *NullableGetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner {
	return &NullableGetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner{value: val, isSet: true}
}

func (v NullableGetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationApplianceUplinksUsageByNetwork200ResponseInnerByUplinkInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


