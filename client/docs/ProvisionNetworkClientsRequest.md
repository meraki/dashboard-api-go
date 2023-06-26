# ProvisionNetworkClientsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clients** | [**[]ProvisionNetworkClientsRequestClientsInner**](ProvisionNetworkClientsRequestClientsInner.md) | The array of clients to provision | 
**DevicePolicy** | **string** | The policy to apply to the specified client. Can be &#39;Group policy&#39;, &#39;Allowed&#39;, &#39;Blocked&#39;, &#39;Per connection&#39; or &#39;Normal&#39;. Required. | 
**GroupPolicyId** | Pointer to **string** | The ID of the desired group policy to apply to the client. Required if &#39;devicePolicy&#39; is set to \&quot;Group policy\&quot;. Otherwise this is ignored. | [optional] 
**PoliciesBySecurityAppliance** | Pointer to [**ProvisionNetworkClientsRequestPoliciesBySecurityAppliance**](ProvisionNetworkClientsRequestPoliciesBySecurityAppliance.md) |  | [optional] 
**PoliciesBySsid** | Pointer to [**ProvisionNetworkClientsRequestPoliciesBySsid**](ProvisionNetworkClientsRequestPoliciesBySsid.md) |  | [optional] 

## Methods

### NewProvisionNetworkClientsRequest

`func NewProvisionNetworkClientsRequest(clients []ProvisionNetworkClientsRequestClientsInner, devicePolicy string, ) *ProvisionNetworkClientsRequest`

NewProvisionNetworkClientsRequest instantiates a new ProvisionNetworkClientsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisionNetworkClientsRequestWithDefaults

`func NewProvisionNetworkClientsRequestWithDefaults() *ProvisionNetworkClientsRequest`

NewProvisionNetworkClientsRequestWithDefaults instantiates a new ProvisionNetworkClientsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClients

`func (o *ProvisionNetworkClientsRequest) GetClients() []ProvisionNetworkClientsRequestClientsInner`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *ProvisionNetworkClientsRequest) GetClientsOk() (*[]ProvisionNetworkClientsRequestClientsInner, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *ProvisionNetworkClientsRequest) SetClients(v []ProvisionNetworkClientsRequestClientsInner)`

SetClients sets Clients field to given value.


### GetDevicePolicy

`func (o *ProvisionNetworkClientsRequest) GetDevicePolicy() string`

GetDevicePolicy returns the DevicePolicy field if non-nil, zero value otherwise.

### GetDevicePolicyOk

`func (o *ProvisionNetworkClientsRequest) GetDevicePolicyOk() (*string, bool)`

GetDevicePolicyOk returns a tuple with the DevicePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePolicy

`func (o *ProvisionNetworkClientsRequest) SetDevicePolicy(v string)`

SetDevicePolicy sets DevicePolicy field to given value.


### GetGroupPolicyId

`func (o *ProvisionNetworkClientsRequest) GetGroupPolicyId() string`

GetGroupPolicyId returns the GroupPolicyId field if non-nil, zero value otherwise.

### GetGroupPolicyIdOk

`func (o *ProvisionNetworkClientsRequest) GetGroupPolicyIdOk() (*string, bool)`

GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPolicyId

`func (o *ProvisionNetworkClientsRequest) SetGroupPolicyId(v string)`

SetGroupPolicyId sets GroupPolicyId field to given value.

### HasGroupPolicyId

`func (o *ProvisionNetworkClientsRequest) HasGroupPolicyId() bool`

HasGroupPolicyId returns a boolean if a field has been set.

### GetPoliciesBySecurityAppliance

`func (o *ProvisionNetworkClientsRequest) GetPoliciesBySecurityAppliance() ProvisionNetworkClientsRequestPoliciesBySecurityAppliance`

GetPoliciesBySecurityAppliance returns the PoliciesBySecurityAppliance field if non-nil, zero value otherwise.

### GetPoliciesBySecurityApplianceOk

`func (o *ProvisionNetworkClientsRequest) GetPoliciesBySecurityApplianceOk() (*ProvisionNetworkClientsRequestPoliciesBySecurityAppliance, bool)`

GetPoliciesBySecurityApplianceOk returns a tuple with the PoliciesBySecurityAppliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoliciesBySecurityAppliance

`func (o *ProvisionNetworkClientsRequest) SetPoliciesBySecurityAppliance(v ProvisionNetworkClientsRequestPoliciesBySecurityAppliance)`

SetPoliciesBySecurityAppliance sets PoliciesBySecurityAppliance field to given value.

### HasPoliciesBySecurityAppliance

`func (o *ProvisionNetworkClientsRequest) HasPoliciesBySecurityAppliance() bool`

HasPoliciesBySecurityAppliance returns a boolean if a field has been set.

### GetPoliciesBySsid

`func (o *ProvisionNetworkClientsRequest) GetPoliciesBySsid() ProvisionNetworkClientsRequestPoliciesBySsid`

GetPoliciesBySsid returns the PoliciesBySsid field if non-nil, zero value otherwise.

### GetPoliciesBySsidOk

`func (o *ProvisionNetworkClientsRequest) GetPoliciesBySsidOk() (*ProvisionNetworkClientsRequestPoliciesBySsid, bool)`

GetPoliciesBySsidOk returns a tuple with the PoliciesBySsid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoliciesBySsid

`func (o *ProvisionNetworkClientsRequest) SetPoliciesBySsid(v ProvisionNetworkClientsRequestPoliciesBySsid)`

SetPoliciesBySsid sets PoliciesBySsid field to given value.

### HasPoliciesBySsid

`func (o *ProvisionNetworkClientsRequest) HasPoliciesBySsid() bool`

HasPoliciesBySsid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


