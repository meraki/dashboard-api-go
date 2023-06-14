# GetNetworkWirelessFailedConnections200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SsidNumber** | Pointer to **int32** | SSID Number | [optional] 
**Vlan** | Pointer to **int32** | LAN | [optional] 
**ClientMac** | Pointer to **string** | Client Mac | [optional] 
**Serial** | Pointer to **string** | Serial Number | [optional] 
**FailureStep** | Pointer to **string** | The failed onboarding step. One of: assoc, auth, dhcp, dns. | [optional] 
**Type** | Pointer to **string** | The failure type in the onboarding step | [optional] 
**Ts** | Pointer to **time.Time** | The timestamp when the client mac failed | [optional] 

## Methods

### NewGetNetworkWirelessFailedConnections200ResponseInner

`func NewGetNetworkWirelessFailedConnections200ResponseInner() *GetNetworkWirelessFailedConnections200ResponseInner`

NewGetNetworkWirelessFailedConnections200ResponseInner instantiates a new GetNetworkWirelessFailedConnections200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkWirelessFailedConnections200ResponseInnerWithDefaults

`func NewGetNetworkWirelessFailedConnections200ResponseInnerWithDefaults() *GetNetworkWirelessFailedConnections200ResponseInner`

NewGetNetworkWirelessFailedConnections200ResponseInnerWithDefaults instantiates a new GetNetworkWirelessFailedConnections200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsidNumber

`func (o *GetNetworkWirelessFailedConnections200ResponseInner) GetSsidNumber() int32`

GetSsidNumber returns the SsidNumber field if non-nil, zero value otherwise.

### GetSsidNumberOk

`func (o *GetNetworkWirelessFailedConnections200ResponseInner) GetSsidNumberOk() (*int32, bool)`

GetSsidNumberOk returns a tuple with the SsidNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidNumber

`func (o *GetNetworkWirelessFailedConnections200ResponseInner) SetSsidNumber(v int32)`

SetSsidNumber sets SsidNumber field to given value.

### HasSsidNumber

`func (o *GetNetworkWirelessFailedConnections200ResponseInner) HasSsidNumber() bool`

HasSsidNumber returns a boolean if a field has been set.

### GetVlan

`func (o *GetNetworkWirelessFailedConnections200ResponseInner) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *GetNetworkWirelessFailedConnections200ResponseInner) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *GetNetworkWirelessFailedConnections200ResponseInner) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *GetNetworkWirelessFailedConnections200ResponseInner) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetClientMac

`func (o *GetNetworkWirelessFailedConnections200ResponseInner) GetClientMac() string`

GetClientMac returns the ClientMac field if non-nil, zero value otherwise.

### GetClientMacOk

`func (o *GetNetworkWirelessFailedConnections200ResponseInner) GetClientMacOk() (*string, bool)`

GetClientMacOk returns a tuple with the ClientMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMac

`func (o *GetNetworkWirelessFailedConnections200ResponseInner) SetClientMac(v string)`

SetClientMac sets ClientMac field to given value.

### HasClientMac

`func (o *GetNetworkWirelessFailedConnections200ResponseInner) HasClientMac() bool`

HasClientMac returns a boolean if a field has been set.

### GetSerial

`func (o *GetNetworkWirelessFailedConnections200ResponseInner) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *GetNetworkWirelessFailedConnections200ResponseInner) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *GetNetworkWirelessFailedConnections200ResponseInner) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *GetNetworkWirelessFailedConnections200ResponseInner) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetFailureStep

`func (o *GetNetworkWirelessFailedConnections200ResponseInner) GetFailureStep() string`

GetFailureStep returns the FailureStep field if non-nil, zero value otherwise.

### GetFailureStepOk

`func (o *GetNetworkWirelessFailedConnections200ResponseInner) GetFailureStepOk() (*string, bool)`

GetFailureStepOk returns a tuple with the FailureStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureStep

`func (o *GetNetworkWirelessFailedConnections200ResponseInner) SetFailureStep(v string)`

SetFailureStep sets FailureStep field to given value.

### HasFailureStep

`func (o *GetNetworkWirelessFailedConnections200ResponseInner) HasFailureStep() bool`

HasFailureStep returns a boolean if a field has been set.

### GetType

`func (o *GetNetworkWirelessFailedConnections200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetNetworkWirelessFailedConnections200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetNetworkWirelessFailedConnections200ResponseInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetNetworkWirelessFailedConnections200ResponseInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTs

`func (o *GetNetworkWirelessFailedConnections200ResponseInner) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetNetworkWirelessFailedConnections200ResponseInner) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetNetworkWirelessFailedConnections200ResponseInner) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetNetworkWirelessFailedConnections200ResponseInner) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


