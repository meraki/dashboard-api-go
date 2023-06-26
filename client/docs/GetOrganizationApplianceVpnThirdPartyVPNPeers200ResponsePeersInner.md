# GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the VPN peer | [optional] 
**PublicIp** | Pointer to **string** | [optional] The public IP of the VPN peer | [optional] 
**RemoteId** | Pointer to **string** | [optional] The remote ID is used to identify the connecting VPN peer. This can either be a valid IPv4 Address, FQDN or User FQDN. | [optional] 
**LocalId** | Pointer to **string** | [optional] The local ID is used to identify the MX to the peer. This will apply to all MXs this peer applies to. | [optional] 
**Secret** | Pointer to **string** | The shared secret with the VPN peer | [optional] 
**PrivateSubnets** | Pointer to **[]string** | The list of the private subnets of the VPN peer | [optional] 
**IpsecPolicies** | Pointer to [**GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies**](GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies.md) |  | [optional] 
**IpsecPoliciesPreset** | Pointer to **string** | One of the following available presets: &#39;default&#39;, &#39;aws&#39;, &#39;azure&#39;. If this is provided, the &#39;ipsecPolicies&#39; parameter is ignored. | [optional] 
**IkeVersion** | Pointer to **string** | [optional] The IKE version to be used for the IPsec VPN peer configuration. Defaults to &#39;1&#39; when omitted. | [optional] [default to "1"]
**NetworkTags** | Pointer to **[]string** | A list of network tags that will connect with this peer. Use [&#39;all&#39;] for all networks. Use [&#39;none&#39;] for no networks. If not included, the default is [&#39;all&#39;]. | [optional] 

## Methods

### NewGetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner

`func NewGetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner() *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner`

NewGetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner instantiates a new GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerWithDefaults

`func NewGetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerWithDefaults() *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner`

NewGetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerWithDefaults instantiates a new GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublicIp

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetRemoteId

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### GetLocalId

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) GetLocalId() string`

GetLocalId returns the LocalId field if non-nil, zero value otherwise.

### GetLocalIdOk

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) GetLocalIdOk() (*string, bool)`

GetLocalIdOk returns a tuple with the LocalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalId

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) SetLocalId(v string)`

SetLocalId sets LocalId field to given value.

### HasLocalId

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) HasLocalId() bool`

HasLocalId returns a boolean if a field has been set.

### GetSecret

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetPrivateSubnets

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) GetPrivateSubnets() []string`

GetPrivateSubnets returns the PrivateSubnets field if non-nil, zero value otherwise.

### GetPrivateSubnetsOk

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) GetPrivateSubnetsOk() (*[]string, bool)`

GetPrivateSubnetsOk returns a tuple with the PrivateSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateSubnets

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) SetPrivateSubnets(v []string)`

SetPrivateSubnets sets PrivateSubnets field to given value.

### HasPrivateSubnets

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) HasPrivateSubnets() bool`

HasPrivateSubnets returns a boolean if a field has been set.

### GetIpsecPolicies

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) GetIpsecPolicies() GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies`

GetIpsecPolicies returns the IpsecPolicies field if non-nil, zero value otherwise.

### GetIpsecPoliciesOk

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) GetIpsecPoliciesOk() (*GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies, bool)`

GetIpsecPoliciesOk returns a tuple with the IpsecPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecPolicies

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) SetIpsecPolicies(v GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInnerIpsecPolicies)`

SetIpsecPolicies sets IpsecPolicies field to given value.

### HasIpsecPolicies

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) HasIpsecPolicies() bool`

HasIpsecPolicies returns a boolean if a field has been set.

### GetIpsecPoliciesPreset

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) GetIpsecPoliciesPreset() string`

GetIpsecPoliciesPreset returns the IpsecPoliciesPreset field if non-nil, zero value otherwise.

### GetIpsecPoliciesPresetOk

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) GetIpsecPoliciesPresetOk() (*string, bool)`

GetIpsecPoliciesPresetOk returns a tuple with the IpsecPoliciesPreset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecPoliciesPreset

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) SetIpsecPoliciesPreset(v string)`

SetIpsecPoliciesPreset sets IpsecPoliciesPreset field to given value.

### HasIpsecPoliciesPreset

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) HasIpsecPoliciesPreset() bool`

HasIpsecPoliciesPreset returns a boolean if a field has been set.

### GetIkeVersion

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) GetIkeVersion() string`

GetIkeVersion returns the IkeVersion field if non-nil, zero value otherwise.

### GetIkeVersionOk

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) GetIkeVersionOk() (*string, bool)`

GetIkeVersionOk returns a tuple with the IkeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeVersion

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) SetIkeVersion(v string)`

SetIkeVersion sets IkeVersion field to given value.

### HasIkeVersion

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) HasIkeVersion() bool`

HasIkeVersion returns a boolean if a field has been set.

### GetNetworkTags

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) GetNetworkTags() []string`

GetNetworkTags returns the NetworkTags field if non-nil, zero value otherwise.

### GetNetworkTagsOk

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) GetNetworkTagsOk() (*[]string, bool)`

GetNetworkTagsOk returns a tuple with the NetworkTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTags

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) SetNetworkTags(v []string)`

SetNetworkTags sets NetworkTags field to given value.

### HasNetworkTags

`func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) HasNetworkTags() bool`

HasNetworkTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


