# CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the custom performance class | 
**MaxLatency** | Pointer to **int32** | Maximum latency in milliseconds | [optional] 
**MaxJitter** | Pointer to **int32** | Maximum jitter in milliseconds | [optional] 
**MaxLossPercentage** | Pointer to **int32** | Maximum percentage of packet loss | [optional] 

## Methods

### NewCreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest

`func NewCreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest(name string, ) *CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest`

NewCreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest instantiates a new CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkApplianceTrafficShapingCustomPerformanceClassRequestWithDefaults

`func NewCreateNetworkApplianceTrafficShapingCustomPerformanceClassRequestWithDefaults() *CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest`

NewCreateNetworkApplianceTrafficShapingCustomPerformanceClassRequestWithDefaults instantiates a new CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) SetName(v string)`

SetName sets Name field to given value.


### GetMaxLatency

`func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) GetMaxLatency() int32`

GetMaxLatency returns the MaxLatency field if non-nil, zero value otherwise.

### GetMaxLatencyOk

`func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) GetMaxLatencyOk() (*int32, bool)`

GetMaxLatencyOk returns a tuple with the MaxLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLatency

`func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) SetMaxLatency(v int32)`

SetMaxLatency sets MaxLatency field to given value.

### HasMaxLatency

`func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) HasMaxLatency() bool`

HasMaxLatency returns a boolean if a field has been set.

### GetMaxJitter

`func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) GetMaxJitter() int32`

GetMaxJitter returns the MaxJitter field if non-nil, zero value otherwise.

### GetMaxJitterOk

`func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) GetMaxJitterOk() (*int32, bool)`

GetMaxJitterOk returns a tuple with the MaxJitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxJitter

`func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) SetMaxJitter(v int32)`

SetMaxJitter sets MaxJitter field to given value.

### HasMaxJitter

`func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) HasMaxJitter() bool`

HasMaxJitter returns a boolean if a field has been set.

### GetMaxLossPercentage

`func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) GetMaxLossPercentage() int32`

GetMaxLossPercentage returns the MaxLossPercentage field if non-nil, zero value otherwise.

### GetMaxLossPercentageOk

`func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) GetMaxLossPercentageOk() (*int32, bool)`

GetMaxLossPercentageOk returns a tuple with the MaxLossPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLossPercentage

`func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) SetMaxLossPercentage(v int32)`

SetMaxLossPercentage sets MaxLossPercentage field to given value.

### HasMaxLossPercentage

`func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) HasMaxLossPercentage() bool`

HasMaxLossPercentage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


