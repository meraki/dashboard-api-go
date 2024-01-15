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

// checks if the GetNetworkWirelessBilling200ResponsePlansInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkWirelessBilling200ResponsePlansInner{}

// GetNetworkWirelessBilling200ResponsePlansInner struct for GetNetworkWirelessBilling200ResponsePlansInner
type GetNetworkWirelessBilling200ResponsePlansInner struct {
	// The id of the pricing plan to update.
	Id *string `json:"id,omitempty"`
	// The price of the billing plan.
	Price *float32 `json:"price,omitempty"`
	BandwidthLimits *GetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimits `json:"bandwidthLimits,omitempty"`
	// The time limit of the pricing plan in minutes.
	TimeLimit *string `json:"timeLimit,omitempty"`
}

// NewGetNetworkWirelessBilling200ResponsePlansInner instantiates a new GetNetworkWirelessBilling200ResponsePlansInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkWirelessBilling200ResponsePlansInner() *GetNetworkWirelessBilling200ResponsePlansInner {
	this := GetNetworkWirelessBilling200ResponsePlansInner{}
	return &this
}

// NewGetNetworkWirelessBilling200ResponsePlansInnerWithDefaults instantiates a new GetNetworkWirelessBilling200ResponsePlansInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkWirelessBilling200ResponsePlansInnerWithDefaults() *GetNetworkWirelessBilling200ResponsePlansInner {
	this := GetNetworkWirelessBilling200ResponsePlansInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNetworkWirelessBilling200ResponsePlansInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessBilling200ResponsePlansInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetNetworkWirelessBilling200ResponsePlansInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetNetworkWirelessBilling200ResponsePlansInner) SetId(v string) {
	o.Id = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *GetNetworkWirelessBilling200ResponsePlansInner) GetPrice() float32 {
	if o == nil || IsNil(o.Price) {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessBilling200ResponsePlansInner) GetPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *GetNetworkWirelessBilling200ResponsePlansInner) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *GetNetworkWirelessBilling200ResponsePlansInner) SetPrice(v float32) {
	o.Price = &v
}

// GetBandwidthLimits returns the BandwidthLimits field value if set, zero value otherwise.
func (o *GetNetworkWirelessBilling200ResponsePlansInner) GetBandwidthLimits() GetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimits {
	if o == nil || IsNil(o.BandwidthLimits) {
		var ret GetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimits
		return ret
	}
	return *o.BandwidthLimits
}

// GetBandwidthLimitsOk returns a tuple with the BandwidthLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessBilling200ResponsePlansInner) GetBandwidthLimitsOk() (*GetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimits, bool) {
	if o == nil || IsNil(o.BandwidthLimits) {
		return nil, false
	}
	return o.BandwidthLimits, true
}

// HasBandwidthLimits returns a boolean if a field has been set.
func (o *GetNetworkWirelessBilling200ResponsePlansInner) HasBandwidthLimits() bool {
	if o != nil && !IsNil(o.BandwidthLimits) {
		return true
	}

	return false
}

// SetBandwidthLimits gets a reference to the given GetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimits and assigns it to the BandwidthLimits field.
func (o *GetNetworkWirelessBilling200ResponsePlansInner) SetBandwidthLimits(v GetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimits) {
	o.BandwidthLimits = &v
}

// GetTimeLimit returns the TimeLimit field value if set, zero value otherwise.
func (o *GetNetworkWirelessBilling200ResponsePlansInner) GetTimeLimit() string {
	if o == nil || IsNil(o.TimeLimit) {
		var ret string
		return ret
	}
	return *o.TimeLimit
}

// GetTimeLimitOk returns a tuple with the TimeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessBilling200ResponsePlansInner) GetTimeLimitOk() (*string, bool) {
	if o == nil || IsNil(o.TimeLimit) {
		return nil, false
	}
	return o.TimeLimit, true
}

// HasTimeLimit returns a boolean if a field has been set.
func (o *GetNetworkWirelessBilling200ResponsePlansInner) HasTimeLimit() bool {
	if o != nil && !IsNil(o.TimeLimit) {
		return true
	}

	return false
}

// SetTimeLimit gets a reference to the given string and assigns it to the TimeLimit field.
func (o *GetNetworkWirelessBilling200ResponsePlansInner) SetTimeLimit(v string) {
	o.TimeLimit = &v
}

func (o GetNetworkWirelessBilling200ResponsePlansInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkWirelessBilling200ResponsePlansInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.BandwidthLimits) {
		toSerialize["bandwidthLimits"] = o.BandwidthLimits
	}
	if !IsNil(o.TimeLimit) {
		toSerialize["timeLimit"] = o.TimeLimit
	}
	return toSerialize, nil
}

type NullableGetNetworkWirelessBilling200ResponsePlansInner struct {
	value *GetNetworkWirelessBilling200ResponsePlansInner
	isSet bool
}

func (v NullableGetNetworkWirelessBilling200ResponsePlansInner) Get() *GetNetworkWirelessBilling200ResponsePlansInner {
	return v.value
}

func (v *NullableGetNetworkWirelessBilling200ResponsePlansInner) Set(val *GetNetworkWirelessBilling200ResponsePlansInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWirelessBilling200ResponsePlansInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWirelessBilling200ResponsePlansInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWirelessBilling200ResponsePlansInner(val *GetNetworkWirelessBilling200ResponsePlansInner) *NullableGetNetworkWirelessBilling200ResponsePlansInner {
	return &NullableGetNetworkWirelessBilling200ResponsePlansInner{value: val, isSet: true}
}

func (v NullableGetNetworkWirelessBilling200ResponsePlansInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWirelessBilling200ResponsePlansInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


