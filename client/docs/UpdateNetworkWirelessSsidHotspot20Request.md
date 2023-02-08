# UpdateNetworkWirelessSsidHotspot20Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether or not Hotspot 2.0 for this SSID is enabled | [optional] 
**Operator** | Pointer to [**UpdateNetworkWirelessSsidHotspot20RequestOperator**](UpdateNetworkWirelessSsidHotspot20RequestOperator.md) |  | [optional] 
**Venue** | Pointer to [**UpdateNetworkWirelessSsidHotspot20RequestVenue**](UpdateNetworkWirelessSsidHotspot20RequestVenue.md) |  | [optional] 
**NetworkAccessType** | Pointer to **string** | The network type of this SSID (&#39;Private network&#39;, &#39;Private network with guest access&#39;, &#39;Chargeable public network&#39;, &#39;Free public network&#39;, &#39;Personal device network&#39;, &#39;Emergency services only network&#39;, &#39;Test or experimental&#39;, &#39;Wildcard&#39;) | [optional] 
**Domains** | Pointer to **[]string** | An array of domain names | [optional] 
**RoamConsortOis** | Pointer to **[]string** | An array of roaming consortium OIs (hexadecimal number 3-5 octets in length) | [optional] 
**MccMncs** | Pointer to [**[]UpdateNetworkWirelessSsidHotspot20RequestMccMncsInner**](UpdateNetworkWirelessSsidHotspot20RequestMccMncsInner.md) | An array of MCC/MNC pairs | [optional] 
**NaiRealms** | Pointer to [**[]UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner**](UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner.md) | An array of NAI realms | [optional] 

## Methods

### NewUpdateNetworkWirelessSsidHotspot20Request

`func NewUpdateNetworkWirelessSsidHotspot20Request() *UpdateNetworkWirelessSsidHotspot20Request`

NewUpdateNetworkWirelessSsidHotspot20Request instantiates a new UpdateNetworkWirelessSsidHotspot20Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkWirelessSsidHotspot20RequestWithDefaults

`func NewUpdateNetworkWirelessSsidHotspot20RequestWithDefaults() *UpdateNetworkWirelessSsidHotspot20Request`

NewUpdateNetworkWirelessSsidHotspot20RequestWithDefaults instantiates a new UpdateNetworkWirelessSsidHotspot20Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *UpdateNetworkWirelessSsidHotspot20Request) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateNetworkWirelessSsidHotspot20Request) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateNetworkWirelessSsidHotspot20Request) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateNetworkWirelessSsidHotspot20Request) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetOperator

`func (o *UpdateNetworkWirelessSsidHotspot20Request) GetOperator() UpdateNetworkWirelessSsidHotspot20RequestOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *UpdateNetworkWirelessSsidHotspot20Request) GetOperatorOk() (*UpdateNetworkWirelessSsidHotspot20RequestOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *UpdateNetworkWirelessSsidHotspot20Request) SetOperator(v UpdateNetworkWirelessSsidHotspot20RequestOperator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *UpdateNetworkWirelessSsidHotspot20Request) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetVenue

`func (o *UpdateNetworkWirelessSsidHotspot20Request) GetVenue() UpdateNetworkWirelessSsidHotspot20RequestVenue`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *UpdateNetworkWirelessSsidHotspot20Request) GetVenueOk() (*UpdateNetworkWirelessSsidHotspot20RequestVenue, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *UpdateNetworkWirelessSsidHotspot20Request) SetVenue(v UpdateNetworkWirelessSsidHotspot20RequestVenue)`

SetVenue sets Venue field to given value.

### HasVenue

`func (o *UpdateNetworkWirelessSsidHotspot20Request) HasVenue() bool`

HasVenue returns a boolean if a field has been set.

### GetNetworkAccessType

`func (o *UpdateNetworkWirelessSsidHotspot20Request) GetNetworkAccessType() string`

GetNetworkAccessType returns the NetworkAccessType field if non-nil, zero value otherwise.

### GetNetworkAccessTypeOk

`func (o *UpdateNetworkWirelessSsidHotspot20Request) GetNetworkAccessTypeOk() (*string, bool)`

GetNetworkAccessTypeOk returns a tuple with the NetworkAccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAccessType

`func (o *UpdateNetworkWirelessSsidHotspot20Request) SetNetworkAccessType(v string)`

SetNetworkAccessType sets NetworkAccessType field to given value.

### HasNetworkAccessType

`func (o *UpdateNetworkWirelessSsidHotspot20Request) HasNetworkAccessType() bool`

HasNetworkAccessType returns a boolean if a field has been set.

### GetDomains

`func (o *UpdateNetworkWirelessSsidHotspot20Request) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *UpdateNetworkWirelessSsidHotspot20Request) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *UpdateNetworkWirelessSsidHotspot20Request) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *UpdateNetworkWirelessSsidHotspot20Request) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetRoamConsortOis

`func (o *UpdateNetworkWirelessSsidHotspot20Request) GetRoamConsortOis() []string`

GetRoamConsortOis returns the RoamConsortOis field if non-nil, zero value otherwise.

### GetRoamConsortOisOk

`func (o *UpdateNetworkWirelessSsidHotspot20Request) GetRoamConsortOisOk() (*[]string, bool)`

GetRoamConsortOisOk returns a tuple with the RoamConsortOis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoamConsortOis

`func (o *UpdateNetworkWirelessSsidHotspot20Request) SetRoamConsortOis(v []string)`

SetRoamConsortOis sets RoamConsortOis field to given value.

### HasRoamConsortOis

`func (o *UpdateNetworkWirelessSsidHotspot20Request) HasRoamConsortOis() bool`

HasRoamConsortOis returns a boolean if a field has been set.

### GetMccMncs

`func (o *UpdateNetworkWirelessSsidHotspot20Request) GetMccMncs() []UpdateNetworkWirelessSsidHotspot20RequestMccMncsInner`

GetMccMncs returns the MccMncs field if non-nil, zero value otherwise.

### GetMccMncsOk

`func (o *UpdateNetworkWirelessSsidHotspot20Request) GetMccMncsOk() (*[]UpdateNetworkWirelessSsidHotspot20RequestMccMncsInner, bool)`

GetMccMncsOk returns a tuple with the MccMncs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMccMncs

`func (o *UpdateNetworkWirelessSsidHotspot20Request) SetMccMncs(v []UpdateNetworkWirelessSsidHotspot20RequestMccMncsInner)`

SetMccMncs sets MccMncs field to given value.

### HasMccMncs

`func (o *UpdateNetworkWirelessSsidHotspot20Request) HasMccMncs() bool`

HasMccMncs returns a boolean if a field has been set.

### GetNaiRealms

`func (o *UpdateNetworkWirelessSsidHotspot20Request) GetNaiRealms() []UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner`

GetNaiRealms returns the NaiRealms field if non-nil, zero value otherwise.

### GetNaiRealmsOk

`func (o *UpdateNetworkWirelessSsidHotspot20Request) GetNaiRealmsOk() (*[]UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner, bool)`

GetNaiRealmsOk returns a tuple with the NaiRealms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaiRealms

`func (o *UpdateNetworkWirelessSsidHotspot20Request) SetNaiRealms(v []UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner)`

SetNaiRealms sets NaiRealms field to given value.

### HasNaiRealms

`func (o *UpdateNetworkWirelessSsidHotspot20Request) HasNaiRealms() bool`

HasNaiRealms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


