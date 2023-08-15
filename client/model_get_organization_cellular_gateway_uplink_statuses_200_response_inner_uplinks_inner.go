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

// checks if the GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner{}

// GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner struct for GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner
type GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner struct {
	// Uplink interface
	Interface *string `json:"interface,omitempty"`
	// Uplink status
	Status *string `json:"status,omitempty"`
	// Uplink IP
	Ip *string `json:"ip,omitempty"`
	// Network Provider
	Provider *string `json:"provider,omitempty"`
	// Public IP
	PublicIp *string `json:"publicIp,omitempty"`
	// Uplink model
	Model *string `json:"model,omitempty"`
	SignalStat *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat `json:"signalStat,omitempty"`
	// Connection Type
	ConnectionType *string `json:"connectionType,omitempty"`
	// Access Point Name
	Apn *string `json:"apn,omitempty"`
	// Gateway IP
	Gateway *string `json:"gateway,omitempty"`
	// Primary DNS IP
	Dns1 *string `json:"dns1,omitempty"`
	// Secondary DNS IP
	Dns2 *string `json:"dns2,omitempty"`
	// Signal Type
	SignalType *string `json:"signalType,omitempty"`
	// Integrated Circuit Card Identification Number
	Iccid *string `json:"iccid,omitempty"`
}

// NewGetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner instantiates a new GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner() *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner {
	this := GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner{}
	return &this
}

// NewGetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerWithDefaults instantiates a new GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerWithDefaults() *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner {
	this := GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner{}
	return &this
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetInterface() string {
	if o == nil || IsNil(o.Interface) {
		var ret string
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetInterfaceOk() (*string, bool) {
	if o == nil || IsNil(o.Interface) {
		return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) HasInterface() bool {
	if o != nil && !IsNil(o.Interface) {
		return true
	}

	return false
}

// SetInterface gets a reference to the given string and assigns it to the Interface field.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) SetInterface(v string) {
	o.Interface = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) SetStatus(v string) {
	o.Status = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) SetIp(v string) {
	o.Ip = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetProvider() string {
	if o == nil || IsNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetProviderOk() (*string, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) SetProvider(v string) {
	o.Provider = &v
}

// GetPublicIp returns the PublicIp field value if set, zero value otherwise.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetPublicIp() string {
	if o == nil || IsNil(o.PublicIp) {
		var ret string
		return ret
	}
	return *o.PublicIp
}

// GetPublicIpOk returns a tuple with the PublicIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetPublicIpOk() (*string, bool) {
	if o == nil || IsNil(o.PublicIp) {
		return nil, false
	}
	return o.PublicIp, true
}

// HasPublicIp returns a boolean if a field has been set.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) HasPublicIp() bool {
	if o != nil && !IsNil(o.PublicIp) {
		return true
	}

	return false
}

// SetPublicIp gets a reference to the given string and assigns it to the PublicIp field.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) SetPublicIp(v string) {
	o.PublicIp = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) SetModel(v string) {
	o.Model = &v
}

// GetSignalStat returns the SignalStat field value if set, zero value otherwise.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetSignalStat() GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat {
	if o == nil || IsNil(o.SignalStat) {
		var ret GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat
		return ret
	}
	return *o.SignalStat
}

// GetSignalStatOk returns a tuple with the SignalStat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetSignalStatOk() (*GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat, bool) {
	if o == nil || IsNil(o.SignalStat) {
		return nil, false
	}
	return o.SignalStat, true
}

// HasSignalStat returns a boolean if a field has been set.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) HasSignalStat() bool {
	if o != nil && !IsNil(o.SignalStat) {
		return true
	}

	return false
}

// SetSignalStat gets a reference to the given GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat and assigns it to the SignalStat field.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) SetSignalStat(v GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat) {
	o.SignalStat = &v
}

// GetConnectionType returns the ConnectionType field value if set, zero value otherwise.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetConnectionType() string {
	if o == nil || IsNil(o.ConnectionType) {
		var ret string
		return ret
	}
	return *o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetConnectionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionType) {
		return nil, false
	}
	return o.ConnectionType, true
}

// HasConnectionType returns a boolean if a field has been set.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) HasConnectionType() bool {
	if o != nil && !IsNil(o.ConnectionType) {
		return true
	}

	return false
}

// SetConnectionType gets a reference to the given string and assigns it to the ConnectionType field.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) SetConnectionType(v string) {
	o.ConnectionType = &v
}

// GetApn returns the Apn field value if set, zero value otherwise.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetApn() string {
	if o == nil || IsNil(o.Apn) {
		var ret string
		return ret
	}
	return *o.Apn
}

// GetApnOk returns a tuple with the Apn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetApnOk() (*string, bool) {
	if o == nil || IsNil(o.Apn) {
		return nil, false
	}
	return o.Apn, true
}

// HasApn returns a boolean if a field has been set.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) HasApn() bool {
	if o != nil && !IsNil(o.Apn) {
		return true
	}

	return false
}

// SetApn gets a reference to the given string and assigns it to the Apn field.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) SetApn(v string) {
	o.Apn = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetGateway() string {
	if o == nil || IsNil(o.Gateway) {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetGatewayOk() (*string, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) SetGateway(v string) {
	o.Gateway = &v
}

// GetDns1 returns the Dns1 field value if set, zero value otherwise.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetDns1() string {
	if o == nil || IsNil(o.Dns1) {
		var ret string
		return ret
	}
	return *o.Dns1
}

// GetDns1Ok returns a tuple with the Dns1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetDns1Ok() (*string, bool) {
	if o == nil || IsNil(o.Dns1) {
		return nil, false
	}
	return o.Dns1, true
}

// HasDns1 returns a boolean if a field has been set.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) HasDns1() bool {
	if o != nil && !IsNil(o.Dns1) {
		return true
	}

	return false
}

// SetDns1 gets a reference to the given string and assigns it to the Dns1 field.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) SetDns1(v string) {
	o.Dns1 = &v
}

// GetDns2 returns the Dns2 field value if set, zero value otherwise.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetDns2() string {
	if o == nil || IsNil(o.Dns2) {
		var ret string
		return ret
	}
	return *o.Dns2
}

// GetDns2Ok returns a tuple with the Dns2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetDns2Ok() (*string, bool) {
	if o == nil || IsNil(o.Dns2) {
		return nil, false
	}
	return o.Dns2, true
}

// HasDns2 returns a boolean if a field has been set.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) HasDns2() bool {
	if o != nil && !IsNil(o.Dns2) {
		return true
	}

	return false
}

// SetDns2 gets a reference to the given string and assigns it to the Dns2 field.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) SetDns2(v string) {
	o.Dns2 = &v
}

// GetSignalType returns the SignalType field value if set, zero value otherwise.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetSignalType() string {
	if o == nil || IsNil(o.SignalType) {
		var ret string
		return ret
	}
	return *o.SignalType
}

// GetSignalTypeOk returns a tuple with the SignalType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetSignalTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SignalType) {
		return nil, false
	}
	return o.SignalType, true
}

// HasSignalType returns a boolean if a field has been set.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) HasSignalType() bool {
	if o != nil && !IsNil(o.SignalType) {
		return true
	}

	return false
}

// SetSignalType gets a reference to the given string and assigns it to the SignalType field.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) SetSignalType(v string) {
	o.SignalType = &v
}

// GetIccid returns the Iccid field value if set, zero value otherwise.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetIccid() string {
	if o == nil || IsNil(o.Iccid) {
		var ret string
		return ret
	}
	return *o.Iccid
}

// GetIccidOk returns a tuple with the Iccid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetIccidOk() (*string, bool) {
	if o == nil || IsNil(o.Iccid) {
		return nil, false
	}
	return o.Iccid, true
}

// HasIccid returns a boolean if a field has been set.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) HasIccid() bool {
	if o != nil && !IsNil(o.Iccid) {
		return true
	}

	return false
}

// SetIccid gets a reference to the given string and assigns it to the Iccid field.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) SetIccid(v string) {
	o.Iccid = &v
}

func (o GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Interface) {
		toSerialize["interface"] = o.Interface
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !IsNil(o.PublicIp) {
		toSerialize["publicIp"] = o.PublicIp
	}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !IsNil(o.SignalStat) {
		toSerialize["signalStat"] = o.SignalStat
	}
	if !IsNil(o.ConnectionType) {
		toSerialize["connectionType"] = o.ConnectionType
	}
	if !IsNil(o.Apn) {
		toSerialize["apn"] = o.Apn
	}
	if !IsNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	if !IsNil(o.Dns1) {
		toSerialize["dns1"] = o.Dns1
	}
	if !IsNil(o.Dns2) {
		toSerialize["dns2"] = o.Dns2
	}
	if !IsNil(o.SignalType) {
		toSerialize["signalType"] = o.SignalType
	}
	if !IsNil(o.Iccid) {
		toSerialize["iccid"] = o.Iccid
	}
	return toSerialize, nil
}

type NullableGetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner struct {
	value *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner
	isSet bool
}

func (v NullableGetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) Get() *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner {
	return v.value
}

func (v *NullableGetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) Set(val *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner(val *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) *NullableGetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner {
	return &NullableGetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner{value: val, isSet: true}
}

func (v NullableGetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


