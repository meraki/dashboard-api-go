# UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the VPN peer | 
**PublicIp** | **string** | The public IP of the VPN peer | 
**RemoteId** | Pointer to **string** | [optional] The remote ID is used to identify the connecting VPN peer. This can either be a valid IPv4 Address, FQDN or User FQDN. | [optional] 
**PrivateSubnets** | **[]string** | The list of the private subnets of the VPN peer | 
**IpsecPolicies** | Pointer to [**UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInnerIpsecPolicies**](UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInnerIpsecPolicies.md) |  | [optional] 
**IpsecPoliciesPreset** | Pointer to **string** | One of the following available presets: &#39;default&#39;, &#39;aws&#39;, &#39;azure&#39;. If this is provided, the &#39;ipsecPolicies&#39; parameter is ignored. | [optional] 
**Secret** | **string** | The shared secret with the VPN peer | 
**IkeVersion** | Pointer to **string** | [optional] The IKE version to be used for the IPsec VPN peer configuration. Defaults to &#39;1&#39; when omitted. | [optional] [default to "1"]
**NetworkTags** | Pointer to **[]string** | A list of network tags that will connect with this peer. Use [&#39;all&#39;] for all networks. Use [&#39;none&#39;] for no networks. If not included, the default is [&#39;all&#39;]. | [optional] 

## Methods

### NewUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner

`func NewUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner(name string, publicIp string, privateSubnets []string, secret string, ) *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner`

NewUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner instantiates a new UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInnerWithDefaults

`func NewUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInnerWithDefaults() *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner`

NewUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInnerWithDefaults instantiates a new UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) SetName(v string)`

SetName sets Name field to given value.


### GetPublicIp

`func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.


### GetRemoteId

`func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### GetPrivateSubnets

`func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) GetPrivateSubnets() []string`

GetPrivateSubnets returns the PrivateSubnets field if non-nil, zero value otherwise.

### GetPrivateSubnetsOk

`func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) GetPrivateSubnetsOk() (*[]string, bool)`

GetPrivateSubnetsOk returns a tuple with the PrivateSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateSubnets

`func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) SetPrivateSubnets(v []string)`

SetPrivateSubnets sets PrivateSubnets field to given value.


### GetIpsecPolicies

`func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) GetIpsecPolicies() UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInnerIpsecPolicies`

GetIpsecPolicies returns the IpsecPolicies field if non-nil, zero value otherwise.

### GetIpsecPoliciesOk

`func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) GetIpsecPoliciesOk() (*UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInnerIpsecPolicies, bool)`

GetIpsecPoliciesOk returns a tuple with the IpsecPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecPolicies

`func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) SetIpsecPolicies(v UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInnerIpsecPolicies)`

SetIpsecPolicies sets IpsecPolicies field to given value.

### HasIpsecPolicies

`func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) HasIpsecPolicies() bool`

HasIpsecPolicies returns a boolean if a field has been set.

### GetIpsecPoliciesPreset

`func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) GetIpsecPoliciesPreset() string`

GetIpsecPoliciesPreset returns the IpsecPoliciesPreset field if non-nil, zero value otherwise.

### GetIpsecPoliciesPresetOk

`func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) GetIpsecPoliciesPresetOk() (*string, bool)`

GetIpsecPoliciesPresetOk returns a tuple with the IpsecPoliciesPreset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecPoliciesPreset

`func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) SetIpsecPoliciesPreset(v string)`

SetIpsecPoliciesPreset sets IpsecPoliciesPreset field to given value.

### HasIpsecPoliciesPreset

`func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) HasIpsecPoliciesPreset() bool`

HasIpsecPoliciesPreset returns a boolean if a field has been set.

### GetSecret

`func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetIkeVersion

`func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) GetIkeVersion() string`

GetIkeVersion returns the IkeVersion field if non-nil, zero value otherwise.

### GetIkeVersionOk

`func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) GetIkeVersionOk() (*string, bool)`

GetIkeVersionOk returns a tuple with the IkeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeVersion

`func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) SetIkeVersion(v string)`

SetIkeVersion sets IkeVersion field to given value.

### HasIkeVersion

`func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) HasIkeVersion() bool`

HasIkeVersion returns a boolean if a field has been set.

### GetNetworkTags

`func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) GetNetworkTags() []string`

GetNetworkTags returns the NetworkTags field if non-nil, zero value otherwise.

### GetNetworkTagsOk

`func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) GetNetworkTagsOk() (*[]string, bool)`

GetNetworkTagsOk returns a tuple with the NetworkTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTags

`func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) SetNetworkTags(v []string)`

SetNetworkTags sets NetworkTags field to given value.

### HasNetworkTags

`func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) HasNetworkTags() bool`

HasNetworkTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


