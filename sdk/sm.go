package meraki

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type SmService service

type GetNetworkSmDevicesQueryParams struct {
	Fields        []string `url:"fields[],omitempty"`      //Additional fields that will be displayed for each device.     The default fields are: id, name, tags, ssid, wifiMac, osName, systemModel, uuid, and serialNumber. The additional fields are: ip,     systemType, availableDeviceCapacity, kioskAppName, biosVersion, lastConnected, missingAppsCount, userSuppliedAddress, location, lastUser,     ownerEmail, ownerUsername, osBuild, publicIp, phoneNumber, diskInfoJson, deviceCapacity, isManaged, hadMdm, isSupervised, meid, imei, iccid,     simCarrierNetwork, cellularDataUsed, isHotspotEnabled, createdAt, batteryEstCharge, quarantined, avName, avRunning, asName, fwName,     isRooted, loginRequired, screenLockEnabled, screenLockDelay, autoLoginDisabled, autoTags, hasMdm, hasDesktopAgent, diskEncryptionEnabled,     hardwareEncryptionCaps, passCodeLock, usesHardwareKeystore, androidSecurityPatchVersion, and url.
	WifiMacs      []string `url:"wifiMacs[],omitempty"`    //Filter devices by wifi mac(s).
	Serials       []string `url:"serials[],omitempty"`     //Filter devices by serial(s).
	IDs           []string `url:"ids[],omitempty"`         //Filter devices by id(s).
	Scope         []string `url:"scope[],omitempty"`       //Specify a scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags.
	PerPage       int      `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter string   `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string   `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
}
type GetNetworkSmDeviceConnectivityQueryParams struct {
	PerPage       int    `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter string `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
}
type GetNetworkSmDeviceDesktopLogsQueryParams struct {
	PerPage       int    `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter string `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
}
type GetNetworkSmDeviceDeviceCommandLogsQueryParams struct {
	PerPage       int    `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter string `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
}
type GetNetworkSmDevicePerformanceHistoryQueryParams struct {
	PerPage       int    `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter string `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
}
type GetNetworkSmTargetGroupsQueryParams struct {
	WithDetails bool `url:"withDetails,omitempty"` //Boolean indicating if the the ids of the devices or users scoped by the target group should be included in the response
}
type GetNetworkSmTargetGroupQueryParams struct {
	WithDetails bool `url:"withDetails,omitempty"` //Boolean indicating if the the ids of the devices or users scoped by the target group should be included in the response
}
type GetNetworkSmTrustedAccessConfigsQueryParams struct {
	PerPage       int    `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100.
	StartingAfter string `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
}
type GetNetworkSmUserAccessDevicesQueryParams struct {
	PerPage       int    `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100.
	StartingAfter string `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
}
type GetNetworkSmUsersQueryParams struct {
	IDs       []string `url:"ids[],omitempty"`       //Filter users by id(s).
	Usernames []string `url:"usernames[],omitempty"` //Filter users by username(s).
	Emails    []string `url:"emails[],omitempty"`    //Filter users by email(s).
	Scope     []string `url:"scope[],omitempty"`     //Specifiy a scope (one of all, none, withAny, withAll, withoutAny, withoutAll) and a set of tags.
}

type ResponseSmCreateNetworkSmBypassActivationLockAttempt interface{}
type ResponseSmGetNetworkSmBypassActivationLockAttempt struct {
	Data   *ResponseSmGetNetworkSmBypassActivationLockAttemptData `json:"data,omitempty"`   //
	ID     string                                                 `json:"id,omitempty"`     //
	Status string                                                 `json:"status,omitempty"` //
}
type ResponseSmGetNetworkSmBypassActivationLockAttemptData struct {
	Status2090938209  *ResponseSmGetNetworkSmBypassActivationLockAttemptData2090938209  `json:"2090938209,omitempty"`  //
	Status38290139892 *ResponseSmGetNetworkSmBypassActivationLockAttemptData38290139892 `json:"38290139892,omitempty"` //
}
type ResponseSmGetNetworkSmBypassActivationLockAttemptData2090938209 struct {
	Errors  []string `json:"errors,omitempty"`  //
	Success *bool    `json:"success,omitempty"` //
}
type ResponseSmGetNetworkSmBypassActivationLockAttemptData38290139892 struct {
	Success *bool `json:"success,omitempty"` //
}
type ResponseSmGetNetworkSmDevices []ResponseItemSmGetNetworkSmDevices // Array of ResponseSmGetNetworkSmDevices
type ResponseItemSmGetNetworkSmDevices struct {
	ID           string   `json:"id,omitempty"`           // The Meraki Id of the device record.
	IP           string   `json:"ip,omitempty"`           // The IP address of the device.
	Name         string   `json:"name,omitempty"`         // The name of the device.
	Notes        string   `json:"notes,omitempty"`        // Notes associated with the device.
	OsName       string   `json:"osName,omitempty"`       // The name of the device OS.
	Serial       string   `json:"serial,omitempty"`       // The device serial.
	SerialNumber string   `json:"serialNumber,omitempty"` // The device serial number.
	SSID         string   `json:"ssid,omitempty"`         // The name of the SSID the device was last connected to.
	SystemModel  string   `json:"systemModel,omitempty"`  // The device model.
	Tags         []string `json:"tags,omitempty"`         // An array of tags associated with the device.
	UUID         string   `json:"uuid,omitempty"`         // The UUID of the device.
	WifiMac      string   `json:"wifiMac,omitempty"`      // The MAC of the device.
}
type ResponseSmCheckinNetworkSmDevices struct {
	IDs []string `json:"ids,omitempty"` // The Meraki Ids of the set of devices.
}
type ResponseSmUpdateNetworkSmDevicesFields []ResponseItemSmUpdateNetworkSmDevicesFields // Array of ResponseSmUpdateNetworkSmDevicesFields
type ResponseItemSmUpdateNetworkSmDevicesFields struct {
	ID      string `json:"id,omitempty"`      // The Meraki Id of the device record.
	Name    string `json:"name,omitempty"`    // The name of the device.
	Notes   string `json:"notes,omitempty"`   // Notes associated with the device.
	Serial  string `json:"serial,omitempty"`  // The device serial.
	WifiMac string `json:"wifiMac,omitempty"` // The MAC of the device.
}
type ResponseSmLockNetworkSmDevices struct {
	IDs []string `json:"ids,omitempty"` // The Meraki Ids of the set of devices.
}
type ResponseSmModifyNetworkSmDevicesTags []ResponseItemSmModifyNetworkSmDevicesTags // Array of ResponseSmModifyNetworkSmDevicesTags
type ResponseItemSmModifyNetworkSmDevicesTags struct {
	ID      string   `json:"id,omitempty"`      // The Meraki Id of the device record.
	Serial  string   `json:"serial,omitempty"`  // The device serial.
	Tags    []string `json:"tags,omitempty"`    // An array of tags associated with the device.
	WifiMac string   `json:"wifiMac,omitempty"` // The MAC of the device.
}
type ResponseSmMoveNetworkSmDevices struct {
	IDs        []string `json:"ids,omitempty"`        // The Meraki Ids of the set of devices.
	NewNetwork string   `json:"newNetwork,omitempty"` // The network to which the devices was moved.
}
type ResponseSmWipeNetworkSmDevices struct {
	ID string `json:"id,omitempty"` // The Meraki Id of the devices.
}
type ResponseSmGetNetworkSmDeviceCellularUsageHistory []ResponseItemSmGetNetworkSmDeviceCellularUsageHistory // Array of ResponseSmGetNetworkSmDeviceCellularUsageHistory
type ResponseItemSmGetNetworkSmDeviceCellularUsageHistory struct {
	Received *float64 `json:"received,omitempty"` // The amount of cellular data received by the device.
	Sent     *float64 `json:"sent,omitempty"`     // The amount of cellular sent received by the device.
	Ts       string   `json:"ts,omitempty"`       // When the cellular usage data was collected.
}
type ResponseSmGetNetworkSmDeviceCerts []ResponseItemSmGetNetworkSmDeviceCerts // Array of ResponseSmGetNetworkSmDeviceCerts
type ResponseItemSmGetNetworkSmDeviceCerts struct {
	CertPem        string `json:"certPem,omitempty"`        // The PEM of the certificate.
	DeviceID       string `json:"deviceId,omitempty"`       // The Meraki managed device Id.
	ID             string `json:"id,omitempty"`             // The Meraki Id of the certificate record.
	Issuer         string `json:"issuer,omitempty"`         // The certificate issuer.
	Name           string `json:"name,omitempty"`           // The name of the certificate.
	NotValidAfter  string `json:"notValidAfter,omitempty"`  // The date after which the certificate is no longer valid.
	NotValidBefore string `json:"notValidBefore,omitempty"` // The date before which the certificate is not valid.
	Subject        string `json:"subject,omitempty"`        // The subject of the certificate.
}
type ResponseSmGetNetworkSmDeviceConnectivity []ResponseItemSmGetNetworkSmDeviceConnectivity // Array of ResponseSmGetNetworkSmDeviceConnectivity
type ResponseItemSmGetNetworkSmDeviceConnectivity struct {
	FirstSeenAt string `json:"firstSeenAt,omitempty"` // When the device was first seen as connected to the internet in each connection.
	LastSeenAt  string `json:"lastSeenAt,omitempty"`  // When the device was last seen as connected to the internet in each connection.
}
type ResponseSmGetNetworkSmDeviceDesktopLogs []ResponseItemSmGetNetworkSmDeviceDesktopLogs // Array of ResponseSmGetNetworkSmDeviceDesktopLogs
type ResponseItemSmGetNetworkSmDeviceDesktopLogs struct {
	DhcpServer    string `json:"dhcpServer,omitempty"`    // The IP address of the DCHP Server.
	DNSServer     string `json:"dnsServer,omitempty"`     // The DNS Server during the connection.
	Gateway       string `json:"gateway,omitempty"`       // The gateway IP the device was connected to.
	IP            string `json:"ip,omitempty"`            // The IP of the device during connection.
	MeasuredAt    string `json:"measuredAt,omitempty"`    // The time the data was measured at.
	NetworkDevice string `json:"networkDevice,omitempty"` // The network device for the device used for connection.
	NetworkDriver string `json:"networkDriver,omitempty"` // The network driver for the device.
	NetworkMTU    string `json:"networkMTU,omitempty"`    // The network max transmission unit.
	PublicIP      string `json:"publicIP,omitempty"`      // The public IP address of the device.
	Subnet        string `json:"subnet,omitempty"`        // The subnet of the device connection.
	Ts            string `json:"ts,omitempty"`            // The time the connection was logged.
	User          string `json:"user,omitempty"`          // The user during connection.
	WifiAuth      string `json:"wifiAuth,omitempty"`      // The type of authentication used by the SSID.
	WifiBssid     string `json:"wifiBssid,omitempty"`     // The MAC of the access point the device is connected to.
	WifiChannel   string `json:"wifiChannel,omitempty"`   // Channel through which the connection is routing.
	WifiNoise     string `json:"wifiNoise,omitempty"`     // The wireless signal power level received by the device.
	WifiRssi      string `json:"wifiRssi,omitempty"`      // The Received Signal Strength Indicator for the device.
	WifiSSID      string `json:"wifiSsid,omitempty"`      // The name of the network the device is connected to.
}
type ResponseSmGetNetworkSmDeviceDeviceCommandLogs []ResponseItemSmGetNetworkSmDeviceDeviceCommandLogs // Array of ResponseSmGetNetworkSmDeviceDeviceCommandLogs
type ResponseItemSmGetNetworkSmDeviceDeviceCommandLogs struct {
	Action        string `json:"action,omitempty"`        // The type of command sent to the device.
	DashboardUser string `json:"dashboardUser,omitempty"` // The Meraki dashboard user who initiated the command.
	Details       string `json:"details,omitempty"`       // A JSON string object containing command details.
	Name          string `json:"name,omitempty"`          // The name of the device to which the command is sent.
	Ts            string `json:"ts,omitempty"`            // The time the command was sent to the device.
}
type ResponseSmGetNetworkSmDeviceDeviceProfiles []ResponseItemSmGetNetworkSmDeviceDeviceProfiles // Array of ResponseSmGetNetworkSmDeviceDeviceProfiles
type ResponseItemSmGetNetworkSmDeviceDeviceProfiles struct {
	DeviceID          string `json:"deviceId,omitempty"`          // The Meraki managed device Id.
	ID                string `json:"id,omitempty"`                // The numerical Meraki Id of the profile.
	IsEncrypted       *bool  `json:"isEncrypted,omitempty"`       // A boolean indicating if the profile is encrypted.
	IsManaged         *bool  `json:"isManaged,omitempty"`         // Whether or not the profile is managed by Meraki.
	Name              string `json:"name,omitempty"`              // The name of the profile.
	ProfileData       string `json:"profileData,omitempty"`       // A string containing a JSON object with the profile data.
	ProfileIDentifier string `json:"profileIdentifier,omitempty"` // The identifier of the profile.
	Version           string `json:"version,omitempty"`           // The verison of the profile.
}
type ResponseSmGetNetworkSmDeviceNetworkAdapters []ResponseItemSmGetNetworkSmDeviceNetworkAdapters // Array of ResponseSmGetNetworkSmDeviceNetworkAdapters
type ResponseItemSmGetNetworkSmDeviceNetworkAdapters struct {
	DhcpServer string `json:"dhcpServer,omitempty"` // The IP address of the DCHP Server.
	DNSServer  string `json:"dnsServer,omitempty"`  // The IP address of the DNS Server.
	Gateway    string `json:"gateway,omitempty"`    // The IP address of the Gateway.
	ID         string `json:"id,omitempty"`         // The Meraki Id of the network adapter record.
	IP         string `json:"ip,omitempty"`         // The IP address of the network adapter.
	Mac        string `json:"mac,omitempty"`        // The MAC associated with the network adapter.
	Name       string `json:"name,omitempty"`       // The name of the newtwork adapter.
	Subnet     string `json:"subnet,omitempty"`     // The subnet for the network adapter.
}
type ResponseSmGetNetworkSmDevicePerformanceHistory []ResponseItemSmGetNetworkSmDevicePerformanceHistory // Array of ResponseSmGetNetworkSmDevicePerformanceHistory
type ResponseItemSmGetNetworkSmDevicePerformanceHistory struct {
	CPUPercentUsed  *float64                                                     `json:"cpuPercentUsed,omitempty"`  // The percentage of CPU used as a decimal format.
	DiskUsage       *ResponseItemSmGetNetworkSmDevicePerformanceHistoryDiskUsage `json:"diskUsage,omitempty"`       // An object containing disk usage details.
	MemActive       *int                                                         `json:"memActive,omitempty"`       // The active RAM on the device.
	MemFree         *int                                                         `json:"memFree,omitempty"`         // Memory that is not yet in use by the system.
	MemInactive     *int                                                         `json:"memInactive,omitempty"`     // The inactive RAM on the device.
	MemWired        *int                                                         `json:"memWired,omitempty"`        // Memory used for core OS functions on the device.
	NetworkReceived *int                                                         `json:"networkReceived,omitempty"` // Network bandwith received.
	NetworkSent     *int                                                         `json:"networkSent,omitempty"`     // Network bandwith transmitted.
	SwapUsed        *int                                                         `json:"swapUsed,omitempty"`        // The amount of space being used on the startup disk to swap unused files to and from RAM.
	Ts              string                                                       `json:"ts,omitempty"`              // The time at which the performance was measured.
}
type ResponseItemSmGetNetworkSmDevicePerformanceHistoryDiskUsage struct {
	C *ResponseItemSmGetNetworkSmDevicePerformanceHistoryDiskUsageC `json:"c,omitempty"` // An object containing current disk usage details.
}
type ResponseItemSmGetNetworkSmDevicePerformanceHistoryDiskUsageC struct {
	Space *int `json:"space,omitempty"` // The available disk space.
	Used  *int `json:"used,omitempty"`  // The used disk space.
}
type ResponseSmGetNetworkSmDeviceRestrictions []ResponseItemSmGetNetworkSmDeviceRestrictions // Array of ResponseSmGetNetworkSmDeviceRestrictions
type ResponseItemSmGetNetworkSmDeviceRestrictions struct {
	Profile      string                                                    `json:"profile,omitempty"`      //
	Restrictions *ResponseItemSmGetNetworkSmDeviceRestrictionsRestrictions `json:"restrictions,omitempty"` //
}
type ResponseItemSmGetNetworkSmDeviceRestrictionsRestrictions struct {
	MyRestriction *ResponseItemSmGetNetworkSmDeviceRestrictionsRestrictionsMyRestriction `json:"myRestriction,omitempty"` //
}
type ResponseItemSmGetNetworkSmDeviceRestrictionsRestrictionsMyRestriction struct {
	Value *bool `json:"value,omitempty"` //
}
type ResponseSmGetNetworkSmDeviceSecurityCenters []ResponseItemSmGetNetworkSmDeviceSecurityCenters // Array of ResponseSmGetNetworkSmDeviceSecurityCenters
type ResponseItemSmGetNetworkSmDeviceSecurityCenters struct {
	AntiVirusName        string `json:"antiVirusName,omitempty"`        // The name of the Antivirus.
	FireWallName         string `json:"fireWallName,omitempty"`         // The name of the Firewall.
	HasAntiVirus         *bool  `json:"hasAntiVirus,omitempty"`         // Boolean indicating if the device has Antivirus.
	HasFireWallInstalled *bool  `json:"hasFireWallInstalled,omitempty"` // Boolean indicating if the device has a Firewall installed.
	ID                   string `json:"id,omitempty"`                   // The Meraki identifier for the security center record.
	IsAutoLoginDisabled  *bool  `json:"isAutoLoginDisabled,omitempty"`  // Boolean indicating if the device has auto login disabled.
	IsDiskEncrypted      *bool  `json:"isDiskEncrypted,omitempty"`      // Boolean indicating if the device has disk encryption.
	IsFireWallEnabled    *bool  `json:"isFireWallEnabled,omitempty"`    // Boolean indicating if the device has a Firewall enabled.
	IsRooted             *bool  `json:"isRooted,omitempty"`             // Boolean indicating if the device is rooted.
	RunningProcs         string `json:"runningProcs,omitempty"`         // A comma seperated list of procs running on the device.
}
type ResponseSmGetNetworkSmDeviceSoftwares []ResponseItemSmGetNetworkSmDeviceSoftwares // Array of ResponseSmGetNetworkSmDeviceSoftwares
type ResponseItemSmGetNetworkSmDeviceSoftwares struct {
	AppID             string `json:"appId,omitempty"`             // The Meraki managed application Id for this record on a particular device.
	BundleSize        *int   `json:"bundleSize,omitempty"`        // The size of the software bundle.
	CreatedAt         string `json:"createdAt,omitempty"`         // When the Meraki record for the software was created.
	DeviceID          string `json:"deviceId,omitempty"`          // The Meraki managed device Id.
	DynamicSize       *int   `json:"dynamicSize,omitempty"`       // The size of the data stored in the application.
	ID                string `json:"id,omitempty"`                // The Meraki software Id.
	IDentifier        string `json:"identifier,omitempty"`        // Software bundle identifier.
	InstalledAt       string `json:"installedAt,omitempty"`       // When the Software was installed on the device.
	IosRedemptionCode *bool  `json:"iosRedemptionCode,omitempty"` // A boolean indicating whether or not an iOS redemption code was used.
	IsManaged         *bool  `json:"isManaged,omitempty"`         // A boolean indicating whether or not the software is managed by Meraki.
	ItunesID          string `json:"itunesId,omitempty"`          // The itunes numerical identifier.
	LicenseKey        string `json:"licenseKey,omitempty"`        // The license key associated with this software installation.
	Name              string `json:"name,omitempty"`              // The name of the software.
	Path              string `json:"path,omitempty"`              // The path on the device where the software record is located.
	RedemptionCode    *int   `json:"redemptionCode,omitempty"`    // The redemption code used for this software.
	ShortVersion      string `json:"shortVersion,omitempty"`      // Short version notation for the software.
	Status            string `json:"status,omitempty"`            // The management status of the software.
	ToInstall         *bool  `json:"toInstall,omitempty"`         // A boolean indicating this software record should be installed on the associated device.
	ToUninstall       *bool  `json:"toUninstall,omitempty"`       // A boolean indicating this software record should be uninstalled on the associated device.
	UninstalledAt     string `json:"uninstalledAt,omitempty"`     // When the record was uninstalled from the device.
	UpdatedAt         string `json:"updatedAt,omitempty"`         // When the record was last updated by Meraki.
	Vendor            string `json:"vendor,omitempty"`            // The vendor of the software.
	Version           string `json:"version,omitempty"`           // Full version notation for the software.
}
type ResponseSmUnenrollNetworkSmDevice interface{}
type ResponseSmGetNetworkSmDeviceWLANLists []ResponseItemSmGetNetworkSmDeviceWLANLists // Array of ResponseSmGetNetworkSmDeviceWlanLists
type ResponseItemSmGetNetworkSmDeviceWLANLists struct {
	CreatedAt string `json:"createdAt,omitempty"` // When the Meraki record for the wlanList was created.
	ID        string `json:"id,omitempty"`        // The Meraki managed Id of the wlanList record.
	Xml       string `json:"xml,omitempty"`       // An XML string containing the WLAN List for the device.
}
type ResponseSmGetNetworkSmProfiles []ResponseItemSmGetNetworkSmProfiles // Array of ResponseSmGetNetworkSmProfiles
type ResponseItemSmGetNetworkSmProfiles struct {
	Description string   `json:"description,omitempty"` // Description of a profile.
	ID          string   `json:"id,omitempty"`          // ID of a profile.
	Name        string   `json:"name,omitempty"`        // Name of a profile.
	Scope       string   `json:"scope,omitempty"`       // Scope of a profile.
	Tags        []string `json:"tags,omitempty"`        // Tags of a profile.
}
type ResponseSmGetNetworkSmTargetGroups []ResponseItemSmGetNetworkSmTargetGroups // Array of ResponseSmGetNetworkSmTargetGroups
type ResponseItemSmGetNetworkSmTargetGroups struct {
	Name  string `json:"name,omitempty"`  //
	Scope string `json:"scope,omitempty"` //
	Tags  string `json:"tags,omitempty"`  //
	Type  string `json:"type,omitempty"`  //
}
type ResponseSmCreateNetworkSmTargetGroup interface{}
type ResponseSmGetNetworkSmTargetGroup struct {
	Name  string `json:"name,omitempty"`  //
	Scope string `json:"scope,omitempty"` //
	Tags  string `json:"tags,omitempty"`  //
	Type  string `json:"type,omitempty"`  //
}
type ResponseSmUpdateNetworkSmTargetGroup interface{}
type ResponseSmGetNetworkSmTrustedAccessConfigs []ResponseItemSmGetNetworkSmTrustedAccessConfigs // Array of ResponseSmGetNetworkSmTrustedAccessConfigs
type ResponseItemSmGetNetworkSmTrustedAccessConfigs struct {
	AccessEndAt   string   `json:"accessEndAt,omitempty"`   // time that access ends
	AccessStartAt string   `json:"accessStartAt,omitempty"` // time that access starts
	ID            string   `json:"id,omitempty"`            // device ID
	Name          string   `json:"name,omitempty"`          // device name
	Scope         string   `json:"scope,omitempty"`         // scope
	SSIDName      string   `json:"ssidName,omitempty"`      // SSID name
	Tags          []string `json:"tags,omitempty"`          // device tags
	TimeboundType string   `json:"timeboundType,omitempty"` // type of access period, either a static range or a dynamic period
}
type ResponseSmGetNetworkSmUserAccessDevices []ResponseItemSmGetNetworkSmUserAccessDevices // Array of ResponseSmGetNetworkSmUserAccessDevices
type ResponseItemSmGetNetworkSmUserAccessDevices struct {
	Email                    string                                                                 `json:"email,omitempty"`                    // user email
	ID                       string                                                                 `json:"id,omitempty"`                       // device ID
	Mac                      string                                                                 `json:"mac,omitempty"`                      // mac address
	Name                     string                                                                 `json:"name,omitempty"`                     // device name
	SystemType               string                                                                 `json:"systemType,omitempty"`               // system type
	Tags                     []string                                                               `json:"tags,omitempty"`                     // device tags
	TrustedAccessConnections *[]ResponseItemSmGetNetworkSmUserAccessDevicesTrustedAccessConnections `json:"trustedAccessConnections,omitempty"` // Array of trusted access configs
	Username                 string                                                                 `json:"username,omitempty"`                 // username
}
type ResponseItemSmGetNetworkSmUserAccessDevicesTrustedAccessConnections struct {
	DownloadedAt          string `json:"downloadedAt,omitempty"`          // time that config was downloaded
	LastConnectedAt       string `json:"lastConnectedAt,omitempty"`       // time of last connection
	ScepCompletedAt       string `json:"scepCompletedAt,omitempty"`       // time that SCEP completed
	TrustedAccessConfigID string `json:"trustedAccessConfigId,omitempty"` // config id
}
type ResponseSmGetNetworkSmUsers []ResponseItemSmGetNetworkSmUsers // Array of ResponseSmGetNetworkSmUsers
type ResponseItemSmGetNetworkSmUsers struct {
	AdGroups               []string `json:"adGroups,omitempty"`               // Active Directory Groups the user belongs to.
	AsmGroups              []string `json:"asmGroups,omitempty"`              // Apple School Manager Groups the user belongs to.
	AzureAdGroups          []string `json:"azureAdGroups,omitempty"`          // Azure Active Directory Groups the user belongs to.
	DisplayName            string   `json:"displayName,omitempty"`            // The user display name.
	Email                  string   `json:"email,omitempty"`                  // User email.
	FullName               string   `json:"fullName,omitempty"`               // User full name.
	HasIDentityCertificate *bool    `json:"hasIdentityCertificate,omitempty"` // A boolean indicating if the user has an associated identity certificate..
	HasPassword            *bool    `json:"hasPassword,omitempty"`            // A boolean denoting if the user has a password associated with the record.
	ID                     string   `json:"id,omitempty"`                     // The Meraki managed Id of the user record.
	IsExternal             *bool    `json:"isExternal,omitempty"`             // Whether the user was created using an external integration, or via the Meraki Dashboard.
	SamlGroups             []string `json:"samlGroups,omitempty"`             // SAML Groups the user belongs to.
	Tags                   string   `json:"tags,omitempty"`                   // The set of tags the user is scoped to.
	UserThumbnail          string   `json:"userThumbnail,omitempty"`          // The url where the users thumbnail is hosted.
	Username               string   `json:"username,omitempty"`               // The users username.
}
type ResponseSmGetNetworkSmUserDeviceProfiles []ResponseItemSmGetNetworkSmUserDeviceProfiles // Array of ResponseSmGetNetworkSmUserDeviceProfiles
type ResponseItemSmGetNetworkSmUserDeviceProfiles struct {
	DeviceID          string `json:"deviceId,omitempty"`          // The Meraki managed device Id.
	ID                string `json:"id,omitempty"`                // The numerical Meraki Id of the profile.
	IsEncrypted       *bool  `json:"isEncrypted,omitempty"`       // A boolean indicating if the profile is encrypted.
	IsManaged         *bool  `json:"isManaged,omitempty"`         // Whether or not the profile is managed by Meraki.
	Name              string `json:"name,omitempty"`              // The name of the profile.
	ProfileData       string `json:"profileData,omitempty"`       // A string containing a JSON object with the profile data.
	ProfileIDentifier string `json:"profileIdentifier,omitempty"` // The identifier of the profile.
	Version           string `json:"version,omitempty"`           // The verison of the profile.
}
type ResponseSmGetNetworkSmUserSoftwares []ResponseItemSmGetNetworkSmUserSoftwares // Array of ResponseSmGetNetworkSmUserSoftwares
type ResponseItemSmGetNetworkSmUserSoftwares struct {
	AppID             string `json:"appId,omitempty"`             // The Meraki managed application Id for this record on a particular device.
	BundleSize        *int   `json:"bundleSize,omitempty"`        // The size of the software bundle.
	CreatedAt         string `json:"createdAt,omitempty"`         // When the Meraki record for the software was created.
	DeviceID          string `json:"deviceId,omitempty"`          // The Meraki managed device Id.
	DynamicSize       *int   `json:"dynamicSize,omitempty"`       // The size of the data stored in the application.
	ID                string `json:"id,omitempty"`                // The Meraki software Id.
	IDentifier        string `json:"identifier,omitempty"`        // Software bundle identifier.
	InstalledAt       string `json:"installedAt,omitempty"`       // When the Software was installed on the device.
	IosRedemptionCode *bool  `json:"iosRedemptionCode,omitempty"` // A boolean indicating whether or not an iOS redemption code was used.
	IsManaged         *bool  `json:"isManaged,omitempty"`         // A boolean indicating whether or not the software is managed by Meraki.
	ItunesID          string `json:"itunesId,omitempty"`          // The itunes numerical identifier.
	LicenseKey        string `json:"licenseKey,omitempty"`        // The license key associated with this software installation.
	Name              string `json:"name,omitempty"`              // The name of the software.
	Path              string `json:"path,omitempty"`              // The path on the device where the software record is located.
	RedemptionCode    *int   `json:"redemptionCode,omitempty"`    // The redemption code used for this software.
	ShortVersion      string `json:"shortVersion,omitempty"`      // Short version notation for the software.
	Status            string `json:"status,omitempty"`            // The management status of the software.
	ToInstall         *bool  `json:"toInstall,omitempty"`         // A boolean indicating this software record should be installed on the associated device.
	ToUninstall       *bool  `json:"toUninstall,omitempty"`       // A boolean indicating this software record should be uninstalled on the associated device.
	UninstalledAt     string `json:"uninstalledAt,omitempty"`     // When the record was uninstalled from the device.
	UpdatedAt         string `json:"updatedAt,omitempty"`         // When the record was last updated by Meraki.
	Vendor            string `json:"vendor,omitempty"`            // The vendor of the software.
	Version           string `json:"version,omitempty"`           // Full version notation for the software.
}
type ResponseSmGetOrganizationSmApnsCert struct {
	Certificate string `json:"certificate,omitempty"` // Organization APNS Certificate used by devices to communication with Apple
}
type ResponseSmGetOrganizationSmVppAccounts []ResponseItemSmGetOrganizationSmVppAccounts // Array of ResponseSmGetOrganizationSmVppAccounts
type ResponseItemSmGetOrganizationSmVppAccounts struct {
	ID              string `json:"id,omitempty"`              // The id of the VPP Account
	VppServiceToken string `json:"vppServiceToken,omitempty"` // The VPP Account's Service Token
}
type ResponseSmGetOrganizationSmVppAccount struct {
	ID              string `json:"id,omitempty"`              // The id of the VPP Account
	VppServiceToken string `json:"vppServiceToken,omitempty"` // The VPP Account's Service Token
}
type RequestSmCreateNetworkSmBypassActivationLockAttempt struct {
	IDs []string `json:"ids,omitempty"` // The ids of the devices to attempt activation lock bypass.
}
type RequestSmCheckinNetworkSmDevices struct {
	IDs      []string `json:"ids,omitempty"`      // The ids of the devices to be checked-in.
	Scope    []string `json:"scope,omitempty"`    // The scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags of the devices to be checked-in.
	Serials  []string `json:"serials,omitempty"`  // The serials of the devices to be checked-in.
	WifiMacs []string `json:"wifiMacs,omitempty"` // The wifiMacs of the devices to be checked-in.
}
type RequestSmUpdateNetworkSmDevicesFields struct {
	DeviceFields *RequestSmUpdateNetworkSmDevicesFieldsDeviceFields `json:"deviceFields,omitempty"` // The new fields of the device. Each field of this object is optional.
	ID           string                                             `json:"id,omitempty"`           // The id of the device to be modified.
	Serial       string                                             `json:"serial,omitempty"`       // The serial of the device to be modified.
	WifiMac      string                                             `json:"wifiMac,omitempty"`      // The wifiMac of the device to be modified.
}
type RequestSmUpdateNetworkSmDevicesFieldsDeviceFields struct {
	Name  string `json:"name,omitempty"`  // New name for the device
	Notes string `json:"notes,omitempty"` // New notes for the device
}
type RequestSmLockNetworkSmDevices struct {
	IDs      []string `json:"ids,omitempty"`      // The ids of the devices to be locked.
	Pin      *int     `json:"pin,omitempty"`      // The pin number for locking macOS devices (a six digit number). Required only for macOS devices.
	Scope    []string `json:"scope,omitempty"`    // The scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags of the devices to be wiped.
	Serials  []string `json:"serials,omitempty"`  // The serials of the devices to be locked.
	WifiMacs []string `json:"wifiMacs,omitempty"` // The wifiMacs of the devices to be locked.
}
type RequestSmModifyNetworkSmDevicesTags struct {
	IDs          []string `json:"ids,omitempty"`          // The ids of the devices to be modified.
	Scope        []string `json:"scope,omitempty"`        // The scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags of the devices to be modified.
	Serials      []string `json:"serials,omitempty"`      // The serials of the devices to be modified.
	Tags         []string `json:"tags,omitempty"`         // The tags to be added, deleted, or updated.
	UpdateAction string   `json:"updateAction,omitempty"` // One of add, delete, or update. Only devices that have been modified will be returned.
	WifiMacs     []string `json:"wifiMacs,omitempty"`     // The wifiMacs of the devices to be modified.
}
type RequestSmMoveNetworkSmDevices struct {
	IDs        []string `json:"ids,omitempty"`        // The ids of the devices to be moved.
	NewNetwork string   `json:"newNetwork,omitempty"` // The new network to which the devices will be moved.
	Scope      []string `json:"scope,omitempty"`      // The scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags of the devices to be moved.
	Serials    []string `json:"serials,omitempty"`    // The serials of the devices to be moved.
	WifiMacs   []string `json:"wifiMacs,omitempty"`   // The wifiMacs of the devices to be moved.
}
type RequestSmWipeNetworkSmDevices struct {
	ID      string `json:"id,omitempty"`      // The id of the device to be wiped.
	Pin     *int   `json:"pin,omitempty"`     // The pin number (a six digit value) for wiping a macOS device. Required only for macOS devices.
	Serial  string `json:"serial,omitempty"`  // The serial of the device to be wiped.
	WifiMac string `json:"wifiMac,omitempty"` // The wifiMac of the device to be wiped.
}
type RequestSmCreateNetworkSmTargetGroup struct {
	Name  string `json:"name,omitempty"`  // The name of this target group
	Scope string `json:"scope,omitempty"` // The scope and tag options of the target group. Comma separated values beginning with one of withAny, withAll, withoutAny, withoutAll, all, none, followed by tags. Default to none if empty.
}
type RequestSmUpdateNetworkSmTargetGroup struct {
	Name  string `json:"name,omitempty"`  // The name of this target group
	Scope string `json:"scope,omitempty"` // The scope and tag options of the target group. Comma separated values beginning with one of withAny, withAll, withoutAny, withoutAll, all, none, followed by tags. Default to none if empty.
}

//GetNetworkSmBypassActivationLockAttempt Bypass activation lock attempt status
/* Bypass activation lock attempt status

@param networkID networkId path parameter. Network ID
@param attemptID attemptId path parameter. Attempt ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sm-bypass-activation-lock-attempt
*/
func (s *SmService) GetNetworkSmBypassActivationLockAttempt(networkID string, attemptID string) (*ResponseSmGetNetworkSmBypassActivationLockAttempt, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/bypassActivationLockAttempts/{attemptId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{attemptId}", fmt.Sprintf("%v", attemptID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSmGetNetworkSmBypassActivationLockAttempt{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSmBypassActivationLockAttempt")
	}

	result := response.Result().(*ResponseSmGetNetworkSmBypassActivationLockAttempt)
	return result, response, err

}

//GetNetworkSmDevices List the devices enrolled in an SM network with various specified fields and filters
/* List the devices enrolled in an SM network with various specified fields and filters

@param networkID networkId path parameter. Network ID
@param getNetworkSmDevicesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sm-devices
*/
func (s *SmService) GetNetworkSmDevices(networkID string, getNetworkSmDevicesQueryParams *GetNetworkSmDevicesQueryParams) (*ResponseSmGetNetworkSmDevices, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(getNetworkSmDevicesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSmGetNetworkSmDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSmDevices")
	}

	result := response.Result().(*ResponseSmGetNetworkSmDevices)
	return result, response, err

}

//GetNetworkSmDeviceCellularUsageHistory Return the client's daily cellular data usage history
/* Return the client's daily cellular data usage history. Usage data is in kilobytes.

@param networkID networkId path parameter. Network ID
@param deviceID deviceId path parameter. Device ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sm-device-cellular-usage-history
*/
func (s *SmService) GetNetworkSmDeviceCellularUsageHistory(networkID string, deviceID string) (*ResponseSmGetNetworkSmDeviceCellularUsageHistory, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/{deviceId}/cellularUsageHistory"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSmGetNetworkSmDeviceCellularUsageHistory{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSmDeviceCellularUsageHistory")
	}

	result := response.Result().(*ResponseSmGetNetworkSmDeviceCellularUsageHistory)
	return result, response, err

}

//GetNetworkSmDeviceCerts List the certs on a device
/* List the certs on a device

@param networkID networkId path parameter. Network ID
@param deviceID deviceId path parameter. Device ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sm-device-certs
*/
func (s *SmService) GetNetworkSmDeviceCerts(networkID string, deviceID string) (*ResponseSmGetNetworkSmDeviceCerts, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/{deviceId}/certs"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSmGetNetworkSmDeviceCerts{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSmDeviceCerts")
	}

	result := response.Result().(*ResponseSmGetNetworkSmDeviceCerts)
	return result, response, err

}

//GetNetworkSmDeviceConnectivity Returns historical connectivity data (whether a device is regularly checking in to Dashboard).
/* Returns historical connectivity data (whether a device is regularly checking in to Dashboard).

@param networkID networkId path parameter. Network ID
@param deviceID deviceId path parameter. Device ID
@param getNetworkSmDeviceConnectivityQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sm-device-connectivity
*/
func (s *SmService) GetNetworkSmDeviceConnectivity(networkID string, deviceID string, getNetworkSmDeviceConnectivityQueryParams *GetNetworkSmDeviceConnectivityQueryParams) (*ResponseSmGetNetworkSmDeviceConnectivity, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/{deviceId}/connectivity"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	queryString, _ := query.Values(getNetworkSmDeviceConnectivityQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSmGetNetworkSmDeviceConnectivity{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSmDeviceConnectivity")
	}

	result := response.Result().(*ResponseSmGetNetworkSmDeviceConnectivity)
	return result, response, err

}

//GetNetworkSmDeviceDesktopLogs Return historical records of various Systems Manager network connection details for desktop devices.
/* Return historical records of various Systems Manager network connection details for desktop devices.

@param networkID networkId path parameter. Network ID
@param deviceID deviceId path parameter. Device ID
@param getNetworkSmDeviceDesktopLogsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sm-device-desktop-logs
*/
func (s *SmService) GetNetworkSmDeviceDesktopLogs(networkID string, deviceID string, getNetworkSmDeviceDesktopLogsQueryParams *GetNetworkSmDeviceDesktopLogsQueryParams) (*ResponseSmGetNetworkSmDeviceDesktopLogs, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/{deviceId}/desktopLogs"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	queryString, _ := query.Values(getNetworkSmDeviceDesktopLogsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSmGetNetworkSmDeviceDesktopLogs{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSmDeviceDesktopLogs")
	}

	result := response.Result().(*ResponseSmGetNetworkSmDeviceDesktopLogs)
	return result, response, err

}

//GetNetworkSmDeviceDeviceCommandLogs Return historical records of commands sent to Systems Manager devices
/* Return historical records of commands sent to Systems Manager devices. Note that this will include the name of the Dashboard user who initiated the command if it was generated by a Dashboard admin rather than the automatic behavior of the system; you may wish to filter this out of any reports.

@param networkID networkId path parameter. Network ID
@param deviceID deviceId path parameter. Device ID
@param getNetworkSmDeviceDeviceCommandLogsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sm-device-device-command-logs
*/
func (s *SmService) GetNetworkSmDeviceDeviceCommandLogs(networkID string, deviceID string, getNetworkSmDeviceDeviceCommandLogsQueryParams *GetNetworkSmDeviceDeviceCommandLogsQueryParams) (*ResponseSmGetNetworkSmDeviceDeviceCommandLogs, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/{deviceId}/deviceCommandLogs"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	queryString, _ := query.Values(getNetworkSmDeviceDeviceCommandLogsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSmGetNetworkSmDeviceDeviceCommandLogs{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSmDeviceDeviceCommandLogs")
	}

	result := response.Result().(*ResponseSmGetNetworkSmDeviceDeviceCommandLogs)
	return result, response, err

}

//GetNetworkSmDeviceDeviceProfiles Get the installed profiles associated with a device
/* Get the installed profiles associated with a device

@param networkID networkId path parameter. Network ID
@param deviceID deviceId path parameter. Device ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sm-device-device-profiles
*/
func (s *SmService) GetNetworkSmDeviceDeviceProfiles(networkID string, deviceID string) (*ResponseSmGetNetworkSmDeviceDeviceProfiles, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/{deviceId}/deviceProfiles"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSmGetNetworkSmDeviceDeviceProfiles{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSmDeviceDeviceProfiles")
	}

	result := response.Result().(*ResponseSmGetNetworkSmDeviceDeviceProfiles)
	return result, response, err

}

//GetNetworkSmDeviceNetworkAdapters List the network adapters of a device
/* List the network adapters of a device

@param networkID networkId path parameter. Network ID
@param deviceID deviceId path parameter. Device ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sm-device-network-adapters
*/
func (s *SmService) GetNetworkSmDeviceNetworkAdapters(networkID string, deviceID string) (*ResponseSmGetNetworkSmDeviceNetworkAdapters, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/{deviceId}/networkAdapters"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSmGetNetworkSmDeviceNetworkAdapters{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSmDeviceNetworkAdapters")
	}

	result := response.Result().(*ResponseSmGetNetworkSmDeviceNetworkAdapters)
	return result, response, err

}

//GetNetworkSmDevicePerformanceHistory Return historical records of various Systems Manager client metrics for desktop devices.
/* Return historical records of various Systems Manager client metrics for desktop devices.

@param networkID networkId path parameter. Network ID
@param deviceID deviceId path parameter. Device ID
@param getNetworkSmDevicePerformanceHistoryQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sm-device-performance-history
*/
func (s *SmService) GetNetworkSmDevicePerformanceHistory(networkID string, deviceID string, getNetworkSmDevicePerformanceHistoryQueryParams *GetNetworkSmDevicePerformanceHistoryQueryParams) (*ResponseSmGetNetworkSmDevicePerformanceHistory, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/{deviceId}/performanceHistory"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	queryString, _ := query.Values(getNetworkSmDevicePerformanceHistoryQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSmGetNetworkSmDevicePerformanceHistory{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSmDevicePerformanceHistory")
	}

	result := response.Result().(*ResponseSmGetNetworkSmDevicePerformanceHistory)
	return result, response, err

}

//GetNetworkSmDeviceRestrictions List the restrictions on a device
/* List the restrictions on a device

@param networkID networkId path parameter. Network ID
@param deviceID deviceId path parameter. Device ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sm-device-restrictions
*/
func (s *SmService) GetNetworkSmDeviceRestrictions(networkID string, deviceID string) (*ResponseSmGetNetworkSmDeviceRestrictions, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/{deviceId}/restrictions"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSmGetNetworkSmDeviceRestrictions{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSmDeviceRestrictions")
	}

	result := response.Result().(*ResponseSmGetNetworkSmDeviceRestrictions)
	return result, response, err

}

//GetNetworkSmDeviceSecurityCenters List the security centers on a device
/* List the security centers on a device

@param networkID networkId path parameter. Network ID
@param deviceID deviceId path parameter. Device ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sm-device-security-centers
*/
func (s *SmService) GetNetworkSmDeviceSecurityCenters(networkID string, deviceID string) (*ResponseSmGetNetworkSmDeviceSecurityCenters, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/{deviceId}/securityCenters"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSmGetNetworkSmDeviceSecurityCenters{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSmDeviceSecurityCenters")
	}

	result := response.Result().(*ResponseSmGetNetworkSmDeviceSecurityCenters)
	return result, response, err

}

//GetNetworkSmDeviceSoftwares Get a list of softwares associated with a device
/* Get a list of softwares associated with a device

@param networkID networkId path parameter. Network ID
@param deviceID deviceId path parameter. Device ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sm-device-softwares
*/
func (s *SmService) GetNetworkSmDeviceSoftwares(networkID string, deviceID string) (*ResponseSmGetNetworkSmDeviceSoftwares, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/{deviceId}/softwares"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSmGetNetworkSmDeviceSoftwares{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSmDeviceSoftwares")
	}

	result := response.Result().(*ResponseSmGetNetworkSmDeviceSoftwares)
	return result, response, err

}

//GetNetworkSmDeviceWLANLists List the saved SSID names on a device
/* List the saved SSID names on a device

@param networkID networkId path parameter. Network ID
@param deviceID deviceId path parameter. Device ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sm-device-wlan-lists
*/
func (s *SmService) GetNetworkSmDeviceWLANLists(networkID string, deviceID string) (*ResponseSmGetNetworkSmDeviceWLANLists, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/{deviceId}/wlanLists"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSmGetNetworkSmDeviceWLANLists{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSmDeviceWlanLists")
	}

	result := response.Result().(*ResponseSmGetNetworkSmDeviceWLANLists)
	return result, response, err

}

//GetNetworkSmProfiles List all profiles in a network
/* List all profiles in a network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sm-profiles
*/
func (s *SmService) GetNetworkSmProfiles(networkID string) (*ResponseSmGetNetworkSmProfiles, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/profiles"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSmGetNetworkSmProfiles{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSmProfiles")
	}

	result := response.Result().(*ResponseSmGetNetworkSmProfiles)
	return result, response, err

}

//GetNetworkSmTargetGroups List the target groups in this network
/* List the target groups in this network

@param networkID networkId path parameter. Network ID
@param getNetworkSmTargetGroupsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sm-target-groups
*/
func (s *SmService) GetNetworkSmTargetGroups(networkID string, getNetworkSmTargetGroupsQueryParams *GetNetworkSmTargetGroupsQueryParams) (*ResponseSmGetNetworkSmTargetGroups, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/targetGroups"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(getNetworkSmTargetGroupsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSmGetNetworkSmTargetGroups{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSmTargetGroups")
	}

	result := response.Result().(*ResponseSmGetNetworkSmTargetGroups)
	return result, response, err

}

//GetNetworkSmTargetGroup Return a target group
/* Return a target group

@param networkID networkId path parameter. Network ID
@param targetGroupID targetGroupId path parameter. Target group ID
@param getNetworkSmTargetGroupQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sm-target-group
*/
func (s *SmService) GetNetworkSmTargetGroup(networkID string, targetGroupID string, getNetworkSmTargetGroupQueryParams *GetNetworkSmTargetGroupQueryParams) (*ResponseSmGetNetworkSmTargetGroup, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/targetGroups/{targetGroupId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{targetGroupId}", fmt.Sprintf("%v", targetGroupID), -1)

	queryString, _ := query.Values(getNetworkSmTargetGroupQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSmGetNetworkSmTargetGroup{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSmTargetGroup")
	}

	result := response.Result().(*ResponseSmGetNetworkSmTargetGroup)
	return result, response, err

}

//GetNetworkSmTrustedAccessConfigs List Trusted Access Configs
/* List Trusted Access Configs

@param networkID networkId path parameter. Network ID
@param getNetworkSmTrustedAccessConfigsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sm-trusted-access-configs
*/
func (s *SmService) GetNetworkSmTrustedAccessConfigs(networkID string, getNetworkSmTrustedAccessConfigsQueryParams *GetNetworkSmTrustedAccessConfigsQueryParams) (*ResponseSmGetNetworkSmTrustedAccessConfigs, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/trustedAccessConfigs"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(getNetworkSmTrustedAccessConfigsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSmGetNetworkSmTrustedAccessConfigs{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSmTrustedAccessConfigs")
	}

	result := response.Result().(*ResponseSmGetNetworkSmTrustedAccessConfigs)
	return result, response, err

}

//GetNetworkSmUserAccessDevices List User Access Devices and its Trusted Access Connections
/* List User Access Devices and its Trusted Access Connections

@param networkID networkId path parameter. Network ID
@param getNetworkSmUserAccessDevicesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sm-user-access-devices
*/
func (s *SmService) GetNetworkSmUserAccessDevices(networkID string, getNetworkSmUserAccessDevicesQueryParams *GetNetworkSmUserAccessDevicesQueryParams) (*ResponseSmGetNetworkSmUserAccessDevices, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/userAccessDevices"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(getNetworkSmUserAccessDevicesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSmGetNetworkSmUserAccessDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSmUserAccessDevices")
	}

	result := response.Result().(*ResponseSmGetNetworkSmUserAccessDevices)
	return result, response, err

}

//GetNetworkSmUsers List the owners in an SM network with various specified fields and filters
/* List the owners in an SM network with various specified fields and filters

@param networkID networkId path parameter. Network ID
@param getNetworkSmUsersQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sm-users
*/
func (s *SmService) GetNetworkSmUsers(networkID string, getNetworkSmUsersQueryParams *GetNetworkSmUsersQueryParams) (*ResponseSmGetNetworkSmUsers, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/users"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(getNetworkSmUsersQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSmGetNetworkSmUsers{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSmUsers")
	}

	result := response.Result().(*ResponseSmGetNetworkSmUsers)
	return result, response, err

}

//GetNetworkSmUserDeviceProfiles Get the profiles associated with a user
/* Get the profiles associated with a user

@param networkID networkId path parameter. Network ID
@param userID userId path parameter. User ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sm-user-device-profiles
*/
func (s *SmService) GetNetworkSmUserDeviceProfiles(networkID string, userID string) (*ResponseSmGetNetworkSmUserDeviceProfiles, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/users/{userId}/deviceProfiles"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{userId}", fmt.Sprintf("%v", userID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSmGetNetworkSmUserDeviceProfiles{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSmUserDeviceProfiles")
	}

	result := response.Result().(*ResponseSmGetNetworkSmUserDeviceProfiles)
	return result, response, err

}

//GetNetworkSmUserSoftwares Get a list of softwares associated with a user
/* Get a list of softwares associated with a user

@param networkID networkId path parameter. Network ID
@param userID userId path parameter. User ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sm-user-softwares
*/
func (s *SmService) GetNetworkSmUserSoftwares(networkID string, userID string) (*ResponseSmGetNetworkSmUserSoftwares, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/users/{userId}/softwares"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{userId}", fmt.Sprintf("%v", userID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSmGetNetworkSmUserSoftwares{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSmUserSoftwares")
	}

	result := response.Result().(*ResponseSmGetNetworkSmUserSoftwares)
	return result, response, err

}

//GetOrganizationSmApnsCert Get the organization's APNS certificate
/* Get the organization's APNS certificate

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-sm-apns-cert
*/
func (s *SmService) GetOrganizationSmApnsCert(organizationID string) (*ResponseSmGetOrganizationSmApnsCert, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/sm/apnsCert"
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSmGetOrganizationSmApnsCert{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationSmApnsCert")
	}

	result := response.Result().(*ResponseSmGetOrganizationSmApnsCert)
	return result, response, err

}

//GetOrganizationSmVppAccounts List the VPP accounts in the organization
/* List the VPP accounts in the organization

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-sm-vpp-accounts
*/
func (s *SmService) GetOrganizationSmVppAccounts(organizationID string) (*ResponseSmGetOrganizationSmVppAccounts, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/sm/vppAccounts"
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSmGetOrganizationSmVppAccounts{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationSmVppAccounts")
	}

	result := response.Result().(*ResponseSmGetOrganizationSmVppAccounts)
	return result, response, err

}

//GetOrganizationSmVppAccount Get a hash containing the unparsed token of the VPP account with the given ID
/* Get a hash containing the unparsed token of the VPP account with the given ID

@param organizationID organizationId path parameter. Organization ID
@param vppAccountID vppAccountId path parameter. Vpp account ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-sm-vpp-account
*/
func (s *SmService) GetOrganizationSmVppAccount(organizationID string, vppAccountID string) (*ResponseSmGetOrganizationSmVppAccount, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/sm/vppAccounts/{vppAccountId}"
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{vppAccountId}", fmt.Sprintf("%v", vppAccountID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSmGetOrganizationSmVppAccount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationSmVppAccount")
	}

	result := response.Result().(*ResponseSmGetOrganizationSmVppAccount)
	return result, response, err

}

//CreateNetworkSmBypassActivationLockAttempt Bypass activation lock attempt
/* Bypass activation lock attempt

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-network-sm-bypass-activation-lock-attempt
*/

func (s *SmService) CreateNetworkSmBypassActivationLockAttempt(networkID string, requestSmCreateNetworkSmBypassActivationLockAttempt *RequestSmCreateNetworkSmBypassActivationLockAttempt) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/bypassActivationLockAttempts"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSmCreateNetworkSmBypassActivationLockAttempt).
		// SetResult(&ResponseSmCreateNetworkSmBypassActivationLockAttempt{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateNetworkSmBypassActivationLockAttempt")
	}

	return response, err

}

//CheckinNetworkSmDevices Force check-in a set of devices
/* Force check-in a set of devices

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!checkin-network-sm-devices
*/

func (s *SmService) CheckinNetworkSmDevices(networkID string, requestSmCheckinNetworkSmDevices *RequestSmCheckinNetworkSmDevices) (*ResponseSmCheckinNetworkSmDevices, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/checkin"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSmCheckinNetworkSmDevices).
		SetResult(&ResponseSmCheckinNetworkSmDevices{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CheckinNetworkSmDevices")
	}

	result := response.Result().(*ResponseSmCheckinNetworkSmDevices)
	return result, response, err

}

//LockNetworkSmDevices Lock a set of devices
/* Lock a set of devices

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!lock-network-sm-devices
*/

func (s *SmService) LockNetworkSmDevices(networkID string, requestSmLockNetworkSmDevices *RequestSmLockNetworkSmDevices) (*ResponseSmLockNetworkSmDevices, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/lock"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSmLockNetworkSmDevices).
		SetResult(&ResponseSmLockNetworkSmDevices{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation LockNetworkSmDevices")
	}

	result := response.Result().(*ResponseSmLockNetworkSmDevices)
	return result, response, err

}

//ModifyNetworkSmDevicesTags Add, delete, or update the tags of a set of devices
/* Add, delete, or update the tags of a set of devices

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!modify-network-sm-devices-tags
*/

func (s *SmService) ModifyNetworkSmDevicesTags(networkID string, requestSmModifyNetworkSmDevicesTags *RequestSmModifyNetworkSmDevicesTags) (*ResponseSmModifyNetworkSmDevicesTags, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/modifyTags"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSmModifyNetworkSmDevicesTags).
		SetResult(&ResponseSmModifyNetworkSmDevicesTags{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ModifyNetworkSmDevicesTags")
	}

	result := response.Result().(*ResponseSmModifyNetworkSmDevicesTags)
	return result, response, err

}

//MoveNetworkSmDevices Move a set of devices to a new network
/* Move a set of devices to a new network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!move-network-sm-devices
*/

func (s *SmService) MoveNetworkSmDevices(networkID string, requestSmMoveNetworkSmDevices *RequestSmMoveNetworkSmDevices) (*ResponseSmMoveNetworkSmDevices, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/move"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSmMoveNetworkSmDevices).
		SetResult(&ResponseSmMoveNetworkSmDevices{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation MoveNetworkSmDevices")
	}

	result := response.Result().(*ResponseSmMoveNetworkSmDevices)
	return result, response, err

}

//WipeNetworkSmDevices Wipe a device
/* Wipe a device

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!wipe-network-sm-devices
*/

func (s *SmService) WipeNetworkSmDevices(networkID string, requestSmWipeNetworkSmDevices *RequestSmWipeNetworkSmDevices) (*ResponseSmWipeNetworkSmDevices, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/wipe"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSmWipeNetworkSmDevices).
		SetResult(&ResponseSmWipeNetworkSmDevices{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation WipeNetworkSmDevices")
	}

	result := response.Result().(*ResponseSmWipeNetworkSmDevices)
	return result, response, err

}

//RefreshNetworkSmDeviceDetails Refresh the details of a device
/* Refresh the details of a device

@param networkID networkId path parameter. Network ID
@param deviceID deviceId path parameter. Device ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!refresh-network-sm-device-details
*/

func (s *SmService) RefreshNetworkSmDeviceDetails(networkID string, deviceID string) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/{deviceId}/refreshDetails"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation RefreshNetworkSmDeviceDetails")
	}

	return response, err

}

//UnenrollNetworkSmDevice Unenroll a device
/* Unenroll a device

@param networkID networkId path parameter. Network ID
@param deviceID deviceId path parameter. Device ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!unenroll-network-sm-device
*/

func (s *SmService) UnenrollNetworkSmDevice(networkID string, deviceID string) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/{deviceId}/unenroll"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").

		// SetResult(&ResponseSmUnenrollNetworkSmDevice{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UnenrollNetworkSmDevice")
	}

	return response, err

}

//CreateNetworkSmTargetGroup Add a target group
/* Add a target group

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-network-sm-target-group
*/

func (s *SmService) CreateNetworkSmTargetGroup(networkID string, requestSmCreateNetworkSmTargetGroup *RequestSmCreateNetworkSmTargetGroup) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/targetGroups"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSmCreateNetworkSmTargetGroup).
		// SetResult(&ResponseSmCreateNetworkSmTargetGroup{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateNetworkSmTargetGroup")
	}

	return response, err

}

//UpdateNetworkSmDevicesFields Modify the fields of a device
/* Modify the fields of a device

@param networkID networkId path parameter. Network ID
*/
func (s *SmService) UpdateNetworkSmDevicesFields(networkID string, requestSmUpdateNetworkSmDevicesFields *RequestSmUpdateNetworkSmDevicesFields) (*ResponseSmUpdateNetworkSmDevicesFields, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/fields"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSmUpdateNetworkSmDevicesFields).
		SetResult(&ResponseSmUpdateNetworkSmDevicesFields{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkSmDevicesFields")
	}

	result := response.Result().(*ResponseSmUpdateNetworkSmDevicesFields)
	return result, response, err

}

//UpdateNetworkSmTargetGroup Update a target group
/* Update a target group

@param networkID networkId path parameter. Network ID
@param targetGroupID targetGroupId path parameter. Target group ID
*/
func (s *SmService) UpdateNetworkSmTargetGroup(networkID string, targetGroupID string, requestSmUpdateNetworkSmTargetGroup *RequestSmUpdateNetworkSmTargetGroup) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/targetGroups/{targetGroupId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{targetGroupId}", fmt.Sprintf("%v", targetGroupID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSmUpdateNetworkSmTargetGroup).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkSmTargetGroup")
	}

	return response, err

}

//DeleteNetworkSmTargetGroup Delete a target group from a network
/* Delete a target group from a network

@param networkID networkId path parameter. Network ID
@param targetGroupID targetGroupId path parameter. Target group ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-network-sm-target-group
*/
func (s *SmService) DeleteNetworkSmTargetGroup(networkID string, targetGroupID string) (*resty.Response, error) {
	//networkID string,targetGroupID string
	path := "/api/v1/networks/{networkId}/sm/targetGroups/{targetGroupId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{targetGroupId}", fmt.Sprintf("%v", targetGroupID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteNetworkSmTargetGroup")
	}

	return response, err

}

//DeleteNetworkSmUserAccessDevice Delete a User Access Device
/* Delete a User Access Device

@param networkID networkId path parameter. Network ID
@param userAccessDeviceID userAccessDeviceId path parameter. User access device ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-network-sm-user-access-device
*/
func (s *SmService) DeleteNetworkSmUserAccessDevice(networkID string, userAccessDeviceID string) (*resty.Response, error) {
	//networkID string,userAccessDeviceID string
	path := "/api/v1/networks/{networkId}/sm/userAccessDevices/{userAccessDeviceId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{userAccessDeviceId}", fmt.Sprintf("%v", userAccessDeviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteNetworkSmUserAccessDevice")
	}

	return response, err

}
