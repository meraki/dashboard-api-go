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

// checks if the GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner{}

// GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner struct for GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner
type GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner struct {
	//     A list of objects describing the definitions of your traffic shaping rule. At least one definition is required. 
	Definitions []UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner `json:"definitions"`
	PerClientBandwidthLimits *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits `json:"perClientBandwidthLimits,omitempty"`
	//     The DSCP tag applied by your rule. null means 'Do not change DSCP tag'.     For a list of possible tag values, use the trafficShaping/dscpTaggingOptions endpoint. 
	DscpTagValue *int32 `json:"dscpTagValue,omitempty"`
	//     The PCP tag applied by your rule. Can be 0 (lowest priority) through 7 (highest priority).     null means 'Do not set PCP tag'. 
	PcpTagValue *int32 `json:"pcpTagValue,omitempty"`
}

// NewGetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner instantiates a new GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner(definitions []UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner) *GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner {
	this := GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner{}
	this.Definitions = definitions
	return &this
}

// NewGetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInnerWithDefaults instantiates a new GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInnerWithDefaults() *GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner {
	this := GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner{}
	return &this
}

// GetDefinitions returns the Definitions field value
func (o *GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner) GetDefinitions() []UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner {
	if o == nil {
		var ret []UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner
		return ret
	}

	return o.Definitions
}

// GetDefinitionsOk returns a tuple with the Definitions field value
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner) GetDefinitionsOk() ([]UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Definitions, true
}

// SetDefinitions sets field value
func (o *GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner) SetDefinitions(v []UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner) {
	o.Definitions = v
}

// GetPerClientBandwidthLimits returns the PerClientBandwidthLimits field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner) GetPerClientBandwidthLimits() UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits {
	if o == nil || IsNil(o.PerClientBandwidthLimits) {
		var ret UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits
		return ret
	}
	return *o.PerClientBandwidthLimits
}

// GetPerClientBandwidthLimitsOk returns a tuple with the PerClientBandwidthLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner) GetPerClientBandwidthLimitsOk() (*UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits, bool) {
	if o == nil || IsNil(o.PerClientBandwidthLimits) {
		return nil, false
	}
	return o.PerClientBandwidthLimits, true
}

// HasPerClientBandwidthLimits returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner) HasPerClientBandwidthLimits() bool {
	if o != nil && !IsNil(o.PerClientBandwidthLimits) {
		return true
	}

	return false
}

// SetPerClientBandwidthLimits gets a reference to the given UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits and assigns it to the PerClientBandwidthLimits field.
func (o *GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner) SetPerClientBandwidthLimits(v UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits) {
	o.PerClientBandwidthLimits = &v
}

// GetDscpTagValue returns the DscpTagValue field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner) GetDscpTagValue() int32 {
	if o == nil || IsNil(o.DscpTagValue) {
		var ret int32
		return ret
	}
	return *o.DscpTagValue
}

// GetDscpTagValueOk returns a tuple with the DscpTagValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner) GetDscpTagValueOk() (*int32, bool) {
	if o == nil || IsNil(o.DscpTagValue) {
		return nil, false
	}
	return o.DscpTagValue, true
}

// HasDscpTagValue returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner) HasDscpTagValue() bool {
	if o != nil && !IsNil(o.DscpTagValue) {
		return true
	}

	return false
}

// SetDscpTagValue gets a reference to the given int32 and assigns it to the DscpTagValue field.
func (o *GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner) SetDscpTagValue(v int32) {
	o.DscpTagValue = &v
}

// GetPcpTagValue returns the PcpTagValue field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner) GetPcpTagValue() int32 {
	if o == nil || IsNil(o.PcpTagValue) {
		var ret int32
		return ret
	}
	return *o.PcpTagValue
}

// GetPcpTagValueOk returns a tuple with the PcpTagValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner) GetPcpTagValueOk() (*int32, bool) {
	if o == nil || IsNil(o.PcpTagValue) {
		return nil, false
	}
	return o.PcpTagValue, true
}

// HasPcpTagValue returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner) HasPcpTagValue() bool {
	if o != nil && !IsNil(o.PcpTagValue) {
		return true
	}

	return false
}

// SetPcpTagValue gets a reference to the given int32 and assigns it to the PcpTagValue field.
func (o *GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner) SetPcpTagValue(v int32) {
	o.PcpTagValue = &v
}

func (o GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["definitions"] = o.Definitions
	if !IsNil(o.PerClientBandwidthLimits) {
		toSerialize["perClientBandwidthLimits"] = o.PerClientBandwidthLimits
	}
	if !IsNil(o.DscpTagValue) {
		toSerialize["dscpTagValue"] = o.DscpTagValue
	}
	if !IsNil(o.PcpTagValue) {
		toSerialize["pcpTagValue"] = o.PcpTagValue
	}
	return toSerialize, nil
}

type NullableGetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner struct {
	value *GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner
	isSet bool
}

func (v NullableGetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner) Get() *GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner {
	return v.value
}

func (v *NullableGetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner) Set(val *GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner(val *GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner) *NullableGetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner {
	return &NullableGetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner{value: val, isSet: true}
}

func (v NullableGetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


