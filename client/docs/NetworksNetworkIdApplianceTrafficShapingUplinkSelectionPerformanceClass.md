# NetworksNetworkIdApplianceTrafficShapingUplinkSelectionPerformanceClass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of this performance class. Must be one of: &#39;builtin&#39; or &#39;custom&#39; | 
**BuiltinPerformanceClassName** | Pointer to **string** | Name of builtin performance class, must be present when performanceClass type is &#39;builtin&#39;, and value must be one of: &#39;VoIP&#39; | [optional] 
**CustomPerformanceClassId** | Pointer to **string** | ID of created custom performance class, must be present when performanceClass type is &#39;custom&#39; | [optional] 

## Methods

### NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionPerformanceClass

`func NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionPerformanceClass(type_ string, ) *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionPerformanceClass`

NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionPerformanceClass instantiates a new NetworksNetworkIdApplianceTrafficShapingUplinkSelectionPerformanceClass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionPerformanceClassWithDefaults

`func NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionPerformanceClassWithDefaults() *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionPerformanceClass`

NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionPerformanceClassWithDefaults instantiates a new NetworksNetworkIdApplianceTrafficShapingUplinkSelectionPerformanceClass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionPerformanceClass) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionPerformanceClass) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionPerformanceClass) SetType(v string)`

SetType sets Type field to given value.


### GetBuiltinPerformanceClassName

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionPerformanceClass) GetBuiltinPerformanceClassName() string`

GetBuiltinPerformanceClassName returns the BuiltinPerformanceClassName field if non-nil, zero value otherwise.

### GetBuiltinPerformanceClassNameOk

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionPerformanceClass) GetBuiltinPerformanceClassNameOk() (*string, bool)`

GetBuiltinPerformanceClassNameOk returns a tuple with the BuiltinPerformanceClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuiltinPerformanceClassName

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionPerformanceClass) SetBuiltinPerformanceClassName(v string)`

SetBuiltinPerformanceClassName sets BuiltinPerformanceClassName field to given value.

### HasBuiltinPerformanceClassName

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionPerformanceClass) HasBuiltinPerformanceClassName() bool`

HasBuiltinPerformanceClassName returns a boolean if a field has been set.

### GetCustomPerformanceClassId

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionPerformanceClass) GetCustomPerformanceClassId() string`

GetCustomPerformanceClassId returns the CustomPerformanceClassId field if non-nil, zero value otherwise.

### GetCustomPerformanceClassIdOk

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionPerformanceClass) GetCustomPerformanceClassIdOk() (*string, bool)`

GetCustomPerformanceClassIdOk returns a tuple with the CustomPerformanceClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPerformanceClassId

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionPerformanceClass) SetCustomPerformanceClassId(v string)`

SetCustomPerformanceClassId sets CustomPerformanceClassId field to given value.

### HasCustomPerformanceClassId

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionPerformanceClass) HasCustomPerformanceClassId() bool`

HasCustomPerformanceClassId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


