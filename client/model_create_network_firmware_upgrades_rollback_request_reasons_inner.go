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

// checks if the CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner{}

// CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner struct for CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner
type CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner struct {
	// Reason for the rollback
	Category string `json:"category"`
	// Additional comment about the rollback
	Comment string `json:"comment"`
}

// NewCreateNetworkFirmwareUpgradesRollbackRequestReasonsInner instantiates a new CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkFirmwareUpgradesRollbackRequestReasonsInner(category string, comment string) *CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner {
	this := CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner{}
	this.Category = category
	this.Comment = comment
	return &this
}

// NewCreateNetworkFirmwareUpgradesRollbackRequestReasonsInnerWithDefaults instantiates a new CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkFirmwareUpgradesRollbackRequestReasonsInnerWithDefaults() *CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner {
	this := CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner{}
	return &this
}

// GetCategory returns the Category field value
func (o *CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner) SetCategory(v string) {
	o.Category = v
}

// GetComment returns the Comment field value
func (o *CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner) GetComment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Comment
}

// GetCommentOk returns a tuple with the Comment field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comment, true
}

// SetComment sets field value
func (o *CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner) SetComment(v string) {
	o.Comment = v
}

func (o CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["category"] = o.Category
	toSerialize["comment"] = o.Comment
	return toSerialize, nil
}

type NullableCreateNetworkFirmwareUpgradesRollbackRequestReasonsInner struct {
	value *CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner
	isSet bool
}

func (v NullableCreateNetworkFirmwareUpgradesRollbackRequestReasonsInner) Get() *CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner {
	return v.value
}

func (v *NullableCreateNetworkFirmwareUpgradesRollbackRequestReasonsInner) Set(val *CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkFirmwareUpgradesRollbackRequestReasonsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkFirmwareUpgradesRollbackRequestReasonsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkFirmwareUpgradesRollbackRequestReasonsInner(val *CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner) *NullableCreateNetworkFirmwareUpgradesRollbackRequestReasonsInner {
	return &NullableCreateNetworkFirmwareUpgradesRollbackRequestReasonsInner{value: val, isSet: true}
}

func (v NullableCreateNetworkFirmwareUpgradesRollbackRequestReasonsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkFirmwareUpgradesRollbackRequestReasonsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


