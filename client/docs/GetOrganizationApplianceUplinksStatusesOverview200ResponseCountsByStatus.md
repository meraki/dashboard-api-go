# GetOrganizationApplianceUplinksStatusesOverview200ResponseCountsByStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **int32** | number of uplinks that are active and working | [optional] 
**Ready** | Pointer to **int32** | number of uplinks that are working but on standby | [optional] 
**Failed** | Pointer to **int32** | number of uplinks that were working but have failed | [optional] 
**Connecting** | Pointer to **int32** | number of uplinks currently connecting | [optional] 
**NotConnected** | Pointer to **int32** | number of uplinks currently where nothing is plugged in | [optional] 

## Methods

### NewGetOrganizationApplianceUplinksStatusesOverview200ResponseCountsByStatus

`func NewGetOrganizationApplianceUplinksStatusesOverview200ResponseCountsByStatus() *GetOrganizationApplianceUplinksStatusesOverview200ResponseCountsByStatus`

NewGetOrganizationApplianceUplinksStatusesOverview200ResponseCountsByStatus instantiates a new GetOrganizationApplianceUplinksStatusesOverview200ResponseCountsByStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationApplianceUplinksStatusesOverview200ResponseCountsByStatusWithDefaults

`func NewGetOrganizationApplianceUplinksStatusesOverview200ResponseCountsByStatusWithDefaults() *GetOrganizationApplianceUplinksStatusesOverview200ResponseCountsByStatus`

NewGetOrganizationApplianceUplinksStatusesOverview200ResponseCountsByStatusWithDefaults instantiates a new GetOrganizationApplianceUplinksStatusesOverview200ResponseCountsByStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *GetOrganizationApplianceUplinksStatusesOverview200ResponseCountsByStatus) GetActive() int32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetOrganizationApplianceUplinksStatusesOverview200ResponseCountsByStatus) GetActiveOk() (*int32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetOrganizationApplianceUplinksStatusesOverview200ResponseCountsByStatus) SetActive(v int32)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetOrganizationApplianceUplinksStatusesOverview200ResponseCountsByStatus) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetReady

`func (o *GetOrganizationApplianceUplinksStatusesOverview200ResponseCountsByStatus) GetReady() int32`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *GetOrganizationApplianceUplinksStatusesOverview200ResponseCountsByStatus) GetReadyOk() (*int32, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *GetOrganizationApplianceUplinksStatusesOverview200ResponseCountsByStatus) SetReady(v int32)`

SetReady sets Ready field to given value.

### HasReady

`func (o *GetOrganizationApplianceUplinksStatusesOverview200ResponseCountsByStatus) HasReady() bool`

HasReady returns a boolean if a field has been set.

### GetFailed

`func (o *GetOrganizationApplianceUplinksStatusesOverview200ResponseCountsByStatus) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *GetOrganizationApplianceUplinksStatusesOverview200ResponseCountsByStatus) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *GetOrganizationApplianceUplinksStatusesOverview200ResponseCountsByStatus) SetFailed(v int32)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *GetOrganizationApplianceUplinksStatusesOverview200ResponseCountsByStatus) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetConnecting

`func (o *GetOrganizationApplianceUplinksStatusesOverview200ResponseCountsByStatus) GetConnecting() int32`

GetConnecting returns the Connecting field if non-nil, zero value otherwise.

### GetConnectingOk

`func (o *GetOrganizationApplianceUplinksStatusesOverview200ResponseCountsByStatus) GetConnectingOk() (*int32, bool)`

GetConnectingOk returns a tuple with the Connecting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnecting

`func (o *GetOrganizationApplianceUplinksStatusesOverview200ResponseCountsByStatus) SetConnecting(v int32)`

SetConnecting sets Connecting field to given value.

### HasConnecting

`func (o *GetOrganizationApplianceUplinksStatusesOverview200ResponseCountsByStatus) HasConnecting() bool`

HasConnecting returns a boolean if a field has been set.

### GetNotConnected

`func (o *GetOrganizationApplianceUplinksStatusesOverview200ResponseCountsByStatus) GetNotConnected() int32`

GetNotConnected returns the NotConnected field if non-nil, zero value otherwise.

### GetNotConnectedOk

`func (o *GetOrganizationApplianceUplinksStatusesOverview200ResponseCountsByStatus) GetNotConnectedOk() (*int32, bool)`

GetNotConnectedOk returns a tuple with the NotConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotConnected

`func (o *GetOrganizationApplianceUplinksStatusesOverview200ResponseCountsByStatus) SetNotConnected(v int32)`

SetNotConnected sets NotConnected field to given value.

### HasNotConnected

`func (o *GetOrganizationApplianceUplinksStatusesOverview200ResponseCountsByStatus) HasNotConnected() bool`

HasNotConnected returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


