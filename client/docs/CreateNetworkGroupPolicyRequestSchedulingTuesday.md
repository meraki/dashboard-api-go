# CreateNetworkGroupPolicyRequestSchedulingTuesday

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Whether the schedule is active (true) or inactive (false) during the time specified between &#39;from&#39; and &#39;to&#39;. Defaults to true. | [optional] 
**From** | Pointer to **string** | The time, from &#39;00:00&#39; to &#39;24:00&#39;. Must be less than the time specified in &#39;to&#39;. Defaults to &#39;00:00&#39;. Only 30 minute increments are allowed. | [optional] 
**To** | Pointer to **string** | The time, from &#39;00:00&#39; to &#39;24:00&#39;. Must be greater than the time specified in &#39;from&#39;. Defaults to &#39;24:00&#39;. Only 30 minute increments are allowed. | [optional] 

## Methods

### NewCreateNetworkGroupPolicyRequestSchedulingTuesday

`func NewCreateNetworkGroupPolicyRequestSchedulingTuesday() *CreateNetworkGroupPolicyRequestSchedulingTuesday`

NewCreateNetworkGroupPolicyRequestSchedulingTuesday instantiates a new CreateNetworkGroupPolicyRequestSchedulingTuesday object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkGroupPolicyRequestSchedulingTuesdayWithDefaults

`func NewCreateNetworkGroupPolicyRequestSchedulingTuesdayWithDefaults() *CreateNetworkGroupPolicyRequestSchedulingTuesday`

NewCreateNetworkGroupPolicyRequestSchedulingTuesdayWithDefaults instantiates a new CreateNetworkGroupPolicyRequestSchedulingTuesday object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *CreateNetworkGroupPolicyRequestSchedulingTuesday) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CreateNetworkGroupPolicyRequestSchedulingTuesday) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CreateNetworkGroupPolicyRequestSchedulingTuesday) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CreateNetworkGroupPolicyRequestSchedulingTuesday) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetFrom

`func (o *CreateNetworkGroupPolicyRequestSchedulingTuesday) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CreateNetworkGroupPolicyRequestSchedulingTuesday) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CreateNetworkGroupPolicyRequestSchedulingTuesday) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *CreateNetworkGroupPolicyRequestSchedulingTuesday) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *CreateNetworkGroupPolicyRequestSchedulingTuesday) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CreateNetworkGroupPolicyRequestSchedulingTuesday) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CreateNetworkGroupPolicyRequestSchedulingTuesday) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *CreateNetworkGroupPolicyRequestSchedulingTuesday) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


