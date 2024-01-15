# GetOrganizationConfigurationChanges200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | Pointer to **time.Time** | Time, in ISO8601 format, when the configuration change was made. | [optional] 
**AdminName** | Pointer to **string** | The name of the admin who made the configuration change. | [optional] 
**AdminEmail** | Pointer to **string** | The email address of the admin who made the configuration change. This attribute may be null. | [optional] 
**AdminId** | Pointer to **string** | The ID of the admin who made the configuration change. This attribute may be null. | [optional] 
**NetworkName** | Pointer to **string** | The name of the network that the configuration change was applied to. This attribute may be null. | [optional] 
**NetworkId** | Pointer to **string** | The ID of the network that the configuration change was applied to. This attribute may be null. | [optional] 
**NetworkUrl** | Pointer to **string** | The url of the network that the configuration change was applied to. This attribute may be null. | [optional] 
**SsidName** | Pointer to **string** | The name of the ssid that the configuration change was applied to, if applicable. This attribute may be null. | [optional] 
**SsidNumber** | Pointer to **int32** | The ssid number that the configuration change was applied to, if applicable. This attribute may be null. | [optional] 
**Page** | Pointer to **string** | The name of the Meraki Dashboard page on which the configuration change was made. | [optional] 
**Label** | Pointer to **string** | Description of the configuration change. | [optional] 
**OldValue** | Pointer to **string** | The value of the configuration, before the change was applied. | [optional] 
**NewValue** | Pointer to **string** | The value of the configuration, after the change was applied. | [optional] 

## Methods

### NewGetOrganizationConfigurationChanges200ResponseInner

`func NewGetOrganizationConfigurationChanges200ResponseInner() *GetOrganizationConfigurationChanges200ResponseInner`

NewGetOrganizationConfigurationChanges200ResponseInner instantiates a new GetOrganizationConfigurationChanges200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationConfigurationChanges200ResponseInnerWithDefaults

`func NewGetOrganizationConfigurationChanges200ResponseInnerWithDefaults() *GetOrganizationConfigurationChanges200ResponseInner`

NewGetOrganizationConfigurationChanges200ResponseInnerWithDefaults instantiates a new GetOrganizationConfigurationChanges200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *GetOrganizationConfigurationChanges200ResponseInner) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetOrganizationConfigurationChanges200ResponseInner) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetOrganizationConfigurationChanges200ResponseInner) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetOrganizationConfigurationChanges200ResponseInner) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetAdminName

`func (o *GetOrganizationConfigurationChanges200ResponseInner) GetAdminName() string`

GetAdminName returns the AdminName field if non-nil, zero value otherwise.

### GetAdminNameOk

`func (o *GetOrganizationConfigurationChanges200ResponseInner) GetAdminNameOk() (*string, bool)`

GetAdminNameOk returns a tuple with the AdminName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminName

`func (o *GetOrganizationConfigurationChanges200ResponseInner) SetAdminName(v string)`

SetAdminName sets AdminName field to given value.

### HasAdminName

`func (o *GetOrganizationConfigurationChanges200ResponseInner) HasAdminName() bool`

HasAdminName returns a boolean if a field has been set.

### GetAdminEmail

`func (o *GetOrganizationConfigurationChanges200ResponseInner) GetAdminEmail() string`

GetAdminEmail returns the AdminEmail field if non-nil, zero value otherwise.

### GetAdminEmailOk

`func (o *GetOrganizationConfigurationChanges200ResponseInner) GetAdminEmailOk() (*string, bool)`

GetAdminEmailOk returns a tuple with the AdminEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminEmail

`func (o *GetOrganizationConfigurationChanges200ResponseInner) SetAdminEmail(v string)`

SetAdminEmail sets AdminEmail field to given value.

### HasAdminEmail

`func (o *GetOrganizationConfigurationChanges200ResponseInner) HasAdminEmail() bool`

HasAdminEmail returns a boolean if a field has been set.

### GetAdminId

`func (o *GetOrganizationConfigurationChanges200ResponseInner) GetAdminId() string`

GetAdminId returns the AdminId field if non-nil, zero value otherwise.

### GetAdminIdOk

`func (o *GetOrganizationConfigurationChanges200ResponseInner) GetAdminIdOk() (*string, bool)`

GetAdminIdOk returns a tuple with the AdminId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminId

`func (o *GetOrganizationConfigurationChanges200ResponseInner) SetAdminId(v string)`

SetAdminId sets AdminId field to given value.

### HasAdminId

`func (o *GetOrganizationConfigurationChanges200ResponseInner) HasAdminId() bool`

HasAdminId returns a boolean if a field has been set.

### GetNetworkName

`func (o *GetOrganizationConfigurationChanges200ResponseInner) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *GetOrganizationConfigurationChanges200ResponseInner) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *GetOrganizationConfigurationChanges200ResponseInner) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *GetOrganizationConfigurationChanges200ResponseInner) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

### GetNetworkId

`func (o *GetOrganizationConfigurationChanges200ResponseInner) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *GetOrganizationConfigurationChanges200ResponseInner) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *GetOrganizationConfigurationChanges200ResponseInner) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *GetOrganizationConfigurationChanges200ResponseInner) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetNetworkUrl

`func (o *GetOrganizationConfigurationChanges200ResponseInner) GetNetworkUrl() string`

GetNetworkUrl returns the NetworkUrl field if non-nil, zero value otherwise.

### GetNetworkUrlOk

`func (o *GetOrganizationConfigurationChanges200ResponseInner) GetNetworkUrlOk() (*string, bool)`

GetNetworkUrlOk returns a tuple with the NetworkUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkUrl

`func (o *GetOrganizationConfigurationChanges200ResponseInner) SetNetworkUrl(v string)`

SetNetworkUrl sets NetworkUrl field to given value.

### HasNetworkUrl

`func (o *GetOrganizationConfigurationChanges200ResponseInner) HasNetworkUrl() bool`

HasNetworkUrl returns a boolean if a field has been set.

### GetSsidName

`func (o *GetOrganizationConfigurationChanges200ResponseInner) GetSsidName() string`

GetSsidName returns the SsidName field if non-nil, zero value otherwise.

### GetSsidNameOk

`func (o *GetOrganizationConfigurationChanges200ResponseInner) GetSsidNameOk() (*string, bool)`

GetSsidNameOk returns a tuple with the SsidName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidName

`func (o *GetOrganizationConfigurationChanges200ResponseInner) SetSsidName(v string)`

SetSsidName sets SsidName field to given value.

### HasSsidName

`func (o *GetOrganizationConfigurationChanges200ResponseInner) HasSsidName() bool`

HasSsidName returns a boolean if a field has been set.

### GetSsidNumber

`func (o *GetOrganizationConfigurationChanges200ResponseInner) GetSsidNumber() int32`

GetSsidNumber returns the SsidNumber field if non-nil, zero value otherwise.

### GetSsidNumberOk

`func (o *GetOrganizationConfigurationChanges200ResponseInner) GetSsidNumberOk() (*int32, bool)`

GetSsidNumberOk returns a tuple with the SsidNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidNumber

`func (o *GetOrganizationConfigurationChanges200ResponseInner) SetSsidNumber(v int32)`

SetSsidNumber sets SsidNumber field to given value.

### HasSsidNumber

`func (o *GetOrganizationConfigurationChanges200ResponseInner) HasSsidNumber() bool`

HasSsidNumber returns a boolean if a field has been set.

### GetPage

`func (o *GetOrganizationConfigurationChanges200ResponseInner) GetPage() string`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetOrganizationConfigurationChanges200ResponseInner) GetPageOk() (*string, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetOrganizationConfigurationChanges200ResponseInner) SetPage(v string)`

SetPage sets Page field to given value.

### HasPage

`func (o *GetOrganizationConfigurationChanges200ResponseInner) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetLabel

`func (o *GetOrganizationConfigurationChanges200ResponseInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetOrganizationConfigurationChanges200ResponseInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetOrganizationConfigurationChanges200ResponseInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetOrganizationConfigurationChanges200ResponseInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetOldValue

`func (o *GetOrganizationConfigurationChanges200ResponseInner) GetOldValue() string`

GetOldValue returns the OldValue field if non-nil, zero value otherwise.

### GetOldValueOk

`func (o *GetOrganizationConfigurationChanges200ResponseInner) GetOldValueOk() (*string, bool)`

GetOldValueOk returns a tuple with the OldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldValue

`func (o *GetOrganizationConfigurationChanges200ResponseInner) SetOldValue(v string)`

SetOldValue sets OldValue field to given value.

### HasOldValue

`func (o *GetOrganizationConfigurationChanges200ResponseInner) HasOldValue() bool`

HasOldValue returns a boolean if a field has been set.

### GetNewValue

`func (o *GetOrganizationConfigurationChanges200ResponseInner) GetNewValue() string`

GetNewValue returns the NewValue field if non-nil, zero value otherwise.

### GetNewValueOk

`func (o *GetOrganizationConfigurationChanges200ResponseInner) GetNewValueOk() (*string, bool)`

GetNewValueOk returns a tuple with the NewValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewValue

`func (o *GetOrganizationConfigurationChanges200ResponseInner) SetNewValue(v string)`

SetNewValue sets NewValue field to given value.

### HasNewValue

`func (o *GetOrganizationConfigurationChanges200ResponseInner) HasNewValue() bool`

HasNewValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


