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

// GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner struct for GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner
type GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner struct {
	// Name of the manufacturer
	Name *string `json:"name,omitempty"`
	Clients *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClients `json:"clients,omitempty"`
	Usage *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage `json:"usage,omitempty"`
}

// NewGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner instantiates a new GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner() *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner {
	this := GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner{}
	return &this
}

// NewGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerWithDefaults instantiates a new GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerWithDefaults() *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner {
	this := GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetClients returns the Clients field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner) GetClients() GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClients {
	if o == nil || isNil(o.Clients) {
		var ret GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClients
		return ret
	}
	return *o.Clients
}

// GetClientsOk returns a tuple with the Clients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner) GetClientsOk() (*GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClients, bool) {
	if o == nil || isNil(o.Clients) {
    return nil, false
	}
	return o.Clients, true
}

// HasClients returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner) HasClients() bool {
	if o != nil && !isNil(o.Clients) {
		return true
	}

	return false
}

// SetClients gets a reference to the given GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClients and assigns it to the Clients field.
func (o *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner) SetClients(v GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClients) {
	o.Clients = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner) GetUsage() GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage {
	if o == nil || isNil(o.Usage) {
		var ret GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner) GetUsageOk() (*GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage, bool) {
	if o == nil || isNil(o.Usage) {
    return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner) HasUsage() bool {
	if o != nil && !isNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage and assigns it to the Usage field.
func (o *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner) SetUsage(v GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage) {
	o.Usage = &v
}

func (o GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Clients) {
		toSerialize["clients"] = o.Clients
	}
	if !isNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	return json.Marshal(toSerialize)
}

type NullableGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner struct {
	value *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner) Get() *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner) Set(val *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner(val *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner) *NullableGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner {
	return &NullableGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


