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
type GetOrganizationSummarySwitchPowerHistoryQueryParams struct {
	T0       string  `url:"t0,omitempty"`       //The beginning of the timespan for the data.
	T1       string  `url:"t1,omitempty"`       //The end of the timespan for the data. t1 can be a maximum of 186 days after t0.
	Timespan float64 `url:"timespan,omitempty"` //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 186 days. The default is 1 day.
}
type GetOrganizationSwitchPortsBySwitchQueryParams struct {
	PerPage                   int      `url:"perPage,omitempty"`                   //The number of entries per page returned. Acceptable range is 3 - 50. Default is 50.
	StartingAfter             string   `url:"startingAfter,omitempty"`             //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore              string   `url:"endingBefore,omitempty"`              //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	ConfigurationUpdatedAfter string   `url:"configurationUpdatedAfter,omitempty"` //Optional parameter to filter items to switches where the configuration has been updated after the given timestamp.
	Mac                       string   `url:"mac,omitempty"`                       //Optional parameter to filter items to switches with MAC addresses that contain the search term or are an exact match.
	Macs                      []string `url:"macs[],omitempty"`                    //Optional parameter to filter items to switches that have one of the provided MAC addresses.
	Name                      string   `url:"name,omitempty"`                      //Optional parameter to filter items to switches with names that contain the search term or are an exact match.
	NetworkIDs                []string `url:"networkIds[],omitempty"`              //Optional parameter to filter items to switches in one of the provided networks.
	PortProfileIDs            []string `url:"portProfileIds[],omitempty"`          //Optional parameter to filter items to switches that contain switchports belonging to one of the specified port profiles.
	Serial                    string   `url:"serial,omitempty"`                    //Optional parameter to filter items to switches with serial number that contains the search term or are an exact match.
	Serials                   []string `url:"serials[],omitempty"`                 //Optional parameter to filter items to switches that have one of the provided serials.
}
type GetOrganizationSwitchPortsClientsOverviewByDeviceQueryParams struct {
	T0                        string   `url:"t0,omitempty"`                        //The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	Timespan                  float64  `url:"timespan,omitempty"`                  //The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
	PerPage                   int      `url:"perPage,omitempty"`                   //The number of entries per page returned. Acceptable range is 3 - 20. Default is 20.
	StartingAfter             string   `url:"startingAfter,omitempty"`             //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore              string   `url:"endingBefore,omitempty"`              //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	ConfigurationUpdatedAfter string   `url:"configurationUpdatedAfter,omitempty"` //Optional parameter to filter items to switches where the configuration has been updated after the given timestamp.
	Mac                       string   `url:"mac,omitempty"`                       //Optional parameter to filter items to switches with MAC addresses that contain the search term or are an exact match.
	Macs                      []string `url:"macs[],omitempty"`                    //Optional parameter to filter items to switches that have one of the provided MAC addresses.
	Name                      string   `url:"name,omitempty"`                      //Optional parameter to filter items to switches with names that contain the search term or are an exact match.
	NetworkIDs                []string `url:"networkIds[],omitempty"`              //Optional parameter to filter items to switches in one of the provided networks.
	PortProfileIDs            []string `url:"portProfileIds[],omitempty"`          //Optional parameter to filter items to switches that contain switchports belonging to one of the specified port profiles.
	Serial                    string   `url:"serial,omitempty"`                    //Optional parameter to filter items to switches with serial number that contains the search term or are an exact match.
	Serials                   []string `url:"serials[],omitempty"`                 //Optional parameter to filter items to switches that have one of the provided serials.
}
type GetOrganizationSwitchPortsOverviewQueryParams struct {
	T0       string  `url:"t0,omitempty"`       //The beginning of the timespan for the data.
	T1       string  `url:"t1,omitempty"`       //The end of the timespan for the data. t1 can be a maximum of 186 days after t0.
	Timespan float64 `url:"timespan,omitempty"` //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 12 hours and be less than or equal to 186 days. The default is 1 day.
}
type GetOrganizationSwitchPortsStatusesBySwitchQueryParams struct {
	PerPage                   int      `url:"perPage,omitempty"`                   //The number of entries per page returned. Acceptable range is 3 - 20. Default is 10.
	StartingAfter             string   `url:"startingAfter,omitempty"`             //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore              string   `url:"endingBefore,omitempty"`              //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	ConfigurationUpdatedAfter string   `url:"configurationUpdatedAfter,omitempty"` //Optional parameter to filter items to switches where the configuration has been updated after the given timestamp.
	Mac                       string   `url:"mac,omitempty"`                       //Optional parameter to filter items to switches with MAC addresses that contain the search term or are an exact match.
	Macs                      []string `url:"macs[],omitempty"`                    //Optional parameter to filter items to switches that have one of the provided MAC addresses.
	Name                      string   `url:"name,omitempty"`                      //Optional parameter to filter items to switches with names that contain the search term or are an exact match.
	NetworkIDs                []string `url:"networkIds[],omitempty"`              //Optional parameter to filter items to switches in one of the provided networks.
	PortProfileIDs            []string `url:"portProfileIds[],omitempty"`          //Optional parameter to filter items to switches that contain switchports belonging to one of the specified port profiles.
	Serial                    string   `url:"serial,omitempty"`                    //Optional parameter to filter items to switches with serial number that contains the search term or are an exact match.
	Serials                   []string `url:"serials[],omitempty"`                 //Optional parameter to filter items to switches that have one of the provided serials.
}
type GetOrganizationSwitchPortsTopologyDiscoveryByDeviceQueryParams struct {
	T0                        string   `url:"t0,omitempty"`                        //The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	Timespan                  float64  `url:"timespan,omitempty"`                  //The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
	PerPage                   int      `url:"perPage,omitempty"`                   //The number of entries per page returned. Acceptable range is 3 - 20. Default is 10.
	StartingAfter             string   `url:"startingAfter,omitempty"`             //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore              string   `url:"endingBefore,omitempty"`              //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	ConfigurationUpdatedAfter string   `url:"configurationUpdatedAfter,omitempty"` //Optional parameter to filter items to switches where the configuration has been updated after the given timestamp.
	Mac                       string   `url:"mac,omitempty"`                       //Optional parameter to filter items to switches with MAC addresses that contain the search term or are an exact match.
	Macs                      []string `url:"macs[],omitempty"`                    //Optional parameter to filter items to switches that have one of the provided MAC addresses.
	Name                      string   `url:"name,omitempty"`                      //Optional parameter to filter items to switches with names that contain the search term or are an exact match.
	NetworkIDs                []string `url:"networkIds[],omitempty"`              //Optional parameter to filter items to switches in one of the provided networks.
	PortProfileIDs            []string `url:"portProfileIds[],omitempty"`          //Optional parameter to filter items to switches that contain switchports belonging to one of the specified port profiles.
	Serial                    string   `url:"serial,omitempty"`                    //Optional parameter to filter items to switches with serial number that contains the search term or are an exact match.
	Serials                   []string `url:"serials[],omitempty"`                 //Optional parameter to filter items to switches that have one of the provided serials.
}

type ResponseSwitchGetDeviceSwitchPorts []ResponseItemSwitchGetDeviceSwitchPorts // Array of ResponseSwitchGetDeviceSwitchPorts
type ResponseItemSwitchGetDeviceSwitchPorts struct {
	AccessPolicyNumber          *int                                                       `json:"accessPolicyNumber,omitempty"`          // The number of a custom access policy to configure on the switch port. Only applicable when 'accessPolicyType' is 'Custom access policy'.
	AccessPolicyType            string                                                     `json:"accessPolicyType,omitempty"`            // The type of the access policy of the switch port. Only applicable to access ports. Can be one of 'Open', 'Custom access policy', 'MAC allow list' or 'Sticky MAC allow list'.
	AdaptivePolicyGroup         *ResponseItemSwitchGetDeviceSwitchPortsAdaptivePolicyGroup `json:"adaptivePolicyGroup,omitempty"`         // The adaptive policy group data of the port.
	AdaptivePolicyGroupID       string                                                     `json:"adaptivePolicyGroupId,omitempty"`       // The adaptive policy group ID that will be used to tag traffic through this switch port. This ID must pre-exist during the configuration, else needs to be created using adaptivePolicy/groups API. Cannot be applied to a port on a switch bound to profile.
	AllowedVLANs                string                                                     `json:"allowedVlans,omitempty"`                // The VLANs allowed on the switch port. Only applicable to trunk ports.
	DaiTrusted                  *bool                                                      `json:"daiTrusted,omitempty"`                  // If true, ARP packets for this port will be considered trusted, and Dynamic ARP Inspection will allow the traffic.
	Dot3Az                      *ResponseItemSwitchGetDeviceSwitchPortsDot3Az              `json:"dot3az,omitempty"`                      // dot3az settings for the port
	Enabled                     *bool                                                      `json:"enabled,omitempty"`                     // The status of the switch port.
	FlexibleStackingEnabled     *bool                                                      `json:"flexibleStackingEnabled,omitempty"`     // For supported switches (e.g. MS420/MS425), whether or not the port has flexible stacking enabled.
	IsolationEnabled            *bool                                                      `json:"isolationEnabled,omitempty"`            // The isolation status of the switch port.
	LinkNegotiation             string                                                     `json:"linkNegotiation,omitempty"`             // The link speed for the switch port.
	LinkNegotiationCapabilities []string                                                   `json:"linkNegotiationCapabilities,omitempty"` // Available link speeds for the switch port.
	MacAllowList                []string                                                   `json:"macAllowList,omitempty"`                // Only devices with MAC addresses specified in this list will have access to this port. Up to 20 MAC addresses can be defined. Only applicable when 'accessPolicyType' is 'MAC allow list'.
	Mirror                      *ResponseItemSwitchGetDeviceSwitchPortsMirror              `json:"mirror,omitempty"`                      // Port mirror
	Module                      *ResponseItemSwitchGetDeviceSwitchPortsModule              `json:"module,omitempty"`                      // Expansion module
	Name                        string                                                     `json:"name,omitempty"`                        // The name of the switch port.
	PeerSgtCapable              *bool                                                      `json:"peerSgtCapable,omitempty"`              // If true, Peer SGT is enabled for traffic through this switch port. Applicable to trunk port only, not access port. Cannot be applied to a port on a switch bound to profile.
	PoeEnabled                  *bool                                                      `json:"poeEnabled,omitempty"`                  // The PoE status of the switch port.
	PortID                      string                                                     `json:"portId,omitempty"`                      // The identifier of the switch port.
	PortScheduleID              string                                                     `json:"portScheduleId,omitempty"`              // The ID of the port schedule. A value of null will clear the port schedule.
	Profile                     *ResponseItemSwitchGetDeviceSwitchPortsProfile             `json:"profile,omitempty"`                     // Profile attributes
	RstpEnabled                 *bool                                                      `json:"rstpEnabled,omitempty"`                 // The rapid spanning tree protocol status.
	Schedule                    *ResponseItemSwitchGetDeviceSwitchPortsSchedule            `json:"schedule,omitempty"`                    // The port schedule data.
	StackwiseVirtual            *ResponseItemSwitchGetDeviceSwitchPortsStackwiseVirtual    `json:"stackwiseVirtual,omitempty"`            // Stackwise Virtual settings for the port
	StickyMacAllowList          []string                                                   `json:"stickyMacAllowList,omitempty"`          // The initial list of MAC addresses for sticky Mac allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
	StickyMacAllowListLimit     *int                                                       `json:"stickyMacAllowListLimit,omitempty"`     // The maximum number of MAC addresses for sticky MAC allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
	StormControlEnabled         *bool                                                      `json:"stormControlEnabled,omitempty"`         // The storm control status of the switch port.
	StpGuard                    string                                                     `json:"stpGuard,omitempty"`                    // The state of the STP guard ('disabled', 'root guard', 'bpdu guard' or 'loop guard').
	Tags                        []string                                                   `json:"tags,omitempty"`                        // The list of tags of the switch port.
	Type                        string                                                     `json:"type,omitempty"`                        // The type of the switch port ('trunk', 'access' or 'stack').
	Udld                        string                                                     `json:"udld,omitempty"`                        // The action to take when Unidirectional Link is detected (Alert only, Enforce). Default configuration is Alert only.
	VLAN                        *int                                                       `json:"vlan,omitempty"`                        // The VLAN of the switch port. For a trunk port, this is the native VLAN. A null value will clear the value set for trunk ports.
	VoiceVLAN                   *int                                                       `json:"voiceVlan,omitempty"`                   // The voice VLAN of the switch port. Only applicable to access ports.
}
type ResponseItemSwitchGetDeviceSwitchPortsAdaptivePolicyGroup struct {
	ID   string `json:"id,omitempty"`   // The ID of the adaptive policy group.
	Name string `json:"name,omitempty"` // The name of the adaptive policy group.
}
type ResponseItemSwitchGetDeviceSwitchPortsDot3Az struct {
	Enabled *bool `json:"enabled,omitempty"` // The Energy Efficient Ethernet status of the switch port.
}
type ResponseItemSwitchGetDeviceSwitchPortsMirror struct {
	Mode string `json:"mode,omitempty"` // The port mirror mode. Can be one of ('Destination port', 'Source port' or 'Not mirroring traffic').
}
type ResponseItemSwitchGetDeviceSwitchPortsModule struct {
	Model string `json:"model,omitempty"` // The model of the expansion module.
}
type ResponseItemSwitchGetDeviceSwitchPortsProfile struct {
	Enabled *bool  `json:"enabled,omitempty"` // When enabled, override this port's configuration with a port profile.
	ID      string `json:"id,omitempty"`      // When enabled, the ID of the port profile used to override the port's configuration.
	Iname   string `json:"iname,omitempty"`   // When enabled, the IName of the profile.
}
type ResponseItemSwitchGetDeviceSwitchPortsSchedule struct {
	ID   string `json:"id,omitempty"`   // The ID of the port schedule.
	Name string `json:"name,omitempty"` // The name of the port schedule.
}
type ResponseItemSwitchGetDeviceSwitchPortsStackwiseVirtual struct {
	IsDualActiveDetector   *bool `json:"isDualActiveDetector,omitempty"`   // For SVL devices, whether or not the port is used for Dual Active Detection.
	IsStackWiseVirtualLink *bool `json:"isStackWiseVirtualLink,omitempty"` // For SVL devices, whether or not the port is used for StackWise Virtual Link.
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
	Poe            *ResponseItemSwitchGetDeviceSwitchPortsStatusesPoe           `json:"poe,omitempty"`            // PoE status of the port.
	PortID         string                                                       `json:"portId,omitempty"`         // The string identifier of this port on the switch. This is commonly just the port number but may contain additional identifying information such as the slot and module-type if the port is located on a port module.
	PowerUsageInWh *float64                                                     `json:"powerUsageInWh,omitempty"` // How much power (in watt-hours) has been delivered by this port during the timespan.
	SecurePort     *ResponseItemSwitchGetDeviceSwitchPortsStatusesSecurePort    `json:"securePort,omitempty"`     // The Secure Port status of the port.
	SpanningTree   *ResponseItemSwitchGetDeviceSwitchPortsStatusesSpanningTree  `json:"spanningTree,omitempty"`   // The Spanning Tree Protocol (STP) information of the connected device.
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
type ResponseItemSwitchGetDeviceSwitchPortsStatusesPoe struct {
	IsAllocated *bool `json:"isAllocated,omitempty"` // Whether the port is drawing power
}
type ResponseItemSwitchGetDeviceSwitchPortsStatusesSecurePort struct {
	Active               *bool                                                                    `json:"active,omitempty"`               // Whether Secure Port is currently active for this port.
	AuthenticationStatus string                                                                   `json:"authenticationStatus,omitempty"` // The current Secure Port status.
	ConfigOverrides      *ResponseItemSwitchGetDeviceSwitchPortsStatusesSecurePortConfigOverrides `json:"configOverrides,omitempty"`      // The configuration overrides applied to this port when Secure Port is active.
	Enabled              *bool                                                                    `json:"enabled,omitempty"`              // Whether Secure Port is turned on for this port.
}
type ResponseItemSwitchGetDeviceSwitchPortsStatusesSecurePortConfigOverrides struct {
	AllowedVLANs string `json:"allowedVlans,omitempty"` // The VLANs allowed on the . Only applicable to trunk ports.
	Type         string `json:"type,omitempty"`         // The type of the  ('trunk', 'access' or 'stack').
	VLAN         *int   `json:"vlan,omitempty"`         // The VLAN of the . For a trunk port, this is the native VLAN. A null value will clear the value set for trunk ports.
	VoiceVLAN    *int   `json:"voiceVlan,omitempty"`    // The voice VLAN of the . Only applicable to access ports.
}
type ResponseItemSwitchGetDeviceSwitchPortsStatusesSpanningTree struct {
	Statuses []string `json:"statuses,omitempty"` // The current Spanning Tree Protocol statuses of the port.
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
	Packets *[]ResponseItemSwitchGetDeviceSwitchPortsStatusesPacketsPackets `json:"packets,omitempty"` // The packet counts on the switch.
	PortID  string                                                          `json:"portId,omitempty"`  // The string identifier of this port on the switch. This is commonly just the port number but may contain additional identifying information such as the slot and module-type if the port is located on a port module.
}
type ResponseItemSwitchGetDeviceSwitchPortsStatusesPacketsPackets struct {
	Desc       string                                                                  `json:"desc,omitempty"`       // The type of packets being counted.
	RatePerSec *ResponseItemSwitchGetDeviceSwitchPortsStatusesPacketsPacketsRatePerSec `json:"ratePerSec,omitempty"` // Packet rates measured in packets per second.
	Recv       *int                                                                    `json:"recv,omitempty"`       // The total count of packets received by the switch during the timespan.
	Sent       *int                                                                    `json:"sent,omitempty"`       // The total count of packets sent by the switch during the timespan.
	Total      *int                                                                    `json:"total,omitempty"`      // The total count of sent and received packets.
}
type ResponseItemSwitchGetDeviceSwitchPortsStatusesPacketsPacketsRatePerSec struct {
	Recv  *int `json:"recv,omitempty"`  // The rate of packets received during the timespan
	Sent  *int `json:"sent,omitempty"`  // The rate of packets sent during the timespan
	Total *int `json:"total,omitempty"` // The rate of all packets sent and received during the timespan
}
type ResponseSwitchGetDeviceSwitchPort struct {
	AccessPolicyNumber          *int                                                  `json:"accessPolicyNumber,omitempty"`          // The number of a custom access policy to configure on the switch port. Only applicable when 'accessPolicyType' is 'Custom access policy'.
	AccessPolicyType            string                                                `json:"accessPolicyType,omitempty"`            // The type of the access policy of the switch port. Only applicable to access ports. Can be one of 'Open', 'Custom access policy', 'MAC allow list' or 'Sticky MAC allow list'.
	AdaptivePolicyGroup         *ResponseSwitchGetDeviceSwitchPortAdaptivePolicyGroup `json:"adaptivePolicyGroup,omitempty"`         // The adaptive policy group data of the port.
	AdaptivePolicyGroupID       string                                                `json:"adaptivePolicyGroupId,omitempty"`       // The adaptive policy group ID that will be used to tag traffic through this switch port. This ID must pre-exist during the configuration, else needs to be created using adaptivePolicy/groups API. Cannot be applied to a port on a switch bound to profile.
	AllowedVLANs                string                                                `json:"allowedVlans,omitempty"`                // The VLANs allowed on the switch port. Only applicable to trunk ports.
	DaiTrusted                  *bool                                                 `json:"daiTrusted,omitempty"`                  // If true, ARP packets for this port will be considered trusted, and Dynamic ARP Inspection will allow the traffic.
	Dot3Az                      *ResponseSwitchGetDeviceSwitchPortDot3Az              `json:"dot3az,omitempty"`                      // dot3az settings for the port
	Enabled                     *bool                                                 `json:"enabled,omitempty"`                     // The status of the switch port.
	FlexibleStackingEnabled     *bool                                                 `json:"flexibleStackingEnabled,omitempty"`     // For supported switches (e.g. MS420/MS425), whether or not the port has flexible stacking enabled.
	IsolationEnabled            *bool                                                 `json:"isolationEnabled,omitempty"`            // The isolation status of the switch port.
	LinkNegotiation             string                                                `json:"linkNegotiation,omitempty"`             // The link speed for the switch port.
	LinkNegotiationCapabilities []string                                              `json:"linkNegotiationCapabilities,omitempty"` // Available link speeds for the switch port.
	MacAllowList                []string                                              `json:"macAllowList,omitempty"`                // Only devices with MAC addresses specified in this list will have access to this port. Up to 20 MAC addresses can be defined. Only applicable when 'accessPolicyType' is 'MAC allow list'.
	Mirror                      *ResponseSwitchGetDeviceSwitchPortMirror              `json:"mirror,omitempty"`                      // Port mirror
	Module                      *ResponseSwitchGetDeviceSwitchPortModule              `json:"module,omitempty"`                      // Expansion module
	Name                        string                                                `json:"name,omitempty"`                        // The name of the switch port.
	PeerSgtCapable              *bool                                                 `json:"peerSgtCapable,omitempty"`              // If true, Peer SGT is enabled for traffic through this switch port. Applicable to trunk port only, not access port. Cannot be applied to a port on a switch bound to profile.
	PoeEnabled                  *bool                                                 `json:"poeEnabled,omitempty"`                  // The PoE status of the switch port.
	PortID                      string                                                `json:"portId,omitempty"`                      // The identifier of the switch port.
	PortScheduleID              string                                                `json:"portScheduleId,omitempty"`              // The ID of the port schedule. A value of null will clear the port schedule.
	Profile                     *ResponseSwitchGetDeviceSwitchPortProfile             `json:"profile,omitempty"`                     // Profile attributes
	RstpEnabled                 *bool                                                 `json:"rstpEnabled,omitempty"`                 // The rapid spanning tree protocol status.
	Schedule                    *ResponseSwitchGetDeviceSwitchPortSchedule            `json:"schedule,omitempty"`                    // The port schedule data.
	StackwiseVirtual            *ResponseSwitchGetDeviceSwitchPortStackwiseVirtual    `json:"stackwiseVirtual,omitempty"`            // Stackwise Virtual settings for the port
	StickyMacAllowList          []string                                              `json:"stickyMacAllowList,omitempty"`          // The initial list of MAC addresses for sticky Mac allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
	StickyMacAllowListLimit     *int                                                  `json:"stickyMacAllowListLimit,omitempty"`     // The maximum number of MAC addresses for sticky MAC allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
	StormControlEnabled         *bool                                                 `json:"stormControlEnabled,omitempty"`         // The storm control status of the switch port.
	StpGuard                    string                                                `json:"stpGuard,omitempty"`                    // The state of the STP guard ('disabled', 'root guard', 'bpdu guard' or 'loop guard').
	Tags                        []string                                              `json:"tags,omitempty"`                        // The list of tags of the switch port.
	Type                        string                                                `json:"type,omitempty"`                        // The type of the switch port ('trunk', 'access' or 'stack').
	Udld                        string                                                `json:"udld,omitempty"`                        // The action to take when Unidirectional Link is detected (Alert only, Enforce). Default configuration is Alert only.
	VLAN                        *int                                                  `json:"vlan,omitempty"`                        // The VLAN of the switch port. For a trunk port, this is the native VLAN. A null value will clear the value set for trunk ports.
	VoiceVLAN                   *int                                                  `json:"voiceVlan,omitempty"`                   // The voice VLAN of the switch port. Only applicable to access ports.
}
type ResponseSwitchGetDeviceSwitchPortAdaptivePolicyGroup struct {
	ID   string `json:"id,omitempty"`   // The ID of the adaptive policy group.
	Name string `json:"name,omitempty"` // The name of the adaptive policy group.
}
type ResponseSwitchGetDeviceSwitchPortDot3Az struct {
	Enabled *bool `json:"enabled,omitempty"` // The Energy Efficient Ethernet status of the switch port.
}
type ResponseSwitchGetDeviceSwitchPortMirror struct {
	Mode string `json:"mode,omitempty"` // The port mirror mode. Can be one of ('Destination port', 'Source port' or 'Not mirroring traffic').
}
type ResponseSwitchGetDeviceSwitchPortModule struct {
	Model string `json:"model,omitempty"` // The model of the expansion module.
}
type ResponseSwitchGetDeviceSwitchPortProfile struct {
	Enabled *bool  `json:"enabled,omitempty"` // When enabled, override this port's configuration with a port profile.
	ID      string `json:"id,omitempty"`      // When enabled, the ID of the port profile used to override the port's configuration.
	Iname   string `json:"iname,omitempty"`   // When enabled, the IName of the profile.
}
type ResponseSwitchGetDeviceSwitchPortSchedule struct {
	ID   string `json:"id,omitempty"`   // The ID of the port schedule.
	Name string `json:"name,omitempty"` // The name of the port schedule.
}
type ResponseSwitchGetDeviceSwitchPortStackwiseVirtual struct {
	IsDualActiveDetector   *bool `json:"isDualActiveDetector,omitempty"`   // For SVL devices, whether or not the port is used for Dual Active Detection.
	IsStackWiseVirtualLink *bool `json:"isStackWiseVirtualLink,omitempty"` // For SVL devices, whether or not the port is used for StackWise Virtual Link.
}
type ResponseSwitchUpdateDeviceSwitchPort struct {
	AccessPolicyNumber          *int                                                     `json:"accessPolicyNumber,omitempty"`          // The number of a custom access policy to configure on the switch port. Only applicable when 'accessPolicyType' is 'Custom access policy'.
	AccessPolicyType            string                                                   `json:"accessPolicyType,omitempty"`            // The type of the access policy of the switch port. Only applicable to access ports. Can be one of 'Open', 'Custom access policy', 'MAC allow list' or 'Sticky MAC allow list'.
	AdaptivePolicyGroup         *ResponseSwitchUpdateDeviceSwitchPortAdaptivePolicyGroup `json:"adaptivePolicyGroup,omitempty"`         // The adaptive policy group data of the port.
	AdaptivePolicyGroupID       string                                                   `json:"adaptivePolicyGroupId,omitempty"`       // The adaptive policy group ID that will be used to tag traffic through this switch port. This ID must pre-exist during the configuration, else needs to be created using adaptivePolicy/groups API. Cannot be applied to a port on a switch bound to profile.
	AllowedVLANs                string                                                   `json:"allowedVlans,omitempty"`                // The VLANs allowed on the switch port. Only applicable to trunk ports.
	DaiTrusted                  *bool                                                    `json:"daiTrusted,omitempty"`                  // If true, ARP packets for this port will be considered trusted, and Dynamic ARP Inspection will allow the traffic.
	Dot3Az                      *ResponseSwitchUpdateDeviceSwitchPortDot3Az              `json:"dot3az,omitempty"`                      // dot3az settings for the port
	Enabled                     *bool                                                    `json:"enabled,omitempty"`                     // The status of the switch port.
	FlexibleStackingEnabled     *bool                                                    `json:"flexibleStackingEnabled,omitempty"`     // For supported switches (e.g. MS420/MS425), whether or not the port has flexible stacking enabled.
	IsolationEnabled            *bool                                                    `json:"isolationEnabled,omitempty"`            // The isolation status of the switch port.
	LinkNegotiation             string                                                   `json:"linkNegotiation,omitempty"`             // The link speed for the switch port.
	LinkNegotiationCapabilities []string                                                 `json:"linkNegotiationCapabilities,omitempty"` // Available link speeds for the switch port.
	MacAllowList                []string                                                 `json:"macAllowList,omitempty"`                // Only devices with MAC addresses specified in this list will have access to this port. Up to 20 MAC addresses can be defined. Only applicable when 'accessPolicyType' is 'MAC allow list'.
	Mirror                      *ResponseSwitchUpdateDeviceSwitchPortMirror              `json:"mirror,omitempty"`                      // Port mirror
	Module                      *ResponseSwitchUpdateDeviceSwitchPortModule              `json:"module,omitempty"`                      // Expansion module
	Name                        string                                                   `json:"name,omitempty"`                        // The name of the switch port.
	PeerSgtCapable              *bool                                                    `json:"peerSgtCapable,omitempty"`              // If true, Peer SGT is enabled for traffic through this switch port. Applicable to trunk port only, not access port. Cannot be applied to a port on a switch bound to profile.
	PoeEnabled                  *bool                                                    `json:"poeEnabled,omitempty"`                  // The PoE status of the switch port.
	PortID                      string                                                   `json:"portId,omitempty"`                      // The identifier of the switch port.
	PortScheduleID              string                                                   `json:"portScheduleId,omitempty"`              // The ID of the port schedule. A value of null will clear the port schedule.
	Profile                     *ResponseSwitchUpdateDeviceSwitchPortProfile             `json:"profile,omitempty"`                     // Profile attributes
	RstpEnabled                 *bool                                                    `json:"rstpEnabled,omitempty"`                 // The rapid spanning tree protocol status.
	Schedule                    *ResponseSwitchUpdateDeviceSwitchPortSchedule            `json:"schedule,omitempty"`                    // The port schedule data.
	StackwiseVirtual            *ResponseSwitchUpdateDeviceSwitchPortStackwiseVirtual    `json:"stackwiseVirtual,omitempty"`            // Stackwise Virtual settings for the port
	StickyMacAllowList          []string                                                 `json:"stickyMacAllowList,omitempty"`          // The initial list of MAC addresses for sticky Mac allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
	StickyMacAllowListLimit     *int                                                     `json:"stickyMacAllowListLimit,omitempty"`     // The maximum number of MAC addresses for sticky MAC allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
	StormControlEnabled         *bool                                                    `json:"stormControlEnabled,omitempty"`         // The storm control status of the switch port.
	StpGuard                    string                                                   `json:"stpGuard,omitempty"`                    // The state of the STP guard ('disabled', 'root guard', 'bpdu guard' or 'loop guard').
	Tags                        []string                                                 `json:"tags,omitempty"`                        // The list of tags of the switch port.
	Type                        string                                                   `json:"type,omitempty"`                        // The type of the switch port ('trunk', 'access' or 'stack').
	Udld                        string                                                   `json:"udld,omitempty"`                        // The action to take when Unidirectional Link is detected (Alert only, Enforce). Default configuration is Alert only.
	VLAN                        *int                                                     `json:"vlan,omitempty"`                        // The VLAN of the switch port. For a trunk port, this is the native VLAN. A null value will clear the value set for trunk ports.
	VoiceVLAN                   *int                                                     `json:"voiceVlan,omitempty"`                   // The voice VLAN of the switch port. Only applicable to access ports.
}
type ResponseSwitchUpdateDeviceSwitchPortAdaptivePolicyGroup struct {
	ID   string `json:"id,omitempty"`   // The ID of the adaptive policy group.
	Name string `json:"name,omitempty"` // The name of the adaptive policy group.
}
type ResponseSwitchUpdateDeviceSwitchPortDot3Az struct {
	Enabled *bool `json:"enabled,omitempty"` // The Energy Efficient Ethernet status of the switch port.
}
type ResponseSwitchUpdateDeviceSwitchPortMirror struct {
	Mode string `json:"mode,omitempty"` // The port mirror mode. Can be one of ('Destination port', 'Source port' or 'Not mirroring traffic').
}
type ResponseSwitchUpdateDeviceSwitchPortModule struct {
	Model string `json:"model,omitempty"` // The model of the expansion module.
}
type ResponseSwitchUpdateDeviceSwitchPortProfile struct {
	Enabled *bool  `json:"enabled,omitempty"` // When enabled, override this port's configuration with a port profile.
	ID      string `json:"id,omitempty"`      // When enabled, the ID of the port profile used to override the port's configuration.
	Iname   string `json:"iname,omitempty"`   // When enabled, the IName of the profile.
}
type ResponseSwitchUpdateDeviceSwitchPortSchedule struct {
	ID   string `json:"id,omitempty"`   // The ID of the port schedule.
	Name string `json:"name,omitempty"` // The name of the port schedule.
}
type ResponseSwitchUpdateDeviceSwitchPortStackwiseVirtual struct {
	IsDualActiveDetector   *bool `json:"isDualActiveDetector,omitempty"`   // For SVL devices, whether or not the port is used for Dual Active Detection.
	IsStackWiseVirtualLink *bool `json:"isStackWiseVirtualLink,omitempty"` // For SVL devices, whether or not the port is used for StackWise Virtual Link.
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
	UplinkV4         *bool                                                           `json:"uplinkV4,omitempty"`         // Whether this is the switch's IPv4 uplink
	UplinkV6         *bool                                                           `json:"uplinkV6,omitempty"`         // Whether this is the switch's IPv6 uplink
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
	UplinkV4         *bool                                                         `json:"uplinkV4,omitempty"`         // Whether this is the switch's IPv4 uplink
	UplinkV6         *bool                                                         `json:"uplinkV6,omitempty"`         // Whether this is the switch's IPv6 uplink
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
	UplinkV4         *bool                                                      `json:"uplinkV4,omitempty"`         // Whether this is the switch's IPv4 uplink
	UplinkV6         *bool                                                      `json:"uplinkV6,omitempty"`         // Whether this is the switch's IPv6 uplink
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
	UplinkV4         *bool                                                         `json:"uplinkV4,omitempty"`         // Whether this is the switch's IPv4 uplink
	UplinkV6         *bool                                                         `json:"uplinkV6,omitempty"`         // Whether this is the switch's IPv6 uplink
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
	BootFileName         string                                                                 `json:"bootFileName,omitempty"`         // The PXE boot server file name for the DHCP server running on the switch stack interface
	BootNextServer       string                                                                 `json:"bootNextServer,omitempty"`       // The PXE boot server IP for the DHCP server running on the switch stack interface
	BootOptionsEnabled   *bool                                                                  `json:"bootOptionsEnabled,omitempty"`   // Enable DHCP boot options to provide PXE boot options configs for the dhcp server running on the switch stack interface
	DhcpLeaseTime        string                                                                 `json:"dhcpLeaseTime,omitempty"`        // The DHCP lease time config for the dhcp server running on the switch stack interface ('30 minutes', '1 hour', '4 hours', '12 hours', '1 day' or '1 week')
	DhcpMode             string                                                                 `json:"dhcpMode,omitempty"`             // The DHCP mode options for the switch stack interface ('dhcpDisabled', 'dhcpRelay' or 'dhcpServer')
	DhcpOptions          *[]ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpDhcpOptions        `json:"dhcpOptions,omitempty"`          // Array of DHCP options consisting of code, type and value for the DHCP server running on the switch stack interface
	DhcpRelayServerIPs   []string                                                               `json:"dhcpRelayServerIps,omitempty"`   // The DHCP relay server IPs to which DHCP packets would get relayed for the switch stack interface
	DNSCustomNameservers []string                                                               `json:"dnsCustomNameservers,omitempty"` // The DHCP name server IPs when DHCP name server option is 'custom'
	DNSNameserversOption string                                                                 `json:"dnsNameserversOption,omitempty"` // The DHCP name server option for the dhcp server running on the switch stack interface ('googlePublicDns', 'openDns' or 'custom')
	FixedIPAssignments   *[]ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpFixedIPAssignments `json:"fixedIpAssignments,omitempty"`   // Array of DHCP reserved IP assignments for the DHCP server running on the switch stack interface
	ReservedIPRanges     *[]ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpReservedIPRanges   `json:"reservedIpRanges,omitempty"`     // Array of DHCP reserved IP assignments for the DHCP server running on the switch stack interface
}
type ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpDhcpOptions struct {
	Code  string `json:"code,omitempty"`  // The code for DHCP option which should be from 2 to 254
	Type  string `json:"type,omitempty"`  // The type of the DHCP option which should be one of ('text', 'ip', 'integer' or 'hex')
	Value string `json:"value,omitempty"` // The value of the DHCP option
}
type ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpFixedIPAssignments struct {
	IP   string `json:"ip,omitempty"`   // The IP address of the client which has fixed IP address assigned to it
	Mac  string `json:"mac,omitempty"`  // The MAC address of the client which has fixed IP address
	Name string `json:"name,omitempty"` // The name of the client which has fixed IP address
}
type ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpReservedIPRanges struct {
	Comment string `json:"comment,omitempty"` // The comment for the reserved IP range
	End     string `json:"end,omitempty"`     // The ending IP address of the reserved IP range
	Start   string `json:"start,omitempty"`   // The starting IP address of the reserved IP range
}
type ResponseSwitchUpdateDeviceSwitchRoutingInterfaceDhcp struct {
	BootFileName         string                                                                    `json:"bootFileName,omitempty"`         // The PXE boot server file name for the DHCP server running on the switch stack interface
	BootNextServer       string                                                                    `json:"bootNextServer,omitempty"`       // The PXE boot server IP for the DHCP server running on the switch stack interface
	BootOptionsEnabled   *bool                                                                     `json:"bootOptionsEnabled,omitempty"`   // Enable DHCP boot options to provide PXE boot options configs for the dhcp server running on the switch stack interface
	DhcpLeaseTime        string                                                                    `json:"dhcpLeaseTime,omitempty"`        // The DHCP lease time config for the dhcp server running on the switch stack interface ('30 minutes', '1 hour', '4 hours', '12 hours', '1 day' or '1 week')
	DhcpMode             string                                                                    `json:"dhcpMode,omitempty"`             // The DHCP mode options for the switch stack interface ('dhcpDisabled', 'dhcpRelay' or 'dhcpServer')
	DhcpOptions          *[]ResponseSwitchUpdateDeviceSwitchRoutingInterfaceDhcpDhcpOptions        `json:"dhcpOptions,omitempty"`          // Array of DHCP options consisting of code, type and value for the DHCP server running on the switch stack interface
	DhcpRelayServerIPs   []string                                                                  `json:"dhcpRelayServerIps,omitempty"`   // The DHCP relay server IPs to which DHCP packets would get relayed for the switch stack interface
	DNSCustomNameservers []string                                                                  `json:"dnsCustomNameservers,omitempty"` // The DHCP name server IPs when DHCP name server option is 'custom'
	DNSNameserversOption string                                                                    `json:"dnsNameserversOption,omitempty"` // The DHCP name server option for the dhcp server running on the switch stack interface ('googlePublicDns', 'openDns' or 'custom')
	FixedIPAssignments   *[]ResponseSwitchUpdateDeviceSwitchRoutingInterfaceDhcpFixedIPAssignments `json:"fixedIpAssignments,omitempty"`   // Array of DHCP reserved IP assignments for the DHCP server running on the switch stack interface
	ReservedIPRanges     *[]ResponseSwitchUpdateDeviceSwitchRoutingInterfaceDhcpReservedIPRanges   `json:"reservedIpRanges,omitempty"`     // Array of DHCP reserved IP assignments for the DHCP server running on the switch stack interface
}
type ResponseSwitchUpdateDeviceSwitchRoutingInterfaceDhcpDhcpOptions struct {
	Code  string `json:"code,omitempty"`  // The code for DHCP option which should be from 2 to 254
	Type  string `json:"type,omitempty"`  // The type of the DHCP option which should be one of ('text', 'ip', 'integer' or 'hex')
	Value string `json:"value,omitempty"` // The value of the DHCP option
}
type ResponseSwitchUpdateDeviceSwitchRoutingInterfaceDhcpFixedIPAssignments struct {
	IP   string `json:"ip,omitempty"`   // The IP address of the client which has fixed IP address assigned to it
	Mac  string `json:"mac,omitempty"`  // The MAC address of the client which has fixed IP address
	Name string `json:"name,omitempty"` // The name of the client which has fixed IP address
}
type ResponseSwitchUpdateDeviceSwitchRoutingInterfaceDhcpReservedIPRanges struct {
	Comment string `json:"comment,omitempty"` // The comment for the reserved IP range
	End     string `json:"end,omitempty"`     // The ending IP address of the reserved IP range
	Start   string `json:"start,omitempty"`   // The starting IP address of the reserved IP range
}
type ResponseSwitchGetDeviceSwitchRoutingStaticRoutes []ResponseItemSwitchGetDeviceSwitchRoutingStaticRoutes // Array of ResponseSwitchGetDeviceSwitchRoutingStaticRoutes
type ResponseItemSwitchGetDeviceSwitchRoutingStaticRoutes struct {
	AdvertiseViaOspfEnabled     *bool  `json:"advertiseViaOspfEnabled,omitempty"`     // Option to advertise static routes via OSPF
	ManagementNextHop           string `json:"managementNextHop,omitempty"`           // Optional fallback IP address for management traffic
	Name                        string `json:"name,omitempty"`                        // The name or description of the layer 3 static route
	NextHopIP                   string `json:"nextHopIp,omitempty"`                   // The IP address of the router to which traffic for this destination network should be sent
	PreferOverOspfRoutesEnabled *bool  `json:"preferOverOspfRoutesEnabled,omitempty"` // Option to prefer static routes over OSPF routes
	StaticRouteID               string `json:"staticRouteId,omitempty"`               // The identifier of a layer 3 static route
	Subnet                      string `json:"subnet,omitempty"`                      // The IP address of the subnetwork specified in CIDR notation (ex. 1.2.3.0/24)
}
type ResponseSwitchCreateDeviceSwitchRoutingStaticRoute struct {
	AdvertiseViaOspfEnabled     *bool  `json:"advertiseViaOspfEnabled,omitempty"`     // Option to advertise static routes via OSPF
	ManagementNextHop           string `json:"managementNextHop,omitempty"`           // Optional fallback IP address for management traffic
	Name                        string `json:"name,omitempty"`                        // The name or description of the layer 3 static route
	NextHopIP                   string `json:"nextHopIp,omitempty"`                   // The IP address of the router to which traffic for this destination network should be sent
	PreferOverOspfRoutesEnabled *bool  `json:"preferOverOspfRoutesEnabled,omitempty"` // Option to prefer static routes over OSPF routes
	StaticRouteID               string `json:"staticRouteId,omitempty"`               // The identifier of a layer 3 static route
	Subnet                      string `json:"subnet,omitempty"`                      // The IP address of the subnetwork specified in CIDR notation (ex. 1.2.3.0/24)
}
type ResponseSwitchGetDeviceSwitchRoutingStaticRoute struct {
	AdvertiseViaOspfEnabled     *bool  `json:"advertiseViaOspfEnabled,omitempty"`     // Option to advertise static routes via OSPF
	ManagementNextHop           string `json:"managementNextHop,omitempty"`           // Optional fallback IP address for management traffic
	Name                        string `json:"name,omitempty"`                        // The name or description of the layer 3 static route
	NextHopIP                   string `json:"nextHopIp,omitempty"`                   // The IP address of the router to which traffic for this destination network should be sent
	PreferOverOspfRoutesEnabled *bool  `json:"preferOverOspfRoutesEnabled,omitempty"` // Option to prefer static routes over OSPF routes
	StaticRouteID               string `json:"staticRouteId,omitempty"`               // The identifier of a layer 3 static route
	Subnet                      string `json:"subnet,omitempty"`                      // The IP address of the subnetwork specified in CIDR notation (ex. 1.2.3.0/24)
}
type ResponseSwitchUpdateDeviceSwitchRoutingStaticRoute struct {
	AdvertiseViaOspfEnabled     *bool  `json:"advertiseViaOspfEnabled,omitempty"`     // Option to advertise static routes via OSPF
	ManagementNextHop           string `json:"managementNextHop,omitempty"`           // Optional fallback IP address for management traffic
	Name                        string `json:"name,omitempty"`                        // The name or description of the layer 3 static route
	NextHopIP                   string `json:"nextHopIp,omitempty"`                   // The IP address of the router to which traffic for this destination network should be sent
	PreferOverOspfRoutesEnabled *bool  `json:"preferOverOspfRoutesEnabled,omitempty"` // Option to prefer static routes over OSPF routes
	StaticRouteID               string `json:"staticRouteId,omitempty"`               // The identifier of a layer 3 static route
	Subnet                      string `json:"subnet,omitempty"`                      // The IP address of the subnetwork specified in CIDR notation (ex. 1.2.3.0/24)
}
type ResponseSwitchGetDeviceSwitchWarmSpare struct {
	Enabled       *bool  `json:"enabled,omitempty"`       // Enable or disable warm spare for a switch
	PrimarySerial string `json:"primarySerial,omitempty"` // Serial number of the primary switch
	SpareSerial   string `json:"spareSerial,omitempty"`   // Serial number of the warm spare switch
}
type ResponseSwitchUpdateDeviceSwitchWarmSpare struct {
	Enabled       *bool  `json:"enabled,omitempty"`       // Enable or disable warm spare for a switch
	PrimarySerial string `json:"primarySerial,omitempty"` // Serial number of the primary switch
	SpareSerial   string `json:"spareSerial,omitempty"`   // Serial number of the warm spare switch
}
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
	AccessPolicyNumber             string                                                                     `json:"accessPolicyNumber,omitempty"`             // Access Type of the policy. Automatically 'Hybrid authentication' when hostMode is 'Multi-Domain'.
	AccessPolicyType               string                                                                     `json:"accessPolicyType,omitempty"`               // Access Type of the policy. Automatically 'Hybrid authentication' when hostMode is 'Multi-Domain'.
	Counts                         *ResponseItemSwitchGetNetworkSwitchAccessPoliciesCounts                    `json:"counts,omitempty"`                         // Counts associated with the access policy
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
	RadiusGroupAttribute           string                                                                     `json:"radiusGroupAttribute,omitempty"`           // Acceptable values are `""` for , or `"11"` for Group Policies ACL
	RadiusServers                  *[]ResponseItemSwitchGetNetworkSwitchAccessPoliciesRadiusServers           `json:"radiusServers,omitempty"`                  // List of RADIUS servers to require connecting devices to authenticate against before granting network access
	RadiusTestingEnabled           *bool                                                                      `json:"radiusTestingEnabled,omitempty"`           // If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers
	URLRedirectWalledGardenEnabled *bool                                                                      `json:"urlRedirectWalledGardenEnabled,omitempty"` // Enable to restrict access for clients to a response_objectific set of IP addresses or hostnames prior to authentication
	URLRedirectWalledGardenRanges  []string                                                                   `json:"urlRedirectWalledGardenRanges,omitempty"`  // IP address ranges, in CIDR notation, to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication
	VoiceVLANClients               *bool                                                                      `json:"voiceVlanClients,omitempty"`               // CDP/LLDP capable voice clients will be able to use this VLAN. Automatically true when hostMode is 'Multi-Domain'.
}
type ResponseItemSwitchGetNetworkSwitchAccessPoliciesCounts struct {
	Ports *ResponseItemSwitchGetNetworkSwitchAccessPoliciesCountsPorts `json:"ports,omitempty"` // Counts associated with ports
}
type ResponseItemSwitchGetNetworkSwitchAccessPoliciesCountsPorts struct {
	WithThisPolicy *int `json:"withThisPolicy,omitempty"` // Number of ports in the network with this policy. For template networks, this is the number of template ports (not child ports) with this policy.
}
type ResponseItemSwitchGetNetworkSwitchAccessPoliciesDot1X struct {
	ControlDirection string `json:"controlDirection,omitempty"` // Supports either 'both' or 'inbound'. Set to 'inbound' to allow unauthorized egress on the switchport. Set to 'both' to control both traffic directions with authorization. Defaults to 'both'
}
type ResponseItemSwitchGetNetworkSwitchAccessPoliciesRadius struct {
	Cache                    *ResponseItemSwitchGetNetworkSwitchAccessPoliciesRadiusCache        `json:"cache,omitempty"`                    // Object for RADIUS Cache Settings
	CriticalAuth             *ResponseItemSwitchGetNetworkSwitchAccessPoliciesRadiusCriticalAuth `json:"criticalAuth,omitempty"`             // Critical auth settings for when authentication is rejected by the RADIUS server
	FailedAuthVLANID         *int                                                                `json:"failedAuthVlanId,omitempty"`         // VLAN that clients will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
	ReAuthenticationInterval *int                                                                `json:"reAuthenticationInterval,omitempty"` // Re-authentication period in seconds. Will be null if hostMode is Multi-Auth
}
type ResponseItemSwitchGetNetworkSwitchAccessPoliciesRadiusCache struct {
	Enabled *bool `json:"enabled,omitempty"` // Enable to cache authorization and authentication responses on the RADIUS server
	Timeout *int  `json:"timeout,omitempty"` // If RADIUS caching is enabled, this value dictates how long the cache will remain in the RADIUS server, in hours, to allow network access without authentication
}
type ResponseItemSwitchGetNetworkSwitchAccessPoliciesRadiusCriticalAuth struct {
	DataVLANID        *int  `json:"dataVlanId,omitempty"`        // VLAN that clients who use data will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
	SuspendPortBounce *bool `json:"suspendPortBounce,omitempty"` // Enable to suspend port bounce when RADIUS servers are unreachable
	VoiceVLANID       *int  `json:"voiceVlanId,omitempty"`       // VLAN that clients who use voice will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
}
type ResponseItemSwitchGetNetworkSwitchAccessPoliciesRadiusAccountingServers struct {
	Host                       string `json:"host,omitempty"`                       // Public IP address of the RADIUS accounting server
	OrganizationRadiusServerID string `json:"organizationRadiusServerId,omitempty"` // Organization wide RADIUS server ID. This value will be empty if this RADIUS server is not an organization wide RADIUS server
	Port                       *int   `json:"port,omitempty"`                       // UDP port that the RADIUS Accounting server listens on for access requests
	ServerID                   string `json:"serverId,omitempty"`                   // Unique ID of the RADIUS accounting server
}
type ResponseItemSwitchGetNetworkSwitchAccessPoliciesRadiusServers struct {
	Host                       string `json:"host,omitempty"`                       // Public IP address of the RADIUS server
	OrganizationRadiusServerID string `json:"organizationRadiusServerId,omitempty"` // Organization wide RADIUS server ID. This value will be empty if this RADIUS server is not an organization wide RADIUS server
	Port                       *int   `json:"port,omitempty"`                       // UDP port that the RADIUS server listens on for access requests
	ServerID                   string `json:"serverId,omitempty"`                   // Unique ID of the RADIUS server
}
type ResponseSwitchCreateNetworkSwitchAccessPolicy struct {
	AccessPolicyType               string                                                                  `json:"accessPolicyType,omitempty"`               // Access Type of the policy. Automatically 'Hybrid authentication' when hostMode is 'Multi-Domain'.
	Counts                         *ResponseSwitchCreateNetworkSwitchAccessPolicyCounts                    `json:"counts,omitempty"`                         // Counts associated with the access policy
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
	RadiusGroupAttribute           string                                                                  `json:"radiusGroupAttribute,omitempty"`           // Acceptable values are `""` for , or `"11"` for Group Policies ACL
	RadiusServers                  *[]ResponseSwitchCreateNetworkSwitchAccessPolicyRadiusServers           `json:"radiusServers,omitempty"`                  // List of RADIUS servers to require connecting devices to authenticate against before granting network access
	RadiusTestingEnabled           *bool                                                                   `json:"radiusTestingEnabled,omitempty"`           // If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers
	URLRedirectWalledGardenEnabled *bool                                                                   `json:"urlRedirectWalledGardenEnabled,omitempty"` // Enable to restrict access for clients to a response_objectific set of IP addresses or hostnames prior to authentication
	URLRedirectWalledGardenRanges  []string                                                                `json:"urlRedirectWalledGardenRanges,omitempty"`  // IP address ranges, in CIDR notation, to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication
	VoiceVLANClients               *bool                                                                   `json:"voiceVlanClients,omitempty"`               // CDP/LLDP capable voice clients will be able to use this VLAN. Automatically true when hostMode is 'Multi-Domain'.
}
type ResponseSwitchCreateNetworkSwitchAccessPolicyCounts struct {
	Ports *ResponseSwitchCreateNetworkSwitchAccessPolicyCountsPorts `json:"ports,omitempty"` // Counts associated with ports
}
type ResponseSwitchCreateNetworkSwitchAccessPolicyCountsPorts struct {
	WithThisPolicy *int `json:"withThisPolicy,omitempty"` // Number of ports in the network with this policy. For template networks, this is the number of template ports (not child ports) with this policy.
}
type ResponseSwitchCreateNetworkSwitchAccessPolicyDot1X struct {
	ControlDirection string `json:"controlDirection,omitempty"` // Supports either 'both' or 'inbound'. Set to 'inbound' to allow unauthorized egress on the switchport. Set to 'both' to control both traffic directions with authorization. Defaults to 'both'
}
type ResponseSwitchCreateNetworkSwitchAccessPolicyRadius struct {
	Cache                    *ResponseSwitchCreateNetworkSwitchAccessPolicyRadiusCache        `json:"cache,omitempty"`                    // Object for RADIUS Cache Settings
	CriticalAuth             *ResponseSwitchCreateNetworkSwitchAccessPolicyRadiusCriticalAuth `json:"criticalAuth,omitempty"`             // Critical auth settings for when authentication is rejected by the RADIUS server
	FailedAuthVLANID         *int                                                             `json:"failedAuthVlanId,omitempty"`         // VLAN that clients will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
	ReAuthenticationInterval *int                                                             `json:"reAuthenticationInterval,omitempty"` // Re-authentication period in seconds. Will be null if hostMode is Multi-Auth
}
type ResponseSwitchCreateNetworkSwitchAccessPolicyRadiusCache struct {
	Enabled *bool `json:"enabled,omitempty"` // Enable to cache authorization and authentication responses on the RADIUS server
	Timeout *int  `json:"timeout,omitempty"` // If RADIUS caching is enabled, this value dictates how long the cache will remain in the RADIUS server, in hours, to allow network access without authentication
}
type ResponseSwitchCreateNetworkSwitchAccessPolicyRadiusCriticalAuth struct {
	DataVLANID        *int  `json:"dataVlanId,omitempty"`        // VLAN that clients who use data will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
	SuspendPortBounce *bool `json:"suspendPortBounce,omitempty"` // Enable to suspend port bounce when RADIUS servers are unreachable
	VoiceVLANID       *int  `json:"voiceVlanId,omitempty"`       // VLAN that clients who use voice will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
}
type ResponseSwitchCreateNetworkSwitchAccessPolicyRadiusAccountingServers struct {
	Host                       string `json:"host,omitempty"`                       // Public IP address of the RADIUS accounting server
	OrganizationRadiusServerID string `json:"organizationRadiusServerId,omitempty"` // Organization wide RADIUS server ID. This value will be empty if this RADIUS server is not an organization wide RADIUS server
	Port                       *int   `json:"port,omitempty"`                       // UDP port that the RADIUS Accounting server listens on for access requests
	ServerID                   string `json:"serverId,omitempty"`                   // Unique ID of the RADIUS accounting server
}
type ResponseSwitchCreateNetworkSwitchAccessPolicyRadiusServers struct {
	Host                       string `json:"host,omitempty"`                       // Public IP address of the RADIUS server
	OrganizationRadiusServerID string `json:"organizationRadiusServerId,omitempty"` // Organization wide RADIUS server ID. This value will be empty if this RADIUS server is not an organization wide RADIUS server
	Port                       *int   `json:"port,omitempty"`                       // UDP port that the RADIUS server listens on for access requests
	ServerID                   string `json:"serverId,omitempty"`                   // Unique ID of the RADIUS server
}
type ResponseSwitchGetNetworkSwitchAccessPolicy struct {
	AccessPolicyNumber             string                                                               `json:"accessPolicyNumber,omitempty"`             //
	AccessPolicyType               string                                                               `json:"accessPolicyType,omitempty"`               // Access Type of the policy. Automatically 'Hybrid authentication' when hostMode is 'Multi-Domain'.
	Counts                         *ResponseSwitchGetNetworkSwitchAccessPolicyCounts                    `json:"counts,omitempty"`                         // Counts associated with the access policy
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
	RadiusGroupAttribute           string                                                               `json:"radiusGroupAttribute,omitempty"`           // Acceptable values are `""` for , or `"11"` for Group Policies ACL
	RadiusServers                  *[]ResponseSwitchGetNetworkSwitchAccessPolicyRadiusServers           `json:"radiusServers,omitempty"`                  // List of RADIUS servers to require connecting devices to authenticate against before granting network access
	RadiusTestingEnabled           *bool                                                                `json:"radiusTestingEnabled,omitempty"`           // If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers
	URLRedirectWalledGardenEnabled *bool                                                                `json:"urlRedirectWalledGardenEnabled,omitempty"` // Enable to restrict access for clients to a response_objectific set of IP addresses or hostnames prior to authentication
	URLRedirectWalledGardenRanges  []string                                                             `json:"urlRedirectWalledGardenRanges,omitempty"`  // IP address ranges, in CIDR notation, to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication
	VoiceVLANClients               *bool                                                                `json:"voiceVlanClients,omitempty"`               // CDP/LLDP capable voice clients will be able to use this VLAN. Automatically true when hostMode is 'Multi-Domain'.
}
type ResponseSwitchGetNetworkSwitchAccessPolicyCounts struct {
	Ports *ResponseSwitchGetNetworkSwitchAccessPolicyCountsPorts `json:"ports,omitempty"` // Counts associated with ports
}
type ResponseSwitchGetNetworkSwitchAccessPolicyCountsPorts struct {
	WithThisPolicy *int `json:"withThisPolicy,omitempty"` // Number of ports in the network with this policy. For template networks, this is the number of template ports (not child ports) with this policy.
}
type ResponseSwitchGetNetworkSwitchAccessPolicyDot1X struct {
	ControlDirection string `json:"controlDirection,omitempty"` // Supports either 'both' or 'inbound'. Set to 'inbound' to allow unauthorized egress on the switchport. Set to 'both' to control both traffic directions with authorization. Defaults to 'both'
}
type ResponseSwitchGetNetworkSwitchAccessPolicyRadius struct {
	Cache                    *ResponseSwitchGetNetworkSwitchAccessPolicyRadiusCache        `json:"cache,omitempty"`                    // Object for RADIUS Cache Settings
	CriticalAuth             *ResponseSwitchGetNetworkSwitchAccessPolicyRadiusCriticalAuth `json:"criticalAuth,omitempty"`             // Critical auth settings for when authentication is rejected by the RADIUS server
	FailedAuthVLANID         *int                                                          `json:"failedAuthVlanId,omitempty"`         // VLAN that clients will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
	ReAuthenticationInterval *int                                                          `json:"reAuthenticationInterval,omitempty"` // Re-authentication period in seconds. Will be null if hostMode is Multi-Auth
}
type ResponseSwitchGetNetworkSwitchAccessPolicyRadiusCache struct {
	Enabled *bool `json:"enabled,omitempty"` // Enable to cache authorization and authentication responses on the RADIUS server
	Timeout *int  `json:"timeout,omitempty"` // If RADIUS caching is enabled, this value dictates how long the cache will remain in the RADIUS server, in hours, to allow network access without authentication
}
type ResponseSwitchGetNetworkSwitchAccessPolicyRadiusCriticalAuth struct {
	DataVLANID        *int  `json:"dataVlanId,omitempty"`        // VLAN that clients who use data will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
	SuspendPortBounce *bool `json:"suspendPortBounce,omitempty"` // Enable to suspend port bounce when RADIUS servers are unreachable
	VoiceVLANID       *int  `json:"voiceVlanId,omitempty"`       // VLAN that clients who use voice will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
}
type ResponseSwitchGetNetworkSwitchAccessPolicyRadiusAccountingServers struct {
	Host                       string `json:"host,omitempty"`                       // Public IP address of the RADIUS accounting server
	OrganizationRadiusServerID string `json:"organizationRadiusServerId,omitempty"` // Organization wide RADIUS server ID. This value will be empty if this RADIUS server is not an organization wide RADIUS server
	Port                       *int   `json:"port,omitempty"`                       // UDP port that the RADIUS Accounting server listens on for access requests
	ServerID                   string `json:"serverId,omitempty"`                   // Unique ID of the RADIUS accounting server
}
type ResponseSwitchGetNetworkSwitchAccessPolicyRadiusServers struct {
	Host                       string `json:"host,omitempty"`                       // Public IP address of the RADIUS server
	OrganizationRadiusServerID string `json:"organizationRadiusServerId,omitempty"` // Organization wide RADIUS server ID. This value will be empty if this RADIUS server is not an organization wide RADIUS server
	Port                       *int   `json:"port,omitempty"`                       // UDP port that the RADIUS server listens on for access requests
	ServerID                   string `json:"serverId,omitempty"`                   // Unique ID of the RADIUS server
}
type ResponseSwitchUpdateNetworkSwitchAccessPolicy struct {
	AccessPolicyType               string                                                                  `json:"accessPolicyType,omitempty"`               // Access Type of the policy. Automatically 'Hybrid authentication' when hostMode is 'Multi-Domain'.
	Counts                         *ResponseSwitchUpdateNetworkSwitchAccessPolicyCounts                    `json:"counts,omitempty"`                         // Counts associated with the access policy
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
	RadiusGroupAttribute           string                                                                  `json:"radiusGroupAttribute,omitempty"`           // Acceptable values are `""` for , or `"11"` for Group Policies ACL
	RadiusServers                  *[]ResponseSwitchUpdateNetworkSwitchAccessPolicyRadiusServers           `json:"radiusServers,omitempty"`                  // List of RADIUS servers to require connecting devices to authenticate against before granting network access
	RadiusTestingEnabled           *bool                                                                   `json:"radiusTestingEnabled,omitempty"`           // If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers
	URLRedirectWalledGardenEnabled *bool                                                                   `json:"urlRedirectWalledGardenEnabled,omitempty"` // Enable to restrict access for clients to a response_objectific set of IP addresses or hostnames prior to authentication
	URLRedirectWalledGardenRanges  []string                                                                `json:"urlRedirectWalledGardenRanges,omitempty"`  // IP address ranges, in CIDR notation, to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication
	VoiceVLANClients               *bool                                                                   `json:"voiceVlanClients,omitempty"`               // CDP/LLDP capable voice clients will be able to use this VLAN. Automatically true when hostMode is 'Multi-Domain'.
}
type ResponseSwitchUpdateNetworkSwitchAccessPolicyCounts struct {
	Ports *ResponseSwitchUpdateNetworkSwitchAccessPolicyCountsPorts `json:"ports,omitempty"` // Counts associated with ports
}
type ResponseSwitchUpdateNetworkSwitchAccessPolicyCountsPorts struct {
	WithThisPolicy *int `json:"withThisPolicy,omitempty"` // Number of ports in the network with this policy. For template networks, this is the number of template ports (not child ports) with this policy.
}
type ResponseSwitchUpdateNetworkSwitchAccessPolicyDot1X struct {
	ControlDirection string `json:"controlDirection,omitempty"` // Supports either 'both' or 'inbound'. Set to 'inbound' to allow unauthorized egress on the switchport. Set to 'both' to control both traffic directions with authorization. Defaults to 'both'
}
type ResponseSwitchUpdateNetworkSwitchAccessPolicyRadius struct {
	Cache                    *ResponseSwitchUpdateNetworkSwitchAccessPolicyRadiusCache        `json:"cache,omitempty"`                    // Object for RADIUS Cache Settings
	CriticalAuth             *ResponseSwitchUpdateNetworkSwitchAccessPolicyRadiusCriticalAuth `json:"criticalAuth,omitempty"`             // Critical auth settings for when authentication is rejected by the RADIUS server
	FailedAuthVLANID         *int                                                             `json:"failedAuthVlanId,omitempty"`         // VLAN that clients will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
	ReAuthenticationInterval *int                                                             `json:"reAuthenticationInterval,omitempty"` // Re-authentication period in seconds. Will be null if hostMode is Multi-Auth
}
type ResponseSwitchUpdateNetworkSwitchAccessPolicyRadiusCache struct {
	Enabled *bool `json:"enabled,omitempty"` // Enable to cache authorization and authentication responses on the RADIUS server
	Timeout *int  `json:"timeout,omitempty"` // If RADIUS caching is enabled, this value dictates how long the cache will remain in the RADIUS server, in hours, to allow network access without authentication
}
type ResponseSwitchUpdateNetworkSwitchAccessPolicyRadiusCriticalAuth struct {
	DataVLANID        *int  `json:"dataVlanId,omitempty"`        // VLAN that clients who use data will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
	SuspendPortBounce *bool `json:"suspendPortBounce,omitempty"` // Enable to suspend port bounce when RADIUS servers are unreachable
	VoiceVLANID       *int  `json:"voiceVlanId,omitempty"`       // VLAN that clients who use voice will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
}
type ResponseSwitchUpdateNetworkSwitchAccessPolicyRadiusAccountingServers struct {
	Host                       string `json:"host,omitempty"`                       // Public IP address of the RADIUS accounting server
	OrganizationRadiusServerID string `json:"organizationRadiusServerId,omitempty"` // Organization wide RADIUS server ID. This value will be empty if this RADIUS server is not an organization wide RADIUS server
	Port                       *int   `json:"port,omitempty"`                       // UDP port that the RADIUS Accounting server listens on for access requests
	ServerID                   string `json:"serverId,omitempty"`                   // Unique ID of the RADIUS accounting server
}
type ResponseSwitchUpdateNetworkSwitchAccessPolicyRadiusServers struct {
	Host                       string `json:"host,omitempty"`                       // Public IP address of the RADIUS server
	OrganizationRadiusServerID string `json:"organizationRadiusServerId,omitempty"` // Organization wide RADIUS server ID. This value will be empty if this RADIUS server is not an organization wide RADIUS server
	Port                       *int   `json:"port,omitempty"`                       // UDP port that the RADIUS server listens on for access requests
	ServerID                   string `json:"serverId,omitempty"`                   // Unique ID of the RADIUS server
}
type ResponseSwitchGetNetworkSwitchAlternateManagementInterface struct {
	Enabled   *bool                                                                 `json:"enabled,omitempty"`   // Boolean value to enable or disable AMI configuration. If enabled, VLAN and protocols must be set
	Protocols []string                                                              `json:"protocols,omitempty"` // Can be one or more of the following values: 'radius', 'snmp' or 'syslog'
	Switches  *[]ResponseSwitchGetNetworkSwitchAlternateManagementInterfaceSwitches `json:"switches,omitempty"`  // Array of switch serial number and IP assignment. If parameter is present, it cannot have empty body. Note: switches parameter is not applicable for template networks, in other words, do not put 'switches' in the body when updating template networks. Also, an empty 'switches' array will remove all previous assignments
	VLANID    *int                                                                  `json:"vlanId,omitempty"`    // Alternate management VLAN, must be between 1 and 4094
}
type ResponseSwitchGetNetworkSwitchAlternateManagementInterfaceSwitches struct {
	AlternateManagementIP string `json:"alternateManagementIp,omitempty"` // Switch alternative management IP. To remove a prior IP setting, provide an empty string
	Gateway               string `json:"gateway,omitempty"`               // Switch gateway must be in IP format. Only and must be specified for Polaris switches
	Serial                string `json:"serial,omitempty"`                // Switch serial number
	SubnetMask            string `json:"subnetMask,omitempty"`            // Switch subnet mask must be in IP format. Only and must be specified for Polaris switches
}
type ResponseSwitchUpdateNetworkSwitchAlternateManagementInterface struct {
	Enabled   *bool                                                                    `json:"enabled,omitempty"`   // Boolean value to enable or disable AMI configuration. If enabled, VLAN and protocols must be set
	Protocols []string                                                                 `json:"protocols,omitempty"` // Can be one or more of the following values: 'radius', 'snmp' or 'syslog'
	Switches  *[]ResponseSwitchUpdateNetworkSwitchAlternateManagementInterfaceSwitches `json:"switches,omitempty"`  // Array of switch serial number and IP assignment. If parameter is present, it cannot have empty body. Note: switches parameter is not applicable for template networks, in other words, do not put 'switches' in the body when updating template networks. Also, an empty 'switches' array will remove all previous assignments
	VLANID    *int                                                                     `json:"vlanId,omitempty"`    // Alternate management VLAN, must be between 1 and 4094
}
type ResponseSwitchUpdateNetworkSwitchAlternateManagementInterfaceSwitches struct {
	AlternateManagementIP string `json:"alternateManagementIp,omitempty"` // Switch alternative management IP. To remove a prior IP setting, provide an empty string
	Gateway               string `json:"gateway,omitempty"`               // Switch gateway must be in IP format. Only and must be specified for Polaris switches
	Serial                string `json:"serial,omitempty"`                // Switch serial number
	SubnetMask            string `json:"subnetMask,omitempty"`            // Switch subnet mask must be in IP format. Only and must be specified for Polaris switches
}
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
	Alerts         *ResponseSwitchGetNetworkSwitchDhcpServerPolicyAlerts        `json:"alerts,omitempty"`         // Email alert settings for DHCP servers
	AllowedServers []string                                                     `json:"allowedServers,omitempty"` // List the MAC addresses of DHCP servers to permit on the network when defaultPolicy is set       to block.An empty array will clear the entries.
	ArpInspection  *ResponseSwitchGetNetworkSwitchDhcpServerPolicyArpInspection `json:"arpInspection,omitempty"`  // Dynamic ARP Inspection settings
	BlockedServers []string                                                     `json:"blockedServers,omitempty"` // List the MAC addresses of DHCP servers to block on the network when defaultPolicy is set       to allow.An empty array will clear the entries.
	DefaultPolicy  string                                                       `json:"defaultPolicy,omitempty"`  // 'allow' or 'block' new DHCP servers. Default value is 'allow'.
}
type ResponseSwitchGetNetworkSwitchDhcpServerPolicyAlerts struct {
	Email *ResponseSwitchGetNetworkSwitchDhcpServerPolicyAlertsEmail `json:"email,omitempty"` // Alert settings for DHCP servers
}
type ResponseSwitchGetNetworkSwitchDhcpServerPolicyAlertsEmail struct {
	Enabled *bool `json:"enabled,omitempty"` // When enabled, send an email if a new DHCP server is seen. Default value is false.
}
type ResponseSwitchGetNetworkSwitchDhcpServerPolicyArpInspection struct {
	Enabled           *bool    `json:"enabled,omitempty"`           // Enable or disable Dynamic ARP Inspection on the network. Default value is false.
	UnsupportedModels []string `json:"unsupportedModels,omitempty"` // List of switch models that does not support dynamic ARP inspection
}
type ResponseSwitchUpdateNetworkSwitchDhcpServerPolicy struct {
	Alerts         *ResponseSwitchUpdateNetworkSwitchDhcpServerPolicyAlerts        `json:"alerts,omitempty"`         // Email alert settings for DHCP servers
	AllowedServers []string                                                        `json:"allowedServers,omitempty"` // List the MAC addresses of DHCP servers to permit on the network when defaultPolicy is set       to block.An empty array will clear the entries.
	ArpInspection  *ResponseSwitchUpdateNetworkSwitchDhcpServerPolicyArpInspection `json:"arpInspection,omitempty"`  // Dynamic ARP Inspection settings
	BlockedServers []string                                                        `json:"blockedServers,omitempty"` // List the MAC addresses of DHCP servers to block on the network when defaultPolicy is set       to allow.An empty array will clear the entries.
	DefaultPolicy  string                                                          `json:"defaultPolicy,omitempty"`  // 'allow' or 'block' new DHCP servers. Default value is 'allow'.
}
type ResponseSwitchUpdateNetworkSwitchDhcpServerPolicyAlerts struct {
	Email *ResponseSwitchUpdateNetworkSwitchDhcpServerPolicyAlertsEmail `json:"email,omitempty"` // Alert settings for DHCP servers
}
type ResponseSwitchUpdateNetworkSwitchDhcpServerPolicyAlertsEmail struct {
	Enabled *bool `json:"enabled,omitempty"` // When enabled, send an email if a new DHCP server is seen. Default value is false.
}
type ResponseSwitchUpdateNetworkSwitchDhcpServerPolicyArpInspection struct {
	Enabled           *bool    `json:"enabled,omitempty"`           // Enable or disable Dynamic ARP Inspection on the network. Default value is false.
	UnsupportedModels []string `json:"unsupportedModels,omitempty"` // List of switch models that does not support dynamic ARP inspection
}
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
	Mappings *[]ResponseSwitchGetNetworkSwitchDscpToCosMappingsMappings `json:"mappings,omitempty"` // An array of DSCP to CoS mappings. An empty array will reset the mappings to default.
}
type ResponseSwitchGetNetworkSwitchDscpToCosMappingsMappings struct {
	Cos   *int   `json:"cos,omitempty"`   // The actual layer-2 CoS queue the DSCP value is mapped to. These are not bits set on outgoing frames. Value can be in the range of 0 to 5 inclusive.
	Dscp  *int   `json:"dscp,omitempty"`  // The Differentiated Services Code Point (DSCP) tag in the IP header that will be mapped to a particular Class-of-Service (CoS) queue. Value can be in the range of 0 to 63 inclusive.
	Title string `json:"title,omitempty"` // Label for the mapping (optional).
}
type ResponseSwitchUpdateNetworkSwitchDscpToCosMappings struct {
	Mappings *[]ResponseSwitchUpdateNetworkSwitchDscpToCosMappingsMappings `json:"mappings,omitempty"` // An array of DSCP to CoS mappings. An empty array will reset the mappings to default.
}
type ResponseSwitchUpdateNetworkSwitchDscpToCosMappingsMappings struct {
	Cos   *int   `json:"cos,omitempty"`   // The actual layer-2 CoS queue the DSCP value is mapped to. These are not bits set on outgoing frames. Value can be in the range of 0 to 5 inclusive.
	Dscp  *int   `json:"dscp,omitempty"`  // The Differentiated Services Code Point (DSCP) tag in the IP header that will be mapped to a particular Class-of-Service (CoS) queue. Value can be in the range of 0 to 63 inclusive.
	Title string `json:"title,omitempty"` // Label for the mapping (optional).
}
type ResponseSwitchGetNetworkSwitchLinkAggregations []ResponseItemSwitchGetNetworkSwitchLinkAggregations // Array of ResponseSwitchGetNetworkSwitchLinkAggregations
type ResponseItemSwitchGetNetworkSwitchLinkAggregations struct {
	ID          string                                                           `json:"id,omitempty"`          // The ID for the link aggregation.
	SwitchPorts *[]ResponseItemSwitchGetNetworkSwitchLinkAggregationsSwitchPorts `json:"switchPorts,omitempty"` // The ID for the link aggregation.
}
type ResponseItemSwitchGetNetworkSwitchLinkAggregationsSwitchPorts struct {
	PortID string `json:"portId,omitempty"` // The ID for the switch port.
	Serial string `json:"serial,omitempty"` // The serial number for the switch port.
}
type ResponseSwitchCreateNetworkSwitchLinkAggregation struct {
	ID          string                                                         `json:"id,omitempty"`          // The ID for the link aggregation.
	SwitchPorts *[]ResponseSwitchCreateNetworkSwitchLinkAggregationSwitchPorts `json:"switchPorts,omitempty"` // The ID for the link aggregation.
}
type ResponseSwitchCreateNetworkSwitchLinkAggregationSwitchPorts struct {
	PortID string `json:"portId,omitempty"` // The ID for the switch port.
	Serial string `json:"serial,omitempty"` // The serial number for the switch port.
}
type ResponseSwitchUpdateNetworkSwitchLinkAggregation struct {
	ID          string                                                         `json:"id,omitempty"`          // The ID for the link aggregation.
	SwitchPorts *[]ResponseSwitchUpdateNetworkSwitchLinkAggregationSwitchPorts `json:"switchPorts,omitempty"` // The ID for the link aggregation.
}
type ResponseSwitchUpdateNetworkSwitchLinkAggregationSwitchPorts struct {
	PortID string `json:"portId,omitempty"` // The ID for the switch port.
	Serial string `json:"serial,omitempty"` // The serial number for the switch port.
}
type ResponseSwitchGetNetworkSwitchMtu struct {
	DefaultMtuSize *int                                          `json:"defaultMtuSize,omitempty"` // MTU size for the entire network. Default value is 9578.
	Overrides      *[]ResponseSwitchGetNetworkSwitchMtuOverrides `json:"overrides,omitempty"`      // Override MTU size for individual switches or switch templates.       An empty array will clear overrides.
}
type ResponseSwitchGetNetworkSwitchMtuOverrides struct {
	MtuSize        *int     `json:"mtuSize,omitempty"`        // MTU size for the switches or switch templates.
	SwitchProfiles []string `json:"switchProfiles,omitempty"` // List of switch template IDs. Applicable only for template network.
	Switches       []string `json:"switches,omitempty"`       // List of switch serials. Applicable only for switch network.
}
type ResponseSwitchUpdateNetworkSwitchMtu struct {
	DefaultMtuSize *int                                             `json:"defaultMtuSize,omitempty"` // MTU size for the entire network. Default value is 9578.
	Overrides      *[]ResponseSwitchUpdateNetworkSwitchMtuOverrides `json:"overrides,omitempty"`      // Override MTU size for individual switches or switch templates.       An empty array will clear overrides.
}
type ResponseSwitchUpdateNetworkSwitchMtuOverrides struct {
	MtuSize        *int     `json:"mtuSize,omitempty"`        // MTU size for the switches or switch templates.
	SwitchProfiles []string `json:"switchProfiles,omitempty"` // List of switch template IDs. Applicable only for template network.
	Switches       []string `json:"switches,omitempty"`       // List of switch serials. Applicable only for switch network.
}
type ResponseSwitchGetNetworkSwitchPortSchedules []ResponseItemSwitchGetNetworkSwitchPortSchedules // Array of ResponseSwitchGetNetworkSwitchPortSchedules
type ResponseItemSwitchGetNetworkSwitchPortSchedules struct {
	ID           string                                                       `json:"id,omitempty"`           // Switch port schedule ID
	Name         string                                                       `json:"name,omitempty"`         // Switch port schedule name
	NetworkID    string                                                       `json:"networkId,omitempty"`    // Network ID
	PortSchedule *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortSchedule `json:"portSchedule,omitempty"` // Port schedule
}
type ResponseItemSwitchGetNetworkSwitchPortSchedulesPortSchedule struct {
	Friday    *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleFriday    `json:"friday,omitempty"`    // Friday schedule
	Monday    *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleMonday    `json:"monday,omitempty"`    // Monday schedule
	Saturday  *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleSaturday  `json:"saturday,omitempty"`  // Saturday schedule
	Sunday    *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleSunday    `json:"sunday,omitempty"`    // Sunday schedule
	Thursday  *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleThursday  `json:"thursday,omitempty"`  // Thursday schedule
	Tuesday   *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleTuesday   `json:"tuesday,omitempty"`   // Tuesday schedule
	Wednesday *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleWednesday `json:"wednesday,omitempty"` // Wednesday schedule
}
type ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleFriday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active or inactive
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'
}
type ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleMonday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active or inactive
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'
}
type ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleSaturday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active or inactive
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'
}
type ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleSunday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active or inactive
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'
}
type ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleThursday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active or inactive
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'
}
type ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleTuesday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active or inactive
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'
}
type ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleWednesday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active or inactive
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'
}
type ResponseSwitchCreateNetworkSwitchPortSchedule struct {
	ID           string                                                     `json:"id,omitempty"`           // Switch port schedule ID
	Name         string                                                     `json:"name,omitempty"`         // Switch port schedule name
	NetworkID    string                                                     `json:"networkId,omitempty"`    // Network ID
	PortSchedule *ResponseSwitchCreateNetworkSwitchPortSchedulePortSchedule `json:"portSchedule,omitempty"` // Port schedule
}
type ResponseSwitchCreateNetworkSwitchPortSchedulePortSchedule struct {
	Friday    *ResponseSwitchCreateNetworkSwitchPortSchedulePortScheduleFriday    `json:"friday,omitempty"`    // Friday schedule
	Monday    *ResponseSwitchCreateNetworkSwitchPortSchedulePortScheduleMonday    `json:"monday,omitempty"`    // Monday schedule
	Saturday  *ResponseSwitchCreateNetworkSwitchPortSchedulePortScheduleSaturday  `json:"saturday,omitempty"`  // Saturday schedule
	Sunday    *ResponseSwitchCreateNetworkSwitchPortSchedulePortScheduleSunday    `json:"sunday,omitempty"`    // Sunday schedule
	Thursday  *ResponseSwitchCreateNetworkSwitchPortSchedulePortScheduleThursday  `json:"thursday,omitempty"`  // Thursday schedule
	Tuesday   *ResponseSwitchCreateNetworkSwitchPortSchedulePortScheduleTuesday   `json:"tuesday,omitempty"`   // Tuesday schedule
	Wednesday *ResponseSwitchCreateNetworkSwitchPortSchedulePortScheduleWednesday `json:"wednesday,omitempty"` // Wednesday schedule
}
type ResponseSwitchCreateNetworkSwitchPortSchedulePortScheduleFriday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active or inactive
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'
}
type ResponseSwitchCreateNetworkSwitchPortSchedulePortScheduleMonday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active or inactive
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'
}
type ResponseSwitchCreateNetworkSwitchPortSchedulePortScheduleSaturday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active or inactive
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'
}
type ResponseSwitchCreateNetworkSwitchPortSchedulePortScheduleSunday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active or inactive
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'
}
type ResponseSwitchCreateNetworkSwitchPortSchedulePortScheduleThursday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active or inactive
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'
}
type ResponseSwitchCreateNetworkSwitchPortSchedulePortScheduleTuesday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active or inactive
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'
}
type ResponseSwitchCreateNetworkSwitchPortSchedulePortScheduleWednesday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active or inactive
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'
}
type ResponseSwitchUpdateNetworkSwitchPortSchedule struct {
	ID           string                                                     `json:"id,omitempty"`           // Switch port schedule ID
	Name         string                                                     `json:"name,omitempty"`         // Switch port schedule name
	NetworkID    string                                                     `json:"networkId,omitempty"`    // Network ID
	PortSchedule *ResponseSwitchUpdateNetworkSwitchPortSchedulePortSchedule `json:"portSchedule,omitempty"` // Port schedule
}
type ResponseSwitchUpdateNetworkSwitchPortSchedulePortSchedule struct {
	Friday    *ResponseSwitchUpdateNetworkSwitchPortSchedulePortScheduleFriday    `json:"friday,omitempty"`    // Friday schedule
	Monday    *ResponseSwitchUpdateNetworkSwitchPortSchedulePortScheduleMonday    `json:"monday,omitempty"`    // Monday schedule
	Saturday  *ResponseSwitchUpdateNetworkSwitchPortSchedulePortScheduleSaturday  `json:"saturday,omitempty"`  // Saturday schedule
	Sunday    *ResponseSwitchUpdateNetworkSwitchPortSchedulePortScheduleSunday    `json:"sunday,omitempty"`    // Sunday schedule
	Thursday  *ResponseSwitchUpdateNetworkSwitchPortSchedulePortScheduleThursday  `json:"thursday,omitempty"`  // Thursday schedule
	Tuesday   *ResponseSwitchUpdateNetworkSwitchPortSchedulePortScheduleTuesday   `json:"tuesday,omitempty"`   // Tuesday schedule
	Wednesday *ResponseSwitchUpdateNetworkSwitchPortSchedulePortScheduleWednesday `json:"wednesday,omitempty"` // Wednesday schedule
}
type ResponseSwitchUpdateNetworkSwitchPortSchedulePortScheduleFriday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active or inactive
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'
}
type ResponseSwitchUpdateNetworkSwitchPortSchedulePortScheduleMonday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active or inactive
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'
}
type ResponseSwitchUpdateNetworkSwitchPortSchedulePortScheduleSaturday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active or inactive
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'
}
type ResponseSwitchUpdateNetworkSwitchPortSchedulePortScheduleSunday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active or inactive
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'
}
type ResponseSwitchUpdateNetworkSwitchPortSchedulePortScheduleThursday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active or inactive
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'
}
type ResponseSwitchUpdateNetworkSwitchPortSchedulePortScheduleTuesday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active or inactive
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'
}
type ResponseSwitchUpdateNetworkSwitchPortSchedulePortScheduleWednesday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active or inactive
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'
}
type ResponseSwitchGetNetworkSwitchQosRules []ResponseItemSwitchGetNetworkSwitchQosRules // Array of ResponseSwitchGetNetworkSwitchQosRules
type ResponseItemSwitchGetNetworkSwitchQosRules struct {
	Dscp         *int   `json:"dscp,omitempty"`         // DSCP tag for the incoming packet. Set this to -1 to trust incoming DSCP. Default value is 0
	DstPort      *int   `json:"dstPort,omitempty"`      // The destination port of the incoming packet. Applicable only if protocol is TCP or UDP.
	DstPortRange string `json:"dstPortRange,omitempty"` // The destination port range of the incoming packet. Applicable only if protocol is set to TCP or UDP. Example: 70-80
	ID           string `json:"id,omitempty"`           // Qos Rule id
	Protocol     string `json:"protocol,omitempty"`     // The protocol of the incoming packet. Can be one of "ANY", "TCP" or "UDP". Default value is "ANY"
	SrcPort      *int   `json:"srcPort,omitempty"`      // The source port of the incoming packet. Applicable only if protocol is TCP or UDP.
	SrcPortRange string `json:"srcPortRange,omitempty"` // The source port range of the incoming packet. Applicable only if protocol is set to TCP or UDP. Example: 70-80
	VLAN         *int   `json:"vlan,omitempty"`         // The VLAN of the incoming packet. A null value will match any VLAN.
}
type ResponseSwitchCreateNetworkSwitchQosRule struct {
	Dscp         *int   `json:"dscp,omitempty"`         // DSCP tag for the incoming packet. Set this to -1 to trust incoming DSCP. Default value is 0
	DstPort      *int   `json:"dstPort,omitempty"`      // The destination port of the incoming packet. Applicable only if protocol is TCP or UDP.
	DstPortRange string `json:"dstPortRange,omitempty"` // The destination port range of the incoming packet. Applicable only if protocol is set to TCP or UDP. Example: 70-80
	ID           string `json:"id,omitempty"`           // Qos Rule id
	Protocol     string `json:"protocol,omitempty"`     // The protocol of the incoming packet. Can be one of "ANY", "TCP" or "UDP". Default value is "ANY"
	SrcPort      *int   `json:"srcPort,omitempty"`      // The source port of the incoming packet. Applicable only if protocol is TCP or UDP.
	SrcPortRange string `json:"srcPortRange,omitempty"` // The source port range of the incoming packet. Applicable only if protocol is set to TCP or UDP. Example: 70-80
	VLAN         *int   `json:"vlan,omitempty"`         // The VLAN of the incoming packet. A null value will match any VLAN.
}
type ResponseSwitchGetNetworkSwitchQosRulesOrder struct {
	RuleIDs []string `json:"ruleIds,omitempty"` // Qos Rule ids
}
type ResponseSwitchUpdateNetworkSwitchQosRulesOrder struct {
	RuleIDs []string `json:"ruleIds,omitempty"` // Qos Rule ids
}
type ResponseSwitchGetNetworkSwitchQosRule struct {
	Dscp         *int   `json:"dscp,omitempty"`         // DSCP tag for the incoming packet. Set this to -1 to trust incoming DSCP. Default value is 0
	DstPort      *int   `json:"dstPort,omitempty"`      // The destination port of the incoming packet. Applicable only if protocol is TCP or UDP.
	DstPortRange string `json:"dstPortRange,omitempty"` // The destination port range of the incoming packet. Applicable only if protocol is set to TCP or UDP. Example: 70-80
	ID           string `json:"id,omitempty"`           // Qos Rule id
	Protocol     string `json:"protocol,omitempty"`     // The protocol of the incoming packet. Can be one of "ANY", "TCP" or "UDP". Default value is "ANY"
	SrcPort      *int   `json:"srcPort,omitempty"`      // The source port of the incoming packet. Applicable only if protocol is TCP or UDP.
	SrcPortRange string `json:"srcPortRange,omitempty"` // The source port range of the incoming packet. Applicable only if protocol is set to TCP or UDP. Example: 70-80
	VLAN         *int   `json:"vlan,omitempty"`         // The VLAN of the incoming packet. A null value will match any VLAN.
}
type ResponseSwitchUpdateNetworkSwitchQosRule struct {
	Dscp         *int   `json:"dscp,omitempty"`         // DSCP tag for the incoming packet. Set this to -1 to trust incoming DSCP. Default value is 0
	DstPort      *int   `json:"dstPort,omitempty"`      // The destination port of the incoming packet. Applicable only if protocol is TCP or UDP.
	DstPortRange string `json:"dstPortRange,omitempty"` // The destination port range of the incoming packet. Applicable only if protocol is set to TCP or UDP. Example: 70-80
	ID           string `json:"id,omitempty"`           // Qos Rule id
	Protocol     string `json:"protocol,omitempty"`     // The protocol of the incoming packet. Can be one of "ANY", "TCP" or "UDP". Default value is "ANY"
	SrcPort      *int   `json:"srcPort,omitempty"`      // The source port of the incoming packet. Applicable only if protocol is TCP or UDP.
	SrcPortRange string `json:"srcPortRange,omitempty"` // The source port range of the incoming packet. Applicable only if protocol is set to TCP or UDP. Example: 70-80
	VLAN         *int   `json:"vlan,omitempty"`         // The VLAN of the incoming packet. A null value will match any VLAN.
}
type ResponseSwitchGetNetworkSwitchRoutingMulticast struct {
	DefaultSettings *ResponseSwitchGetNetworkSwitchRoutingMulticastDefaultSettings `json:"defaultSettings,omitempty"` // Default multicast setting for entire network. IGMP snooping and Flood unknown       multicast traffic settings are enabled by default.
	Overrides       *[]ResponseSwitchGetNetworkSwitchRoutingMulticastOverrides     `json:"overrides,omitempty"`       // Array of paired switches/stacks/profiles and corresponding multicast settings.       An empty array will clear the multicast settings.
}
type ResponseSwitchGetNetworkSwitchRoutingMulticastDefaultSettings struct {
	FloodUnknownMulticastTrafficEnabled *bool `json:"floodUnknownMulticastTrafficEnabled,omitempty"` // Flood unknown multicast traffic enabled for the entire network
	IgmpSnoopingEnabled                 *bool `json:"igmpSnoopingEnabled,omitempty"`                 // IGMP snooping enabled for the entire network
}
type ResponseSwitchGetNetworkSwitchRoutingMulticastOverrides struct {
	FloodUnknownMulticastTrafficEnabled *bool    `json:"floodUnknownMulticastTrafficEnabled,omitempty"` // Flood unknown multicast traffic enabled for switches, switch stacks or switch templates
	IgmpSnoopingEnabled                 *bool    `json:"igmpSnoopingEnabled,omitempty"`                 // IGMP snooping enabled for switches, switch stacks or switch templates
	Stacks                              []string `json:"stacks,omitempty"`                              // (optional) List of switch stack ids for non-template network
	SwitchProfiles                      []string `json:"switchProfiles,omitempty"`                      // (optional) List of switch templates ids for template network
	Switches                            []string `json:"switches,omitempty"`                            // (optional) List of switch serials for non-template network
}
type ResponseSwitchUpdateNetworkSwitchRoutingMulticast struct {
	DefaultSettings *ResponseSwitchUpdateNetworkSwitchRoutingMulticastDefaultSettings `json:"defaultSettings,omitempty"` // Default multicast setting for entire network. IGMP snooping and Flood unknown       multicast traffic settings are enabled by default.
	Overrides       *[]ResponseSwitchUpdateNetworkSwitchRoutingMulticastOverrides     `json:"overrides,omitempty"`       // Array of paired switches/stacks/profiles and corresponding multicast settings.       An empty array will clear the multicast settings.
}
type ResponseSwitchUpdateNetworkSwitchRoutingMulticastDefaultSettings struct {
	FloodUnknownMulticastTrafficEnabled *bool `json:"floodUnknownMulticastTrafficEnabled,omitempty"` // Flood unknown multicast traffic enabled for the entire network
	IgmpSnoopingEnabled                 *bool `json:"igmpSnoopingEnabled,omitempty"`                 // IGMP snooping enabled for the entire network
}
type ResponseSwitchUpdateNetworkSwitchRoutingMulticastOverrides struct {
	FloodUnknownMulticastTrafficEnabled *bool    `json:"floodUnknownMulticastTrafficEnabled,omitempty"` // Flood unknown multicast traffic enabled for switches, switch stacks or switch templates
	IgmpSnoopingEnabled                 *bool    `json:"igmpSnoopingEnabled,omitempty"`                 // IGMP snooping enabled for switches, switch stacks or switch templates
	Stacks                              []string `json:"stacks,omitempty"`                              // (optional) List of switch stack ids for non-template network
	SwitchProfiles                      []string `json:"switchProfiles,omitempty"`                      // (optional) List of switch templates ids for template network
	Switches                            []string `json:"switches,omitempty"`                            // (optional) List of switch serials for non-template network
}
type ResponseSwitchGetNetworkSwitchRoutingMulticastRendezvousPoints []ResponseItemSwitchGetNetworkSwitchRoutingMulticastRendezvousPoints // Array of ResponseSwitchGetNetworkSwitchRoutingMulticastRendezvousPoints
type ResponseItemSwitchGetNetworkSwitchRoutingMulticastRendezvousPoints struct {
	InterfaceIP       string `json:"interfaceIp,omitempty"`       // The IP address of the interface to use.
	InterfaceName     string `json:"interfaceName,omitempty"`     // The name of the interface to use.
	MulticastGroup    string `json:"multicastGroup,omitempty"`    // 'Any', or the IP address of a multicast group.
	RendezvousPointID string `json:"rendezvousPointId,omitempty"` // The id.
	Serial            string `json:"serial,omitempty"`            // The serial.
}
type ResponseSwitchCreateNetworkSwitchRoutingMulticastRendezvousPoint struct {
	InterfaceIP       string `json:"interfaceIp,omitempty"`       // The IP address of the interface to use.
	InterfaceName     string `json:"interfaceName,omitempty"`     // The name of the interface to use.
	MulticastGroup    string `json:"multicastGroup,omitempty"`    // 'Any', or the IP address of a multicast group.
	RendezvousPointID string `json:"rendezvousPointId,omitempty"` // The id.
	Serial            string `json:"serial,omitempty"`            // The serial.
}
type ResponseSwitchGetNetworkSwitchRoutingMulticastRendezvousPoint struct {
	InterfaceIP       string `json:"interfaceIp,omitempty"`       // The IP address of the interface to use.
	InterfaceName     string `json:"interfaceName,omitempty"`     // The name of the interface to use.
	MulticastGroup    string `json:"multicastGroup,omitempty"`    // 'Any', or the IP address of a multicast group.
	RendezvousPointID string `json:"rendezvousPointId,omitempty"` // The id.
	Serial            string `json:"serial,omitempty"`            // The serial.
}
type ResponseSwitchUpdateNetworkSwitchRoutingMulticastRendezvousPoint struct {
	InterfaceIP       string `json:"interfaceIp,omitempty"`       // The IP address of the interface to use.
	InterfaceName     string `json:"interfaceName,omitempty"`     // The name of the interface to use.
	MulticastGroup    string `json:"multicastGroup,omitempty"`    // 'Any', or the IP address of a multicast group.
	RendezvousPointID string `json:"rendezvousPointId,omitempty"` // The id.
	Serial            string `json:"serial,omitempty"`            // The serial.
}
type ResponseSwitchGetNetworkSwitchRoutingOspf struct {
	Areas                    *[]ResponseSwitchGetNetworkSwitchRoutingOspfAreas              `json:"areas,omitempty"`                    // OSPF areas
	DeadTimerInSeconds       *int                                                           `json:"deadTimerInSeconds,omitempty"`       // Time interval to determine when the peer will be declared inactive/dead. Value must be between 1 and 65535
	Enabled                  *bool                                                          `json:"enabled,omitempty"`                  // Boolean value to enable or disable OSPF routing. OSPF routing is disabled by default.
	HelloTimerInSeconds      *int                                                           `json:"helloTimerInSeconds,omitempty"`      // Time interval in seconds at which hello packet will be sent to OSPF neighbors to maintain connectivity. Value must be between 1 and 255. Default is 10 seconds.
	Md5AuthenticationEnabled *bool                                                          `json:"md5AuthenticationEnabled,omitempty"` // Boolean value to enable or disable MD5 authentication. MD5 authentication is disabled by default.
	Md5AuthenticationKey     *ResponseSwitchGetNetworkSwitchRoutingOspfMd5AuthenticationKey `json:"md5AuthenticationKey,omitempty"`     // MD5 authentication credentials. This param is only relevant if md5AuthenticationEnabled is true
	V3                       *ResponseSwitchGetNetworkSwitchRoutingOspfV3                   `json:"v3,omitempty"`                       // OSPF v3 configuration
}
type ResponseSwitchGetNetworkSwitchRoutingOspfAreas struct {
	AreaID   *int   `json:"areaId,omitempty"`   // OSPF area ID
	AreaName string `json:"areaName,omitempty"` // Name of the OSPF area
	AreaType string `json:"areaType,omitempty"` // Area types in OSPF. Must be one of: ["normal", "stub", "nssa"]
}
type ResponseSwitchGetNetworkSwitchRoutingOspfMd5AuthenticationKey struct {
	ID         *int   `json:"id,omitempty"`         // MD5 authentication key index. Key index must be between 1 to 255
	Passphrase string `json:"passphrase,omitempty"` // MD5 authentication passphrase
}
type ResponseSwitchGetNetworkSwitchRoutingOspfV3 struct {
	Areas               *[]ResponseSwitchGetNetworkSwitchRoutingOspfV3Areas `json:"areas,omitempty"`               // OSPF v3 areas
	DeadTimerInSeconds  *int                                                `json:"deadTimerInSeconds,omitempty"`  // Time interval to determine when the peer will be declared inactive/dead. Value must be between 1 and 65535
	Enabled             *bool                                               `json:"enabled,omitempty"`             // Boolean value to enable or disable V3 OSPF routing. OSPF V3 routing is disabled by default.
	HelloTimerInSeconds *int                                                `json:"helloTimerInSeconds,omitempty"` // Time interval in seconds at which hello packet will be sent to OSPF neighbors to maintain connectivity. Value must be between 1 and 255. Default is 10 seconds.
}
type ResponseSwitchGetNetworkSwitchRoutingOspfV3Areas struct {
	AreaID   *int   `json:"areaId,omitempty"`   // OSPF area ID
	AreaName string `json:"areaName,omitempty"` // Name of the OSPF area
	AreaType string `json:"areaType,omitempty"` // Area types in OSPF. Must be one of: ["normal", "stub", "nssa"]
}
type ResponseSwitchUpdateNetworkSwitchRoutingOspf struct {
	Areas                    *[]ResponseSwitchUpdateNetworkSwitchRoutingOspfAreas              `json:"areas,omitempty"`                    // OSPF areas
	DeadTimerInSeconds       *int                                                              `json:"deadTimerInSeconds,omitempty"`       // Time interval to determine when the peer will be declared inactive/dead. Value must be between 1 and 65535
	Enabled                  *bool                                                             `json:"enabled,omitempty"`                  // Boolean value to enable or disable OSPF routing. OSPF routing is disabled by default.
	HelloTimerInSeconds      *int                                                              `json:"helloTimerInSeconds,omitempty"`      // Time interval in seconds at which hello packet will be sent to OSPF neighbors to maintain connectivity. Value must be between 1 and 255. Default is 10 seconds.
	Md5AuthenticationEnabled *bool                                                             `json:"md5AuthenticationEnabled,omitempty"` // Boolean value to enable or disable MD5 authentication. MD5 authentication is disabled by default.
	Md5AuthenticationKey     *ResponseSwitchUpdateNetworkSwitchRoutingOspfMd5AuthenticationKey `json:"md5AuthenticationKey,omitempty"`     // MD5 authentication credentials. This param is only relevant if md5AuthenticationEnabled is true
	V3                       *ResponseSwitchUpdateNetworkSwitchRoutingOspfV3                   `json:"v3,omitempty"`                       // OSPF v3 configuration
}
type ResponseSwitchUpdateNetworkSwitchRoutingOspfAreas struct {
	AreaID   *int   `json:"areaId,omitempty"`   // OSPF area ID
	AreaName string `json:"areaName,omitempty"` // Name of the OSPF area
	AreaType string `json:"areaType,omitempty"` // Area types in OSPF. Must be one of: ["normal", "stub", "nssa"]
}
type ResponseSwitchUpdateNetworkSwitchRoutingOspfMd5AuthenticationKey struct {
	ID         *int   `json:"id,omitempty"`         // MD5 authentication key index. Key index must be between 1 to 255
	Passphrase string `json:"passphrase,omitempty"` // MD5 authentication passphrase
}
type ResponseSwitchUpdateNetworkSwitchRoutingOspfV3 struct {
	Areas               *[]ResponseSwitchUpdateNetworkSwitchRoutingOspfV3Areas `json:"areas,omitempty"`               // OSPF v3 areas
	DeadTimerInSeconds  *int                                                   `json:"deadTimerInSeconds,omitempty"`  // Time interval to determine when the peer will be declared inactive/dead. Value must be between 1 and 65535
	Enabled             *bool                                                  `json:"enabled,omitempty"`             // Boolean value to enable or disable V3 OSPF routing. OSPF V3 routing is disabled by default.
	HelloTimerInSeconds *int                                                   `json:"helloTimerInSeconds,omitempty"` // Time interval in seconds at which hello packet will be sent to OSPF neighbors to maintain connectivity. Value must be between 1 and 255. Default is 10 seconds.
}
type ResponseSwitchUpdateNetworkSwitchRoutingOspfV3Areas struct {
	AreaID   *int   `json:"areaId,omitempty"`   // OSPF area ID
	AreaName string `json:"areaName,omitempty"` // Name of the OSPF area
	AreaType string `json:"areaType,omitempty"` // Area types in OSPF. Must be one of: ["normal", "stub", "nssa"]
}
type ResponseSwitchGetNetworkSwitchSettings struct {
	MacBlocklist         *ResponseSwitchGetNetworkSwitchSettingsMacBlocklist         `json:"macBlocklist,omitempty"`         // MAC blocklist
	PowerExceptions      *[]ResponseSwitchGetNetworkSwitchSettingsPowerExceptions    `json:"powerExceptions,omitempty"`      // Exceptions on a per switch basis to "useCombinedPower"
	UplinkClientSampling *ResponseSwitchGetNetworkSwitchSettingsUplinkClientSampling `json:"uplinkClientSampling,omitempty"` // Uplink client sampling
	UseCombinedPower     *bool                                                       `json:"useCombinedPower,omitempty"`     // The use Combined Power as the default behavior of secondary power supplies on supported devices.
	VLAN                 *int                                                        `json:"vlan,omitempty"`                 // Management VLAN
}
type ResponseSwitchGetNetworkSwitchSettingsMacBlocklist struct {
	Enabled *bool `json:"enabled,omitempty"` // Enable MAC blocklist for switches in the network
}
type ResponseSwitchGetNetworkSwitchSettingsPowerExceptions struct {
	PowerType string `json:"powerType,omitempty"` // Per switch exception (combined, redundant, useNetworkSetting)
	Serial    string `json:"serial,omitempty"`    // Serial number of the switch
}
type ResponseSwitchGetNetworkSwitchSettingsUplinkClientSampling struct {
	Enabled *bool `json:"enabled,omitempty"` // Enable client sampling on uplink
}
type ResponseSwitchUpdateNetworkSwitchSettings struct {
	MacBlocklist         *ResponseSwitchUpdateNetworkSwitchSettingsMacBlocklist         `json:"macBlocklist,omitempty"`         // MAC blocklist
	PowerExceptions      *[]ResponseSwitchUpdateNetworkSwitchSettingsPowerExceptions    `json:"powerExceptions,omitempty"`      // Exceptions on a per switch basis to "useCombinedPower"
	UplinkClientSampling *ResponseSwitchUpdateNetworkSwitchSettingsUplinkClientSampling `json:"uplinkClientSampling,omitempty"` // Uplink client sampling
	UseCombinedPower     *bool                                                          `json:"useCombinedPower,omitempty"`     // The use Combined Power as the default behavior of secondary power supplies on supported devices.
	VLAN                 *int                                                           `json:"vlan,omitempty"`                 // Management VLAN
}
type ResponseSwitchUpdateNetworkSwitchSettingsMacBlocklist struct {
	Enabled *bool `json:"enabled,omitempty"` // Enable MAC blocklist for switches in the network
}
type ResponseSwitchUpdateNetworkSwitchSettingsPowerExceptions struct {
	PowerType string `json:"powerType,omitempty"` // Per switch exception (combined, redundant, useNetworkSetting)
	Serial    string `json:"serial,omitempty"`    // Serial number of the switch
}
type ResponseSwitchUpdateNetworkSwitchSettingsUplinkClientSampling struct {
	Enabled *bool `json:"enabled,omitempty"` // Enable client sampling on uplink
}
type ResponseSwitchGetNetworkSwitchStacks []ResponseItemSwitchGetNetworkSwitchStacks // Array of ResponseSwitchGetNetworkSwitchStacks
type ResponseItemSwitchGetNetworkSwitchStacks struct {
	ID            string                                             `json:"id,omitempty"`            // ID of the Switch stack
	IsMonitorOnly *bool                                              `json:"isMonitorOnly,omitempty"` // Tells if stack is Monitored Stack.
	Members       *[]ResponseItemSwitchGetNetworkSwitchStacksMembers `json:"members,omitempty"`       // Members of the Stack
	Name          string                                             `json:"name,omitempty"`          // Name of the Switch stack
	Serials       []string                                           `json:"serials,omitempty"`       // Serials of the switches in the switch stack
}
type ResponseItemSwitchGetNetworkSwitchStacksMembers struct {
	Mac    string `json:"mac,omitempty"`    // MAC address of the device
	Model  string `json:"model,omitempty"`  // Model of the device
	Name   string `json:"name,omitempty"`   // Name of the device
	Role   string `json:"role,omitempty"`   // Role of the device
	Serial string `json:"serial,omitempty"` // Serial number of the device
}
type ResponseSwitchCreateNetworkSwitchStack struct {
	ID            string                                           `json:"id,omitempty"`            // ID of the Switch stack
	IsMonitorOnly *bool                                            `json:"isMonitorOnly,omitempty"` // Tells if stack is Monitored Stack.
	Members       *[]ResponseSwitchCreateNetworkSwitchStackMembers `json:"members,omitempty"`       // Members of the Stack
	Name          string                                           `json:"name,omitempty"`          // Name of the Switch stack
	Serials       []string                                         `json:"serials,omitempty"`       // Serials of the switches in the switch stack
}
type ResponseSwitchCreateNetworkSwitchStackMembers struct {
	Mac    string `json:"mac,omitempty"`    // MAC address of the device
	Model  string `json:"model,omitempty"`  // Model of the device
	Name   string `json:"name,omitempty"`   // Name of the device
	Role   string `json:"role,omitempty"`   // Role of the device
	Serial string `json:"serial,omitempty"` // Serial number of the device
}
type ResponseSwitchGetNetworkSwitchStack struct {
	ID            string                                        `json:"id,omitempty"`            // ID of the Switch stack
	IsMonitorOnly *bool                                         `json:"isMonitorOnly,omitempty"` // Tells if stack is Monitored Stack.
	Members       *[]ResponseSwitchGetNetworkSwitchStackMembers `json:"members,omitempty"`       // Members of the Stack
	Name          string                                        `json:"name,omitempty"`          // Name of the Switch stack
	Serials       []string                                      `json:"serials,omitempty"`       // Serials of the switches in the switch stack
}
type ResponseSwitchGetNetworkSwitchStackMembers struct {
	Mac    string `json:"mac,omitempty"`    // MAC address of the device
	Model  string `json:"model,omitempty"`  // Model of the device
	Name   string `json:"name,omitempty"`   // Name of the device
	Role   string `json:"role,omitempty"`   // Role of the device
	Serial string `json:"serial,omitempty"` // Serial number of the device
}
type ResponseSwitchAddNetworkSwitchStack struct {
	ID            string                                        `json:"id,omitempty"`            // ID of the Switch stack
	IsMonitorOnly *bool                                         `json:"isMonitorOnly,omitempty"` // Tells if stack is Monitored Stack.
	Members       *[]ResponseSwitchAddNetworkSwitchStackMembers `json:"members,omitempty"`       // Members of the Stack
	Name          string                                        `json:"name,omitempty"`          // Name of the Switch stack
	Serials       []string                                      `json:"serials,omitempty"`       // Serials of the switches in the switch stack
}
type ResponseSwitchAddNetworkSwitchStackMembers struct {
	Mac    string `json:"mac,omitempty"`    // MAC address of the device
	Model  string `json:"model,omitempty"`  // Model of the device
	Name   string `json:"name,omitempty"`   // Name of the device
	Role   string `json:"role,omitempty"`   // Role of the device
	Serial string `json:"serial,omitempty"` // Serial number of the device
}
type ResponseSwitchRemoveNetworkSwitchStack struct {
	ID            string                                           `json:"id,omitempty"`            // ID of the Switch stack
	IsMonitorOnly *bool                                            `json:"isMonitorOnly,omitempty"` // Tells if stack is Monitored Stack.
	Members       *[]ResponseSwitchRemoveNetworkSwitchStackMembers `json:"members,omitempty"`       // Members of the Stack
	Name          string                                           `json:"name,omitempty"`          // Name of the Switch stack
	Serials       []string                                         `json:"serials,omitempty"`       // Serials of the switches in the switch stack
}
type ResponseSwitchRemoveNetworkSwitchStackMembers struct {
	Mac    string `json:"mac,omitempty"`    // MAC address of the device
	Model  string `json:"model,omitempty"`  // Model of the device
	Name   string `json:"name,omitempty"`   // Name of the device
	Role   string `json:"role,omitempty"`   // Role of the device
	Serial string `json:"serial,omitempty"` // Serial number of the device
}
type ResponseSwitchGetNetworkSwitchStackRoutingInterfaces []ResponseItemSwitchGetNetworkSwitchStackRoutingInterfaces // Array of ResponseSwitchGetNetworkSwitchStackRoutingInterfaces
type ResponseItemSwitchGetNetworkSwitchStackRoutingInterfaces struct {
	DefaultGateway   string                                                                `json:"defaultGateway,omitempty"`   // IPv4 default gateway
	InterfaceID      string                                                                `json:"interfaceId,omitempty"`      // The id
	InterfaceIP      string                                                                `json:"interfaceIp,omitempty"`      // IPv4 address
	IPv6             *ResponseItemSwitchGetNetworkSwitchStackRoutingInterfacesIPv6         `json:"ipv6,omitempty"`             // IPv6 addressing
	MulticastRouting string                                                                `json:"multicastRouting,omitempty"` // Multicast routing status
	Name             string                                                                `json:"name,omitempty"`             // The name
	OspfSettings     *ResponseItemSwitchGetNetworkSwitchStackRoutingInterfacesOspfSettings `json:"ospfSettings,omitempty"`     // IPv4 OSPF Settings
	OspfV3           *ResponseItemSwitchGetNetworkSwitchStackRoutingInterfacesOspfV3       `json:"ospfV3,omitempty"`           // IPv6 OSPF Settings
	Subnet           string                                                                `json:"subnet,omitempty"`           // IPv4 subnet
	UplinkV4         *bool                                                                 `json:"uplinkV4,omitempty"`         // Whether this is the switch's IPv4 uplink
	UplinkV6         *bool                                                                 `json:"uplinkV6,omitempty"`         // Whether this is the switch's IPv6 uplink
	VLANID           *int                                                                  `json:"vlanId,omitempty"`           // VLAN id
}
type ResponseItemSwitchGetNetworkSwitchStackRoutingInterfacesIPv6 struct {
	Address        string `json:"address,omitempty"`        // IPv6 address
	AssignmentMode string `json:"assignmentMode,omitempty"` // Assignment mode
	Gateway        string `json:"gateway,omitempty"`        // IPv6 gateway
	Prefix         string `json:"prefix,omitempty"`         // IPv6 subnet
}
type ResponseItemSwitchGetNetworkSwitchStackRoutingInterfacesOspfSettings struct {
	Area             string `json:"area,omitempty"`             // Area id
	Cost             *int   `json:"cost,omitempty"`             // OSPF Cost
	IsPassiveEnabled *bool  `json:"isPassiveEnabled,omitempty"` // Disable sending Hello packets on this interface's IPv4 area
}
type ResponseItemSwitchGetNetworkSwitchStackRoutingInterfacesOspfV3 struct {
	Area             string `json:"area,omitempty"`             // Area id
	Cost             *int   `json:"cost,omitempty"`             // OSPF Cost
	IsPassiveEnabled *bool  `json:"isPassiveEnabled,omitempty"` // Disable sending Hello packets on this interface's IPv6 area
}
type ResponseSwitchCreateNetworkSwitchStackRoutingInterface struct {
	DefaultGateway   string                                                              `json:"defaultGateway,omitempty"`   // IPv4 default gateway
	InterfaceID      string                                                              `json:"interfaceId,omitempty"`      // The id
	InterfaceIP      string                                                              `json:"interfaceIp,omitempty"`      // IPv4 address
	IPv6             *ResponseSwitchCreateNetworkSwitchStackRoutingInterfaceIPv6         `json:"ipv6,omitempty"`             // IPv6 addressing
	MulticastRouting string                                                              `json:"multicastRouting,omitempty"` // Multicast routing status
	Name             string                                                              `json:"name,omitempty"`             // The name
	OspfSettings     *ResponseSwitchCreateNetworkSwitchStackRoutingInterfaceOspfSettings `json:"ospfSettings,omitempty"`     // IPv4 OSPF Settings
	OspfV3           *ResponseSwitchCreateNetworkSwitchStackRoutingInterfaceOspfV3       `json:"ospfV3,omitempty"`           // IPv6 OSPF Settings
	Subnet           string                                                              `json:"subnet,omitempty"`           // IPv4 subnet
	UplinkV4         *bool                                                               `json:"uplinkV4,omitempty"`         // Whether this is the switch's IPv4 uplink
	UplinkV6         *bool                                                               `json:"uplinkV6,omitempty"`         // Whether this is the switch's IPv6 uplink
	VLANID           *int                                                                `json:"vlanId,omitempty"`           // VLAN id
}
type ResponseSwitchCreateNetworkSwitchStackRoutingInterfaceIPv6 struct {
	Address        string `json:"address,omitempty"`        // IPv6 address
	AssignmentMode string `json:"assignmentMode,omitempty"` // Assignment mode
	Gateway        string `json:"gateway,omitempty"`        // IPv6 gateway
	Prefix         string `json:"prefix,omitempty"`         // IPv6 subnet
}
type ResponseSwitchCreateNetworkSwitchStackRoutingInterfaceOspfSettings struct {
	Area             string `json:"area,omitempty"`             // Area id
	Cost             *int   `json:"cost,omitempty"`             // OSPF Cost
	IsPassiveEnabled *bool  `json:"isPassiveEnabled,omitempty"` // Disable sending Hello packets on this interface's IPv4 area
}
type ResponseSwitchCreateNetworkSwitchStackRoutingInterfaceOspfV3 struct {
	Area             string `json:"area,omitempty"`             // Area id
	Cost             *int   `json:"cost,omitempty"`             // OSPF Cost
	IsPassiveEnabled *bool  `json:"isPassiveEnabled,omitempty"` // Disable sending Hello packets on this interface's IPv6 area
}
type ResponseSwitchGetNetworkSwitchStackRoutingInterface struct {
	DefaultGateway   string                                                           `json:"defaultGateway,omitempty"`   // IPv4 default gateway
	InterfaceID      string                                                           `json:"interfaceId,omitempty"`      // The id
	InterfaceIP      string                                                           `json:"interfaceIp,omitempty"`      // IPv4 address
	IPv6             *ResponseSwitchGetNetworkSwitchStackRoutingInterfaceIPv6         `json:"ipv6,omitempty"`             // IPv6 addressing
	MulticastRouting string                                                           `json:"multicastRouting,omitempty"` // Multicast routing status
	Name             string                                                           `json:"name,omitempty"`             // The name
	OspfSettings     *ResponseSwitchGetNetworkSwitchStackRoutingInterfaceOspfSettings `json:"ospfSettings,omitempty"`     // IPv4 OSPF Settings
	OspfV3           *ResponseSwitchGetNetworkSwitchStackRoutingInterfaceOspfV3       `json:"ospfV3,omitempty"`           // IPv6 OSPF Settings
	Subnet           string                                                           `json:"subnet,omitempty"`           // IPv4 subnet
	UplinkV4         *bool                                                            `json:"uplinkV4,omitempty"`         // Whether this is the switch's IPv4 uplink
	UplinkV6         *bool                                                            `json:"uplinkV6,omitempty"`         // Whether this is the switch's IPv6 uplink
	VLANID           *int                                                             `json:"vlanId,omitempty"`           // VLAN id
}
type ResponseSwitchGetNetworkSwitchStackRoutingInterfaceIPv6 struct {
	Address        string `json:"address,omitempty"`        // IPv6 address
	AssignmentMode string `json:"assignmentMode,omitempty"` // Assignment mode
	Gateway        string `json:"gateway,omitempty"`        // IPv6 gateway
	Prefix         string `json:"prefix,omitempty"`         // IPv6 subnet
}
type ResponseSwitchGetNetworkSwitchStackRoutingInterfaceOspfSettings struct {
	Area             string `json:"area,omitempty"`             // Area id
	Cost             *int   `json:"cost,omitempty"`             // OSPF Cost
	IsPassiveEnabled *bool  `json:"isPassiveEnabled,omitempty"` // Disable sending Hello packets on this interface's IPv4 area
}
type ResponseSwitchGetNetworkSwitchStackRoutingInterfaceOspfV3 struct {
	Area             string `json:"area,omitempty"`             // Area id
	Cost             *int   `json:"cost,omitempty"`             // OSPF Cost
	IsPassiveEnabled *bool  `json:"isPassiveEnabled,omitempty"` // Disable sending Hello packets on this interface's IPv6 area
}
type ResponseSwitchUpdateNetworkSwitchStackRoutingInterface struct {
	InterfaceID      string                                                              `json:"interfaceId,omitempty"`      // The id
	InterfaceIP      string                                                              `json:"interfaceIp,omitempty"`      // IPv4 address
	IPv6             *ResponseSwitchUpdateNetworkSwitchStackRoutingInterfaceIPv6         `json:"ipv6,omitempty"`             // IPv6 addressing
	MulticastRouting string                                                              `json:"multicastRouting,omitempty"` // Multicast routing status
	Name             string                                                              `json:"name,omitempty"`             // The name
	OspfSettings     *ResponseSwitchUpdateNetworkSwitchStackRoutingInterfaceOspfSettings `json:"ospfSettings,omitempty"`     // IPv4 OSPF Settings
	OspfV3           *ResponseSwitchUpdateNetworkSwitchStackRoutingInterfaceOspfV3       `json:"ospfV3,omitempty"`           // IPv6 OSPF Settings
	Subnet           string                                                              `json:"subnet,omitempty"`           // IPv4 subnet
	UplinkV4         *bool                                                               `json:"uplinkV4,omitempty"`         // Whether this is the switch's IPv4 uplink
	UplinkV6         *bool                                                               `json:"uplinkV6,omitempty"`         // Whether this is the switch's IPv6 uplink
	VLANID           *int                                                                `json:"vlanId,omitempty"`           // VLAN id
}
type ResponseSwitchUpdateNetworkSwitchStackRoutingInterfaceIPv6 struct {
	Address        string `json:"address,omitempty"`        // IPv6 address
	AssignmentMode string `json:"assignmentMode,omitempty"` // Assignment mode
	Gateway        string `json:"gateway,omitempty"`        // IPv6 gateway
	Prefix         string `json:"prefix,omitempty"`         // IPv6 subnet
}
type ResponseSwitchUpdateNetworkSwitchStackRoutingInterfaceOspfSettings struct {
	Area             string `json:"area,omitempty"`             // Area id
	Cost             *int   `json:"cost,omitempty"`             // OSPF Cost
	IsPassiveEnabled *bool  `json:"isPassiveEnabled,omitempty"` // Disable sending Hello packets on this interface's IPv4 area
}
type ResponseSwitchUpdateNetworkSwitchStackRoutingInterfaceOspfV3 struct {
	Area             string `json:"area,omitempty"`             // Area id
	Cost             *int   `json:"cost,omitempty"`             // OSPF Cost
	IsPassiveEnabled *bool  `json:"isPassiveEnabled,omitempty"` // Disable sending Hello packets on this interface's IPv6 area
}
type ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcp struct {
	BootFileName         string                                                                       `json:"bootFileName,omitempty"`         // The PXE boot server file name for the DHCP server running on the switch stack interface
	BootNextServer       string                                                                       `json:"bootNextServer,omitempty"`       // The PXE boot server IP for the DHCP server running on the switch stack interface
	BootOptionsEnabled   *bool                                                                        `json:"bootOptionsEnabled,omitempty"`   // Enable DHCP boot options to provide PXE boot options configs for the dhcp server running on the switch stack interface
	DhcpLeaseTime        string                                                                       `json:"dhcpLeaseTime,omitempty"`        // The DHCP lease time config for the dhcp server running on the switch stack interface ('30 minutes', '1 hour', '4 hours', '12 hours', '1 day' or '1 week')
	DhcpMode             string                                                                       `json:"dhcpMode,omitempty"`             // The DHCP mode options for the switch stack interface ('dhcpDisabled', 'dhcpRelay' or 'dhcpServer')
	DhcpOptions          *[]ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpDhcpOptions        `json:"dhcpOptions,omitempty"`          // Array of DHCP options consisting of code, type and value for the DHCP server running on the switch stack interface
	DhcpRelayServerIPs   []string                                                                     `json:"dhcpRelayServerIps,omitempty"`   // The DHCP relay server IPs to which DHCP packets would get relayed for the switch stack interface
	DNSCustomNameservers []string                                                                     `json:"dnsCustomNameservers,omitempty"` // The DHCP name server IPs when DHCP name server option is 'custom'
	DNSNameserversOption string                                                                       `json:"dnsNameserversOption,omitempty"` // The DHCP name server option for the dhcp server running on the switch stack interface ('googlePublicDns', 'openDns' or 'custom')
	FixedIPAssignments   *[]ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpFixedIPAssignments `json:"fixedIpAssignments,omitempty"`   // Array of DHCP reserved IP assignments for the DHCP server running on the switch stack interface
	ReservedIPRanges     *[]ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpReservedIPRanges   `json:"reservedIpRanges,omitempty"`     // Array of DHCP reserved IP assignments for the DHCP server running on the switch stack interface
}
type ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpDhcpOptions struct {
	Code  string `json:"code,omitempty"`  // The code for DHCP option which should be from 2 to 254
	Type  string `json:"type,omitempty"`  // The type of the DHCP option which should be one of ('text', 'ip', 'integer' or 'hex')
	Value string `json:"value,omitempty"` // The value of the DHCP option
}
type ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpFixedIPAssignments struct {
	IP   string `json:"ip,omitempty"`   // The IP address of the client which has fixed IP address assigned to it
	Mac  string `json:"mac,omitempty"`  // The MAC address of the client which has fixed IP address
	Name string `json:"name,omitempty"` // The name of the client which has fixed IP address
}
type ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpReservedIPRanges struct {
	Comment string `json:"comment,omitempty"` // The comment for the reserved IP range
	End     string `json:"end,omitempty"`     // The ending IP address of the reserved IP range
	Start   string `json:"start,omitempty"`   // The starting IP address of the reserved IP range
}
type ResponseSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcp struct {
	BootFileName         string                                                                          `json:"bootFileName,omitempty"`         // The PXE boot server file name for the DHCP server running on the switch stack interface
	BootNextServer       string                                                                          `json:"bootNextServer,omitempty"`       // The PXE boot server IP for the DHCP server running on the switch stack interface
	BootOptionsEnabled   *bool                                                                           `json:"bootOptionsEnabled,omitempty"`   // Enable DHCP boot options to provide PXE boot options configs for the dhcp server running on the switch stack interface
	DhcpLeaseTime        string                                                                          `json:"dhcpLeaseTime,omitempty"`        // The DHCP lease time config for the dhcp server running on the switch stack interface ('30 minutes', '1 hour', '4 hours', '12 hours', '1 day' or '1 week')
	DhcpMode             string                                                                          `json:"dhcpMode,omitempty"`             // The DHCP mode options for the switch stack interface ('dhcpDisabled', 'dhcpRelay' or 'dhcpServer')
	DhcpOptions          *[]ResponseSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcpDhcpOptions        `json:"dhcpOptions,omitempty"`          // Array of DHCP options consisting of code, type and value for the DHCP server running on the switch stack interface
	DhcpRelayServerIPs   []string                                                                        `json:"dhcpRelayServerIps,omitempty"`   // The DHCP relay server IPs to which DHCP packets would get relayed for the switch stack interface
	DNSCustomNameservers []string                                                                        `json:"dnsCustomNameservers,omitempty"` // The DHCP name server IPs when DHCP name server option is 'custom'
	DNSNameserversOption string                                                                          `json:"dnsNameserversOption,omitempty"` // The DHCP name server option for the dhcp server running on the switch stack interface ('googlePublicDns', 'openDns' or 'custom')
	FixedIPAssignments   *[]ResponseSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcpFixedIPAssignments `json:"fixedIpAssignments,omitempty"`   // Array of DHCP reserved IP assignments for the DHCP server running on the switch stack interface
	ReservedIPRanges     *[]ResponseSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcpReservedIPRanges   `json:"reservedIpRanges,omitempty"`     // Array of DHCP reserved IP assignments for the DHCP server running on the switch stack interface
}
type ResponseSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcpDhcpOptions struct {
	Code  string `json:"code,omitempty"`  // The code for DHCP option which should be from 2 to 254
	Type  string `json:"type,omitempty"`  // The type of the DHCP option which should be one of ('text', 'ip', 'integer' or 'hex')
	Value string `json:"value,omitempty"` // The value of the DHCP option
}
type ResponseSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcpFixedIPAssignments struct {
	IP   string `json:"ip,omitempty"`   // The IP address of the client which has fixed IP address assigned to it
	Mac  string `json:"mac,omitempty"`  // The MAC address of the client which has fixed IP address
	Name string `json:"name,omitempty"` // The name of the client which has fixed IP address
}
type ResponseSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcpReservedIPRanges struct {
	Comment string `json:"comment,omitempty"` // The comment for the reserved IP range
	End     string `json:"end,omitempty"`     // The ending IP address of the reserved IP range
	Start   string `json:"start,omitempty"`   // The starting IP address of the reserved IP range
}
type ResponseSwitchGetNetworkSwitchStackRoutingStaticRoutes []ResponseItemSwitchGetNetworkSwitchStackRoutingStaticRoutes // Array of ResponseSwitchGetNetworkSwitchStackRoutingStaticRoutes
type ResponseItemSwitchGetNetworkSwitchStackRoutingStaticRoutes struct {
	AdvertiseViaOspfEnabled     *bool  `json:"advertiseViaOspfEnabled,omitempty"`     // Option to advertise static routes via OSPF
	ManagementNextHop           string `json:"managementNextHop,omitempty"`           // Optional fallback IP address for management traffic
	Name                        string `json:"name,omitempty"`                        // The name or description of the layer 3 static route
	NextHopIP                   string `json:"nextHopIp,omitempty"`                   // The IP address of the router to which traffic for this destination network should be sent
	PreferOverOspfRoutesEnabled *bool  `json:"preferOverOspfRoutesEnabled,omitempty"` // Option to prefer static routes over OSPF routes
	StaticRouteID               string `json:"staticRouteId,omitempty"`               // The identifier of a layer 3 static route
	Subnet                      string `json:"subnet,omitempty"`                      // The IP address of the subnetwork specified in CIDR notation (ex. 1.2.3.0/24)
}
type ResponseSwitchCreateNetworkSwitchStackRoutingStaticRoute struct {
	AdvertiseViaOspfEnabled     *bool  `json:"advertiseViaOspfEnabled,omitempty"`     // Option to advertise static routes via OSPF
	ManagementNextHop           string `json:"managementNextHop,omitempty"`           // Optional fallback IP address for management traffic
	Name                        string `json:"name,omitempty"`                        // The name or description of the layer 3 static route
	NextHopIP                   string `json:"nextHopIp,omitempty"`                   // The IP address of the router to which traffic for this destination network should be sent
	PreferOverOspfRoutesEnabled *bool  `json:"preferOverOspfRoutesEnabled,omitempty"` // Option to prefer static routes over OSPF routes
	StaticRouteID               string `json:"staticRouteId,omitempty"`               // The identifier of a layer 3 static route
	Subnet                      string `json:"subnet,omitempty"`                      // The IP address of the subnetwork specified in CIDR notation (ex. 1.2.3.0/24)
}
type ResponseSwitchGetNetworkSwitchStackRoutingStaticRoute struct {
	AdvertiseViaOspfEnabled     *bool  `json:"advertiseViaOspfEnabled,omitempty"`     // Option to advertise static routes via OSPF
	ManagementNextHop           string `json:"managementNextHop,omitempty"`           // Optional fallback IP address for management traffic
	Name                        string `json:"name,omitempty"`                        // The name or description of the layer 3 static route
	NextHopIP                   string `json:"nextHopIp,omitempty"`                   // The IP address of the router to which traffic for this destination network should be sent
	PreferOverOspfRoutesEnabled *bool  `json:"preferOverOspfRoutesEnabled,omitempty"` // Option to prefer static routes over OSPF routes
	StaticRouteID               string `json:"staticRouteId,omitempty"`               // The identifier of a layer 3 static route
	Subnet                      string `json:"subnet,omitempty"`                      // The IP address of the subnetwork specified in CIDR notation (ex. 1.2.3.0/24)
}
type ResponseSwitchUpdateNetworkSwitchStackRoutingStaticRoute struct {
	AdvertiseViaOspfEnabled     *bool  `json:"advertiseViaOspfEnabled,omitempty"`     // Option to advertise static routes via OSPF
	ManagementNextHop           string `json:"managementNextHop,omitempty"`           // Optional fallback IP address for management traffic
	Name                        string `json:"name,omitempty"`                        // The name or description of the layer 3 static route
	NextHopIP                   string `json:"nextHopIp,omitempty"`                   // The IP address of the router to which traffic for this destination network should be sent
	PreferOverOspfRoutesEnabled *bool  `json:"preferOverOspfRoutesEnabled,omitempty"` // Option to prefer static routes over OSPF routes
	StaticRouteID               string `json:"staticRouteId,omitempty"`               // The identifier of a layer 3 static route
	Subnet                      string `json:"subnet,omitempty"`                      // The IP address of the subnetwork specified in CIDR notation (ex. 1.2.3.0/24)
}
type ResponseSwitchGetNetworkSwitchStormControl struct {
	BroadcastThreshold                   *int     `json:"broadcastThreshold,omitempty"`                   // Broadcast threshold.
	MulticastThreshold                   *int     `json:"multicastThreshold,omitempty"`                   // Multicast threshold.
	TreatTheseTrafficTypesAsOneThreshold []string `json:"treatTheseTrafficTypesAsOneThreshold,omitempty"` // Grouped traffic types
	UnknownUnicastThreshold              *int     `json:"unknownUnicastThreshold,omitempty"`              // Unknown Unicast threshold.
}
type ResponseSwitchUpdateNetworkSwitchStormControl struct {
	BroadcastThreshold                   *int     `json:"broadcastThreshold,omitempty"`                   // Broadcast threshold.
	MulticastThreshold                   *int     `json:"multicastThreshold,omitempty"`                   // Multicast threshold.
	TreatTheseTrafficTypesAsOneThreshold []string `json:"treatTheseTrafficTypesAsOneThreshold,omitempty"` // Grouped traffic types
	UnknownUnicastThreshold              *int     `json:"unknownUnicastThreshold,omitempty"`              // Unknown Unicast threshold.
}
type ResponseSwitchGetNetworkSwitchStp struct {
	RstpEnabled       *bool                                                 `json:"rstpEnabled,omitempty"`       // The spanning tree protocol status in network
	StpBridgePriority *[]ResponseSwitchGetNetworkSwitchStpStpBridgePriority `json:"stpBridgePriority,omitempty"` // STP bridge priority for switches/stacks or switch templates. An empty array will clear the STP bridge priority settings.
}
type ResponseSwitchGetNetworkSwitchStpStpBridgePriority struct {
	Stacks         []string `json:"stacks,omitempty"`         // List of stack IDs
	StpPriority    *int     `json:"stpPriority,omitempty"`    // STP priority for switch, stacks, or switch templates
	SwitchProfiles []string `json:"switchProfiles,omitempty"` // List of switch template IDs
	Switches       []string `json:"switches,omitempty"`       // List of switch serial numbers
}
type ResponseSwitchUpdateNetworkSwitchStp struct {
	RstpEnabled       *bool                                                    `json:"rstpEnabled,omitempty"`       // The spanning tree protocol status in network
	StpBridgePriority *[]ResponseSwitchUpdateNetworkSwitchStpStpBridgePriority `json:"stpBridgePriority,omitempty"` // STP bridge priority for switches/stacks or switch templates. An empty array will clear the STP bridge priority settings.
}
type ResponseSwitchUpdateNetworkSwitchStpStpBridgePriority struct {
	Stacks         []string `json:"stacks,omitempty"`         // List of stack IDs
	StpPriority    *int     `json:"stpPriority,omitempty"`    // STP priority for switch, stacks, or switch templates
	SwitchProfiles []string `json:"switchProfiles,omitempty"` // List of switch template IDs
	Switches       []string `json:"switches,omitempty"`       // List of switch serial numbers
}
type ResponseSwitchGetOrganizationConfigTemplateSwitchProfiles []ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfiles // Array of ResponseSwitchGetOrganizationConfigTemplateSwitchProfiles
type ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfiles struct {
	Model           string `json:"model,omitempty"`           // Switch model
	Name            string `json:"name,omitempty"`            // Switch template name
	SwitchProfileID string `json:"switchProfileId,omitempty"` // Switch template id
}
type ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePorts []ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePorts // Array of ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePorts
type ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePorts struct {
	AccessPolicyNumber          *int                                                                               `json:"accessPolicyNumber,omitempty"`          // The number of a custom access policy to configure on the switch template port. Only applicable when 'accessPolicyType' is 'Custom access policy'.
	AccessPolicyType            string                                                                             `json:"accessPolicyType,omitempty"`            // The type of the access policy of the switch template port. Only applicable to access ports. Can be one of 'Open', 'Custom access policy', 'MAC allow list' or 'Sticky MAC allow list'.
	AllowedVLANs                string                                                                             `json:"allowedVlans,omitempty"`                // The VLANs allowed on the switch template port. Only applicable to trunk ports.
	DaiTrusted                  *bool                                                                              `json:"daiTrusted,omitempty"`                  // If true, ARP packets for this port will be considered trusted, and Dynamic ARP Inspection will allow the traffic.
	Dot3Az                      *ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePortsDot3Az           `json:"dot3az,omitempty"`                      // dot3az settings for the port
	Enabled                     *bool                                                                              `json:"enabled,omitempty"`                     // The status of the switch template port.
	FlexibleStackingEnabled     *bool                                                                              `json:"flexibleStackingEnabled,omitempty"`     // For supported switches (e.g. MS420/MS425), whether or not the port has flexible stacking enabled.
	IsolationEnabled            *bool                                                                              `json:"isolationEnabled,omitempty"`            // The isolation status of the switch template port.
	LinkNegotiation             string                                                                             `json:"linkNegotiation,omitempty"`             // The link speed for the switch template port.
	LinkNegotiationCapabilities []string                                                                           `json:"linkNegotiationCapabilities,omitempty"` // Available link speeds for the switch template port.
	MacAllowList                []string                                                                           `json:"macAllowList,omitempty"`                // Only devices with MAC addresses specified in this list will have access to this port. Up to 20 MAC addresses can be defined. Only applicable when 'accessPolicyType' is 'MAC allow list'.
	Mirror                      *ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePortsMirror           `json:"mirror,omitempty"`                      // Port mirror
	Module                      *ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePortsModule           `json:"module,omitempty"`                      // Expansion module
	Name                        string                                                                             `json:"name,omitempty"`                        // The name of the switch template port.
	PoeEnabled                  *bool                                                                              `json:"poeEnabled,omitempty"`                  // The PoE status of the switch template port.
	PortID                      string                                                                             `json:"portId,omitempty"`                      // The identifier of the switch template port.
	PortScheduleID              string                                                                             `json:"portScheduleId,omitempty"`              // The ID of the port schedule. A value of null will clear the port schedule.
	Profile                     *ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePortsProfile          `json:"profile,omitempty"`                     // Profile attributes
	RstpEnabled                 *bool                                                                              `json:"rstpEnabled,omitempty"`                 // The rapid spanning tree protocol status.
	Schedule                    *ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePortsSchedule         `json:"schedule,omitempty"`                    // The port schedule data.
	StackwiseVirtual            *ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePortsStackwiseVirtual `json:"stackwiseVirtual,omitempty"`            // Stackwise Virtual settings for the port
	StickyMacAllowList          []string                                                                           `json:"stickyMacAllowList,omitempty"`          // The initial list of MAC addresses for sticky Mac allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
	StickyMacAllowListLimit     *int                                                                               `json:"stickyMacAllowListLimit,omitempty"`     // The maximum number of MAC addresses for sticky MAC allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
	StormControlEnabled         *bool                                                                              `json:"stormControlEnabled,omitempty"`         // The storm control status of the switch template port.
	StpGuard                    string                                                                             `json:"stpGuard,omitempty"`                    // The state of the STP guard ('disabled', 'root guard', 'bpdu guard' or 'loop guard').
	Tags                        []string                                                                           `json:"tags,omitempty"`                        // The list of tags of the switch template port.
	Type                        string                                                                             `json:"type,omitempty"`                        // The type of the switch template port ('trunk', 'access' or 'stack').
	Udld                        string                                                                             `json:"udld,omitempty"`                        // The action to take when Unidirectional Link is detected (Alert only, Enforce). Default configuration is Alert only.
	VLAN                        *int                                                                               `json:"vlan,omitempty"`                        // The VLAN of the switch template port. For a trunk port, this is the native VLAN. A null value will clear the value set for trunk ports.
	VoiceVLAN                   *int                                                                               `json:"voiceVlan,omitempty"`                   // The voice VLAN of the switch template port. Only applicable to access ports.
}
type ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePortsDot3Az struct {
	Enabled *bool `json:"enabled,omitempty"` // The Energy Efficient Ethernet status of the switch template port.
}
type ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePortsMirror struct {
	Mode string `json:"mode,omitempty"` // The port mirror mode. Can be one of ('Destination port', 'Source port' or 'Not mirroring traffic').
}
type ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePortsModule struct {
	Model string `json:"model,omitempty"` // The model of the expansion module.
}
type ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePortsProfile struct {
	Enabled *bool  `json:"enabled,omitempty"` // When enabled, override this port's configuration with a port profile.
	ID      string `json:"id,omitempty"`      // When enabled, the ID of the port profile used to override the port's configuration.
	Iname   string `json:"iname,omitempty"`   // When enabled, the IName of the profile.
}
type ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePortsSchedule struct {
	ID   string `json:"id,omitempty"`   // The ID of the port schedule.
	Name string `json:"name,omitempty"` // The name of the port schedule.
}
type ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePortsStackwiseVirtual struct {
	IsDualActiveDetector   *bool `json:"isDualActiveDetector,omitempty"`   // For SVL devices, whether or not the port is used for Dual Active Detection.
	IsStackWiseVirtualLink *bool `json:"isStackWiseVirtualLink,omitempty"` // For SVL devices, whether or not the port is used for StackWise Virtual Link.
}
type ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePort struct {
	AccessPolicyNumber          *int                                                                          `json:"accessPolicyNumber,omitempty"`          // The number of a custom access policy to configure on the switch template port. Only applicable when 'accessPolicyType' is 'Custom access policy'.
	AccessPolicyType            string                                                                        `json:"accessPolicyType,omitempty"`            // The type of the access policy of the switch template port. Only applicable to access ports. Can be one of 'Open', 'Custom access policy', 'MAC allow list' or 'Sticky MAC allow list'.
	AllowedVLANs                string                                                                        `json:"allowedVlans,omitempty"`                // The VLANs allowed on the switch template port. Only applicable to trunk ports.
	DaiTrusted                  *bool                                                                         `json:"daiTrusted,omitempty"`                  // If true, ARP packets for this port will be considered trusted, and Dynamic ARP Inspection will allow the traffic.
	Dot3Az                      *ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortDot3Az           `json:"dot3az,omitempty"`                      // dot3az settings for the port
	Enabled                     *bool                                                                         `json:"enabled,omitempty"`                     // The status of the switch template port.
	FlexibleStackingEnabled     *bool                                                                         `json:"flexibleStackingEnabled,omitempty"`     // For supported switches (e.g. MS420/MS425), whether or not the port has flexible stacking enabled.
	IsolationEnabled            *bool                                                                         `json:"isolationEnabled,omitempty"`            // The isolation status of the switch template port.
	LinkNegotiation             string                                                                        `json:"linkNegotiation,omitempty"`             // The link speed for the switch template port.
	LinkNegotiationCapabilities []string                                                                      `json:"linkNegotiationCapabilities,omitempty"` // Available link speeds for the switch template port.
	MacAllowList                []string                                                                      `json:"macAllowList,omitempty"`                // Only devices with MAC addresses specified in this list will have access to this port. Up to 20 MAC addresses can be defined. Only applicable when 'accessPolicyType' is 'MAC allow list'.
	Mirror                      *ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortMirror           `json:"mirror,omitempty"`                      // Port mirror
	Module                      *ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortModule           `json:"module,omitempty"`                      // Expansion module
	Name                        string                                                                        `json:"name,omitempty"`                        // The name of the switch template port.
	PoeEnabled                  *bool                                                                         `json:"poeEnabled,omitempty"`                  // The PoE status of the switch template port.
	PortID                      string                                                                        `json:"portId,omitempty"`                      // The identifier of the switch template port.
	PortScheduleID              string                                                                        `json:"portScheduleId,omitempty"`              // The ID of the port schedule. A value of null will clear the port schedule.
	Profile                     *ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortProfile          `json:"profile,omitempty"`                     // Profile attributes
	RstpEnabled                 *bool                                                                         `json:"rstpEnabled,omitempty"`                 // The rapid spanning tree protocol status.
	Schedule                    *ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortSchedule         `json:"schedule,omitempty"`                    // The port schedule data.
	StackwiseVirtual            *ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortStackwiseVirtual `json:"stackwiseVirtual,omitempty"`            // Stackwise Virtual settings for the port
	StickyMacAllowList          []string                                                                      `json:"stickyMacAllowList,omitempty"`          // The initial list of MAC addresses for sticky Mac allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
	StickyMacAllowListLimit     *int                                                                          `json:"stickyMacAllowListLimit,omitempty"`     // The maximum number of MAC addresses for sticky MAC allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
	StormControlEnabled         *bool                                                                         `json:"stormControlEnabled,omitempty"`         // The storm control status of the switch template port.
	StpGuard                    string                                                                        `json:"stpGuard,omitempty"`                    // The state of the STP guard ('disabled', 'root guard', 'bpdu guard' or 'loop guard').
	Tags                        []string                                                                      `json:"tags,omitempty"`                        // The list of tags of the switch template port.
	Type                        string                                                                        `json:"type,omitempty"`                        // The type of the switch template port ('trunk', 'access' or 'stack').
	Udld                        string                                                                        `json:"udld,omitempty"`                        // The action to take when Unidirectional Link is detected (Alert only, Enforce). Default configuration is Alert only.
	VLAN                        *int                                                                          `json:"vlan,omitempty"`                        // The VLAN of the switch template port. For a trunk port, this is the native VLAN. A null value will clear the value set for trunk ports.
	VoiceVLAN                   *int                                                                          `json:"voiceVlan,omitempty"`                   // The voice VLAN of the switch template port. Only applicable to access ports.
}
type ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortDot3Az struct {
	Enabled *bool `json:"enabled,omitempty"` // The Energy Efficient Ethernet status of the switch template port.
}
type ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortMirror struct {
	Mode string `json:"mode,omitempty"` // The port mirror mode. Can be one of ('Destination port', 'Source port' or 'Not mirroring traffic').
}
type ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortModule struct {
	Model string `json:"model,omitempty"` // The model of the expansion module.
}
type ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortProfile struct {
	Enabled *bool  `json:"enabled,omitempty"` // When enabled, override this port's configuration with a port profile.
	ID      string `json:"id,omitempty"`      // When enabled, the ID of the port profile used to override the port's configuration.
	Iname   string `json:"iname,omitempty"`   // When enabled, the IName of the profile.
}
type ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortSchedule struct {
	ID   string `json:"id,omitempty"`   // The ID of the port schedule.
	Name string `json:"name,omitempty"` // The name of the port schedule.
}
type ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortStackwiseVirtual struct {
	IsDualActiveDetector   *bool `json:"isDualActiveDetector,omitempty"`   // For SVL devices, whether or not the port is used for Dual Active Detection.
	IsStackWiseVirtualLink *bool `json:"isStackWiseVirtualLink,omitempty"` // For SVL devices, whether or not the port is used for StackWise Virtual Link.
}
type ResponseSwitchUpdateOrganizationConfigTemplateSwitchProfilePort struct {
	AccessPolicyNumber          *int                                                                             `json:"accessPolicyNumber,omitempty"`          // The number of a custom access policy to configure on the switch template port. Only applicable when 'accessPolicyType' is 'Custom access policy'.
	AccessPolicyType            string                                                                           `json:"accessPolicyType,omitempty"`            // The type of the access policy of the switch template port. Only applicable to access ports. Can be one of 'Open', 'Custom access policy', 'MAC allow list' or 'Sticky MAC allow list'.
	AllowedVLANs                string                                                                           `json:"allowedVlans,omitempty"`                // The VLANs allowed on the switch template port. Only applicable to trunk ports.
	DaiTrusted                  *bool                                                                            `json:"daiTrusted,omitempty"`                  // If true, ARP packets for this port will be considered trusted, and Dynamic ARP Inspection will allow the traffic.
	Dot3Az                      *ResponseSwitchUpdateOrganizationConfigTemplateSwitchProfilePortDot3Az           `json:"dot3az,omitempty"`                      // dot3az settings for the port
	Enabled                     *bool                                                                            `json:"enabled,omitempty"`                     // The status of the switch template port.
	FlexibleStackingEnabled     *bool                                                                            `json:"flexibleStackingEnabled,omitempty"`     // For supported switches (e.g. MS420/MS425), whether or not the port has flexible stacking enabled.
	IsolationEnabled            *bool                                                                            `json:"isolationEnabled,omitempty"`            // The isolation status of the switch template port.
	LinkNegotiation             string                                                                           `json:"linkNegotiation,omitempty"`             // The link speed for the switch template port.
	LinkNegotiationCapabilities []string                                                                         `json:"linkNegotiationCapabilities,omitempty"` // Available link speeds for the switch template port.
	MacAllowList                []string                                                                         `json:"macAllowList,omitempty"`                // Only devices with MAC addresses specified in this list will have access to this port. Up to 20 MAC addresses can be defined. Only applicable when 'accessPolicyType' is 'MAC allow list'.
	Mirror                      *ResponseSwitchUpdateOrganizationConfigTemplateSwitchProfilePortMirror           `json:"mirror,omitempty"`                      // Port mirror
	Module                      *ResponseSwitchUpdateOrganizationConfigTemplateSwitchProfilePortModule           `json:"module,omitempty"`                      // Expansion module
	Name                        string                                                                           `json:"name,omitempty"`                        // The name of the switch template port.
	PoeEnabled                  *bool                                                                            `json:"poeEnabled,omitempty"`                  // The PoE status of the switch template port.
	PortID                      string                                                                           `json:"portId,omitempty"`                      // The identifier of the switch template port.
	PortScheduleID              string                                                                           `json:"portScheduleId,omitempty"`              // The ID of the port schedule. A value of null will clear the port schedule.
	Profile                     *ResponseSwitchUpdateOrganizationConfigTemplateSwitchProfilePortProfile          `json:"profile,omitempty"`                     // Profile attributes
	RstpEnabled                 *bool                                                                            `json:"rstpEnabled,omitempty"`                 // The rapid spanning tree protocol status.
	Schedule                    *ResponseSwitchUpdateOrganizationConfigTemplateSwitchProfilePortSchedule         `json:"schedule,omitempty"`                    // The port schedule data.
	StackwiseVirtual            *ResponseSwitchUpdateOrganizationConfigTemplateSwitchProfilePortStackwiseVirtual `json:"stackwiseVirtual,omitempty"`            // Stackwise Virtual settings for the port
	StickyMacAllowList          []string                                                                         `json:"stickyMacAllowList,omitempty"`          // The initial list of MAC addresses for sticky Mac allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
	StickyMacAllowListLimit     *int                                                                             `json:"stickyMacAllowListLimit,omitempty"`     // The maximum number of MAC addresses for sticky MAC allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
	StormControlEnabled         *bool                                                                            `json:"stormControlEnabled,omitempty"`         // The storm control status of the switch template port.
	StpGuard                    string                                                                           `json:"stpGuard,omitempty"`                    // The state of the STP guard ('disabled', 'root guard', 'bpdu guard' or 'loop guard').
	Tags                        []string                                                                         `json:"tags,omitempty"`                        // The list of tags of the switch template port.
	Type                        string                                                                           `json:"type,omitempty"`                        // The type of the switch template port ('trunk', 'access' or 'stack').
	Udld                        string                                                                           `json:"udld,omitempty"`                        // The action to take when Unidirectional Link is detected (Alert only, Enforce). Default configuration is Alert only.
	VLAN                        *int                                                                             `json:"vlan,omitempty"`                        // The VLAN of the switch template port. For a trunk port, this is the native VLAN. A null value will clear the value set for trunk ports.
	VoiceVLAN                   *int                                                                             `json:"voiceVlan,omitempty"`                   // The voice VLAN of the switch template port. Only applicable to access ports.
}
type ResponseSwitchUpdateOrganizationConfigTemplateSwitchProfilePortDot3Az struct {
	Enabled *bool `json:"enabled,omitempty"` // The Energy Efficient Ethernet status of the switch template port.
}
type ResponseSwitchUpdateOrganizationConfigTemplateSwitchProfilePortMirror struct {
	Mode string `json:"mode,omitempty"` // The port mirror mode. Can be one of ('Destination port', 'Source port' or 'Not mirroring traffic').
}
type ResponseSwitchUpdateOrganizationConfigTemplateSwitchProfilePortModule struct {
	Model string `json:"model,omitempty"` // The model of the expansion module.
}
type ResponseSwitchUpdateOrganizationConfigTemplateSwitchProfilePortProfile struct {
	Enabled *bool  `json:"enabled,omitempty"` // When enabled, override this port's configuration with a port profile.
	ID      string `json:"id,omitempty"`      // When enabled, the ID of the port profile used to override the port's configuration.
	Iname   string `json:"iname,omitempty"`   // When enabled, the IName of the profile.
}
type ResponseSwitchUpdateOrganizationConfigTemplateSwitchProfilePortSchedule struct {
	ID   string `json:"id,omitempty"`   // The ID of the port schedule.
	Name string `json:"name,omitempty"` // The name of the port schedule.
}
type ResponseSwitchUpdateOrganizationConfigTemplateSwitchProfilePortStackwiseVirtual struct {
	IsDualActiveDetector   *bool `json:"isDualActiveDetector,omitempty"`   // For SVL devices, whether or not the port is used for Dual Active Detection.
	IsStackWiseVirtualLink *bool `json:"isStackWiseVirtualLink,omitempty"` // For SVL devices, whether or not the port is used for StackWise Virtual Link.
}
type ResponseSwitchGetOrganizationSummarySwitchPowerHistory []ResponseItemSwitchGetOrganizationSummarySwitchPowerHistory // Array of ResponseSwitchGetOrganizationSummarySwitchPowerHistory
type ResponseItemSwitchGetOrganizationSummarySwitchPowerHistory struct {
	Draw *float64 `json:"draw,omitempty"` // The PoE power draw in watts for all switch ports in the organization for the given interval.
	Ts   string   `json:"ts,omitempty"`   // Timestamp of the start of the interval.
}
type ResponseSwitchCloneOrganizationSwitchDevices struct {
	SourceSerial  string   `json:"sourceSerial,omitempty"`  // Serial number of the source switch (must be on a network not bound to a template)
	TargetSerials []string `json:"targetSerials,omitempty"` // Array of serial numbers of one or more target switches (must be on a network not bound to a template)
}
type ResponseSwitchGetOrganizationSwitchPortsBySwitch struct {
	Mac     string                                                   `json:"mac,omitempty"`     // The MAC address of the switch.
	Model   string                                                   `json:"model,omitempty"`   // The model of the switch.
	Name    string                                                   `json:"name,omitempty"`    // The name of the switch.
	Network *ResponseSwitchGetOrganizationSwitchPortsBySwitchNetwork `json:"network,omitempty"` // Identifying information of the switch's network.
	Ports   *[]ResponseSwitchGetOrganizationSwitchPortsBySwitchPorts `json:"ports,omitempty"`   // Ports belonging to the switch
	Serial  string                                                   `json:"serial,omitempty"`  // The serial number of the switch.
}
type ResponseSwitchGetOrganizationSwitchPortsBySwitchNetwork struct {
	ID   string `json:"id,omitempty"`   // The ID of the network.
	Name string `json:"name,omitempty"` // The name of the network.
}
type ResponseSwitchGetOrganizationSwitchPortsBySwitchPorts struct {
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
	Type                    string   `json:"type,omitempty"`                    // The type of the switch port ('trunk', 'access' or 'stack').
	VLAN                    *int     `json:"vlan,omitempty"`                    // The VLAN of the switch port. For a trunk port, this is the native VLAN. A null value will clear the value set for trunk ports.
	VoiceVLAN               *int     `json:"voiceVlan,omitempty"`               // The voice VLAN of the switch port. Only applicable to access ports.
}
type ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDevice struct {
	Items *[]ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceItems `json:"items,omitempty"` // Switches
	Meta  *ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceMeta    `json:"meta,omitempty"`  // Metadata relevant to the paginated dataset
}
type ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceItems struct {
	Mac     string                                                                       `json:"mac,omitempty"`     // The MAC address of the switch.
	Model   string                                                                       `json:"model,omitempty"`   // The model of the switch.
	Name    string                                                                       `json:"name,omitempty"`    // The name of the switch.
	Network *ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceItemsNetwork `json:"network,omitempty"` // Identifying information of the switch's network.
	Ports   *[]ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceItemsPorts `json:"ports,omitempty"`   // The number of online clients of the ports on the switch.
	Serial  string                                                                       `json:"serial,omitempty"`  // The serial number of the switch.
}
type ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceItemsNetwork struct {
	ID   string `json:"id,omitempty"`   // The ID of the network.
	Name string `json:"name,omitempty"` // The name of the network.
}
type ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceItemsPorts struct {
	Counts *ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceItemsPortsCounts `json:"counts,omitempty"` // Number of clients on the port in a given time.
	PortID string                                                                           `json:"portId,omitempty"` // The string identifier of this port on the switch. This is commonly just the port number but may contain additional identifying information such as the slot and module-type if the port is located on a port module.
}
type ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceItemsPortsCounts struct {
	ByStatus *ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceItemsPortsCountsByStatus `json:"byStatus,omitempty"` // Associated client count on access point by status.
}
type ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceItemsPortsCountsByStatus struct {
	Online *int `json:"online,omitempty"` // Active client count.
}
type ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceMeta struct {
	Counts *ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceMetaCounts `json:"counts,omitempty"` // Counts relating to the paginated dataset
}
type ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceMetaCounts struct {
	Items *ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceMetaCountsItems `json:"items,omitempty"` // Counts relating to the paginated items
}
type ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceMetaCountsItems struct {
	Remaining *int `json:"remaining,omitempty"` // The number of items in the dataset that are available on subsequent pages
	Total     *int `json:"total,omitempty"`     // The total number of items in the dataset
}
type ResponseSwitchGetOrganizationSwitchPortsOverview struct {
	Counts *ResponseSwitchGetOrganizationSwitchPortsOverviewCounts `json:"counts,omitempty"` // The count data of all ports
}
type ResponseSwitchGetOrganizationSwitchPortsOverviewCounts struct {
	ByStatus *ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatus `json:"byStatus,omitempty"` // The count data, indexed by active or inactive status
	Total    *int                                                            `json:"total,omitempty"`    // The total number of ports
}
type ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatus struct {
	Active   *ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusActive   `json:"active,omitempty"`   // The count data for active ports
	Inactive *ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusInactive `json:"inactive,omitempty"` // The count data for inactive ports
}
type ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusActive struct {
	ByMediaAndLinkSpeed *ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusActiveByMediaAndLinkSpeed `json:"byMediaAndLinkSpeed,omitempty"` // The active count data, indexed by media type (RJ45 or SFP)
	Total               *int                                                                                     `json:"total,omitempty"`               // The total number of active ports
}
type ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusActiveByMediaAndLinkSpeed struct {
	Rj45 *ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusActiveByMediaAndLinkSpeedRj45 `json:"rj45,omitempty"` // The count data for RJ45 ports, indexed by speed in Mb
	Sfp  *ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusActiveByMediaAndLinkSpeedSfp  `json:"sfp,omitempty"`  // The count data for SFP ports, indexed by speed in Mb
}
type ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusActiveByMediaAndLinkSpeedRj45 struct {
	Status10    *int `json:"10,omitempty"`    // The number of active 10 Mbps RJ45 ports
	Status100   *int `json:"100,omitempty"`   // The number of active 100 Mbps RJ45 ports
	Status1000  *int `json:"1000,omitempty"`  // The number of active 1 Gbps RJ45 ports
	Status10000 *int `json:"10000,omitempty"` // The number of active 10 Gbps RJ45 ports
	Status2500  *int `json:"2500,omitempty"`  // The number of active 2 Gbps RJ45 ports
	Status5000  *int `json:"5000,omitempty"`  // The number of active 5 Gbps RJ45 ports
	Total       *int `json:"total,omitempty"` // The total number of active RJ45 ports
}
type ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusActiveByMediaAndLinkSpeedSfp struct {
	Status100    *int `json:"100,omitempty"`    // The number of active 100 Mbps SFP ports
	Status1000   *int `json:"1000,omitempty"`   // The number of active 1 Gbps SFP ports
	Status10000  *int `json:"10000,omitempty"`  // The number of active 10 Gbps SFP ports
	Status100000 *int `json:"100000,omitempty"` // The number of active 100 Gbps SFP ports
	Status20000  *int `json:"20000,omitempty"`  // The number of active 20 Gbps SFP ports
	Status25000  *int `json:"25000,omitempty"`  // The number of active 25 Gbps SFP ports
	Status40000  *int `json:"40000,omitempty"`  // The number of active 40 Gbps SFP ports
	Status50000  *int `json:"50000,omitempty"`  // The number of active 50 Gbps SFP ports
	Total        *int `json:"total,omitempty"`  // The total number of active SFP ports
}
type ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusInactive struct {
	ByMedia *ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusInactiveByMedia `json:"byMedia,omitempty"` // The inactive count data, indexed by media type (RJ45 or SFP)
	Total   *int                                                                           `json:"total,omitempty"`   // The total number of inactive ports
}
type ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusInactiveByMedia struct {
	Rj45 *ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusInactiveByMediaRj45 `json:"rj45,omitempty"` // The count data for inactive RJ45 ports
	Sfp  *ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusInactiveByMediaSfp  `json:"sfp,omitempty"`  // The count data for inactive SFP ports
}
type ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusInactiveByMediaRj45 struct {
	Total *int `json:"total,omitempty"` // The total number of inactive RJ45 ports
}
type ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusInactiveByMediaSfp struct {
	Total *int `json:"total,omitempty"` // The total number of inactive SFP ports
}
type ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitch struct {
	Items *[]ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchItems `json:"items,omitempty"` // Switches
	Meta  *ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchMeta    `json:"meta,omitempty"`  // Metadata relevant to the paginated dataset
}
type ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchItems struct {
	Mac     string                                                                `json:"mac,omitempty"`     // The MAC address of the switch.
	Model   string                                                                `json:"model,omitempty"`   // The model of the switch.
	Name    string                                                                `json:"name,omitempty"`    // The name of the switch.
	Network *ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchItemsNetwork `json:"network,omitempty"` // Identifying information of the switch's network.
	Ports   *[]ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchItemsPorts `json:"ports,omitempty"`   // The statuses of the ports on the switch.
	Serial  string                                                                `json:"serial,omitempty"`  // The serial number of the switch.
}
type ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchItemsNetwork struct {
	ID   string `json:"id,omitempty"`   // The ID of the network.
	Name string `json:"name,omitempty"` // The name of the network.
}
type ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchItemsPorts struct {
	Duplex       string                                                                          `json:"duplex,omitempty"`       // The current duplex of a connected port.
	Enabled      *bool                                                                           `json:"enabled,omitempty"`      // Whether the port is configured to be enabled.
	Errors       []string                                                                        `json:"errors,omitempty"`       // All errors present on the port.
	IsUplink     *bool                                                                           `json:"isUplink,omitempty"`     // Whether the port is the switch's uplink.
	Poe          *ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchItemsPortsPoe          `json:"poe,omitempty"`          // PoE status of the port.
	PortID       string                                                                          `json:"portId,omitempty"`       // The string identifier of this port on the switch. This is commonly just the port number but may contain additional identifying information such as the slot and module-type if the port is located on a port module.
	SecurePort   *ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchItemsPortsSecurePort   `json:"securePort,omitempty"`   // The Secure Port status of the port.
	SpanningTree *ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchItemsPortsSpanningTree `json:"spanningTree,omitempty"` // The Spanning Tree Protocol (STP) information of the connected device.
	Speed        string                                                                          `json:"speed,omitempty"`        // The current data transfer rate which the port is operating at.
	Status       string                                                                          `json:"status,omitempty"`       // The current connection status of the port.
	Warnings     []string                                                                        `json:"warnings,omitempty"`     // All warnings present on the port.
}
type ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchItemsPortsPoe struct {
	IsAllocated *bool `json:"isAllocated,omitempty"` // Whether the port is drawing power
}
type ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchItemsPortsSecurePort struct {
	Active               *bool  `json:"active,omitempty"`               // Whether Secure Port is currently active for this port.
	AuthenticationStatus string `json:"authenticationStatus,omitempty"` // The current Secure Port status.
}
type ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchItemsPortsSpanningTree struct {
	Statuses []string `json:"statuses,omitempty"` // The current Spanning Tree Protocol statuses of the port.
}
type ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchMeta struct {
	Counts *ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchMetaCounts `json:"counts,omitempty"` // Counts relating to the paginated dataset
}
type ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchMetaCounts struct {
	Items *ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchMetaCountsItems `json:"items,omitempty"` // Counts relating to the paginated items
}
type ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchMetaCountsItems struct {
	Remaining *int `json:"remaining,omitempty"` // The number of items in the dataset that are available on subsequent pages
	Total     *int `json:"total,omitempty"`     // The total number of items in the dataset
}
type ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDevice struct {
	Items *[]ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceItems `json:"items,omitempty"` // Switches
	Meta  *ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceMeta    `json:"meta,omitempty"`  // Metadata relevant to the paginated dataset
}
type ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceItems struct {
	Mac     string                                                                         `json:"mac,omitempty"`     // The MAC address of the switch.
	Model   string                                                                         `json:"model,omitempty"`   // The model of the switch.
	Name    string                                                                         `json:"name,omitempty"`    // The name of the switch.
	Network *ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceItemsNetwork `json:"network,omitempty"` // Identifying information of the switch's network.
	Ports   *[]ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceItemsPorts `json:"ports,omitempty"`   // Ports belonging to the switch with LLDP/CDP discovery info.
	Serial  string                                                                         `json:"serial,omitempty"`  // The serial number of the switch.
}
type ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceItemsNetwork struct {
	ID   string `json:"id,omitempty"`   // The ID of the network.
	Name string `json:"name,omitempty"` // The name of the network.
}
type ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceItemsPorts struct {
	Cdp           *[]ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceItemsPortsCdp  `json:"cdp,omitempty"`           // The Cisco Discovery Protocol (CDP) information of the connected device.
	LastUpdatedAt string                                                                             `json:"lastUpdatedAt,omitempty"` // Timestamp for most recent discovery info on this port.
	Lldp          *[]ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceItemsPortsLldp `json:"lldp,omitempty"`          // The Link Layer Discovery Protocol (LLDP) information of the connected device.
	PortID        string                                                                             `json:"portId,omitempty"`        // The string identifier of this port on the switch. This is commonly just the port number but may contain additional identifying information such as the slot and module-type if the port is located on a port module.
}
type ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceItemsPortsCdp struct {
	Name  string `json:"name,omitempty"`  // CDP RFC/official name of TLV
	Value string `json:"value,omitempty"` // Value of the named TLV.
}
type ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceItemsPortsLldp struct {
	Name  string `json:"name,omitempty"`  // LLDP RFC/official name of TLV
	Value string `json:"value,omitempty"` // Value of the named TLV.
}
type ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceMeta struct {
	Counts *ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceMetaCounts `json:"counts,omitempty"` // Counts relating to the paginated dataset
}
type ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceMetaCounts struct {
	Items *ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceMetaCountsItems `json:"items,omitempty"` // Counts relating to the paginated items
}
type ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceMetaCountsItems struct {
	Remaining *int `json:"remaining,omitempty"` // The number of items in the dataset that are available on subsequent pages
	Total     *int `json:"total,omitempty"`     // The total number of items in the dataset
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
	Dot3Az                  *RequestSwitchUpdateDeviceSwitchPortDot3Az  `json:"dot3az,omitempty"`                  // dot3az settings for the port
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
	Type                    string                                      `json:"type,omitempty"`                    // The type of the switch port ('trunk', 'access' or 'stack').
	Udld                    string                                      `json:"udld,omitempty"`                    // The action to take when Unidirectional Link is detected (Alert only, Enforce). Default configuration is Alert only.
	VLAN                    *int                                        `json:"vlan,omitempty"`                    // The VLAN of the switch port. For a trunk port, this is the native VLAN. A null value will clear the value set for trunk ports.
	VoiceVLAN               *int                                        `json:"voiceVlan,omitempty"`               // The voice VLAN of the switch port. Only applicable to access ports.
}
type RequestSwitchUpdateDeviceSwitchPortDot3Az struct {
	Enabled *bool `json:"enabled,omitempty"` // The Energy Efficient Ethernet status of the switch port.
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
type RequestSwitchUpdateDeviceSwitchRoutingInterface struct {
	DefaultGateway   string                                                       `json:"defaultGateway,omitempty"`   // The next hop for any traffic that isn't going to a directly connected subnet or over a static route.         This IP address must exist in a subnet with a routed interface. Required if this is the first IPv4 interface.
	InterfaceIP      string                                                       `json:"interfaceIp,omitempty"`      // The IP address this switch will use for layer 3 routing on this VLAN or subnet. This cannot be the same         as the switch's management IP.
	IPv6             *RequestSwitchUpdateDeviceSwitchRoutingInterfaceIPv6         `json:"ipv6,omitempty"`             // The IPv6 settings of the interface.
	MulticastRouting string                                                       `json:"multicastRouting,omitempty"` // Enable multicast support if, multicast routing between VLANs is required. Options are:         'disabled', 'enabled' or 'IGMP snooping querier'. Default is 'disabled'.
	Name             string                                                       `json:"name,omitempty"`             // A friendly name or description for the interface or VLAN.
	OspfSettings     *RequestSwitchUpdateDeviceSwitchRoutingInterfaceOspfSettings `json:"ospfSettings,omitempty"`     // The OSPF routing settings of the interface.
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
type RequestSwitchUpdateDeviceSwitchRoutingInterfaceDhcp struct {
	BootFileName         string                                                                   `json:"bootFileName,omitempty"`         // The PXE boot server filename for the DHCP server running on the switch interface
	BootNextServer       string                                                                   `json:"bootNextServer,omitempty"`       // The PXE boot server IP for the DHCP server running on the switch interface
	BootOptionsEnabled   *bool                                                                    `json:"bootOptionsEnabled,omitempty"`   // Enable DHCP boot options to provide PXE boot options configs for the dhcp server running on the switch         interface
	DhcpLeaseTime        string                                                                   `json:"dhcpLeaseTime,omitempty"`        // The DHCP lease time config for the dhcp server running on switch interface         ('30 minutes', '1 hour', '4 hours', '12 hours', '1 day' or '1 week')
	DhcpMode             string                                                                   `json:"dhcpMode,omitempty"`             // The DHCP mode options for the switch interface        ('dhcpDisabled', 'dhcpRelay' or 'dhcpServer')
	DhcpOptions          *[]RequestSwitchUpdateDeviceSwitchRoutingInterfaceDhcpDhcpOptions        `json:"dhcpOptions,omitempty"`          // Array of DHCP options consisting of code, type and value for the DHCP server running on the switch interface
	DhcpRelayServerIPs   []string                                                                 `json:"dhcpRelayServerIps,omitempty"`   // The DHCP relay server IPs to which DHCP packets would get relayed for the switch interface
	DNSCustomNameservers []string                                                                 `json:"dnsCustomNameservers,omitempty"` // The DHCP name server IPs when DHCP name server option is         'custom'
	DNSNameserversOption string                                                                   `json:"dnsNameserversOption,omitempty"` // The DHCP name server option for the dhcp server running on the switch interface         ('googlePublicDns', 'openDns' or 'custom')
	FixedIPAssignments   *[]RequestSwitchUpdateDeviceSwitchRoutingInterfaceDhcpFixedIPAssignments `json:"fixedIpAssignments,omitempty"`   // Array of DHCP fixed IP assignments for the DHCP server running on the switch interface
	ReservedIPRanges     *[]RequestSwitchUpdateDeviceSwitchRoutingInterfaceDhcpReservedIPRanges   `json:"reservedIpRanges,omitempty"`     // Array of DHCP reserved IP assignments for the DHCP server running on the switch interface
}
type RequestSwitchUpdateDeviceSwitchRoutingInterfaceDhcpDhcpOptions struct {
	Code  string `json:"code,omitempty"`  // The code for DHCP option which should be from 2 to 254
	Type  string `json:"type,omitempty"`  // The type of the DHCP option which should be one of           ('text', 'ip', 'integer' or 'hex')
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
	ManagementNextHop           string `json:"managementNextHop,omitempty"`           // Optional fallback IP address for management traffic
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
	RadiusGroupAttribute           string                                                                 `json:"radiusGroupAttribute,omitempty"`           // Acceptable values are `""` for , or `"11"` for Group Policies ACL
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
	Cache                    *RequestSwitchCreateNetworkSwitchAccessPolicyRadiusCache        `json:"cache,omitempty"`                    // Object for RADIUS Cache Settings
	CriticalAuth             *RequestSwitchCreateNetworkSwitchAccessPolicyRadiusCriticalAuth `json:"criticalAuth,omitempty"`             // Critical auth settings for when authentication is rejected by the RADIUS server
	FailedAuthVLANID         *int                                                            `json:"failedAuthVlanId,omitempty"`         // VLAN that clients will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
	ReAuthenticationInterval *int                                                            `json:"reAuthenticationInterval,omitempty"` // Re-authentication period in seconds. Will be null if hostMode is Multi-Auth
}
type RequestSwitchCreateNetworkSwitchAccessPolicyRadiusCache struct {
	Enabled *bool `json:"enabled,omitempty"` // Enable to cache authorization and authentication responses on the RADIUS server
	Timeout *int  `json:"timeout,omitempty"` // If RADIUS caching is enabled, this value dictates how long the cache will remain in the RADIUS server, in hours, to allow network access without authentication
}
type RequestSwitchCreateNetworkSwitchAccessPolicyRadiusCriticalAuth struct {
	DataVLANID        *int  `json:"dataVlanId,omitempty"`        // VLAN that clients who use data will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
	SuspendPortBounce *bool `json:"suspendPortBounce,omitempty"` // Enable to suspend port bounce when RADIUS servers are unreachable
	VoiceVLANID       *int  `json:"voiceVlanId,omitempty"`       // VLAN that clients who use voice will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
}
type RequestSwitchCreateNetworkSwitchAccessPolicyRadiusAccountingServers struct {
	Host                       string `json:"host,omitempty"`                       // Public IP address of the RADIUS accounting server
	OrganizationRadiusServerID string `json:"organizationRadiusServerId,omitempty"` // Organization wide RADIUS server ID. If this field is provided, the host, port and secret field will be ignored
	Port                       *int   `json:"port,omitempty"`                       // UDP port that the RADIUS Accounting server listens on for access requests
	Secret                     string `json:"secret,omitempty"`                     // RADIUS client shared secret
}
type RequestSwitchCreateNetworkSwitchAccessPolicyRadiusServers struct {
	Host                       string `json:"host,omitempty"`                       // Public IP address of the RADIUS server
	OrganizationRadiusServerID string `json:"organizationRadiusServerId,omitempty"` // Organization wide RADIUS server ID. If this field is provided, the host, port and secret field will be ignored
	Port                       *int   `json:"port,omitempty"`                       // UDP port that the RADIUS server listens on for access requests
	Secret                     string `json:"secret,omitempty"`                     // RADIUS client shared secret
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
	RadiusGroupAttribute           string                                                                 `json:"radiusGroupAttribute,omitempty"`           // Acceptable values are `""` for , or `"11"` for Group Policies ACL
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
	Cache                    *RequestSwitchUpdateNetworkSwitchAccessPolicyRadiusCache        `json:"cache,omitempty"`                    // Object for RADIUS Cache Settings
	CriticalAuth             *RequestSwitchUpdateNetworkSwitchAccessPolicyRadiusCriticalAuth `json:"criticalAuth,omitempty"`             // Critical auth settings for when authentication is rejected by the RADIUS server
	FailedAuthVLANID         *int                                                            `json:"failedAuthVlanId,omitempty"`         // VLAN that clients will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
	ReAuthenticationInterval *int                                                            `json:"reAuthenticationInterval,omitempty"` // Re-authentication period in seconds. Will be null if hostMode is Multi-Auth
}
type RequestSwitchUpdateNetworkSwitchAccessPolicyRadiusCache struct {
	Enabled *bool `json:"enabled,omitempty"` // Enable to cache authorization and authentication responses on the RADIUS server
	Timeout *int  `json:"timeout,omitempty"` // If RADIUS caching is enabled, this value dictates how long the cache will remain in the RADIUS server, in hours, to allow network access without authentication
}
type RequestSwitchUpdateNetworkSwitchAccessPolicyRadiusCriticalAuth struct {
	DataVLANID        *int  `json:"dataVlanId,omitempty"`        // VLAN that clients who use data will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
	SuspendPortBounce *bool `json:"suspendPortBounce,omitempty"` // Enable to suspend port bounce when RADIUS servers are unreachable
	VoiceVLANID       *int  `json:"voiceVlanId,omitempty"`       // VLAN that clients who use voice will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
}
type RequestSwitchUpdateNetworkSwitchAccessPolicyRadiusAccountingServers struct {
	Host                       string `json:"host,omitempty"`                       // Public IP address of the RADIUS accounting server
	OrganizationRadiusServerID string `json:"organizationRadiusServerId,omitempty"` // Organization wide RADIUS server ID. If this field is provided, the host, port and secret field will be ignored
	Port                       *int   `json:"port,omitempty"`                       // UDP port that the RADIUS Accounting server listens on for access requests
	Secret                     string `json:"secret,omitempty"`                     // RADIUS client shared secret
	ServerID                   string `json:"serverId,omitempty"`                   // Unique ID of the RADIUS accounting server. When provided, the existing RADIUS server will be updated instead of creating a new one
}
type RequestSwitchUpdateNetworkSwitchAccessPolicyRadiusServers struct {
	Host                       string `json:"host,omitempty"`                       // Public IP address of the RADIUS server
	OrganizationRadiusServerID string `json:"organizationRadiusServerId,omitempty"` // Organization wide RADIUS server ID. If this field is provided, the host, port and secret field will be ignored
	Port                       *int   `json:"port,omitempty"`                       // UDP port that the RADIUS server listens on for access requests
	Secret                     string `json:"secret,omitempty"`                     // RADIUS client shared secret
	ServerID                   string `json:"serverId,omitempty"`                   // Unique ID of the RADIUS server. When provided, the existing RADIUS server will be updated instead of creating a new one
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
	Overrides      *[]RequestSwitchUpdateNetworkSwitchMtuOverrides `json:"overrides,omitempty"`      // Override MTU size for individual switches or switch templates. An empty array will clear overrides.
}
type RequestSwitchUpdateNetworkSwitchMtuOverrides struct {
	MtuSize        *int     `json:"mtuSize,omitempty"`        // MTU size for the switches or switch templates.
	SwitchProfiles []string `json:"switchProfiles,omitempty"` // List of switch template IDs. Applicable only for template network.
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
	Dscp         *int   `json:"dscp,omitempty"`         // DSCP tag for the incoming packet. Set this to -1 to trust incoming DSCP. Default value is 0
	DstPort      *int   `json:"dstPort,omitempty"`      // The destination port of the incoming packet. Applicable only if protocol is TCP or UDP.
	DstPortRange string `json:"dstPortRange,omitempty"` // The destination port range of the incoming packet. Applicable only if protocol is set to TCP or UDP.
	Protocol     string `json:"protocol,omitempty"`     // The protocol of the incoming packet. Default value is "ANY"
	SrcPort      *int   `json:"srcPort,omitempty"`      // The source port of the incoming packet. Applicable only if protocol is TCP or UDP.
	SrcPortRange string `json:"srcPortRange,omitempty"` // The source port range of the incoming packet. Applicable only if protocol is set to TCP or UDP.
	VLAN         *int   `json:"vlan,omitempty"`         // The VLAN of the incoming packet. A null value will match any VLAN.
}
type RequestSwitchUpdateNetworkSwitchQosRulesOrder struct {
	RuleIDs []string `json:"ruleIds,omitempty"` // A list of quality of service rule IDs arranged in order in which they should be processed by the switch.
}
type RequestSwitchUpdateNetworkSwitchQosRule struct {
	Dscp         *int   `json:"dscp,omitempty"`         // DSCP tag that should be assigned to incoming packet. Set this to -1 to trust incoming DSCP. Default value is 0
	DstPort      *int   `json:"dstPort,omitempty"`      // The destination port of the incoming packet. Applicable only if protocol is TCP or UDP.
	DstPortRange string `json:"dstPortRange,omitempty"` // The destination port range of the incoming packet. Applicable only if protocol is set to TCP or UDP.
	Protocol     string `json:"protocol,omitempty"`     // The protocol of the incoming packet. Default value is "ANY"
	SrcPort      *int   `json:"srcPort,omitempty"`      // The source port of the incoming packet. Applicable only if protocol is TCP or UDP.
	SrcPortRange string `json:"srcPortRange,omitempty"` // The source port range of the incoming packet. Applicable only if protocol is set to TCP or UDP.
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
	FloodUnknownMulticastTrafficEnabled *bool    `json:"floodUnknownMulticastTrafficEnabled,omitempty"` // Flood unknown multicast traffic setting for switches, switch stacks or switch templates
	IgmpSnoopingEnabled                 *bool    `json:"igmpSnoopingEnabled,omitempty"`                 // IGMP snooping setting for switches, switch stacks or switch templates
	Stacks                              []string `json:"stacks,omitempty"`                              // List of switch stack ids for non-template network
	SwitchProfiles                      []string `json:"switchProfiles,omitempty"`                      // List of switch templates ids for template network
	Switches                            []string `json:"switches,omitempty"`                            // List of switch serials for non-template network
}
type RequestSwitchCreateNetworkSwitchRoutingMulticastRendezvousPoint struct {
	InterfaceIP    string `json:"interfaceIp,omitempty"`    // The IP address of the interface where the RP needs to be created.
	MulticastGroup string `json:"multicastGroup,omitempty"` // 'Any', or the IP address of a multicast group
}
type RequestSwitchUpdateNetworkSwitchRoutingMulticastRendezvousPoint struct {
	InterfaceIP    string `json:"interfaceIp,omitempty"`    // The IP address of the interface where the RP needs to be created.
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
	MacBlocklist         *RequestSwitchUpdateNetworkSwitchSettingsMacBlocklist         `json:"macBlocklist,omitempty"`         // MAC blocklist
	PowerExceptions      *[]RequestSwitchUpdateNetworkSwitchSettingsPowerExceptions    `json:"powerExceptions,omitempty"`      // Exceptions on a per switch basis to "useCombinedPower"
	UplinkClientSampling *RequestSwitchUpdateNetworkSwitchSettingsUplinkClientSampling `json:"uplinkClientSampling,omitempty"` // Uplink client sampling
	UseCombinedPower     *bool                                                         `json:"useCombinedPower,omitempty"`     // The use Combined Power as the default behavior of secondary power supplies on supported devices.
	VLAN                 *int                                                          `json:"vlan,omitempty"`                 // Management VLAN
}
type RequestSwitchUpdateNetworkSwitchSettingsMacBlocklist struct {
	Enabled *bool `json:"enabled,omitempty"` // Enable MAC blocklist
}
type RequestSwitchUpdateNetworkSwitchSettingsPowerExceptions struct {
	PowerType string `json:"powerType,omitempty"` // Per switch exception (combined, redundant, useNetworkSetting)
	Serial    string `json:"serial,omitempty"`    // Serial number of the switch
}
type RequestSwitchUpdateNetworkSwitchSettingsUplinkClientSampling struct {
	Enabled *bool `json:"enabled,omitempty"` // Enable uplink client sampling
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
	BootOptionsEnabled   *bool                                                                          `json:"bootOptionsEnabled,omitempty"`   // Enable DHCP boot options to provide PXE boot options configs for the dhcp server running on the switch         stack interface
	DhcpLeaseTime        string                                                                         `json:"dhcpLeaseTime,omitempty"`        // The DHCP lease time config for the dhcp server running on switch stack interface         ('30 minutes', '1 hour', '4 hours', '12 hours', '1 day' or '1 week')
	DhcpMode             string                                                                         `json:"dhcpMode,omitempty"`             // The DHCP mode options for the switch stack interface         ('dhcpDisabled', 'dhcpRelay' or 'dhcpServer')
	DhcpOptions          *[]RequestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcpDhcpOptions        `json:"dhcpOptions,omitempty"`          // Array of DHCP options consisting of code, type and value for the DHCP server running on the         switch stack interface
	DhcpRelayServerIPs   []string                                                                       `json:"dhcpRelayServerIps,omitempty"`   // The DHCP relay server IPs to which DHCP packets would get relayed for the switch stack interface
	DNSCustomNameservers []string                                                                       `json:"dnsCustomNameservers,omitempty"` // The DHCP name server IPs when DHCP name server option is '         custom'
	DNSNameserversOption string                                                                         `json:"dnsNameserversOption,omitempty"` // The DHCP name server option for the dhcp server running on the switch stack interface         ('googlePublicDns', 'openDns' or 'custom')
	FixedIPAssignments   *[]RequestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcpFixedIPAssignments `json:"fixedIpAssignments,omitempty"`   // Array of DHCP fixed IP assignments for the DHCP server running on the switch stack interface
	ReservedIPRanges     *[]RequestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcpReservedIPRanges   `json:"reservedIpRanges,omitempty"`     // Array of DHCP reserved IP assignments for the DHCP server running on the switch stack interface
}
type RequestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcpDhcpOptions struct {
	Code  string `json:"code,omitempty"`  // The code for DHCP option which should be from 2 to 254
	Type  string `json:"type,omitempty"`  // The type of the DHCP option which should be one of           ('text', 'ip', 'integer' or 'hex')
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
	ManagementNextHop           string `json:"managementNextHop,omitempty"`           // Optional fallback IP address for management traffic
	Name                        string `json:"name,omitempty"`                        // Name or description for layer 3 static route
	NextHopIP                   string `json:"nextHopIp,omitempty"`                   // IP address of the next hop device to which the device sends its traffic for the subnet
	PreferOverOspfRoutesEnabled *bool  `json:"preferOverOspfRoutesEnabled,omitempty"` // Option to prefer static route over OSPF routes
	Subnet                      string `json:"subnet,omitempty"`                      // The subnet which is routed via this static route and should be specified in CIDR notation (ex. 1.2.3.0/24)
}
type RequestSwitchUpdateNetworkSwitchStormControl struct {
	BroadcastThreshold                   *int     `json:"broadcastThreshold,omitempty"`                   // Percentage (1 to 99) of total available port bandwidth for broadcast traffic type. Default value 100 percent rate is to clear the configuration.
	MulticastThreshold                   *int     `json:"multicastThreshold,omitempty"`                   // Percentage (1 to 99) of total available port bandwidth for multicast traffic type. Default value 100 percent rate is to clear the configuration.
	TreatTheseTrafficTypesAsOneThreshold []string `json:"treatTheseTrafficTypesAsOneThreshold,omitempty"` // Grouped traffic types
	UnknownUnicastThreshold              *int     `json:"unknownUnicastThreshold,omitempty"`              // Percentage (1 to 99) of total available port bandwidth for unknown unicast (dlf-destination lookup failure) traffic type. Default value 100 percent rate is to clear the configuration.
}
type RequestSwitchUpdateNetworkSwitchStp struct {
	RstpEnabled       *bool                                                   `json:"rstpEnabled,omitempty"`       // The spanning tree protocol status in network
	StpBridgePriority *[]RequestSwitchUpdateNetworkSwitchStpStpBridgePriority `json:"stpBridgePriority,omitempty"` // STP bridge priority for switches/stacks or switch templates. An empty array will clear the STP bridge priority settings.
}
type RequestSwitchUpdateNetworkSwitchStpStpBridgePriority struct {
	Stacks         []string `json:"stacks,omitempty"`         // List of stack IDs
	StpPriority    *int     `json:"stpPriority,omitempty"`    // STP priority for switch, stacks, or switch templates
	SwitchProfiles []string `json:"switchProfiles,omitempty"` // List of switch template IDs
	Switches       []string `json:"switches,omitempty"`       // List of switch serial numbers
}
type RequestSwitchUpdateOrganizationConfigTemplateSwitchProfilePort struct {
	AccessPolicyNumber      *int                                                                   `json:"accessPolicyNumber,omitempty"`      // The number of a custom access policy to configure on the switch template port. Only applicable when 'accessPolicyType' is 'Custom access policy'.
	AccessPolicyType        string                                                                 `json:"accessPolicyType,omitempty"`        // The type of the access policy of the switch template port. Only applicable to access ports. Can be one of 'Open', 'Custom access policy', 'MAC allow list' or 'Sticky MAC allow list'.
	AllowedVLANs            string                                                                 `json:"allowedVlans,omitempty"`            // The VLANs allowed on the switch template port. Only applicable to trunk ports.
	DaiTrusted              *bool                                                                  `json:"daiTrusted,omitempty"`              // If true, ARP packets for this port will be considered trusted, and Dynamic ARP Inspection will allow the traffic.
	Dot3Az                  *RequestSwitchUpdateOrganizationConfigTemplateSwitchProfilePortDot3Az  `json:"dot3az,omitempty"`                  // dot3az settings for the port
	Enabled                 *bool                                                                  `json:"enabled,omitempty"`                 // The status of the switch template port.
	FlexibleStackingEnabled *bool                                                                  `json:"flexibleStackingEnabled,omitempty"` // For supported switches (e.g. MS420/MS425), whether or not the port has flexible stacking enabled.
	IsolationEnabled        *bool                                                                  `json:"isolationEnabled,omitempty"`        // The isolation status of the switch template port.
	LinkNegotiation         string                                                                 `json:"linkNegotiation,omitempty"`         // The link speed for the switch template port.
	MacAllowList            []string                                                               `json:"macAllowList,omitempty"`            // Only devices with MAC addresses specified in this list will have access to this port. Up to 20 MAC addresses can be defined. Only applicable when 'accessPolicyType' is 'MAC allow list'.
	Name                    string                                                                 `json:"name,omitempty"`                    // The name of the switch template port.
	PoeEnabled              *bool                                                                  `json:"poeEnabled,omitempty"`              // The PoE status of the switch template port.
	PortScheduleID          string                                                                 `json:"portScheduleId,omitempty"`          // The ID of the port schedule. A value of null will clear the port schedule.
	Profile                 *RequestSwitchUpdateOrganizationConfigTemplateSwitchProfilePortProfile `json:"profile,omitempty"`                 // Profile attributes
	RstpEnabled             *bool                                                                  `json:"rstpEnabled,omitempty"`             // The rapid spanning tree protocol status.
	StickyMacAllowList      []string                                                               `json:"stickyMacAllowList,omitempty"`      // The initial list of MAC addresses for sticky Mac allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
	StickyMacAllowListLimit *int                                                                   `json:"stickyMacAllowListLimit,omitempty"` // The maximum number of MAC addresses for sticky MAC allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
	StormControlEnabled     *bool                                                                  `json:"stormControlEnabled,omitempty"`     // The storm control status of the switch template port.
	StpGuard                string                                                                 `json:"stpGuard,omitempty"`                // The state of the STP guard ('disabled', 'root guard', 'bpdu guard' or 'loop guard').
	Tags                    []string                                                               `json:"tags,omitempty"`                    // The list of tags of the switch template port.
	Type                    string                                                                 `json:"type,omitempty"`                    // The type of the switch template port ('trunk', 'access' or 'stack').
	Udld                    string                                                                 `json:"udld,omitempty"`                    // The action to take when Unidirectional Link is detected (Alert only, Enforce). Default configuration is Alert only.
	VLAN                    *int                                                                   `json:"vlan,omitempty"`                    // The VLAN of the switch template port. For a trunk port, this is the native VLAN. A null value will clear the value set for trunk ports.
	VoiceVLAN               *int                                                                   `json:"voiceVlan,omitempty"`               // The voice VLAN of the switch template port. Only applicable to access ports.
}
type RequestSwitchUpdateOrganizationConfigTemplateSwitchProfilePortDot3Az struct {
	Enabled *bool `json:"enabled,omitempty"` // The Energy Efficient Ethernet status of the switch template port.
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


*/
func (s *SwitchService) GetDeviceSwitchPorts(serial string) (*ResponseSwitchGetDeviceSwitchPorts, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/switch/ports"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) GetDeviceSwitchPortsStatuses(serial string, getDeviceSwitchPortsStatusesQueryParams *GetDeviceSwitchPortsStatusesQueryParams) (*ResponseSwitchGetDeviceSwitchPortsStatuses, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/switch/ports/statuses"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) GetDeviceSwitchPortsStatusesPackets(serial string, getDeviceSwitchPortsStatusesPacketsQueryParams *GetDeviceSwitchPortsStatusesPacketsQueryParams) (*ResponseSwitchGetDeviceSwitchPortsStatusesPackets, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/switch/ports/statuses/packets"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) GetDeviceSwitchPort(serial string, portID string) (*ResponseSwitchGetDeviceSwitchPort, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/switch/ports/{portId}"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) GetDeviceSwitchRoutingInterfaces(serial string) (*ResponseSwitchGetDeviceSwitchRoutingInterfaces, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/switch/routing/interfaces"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) GetDeviceSwitchRoutingInterface(serial string, interfaceID string) (*ResponseSwitchGetDeviceSwitchRoutingInterface, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/switch/routing/interfaces/{interfaceId}"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) GetDeviceSwitchRoutingInterfaceDhcp(serial string, interfaceID string) (*ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcp, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/switch/routing/interfaces/{interfaceId}/dhcp"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) GetDeviceSwitchRoutingStaticRoutes(serial string) (*ResponseSwitchGetDeviceSwitchRoutingStaticRoutes, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/switch/routing/staticRoutes"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) GetDeviceSwitchRoutingStaticRoute(serial string, staticRouteID string) (*ResponseSwitchGetDeviceSwitchRoutingStaticRoute, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/switch/routing/staticRoutes/{staticRouteId}"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) GetDeviceSwitchWarmSpare(serial string) (*ResponseSwitchGetDeviceSwitchWarmSpare, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/switch/warmSpare"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) GetNetworkSwitchAccessControlLists(networkID string) (*ResponseSwitchGetNetworkSwitchAccessControlLists, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/accessControlLists"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) GetNetworkSwitchAccessPolicies(networkID string) (*ResponseSwitchGetNetworkSwitchAccessPolicies, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/accessPolicies"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) GetNetworkSwitchAccessPolicy(networkID string, accessPolicyNumber string) (*ResponseSwitchGetNetworkSwitchAccessPolicy, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/accessPolicies/{accessPolicyNumber}"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) GetNetworkSwitchAlternateManagementInterface(networkID string) (*ResponseSwitchGetNetworkSwitchAlternateManagementInterface, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/alternateManagementInterface"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) GetNetworkSwitchDhcpV4ServersSeen(networkID string, getNetworkSwitchDhcpV4ServersSeenQueryParams *GetNetworkSwitchDhcpV4ServersSeenQueryParams) (*ResponseSwitchGetNetworkSwitchDhcpV4ServersSeen, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/dhcp/v4/servers/seen"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) GetNetworkSwitchDhcpServerPolicy(networkID string) (*ResponseSwitchGetNetworkSwitchDhcpServerPolicy, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/dhcpServerPolicy"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers(networkID string, getNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersQueryParams *GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersQueryParams) (*ResponseSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/dhcpServerPolicy/arpInspection/trustedServers"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice(networkID string, getNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceQueryParams *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceQueryParams) (*ResponseSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/dhcpServerPolicy/arpInspection/warnings/byDevice"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) GetNetworkSwitchDscpToCosMappings(networkID string) (*ResponseSwitchGetNetworkSwitchDscpToCosMappings, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/dscpToCosMappings"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) GetNetworkSwitchLinkAggregations(networkID string) (*ResponseSwitchGetNetworkSwitchLinkAggregations, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/linkAggregations"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) GetNetworkSwitchMtu(networkID string) (*ResponseSwitchGetNetworkSwitchMtu, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/mtu"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) GetNetworkSwitchPortSchedules(networkID string) (*ResponseSwitchGetNetworkSwitchPortSchedules, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/portSchedules"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) GetNetworkSwitchQosRules(networkID string) (*ResponseSwitchGetNetworkSwitchQosRules, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/qosRules"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) GetNetworkSwitchQosRulesOrder(networkID string) (*ResponseSwitchGetNetworkSwitchQosRulesOrder, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/qosRules/order"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) GetNetworkSwitchQosRule(networkID string, qosRuleID string) (*ResponseSwitchGetNetworkSwitchQosRule, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/qosRules/{qosRuleId}"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) GetNetworkSwitchRoutingMulticast(networkID string) (*ResponseSwitchGetNetworkSwitchRoutingMulticast, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/routing/multicast"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) GetNetworkSwitchRoutingMulticastRendezvousPoints(networkID string) (*ResponseSwitchGetNetworkSwitchRoutingMulticastRendezvousPoints, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/routing/multicast/rendezvousPoints"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) GetNetworkSwitchRoutingMulticastRendezvousPoint(networkID string, rendezvousPointID string) (*ResponseSwitchGetNetworkSwitchRoutingMulticastRendezvousPoint, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/routing/multicast/rendezvousPoints/{rendezvousPointId}"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) GetNetworkSwitchRoutingOspf(networkID string) (*ResponseSwitchGetNetworkSwitchRoutingOspf, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/routing/ospf"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) GetNetworkSwitchSettings(networkID string) (*ResponseSwitchGetNetworkSwitchSettings, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/settings"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) GetNetworkSwitchStacks(networkID string) (*ResponseSwitchGetNetworkSwitchStacks, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/stacks"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) GetNetworkSwitchStack(networkID string, switchStackID string) (*ResponseSwitchGetNetworkSwitchStack, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/stacks/{switchStackId}"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) GetNetworkSwitchStackRoutingInterfaces(networkID string, switchStackID string) (*ResponseSwitchGetNetworkSwitchStackRoutingInterfaces, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) GetNetworkSwitchStackRoutingInterface(networkID string, switchStackID string, interfaceID string) (*ResponseSwitchGetNetworkSwitchStackRoutingInterface, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId}"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) GetNetworkSwitchStackRoutingInterfaceDhcp(networkID string, switchStackID string, interfaceID string) (*ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcp, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId}/dhcp"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) GetNetworkSwitchStackRoutingStaticRoutes(networkID string, switchStackID string) (*ResponseSwitchGetNetworkSwitchStackRoutingStaticRoutes, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/stacks/{switchStackId}/routing/staticRoutes"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) GetNetworkSwitchStackRoutingStaticRoute(networkID string, switchStackID string, staticRouteID string) (*ResponseSwitchGetNetworkSwitchStackRoutingStaticRoute, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/stacks/{switchStackId}/routing/staticRoutes/{staticRouteId}"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) GetNetworkSwitchStormControl(networkID string) (*ResponseSwitchGetNetworkSwitchStormControl, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/stormControl"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) GetNetworkSwitchStp(networkID string) (*ResponseSwitchGetNetworkSwitchStp, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/stp"
	s.rateLimiterBucket.Wait(1)
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

//GetOrganizationConfigTemplateSwitchProfiles List the switch templates for your switch template configuration
/* List the switch templates for your switch template configuration

@param organizationID organizationId path parameter. Organization ID
@param configTemplateID configTemplateId path parameter. Config template ID


*/
func (s *SwitchService) GetOrganizationConfigTemplateSwitchProfiles(organizationID string, configTemplateID string) (*ResponseSwitchGetOrganizationConfigTemplateSwitchProfiles, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/configTemplates/{configTemplateId}/switch/profiles"
	s.rateLimiterBucket.Wait(1)
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

//GetOrganizationConfigTemplateSwitchProfilePorts Return all the ports of a switch template
/* Return all the ports of a switch template

@param organizationID organizationId path parameter. Organization ID
@param configTemplateID configTemplateId path parameter. Config template ID
@param profileID profileId path parameter. Profile ID


*/
func (s *SwitchService) GetOrganizationConfigTemplateSwitchProfilePorts(organizationID string, configTemplateID string, profileID string) (*ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePorts, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/configTemplates/{configTemplateId}/switch/profiles/{profileId}/ports"
	s.rateLimiterBucket.Wait(1)
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

//GetOrganizationConfigTemplateSwitchProfilePort Return a switch template port
/* Return a switch template port

@param organizationID organizationId path parameter. Organization ID
@param configTemplateID configTemplateId path parameter. Config template ID
@param profileID profileId path parameter. Profile ID
@param portID portId path parameter. Port ID


*/
func (s *SwitchService) GetOrganizationConfigTemplateSwitchProfilePort(organizationID string, configTemplateID string, profileID string, portID string) (*ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePort, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/configTemplates/{configTemplateId}/switch/profiles/{profileId}/ports/{portId}"
	s.rateLimiterBucket.Wait(1)
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

//GetOrganizationSummarySwitchPowerHistory Returns the total PoE power draw for all switch ports in the organization over the requested timespan (by default the last 24 hours)
/* Returns the total PoE power draw for all switch ports in the organization over the requested timespan (by default the last 24 hours). The returned array is a newest-first list of intervals. The time between intervals depends on the requested timespan with 20 minute intervals used for timespans up to 1 day, 4 hour intervals used for timespans up to 2 weeks, and 1 day intervals for timespans larger than 2 weeks.

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationSummarySwitchPowerHistoryQueryParams Filtering parameter


*/
func (s *SwitchService) GetOrganizationSummarySwitchPowerHistory(organizationID string, getOrganizationSummarySwitchPowerHistoryQueryParams *GetOrganizationSummarySwitchPowerHistoryQueryParams) (*ResponseSwitchGetOrganizationSummarySwitchPowerHistory, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/summary/switch/power/history"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationSummarySwitchPowerHistoryQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSwitchGetOrganizationSummarySwitchPowerHistory{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationSummarySwitchPowerHistory")
	}

	result := response.Result().(*ResponseSwitchGetOrganizationSummarySwitchPowerHistory)
	return result, response, err

}

//GetOrganizationSwitchPortsBySwitch List the switchports in an organization by switch
/* List the switchports in an organization by switch

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationSwitchPortsBySwitchQueryParams Filtering parameter


*/
func (s *SwitchService) GetOrganizationSwitchPortsBySwitch(organizationID string, getOrganizationSwitchPortsBySwitchQueryParams *GetOrganizationSwitchPortsBySwitchQueryParams) (*ResponseSwitchGetOrganizationSwitchPortsBySwitch, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/switch/ports/bySwitch"
	s.rateLimiterBucket.Wait(1)
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

//GetOrganizationSwitchPortsClientsOverviewByDevice List the number of clients for all switchports with at least one online client in an organization.
/* List the number of clients for all switchports with at least one online client in an organization.

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationSwitchPortsClientsOverviewByDeviceQueryParams Filtering parameter


*/
func (s *SwitchService) GetOrganizationSwitchPortsClientsOverviewByDevice(organizationID string, getOrganizationSwitchPortsClientsOverviewByDeviceQueryParams *GetOrganizationSwitchPortsClientsOverviewByDeviceQueryParams) (*ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDevice, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/switch/ports/clients/overview/byDevice"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationSwitchPortsClientsOverviewByDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationSwitchPortsClientsOverviewByDevice")
	}

	result := response.Result().(*ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDevice)
	return result, response, err

}

//GetOrganizationSwitchPortsOverview Returns the counts of all active ports for the requested timespan, grouped by speed
/* Returns the counts of all active ports for the requested timespan, grouped by speed. An active port is a port that at any point during the timeframe is observed to be connected to a responsive device and isn't configured to be disabled. For a port that is observed at multiple speeds during the timeframe, it will be counted at the highest speed observed. The number of inactive ports, and the total number of ports are also provided. Only ports on switches online during the timeframe will be represented and a port is only guaranteed to be present if its switch was online for at least 6 hours of the timeframe.

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationSwitchPortsOverviewQueryParams Filtering parameter


*/
func (s *SwitchService) GetOrganizationSwitchPortsOverview(organizationID string, getOrganizationSwitchPortsOverviewQueryParams *GetOrganizationSwitchPortsOverviewQueryParams) (*ResponseSwitchGetOrganizationSwitchPortsOverview, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/switch/ports/overview"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationSwitchPortsOverviewQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSwitchGetOrganizationSwitchPortsOverview{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationSwitchPortsOverview")
	}

	result := response.Result().(*ResponseSwitchGetOrganizationSwitchPortsOverview)
	return result, response, err

}

//GetOrganizationSwitchPortsStatusesBySwitch List the switchports in an organization
/* List the switchports in an organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationSwitchPortsStatusesBySwitchQueryParams Filtering parameter


*/
func (s *SwitchService) GetOrganizationSwitchPortsStatusesBySwitch(organizationID string, getOrganizationSwitchPortsStatusesBySwitchQueryParams *GetOrganizationSwitchPortsStatusesBySwitchQueryParams) (*ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitch, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/switch/ports/statuses/bySwitch"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationSwitchPortsStatusesBySwitchQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitch{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationSwitchPortsStatusesBySwitch")
	}

	result := response.Result().(*ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitch)
	return result, response, err

}

//GetOrganizationSwitchPortsTopologyDiscoveryByDevice List most recently seen LLDP/CDP discovery and topology information per switch port in an organization.
/* List most recently seen LLDP/CDP discovery and topology information per switch port in an organization.

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationSwitchPortsTopologyDiscoveryByDeviceQueryParams Filtering parameter


*/
func (s *SwitchService) GetOrganizationSwitchPortsTopologyDiscoveryByDevice(organizationID string, getOrganizationSwitchPortsTopologyDiscoveryByDeviceQueryParams *GetOrganizationSwitchPortsTopologyDiscoveryByDeviceQueryParams) (*ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDevice, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/switch/ports/topology/discovery/byDevice"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationSwitchPortsTopologyDiscoveryByDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationSwitchPortsTopologyDiscoveryByDevice")
	}

	result := response.Result().(*ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDevice)
	return result, response, err

}

//CycleDeviceSwitchPorts Cycle a set of switch ports
/* Cycle a set of switch ports

@param serial serial path parameter.


*/

func (s *SwitchService) CycleDeviceSwitchPorts(serial string, requestSwitchCycleDeviceSwitchPorts *RequestSwitchCycleDeviceSwitchPorts) (*ResponseSwitchCycleDeviceSwitchPorts, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/switch/ports/cycle"
	s.rateLimiterBucket.Wait(1)
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


*/

func (s *SwitchService) CreateDeviceSwitchRoutingInterface(serial string, requestSwitchCreateDeviceSwitchRoutingInterface *RequestSwitchCreateDeviceSwitchRoutingInterface) (*ResponseSwitchCreateDeviceSwitchRoutingInterface, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/switch/routing/interfaces"
	s.rateLimiterBucket.Wait(1)
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


*/

func (s *SwitchService) CreateDeviceSwitchRoutingStaticRoute(serial string, requestSwitchCreateDeviceSwitchRoutingStaticRoute *RequestSwitchCreateDeviceSwitchRoutingStaticRoute) (*ResponseSwitchCreateDeviceSwitchRoutingStaticRoute, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/switch/routing/staticRoutes"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchCreateDeviceSwitchRoutingStaticRoute).
		SetResult(&ResponseSwitchCreateDeviceSwitchRoutingStaticRoute{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateDeviceSwitchRoutingStaticRoute")
	}

	result := response.Result().(*ResponseSwitchCreateDeviceSwitchRoutingStaticRoute)
	return result, response, err

}

//CreateNetworkSwitchAccessPolicy Create an access policy for a switch network
/* Create an access policy for a switch network. If you would like to enable Meraki Authentication, set radiusServers to empty array.

@param networkID networkId path parameter. Network ID


*/

func (s *SwitchService) CreateNetworkSwitchAccessPolicy(networkID string, requestSwitchCreateNetworkSwitchAccessPolicy *RequestSwitchCreateNetworkSwitchAccessPolicy) (*ResponseSwitchCreateNetworkSwitchAccessPolicy, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/accessPolicies"
	s.rateLimiterBucket.Wait(1)
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


*/

func (s *SwitchService) CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(networkID string, requestSwitchCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer *RequestSwitchCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer) (*ResponseSwitchCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/dhcpServerPolicy/arpInspection/trustedServers"
	s.rateLimiterBucket.Wait(1)
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


*/

func (s *SwitchService) CreateNetworkSwitchLinkAggregation(networkID string, requestSwitchCreateNetworkSwitchLinkAggregation *RequestSwitchCreateNetworkSwitchLinkAggregation) (*ResponseSwitchCreateNetworkSwitchLinkAggregation, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/linkAggregations"
	s.rateLimiterBucket.Wait(1)
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


*/

func (s *SwitchService) CreateNetworkSwitchPortSchedule(networkID string, requestSwitchCreateNetworkSwitchPortSchedule *RequestSwitchCreateNetworkSwitchPortSchedule) (*ResponseSwitchCreateNetworkSwitchPortSchedule, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/portSchedules"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchCreateNetworkSwitchPortSchedule).
		SetResult(&ResponseSwitchCreateNetworkSwitchPortSchedule{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateNetworkSwitchPortSchedule")
	}

	result := response.Result().(*ResponseSwitchCreateNetworkSwitchPortSchedule)
	return result, response, err

}

//CreateNetworkSwitchQosRule Add a quality of service rule
/* Add a quality of service rule

@param networkID networkId path parameter. Network ID


*/

func (s *SwitchService) CreateNetworkSwitchQosRule(networkID string, requestSwitchCreateNetworkSwitchQosRule *RequestSwitchCreateNetworkSwitchQosRule) (*ResponseSwitchCreateNetworkSwitchQosRule, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/qosRules"
	s.rateLimiterBucket.Wait(1)
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


*/

func (s *SwitchService) CreateNetworkSwitchRoutingMulticastRendezvousPoint(networkID string, requestSwitchCreateNetworkSwitchRoutingMulticastRendezvousPoint *RequestSwitchCreateNetworkSwitchRoutingMulticastRendezvousPoint) (*ResponseSwitchCreateNetworkSwitchRoutingMulticastRendezvousPoint, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/routing/multicast/rendezvousPoints"
	s.rateLimiterBucket.Wait(1)
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

//CreateNetworkSwitchStack Create a switch stack
/* Create a switch stack

@param networkID networkId path parameter. Network ID


*/

func (s *SwitchService) CreateNetworkSwitchStack(networkID string, requestSwitchCreateNetworkSwitchStack *RequestSwitchCreateNetworkSwitchStack) (*ResponseSwitchCreateNetworkSwitchStack, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/stacks"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchCreateNetworkSwitchStack).
		SetResult(&ResponseSwitchCreateNetworkSwitchStack{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateNetworkSwitchStack")
	}

	result := response.Result().(*ResponseSwitchCreateNetworkSwitchStack)
	return result, response, err

}

//AddNetworkSwitchStack Add a switch to a stack
/* Add a switch to a stack

@param networkID networkId path parameter. Network ID
@param switchStackID switchStackId path parameter. Switch stack ID


*/

func (s *SwitchService) AddNetworkSwitchStack(networkID string, switchStackID string, requestSwitchAddNetworkSwitchStack *RequestSwitchAddNetworkSwitchStack) (*ResponseSwitchAddNetworkSwitchStack, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/stacks/{switchStackId}/add"
	s.rateLimiterBucket.Wait(1)
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


*/

func (s *SwitchService) RemoveNetworkSwitchStack(networkID string, switchStackID string, requestSwitchRemoveNetworkSwitchStack *RequestSwitchRemoveNetworkSwitchStack) (*ResponseSwitchRemoveNetworkSwitchStack, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/stacks/{switchStackId}/remove"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{switchStackId}", fmt.Sprintf("%v", switchStackID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchRemoveNetworkSwitchStack).
		SetResult(&ResponseSwitchRemoveNetworkSwitchStack{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation RemoveNetworkSwitchStack")
	}

	result := response.Result().(*ResponseSwitchRemoveNetworkSwitchStack)
	return result, response, err

}

//CreateNetworkSwitchStackRoutingInterface Create a layer 3 interface for a switch stack
/* Create a layer 3 interface for a switch stack

@param networkID networkId path parameter. Network ID
@param switchStackID switchStackId path parameter. Switch stack ID


*/

func (s *SwitchService) CreateNetworkSwitchStackRoutingInterface(networkID string, switchStackID string, requestSwitchCreateNetworkSwitchStackRoutingInterface *RequestSwitchCreateNetworkSwitchStackRoutingInterface) (*ResponseSwitchCreateNetworkSwitchStackRoutingInterface, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{switchStackId}", fmt.Sprintf("%v", switchStackID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchCreateNetworkSwitchStackRoutingInterface).
		SetResult(&ResponseSwitchCreateNetworkSwitchStackRoutingInterface{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateNetworkSwitchStackRoutingInterface")
	}

	result := response.Result().(*ResponseSwitchCreateNetworkSwitchStackRoutingInterface)
	return result, response, err

}

//CreateNetworkSwitchStackRoutingStaticRoute Create a layer 3 static route for a switch stack
/* Create a layer 3 static route for a switch stack

@param networkID networkId path parameter. Network ID
@param switchStackID switchStackId path parameter. Switch stack ID


*/

func (s *SwitchService) CreateNetworkSwitchStackRoutingStaticRoute(networkID string, switchStackID string, requestSwitchCreateNetworkSwitchStackRoutingStaticRoute *RequestSwitchCreateNetworkSwitchStackRoutingStaticRoute) (*ResponseSwitchCreateNetworkSwitchStackRoutingStaticRoute, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/stacks/{switchStackId}/routing/staticRoutes"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{switchStackId}", fmt.Sprintf("%v", switchStackID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchCreateNetworkSwitchStackRoutingStaticRoute).
		SetResult(&ResponseSwitchCreateNetworkSwitchStackRoutingStaticRoute{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateNetworkSwitchStackRoutingStaticRoute")
	}

	result := response.Result().(*ResponseSwitchCreateNetworkSwitchStackRoutingStaticRoute)
	return result, response, err

}

//CloneOrganizationSwitchDevices Clone port-level and some switch-level configuration settings from a source switch to one or more target switches
/* Clone port-level and some switch-level configuration settings from a source switch to one or more target switches. Cloned settings include: Aggregation Groups, Power Settings, Multicast Settings, MTU Configuration, STP Bridge priority, Port Mirroring

@param organizationID organizationId path parameter. Organization ID


*/

func (s *SwitchService) CloneOrganizationSwitchDevices(organizationID string, requestSwitchCloneOrganizationSwitchDevices *RequestSwitchCloneOrganizationSwitchDevices) (*ResponseSwitchCloneOrganizationSwitchDevices, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/switch/devices/clone"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchCloneOrganizationSwitchDevices).
		SetResult(&ResponseSwitchCloneOrganizationSwitchDevices{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CloneOrganizationSwitchDevices")
	}

	result := response.Result().(*ResponseSwitchCloneOrganizationSwitchDevices)
	return result, response, err

}

//UpdateDeviceSwitchPort Update a switch port
/* Update a switch port

@param serial serial path parameter.
@param portID portId path parameter. Port ID
*/
func (s *SwitchService) UpdateDeviceSwitchPort(serial string, portID string, requestSwitchUpdateDeviceSwitchPort *RequestSwitchUpdateDeviceSwitchPort) (*ResponseSwitchUpdateDeviceSwitchPort, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/switch/ports/{portId}"
	s.rateLimiterBucket.Wait(1)
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
	s.rateLimiterBucket.Wait(1)
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
func (s *SwitchService) UpdateDeviceSwitchRoutingInterfaceDhcp(serial string, interfaceID string, requestSwitchUpdateDeviceSwitchRoutingInterfaceDhcp *RequestSwitchUpdateDeviceSwitchRoutingInterfaceDhcp) (*ResponseSwitchUpdateDeviceSwitchRoutingInterfaceDhcp, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/switch/routing/interfaces/{interfaceId}/dhcp"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)
	path = strings.Replace(path, "{interfaceId}", fmt.Sprintf("%v", interfaceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateDeviceSwitchRoutingInterfaceDhcp).
		SetResult(&ResponseSwitchUpdateDeviceSwitchRoutingInterfaceDhcp{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateDeviceSwitchRoutingInterfaceDhcp")
	}

	result := response.Result().(*ResponseSwitchUpdateDeviceSwitchRoutingInterfaceDhcp)
	return result, response, err

}

//UpdateDeviceSwitchRoutingStaticRoute Update a layer 3 static route for a switch
/* Update a layer 3 static route for a switch

@param serial serial path parameter.
@param staticRouteID staticRouteId path parameter. Static route ID
*/
func (s *SwitchService) UpdateDeviceSwitchRoutingStaticRoute(serial string, staticRouteID string, requestSwitchUpdateDeviceSwitchRoutingStaticRoute *RequestSwitchUpdateDeviceSwitchRoutingStaticRoute) (*ResponseSwitchUpdateDeviceSwitchRoutingStaticRoute, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/switch/routing/staticRoutes/{staticRouteId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)
	path = strings.Replace(path, "{staticRouteId}", fmt.Sprintf("%v", staticRouteID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateDeviceSwitchRoutingStaticRoute).
		SetResult(&ResponseSwitchUpdateDeviceSwitchRoutingStaticRoute{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateDeviceSwitchRoutingStaticRoute")
	}

	result := response.Result().(*ResponseSwitchUpdateDeviceSwitchRoutingStaticRoute)
	return result, response, err

}

//UpdateDeviceSwitchWarmSpare Update warm spare configuration for a switch
/* Update warm spare configuration for a switch. The spare will use the same L3 configuration as the primary. Note that this will irreversibly destroy any existing L3 configuration on the spare.

@param serial serial path parameter.
*/
func (s *SwitchService) UpdateDeviceSwitchWarmSpare(serial string, requestSwitchUpdateDeviceSwitchWarmSpare *RequestSwitchUpdateDeviceSwitchWarmSpare) (*ResponseSwitchUpdateDeviceSwitchWarmSpare, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/switch/warmSpare"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateDeviceSwitchWarmSpare).
		SetResult(&ResponseSwitchUpdateDeviceSwitchWarmSpare{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateDeviceSwitchWarmSpare")
	}

	result := response.Result().(*ResponseSwitchUpdateDeviceSwitchWarmSpare)
	return result, response, err

}

//UpdateNetworkSwitchAccessControlLists Update the access control lists for a MS network
/* Update the access control lists for a MS network

@param networkID networkId path parameter. Network ID
*/
func (s *SwitchService) UpdateNetworkSwitchAccessControlLists(networkID string, requestSwitchUpdateNetworkSwitchAccessControlLists *RequestSwitchUpdateNetworkSwitchAccessControlLists) (*ResponseSwitchUpdateNetworkSwitchAccessControlLists, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/accessControlLists"
	s.rateLimiterBucket.Wait(1)
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
	s.rateLimiterBucket.Wait(1)
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
func (s *SwitchService) UpdateNetworkSwitchAlternateManagementInterface(networkID string, requestSwitchUpdateNetworkSwitchAlternateManagementInterface *RequestSwitchUpdateNetworkSwitchAlternateManagementInterface) (*ResponseSwitchUpdateNetworkSwitchAlternateManagementInterface, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/alternateManagementInterface"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateNetworkSwitchAlternateManagementInterface).
		SetResult(&ResponseSwitchUpdateNetworkSwitchAlternateManagementInterface{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkSwitchAlternateManagementInterface")
	}

	result := response.Result().(*ResponseSwitchUpdateNetworkSwitchAlternateManagementInterface)
	return result, response, err

}

//UpdateNetworkSwitchDhcpServerPolicy Update the DHCP server settings
/* Update the DHCP server settings. Blocked/allowed servers are only applied when default policy is allow/block, respectively

@param networkID networkId path parameter. Network ID
*/
func (s *SwitchService) UpdateNetworkSwitchDhcpServerPolicy(networkID string, requestSwitchUpdateNetworkSwitchDhcpServerPolicy *RequestSwitchUpdateNetworkSwitchDhcpServerPolicy) (*ResponseSwitchUpdateNetworkSwitchDhcpServerPolicy, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/dhcpServerPolicy"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateNetworkSwitchDhcpServerPolicy).
		SetResult(&ResponseSwitchUpdateNetworkSwitchDhcpServerPolicy{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkSwitchDhcpServerPolicy")
	}

	result := response.Result().(*ResponseSwitchUpdateNetworkSwitchDhcpServerPolicy)
	return result, response, err

}

//UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer Update a server that is trusted by Dynamic ARP Inspection on this network
/* Update a server that is trusted by Dynamic ARP Inspection on this network

@param networkID networkId path parameter. Network ID
@param trustedServerID trustedServerId path parameter. Trusted server ID
*/
func (s *SwitchService) UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(networkID string, trustedServerID string, requestSwitchUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer *RequestSwitchUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer) (*ResponseSwitchUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/dhcpServerPolicy/arpInspection/trustedServers/{trustedServerId}"
	s.rateLimiterBucket.Wait(1)
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
func (s *SwitchService) UpdateNetworkSwitchDscpToCosMappings(networkID string, requestSwitchUpdateNetworkSwitchDscpToCosMappings *RequestSwitchUpdateNetworkSwitchDscpToCosMappings) (*ResponseSwitchUpdateNetworkSwitchDscpToCosMappings, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/dscpToCosMappings"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateNetworkSwitchDscpToCosMappings).
		SetResult(&ResponseSwitchUpdateNetworkSwitchDscpToCosMappings{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkSwitchDscpToCosMappings")
	}

	result := response.Result().(*ResponseSwitchUpdateNetworkSwitchDscpToCosMappings)
	return result, response, err

}

//UpdateNetworkSwitchLinkAggregation Update a link aggregation group
/* Update a link aggregation group

@param networkID networkId path parameter. Network ID
@param linkAggregationID linkAggregationId path parameter. Link aggregation ID
*/
func (s *SwitchService) UpdateNetworkSwitchLinkAggregation(networkID string, linkAggregationID string, requestSwitchUpdateNetworkSwitchLinkAggregation *RequestSwitchUpdateNetworkSwitchLinkAggregation) (*ResponseSwitchUpdateNetworkSwitchLinkAggregation, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/linkAggregations/{linkAggregationId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{linkAggregationId}", fmt.Sprintf("%v", linkAggregationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateNetworkSwitchLinkAggregation).
		SetResult(&ResponseSwitchUpdateNetworkSwitchLinkAggregation{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkSwitchLinkAggregation")
	}

	result := response.Result().(*ResponseSwitchUpdateNetworkSwitchLinkAggregation)
	return result, response, err

}

//UpdateNetworkSwitchMtu Update the MTU configuration
/* Update the MTU configuration

@param networkID networkId path parameter. Network ID
*/
func (s *SwitchService) UpdateNetworkSwitchMtu(networkID string, requestSwitchUpdateNetworkSwitchMtu *RequestSwitchUpdateNetworkSwitchMtu) (*ResponseSwitchUpdateNetworkSwitchMtu, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/mtu"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateNetworkSwitchMtu).
		SetResult(&ResponseSwitchUpdateNetworkSwitchMtu{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkSwitchMtu")
	}

	result := response.Result().(*ResponseSwitchUpdateNetworkSwitchMtu)
	return result, response, err

}

//UpdateNetworkSwitchPortSchedule Update a switch port schedule
/* Update a switch port schedule

@param networkID networkId path parameter. Network ID
@param portScheduleID portScheduleId path parameter. Port schedule ID
*/
func (s *SwitchService) UpdateNetworkSwitchPortSchedule(networkID string, portScheduleID string, requestSwitchUpdateNetworkSwitchPortSchedule *RequestSwitchUpdateNetworkSwitchPortSchedule) (*ResponseSwitchUpdateNetworkSwitchPortSchedule, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/portSchedules/{portScheduleId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{portScheduleId}", fmt.Sprintf("%v", portScheduleID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateNetworkSwitchPortSchedule).
		SetResult(&ResponseSwitchUpdateNetworkSwitchPortSchedule{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkSwitchPortSchedule")
	}

	result := response.Result().(*ResponseSwitchUpdateNetworkSwitchPortSchedule)
	return result, response, err

}

//UpdateNetworkSwitchQosRulesOrder Update the order in which the rules should be processed by the switch
/* Update the order in which the rules should be processed by the switch

@param networkID networkId path parameter. Network ID
*/
func (s *SwitchService) UpdateNetworkSwitchQosRulesOrder(networkID string, requestSwitchUpdateNetworkSwitchQosRulesOrder *RequestSwitchUpdateNetworkSwitchQosRulesOrder) (*ResponseSwitchUpdateNetworkSwitchQosRulesOrder, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/qosRules/order"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateNetworkSwitchQosRulesOrder).
		SetResult(&ResponseSwitchUpdateNetworkSwitchQosRulesOrder{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkSwitchQosRulesOrder")
	}

	result := response.Result().(*ResponseSwitchUpdateNetworkSwitchQosRulesOrder)
	return result, response, err

}

//UpdateNetworkSwitchQosRule Update a quality of service rule
/* Update a quality of service rule

@param networkID networkId path parameter. Network ID
@param qosRuleID qosRuleId path parameter. Qos rule ID
*/
func (s *SwitchService) UpdateNetworkSwitchQosRule(networkID string, qosRuleID string, requestSwitchUpdateNetworkSwitchQosRule *RequestSwitchUpdateNetworkSwitchQosRule) (*ResponseSwitchUpdateNetworkSwitchQosRule, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/qosRules/{qosRuleId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{qosRuleId}", fmt.Sprintf("%v", qosRuleID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateNetworkSwitchQosRule).
		SetResult(&ResponseSwitchUpdateNetworkSwitchQosRule{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkSwitchQosRule")
	}

	result := response.Result().(*ResponseSwitchUpdateNetworkSwitchQosRule)
	return result, response, err

}

//UpdateNetworkSwitchRoutingMulticast Update multicast settings for a network
/* Update multicast settings for a network

@param networkID networkId path parameter. Network ID
*/
func (s *SwitchService) UpdateNetworkSwitchRoutingMulticast(networkID string, requestSwitchUpdateNetworkSwitchRoutingMulticast *RequestSwitchUpdateNetworkSwitchRoutingMulticast) (*ResponseSwitchUpdateNetworkSwitchRoutingMulticast, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/routing/multicast"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateNetworkSwitchRoutingMulticast).
		SetResult(&ResponseSwitchUpdateNetworkSwitchRoutingMulticast{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkSwitchRoutingMulticast")
	}

	result := response.Result().(*ResponseSwitchUpdateNetworkSwitchRoutingMulticast)
	return result, response, err

}

//UpdateNetworkSwitchRoutingMulticastRendezvousPoint Update a multicast rendezvous point
/* Update a multicast rendezvous point

@param networkID networkId path parameter. Network ID
@param rendezvousPointID rendezvousPointId path parameter. Rendezvous point ID
*/
func (s *SwitchService) UpdateNetworkSwitchRoutingMulticastRendezvousPoint(networkID string, rendezvousPointID string, requestSwitchUpdateNetworkSwitchRoutingMulticastRendezvousPoint *RequestSwitchUpdateNetworkSwitchRoutingMulticastRendezvousPoint) (*ResponseSwitchUpdateNetworkSwitchRoutingMulticastRendezvousPoint, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/routing/multicast/rendezvousPoints/{rendezvousPointId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{rendezvousPointId}", fmt.Sprintf("%v", rendezvousPointID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateNetworkSwitchRoutingMulticastRendezvousPoint).
		SetResult(&ResponseSwitchUpdateNetworkSwitchRoutingMulticastRendezvousPoint{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkSwitchRoutingMulticastRendezvousPoint")
	}

	result := response.Result().(*ResponseSwitchUpdateNetworkSwitchRoutingMulticastRendezvousPoint)
	return result, response, err

}

//UpdateNetworkSwitchRoutingOspf Update layer 3 OSPF routing configuration
/* Update layer 3 OSPF routing configuration

@param networkID networkId path parameter. Network ID
*/
func (s *SwitchService) UpdateNetworkSwitchRoutingOspf(networkID string, requestSwitchUpdateNetworkSwitchRoutingOspf *RequestSwitchUpdateNetworkSwitchRoutingOspf) (*ResponseSwitchUpdateNetworkSwitchRoutingOspf, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/routing/ospf"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateNetworkSwitchRoutingOspf).
		SetResult(&ResponseSwitchUpdateNetworkSwitchRoutingOspf{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkSwitchRoutingOspf")
	}

	result := response.Result().(*ResponseSwitchUpdateNetworkSwitchRoutingOspf)
	return result, response, err

}

//UpdateNetworkSwitchSettings Update switch network settings
/* Update switch network settings

@param networkID networkId path parameter. Network ID
*/
func (s *SwitchService) UpdateNetworkSwitchSettings(networkID string, requestSwitchUpdateNetworkSwitchSettings *RequestSwitchUpdateNetworkSwitchSettings) (*ResponseSwitchUpdateNetworkSwitchSettings, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/settings"
	s.rateLimiterBucket.Wait(1)
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
func (s *SwitchService) UpdateNetworkSwitchStackRoutingInterface(networkID string, switchStackID string, interfaceID string, requestSwitchUpdateNetworkSwitchStackRoutingInterface *RequestSwitchUpdateNetworkSwitchStackRoutingInterface) (*ResponseSwitchUpdateNetworkSwitchStackRoutingInterface, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{switchStackId}", fmt.Sprintf("%v", switchStackID), -1)
	path = strings.Replace(path, "{interfaceId}", fmt.Sprintf("%v", interfaceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateNetworkSwitchStackRoutingInterface).
		SetResult(&ResponseSwitchUpdateNetworkSwitchStackRoutingInterface{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkSwitchStackRoutingInterface")
	}

	result := response.Result().(*ResponseSwitchUpdateNetworkSwitchStackRoutingInterface)
	return result, response, err

}

//UpdateNetworkSwitchStackRoutingInterfaceDhcp Update a layer 3 interface DHCP configuration for a switch stack
/* Update a layer 3 interface DHCP configuration for a switch stack

@param networkID networkId path parameter. Network ID
@param switchStackID switchStackId path parameter. Switch stack ID
@param interfaceID interfaceId path parameter. Interface ID
*/
func (s *SwitchService) UpdateNetworkSwitchStackRoutingInterfaceDhcp(networkID string, switchStackID string, interfaceID string, requestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcp *RequestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcp) (*ResponseSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcp, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId}/dhcp"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{switchStackId}", fmt.Sprintf("%v", switchStackID), -1)
	path = strings.Replace(path, "{interfaceId}", fmt.Sprintf("%v", interfaceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcp).
		SetResult(&ResponseSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcp{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkSwitchStackRoutingInterfaceDhcp")
	}

	result := response.Result().(*ResponseSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcp)
	return result, response, err

}

//UpdateNetworkSwitchStackRoutingStaticRoute Update a layer 3 static route for a switch stack
/* Update a layer 3 static route for a switch stack

@param networkID networkId path parameter. Network ID
@param switchStackID switchStackId path parameter. Switch stack ID
@param staticRouteID staticRouteId path parameter. Static route ID
*/
func (s *SwitchService) UpdateNetworkSwitchStackRoutingStaticRoute(networkID string, switchStackID string, staticRouteID string, requestSwitchUpdateNetworkSwitchStackRoutingStaticRoute *RequestSwitchUpdateNetworkSwitchStackRoutingStaticRoute) (*ResponseSwitchUpdateNetworkSwitchStackRoutingStaticRoute, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/stacks/{switchStackId}/routing/staticRoutes/{staticRouteId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{switchStackId}", fmt.Sprintf("%v", switchStackID), -1)
	path = strings.Replace(path, "{staticRouteId}", fmt.Sprintf("%v", staticRouteID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateNetworkSwitchStackRoutingStaticRoute).
		SetResult(&ResponseSwitchUpdateNetworkSwitchStackRoutingStaticRoute{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkSwitchStackRoutingStaticRoute")
	}

	result := response.Result().(*ResponseSwitchUpdateNetworkSwitchStackRoutingStaticRoute)
	return result, response, err

}

//UpdateNetworkSwitchStormControl Update the storm control configuration for a switch network
/* Update the storm control configuration for a switch network

@param networkID networkId path parameter. Network ID
*/
func (s *SwitchService) UpdateNetworkSwitchStormControl(networkID string, requestSwitchUpdateNetworkSwitchStormControl *RequestSwitchUpdateNetworkSwitchStormControl) (*ResponseSwitchUpdateNetworkSwitchStormControl, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/stormControl"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateNetworkSwitchStormControl).
		SetResult(&ResponseSwitchUpdateNetworkSwitchStormControl{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkSwitchStormControl")
	}

	result := response.Result().(*ResponseSwitchUpdateNetworkSwitchStormControl)
	return result, response, err

}

//UpdateNetworkSwitchStp Updates STP settings
/* Updates STP settings

@param networkID networkId path parameter. Network ID
*/
func (s *SwitchService) UpdateNetworkSwitchStp(networkID string, requestSwitchUpdateNetworkSwitchStp *RequestSwitchUpdateNetworkSwitchStp) (*ResponseSwitchUpdateNetworkSwitchStp, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/switch/stp"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSwitchUpdateNetworkSwitchStp).
		SetResult(&ResponseSwitchUpdateNetworkSwitchStp{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkSwitchStp")
	}

	result := response.Result().(*ResponseSwitchUpdateNetworkSwitchStp)
	return result, response, err

}

//UpdateOrganizationConfigTemplateSwitchProfilePort Update a switch template port
/* Update a switch template port

@param organizationID organizationId path parameter. Organization ID
@param configTemplateID configTemplateId path parameter. Config template ID
@param profileID profileId path parameter. Profile ID
@param portID portId path parameter. Port ID
*/
func (s *SwitchService) UpdateOrganizationConfigTemplateSwitchProfilePort(organizationID string, configTemplateID string, profileID string, portID string, requestSwitchUpdateOrganizationConfigTemplateSwitchProfilePort *RequestSwitchUpdateOrganizationConfigTemplateSwitchProfilePort) (*ResponseSwitchUpdateOrganizationConfigTemplateSwitchProfilePort, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/configTemplates/{configTemplateId}/switch/profiles/{profileId}/ports/{portId}"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) DeleteDeviceSwitchRoutingInterface(serial string, interfaceID string) (*resty.Response, error) {
	//serial string,interfaceID string
	path := "/api/v1/devices/{serial}/switch/routing/interfaces/{interfaceId}"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) DeleteDeviceSwitchRoutingStaticRoute(serial string, staticRouteID string) (*resty.Response, error) {
	//serial string,staticRouteID string
	path := "/api/v1/devices/{serial}/switch/routing/staticRoutes/{staticRouteId}"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) DeleteNetworkSwitchAccessPolicy(networkID string, accessPolicyNumber string) (*resty.Response, error) {
	//networkID string,accessPolicyNumber string
	path := "/api/v1/networks/{networkId}/switch/accessPolicies/{accessPolicyNumber}"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(networkID string, trustedServerID string) (*resty.Response, error) {
	//networkID string,trustedServerID string
	path := "/api/v1/networks/{networkId}/switch/dhcpServerPolicy/arpInspection/trustedServers/{trustedServerId}"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) DeleteNetworkSwitchLinkAggregation(networkID string, linkAggregationID string) (*resty.Response, error) {
	//networkID string,linkAggregationID string
	path := "/api/v1/networks/{networkId}/switch/linkAggregations/{linkAggregationId}"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) DeleteNetworkSwitchPortSchedule(networkID string, portScheduleID string) (*resty.Response, error) {
	//networkID string,portScheduleID string
	path := "/api/v1/networks/{networkId}/switch/portSchedules/{portScheduleId}"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) DeleteNetworkSwitchQosRule(networkID string, qosRuleID string) (*resty.Response, error) {
	//networkID string,qosRuleID string
	path := "/api/v1/networks/{networkId}/switch/qosRules/{qosRuleId}"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) DeleteNetworkSwitchRoutingMulticastRendezvousPoint(networkID string, rendezvousPointID string) (*resty.Response, error) {
	//networkID string,rendezvousPointID string
	path := "/api/v1/networks/{networkId}/switch/routing/multicast/rendezvousPoints/{rendezvousPointId}"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) DeleteNetworkSwitchStack(networkID string, switchStackID string) (*resty.Response, error) {
	//networkID string,switchStackID string
	path := "/api/v1/networks/{networkId}/switch/stacks/{switchStackId}"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) DeleteNetworkSwitchStackRoutingInterface(networkID string, switchStackID string, interfaceID string) (*resty.Response, error) {
	//networkID string,switchStackID string,interfaceID string
	path := "/api/v1/networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId}"
	s.rateLimiterBucket.Wait(1)
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


*/
func (s *SwitchService) DeleteNetworkSwitchStackRoutingStaticRoute(networkID string, switchStackID string, staticRouteID string) (*resty.Response, error) {
	//networkID string,switchStackID string,staticRouteID string
	path := "/api/v1/networks/{networkId}/switch/stacks/{switchStackId}/routing/staticRoutes/{staticRouteId}"
	s.rateLimiterBucket.Wait(1)
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
