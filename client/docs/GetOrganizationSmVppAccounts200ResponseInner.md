# GetOrganizationSmVppAccounts200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VppAccountId** | Pointer to **string** | The id of the VPP Account | [optional] 
**ContentToken** | Pointer to **string** | The VPP service token | [optional] 
**Email** | Pointer to **string** | The email address associated with the VPP account | [optional] 
**Name** | Pointer to **string** | The name of the VPP account | [optional] 
**AllowedAdmins** | Pointer to **string** | The allowed admins for the VPP account | [optional] 
**NetworkIdAdmins** | Pointer to **string** | The network IDs of the admins for the VPP account | [optional] 
**AssignableNetworks** | Pointer to **string** | The assignable networks for the VPP account | [optional] 
**AssignableNetworkIds** | Pointer to **[]string** | The network IDs of the assignable networks for the VPP account | [optional] 
**VppLocationId** | Pointer to **string** | The VPP location ID | [optional] 
**VppLocationName** | Pointer to **string** | The VPP location name | [optional] 
**LastSyncedAt** | Pointer to **string** | The last time the VPP account was synced | [optional] 
**LastForceSyncedAt** | Pointer to **string** | The last time the VPP account was force synced | [optional] 
**ParsedToken** | Pointer to [**GetOrganizationSmVppAccounts200ResponseInnerParsedToken**](GetOrganizationSmVppAccounts200ResponseInnerParsedToken.md) |  | [optional] 
**Id** | Pointer to **string** | The id of the VPP Account | [optional] 
**VppServiceToken** | Pointer to **string** | The VPP Account&#39;s Service Token | [optional] 

## Methods

### NewGetOrganizationSmVppAccounts200ResponseInner

`func NewGetOrganizationSmVppAccounts200ResponseInner() *GetOrganizationSmVppAccounts200ResponseInner`

NewGetOrganizationSmVppAccounts200ResponseInner instantiates a new GetOrganizationSmVppAccounts200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationSmVppAccounts200ResponseInnerWithDefaults

`func NewGetOrganizationSmVppAccounts200ResponseInnerWithDefaults() *GetOrganizationSmVppAccounts200ResponseInner`

NewGetOrganizationSmVppAccounts200ResponseInnerWithDefaults instantiates a new GetOrganizationSmVppAccounts200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVppAccountId

`func (o *GetOrganizationSmVppAccounts200ResponseInner) GetVppAccountId() string`

GetVppAccountId returns the VppAccountId field if non-nil, zero value otherwise.

### GetVppAccountIdOk

`func (o *GetOrganizationSmVppAccounts200ResponseInner) GetVppAccountIdOk() (*string, bool)`

GetVppAccountIdOk returns a tuple with the VppAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVppAccountId

`func (o *GetOrganizationSmVppAccounts200ResponseInner) SetVppAccountId(v string)`

SetVppAccountId sets VppAccountId field to given value.

### HasVppAccountId

`func (o *GetOrganizationSmVppAccounts200ResponseInner) HasVppAccountId() bool`

HasVppAccountId returns a boolean if a field has been set.

### GetContentToken

`func (o *GetOrganizationSmVppAccounts200ResponseInner) GetContentToken() string`

GetContentToken returns the ContentToken field if non-nil, zero value otherwise.

### GetContentTokenOk

`func (o *GetOrganizationSmVppAccounts200ResponseInner) GetContentTokenOk() (*string, bool)`

GetContentTokenOk returns a tuple with the ContentToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentToken

`func (o *GetOrganizationSmVppAccounts200ResponseInner) SetContentToken(v string)`

SetContentToken sets ContentToken field to given value.

### HasContentToken

`func (o *GetOrganizationSmVppAccounts200ResponseInner) HasContentToken() bool`

HasContentToken returns a boolean if a field has been set.

### GetEmail

`func (o *GetOrganizationSmVppAccounts200ResponseInner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetOrganizationSmVppAccounts200ResponseInner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetOrganizationSmVppAccounts200ResponseInner) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetOrganizationSmVppAccounts200ResponseInner) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *GetOrganizationSmVppAccounts200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetOrganizationSmVppAccounts200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetOrganizationSmVppAccounts200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetOrganizationSmVppAccounts200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAllowedAdmins

`func (o *GetOrganizationSmVppAccounts200ResponseInner) GetAllowedAdmins() string`

GetAllowedAdmins returns the AllowedAdmins field if non-nil, zero value otherwise.

### GetAllowedAdminsOk

`func (o *GetOrganizationSmVppAccounts200ResponseInner) GetAllowedAdminsOk() (*string, bool)`

GetAllowedAdminsOk returns a tuple with the AllowedAdmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAdmins

`func (o *GetOrganizationSmVppAccounts200ResponseInner) SetAllowedAdmins(v string)`

SetAllowedAdmins sets AllowedAdmins field to given value.

### HasAllowedAdmins

`func (o *GetOrganizationSmVppAccounts200ResponseInner) HasAllowedAdmins() bool`

HasAllowedAdmins returns a boolean if a field has been set.

### GetNetworkIdAdmins

`func (o *GetOrganizationSmVppAccounts200ResponseInner) GetNetworkIdAdmins() string`

GetNetworkIdAdmins returns the NetworkIdAdmins field if non-nil, zero value otherwise.

### GetNetworkIdAdminsOk

`func (o *GetOrganizationSmVppAccounts200ResponseInner) GetNetworkIdAdminsOk() (*string, bool)`

GetNetworkIdAdminsOk returns a tuple with the NetworkIdAdmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIdAdmins

`func (o *GetOrganizationSmVppAccounts200ResponseInner) SetNetworkIdAdmins(v string)`

SetNetworkIdAdmins sets NetworkIdAdmins field to given value.

### HasNetworkIdAdmins

`func (o *GetOrganizationSmVppAccounts200ResponseInner) HasNetworkIdAdmins() bool`

HasNetworkIdAdmins returns a boolean if a field has been set.

### GetAssignableNetworks

`func (o *GetOrganizationSmVppAccounts200ResponseInner) GetAssignableNetworks() string`

GetAssignableNetworks returns the AssignableNetworks field if non-nil, zero value otherwise.

### GetAssignableNetworksOk

`func (o *GetOrganizationSmVppAccounts200ResponseInner) GetAssignableNetworksOk() (*string, bool)`

GetAssignableNetworksOk returns a tuple with the AssignableNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignableNetworks

`func (o *GetOrganizationSmVppAccounts200ResponseInner) SetAssignableNetworks(v string)`

SetAssignableNetworks sets AssignableNetworks field to given value.

### HasAssignableNetworks

`func (o *GetOrganizationSmVppAccounts200ResponseInner) HasAssignableNetworks() bool`

HasAssignableNetworks returns a boolean if a field has been set.

### GetAssignableNetworkIds

`func (o *GetOrganizationSmVppAccounts200ResponseInner) GetAssignableNetworkIds() []string`

GetAssignableNetworkIds returns the AssignableNetworkIds field if non-nil, zero value otherwise.

### GetAssignableNetworkIdsOk

`func (o *GetOrganizationSmVppAccounts200ResponseInner) GetAssignableNetworkIdsOk() (*[]string, bool)`

GetAssignableNetworkIdsOk returns a tuple with the AssignableNetworkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignableNetworkIds

`func (o *GetOrganizationSmVppAccounts200ResponseInner) SetAssignableNetworkIds(v []string)`

SetAssignableNetworkIds sets AssignableNetworkIds field to given value.

### HasAssignableNetworkIds

`func (o *GetOrganizationSmVppAccounts200ResponseInner) HasAssignableNetworkIds() bool`

HasAssignableNetworkIds returns a boolean if a field has been set.

### GetVppLocationId

`func (o *GetOrganizationSmVppAccounts200ResponseInner) GetVppLocationId() string`

GetVppLocationId returns the VppLocationId field if non-nil, zero value otherwise.

### GetVppLocationIdOk

`func (o *GetOrganizationSmVppAccounts200ResponseInner) GetVppLocationIdOk() (*string, bool)`

GetVppLocationIdOk returns a tuple with the VppLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVppLocationId

`func (o *GetOrganizationSmVppAccounts200ResponseInner) SetVppLocationId(v string)`

SetVppLocationId sets VppLocationId field to given value.

### HasVppLocationId

`func (o *GetOrganizationSmVppAccounts200ResponseInner) HasVppLocationId() bool`

HasVppLocationId returns a boolean if a field has been set.

### GetVppLocationName

`func (o *GetOrganizationSmVppAccounts200ResponseInner) GetVppLocationName() string`

GetVppLocationName returns the VppLocationName field if non-nil, zero value otherwise.

### GetVppLocationNameOk

`func (o *GetOrganizationSmVppAccounts200ResponseInner) GetVppLocationNameOk() (*string, bool)`

GetVppLocationNameOk returns a tuple with the VppLocationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVppLocationName

`func (o *GetOrganizationSmVppAccounts200ResponseInner) SetVppLocationName(v string)`

SetVppLocationName sets VppLocationName field to given value.

### HasVppLocationName

`func (o *GetOrganizationSmVppAccounts200ResponseInner) HasVppLocationName() bool`

HasVppLocationName returns a boolean if a field has been set.

### GetLastSyncedAt

`func (o *GetOrganizationSmVppAccounts200ResponseInner) GetLastSyncedAt() string`

GetLastSyncedAt returns the LastSyncedAt field if non-nil, zero value otherwise.

### GetLastSyncedAtOk

`func (o *GetOrganizationSmVppAccounts200ResponseInner) GetLastSyncedAtOk() (*string, bool)`

GetLastSyncedAtOk returns a tuple with the LastSyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncedAt

`func (o *GetOrganizationSmVppAccounts200ResponseInner) SetLastSyncedAt(v string)`

SetLastSyncedAt sets LastSyncedAt field to given value.

### HasLastSyncedAt

`func (o *GetOrganizationSmVppAccounts200ResponseInner) HasLastSyncedAt() bool`

HasLastSyncedAt returns a boolean if a field has been set.

### GetLastForceSyncedAt

`func (o *GetOrganizationSmVppAccounts200ResponseInner) GetLastForceSyncedAt() string`

GetLastForceSyncedAt returns the LastForceSyncedAt field if non-nil, zero value otherwise.

### GetLastForceSyncedAtOk

`func (o *GetOrganizationSmVppAccounts200ResponseInner) GetLastForceSyncedAtOk() (*string, bool)`

GetLastForceSyncedAtOk returns a tuple with the LastForceSyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastForceSyncedAt

`func (o *GetOrganizationSmVppAccounts200ResponseInner) SetLastForceSyncedAt(v string)`

SetLastForceSyncedAt sets LastForceSyncedAt field to given value.

### HasLastForceSyncedAt

`func (o *GetOrganizationSmVppAccounts200ResponseInner) HasLastForceSyncedAt() bool`

HasLastForceSyncedAt returns a boolean if a field has been set.

### GetParsedToken

`func (o *GetOrganizationSmVppAccounts200ResponseInner) GetParsedToken() GetOrganizationSmVppAccounts200ResponseInnerParsedToken`

GetParsedToken returns the ParsedToken field if non-nil, zero value otherwise.

### GetParsedTokenOk

`func (o *GetOrganizationSmVppAccounts200ResponseInner) GetParsedTokenOk() (*GetOrganizationSmVppAccounts200ResponseInnerParsedToken, bool)`

GetParsedTokenOk returns a tuple with the ParsedToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParsedToken

`func (o *GetOrganizationSmVppAccounts200ResponseInner) SetParsedToken(v GetOrganizationSmVppAccounts200ResponseInnerParsedToken)`

SetParsedToken sets ParsedToken field to given value.

### HasParsedToken

`func (o *GetOrganizationSmVppAccounts200ResponseInner) HasParsedToken() bool`

HasParsedToken returns a boolean if a field has been set.

### GetId

`func (o *GetOrganizationSmVppAccounts200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetOrganizationSmVppAccounts200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetOrganizationSmVppAccounts200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetOrganizationSmVppAccounts200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVppServiceToken

`func (o *GetOrganizationSmVppAccounts200ResponseInner) GetVppServiceToken() string`

GetVppServiceToken returns the VppServiceToken field if non-nil, zero value otherwise.

### GetVppServiceTokenOk

`func (o *GetOrganizationSmVppAccounts200ResponseInner) GetVppServiceTokenOk() (*string, bool)`

GetVppServiceTokenOk returns a tuple with the VppServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVppServiceToken

`func (o *GetOrganizationSmVppAccounts200ResponseInner) SetVppServiceToken(v string)`

SetVppServiceToken sets VppServiceToken field to given value.

### HasVppServiceToken

`func (o *GetOrganizationSmVppAccounts200ResponseInner) HasVppServiceToken() bool`

HasVppServiceToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


