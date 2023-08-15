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

// checks if the GetNetworkSettings200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSettings200Response{}

// GetNetworkSettings200Response struct for GetNetworkSettings200Response
type GetNetworkSettings200Response struct {
	// Enables / disables the local device status pages (<a target='_blank' href='http://my.meraki.com/'>my.meraki.com, </a><a target='_blank' href='http://ap.meraki.com/'>ap.meraki.com, </a><a target='_blank' href='http://switch.meraki.com/'>switch.meraki.com, </a><a target='_blank' href='http://wired.meraki.com/'>wired.meraki.com</a>). Optional (defaults to false)
	LocalStatusPageEnabled *bool `json:"localStatusPageEnabled,omitempty"`
	// Enables / disables access to the device status page (<a target='_blank'>http://[device's LAN IP])</a>. Optional. Can only be set if localStatusPageEnabled is set to true
	RemoteStatusPageEnabled *bool `json:"remoteStatusPageEnabled,omitempty"`
	LocalStatusPage *GetNetworkSettings200ResponseLocalStatusPage `json:"localStatusPage,omitempty"`
	SecurePort *GetNetworkSettings200ResponseSecurePort `json:"securePort,omitempty"`
	Fips *GetNetworkSettings200ResponseFips `json:"fips,omitempty"`
	NamedVlans *GetNetworkSettings200ResponseNamedVlans `json:"namedVlans,omitempty"`
	ClientPrivacy *GetNetworkSettings200ResponseClientPrivacy `json:"clientPrivacy,omitempty"`
}

// NewGetNetworkSettings200Response instantiates a new GetNetworkSettings200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSettings200Response() *GetNetworkSettings200Response {
	this := GetNetworkSettings200Response{}
	return &this
}

// NewGetNetworkSettings200ResponseWithDefaults instantiates a new GetNetworkSettings200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSettings200ResponseWithDefaults() *GetNetworkSettings200Response {
	this := GetNetworkSettings200Response{}
	return &this
}

// GetLocalStatusPageEnabled returns the LocalStatusPageEnabled field value if set, zero value otherwise.
func (o *GetNetworkSettings200Response) GetLocalStatusPageEnabled() bool {
	if o == nil || IsNil(o.LocalStatusPageEnabled) {
		var ret bool
		return ret
	}
	return *o.LocalStatusPageEnabled
}

// GetLocalStatusPageEnabledOk returns a tuple with the LocalStatusPageEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSettings200Response) GetLocalStatusPageEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.LocalStatusPageEnabled) {
		return nil, false
	}
	return o.LocalStatusPageEnabled, true
}

// HasLocalStatusPageEnabled returns a boolean if a field has been set.
func (o *GetNetworkSettings200Response) HasLocalStatusPageEnabled() bool {
	if o != nil && !IsNil(o.LocalStatusPageEnabled) {
		return true
	}

	return false
}

// SetLocalStatusPageEnabled gets a reference to the given bool and assigns it to the LocalStatusPageEnabled field.
func (o *GetNetworkSettings200Response) SetLocalStatusPageEnabled(v bool) {
	o.LocalStatusPageEnabled = &v
}

// GetRemoteStatusPageEnabled returns the RemoteStatusPageEnabled field value if set, zero value otherwise.
func (o *GetNetworkSettings200Response) GetRemoteStatusPageEnabled() bool {
	if o == nil || IsNil(o.RemoteStatusPageEnabled) {
		var ret bool
		return ret
	}
	return *o.RemoteStatusPageEnabled
}

// GetRemoteStatusPageEnabledOk returns a tuple with the RemoteStatusPageEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSettings200Response) GetRemoteStatusPageEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.RemoteStatusPageEnabled) {
		return nil, false
	}
	return o.RemoteStatusPageEnabled, true
}

// HasRemoteStatusPageEnabled returns a boolean if a field has been set.
func (o *GetNetworkSettings200Response) HasRemoteStatusPageEnabled() bool {
	if o != nil && !IsNil(o.RemoteStatusPageEnabled) {
		return true
	}

	return false
}

// SetRemoteStatusPageEnabled gets a reference to the given bool and assigns it to the RemoteStatusPageEnabled field.
func (o *GetNetworkSettings200Response) SetRemoteStatusPageEnabled(v bool) {
	o.RemoteStatusPageEnabled = &v
}

// GetLocalStatusPage returns the LocalStatusPage field value if set, zero value otherwise.
func (o *GetNetworkSettings200Response) GetLocalStatusPage() GetNetworkSettings200ResponseLocalStatusPage {
	if o == nil || IsNil(o.LocalStatusPage) {
		var ret GetNetworkSettings200ResponseLocalStatusPage
		return ret
	}
	return *o.LocalStatusPage
}

// GetLocalStatusPageOk returns a tuple with the LocalStatusPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSettings200Response) GetLocalStatusPageOk() (*GetNetworkSettings200ResponseLocalStatusPage, bool) {
	if o == nil || IsNil(o.LocalStatusPage) {
		return nil, false
	}
	return o.LocalStatusPage, true
}

// HasLocalStatusPage returns a boolean if a field has been set.
func (o *GetNetworkSettings200Response) HasLocalStatusPage() bool {
	if o != nil && !IsNil(o.LocalStatusPage) {
		return true
	}

	return false
}

// SetLocalStatusPage gets a reference to the given GetNetworkSettings200ResponseLocalStatusPage and assigns it to the LocalStatusPage field.
func (o *GetNetworkSettings200Response) SetLocalStatusPage(v GetNetworkSettings200ResponseLocalStatusPage) {
	o.LocalStatusPage = &v
}

// GetSecurePort returns the SecurePort field value if set, zero value otherwise.
func (o *GetNetworkSettings200Response) GetSecurePort() GetNetworkSettings200ResponseSecurePort {
	if o == nil || IsNil(o.SecurePort) {
		var ret GetNetworkSettings200ResponseSecurePort
		return ret
	}
	return *o.SecurePort
}

// GetSecurePortOk returns a tuple with the SecurePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSettings200Response) GetSecurePortOk() (*GetNetworkSettings200ResponseSecurePort, bool) {
	if o == nil || IsNil(o.SecurePort) {
		return nil, false
	}
	return o.SecurePort, true
}

// HasSecurePort returns a boolean if a field has been set.
func (o *GetNetworkSettings200Response) HasSecurePort() bool {
	if o != nil && !IsNil(o.SecurePort) {
		return true
	}

	return false
}

// SetSecurePort gets a reference to the given GetNetworkSettings200ResponseSecurePort and assigns it to the SecurePort field.
func (o *GetNetworkSettings200Response) SetSecurePort(v GetNetworkSettings200ResponseSecurePort) {
	o.SecurePort = &v
}

// GetFips returns the Fips field value if set, zero value otherwise.
func (o *GetNetworkSettings200Response) GetFips() GetNetworkSettings200ResponseFips {
	if o == nil || IsNil(o.Fips) {
		var ret GetNetworkSettings200ResponseFips
		return ret
	}
	return *o.Fips
}

// GetFipsOk returns a tuple with the Fips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSettings200Response) GetFipsOk() (*GetNetworkSettings200ResponseFips, bool) {
	if o == nil || IsNil(o.Fips) {
		return nil, false
	}
	return o.Fips, true
}

// HasFips returns a boolean if a field has been set.
func (o *GetNetworkSettings200Response) HasFips() bool {
	if o != nil && !IsNil(o.Fips) {
		return true
	}

	return false
}

// SetFips gets a reference to the given GetNetworkSettings200ResponseFips and assigns it to the Fips field.
func (o *GetNetworkSettings200Response) SetFips(v GetNetworkSettings200ResponseFips) {
	o.Fips = &v
}

// GetNamedVlans returns the NamedVlans field value if set, zero value otherwise.
func (o *GetNetworkSettings200Response) GetNamedVlans() GetNetworkSettings200ResponseNamedVlans {
	if o == nil || IsNil(o.NamedVlans) {
		var ret GetNetworkSettings200ResponseNamedVlans
		return ret
	}
	return *o.NamedVlans
}

// GetNamedVlansOk returns a tuple with the NamedVlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSettings200Response) GetNamedVlansOk() (*GetNetworkSettings200ResponseNamedVlans, bool) {
	if o == nil || IsNil(o.NamedVlans) {
		return nil, false
	}
	return o.NamedVlans, true
}

// HasNamedVlans returns a boolean if a field has been set.
func (o *GetNetworkSettings200Response) HasNamedVlans() bool {
	if o != nil && !IsNil(o.NamedVlans) {
		return true
	}

	return false
}

// SetNamedVlans gets a reference to the given GetNetworkSettings200ResponseNamedVlans and assigns it to the NamedVlans field.
func (o *GetNetworkSettings200Response) SetNamedVlans(v GetNetworkSettings200ResponseNamedVlans) {
	o.NamedVlans = &v
}

// GetClientPrivacy returns the ClientPrivacy field value if set, zero value otherwise.
func (o *GetNetworkSettings200Response) GetClientPrivacy() GetNetworkSettings200ResponseClientPrivacy {
	if o == nil || IsNil(o.ClientPrivacy) {
		var ret GetNetworkSettings200ResponseClientPrivacy
		return ret
	}
	return *o.ClientPrivacy
}

// GetClientPrivacyOk returns a tuple with the ClientPrivacy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSettings200Response) GetClientPrivacyOk() (*GetNetworkSettings200ResponseClientPrivacy, bool) {
	if o == nil || IsNil(o.ClientPrivacy) {
		return nil, false
	}
	return o.ClientPrivacy, true
}

// HasClientPrivacy returns a boolean if a field has been set.
func (o *GetNetworkSettings200Response) HasClientPrivacy() bool {
	if o != nil && !IsNil(o.ClientPrivacy) {
		return true
	}

	return false
}

// SetClientPrivacy gets a reference to the given GetNetworkSettings200ResponseClientPrivacy and assigns it to the ClientPrivacy field.
func (o *GetNetworkSettings200Response) SetClientPrivacy(v GetNetworkSettings200ResponseClientPrivacy) {
	o.ClientPrivacy = &v
}

func (o GetNetworkSettings200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSettings200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LocalStatusPageEnabled) {
		toSerialize["localStatusPageEnabled"] = o.LocalStatusPageEnabled
	}
	if !IsNil(o.RemoteStatusPageEnabled) {
		toSerialize["remoteStatusPageEnabled"] = o.RemoteStatusPageEnabled
	}
	if !IsNil(o.LocalStatusPage) {
		toSerialize["localStatusPage"] = o.LocalStatusPage
	}
	if !IsNil(o.SecurePort) {
		toSerialize["securePort"] = o.SecurePort
	}
	if !IsNil(o.Fips) {
		toSerialize["fips"] = o.Fips
	}
	if !IsNil(o.NamedVlans) {
		toSerialize["namedVlans"] = o.NamedVlans
	}
	if !IsNil(o.ClientPrivacy) {
		toSerialize["clientPrivacy"] = o.ClientPrivacy
	}
	return toSerialize, nil
}

type NullableGetNetworkSettings200Response struct {
	value *GetNetworkSettings200Response
	isSet bool
}

func (v NullableGetNetworkSettings200Response) Get() *GetNetworkSettings200Response {
	return v.value
}

func (v *NullableGetNetworkSettings200Response) Set(val *GetNetworkSettings200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSettings200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSettings200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSettings200Response(val *GetNetworkSettings200Response) *NullableGetNetworkSettings200Response {
	return &NullableGetNetworkSettings200Response{value: val, isSet: true}
}

func (v NullableGetNetworkSettings200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSettings200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


