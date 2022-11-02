# GetOrganizationUplinksStatuses200ResponseInnerUplinksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interface** | Pointer to **string** | Uplink interface | [optional] 
**Status** | Pointer to **string** | Uplink status | [optional] 
**Ip** | Pointer to **string** | Uplink IP | [optional] 
**Gateway** | Pointer to **string** | Gateway IP | [optional] 
**PublicIp** | Pointer to **string** | Public IP | [optional] 
**PrimaryDns** | Pointer to **string** | Primary DNS IP | [optional] 
**SecondaryDns** | Pointer to **string** | Secondary DNS IP | [optional] 
**IpAssignedBy** | Pointer to **string** | The way in which the IP is assigned | [optional] 
**Provider** | Pointer to **string** | Network Provider | [optional] 
**SignalStat** | Pointer to [**GetOrganizationUplinksStatuses200ResponseInnerUplinksInnerSignalStat**](GetOrganizationUplinksStatuses200ResponseInnerUplinksInnerSignalStat.md) |  | [optional] 
**ConnectionType** | Pointer to **string** | Connection Type | [optional] 
**Apn** | Pointer to **string** | Access Point Name | [optional] 
**Dns1** | Pointer to **string** | Primary DNS IP | [optional] 
**Dns2** | Pointer to **string** | Secondary DNS IP | [optional] 
**SignalType** | Pointer to **string** | Signal Type | [optional] 
**Iccid** | Pointer to **string** | Integrated Circuit Card Identification Number | [optional] 

## Methods

### NewGetOrganizationUplinksStatuses200ResponseInnerUplinksInner

`func NewGetOrganizationUplinksStatuses200ResponseInnerUplinksInner() *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner`

NewGetOrganizationUplinksStatuses200ResponseInnerUplinksInner instantiates a new GetOrganizationUplinksStatuses200ResponseInnerUplinksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationUplinksStatuses200ResponseInnerUplinksInnerWithDefaults

`func NewGetOrganizationUplinksStatuses200ResponseInnerUplinksInnerWithDefaults() *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner`

NewGetOrganizationUplinksStatuses200ResponseInnerUplinksInnerWithDefaults instantiates a new GetOrganizationUplinksStatuses200ResponseInnerUplinksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterface

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetStatus

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIp

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetGateway

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetPublicIp

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetPrimaryDns

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) GetPrimaryDns() string`

GetPrimaryDns returns the PrimaryDns field if non-nil, zero value otherwise.

### GetPrimaryDnsOk

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) GetPrimaryDnsOk() (*string, bool)`

GetPrimaryDnsOk returns a tuple with the PrimaryDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDns

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) SetPrimaryDns(v string)`

SetPrimaryDns sets PrimaryDns field to given value.

### HasPrimaryDns

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) HasPrimaryDns() bool`

HasPrimaryDns returns a boolean if a field has been set.

### GetSecondaryDns

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) GetSecondaryDns() string`

GetSecondaryDns returns the SecondaryDns field if non-nil, zero value otherwise.

### GetSecondaryDnsOk

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) GetSecondaryDnsOk() (*string, bool)`

GetSecondaryDnsOk returns a tuple with the SecondaryDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDns

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) SetSecondaryDns(v string)`

SetSecondaryDns sets SecondaryDns field to given value.

### HasSecondaryDns

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) HasSecondaryDns() bool`

HasSecondaryDns returns a boolean if a field has been set.

### GetIpAssignedBy

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) GetIpAssignedBy() string`

GetIpAssignedBy returns the IpAssignedBy field if non-nil, zero value otherwise.

### GetIpAssignedByOk

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) GetIpAssignedByOk() (*string, bool)`

GetIpAssignedByOk returns a tuple with the IpAssignedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAssignedBy

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) SetIpAssignedBy(v string)`

SetIpAssignedBy sets IpAssignedBy field to given value.

### HasIpAssignedBy

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) HasIpAssignedBy() bool`

HasIpAssignedBy returns a boolean if a field has been set.

### GetProvider

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetSignalStat

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) GetSignalStat() GetOrganizationUplinksStatuses200ResponseInnerUplinksInnerSignalStat`

GetSignalStat returns the SignalStat field if non-nil, zero value otherwise.

### GetSignalStatOk

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) GetSignalStatOk() (*GetOrganizationUplinksStatuses200ResponseInnerUplinksInnerSignalStat, bool)`

GetSignalStatOk returns a tuple with the SignalStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalStat

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) SetSignalStat(v GetOrganizationUplinksStatuses200ResponseInnerUplinksInnerSignalStat)`

SetSignalStat sets SignalStat field to given value.

### HasSignalStat

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) HasSignalStat() bool`

HasSignalStat returns a boolean if a field has been set.

### GetConnectionType

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetApn

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) GetApn() string`

GetApn returns the Apn field if non-nil, zero value otherwise.

### GetApnOk

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) GetApnOk() (*string, bool)`

GetApnOk returns a tuple with the Apn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApn

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) SetApn(v string)`

SetApn sets Apn field to given value.

### HasApn

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) HasApn() bool`

HasApn returns a boolean if a field has been set.

### GetDns1

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) GetDns1() string`

GetDns1 returns the Dns1 field if non-nil, zero value otherwise.

### GetDns1Ok

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) GetDns1Ok() (*string, bool)`

GetDns1Ok returns a tuple with the Dns1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns1

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) SetDns1(v string)`

SetDns1 sets Dns1 field to given value.

### HasDns1

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) HasDns1() bool`

HasDns1 returns a boolean if a field has been set.

### GetDns2

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) GetDns2() string`

GetDns2 returns the Dns2 field if non-nil, zero value otherwise.

### GetDns2Ok

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) GetDns2Ok() (*string, bool)`

GetDns2Ok returns a tuple with the Dns2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns2

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) SetDns2(v string)`

SetDns2 sets Dns2 field to given value.

### HasDns2

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) HasDns2() bool`

HasDns2 returns a boolean if a field has been set.

### GetSignalType

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) GetSignalType() string`

GetSignalType returns the SignalType field if non-nil, zero value otherwise.

### GetSignalTypeOk

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) GetSignalTypeOk() (*string, bool)`

GetSignalTypeOk returns a tuple with the SignalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalType

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) SetSignalType(v string)`

SetSignalType sets SignalType field to given value.

### HasSignalType

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) HasSignalType() bool`

HasSignalType returns a boolean if a field has been set.

### GetIccid

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) GetIccid() string`

GetIccid returns the Iccid field if non-nil, zero value otherwise.

### GetIccidOk

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) GetIccidOk() (*string, bool)`

GetIccidOk returns a tuple with the Iccid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIccid

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) SetIccid(v string)`

SetIccid sets Iccid field to given value.

### HasIccid

`func (o *GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) HasIccid() bool`

HasIccid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


