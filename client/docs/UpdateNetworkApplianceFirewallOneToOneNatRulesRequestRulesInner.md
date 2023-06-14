# UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A descriptive name for the rule | [optional] 
**PublicIp** | Pointer to **string** | The IP address that will be used to access the internal resource from the WAN | [optional] 
**LanIp** | **string** | The IP address of the server or device that hosts the internal resource that you wish to make available on the WAN | 
**Uplink** | Pointer to **string** | The physical WAN interface on which the traffic will arrive (&#39;internet1&#39; or, if available, &#39;internet2&#39;) | [optional] 
**AllowedInbound** | Pointer to [**[]UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner**](UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner.md) | The ports this mapping will provide access on, and the remote IPs that will be allowed access to the resource | [optional] 

## Methods

### NewUpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner

`func NewUpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner(lanIp string, ) *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner`

NewUpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner instantiates a new UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerWithDefaults

`func NewUpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerWithDefaults() *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner`

NewUpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerWithDefaults instantiates a new UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublicIp

`func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetLanIp

`func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) GetLanIp() string`

GetLanIp returns the LanIp field if non-nil, zero value otherwise.

### GetLanIpOk

`func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) GetLanIpOk() (*string, bool)`

GetLanIpOk returns a tuple with the LanIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanIp

`func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) SetLanIp(v string)`

SetLanIp sets LanIp field to given value.


### GetUplink

`func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) GetUplink() string`

GetUplink returns the Uplink field if non-nil, zero value otherwise.

### GetUplinkOk

`func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) GetUplinkOk() (*string, bool)`

GetUplinkOk returns a tuple with the Uplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplink

`func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) SetUplink(v string)`

SetUplink sets Uplink field to given value.

### HasUplink

`func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) HasUplink() bool`

HasUplink returns a boolean if a field has been set.

### GetAllowedInbound

`func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) GetAllowedInbound() []UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner`

GetAllowedInbound returns the AllowedInbound field if non-nil, zero value otherwise.

### GetAllowedInboundOk

`func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) GetAllowedInboundOk() (*[]UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner, bool)`

GetAllowedInboundOk returns a tuple with the AllowedInbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedInbound

`func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) SetAllowedInbound(v []UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner)`

SetAllowedInbound sets AllowedInbound field to given value.

### HasAllowedInbound

`func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) HasAllowedInbound() bool`

HasAllowedInbound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


