/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// GetOrganizationLicensingCotermLicenses200ResponseInner struct for GetOrganizationLicensingCotermLicenses200ResponseInner
type GetOrganizationLicensingCotermLicenses200ResponseInner struct {
	// The key of the license
	Key *string `json:"key,omitempty"`
	// The ID of the organization that the license is claimed in
	OrganizationId *string `json:"organizationId,omitempty"`
	// The duration (term length) of the license, measured in days
	Duration *int32 `json:"duration,omitempty"`
	// The operation mode of the license when it was claimed
	Mode *string `json:"mode,omitempty"`
	// When the license's term began (approximately the date when the license was created)
	StartedAt *time.Time `json:"startedAt,omitempty"`
	// When the license was claimed into the organization
	ClaimedAt *time.Time `json:"claimedAt,omitempty"`
	// Flag to indicated that the license is invalidated
	Invalidated *bool `json:"invalidated,omitempty"`
	// When the license was invalidated. Will be null for active licenses
	InvalidatedAt *time.Time `json:"invalidatedAt,omitempty"`
	// Flag to indicate if the license is expired
	Expired *bool `json:"expired,omitempty"`
	// The editions of the license for each relevant product type
	Editions []GetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner `json:"editions,omitempty"`
	// The counts of the license by model type
	Counts []GetOrganizationLicensingCotermLicenses200ResponseInnerCountsInner `json:"counts,omitempty"`
}

// NewGetOrganizationLicensingCotermLicenses200ResponseInner instantiates a new GetOrganizationLicensingCotermLicenses200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationLicensingCotermLicenses200ResponseInner() *GetOrganizationLicensingCotermLicenses200ResponseInner {
	this := GetOrganizationLicensingCotermLicenses200ResponseInner{}
	return &this
}

// NewGetOrganizationLicensingCotermLicenses200ResponseInnerWithDefaults instantiates a new GetOrganizationLicensingCotermLicenses200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationLicensingCotermLicenses200ResponseInnerWithDefaults() *GetOrganizationLicensingCotermLicenses200ResponseInner {
	this := GetOrganizationLicensingCotermLicenses200ResponseInner{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetKey() string {
	if o == nil || isNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetKeyOk() (*string, bool) {
	if o == nil || isNil(o.Key) {
    return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) HasKey() bool {
	if o != nil && !isNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) SetKey(v string) {
	o.Key = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetOrganizationId() string {
	if o == nil || isNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetOrganizationIdOk() (*string, bool) {
	if o == nil || isNil(o.OrganizationId) {
    return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) HasOrganizationId() bool {
	if o != nil && !isNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetDuration() int32 {
	if o == nil || isNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetDurationOk() (*int32, bool) {
	if o == nil || isNil(o.Duration) {
    return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) HasDuration() bool {
	if o != nil && !isNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) SetDuration(v int32) {
	o.Duration = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetMode() string {
	if o == nil || isNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetModeOk() (*string, bool) {
	if o == nil || isNil(o.Mode) {
    return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) HasMode() bool {
	if o != nil && !isNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) SetMode(v string) {
	o.Mode = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetStartedAt() time.Time {
	if o == nil || isNil(o.StartedAt) {
		var ret time.Time
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetStartedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartedAt) {
    return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) HasStartedAt() bool {
	if o != nil && !isNil(o.StartedAt) {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given time.Time and assigns it to the StartedAt field.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) SetStartedAt(v time.Time) {
	o.StartedAt = &v
}

// GetClaimedAt returns the ClaimedAt field value if set, zero value otherwise.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetClaimedAt() time.Time {
	if o == nil || isNil(o.ClaimedAt) {
		var ret time.Time
		return ret
	}
	return *o.ClaimedAt
}

// GetClaimedAtOk returns a tuple with the ClaimedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetClaimedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.ClaimedAt) {
    return nil, false
	}
	return o.ClaimedAt, true
}

// HasClaimedAt returns a boolean if a field has been set.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) HasClaimedAt() bool {
	if o != nil && !isNil(o.ClaimedAt) {
		return true
	}

	return false
}

// SetClaimedAt gets a reference to the given time.Time and assigns it to the ClaimedAt field.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) SetClaimedAt(v time.Time) {
	o.ClaimedAt = &v
}

// GetInvalidated returns the Invalidated field value if set, zero value otherwise.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetInvalidated() bool {
	if o == nil || isNil(o.Invalidated) {
		var ret bool
		return ret
	}
	return *o.Invalidated
}

// GetInvalidatedOk returns a tuple with the Invalidated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetInvalidatedOk() (*bool, bool) {
	if o == nil || isNil(o.Invalidated) {
    return nil, false
	}
	return o.Invalidated, true
}

// HasInvalidated returns a boolean if a field has been set.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) HasInvalidated() bool {
	if o != nil && !isNil(o.Invalidated) {
		return true
	}

	return false
}

// SetInvalidated gets a reference to the given bool and assigns it to the Invalidated field.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) SetInvalidated(v bool) {
	o.Invalidated = &v
}

// GetInvalidatedAt returns the InvalidatedAt field value if set, zero value otherwise.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetInvalidatedAt() time.Time {
	if o == nil || isNil(o.InvalidatedAt) {
		var ret time.Time
		return ret
	}
	return *o.InvalidatedAt
}

// GetInvalidatedAtOk returns a tuple with the InvalidatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetInvalidatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.InvalidatedAt) {
    return nil, false
	}
	return o.InvalidatedAt, true
}

// HasInvalidatedAt returns a boolean if a field has been set.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) HasInvalidatedAt() bool {
	if o != nil && !isNil(o.InvalidatedAt) {
		return true
	}

	return false
}

// SetInvalidatedAt gets a reference to the given time.Time and assigns it to the InvalidatedAt field.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) SetInvalidatedAt(v time.Time) {
	o.InvalidatedAt = &v
}

// GetExpired returns the Expired field value if set, zero value otherwise.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetExpired() bool {
	if o == nil || isNil(o.Expired) {
		var ret bool
		return ret
	}
	return *o.Expired
}

// GetExpiredOk returns a tuple with the Expired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetExpiredOk() (*bool, bool) {
	if o == nil || isNil(o.Expired) {
    return nil, false
	}
	return o.Expired, true
}

// HasExpired returns a boolean if a field has been set.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) HasExpired() bool {
	if o != nil && !isNil(o.Expired) {
		return true
	}

	return false
}

// SetExpired gets a reference to the given bool and assigns it to the Expired field.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) SetExpired(v bool) {
	o.Expired = &v
}

// GetEditions returns the Editions field value if set, zero value otherwise.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetEditions() []GetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner {
	if o == nil || isNil(o.Editions) {
		var ret []GetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner
		return ret
	}
	return o.Editions
}

// GetEditionsOk returns a tuple with the Editions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetEditionsOk() ([]GetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner, bool) {
	if o == nil || isNil(o.Editions) {
    return nil, false
	}
	return o.Editions, true
}

// HasEditions returns a boolean if a field has been set.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) HasEditions() bool {
	if o != nil && !isNil(o.Editions) {
		return true
	}

	return false
}

// SetEditions gets a reference to the given []GetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner and assigns it to the Editions field.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) SetEditions(v []GetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner) {
	o.Editions = v
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetCounts() []GetOrganizationLicensingCotermLicenses200ResponseInnerCountsInner {
	if o == nil || isNil(o.Counts) {
		var ret []GetOrganizationLicensingCotermLicenses200ResponseInnerCountsInner
		return ret
	}
	return o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetCountsOk() ([]GetOrganizationLicensingCotermLicenses200ResponseInnerCountsInner, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given []GetOrganizationLicensingCotermLicenses200ResponseInnerCountsInner and assigns it to the Counts field.
func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) SetCounts(v []GetOrganizationLicensingCotermLicenses200ResponseInnerCountsInner) {
	o.Counts = v
}

func (o GetOrganizationLicensingCotermLicenses200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !isNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if !isNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !isNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !isNil(o.StartedAt) {
		toSerialize["startedAt"] = o.StartedAt
	}
	if !isNil(o.ClaimedAt) {
		toSerialize["claimedAt"] = o.ClaimedAt
	}
	if !isNil(o.Invalidated) {
		toSerialize["invalidated"] = o.Invalidated
	}
	if !isNil(o.InvalidatedAt) {
		toSerialize["invalidatedAt"] = o.InvalidatedAt
	}
	if !isNil(o.Expired) {
		toSerialize["expired"] = o.Expired
	}
	if !isNil(o.Editions) {
		toSerialize["editions"] = o.Editions
	}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableGetOrganizationLicensingCotermLicenses200ResponseInner struct {
	value *GetOrganizationLicensingCotermLicenses200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationLicensingCotermLicenses200ResponseInner) Get() *GetOrganizationLicensingCotermLicenses200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationLicensingCotermLicenses200ResponseInner) Set(val *GetOrganizationLicensingCotermLicenses200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationLicensingCotermLicenses200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationLicensingCotermLicenses200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationLicensingCotermLicenses200ResponseInner(val *GetOrganizationLicensingCotermLicenses200ResponseInner) *NullableGetOrganizationLicensingCotermLicenses200ResponseInner {
	return &NullableGetOrganizationLicensingCotermLicenses200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationLicensingCotermLicenses200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationLicensingCotermLicenses200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

