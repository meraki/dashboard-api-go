# CreateNetworkApplianceRfProfileRequestPerSsidSettings1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BandOperationMode** | Pointer to **string** | Choice between &#39;dual&#39;, &#39;2.4ghz&#39;, &#39;5ghz&#39;, &#39;6ghz&#39; or &#39;multi&#39;. | [optional] 
**BandSteeringEnabled** | Pointer to **bool** | Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false. | [optional] 

## Methods

### NewCreateNetworkApplianceRfProfileRequestPerSsidSettings1

`func NewCreateNetworkApplianceRfProfileRequestPerSsidSettings1() *CreateNetworkApplianceRfProfileRequestPerSsidSettings1`

NewCreateNetworkApplianceRfProfileRequestPerSsidSettings1 instantiates a new CreateNetworkApplianceRfProfileRequestPerSsidSettings1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkApplianceRfProfileRequestPerSsidSettings1WithDefaults

`func NewCreateNetworkApplianceRfProfileRequestPerSsidSettings1WithDefaults() *CreateNetworkApplianceRfProfileRequestPerSsidSettings1`

NewCreateNetworkApplianceRfProfileRequestPerSsidSettings1WithDefaults instantiates a new CreateNetworkApplianceRfProfileRequestPerSsidSettings1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandOperationMode

`func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings1) GetBandOperationMode() string`

GetBandOperationMode returns the BandOperationMode field if non-nil, zero value otherwise.

### GetBandOperationModeOk

`func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings1) GetBandOperationModeOk() (*string, bool)`

GetBandOperationModeOk returns a tuple with the BandOperationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandOperationMode

`func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings1) SetBandOperationMode(v string)`

SetBandOperationMode sets BandOperationMode field to given value.

### HasBandOperationMode

`func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings1) HasBandOperationMode() bool`

HasBandOperationMode returns a boolean if a field has been set.

### GetBandSteeringEnabled

`func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings1) GetBandSteeringEnabled() bool`

GetBandSteeringEnabled returns the BandSteeringEnabled field if non-nil, zero value otherwise.

### GetBandSteeringEnabledOk

`func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings1) GetBandSteeringEnabledOk() (*bool, bool)`

GetBandSteeringEnabledOk returns a tuple with the BandSteeringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandSteeringEnabled

`func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings1) SetBandSteeringEnabled(v bool)`

SetBandSteeringEnabled sets BandSteeringEnabled field to given value.

### HasBandSteeringEnabled

`func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings1) HasBandSteeringEnabled() bool`

HasBandSteeringEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


