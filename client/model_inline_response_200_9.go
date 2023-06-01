/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 22 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.31.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse2009 struct for InlineResponse2009
type InlineResponse2009 struct {
	// Network ID
	Id *string `json:"id,omitempty"`
	// Organization ID
	OrganizationId *string `json:"organizationId,omitempty"`
	// Network name
	Name *string `json:"name,omitempty"`
	// List of the product types that the network supports
	ProductTypes []string `json:"productTypes,omitempty"`
	// Timezone of the network
	TimeZone *string `json:"timeZone,omitempty"`
	// Network tags
	Tags []string `json:"tags,omitempty"`
	// Enrollment string for the network
	EnrollmentString *string `json:"enrollmentString,omitempty"`
	// URL to the network Dashboard UI
	Url *string `json:"url,omitempty"`
	// Notes for the network
	Notes *string `json:"notes,omitempty"`
	// If the network is bound to a config template
	IsBoundToConfigTemplate *bool `json:"isBoundToConfigTemplate,omitempty"`
}

// NewInlineResponse2009 instantiates a new InlineResponse2009 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2009() *InlineResponse2009 {
	this := InlineResponse2009{}
	return &this
}

// NewInlineResponse2009WithDefaults instantiates a new InlineResponse2009 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2009WithDefaults() *InlineResponse2009 {
	this := InlineResponse2009{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse2009) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2009) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse2009) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse2009) SetId(v string) {
	o.Id = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *InlineResponse2009) GetOrganizationId() string {
	if o == nil || isNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2009) GetOrganizationIdOk() (*string, bool) {
	if o == nil || isNil(o.OrganizationId) {
    return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *InlineResponse2009) HasOrganizationId() bool {
	if o != nil && !isNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *InlineResponse2009) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse2009) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2009) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse2009) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse2009) SetName(v string) {
	o.Name = &v
}

// GetProductTypes returns the ProductTypes field value if set, zero value otherwise.
func (o *InlineResponse2009) GetProductTypes() []string {
	if o == nil || isNil(o.ProductTypes) {
		var ret []string
		return ret
	}
	return o.ProductTypes
}

// GetProductTypesOk returns a tuple with the ProductTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2009) GetProductTypesOk() ([]string, bool) {
	if o == nil || isNil(o.ProductTypes) {
    return nil, false
	}
	return o.ProductTypes, true
}

// HasProductTypes returns a boolean if a field has been set.
func (o *InlineResponse2009) HasProductTypes() bool {
	if o != nil && !isNil(o.ProductTypes) {
		return true
	}

	return false
}

// SetProductTypes gets a reference to the given []string and assigns it to the ProductTypes field.
func (o *InlineResponse2009) SetProductTypes(v []string) {
	o.ProductTypes = v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *InlineResponse2009) GetTimeZone() string {
	if o == nil || isNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2009) GetTimeZoneOk() (*string, bool) {
	if o == nil || isNil(o.TimeZone) {
    return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *InlineResponse2009) HasTimeZone() bool {
	if o != nil && !isNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *InlineResponse2009) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineResponse2009) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2009) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineResponse2009) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *InlineResponse2009) SetTags(v []string) {
	o.Tags = v
}

// GetEnrollmentString returns the EnrollmentString field value if set, zero value otherwise.
func (o *InlineResponse2009) GetEnrollmentString() string {
	if o == nil || isNil(o.EnrollmentString) {
		var ret string
		return ret
	}
	return *o.EnrollmentString
}

// GetEnrollmentStringOk returns a tuple with the EnrollmentString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2009) GetEnrollmentStringOk() (*string, bool) {
	if o == nil || isNil(o.EnrollmentString) {
    return nil, false
	}
	return o.EnrollmentString, true
}

// HasEnrollmentString returns a boolean if a field has been set.
func (o *InlineResponse2009) HasEnrollmentString() bool {
	if o != nil && !isNil(o.EnrollmentString) {
		return true
	}

	return false
}

// SetEnrollmentString gets a reference to the given string and assigns it to the EnrollmentString field.
func (o *InlineResponse2009) SetEnrollmentString(v string) {
	o.EnrollmentString = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *InlineResponse2009) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2009) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *InlineResponse2009) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *InlineResponse2009) SetUrl(v string) {
	o.Url = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *InlineResponse2009) GetNotes() string {
	if o == nil || isNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2009) GetNotesOk() (*string, bool) {
	if o == nil || isNil(o.Notes) {
    return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *InlineResponse2009) HasNotes() bool {
	if o != nil && !isNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *InlineResponse2009) SetNotes(v string) {
	o.Notes = &v
}

// GetIsBoundToConfigTemplate returns the IsBoundToConfigTemplate field value if set, zero value otherwise.
func (o *InlineResponse2009) GetIsBoundToConfigTemplate() bool {
	if o == nil || isNil(o.IsBoundToConfigTemplate) {
		var ret bool
		return ret
	}
	return *o.IsBoundToConfigTemplate
}

// GetIsBoundToConfigTemplateOk returns a tuple with the IsBoundToConfigTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2009) GetIsBoundToConfigTemplateOk() (*bool, bool) {
	if o == nil || isNil(o.IsBoundToConfigTemplate) {
    return nil, false
	}
	return o.IsBoundToConfigTemplate, true
}

// HasIsBoundToConfigTemplate returns a boolean if a field has been set.
func (o *InlineResponse2009) HasIsBoundToConfigTemplate() bool {
	if o != nil && !isNil(o.IsBoundToConfigTemplate) {
		return true
	}

	return false
}

// SetIsBoundToConfigTemplate gets a reference to the given bool and assigns it to the IsBoundToConfigTemplate field.
func (o *InlineResponse2009) SetIsBoundToConfigTemplate(v bool) {
	o.IsBoundToConfigTemplate = &v
}

func (o InlineResponse2009) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.ProductTypes) {
		toSerialize["productTypes"] = o.ProductTypes
	}
	if !isNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !isNil(o.EnrollmentString) {
		toSerialize["enrollmentString"] = o.EnrollmentString
	}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !isNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !isNil(o.IsBoundToConfigTemplate) {
		toSerialize["isBoundToConfigTemplate"] = o.IsBoundToConfigTemplate
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2009 struct {
	value *InlineResponse2009
	isSet bool
}

func (v NullableInlineResponse2009) Get() *InlineResponse2009 {
	return v.value
}

func (v *NullableInlineResponse2009) Set(val *InlineResponse2009) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2009) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2009) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2009(val *InlineResponse2009) *NullableInlineResponse2009 {
	return &NullableInlineResponse2009{value: val, isSet: true}
}

func (v NullableInlineResponse2009) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2009) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

