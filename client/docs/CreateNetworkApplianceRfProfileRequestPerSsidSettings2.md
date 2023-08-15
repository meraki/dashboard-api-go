# CreateNetworkApplianceRfProfileRequestPerSsidSettings2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BandOperationMode** | Pointer to **string** | Choice between &#39;dual&#39;, &#39;2.4ghz&#39;, &#39;5ghz&#39;, &#39;6ghz&#39; or &#39;multi&#39;. | [optional] 
**BandSteeringEnabled** | Pointer to **bool** | Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false. | [optional] 

## Methods

### NewCreateNetworkApplianceRfProfileRequestPerSsidSettings2

`func NewCreateNetworkApplianceRfProfileRequestPerSsidSettings2() *CreateNetworkApplianceRfProfileRequestPerSsidSettings2`

NewCreateNetworkApplianceRfProfileRequestPerSsidSettings2 instantiates a new CreateNetworkApplianceRfProfileRequestPerSsidSettings2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkApplianceRfProfileRequestPerSsidSettings2WithDefaults

`func NewCreateNetworkApplianceRfProfileRequestPerSsidSettings2WithDefaults() *CreateNetworkApplianceRfProfileRequestPerSsidSettings2`

NewCreateNetworkApplianceRfProfileRequestPerSsidSettings2WithDefaults instantiates a new CreateNetworkApplianceRfProfileRequestPerSsidSettings2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandOperationMode

`func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings2) GetBandOperationMode() string`

GetBandOperationMode returns the BandOperationMode field if non-nil, zero value otherwise.

### GetBandOperationModeOk

`func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings2) GetBandOperationModeOk() (*string, bool)`

GetBandOperationModeOk returns a tuple with the BandOperationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandOperationMode

`func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings2) SetBandOperationMode(v string)`

SetBandOperationMode sets BandOperationMode field to given value.

### HasBandOperationMode

`func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings2) HasBandOperationMode() bool`

HasBandOperationMode returns a boolean if a field has been set.

### GetBandSteeringEnabled

`func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings2) GetBandSteeringEnabled() bool`

GetBandSteeringEnabled returns the BandSteeringEnabled field if non-nil, zero value otherwise.

### GetBandSteeringEnabledOk

`func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings2) GetBandSteeringEnabledOk() (*bool, bool)`

GetBandSteeringEnabledOk returns a tuple with the BandSteeringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandSteeringEnabled

`func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings2) SetBandSteeringEnabled(v bool)`

SetBandSteeringEnabled sets BandSteeringEnabled field to given value.

### HasBandSteeringEnabled

`func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings2) HasBandSteeringEnabled() bool`

HasBandSteeringEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


