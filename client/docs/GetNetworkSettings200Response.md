# GetNetworkSettings200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalStatusPageEnabled** | Pointer to **bool** | Enables / disables the local device status pages (&lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;http://my.meraki.com/&#39;&gt;my.meraki.com, &lt;/a&gt;&lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;http://ap.meraki.com/&#39;&gt;ap.meraki.com, &lt;/a&gt;&lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;http://switch.meraki.com/&#39;&gt;switch.meraki.com, &lt;/a&gt;&lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;http://wired.meraki.com/&#39;&gt;wired.meraki.com&lt;/a&gt;). Optional (defaults to false) | [optional] 
**RemoteStatusPageEnabled** | Pointer to **bool** | Enables / disables access to the device status page (&lt;a target&#x3D;&#39;_blank&#39;&gt;http://[device&#39;s LAN IP])&lt;/a&gt;. Optional. Can only be set if localStatusPageEnabled is set to true | [optional] 
**LocalStatusPage** | Pointer to [**GetNetworkSettings200ResponseLocalStatusPage**](GetNetworkSettings200ResponseLocalStatusPage.md) |  | [optional] 
**SecurePort** | Pointer to [**GetNetworkSettings200ResponseSecurePort**](GetNetworkSettings200ResponseSecurePort.md) |  | [optional] 
**Fips** | Pointer to [**GetNetworkSettings200ResponseFips**](GetNetworkSettings200ResponseFips.md) |  | [optional] 
**NamedVlans** | Pointer to [**GetNetworkSettings200ResponseNamedVlans**](GetNetworkSettings200ResponseNamedVlans.md) |  | [optional] 
**ClientPrivacy** | Pointer to [**GetNetworkSettings200ResponseClientPrivacy**](GetNetworkSettings200ResponseClientPrivacy.md) |  | [optional] 

## Methods

### NewGetNetworkSettings200Response

`func NewGetNetworkSettings200Response() *GetNetworkSettings200Response`

NewGetNetworkSettings200Response instantiates a new GetNetworkSettings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSettings200ResponseWithDefaults

`func NewGetNetworkSettings200ResponseWithDefaults() *GetNetworkSettings200Response`

NewGetNetworkSettings200ResponseWithDefaults instantiates a new GetNetworkSettings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalStatusPageEnabled

`func (o *GetNetworkSettings200Response) GetLocalStatusPageEnabled() bool`

GetLocalStatusPageEnabled returns the LocalStatusPageEnabled field if non-nil, zero value otherwise.

### GetLocalStatusPageEnabledOk

`func (o *GetNetworkSettings200Response) GetLocalStatusPageEnabledOk() (*bool, bool)`

GetLocalStatusPageEnabledOk returns a tuple with the LocalStatusPageEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalStatusPageEnabled

`func (o *GetNetworkSettings200Response) SetLocalStatusPageEnabled(v bool)`

SetLocalStatusPageEnabled sets LocalStatusPageEnabled field to given value.

### HasLocalStatusPageEnabled

`func (o *GetNetworkSettings200Response) HasLocalStatusPageEnabled() bool`

HasLocalStatusPageEnabled returns a boolean if a field has been set.

### GetRemoteStatusPageEnabled

`func (o *GetNetworkSettings200Response) GetRemoteStatusPageEnabled() bool`

GetRemoteStatusPageEnabled returns the RemoteStatusPageEnabled field if non-nil, zero value otherwise.

### GetRemoteStatusPageEnabledOk

`func (o *GetNetworkSettings200Response) GetRemoteStatusPageEnabledOk() (*bool, bool)`

GetRemoteStatusPageEnabledOk returns a tuple with the RemoteStatusPageEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteStatusPageEnabled

`func (o *GetNetworkSettings200Response) SetRemoteStatusPageEnabled(v bool)`

SetRemoteStatusPageEnabled sets RemoteStatusPageEnabled field to given value.

### HasRemoteStatusPageEnabled

`func (o *GetNetworkSettings200Response) HasRemoteStatusPageEnabled() bool`

HasRemoteStatusPageEnabled returns a boolean if a field has been set.

### GetLocalStatusPage

`func (o *GetNetworkSettings200Response) GetLocalStatusPage() GetNetworkSettings200ResponseLocalStatusPage`

GetLocalStatusPage returns the LocalStatusPage field if non-nil, zero value otherwise.

### GetLocalStatusPageOk

`func (o *GetNetworkSettings200Response) GetLocalStatusPageOk() (*GetNetworkSettings200ResponseLocalStatusPage, bool)`

GetLocalStatusPageOk returns a tuple with the LocalStatusPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalStatusPage

`func (o *GetNetworkSettings200Response) SetLocalStatusPage(v GetNetworkSettings200ResponseLocalStatusPage)`

SetLocalStatusPage sets LocalStatusPage field to given value.

### HasLocalStatusPage

`func (o *GetNetworkSettings200Response) HasLocalStatusPage() bool`

HasLocalStatusPage returns a boolean if a field has been set.

### GetSecurePort

`func (o *GetNetworkSettings200Response) GetSecurePort() GetNetworkSettings200ResponseSecurePort`

GetSecurePort returns the SecurePort field if non-nil, zero value otherwise.

### GetSecurePortOk

`func (o *GetNetworkSettings200Response) GetSecurePortOk() (*GetNetworkSettings200ResponseSecurePort, bool)`

GetSecurePortOk returns a tuple with the SecurePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurePort

`func (o *GetNetworkSettings200Response) SetSecurePort(v GetNetworkSettings200ResponseSecurePort)`

SetSecurePort sets SecurePort field to given value.

### HasSecurePort

`func (o *GetNetworkSettings200Response) HasSecurePort() bool`

HasSecurePort returns a boolean if a field has been set.

### GetFips

`func (o *GetNetworkSettings200Response) GetFips() GetNetworkSettings200ResponseFips`

GetFips returns the Fips field if non-nil, zero value otherwise.

### GetFipsOk

`func (o *GetNetworkSettings200Response) GetFipsOk() (*GetNetworkSettings200ResponseFips, bool)`

GetFipsOk returns a tuple with the Fips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFips

`func (o *GetNetworkSettings200Response) SetFips(v GetNetworkSettings200ResponseFips)`

SetFips sets Fips field to given value.

### HasFips

`func (o *GetNetworkSettings200Response) HasFips() bool`

HasFips returns a boolean if a field has been set.

### GetNamedVlans

`func (o *GetNetworkSettings200Response) GetNamedVlans() GetNetworkSettings200ResponseNamedVlans`

GetNamedVlans returns the NamedVlans field if non-nil, zero value otherwise.

### GetNamedVlansOk

`func (o *GetNetworkSettings200Response) GetNamedVlansOk() (*GetNetworkSettings200ResponseNamedVlans, bool)`

GetNamedVlansOk returns a tuple with the NamedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamedVlans

`func (o *GetNetworkSettings200Response) SetNamedVlans(v GetNetworkSettings200ResponseNamedVlans)`

SetNamedVlans sets NamedVlans field to given value.

### HasNamedVlans

`func (o *GetNetworkSettings200Response) HasNamedVlans() bool`

HasNamedVlans returns a boolean if a field has been set.

### GetClientPrivacy

`func (o *GetNetworkSettings200Response) GetClientPrivacy() GetNetworkSettings200ResponseClientPrivacy`

GetClientPrivacy returns the ClientPrivacy field if non-nil, zero value otherwise.

### GetClientPrivacyOk

`func (o *GetNetworkSettings200Response) GetClientPrivacyOk() (*GetNetworkSettings200ResponseClientPrivacy, bool)`

GetClientPrivacyOk returns a tuple with the ClientPrivacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPrivacy

`func (o *GetNetworkSettings200Response) SetClientPrivacy(v GetNetworkSettings200ResponseClientPrivacy)`

SetClientPrivacy sets ClientPrivacy field to given value.

### HasClientPrivacy

`func (o *GetNetworkSettings200Response) HasClientPrivacy() bool`

HasClientPrivacy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


