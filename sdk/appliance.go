package meraki

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type ApplianceService service

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
type GetOrganizationApplianceSecurityEventsQueryParams struct {
	T0            string  `url:"t0,omitempty"`            //The beginning of the timespan for the data. Data is gathered after the specified t0 value. The maximum lookback period is 365 days from today.
	T1            string  `url:"t1,omitempty"`            //The end of the timespan for the data. t1 can be a maximum of 365 days after t0.
	Timespan      float64 `url:"timespan,omitempty"`      //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 365 days. The default is 31 days.
	PerPage       int     `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100.
	StartingAfter string  `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string  `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	SortOrder     string  `url:"sortOrder,omitempty"`     //Sorted order of security events based on event detection time. Order options are 'ascending' or 'descending'. Default is ascending order.
}
type GetOrganizationApplianceUplinkStatusesQueryParams struct {
	PerPage       int      `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter string   `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string   `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	NetworkIDs    []string `url:"networkIds[],omitempty"`  //A list of network IDs. The returned devices will be filtered to only include these networks.
	Serials       []string `url:"serials[],omitempty"`     //A list of serial numbers. The returned devices will be filtered to only include these serials.
	Iccids        []string `url:"iccids[],omitempty"`      //A list of ICCIDs. The returned devices will be filtered to only include these ICCIDs.
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
	FreeCount *int   `json:"freeCount,omitempty"` //
	Subnet    string `json:"subnet,omitempty"`    //
	UsedCount *int   `json:"usedCount,omitempty"` //
	VLANID    *int   `json:"vlanId,omitempty"`    //
}
type ResponseApplianceGetDeviceAppliancePerformance struct {
	PerfScore *int `json:"perfScore,omitempty"` //
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
	Destinations *[]ResponseApplianceGetNetworkApplianceConnectivityMonitoringDestinationsDestinations `json:"destinations,omitempty"` //
}
type ResponseApplianceGetNetworkApplianceConnectivityMonitoringDestinationsDestinations struct {
	Default     *bool  `json:"default,omitempty"`     //
	Description string `json:"description,omitempty"` //
	IP          string `json:"ip,omitempty"`          //
}
type ResponseApplianceUpdateNetworkApplianceConnectivityMonitoringDestinations interface{}
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
	Access     string   `json:"access,omitempty"`     //
	AllowedIPs []string `json:"allowedIps,omitempty"` //
	Service    string   `json:"service,omitempty"`    //
}
type ResponseApplianceGetNetworkApplianceFirewallFirewalledService struct {
	Access     string   `json:"access,omitempty"`     //
	AllowedIPs []string `json:"allowedIps,omitempty"` //
	Service    string   `json:"service,omitempty"`    //
}
type ResponseApplianceUpdateNetworkApplianceFirewallFirewalledService interface{}
type ResponseApplianceGetNetworkApplianceFirewallInboundCellularFirewallRules []ResponseItemApplianceGetNetworkApplianceFirewallInboundCellularFirewallRules // Array of ResponseApplianceGetNetworkApplianceFirewallInboundCellularFirewallRules
type ResponseItemApplianceGetNetworkApplianceFirewallInboundCellularFirewallRules struct {
	Comment       string `json:"comment,omitempty"`       //
	DestCidr      string `json:"destCidr,omitempty"`      //
	DestPort      string `json:"destPort,omitempty"`      //
	Policy        string `json:"policy,omitempty"`        //
	Protocol      string `json:"protocol,omitempty"`      //
	SrcCidr       string `json:"srcCidr,omitempty"`       //
	SrcPort       string `json:"srcPort,omitempty"`       //
	SyslogEnabled *bool  `json:"syslogEnabled,omitempty"` //
}
type ResponseApplianceUpdateNetworkApplianceFirewallInboundCellularFirewallRules []ResponseItemApplianceUpdateNetworkApplianceFirewallInboundCellularFirewallRules // Array of ResponseApplianceUpdateNetworkApplianceFirewallInboundCellularFirewallRules
type ResponseItemApplianceUpdateNetworkApplianceFirewallInboundCellularFirewallRules interface{}
type ResponseApplianceGetNetworkApplianceFirewallInboundFirewallRules struct {
	Rules             *[]ResponseApplianceGetNetworkApplianceFirewallInboundFirewallRulesRules `json:"rules,omitempty"`             //
	SyslogDefaultRule *bool                                                                    `json:"syslogDefaultRule,omitempty"` //
}
type ResponseApplianceGetNetworkApplianceFirewallInboundFirewallRulesRules struct {
	Comment       string `json:"comment,omitempty"`       //
	DestCidr      string `json:"destCidr,omitempty"`      //
	DestPort      string `json:"destPort,omitempty"`      //
	Policy        string `json:"policy,omitempty"`        //
	Protocol      string `json:"protocol,omitempty"`      //
	SrcCidr       string `json:"srcCidr,omitempty"`       //
	SrcPort       string `json:"srcPort,omitempty"`       //
	SyslogEnabled *bool  `json:"syslogEnabled,omitempty"` //
}
type ResponseApplianceUpdateNetworkApplianceFirewallInboundFirewallRules interface{}
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
	Policy string `json:"policy,omitempty"` //
	Type   string `json:"type,omitempty"`   //
	Value  string `json:"value,omitempty"`  //
}
type ResponseApplianceUpdateNetworkApplianceFirewallL7FirewallRules interface{}
type ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories struct {
	ApplicationCategories *[]ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesApplicationCategoriesApplicationCategories `json:"applicationCategories,omitempty"` //
}
type ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesApplicationCategoriesApplicationCategories struct {
	Applications *[]ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesApplicationCategoriesApplicationCategoriesApplications `json:"applications,omitempty"` //
	ID           string                                                                                                               `json:"id,omitempty"`           //
	Name         string                                                                                                               `json:"name,omitempty"`         //
}
type ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesApplicationCategoriesApplicationCategoriesApplications struct {
	ID   string `json:"id,omitempty"`   //
	Name string `json:"name,omitempty"` //
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
	Rules *[]ResponseApplianceGetNetworkApplianceFirewallPortForwardingRulesRules `json:"rules,omitempty"` //
}
type ResponseApplianceGetNetworkApplianceFirewallPortForwardingRulesRules struct {
	AllowedIPs []string `json:"allowedIps,omitempty"` //
	LanIP      string   `json:"lanIp,omitempty"`      //
	LocalPort  string   `json:"localPort,omitempty"`  //
	Name       string   `json:"name,omitempty"`       //
	Protocol   string   `json:"protocol,omitempty"`   //
	PublicPort string   `json:"publicPort,omitempty"` //
	Uplink     string   `json:"uplink,omitempty"`     //
}
type ResponseApplianceUpdateNetworkApplianceFirewallPortForwardingRules interface{}
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
	IDsRulesets       string                                                                  `json:"idsRulesets,omitempty"`       //
	Mode              string                                                                  `json:"mode,omitempty"`              //
	ProtectedNetworks *ResponseApplianceGetNetworkApplianceSecurityIntrusionProtectedNetworks `json:"protectedNetworks,omitempty"` //
}
type ResponseApplianceGetNetworkApplianceSecurityIntrusionProtectedNetworks struct {
	ExcludedCidr []string `json:"excludedCidr,omitempty"` //
	IncludedCidr []string `json:"includedCidr,omitempty"` //
	UseDefault   *bool    `json:"useDefault,omitempty"`   //
}
type ResponseApplianceUpdateNetworkApplianceSecurityIntrusion interface{}
type ResponseApplianceGetNetworkApplianceSecurityMalware struct {
	AllowedFiles *[]ResponseApplianceGetNetworkApplianceSecurityMalwareAllowedFiles `json:"allowedFiles,omitempty"` //
	AllowedURLs  *[]ResponseApplianceGetNetworkApplianceSecurityMalwareAllowedURLs  `json:"allowedUrls,omitempty"`  //
	Mode         string                                                             `json:"mode,omitempty"`         //
}
type ResponseApplianceGetNetworkApplianceSecurityMalwareAllowedFiles struct {
	Comment string `json:"comment,omitempty"` //
	Sha256  string `json:"sha256,omitempty"`  //
}
type ResponseApplianceGetNetworkApplianceSecurityMalwareAllowedURLs struct {
	Comment string `json:"comment,omitempty"` //
	URL     string `json:"url,omitempty"`     //
}
type ResponseApplianceUpdateNetworkApplianceSecurityMalware interface{}
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
	Enabled            *bool                                                                   `json:"enabled,omitempty"`            //
	FixedIPAssignments *ResponseItemApplianceGetNetworkApplianceStaticRoutesFixedIPAssignments `json:"fixedIpAssignments,omitempty"` //
	GatewayIP          string                                                                  `json:"gatewayIp,omitempty"`          //
	ID                 string                                                                  `json:"id,omitempty"`                 //
	IPVersion          *int                                                                    `json:"ipVersion,omitempty"`          //
	Name               string                                                                  `json:"name,omitempty"`               //
	NetworkID          string                                                                  `json:"networkId,omitempty"`          //
	ReservedIPRanges   *[]ResponseItemApplianceGetNetworkApplianceStaticRoutesReservedIPRanges `json:"reservedIpRanges,omitempty"`   //
	Subnet             string                                                                  `json:"subnet,omitempty"`             //
}
type ResponseItemApplianceGetNetworkApplianceStaticRoutesFixedIPAssignments struct {
	Status223344556677 *ResponseItemApplianceGetNetworkApplianceStaticRoutesFixedIPAssignments223344556677 `json:"22:33:44:55:66:77,omitempty"` //
}
type ResponseItemApplianceGetNetworkApplianceStaticRoutesFixedIPAssignments223344556677 struct {
	IP   string `json:"ip,omitempty"`   //
	Name string `json:"name,omitempty"` //
}
type ResponseItemApplianceGetNetworkApplianceStaticRoutesReservedIPRanges struct {
	Comment string `json:"comment,omitempty"` //
	End     string `json:"end,omitempty"`     //
	Start   string `json:"start,omitempty"`   //
}
type ResponseApplianceCreateNetworkApplianceStaticRoute interface{}
type ResponseApplianceGetNetworkApplianceStaticRoute struct {
	Enabled            *bool                                                              `json:"enabled,omitempty"`            //
	FixedIPAssignments *ResponseApplianceGetNetworkApplianceStaticRouteFixedIPAssignments `json:"fixedIpAssignments,omitempty"` //
	GatewayIP          string                                                             `json:"gatewayIp,omitempty"`          //
	ID                 string                                                             `json:"id,omitempty"`                 //
	IPVersion          *int                                                               `json:"ipVersion,omitempty"`          //
	Name               string                                                             `json:"name,omitempty"`               //
	NetworkID          string                                                             `json:"networkId,omitempty"`          //
	ReservedIPRanges   *[]ResponseApplianceGetNetworkApplianceStaticRouteReservedIPRanges `json:"reservedIpRanges,omitempty"`   //
	Subnet             string                                                             `json:"subnet,omitempty"`             //
}
type ResponseApplianceGetNetworkApplianceStaticRouteFixedIPAssignments struct {
	Status223344556677 *ResponseApplianceGetNetworkApplianceStaticRouteFixedIPAssignments223344556677 `json:"22:33:44:55:66:77,omitempty"` //
}
type ResponseApplianceGetNetworkApplianceStaticRouteFixedIPAssignments223344556677 struct {
	IP   string `json:"ip,omitempty"`   //
	Name string `json:"name,omitempty"` //
}
type ResponseApplianceGetNetworkApplianceStaticRouteReservedIPRanges struct {
	Comment string `json:"comment,omitempty"` //
	End     string `json:"end,omitempty"`     //
	Start   string `json:"start,omitempty"`   //
}
type ResponseApplianceUpdateNetworkApplianceStaticRoute interface{}
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
	CustomPerformanceClassID string `json:"customPerformanceClassId,omitempty"` //
	MaxJitter                *int   `json:"maxJitter,omitempty"`                //
	MaxLatency               *int   `json:"maxLatency,omitempty"`               //
	MaxLossPercentage        *int   `json:"maxLossPercentage,omitempty"`        //
	Name                     string `json:"name,omitempty"`                     //
}
type ResponseApplianceCreateNetworkApplianceTrafficShapingCustomPerformanceClass interface{}
type ResponseApplianceGetNetworkApplianceTrafficShapingCustomPerformanceClass struct {
	CustomPerformanceClassID string `json:"customPerformanceClassId,omitempty"` //
	MaxJitter                *int   `json:"maxJitter,omitempty"`                //
	MaxLatency               *int   `json:"maxLatency,omitempty"`               //
	MaxLossPercentage        *int   `json:"maxLossPercentage,omitempty"`        //
	Name                     string `json:"name,omitempty"`                     //
}
type ResponseApplianceUpdateNetworkApplianceTrafficShapingCustomPerformanceClass interface{}
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
	Type  string `json:"type,omitempty"`  //
	Value string `json:"value,omitempty"` //
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
	Cidr string `json:"cidr,omitempty"` // CIDR format address (e.g."192.168.10.1", which is the same as "192.168.10.1/32"), or "any"
	Port string `json:"port,omitempty"` // E.g.: "any", "0" (also means "any"), "8080", "1-1024"
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
	Cidr string `json:"cidr,omitempty"` // CIDR format address (e.g."192.168.10.1", which is the same as "192.168.10.1/32"), or "any"
	Port string `json:"port,omitempty"` // E.g.: "any", "0" (also means "any"), "8080", "1-1024"
}
type ResponseApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueSource struct {
	Cidr string `json:"cidr,omitempty"` // CIDR format address (e.g."192.168.10.1", which is the same as "192.168.10.1/32"), or "any". Cannot be used in combination with the "vlan" property
	Host *int   `json:"host,omitempty"` // Host ID in the VLAN. Should not exceed the VLAN subnet capacity. Must be used along with the "vlan" property and is currently only available under a template network.
	Port string `json:"port,omitempty"` // E.g.: "any", "0" (also means "any"), "8080", "1-1024"
	VLAN *int   `json:"vlan,omitempty"` // VLAN ID of the configured VLAN in the Meraki network. Cannot be used in combination with the "cidr" property and is currently only available under a template network.
}
type ResponseApplianceGetNetworkApplianceUplinksUsageHistory []ResponseItemApplianceGetNetworkApplianceUplinksUsageHistory // Array of ResponseApplianceGetNetworkApplianceUplinksUsageHistory
type ResponseItemApplianceGetNetworkApplianceUplinksUsageHistory struct {
	ByInterface *[]ResponseItemApplianceGetNetworkApplianceUplinksUsageHistoryByInterface `json:"byInterface,omitempty"` //
	EndTime     string                                                                    `json:"endTime,omitempty"`     //
	StartTime   string                                                                    `json:"startTime,omitempty"`   //
}
type ResponseItemApplianceGetNetworkApplianceUplinksUsageHistoryByInterface struct {
	Interface string `json:"interface,omitempty"` //
	Received  *int   `json:"received,omitempty"`  //
	Sent      *int   `json:"sent,omitempty"`      //
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
	VLANsEnabled *bool `json:"vlansEnabled,omitempty"` //
}
type ResponseApplianceUpdateNetworkApplianceVLANsSettings interface{}
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
	ID                     string                                                         `json:"id,omitempty"`                     // The VLAN ID of the VLAN
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
	AsNumber      *int                                                   `json:"asNumber,omitempty"`      //
	Enabled       *bool                                                  `json:"enabled,omitempty"`       //
	IbgpHoldTimer *int                                                   `json:"ibgpHoldTimer,omitempty"` //
	Neighbors     *[]ResponseApplianceGetNetworkApplianceVpnBgpNeighbors `json:"neighbors,omitempty"`     //
}
type ResponseApplianceGetNetworkApplianceVpnBgpNeighbors struct {
	AllowTransit   *bool  `json:"allowTransit,omitempty"`   //
	EbgpHoldTimer  *int   `json:"ebgpHoldTimer,omitempty"`  //
	EbgpMultihop   *int   `json:"ebgpMultihop,omitempty"`   //
	IP             string `json:"ip,omitempty"`             //
	ReceiveLimit   *int   `json:"receiveLimit,omitempty"`   //
	RemoteAsNumber *int   `json:"remoteAsNumber,omitempty"` //
}
type ResponseApplianceUpdateNetworkApplianceVpnBgp interface{}
type ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpn struct {
	Hubs    *[]ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnHubs    `json:"hubs,omitempty"`    // The list of VPN hubs, in order of preference.
	Mode    string                                                         `json:"mode,omitempty"`    // The site-to-site VPN mode.
	Subnets *[]ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnets `json:"subnets,omitempty"` // The list of subnets and their VPN presence.
}
type ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnHubs struct {
	HubID           string `json:"hubId,omitempty"`           // The network ID of the hub.
	UseDefaultRoute *bool  `json:"useDefaultRoute,omitempty"` // Indicates whether default route traffic should be sent to this hub.
}
type ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnets struct {
	LocalSubnet string `json:"localSubnet,omitempty"` // The CIDR notation subnet used within the VPN
	UseVpn      *bool  `json:"useVpn,omitempty"`      // Indicates the presence of the subnet in the VPN
}
type ResponseApplianceUpdateNetworkApplianceVpnSiteToSiteVpn struct {
	Hubs    *[]ResponseApplianceUpdateNetworkApplianceVpnSiteToSiteVpnHubs    `json:"hubs,omitempty"`    // The list of VPN hubs, in order of preference.
	Mode    string                                                            `json:"mode,omitempty"`    // The site-to-site VPN mode.
	Subnets *[]ResponseApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnets `json:"subnets,omitempty"` // The list of subnets and their VPN presence.
}
type ResponseApplianceUpdateNetworkApplianceVpnSiteToSiteVpnHubs struct {
	HubID           string `json:"hubId,omitempty"`           // The network ID of the hub.
	UseDefaultRoute *bool  `json:"useDefaultRoute,omitempty"` // Indicates whether default route traffic should be sent to this hub.
}
type ResponseApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnets struct {
	LocalSubnet string `json:"localSubnet,omitempty"` // The CIDR notation subnet used within the VPN
	UseVpn      *bool  `json:"useVpn,omitempty"`      // Indicates the presence of the subnet in the VPN
}
type ResponseApplianceGetNetworkApplianceWarmSpare struct {
	Enabled       *bool                                              `json:"enabled,omitempty"`       //
	PrimarySerial string                                             `json:"primarySerial,omitempty"` //
	SpareSerial   string                                             `json:"spareSerial,omitempty"`   //
	UplinkMode    string                                             `json:"uplinkMode,omitempty"`    //
	Wan1          *ResponseApplianceGetNetworkApplianceWarmSpareWan1 `json:"wan1,omitempty"`          //
	Wan2          *ResponseApplianceGetNetworkApplianceWarmSpareWan2 `json:"wan2,omitempty"`          //
}
type ResponseApplianceGetNetworkApplianceWarmSpareWan1 struct {
	IP     string `json:"ip,omitempty"`     //
	Subnet string `json:"subnet,omitempty"` //
}
type ResponseApplianceGetNetworkApplianceWarmSpareWan2 struct {
	IP     string `json:"ip,omitempty"`     //
	Subnet string `json:"subnet,omitempty"` //
}
type ResponseApplianceUpdateNetworkApplianceWarmSpare interface{}
type ResponseApplianceSwapNetworkApplianceWarmSpare interface{}
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
type ResponseApplianceGetOrganizationApplianceUplinkStatuses []ResponseItemApplianceGetOrganizationApplianceUplinkStatuses // Array of ResponseApplianceGetOrganizationApplianceUplinkStatuses
type ResponseItemApplianceGetOrganizationApplianceUplinkStatuses struct {
	HighAvailability *ResponseItemApplianceGetOrganizationApplianceUplinkStatusesHighAvailability `json:"highAvailability,omitempty"` //
	LastReportedAt   string                                                                       `json:"lastReportedAt,omitempty"`   //
	Model            string                                                                       `json:"model,omitempty"`            //
	NetworkID        string                                                                       `json:"networkId,omitempty"`        //
	Serial           string                                                                       `json:"serial,omitempty"`           //
	Uplinks          *[]ResponseItemApplianceGetOrganizationApplianceUplinkStatusesUplinks        `json:"uplinks,omitempty"`          //
}
type ResponseItemApplianceGetOrganizationApplianceUplinkStatusesHighAvailability struct {
	Enabled *bool  `json:"enabled,omitempty"` //
	Role    string `json:"role,omitempty"`    //
}
type ResponseItemApplianceGetOrganizationApplianceUplinkStatusesUplinks struct {
	Gateway      string `json:"gateway,omitempty"`      //
	Interface    string `json:"interface,omitempty"`    //
	IP           string `json:"ip,omitempty"`           //
	IPAssignedBy string `json:"ipAssignedBy,omitempty"` //
	PrimaryDNS   string `json:"primaryDns,omitempty"`   //
	PublicIP     string `json:"publicIp,omitempty"`     //
	SecondaryDNS string `json:"secondaryDns,omitempty"` //
	Status       string `json:"status,omitempty"`       //
}
type ResponseApplianceGetOrganizationApplianceVpnStats []ResponseItemApplianceGetOrganizationApplianceVpnStats // Array of ResponseApplianceGetOrganizationApplianceVpnStats
type ResponseItemApplianceGetOrganizationApplianceVpnStats struct {
	MerakiVpnpeers *[]ResponseItemApplianceGetOrganizationApplianceVpnStatsMerakiVpnpeers `json:"merakiVpnPeers,omitempty"` //
	NetworkID      string                                                                 `json:"networkId,omitempty"`      //
	NetworkName    string                                                                 `json:"networkName,omitempty"`    //
}
type ResponseItemApplianceGetOrganizationApplianceVpnStatsMerakiVpnpeers struct {
	JitterSummaries         *[]ResponseItemApplianceGetOrganizationApplianceVpnStatsMerakiVpnpeersJitterSummaries         `json:"jitterSummaries,omitempty"`         //
	LatencySummaries        *[]ResponseItemApplianceGetOrganizationApplianceVpnStatsMerakiVpnpeersLatencySummaries        `json:"latencySummaries,omitempty"`        //
	LossPercentageSummaries *[]ResponseItemApplianceGetOrganizationApplianceVpnStatsMerakiVpnpeersLossPercentageSummaries `json:"lossPercentageSummaries,omitempty"` //
	MosSummaries            *[]ResponseItemApplianceGetOrganizationApplianceVpnStatsMerakiVpnpeersMosSummaries            `json:"mosSummaries,omitempty"`            //
	NetworkID               string                                                                                        `json:"networkId,omitempty"`               //
	NetworkName             string                                                                                        `json:"networkName,omitempty"`             //
	UsageSummary            *ResponseItemApplianceGetOrganizationApplianceVpnStatsMerakiVpnpeersUsageSummary              `json:"usageSummary,omitempty"`            //
}
type ResponseItemApplianceGetOrganizationApplianceVpnStatsMerakiVpnpeersJitterSummaries struct {
	AvgJitter      *float64 `json:"avgJitter,omitempty"`      //
	MaxJitter      *float64 `json:"maxJitter,omitempty"`      //
	MinJitter      *int     `json:"minJitter,omitempty"`      //
	ReceiverUplink string   `json:"receiverUplink,omitempty"` //
	SenderUplink   string   `json:"senderUplink,omitempty"`   //
}
type ResponseItemApplianceGetOrganizationApplianceVpnStatsMerakiVpnpeersLatencySummaries struct {
	AvgLatencyMs   *int   `json:"avgLatencyMs,omitempty"`   //
	MaxLatencyMs   *int   `json:"maxLatencyMs,omitempty"`   //
	MinLatencyMs   *int   `json:"minLatencyMs,omitempty"`   //
	ReceiverUplink string `json:"receiverUplink,omitempty"` //
	SenderUplink   string `json:"senderUplink,omitempty"`   //
}
type ResponseItemApplianceGetOrganizationApplianceVpnStatsMerakiVpnpeersLossPercentageSummaries struct {
	AvgLossPercentage *int     `json:"avgLossPercentage,omitempty"` //
	MaxLossPercentage *float64 `json:"maxLossPercentage,omitempty"` //
	MinLossPercentage *int     `json:"minLossPercentage,omitempty"` //
	ReceiverUplink    string   `json:"receiverUplink,omitempty"`    //
	SenderUplink      string   `json:"senderUplink,omitempty"`      //
}
type ResponseItemApplianceGetOrganizationApplianceVpnStatsMerakiVpnpeersMosSummaries struct {
	AvgMos         *float64 `json:"avgMos,omitempty"`         //
	MaxMos         *float64 `json:"maxMos,omitempty"`         //
	MinMos         *int     `json:"minMos,omitempty"`         //
	ReceiverUplink string   `json:"receiverUplink,omitempty"` //
	SenderUplink   string   `json:"senderUplink,omitempty"`   //
}
type ResponseItemApplianceGetOrganizationApplianceVpnStatsMerakiVpnpeersUsageSummary struct {
	ReceivedInKilobytes *int `json:"receivedInKilobytes,omitempty"` //
	SentInKilobytes     *int `json:"sentInKilobytes,omitempty"`     //
}
type ResponseApplianceGetOrganizationApplianceVpnStatuses []ResponseItemApplianceGetOrganizationApplianceVpnStatuses // Array of ResponseApplianceGetOrganizationApplianceVpnStatuses
type ResponseItemApplianceGetOrganizationApplianceVpnStatuses struct {
	DeviceSerial       string                                                                        `json:"deviceSerial,omitempty"`       //
	DeviceStatus       string                                                                        `json:"deviceStatus,omitempty"`       //
	ExportedSubnets    *[]ResponseItemApplianceGetOrganizationApplianceVpnStatusesExportedSubnets    `json:"exportedSubnets,omitempty"`    //
	MerakiVpnpeers     *[]ResponseItemApplianceGetOrganizationApplianceVpnStatusesMerakiVpnpeers     `json:"merakiVpnPeers,omitempty"`     //
	NetworkID          string                                                                        `json:"networkId,omitempty"`          //
	NetworkName        string                                                                        `json:"networkName,omitempty"`        //
	ThirdPartyVpnpeers *[]ResponseItemApplianceGetOrganizationApplianceVpnStatusesThirdPartyVpnpeers `json:"thirdPartyVpnPeers,omitempty"` //
	Uplinks            *[]ResponseItemApplianceGetOrganizationApplianceVpnStatusesUplinks            `json:"uplinks,omitempty"`            //
	VpnMode            string                                                                        `json:"vpnMode,omitempty"`            //
}
type ResponseItemApplianceGetOrganizationApplianceVpnStatusesExportedSubnets struct {
	Name   string `json:"name,omitempty"`   //
	Subnet string `json:"subnet,omitempty"` //
}
type ResponseItemApplianceGetOrganizationApplianceVpnStatusesMerakiVpnpeers struct {
	NetworkID    string `json:"networkId,omitempty"`    //
	NetworkName  string `json:"networkName,omitempty"`  //
	Reachability string `json:"reachability,omitempty"` //
}
type ResponseItemApplianceGetOrganizationApplianceVpnStatusesThirdPartyVpnpeers struct {
	Name         string `json:"name,omitempty"`         //
	PublicIP     string `json:"publicIp,omitempty"`     //
	Reachability string `json:"reachability,omitempty"` //
}
type ResponseItemApplianceGetOrganizationApplianceVpnStatusesUplinks struct {
	Interface string `json:"interface,omitempty"` //
	PublicIP  string `json:"publicIp,omitempty"`  //
}
type ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeers struct {
	Peers *[]ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeersPeers `json:"peers,omitempty"` // The list of VPN peers
}
type ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeersPeers struct {
	IkeVersion          string                                                                            `json:"ikeVersion,omitempty"`          // [optional] The IKE version to be used for the IPsec VPN peer configuration. Defaults to '1' when omitted.
	IPsecPolicies       *ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeersPeersIPsecPolicies `json:"ipsecPolicies,omitempty"`       // Custom IPSec policies for the VPN peer. If not included and a preset has not been chosen, the default preset for IPSec policies will be used.
	IPsecPoliciesPreset string                                                                            `json:"ipsecPoliciesPreset,omitempty"` // One of the following available presets: 'default', 'aws', 'azure'. If this is provided, the 'ipsecPolicies' parameter is ignored.
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
	IPsecPoliciesPreset string                                                                               `json:"ipsecPoliciesPreset,omitempty"` // One of the following available presets: 'default', 'aws', 'azure'. If this is provided, the 'ipsecPolicies' parameter is ignored.
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
	Rules *[]ResponseApplianceGetOrganizationApplianceVpnVpnFirewallRulesRules `json:"rules,omitempty"` //
}
type ResponseApplianceGetOrganizationApplianceVpnVpnFirewallRulesRules struct {
	Comment       string `json:"comment,omitempty"`       //
	DestCidr      string `json:"destCidr,omitempty"`      //
	DestPort      string `json:"destPort,omitempty"`      //
	Policy        string `json:"policy,omitempty"`        //
	Protocol      string `json:"protocol,omitempty"`      //
	SrcCidr       string `json:"srcCidr,omitempty"`       //
	SrcPort       string `json:"srcPort,omitempty"`       //
	SyslogEnabled *bool  `json:"syslogEnabled,omitempty"` //
}
type ResponseApplianceUpdateOrganizationApplianceVpnVpnFirewallRules struct {
	Rules *[]ResponseApplianceUpdateOrganizationApplianceVpnVpnFirewallRulesRules `json:"rules,omitempty"` // List of firewall rules
}
type ResponseApplianceUpdateOrganizationApplianceVpnVpnFirewallRulesRules struct {
	Comment       string `json:"comment,omitempty"`       // Description of the rule
	DestCidr      string `json:"destCidr,omitempty"`      // List of destination IP addresses
	DestPort      string `json:"destPort,omitempty"`      // List of destination ports
	Policy        string `json:"policy,omitempty"`        // 'Deny' traffic specified by this rule
	Protocol      string `json:"protocol,omitempty"`      // The type of protocol
	SrcCidr       string `json:"srcCidr,omitempty"`       // List of source IP addresses
	SrcPort       string `json:"srcPort,omitempty"`       // List of source ports
	SyslogEnabled *bool  `json:"syslogEnabled,omitempty"` // Flag indicating whether the rule should be logged to syslog
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
	Policy string `json:"policy,omitempty"` // 'Deny' traffic specified by this rule
	Type   string `json:"type,omitempty"`   // Type of the L7 rule. One of: 'application', 'applicationCategory', 'host', 'port', 'ipRange'
	Value  string `json:"value,omitempty"`  // The 'value' of what you want to block. Format of 'value' varies depending on type of the rule. The application categories and application ids can be retrieved from the the 'MX L7 application categories' endpoint. The countries follow the two-letter ISO 3166-1 alpha-2 format.
}
type RequestApplianceUpdateNetworkApplianceFirewallOneToManyNatRules struct {
	Rules *[]RequestApplianceUpdateNetworkApplianceFirewallOneToManyNatRulesRules `json:"rules,omitempty"` // An array of 1:Many nat rules
}
type RequestApplianceUpdateNetworkApplianceFirewallOneToManyNatRulesRules struct {
	PortRules *[]RequestApplianceUpdateNetworkApplianceFirewallOneToManyNatRulesRulesPortRules `json:"portRules,omitempty"` // An array of associated forwarding rules
	PublicIP  string                                                                           `json:"publicIp,omitempty"`  // The IP address that will be used to access the internal resource from the WAN
	Uplink    string                                                                           `json:"uplink,omitempty"`    // The physical WAN interface on which the traffic will arrive ('internet1' or, if available, 'internet2')
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
	Uplink         string                                                                               `json:"uplink,omitempty"`         // The physical WAN interface on which the traffic will arrive ('internet1' or, if available, 'internet2')
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
type RequestApplianceUpdateNetworkApplianceSSIDRadiusServers struct {
	Host   string `json:"host,omitempty"`   // The IP address of your RADIUS server.
	Port   *int   `json:"port,omitempty"`   // The UDP port your RADIUS servers listens on for Access-requests.
	Secret string `json:"secret,omitempty"` // The RADIUS client shared secret.
}
type RequestApplianceCreateNetworkApplianceStaticRoute struct {
	GatewayIP     string `json:"gatewayIp,omitempty"`     // The gateway IP (next hop) of the static route
	GatewayVLANID string `json:"gatewayVlanId,omitempty"` // The gateway IP (next hop) VLAN ID of the static route
	Name          string `json:"name,omitempty"`          // The name of the new static route
	Subnet        string `json:"subnet,omitempty"`        // The subnet of the static route
}
type RequestApplianceUpdateNetworkApplianceStaticRoute struct {
	Enabled            *bool                                                                `json:"enabled,omitempty"`            // The enabled state of the static route
	FixedIPAssignments *RequestApplianceUpdateNetworkApplianceStaticRouteFixedIPAssignments `json:"fixedIpAssignments,omitempty"` // The DHCP fixed IP assignments on the static route. This should be an object that contains mappings from MAC addresses to objects that themselves each contain "ip" and "name" string fields. See the sample request/response for more details.
	GatewayIP          string                                                               `json:"gatewayIp,omitempty"`          // The gateway IP (next hop) of the static route
	GatewayVLANID      string                                                               `json:"gatewayVlanId,omitempty"`      // The gateway IP (next hop) VLAN ID of the static route
	Name               string                                                               `json:"name,omitempty"`               // The name of the static route
	ReservedIPRanges   *[]RequestApplianceUpdateNetworkApplianceStaticRouteReservedIPRanges `json:"reservedIpRanges,omitempty"`   // The DHCP reserved IP ranges on the static route
	Subnet             string                                                               `json:"subnet,omitempty"`             // The subnet of the static route
}
type RequestApplianceUpdateNetworkApplianceStaticRouteFixedIPAssignments interface{}
type RequestApplianceUpdateNetworkApplianceStaticRouteReservedIPRanges struct {
	Comment string `json:"comment,omitempty"` // A text comment for the reserved range
	End     string `json:"end,omitempty"`     // The last IP in the reserved range
	Start   string `json:"start,omitempty"`   // The first IP in the reserved range
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
	Type  string `json:"type,omitempty"`  // The type of definition. Can be one of 'application', 'applicationCategory', 'host', 'port', 'ipRange' or 'localNet'.
	Value string `json:"value,omitempty"` //     If "type" is 'host', 'port', 'ipRange' or 'localNet', then "value" must be a string, matching either     a hostname (e.g. "somesite.com"), a port (e.g. 8080), or an IP range ("192.1.0.0",     "192.1.0.0/16", or "10.1.0.0/16:80"). 'localNet' also supports CIDR notation, excluding     custom ports.      If "type" is 'application' or 'applicationCategory', then "value" must be an object     with the structure { "id": "meraki:layer7/..." }, where "id" is the application category or     application ID (for a list of IDs for your network, use the trafficShaping/applicationCategories     endpoint).
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
type RequestApplianceCreateNetworkApplianceVLAN struct {
	ApplianceIP      string                                                   `json:"applianceIp,omitempty"`      // The local IP of the appliance on the VLAN
	Cidr             string                                                   `json:"cidr,omitempty"`             // CIDR of the pool of subnets. Applicable only for template network. Each network bound to the template will automatically pick a subnet from this pool to build its own VLAN.
	GroupPolicyID    string                                                   `json:"groupPolicyId,omitempty"`    // The id of the desired group policy to apply to the VLAN
	ID               string                                                   `json:"id,omitempty"`               // The VLAN ID of the new VLAN (must be between 1 and 4094)
	IPv6             *RequestApplianceCreateNetworkApplianceVLANIPv6          `json:"ipv6,omitempty"`             // IPv6 configuration on the VLAN
	MandatoryDhcp    *RequestApplianceCreateNetworkApplianceVLANMandatoryDhcp `json:"mandatoryDhcp,omitempty"`    // Mandatory DHCP will enforce that clients connecting to this VLAN must use the IP address assigned by the DHCP server. Clients who use a static IP address won't be able to associate. Only available on firmware versions 17.0 and above
	Mask             *int                                                     `json:"mask,omitempty"`             // Mask used for the subnet of all bound to the template networks. Applicable only for template network.
	Name             string                                                   `json:"name,omitempty"`             // The name of the new VLAN
	Subnet           string                                                   `json:"subnet,omitempty"`           // The subnet of the VLAN
	TemplateVLANType string                                                   `json:"templateVlanType,omitempty"` // Type of subnetting of the VLAN. Applicable only for template network.
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
	IbgpHoldTimer *int                                                     `json:"ibgpHoldTimer,omitempty"` // The IBGP holdtimer in seconds. The IBGP holdtimer must be an integer between 12 and 240. When absent, this field is not updated. If no value exists then it defaults to 240.
	Neighbors     *[]RequestApplianceUpdateNetworkApplianceVpnBgpNeighbors `json:"neighbors,omitempty"`     // List of BGP neighbors. This list replaces the existing set of neighbors. When absent, this field is not updated.
}
type RequestApplianceUpdateNetworkApplianceVpnBgpNeighbors struct {
	AllowTransit    *bool                                                                `json:"allowTransit,omitempty"`    // When this feature is on, the Meraki device will advertise routes learned from other Autonomous Systems, thereby allowing traffic between Autonomous Systems to transit this AS. When absent, it defaults to false.
	Authentication  *RequestApplianceUpdateNetworkApplianceVpnBgpNeighborsAuthentication `json:"authentication,omitempty"`  // Authentication settings between BGP peers.
	EbgpHoldTimer   *int                                                                 `json:"ebgpHoldTimer,omitempty"`   // The EBGP hold timer in seconds for each neighbor. The EBGP hold timer must be an integer between 12 and 240.
	EbgpMultihop    *int                                                                 `json:"ebgpMultihop,omitempty"`    // Configure this if the neighbor is not adjacent. The EBGP multi-hop must be an integer between 1 and 255.
	IP              string                                                               `json:"ip,omitempty"`              // The IPv4 address of the neighbor
	IPv6            *RequestApplianceUpdateNetworkApplianceVpnBgpNeighborsIPv6           `json:"ipv6,omitempty"`            // Information regarding IPv6 address of the neighbor, Required if *ip* is not present.
	NextHopIP       string                                                               `json:"nextHopIp,omitempty"`       // The IPv4 address of the remote BGP peer that will establish a TCP session with the local MX.
	ReceiveLimit    *int                                                                 `json:"receiveLimit,omitempty"`    // The receive limit is the maximum number of routes that can be received from any BGP peer. The receive limit must be an integer between 0 and 4294967295. When absent, it defaults to 0.
	RemoteAsNumber  *int                                                                 `json:"remoteAsNumber,omitempty"`  // Remote ASN of the neighbor. The remote ASN must be an integer between 1 and 4294967295.
	SourceInterface string                                                               `json:"sourceInterface,omitempty"` // The output interface for peering with the remote BGP peer. Valid values are: 'wired0', 'wired1' or 'vlan{VLAN ID}'(e.g. 'vlan123').
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
	Subnets *[]RequestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnets `json:"subnets,omitempty"` // The list of subnets and their VPN presence.
}
type RequestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnHubs struct {
	HubID           string `json:"hubId,omitempty"`           // The network ID of the hub.
	UseDefaultRoute *bool  `json:"useDefaultRoute,omitempty"` // Only valid in 'spoke' mode. Indicates whether default route traffic should be sent to this hub.
}
type RequestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnets struct {
	LocalSubnet string `json:"localSubnet,omitempty"` // The CIDR notation subnet used within the VPN
	UseVpn      *bool  `json:"useVpn,omitempty"`      // Indicates the presence of the subnet in the VPN
}
type RequestApplianceUpdateNetworkApplianceWarmSpare struct {
	Enabled     *bool  `json:"enabled,omitempty"`     // Enable warm spare
	SpareSerial string `json:"spareSerial,omitempty"` // Serial number of the warm spare appliance
	UplinkMode  string `json:"uplinkMode,omitempty"`  // Uplink mode, either virtual or public
	VirtualIP1  string `json:"virtualIp1,omitempty"`  // The WAN 1 shared IP
	VirtualIP2  string `json:"virtualIp2,omitempty"`  // The WAN 2 shared IP
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
	IPsecPoliciesPreset string                                                                              `json:"ipsecPoliciesPreset,omitempty"` // One of the following available presets: 'default', 'aws', 'azure'. If this is provided, the 'ipsecPolicies' parameter is ignored.
	LocalID             string                                                                              `json:"localId,omitempty"`             // [optional] The local ID is used to identify the MX to the peer. This will apply to all MXs this peer applies to.
	Name                string                                                                              `json:"name,omitempty"`                // The name of the VPN peer
	NetworkTags         []string                                                                            `json:"networkTags,omitempty"`         // A list of network tags that will connect with this peer. Use ['all'] for all networks. Use ['none'] for no networks. If not included, the default is ['all'].
	PrivateSubnets      []string                                                                            `json:"privateSubnets,omitempty"`      // The list of the private subnets of the VPN peer
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

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-appliance-dhcp-subnets
*/
func (s *ApplianceService) GetDeviceApplianceDhcpSubnets(serial string) (*ResponseApplianceGetDeviceApplianceDhcpSubnets, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/appliance/dhcp/subnets"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetDeviceApplianceDhcpSubnets{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceApplianceDhcpSubnets")
	}

	result := response.Result().(*ResponseApplianceGetDeviceApplianceDhcpSubnets)
	return result, response, err

}

//GetDeviceAppliancePerformance Return the performance score for a single MX
/* Return the performance score for a single MX. Only primary MX devices supported. If no data is available, a 204 error code is returned.

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-appliance-performance
*/
func (s *ApplianceService) GetDeviceAppliancePerformance(serial string) (*ResponseApplianceGetDeviceAppliancePerformance, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/appliance/performance"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetDeviceAppliancePerformance{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceAppliancePerformance")
	}

	result := response.Result().(*ResponseApplianceGetDeviceAppliancePerformance)
	return result, response, err

}

//GetDeviceAppliancePrefixesDelegated Return current delegated IPv6 prefixes on an appliance.
/* Return current delegated IPv6 prefixes on an appliance.

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-appliance-prefixes-delegated
*/
func (s *ApplianceService) GetDeviceAppliancePrefixesDelegated(serial string) (*ResponseApplianceGetDeviceAppliancePrefixesDelegated, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/appliance/prefixes/delegated"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetDeviceAppliancePrefixesDelegated{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceAppliancePrefixesDelegated")
	}

	result := response.Result().(*ResponseApplianceGetDeviceAppliancePrefixesDelegated)
	return result, response, err

}

//GetDeviceAppliancePrefixesDelegatedVLANAssignments Return prefixes assigned to all IPv6 enabled VLANs on an appliance.
/* Return prefixes assigned to all IPv6 enabled VLANs on an appliance.

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-appliance-prefixes-delegated-vlan-assignments
*/
func (s *ApplianceService) GetDeviceAppliancePrefixesDelegatedVLANAssignments(serial string) (*ResponseApplianceGetDeviceAppliancePrefixesDelegatedVLANAssignments, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/appliance/prefixes/delegated/vlanAssignments"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetDeviceAppliancePrefixesDelegatedVLANAssignments{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceAppliancePrefixesDelegatedVlanAssignments")
	}

	result := response.Result().(*ResponseApplianceGetDeviceAppliancePrefixesDelegatedVLANAssignments)
	return result, response, err

}

//GetDeviceApplianceUplinksSettings Return the uplink settings for an MX appliance
/* Return the uplink settings for an MX appliance

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-appliance-uplinks-settings
*/
func (s *ApplianceService) GetDeviceApplianceUplinksSettings(serial string) (*ResponseApplianceGetDeviceApplianceUplinksSettings, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/appliance/uplinks/settings"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetDeviceApplianceUplinksSettings{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceApplianceUplinksSettings")
	}

	result := response.Result().(*ResponseApplianceGetDeviceApplianceUplinksSettings)
	return result, response, err

}

//GetNetworkApplianceClientSecurityEvents List the security events for a client
/* List the security events for a client. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.

@param networkID networkId path parameter. Network ID
@param clientID clientId path parameter. Client ID
@param getNetworkApplianceClientSecurityEventsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-appliance-client-security-events
*/
func (s *ApplianceService) GetNetworkApplianceClientSecurityEvents(networkID string, clientID string, getNetworkApplianceClientSecurityEventsQueryParams *GetNetworkApplianceClientSecurityEventsQueryParams) (*ResponseApplianceGetNetworkApplianceClientSecurityEvents, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/clients/{clientId}/security/events"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{clientId}", fmt.Sprintf("%v", clientID), -1)

	queryString, _ := query.Values(getNetworkApplianceClientSecurityEventsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseApplianceGetNetworkApplianceClientSecurityEvents{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkApplianceClientSecurityEvents")
	}

	result := response.Result().(*ResponseApplianceGetNetworkApplianceClientSecurityEvents)
	return result, response, err

}

//GetNetworkApplianceConnectivityMonitoringDestinations Return the connectivity testing destinations for an MX network
/* Return the connectivity testing destinations for an MX network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-appliance-connectivity-monitoring-destinations
*/
func (s *ApplianceService) GetNetworkApplianceConnectivityMonitoringDestinations(networkID string) (*ResponseApplianceGetNetworkApplianceConnectivityMonitoringDestinations, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/connectivityMonitoringDestinations"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetNetworkApplianceConnectivityMonitoringDestinations{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkApplianceConnectivityMonitoringDestinations")
	}

	result := response.Result().(*ResponseApplianceGetNetworkApplianceConnectivityMonitoringDestinations)
	return result, response, err

}

//GetNetworkApplianceContentFiltering Return the content filtering settings for an MX network
/* Return the content filtering settings for an MX network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-appliance-content-filtering
*/
func (s *ApplianceService) GetNetworkApplianceContentFiltering(networkID string) (*ResponseApplianceGetNetworkApplianceContentFiltering, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/contentFiltering"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetNetworkApplianceContentFiltering{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkApplianceContentFiltering")
	}

	result := response.Result().(*ResponseApplianceGetNetworkApplianceContentFiltering)
	return result, response, err

}

//GetNetworkApplianceContentFilteringCategories List all available content filtering categories for an MX network
/* List all available content filtering categories for an MX network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-appliance-content-filtering-categories
*/
func (s *ApplianceService) GetNetworkApplianceContentFilteringCategories(networkID string) (*ResponseApplianceGetNetworkApplianceContentFilteringCategories, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/contentFiltering/categories"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetNetworkApplianceContentFilteringCategories{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkApplianceContentFilteringCategories")
	}

	result := response.Result().(*ResponseApplianceGetNetworkApplianceContentFilteringCategories)
	return result, response, err

}

//GetNetworkApplianceFirewallCellularFirewallRules Return the cellular firewall rules for an MX network
/* Return the cellular firewall rules for an MX network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-appliance-firewall-cellular-firewall-rules
*/
func (s *ApplianceService) GetNetworkApplianceFirewallCellularFirewallRules(networkID string) (*ResponseApplianceGetNetworkApplianceFirewallCellularFirewallRules, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/cellularFirewallRules"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetNetworkApplianceFirewallCellularFirewallRules{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkApplianceFirewallCellularFirewallRules")
	}

	result := response.Result().(*ResponseApplianceGetNetworkApplianceFirewallCellularFirewallRules)
	return result, response, err

}

//GetNetworkApplianceFirewallFirewalledServices List the appliance services and their accessibility rules
/* List the appliance services and their accessibility rules

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-appliance-firewall-firewalled-services
*/
func (s *ApplianceService) GetNetworkApplianceFirewallFirewalledServices(networkID string) (*ResponseApplianceGetNetworkApplianceFirewallFirewalledServices, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/firewalledServices"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetNetworkApplianceFirewallFirewalledServices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkApplianceFirewallFirewalledServices")
	}

	result := response.Result().(*ResponseApplianceGetNetworkApplianceFirewallFirewalledServices)
	return result, response, err

}

//GetNetworkApplianceFirewallFirewalledService Return the accessibility settings of the given service ('ICMP', 'web', or 'SNMP')
/* Return the accessibility settings of the given service ('ICMP', 'web', or 'SNMP')

@param networkID networkId path parameter. Network ID
@param service service path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-appliance-firewall-firewalled-service
*/
func (s *ApplianceService) GetNetworkApplianceFirewallFirewalledService(networkID string, service string) (*ResponseApplianceGetNetworkApplianceFirewallFirewalledService, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/firewalledServices/{service}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{service}", fmt.Sprintf("%v", service), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetNetworkApplianceFirewallFirewalledService{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkApplianceFirewallFirewalledService")
	}

	result := response.Result().(*ResponseApplianceGetNetworkApplianceFirewallFirewalledService)
	return result, response, err

}

//GetNetworkApplianceFirewallInboundCellularFirewallRules Return the inbound cellular firewall rules for an MX network
/* Return the inbound cellular firewall rules for an MX network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-appliance-firewall-inbound-cellular-firewall-rules
*/
func (s *ApplianceService) GetNetworkApplianceFirewallInboundCellularFirewallRules(networkID string) (*ResponseApplianceGetNetworkApplianceFirewallInboundCellularFirewallRules, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/inboundCellularFirewallRules"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetNetworkApplianceFirewallInboundCellularFirewallRules{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkApplianceFirewallInboundCellularFirewallRules")
	}

	result := response.Result().(*ResponseApplianceGetNetworkApplianceFirewallInboundCellularFirewallRules)
	return result, response, err

}

//GetNetworkApplianceFirewallInboundFirewallRules Return the inbound firewall rules for an MX network
/* Return the inbound firewall rules for an MX network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-appliance-firewall-inbound-firewall-rules
*/
func (s *ApplianceService) GetNetworkApplianceFirewallInboundFirewallRules(networkID string) (*ResponseApplianceGetNetworkApplianceFirewallInboundFirewallRules, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/inboundFirewallRules"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetNetworkApplianceFirewallInboundFirewallRules{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkApplianceFirewallInboundFirewallRules")
	}

	result := response.Result().(*ResponseApplianceGetNetworkApplianceFirewallInboundFirewallRules)
	return result, response, err

}

//GetNetworkApplianceFirewallL3FirewallRules Return the L3 firewall rules for an MX network
/* Return the L3 firewall rules for an MX network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-appliance-firewall-l3-firewall-rules
*/
func (s *ApplianceService) GetNetworkApplianceFirewallL3FirewallRules(networkID string) (*ResponseApplianceGetNetworkApplianceFirewallL3FirewallRules, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/l3FirewallRules"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetNetworkApplianceFirewallL3FirewallRules{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkApplianceFirewallL3FirewallRules")
	}

	result := response.Result().(*ResponseApplianceGetNetworkApplianceFirewallL3FirewallRules)
	return result, response, err

}

//GetNetworkApplianceFirewallL7FirewallRules List the MX L7 firewall rules for an MX network
/* List the MX L7 firewall rules for an MX network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-appliance-firewall-l7-firewall-rules
*/
func (s *ApplianceService) GetNetworkApplianceFirewallL7FirewallRules(networkID string) (*ResponseApplianceGetNetworkApplianceFirewallL7FirewallRules, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/l7FirewallRules"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetNetworkApplianceFirewallL7FirewallRules{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkApplianceFirewallL7FirewallRules")
	}

	result := response.Result().(*ResponseApplianceGetNetworkApplianceFirewallL7FirewallRules)
	return result, response, err

}

//GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories Return the L7 firewall application categories and their associated applications for an MX network
/* Return the L7 firewall application categories and their associated applications for an MX network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-appliance-firewall-l7-firewall-rules-application-categories
*/
func (s *ApplianceService) GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories(networkID string) (*ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/l7FirewallRules/applicationCategories"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories")
	}

	result := response.Result().(*ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories)
	return result, response, err

}

//GetNetworkApplianceFirewallOneToManyNatRules Return the 1:Many NAT mapping rules for an MX network
/* Return the 1:Many NAT mapping rules for an MX network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-appliance-firewall-one-to-many-nat-rules
*/
func (s *ApplianceService) GetNetworkApplianceFirewallOneToManyNatRules(networkID string) (*ResponseApplianceGetNetworkApplianceFirewallOneToManyNatRules, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/oneToManyNatRules"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetNetworkApplianceFirewallOneToManyNatRules{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkApplianceFirewallOneToManyNatRules")
	}

	result := response.Result().(*ResponseApplianceGetNetworkApplianceFirewallOneToManyNatRules)
	return result, response, err

}

//GetNetworkApplianceFirewallOneToOneNatRules Return the 1:1 NAT mapping rules for an MX network
/* Return the 1:1 NAT mapping rules for an MX network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-appliance-firewall-one-to-one-nat-rules
*/
func (s *ApplianceService) GetNetworkApplianceFirewallOneToOneNatRules(networkID string) (*ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRules, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/oneToOneNatRules"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRules{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkApplianceFirewallOneToOneNatRules")
	}

	result := response.Result().(*ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRules)
	return result, response, err

}

//GetNetworkApplianceFirewallPortForwardingRules Return the port forwarding rules for an MX network
/* Return the port forwarding rules for an MX network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-appliance-firewall-port-forwarding-rules
*/
func (s *ApplianceService) GetNetworkApplianceFirewallPortForwardingRules(networkID string) (*ResponseApplianceGetNetworkApplianceFirewallPortForwardingRules, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/portForwardingRules"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetNetworkApplianceFirewallPortForwardingRules{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkApplianceFirewallPortForwardingRules")
	}

	result := response.Result().(*ResponseApplianceGetNetworkApplianceFirewallPortForwardingRules)
	return result, response, err

}

//GetNetworkApplianceFirewallSettings Return the firewall settings for this network
/* Return the firewall settings for this network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-appliance-firewall-settings
*/
func (s *ApplianceService) GetNetworkApplianceFirewallSettings(networkID string) (*ResponseApplianceGetNetworkApplianceFirewallSettings, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/settings"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetNetworkApplianceFirewallSettings{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkApplianceFirewallSettings")
	}

	result := response.Result().(*ResponseApplianceGetNetworkApplianceFirewallSettings)
	return result, response, err

}

//GetNetworkAppliancePorts List per-port VLAN settings for all ports of a MX.
/* List per-port VLAN settings for all ports of a MX.

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-appliance-ports
*/
func (s *ApplianceService) GetNetworkAppliancePorts(networkID string) (*ResponseApplianceGetNetworkAppliancePorts, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/ports"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetNetworkAppliancePorts{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkAppliancePorts")
	}

	result := response.Result().(*ResponseApplianceGetNetworkAppliancePorts)
	return result, response, err

}

//GetNetworkAppliancePort Return per-port VLAN settings for a single MX port.
/* Return per-port VLAN settings for a single MX port.

@param networkID networkId path parameter. Network ID
@param portID portId path parameter. Port ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-appliance-port
*/
func (s *ApplianceService) GetNetworkAppliancePort(networkID string, portID string) (*ResponseApplianceGetNetworkAppliancePort, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/ports/{portId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{portId}", fmt.Sprintf("%v", portID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetNetworkAppliancePort{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkAppliancePort")
	}

	result := response.Result().(*ResponseApplianceGetNetworkAppliancePort)
	return result, response, err

}

//GetNetworkAppliancePrefixesDelegatedStatics List static delegated prefixes for a network
/* List static delegated prefixes for a network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-appliance-prefixes-delegated-statics
*/
func (s *ApplianceService) GetNetworkAppliancePrefixesDelegatedStatics(networkID string) (*ResponseApplianceGetNetworkAppliancePrefixesDelegatedStatics, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/prefixes/delegated/statics"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetNetworkAppliancePrefixesDelegatedStatics{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkAppliancePrefixesDelegatedStatics")
	}

	result := response.Result().(*ResponseApplianceGetNetworkAppliancePrefixesDelegatedStatics)
	return result, response, err

}

//GetNetworkAppliancePrefixesDelegatedStatic Return a static delegated prefix from a network
/* Return a static delegated prefix from a network

@param networkID networkId path parameter. Network ID
@param staticDelegatedPrefixID staticDelegatedPrefixId path parameter. Static delegated prefix ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-appliance-prefixes-delegated-static
*/
func (s *ApplianceService) GetNetworkAppliancePrefixesDelegatedStatic(networkID string, staticDelegatedPrefixID string) (*ResponseApplianceGetNetworkAppliancePrefixesDelegatedStatic, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/prefixes/delegated/statics/{staticDelegatedPrefixId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{staticDelegatedPrefixId}", fmt.Sprintf("%v", staticDelegatedPrefixID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetNetworkAppliancePrefixesDelegatedStatic{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkAppliancePrefixesDelegatedStatic")
	}

	result := response.Result().(*ResponseApplianceGetNetworkAppliancePrefixesDelegatedStatic)
	return result, response, err

}

//GetNetworkApplianceSecurityEvents List the security events for a network
/* List the security events for a network

@param networkID networkId path parameter. Network ID
@param getNetworkApplianceSecurityEventsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-appliance-security-events
*/
func (s *ApplianceService) GetNetworkApplianceSecurityEvents(networkID string, getNetworkApplianceSecurityEventsQueryParams *GetNetworkApplianceSecurityEventsQueryParams) (*ResponseApplianceGetNetworkApplianceSecurityEvents, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/security/events"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(getNetworkApplianceSecurityEventsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseApplianceGetNetworkApplianceSecurityEvents{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkApplianceSecurityEvents")
	}

	result := response.Result().(*ResponseApplianceGetNetworkApplianceSecurityEvents)
	return result, response, err

}

//GetNetworkApplianceSecurityIntrusion Returns all supported intrusion settings for an MX network
/* Returns all supported intrusion settings for an MX network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-appliance-security-intrusion
*/
func (s *ApplianceService) GetNetworkApplianceSecurityIntrusion(networkID string) (*ResponseApplianceGetNetworkApplianceSecurityIntrusion, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/security/intrusion"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetNetworkApplianceSecurityIntrusion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkApplianceSecurityIntrusion")
	}

	result := response.Result().(*ResponseApplianceGetNetworkApplianceSecurityIntrusion)
	return result, response, err

}

//GetNetworkApplianceSecurityMalware Returns all supported malware settings for an MX network
/* Returns all supported malware settings for an MX network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-appliance-security-malware
*/
func (s *ApplianceService) GetNetworkApplianceSecurityMalware(networkID string) (*ResponseApplianceGetNetworkApplianceSecurityMalware, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/security/malware"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetNetworkApplianceSecurityMalware{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkApplianceSecurityMalware")
	}

	result := response.Result().(*ResponseApplianceGetNetworkApplianceSecurityMalware)
	return result, response, err

}

//GetNetworkApplianceSettings Return the appliance settings for a network
/* Return the appliance settings for a network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-appliance-settings
*/
func (s *ApplianceService) GetNetworkApplianceSettings(networkID string) (*ResponseApplianceGetNetworkApplianceSettings, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/settings"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetNetworkApplianceSettings{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkApplianceSettings")
	}

	result := response.Result().(*ResponseApplianceGetNetworkApplianceSettings)
	return result, response, err

}

//GetNetworkApplianceSingleLan Return single LAN configuration
/* Return single LAN configuration

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-appliance-single-lan
*/
func (s *ApplianceService) GetNetworkApplianceSingleLan(networkID string) (*ResponseApplianceGetNetworkApplianceSingleLan, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/singleLan"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetNetworkApplianceSingleLan{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkApplianceSingleLan")
	}

	result := response.Result().(*ResponseApplianceGetNetworkApplianceSingleLan)
	return result, response, err

}

//GetNetworkApplianceSSIDs List the MX SSIDs in a network
/* List the MX SSIDs in a network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-appliance-ssids
*/
func (s *ApplianceService) GetNetworkApplianceSSIDs(networkID string) (*ResponseApplianceGetNetworkApplianceSSIDs, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/ssids"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetNetworkApplianceSSIDs{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkApplianceSsids")
	}

	result := response.Result().(*ResponseApplianceGetNetworkApplianceSSIDs)
	return result, response, err

}

//GetNetworkApplianceSSID Return a single MX SSID
/* Return a single MX SSID

@param networkID networkId path parameter. Network ID
@param number number path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-appliance-ssid
*/
func (s *ApplianceService) GetNetworkApplianceSSID(networkID string, number string) (*ResponseApplianceGetNetworkApplianceSSID, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/ssids/{number}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{number}", fmt.Sprintf("%v", number), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetNetworkApplianceSSID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkApplianceSsid")
	}

	result := response.Result().(*ResponseApplianceGetNetworkApplianceSSID)
	return result, response, err

}

//GetNetworkApplianceStaticRoutes List the static routes for an MX or teleworker network
/* List the static routes for an MX or teleworker network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-appliance-static-routes
*/
func (s *ApplianceService) GetNetworkApplianceStaticRoutes(networkID string) (*ResponseApplianceGetNetworkApplianceStaticRoutes, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/staticRoutes"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetNetworkApplianceStaticRoutes{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkApplianceStaticRoutes")
	}

	result := response.Result().(*ResponseApplianceGetNetworkApplianceStaticRoutes)
	return result, response, err

}

//GetNetworkApplianceStaticRoute Return a static route for an MX or teleworker network
/* Return a static route for an MX or teleworker network

@param networkID networkId path parameter. Network ID
@param staticRouteID staticRouteId path parameter. Static route ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-appliance-static-route
*/
func (s *ApplianceService) GetNetworkApplianceStaticRoute(networkID string, staticRouteID string) (*ResponseApplianceGetNetworkApplianceStaticRoute, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/staticRoutes/{staticRouteId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{staticRouteId}", fmt.Sprintf("%v", staticRouteID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetNetworkApplianceStaticRoute{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkApplianceStaticRoute")
	}

	result := response.Result().(*ResponseApplianceGetNetworkApplianceStaticRoute)
	return result, response, err

}

//GetNetworkApplianceTrafficShaping Display the traffic shaping settings for an MX network
/* Display the traffic shaping settings for an MX network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-appliance-traffic-shaping
*/
func (s *ApplianceService) GetNetworkApplianceTrafficShaping(networkID string) (*ResponseApplianceGetNetworkApplianceTrafficShaping, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/trafficShaping"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetNetworkApplianceTrafficShaping{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkApplianceTrafficShaping")
	}

	result := response.Result().(*ResponseApplianceGetNetworkApplianceTrafficShaping)
	return result, response, err

}

//GetNetworkApplianceTrafficShapingCustomPerformanceClasses List all custom performance classes for an MX network
/* List all custom performance classes for an MX network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-appliance-traffic-shaping-custom-performance-classes
*/
func (s *ApplianceService) GetNetworkApplianceTrafficShapingCustomPerformanceClasses(networkID string) (*ResponseApplianceGetNetworkApplianceTrafficShapingCustomPerformanceClasses, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/trafficShaping/customPerformanceClasses"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetNetworkApplianceTrafficShapingCustomPerformanceClasses{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkApplianceTrafficShapingCustomPerformanceClasses")
	}

	result := response.Result().(*ResponseApplianceGetNetworkApplianceTrafficShapingCustomPerformanceClasses)
	return result, response, err

}

//GetNetworkApplianceTrafficShapingCustomPerformanceClass Return a custom performance class for an MX network
/* Return a custom performance class for an MX network

@param networkID networkId path parameter. Network ID
@param customPerformanceClassID customPerformanceClassId path parameter. Custom performance class ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-appliance-traffic-shaping-custom-performance-class
*/
func (s *ApplianceService) GetNetworkApplianceTrafficShapingCustomPerformanceClass(networkID string, customPerformanceClassID string) (*ResponseApplianceGetNetworkApplianceTrafficShapingCustomPerformanceClass, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/trafficShaping/customPerformanceClasses/{customPerformanceClassId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{customPerformanceClassId}", fmt.Sprintf("%v", customPerformanceClassID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetNetworkApplianceTrafficShapingCustomPerformanceClass{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkApplianceTrafficShapingCustomPerformanceClass")
	}

	result := response.Result().(*ResponseApplianceGetNetworkApplianceTrafficShapingCustomPerformanceClass)
	return result, response, err

}

//GetNetworkApplianceTrafficShapingRules Display the traffic shaping settings rules for an MX network
/* Display the traffic shaping settings rules for an MX network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-appliance-traffic-shaping-rules
*/
func (s *ApplianceService) GetNetworkApplianceTrafficShapingRules(networkID string) (*ResponseApplianceGetNetworkApplianceTrafficShapingRules, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/trafficShaping/rules"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetNetworkApplianceTrafficShapingRules{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkApplianceTrafficShapingRules")
	}

	result := response.Result().(*ResponseApplianceGetNetworkApplianceTrafficShapingRules)
	return result, response, err

}

//GetNetworkApplianceTrafficShapingUplinkBandwidth Returns the uplink bandwidth limits for your MX network
/* Returns the uplink bandwidth limits for your MX network. This may not reflect the affected device's hardware capabilities.  For more information on your device's hardware capabilities, please consult our MX Family Datasheet [https://meraki.cisco.com/product-collateral/mx-family-datasheet/?file]

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-appliance-traffic-shaping-uplink-bandwidth
*/
func (s *ApplianceService) GetNetworkApplianceTrafficShapingUplinkBandwidth(networkID string) (*ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidth, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/trafficShaping/uplinkBandwidth"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidth{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkApplianceTrafficShapingUplinkBandwidth")
	}

	result := response.Result().(*ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidth)
	return result, response, err

}

//GetNetworkApplianceTrafficShapingUplinkSelection Show uplink selection settings for an MX network
/* Show uplink selection settings for an MX network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-appliance-traffic-shaping-uplink-selection
*/
func (s *ApplianceService) GetNetworkApplianceTrafficShapingUplinkSelection(networkID string) (*ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelection, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/trafficShaping/uplinkSelection"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelection{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkApplianceTrafficShapingUplinkSelection")
	}

	result := response.Result().(*ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelection)
	return result, response, err

}

//GetNetworkApplianceUplinksUsageHistory Get the sent and received bytes for each uplink of a network.
/* Get the sent and received bytes for each uplink of a network.

@param networkID networkId path parameter. Network ID
@param getNetworkApplianceUplinksUsageHistoryQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-appliance-uplinks-usage-history
*/
func (s *ApplianceService) GetNetworkApplianceUplinksUsageHistory(networkID string, getNetworkApplianceUplinksUsageHistoryQueryParams *GetNetworkApplianceUplinksUsageHistoryQueryParams) (*ResponseApplianceGetNetworkApplianceUplinksUsageHistory, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/uplinks/usageHistory"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(getNetworkApplianceUplinksUsageHistoryQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseApplianceGetNetworkApplianceUplinksUsageHistory{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkApplianceUplinksUsageHistory")
	}

	result := response.Result().(*ResponseApplianceGetNetworkApplianceUplinksUsageHistory)
	return result, response, err

}

//GetNetworkApplianceVLANs List the VLANs for an MX network
/* List the VLANs for an MX network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-appliance-vlans
*/
func (s *ApplianceService) GetNetworkApplianceVLANs(networkID string) (*ResponseApplianceGetNetworkApplianceVLANs, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/vlans"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetNetworkApplianceVLANs{}).
		SetError(&Error).
		Get(path)

	// log.Printf("[DEBUG] Response: %v", response)
	log.Printf("[DEBUG] Error: %v", err)
	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkApplianceVlans")
	}

	log.Printf("[DEBUG] Selected method: GetDeviceCameraAnalyticsLive")
	result := response.Result().(*ResponseApplianceGetNetworkApplianceVLANs)
	log.Println("Result: ", result)
	return result, response, err

}

//GetNetworkApplianceVLANsSettings Returns the enabled status of VLANs for the network
/* Returns the enabled status of VLANs for the network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-appliance-vlans-settings
*/
func (s *ApplianceService) GetNetworkApplianceVLANsSettings(networkID string) (*ResponseApplianceGetNetworkApplianceVLANsSettings, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/vlans/settings"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetNetworkApplianceVLANsSettings{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkApplianceVlansSettings")
	}

	result := response.Result().(*ResponseApplianceGetNetworkApplianceVLANsSettings)
	return result, response, err

}

//GetNetworkApplianceVLAN Return a VLAN
/* Return a VLAN

@param networkID networkId path parameter. Network ID
@param vlanID vlanId path parameter. Vlan ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-appliance-vlan
*/
func (s *ApplianceService) GetNetworkApplianceVLAN(networkID string, vlanID string) (*ResponseApplianceGetNetworkApplianceVLAN, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/vlans/{vlanId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{vlanId}", fmt.Sprintf("%v", vlanID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetNetworkApplianceVLAN{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkApplianceVlan")
	}

	result := response.Result().(*ResponseApplianceGetNetworkApplianceVLAN)
	return result, response, err

}

//GetNetworkApplianceVpnBgp Return a Hub BGP Configuration
/* Return a Hub BGP Configuration

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-appliance-vpn-bgp
*/
func (s *ApplianceService) GetNetworkApplianceVpnBgp(networkID string) (*ResponseApplianceGetNetworkApplianceVpnBgp, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/vpn/bgp"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetNetworkApplianceVpnBgp{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkApplianceVpnBgp")
	}

	result := response.Result().(*ResponseApplianceGetNetworkApplianceVpnBgp)
	return result, response, err

}

//GetNetworkApplianceVpnSiteToSiteVpn Return the site-to-site VPN settings of a network
/* Return the site-to-site VPN settings of a network. Only valid for MX networks.

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-appliance-vpn-site-to-site-vpn
*/
func (s *ApplianceService) GetNetworkApplianceVpnSiteToSiteVpn(networkID string) (*ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpn, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/vpn/siteToSiteVpn"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpn{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkApplianceVpnSiteToSiteVpn")
	}

	result := response.Result().(*ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpn)
	return result, response, err

}

//GetNetworkApplianceWarmSpare Return MX warm spare settings
/* Return MX warm spare settings

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-appliance-warm-spare
*/
func (s *ApplianceService) GetNetworkApplianceWarmSpare(networkID string) (*ResponseApplianceGetNetworkApplianceWarmSpare, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/warmSpare"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetNetworkApplianceWarmSpare{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkApplianceWarmSpare")
	}

	result := response.Result().(*ResponseApplianceGetNetworkApplianceWarmSpare)
	return result, response, err

}

//GetOrganizationApplianceSecurityEvents List the security events for an organization
/* List the security events for an organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationApplianceSecurityEventsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-appliance-security-events
*/
func (s *ApplianceService) GetOrganizationApplianceSecurityEvents(organizationID string, getOrganizationApplianceSecurityEventsQueryParams *GetOrganizationApplianceSecurityEventsQueryParams) (*ResponseApplianceGetOrganizationApplianceSecurityEvents, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/appliance/security/events"
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationApplianceSecurityEventsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseApplianceGetOrganizationApplianceSecurityEvents{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationApplianceSecurityEvents")
	}

	result := response.Result().(*ResponseApplianceGetOrganizationApplianceSecurityEvents)
	return result, response, err

}

//GetOrganizationApplianceSecurityIntrusion Returns all supported intrusion settings for an organization
/* Returns all supported intrusion settings for an organization

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-appliance-security-intrusion
*/
func (s *ApplianceService) GetOrganizationApplianceSecurityIntrusion(organizationID string) (*ResponseApplianceGetOrganizationApplianceSecurityIntrusion, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/appliance/security/intrusion"
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetOrganizationApplianceSecurityIntrusion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationApplianceSecurityIntrusion")
	}

	result := response.Result().(*ResponseApplianceGetOrganizationApplianceSecurityIntrusion)
	return result, response, err

}

//GetOrganizationApplianceUplinkStatuses List the uplink status of every Meraki MX and Z series appliances in the organization
/* List the uplink status of every Meraki MX and Z series appliances in the organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationApplianceUplinkStatusesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-appliance-uplink-statuses
*/
func (s *ApplianceService) GetOrganizationApplianceUplinkStatuses(organizationID string, getOrganizationApplianceUplinkStatusesQueryParams *GetOrganizationApplianceUplinkStatusesQueryParams) (*ResponseApplianceGetOrganizationApplianceUplinkStatuses, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/appliance/uplink/statuses"
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationApplianceUplinkStatusesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseApplianceGetOrganizationApplianceUplinkStatuses{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationApplianceUplinkStatuses")
	}

	result := response.Result().(*ResponseApplianceGetOrganizationApplianceUplinkStatuses)
	return result, response, err

}

//GetOrganizationApplianceVpnStats Show VPN history stat for networks in an organization
/* Show VPN history stat for networks in an organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationApplianceVpnStatsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-appliance-vpn-stats
*/
func (s *ApplianceService) GetOrganizationApplianceVpnStats(organizationID string, getOrganizationApplianceVpnStatsQueryParams *GetOrganizationApplianceVpnStatsQueryParams) (*ResponseApplianceGetOrganizationApplianceVpnStats, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/appliance/vpn/stats"
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationApplianceVpnStatsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseApplianceGetOrganizationApplianceVpnStats{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationApplianceVpnStats")
	}

	result := response.Result().(*ResponseApplianceGetOrganizationApplianceVpnStats)
	return result, response, err

}

//GetOrganizationApplianceVpnStatuses Show VPN status for networks in an organization
/* Show VPN status for networks in an organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationApplianceVpnStatusesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-appliance-vpn-statuses
*/
func (s *ApplianceService) GetOrganizationApplianceVpnStatuses(organizationID string, getOrganizationApplianceVpnStatusesQueryParams *GetOrganizationApplianceVpnStatusesQueryParams) (*ResponseApplianceGetOrganizationApplianceVpnStatuses, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/appliance/vpn/statuses"
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationApplianceVpnStatusesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseApplianceGetOrganizationApplianceVpnStatuses{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationApplianceVpnStatuses")
	}

	result := response.Result().(*ResponseApplianceGetOrganizationApplianceVpnStatuses)
	return result, response, err

}

//GetOrganizationApplianceVpnThirdPartyVpnpeers Return the third party VPN peers for an organization
/* Return the third party VPN peers for an organization

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-appliance-vpn-third-party-vpnpeers
*/
func (s *ApplianceService) GetOrganizationApplianceVpnThirdPartyVpnpeers(organizationID string) (*ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeers, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/appliance/vpn/thirdPartyVPNPeers"
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeers{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationApplianceVpnThirdPartyVpnpeers")
	}

	result := response.Result().(*ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeers)
	return result, response, err

}

//GetOrganizationApplianceVpnVpnFirewallRules Return the firewall rules for an organization's site-to-site VPN
/* Return the firewall rules for an organization's site-to-site VPN

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-appliance-vpn-vpn-firewall-rules
*/
func (s *ApplianceService) GetOrganizationApplianceVpnVpnFirewallRules(organizationID string) (*ResponseApplianceGetOrganizationApplianceVpnVpnFirewallRules, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/appliance/vpn/vpnFirewallRules"
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceGetOrganizationApplianceVpnVpnFirewallRules{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationApplianceVpnVpnFirewallRules")
	}

	result := response.Result().(*ResponseApplianceGetOrganizationApplianceVpnVpnFirewallRules)
	return result, response, err

}

//CreateDeviceApplianceVmxAuthenticationToken Generate a new vMX authentication token
/* Generate a new vMX authentication token

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-device-appliance-vmx-authentication-token
*/

func (s *ApplianceService) CreateDeviceApplianceVmxAuthenticationToken(serial string) (*ResponseApplianceCreateDeviceApplianceVmxAuthenticationToken, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/appliance/vmx/authenticationToken"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplianceCreateDeviceApplianceVmxAuthenticationToken{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateDeviceApplianceVmxAuthenticationToken")
	}

	result := response.Result().(*ResponseApplianceCreateDeviceApplianceVmxAuthenticationToken)
	return result, response, err

}

//CreateNetworkAppliancePrefixesDelegatedStatic Add a static delegated prefix from a network
/* Add a static delegated prefix from a network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-network-appliance-prefixes-delegated-static
*/

func (s *ApplianceService) CreateNetworkAppliancePrefixesDelegatedStatic(networkID string, requestApplianceCreateNetworkAppliancePrefixesDelegatedStatic *RequestApplianceCreateNetworkAppliancePrefixesDelegatedStatic) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/prefixes/delegated/statics"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplianceCreateNetworkAppliancePrefixesDelegatedStatic).
		// SetResult(&ResponseApplianceCreateNetworkAppliancePrefixesDelegatedStatic{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateNetworkAppliancePrefixesDelegatedStatic")
	}

	return response, err

}

//CreateNetworkApplianceStaticRoute Add a static route for an MX or teleworker network
/* Add a static route for an MX or teleworker network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-network-appliance-static-route
*/

func (s *ApplianceService) CreateNetworkApplianceStaticRoute(networkID string, requestApplianceCreateNetworkApplianceStaticRoute *RequestApplianceCreateNetworkApplianceStaticRoute) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/staticRoutes"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplianceCreateNetworkApplianceStaticRoute).
		// SetResult(&ResponseApplianceCreateNetworkApplianceStaticRoute{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateNetworkApplianceStaticRoute")
	}

	return response, err

}

//CreateNetworkApplianceTrafficShapingCustomPerformanceClass Add a custom performance class for an MX network
/* Add a custom performance class for an MX network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-network-appliance-traffic-shaping-custom-performance-class
*/

func (s *ApplianceService) CreateNetworkApplianceTrafficShapingCustomPerformanceClass(networkID string, requestApplianceCreateNetworkApplianceTrafficShapingCustomPerformanceClass *RequestApplianceCreateNetworkApplianceTrafficShapingCustomPerformanceClass) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/trafficShaping/customPerformanceClasses"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplianceCreateNetworkApplianceTrafficShapingCustomPerformanceClass).
		// SetResult(&ResponseApplianceCreateNetworkApplianceTrafficShapingCustomPerformanceClass{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateNetworkApplianceTrafficShapingCustomPerformanceClass")
	}

	return response, err

}

//CreateNetworkApplianceVLAN Add a VLAN
/* Add a VLAN

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-network-appliance-vlan
*/

func (s *ApplianceService) CreateNetworkApplianceVLAN(networkID string, requestApplianceCreateNetworkApplianceVlan *RequestApplianceCreateNetworkApplianceVLAN) (*ResponseApplianceCreateNetworkApplianceVLAN, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/vlans"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplianceCreateNetworkApplianceVlan).
		SetResult(&ResponseApplianceCreateNetworkApplianceVLAN{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateNetworkApplianceVlan")
	}

	result := response.Result().(*ResponseApplianceCreateNetworkApplianceVLAN)
	return result, response, err

}

//SwapNetworkApplianceWarmSpare Swap MX primary and warm spare appliances
/* Swap MX primary and warm spare appliances

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!swap-network-appliance-warm-spare
*/

func (s *ApplianceService) SwapNetworkApplianceWarmSpare(networkID string) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/warmSpare/swap"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").

		// SetResult(&ResponseApplianceSwapNetworkApplianceWarmSpare{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation SwapNetworkApplianceWarmSpare")
	}

	return response, err

}

//UpdateDeviceApplianceUplinksSettings Update the uplink settings for an MX appliance
/* Update the uplink settings for an MX appliance

@param serial serial path parameter.
*/
func (s *ApplianceService) UpdateDeviceApplianceUplinksSettings(serial string, requestApplianceUpdateDeviceApplianceUplinksSettings *RequestApplianceUpdateDeviceApplianceUplinksSettings) (*ResponseApplianceUpdateDeviceApplianceUplinksSettings, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/appliance/uplinks/settings"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplianceUpdateDeviceApplianceUplinksSettings).
		SetResult(&ResponseApplianceUpdateDeviceApplianceUplinksSettings{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateDeviceApplianceUplinksSettings")
	}

	result := response.Result().(*ResponseApplianceUpdateDeviceApplianceUplinksSettings)
	return result, response, err

}

//UpdateNetworkApplianceConnectivityMonitoringDestinations Update the connectivity testing destinations for an MX network
/* Update the connectivity testing destinations for an MX network

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceConnectivityMonitoringDestinations(networkID string, requestApplianceUpdateNetworkApplianceConnectivityMonitoringDestinations *RequestApplianceUpdateNetworkApplianceConnectivityMonitoringDestinations) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/connectivityMonitoringDestinations"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplianceUpdateNetworkApplianceConnectivityMonitoringDestinations).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkApplianceConnectivityMonitoringDestinations")
	}

	return response, err

}

//UpdateNetworkApplianceContentFiltering Update the content filtering settings for an MX network
/* Update the content filtering settings for an MX network

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceContentFiltering(networkID string, requestApplianceUpdateNetworkApplianceContentFiltering *RequestApplianceUpdateNetworkApplianceContentFiltering) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/contentFiltering"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplianceUpdateNetworkApplianceContentFiltering).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkApplianceContentFiltering")
	}

	return response, err

}

//UpdateNetworkApplianceFirewallCellularFirewallRules Update the cellular firewall rules of an MX network
/* Update the cellular firewall rules of an MX network

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceFirewallCellularFirewallRules(networkID string, requestApplianceUpdateNetworkApplianceFirewallCellularFirewallRules *RequestApplianceUpdateNetworkApplianceFirewallCellularFirewallRules) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/cellularFirewallRules"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplianceUpdateNetworkApplianceFirewallCellularFirewallRules).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkApplianceFirewallCellularFirewallRules")
	}

	return response, err

}

//UpdateNetworkApplianceFirewallFirewalledService Updates the accessibility settings for the given service ('ICMP', 'web', or 'SNMP')
/* Updates the accessibility settings for the given service ('ICMP', 'web', or 'SNMP')

@param networkID networkId path parameter. Network ID
@param service service path parameter.
*/
func (s *ApplianceService) UpdateNetworkApplianceFirewallFirewalledService(networkID string, service string, requestApplianceUpdateNetworkApplianceFirewallFirewalledService *RequestApplianceUpdateNetworkApplianceFirewallFirewalledService) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/firewalledServices/{service}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{service}", fmt.Sprintf("%v", service), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplianceUpdateNetworkApplianceFirewallFirewalledService).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkApplianceFirewallFirewalledService")
	}

	return response, err

}

//UpdateNetworkApplianceFirewallInboundCellularFirewallRules Update the inbound cellular firewall rules of an MX network
/* Update the inbound cellular firewall rules of an MX network

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceFirewallInboundCellularFirewallRules(networkID string, requestApplianceUpdateNetworkApplianceFirewallInboundCellularFirewallRules *RequestApplianceUpdateNetworkApplianceFirewallInboundCellularFirewallRules) (*ResponseApplianceUpdateNetworkApplianceFirewallInboundCellularFirewallRules, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/inboundCellularFirewallRules"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplianceUpdateNetworkApplianceFirewallInboundCellularFirewallRules).
		SetResult(&ResponseApplianceUpdateNetworkApplianceFirewallInboundCellularFirewallRules{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkApplianceFirewallInboundCellularFirewallRules")
	}

	result := response.Result().(*ResponseApplianceUpdateNetworkApplianceFirewallInboundCellularFirewallRules)
	return result, response, err

}

//UpdateNetworkApplianceFirewallInboundFirewallRules Update the inbound firewall rules of an MX network
/* Update the inbound firewall rules of an MX network

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceFirewallInboundFirewallRules(networkID string, requestApplianceUpdateNetworkApplianceFirewallInboundFirewallRules *RequestApplianceUpdateNetworkApplianceFirewallInboundFirewallRules) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/inboundFirewallRules"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplianceUpdateNetworkApplianceFirewallInboundFirewallRules).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkApplianceFirewallInboundFirewallRules")
	}

	return response, err

}

//UpdateNetworkApplianceFirewallL3FirewallRules Update the L3 firewall rules of an MX network
/* Update the L3 firewall rules of an MX network

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceFirewallL3FirewallRules(networkID string, requestApplianceUpdateNetworkApplianceFirewallL3FirewallRules *RequestApplianceUpdateNetworkApplianceFirewallL3FirewallRules) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/l3FirewallRules"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplianceUpdateNetworkApplianceFirewallL3FirewallRules).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkApplianceFirewallL3FirewallRules")
	}

	return response, err

}

//UpdateNetworkApplianceFirewallL7FirewallRules Update the MX L7 firewall rules for an MX network
/* Update the MX L7 firewall rules for an MX network

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceFirewallL7FirewallRules(networkID string, requestApplianceUpdateNetworkApplianceFirewallL7FirewallRules *RequestApplianceUpdateNetworkApplianceFirewallL7FirewallRules) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/l7FirewallRules"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplianceUpdateNetworkApplianceFirewallL7FirewallRules).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkApplianceFirewallL7FirewallRules")
	}

	return response, err

}

//UpdateNetworkApplianceFirewallOneToManyNatRules Set the 1:Many NAT mapping rules for an MX network
/* Set the 1:Many NAT mapping rules for an MX network

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceFirewallOneToManyNatRules(networkID string, requestApplianceUpdateNetworkApplianceFirewallOneToManyNatRules *RequestApplianceUpdateNetworkApplianceFirewallOneToManyNatRules) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/oneToManyNatRules"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplianceUpdateNetworkApplianceFirewallOneToManyNatRules).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkApplianceFirewallOneToManyNatRules")
	}

	return response, err

}

//UpdateNetworkApplianceFirewallOneToOneNatRules Set the 1:1 NAT mapping rules for an MX network
/* Set the 1:1 NAT mapping rules for an MX network

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceFirewallOneToOneNatRules(networkID string, requestApplianceUpdateNetworkApplianceFirewallOneToOneNatRules *RequestApplianceUpdateNetworkApplianceFirewallOneToOneNatRules) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/oneToOneNatRules"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplianceUpdateNetworkApplianceFirewallOneToOneNatRules).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkApplianceFirewallOneToOneNatRules")
	}

	return response, err

}

//UpdateNetworkApplianceFirewallPortForwardingRules Update the port forwarding rules for an MX network
/* Update the port forwarding rules for an MX network

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceFirewallPortForwardingRules(networkID string, requestApplianceUpdateNetworkApplianceFirewallPortForwardingRules *RequestApplianceUpdateNetworkApplianceFirewallPortForwardingRules) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/portForwardingRules"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplianceUpdateNetworkApplianceFirewallPortForwardingRules).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkApplianceFirewallPortForwardingRules")
	}

	return response, err

}

//UpdateNetworkApplianceFirewallSettings Update the firewall settings for this network
/* Update the firewall settings for this network

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceFirewallSettings(networkID string, requestApplianceUpdateNetworkApplianceFirewallSettings *RequestApplianceUpdateNetworkApplianceFirewallSettings) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/firewall/settings"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplianceUpdateNetworkApplianceFirewallSettings).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkApplianceFirewallSettings")
	}

	return response, err

}

//UpdateNetworkAppliancePort Update the per-port VLAN settings for a single MX port.
/* Update the per-port VLAN settings for a single MX port.

@param networkID networkId path parameter. Network ID
@param portID portId path parameter. Port ID
*/
func (s *ApplianceService) UpdateNetworkAppliancePort(networkID string, portID string, requestApplianceUpdateNetworkAppliancePort *RequestApplianceUpdateNetworkAppliancePort) (*ResponseApplianceUpdateNetworkAppliancePort, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/ports/{portId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{portId}", fmt.Sprintf("%v", portID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplianceUpdateNetworkAppliancePort).
		SetResult(&ResponseApplianceUpdateNetworkAppliancePort{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkAppliancePort")
	}

	result := response.Result().(*ResponseApplianceUpdateNetworkAppliancePort)
	return result, response, err

}

//UpdateNetworkAppliancePrefixesDelegatedStatic Update a static delegated prefix from a network
/* Update a static delegated prefix from a network

@param networkID networkId path parameter. Network ID
@param staticDelegatedPrefixID staticDelegatedPrefixId path parameter. Static delegated prefix ID
*/
func (s *ApplianceService) UpdateNetworkAppliancePrefixesDelegatedStatic(networkID string, staticDelegatedPrefixID string, requestApplianceUpdateNetworkAppliancePrefixesDelegatedStatic *RequestApplianceUpdateNetworkAppliancePrefixesDelegatedStatic) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/prefixes/delegated/statics/{staticDelegatedPrefixId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{staticDelegatedPrefixId}", fmt.Sprintf("%v", staticDelegatedPrefixID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplianceUpdateNetworkAppliancePrefixesDelegatedStatic).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkAppliancePrefixesDelegatedStatic")
	}

	return response, err

}

//UpdateNetworkApplianceSecurityIntrusion Set the supported intrusion settings for an MX network
/* Set the supported intrusion settings for an MX network

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceSecurityIntrusion(networkID string, requestApplianceUpdateNetworkApplianceSecurityIntrusion *RequestApplianceUpdateNetworkApplianceSecurityIntrusion) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/security/intrusion"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplianceUpdateNetworkApplianceSecurityIntrusion).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkApplianceSecurityIntrusion")
	}

	return response, err

}

//UpdateNetworkApplianceSecurityMalware Set the supported malware settings for an MX network
/* Set the supported malware settings for an MX network

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceSecurityMalware(networkID string, requestApplianceUpdateNetworkApplianceSecurityMalware *RequestApplianceUpdateNetworkApplianceSecurityMalware) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/security/malware"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplianceUpdateNetworkApplianceSecurityMalware).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkApplianceSecurityMalware")
	}

	return response, err

}

//UpdateNetworkApplianceSettings Update the appliance settings for a network
/* Update the appliance settings for a network

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceSettings(networkID string, requestApplianceUpdateNetworkApplianceSettings *RequestApplianceUpdateNetworkApplianceSettings) (*ResponseApplianceUpdateNetworkApplianceSettings, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/settings"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplianceUpdateNetworkApplianceSettings).
		SetResult(&ResponseApplianceUpdateNetworkApplianceSettings{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkApplianceSettings")
	}

	result := response.Result().(*ResponseApplianceUpdateNetworkApplianceSettings)
	return result, response, err

}

//UpdateNetworkApplianceSingleLan Update single LAN configuration
/* Update single LAN configuration

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceSingleLan(networkID string, requestApplianceUpdateNetworkApplianceSingleLan *RequestApplianceUpdateNetworkApplianceSingleLan) (*ResponseApplianceUpdateNetworkApplianceSingleLan, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/singleLan"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplianceUpdateNetworkApplianceSingleLan).
		SetResult(&ResponseApplianceUpdateNetworkApplianceSingleLan{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkApplianceSingleLan")
	}

	result := response.Result().(*ResponseApplianceUpdateNetworkApplianceSingleLan)
	return result, response, err

}

//UpdateNetworkApplianceSSID Update the attributes of an MX SSID
/* Update the attributes of an MX SSID

@param networkID networkId path parameter. Network ID
@param number number path parameter.
*/
func (s *ApplianceService) UpdateNetworkApplianceSSID(networkID string, number string, requestApplianceUpdateNetworkApplianceSsid *RequestApplianceUpdateNetworkApplianceSSID) (*ResponseApplianceUpdateNetworkApplianceSSID, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/ssids/{number}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{number}", fmt.Sprintf("%v", number), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplianceUpdateNetworkApplianceSsid).
		SetResult(&ResponseApplianceUpdateNetworkApplianceSSID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkApplianceSsid")
	}

	result := response.Result().(*ResponseApplianceUpdateNetworkApplianceSSID)
	return result, response, err

}

//UpdateNetworkApplianceStaticRoute Update a static route for an MX or teleworker network
/* Update a static route for an MX or teleworker network

@param networkID networkId path parameter. Network ID
@param staticRouteID staticRouteId path parameter. Static route ID
*/
func (s *ApplianceService) UpdateNetworkApplianceStaticRoute(networkID string, staticRouteID string, requestApplianceUpdateNetworkApplianceStaticRoute *RequestApplianceUpdateNetworkApplianceStaticRoute) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/staticRoutes/{staticRouteId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{staticRouteId}", fmt.Sprintf("%v", staticRouteID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplianceUpdateNetworkApplianceStaticRoute).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkApplianceStaticRoute")
	}

	return response, err

}

//UpdateNetworkApplianceTrafficShaping Update the traffic shaping settings for an MX network
/* Update the traffic shaping settings for an MX network

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceTrafficShaping(networkID string, requestApplianceUpdateNetworkApplianceTrafficShaping *RequestApplianceUpdateNetworkApplianceTrafficShaping) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/trafficShaping"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplianceUpdateNetworkApplianceTrafficShaping).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkApplianceTrafficShaping")
	}

	return response, err

}

//UpdateNetworkApplianceTrafficShapingCustomPerformanceClass Update a custom performance class for an MX network
/* Update a custom performance class for an MX network

@param networkID networkId path parameter. Network ID
@param customPerformanceClassID customPerformanceClassId path parameter. Custom performance class ID
*/
func (s *ApplianceService) UpdateNetworkApplianceTrafficShapingCustomPerformanceClass(networkID string, customPerformanceClassID string, requestApplianceUpdateNetworkApplianceTrafficShapingCustomPerformanceClass *RequestApplianceUpdateNetworkApplianceTrafficShapingCustomPerformanceClass) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/trafficShaping/customPerformanceClasses/{customPerformanceClassId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{customPerformanceClassId}", fmt.Sprintf("%v", customPerformanceClassID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplianceUpdateNetworkApplianceTrafficShapingCustomPerformanceClass).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkApplianceTrafficShapingCustomPerformanceClass")
	}

	return response, err

}

//UpdateNetworkApplianceTrafficShapingRules Update the traffic shaping settings rules for an MX network
/* Update the traffic shaping settings rules for an MX network

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceTrafficShapingRules(networkID string, requestApplianceUpdateNetworkApplianceTrafficShapingRules *RequestApplianceUpdateNetworkApplianceTrafficShapingRules) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/trafficShaping/rules"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplianceUpdateNetworkApplianceTrafficShapingRules).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkApplianceTrafficShapingRules")
	}

	return response, err

}

//UpdateNetworkApplianceTrafficShapingUplinkBandwidth Updates the uplink bandwidth settings for your MX network.
/* Updates the uplink bandwidth settings for your MX network.

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceTrafficShapingUplinkBandwidth(networkID string, requestApplianceUpdateNetworkApplianceTrafficShapingUplinkBandwidth *RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkBandwidth) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/trafficShaping/uplinkBandwidth"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplianceUpdateNetworkApplianceTrafficShapingUplinkBandwidth).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkApplianceTrafficShapingUplinkBandwidth")
	}

	return response, err

}

//UpdateNetworkApplianceTrafficShapingUplinkSelection Update uplink selection settings for an MX network
/* Update uplink selection settings for an MX network

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceTrafficShapingUplinkSelection(networkID string, requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelection *RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelection) (*ResponseApplianceUpdateNetworkApplianceTrafficShapingUplinkSelection, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/trafficShaping/uplinkSelection"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelection).
		SetResult(&ResponseApplianceUpdateNetworkApplianceTrafficShapingUplinkSelection{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkApplianceTrafficShapingUplinkSelection")
	}

	result := response.Result().(*ResponseApplianceUpdateNetworkApplianceTrafficShapingUplinkSelection)
	return result, response, err

}

//UpdateNetworkApplianceVLANsSettings Enable/Disable VLANs for the given network
/* Enable/Disable VLANs for the given network

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceVLANsSettings(networkID string, requestApplianceUpdateNetworkApplianceVlansSettings *RequestApplianceUpdateNetworkApplianceVLANsSettings) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/vlans/settings"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplianceUpdateNetworkApplianceVlansSettings).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkApplianceVlansSettings")
	}

	return response, err

}

//UpdateNetworkApplianceVLAN Update a VLAN
/* Update a VLAN

@param networkID networkId path parameter. Network ID
@param vlanID vlanId path parameter. Vlan ID
*/
func (s *ApplianceService) UpdateNetworkApplianceVLAN(networkID string, vlanID string, requestApplianceUpdateNetworkApplianceVlan *RequestApplianceUpdateNetworkApplianceVLAN) (*ResponseApplianceUpdateNetworkApplianceVLAN, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/vlans/{vlanId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{vlanId}", fmt.Sprintf("%v", vlanID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplianceUpdateNetworkApplianceVlan).
		SetResult(&ResponseApplianceUpdateNetworkApplianceVLAN{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkApplianceVlan")
	}

	result := response.Result().(*ResponseApplianceUpdateNetworkApplianceVLAN)
	return result, response, err

}

//UpdateNetworkApplianceVpnBgp Update a Hub BGP Configuration
/* Update a Hub BGP Configuration

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceVpnBgp(networkID string, requestApplianceUpdateNetworkApplianceVpnBgp *RequestApplianceUpdateNetworkApplianceVpnBgp) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/vpn/bgp"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplianceUpdateNetworkApplianceVpnBgp).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkApplianceVpnBgp")
	}

	return response, err

}

//UpdateNetworkApplianceVpnSiteToSiteVpn Update the site-to-site VPN settings of a network
/* Update the site-to-site VPN settings of a network. Only valid for MX networks in NAT mode.

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceVpnSiteToSiteVpn(networkID string, requestApplianceUpdateNetworkApplianceVpnSiteToSiteVpn *RequestApplianceUpdateNetworkApplianceVpnSiteToSiteVpn) (*ResponseApplianceUpdateNetworkApplianceVpnSiteToSiteVpn, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/vpn/siteToSiteVpn"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplianceUpdateNetworkApplianceVpnSiteToSiteVpn).
		SetResult(&ResponseApplianceUpdateNetworkApplianceVpnSiteToSiteVpn{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkApplianceVpnSiteToSiteVpn")
	}

	result := response.Result().(*ResponseApplianceUpdateNetworkApplianceVpnSiteToSiteVpn)
	return result, response, err

}

//UpdateNetworkApplianceWarmSpare Update MX warm spare settings
/* Update MX warm spare settings

@param networkID networkId path parameter. Network ID
*/
func (s *ApplianceService) UpdateNetworkApplianceWarmSpare(networkID string, requestApplianceUpdateNetworkApplianceWarmSpare *RequestApplianceUpdateNetworkApplianceWarmSpare) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/appliance/warmSpare"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplianceUpdateNetworkApplianceWarmSpare).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkApplianceWarmSpare")
	}

	return response, err

}

//UpdateOrganizationApplianceSecurityIntrusion Sets supported intrusion settings for an organization
/* Sets supported intrusion settings for an organization

@param organizationID organizationId path parameter. Organization ID
*/
func (s *ApplianceService) UpdateOrganizationApplianceSecurityIntrusion(organizationID string, requestApplianceUpdateOrganizationApplianceSecurityIntrusion *RequestApplianceUpdateOrganizationApplianceSecurityIntrusion) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/appliance/security/intrusion"
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplianceUpdateOrganizationApplianceSecurityIntrusion).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateOrganizationApplianceSecurityIntrusion")
	}

	return response, err

}

//UpdateOrganizationApplianceVpnThirdPartyVpnpeers Update the third party VPN peers for an organization
/* Update the third party VPN peers for an organization

@param organizationID organizationId path parameter. Organization ID
*/
func (s *ApplianceService) UpdateOrganizationApplianceVpnThirdPartyVpnpeers(organizationID string, requestApplianceUpdateOrganizationApplianceVpnThirdPartyVPNPeers *RequestApplianceUpdateOrganizationApplianceVpnThirdPartyVpnpeers) (*ResponseApplianceUpdateOrganizationApplianceVpnThirdPartyVpnpeers, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/appliance/vpn/thirdPartyVPNPeers"
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplianceUpdateOrganizationApplianceVpnThirdPartyVPNPeers).
		SetResult(&ResponseApplianceUpdateOrganizationApplianceVpnThirdPartyVpnpeers{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateOrganizationApplianceVpnThirdPartyVpnpeers")
	}

	result := response.Result().(*ResponseApplianceUpdateOrganizationApplianceVpnThirdPartyVpnpeers)
	return result, response, err

}

//UpdateOrganizationApplianceVpnVpnFirewallRules Update the firewall rules of an organization's site-to-site VPN
/* Update the firewall rules of an organization's site-to-site VPN

@param organizationID organizationId path parameter. Organization ID
*/
func (s *ApplianceService) UpdateOrganizationApplianceVpnVpnFirewallRules(organizationID string, requestApplianceUpdateOrganizationApplianceVpnVpnFirewallRules *RequestApplianceUpdateOrganizationApplianceVpnVpnFirewallRules) (*ResponseApplianceUpdateOrganizationApplianceVpnVpnFirewallRules, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/appliance/vpn/vpnFirewallRules"
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplianceUpdateOrganizationApplianceVpnVpnFirewallRules).
		SetResult(&ResponseApplianceUpdateOrganizationApplianceVpnVpnFirewallRules{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateOrganizationApplianceVpnVpnFirewallRules")
	}

	result := response.Result().(*ResponseApplianceUpdateOrganizationApplianceVpnVpnFirewallRules)
	return result, response, err

}

//DeleteNetworkAppliancePrefixesDelegatedStatic Delete a static delegated prefix from a network
/* Delete a static delegated prefix from a network

@param networkID networkId path parameter. Network ID
@param staticDelegatedPrefixID staticDelegatedPrefixId path parameter. Static delegated prefix ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-network-appliance-prefixes-delegated-static
*/
func (s *ApplianceService) DeleteNetworkAppliancePrefixesDelegatedStatic(networkID string, staticDelegatedPrefixID string) (*resty.Response, error) {
	//networkID string,staticDelegatedPrefixID string
	path := "/api/v1/networks/{networkId}/appliance/prefixes/delegated/statics/{staticDelegatedPrefixId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{staticDelegatedPrefixId}", fmt.Sprintf("%v", staticDelegatedPrefixID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteNetworkAppliancePrefixesDelegatedStatic")
	}

	return response, err

}

//DeleteNetworkApplianceStaticRoute Delete a static route from an MX or teleworker network
/* Delete a static route from an MX or teleworker network

@param networkID networkId path parameter. Network ID
@param staticRouteID staticRouteId path parameter. Static route ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-network-appliance-static-route
*/
func (s *ApplianceService) DeleteNetworkApplianceStaticRoute(networkID string, staticRouteID string) (*resty.Response, error) {
	//networkID string,staticRouteID string
	path := "/api/v1/networks/{networkId}/appliance/staticRoutes/{staticRouteId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
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
		return response, fmt.Errorf("error with operation DeleteNetworkApplianceStaticRoute")
	}

	return response, err

}

//DeleteNetworkApplianceTrafficShapingCustomPerformanceClass Delete a custom performance class from an MX network
/* Delete a custom performance class from an MX network

@param networkID networkId path parameter. Network ID
@param customPerformanceClassID customPerformanceClassId path parameter. Custom performance class ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-network-appliance-traffic-shaping-custom-performance-class
*/
func (s *ApplianceService) DeleteNetworkApplianceTrafficShapingCustomPerformanceClass(networkID string, customPerformanceClassID string) (*resty.Response, error) {
	//networkID string,customPerformanceClassID string
	path := "/api/v1/networks/{networkId}/appliance/trafficShaping/customPerformanceClasses/{customPerformanceClassId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{customPerformanceClassId}", fmt.Sprintf("%v", customPerformanceClassID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteNetworkApplianceTrafficShapingCustomPerformanceClass")
	}

	return response, err

}

//DeleteNetworkApplianceVLAN Delete a VLAN from a network
/* Delete a VLAN from a network

@param networkID networkId path parameter. Network ID
@param vlanID vlanId path parameter. Vlan ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-network-appliance-vlan
*/
func (s *ApplianceService) DeleteNetworkApplianceVLAN(networkID string, vlanID string) (*resty.Response, error) {
	//networkID string,vlanID string
	path := "/api/v1/networks/{networkId}/appliance/vlans/{vlanId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{vlanId}", fmt.Sprintf("%v", vlanID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteNetworkApplianceVlan")
	}

	return response, err

}
