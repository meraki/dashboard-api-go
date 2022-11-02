/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// UpdateNetworkMqttBrokerRequest struct for UpdateNetworkMqttBrokerRequest
type UpdateNetworkMqttBrokerRequest struct {
	// Name of the MQTT broker.
	Name *string `json:"name,omitempty"`
	// Host name/IP address where the MQTT broker runs.
	Host *string `json:"host,omitempty"`
	// Host port though which the MQTT broker can be reached.
	Port *int32 `json:"port,omitempty"`
	Security *CreateNetworkMqttBrokerRequestSecurity `json:"security,omitempty"`
	// Authentication settings of the MQTT broker
	Authentication map[string]interface{} `json:"authentication,omitempty"`
}

// NewUpdateNetworkMqttBrokerRequest instantiates a new UpdateNetworkMqttBrokerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkMqttBrokerRequest() *UpdateNetworkMqttBrokerRequest {
	this := UpdateNetworkMqttBrokerRequest{}
	return &this
}

// NewUpdateNetworkMqttBrokerRequestWithDefaults instantiates a new UpdateNetworkMqttBrokerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkMqttBrokerRequestWithDefaults() *UpdateNetworkMqttBrokerRequest {
	this := UpdateNetworkMqttBrokerRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateNetworkMqttBrokerRequest) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkMqttBrokerRequest) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateNetworkMqttBrokerRequest) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateNetworkMqttBrokerRequest) SetName(v string) {
	o.Name = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *UpdateNetworkMqttBrokerRequest) GetHost() string {
	if o == nil || isNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkMqttBrokerRequest) GetHostOk() (*string, bool) {
	if o == nil || isNil(o.Host) {
    return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *UpdateNetworkMqttBrokerRequest) HasHost() bool {
	if o != nil && !isNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *UpdateNetworkMqttBrokerRequest) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *UpdateNetworkMqttBrokerRequest) GetPort() int32 {
	if o == nil || isNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkMqttBrokerRequest) GetPortOk() (*int32, bool) {
	if o == nil || isNil(o.Port) {
    return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *UpdateNetworkMqttBrokerRequest) HasPort() bool {
	if o != nil && !isNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *UpdateNetworkMqttBrokerRequest) SetPort(v int32) {
	o.Port = &v
}

// GetSecurity returns the Security field value if set, zero value otherwise.
func (o *UpdateNetworkMqttBrokerRequest) GetSecurity() CreateNetworkMqttBrokerRequestSecurity {
	if o == nil || isNil(o.Security) {
		var ret CreateNetworkMqttBrokerRequestSecurity
		return ret
	}
	return *o.Security
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkMqttBrokerRequest) GetSecurityOk() (*CreateNetworkMqttBrokerRequestSecurity, bool) {
	if o == nil || isNil(o.Security) {
    return nil, false
	}
	return o.Security, true
}

// HasSecurity returns a boolean if a field has been set.
func (o *UpdateNetworkMqttBrokerRequest) HasSecurity() bool {
	if o != nil && !isNil(o.Security) {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given CreateNetworkMqttBrokerRequestSecurity and assigns it to the Security field.
func (o *UpdateNetworkMqttBrokerRequest) SetSecurity(v CreateNetworkMqttBrokerRequestSecurity) {
	o.Security = &v
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *UpdateNetworkMqttBrokerRequest) GetAuthentication() map[string]interface{} {
	if o == nil || isNil(o.Authentication) {
		var ret map[string]interface{}
		return ret
	}
	return o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkMqttBrokerRequest) GetAuthenticationOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Authentication) {
    return map[string]interface{}{}, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *UpdateNetworkMqttBrokerRequest) HasAuthentication() bool {
	if o != nil && !isNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given map[string]interface{} and assigns it to the Authentication field.
func (o *UpdateNetworkMqttBrokerRequest) SetAuthentication(v map[string]interface{}) {
	o.Authentication = v
}

func (o UpdateNetworkMqttBrokerRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !isNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !isNil(o.Security) {
		toSerialize["security"] = o.Security
	}
	if !isNil(o.Authentication) {
		toSerialize["authentication"] = o.Authentication
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkMqttBrokerRequest struct {
	value *UpdateNetworkMqttBrokerRequest
	isSet bool
}

func (v NullableUpdateNetworkMqttBrokerRequest) Get() *UpdateNetworkMqttBrokerRequest {
	return v.value
}

func (v *NullableUpdateNetworkMqttBrokerRequest) Set(val *UpdateNetworkMqttBrokerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkMqttBrokerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkMqttBrokerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkMqttBrokerRequest(val *UpdateNetworkMqttBrokerRequest) *NullableUpdateNetworkMqttBrokerRequest {
	return &NullableUpdateNetworkMqttBrokerRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkMqttBrokerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkMqttBrokerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


