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

// checks if the GetOrganizationActionBatches200ResponseInnerActionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationActionBatches200ResponseInnerActionsInner{}

// GetOrganizationActionBatches200ResponseInnerActionsInner struct for GetOrganizationActionBatches200ResponseInnerActionsInner
type GetOrganizationActionBatches200ResponseInnerActionsInner struct {
	// Unique identifier for the resource to be acted on
	Resource string `json:"resource"`
	// The operation to be used by this action
	Operation string `json:"operation"`
	// Data provided in the body of the Action. Contents depend on the Action type
	Body map[string]interface{} `json:"body,omitempty"`
}

// NewGetOrganizationActionBatches200ResponseInnerActionsInner instantiates a new GetOrganizationActionBatches200ResponseInnerActionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationActionBatches200ResponseInnerActionsInner(resource string, operation string) *GetOrganizationActionBatches200ResponseInnerActionsInner {
	this := GetOrganizationActionBatches200ResponseInnerActionsInner{}
	this.Resource = resource
	this.Operation = operation
	return &this
}

// NewGetOrganizationActionBatches200ResponseInnerActionsInnerWithDefaults instantiates a new GetOrganizationActionBatches200ResponseInnerActionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationActionBatches200ResponseInnerActionsInnerWithDefaults() *GetOrganizationActionBatches200ResponseInnerActionsInner {
	this := GetOrganizationActionBatches200ResponseInnerActionsInner{}
	return &this
}

// GetResource returns the Resource field value
func (o *GetOrganizationActionBatches200ResponseInnerActionsInner) GetResource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *GetOrganizationActionBatches200ResponseInnerActionsInner) GetResourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *GetOrganizationActionBatches200ResponseInnerActionsInner) SetResource(v string) {
	o.Resource = v
}

// GetOperation returns the Operation field value
func (o *GetOrganizationActionBatches200ResponseInnerActionsInner) GetOperation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *GetOrganizationActionBatches200ResponseInnerActionsInner) GetOperationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *GetOrganizationActionBatches200ResponseInnerActionsInner) SetOperation(v string) {
	o.Operation = v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *GetOrganizationActionBatches200ResponseInnerActionsInner) GetBody() map[string]interface{} {
	if o == nil || IsNil(o.Body) {
		var ret map[string]interface{}
		return ret
	}
	return o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationActionBatches200ResponseInnerActionsInner) GetBodyOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Body) {
		return map[string]interface{}{}, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *GetOrganizationActionBatches200ResponseInnerActionsInner) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given map[string]interface{} and assigns it to the Body field.
func (o *GetOrganizationActionBatches200ResponseInnerActionsInner) SetBody(v map[string]interface{}) {
	o.Body = v
}

func (o GetOrganizationActionBatches200ResponseInnerActionsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationActionBatches200ResponseInnerActionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resource"] = o.Resource
	toSerialize["operation"] = o.Operation
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	return toSerialize, nil
}

type NullableGetOrganizationActionBatches200ResponseInnerActionsInner struct {
	value *GetOrganizationActionBatches200ResponseInnerActionsInner
	isSet bool
}

func (v NullableGetOrganizationActionBatches200ResponseInnerActionsInner) Get() *GetOrganizationActionBatches200ResponseInnerActionsInner {
	return v.value
}

func (v *NullableGetOrganizationActionBatches200ResponseInnerActionsInner) Set(val *GetOrganizationActionBatches200ResponseInnerActionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationActionBatches200ResponseInnerActionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationActionBatches200ResponseInnerActionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationActionBatches200ResponseInnerActionsInner(val *GetOrganizationActionBatches200ResponseInnerActionsInner) *NullableGetOrganizationActionBatches200ResponseInnerActionsInner {
	return &NullableGetOrganizationActionBatches200ResponseInnerActionsInner{value: val, isSet: true}
}

func (v NullableGetOrganizationActionBatches200ResponseInnerActionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationActionBatches200ResponseInnerActionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


