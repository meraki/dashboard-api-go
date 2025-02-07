package meraki

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type CellularGatewayService service

type GetOrganizationCellularGatewayEsimsInventoryQueryParams struct {
	Eids []string `url:"eids[],omitempty"` //Optional parameter to filter the results by EID.
}
type GetOrganizationCellularGatewayEsimsServiceProvidersAccountsQueryParams struct {
	AccountIDs []string `url:"accountIds[],omitempty"` //Optional parameter to filter the results by service provider account IDs.
}
type GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansQueryParams struct {
	AccountIDs []string `url:"accountIds[],omitempty"` //Account IDs that communication plans will be fetched for
}
type GetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlansQueryParams struct {
	AccountIDs []string `url:"accountIds[],omitempty"` //Account IDs that rate plans will be fetched for
}
type GetOrganizationCellularGatewayUplinkStatusesQueryParams struct {
	PerPage       int      `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter string   `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string   `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	NetworkIDs    []string `url:"networkIds[],omitempty"`  //A list of network IDs. The returned devices will be filtered to only include these networks.
	Serials       []string `url:"serials[],omitempty"`     //A list of serial numbers. The returned devices will be filtered to only include these serials.
	Iccids        []string `url:"iccids[],omitempty"`      //A list of ICCIDs. The returned devices will be filtered to only include these ICCIDs.
}

type ResponseCellularGatewayGetDeviceCellularGatewayLan struct {
	DeviceLanIP        string                                                                  `json:"deviceLanIp,omitempty"`        // Lan IP of the MG
	DeviceName         string                                                                  `json:"deviceName,omitempty"`         // Name of the MG.
	DeviceSubnet       string                                                                  `json:"deviceSubnet,omitempty"`       // Subnet configuration of the MG.
	FixedIPAssignments *[]ResponseCellularGatewayGetDeviceCellularGatewayLanFixedIPAssignments `json:"fixedIpAssignments,omitempty"` // list of all fixed IP assignments for a single MG
	ReservedIPRanges   *[]ResponseCellularGatewayGetDeviceCellularGatewayLanReservedIPRanges   `json:"reservedIpRanges,omitempty"`   // list of all reserved IP ranges for a single MG
}
type ResponseCellularGatewayGetDeviceCellularGatewayLanFixedIPAssignments struct {
	IP   string `json:"ip,omitempty"`   // The IP address you want to assign to a specific server or device
	Mac  string `json:"mac,omitempty"`  // The MAC address of the server or device that hosts the internal resource that you wish to receive the specified IP address
	Name string `json:"name,omitempty"` // A descriptive name of the assignment
}
type ResponseCellularGatewayGetDeviceCellularGatewayLanReservedIPRanges struct {
	Comment string `json:"comment,omitempty"` // Comment explaining the reserved IP range
	End     string `json:"end,omitempty"`     // Ending IP included in the reserved range of IPs
	Start   string `json:"start,omitempty"`   // Starting IP included in the reserved range of IPs
}
type ResponseCellularGatewayUpdateDeviceCellularGatewayLan struct {
	DeviceLanIP        string                                                                     `json:"deviceLanIp,omitempty"`        // Lan IP of the MG
	DeviceName         string                                                                     `json:"deviceName,omitempty"`         // Name of the MG.
	DeviceSubnet       string                                                                     `json:"deviceSubnet,omitempty"`       // Subnet configuration of the MG.
	FixedIPAssignments *[]ResponseCellularGatewayUpdateDeviceCellularGatewayLanFixedIPAssignments `json:"fixedIpAssignments,omitempty"` // list of all fixed IP assignments for a single MG
	ReservedIPRanges   *[]ResponseCellularGatewayUpdateDeviceCellularGatewayLanReservedIPRanges   `json:"reservedIpRanges,omitempty"`   // list of all reserved IP ranges for a single MG
}
type ResponseCellularGatewayUpdateDeviceCellularGatewayLanFixedIPAssignments struct {
	IP   string `json:"ip,omitempty"`   // The IP address you want to assign to a specific server or device
	Mac  string `json:"mac,omitempty"`  // The MAC address of the server or device that hosts the internal resource that you wish to receive the specified IP address
	Name string `json:"name,omitempty"` // A descriptive name of the assignment
}
type ResponseCellularGatewayUpdateDeviceCellularGatewayLanReservedIPRanges struct {
	Comment string `json:"comment,omitempty"` // Comment explaining the reserved IP range
	End     string `json:"end,omitempty"`     // Ending IP included in the reserved range of IPs
	Start   string `json:"start,omitempty"`   // Starting IP included in the reserved range of IPs
}
type ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRules struct {
	Rules *[]ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRulesRules `json:"rules,omitempty"` // An array of port forwarding params
}
type ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRulesRules struct {
	Access     string   `json:"access,omitempty"`     // *any* or *restricted*. Specify the right to make inbound connections on the specified ports or port ranges. If *restricted*, a list of allowed IPs is mandatory.
	AllowedIPs []string `json:"allowedIps,omitempty"` // An array of ranges of WAN IP addresses that are allowed to make inbound connections on the specified ports or port ranges.
	LanIP      string   `json:"lanIp,omitempty"`      // The IP address of the server or device that hosts the internal resource that you wish to make available on the WAN
	LocalPort  string   `json:"localPort,omitempty"`  // A port or port ranges that will receive the forwarded traffic from the WAN
	Name       string   `json:"name,omitempty"`       // A descriptive name for the rule
	Protocol   string   `json:"protocol,omitempty"`   // TCP or UDP
	PublicPort string   `json:"publicPort,omitempty"` // A port or port ranges that will be forwarded to the host on the LAN
}
type ResponseCellularGatewayUpdateDeviceCellularGatewayPortForwardingRules struct {
	Rules *[]ResponseCellularGatewayUpdateDeviceCellularGatewayPortForwardingRulesRules `json:"rules,omitempty"` // An array of port forwarding params
}
type ResponseCellularGatewayUpdateDeviceCellularGatewayPortForwardingRulesRules struct {
	Access     string   `json:"access,omitempty"`     // *any* or *restricted*. Specify the right to make inbound connections on the specified ports or port ranges. If *restricted*, a list of allowed IPs is mandatory.
	AllowedIPs []string `json:"allowedIps,omitempty"` // An array of ranges of WAN IP addresses that are allowed to make inbound connections on the specified ports or port ranges.
	LanIP      string   `json:"lanIp,omitempty"`      // The IP address of the server or device that hosts the internal resource that you wish to make available on the WAN
	LocalPort  string   `json:"localPort,omitempty"`  // A port or port ranges that will receive the forwarded traffic from the WAN
	Name       string   `json:"name,omitempty"`       // A descriptive name for the rule
	Protocol   string   `json:"protocol,omitempty"`   // TCP or UDP
	PublicPort string   `json:"publicPort,omitempty"` // A port or port ranges that will be forwarded to the host on the LAN
}
type ResponseCellularGatewayGetNetworkCellularGatewayConnectivityMonitoringDestinations struct {
	Destinations *[]ResponseCellularGatewayGetNetworkCellularGatewayConnectivityMonitoringDestinationsDestinations `json:"destinations,omitempty"` // The list of connectivity monitoring destinations
}
type ResponseCellularGatewayGetNetworkCellularGatewayConnectivityMonitoringDestinationsDestinations struct {
	Default     *bool  `json:"default,omitempty"`     // Boolean indicating whether this is the default testing destination (true) or not (false). Defaults to false. Only one default is allowed
	Description string `json:"description,omitempty"` // Description of the testing destination. Optional, defaults to an empty string
	IP          string `json:"ip,omitempty"`          // The IP address to test connectivity with
}
type ResponseCellularGatewayUpdateNetworkCellularGatewayConnectivityMonitoringDestinations struct {
	Destinations *[]ResponseCellularGatewayUpdateNetworkCellularGatewayConnectivityMonitoringDestinationsDestinations `json:"destinations,omitempty"` // The list of connectivity monitoring destinations
}
type ResponseCellularGatewayUpdateNetworkCellularGatewayConnectivityMonitoringDestinationsDestinations struct {
	Default     *bool  `json:"default,omitempty"`     // Boolean indicating whether this is the default testing destination (true) or not (false). Defaults to false. Only one default is allowed
	Description string `json:"description,omitempty"` // Description of the testing destination. Optional, defaults to an empty string
	IP          string `json:"ip,omitempty"`          // The IP address to test connectivity with
}
type ResponseCellularGatewayGetNetworkCellularGatewayDhcp struct {
	DhcpLeaseTime        string   `json:"dhcpLeaseTime,omitempty"`        // DHCP Lease time for all MG in the network.
	DNSCustomNameservers []string `json:"dnsCustomNameservers,omitempty"` // List of fixed IPs representing the the DNS Name servers when the mode is 'custom'.
	DNSNameservers       string   `json:"dnsNameservers,omitempty"`       // DNS name servers mode for all MG in the network.
}
type ResponseCellularGatewayUpdateNetworkCellularGatewayDhcp struct {
	DhcpLeaseTime        string   `json:"dhcpLeaseTime,omitempty"`        // DHCP Lease time for all MG in the network.
	DNSCustomNameservers []string `json:"dnsCustomNameservers,omitempty"` // List of fixed IPs representing the the DNS Name servers when the mode is 'custom'.
	DNSNameservers       string   `json:"dnsNameservers,omitempty"`       // DNS name servers mode for all MG in the network.
}
type ResponseCellularGatewayGetNetworkCellularGatewaySubnetPool struct {
	Cidr           string                                                               `json:"cidr,omitempty"`           // CIDR of the pool of subnets. Each MG in this network will automatically pick a subnet from this pool.
	DeploymentMode string                                                               `json:"deploymentMode,omitempty"` // Deployment mode for the cellular gateways in the network. (Passthrough/Routed)
	Mask           *int                                                                 `json:"mask,omitempty"`           // Mask used for the subnet of all MGs in  this network.
	Subnets        *[]ResponseCellularGatewayGetNetworkCellularGatewaySubnetPoolSubnets `json:"subnets,omitempty"`        // List of subnets of all MGs in this network.
}
type ResponseCellularGatewayGetNetworkCellularGatewaySubnetPoolSubnets struct {
	ApplianceIP string `json:"applianceIp,omitempty"` // Appliance IP of the MG device.
	Name        string `json:"name,omitempty"`        // Name of the MG.
	Serial      string `json:"serial,omitempty"`      // Serial Number of the MG.
	Subnet      string `json:"subnet,omitempty"`      // Subnet of MG device.
}
type ResponseCellularGatewayUpdateNetworkCellularGatewaySubnetPool struct {
	Cidr           string                                                                  `json:"cidr,omitempty"`           // CIDR of the pool of subnets. Each MG in this network will automatically pick a subnet from this pool.
	DeploymentMode string                                                                  `json:"deploymentMode,omitempty"` // Deployment mode for the cellular gateways in the network. (Passthrough/Routed)
	Mask           *int                                                                    `json:"mask,omitempty"`           // Mask used for the subnet of all MGs in  this network.
	Subnets        *[]ResponseCellularGatewayUpdateNetworkCellularGatewaySubnetPoolSubnets `json:"subnets,omitempty"`        // List of subnets of all MGs in this network.
}
type ResponseCellularGatewayUpdateNetworkCellularGatewaySubnetPoolSubnets struct {
	ApplianceIP string `json:"applianceIp,omitempty"` // Appliance IP of the MG device.
	Name        string `json:"name,omitempty"`        // Name of the MG.
	Serial      string `json:"serial,omitempty"`      // Serial Number of the MG.
	Subnet      string `json:"subnet,omitempty"`      // Subnet of MG device.
}
type ResponseCellularGatewayGetNetworkCellularGatewayUplink struct {
	BandwidthLimits *ResponseCellularGatewayGetNetworkCellularGatewayUplinkBandwidthLimits `json:"bandwidthLimits,omitempty"` // The bandwidth settings for the 'cellular' uplink
}
type ResponseCellularGatewayGetNetworkCellularGatewayUplinkBandwidthLimits struct {
	LimitDown *int `json:"limitDown,omitempty"` // The maximum download limit (integer, in Kbps). 'null' indicates no limit.
	LimitUp   *int `json:"limitUp,omitempty"`   // The maximum upload limit (integer, in Kbps). 'null' indicates no limit.
}
type ResponseCellularGatewayUpdateNetworkCellularGatewayUplink struct {
	BandwidthLimits *ResponseCellularGatewayUpdateNetworkCellularGatewayUplinkBandwidthLimits `json:"bandwidthLimits,omitempty"` // The bandwidth settings for the 'cellular' uplink
}
type ResponseCellularGatewayUpdateNetworkCellularGatewayUplinkBandwidthLimits struct {
	LimitDown *int `json:"limitDown,omitempty"` // The maximum download limit (integer, in Kbps). 'null' indicates no limit.
	LimitUp   *int `json:"limitUp,omitempty"`   // The maximum upload limit (integer, in Kbps). 'null' indicates no limit.
}
type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventory []ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsInventory // Array of ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventory
type ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsInventory struct {
	Items *[]ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsInventoryItems `json:"items,omitempty"` // List of eSIM Devices
	Meta  *ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsInventoryMeta    `json:"meta,omitempty"`  // Meta details about the result
}
type ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsInventoryItems struct {
	Active        *bool                                                                                   `json:"active,omitempty"`        // Whether eSIM is currently active SIM on Device
	Device        *ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsInventoryItemsDevice     `json:"device,omitempty"`        // Meraki Device properties
	Eid           string                                                                                  `json:"eid,omitempty"`           // eSIM EID
	LastUpdatedAt string                                                                                  `json:"lastUpdatedAt,omitempty"` // Last update of eSIM
	Network       *ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsInventoryItemsNetwork    `json:"network,omitempty"`       // Meraki Network properties
	Profiles      *[]ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsInventoryItemsProfiles `json:"profiles,omitempty"`      // eSIM Profile Information
}
type ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsInventoryItemsDevice struct {
	Model  string `json:"model,omitempty"`  // Device model
	Name   string `json:"name,omitempty"`   // Device name
	Serial string `json:"serial,omitempty"` // Device serial number
	Status string `json:"status,omitempty"` // Device status
	URL    string `json:"url,omitempty"`    // Device URL
}
type ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsInventoryItemsNetwork struct {
	ID string `json:"id,omitempty"` // Network ID for this eSIM
}
type ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsInventoryItemsProfiles struct {
	CustomApns      []string                                                                                             `json:"customApns,omitempty"`      // Available custom APNs for the profile
	Iccid           string                                                                                               `json:"iccid,omitempty"`           // eSIM profile ID
	ServiceProvider *ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsInventoryItemsProfilesServiceProvider `json:"serviceProvider,omitempty"` // Service Provider information
	Status          string                                                                                               `json:"status,omitempty"`          // eSIM profile status
}
type ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsInventoryItemsProfilesServiceProvider struct {
	Name  string                                                                                                      `json:"name,omitempty"`  // Service Provider name
	Plans *[]ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsInventoryItemsProfilesServiceProviderPlans `json:"plans,omitempty"` // Plans currently active on the eSIM
}
type ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsInventoryItemsProfilesServiceProviderPlans struct {
	Name string `json:"name,omitempty"` // Plan name
	Type string `json:"type,omitempty"` // Plan type (communication, rate)
}
type ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsInventoryMeta struct {
	Counts *ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsInventoryMetaCounts `json:"counts,omitempty"` // Counts of involved entities
}
type ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsInventoryMetaCounts struct {
	Items *ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsInventoryMetaCountsItems `json:"items,omitempty"` // Count of eSIM Devices available
}
type ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsInventoryMetaCountsItems struct {
	Remaining *int `json:"remaining,omitempty"` // Remaining number of eSIM Devices
	Total     *int `json:"total,omitempty"`     // Total number of eSIM Devices
}
type ResponseCellularGatewayUpdateOrganizationCellularGatewayEsimsInventory struct {
	Active        *bool                                                                             `json:"active,omitempty"`        // Whether eSIM is currently active SIM on Device
	Device        *ResponseCellularGatewayUpdateOrganizationCellularGatewayEsimsInventoryDevice     `json:"device,omitempty"`        // Meraki Device properties
	Eid           string                                                                            `json:"eid,omitempty"`           // eSIM EID
	LastUpdatedAt string                                                                            `json:"lastUpdatedAt,omitempty"` // Last update of eSIM
	Network       *ResponseCellularGatewayUpdateOrganizationCellularGatewayEsimsInventoryNetwork    `json:"network,omitempty"`       // Meraki Network properties
	Profiles      *[]ResponseCellularGatewayUpdateOrganizationCellularGatewayEsimsInventoryProfiles `json:"profiles,omitempty"`      // eSIM Profile Information
}
type ResponseCellularGatewayUpdateOrganizationCellularGatewayEsimsInventoryDevice struct {
	Model  string `json:"model,omitempty"`  // Device model
	Name   string `json:"name,omitempty"`   // Device name
	Serial string `json:"serial,omitempty"` // Device serial number
	Status string `json:"status,omitempty"` // Device status
	URL    string `json:"url,omitempty"`    // Device URL
}
type ResponseCellularGatewayUpdateOrganizationCellularGatewayEsimsInventoryNetwork struct {
	ID string `json:"id,omitempty"` // Network ID for this eSIM
}
type ResponseCellularGatewayUpdateOrganizationCellularGatewayEsimsInventoryProfiles struct {
	CustomApns      []string                                                                                       `json:"customApns,omitempty"`      // Available custom APNs for the profile
	Iccid           string                                                                                         `json:"iccid,omitempty"`           // eSIM profile ID
	ServiceProvider *ResponseCellularGatewayUpdateOrganizationCellularGatewayEsimsInventoryProfilesServiceProvider `json:"serviceProvider,omitempty"` // Service Provider information
	Status          string                                                                                         `json:"status,omitempty"`          // eSIM profile status
}
type ResponseCellularGatewayUpdateOrganizationCellularGatewayEsimsInventoryProfilesServiceProvider struct {
	Name  string                                                                                                `json:"name,omitempty"`  // Service Provider name
	Plans *[]ResponseCellularGatewayUpdateOrganizationCellularGatewayEsimsInventoryProfilesServiceProviderPlans `json:"plans,omitempty"` // Plans currently active on the eSIM
}
type ResponseCellularGatewayUpdateOrganizationCellularGatewayEsimsInventoryProfilesServiceProviderPlans struct {
	Name string `json:"name,omitempty"` // Plan name
	Type string `json:"type,omitempty"` // Plan type (communication, rate)
}
type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProviders struct {
	Items *[]ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersItems `json:"items,omitempty"` // List Cellular Service Providers
	Meta  *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersMeta    `json:"meta,omitempty"`  // Meta details about the result
}
type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersItems struct {
	IsBootstrap *bool                                                                                 `json:"isBootstrap,omitempty"` // Indicates if service provider is the bootstrap provider.
	Logo        *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersItemsLogo  `json:"logo,omitempty"`        // Service Provider logo data.
	Name        string                                                                                `json:"name,omitempty"`        // Service provider name.
	Terms       *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersItemsTerms `json:"terms,omitempty"`       // Service provider terms.
}
type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersItemsLogo struct {
	URL string `json:"url,omitempty"` // URL of service provider's logo.
}
type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersItemsTerms struct {
	Content string `json:"content,omitempty"` // URL of service provider's terms.
	Name    string `json:"name,omitempty"`    // Label for service provider's terms.
}
type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersMeta struct {
	Counts *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersMetaCounts `json:"counts,omitempty"` // Counts of involved entities
}
type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersMetaCounts struct {
	Items *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersMetaCountsItems `json:"items,omitempty"` // Service Providers available
}
type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersMetaCountsItems struct {
	Remaining *int `json:"remaining,omitempty"` // Remaining number of Service Providers
	Total     *int `json:"total,omitempty"`     // Total number of Service Providers
}
type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccounts ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccounts // Array of ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccounts
type ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccounts struct {
	Items *[]ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsItems `json:"items,omitempty"` // IList of Cellular Service Provider Accounts
	Meta  *ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsMeta    `json:"meta,omitempty"`  // Meta details about the result
}
type ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsItems struct {
	AccountID       string                                                                                                      `json:"accountId,omitempty"`       // Service provider account ID
	LastUpdatedAt   string                                                                                                      `json:"lastUpdatedAt,omitempty"`   // Last updated at
	ServiceProvider *ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsItemsServiceProvider `json:"serviceProvider,omitempty"` // Service provider data.
	Title           string                                                                                                      `json:"title,omitempty"`           // Service provider account name
	Username        string                                                                                                      `json:"username,omitempty"`        // Service provider account username
}
type ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsItemsServiceProvider struct {
	Logo *ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsItemsServiceProviderLogo `json:"logo,omitempty"` // Service provider logo data.
	Name string                                                                                                          `json:"name,omitempty"` // Name of the service provider.
}
type ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsItemsServiceProviderLogo struct {
	URL string `json:"url,omitempty"` // Service Provider logo url.
}
type ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsMeta struct {
	Counts *ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsMetaCounts `json:"counts,omitempty"` // Counts of involved entities
}
type ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsMetaCounts struct {
	Items *ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsMetaCountsItems `json:"items,omitempty"` // Count of Cellular Service Providers available
}
type ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsMetaCountsItems struct {
	Remaining *int `json:"remaining,omitempty"` // Remaining number of Cellular Service Providers
	Total     *int `json:"total,omitempty"`     // Total number of Cellular Service Providers
}
type ResponseCellularGatewayCreateOrganizationCellularGatewayEsimsServiceProvidersAccount struct {
	AccountID       string                                                                                               `json:"accountId,omitempty"`       // Service provider account ID
	LastUpdatedAt   string                                                                                               `json:"lastUpdatedAt,omitempty"`   // Last updated at
	ServiceProvider *ResponseCellularGatewayCreateOrganizationCellularGatewayEsimsServiceProvidersAccountServiceProvider `json:"serviceProvider,omitempty"` // Service provider data.
	Title           string                                                                                               `json:"title,omitempty"`           // Service provider account name
	Username        string                                                                                               `json:"username,omitempty"`        // Service provider account username
}
type ResponseCellularGatewayCreateOrganizationCellularGatewayEsimsServiceProvidersAccountServiceProvider struct {
	Logo *ResponseCellularGatewayCreateOrganizationCellularGatewayEsimsServiceProvidersAccountServiceProviderLogo `json:"logo,omitempty"` // Service provider logo data.
	Name string                                                                                                   `json:"name,omitempty"` // Name of the service provider.
}
type ResponseCellularGatewayCreateOrganizationCellularGatewayEsimsServiceProvidersAccountServiceProviderLogo struct {
	URL string `json:"url,omitempty"` // Service Provider logo url.
}
type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans struct {
	Items *[]ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansItems `json:"items,omitempty"` // List of Cellular Service Provider Communication Plans
	Meta  *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansMeta    `json:"meta,omitempty"`  // Meta details about the result
}
type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansItems struct {
	AccountID string                                                                                                           `json:"accountId,omitempty"` // Account ID of plans to be fetched
	Apns      *[]ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansItemsApns `json:"apns,omitempty"`      // Available APNs
	Name      string                                                                                                           `json:"name,omitempty"`      // Communication plan name
}
type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansItemsApns struct {
	Name string `json:"name,omitempty"` // APN name
}
type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansMeta struct {
	Counts *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansMetaCounts `json:"counts,omitempty"` // Counts of involved entities
}
type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansMetaCounts struct {
	Items *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansMetaCountsItems `json:"items,omitempty"` // Count of Communication Plans available
}
type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansMetaCountsItems struct {
	Remaining *int `json:"remaining,omitempty"` // Remaining number of Communication Plans
	Total     *int `json:"total,omitempty"`     // Total number of Communication Plans
}
type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans struct {
	Items *[]ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlansItems `json:"items,omitempty"` // List of Cellular Service Provider Rate Plans
	Meta  *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlansMeta    `json:"meta,omitempty"`  // Meta details about the result
}
type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlansItems struct {
	AccountID string `json:"accountId,omitempty"` // Account ID of plans to be fetched
	Name      string `json:"name,omitempty"`      // Rate plan name
}
type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlansMeta struct {
	Counts *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlansMetaCounts `json:"counts,omitempty"` // Counts of involved entities
}
type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlansMetaCounts struct {
	Items *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlansMetaCountsItems `json:"items,omitempty"` // Count of Rate Plans available
}
type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlansMetaCountsItems struct {
	Remaining *int `json:"remaining,omitempty"` // Remaining number of Rate Plans
	Total     *int `json:"total,omitempty"`     // Total number of Rate Plans
}
type ResponseCellularGatewayUpdateOrganizationCellularGatewayEsimsServiceProvidersAccount struct {
	AccountID       string                                                                                               `json:"accountId,omitempty"`       // Service provider account ID
	LastUpdatedAt   string                                                                                               `json:"lastUpdatedAt,omitempty"`   // Last updated at
	ServiceProvider *ResponseCellularGatewayUpdateOrganizationCellularGatewayEsimsServiceProvidersAccountServiceProvider `json:"serviceProvider,omitempty"` // Service provider data.
	Title           string                                                                                               `json:"title,omitempty"`           // Service provider account name
	Username        string                                                                                               `json:"username,omitempty"`        // Service provider account username
}
type ResponseCellularGatewayUpdateOrganizationCellularGatewayEsimsServiceProvidersAccountServiceProvider struct {
	Logo *ResponseCellularGatewayUpdateOrganizationCellularGatewayEsimsServiceProvidersAccountServiceProviderLogo `json:"logo,omitempty"` // Service provider logo data.
	Name string                                                                                                   `json:"name,omitempty"` // Name of the service provider.
}
type ResponseCellularGatewayUpdateOrganizationCellularGatewayEsimsServiceProvidersAccountServiceProviderLogo struct {
	URL string `json:"url,omitempty"` // Service Provider logo url.
}
type ResponseCellularGatewayCreateOrganizationCellularGatewayEsimsSwap struct {
	Eid    string `json:"eid,omitempty"`    // eSIM EID
	Iccid  string `json:"iccid,omitempty"`  // eSIM ICCID
	Status string `json:"status,omitempty"` // Swap status
}
type ResponseCellularGatewayUpdateOrganizationCellularGatewayEsimsSwap struct {
	Eid    string `json:"eid,omitempty"`    // eSIM EID
	Iccid  string `json:"iccid,omitempty"`  // eSIM ICCID
	Status string `json:"status,omitempty"` // Swap status
}
type ResponseCellularGatewayGetOrganizationCellularGatewayUplinkStatuses []ResponseItemCellularGatewayGetOrganizationCellularGatewayUplinkStatuses // Array of ResponseCellularGatewayGetOrganizationCellularGatewayUplinkStatuses
type ResponseItemCellularGatewayGetOrganizationCellularGatewayUplinkStatuses struct {
	LastReportedAt string                                                                            `json:"lastReportedAt,omitempty"` // Last reported time for the device
	Model          string                                                                            `json:"model,omitempty"`          // Device model
	NetworkID      string                                                                            `json:"networkId,omitempty"`      // Network Id
	Serial         string                                                                            `json:"serial,omitempty"`         // Serial number of the device
	Uplinks        *[]ResponseItemCellularGatewayGetOrganizationCellularGatewayUplinkStatusesUplinks `json:"uplinks,omitempty"`        // Uplinks info
}
type ResponseItemCellularGatewayGetOrganizationCellularGatewayUplinkStatusesUplinks struct {
	Apn            string                                                                                    `json:"apn,omitempty"`            // Access Point Name
	ConnectionType string                                                                                    `json:"connectionType,omitempty"` // Connection Type
	DNS1           string                                                                                    `json:"dns1,omitempty"`           // Primary DNS IP
	DNS2           string                                                                                    `json:"dns2,omitempty"`           // Secondary DNS IP
	Gateway        string                                                                                    `json:"gateway,omitempty"`        // Gateway IP
	Iccid          string                                                                                    `json:"iccid,omitempty"`          // Integrated Circuit Card Identification Number
	Interface      string                                                                                    `json:"interface,omitempty"`      // Uplink interface
	IP             string                                                                                    `json:"ip,omitempty"`             // Uplink IP
	Model          string                                                                                    `json:"model,omitempty"`          // Uplink model
	Provider       string                                                                                    `json:"provider,omitempty"`       // Network Provider
	PublicIP       string                                                                                    `json:"publicIp,omitempty"`       // Public IP
	SignalStat     *ResponseItemCellularGatewayGetOrganizationCellularGatewayUplinkStatusesUplinksSignalStat `json:"signalStat,omitempty"`     // Tower Signal Status
	SignalType     string                                                                                    `json:"signalType,omitempty"`     // Signal Type
	Status         string                                                                                    `json:"status,omitempty"`         // Uplink status
}
type ResponseItemCellularGatewayGetOrganizationCellularGatewayUplinkStatusesUplinksSignalStat struct {
	Rsrp string `json:"rsrp,omitempty"` // Reference Signal Received Power
	Rsrq string `json:"rsrq,omitempty"` // Reference Signal Received Quality
}
type RequestCellularGatewayUpdateDeviceCellularGatewayLan struct {
	FixedIPAssignments *[]RequestCellularGatewayUpdateDeviceCellularGatewayLanFixedIPAssignments `json:"fixedIpAssignments,omitempty"` // list of all fixed IP assignments for a single MG
	ReservedIPRanges   *[]RequestCellularGatewayUpdateDeviceCellularGatewayLanReservedIPRanges   `json:"reservedIpRanges,omitempty"`   // list of all reserved IP ranges for a single MG
}
type RequestCellularGatewayUpdateDeviceCellularGatewayLanFixedIPAssignments struct {
	IP   string `json:"ip,omitempty"`   // The IP address you want to assign to a specific server or device
	Mac  string `json:"mac,omitempty"`  // The MAC address of the server or device that hosts the internal resource that you wish to receive the specified IP address
	Name string `json:"name,omitempty"` // A descriptive name of the assignment
}
type RequestCellularGatewayUpdateDeviceCellularGatewayLanReservedIPRanges struct {
	Comment string `json:"comment,omitempty"` // Comment explaining the reserved IP range
	End     string `json:"end,omitempty"`     // Ending IP included in the reserved range of IPs
	Start   string `json:"start,omitempty"`   // Starting IP included in the reserved range of IPs
}
type RequestCellularGatewayUpdateDeviceCellularGatewayPortForwardingRules struct {
	Rules *[]RequestCellularGatewayUpdateDeviceCellularGatewayPortForwardingRulesRules `json:"rules,omitempty"` // An array of port forwarding params
}
type RequestCellularGatewayUpdateDeviceCellularGatewayPortForwardingRulesRules struct {
	Access     string   `json:"access,omitempty"`     // *any* or *restricted*. Specify the right to make inbound connections on the specified ports or port ranges. If *restricted*, a list of allowed IPs is mandatory.
	AllowedIPs []string `json:"allowedIps,omitempty"` // An array of ranges of WAN IP addresses that are allowed to make inbound connections on the specified ports or port ranges.
	LanIP      string   `json:"lanIp,omitempty"`      // The IP address of the server or device that hosts the internal resource that you wish to make available on the WAN
	LocalPort  string   `json:"localPort,omitempty"`  // A port or port ranges that will receive the forwarded traffic from the WAN
	Name       string   `json:"name,omitempty"`       // A descriptive name for the rule
	Protocol   string   `json:"protocol,omitempty"`   // TCP or UDP
	PublicPort string   `json:"publicPort,omitempty"` // A port or port ranges that will be forwarded to the host on the LAN
}
type RequestCellularGatewayUpdateNetworkCellularGatewayConnectivityMonitoringDestinations struct {
	Destinations *[]RequestCellularGatewayUpdateNetworkCellularGatewayConnectivityMonitoringDestinationsDestinations `json:"destinations,omitempty"` // The list of connectivity monitoring destinations
}
type RequestCellularGatewayUpdateNetworkCellularGatewayConnectivityMonitoringDestinationsDestinations struct {
	Default     *bool  `json:"default,omitempty"`     // Boolean indicating whether this is the default testing destination (true) or not (false). Defaults to false. Only one default is allowed
	Description string `json:"description,omitempty"` // Description of the testing destination. Optional, defaults to an empty string
	IP          string `json:"ip,omitempty"`          // The IP address to test connectivity with
}
type RequestCellularGatewayUpdateNetworkCellularGatewayDhcp struct {
	DhcpLeaseTime        string   `json:"dhcpLeaseTime,omitempty"`        // DHCP Lease time for all MG of the network. Possible values are '30 minutes', '1 hour', '4 hours', '12 hours', '1 day' or '1 week'.
	DNSCustomNameservers []string `json:"dnsCustomNameservers,omitempty"` // list of fixed IPs representing the the DNS Name servers when the mode is 'custom'
	DNSNameservers       string   `json:"dnsNameservers,omitempty"`       // DNS name servers mode for all MG of the network. Possible values are: 'upstream_dns', 'google_dns', 'opendns', 'custom'.
}
type RequestCellularGatewayUpdateNetworkCellularGatewaySubnetPool struct {
	Cidr string `json:"cidr,omitempty"` // CIDR of the pool of subnets. Each MG in this network will automatically pick a subnet from this pool.
	Mask *int   `json:"mask,omitempty"` // Mask used for the subnet of all MGs in  this network.
}
type RequestCellularGatewayUpdateNetworkCellularGatewayUplink struct {
	BandwidthLimits *RequestCellularGatewayUpdateNetworkCellularGatewayUplinkBandwidthLimits `json:"bandwidthLimits,omitempty"` // The bandwidth settings for the 'cellular' uplink
}
type RequestCellularGatewayUpdateNetworkCellularGatewayUplinkBandwidthLimits struct {
	LimitDown *int `json:"limitDown,omitempty"` // The maximum download limit (integer, in Kbps). null indicates no limit
	LimitUp   *int `json:"limitUp,omitempty"`   // The maximum upload limit (integer, in Kbps). null indicates no limit
}
type RequestCellularGatewayUpdateOrganizationCellularGatewayEsimsInventory struct {
	Status string `json:"status,omitempty"` // Status the eSIM will be updated to
}
type RequestCellularGatewayCreateOrganizationCellularGatewayEsimsServiceProvidersAccount struct {
	AccountID       string                                                                                              `json:"accountId,omitempty"`       // Service provider account ID
	APIKey          string                                                                                              `json:"apiKey,omitempty"`          // Service provider account API key
	ServiceProvider *RequestCellularGatewayCreateOrganizationCellularGatewayEsimsServiceProvidersAccountServiceProvider `json:"serviceProvider,omitempty"` // Service Provider information
	Title           string                                                                                              `json:"title,omitempty"`           // Service provider account name
	Username        string                                                                                              `json:"username,omitempty"`        // Service provider account username
}
type RequestCellularGatewayCreateOrganizationCellularGatewayEsimsServiceProvidersAccountServiceProvider struct {
	Name string `json:"name,omitempty"` // Service provider name
}
type RequestCellularGatewayUpdateOrganizationCellularGatewayEsimsServiceProvidersAccount struct {
	APIKey string `json:"apiKey,omitempty"` // Service provider account API key
	Title  string `json:"title,omitempty"`  // Service provider account name used on the Meraki UI
}
type RequestCellularGatewayCreateOrganizationCellularGatewayEsimsSwap struct {
	Swaps *[]RequestCellularGatewayCreateOrganizationCellularGatewayEsimsSwapSwaps `json:"swaps,omitempty"` // Each object represents a swap for one eSIM
}
type RequestCellularGatewayCreateOrganizationCellularGatewayEsimsSwapSwaps struct {
	Eid    string                                                                       `json:"eid,omitempty"`    // eSIM EID
	Target *RequestCellularGatewayCreateOrganizationCellularGatewayEsimsSwapSwapsTarget `json:"target,omitempty"` // Target Profile attributes
}
type RequestCellularGatewayCreateOrganizationCellularGatewayEsimsSwapSwapsTarget struct {
	AccountID         string `json:"accountId,omitempty"`         // ID of the target account; can be the account currently tied to the eSIM
	CommunicationPlan string `json:"communicationPlan,omitempty"` // Name of the target communication plan
	RatePlan          string `json:"ratePlan,omitempty"`          // Name of the target rate plan
}

//GetDeviceCellularGatewayLan Show the LAN Settings of a MG
/* Show the LAN Settings of a MG

@param serial serial path parameter.


*/

func (s *CellularGatewayService) GetDeviceCellularGatewayLan(serial string) (*ResponseCellularGatewayGetDeviceCellularGatewayLan, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/cellularGateway/lan"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCellularGatewayGetDeviceCellularGatewayLan{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceCellularGatewayLan")
	}

	result := response.Result().(*ResponseCellularGatewayGetDeviceCellularGatewayLan)
	return result, response, err

}

//GetDeviceCellularGatewayPortForwardingRules Returns the port forwarding rules for a single MG.
/* Returns the port forwarding rules for a single MG.

@param serial serial path parameter.


*/

func (s *CellularGatewayService) GetDeviceCellularGatewayPortForwardingRules(serial string) (*ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRules, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/cellularGateway/portForwardingRules"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRules{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceCellularGatewayPortForwardingRules")
	}

	result := response.Result().(*ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRules)
	return result, response, err

}

//GetNetworkCellularGatewayConnectivityMonitoringDestinations Return the connectivity testing destinations for an MG network
/* Return the connectivity testing destinations for an MG network

@param networkID networkId path parameter. Network ID


*/

func (s *CellularGatewayService) GetNetworkCellularGatewayConnectivityMonitoringDestinations(networkID string) (*ResponseCellularGatewayGetNetworkCellularGatewayConnectivityMonitoringDestinations, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/cellularGateway/connectivityMonitoringDestinations"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCellularGatewayGetNetworkCellularGatewayConnectivityMonitoringDestinations{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkCellularGatewayConnectivityMonitoringDestinations")
	}

	result := response.Result().(*ResponseCellularGatewayGetNetworkCellularGatewayConnectivityMonitoringDestinations)
	return result, response, err

}

//GetNetworkCellularGatewayDhcp List common DHCP settings of MGs
/* List common DHCP settings of MGs

@param networkID networkId path parameter. Network ID


*/

func (s *CellularGatewayService) GetNetworkCellularGatewayDhcp(networkID string) (*ResponseCellularGatewayGetNetworkCellularGatewayDhcp, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/cellularGateway/dhcp"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCellularGatewayGetNetworkCellularGatewayDhcp{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkCellularGatewayDhcp")
	}

	result := response.Result().(*ResponseCellularGatewayGetNetworkCellularGatewayDhcp)
	return result, response, err

}

//GetNetworkCellularGatewaySubnetPool Return the subnet pool and mask configured for MGs in the network.
/* Return the subnet pool and mask configured for MGs in the network.

@param networkID networkId path parameter. Network ID


*/

func (s *CellularGatewayService) GetNetworkCellularGatewaySubnetPool(networkID string) (*ResponseCellularGatewayGetNetworkCellularGatewaySubnetPool, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/cellularGateway/subnetPool"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCellularGatewayGetNetworkCellularGatewaySubnetPool{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkCellularGatewaySubnetPool")
	}

	result := response.Result().(*ResponseCellularGatewayGetNetworkCellularGatewaySubnetPool)
	return result, response, err

}

//GetNetworkCellularGatewayUplink Returns the uplink settings for your MG network.
/* Returns the uplink settings for your MG network.

@param networkID networkId path parameter. Network ID


*/

func (s *CellularGatewayService) GetNetworkCellularGatewayUplink(networkID string) (*ResponseCellularGatewayGetNetworkCellularGatewayUplink, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/cellularGateway/uplink"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCellularGatewayGetNetworkCellularGatewayUplink{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkCellularGatewayUplink")
	}

	result := response.Result().(*ResponseCellularGatewayGetNetworkCellularGatewayUplink)
	return result, response, err

}

//GetOrganizationCellularGatewayEsimsInventory The eSIM inventory of a given organization.
/* The eSIM inventory of a given organization.

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationCellularGatewayEsimsInventoryQueryParams Filtering parameter


*/

func (s *CellularGatewayService) GetOrganizationCellularGatewayEsimsInventory(organizationID string, getOrganizationCellularGatewayEsimsInventoryQueryParams *GetOrganizationCellularGatewayEsimsInventoryQueryParams) (*ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventory, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/cellularGateway/esims/inventory"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationCellularGatewayEsimsInventoryQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventory{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationCellularGatewayEsimsInventory")
	}

	result := response.Result().(*ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventory)
	return result, response, err

}

//GetOrganizationCellularGatewayEsimsServiceProviders Service providers customers can add accounts for.
/* Service providers customers can add accounts for.

@param organizationID organizationId path parameter. Organization ID


*/

func (s *CellularGatewayService) GetOrganizationCellularGatewayEsimsServiceProviders(organizationID string) (*ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProviders, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/cellularGateway/esims/serviceProviders"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProviders{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationCellularGatewayEsimsServiceProviders")
	}

	result := response.Result().(*ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProviders)
	return result, response, err

}

//GetOrganizationCellularGatewayEsimsServiceProvidersAccounts Inventory of service provider accounts tied to the organization.
/* Inventory of service provider accounts tied to the organization.

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationCellularGatewayEsimsServiceProvidersAccountsQueryParams Filtering parameter


*/

func (s *CellularGatewayService) GetOrganizationCellularGatewayEsimsServiceProvidersAccounts(organizationID string, getOrganizationCellularGatewayEsimsServiceProvidersAccountsQueryParams *GetOrganizationCellularGatewayEsimsServiceProvidersAccountsQueryParams) (*ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccounts, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/cellularGateway/esims/serviceProviders/accounts"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationCellularGatewayEsimsServiceProvidersAccountsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccounts{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationCellularGatewayEsimsServiceProvidersAccounts")
	}

	result := response.Result().(*ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccounts)
	return result, response, err

}

//GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans The communication plans available for a given provider.
/* The communication plans available for a given provider.

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansQueryParams Filtering parameter


*/

func (s *CellularGatewayService) GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans(organizationID string, getOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansQueryParams *GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansQueryParams) (*ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/cellularGateway/esims/serviceProviders/accounts/communicationPlans"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans")
	}

	result := response.Result().(*ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans)
	return result, response, err

}

//GetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans The rate plans available for a given provider.
/* The rate plans available for a given provider.

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlansQueryParams Filtering parameter


*/

func (s *CellularGatewayService) GetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans(organizationID string, getOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlansQueryParams *GetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlansQueryParams) (*ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/cellularGateway/esims/serviceProviders/accounts/ratePlans"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlansQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans")
	}

	result := response.Result().(*ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans)
	return result, response, err

}

//GetOrganizationCellularGatewayUplinkStatuses List the uplink status of every Meraki MG cellular gateway in the organization
/* List the uplink status of every Meraki MG cellular gateway in the organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationCellularGatewayUplinkStatusesQueryParams Filtering parameter


*/

func (s *CellularGatewayService) GetOrganizationCellularGatewayUplinkStatuses(organizationID string, getOrganizationCellularGatewayUplinkStatusesQueryParams *GetOrganizationCellularGatewayUplinkStatusesQueryParams) (*ResponseCellularGatewayGetOrganizationCellularGatewayUplinkStatuses, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/cellularGateway/uplink/statuses"
	s.rateLimiterBucket.Wait(1)

	if getOrganizationCellularGatewayUplinkStatusesQueryParams != nil && getOrganizationCellularGatewayUplinkStatusesQueryParams.PerPage == -1 {
		var result *ResponseCellularGatewayGetOrganizationCellularGatewayUplinkStatuses
		println("Paginate")
		getOrganizationCellularGatewayUplinkStatusesQueryParams.PerPage = PAGINATION_PER_PAGE
		result2, response, err := Paginate(s.GetOrganizationCellularGatewayUplinkStatusesPaginate, organizationID, "", getOrganizationCellularGatewayUplinkStatusesQueryParams)
		if err != nil {
			return nil, nil, err
		}
		jsonResult, err := json.Marshal(result2)
		// Verficar el error
		if err != nil {
			return nil, nil, err
		}
		var paginatedResponse []any
		err = json.Unmarshal(jsonResult, &paginatedResponse)
		// for para recorrer "paginatedResponse"
		for i := 0; i < len(paginatedResponse); i++ {
			var resultTmp *ResponseCellularGatewayGetOrganizationCellularGatewayUplinkStatuses
			jsonResult2, _ := json.Marshal(paginatedResponse[i])
			err = json.Unmarshal(jsonResult2, &resultTmp)
			// Verificar si result es nil, si lo es inicialiarlo
			if result == nil {
				result = resultTmp
			} else {
				*result = append(*result, *resultTmp...)
			}
		}
		return result, response, err
	}
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationCellularGatewayUplinkStatusesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseCellularGatewayGetOrganizationCellularGatewayUplinkStatuses{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationCellularGatewayUplinkStatuses")
	}

	result := response.Result().(*ResponseCellularGatewayGetOrganizationCellularGatewayUplinkStatuses)
	return result, response, err

}
func (s *CellularGatewayService) GetOrganizationCellularGatewayUplinkStatusesPaginate(organizationID string, getOrganizationCellularGatewayUplinkStatusesQueryParams any) (any, *resty.Response, error) {
	getOrganizationCellularGatewayUplinkStatusesQueryParamsConverted := getOrganizationCellularGatewayUplinkStatusesQueryParams.(*GetOrganizationCellularGatewayUplinkStatusesQueryParams)

	return s.GetOrganizationCellularGatewayUplinkStatuses(organizationID, getOrganizationCellularGatewayUplinkStatusesQueryParamsConverted)
}

//CreateOrganizationCellularGatewayEsimsServiceProvidersAccount Add a service provider account.
/* Add a service provider account.

@param organizationID organizationId path parameter. Organization ID


*/

func (s *CellularGatewayService) CreateOrganizationCellularGatewayEsimsServiceProvidersAccount(organizationID string, requestCellularGatewayCreateOrganizationCellularGatewayEsimsServiceProvidersAccount *RequestCellularGatewayCreateOrganizationCellularGatewayEsimsServiceProvidersAccount) (*ResponseCellularGatewayCreateOrganizationCellularGatewayEsimsServiceProvidersAccount, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/cellularGateway/esims/serviceProviders/accounts"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCellularGatewayCreateOrganizationCellularGatewayEsimsServiceProvidersAccount).
		SetResult(&ResponseCellularGatewayCreateOrganizationCellularGatewayEsimsServiceProvidersAccount{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateOrganizationCellularGatewayEsimsServiceProvidersAccount")
	}

	result := response.Result().(*ResponseCellularGatewayCreateOrganizationCellularGatewayEsimsServiceProvidersAccount)
	return result, response, err

}

//CreateOrganizationCellularGatewayEsimsSwap Swap which profile an eSIM uses.
/* Swap which profile an eSIM uses.

@param organizationID organizationId path parameter. Organization ID


*/

func (s *CellularGatewayService) CreateOrganizationCellularGatewayEsimsSwap(organizationID string, requestCellularGatewayCreateOrganizationCellularGatewayEsimsSwap *RequestCellularGatewayCreateOrganizationCellularGatewayEsimsSwap) (*ResponseCellularGatewayCreateOrganizationCellularGatewayEsimsSwap, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/cellularGateway/esims/swap"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCellularGatewayCreateOrganizationCellularGatewayEsimsSwap).
		SetResult(&ResponseCellularGatewayCreateOrganizationCellularGatewayEsimsSwap{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateOrganizationCellularGatewayEsimsSwap")
	}

	result := response.Result().(*ResponseCellularGatewayCreateOrganizationCellularGatewayEsimsSwap)
	return result, response, err

}

//UpdateDeviceCellularGatewayLan Update the LAN Settings for a single MG.
/* Update the LAN Settings for a single MG.

@param serial serial path parameter.
*/
func (s *CellularGatewayService) UpdateDeviceCellularGatewayLan(serial string, requestCellularGatewayUpdateDeviceCellularGatewayLan *RequestCellularGatewayUpdateDeviceCellularGatewayLan) (*ResponseCellularGatewayUpdateDeviceCellularGatewayLan, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/cellularGateway/lan"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCellularGatewayUpdateDeviceCellularGatewayLan).
		SetResult(&ResponseCellularGatewayUpdateDeviceCellularGatewayLan{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateDeviceCellularGatewayLan")
	}

	result := response.Result().(*ResponseCellularGatewayUpdateDeviceCellularGatewayLan)
	return result, response, err

}

//UpdateDeviceCellularGatewayPortForwardingRules Updates the port forwarding rules for a single MG.
/* Updates the port forwarding rules for a single MG.

@param serial serial path parameter.
*/
func (s *CellularGatewayService) UpdateDeviceCellularGatewayPortForwardingRules(serial string, requestCellularGatewayUpdateDeviceCellularGatewayPortForwardingRules *RequestCellularGatewayUpdateDeviceCellularGatewayPortForwardingRules) (*ResponseCellularGatewayUpdateDeviceCellularGatewayPortForwardingRules, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/cellularGateway/portForwardingRules"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCellularGatewayUpdateDeviceCellularGatewayPortForwardingRules).
		SetResult(&ResponseCellularGatewayUpdateDeviceCellularGatewayPortForwardingRules{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateDeviceCellularGatewayPortForwardingRules")
	}

	result := response.Result().(*ResponseCellularGatewayUpdateDeviceCellularGatewayPortForwardingRules)
	return result, response, err

}

//UpdateNetworkCellularGatewayConnectivityMonitoringDestinations Update the connectivity testing destinations for an MG network
/* Update the connectivity testing destinations for an MG network

@param networkID networkId path parameter. Network ID
*/
func (s *CellularGatewayService) UpdateNetworkCellularGatewayConnectivityMonitoringDestinations(networkID string, requestCellularGatewayUpdateNetworkCellularGatewayConnectivityMonitoringDestinations *RequestCellularGatewayUpdateNetworkCellularGatewayConnectivityMonitoringDestinations) (*ResponseCellularGatewayUpdateNetworkCellularGatewayConnectivityMonitoringDestinations, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/cellularGateway/connectivityMonitoringDestinations"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCellularGatewayUpdateNetworkCellularGatewayConnectivityMonitoringDestinations).
		SetResult(&ResponseCellularGatewayUpdateNetworkCellularGatewayConnectivityMonitoringDestinations{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkCellularGatewayConnectivityMonitoringDestinations")
	}

	result := response.Result().(*ResponseCellularGatewayUpdateNetworkCellularGatewayConnectivityMonitoringDestinations)
	return result, response, err

}

//UpdateNetworkCellularGatewayDhcp Update common DHCP settings of MGs
/* Update common DHCP settings of MGs

@param networkID networkId path parameter. Network ID
*/
func (s *CellularGatewayService) UpdateNetworkCellularGatewayDhcp(networkID string, requestCellularGatewayUpdateNetworkCellularGatewayDhcp *RequestCellularGatewayUpdateNetworkCellularGatewayDhcp) (*ResponseCellularGatewayUpdateNetworkCellularGatewayDhcp, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/cellularGateway/dhcp"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCellularGatewayUpdateNetworkCellularGatewayDhcp).
		SetResult(&ResponseCellularGatewayUpdateNetworkCellularGatewayDhcp{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkCellularGatewayDhcp")
	}

	result := response.Result().(*ResponseCellularGatewayUpdateNetworkCellularGatewayDhcp)
	return result, response, err

}

//UpdateNetworkCellularGatewaySubnetPool Update the subnet pool and mask configuration for MGs in the network.
/* Update the subnet pool and mask configuration for MGs in the network.

@param networkID networkId path parameter. Network ID
*/
func (s *CellularGatewayService) UpdateNetworkCellularGatewaySubnetPool(networkID string, requestCellularGatewayUpdateNetworkCellularGatewaySubnetPool *RequestCellularGatewayUpdateNetworkCellularGatewaySubnetPool) (*ResponseCellularGatewayUpdateNetworkCellularGatewaySubnetPool, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/cellularGateway/subnetPool"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCellularGatewayUpdateNetworkCellularGatewaySubnetPool).
		SetResult(&ResponseCellularGatewayUpdateNetworkCellularGatewaySubnetPool{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkCellularGatewaySubnetPool")
	}

	result := response.Result().(*ResponseCellularGatewayUpdateNetworkCellularGatewaySubnetPool)
	return result, response, err

}

//UpdateNetworkCellularGatewayUplink Updates the uplink settings for your MG network.
/* Updates the uplink settings for your MG network.

@param networkID networkId path parameter. Network ID
*/
func (s *CellularGatewayService) UpdateNetworkCellularGatewayUplink(networkID string, requestCellularGatewayUpdateNetworkCellularGatewayUplink *RequestCellularGatewayUpdateNetworkCellularGatewayUplink) (*ResponseCellularGatewayUpdateNetworkCellularGatewayUplink, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/cellularGateway/uplink"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCellularGatewayUpdateNetworkCellularGatewayUplink).
		SetResult(&ResponseCellularGatewayUpdateNetworkCellularGatewayUplink{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkCellularGatewayUplink")
	}

	result := response.Result().(*ResponseCellularGatewayUpdateNetworkCellularGatewayUplink)
	return result, response, err

}

//UpdateOrganizationCellularGatewayEsimsInventory Toggle the status of an eSIM
/* Toggle the status of an eSIM

@param organizationID organizationId path parameter. Organization ID
@param id id path parameter.
*/
func (s *CellularGatewayService) UpdateOrganizationCellularGatewayEsimsInventory(organizationID string, id string, requestCellularGatewayUpdateOrganizationCellularGatewayEsimsInventory *RequestCellularGatewayUpdateOrganizationCellularGatewayEsimsInventory) (*ResponseCellularGatewayUpdateOrganizationCellularGatewayEsimsInventory, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/cellularGateway/esims/inventory/{id}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCellularGatewayUpdateOrganizationCellularGatewayEsimsInventory).
		SetResult(&ResponseCellularGatewayUpdateOrganizationCellularGatewayEsimsInventory{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateOrganizationCellularGatewayEsimsInventory")
	}

	result := response.Result().(*ResponseCellularGatewayUpdateOrganizationCellularGatewayEsimsInventory)
	return result, response, err

}

//UpdateOrganizationCellularGatewayEsimsServiceProvidersAccount Edit service provider account info stored in Meraki's database.
/* Edit service provider account info stored in Meraki's database.

@param organizationID organizationId path parameter. Organization ID
@param accountID accountId path parameter. Account ID
*/
func (s *CellularGatewayService) UpdateOrganizationCellularGatewayEsimsServiceProvidersAccount(organizationID string, accountID string, requestCellularGatewayUpdateOrganizationCellularGatewayEsimsServiceProvidersAccount *RequestCellularGatewayUpdateOrganizationCellularGatewayEsimsServiceProvidersAccount) (*ResponseCellularGatewayUpdateOrganizationCellularGatewayEsimsServiceProvidersAccount, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/cellularGateway/esims/serviceProviders/accounts/{accountId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{accountId}", fmt.Sprintf("%v", accountID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCellularGatewayUpdateOrganizationCellularGatewayEsimsServiceProvidersAccount).
		SetResult(&ResponseCellularGatewayUpdateOrganizationCellularGatewayEsimsServiceProvidersAccount{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateOrganizationCellularGatewayEsimsServiceProvidersAccount")
	}

	result := response.Result().(*ResponseCellularGatewayUpdateOrganizationCellularGatewayEsimsServiceProvidersAccount)
	return result, response, err

}

//UpdateOrganizationCellularGatewayEsimsSwap Get the status of a profile swap.
/* Get the status of a profile swap.

@param id id path parameter. eSIM EID
@param organizationID organizationId path parameter. Organization ID
*/
func (s *CellularGatewayService) UpdateOrganizationCellularGatewayEsimsSwap(id string, organizationID string) (*ResponseCellularGatewayUpdateOrganizationCellularGatewayEsimsSwap, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/cellularGateway/esims/swap/{id}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCellularGatewayUpdateOrganizationCellularGatewayEsimsSwap{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateOrganizationCellularGatewayEsimsSwap")
	}

	result := response.Result().(*ResponseCellularGatewayUpdateOrganizationCellularGatewayEsimsSwap)
	return result, response, err

}

//DeleteOrganizationCellularGatewayEsimsServiceProvidersAccount Remove a service provider account's integration with the Dashboard.
/* Remove a service provider account's integration with the Dashboard.

@param organizationID organizationId path parameter. Organization ID
@param accountID accountId path parameter. Account ID


*/
func (s *CellularGatewayService) DeleteOrganizationCellularGatewayEsimsServiceProvidersAccount(organizationID string, accountID string) (*resty.Response, error) {
	//organizationID string,accountID string
	path := "/api/v1/organizations/{organizationId}/cellularGateway/esims/serviceProviders/accounts/{accountId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{accountId}", fmt.Sprintf("%v", accountID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteOrganizationCellularGatewayEsimsServiceProvidersAccount")
	}

	return response, err

}
