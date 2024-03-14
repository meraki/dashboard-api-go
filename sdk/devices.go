package meraki

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type DevicesService service

type GetDeviceClientsQueryParams struct {
	T0       string  `url:"t0,omitempty"`       //The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	Timespan float64 `url:"timespan,omitempty"` //The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
}
type GetDeviceLossAndLatencyHistoryQueryParams struct {
	T0         string  `url:"t0,omitempty"`         //The beginning of the timespan for the data. The maximum lookback period is 60 days from today.
	T1         string  `url:"t1,omitempty"`         //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan   float64 `url:"timespan,omitempty"`   //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
	Resolution int     `url:"resolution,omitempty"` //The time resolution in seconds for returned data. The valid resolutions are: 60, 600, 3600, 86400. The default is 60.
	Uplink     string  `url:"uplink,omitempty"`     //The WAN uplink used to obtain the requested stats. Valid uplinks are wan1, wan2, cellular. The default is wan1.
	IP         string  `url:"ip,omitempty"`         //The destination IP used to obtain the requested stats. This is required.
}

type ResponseDevicesGetDevice struct {
	Address        string                                  `json:"address,omitempty"`        //
	BeaconIDParams *ResponseDevicesGetDeviceBeaconIDParams `json:"beaconIdParams,omitempty"` //
	Firmware       string                                  `json:"firmware,omitempty"`       //
	FloorPlanID    string                                  `json:"floorPlanId,omitempty"`    //
	LanIP          string                                  `json:"lanIp,omitempty"`          //
	Lat            *float64                                `json:"lat,omitempty"`            //
	Lng            *float64                                `json:"lng,omitempty"`            //
	Mac            string                                  `json:"mac,omitempty"`            //
	Model          string                                  `json:"model,omitempty"`          //
	Name           string                                  `json:"name,omitempty"`           //
	NetworkID      string                                  `json:"networkId,omitempty"`      //
	Notes          string                                  `json:"notes,omitempty"`          //
	Serial         string                                  `json:"serial,omitempty"`         //
	Tags           []string                                `json:"tags,omitempty"`           //
}
type ResponseDevicesGetDeviceBeaconIDParams struct {
	Major *int   `json:"major,omitempty"` //
	Minor *int   `json:"minor,omitempty"` //
	UUID  string `json:"uuid,omitempty"`  //
}
type ResponseDevicesUpdateDevice interface{}
type ResponseDevicesBlinkDeviceLeds struct {
	Duration *int `json:"duration,omitempty"` //
	Duty     *int `json:"duty,omitempty"`     //
	Period   *int `json:"period,omitempty"`   //
}
type ResponseDevicesGetDeviceCellularSims struct {
	Sims *[]ResponseDevicesGetDeviceCellularSimsSims `json:"sims,omitempty"` //
}
type ResponseDevicesGetDeviceCellularSimsSims struct {
	Apns      *[]ResponseDevicesGetDeviceCellularSimsSimsApns `json:"apns,omitempty"`      //
	IsPrimary *bool                                           `json:"isPrimary,omitempty"` //
	Slot      string                                          `json:"slot,omitempty"`      //
}
type ResponseDevicesGetDeviceCellularSimsSimsApns struct {
	AllowedIPTypes []string                                                    `json:"allowedIpTypes,omitempty"` //
	Authentication *ResponseDevicesGetDeviceCellularSimsSimsApnsAuthentication `json:"authentication,omitempty"` //
	Name           string                                                      `json:"name,omitempty"`           //
}
type ResponseDevicesGetDeviceCellularSimsSimsApnsAuthentication struct {
	Type     string `json:"type,omitempty"`     //
	Username string `json:"username,omitempty"` //
}
type ResponseDevicesUpdateDeviceCellularSims interface{}
type ResponseDevicesGetDeviceClients []ResponseItemDevicesGetDeviceClients // Array of ResponseDevicesGetDeviceClients
type ResponseItemDevicesGetDeviceClients struct {
	Description  string                                    `json:"description,omitempty"`  //
	DhcpHostname string                                    `json:"dhcpHostname,omitempty"` //
	ID           string                                    `json:"id,omitempty"`           //
	IP           string                                    `json:"ip,omitempty"`           //
	Mac          string                                    `json:"mac,omitempty"`          //
	MdnsName     string                                    `json:"mdnsName,omitempty"`     //
	NamedVLAN    string                                    `json:"namedVlan,omitempty"`    //
	Usage        *ResponseItemDevicesGetDeviceClientsUsage `json:"usage,omitempty"`        //
	User         string                                    `json:"user,omitempty"`         //
	VLAN         string                                    `json:"vlan,omitempty"`         //
}
type ResponseItemDevicesGetDeviceClientsUsage struct {
	Recv *int `json:"recv,omitempty"` //
	Sent *int `json:"sent,omitempty"` //
}
type ResponseDevicesCreateDeviceLiveToolsPing struct {
	PingID  string                                           `json:"pingId,omitempty"`  // Id to check the status of your ping request.
	Request *ResponseDevicesCreateDeviceLiveToolsPingRequest `json:"request,omitempty"` // Ping request parameters
	Status  string                                           `json:"status,omitempty"`  // Status of the ping request.
	URL     string                                           `json:"url,omitempty"`     // GET this url to check the status of your ping request.
}
type ResponseDevicesCreateDeviceLiveToolsPingRequest struct {
	Count  *int   `json:"count,omitempty"`  // Number of pings to send
	Serial string `json:"serial,omitempty"` // Device serial number
	Target string `json:"target,omitempty"` // IP address or FQDN to ping
}
type ResponseDevicesGetDeviceLiveToolsPing struct {
	PingID  string                                        `json:"pingId,omitempty"`  // Id to check the status of your ping request.
	Request *ResponseDevicesGetDeviceLiveToolsPingRequest `json:"request,omitempty"` // Ping request parameters
	Results *ResponseDevicesGetDeviceLiveToolsPingResults `json:"results,omitempty"` // Results of the ping request.
	Status  string                                        `json:"status,omitempty"`  // Status of the ping request.
	URL     string                                        `json:"url,omitempty"`     // GET this url to check the status of your ping request.
}
type ResponseDevicesGetDeviceLiveToolsPingRequest struct {
	Count  *int   `json:"count,omitempty"`  // Number of pings to send
	Serial string `json:"serial,omitempty"` // Device serial number
	Target string `json:"target,omitempty"` // IP address or FQDN to ping
}
type ResponseDevicesGetDeviceLiveToolsPingResults struct {
	Latencies *ResponseDevicesGetDeviceLiveToolsPingResultsLatencies `json:"latencies,omitempty"` // Packet latency stats
	Loss      *ResponseDevicesGetDeviceLiveToolsPingResultsLoss      `json:"loss,omitempty"`      // Lost packets
	Received  *int                                                   `json:"received,omitempty"`  // Number of packets received
	Replies   *[]ResponseDevicesGetDeviceLiveToolsPingResultsReplies `json:"replies,omitempty"`   // Received packets
	Sent      *int                                                   `json:"sent,omitempty"`      // Number of packets sent
}
type ResponseDevicesGetDeviceLiveToolsPingResultsLatencies struct {
	Average *float64 `json:"average,omitempty"` // Average latency
	Maximum *float64 `json:"maximum,omitempty"` // Maximum latency
	Minimum *float64 `json:"minimum,omitempty"` // Minimum latency
}
type ResponseDevicesGetDeviceLiveToolsPingResultsLoss struct {
	Percentage *float64 `json:"percentage,omitempty"` // Percentage of packets lost
}
type ResponseDevicesGetDeviceLiveToolsPingResultsReplies struct {
	Latency    *float64 `json:"latency,omitempty"`    // Latency of the packet in milliseconds
	SequenceID *int     `json:"sequenceId,omitempty"` // Sequence ID of the packet
	Size       *int     `json:"size,omitempty"`       // Size of the packet in bytes
}
type ResponseDevicesCreateDeviceLiveToolsPingDevice struct {
	PingID  string                                                 `json:"pingId,omitempty"`  // Id to check the status of your ping request.
	Request *ResponseDevicesCreateDeviceLiveToolsPingDeviceRequest `json:"request,omitempty"` // Ping request parameters
	Status  string                                                 `json:"status,omitempty"`  // Status of the ping request.
	URL     string                                                 `json:"url,omitempty"`     // GET this url to check the status of your ping request.
}
type ResponseDevicesCreateDeviceLiveToolsPingDeviceRequest struct {
	Count  *int   `json:"count,omitempty"`  // Number of pings to send
	Serial string `json:"serial,omitempty"` // Device serial number
	Target string `json:"target,omitempty"` // IP address or FQDN to ping
}
type ResponseDevicesGetDeviceLiveToolsPingDevice struct {
	PingID  string                                              `json:"pingId,omitempty"`  // Id to check the status of your ping request.
	Request *ResponseDevicesGetDeviceLiveToolsPingDeviceRequest `json:"request,omitempty"` // Ping request parameters
	Results *ResponseDevicesGetDeviceLiveToolsPingDeviceResults `json:"results,omitempty"` // Results of the ping request.
	Status  string                                              `json:"status,omitempty"`  // Status of the ping request.
	URL     string                                              `json:"url,omitempty"`     // GET this url to check the status of your ping request.
}
type ResponseDevicesGetDeviceLiveToolsPingDeviceRequest struct {
	Count  *int   `json:"count,omitempty"`  // Number of pings to send
	Serial string `json:"serial,omitempty"` // Device serial number
	Target string `json:"target,omitempty"` // IP address or FQDN to ping
}
type ResponseDevicesGetDeviceLiveToolsPingDeviceResults struct {
	Latencies *ResponseDevicesGetDeviceLiveToolsPingDeviceResultsLatencies `json:"latencies,omitempty"` // Packet latency stats
	Loss      *ResponseDevicesGetDeviceLiveToolsPingDeviceResultsLoss      `json:"loss,omitempty"`      // Lost packets
	Received  *int                                                         `json:"received,omitempty"`  // Number of packets received
	Replies   *[]ResponseDevicesGetDeviceLiveToolsPingDeviceResultsReplies `json:"replies,omitempty"`   // Received packets
	Sent      *int                                                         `json:"sent,omitempty"`      // Number of packets sent
}
type ResponseDevicesGetDeviceLiveToolsPingDeviceResultsLatencies struct {
	Average *float64 `json:"average,omitempty"` // Average latency
	Maximum *float64 `json:"maximum,omitempty"` // Maximum latency
	Minimum *float64 `json:"minimum,omitempty"` // Minimum latency
}
type ResponseDevicesGetDeviceLiveToolsPingDeviceResultsLoss struct {
	Percentage *float64 `json:"percentage,omitempty"` // Percentage of packets lost
}
type ResponseDevicesGetDeviceLiveToolsPingDeviceResultsReplies struct {
	Latency    *float64 `json:"latency,omitempty"`    // Latency of the packet in milliseconds
	SequenceID *int     `json:"sequenceId,omitempty"` // Sequence ID of the packet
	Size       *int     `json:"size,omitempty"`       // Size of the packet in bytes
}
type ResponseDevicesGetDeviceLldpCdp struct {
	Ports     *ResponseDevicesGetDeviceLldpCdpPorts `json:"ports,omitempty"`     //
	SourceMac string                                `json:"sourceMac,omitempty"` //
}
type ResponseDevicesGetDeviceLldpCdpPorts struct {
	Status12 *ResponseDevicesGetDeviceLldpCdpPorts12 `json:"12,omitempty"` //
	Status8  *ResponseDevicesGetDeviceLldpCdpPorts8  `json:"8,omitempty"`  //
}
type ResponseDevicesGetDeviceLldpCdpPorts12 struct {
	Cdp  *ResponseDevicesGetDeviceLldpCdpPorts12Cdp  `json:"cdp,omitempty"`  //
	Lldp *ResponseDevicesGetDeviceLldpCdpPorts12Lldp `json:"lldp,omitempty"` //
}
type ResponseDevicesGetDeviceLldpCdpPorts12Cdp struct {
	Address    string `json:"address,omitempty"`    //
	DeviceID   string `json:"deviceId,omitempty"`   //
	PortID     string `json:"portId,omitempty"`     //
	SourcePort string `json:"sourcePort,omitempty"` //
}
type ResponseDevicesGetDeviceLldpCdpPorts12Lldp struct {
	ManagementAddress string `json:"managementAddress,omitempty"` //
	PortID            string `json:"portId,omitempty"`            //
	SourcePort        string `json:"sourcePort,omitempty"`        //
	SystemName        string `json:"systemName,omitempty"`        //
}
type ResponseDevicesGetDeviceLldpCdpPorts8 struct {
	Cdp *ResponseDevicesGetDeviceLldpCdpPorts8Cdp `json:"cdp,omitempty"` //
}
type ResponseDevicesGetDeviceLldpCdpPorts8Cdp struct {
	Address    string `json:"address,omitempty"`    //
	DeviceID   string `json:"deviceId,omitempty"`   //
	PortID     string `json:"portId,omitempty"`     //
	SourcePort string `json:"sourcePort,omitempty"` //
}
type ResponseDevicesGetDeviceLossAndLatencyHistory []ResponseItemDevicesGetDeviceLossAndLatencyHistory // Array of ResponseDevicesGetDeviceLossAndLatencyHistory
type ResponseItemDevicesGetDeviceLossAndLatencyHistory struct {
	EndTime     string   `json:"endTime,omitempty"`     //
	Goodput     *int     `json:"goodput,omitempty"`     //
	Jitter      *float64 `json:"jitter,omitempty"`      //
	LatencyMs   *float64 `json:"latencyMs,omitempty"`   //
	LossPercent *float64 `json:"lossPercent,omitempty"` //
	StartTime   string   `json:"startTime,omitempty"`   //
}
type ResponseDevicesGetDeviceManagementInterface struct {
	DdnsHostnames *ResponseDevicesGetDeviceManagementInterfaceDdnsHostnames `json:"ddnsHostnames,omitempty"` //
	Wan1          *ResponseDevicesGetDeviceManagementInterfaceWan1          `json:"wan1,omitempty"`          //
	Wan2          *ResponseDevicesGetDeviceManagementInterfaceWan2          `json:"wan2,omitempty"`          //
}
type ResponseDevicesGetDeviceManagementInterfaceDdnsHostnames struct {
	ActiveDdnsHostname string `json:"activeDdnsHostname,omitempty"` //
	DdnsHostnameWan1   string `json:"ddnsHostnameWan1,omitempty"`   //
	DdnsHostnameWan2   string `json:"ddnsHostnameWan2,omitempty"`   //
}
type ResponseDevicesGetDeviceManagementInterfaceWan1 struct {
	StaticDNS        []string `json:"staticDns,omitempty"`        //
	StaticGatewayIP  string   `json:"staticGatewayIp,omitempty"`  //
	StaticIP         string   `json:"staticIp,omitempty"`         //
	StaticSubnetMask string   `json:"staticSubnetMask,omitempty"` //
	UsingStaticIP    *bool    `json:"usingStaticIp,omitempty"`    //
	VLAN             *int     `json:"vlan,omitempty"`             //
	WanEnabled       string   `json:"wanEnabled,omitempty"`       //
}
type ResponseDevicesGetDeviceManagementInterfaceWan2 struct {
	UsingStaticIP *bool  `json:"usingStaticIp,omitempty"` //
	VLAN          *int   `json:"vlan,omitempty"`          //
	WanEnabled    string `json:"wanEnabled,omitempty"`    //
}
type ResponseDevicesUpdateDeviceManagementInterface interface{}
type ResponseDevicesRebootDevice interface{}
type ResponseDevicesGetNetworkDevices []ResponseItemDevicesGetNetworkDevices // Array of ResponseDevicesGetNetworkDevices
type ResponseItemDevicesGetNetworkDevices struct {
	Address        string                                              `json:"address,omitempty"`        //
	BeaconIDParams *ResponseItemDevicesGetNetworkDevicesBeaconIDParams `json:"beaconIdParams,omitempty"` //
	Firmware       string                                              `json:"firmware,omitempty"`       //
	FloorPlanID    string                                              `json:"floorPlanId,omitempty"`    //
	LanIP          string                                              `json:"lanIp,omitempty"`          //
	Lat            *float64                                            `json:"lat,omitempty"`            //
	Lng            *float64                                            `json:"lng,omitempty"`            //
	Mac            string                                              `json:"mac,omitempty"`            //
	Model          string                                              `json:"model,omitempty"`          //
	Name           string                                              `json:"name,omitempty"`           //
	NetworkID      string                                              `json:"networkId,omitempty"`      //
	Notes          string                                              `json:"notes,omitempty"`          //
	Serial         string                                              `json:"serial,omitempty"`         //
	Tags           string                                              `json:"tags,omitempty"`           //
}
type ResponseItemDevicesGetNetworkDevicesBeaconIDParams struct {
	Major *int   `json:"major,omitempty"` //
	Minor *int   `json:"minor,omitempty"` //
	UUID  string `json:"uuid,omitempty"`  //
}
type ResponseDevicesVmxNetworkDevicesClaim interface{}
type ResponseDevicesCheckinNetworkSmDevices struct {
	IDs []string `json:"ids,omitempty"` // The Meraki Ids of the set of devices.
}
type ResponseDevicesUpdateNetworkSmDevicesFields []ResponseItemDevicesUpdateNetworkSmDevicesFields // Array of ResponseDevicesUpdateNetworkSmDevicesFields
type ResponseItemDevicesUpdateNetworkSmDevicesFields struct {
	ID      string `json:"id,omitempty"`      // The Meraki Id of the device record.
	Name    string `json:"name,omitempty"`    // The name of the device.
	Notes   string `json:"notes,omitempty"`   // Notes associated with the device.
	Serial  string `json:"serial,omitempty"`  // The device serial.
	WifiMac string `json:"wifiMac,omitempty"` // The MAC of the device.
}
type ResponseDevicesLockNetworkSmDevices struct {
	IDs []string `json:"ids,omitempty"` // The Meraki Ids of the set of devices.
}
type ResponseDevicesModifyNetworkSmDevicesTags []ResponseItemDevicesModifyNetworkSmDevicesTags // Array of ResponseDevicesModifyNetworkSmDevicesTags
type ResponseItemDevicesModifyNetworkSmDevicesTags struct {
	ID      string   `json:"id,omitempty"`      // The Meraki Id of the device record.
	Serial  string   `json:"serial,omitempty"`  // The device serial.
	Tags    []string `json:"tags,omitempty"`    // An array of tags associated with the device.
	WifiMac string   `json:"wifiMac,omitempty"` // The MAC of the device.
}
type ResponseDevicesMoveNetworkSmDevices struct {
	IDs        []string `json:"ids,omitempty"`        // The Meraki Ids of the set of devices.
	NewNetwork string   `json:"newNetwork,omitempty"` // The network to which the devices was moved.
}
type ResponseDevicesWipeNetworkSmDevices struct {
	ID string `json:"id,omitempty"` // The Meraki Id of the devices.
}
type ResponseDevicesGetNetworkSmDeviceCellularUsageHistory []ResponseItemDevicesGetNetworkSmDeviceCellularUsageHistory // Array of ResponseDevicesGetNetworkSmDeviceCellularUsageHistory
type ResponseItemDevicesGetNetworkSmDeviceCellularUsageHistory struct {
	Received *float64 `json:"received,omitempty"` // The amount of cellular data received by the device.
	Sent     *float64 `json:"sent,omitempty"`     // The amount of cellular sent received by the device.
	Ts       string   `json:"ts,omitempty"`       // When the cellular usage data was collected.
}
type ResponseDevicesGetNetworkSmDeviceCerts []ResponseItemDevicesGetNetworkSmDeviceCerts // Array of ResponseDevicesGetNetworkSmDeviceCerts
type ResponseItemDevicesGetNetworkSmDeviceCerts struct {
	CertPem        string `json:"certPem,omitempty"`        // The PEM of the certificate.
	DeviceID       string `json:"deviceId,omitempty"`       // The Meraki managed device Id.
	ID             string `json:"id,omitempty"`             // The Meraki Id of the certificate record.
	Issuer         string `json:"issuer,omitempty"`         // The certificate issuer.
	Name           string `json:"name,omitempty"`           // The name of the certificate.
	NotValidAfter  string `json:"notValidAfter,omitempty"`  // The date after which the certificate is no longer valid.
	NotValidBefore string `json:"notValidBefore,omitempty"` // The date before which the certificate is not valid.
	Subject        string `json:"subject,omitempty"`        // The subject of the certificate.
}
type ResponseDevicesGetNetworkSmDeviceDeviceProfiles []ResponseItemDevicesGetNetworkSmDeviceDeviceProfiles // Array of ResponseDevicesGetNetworkSmDeviceDeviceProfiles
type ResponseItemDevicesGetNetworkSmDeviceDeviceProfiles struct {
	DeviceID          string `json:"deviceId,omitempty"`          // The Meraki managed device Id.
	ID                string `json:"id,omitempty"`                // The numerical Meraki Id of the profile.
	IsEncrypted       *bool  `json:"isEncrypted,omitempty"`       // A boolean indicating if the profile is encrypted.
	IsManaged         *bool  `json:"isManaged,omitempty"`         // Whether or not the profile is managed by Meraki.
	Name              string `json:"name,omitempty"`              // The name of the profile.
	ProfileData       string `json:"profileData,omitempty"`       // A string containing a JSON object with the profile data.
	ProfileIDentifier string `json:"profileIdentifier,omitempty"` // The identifier of the profile.
	Version           string `json:"version,omitempty"`           // The verison of the profile.
}
type ResponseDevicesGetNetworkSmDeviceNetworkAdapters []ResponseItemDevicesGetNetworkSmDeviceNetworkAdapters // Array of ResponseDevicesGetNetworkSmDeviceNetworkAdapters
type ResponseItemDevicesGetNetworkSmDeviceNetworkAdapters struct {
	DhcpServer string `json:"dhcpServer,omitempty"` // The IP address of the DCHP Server.
	DNSServer  string `json:"dnsServer,omitempty"`  // The IP address of the DNS Server.
	Gateway    string `json:"gateway,omitempty"`    // The IP address of the Gateway.
	ID         string `json:"id,omitempty"`         // The Meraki Id of the network adapter record.
	IP         string `json:"ip,omitempty"`         // The IP address of the network adapter.
	Mac        string `json:"mac,omitempty"`        // The MAC associated with the network adapter.
	Name       string `json:"name,omitempty"`       // The name of the newtwork adapter.
	Subnet     string `json:"subnet,omitempty"`     // The subnet for the network adapter.
}
type ResponseDevicesGetNetworkSmDeviceRestrictions []ResponseItemDevicesGetNetworkSmDeviceRestrictions // Array of ResponseDevicesGetNetworkSmDeviceRestrictions
type ResponseItemDevicesGetNetworkSmDeviceRestrictions struct {
	Profile      string                                                         `json:"profile,omitempty"`      //
	Restrictions *ResponseItemDevicesGetNetworkSmDeviceRestrictionsRestrictions `json:"restrictions,omitempty"` //
}
type ResponseItemDevicesGetNetworkSmDeviceRestrictionsRestrictions struct {
	MyRestriction *ResponseItemDevicesGetNetworkSmDeviceRestrictionsRestrictionsMyRestriction `json:"myRestriction,omitempty"` //
}
type ResponseItemDevicesGetNetworkSmDeviceRestrictionsRestrictionsMyRestriction struct {
	Value *bool `json:"value,omitempty"` //
}
type ResponseDevicesGetNetworkSmDeviceSecurityCenters []ResponseItemDevicesGetNetworkSmDeviceSecurityCenters // Array of ResponseDevicesGetNetworkSmDeviceSecurityCenters
type ResponseItemDevicesGetNetworkSmDeviceSecurityCenters struct {
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
type ResponseDevicesGetNetworkSmDeviceSoftwares []ResponseItemDevicesGetNetworkSmDeviceSoftwares // Array of ResponseDevicesGetNetworkSmDeviceSoftwares
type ResponseItemDevicesGetNetworkSmDeviceSoftwares struct {
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
type ResponseDevicesUnenrollNetworkSmDevice interface{}
type ResponseDevicesGetNetworkSmDeviceWLANLists []ResponseItemDevicesGetNetworkSmDeviceWLANLists // Array of ResponseDevicesGetNetworkSmDeviceWlanLists
type ResponseItemDevicesGetNetworkSmDeviceWLANLists struct {
	CreatedAt string `json:"createdAt,omitempty"` // When the Meraki record for the wlanList was created.
	ID        string `json:"id,omitempty"`        // The Meraki managed Id of the wlanList record.
	Xml       string `json:"xml,omitempty"`       // An XML string containing the WLAN List for the device.
}
type ResponseDevicesGetOrganizationInventoryDevice struct {
	ClaimedAt             string   `json:"claimedAt,omitempty"`             // Claimed time of the device
	LicenseExpirationDate string   `json:"licenseExpirationDate,omitempty"` // License expiration date of the device
	Mac                   string   `json:"mac,omitempty"`                   // MAC address of the device
	Model                 string   `json:"model,omitempty"`                 // Model type of the device
	Name                  string   `json:"name,omitempty"`                  // Name of the device
	NetworkID             string   `json:"networkId,omitempty"`             // Network Id of the device
	OrderNumber           string   `json:"orderNumber,omitempty"`           // Order number of the device
	ProductType           string   `json:"productType,omitempty"`           // Product type of the device
	Serial                string   `json:"serial,omitempty"`                // Serial number of the device
	Tags                  []string `json:"tags,omitempty"`                  // Device tags
}
type ResponseDevicesCloneOrganizationSwitchDevices interface{}
type RequestDevicesUpdateDevice struct {
	Address         string   `json:"address,omitempty"`         // The address of a device
	FloorPlanID     string   `json:"floorPlanId,omitempty"`     // The floor plan to associate to this device. null disassociates the device from the floorplan.
	Lat             *float64 `json:"lat,omitempty"`             // The latitude of a device
	Lng             *float64 `json:"lng,omitempty"`             // The longitude of a device
	Mac             string   `json:"mac,omitempty"`             // Mac.
	MoveMapMarker   *bool    `json:"moveMapMarker,omitempty"`   // Whether or not to set the latitude and longitude of a device based on the new address. Only applies when lat and lng are not specified.
	Name            string   `json:"name,omitempty"`            // The name of a device
	Notes           string   `json:"notes,omitempty"`           // The notes for the device. String. Limited to 255 characters.
	SwitchProfileID string   `json:"switchProfileId,omitempty"` // The ID of a switch profile to bind to the device (for available switch profiles, see the 'Switch Profiles' endpoint). Use null to unbind the switch device from the current profile. For a device to be bindable to a switch profile, it must (1) be a switch, and (2) belong to a network that is bound to a configuration template.
	Tags            []string `json:"tags,omitempty"`            // The list of tags of a device
}
type RequestDevicesBlinkDeviceLeds struct {
	Duration *int `json:"duration,omitempty"` // The duration in seconds. Must be between 5 and 120. Default is 20 seconds
	Duty     *int `json:"duty,omitempty"`     // The duty cycle as the percent active. Must be between 10 and 90. Default is 50.
	Period   *int `json:"period,omitempty"`   // The period in milliseconds. Must be between 100 and 1000. Default is 160 milliseconds
}
type RequestDevicesUpdateDeviceCellularSims struct {
	SimFailover *RequestDevicesUpdateDeviceCellularSimsSimFailover `json:"simFailover,omitempty"` // SIM Failover settings.
	Sims        *[]RequestDevicesUpdateDeviceCellularSimsSims      `json:"sims,omitempty"`        // List of SIMs. If a SIM was previously configured and not specified in this request, it will remain unchanged.
}
type RequestDevicesUpdateDeviceCellularSimsSimFailover struct {
	Enabled *bool `json:"enabled,omitempty"` // Failover to secondary SIM (optional)
}
type RequestDevicesUpdateDeviceCellularSimsSims struct {
	Apns      *[]RequestDevicesUpdateDeviceCellularSimsSimsApns `json:"apns,omitempty"`      // APN configurations. If empty, the default APN will be used.
	IsPrimary *bool                                             `json:"isPrimary,omitempty"` // If true, this SIM is used for boot. Must be true on single-sim devices.
	Slot      string                                            `json:"slot,omitempty"`      // SIM slot being configured. Must be 'sim1' on single-sim devices.
}
type RequestDevicesUpdateDeviceCellularSimsSimsApns struct {
	AllowedIPTypes []string                                                      `json:"allowedIpTypes,omitempty"` // IP versions to support (permitted values include 'ipv4', 'ipv6').
	Authentication *RequestDevicesUpdateDeviceCellularSimsSimsApnsAuthentication `json:"authentication,omitempty"` // APN authentication configurations.
	Name           string                                                        `json:"name,omitempty"`           // APN name.
}
type RequestDevicesUpdateDeviceCellularSimsSimsApnsAuthentication struct {
	Password string `json:"password,omitempty"` // APN password, if type is set (if APN password is not supplied, the password is left unchanged).
	Type     string `json:"type,omitempty"`     // APN auth type.
	Username string `json:"username,omitempty"` // APN username, if type is set.
}
type RequestDevicesCreateDeviceLiveToolsPing struct {
	Count  *int   `json:"count,omitempty"`  // Count parameter to pass to ping. [1..5], default 5
	Target string `json:"target,omitempty"` // FQDN, IPv4 or IPv6 address
}
type RequestDevicesCreateDeviceLiveToolsPingDevice struct {
	Count *int `json:"count,omitempty"` // Count parameter to pass to ping. [1..5], default 5
}
type RequestDevicesUpdateDeviceManagementInterface struct {
	Wan1 *RequestDevicesUpdateDeviceManagementInterfaceWan1 `json:"wan1,omitempty"` // WAN 1 settings
	Wan2 *RequestDevicesUpdateDeviceManagementInterfaceWan2 `json:"wan2,omitempty"` // WAN 2 settings (only for MX devices)
}
type RequestDevicesUpdateDeviceManagementInterfaceWan1 struct {
	StaticDNS        []string `json:"staticDns,omitempty"`        // Up to two DNS IPs.
	StaticGatewayIP  string   `json:"staticGatewayIp,omitempty"`  // The IP of the gateway on the WAN.
	StaticIP         string   `json:"staticIp,omitempty"`         // The IP the device should use on the WAN.
	StaticSubnetMask string   `json:"staticSubnetMask,omitempty"` // The subnet mask for the WAN.
	UsingStaticIP    *bool    `json:"usingStaticIp,omitempty"`    // Configure the interface to have static IP settings or use DHCP.
	VLAN             *int     `json:"vlan,omitempty"`             // The VLAN that management traffic should be tagged with. Applies whether usingStaticIp is true or false.
	WanEnabled       string   `json:"wanEnabled,omitempty"`       // Enable or disable the interface (only for MX devices). Valid values are 'enabled', 'disabled', and 'not configured'.
}
type RequestDevicesUpdateDeviceManagementInterfaceWan2 struct {
	StaticDNS        []string `json:"staticDns,omitempty"`        // Up to two DNS IPs.
	StaticGatewayIP  string   `json:"staticGatewayIp,omitempty"`  // The IP of the gateway on the WAN.
	StaticIP         string   `json:"staticIp,omitempty"`         // The IP the device should use on the WAN.
	StaticSubnetMask string   `json:"staticSubnetMask,omitempty"` // The subnet mask for the WAN.
	UsingStaticIP    *bool    `json:"usingStaticIp,omitempty"`    // Configure the interface to have static IP settings or use DHCP.
	VLAN             *int     `json:"vlan,omitempty"`             // The VLAN that management traffic should be tagged with. Applies whether usingStaticIp is true or false.
	WanEnabled       string   `json:"wanEnabled,omitempty"`       // Enable or disable the interface (only for MX devices). Valid values are 'enabled', 'disabled', and 'not configured'.
}
type RequestDevicesClaimNetworkDevices struct {
	Serials []string `json:"serials,omitempty"` // A list of serials of devices to claim
}
type RequestDevicesVmxNetworkDevicesClaim struct {
	Size string `json:"size,omitempty"` // The size of the vMX you claim. It can be one of: small, medium, large, 100
}
type RequestDevicesRemoveNetworkDevices struct {
	Serial string `json:"serial,omitempty"` // The serial of a device
}
type RequestDevicesCheckinNetworkSmDevices struct {
	IDs      []string `json:"ids,omitempty"`      // The ids of the devices to be checked-in.
	Scope    []string `json:"scope,omitempty"`    // The scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags of the devices to be checked-in.
	Serials  []string `json:"serials,omitempty"`  // The serials of the devices to be checked-in.
	WifiMacs []string `json:"wifiMacs,omitempty"` // The wifiMacs of the devices to be checked-in.
}
type RequestDevicesUpdateNetworkSmDevicesFields struct {
	DeviceFields *RequestDevicesUpdateNetworkSmDevicesFieldsDeviceFields `json:"deviceFields,omitempty"` // The new fields of the device. Each field of this object is optional.
	ID           string                                                  `json:"id,omitempty"`           // The id of the device to be modified.
	Serial       string                                                  `json:"serial,omitempty"`       // The serial of the device to be modified.
	WifiMac      string                                                  `json:"wifiMac,omitempty"`      // The wifiMac of the device to be modified.
}
type RequestDevicesUpdateNetworkSmDevicesFieldsDeviceFields struct {
	Name  string `json:"name,omitempty"`  // New name for the device
	Notes string `json:"notes,omitempty"` // New notes for the device
}
type RequestDevicesLockNetworkSmDevices struct {
	IDs      []string `json:"ids,omitempty"`      // The ids of the devices to be locked.
	Pin      *int     `json:"pin,omitempty"`      // The pin number for locking macOS devices (a six digit number). Required only for macOS devices.
	Scope    []string `json:"scope,omitempty"`    // The scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags of the devices to be wiped.
	Serials  []string `json:"serials,omitempty"`  // The serials of the devices to be locked.
	WifiMacs []string `json:"wifiMacs,omitempty"` // The wifiMacs of the devices to be locked.
}
type RequestDevicesModifyNetworkSmDevicesTags struct {
	IDs          []string `json:"ids,omitempty"`          // The ids of the devices to be modified.
	Scope        []string `json:"scope,omitempty"`        // The scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags of the devices to be modified.
	Serials      []string `json:"serials,omitempty"`      // The serials of the devices to be modified.
	Tags         []string `json:"tags,omitempty"`         // The tags to be added, deleted, or updated.
	UpdateAction string   `json:"updateAction,omitempty"` // One of add, delete, or update. Only devices that have been modified will be returned.
	WifiMacs     []string `json:"wifiMacs,omitempty"`     // The wifiMacs of the devices to be modified.
}
type RequestDevicesMoveNetworkSmDevices struct {
	IDs        []string `json:"ids,omitempty"`        // The ids of the devices to be moved.
	NewNetwork string   `json:"newNetwork,omitempty"` // The new network to which the devices will be moved.
	Scope      []string `json:"scope,omitempty"`      // The scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags of the devices to be moved.
	Serials    []string `json:"serials,omitempty"`    // The serials of the devices to be moved.
	WifiMacs   []string `json:"wifiMacs,omitempty"`   // The wifiMacs of the devices to be moved.
}
type RequestDevicesWipeNetworkSmDevices struct {
	ID      string `json:"id,omitempty"`      // The id of the device to be wiped.
	Pin     *int   `json:"pin,omitempty"`     // The pin number (a six digit value) for wiping a macOS device. Required only for macOS devices.
	Serial  string `json:"serial,omitempty"`  // The serial of the device to be wiped.
	WifiMac string `json:"wifiMac,omitempty"` // The wifiMac of the device to be wiped.
}
type RequestOrganizationsCloneOrganizationSwitchDevices struct {
	SourceSerial  string   `json:"sourceSerial,omitempty"`  // Serial number of the source switch (must be on a network not bound to a template)
	TargetSerials []string `json:"targetSerials,omitempty"` // Array of serial numbers of one or more target switches (must be on a network not bound to a template)
}

//GetDevice Return a single device
/* Return a single device

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device
*/
func (s *DevicesService) GetDevice(serial string) (*ResponseDevicesGetDevice, *resty.Response, error) {
	path := "/api/v1/devices/{serial}"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDevice")
	}

	result := response.Result().(*ResponseDevicesGetDevice)
	return result, response, err

}

//GetDeviceCellularSims Return the SIM and APN configurations for a cellular device.
/* Return the SIM and APN configurations for a cellular device.

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-cellular-sims
*/
func (s *DevicesService) GetDeviceCellularSims(serial string) (*ResponseDevicesGetDeviceCellularSims, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/cellular/sims"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceCellularSims{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceCellularSims")
	}

	result := response.Result().(*ResponseDevicesGetDeviceCellularSims)
	return result, response, err

}

//GetDeviceClients List the clients of a device, up to a maximum of a month ago
/* List the clients of a device, up to a maximum of a month ago. The usage of each client is returned in kilobytes. If the device is a switch, the switchport is returned; otherwise the switchport field is null.

@param serial serial path parameter.
@param getDeviceClientsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-clients
*/
func (s *DevicesService) GetDeviceClients(serial string, getDeviceClientsQueryParams *GetDeviceClientsQueryParams) (*ResponseDevicesGetDeviceClients, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/clients"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	queryString, _ := query.Values(getDeviceClientsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetDeviceClients{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceClients")
	}

	result := response.Result().(*ResponseDevicesGetDeviceClients)
	return result, response, err

}

//GetDeviceLiveToolsPing Return a ping job
/* Return a ping job. Latency unit in response is in milliseconds. Size is in bytes.

@param serial serial path parameter.
@param id id path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-live-tools-ping
*/
func (s *DevicesService) GetDeviceLiveToolsPing(serial string, id string) (*ResponseDevicesGetDeviceLiveToolsPing, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/liveTools/ping/{id}"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceLiveToolsPing{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceLiveToolsPing")
	}

	result := response.Result().(*ResponseDevicesGetDeviceLiveToolsPing)
	return result, response, err

}

//GetDeviceLiveToolsPingDevice Return a ping device job
/* Return a ping device job. Latency unit in response is in milliseconds. Size is in bytes.

@param serial serial path parameter.
@param id id path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-live-tools-ping-device
*/
func (s *DevicesService) GetDeviceLiveToolsPingDevice(serial string, id string) (*ResponseDevicesGetDeviceLiveToolsPingDevice, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/liveTools/pingDevice/{id}"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceLiveToolsPingDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceLiveToolsPingDevice")
	}

	result := response.Result().(*ResponseDevicesGetDeviceLiveToolsPingDevice)
	return result, response, err

}

//GetDeviceLldpCdp List LLDP and CDP information for a device
/* List LLDP and CDP information for a device

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-lldp-cdp
*/
func (s *DevicesService) GetDeviceLldpCdp(serial string) (*ResponseDevicesGetDeviceLldpCdp, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/lldpCdp"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceLldpCdp{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceLldpCdp")
	}

	result := response.Result().(*ResponseDevicesGetDeviceLldpCdp)
	return result, response, err

}

//GetDeviceLossAndLatencyHistory Get the uplink loss percentage and latency in milliseconds, and goodput in kilobits per second for a wired network device.
/* Get the uplink loss percentage and latency in milliseconds, and goodput in kilobits per second for a wired network device.

@param serial serial path parameter.
@param getDeviceLossAndLatencyHistoryQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-loss-and-latency-history
*/
func (s *DevicesService) GetDeviceLossAndLatencyHistory(serial string, getDeviceLossAndLatencyHistoryQueryParams *GetDeviceLossAndLatencyHistoryQueryParams) (*ResponseDevicesGetDeviceLossAndLatencyHistory, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/lossAndLatencyHistory"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	queryString, _ := query.Values(getDeviceLossAndLatencyHistoryQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetDeviceLossAndLatencyHistory{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceLossAndLatencyHistory")
	}

	result := response.Result().(*ResponseDevicesGetDeviceLossAndLatencyHistory)
	return result, response, err

}

//GetDeviceManagementInterface Return the management interface settings for a device
/* Return the management interface settings for a device

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-management-interface
*/
func (s *DevicesService) GetDeviceManagementInterface(serial string) (*ResponseDevicesGetDeviceManagementInterface, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/managementInterface"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceManagementInterface{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceManagementInterface")
	}

	result := response.Result().(*ResponseDevicesGetDeviceManagementInterface)
	return result, response, err

}

//GetNetworkDevices List the devices in a network
/* List the devices in a network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-devices
*/
func (s *DevicesService) GetNetworkDevices(networkID string) (*ResponseDevicesGetNetworkDevices, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/devices"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetNetworkDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkDevices")
	}

	result := response.Result().(*ResponseDevicesGetNetworkDevices)
	return result, response, err

}

//GetNetworkSmDeviceCellularUsageHistory Return the client's daily cellular data usage history
/* Return the client's daily cellular data usage history. Usage data is in kilobytes.

@param networkID networkId path parameter. Network ID
@param deviceID deviceId path parameter. Device ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sm-device-cellular-usage-history
*/
func (s *DevicesService) GetNetworkSmDeviceCellularUsageHistory(networkID string, deviceID string) (*ResponseDevicesGetNetworkSmDeviceCellularUsageHistory, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/{deviceId}/cellularUsageHistory"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetNetworkSmDeviceCellularUsageHistory{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSmDeviceCellularUsageHistory")
	}

	result := response.Result().(*ResponseDevicesGetNetworkSmDeviceCellularUsageHistory)
	return result, response, err

}

//GetNetworkSmDeviceCerts List the certs on a device
/* List the certs on a device

@param networkID networkId path parameter. Network ID
@param deviceID deviceId path parameter. Device ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sm-device-certs
*/
func (s *DevicesService) GetNetworkSmDeviceCerts(networkID string, deviceID string) (*ResponseDevicesGetNetworkSmDeviceCerts, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/{deviceId}/certs"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetNetworkSmDeviceCerts{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSmDeviceCerts")
	}

	result := response.Result().(*ResponseDevicesGetNetworkSmDeviceCerts)
	return result, response, err

}

//GetNetworkSmDeviceDeviceProfiles Get the installed profiles associated with a device
/* Get the installed profiles associated with a device

@param networkID networkId path parameter. Network ID
@param deviceID deviceId path parameter. Device ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sm-device-device-profiles
*/
func (s *DevicesService) GetNetworkSmDeviceDeviceProfiles(networkID string, deviceID string) (*ResponseDevicesGetNetworkSmDeviceDeviceProfiles, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/{deviceId}/deviceProfiles"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetNetworkSmDeviceDeviceProfiles{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSmDeviceDeviceProfiles")
	}

	result := response.Result().(*ResponseDevicesGetNetworkSmDeviceDeviceProfiles)
	return result, response, err

}

//GetNetworkSmDeviceNetworkAdapters List the network adapters of a device
/* List the network adapters of a device

@param networkID networkId path parameter. Network ID
@param deviceID deviceId path parameter. Device ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sm-device-network-adapters
*/
func (s *DevicesService) GetNetworkSmDeviceNetworkAdapters(networkID string, deviceID string) (*ResponseDevicesGetNetworkSmDeviceNetworkAdapters, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/{deviceId}/networkAdapters"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetNetworkSmDeviceNetworkAdapters{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSmDeviceNetworkAdapters")
	}

	result := response.Result().(*ResponseDevicesGetNetworkSmDeviceNetworkAdapters)
	return result, response, err

}

//GetNetworkSmDeviceRestrictions List the restrictions on a device
/* List the restrictions on a device

@param networkID networkId path parameter. Network ID
@param deviceID deviceId path parameter. Device ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sm-device-restrictions
*/
func (s *DevicesService) GetNetworkSmDeviceRestrictions(networkID string, deviceID string) (*ResponseDevicesGetNetworkSmDeviceRestrictions, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/{deviceId}/restrictions"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetNetworkSmDeviceRestrictions{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSmDeviceRestrictions")
	}

	result := response.Result().(*ResponseDevicesGetNetworkSmDeviceRestrictions)
	return result, response, err

}

//GetNetworkSmDeviceSecurityCenters List the security centers on a device
/* List the security centers on a device

@param networkID networkId path parameter. Network ID
@param deviceID deviceId path parameter. Device ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sm-device-security-centers
*/
func (s *DevicesService) GetNetworkSmDeviceSecurityCenters(networkID string, deviceID string) (*ResponseDevicesGetNetworkSmDeviceSecurityCenters, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/{deviceId}/securityCenters"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetNetworkSmDeviceSecurityCenters{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSmDeviceSecurityCenters")
	}

	result := response.Result().(*ResponseDevicesGetNetworkSmDeviceSecurityCenters)
	return result, response, err

}

//GetNetworkSmDeviceSoftwares Get a list of softwares associated with a device
/* Get a list of softwares associated with a device

@param networkID networkId path parameter. Network ID
@param deviceID deviceId path parameter. Device ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sm-device-softwares
*/
func (s *DevicesService) GetNetworkSmDeviceSoftwares(networkID string, deviceID string) (*ResponseDevicesGetNetworkSmDeviceSoftwares, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/{deviceId}/softwares"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetNetworkSmDeviceSoftwares{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSmDeviceSoftwares")
	}

	result := response.Result().(*ResponseDevicesGetNetworkSmDeviceSoftwares)
	return result, response, err

}

//GetNetworkSmDeviceWLANLists List the saved SSID names on a device
/* List the saved SSID names on a device

@param networkID networkId path parameter. Network ID
@param deviceID deviceId path parameter. Device ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sm-device-wlan-lists
*/
func (s *DevicesService) GetNetworkSmDeviceWLANLists(networkID string, deviceID string) (*ResponseDevicesGetNetworkSmDeviceWLANLists, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/{deviceId}/wlanLists"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetNetworkSmDeviceWLANLists{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSmDeviceWlanLists")
	}

	result := response.Result().(*ResponseDevicesGetNetworkSmDeviceWLANLists)
	return result, response, err

}

//GetOrganizationInventoryDevice Return a single device from the inventory of an organization
/* Return a single device from the inventory of an organization

@param organizationID organizationId path parameter. Organization ID
@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-inventory-device
*/
func (s *DevicesService) GetOrganizationInventoryDevice(organizationID string, serial string) (*ResponseDevicesGetOrganizationInventoryDevice, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/inventory/devices/{serial}"
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetOrganizationInventoryDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationInventoryDevice")
	}

	result := response.Result().(*ResponseDevicesGetOrganizationInventoryDevice)
	return result, response, err

}

//BlinkDeviceLeds Blink the LEDs on a device
/* Blink the LEDs on a device

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!blink-device-leds
*/

func (s *DevicesService) BlinkDeviceLeds(serial string, requestDevicesBlinkDeviceLeds *RequestDevicesBlinkDeviceLeds) (*ResponseDevicesBlinkDeviceLeds, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/blinkLeds"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesBlinkDeviceLeds).
		SetResult(&ResponseDevicesBlinkDeviceLeds{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation BlinkDeviceLeds")
	}

	result := response.Result().(*ResponseDevicesBlinkDeviceLeds)
	return result, response, err

}

//CreateDeviceLiveToolsPing Enqueue a job to ping a target host from the device
/* Enqueue a job to ping a target host from the device

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-device-live-tools-ping
*/

func (s *DevicesService) CreateDeviceLiveToolsPing(serial string, requestDevicesCreateDeviceLiveToolsPing *RequestDevicesCreateDeviceLiveToolsPing) (*ResponseDevicesCreateDeviceLiveToolsPing, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/liveTools/ping"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesCreateDeviceLiveToolsPing).
		SetResult(&ResponseDevicesCreateDeviceLiveToolsPing{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateDeviceLiveToolsPing")
	}

	result := response.Result().(*ResponseDevicesCreateDeviceLiveToolsPing)
	return result, response, err

}

//CreateDeviceLiveToolsPingDevice Enqueue a job to check connectivity status to the device
/* Enqueue a job to check connectivity status to the device

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-device-live-tools-ping-device
*/

func (s *DevicesService) CreateDeviceLiveToolsPingDevice(serial string, requestDevicesCreateDeviceLiveToolsPingDevice *RequestDevicesCreateDeviceLiveToolsPingDevice) (*ResponseDevicesCreateDeviceLiveToolsPingDevice, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/liveTools/pingDevice"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesCreateDeviceLiveToolsPingDevice).
		SetResult(&ResponseDevicesCreateDeviceLiveToolsPingDevice{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateDeviceLiveToolsPingDevice")
	}

	result := response.Result().(*ResponseDevicesCreateDeviceLiveToolsPingDevice)
	return result, response, err

}

//RebootDevice Reboot a device
/* Reboot a device

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!reboot-device
*/

func (s *DevicesService) RebootDevice(serial string) (*resty.Response, error) {
	path := "/api/v1/devices/{serial}/reboot"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").

		// SetResult(&ResponseDevicesRebootDevice{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation RebootDevice")
	}

	return response, err

}

//ClaimNetworkDevices Claim devices into a network. (Note: for recently claimed devices, it may take a few minutes for API requsts against that device to succeed)
/* Claim devices into a network. (Note: for recently claimed devices, it may take a few minutes for API requsts against that device to succeed)

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!claim-network-devices
*/

func (s *DevicesService) ClaimNetworkDevices(networkID string, requestDevicesClaimNetworkDevices *RequestDevicesClaimNetworkDevices) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/devices/claim"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesClaimNetworkDevices).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation ClaimNetworkDevices")
	}

	return response, err

}

//VmxNetworkDevicesClaim Claim a vMX into a network
/* Claim a vMX into a network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!vmx-network-devices-claim
*/

func (s *DevicesService) VmxNetworkDevicesClaim(networkID string, requestDevicesVmxNetworkDevicesClaim *RequestDevicesVmxNetworkDevicesClaim) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/devices/claim/vmx"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesVmxNetworkDevicesClaim).
		// SetResult(&ResponseDevicesVmxNetworkDevicesClaim{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation VmxNetworkDevicesClaim")
	}

	return response, err

}

//RemoveNetworkDevices Remove a single device
/* Remove a single device

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!remove-network-devices
*/

func (s *DevicesService) RemoveNetworkDevices(networkID string, requestDevicesRemoveNetworkDevices *RequestDevicesRemoveNetworkDevices) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/devices/remove"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesRemoveNetworkDevices).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation RemoveNetworkDevices")
	}

	return response, err

}

//CheckinNetworkSmDevices Force check-in a set of devices
/* Force check-in a set of devices

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!checkin-network-sm-devices
*/

func (s *DevicesService) CheckinNetworkSmDevices(networkID string, requestDevicesCheckinNetworkSmDevices *RequestDevicesCheckinNetworkSmDevices) (*ResponseDevicesCheckinNetworkSmDevices, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/checkin"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesCheckinNetworkSmDevices).
		SetResult(&ResponseDevicesCheckinNetworkSmDevices{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CheckinNetworkSmDevices")
	}

	result := response.Result().(*ResponseDevicesCheckinNetworkSmDevices)
	return result, response, err

}

//LockNetworkSmDevices Lock a set of devices
/* Lock a set of devices

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!lock-network-sm-devices
*/

func (s *DevicesService) LockNetworkSmDevices(networkID string, requestDevicesLockNetworkSmDevices *RequestDevicesLockNetworkSmDevices) (*ResponseDevicesLockNetworkSmDevices, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/lock"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesLockNetworkSmDevices).
		SetResult(&ResponseDevicesLockNetworkSmDevices{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation LockNetworkSmDevices")
	}

	result := response.Result().(*ResponseDevicesLockNetworkSmDevices)
	return result, response, err

}

//ModifyNetworkSmDevicesTags Add, delete, or update the tags of a set of devices
/* Add, delete, or update the tags of a set of devices

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!modify-network-sm-devices-tags
*/

func (s *DevicesService) ModifyNetworkSmDevicesTags(networkID string, requestDevicesModifyNetworkSmDevicesTags *RequestDevicesModifyNetworkSmDevicesTags) (*ResponseDevicesModifyNetworkSmDevicesTags, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/modifyTags"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesModifyNetworkSmDevicesTags).
		SetResult(&ResponseDevicesModifyNetworkSmDevicesTags{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ModifyNetworkSmDevicesTags")
	}

	result := response.Result().(*ResponseDevicesModifyNetworkSmDevicesTags)
	return result, response, err

}

//MoveNetworkSmDevices Move a set of devices to a new network
/* Move a set of devices to a new network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!move-network-sm-devices
*/

func (s *DevicesService) MoveNetworkSmDevices(networkID string, requestDevicesMoveNetworkSmDevices *RequestDevicesMoveNetworkSmDevices) (*ResponseDevicesMoveNetworkSmDevices, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/move"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesMoveNetworkSmDevices).
		SetResult(&ResponseDevicesMoveNetworkSmDevices{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation MoveNetworkSmDevices")
	}

	result := response.Result().(*ResponseDevicesMoveNetworkSmDevices)
	return result, response, err

}

//WipeNetworkSmDevices Wipe a device
/* Wipe a device

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!wipe-network-sm-devices
*/

func (s *DevicesService) WipeNetworkSmDevices(networkID string, requestDevicesWipeNetworkSmDevices *RequestDevicesWipeNetworkSmDevices) (*ResponseDevicesWipeNetworkSmDevices, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/wipe"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesWipeNetworkSmDevices).
		SetResult(&ResponseDevicesWipeNetworkSmDevices{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation WipeNetworkSmDevices")
	}

	result := response.Result().(*ResponseDevicesWipeNetworkSmDevices)
	return result, response, err

}

//RefreshNetworkSmDeviceDetails Refresh the details of a device
/* Refresh the details of a device

@param networkID networkId path parameter. Network ID
@param deviceID deviceId path parameter. Device ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!refresh-network-sm-device-details
*/

func (s *DevicesService) RefreshNetworkSmDeviceDetails(networkID string, deviceID string) (*resty.Response, error) {
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

func (s *DevicesService) UnenrollNetworkSmDevice(networkID string, deviceID string) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/{deviceId}/unenroll"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").

		// SetResult(&ResponseDevicesUnenrollNetworkSmDevice{}).
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

//CloneOrganizationSwitchDevices Clone port-level and some switch-level configuration settings from a source switch to one or more target switches
/* Clone port-level and some switch-level configuration settings from a source switch to one or more target switches. Cloned settings include: Aggregation Groups, Power Settings, Multicast Settings, MTU Configuration, STP Bridge priority, Port Mirroring

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!clone-organization-switch-devices
*/

func (s *DevicesService) CloneOrganizationSwitchDevices(organizationID string) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/switch/devices/clone"
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation CloneOrganizationSwitchDevices")
	}

	return response, err

}

//UpdateDevice Update the attributes of a device
/* Update the attributes of a device

@param serial serial path parameter.
*/
func (s *DevicesService) UpdateDevice(serial string, requestDevicesUpdateDevice *RequestDevicesUpdateDevice) (*resty.Response, error) {
	path := "/api/v1/devices/{serial}"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesUpdateDevice).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateDevice")
	}

	return response, err

}

//UpdateDeviceCellularSims Updates the SIM and APN configurations for a cellular device.
/* Updates the SIM and APN configurations for a cellular device.

@param serial serial path parameter.
*/
func (s *DevicesService) UpdateDeviceCellularSims(serial string, requestDevicesUpdateDeviceCellularSims *RequestDevicesUpdateDeviceCellularSims) (*resty.Response, error) {
	path := "/api/v1/devices/{serial}/cellular/sims"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesUpdateDeviceCellularSims).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateDeviceCellularSims")
	}

	return response, err

}

//UpdateDeviceManagementInterface Update the management interface settings for a device
/* Update the management interface settings for a device

@param serial serial path parameter.
*/
func (s *DevicesService) UpdateDeviceManagementInterface(serial string, requestDevicesUpdateDeviceManagementInterface *RequestDevicesUpdateDeviceManagementInterface) (*resty.Response, error) {
	path := "/api/v1/devices/{serial}/managementInterface"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesUpdateDeviceManagementInterface).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateDeviceManagementInterface")
	}

	return response, err

}

//UpdateNetworkSmDevicesFields Modify the fields of a device
/* Modify the fields of a device

@param networkID networkId path parameter. Network ID
*/
func (s *DevicesService) UpdateNetworkSmDevicesFields(networkID string, requestDevicesUpdateNetworkSmDevicesFields *RequestDevicesUpdateNetworkSmDevicesFields) (*ResponseDevicesUpdateNetworkSmDevicesFields, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/fields"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesUpdateNetworkSmDevicesFields).
		SetResult(&ResponseDevicesUpdateNetworkSmDevicesFields{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkSmDevicesFields")
	}

	result := response.Result().(*ResponseDevicesUpdateNetworkSmDevicesFields)
	return result, response, err

}
