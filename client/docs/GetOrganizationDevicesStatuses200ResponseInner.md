# GetOrganizationDevicesStatuses200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Device Name | [optional] 
**Serial** | Pointer to **string** | Device Serial Number | [optional] 
**Mac** | Pointer to **string** | MAC Address | [optional] 
**PublicIp** | Pointer to **string** | Public IP Address | [optional] 
**NetworkId** | Pointer to **string** | Network ID | [optional] 
**Status** | Pointer to **string** | Device Status | [optional] 
**LastReportedAt** | Pointer to **string** | Device Last Reported Location | [optional] 
**LanIp** | Pointer to **string** | LAN IP Address | [optional] 
**Gateway** | Pointer to **string** | IP Gateway | [optional] 
**IpType** | Pointer to **string** | IP Type | [optional] 
**PrimaryDns** | Pointer to **string** | Primary DNS | [optional] 
**SecondaryDns** | Pointer to **string** | Secondary DNS | [optional] 
**ProductType** | Pointer to **string** | Product Type | [optional] 
**Components** | Pointer to [**GetOrganizationDevicesStatuses200ResponseInnerComponents**](GetOrganizationDevicesStatuses200ResponseInnerComponents.md) |  | [optional] 
**Model** | Pointer to **string** | Model | [optional] 
**Tags** | Pointer to **[]string** | Tags | [optional] 

## Methods

### NewGetOrganizationDevicesStatuses200ResponseInner

`func NewGetOrganizationDevicesStatuses200ResponseInner() *GetOrganizationDevicesStatuses200ResponseInner`

NewGetOrganizationDevicesStatuses200ResponseInner instantiates a new GetOrganizationDevicesStatuses200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationDevicesStatuses200ResponseInnerWithDefaults

`func NewGetOrganizationDevicesStatuses200ResponseInnerWithDefaults() *GetOrganizationDevicesStatuses200ResponseInner`

NewGetOrganizationDevicesStatuses200ResponseInnerWithDefaults instantiates a new GetOrganizationDevicesStatuses200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetOrganizationDevicesStatuses200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetOrganizationDevicesStatuses200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetOrganizationDevicesStatuses200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetOrganizationDevicesStatuses200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSerial

`func (o *GetOrganizationDevicesStatuses200ResponseInner) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *GetOrganizationDevicesStatuses200ResponseInner) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *GetOrganizationDevicesStatuses200ResponseInner) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *GetOrganizationDevicesStatuses200ResponseInner) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetMac

`func (o *GetOrganizationDevicesStatuses200ResponseInner) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *GetOrganizationDevicesStatuses200ResponseInner) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *GetOrganizationDevicesStatuses200ResponseInner) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *GetOrganizationDevicesStatuses200ResponseInner) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetPublicIp

`func (o *GetOrganizationDevicesStatuses200ResponseInner) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *GetOrganizationDevicesStatuses200ResponseInner) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *GetOrganizationDevicesStatuses200ResponseInner) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *GetOrganizationDevicesStatuses200ResponseInner) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetNetworkId

`func (o *GetOrganizationDevicesStatuses200ResponseInner) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *GetOrganizationDevicesStatuses200ResponseInner) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *GetOrganizationDevicesStatuses200ResponseInner) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *GetOrganizationDevicesStatuses200ResponseInner) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetStatus

`func (o *GetOrganizationDevicesStatuses200ResponseInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetOrganizationDevicesStatuses200ResponseInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetOrganizationDevicesStatuses200ResponseInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetOrganizationDevicesStatuses200ResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLastReportedAt

`func (o *GetOrganizationDevicesStatuses200ResponseInner) GetLastReportedAt() string`

GetLastReportedAt returns the LastReportedAt field if non-nil, zero value otherwise.

### GetLastReportedAtOk

`func (o *GetOrganizationDevicesStatuses200ResponseInner) GetLastReportedAtOk() (*string, bool)`

GetLastReportedAtOk returns a tuple with the LastReportedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReportedAt

`func (o *GetOrganizationDevicesStatuses200ResponseInner) SetLastReportedAt(v string)`

SetLastReportedAt sets LastReportedAt field to given value.

### HasLastReportedAt

`func (o *GetOrganizationDevicesStatuses200ResponseInner) HasLastReportedAt() bool`

HasLastReportedAt returns a boolean if a field has been set.

### GetLanIp

`func (o *GetOrganizationDevicesStatuses200ResponseInner) GetLanIp() string`

GetLanIp returns the LanIp field if non-nil, zero value otherwise.

### GetLanIpOk

`func (o *GetOrganizationDevicesStatuses200ResponseInner) GetLanIpOk() (*string, bool)`

GetLanIpOk returns a tuple with the LanIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanIp

`func (o *GetOrganizationDevicesStatuses200ResponseInner) SetLanIp(v string)`

SetLanIp sets LanIp field to given value.

### HasLanIp

`func (o *GetOrganizationDevicesStatuses200ResponseInner) HasLanIp() bool`

HasLanIp returns a boolean if a field has been set.

### GetGateway

`func (o *GetOrganizationDevicesStatuses200ResponseInner) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *GetOrganizationDevicesStatuses200ResponseInner) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *GetOrganizationDevicesStatuses200ResponseInner) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *GetOrganizationDevicesStatuses200ResponseInner) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetIpType

`func (o *GetOrganizationDevicesStatuses200ResponseInner) GetIpType() string`

GetIpType returns the IpType field if non-nil, zero value otherwise.

### GetIpTypeOk

`func (o *GetOrganizationDevicesStatuses200ResponseInner) GetIpTypeOk() (*string, bool)`

GetIpTypeOk returns a tuple with the IpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpType

`func (o *GetOrganizationDevicesStatuses200ResponseInner) SetIpType(v string)`

SetIpType sets IpType field to given value.

### HasIpType

`func (o *GetOrganizationDevicesStatuses200ResponseInner) HasIpType() bool`

HasIpType returns a boolean if a field has been set.

### GetPrimaryDns

`func (o *GetOrganizationDevicesStatuses200ResponseInner) GetPrimaryDns() string`

GetPrimaryDns returns the PrimaryDns field if non-nil, zero value otherwise.

### GetPrimaryDnsOk

`func (o *GetOrganizationDevicesStatuses200ResponseInner) GetPrimaryDnsOk() (*string, bool)`

GetPrimaryDnsOk returns a tuple with the PrimaryDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDns

`func (o *GetOrganizationDevicesStatuses200ResponseInner) SetPrimaryDns(v string)`

SetPrimaryDns sets PrimaryDns field to given value.

### HasPrimaryDns

`func (o *GetOrganizationDevicesStatuses200ResponseInner) HasPrimaryDns() bool`

HasPrimaryDns returns a boolean if a field has been set.

### GetSecondaryDns

`func (o *GetOrganizationDevicesStatuses200ResponseInner) GetSecondaryDns() string`

GetSecondaryDns returns the SecondaryDns field if non-nil, zero value otherwise.

### GetSecondaryDnsOk

`func (o *GetOrganizationDevicesStatuses200ResponseInner) GetSecondaryDnsOk() (*string, bool)`

GetSecondaryDnsOk returns a tuple with the SecondaryDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDns

`func (o *GetOrganizationDevicesStatuses200ResponseInner) SetSecondaryDns(v string)`

SetSecondaryDns sets SecondaryDns field to given value.

### HasSecondaryDns

`func (o *GetOrganizationDevicesStatuses200ResponseInner) HasSecondaryDns() bool`

HasSecondaryDns returns a boolean if a field has been set.

### GetProductType

`func (o *GetOrganizationDevicesStatuses200ResponseInner) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *GetOrganizationDevicesStatuses200ResponseInner) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *GetOrganizationDevicesStatuses200ResponseInner) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *GetOrganizationDevicesStatuses200ResponseInner) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetComponents

`func (o *GetOrganizationDevicesStatuses200ResponseInner) GetComponents() GetOrganizationDevicesStatuses200ResponseInnerComponents`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *GetOrganizationDevicesStatuses200ResponseInner) GetComponentsOk() (*GetOrganizationDevicesStatuses200ResponseInnerComponents, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *GetOrganizationDevicesStatuses200ResponseInner) SetComponents(v GetOrganizationDevicesStatuses200ResponseInnerComponents)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *GetOrganizationDevicesStatuses200ResponseInner) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetModel

`func (o *GetOrganizationDevicesStatuses200ResponseInner) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *GetOrganizationDevicesStatuses200ResponseInner) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *GetOrganizationDevicesStatuses200ResponseInner) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *GetOrganizationDevicesStatuses200ResponseInner) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetTags

`func (o *GetOrganizationDevicesStatuses200ResponseInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetOrganizationDevicesStatuses200ResponseInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetOrganizationDevicesStatuses200ResponseInner) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetOrganizationDevicesStatuses200ResponseInner) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


