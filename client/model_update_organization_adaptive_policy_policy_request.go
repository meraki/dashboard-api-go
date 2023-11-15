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

// checks if the UpdateOrganizationAdaptivePolicyPolicyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOrganizationAdaptivePolicyPolicyRequest{}

// UpdateOrganizationAdaptivePolicyPolicyRequest struct for UpdateOrganizationAdaptivePolicyPolicyRequest
type UpdateOrganizationAdaptivePolicyPolicyRequest struct {
	SourceGroup *CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup `json:"sourceGroup,omitempty"`
	DestinationGroup *CreateOrganizationAdaptivePolicyPolicyRequestDestinationGroup `json:"destinationGroup,omitempty"`
	// An ordered array of adaptive policy ACLs (each requires one unique attribute) that apply to this policy
	Acls []CreateOrganizationAdaptivePolicyPolicyRequestAclsInner `json:"acls,omitempty"`
	// The rule to apply if there is no matching ACL
	LastEntryRule *string `json:"lastEntryRule,omitempty"`
}

// NewUpdateOrganizationAdaptivePolicyPolicyRequest instantiates a new UpdateOrganizationAdaptivePolicyPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrganizationAdaptivePolicyPolicyRequest() *UpdateOrganizationAdaptivePolicyPolicyRequest {
	this := UpdateOrganizationAdaptivePolicyPolicyRequest{}
	return &this
}

// NewUpdateOrganizationAdaptivePolicyPolicyRequestWithDefaults instantiates a new UpdateOrganizationAdaptivePolicyPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrganizationAdaptivePolicyPolicyRequestWithDefaults() *UpdateOrganizationAdaptivePolicyPolicyRequest {
	this := UpdateOrganizationAdaptivePolicyPolicyRequest{}
	return &this
}

// GetSourceGroup returns the SourceGroup field value if set, zero value otherwise.
func (o *UpdateOrganizationAdaptivePolicyPolicyRequest) GetSourceGroup() CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup {
	if o == nil || IsNil(o.SourceGroup) {
		var ret CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup
		return ret
	}
	return *o.SourceGroup
}

// GetSourceGroupOk returns a tuple with the SourceGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationAdaptivePolicyPolicyRequest) GetSourceGroupOk() (*CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup, bool) {
	if o == nil || IsNil(o.SourceGroup) {
		return nil, false
	}
	return o.SourceGroup, true
}

// HasSourceGroup returns a boolean if a field has been set.
func (o *UpdateOrganizationAdaptivePolicyPolicyRequest) HasSourceGroup() bool {
	if o != nil && !IsNil(o.SourceGroup) {
		return true
	}

	return false
}

// SetSourceGroup gets a reference to the given CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup and assigns it to the SourceGroup field.
func (o *UpdateOrganizationAdaptivePolicyPolicyRequest) SetSourceGroup(v CreateOrganizationAdaptivePolicyPolicyRequestSourceGroup) {
	o.SourceGroup = &v
}

// GetDestinationGroup returns the DestinationGroup field value if set, zero value otherwise.
func (o *UpdateOrganizationAdaptivePolicyPolicyRequest) GetDestinationGroup() CreateOrganizationAdaptivePolicyPolicyRequestDestinationGroup {
	if o == nil || IsNil(o.DestinationGroup) {
		var ret CreateOrganizationAdaptivePolicyPolicyRequestDestinationGroup
		return ret
	}
	return *o.DestinationGroup
}

// GetDestinationGroupOk returns a tuple with the DestinationGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationAdaptivePolicyPolicyRequest) GetDestinationGroupOk() (*CreateOrganizationAdaptivePolicyPolicyRequestDestinationGroup, bool) {
	if o == nil || IsNil(o.DestinationGroup) {
		return nil, false
	}
	return o.DestinationGroup, true
}

// HasDestinationGroup returns a boolean if a field has been set.
func (o *UpdateOrganizationAdaptivePolicyPolicyRequest) HasDestinationGroup() bool {
	if o != nil && !IsNil(o.DestinationGroup) {
		return true
	}

	return false
}

// SetDestinationGroup gets a reference to the given CreateOrganizationAdaptivePolicyPolicyRequestDestinationGroup and assigns it to the DestinationGroup field.
func (o *UpdateOrganizationAdaptivePolicyPolicyRequest) SetDestinationGroup(v CreateOrganizationAdaptivePolicyPolicyRequestDestinationGroup) {
	o.DestinationGroup = &v
}

// GetAcls returns the Acls field value if set, zero value otherwise.
func (o *UpdateOrganizationAdaptivePolicyPolicyRequest) GetAcls() []CreateOrganizationAdaptivePolicyPolicyRequestAclsInner {
	if o == nil || IsNil(o.Acls) {
		var ret []CreateOrganizationAdaptivePolicyPolicyRequestAclsInner
		return ret
	}
	return o.Acls
}

// GetAclsOk returns a tuple with the Acls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationAdaptivePolicyPolicyRequest) GetAclsOk() ([]CreateOrganizationAdaptivePolicyPolicyRequestAclsInner, bool) {
	if o == nil || IsNil(o.Acls) {
		return nil, false
	}
	return o.Acls, true
}

// HasAcls returns a boolean if a field has been set.
func (o *UpdateOrganizationAdaptivePolicyPolicyRequest) HasAcls() bool {
	if o != nil && !IsNil(o.Acls) {
		return true
	}

	return false
}

// SetAcls gets a reference to the given []CreateOrganizationAdaptivePolicyPolicyRequestAclsInner and assigns it to the Acls field.
func (o *UpdateOrganizationAdaptivePolicyPolicyRequest) SetAcls(v []CreateOrganizationAdaptivePolicyPolicyRequestAclsInner) {
	o.Acls = v
}

// GetLastEntryRule returns the LastEntryRule field value if set, zero value otherwise.
func (o *UpdateOrganizationAdaptivePolicyPolicyRequest) GetLastEntryRule() string {
	if o == nil || IsNil(o.LastEntryRule) {
		var ret string
		return ret
	}
	return *o.LastEntryRule
}

// GetLastEntryRuleOk returns a tuple with the LastEntryRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationAdaptivePolicyPolicyRequest) GetLastEntryRuleOk() (*string, bool) {
	if o == nil || IsNil(o.LastEntryRule) {
		return nil, false
	}
	return o.LastEntryRule, true
}

// HasLastEntryRule returns a boolean if a field has been set.
func (o *UpdateOrganizationAdaptivePolicyPolicyRequest) HasLastEntryRule() bool {
	if o != nil && !IsNil(o.LastEntryRule) {
		return true
	}

	return false
}

// SetLastEntryRule gets a reference to the given string and assigns it to the LastEntryRule field.
func (o *UpdateOrganizationAdaptivePolicyPolicyRequest) SetLastEntryRule(v string) {
	o.LastEntryRule = &v
}

func (o UpdateOrganizationAdaptivePolicyPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOrganizationAdaptivePolicyPolicyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SourceGroup) {
		toSerialize["sourceGroup"] = o.SourceGroup
	}
	if !IsNil(o.DestinationGroup) {
		toSerialize["destinationGroup"] = o.DestinationGroup
	}
	if !IsNil(o.Acls) {
		toSerialize["acls"] = o.Acls
	}
	if !IsNil(o.LastEntryRule) {
		toSerialize["lastEntryRule"] = o.LastEntryRule
	}
	return toSerialize, nil
}

type NullableUpdateOrganizationAdaptivePolicyPolicyRequest struct {
	value *UpdateOrganizationAdaptivePolicyPolicyRequest
	isSet bool
}

func (v NullableUpdateOrganizationAdaptivePolicyPolicyRequest) Get() *UpdateOrganizationAdaptivePolicyPolicyRequest {
	return v.value
}

func (v *NullableUpdateOrganizationAdaptivePolicyPolicyRequest) Set(val *UpdateOrganizationAdaptivePolicyPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrganizationAdaptivePolicyPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrganizationAdaptivePolicyPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrganizationAdaptivePolicyPolicyRequest(val *UpdateOrganizationAdaptivePolicyPolicyRequest) *NullableUpdateOrganizationAdaptivePolicyPolicyRequest {
	return &NullableUpdateOrganizationAdaptivePolicyPolicyRequest{value: val, isSet: true}
}

func (v NullableUpdateOrganizationAdaptivePolicyPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrganizationAdaptivePolicyPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


