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

// checks if the GetNetworkHealthAlerts200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkHealthAlerts200ResponseInner{}

// GetNetworkHealthAlerts200ResponseInner struct for GetNetworkHealthAlerts200ResponseInner
type GetNetworkHealthAlerts200ResponseInner struct {
	// Alert identifier. Value can be empty
	Id *string `json:"id,omitempty"`
	// Category of the alert
	Category *string `json:"category,omitempty"`
	// Alert type
	Type *string `json:"type,omitempty"`
	// Severity of the alert
	Severity *string `json:"severity,omitempty"`
	Scope *GetNetworkHealthAlerts200ResponseInnerScope `json:"scope,omitempty"`
}

// NewGetNetworkHealthAlerts200ResponseInner instantiates a new GetNetworkHealthAlerts200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkHealthAlerts200ResponseInner() *GetNetworkHealthAlerts200ResponseInner {
	this := GetNetworkHealthAlerts200ResponseInner{}
	return &this
}

// NewGetNetworkHealthAlerts200ResponseInnerWithDefaults instantiates a new GetNetworkHealthAlerts200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkHealthAlerts200ResponseInnerWithDefaults() *GetNetworkHealthAlerts200ResponseInner {
	this := GetNetworkHealthAlerts200ResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNetworkHealthAlerts200ResponseInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkHealthAlerts200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetNetworkHealthAlerts200ResponseInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetNetworkHealthAlerts200ResponseInner) SetId(v string) {
	o.Id = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *GetNetworkHealthAlerts200ResponseInner) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkHealthAlerts200ResponseInner) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *GetNetworkHealthAlerts200ResponseInner) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *GetNetworkHealthAlerts200ResponseInner) SetCategory(v string) {
	o.Category = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetNetworkHealthAlerts200ResponseInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkHealthAlerts200ResponseInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetNetworkHealthAlerts200ResponseInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetNetworkHealthAlerts200ResponseInner) SetType(v string) {
	o.Type = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *GetNetworkHealthAlerts200ResponseInner) GetSeverity() string {
	if o == nil || IsNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkHealthAlerts200ResponseInner) GetSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *GetNetworkHealthAlerts200ResponseInner) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *GetNetworkHealthAlerts200ResponseInner) SetSeverity(v string) {
	o.Severity = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *GetNetworkHealthAlerts200ResponseInner) GetScope() GetNetworkHealthAlerts200ResponseInnerScope {
	if o == nil || IsNil(o.Scope) {
		var ret GetNetworkHealthAlerts200ResponseInnerScope
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkHealthAlerts200ResponseInner) GetScopeOk() (*GetNetworkHealthAlerts200ResponseInnerScope, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *GetNetworkHealthAlerts200ResponseInner) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given GetNetworkHealthAlerts200ResponseInnerScope and assigns it to the Scope field.
func (o *GetNetworkHealthAlerts200ResponseInner) SetScope(v GetNetworkHealthAlerts200ResponseInnerScope) {
	o.Scope = &v
}

func (o GetNetworkHealthAlerts200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkHealthAlerts200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	return toSerialize, nil
}

type NullableGetNetworkHealthAlerts200ResponseInner struct {
	value *GetNetworkHealthAlerts200ResponseInner
	isSet bool
}

func (v NullableGetNetworkHealthAlerts200ResponseInner) Get() *GetNetworkHealthAlerts200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkHealthAlerts200ResponseInner) Set(val *GetNetworkHealthAlerts200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkHealthAlerts200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkHealthAlerts200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkHealthAlerts200ResponseInner(val *GetNetworkHealthAlerts200ResponseInner) *NullableGetNetworkHealthAlerts200ResponseInner {
	return &NullableGetNetworkHealthAlerts200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkHealthAlerts200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkHealthAlerts200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


