# GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interface** | Pointer to **string** | Uplink interface | [optional] 
**Status** | Pointer to **string** | Uplink status | [optional] 
**Ip** | Pointer to **string** | Uplink IP | [optional] 
**Provider** | Pointer to **string** | Network Provider | [optional] 
**PublicIp** | Pointer to **string** | Public IP | [optional] 
**Model** | Pointer to **string** | Uplink model | [optional] 
**SignalStat** | Pointer to [**GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat**](GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat.md) |  | [optional] 
**ConnectionType** | Pointer to **string** | Connection Type | [optional] 
**Apn** | Pointer to **string** | Access Point Name | [optional] 
**Gateway** | Pointer to **string** | Gateway IP | [optional] 
**Dns1** | Pointer to **string** | Primary DNS IP | [optional] 
**Dns2** | Pointer to **string** | Secondary DNS IP | [optional] 
**SignalType** | Pointer to **string** | Signal Type | [optional] 
**Iccid** | Pointer to **string** | Integrated Circuit Card Identification Number | [optional] 

## Methods

### NewGetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner

`func NewGetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner() *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner`

NewGetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner instantiates a new GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerWithDefaults

`func NewGetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerWithDefaults() *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner`

NewGetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerWithDefaults instantiates a new GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterface

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetStatus

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIp

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetProvider

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetPublicIp

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetModel

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetSignalStat

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetSignalStat() GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat`

GetSignalStat returns the SignalStat field if non-nil, zero value otherwise.

### GetSignalStatOk

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetSignalStatOk() (*GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat, bool)`

GetSignalStatOk returns a tuple with the SignalStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalStat

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) SetSignalStat(v GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat)`

SetSignalStat sets SignalStat field to given value.

### HasSignalStat

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) HasSignalStat() bool`

HasSignalStat returns a boolean if a field has been set.

### GetConnectionType

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetApn

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetApn() string`

GetApn returns the Apn field if non-nil, zero value otherwise.

### GetApnOk

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetApnOk() (*string, bool)`

GetApnOk returns a tuple with the Apn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApn

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) SetApn(v string)`

SetApn sets Apn field to given value.

### HasApn

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) HasApn() bool`

HasApn returns a boolean if a field has been set.

### GetGateway

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetDns1

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetDns1() string`

GetDns1 returns the Dns1 field if non-nil, zero value otherwise.

### GetDns1Ok

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetDns1Ok() (*string, bool)`

GetDns1Ok returns a tuple with the Dns1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns1

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) SetDns1(v string)`

SetDns1 sets Dns1 field to given value.

### HasDns1

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) HasDns1() bool`

HasDns1 returns a boolean if a field has been set.

### GetDns2

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetDns2() string`

GetDns2 returns the Dns2 field if non-nil, zero value otherwise.

### GetDns2Ok

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetDns2Ok() (*string, bool)`

GetDns2Ok returns a tuple with the Dns2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns2

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) SetDns2(v string)`

SetDns2 sets Dns2 field to given value.

### HasDns2

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) HasDns2() bool`

HasDns2 returns a boolean if a field has been set.

### GetSignalType

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetSignalType() string`

GetSignalType returns the SignalType field if non-nil, zero value otherwise.

### GetSignalTypeOk

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetSignalTypeOk() (*string, bool)`

GetSignalTypeOk returns a tuple with the SignalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalType

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) SetSignalType(v string)`

SetSignalType sets SignalType field to given value.

### HasSignalType

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) HasSignalType() bool`

HasSignalType returns a boolean if a field has been set.

### GetIccid

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetIccid() string`

GetIccid returns the Iccid field if non-nil, zero value otherwise.

### GetIccidOk

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) GetIccidOk() (*string, bool)`

GetIccidOk returns a tuple with the Iccid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIccid

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) SetIccid(v string)`

SetIccid sets Iccid field to given value.

### HasIccid

`func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInner) HasIccid() bool`

HasIccid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


