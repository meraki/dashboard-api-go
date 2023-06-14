# GetOrganizationDevicesStatuses200Response

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
**Components** | Pointer to [**GetOrganizationDevicesStatuses200ResponseComponents**](GetOrganizationDevicesStatuses200ResponseComponents.md) |  | [optional] 
**Model** | Pointer to **string** | Model | [optional] 
**Tags** | Pointer to **[]string** | Tags | [optional] 

## Methods

### NewGetOrganizationDevicesStatuses200Response

`func NewGetOrganizationDevicesStatuses200Response() *GetOrganizationDevicesStatuses200Response`

NewGetOrganizationDevicesStatuses200Response instantiates a new GetOrganizationDevicesStatuses200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationDevicesStatuses200ResponseWithDefaults

`func NewGetOrganizationDevicesStatuses200ResponseWithDefaults() *GetOrganizationDevicesStatuses200Response`

NewGetOrganizationDevicesStatuses200ResponseWithDefaults instantiates a new GetOrganizationDevicesStatuses200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetOrganizationDevicesStatuses200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetOrganizationDevicesStatuses200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetOrganizationDevicesStatuses200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetOrganizationDevicesStatuses200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSerial

`func (o *GetOrganizationDevicesStatuses200Response) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *GetOrganizationDevicesStatuses200Response) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *GetOrganizationDevicesStatuses200Response) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *GetOrganizationDevicesStatuses200Response) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetMac

`func (o *GetOrganizationDevicesStatuses200Response) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *GetOrganizationDevicesStatuses200Response) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *GetOrganizationDevicesStatuses200Response) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *GetOrganizationDevicesStatuses200Response) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetPublicIp

`func (o *GetOrganizationDevicesStatuses200Response) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *GetOrganizationDevicesStatuses200Response) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *GetOrganizationDevicesStatuses200Response) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *GetOrganizationDevicesStatuses200Response) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetNetworkId

`func (o *GetOrganizationDevicesStatuses200Response) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *GetOrganizationDevicesStatuses200Response) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *GetOrganizationDevicesStatuses200Response) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *GetOrganizationDevicesStatuses200Response) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetStatus

`func (o *GetOrganizationDevicesStatuses200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetOrganizationDevicesStatuses200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetOrganizationDevicesStatuses200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetOrganizationDevicesStatuses200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLastReportedAt

`func (o *GetOrganizationDevicesStatuses200Response) GetLastReportedAt() string`

GetLastReportedAt returns the LastReportedAt field if non-nil, zero value otherwise.

### GetLastReportedAtOk

`func (o *GetOrganizationDevicesStatuses200Response) GetLastReportedAtOk() (*string, bool)`

GetLastReportedAtOk returns a tuple with the LastReportedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReportedAt

`func (o *GetOrganizationDevicesStatuses200Response) SetLastReportedAt(v string)`

SetLastReportedAt sets LastReportedAt field to given value.

### HasLastReportedAt

`func (o *GetOrganizationDevicesStatuses200Response) HasLastReportedAt() bool`

HasLastReportedAt returns a boolean if a field has been set.

### GetLanIp

`func (o *GetOrganizationDevicesStatuses200Response) GetLanIp() string`

GetLanIp returns the LanIp field if non-nil, zero value otherwise.

### GetLanIpOk

`func (o *GetOrganizationDevicesStatuses200Response) GetLanIpOk() (*string, bool)`

GetLanIpOk returns a tuple with the LanIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanIp

`func (o *GetOrganizationDevicesStatuses200Response) SetLanIp(v string)`

SetLanIp sets LanIp field to given value.

### HasLanIp

`func (o *GetOrganizationDevicesStatuses200Response) HasLanIp() bool`

HasLanIp returns a boolean if a field has been set.

### GetGateway

`func (o *GetOrganizationDevicesStatuses200Response) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *GetOrganizationDevicesStatuses200Response) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *GetOrganizationDevicesStatuses200Response) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *GetOrganizationDevicesStatuses200Response) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetIpType

`func (o *GetOrganizationDevicesStatuses200Response) GetIpType() string`

GetIpType returns the IpType field if non-nil, zero value otherwise.

### GetIpTypeOk

`func (o *GetOrganizationDevicesStatuses200Response) GetIpTypeOk() (*string, bool)`

GetIpTypeOk returns a tuple with the IpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpType

`func (o *GetOrganizationDevicesStatuses200Response) SetIpType(v string)`

SetIpType sets IpType field to given value.

### HasIpType

`func (o *GetOrganizationDevicesStatuses200Response) HasIpType() bool`

HasIpType returns a boolean if a field has been set.

### GetPrimaryDns

`func (o *GetOrganizationDevicesStatuses200Response) GetPrimaryDns() string`

GetPrimaryDns returns the PrimaryDns field if non-nil, zero value otherwise.

### GetPrimaryDnsOk

`func (o *GetOrganizationDevicesStatuses200Response) GetPrimaryDnsOk() (*string, bool)`

GetPrimaryDnsOk returns a tuple with the PrimaryDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDns

`func (o *GetOrganizationDevicesStatuses200Response) SetPrimaryDns(v string)`

SetPrimaryDns sets PrimaryDns field to given value.

### HasPrimaryDns

`func (o *GetOrganizationDevicesStatuses200Response) HasPrimaryDns() bool`

HasPrimaryDns returns a boolean if a field has been set.

### GetSecondaryDns

`func (o *GetOrganizationDevicesStatuses200Response) GetSecondaryDns() string`

GetSecondaryDns returns the SecondaryDns field if non-nil, zero value otherwise.

### GetSecondaryDnsOk

`func (o *GetOrganizationDevicesStatuses200Response) GetSecondaryDnsOk() (*string, bool)`

GetSecondaryDnsOk returns a tuple with the SecondaryDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDns

`func (o *GetOrganizationDevicesStatuses200Response) SetSecondaryDns(v string)`

SetSecondaryDns sets SecondaryDns field to given value.

### HasSecondaryDns

`func (o *GetOrganizationDevicesStatuses200Response) HasSecondaryDns() bool`

HasSecondaryDns returns a boolean if a field has been set.

### GetProductType

`func (o *GetOrganizationDevicesStatuses200Response) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *GetOrganizationDevicesStatuses200Response) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *GetOrganizationDevicesStatuses200Response) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *GetOrganizationDevicesStatuses200Response) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetComponents

`func (o *GetOrganizationDevicesStatuses200Response) GetComponents() GetOrganizationDevicesStatuses200ResponseComponents`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *GetOrganizationDevicesStatuses200Response) GetComponentsOk() (*GetOrganizationDevicesStatuses200ResponseComponents, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *GetOrganizationDevicesStatuses200Response) SetComponents(v GetOrganizationDevicesStatuses200ResponseComponents)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *GetOrganizationDevicesStatuses200Response) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetModel

`func (o *GetOrganizationDevicesStatuses200Response) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *GetOrganizationDevicesStatuses200Response) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *GetOrganizationDevicesStatuses200Response) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *GetOrganizationDevicesStatuses200Response) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetTags

`func (o *GetOrganizationDevicesStatuses200Response) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetOrganizationDevicesStatuses200Response) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetOrganizationDevicesStatuses200Response) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetOrganizationDevicesStatuses200Response) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


