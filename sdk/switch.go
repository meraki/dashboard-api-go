package meraki

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type SwitchService service

type GetDeviceSwitchPortsStatusesQueryParams struct {
	T0       string  `url:"t0,omitempty"`       //The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	Timespan float64 `url:"timespan,omitempty"` //The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
}
type GetDeviceSwitchPortsStatusesPacketsQueryParams struct {
	T0       string  `url:"t0,omitempty"`       //The beginning of the timespan for the data. The maximum lookback period is 1 day from today.
	Timespan float64 `url:"timespan,omitempty"` //The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 1 day. The default is 1 day.
}
type GetNetworkSwitchDhcpV4ServersSeenQueryParams struct {
	T0            string  `url:"t0,omitempty"`            //The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	Timespan      float64 `url:"timespan,omitempty"`      //The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
	PerPage       int     `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter string  `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string  `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
}
type GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersQueryParams struct {
	PerPage       int    `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter string `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
}
type GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceQueryParams struct {
	PerPage       int    `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter string `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
}
type GetOrganizationSwitchPortsBySwitchQueryParams struct {
	PerPage                   int      `url:"perPage,omitempty"`                   //The number of entries per page returned. Acceptable range is 3 - 50. Default is 50.
	StartingAfter             string   `url:"startingAfter,omitempty"`             //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore              string   `url:"endingBefore,omitempty"`              //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	NetworkIDs                []string `url:"networkIds[],omitempty"`              //Optional parameter to filter switchports by network.
	PortProfileIDs            []string `url:"portProfileIds[],omitempty"`          //Optional parameter to filter switchports belonging to the specified port profiles.
	Name                      string   `url:"name,omitempty"`                      //Optional parameter to filter switchports belonging to switches by name. All returned switches will have a name that contains the search term or is an exact match.
	Mac                       string   `url:"mac,omitempty"`                       //Optional parameter to filter switchports belonging to switches by MAC address. All returned switches will have a MAC address that contains the search term or is an exact match.
	Macs                      []string `url:"macs[],omitempty"`                    //Optional parameter to filter switchports by one or more MAC addresses belonging to devices. All switchports returned belong to MAC addresses of switches that are an exact match.
	Serial                    string   `url:"serial,omitempty"`                    //Optional parameter to filter switchports belonging to switches by serial number. All returned switches will have a serial number that contains the search term or is an exact match.
	Serials                   []string `url:"serials[],omitempty"`                 //Optional parameter to filter switchports belonging to switches with one or more serial numbers. All switchports returned belong to serial numbers of switches that are an exact match.
	ConfigurationUpdatedAfter string   `url:"configurationUpdatedAfter,omitempty"` //Optional parameter to filter results by switches where the configuration has been updated after the given timestamp.
}

type ResponseSwitchGetDeviceSwitchPorts []ResponseItemSwitchGetDeviceSwitchPorts // Array of ResponseSwitchGetDeviceSwitchPorts
type ResponseItemSwitchGetDeviceSwitchPorts struct {
	AccessPolicyNumber          *int                                           `json:"accessPolicyNumber,omitempty"`          // The number of a custom access policy to configure on the switch port. Only applicable when 'accessPolicyType' is 'Custom access policy'.
	AccessPolicyType            string                                         `json:"accessPolicyType,omitempty"`            // The type of the access policy of the switch port. Only applicable to access ports. Can be one of 'Open', 'Custom access policy', 'MAC allow list' or 'Sticky MAC allow list'.
	AdaptivePolicyGroupID       string                                         `json:"adaptivePolicyGroupId,omitempty"`       // The adaptive policy group ID that will be used to tag traffic through this switch port. This ID must pre-exist during the configuration, else needs to be created using adaptivePolicy/groups API. Cannot be applied to a port on a switch bound to profile.
	AllowedVLANs                string                                         `json:"allowedVlans,omitempty"`                // The VLANs allowed on the switch port. Only applicable to trunk ports.
	DaiTrusted                  *bool                                          `json:"daiTrusted,omitempty"`                  // If true, ARP packets for this port will be considered trusted, and Dynamic ARP Inspection will allow the traffic.
	Enabled                     *bool                                          `json:"enabled,omitempty"`                     // The status of the switch port.
	FlexibleStackingEnabled     *bool                                          `json:"flexibleStackingEnabled,omitempty"`     // For supported switches (e.g. MS420/MS425), whether or not the port has flexible stacking enabled.
	IsolationEnabled            *bool                                          `json:"isolationEnabled,omitempty"`            // The isolation status of the switch port.
	LinkNegotiation             string                                         `json:"linkNegotiation,omitempty"`             // The link speed for the switch port.
	LinkNegotiationCapabilities []string                                       `json:"linkNegotiationCapabilities,omitempty"` // Available link speeds for the switch port.
	MacAllowList                []string                                       `json:"macAllowList,omitempty"`                // Only devices with MAC addresses specified in this list will have access to this port. Up to 20 MAC addresses can be defined. Only applicable when 'accessPolicyType' is 'MAC allow list'.
	Name                        string                                         `json:"name,omitempty"`                        // The name of the switch port.
	PeerSgtCapable              *bool                                          `json:"peerSgtCapable,omitempty"`              // If true, Peer SGT is enabled for traffic through this switch port. Applicable to trunk port only, not access port. Cannot be applied to a port on a switch bound to profile.
	PoeEnabled                  *bool                                          `json:"poeEnabled,omitempty"`                  // The PoE status of the switch port.
	PortID                      string                                         `json:"portId,omitempty"`                      // The identifier of the switch port.
	PortScheduleID              string                                         `json:"portScheduleId,omitempty"`              // The ID of the port schedule. A value of null will clear the port schedule.
	Profile                     *ResponseItemSwitchGetDeviceSwitchPortsProfile `json:"profile,omitempty"`                     // Profile attributes
	RstpEnabled                 *bool                                          `json:"rstpEnabled,omitempty"`                 // The rapid spanning tree protocol status.
	StickyMacAllowList          []string                                       `json:"stickyMacAllowList,omitempty"`          // The initial list of MAC addresses for sticky Mac allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
	StickyMacAllowListLimit     *int                                           `json:"stickyMacAllowListLimit,omitempty"`     // The maximum number of MAC addresses for sticky MAC allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
	StormControlEnabled         *bool                                          `json:"stormControlEnabled,omitempty"`         // The storm control status of the switch port.
	StpGuard                    string                                         `json:"stpGuard,omitempty"`                    // The state of the STP guard ('disabled', 'root guard', 'bpdu guard' or 'loop guard').
	Tags                        []string                                       `json:"tags,omitempty"`                        // The list of tags of the switch port.
	Type                        string                                         `json:"type,omitempty"`                        // The type of the switch port ('trunk' or 'access').
	Udld                        string                                         `json:"udld,omitempty"`                        // The action to take when Unidirectional Link is detected (Alert only, Enforce). Default configuration is Alert only.
	VLAN                        *int                                           `json:"vlan,omitempty"`                        // The VLAN of the switch port. A null value will clear the value set for trunk ports.
	VoiceVLAN                   *int                                           `json:"voiceVlan,omitempty"`                   // The voice VLAN of the switch port. Only applicable to access ports.
}
type ResponseItemSwitchGetDeviceSwitchPortsProfile struct {
	Enabled *bool  `json:"enabled,omitempty"` // When enabled, override this port's configuration with a port profile.
	ID      string `json:"id,omitempty"`      // When enabled, the ID of the port profile used to override the port's configuration.
	Iname   string `json:"iname,omitempty"`   // When enabled, the IName of the profile.
}
type ResponseSwitchCycleDeviceSwitchPorts struct {
	Ports []string `json:"ports,omitempty"` // List of switch ports
}
type ResponseSwitchGetDeviceSwitchPortsStatuses []ResponseItemSwitchGetDeviceSwitchPortsStatuses // Array of ResponseSwitchGetDeviceSwitchPortsStatuses
type ResponseItemSwitchGetDeviceSwitchPortsStatuses struct {
	Cdp            *ResponseItemSwitchGetDeviceSwitchPortsStatusesCdp           `json:"cdp,omitempty"`            // The Cisco Discovery Protocol (CDP) information of the connected device.
	ClientCount    *int                                                         `json:"clientCount,omitempty"`    // The number of clients connected through this port.
	Duplex         string                                                       `json:"duplex,omitempty"`         // The current duplex of a connected port.
	Enabled        *bool                                                        `json:"enabled,omitempty"`        // Whether the port is configured to be enabled.
	Errors         []string                                                     `json:"errors,omitempty"`         // All errors present on the port.
	IsUplink       *bool                                                        `json:"isUplink,omitempty"`       // Whether the port is the switch's uplink.
	Lldp           *ResponseItemSwitchGetDeviceSwitchPortsStatusesLldp          `json:"lldp,omitempty"`           // The Link Layer Discovery Protocol (LLDP) information of the connected device.
	PortID         string                                                       `json:"portId,omitempty"`         // The string identifier of this port on the switch. This is commonly just the port number but may contain additional identifying information such as the slot and module-type if the port is located on a port module.
	PowerUsageInWh *float64                                                     `json:"powerUsageInWh,omitempty"` // How much power (in watt-hours) has been delivered by this port during the timespan.
	SecurePort     *ResponseItemSwitchGetDeviceSwitchPortsStatusesSecurePort    `json:"securePort,omitempty"`     // The Secure Port status of the port.
	Speed          string                                                       `json:"speed,omitempty"`          // The current data transfer rate which the port is operating at.
	Status         string                                                       `json:"status,omitempty"`         // The current connection status of the port.
	TrafficInKbps  *ResponseItemSwitchGetDeviceSwitchPortsStatusesTrafficInKbps `json:"trafficInKbps,omitempty"`  // A breakdown of the average speed of data that has passed through this port during the timespan.
	UsageInKb      *ResponseItemSwitchGetDeviceSwitchPortsStatusesUsageInKb     `json:"usageInKb,omitempty"`      // A breakdown of how many kilobytes have passed through this port during the timespan.
	Warnings       []string                                                     `json:"warnings,omitempty"`       // All warnings present on the port.
}
type ResponseItemSwitchGetDeviceSwitchPortsStatusesCdp struct {
	Address             string `json:"address,omitempty"`             // Contains network addresses of both receiving and sending devices.
	Capabilities        string `json:"capabilities,omitempty"`        // Identifies the device type, which indicates the functional capabilities of the device.
	DeviceID            string `json:"deviceId,omitempty"`            // Identifies the device name.
	ManagementAddress   string `json:"managementAddress,omitempty"`   // The device's management IP.
	NativeVLAN          *int   `json:"nativeVlan,omitempty"`          // Indicates, per interface, the assumed VLAN for untagged packets on the interface.
	Platform            string `json:"platform,omitempty"`            // Identifies the hardware platform of the device.
	PortID              string `json:"portId,omitempty"`              // Identifies the port from which the CDP packet was sent.
	SystemName          string `json:"systemName,omitempty"`          // The system name.
	Version             string `json:"version,omitempty"`             // Contains the device software release information.
	VtpManagementDomain string `json:"vtpManagementDomain,omitempty"` // Advertises the configured VLAN Trunking Protocl (VTP)-management-domain name of the system.
}
type ResponseItemSwitchGetDeviceSwitchPortsStatusesLldp struct {
	ChassisID          string `json:"chassisId,omitempty"`          // The device's chassis ID.
	ManagementAddress  string `json:"managementAddress,omitempty"`  // The device's management IP.
	ManagementVLAN     *int   `json:"managementVlan,omitempty"`     // The device's management VLAN.
	PortDescription    string `json:"portDescription,omitempty"`    // Description of the port from which the LLDP packet was sent.
	PortID             string `json:"portId,omitempty"`             // Identifies the port from which the LLDP packet was sent
	PortVLAN           *int   `json:"portVlan,omitempty"`           // The port's VLAN.
	SystemCapabilities string `json:"systemCapabilities,omitempty"` // Identifies the device type, which indicates the functional capabilities of the device.
	SystemDescription  string `json:"systemDescription,omitempty"`  // The device's system description.
	SystemName         string `json:"systemName,omitempty"`         // The device's system name.
}
type ResponseItemSwitchGetDeviceSwitchPortsStatusesSecurePort struct {
	Active               *bool                                                                    `json:"active,omitempty"`               // Whether Secure Port is currently active for this port.
	AuthenticationStatus string                                                                   `json:"authenticationStatus,omitempty"` // The current Secure Port status.
	ConfigOverrides      *ResponseItemSwitchGetDeviceSwitchPortsStatusesSecurePortConfigOverrides `json:"configOverrides,omitempty"`      // The configuration overrides applied to this port when Secure Port is active.
	Enabled              *bool                                                                    `json:"enabled,omitempty"`              // Whether Secure Port is turned on for this port.
}
type ResponseItemSwitchGetDeviceSwitchPortsStatusesSecurePortConfigOverrides struct {
	AllowedVLANs string `json:"allowedVlans,omitempty"` // The VLANs allowed on the . Only applicable to trunk ports.
	Type         string `json:"type,omitempty"`         // The type of the  ('trunk' or 'access').
	VLAN         *int   `json:"vlan,omitempty"`         // The VLAN of the . A null value will clear the value set for trunk ports.
	VoiceVLAN    *int   `json:"voiceVlan,omitempty"`    // The voice VLAN of the . Only applicable to access ports.
}
type ResponseItemSwitchGetDeviceSwitchPortsStatusesTrafficInKbps struct {
	Recv  *float64 `json:"recv,omitempty"`  // The average speed of the data received (in kilobits-per-second).
	Sent  *float64 `json:"sent,omitempty"`  // The average speed of the data sent (in kilobits-per-second).
	Total *float64 `json:"total,omitempty"` // The average speed of the data sent and received (in kilobits-per-second).
}
type ResponseItemSwitchGetDeviceSwitchPortsStatusesUsageInKb struct {
	Recv  *int `json:"recv,omitempty"`  // The amount of data received (in kilobytes).
	Sent  *int `json:"sent,omitempty"`  // The amount of data sent (in kilobytes).
	Total *int `json:"total,omitempty"` // The total amount of data sent and received (in kilobytes).
}
type ResponseSwitchGetDeviceSwitchPortsStatusesPackets []ResponseItemSwitchGetDeviceSwitchPortsStatusesPackets // Array of ResponseSwitchGetDeviceSwitchPortsStatusesPackets
type ResponseItemSwitchGetDeviceSwitchPortsStatusesPackets struct {
	Packets *[]ResponseItemSwitchGetDeviceSwitchPortsStatusesPacketsPackets `json:"packets,omitempty"` //
	PortID  string                                                          `json:"portId,omitempty"`  //
}
type ResponseItemSwitchGetDeviceSwitchPortsStatusesPacketsPackets struct {
	Desc       string                                                                  `json:"desc,omitempty"`       //
	RatePerSec *ResponseItemSwitchGetDeviceSwitchPortsStatusesPacketsPacketsRatePerSec `json:"ratePerSec,omitempty"` //
	Recv       *int                                                                    `json:"recv,omitempty"`       //
	Sent       *int                                                                    `json:"sent,omitempty"`       //
	Total      *int                                                                    `json:"total,omitempty"`      //
}
type ResponseItemSwitchGetDeviceSwitchPortsStatusesPacketsPacketsRatePerSec struct {
	Recv  *int `json:"recv,omitempty"`  //
	Sent  *int `json:"sent,omitempty"`  //
	Total *int `json:"total,omitempty"` //
}
type ResponseSwitchGetDeviceSwitchPort struct {
	AccessPolicyNumber          *int                                      `json:"accessPolicyNumber,omitempty"`          // The number of a custom access policy to configure on the switch port. Only applicable when 'accessPolicyType' is 'Custom access policy'.
	AccessPolicyType            string                                    `json:"accessPolicyType,omitempty"`            // The type of the access policy of the switch port. Only applicable to access ports. Can be one of 'Open', 'Custom access policy', 'MAC allow list' or 'Sticky MAC allow list'.
	AdaptivePolicyGroupID       string                                    `json:"adaptivePolicyGroupId,omitempty"`       // The adaptive policy group ID that will be used to tag traffic through this switch port. This ID must pre-exist during the configuration, else needs to be created using adaptivePolicy/groups API. Cannot be applied to a port on a switch bound to profile.
	AllowedVLANs                string                                    `json:"allowedVlans,omitempty"`                // The VLANs allowed on the switch port. Only applicable to trunk ports.
	DaiTrusted                  *bool                                     `json:"daiTrusted,omitempty"`                  // If true, ARP packets for this port will be considered trusted, and Dynamic ARP Inspection will allow the traffic.
	Enabled                     *bool                                     `json:"enabled,omitempty"`                     // The status of the switch port.
	FlexibleStackingEnabled     *bool                                     `json:"flexibleStackingEnabled,omitempty"`     // For supported switches (e.g. MS420/MS425), whether or not the port has flexible stacking enabled.
	IsolationEnabled            *bool                                     `json:"isolationEnabled,omitempty"`            // The isolation status of the switch port.
	LinkNegotiation             string                                    `json:"linkNegotiation,omitempty"`             // The link speed for the switch port.
	LinkNegotiationCapabilities []string                                  `json:"linkNegotiationCapabilities,omitempty"` // Available link speeds for the switch port.
	MacAllowList                []string                                  `json:"macAllowList,omitempty"`                // Only devices with MAC addresses specified in this list will have access to this port. Up to 20 MAC addresses can be defined. Only applicable when 'accessPolicyType' is 'MAC allow list'.
	Name                        string                                    `json:"name,omitempty"`                        // The name of the switch port.
	PeerSgtCapable              *bool                                     `json:"peerSgtCapable,omitempty"`              // If true, Peer SGT is enabled for traffic through this switch port. Applicable to trunk port only, not access port. Cannot be applied to a port on a switch bound to profile.
	PoeEnabled                  *bool                                     `json:"poeEnabled,omitempty"`                  // The PoE status of the switch port.
	PortID                      string                                    `json:"portId,omitempty"`                      // The identifier of the switch port.
	PortScheduleID              string                                    `json:"portScheduleId,omitempty"`              // The ID of the port schedule. A value of null will clear the port schedule.
	Profile                     *ResponseSwitchGetDeviceSwitchPortProfile `json:"profile,omitempty"`                     // Profile attributes
	RstpEnabled                 *bool                                     `json:"rstpEnabled,omitempty"`                 // The rapid spanning tree protocol status.
	StickyMacAllowList          []string                                  `json:"stickyMacAllowList,omitempty"`          // The initial list of MAC addresses for sticky Mac allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
	StickyMacAllowListLimit     *int                                      `json:"stickyMacAllowListLimit,omitempty"`     // The maximum number of MAC addresses for sticky MAC allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
	StormControlEnabled         *bool                                     `json:"stormControlEnabled,omitempty"`         // The storm control status of the switch port.
	StpGuard                    string                                    `json:"stpGuard,omitempty"`                    // The state of the STP guard ('disabled', 'root guard', 'bpdu guard' or 'loop guard').
	Tags                        []string                                  `json:"tags,omitempty"`                        // The list of tags of the switch port.
	Type                        string                                    `json:"type,omitempty"`                        // The type of the switch port ('trunk' or 'access').
	Udld                        string                                    `json:"udld,omitempty"`                        // The action to take when Unidirectional Link is detected (Alert only, Enforce). Default configuration is Alert only.
	VLAN                        *int                                      `json:"vlan,omitempty"`                        // The VLAN of the switch port. A null value will clear the value set for trunk ports.
	VoiceVLAN                   *int                                      `json:"voiceVlan,omitempty"`                   // The voice VLAN of the switch port. Only applicable to access ports.
}
type ResponseSwitchGetDeviceSwitchPortProfile struct {
	Enabled *bool  `json:"enabled,omitempty"` // When enabled, override this port's configuration with a port profile.
	ID      string `json:"id,omitempty"`      // When enabled, the ID of the port profile used to override the port's configuration.
	Iname   string `json:"iname,omitempty"`   // When enabled, the IName of the profile.
}
type ResponseSwitchUpdateDeviceSwitchPort struct {
	AccessPolicyNumber          *int                                         `json:"accessPolicyNumber,omitempty"`          // The number of a custom access policy to configure on the switch port. Only applicable when 'accessPolicyType' is 'Custom access policy'.
	AccessPolicyType            string                                       `json:"accessPolicyType,omitempty"`            // The type of the access policy of the switch port. Only applicable to access ports. Can be one of 'Open', 'Custom access policy', 'MAC allow list' or 'Sticky MAC allow list'.
	AdaptivePolicyGroupID       string                                       `json:"adaptivePolicyGroupId,omitempty"`       // The adaptive policy group ID that will be used to tag traffic through this switch port. This ID must pre-exist during the configuration, else needs to be created using adaptivePolicy/groups API. Cannot be applied to a port on a switch bound to profile.
	AllowedVLANs                string                                       `json:"allowedVlans,omitempty"`                // The VLANs allowed on the switch port. Only applicable to trunk ports.
	DaiTrusted                  *bool                                        `json:"daiTrusted,omitempty"`                  // If true, ARP packets for this port will be considered trusted, and Dynamic ARP Inspection will allow the traffic.
	Enabled                     *bool                                        `json:"enabled,omitempty"`                     // The status of the switch port.
	FlexibleStackingEnabled     *bool                                        `json:"flexibleStackingEnabled,omitempty"`     // For supported switches (e.g. MS420/MS425), whether or not the port has flexible stacking enabled.
	IsolationEnabled            *bool                                        `json:"isolationEnabled,omitempty"`            // The isolation status of the switch port.
	LinkNegotiation             string                                       `json:"linkNegotiation,omitempty"`             // The link speed for the switch port.
	LinkNegotiationCapabilities []string                                     `json:"linkNegotiationCapabilities,omitempty"` // Available link speeds for the switch port.
	MacAllowList                []string                                     `json:"macAllowList,omitempty"`                // Only devices with MAC addresses specified in this list will have access to this port. Up to 20 MAC addresses can be defined. Only applicable when 'accessPolicyType' is 'MAC allow list'.
	Name                        string                                       `json:"name,omitempty"`                        // The name of the switch port.
	PeerSgtCapable              *bool                                        `json:"peerSgtCapable,omitempty"`              // If true, Peer SGT is enabled for traffic through this switch port. Applicable to trunk port only, not access port. Cannot be applied to a port on a switch bound to profile.
	PoeEnabled                  *bool                                        `json:"poeEnabled,omitempty"`                  // The PoE status of the switch port.
	PortID                      string                                       `json:"portId,omitempty"`                      // The identifier of the switch port.
	PortScheduleID              string                                       `json:"portScheduleId,omitempty"`              // The ID of the port schedule. A value of null will clear the port schedule.
	Profile                     *ResponseSwitchUpdateDeviceSwitchPortProfile `json:"profile,omitempty"`                     // Profile attributes
	RstpEnabled                 *bool                                        `json:"rstpEnabled,omitempty"`                 // The rapid spanning tree protocol status.
	StickyMacAllowList          []string                                     `json:"stickyMacAllowList,omitempty"`          // The initial list of MAC addresses for sticky Mac allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
	StickyMacAllowListLimit     *int                                         `json:"stickyMacAllowListLimit,omitempty"`     // The maximum number of MAC addresses for sticky MAC allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
	StormControlEnabled         *bool                                        `json:"stormControlEnabled,omitempty"`         // The storm control status of the switch port.
	StpGuard                    string                                       `json:"stpGuard,omitempty"`                    // The state of the STP guard ('disabled', 'root guard', 'bpdu guard' or 'loop guard').
	Tags                        []string                                     `json:"tags,omitempty"`                        // The list of tags of the switch port.
	Type                        string                                       `json:"type,omitempty"`                        // The type of the switch port ('trunk' or 'access').
	Udld                        string                                       `json:"udld,omitempty"`                        // The action to take when Unidirectional Link is detected (Alert only, Enforce). Default configuration is Alert only.
	VLAN                        *int                                         `json:"vlan,omitempty"`                        // The VLAN of the switch port. A null value will clear the value set for trunk ports.
	VoiceVLAN                   *int                                         `json:"voiceVlan,omitempty"`                   // The voice VLAN of the switch port. Only applicable to access ports.
}
type ResponseSwitchUpdateDeviceSwitchPortProfile struct {
	Enabled *bool  `json:"enabled,omitempty"` // When enabled, override this port's configuration with a port profile.
	ID      string `json:"id,omitempty"`      // When enabled, the ID of the port profile used to override the port's configuration.
	Iname   string `json:"iname,omitempty"`   // When enabled, the IName of the profile.
}
type ResponseSwitchGetDeviceSwitchRoutingInterfaces []ResponseItemSwitchGetDeviceSwitchRoutingInterfaces // Array of ResponseSwitchGetDeviceSwitchRoutingInterfaces
type ResponseItemSwitchGetDeviceSwitchRoutingInterfaces struct {
	DefaultGateway   string                                                          `json:"defaultGateway,omitempty"`   // IPv4 default gateway
	InterfaceID      string                                                          `json:"interfaceId,omitempty"`      // The id
	InterfaceIP      string                                                          `json:"interfaceIp,omitempty"`      // IPv4 address
	IPv6             *ResponseItemSwitchGetDeviceSwitchRoutingInterfacesIPv6         `json:"ipv6,omitempty"`             // IPv6 addressing
	MulticastRouting string                                                          `json:"multicastRouting,omitempty"` // Multicast routing status
	Name             string                                                          `json:"name,omitempty"`             // The name
	OspfSettings     *ResponseItemSwitchGetDeviceSwitchRoutingInterfacesOspfSettings `json:"ospfSettings,omitempty"`     // IPv4 OSPF Settings
	OspfV3           *ResponseItemSwitchGetDeviceSwitchRoutingInterfacesOspfV3       `json:"ospfV3,omitempty"`           // IPv6 OSPF Settings
	Subnet           string                                                          `json:"subnet,omitempty"`           // IPv4 subnet
	VLANID           *int                                                            `json:"vlanId,omitempty"`           // VLAN id
}
type ResponseItemSwitchGetDeviceSwitchRoutingInterfacesIPv6 struct {
	Address        string `json:"address,omitempty"`        // IPv6 address
	AssignmentMode string `json:"assignmentMode,omitempty"` // Assignment mode
	Gateway        string `json:"gateway,omitempty"`        // IPv6 gateway
	Prefix         string `json:"prefix,omitempty"`         // IPv6 subnet
}
type ResponseItemSwitchGetDeviceSwitchRoutingInterfacesOspfSettings struct {
	Area             string `json:"area,omitempty"`             // Area id
	Cost             *int   `json:"cost,omitempty"`             // OSPF Cost
	IsPassiveEnabled *bool  `json:"isPassiveEnabled,omitempty"` // Disable sending Hello packets on this interface's IPv4 area
}
type ResponseItemSwitchGetDeviceSwitchRoutingInterfacesOspfV3 struct {
	Area             string `json:"area,omitempty"`             // Area id
	Cost             *int   `json:"cost,omitempty"`             // OSPF Cost
	IsPassiveEnabled *bool  `json:"isPassiveEnabled,omitempty"` // Disable sending Hello packets on this interface's IPv6 area
}
type ResponseSwitchCreateDeviceSwitchRoutingInterface struct {
	DefaultGateway   string                                                        `json:"defaultGateway,omitempty"`   // IPv4 default gateway
	InterfaceID      string                                                        `json:"interfaceId,omitempty"`      // The id
	InterfaceIP      string                                                        `json:"interfaceIp,omitempty"`      // IPv4 address
	IPv6             *ResponseSwitchCreateDeviceSwitchRoutingInterfaceIPv6         `json:"ipv6,omitempty"`             // IPv6 addressing
	MulticastRouting string                                                        `json:"multicastRouting,omitempty"` // Multicast routing status
	Name             string                                                        `json:"name,omitempty"`             // The name
	OspfSettings     *ResponseSwitchCreateDeviceSwitchRoutingInterfaceOspfSettings `json:"ospfSettings,omitempty"`     // IPv4 OSPF Settings
	OspfV3           *ResponseSwitchCreateDeviceSwitchRoutingInterfaceOspfV3       `json:"ospfV3,omitempty"`           // IPv6 OSPF Settings
	Subnet           string                                                        `json:"subnet,omitempty"`           // IPv4 subnet
	VLANID           *int                                                          `json:"vlanId,omitempty"`           // VLAN id
}
type ResponseSwitchCreateDeviceSwitchRoutingInterfaceIPv6 struct {
	Address        string `json:"address,omitempty"`        // IPv6 address
	AssignmentMode string `json:"assignmentMode,omitempty"` // Assignment mode
	Gateway        string `json:"gateway,omitempty"`        // IPv6 gateway
	Prefix         string `json:"prefix,omitempty"`         // IPv6 subnet
}
type ResponseSwitchCreateDeviceSwitchRoutingInterfaceOspfSettings struct {
	Area             string `json:"area,omitempty"`             // Area id
	Cost             *int   `json:"cost,omitempty"`             // OSPF Cost
	IsPassiveEnabled *bool  `json:"isPassiveEnabled,omitempty"` // Disable sending Hello packets on this interface's IPv4 area
}
type ResponseSwitchCreateDeviceSwitchRoutingInterfaceOspfV3 struct {
	Area             string `json:"area,omitempty"`             // Area id
	Cost             *int   `json:"cost,omitempty"`             // OSPF Cost
	IsPassiveEnabled *bool  `json:"isPassiveEnabled,omitempty"` // Disable sending Hello packets on this interface's IPv6 area
}
type ResponseSwitchGetDeviceSwitchRoutingInterface struct {
	DefaultGateway   string                                                     `json:"defaultGateway,omitempty"`   // IPv4 default gateway
	InterfaceID      string                                                     `json:"interfaceId,omitempty"`      // The id
	InterfaceIP      string                                                     `json:"interfaceIp,omitempty"`      // IPv4 address
	IPv6             *ResponseSwitchGetDeviceSwitchRoutingInterfaceIPv6         `json:"ipv6,omitempty"`             // IPv6 addressing
	MulticastRouting string                                                     `json:"multicastRouting,omitempty"` // Multicast routing status
	Name             string                                                     `json:"name,omitempty"`             // The name
	OspfSettings     *ResponseSwitchGetDeviceSwitchRoutingInterfaceOspfSettings `json:"ospfSettings,omitempty"`     // IPv4 OSPF Settings
	OspfV3           *ResponseSwitchGetDeviceSwitchRoutingInterfaceOspfV3       `json:"ospfV3,omitempty"`           // IPv6 OSPF Settings
	Subnet           string                                                     `json:"subnet,omitempty"`           // IPv4 subnet
	VLANID           *int                                                       `json:"vlanId,omitempty"`           // VLAN id
}
type ResponseSwitchGetDeviceSwitchRoutingInterfaceIPv6 struct {
	Address        string `json:"address,omitempty"`        // IPv6 address
	AssignmentMode string `json:"assignmentMode,omitempty"` // Assignment mode
	Gateway        string `json:"gateway,omitempty"`        // IPv6 gateway
	Prefix         string `json:"prefix,omitempty"`         // IPv6 subnet
}
type ResponseSwitchGetDeviceSwitchRoutingInterfaceOspfSettings struct {
	Area             string `json:"area,omitempty"`             // Area id
	Cost             *int   `json:"cost,omitempty"`             // OSPF Cost
	IsPassiveEnabled *bool  `json:"isPassiveEnabled,omitempty"` // Disable sending Hello packets on this interface's IPv4 area
}
type ResponseSwitchGetDeviceSwitchRoutingInterfaceOspfV3 struct {
	Area             string `json:"area,omitempty"`             // Area id
	Cost             *int   `json:"cost,omitempty"`             // OSPF Cost
	IsPassiveEnabled *bool  `json:"isPassiveEnabled,omitempty"` // Disable sending Hello packets on this interface's IPv6 area
}
type ResponseSwitchUpdateDeviceSwitchRoutingInterface struct {
	DefaultGateway   string                                                        `json:"defaultGateway,omitempty"`   // IPv4 default gateway
	InterfaceID      string                                                        `json:"interfaceId,omitempty"`      // The id
	InterfaceIP      string                                                        `json:"interfaceIp,omitempty"`      // IPv4 address
	IPv6             *ResponseSwitchUpdateDeviceSwitchRoutingInterfaceIPv6         `json:"ipv6,omitempty"`             // IPv6 addressing
	MulticastRouting string                                                        `json:"multicastRouting,omitempty"` // Multicast routing status
	Name             string                                                        `json:"name,omitempty"`             // The name
	OspfSettings     *ResponseSwitchUpdateDeviceSwitchRoutingInterfaceOspfSettings `json:"ospfSettings,omitempty"`     // IPv4 OSPF Settings
	OspfV3           *ResponseSwitchUpdateDeviceSwitchRoutingInterfaceOspfV3       `json:"ospfV3,omitempty"`           // IPv6 OSPF Settings
	Subnet           string                                                        `json:"subnet,omitempty"`           // IPv4 subnet
	VLANID           *int                                                          `json:"vlanId,omitempty"`           // VLAN id
}
type ResponseSwitchUpdateDeviceSwitchRoutingInterfaceIPv6 struct {
	Address        string `json:"address,omitempty"`        // IPv6 address
	AssignmentMode string `json:"assignmentMode,omitempty"` // Assignment mode
	Gateway        string `json:"gateway,omitempty"`        // IPv6 gateway
	Prefix         string `json:"prefix,omitempty"`         // IPv6 subnet
}
type ResponseSwitchUpdateDeviceSwitchRoutingInterfaceOspfSettings struct {
	Area             string `json:"area,omitempty"`             // Area id
	Cost             *int   `json:"cost,omitempty"`             // OSPF Cost
	IsPassiveEnabled *bool  `json:"isPassiveEnabled,omitempty"` // Disable sending Hello packets on this interface's IPv4 area
}
type ResponseSwitchUpdateDeviceSwitchRoutingInterfaceOspfV3 struct {
	Area             string `json:"area,omitempty"`             // Area id
	Cost             *int   `json:"cost,omitempty"`             // OSPF Cost
	IsPassiveEnabled *bool  `json:"isPassiveEnabled,omitempty"` // Disable sending Hello packets on this interface's IPv6 area
}
type ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcp struct {
	BootFileName         string                                                                 `json:"bootFileName,omitempty"`         //
	BootNextServer       string                                                                 `json:"bootNextServer,omitempty"`       //
	BootOptionsEnabled   *bool                                                                  `json:"bootOptionsEnabled,omitempty"`   //
	DhcpLeaseTime        string                                                                 `json:"dhcpLeaseTime,omitempty"`        //
	DhcpMode             string                                                                 `json:"dhcpMode,omitempty"`             //
	DhcpOptions          *[]ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpDhcpOptions        `json:"dhcpOptions,omitempty"`          //
	DNSCustomNameservers []string                                                               `json:"dnsCustomNameservers,omitempty"` //
	DNSNameserversOption string                                                                 `json:"dnsNameserversOption,omitempty"` //
	FixedIPAssignments   *[]ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpFixedIPAssignments `json:"fixedIpAssignments,omitempty"`   //
	ReservedIPRanges     *[]ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpReservedIPRanges   `json:"reservedIpRanges,omitempty"`     //
}
type ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpDhcpOptions struct {
	Code  string `json:"code,omitempty"`  //
	Type  string `json:"type,omitempty"`  //
	Value string `json:"value,omitempty"` //
}
type ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpFixedIPAssignments struct {
	IP   string `json:"ip,omitempty"`   //
	Mac  string `json:"mac,omitempty"`  //
	Name string `json:"name,omitempty"` //
}
type ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpReservedIPRanges struct {
	Comment string `json:"comment,omitempty"` //
	End     string `json:"end,omitempty"`     //
	Start   string `json:"start,omitempty"`   //
}
type ResponseSwitchUpdateDeviceSwitchRoutingInterfaceDhcp interface{}
type ResponseSwitchGetDeviceSwitchRoutingStaticRoutes []ResponseItemSwitchGetDeviceSwitchRoutingStaticRoutes // Array of ResponseSwitchGetDeviceSwitchRoutingStaticRoutes
type ResponseItemSwitchGetDeviceSwitchRoutingStaticRoutes struct {
	AdvertiseViaOspfEnabled     *bool  `json:"advertiseViaOspfEnabled,omitempty"`     //
	Name                        string `json:"name,omitempty"`                        //
	NextHopIP                   string `json:"nextHopIp,omitempty"`                   //
	PreferOverOspfRoutesEnabled *bool  `json:"preferOverOspfRoutesEnabled,omitempty"` //
	StaticRouteID               string `json:"staticRouteId,omitempty"`               //
	Subnet                      string `json:"subnet,omitempty"`                      //
}
type ResponseSwitchCreateDeviceSwitchRoutingStaticRoute interface{}
type ResponseSwitchGetDeviceSwitchRoutingStaticRoute struct {
	AdvertiseViaOspfEnabled     *bool  `json:"advertiseViaOspfEnabled,omitempty"`     // Option to advertise static routes via OSPF
	Name                        string `json:"name,omitempty"`                        // The name or description of the layer 3 static route
	NextHopIP                   string `json:"nextHopIp,omitempty"`                   //  The IP address of the router to which traffic for this destination network should be sent
	PreferOverOspfRoutesEnabled *bool  `json:"preferOverOspfRoutesEnabled,omitempty"` // Option to prefer static routes over OSPF routes
	StaticRouteID               string `json:"staticRouteId,omitempty"`               // The identifier of a layer 3 static route
	Subnet                      string `json:"subnet,omitempty"`                      // The IP address of the subnetwork specified in CIDR notation (ex. 1.2.3.0/24)
}
type ResponseSwitchUpdateDeviceSwitchRoutingStaticRoute interface{}
type ResponseSwitchGetDeviceSwitchWarmSpare struct {
	Enabled       *bool  `json:"enabled,omitempty"`       //
	PrimarySerial string `json:"primarySerial,omitempty"` //
	SpareSerial   string `json:"spareSerial,omitempty"`   //
}
type ResponseSwitchUpdateDeviceSwitchWarmSpare interface{}
type ResponseSwitchGetNetworkSwitchAccessControlLists struct {
	Rules *[]ResponseSwitchGetNetworkSwitchAccessControlListsRules `json:"rules,omitempty"` // An ordered array of the access control list rules
}
type ResponseSwitchGetNetworkSwitchAccessControlListsRules struct {
	Comment   string `json:"comment,omitempty"`   // Description of the rule (optional)
	DstCidr   string `json:"dstCidr,omitempty"`   // Destination IP address (in IP or CIDR notation)
	DstPort   string `json:"dstPort,omitempty"`   // Destination port
	IPVersion string `json:"ipVersion,omitempty"` // IP address version
	Policy    string `json:"policy,omitempty"`    // 'allow' or 'deny' traffic specified by this rule
	Protocol  string `json:"protocol,omitempty"`  // The type of protocol
	SrcCidr   string `json:"srcCidr,omitempty"`   // Source IP address (in IP or CIDR notation)
	SrcPort   string `json:"srcPort,omitempty"`   // Source port
	VLAN      string `json:"vlan,omitempty"`      // ncoming traffic VLAN
}
type ResponseSwitchUpdateNetworkSwitchAccessControlLists struct {
	Rules *[]ResponseSwitchUpdateNetworkSwitchAccessControlListsRules `json:"rules,omitempty"` // An ordered array of the access control list rules
}
type ResponseSwitchUpdateNetworkSwitchAccessControlListsRules struct {
	Comment   string `json:"comment,omitempty"`   // Description of the rule (optional)
	DstCidr   string `json:"dstCidr,omitempty"`   // Destination IP address (in IP or CIDR notation)
	DstPort   string `json:"dstPort,omitempty"`   // Destination port
	IPVersion string `json:"ipVersion,omitempty"` // IP address version
	Policy    string `json:"policy,omitempty"`    // 'allow' or 'deny' traffic specified by this rule
	Protocol  string `json:"protocol,omitempty"`  // The type of protocol
	SrcCidr   string `json:"srcCidr,omitempty"`   // Source IP address (in IP or CIDR notation)
	SrcPort   string `json:"srcPort,omitempty"`   // Source port
	VLAN      string `json:"vlan,omitempty"`      // ncoming traffic VLAN
}
type ResponseSwitchGetNetworkSwitchAccessPolicies []ResponseItemSwitchGetNetworkSwitchAccessPolicies // Array of ResponseSwitchGetNetworkSwitchAccessPolicies
type ResponseItemSwitchGetNetworkSwitchAccessPolicies struct {
	AccessPolicyType               string                                                                     `json:"accessPolicyType,omitempty"`               // Access Type of the policy. Automatically 'Hybrid authentication' when hostMode is 'Multi-Domain'.
	Dot1X                          *ResponseItemSwitchGetNetworkSwitchAccessPoliciesDot1X                     `json:"dot1x,omitempty"`                          // 802.1x Settings
	GuestPortBouncing              *bool                                                                      `json:"guestPortBouncing,omitempty"`              // If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers
	GuestVLANID                    *int                                                                       `json:"guestVlanId,omitempty"`                    // ID for the guest VLAN allow unauthorized devices access to limited network resources
	HostMode                       string                                                                     `json:"hostMode,omitempty"`                       // Choose the Host Mode for the access policy.
	IncreaseAccessSpeed            *bool                                                                      `json:"increaseAccessSpeed,omitempty"`            // Enabling this option will make switches execute 802.1X and MAC-bypass authentication simultaneously so that clients authenticate faster. Only required when accessPolicyType is 'Hybrid Authentication.
	Name                           string                                                                     `json:"name,omitempty"`                           // Name of the access policy
	Radius                         *ResponseItemSwitchGetNetworkSwitchAccessPoliciesRadius                    `json:"radius,omitempty"`                         // Object for RADIUS Settings
	RadiusAccountingEnabled        *bool                                                                      `json:"radiusAccountingEnabled,omitempty"`        // Enable to send start, interim-update and stop messages to a configured RADIUS accounting server for tracking connected clients
	RadiusAccountingServers        *[]ResponseItemSwitchGetNetworkSwitchAccessPoliciesRadiusAccountingServers `json:"radiusAccountingServers,omitempty"`        // List of RADIUS accounting servers to require connecting devices to authenticate against before granting network access
	RadiusCoaSupportEnabled        *bool                                                                      `json:"radiusCoaSupportEnabled,omitempty"`        // Change of authentication for RADIUS re-authentication and disconnection
	RadiusGroupAttribute           string                                                                     `json:"radiusGroupAttribute,omitempty"`           // Acceptable values are *""* for , or *"11"* for Group Policies ACL
	RadiusServers                  *[]ResponseItemSwitchGetNetworkSwitchAccessPoliciesRadiusServers           `json:"radiusServers,omitempty"`                  // List of RADIUS servers to require connecting devices to authenticate against before granting network access
	RadiusTestingEnabled           *bool                                                                      `json:"radiusTestingEnabled,omitempty"`           // If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers
	URLRedirectWalledGardenEnabled *bool                                                                      `json:"urlRedirectWalledGardenEnabled,omitempty"` // Enable to restrict access for clients to a response_objectific set of IP addresses or hostnames prior to authentication
	URLRedirectWalledGardenRanges  []string                                                                   `json:"urlRedirectWalledGardenRanges,omitempty"`  // IP address ranges, in CIDR notation, to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication
	VoiceVLANClients               *bool                                                                      `json:"voiceVlanClients,omitempty"`               // CDP/LLDP capable voice clients will be able to use this VLAN. Automatically true when hostMode is 'Multi-Domain'.
}
type ResponseItemSwitchGetNetworkSwitchAccessPoliciesDot1X struct {
	ControlDirection string `json:"controlDirection,omitempty"` // Supports either 'both' or 'inbound'. Set to 'inbound' to allow unauthorized egress on the switchport. Set to 'both' to control both traffic directions with authorization. Defaults to 'both'
}
type ResponseItemSwitchGetNetworkSwitchAccessPoliciesRadius struct {
	CriticalAuth             *ResponseItemSwitchGetNetworkSwitchAccessPoliciesRadiusCriticalAuth `json:"criticalAuth,omitempty"`             // Critical auth settings for when authentication is rejected by the RADIUS server
	FailedAuthVLANID         *int                                                                `json:"failedAuthVlanId,omitempty"`         // VLAN that clients will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
	ReAuthenticationInterval *int                                                                `json:"reAuthenticationInterval,omitempty"` // Re-authentication period in seconds. Will be null if hostMode is Multi-Auth
}
type ResponseItemSwitchGetNetworkSwitchAccessPoliciesRadiusCriticalAuth struct {
	DataVLANID        *int  `json:"dataVlanId,omitempty"`        // VLAN that clients who use data will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
	SuspendPortBounce *bool `json:"suspendPortBounce,omitempty"` // Enable to suspend port bounce when RADIUS servers are unreachable
	VoiceVLANID       *int  `json:"voiceVlanId,omitempty"`       // VLAN that clients who use voice will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
}
type ResponseItemSwitchGetNetworkSwitchAccessPoliciesRadiusAccountingServers struct {
	Host string `json:"host,omitempty"` // Public IP address of the RADIUS accounting server
	Port *int   `json:"port,omitempty"` // UDP port that the RADIUS Accounting server listens on for access requests
}
type ResponseItemSwitchGetNetworkSwitchAccessPoliciesRadiusServers struct {
	Host string `json:"host,omitempty"` // Public IP address of the RADIUS server
	Port *int   `json:"port,omitempty"` // UDP port that the RADIUS server listens on for access requests
}
type ResponseSwitchCreateNetworkSwitchAccessPolicy struct {
	AccessPolicyType               string                                                                  `json:"accessPolicyType,omitempty"`               // Access Type of the policy. Automatically 'Hybrid authentication' when hostMode is 'Multi-Domain'.
	Dot1X                          *ResponseSwitchCreateNetworkSwitchAccessPolicyDot1X                     `json:"dot1x,omitempty"`                          // 802.1x Settings
	GuestPortBouncing              *bool                                                                   `json:"guestPortBouncing,omitempty"`              // If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers
	GuestVLANID                    *int                                                                    `json:"guestVlanId,omitempty"`                    // ID for the guest VLAN allow unauthorized devices access to limited network resources
	HostMode                       string                                                                  `json:"hostMode,omitempty"`                       // Choose the Host Mode for the access policy.
	IncreaseAccessSpeed            *bool                                                                   `json:"increaseAccessSpeed,omitempty"`            // Enabling this option will make switches execute 802.1X and MAC-bypass authentication simultaneously so that clients authenticate faster. Only required when accessPolicyType is 'Hybrid Authentication.
	Name                           string                                                                  `json:"name,omitempty"`                           // Name of the access policy
	Radius                         *ResponseSwitchCreateNetworkSwitchAccessPolicyRadius                    `json:"radius,omitempty"`                         // Object for RADIUS Settings
	RadiusAccountingEnabled        *bool                                                                   `json:"radiusAccountingEnabled,omitempty"`        // Enable to send start, interim-update and stop messages to a configured RADIUS accounting server for tracking connected clients
	RadiusAccountingServers        *[]ResponseSwitchCreateNetworkSwitchAccessPolicyRadiusAccountingServers `json:"radiusAccountingServers,omitempty"`        // List of RADIUS accounting servers to require connecting devices to authenticate against before granting network access
	RadiusCoaSupportEnabled        *bool                                                                   `json:"radiusCoaSupportEnabled,omitempty"`        // Change of authentication for RADIUS re-authentication and disconnection
	RadiusGroupAttribute           string                                                                  `json:"radiusGroupAttribute,omitempty"`           // Acceptable values are *""* for , or *"11"* for Group Policies ACL
	RadiusServers                  *[]ResponseSwitchCreateNetworkSwitchAccessPolicyRadiusServers           `json:"radiusServers,omitempty"`                  // List of RADIUS servers to require connecting devices to authenticate against before granting network access
	RadiusTestingEnabled           *bool                                                                   `json:"radiusTestingEnabled,omitempty"`           // If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers
	URLRedirectWalledGardenEnabled *bool                                                                   `json:"urlRedirectWalledGardenEnabled,omitempty"` // Enable to restrict access for clients to a response_objectific set of IP addresses or hostnames prior to authentication
	URLRedirectWalledGardenRanges  []string                                                                `json:"urlRedirectWalledGardenRanges,omitempty"`  // IP address ranges, in CIDR notation, to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication
	VoiceVLANClients               *bool                                                                   `json:"voiceVlanClients,omitempty"`               // CDP/LLDP capable voice clients will be able to use this VLAN. Automatically true when hostMode is 'Multi-Domain'.
}
type ResponseSwitchCreateNetworkSwitchAccessPolicyDot1X struct {
	ControlDirection string `json:"controlDirection,omitempty"` // Supports either 'both' or 'inbound'. Set to 'inbound' to allow unauthorized egress on the switchport. Set to 'both' to control both traffic directions with authorization. Defaults to 'both'
}
type ResponseSwitchCreateNetworkSwitchAccessPolicyRadius struct {
	CriticalAuth             *ResponseSwitchCreateNetworkSwitchAccessPolicyRadiusCriticalAuth `json:"criticalAuth,omitempty"`             // Critical auth settings for when authentication is rejected by the RADIUS server
	FailedAuthVLANID         *int                                                             `json:"failedAuthVlanId,omitempty"`         // VLAN that clients will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
	ReAuthenticationInterval *int                                                             `json:"reAuthenticationInterval,omitempty"` // Re-authentication period in seconds. Will be null if hostMode is Multi-Auth
}
type ResponseSwitchCreateNetworkSwitchAccessPolicyRadiusCriticalAuth struct {
	DataVLANID        *int  `json:"dataVlanId,omitempty"`        // VLAN that clients who use data will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
	SuspendPortBounce *bool `json:"suspendPortBounce,omitempty"` // Enable to suspend port bounce when RADIUS servers are unreachable
	VoiceVLANID       *int  `json:"voiceVlanId,omitempty"`       // VLAN that clients who use voice will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
}
type ResponseSwitchCreateNetworkSwitchAccessPolicyRadiusAccountingServers struct {
	Host string `json:"host,omitempty"` // Public IP address of the RADIUS accounting server
	Port *int   `json:"port,omitempty"` // UDP port that the RADIUS Accounting server listens on for access requests
}
type ResponseSwitchCreateNetworkSwitchAccessPolicyRadiusServers struct {
	Host string `json:"host,omitempty"` // Public IP address of the RADIUS server
	Port *int   `json:"port,omitempty"` // UDP port that the RADIUS server listens on for access requests
}
type ResponseSwitchGetNetworkSwitchAccessPolicy struct {
	AccessPolicyType               string                                                               `json:"accessPolicyType,omitempty"`               // Access Type of the policy. Automatically 'Hybrid authentication' when hostMode is 'Multi-Domain'.
	Dot1X                          *ResponseSwitchGetNetworkSwitchAccessPolicyDot1X                     `json:"dot1x,omitempty"`                          // 802.1x Settings
	GuestPortBouncing              *bool                                                                `json:"guestPortBouncing,omitempty"`              // If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers
	GuestVLANID                    *int                                                                 `json:"guestVlanId,omitempty"`                    // ID for the guest VLAN allow unauthorized devices access to limited network resources
	HostMode                       string                                                               `json:"hostMode,omitempty"`                       // Choose the Host Mode for the access policy.
	IncreaseAccessSpeed            *bool                                                                `json:"increaseAccessSpeed,omitempty"`            // Enabling this option will make switches execute 802.1X and MAC-bypass authentication simultaneously so that clients authenticate faster. Only required when accessPolicyType is 'Hybrid Authentication.
	Name                           string                                                               `json:"name,omitempty"`                           // Name of the access policy
	Radius                         *ResponseSwitchGetNetworkSwitchAccessPolicyRadius                    `json:"radius,omitempty"`                         // Object for RADIUS Settings
	RadiusAccountingEnabled        *bool                                                                `json:"radiusAccountingEnabled,omitempty"`        // Enable to send start, interim-update and stop messages to a configured RADIUS accounting server for tracking connected clients
	RadiusAccountingServers        *[]ResponseSwitchGetNetworkSwitchAccessPolicyRadiusAccountingServers `json:"radiusAccountingServers,omitempty"`        // List of RADIUS accounting servers to require connecting devices to authenticate against before granting network access
	RadiusCoaSupportEnabled        *bool                                                                `json:"radiusCoaSupportEnabled,omitempty"`        // Change of authentication for RADIUS re-authentication and disconnection
	RadiusGroupAttribute           string                                                               `json:"radiusGroupAttribute,omitempty"`           // Acceptable values are *""* for , or *"11"* for Group Policies ACL
	RadiusServers                  *[]ResponseSwitchGetNetworkSwitchAccessPolicyRadiusServers           `json:"radiusServers,omitempty"`                  // List of RADIUS servers to require connecting devices to authenticate against before granting network access
	RadiusTestingEnabled           *bool                                                                `json:"radiusTestingEnabled,omitempty"`           // If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers
	URLRedirectWalledGardenEnabled *bool                                                                `json:"urlRedirectWalledGardenEnabled,omitempty"` // Enable to restrict access for clients to a response_objectific set of IP addresses or hostnames prior to authentication
	URLRedirectWalledGardenRanges  []string                                                             `json:"urlRedirectWalledGardenRanges,omitempty"`  // IP address ranges, in CIDR notation, to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication
	VoiceVLANClients               *bool                                                                `json:"voiceVlanClients,omitempty"`               // CDP/LLDP capable voice clients will be able to use this VLAN. Automatically true when hostMode is 'Multi-Domain'.
}
type ResponseSwitchGetNetworkSwitchAccessPolicyDot1X struct {
	ControlDirection string `json:"controlDirection,omitempty"` // Supports either 'both' or 'inbound'. Set to 'inbound' to allow unauthorized egress on the switchport. Set to 'both' to control both traffic directions with authorization. Defaults to 'both'
}
type ResponseSwitchGetNetworkSwitchAccessPolicyRadius struct {
	CriticalAuth             *ResponseSwitchGetNetworkSwitchAccessPolicyRadiusCriticalAuth `json:"criticalAuth,omitempty"`             // Critical auth settings for when authentication is rejected by the RADIUS server
	FailedAuthVLANID         *int                                                          `json:"failedAuthVlanId,omitempty"`         // VLAN that clients will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
	ReAuthenticationInterval *int                                                          `json:"reAuthenticationInterval,omitempty"` // Re-authentication period in seconds. Will be null if hostMode is Multi-Auth
}
type ResponseSwitchGetNetworkSwitchAccessPolicyRadiusCriticalAuth struct {
	DataVLANID        *int  `json:"dataVlanId,omitempty"`        // VLAN that clients who use data will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
	SuspendPortBounce *bool `json:"suspendPortBounce,omitempty"` // Enable to suspend port bounce when RADIUS servers are unreachable
	VoiceVLANID       *int  `json:"voiceVlanId,omitempty"`       // VLAN that clients who use voice will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
}
type ResponseSwitchGetNetworkSwitchAccessPolicyRadiusAccountingServers struct {
	Host string `json:"host,omitempty"` // Public IP address of the RADIUS accounting server
	Port *int   `json:"port,omitempty"` // UDP port that the RADIUS Accounting server listens on for access requests
}
type ResponseSwitchGetNetworkSwitchAccessPolicyRadiusServers struct {
	Host string `json:"host,omitempty"` // Public IP address of the RADIUS server
	Port *int   `json:"port,omitempty"` // UDP port that the RADIUS server listens on for access requests
}
type ResponseSwitchUpdateNetworkSwitchAccessPolicy struct {
	AccessPolicyType               string                                                                  `json:"accessPolicyType,omitempty"`               // Access Type of the policy. Automatically 'Hybrid authentication' when hostMode is 'Multi-Domain'.
	Dot1X                          *ResponseSwitchUpdateNetworkSwitchAccessPolicyDot1X                     `json:"dot1x,omitempty"`                          // 802.1x Settings
	GuestPortBouncing              *bool                                                                   `json:"guestPortBouncing,omitempty"`              // If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers
	GuestVLANID                    *int                                                                    `json:"guestVlanId,omitempty"`                    // ID for the guest VLAN allow unauthorized devices access to limited network resources
	HostMode                       string                                                                  `json:"hostMode,omitempty"`                       // Choose the Host Mode for the access policy.
	IncreaseAccessSpeed            *bool                                                                   `json:"increaseAccessSpeed,omitempty"`            // Enabling this option will make switches execute 802.1X and MAC-bypass authentication simultaneously so that clients authenticate faster. Only required when accessPolicyType is 'Hybrid Authentication.
	Name                           string                                                                  `json:"name,omitempty"`                           // Name of the access policy
	Radius                         *ResponseSwitchUpdateNetworkSwitchAccessPolicyRadius                    `json:"radius,omitempty"`                         // Object for RADIUS Settings
	RadiusAccountingEnabled        *bool                                                                   `json:"radiusAccountingEnabled,omitempty"`        // Enable to send start, interim-update and stop messages to a configured RADIUS accounting server for tracking connected clients
	RadiusAccountingServers        *[]ResponseSwitchUpdateNetworkSwitchAccessPolicyRadiusAccountingServers `json:"radiusAccountingServers,omitempty"`        // List of RADIUS accounting servers to require connecting devices to authenticate against before granting network access
	RadiusCoaSupportEnabled        *bool                                                                   `json:"radiusCoaSupportEnabled,omitempty"`        // Change of authentication for RADIUS re-authentication and disconnection
	RadiusGroupAttribute           string                                                                  `json:"radiusGroupAttribute,omitempty"`           // Acceptable values are *""* for , or *"11"* for Group Policies ACL
	RadiusServers                  *[]ResponseSwitchUpdateNetworkSwitchAccessPolicyRadiusServers           `json:"radiusServers,omitempty"`                  // List of RADIUS servers to require connecting devices to authenticate against before granting network access
	RadiusTestingEnabled           *bool                                                                   `json:"radiusTestingEnabled,omitempty"`           // If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers
	URLRedirectWalledGardenEnabled *bool                                                                   `json:"urlRedirectWalledGardenEnabled,omitempty"` // Enable to restrict access for clients to a response_objectific set of IP addresses or hostnames prior to authentication
	URLRedirectWalledGardenRanges  []string                                                                `json:"urlRedirectWalledGardenRanges,omitempty"`  // IP address ranges, in CIDR notation, to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication
	VoiceVLANClients               *bool                                                                   `json:"voiceVlanClients,omitempty"`               // CDP/LLDP capable voice clients will be able to use this VLAN. Automatically true when hostMode is 'Multi-Domain'.
}
type ResponseSwitchUpdateNetworkSwitchAccessPolicyDot1X struct {
	ControlDirection string `json:"controlDirection,omitempty"` // Supports either 'both' or 'inbound'. Set to 'inbound' to allow unauthorized egress on the switchport. Set to 'both' to control both traffic directions with authorization. Defaults to 'both'
}
type ResponseSwitchUpdateNetworkSwitchAccessPolicyRadius struct {
	CriticalAuth             *ResponseSwitchUpdateNetworkSwitchAccessPolicyRadiusCriticalAuth `json:"criticalAuth,omitempty"`             // Critical auth settings for when authentication is rejected by the RADIUS server
	FailedAuthVLANID         *int                                                             `json:"failedAuthVlanId,omitempty"`         // VLAN that clients will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
	ReAuthenticationInterval *int                                                             `json:"reAuthenticationInterval,omitempty"` // Re-authentication period in seconds. Will be null if hostMode is Multi-Auth
}
type ResponseSwitchUpdateNetworkSwitchAccessPolicyRadiusCriticalAuth struct {
	DataVLANID        *int  `json:"dataVlanId,omitempty"`        // VLAN that clients who use data will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
	SuspendPortBounce *bool `json:"suspendPortBounce,omitempty"` // Enable to suspend port bounce when RADIUS servers are unreachable
	VoiceVLANID       *int  `json:"voiceVlanId,omitempty"`       // VLAN that clients who use voice will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
}
type ResponseSwitchUpdateNetworkSwitchAccessPolicyRadiusAccountingServers struct {
	Host string `json:"host,omitempty"` // Public IP address of the RADIUS accounting server
	Port *int   `json:"port,omitempty"` // UDP port that the RADIUS Accounting server listens on for access requests
}
type ResponseSwitchUpdateNetworkSwitchAccessPolicyRadiusServers struct {
	Host string `json:"host,omitempty"` // Public IP address of the RADIUS server
	Port *int   `json:"port,omitempty"` // UDP port that the RADIUS server listens on for access requests
}
type ResponseSwitchGetNetworkSwitchAlternateManagementInterface struct {
	Enabled   *bool                                                                 `json:"enabled,omitempty"`   //
	Protocols []string                                                              `json:"protocols,omitempty"` //
	Switches  *[]ResponseSwitchGetNetworkSwitchAlternateManagementInterfaceSwitches `json:"switches,omitempty"`  //
	VLANID    *int                                                                  `json:"vlanId,omitempty"`    //
}
type ResponseSwitchGetNetworkSwitchAlternateManagementInterfaceSwitches struct {
	AlternateManagementIP string `json:"alternateManagementIp,omitempty"` //
	Gateway               string `json:"gateway,omitempty"`               //
	Serial                string `json:"serial,omitempty"`                //
	SubnetMask            string `json:"subnetMask,omitempty"`            //
}
type ResponseSwitchUpdateNetworkSwitchAlternateManagementInterface interface{}
type ResponseSwitchGetNetworkSwitchDhcpV4ServersSeen []ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeen // Array of ResponseSwitchGetNetworkSwitchDhcpV4ServersSeen
type ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeen struct {
	ClientID     string                                                         `json:"clientId,omitempty"`     // Client id of the server if available.
	Device       *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenDevice     `json:"device,omitempty"`       // Attributes of the server when it's a device.
	IPv4         *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenIPv4       `json:"ipv4,omitempty"`         // IPv4 attributes of the server.
	IsAllowed    *bool                                                          `json:"isAllowed,omitempty"`    // Whether the server is allowed or blocked. Always true for configured servers.
	IsConfigured *bool                                                          `json:"isConfigured,omitempty"` // Whether the server is configured.
	LastAck      *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastAck    `json:"lastAck,omitempty"`      // Attributes of the server's last ack.
	LastPacket   *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacket `json:"lastPacket,omitempty"`   // Last packet the server received.
	LastSeenAt   string                                                         `json:"lastSeenAt,omitempty"`   // Last time the server was seen.
	Mac          string                                                         `json:"mac,omitempty"`          // Mac address of the server.
	SeenBy       *[]ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenSeenBy   `json:"seenBy,omitempty"`       // Devices that saw the server.
	Type         string                                                         `json:"type,omitempty"`         // server type. Can be a 'device', 'stack', or 'discovered' (i.e client).
	VLAN         *int                                                           `json:"vlan,omitempty"`         // Vlan id of the server.
}
type ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenDevice struct {
	Interface *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenDeviceInterface `json:"interface,omitempty"` // Interface attributes of the server. Only for configured servers.
	Name      string                                                              `json:"name,omitempty"`      // Device name.
	Serial    string                                                              `json:"serial,omitempty"`    // Device serial.
	URL       string                                                              `json:"url,omitempty"`       // Url link to device.
}
type ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenDeviceInterface struct {
	Name string `json:"name,omitempty"` // Interface name.
	URL  string `json:"url,omitempty"`  // Url link to interface.
}
type ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenIPv4 struct {
	Address string `json:"address,omitempty"` // IPv4 address of the server.
	Gateway string `json:"gateway,omitempty"` // IPv4 gateway address of the server.
	Subnet  string `json:"subnet,omitempty"`  // Subnet of the server.
}
type ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastAck struct {
	IPv4 *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastAckIPv4 `json:"ipv4,omitempty"` // IPv4 attributes of the last ack.
	Ts   string                                                          `json:"ts,omitempty"`   // Last time the server was acked.
}
type ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastAckIPv4 struct {
	Address string `json:"address,omitempty"` // IPv4 address of the last ack.
}
type ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacket struct {
	Destination *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketDestination `json:"destination,omitempty"` // Destination of the packet.
	Ethernet    *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketEthernet    `json:"ethernet,omitempty"`    // Additional ethernet attributes of the packet.
	Fields      *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketFields      `json:"fields,omitempty"`      // DHCP-specific fields of the packet.
	IP          *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketIP          `json:"ip,omitempty"`          // Additional IP attributes of the packet.
	Source      *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketSource      `json:"source,omitempty"`      // Source of the packet.
	Type        string                                                                    `json:"type,omitempty"`        // Packet type.
	UDP         *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketUDP         `json:"udp,omitempty"`         // UDP attributes of the packet.
}
type ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketDestination struct {
	IPv4 *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketDestinationIPv4 `json:"ipv4,omitempty"` // Destination ipv4 attributes of the packet.
	Mac  string                                                                        `json:"mac,omitempty"`  // Destination mac address of the packet.
	Port *int                                                                          `json:"port,omitempty"` // Destination port of the packet.
}
type ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketDestinationIPv4 struct {
	Address string `json:"address,omitempty"` // Destination ipv4 address of the packet.
}
type ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketEthernet struct {
	Type string `json:"type,omitempty"` // Ethernet type of the packet.
}
type ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketFields struct {
	Chaddr      string                                                                        `json:"chaddr,omitempty"`      // Client hardware address of the packet.
	Ciaddr      string                                                                        `json:"ciaddr,omitempty"`      // Client IP address of the packet.
	Flags       string                                                                        `json:"flags,omitempty"`       // Packet flags.
	Giaddr      string                                                                        `json:"giaddr,omitempty"`      // Gateway IP address of the packet.
	Hlen        *int                                                                          `json:"hlen,omitempty"`        // Hardware length of the packet.
	Hops        *int                                                                          `json:"hops,omitempty"`        // Number of hops the packet took.
	Htype       *int                                                                          `json:"htype,omitempty"`       // Hardware type code of the packet.
	MagicCookie string                                                                        `json:"magicCookie,omitempty"` // Magic cookie of the packet.
	Op          *int                                                                          `json:"op,omitempty"`          // Operation code of the packet.
	Options     *[]ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketFieldsOptions `json:"options,omitempty"`     // Additional DHCP options of the packet.
	Secs        *int                                                                          `json:"secs,omitempty"`        // Number of seconds since receiving the packet.
	Siaddr      string                                                                        `json:"siaddr,omitempty"`      // Server IP address of the packet.
	Sname       string                                                                        `json:"sname,omitempty"`       // Server identifier address of the packet.
	Xid         string                                                                        `json:"xid,omitempty"`         // Transaction id of the packet.
	Yiaddr      string                                                                        `json:"yiaddr,omitempty"`      // Assigned IP address of the packet.
}
type ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketFieldsOptions struct {
	Name  string `json:"name,omitempty"`  // Option name.
	Value string `json:"value,omitempty"` // Option value.
}
type ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketIP struct {
	Dscp         *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketIPDscp `json:"dscp,omitempty"`         // DSCP attributes of the packet.
	HeaderLength *int                                                                 `json:"headerLength,omitempty"` // IP header length of the packet.
	ID           string                                                               `json:"id,omitempty"`           // IP ID of the packet.
	Length       *int                                                                 `json:"length,omitempty"`       // IP length of the packet.
	Protocol     *int                                                                 `json:"protocol,omitempty"`     // IP protocol number of the packet.
	Ttl          *int                                                                 `json:"ttl,omitempty"`          // Time to live of the packet.
	Version      *int                                                                 `json:"version,omitempty"`      // IP version of the packet.
}
type ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketIPDscp struct {
	Ecn *int `json:"ecn,omitempty"` // ECN number of the packet.
	Tag *int `json:"tag,omitempty"` // DSCP tag number of the packet.
}
type ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketSource struct {
	IPv4 *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketSourceIPv4 `json:"ipv4,omitempty"` // Source ipv4 attributes of the packet.
	Mac  string                                                                   `json:"mac,omitempty"`  // Source mac address of the packet.
	Port *int                                                                     `json:"port,omitempty"` // Source port of the packet.
}
type ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketSourceIPv4 struct {
	Address string `json:"address,omitempty"` // Source ipv4 address of the packet.
}
type ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketUDP struct {
	Checksum string `json:"checksum,omitempty"` // UDP checksum of the packet.
	Length   *int   `json:"length,omitempty"`   // UDP length of the packet.
}
type ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenSeenBy struct {
	Name   string `json:"name,omitempty"`   // Device name.
	Serial string `json:"serial,omitempty"` // Device serial.
	URL    string `json:"url,omitempty"`    // Url link to device.
}
type ResponseSwitchGetNetworkSwitchDhcpServerPolicy struct {
	Alerts         *ResponseSwitchGetNetworkSwitchDhcpServerPolicyAlerts        `json:"alerts,omitempty"`         //
	AllowedServers []string                                                     `json:"allowedServers,omitempty"` //
	ArpInspection  *ResponseSwitchGetNetworkSwitchDhcpServerPolicyArpInspection `json:"arpInspection,omitempty"`  //
	BlockedServers []string                                                     `json:"blockedServers,omitempty"` //
	DefaultPolicy  string                                                       `json:"defaultPolicy,omitempty"`  //
}
type ResponseSwitchGetNetworkSwitchDhcpServerPolicyAlerts struct {
	Email *ResponseSwitchGetNetworkSwitchDhcpServerPolicyAlertsEmail `json:"email,omitempty"` //
}
type ResponseSwitchGetNetworkSwitchDhcpServerPolicyAlertsEmail struct {
	Enabled *bool `json:"enabled,omitempty"` //
}
type ResponseSwitchGetNetworkSwitchDhcpServerPolicyArpInspection struct {
	Enabled           *bool    `json:"enabled,omitempty"`           //
	UnsupportedModels []string `json:"unsupportedModels,omitempty"` //
}
type ResponseSwitchUpdateNetworkSwitchDhcpServerPolicy interface{}
type ResponseSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers []ResponseItemSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers // Array of ResponseSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers
type ResponseItemSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers struct {
	IPv4            *ResponseItemSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersIPv4 `json:"ipv4,omitempty"`            // IPv4 attributes of the trusted server.
	Mac             string                                                                             `json:"mac,omitempty"`             // Mac address of the trusted server.
	TrustedServerID string                                                                             `json:"trustedServerId,omitempty"` // ID of the trusted server.
	VLAN            *int                                                                               `json:"vlan,omitempty"`            // Vlan ID of the trusted server.
}
type ResponseItemSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersIPv4 struct {
	Address string `json:"address,omitempty"` // IPv4 address of the trusted server.
}
type ResponseSwitchCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer struct {
	IPv4            *ResponseSwitchCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerIPv4 `json:"ipv4,omitempty"`            // IPv4 attributes of the trusted server.
	Mac             string                                                                           `json:"mac,omitempty"`             // Mac address of the trusted server.
	TrustedServerID string                                                                           `json:"trustedServerId,omitempty"` // ID of the trusted server.
	VLAN            *int                                                                             `json:"vlan,omitempty"`            // Vlan ID of the trusted server.
}
type ResponseSwitchCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerIPv4 struct {
	Address string `json:"address,omitempty"` // IPv4 address of the trusted server.
}
type ResponseSwitchUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer struct {
	IPv4            *ResponseSwitchUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerIPv4 `json:"ipv4,omitempty"`            // IPv4 attributes of the trusted server.
	Mac             string                                                                           `json:"mac,omitempty"`             // Mac address of the trusted server.
	TrustedServerID string                                                                           `json:"trustedServerId,omitempty"` // ID of the trusted server.
	VLAN            *int                                                                             `json:"vlan,omitempty"`            // Vlan ID of the trusted server.
}
type ResponseSwitchUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerIPv4 struct {
	Address string `json:"address,omitempty"` // IPv4 address of the trusted server.
}
type ResponseSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice []ResponseItemSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice // Array of ResponseSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice
type ResponseItemSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice struct {
	HasTrustedPort     *bool  `json:"hasTrustedPort,omitempty"`     // Whether this switch has a trusted DAI port. Always false if supportsInspection is false.
	Name               string `json:"name,omitempty"`               // Switch name.
	Serial             string `json:"serial,omitempty"`             // Switch serial.
	SupportsInspection *bool  `json:"supportsInspection,omitempty"` // Whether this switch supports Dynamic ARP Inspection.
	URL                string `json:"url,omitempty"`                // Url link to switch.
}
type ResponseSwitchGetNetworkSwitchDscpToCosMappings struct {
	Mappings *[]ResponseSwitchGetNetworkSwitchDscpToCosMappingsMappings `json:"mappings,omitempty"` //
}
type ResponseSwitchGetNetworkSwitchDscpToCosMappingsMappings struct {
	Cos   *int   `json:"cos,omitempty"`   //
	Dscp  *int   `json:"dscp,omitempty"`  //
	Title string `json:"title,omitempty"` //
}
type ResponseSwitchUpdateNetworkSwitchDscpToCosMappings interface{}
type ResponseSwitchGetNetworkSwitchLinkAggregations []ResponseItemSwitchGetNetworkSwitchLinkAggregations // Array of ResponseSwitchGetNetworkSwitchLinkAggregations
type ResponseItemSwitchGetNetworkSwitchLinkAggregations struct {
	ID          string                                                           `json:"id,omitempty"`          //
	SwitchPorts *[]ResponseItemSwitchGetNetworkSwitchLinkAggregationsSwitchPorts `json:"switchPorts,omitempty"` // Array of switch or stack ports for creating aggregation group. Minimum 2 and maximum 8 ports are supported.
}
type ResponseItemSwitchGetNetworkSwitchLinkAggregationsSwitchPorts struct {
	PortID string `json:"portId,omitempty"` // Port identifier of switch port. For modules, the identifier is "SlotNumber_ModuleType_PortNumber" (Ex: "1_8X10G_1"), otherwise it is just the port number (Ex: "8").
	Serial string `json:"serial,omitempty"` // Serial number of the switch.
}
type ResponseSwitchCreateNetworkSwitchLinkAggregation struct {
	ID          string                                                         `json:"id,omitempty"`          //
	SwitchPorts *[]ResponseSwitchCreateNetworkSwitchLinkAggregationSwitchPorts `json:"switchPorts,omitempty"` //
}
type ResponseSwitchCreateNetworkSwitchLinkAggregationSwitchPorts struct {
	PortID string `json:"portId,omitempty"` //
	Serial string `json:"serial,omitempty"` //
}
type ResponseSwitchUpdateNetworkSwitchLinkAggregation interface{}
type ResponseSwitchGetNetworkSwitchMtu struct {
	DefaultMtuSize *int                                          `json:"defaultMtuSize,omitempty"` // MTU size for the entire network. Default value is 9578.
	Overrides      *[]ResponseSwitchGetNetworkSwitchMtuOverrides `json:"overrides,omitempty"`      // Override MTU size for individual switches or switch profiles.       An empty array will clear overrides.
}
type ResponseSwitchGetNetworkSwitchMtuOverrides struct {
	MtuSize        *int     `json:"mtuSize,omitempty"`        // MTU size for the switches or switch profiles.
	SwitchProfiles []string `json:"switchProfiles,omitempty"` // List of switch profile IDs. Applicable only for template network.
	Switches       []string `json:"switches,omitempty"`       // List of switch serials. Applicable only for switch network.
}
type ResponseSwitchUpdateNetworkSwitchMtu interface{}
type ResponseSwitchGetNetworkSwitchPortSchedules []ResponseItemSwitchGetNetworkSwitchPortSchedules // Array of ResponseSwitchGetNetworkSwitchPortSchedules
type ResponseItemSwitchGetNetworkSwitchPortSchedules struct {
	ID           string                                                       `json:"id,omitempty"`           //
	Name         string                                                       `json:"name,omitempty"`         // The name for your port schedule. Required
	PortSchedule *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortSchedule `json:"portSchedule,omitempty"` //     The schedule for switch port scheduling. Schedules are applied to days of the week.     When it's empty, default schedule with all days of a week are configured.     Any unspecified day in the schedule is added as a default schedule configuration of the day.
}
type ResponseItemSwitchGetNetworkSwitchPortSchedulesPortSchedule struct {
	Friday    *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleFriday    `json:"friday,omitempty"`    // The schedule object for Friday.
	Monday    *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleMonday    `json:"monday,omitempty"`    // The schedule object for Monday.
	Saturday  *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleSaturday  `json:"saturday,omitempty"`  // The schedule object for Saturday.
	Sunday    *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleSunday    `json:"sunday,omitempty"`    // The schedule object for Sunday.
	Thursday  *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleThursday  `json:"thursday,omitempty"`  // The schedule object for Thursday.
	Tuesday   *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleTuesday   `json:"tuesday,omitempty"`   // The schedule object for Tuesday.
	Wednesday *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleWednesday `json:"wednesday,omitempty"` // The schedule object for Wednesday.
}
type ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleFriday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
}
type ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleMonday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
}
type ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleSaturday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
}
type ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleSunday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
}
type ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleThursday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
}
type ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleTuesday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
}
type ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleWednesday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
}
type ResponseSwitchCreateNetworkSwitchPortSchedule interface{}
type ResponseSwitchUpdateNetworkSwitchPortSchedule interface{}
type ResponseSwitchGetNetworkSwitchQosRules []ResponseItemSwitchGetNetworkSwitchQosRules // Array of ResponseSwitchGetNetworkSwitchQosRules
type ResponseItemSwitchGetNetworkSwitchQosRules struct {
	Dscp         *int   `json:"dscp,omitempty"`         //
	DstPort      *int   `json:"dstPort,omitempty"`      //
	DstPortRange string `json:"dstPortRange,omitempty"` //
	ID           string `json:"id,omitempty"`           //
	Protocol     string `json:"protocol,omitempty"`     //
	SrcPort      *int   `json:"srcPort,omitempty"`      //
	SrcPortRange string `json:"srcPortRange,omitempty"` //
	VLAN         *int   `json:"vlan,omitempty"`         //
}
type ResponseSwitchCreateNetworkSwitchQosRule struct {
	Dscp         *int   `json:"dscp,omitempty"`         //
	DstPort      *int   `json:"dstPort,omitempty"`      //
	DstPortRange string `json:"dstPortRange,omitempty"` //
	ID           string `json:"id,omitempty"`           //
	Protocol     string `json:"protocol,omitempty"`     //
	SrcPort      *int   `json:"srcPort,omitempty"`      //
	SrcPortRange string `json:"srcPortRange,omitempty"` //
	VLAN         *int   `json:"vlan,omitempty"`         //
}
type ResponseSwitchGetNetworkSwitchQosRulesOrder struct {
	RuleIDs []string `json:"ruleIds,omitempty"` //
}
type ResponseSwitchUpdateNetworkSwitchQosRulesOrder interface{}
type ResponseSwitchGetNetworkSwitchQosRule struct {
	Dscp         *int   `json:"dscp,omitempty"`         //
	DstPort      *int   `json:"dstPort,omitempty"`      //
	DstPortRange string `json:"dstPortRange,omitempty"` //
	ID           string `json:"id,omitempty"`           //
	Protocol     string `json:"protocol,omitempty"`     //
	SrcPort      *int   `json:"srcPort,omitempty"`      //
	SrcPortRange string `json:"srcPortRange,omitempty"` //
	VLAN         *int   `json:"vlan,omitempty"`         //
}
type ResponseSwitchUpdateNetworkSwitchQosRule interface{}
type ResponseSwitchGetNetworkSwitchRoutingMulticast struct {
	DefaultSettings *ResponseSwitchGetNetworkSwitchRoutingMulticastDefaultSettings `json:"defaultSettings,omitempty"` //
	Overrides       *[]ResponseSwitchGetNetworkSwitchRoutingMulticastOverrides     `json:"overrides,omitempty"`       //
}
type ResponseSwitchGetNetworkSwitchRoutingMulticastDefaultSettings struct {
	FloodUnknownMulticastTrafficEnabled *bool `json:"floodUnknownMulticastTrafficEnabled,omitempty"` //
	IgmpSnoopingEnabled                 *bool `json:"igmpSnoopingEnabled,omitempty"`                 //
}
type ResponseSwitchGetNetworkSwitchRoutingMulticastOverrides struct {
	FloodUnknownMulticastTrafficEnabled *bool    `json:"floodUnknownMulticastTrafficEnabled,omitempty"` //
	IgmpSnoopingEnabled                 *bool    `json:"igmpSnoopingEnabled,omitempty"`                 //
	Switches                            []string `json:"switches,omitempty"`                            //
}
type ResponseSwitchUpdateNetworkSwitchRoutingMulticast interface{}
type ResponseSwitchGetNetworkSwitchRoutingMulticastRendezvousPoints []ResponseItemSwitchGetNetworkSwitchRoutingMulticastRendezvousPoints // Array of ResponseSwitchGetNetworkSwitchRoutingMulticastRendezvousPoints
type ResponseItemSwitchGetNetworkSwitchRoutingMulticastRendezvousPoints struct {
	InterfaceIP       string `json:"interfaceIp,omitempty"`       //
	InterfaceName     string `json:"interfaceName,omitempty"`     //
	MulticastGroup    string `json:"multicastGroup,omitempty"`    //
	RendezvousPointID string `json:"rendezvousPointId,omitempty"` //
	Serial            string `json:"serial,omitempty"`            //
}
type ResponseSwitchCreateNetworkSwitchRoutingMulticastRendezvousPoint struct {
	InterfaceIP       string `json:"interfaceIp,omitempty"`       //
	MulticastGroup    string `json:"multicastGroup,omitempty"`    //
	RendezvousPointID string `json:"rendezvousPointId,omitempty"` //
}
type ResponseSwitchGetNetworkSwitchRoutingMulticastRendezvousPoint struct {
	InterfaceIP       string `json:"interfaceIp,omitempty"`       //
	InterfaceName     string `json:"interfaceName,omitempty"`     //
	MulticastGroup    string `json:"multicastGroup,omitempty"`    //
	RendezvousPointID string `json:"rendezvousPointId,omitempty"` //
	Serial            string `json:"serial,omitempty"`            //
}
type ResponseSwitchUpdateNetworkSwitchRoutingMulticastRendezvousPoint interface{}
type ResponseSwitchGetNetworkSwitchRoutingOspf struct {
	Areas                    *[]ResponseSwitchGetNetworkSwitchRoutingOspfAreas              `json:"areas,omitempty"`                    //
	DeadTimerInSeconds       *int                                                           `json:"deadTimerInSeconds,omitempty"`       //
	Enabled                  *bool                                                          `json:"enabled,omitempty"`                  //
	HelloTimerInSeconds      *int                                                           `json:"helloTimerInSeconds,omitempty"`      //
	Md5AuthenticationEnabled *bool                                                          `json:"md5AuthenticationEnabled,omitempty"` //
	Md5AuthenticationKey     *ResponseSwitchGetNetworkSwitchRoutingOspfMd5AuthenticationKey `json:"md5AuthenticationKey,omitempty"`     //
	V3                       *ResponseSwitchGetNetworkSwitchRoutingOspfV3                   `json:"v3,omitempty"`                       //
}
type ResponseSwitchGetNetworkSwitchRoutingOspfAreas struct {
	AreaID   string `json:"areaId,omitempty"`   //
	AreaName string `json:"areaName,omitempty"` //
	AreaType string `json:"areaType,omitempty"` //
}
type ResponseSwitchGetNetworkSwitchRoutingOspfMd5AuthenticationKey struct {
	ID         *int   `json:"id,omitempty"`         //
	Passphrase string `json:"passphrase,omitempty"` //
}
type ResponseSwitchGetNetworkSwitchRoutingOspfV3 struct {
	Areas               *[]ResponseSwitchGetNetworkSwitchRoutingOspfV3Areas `json:"areas,omitempty"`               //
	DeadTimerInSeconds  *int                                                `json:"deadTimerInSeconds,omitempty"`  //
	Enabled             *bool                                               `json:"enabled,omitempty"`             //
	HelloTimerInSeconds *int                                                `json:"helloTimerInSeconds,omitempty"` //
}
type ResponseSwitchGetNetworkSwitchRoutingOspfV3Areas struct {
	AreaID   string `json:"areaId,omitempty"`   //
	AreaName string `json:"areaName,omitempty"` //
	AreaType string `json:"areaType,omitempty"` //
}
type ResponseSwitchUpdateNetworkSwitchRoutingOspf interface{}
type ResponseSwitchGetNetworkSwitchSettings struct {
	PowerExceptions  *[]ResponseSwitchGetNetworkSwitchSettingsPowerExceptions `json:"powerExceptions,omitempty"`  // Exceptions on a per switch basis to "useCombinedPower"
	UseCombinedPower *bool                                                    `json:"useCombinedPower,omitempty"` // The use Combined Power as the default behavior of secondary power supplies on supported devices.
	VLAN             *int                                                     `json:"vlan,omitempty"`             // Management VLAN
}
type ResponseSwitchGetNetworkSwitchSettingsPowerExceptions struct {
	PowerType string `json:"powerType,omitempty"` // Per switch exception (combined, redundant, useNetworkSetting)
	Serial    string `json:"serial,omitempty"`    // Serial number of the switch
}
type ResponseSwitchUpdateNetworkSwitchSettings struct {
	PowerExceptions  *[]ResponseSwitchUpdateNetworkSwitchSettingsPowerExceptions `json:"powerExceptions,omitempty"`  // Exceptions on a per switch basis to "useCombinedPower"
	UseCombinedPower *bool                                                       `json:"useCombinedPower,omitempty"` // The use Combined Power as the default behavior of secondary power supplies on supported devices.
	VLAN             *int                                                        `json:"vlan,omitempty"`             // Management VLAN
}
type ResponseSwitchUpdateNetworkSwitchSettingsPowerExceptions struct {
	PowerType string `json:"powerType,omitempty"` // Per switch exception (combined, redundant, useNetworkSetting)
	Serial    string `json:"serial,omitempty"`    // Serial number of the switch
}
type ResponseSwitchGetNetworkSwitchStacks []ResponseItemSwitchGetNetworkSwitchStacks // Array of ResponseSwitchGetNetworkSwitchStacks
type ResponseItemSwitchGetNetworkSwitchStacks struct {
	ID      string   `json:"id,omitempty"`      //
	Name    string   `json:"name,omitempty"`    //
	Serials []string `json:"serials,omitempty"` //
}
type ResponseSwitchCreateNetworkSwitchStack interface{}
type ResponseSwitchGetNetworkSwitchStack struct {
	ID      string   `json:"id,omitempty"`      // Switch stacks id
	Name    string   `json:"name,omitempty"`    // Switch stacks name
	Serials []string `json:"serials,omitempty"` // Serials of the switches in the switch stack
}
type ResponseSwitchAddNetworkSwitchStack struct {
	ID      string   `json:"id,omitempty"`      // Switch stacks id
	Name    string   `json:"name,omitempty"`    // Switch stacks name
	Serials []string `json:"serials,omitempty"` // Serials of the switches in the switch stack
}
type ResponseSwitchRemoveNetworkSwitchStack interface{}
type ResponseSwitchGetNetworkSwitchStackRoutingInterfaces []ResponseItemSwitchGetNetworkSwitchStackRoutingInterfaces // Array of ResponseSwitchGetNetworkSwitchStackRoutingInterfaces
type ResponseItemSwitchGetNetworkSwitchStackRoutingInterfaces struct {
	DefaultGateway   string                                                                `json:"defaultGateway,omitempty"`   //
	InterfaceID      string                                                                `json:"interfaceId,omitempty"`      //
	InterfaceIP      string                                                                `json:"interfaceIp,omitempty"`      //
	IPv6             *ResponseItemSwitchGetNetworkSwitchStackRoutingInterfacesIPv6         `json:"ipv6,omitempty"`             //
	MulticastRouting string                                                                `json:"multicastRouting,omitempty"` //
	Name             string                                                                `json:"name,omitempty"`             //
	OspfSettings     *ResponseItemSwitchGetNetworkSwitchStackRoutingInterfacesOspfSettings `json:"ospfSettings,omitempty"`     //
	OspfV3           *ResponseItemSwitchGetNetworkSwitchStackRoutingInterfacesOspfV3       `json:"ospfV3,omitempty"`           //
	Subnet           string                                                                `json:"subnet,omitempty"`           //
	VLANID           *int                                                                  `json:"vlanId,omitempty"`           //
}
type ResponseItemSwitchGetNetworkSwitchStackRoutingInterfacesIPv6 struct {
	Address        string `json:"address,omitempty"`        //
	AssignmentMode string `json:"assignmentMode,omitempty"` //
	Gateway        string `json:"gateway,omitempty"`        //
	Prefix         string `json:"prefix,omitempty"`         //
}
type ResponseItemSwitchGetNetworkSwitchStackRoutingInterfacesOspfSettings struct {
	Area             string `json:"area,omitempty"`             //
	Cost             *int   `json:"cost,omitempty"`             //
	IsPassiveEnabled *bool  `json:"isPassiveEnabled,omitempty"` //
}
type ResponseItemSwitchGetNetworkSwitchStackRoutingInterfacesOspfV3 struct {
	Area             string `json:"area,omitempty"`             //
	Cost             *int   `json:"cost,omitempty"`             //
	IsPassiveEnabled *bool  `json:"isPassiveEnabled,omitempty"` //
}
type ResponseSwitchCreateNetworkSwitchStackRoutingInterface interface{}
type ResponseSwitchGetNetworkSwitchStackRoutingInterface struct {
	DefaultGateway   string                                                           `json:"defaultGateway,omitempty"`   //
	InterfaceID      string                                                           `json:"interfaceId,omitempty"`      //
	InterfaceIP      string                                                           `json:"interfaceIp,omitempty"`      //
	IPv6             *ResponseSwitchGetNetworkSwitchStackRoutingInterfaceIPv6         `json:"ipv6,omitempty"`             //
	MulticastRouting string                                                           `json:"multicastRouting,omitempty"` //
	Name             string                                                           `json:"name,omitempty"`             //
	OspfSettings     *ResponseSwitchGetNetworkSwitchStackRoutingInterfaceOspfSettings `json:"ospfSettings,omitempty"`     //
	OspfV3           *ResponseSwitchGetNetworkSwitchStackRoutingInterfaceOspfV3       `json:"ospfV3,omitempty"`           //
	Subnet           string                                                           `json:"subnet,omitempty"`           //
	VLANID           *int                                                             `json:"vlanId,omitempty"`           //
}
type ResponseSwitchGetNetworkSwitchStackRoutingInterfaceIPv6 struct {
	Address        string `json:"address,omitempty"`        //
	AssignmentMode string `json:"assignmentMode,omitempty"` //
	Gateway        string `json:"gateway,omitempty"`        //
	Prefix         string `json:"prefix,omitempty"`         //
}
type ResponseSwitchGetNetworkSwitchStackRoutingInterfaceOspfSettings struct {
	Area             string `json:"area,omitempty"`             //
	Cost             *int   `json:"cost,omitempty"`             //
	IsPassiveEnabled *bool  `json:"isPassiveEnabled,omitempty"` //
}
type ResponseSwitchGetNetworkSwitchStackRoutingInterfaceOspfV3 struct {
	Area             string `json:"area,omitempty"`             //
	Cost             *int   `json:"cost,omitempty"`             //
	IsPassiveEnabled *bool  `json:"isPassiveEnabled,omitempty"` //
}
type ResponseSwitchUpdateNetworkSwitchStackRoutingInterface interface{}
type ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcp struct {
	BootFileName         string                                                                       `json:"bootFileName,omitempty"`         //
	BootNextServer       string                                                                       `json:"bootNextServer,omitempty"`       //
	BootOptionsEnabled   *bool                                                                        `json:"bootOptionsEnabled,omitempty"`   //
	DhcpLeaseTime        string                                                                       `json:"dhcpLeaseTime,omitempty"`        //
	DhcpMode             string                                                                       `json:"dhcpMode,omitempty"`             //
	DhcpOptions          *[]ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpDhcpOptions        `json:"dhcpOptions,omitempty"`          //
	DNSCustomNameservers []string                                                                     `json:"dnsCustomNameservers,omitempty"` //
	DNSNameserversOption string                                                                       `json:"dnsNameserversOption,omitempty"` //
	FixedIPAssignments   *[]ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpFixedIPAssignments `json:"fixedIpAssignments,omitempty"`   //
	ReservedIPRanges     *[]ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpReservedIPRanges   `json:"reservedIpRanges,omitempty"`     //
}
type ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpDhcpOptions struct {
	Code  string `json:"code,omitempty"`  //
	Type  string `json:"type,omitempty"`  //
	Value string `json:"value,omitempty"` //
}
type ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpFixedIPAssignments struct {
	IP   string `json:"ip,omitempty"`   //
	Mac  string `json:"mac,omitempty"`  //
	Name string `json:"name,omitempty"` //
}
type ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpReservedIPRanges struct {
	Comment string `json:"comment,omitempty"` //
	End     string `json:"end,omitempty"`     //
	Start   string `json:"start,omitempty"`   //
}
type ResponseSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcp interface{}
type ResponseSwitchGetNetworkSwitchStackRoutingStaticRoutes []ResponseItemSwitchGetNetworkSwitchStackRoutingStaticRoutes // Array of ResponseSwitchGetNetworkSwitchStackRoutingStaticRoutes
type ResponseItemSwitchGetNetworkSwitchStackRoutingStaticRoutes struct {
	AdvertiseViaOspfEnabled     *bool  `json:"advertiseViaOspfEnabled,omitempty"`     //
	Name                        string `json:"name,omitempty"`                        //
	NextHopIP                   string `json:"nextHopIp,omitempty"`                   //
	PreferOverOspfRoutesEnabled *bool  `json:"preferOverOspfRoutesEnabled,omitempty"` //
	StaticRouteID               string `json:"staticRouteId,omitempty"`               //
	Subnet                      string `json:"subnet,omitempty"`                      //
}
type ResponseSwitchCreateNetworkSwitchStackRoutingStaticRoute interface{}
type ResponseSwitchGetNetworkSwitchStackRoutingStaticRoute struct {
	AdvertiseViaOspfEnabled     *bool  `json:"advertiseViaOspfEnabled,omitempty"`     //
	Name                        string `json:"name,omitempty"`                        //
	NextHopIP                   string `json:"nextHopIp,omitempty"`                   //
	PreferOverOspfRoutesEnabled *bool  `json:"preferOverOspfRoutesEnabled,omitempty"` //
	StaticRouteID               string `json:"staticRouteId,omitempty"`               //
	Subnet                      string `json:"subnet,omitempty"`                      //
}
type ResponseSwitchUpdateNetworkSwitchStackRoutingStaticRoute interface{}
type ResponseSwitchGetNetworkSwitchStormControl struct {
	BroadcastThreshold      *int `json:"broadcastThreshold,omitempty"`      // Broadcast threshold.
	MulticastThreshold      *int `json:"multicastThreshold,omitempty"`      // Multicast threshold.
	UnknownUnicastThreshold *int `json:"unknownUnicastThreshold,omitempty"` // Unknown Unicast threshold.
}
type ResponseSwitchUpdateNetworkSwitchStormControl interface{}
type ResponseSwitchGetNetworkSwitchStp struct {
	RstpEnabled       *bool                                                 `json:"rstpEnabled,omitempty"`       //
	StpBridgePriority *[]ResponseSwitchGetNetworkSwitchStpStpBridgePriority `json:"stpBridgePriority,omitempty"` //
}
type ResponseSwitchGetNetworkSwitchStpStpBridgePriority struct {
	StpPriority *int     `json:"stpPriority,omitempty"` //
	Switches    []string `json:"switches,omitempty"`    //
}
type ResponseSwitchUpdateNetworkSwitchStp interface{}
type ResponseSwitchGetOrganizationConfigTemplateSwitchProfiles struct {
	Model           string `json:"model,omitempty"`           // Switch model
	Name            string `json:"name,omitempty"`            // Switch profile name
	SwitchProfileID string `json:"switchProfileId,omitempty"` // Switch profile id
}
type ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePorts []ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePorts // Array of ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePorts
type ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePorts struct {
	AccessPolicyNumber          *int                                                                      `json:"accessPolicyNumber,omitempty"`          // The number of a custom access policy to configure on the switch profile port. Only applicable when 'accessPolicyType' is 'Custom access policy'.
	AccessPolicyType            string                                                                    `json:"accessPolicyType,omitempty"`            // The type of the access policy of the switch profile port. Only applicable to access ports. Can be one of 'Open', 'Custom access policy', 'MAC allow list' or 'Sticky MAC allow list'.
	AllowedVLANs                string                                                                    `json:"allowedVlans,omitempty"`                // The VLANs allowed on the switch profile port. Only applicable to trunk ports.
	DaiTrusted                  *bool                                                                     `json:"daiTrusted,omitempty"`                  // If true, ARP packets for this port will be considered trusted, and Dynamic ARP Inspection will allow the traffic.
	Enabled                     *bool                                                                     `json:"enabled,omitempty"`                     // The status of the switch profile port.
	FlexibleStackingEnabled     *bool                                                                     `json:"flexibleStackingEnabled,omitempty"`     // For supported switches (e.g. MS420/MS425), whether or not the port has flexible stacking enabled.
	IsolationEnabled            *bool                                                                     `json:"isolationEnabled,omitempty"`            // The isolation status of the switch profile port.
	LinkNegotiation             string                                                                    `json:"linkNegotiation,omitempty"`             // The link speed for the switch profile port.
	LinkNegotiationCapabilities []string                                                                  `json:"linkNegotiationCapabilities,omitempty"` // Available link speeds for the switch profile port.
	MacAllowList                []string                                                                  `json:"macAllowList,omitempty"`                // Only devices with MAC addresses specified in this list will have access to this port. Up to 20 MAC addresses can be defined. Only applicable when 'accessPolicyType' is 'MAC allow list'.
	Name                        string                                                                    `json:"name,omitempty"`                        // The name of the switch profile port.
	PoeEnabled                  *bool                                                                     `json:"poeEnabled,omitempty"`                  // The PoE status of the switch profile port.
	PortID                      string                                                                    `json:"portId,omitempty"`                      // The identifier of the switch profile port.
	PortScheduleID              string                                                                    `json:"portScheduleId,omitempty"`              // The ID of the port schedule. A value of null will clear the port schedule.
	Profile                     *ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePortsProfile `json:"profile,omitempty"`                     // Profile attributes
	RstpEnabled                 *bool                                                                     `json:"rstpEnabled,omitempty"`                 // The rapid spanning tree protocol status.
	StickyMacAllowList          []string                                                                  `json:"stickyMacAllowList,omitempty"`          // The initial list of MAC addresses for sticky Mac allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
	StickyMacAllowListLimit     *int                                                                      `json:"stickyMacAllowListLimit,omitempty"`     // The maximum number of MAC addresses for sticky MAC allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
	StormControlEnabled         *bool                                                                     `json:"stormControlEnabled,omitempty"`         // The storm control status of the switch profile port.
	StpGuard                    string                                                                    `json:"stpGuard,omitempty"`                    // The state of the STP guard ('disabled', 'root guard', 'bpdu guard' or 'loop guard').
	Tags                        []string                                                                  `json:"tags,omitempty"`                        // The list of tags of the switch profile port.
	Type                        string                                                                    `json:"type,omitempty"`                        // The type of the switch profile port ('trunk' or 'access').
	Udld                        string                                                                    `json:"udld,omitempty"`                        // The action to take when Unidirectional Link is detected (Alert only, Enforce). Default configuration is Alert only.
	VLAN                        *int                                                                      `json:"vlan,omitempty"`                        // The VLAN of the switch profile port. A null value will clear the value set for trunk ports.
	VoiceVLAN                   *int                                                                      `json:"voiceVlan,omitempty"`                   // The voice VLAN of the switch profile port. Only applicable to access ports.
}
type ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePortsProfile struct {
	Enabled *bool  `json:"enabled,omitempty"` // When enabled, override this port's configuration with a port profile.
	ID      string `json:"id,omitempty"`      // When enabled, the ID of the port profile used to override the port's configuration.
	Iname   string `json:"iname,omitempty"`   // When enabled, the IName of the profile.
}
type ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePort struct {
	AccessPolicyNumber          *int                                                                 `json:"accessPolicyNumber,omitempty"`          // The number of a custom access policy to configure on the switch profile port. Only applicable when 'accessPolicyType' is 'Custom access policy'.
	AccessPolicyType            string                                                               `json:"accessPolicyType,omitempty"`            // The type of the access policy of the switch profile port. Only applicable to access ports. Can be one of 'Open', 'Custom access policy', 'MAC allow list' or 'Sticky MAC allow list'.
	AllowedVLANs                string                                                               `json:"allowedVlans,omitempty"`                // The VLANs allowed on the switch profile port. Only applicable to trunk ports.
	DaiTrusted                  *bool                                                                `json:"daiTrusted,omitempty"`                  // If true, ARP packets for this port will be considered trusted, and Dynamic ARP Inspection will allow the traffic.
	Enabled                     *bool                                                                `json:"enabled,omitempty"`                     // The status of the switch profile port.
	FlexibleStackingEnabled     *bool                                                                `json:"flexibleStackingEnabled,omitempty"`     // For supported switches (e.g. MS420/MS425), whether or not the port has flexible stacking enabled.
	IsolationEnabled            *bool                                                                `json:"isolationEnabled,omitempty"`            // The isolation status of the switch profile port.
	LinkNegotiation             string                                                               `json:"linkNegotiation,omitempty"`             // The link speed for the switch profile port.
	LinkNegotiationCapabilities []string                                                             `json:"linkNegotiationCapabilities,omitempty"` // Available link speeds for the switch profile port.
	MacAllowList                []string                                                             `json:"macAllowList,omitempty"`                // Only devices with MAC addresses specified in this list will have access to this port. Up to 20 MAC addresses can be defined. Only applicable when 'accessPolicyType' is 'MAC allow list'.
	Name                        string                                                               `json:"name,omitempty"`                        // The name of the switch profile port.
	PoeEnabled                  *bool                                                                `json:"poeEnabled,omitempty"`                  // The PoE status of the switch profile port.
	PortID                      string                                                               `json:"portId,omitempty"`                      // The identifier of the switch profile port.
	PortScheduleID              string                                                               `json:"portScheduleId,omitempty"`              // The ID of the port schedule. A value of null will clear the port schedule.
	Profile                     *ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortProfile `json:"profile,omitempty"`                     // Profile attributes
	RstpEnabled                 *bool                                                                `json:"rstpEnabled,omitempty"`                 // The rapid spanning tree protocol status.
	StickyMacAllowList          []string                                                             `json:"stickyMacAllowList,omitempty"`          // The initial list of MAC addresses for sticky Mac allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
	StickyMacAllowListLimit     *int                                                                 `json:"stickyMacAllowListLimit,omitempty"`     // The maximum number of MAC addresses for sticky MAC allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
	StormControlEnabled         *bool                                                                `json:"stormControlEnabled,omitempty"`         // The storm control status of the switch profile port.
	StpGuard                    string                                                               `json:"stpGuard,omitempty"`                    // The state of the STP guard ('disabled', 'root guard', 'bpdu guard' or 'loop guard').
	Tags                        []string                                                             `json:"tags,omitempty"`                        // The list of tags of the switch profile port.
	Type                        string                                                               `json:"type,omitempty"`                        // The type of the switch profile port ('trunk' or 'access').
	Udld                        string                                                               `json:"udld,omitempty"`                        // The action to take when Unidirectional Link is detected (Alert only, Enforce). Default configuration is Alert only.
	VLAN                        *int                                                                 `json:"vlan,omitempty"`                        // The VLAN of the switch profile port. A null value will clear the value set for trunk ports.
	VoiceVLAN                   *int                                                                 `json:"voiceVlan,omitempty"`                   // The voice VLAN of the switch profile port. Only applicable to access ports.
}
type ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortProfile struct {
	Enabled *bool  `json:"enabled,omitempty"` // When enabled, override this port's configuration with a port profile.
	ID      string `json:"id,omitempty"`      // When enabled, the ID of the port profile used to override the port's configuration.
	Iname   string `json:"iname,omitempty"`   // When enabled, the IName of the profile.
}
type ResponseSwitchUpdateOrganizationConfigTemplateSwitchProfilePort struct {
	AccessPolicyNumber          *int                                                                    `json:"accessPolicyNumber,omitempty"`          // The number of a custom access policy to configure on the switch profile port. Only applicable when 'accessPolicyType' is 'Custom access policy'.
	AccessPolicyType            string                                                                  `json:"accessPolicyType,omitempty"`            // The type of the access policy of the switch profile port. Only applicable to access ports. Can be one of 'Open', 'Custom access policy', 'MAC allow list' or 'Sticky MAC allow list'.
	AllowedVLANs                string                                                                  `json:"allowedVlans,omitempty"`                // The VLANs allowed on the switch profile port. Only applicable to trunk ports.
	DaiTrusted                  *bool                                                                   `json:"daiTrusted,omitempty"`                  // If true, ARP packets for this port will be considered trusted, and Dynamic ARP Inspection will allow the traffic.
	Enabled                     *bool                                                                   `json:"enabled,omitempty"`                     // The status of the switch profile port.
	FlexibleStackingEnabled     *bool                                                                   `json:"flexibleStackingEnabled,omitempty"`     // For supported switches (e.g. MS420/MS425), whether or not the port has flexible stacking enabled.
	IsolationEnabled            *bool                                                                   `json:"isolationEnabled,omitempty"`            // The isolation status of the switch profile port.
	LinkNegotiation             string                                                                  `json:"linkNegotiation,omitempty"`             // The link speed for the switch profile port.
	LinkNegotiationCapabilities []string                                                                `json:"linkNegotiationCapabilities,omitempty"` // Available link speeds for the switch profile port.
	MacAllowList                []string                                                                `json:"macAllowList,omitempty"`                // Only devices with MAC addresses specified in this list will have access to this port. Up to 20 MAC addresses can be defined. Only applicable when 'accessPolicyType' is 'MAC allow list'.
	Name                        string                                                                  `json:"name,omitempty"`                        // The name of the switch profile port.
	PoeEnabled                  *bool                                                                   `json:"poeEnabled,omitempty"`                  // The PoE status of the switch profile port.
	PortID                      string                                                                  `json:"portId,omitempty"`                      // The identifier of the switch profile port.
	PortScheduleID              string                                                                  `json:"portScheduleId,omitempty"`              // The ID of the port schedule. A value of null will clear the port schedule.
	Profile                     *ResponseSwitchUpdateOrganizationConfigTemplateSwitchProfilePortProfile `json:"profile,omitempty"`                     // Profile attributes
	RstpEnabled                 *bool                                                                   `json:"rstpEnabled,omitempty"`                 // The rapid spanning tree protocol status.
	StickyMacAllowList          []string                                                                `json:"stickyMacAllowList,omitempty"`          // The initial list of MAC addresses for sticky Mac allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
	StickyMacAllowListLimit     *int                                                                    `json:"stickyMacAllowListLimit,omitempty"`     // The maximum number of MAC addresses for sticky MAC allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
	StormControlEnabled         *bool                                                                   `json:"stormControlEnabled,omitempty"`         // The storm control status of the switch profile port.
	StpGuard                    string                                                                  `json:"stpGuard,omitempty"`                    // The state of the STP guard ('disabled', 'root guard', 'bpdu guard' or 'loop guard').
	Tags                        []string                                                                `json:"tags,omitempty"`                        // The list of tags of the switch profile port.
	Type                        string                                                                  `json:"type,omitempty"`                        // The type of the switch profile port ('trunk' or 'access').
	Udld                        string                                                                  `json:"udld,omitempty"`                        // The action to take when Unidirectional Link is detected (Alert only, Enforce). Default configuration is Alert only.
	VLAN                        *int                                                                    `json:"vlan,omitempty"`                        // The VLAN of the switch profile port. A null value will clear the value set for trunk ports.
	VoiceVLAN                   *int                                                                    `json:"voiceVlan,omitempty"`                   // The voice VLAN of the switch profile port. Only applicable to access ports.
}
type ResponseSwitchUpdateOrganizationConfigTemplateSwitchProfilePortProfile struct {
	Enabled *bool  `json:"enabled,omitempty"` // When enabled, override this port's configuration with a port profile.
	ID      string `json:"id,omitempty"`      // When enabled, the ID of the port profile used to override the port's configuration.
	Iname   string `json:"iname,omitempty"`   // When enabled, the IName of the profile.
}
type ResponseSwitchCloneOrganizationSwitchDevices interface{}
type ResponseSwitchGetOrganizationSwitchPortsBySwitch []ResponseItemSwitchGetOrganizationSwitchPortsBySwitch // Array of ResponseSwitchGetOrganizationSwitchPortsBySwitch
type ResponseItemSwitchGetOrganizationSwitchPortsBySwitch struct {
	Mac     string                                                       `json:"mac,omitempty"`     // The MAC address of the switch.
	Model   string                                                       `json:"model,omitempty"`   // The model of the switch.
	Name    string                                                       `json:"name,omitempty"`    // The name of the switch.
	Network *ResponseItemSwitchGetOrganizationSwitchPortsBySwitchNetwork `json:"network,omitempty"` // Identifying information of the switch's network.
	Ports   *[]ResponseItemSwitchGetOrganizationSwitchPortsBySwitchPorts `json:"ports,omitempty"`   // Ports belonging to the switch
	Serial  string                                                       `json:"serial,omitempty"`  // The serial number of the switch.
}
type ResponseItemSwitchGetOrganizationSwitchPortsBySwitchNetwork struct {
	ID   string `json:"id,omitempty"`   // The ID of the network.
	Name string `json:"name,omitempty"` // The name of the network.
}
type ResponseItemSwitchGetOrganizationSwitchPortsBySwitchPorts struct {
	AccessPolicyType        string   `json:"accessPolicyType,omitempty"`        // The type of the access policy of the switch port. Only applicable to access ports. Can be one of 'Open', 'Custom access policy', 'MAC allow list' or 'Sticky MAC allow list'.
	AllowedVLANs            string   `json:"allowedVlans,omitempty"`            // The VLANs allowed on the switch port. Only applicable to trunk ports.
	Enabled                 *bool    `json:"enabled,omitempty"`                 // The status of the switch port.
	LinkNegotiation         string   `json:"linkNegotiation,omitempty"`         // The link speed for the switch port.
	Name                    string   `json:"name,omitempty"`                    // The name of the switch port.
	PoeEnabled              *bool    `json:"poeEnabled,omitempty"`              // The PoE status of the switch port.
	PortID                  string   `json:"portId,omitempty"`                  // The identifier of the switch port.
	RstpEnabled             *bool    `json:"rstpEnabled,omitempty"`             // The rapid spanning tree protocol status.
	StickyMacAllowList      []string `json:"stickyMacAllowList,omitempty"`      // The initial list of MAC addresses for sticky Mac allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
	StickyMacAllowListLimit *int     `json:"stickyMacAllowListLimit,omitempty"` // The maximum number of MAC addresses for sticky MAC allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
	StpGuard                string   `json:"stpGuard,omitempty"`                // The state of the STP guard ('disabled', 'root guard', 'bpdu guard' or 'loop guard').
	Tags                    []string `json:"tags,omitempty"`                    // The list of tags of the switch port.
	Type                    string   `json:"type,omitempty"`                    // The type of the switch port ('trunk' or 'access').
	VLAN                    *int     `json:"vlan,omitempty"`                    // The VLAN of the switch port. A null value will clear the value set for trunk ports.
	VoiceVLAN               *int     `json:"voiceVlan,omitempty"`               // The voice VLAN of the switch port. Only applicable to access ports.
}
type RequestSwitchCycleDeviceSwitchPorts struct {
	Ports []string `json:"ports,omitempty"` // List of switch ports
}
type RequestSwitchUpdateDeviceSwitchPort struct {
	AccessPolicyNumber      *int                                        `json:"accessPolicyNumber,omitempty"`      // The number of a custom access policy to configure on the switch port. Only applicable when 'accessPolicyType' is 'Custom access policy'.
	AccessPolicyType        string                                      `json:"accessPolicyType,omitempty"`        // The type of the access policy of the switch port. Only applicable to access ports. Can be one of 'Open', 'Custom access policy', 'MAC allow list' or 'Sticky MAC allow list'.
	AdaptivePolicyGroupID   string                                      `json:"adaptivePolicyGroupId,omitempty"`   // The adaptive policy group ID that will be used to tag traffic through this switch port. This ID must pre-exist during the configuration, else needs to be created using adaptivePolicy/groups API. Cannot be applied to a port on a switch bound to profile.
	AllowedVLANs            string                                      `json:"allowedVlans,omitempty"`            // The VLANs allowed on the switch port. Only applicable to trunk ports.
	DaiTrusted              *bool                                       `json:"daiTrusted,omitempty"`              // If true, ARP packets for this port will be considered trusted, and Dynamic ARP Inspection will allow the traffic.
	Enabled                 *bool                                       `json:"enabled,omitempty"`                 // The status of the switch port.
	FlexibleStackingEnabled *bool                                       `json:"flexibleStackingEnabled,omitempty"` // For supported switches (e.g. MS420/MS425), whether or not the port has flexible stacking enabled.
	IsolationEnabled        *bool                                       `json:"isolationEnabled,omitempty"`        // The isolation status of the switch port.
	LinkNegotiation         string                                      `json:"linkNegotiation,omitempty"`         // The link speed for the switch port.
	MacAllowList            []string                                    `json:"macAllowList,omitempty"`            // Only devices with MAC addresses specified in this list will have access to this port. Up to 20 MAC addresses can be defined. Only applicable when 'accessPolicyType' is 'MAC allow list'.
	Name                    string                                      `json:"name,omitempty"`                    // The name of the switch port.
	PeerSgtCapable          *bool                                       `json:"peerSgtCapable,omitempty"`          // If true, Peer SGT is enabled for traffic through this switch port. Applicable to trunk port only, not access port. Cannot be applied to a port on a switch bound to profile.
	PoeEnabled              *bool                                       `json:"poeEnabled,omitempty"`              // The PoE status of the switch port.
	PortScheduleID          string                                      `json:"portScheduleId,omitempty"`          // The ID of the port schedule. A value of null will clear the port schedule.
	Profile                 *RequestSwitchUpdateDeviceSwitchPortProfile `json:"profile,omitempty"`                 // Profile attributes
	RstpEnabled             *bool                                       `json:"rstpEnabled,omitempty"`             // The rapid spanning tree protocol status.
	StickyMacAllowList      []string                                    `json:"stickyMacAllowList,omitempty"`      // The initial list of MAC addresses for sticky Mac allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
	StickyMacAllowListLimit *int                                        `json:"stickyMacAllowListLimit,omitempty"` // The maximum number of MAC addresses for sticky MAC allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
	StormControlEnabled     *bool                                       `json:"stormControlEnabled,omitempty"`     // The storm control status of the switch port.
	StpGuard                string                                      `json:"stpGuard,omitempty"`                // The state of the STP guard ('disabled', 'root guard', 'bpdu guard' or 'loop guard').
	Tags                    []string                                    `json:"tags,omitempty"`                    // The list of tags of the switch port.
	Type                    string                                      `json:"type,omitempty"`                    // The type of the switch port ('trunk' or 'access').
	Udld                    string                                      `json:"udld,omitempty"`                    // The action to take when Unidirectional Link is detected (Alert only, Enforce). Default configuration is Alert only.
	VLAN                    *int                                        `json:"vlan,omitempty"`                    // The VLAN of the switch port. A null value will clear the value set for trunk ports.
	VoiceVLAN               *int                                        `json:"voiceVlan,omitempty"`               // The voice VLAN of the switch port. Only applicable to access ports.
}
type RequestSwitchUpdateDeviceSwitchPortProfile struct {
	Enabled *bool  `json:"enabled,omitempty"` // When enabled, override this port's configuration with a port profile.
	ID      string `json:"id,omitempty"`      // When enabled, the ID of the port profile used to override the port's configuration.
	Iname   string `json:"iname,omitempty"`   // When enabled, the IName of the profile.
}
type RequestSwitchCreateDeviceSwitchRoutingInterface struct {
	DefaultGateway   string                                                       `json:"defaultGateway,omitempty"`   // The next hop for any traffic that isn't going to a directly connected subnet or over a static route.         This IP address must exist in a subnet with a routed interface. Required if this is the first IPv4 interface.
	InterfaceIP      string                                                       `json:"interfaceIp,omitempty"`      // The IP address this switch will use for layer 3 routing on this VLAN or subnet. This cannot be the same         as the switch's management IP.
	IPv6             *RequestSwitchCreateDeviceSwitchRoutingInterfaceIPv6         `json:"ipv6,omitempty"`             // The IPv6 settings of the interface.
	MulticastRouting string                                                       `json:"multicastRouting,omitempty"` // Enable multicast support if, multicast routing between VLANs is required. Options are:         'disabled', 'enabled' or 'IGMP snooping querier'. Default is 'disabled'.
	Name             string                                                       `json:"name,omitempty"`             // A friendly name or description for the interface or VLAN.
	OspfSettings     *RequestSwitchCreateDeviceSwitchRoutingInterfaceOspfSettings `json:"ospfSettings,omitempty"`     // The OSPF routing settings of the interface.
	OspfV3           *RequestSwitchCreateDeviceSwitchRoutingInterfaceOspfV3       `json:"ospfV3,omitempty"`           // The OSPFv3 routing settings of the interface.
	Subnet           string                                                       `json:"subnet,omitempty"`           // The network that this routed interface is on, in CIDR notation (ex. 10.1.1.0/24).
	VLANID           *int                                                         `json:"vlanId,omitempty"`           // The VLAN this routed interface is on. VLAN must be between 1 and 4094.
}
type RequestSwitchCreateDeviceSwitchRoutingInterfaceIPv6 struct {
	Address        string `json:"address,omitempty"`        // The IPv6 address of the interface. Required if assignmentMode is 'static'. Must not be included if           assignmentMode is 'eui-64'.
	AssignmentMode string `json:"assignmentMode,omitempty"` // The IPv6 assignment mode for the interface. Can be either 'eui-64' or 'static'.
	Gateway        string `json:"gateway,omitempty"`        // The IPv6 default gateway of the interface. Required if prefix is defined and this is the first           interface with IPv6 configured for the switch.
	Prefix         string `json:"prefix,omitempty"`         // The IPv6 prefix of the interface. Required if IPv6 object is included.
}
type RequestSwitchCreateDeviceSwitchRoutingInterfaceOspfSettings struct {
	Area             string `json:"area,omitempty"`             // The OSPF area to which this interface should belong. Can be either 'disabled' or the identifier of an           existing OSPF area. Defaults to 'disabled'.
	Cost             *int   `json:"cost,omitempty"`             // The path cost for this interface. Defaults to 1, but can be increased up to 65535           to give lower priority.
	IsPassiveEnabled *bool  `json:"isPassiveEnabled,omitempty"` // When enabled, OSPF will not run on the interface, but the subnet will still be advertised.
}
type RequestSwitchCreateDeviceSwitchRoutingInterfaceOspfV3 struct {
	Area             string `json:"area,omitempty"`             // The OSPFv3 area to which this interface should belong. Can be either 'disabled' or the identifier of an           existing OSPFv3 area. Defaults to 'disabled'.
	Cost             *int   `json:"cost,omitempty"`             // The path cost for this interface. Defaults to 1, but can be increased up to 65535           to give lower priority.
	IsPassiveEnabled *bool  `json:"isPassiveEnabled,omitempty"` // When enabled, OSPFv3 will not run on the interface, but the subnet will still be advertised.
}
type RequestSwitchUpdateDeviceSwitchRoutingInterface struct {
	DefaultGateway   string                                                       `json:"defaultGateway,omitempty"`   // The next hop for any traffic that isn't going to a directly connected subnet or over a static route.         This IP address must exist in a subnet with a routed interface. Required if this is the first IPv4 interface.
	InterfaceIP      string                                                       `json:"interfaceIp,omitempty"`      // The IP address this switch will use for layer 3 routing on this VLAN or subnet. This cannot be the same         as the switch's management IP.
	IPv6             *RequestSwitchUpdateDeviceSwitchRoutingInterfaceIPv6         `json:"ipv6,omitempty"`             // The IPv6 settings of the interface.
	MulticastRouting string                                                       `json:"multicastRouting,omitempty"` // Enable multicast support if, multicast routing between VLANs is required. Options are:         'disabled', 'enabled' or 'IGMP snooping querier'. Default is 'disabled'.
	Name             string                                                       `json:"name,omitempty"`             // A friendly name or description for the interface or VLAN.
	OspfSettings     *RequestSwitchUpdateDeviceSwitchRoutingInterfaceOspfSettings `json:"ospfSettings,omitempty"`     // The OSPF routing settings of the interface.
	OspfV3           *RequestSwitchUpdateDeviceSwitchRoutingInterfaceOspfV3       `json:"ospfV3,omitempty"`           // The OSPFv3 routing settings of the interface.
	Subnet           string                                                       `json:"subnet,omitempty"`           // The network that this routed interface is on, in CIDR notation (ex. 10.1.1.0/24).
	VLANID           *int                                                         `json:"vlanId,omitempty"`           // The VLAN this routed interface is on. VLAN must be between 1 and 4094.
}
type RequestSwitchUpdateDeviceSwitchRoutingInterfaceIPv6 struct {
	Address        string `json:"address,omitempty"`        // The IPv6 address of the interface. Required if assignmentMode is 'static'. Must not be included if           assignmentMode is 'eui-64'.
	AssignmentMode string `json:"assignmentMode,omitempty"` // The IPv6 assignment mode for the interface. Can be either 'eui-64' or 'static'.
	Gateway        string `json:"gateway,omitempty"`        // The IPv6 default gateway of the interface. Required if prefix is defined and this is the first           interface with IPv6 configured for the switch.
	Prefix         string `json:"prefix,omitempty"`         // The IPv6 prefix of the interface. Required if IPv6 object is included.
}
type RequestSwitchUpdateDeviceSwitchRoutingInterfaceOspfSettings struct {
	Area             string `json:"area,omitempty"`             // The OSPF area to which this interface should belong. Can be either 'disabled' or the identifier of an           existing OSPF area. Defaults to 'disabled'.
	Cost             *int   `json:"cost,omitempty"`             // The path cost for this interface. Defaults to 1, but can be increased up to 65535           to give lower priority.
	IsPassiveEnabled *bool  `json:"isPassiveEnabled,omitempty"` // When enabled, OSPF will not run on the interface, but the subnet will still be advertised.
}
type RequestSwitchUpdateDeviceSwitchRoutingInterfaceOspfV3 struct {
	Area             string `json:"area,omitempty"`             // The OSPFv3 area to which this interface should belong. Can be either 'disabled' or the identifier of an           existing OSPFv3 area. Defaults to 'disabled'.
	Cost             *int   `json:"cost,omitempty"`             // The path cost for this interface. Defaults to 1, but can be increased up to 65535           to give lower priority.
	IsPassiveEnabled *bool  `json:"isPassiveEnabled,omitempty"` // When enabled, OSPFv3 will not run on the interface, but the subnet will still be advertised.
}
type RequestSwitchUpdateDeviceSwitchRoutingInterfaceDhcp struct {
	BootFileName         string                                                                   `json:"bootFileName,omitempty"`         // The PXE boot server filename for the DHCP server running on the switch interface
	BootNextServer       string                                                                   `json:"bootNextServer,omitempty"`       // The PXE boot server IP for the DHCP server running on the switch interface
	BootOptionsEnabled   *bool                                                                    `json:"bootOptionsEnabled,omitempty"`   // Enable DHCP boot options to provide PXE boot options configs for the dhcp server running on the switch interface
	DhcpLeaseTime        string                                                                   `json:"dhcpLeaseTime,omitempty"`        // The DHCP lease time config for the dhcp server running on switch interface ('30 minutes', '1 hour', '4 hours', '12 hours', '1 day' or '1 week')
	DhcpMode             string                                                                   `json:"dhcpMode,omitempty"`             // The DHCP mode options for the switch interface ('dhcpDisabled', 'dhcpRelay' or 'dhcpServer')
	DhcpOptions          *[]RequestSwitchUpdateDeviceSwitchRoutingInterfaceDhcpDhcpOptions        `json:"dhcpOptions,omitempty"`          // Array of DHCP options consisting of code, type and value for the DHCP server running on the switch interface
	DhcpRelayServerIPs   []string                                                                 `json:"dhcpRelayServerIps,omitempty"`   // The DHCP relay server IPs to which DHCP packets would get relayed for the switch interface
	DNSCustomNameservers []string                                                                 `json:"dnsCustomNameservers,omitempty"` // The DHCP name server IPs when DHCP name server option is 'custom'
	DNSNameserversOption string                                                                   `json:"dnsNameserversOption,omitempty"` // The DHCP name server option for the dhcp server running on the switch interface ('googlePublicDns', 'openDns' or 'custom')
	FixedIPAssignments   *[]RequestSwitchUpdateDeviceSwitchRoutingInterfaceDhcpFixedIPAssignments `json:"fixedIpAssignments,omitempty"`   // Array of DHCP fixed IP assignments for the DHCP server running on the switch interface
	ReservedIPRanges     *[]RequestSwitchUpdateDeviceSwitchRoutingInterfaceDhcpReservedIPRanges   `json:"reservedIpRanges,omitempty"`     // Array of DHCP reserved IP assignments for the DHCP server running on the switch interface
}
type RequestSwitchUpdateDeviceSwitchRoutingInterfaceDhcpDhcpOptions struct {
	Code  string `json:"code,omitempty"`  // The code for DHCP option which should be from 2 to 254
	Type  string `json:"type,omitempty"`  // The type of the DHCP option which should be one of ('text', 'ip', 'integer' or 'hex')
	Value string `json:"value,omitempty"` // The value of the DHCP option
}
type RequestSwitchUpdateDeviceSwitchRoutingInterfaceDhcpFixedIPAssignments struct {
	IP   string `json:"ip,omitempty"`   // The IP address of the client which has fixed IP address assigned to it
	Mac  string `json:"mac,omitempty"`  // The MAC address of the client which has fixed IP address
	Name string `json:"name,omitempty"` // The name of the client which has fixed IP address
}
type RequestSwitchUpdateDeviceSwitchRoutingInterfaceDhcpReservedIPRanges struct {
	Comment string `json:"comment,omitempty"` // The comment for the reserved IP range
	End     string `json:"end,omitempty"`     // The ending IP address of the reserved IP range
	Start   string `json:"start,omitempty"`   // The starting IP address of the reserved IP range
}
type RequestSwitchCreateDeviceSwitchRoutingStaticRoute struct {
	AdvertiseViaOspfEnabled     *bool  `json:"advertiseViaOspfEnabled,omitempty"`     // Option to advertise static route via OSPF
	Name                        string `json:"name,omitempty"`                        // Name or description for layer 3 static route
	NextHopIP                   string `json:"nextHopIp,omitempty"`                   // IP address of the next hop device to which the device sends its traffic for the subnet
	PreferOverOspfRoutesEnabled *bool  `json:"preferOverOspfRoutesEnabled,omitempty"` // Option to prefer static route over OSPF routes
	Subnet                      string `json:"subnet,omitempty"`                      // The subnet which is routed via this static route and should be specified in CIDR notation (ex. 1.2.3.0/24)
}
type RequestSwitchUpdateDeviceSwitchRoutingStaticRoute struct {
	AdvertiseViaOspfEnabled     *bool  `json:"advertiseViaOspfEnabled,omitempty"`     // Option to advertise static route via OSPF
	Name                        string `json:"name,omitempty"`                        // Name or description for layer 3 static route
	NextHopIP                   string `json:"nextHopIp,omitempty"`                   // IP address of the next hop device to which the device sends its traffic for the subnet
	PreferOverOspfRoutesEnabled *bool  `json:"preferOverOspfRoutesEnabled,omitempty"` // Option to prefer static route over OSPF routes
	Subnet                      string `json:"subnet,omitempty"`                      // The subnet which is routed via this static route and should be specified in CIDR notation (ex. 1.2.3.0/24)
}
type RequestSwitchUpdateDeviceSwitchWarmSpare struct {
	Enabled     *bool  `json:"enabled,omitempty"`     // Enable or disable warm spare for a switch
	SpareSerial string `json:"spareSerial,omitempty"` // Serial number of the warm spare switch
}
type RequestSwitchUpdateNetworkSwitchAccessControlLists struct {
	Rules *[]RequestSwitchUpdateNetworkSwitchAccessControlListsRules `json:"rules,omitempty"` // An ordered array of the access control list rules (not including the default rule). An empty array will clear the rules.
}
type RequestSwitchUpdateNetworkSwitchAccessControlListsRules struct {
	Comment   string `json:"comment,omitempty"`   // Description of the rule (optional).
	DstCidr   string `json:"dstCidr,omitempty"`   // Destination IP address (in IP or CIDR notation) or 'any'.
	DstPort   string `json:"dstPort,omitempty"`   // Destination port. Must be in the range of 1-65535 or 'any'. Default is 'any'.
	IPVersion string `json:"ipVersion,omitempty"` // IP address version (must be 'any', 'ipv4' or 'ipv6'). Applicable only if network supports IPv6. Default value is 'ipv4'.
	Policy    string `json:"policy,omitempty"`    // 'allow' or 'deny' traffic specified by this rule.
	Protocol  string `json:"protocol,omitempty"`  // The type of protocol (must be 'tcp', 'udp', or 'any').
	SrcCidr   string `json:"srcCidr,omitempty"`   // Source IP address (in IP or CIDR notation) or 'any'.
	SrcPort   string `json:"srcPort,omitempty"`   // Source port. Must be in the range  of 1-65535 or 'any'. Default is 'any'.
	VLAN      string `json:"vlan,omitempty"`      // Incoming traffic VLAN. Must be in the range of 1-4095 or 'any'. Default is 'any'.
}
type RequestSwitchCreateNetworkSwitchAccessPolicy struct {
	AccessPolicyType               string                                                                 `json:"accessPolicyType,omitempty"`               // Access Type of the policy. Automatically 'Hybrid authentication' when hostMode is 'Multi-Domain'.
	Dot1X                          *RequestSwitchCreateNetworkSwitchAccessPolicyDot1X                     `json:"dot1x,omitempty"`                          // 802.1x Settings
	GuestPortBouncing              *bool                                                                  `json:"guestPortBouncing,omitempty"`              // If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers
	GuestVLANID                    *int                                                                   `json:"guestVlanId,omitempty"`                    // ID for the guest VLAN allow unauthorized devices access to limited network resources
	HostMode                       string                                                                 `json:"hostMode,omitempty"`                       // Choose the Host Mode for the access policy.
	IncreaseAccessSpeed            *bool                                                                  `json:"increaseAccessSpeed,omitempty"`            // Enabling this option will make switches execute 802.1X and MAC-bypass authentication simultaneously so that clients authenticate faster. Only required when accessPolicyType is 'Hybrid Authentication.
	Name                           string                                                                 `json:"name,omitempty"`                           // Name of the access policy
	Radius                         *RequestSwitchCreateNetworkSwitchAccessPolicyRadius                    `json:"radius,omitempty"`                         // Object for RADIUS Settings
	RadiusAccountingEnabled        *bool                                                                  `json:"radiusAccountingEnabled,omitempty"`        // Enable to send start, interim-update and stop messages to a configured RADIUS accounting server for tracking connected clients
	RadiusAccountingServers        *[]RequestSwitchCreateNetworkSwitchAccessPolicyRadiusAccountingServers `json:"radiusAccountingServers,omitempty"`        // List of RADIUS accounting servers to require connecting devices to authenticate against before granting network access
	RadiusCoaSupportEnabled        *bool                                                                  `json:"radiusCoaSupportEnabled,omitempty"`        // Change of authentication for RADIUS re-authentication and disconnection
	RadiusGroupAttribute           string                                                                 `json:"radiusGroupAttribute,omitempty"`           // Acceptable values are *""* for , or *"11"* for Group Policies ACL
	RadiusServers                  *[]RequestSwitchCreateNetworkSwitchAccessPolicyRadiusServers           `json:"radiusServers,omitempty"`                  // List of RADIUS servers to require connecting devices to authenticate against before granting network access
	RadiusTestingEnabled           *bool                                                                  `json:"radiusTestingEnabled,omitempty"`           // If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers
	URLRedirectWalledGardenEnabled *bool                                                                  `json:"urlRedirectWalledGardenEnabled,omitempty"` // Enable to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication
	URLRedirectWalledGardenRanges  []string                                                               `json:"urlRedirectWalledGardenRanges,omitempty"`  // IP address ranges, in CIDR notation, to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication
	VoiceVLANClients               *bool                                                                  `json:"voiceVlanClients,omitempty"`               // CDP/LLDP capable voice clients will be able to use this VLAN. Automatically true when hostMode is 'Multi-Domain'.
}
type RequestSwitchCreateNetworkSwitchAccessPolicyDot1X struct {
	ControlDirection string `json:"controlDirection,omitempty"` // Supports either 'both' or 'inbound'. Set to 'inbound' to allow unauthorized egress on the switchport. Set to 'both' to control both traffic directions with authorization. Defaults to 'both'
}
type RequestSwitchCreateNetworkSwitchAccessPolicyRadius struct {
	CriticalAuth             *RequestSwitchCreateNetworkSwitchAccessPolicyRadiusCriticalAuth `json:"criticalAuth,omitempty"`             // Critical auth settings for when authentication is rejected by the RADIUS server
	FailedAuthVLANID         *int                                                            `json:"failedAuthVlanId,omitempty"`         // VLAN that clients will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
	ReAuthenticationInterval *int                                                            `json:"reAuthenticationInterval,omitempty"` // Re-authentication period in seconds. Will be null if hostMode is Multi-Auth
}
type RequestSwitchCreateNetworkSwitchAccessPolicyRadiusCriticalAuth struct {
	DataVLANID        *int  `json:"dataVlanId,omitempty"`        // VLAN that clients who use data will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
	SuspendPortBounce *bool `json:"suspendPortBounce,omitempty"` // Enable to suspend port bounce when RADIUS servers are unreachable
	VoiceVLANID       *int  `json:"voiceVlanId,omitempty"`       // VLAN that clients who use voice will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
}
type RequestSwitchCreateNetworkSwitchAccessPolicyRadiusAccountingServers struct {
	Host   string `json:"host,omitempty"`   // Public IP address of the RADIUS accounting server
	Port   *int   `json:"port,omitempty"`   // UDP port that the RADIUS Accounting server listens on for access requests
	Secret string `json:"secret,omitempty"` // RADIUS client shared secret
}
type RequestSwitchCreateNetworkSwitchAccessPolicyRadiusServers struct {
	Host   string `json:"host,omitempty"`   // Public IP address of the RADIUS server
	Port   *int   `json:"port,omitempty"`   // UDP port that the RADIUS server listens on for access requests
	Secret string `json:"secret,omitempty"` // RADIUS client shared secret
}
type RequestSwitchUpdateNetworkSwitchAccessPolicy struct {
	AccessPolicyType               string                                                                 `json:"accessPolicyType,omitempty"`               // Access Type of the policy. Automatically 'Hybrid authentication' when hostMode is 'Multi-Domain'.
	Dot1X                          *RequestSwitchUpdateNetworkSwitchAccessPolicyDot1X                     `json:"dot1x,omitempty"`                          // 802.1x Settings
	GuestPortBouncing              *bool                                                                  `json:"guestPortBouncing,omitempty"`              // If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers
	GuestVLANID                    *int                                                                   `json:"guestVlanId,omitempty"`                    // ID for the guest VLAN allow unauthorized devices access to limited network resources
	HostMode                       string                                                                 `json:"hostMode,omitempty"`                       // Choose the Host Mode for the access policy.
	IncreaseAccessSpeed            *bool                                                                  `json:"increaseAccessSpeed,omitempty"`            // Enabling this option will make switches execute 802.1X and MAC-bypass authentication simultaneously so that clients authenticate faster. Only required when accessPolicyType is 'Hybrid Authentication.
	Name                           string                                                                 `json:"name,omitempty"`                           // Name of the access policy
	Radius                         *RequestSwitchUpdateNetworkSwitchAccessPolicyRadius                    `json:"radius,omitempty"`                         // Object for RADIUS Settings
	RadiusAccountingEnabled        *bool                                                                  `json:"radiusAccountingEnabled,omitempty"`        // Enable to send start, interim-update and stop messages to a configured RADIUS accounting server for tracking connected clients
	RadiusAccountingServers        *[]RequestSwitchUpdateNetworkSwitchAccessPolicyRadiusAccountingServers `json:"radiusAccountingServers,omitempty"`        // List of RADIUS accounting servers to require connecting devices to authenticate against before granting network access
	RadiusCoaSupportEnabled        *bool                                                                  `json:"radiusCoaSupportEnabled,omitempty"`        // Change of authentication for RADIUS re-authentication and disconnection
	RadiusGroupAttribute           string                                                                 `json:"radiusGroupAttribute,omitempty"`           // Acceptable values are *""* for , or *"11"* for Group Policies ACL
	RadiusServers                  *[]RequestSwitchUpdateNetworkSwitchAccessPolicyRadiusServers           `json:"radiusServers,omitempty"`                  // List of RADIUS servers to require connecting devices to authenticate against before granting network access
	RadiusTestingEnabled           *bool                                                                  `json:"radiusTestingEnabled,omitempty"`           // If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers
	URLRedirectWalledGardenEnabled *bool                                                                  `json:"urlRedirectWalledGardenEnabled,omitempty"` // Enable to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication
	URLRedirectWalledGardenRanges  []string                                                               `json:"urlRedirectWalledGardenRanges,omitempty"`  // IP address ranges, in CIDR notation, to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication
	VoiceVLANClients               *bool                                                                  `json:"voiceVlanClients,omitempty"`               // CDP/LLDP capable voice clients will be able to use this VLAN. Automatically true when hostMode is 'Multi-Domain'.
}
type RequestSwitchUpdateNetworkSwitchAccessPolicyDot1X struct {
	ControlDirection string `json:"controlDirection,omitempty"` // Supports either 'both' or 'inbound'. Set to 'inbound' to allow unauthorized egress on the switchport. Set to 'both' to control both traffic directions with authorization. Defaults to 'both'
}
type RequestSwitchUpdateNetworkSwitchAccessPolicyRadius struct {
	CriticalAuth             *RequestSwitchUpdateNetworkSwitchAccessPolicyRadiusCriticalAuth `json:"criticalAuth,omitempty"`             // Critical auth settings for when authentication is rejected by the RADIUS server
	FailedAuthVLANID         *int                                                            `json:"failedAuthVlanId,omitempty"`         // VLAN that clients will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
	ReAuthenticationInterval *int                                                            `json:"reAuthenticationInterval,omitempty"` // Re-authentication period in seconds. Will be null if hostMode is Multi-Auth
}
type RequestSwitchUpdateNetworkSwitchAccessPolicyRadiusCriticalAuth struct {
	DataVLANID        *int  `json:"dataVlanId,omitempty"`        // VLAN that clients who use data will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
	SuspendPortBounce *bool `json:"suspendPortBounce,omitempty"` // Enable to suspend port bounce when RADIUS servers are unreachable
	VoiceVLANID       *int  `json:"voiceVlanId,omitempty"`       // VLAN that clients who use voice will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
}
type RequestSwitchUpdateNetworkSwitchAccessPolicyRadiusAccountingServers struct {
	Host   string `json:"host,omitempty"`   // Public IP address of the RADIUS accounting server
	Port   *int   `json:"port,omitempty"`   // UDP port that the RADIUS Accounting server listens on for access requests
	Secret string `json:"secret,omitempty"` // RADIUS client shared secret
}
type RequestSwitchUpdateNetworkSwitchAccessPolicyRadiusServers struct {
	Host   string `json:"host,omitempty"`   // Public IP address of the RADIUS server
	Port   *int   `json:"port,omitempty"`   // UDP port that the RADIUS server listens on for access requests
	Secret string `json:"secret,omitempty"` // RADIUS client shared secret
}
type RequestSwitchUpdateNetworkSwitchAlternateManagementInterface struct {
	Enabled   *bool                                                                   `json:"enabled,omitempty"`   // Boolean value to enable or disable AMI configuration. If enabled, VLAN and protocols must be set
	Protocols []string                                                                `json:"protocols,omitempty"` // Can be one or more of the following values: 'radius', 'snmp' or 'syslog'
	Switches  *[]RequestSwitchUpdateNetworkSwitchAlternateManagementInterfaceSwitches `json:"switches,omitempty"`  // Array of switch serial number and IP assignment. If parameter is present, it cannot have empty body. Note: switches parameter is not applicable for template networks, in other words, do not put 'switches' in the body when updating template networks. Also, an empty 'switches' array will remove all previous assignments
	VLANID    *int                                                                    `json:"vlanId,omitempty"`    // Alternate management VLAN, must be between 1 and 4094
}
type RequestSwitchUpdateNetworkSwitchAlternateManagementInterfaceSwitches struct {
	AlternateManagementIP string `json:"alternateManagementIp,omitempty"` // Switch alternative management IP. To remove a prior IP setting, provide an empty string
	Gateway               string `json:"gateway,omitempty"`               // Switch gateway must be in IP format. Only and must be specified for Polaris switches
	Serial                string `json:"serial,omitempty"`                // Switch serial number
	SubnetMask            string `json:"subnetMask,omitempty"`            // Switch subnet mask must be in IP format. Only and must be specified for Polaris switches
}
type RequestSwitchUpdateNetworkSwitchDhcpServerPolicy struct {
	Alerts         *RequestSwitchUpdateNetworkSwitchDhcpServerPolicyAlerts        `json:"alerts,omitempty"`         // Alert settings for DHCP servers
	AllowedServers []string                                                       `json:"allowedServers,omitempty"` // List the MAC addresses of DHCP servers to permit on the network when defaultPolicy is set to block. An empty array will clear the entries.
	ArpInspection  *RequestSwitchUpdateNetworkSwitchDhcpServerPolicyArpInspection `json:"arpInspection,omitempty"`  // Dynamic ARP Inspection settings
	BlockedServers []string                                                       `json:"blockedServers,omitempty"` // List the MAC addresses of DHCP servers to block on the network when defaultPolicy is set to allow. An empty array will clear the entries.
	DefaultPolicy  string                                                         `json:"defaultPolicy,omitempty"`  // 'allow' or 'block' new DHCP servers. Default value is 'allow'.
}
type RequestSwitchUpdateNetworkSwitchDhcpServerPolicyAlerts struct {
	Email *RequestSwitchUpdateNetworkSwitchDhcpServerPolicyAlertsEmail `json:"email,omitempty"` // Email alert settings for DHCP servers
}
type RequestSwitchUpdateNetworkSwitchDhcpServerPolicyAlertsEmail struct {
	Enabled *bool `json:"enabled,omitempty"` // When enabled, send an email if a new DHCP server is seen. Default value is false.
}
type RequestSwitchUpdateNetworkSwitchDhcpServerPolicyArpInspection struct {
	Enabled *bool `json:"enabled,omitempty"` // Enable or disable Dynamic ARP Inspection on the network. Default value is false.
}
type RequestSwitchCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer struct {
	IPv4 *RequestSwitchCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerIPv4 `json:"ipv4,omitempty"` // The IPv4 attributes of the trusted server being added
	Mac  string                                                                          `json:"mac,omitempty"`  // The mac address of the trusted server being added
	VLAN *int                                                                            `json:"vlan,omitempty"` // The VLAN of the trusted server being added. It must be between 1 and 4094
}
type RequestSwitchCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerIPv4 struct {
	Address string `json:"address,omitempty"` // The IPv4 address of the trusted server being added
}
type RequestSwitchUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer struct {
	IPv4 *RequestSwitchUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerIPv4 `json:"ipv4,omitempty"` // The updated IPv4 attributes of the trusted server
	Mac  string                                                                          `json:"mac,omitempty"`  // The updated mac address of the trusted server
	VLAN *int                                                                            `json:"vlan,omitempty"` // The updated VLAN of the trusted server. It must be between 1 and 4094
}
type RequestSwitchUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerIPv4 struct {
	Address string `json:"address,omitempty"` // The updated IPv4 address of the trusted server
}
type RequestSwitchUpdateNetworkSwitchDscpToCosMappings struct {
	Mappings *[]RequestSwitchUpdateNetworkSwitchDscpToCosMappingsMappings `json:"mappings,omitempty"` // An array of DSCP to CoS mappings. An empty array will reset the mappings to default.
}
type RequestSwitchUpdateNetworkSwitchDscpToCosMappingsMappings struct {
	Cos   *int   `json:"cos,omitempty"`   // The actual layer-2 CoS queue the DSCP value is mapped to. These are not bits set on outgoing frames. Value can be in the range of 0 to 5 inclusive.
	Dscp  *int   `json:"dscp,omitempty"`  // The Differentiated Services Code Point (DSCP) tag in the IP header that will be mapped to a particular Class-of-Service (CoS) queue. Value can be in the range of 0 to 63 inclusive.
	Title string `json:"title,omitempty"` // Label for the mapping (optional).
}
type RequestSwitchCreateNetworkSwitchLinkAggregation struct {
	SwitchPorts        *[]RequestSwitchCreateNetworkSwitchLinkAggregationSwitchPorts        `json:"switchPorts,omitempty"`        // Array of switch or stack ports for creating aggregation group. Minimum 2 and maximum 8 ports are supported.
	SwitchProfilePorts *[]RequestSwitchCreateNetworkSwitchLinkAggregationSwitchProfilePorts `json:"switchProfilePorts,omitempty"` // Array of switch profile ports for creating aggregation group. Minimum 2 and maximum 8 ports are supported.
}
type RequestSwitchCreateNetworkSwitchLinkAggregationSwitchPorts struct {
	PortID string `json:"portId,omitempty"` // Port identifier of switch port. For modules, the identifier is "SlotNumber_ModuleType_PortNumber" (Ex: "1_8X10G_1"), otherwise it is just the port number (Ex: "8").
	Serial string `json:"serial,omitempty"` // Serial number of the switch.
}
type RequestSwitchCreateNetworkSwitchLinkAggregationSwitchProfilePorts struct {
	PortID  string `json:"portId,omitempty"`  // Port identifier of switch port. For modules, the identifier is "SlotNumber_ModuleType_PortNumber" (Ex: "1_8X10G_1"), otherwise it is just the port number (Ex: "8").
	Profile string `json:"profile,omitempty"` // Profile identifier.
}
type RequestSwitchUpdateNetworkSwitchLinkAggregation struct {
	SwitchPorts        *[]RequestSwitchUpdateNetworkSwitchLinkAggregationSwitchPorts        `json:"switchPorts,omitempty"`        // Array of switch or stack ports for updating aggregation group. Minimum 2 and maximum 8 ports are supported.
	SwitchProfilePorts *[]RequestSwitchUpdateNetworkSwitchLinkAggregationSwitchProfilePorts `json:"switchProfilePorts,omitempty"` // Array of switch profile ports for updating aggregation group. Minimum 2 and maximum 8 ports are supported.
}
type RequestSwitchUpdateNetworkSwitchLinkAggregationSwitchPorts struct {
	PortID string `json:"portId,omitempty"` // Port identifier of switch port. For modules, the identifier is "SlotNumber_ModuleType_PortNumber" (Ex: "1_8X10G_1"), otherwise it is just the port number (Ex: "8").
	Serial string `json:"serial,omitempty"` // Serial number of the switch.
}
type RequestSwitchUpdateNetworkSwitchLinkAggregationSwitchProfilePorts struct {
	PortID  string `json:"portId,omitempty"`  // Port identifier of switch port. For modules, the identifier is "SlotNumber_ModuleType_PortNumber" (Ex: "1_8X10G_1"), otherwise it is just the port number (Ex: "8").
	Profile string `json:"profile,omitempty"` // Profile identifier.
}
type RequestSwitchUpdateNetworkSwitchMtu struct {
	DefaultMtuSize *int                                            `json:"defaultMtuSize,omitempty"` // MTU size for the entire network. Default value is 9578.
	Overrides      *[]RequestSwitchUpdateNetworkSwitchMtuOverrides `json:"overrides,omitempty"`      // Override MTU size for individual switches or switch profiles. An empty array will clear overrides.
}
type RequestSwitchUpdateNetworkSwitchMtuOverrides struct {
	MtuSize        *int     `json:"mtuSize,omitempty"`        // MTU size for the switches or switch profiles.
	SwitchProfiles []string `json:"switchProfiles,omitempty"` // List of switch profile IDs. Applicable only for template network.
	Switches       []string `json:"switches,omitempty"`       // List of switch serials. Applicable only for switch network.
}
type RequestSwitchCreateNetworkSwitchPortSchedule struct {
	Name         string                                                    `json:"name,omitempty"`         // The name for your port schedule. Required
	PortSchedule *RequestSwitchCreateNetworkSwitchPortSchedulePortSchedule `json:"portSchedule,omitempty"` //     The schedule for switch port scheduling. Schedules are applied to days of the week.     When it's empty, default schedule with all days of a week are configured.     Any unspecified day in the schedule is added as a default schedule configuration of the day.
}
type RequestSwitchCreateNetworkSwitchPortSchedulePortSchedule struct {
	Friday    *RequestSwitchCreateNetworkSwitchPortSchedulePortScheduleFriday    `json:"friday,omitempty"`    // The schedule object for Friday.
	Monday    *RequestSwitchCreateNetworkSwitchPortSchedulePortScheduleMonday    `json:"monday,omitempty"`    // The schedule object for Monday.
	Saturday  *RequestSwitchCreateNetworkSwitchPortSchedulePortScheduleSaturday  `json:"saturday,omitempty"`  // The schedule object for Saturday.
	Sunday    *RequestSwitchCreateNetworkSwitchPortSchedulePortScheduleSunday    `json:"sunday,omitempty"`    // The schedule object for Sunday.
	Thursday  *RequestSwitchCreateNetworkSwitchPortSchedulePortScheduleThursday  `json:"thursday,omitempty"`  // The schedule object for Thursday.
	Tuesday   *RequestSwitchCreateNetworkSwitchPortSchedulePortScheduleTuesday   `json:"tuesday,omitempty"`   // The schedule object for Tuesday.
	Wednesday *RequestSwitchCreateNetworkSwitchPortSchedulePortScheduleWednesday `json:"wednesday,omitempty"` // The schedule object for Wednesday.
}
type RequestSwitchCreateNetworkSwitchPortSchedulePortScheduleFriday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
}
type RequestSwitchCreateNetworkSwitchPortSchedulePortScheduleMonday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
}
type RequestSwitchCreateNetworkSwitchPortSchedulePortScheduleSaturday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
}
type RequestSwitchCreateNetworkSwitchPortSchedulePortScheduleSunday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
}
type RequestSwitchCreateNetworkSwitchPortSchedulePortScheduleThursday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
}
type RequestSwitchCreateNetworkSwitchPortSchedulePortScheduleTuesday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
}
type RequestSwitchCreateNetworkSwitchPortSchedulePortScheduleWednesday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
}
type RequestSwitchUpdateNetworkSwitchPortSchedule struct {
	Name         string                                                    `json:"name,omitempty"`         // The name for your port schedule.
	PortSchedule *RequestSwitchUpdateNetworkSwitchPortSchedulePortSchedule `json:"portSchedule,omitempty"` //     The schedule for switch port scheduling. Schedules are applied to days of the week.     When it's empty, default schedule with all days of a week are configured.     Any unspecified day in the schedule is added as a default schedule configuration of the day.
}
type RequestSwitchUpdateNetworkSwitchPortSchedulePortSchedule struct {
	Friday    *RequestSwitchUpdateNetworkSwitchPortSchedulePortScheduleFriday    `json:"friday,omitempty"`    // The schedule object for Friday.
	Monday    *RequestSwitchUpdateNetworkSwitchPortSchedulePortScheduleMonday    `json:"monday,omitempty"`    // The schedule object for Monday.
	Saturday  *RequestSwitchUpdateNetworkSwitchPortSchedulePortScheduleSaturday  `json:"saturday,omitempty"`  // The schedule object for Saturday.
	Sunday    *RequestSwitchUpdateNetworkSwitchPortSchedulePortScheduleSunday    `json:"sunday,omitempty"`    // The schedule object for Sunday.
	Thursday  *RequestSwitchUpdateNetworkSwitchPortSchedulePortScheduleThursday  `json:"thursday,omitempty"`  // The schedule object for Thursday.
	Tuesday   *RequestSwitchUpdateNetworkSwitchPortSchedulePortScheduleTuesday   `json:"tuesday,omitempty"`   // The schedule object for Tuesday.
	Wednesday *RequestSwitchUpdateNetworkSwitchPortSchedulePortScheduleWednesday `json:"wednesday,omitempty"` // The schedule object for Wednesday.
}
type RequestSwitchUpdateNetworkSwitchPortSchedulePortScheduleFriday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
}
type RequestSwitchUpdateNetworkSwitchPortSchedulePortScheduleMonday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
}
type RequestSwitchUpdateNetworkSwitchPortSchedulePortScheduleSaturday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
}
type RequestSwitchUpdateNetworkSwitchPortSchedulePortScheduleSunday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
}
type RequestSwitchUpdateNetworkSwitchPortSchedulePortScheduleThursday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
}
type RequestSwitchUpdateNetworkSwitchPortSchedulePortScheduleTuesday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
}
type RequestSwitchUpdateNetworkSwitchPortSchedulePortScheduleWednesday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
}
type RequestSwitchCreateNetworkSwitchQosRule struct {
	Dscp         *int   `json:"dscp,omitempty"`         // DSCP tag. Set this to -1 to trust incoming DSCP. Default value is 0
	DstPort      *int   `json:"dstPort,omitempty"`      // The destination port of the incoming packet. Applicable only if protocol is TCP or UDP.
	DstPortRange string `json:"dstPortRange,omitempty"` // The destination port range of the incoming packet. Applicable only if protocol is set to TCP or UDP. Example: 70-80
	Protocol     string `json:"protocol,omitempty"`     // The protocol of the incoming packet. Can be one of "ANY", "TCP" or "UDP". Default value is "ANY"
	SrcPort      *int   `json:"srcPort,omitempty"`      // The source port of the incoming packet. Applicable only if protocol is TCP or UDP.
	SrcPortRange string `json:"srcPortRange,omitempty"` // The source port range of the incoming packet. Applicable only if protocol is set to TCP or UDP. Example: 70-80
	VLAN         *int   `json:"vlan,omitempty"`         // The VLAN of the incoming packet. A null value will match any VLAN.
}
type RequestSwitchUpdateNetworkSwitchQosRulesOrder struct {
	RuleIDs []string `json:"ruleIds,omitempty"` // A list of quality of service rule IDs arranged in order in which they should be processed by the switch.
}
type RequestSwitchUpdateNetworkSwitchQosRule struct {
	Dscp         *int   `json:"dscp,omitempty"`         // DSCP tag that should be assigned to incoming packet. Set this to -1 to trust incoming DSCP. Default value is 0.
	DstPort      *int   `json:"dstPort,omitempty"`      // The destination port of the incoming packet. Applicable only if protocol is TCP or UDP.
	DstPortRange string `json:"dstPortRange,omitempty"` // The destination port range of the incoming packet. Applicable only if protocol is set to TCP or UDP. Example: 70-80
	Protocol     string `json:"protocol,omitempty"`     // The protocol of the incoming packet. Can be one of "ANY", "TCP" or "UDP". Default value is "ANY".
	SrcPort      *int   `json:"srcPort,omitempty"`      // The source port of the incoming packet. Applicable only if protocol is TCP or UDP.
	SrcPortRange string `json:"srcPortRange,omitempty"` // The source port range of the incoming packet. Applicable only if protocol is set to TCP or UDP. Example: 70-80
	VLAN         *int   `json:"vlan,omitempty"`         // The VLAN of the incoming packet. A null value will match any VLAN.
}
type RequestSwitchUpdateNetworkSwitchRoutingMulticast struct {
	DefaultSettings *RequestSwitchUpdateNetworkSwitchRoutingMulticastDefaultSettings `json:"defaultSettings,omitempty"` // Default multicast setting for entire network. IGMP snooping and Flood unknown multicast traffic settings are enabled by default.
	Overrides       *[]RequestSwitchUpdateNetworkSwitchRoutingMulticastOverrides     `json:"overrides,omitempty"`       // Array of paired switches/stacks/profiles and corresponding multicast settings. An empty array will clear the multicast settings.
}
type RequestSwitchUpdateNetworkSwitchRoutingMulticastDefaultSettings struct {
	FloodUnknownMulticastTrafficEnabled *bool `json:"floodUnknownMulticastTrafficEnabled,omitempty"` // Flood unknown multicast traffic setting for entire network
	IgmpSnoopingEnabled                 *bool `json:"igmpSnoopingEnabled,omitempty"`                 // IGMP snooping setting for entire network
}
type RequestSwitchUpdateNetworkSwitchRoutingMulticastOverrides struct {
	FloodUnknownMulticastTrafficEnabled *bool    `json:"floodUnknownMulticastTrafficEnabled,omitempty"` // Flood unknown multicast traffic setting for switches, switch stacks or switch profiles
	IgmpSnoopingEnabled                 *bool    `json:"igmpSnoopingEnabled,omitempty"`                 // IGMP snooping setting for switches, switch stacks or switch profiles
	Stacks                              []string `json:"stacks,omitempty"`                              // List of switch stack ids for non-template network
	SwitchProfiles                      []string `json:"switchProfiles,omitempty"`                      // List of switch profiles ids for template network
	Switches                            []string `json:"switches,omitempty"`                            // List of switch serials for non-template network
}
type RequestSwitchCreateNetworkSwitchRoutingMulticastRendezvousPoint struct {
	InterfaceIP    string `json:"interfaceIp,omitempty"`    // The IP address of the interface where the RP needs to be created.
	MulticastGroup string `json:"multicastGroup,omitempty"` // 'Any', or the IP address of a multicast group
}
type RequestSwitchUpdateNetworkSwitchRoutingMulticastRendezvousPoint struct {
	InterfaceIP    string `json:"interfaceIp,omitempty"`    // The IP address of the interface to use
	MulticastGroup string `json:"multicastGroup,omitempty"` // 'Any', or the IP address of a multicast group
}
type RequestSwitchUpdateNetworkSwitchRoutingOspf struct {
	Areas                    *[]RequestSwitchUpdateNetworkSwitchRoutingOspfAreas              `json:"areas,omitempty"`                    // OSPF areas
	DeadTimerInSeconds       *int                                                             `json:"deadTimerInSeconds,omitempty"`       // Time interval to determine when the peer will be declared inactive/dead. Value must be between 1 and 65535
	Enabled                  *bool                                                            `json:"enabled,omitempty"`                  // Boolean value to enable or disable OSPF routing. OSPF routing is disabled by default.
	HelloTimerInSeconds      *int                                                             `json:"helloTimerInSeconds,omitempty"`      // Time interval in seconds at which hello packet will be sent to OSPF neighbors to maintain connectivity. Value must be between 1 and 255. Default is 10 seconds.
	Md5AuthenticationEnabled *bool                                                            `json:"md5AuthenticationEnabled,omitempty"` // Boolean value to enable or disable MD5 authentication. MD5 authentication is disabled by default.
	Md5AuthenticationKey     *RequestSwitchUpdateNetworkSwitchRoutingOspfMd5AuthenticationKey `json:"md5AuthenticationKey,omitempty"`     // MD5 authentication credentials. This param is only relevant if md5AuthenticationEnabled is true
	V3                       *RequestSwitchUpdateNetworkSwitchRoutingOspfV3                   `json:"v3,omitempty"`                       // OSPF v3 configuration
}
type RequestSwitchUpdateNetworkSwitchRoutingOspfAreas struct {
	AreaID   string `json:"areaId,omitempty"`   // OSPF area ID
	AreaName string `json:"areaName,omitempty"` // Name of the OSPF area
	AreaType string `json:"areaType,omitempty"` // Area types in OSPF. Must be one of: ["normal", "stub", "nssa"]
}
type RequestSwitchUpdateNetworkSwitchRoutingOspfMd5AuthenticationKey struct {
	ID         *int   `json:"id,omitempty"`         // MD5 authentication key index. Key index must be between 1 to 255
	Passphrase string `json:"passphrase,omitempty"` // MD5 authentication passphrase
}
type RequestSwitchUpdateNetworkSwitchRoutingOspfV3 struct {
	Areas               *[]RequestSwitchUpdateNetworkSwitchRoutingOspfV3Areas `json:"areas,omitempty"`               // OSPF v3 areas
	DeadTimerInSeconds  *int                                                  `json:"deadTimerInSeconds,omitempty"`  // Time interval to determine when the peer will be declared inactive/dead. Value must be between 1 and 65535
	Enabled             *bool                                                 `json:"enabled,omitempty"`             // Boolean value to enable or disable V3 OSPF routing. OSPF V3 routing is disabled by default.
	HelloTimerInSeconds *int                                                  `json:"helloTimerInSeconds,omitempty"` // Time interval in seconds at which hello packet will be sent to OSPF neighbors to maintain connectivity. Value must be between 1 and 255. Default is 10 seconds.
}
type RequestSwitchUpdateNetworkSwitchRoutingOspfV3Areas struct {
	AreaID   string `json:"areaId,omitempty"`   // OSPF area ID
	AreaName string `json:"areaName,omitempty"` // Name of the OSPF area
	AreaType string `json:"areaType,omitempty"` // Area types in OSPF. Must be one of: ["normal", "stub", "nssa"]
}
type RequestSwitchUpdateNetworkSwitchSettings struct {
	PowerExceptions  *[]RequestSwitchUpdateNetworkSwitchSettingsPowerExceptions `json:"powerExceptions,omitempty"`  // Exceptions on a per switch basis to "useCombinedPower"
	UseCombinedPower *bool                                                      `json:"useCombinedPower,omitempty"` // The use Combined Power as the default behavior of secondary power supplies on supported devices.
	VLAN             *int                                                       `json:"vlan,omitempty"`             // Management VLAN
}
type RequestSwitchUpdateNetworkSwitchSettingsPowerExceptions struct {
	PowerType string `json:"powerType,omitempty"` // Per switch exception (combined, redundant, useNetworkSetting)
	Serial    string `json:"serial,omitempty"`    // Serial number of the switch
}
type RequestSwitchCreateNetworkSwitchStack struct {
	Name    string   `json:"name,omitempty"`    // The name of the new stack
	Serials []string `json:"serials,omitempty"` // An array of switch serials to be added into the new stack
}
type RequestSwitchAddNetworkSwitchStack struct {
	Serial string `json:"serial,omitempty"` // The serial of the switch to be added
}
type RequestSwitchRemoveNetworkSwitchStack struct {
	Serial string `json:"serial,omitempty"` // The serial of the switch to be removed
}
type RequestSwitchCreateNetworkSwitchStackRoutingInterface struct {
	DefaultGateway   string                                                             `json:"defaultGateway,omitempty"`   // The next hop for any traffic that isn't going to a directly connected subnet or over a static route. This IP address must exist in a subnet with a routed interface.
	InterfaceIP      string                                                             `json:"interfaceIp,omitempty"`      // The IP address this switch stack will use for layer 3 routing on this VLAN or subnet. This cannot be the same as the switch's management IP.
	IPv6             *RequestSwitchCreateNetworkSwitchStackRoutingInterfaceIPv6         `json:"ipv6,omitempty"`             // The IPv6 settings of the interface.
	MulticastRouting string                                                             `json:"multicastRouting,omitempty"` // Enable multicast support if, multicast routing between VLANs is required. Options are, 'disabled', 'enabled' or 'IGMP snooping querier'. Default is 'disabled'.
	Name             string                                                             `json:"name,omitempty"`             // A friendly name or description for the interface or VLAN.
	OspfSettings     *RequestSwitchCreateNetworkSwitchStackRoutingInterfaceOspfSettings `json:"ospfSettings,omitempty"`     // The OSPF routing settings of the interface.
	Subnet           string                                                             `json:"subnet,omitempty"`           // The network that this routed interface is on, in CIDR notation (ex. 10.1.1.0/24).
	VLANID           *int                                                               `json:"vlanId,omitempty"`           // The VLAN this routed interface is on. VLAN must be between 1 and 4094.
}
type RequestSwitchCreateNetworkSwitchStackRoutingInterfaceIPv6 struct {
	Address        string `json:"address,omitempty"`        // The IPv6 address of the interface. Required if assignmentMode is 'static'. Must not be included if assignmentMode is 'eui-64'.
	AssignmentMode string `json:"assignmentMode,omitempty"` // The IPv6 assignment mode for the interface. Can be either 'eui-64' or 'static'.
	Gateway        string `json:"gateway,omitempty"`        // The IPv6 default gateway of the interface. Required if prefix is defined and this is the first interface with IPv6 configured for the stack.
	Prefix         string `json:"prefix,omitempty"`         // The IPv6 prefix of the interface. Required if IPv6 object is included.
}
type RequestSwitchCreateNetworkSwitchStackRoutingInterfaceOspfSettings struct {
	Area             string `json:"area,omitempty"`             // The OSPF area to which this interface should belong. Can be either 'disabled' or the identifier of an existing OSPF area. Defaults to 'disabled'.
	Cost             *int   `json:"cost,omitempty"`             // The path cost for this interface. Defaults to 1, but can be increased up to 65535 to give lower priority.
	IsPassiveEnabled *bool  `json:"isPassiveEnabled,omitempty"` // When enabled, OSPF will not run on the interface, but the subnet will still be advertised.
}
type RequestSwitchUpdateNetworkSwitchStackRoutingInterface struct {
	DefaultGateway   string                                                             `json:"defaultGateway,omitempty"`   // The next hop for any traffic that isn't going to a directly connected subnet or over a static route. This IP address must exist in a subnet with a routed interface.
	InterfaceIP      string                                                             `json:"interfaceIp,omitempty"`      // The IP address this switch stack will use for layer 3 routing on this VLAN or subnet. This cannot be the same as the switch's management IP.
	IPv6             *RequestSwitchUpdateNetworkSwitchStackRoutingInterfaceIPv6         `json:"ipv6,omitempty"`             // The IPv6 settings of the interface.
	MulticastRouting string                                                             `json:"multicastRouting,omitempty"` // Enable multicast support if, multicast routing between VLANs is required. Options are, 'disabled', 'enabled' or 'IGMP snooping querier'.
	Name             string                                                             `json:"name,omitempty"`             // A friendly name or description for the interface or VLAN.
	OspfSettings     *RequestSwitchUpdateNetworkSwitchStackRoutingInterfaceOspfSettings `json:"ospfSettings,omitempty"`     // The OSPF routing settings of the interface.
	Subnet           string                                                             `json:"subnet,omitempty"`           // The network that this routed interface is on, in CIDR notation (ex. 10.1.1.0/24).
	VLANID           *int                                                               `json:"vlanId,omitempty"`           // The VLAN this routed interface is on. VLAN must be between 1 and 4094.
}
type RequestSwitchUpdateNetworkSwitchStackRoutingInterfaceIPv6 struct {
	Address        string `json:"address,omitempty"`        // The IPv6 address of the interface. Required if assignmentMode is included and set as 'static'. Must not be included if assignmentMode is 'eui-64'.
	AssignmentMode string `json:"assignmentMode,omitempty"` // The IPv6 assignment mode for the interface. Can be either 'eui-64' or 'static'.
	Gateway        string `json:"gateway,omitempty"`        // The IPv6 default gateway of the interface. Required if prefix is defined and this is the first interface with IPv6 configured for the stack.
	Prefix         string `json:"prefix,omitempty"`         // The IPv6 prefix of the interface. Required if IPv6 object is included and interface does not already have ipv6.prefix configured
}
type RequestSwitchUpdateNetworkSwitchStackRoutingInterfaceOspfSettings struct {
	Area             string `json:"area,omitempty"`             // The OSPF area to which this interface should belong. Can be either 'disabled' or the identifier of an existing OSPF area.
	Cost             *int   `json:"cost,omitempty"`             // The path cost for this interface. Defaults to 1, but can be increased up to 65535 to give lower priority.
	IsPassiveEnabled *bool  `json:"isPassiveEnabled,omitempty"` // When enabled, OSPF will not run on the interface, but the subnet will still be advertised.
}
type RequestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcp struct {
	BootFileName         string                                                                         `json:"bootFileName,omitempty"`         // The PXE boot server file name for the DHCP server running on the switch stack interface
	BootNextServer       string                                                                         `json:"bootNextServer,omitempty"`       // The PXE boot server IP for the DHCP server running on the switch stack interface
	BootOptionsEnabled   *bool                                                                          `json:"bootOptionsEnabled,omitempty"`   // Enable DHCP boot options to provide PXE boot options configs for the dhcp server running on the switch stack interface
	DhcpLeaseTime        string                                                                         `json:"dhcpLeaseTime,omitempty"`        // The DHCP lease time config for the dhcp server running on switch stack interface ('30 minutes', '1 hour', '4 hours', '12 hours', '1 day' or '1 week')
	DhcpMode             string                                                                         `json:"dhcpMode,omitempty"`             // The DHCP mode options for the switch stack interface ('dhcpDisabled', 'dhcpRelay' or 'dhcpServer')
	DhcpOptions          *[]RequestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcpDhcpOptions        `json:"dhcpOptions,omitempty"`          // Array of DHCP options consisting of code, type and value for the DHCP server running on the switch stack interface
	DhcpRelayServerIPs   []string                                                                       `json:"dhcpRelayServerIps,omitempty"`   // The DHCP relay server IPs to which DHCP packets would get relayed for the switch stack interface
	DNSCustomNameservers []string                                                                       `json:"dnsCustomNameservers,omitempty"` // The DHCP name server IPs when DHCP name server option is 'custom'
	DNSNameserversOption string                                                                         `json:"dnsNameserversOption,omitempty"` // The DHCP name server option for the dhcp server running on the switch stack interface ('googlePublicDns', 'openDns' or 'custom')
	FixedIPAssignments   *[]RequestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcpFixedIPAssignments `json:"fixedIpAssignments,omitempty"`   // Array of DHCP fixed IP assignments for the DHCP server running on the switch stack interface
	ReservedIPRanges     *[]RequestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcpReservedIPRanges   `json:"reservedIpRanges,omitempty"`     // Array of DHCP reserved IP assignments for the DHCP server running on the switch stack interface
}
type RequestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcpDhcpOptions struct {
	Code  string `json:"code,omitempty"`  // The code for DHCP option which should be from 2 to 254
	Type  string `json:"type,omitempty"`  // The type of the DHCP option which should be one of ('text', 'ip', 'integer' or 'hex')
	Value string `json:"value,omitempty"` // The value of the DHCP option
}
type RequestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcpFixedIPAssignments struct {
	IP   string `json:"ip,omitempty"`   // The IP address of the client which has fixed IP address assigned to it
	Mac  string `json:"mac,omitempty"`  // The MAC address of the client which has fixed IP address
	Name string `json:"name,omitempty"` // The name of the client which has fixed IP address
}
type RequestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcpReservedIPRanges struct {
	Comment string `json:"comment,omitempty"` // The comment for the reserved IP range
	End     string `json:"end,omitempty"`     // The ending IP address of the reserved IP range
	Start   string `json:"start,omitempty"`   // The starting IP address of the reserved IP range
}
type RequestSwitchCreateNetworkSwitchStackRoutingStaticRoute struct {
	AdvertiseViaOspfEnabled     *bool  `json:"advertiseViaOspfEnabled,omitempty"`     // Option to advertise static route via OSPF
	Name                        string `json:"name,omitempty"`                        // Name or description for layer 3 static route
	NextHopIP                   string `json:"nextHopIp,omitempty"`                   // IP address of the next hop device to which the device sends its traffic for the subnet
	PreferOverOspfRoutesEnabled *bool  `json:"preferOverOspfRoutesEnabled,omitempty"` // Option to prefer static route over OSPF routes
	Subnet                      string `json:"subnet,omitempty"`                      // The subnet which is routed via this static route and should be specified in CIDR notation (ex. 1.2.3.0/24)
}
type RequestSwitchUpdateNetworkSwitchStackRoutingStaticRoute struct {
	AdvertiseViaOspfEnabled     *bool  `json:"advertiseViaOspfEnabled,omitempty"`     // Option to advertise static route via OSPF
	Name                        string `json:"name,omitempty"`                        // Name or description for layer 3 static route
	NextHopIP                   string `json:"nextHopIp,omitempty"`                   // IP address of the next hop device to which the device sends its traffic for the subnet
	PreferOverOspfRoutesEnabled *bool  `json:"preferOverOspfRoutesEnabled,omitempty"` // Option to prefer static route over OSPF routes
	Subnet                      string `json:"subnet,omitempty"`                      // The subnet which is routed via this static route and should be specified in CIDR notation (ex. 1.2.3.0/24)
}
type RequestSwitchUpdateNetworkSwitchStormControl struct {
	BroadcastThreshold      *int `json:"broadcastThreshold,omitempty"`      // Percentage (1 to 99) of total available port bandwidth for broadcast traffic type. Default value 100 percent rate is to clear the configuration.
	MulticastThreshold      *int `json:"multicastThreshold,omitempty"`      // Percentage (1 to 99) of total available port bandwidth for multicast traffic type. Default value 100 percent rate is to clear the configuration.
	UnknownUnicastThreshold *int `json:"unknownUnicastThreshold,omitempty"` // Percentage (1 to 99) of total available port bandwidth for unknown unicast (dlf-destination lookup failure) traffic type. Default value 100 percent rate is to clear the configuration.
}
type RequestSwitchUpdateNetworkSwitchStp struct {
	RstpEnabled       *bool                                                   `json:"rstpEnabled,omitempty"`       // The spanning tree protocol status in network
	StpBridgePriority *[]RequestSwitchUpdateNetworkSwitchStpStpBridgePriority `json:"stpBridgePriority,omitempty"` // STP bridge priority for switches/stacks or switch profiles. An empty array will clear the STP bridge priority settings.
}
type RequestSwitchUpdateNetworkSwitchStpStpBridgePriority struct {
	Stacks         []string `json:"stacks,omitempty"`         // List of stack IDs
	StpPriority    *int     `json:"stpPriority,omitempty"`    // STP priority for switch, stacks, or switch profiles
	SwitchProfiles []string `json:"switchProfiles,omitempty"` // List of switch profile IDs
	Switches       []string `json:"switches,omitempty"`       // List of switch serial numbers
}
type RequestSwitchUpdateOrganizationConfigTemplateSwitchProfilePort struct {
	AccessPolicyNumber      *int                                                                   `json:"accessPolicyNumber,omitempty"`      // The number of a custom access policy to configure on the switch profile port. Only applicable when 'accessPolicyType' is 'Custom access policy'.
	AccessPolicyType        string                                                                 `json:"accessPolicyType,omitempty"`        // The type of the access policy of the switch profile port. Only applicable to access ports. Can be one of 'Open', 'Custom access policy', 'MAC allow list' or 'Sticky MAC allow list'.
	AllowedVLANs            string                                                                 `json:"allowedVlans,omitempty"`            // The VLANs allowed on the switch profile port. Only applicable to trunk ports.
	DaiTrusted              *bool                                                                  `json:"daiTrusted,omitempty"`              // If true, ARP packets for this port will be considered trusted, and Dynamic ARP Inspection will allow the traffic.
	Enabled                 *bool                                                                  `json:"enabled,omitempty"`                 // The status of the switch profile port.
	FlexibleStackingEnabled *bool                                                                  `json:"flexibleStackingEnabled,omitempty"` // For supported switches (e.g. MS420/MS425), whether or not the port has flexible stacking enabled.
	IsolationEnabled        *bool                                                                  `json:"isolationEnabled,omitempty"`        // The isolation status of the switch profile port.
	LinkNegotiation         string                                                                 `json:"linkNegotiation,omitempty"`         // The link speed for the switch profile port.
	MacAllowList            []string                                                               `json:"macAllowList,omitempty"`            // Only devices with MAC addresses specified in this list will have access to this port. Up to 20 MAC addresses can be defined. Only applicable when 'accessPolicyType' is 'MAC allow list'.
	Name                    string                                                                 `json:"name,omitempty"`                    // The name of the switch profile port.
	PoeEnabled              *bool                                                                  `json:"poeEnabled,omitempty"`              // The PoE status of the switch profile port.
	PortScheduleID          string                                                                 `json:"portScheduleId,omitempty"`          // The ID of the port schedule. A value of null will clear the port schedule.
	Profile                 *RequestSwitchUpdateOrganizationConfigTemplateSwitchProfilePortProfile `json:"profile,omitempty"`                 // Profile attributes
	RstpEnabled             *bool                                                                  `json:"rstpEnabled,omitempty"`             // The rapid spanning tree protocol status.
	StickyMacAllowList      []string                                                               `json:"stickyMacAllowList,omitempty"`      // The initial list of MAC addresses for sticky Mac allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
	StickyMacAllowListLimit *int                                                                   `json:"stickyMacAllowListLimit,omitempty"` // The maximum number of MAC addresses for sticky MAC allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
	StormControlEnabled     *bool                                                                  `json:"stormControlEnabled,omitempty"`     // The storm control status of the switch profile port.
	StpGuard                string                                                                 `json:"stpGuard,omitempty"`                // The state of the STP guard ('disabled', 'root guard', 'bpdu guard' or 'loop guard').
	Tags                    []string                                                               `json:"tags,omitempty"`                    // The list of tags of the switch profile port.
	Type                    string                                                                 `json:"type,omitempty"`                    // The type of the switch profile port ('trunk' or 'access').
	Udld                    string                                                                 `json:"udld,omitempty"`                    // The action to take when Unidirectional Link is detected (Alert only, Enforce). Default configuration is Alert only.
	VLAN                    *int                                                                   `json:"vlan,omitempty"`                    // The VLAN of the switch profile port. A null value will clear the value set for trunk ports.
	VoiceVLAN               *int                                                                   `json:"voiceVlan,omitempty"`               // The voice VLAN of the switch profile port. Only applicable to access ports.
}
type RequestSwitchUpdateOrganizationConfigTemplateSwitchProfilePortProfile struct {
	Enabled *bool  `json:"enabled,omitempty"` // When enabled, override this port's configuration with a port profile.
	ID      string `json:"id,omitempty"`      // When enabled, the ID of the port profile used to override the port's configuration.
	Iname   string `json:"iname,omitempty"`   // When enabled, the IName of the profile.
}
type RequestSwitchCloneOrganizationSwitchDevices struct {
	SourceSerial  string   `json:"sourceSerial,omitempty"`  // Serial number of the source switch (must be on a network not bound to a template)
	TargetSerials []string `json:"targetSerials,omitempty"` // Array of serial numbers of one or more target switches (must be on a network not bound to a template)
}

//GetDeviceSwitchPorts List the switch ports for a switch
/* List the switch ports for a switch

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-switch-ports
*/
func (s *SwitchService) GetDeviceSwitchPorts(serial string) (*ResponseSwitchGetDeviceSwitchPorts, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/switch/ports"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSwitchGetDeviceSwitchPorts{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceSwitchPorts")
	}

	result := response.Result().(*ResponseSwitchGetDeviceSwitchPorts)
	return result, response, err

}

//GetDeviceSwitchPortsStatuses Return the status for all the ports of a switch
/* Return the status for all the ports of a switch

@param serial serial path parameter.
@param getDeviceSwitchPortsStatusesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-switch-ports-statuses
*/
func (s *SwitchService) GetDeviceSwitchPortsStatuses(serial string, getDeviceSwitchPortsStatusesQueryParams *GetDeviceSwitchPortsStatusesQueryParams) (*ResponseSwitchGetDeviceSwitchPortsStatuses, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/switch/ports/statuses"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	queryString, _ := query.Values(getDeviceSwitchPortsStatusesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSwitchGetDeviceSwitchPortsStatuses{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceSwitchPortsStatuses")
	}

	result := response.Result().(*ResponseSwitchGetDeviceSwitchPortsStatuses)
	return result, response, err

}

//GetDeviceSwitchPortsStatusesPackets Return the packet counters for all the ports of a switch
/* Return the packet counters for all the ports of a switch

@param serial serial path parameter.
@param getDeviceSwitchPortsStatusesPacketsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-switch-ports-statuses-packets
*/
func (s *SwitchService) GetDeviceSwitchPortsStatusesPackets(serial string, getDeviceSwitchPortsStatusesPacketsQueryParams *GetDeviceSwitchPortsStatusesPacketsQueryParams) (*ResponseSwitchGetDeviceSwitchPortsStatusesPackets, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/switch/ports/statuses/packets"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	queryString, _ := query.Values(getDeviceSwitchPortsStatusesPacketsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSwitchGetDeviceSwitchPortsStatusesPackets{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceSwitchPortsStatusesPackets")
	}

	result := response.Result().(*ResponseSwitchGetDeviceSwitchPortsStatusesPackets)
	return result, response, err

}

//GetDeviceSwitchPort Return a switch port
/* Return a switch port

@param serial serial path parameter.
@param portID portId path parameter. Port ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-switch-port
*/
func (s *SwitchService) GetDeviceSwitchPort(serial string, portID string) (*ResponseSwitchGetDeviceSwitchPort, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/switch/ports/{portId}"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)
	path = strings.Replace(path, "{portId}", fmt.Sprintf("%v", portID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSwitchGetDeviceSwitchPort{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceSwitchPort")
	}

	result := response.Result().(*ResponseSwitchGetDeviceSwitchPort)
	return result, response, err

}

//GetDeviceSwitchRoutingInterfaces List layer 3 interfaces for a switch
/* List layer 3 interfaces for a switch. Those for a stack may be found under switch stack routing.

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-switch-routing-interfaces
*/
func (s *SwitchService) GetDeviceSwitchRoutingInterfaces(serial string) (*ResponseSwitchGetDeviceSwitchRoutingInterfaces, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/switch/routing/interfaces"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSwitchGetDeviceSwitchRoutingInterfaces{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceSwitchRoutingInterfaces")
	}

	result := response.Result().(*ResponseSwitchGetDeviceSwitchRoutingInterfaces)
	return result, response, err

}

//GetDeviceSwitchRoutingInterface Return a layer 3 interface for a switch
/* Return a layer 3 interface for a switch

@param serial serial path parameter.
@param interfaceID interfaceId path parameter. Interface ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-switch-routing-interface
*/
func (s *SwitchService) GetDeviceSwitchRoutingInterface(serial string, interfaceID string) (*ResponseSwitchGetDeviceSwitchRoutingInterface, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/switch/routing/interfaces/{interfaceId}"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)
	path = strings.Replace(path, "{interfaceId}", fmt.Sprintf("%v", interfaceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSwitchGetDeviceSwitchRoutingInterface{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceSwitchRoutingInterface")
	}

	result := response.Result().(*ResponseSwitchGetDeviceSwitchRoutingInterface)
	return result, response, err

}

//GetDeviceSwitchRoutingInterfaceDhcp Return a layer 3 interface DHCP configuration for a switch
/* Return a layer 3 interface DHCP configuration for a switch

@param serial serial path parameter.
@param interfaceID interfaceId path parameter. Interface ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-switch-routing-interface-dhcp
*/
func (s *SwitchService) GetDeviceSwitchRoutingInterfaceDhcp(serial string, interfaceID string) (*ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcp, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/switch/routing/interfaces/{interfaceId}/dhcp"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)
	path = strings.Replace(path, "{interfaceId}", fmt.Sprintf("%v", interfaceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcp{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceSwitchRoutingInterfaceDhcp")
	}

	result := response.Result().(*ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcp)
	return result, response, err

}

//GetDeviceSwitchRoutingStaticRoutes List layer 3 static routes for a switch
/* List layer 3 static routes for a switch

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-switch-routing-static-routes
*/
func (s *SwitchService) GetDeviceSwitchRoutingStaticRoutes(serial string) (*ResponseSwitchGetDeviceSwitchRoutingStaticRoutes, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/switch/routing/staticRoutes"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSwitchGetDeviceSwitchRoutingStaticRoutes{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceSwitchRoutingStaticRoutes")
	}

	result := response.Result().(*ResponseSwitchGetDeviceSwitchRoutingStaticRoutes)
	return result, response, err

}

//GetDeviceSwitchRoutingStaticRoute Return a layer 3 static route for a switch
/* Return a layer 3 static route for a switch

@param serial serial path parameter.
@param staticRouteID staticRouteId path parameter. Static route ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-switch-routing-static-route
*/
func (s *SwitchService) GetDeviceSwitchRoutingStaticRoute(serial string, staticRouteID string) (*ResponseSwitchGetDeviceSwitchRoutingStaticRoute, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/switch/routing/staticRoutes/{staticRouteId}"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)
	path = strings.Replace(path, "{staticRouteId}", fmt.Sprintf("%v", staticRouteID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSwitchGetDeviceSwitchRoutingStaticRoute{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceSwitchRoutingStaticRoute")
	}

	result := response.Result().(*ResponseSwitchGetDeviceSwitchRoutingStaticRoute)
	return result, response, err

}

//GetDeviceSwitchWarmSpare Return warm spare configuration for a switch
/* Return warm spare configuration for a switch

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-switch-warm-spare
*/
func (s *SwitchService) GetDeviceSwitchWarmSpare(serial string) (*ResponseSwitchGetDeviceSwitchWarmSpare, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/switch/warmSpare"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSwitchGetDeviceSwitchWarmSpare{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceSwitchWarmSpare")
	}

	result := response.Result().(*ResponseSwitchGetDeviceSwitchWarmSpare)
	return result, response, err

}

//GetNetworkSwitchAccessControlLists Return the access control lists for a MS network
/* Return the access control lists for a MS network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-switch-access-control-lists
*/
func (s *SwitchService) GetNetworkSwitchAccessControlLists(networkID string) (*ResponseSwitchGetNetworkSwitchAccessControlLists, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/accessControlLists"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSwitchGetNetworkSwitchAccessControlLists{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSwitchAccessControlLists")
	}

	result := response.Result().(*ResponseSwitchGetNetworkSwitchAccessControlLists)
	return result, response, err

}

//GetNetworkSwitchAccessPolicies List the access policies for a switch network
/* List the access policies for a switch network. Only returns access policies with 'my RADIUS server' as authentication method

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-switch-access-policies
*/
func (s *SwitchService) GetNetworkSwitchAccessPolicies(networkID string) (*ResponseSwitchGetNetworkSwitchAccessPolicies, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/accessPolicies"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSwitchGetNetworkSwitchAccessPolicies{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSwitchAccessPolicies")
	}

	result := response.Result().(*ResponseSwitchGetNetworkSwitchAccessPolicies)
	return result, response, err

}

//GetNetworkSwitchAccessPolicy Return a specific access policy for a switch network
/* Return a specific access policy for a switch network

@param networkID networkId path parameter. Network ID
@param accessPolicyNumber accessPolicyNumber path parameter. Access policy number

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-switch-access-policy
*/
func (s *SwitchService) GetNetworkSwitchAccessPolicy(networkID string, accessPolicyNumber string) (*ResponseSwitchGetNetworkSwitchAccessPolicy, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/accessPolicies/{accessPolicyNumber}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{accessPolicyNumber}", fmt.Sprintf("%v", accessPolicyNumber), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSwitchGetNetworkSwitchAccessPolicy{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSwitchAccessPolicy")
	}

	result := response.Result().(*ResponseSwitchGetNetworkSwitchAccessPolicy)
	return result, response, err

}

//GetNetworkSwitchAlternateManagementInterface Return the switch alternate management interface for the network
/* Return the switch alternate management interface for the network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-switch-alternate-management-interface
*/
func (s *SwitchService) GetNetworkSwitchAlternateManagementInterface(networkID string) (*ResponseSwitchGetNetworkSwitchAlternateManagementInterface, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/alternateManagementInterface"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSwitchGetNetworkSwitchAlternateManagementInterface{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSwitchAlternateManagementInterface")
	}

	result := response.Result().(*ResponseSwitchGetNetworkSwitchAlternateManagementInterface)
	return result, response, err

}

//GetNetworkSwitchDhcpV4ServersSeen Return the network's DHCPv4 servers seen within the selected timeframe (default 1 day)
/* Return the network's DHCPv4 servers seen within the selected timeframe (default 1 day)

@param networkID networkId path parameter. Network ID
@param getNetworkSwitchDhcpV4ServersSeenQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-switch-dhcp-v4-servers-seen
*/
func (s *SwitchService) GetNetworkSwitchDhcpV4ServersSeen(networkID string, getNetworkSwitchDhcpV4ServersSeenQueryParams *GetNetworkSwitchDhcpV4ServersSeenQueryParams) (*ResponseSwitchGetNetworkSwitchDhcpV4ServersSeen, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/dhcp/v4/servers/seen"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(getNetworkSwitchDhcpV4ServersSeenQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSwitchGetNetworkSwitchDhcpV4ServersSeen{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSwitchDhcpV4ServersSeen")
	}

	result := response.Result().(*ResponseSwitchGetNetworkSwitchDhcpV4ServersSeen)
	return result, response, err

}

//GetNetworkSwitchDhcpServerPolicy Return the DHCP server settings
/* Return the DHCP server settings. Blocked/allowed servers are only applied when default policy is allow/block, respectively

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-switch-dhcp-server-policy
*/
func (s *SwitchService) GetNetworkSwitchDhcpServerPolicy(networkID string) (*ResponseSwitchGetNetworkSwitchDhcpServerPolicy, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/dhcpServerPolicy"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSwitchGetNetworkSwitchDhcpServerPolicy{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSwitchDhcpServerPolicy")
	}

	result := response.Result().(*ResponseSwitchGetNetworkSwitchDhcpServerPolicy)
	return result, response, err

}

//GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers Return the list of servers trusted by Dynamic ARP Inspection on this network
/* Return the list of servers trusted by Dynamic ARP Inspection on this network. These are also known as allow listed snoop entries

@param networkID networkId path parameter. Network ID
@param getNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-switch-dhcp-server-policy-arp-inspection-trusted-servers
*/
func (s *SwitchService) GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers(networkID string, getNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersQueryParams *GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersQueryParams) (*ResponseSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/dhcpServerPolicy/arpInspection/trustedServers"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(getNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers")
	}

	result := response.Result().(*ResponseSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers)
	return result, response, err

}

//GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice Return the devices that have a Dynamic ARP Inspection warning and their warnings
/* Return the devices that have a Dynamic ARP Inspection warning and their warnings

@param networkID networkId path parameter. Network ID
@param getNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-switch-dhcp-server-policy-arp-inspection-warnings-by-device
*/
func (s *SwitchService) GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice(networkID string, getNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceQueryParams *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceQueryParams) (*ResponseSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/dhcpServerPolicy/arpInspection/warnings/byDevice"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(getNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice")
	}

	result := response.Result().(*ResponseSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice)
	return result, response, err

}

//GetNetworkSwitchDscpToCosMappings Return the DSCP to CoS mappings
/* Return the DSCP to CoS mappings

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-switch-dscp-to-cos-mappings
*/
func (s *SwitchService) GetNetworkSwitchDscpToCosMappings(networkID string) (*ResponseSwitchGetNetworkSwitchDscpToCosMappings, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/dscpToCosMappings"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSwitchGetNetworkSwitchDscpToCosMappings{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSwitchDscpToCosMappings")
	}

	result := response.Result().(*ResponseSwitchGetNetworkSwitchDscpToCosMappings)
	return result, response, err

}

//GetNetworkSwitchLinkAggregations List link aggregation groups
/* List link aggregation groups

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-switch-link-aggregations
*/
func (s *SwitchService) GetNetworkSwitchLinkAggregations(networkID string) (*ResponseSwitchGetNetworkSwitchLinkAggregations, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/linkAggregations"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSwitchGetNetworkSwitchLinkAggregations{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSwitchLinkAggregations")
	}

	result := response.Result().(*ResponseSwitchGetNetworkSwitchLinkAggregations)
	return result, response, err

}

//GetNetworkSwitchMtu Return the MTU configuration
/* Return the MTU configuration

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-switch-mtu
*/
func (s *SwitchService) GetNetworkSwitchMtu(networkID string) (*ResponseSwitchGetNetworkSwitchMtu, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/mtu"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSwitchGetNetworkSwitchMtu{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSwitchMtu")
	}

	result := response.Result().(*ResponseSwitchGetNetworkSwitchMtu)
	return result, response, err

}

//GetNetworkSwitchPortSchedules List switch port schedules
/* List switch port schedules

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-switch-port-schedules
*/
func (s *SwitchService) GetNetworkSwitchPortSchedules(networkID string) (*ResponseSwitchGetNetworkSwitchPortSchedules, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/portSchedules"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSwitchGetNetworkSwitchPortSchedules{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSwitchPortSchedules")
	}

	result := response.Result().(*ResponseSwitchGetNetworkSwitchPortSchedules)
	return result, response, err

}

//GetNetworkSwitchQosRules List quality of service rules
/* List quality of service rules

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-switch-qos-rules
*/
func (s *SwitchService) GetNetworkSwitchQosRules(networkID string) (*ResponseSwitchGetNetworkSwitchQosRules, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/qosRules"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSwitchGetNetworkSwitchQosRules{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSwitchQosRules")
	}

	result := response.Result().(*ResponseSwitchGetNetworkSwitchQosRules)
	return result, response, err

}

//GetNetworkSwitchQosRulesOrder Return the quality of service rule IDs by order in which they will be processed by the switch
/* Return the quality of service rule IDs by order in which they will be processed by the switch

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-switch-qos-rules-order
*/
func (s *SwitchService) GetNetworkSwitchQosRulesOrder(networkID string) (*ResponseSwitchGetNetworkSwitchQosRulesOrder, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/qosRules/order"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSwitchGetNetworkSwitchQosRulesOrder{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSwitchQosRulesOrder")
	}

	result := response.Result().(*ResponseSwitchGetNetworkSwitchQosRulesOrder)
	return result, response, err

}

//GetNetworkSwitchQosRule Return a quality of service rule
/* Return a quality of service rule

@param networkID networkId path parameter. Network ID
@param qosRuleID qosRuleId path parameter. Qos rule ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-switch-qos-rule
*/
func (s *SwitchService) GetNetworkSwitchQosRule(networkID string, qosRuleID string) (*ResponseSwitchGetNetworkSwitchQosRule, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/qosRules/{qosRuleId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{qosRuleId}", fmt.Sprintf("%v", qosRuleID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSwitchGetNetworkSwitchQosRule{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSwitchQosRule")
	}

	result := response.Result().(*ResponseSwitchGetNetworkSwitchQosRule)
	return result, response, err

}

//GetNetworkSwitchRoutingMulticast Return multicast settings for a network
/* Return multicast settings for a network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-switch-routing-multicast
*/
func (s *SwitchService) GetNetworkSwitchRoutingMulticast(networkID string) (*ResponseSwitchGetNetworkSwitchRoutingMulticast, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/routing/multicast"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSwitchGetNetworkSwitchRoutingMulticast{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSwitchRoutingMulticast")
	}

	result := response.Result().(*ResponseSwitchGetNetworkSwitchRoutingMulticast)
	return result, response, err

}

//GetNetworkSwitchRoutingMulticastRendezvousPoints List multicast rendezvous points
/* List multicast rendezvous points

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-switch-routing-multicast-rendezvous-points
*/
func (s *SwitchService) GetNetworkSwitchRoutingMulticastRendezvousPoints(networkID string) (*ResponseSwitchGetNetworkSwitchRoutingMulticastRendezvousPoints, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/routing/multicast/rendezvousPoints"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSwitchGetNetworkSwitchRoutingMulticastRendezvousPoints{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSwitchRoutingMulticastRendezvousPoints")
	}

	result := response.Result().(*ResponseSwitchGetNetworkSwitchRoutingMulticastRendezvousPoints)
	return result, response, err

}

//GetNetworkSwitchRoutingMulticastRendezvousPoint Return a multicast rendezvous point
/* Return a multicast rendezvous point

@param networkID networkId path parameter. Network ID
@param rendezvousPointID rendezvousPointId path parameter. Rendezvous point ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-switch-routing-multicast-rendezvous-point
*/
func (s *SwitchService) GetNetworkSwitchRoutingMulticastRendezvousPoint(networkID string, rendezvousPointID string) (*ResponseSwitchGetNetworkSwitchRoutingMulticastRendezvousPoint, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/routing/multicast/rendezvousPoints/{rendezvousPointId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{rendezvousPointId}", fmt.Sprintf("%v", rendezvousPointID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSwitchGetNetworkSwitchRoutingMulticastRendezvousPoint{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSwitchRoutingMulticastRendezvousPoint")
	}

	result := response.Result().(*ResponseSwitchGetNetworkSwitchRoutingMulticastRendezvousPoint)
	return result, response, err

}

//GetNetworkSwitchRoutingOspf Return layer 3 OSPF routing configuration
/* Return layer 3 OSPF routing configuration

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-switch-routing-ospf
*/
func (s *SwitchService) GetNetworkSwitchRoutingOspf(networkID string) (*ResponseSwitchGetNetworkSwitchRoutingOspf, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/routing/ospf"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSwitchGetNetworkSwitchRoutingOspf{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSwitchRoutingOspf")
	}

	result := response.Result().(*ResponseSwitchGetNetworkSwitchRoutingOspf)
	return result, response, err

}

//GetNetworkSwitchSettings Returns the switch network settings
/* Returns the switch network settings

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-switch-settings
*/
func (s *SwitchService) GetNetworkSwitchSettings(networkID string) (*ResponseSwitchGetNetworkSwitchSettings, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/settings"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSwitchGetNetworkSwitchSettings{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSwitchSettings")
	}

	result := response.Result().(*ResponseSwitchGetNetworkSwitchSettings)
	return result, response, err

}

//GetNetworkSwitchStacks List the switch stacks in a network
/* List the switch stacks in a network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-switch-stacks
*/
func (s *SwitchService) GetNetworkSwitchStacks(networkID string) (*ResponseSwitchGetNetworkSwitchStacks, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/stacks"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSwitchGetNetworkSwitchStacks{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSwitchStacks")
	}

	result := response.Result().(*ResponseSwitchGetNetworkSwitchStacks)
	return result, response, err

}

//GetNetworkSwitchStack Show a switch stack
/* Show a switch stack

@param networkID networkId path parameter. Network ID
@param switchStackID switchStackId path parameter. Switch stack ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-switch-stack
*/
func (s *SwitchService) GetNetworkSwitchStack(networkID string, switchStackID string) (*ResponseSwitchGetNetworkSwitchStack, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/stacks/{switchStackId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{switchStackId}", fmt.Sprintf("%v", switchStackID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSwitchGetNetworkSwitchStack{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSwitchStack")
	}

	result := response.Result().(*ResponseSwitchGetNetworkSwitchStack)
	return result, response, err

}

//GetNetworkSwitchStackRoutingInterfaces List layer 3 interfaces for a switch stack
/* List layer 3 interfaces for a switch stack

@param networkID networkId path parameter. Network ID
@param switchStackID switchStackId path parameter. Switch stack ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-switch-stack-routing-interfaces
*/
func (s *SwitchService) GetNetworkSwitchStackRoutingInterfaces(networkID string, switchStackID string) (*ResponseSwitchGetNetworkSwitchStackRoutingInterfaces, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{switchStackId}", fmt.Sprintf("%v", switchStackID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSwitchGetNetworkSwitchStackRoutingInterfaces{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSwitchStackRoutingInterfaces")
	}

	result := response.Result().(*ResponseSwitchGetNetworkSwitchStackRoutingInterfaces)
	return result, response, err

}

//GetNetworkSwitchStackRoutingInterface Return a layer 3 interface from a switch stack
/* Return a layer 3 interface from a switch stack

@param networkID networkId path parameter. Network ID
@param switchStackID switchStackId path parameter. Switch stack ID
@param interfaceID interfaceId path parameter. Interface ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-switch-stack-routing-interface
*/
func (s *SwitchService) GetNetworkSwitchStackRoutingInterface(networkID string, switchStackID string, interfaceID string) (*ResponseSwitchGetNetworkSwitchStackRoutingInterface, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{switchStackId}", fmt.Sprintf("%v", switchStackID), -1)
	path = strings.Replace(path, "{interfaceId}", fmt.Sprintf("%v", interfaceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSwitchGetNetworkSwitchStackRoutingInterface{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSwitchStackRoutingInterface")
	}

	result := response.Result().(*ResponseSwitchGetNetworkSwitchStackRoutingInterface)
	return result, response, err

}

//GetNetworkSwitchStackRoutingInterfaceDhcp Return a layer 3 interface DHCP configuration for a switch stack
/* Return a layer 3 interface DHCP configuration for a switch stack

@param networkID networkId path parameter. Network ID
@param switchStackID switchStackId path parameter. Switch stack ID
@param interfaceID interfaceId path parameter. Interface ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-switch-stack-routing-interface-dhcp
*/
func (s *SwitchService) GetNetworkSwitchStackRoutingInterfaceDhcp(networkID string, switchStackID string, interfaceID string) (*ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcp, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId}/dhcp"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{switchStackId}", fmt.Sprintf("%v", switchStackID), -1)
	path = strings.Replace(path, "{interfaceId}", fmt.Sprintf("%v", interfaceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcp{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSwitchStackRoutingInterfaceDhcp")
	}

	result := response.Result().(*ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcp)
	return result, response, err

}

//GetNetworkSwitchStackRoutingStaticRoutes List layer 3 static routes for a switch stack
/* List layer 3 static routes for a switch stack

@param networkID networkId path parameter. Network ID
@param switchStackID switchStackId path parameter. Switch stack ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-switch-stack-routing-static-routes
*/
func (s *SwitchService) GetNetworkSwitchStackRoutingStaticRoutes(networkID string, switchStackID string) (*ResponseSwitchGetNetworkSwitchStackRoutingStaticRoutes, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/stacks/{switchStackId}/routing/staticRoutes"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{switchStackId}", fmt.Sprintf("%v", switchStackID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSwitchGetNetworkSwitchStackRoutingStaticRoutes{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSwitchStackRoutingStaticRoutes")
	}

	result := response.Result().(*ResponseSwitchGetNetworkSwitchStackRoutingStaticRoutes)
	return result, response, err

}

//GetNetworkSwitchStackRoutingStaticRoute Return a layer 3 static route for a switch stack
/* Return a layer 3 static route for a switch stack

@param networkID networkId path parameter. Network ID
@param switchStackID switchStackId path parameter. Switch stack ID
@param staticRouteID staticRouteId path parameter. Static route ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-switch-stack-routing-static-route
*/
func (s *SwitchService) GetNetworkSwitchStackRoutingStaticRoute(networkID string, switchStackID string, staticRouteID string) (*ResponseSwitchGetNetworkSwitchStackRoutingStaticRoute, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/stacks/{switchStackId}/routing/staticRoutes/{staticRouteId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{switchStackId}", fmt.Sprintf("%v", switchStackID), -1)
	path = strings.Replace(path, "{staticRouteId}", fmt.Sprintf("%v", staticRouteID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSwitchGetNetworkSwitchStackRoutingStaticRoute{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSwitchStackRoutingStaticRoute")
	}

	result := response.Result().(*ResponseSwitchGetNetworkSwitchStackRoutingStaticRoute)
	return result, response, err

}

//GetNetworkSwitchStormControl Return the storm control configuration for a switch network
/* Return the storm control configuration for a switch network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-switch-storm-control
*/
func (s *SwitchService) GetNetworkSwitchStormControl(networkID string) (*ResponseSwitchGetNetworkSwitchStormControl, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/stormControl"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSwitchGetNetworkSwitchStormControl{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSwitchStormControl")
	}

	result := response.Result().(*ResponseSwitchGetNetworkSwitchStormControl)
	return result, response, err

}

//GetNetworkSwitchStp Returns STP settings
/* Returns STP settings

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-switch-stp
*/
func (s *SwitchService) GetNetworkSwitchStp(networkID string) (*ResponseSwitchGetNetworkSwitchStp, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/stp"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSwitchGetNetworkSwitchStp{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSwitchStp")
	}

	result := response.Result().(*ResponseSwitchGetNetworkSwitchStp)
	return result, response, err

}

//GetOrganizationConfigTemplateSwitchProfiles List the switch profiles for your switch template configuration
/* List the switch profiles for your switch template configuration

@param organizationID organizationId path parameter. Organization ID
@param configTemplateID configTemplateId path parameter. Config template ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-config-template-switch-profiles
*/
func (s *SwitchService) GetOrganizationConfigTemplateSwitchProfiles(organizationID string, configTemplateID string) (*ResponseSwitchGetOrganizationConfigTemplateSwitchProfiles, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/configTemplates/{configTemplateId}/switch/profiles"
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{configTemplateId}", fmt.Sprintf("%v", configTemplateID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSwitchGetOrganizationConfigTemplateSwitchProfiles{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationConfigTemplateSwitchProfiles")
	}

	result := response.Result().(*ResponseSwitchGetOrganizationConfigTemplateSwitchProfiles)
	return result, response, err

}

//GetOrganizationConfigTemplateSwitchProfilePorts Return all the ports of a switch profile
/* Return all the ports of a switch profile

@param organizationID organizationId path parameter. Organization ID
@param configTemplateID configTemplateId path parameter. Config template ID
@param profileID profileId path parameter. Profile ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-config-template-switch-profile-ports
*/
func (s *SwitchService) GetOrganizationConfigTemplateSwitchProfilePorts(organizationID string, configTemplateID string, profileID string) (*ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePorts, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/configTemplates/{configTemplateId}/switch/profiles/{profileId}/ports"
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{configTemplateId}", fmt.Sprintf("%v", configTemplateID), -1)
	path = strings.Replace(path, "{profileId}", fmt.Sprintf("%v", profileID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePorts{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationConfigTemplateSwitchProfilePorts")
	}

	result := response.Result().(*ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePorts)
	return result, response, err

}

//GetOrganizationConfigTemplateSwitchProfilePort Return a switch profile port
/* Return a switch profile port

@param organizationID organizationId path parameter. Organization ID
@param configTemplateID configTemplateId path parameter. Config template ID
@param profileID profileId path parameter. Profile ID
@param portID portId path parameter. Port ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-config-template-switch-profile-port
*/
func (s *SwitchService) GetOrganizationConfigTemplateSwitchProfilePort(organizationID string, configTemplateID string, profileID string, portID string) (*ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePort, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/configTemplates/{configTemplateId}/switch/profiles/{profileId}/ports/{portId}"
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{configTemplateId}", fmt.Sprintf("%v", configTemplateID), -1)
	path = strings.Replace(path, "{profileId}", fmt.Sprintf("%v", profileID), -1)
	path = strings.Replace(path, "{portId}", fmt.Sprintf("%v", portID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePort{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationConfigTemplateSwitchProfilePort")
	}

	result := response.Result().(*ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePort)
	return result, response, err

}

//GetOrganizationSwitchPortsBySwitch List the switchports in an organization by switch
/* List the switchports in an organization by switch

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationSwitchPortsBySwitchQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-switch-ports-by-switch
*/
func (s *SwitchService) GetOrganizationSwitchPortsBySwitch(organizationID string, getOrganizationSwitchPortsBySwitchQueryParams *GetOrganizationSwitchPortsBySwitchQueryParams) (*ResponseSwitchGetOrganizationSwitchPortsBySwitch, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/switch/ports/bySwitch"
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationSwitchPortsBySwitchQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSwitchGetOrganizationSwitchPortsBySwitch{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationSwitchPortsBySwitch")
	}

	result := response.Result().(*ResponseSwitchGetOrganizationSwitchPortsBySwitch)
	return result, response, err

}

//CycleDeviceSwitchPorts Cycle a set of switch ports
/* Cycle a set of switch ports

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!cycle-device-switch-ports
*/

func (s *SwitchService) CycleDeviceSwitchPorts(serial string, requestSwitchCycleDeviceSwitchPorts *RequestSwitchCycleDeviceSwitchPorts) (*ResponseSwitchCycleDeviceSwitchPorts, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/switch/ports/cycle"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchCycleDeviceSwitchPorts).
		SetResult(&ResponseSwitchCycleDeviceSwitchPorts{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CycleDeviceSwitchPorts")
	}

	result := response.Result().(*ResponseSwitchCycleDeviceSwitchPorts)
	return result, response, err

}

//CreateDeviceSwitchRoutingInterface Create a layer 3 interface for a switch
/* Create a layer 3 interface for a switch

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-device-switch-routing-interface
*/

func (s *SwitchService) CreateDeviceSwitchRoutingInterface(serial string, requestSwitchCreateDeviceSwitchRoutingInterface *RequestSwitchCreateDeviceSwitchRoutingInterface) (*ResponseSwitchCreateDeviceSwitchRoutingInterface, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/switch/routing/interfaces"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchCreateDeviceSwitchRoutingInterface).
		SetResult(&ResponseSwitchCreateDeviceSwitchRoutingInterface{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateDeviceSwitchRoutingInterface")
	}

	result := response.Result().(*ResponseSwitchCreateDeviceSwitchRoutingInterface)
	return result, response, err

}

//CreateDeviceSwitchRoutingStaticRoute Create a layer 3 static route for a switch
/* Create a layer 3 static route for a switch

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-device-switch-routing-static-route
*/

func (s *SwitchService) CreateDeviceSwitchRoutingStaticRoute(serial string, requestSwitchCreateDeviceSwitchRoutingStaticRoute *RequestSwitchCreateDeviceSwitchRoutingStaticRoute) (*resty.Response, error) {
	path := "/api/v1/devices/{serial}/switch/routing/staticRoutes"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchCreateDeviceSwitchRoutingStaticRoute).
		// SetResult(&ResponseSwitchCreateDeviceSwitchRoutingStaticRoute{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateDeviceSwitchRoutingStaticRoute")
	}

	return response, err

}

//CreateNetworkSwitchAccessPolicy Create an access policy for a switch network
/* Create an access policy for a switch network. If you would like to enable Meraki Authentication, set radiusServers to empty array.

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-network-switch-access-policy
*/

func (s *SwitchService) CreateNetworkSwitchAccessPolicy(networkID string, requestSwitchCreateNetworkSwitchAccessPolicy *RequestSwitchCreateNetworkSwitchAccessPolicy) (*ResponseSwitchCreateNetworkSwitchAccessPolicy, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/accessPolicies"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchCreateNetworkSwitchAccessPolicy).
		SetResult(&ResponseSwitchCreateNetworkSwitchAccessPolicy{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateNetworkSwitchAccessPolicy")
	}

	result := response.Result().(*ResponseSwitchCreateNetworkSwitchAccessPolicy)
	return result, response, err

}

//CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer Add a server to be trusted by Dynamic ARP Inspection on this network
/* Add a server to be trusted by Dynamic ARP Inspection on this network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-network-switch-dhcp-server-policy-arp-inspection-trusted-server
*/

func (s *SwitchService) CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(networkID string, requestSwitchCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer *RequestSwitchCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer) (*ResponseSwitchCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/dhcpServerPolicy/arpInspection/trustedServers"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer).
		SetResult(&ResponseSwitchCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer")
	}

	result := response.Result().(*ResponseSwitchCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer)
	return result, response, err

}

//CreateNetworkSwitchLinkAggregation Create a link aggregation group
/* Create a link aggregation group

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-network-switch-link-aggregation
*/

func (s *SwitchService) CreateNetworkSwitchLinkAggregation(networkID string, requestSwitchCreateNetworkSwitchLinkAggregation *RequestSwitchCreateNetworkSwitchLinkAggregation) (*ResponseSwitchCreateNetworkSwitchLinkAggregation, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/linkAggregations"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchCreateNetworkSwitchLinkAggregation).
		SetResult(&ResponseSwitchCreateNetworkSwitchLinkAggregation{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateNetworkSwitchLinkAggregation")
	}

	result := response.Result().(*ResponseSwitchCreateNetworkSwitchLinkAggregation)
	return result, response, err

}

//CreateNetworkSwitchPortSchedule Add a switch port schedule
/* Add a switch port schedule

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-network-switch-port-schedule
*/

func (s *SwitchService) CreateNetworkSwitchPortSchedule(networkID string, requestSwitchCreateNetworkSwitchPortSchedule *RequestSwitchCreateNetworkSwitchPortSchedule) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/portSchedules"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchCreateNetworkSwitchPortSchedule).
		// SetResult(&ResponseSwitchCreateNetworkSwitchPortSchedule{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateNetworkSwitchPortSchedule")
	}

	return response, err

}

//CreateNetworkSwitchQosRule Add a quality of service rule
/* Add a quality of service rule

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-network-switch-qos-rule
*/

func (s *SwitchService) CreateNetworkSwitchQosRule(networkID string, requestSwitchCreateNetworkSwitchQosRule *RequestSwitchCreateNetworkSwitchQosRule) (*ResponseSwitchCreateNetworkSwitchQosRule, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/qosRules"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchCreateNetworkSwitchQosRule).
		SetResult(&ResponseSwitchCreateNetworkSwitchQosRule{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateNetworkSwitchQosRule")
	}

	result := response.Result().(*ResponseSwitchCreateNetworkSwitchQosRule)
	return result, response, err

}

//CreateNetworkSwitchRoutingMulticastRendezvousPoint Create a multicast rendezvous point
/* Create a multicast rendezvous point

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-network-switch-routing-multicast-rendezvous-point
*/

func (s *SwitchService) CreateNetworkSwitchRoutingMulticastRendezvousPoint(networkID string, requestSwitchCreateNetworkSwitchRoutingMulticastRendezvousPoint *RequestSwitchCreateNetworkSwitchRoutingMulticastRendezvousPoint) (*ResponseSwitchCreateNetworkSwitchRoutingMulticastRendezvousPoint, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/routing/multicast/rendezvousPoints"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchCreateNetworkSwitchRoutingMulticastRendezvousPoint).
		SetResult(&ResponseSwitchCreateNetworkSwitchRoutingMulticastRendezvousPoint{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateNetworkSwitchRoutingMulticastRendezvousPoint")
	}

	result := response.Result().(*ResponseSwitchCreateNetworkSwitchRoutingMulticastRendezvousPoint)
	return result, response, err

}

//CreateNetworkSwitchStack Create a stack
/* Create a stack

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-network-switch-stack
*/

func (s *SwitchService) CreateNetworkSwitchStack(networkID string, requestSwitchCreateNetworkSwitchStack *RequestSwitchCreateNetworkSwitchStack) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/stacks"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchCreateNetworkSwitchStack).
		// SetResult(&ResponseSwitchCreateNetworkSwitchStack{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateNetworkSwitchStack")
	}

	return response, err

}

//AddNetworkSwitchStack Add a switch to a stack
/* Add a switch to a stack

@param networkID networkId path parameter. Network ID
@param switchStackID switchStackId path parameter. Switch stack ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-network-switch-stack
*/

func (s *SwitchService) AddNetworkSwitchStack(networkID string, switchStackID string, requestSwitchAddNetworkSwitchStack *RequestSwitchAddNetworkSwitchStack) (*ResponseSwitchAddNetworkSwitchStack, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/stacks/{switchStackId}/add"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{switchStackId}", fmt.Sprintf("%v", switchStackID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchAddNetworkSwitchStack).
		SetResult(&ResponseSwitchAddNetworkSwitchStack{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation AddNetworkSwitchStack")
	}

	result := response.Result().(*ResponseSwitchAddNetworkSwitchStack)
	return result, response, err

}

//RemoveNetworkSwitchStack Remove a switch from a stack
/* Remove a switch from a stack

@param networkID networkId path parameter. Network ID
@param switchStackID switchStackId path parameter. Switch stack ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!remove-network-switch-stack
*/

func (s *SwitchService) RemoveNetworkSwitchStack(networkID string, switchStackID string, requestSwitchRemoveNetworkSwitchStack *RequestSwitchRemoveNetworkSwitchStack) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/stacks/{switchStackId}/remove"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{switchStackId}", fmt.Sprintf("%v", switchStackID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchRemoveNetworkSwitchStack).
		// SetResult(&ResponseSwitchRemoveNetworkSwitchStack{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation RemoveNetworkSwitchStack")
	}

	return response, err

}

//CreateNetworkSwitchStackRoutingInterface Create a layer 3 interface for a switch stack
/* Create a layer 3 interface for a switch stack

@param networkID networkId path parameter. Network ID
@param switchStackID switchStackId path parameter. Switch stack ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-network-switch-stack-routing-interface
*/

func (s *SwitchService) CreateNetworkSwitchStackRoutingInterface(networkID string, switchStackID string, requestSwitchCreateNetworkSwitchStackRoutingInterface *RequestSwitchCreateNetworkSwitchStackRoutingInterface) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{switchStackId}", fmt.Sprintf("%v", switchStackID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchCreateNetworkSwitchStackRoutingInterface).
		// SetResult(&ResponseSwitchCreateNetworkSwitchStackRoutingInterface{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateNetworkSwitchStackRoutingInterface")
	}

	return response, err

}

//CreateNetworkSwitchStackRoutingStaticRoute Create a layer 3 static route for a switch stack
/* Create a layer 3 static route for a switch stack

@param networkID networkId path parameter. Network ID
@param switchStackID switchStackId path parameter. Switch stack ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-network-switch-stack-routing-static-route
*/

func (s *SwitchService) CreateNetworkSwitchStackRoutingStaticRoute(networkID string, switchStackID string, requestSwitchCreateNetworkSwitchStackRoutingStaticRoute *RequestSwitchCreateNetworkSwitchStackRoutingStaticRoute) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/stacks/{switchStackId}/routing/staticRoutes"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{switchStackId}", fmt.Sprintf("%v", switchStackID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchCreateNetworkSwitchStackRoutingStaticRoute).
		// SetResult(&ResponseSwitchCreateNetworkSwitchStackRoutingStaticRoute{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateNetworkSwitchStackRoutingStaticRoute")
	}

	return response, err

}

//CloneOrganizationSwitchDevices Clone port-level and some switch-level configuration settings from a source switch to one or more target switches
/* Clone port-level and some switch-level configuration settings from a source switch to one or more target switches. Cloned settings include: Aggregation Groups, Power Settings, Multicast Settings, MTU Configuration, STP Bridge priority, Port Mirroring

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!clone-organization-switch-devices
*/

func (s *SwitchService) CloneOrganizationSwitchDevices(organizationID string, requestSwitchCloneOrganizationSwitchDevices *RequestSwitchCloneOrganizationSwitchDevices) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/switch/devices/clone"
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchCloneOrganizationSwitchDevices).
		// SetResult(&ResponseSwitchCloneOrganizationSwitchDevices{}).
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

//UpdateDeviceSwitchPort Update a switch port
/* Update a switch port

@param serial serial path parameter.
@param portID portId path parameter. Port ID
*/
func (s *SwitchService) UpdateDeviceSwitchPort(serial string, portID string, requestSwitchUpdateDeviceSwitchPort *RequestSwitchUpdateDeviceSwitchPort) (*ResponseSwitchUpdateDeviceSwitchPort, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/switch/ports/{portId}"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)
	path = strings.Replace(path, "{portId}", fmt.Sprintf("%v", portID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateDeviceSwitchPort).
		SetResult(&ResponseSwitchUpdateDeviceSwitchPort{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateDeviceSwitchPort")
	}

	result := response.Result().(*ResponseSwitchUpdateDeviceSwitchPort)
	return result, response, err

}

//UpdateDeviceSwitchRoutingInterface Update a layer 3 interface for a switch
/* Update a layer 3 interface for a switch

@param serial serial path parameter.
@param interfaceID interfaceId path parameter. Interface ID
*/
func (s *SwitchService) UpdateDeviceSwitchRoutingInterface(serial string, interfaceID string, requestSwitchUpdateDeviceSwitchRoutingInterface *RequestSwitchUpdateDeviceSwitchRoutingInterface) (*ResponseSwitchUpdateDeviceSwitchRoutingInterface, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/switch/routing/interfaces/{interfaceId}"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)
	path = strings.Replace(path, "{interfaceId}", fmt.Sprintf("%v", interfaceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateDeviceSwitchRoutingInterface).
		SetResult(&ResponseSwitchUpdateDeviceSwitchRoutingInterface{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateDeviceSwitchRoutingInterface")
	}

	result := response.Result().(*ResponseSwitchUpdateDeviceSwitchRoutingInterface)
	return result, response, err

}

//UpdateDeviceSwitchRoutingInterfaceDhcp Update a layer 3 interface DHCP configuration for a switch
/* Update a layer 3 interface DHCP configuration for a switch

@param serial serial path parameter.
@param interfaceID interfaceId path parameter. Interface ID
*/
func (s *SwitchService) UpdateDeviceSwitchRoutingInterfaceDhcp(serial string, interfaceID string, requestSwitchUpdateDeviceSwitchRoutingInterfaceDhcp *RequestSwitchUpdateDeviceSwitchRoutingInterfaceDhcp) (*resty.Response, error) {
	path := "/api/v1/devices/{serial}/switch/routing/interfaces/{interfaceId}/dhcp"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)
	path = strings.Replace(path, "{interfaceId}", fmt.Sprintf("%v", interfaceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateDeviceSwitchRoutingInterfaceDhcp).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateDeviceSwitchRoutingInterfaceDhcp")
	}

	return response, err

}

//UpdateDeviceSwitchRoutingStaticRoute Update a layer 3 static route for a switch
/* Update a layer 3 static route for a switch

@param serial serial path parameter.
@param staticRouteID staticRouteId path parameter. Static route ID
*/
func (s *SwitchService) UpdateDeviceSwitchRoutingStaticRoute(serial string, staticRouteID string, requestSwitchUpdateDeviceSwitchRoutingStaticRoute *RequestSwitchUpdateDeviceSwitchRoutingStaticRoute) (*resty.Response, error) {
	path := "/api/v1/devices/{serial}/switch/routing/staticRoutes/{staticRouteId}"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)
	path = strings.Replace(path, "{staticRouteId}", fmt.Sprintf("%v", staticRouteID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateDeviceSwitchRoutingStaticRoute).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateDeviceSwitchRoutingStaticRoute")
	}

	return response, err

}

//UpdateDeviceSwitchWarmSpare Update warm spare configuration for a switch
/* Update warm spare configuration for a switch. The spare will use the same L3 configuration as the primary. Note that this will irreversibly destroy any existing L3 configuration on the spare.

@param serial serial path parameter.
*/
func (s *SwitchService) UpdateDeviceSwitchWarmSpare(serial string, requestSwitchUpdateDeviceSwitchWarmSpare *RequestSwitchUpdateDeviceSwitchWarmSpare) (*resty.Response, error) {
	path := "/api/v1/devices/{serial}/switch/warmSpare"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateDeviceSwitchWarmSpare).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateDeviceSwitchWarmSpare")
	}

	return response, err

}

//UpdateNetworkSwitchAccessControlLists Update the access control lists for a MS network
/* Update the access control lists for a MS network

@param networkID networkId path parameter. Network ID
*/
func (s *SwitchService) UpdateNetworkSwitchAccessControlLists(networkID string, requestSwitchUpdateNetworkSwitchAccessControlLists *RequestSwitchUpdateNetworkSwitchAccessControlLists) (*ResponseSwitchUpdateNetworkSwitchAccessControlLists, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/accessControlLists"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateNetworkSwitchAccessControlLists).
		SetResult(&ResponseSwitchUpdateNetworkSwitchAccessControlLists{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkSwitchAccessControlLists")
	}

	result := response.Result().(*ResponseSwitchUpdateNetworkSwitchAccessControlLists)
	return result, response, err

}

//UpdateNetworkSwitchAccessPolicy Update an access policy for a switch network
/* Update an access policy for a switch network. If you would like to enable Meraki Authentication, set radiusServers to empty array.

@param networkID networkId path parameter. Network ID
@param accessPolicyNumber accessPolicyNumber path parameter. Access policy number
*/
func (s *SwitchService) UpdateNetworkSwitchAccessPolicy(networkID string, accessPolicyNumber string, requestSwitchUpdateNetworkSwitchAccessPolicy *RequestSwitchUpdateNetworkSwitchAccessPolicy) (*ResponseSwitchUpdateNetworkSwitchAccessPolicy, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/accessPolicies/{accessPolicyNumber}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{accessPolicyNumber}", fmt.Sprintf("%v", accessPolicyNumber), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateNetworkSwitchAccessPolicy).
		SetResult(&ResponseSwitchUpdateNetworkSwitchAccessPolicy{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkSwitchAccessPolicy")
	}

	result := response.Result().(*ResponseSwitchUpdateNetworkSwitchAccessPolicy)
	return result, response, err

}

//UpdateNetworkSwitchAlternateManagementInterface Update the switch alternate management interface for the network
/* Update the switch alternate management interface for the network

@param networkID networkId path parameter. Network ID
*/
func (s *SwitchService) UpdateNetworkSwitchAlternateManagementInterface(networkID string, requestSwitchUpdateNetworkSwitchAlternateManagementInterface *RequestSwitchUpdateNetworkSwitchAlternateManagementInterface) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/alternateManagementInterface"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateNetworkSwitchAlternateManagementInterface).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkSwitchAlternateManagementInterface")
	}

	return response, err

}

//UpdateNetworkSwitchDhcpServerPolicy Update the DHCP server settings
/* Update the DHCP server settings. Blocked/allowed servers are only applied when default policy is allow/block, respectively

@param networkID networkId path parameter. Network ID
*/
func (s *SwitchService) UpdateNetworkSwitchDhcpServerPolicy(networkID string, requestSwitchUpdateNetworkSwitchDhcpServerPolicy *RequestSwitchUpdateNetworkSwitchDhcpServerPolicy) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/dhcpServerPolicy"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateNetworkSwitchDhcpServerPolicy).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkSwitchDhcpServerPolicy")
	}

	return response, err

}

//UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer Update a server that is trusted by Dynamic ARP Inspection on this network
/* Update a server that is trusted by Dynamic ARP Inspection on this network

@param networkID networkId path parameter. Network ID
@param trustedServerID trustedServerId path parameter. Trusted server ID
*/
func (s *SwitchService) UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(networkID string, trustedServerID string, requestSwitchUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer *RequestSwitchUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer) (*ResponseSwitchUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/dhcpServerPolicy/arpInspection/trustedServers/{trustedServerId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{trustedServerId}", fmt.Sprintf("%v", trustedServerID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer).
		SetResult(&ResponseSwitchUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer")
	}

	result := response.Result().(*ResponseSwitchUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer)
	return result, response, err

}

//UpdateNetworkSwitchDscpToCosMappings Update the DSCP to CoS mappings
/* Update the DSCP to CoS mappings

@param networkID networkId path parameter. Network ID
*/
func (s *SwitchService) UpdateNetworkSwitchDscpToCosMappings(networkID string, requestSwitchUpdateNetworkSwitchDscpToCosMappings *RequestSwitchUpdateNetworkSwitchDscpToCosMappings) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/dscpToCosMappings"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateNetworkSwitchDscpToCosMappings).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkSwitchDscpToCosMappings")
	}

	return response, err

}

//UpdateNetworkSwitchLinkAggregation Update a link aggregation group
/* Update a link aggregation group

@param networkID networkId path parameter. Network ID
@param linkAggregationID linkAggregationId path parameter. Link aggregation ID
*/
func (s *SwitchService) UpdateNetworkSwitchLinkAggregation(networkID string, linkAggregationID string, requestSwitchUpdateNetworkSwitchLinkAggregation *RequestSwitchUpdateNetworkSwitchLinkAggregation) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/linkAggregations/{linkAggregationId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{linkAggregationId}", fmt.Sprintf("%v", linkAggregationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateNetworkSwitchLinkAggregation).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkSwitchLinkAggregation")
	}

	return response, err

}

//UpdateNetworkSwitchMtu Update the MTU configuration
/* Update the MTU configuration

@param networkID networkId path parameter. Network ID
*/
func (s *SwitchService) UpdateNetworkSwitchMtu(networkID string, requestSwitchUpdateNetworkSwitchMtu *RequestSwitchUpdateNetworkSwitchMtu) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/mtu"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateNetworkSwitchMtu).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkSwitchMtu")
	}

	return response, err

}

//UpdateNetworkSwitchPortSchedule Update a switch port schedule
/* Update a switch port schedule

@param networkID networkId path parameter. Network ID
@param portScheduleID portScheduleId path parameter. Port schedule ID
*/
func (s *SwitchService) UpdateNetworkSwitchPortSchedule(networkID string, portScheduleID string, requestSwitchUpdateNetworkSwitchPortSchedule *RequestSwitchUpdateNetworkSwitchPortSchedule) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/portSchedules/{portScheduleId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{portScheduleId}", fmt.Sprintf("%v", portScheduleID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateNetworkSwitchPortSchedule).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkSwitchPortSchedule")
	}

	return response, err

}

//UpdateNetworkSwitchQosRulesOrder Update the order in which the rules should be processed by the switch
/* Update the order in which the rules should be processed by the switch

@param networkID networkId path parameter. Network ID
*/
func (s *SwitchService) UpdateNetworkSwitchQosRulesOrder(networkID string, requestSwitchUpdateNetworkSwitchQosRulesOrder *RequestSwitchUpdateNetworkSwitchQosRulesOrder) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/qosRules/order"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateNetworkSwitchQosRulesOrder).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkSwitchQosRulesOrder")
	}

	return response, err

}

//UpdateNetworkSwitchQosRule Update a quality of service rule
/* Update a quality of service rule

@param networkID networkId path parameter. Network ID
@param qosRuleID qosRuleId path parameter. Qos rule ID
*/
func (s *SwitchService) UpdateNetworkSwitchQosRule(networkID string, qosRuleID string, requestSwitchUpdateNetworkSwitchQosRule *RequestSwitchUpdateNetworkSwitchQosRule) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/qosRules/{qosRuleId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{qosRuleId}", fmt.Sprintf("%v", qosRuleID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateNetworkSwitchQosRule).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkSwitchQosRule")
	}

	return response, err

}

//UpdateNetworkSwitchRoutingMulticast Update multicast settings for a network
/* Update multicast settings for a network

@param networkID networkId path parameter. Network ID
*/
func (s *SwitchService) UpdateNetworkSwitchRoutingMulticast(networkID string, requestSwitchUpdateNetworkSwitchRoutingMulticast *RequestSwitchUpdateNetworkSwitchRoutingMulticast) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/routing/multicast"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateNetworkSwitchRoutingMulticast).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkSwitchRoutingMulticast")
	}

	return response, err

}

//UpdateNetworkSwitchRoutingMulticastRendezvousPoint Update a multicast rendezvous point
/* Update a multicast rendezvous point

@param networkID networkId path parameter. Network ID
@param rendezvousPointID rendezvousPointId path parameter. Rendezvous point ID
*/
func (s *SwitchService) UpdateNetworkSwitchRoutingMulticastRendezvousPoint(networkID string, rendezvousPointID string, requestSwitchUpdateNetworkSwitchRoutingMulticastRendezvousPoint *RequestSwitchUpdateNetworkSwitchRoutingMulticastRendezvousPoint) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/routing/multicast/rendezvousPoints/{rendezvousPointId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{rendezvousPointId}", fmt.Sprintf("%v", rendezvousPointID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateNetworkSwitchRoutingMulticastRendezvousPoint).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkSwitchRoutingMulticastRendezvousPoint")
	}

	return response, err

}

//UpdateNetworkSwitchRoutingOspf Update layer 3 OSPF routing configuration
/* Update layer 3 OSPF routing configuration

@param networkID networkId path parameter. Network ID
*/
func (s *SwitchService) UpdateNetworkSwitchRoutingOspf(networkID string, requestSwitchUpdateNetworkSwitchRoutingOspf *RequestSwitchUpdateNetworkSwitchRoutingOspf) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/routing/ospf"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateNetworkSwitchRoutingOspf).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkSwitchRoutingOspf")
	}

	return response, err

}

//UpdateNetworkSwitchSettings Update switch network settings
/* Update switch network settings

@param networkID networkId path parameter. Network ID
*/
func (s *SwitchService) UpdateNetworkSwitchSettings(networkID string, requestSwitchUpdateNetworkSwitchSettings *RequestSwitchUpdateNetworkSwitchSettings) (*ResponseSwitchUpdateNetworkSwitchSettings, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/settings"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateNetworkSwitchSettings).
		SetResult(&ResponseSwitchUpdateNetworkSwitchSettings{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkSwitchSettings")
	}

	result := response.Result().(*ResponseSwitchUpdateNetworkSwitchSettings)
	return result, response, err

}

//UpdateNetworkSwitchStackRoutingInterface Update a layer 3 interface for a switch stack
/* Update a layer 3 interface for a switch stack

@param networkID networkId path parameter. Network ID
@param switchStackID switchStackId path parameter. Switch stack ID
@param interfaceID interfaceId path parameter. Interface ID
*/
func (s *SwitchService) UpdateNetworkSwitchStackRoutingInterface(networkID string, switchStackID string, interfaceID string, requestSwitchUpdateNetworkSwitchStackRoutingInterface *RequestSwitchUpdateNetworkSwitchStackRoutingInterface) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{switchStackId}", fmt.Sprintf("%v", switchStackID), -1)
	path = strings.Replace(path, "{interfaceId}", fmt.Sprintf("%v", interfaceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateNetworkSwitchStackRoutingInterface).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkSwitchStackRoutingInterface")
	}

	return response, err

}

//UpdateNetworkSwitchStackRoutingInterfaceDhcp Update a layer 3 interface DHCP configuration for a switch stack
/* Update a layer 3 interface DHCP configuration for a switch stack

@param networkID networkId path parameter. Network ID
@param switchStackID switchStackId path parameter. Switch stack ID
@param interfaceID interfaceId path parameter. Interface ID
*/
func (s *SwitchService) UpdateNetworkSwitchStackRoutingInterfaceDhcp(networkID string, switchStackID string, interfaceID string, requestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcp *RequestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcp) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId}/dhcp"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{switchStackId}", fmt.Sprintf("%v", switchStackID), -1)
	path = strings.Replace(path, "{interfaceId}", fmt.Sprintf("%v", interfaceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcp).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkSwitchStackRoutingInterfaceDhcp")
	}

	return response, err

}

//UpdateNetworkSwitchStackRoutingStaticRoute Update a layer 3 static route for a switch stack
/* Update a layer 3 static route for a switch stack

@param networkID networkId path parameter. Network ID
@param switchStackID switchStackId path parameter. Switch stack ID
@param staticRouteID staticRouteId path parameter. Static route ID
*/
func (s *SwitchService) UpdateNetworkSwitchStackRoutingStaticRoute(networkID string, switchStackID string, staticRouteID string, requestSwitchUpdateNetworkSwitchStackRoutingStaticRoute *RequestSwitchUpdateNetworkSwitchStackRoutingStaticRoute) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/stacks/{switchStackId}/routing/staticRoutes/{staticRouteId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{switchStackId}", fmt.Sprintf("%v", switchStackID), -1)
	path = strings.Replace(path, "{staticRouteId}", fmt.Sprintf("%v", staticRouteID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateNetworkSwitchStackRoutingStaticRoute).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkSwitchStackRoutingStaticRoute")
	}

	return response, err

}

//UpdateNetworkSwitchStormControl Update the storm control configuration for a switch network
/* Update the storm control configuration for a switch network

@param networkID networkId path parameter. Network ID
*/
func (s *SwitchService) UpdateNetworkSwitchStormControl(networkID string, requestSwitchUpdateNetworkSwitchStormControl *RequestSwitchUpdateNetworkSwitchStormControl) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/stormControl"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateNetworkSwitchStormControl).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkSwitchStormControl")
	}

	return response, err

}

//UpdateNetworkSwitchStp Updates STP settings
/* Updates STP settings

@param networkID networkId path parameter. Network ID
*/
func (s *SwitchService) UpdateNetworkSwitchStp(networkID string, requestSwitchUpdateNetworkSwitchStp *RequestSwitchUpdateNetworkSwitchStp) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/stp"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateNetworkSwitchStp).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkSwitchStp")
	}

	return response, err

}

//UpdateOrganizationConfigTemplateSwitchProfilePort Update a switch profile port
/* Update a switch profile port

@param organizationID organizationId path parameter. Organization ID
@param configTemplateID configTemplateId path parameter. Config template ID
@param profileID profileId path parameter. Profile ID
@param portID portId path parameter. Port ID
*/
func (s *SwitchService) UpdateOrganizationConfigTemplateSwitchProfilePort(organizationID string, configTemplateID string, profileID string, portID string, requestSwitchUpdateOrganizationConfigTemplateSwitchProfilePort *RequestSwitchUpdateOrganizationConfigTemplateSwitchProfilePort) (*ResponseSwitchUpdateOrganizationConfigTemplateSwitchProfilePort, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/configTemplates/{configTemplateId}/switch/profiles/{profileId}/ports/{portId}"
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{configTemplateId}", fmt.Sprintf("%v", configTemplateID), -1)
	path = strings.Replace(path, "{profileId}", fmt.Sprintf("%v", profileID), -1)
	path = strings.Replace(path, "{portId}", fmt.Sprintf("%v", portID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateOrganizationConfigTemplateSwitchProfilePort).
		SetResult(&ResponseSwitchUpdateOrganizationConfigTemplateSwitchProfilePort{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateOrganizationConfigTemplateSwitchProfilePort")
	}

	result := response.Result().(*ResponseSwitchUpdateOrganizationConfigTemplateSwitchProfilePort)
	return result, response, err

}

//DeleteDeviceSwitchRoutingInterface Delete a layer 3 interface from the switch
/* Delete a layer 3 interface from the switch

@param serial serial path parameter.
@param interfaceID interfaceId path parameter. Interface ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-device-switch-routing-interface
*/
func (s *SwitchService) DeleteDeviceSwitchRoutingInterface(serial string, interfaceID string) (*resty.Response, error) {
	//serial string,interfaceID string
	path := "/api/v1/devices/{serial}/switch/routing/interfaces/{interfaceId}"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)
	path = strings.Replace(path, "{interfaceId}", fmt.Sprintf("%v", interfaceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteDeviceSwitchRoutingInterface")
	}

	return response, err

}

//DeleteDeviceSwitchRoutingStaticRoute Delete a layer 3 static route for a switch
/* Delete a layer 3 static route for a switch

@param serial serial path parameter.
@param staticRouteID staticRouteId path parameter. Static route ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-device-switch-routing-static-route
*/
func (s *SwitchService) DeleteDeviceSwitchRoutingStaticRoute(serial string, staticRouteID string) (*resty.Response, error) {
	//serial string,staticRouteID string
	path := "/api/v1/devices/{serial}/switch/routing/staticRoutes/{staticRouteId}"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)
	path = strings.Replace(path, "{staticRouteId}", fmt.Sprintf("%v", staticRouteID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteDeviceSwitchRoutingStaticRoute")
	}

	return response, err

}

//DeleteNetworkSwitchAccessPolicy Delete an access policy for a switch network
/* Delete an access policy for a switch network

@param networkID networkId path parameter. Network ID
@param accessPolicyNumber accessPolicyNumber path parameter. Access policy number

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-network-switch-access-policy
*/
func (s *SwitchService) DeleteNetworkSwitchAccessPolicy(networkID string, accessPolicyNumber string) (*resty.Response, error) {
	//networkID string,accessPolicyNumber string
	path := "/api/v1/networks/{networkId}/switch/accessPolicies/{accessPolicyNumber}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{accessPolicyNumber}", fmt.Sprintf("%v", accessPolicyNumber), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteNetworkSwitchAccessPolicy")
	}

	return response, err

}

//DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer Remove a server from being trusted by Dynamic ARP Inspection on this network
/* Remove a server from being trusted by Dynamic ARP Inspection on this network

@param networkID networkId path parameter. Network ID
@param trustedServerID trustedServerId path parameter. Trusted server ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-network-switch-dhcp-server-policy-arp-inspection-trusted-server
*/
func (s *SwitchService) DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(networkID string, trustedServerID string) (*resty.Response, error) {
	//networkID string,trustedServerID string
	path := "/api/v1/networks/{networkId}/switch/dhcpServerPolicy/arpInspection/trustedServers/{trustedServerId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{trustedServerId}", fmt.Sprintf("%v", trustedServerID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer")
	}

	return response, err

}

//DeleteNetworkSwitchLinkAggregation Split a link aggregation group into separate ports
/* Split a link aggregation group into separate ports

@param networkID networkId path parameter. Network ID
@param linkAggregationID linkAggregationId path parameter. Link aggregation ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-network-switch-link-aggregation
*/
func (s *SwitchService) DeleteNetworkSwitchLinkAggregation(networkID string, linkAggregationID string) (*resty.Response, error) {
	//networkID string,linkAggregationID string
	path := "/api/v1/networks/{networkId}/switch/linkAggregations/{linkAggregationId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{linkAggregationId}", fmt.Sprintf("%v", linkAggregationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteNetworkSwitchLinkAggregation")
	}

	return response, err

}

//DeleteNetworkSwitchPortSchedule Delete a switch port schedule
/* Delete a switch port schedule

@param networkID networkId path parameter. Network ID
@param portScheduleID portScheduleId path parameter. Port schedule ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-network-switch-port-schedule
*/
func (s *SwitchService) DeleteNetworkSwitchPortSchedule(networkID string, portScheduleID string) (*resty.Response, error) {
	//networkID string,portScheduleID string
	path := "/api/v1/networks/{networkId}/switch/portSchedules/{portScheduleId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{portScheduleId}", fmt.Sprintf("%v", portScheduleID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteNetworkSwitchPortSchedule")
	}

	return response, err

}

//DeleteNetworkSwitchQosRule Delete a quality of service rule
/* Delete a quality of service rule

@param networkID networkId path parameter. Network ID
@param qosRuleID qosRuleId path parameter. Qos rule ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-network-switch-qos-rule
*/
func (s *SwitchService) DeleteNetworkSwitchQosRule(networkID string, qosRuleID string) (*resty.Response, error) {
	//networkID string,qosRuleID string
	path := "/api/v1/networks/{networkId}/switch/qosRules/{qosRuleId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{qosRuleId}", fmt.Sprintf("%v", qosRuleID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteNetworkSwitchQosRule")
	}

	return response, err

}

//DeleteNetworkSwitchRoutingMulticastRendezvousPoint Delete a multicast rendezvous point
/* Delete a multicast rendezvous point

@param networkID networkId path parameter. Network ID
@param rendezvousPointID rendezvousPointId path parameter. Rendezvous point ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-network-switch-routing-multicast-rendezvous-point
*/
func (s *SwitchService) DeleteNetworkSwitchRoutingMulticastRendezvousPoint(networkID string, rendezvousPointID string) (*resty.Response, error) {
	//networkID string,rendezvousPointID string
	path := "/api/v1/networks/{networkId}/switch/routing/multicast/rendezvousPoints/{rendezvousPointId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{rendezvousPointId}", fmt.Sprintf("%v", rendezvousPointID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteNetworkSwitchRoutingMulticastRendezvousPoint")
	}

	return response, err

}

//DeleteNetworkSwitchStack Delete a stack
/* Delete a stack

@param networkID networkId path parameter. Network ID
@param switchStackID switchStackId path parameter. Switch stack ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-network-switch-stack
*/
func (s *SwitchService) DeleteNetworkSwitchStack(networkID string, switchStackID string) (*resty.Response, error) {
	//networkID string,switchStackID string
	path := "/api/v1/networks/{networkId}/switch/stacks/{switchStackId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{switchStackId}", fmt.Sprintf("%v", switchStackID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteNetworkSwitchStack")
	}

	return response, err

}

//DeleteNetworkSwitchStackRoutingInterface Delete a layer 3 interface from a switch stack
/* Delete a layer 3 interface from a switch stack

@param networkID networkId path parameter. Network ID
@param switchStackID switchStackId path parameter. Switch stack ID
@param interfaceID interfaceId path parameter. Interface ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-network-switch-stack-routing-interface
*/
func (s *SwitchService) DeleteNetworkSwitchStackRoutingInterface(networkID string, switchStackID string, interfaceID string) (*resty.Response, error) {
	//networkID string,switchStackID string,interfaceID string
	path := "/api/v1/networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{switchStackId}", fmt.Sprintf("%v", switchStackID), -1)
	path = strings.Replace(path, "{interfaceId}", fmt.Sprintf("%v", interfaceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteNetworkSwitchStackRoutingInterface")
	}

	return response, err

}

//DeleteNetworkSwitchStackRoutingStaticRoute Delete a layer 3 static route for a switch stack
/* Delete a layer 3 static route for a switch stack

@param networkID networkId path parameter. Network ID
@param switchStackID switchStackId path parameter. Switch stack ID
@param staticRouteID staticRouteId path parameter. Static route ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-network-switch-stack-routing-static-route
*/
func (s *SwitchService) DeleteNetworkSwitchStackRoutingStaticRoute(networkID string, switchStackID string, staticRouteID string) (*resty.Response, error) {
	//networkID string,switchStackID string,staticRouteID string
	path := "/api/v1/networks/{networkId}/switch/stacks/{switchStackId}/routing/staticRoutes/{staticRouteId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{switchStackId}", fmt.Sprintf("%v", switchStackID), -1)
	path = strings.Replace(path, "{staticRouteId}", fmt.Sprintf("%v", staticRouteID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteNetworkSwitchStackRoutingStaticRoute")
	}

	return response, err

}
