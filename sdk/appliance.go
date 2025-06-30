package meraki

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type ApplianceService service

type GetDeviceAppliancePerformanceQueryParams struct {
	T0       string  `url:"t0,omitempty"`       //The beginning of the timespan for the data. The maximum lookback period is 30 days from today.
	T1       string  `url:"t1,omitempty"`       //The end of the timespan for the data. t1 can be a maximum of 14 days after t0.
	Timespan float64 `url:"timespan,omitempty"` //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 30 minutes and be less than or equal to 14 days. The default is 30 minutes.
}
type GetNetworkApplianceClientSecurityEventsQueryParams struct {
	T0            string  `url:"t0,omitempty"`            //The beginning of the timespan for the data. Data is gathered after the specified t0 value. The maximum lookback period is 791 days from today.
	T1            string  `url:"t1,omitempty"`            //The end of the timespan for the data. t1 can be a maximum of 791 days after t0.
	Timespan      float64 `url:"timespan,omitempty"`      //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 791 days. The default is 31 days.
	PerPage       int     `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100.
	StartingAfter string  `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string  `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	SortOrder     string  `url:"sortOrder,omitempty"`     //Sorted order of security events based on event detection time. Order options are 'ascending' or 'descending'. Default is ascending order.
}
type GetNetworkApplianceSecurityEventsQueryParams struct {
	T0            string  `url:"t0,omitempty"`            //The beginning of the timespan for the data. Data is gathered after the specified t0 value. The maximum lookback period is 365 days from today.
	T1            string  `url:"t1,omitempty"`            //The end of the timespan for the data. t1 can be a maximum of 365 days after t0.
	Timespan      float64 `url:"timespan,omitempty"`      //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 365 days. The default is 31 days.
	PerPage       int     `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100.
	StartingAfter string  `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string  `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	SortOrder     string  `url:"sortOrder,omitempty"`     //Sorted order of security events based on event detection time. Order options are 'ascending' or 'descending'. Default is ascending order.
}
type GetNetworkApplianceUplinksUsageHistoryQueryParams struct {
	T0         string  `url:"t0,omitempty"`         //The beginning of the timespan for the data. The maximum lookback period is 365 days from today.
	T1         string  `url:"t1,omitempty"`         //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan   float64 `url:"timespan,omitempty"`   //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 10 minutes.
	Resolution int     `url:"resolution,omitempty"` //The time resolution in seconds for returned data. The valid resolutions are: 60, 300, 600, 1800, 3600, 86400. The default is 60.
}
type GetOrganizationApplianceDNSLocalProfilesQueryParams struct {
	ProfileIDs []string `url:"profileIds[],omitempty"` //Optional parameter to filter the results by profile IDs
}
type GetOrganizationApplianceDNSLocalProfilesAssignmentsQueryParams struct {
	ProfileIDs []string `url:"profileIds[],omitempty"` //Optional parameter to filter the results by profile IDs
	NetworkIDs []string `url:"networkIds[],omitempty"` //Optional parameter to filter the results by network IDs
}
type GetOrganizationApplianceDNSLocalRecordsQueryParams struct {
	ProfileIDs []string `url:"profileIds[],omitempty"` //Optional parameter to filter the results by profile IDs
}
type GetOrganizationApplianceDNSSplitProfilesQueryParams struct {
	ProfileIDs []string `url:"profileIds[],omitempty"` //Optional parameter to filter the results by profile IDs
}
type GetOrganizationApplianceDNSSplitProfilesAssignmentsQueryParams struct {
	ProfileIDs []string `url:"profileIds[],omitempty"` //Optional parameter to filter the results by profile IDs
	NetworkIDs []string `url:"networkIds[],omitempty"` //Optional parameter to filter the results by network IDs
}
type GetOrganizationApplianceFirewallMulticastForwardingByNetworkQueryParams struct {
	PerPage       int      `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter string   `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string   `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	NetworkIDs    []string `url:"networkIds[],omitempty"`  //Optional parameter to filter the results by network IDs
}
type GetOrganizationApplianceSecurityEventsQueryParams struct {
	T0            string  `url:"t0,omitempty"`            //The beginning of the timespan for the data. Data is gathered after the specified t0 value. The maximum lookback period is 365 days from today.
	T1            string  `url:"t1,omitempty"`            //The end of the timespan for the data. t1 can be a maximum of 365 days after t0.
	Timespan      float64 `url:"timespan,omitempty"`      //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 365 days. The default is 31 days.
	PerPage       int     `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100.
	StartingAfter string  `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string  `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	SortOrder     string  `url:"sortOrder,omitempty"`     //Sorted order of security events based on event detection time. Order options are 'ascending' or 'descending'. Default is ascending order.
}
type GetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkQueryParams struct {
	PerPage       int      `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50.
	StartingAfter string   `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string   `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	NetworkIDs    []string `url:"networkIds[],omitempty"`  //Optional parameter to filter the results by network IDs
}
type GetOrganizationApplianceUplinkStatusesQueryParams struct {
	PerPage       int      `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter string   `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string   `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	NetworkIDs    []string `url:"networkIds[],omitempty"`  //A list of network IDs. The returned devices will be filtered to only include these networks.
	Serials       []string `url:"serials[],omitempty"`     //A list of serial numbers. The returned devices will be filtered to only include these serials.
	Iccids        []string `url:"iccids[],omitempty"`      //A list of ICCIDs. The returned devices will be filtered to only include these ICCIDs.
}
type GetOrganizationApplianceUplinksUsageByNetworkQueryParams struct {
	T0       string  `url:"t0,omitempty"`       //The beginning of the timespan for the data. The maximum lookback period is 365 days from today.
	T1       string  `url:"t1,omitempty"`       //The end of the timespan for the data. t1 can be a maximum of 14 days after t0.
	Timespan float64 `url:"timespan,omitempty"` //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 14 days. The default is 1 day.
}
type GetOrganizationApplianceVpnStatsQueryParams struct {
	PerPage       int      `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 300. Default is 300.
	StartingAfter string   `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string   `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	NetworkIDs    []string `url:"networkIds[],omitempty"`  //A list of Meraki network IDs to filter results to contain only specified networks. E.g.: networkIds[]=N_12345678&networkIds[]=L_3456
	T0            string   `url:"t0,omitempty"`            //The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	T1            string   `url:"t1,omitempty"`            //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan      float64  `url:"timespan,omitempty"`      //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
}
type GetOrganizationApplianceVpnStatusesQueryParams struct {
	PerPage       int      `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 300. Default is 300.
	StartingAfter string   `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string   `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	NetworkIDs    []string `url:"networkIds[],omitempty"`  //A list of Meraki network IDs to filter results to contain only specified networks. E.g.: networkIds[]=N_12345678&networkIds[]=L_3456
}

type ResponseApplianceGetDeviceApplianceDhcpSubnets []ResponseItemApplianceGetDeviceApplianceDhcpSubnets // Array of ResponseApplianceGetDeviceApplianceDhcpSubnets
type ResponseItemApplianceGetDeviceApplianceDhcpSubnets struct {
	FreeCount *int   `json:"freeCount,omitempty"` // Count of free IP addresses in subnet
	Subnet    string `json:"subnet,omitempty"`    // Subnet
	UsedCount *int   `json:"usedCount,omitempty"` // Count of used IP addresses in subnet
	VLANID    *int   `json:"vlanId,omitempty"`    // VLAN ID
}
type ResponseApplianceGetDeviceAppliancePerformance struct {
	PerfScore *float64 `json:"perfScore,omitempty"` // Utilization for the MX device
}
type ResponseApplianceGetDeviceAppliancePrefixesDelegated []ResponseItemApplianceGetDeviceAppliancePrefixesDelegated // Array of ResponseApplianceGetDeviceAppliancePrefixesDelegated
type ResponseItemApplianceGetDeviceAppliancePrefixesDelegated struct {
	Counts      *ResponseItemApplianceGetDeviceAppliancePrefixesDelegatedCounts `json:"counts,omitempty"`      //
	Description string                                                          `json:"description,omitempty"` //
	ExpiresAt   string                                                          `json:"expiresAt,omitempty"`   //
	IsPreferred *bool                                                           `json:"isPreferred,omitempty"` //
	Method      string                                                          `json:"method,omitempty"`      //
	Origin      *ResponseItemApplianceGetDeviceAppliancePrefixesDelegatedOrigin `json:"origin,omitempty"`      //
	Prefix      string                                                          `json:"prefix,omitempty"`      //
}
type ResponseItemApplianceGetDeviceAppliancePrefixesDelegatedCounts struct {
	Assigned  *int `json:"assigned,omitempty"`  //
	Available *int `json:"available,omitempty"` //
}
type ResponseItemApplianceGetDeviceAppliancePrefixesDelegatedOrigin struct {
	Interface string `json:"interface,omitempty"` //
}
type ResponseApplianceGetDeviceAppliancePrefixesDelegatedVLANAssignments []ResponseItemApplianceGetDeviceAppliancePrefixesDelegatedVLANAssignments // Array of ResponseApplianceGetDeviceAppliancePrefixesDelegatedVlanAssignments
type ResponseItemApplianceGetDeviceAppliancePrefixesDelegatedVLANAssignments struct {
	IPv6   *ResponseItemApplianceGetDeviceAppliancePrefixesDelegatedVLANAssignmentsIPv6   `json:"ipv6,omitempty"`   //
	Origin *ResponseItemApplianceGetDeviceAppliancePrefixesDelegatedVLANAssignmentsOrigin `json:"origin,omitempty"` //
	Status string                                                                         `json:"status,omitempty"` //
	VLAN   *ResponseItemApplianceGetDeviceAppliancePrefixesDelegatedVLANAssignmentsVLAN   `json:"vlan,omitempty"`   //
}
type ResponseItemApplianceGetDeviceAppliancePrefixesDelegatedVLANAssignmentsIPv6 struct {
	Address                string                                                                                             `json:"address,omitempty"`                //
	LinkLocal              *ResponseItemApplianceGetDeviceAppliancePrefixesDelegatedVLANAssignmentsIPv6LinkLocal              `json:"linkLocal,omitempty"`              //
	Prefix                 string                                                                                             `json:"prefix,omitempty"`                 //
	SolicitedNodeMulticast *ResponseItemApplianceGetDeviceAppliancePrefixesDelegatedVLANAssignmentsIPv6SolicitedNodeMulticast `json:"solicitedNodeMulticast,omitempty"` //
}
type ResponseItemApplianceGetDeviceAppliancePrefixesDelegatedVLANAssignmentsIPv6LinkLocal struct {
	Address string `json:"address,omitempty"` //
}
type ResponseItemApplianceGetDeviceAppliancePrefixesDelegatedVLANAssignmentsIPv6SolicitedNodeMulticast struct {
	Address string `json:"address,omitempty"` //
}
type ResponseItemApplianceGetDeviceAppliancePrefixesDelegatedVLANAssignmentsOrigin struct {
	Interface string `json:"interface,omitempty"` //
	Prefix    string `json:"prefix,omitempty"`    //
}
type ResponseItemApplianceGetDeviceAppliancePrefixesDelegatedVLANAssignmentsVLAN struct {
	ID   *int   `json:"id,omitempty"`   //
	Name string `json:"name,omitempty"` //
}
type ResponseApplianceGetDeviceApplianceRadioSettings struct {
	FiveGhzSettings    *ResponseApplianceGetDeviceApplianceRadioSettingsFiveGhzSettings    `json:"fiveGhzSettings,omitempty"`    // Manual radio settings for 5 GHz
	RfProfileID        string                                                              `json:"rfProfileId,omitempty"`        // RF Profile ID
	Serial             string                                                              `json:"serial,omitempty"`             // The device serial
	TwoFourGhzSettings *ResponseApplianceGetDeviceApplianceRadioSettingsTwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"` // Manual radio settings for 2.4 GHz
}
type ResponseApplianceGetDeviceApplianceRadioSettingsFiveGhzSettings struct {
	Channel      *int `json:"channel,omitempty"`      // Manual channel for 5 GHz
	ChannelWidth *int `json:"channelWidth,omitempty"` // Manual channel width for 5 GHz
	TargetPower  *int `json:"targetPower,omitempty"`  // Manual target power for 5 GHz
}
type ResponseApplianceGetDeviceApplianceRadioSettingsTwoFourGhzSettings struct {
	Channel     *int `json:"channel,omitempty"`     // Manual channel for 2.4 GHz
	TargetPower *int `json:"targetPower,omitempty"` // Manual target power for 2.4 GHz
}
type ResponseApplianceUpdateDeviceApplianceRadioSettings struct {
	FiveGhzSettings    *ResponseApplianceUpdateDeviceApplianceRadioSettingsFiveGhzSettings    `json:"fiveGhzSettings,omitempty"`    // Manual radio settings for 5 GHz
	RfProfileID        string                                                                 `json:"rfProfileId,omitempty"`        // RF Profile ID
	Serial             string                                                                 `json:"serial,omitempty"`             // The device serial
	TwoFourGhzSettings *ResponseApplianceUpdateDeviceApplianceRadioSettingsTwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"` // Manual radio settings for 2.4 GHz
}
type ResponseApplianceUpdateDeviceApplianceRadioSettingsFiveGhzSettings struct {
	Channel      *int `json:"channel,omitempty"`      // Manual channel for 5 GHz
	ChannelWidth *int `json:"channelWidth,omitempty"` // Manual channel width for 5 GHz
	TargetPower  *int `json:"targetPower,omitempty"`  // Manual target power for 5 GHz
}
type ResponseApplianceUpdateDeviceApplianceRadioSettingsTwoFourGhzSettings struct {
	Channel     *int `json:"channel,omitempty"`     // Manual channel for 2.4 GHz
	TargetPower *int `json:"targetPower,omitempty"` // Manual target power for 2.4 GHz
}
type ResponseApplianceGetDeviceApplianceUplinksSettings struct {
	Interfaces *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfaces `json:"interfaces,omitempty"` // Interface settings.
}
type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfaces struct {
	Wan1 *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1 `json:"wan1,omitempty"` // WAN 1 settings.
	Wan2 *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2 `json:"wan2,omitempty"` // WAN 2 settings.
}
type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1 struct {
	Enabled     *bool                                                                        `json:"enabled,omitempty"`     // Enable or disable the interface.
	Pppoe       *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1Pppoe       `json:"pppoe,omitempty"`       // Configuration options for PPPoE.
	Svis        *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1Svis        `json:"svis,omitempty"`        // SVI settings by protocol.
	VLANTagging *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1VLANTagging `json:"vlanTagging,omitempty"` // VLAN tagging settings.
}
type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1Pppoe struct {
	Authentication *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1PppoeAuthentication `json:"authentication,omitempty"` // Settings for PPPoE Authentication.
	Enabled        *bool                                                                                `json:"enabled,omitempty"`        // Whether PPPoE is enabled.
}
type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1PppoeAuthentication struct {
	Enabled  *bool  `json:"enabled,omitempty"`  // Whether PPPoE authentication is enabled.
	Username string `json:"username,omitempty"` // Username for PPPoE authentication.
}
type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1Svis struct {
	IPv4 *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv4 `json:"ipv4,omitempty"` // IPv4 settings for static/dynamic mode.
	IPv6 *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv6 `json:"ipv6,omitempty"` // IPv6 settings for static/dynamic mode.
}
type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv4 struct {
	Address        string                                                                               `json:"address,omitempty"`        // IP address and subnet mask when in static mode.
	AssignmentMode string                                                                               `json:"assignmentMode,omitempty"` // The assignment mode for this SVI. Applies only when PPPoE is disabled.
	Gateway        string                                                                               `json:"gateway,omitempty"`        // Gateway IP address when in static mode.
	Nameservers    *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv4Nameservers `json:"nameservers,omitempty"`    // The nameserver settings for this SVI.
}
type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv4Nameservers struct {
	Addresses []string `json:"addresses,omitempty"` // Up to 2 nameserver addresses to use, ordered in priority from highest to lowest priority.
}
type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv6 struct {
	Address        string                                                                               `json:"address,omitempty"`        // Static address that will override the one(s) received by SLAAC.
	AssignmentMode string                                                                               `json:"assignmentMode,omitempty"` // The assignment mode for this SVI. Applies only when PPPoE is disabled.
	Gateway        string                                                                               `json:"gateway,omitempty"`        // Static gateway that will override the one received by autoconf.
	Nameservers    *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv6Nameservers `json:"nameservers,omitempty"`    // The nameserver settings for this SVI.
}
type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv6Nameservers struct {
	Addresses []string `json:"addresses,omitempty"` // Up to 2 nameserver addresses to use, ordered in priority from highest to lowest priority.
}
type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1VLANTagging struct {
	Enabled *bool `json:"enabled,omitempty"` // Whether VLAN tagging is enabled.
	VLANID  *int  `json:"vlanId,omitempty"`  // The ID of the VLAN to use for VLAN tagging.
}
type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2 struct {
	Enabled     *bool                                                                        `json:"enabled,omitempty"`     // Enable or disable the interface.
	Pppoe       *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2Pppoe       `json:"pppoe,omitempty"`       // Configuration options for PPPoE.
	Svis        *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2Svis        `json:"svis,omitempty"`        // SVI settings by protocol.
	VLANTagging *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2VLANTagging `json:"vlanTagging,omitempty"` // VLAN tagging settings.
}
type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2Pppoe struct {
	Authentication *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2PppoeAuthentication `json:"authentication,omitempty"` // Settings for PPPoE Authentication.
	Enabled        *bool                                                                                `json:"enabled,omitempty"`        // Whether PPPoE is enabled.
}
type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2PppoeAuthentication struct {
	Enabled  *bool  `json:"enabled,omitempty"`  // Whether PPPoE authentication is enabled.
	Username string `json:"username,omitempty"` // Username for PPPoE authentication.
}
type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2Svis struct {
	IPv4 *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv4 `json:"ipv4,omitempty"` // IPv4 settings for static/dynamic mode.
	IPv6 *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv6 `json:"ipv6,omitempty"` // IPv6 settings for static/dynamic mode.
}
type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv4 struct {
	Address        string                                                                               `json:"address,omitempty"`        // IP address and subnet mask when in static mode.
	AssignmentMode string                                                                               `json:"assignmentMode,omitempty"` // The assignment mode for this SVI. Applies only when PPPoE is disabled.
	Gateway        string                                                                               `json:"gateway,omitempty"`        // Gateway IP address when in static mode.
	Nameservers    *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv4Nameservers `json:"nameservers,omitempty"`    // The nameserver settings for this SVI.
}
type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv4Nameservers struct {
	Addresses []string `json:"addresses,omitempty"` // Up to 2 nameserver addresses to use, ordered in priority from highest to lowest priority.
}
type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv6 struct {
	Address        string                                                                               `json:"address,omitempty"`        // Static address that will override the one(s) received by SLAAC.
	AssignmentMode string                                                                               `json:"assignmentMode,omitempty"` // The assignment mode for this SVI. Applies only when PPPoE is disabled.
	Gateway        string                                                                               `json:"gateway,omitempty"`        // Static gateway that will override the one received by autoconf.
	Nameservers    *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv6Nameservers `json:"nameservers,omitempty"`    // The nameserver settings for this SVI.
}
type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv6Nameservers struct {
	Addresses []string `json:"addresses,omitempty"` // Up to 2 nameserver addresses to use, ordered in priority from highest to lowest priority.
}
type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2VLANTagging struct {
	Enabled *bool `json:"enabled,omitempty"` // Whether VLAN tagging is enabled.
	VLANID  *int  `json:"vlanId,omitempty"`  // The ID of the VLAN to use for VLAN tagging.
}
type ResponseApplianceUpdateDeviceApplianceUplinksSettings struct {
	Interfaces *ResponseApplianceUpdateDeviceApplianceUplinksSettingsInterfaces `json:"interfaces,omitempty"` // Interface settings.
}
type ResponseApplianceUpdateDeviceApplianceUplinksSettingsInterfaces struct {
	Wan1 *ResponseApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1 `json:"wan1,omitempty"` // WAN 1 settings.
	Wan2 *ResponseApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2 `json:"wan2,omitempty"` // WAN 2 settings.
}
type ResponseApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1 struct {
	Enabled     *bool                                                                           `json:"enabled,omitempty"`     // Enable or disable the interface.
	Pppoe       *ResponseApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1Pppoe       `json:"pppoe,omitempty"`       // Configuration options for PPPoE.
	Svis        *ResponseApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1Svis        `json:"svis,omitempty"`        // SVI settings by protocol.
	VLANTagging *ResponseApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1VLANTagging `json:"vlanTagging,omitempty"` // VLAN tagging settings.
}
type ResponseApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1Pppoe struct {
	Authentication *ResponseApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1PppoeAuthentication `json:"authentication,omitempty"` // Settings for PPPoE Authentication.
	Enabled        *bool                                                                                   `json:"enabled,omitempty"`        // Whether PPPoE is enabled.
}
type ResponseApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1PppoeAuthentication struct {
	Enabled  *bool  `json:"enabled,omitempty"`  // Whether PPPoE authentication is enabled.
	Username string `json:"username,omitempty"` // Username for PPPoE authentication.
}
type ResponseApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1Svis struct {
	IPv4 *ResponseApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv4 `json:"ipv4,omitempty"` // IPv4 settings for static/dynamic mode.
	IPv6 *ResponseApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv6 `json:"ipv6,omitempty"` // IPv6 settings for static/dynamic mode.
}
type ResponseApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv4 struct {
	Address        string                                                                                  `json:"address,omitempty"`        // IP address and subnet mask when in static mode.
	AssignmentMode string                                                                                  `json:"assignmentMode,omitempty"` // The assignment mode for this SVI. Applies only when PPPoE is disabled.
	Gateway        string                                                                                  `json:"gateway,omitempty"`        // Gateway IP address when in static mode.
	Nameservers    *ResponseApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv4Nameservers `json:"nameservers,omitempty"`    // The nameserver settings for this SVI.
}
type ResponseApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv4Nameservers struct {
	Addresses []string `json:"addresses,omitempty"` // Up to 2 nameserver addresses to use, ordered in priority from highest to lowest priority.
}
type ResponseApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv6 struct {
	Address        string                                                                                  `json:"address,omitempty"`        // Static address that will override the one(s) received by SLAAC.
	AssignmentMode string                                                                                  `json:"assignmentMode,omitempty"` // The assignment mode for this SVI. Applies only when PPPoE is disabled.
	Gateway        string                                                                                  `json:"gateway,omitempty"`        // Static gateway that will override the one received by autoconf.
	Nameservers    *ResponseApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv6Nameservers `json:"nameservers,omitempty"`    // The nameserver settings for this SVI.
}
type ResponseApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv6Nameservers struct {
	Addresses []string `json:"addresses,omitempty"` // Up to 2 nameserver addresses to use, ordered in priority from highest to lowest priority.
}
type ResponseApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1VLANTagging struct {
	Enabled *bool `json:"enabled,omitempty"` // Whether VLAN tagging is enabled.
	VLANID  *int  `json:"vlanId,omitempty"`  // The ID of the VLAN to use for VLAN tagging.
}
type ResponseApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2 struct {
	Enabled     *bool                                                                           `json:"enabled,omitempty"`     // Enable or disable the interface.
	Pppoe       *ResponseApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2Pppoe       `json:"pppoe,omitempty"`       // Configuration options for PPPoE.
	Svis        *ResponseApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2Svis        `json:"svis,omitempty"`        // SVI settings by protocol.
	VLANTagging *ResponseApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2VLANTagging `json:"vlanTagging,omitempty"` // VLAN tagging settings.
}
type ResponseApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2Pppoe struct {
	Authentication *ResponseApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2PppoeAuthentication `json:"authentication,omitempty"` // Settings for PPPoE Authentication.
	Enabled        *bool                                                                                   `json:"enabled,omitempty"`        // Whether PPPoE is enabled.
}
type ResponseApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2PppoeAuthentication struct {
	Enabled  *bool  `json:"enabled,omitempty"`  // Whether PPPoE authentication is enabled.
	Username string `json:"username,omitempty"` // Username for PPPoE authentication.
}
type ResponseApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2Svis struct {
	IPv4 *ResponseApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv4 `json:"ipv4,omitempty"` // IPv4 settings for static/dynamic mode.
	IPv6 *ResponseApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv6 `json:"ipv6,omitempty"` // IPv6 settings for static/dynamic mode.
}
type ResponseApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv4 struct {
	Address        string                                                                                  `json:"address,omitempty"`        // IP address and subnet mask when in static mode.
	AssignmentMode string                                                                                  `json:"assignmentMode,omitempty"` // The assignment mode for this SVI. Applies only when PPPoE is disabled.
	Gateway        string                                                                                  `json:"gateway,omitempty"`        // Gateway IP address when in static mode.
	Nameservers    *ResponseApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv4Nameservers `json:"nameservers,omitempty"`    // The nameserver settings for this SVI.
}
type ResponseApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv4Nameservers struct {
	Addresses []string `json:"addresses,omitempty"` // Up to 2 nameserver addresses to use, ordered in priority from highest to lowest priority.
}
type ResponseApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv6 struct {
	Address        string                                                                                  `json:"address,omitempty"`        // Static address that will override the one(s) received by SLAAC.
	AssignmentMode string                                                                                  `json:"assignmentMode,omitempty"` // The assignment mode for this SVI. Applies only when PPPoE is disabled.
	Gateway        string                                                                                  `json:"gateway,omitempty"`        // Static gateway that will override the one received by autoconf.
	Nameservers    *ResponseApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv6Nameservers `json:"nameservers,omitempty"`    // The nameserver settings for this SVI.
}
type ResponseApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv6Nameservers struct {
	Addresses []string `json:"addresses,omitempty"` // Up to 2 nameserver addresses to use, ordered in priority from highest to lowest priority.
}
type ResponseApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2VLANTagging struct {
	Enabled *bool `json:"enabled,omitempty"` // Whether VLAN tagging is enabled.
	VLANID  *int  `json:"vlanId,omitempty"`  // The ID of the VLAN to use for VLAN tagging.
}
type ResponseApplianceCreateDeviceApplianceVmxAuthenticationToken struct {
	ExpiresAt string `json:"expiresAt,omitempty"` // The expiration time for the token, in ISO 8601 format
	Token     string `json:"token,omitempty"`     // The newly generated authentication token for the vMX instance
}
type ResponseApplianceGetNetworkApplianceClientSecurityEvents []ResponseItemApplianceGetNetworkApplianceClientSecurityEvents // Array of ResponseApplianceGetNetworkApplianceClientSecurityEvents
type ResponseItemApplianceGetNetworkApplianceClientSecurityEvents struct {
	Action          string `json:"action,omitempty"`          //
	CanonicalName   string `json:"canonicalName,omitempty"`   //
	ClientIP        string `json:"clientIp,omitempty"`        //
	ClientMac       string `json:"clientMac,omitempty"`       //
	ClientName      string `json:"clientName,omitempty"`      //
	DestIP          string `json:"destIp,omitempty"`          //
	DestinationPort *int   `json:"destinationPort,omitempty"` //
	Disposition     string `json:"disposition,omitempty"`     //
	EventType       string `json:"eventType,omitempty"`       //
	FileHash        string `json:"fileHash,omitempty"`        //
	FileSizeBytes   *int   `json:"fileSizeBytes,omitempty"`   //
	FileType        string `json:"fileType,omitempty"`        //
	Protocol        string `json:"protocol,omitempty"`        //
	SrcIP           string `json:"srcIp,omitempty"`           //
	Ts              string `json:"ts,omitempty"`              //
	URI             string `json:"uri,omitempty"`             //
}
type ResponseApplianceGetNetworkApplianceConnectivityMonitoringDestinations struct {
	Destinations *[]ResponseApplianceGetNetworkApplianceConnectivityMonitoringDestinationsDestinations `json:"destinations,omitempty"` // The list of connectivity monitoring destinations
}
type ResponseApplianceGetNetworkApplianceConnectivityMonitoringDestinationsDestinations struct {
	Default     *bool  `json:"default,omitempty"`     // Boolean indicating whether this is the default testing destination (true) or not (false). Defaults to false. Only one default is allowed
	Description string `json:"description,omitempty"` // Description of the testing destination. Optional, defaults to an empty string
	IP          string `json:"ip,omitempty"`          // The IP address to test connectivity with
}
type ResponseApplianceUpdateNetworkApplianceConnectivityMonitoringDestinations struct {
	Destinations *[]ResponseApplianceUpdateNetworkApplianceConnectivityMonitoringDestinationsDestinations `json:"destinations,omitempty"` // The list of connectivity monitoring destinations
}
type ResponseApplianceUpdateNetworkApplianceConnectivityMonitoringDestinationsDestinations struct {
	Default     *bool  `json:"default,omitempty"`     // Boolean indicating whether this is the default testing destination (true) or not (false). Defaults to false. Only one default is allowed
	Description string `json:"description,omitempty"` // Description of the testing destination. Optional, defaults to an empty string
	IP          string `json:"ip,omitempty"`          // The IP address to test connectivity with
}
type ResponseApplianceGetNetworkApplianceContentFiltering struct {
	AllowedURLPatterns   []string                                                                    `json:"allowedUrlPatterns,omitempty"`   //
	BlockedURLCategories *[]ResponseApplianceGetNetworkApplianceContentFilteringBlockedURLCategories `json:"blockedUrlCategories,omitempty"` //
	BlockedURLPatterns   []string                                                                    `json:"blockedUrlPatterns,omitempty"`   //
	URLCategoryListSize  string                                                                      `json:"urlCategoryListSize,omitempty"`  //
}
type ResponseApplianceGetNetworkApplianceContentFilteringBlockedURLCategories struct {
	ID   string `json:"id,omitempty"`   //
	Name string `json:"name,omitempty"` //
}
type ResponseApplianceUpdateNetworkApplianceContentFiltering interface{}
type ResponseApplianceGetNetworkApplianceContentFilteringCategories struct {
	Categories *[]ResponseApplianceGetNetworkApplianceContentFilteringCategoriesCategories `json:"categories,omitempty"` //
}
type ResponseApplianceGetNetworkApplianceContentFilteringCategoriesCategories struct {
	ID   string `json:"id,omitempty"`   //
	Name string `json:"name,omitempty"` //
}
type ResponseApplianceGetNetworkApplianceFirewallCellularFirewallRules struct {
	Rules *[]ResponseApplianceGetNetworkApplianceFirewallCellularFirewallRulesRules `json:"rules,omitempty"` //
}
type ResponseApplianceGetNetworkApplianceFirewallCellularFirewallRulesRules struct {
	Comment       string `json:"comment,omitempty"`       //
	DestCidr      string `json:"destCidr,omitempty"`      //
	DestPort      string `json:"destPort,omitempty"`      //
	Policy        string `json:"policy,omitempty"`        //
	Protocol      string `json:"protocol,omitempty"`      //
	SrcCidr       string `json:"srcCidr,omitempty"`       //
	SrcPort       string `json:"srcPort,omitempty"`       //
	SyslogEnabled *bool  `json:"syslogEnabled,omitempty"` //
}
type ResponseApplianceUpdateNetworkApplianceFirewallCellularFirewallRules interface{}
type ResponseApplianceGetNetworkApplianceFirewallFirewalledServices []ResponseItemApplianceGetNetworkApplianceFirewallFirewalledServices // Array of ResponseApplianceGetNetworkApplianceFirewallFirewalledServices
type ResponseItemApplianceGetNetworkApplianceFirewallFirewalledServices struct {
	Access     string   `json:"access,omitempty"`     // A string indicating the rule for which IPs are allowed to use the specified service
	AllowedIPs []string `json:"allowedIps,omitempty"` // An array of allowed IPs that can access the service
	Service    string   `json:"service,omitempty"`    // Appliance service name
}
type ResponseApplianceGetNetworkApplianceFirewallFirewalledService struct {
	Access     string   `json:"access,omitempty"`     // A string indicating the rule for which IPs are allowed to use the specified service
	AllowedIPs []string `json:"allowedIps,omitempty"` // An array of allowed IPs that can access the service
	Service    string   `json:"service,omitempty"`    // Appliance service name
}
type ResponseApplianceUpdateNetworkApplianceFirewallFirewalledService struct {
	Access     string   `json:"access,omitempty"`     // A string indicating the rule for which IPs are allowed to use the specified service
	AllowedIPs []string `json:"allowedIps,omitempty"` // An array of allowed IPs that can access the service
	Service    string   `json:"service,omitempty"`    // Appliance service name
}
type ResponseApplianceGetNetworkApplianceFirewallInboundCellularFirewallRules struct {
	Rules *[]ResponseApplianceGetNetworkApplianceFirewallInboundCellularFirewallRulesRules `json:"rules,omitempty"` // An ordered array of the firewall rules (not including the default rule)
}
type ResponseApplianceGetNetworkApplianceFirewallInboundCellularFirewallRulesRules struct {
	Comment       string `json:"comment,omitempty"`       // Description of the rule (optional)
	DestCidr      string `json:"destCidr,omitempty"`      // Comma-separated list of destination IP address(es) (in IP or CIDR notation), fully-qualified domain names (FQDN) or 'any'
	DestPort      string `json:"destPort,omitempty"`      // Comma-separated list of destination port(s) (integer in the range 1-65535), or 'any'
	Policy        string `json:"policy,omitempty"`        // 'allow' or 'deny' traffic specified by this rule
	Protocol      string `json:"protocol,omitempty"`      // The type of protocol (must be 'tcp', 'udp', 'icmp', 'icmp6' or 'any')
	SrcCidr       string `json:"srcCidr,omitempty"`       // Comma-separated list of source IP address(es) (in IP or CIDR notation), or 'any' (note: FQDN not supported for source addresses)
	SrcPort       string `json:"srcPort,omitempty"`       // Comma-separated list of source port(s) (integer in the range 1-65535), or 'any'
	SyslogEnabled *bool  `json:"syslogEnabled,omitempty"` // Log this rule to syslog (true or false, boolean value) - only applicable if a syslog has been configured (optional)
}
type ResponseApplianceUpdateNetworkApplianceFirewallInboundCellularFirewallRules struct {
	Rules *[]ResponseApplianceUpdateNetworkApplianceFirewallInboundCellularFirewallRulesRules `json:"rules,omitempty"` // An ordered array of the firewall rules (not including the default rule)
}
type ResponseApplianceUpdateNetworkApplianceFirewallInboundCellularFirewallRulesRules struct {
	Comment       string `json:"comment,omitempty"`       // Description of the rule (optional)
	DestCidr      string `json:"destCidr,omitempty"`      // Comma-separated list of destination IP address(es) (in IP or CIDR notation), fully-qualified domain names (FQDN) or 'any'
	DestPort      string `json:"destPort,omitempty"`      // Comma-separated list of destination port(s) (integer in the range 1-65535), or 'any'
	Policy        string `json:"policy,omitempty"`        // 'allow' or 'deny' traffic specified by this rule
	Protocol      string `json:"protocol,omitempty"`      // The type of protocol (must be 'tcp', 'udp', 'icmp', 'icmp6' or 'any')
	SrcCidr       string `json:"srcCidr,omitempty"`       // Comma-separated list of source IP address(es) (in IP or CIDR notation), or 'any' (note: FQDN not supported for source addresses)
	SrcPort       string `json:"srcPort,omitempty"`       // Comma-separated list of source port(s) (integer in the range 1-65535), or 'any'
	SyslogEnabled *bool  `json:"syslogEnabled,omitempty"` // Log this rule to syslog (true or false, boolean value) - only applicable if a syslog has been configured (optional)
}
type ResponseApplianceGetNetworkApplianceFirewallInboundFirewallRules struct {
	Rules             *[]ResponseApplianceGetNetworkApplianceFirewallInboundFirewallRulesRules `json:"rules,omitempty"`             // An ordered array of the firewall rules (not including the default rule)
	SyslogDefaultRule *bool                                                                    `json:"syslogDefaultRule,omitempty"` // Log the special default rule (boolean value - enable only if you've configured a syslog server) (optional)
}
type ResponseApplianceGetNetworkApplianceFirewallInboundFirewallRulesRules struct {
	Comment       string `json:"comment,omitempty"`       // Description of the rule (optional)
	DestCidr      string `json:"destCidr,omitempty"`      // Comma-separated list of destination IP address(es) (in IP or CIDR notation), fully-qualified domain names (FQDN) or 'any'
	DestPort      string `json:"destPort,omitempty"`      // Comma-separated list of destination port(s) (integer in the range 1-65535), or 'any'
	Policy        string `json:"policy,omitempty"`        // 'allow' or 'deny' traffic specified by this rule
	Protocol      string `json:"protocol,omitempty"`      // The type of protocol (must be 'tcp', 'udp', 'icmp', 'icmp6' or 'any')
	SrcCidr       string `json:"srcCidr,omitempty"`       // Comma-separated list of source IP address(es) (in IP or CIDR notation), or 'any' (note: FQDN not supported for source addresses)
	SrcPort       string `json:"srcPort,omitempty"`       // Comma-separated list of source port(s) (integer in the range 1-65535), or 'any'
	SyslogEnabled *bool  `json:"syslogEnabled,omitempty"` // Log this rule to syslog (true or false, boolean value) - only applicable if a syslog has been configured (optional)
}
type ResponseApplianceUpdateNetworkApplianceFirewallInboundFirewallRules struct {
	Rules             *[]ResponseApplianceUpdateNetworkApplianceFirewallInboundFirewallRulesRules `json:"rules,omitempty"`             // An ordered array of the firewall rules (not including the default rule)
	SyslogDefaultRule *bool                                                                       `json:"syslogDefaultRule,omitempty"` // Log the special default rule (boolean value - enable only if you've configured a syslog server) (optional)
}
type ResponseApplianceUpdateNetworkApplianceFirewallInboundFirewallRulesRules struct {
	Comment       string `json:"comment,omitempty"`       // Description of the rule (optional)
	DestCidr      string `json:"destCidr,omitempty"`      // Comma-separated list of destination IP address(es) (in IP or CIDR notation), fully-qualified domain names (FQDN) or 'any'
	DestPort      string `json:"destPort,omitempty"`      // Comma-separated list of destination port(s) (integer in the range 1-65535), or 'any'
	Policy        string `json:"policy,omitempty"`        // 'allow' or 'deny' traffic specified by this rule
	Protocol      string `json:"protocol,omitempty"`      // The type of protocol (must be 'tcp', 'udp', 'icmp', 'icmp6' or 'any')
	SrcCidr       string `json:"srcCidr,omitempty"`       // Comma-separated list of source IP address(es) (in IP or CIDR notation), or 'any' (note: FQDN not supported for source addresses)
	SrcPort       string `json:"srcPort,omitempty"`       // Comma-separated list of source port(s) (integer in the range 1-65535), or 'any'
	SyslogEnabled *bool  `json:"syslogEnabled,omitempty"` // Log this rule to syslog (true or false, boolean value) - only applicable if a syslog has been configured (optional)
}
type ResponseApplianceGetNetworkApplianceFirewallL3FirewallRules struct {
	Rules *[]ResponseApplianceGetNetworkApplianceFirewallL3FirewallRulesRules `json:"rules,omitempty"` //
}
type ResponseApplianceGetNetworkApplianceFirewallL3FirewallRulesRules struct {
	Comment       string `json:"comment,omitempty"`       //
	DestCidr      string `json:"destCidr,omitempty"`      //
	DestPort      string `json:"destPort,omitempty"`      //
	Policy        string `json:"policy,omitempty"`        //
	Protocol      string `json:"protocol,omitempty"`      //
	SrcCidr       string `json:"srcCidr,omitempty"`       //
	SrcPort       string `json:"srcPort,omitempty"`       //
	SyslogEnabled *bool  `json:"syslogEnabled,omitempty"` //
}
type ResponseApplianceUpdateNetworkApplianceFirewallL3FirewallRules interface{}
type ResponseApplianceGetNetworkApplianceFirewallL7FirewallRules struct {
	Rules *[]ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesRules `json:"rules,omitempty"` //
}
type ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesRules struct {
	Policy    string                                                                    `json:"policy,omitempty"` //
	Type      string                                                                    `json:"type,omitempty"`   //
	Value     *string                                                                   //
	ValueObj  *ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesRulesValueObj //
	ValueList *[]string                                                                 //
}
type ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesRulesValueObj struct {
	ID   string `json:"id,omitempty"`   //
	Name string `json:"name,omitempty"` //
}

func (r *ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesRules) UnmarshalJSON(data []byte) error {
	type Alias ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesRules
	aux := &struct {
		Value interface{} `json:"value"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	switch v := aux.Value.(type) {
	case string:
		r.Value = &v
	case []interface{}:
		strList := make([]string, len(v))
		for i, item := range v {
			strList[i] = item.(string)
		}
		r.ValueList = &strList
	case map[string]interface{}:
		valueObj := &ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesRulesValueObj{
			ID:   v["id"].(string),
			Name: v["name"].(string),
		}
		r.ValueObj = valueObj
	}
	return nil
}

type ResponseApplianceUpdateNetworkApplianceFirewallL7FirewallRules interface{}
type ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories struct {
	ApplicationCategories *[]ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesApplicationCategoriesApplicationCategories `json:"applicationCategories,omitempty"` //  The L7 firewall application categories and their associated applications for an MX network
}
type ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesApplicationCategoriesApplicationCategories struct {
	Applications *[]ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesApplicationCategoriesApplicationCategoriesApplications `json:"applications,omitempty"` // Details of the associated applications
	ID           string                                                                                                               `json:"id,omitempty"`           // The id of the category
	Name         string                                                                                                               `json:"name,omitempty"`         // The name of the category
}
type ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesApplicationCategoriesApplicationCategoriesApplications struct {
	ID   string `json:"id,omitempty"`   // The id of the application
	Name string `json:"name,omitempty"` // The name of the application
}
type ResponseApplianceUpdateNetworkApplianceFirewallMulticastForwarding struct {
	Network *ResponseApplianceUpdateNetworkApplianceFirewallMulticastForwardingNetwork `json:"network,omitempty"` // Network details
	Rules   *[]ResponseApplianceUpdateNetworkApplianceFirewallMulticastForwardingRules `json:"rules,omitempty"`   // Static multicast forwarding rules.
}
type ResponseApplianceUpdateNetworkApplianceFirewallMulticastForwardingNetwork struct {
	ID   string `json:"id,omitempty"`   // ID of the network whose multicast forwarding settings are returned.
	Name string `json:"name,omitempty"` // Name of the network whose multicast forwarding settings are returned.
}
type ResponseApplianceUpdateNetworkApplianceFirewallMulticastForwardingRules struct {
	Address     string   `json:"address,omitempty"`     // IP address
	Description string   `json:"description,omitempty"` // Forwarding rule description.
	VLANIDs     []string `json:"vlanIds,omitempty"`     // List of VLAN IDs
}
type ResponseApplianceGetNetworkApplianceFirewallOneToManyNatRules struct {
	Rules *[]ResponseApplianceGetNetworkApplianceFirewallOneToManyNatRulesRules `json:"rules,omitempty"` //
}
type ResponseApplianceGetNetworkApplianceFirewallOneToManyNatRulesRules struct {
	PortRules *[]ResponseApplianceGetNetworkApplianceFirewallOneToManyNatRulesRulesPortRules `json:"portRules,omitempty"` //
	PublicIP  string                                                                         `json:"publicIp,omitempty"`  //
	Uplink    string                                                                         `json:"uplink,omitempty"`    //
}
type ResponseApplianceGetNetworkApplianceFirewallOneToManyNatRulesRulesPortRules struct {
	AllowedIPs []string `json:"allowedIps,omitempty"` //
	LocalIP    string   `json:"localIp,omitempty"`    //
	LocalPort  string   `json:"localPort,omitempty"`  //
	Name       string   `json:"name,omitempty"`       //
	Protocol   string   `json:"protocol,omitempty"`   //
	PublicPort string   `json:"publicPort,omitempty"` //
}
type ResponseApplianceUpdateNetworkApplianceFirewallOneToManyNatRules interface{}
type ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRules struct {
	Rules *[]ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRulesRules `json:"rules,omitempty"` //
}
type ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRulesRules struct {
	AllowedInbound *[]ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRulesRulesAllowedInbound `json:"allowedInbound,omitempty"` //
	LanIP          string                                                                             `json:"lanIp,omitempty"`          //
	Name           string                                                                             `json:"name,omitempty"`           //
	PublicIP       string                                                                             `json:"publicIp,omitempty"`       //
	Uplink         string                                                                             `json:"uplink,omitempty"`         //
}
type ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRulesRulesAllowedInbound struct {
	AllowedIPs       []string `json:"allowedIps,omitempty"`       //
	DestinationPorts []string `json:"destinationPorts,omitempty"` //
	Protocol         string   `json:"protocol,omitempty"`         //
}
type ResponseApplianceUpdateNetworkApplianceFirewallOneToOneNatRules interface{}
type ResponseApplianceGetNetworkApplianceFirewallPortForwardingRules struct {
	Rules *[]ResponseApplianceGetNetworkApplianceFirewallPortForwardingRulesRules `json:"rules,omitempty"` // An array of port forwarding rules
}
type ResponseApplianceGetNetworkApplianceFirewallPortForwardingRulesRules struct {
	AllowedIPs []string `json:"allowedIps,omitempty"` // An array of ranges of WAN IP addresses that are allowed to make inbound connections on the specified ports or port ranges (or any)
	LanIP      string   `json:"lanIp,omitempty"`      // IP address of the device subject to port forwarding
	LocalPort  string   `json:"localPort,omitempty"`  // The port or port range that receives forwarded traffic from the WAN
	Name       string   `json:"name,omitempty"`       // Name of the rule
	Protocol   string   `json:"protocol,omitempty"`   // Protocol the rule applies to
	PublicPort string   `json:"publicPort,omitempty"` // The port or port range forwarded to the host on the LAN
	Uplink     string   `json:"uplink,omitempty"`     // The physical WAN interface on which the traffic arrives; allowed values vary by appliance model and configuration
}
type ResponseApplianceUpdateNetworkApplianceFirewallPortForwardingRules struct {
	Rules *[]ResponseApplianceUpdateNetworkApplianceFirewallPortForwardingRulesRules `json:"rules,omitempty"` // An array of port forwarding rules
}
type ResponseApplianceUpdateNetworkApplianceFirewallPortForwardingRulesRules struct {
	AllowedIPs []string `json:"allowedIps,omitempty"` // An array of ranges of WAN IP addresses that are allowed to make inbound connections on the specified ports or port ranges (or any)
	LanIP      string   `json:"lanIp,omitempty"`      // IP address of the device subject to port forwarding
	LocalPort  string   `json:"localPort,omitempty"`  // The port or port range that receives forwarded traffic from the WAN
	Name       string   `json:"name,omitempty"`       // Name of the rule
	Protocol   string   `json:"protocol,omitempty"`   // Protocol the rule applies to
	PublicPort string   `json:"publicPort,omitempty"` // The port or port range forwarded to the host on the LAN
	Uplink     string   `json:"uplink,omitempty"`     // The physical WAN interface on which the traffic arrives; allowed values vary by appliance model and configuration
}
type ResponseApplianceGetNetworkApplianceFirewallSettings struct {
	SpoofingProtection *ResponseApplianceGetNetworkApplianceFirewallSettingsSpoofingProtection `json:"spoofingProtection,omitempty"` //
}
type ResponseApplianceGetNetworkApplianceFirewallSettingsSpoofingProtection struct {
	IPSourceGuard *ResponseApplianceGetNetworkApplianceFirewallSettingsSpoofingProtectionIPSourceGuard `json:"ipSourceGuard,omitempty"` //
}
type ResponseApplianceGetNetworkApplianceFirewallSettingsSpoofingProtectionIPSourceGuard struct {
	Mode string `json:"mode,omitempty"` //
}
type ResponseApplianceUpdateNetworkApplianceFirewallSettings interface{}
type ResponseApplianceGetNetworkAppliancePorts []ResponseItemApplianceGetNetworkAppliancePorts // Array of ResponseApplianceGetNetworkAppliancePorts
type ResponseItemApplianceGetNetworkAppliancePorts struct {
	AccessPolicy        string `json:"accessPolicy,omitempty"`        // The name of the policy. Only applicable to Access ports.
	AllowedVLANs        string `json:"allowedVlans,omitempty"`        // Comma-delimited list of the VLAN ID's allowed on the port, or 'all' to permit all VLAN's on the port.
	DropUntaggedTraffic *bool  `json:"dropUntaggedTraffic,omitempty"` // Whether the trunk port can drop all untagged traffic.
	Enabled             *bool  `json:"enabled,omitempty"`             // The status of the port
	Number              *int   `json:"number,omitempty"`              // Number of the port
	Type                string `json:"type,omitempty"`                // The type of the port: 'access' or 'trunk'.
	VLAN                *int   `json:"vlan,omitempty"`                // Native VLAN when the port is in Trunk mode. Access VLAN when the port is in Access mode.
}
type ResponseApplianceGetNetworkAppliancePort struct {
	AccessPolicy        string `json:"accessPolicy,omitempty"`        // The name of the policy. Only applicable to Access ports.
	AllowedVLANs        string `json:"allowedVlans,omitempty"`        // Comma-delimited list of the VLAN ID's allowed on the port, or 'all' to permit all VLAN's on the port.
	DropUntaggedTraffic *bool  `json:"dropUntaggedTraffic,omitempty"` // Whether the trunk port can drop all untagged traffic.
	Enabled             *bool  `json:"enabled,omitempty"`             // The status of the port
	Number              *int   `json:"number,omitempty"`              // Number of the port
	Type                string `json:"type,omitempty"`                // The type of the port: 'access' or 'trunk'.
	VLAN                *int   `json:"vlan,omitempty"`                // Native VLAN when the port is in Trunk mode. Access VLAN when the port is in Access mode.
}
type ResponseApplianceUpdateNetworkAppliancePort struct {
	AccessPolicy        string `json:"accessPolicy,omitempty"`        // The name of the policy. Only applicable to Access ports.
	AllowedVLANs        string `json:"allowedVlans,omitempty"`        // Comma-delimited list of the VLAN ID's allowed on the port, or 'all' to permit all VLAN's on the port.
	DropUntaggedTraffic *bool  `json:"dropUntaggedTraffic,omitempty"` // Whether the trunk port can drop all untagged traffic.
	Enabled             *bool  `json:"enabled,omitempty"`             // The status of the port
	Number              *int   `json:"number,omitempty"`              // Number of the port
	Type                string `json:"type,omitempty"`                // The type of the port: 'access' or 'trunk'.
	VLAN                *int   `json:"vlan,omitempty"`                // Native VLAN when the port is in Trunk mode. Access VLAN when the port is in Access mode.
}
type ResponseApplianceGetNetworkAppliancePrefixesDelegatedStatics []ResponseItemApplianceGetNetworkAppliancePrefixesDelegatedStatics // Array of ResponseApplianceGetNetworkAppliancePrefixesDelegatedStatics
type ResponseItemApplianceGetNetworkAppliancePrefixesDelegatedStatics struct {
	CreatedAt               string                                                                  `json:"createdAt,omitempty"`               // Prefix creation time.
	Description             string                                                                  `json:"description,omitempty"`             // Identifying description for the prefix.
	Origin                  *ResponseItemApplianceGetNetworkAppliancePrefixesDelegatedStaticsOrigin `json:"origin,omitempty"`                  // WAN1/WAN2/Independent prefix.
	Prefix                  string                                                                  `json:"prefix,omitempty"`                  // IPv6 prefix/prefix length.
	StaticDelegatedPrefixID string                                                                  `json:"staticDelegatedPrefixId,omitempty"` // Static delegated prefix id.
	UpdatedAt               string                                                                  `json:"updatedAt,omitempty"`               // Prefix Updated time.
}
type ResponseItemApplianceGetNetworkAppliancePrefixesDelegatedStaticsOrigin struct {
	Interfaces []string `json:"interfaces,omitempty"` // Uplink provided or independent
	Type       string   `json:"type,omitempty"`       // Origin type
}
type ResponseApplianceCreateNetworkAppliancePrefixesDelegatedStatic interface{}
type ResponseApplianceGetNetworkAppliancePrefixesDelegatedStatic struct {
	CreatedAt               string                                                             `json:"createdAt,omitempty"`               // Prefix creation time.
	Description             string                                                             `json:"description,omitempty"`             // Identifying description for the prefix.
	Origin                  *ResponseApplianceGetNetworkAppliancePrefixesDelegatedStaticOrigin `json:"origin,omitempty"`                  // WAN1/WAN2/Independent prefix.
	Prefix                  string                                                             `json:"prefix,omitempty"`                  // IPv6 prefix/prefix length.
	StaticDelegatedPrefixID string                                                             `json:"staticDelegatedPrefixId,omitempty"` // Static delegated prefix id.
	UpdatedAt               string                                                             `json:"updatedAt,omitempty"`               // Prefix Updated time.
}
type ResponseApplianceGetNetworkAppliancePrefixesDelegatedStaticOrigin struct {
	Interfaces []string `json:"interfaces,omitempty"` // Uplink provided or independent
	Type       string   `json:"type,omitempty"`       // Origin type
}
type ResponseApplianceUpdateNetworkAppliancePrefixesDelegatedStatic interface{}
type ResponseApplianceGetNetworkApplianceRfProfiles struct {
	Assigned *[]ResponseApplianceGetNetworkApplianceRfProfilesAssigned `json:"assigned,omitempty"` // RF Profiles
}
type ResponseApplianceGetNetworkApplianceRfProfilesAssigned struct {
	FiveGhzSettings    *ResponseApplianceGetNetworkApplianceRfProfilesAssignedFiveGhzSettings    `json:"fiveGhzSettings,omitempty"`    // Settings related to 5Ghz band.
	ID                 string                                                                    `json:"id,omitempty"`                 // ID of the RF Profile.
	Name               string                                                                    `json:"name,omitempty"`               // The name of the profile.
	NetworkID          string                                                                    `json:"networkId,omitempty"`          // ID of network this RF Profile belongs in.
	PerSSIDSettings    *ResponseApplianceGetNetworkApplianceRfProfilesAssignedPerSSIDSettings    `json:"perSsidSettings,omitempty"`    // Per-SSID radio settings by number.
	TwoFourGhzSettings *ResponseApplianceGetNetworkApplianceRfProfilesAssignedTwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"` // Settings related to 2.4Ghz band.
}
type ResponseApplianceGetNetworkApplianceRfProfilesAssignedFiveGhzSettings struct {
	AxEnabled  *bool `json:"axEnabled,omitempty"`  // Whether ax radio on 5Ghz band is on or off.
	MinBitrate *int  `json:"minBitrate,omitempty"` // Min bitrate (Mbps) of 2.4Ghz band.
}
type ResponseApplianceGetNetworkApplianceRfProfilesAssignedPerSSIDSettings struct {
	Status1 *ResponseApplianceGetNetworkApplianceRfProfilesAssignedPerSSIDSettings1 `json:"1,omitempty"` // Settings for SSID 1.
	Status2 *ResponseApplianceGetNetworkApplianceRfProfilesAssignedPerSSIDSettings2 `json:"2,omitempty"` // Settings for SSID 2.
	Status3 *ResponseApplianceGetNetworkApplianceRfProfilesAssignedPerSSIDSettings3 `json:"3,omitempty"` // Settings for SSID 3.
	Status4 *ResponseApplianceGetNetworkApplianceRfProfilesAssignedPerSSIDSettings4 `json:"4,omitempty"` // Settings for SSID 4.
}
type ResponseApplianceGetNetworkApplianceRfProfilesAssignedPerSSIDSettings1 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Band mode of this SSID
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Whether this SSID steers clients to the most open band between 2.4 GHz and 5 GHz.
}
type ResponseApplianceGetNetworkApplianceRfProfilesAssignedPerSSIDSettings2 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Band mode of this SSID
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Whether this SSID steers clients to the most open band between 2.4 GHz and 5 GHz.
}
type ResponseApplianceGetNetworkApplianceRfProfilesAssignedPerSSIDSettings3 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Band mode of this SSID
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Whether this SSID steers clients to the most open band between 2.4 GHz and 5 GHz.
}
type ResponseApplianceGetNetworkApplianceRfProfilesAssignedPerSSIDSettings4 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Band mode of this SSID
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Whether this SSID steers clients to the most open band between 2.4 GHz and 5 GHz.
}
type ResponseApplianceGetNetworkApplianceRfProfilesAssignedTwoFourGhzSettings struct {
	AxEnabled  *bool    `json:"axEnabled,omitempty"`  // Whether ax radio on 2.4Ghz band is on or off.
	MinBitrate *float64 `json:"minBitrate,omitempty"` // Min bitrate (Mbps) of 2.4Ghz band.
}
type ResponseApplianceCreateNetworkApplianceRfProfile struct {
	FiveGhzSettings    *ResponseApplianceCreateNetworkApplianceRfProfileFiveGhzSettings    `json:"fiveGhzSettings,omitempty"`    // Settings related to 5Ghz band.
	ID                 string                                                              `json:"id,omitempty"`                 // ID of the RF Profile.
	Name               string                                                              `json:"name,omitempty"`               // The name of the profile.
	NetworkID          string                                                              `json:"networkId,omitempty"`          // ID of network this RF Profile belongs in.
	PerSSIDSettings    *ResponseApplianceCreateNetworkApplianceRfProfilePerSSIDSettings    `json:"perSsidSettings,omitempty"`    // Per-SSID radio settings by number.
	TwoFourGhzSettings *ResponseApplianceCreateNetworkApplianceRfProfileTwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"` // Settings related to 2.4Ghz band.
}
type ResponseApplianceCreateNetworkApplianceRfProfileFiveGhzSettings struct {
	AxEnabled  *bool `json:"axEnabled,omitempty"`  // Whether ax radio on 5Ghz band is on or off.
	MinBitrate *int  `json:"minBitrate,omitempty"` // Min bitrate (Mbps) of 2.4Ghz band.
}
type ResponseApplianceCreateNetworkApplianceRfProfilePerSSIDSettings struct {
	Status1 *ResponseApplianceCreateNetworkApplianceRfProfilePerSSIDSettings1 `json:"1,omitempty"` // Settings for SSID 1.
	Status2 *ResponseApplianceCreateNetworkApplianceRfProfilePerSSIDSettings2 `json:"2,omitempty"` // Settings for SSID 2.
	Status3 *ResponseApplianceCreateNetworkApplianceRfProfilePerSSIDSettings3 `json:"3,omitempty"` // Settings for SSID 3.
	Status4 *ResponseApplianceCreateNetworkApplianceRfProfilePerSSIDSettings4 `json:"4,omitempty"` // Settings for SSID 4.
}
type ResponseApplianceCreateNetworkApplianceRfProfilePerSSIDSettings1 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Band mode of this SSID
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Whether this SSID steers clients to the most open band between 2.4 GHz and 5 GHz.
}
type ResponseApplianceCreateNetworkApplianceRfProfilePerSSIDSettings2 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Band mode of this SSID
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Whether this SSID steers clients to the most open band between 2.4 GHz and 5 GHz.
}
type ResponseApplianceCreateNetworkApplianceRfProfilePerSSIDSettings3 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Band mode of this SSID
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Whether this SSID steers clients to the most open band between 2.4 GHz and 5 GHz.
}
type ResponseApplianceCreateNetworkApplianceRfProfilePerSSIDSettings4 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Band mode of this SSID
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Whether this SSID steers clients to the most open band between 2.4 GHz and 5 GHz.
}
type ResponseApplianceCreateNetworkApplianceRfProfileTwoFourGhzSettings struct {
	AxEnabled  *bool    `json:"axEnabled,omitempty"`  // Whether ax radio on 2.4Ghz band is on or off.
	MinBitrate *float64 `json:"minBitrate,omitempty"` // Min bitrate (Mbps) of 2.4Ghz band.
}
type ResponseApplianceGetNetworkApplianceRfProfile struct {
	FiveGhzSettings    *ResponseApplianceGetNetworkApplianceRfProfileFiveGhzSettings    `json:"fiveGhzSettings,omitempty"`    // Settings related to 5Ghz band.
	ID                 string                                                           `json:"id,omitempty"`                 // ID of the RF Profile.
	Name               string                                                           `json:"name,omitempty"`               // The name of the profile.
	NetworkID          string                                                           `json:"networkId,omitempty"`          // ID of network this RF Profile belongs in.
	PerSSIDSettings    *ResponseApplianceGetNetworkApplianceRfProfilePerSSIDSettings    `json:"perSsidSettings,omitempty"`    // Per-SSID radio settings by number.
	TwoFourGhzSettings *ResponseApplianceGetNetworkApplianceRfProfileTwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"` // Settings related to 2.4Ghz band.
}
type ResponseApplianceGetNetworkApplianceRfProfileFiveGhzSettings struct {
	AxEnabled  *bool `json:"axEnabled,omitempty"`  // Whether ax radio on 5Ghz band is on or off.
	MinBitrate *int  `json:"minBitrate,omitempty"` // Min bitrate (Mbps) of 2.4Ghz band.
}
type ResponseApplianceGetNetworkApplianceRfProfilePerSSIDSettings struct {
	Status1 *ResponseApplianceGetNetworkApplianceRfProfilePerSSIDSettings1 `json:"1,omitempty"` // Settings for SSID 1.
	Status2 *ResponseApplianceGetNetworkApplianceRfProfilePerSSIDSettings2 `json:"2,omitempty"` // Settings for SSID 2.
	Status3 *ResponseApplianceGetNetworkApplianceRfProfilePerSSIDSettings3 `json:"3,omitempty"` // Settings for SSID 3.
	Status4 *ResponseApplianceGetNetworkApplianceRfProfilePerSSIDSettings4 `json:"4,omitempty"` // Settings for SSID 4.
}
type ResponseApplianceGetNetworkApplianceRfProfilePerSSIDSettings1 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Band mode of this SSID
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Whether this SSID steers clients to the most open band between 2.4 GHz and 5 GHz.
}
type ResponseApplianceGetNetworkApplianceRfProfilePerSSIDSettings2 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Band mode of this SSID
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Whether this SSID steers clients to the most open band between 2.4 GHz and 5 GHz.
}
type ResponseApplianceGetNetworkApplianceRfProfilePerSSIDSettings3 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Band mode of this SSID
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Whether this SSID steers clients to the most open band between 2.4 GHz and 5 GHz.
}
type ResponseApplianceGetNetworkApplianceRfProfilePerSSIDSettings4 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Band mode of this SSID
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Whether this SSID steers clients to the most open band between 2.4 GHz and 5 GHz.
}
type ResponseApplianceGetNetworkApplianceRfProfileTwoFourGhzSettings struct {
	AxEnabled  *bool    `json:"axEnabled,omitempty"`  // Whether ax radio on 2.4Ghz band is on or off.
	MinBitrate *float64 `json:"minBitrate,omitempty"` // Min bitrate (Mbps) of 2.4Ghz band.
}
type ResponseApplianceUpdateNetworkApplianceRfProfile struct {
	FiveGhzSettings    *ResponseApplianceUpdateNetworkApplianceRfProfileFiveGhzSettings    `json:"fiveGhzSettings,omitempty"`    // Settings related to 5Ghz band.
	ID                 string                                                              `json:"id,omitempty"`                 // ID of the RF Profile.
	Name               string                                                              `json:"name,omitempty"`               // The name of the profile.
	NetworkID          string                                                              `json:"networkId,omitempty"`          // ID of network this RF Profile belongs in.
	PerSSIDSettings    *ResponseApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings    `json:"perSsidSettings,omitempty"`    // Per-SSID radio settings by number.
	TwoFourGhzSettings *ResponseApplianceUpdateNetworkApplianceRfProfileTwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"` // Settings related to 2.4Ghz band.
}
type ResponseApplianceUpdateNetworkApplianceRfProfileFiveGhzSettings struct {
	AxEnabled  *bool `json:"axEnabled,omitempty"`  // Whether ax radio on 5Ghz band is on or off.
	MinBitrate *int  `json:"minBitrate,omitempty"` // Min bitrate (Mbps) of 2.4Ghz band.
}
type ResponseApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings struct {
	Status1 *ResponseApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings1 `json:"1,omitempty"` // Settings for SSID 1.
	Status2 *ResponseApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings2 `json:"2,omitempty"` // Settings for SSID 2.
	Status3 *ResponseApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings3 `json:"3,omitempty"` // Settings for SSID 3.
	Status4 *ResponseApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings4 `json:"4,omitempty"` // Settings for SSID 4.
}
type ResponseApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings1 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Band mode of this SSID
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Whether this SSID steers clients to the most open band between 2.4 GHz and 5 GHz.
}
type ResponseApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings2 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Band mode of this SSID
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Whether this SSID steers clients to the most open band between 2.4 GHz and 5 GHz.
}
type ResponseApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings3 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Band mode of this SSID
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Whether this SSID steers clients to the most open band between 2.4 GHz and 5 GHz.
}
type ResponseApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings4 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Band mode of this SSID
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Whether this SSID steers clients to the most open band between 2.4 GHz and 5 GHz.
}
type ResponseApplianceUpdateNetworkApplianceRfProfileTwoFourGhzSettings struct {
	AxEnabled  *bool    `json:"axEnabled,omitempty"`  // Whether ax radio on 2.4Ghz band is on or off.
	MinBitrate *float64 `json:"minBitrate,omitempty"` // Min bitrate (Mbps) of 2.4Ghz band.
}
type ResponseApplianceUpdateNetworkApplianceSdwanInternetPolicies struct {
	WanTrafficUplinkPreferences *[]ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences `json:"wanTrafficUplinkPreferences,omitempty"` // policies with respective traffic filters for an MX network
}
type ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences struct {
	FailOverCriterion string                                                                                                   `json:"failOverCriterion,omitempty"` // WAN failover and failback behavior
	PerformanceClass  *ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesPerformanceClass `json:"performanceClass,omitempty"`  // Performance class setting for uplink preference rule
	PreferredUplink   string                                                                                                   `json:"preferredUplink,omitempty"`   // Preferred uplink for uplink preference rule. Must be one of: 'wan1', 'wan2', 'bestForVoIP', 'loadBalancing' or 'defaultUplink'
	TrafficFilters    *[]ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFilters `json:"trafficFilters,omitempty"`    // Traffic filters
}
type ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesPerformanceClass struct {
	BuiltinPerformanceClassName string `json:"builtinPerformanceClassName,omitempty"` // Name of builtin performance class. Must be present when performanceClass type is 'builtin' and value must be one of: 'VoIP'
	CustomPerformanceClassID    string `json:"customPerformanceClassId,omitempty"`    // ID of created custom performance class, must be present when performanceClass type is "custom"
	Type                        string `json:"type,omitempty"`                        // Type of this performance class. Must be one of: 'builtin' or 'custom'
}
type ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFilters struct {
	Type  string                                                                                                      `json:"type,omitempty"`  // Traffic filter type. Must be 'custom', 'major_application', 'application (NBAR)', if type is 'application', you can pass either an NBAR App Category or Application
	Value *ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValue `json:"value,omitempty"` // Value of traffic filter
}
type ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValue struct {
	Destination *ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueDestination `json:"destination,omitempty"` // Destination of 'custom' type traffic filter
	Protocol    string                                                                                                                 `json:"protocol,omitempty"`    // Protocol of the traffic filter. Must be one of: 'tcp', 'udp', 'icmp6' or 'any'
	Source      *ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueSource      `json:"source,omitempty"`      // Source of traffic filter
}
type ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueDestination struct {
	Applications *[]ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueDestinationApplications `json:"applications,omitempty"` // list of application objects (either majorApplication or nbar)
	Cidr         string                                                                                                                               `json:"cidr,omitempty"`         // CIDR format address (e.g."192.168.10.1", which is the same as "192.168.10.1/32"), or "any"
	Port         string                                                                                                                               `json:"port,omitempty"`         // E.g.: "any", "0" (also means "any"), "8080", "1-1024"
}
type ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueDestinationApplications struct {
	ID   string `json:"id,omitempty"`   // Id of the major application, or a list of NBAR Application Category or Application selections
	Name string `json:"name,omitempty"` // Name of the major application or application category selected
	Type string `json:"type,omitempty"` // app type (major or nbar)
}
type ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueSource struct {
	Cidr string `json:"cidr,omitempty"` // CIDR format address (e.g."192.168.10.1", which is the same as "192.168.10.1/32"), or "any". Cannot be used in combination with the "vlan" property
	Host *int   `json:"host,omitempty"` // Host ID in the VLAN. Should not exceed the VLAN subnet capacity. Must be used along with the "vlan" property and is currently only available under a template network.
	Port string `json:"port,omitempty"` // E.g.: "any", "0" (also means "any"), "8080", "1-1024"
	VLAN *int   `json:"vlan,omitempty"` // VLAN ID of the configured VLAN in the Meraki network. Cannot be used in combination with the "cidr" property and is currently only available under a template network.
}
type ResponseApplianceGetNetworkApplianceSecurityEvents []ResponseItemApplianceGetNetworkApplianceSecurityEvents // Array of ResponseApplianceGetNetworkApplianceSecurityEvents
type ResponseItemApplianceGetNetworkApplianceSecurityEvents struct {
	Action          string `json:"action,omitempty"`          //
	CanonicalName   string `json:"canonicalName,omitempty"`   //
	ClientIP        string `json:"clientIp,omitempty"`        //
	ClientMac       string `json:"clientMac,omitempty"`       //
	ClientName      string `json:"clientName,omitempty"`      //
	DestIP          string `json:"destIp,omitempty"`          //
	DestinationPort *int   `json:"destinationPort,omitempty"` //
	Disposition     string `json:"disposition,omitempty"`     //
	EventType       string `json:"eventType,omitempty"`       //
	FileHash        string `json:"fileHash,omitempty"`        //
	FileSizeBytes   *int   `json:"fileSizeBytes,omitempty"`   //
	FileType        string `json:"fileType,omitempty"`        //
	Protocol        string `json:"protocol,omitempty"`        //
	SrcIP           string `json:"srcIp,omitempty"`           //
	Ts              string `json:"ts,omitempty"`              //
	URI             string `json:"uri,omitempty"`             //
}
type ResponseApplianceGetNetworkApplianceSecurityIntrusion struct {
	IDsRulesets       string                                                                  `json:"idsRulesets,omitempty"`       // Intrusion detection ruleset
	Mode              string                                                                  `json:"mode,omitempty"`              // Intrusion detection mode
	ProtectedNetworks *ResponseApplianceGetNetworkApplianceSecurityIntrusionProtectedNetworks `json:"protectedNetworks,omitempty"` // Networks included in and excluded from the detection engine
}
type ResponseApplianceGetNetworkApplianceSecurityIntrusionProtectedNetworks struct {
	ExcludedCidr []string `json:"excludedCidr,omitempty"` // List of IP addresses or subnets being excluded from protection
	IncludedCidr []string `json:"includedCidr,omitempty"` // List of IP addresses or subnets being protected
	UseDefault   *bool    `json:"useDefault,omitempty"`   // Whether special IPv4 addresses should be used (see: https://tools.ietf.org/html/rfc5735)
}
type ResponseApplianceUpdateNetworkApplianceSecurityIntrusion struct {
	IDsRulesets       string                                                                     `json:"idsRulesets,omitempty"`       // Intrusion detection ruleset
	Mode              string                                                                     `json:"mode,omitempty"`              // Intrusion detection mode
	ProtectedNetworks *ResponseApplianceUpdateNetworkApplianceSecurityIntrusionProtectedNetworks `json:"protectedNetworks,omitempty"` // Networks included in and excluded from the detection engine
}
type ResponseApplianceUpdateNetworkApplianceSecurityIntrusionProtectedNetworks struct {
	ExcludedCidr []string `json:"excludedCidr,omitempty"` // List of IP addresses or subnets being excluded from protection
	IncludedCidr []string `json:"includedCidr,omitempty"` // List of IP addresses or subnets being protected
	UseDefault   *bool    `json:"useDefault,omitempty"`   // Whether special IPv4 addresses should be used (see: https://tools.ietf.org/html/rfc5735)
}
type ResponseApplianceGetNetworkApplianceSecurityMalware struct {
	AllowedFiles *[]ResponseApplianceGetNetworkApplianceSecurityMalwareAllowedFiles `json:"allowedFiles,omitempty"` // Sha256 digests of files permitted by the malware detection engine
	AllowedURLs  *[]ResponseApplianceGetNetworkApplianceSecurityMalwareAllowedURLs  `json:"allowedUrls,omitempty"`  // URLs permitted by the malware detection engine
	Mode         string                                                             `json:"mode,omitempty"`         // Current status of malware prevention
}
type ResponseApplianceGetNetworkApplianceSecurityMalwareAllowedFiles struct {
	Comment string `json:"comment,omitempty"` // Comment about the allowed file
	Sha256  string `json:"sha256,omitempty"`  // The sha256 digest of allowed file
}
type ResponseApplianceGetNetworkApplianceSecurityMalwareAllowedURLs struct {
	Comment string `json:"comment,omitempty"` // Comment about the allowed URL
	URL     string `json:"url,omitempty"`     // The allowed URL
}
type ResponseApplianceUpdateNetworkApplianceSecurityMalware struct {
	AllowedFiles *[]ResponseApplianceUpdateNetworkApplianceSecurityMalwareAllowedFiles `json:"allowedFiles,omitempty"` // Sha256 digests of files permitted by the malware detection engine
	AllowedURLs  *[]ResponseApplianceUpdateNetworkApplianceSecurityMalwareAllowedURLs  `json:"allowedUrls,omitempty"`  // URLs permitted by the malware detection engine
	Mode         string                                                                `json:"mode,omitempty"`         // Current status of malware prevention
}
type ResponseApplianceUpdateNetworkApplianceSecurityMalwareAllowedFiles struct {
	Comment string `json:"comment,omitempty"` // Comment about the allowed file
	Sha256  string `json:"sha256,omitempty"`  // The sha256 digest of allowed file
}
type ResponseApplianceUpdateNetworkApplianceSecurityMalwareAllowedURLs struct {
	Comment string `json:"comment,omitempty"` // Comment about the allowed URL
	URL     string `json:"url,omitempty"`     // The allowed URL
}
type ResponseApplianceGetNetworkApplianceSettings struct {
	ClientTrackingMethod string                                                  `json:"clientTrackingMethod,omitempty"` // Client tracking method of a network
	DeploymentMode       string                                                  `json:"deploymentMode,omitempty"`       // Deployment mode of a network
	DynamicDNS           *ResponseApplianceGetNetworkApplianceSettingsDynamicDNS `json:"dynamicDns,omitempty"`           // Dynamic DNS settings for a network
}
type ResponseApplianceGetNetworkApplianceSettingsDynamicDNS struct {
	Enabled *bool  `json:"enabled,omitempty"` // Dynamic DNS enabled
	Prefix  string `json:"prefix,omitempty"`  // Dynamic DNS url prefix. DDNS must be enabled to update
	URL     string `json:"url,omitempty"`     // Dynamic DNS url. DDNS must be enabled to update
}
type ResponseApplianceUpdateNetworkApplianceSettings struct {
	ClientTrackingMethod string                                                     `json:"clientTrackingMethod,omitempty"` // Client tracking method of a network
	DeploymentMode       string                                                     `json:"deploymentMode,omitempty"`       // Deployment mode of a network
	DynamicDNS           *ResponseApplianceUpdateNetworkApplianceSettingsDynamicDNS `json:"dynamicDns,omitempty"`           // Dynamic DNS settings for a network
}
type ResponseApplianceUpdateNetworkApplianceSettingsDynamicDNS struct {
	Enabled *bool  `json:"enabled,omitempty"` // Dynamic DNS enabled
	Prefix  string `json:"prefix,omitempty"`  // Dynamic DNS url prefix. DDNS must be enabled to update
	URL     string `json:"url,omitempty"`     // Dynamic DNS url. DDNS must be enabled to update
}
type ResponseApplianceGetNetworkApplianceSingleLan struct {
	ApplianceIP   string                                                      `json:"applianceIp,omitempty"`   // The local IP of the appliance on the single LAN
	IPv6          *ResponseApplianceGetNetworkApplianceSingleLanIPv6          `json:"ipv6,omitempty"`          // IPv6 configuration on the single LAN
	MandatoryDhcp *ResponseApplianceGetNetworkApplianceSingleLanMandatoryDhcp `json:"mandatoryDhcp,omitempty"` // Mandatory DHCP will enforce that clients connecting to this single LAN must use the IP address assigned by the DHCP server. Clients who use a static IP address won't be able to associate. Only available on firmware versions 17.0 and above
	Subnet        string                                                      `json:"subnet,omitempty"`        // The subnet of the single LAN
}
type ResponseApplianceGetNetworkApplianceSingleLanIPv6 struct {
	Enabled           *bool                                                                 `json:"enabled,omitempty"`           // Enable IPv6 on single LAN
	PrefixAssignments *[]ResponseApplianceGetNetworkApplianceSingleLanIPv6PrefixAssignments `json:"prefixAssignments,omitempty"` // Prefix assignments on the single LAN
}
type ResponseApplianceGetNetworkApplianceSingleLanIPv6PrefixAssignments struct {
	Autonomous         *bool                                                                     `json:"autonomous,omitempty"`         // Auto assign a /64 prefix from the origin to the single LAN
	Origin             *ResponseApplianceGetNetworkApplianceSingleLanIPv6PrefixAssignmentsOrigin `json:"origin,omitempty"`             // The origin of the prefix
	StaticApplianceIP6 string                                                                    `json:"staticApplianceIp6,omitempty"` // Manual configuration of the IPv6 Appliance IP
	StaticPrefix       string                                                                    `json:"staticPrefix,omitempty"`       // Manual configuration of a /64 prefix on the single LAN
}
type ResponseApplianceGetNetworkApplianceSingleLanIPv6PrefixAssignmentsOrigin struct {
	Interfaces []string `json:"interfaces,omitempty"` // Interfaces associated with the prefix
	Type       string   `json:"type,omitempty"`       // Type of the origin
}
type ResponseApplianceGetNetworkApplianceSingleLanMandatoryDhcp struct {
	Enabled *bool `json:"enabled,omitempty"` // Enable Mandatory DHCP on single LAN.
}
type ResponseApplianceUpdateNetworkApplianceSingleLan struct {
	ApplianceIP   string                                                         `json:"applianceIp,omitempty"`   // The local IP of the appliance on the single LAN
	IPv6          *ResponseApplianceUpdateNetworkApplianceSingleLanIPv6          `json:"ipv6,omitempty"`          // IPv6 configuration on the single LAN
	MandatoryDhcp *ResponseApplianceUpdateNetworkApplianceSingleLanMandatoryDhcp `json:"mandatoryDhcp,omitempty"` // Mandatory DHCP will enforce that clients connecting to this single LAN must use the IP address assigned by the DHCP server. Clients who use a static IP address won't be able to associate. Only available on firmware versions 17.0 and above
	Subnet        string                                                         `json:"subnet,omitempty"`        // The subnet of the single LAN
}
type ResponseApplianceUpdateNetworkApplianceSingleLanIPv6 struct {
	Enabled           *bool                                                                    `json:"enabled,omitempty"`           // Enable IPv6 on single LAN
	PrefixAssignments *[]ResponseApplianceUpdateNetworkApplianceSingleLanIPv6PrefixAssignments `json:"prefixAssignments,omitempty"` // Prefix assignments on the single LAN
}
type ResponseApplianceUpdateNetworkApplianceSingleLanIPv6PrefixAssignments struct {
	Autonomous         *bool                                                                        `json:"autonomous,omitempty"`         // Auto assign a /64 prefix from the origin to the single LAN
	Origin             *ResponseApplianceUpdateNetworkApplianceSingleLanIPv6PrefixAssignmentsOrigin `json:"origin,omitempty"`             // The origin of the prefix
	StaticApplianceIP6 string                                                                       `json:"staticApplianceIp6,omitempty"` // Manual configuration of the IPv6 Appliance IP
	StaticPrefix       string                                                                       `json:"staticPrefix,omitempty"`       // Manual configuration of a /64 prefix on the single LAN
}
type ResponseApplianceUpdateNetworkApplianceSingleLanIPv6PrefixAssignmentsOrigin struct {
	Interfaces []string `json:"interfaces,omitempty"` // Interfaces associated with the prefix
	Type       string   `json:"type,omitempty"`       // Type of the origin
}
type ResponseApplianceUpdateNetworkApplianceSingleLanMandatoryDhcp struct {
	Enabled *bool `json:"enabled,omitempty"` // Enable Mandatory DHCP on single LAN.
}
type ResponseApplianceGetNetworkApplianceSSIDs []ResponseItemApplianceGetNetworkApplianceSSIDs // Array of ResponseApplianceGetNetworkApplianceSsids
type ResponseItemApplianceGetNetworkApplianceSSIDs struct {
	AuthMode          string                                                        `json:"authMode,omitempty"`          // The association control method for the SSID.
	DefaultVLANID     *int                                                          `json:"defaultVlanId,omitempty"`     // The VLAN ID of the VLAN associated to this SSID.
	Enabled           *bool                                                         `json:"enabled,omitempty"`           // Whether or not the SSID is enabled.
	EncryptionMode    string                                                        `json:"encryptionMode,omitempty"`    // The psk encryption mode for the SSID.
	Name              string                                                        `json:"name,omitempty"`              // The name of the SSID.
	Number            *int                                                          `json:"number,omitempty"`            // The number of the SSID.
	RadiusServers     *[]ResponseItemApplianceGetNetworkApplianceSSIDsRadiusServers `json:"radiusServers,omitempty"`     // The RADIUS 802.1x servers to be used for authentication.
	Visible           *bool                                                         `json:"visible,omitempty"`           // Boolean indicating whether the MX should advertise or hide this SSID.
	WpaEncryptionMode string                                                        `json:"wpaEncryptionMode,omitempty"` // WPA encryption mode for the SSID.
}
type ResponseItemApplianceGetNetworkApplianceSSIDsRadiusServers struct {
	Host string `json:"host,omitempty"` // The IP address of your RADIUS server.
	Port *int   `json:"port,omitempty"` // The UDP port your RADIUS servers listens on for Access-requests.
}
type ResponseApplianceGetNetworkApplianceSSID struct {
	AuthMode          string                                                   `json:"authMode,omitempty"`          // The association control method for the SSID.
	DefaultVLANID     *int                                                     `json:"defaultVlanId,omitempty"`     // The VLAN ID of the VLAN associated to this SSID.
	Enabled           *bool                                                    `json:"enabled,omitempty"`           // Whether or not the SSID is enabled.
	EncryptionMode    string                                                   `json:"encryptionMode,omitempty"`    // The psk encryption mode for the SSID.
	Name              string                                                   `json:"name,omitempty"`              // The name of the SSID.
	Number            *int                                                     `json:"number,omitempty"`            // The number of the SSID.
	RadiusServers     *[]ResponseApplianceGetNetworkApplianceSSIDRadiusServers `json:"radiusServers,omitempty"`     // The RADIUS 802.1x servers to be used for authentication.
	Visible           *bool                                                    `json:"visible,omitempty"`           // Boolean indicating whether the MX should advertise or hide this SSID.
	WpaEncryptionMode string                                                   `json:"wpaEncryptionMode,omitempty"` // WPA encryption mode for the SSID.
}
type ResponseApplianceGetNetworkApplianceSSIDRadiusServers struct {
	Host string `json:"host,omitempty"` // The IP address of your RADIUS server.
	Port *int   `json:"port,omitempty"` // The UDP port your RADIUS servers listens on for Access-requests.
}
type ResponseApplianceUpdateNetworkApplianceSSID struct {
	AuthMode          string                                                      `json:"authMode,omitempty"`          // The association control method for the SSID.
	DefaultVLANID     *int                                                        `json:"defaultVlanId,omitempty"`     // The VLAN ID of the VLAN associated to this SSID.
	Enabled           *bool                                                       `json:"enabled,omitempty"`           // Whether or not the SSID is enabled.
	EncryptionMode    string                                                      `json:"encryptionMode,omitempty"`    // The psk encryption mode for the SSID.
	Name              string                                                      `json:"name,omitempty"`              // The name of the SSID.
	Number            *int                                                        `json:"number,omitempty"`            // The number of the SSID.
	RadiusServers     *[]ResponseApplianceUpdateNetworkApplianceSSIDRadiusServers `json:"radiusServers,omitempty"`     // The RADIUS 802.1x servers to be used for authentication.
	Visible           *bool                                                       `json:"visible,omitempty"`           // Boolean indicating whether the MX should advertise or hide this SSID.
	WpaEncryptionMode string                                                      `json:"wpaEncryptionMode,omitempty"` // WPA encryption mode for the SSID.
}
type ResponseApplianceUpdateNetworkApplianceSSIDRadiusServers struct {
	Host string `json:"host,omitempty"` // The IP address of your RADIUS server.
	Port *int   `json:"port,omitempty"` // The UDP port your RADIUS servers listens on for Access-requests.
}
type ResponseApplianceGetNetworkApplianceStaticRoutes []ResponseItemApplianceGetNetworkApplianceStaticRoutes // Array of ResponseApplianceGetNetworkApplianceStaticRoutes
type ResponseItemApplianceGetNetworkApplianceStaticRoutes struct {
	Enabled            *bool                                                                   `json:"enabled,omitempty"`            // Whether the route is enabled or not
	FixedIPAssignments *ResponseItemApplianceGetNetworkApplianceStaticRoutesFixedIPAssignments `json:"fixedIpAssignments,omitempty"` // Fixed DHCP IP assignments on the route
	GatewayIP          string                                                                  `json:"gatewayIp,omitempty"`          // Gateway IP address (next hop)
	GatewayVLANID      *int                                                                    `json:"gatewayVlanId,omitempty"`      // Gateway VLAN ID
	ID                 string                                                                  `json:"id,omitempty"`                 // Route ID
	IPVersion          *int                                                                    `json:"ipVersion,omitempty"`          // IP protocol version
	Name               string                                                                  `json:"name,omitempty"`               // Name of the route
	NetworkID          string                                                                  `json:"networkId,omitempty"`          // Network ID
	ReservedIPRanges   *[]ResponseItemApplianceGetNetworkApplianceStaticRoutesReservedIPRanges `json:"reservedIpRanges,omitempty"`   // DHCP reserved IP ranges
	Subnet             string                                                                  `json:"subnet,omitempty"`             // Subnet of the route
}
type ResponseItemApplianceGetNetworkApplianceStaticRoutesFixedIPAssignments struct {
	Status223344556677 *ResponseItemApplianceGetNetworkApplianceStaticRoutesFixedIPAssignments223344556677 `json:"22:33:44:55:66:77,omitempty"` //
}
type ResponseItemApplianceGetNetworkApplianceStaticRoutesFixedIPAssignments223344556677 struct {
	IP   string `json:"ip,omitempty"`   //
	Name string `json:"name,omitempty"` //
}
type ResponseItemApplianceGetNetworkApplianceStaticRoutesReservedIPRanges struct {
	Comment string `json:"comment,omitempty"` // Description of the range
	End     string `json:"end,omitempty"`     // Last address in the reserved range
	Start   string `json:"start,omitempty"`   // First address in the reserved range
}
type ResponseApplianceCreateNetworkApplianceStaticRoute struct {
	Enabled            *bool                                                                 `json:"enabled,omitempty"`            // Whether the route is enabled or not
	FixedIPAssignments *ResponseApplianceCreateNetworkApplianceStaticRouteFixedIPAssignments `json:"fixedIpAssignments,omitempty"` // Fixed DHCP IP assignments on the route
	GatewayIP          string                                                                `json:"gatewayIp,omitempty"`          // Gateway IP address (next hop)
	GatewayVLANID      *int                                                                  `json:"gatewayVlanId,omitempty"`      // Gateway VLAN ID
	ID                 string                                                                `json:"id,omitempty"`                 // Route ID
	IPVersion          *int                                                                  `json:"ipVersion,omitempty"`          // IP protocol version
	Name               string                                                                `json:"name,omitempty"`               // Name of the route
	NetworkID          string                                                                `json:"networkId,omitempty"`          // Network ID
	ReservedIPRanges   *[]ResponseApplianceCreateNetworkApplianceStaticRouteReservedIPRanges `json:"reservedIpRanges,omitempty"`   // DHCP reserved IP ranges
	Subnet             string                                                                `json:"subnet,omitempty"`             // Subnet of the route
}
type ResponseApplianceCreateNetworkApplianceStaticRouteFixedIPAssignments interface{}
type ResponseApplianceCreateNetworkApplianceStaticRouteReservedIPRanges struct {
	Comment string `json:"comment,omitempty"` // Description of the range
	End     string `json:"end,omitempty"`     // Last address in the reserved range
	Start   string `json:"start,omitempty"`   // First address in the reserved range
}
type ResponseApplianceGetNetworkApplianceStaticRoute struct {
	Enabled            *bool                                                              `json:"enabled,omitempty"`            // Whether the route is enabled or not
	FixedIPAssignments *ResponseApplianceGetNetworkApplianceStaticRouteFixedIPAssignments `json:"fixedIpAssignments,omitempty"` // Fixed DHCP IP assignments on the route
	GatewayIP          string                                                             `json:"gatewayIp,omitempty"`          // Gateway IP address (next hop)
	GatewayVLANID      *int                                                               `json:"gatewayVlanId,omitempty"`      // Gateway VLAN ID
	ID                 string                                                             `json:"id,omitempty"`                 // Route ID
	IPVersion          *int                                                               `json:"ipVersion,omitempty"`          // IP protocol version
	Name               string                                                             `json:"name,omitempty"`               // Name of the route
	NetworkID          string                                                             `json:"networkId,omitempty"`          // Network ID
	ReservedIPRanges   *[]ResponseApplianceGetNetworkApplianceStaticRouteReservedIPRanges `json:"reservedIpRanges,omitempty"`   // DHCP reserved IP ranges
	Subnet             string                                                             `json:"subnet,omitempty"`             // Subnet of the route
}
type ResponseApplianceGetNetworkApplianceStaticRouteFixedIPAssignments struct {
	Status223344556677 *ResponseApplianceGetNetworkApplianceStaticRouteFixedIPAssignments223344556677 `json:"22:33:44:55:66:77,omitempty"` //
}
type ResponseApplianceGetNetworkApplianceStaticRouteFixedIPAssignments223344556677 struct {
	IP   string `json:"ip,omitempty"`   //
	Name string `json:"name,omitempty"` //
}
type ResponseApplianceGetNetworkApplianceStaticRouteReservedIPRanges struct {
	Comment string `json:"comment,omitempty"` // Description of the range
	End     string `json:"end,omitempty"`     // Last address in the reserved range
	Start   string `json:"start,omitempty"`   // First address in the reserved range
}
type ResponseApplianceUpdateNetworkApplianceStaticRoute struct {
	Enabled            *bool                                                                 `json:"enabled,omitempty"`            // Whether the route is enabled or not
	FixedIPAssignments *ResponseApplianceUpdateNetworkApplianceStaticRouteFixedIPAssignments `json:"fixedIpAssignments,omitempty"` // Fixed DHCP IP assignments on the route
	GatewayIP          string                                                                `json:"gatewayIp,omitempty"`          // Gateway IP address (next hop)
	GatewayVLANID      *int                                                                  `json:"gatewayVlanId,omitempty"`      // Gateway VLAN ID
	ID                 string                                                                `json:"id,omitempty"`                 // Route ID
	IPVersion          *int                                                                  `json:"ipVersion,omitempty"`          // IP protocol version
	Name               string                                                                `json:"name,omitempty"`               // Name of the route
	NetworkID          string                                                                `json:"networkId,omitempty"`          // Network ID
	ReservedIPRanges   *[]ResponseApplianceUpdateNetworkApplianceStaticRouteReservedIPRanges `json:"reservedIpRanges,omitempty"`   // DHCP reserved IP ranges
	Subnet             string                                                                `json:"subnet,omitempty"`             // Subnet of the route
}
type ResponseApplianceUpdateNetworkApplianceStaticRouteFixedIPAssignments interface{}
type ResponseApplianceUpdateNetworkApplianceStaticRouteReservedIPRanges struct {
	Comment string `json:"comment,omitempty"` // Description of the range
	End     string `json:"end,omitempty"`     // Last address in the reserved range
	Start   string `json:"start,omitempty"`   // First address in the reserved range
}
type ResponseApplianceGetNetworkApplianceTrafficShaping struct {
	GlobalBandwidthLimits *ResponseApplianceGetNetworkApplianceTrafficShapingGlobalBandwidthLimits `json:"globalBandwidthLimits,omitempty"` //
}
type ResponseApplianceGetNetworkApplianceTrafficShapingGlobalBandwidthLimits struct {
	LimitDown *int `json:"limitDown,omitempty"` //
	LimitUp   *int `json:"limitUp,omitempty"`   //
}
type ResponseApplianceUpdateNetworkApplianceTrafficShaping interface{}
type ResponseApplianceGetNetworkApplianceTrafficShapingCustomPerformanceClasses []ResponseItemApplianceGetNetworkApplianceTrafficShapingCustomPerformanceClasses // Array of ResponseApplianceGetNetworkApplianceTrafficShapingCustomPerformanceClasses
type ResponseItemApplianceGetNetworkApplianceTrafficShapingCustomPerformanceClasses struct {
	CustomPerformanceClassID string `json:"customPerformanceClassId,omitempty"` // ID of the custom performance class
	MaxJitter                *int   `json:"maxJitter,omitempty"`                // Maximum jitter in milliseconds
	MaxLatency               *int   `json:"maxLatency,omitempty"`               // Maximum latency in milliseconds
	MaxLossPercentage        *int   `json:"maxLossPercentage,omitempty"`        // Maximum percentage of packet loss
	Name                     string `json:"name,omitempty"`                     // Name of the custom performance class
}
type ResponseApplianceCreateNetworkApplianceTrafficShapingCustomPerformanceClass struct {
	CustomPerformanceClassID string `json:"customPerformanceClassId,omitempty"` // ID of the custom performance class
	MaxJitter                *int   `json:"maxJitter,omitempty"`                // Maximum jitter in milliseconds
	MaxLatency               *int   `json:"maxLatency,omitempty"`               // Maximum latency in milliseconds
	MaxLossPercentage        *int   `json:"maxLossPercentage,omitempty"`        // Maximum percentage of packet loss
	Name                     string `json:"name,omitempty"`                     // Name of the custom performance class
}
type ResponseApplianceGetNetworkApplianceTrafficShapingCustomPerformanceClass struct {
	CustomPerformanceClassID string `json:"customPerformanceClassId,omitempty"` // ID of the custom performance class
	MaxJitter                *int   `json:"maxJitter,omitempty"`                // Maximum jitter in milliseconds
	MaxLatency               *int   `json:"maxLatency,omitempty"`               // Maximum latency in milliseconds
	MaxLossPercentage        *int   `json:"maxLossPercentage,omitempty"`        // Maximum percentage of packet loss
	Name                     string `json:"name,omitempty"`                     // Name of the custom performance class
}
type ResponseApplianceUpdateNetworkApplianceTrafficShapingCustomPerformanceClass struct {
	CustomPerformanceClassID string `json:"customPerformanceClassId,omitempty"` // ID of the custom performance class
	MaxJitter                *int   `json:"maxJitter,omitempty"`                // Maximum jitter in milliseconds
	MaxLatency               *int   `json:"maxLatency,omitempty"`               // Maximum latency in milliseconds
	MaxLossPercentage        *int   `json:"maxLossPercentage,omitempty"`        // Maximum percentage of packet loss
	Name                     string `json:"name,omitempty"`                     // Name of the custom performance class
}
type ResponseApplianceGetNetworkApplianceTrafficShapingRules struct {
	DefaultRulesEnabled *bool                                                           `json:"defaultRulesEnabled,omitempty"` //
	Rules               *[]ResponseApplianceGetNetworkApplianceTrafficShapingRulesRules `json:"rules,omitempty"`               //
}
type ResponseApplianceGetNetworkApplianceTrafficShapingRulesRules struct {
	Definitions              *[]ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesDefinitions            `json:"definitions,omitempty"`              //
	DscpTagValue             *int                                                                                  `json:"dscpTagValue,omitempty"`             //
	PerClientBandwidthLimits *ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesPerClientBandwidthLimits `json:"perClientBandwidthLimits,omitempty"` //
	Priority                 string                                                                                `json:"priority,omitempty"`                 //
}
type ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesDefinitions struct {
	Type      string                                                                    `json:"type,omitempty"` //
	Value     *string                                                                   //
	ValueObj  *ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesRulesValueObj //
	ValueList *[]string                                                                 //
}

func (r *ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesDefinitions) UnmarshalJSON(data []byte) error {
	type Alias ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesDefinitions
	aux := &struct {
		Value interface{} `json:"value"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	switch v := aux.Value.(type) {
	case string:
		r.Value = &v
	case []interface{}:
		strList := make([]string, len(v))
		for i, item := range v {
			strList[i] = item.(string)
		}
		r.ValueList = &strList
	case map[string]interface{}:
		valueObj := &ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesRulesValueObj{
			ID:   v["id"].(string),
			Name: v["name"].(string),
		}
		r.ValueObj = valueObj
	}
	return nil
}

type ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesPerClientBandwidthLimits struct {
	BandwidthLimits *ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesPerClientBandwidthLimitsBandwidthLimits `json:"bandwidthLimits,omitempty"` //
	Settings        string                                                                                               `json:"settings,omitempty"`        //
}
type ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesPerClientBandwidthLimitsBandwidthLimits struct {
	LimitDown *int `json:"limitDown,omitempty"` //
	LimitUp   *int `json:"limitUp,omitempty"`   //
}
type ResponseApplianceUpdateNetworkApplianceTrafficShapingRules interface{}
type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidth struct {
	BandwidthLimits *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimits `json:"bandwidthLimits,omitempty"` // A hash uplink keys and their configured settings for the Appliance
}
type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimits struct {
	Cellular *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular `json:"cellular,omitempty"` // uplink cellular configured limits [optional]
	Wan1     *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan1     `json:"wan1,omitempty"`     // uplink wan1 configured limits [optional]
	Wan2     *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2     `json:"wan2,omitempty"`     // uplink wan2 configured limits [optional]
}
type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular struct {
	LimitDown *int `json:"limitDown,omitempty"` // configured DOWN limit for the uplink (in Kbps).  Null indicated unlimited
	LimitUp   *int `json:"limitUp,omitempty"`   // configured UP limit for the uplink (in Kbps).  Null indicated unlimited
}
type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan1 struct {
	LimitDown *int `json:"limitDown,omitempty"` // configured DOWN limit for the uplink (in Kbps).  Null indicated unlimited
	LimitUp   *int `json:"limitUp,omitempty"`   // configured UP limit for the uplink (in Kbps).  Null indicated unlimited
}
type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2 struct {
	LimitDown *int `json:"limitDown,omitempty"` // configured DOWN limit for the uplink (in Kbps).  Null indicated unlimited
	LimitUp   *int `json:"limitUp,omitempty"`   // configured UP limit for the uplink (in Kbps).  Null indicated unlimited
}
type ResponseApplianceUpdateNetworkApplianceTrafficShapingUplinkBandwidth interface{}
type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelection struct {
	ActiveActiveAutoVpnEnabled  *bool                                                                                           `json:"activeActiveAutoVpnEnabled,omitempty"`  // Whether active-active AutoVPN is enabled
	DefaultUplink               string                                                                                          `json:"defaultUplink,omitempty"`               // The default uplink. Must be one of: 'wan1' or 'wan2'
	FailoverAndFailback         *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionFailoverAndFailback           `json:"failoverAndFailback,omitempty"`         // WAN failover and failback
	LoadBalancingEnabled        *bool                                                                                           `json:"loadBalancingEnabled,omitempty"`        // Whether load balancing is enabled
	VpnTrafficUplinkPreferences *[]ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences `json:"vpnTrafficUplinkPreferences,omitempty"` // Uplink preference rules for VPN traffic
	WanTrafficUplinkPreferences *[]ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences `json:"wanTrafficUplinkPreferences,omitempty"` // Uplink preference rules for WAN traffic
}
type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionFailoverAndFailback struct {
	Immediate *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate `json:"immediate,omitempty"` // Immediate WAN failover and failback
}
type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate struct {
	Enabled *bool `json:"enabled,omitempty"` // Whether immediate WAN failover and failback is enabled
}
type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences struct {
	FailOverCriterion string                                                                                                        `json:"failOverCriterion,omitempty"` // Fail over criterion for uplink preference rule. Must be one of: 'poorPerformance' or 'uplinkDown'
	PerformanceClass  *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesPerformanceClass `json:"performanceClass,omitempty"`  // Performance class setting for uplink preference rule
	PreferredUplink   string                                                                                                        `json:"preferredUplink,omitempty"`   // Preferred uplink for uplink preference rule. Must be one of: 'wan1', 'wan2', 'bestForVoIP', 'loadBalancing' or 'defaultUplink'
	TrafficFilters    *[]ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFilters `json:"trafficFilters,omitempty"`    // Traffic filters
}
type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesPerformanceClass struct {
	BuiltinPerformanceClassName string `json:"builtinPerformanceClassName,omitempty"` // Name of builtin performance class. Must be present when performanceClass type is 'builtin' and value must be one of: 'VoIP'
	CustomPerformanceClassID    string `json:"customPerformanceClassId,omitempty"`    // ID of created custom performance class, must be present when performanceClass type is "custom"
	Type                        string `json:"type,omitempty"`                        // Type of this performance class. Must be one of: 'builtin' or 'custom'
}
type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFilters struct {
	Type  string                                                                                                           `json:"type,omitempty"`  // Traffic filter type. Must be one of: 'applicationCategory', 'application' or 'custom'
	Value *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValue `json:"value,omitempty"` // Value of traffic filter
}
type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValue struct {
	Destination *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValueDestination `json:"destination,omitempty"` // Destination of 'custom' type traffic filter
	ID          string                                                                                                                      `json:"id,omitempty"`          // ID of 'applicationCategory' or 'application' type traffic filter
	Protocol    string                                                                                                                      `json:"protocol,omitempty"`    // Protocol of 'custom' type traffic filter. Must be one of: 'tcp', 'udp', 'icmp', 'icmp6' or 'any'
	Source      *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValueSource      `json:"source,omitempty"`      // Source of 'custom' type traffic filter
}
type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValueDestination struct {
	Cidr    string `json:"cidr,omitempty"`    // CIDR format address (e.g."192.168.10.1", which is the same as "192.168.10.1/32"), or "any". Cannot be used in combination with the "vlan" or "fqdn" property
	Fqdn    string `json:"fqdn,omitempty"`    // FQDN format address. Cannot be used in combination with the "cidr" or "fqdn" property and is currently only available in the "destination" object of the "vpnTrafficUplinkPreference" object. E.g.: "www.google.com"
	Host    *int   `json:"host,omitempty"`    // Host ID in the VLAN. Should not exceed the VLAN subnet capacity. Must be used along with the "vlan" property and is currently only available under a template network.
	Network string `json:"network,omitempty"` // Meraki network ID. Currently only available under a template network, and the value should be ID of either same template network, or another template network currently. E.g.: "L_12345678".
	Port    string `json:"port,omitempty"`    // E.g.: "any", "0" (also means "any"), "8080", "1-1024"
	VLAN    *int   `json:"vlan,omitempty"`    // VLAN ID of the configured VLAN in the Meraki network. Cannot be used in combination with the "cidr" or "fqdn" property and is currently only available under a template network.
}
type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValueSource struct {
	Cidr    string `json:"cidr,omitempty"`    // CIDR format address (e.g."192.168.10.1", which is the same as "192.168.10.1/32"), or "any". Cannot be used in combination with the "vlan" property
	Host    *int   `json:"host,omitempty"`    // Host ID in the VLAN. Should not exceed the VLAN subnet capacity. Must be used along with the "vlan" property and is currently only available under a template network.
	Network string `json:"network,omitempty"` // Meraki network ID. Currently only available under a template network, and the value should be ID of either same template network, or another template network currently. E.g.: "L_12345678".
	Port    string `json:"port,omitempty"`    // E.g.: "any", "0" (also means "any"), "8080", "1-1024"
	VLAN    *int   `json:"vlan,omitempty"`    // VLAN ID of the configured VLAN in the Meraki network. Cannot be used in combination with the "cidr" property and is currently only available under a template network.
}
type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences struct {
	PreferredUplink string                                                                                                        `json:"preferredUplink,omitempty"` // Preferred uplink for uplink preference rule. Must be one of: 'wan1' or 'wan2'
	TrafficFilters  *[]ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFilters `json:"trafficFilters,omitempty"`  // Traffic filters
}
type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFilters struct {
	Type  string                                                                                                           `json:"type,omitempty"`  // Traffic filter type. Must be "custom"
	Value *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValue `json:"value,omitempty"` // Value of traffic filter
}
type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValue struct {
	Destination *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueDestination `json:"destination,omitempty"` // Destination of 'custom' type traffic filter
	Protocol    string                                                                                                                      `json:"protocol,omitempty"`    // Protocol of 'custom' type traffic filter. Must be one of: 'tcp', 'udp', 'icmp6' or 'any'
	Source      *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueSource      `json:"source,omitempty"`      // Source of 'custom' type traffic filter
}
type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueDestination struct {
	Applications *[]ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueDestinationApplications `json:"applications,omitempty"` // list of application objects (either majorApplication or nbar)
	Cidr         string                                                                                                                                    `json:"cidr,omitempty"`         // CIDR format address (e.g."192.168.10.1", which is the same as "192.168.10.1/32"), or "any"
	Port         string                                                                                                                                    `json:"port,omitempty"`         // E.g.: "any", "0" (also means "any"), "8080", "1-1024"
}
type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueDestinationApplications struct {
	ID   string `json:"id,omitempty"`   // Id of the major application, or a list of NBAR Application Category or Application selections
	Name string `json:"name,omitempty"` // Name of the major application or application category selected
	Type string `json:"type,omitempty"` // app type (major or nbar)
}
type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueSource struct {
	Cidr string `json:"cidr,omitempty"` // CIDR format address (e.g."192.168.10.1", which is the same as "192.168.10.1/32"), or "any". Cannot be used in combination with the "vlan" property
	Host *int   `json:"host,omitempty"` // Host ID in the VLAN. Should not exceed the VLAN subnet capacity. Must be used along with the "vlan" property and is currently only available under a template network.
	Port string `json:"port,omitempty"` // E.g.: "any", "0" (also means "any"), "8080", "1-1024"
	VLAN *int   `json:"vlan,omitempty"` // VLAN ID of the configured VLAN in the Meraki network. Cannot be used in combination with the "cidr" property and is currently only available under a template network.
}
type ResponseApplianceUpdateNetworkApplianceTrafficShapingUplinkSelection struct {
	ActiveActiveAutoVpnEnabled  *bool                                                                                              `json:"activeActiveAutoVpnEnabled,omitempty"`  // Whether active-active AutoVPN is enabled
	DefaultUplink               string                                                                                             `json:"defaultUplink,omitempty"`               // The default uplink. Must be one of: 'wan1' or 'wan2'
	FailoverAndFailback         *ResponseApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionFailoverAndFailback           `json:"failoverAndFailback,omitempty"`         // WAN failover and failback
	LoadBalancingEnabled        *bool                                                                                              `json:"loadBalancingEnabled,omitempty"`        // Whether load balancing is enabled
	VpnTrafficUplinkPreferences *[]ResponseApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences `json:"vpnTrafficUplinkPreferences,omitempty"` // Uplink preference rules for VPN traffic
	WanTrafficUplinkPreferences *[]ResponseApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences `json:"wanTrafficUplinkPreferences,omitempty"` // Uplink preference rules for WAN traffic
}
type ResponseApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionFailoverAndFailback struct {
	Immediate *ResponseApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate `json:"immediate,omitempty"` // Immediate WAN failover and failback
}
type ResponseApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate struct {
	Enabled *bool `json:"enabled,omitempty"` // Whether immediate WAN failover and failback is enabled
}
type ResponseApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences struct {
	FailOverCriterion string                                                                                                           `json:"failOverCriterion,omitempty"` // Fail over criterion for uplink preference rule. Must be one of: 'poorPerformance' or 'uplinkDown'
	PerformanceClass  *ResponseApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesPerformanceClass `json:"performanceClass,omitempty"`  // Performance class setting for uplink preference rule
	PreferredUplink   string                                                                                                           `json:"preferredUplink,omitempty"`   // Preferred uplink for uplink preference rule. Must be one of: 'wan1', 'wan2', 'bestForVoIP', 'loadBalancing' or 'defaultUplink'
	TrafficFilters    *[]ResponseApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFilters `json:"trafficFilters,omitempty"`    // Traffic filters
}
type ResponseApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesPerformanceClass struct {
	BuiltinPerformanceClassName string `json:"builtinPerformanceClassName,omitempty"` // Name of builtin performance class. Must be present when performanceClass type is 'builtin' and value must be one of: 'VoIP'
	CustomPerformanceClassID    string `json:"customPerformanceClassId,omitempty"`    // ID of created custom performance class, must be present when performanceClass type is "custom"
	Type                        string `json:"type,omitempty"`                        // Type of this performance class. Must be one of: 'builtin' or 'custom'
}
type ResponseApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFilters struct {
	Type  string                                                                                                              `json:"type,omitempty"`  // Traffic filter type. Must be one of: 'applicationCategory', 'application' or 'custom'
	Value *ResponseApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValue `json:"value,omitempty"` // Value of traffic filter
}
type ResponseApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValue struct {
	Destination *ResponseApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValueDestination `json:"destination,omitempty"` // Destination of 'custom' type traffic filter
	ID          string                                                                                                                         `json:"id,omitempty"`          // ID of 'applicationCategory' or 'application' type traffic filter
	Protocol    string                                                                                                                         `json:"protocol,omitempty"`    // Protocol of 'custom' type traffic filter. Must be one of: 'tcp', 'udp', 'icmp', 'icmp6' or 'any'
	Source      *ResponseApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValueSource      `json:"source,omitempty"`      // Source of 'custom' type traffic filter
}
type ResponseApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValueDestination struct {
	Cidr    string `json:"cidr,omitempty"`    // CIDR format address (e.g."192.168.10.1", which is the same as "192.168.10.1/32"), or "any". Cannot be used in combination with the "vlan" or "fqdn" property
	Fqdn    string `json:"fqdn,omitempty"`    // FQDN format address. Cannot be used in combination with the "cidr" or "fqdn" property and is currently only available in the "destination" object of the "vpnTrafficUplinkPreference" object. E.g.: "www.google.com"
	Host    *int   `json:"host,omitempty"`    // Host ID in the VLAN. Should not exceed the VLAN subnet capacity. Must be used along with the "vlan" property and is currently only available under a template network.
	Network string `json:"network,omitempty"` // Meraki network ID. Currently only available under a template network, and the value should be ID of either same template network, or another template network currently. E.g.: "L_12345678".
	Port    string `json:"port,omitempty"`    // E.g.: "any", "0" (also means "any"), "8080", "1-1024"
	VLAN    *int   `json:"vlan,omitempty"`    // VLAN ID of the configured VLAN in the Meraki network. Cannot be used in combination with the "cidr" or "fqdn" property and is currently only available under a template network.
}
type ResponseApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValueSource struct {
	Cidr    string `json:"cidr,omitempty"`    // CIDR format address (e.g."192.168.10.1", which is the same as "192.168.10.1/32"), or "any". Cannot be used in combination with the "vlan" property
	Host    *int   `json:"host,omitempty"`    // Host ID in the VLAN. Should not exceed the VLAN subnet capacity. Must be used along with the "vlan" property and is currently only available under a template network.
	Network string `json:"network,omitempty"` // Meraki network ID. Currently only available under a template network, and the value should be ID of either same template network, or another template network currently. E.g.: "L_12345678".
	Port    string `json:"port,omitempty"`    // E.g.: "any", "0" (also means "any"), "8080", "1-1024"
	VLAN    *int   `json:"vlan,omitempty"`    // VLAN ID of the configured VLAN in the Meraki network. Cannot be used in combination with the "cidr" property and is currently only available under a template network.
}
type ResponseApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences struct {
	PreferredUplink string                                                                                                           `json:"preferredUplink,omitempty"` // Preferred uplink for uplink preference rule. Must be one of: 'wan1' or 'wan2'
	TrafficFilters  *[]ResponseApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFilters `json:"trafficFilters,omitempty"`  // Traffic filters
}
type ResponseApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFilters struct {
	Type  string                                                                                                              `json:"type,omitempty"`  // Traffic filter type. Must be "custom"
	Value *ResponseApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValue `json:"value,omitempty"` // Value of traffic filter
}
type ResponseApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValue struct {
	Destination *ResponseApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueDestination `json:"destination,omitempty"` // Destination of 'custom' type traffic filter
	Protocol    string                                                                                                                         `json:"protocol,omitempty"`    // Protocol of 'custom' type traffic filter. Must be one of: 'tcp', 'udp', 'icmp6' or 'any'
	Source      *ResponseApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueSource      `json:"source,omitempty"`      // Source of 'custom' type traffic filter
}
type ResponseApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueDestination struct {
	Applications *[]ResponseApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueDestinationApplications `json:"applications,omitempty"` // list of application objects (either majorApplication or nbar)
	Cidr         string                                                                                                                                       `json:"cidr,omitempty"`         // CIDR format address (e.g."192.168.10.1", which is the same as "192.168.10.1/32"), or "any"
	Port         string                                                                                                                                       `json:"port,omitempty"`         // E.g.: "any", "0" (also means "any"), "8080", "1-1024"
}
type ResponseApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueDestinationApplications struct {
	ID   string `json:"id,omitempty"`   // Id of the major application, or a list of NBAR Application Category or Application selections
	Name string `json:"name,omitempty"` // Name of the major application or application category selected
	Type string `json:"type,omitempty"` // app type (major or nbar)
}
type ResponseApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueSource struct {
	Cidr string `json:"cidr,omitempty"` // CIDR format address (e.g."192.168.10.1", which is the same as "192.168.10.1/32"), or "any". Cannot be used in combination with the "vlan" property
	Host *int   `json:"host,omitempty"` // Host ID in the VLAN. Should not exceed the VLAN subnet capacity. Must be used along with the "vlan" property and is currently only available under a template network.
	Port string `json:"port,omitempty"` // E.g.: "any", "0" (also means "any"), "8080", "1-1024"
	VLAN *int   `json:"vlan,omitempty"` // VLAN ID of the configured VLAN in the Meraki network. Cannot be used in combination with the "cidr" property and is currently only available under a template network.
}
type ResponseApplianceUpdateNetworkApplianceTrafficShapingVpnExclusions struct {
	Custom            *[]ResponseApplianceUpdateNetworkApplianceTrafficShapingVpnExclusionsCustom            `json:"custom,omitempty"`            // Custom VPN exclusion rules.
	MajorApplications *[]ResponseApplianceUpdateNetworkApplianceTrafficShapingVpnExclusionsMajorApplications `json:"majorApplications,omitempty"` // Major Application based VPN exclusion rules.
	NetworkID         string                                                                                 `json:"networkId,omitempty"`         // ID of the network whose VPN exclusion rules are returned.
	NetworkName       string                                                                                 `json:"networkName,omitempty"`       // Name of the network whose VPN exclusion rules are returned.
}
type ResponseApplianceUpdateNetworkApplianceTrafficShapingVpnExclusionsCustom struct {
	Destination string `json:"destination,omitempty"` // Destination address; hostname required for DNS, IPv4 otherwise.
	Port        string `json:"port,omitempty"`        // Destination port.
	Protocol    string `json:"protocol,omitempty"`    // Protocol.
}
type ResponseApplianceUpdateNetworkApplianceTrafficShapingVpnExclusionsMajorApplications struct {
	ID   string `json:"id,omitempty"`   // Application's Meraki ID.
	Name string `json:"name,omitempty"` // Application's name.
}
type ResponseApplianceGetNetworkApplianceUplinksUsageHistory []ResponseItemApplianceGetNetworkApplianceUplinksUsageHistory // Array of ResponseApplianceGetNetworkApplianceUplinksUsageHistory
type ResponseItemApplianceGetNetworkApplianceUplinksUsageHistory struct {
	ByInterface *[]ResponseItemApplianceGetNetworkApplianceUplinksUsageHistoryByInterface `json:"byInterface,omitempty"` // List of usage data for each interface
	EndTime     string                                                                    `json:"endTime,omitempty"`     // End time of interval
	StartTime   string                                                                    `json:"startTime,omitempty"`   // Start time of interval
}
type ResponseItemApplianceGetNetworkApplianceUplinksUsageHistoryByInterface struct {
	Interface string `json:"interface,omitempty"` // Uplink interface
	Received  *int   `json:"received,omitempty"`  // Amount of data received in bytes
	Sent      *int   `json:"sent,omitempty"`      // Amount of data sent in bytes
}
type ResponseApplianceGetNetworkApplianceVLANs []ResponseItemApplianceGetNetworkApplianceVLANs // Array of ResponseApplianceGetNetworkApplianceVlans
type ResponseItemApplianceGetNetworkApplianceVLANs struct {
	ApplianceIP            string                                                           `json:"applianceIp,omitempty"`            // The local IP of the appliance on the VLAN
	Cidr                   string                                                           `json:"cidr,omitempty"`                   // CIDR of the pool of subnets. Applicable only for template network. Each network bound to the template will automatically pick a subnet from this pool to build its own VLAN.
	DhcpBootFilename       string                                                           `json:"dhcpBootFilename,omitempty"`       // DHCP boot option for boot filename
	DhcpBootNextServer     string                                                           `json:"dhcpBootNextServer,omitempty"`     // DHCP boot option to direct boot clients to the server to load the boot file from
	DhcpBootOptionsEnabled *bool                                                            `json:"dhcpBootOptionsEnabled,omitempty"` // Use DHCP boot options specified in other properties
	DhcpHandling           string                                                           `json:"dhcpHandling,omitempty"`           // The appliance's handling of DHCP requests on this VLAN. One of: 'Run a DHCP server', 'Relay DHCP to another server' or 'Do not respond to DHCP requests'
	DhcpLeaseTime          string                                                           `json:"dhcpLeaseTime,omitempty"`          // The term of DHCP leases if the appliance is running a DHCP server on this VLAN. One of: '30 minutes', '1 hour', '4 hours', '12 hours', '1 day' or '1 week'
	DhcpOptions            *[]ResponseItemApplianceGetNetworkApplianceVLANsDhcpOptions      `json:"dhcpOptions,omitempty"`            // The list of DHCP options that will be included in DHCP responses. Each object in the list should have "code", "type", and "value" properties.
	DhcpRelayServerIPs     []string                                                         `json:"dhcpRelayServerIps,omitempty"`     // The IPs of the DHCP servers that DHCP requests should be relayed to
	DNSNameservers         string                                                           `json:"dnsNameservers,omitempty"`         // The DNS nameservers used for DHCP responses, either "upstream_dns", "google_dns", "opendns", or a newline seperated string of IP addresses or domain names
	FixedIPAssignments     *ResponseItemApplianceGetNetworkApplianceVLANsFixedIPAssignments `json:"fixedIpAssignments,omitempty"`     // The DHCP fixed IP assignments on the VLAN. This should be an object that contains mappings from MAC addresses to objects that themselves each contain "ip" and "name" string fields. See the sample request/response for more details.
	GroupPolicyID          string                                                           `json:"groupPolicyId,omitempty"`          // The id of the desired group policy to apply to the VLAN
	ID                     *int                                                             `json:"id,omitempty"`                     // The VLAN ID of the VLAN
	InterfaceID            string                                                           `json:"interfaceId,omitempty"`            // The interface ID of the VLAN
	IPv6                   *ResponseItemApplianceGetNetworkApplianceVLANsIPv6               `json:"ipv6,omitempty"`                   // IPv6 configuration on the VLAN
	MandatoryDhcp          *ResponseItemApplianceGetNetworkApplianceVLANsMandatoryDhcp      `json:"mandatoryDhcp,omitempty"`          // Mandatory DHCP will enforce that clients connecting to this VLAN must use the IP address assigned by the DHCP server. Clients who use a static IP address won't be able to associate. Only available on firmware versions 17.0 and above
	Mask                   *int                                                             `json:"mask,omitempty"`                   // Mask used for the subnet of all bound to the template networks. Applicable only for template network.
	Name                   string                                                           `json:"name,omitempty"`                   // The name of the VLAN
	ReservedIPRanges       *[]ResponseItemApplianceGetNetworkApplianceVLANsReservedIPRanges `json:"reservedIpRanges,omitempty"`       // The DHCP reserved IP ranges on the VLAN
	Subnet                 string                                                           `json:"subnet,omitempty"`                 // The subnet of the VLAN
	TemplateVLANType       string                                                           `json:"templateVlanType,omitempty"`       // Type of subnetting of the VLAN. Applicable only for template network.
	VpnNatSubnet           string                                                           `json:"vpnNatSubnet,omitempty"`           // The translated VPN subnet if VPN and VPN subnet translation are enabled on the VLAN
}
type ResponseItemApplianceGetNetworkApplianceVLANsDhcpOptions struct {
	Code  string `json:"code,omitempty"`  // The code for the DHCP option. This should be an integer between 2 and 254.
	Type  string `json:"type,omitempty"`  // The type for the DHCP option. One of: 'text', 'ip', 'hex' or 'integer'
	Value string `json:"value,omitempty"` // The value for the DHCP option
}
type ResponseItemApplianceGetNetworkApplianceVLANsFixedIPAssignments interface{}
type ResponseItemApplianceGetNetworkApplianceVLANsIPv6 struct {
	Enabled           *bool                                                                 `json:"enabled,omitempty"`           // Enable IPv6 on VLAN
	PrefixAssignments *[]ResponseItemApplianceGetNetworkApplianceVLANsIPv6PrefixAssignments `json:"prefixAssignments,omitempty"` // Prefix assignments on the VLAN
}
type ResponseItemApplianceGetNetworkApplianceVLANsIPv6PrefixAssignments struct {
	Autonomous         *bool                                                                     `json:"autonomous,omitempty"`         // Auto assign a /64 prefix from the origin to the VLAN
	Origin             *ResponseItemApplianceGetNetworkApplianceVLANsIPv6PrefixAssignmentsOrigin `json:"origin,omitempty"`             // The origin of the prefix
	StaticApplianceIP6 string                                                                    `json:"staticApplianceIp6,omitempty"` // Manual configuration of the IPv6 Appliance IP
	StaticPrefix       string                                                                    `json:"staticPrefix,omitempty"`       // Manual configuration of a /64 prefix on the VLAN
}
type ResponseItemApplianceGetNetworkApplianceVLANsIPv6PrefixAssignmentsOrigin struct {
	Interfaces []string `json:"interfaces,omitempty"` // Interfaces associated with the prefix
	Type       string   `json:"type,omitempty"`       // Type of the origin
}
type ResponseItemApplianceGetNetworkApplianceVLANsMandatoryDhcp struct {
	Enabled *bool `json:"enabled,omitempty"` // Enable Mandatory DHCP on VLAN.
}
type ResponseItemApplianceGetNetworkApplianceVLANsReservedIPRanges struct {
	Comment string `json:"comment,omitempty"` // A text comment for the reserved range
	End     string `json:"end,omitempty"`     // The last IP in the reserved range
	Start   string `json:"start,omitempty"`   // The first IP in the reserved range
}
type ResponseApplianceCreateNetworkApplianceVLAN struct {
	ApplianceIP      string                                                    `json:"applianceIp,omitempty"`      // The local IP of the appliance on the VLAN
	Cidr             string                                                    `json:"cidr,omitempty"`             // CIDR of the pool of subnets. Applicable only for template network. Each network bound to the template will automatically pick a subnet from this pool to build its own VLAN.
	GroupPolicyID    string                                                    `json:"groupPolicyId,omitempty"`    // The id of the desired group policy to apply to the VLAN
	ID               *int                                                      `json:"id,omitempty"`               // The VLAN ID of the VLAN
	InterfaceID      string                                                    `json:"interfaceId,omitempty"`      // The interface ID of the VLAN
	IPv6             *ResponseApplianceCreateNetworkApplianceVLANIPv6          `json:"ipv6,omitempty"`             // IPv6 configuration on the VLAN
	MandatoryDhcp    *ResponseApplianceCreateNetworkApplianceVLANMandatoryDhcp `json:"mandatoryDhcp,omitempty"`    // Mandatory DHCP will enforce that clients connecting to this VLAN must use the IP address assigned by the DHCP server. Clients who use a static IP address won't be able to associate. Only available on firmware versions 17.0 and above
	Mask             *int                                                      `json:"mask,omitempty"`             // Mask used for the subnet of all bound to the template networks. Applicable only for template network.
	Name             string                                                    `json:"name,omitempty"`             // The name of the VLAN
	Subnet           string                                                    `json:"subnet,omitempty"`           // The subnet of the VLAN
	TemplateVLANType string                                                    `json:"templateVlanType,omitempty"` // Type of subnetting of the VLAN. Applicable only for template network.
}
type ResponseApplianceCreateNetworkApplianceVLANIPv6 struct {
	Enabled           *bool                                                               `json:"enabled,omitempty"`           // Enable IPv6 on VLAN
	PrefixAssignments *[]ResponseApplianceCreateNetworkApplianceVLANIPv6PrefixAssignments `json:"prefixAssignments,omitempty"` // Prefix assignments on the VLAN
}
type ResponseApplianceCreateNetworkApplianceVLANIPv6PrefixAssignments struct {
	Autonomous         *bool                                                                   `json:"autonomous,omitempty"`         // Auto assign a /64 prefix from the origin to the VLAN
	Origin             *ResponseApplianceCreateNetworkApplianceVLANIPv6PrefixAssignmentsOrigin `json:"origin,omitempty"`             // The origin of the prefix
	StaticApplianceIP6 string                                                                  `json:"staticApplianceIp6,omitempty"` // Manual configuration of the IPv6 Appliance IP
	StaticPrefix       string                                                                  `json:"staticPrefix,omitempty"`       // Manual configuration of a /64 prefix on the VLAN
}
type ResponseApplianceCreateNetworkApplianceVLANIPv6PrefixAssignmentsOrigin struct {
	Interfaces []string `json:"interfaces,omitempty"` // Interfaces associated with the prefix
	Type       string   `json:"type,omitempty"`       // Type of the origin
}
type ResponseApplianceCreateNetworkApplianceVLANMandatoryDhcp struct {
	Enabled *bool `json:"enabled,omitempty"` // Enable Mandatory DHCP on VLAN.
}
type ResponseApplianceGetNetworkApplianceVLANsSettings struct {
	VLANsEnabled *bool `json:"vlansEnabled,omitempty"` // Boolean indicating whether VLANs are enabled (true) or disabled (false) for the network
}
type ResponseApplianceUpdateNetworkApplianceVLANsSettings struct {
	VLANsEnabled *bool `json:"vlansEnabled,omitempty"` // Boolean indicating whether VLANs are enabled (true) or disabled (false) for the network
}
type ResponseApplianceGetNetworkApplianceVLAN struct {
	ApplianceIP            string                                                      `json:"applianceIp,omitempty"`            // The local IP of the appliance on the VLAN
	Cidr                   string                                                      `json:"cidr,omitempty"`                   // CIDR of the pool of subnets. Applicable only for template network. Each network bound to the template will automatically pick a subnet from this pool to build its own VLAN.
	DhcpBootFilename       string                                                      `json:"dhcpBootFilename,omitempty"`       // DHCP boot option for boot filename
	DhcpBootNextServer     string                                                      `json:"dhcpBootNextServer,omitempty"`     // DHCP boot option to direct boot clients to the server to load the boot file from
	DhcpBootOptionsEnabled *bool                                                       `json:"dhcpBootOptionsEnabled,omitempty"` // Use DHCP boot options specified in other properties
	DhcpHandling           string                                                      `json:"dhcpHandling,omitempty"`           // The appliance's handling of DHCP requests on this VLAN. One of: 'Run a DHCP server', 'Relay DHCP to another server' or 'Do not respond to DHCP requests'
	DhcpLeaseTime          string                                                      `json:"dhcpLeaseTime,omitempty"`          // The term of DHCP leases if the appliance is running a DHCP server on this VLAN. One of: '30 minutes', '1 hour', '4 hours', '12 hours', '1 day' or '1 week'
	DhcpOptions            *[]ResponseApplianceGetNetworkApplianceVLANDhcpOptions      `json:"dhcpOptions,omitempty"`            // The list of DHCP options that will be included in DHCP responses. Each object in the list should have "code", "type", and "value" properties.
	DhcpRelayServerIPs     []string                                                    `json:"dhcpRelayServerIps,omitempty"`     // The IPs of the DHCP servers that DHCP requests should be relayed to
	DNSNameservers         string                                                      `json:"dnsNameservers,omitempty"`         // The DNS nameservers used for DHCP responses, either "upstream_dns", "google_dns", "opendns", or a newline seperated string of IP addresses or domain names
	FixedIPAssignments     *ResponseApplianceGetNetworkApplianceVLANFixedIPAssignments `json:"fixedIpAssignments,omitempty"`     // The DHCP fixed IP assignments on the VLAN. This should be an object that contains mappings from MAC addresses to objects that themselves each contain "ip" and "name" string fields. See the sample request/response for more details.
	GroupPolicyID          string                                                      `json:"groupPolicyId,omitempty"`          // The id of the desired group policy to apply to the VLAN
	ID                     *int                                                        `json:"id,omitempty"`                     // The VLAN ID of the VLAN
	InterfaceID            string                                                      `json:"interfaceId,omitempty"`            // The interface ID of the VLAN
	IPv6                   *ResponseApplianceGetNetworkApplianceVLANIPv6               `json:"ipv6,omitempty"`                   // IPv6 configuration on the VLAN
	MandatoryDhcp          *ResponseApplianceGetNetworkApplianceVLANMandatoryDhcp      `json:"mandatoryDhcp,omitempty"`          // Mandatory DHCP will enforce that clients connecting to this VLAN must use the IP address assigned by the DHCP server. Clients who use a static IP address won't be able to associate. Only available on firmware versions 17.0 and above
	Mask                   *int                                                        `json:"mask,omitempty"`                   // Mask used for the subnet of all bound to the template networks. Applicable only for template network.
	Name                   string                                                      `json:"name,omitempty"`                   // The name of the VLAN
	ReservedIPRanges       *[]ResponseApplianceGetNetworkApplianceVLANReservedIPRanges `json:"reservedIpRanges,omitempty"`       // The DHCP reserved IP ranges on the VLAN
	Subnet                 string                                                      `json:"subnet,omitempty"`                 // The subnet of the VLAN
	TemplateVLANType       string                                                      `json:"templateVlanType,omitempty"`       // Type of subnetting of the VLAN. Applicable only for template network.
	VpnNatSubnet           string                                                      `json:"vpnNatSubnet,omitempty"`           // The translated VPN subnet if VPN and VPN subnet translation are enabled on the VLAN
}
type ResponseApplianceGetNetworkApplianceVLANDhcpOptions struct {
	Code  string `json:"code,omitempty"`  // The code for the DHCP option. This should be an integer between 2 and 254.
	Type  string `json:"type,omitempty"`  // The type for the DHCP option. One of: 'text', 'ip', 'hex' or 'integer'
	Value string `json:"value,omitempty"` // The value for the DHCP option
}
type ResponseApplianceGetNetworkApplianceVLANFixedIPAssignments interface{}
type ResponseApplianceGetNetworkApplianceVLANIPv6 struct {
	Enabled           *bool                                                            `json:"enabled,omitempty"`           // Enable IPv6 on VLAN
	PrefixAssignments *[]ResponseApplianceGetNetworkApplianceVLANIPv6PrefixAssignments `json:"prefixAssignments,omitempty"` // Prefix assignments on the VLAN
}
type ResponseApplianceGetNetworkApplianceVLANIPv6PrefixAssignments struct {
	Autonomous         *bool                                                                `json:"autonomous,omitempty"`         // Auto assign a /64 prefix from the origin to the VLAN
	Origin             *ResponseApplianceGetNetworkApplianceVLANIPv6PrefixAssignmentsOrigin `json:"origin,omitempty"`             // The origin of the prefix
	StaticApplianceIP6 string                                                               `json:"staticApplianceIp6,omitempty"` // Manual configuration of the IPv6 Appliance IP
	StaticPrefix       string                                                               `json:"staticPrefix,omitempty"`       // Manual configuration of a /64 prefix on the VLAN
}
type ResponseApplianceGetNetworkApplianceVLANIPv6PrefixAssignmentsOrigin struct {
	Interfaces []string `json:"interfaces,omitempty"` // Interfaces associated with the prefix
	Type       string   `json:"type,omitempty"`       // Type of the origin
}
type ResponseApplianceGetNetworkApplianceVLANMandatoryDhcp struct {
	Enabled *bool `json:"enabled,omitempty"` // Enable Mandatory DHCP on VLAN.
}
type ResponseApplianceGetNetworkApplianceVLANReservedIPRanges struct {
	Comment string `json:"comment,omitempty"` // A text comment for the reserved range
	End     string `json:"end,omitempty"`     // The last IP in the reserved range
	Start   string `json:"start,omitempty"`   // The first IP in the reserved range
}
type ResponseApplianceUpdateNetworkApplianceVLAN struct {
	ApplianceIP            string                                                         `json:"applianceIp,omitempty"`            // The local IP of the appliance on the VLAN
	Cidr                   string                                                         `json:"cidr,omitempty"`                   // CIDR of the pool of subnets. Applicable only for template network. Each network bound to the template will automatically pick a subnet from this pool to build its own VLAN.
	DhcpBootFilename       string                                                         `json:"dhcpBootFilename,omitempty"`       // DHCP boot option for boot filename
	DhcpBootNextServer     string                                                         `json:"dhcpBootNextServer,omitempty"`     // DHCP boot option to direct boot clients to the server to load the boot file from
	DhcpBootOptionsEnabled *bool                                                          `json:"dhcpBootOptionsEnabled,omitempty"` // Use DHCP boot options specified in other properties
	DhcpHandling           string                                                         `json:"dhcpHandling,omitempty"`           // The appliance's handling of DHCP requests on this VLAN. One of: 'Run a DHCP server', 'Relay DHCP to another server' or 'Do not respond to DHCP requests'
	DhcpLeaseTime          string                                                         `json:"dhcpLeaseTime,omitempty"`          // The term of DHCP leases if the appliance is running a DHCP server on this VLAN. One of: '30 minutes', '1 hour', '4 hours', '12 hours', '1 day' or '1 week'
	DhcpOptions            *[]ResponseApplianceUpdateNetworkApplianceVLANDhcpOptions      `json:"dhcpOptions,omitempty"`            // The list of DHCP options that will be included in DHCP responses. Each object in the list should have "code", "type", and "value" properties.
	DhcpRelayServerIPs     []string                                                       `json:"dhcpRelayServerIps,omitempty"`     // The IPs of the DHCP servers that DHCP requests should be relayed to
	DNSNameservers         string                                                         `json:"dnsNameservers,omitempty"`         // The DNS nameservers used for DHCP responses, either "upstream_dns", "google_dns", "opendns", or a newline seperated string of IP addresses or domain names
	FixedIPAssignments     *ResponseApplianceUpdateNetworkApplianceVLANFixedIPAssignments `json:"fixedIpAssignments,omitempty"`     // The DHCP fixed IP assignments on the VLAN. This should be an object that contains mappings from MAC addresses to objects that themselves each contain "ip" and "name" string fields. See the sample request/response for more details.
	GroupPolicyID          string                                                         `json:"groupPolicyId,omitempty"`          // The id of the desired group policy to apply to the VLAN
	ID                     *int                                                           `json:"id,omitempty"`                     // The VLAN ID of the VLAN
	InterfaceID            string                                                         `json:"interfaceId,omitempty"`            // The interface ID of the VLAN
	IPv6                   *ResponseApplianceUpdateNetworkApplianceVLANIPv6               `json:"ipv6,omitempty"`                   // IPv6 configuration on the VLAN
	MandatoryDhcp          *ResponseApplianceUpdateNetworkApplianceVLANMandatoryDhcp      `json:"mandatoryDhcp,omitempty"`          // Mandatory DHCP will enforce that clients connecting to this VLAN must use the IP address assigned by the DHCP server. Clients who use a static IP address won't be able to associate. Only available on firmware versions 17.0 and above
	Mask                   *int                                                           `json:"mask,omitempty"`                   // Mask used for the subnet of all bound to the template networks. Applicable only for template network.
	Name                   string                                                         `json:"name,omitempty"`                   // The name of the VLAN
	ReservedIPRanges       *[]ResponseApplianceUpdateNetworkApplianceVLANReservedIPRanges `json:"reservedIpRanges,omitempty"`       // The DHCP reserved IP ranges on the VLAN
	Subnet                 string                                                         `json:"subnet,omitempty"`                 // The subnet of the VLAN
	TemplateVLANType       string                                                         `json:"templateVlanType,omitempty"`       // Type of subnetting of the VLAN. Applicable only for template network.
	VpnNatSubnet           string                                                         `json:"vpnNatSubnet,omitempty"`           // The translated VPN subnet if VPN and VPN subnet translation are enabled on the VLAN
}
type ResponseApplianceUpdateNetworkApplianceVLANDhcpOptions struct {
	Code  string `json:"code,omitempty"`  // The code for the DHCP option. This should be an integer between 2 and 254.
	Type  string `json:"type,omitempty"`  // The type for the DHCP option. One of: 'text', 'ip', 'hex' or 'integer'
	Value string `json:"value,omitempty"` // The value for the DHCP option
}
type ResponseApplianceUpdateNetworkApplianceVLANFixedIPAssignments interface{}
type ResponseApplianceUpdateNetworkApplianceVLANIPv6 struct {
	Enabled           *bool                                                               `json:"enabled,omitempty"`           // Enable IPv6 on VLAN
	PrefixAssignments *[]ResponseApplianceUpdateNetworkApplianceVLANIPv6PrefixAssignments `json:"prefixAssignments,omitempty"` // Prefix assignments on the VLAN
}
type ResponseApplianceUpdateNetworkApplianceVLANIPv6PrefixAssignments struct {
	Autonomous         *bool                                                                   `json:"autonomous,omitempty"`         // Auto assign a /64 prefix from the origin to the VLAN
	Origin             *ResponseApplianceUpdateNetworkApplianceVLANIPv6PrefixAssignmentsOrigin `json:"origin,omitempty"`             // The origin of the prefix
	StaticApplianceIP6 string                                                                  `json:"staticApplianceIp6,omitempty"` // Manual configuration of the IPv6 Appliance IP
	StaticPrefix       string                                                                  `json:"staticPrefix,omitempty"`       // Manual configuration of a /64 prefix on the VLAN
}
type ResponseApplianceUpdateNetworkApplianceVLANIPv6PrefixAssignmentsOrigin struct {
	Interfaces []string `json:"interfaces,omitempty"` // Interfaces associated with the prefix
	Type       string   `json:"type,omitempty"`       // Type of the origin
}
type ResponseApplianceUpdateNetworkApplianceVLANMandatoryDhcp struct {
	Enabled *bool `json:"enabled,omitempty"` // Enable Mandatory DHCP on VLAN.
}
type ResponseApplianceUpdateNetworkApplianceVLANReservedIPRanges struct {
	Comment string `json:"comment,omitempty"` // A text comment for the reserved range
	End     string `json:"end,omitempty"`     // The last IP in the reserved range
	Start   string `json:"start,omitempty"`   // The first IP in the reserved range
}
type ResponseApplianceGetNetworkApplianceVpnBgp struct {
	AsNumber      *int                                                   `json:"asNumber,omitempty"`      // The number of the Autonomous System to which the appliance belongs
	Enabled       *bool                                                  `json:"enabled,omitempty"`       // Whether BGP is enabled on the appliance
	IbgpHoldTimer *int                                                   `json:"ibgpHoldTimer,omitempty"` // The iBGP hold time in seconds
	Neighbors     *[]ResponseApplianceGetNetworkApplianceVpnBgpNeighbors `json:"neighbors,omitempty"`     // List of eBGP neighbor configurations
}
type ResponseApplianceGetNetworkApplianceVpnBgpNeighbors struct {
	AllowTransit    *bool                                                              `json:"allowTransit,omitempty"`    // Whether the appliance will advertise routes learned from other Autonomous Systems
	Authentication  *ResponseApplianceGetNetworkApplianceVpnBgpNeighborsAuthentication `json:"authentication,omitempty"`  // Authentication settings between BGP peers
	EbgpHoldTimer   *int                                                               `json:"ebgpHoldTimer,omitempty"`   // The eBGP hold time in seconds for the neighbor
	EbgpMultihop    *int                                                               `json:"ebgpMultihop,omitempty"`    // The number of hops the appliance must traverse to establish a peering relationship with the neighbor
	IP              string                                                             `json:"ip,omitempty"`              // The IPv4 address of the neighbor
	IPv6            *ResponseApplianceGetNetworkApplianceVpnBgpNeighborsIPv6           `json:"ipv6,omitempty"`            // Information regarding IPv6 address of the neighbor
	NextHopIP       string                                                             `json:"nextHopIp,omitempty"`       // The IPv4 address of the neighbor that will establish a TCP session with the appliance
	ReceiveLimit    *int                                                               `json:"receiveLimit,omitempty"`    // The maximum number of routes that the appliance can receive from the neighbor
	RemoteAsNumber  *int                                                               `json:"remoteAsNumber,omitempty"`  // Remote AS number of the neighbor
	SourceInterface string                                                             `json:"sourceInterface,omitempty"` // The output interface the appliance uses to establish a peering relationship with the neighbor
	TtlSecurity     *ResponseApplianceGetNetworkApplianceVpnBgpNeighborsTtlSecurity    `json:"ttlSecurity,omitempty"`     // Settings for BGP TTL security to protect BGP peering sessions from forged IP attacks
}
type ResponseApplianceGetNetworkApplianceVpnBgpNeighborsAuthentication struct {
	Password string `json:"password,omitempty"` // Password to configure MD5 authentication between BGP peers
}
type ResponseApplianceGetNetworkApplianceVpnBgpNeighborsIPv6 struct {
	Address string `json:"address,omitempty"` // The IPv6 address of the neighbor
}
type ResponseApplianceGetNetworkApplianceVpnBgpNeighborsTtlSecurity struct {
	Enabled *bool `json:"enabled,omitempty"` // Whether BGP TTL security is enabled
}
type ResponseApplianceUpdateNetworkApplianceVpnBgp struct {
	AsNumber      *int                                                      `json:"asNumber,omitempty"`      // The number of the Autonomous System to which the appliance belongs
	Enabled       *bool                                                     `json:"enabled,omitempty"`       // Whether BGP is enabled on the appliance
	IbgpHoldTimer *int                                                      `json:"ibgpHoldTimer,omitempty"` // The iBGP hold time in seconds
	Neighbors     *[]ResponseApplianceUpdateNetworkApplianceVpnBgpNeighbors `json:"neighbors,omitempty"`     // List of eBGP neighbor configurations
}
type ResponseApplianceUpdateNetworkApplianceVpnBgpNeighbors struct {
	AllowTransit    *bool                                                                 `json:"allowTransit,omitempty"`    // Whether the appliance will advertise routes learned from other Autonomous Systems
	Authentication  *ResponseApplianceUpdateNetworkApplianceVpnBgpNeighborsAuthentication `json:"authentication,omitempty"`  // Authentication settings between BGP peers
	EbgpHoldTimer   *int                                                                  `json:"ebgpHoldTimer,omitempty"`   // The eBGP hold time in seconds for the neighbor
	EbgpMultihop    *int                                                                  `json:"ebgpMultihop,omitempty"`    // The number of hops the appliance must traverse to establish a peering relationship with the neighbor
	IP              string                                                                `json:"ip,omitempty"`              // The IPv4 address of the neighbor
	IPv6            *ResponseApplianceUpdateNetworkApplianceVpnBgpNeighborsIPv6           `json:"ipv6,omitempty"`            // Information regarding IPv6 address of the neighbor
	NextHopIP       string                                                                `json:"nextHopIp,omitempty"`       // The IPv4 address of the neighbor that will establish a TCP session with the appliance
	ReceiveLimit    *int                                                                  `json:"receiveLimit,omitempty"`    // The maximum number of routes that the appliance can receive from the neighbor
	RemoteAsNumber  *int                                                                  `json:"remoteAsNumber,omitempty"`  // Remote AS number of the neighbor
	SourceInterface string                                                                `json:"sourceInterface,omitempty"` // The output interface the appliance uses to establish a peering relationship with the neighbor
	TtlSecurity     *ResponseApplianceUpdateNetworkApplianceVpnBgpNeighborsTtlSecurity    `json:"ttlSecurity,omitempty"`     // Settings for BGP TTL security to protect BGP peering sessions from forged IP attacks
}
type ResponseApplianceUpdateNetworkApplianceVpnBgpNeighborsAuthentication struct {
	Password string `json:"password,omitempty"` // Password to configure MD5 authentication between BGP peers
}
type ResponseApplianceUpdateNetworkApplianceVpnBgpNeighborsIPv6 struct {
	Address string `json:"address,omitempty"` // The IPv6 address of the neighbor
}
type ResponseApplianceUpdateNetworkApplianceVpnBgpNeighborsTtlSecurity struct {
	Enabled *bool `json:"enabled,omitempty"` // Whether BGP TTL security is enabled
}
type ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpn struct {
	Hubs    *[]ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnHubs    `json:"hubs,omitempty"`    // The list of VPN hubs, in order of preference.
	Mode    string                                                         `json:"mode,omitempty"`    // The site-to-site VPN mode.
	Subnet  *ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnet    `json:"subnet,omitempty"`  // Configuration of subnet features
	Subnets *[]ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnets `json:"subnets,omitempty"` // The list of subnets and their VPN presence.
}
type ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnHubs struct {
	HubID           string `json:"hubId,omitempty"`           // The network ID of the hub.
	UseDefaultRoute *bool  `json:"useDefaultRoute,omitempty"` // Indicates whether default route traffic should be sent to this hub.
}
type ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnet struct {
	Nat *ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnetNat `json:"nat,omitempty"` // Configuration of NAT for subnets
}
type ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnetNat struct {
	IsAllowed *bool `json:"isAllowed,omitempty"` // If enabled, VPN subnet translation can be used to translate any local subnets that are allowed to use the VPN into a new subnet with the same number of addresses.
}
type ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnets struct {
	LocalSubnet string                                                          `json:"localSubnet,omitempty"` // The CIDR notation subnet used within the VPN
	Nat         *ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnetsNat `json:"nat,omitempty"`         // Configuration of NAT for the subnet
	UseVpn      *bool                                                           `json:"useVpn,omitempty"`      // Indicates the presence of the subnet in the VPN
}
type ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnetsNat struct {
	Enabled      *bool  `json:"enabled,omitempty"`      // Whether or not VPN subnet translation is enabled for the subnet
	RemoteSubnet string `json:"remoteSubnet,omitempty"` // The translated subnet to be used in the VPN
}
type ResponseApplianceUpdateNetworkApplianceVpnSiteToSiteVpn struct {
	Hubs    *[]ResponseApplianceUpdateNetworkApplianceVpnSiteToSiteVpnHubs    `json:"hubs,omitempty"`    // The list of VPN hubs, in order of preference.
	Mode    string                                                            `json:"mode,omitempty"`    // The site-to-site VPN mode.
	Subnet  *ResponseApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnet    `json:"subnet,omitempty"`  // Configuration of subnet features
	Subnets *[]ResponseApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnets `json:"subnets,omitempty"` // The list of subnets and their VPN presence.
}
type ResponseApplianceUpdateNetworkApplianceVpnSiteToSiteVpnHubs struct {
	HubID           string `json:"hubId,omitempty"`           // The network ID of the hub.
	UseDefaultRoute *bool  `json:"useDefaultRoute,omitempty"` // Indicates whether default route traffic should be sent to this hub.
}
type ResponseApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnet struct {
	Nat *ResponseApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnetNat `json:"nat,omitempty"` // Configuration of NAT for subnets
}
type ResponseApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnetNat struct {
	IsAllowed *bool `json:"isAllowed,omitempty"` // If enabled, VPN subnet translation can be used to translate any local subnets that are allowed to use the VPN into a new subnet with the same number of addresses.
}
type ResponseApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnets struct {
	LocalSubnet string                                                             `json:"localSubnet,omitempty"` // The CIDR notation subnet used within the VPN
	Nat         *ResponseApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnetsNat `json:"nat,omitempty"`         // Configuration of NAT for the subnet
	UseVpn      *bool                                                              `json:"useVpn,omitempty"`      // Indicates the presence of the subnet in the VPN
}
type ResponseApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnetsNat struct {
	Enabled      *bool  `json:"enabled,omitempty"`      // Whether or not VPN subnet translation is enabled for the subnet
	RemoteSubnet string `json:"remoteSubnet,omitempty"` // The translated subnet to be used in the VPN
}
type ResponseApplianceGetNetworkApplianceWarmSpare struct {
	Enabled       *bool                                              `json:"enabled,omitempty"`       // Is the warm spare enabled
	PrimarySerial string                                             `json:"primarySerial,omitempty"` // Serial number of the primary appliance
	SpareSerial   string                                             `json:"spareSerial,omitempty"`   // Serial number of the warm spare appliance
	UplinkMode    string                                             `json:"uplinkMode,omitempty"`    // Uplink mode, either virtual or public
	Wan1          *ResponseApplianceGetNetworkApplianceWarmSpareWan1 `json:"wan1,omitempty"`          // WAN 1 IP and subnet
	Wan2          *ResponseApplianceGetNetworkApplianceWarmSpareWan2 `json:"wan2,omitempty"`          // WAN 2 IP and subnet
}
type ResponseApplianceGetNetworkApplianceWarmSpareWan1 struct {
	IP     string `json:"ip,omitempty"`     // IP address used for WAN 1
	Subnet string `json:"subnet,omitempty"` // Subnet used for WAN 1
}
type ResponseApplianceGetNetworkApplianceWarmSpareWan2 struct {
	IP     string `json:"ip,omitempty"`     // IP address used for WAN 2
	Subnet string `json:"subnet,omitempty"` // Subnet used for WAN 2
}
type ResponseApplianceUpdateNetworkApplianceWarmSpare struct {
	Enabled       *bool                                                 `json:"enabled,omitempty"`       // Is the warm spare enabled
	PrimarySerial string                                                `json:"primarySerial,omitempty"` // Serial number of the primary appliance
	SpareSerial   string                                                `json:"spareSerial,omitempty"`   // Serial number of the warm spare appliance
	UplinkMode    string                                                `json:"uplinkMode,omitempty"`    // Uplink mode, either virtual or public
	Wan1          *ResponseApplianceUpdateNetworkApplianceWarmSpareWan1 `json:"wan1,omitempty"`          // WAN 1 IP and subnet
	Wan2          *ResponseApplianceUpdateNetworkApplianceWarmSpareWan2 `json:"wan2,omitempty"`          // WAN 2 IP and subnet
}
type ResponseApplianceUpdateNetworkApplianceWarmSpareWan1 struct {
	IP     string `json:"ip,omitempty"`     // IP address used for WAN 1
	Subnet string `json:"subnet,omitempty"` // Subnet used for WAN 1
}
type ResponseApplianceUpdateNetworkApplianceWarmSpareWan2 struct {
	IP     string `json:"ip,omitempty"`     // IP address used for WAN 2
	Subnet string `json:"subnet,omitempty"` // Subnet used for WAN 2
}
type ResponseApplianceSwapNetworkApplianceWarmSpare struct {
	Enabled       *bool                                               `json:"enabled,omitempty"`       // Is the warm spare enabled
	PrimarySerial string                                              `json:"primarySerial,omitempty"` // Serial number of the primary appliance
	SpareSerial   string                                              `json:"spareSerial,omitempty"`   // Serial number of the warm spare appliance
	UplinkMode    string                                              `json:"uplinkMode,omitempty"`    // Uplink mode, either virtual or public
	Wan1          *ResponseApplianceSwapNetworkApplianceWarmSpareWan1 `json:"wan1,omitempty"`          // WAN 1 IP and subnet
	Wan2          *ResponseApplianceSwapNetworkApplianceWarmSpareWan2 `json:"wan2,omitempty"`          // WAN 2 IP and subnet
}
type ResponseApplianceSwapNetworkApplianceWarmSpareWan1 struct {
	IP     string `json:"ip,omitempty"`     // IP address used for WAN 1
	Subnet string `json:"subnet,omitempty"` // Subnet used for WAN 1
}
type ResponseApplianceSwapNetworkApplianceWarmSpareWan2 struct {
	IP     string `json:"ip,omitempty"`     // IP address used for WAN 2
	Subnet string `json:"subnet,omitempty"` // Subnet used for WAN 2
}
type ResponseApplianceGetOrganizationApplianceDNSLocalProfiles []ResponseItemApplianceGetOrganizationApplianceDNSLocalProfiles // Array of ResponseApplianceGetOrganizationApplianceDnsLocalProfiles
type ResponseItemApplianceGetOrganizationApplianceDNSLocalProfiles struct {
	Name      string `json:"name,omitempty"`      // Name of profile
	ProfileID string `json:"profileId,omitempty"` // Profile ID
}
type ResponseApplianceCreateOrganizationApplianceDNSLocalProfile struct {
	Name      string `json:"name,omitempty"`      // Name of profile
	ProfileID string `json:"profileId,omitempty"` // Profile ID
}
type ResponseApplianceGetOrganizationApplianceDNSLocalProfilesAssignments struct {
	Items *[]ResponseApplianceGetOrganizationApplianceDNSLocalProfilesAssignmentsItems `json:"items,omitempty"` // List of local DNS profile assignment
	Meta  *ResponseApplianceGetOrganizationApplianceDNSLocalProfilesAssignmentsMeta    `json:"meta,omitempty"`  // Metadata relevant to the paginated dataset
}
type ResponseApplianceGetOrganizationApplianceDNSLocalProfilesAssignmentsItems struct {
	AssignmentID string                                                                            `json:"assignmentId,omitempty"` // ID of the assignment
	Network      *ResponseApplianceGetOrganizationApplianceDNSLocalProfilesAssignmentsItemsNetwork `json:"network,omitempty"`      // The network attached to the profile
	Profile      *ResponseApplianceGetOrganizationApplianceDNSLocalProfilesAssignmentsItemsProfile `json:"profile,omitempty"`      // The profile the network is attached to
}
type ResponseApplianceGetOrganizationApplianceDNSLocalProfilesAssignmentsItemsNetwork struct {
	ID string `json:"id,omitempty"` // ID of the network
}
type ResponseApplianceGetOrganizationApplianceDNSLocalProfilesAssignmentsItemsProfile struct {
	ID string `json:"id,omitempty"` // ID of the profile
}
type ResponseApplianceGetOrganizationApplianceDNSLocalProfilesAssignmentsMeta struct {
	Counts *ResponseApplianceGetOrganizationApplianceDNSLocalProfilesAssignmentsMetaCounts `json:"counts,omitempty"` // Counts relating to the paginated dataset
}
type ResponseApplianceGetOrganizationApplianceDNSLocalProfilesAssignmentsMetaCounts struct {
	Items *ResponseApplianceGetOrganizationApplianceDNSLocalProfilesAssignmentsMetaCountsItems `json:"items,omitempty"` // Counts relating to the paginated items
}
type ResponseApplianceGetOrganizationApplianceDNSLocalProfilesAssignmentsMetaCountsItems struct {
	Remaining *int `json:"remaining,omitempty"` // The number of items in the dataset that are available on subsequent pages
	Total     *int `json:"total,omitempty"`     // The total number of items in the dataset
}
type ResponseApplianceBulkOrganizationApplianceDNSLocalProfilesAssignmentsCreate struct {
	Items *[]ResponseApplianceBulkOrganizationApplianceDNSLocalProfilesAssignmentsCreateItems `json:"items,omitempty"` // List of local DNS profile assignment
}
type ResponseApplianceBulkOrganizationApplianceDNSLocalProfilesAssignmentsCreateItems struct {
	AssignmentID string                                                                                   `json:"assignmentId,omitempty"` // ID of the assignment
	Network      *ResponseApplianceBulkOrganizationApplianceDNSLocalProfilesAssignmentsCreateItemsNetwork `json:"network,omitempty"`      // The network attached to the profile
	Profile      *ResponseApplianceBulkOrganizationApplianceDNSLocalProfilesAssignmentsCreateItemsProfile `json:"profile,omitempty"`      // The profile the network is attached to
}
type ResponseApplianceBulkOrganizationApplianceDNSLocalProfilesAssignmentsCreateItemsNetwork struct {
	ID string `json:"id,omitempty"` // ID of the network
}
type ResponseApplianceBulkOrganizationApplianceDNSLocalProfilesAssignmentsCreateItemsProfile struct {
	ID string `json:"id,omitempty"` // ID of the profile
}
type ResponseApplianceCreateOrganizationApplianceDNSLocalProfilesAssignmentsBulkDelete struct {
	Items *[]ResponseApplianceCreateOrganizationApplianceDNSLocalProfilesAssignmentsBulkDeleteItems `json:"items,omitempty"` // List of local DNS profile assignment
}
type ResponseApplianceCreateOrganizationApplianceDNSLocalProfilesAssignmentsBulkDeleteItems struct {
	AssignmentID string                                                                                         `json:"assignmentId,omitempty"` // ID of the assignment
	Network      *ResponseApplianceCreateOrganizationApplianceDNSLocalProfilesAssignmentsBulkDeleteItemsNetwork `json:"network,omitempty"`      // The network attached to the profile
	Profile      *ResponseApplianceCreateOrganizationApplianceDNSLocalProfilesAssignmentsBulkDeleteItemsProfile `json:"profile,omitempty"`      // The profile the network is attached to
}
type ResponseApplianceCreateOrganizationApplianceDNSLocalProfilesAssignmentsBulkDeleteItemsNetwork struct {
	ID string `json:"id,omitempty"` // ID of the network
}
type ResponseApplianceCreateOrganizationApplianceDNSLocalProfilesAssignmentsBulkDeleteItemsProfile struct {
	ID string `json:"id,omitempty"` // ID of the profile
}
type ResponseApplianceUpdateOrganizationApplianceDNSLocalProfile struct {
	Name      string `json:"name,omitempty"`      // Name of profile
	ProfileID string `json:"profileId,omitempty"` // Profile ID
}
type ResponseApplianceGetOrganizationApplianceDNSLocalRecords []ResponseItemApplianceGetOrganizationApplianceDNSLocalRecords // Array of ResponseApplianceGetOrganizationApplianceDnsLocalRecords
type ResponseItemApplianceGetOrganizationApplianceDNSLocalRecords struct {
	Address  string                                                               `json:"address,omitempty"`  // IP for the DNS record
	Hostname string                                                               `json:"hostname,omitempty"` // Hostname for the DNS record
	Profile  *ResponseItemApplianceGetOrganizationApplianceDNSLocalRecordsProfile `json:"profile,omitempty"`  // The profile the DNS record is associated with
	RecordID string                                                               `json:"recordId,omitempty"` // Record ID
}
type ResponseItemApplianceGetOrganizationApplianceDNSLocalRecordsProfile struct {
	ID string `json:"id,omitempty"` // Profile ID
}
type ResponseApplianceCreateOrganizationApplianceDNSLocalRecord []ResponseItemApplianceCreateOrganizationApplianceDNSLocalRecord // Array of ResponseApplianceCreateOrganizationApplianceDnsLocalRecord
type ResponseItemApplianceCreateOrganizationApplianceDNSLocalRecord struct {
	Address  string                                                                 `json:"address,omitempty"`  // IP for the DNS record
	Hostname string                                                                 `json:"hostname,omitempty"` // Hostname for the DNS record
	Profile  *ResponseItemApplianceCreateOrganizationApplianceDNSLocalRecordProfile `json:"profile,omitempty"`  // The profile the DNS record is associated with
	RecordID string                                                                 `json:"recordId,omitempty"` // Record ID
}
type ResponseItemApplianceCreateOrganizationApplianceDNSLocalRecordProfile struct {
	ID string `json:"id,omitempty"` // Profile ID
}
type ResponseApplianceUpdateOrganizationApplianceDNSLocalRecord struct {
	Address  string                                                             `json:"address,omitempty"`  // IP for the DNS record
	Hostname string                                                             `json:"hostname,omitempty"` // Hostname for the DNS record
	Profile  *ResponseApplianceUpdateOrganizationApplianceDNSLocalRecordProfile `json:"profile,omitempty"`  // The profile the DNS record is associated with
	RecordID string                                                             `json:"recordId,omitempty"` // Record ID
}
type ResponseApplianceUpdateOrganizationApplianceDNSLocalRecordProfile struct {
	ID string `json:"id,omitempty"` // Profile ID
}
type ResponseApplianceGetOrganizationApplianceDNSSplitProfiles []ResponseItemApplianceGetOrganizationApplianceDNSSplitProfiles // Array of ResponseApplianceGetOrganizationApplianceDnsSplitProfiles
type ResponseItemApplianceGetOrganizationApplianceDNSSplitProfiles struct {
	Hostnames   []string                                                                  `json:"hostnames,omitempty"`   // The hostname patterns to match for redirection. For more information on Split DNS hostname pattern formatting, please consult the Split DNS KB.
	Name        string                                                                    `json:"name,omitempty"`        // Name of profile
	Nameservers *ResponseItemApplianceGetOrganizationApplianceDNSSplitProfilesNameservers `json:"nameservers,omitempty"` // Contains the nameserver information for redirection.
	ProfileID   string                                                                    `json:"profileId,omitempty"`   // Profile ID
}
type ResponseItemApplianceGetOrganizationApplianceDNSSplitProfilesNameservers struct {
	Addresses []string `json:"addresses,omitempty"` // The nameserver address(es) to use for redirection. A maximum of one address is supported.
}
type ResponseApplianceCreateOrganizationApplianceDNSSplitProfile struct {
	Hostnames   []string                                                                `json:"hostnames,omitempty"`   // The hostname patterns to match for redirection. For more information on Split DNS hostname pattern formatting, please consult the Split DNS KB.
	Name        string                                                                  `json:"name,omitempty"`        // Name of profile
	Nameservers *ResponseApplianceCreateOrganizationApplianceDNSSplitProfileNameservers `json:"nameservers,omitempty"` // Contains the nameserver information for redirection.
	ProfileID   string                                                                  `json:"profileId,omitempty"`   // Profile ID
}
type ResponseApplianceCreateOrganizationApplianceDNSSplitProfileNameservers struct {
	Addresses []string `json:"addresses,omitempty"` // The nameserver address(es) to use for redirection. A maximum of one address is supported.
}
type ResponseApplianceGetOrganizationApplianceDNSSplitProfilesAssignments struct {
	Items *[]ResponseApplianceGetOrganizationApplianceDNSSplitProfilesAssignmentsItems `json:"items,omitempty"` // List of split DNS profile assignment
	Meta  *ResponseApplianceGetOrganizationApplianceDNSSplitProfilesAssignmentsMeta    `json:"meta,omitempty"`  // Metadata relevant to the paginated dataset
}
type ResponseApplianceGetOrganizationApplianceDNSSplitProfilesAssignmentsItems struct {
	AssignmentID string                                                                            `json:"assignmentId,omitempty"` // ID of the assignment
	Network      *ResponseApplianceGetOrganizationApplianceDNSSplitProfilesAssignmentsItemsNetwork `json:"network,omitempty"`      // The network attached to the profile
	Profile      *ResponseApplianceGetOrganizationApplianceDNSSplitProfilesAssignmentsItemsProfile `json:"profile,omitempty"`      // The profile the network is attached to
}
type ResponseApplianceGetOrganizationApplianceDNSSplitProfilesAssignmentsItemsNetwork struct {
	ID string `json:"id,omitempty"` // ID of the network
}
type ResponseApplianceGetOrganizationApplianceDNSSplitProfilesAssignmentsItemsProfile struct {
	ID string `json:"id,omitempty"` // ID of the profile
}
type ResponseApplianceGetOrganizationApplianceDNSSplitProfilesAssignmentsMeta struct {
	Counts *ResponseApplianceGetOrganizationApplianceDNSSplitProfilesAssignmentsMetaCounts `json:"counts,omitempty"` // Counts relating to the paginated dataset
}
type ResponseApplianceGetOrganizationApplianceDNSSplitProfilesAssignmentsMetaCounts struct {
	Items *ResponseApplianceGetOrganizationApplianceDNSSplitProfilesAssignmentsMetaCountsItems `json:"items,omitempty"` // Counts relating to the paginated items
}
type ResponseApplianceGetOrganizationApplianceDNSSplitProfilesAssignmentsMetaCountsItems struct {
	Remaining *int `json:"remaining,omitempty"` // The number of items in the dataset that are available on subsequent pages
	Total     *int `json:"total,omitempty"`     // The total number of items in the dataset
}
type ResponseApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreate struct {
	Items *[]ResponseApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreateItems `json:"items,omitempty"` // List of split DNS profile assignment
}
type ResponseApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreateItems struct {
	AssignmentID string                                                                                         `json:"assignmentId,omitempty"` // ID of the assignment
	Network      *ResponseApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreateItemsNetwork `json:"network,omitempty"`      // The network attached to the profile
	Profile      *ResponseApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreateItemsProfile `json:"profile,omitempty"`      // The profile the network is attached to
}
type ResponseApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreateItemsNetwork struct {
	ID string `json:"id,omitempty"` // ID of the network
}
type ResponseApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreateItemsProfile struct {
	ID string `json:"id,omitempty"` // ID of the profile
}
type ResponseApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkDelete struct {
	Items *[]ResponseApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkDeleteItems `json:"items,omitempty"` // List of split DNS profile assignment
}
type ResponseApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkDeleteItems struct {
	AssignmentID string                                                                                         `json:"assignmentId,omitempty"` // ID of the assignment
	Network      *ResponseApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkDeleteItemsNetwork `json:"network,omitempty"`      // The network attached to the profile
	Profile      *ResponseApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkDeleteItemsProfile `json:"profile,omitempty"`      // The profile the network is attached to
}
type ResponseApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkDeleteItemsNetwork struct {
	ID string `json:"id,omitempty"` // ID of the network
}
type ResponseApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkDeleteItemsProfile struct {
	ID string `json:"id,omitempty"` // ID of the profile
}
type ResponseApplianceUpdateOrganizationApplianceDNSSplitProfile struct {
	Hostnames   []string                                                                `json:"hostnames,omitempty"`   // The hostname patterns to match for redirection. For more information on Split DNS hostname pattern formatting, please consult the Split DNS KB.
	Name        string                                                                  `json:"name,omitempty"`        // Name of profile
	Nameservers *ResponseApplianceUpdateOrganizationApplianceDNSSplitProfileNameservers `json:"nameservers,omitempty"` // Contains the nameserver information for redirection.
	ProfileID   string                                                                  `json:"profileId,omitempty"`   // Profile ID
}
type ResponseApplianceUpdateOrganizationApplianceDNSSplitProfileNameservers struct {
	Addresses []string `json:"addresses,omitempty"` // The nameserver address(es) to use for redirection. A maximum of one address is supported.
}
type ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetwork struct {
	Items *[]ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetworkItems `json:"items,omitempty"` // List of networks with multicast static forwarding rules
	Meta  *ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetworkMeta    `json:"meta,omitempty"`  // Metadata relevant to the paginated dataset
}
type ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetworkItems struct {
	Network *ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetworkItemsNetwork `json:"network,omitempty"` // Network details
	Rules   *[]ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetworkItemsRules `json:"rules,omitempty"`   // Static multicast forwarding rules.
}
type ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetworkItemsNetwork struct {
	ID   string `json:"id,omitempty"`   // ID of the network whose multicast forwarding settings are returned.
	Name string `json:"name,omitempty"` // Name of the network whose multicast forwarding settings are returned.
}
type ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetworkItemsRules struct {
	Address     string   `json:"address,omitempty"`     // IP address
	Description string   `json:"description,omitempty"` // Forwarding rule description.
	VLANIDs     []string `json:"vlanIds,omitempty"`     // List of VLAN IDs
}
type ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetworkMeta struct {
	Counts *ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetworkMetaCounts `json:"counts,omitempty"` // Counts relating to the paginated dataset
}
type ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetworkMetaCounts struct {
	Items *ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetworkMetaCountsItems `json:"items,omitempty"` // Counts relating to the paginated items
}
type ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetworkMetaCountsItems struct {
	Remaining *int `json:"remaining,omitempty"` // The number of items in the dataset that are available on subsequent pages
	Total     *int `json:"total,omitempty"`     // The total number of items in the dataset
}
type ResponseApplianceGetOrganizationApplianceSecurityEvents []ResponseItemApplianceGetOrganizationApplianceSecurityEvents // Array of ResponseApplianceGetOrganizationApplianceSecurityEvents
type ResponseItemApplianceGetOrganizationApplianceSecurityEvents struct {
	Action          string `json:"action,omitempty"`          //
	CanonicalName   string `json:"canonicalName,omitempty"`   //
	ClientIP        string `json:"clientIp,omitempty"`        //
	ClientMac       string `json:"clientMac,omitempty"`       //
	ClientName      string `json:"clientName,omitempty"`      //
	DestIP          string `json:"destIp,omitempty"`          //
	DestinationPort *int   `json:"destinationPort,omitempty"` //
	Disposition     string `json:"disposition,omitempty"`     //
	EventType       string `json:"eventType,omitempty"`       //
	FileHash        string `json:"fileHash,omitempty"`        //
	FileSizeBytes   *int   `json:"fileSizeBytes,omitempty"`   //
	FileType        string `json:"fileType,omitempty"`        //
	Protocol        string `json:"protocol,omitempty"`        //
	SrcIP           string `json:"srcIp,omitempty"`           //
	Ts              string `json:"ts,omitempty"`              //
	URI             string `json:"uri,omitempty"`             //
}
type ResponseApplianceGetOrganizationApplianceSecurityIntrusion struct {
	AllowedRules *[]ResponseApplianceGetOrganizationApplianceSecurityIntrusionAllowedRules `json:"allowedRules,omitempty"` //
}
type ResponseApplianceGetOrganizationApplianceSecurityIntrusionAllowedRules struct {
	Message string `json:"message,omitempty"` //
	RuleID  string `json:"ruleId,omitempty"`  //
}
type ResponseApplianceUpdateOrganizationApplianceSecurityIntrusion interface{}
type ResponseApplianceGetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork struct {
	Items *[]ResponseApplianceGetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkItems `json:"items,omitempty"` // VPN exclusion rules by network
}
type ResponseApplianceGetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkItems struct {
	Custom            *[]ResponseApplianceGetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkItemsCustom            `json:"custom,omitempty"`            // Custom VPN exclusion rules.
	MajorApplications *[]ResponseApplianceGetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkItemsMajorApplications `json:"majorApplications,omitempty"` // Major Application based VPN exclusion rules.
	NetworkID         string                                                                                                 `json:"networkId,omitempty"`         // ID of the network whose VPN exclusion rules are returned.
	NetworkName       string                                                                                                 `json:"networkName,omitempty"`       // Name of the network whose VPN exclusion rules are returned.
}
type ResponseApplianceGetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkItemsCustom struct {
	Destination string `json:"destination,omitempty"` // Destination address; hostname required for DNS, IPv4 otherwise.
	Port        string `json:"port,omitempty"`        // Destination port.
	Protocol    string `json:"protocol,omitempty"`    // Protocol.
}
type ResponseApplianceGetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkItemsMajorApplications struct {
	ID   string `json:"id,omitempty"`   // Application's Meraki ID.
	Name string `json:"name,omitempty"` // Application's name.
}
type ResponseApplianceGetOrganizationApplianceUplinkStatuses []ResponseItemApplianceGetOrganizationApplianceUplinkStatuses // Array of ResponseApplianceGetOrganizationApplianceUplinkStatuses
type ResponseItemApplianceGetOrganizationApplianceUplinkStatuses struct {
	HighAvailability *ResponseItemApplianceGetOrganizationApplianceUplinkStatusesHighAvailability `json:"highAvailability,omitempty"` // Device High Availability Capabilities
	LastReportedAt   string                                                                       `json:"lastReportedAt,omitempty"`   // Last reported time for the device
	Model            string                                                                       `json:"model,omitempty"`            // The uplink model
	NetworkID        string                                                                       `json:"networkId,omitempty"`        // Network identifier
	Serial           string                                                                       `json:"serial,omitempty"`           // The uplink serial
	Uplinks          *[]ResponseItemApplianceGetOrganizationApplianceUplinkStatusesUplinks        `json:"uplinks,omitempty"`          // Uplinks
}
type ResponseItemApplianceGetOrganizationApplianceUplinkStatusesHighAvailability struct {
	Enabled *bool  `json:"enabled,omitempty"` // Indicates whether High Availability is enabled for the device. For devices that do not support HA, this will be 'false'
	Role    string `json:"role,omitempty"`    // The HA role of the device on the network. For devices that do not support HA, this will be 'primary'
}
type ResponseItemApplianceGetOrganizationApplianceUplinkStatusesUplinks struct {
	Gateway      string `json:"gateway,omitempty"`      // Gateway IP
	Interface    string `json:"interface,omitempty"`    // Uplink interface
	IP           string `json:"ip,omitempty"`           // Uplink IP
	IPAssignedBy string `json:"ipAssignedBy,omitempty"` // The way in which the IP is assigned
	PrimaryDNS   string `json:"primaryDns,omitempty"`   // Primary DNS IP
	PublicIP     string `json:"publicIp,omitempty"`     // Public IP
	SecondaryDNS string `json:"secondaryDns,omitempty"` // Secondary DNS IP
	Status       string `json:"status,omitempty"`       // Uplink status
}
type ResponseApplianceGetOrganizationApplianceUplinksStatusesOverview struct {
	Counts *ResponseApplianceGetOrganizationApplianceUplinksStatusesOverviewCounts `json:"counts,omitempty"` // counts
}
type ResponseApplianceGetOrganizationApplianceUplinksStatusesOverviewCounts struct {
	ByStatus *ResponseApplianceGetOrganizationApplianceUplinksStatusesOverviewCountsByStatus `json:"byStatus,omitempty"` // byStatus
}
type ResponseApplianceGetOrganizationApplianceUplinksStatusesOverviewCountsByStatus struct {
	Active       *int `json:"active,omitempty"`       // number of uplinks that are active and working
	Connecting   *int `json:"connecting,omitempty"`   // number of uplinks currently connecting
	Failed       *int `json:"failed,omitempty"`       // number of uplinks that were working but have failed
	NotConnected *int `json:"notConnected,omitempty"` // number of uplinks currently where nothing is plugged in
	Ready        *int `json:"ready,omitempty"`        // number of uplinks that are working but on standby
}
type ResponseApplianceGetOrganizationApplianceUplinksUsageByNetwork []ResponseItemApplianceGetOrganizationApplianceUplinksUsageByNetwork // Array of ResponseApplianceGetOrganizationApplianceUplinksUsageByNetwork
type ResponseItemApplianceGetOrganizationApplianceUplinksUsageByNetwork struct {
	ByUplink  *[]ResponseItemApplianceGetOrganizationApplianceUplinksUsageByNetworkByUplink `json:"byUplink,omitempty"`  // Uplink usage
	Name      string                                                                        `json:"name,omitempty"`      // Network name
	NetworkID string                                                                        `json:"networkId,omitempty"` // Network identifier
}
type ResponseItemApplianceGetOrganizationApplianceUplinksUsageByNetworkByUplink struct {
	Interface string `json:"interface,omitempty"` // Uplink name
	Received  *int   `json:"received,omitempty"`  // Bytes received
	Sent      *int   `json:"sent,omitempty"`      // Bytes sent
	Serial    string `json:"serial,omitempty"`    // Uplink serial
}
type ResponseApplianceGetOrganizationApplianceVpnStats []ResponseItemApplianceGetOrganizationApplianceVpnStats // Array of ResponseApplianceGetOrganizationApplianceVpnStats
type ResponseItemApplianceGetOrganizationApplianceVpnStats struct {
	MerakiVpnpeers *[]ResponseItemApplianceGetOrganizationApplianceVpnStatsMerakiVpnpeers `json:"merakiVpnPeers,omitempty"` // List of VPN peers with their summaries
	NetworkID      string                                                                 `json:"networkId,omitempty"`      // Network ID
	NetworkName    string                                                                 `json:"networkName,omitempty"`    // Network name
}
type ResponseItemApplianceGetOrganizationApplianceVpnStatsMerakiVpnpeers struct {
	JitterSummaries         *[]ResponseItemApplianceGetOrganizationApplianceVpnStatsMerakiVpnpeersJitterSummaries         `json:"jitterSummaries,omitempty"`         // List of jitter summaries
	LatencySummaries        *[]ResponseItemApplianceGetOrganizationApplianceVpnStatsMerakiVpnpeersLatencySummaries        `json:"latencySummaries,omitempty"`        // List of latency summaries
	LossPercentageSummaries *[]ResponseItemApplianceGetOrganizationApplianceVpnStatsMerakiVpnpeersLossPercentageSummaries `json:"lossPercentageSummaries,omitempty"` // List of loss percentage summaries
	MosSummaries            *[]ResponseItemApplianceGetOrganizationApplianceVpnStatsMerakiVpnpeersMosSummaries            `json:"mosSummaries,omitempty"`            // List of MOS (Mean opinion score) summaries
	NetworkID               string                                                                                        `json:"networkId,omitempty"`               // Network ID of the VPN peer
	NetworkName             string                                                                                        `json:"networkName,omitempty"`             // Network name of the VPN peer
	UsageSummary            *ResponseItemApplianceGetOrganizationApplianceVpnStatsMerakiVpnpeersUsageSummary              `json:"usageSummary,omitempty"`            // Usage summary
}
type ResponseItemApplianceGetOrganizationApplianceVpnStatsMerakiVpnpeersJitterSummaries struct {
	AvgJitter      *float64 `json:"avgJitter,omitempty"`      // Average jitter
	MaxJitter      *float64 `json:"maxJitter,omitempty"`      // Maximum jitter
	MinJitter      *float64 `json:"minJitter,omitempty"`      // Minimum jitter
	ReceiverUplink string   `json:"receiverUplink,omitempty"` // Receiver uplink interface
	SenderUplink   string   `json:"senderUplink,omitempty"`   // Sender uplink interface
}
type ResponseItemApplianceGetOrganizationApplianceVpnStatsMerakiVpnpeersLatencySummaries struct {
	AvgLatencyMs   *int   `json:"avgLatencyMs,omitempty"`   // Average latency in milliseconds
	MaxLatencyMs   *int   `json:"maxLatencyMs,omitempty"`   // Maximum latency in milliseconds
	MinLatencyMs   *int   `json:"minLatencyMs,omitempty"`   // Minimum latency in milliseconds
	ReceiverUplink string `json:"receiverUplink,omitempty"` // Receiver uplink interface
	SenderUplink   string `json:"senderUplink,omitempty"`   // Sender uplink interface
}
type ResponseItemApplianceGetOrganizationApplianceVpnStatsMerakiVpnpeersLossPercentageSummaries struct {
	AvgLossPercentage *float64 `json:"avgLossPercentage,omitempty"` // Average loss percentage
	MaxLossPercentage *float64 `json:"maxLossPercentage,omitempty"` // Maximum loss percentage
	MinLossPercentage *float64 `json:"minLossPercentage,omitempty"` // Minimum loss percentage
	ReceiverUplink    string   `json:"receiverUplink,omitempty"`    // Receiver uplink interface
	SenderUplink      string   `json:"senderUplink,omitempty"`      // Sender uplink interface
}
type ResponseItemApplianceGetOrganizationApplianceVpnStatsMerakiVpnpeersMosSummaries struct {
	AvgMos         *float64 `json:"avgMos,omitempty"`         // Average MOS (Mean opinion score)
	MaxMos         *float64 `json:"maxMos,omitempty"`         // Maximum MOS (Mean opinion score
	MinMos         *float64 `json:"minMos,omitempty"`         // Minimum MOS (Mean opinion score
	ReceiverUplink string   `json:"receiverUplink,omitempty"` // Receiver uplink interface
	SenderUplink   string   `json:"senderUplink,omitempty"`   // Sender uplink interface
}
type ResponseItemApplianceGetOrganizationApplianceVpnStatsMerakiVpnpeersUsageSummary struct {
	ReceivedInKilobytes *int `json:"receivedInKilobytes,omitempty"` // The amount of data received (in kilobytes)
	SentInKilobytes     *int `json:"sentInKilobytes,omitempty"`     // The amount of data sent (in kilobytes)
}
type ResponseApplianceGetOrganizationApplianceVpnStatuses struct {
	Vpnstatusentities *[]ResponseApplianceGetOrganizationApplianceVpnStatusesVpnstatusentities `json:"vpnstatusentities,omitempty"` // The list of VPN Status for networks
}
type ResponseApplianceGetOrganizationApplianceVpnStatusesVpnstatusentities struct {
	DeviceSerial       string                                                                                     `json:"deviceSerial,omitempty"`       // Serial number of the device
	DeviceStatus       string                                                                                     `json:"deviceStatus,omitempty"`       // Device Status
	ExportedSubnets    *[]ResponseApplianceGetOrganizationApplianceVpnStatusesVpnstatusentitiesExportedSubnets    `json:"exportedSubnets,omitempty"`    // List of Exported Subnets
	MerakiVpnpeers     *[]ResponseApplianceGetOrganizationApplianceVpnStatusesVpnstatusentitiesMerakiVpnpeers     `json:"merakiVpnPeers,omitempty"`     // Meraki VPN Peers
	NetworkID          string                                                                                     `json:"networkId,omitempty"`          // Network Id
	NetworkName        string                                                                                     `json:"networkName,omitempty"`        // Network name
	ThirdPartyVpnpeers *[]ResponseApplianceGetOrganizationApplianceVpnStatusesVpnstatusentitiesThirdPartyVpnpeers `json:"thirdPartyVpnPeers,omitempty"` // Third Party VPN Peers
	Uplinks            *[]ResponseApplianceGetOrganizationApplianceVpnStatusesVpnstatusentitiesUplinks            `json:"uplinks,omitempty"`            // List of Uplink Information
	VpnMode            string                                                                                     `json:"vpnMode,omitempty"`            // VPN Mode
}
type ResponseApplianceGetOrganizationApplianceVpnStatusesVpnstatusentitiesExportedSubnets struct {
	Name   string `json:"name,omitempty"`   // Name of the subnet
	Subnet string `json:"subnet,omitempty"` // Subnet
}
type ResponseApplianceGetOrganizationApplianceVpnStatusesVpnstatusentitiesMerakiVpnpeers struct {
	NetworkID    string `json:"networkId,omitempty"`    // Network ID
	NetworkName  string `json:"networkName,omitempty"`  // Network Name
	Reachability string `json:"reachability,omitempty"` // Reachability
}
type ResponseApplianceGetOrganizationApplianceVpnStatusesVpnstatusentitiesThirdPartyVpnpeers struct {
	Name         string `json:"name,omitempty"`         // Name of the peer
	PublicIP     string `json:"publicIp,omitempty"`     // Public IP of the peer
	Reachability string `json:"reachability,omitempty"` // Reachability
}
type ResponseApplianceGetOrganizationApplianceVpnStatusesVpnstatusentitiesUplinks struct {
	Interface string `json:"interface,omitempty"` // Uplink Interface Name
	PublicIP  string `json:"publicIp,omitempty"`  // Uplink IP address (in IP or CIDR notation)
}
type ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeers struct {
	Peers *[]ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeersPeers `json:"peers,omitempty"` // The list of VPN peers
}
type ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeersPeers struct {
	IkeVersion          string                                                                            `json:"ikeVersion,omitempty"`          // [optional] The IKE version to be used for the IPsec VPN peer configuration. Defaults to '1' when omitted.
	IPsecPolicies       *ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeersPeersIPsecPolicies `json:"ipsecPolicies,omitempty"`       // Custom IPSec policies for the VPN peer. If not included and a preset has not been chosen, the default preset for IPSec policies will be used.
	IPsecPoliciesPreset string                                                                            `json:"ipsecPoliciesPreset,omitempty"` // One of the following available presets: 'default', 'aws', 'azure', 'umbrella', 'zscaler'. If this is provided, the 'ipsecPolicies' parameter is ignored.
	LocalID             string                                                                            `json:"localId,omitempty"`             // [optional] The local ID is used to identify the MX to the peer. This will apply to all MXs this peer applies to.
	Name                string                                                                            `json:"name,omitempty"`                // The name of the VPN peer
	NetworkTags         []string                                                                          `json:"networkTags,omitempty"`         // A list of network tags that will connect with this peer. Use ['all'] for all networks. Use ['none'] for no networks. If not included, the default is ['all'].
	PrivateSubnets      []string                                                                          `json:"privateSubnets,omitempty"`      // The list of the private subnets of the VPN peer
	PublicIP            string                                                                            `json:"publicIp,omitempty"`            // [optional] The public IP of the VPN peer
	RemoteID            string                                                                            `json:"remoteId,omitempty"`            // [optional] The remote ID is used to identify the connecting VPN peer. This can either be a valid IPv4 Address, FQDN or User FQDN.
	Secret              string                                                                            `json:"secret,omitempty"`              // The shared secret with the VPN peer
}
type ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeersPeersIPsecPolicies struct {
	ChildAuthAlgo         []string `json:"childAuthAlgo,omitempty"`         // This is the authentication algorithms to be used in Phase 2. The value should be an array with one of the following algorithms: 'sha256', 'sha1', 'md5'
	ChildCipherAlgo       []string `json:"childCipherAlgo,omitempty"`       // This is the cipher algorithms to be used in Phase 2. The value should be an array with one or more of the following algorithms: 'aes256', 'aes192', 'aes128', 'tripledes', 'des', 'null'
	ChildLifetime         *int     `json:"childLifetime,omitempty"`         // The lifetime of the Phase 2 SA in seconds.
	ChildPfsGroup         []string `json:"childPfsGroup,omitempty"`         // This is the Diffie-Hellman group to be used for Perfect Forward Secrecy in Phase 2. The value should be an array with one of the following values: 'disabled','group14', 'group5', 'group2', 'group1'
	IkeAuthAlgo           []string `json:"ikeAuthAlgo,omitempty"`           // This is the authentication algorithm to be used in Phase 1. The value should be an array with one of the following algorithms: 'sha256', 'sha1', 'md5'
	IkeCipherAlgo         []string `json:"ikeCipherAlgo,omitempty"`         // This is the cipher algorithm to be used in Phase 1. The value should be an array with one of the following algorithms: 'aes256', 'aes192', 'aes128', 'tripledes', 'des'
	IkeDiffieHellmanGroup []string `json:"ikeDiffieHellmanGroup,omitempty"` // This is the Diffie-Hellman group to be used in Phase 1. The value should be an array with one of the following algorithms: 'group14', 'group5', 'group2', 'group1'
	IkeLifetime           *int     `json:"ikeLifetime,omitempty"`           // The lifetime of the Phase 1 SA in seconds.
	IkePrfAlgo            []string `json:"ikePrfAlgo,omitempty"`            // [optional] This is the pseudo-random function to be used in IKE_SA. The value should be an array with one of the following algorithms: 'prfsha256', 'prfsha1', 'prfmd5', 'default'. The 'default' option can be used to default to the Authentication algorithm.
}
type ResponseApplianceUpdateOrganizationApplianceVpnThirdPartyVpnpeers struct {
	Peers *[]ResponseApplianceUpdateOrganizationApplianceVpnThirdPartyVpnpeersPeers `json:"peers,omitempty"` // The list of VPN peers
}
type ResponseApplianceUpdateOrganizationApplianceVpnThirdPartyVpnpeersPeers struct {
	IkeVersion          string                                                                               `json:"ikeVersion,omitempty"`          // [optional] The IKE version to be used for the IPsec VPN peer configuration. Defaults to '1' when omitted.
	IPsecPolicies       *ResponseApplianceUpdateOrganizationApplianceVpnThirdPartyVpnpeersPeersIPsecPolicies `json:"ipsecPolicies,omitempty"`       // Custom IPSec policies for the VPN peer. If not included and a preset has not been chosen, the default preset for IPSec policies will be used.
	IPsecPoliciesPreset string                                                                               `json:"ipsecPoliciesPreset,omitempty"` // One of the following available presets: 'default', 'aws', 'azure', 'umbrella', 'zscaler'. If this is provided, the 'ipsecPolicies' parameter is ignored.
	LocalID             string                                                                               `json:"localId,omitempty"`             // [optional] The local ID is used to identify the MX to the peer. This will apply to all MXs this peer applies to.
	Name                string                                                                               `json:"name,omitempty"`                // The name of the VPN peer
	NetworkTags         []string                                                                             `json:"networkTags,omitempty"`         // A list of network tags that will connect with this peer. Use ['all'] for all networks. Use ['none'] for no networks. If not included, the default is ['all'].
	PrivateSubnets      []string                                                                             `json:"privateSubnets,omitempty"`      // The list of the private subnets of the VPN peer
	PublicIP            string                                                                               `json:"publicIp,omitempty"`            // [optional] The public IP of the VPN peer
	RemoteID            string                                                                               `json:"remoteId,omitempty"`            // [optional] The remote ID is used to identify the connecting VPN peer. This can either be a valid IPv4 Address, FQDN or User FQDN.
	Secret              string                                                                               `json:"secret,omitempty"`              // The shared secret with the VPN peer
}
type ResponseApplianceUpdateOrganizationApplianceVpnThirdPartyVpnpeersPeersIPsecPolicies struct {
	ChildAuthAlgo         []string `json:"childAuthAlgo,omitempty"`         // This is the authentication algorithms to be used in Phase 2. The value should be an array with one of the following algorithms: 'sha256', 'sha1', 'md5'
	ChildCipherAlgo       []string `json:"childCipherAlgo,omitempty"`       // This is the cipher algorithms to be used in Phase 2. The value should be an array with one or more of the following algorithms: 'aes256', 'aes192', 'aes128', 'tripledes', 'des', 'null'
	ChildLifetime         *int     `json:"childLifetime,omitempty"`         // The lifetime of the Phase 2 SA in seconds.
	ChildPfsGroup         []string `json:"childPfsGroup,omitempty"`         // This is the Diffie-Hellman group to be used for Perfect Forward Secrecy in Phase 2. The value should be an array with one of the following values: 'disabled','group14', 'group5', 'group2', 'group1'
	IkeAuthAlgo           []string `json:"ikeAuthAlgo,omitempty"`           // This is the authentication algorithm to be used in Phase 1. The value should be an array with one of the following algorithms: 'sha256', 'sha1', 'md5'
	IkeCipherAlgo         []string `json:"ikeCipherAlgo,omitempty"`         // This is the cipher algorithm to be used in Phase 1. The value should be an array with one of the following algorithms: 'aes256', 'aes192', 'aes128', 'tripledes', 'des'
	IkeDiffieHellmanGroup []string `json:"ikeDiffieHellmanGroup,omitempty"` // This is the Diffie-Hellman group to be used in Phase 1. The value should be an array with one of the following algorithms: 'group14', 'group5', 'group2', 'group1'
	IkeLifetime           *int     `json:"ikeLifetime,omitempty"`           // The lifetime of the Phase 1 SA in seconds.
	IkePrfAlgo            []string `json:"ikePrfAlgo,omitempty"`            // [optional] This is the pseudo-random function to be used in IKE_SA. The value should be an array with one of the following algorithms: 'prfsha256', 'prfsha1', 'prfmd5', 'default'. The 'default' option can be used to default to the Authentication algorithm.
}
type ResponseApplianceGetOrganizationApplianceVpnVpnFirewallRules struct {
	Rules *[]ResponseApplianceGetOrganizationApplianceVpnVpnFirewallRulesRules `json:"rules,omitempty"` // An ordered array of the firewall rules (not including the default rule)
}
type ResponseApplianceGetOrganizationApplianceVpnVpnFirewallRulesRules struct {
	Comment       string `json:"comment,omitempty"`       // Description of the rule (optional)
	DestCidr      string `json:"destCidr,omitempty"`      // Comma-separated list of destination IP address(es) (in IP or CIDR notation), fully-qualified domain names (FQDN) or 'any'
	DestPort      string `json:"destPort,omitempty"`      // Comma-separated list of destination port(s) (integer in the range 1-65535), or 'any'
	Policy        string `json:"policy,omitempty"`        // 'allow' or 'deny' traffic specified by this rule
	Protocol      string `json:"protocol,omitempty"`      // The type of protocol (must be 'tcp', 'udp', 'icmp', 'icmp6' or 'any')
	SrcCidr       string `json:"srcCidr,omitempty"`       // Comma-separated list of source IP address(es) (in IP or CIDR notation), or 'any' (note: FQDN not supported for source addresses)
	SrcPort       string `json:"srcPort,omitempty"`       // Comma-separated list of source port(s) (integer in the range 1-65535), or 'any'
	SyslogEnabled *bool  `json:"syslogEnabled,omitempty"` // Log this rule to syslog (true or false, boolean value) - only applicable if a syslog has been configured (optional)
}
type ResponseApplianceUpdateOrganizationApplianceVpnVpnFirewallRules struct {
	Rules *[]ResponseApplianceUpdateOrganizationApplianceVpnVpnFirewallRulesRules `json:"rules,omitempty"` // An ordered array of the firewall rules (not including the default rule)
}
type ResponseApplianceUpdateOrganizationApplianceVpnVpnFirewallRulesRules struct {
	Comment       string `json:"comment,omitempty"`       // Description of the rule (optional)
	DestCidr      string `json:"destCidr,omitempty"`      // Comma-separated list of destination IP address(es) (in IP or CIDR notation), fully-qualified domain names (FQDN) or 'any'
	DestPort      string `json:"destPort,omitempty"`      // Comma-separated list of destination port(s) (integer in the range 1-65535), or 'any'
	Policy        string `json:"policy,omitempty"`        // 'allow' or 'deny' traffic specified by this rule
	Protocol      string `json:"protocol,omitempty"`      // The type of protocol (must be 'tcp', 'udp', 'icmp', 'icmp6' or 'any')
	SrcCidr       string `json:"srcCidr,omitempty"`       // Comma-separated list of source IP address(es) (in IP or CIDR notation), or 'any' (note: FQDN not supported for source addresses)
	SrcPort       string `json:"srcPort,omitempty"`       // Comma-separated list of source port(s) (integer in the range 1-65535), or 'any'
	SyslogEnabled *bool  `json:"syslogEnabled,omitempty"` // Log this rule to syslog (true or false, boolean value) - only applicable if a syslog has been configured (optional)
}
type RequestApplianceUpdateDeviceApplianceRadioSettings struct {
	FiveGhzSettings    *RequestApplianceUpdateDeviceApplianceRadioSettingsFiveGhzSettings    `json:"fiveGhzSettings,omitempty"`    // Manual radio settings for 5 GHz.
	RfProfileID        string                                                                `json:"rfProfileId,omitempty"`        // The ID of an RF profile to assign to the device. If the value of this parameter is null, the appropriate basic RF profile (indoor or outdoor) will be assigned to the device. Assigning an RF profile will clear ALL manually configured overrides on the device (channel width, channel, power).
	TwoFourGhzSettings *RequestApplianceUpdateDeviceApplianceRadioSettingsTwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"` // Manual radio settings for 2.4 GHz.
}
type RequestApplianceUpdateDeviceApplianceRadioSettingsFiveGhzSettings struct {
	Channel      *int `json:"channel,omitempty"`      // Sets a manual channel for 5 GHz. Can be '36', '40', '44', '48', '52', '56', '60', '64', '100', '104', '108', '112', '116', '120', '124', '128', '132', '136', '140', '144', '149', '153', '157', '161', '165', '169', '173' or '177' or null for using auto channel.
	ChannelWidth *int `json:"channelWidth,omitempty"` // Sets a manual channel width for 5 GHz. Can be '0', '20', '40', '80' or '160' or null for using auto channel width.
	TargetPower  *int `json:"targetPower,omitempty"`  // Set a manual target power for 5 GHz (dBm). Enter null for using auto power range.
}
type RequestApplianceUpdateDeviceApplianceRadioSettingsTwoFourGhzSettings struct {
	Channel     *int `json:"channel,omitempty"`     // Sets a manual channel for 2.4 GHz. Can be '1', '2', '3', '4', '5', '6', '7', '8', '9', '10', '11', '12', '13' or '14' or null for using auto channel.
	TargetPower *int `json:"targetPower,omitempty"` // Set a manual target power for 2.4 GHz (dBm). Enter null for using auto power range.
}
type RequestApplianceUpdateDeviceApplianceUplinksSettings struct {
	Interfaces *RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfaces `json:"interfaces,omitempty"` // Interface settings.
}
type RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfaces struct {
	Wan1 *RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1 `json:"wan1,omitempty"` // WAN 1 settings.
	Wan2 *RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2 `json:"wan2,omitempty"` // WAN 2 settings.
}
type RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1 struct {
	Enabled     *bool                                                                          `json:"enabled,omitempty"`     // Enable or disable the interface.
	Pppoe       *RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1Pppoe       `json:"pppoe,omitempty"`       // Configuration options for PPPoE.
	Svis        *RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1Svis        `json:"svis,omitempty"`        // SVI settings by protocol.
	VLANTagging *RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1VLANTagging `json:"vlanTagging,omitempty"` // VLAN tagging settings.
}
type RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1Pppoe struct {
	Authentication *RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1PppoeAuthentication `json:"authentication,omitempty"` // Settings for PPPoE Authentication.
	Enabled        *bool                                                                                  `json:"enabled,omitempty"`        // Whether PPPoE is enabled.
}
type RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1PppoeAuthentication struct {
	Enabled  *bool  `json:"enabled,omitempty"`  // Whether PPPoE authentication is enabled.
	Password string `json:"password,omitempty"` // Password for PPPoE authentication. This parameter is not returned.
	Username string `json:"username,omitempty"` // Username for PPPoE authentication.
}
type RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1Svis struct {
	IPv4 *RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv4 `json:"ipv4,omitempty"` // IPv4 settings for static/dynamic mode.
	IPv6 *RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv6 `json:"ipv6,omitempty"` // IPv6 settings for static/dynamic mode.
}
type RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv4 struct {
	Address        string                                                                                 `json:"address,omitempty"`        // IP address and subnet mask when in static mode.
	AssignmentMode string                                                                                 `json:"assignmentMode,omitempty"` // The assignment mode for this SVI. Applies only when PPPoE is disabled.
	Gateway        string                                                                                 `json:"gateway,omitempty"`        // Gateway IP address when in static mode.
	Nameservers    *RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv4Nameservers `json:"nameservers,omitempty"`    // The nameserver settings for this SVI.
}
type RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv4Nameservers struct {
	Addresses []string `json:"addresses,omitempty"` // Up to 2 nameserver addresses to use, ordered in priority from highest to lowest priority.
}
type RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv6 struct {
	Address        string                                                                                 `json:"address,omitempty"`        // Static address that will override the one(s) received by SLAAC.
	AssignmentMode string                                                                                 `json:"assignmentMode,omitempty"` // The assignment mode for this SVI. Applies only when PPPoE is disabled.
	Gateway        string                                                                                 `json:"gateway,omitempty"`        // Static gateway that will override the one received by autoconf.
	Nameservers    *RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv6Nameservers `json:"nameservers,omitempty"`    // The nameserver settings for this SVI.
}
type RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv6Nameservers struct {
	Addresses []string `json:"addresses,omitempty"` // Up to 2 nameserver addresses to use, ordered in priority from highest to lowest priority.
}
type RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1VLANTagging struct {
	Enabled *bool `json:"enabled,omitempty"` // Whether VLAN tagging is enabled.
	VLANID  *int  `json:"vlanId,omitempty"`  // The ID of the VLAN to use for VLAN tagging.
}
type RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2 struct {
	Enabled     *bool                                                                          `json:"enabled,omitempty"`     // Enable or disable the interface.
	Pppoe       *RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2Pppoe       `json:"pppoe,omitempty"`       // Configuration options for PPPoE.
	Svis        *RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2Svis        `json:"svis,omitempty"`        // SVI settings by protocol.
	VLANTagging *RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2VLANTagging `json:"vlanTagging,omitempty"` // VLAN tagging settings.
}
type RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2Pppoe struct {
	Authentication *RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2PppoeAuthentication `json:"authentication,omitempty"` // Settings for PPPoE Authentication.
	Enabled        *bool                                                                                  `json:"enabled,omitempty"`        // Whether PPPoE is enabled.
}
type RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2PppoeAuthentication struct {
	Enabled  *bool  `json:"enabled,omitempty"`  // Whether PPPoE authentication is enabled.
	Password string `json:"password,omitempty"` // Password for PPPoE authentication. This parameter is not returned.
	Username string `json:"username,omitempty"` // Username for PPPoE authentication.
}
type RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2Svis struct {
	IPv4 *RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv4 `json:"ipv4,omitempty"` // IPv4 settings for static/dynamic mode.
	IPv6 *RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv6 `json:"ipv6,omitempty"` // IPv6 settings for static/dynamic mode.
}
type RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv4 struct {
	Address        string                                                                                 `json:"address,omitempty"`        // IP address and subnet mask when in static mode.
	AssignmentMode string                                                                                 `json:"assignmentMode,omitempty"` // The assignment mode for this SVI. Applies only when PPPoE is disabled.
	Gateway        string                                                                                 `json:"gateway,omitempty"`        // Gateway IP address when in static mode.
	Nameservers    *RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv4Nameservers `json:"nameservers,omitempty"`    // The nameserver settings for this SVI.
}
type RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv4Nameservers struct {
	Addresses []string `json:"addresses,omitempty"` // Up to 2 nameserver addresses to use, ordered in priority from highest to lowest priority.
}
type RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv6 struct {
	Address        string                                                                                 `json:"address,omitempty"`        // Static address that will override the one(s) received by SLAAC.
	AssignmentMode string                                                                                 `json:"assignmentMode,omitempty"` // The assignment mode for this SVI. Applies only when PPPoE is disabled.
	Gateway        string                                                                                 `json:"gateway,omitempty"`        // Static gateway that will override the one received by autoconf.
	Nameservers    *RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv6Nameservers `json:"nameservers,omitempty"`    // The nameserver settings for this SVI.
}
type RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv6Nameservers struct {
	Addresses []string `json:"addresses,omitempty"` // Up to 2 nameserver addresses to use, ordered in priority from highest to lowest priority.
}
type RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2VLANTagging struct {
	Enabled *bool `json:"enabled,omitempty"` // Whether VLAN tagging is enabled.
	VLANID  *int  `json:"vlanId,omitempty"`  // The ID of the VLAN to use for VLAN tagging.
}
type RequestApplianceUpdateNetworkApplianceConnectivityMonitoringDestinations struct {
	Destinations *[]RequestApplianceUpdateNetworkApplianceConnectivityMonitoringDestinationsDestinations `json:"destinations,omitempty"` // The list of connectivity monitoring destinations
}
type RequestApplianceUpdateNetworkApplianceConnectivityMonitoringDestinationsDestinations struct {
	Default     *bool  `json:"default,omitempty"`     // Boolean indicating whether this is the default testing destination (true) or not (false). Defaults to false. Only one default is allowed
	Description string `json:"description,omitempty"` // Description of the testing destination. Optional, defaults to null
	IP          string `json:"ip,omitempty"`          // The IP address to test connectivity with
}
type RequestApplianceUpdateNetworkApplianceContentFiltering struct {
	AllowedURLPatterns   []string `json:"allowedUrlPatterns,omitempty"`   // A list of URL patterns that are allowed
	BlockedURLCategories []string `json:"blockedUrlCategories,omitempty"` // A list of URL categories to block
	BlockedURLPatterns   []string `json:"blockedUrlPatterns,omitempty"`   // A list of URL patterns that are blocked
	URLCategoryListSize  string   `json:"urlCategoryListSize,omitempty"`  // URL category list size which is either 'topSites' or 'fullList'
}
type RequestApplianceUpdateNetworkApplianceFirewallCellularFirewallRules struct {
	Rules *[]RequestApplianceUpdateNetworkApplianceFirewallCellularFirewallRulesRules `json:"rules,omitempty"` // An ordered array of the firewall rules (not including the default rule)
}
type RequestApplianceUpdateNetworkApplianceFirewallCellularFirewallRulesRules struct {
	Comment       string `json:"comment,omitempty"`       // Description of the rule (optional)
	DestCidr      string `json:"destCidr,omitempty"`      // Comma-separated list of destination IP address(es) (in IP or CIDR notation), fully-qualified domain names (FQDN) or 'any'
	DestPort      string `json:"destPort,omitempty"`      // Comma-separated list of destination port(s) (integer in the range 1-65535), or 'any'
	Policy        string `json:"policy,omitempty"`        // 'allow' or 'deny' traffic specified by this rule
	Protocol      string `json:"protocol,omitempty"`      // The type of protocol (must be 'tcp', 'udp', 'icmp', 'icmp6' or 'any')
	SrcCidr       string `json:"srcCidr,omitempty"`       // Comma-separated list of source IP address(es) (in IP or CIDR notation), or 'any' (note: FQDN not supported for source addresses)
	SrcPort       string `json:"srcPort,omitempty"`       // Comma-separated list of source port(s) (integer in the range 1-65535), or 'any'
	SyslogEnabled *bool  `json:"syslogEnabled,omitempty"` // Log this rule to syslog (true or false, boolean value) - only applicable if a syslog has been configured (optional)
}
type RequestApplianceUpdateNetworkApplianceFirewallFirewalledService struct {
	Access     string   `json:"access,omitempty"`     // A string indicating the rule for which IPs are allowed to use the specified service. Acceptable values are "blocked" (no remote IPs can access the service), "restricted" (only allowed IPs can access the service), and "unrestriced" (any remote IP can access the service). This field is required
	AllowedIPs []string `json:"allowedIps,omitempty"` // An array of allowed IPs that can access the service. This field is required if "access" is set to "restricted". Otherwise this field is ignored
}
type RequestApplianceUpdateNetworkApplianceFirewallInboundCellularFirewallRules struct {
	Rules *[]RequestApplianceUpdateNetworkApplianceFirewallInboundCellularFirewallRulesRules `json:"rules,omitempty"` // An ordered array of the firewall rules (not including the default rule)
}
type RequestApplianceUpdateNetworkApplianceFirewallInboundCellularFirewallRulesRules struct {
	Comment       string `json:"comment,omitempty"`       // Description of the rule (optional)
	DestCidr      string `json:"destCidr,omitempty"`      // Comma-separated list of destination IP address(es) (in IP or CIDR notation), fully-qualified domain names (FQDN) or 'any'
	DestPort      string `json:"destPort,omitempty"`      // Comma-separated list of destination port(s) (integer in the range 1-65535), or 'any'
	Policy        string `json:"policy,omitempty"`        // 'allow' or 'deny' traffic specified by this rule
	Protocol      string `json:"protocol,omitempty"`      // The type of protocol (must be 'tcp', 'udp', 'icmp', 'icmp6' or 'any')
	SrcCidr       string `json:"srcCidr,omitempty"`       // Comma-separated list of source IP address(es) (in IP or CIDR notation), or 'any' (note: FQDN not supported for source addresses)
	SrcPort       string `json:"srcPort,omitempty"`       // Comma-separated list of source port(s) (integer in the range 1-65535), or 'any'
	SyslogEnabled *bool  `json:"syslogEnabled,omitempty"` // Log this rule to syslog (true or false, boolean value) - only applicable if a syslog has been configured (optional)
}
type RequestApplianceUpdateNetworkApplianceFirewallInboundFirewallRules struct {
	Rules             *[]RequestApplianceUpdateNetworkApplianceFirewallInboundFirewallRulesRules `json:"rules,omitempty"`             // An ordered array of the firewall rules (not including the default rule)
	SyslogDefaultRule *bool                                                                      `json:"syslogDefaultRule,omitempty"` // Log the special default rule (boolean value - enable only if you've configured a syslog server) (optional)
}
type RequestApplianceUpdateNetworkApplianceFirewallInboundFirewallRulesRules struct {
	Comment       string `json:"comment,omitempty"`       // Description of the rule (optional)
	DestCidr      string `json:"destCidr,omitempty"`      // Comma-separated list of destination IP address(es) (in IP or CIDR notation), fully-qualified domain names (FQDN) or 'any'
	DestPort      string `json:"destPort,omitempty"`      // Comma-separated list of destination port(s) (integer in the range 1-65535), or 'any'
	Policy        string `json:"policy,omitempty"`        // 'allow' or 'deny' traffic specified by this rule
	Protocol      string `json:"protocol,omitempty"`      // The type of protocol (must be 'tcp', 'udp', 'icmp', 'icmp6' or 'any')
	SrcCidr       string `json:"srcCidr,omitempty"`       // Comma-separated list of source IP address(es) (in IP or CIDR notation), or 'any' (note: FQDN not supported for source addresses)
	SrcPort       string `json:"srcPort,omitempty"`       // Comma-separated list of source port(s) (integer in the range 1-65535), or 'any'
	SyslogEnabled *bool  `json:"syslogEnabled,omitempty"` // Log this rule to syslog (true or false, boolean value) - only applicable if a syslog has been configured (optional)
}
type RequestApplianceUpdateNetworkApplianceFirewallL3FirewallRules struct {
	Rules             *[]RequestApplianceUpdateNetworkApplianceFirewallL3FirewallRulesRules `json:"rules,omitempty"`             // An ordered array of the firewall rules (not including the default rule)
	SyslogDefaultRule *bool                                                                 `json:"syslogDefaultRule,omitempty"` // Log the special default rule (boolean value - enable only if you've configured a syslog server) (optional)
}
type RequestApplianceUpdateNetworkApplianceFirewallL3FirewallRulesRules struct {
	Comment       string `json:"comment,omitempty"`       // Description of the rule (optional)
	DestCidr      string `json:"destCidr,omitempty"`      // Comma-separated list of destination IP address(es) (in IP or CIDR notation), fully-qualified domain names (FQDN) or 'any'
	DestPort      string `json:"destPort,omitempty"`      // Comma-separated list of destination port(s) (integer in the range 1-65535), or 'any'
	Policy        string `json:"policy,omitempty"`        // 'allow' or 'deny' traffic specified by this rule
	Protocol      string `json:"protocol,omitempty"`      // The type of protocol (must be 'tcp', 'udp', 'icmp', 'icmp6' or 'any')
	SrcCidr       string `json:"srcCidr,omitempty"`       // Comma-separated list of source IP address(es) (in IP or CIDR notation), or 'any' (note: FQDN not supported for source addresses)
	SrcPort       string `json:"srcPort,omitempty"`       // Comma-separated list of source port(s) (integer in the range 1-65535), or 'any'
	SyslogEnabled *bool  `json:"syslogEnabled,omitempty"` // Log this rule to syslog (true or false, boolean value) - only applicable if a syslog has been configured (optional)
}
type RequestApplianceUpdateNetworkApplianceFirewallL7FirewallRules struct {
	Rules *[]RequestApplianceUpdateNetworkApplianceFirewallL7FirewallRulesRules `json:"rules,omitempty"` // An ordered array of the MX L7 firewall rules
}
type RequestApplianceUpdateNetworkApplianceFirewallL7FirewallRulesRules struct {
	Policy string      `json:"policy,omitempty"` // 'Deny' traffic specified by this rule
	Type   string      `json:"type,omitempty"`   // Type of the L7 rule. One of: 'application', 'applicationCategory', 'host', 'port', 'ipRange'
	Value  interface{} `json:"value,omitempty"`  // The 'value' of what you want to block. Format of 'value' varies depending on type of the rule. The application categories and application ids can be retrieved from the the 'MX L7 application categories' endpoint. The countries follow the two-letter ISO 3166-1 alpha-2 format.
}
type RequestApplianceUpdateNetworkApplianceFirewallL7FirewallRulesRulesValue struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
type RequestApplianceUpdateNetworkApplianceFirewallMulticastForwarding struct {
	Rules *[]RequestApplianceUpdateNetworkApplianceFirewallMulticastForwardingRules `json:"rules,omitempty"` // Static multicast forwarding rules. Pass an empty array to clear all rules.
}
type RequestApplianceUpdateNetworkApplianceFirewallMulticastForwardingRules struct {
	Address     string   `json:"address,omitempty"`     // IP address
	Description string   `json:"description,omitempty"` // Forwarding rule description.
	VLANIDs     []string `json:"vlanIds,omitempty"`     // List of VLAN IDs
}
type RequestApplianceUpdateNetworkApplianceFirewallOneToManyNatRules struct {
	Rules *[]RequestApplianceUpdateNetworkApplianceFirewallOneToManyNatRulesRules `json:"rules,omitempty"` // An array of 1:Many nat rules
}
type RequestApplianceUpdateNetworkApplianceFirewallOneToManyNatRulesRules struct {
	PortRules *[]RequestApplianceUpdateNetworkApplianceFirewallOneToManyNatRulesRulesPortRules `json:"portRules,omitempty"` // An array of associated forwarding rules
	PublicIP  string                                                                           `json:"publicIp,omitempty"`  // The IP address that will be used to access the internal resource from the WAN
	Uplink    string                                                                           `json:"uplink,omitempty"`    // The physical WAN interface on which the traffic will arrive, formatted as 'internetN' where N is an integer representing a valid uplink for the network's appliance
}
type RequestApplianceUpdateNetworkApplianceFirewallOneToManyNatRulesRulesPortRules struct {
	AllowedIPs []string `json:"allowedIps,omitempty"` // Remote IP addresses or ranges that are permitted to access the internal resource via this port forwarding rule, or 'any'
	LocalIP    string   `json:"localIp,omitempty"`    // Local IP address to which traffic will be forwarded
	LocalPort  string   `json:"localPort,omitempty"`  // Destination port of the forwarded traffic that will be sent from the MX to the specified host on the LAN. If you simply wish to forward the traffic without translating the port, this should be the same as the Public port
	Name       string   `json:"name,omitempty"`       // A description of the rule
	Protocol   string   `json:"protocol,omitempty"`   // 'tcp' or 'udp'
	PublicPort string   `json:"publicPort,omitempty"` // Destination port of the traffic that is arriving on the WAN
}
type RequestApplianceUpdateNetworkApplianceFirewallOneToOneNatRules struct {
	Rules *[]RequestApplianceUpdateNetworkApplianceFirewallOneToOneNatRulesRules `json:"rules,omitempty"` // An array of 1:1 nat rules
}
type RequestApplianceUpdateNetworkApplianceFirewallOneToOneNatRulesRules struct {
	AllowedInbound *[]RequestApplianceUpdateNetworkApplianceFirewallOneToOneNatRulesRulesAllowedInbound `json:"allowedInbound,omitempty"` // The ports this mapping will provide access on, and the remote IPs that will be allowed access to the resource
	LanIP          string                                                                               `json:"lanIp,omitempty"`          // The IP address of the server or device that hosts the internal resource that you wish to make available on the WAN
	Name           string                                                                               `json:"name,omitempty"`           // A descriptive name for the rule
	PublicIP       string                                                                               `json:"publicIp,omitempty"`       // The IP address that will be used to access the internal resource from the WAN
	Uplink         string                                                                               `json:"uplink,omitempty"`         // The physical WAN interface on which the traffic will arrive, formatted as 'internetN' where N is an integer representing a valid uplink for the network's appliance
}
type RequestApplianceUpdateNetworkApplianceFirewallOneToOneNatRulesRulesAllowedInbound struct {
	AllowedIPs       []string `json:"allowedIps,omitempty"`       // An array of ranges of WAN IP addresses that are allowed to make inbound connections on the specified ports or port ranges, or 'any'
	DestinationPorts []string `json:"destinationPorts,omitempty"` // An array of ports or port ranges that will be forwarded to the host on the LAN
	Protocol         string   `json:"protocol,omitempty"`         // Either of the following: 'tcp', 'udp', 'icmp-ping' or 'any'
}
type RequestApplianceUpdateNetworkApplianceFirewallPortForwardingRules struct {
	Rules *[]RequestApplianceUpdateNetworkApplianceFirewallPortForwardingRulesRules `json:"rules,omitempty"` // An array of port forwarding params
}
type RequestApplianceUpdateNetworkApplianceFirewallPortForwardingRulesRules struct {
	AllowedIPs []string `json:"allowedIps,omitempty"` // An array of ranges of WAN IP addresses that are allowed to make inbound connections on the specified ports or port ranges (or any)
	LanIP      string   `json:"lanIp,omitempty"`      // The IP address of the server or device that hosts the internal resource that you wish to make available on the WAN
	LocalPort  string   `json:"localPort,omitempty"`  // A port or port ranges that will receive the forwarded traffic from the WAN
	Name       string   `json:"name,omitempty"`       // A descriptive name for the rule
	Protocol   string   `json:"protocol,omitempty"`   // TCP or UDP
	PublicPort string   `json:"publicPort,omitempty"` // A port or port ranges that will be forwarded to the host on the LAN
	Uplink     string   `json:"uplink,omitempty"`     // The physical WAN interface on which the traffic will arrive ('internet1' or, if available, 'internet2' or 'both')
}
type RequestApplianceUpdateNetworkApplianceFirewallSettings struct {
	SpoofingProtection *RequestApplianceUpdateNetworkApplianceFirewallSettingsSpoofingProtection `json:"spoofingProtection,omitempty"` // Spoofing protection settings
}
type RequestApplianceUpdateNetworkApplianceFirewallSettingsSpoofingProtection struct {
	IPSourceGuard *RequestApplianceUpdateNetworkApplianceFirewallSettingsSpoofingProtectionIPSourceGuard `json:"ipSourceGuard,omitempty"` // IP source address spoofing settings
}
type RequestApplianceUpdateNetworkApplianceFirewallSettingsSpoofingProtectionIPSourceGuard struct {
	Mode string `json:"mode,omitempty"` // Mode of protection
}
type RequestApplianceUpdateNetworkAppliancePort struct {
	AccessPolicy        string `json:"accessPolicy,omitempty"`        // The name of the policy. Only applicable to Access ports. Valid values are: 'open', '8021x-radius', 'mac-radius', 'hybris-radius' for MX64 or Z3 or any MX supporting the per port authentication feature. Otherwise, 'open' is the only valid value and 'open' is the default value if the field is missing.
	AllowedVLANs        string `json:"allowedVlans,omitempty"`        // Comma-delimited list of the VLAN ID's allowed on the port, or 'all' to permit all VLAN's on the port.
	DropUntaggedTraffic *bool  `json:"dropUntaggedTraffic,omitempty"` // Trunk port can Drop all Untagged traffic. When true, no VLAN is required. Access ports cannot have dropUntaggedTraffic set to true.
	Enabled             *bool  `json:"enabled,omitempty"`             // The status of the port
	Type                string `json:"type,omitempty"`                // The type of the port: 'access' or 'trunk'.
	VLAN                *int   `json:"vlan,omitempty"`                // Native VLAN when the port is in Trunk mode. Access VLAN when the port is in Access mode.
}
type RequestApplianceCreateNetworkAppliancePrefixesDelegatedStatic struct {
	Description string                                                               `json:"description,omitempty"` // A name or description for the prefix
	Origin      *RequestApplianceCreateNetworkAppliancePrefixesDelegatedStaticOrigin `json:"origin,omitempty"`      // The origin of the prefix
	Prefix      string                                                               `json:"prefix,omitempty"`      // A static IPv6 prefix
}
type RequestApplianceCreateNetworkAppliancePrefixesDelegatedStaticOrigin struct {
	Interfaces []string `json:"interfaces,omitempty"` // Interfaces associated with the prefix
	Type       string   `json:"type,omitempty"`       // Type of the origin
}
type RequestApplianceUpdateNetworkAppliancePrefixesDelegatedStatic struct {
	Description string                                                               `json:"description,omitempty"` // A name or description for the prefix
	Origin      *RequestApplianceUpdateNetworkAppliancePrefixesDelegatedStaticOrigin `json:"origin,omitempty"`      // The origin of the prefix
	Prefix      string                                                               `json:"prefix,omitempty"`      // A static IPv6 prefix
}
type RequestApplianceUpdateNetworkAppliancePrefixesDelegatedStaticOrigin struct {
	Interfaces []string `json:"interfaces,omitempty"` // Interfaces associated with the prefix
	Type       string   `json:"type,omitempty"`       // Type of the origin
}
type RequestApplianceCreateNetworkApplianceRfProfile struct {
	FiveGhzSettings    *RequestApplianceCreateNetworkApplianceRfProfileFiveGhzSettings    `json:"fiveGhzSettings,omitempty"`    // Settings related to 5Ghz band
	Name               string                                                             `json:"name,omitempty"`               // The name of the new profile. Must be unique. This param is required on creation.
	PerSSIDSettings    *RequestApplianceCreateNetworkApplianceRfProfilePerSSIDSettings    `json:"perSsidSettings,omitempty"`    // Per-SSID radio settings by number.
	TwoFourGhzSettings *RequestApplianceCreateNetworkApplianceRfProfileTwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"` // Settings related to 2.4Ghz band
}
type RequestApplianceCreateNetworkApplianceRfProfileFiveGhzSettings struct {
	AxEnabled  *bool `json:"axEnabled,omitempty"`  // Determines whether ax radio on 5Ghz band is on or off. Can be either true or false. If false, we highly recommend disabling band steering. Defaults to true.
	MinBitrate *int  `json:"minBitrate,omitempty"` // Sets min bitrate (Mbps) of 5Ghz band. Can be one of '6', '9', '12', '18', '24', '36', '48' or '54'. Defaults to 12.
}
type RequestApplianceCreateNetworkApplianceRfProfilePerSSIDSettings struct {
	Status1 *RequestApplianceCreateNetworkApplianceRfProfilePerSSIDSettings1 `json:"1,omitempty"` // Settings for SSID 1
	Status2 *RequestApplianceCreateNetworkApplianceRfProfilePerSSIDSettings2 `json:"2,omitempty"` // Settings for SSID 2
	Status3 *RequestApplianceCreateNetworkApplianceRfProfilePerSSIDSettings3 `json:"3,omitempty"` // Settings for SSID 3
	Status4 *RequestApplianceCreateNetworkApplianceRfProfilePerSSIDSettings4 `json:"4,omitempty"` // Settings for SSID 4
}
type RequestApplianceCreateNetworkApplianceRfProfilePerSSIDSettings1 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
}
type RequestApplianceCreateNetworkApplianceRfProfilePerSSIDSettings2 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
}
type RequestApplianceCreateNetworkApplianceRfProfilePerSSIDSettings3 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
}
type RequestApplianceCreateNetworkApplianceRfProfilePerSSIDSettings4 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
}
type RequestApplianceCreateNetworkApplianceRfProfileTwoFourGhzSettings struct {
	AxEnabled  *bool    `json:"axEnabled,omitempty"`  // Determines whether ax radio on 2.4Ghz band is on or off. Can be either true or false. If false, we highly recommend disabling band steering. Defaults to true.
	MinBitrate *float64 `json:"minBitrate,omitempty"` // Sets min bitrate (Mbps) of 2.4Ghz band. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'. Defaults to 11.
}
type RequestApplianceUpdateNetworkApplianceRfProfile struct {
	FiveGhzSettings    *RequestApplianceUpdateNetworkApplianceRfProfileFiveGhzSettings    `json:"fiveGhzSettings,omitempty"`    // Settings related to 5Ghz band
	Name               string                                                             `json:"name,omitempty"`               // The name of the new profile. Must be unique.
	PerSSIDSettings    *RequestApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings    `json:"perSsidSettings,omitempty"`    // Per-SSID radio settings by number.
	TwoFourGhzSettings *RequestApplianceUpdateNetworkApplianceRfProfileTwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"` // Settings related to 2.4Ghz band
}
type RequestApplianceUpdateNetworkApplianceRfProfileFiveGhzSettings struct {
	AxEnabled  *bool `json:"axEnabled,omitempty"`  // Determines whether ax radio on 5Ghz band is on or off. Can be either true or false. If false, we highly recommend disabling band steering.
	MinBitrate *int  `json:"minBitrate,omitempty"` // Sets min bitrate (Mbps) of 5Ghz band. Can be one of '6', '9', '12', '18', '24', '36', '48' or '54'.
}
type RequestApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings struct {
	Status1 *RequestApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings1 `json:"1,omitempty"` // Settings for SSID 1
	Status2 *RequestApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings2 `json:"2,omitempty"` // Settings for SSID 2
	Status3 *RequestApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings3 `json:"3,omitempty"` // Settings for SSID 3
	Status4 *RequestApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings4 `json:"4,omitempty"` // Settings for SSID 4
}
type RequestApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings1 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
}
type RequestApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings2 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
}
type RequestApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings3 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
}
type RequestApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings4 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
}
type RequestApplianceUpdateNetworkApplianceRfProfileTwoFourGhzSettings struct {
	AxEnabled  *bool    `json:"axEnabled,omitempty"`  // Determines whether ax radio on 2.4Ghz band is on or off. Can be either true or false. If false, we highly recommend disabling band steering.
	MinBitrate *float64 `json:"minBitrate,omitempty"` // Sets min bitrate (Mbps) of 2.4Ghz band. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestApplianceUpdateNetworkApplianceSdwanInternetPolicies struct {
	WanTrafficUplinkPreferences *[]RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences `json:"wanTrafficUplinkPreferences,omitempty"` // policies with respective traffic filters for an MX network
}
type RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences struct {
	FailOverCriterion string                                                                                                  `json:"failOverCriterion,omitempty"` // WAN failover and failback behavior
	PerformanceClass  *RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesPerformanceClass `json:"performanceClass,omitempty"`  // Performance class setting for uplink preference rule
	PreferredUplink   string                                                                                                  `json:"preferredUplink,omitempty"`   // Preferred uplink for uplink preference rule. Must be one of: 'wan1', 'wan2', 'bestForVoIP', 'loadBalancing' or 'defaultUplink'
	TrafficFilters    *[]RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFilters `json:"trafficFilters,omitempty"`    // Traffic filters
}
type RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesPerformanceClass struct {
	BuiltinPerformanceClassName string `json:"builtinPerformanceClassName,omitempty"` // Name of builtin performance class. Must be present when performanceClass type is 'builtin' and value must be one of: 'VoIP'
	CustomPerformanceClassID    string `json:"customPerformanceClassId,omitempty"`    // ID of created custom performance class, must be present when performanceClass type is "custom"
	Type                        string `json:"type,omitempty"`                        // Type of this performance class. Must be one of: 'builtin' or 'custom'
}
type RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFilters struct {
	Type  string                                                                                                     `json:"type,omitempty"`  // Traffic filter type. Must be 'custom', 'major_application', 'application (NBAR)', if type is 'application', you can pass either an NBAR App Category or Application
	Value *RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValue `json:"value,omitempty"` // Value of traffic filter
}
type RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValue struct {
	Destination *RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueDestination `json:"destination,omitempty"` // Destination of 'custom' type traffic filter
	Protocol    string                                                                                                                `json:"protocol,omitempty"`    // Protocol of the traffic filter. Must be one of: 'tcp', 'udp', 'icmp6' or 'any'
	Source      *RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueSource      `json:"source,omitempty"`      // Source of traffic filter
}
type RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueDestination struct {
	Applications *[]RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueDestinationApplications `json:"applications,omitempty"` // list of application objects (either majorApplication or nbar)
	Cidr         string                                                                                                                              `json:"cidr,omitempty"`         // CIDR format address (e.g."192.168.10.1", which is the same as "192.168.10.1/32"), or "any"
	Port         string                                                                                                                              `json:"port,omitempty"`         // E.g.: "any", "0" (also means "any"), "8080", "1-1024"
}
type RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueDestinationApplications struct {
	ID   string `json:"id,omitempty"`   // Id of the major application, or a list of NBAR Application Category or Application selections
	Name string `json:"name,omitempty"` // Name of the major application or application category selected
	Type string `json:"type,omitempty"` // app type (major or nbar)
}
type RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueSource struct {
	Cidr string `json:"cidr,omitempty"` // CIDR format address (e.g."192.168.10.1", which is the same as "192.168.10.1/32"), or "any". Cannot be used in combination with the "vlan" property
	Host *int   `json:"host,omitempty"` // Host ID in the VLAN. Should not exceed the VLAN subnet capacity. Must be used along with the "vlan" property and is currently only available under a template network.
	Port string `json:"port,omitempty"` // E.g.: "any", "0" (also means "any"), "8080", "1-1024"
	VLAN *int   `json:"vlan,omitempty"` // VLAN ID of the configured VLAN in the Meraki network. Cannot be used in combination with the "cidr" property and is currently only available under a template network.
}
type RequestApplianceUpdateNetworkApplianceSecurityIntrusion struct {
	IDsRulesets       string                                                                    `json:"idsRulesets,omitempty"`       // Set the detection ruleset 'connectivity'/'balanced'/'security' (optional - omitting will leave current config unchanged). Default value is 'balanced' if none currently saved
	Mode              string                                                                    `json:"mode,omitempty"`              // Set mode to 'disabled'/'detection'/'prevention' (optional - omitting will leave current config unchanged)
	ProtectedNetworks *RequestApplianceUpdateNetworkApplianceSecurityIntrusionProtectedNetworks `json:"protectedNetworks,omitempty"` // Set the included/excluded networks from the intrusion engine (optional - omitting will leave current config unchanged). This is available only in 'passthrough' mode
}
type RequestApplianceUpdateNetworkApplianceSecurityIntrusionProtectedNetworks struct {
	ExcludedCidr []string `json:"excludedCidr,omitempty"` // list of IP addresses or subnets being excluded from protection (required if 'useDefault' is false)
	IncludedCidr []string `json:"includedCidr,omitempty"` // list of IP addresses or subnets being protected (required if 'useDefault' is false)
	UseDefault   *bool    `json:"useDefault,omitempty"`   // true/false whether to use special IPv4 addresses: https://tools.ietf.org/html/rfc5735 (required). Default value is true if none currently saved
}
type RequestApplianceUpdateNetworkApplianceSecurityMalware struct {
	AllowedFiles *[]RequestApplianceUpdateNetworkApplianceSecurityMalwareAllowedFiles `json:"allowedFiles,omitempty"` // The sha256 digests of files that should be permitted by the malware detection engine. If omitted, the current config will remain unchanged. This is available only if your network supports AMP allow listing
	AllowedURLs  *[]RequestApplianceUpdateNetworkApplianceSecurityMalwareAllowedURLs  `json:"allowedUrls,omitempty"`  // The urls that should be permitted by the malware detection engine. If omitted, the current config will remain unchanged. This is available only if your network supports AMP allow listing
	Mode         string                                                               `json:"mode,omitempty"`         // Set mode to 'enabled' to enable malware prevention, otherwise 'disabled'
}
type RequestApplianceUpdateNetworkApplianceSecurityMalwareAllowedFiles struct {
	Comment string `json:"comment,omitempty"` // Comment about the allowed entity
	Sha256  string `json:"sha256,omitempty"`  // The file sha256 hash to allow
}
type RequestApplianceUpdateNetworkApplianceSecurityMalwareAllowedURLs struct {
	Comment string `json:"comment,omitempty"` // Comment about the allowed entity
	URL     string `json:"url,omitempty"`     // The url to allow
}
type RequestApplianceUpdateNetworkApplianceSettings struct {
	ClientTrackingMethod string                                                    `json:"clientTrackingMethod,omitempty"` // Client tracking method of a network
	DeploymentMode       string                                                    `json:"deploymentMode,omitempty"`       // Deployment mode of a network
	DynamicDNS           *RequestApplianceUpdateNetworkApplianceSettingsDynamicDNS `json:"dynamicDns,omitempty"`           // Dynamic DNS settings for a network
}
type RequestApplianceUpdateNetworkApplianceSettingsDynamicDNS struct {
	Enabled *bool  `json:"enabled,omitempty"` // Dynamic DNS enabled
	Prefix  string `json:"prefix,omitempty"`  // Dynamic DNS url prefix. DDNS must be enabled to update
}
type RequestApplianceUpdateNetworkApplianceSingleLan struct {
	ApplianceIP   string                                                        `json:"applianceIp,omitempty"`   // The appliance IP address of the single LAN
	IPv6          *RequestApplianceUpdateNetworkApplianceSingleLanIPv6          `json:"ipv6,omitempty"`          // IPv6 configuration on the VLAN
	MandatoryDhcp *RequestApplianceUpdateNetworkApplianceSingleLanMandatoryDhcp `json:"mandatoryDhcp,omitempty"` // Mandatory DHCP will enforce that clients connecting to this LAN must use the IP address assigned by the DHCP server. Clients who use a static IP address won't be able to associate. Only available on firmware versions 17.0 and above
	Subnet        string                                                        `json:"subnet,omitempty"`        // The subnet of the single LAN configuration
}
type RequestApplianceUpdateNetworkApplianceSingleLanIPv6 struct {
	Enabled           *bool                                                                   `json:"enabled,omitempty"`           // Enable IPv6 on VLAN.
	PrefixAssignments *[]RequestApplianceUpdateNetworkApplianceSingleLanIPv6PrefixAssignments `json:"prefixAssignments,omitempty"` // Prefix assignments on the VLAN
}
type RequestApplianceUpdateNetworkApplianceSingleLanIPv6PrefixAssignments struct {
	Autonomous         *bool                                                                       `json:"autonomous,omitempty"`         // Auto assign a /64 prefix from the origin to the VLAN
	Origin             *RequestApplianceUpdateNetworkApplianceSingleLanIPv6PrefixAssignmentsOrigin `json:"origin,omitempty"`             // The origin of the prefix
	StaticApplianceIP6 string                                                                      `json:"staticApplianceIp6,omitempty"` // Manual configuration of the IPv6 Appliance IP
	StaticPrefix       string                                                                      `json:"staticPrefix,omitempty"`       // Manual configuration of a /64 prefix on the VLAN
}
type RequestApplianceUpdateNetworkApplianceSingleLanIPv6PrefixAssignmentsOrigin struct {
	Interfaces []string `json:"interfaces,omitempty"` // Interfaces associated with the prefix
	Type       string   `json:"type,omitempty"`       // Type of the origin
}
type RequestApplianceUpdateNetworkApplianceSingleLanMandatoryDhcp struct {
	Enabled *bool `json:"enabled,omitempty"` // Enable Mandatory DHCP on LAN.
}
type RequestApplianceUpdateNetworkApplianceSSID struct {
	AuthMode                     string                                                                  `json:"authMode,omitempty"`                     // The association control method for the SSID ('open', 'psk', '8021x-meraki' or '8021x-radius').
	DefaultVLANID                *int                                                                    `json:"defaultVlanId,omitempty"`                // The VLAN ID of the VLAN associated to this SSID. This parameter is only valid if the network is in routed mode.
	DhcpEnforcedDeauthentication *RequestApplianceUpdateNetworkApplianceSSIDDhcpEnforcedDeauthentication `json:"dhcpEnforcedDeauthentication,omitempty"` // DHCP Enforced Deauthentication enables the disassociation of wireless clients in addition to Mandatory DHCP. This param is only valid on firmware versions >= MX 17.0 where the associated LAN has Mandatory DHCP Enabled
	Dot11W                       *RequestApplianceUpdateNetworkApplianceSSIDDot11W                       `json:"dot11w,omitempty"`                       // The current setting for Protected Management Frames (802.11w).
	Enabled                      *bool                                                                   `json:"enabled,omitempty"`                      // Whether or not the SSID is enabled.
	EncryptionMode               string                                                                  `json:"encryptionMode,omitempty"`               // The psk encryption mode for the SSID ('wep' or 'wpa'). This param is only valid if the authMode is 'psk'.
	Name                         string                                                                  `json:"name,omitempty"`                         // The name of the SSID.
	Psk                          string                                                                  `json:"psk,omitempty"`                          // The passkey for the SSID. This param is only valid if the authMode is 'psk'.
	RadiusServers                *[]RequestApplianceUpdateNetworkApplianceSSIDRadiusServers              `json:"radiusServers,omitempty"`                // The RADIUS 802.1x servers to be used for authentication. This param is only valid if the authMode is '8021x-radius'.
	Visible                      *bool                                                                   `json:"visible,omitempty"`                      // Boolean indicating whether the MX should advertise or hide this SSID.
	WpaEncryptionMode            string                                                                  `json:"wpaEncryptionMode,omitempty"`            // The types of WPA encryption. ('WPA1 and WPA2', 'WPA2 only', 'WPA3 Transition Mode' or 'WPA3 only'). This param is only valid if (1) the authMode is 'psk' & the encryptionMode is 'wpa' OR (2) the authMode is '8021x-meraki' OR (3) the authMode is '8021x-radius'
}
type RequestApplianceUpdateNetworkApplianceSSIDDhcpEnforcedDeauthentication struct {
	Enabled *bool `json:"enabled,omitempty"` // Enable DCHP Enforced Deauthentication on the SSID.
}
type RequestApplianceUpdateNetworkApplianceSSIDDot11W struct {
	Enabled  *bool `json:"enabled,omitempty"`  // Whether 802.11w is enabled or not.
	Required *bool `json:"required,omitempty"` // (Optional) Whether 802.11w is required or not.
}
type RequestApplianceUpdateNetworkApplianceSSIDRadiusServers struct {
	Host   string `json:"host,omitempty"`   // The IP address of your RADIUS server.
	Port   *int   `json:"port,omitempty"`   // The UDP port your RADIUS servers listens on for Access-requests.
	Secret string `json:"secret,omitempty"` // The RADIUS client shared secret.
}
type RequestApplianceCreateNetworkApplianceStaticRoute struct {
	GatewayIP     string `json:"gatewayIp,omitempty"`     // Gateway IP address (next hop)
	GatewayVLANID string `json:"gatewayVlanId,omitempty"` // Gateway VLAN ID
	Name          string `json:"name,omitempty"`          // Name of the route
	Subnet        string `json:"subnet,omitempty"`        // Subnet of the route
}
type RequestApplianceUpdateNetworkApplianceStaticRoute struct {
	Enabled            *bool                                                                `json:"enabled,omitempty"`            // Whether the route should be enabled or not
	FixedIPAssignments *RequestApplianceUpdateNetworkApplianceStaticRouteFixedIPAssignments `json:"fixedIpAssignments,omitempty"` // Fixed DHCP IP assignments on the route
	GatewayIP          string                                                               `json:"gatewayIp,omitempty"`          // Gateway IP address (next hop)
	GatewayVLANID      string                                                               `json:"gatewayVlanId,omitempty"`      // Gateway VLAN ID
	Name               string                                                               `json:"name,omitempty"`               // Name of the route
	ReservedIPRanges   *[]RequestApplianceUpdateNetworkApplianceStaticRouteReservedIPRanges `json:"reservedIpRanges,omitempty"`   // DHCP reserved IP ranges
	Subnet             string                                                               `json:"subnet,omitempty"`             // Subnet of the route
}
type RequestApplianceUpdateNetworkApplianceStaticRouteFixedIPAssignments interface{}
type RequestApplianceUpdateNetworkApplianceStaticRouteReservedIPRanges struct {
	Comment string `json:"comment,omitempty"` // Description of the range
	End     string `json:"end,omitempty"`     // Last address in the reserved range
	Start   string `json:"start,omitempty"`   // First address in the reserved range
}
type RequestApplianceUpdateNetworkApplianceTrafficShaping struct {
	GlobalBandwidthLimits *RequestApplianceUpdateNetworkApplianceTrafficShapingGlobalBandwidthLimits `json:"globalBandwidthLimits,omitempty"` // Global per-client bandwidth limit
}
type RequestApplianceUpdateNetworkApplianceTrafficShapingGlobalBandwidthLimits struct {
	LimitDown *int `json:"limitDown,omitempty"` // The download bandwidth limit in Kbps. (0 represents no limit.)
	LimitUp   *int `json:"limitUp,omitempty"`   // The upload bandwidth limit in Kbps. (0 represents no limit.)
}
type RequestApplianceCreateNetworkApplianceTrafficShapingCustomPerformanceClass struct {
	MaxJitter         *int   `json:"maxJitter,omitempty"`         // Maximum jitter in milliseconds
	MaxLatency        *int   `json:"maxLatency,omitempty"`        // Maximum latency in milliseconds
	MaxLossPercentage *int   `json:"maxLossPercentage,omitempty"` // Maximum percentage of packet loss
	Name              string `json:"name,omitempty"`              // Name of the custom performance class
}
type RequestApplianceUpdateNetworkApplianceTrafficShapingCustomPerformanceClass struct {
	MaxJitter         *int   `json:"maxJitter,omitempty"`         // Maximum jitter in milliseconds
	MaxLatency        *int   `json:"maxLatency,omitempty"`        // Maximum latency in milliseconds
	MaxLossPercentage *int   `json:"maxLossPercentage,omitempty"` // Maximum percentage of packet loss
	Name              string `json:"name,omitempty"`              // Name of the custom performance class
}
type RequestApplianceUpdateNetworkApplianceTrafficShapingRules struct {
	DefaultRulesEnabled *bool                                                             `json:"defaultRulesEnabled,omitempty"` // Whether default traffic shaping rules are enabled (true) or disabled (false). There are 4 default rules, which can be seen on your network's traffic shaping page. Note that default rules count against the rule limit of 8.
	Rules               *[]RequestApplianceUpdateNetworkApplianceTrafficShapingRulesRules `json:"rules,omitempty"`               //     An array of traffic shaping rules. Rules are applied in the order that     they are specified in. An empty list (or null) means no rules. Note that     you are allowed a maximum of 8 rules.
}
type RequestApplianceUpdateNetworkApplianceTrafficShapingRulesRules struct {
	Definitions              *[]RequestApplianceUpdateNetworkApplianceTrafficShapingRulesRulesDefinitions            `json:"definitions,omitempty"`              //     A list of objects describing the definitions of your traffic shaping rule. At least one definition is required.
	DscpTagValue             *int                                                                                    `json:"dscpTagValue,omitempty"`             //     The DSCP tag applied by your rule. null means 'Do not change DSCP tag'.     For a list of possible tag values, use the trafficShaping/dscpTaggingOptions endpoint.
	PerClientBandwidthLimits *RequestApplianceUpdateNetworkApplianceTrafficShapingRulesRulesPerClientBandwidthLimits `json:"perClientBandwidthLimits,omitempty"` //     An object describing the bandwidth settings for your rule.
	Priority                 string                                                                                  `json:"priority,omitempty"`                 //     A string, indicating the priority level for packets bound to your rule.     Can be 'low', 'normal' or 'high'.
}
type RequestApplianceUpdateNetworkApplianceTrafficShapingRulesRulesDefinitions struct {
	Type  string      `json:"type,omitempty"`  // The type of definition. Can be one of 'application', 'applicationCategory', 'host', 'port', 'ipRange' or 'localNet'.
	Value interface{} `json:"value,omitempty"` //     If "type" is 'host', 'port', 'ipRange' or 'localNet', then "value" must be a string, matching either     a hostname (e.g. "somesite.com"), a port (e.g. 8080), or an IP range ("192.1.0.0",     "192.1.0.0/16", or "10.1.0.0/16:80"). 'localNet' also supports CIDR notation, excluding     custom ports.      If "type" is 'application' or 'applicationCategory', then "value" must be an object     with the structure { "id": "meraki:layer7/..." }, where "id" is the application category or     application ID (for a list of IDs for your network, use the trafficShaping/applicationCategories     endpoint).
}
type RequestApplianceUpdateNetworkApplianceTrafficShapingRulesRulesPerClientBandwidthLimits struct {
	BandwidthLimits *RequestApplianceUpdateNetworkApplianceTrafficShapingRulesRulesPerClientBandwidthLimitsBandwidthLimits `json:"bandwidthLimits,omitempty"` // The bandwidth limits object, specifying the upload ('limitUp') and download ('limitDown') speed in Kbps. These are only enforced if 'settings' is set to 'custom'.
	Settings        string                                                                                                 `json:"settings,omitempty"`        // How bandwidth limits are applied by your rule. Can be one of 'network default', 'ignore' or 'custom'.
}
type RequestApplianceUpdateNetworkApplianceTrafficShapingRulesRulesPerClientBandwidthLimitsBandwidthLimits struct {
	LimitDown *int `json:"limitDown,omitempty"` // The maximum download limit (integer, in Kbps).
	LimitUp   *int `json:"limitUp,omitempty"`   // The maximum upload limit (integer, in Kbps).
}
type RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkBandwidth struct {
	BandwidthLimits *RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimits `json:"bandwidthLimits,omitempty"` // A mapping of uplinks to their bandwidth settings (be sure to check which uplinks are supported for your network)
}
type RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimits struct {
	Cellular *RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular `json:"cellular,omitempty"` // The bandwidth settings for the 'cellular' uplink
	Wan1     *RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan1     `json:"wan1,omitempty"`     // The bandwidth settings for the 'wan1' uplink
	Wan2     *RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2     `json:"wan2,omitempty"`     // The bandwidth settings for the 'wan2' uplink
}
type RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular struct {
	LimitDown *int `json:"limitDown,omitempty"` // The maximum download limit (integer, in Kbps). null indicates no limit
	LimitUp   *int `json:"limitUp,omitempty"`   // The maximum upload limit (integer, in Kbps). null indicates no limit
}
type RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan1 struct {
	LimitDown *int `json:"limitDown,omitempty"` // The maximum download limit (integer, in Kbps). null indicates no limit
	LimitUp   *int `json:"limitUp,omitempty"`   // The maximum upload limit (integer, in Kbps). null indicates no limit
}
type RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2 struct {
	LimitDown *int `json:"limitDown,omitempty"` // The maximum download limit (integer, in Kbps). null indicates no limit
	LimitUp   *int `json:"limitUp,omitempty"`   // The maximum upload limit (integer, in Kbps). null indicates no limit
}
type RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelection struct {
	ActiveActiveAutoVpnEnabled  *bool                                                                                             `json:"activeActiveAutoVpnEnabled,omitempty"`  // Toggle for enabling or disabling active-active AutoVPN
	DefaultUplink               string                                                                                            `json:"defaultUplink,omitempty"`               // The default uplink. Must be one of: 'wan1' or 'wan2'
	FailoverAndFailback         *RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionFailoverAndFailback           `json:"failoverAndFailback,omitempty"`         // WAN failover and failback behavior
	LoadBalancingEnabled        *bool                                                                                             `json:"loadBalancingEnabled,omitempty"`        // Toggle for enabling or disabling load balancing
	VpnTrafficUplinkPreferences *[]RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences `json:"vpnTrafficUplinkPreferences,omitempty"` // Array of uplink preference rules for VPN traffic
	WanTrafficUplinkPreferences *[]RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences `json:"wanTrafficUplinkPreferences,omitempty"` // Array of uplink preference rules for WAN traffic
}
type RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionFailoverAndFailback struct {
	Immediate *RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate `json:"immediate,omitempty"` // Immediate WAN transition terminates all flows (new and existing) on current WAN when it is deemed unreliable.
}
type RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate struct {
	Enabled *bool `json:"enabled,omitempty"` // Toggle for enabling or disabling immediate WAN failover and failback
}
type RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences struct {
	FailOverCriterion string                                                                                                          `json:"failOverCriterion,omitempty"` // Fail over criterion for this uplink preference rule. Must be one of: 'poorPerformance' or 'uplinkDown'
	PerformanceClass  *RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesPerformanceClass `json:"performanceClass,omitempty"`  // Performance class setting for this uplink preference rule
	PreferredUplink   string                                                                                                          `json:"preferredUplink,omitempty"`   // Preferred uplink for this uplink preference rule. Must be one of: 'wan1', 'wan2', 'bestForVoIP', 'loadBalancing' or 'defaultUplink'
	TrafficFilters    *[]RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFilters `json:"trafficFilters,omitempty"`    // Array of traffic filters for this uplink preference rule
}
type RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesPerformanceClass struct {
	BuiltinPerformanceClassName string `json:"builtinPerformanceClassName,omitempty"` // Name of builtin performance class, must be present when performanceClass type is 'builtin', and value must be one of: 'VoIP'
	CustomPerformanceClassID    string `json:"customPerformanceClassId,omitempty"`    // ID of created custom performance class, must be present when performanceClass type is 'custom'
	Type                        string `json:"type,omitempty"`                        // Type of this performance class. Must be one of: 'builtin' or 'custom'
}
type RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFilters struct {
	Type  string                                                                                                             `json:"type,omitempty"`  // Type of this traffic filter. Must be one of: 'applicationCategory', 'application' or 'custom'
	Value *RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValue `json:"value,omitempty"` // Value object of this traffic filter
}
type RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValue struct {
	Destination *RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValueDestination `json:"destination,omitempty"` // Destination of this custom type traffic filter
	ID          string                                                                                                                        `json:"id,omitempty"`          // ID of this applicationCategory or application type traffic filter. E.g.: "meraki:layer7/category/1", "meraki:layer7/application/4"
	Protocol    string                                                                                                                        `json:"protocol,omitempty"`    // Protocol of this custom type traffic filter. Must be one of: 'tcp', 'udp', 'icmp', 'icmp6' or 'any'
	Source      *RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValueSource      `json:"source,omitempty"`      // Source of this custom type traffic filter
}
type RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValueDestination struct {
	Cidr    string `json:"cidr,omitempty"`    // CIDR format address, or "any". E.g.: "192.168.10.0/24",  "192.168.10.1" (same as "192.168.10.1/32"), "0.0.0.0/0" (same as "any")
	Fqdn    string `json:"fqdn,omitempty"`    // FQDN format address. Currently only availabe in 'destination' of 'vpnTrafficUplinkPreference' object. E.g.: 'www.google.com'
	Host    *int   `json:"host,omitempty"`    // Host ID in the VLAN, should be used along with 'vlan', and not exceed the vlan subnet capacity. Currently only available under a template network.
	Network string `json:"network,omitempty"` // Meraki network ID. Currently only available under a template network, and the value should be ID of either same template network, or another template network currently. E.g.: "L_12345678".
	Port    string `json:"port,omitempty"`    // E.g.: "any", "0" (also means "any"), "8080", "1-1024"
	VLAN    *int   `json:"vlan,omitempty"`    // VLAN ID of the configured VLAN in the Meraki network. Currently only available under a template network.
}
type RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValueSource struct {
	Cidr    string `json:"cidr,omitempty"`    // CIDR format address, or "any". E.g.: "192.168.10.0/24",  "192.168.10.1" (same as "192.168.10.1/32"), "0.0.0.0/0" (same as "any")
	Host    *int   `json:"host,omitempty"`    // Host ID in the VLAN, should be used along with 'vlan', and not exceed the vlan subnet capacity. Currently only available under a template network.
	Network string `json:"network,omitempty"` // Meraki network ID. Currently only available under a template network, and the value should be ID of either same template network, or another template network currently. E.g.: "L_12345678".
	Port    string `json:"port,omitempty"`    // E.g.: "any", "0" (also means "any"), "8080", "1-1024"
	VLAN    *int   `json:"vlan,omitempty"`    // VLAN ID of the configured VLAN in the Meraki network. Currently only available under a template network.
}
type RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences struct {
	PreferredUplink string                                                                                                          `json:"preferredUplink,omitempty"` // Preferred uplink for this uplink preference rule. Must be one of: 'wan1' or 'wan2'
	TrafficFilters  *[]RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFilters `json:"trafficFilters,omitempty"`  // Array of traffic filters for this uplink preference rule
}
type RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFilters struct {
	Type  string                                                                                                             `json:"type,omitempty"`  // Type of this traffic filter. Must be one of: 'custom'
	Value *RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValue `json:"value,omitempty"` // Value object of this traffic filter
}
type RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValue struct {
	Destination *RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueDestination `json:"destination,omitempty"` // Destination of this custom type traffic filter
	Protocol    string                                                                                                                        `json:"protocol,omitempty"`    // Protocol of this custom type traffic filter. Must be one of: 'tcp', 'udp', 'icmp6' or 'any'
	Source      *RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueSource      `json:"source,omitempty"`      // Source of this custom type traffic filter
}
type RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueDestination struct {
	Cidr string `json:"cidr,omitempty"` // CIDR format address, or "any". E.g.: "192.168.10.0/24",  "192.168.10.1" (same as "192.168.10.1/32"), "0.0.0.0/0" (same as "any")
	Port string `json:"port,omitempty"` // E.g.: "any", "0" (also means "any"), "8080", "1-1024"
}
type RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueSource struct {
	Cidr string `json:"cidr,omitempty"` // CIDR format address, or "any". E.g.: "192.168.10.0/24",  "192.168.10.1" (same as "192.168.10.1/32"), "0.0.0.0/0" (same as "any")
	Host *int   `json:"host,omitempty"` // Host ID in the VLAN, should be used along with 'vlan', and not exceed the vlan subnet capacity. Currently only available under a template network.
	Port string `json:"port,omitempty"` // E.g.: "any", "0" (also means "any"), "8080", "1-1024"
	VLAN *int   `json:"vlan,omitempty"` // VLAN ID of the configured VLAN in the Meraki network. Currently only available under a template network.
}
type RequestApplianceUpdateNetworkApplianceTrafficShapingVpnExclusions struct {
	Custom            *[]RequestApplianceUpdateNetworkApplianceTrafficShapingVpnExclusionsCustom            `json:"custom,omitempty"`            // Custom VPN exclusion rules. Pass an empty array to clear existing rules.
	MajorApplications *[]RequestApplianceUpdateNetworkApplianceTrafficShapingVpnExclusionsMajorApplications `json:"majorApplications,omitempty"` // Major Application based VPN exclusion rules. Pass an empty array to clear existing rules.
}
type RequestApplianceUpdateNetworkApplianceTrafficShapingVpnExclusionsCustom struct {
	Destination string `json:"destination,omitempty"` // Destination address; hostname required for DNS, IPv4 otherwise.
	Port        string `json:"port,omitempty"`        // Destination port.
	Protocol    string `json:"protocol,omitempty"`    // Protocol.
}
type RequestApplianceUpdateNetworkApplianceTrafficShapingVpnExclusionsMajorApplications struct {
	ID   string `json:"id,omitempty"`   // Application's Meraki ID.
	Name string `json:"name,omitempty"` // Application's name.
}
type RequestApplianceCreateNetworkApplianceVLAN struct {
	ApplianceIP            string                                                   `json:"applianceIp,omitempty"`            // The local IP of the appliance on the VLAN
	Cidr                   string                                                   `json:"cidr,omitempty"`                   // CIDR of the pool of subnets. Applicable only for template network. Each network bound to the template will automatically pick a subnet from this pool to build its own VLAN.
	DhcpBootOptionsEnabled *bool                                                    `json:"dhcpBootOptionsEnabled,omitempty"` // Use DHCP boot options specified in other properties
	DhcpHandling           string                                                   `json:"dhcpHandling,omitempty"`           // The appliance's handling of DHCP requests on this VLAN. One of: 'Run a DHCP server', 'Relay DHCP to another server' or 'Do not respond to DHCP requests'
	DhcpLeaseTime          string                                                   `json:"dhcpLeaseTime,omitempty"`          // The term of DHCP leases if the appliance is running a DHCP server on this VLAN. One of: '30 minutes', '1 hour', '4 hours', '12 hours', '1 day' or '1 week'
	DhcpOptions            *[]RequestApplianceCreateNetworkApplianceVLANDhcpOptions `json:"dhcpOptions,omitempty"`            // The list of DHCP options that will be included in DHCP responses. Each object in the list should have "code", "type", and "value" properties.
	GroupPolicyID          string                                                   `json:"groupPolicyId,omitempty"`          // The id of the desired group policy to apply to the VLAN
	ID                     string                                                   `json:"id,omitempty"`                     // The VLAN ID of the new VLAN (must be between 1 and 4094)
	IPv6                   *RequestApplianceCreateNetworkApplianceVLANIPv6          `json:"ipv6,omitempty"`                   // IPv6 configuration on the VLAN
	MandatoryDhcp          *RequestApplianceCreateNetworkApplianceVLANMandatoryDhcp `json:"mandatoryDhcp,omitempty"`          // Mandatory DHCP will enforce that clients connecting to this VLAN must use the IP address assigned by the DHCP server. Clients who use a static IP address won't be able to associate. Only available on firmware versions 17.0 and above
	Mask                   *int                                                     `json:"mask,omitempty"`                   // Mask used for the subnet of all bound to the template networks. Applicable only for template network.
	Name                   string                                                   `json:"name,omitempty"`                   // The name of the new VLAN
	Subnet                 string                                                   `json:"subnet,omitempty"`                 // The subnet of the VLAN
	TemplateVLANType       string                                                   `json:"templateVlanType,omitempty"`       // Type of subnetting of the VLAN. Applicable only for template network.
}
type RequestApplianceCreateNetworkApplianceVLANDhcpOptions struct {
	Code  string `json:"code,omitempty"`  // The code for the DHCP option. This should be an integer between 2 and 254.
	Type  string `json:"type,omitempty"`  // The type for the DHCP option. One of: 'text', 'ip', 'hex' or 'integer'
	Value string `json:"value,omitempty"` // The value for the DHCP option
}
type RequestApplianceCreateNetworkApplianceVLANIPv6 struct {
	Enabled           *bool                                                              `json:"enabled,omitempty"`           // Enable IPv6 on VLAN.
	PrefixAssignments *[]RequestApplianceCreateNetworkApplianceVLANIPv6PrefixAssignments `json:"prefixAssignments,omitempty"` // Prefix assignments on the VLAN
}
type RequestApplianceCreateNetworkApplianceVLANIPv6PrefixAssignments struct {
	Autonomous         *bool                                                                  `json:"autonomous,omitempty"`         // Auto assign a /64 prefix from the origin to the VLAN
	Origin             *RequestApplianceCreateNetworkApplianceVLANIPv6PrefixAssignmentsOrigin `json:"origin,omitempty"`             // The origin of the prefix
	StaticApplianceIP6 string                                                                 `json:"staticApplianceIp6,omitempty"` // Manual configuration of the IPv6 Appliance IP
	StaticPrefix       string                                                                 `json:"staticPrefix,omitempty"`       // Manual configuration of a /64 prefix on the VLAN
}
type RequestApplianceCreateNetworkApplianceVLANIPv6PrefixAssignmentsOrigin struct {
	Interfaces []string `json:"interfaces,omitempty"` // Interfaces associated with the prefix
	Type       string   `json:"type,omitempty"`       // Type of the origin
}
type RequestApplianceCreateNetworkApplianceVLANMandatoryDhcp struct {
	Enabled *bool `json:"enabled,omitempty"` // Enable Mandatory DHCP on VLAN.
}
type RequestApplianceUpdateNetworkApplianceVLANsSettings struct {
	VLANsEnabled *bool `json:"vlansEnabled,omitempty"` // Boolean indicating whether to enable (true) or disable (false) VLANs for the network
}
type RequestApplianceUpdateNetworkApplianceVLAN struct {
	ApplianceIP            string                                                        `json:"applianceIp,omitempty"`            // The local IP of the appliance on the VLAN
	Cidr                   string                                                        `json:"cidr,omitempty"`                   // CIDR of the pool of subnets. Applicable only for template network. Each network bound to the template will automatically pick a subnet from this pool to build its own VLAN.
	DhcpBootFilename       string                                                        `json:"dhcpBootFilename,omitempty"`       // DHCP boot option for boot filename
	DhcpBootNextServer     string                                                        `json:"dhcpBootNextServer,omitempty"`     // DHCP boot option to direct boot clients to the server to load the boot file from
	DhcpBootOptionsEnabled *bool                                                         `json:"dhcpBootOptionsEnabled,omitempty"` // Use DHCP boot options specified in other properties
	DhcpHandling           string                                                        `json:"dhcpHandling,omitempty"`           // The appliance's handling of DHCP requests on this VLAN. One of: 'Run a DHCP server', 'Relay DHCP to another server' or 'Do not respond to DHCP requests'
	DhcpLeaseTime          string                                                        `json:"dhcpLeaseTime,omitempty"`          // The term of DHCP leases if the appliance is running a DHCP server on this VLAN. One of: '30 minutes', '1 hour', '4 hours', '12 hours', '1 day' or '1 week'
	DhcpOptions            *[]RequestApplianceUpdateNetworkApplianceVLANDhcpOptions      `json:"dhcpOptions,omitempty"`            // The list of DHCP options that will be included in DHCP responses. Each object in the list should have "code", "type", and "value" properties.
	DhcpRelayServerIPs     []string                                                      `json:"dhcpRelayServerIps,omitempty"`     // The IPs of the DHCP servers that DHCP requests should be relayed to
	DNSNameservers         string                                                        `json:"dnsNameservers,omitempty"`         // The DNS nameservers used for DHCP responses, either "upstream_dns", "google_dns", "opendns", or a newline seperated string of IP addresses or domain names
	FixedIPAssignments     *RequestApplianceUpdateNetworkApplianceVLANFixedIPAssignments `json:"fixedIpAssignments,omitempty"`     // The DHCP fixed IP assignments on the VLAN. This should be an object that contains mappings from MAC addresses to objects that themselves each contain "ip" and "name" string fields. See the sample request/response for more details.
	GroupPolicyID          string                                                        `json:"groupPolicyId,omitempty"`          // The id of the desired group policy to apply to the VLAN
	IPv6                   *RequestApplianceUpdateNetworkApplianceVLANIPv6               `json:"ipv6,omitempty"`                   // IPv6 configuration on the VLAN
	MandatoryDhcp          *RequestApplianceUpdateNetworkApplianceVLANMandatoryDhcp      `json:"mandatoryDhcp,omitempty"`          // Mandatory DHCP will enforce that clients connecting to this VLAN must use the IP address assigned by the DHCP server. Clients who use a static IP address won't be able to associate. Only available on firmware versions 17.0 and above
	Mask                   *int                                                          `json:"mask,omitempty"`                   // Mask used for the subnet of all bound to the template networks. Applicable only for template network.
	Name                   string                                                        `json:"name,omitempty"`                   // The name of the VLAN
	ReservedIPRanges       *[]RequestApplianceUpdateNetworkApplianceVLANReservedIPRanges `json:"reservedIpRanges,omitempty"`       // The DHCP reserved IP ranges on the VLAN
	Subnet                 string                                                        `json:"subnet,omitempty"`                 // The subnet of the VLAN
	TemplateVLANType       string                                                        `json:"templateVlanType,omitempty"`       // Type of subnetting of the VLAN. Applicable only for template network.
	VpnNatSubnet           string                                                        `json:"vpnNatSubnet,omitempty"`           // The translated VPN subnet if VPN and VPN subnet translation are enabled on the VLAN
}
type RequestApplianceUpdateNetworkApplianceVLANDhcpOptions struct {
	Code  string `json:"code,omitempty"`  // The code for the DHCP option. This should be an integer between 2 and 254.
	Type  string `json:"type,omitempty"`  // The type for the DHCP option. One of: 'text', 'ip', 'hex' or 'integer'
	Value string `json:"value,omitempty"` // The value for the DHCP option
}
type RequestApplianceUpdateNetworkApplianceVLANFixedIPAssignments interface{}
type RequestApplianceUpdateNetworkApplianceVLANIPv6 struct {
	Enabled           *bool                                                              `json:"enabled,omitempty"`           // Enable IPv6 on VLAN.
	PrefixAssignments *[]RequestApplianceUpdateNetworkApplianceVLANIPv6PrefixAssignments `json:"prefixAssignments,omitempty"` // Prefix assignments on the VLAN
}
type RequestApplianceUpdateNetworkApplianceVLANIPv6PrefixAssignments struct {
	Autonomous         *bool                                                                  `json:"autonomous,omitempty"`         // Auto assign a /64 prefix from the origin to the VLAN
	Origin             *RequestApplianceUpdateNetworkApplianceVLANIPv6PrefixAssignmentsOrigin `json:"origin,omitempty"`             // The origin of the prefix
	StaticApplianceIP6 string                                                                 `json:"staticApplianceIp6,omitempty"` // Manual configuration of the IPv6 Appliance IP
	StaticPrefix       string                                                                 `json:"staticPrefix,omitempty"`       // Manual configuration of a /64 prefix on the VLAN
}
type RequestApplianceUpdateNetworkApplianceVLANIPv6PrefixAssignmentsOrigin struct {
	Interfaces []string `json:"interfaces,omitempty"` // Interfaces associated with the prefix
	Type       string   `json:"type,omitempty"`       // Type of the origin
}
type RequestApplianceUpdateNetworkApplianceVLANMandatoryDhcp struct {
	Enabled *bool `json:"enabled,omitempty"` // Enable Mandatory DHCP on VLAN.
}
type RequestApplianceUpdateNetworkApplianceVLANReservedIPRanges struct {
	Comment string `json:"comment,omitempty"` // A text comment for the reserved range
	End     string `json:"end,omitempty"`     // The last IP in the reserved range
	Start   string `json:"start,omitempty"`   // The first IP in the reserved range
}
type RequestApplianceUpdateNetworkApplianceVpnBgp struct {
	AsNumber      *int                                                     `json:"asNumber,omitempty"`      // An Autonomous System Number (ASN) is required if you are to run BGP and peer with another BGP Speaker outside of the Auto VPN domain. This ASN will be applied to the entire Auto VPN domain. The entire 4-byte ASN range is supported. So, the ASN must be an integer between 1 and 4294967295. When absent, this field is not updated. If no value exists then it defaults to 64512.
	Enabled       *bool                                                    `json:"enabled,omitempty"`       // Boolean value to enable or disable the BGP configuration. When BGP is enabled, the asNumber (ASN) will be autopopulated with the preconfigured ASN at other Hubs or a default value if there is no ASN configured.
	IbgpHoldTimer *int                                                     `json:"ibgpHoldTimer,omitempty"` // The iBGP holdtimer in seconds. The iBGP holdtimer must be an integer between 12 and 240. When absent, this field is not updated. If no value exists then it defaults to 240.
	Neighbors     *[]RequestApplianceUpdateNetworkApplianceVpnBgpNeighbors `json:"neighbors,omitempty"`     // List of BGP neighbors. This list replaces the existing set of neighbors. When absent, this field is not updated.
}
type RequestApplianceUpdateNetworkApplianceVpnBgpNeighbors struct {
	AllowTransit    *bool                                                                `json:"allowTransit,omitempty"`    // When this feature is on, the Meraki device will advertise routes learned from other Autonomous Systems, thereby allowing traffic between Autonomous Systems to transit this AS. When absent, it defaults to false.
	Authentication  *RequestApplianceUpdateNetworkApplianceVpnBgpNeighborsAuthentication `json:"authentication,omitempty"`  // Authentication settings between BGP peers.
	EbgpHoldTimer   *int                                                                 `json:"ebgpHoldTimer,omitempty"`   // The eBGP hold timer in seconds for each neighbor. The eBGP hold timer must be an integer between 12 and 240.
	EbgpMultihop    *int                                                                 `json:"ebgpMultihop,omitempty"`    // Configure this if the neighbor is not adjacent. The eBGP multi-hop must be an integer between 1 and 255.
	IP              string                                                               `json:"ip,omitempty"`              // The IPv4 address of the neighbor
	IPv6            *RequestApplianceUpdateNetworkApplianceVpnBgpNeighborsIPv6           `json:"ipv6,omitempty"`            // Information regarding IPv6 address of the neighbor, Required if `ip` is not present.
	NextHopIP       string                                                               `json:"nextHopIp,omitempty"`       // The IPv4 address of the remote BGP peer that will establish a TCP session with the local MX.
	ReceiveLimit    *int                                                                 `json:"receiveLimit,omitempty"`    // The receive limit is the maximum number of routes that can be received from any BGP peer. The receive limit must be an integer between 0 and 2147483647. When absent, it defaults to 0.
	RemoteAsNumber  *int                                                                 `json:"remoteAsNumber,omitempty"`  // Remote ASN of the neighbor. The remote ASN must be an integer between 1 and 4294967295.
	SourceInterface string                                                               `json:"sourceInterface,omitempty"` // The output interface for peering with the remote BGP peer. Valid values are: 'wan{NUMBER}' (e.g. 'wan3') or 'vlan{VLAN ID}' (e.g. 'vlan123').
	TtlSecurity     *RequestApplianceUpdateNetworkApplianceVpnBgpNeighborsTtlSecurity    `json:"ttlSecurity,omitempty"`     // Settings for BGP TTL security to protect BGP peering sessions from forged IP attacks.
}
type RequestApplianceUpdateNetworkApplianceVpnBgpNeighborsAuthentication struct {
	Password string `json:"password,omitempty"` // Password to configure MD5 authentication between BGP peers.
}
type RequestApplianceUpdateNetworkApplianceVpnBgpNeighborsIPv6 struct {
	Address string `json:"address,omitempty"` // The IPv6 address of the neighbor.
}
type RequestApplianceUpdateNetworkApplianceVpnBgpNeighborsTtlSecurity struct {
	Enabled *bool `json:"enabled,omitempty"` // Boolean value to enable or disable BGP TTL security.
}
type RequestApplianceUpdateNetworkApplianceVpnSiteToSiteVpn struct {
	Hubs    *[]RequestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnHubs    `json:"hubs,omitempty"`    // The list of VPN hubs, in order of preference. In spoke mode, at least 1 hub is required.
	Mode    string                                                           `json:"mode,omitempty"`    // The site-to-site VPN mode. Can be one of 'none', 'spoke' or 'hub'
	Subnet  *RequestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnet    `json:"subnet,omitempty"`  // Configuration of subnet features
	Subnets *[]RequestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnets `json:"subnets,omitempty"` // The list of subnets and their VPN presence.
}
type RequestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnHubs struct {
	HubID           string `json:"hubId,omitempty"`           // The network ID of the hub.
	UseDefaultRoute *bool  `json:"useDefaultRoute,omitempty"` // Only valid in 'spoke' mode. Indicates whether default route traffic should be sent to this hub.
}
type RequestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnet struct {
	Nat *RequestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnetNat `json:"nat,omitempty"` // Configuration of NAT for subnets
}
type RequestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnetNat struct {
	IsAllowed *bool `json:"isAllowed,omitempty"` // If enabled, VPN subnet translation can be used to translate any local subnets that are allowed to use the VPN into a new subnet with the same number of addresses.
}
type RequestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnets struct {
	LocalSubnet string                                                            `json:"localSubnet,omitempty"` // The CIDR notation subnet used within the VPN
	Nat         *RequestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnetsNat `json:"nat,omitempty"`         // Configuration of NAT for the subnet
	UseVpn      *bool                                                             `json:"useVpn,omitempty"`      // Indicates the presence of the subnet in the VPN
}
type RequestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnetsNat struct {
	Enabled      *bool  `json:"enabled,omitempty"`      // Whether or not VPN subnet translation is enabled for the subnet
	RemoteSubnet string `json:"remoteSubnet,omitempty"` // The translated subnet to be used in the VPN
}
type RequestApplianceUpdateNetworkApplianceWarmSpare struct {
	Enabled     *bool  `json:"enabled,omitempty"`     // Enable warm spare
	SpareSerial string `json:"spareSerial,omitempty"` // Serial number of the warm spare appliance
	UplinkMode  string `json:"uplinkMode,omitempty"`  // Uplink mode, either virtual or public
	VirtualIP1  string `json:"virtualIp1,omitempty"`  // The WAN 1 shared IP
	VirtualIP2  string `json:"virtualIp2,omitempty"`  // The WAN 2 shared IP
}
type RequestApplianceCreateOrganizationApplianceDNSLocalProfile struct {
	Name string `json:"name,omitempty"` // Name of profile
}
type RequestApplianceBulkOrganizationApplianceDNSLocalProfilesAssignmentsCreate struct {
	Items *[]RequestApplianceBulkOrganizationApplianceDNSLocalProfilesAssignmentsCreateItems `json:"items,omitempty"` // List containing the network ID and Profile ID
}
type RequestApplianceBulkOrganizationApplianceDNSLocalProfilesAssignmentsCreateItems struct {
	Network *RequestApplianceBulkOrganizationApplianceDNSLocalProfilesAssignmentsCreateItemsNetwork `json:"network,omitempty"` // The network attached to the profile
	Profile *RequestApplianceBulkOrganizationApplianceDNSLocalProfilesAssignmentsCreateItemsProfile `json:"profile,omitempty"` // The profile the network is attached to
}
type RequestApplianceBulkOrganizationApplianceDNSLocalProfilesAssignmentsCreateItemsNetwork struct {
	ID string `json:"id,omitempty"` // ID of the network
}
type RequestApplianceBulkOrganizationApplianceDNSLocalProfilesAssignmentsCreateItemsProfile struct {
	ID string `json:"id,omitempty"` // ID of the profile
}
type RequestApplianceCreateOrganizationApplianceDNSLocalProfilesAssignmentsBulkDelete struct {
	Items *[]RequestApplianceCreateOrganizationApplianceDNSLocalProfilesAssignmentsBulkDeleteItems `json:"items,omitempty"` // List containing the assignment ID
}
type RequestApplianceCreateOrganizationApplianceDNSLocalProfilesAssignmentsBulkDeleteItems struct {
	AssignmentID string `json:"assignmentId,omitempty"` // ID of the assignment
}
type RequestApplianceUpdateOrganizationApplianceDNSLocalProfile struct {
	Name string `json:"name,omitempty"` // Name of profile
}
type RequestApplianceCreateOrganizationApplianceDNSLocalRecord struct {
	Address  string                                                            `json:"address,omitempty"`  // IP for the DNS record
	Hostname string                                                            `json:"hostname,omitempty"` // Hostname for the DNS record
	Profile  *RequestApplianceCreateOrganizationApplianceDNSLocalRecordProfile `json:"profile,omitempty"`  // The profile the DNS record is associated with
}
type RequestApplianceCreateOrganizationApplianceDNSLocalRecordProfile struct {
	ID string `json:"id,omitempty"` // Profile ID
}
type RequestApplianceUpdateOrganizationApplianceDNSLocalRecord struct {
	Address  string                                                            `json:"address,omitempty"`  // IP for the DNS record
	Hostname string                                                            `json:"hostname,omitempty"` // Hostname for the DNS record
	Profile  *RequestApplianceUpdateOrganizationApplianceDNSLocalRecordProfile `json:"profile,omitempty"`  // The profile the DNS record is associated with
}
type RequestApplianceUpdateOrganizationApplianceDNSLocalRecordProfile struct {
	ID string `json:"id,omitempty"` // Profile ID
}
type RequestApplianceCreateOrganizationApplianceDNSSplitProfile struct {
	Hostnames   []string                                                               `json:"hostnames,omitempty"`   // The hostname patterns to match for redirection. For more information on Split DNS hostname pattern formatting, please consult the Split DNS KB.
	Name        string                                                                 `json:"name,omitempty"`        // Name of profile
	Nameservers *RequestApplianceCreateOrganizationApplianceDNSSplitProfileNameservers `json:"nameservers,omitempty"` // Contains the nameserver information for redirection.
}
type RequestApplianceCreateOrganizationApplianceDNSSplitProfileNameservers struct {
	Addresses []string `json:"addresses,omitempty"` // The nameserver address(es) to use for redirection. A maximum of one address is supported.
}
type RequestApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreate struct {
	Items *[]RequestApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreateItems `json:"items,omitempty"` // List containing the network ID and Profile ID
}
type RequestApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreateItems struct {
	Network *RequestApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreateItemsNetwork `json:"network,omitempty"` // The network attached to the profile
	Profile *RequestApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreateItemsProfile `json:"profile,omitempty"` // The profile the network is attached to
}
type RequestApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreateItemsNetwork struct {
	ID string `json:"id,omitempty"` // ID of the network
}
type RequestApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreateItemsProfile struct {
	ID string `json:"id,omitempty"` // ID of the profile
}
type RequestApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkDelete struct {
	Items *[]RequestApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkDeleteItems `json:"items,omitempty"` // List containing the assignment ID
}
type RequestApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkDeleteItems struct {
	AssignmentID string `json:"assignmentId,omitempty"` // ID of the assignment
}
type RequestApplianceUpdateOrganizationApplianceDNSSplitProfile struct {
	Hostnames   []string                                                               `json:"hostnames,omitempty"`   // The hostname patterns to match for redirection. For more information on Split DNS hostname pattern formatting, please consult the Split DNS KB.
	Name        string                                                                 `json:"name,omitempty"`        // Name of profile
	Nameservers *RequestApplianceUpdateOrganizationApplianceDNSSplitProfileNameservers `json:"nameservers,omitempty"` // Contains the nameserver information for redirection.
}
type RequestApplianceUpdateOrganizationApplianceDNSSplitProfileNameservers struct {
	Addresses []string `json:"addresses,omitempty"` // The nameserver address(es) to use for redirection. A maximum of one address is supported.
}
type RequestApplianceUpdateOrganizationApplianceSecurityIntrusion struct {
	AllowedRules *[]RequestApplianceUpdateOrganizationApplianceSecurityIntrusionAllowedRules `json:"allowedRules,omitempty"` // Sets a list of specific SNORT signatures to allow
}
type RequestApplianceUpdateOrganizationApplianceSecurityIntrusionAllowedRules struct {
	Message string `json:"message,omitempty"` // Message is optional and is ignored on a PUT call. It is allowed in order for PUT to be compatible with GET
	RuleID  string `json:"ruleId,omitempty"`  // A rule identifier of the format meraki:intrusion/snort/GID/<gid>/SID/<sid>. gid and sid can be obtained from either https://www.snort.org/rule-docs or as ruleIds from the security events in /organization/[orgId]/securityEvents
}
type RequestApplianceUpdateOrganizationApplianceVpnThirdPartyVpnpeers struct {
	Peers *[]RequestApplianceUpdateOrganizationApplianceVpnThirdPartyVpnpeersPeers `json:"peers,omitempty"` // The list of VPN peers
}
type RequestApplianceUpdateOrganizationApplianceVpnThirdPartyVpnpeersPeers struct {
	IkeVersion          string                                                                              `json:"ikeVersion,omitempty"`          // [optional] The IKE version to be used for the IPsec VPN peer configuration. Defaults to '1' when omitted.
	IPsecPolicies       *RequestApplianceUpdateOrganizationApplianceVpnThirdPartyVpnpeersPeersIPsecPolicies `json:"ipsecPolicies,omitempty"`       // Custom IPSec policies for the VPN peer. If not included and a preset has not been chosen, the default preset for IPSec policies will be used.
	IPsecPoliciesPreset string                                                                              `json:"ipsecPoliciesPreset,omitempty"` // One of the following available presets: 'default', 'aws', 'azure', 'umbrella', 'zscaler'. If this is provided, the 'ipsecPolicies' parameter is ignored.
	LocalID             string                                                                              `json:"localId,omitempty"`             // [optional] The local ID is used to identify the MX to the peer. This will apply to all MXs this peer applies to.
	Name                string                                                                              `json:"name,omitempty"`                // The name of the VPN peer
	NetworkTags         []string                                                                            `json:"networkTags,omitempty"`         // A list of network tags that will connect with this peer. Use ['all'] for all networks. Use ['none'] for no networks. If not included, the default is ['all'].
	PrivateSubnets      []string                                                                            `json:"privateSubnets,omitempty"`      // The list of the private subnets of the VPN peer
	PublicHostname      string                                                                              `json:"publicHostname,omitempty"`      // [optional] The public hostname of the VPN peer
	PublicIP            string                                                                              `json:"publicIp,omitempty"`            // [optional] The public IP of the VPN peer
	RemoteID            string                                                                              `json:"remoteId,omitempty"`            // [optional] The remote ID is used to identify the connecting VPN peer. This can either be a valid IPv4 Address, FQDN or User FQDN.
	Secret              string                                                                              `json:"secret,omitempty"`              // The shared secret with the VPN peer
}
type RequestApplianceUpdateOrganizationApplianceVpnThirdPartyVpnpeersPeersIPsecPolicies struct {
	ChildAuthAlgo         []string `json:"childAuthAlgo,omitempty"`         // This is the authentication algorithms to be used in Phase 2. The value should be an array with one of the following algorithms: 'sha256', 'sha1', 'md5'
	ChildCipherAlgo       []string `json:"childCipherAlgo,omitempty"`       // This is the cipher algorithms to be used in Phase 2. The value should be an array with one or more of the following algorithms: 'aes256', 'aes192', 'aes128', 'tripledes', 'des', 'null'
	ChildLifetime         *int     `json:"childLifetime,omitempty"`         // The lifetime of the Phase 2 SA in seconds.
	ChildPfsGroup         []string `json:"childPfsGroup,omitempty"`         // This is the Diffie-Hellman group to be used for Perfect Forward Secrecy in Phase 2. The value should be an array with one of the following values: 'disabled','group14', 'group5', 'group2', 'group1'
	IkeAuthAlgo           []string `json:"ikeAuthAlgo,omitempty"`           // This is the authentication algorithm to be used in Phase 1. The value should be an array with one of the following algorithms: 'sha256', 'sha1', 'md5'
	IkeCipherAlgo         []string `json:"ikeCipherAlgo,omitempty"`         // This is the cipher algorithm to be used in Phase 1. The value should be an array with one of the following algorithms: 'aes256', 'aes192', 'aes128', 'tripledes', 'des'
	IkeDiffieHellmanGroup []string `json:"ikeDiffieHellmanGroup,omitempty"` // This is the Diffie-Hellman group to be used in Phase 1. The value should be an array with one of the following algorithms: 'group14', 'group5', 'group2', 'group1'
	IkeLifetime           *int     `json:"ikeLifetime,omitempty"`           // The lifetime of the Phase 1 SA in seconds.
	IkePrfAlgo            []string `json:"ikePrfAlgo,omitempty"`            // [optional] This is the pseudo-random function to be used in IKE_SA. The value should be an array with one of the following algorithms: 'prfsha256', 'prfsha1', 'prfmd5', 'default'. The 'default' option can be used to default to the Authentication algorithm.
}
type RequestApplianceUpdateOrganizationApplianceVpnVpnFirewallRules struct {
	Rules             *[]RequestApplianceUpdateOrganizationApplianceVpnVpnFirewallRulesRules `json:"rules,omitempty"`             // An ordered array of the firewall rules (not including the default rule)
	SyslogDefaultRule *bool                                                                  `json:"syslogDefaultRule,omitempty"` // Log the special default rule (boolean value - enable only if you've configured a syslog server) (optional)
}
type RequestApplianceUpdateOrganizationApplianceVpnVpnFirewallRulesRules struct {
	Comment       string `json:"comment,omitempty"`       // Description of the rule (optional)
	DestCidr      string `json:"destCidr,omitempty"`      // Comma-separated list of destination IP address(es) (in IP or CIDR notation) or 'any' (FQDN not supported)
	DestPort      string `json:"destPort,omitempty"`      // Comma-separated list of destination port(s) (integer in the range 1-65535), or 'any'
	Policy        string `json:"policy,omitempty"`        // 'allow' or 'deny' traffic specified by this rule
	Protocol      string `json:"protocol,omitempty"`      // The type of protocol (must be 'tcp', 'udp', 'icmp', 'icmp6' or 'any')
	SrcCidr       string `json:"srcCidr,omitempty"`       // Comma-separated list of source IP address(es) (in IP or CIDR notation), or 'any' (FQDN not supported)
	SrcPort       string `json:"srcPort,omitempty"`       // Comma-separated list of source port(s) (integer in the range 1-65535), or 'any'
	SyslogEnabled *bool  `json:"syslogEnabled,omitempty"` // Log this rule to syslog (true or false, boolean value) - only applicable if a syslog has been configured (optional)
}

//GetDeviceApplianceDhcpSubnets Return the DHCP subnet information for an appliance
/* Return the DHCP subnet information for an appliance

@param serial serial path parameter.


*/

func (s *ApplianceService) GetDeviceApplianceDhcpSubnets(serial string) (*ResponseApplianceGetDeviceApplianceDhcpSubnets, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/appliance/dhcp/subnets"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetDeviceApplianceDhcpSubnets](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetDeviceAppliancePerformance Return the performance score for a single MX
/* Return the performance score for a single MX. Only primary MX devices supported. If no data is available, a 204 error code is returned.

@param serial serial path parameter.
@param getDeviceAppliancePerformanceQueryParams Filtering parameter


*/

func (s *ApplianceService) GetDeviceAppliancePerformance(serial string, getDeviceAppliancePerformanceQueryParams *GetDeviceAppliancePerformanceQueryParams) (*ResponseApplianceGetDeviceAppliancePerformance, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/appliance/performance"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetDeviceAppliancePerformance](
		func() (*resty.Response, error) {
			return GET(path, s.client, getDeviceAppliancePerformanceQueryParams, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetDeviceAppliancePrefixesDelegated Return current delegated IPv6 prefixes on an appliance.
/* Return current delegated IPv6 prefixes on an appliance.

@param serial serial path parameter.


*/

func (s *ApplianceService) GetDeviceAppliancePrefixesDelegated(serial string) (*ResponseApplianceGetDeviceAppliancePrefixesDelegated, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/appliance/prefixes/delegated"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetDeviceAppliancePrefixesDelegated](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetDeviceAppliancePrefixesDelegatedVLANAssignments Return prefixes assigned to all IPv6 enabled VLANs on an appliance.
/* Return prefixes assigned to all IPv6 enabled VLANs on an appliance.

@param serial serial path parameter.


*/

func (s *ApplianceService) GetDeviceAppliancePrefixesDelegatedVLANAssignments(serial string) (*ResponseApplianceGetDeviceAppliancePrefixesDelegatedVLANAssignments, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/appliance/prefixes/delegated/vlanAssignments"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetDeviceAppliancePrefixesDelegatedVLANAssignments](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetDeviceApplianceRadioSettings Return the radio settings of an appliance
/* Return the radio settings of an appliance

@param serial serial path parameter.


*/

func (s *ApplianceService) GetDeviceApplianceRadioSettings(serial string) (*ResponseApplianceGetDeviceApplianceRadioSettings, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/appliance/radio/settings"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetDeviceApplianceRadioSettings](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetDeviceApplianceUplinksSettings Return the uplink settings for an MX appliance
/* Return the uplink settings for an MX appliance

@param serial serial path parameter.


*/

func (s *ApplianceService) GetDeviceApplianceUplinksSettings(serial string) (*ResponseApplianceGetDeviceApplianceUplinksSettings, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/appliance/uplinks/settings"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetDeviceApplianceUplinksSettings](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetNetworkApplianceClientSecurityEvents List the security events for a client
/* List the security events for a client. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.

@param networkID networkId path parameter. Network ID
@param clientID clientId path parameter. Client ID
@param getNetworkApplianceClientSecurityEventsQueryParams Filtering parameter


*/

func (s *ApplianceService) GetNetworkApplianceClientSecurityEvents(networkID string, clientID string, getNetworkApplianceClientSecurityEventsQueryParams *GetNetworkApplianceClientSecurityEventsQueryParams) (*ResponseApplianceGetNetworkApplianceClientSecurityEvents, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/clients/{clientId}/security/events"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{clientId}", fmt.Sprintf("%v", clientID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkApplianceClientSecurityEvents](
		func() (*resty.Response, error) {
			return GET(path, s.client, getNetworkApplianceClientSecurityEventsQueryParams, &HeaderDefault)
		},
		s.client,
		func(dst, src ResponseApplianceGetNetworkApplianceClientSecurityEvents) ResponseApplianceGetNetworkApplianceClientSecurityEvents {
			dst = append(dst, src...)
			return dst
		},
		func() bool {
			if getNetworkApplianceClientSecurityEventsQueryParams != nil {
				return getNetworkApplianceClientSecurityEventsQueryParams.PerPage == -1
			}
			return false
		}(),
	)

}

//GetNetworkApplianceConnectivityMonitoringDestinations Return the connectivity testing destinations for an MX network
/* Return the connectivity testing destinations for an MX network

@param networkID networkId path parameter. Network ID


*/

func (s *ApplianceService) GetNetworkApplianceConnectivityMonitoringDestinations(networkID string) (*ResponseApplianceGetNetworkApplianceConnectivityMonitoringDestinations, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/connectivityMonitoringDestinations"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkApplianceConnectivityMonitoringDestinations](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetNetworkApplianceContentFiltering Return the content filtering settings for an MX network
/* Return the content filtering settings for an MX network

@param networkID networkId path parameter. Network ID


*/

func (s *ApplianceService) GetNetworkApplianceContentFiltering(networkID string) (*ResponseApplianceGetNetworkApplianceContentFiltering, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/contentFiltering"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkApplianceContentFiltering](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetNetworkApplianceContentFilteringCategories List all available content filtering categories for an MX network
/* List all available content filtering categories for an MX network

@param networkID networkId path parameter. Network ID


*/

func (s *ApplianceService) GetNetworkApplianceContentFilteringCategories(networkID string) (*ResponseApplianceGetNetworkApplianceContentFilteringCategories, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/contentFiltering/categories"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkApplianceContentFilteringCategories](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetNetworkApplianceFirewallCellularFirewallRules Return the cellular firewall rules for an MX network
/* Return the cellular firewall rules for an MX network

@param networkID networkId path parameter. Network ID


*/

func (s *ApplianceService) GetNetworkApplianceFirewallCellularFirewallRules(networkID string) (*ResponseApplianceGetNetworkApplianceFirewallCellularFirewallRules, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/cellularFirewallRules"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkApplianceFirewallCellularFirewallRules](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetNetworkApplianceFirewallFirewalledServices List the appliance services and their accessibility rules
/* List the appliance services and their accessibility rules

@param networkID networkId path parameter. Network ID


*/

func (s *ApplianceService) GetNetworkApplianceFirewallFirewalledServices(networkID string) (*ResponseApplianceGetNetworkApplianceFirewallFirewalledServices, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/firewalledServices"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkApplianceFirewallFirewalledServices](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetNetworkApplianceFirewallFirewalledService Return the accessibility settings of the given service ('ICMP', 'web', or 'SNMP')
/* Return the accessibility settings of the given service ('ICMP', 'web', or 'SNMP')

@param networkID networkId path parameter. Network ID
@param service service path parameter.


*/

func (s *ApplianceService) GetNetworkApplianceFirewallFirewalledService(networkID string, service string) (*ResponseApplianceGetNetworkApplianceFirewallFirewalledService, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/firewalledServices/{service}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{service}", fmt.Sprintf("%v", service), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkApplianceFirewallFirewalledService](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetNetworkApplianceFirewallInboundCellularFirewallRules Return the inbound cellular firewall rules for an MX network
/* Return the inbound cellular firewall rules for an MX network

@param networkID networkId path parameter. Network ID


*/

func (s *ApplianceService) GetNetworkApplianceFirewallInboundCellularFirewallRules(networkID string) (*ResponseApplianceGetNetworkApplianceFirewallInboundCellularFirewallRules, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/inboundCellularFirewallRules"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkApplianceFirewallInboundCellularFirewallRules](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetNetworkApplianceFirewallInboundFirewallRules Return the inbound firewall rules for an MX network
/* Return the inbound firewall rules for an MX network

@param networkID networkId path parameter. Network ID


*/

func (s *ApplianceService) GetNetworkApplianceFirewallInboundFirewallRules(networkID string) (*ResponseApplianceGetNetworkApplianceFirewallInboundFirewallRules, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/inboundFirewallRules"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkApplianceFirewallInboundFirewallRules](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetNetworkApplianceFirewallL3FirewallRules Return the L3 firewall rules for an MX network
/* Return the L3 firewall rules for an MX network

@param networkID networkId path parameter. Network ID


*/

func (s *ApplianceService) GetNetworkApplianceFirewallL3FirewallRules(networkID string) (*ResponseApplianceGetNetworkApplianceFirewallL3FirewallRules, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/l3FirewallRules"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkApplianceFirewallL3FirewallRules](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetNetworkApplianceFirewallL7FirewallRules List the MX L7 firewall rules for an MX network
/* List the MX L7 firewall rules for an MX network

@param networkID networkId path parameter. Network ID


*/

func (s *ApplianceService) GetNetworkApplianceFirewallL7FirewallRules(networkID string) (*ResponseApplianceGetNetworkApplianceFirewallL7FirewallRules, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/l7FirewallRules"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkApplianceFirewallL7FirewallRules](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories Return the L7 firewall application categories and their associated applications for an MX network
/* Return the L7 firewall application categories and their associated applications for an MX network

@param networkID networkId path parameter. Network ID


*/

func (s *ApplianceService) GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories(networkID string) (*ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/l7FirewallRules/applicationCategories"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetNetworkApplianceFirewallOneToManyNatRules Return the 1:Many NAT mapping rules for an MX network
/* Return the 1:Many NAT mapping rules for an MX network

@param networkID networkId path parameter. Network ID


*/

func (s *ApplianceService) GetNetworkApplianceFirewallOneToManyNatRules(networkID string) (*ResponseApplianceGetNetworkApplianceFirewallOneToManyNatRules, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/oneToManyNatRules"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkApplianceFirewallOneToManyNatRules](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetNetworkApplianceFirewallOneToOneNatRules Return the 1:1 NAT mapping rules for an MX network
/* Return the 1:1 NAT mapping rules for an MX network

@param networkID networkId path parameter. Network ID


*/

func (s *ApplianceService) GetNetworkApplianceFirewallOneToOneNatRules(networkID string) (*ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRules, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/oneToOneNatRules"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRules](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetNetworkApplianceFirewallPortForwardingRules Return the port forwarding rules for an MX network
/* Return the port forwarding rules for an MX network

@param networkID networkId path parameter. Network ID


*/

func (s *ApplianceService) GetNetworkApplianceFirewallPortForwardingRules(networkID string) (*ResponseApplianceGetNetworkApplianceFirewallPortForwardingRules, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/portForwardingRules"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkApplianceFirewallPortForwardingRules](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetNetworkApplianceFirewallSettings Return the firewall settings for this network
/* Return the firewall settings for this network

@param networkID networkId path parameter. Network ID


*/

func (s *ApplianceService) GetNetworkApplianceFirewallSettings(networkID string) (*ResponseApplianceGetNetworkApplianceFirewallSettings, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/settings"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkApplianceFirewallSettings](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetNetworkAppliancePorts List per-port VLAN settings for all ports of a MX.
/* List per-port VLAN settings for all ports of a MX.

@param networkID networkId path parameter. Network ID


*/

func (s *ApplianceService) GetNetworkAppliancePorts(networkID string) (*ResponseApplianceGetNetworkAppliancePorts, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/ports"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkAppliancePorts](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetNetworkAppliancePort Return per-port VLAN settings for a single MX port.
/* Return per-port VLAN settings for a single MX port.

@param networkID networkId path parameter. Network ID
@param portID portId path parameter. Port ID


*/

func (s *ApplianceService) GetNetworkAppliancePort(networkID string, portID string) (*ResponseApplianceGetNetworkAppliancePort, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/ports/{portId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{portId}", fmt.Sprintf("%v", portID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkAppliancePort](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetNetworkAppliancePrefixesDelegatedStatics List static delegated prefixes for a network
/* List static delegated prefixes for a network

@param networkID networkId path parameter. Network ID


*/

func (s *ApplianceService) GetNetworkAppliancePrefixesDelegatedStatics(networkID string) (*ResponseApplianceGetNetworkAppliancePrefixesDelegatedStatics, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/prefixes/delegated/statics"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkAppliancePrefixesDelegatedStatics](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetNetworkAppliancePrefixesDelegatedStatic Return a static delegated prefix from a network
/* Return a static delegated prefix from a network

@param networkID networkId path parameter. Network ID
@param staticDelegatedPrefixID staticDelegatedPrefixId path parameter. Static delegated prefix ID


*/

func (s *ApplianceService) GetNetworkAppliancePrefixesDelegatedStatic(networkID string, staticDelegatedPrefixID string) (*ResponseApplianceGetNetworkAppliancePrefixesDelegatedStatic, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/prefixes/delegated/statics/{staticDelegatedPrefixId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{staticDelegatedPrefixId}", fmt.Sprintf("%v", staticDelegatedPrefixID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkAppliancePrefixesDelegatedStatic](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetNetworkApplianceRfProfiles List the RF profiles for this network
/* List the RF profiles for this network

@param networkID networkId path parameter. Network ID


*/

func (s *ApplianceService) GetNetworkApplianceRfProfiles(networkID string) (*ResponseApplianceGetNetworkApplianceRfProfiles, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/rfProfiles"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkApplianceRfProfiles](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetNetworkApplianceRfProfile Return a RF profile
/* Return a RF profile

@param networkID networkId path parameter. Network ID
@param rfProfileID rfProfileId path parameter. Rf profile ID


*/

func (s *ApplianceService) GetNetworkApplianceRfProfile(networkID string, rfProfileID string) (*ResponseApplianceGetNetworkApplianceRfProfile, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/rfProfiles/{rfProfileId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{rfProfileId}", fmt.Sprintf("%v", rfProfileID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkApplianceRfProfile](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetNetworkApplianceSecurityEvents List the security events for a network
/* List the security events for a network

@param networkID networkId path parameter. Network ID
@param getNetworkApplianceSecurityEventsQueryParams Filtering parameter


*/

func (s *ApplianceService) GetNetworkApplianceSecurityEvents(networkID string, getNetworkApplianceSecurityEventsQueryParams *GetNetworkApplianceSecurityEventsQueryParams) (*ResponseApplianceGetNetworkApplianceSecurityEvents, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/security/events"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkApplianceSecurityEvents](
		func() (*resty.Response, error) {
			return GET(path, s.client, getNetworkApplianceSecurityEventsQueryParams, &HeaderDefault)
		},
		s.client,
		func(dst, src ResponseApplianceGetNetworkApplianceSecurityEvents) ResponseApplianceGetNetworkApplianceSecurityEvents {
			dst = append(dst, src...)
			return dst
		},
		func() bool {
			if getNetworkApplianceSecurityEventsQueryParams != nil {
				return getNetworkApplianceSecurityEventsQueryParams.PerPage == -1
			}
			return false
		}(),
	)

}

//GetNetworkApplianceSecurityIntrusion Returns all supported intrusion settings for an MX network
/* Returns all supported intrusion settings for an MX network

@param networkID networkId path parameter. Network ID


*/

func (s *ApplianceService) GetNetworkApplianceSecurityIntrusion(networkID string) (*ResponseApplianceGetNetworkApplianceSecurityIntrusion, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/security/intrusion"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkApplianceSecurityIntrusion](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetNetworkApplianceSecurityMalware Returns all supported malware settings for an MX network
/* Returns all supported malware settings for an MX network

@param networkID networkId path parameter. Network ID


*/

func (s *ApplianceService) GetNetworkApplianceSecurityMalware(networkID string) (*ResponseApplianceGetNetworkApplianceSecurityMalware, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/security/malware"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkApplianceSecurityMalware](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetNetworkApplianceSettings Return the appliance settings for a network
/* Return the appliance settings for a network

@param networkID networkId path parameter. Network ID


*/

func (s *ApplianceService) GetNetworkApplianceSettings(networkID string) (*ResponseApplianceGetNetworkApplianceSettings, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/settings"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkApplianceSettings](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetNetworkApplianceSingleLan Return single LAN configuration
/* Return single LAN configuration

@param networkID networkId path parameter. Network ID


*/

func (s *ApplianceService) GetNetworkApplianceSingleLan(networkID string) (*ResponseApplianceGetNetworkApplianceSingleLan, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/singleLan"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkApplianceSingleLan](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetNetworkApplianceSSIDs List the MX SSIDs in a network
/* List the MX SSIDs in a network

@param networkID networkId path parameter. Network ID


*/

func (s *ApplianceService) GetNetworkApplianceSSIDs(networkID string) (*ResponseApplianceGetNetworkApplianceSSIDs, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/ssids"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkApplianceSSIDs](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetNetworkApplianceSSID Return a single MX SSID
/* Return a single MX SSID

@param networkID networkId path parameter. Network ID
@param number number path parameter.


*/

func (s *ApplianceService) GetNetworkApplianceSSID(networkID string, number string) (*ResponseApplianceGetNetworkApplianceSSID, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/ssids/{number}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{number}", fmt.Sprintf("%v", number), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkApplianceSSID](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetNetworkApplianceStaticRoutes List the static routes for an MX or teleworker network
/* List the static routes for an MX or teleworker network

@param networkID networkId path parameter. Network ID


*/

func (s *ApplianceService) GetNetworkApplianceStaticRoutes(networkID string) (*ResponseApplianceGetNetworkApplianceStaticRoutes, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/staticRoutes"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkApplianceStaticRoutes](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetNetworkApplianceStaticRoute Return a static route for an MX or teleworker network
/* Return a static route for an MX or teleworker network

@param networkID networkId path parameter. Network ID
@param staticRouteID staticRouteId path parameter. Static route ID


*/

func (s *ApplianceService) GetNetworkApplianceStaticRoute(networkID string, staticRouteID string) (*ResponseApplianceGetNetworkApplianceStaticRoute, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/staticRoutes/{staticRouteId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{staticRouteId}", fmt.Sprintf("%v", staticRouteID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkApplianceStaticRoute](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetNetworkApplianceTrafficShaping Display the traffic shaping settings for an MX network
/* Display the traffic shaping settings for an MX network

@param networkID networkId path parameter. Network ID


*/

func (s *ApplianceService) GetNetworkApplianceTrafficShaping(networkID string) (*ResponseApplianceGetNetworkApplianceTrafficShaping, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/trafficShaping"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkApplianceTrafficShaping](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetNetworkApplianceTrafficShapingCustomPerformanceClasses List all custom performance classes for an MX network
/* List all custom performance classes for an MX network

@param networkID networkId path parameter. Network ID


*/

func (s *ApplianceService) GetNetworkApplianceTrafficShapingCustomPerformanceClasses(networkID string) (*ResponseApplianceGetNetworkApplianceTrafficShapingCustomPerformanceClasses, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/trafficShaping/customPerformanceClasses"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkApplianceTrafficShapingCustomPerformanceClasses](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetNetworkApplianceTrafficShapingCustomPerformanceClass Return a custom performance class for an MX network
/* Return a custom performance class for an MX network

@param networkID networkId path parameter. Network ID
@param customPerformanceClassID customPerformanceClassId path parameter. Custom performance class ID


*/

func (s *ApplianceService) GetNetworkApplianceTrafficShapingCustomPerformanceClass(networkID string, customPerformanceClassID string) (*ResponseApplianceGetNetworkApplianceTrafficShapingCustomPerformanceClass, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/trafficShaping/customPerformanceClasses/{customPerformanceClassId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{customPerformanceClassId}", fmt.Sprintf("%v", customPerformanceClassID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkApplianceTrafficShapingCustomPerformanceClass](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetNetworkApplianceTrafficShapingRules Display the traffic shaping settings rules for an MX network
/* Display the traffic shaping settings rules for an MX network

@param networkID networkId path parameter. Network ID


*/

func (s *ApplianceService) GetNetworkApplianceTrafficShapingRules(networkID string) (*ResponseApplianceGetNetworkApplianceTrafficShapingRules, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/trafficShaping/rules"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkApplianceTrafficShapingRules](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetNetworkApplianceTrafficShapingUplinkBandwidth Returns the uplink bandwidth limits for your MX network
/* Returns the uplink bandwidth limits for your MX network. This may not reflect the affected device's hardware capabilities.  For more information on your device's hardware capabilities, please consult our MX Family Datasheet [https://meraki.cisco.com/product-collateral/mx-family-datasheet/?file]

@param networkID networkId path parameter. Network ID


*/

func (s *ApplianceService) GetNetworkApplianceTrafficShapingUplinkBandwidth(networkID string) (*ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidth, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/trafficShaping/uplinkBandwidth"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidth](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetNetworkApplianceTrafficShapingUplinkSelection Show uplink selection settings for an MX network
/* Show uplink selection settings for an MX network

@param networkID networkId path parameter. Network ID


*/

func (s *ApplianceService) GetNetworkApplianceTrafficShapingUplinkSelection(networkID string) (*ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelection, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/trafficShaping/uplinkSelection"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelection](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetNetworkApplianceUplinksUsageHistory Get the sent and received bytes for each uplink of a network.
/* Get the sent and received bytes for each uplink of a network.

@param networkID networkId path parameter. Network ID
@param getNetworkApplianceUplinksUsageHistoryQueryParams Filtering parameter


*/

func (s *ApplianceService) GetNetworkApplianceUplinksUsageHistory(networkID string, getNetworkApplianceUplinksUsageHistoryQueryParams *GetNetworkApplianceUplinksUsageHistoryQueryParams) (*ResponseApplianceGetNetworkApplianceUplinksUsageHistory, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/uplinks/usageHistory"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkApplianceUplinksUsageHistory](
		func() (*resty.Response, error) {
			return GET(path, s.client, getNetworkApplianceUplinksUsageHistoryQueryParams, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetNetworkApplianceVLANs List the VLANs for an MX network
/* List the VLANs for an MX network

@param networkID networkId path parameter. Network ID


*/

func (s *ApplianceService) GetNetworkApplianceVLANs(networkID string) (*ResponseApplianceGetNetworkApplianceVLANs, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/vlans"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkApplianceVLANs](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetNetworkApplianceVLANsSettings Returns the enabled status of VLANs for the network
/* Returns the enabled status of VLANs for the network

@param networkID networkId path parameter. Network ID


*/

func (s *ApplianceService) GetNetworkApplianceVLANsSettings(networkID string) (*ResponseApplianceGetNetworkApplianceVLANsSettings, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/vlans/settings"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkApplianceVLANsSettings](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetNetworkApplianceVLAN Return a VLAN
/* Return a VLAN

@param networkID networkId path parameter. Network ID
@param vlanID vlanId path parameter. Vlan ID


*/

func (s *ApplianceService) GetNetworkApplianceVLAN(networkID string, vlanID string) (*ResponseApplianceGetNetworkApplianceVLAN, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/vlans/{vlanId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{vlanId}", fmt.Sprintf("%v", vlanID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkApplianceVLAN](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetNetworkApplianceVpnBgp Return a Hub BGP Configuration
/* Return a Hub BGP Configuration

@param networkID networkId path parameter. Network ID


*/

func (s *ApplianceService) GetNetworkApplianceVpnBgp(networkID string) (*ResponseApplianceGetNetworkApplianceVpnBgp, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/vpn/bgp"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkApplianceVpnBgp](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetNetworkApplianceVpnSiteToSiteVpn Return the site-to-site VPN settings of a network
/* Return the site-to-site VPN settings of a network. Only valid for MX networks.

@param networkID networkId path parameter. Network ID


*/

func (s *ApplianceService) GetNetworkApplianceVpnSiteToSiteVpn(networkID string) (*ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpn, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/vpn/siteToSiteVpn"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpn](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetNetworkApplianceWarmSpare Return MX warm spare settings
/* Return MX warm spare settings

@param networkID networkId path parameter. Network ID


*/

func (s *ApplianceService) GetNetworkApplianceWarmSpare(networkID string) (*ResponseApplianceGetNetworkApplianceWarmSpare, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/warmSpare"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetNetworkApplianceWarmSpare](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetOrganizationApplianceDNSLocalProfiles Fetch the local DNS profiles used in the organization
/* Fetch the local DNS profiles used in the organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationApplianceDnsLocalProfilesQueryParams Filtering parameter


*/

func (s *ApplianceService) GetOrganizationApplianceDNSLocalProfiles(organizationID string, getOrganizationApplianceDnsLocalProfilesQueryParams *GetOrganizationApplianceDNSLocalProfilesQueryParams) (*ResponseApplianceGetOrganizationApplianceDNSLocalProfiles, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/appliance/dns/local/profiles"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetOrganizationApplianceDNSLocalProfiles](
		func() (*resty.Response, error) {
			return GET(path, s.client, getOrganizationApplianceDnsLocalProfilesQueryParams, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetOrganizationApplianceDNSLocalProfilesAssignments Fetch the local DNS profile assignments in the organization
/* Fetch the local DNS profile assignments in the organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationApplianceDnsLocalProfilesAssignmentsQueryParams Filtering parameter


*/

func (s *ApplianceService) GetOrganizationApplianceDNSLocalProfilesAssignments(organizationID string, getOrganizationApplianceDnsLocalProfilesAssignmentsQueryParams *GetOrganizationApplianceDNSLocalProfilesAssignmentsQueryParams) (*ResponseApplianceGetOrganizationApplianceDNSLocalProfilesAssignments, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/appliance/dns/local/profiles/assignments"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetOrganizationApplianceDNSLocalProfilesAssignments](
		func() (*resty.Response, error) {
			return GET(path, s.client, getOrganizationApplianceDnsLocalProfilesAssignmentsQueryParams, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetOrganizationApplianceDNSLocalRecords Fetch the DNS records used in local DNS profiles
/* Fetch the DNS records used in local DNS profiles

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationApplianceDnsLocalRecordsQueryParams Filtering parameter


*/

func (s *ApplianceService) GetOrganizationApplianceDNSLocalRecords(organizationID string, getOrganizationApplianceDnsLocalRecordsQueryParams *GetOrganizationApplianceDNSLocalRecordsQueryParams) (*ResponseApplianceGetOrganizationApplianceDNSLocalRecords, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/appliance/dns/local/records"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetOrganizationApplianceDNSLocalRecords](
		func() (*resty.Response, error) {
			return GET(path, s.client, getOrganizationApplianceDnsLocalRecordsQueryParams, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetOrganizationApplianceDNSSplitProfiles Fetch the split DNS profiles used in the organization
/* Fetch the split DNS profiles used in the organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationApplianceDnsSplitProfilesQueryParams Filtering parameter


*/

func (s *ApplianceService) GetOrganizationApplianceDNSSplitProfiles(organizationID string, getOrganizationApplianceDnsSplitProfilesQueryParams *GetOrganizationApplianceDNSSplitProfilesQueryParams) (*ResponseApplianceGetOrganizationApplianceDNSSplitProfiles, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/appliance/dns/split/profiles"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetOrganizationApplianceDNSSplitProfiles](
		func() (*resty.Response, error) {
			return GET(path, s.client, getOrganizationApplianceDnsSplitProfilesQueryParams, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetOrganizationApplianceDNSSplitProfilesAssignments Fetch the split DNS profile assignments in the organization
/* Fetch the split DNS profile assignments in the organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationApplianceDnsSplitProfilesAssignmentsQueryParams Filtering parameter


*/

func (s *ApplianceService) GetOrganizationApplianceDNSSplitProfilesAssignments(organizationID string, getOrganizationApplianceDnsSplitProfilesAssignmentsQueryParams *GetOrganizationApplianceDNSSplitProfilesAssignmentsQueryParams) (*ResponseApplianceGetOrganizationApplianceDNSSplitProfilesAssignments, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/appliance/dns/split/profiles/assignments"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetOrganizationApplianceDNSSplitProfilesAssignments](
		func() (*resty.Response, error) {
			return GET(path, s.client, getOrganizationApplianceDnsSplitProfilesAssignmentsQueryParams, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetOrganizationApplianceFirewallMulticastForwardingByNetwork List Static Multicasting forwarding settings for MX networks
/* List Static Multicasting forwarding settings for MX networks

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationApplianceFirewallMulticastForwardingByNetworkQueryParams Filtering parameter


*/

func (s *ApplianceService) GetOrganizationApplianceFirewallMulticastForwardingByNetwork(organizationID string, getOrganizationApplianceFirewallMulticastForwardingByNetworkQueryParams *GetOrganizationApplianceFirewallMulticastForwardingByNetworkQueryParams) (*ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetwork, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/appliance/firewall/multicastForwarding/byNetwork"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetwork](
		func() (*resty.Response, error) {
			return GET(path, s.client, getOrganizationApplianceFirewallMulticastForwardingByNetworkQueryParams, &HeaderDefault)
		},
		s.client,
		func(dst, src ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetwork) ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetwork {
			*dst.Items = append(*dst.Items, *src.Items...)
			return dst
		},
		func() bool {
			if getOrganizationApplianceFirewallMulticastForwardingByNetworkQueryParams != nil {
				return getOrganizationApplianceFirewallMulticastForwardingByNetworkQueryParams.PerPage == -1
			}
			return false
		}(),
	)

}

//GetOrganizationApplianceSecurityEvents List the security events for an organization
/* List the security events for an organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationApplianceSecurityEventsQueryParams Filtering parameter


*/

func (s *ApplianceService) GetOrganizationApplianceSecurityEvents(organizationID string, getOrganizationApplianceSecurityEventsQueryParams *GetOrganizationApplianceSecurityEventsQueryParams) (*ResponseApplianceGetOrganizationApplianceSecurityEvents, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/appliance/security/events"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetOrganizationApplianceSecurityEvents](
		func() (*resty.Response, error) {
			return GET(path, s.client, getOrganizationApplianceSecurityEventsQueryParams, &HeaderDefault)
		},
		s.client,
		func(dst, src ResponseApplianceGetOrganizationApplianceSecurityEvents) ResponseApplianceGetOrganizationApplianceSecurityEvents {
			dst = append(dst, src...)
			return dst
		},
		func() bool {
			if getOrganizationApplianceSecurityEventsQueryParams != nil {
				return getOrganizationApplianceSecurityEventsQueryParams.PerPage == -1
			}
			return false
		}(),
	)

}

//GetOrganizationApplianceSecurityIntrusion Returns all supported intrusion settings for an organization
/* Returns all supported intrusion settings for an organization

@param organizationID organizationId path parameter. Organization ID


*/

func (s *ApplianceService) GetOrganizationApplianceSecurityIntrusion(organizationID string) (*ResponseApplianceGetOrganizationApplianceSecurityIntrusion, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/appliance/security/intrusion"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetOrganizationApplianceSecurityIntrusion](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork Display VPN exclusion rules for MX networks.
/* Display VPN exclusion rules for MX networks.

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationApplianceTrafficShapingVpnExclusionsByNetworkQueryParams Filtering parameter


*/

func (s *ApplianceService) GetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork(organizationID string, getOrganizationApplianceTrafficShapingVpnExclusionsByNetworkQueryParams *GetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkQueryParams) (*ResponseApplianceGetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/appliance/trafficShaping/vpnExclusions/byNetwork"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork](
		func() (*resty.Response, error) {
			return GET(path, s.client, getOrganizationApplianceTrafficShapingVpnExclusionsByNetworkQueryParams, &HeaderDefault)
		},
		s.client,
		func(dst, src ResponseApplianceGetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork) ResponseApplianceGetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork {
			*dst.Items = append(*dst.Items, *src.Items...)
			return dst
		},
		func() bool {
			if getOrganizationApplianceTrafficShapingVpnExclusionsByNetworkQueryParams != nil {
				return getOrganizationApplianceTrafficShapingVpnExclusionsByNetworkQueryParams.PerPage == -1
			}
			return false
		}(),
	)

}

//GetOrganizationApplianceUplinkStatuses List the uplink status of every Meraki MX and Z series appliances in the organization
/* List the uplink status of every Meraki MX and Z series appliances in the organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationApplianceUplinkStatusesQueryParams Filtering parameter


*/

func (s *ApplianceService) GetOrganizationApplianceUplinkStatuses(organizationID string, getOrganizationApplianceUplinkStatusesQueryParams *GetOrganizationApplianceUplinkStatusesQueryParams) (*ResponseApplianceGetOrganizationApplianceUplinkStatuses, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/appliance/uplink/statuses"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	// Other way

	return doWithRetriesAndResult(
		func() (*resty.Response, error) {
			return GET(path, s.client, getOrganizationApplianceUplinkStatusesQueryParams, &HeaderDefault)
		},
		s.client,
		func(dst, src ResponseApplianceGetOrganizationApplianceUplinkStatuses) ResponseApplianceGetOrganizationApplianceUplinkStatuses {
			dst = append(dst, src...)
			return dst
		},
		func() bool {
			if getOrganizationApplianceUplinkStatusesQueryParams != nil {
				return getOrganizationApplianceUplinkStatusesQueryParams.PerPage == -1
			}
			return false
		}(),
	)

}

//GetOrganizationApplianceUplinksStatusesOverview Returns an overview of uplink statuses
/* Returns an overview of uplink statuses

@param organizationID organizationId path parameter. Organization ID


*/

func (s *ApplianceService) GetOrganizationApplianceUplinksStatusesOverview(organizationID string) (*ResponseApplianceGetOrganizationApplianceUplinksStatusesOverview, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/appliance/uplinks/statuses/overview"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetOrganizationApplianceUplinksStatusesOverview](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetOrganizationApplianceUplinksUsageByNetwork Get the sent and received bytes for each uplink of all MX and Z networks within an organization
/* Get the sent and received bytes for each uplink of all MX and Z networks within an organization. If more than one device was active during the specified timespan, then the sent and received bytes will be aggregated by interface.

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationApplianceUplinksUsageByNetworkQueryParams Filtering parameter


*/

func (s *ApplianceService) GetOrganizationApplianceUplinksUsageByNetwork(organizationID string, getOrganizationApplianceUplinksUsageByNetworkQueryParams *GetOrganizationApplianceUplinksUsageByNetworkQueryParams) (*ResponseApplianceGetOrganizationApplianceUplinksUsageByNetwork, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/appliance/uplinks/usage/byNetwork"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetOrganizationApplianceUplinksUsageByNetwork](
		func() (*resty.Response, error) {
			return GET(path, s.client, getOrganizationApplianceUplinksUsageByNetworkQueryParams, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetOrganizationApplianceVpnStats Show VPN history stat for networks in an organization
/* Show VPN history stat for networks in an organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationApplianceVpnStatsQueryParams Filtering parameter


*/

func (s *ApplianceService) GetOrganizationApplianceVpnStats(organizationID string, getOrganizationApplianceVpnStatsQueryParams *GetOrganizationApplianceVpnStatsQueryParams) (*ResponseApplianceGetOrganizationApplianceVpnStats, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/appliance/vpn/stats"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetOrganizationApplianceVpnStats](
		func() (*resty.Response, error) {
			return GET(path, s.client, getOrganizationApplianceVpnStatsQueryParams, &HeaderDefault)
		},
		s.client,
		func(dst, src ResponseApplianceGetOrganizationApplianceVpnStats) ResponseApplianceGetOrganizationApplianceVpnStats {
			dst = append(dst, src...)
			return dst
		},
		func() bool {
			if getOrganizationApplianceVpnStatsQueryParams != nil {
				return getOrganizationApplianceVpnStatsQueryParams.PerPage == -1
			}
			return false
		}(),
	)

}

//GetOrganizationApplianceVpnStatuses Show VPN status for networks in an organization
/* Show VPN status for networks in an organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationApplianceVpnStatusesQueryParams Filtering parameter


*/

func (s *ApplianceService) GetOrganizationApplianceVpnStatuses(organizationID string, getOrganizationApplianceVpnStatusesQueryParams *GetOrganizationApplianceVpnStatusesQueryParams) (*ResponseApplianceGetOrganizationApplianceVpnStatuses, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/appliance/vpn/statuses"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetOrganizationApplianceVpnStatuses](
		func() (*resty.Response, error) {
			return GET(path, s.client, getOrganizationApplianceVpnStatusesQueryParams, &HeaderDefault)
		},
		s.client,
		func(dst, src ResponseApplianceGetOrganizationApplianceVpnStatuses) ResponseApplianceGetOrganizationApplianceVpnStatuses {
			*dst.Vpnstatusentities = append(*dst.Vpnstatusentities, *src.Vpnstatusentities...) // Total arrays: 1
			return dst
		},
		func() bool {
			if getOrganizationApplianceVpnStatusesQueryParams != nil {
				return getOrganizationApplianceVpnStatusesQueryParams.PerPage == -1
			}
			return false
		}(),
	)

}

//GetOrganizationApplianceVpnThirdPartyVpnpeers Return the third party VPN peers for an organization
/* Return the third party VPN peers for an organization

@param organizationID organizationId path parameter. Organization ID


*/

func (s *ApplianceService) GetOrganizationApplianceVpnThirdPartyVpnpeers(organizationID string) (*ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeers, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/appliance/vpn/thirdPartyVPNPeers"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeers](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//GetOrganizationApplianceVpnVpnFirewallRules Return the firewall rules for an organization's site-to-site VPN
/* Return the firewall rules for an organization's site-to-site VPN

@param organizationID organizationId path parameter. Organization ID


*/

func (s *ApplianceService) GetOrganizationApplianceVpnVpnFirewallRules(organizationID string) (*ResponseApplianceGetOrganizationApplianceVpnVpnFirewallRules, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/appliance/vpn/vpnFirewallRules"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceGetOrganizationApplianceVpnVpnFirewallRules](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		nil,
	)

}

//CreateDeviceApplianceVmxAuthenticationToken Generate a new vMX authentication token
/* Generate a new vMX authentication token

@param serial serial path parameter.


*/

func (s *ApplianceService) CreateDeviceApplianceVmxAuthenticationToken(serial string) (*ResponseApplianceCreateDeviceApplianceVmxAuthenticationToken, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/appliance/vmx/authenticationToken"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	// Past way

	return doWithRetriesAndResult[ResponseApplianceCreateDeviceApplianceVmxAuthenticationToken](
		func() (*resty.Response, error) {
			return POST(path, s.client, nil, nil)
		},
		s.client,
		nil,
	)

}

//CreateNetworkAppliancePrefixesDelegatedStatic Add a static delegated prefix from a network
/* Add a static delegated prefix from a network

@param networkID networkId path parameter. Network ID


*/

func (s *ApplianceService) CreateNetworkAppliancePrefixesDelegatedStatic(networkID string, requestApplianceCreateNetworkAppliancePrefixesDelegatedStatic *RequestApplianceCreateNetworkAppliancePrefixesDelegatedStatic) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/prefixes/delegated/statics"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Past way

	return doWithRetriesAndNotResult(
		func() (*resty.Response, error) {
			return POST(path, s.client, requestApplianceCreateNetworkAppliancePrefixesDelegatedStatic, nil)
		},
	)

}

//CreateNetworkApplianceRfProfile Creates new RF profile for this network
/* Creates new RF profile for this network

@param networkID networkId path parameter. Network ID


*/

func (s *ApplianceService) CreateNetworkApplianceRfProfile(networkID string, requestApplianceCreateNetworkApplianceRfProfile *RequestApplianceCreateNetworkApplianceRfProfile) (*ResponseApplianceCreateNetworkApplianceRfProfile, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/rfProfiles"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Past way

	return doWithRetriesAndResult[ResponseApplianceCreateNetworkApplianceRfProfile](
		func() (*resty.Response, error) {
			return POST(path, s.client, requestApplianceCreateNetworkApplianceRfProfile, nil)
		},
		s.client,
		nil,
	)

}

//CreateNetworkApplianceStaticRoute Add a static route for an MX or teleworker network
/* Add a static route for an MX or teleworker network

@param networkID networkId path parameter. Network ID


*/

func (s *ApplianceService) CreateNetworkApplianceStaticRoute(networkID string, requestApplianceCreateNetworkApplianceStaticRoute *RequestApplianceCreateNetworkApplianceStaticRoute) (*ResponseApplianceCreateNetworkApplianceStaticRoute, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/staticRoutes"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Past way

	return doWithRetriesAndResult[ResponseApplianceCreateNetworkApplianceStaticRoute](
		func() (*resty.Response, error) {
			return POST(path, s.client, requestApplianceCreateNetworkApplianceStaticRoute, nil)
		},
		s.client,
		nil,
	)

}

//CreateNetworkApplianceTrafficShapingCustomPerformanceClass Add a custom performance class for an MX network
/* Add a custom performance class for an MX network

@param networkID networkId path parameter. Network ID


*/

func (s *ApplianceService) CreateNetworkApplianceTrafficShapingCustomPerformanceClass(networkID string, requestApplianceCreateNetworkApplianceTrafficShapingCustomPerformanceClass *RequestApplianceCreateNetworkApplianceTrafficShapingCustomPerformanceClass) (*ResponseApplianceCreateNetworkApplianceTrafficShapingCustomPerformanceClass, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/trafficShaping/customPerformanceClasses"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Past way

	return doWithRetriesAndResult[ResponseApplianceCreateNetworkApplianceTrafficShapingCustomPerformanceClass](
		func() (*resty.Response, error) {
			return POST(path, s.client, requestApplianceCreateNetworkApplianceTrafficShapingCustomPerformanceClass, nil)
		},
		s.client,
		nil,
	)

}

//CreateNetworkApplianceVLAN Add a VLAN
/* Add a VLAN

@param networkID networkId path parameter. Network ID


*/

func (s *ApplianceService) CreateNetworkApplianceVLAN(networkID string, requestApplianceCreateNetworkApplianceVlan *RequestApplianceCreateNetworkApplianceVLAN) (*ResponseApplianceCreateNetworkApplianceVLAN, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/vlans"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Past way

	return doWithRetriesAndResult[ResponseApplianceCreateNetworkApplianceVLAN](
		func() (*resty.Response, error) {
			return POST(path, s.client, requestApplianceCreateNetworkApplianceVlan, nil)
		},
		s.client,
		nil,
	)

}

//SwapNetworkApplianceWarmSpare Swap MX primary and warm spare appliances
/* Swap MX primary and warm spare appliances

@param networkID networkId path parameter. Network ID


*/

func (s *ApplianceService) SwapNetworkApplianceWarmSpare(networkID string) (*ResponseApplianceSwapNetworkApplianceWarmSpare, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/warmSpare/swap"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Past way

	return doWithRetriesAndResult[ResponseApplianceSwapNetworkApplianceWarmSpare](
		func() (*resty.Response, error) {
			return POST(path, s.client, nil, nil)
		},
		s.client,
		nil,
	)

}

//CreateOrganizationApplianceDNSLocalProfile Create a new local DNS profile
/* Create a new local DNS profile

@param organizationID organizationId path parameter. Organization ID


*/

func (s *ApplianceService) CreateOrganizationApplianceDNSLocalProfile(organizationID string, requestApplianceCreateOrganizationApplianceDnsLocalProfile *RequestApplianceCreateOrganizationApplianceDNSLocalProfile) (*ResponseApplianceCreateOrganizationApplianceDNSLocalProfile, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/appliance/dns/local/profiles"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	// Past way

	return doWithRetriesAndResult[ResponseApplianceCreateOrganizationApplianceDNSLocalProfile](
		func() (*resty.Response, error) {
			return POST(path, s.client, requestApplianceCreateOrganizationApplianceDnsLocalProfile, nil)
		},
		s.client,
		nil,
	)

}

//BulkOrganizationApplianceDNSLocalProfilesAssignmentsCreate Assign the local DNS profile to networks in the organization
/* Assign the local DNS profile to networks in the organization

@param organizationID organizationId path parameter. Organization ID


*/

func (s *ApplianceService) BulkOrganizationApplianceDNSLocalProfilesAssignmentsCreate(organizationID string, requestApplianceBulkOrganizationApplianceDnsLocalProfilesAssignmentsCreate *RequestApplianceBulkOrganizationApplianceDNSLocalProfilesAssignmentsCreate) (*ResponseApplianceBulkOrganizationApplianceDNSLocalProfilesAssignmentsCreate, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/appliance/dns/local/profiles/assignments/bulkCreate"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	// Past way

	return doWithRetriesAndResult[ResponseApplianceBulkOrganizationApplianceDNSLocalProfilesAssignmentsCreate](
		func() (*resty.Response, error) {
			return POST(path, s.client, requestApplianceBulkOrganizationApplianceDnsLocalProfilesAssignmentsCreate, nil)
		},
		s.client,
		nil,
	)

}

//CreateOrganizationApplianceDNSLocalProfilesAssignmentsBulkDelete Unassign the local DNS profile to networks in the organization
/* Unassign the local DNS profile to networks in the organization

@param organizationID organizationId path parameter. Organization ID


*/

func (s *ApplianceService) CreateOrganizationApplianceDNSLocalProfilesAssignmentsBulkDelete(organizationID string, requestApplianceCreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete *RequestApplianceCreateOrganizationApplianceDNSLocalProfilesAssignmentsBulkDelete) (*ResponseApplianceCreateOrganizationApplianceDNSLocalProfilesAssignmentsBulkDelete, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/appliance/dns/local/profiles/assignments/bulkDelete"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	// Past way

	return doWithRetriesAndResult[ResponseApplianceCreateOrganizationApplianceDNSLocalProfilesAssignmentsBulkDelete](
		func() (*resty.Response, error) {
			return POST(path, s.client, requestApplianceCreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete, nil)
		},
		s.client,
		nil,
	)

}

//CreateOrganizationApplianceDNSLocalRecord Create a new local DNS record
/* Create a new local DNS record

@param organizationID organizationId path parameter. Organization ID


*/

func (s *ApplianceService) CreateOrganizationApplianceDNSLocalRecord(organizationID string, requestApplianceCreateOrganizationApplianceDnsLocalRecord *RequestApplianceCreateOrganizationApplianceDNSLocalRecord) (*ResponseApplianceCreateOrganizationApplianceDNSLocalRecord, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/appliance/dns/local/records"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	// Past way

	return doWithRetriesAndResult[ResponseApplianceCreateOrganizationApplianceDNSLocalRecord](
		func() (*resty.Response, error) {
			return POST(path, s.client, requestApplianceCreateOrganizationApplianceDnsLocalRecord, nil)
		},
		s.client,
		nil,
	)

}

//CreateOrganizationApplianceDNSSplitProfile Create a new split DNS profile
/* Create a new split DNS profile

@param organizationID organizationId path parameter. Organization ID


*/

func (s *ApplianceService) CreateOrganizationApplianceDNSSplitProfile(organizationID string, requestApplianceCreateOrganizationApplianceDnsSplitProfile *RequestApplianceCreateOrganizationApplianceDNSSplitProfile) (*ResponseApplianceCreateOrganizationApplianceDNSSplitProfile, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/appliance/dns/split/profiles"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	// Past way

	return doWithRetriesAndResult[ResponseApplianceCreateOrganizationApplianceDNSSplitProfile](
		func() (*resty.Response, error) {
			return POST(path, s.client, requestApplianceCreateOrganizationApplianceDnsSplitProfile, nil)
		},
		s.client,
		nil,
	)

}

//CreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreate Assign the split DNS profile to networks in the organization
/* Assign the split DNS profile to networks in the organization

@param organizationID organizationId path parameter. Organization ID


*/

func (s *ApplianceService) CreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreate(organizationID string, requestApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate *RequestApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreate) (*ResponseApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreate, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/appliance/dns/split/profiles/assignments/bulkCreate"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	// Past way

	return doWithRetriesAndResult[ResponseApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreate](
		func() (*resty.Response, error) {
			return POST(path, s.client, requestApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate, nil)
		},
		s.client,
		nil,
	)

}

//CreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkDelete Unassign the split DNS profile to networks in the organization
/* Unassign the split DNS profile to networks in the organization

@param organizationID organizationId path parameter. Organization ID


*/

func (s *ApplianceService) CreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkDelete(organizationID string, requestApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete *RequestApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkDelete) (*ResponseApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkDelete, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/appliance/dns/split/profiles/assignments/bulkDelete"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	// Past way

	return doWithRetriesAndResult[ResponseApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkDelete](
		func() (*resty.Response, error) {
			return POST(path, s.client, requestApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete, nil)
		},
		s.client,
		nil,
	)

}

//UpdateDeviceApplianceRadioSettings Update the radio settings of an appliance
/* Update the radio settings of an appliance

@param serial serial path parameter.
*/
func (s *ApplianceService) UpdateDeviceApplianceRadioSettings(serial string, requestApplianceUpdateDeviceApplianceRadioSettings *RequestApplianceUpdateDeviceApplianceRadioSettings) (*ResponseApplianceUpdateDeviceApplianceRadioSettings, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/appliance/radio/settings"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceUpdateDeviceApplianceRadioSettings](
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestApplianceUpdateDeviceApplianceRadioSettings)
		},
		s.client,
		nil,
	)

}

//UpdateDeviceApplianceUplinksSettings Update the uplink settings for an MX appliance
/* Update the uplink settings for an MX appliance

@param serial serial path parameter.
*/
func (s *ApplianceService) UpdateDeviceApplianceUplinksSettings(serial string, requestApplianceUpdateDeviceApplianceUplinksSettings *RequestApplianceUpdateDeviceApplianceUplinksSettings) (*ResponseApplianceUpdateDeviceApplianceUplinksSettings, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/appliance/uplinks/settings"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceUpdateDeviceApplianceUplinksSettings](
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestApplianceUpdateDeviceApplianceUplinksSettings)
		},
		s.client,
		nil,
	)

}

//UpdateNetworkApplianceConnectivityMonitoringDestinations Update the connectivity testing destinations for an MX network
/* Update the connectivity testing destinations for an MX network

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceConnectivityMonitoringDestinations(networkID string, requestApplianceUpdateNetworkApplianceConnectivityMonitoringDestinations *RequestApplianceUpdateNetworkApplianceConnectivityMonitoringDestinations) (*ResponseApplianceUpdateNetworkApplianceConnectivityMonitoringDestinations, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/connectivityMonitoringDestinations"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceUpdateNetworkApplianceConnectivityMonitoringDestinations](
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestApplianceUpdateNetworkApplianceConnectivityMonitoringDestinations)
		},
		s.client,
		nil,
	)

}

//UpdateNetworkApplianceContentFiltering Update the content filtering settings for an MX network
/* Update the content filtering settings for an MX network

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceContentFiltering(networkID string, requestApplianceUpdateNetworkApplianceContentFiltering *RequestApplianceUpdateNetworkApplianceContentFiltering) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/contentFiltering"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndNotResult(
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestApplianceUpdateNetworkApplianceContentFiltering)
		},
	)

}

//UpdateNetworkApplianceFirewallCellularFirewallRules Update the cellular firewall rules of an MX network
/* Update the cellular firewall rules of an MX network

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceFirewallCellularFirewallRules(networkID string, requestApplianceUpdateNetworkApplianceFirewallCellularFirewallRules *RequestApplianceUpdateNetworkApplianceFirewallCellularFirewallRules) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/cellularFirewallRules"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndNotResult(
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestApplianceUpdateNetworkApplianceFirewallCellularFirewallRules)
		},
	)

}

//UpdateNetworkApplianceFirewallFirewalledService Updates the accessibility settings for the given service ('ICMP', 'web', or 'SNMP')
/* Updates the accessibility settings for the given service ('ICMP', 'web', or 'SNMP')

@param networkID networkId path parameter. Network ID
@param service service path parameter.
*/
func (s *ApplianceService) UpdateNetworkApplianceFirewallFirewalledService(networkID string, service string, requestApplianceUpdateNetworkApplianceFirewallFirewalledService *RequestApplianceUpdateNetworkApplianceFirewallFirewalledService) (*ResponseApplianceUpdateNetworkApplianceFirewallFirewalledService, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/firewalledServices/{service}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{service}", fmt.Sprintf("%v", service), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceUpdateNetworkApplianceFirewallFirewalledService](
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestApplianceUpdateNetworkApplianceFirewallFirewalledService)
		},
		s.client,
		nil,
	)

}

//UpdateNetworkApplianceFirewallInboundCellularFirewallRules Update the inbound cellular firewall rules of an MX network
/* Update the inbound cellular firewall rules of an MX network

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceFirewallInboundCellularFirewallRules(networkID string, requestApplianceUpdateNetworkApplianceFirewallInboundCellularFirewallRules *RequestApplianceUpdateNetworkApplianceFirewallInboundCellularFirewallRules) (*ResponseApplianceUpdateNetworkApplianceFirewallInboundCellularFirewallRules, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/inboundCellularFirewallRules"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceUpdateNetworkApplianceFirewallInboundCellularFirewallRules](
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestApplianceUpdateNetworkApplianceFirewallInboundCellularFirewallRules)
		},
		s.client,
		nil,
	)

}

//UpdateNetworkApplianceFirewallInboundFirewallRules Update the inbound firewall rules of an MX network
/* Update the inbound firewall rules of an MX network

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceFirewallInboundFirewallRules(networkID string, requestApplianceUpdateNetworkApplianceFirewallInboundFirewallRules *RequestApplianceUpdateNetworkApplianceFirewallInboundFirewallRules) (*ResponseApplianceUpdateNetworkApplianceFirewallInboundFirewallRules, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/inboundFirewallRules"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceUpdateNetworkApplianceFirewallInboundFirewallRules](
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestApplianceUpdateNetworkApplianceFirewallInboundFirewallRules)
		},
		s.client,
		nil,
	)

}

//UpdateNetworkApplianceFirewallL3FirewallRules Update the L3 firewall rules of an MX network
/* Update the L3 firewall rules of an MX network

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceFirewallL3FirewallRules(networkID string, requestApplianceUpdateNetworkApplianceFirewallL3FirewallRules *RequestApplianceUpdateNetworkApplianceFirewallL3FirewallRules) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/l3FirewallRules"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndNotResult(
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestApplianceUpdateNetworkApplianceFirewallL3FirewallRules)
		},
	)

}

//UpdateNetworkApplianceFirewallL7FirewallRules Update the MX L7 firewall rules for an MX network
/* Update the MX L7 firewall rules for an MX network

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceFirewallL7FirewallRules(networkID string, requestApplianceUpdateNetworkApplianceFirewallL7FirewallRules *RequestApplianceUpdateNetworkApplianceFirewallL7FirewallRules) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/l7FirewallRules"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndNotResult(
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestApplianceUpdateNetworkApplianceFirewallL7FirewallRules)
		},
	)

}

//UpdateNetworkApplianceFirewallMulticastForwarding Update static multicast forward rules for a network
/* Update static multicast forward rules for a network

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceFirewallMulticastForwarding(networkID string, requestApplianceUpdateNetworkApplianceFirewallMulticastForwarding *RequestApplianceUpdateNetworkApplianceFirewallMulticastForwarding) (*ResponseApplianceUpdateNetworkApplianceFirewallMulticastForwarding, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/multicastForwarding"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceUpdateNetworkApplianceFirewallMulticastForwarding](
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestApplianceUpdateNetworkApplianceFirewallMulticastForwarding)
		},
		s.client,
		nil,
	)

}

//UpdateNetworkApplianceFirewallOneToManyNatRules Set the 1:Many NAT mapping rules for an MX network
/* Set the 1:Many NAT mapping rules for an MX network

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceFirewallOneToManyNatRules(networkID string, requestApplianceUpdateNetworkApplianceFirewallOneToManyNatRules *RequestApplianceUpdateNetworkApplianceFirewallOneToManyNatRules) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/oneToManyNatRules"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndNotResult(
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestApplianceUpdateNetworkApplianceFirewallOneToManyNatRules)
		},
	)

}

//UpdateNetworkApplianceFirewallOneToOneNatRules Set the 1:1 NAT mapping rules for an MX network
/* Set the 1:1 NAT mapping rules for an MX network

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceFirewallOneToOneNatRules(networkID string, requestApplianceUpdateNetworkApplianceFirewallOneToOneNatRules *RequestApplianceUpdateNetworkApplianceFirewallOneToOneNatRules) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/oneToOneNatRules"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndNotResult(
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestApplianceUpdateNetworkApplianceFirewallOneToOneNatRules)
		},
	)

}

//UpdateNetworkApplianceFirewallPortForwardingRules Update the port forwarding rules for an MX network
/* Update the port forwarding rules for an MX network

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceFirewallPortForwardingRules(networkID string, requestApplianceUpdateNetworkApplianceFirewallPortForwardingRules *RequestApplianceUpdateNetworkApplianceFirewallPortForwardingRules) (*ResponseApplianceUpdateNetworkApplianceFirewallPortForwardingRules, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/portForwardingRules"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceUpdateNetworkApplianceFirewallPortForwardingRules](
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestApplianceUpdateNetworkApplianceFirewallPortForwardingRules)
		},
		s.client,
		nil,
	)

}

//UpdateNetworkApplianceFirewallSettings Update the firewall settings for this network
/* Update the firewall settings for this network

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceFirewallSettings(networkID string, requestApplianceUpdateNetworkApplianceFirewallSettings *RequestApplianceUpdateNetworkApplianceFirewallSettings) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/settings"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndNotResult(
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestApplianceUpdateNetworkApplianceFirewallSettings)
		},
	)

}

//UpdateNetworkAppliancePort Update the per-port VLAN settings for a single MX port.
/* Update the per-port VLAN settings for a single MX port.

@param networkID networkId path parameter. Network ID
@param portID portId path parameter. Port ID
*/
func (s *ApplianceService) UpdateNetworkAppliancePort(networkID string, portID string, requestApplianceUpdateNetworkAppliancePort *RequestApplianceUpdateNetworkAppliancePort) (*ResponseApplianceUpdateNetworkAppliancePort, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/ports/{portId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{portId}", fmt.Sprintf("%v", portID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceUpdateNetworkAppliancePort](
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestApplianceUpdateNetworkAppliancePort)
		},
		s.client,
		nil,
	)

}

//UpdateNetworkAppliancePrefixesDelegatedStatic Update a static delegated prefix from a network
/* Update a static delegated prefix from a network

@param networkID networkId path parameter. Network ID
@param staticDelegatedPrefixID staticDelegatedPrefixId path parameter. Static delegated prefix ID
*/
func (s *ApplianceService) UpdateNetworkAppliancePrefixesDelegatedStatic(networkID string, staticDelegatedPrefixID string, requestApplianceUpdateNetworkAppliancePrefixesDelegatedStatic *RequestApplianceUpdateNetworkAppliancePrefixesDelegatedStatic) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/prefixes/delegated/statics/{staticDelegatedPrefixId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{staticDelegatedPrefixId}", fmt.Sprintf("%v", staticDelegatedPrefixID), -1)

	// Other way

	return doWithRetriesAndNotResult(
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestApplianceUpdateNetworkAppliancePrefixesDelegatedStatic)
		},
	)

}

//UpdateNetworkApplianceRfProfile Updates specified RF profile for this network
/* Updates specified RF profile for this network

@param networkID networkId path parameter. Network ID
@param rfProfileID rfProfileId path parameter. Rf profile ID
*/
func (s *ApplianceService) UpdateNetworkApplianceRfProfile(networkID string, rfProfileID string, requestApplianceUpdateNetworkApplianceRfProfile *RequestApplianceUpdateNetworkApplianceRfProfile) (*ResponseApplianceUpdateNetworkApplianceRfProfile, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/rfProfiles/{rfProfileId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{rfProfileId}", fmt.Sprintf("%v", rfProfileID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceUpdateNetworkApplianceRfProfile](
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestApplianceUpdateNetworkApplianceRfProfile)
		},
		s.client,
		nil,
	)

}

//UpdateNetworkApplianceSdwanInternetPolicies Update SDWAN internet traffic preferences for an MX network
/* Update SDWAN internet traffic preferences for an MX network

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceSdwanInternetPolicies(networkID string, requestApplianceUpdateNetworkApplianceSdwanInternetPolicies *RequestApplianceUpdateNetworkApplianceSdwanInternetPolicies) (*ResponseApplianceUpdateNetworkApplianceSdwanInternetPolicies, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/sdwan/internetPolicies"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceUpdateNetworkApplianceSdwanInternetPolicies](
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestApplianceUpdateNetworkApplianceSdwanInternetPolicies)
		},
		s.client,
		nil,
	)

}

//UpdateNetworkApplianceSecurityIntrusion Set the supported intrusion settings for an MX network
/* Set the supported intrusion settings for an MX network

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceSecurityIntrusion(networkID string, requestApplianceUpdateNetworkApplianceSecurityIntrusion *RequestApplianceUpdateNetworkApplianceSecurityIntrusion) (*ResponseApplianceUpdateNetworkApplianceSecurityIntrusion, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/security/intrusion"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceUpdateNetworkApplianceSecurityIntrusion](
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestApplianceUpdateNetworkApplianceSecurityIntrusion)
		},
		s.client,
		nil,
	)

}

//UpdateNetworkApplianceSecurityMalware Set the supported malware settings for an MX network
/* Set the supported malware settings for an MX network

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceSecurityMalware(networkID string, requestApplianceUpdateNetworkApplianceSecurityMalware *RequestApplianceUpdateNetworkApplianceSecurityMalware) (*ResponseApplianceUpdateNetworkApplianceSecurityMalware, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/security/malware"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceUpdateNetworkApplianceSecurityMalware](
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestApplianceUpdateNetworkApplianceSecurityMalware)
		},
		s.client,
		nil,
	)

}

//UpdateNetworkApplianceSettings Update the appliance settings for a network
/* Update the appliance settings for a network

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceSettings(networkID string, requestApplianceUpdateNetworkApplianceSettings *RequestApplianceUpdateNetworkApplianceSettings) (*ResponseApplianceUpdateNetworkApplianceSettings, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/settings"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceUpdateNetworkApplianceSettings](
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestApplianceUpdateNetworkApplianceSettings)
		},
		s.client,
		nil,
	)

}

//UpdateNetworkApplianceSingleLan Update single LAN configuration
/* Update single LAN configuration

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceSingleLan(networkID string, requestApplianceUpdateNetworkApplianceSingleLan *RequestApplianceUpdateNetworkApplianceSingleLan) (*ResponseApplianceUpdateNetworkApplianceSingleLan, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/singleLan"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceUpdateNetworkApplianceSingleLan](
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestApplianceUpdateNetworkApplianceSingleLan)
		},
		s.client,
		nil,
	)

}

//UpdateNetworkApplianceSSID Update the attributes of an MX SSID
/* Update the attributes of an MX SSID

@param networkID networkId path parameter. Network ID
@param number number path parameter.
*/
func (s *ApplianceService) UpdateNetworkApplianceSSID(networkID string, number string, requestApplianceUpdateNetworkApplianceSsid *RequestApplianceUpdateNetworkApplianceSSID) (*ResponseApplianceUpdateNetworkApplianceSSID, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/ssids/{number}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{number}", fmt.Sprintf("%v", number), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceUpdateNetworkApplianceSSID](
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestApplianceUpdateNetworkApplianceSsid)
		},
		s.client,
		nil,
	)

}

//UpdateNetworkApplianceStaticRoute Update a static route for an MX or teleworker network
/* Update a static route for an MX or teleworker network

@param networkID networkId path parameter. Network ID
@param staticRouteID staticRouteId path parameter. Static route ID
*/
func (s *ApplianceService) UpdateNetworkApplianceStaticRoute(networkID string, staticRouteID string, requestApplianceUpdateNetworkApplianceStaticRoute *RequestApplianceUpdateNetworkApplianceStaticRoute) (*ResponseApplianceUpdateNetworkApplianceStaticRoute, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/staticRoutes/{staticRouteId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{staticRouteId}", fmt.Sprintf("%v", staticRouteID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceUpdateNetworkApplianceStaticRoute](
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestApplianceUpdateNetworkApplianceStaticRoute)
		},
		s.client,
		nil,
	)

}

//UpdateNetworkApplianceTrafficShaping Update the traffic shaping settings for an MX network
/* Update the traffic shaping settings for an MX network

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceTrafficShaping(networkID string, requestApplianceUpdateNetworkApplianceTrafficShaping *RequestApplianceUpdateNetworkApplianceTrafficShaping) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/trafficShaping"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndNotResult(
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestApplianceUpdateNetworkApplianceTrafficShaping)
		},
	)

}

//UpdateNetworkApplianceTrafficShapingCustomPerformanceClass Update a custom performance class for an MX network
/* Update a custom performance class for an MX network

@param networkID networkId path parameter. Network ID
@param customPerformanceClassID customPerformanceClassId path parameter. Custom performance class ID
*/
func (s *ApplianceService) UpdateNetworkApplianceTrafficShapingCustomPerformanceClass(networkID string, customPerformanceClassID string, requestApplianceUpdateNetworkApplianceTrafficShapingCustomPerformanceClass *RequestApplianceUpdateNetworkApplianceTrafficShapingCustomPerformanceClass) (*ResponseApplianceUpdateNetworkApplianceTrafficShapingCustomPerformanceClass, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/trafficShaping/customPerformanceClasses/{customPerformanceClassId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{customPerformanceClassId}", fmt.Sprintf("%v", customPerformanceClassID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceUpdateNetworkApplianceTrafficShapingCustomPerformanceClass](
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestApplianceUpdateNetworkApplianceTrafficShapingCustomPerformanceClass)
		},
		s.client,
		nil,
	)

}

//UpdateNetworkApplianceTrafficShapingRules Update the traffic shaping settings rules for an MX network
/* Update the traffic shaping settings rules for an MX network

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceTrafficShapingRules(networkID string, requestApplianceUpdateNetworkApplianceTrafficShapingRules *RequestApplianceUpdateNetworkApplianceTrafficShapingRules) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/trafficShaping/rules"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndNotResult(
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestApplianceUpdateNetworkApplianceTrafficShapingRules)
		},
	)

}

//UpdateNetworkApplianceTrafficShapingUplinkBandwidth Updates the uplink bandwidth settings for your MX network.
/* Updates the uplink bandwidth settings for your MX network.

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceTrafficShapingUplinkBandwidth(networkID string, requestApplianceUpdateNetworkApplianceTrafficShapingUplinkBandwidth *RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkBandwidth) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/trafficShaping/uplinkBandwidth"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndNotResult(
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestApplianceUpdateNetworkApplianceTrafficShapingUplinkBandwidth)
		},
	)

}

//UpdateNetworkApplianceTrafficShapingUplinkSelection Update uplink selection settings for an MX network
/* Update uplink selection settings for an MX network

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceTrafficShapingUplinkSelection(networkID string, requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelection *RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelection) (*ResponseApplianceUpdateNetworkApplianceTrafficShapingUplinkSelection, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/trafficShaping/uplinkSelection"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceUpdateNetworkApplianceTrafficShapingUplinkSelection](
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelection)
		},
		s.client,
		nil,
	)

}

//UpdateNetworkApplianceTrafficShapingVpnExclusions Update VPN exclusion rules for an MX network.
/* Update VPN exclusion rules for an MX network.

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceTrafficShapingVpnExclusions(networkID string, requestApplianceUpdateNetworkApplianceTrafficShapingVpnExclusions *RequestApplianceUpdateNetworkApplianceTrafficShapingVpnExclusions) (*ResponseApplianceUpdateNetworkApplianceTrafficShapingVpnExclusions, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/trafficShaping/vpnExclusions"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceUpdateNetworkApplianceTrafficShapingVpnExclusions](
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestApplianceUpdateNetworkApplianceTrafficShapingVpnExclusions)
		},
		s.client,
		nil,
	)

}

//UpdateNetworkApplianceVLANsSettings Enable/Disable VLANs for the given network
/* Enable/Disable VLANs for the given network

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceVLANsSettings(networkID string, requestApplianceUpdateNetworkApplianceVlansSettings *RequestApplianceUpdateNetworkApplianceVLANsSettings) (*ResponseApplianceUpdateNetworkApplianceVLANsSettings, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/vlans/settings"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceUpdateNetworkApplianceVLANsSettings](
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestApplianceUpdateNetworkApplianceVlansSettings)
		},
		s.client,
		nil,
	)

}

//UpdateNetworkApplianceVLAN Update a VLAN
/* Update a VLAN

@param networkID networkId path parameter. Network ID
@param vlanID vlanId path parameter. Vlan ID
*/
func (s *ApplianceService) UpdateNetworkApplianceVLAN(networkID string, vlanID string, requestApplianceUpdateNetworkApplianceVlan *RequestApplianceUpdateNetworkApplianceVLAN) (*ResponseApplianceUpdateNetworkApplianceVLAN, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/vlans/{vlanId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{vlanId}", fmt.Sprintf("%v", vlanID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceUpdateNetworkApplianceVLAN](
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestApplianceUpdateNetworkApplianceVlan)
		},
		s.client,
		nil,
	)

}

//UpdateNetworkApplianceVpnBgp Update a Hub BGP Configuration
/* Update a Hub BGP Configuration

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceVpnBgp(networkID string, requestApplianceUpdateNetworkApplianceVpnBgp *RequestApplianceUpdateNetworkApplianceVpnBgp) (*ResponseApplianceUpdateNetworkApplianceVpnBgp, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/vpn/bgp"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceUpdateNetworkApplianceVpnBgp](
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestApplianceUpdateNetworkApplianceVpnBgp)
		},
		s.client,
		nil,
	)

}

//UpdateNetworkApplianceVpnSiteToSiteVpn Update the site-to-site VPN settings of a network
/* Update the site-to-site VPN settings of a network. Only valid for MX networks in NAT mode.

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceVpnSiteToSiteVpn(networkID string, requestApplianceUpdateNetworkApplianceVpnSiteToSiteVpn *RequestApplianceUpdateNetworkApplianceVpnSiteToSiteVpn) (*ResponseApplianceUpdateNetworkApplianceVpnSiteToSiteVpn, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/vpn/siteToSiteVpn"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceUpdateNetworkApplianceVpnSiteToSiteVpn](
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestApplianceUpdateNetworkApplianceVpnSiteToSiteVpn)
		},
		s.client,
		nil,
	)

}

//UpdateNetworkApplianceWarmSpare Update MX warm spare settings
/* Update MX warm spare settings

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceWarmSpare(networkID string, requestApplianceUpdateNetworkApplianceWarmSpare *RequestApplianceUpdateNetworkApplianceWarmSpare) (*ResponseApplianceUpdateNetworkApplianceWarmSpare, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/warmSpare"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceUpdateNetworkApplianceWarmSpare](
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestApplianceUpdateNetworkApplianceWarmSpare)
		},
		s.client,
		nil,
	)

}

//UpdateOrganizationApplianceDNSLocalProfile Update a local DNS profile
/* Update a local DNS profile

@param organizationID organizationId path parameter. Organization ID
@param profileID profileId path parameter. Profile ID
*/
func (s *ApplianceService) UpdateOrganizationApplianceDNSLocalProfile(organizationID string, profileID string, requestApplianceUpdateOrganizationApplianceDnsLocalProfile *RequestApplianceUpdateOrganizationApplianceDNSLocalProfile) (*ResponseApplianceUpdateOrganizationApplianceDNSLocalProfile, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/appliance/dns/local/profiles/{profileId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{profileId}", fmt.Sprintf("%v", profileID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceUpdateOrganizationApplianceDNSLocalProfile](
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestApplianceUpdateOrganizationApplianceDnsLocalProfile)
		},
		s.client,
		nil,
	)

}

//UpdateOrganizationApplianceDNSLocalRecord Updates a local DNS record
/* Updates a local DNS record

@param organizationID organizationId path parameter. Organization ID
@param recordID recordId path parameter. Record ID
*/
func (s *ApplianceService) UpdateOrganizationApplianceDNSLocalRecord(organizationID string, recordID string, requestApplianceUpdateOrganizationApplianceDnsLocalRecord *RequestApplianceUpdateOrganizationApplianceDNSLocalRecord) (*ResponseApplianceUpdateOrganizationApplianceDNSLocalRecord, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/appliance/dns/local/records/{recordId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{recordId}", fmt.Sprintf("%v", recordID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceUpdateOrganizationApplianceDNSLocalRecord](
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestApplianceUpdateOrganizationApplianceDnsLocalRecord)
		},
		s.client,
		nil,
	)

}

//UpdateOrganizationApplianceDNSSplitProfile Update a split DNS profile
/* Update a split DNS profile

@param organizationID organizationId path parameter. Organization ID
@param profileID profileId path parameter. Profile ID
*/
func (s *ApplianceService) UpdateOrganizationApplianceDNSSplitProfile(organizationID string, profileID string, requestApplianceUpdateOrganizationApplianceDnsSplitProfile *RequestApplianceUpdateOrganizationApplianceDNSSplitProfile) (*ResponseApplianceUpdateOrganizationApplianceDNSSplitProfile, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/appliance/dns/split/profiles/{profileId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{profileId}", fmt.Sprintf("%v", profileID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceUpdateOrganizationApplianceDNSSplitProfile](
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestApplianceUpdateOrganizationApplianceDnsSplitProfile)
		},
		s.client,
		nil,
	)

}

//UpdateOrganizationApplianceSecurityIntrusion Sets supported intrusion settings for an organization
/* Sets supported intrusion settings for an organization

@param organizationID organizationId path parameter. Organization ID
*/
func (s *ApplianceService) UpdateOrganizationApplianceSecurityIntrusion(organizationID string, requestApplianceUpdateOrganizationApplianceSecurityIntrusion *RequestApplianceUpdateOrganizationApplianceSecurityIntrusion) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/appliance/security/intrusion"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	// Other way

	return doWithRetriesAndNotResult(
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestApplianceUpdateOrganizationApplianceSecurityIntrusion)
		},
	)

}

//UpdateOrganizationApplianceVpnThirdPartyVpnpeers Update the third party VPN peers for an organization
/* Update the third party VPN peers for an organization

@param organizationID organizationId path parameter. Organization ID
*/
func (s *ApplianceService) UpdateOrganizationApplianceVpnThirdPartyVpnpeers(organizationID string, requestApplianceUpdateOrganizationApplianceVpnThirdPartyVpnpeers *RequestApplianceUpdateOrganizationApplianceVpnThirdPartyVpnpeers) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/appliance/vpn/thirdPartyVPNPeers"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	// Other way

	return doWithRetriesAndNotResult(
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestApplianceUpdateOrganizationApplianceVpnThirdPartyVpnpeers)
		},
	)

}

//UpdateOrganizationApplianceVpnVpnFirewallRules Update the firewall rules of an organization's site-to-site VPN
/* Update the firewall rules of an organization's site-to-site VPN

@param organizationID organizationId path parameter. Organization ID
*/
func (s *ApplianceService) UpdateOrganizationApplianceVpnVpnFirewallRules(organizationID string, requestApplianceUpdateOrganizationApplianceVpnVpnFirewallRules *RequestApplianceUpdateOrganizationApplianceVpnVpnFirewallRules) (*ResponseApplianceUpdateOrganizationApplianceVpnVpnFirewallRules, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/appliance/vpn/vpnFirewallRules"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseApplianceUpdateOrganizationApplianceVpnVpnFirewallRules](
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestApplianceUpdateOrganizationApplianceVpnVpnFirewallRules)
		},
		s.client,
		nil,
	)

}

//DeleteNetworkAppliancePrefixesDelegatedStatic Delete a static delegated prefix from a network
/* Delete a static delegated prefix from a network

@param networkID networkId path parameter. Network ID
@param staticDelegatedPrefixID staticDelegatedPrefixId path parameter. Static delegated prefix ID


*/
func (s *ApplianceService) DeleteNetworkAppliancePrefixesDelegatedStatic(networkID string, staticDelegatedPrefixID string) (*resty.Response, error) {
	//networkID string,staticDelegatedPrefixID string
	path := "/api/v1/networks/{networkId}/appliance/prefixes/delegated/statics/{staticDelegatedPrefixId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{staticDelegatedPrefixId}", fmt.Sprintf("%v", staticDelegatedPrefixID), -1)

	return doWithRetriesAndNotResult(
		func() (*resty.Response, error) {
			return DELETE(path, s.client, &QueryParamsDefault)
		},
	)
}

//DeleteNetworkApplianceRfProfile Delete a RF Profile
/* Delete a RF Profile

@param networkID networkId path parameter. Network ID
@param rfProfileID rfProfileId path parameter. Rf profile ID


*/
func (s *ApplianceService) DeleteNetworkApplianceRfProfile(networkID string, rfProfileID string) (*resty.Response, error) {
	//networkID string,rfProfileID string
	path := "/api/v1/networks/{networkId}/appliance/rfProfiles/{rfProfileId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{rfProfileId}", fmt.Sprintf("%v", rfProfileID), -1)

	return doWithRetriesAndNotResult(
		func() (*resty.Response, error) {
			return DELETE(path, s.client, &QueryParamsDefault)
		},
	)
}

//DeleteNetworkApplianceStaticRoute Delete a static route from an MX or teleworker network
/* Delete a static route from an MX or teleworker network

@param networkID networkId path parameter. Network ID
@param staticRouteID staticRouteId path parameter. Static route ID


*/
func (s *ApplianceService) DeleteNetworkApplianceStaticRoute(networkID string, staticRouteID string) (*resty.Response, error) {
	//networkID string,staticRouteID string
	path := "/api/v1/networks/{networkId}/appliance/staticRoutes/{staticRouteId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{staticRouteId}", fmt.Sprintf("%v", staticRouteID), -1)

	return doWithRetriesAndNotResult(
		func() (*resty.Response, error) {
			return DELETE(path, s.client, &QueryParamsDefault)
		},
	)
}

//DeleteNetworkApplianceTrafficShapingCustomPerformanceClass Delete a custom performance class from an MX network
/* Delete a custom performance class from an MX network

@param networkID networkId path parameter. Network ID
@param customPerformanceClassID customPerformanceClassId path parameter. Custom performance class ID


*/
func (s *ApplianceService) DeleteNetworkApplianceTrafficShapingCustomPerformanceClass(networkID string, customPerformanceClassID string) (*resty.Response, error) {
	//networkID string,customPerformanceClassID string
	path := "/api/v1/networks/{networkId}/appliance/trafficShaping/customPerformanceClasses/{customPerformanceClassId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{customPerformanceClassId}", fmt.Sprintf("%v", customPerformanceClassID), -1)

	return doWithRetriesAndNotResult(
		func() (*resty.Response, error) {
			return DELETE(path, s.client, &QueryParamsDefault)
		},
	)
}

//DeleteNetworkApplianceVLAN Delete a VLAN from a network
/* Delete a VLAN from a network

@param networkID networkId path parameter. Network ID
@param vlanID vlanId path parameter. Vlan ID


*/
func (s *ApplianceService) DeleteNetworkApplianceVLAN(networkID string, vlanID string) (*resty.Response, error) {
	//networkID string,vlanID string
	path := "/api/v1/networks/{networkId}/appliance/vlans/{vlanId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{vlanId}", fmt.Sprintf("%v", vlanID), -1)

	return doWithRetriesAndNotResult(
		func() (*resty.Response, error) {
			return DELETE(path, s.client, &QueryParamsDefault)
		},
	)
}

//DeleteOrganizationApplianceDNSLocalProfile Deletes a local DNS profile
/* Deletes a local DNS profile

@param organizationID organizationId path parameter. Organization ID
@param profileID profileId path parameter. Profile ID


*/
func (s *ApplianceService) DeleteOrganizationApplianceDNSLocalProfile(organizationID string, profileID string) (*resty.Response, error) {
	//organizationID string,profileID string
	path := "/api/v1/organizations/{organizationId}/appliance/dns/local/profiles/{profileId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{profileId}", fmt.Sprintf("%v", profileID), -1)

	return doWithRetriesAndNotResult(
		func() (*resty.Response, error) {
			return DELETE(path, s.client, &QueryParamsDefault)
		},
	)
}

//DeleteOrganizationApplianceDNSLocalRecord Deletes a local DNS record
/* Deletes a local DNS record

@param organizationID organizationId path parameter. Organization ID
@param recordID recordId path parameter. Record ID


*/
func (s *ApplianceService) DeleteOrganizationApplianceDNSLocalRecord(organizationID string, recordID string) (*resty.Response, error) {
	//organizationID string,recordID string
	path := "/api/v1/organizations/{organizationId}/appliance/dns/local/records/{recordId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{recordId}", fmt.Sprintf("%v", recordID), -1)

	return doWithRetriesAndNotResult(
		func() (*resty.Response, error) {
			return DELETE(path, s.client, &QueryParamsDefault)
		},
	)
}

//DeleteOrganizationApplianceDNSSplitProfile Deletes a split DNS profile
/* Deletes a split DNS profile

@param organizationID organizationId path parameter. Organization ID
@param profileID profileId path parameter. Profile ID


*/
func (s *ApplianceService) DeleteOrganizationApplianceDNSSplitProfile(organizationID string, profileID string) (*resty.Response, error) {
	//organizationID string,profileID string
	path := "/api/v1/organizations/{organizationId}/appliance/dns/split/profiles/{profileId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{profileId}", fmt.Sprintf("%v", profileID), -1)

	return doWithRetriesAndNotResult(
		func() (*resty.Response, error) {
			return DELETE(path, s.client, &QueryParamsDefault)
		},
	)
}
