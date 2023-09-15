# UpdateNetworkApplianceTrafficShapingVpnExclusions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | **string** | ID of the network whose VPN exclusion rules are returned. | 
**NetworkName** | **string** | Name of the network whose VPN exclusion rules are returned. | 
**Custom** | [**[]UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner**](UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner.md) | Custom VPN exclusion rules. | 
**MajorApplications** | [**[]UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner**](UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner.md) | Major Application based VPN exclusion rules. | 

## Methods

### NewUpdateNetworkApplianceTrafficShapingVpnExclusions200Response

`func NewUpdateNetworkApplianceTrafficShapingVpnExclusions200Response(networkId string, networkName string, custom []UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner, majorApplications []UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner, ) *UpdateNetworkApplianceTrafficShapingVpnExclusions200Response`

NewUpdateNetworkApplianceTrafficShapingVpnExclusions200Response instantiates a new UpdateNetworkApplianceTrafficShapingVpnExclusions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseWithDefaults

`func NewUpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseWithDefaults() *UpdateNetworkApplianceTrafficShapingVpnExclusions200Response`

NewUpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseWithDefaults instantiates a new UpdateNetworkApplianceTrafficShapingVpnExclusions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *UpdateNetworkApplianceTrafficShapingVpnExclusions200Response) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *UpdateNetworkApplianceTrafficShapingVpnExclusions200Response) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *UpdateNetworkApplianceTrafficShapingVpnExclusions200Response) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.


### GetNetworkName

`func (o *UpdateNetworkApplianceTrafficShapingVpnExclusions200Response) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *UpdateNetworkApplianceTrafficShapingVpnExclusions200Response) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *UpdateNetworkApplianceTrafficShapingVpnExclusions200Response) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.


### GetCustom

`func (o *UpdateNetworkApplianceTrafficShapingVpnExclusions200Response) GetCustom() []UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *UpdateNetworkApplianceTrafficShapingVpnExclusions200Response) GetCustomOk() (*[]UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *UpdateNetworkApplianceTrafficShapingVpnExclusions200Response) SetCustom(v []UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseCustomInner)`

SetCustom sets Custom field to given value.


### GetMajorApplications

`func (o *UpdateNetworkApplianceTrafficShapingVpnExclusions200Response) GetMajorApplications() []UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner`

GetMajorApplications returns the MajorApplications field if non-nil, zero value otherwise.

### GetMajorApplicationsOk

`func (o *UpdateNetworkApplianceTrafficShapingVpnExclusions200Response) GetMajorApplicationsOk() (*[]UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner, bool)`

GetMajorApplicationsOk returns a tuple with the MajorApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorApplications

`func (o *UpdateNetworkApplianceTrafficShapingVpnExclusions200Response) SetMajorApplications(v []UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner)`

SetMajorApplications sets MajorApplications field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


