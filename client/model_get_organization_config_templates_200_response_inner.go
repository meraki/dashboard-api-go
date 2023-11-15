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

// checks if the GetOrganizationConfigTemplates200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationConfigTemplates200ResponseInner{}

// GetOrganizationConfigTemplates200ResponseInner struct for GetOrganizationConfigTemplates200ResponseInner
type GetOrganizationConfigTemplates200ResponseInner struct {
	// The ID of the network or config template to copy configuration from
	Id *string `json:"id,omitempty"`
	// The name of the configuration template
	Name *string `json:"name,omitempty"`
	// The product types of the configuration template
	ProductTypes []string `json:"productTypes,omitempty"`
	// The timezone of the configuration template. For a list of allowed timezones, please see the 'TZ' column in the table in <a target='_blank' href='https://en.wikipedia.org/wiki/List_of_tz_database_time_zones'>this article</a>. Not applicable if copying from existing network or template
	TimeZone *string `json:"timeZone,omitempty"`
}

// NewGetOrganizationConfigTemplates200ResponseInner instantiates a new GetOrganizationConfigTemplates200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationConfigTemplates200ResponseInner() *GetOrganizationConfigTemplates200ResponseInner {
	this := GetOrganizationConfigTemplates200ResponseInner{}
	return &this
}

// NewGetOrganizationConfigTemplates200ResponseInnerWithDefaults instantiates a new GetOrganizationConfigTemplates200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationConfigTemplates200ResponseInnerWithDefaults() *GetOrganizationConfigTemplates200ResponseInner {
	this := GetOrganizationConfigTemplates200ResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetOrganizationConfigTemplates200ResponseInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationConfigTemplates200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetOrganizationConfigTemplates200ResponseInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetOrganizationConfigTemplates200ResponseInner) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetOrganizationConfigTemplates200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationConfigTemplates200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetOrganizationConfigTemplates200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetOrganizationConfigTemplates200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetProductTypes returns the ProductTypes field value if set, zero value otherwise.
func (o *GetOrganizationConfigTemplates200ResponseInner) GetProductTypes() []string {
	if o == nil || IsNil(o.ProductTypes) {
		var ret []string
		return ret
	}
	return o.ProductTypes
}

// GetProductTypesOk returns a tuple with the ProductTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationConfigTemplates200ResponseInner) GetProductTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.ProductTypes) {
		return nil, false
	}
	return o.ProductTypes, true
}

// HasProductTypes returns a boolean if a field has been set.
func (o *GetOrganizationConfigTemplates200ResponseInner) HasProductTypes() bool {
	if o != nil && !IsNil(o.ProductTypes) {
		return true
	}

	return false
}

// SetProductTypes gets a reference to the given []string and assigns it to the ProductTypes field.
func (o *GetOrganizationConfigTemplates200ResponseInner) SetProductTypes(v []string) {
	o.ProductTypes = v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *GetOrganizationConfigTemplates200ResponseInner) GetTimeZone() string {
	if o == nil || IsNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationConfigTemplates200ResponseInner) GetTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *GetOrganizationConfigTemplates200ResponseInner) HasTimeZone() bool {
	if o != nil && !IsNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *GetOrganizationConfigTemplates200ResponseInner) SetTimeZone(v string) {
	o.TimeZone = &v
}

func (o GetOrganizationConfigTemplates200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationConfigTemplates200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ProductTypes) {
		toSerialize["productTypes"] = o.ProductTypes
	}
	if !IsNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}
	return toSerialize, nil
}

type NullableGetOrganizationConfigTemplates200ResponseInner struct {
	value *GetOrganizationConfigTemplates200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationConfigTemplates200ResponseInner) Get() *GetOrganizationConfigTemplates200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationConfigTemplates200ResponseInner) Set(val *GetOrganizationConfigTemplates200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationConfigTemplates200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationConfigTemplates200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationConfigTemplates200ResponseInner(val *GetOrganizationConfigTemplates200ResponseInner) *NullableGetOrganizationConfigTemplates200ResponseInner {
	return &NullableGetOrganizationConfigTemplates200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationConfigTemplates200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationConfigTemplates200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


