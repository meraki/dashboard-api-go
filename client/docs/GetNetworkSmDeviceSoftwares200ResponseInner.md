# GetNetworkSmDeviceSoftwares200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **string** | The Meraki managed application Id for this record on a particular device. | [optional] 
**BundleSize** | Pointer to **int32** | The size of the software bundle. | [optional] 
**CreatedAt** | Pointer to **string** | When the Meraki record for the software was created. | [optional] 
**DeviceId** | Pointer to **string** | The Meraki managed device Id. | [optional] 
**DynamicSize** | Pointer to **int32** | The size of the data stored in the application. | [optional] 
**Id** | Pointer to **string** | The Meraki software Id. | [optional] 
**Identifier** | Pointer to **string** | Software bundle identifier. | [optional] 
**InstalledAt** | Pointer to **string** | When the Software was installed on the device. | [optional] 
**ToInstall** | Pointer to **bool** | A boolean indicating this software record should be installed on the associated device. | [optional] 
**IosRedemptionCode** | Pointer to **bool** | A boolean indicating whether or not an iOS redemption code was used. | [optional] 
**IsManaged** | Pointer to **bool** | A boolean indicating whether or not the software is managed by Meraki. | [optional] 
**ItunesId** | Pointer to **string** | The itunes numerical identifier. | [optional] 
**LicenseKey** | Pointer to **string** | The license key associated with this software installation. | [optional] 
**Name** | Pointer to **string** | The name of the software. | [optional] 
**Path** | Pointer to **string** | The path on the device where the software record is located. | [optional] 
**RedemptionCode** | Pointer to **int32** | The redemption code used for this software. | [optional] 
**ShortVersion** | Pointer to **string** | Short version notation for the software. | [optional] 
**Status** | Pointer to **string** | The management status of the software. | [optional] 
**ToUninstall** | Pointer to **bool** | A boolean indicating this software record should be uninstalled on the associated device. | [optional] 
**UninstalledAt** | Pointer to **string** | When the record was uninstalled from the device. | [optional] 
**UpdatedAt** | Pointer to **string** | When the record was last updated by Meraki. | [optional] 
**Vendor** | Pointer to **string** | The vendor of the software. | [optional] 
**Version** | Pointer to **string** | Full version notation for the software. | [optional] 

## Methods

### NewGetNetworkSmDeviceSoftwares200ResponseInner

`func NewGetNetworkSmDeviceSoftwares200ResponseInner() *GetNetworkSmDeviceSoftwares200ResponseInner`

NewGetNetworkSmDeviceSoftwares200ResponseInner instantiates a new GetNetworkSmDeviceSoftwares200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSmDeviceSoftwares200ResponseInnerWithDefaults

`func NewGetNetworkSmDeviceSoftwares200ResponseInnerWithDefaults() *GetNetworkSmDeviceSoftwares200ResponseInner`

NewGetNetworkSmDeviceSoftwares200ResponseInnerWithDefaults instantiates a new GetNetworkSmDeviceSoftwares200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetBundleSize

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetBundleSize() int32`

GetBundleSize returns the BundleSize field if non-nil, zero value otherwise.

### GetBundleSizeOk

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetBundleSizeOk() (*int32, bool)`

GetBundleSizeOk returns a tuple with the BundleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleSize

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetBundleSize(v int32)`

SetBundleSize sets BundleSize field to given value.

### HasBundleSize

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasBundleSize() bool`

HasBundleSize returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeviceId

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDynamicSize

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetDynamicSize() int32`

GetDynamicSize returns the DynamicSize field if non-nil, zero value otherwise.

### GetDynamicSizeOk

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetDynamicSizeOk() (*int32, bool)`

GetDynamicSizeOk returns a tuple with the DynamicSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicSize

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetDynamicSize(v int32)`

SetDynamicSize sets DynamicSize field to given value.

### HasDynamicSize

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasDynamicSize() bool`

HasDynamicSize returns a boolean if a field has been set.

### GetId

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifier

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetInstalledAt

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetInstalledAt() string`

GetInstalledAt returns the InstalledAt field if non-nil, zero value otherwise.

### GetInstalledAtOk

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetInstalledAtOk() (*string, bool)`

GetInstalledAtOk returns a tuple with the InstalledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledAt

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetInstalledAt(v string)`

SetInstalledAt sets InstalledAt field to given value.

### HasInstalledAt

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasInstalledAt() bool`

HasInstalledAt returns a boolean if a field has been set.

### GetToInstall

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetToInstall() bool`

GetToInstall returns the ToInstall field if non-nil, zero value otherwise.

### GetToInstallOk

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetToInstallOk() (*bool, bool)`

GetToInstallOk returns a tuple with the ToInstall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToInstall

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetToInstall(v bool)`

SetToInstall sets ToInstall field to given value.

### HasToInstall

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasToInstall() bool`

HasToInstall returns a boolean if a field has been set.

### GetIosRedemptionCode

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetIosRedemptionCode() bool`

GetIosRedemptionCode returns the IosRedemptionCode field if non-nil, zero value otherwise.

### GetIosRedemptionCodeOk

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetIosRedemptionCodeOk() (*bool, bool)`

GetIosRedemptionCodeOk returns a tuple with the IosRedemptionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosRedemptionCode

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetIosRedemptionCode(v bool)`

SetIosRedemptionCode sets IosRedemptionCode field to given value.

### HasIosRedemptionCode

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasIosRedemptionCode() bool`

HasIosRedemptionCode returns a boolean if a field has been set.

### GetIsManaged

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetIsManaged() bool`

GetIsManaged returns the IsManaged field if non-nil, zero value otherwise.

### GetIsManagedOk

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetIsManagedOk() (*bool, bool)`

GetIsManagedOk returns a tuple with the IsManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManaged

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetIsManaged(v bool)`

SetIsManaged sets IsManaged field to given value.

### HasIsManaged

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasIsManaged() bool`

HasIsManaged returns a boolean if a field has been set.

### GetItunesId

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetItunesId() string`

GetItunesId returns the ItunesId field if non-nil, zero value otherwise.

### GetItunesIdOk

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetItunesIdOk() (*string, bool)`

GetItunesIdOk returns a tuple with the ItunesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItunesId

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetItunesId(v string)`

SetItunesId sets ItunesId field to given value.

### HasItunesId

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasItunesId() bool`

HasItunesId returns a boolean if a field has been set.

### GetLicenseKey

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetLicenseKey() string`

GetLicenseKey returns the LicenseKey field if non-nil, zero value otherwise.

### GetLicenseKeyOk

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetLicenseKeyOk() (*string, bool)`

GetLicenseKeyOk returns a tuple with the LicenseKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseKey

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetLicenseKey(v string)`

SetLicenseKey sets LicenseKey field to given value.

### HasLicenseKey

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasLicenseKey() bool`

HasLicenseKey returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRedemptionCode

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetRedemptionCode() int32`

GetRedemptionCode returns the RedemptionCode field if non-nil, zero value otherwise.

### GetRedemptionCodeOk

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetRedemptionCodeOk() (*int32, bool)`

GetRedemptionCodeOk returns a tuple with the RedemptionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedemptionCode

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetRedemptionCode(v int32)`

SetRedemptionCode sets RedemptionCode field to given value.

### HasRedemptionCode

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasRedemptionCode() bool`

HasRedemptionCode returns a boolean if a field has been set.

### GetShortVersion

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetShortVersion() string`

GetShortVersion returns the ShortVersion field if non-nil, zero value otherwise.

### GetShortVersionOk

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetShortVersionOk() (*string, bool)`

GetShortVersionOk returns a tuple with the ShortVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortVersion

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetShortVersion(v string)`

SetShortVersion sets ShortVersion field to given value.

### HasShortVersion

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasShortVersion() bool`

HasShortVersion returns a boolean if a field has been set.

### GetStatus

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetToUninstall

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetToUninstall() bool`

GetToUninstall returns the ToUninstall field if non-nil, zero value otherwise.

### GetToUninstallOk

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetToUninstallOk() (*bool, bool)`

GetToUninstallOk returns a tuple with the ToUninstall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToUninstall

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetToUninstall(v bool)`

SetToUninstall sets ToUninstall field to given value.

### HasToUninstall

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasToUninstall() bool`

HasToUninstall returns a boolean if a field has been set.

### GetUninstalledAt

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetUninstalledAt() string`

GetUninstalledAt returns the UninstalledAt field if non-nil, zero value otherwise.

### GetUninstalledAtOk

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetUninstalledAtOk() (*string, bool)`

GetUninstalledAtOk returns a tuple with the UninstalledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninstalledAt

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetUninstalledAt(v string)`

SetUninstalledAt sets UninstalledAt field to given value.

### HasUninstalledAt

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasUninstalledAt() bool`

HasUninstalledAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVendor

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVersion

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


