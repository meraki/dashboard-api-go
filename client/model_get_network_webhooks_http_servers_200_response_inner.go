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

// checks if the GetNetworkWebhooksHttpServers200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkWebhooksHttpServers200ResponseInner{}

// GetNetworkWebhooksHttpServers200ResponseInner struct for GetNetworkWebhooksHttpServers200ResponseInner
type GetNetworkWebhooksHttpServers200ResponseInner struct {
	// A Base64 encoded ID.
	Id *string `json:"id,omitempty"`
	// A name for easy reference to the HTTP server
	Name *string `json:"name,omitempty"`
	// The URL of the HTTP server.
	Url *string `json:"url,omitempty"`
	// A Meraki network ID.
	NetworkId *string `json:"networkId,omitempty"`
	PayloadTemplate *GetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate `json:"payloadTemplate,omitempty"`
}

// NewGetNetworkWebhooksHttpServers200ResponseInner instantiates a new GetNetworkWebhooksHttpServers200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkWebhooksHttpServers200ResponseInner() *GetNetworkWebhooksHttpServers200ResponseInner {
	this := GetNetworkWebhooksHttpServers200ResponseInner{}
	return &this
}

// NewGetNetworkWebhooksHttpServers200ResponseInnerWithDefaults instantiates a new GetNetworkWebhooksHttpServers200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkWebhooksHttpServers200ResponseInnerWithDefaults() *GetNetworkWebhooksHttpServers200ResponseInner {
	this := GetNetworkWebhooksHttpServers200ResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNetworkWebhooksHttpServers200ResponseInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWebhooksHttpServers200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetNetworkWebhooksHttpServers200ResponseInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetNetworkWebhooksHttpServers200ResponseInner) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkWebhooksHttpServers200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWebhooksHttpServers200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkWebhooksHttpServers200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkWebhooksHttpServers200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GetNetworkWebhooksHttpServers200ResponseInner) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWebhooksHttpServers200ResponseInner) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GetNetworkWebhooksHttpServers200ResponseInner) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GetNetworkWebhooksHttpServers200ResponseInner) SetUrl(v string) {
	o.Url = &v
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *GetNetworkWebhooksHttpServers200ResponseInner) GetNetworkId() string {
	if o == nil || IsNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWebhooksHttpServers200ResponseInner) GetNetworkIdOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkId) {
		return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *GetNetworkWebhooksHttpServers200ResponseInner) HasNetworkId() bool {
	if o != nil && !IsNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *GetNetworkWebhooksHttpServers200ResponseInner) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetPayloadTemplate returns the PayloadTemplate field value if set, zero value otherwise.
func (o *GetNetworkWebhooksHttpServers200ResponseInner) GetPayloadTemplate() GetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate {
	if o == nil || IsNil(o.PayloadTemplate) {
		var ret GetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate
		return ret
	}
	return *o.PayloadTemplate
}

// GetPayloadTemplateOk returns a tuple with the PayloadTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWebhooksHttpServers200ResponseInner) GetPayloadTemplateOk() (*GetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate, bool) {
	if o == nil || IsNil(o.PayloadTemplate) {
		return nil, false
	}
	return o.PayloadTemplate, true
}

// HasPayloadTemplate returns a boolean if a field has been set.
func (o *GetNetworkWebhooksHttpServers200ResponseInner) HasPayloadTemplate() bool {
	if o != nil && !IsNil(o.PayloadTemplate) {
		return true
	}

	return false
}

// SetPayloadTemplate gets a reference to the given GetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate and assigns it to the PayloadTemplate field.
func (o *GetNetworkWebhooksHttpServers200ResponseInner) SetPayloadTemplate(v GetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate) {
	o.PayloadTemplate = &v
}

func (o GetNetworkWebhooksHttpServers200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkWebhooksHttpServers200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !IsNil(o.PayloadTemplate) {
		toSerialize["payloadTemplate"] = o.PayloadTemplate
	}
	return toSerialize, nil
}

type NullableGetNetworkWebhooksHttpServers200ResponseInner struct {
	value *GetNetworkWebhooksHttpServers200ResponseInner
	isSet bool
}

func (v NullableGetNetworkWebhooksHttpServers200ResponseInner) Get() *GetNetworkWebhooksHttpServers200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkWebhooksHttpServers200ResponseInner) Set(val *GetNetworkWebhooksHttpServers200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWebhooksHttpServers200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWebhooksHttpServers200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWebhooksHttpServers200ResponseInner(val *GetNetworkWebhooksHttpServers200ResponseInner) *NullableGetNetworkWebhooksHttpServers200ResponseInner {
	return &NullableGetNetworkWebhooksHttpServers200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkWebhooksHttpServers200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWebhooksHttpServers200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


