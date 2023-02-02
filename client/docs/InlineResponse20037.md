# InlineResponse20037

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalStatusPageEnabled** | Pointer to **bool** | Enables / disables the local device status pages (&lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;http://my.meraki.com/&#39;&gt;my.meraki.com, &lt;/a&gt;&lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;http://ap.meraki.com/&#39;&gt;ap.meraki.com, &lt;/a&gt;&lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;http://switch.meraki.com/&#39;&gt;switch.meraki.com, &lt;/a&gt;&lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;http://wired.meraki.com/&#39;&gt;wired.meraki.com&lt;/a&gt;). Optional (defaults to false) | [optional] 
**RemoteStatusPageEnabled** | Pointer to **bool** | Enables / disables access to the device status page (&lt;a target&#x3D;&#39;_blank&#39;&gt;http://[device&#39;s LAN IP])&lt;/a&gt;. Optional. Can only be set if localStatusPageEnabled is set to true | [optional] 
**LocalStatusPage** | Pointer to [**InlineResponse20037LocalStatusPage**](InlineResponse20037LocalStatusPage.md) |  | [optional] 
**SecurePort** | Pointer to [**InlineResponse20037SecurePort**](InlineResponse20037SecurePort.md) |  | [optional] 
**Fips** | Pointer to [**InlineResponse20037Fips**](InlineResponse20037Fips.md) |  | [optional] 
**NamedVlans** | Pointer to [**InlineResponse20037NamedVlans**](InlineResponse20037NamedVlans.md) |  | [optional] 
**ClientPrivacy** | Pointer to [**InlineResponse20037ClientPrivacy**](InlineResponse20037ClientPrivacy.md) |  | [optional] 

## Methods

### NewInlineResponse20037

`func NewInlineResponse20037() *InlineResponse20037`

NewInlineResponse20037 instantiates a new InlineResponse20037 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20037WithDefaults

`func NewInlineResponse20037WithDefaults() *InlineResponse20037`

NewInlineResponse20037WithDefaults instantiates a new InlineResponse20037 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalStatusPageEnabled

`func (o *InlineResponse20037) GetLocalStatusPageEnabled() bool`

GetLocalStatusPageEnabled returns the LocalStatusPageEnabled field if non-nil, zero value otherwise.

### GetLocalStatusPageEnabledOk

`func (o *InlineResponse20037) GetLocalStatusPageEnabledOk() (*bool, bool)`

GetLocalStatusPageEnabledOk returns a tuple with the LocalStatusPageEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalStatusPageEnabled

`func (o *InlineResponse20037) SetLocalStatusPageEnabled(v bool)`

SetLocalStatusPageEnabled sets LocalStatusPageEnabled field to given value.

### HasLocalStatusPageEnabled

`func (o *InlineResponse20037) HasLocalStatusPageEnabled() bool`

HasLocalStatusPageEnabled returns a boolean if a field has been set.

### GetRemoteStatusPageEnabled

`func (o *InlineResponse20037) GetRemoteStatusPageEnabled() bool`

GetRemoteStatusPageEnabled returns the RemoteStatusPageEnabled field if non-nil, zero value otherwise.

### GetRemoteStatusPageEnabledOk

`func (o *InlineResponse20037) GetRemoteStatusPageEnabledOk() (*bool, bool)`

GetRemoteStatusPageEnabledOk returns a tuple with the RemoteStatusPageEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteStatusPageEnabled

`func (o *InlineResponse20037) SetRemoteStatusPageEnabled(v bool)`

SetRemoteStatusPageEnabled sets RemoteStatusPageEnabled field to given value.

### HasRemoteStatusPageEnabled

`func (o *InlineResponse20037) HasRemoteStatusPageEnabled() bool`

HasRemoteStatusPageEnabled returns a boolean if a field has been set.

### GetLocalStatusPage

`func (o *InlineResponse20037) GetLocalStatusPage() InlineResponse20037LocalStatusPage`

GetLocalStatusPage returns the LocalStatusPage field if non-nil, zero value otherwise.

### GetLocalStatusPageOk

`func (o *InlineResponse20037) GetLocalStatusPageOk() (*InlineResponse20037LocalStatusPage, bool)`

GetLocalStatusPageOk returns a tuple with the LocalStatusPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalStatusPage

`func (o *InlineResponse20037) SetLocalStatusPage(v InlineResponse20037LocalStatusPage)`

SetLocalStatusPage sets LocalStatusPage field to given value.

### HasLocalStatusPage

`func (o *InlineResponse20037) HasLocalStatusPage() bool`

HasLocalStatusPage returns a boolean if a field has been set.

### GetSecurePort

`func (o *InlineResponse20037) GetSecurePort() InlineResponse20037SecurePort`

GetSecurePort returns the SecurePort field if non-nil, zero value otherwise.

### GetSecurePortOk

`func (o *InlineResponse20037) GetSecurePortOk() (*InlineResponse20037SecurePort, bool)`

GetSecurePortOk returns a tuple with the SecurePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurePort

`func (o *InlineResponse20037) SetSecurePort(v InlineResponse20037SecurePort)`

SetSecurePort sets SecurePort field to given value.

### HasSecurePort

`func (o *InlineResponse20037) HasSecurePort() bool`

HasSecurePort returns a boolean if a field has been set.

### GetFips

`func (o *InlineResponse20037) GetFips() InlineResponse20037Fips`

GetFips returns the Fips field if non-nil, zero value otherwise.

### GetFipsOk

`func (o *InlineResponse20037) GetFipsOk() (*InlineResponse20037Fips, bool)`

GetFipsOk returns a tuple with the Fips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFips

`func (o *InlineResponse20037) SetFips(v InlineResponse20037Fips)`

SetFips sets Fips field to given value.

### HasFips

`func (o *InlineResponse20037) HasFips() bool`

HasFips returns a boolean if a field has been set.

### GetNamedVlans

`func (o *InlineResponse20037) GetNamedVlans() InlineResponse20037NamedVlans`

GetNamedVlans returns the NamedVlans field if non-nil, zero value otherwise.

### GetNamedVlansOk

`func (o *InlineResponse20037) GetNamedVlansOk() (*InlineResponse20037NamedVlans, bool)`

GetNamedVlansOk returns a tuple with the NamedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamedVlans

`func (o *InlineResponse20037) SetNamedVlans(v InlineResponse20037NamedVlans)`

SetNamedVlans sets NamedVlans field to given value.

### HasNamedVlans

`func (o *InlineResponse20037) HasNamedVlans() bool`

HasNamedVlans returns a boolean if a field has been set.

### GetClientPrivacy

`func (o *InlineResponse20037) GetClientPrivacy() InlineResponse20037ClientPrivacy`

GetClientPrivacy returns the ClientPrivacy field if non-nil, zero value otherwise.

### GetClientPrivacyOk

`func (o *InlineResponse20037) GetClientPrivacyOk() (*InlineResponse20037ClientPrivacy, bool)`

GetClientPrivacyOk returns a tuple with the ClientPrivacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPrivacy

`func (o *InlineResponse20037) SetClientPrivacy(v InlineResponse20037ClientPrivacy)`

SetClientPrivacy sets ClientPrivacy field to given value.

### HasClientPrivacy

`func (o *InlineResponse20037) HasClientPrivacy() bool`

HasClientPrivacy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


