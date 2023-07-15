# CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **string** | Model of the AP | [optional] 
**Bands** | Pointer to **[]string** | Band to use for each flex radio. For example, [&#39;6&#39;] will set the AP&#39;s first flex radio to 6 GHz | [optional] 

## Methods

### NewCreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner

`func NewCreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner() *CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner`

NewCreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner instantiates a new CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkWirelessRfProfileRequestFlexRadiosByModelInnerWithDefaults

`func NewCreateNetworkWirelessRfProfileRequestFlexRadiosByModelInnerWithDefaults() *CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner`

NewCreateNetworkWirelessRfProfileRequestFlexRadiosByModelInnerWithDefaults instantiates a new CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetBands

`func (o *CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner) GetBands() []string`

GetBands returns the Bands field if non-nil, zero value otherwise.

### GetBandsOk

`func (o *CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner) GetBandsOk() (*[]string, bool)`

GetBandsOk returns a tuple with the Bands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBands

`func (o *CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner) SetBands(v []string)`

SetBands sets Bands field to given value.

### HasBands

`func (o *CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner) HasBands() bool`

HasBands returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


