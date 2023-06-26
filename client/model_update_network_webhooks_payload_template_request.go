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

// UpdateNetworkWebhooksPayloadTemplateRequest struct for UpdateNetworkWebhooksPayloadTemplateRequest
type UpdateNetworkWebhooksPayloadTemplateRequest struct {
	// The name of the template
	Name *string `json:"name,omitempty"`
	// The liquid template used for the body of the webhook message.
	Body *string `json:"body,omitempty"`
	// The liquid template used with the webhook headers.
	Headers []CreateNetworkWebhooksPayloadTemplateRequestHeadersInner `json:"headers,omitempty"`
	// A file containing liquid template used for the body of the webhook message.
	BodyFile *string `json:"bodyFile,omitempty"`
	// A file containing the liquid template used with the webhook headers.
	HeadersFile *string `json:"headersFile,omitempty"`
}

// NewUpdateNetworkWebhooksPayloadTemplateRequest instantiates a new UpdateNetworkWebhooksPayloadTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWebhooksPayloadTemplateRequest() *UpdateNetworkWebhooksPayloadTemplateRequest {
	this := UpdateNetworkWebhooksPayloadTemplateRequest{}
	return &this
}

// NewUpdateNetworkWebhooksPayloadTemplateRequestWithDefaults instantiates a new UpdateNetworkWebhooksPayloadTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWebhooksPayloadTemplateRequestWithDefaults() *UpdateNetworkWebhooksPayloadTemplateRequest {
	this := UpdateNetworkWebhooksPayloadTemplateRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateNetworkWebhooksPayloadTemplateRequest) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWebhooksPayloadTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateNetworkWebhooksPayloadTemplateRequest) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateNetworkWebhooksPayloadTemplateRequest) SetName(v string) {
	o.Name = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *UpdateNetworkWebhooksPayloadTemplateRequest) GetBody() string {
	if o == nil || isNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWebhooksPayloadTemplateRequest) GetBodyOk() (*string, bool) {
	if o == nil || isNil(o.Body) {
    return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *UpdateNetworkWebhooksPayloadTemplateRequest) HasBody() bool {
	if o != nil && !isNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *UpdateNetworkWebhooksPayloadTemplateRequest) SetBody(v string) {
	o.Body = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *UpdateNetworkWebhooksPayloadTemplateRequest) GetHeaders() []CreateNetworkWebhooksPayloadTemplateRequestHeadersInner {
	if o == nil || isNil(o.Headers) {
		var ret []CreateNetworkWebhooksPayloadTemplateRequestHeadersInner
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWebhooksPayloadTemplateRequest) GetHeadersOk() ([]CreateNetworkWebhooksPayloadTemplateRequestHeadersInner, bool) {
	if o == nil || isNil(o.Headers) {
    return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *UpdateNetworkWebhooksPayloadTemplateRequest) HasHeaders() bool {
	if o != nil && !isNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []CreateNetworkWebhooksPayloadTemplateRequestHeadersInner and assigns it to the Headers field.
func (o *UpdateNetworkWebhooksPayloadTemplateRequest) SetHeaders(v []CreateNetworkWebhooksPayloadTemplateRequestHeadersInner) {
	o.Headers = v
}

// GetBodyFile returns the BodyFile field value if set, zero value otherwise.
func (o *UpdateNetworkWebhooksPayloadTemplateRequest) GetBodyFile() string {
	if o == nil || isNil(o.BodyFile) {
		var ret string
		return ret
	}
	return *o.BodyFile
}

// GetBodyFileOk returns a tuple with the BodyFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWebhooksPayloadTemplateRequest) GetBodyFileOk() (*string, bool) {
	if o == nil || isNil(o.BodyFile) {
    return nil, false
	}
	return o.BodyFile, true
}

// HasBodyFile returns a boolean if a field has been set.
func (o *UpdateNetworkWebhooksPayloadTemplateRequest) HasBodyFile() bool {
	if o != nil && !isNil(o.BodyFile) {
		return true
	}

	return false
}

// SetBodyFile gets a reference to the given string and assigns it to the BodyFile field.
func (o *UpdateNetworkWebhooksPayloadTemplateRequest) SetBodyFile(v string) {
	o.BodyFile = &v
}

// GetHeadersFile returns the HeadersFile field value if set, zero value otherwise.
func (o *UpdateNetworkWebhooksPayloadTemplateRequest) GetHeadersFile() string {
	if o == nil || isNil(o.HeadersFile) {
		var ret string
		return ret
	}
	return *o.HeadersFile
}

// GetHeadersFileOk returns a tuple with the HeadersFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWebhooksPayloadTemplateRequest) GetHeadersFileOk() (*string, bool) {
	if o == nil || isNil(o.HeadersFile) {
    return nil, false
	}
	return o.HeadersFile, true
}

// HasHeadersFile returns a boolean if a field has been set.
func (o *UpdateNetworkWebhooksPayloadTemplateRequest) HasHeadersFile() bool {
	if o != nil && !isNil(o.HeadersFile) {
		return true
	}

	return false
}

// SetHeadersFile gets a reference to the given string and assigns it to the HeadersFile field.
func (o *UpdateNetworkWebhooksPayloadTemplateRequest) SetHeadersFile(v string) {
	o.HeadersFile = &v
}

func (o UpdateNetworkWebhooksPayloadTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if !isNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	if !isNil(o.BodyFile) {
		toSerialize["bodyFile"] = o.BodyFile
	}
	if !isNil(o.HeadersFile) {
		toSerialize["headersFile"] = o.HeadersFile
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkWebhooksPayloadTemplateRequest struct {
	value *UpdateNetworkWebhooksPayloadTemplateRequest
	isSet bool
}

func (v NullableUpdateNetworkWebhooksPayloadTemplateRequest) Get() *UpdateNetworkWebhooksPayloadTemplateRequest {
	return v.value
}

func (v *NullableUpdateNetworkWebhooksPayloadTemplateRequest) Set(val *UpdateNetworkWebhooksPayloadTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWebhooksPayloadTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWebhooksPayloadTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWebhooksPayloadTemplateRequest(val *UpdateNetworkWebhooksPayloadTemplateRequest) *NullableUpdateNetworkWebhooksPayloadTemplateRequest {
	return &NullableUpdateNetworkWebhooksPayloadTemplateRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkWebhooksPayloadTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWebhooksPayloadTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


