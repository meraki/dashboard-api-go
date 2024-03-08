package meraki

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type NetworksService service

type GetNetworkAlertsHistoryQueryParams struct {
	PerPage       int    `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100.
	StartingAfter string `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
}
type GetNetworkBluetoothClientsQueryParams struct {
	T0                         string  `url:"t0,omitempty"`                         //The beginning of the timespan for the data. The maximum lookback period is 7 days from today.
	Timespan                   float64 `url:"timespan,omitempty"`                   //The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 7 days. The default is 1 day.
	PerPage                    int     `url:"perPage,omitempty"`                    //The number of entries per page returned. Acceptable range is 5 - 1000. Default is 10.
	StartingAfter              string  `url:"startingAfter,omitempty"`              //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore               string  `url:"endingBefore,omitempty"`               //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	IncludeConnectivityHistory bool    `url:"includeConnectivityHistory,omitempty"` //Include the connectivity history for this client
}
type GetNetworkBluetoothClientQueryParams struct {
	IncludeConnectivityHistory  bool `url:"includeConnectivityHistory,omitempty"`  //Include the connectivity history for this client
	ConnectivityHistoryTimespan int  `url:"connectivityHistoryTimespan,omitempty"` //The timespan, in seconds, for the connectivityHistory data. By default 1 day, 86400, will be used.
}
type GetNetworkClientsQueryParams struct {
	T0                      string   `url:"t0,omitempty"`                        //The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	Timespan                float64  `url:"timespan,omitempty"`                  //The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
	PerPage                 int      `url:"perPage,omitempty"`                   //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 10.
	StartingAfter           string   `url:"startingAfter,omitempty"`             //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore            string   `url:"endingBefore,omitempty"`              //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	Statuses                []string `url:"statuses[],omitempty"`                //Filters clients based on status. Can be one of 'Online' or 'Offline'.
	IP                      string   `url:"ip,omitempty"`                        //Filters clients based on a partial or full match for the ip address field.
	IP6                     string   `url:"ip6,omitempty"`                       //Filters clients based on a partial or full match for the ip6 address field.
	IP6Local                string   `url:"ip6Local,omitempty"`                  //Filters clients based on a partial or full match for the ip6Local address field.
	Mac                     string   `url:"mac,omitempty"`                       //Filters clients based on a partial or full match for the mac address field.
	Os                      string   `url:"os,omitempty"`                        //Filters clients based on a partial or full match for the os (operating system) field.
	PskGroup                string   `url:"pskGroup,omitempty"`                  //Filters clients based on partial or full match for the iPSK name field.
	Description             string   `url:"description,omitempty"`               //Filters clients based on a partial or full match for the description field.
	VLAN                    string   `url:"vlan,omitempty"`                      //Filters clients based on the full match for the VLAN field.
	RecentDeviceConnections []string `url:"recentDeviceConnections[],omitempty"` //Filters clients based on recent connection type. Can be one of 'Wired' or 'Wireless'.
}
type GetNetworkClientsApplicationUsageQueryParams struct {
	Clients       string  `url:"clients,omitempty"`       //A list of client keys, MACs or IPs separated by comma.
	SSIDNumber    int     `url:"ssidNumber,omitempty"`    //An SSID number to include. If not specified, eveusage histories application usagents for all SSIDs will be returned.
	PerPage       int     `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000.
	StartingAfter string  `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string  `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	T0            string  `url:"t0,omitempty"`            //The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	T1            string  `url:"t1,omitempty"`            //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan      float64 `url:"timespan,omitempty"`      //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
}
type GetNetworkClientsBandwidthUsageHistoryQueryParams struct {
	T0            string  `url:"t0,omitempty"`            //The beginning of the timespan for the data. The maximum lookback period is 30 days from today.
	T1            string  `url:"t1,omitempty"`            //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan      float64 `url:"timespan,omitempty"`      //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
	PerPage       int     `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter string  `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string  `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
}
type GetNetworkClientsOverviewQueryParams struct {
	T0         string  `url:"t0,omitempty"`         //The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	T1         string  `url:"t1,omitempty"`         //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan   float64 `url:"timespan,omitempty"`   //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
	Resolution int     `url:"resolution,omitempty"` //The time resolution in seconds for returned data. The valid resolutions are: 7200, 86400, 604800, 2592000. The default is 604800.
}
type GetNetworkClientsUsageHistoriesQueryParams struct {
	Clients       string  `url:"clients,omitempty"`       //A list of client keys, MACs or IPs separated by comma.
	SSIDNumber    int     `url:"ssidNumber,omitempty"`    //An SSID number to include. If not specified, events for all SSIDs will be returned.
	PerPage       int     `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000.
	StartingAfter string  `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string  `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	T0            string  `url:"t0,omitempty"`            //The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	T1            string  `url:"t1,omitempty"`            //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan      float64 `url:"timespan,omitempty"`      //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
}
type GetNetworkClientTrafficHistoryQueryParams struct {
	PerPage       int    `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000.
	StartingAfter string `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
}
type GetNetworkEventsQueryParams struct {
	ProductType        string   `url:"productType,omitempty"`          //The product type to fetch events for. This parameter is required for networks with multiple device types. Valid types are wireless, appliance, switch, systemsManager, camera, and cellularGateway
	IncludedEventTypes []string `url:"includedEventTypes[],omitempty"` //A list of event types. The returned events will be filtered to only include events with these types.
	ExcludedEventTypes []string `url:"excludedEventTypes[],omitempty"` //A list of event types. The returned events will be filtered to exclude events with these types.
	DeviceMac          string   `url:"deviceMac,omitempty"`            //The MAC address of the Meraki device which the list of events will be filtered with
	DeviceSerial       string   `url:"deviceSerial,omitempty"`         //The serial of the Meraki device which the list of events will be filtered with
	DeviceName         string   `url:"deviceName,omitempty"`           //The name of the Meraki device which the list of events will be filtered with
	ClientIP           string   `url:"clientIp,omitempty"`             //The IP of the client which the list of events will be filtered with. Only supported for track-by-IP networks.
	ClientMac          string   `url:"clientMac,omitempty"`            //The MAC address of the client which the list of events will be filtered with. Only supported for track-by-MAC networks.
	ClientName         string   `url:"clientName,omitempty"`           //The name, or partial name, of the client which the list of events will be filtered with
	SmDeviceMac        string   `url:"smDeviceMac,omitempty"`          //The MAC address of the Systems Manager device which the list of events will be filtered with
	SmDeviceName       string   `url:"smDeviceName,omitempty"`         //The name of the Systems Manager device which the list of events will be filtered with
	PerPage            int      `url:"perPage,omitempty"`              //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 10.
	StartingAfter      string   `url:"startingAfter,omitempty"`        //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore       string   `url:"endingBefore,omitempty"`         //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
}
type GetNetworkNetworkHealthChannelUtilizationQueryParams struct {
	T0            string  `url:"t0,omitempty"`            //The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	T1            string  `url:"t1,omitempty"`            //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan      float64 `url:"timespan,omitempty"`      //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
	Resolution    int     `url:"resolution,omitempty"`    //The time resolution in seconds for returned data. The valid resolutions are: 600. The default is 600.
	PerPage       int     `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 100. Default is 10.
	StartingAfter string  `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string  `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
}
type GetNetworkPiiPiiKeysQueryParams struct {
	Username     string `url:"username,omitempty"`     //The username of a Systems Manager user
	Email        string `url:"email,omitempty"`        //The email of a network user account or a Systems Manager device
	Mac          string `url:"mac,omitempty"`          //The MAC of a network client device or a Systems Manager device
	Serial       string `url:"serial,omitempty"`       //The serial of a Systems Manager device
	Imei         string `url:"imei,omitempty"`         //The IMEI of a Systems Manager device
	BluetoothMac string `url:"bluetoothMac,omitempty"` //The MAC of a Bluetooth client
}
type GetNetworkPiiSmDevicesForKeyQueryParams struct {
	Username     string `url:"username,omitempty"`     //The username of a Systems Manager user
	Email        string `url:"email,omitempty"`        //The email of a network user account or a Systems Manager device
	Mac          string `url:"mac,omitempty"`          //The MAC of a network client device or a Systems Manager device
	Serial       string `url:"serial,omitempty"`       //The serial of a Systems Manager device
	Imei         string `url:"imei,omitempty"`         //The IMEI of a Systems Manager device
	BluetoothMac string `url:"bluetoothMac,omitempty"` //The MAC of a Bluetooth client
}
type GetNetworkPiiSmOwnersForKeyQueryParams struct {
	Username     string `url:"username,omitempty"`     //The username of a Systems Manager user
	Email        string `url:"email,omitempty"`        //The email of a network user account or a Systems Manager device
	Mac          string `url:"mac,omitempty"`          //The MAC of a network client device or a Systems Manager device
	Serial       string `url:"serial,omitempty"`       //The serial of a Systems Manager device
	Imei         string `url:"imei,omitempty"`         //The IMEI of a Systems Manager device
	BluetoothMac string `url:"bluetoothMac,omitempty"` //The MAC of a Bluetooth client
}
type GetNetworkPoliciesByClientQueryParams struct {
	PerPage       int     `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50.
	StartingAfter string  `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string  `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	T0            string  `url:"t0,omitempty"`            //The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	Timespan      float64 `url:"timespan,omitempty"`      //The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
}
type GetNetworkSplashLoginAttemptsQueryParams struct {
	SSIDNumber      int    `url:"ssidNumber,omitempty"`      //Only return the login attempts for the specified SSID
	LoginIDentifier string `url:"loginIdentifier,omitempty"` //The username, email, or phone number used during login
	Timespan        int    `url:"timespan,omitempty"`        //The timespan, in seconds, for the login attempts. The period will be from [timespan] seconds ago until now. The maximum timespan is 3 months
}
type GetNetworkTrafficQueryParams struct {
	T0         string  `url:"t0,omitempty"`         //The beginning of the timespan for the data. The maximum lookback period is 30 days from today.
	Timespan   float64 `url:"timespan,omitempty"`   //The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 30 days.
	DeviceType string  `url:"deviceType,omitempty"` //Filter the data by device type: 'combined', 'wireless', 'switch' or 'appliance'. Defaults to 'combined'. When using 'combined', for each rule the data will come from the device type with the most usage.
}

type ResponseNetworksGetNetwork struct {
	EnrollmentString        string   `json:"enrollmentString,omitempty"`        // Enrollment string for the network
	ID                      string   `json:"id,omitempty"`                      // Network ID
	IsBoundToConfigTemplate *bool    `json:"isBoundToConfigTemplate,omitempty"` // If the network is bound to a config template
	Name                    string   `json:"name,omitempty"`                    // Network name
	Notes                   string   `json:"notes,omitempty"`                   // Notes for the network
	OrganizationID          string   `json:"organizationId,omitempty"`          // Organization ID
	ProductTypes            []string `json:"productTypes,omitempty"`            // List of the product types that the network supports
	Tags                    []string `json:"tags,omitempty"`                    // Network tags
	TimeZone                string   `json:"timeZone,omitempty"`                // Timezone of the network
	URL                     string   `json:"url,omitempty"`                     // URL to the network Dashboard UI
}
type ResponseNetworksUpdateNetwork struct {
	EnrollmentString        string   `json:"enrollmentString,omitempty"`        // Enrollment string for the network
	ID                      string   `json:"id,omitempty"`                      // Network ID
	IsBoundToConfigTemplate *bool    `json:"isBoundToConfigTemplate,omitempty"` // If the network is bound to a config template
	Name                    string   `json:"name,omitempty"`                    // Network name
	Notes                   string   `json:"notes,omitempty"`                   // Notes for the network
	OrganizationID          string   `json:"organizationId,omitempty"`          // Organization ID
	ProductTypes            []string `json:"productTypes,omitempty"`            // List of the product types that the network supports
	Tags                    []string `json:"tags,omitempty"`                    // Network tags
	TimeZone                string   `json:"timeZone,omitempty"`                // Timezone of the network
	URL                     string   `json:"url,omitempty"`                     // URL to the network Dashboard UI
}
type ResponseNetworksGetNetworkAlertsHistory []ResponseItemNetworksGetNetworkAlertsHistory // Array of ResponseNetworksGetNetworkAlertsHistory
type ResponseItemNetworksGetNetworkAlertsHistory struct {
	AlertData    *ResponseItemNetworksGetNetworkAlertsHistoryAlertData    `json:"alertData,omitempty"`    // relevant data about the event that caused the alert
	AlertType    string                                                   `json:"alertType,omitempty"`    // user friendly alert type
	AlertTypeID  string                                                   `json:"alertTypeId,omitempty"`  // type of alert
	Destinations *ResponseItemNetworksGetNetworkAlertsHistoryDestinations `json:"destinations,omitempty"` // the destinations this alert is configured to be delivered to
	Device       *ResponseItemNetworksGetNetworkAlertsHistoryDevice       `json:"device,omitempty"`       // info related to the device that caused the alert
	OccurredAt   string                                                   `json:"occurredAt,omitempty"`   // time when the event occurred
}
type ResponseItemNetworksGetNetworkAlertsHistoryAlertData interface{}
type ResponseItemNetworksGetNetworkAlertsHistoryDestinations struct {
	Email   *ResponseItemNetworksGetNetworkAlertsHistoryDestinationsEmail   `json:"email,omitempty"`   // email destinations for this alert
	Push    *ResponseItemNetworksGetNetworkAlertsHistoryDestinationsPush    `json:"push,omitempty"`    // push destinations for this alert
	Sms     *ResponseItemNetworksGetNetworkAlertsHistoryDestinationsSms     `json:"sms,omitempty"`     // sms destinations for this alert
	Webhook *ResponseItemNetworksGetNetworkAlertsHistoryDestinationsWebhook `json:"webhook,omitempty"` // webhook destinations for this alert
}
type ResponseItemNetworksGetNetworkAlertsHistoryDestinationsEmail struct {
	SentAt string `json:"sentAt,omitempty"` // time when the alert was sent to the user(s) for this channel
}
type ResponseItemNetworksGetNetworkAlertsHistoryDestinationsPush struct {
	SentAt string `json:"sentAt,omitempty"` // time when the alert was sent to the user(s) for this channel
}
type ResponseItemNetworksGetNetworkAlertsHistoryDestinationsSms struct {
	SentAt string `json:"sentAt,omitempty"` // time when the alert was sent to the user(s) for this channel
}
type ResponseItemNetworksGetNetworkAlertsHistoryDestinationsWebhook struct {
	SentAt string `json:"sentAt,omitempty"` // time when the alert was sent to the user(s) for this channel
}
type ResponseItemNetworksGetNetworkAlertsHistoryDevice struct {
	Serial string `json:"serial,omitempty"` // device serial
}
type ResponseNetworksGetNetworkAlertsSettings struct {
	Alerts              *[]ResponseNetworksGetNetworkAlertsSettingsAlerts            `json:"alerts,omitempty"`              //
	DefaultDestinations *ResponseNetworksGetNetworkAlertsSettingsDefaultDestinations `json:"defaultDestinations,omitempty"` //
}
type ResponseNetworksGetNetworkAlertsSettingsAlerts struct {
	AlertDestinations *ResponseNetworksGetNetworkAlertsSettingsAlertsAlertDestinations `json:"alertDestinations,omitempty"` //
	Enabled           *bool                                                            `json:"enabled,omitempty"`           //
	Filters           *ResponseNetworksGetNetworkAlertsSettingsAlertsFilters           `json:"filters,omitempty"`           //
	Type              string                                                           `json:"type,omitempty"`              //
}
type ResponseNetworksGetNetworkAlertsSettingsAlertsAlertDestinations struct {
	AllAdmins     *bool    `json:"allAdmins,omitempty"`     //
	Emails        []string `json:"emails,omitempty"`        //
	HTTPServerIDs []string `json:"httpServerIds,omitempty"` //
	SNMP          *bool    `json:"snmp,omitempty"`          //
}
type ResponseNetworksGetNetworkAlertsSettingsAlertsFilters struct {
	Timeout *int `json:"timeout,omitempty"` //
}
type ResponseNetworksGetNetworkAlertsSettingsDefaultDestinations struct {
	AllAdmins     *bool    `json:"allAdmins,omitempty"`     //
	Emails        []string `json:"emails,omitempty"`        //
	HTTPServerIDs []string `json:"httpServerIds,omitempty"` //
	SNMP          *bool    `json:"snmp,omitempty"`          //
}
type ResponseNetworksUpdateNetworkAlertsSettings interface{}
type ResponseNetworksBindNetwork interface{}
type ResponseNetworksGetNetworkBluetoothClients []ResponseItemNetworksGetNetworkBluetoothClients // Array of ResponseNetworksGetNetworkBluetoothClients
type ResponseItemNetworksGetNetworkBluetoothClients struct {
	DeviceName      string   `json:"deviceName,omitempty"`      //
	ID              string   `json:"id,omitempty"`              //
	InSightAlert    *bool    `json:"inSightAlert,omitempty"`    //
	LastSeen        *int     `json:"lastSeen,omitempty"`        //
	Mac             string   `json:"mac,omitempty"`             //
	Manufacturer    string   `json:"manufacturer,omitempty"`    //
	Name            string   `json:"name,omitempty"`            //
	NetworkID       string   `json:"networkId,omitempty"`       //
	OutOfSightAlert *bool    `json:"outOfSightAlert,omitempty"` //
	SeenByDeviceMac string   `json:"seenByDeviceMac,omitempty"` //
	Tags            []string `json:"tags,omitempty"`            //
}
type ResponseNetworksGetNetworkBluetoothClient struct {
	DeviceName      string   `json:"deviceName,omitempty"`      //
	ID              string   `json:"id,omitempty"`              //
	InSightAlert    *bool    `json:"inSightAlert,omitempty"`    //
	LastSeen        *int     `json:"lastSeen,omitempty"`        //
	Mac             string   `json:"mac,omitempty"`             //
	Manufacturer    string   `json:"manufacturer,omitempty"`    //
	Name            string   `json:"name,omitempty"`            //
	NetworkID       string   `json:"networkId,omitempty"`       //
	OutOfSightAlert *bool    `json:"outOfSightAlert,omitempty"` //
	SeenByDeviceMac string   `json:"seenByDeviceMac,omitempty"` //
	Tags            []string `json:"tags,omitempty"`            //
}
type ResponseNetworksGetNetworkClients struct {
	AdaptivePolicyGroup    string                                  `json:"adaptivePolicyGroup,omitempty"`    // The adaptive policy group of the client
	Description            string                                  `json:"description,omitempty"`            // Short description of the client
	DeviceTypePrediction   string                                  `json:"deviceTypePrediction,omitempty"`   // Prediction of the client's device type
	FirstSeen              *int                                    `json:"firstSeen,omitempty"`              // Timestamp client was first seen in the network
	GroupPolicy8021X       string                                  `json:"groupPolicy8021x,omitempty"`       // 802.1x group policy of the client
	ID                     string                                  `json:"id,omitempty"`                     // The ID of the client
	IP                     string                                  `json:"ip,omitempty"`                     // The IP address of the client
	IP6                    string                                  `json:"ip6,omitempty"`                    // The IPv6 address of the client
	IP6Local               string                                  `json:"ip6Local,omitempty"`               // Local IPv6 address of the client
	LastSeen               *int                                    `json:"lastSeen,omitempty"`               // Timestamp client was last seen in the network
	Mac                    string                                  `json:"mac,omitempty"`                    // The MAC address of the client
	Manufacturer           string                                  `json:"manufacturer,omitempty"`           // Manufacturer of the client
	NamedVLAN              string                                  `json:"namedVlan,omitempty"`              // Named VLAN of the client
	Notes                  string                                  `json:"notes,omitempty"`                  // Notes on the client
	Os                     string                                  `json:"os,omitempty"`                     // The operating system of the client
	PskGroup               string                                  `json:"pskGroup,omitempty"`               // iPSK name of the client
	RecentDeviceConnection string                                  `json:"recentDeviceConnection,omitempty"` // Client's most recent connection type
	RecentDeviceMac        string                                  `json:"recentDeviceMac,omitempty"`        // The MAC address of the node that the device was last connected to
	RecentDeviceName       string                                  `json:"recentDeviceName,omitempty"`       // The name of the node the device was last connected to
	RecentDeviceSerial     string                                  `json:"recentDeviceSerial,omitempty"`     // The serial of the node the device was last connected to
	SmInstalled            *bool                                   `json:"smInstalled,omitempty"`            // Status of SM for the client
	SSID                   string                                  `json:"ssid,omitempty"`                   // The name of the SSID that the client is connected to
	Status                 string                                  `json:"status,omitempty"`                 // The connection status of the client
	Switchport             string                                  `json:"switchport,omitempty"`             // The switch port that the client is connected to
	Usage                  *ResponseNetworksGetNetworkClientsUsage `json:"usage,omitempty"`                  // Usage, sent and received
	User                   string                                  `json:"user,omitempty"`                   // The username of the user of the client
	VLAN                   string                                  `json:"vlan,omitempty"`                   // The name of the VLAN that the client is connected to
	WirelessCapabilities   string                                  `json:"wirelessCapabilities,omitempty"`   // Wireless capabilities of the client
}
type ResponseNetworksGetNetworkClientsUsage struct {
	Recv *float64 `json:"recv,omitempty"` // Usage received by the client
	Sent *float64 `json:"sent,omitempty"` // Usage sent by the client
}
type ResponseNetworksGetNetworkClientsApplicationUsage []ResponseItemNetworksGetNetworkClientsApplicationUsage // Array of ResponseNetworksGetNetworkClientsApplicationUsage
type ResponseItemNetworksGetNetworkClientsApplicationUsage struct {
	ApplicationUsage *[]ResponseItemNetworksGetNetworkClientsApplicationUsageApplicationUsage `json:"applicationUsage,omitempty"` //
	ClientID         string                                                                   `json:"clientId,omitempty"`         //
	ClientIP         string                                                                   `json:"clientIp,omitempty"`         //
	ClientMac        string                                                                   `json:"clientMac,omitempty"`        //
}
type ResponseItemNetworksGetNetworkClientsApplicationUsageApplicationUsage struct {
	Application string `json:"application,omitempty"` //
	Recv        *int   `json:"recv,omitempty"`        //
	Sent        *int   `json:"sent,omitempty"`        //
}
type ResponseNetworksGetNetworkClientsBandwidthUsageHistory []ResponseItemNetworksGetNetworkClientsBandwidthUsageHistory // Array of ResponseNetworksGetNetworkClientsBandwidthUsageHistory
type ResponseItemNetworksGetNetworkClientsBandwidthUsageHistory struct {
	Total *int   `json:"total,omitempty"` //
	Ts    string `json:"ts,omitempty"`    //
}
type ResponseNetworksGetNetworkClientsOverview struct {
	Counts *ResponseNetworksGetNetworkClientsOverviewCounts `json:"counts,omitempty"` //
	Usages *ResponseNetworksGetNetworkClientsOverviewUsages `json:"usages,omitempty"` //
}
type ResponseNetworksGetNetworkClientsOverviewCounts struct {
	Total          *int `json:"total,omitempty"`          //
	WithHeavyUsage *int `json:"withHeavyUsage,omitempty"` //
}
type ResponseNetworksGetNetworkClientsOverviewUsages struct {
	Average               *int `json:"average,omitempty"`               //
	WithHeavyUsageAverage *int `json:"withHeavyUsageAverage,omitempty"` //
}
type ResponseNetworksProvisionNetworkClients interface{}
type ResponseNetworksGetNetworkClientsUsageHistories []ResponseItemNetworksGetNetworkClientsUsageHistories // Array of ResponseNetworksGetNetworkClientsUsageHistories
type ResponseItemNetworksGetNetworkClientsUsageHistories struct {
	ClientID     string                                                             `json:"clientId,omitempty"`     //
	ClientIP     string                                                             `json:"clientIp,omitempty"`     //
	ClientMac    string                                                             `json:"clientMac,omitempty"`    //
	UsageHistory *[]ResponseItemNetworksGetNetworkClientsUsageHistoriesUsageHistory `json:"usageHistory,omitempty"` //
}
type ResponseItemNetworksGetNetworkClientsUsageHistoriesUsageHistory struct {
	Recv *int   `json:"recv,omitempty"` //
	Sent *int   `json:"sent,omitempty"` //
	Ts   string `json:"ts,omitempty"`   //
}
type ResponseNetworksGetNetworkClient struct {
	Cdp                  *[][]string                                             `json:"cdp,omitempty"`                  // The Cisco discover protocol settings for the client
	ClientVpnConnections *[]ResponseNetworksGetNetworkClientClientVpnConnections `json:"clientVpnConnections,omitempty"` // VPN connections associated with the client
	Description          string                                                  `json:"description,omitempty"`          // Short description of the client
	FirstSeen            *int                                                    `json:"firstSeen,omitempty"`            // Timestamp client was first seen in the network
	ID                   string                                                  `json:"id,omitempty"`                   // The ID of the client
	IP                   string                                                  `json:"ip,omitempty"`                   // The IP address of the client
	IP6                  string                                                  `json:"ip6,omitempty"`                  // The IPv6 address of the client
	LastSeen             *int                                                    `json:"lastSeen,omitempty"`             // Timestamp client was last seen in the network
	Lldp                 *[][]string                                             `json:"lldp,omitempty"`                 // The link layer discover protocol settings for the client
	Mac                  string                                                  `json:"mac,omitempty"`                  // The MAC address of the client
	Manufacturer         string                                                  `json:"manufacturer,omitempty"`         // Manufacturer of the client
	Os                   string                                                  `json:"os,omitempty"`                   // The operating system of the client
	RecentDeviceMac      string                                                  `json:"recentDeviceMac,omitempty"`      // The MAC address of the node that the device was last connected to
	SmInstalled          *bool                                                   `json:"smInstalled,omitempty"`          // Status of SM for the client
	SSID                 string                                                  `json:"ssid,omitempty"`                 // The name of the SSID that the client is connected to
	Status               string                                                  `json:"status,omitempty"`               // The connection status of the client
	Switchport           string                                                  `json:"switchport,omitempty"`           // The switch port that the client is connected to
	User                 string                                                  `json:"user,omitempty"`                 // The username of the user of the client
	VLAN                 string                                                  `json:"vlan,omitempty"`                 // The name of the VLAN that the client is connected to
	WirelessCapabilities string                                                  `json:"wirelessCapabilities,omitempty"` // Wireless capabilities of the client
}
type ResponseNetworksGetNetworkClientClientVpnConnections struct {
	ConnectedAt    *int   `json:"connectedAt,omitempty"`    // The time the client last connected to the VPN
	DisconnectedAt *int   `json:"disconnectedAt,omitempty"` // The time the client last disconnectd from the VPN
	RemoteIP       string `json:"remoteIp,omitempty"`       // The IP address of the VPN the client last connected to
}
type ResponseNetworksGetNetworkClientPolicy struct {
	DevicePolicy  string `json:"devicePolicy,omitempty"`  //
	GroupPolicyID string `json:"groupPolicyId,omitempty"` //
	Mac           string `json:"mac,omitempty"`           //
}
type ResponseNetworksUpdateNetworkClientPolicy interface{}
type ResponseNetworksGetNetworkClientSplashAuthorizationStatus struct {
	SSIDs *ResponseNetworksGetNetworkClientSplashAuthorizationStatusSSIDs `json:"ssids,omitempty"` //
}
type ResponseNetworksGetNetworkClientSplashAuthorizationStatusSSIDs struct {
	Status0 *ResponseNetworksGetNetworkClientSplashAuthorizationStatusSSIDs0 `json:"0,omitempty"` //
	Status2 *ResponseNetworksGetNetworkClientSplashAuthorizationStatusSSIDs2 `json:"2,omitempty"` //
}
type ResponseNetworksGetNetworkClientSplashAuthorizationStatusSSIDs0 struct {
	AuthorizedAt string `json:"authorizedAt,omitempty"` //
	ExpiresAt    string `json:"expiresAt,omitempty"`    //
	IsAuthorized *bool  `json:"isAuthorized,omitempty"` //
}
type ResponseNetworksGetNetworkClientSplashAuthorizationStatusSSIDs2 struct {
	IsAuthorized *bool `json:"isAuthorized,omitempty"` //
}
type ResponseNetworksUpdateNetworkClientSplashAuthorizationStatus interface{}
type ResponseNetworksGetNetworkClientTrafficHistory []ResponseItemNetworksGetNetworkClientTrafficHistory // Array of ResponseNetworksGetNetworkClientTrafficHistory
type ResponseItemNetworksGetNetworkClientTrafficHistory struct {
	ActiveSeconds *int   `json:"activeSeconds,omitempty"` //
	Application   string `json:"application,omitempty"`   //
	Destination   string `json:"destination,omitempty"`   //
	NumFlows      *int   `json:"numFlows,omitempty"`      //
	Port          *int   `json:"port,omitempty"`          //
	Protocol      string `json:"protocol,omitempty"`      //
	Recv          *int   `json:"recv,omitempty"`          //
	Sent          *int   `json:"sent,omitempty"`          //
	Ts            string `json:"ts,omitempty"`            //
}
type ResponseNetworksGetNetworkClientUsageHistory []ResponseItemNetworksGetNetworkClientUsageHistory // Array of ResponseNetworksGetNetworkClientUsageHistory
type ResponseItemNetworksGetNetworkClientUsageHistory struct {
	Received *int   `json:"received,omitempty"` //
	Sent     *int   `json:"sent,omitempty"`     //
	Ts       string `json:"ts,omitempty"`       //
}
type ResponseNetworksGetNetworkDevices []ResponseItemNetworksGetNetworkDevices // Array of ResponseNetworksGetNetworkDevices
type ResponseItemNetworksGetNetworkDevices struct {
	Address        string                                               `json:"address,omitempty"`        //
	BeaconIDParams *ResponseItemNetworksGetNetworkDevicesBeaconIDParams `json:"beaconIdParams,omitempty"` //
	Firmware       string                                               `json:"firmware,omitempty"`       //
	FloorPlanID    string                                               `json:"floorPlanId,omitempty"`    //
	LanIP          string                                               `json:"lanIp,omitempty"`          //
	Lat            *float64                                             `json:"lat,omitempty"`            //
	Lng            *float64                                             `json:"lng,omitempty"`            //
	Mac            string                                               `json:"mac,omitempty"`            //
	Model          string                                               `json:"model,omitempty"`          //
	Name           string                                               `json:"name,omitempty"`           //
	NetworkID      string                                               `json:"networkId,omitempty"`      //
	Notes          string                                               `json:"notes,omitempty"`          //
	Serial         string                                               `json:"serial,omitempty"`         //
	Tags           string                                               `json:"tags,omitempty"`           //
}
type ResponseItemNetworksGetNetworkDevicesBeaconIDParams struct {
	Major *int   `json:"major,omitempty"` //
	Minor *int   `json:"minor,omitempty"` //
	UUID  string `json:"uuid,omitempty"`  //
}
type ResponseNetworksVmxNetworkDevicesClaim interface{}
type ResponseNetworksGetNetworkEvents struct {
	Events      *[]ResponseNetworksGetNetworkEventsEvents `json:"events,omitempty"`      // An array of events that took place in the network.
	Message     string                                    `json:"message,omitempty"`     // A message regarding the events sent. Usually 'null' unless there are no events
	PageEndAt   string                                    `json:"pageEndAt,omitempty"`   // An UTC ISO8601 string of the latest occured at time of the listed events of the page.
	PageStartAt string                                    `json:"pageStartAt,omitempty"` // An UTC ISO8601 string of the earliest occured at time of the listed events of the page.
}
type ResponseNetworksGetNetworkEventsEvents struct {
	Category          string                                           `json:"category,omitempty"`          // The category that the event type belongs to
	ClientDescription string                                           `json:"clientDescription,omitempty"` // A description of the client. This is usually the client's device name.
	ClientID          string                                           `json:"clientId,omitempty"`          // A string identifying the client. This could be a client's MAC or IP address
	ClientMac         string                                           `json:"clientMac,omitempty"`         // The client's MAC address.
	Description       string                                           `json:"description,omitempty"`       // A description of the event the happened.
	DeviceName        string                                           `json:"deviceName,omitempty"`        // The name of the device. Only shown if the device is an access point.
	DeviceSerial      string                                           `json:"deviceSerial,omitempty"`      // The serial number of the device. Only shown if the device is an access point.
	EventData         *ResponseNetworksGetNetworkEventsEventsEventData `json:"eventData,omitempty"`         // An object containing more data related to the event.
	NetworkID         string                                           `json:"networkId,omitempty"`         // The ID of the network.
	OccurredAt        string                                           `json:"occurredAt,omitempty"`        // An UTC ISO8601 string of the time the event occurred at.
	SSIDNumber        *int                                             `json:"ssidNumber,omitempty"`        // The SSID number of the device.
	Type              string                                           `json:"type,omitempty"`              // The type of event being listed.
}
type ResponseNetworksGetNetworkEventsEventsEventData struct {
	Aid       string `json:"aid,omitempty"`        // The association ID of the client.
	Channel   string `json:"channel,omitempty"`    // The radio channel the client is connecting to.
	ClientIP  string `json:"client_ip,omitempty"`  // The client's IP address
	ClientMac string `json:"client_mac,omitempty"` // The client's MAC address
	Radio     string `json:"radio,omitempty"`      // The radio band number the client is trying to connect to.
	Rssi      string `json:"rssi,omitempty"`       // The current received signal strength indication (RSSI) of the client connected to an AP.
	Vap       string `json:"vap,omitempty"`        // The virtual access point (VAP) number the client is connecting to.
}
type ResponseNetworksGetNetworkEventsEventTypes []ResponseItemNetworksGetNetworkEventsEventTypes // Array of ResponseNetworksGetNetworkEventsEventTypes
type ResponseItemNetworksGetNetworkEventsEventTypes struct {
	Category    string `json:"category,omitempty"`    // Event category
	Description string `json:"description,omitempty"` // Description of the event
	Type        string `json:"type,omitempty"`        // Event type
}
type ResponseNetworksGetNetworkFirmwareUpgrades struct {
	Products      *ResponseNetworksGetNetworkFirmwareUpgradesProducts      `json:"products,omitempty"`      // The network devices to be updated
	Timezone      string                                                   `json:"timezone,omitempty"`      // The timezone for the network
	UpgradeWindow *ResponseNetworksGetNetworkFirmwareUpgradesUpgradeWindow `json:"upgradeWindow,omitempty"` // Upgrade window for devices in network
}
type ResponseNetworksGetNetworkFirmwareUpgradesProducts struct {
	Appliance       *ResponseNetworksGetNetworkFirmwareUpgradesProductsAppliance       `json:"appliance,omitempty"`       // The network device to be updated
	Camera          *ResponseNetworksGetNetworkFirmwareUpgradesProductsCamera          `json:"camera,omitempty"`          // The network device to be updated
	CellularGateway *ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGateway `json:"cellularGateway,omitempty"` // The network device to be updated
	Sensor          *ResponseNetworksGetNetworkFirmwareUpgradesProductsSensor          `json:"sensor,omitempty"`          // The network device to be updated
	Switch          *ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitch          `json:"switch,omitempty"`          // The network device to be updated
	Wireless        *ResponseNetworksGetNetworkFirmwareUpgradesProductsWireless        `json:"wireless,omitempty"`        // The network device to be updated
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsAppliance struct {
	AvailableVersions            *[]ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceAvailableVersions `json:"availableVersions,omitempty"`            // Firmware versions available for upgrade
	CurrentVersion               *ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceCurrentVersion      `json:"currentVersion,omitempty"`               // Details of the current version on the device
	LastUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceLastUpgrade         `json:"lastUpgrade,omitempty"`                  // Details of the last firmware upgrade on the device
	NextUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceNextUpgrade         `json:"nextUpgrade,omitempty"`                  // Details of the next firmware upgrade on the device
	ParticipateInNextBetaRelease *bool                                                                           `json:"participateInNextBetaRelease,omitempty"` // Whether or not the network wants beta firmware
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceAvailableVersions struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceCurrentVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceLastUpgrade struct {
	FromVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceLastUpgradeFromVersion `json:"fromVersion,omitempty"` // Details of the version the device upgraded from
	Time        string                                                                             `json:"time,omitempty"`        // Timestamp of the last successful firmware upgrade
	ToVersion   *ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceLastUpgradeToVersion   `json:"toVersion,omitempty"`   // Details of the version the device upgraded to
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceLastUpgradeFromVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceLastUpgradeToVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceNextUpgrade struct {
	Time      string                                                                           `json:"time,omitempty"`      // Timestamp of the next scheduled firmware upgrade
	ToVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceNextUpgradeToVersion `json:"toVersion,omitempty"` // Details of the version the device will upgrade to if it exists
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceNextUpgradeToVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsCamera struct {
	AvailableVersions            *[]ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraAvailableVersions `json:"availableVersions,omitempty"`            // Firmware versions available for upgrade
	CurrentVersion               *ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraCurrentVersion      `json:"currentVersion,omitempty"`               // Details of the current version on the device
	LastUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraLastUpgrade         `json:"lastUpgrade,omitempty"`                  // Details of the last firmware upgrade on the device
	NextUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraNextUpgrade         `json:"nextUpgrade,omitempty"`                  // Details of the next firmware upgrade on the device
	ParticipateInNextBetaRelease *bool                                                                        `json:"participateInNextBetaRelease,omitempty"` // Whether or not the network wants beta firmware
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraAvailableVersions struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraCurrentVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraLastUpgrade struct {
	FromVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraLastUpgradeFromVersion `json:"fromVersion,omitempty"` // Details of the version the device upgraded from
	Time        string                                                                          `json:"time,omitempty"`        // Timestamp of the last successful firmware upgrade
	ToVersion   *ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraLastUpgradeToVersion   `json:"toVersion,omitempty"`   // Details of the version the device upgraded to
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraLastUpgradeFromVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraLastUpgradeToVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraNextUpgrade struct {
	Time      string                                                                        `json:"time,omitempty"`      // Timestamp of the next scheduled firmware upgrade
	ToVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraNextUpgradeToVersion `json:"toVersion,omitempty"` // Details of the version the device will upgrade to if it exists
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraNextUpgradeToVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGateway struct {
	AvailableVersions            *[]ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayAvailableVersions `json:"availableVersions,omitempty"`            // Firmware versions available for upgrade
	CurrentVersion               *ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayCurrentVersion      `json:"currentVersion,omitempty"`               // Details of the current version on the device
	LastUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayLastUpgrade         `json:"lastUpgrade,omitempty"`                  // Details of the last firmware upgrade on the device
	NextUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayNextUpgrade         `json:"nextUpgrade,omitempty"`                  // Details of the next firmware upgrade on the device
	ParticipateInNextBetaRelease *bool                                                                                 `json:"participateInNextBetaRelease,omitempty"` // Whether or not the network wants beta firmware
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayAvailableVersions struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayCurrentVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayLastUpgrade struct {
	FromVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayLastUpgradeFromVersion `json:"fromVersion,omitempty"` // Details of the version the device upgraded from
	Time        string                                                                                   `json:"time,omitempty"`        // Timestamp of the last successful firmware upgrade
	ToVersion   *ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayLastUpgradeToVersion   `json:"toVersion,omitempty"`   // Details of the version the device upgraded to
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayLastUpgradeFromVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayLastUpgradeToVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayNextUpgrade struct {
	Time      string                                                                                 `json:"time,omitempty"`      // Timestamp of the next scheduled firmware upgrade
	ToVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayNextUpgradeToVersion `json:"toVersion,omitempty"` // Details of the version the device will upgrade to if it exists
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayNextUpgradeToVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsSensor struct {
	AvailableVersions            *[]ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorAvailableVersions `json:"availableVersions,omitempty"`            // Firmware versions available for upgrade
	CurrentVersion               *ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorCurrentVersion      `json:"currentVersion,omitempty"`               // Details of the current version on the device
	LastUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorLastUpgrade         `json:"lastUpgrade,omitempty"`                  // Details of the last firmware upgrade on the device
	NextUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorNextUpgrade         `json:"nextUpgrade,omitempty"`                  // Details of the next firmware upgrade on the device
	ParticipateInNextBetaRelease *bool                                                                        `json:"participateInNextBetaRelease,omitempty"` // Whether or not the network wants beta firmware
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorAvailableVersions struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorCurrentVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorLastUpgrade struct {
	FromVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorLastUpgradeFromVersion `json:"fromVersion,omitempty"` // Details of the version the device upgraded from
	Time        string                                                                          `json:"time,omitempty"`        // Timestamp of the last successful firmware upgrade
	ToVersion   *ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorLastUpgradeToVersion   `json:"toVersion,omitempty"`   // Details of the version the device upgraded to
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorLastUpgradeFromVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorLastUpgradeToVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorNextUpgrade struct {
	Time      string                                                                        `json:"time,omitempty"`      // Timestamp of the next scheduled firmware upgrade
	ToVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorNextUpgradeToVersion `json:"toVersion,omitempty"` // Details of the version the device will upgrade to if it exists
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorNextUpgradeToVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitch struct {
	AvailableVersions            *[]ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchAvailableVersions `json:"availableVersions,omitempty"`            // Firmware versions available for upgrade
	CurrentVersion               *ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchCurrentVersion      `json:"currentVersion,omitempty"`               // Details of the current version on the device
	LastUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchLastUpgrade         `json:"lastUpgrade,omitempty"`                  // Details of the last firmware upgrade on the device
	NextUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchNextUpgrade         `json:"nextUpgrade,omitempty"`                  // Details of the next firmware upgrade on the device
	ParticipateInNextBetaRelease *bool                                                                        `json:"participateInNextBetaRelease,omitempty"` // Whether or not the network wants beta firmware
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchAvailableVersions struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchCurrentVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchLastUpgrade struct {
	FromVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchLastUpgradeFromVersion `json:"fromVersion,omitempty"` // Details of the version the device upgraded from
	Time        string                                                                          `json:"time,omitempty"`        // Timestamp of the last successful firmware upgrade
	ToVersion   *ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchLastUpgradeToVersion   `json:"toVersion,omitempty"`   // Details of the version the device upgraded to
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchLastUpgradeFromVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchLastUpgradeToVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchNextUpgrade struct {
	Time      string                                                                        `json:"time,omitempty"`      // Timestamp of the next scheduled firmware upgrade
	ToVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchNextUpgradeToVersion `json:"toVersion,omitempty"` // Details of the version the device will upgrade to if it exists
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchNextUpgradeToVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsWireless struct {
	AvailableVersions            *[]ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessAvailableVersions `json:"availableVersions,omitempty"`            // Firmware versions available for upgrade
	CurrentVersion               *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessCurrentVersion      `json:"currentVersion,omitempty"`               // Details of the current version on the device
	LastUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessLastUpgrade         `json:"lastUpgrade,omitempty"`                  // Details of the last firmware upgrade on the device
	NextUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessNextUpgrade         `json:"nextUpgrade,omitempty"`                  // Details of the next firmware upgrade on the device
	ParticipateInNextBetaRelease *bool                                                                          `json:"participateInNextBetaRelease,omitempty"` // Whether or not the network wants beta firmware
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessAvailableVersions struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessCurrentVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessLastUpgrade struct {
	FromVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessLastUpgradeFromVersion `json:"fromVersion,omitempty"` // Details of the version the device upgraded from
	Time        string                                                                            `json:"time,omitempty"`        // Timestamp of the last successful firmware upgrade
	ToVersion   *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessLastUpgradeToVersion   `json:"toVersion,omitempty"`   // Details of the version the device upgraded to
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessLastUpgradeFromVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessLastUpgradeToVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessNextUpgrade struct {
	Time      string                                                                          `json:"time,omitempty"`      // Timestamp of the next scheduled firmware upgrade
	ToVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessNextUpgradeToVersion `json:"toVersion,omitempty"` // Details of the version the device will upgrade to if it exists
}
type ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessNextUpgradeToVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksGetNetworkFirmwareUpgradesUpgradeWindow struct {
	DayOfWeek string `json:"dayOfWeek,omitempty"` // Day of the week
	HourOfDay string `json:"hourOfDay,omitempty"` // Hour of the day
}
type ResponseNetworksUpdateNetworkFirmwareUpgrades struct {
	Products      *ResponseNetworksUpdateNetworkFirmwareUpgradesProducts      `json:"products,omitempty"`      // The network devices to be updated
	Timezone      string                                                      `json:"timezone,omitempty"`      // The timezone for the network
	UpgradeWindow *ResponseNetworksUpdateNetworkFirmwareUpgradesUpgradeWindow `json:"upgradeWindow,omitempty"` // Upgrade window for devices in network
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProducts struct {
	Appliance       *ResponseNetworksUpdateNetworkFirmwareUpgradesProductsAppliance       `json:"appliance,omitempty"`       // The network device to be updated
	Camera          *ResponseNetworksUpdateNetworkFirmwareUpgradesProductsCamera          `json:"camera,omitempty"`          // The network device to be updated
	CellularGateway *ResponseNetworksUpdateNetworkFirmwareUpgradesProductsCellularGateway `json:"cellularGateway,omitempty"` // The network device to be updated
	Sensor          *ResponseNetworksUpdateNetworkFirmwareUpgradesProductsSensor          `json:"sensor,omitempty"`          // The network device to be updated
	Switch          *ResponseNetworksUpdateNetworkFirmwareUpgradesProductsSwitch          `json:"switch,omitempty"`          // The network device to be updated
	Wireless        *ResponseNetworksUpdateNetworkFirmwareUpgradesProductsWireless        `json:"wireless,omitempty"`        // The network device to be updated
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsAppliance struct {
	AvailableVersions            *[]ResponseNetworksUpdateNetworkFirmwareUpgradesProductsApplianceAvailableVersions `json:"availableVersions,omitempty"`            // Firmware versions available for upgrade
	CurrentVersion               *ResponseNetworksUpdateNetworkFirmwareUpgradesProductsApplianceCurrentVersion      `json:"currentVersion,omitempty"`               // Details of the current version on the device
	LastUpgrade                  *ResponseNetworksUpdateNetworkFirmwareUpgradesProductsApplianceLastUpgrade         `json:"lastUpgrade,omitempty"`                  // Details of the last firmware upgrade on the device
	NextUpgrade                  *ResponseNetworksUpdateNetworkFirmwareUpgradesProductsApplianceNextUpgrade         `json:"nextUpgrade,omitempty"`                  // Details of the next firmware upgrade on the device
	ParticipateInNextBetaRelease *bool                                                                              `json:"participateInNextBetaRelease,omitempty"` // Whether or not the network wants beta firmware
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsApplianceAvailableVersions struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsApplianceCurrentVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsApplianceLastUpgrade struct {
	FromVersion *ResponseNetworksUpdateNetworkFirmwareUpgradesProductsApplianceLastUpgradeFromVersion `json:"fromVersion,omitempty"` // Details of the version the device upgraded from
	Time        string                                                                                `json:"time,omitempty"`        // Timestamp of the last successful firmware upgrade
	ToVersion   *ResponseNetworksUpdateNetworkFirmwareUpgradesProductsApplianceLastUpgradeToVersion   `json:"toVersion,omitempty"`   // Details of the version the device upgraded to
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsApplianceLastUpgradeFromVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsApplianceLastUpgradeToVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsApplianceNextUpgrade struct {
	Time      string                                                                              `json:"time,omitempty"`      // Timestamp of the next scheduled firmware upgrade
	ToVersion *ResponseNetworksUpdateNetworkFirmwareUpgradesProductsApplianceNextUpgradeToVersion `json:"toVersion,omitempty"` // Details of the version the device will upgrade to if it exists
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsApplianceNextUpgradeToVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsCamera struct {
	AvailableVersions            *[]ResponseNetworksUpdateNetworkFirmwareUpgradesProductsCameraAvailableVersions `json:"availableVersions,omitempty"`            // Firmware versions available for upgrade
	CurrentVersion               *ResponseNetworksUpdateNetworkFirmwareUpgradesProductsCameraCurrentVersion      `json:"currentVersion,omitempty"`               // Details of the current version on the device
	LastUpgrade                  *ResponseNetworksUpdateNetworkFirmwareUpgradesProductsCameraLastUpgrade         `json:"lastUpgrade,omitempty"`                  // Details of the last firmware upgrade on the device
	NextUpgrade                  *ResponseNetworksUpdateNetworkFirmwareUpgradesProductsCameraNextUpgrade         `json:"nextUpgrade,omitempty"`                  // Details of the next firmware upgrade on the device
	ParticipateInNextBetaRelease *bool                                                                           `json:"participateInNextBetaRelease,omitempty"` // Whether or not the network wants beta firmware
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsCameraAvailableVersions struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsCameraCurrentVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsCameraLastUpgrade struct {
	FromVersion *ResponseNetworksUpdateNetworkFirmwareUpgradesProductsCameraLastUpgradeFromVersion `json:"fromVersion,omitempty"` // Details of the version the device upgraded from
	Time        string                                                                             `json:"time,omitempty"`        // Timestamp of the last successful firmware upgrade
	ToVersion   *ResponseNetworksUpdateNetworkFirmwareUpgradesProductsCameraLastUpgradeToVersion   `json:"toVersion,omitempty"`   // Details of the version the device upgraded to
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsCameraLastUpgradeFromVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsCameraLastUpgradeToVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsCameraNextUpgrade struct {
	Time      string                                                                           `json:"time,omitempty"`      // Timestamp of the next scheduled firmware upgrade
	ToVersion *ResponseNetworksUpdateNetworkFirmwareUpgradesProductsCameraNextUpgradeToVersion `json:"toVersion,omitempty"` // Details of the version the device will upgrade to if it exists
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsCameraNextUpgradeToVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsCellularGateway struct {
	AvailableVersions            *[]ResponseNetworksUpdateNetworkFirmwareUpgradesProductsCellularGatewayAvailableVersions `json:"availableVersions,omitempty"`            // Firmware versions available for upgrade
	CurrentVersion               *ResponseNetworksUpdateNetworkFirmwareUpgradesProductsCellularGatewayCurrentVersion      `json:"currentVersion,omitempty"`               // Details of the current version on the device
	LastUpgrade                  *ResponseNetworksUpdateNetworkFirmwareUpgradesProductsCellularGatewayLastUpgrade         `json:"lastUpgrade,omitempty"`                  // Details of the last firmware upgrade on the device
	NextUpgrade                  *ResponseNetworksUpdateNetworkFirmwareUpgradesProductsCellularGatewayNextUpgrade         `json:"nextUpgrade,omitempty"`                  // Details of the next firmware upgrade on the device
	ParticipateInNextBetaRelease *bool                                                                                    `json:"participateInNextBetaRelease,omitempty"` // Whether or not the network wants beta firmware
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsCellularGatewayAvailableVersions struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsCellularGatewayCurrentVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsCellularGatewayLastUpgrade struct {
	FromVersion *ResponseNetworksUpdateNetworkFirmwareUpgradesProductsCellularGatewayLastUpgradeFromVersion `json:"fromVersion,omitempty"` // Details of the version the device upgraded from
	Time        string                                                                                      `json:"time,omitempty"`        // Timestamp of the last successful firmware upgrade
	ToVersion   *ResponseNetworksUpdateNetworkFirmwareUpgradesProductsCellularGatewayLastUpgradeToVersion   `json:"toVersion,omitempty"`   // Details of the version the device upgraded to
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsCellularGatewayLastUpgradeFromVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsCellularGatewayLastUpgradeToVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsCellularGatewayNextUpgrade struct {
	Time      string                                                                                    `json:"time,omitempty"`      // Timestamp of the next scheduled firmware upgrade
	ToVersion *ResponseNetworksUpdateNetworkFirmwareUpgradesProductsCellularGatewayNextUpgradeToVersion `json:"toVersion,omitempty"` // Details of the version the device will upgrade to if it exists
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsCellularGatewayNextUpgradeToVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsSensor struct {
	AvailableVersions            *[]ResponseNetworksUpdateNetworkFirmwareUpgradesProductsSensorAvailableVersions `json:"availableVersions,omitempty"`            // Firmware versions available for upgrade
	CurrentVersion               *ResponseNetworksUpdateNetworkFirmwareUpgradesProductsSensorCurrentVersion      `json:"currentVersion,omitempty"`               // Details of the current version on the device
	LastUpgrade                  *ResponseNetworksUpdateNetworkFirmwareUpgradesProductsSensorLastUpgrade         `json:"lastUpgrade,omitempty"`                  // Details of the last firmware upgrade on the device
	NextUpgrade                  *ResponseNetworksUpdateNetworkFirmwareUpgradesProductsSensorNextUpgrade         `json:"nextUpgrade,omitempty"`                  // Details of the next firmware upgrade on the device
	ParticipateInNextBetaRelease *bool                                                                           `json:"participateInNextBetaRelease,omitempty"` // Whether or not the network wants beta firmware
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsSensorAvailableVersions struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsSensorCurrentVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsSensorLastUpgrade struct {
	FromVersion *ResponseNetworksUpdateNetworkFirmwareUpgradesProductsSensorLastUpgradeFromVersion `json:"fromVersion,omitempty"` // Details of the version the device upgraded from
	Time        string                                                                             `json:"time,omitempty"`        // Timestamp of the last successful firmware upgrade
	ToVersion   *ResponseNetworksUpdateNetworkFirmwareUpgradesProductsSensorLastUpgradeToVersion   `json:"toVersion,omitempty"`   // Details of the version the device upgraded to
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsSensorLastUpgradeFromVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsSensorLastUpgradeToVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsSensorNextUpgrade struct {
	Time      string                                                                           `json:"time,omitempty"`      // Timestamp of the next scheduled firmware upgrade
	ToVersion *ResponseNetworksUpdateNetworkFirmwareUpgradesProductsSensorNextUpgradeToVersion `json:"toVersion,omitempty"` // Details of the version the device will upgrade to if it exists
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsSensorNextUpgradeToVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsSwitch struct {
	AvailableVersions            *[]ResponseNetworksUpdateNetworkFirmwareUpgradesProductsSwitchAvailableVersions `json:"availableVersions,omitempty"`            // Firmware versions available for upgrade
	CurrentVersion               *ResponseNetworksUpdateNetworkFirmwareUpgradesProductsSwitchCurrentVersion      `json:"currentVersion,omitempty"`               // Details of the current version on the device
	LastUpgrade                  *ResponseNetworksUpdateNetworkFirmwareUpgradesProductsSwitchLastUpgrade         `json:"lastUpgrade,omitempty"`                  // Details of the last firmware upgrade on the device
	NextUpgrade                  *ResponseNetworksUpdateNetworkFirmwareUpgradesProductsSwitchNextUpgrade         `json:"nextUpgrade,omitempty"`                  // Details of the next firmware upgrade on the device
	ParticipateInNextBetaRelease *bool                                                                           `json:"participateInNextBetaRelease,omitempty"` // Whether or not the network wants beta firmware
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsSwitchAvailableVersions struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsSwitchCurrentVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsSwitchLastUpgrade struct {
	FromVersion *ResponseNetworksUpdateNetworkFirmwareUpgradesProductsSwitchLastUpgradeFromVersion `json:"fromVersion,omitempty"` // Details of the version the device upgraded from
	Time        string                                                                             `json:"time,omitempty"`        // Timestamp of the last successful firmware upgrade
	ToVersion   *ResponseNetworksUpdateNetworkFirmwareUpgradesProductsSwitchLastUpgradeToVersion   `json:"toVersion,omitempty"`   // Details of the version the device upgraded to
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsSwitchLastUpgradeFromVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsSwitchLastUpgradeToVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsSwitchNextUpgrade struct {
	Time      string                                                                           `json:"time,omitempty"`      // Timestamp of the next scheduled firmware upgrade
	ToVersion *ResponseNetworksUpdateNetworkFirmwareUpgradesProductsSwitchNextUpgradeToVersion `json:"toVersion,omitempty"` // Details of the version the device will upgrade to if it exists
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsSwitchNextUpgradeToVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsWireless struct {
	AvailableVersions            *[]ResponseNetworksUpdateNetworkFirmwareUpgradesProductsWirelessAvailableVersions `json:"availableVersions,omitempty"`            // Firmware versions available for upgrade
	CurrentVersion               *ResponseNetworksUpdateNetworkFirmwareUpgradesProductsWirelessCurrentVersion      `json:"currentVersion,omitempty"`               // Details of the current version on the device
	LastUpgrade                  *ResponseNetworksUpdateNetworkFirmwareUpgradesProductsWirelessLastUpgrade         `json:"lastUpgrade,omitempty"`                  // Details of the last firmware upgrade on the device
	NextUpgrade                  *ResponseNetworksUpdateNetworkFirmwareUpgradesProductsWirelessNextUpgrade         `json:"nextUpgrade,omitempty"`                  // Details of the next firmware upgrade on the device
	ParticipateInNextBetaRelease *bool                                                                             `json:"participateInNextBetaRelease,omitempty"` // Whether or not the network wants beta firmware
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsWirelessAvailableVersions struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsWirelessCurrentVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsWirelessLastUpgrade struct {
	FromVersion *ResponseNetworksUpdateNetworkFirmwareUpgradesProductsWirelessLastUpgradeFromVersion `json:"fromVersion,omitempty"` // Details of the version the device upgraded from
	Time        string                                                                               `json:"time,omitempty"`        // Timestamp of the last successful firmware upgrade
	ToVersion   *ResponseNetworksUpdateNetworkFirmwareUpgradesProductsWirelessLastUpgradeToVersion   `json:"toVersion,omitempty"`   // Details of the version the device upgraded to
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsWirelessLastUpgradeFromVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsWirelessLastUpgradeToVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsWirelessNextUpgrade struct {
	Time      string                                                                             `json:"time,omitempty"`      // Timestamp of the next scheduled firmware upgrade
	ToVersion *ResponseNetworksUpdateNetworkFirmwareUpgradesProductsWirelessNextUpgradeToVersion `json:"toVersion,omitempty"` // Details of the version the device will upgrade to if it exists
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesProductsWirelessNextUpgradeToVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesUpgradeWindow struct {
	DayOfWeek string `json:"dayOfWeek,omitempty"` // Day of the week
	HourOfDay string `json:"hourOfDay,omitempty"` // Hour of the day
}
type ResponseNetworksCreateNetworkFirmwareUpgradesRollback struct {
	Product        string                                                          `json:"product,omitempty"`        // Product type to rollback (if the network is a combined network)
	Reasons        *[]ResponseNetworksCreateNetworkFirmwareUpgradesRollbackReasons `json:"reasons,omitempty"`        // Reasons for the rollback
	Status         string                                                          `json:"status,omitempty"`         // Status of the rollback
	Time           string                                                          `json:"time,omitempty"`           // Scheduled time for the rollback
	ToVersion      *ResponseNetworksCreateNetworkFirmwareUpgradesRollbackToVersion `json:"toVersion,omitempty"`      // Version to downgrade to (if the network has firmware flexibility)
	UpgradeBatchID string                                                          `json:"upgradeBatchId,omitempty"` // Batch ID of the firmware rollback
}
type ResponseNetworksCreateNetworkFirmwareUpgradesRollbackReasons struct {
	Category string `json:"category,omitempty"` // Reason for the rollback
	Comment  string `json:"comment,omitempty"`  // Additional comment about the rollback
}
type ResponseNetworksCreateNetworkFirmwareUpgradesRollbackToVersion struct {
	Firmware    string `json:"firmware,omitempty"`    // Name of the firmware version
	ID          string `json:"id,omitempty"`          // Firmware version identifier
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseNetworksGetNetworkFirmwareUpgradesStagedEvents struct {
	Products *ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsProducts  `json:"products,omitempty"` // The network devices to be updated
	Reasons  *[]ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsReasons `json:"reasons,omitempty"`  // Reasons for the rollback
	Stages   *[]ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsStages  `json:"stages,omitempty"`   // The ordered stages in the network
}
type ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsProducts struct {
	Switch *ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsProductsSwitch `json:"switch,omitempty"` // The Switch network to be updated
}
type ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsProductsSwitch struct {
	NextUpgrade *ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade `json:"nextUpgrade,omitempty"` // Details of the next firmware upgrade
}
type ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade struct {
	ToVersion *ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion `json:"toVersion,omitempty"` // Details of the version the device will upgrade to
}
type ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion struct {
	ID        string `json:"id,omitempty"`        // Id of the Version being upgraded to
	ShortName string `json:"shortName,omitempty"` // Firmware version short name
}
type ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsReasons struct {
	Category string `json:"category,omitempty"` // Reason for the rollback
	Comment  string `json:"comment,omitempty"`  // Additional comment about the rollback
}
type ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsStages struct {
	Group      *ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsStagesGroup      `json:"group,omitempty"`      // The staged upgrade group
	Milestones *ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsStagesMilestones `json:"milestones,omitempty"` // The Staged Upgrade Milestones for the stage
	Status     string                                                                  `json:"status,omitempty"`     // Current upgrade status of the group
}
type ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsStagesGroup struct {
	Description string `json:"description,omitempty"` // Description of the Staged Upgrade Group
	ID          string `json:"id,omitempty"`          // Id of the Staged Upgrade Group
	Name        string `json:"name,omitempty"`        // Name of the Staged Upgrade Group
}
type ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsStagesMilestones struct {
	CanceledAt   string `json:"canceledAt,omitempty"`   // Time that the group was canceled
	CompletedAt  string `json:"completedAt,omitempty"`  // Finish time for the group
	ScheduledFor string `json:"scheduledFor,omitempty"` // Scheduled start time for the group
	StartedAt    string `json:"startedAt,omitempty"`    // Start time for the group
}
type ResponseNetworksCreateNetworkFirmwareUpgradesStagedEvent struct {
	Products *ResponseNetworksCreateNetworkFirmwareUpgradesStagedEventProducts  `json:"products,omitempty"` // The network devices to be updated
	Reasons  *[]ResponseNetworksCreateNetworkFirmwareUpgradesStagedEventReasons `json:"reasons,omitempty"`  // Reasons for the rollback
	Stages   *[]ResponseNetworksCreateNetworkFirmwareUpgradesStagedEventStages  `json:"stages,omitempty"`   // The ordered stages in the network
}
type ResponseNetworksCreateNetworkFirmwareUpgradesStagedEventProducts struct {
	Switch *ResponseNetworksCreateNetworkFirmwareUpgradesStagedEventProductsSwitch `json:"switch,omitempty"` // The Switch network to be updated
}
type ResponseNetworksCreateNetworkFirmwareUpgradesStagedEventProductsSwitch struct {
	NextUpgrade *ResponseNetworksCreateNetworkFirmwareUpgradesStagedEventProductsSwitchNextUpgrade `json:"nextUpgrade,omitempty"` // Details of the next firmware upgrade
}
type ResponseNetworksCreateNetworkFirmwareUpgradesStagedEventProductsSwitchNextUpgrade struct {
	ToVersion *ResponseNetworksCreateNetworkFirmwareUpgradesStagedEventProductsSwitchNextUpgradeToVersion `json:"toVersion,omitempty"` // Details of the version the device will upgrade to
}
type ResponseNetworksCreateNetworkFirmwareUpgradesStagedEventProductsSwitchNextUpgradeToVersion struct {
	ID        string `json:"id,omitempty"`        // Id of the Version being upgraded to
	ShortName string `json:"shortName,omitempty"` // Firmware version short name
}
type ResponseNetworksCreateNetworkFirmwareUpgradesStagedEventReasons struct {
	Category string `json:"category,omitempty"` // Reason for the rollback
	Comment  string `json:"comment,omitempty"`  // Additional comment about the rollback
}
type ResponseNetworksCreateNetworkFirmwareUpgradesStagedEventStages struct {
	Group      *ResponseNetworksCreateNetworkFirmwareUpgradesStagedEventStagesGroup      `json:"group,omitempty"`      // The staged upgrade group
	Milestones *ResponseNetworksCreateNetworkFirmwareUpgradesStagedEventStagesMilestones `json:"milestones,omitempty"` // The Staged Upgrade Milestones for the stage
	Status     string                                                                    `json:"status,omitempty"`     // Current upgrade status of the group
}
type ResponseNetworksCreateNetworkFirmwareUpgradesStagedEventStagesGroup struct {
	Description string `json:"description,omitempty"` // Description of the Staged Upgrade Group
	ID          string `json:"id,omitempty"`          // Id of the Staged Upgrade Group
	Name        string `json:"name,omitempty"`        // Name of the Staged Upgrade Group
}
type ResponseNetworksCreateNetworkFirmwareUpgradesStagedEventStagesMilestones struct {
	CanceledAt   string `json:"canceledAt,omitempty"`   // Time that the group was canceled
	CompletedAt  string `json:"completedAt,omitempty"`  // Finish time for the group
	ScheduledFor string `json:"scheduledFor,omitempty"` // Scheduled start time for the group
	StartedAt    string `json:"startedAt,omitempty"`    // Start time for the group
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesStagedEvents struct {
	Products *ResponseNetworksUpdateNetworkFirmwareUpgradesStagedEventsProducts  `json:"products,omitempty"` // The network devices to be updated
	Reasons  *[]ResponseNetworksUpdateNetworkFirmwareUpgradesStagedEventsReasons `json:"reasons,omitempty"`  // Reasons for the rollback
	Stages   *[]ResponseNetworksUpdateNetworkFirmwareUpgradesStagedEventsStages  `json:"stages,omitempty"`   // The ordered stages in the network
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesStagedEventsProducts struct {
	Switch *ResponseNetworksUpdateNetworkFirmwareUpgradesStagedEventsProductsSwitch `json:"switch,omitempty"` // The Switch network to be updated
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesStagedEventsProductsSwitch struct {
	NextUpgrade *ResponseNetworksUpdateNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade `json:"nextUpgrade,omitempty"` // Details of the next firmware upgrade
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade struct {
	ToVersion *ResponseNetworksUpdateNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion `json:"toVersion,omitempty"` // Details of the version the device will upgrade to
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion struct {
	ID        string `json:"id,omitempty"`        // Id of the Version being upgraded to
	ShortName string `json:"shortName,omitempty"` // Firmware version short name
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesStagedEventsReasons struct {
	Category string `json:"category,omitempty"` // Reason for the rollback
	Comment  string `json:"comment,omitempty"`  // Additional comment about the rollback
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesStagedEventsStages struct {
	Group      *ResponseNetworksUpdateNetworkFirmwareUpgradesStagedEventsStagesGroup      `json:"group,omitempty"`      // The staged upgrade group
	Milestones *ResponseNetworksUpdateNetworkFirmwareUpgradesStagedEventsStagesMilestones `json:"milestones,omitempty"` // The Staged Upgrade Milestones for the stage
	Status     string                                                                     `json:"status,omitempty"`     // Current upgrade status of the group
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesStagedEventsStagesGroup struct {
	Description string `json:"description,omitempty"` // Description of the Staged Upgrade Group
	ID          string `json:"id,omitempty"`          // Id of the Staged Upgrade Group
	Name        string `json:"name,omitempty"`        // Name of the Staged Upgrade Group
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesStagedEventsStagesMilestones struct {
	CanceledAt   string `json:"canceledAt,omitempty"`   // Time that the group was canceled
	CompletedAt  string `json:"completedAt,omitempty"`  // Finish time for the group
	ScheduledFor string `json:"scheduledFor,omitempty"` // Scheduled start time for the group
	StartedAt    string `json:"startedAt,omitempty"`    // Start time for the group
}
type ResponseNetworksDeferNetworkFirmwareUpgradesStagedEvents struct {
	Products *ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsProducts  `json:"products,omitempty"` // The network devices to be updated
	Reasons  *[]ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsReasons `json:"reasons,omitempty"`  // Reasons for the rollback
	Stages   *[]ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsStages  `json:"stages,omitempty"`   // The ordered stages in the network
}
type ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsProducts struct {
	Switch *ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsProductsSwitch `json:"switch,omitempty"` // The Switch network to be updated
}
type ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsProductsSwitch struct {
	NextUpgrade *ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade `json:"nextUpgrade,omitempty"` // Details of the next firmware upgrade
}
type ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade struct {
	ToVersion *ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion `json:"toVersion,omitempty"` // Details of the version the device will upgrade to
}
type ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion struct {
	ID        string `json:"id,omitempty"`        // Id of the Version being upgraded to
	ShortName string `json:"shortName,omitempty"` // Firmware version short name
}
type ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsReasons struct {
	Category string `json:"category,omitempty"` // Reason for the rollback
	Comment  string `json:"comment,omitempty"`  // Additional comment about the rollback
}
type ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsStages struct {
	Group      *ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsStagesGroup      `json:"group,omitempty"`      // The staged upgrade group
	Milestones *ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsStagesMilestones `json:"milestones,omitempty"` // The Staged Upgrade Milestones for the stage
	Status     string                                                                    `json:"status,omitempty"`     // Current upgrade status of the group
}
type ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsStagesGroup struct {
	Description string `json:"description,omitempty"` // Description of the Staged Upgrade Group
	ID          string `json:"id,omitempty"`          // Id of the Staged Upgrade Group
	Name        string `json:"name,omitempty"`        // Name of the Staged Upgrade Group
}
type ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsStagesMilestones struct {
	CanceledAt   string `json:"canceledAt,omitempty"`   // Time that the group was canceled
	CompletedAt  string `json:"completedAt,omitempty"`  // Finish time for the group
	ScheduledFor string `json:"scheduledFor,omitempty"` // Scheduled start time for the group
	StartedAt    string `json:"startedAt,omitempty"`    // Start time for the group
}
type ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEvents struct {
	Products *ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsProducts  `json:"products,omitempty"` // The network devices to be updated
	Reasons  *[]ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsReasons `json:"reasons,omitempty"`  // Reasons for the rollback
	Stages   *[]ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStages  `json:"stages,omitempty"`   // The ordered stages in the network
}
type ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsProducts struct {
	Switch *ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsProductsSwitch `json:"switch,omitempty"` // The Switch network to be updated
}
type ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsProductsSwitch struct {
	NextUpgrade *ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade `json:"nextUpgrade,omitempty"` // Details of the next firmware upgrade
}
type ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade struct {
	ToVersion *ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion `json:"toVersion,omitempty"` // Details of the version the device will upgrade to
}
type ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion struct {
	ID        string `json:"id,omitempty"`        // Id of the Version being upgraded to
	ShortName string `json:"shortName,omitempty"` // Firmware version short name
}
type ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsReasons struct {
	Category string `json:"category,omitempty"` // Reason for the rollback
	Comment  string `json:"comment,omitempty"`  // Additional comment about the rollback
}
type ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStages struct {
	Group      *ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStagesGroup      `json:"group,omitempty"`      // The staged upgrade group
	Milestones *ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStagesMilestones `json:"milestones,omitempty"` // The Staged Upgrade Milestones for the stage
	Status     string                                                                        `json:"status,omitempty"`     // Current upgrade status of the group
}
type ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStagesGroup struct {
	Description string `json:"description,omitempty"` // Description of the Staged Upgrade Group
	ID          string `json:"id,omitempty"`          // Id of the Staged Upgrade Group
	Name        string `json:"name,omitempty"`        // Name of the Staged Upgrade Group
}
type ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStagesMilestones struct {
	CanceledAt   string `json:"canceledAt,omitempty"`   // Time that the group was canceled
	CompletedAt  string `json:"completedAt,omitempty"`  // Finish time for the group
	ScheduledFor string `json:"scheduledFor,omitempty"` // Scheduled start time for the group
	StartedAt    string `json:"startedAt,omitempty"`    // Start time for the group
}
type ResponseNetworksGetNetworkFirmwareUpgradesStagedGroups []ResponseItemNetworksGetNetworkFirmwareUpgradesStagedGroups // Array of ResponseNetworksGetNetworkFirmwareUpgradesStagedGroups
type ResponseItemNetworksGetNetworkFirmwareUpgradesStagedGroups struct {
	AssignedDevices *ResponseItemNetworksGetNetworkFirmwareUpgradesStagedGroupsAssignedDevices `json:"assignedDevices,omitempty"` // The devices and Switch Stacks assigned to the Group
	Description     string                                                                     `json:"description,omitempty"`     // Description of the Staged Upgrade Group
	GroupID         string                                                                     `json:"groupId,omitempty"`         // Id of staged upgrade group
	IsDefault       *bool                                                                      `json:"isDefault,omitempty"`       // Boolean indicating the default Group. Any device that does not have a group explicitly assigned will upgrade with this group
	Name            string                                                                     `json:"name,omitempty"`            // Name of the Staged Upgrade Group
}
type ResponseItemNetworksGetNetworkFirmwareUpgradesStagedGroupsAssignedDevices struct {
	Devices      *[]ResponseItemNetworksGetNetworkFirmwareUpgradesStagedGroupsAssignedDevicesDevices      `json:"devices,omitempty"`      // Data Array of Devices containing the name and serial
	SwitchStacks *[]ResponseItemNetworksGetNetworkFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks `json:"switchStacks,omitempty"` // Data Array of Switch Stacks containing the name and id
}
type ResponseItemNetworksGetNetworkFirmwareUpgradesStagedGroupsAssignedDevicesDevices struct {
	Name   string `json:"name,omitempty"`   // Name of the device
	Serial string `json:"serial,omitempty"` // Serial of the device
}
type ResponseItemNetworksGetNetworkFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks struct {
	ID   string `json:"id,omitempty"`   // ID of the Switch Stack
	Name string `json:"name,omitempty"` // Name of the Switch Stack
}
type ResponseNetworksCreateNetworkFirmwareUpgradesStagedGroup interface{}
type ResponseNetworksGetNetworkFirmwareUpgradesStagedGroup struct {
	AssignedDevices *ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupAssignedDevices `json:"assignedDevices,omitempty"` // The devices and Switch Stacks assigned to the Group
	Description     string                                                                `json:"description,omitempty"`     // Description of the Staged Upgrade Group
	GroupID         string                                                                `json:"groupId,omitempty"`         // Id of staged upgrade group
	IsDefault       *bool                                                                 `json:"isDefault,omitempty"`       // Boolean indicating the default Group. Any device that does not have a group explicitly assigned will upgrade with this group
	Name            string                                                                `json:"name,omitempty"`            // Name of the Staged Upgrade Group
}
type ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupAssignedDevices struct {
	Devices      *[]ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupAssignedDevicesDevices      `json:"devices,omitempty"`      // Data Array of Devices containing the name and serial
	SwitchStacks *[]ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupAssignedDevicesSwitchStacks `json:"switchStacks,omitempty"` // Data Array of Switch Stacks containing the name and id
}
type ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupAssignedDevicesDevices struct {
	Name   string `json:"name,omitempty"`   // Name of the device
	Serial string `json:"serial,omitempty"` // Serial of the device
}
type ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupAssignedDevicesSwitchStacks struct {
	ID   string `json:"id,omitempty"`   // ID of the Switch Stack
	Name string `json:"name,omitempty"` // Name of the Switch Stack
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesStagedGroup interface{}
type ResponseNetworksGetNetworkFirmwareUpgradesStagedStages []ResponseItemNetworksGetNetworkFirmwareUpgradesStagedStages // Array of ResponseNetworksGetNetworkFirmwareUpgradesStagedStages
type ResponseItemNetworksGetNetworkFirmwareUpgradesStagedStages struct {
	Group *ResponseItemNetworksGetNetworkFirmwareUpgradesStagedStagesGroup `json:"group,omitempty"` // The Staged Upgrade Group
}
type ResponseItemNetworksGetNetworkFirmwareUpgradesStagedStagesGroup struct {
	Description string `json:"description,omitempty"` // Description of the Staged Upgrade Group
	ID          string `json:"id,omitempty"`          // Id of the Staged Upgrade Group
	Name        string `json:"name,omitempty"`        // Name of the Staged Upgrade Group
}
type ResponseNetworksUpdateNetworkFirmwareUpgradesStagedStages []ResponseItemNetworksUpdateNetworkFirmwareUpgradesStagedStages // Array of ResponseNetworksUpdateNetworkFirmwareUpgradesStagedStages
type ResponseItemNetworksUpdateNetworkFirmwareUpgradesStagedStages struct {
	Group *ResponseItemNetworksUpdateNetworkFirmwareUpgradesStagedStagesGroup `json:"group,omitempty"` // The Staged Upgrade Group
}
type ResponseItemNetworksUpdateNetworkFirmwareUpgradesStagedStagesGroup struct {
	Description string `json:"description,omitempty"` // Description of the Staged Upgrade Group
	ID          string `json:"id,omitempty"`          // Id of the Staged Upgrade Group
	Name        string `json:"name,omitempty"`        // Name of the Staged Upgrade Group
}
type ResponseNetworksGetNetworkFloorPlans []ResponseItemNetworksGetNetworkFloorPlans // Array of ResponseNetworksGetNetworkFloorPlans
type ResponseItemNetworksGetNetworkFloorPlans struct {
	BottomLeftCorner  *ResponseItemNetworksGetNetworkFloorPlansBottomLeftCorner  `json:"bottomLeftCorner,omitempty"`  //
	BottomRightCorner *ResponseItemNetworksGetNetworkFloorPlansBottomRightCorner `json:"bottomRightCorner,omitempty"` //
	Center            *ResponseItemNetworksGetNetworkFloorPlansCenter            `json:"center,omitempty"`            //
	Devices           *[]ResponseItemNetworksGetNetworkFloorPlansDevices         `json:"devices,omitempty"`           //
	FloorPlanID       string                                                     `json:"floorPlanId,omitempty"`       //
	Height            *float64                                                   `json:"height,omitempty"`            //
	ImageExtension    string                                                     `json:"imageExtension,omitempty"`    //
	ImageMd5          string                                                     `json:"imageMd5,omitempty"`          //
	ImageURL          string                                                     `json:"imageUrl,omitempty"`          //
	ImageURLExpiresAt string                                                     `json:"imageUrlExpiresAt,omitempty"` //
	Name              string                                                     `json:"name,omitempty"`              //
	TopLeftCorner     *ResponseItemNetworksGetNetworkFloorPlansTopLeftCorner     `json:"topLeftCorner,omitempty"`     //
	TopRightCorner    *ResponseItemNetworksGetNetworkFloorPlansTopRightCorner    `json:"topRightCorner,omitempty"`    //
	Width             *int                                                       `json:"width,omitempty"`             //
}
type ResponseItemNetworksGetNetworkFloorPlansBottomLeftCorner struct {
	Lat *float64 `json:"lat,omitempty"` //
	Lng *float64 `json:"lng,omitempty"` //
}
type ResponseItemNetworksGetNetworkFloorPlansBottomRightCorner struct {
	Lat *float64 `json:"lat,omitempty"` //
	Lng *float64 `json:"lng,omitempty"` //
}
type ResponseItemNetworksGetNetworkFloorPlansCenter struct {
	Lat *float64 `json:"lat,omitempty"` //
	Lng *float64 `json:"lng,omitempty"` //
}
type ResponseItemNetworksGetNetworkFloorPlansDevices struct {
	Address        string                                                         `json:"address,omitempty"`        //
	BeaconIDParams *ResponseItemNetworksGetNetworkFloorPlansDevicesBeaconIDParams `json:"beaconIdParams,omitempty"` //
	Firmware       string                                                         `json:"firmware,omitempty"`       //
	FloorPlanID    string                                                         `json:"floorPlanId,omitempty"`    //
	LanIP          string                                                         `json:"lanIp,omitempty"`          //
	Lat            *float64                                                       `json:"lat,omitempty"`            //
	Lng            *float64                                                       `json:"lng,omitempty"`            //
	Mac            string                                                         `json:"mac,omitempty"`            //
	Model          string                                                         `json:"model,omitempty"`          //
	Name           string                                                         `json:"name,omitempty"`           //
	NetworkID      string                                                         `json:"networkId,omitempty"`      //
	Notes          string                                                         `json:"notes,omitempty"`          //
	Serial         string                                                         `json:"serial,omitempty"`         //
	Tags           []string                                                       `json:"tags,omitempty"`           //
}
type ResponseItemNetworksGetNetworkFloorPlansDevicesBeaconIDParams struct {
	Major *int   `json:"major,omitempty"` //
	Minor *int   `json:"minor,omitempty"` //
	UUID  string `json:"uuid,omitempty"`  //
}
type ResponseItemNetworksGetNetworkFloorPlansTopLeftCorner struct {
	Lat *float64 `json:"lat,omitempty"` //
	Lng *float64 `json:"lng,omitempty"` //
}
type ResponseItemNetworksGetNetworkFloorPlansTopRightCorner struct {
	Lat *float64 `json:"lat,omitempty"` //
	Lng *float64 `json:"lng,omitempty"` //
}
type ResponseNetworksCreateNetworkFloorPlan interface{}
type ResponseNetworksGetNetworkFloorPlan struct {
	BottomLeftCorner  *ResponseNetworksGetNetworkFloorPlanBottomLeftCorner  `json:"bottomLeftCorner,omitempty"`  //
	BottomRightCorner *ResponseNetworksGetNetworkFloorPlanBottomRightCorner `json:"bottomRightCorner,omitempty"` //
	Center            *ResponseNetworksGetNetworkFloorPlanCenter            `json:"center,omitempty"`            //
	Devices           *[]ResponseNetworksGetNetworkFloorPlanDevices         `json:"devices,omitempty"`           //
	FloorPlanID       string                                                `json:"floorPlanId,omitempty"`       //
	Height            *float64                                              `json:"height,omitempty"`            //
	ImageExtension    string                                                `json:"imageExtension,omitempty"`    //
	ImageMd5          string                                                `json:"imageMd5,omitempty"`          //
	ImageURL          string                                                `json:"imageUrl,omitempty"`          //
	ImageURLExpiresAt string                                                `json:"imageUrlExpiresAt,omitempty"` //
	Name              string                                                `json:"name,omitempty"`              //
	TopLeftCorner     *ResponseNetworksGetNetworkFloorPlanTopLeftCorner     `json:"topLeftCorner,omitempty"`     //
	TopRightCorner    *ResponseNetworksGetNetworkFloorPlanTopRightCorner    `json:"topRightCorner,omitempty"`    //
	Width             *int                                                  `json:"width,omitempty"`             //
}
type ResponseNetworksGetNetworkFloorPlanBottomLeftCorner struct {
	Lat *float64 `json:"lat,omitempty"` //
	Lng *float64 `json:"lng,omitempty"` //
}
type ResponseNetworksGetNetworkFloorPlanBottomRightCorner struct {
	Lat *float64 `json:"lat,omitempty"` //
	Lng *float64 `json:"lng,omitempty"` //
}
type ResponseNetworksGetNetworkFloorPlanCenter struct {
	Lat *float64 `json:"lat,omitempty"` //
	Lng *float64 `json:"lng,omitempty"` //
}
type ResponseNetworksGetNetworkFloorPlanDevices struct {
	Address        string                                                    `json:"address,omitempty"`        //
	BeaconIDParams *ResponseNetworksGetNetworkFloorPlanDevicesBeaconIDParams `json:"beaconIdParams,omitempty"` //
	Firmware       string                                                    `json:"firmware,omitempty"`       //
	FloorPlanID    string                                                    `json:"floorPlanId,omitempty"`    //
	LanIP          string                                                    `json:"lanIp,omitempty"`          //
	Lat            *float64                                                  `json:"lat,omitempty"`            //
	Lng            *float64                                                  `json:"lng,omitempty"`            //
	Mac            string                                                    `json:"mac,omitempty"`            //
	Model          string                                                    `json:"model,omitempty"`          //
	Name           string                                                    `json:"name,omitempty"`           //
	NetworkID      string                                                    `json:"networkId,omitempty"`      //
	Notes          string                                                    `json:"notes,omitempty"`          //
	Serial         string                                                    `json:"serial,omitempty"`         //
	Tags           []string                                                  `json:"tags,omitempty"`           //
}
type ResponseNetworksGetNetworkFloorPlanDevicesBeaconIDParams struct {
	Major *int   `json:"major,omitempty"` //
	Minor *int   `json:"minor,omitempty"` //
	UUID  string `json:"uuid,omitempty"`  //
}
type ResponseNetworksGetNetworkFloorPlanTopLeftCorner struct {
	Lat *float64 `json:"lat,omitempty"` //
	Lng *float64 `json:"lng,omitempty"` //
}
type ResponseNetworksGetNetworkFloorPlanTopRightCorner struct {
	Lat *float64 `json:"lat,omitempty"` //
	Lng *float64 `json:"lng,omitempty"` //
}
type ResponseNetworksUpdateNetworkFloorPlan interface{}
type ResponseNetworksGetNetworkGroupPolicies []ResponseItemNetworksGetNetworkGroupPolicies // Array of ResponseNetworksGetNetworkGroupPolicies
type ResponseItemNetworksGetNetworkGroupPolicies struct {
	Bandwidth                 *ResponseItemNetworksGetNetworkGroupPoliciesBandwidth                 `json:"bandwidth,omitempty"`                 //
	BonjourForwarding         *ResponseItemNetworksGetNetworkGroupPoliciesBonjourForwarding         `json:"bonjourForwarding,omitempty"`         //
	ContentFiltering          *ResponseItemNetworksGetNetworkGroupPoliciesContentFiltering          `json:"contentFiltering,omitempty"`          //
	FirewallAndTrafficShaping *ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShaping `json:"firewallAndTrafficShaping,omitempty"` //
	GroupPolicyID             string                                                                `json:"groupPolicyId,omitempty"`             //
	Name                      string                                                                `json:"name,omitempty"`                      //
	Scheduling                *ResponseItemNetworksGetNetworkGroupPoliciesScheduling                `json:"scheduling,omitempty"`                //
	SplashAuthSettings        string                                                                `json:"splashAuthSettings,omitempty"`        //
	VLANTagging               *ResponseItemNetworksGetNetworkGroupPoliciesVLANTagging               `json:"vlanTagging,omitempty"`               //
}
type ResponseItemNetworksGetNetworkGroupPoliciesBandwidth struct {
	BandwidthLimits *ResponseItemNetworksGetNetworkGroupPoliciesBandwidthBandwidthLimits `json:"bandwidthLimits,omitempty"` //
	Settings        string                                                               `json:"settings,omitempty"`        //
}
type ResponseItemNetworksGetNetworkGroupPoliciesBandwidthBandwidthLimits struct {
	LimitDown *int `json:"limitDown,omitempty"` //
	LimitUp   *int `json:"limitUp,omitempty"`   //
}
type ResponseItemNetworksGetNetworkGroupPoliciesBonjourForwarding struct {
	Rules    *[]ResponseItemNetworksGetNetworkGroupPoliciesBonjourForwardingRules `json:"rules,omitempty"`    //
	Settings string                                                               `json:"settings,omitempty"` //
}
type ResponseItemNetworksGetNetworkGroupPoliciesBonjourForwardingRules struct {
	Description string   `json:"description,omitempty"` //
	Services    []string `json:"services,omitempty"`    //
	VLANID      string   `json:"vlanId,omitempty"`      //
}
type ResponseItemNetworksGetNetworkGroupPoliciesContentFiltering struct {
	AllowedURLPatterns   *ResponseItemNetworksGetNetworkGroupPoliciesContentFilteringAllowedURLPatterns   `json:"allowedUrlPatterns,omitempty"`   //
	BlockedURLCategories *ResponseItemNetworksGetNetworkGroupPoliciesContentFilteringBlockedURLCategories `json:"blockedUrlCategories,omitempty"` //
	BlockedURLPatterns   *ResponseItemNetworksGetNetworkGroupPoliciesContentFilteringBlockedURLPatterns   `json:"blockedUrlPatterns,omitempty"`   //
}
type ResponseItemNetworksGetNetworkGroupPoliciesContentFilteringAllowedURLPatterns struct {
	Patterns []string `json:"patterns,omitempty"` //
	Settings string   `json:"settings,omitempty"` //
}
type ResponseItemNetworksGetNetworkGroupPoliciesContentFilteringBlockedURLCategories struct {
	Categories []string `json:"categories,omitempty"` //
	Settings   string   `json:"settings,omitempty"`   //
}
type ResponseItemNetworksGetNetworkGroupPoliciesContentFilteringBlockedURLPatterns struct {
	Patterns []string `json:"patterns,omitempty"` //
	Settings string   `json:"settings,omitempty"` //
}
type ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShaping struct {
	L3FirewallRules     *[]ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShapingL3FirewallRules     `json:"l3FirewallRules,omitempty"`     //
	L7FirewallRules     *[]ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShapingL7FirewallRules     `json:"l7FirewallRules,omitempty"`     //
	Settings            string                                                                                     `json:"settings,omitempty"`            //
	TrafficShapingRules *[]ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules `json:"trafficShapingRules,omitempty"` //
}
type ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShapingL3FirewallRules struct {
	Comment  string `json:"comment,omitempty"`  //
	DestCidr string `json:"destCidr,omitempty"` //
	DestPort string `json:"destPort,omitempty"` //
	Policy   string `json:"policy,omitempty"`   //
	Protocol string `json:"protocol,omitempty"` //
}
type ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShapingL7FirewallRules struct {
	Policy string `json:"policy,omitempty"` //
	Type   string `json:"type,omitempty"`   //
	Value  string `json:"value,omitempty"`  //
}
type ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules struct {
	Definitions              *[]ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShapingTrafficShapingRulesDefinitions            `json:"definitions,omitempty"`              //
	DscpTagValue             *int                                                                                                             `json:"dscpTagValue,omitempty"`             //
	PcpTagValue              *int                                                                                                             `json:"pcpTagValue,omitempty"`              //
	PerClientBandwidthLimits *ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimits `json:"perClientBandwidthLimits,omitempty"` //
}
type ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShapingTrafficShapingRulesDefinitions struct {
	Type  string `json:"type,omitempty"`  //
	Value string `json:"value,omitempty"` //
}
type ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimits struct {
	BandwidthLimits *ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimitsBandwidthLimits `json:"bandwidthLimits,omitempty"` //
	Settings        string                                                                                                                          `json:"settings,omitempty"`        //
}
type ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimitsBandwidthLimits struct {
	LimitDown *int `json:"limitDown,omitempty"` //
	LimitUp   *int `json:"limitUp,omitempty"`   //
}
type ResponseItemNetworksGetNetworkGroupPoliciesScheduling struct {
	Enabled   *bool                                                           `json:"enabled,omitempty"`   //
	Friday    *ResponseItemNetworksGetNetworkGroupPoliciesSchedulingFriday    `json:"friday,omitempty"`    //
	Monday    *ResponseItemNetworksGetNetworkGroupPoliciesSchedulingMonday    `json:"monday,omitempty"`    //
	Saturday  *ResponseItemNetworksGetNetworkGroupPoliciesSchedulingSaturday  `json:"saturday,omitempty"`  //
	Sunday    *ResponseItemNetworksGetNetworkGroupPoliciesSchedulingSunday    `json:"sunday,omitempty"`    //
	Thursday  *ResponseItemNetworksGetNetworkGroupPoliciesSchedulingThursday  `json:"thursday,omitempty"`  //
	Tuesday   *ResponseItemNetworksGetNetworkGroupPoliciesSchedulingTuesday   `json:"tuesday,omitempty"`   //
	Wednesday *ResponseItemNetworksGetNetworkGroupPoliciesSchedulingWednesday `json:"wednesday,omitempty"` //
}
type ResponseItemNetworksGetNetworkGroupPoliciesSchedulingFriday struct {
	Active *bool  `json:"active,omitempty"` //
	From   string `json:"from,omitempty"`   //
	To     string `json:"to,omitempty"`     //
}
type ResponseItemNetworksGetNetworkGroupPoliciesSchedulingMonday struct {
	Active *bool  `json:"active,omitempty"` //
	From   string `json:"from,omitempty"`   //
	To     string `json:"to,omitempty"`     //
}
type ResponseItemNetworksGetNetworkGroupPoliciesSchedulingSaturday struct {
	Active *bool  `json:"active,omitempty"` //
	From   string `json:"from,omitempty"`   //
	To     string `json:"to,omitempty"`     //
}
type ResponseItemNetworksGetNetworkGroupPoliciesSchedulingSunday struct {
	Active *bool  `json:"active,omitempty"` //
	From   string `json:"from,omitempty"`   //
	To     string `json:"to,omitempty"`     //
}
type ResponseItemNetworksGetNetworkGroupPoliciesSchedulingThursday struct {
	Active *bool  `json:"active,omitempty"` //
	From   string `json:"from,omitempty"`   //
	To     string `json:"to,omitempty"`     //
}
type ResponseItemNetworksGetNetworkGroupPoliciesSchedulingTuesday struct {
	Active *bool  `json:"active,omitempty"` //
	From   string `json:"from,omitempty"`   //
	To     string `json:"to,omitempty"`     //
}
type ResponseItemNetworksGetNetworkGroupPoliciesSchedulingWednesday struct {
	Active *bool  `json:"active,omitempty"` //
	From   string `json:"from,omitempty"`   //
	To     string `json:"to,omitempty"`     //
}
type ResponseItemNetworksGetNetworkGroupPoliciesVLANTagging struct {
	Settings string `json:"settings,omitempty"` //
	VLANID   string `json:"vlanId,omitempty"`   //
}
type ResponseNetworksCreateNetworkGroupPolicy interface{}
type ResponseNetworksGetNetworkGroupPolicy struct {
	Bandwidth                 *ResponseNetworksGetNetworkGroupPolicyBandwidth                 `json:"bandwidth,omitempty"`                 //
	BonjourForwarding         *ResponseNetworksGetNetworkGroupPolicyBonjourForwarding         `json:"bonjourForwarding,omitempty"`         //
	ContentFiltering          *ResponseNetworksGetNetworkGroupPolicyContentFiltering          `json:"contentFiltering,omitempty"`          //
	FirewallAndTrafficShaping *ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShaping `json:"firewallAndTrafficShaping,omitempty"` //
	GroupPolicyID             string                                                          `json:"groupPolicyId,omitempty"`             //
	Name                      string                                                          `json:"name,omitempty"`                      //
	Scheduling                *ResponseNetworksGetNetworkGroupPolicyScheduling                `json:"scheduling,omitempty"`                //
	SplashAuthSettings        string                                                          `json:"splashAuthSettings,omitempty"`        //
	VLANTagging               *ResponseNetworksGetNetworkGroupPolicyVLANTagging               `json:"vlanTagging,omitempty"`               //
}
type ResponseNetworksGetNetworkGroupPolicyBandwidth struct {
	BandwidthLimits *ResponseNetworksGetNetworkGroupPolicyBandwidthBandwidthLimits `json:"bandwidthLimits,omitempty"` //
	Settings        string                                                         `json:"settings,omitempty"`        //
}
type ResponseNetworksGetNetworkGroupPolicyBandwidthBandwidthLimits struct {
	LimitDown *int `json:"limitDown,omitempty"` //
	LimitUp   *int `json:"limitUp,omitempty"`   //
}
type ResponseNetworksGetNetworkGroupPolicyBonjourForwarding struct {
	Rules    *[]ResponseNetworksGetNetworkGroupPolicyBonjourForwardingRules `json:"rules,omitempty"`    //
	Settings string                                                         `json:"settings,omitempty"` //
}
type ResponseNetworksGetNetworkGroupPolicyBonjourForwardingRules struct {
	Description string   `json:"description,omitempty"` //
	Services    []string `json:"services,omitempty"`    //
	VLANID      string   `json:"vlanId,omitempty"`      //
}
type ResponseNetworksGetNetworkGroupPolicyContentFiltering struct {
	AllowedURLPatterns   *ResponseNetworksGetNetworkGroupPolicyContentFilteringAllowedURLPatterns   `json:"allowedUrlPatterns,omitempty"`   //
	BlockedURLCategories *ResponseNetworksGetNetworkGroupPolicyContentFilteringBlockedURLCategories `json:"blockedUrlCategories,omitempty"` //
	BlockedURLPatterns   *ResponseNetworksGetNetworkGroupPolicyContentFilteringBlockedURLPatterns   `json:"blockedUrlPatterns,omitempty"`   //
}
type ResponseNetworksGetNetworkGroupPolicyContentFilteringAllowedURLPatterns struct {
	Patterns []string `json:"patterns,omitempty"` //
	Settings string   `json:"settings,omitempty"` //
}
type ResponseNetworksGetNetworkGroupPolicyContentFilteringBlockedURLCategories struct {
	Categories []string `json:"categories,omitempty"` //
	Settings   string   `json:"settings,omitempty"`   //
}
type ResponseNetworksGetNetworkGroupPolicyContentFilteringBlockedURLPatterns struct {
	Patterns []string `json:"patterns,omitempty"` //
	Settings string   `json:"settings,omitempty"` //
}
type ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShaping struct {
	L3FirewallRules     *[]ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRules     `json:"l3FirewallRules,omitempty"`     //
	L7FirewallRules     *[]ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules     `json:"l7FirewallRules,omitempty"`     //
	Settings            string                                                                               `json:"settings,omitempty"`            //
	TrafficShapingRules *[]ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRules `json:"trafficShapingRules,omitempty"` //
}
type ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRules struct {
	Comment  string `json:"comment,omitempty"`  //
	DestCidr string `json:"destCidr,omitempty"` //
	DestPort string `json:"destPort,omitempty"` //
	Policy   string `json:"policy,omitempty"`   //
	Protocol string `json:"protocol,omitempty"` //
}
type ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules struct {
	Policy string `json:"policy,omitempty"` //
	Type   string `json:"type,omitempty"`   //
	Value  string `json:"value,omitempty"`  //
}
type ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRules struct {
	Definitions              *[]ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions            `json:"definitions,omitempty"`              //
	DscpTagValue             *int                                                                                                       `json:"dscpTagValue,omitempty"`             //
	PcpTagValue              *int                                                                                                       `json:"pcpTagValue,omitempty"`              //
	PerClientBandwidthLimits *ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimits `json:"perClientBandwidthLimits,omitempty"` //
}
type ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions struct {
	Type  string `json:"type,omitempty"`  //
	Value string `json:"value,omitempty"` //
}
type ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimits struct {
	BandwidthLimits *ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimitsBandwidthLimits `json:"bandwidthLimits,omitempty"` //
	Settings        string                                                                                                                    `json:"settings,omitempty"`        //
}
type ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimitsBandwidthLimits struct {
	LimitDown *int `json:"limitDown,omitempty"` //
	LimitUp   *int `json:"limitUp,omitempty"`   //
}
type ResponseNetworksGetNetworkGroupPolicyScheduling struct {
	Enabled   *bool                                                     `json:"enabled,omitempty"`   //
	Friday    *ResponseNetworksGetNetworkGroupPolicySchedulingFriday    `json:"friday,omitempty"`    //
	Monday    *ResponseNetworksGetNetworkGroupPolicySchedulingMonday    `json:"monday,omitempty"`    //
	Saturday  *ResponseNetworksGetNetworkGroupPolicySchedulingSaturday  `json:"saturday,omitempty"`  //
	Sunday    *ResponseNetworksGetNetworkGroupPolicySchedulingSunday    `json:"sunday,omitempty"`    //
	Thursday  *ResponseNetworksGetNetworkGroupPolicySchedulingThursday  `json:"thursday,omitempty"`  //
	Tuesday   *ResponseNetworksGetNetworkGroupPolicySchedulingTuesday   `json:"tuesday,omitempty"`   //
	Wednesday *ResponseNetworksGetNetworkGroupPolicySchedulingWednesday `json:"wednesday,omitempty"` //
}
type ResponseNetworksGetNetworkGroupPolicySchedulingFriday struct {
	Active *bool  `json:"active,omitempty"` //
	From   string `json:"from,omitempty"`   //
	To     string `json:"to,omitempty"`     //
}
type ResponseNetworksGetNetworkGroupPolicySchedulingMonday struct {
	Active *bool  `json:"active,omitempty"` //
	From   string `json:"from,omitempty"`   //
	To     string `json:"to,omitempty"`     //
}
type ResponseNetworksGetNetworkGroupPolicySchedulingSaturday struct {
	Active *bool  `json:"active,omitempty"` //
	From   string `json:"from,omitempty"`   //
	To     string `json:"to,omitempty"`     //
}
type ResponseNetworksGetNetworkGroupPolicySchedulingSunday struct {
	Active *bool  `json:"active,omitempty"` //
	From   string `json:"from,omitempty"`   //
	To     string `json:"to,omitempty"`     //
}
type ResponseNetworksGetNetworkGroupPolicySchedulingThursday struct {
	Active *bool  `json:"active,omitempty"` //
	From   string `json:"from,omitempty"`   //
	To     string `json:"to,omitempty"`     //
}
type ResponseNetworksGetNetworkGroupPolicySchedulingTuesday struct {
	Active *bool  `json:"active,omitempty"` //
	From   string `json:"from,omitempty"`   //
	To     string `json:"to,omitempty"`     //
}
type ResponseNetworksGetNetworkGroupPolicySchedulingWednesday struct {
	Active *bool  `json:"active,omitempty"` //
	From   string `json:"from,omitempty"`   //
	To     string `json:"to,omitempty"`     //
}
type ResponseNetworksGetNetworkGroupPolicyVLANTagging struct {
	Settings string `json:"settings,omitempty"` //
	VLANID   string `json:"vlanId,omitempty"`   //
}
type ResponseNetworksUpdateNetworkGroupPolicy interface{}
type ResponseNetworksGetNetworkHealthAlerts []ResponseItemNetworksGetNetworkHealthAlerts // Array of ResponseNetworksGetNetworkHealthAlerts
type ResponseItemNetworksGetNetworkHealthAlerts struct {
	Category string                                           `json:"category,omitempty"` // Category of the alert
	ID       string                                           `json:"id,omitempty"`       // Alert identifier. Value can be empty
	Scope    *ResponseItemNetworksGetNetworkHealthAlertsScope `json:"scope,omitempty"`    // The scope of the alert
	Severity string                                           `json:"severity,omitempty"` // Severity of the alert
	Type     string                                           `json:"type,omitempty"`     // Alert type
}
type ResponseItemNetworksGetNetworkHealthAlertsScope struct {
	Applications *[]ResponseItemNetworksGetNetworkHealthAlertsScopeApplications `json:"applications,omitempty"` // Applications related to the alert
	Devices      *[]ResponseItemNetworksGetNetworkHealthAlertsScopeDevices      `json:"devices,omitempty"`      // Devices related to the alert
	Peers        *[]ResponseItemNetworksGetNetworkHealthAlertsScopePeers        `json:"peers,omitempty"`        // Peers related to the alert
}
type ResponseItemNetworksGetNetworkHealthAlertsScopeApplications struct {
	Name string `json:"name,omitempty"` // Name of the application
	URL  string `json:"url,omitempty"`  // URL to the application
}
type ResponseItemNetworksGetNetworkHealthAlertsScopeDevices struct {
	Clients     *[]ResponseItemNetworksGetNetworkHealthAlertsScopeDevicesClients `json:"clients,omitempty"`     // Clients related to the device
	Lldp        *ResponseItemNetworksGetNetworkHealthAlertsScopeDevicesLldp      `json:"lldp,omitempty"`        // Lldp information
	Mac         string                                                           `json:"mac,omitempty"`         // The mac address of the device
	Name        string                                                           `json:"name,omitempty"`        // Name of the device
	ProductType string                                                           `json:"productType,omitempty"` // Product type of the device
	Serial      string                                                           `json:"serial,omitempty"`      // Serial number of the device
	URL         string                                                           `json:"url,omitempty"`         // URL to the device
}
type ResponseItemNetworksGetNetworkHealthAlertsScopeDevicesClients struct {
	Mac string `json:"mac,omitempty"` // Mac address of the client
}
type ResponseItemNetworksGetNetworkHealthAlertsScopeDevicesLldp struct {
	PortID string `json:"portId,omitempty"` // Port Id
}
type ResponseItemNetworksGetNetworkHealthAlertsScopePeers struct {
	Network *ResponseItemNetworksGetNetworkHealthAlertsScopePeersNetwork `json:"network,omitempty"` // Network of the peer
	URL     string                                                       `json:"url,omitempty"`     // URL to the peer
}
type ResponseItemNetworksGetNetworkHealthAlertsScopePeersNetwork struct {
	ID   string `json:"id,omitempty"`   // Id of the network
	Name string `json:"name,omitempty"` // Name of the network
}
type ResponseNetworksGetNetworkMerakiAuthUsers []ResponseItemNetworksGetNetworkMerakiAuthUsers // Array of ResponseNetworksGetNetworkMerakiAuthUsers
type ResponseItemNetworksGetNetworkMerakiAuthUsers struct {
	AccountType    string                                                         `json:"accountType,omitempty"`    // Authorization type for user.
	Authorizations *[]ResponseItemNetworksGetNetworkMerakiAuthUsersAuthorizations `json:"authorizations,omitempty"` // User authorization info
	CreatedAt      string                                                         `json:"createdAt,omitempty"`      // Creation time of the user
	Email          string                                                         `json:"email,omitempty"`          // Email address of the user
	ID             string                                                         `json:"id,omitempty"`             // Meraki auth user id
	IsAdmin        *bool                                                          `json:"isAdmin,omitempty"`        // Whether or not the user is a Dashboard administrator
	Name           string                                                         `json:"name,omitempty"`           // Name of the user
}
type ResponseItemNetworksGetNetworkMerakiAuthUsersAuthorizations struct {
	AuthorizedByEmail string `json:"authorizedByEmail,omitempty"` // User is authorized by the account email address
	AuthorizedByName  string `json:"authorizedByName,omitempty"`  // User is authorized by the account name
	AuthorizedZone    string `json:"authorizedZone,omitempty"`    // Authorized zone of the user
	ExpiresAt         string `json:"expiresAt,omitempty"`         // Authorization expiration time
	SSIDNumber        *int   `json:"ssidNumber,omitempty"`        // SSID number
}
type ResponseNetworksCreateNetworkMerakiAuthUser struct {
	AccountType    string                                                       `json:"accountType,omitempty"`    // Authorization type for user.
	Authorizations *[]ResponseNetworksCreateNetworkMerakiAuthUserAuthorizations `json:"authorizations,omitempty"` // User authorization info
	CreatedAt      string                                                       `json:"createdAt,omitempty"`      // Creation time of the user
	Email          string                                                       `json:"email,omitempty"`          // Email address of the user
	ID             string                                                       `json:"id,omitempty"`             // Meraki auth user id
	IsAdmin        *bool                                                        `json:"isAdmin,omitempty"`        // Whether or not the user is a Dashboard administrator
	Name           string                                                       `json:"name,omitempty"`           // Name of the user
}
type ResponseNetworksCreateNetworkMerakiAuthUserAuthorizations struct {
	AuthorizedByEmail string `json:"authorizedByEmail,omitempty"` // User is authorized by the account email address
	AuthorizedByName  string `json:"authorizedByName,omitempty"`  // User is authorized by the account name
	AuthorizedZone    string `json:"authorizedZone,omitempty"`    // Authorized zone of the user
	ExpiresAt         string `json:"expiresAt,omitempty"`         // Authorization expiration time
	SSIDNumber        *int   `json:"ssidNumber,omitempty"`        // SSID number
}
type ResponseNetworksGetNetworkMerakiAuthUser struct {
	AccountType    string                                                    `json:"accountType,omitempty"`    // Authorization type for user.
	Authorizations *[]ResponseNetworksGetNetworkMerakiAuthUserAuthorizations `json:"authorizations,omitempty"` // User authorization info
	CreatedAt      string                                                    `json:"createdAt,omitempty"`      // Creation time of the user
	Email          string                                                    `json:"email,omitempty"`          // Email address of the user
	ID             string                                                    `json:"id,omitempty"`             // Meraki auth user id
	IsAdmin        *bool                                                     `json:"isAdmin,omitempty"`        // Whether or not the user is a Dashboard administrator
	Name           string                                                    `json:"name,omitempty"`           // Name of the user
}
type ResponseNetworksGetNetworkMerakiAuthUserAuthorizations struct {
	AuthorizedByEmail string `json:"authorizedByEmail,omitempty"` // User is authorized by the account email address
	AuthorizedByName  string `json:"authorizedByName,omitempty"`  // User is authorized by the account name
	AuthorizedZone    string `json:"authorizedZone,omitempty"`    // Authorized zone of the user
	ExpiresAt         string `json:"expiresAt,omitempty"`         // Authorization expiration time
	SSIDNumber        *int   `json:"ssidNumber,omitempty"`        // SSID number
}
type ResponseNetworksUpdateNetworkMerakiAuthUser struct {
	AccountType    string                                                       `json:"accountType,omitempty"`    // Authorization type for user.
	Authorizations *[]ResponseNetworksUpdateNetworkMerakiAuthUserAuthorizations `json:"authorizations,omitempty"` // User authorization info
	CreatedAt      string                                                       `json:"createdAt,omitempty"`      // Creation time of the user
	Email          string                                                       `json:"email,omitempty"`          // Email address of the user
	ID             string                                                       `json:"id,omitempty"`             // Meraki auth user id
	IsAdmin        *bool                                                        `json:"isAdmin,omitempty"`        // Whether or not the user is a Dashboard administrator
	Name           string                                                       `json:"name,omitempty"`           // Name of the user
}
type ResponseNetworksUpdateNetworkMerakiAuthUserAuthorizations struct {
	AuthorizedByEmail string `json:"authorizedByEmail,omitempty"` // User is authorized by the account email address
	AuthorizedByName  string `json:"authorizedByName,omitempty"`  // User is authorized by the account name
	AuthorizedZone    string `json:"authorizedZone,omitempty"`    // Authorized zone of the user
	ExpiresAt         string `json:"expiresAt,omitempty"`         // Authorization expiration time
	SSIDNumber        *int   `json:"ssidNumber,omitempty"`        // SSID number
}
type ResponseNetworksGetNetworkMqttBrokers []ResponseItemNetworksGetNetworkMqttBrokers // Array of ResponseNetworksGetNetworkMqttBrokers
type ResponseItemNetworksGetNetworkMqttBrokers struct {
	Authentication *ResponseItemNetworksGetNetworkMqttBrokersAuthentication `json:"authentication,omitempty"` //
	Host           string                                                   `json:"host,omitempty"`           //
	ID             string                                                   `json:"id,omitempty"`             //
	Name           string                                                   `json:"name,omitempty"`           //
	Port           *int                                                     `json:"port,omitempty"`           //
	Security       *ResponseItemNetworksGetNetworkMqttBrokersSecurity       `json:"security,omitempty"`       //
}
type ResponseItemNetworksGetNetworkMqttBrokersAuthentication struct {
	Username string `json:"username,omitempty"` //
}
type ResponseItemNetworksGetNetworkMqttBrokersSecurity struct {
	Mode string                                                `json:"mode,omitempty"` //
	Tls  *ResponseItemNetworksGetNetworkMqttBrokersSecurityTls `json:"tls,omitempty"`  //
}
type ResponseItemNetworksGetNetworkMqttBrokersSecurityTls struct {
	HasCaCertificate *bool `json:"hasCaCertificate,omitempty"` //
	VerifyHostnames  *bool `json:"verifyHostnames,omitempty"`  //
}
type ResponseNetworksCreateNetworkMqttBroker interface{}
type ResponseNetworksGetNetworkMqttBroker struct {
	Authentication *ResponseNetworksGetNetworkMqttBrokerAuthentication `json:"authentication,omitempty"` //
	Host           string                                              `json:"host,omitempty"`           //
	ID             string                                              `json:"id,omitempty"`             //
	Name           string                                              `json:"name,omitempty"`           //
	Port           *int                                                `json:"port,omitempty"`           //
	Security       *ResponseNetworksGetNetworkMqttBrokerSecurity       `json:"security,omitempty"`       //
}
type ResponseNetworksGetNetworkMqttBrokerAuthentication struct {
	Username string `json:"username,omitempty"` //
}
type ResponseNetworksGetNetworkMqttBrokerSecurity struct {
	Mode string                                           `json:"mode,omitempty"` //
	Tls  *ResponseNetworksGetNetworkMqttBrokerSecurityTls `json:"tls,omitempty"`  //
}
type ResponseNetworksGetNetworkMqttBrokerSecurityTls struct {
	HasCaCertificate *bool `json:"hasCaCertificate,omitempty"` //
	VerifyHostnames  *bool `json:"verifyHostnames,omitempty"`  //
}
type ResponseNetworksUpdateNetworkMqttBroker interface{}
type ResponseNetworksGetNetworkNetflow struct {
	CollectorIP      string `json:"collectorIp,omitempty"`      //
	CollectorPort    *int   `json:"collectorPort,omitempty"`    //
	EtaDstPort       *int   `json:"etaDstPort,omitempty"`       //
	EtaEnabled       *bool  `json:"etaEnabled,omitempty"`       //
	ReportingEnabled *bool  `json:"reportingEnabled,omitempty"` //
}
type ResponseNetworksUpdateNetworkNetflow interface{}
type ResponseNetworksGetNetworkNetworkHealthChannelUtilization []ResponseItemNetworksGetNetworkNetworkHealthChannelUtilization // Array of ResponseNetworksGetNetworkNetworkHealthChannelUtilization
type ResponseItemNetworksGetNetworkNetworkHealthChannelUtilization struct {
	Model  string                                                                `json:"model,omitempty"`  //
	Serial string                                                                `json:"serial,omitempty"` //
	Tags   string                                                                `json:"tags,omitempty"`   //
	Wifi0  *[]ResponseItemNetworksGetNetworkNetworkHealthChannelUtilizationWifi0 `json:"wifi0,omitempty"`  //
	Wifi1  *[]ResponseItemNetworksGetNetworkNetworkHealthChannelUtilizationWifi1 `json:"wifi1,omitempty"`  //
}
type ResponseItemNetworksGetNetworkNetworkHealthChannelUtilizationWifi0 struct {
	EndTime             string   `json:"endTime,omitempty"`             //
	StartTime           string   `json:"startTime,omitempty"`           //
	Utilization80211    *int     `json:"utilization80211,omitempty"`    //
	UtilizationNon80211 *float64 `json:"utilizationNon80211,omitempty"` //
	UtilizationTotal    *float64 `json:"utilizationTotal,omitempty"`    //
}
type ResponseItemNetworksGetNetworkNetworkHealthChannelUtilizationWifi1 struct {
	EndTime             string   `json:"endTime,omitempty"`             //
	StartTime           string   `json:"startTime,omitempty"`           //
	Utilization80211    *int     `json:"utilization80211,omitempty"`    //
	UtilizationNon80211 *float64 `json:"utilizationNon80211,omitempty"` //
	UtilizationTotal    *float64 `json:"utilizationTotal,omitempty"`    //
}
type ResponseNetworksGetNetworkPiiPiiKeys struct {
	N1234 *ResponseNetworksGetNetworkPiiPiiKeysN1234 `json:"N_1234,omitempty"` //
}
type ResponseNetworksGetNetworkPiiPiiKeysN1234 struct {
	BluetoothMacs []string `json:"bluetoothMacs,omitempty"` //
	Emails        []string `json:"emails,omitempty"`        //
	Imeis         []string `json:"imeis,omitempty"`         //
	Macs          []string `json:"macs,omitempty"`          //
	Serials       []string `json:"serials,omitempty"`       //
	Usernames     []string `json:"usernames,omitempty"`     //
}
type ResponseNetworksGetNetworkPiiRequests []ResponseItemNetworksGetNetworkPiiRequests // Array of ResponseNetworksGetNetworkPiiRequests
type ResponseItemNetworksGetNetworkPiiRequests struct {
	CompletedAt      *int   `json:"completedAt,omitempty"`      //
	CreatedAt        *int   `json:"createdAt,omitempty"`        //
	Datasets         string `json:"datasets,omitempty"`         //
	ID               string `json:"id,omitempty"`               //
	Mac              string `json:"mac,omitempty"`              //
	NetworkID        string `json:"networkId,omitempty"`        //
	OrganizationWide *bool  `json:"organizationWide,omitempty"` //
	Status           string `json:"status,omitempty"`           //
	Type             string `json:"type,omitempty"`             //
}
type ResponseNetworksCreateNetworkPiiRequest interface{}
type ResponseNetworksGetNetworkPiiRequest struct {
	CompletedAt      *int   `json:"completedAt,omitempty"`      //
	CreatedAt        *int   `json:"createdAt,omitempty"`        //
	Datasets         string `json:"datasets,omitempty"`         //
	ID               string `json:"id,omitempty"`               //
	Mac              string `json:"mac,omitempty"`              //
	NetworkID        string `json:"networkId,omitempty"`        //
	OrganizationWide *bool  `json:"organizationWide,omitempty"` //
	Status           string `json:"status,omitempty"`           //
	Type             string `json:"type,omitempty"`             //
}
type ResponseNetworksGetNetworkPiiSmDevicesForKey struct {
	N1234 []string `json:"N_1234,omitempty"` //
}
type ResponseNetworksGetNetworkPiiSmOwnersForKey struct {
	N1234 []string `json:"N_1234,omitempty"` //
}
type ResponseNetworksGetNetworkPoliciesByClient []ResponseItemNetworksGetNetworkPoliciesByClient // Array of ResponseNetworksGetNetworkPoliciesByClient
type ResponseItemNetworksGetNetworkPoliciesByClient struct {
	Assigned *[]ResponseItemNetworksGetNetworkPoliciesByClientAssigned `json:"assigned,omitempty"` // Assigned policies
	ClientID string                                                    `json:"clientId,omitempty"` // ID of client
	Name     string                                                    `json:"name,omitempty"`     // Name of client
}
type ResponseItemNetworksGetNetworkPoliciesByClientAssigned struct {
	GroupPolicyID string                                                        `json:"groupPolicyId,omitempty"` // id of policy
	Name          string                                                        `json:"name,omitempty"`          // name of policy
	SSID          *[]ResponseItemNetworksGetNetworkPoliciesByClientAssignedSSID `json:"ssid,omitempty"`          // ssid
	Type          string                                                        `json:"type,omitempty"`          // type of policy
}
type ResponseItemNetworksGetNetworkPoliciesByClientAssignedSSID struct {
	SSIDNumber *int `json:"ssidNumber,omitempty"` // number of ssid
}
type ResponseNetworksGetNetworkSettings struct {
	ClientPrivacy           *ResponseNetworksGetNetworkSettingsClientPrivacy   `json:"clientPrivacy,omitempty"`           // Privacy settings
	Fips                    *ResponseNetworksGetNetworkSettingsFips            `json:"fips,omitempty"`                    // A hash of FIPS options applied to the Network
	LocalStatusPage         *ResponseNetworksGetNetworkSettingsLocalStatusPage `json:"localStatusPage,omitempty"`         // A hash of Local Status page(s)' authentication options applied to the Network.
	LocalStatusPageEnabled  *bool                                              `json:"localStatusPageEnabled,omitempty"`  // Enables / disables the local device status pages (<a target='_blank' href='http://my.meraki.com/'>my.meraki.com, </a><a target='_blank' href='http://ap.meraki.com/'>ap.meraki.com, </a><a target='_blank' href='http://switch.meraki.com/'>switch.meraki.com, </a><a target='_blank' href='http://wired.meraki.com/'>wired.meraki.com</a>). Optional (defaults to false)
	NamedVLANs              *ResponseNetworksGetNetworkSettingsNamedVLANs      `json:"namedVlans,omitempty"`              // A hash of Named VLANs options applied to the Network.
	RemoteStatusPageEnabled *bool                                              `json:"remoteStatusPageEnabled,omitempty"` // Enables / disables access to the device status page (<a target='_blank'>http://[device's LAN IP])</a>. Optional. Can only be set if localStatusPageEnabled is set to true
	SecurePort              *ResponseNetworksGetNetworkSettingsSecurePort      `json:"securePort,omitempty"`              // A hash of SecureConnect options applied to the Network.
}
type ResponseNetworksGetNetworkSettingsClientPrivacy struct {
	ExpireDataBefore    string `json:"expireDataBefore,omitempty"`    // The date to expire the data before
	ExpireDataOlderThan *int   `json:"expireDataOlderThan,omitempty"` // The number of days, weeks, or months in Epoch time to expire the data before
}
type ResponseNetworksGetNetworkSettingsFips struct {
	Enabled *bool `json:"enabled,omitempty"` // Enables / disables FIPS on the network.
}
type ResponseNetworksGetNetworkSettingsLocalStatusPage struct {
	Authentication *ResponseNetworksGetNetworkSettingsLocalStatusPageAuthentication `json:"authentication,omitempty"` // A hash of Local Status page(s)' authentication options applied to the Network.
}
type ResponseNetworksGetNetworkSettingsLocalStatusPageAuthentication struct {
	Enabled  *bool  `json:"enabled,omitempty"`  // Enables / disables the authentication on Local Status page(s).
	Username string `json:"username,omitempty"` // The username used for Local Status Page(s).
}
type ResponseNetworksGetNetworkSettingsNamedVLANs struct {
	Enabled *bool `json:"enabled,omitempty"` // Enables / disables Named VLANs on the Network.
}
type ResponseNetworksGetNetworkSettingsSecurePort struct {
	Enabled *bool `json:"enabled,omitempty"` // Enables / disables SecureConnect on the network. Optional.
}
type ResponseNetworksUpdateNetworkSettings struct {
	ClientPrivacy           *ResponseNetworksUpdateNetworkSettingsClientPrivacy   `json:"clientPrivacy,omitempty"`           // Privacy settings
	Fips                    *ResponseNetworksUpdateNetworkSettingsFips            `json:"fips,omitempty"`                    // A hash of FIPS options applied to the Network
	LocalStatusPage         *ResponseNetworksUpdateNetworkSettingsLocalStatusPage `json:"localStatusPage,omitempty"`         // A hash of Local Status page(s)' authentication options applied to the Network.
	LocalStatusPageEnabled  *bool                                                 `json:"localStatusPageEnabled,omitempty"`  // Enables / disables the local device status pages (<a target='_blank' href='http://my.meraki.com/'>my.meraki.com, </a><a target='_blank' href='http://ap.meraki.com/'>ap.meraki.com, </a><a target='_blank' href='http://switch.meraki.com/'>switch.meraki.com, </a><a target='_blank' href='http://wired.meraki.com/'>wired.meraki.com</a>). Optional (defaults to false)
	NamedVLANs              *ResponseNetworksUpdateNetworkSettingsNamedVLANs      `json:"namedVlans,omitempty"`              // A hash of Named VLANs options applied to the Network.
	RemoteStatusPageEnabled *bool                                                 `json:"remoteStatusPageEnabled,omitempty"` // Enables / disables access to the device status page (<a target='_blank'>http://[device's LAN IP])</a>. Optional. Can only be set if localStatusPageEnabled is set to true
	SecurePort              *ResponseNetworksUpdateNetworkSettingsSecurePort      `json:"securePort,omitempty"`              // A hash of SecureConnect options applied to the Network.
}
type ResponseNetworksUpdateNetworkSettingsClientPrivacy struct {
	ExpireDataBefore    string `json:"expireDataBefore,omitempty"`    // The date to expire the data before
	ExpireDataOlderThan *int   `json:"expireDataOlderThan,omitempty"` // The number of days, weeks, or months in Epoch time to expire the data before
}
type ResponseNetworksUpdateNetworkSettingsFips struct {
	Enabled *bool `json:"enabled,omitempty"` // Enables / disables FIPS on the network.
}
type ResponseNetworksUpdateNetworkSettingsLocalStatusPage struct {
	Authentication *ResponseNetworksUpdateNetworkSettingsLocalStatusPageAuthentication `json:"authentication,omitempty"` // A hash of Local Status page(s)' authentication options applied to the Network.
}
type ResponseNetworksUpdateNetworkSettingsLocalStatusPageAuthentication struct {
	Enabled  *bool  `json:"enabled,omitempty"`  // Enables / disables the authentication on Local Status page(s).
	Username string `json:"username,omitempty"` // The username used for Local Status Page(s).
}
type ResponseNetworksUpdateNetworkSettingsNamedVLANs struct {
	Enabled *bool `json:"enabled,omitempty"` // Enables / disables Named VLANs on the Network.
}
type ResponseNetworksUpdateNetworkSettingsSecurePort struct {
	Enabled *bool `json:"enabled,omitempty"` // Enables / disables SecureConnect on the network. Optional.
}
type ResponseNetworksGetNetworkSNMP struct {
	Access string                                 `json:"access,omitempty"` //
	Users  *[]ResponseNetworksGetNetworkSNMPUsers `json:"users,omitempty"`  //
}
type ResponseNetworksGetNetworkSNMPUsers struct {
	Passphrase string `json:"passphrase,omitempty"` //
	Username   string `json:"username,omitempty"`   //
}
type ResponseNetworksUpdateNetworkSNMP interface{}
type ResponseNetworksGetNetworkSplashLoginAttempts []ResponseItemNetworksGetNetworkSplashLoginAttempts // Array of ResponseNetworksGetNetworkSplashLoginAttempts
type ResponseItemNetworksGetNetworkSplashLoginAttempts struct {
	Authorization    string `json:"authorization,omitempty"`    //
	ClientID         string `json:"clientId,omitempty"`         //
	ClientMac        string `json:"clientMac,omitempty"`        //
	GatewayDeviceMac string `json:"gatewayDeviceMac,omitempty"` //
	Login            string `json:"login,omitempty"`            //
	LoginAt          string `json:"loginAt,omitempty"`          //
	Name             string `json:"name,omitempty"`             //
	SSID             string `json:"ssid,omitempty"`             //
}
type ResponseNetworksSplitNetwork struct {
	ResultingNetworks *[]ResponseNetworksSplitNetworkResultingNetworks `json:"resultingNetworks,omitempty"` // Networks after the split
}
type ResponseNetworksSplitNetworkResultingNetworks struct {
	EnrollmentString        string   `json:"enrollmentString,omitempty"`        // Enrollment string for the network
	ID                      string   `json:"id,omitempty"`                      // Network ID
	IsBoundToConfigTemplate *bool    `json:"isBoundToConfigTemplate,omitempty"` // If the network is bound to a config template
	Name                    string   `json:"name,omitempty"`                    // Network name
	Notes                   string   `json:"notes,omitempty"`                   // Notes for the network
	OrganizationID          string   `json:"organizationId,omitempty"`          // Organization ID
	ProductTypes            []string `json:"productTypes,omitempty"`            // List of the product types that the network supports
	Tags                    []string `json:"tags,omitempty"`                    // Network tags
	TimeZone                string   `json:"timeZone,omitempty"`                // Timezone of the network
	URL                     string   `json:"url,omitempty"`                     // URL to the network Dashboard UI
}
type ResponseNetworksGetNetworkSyslogServers struct {
	Servers *[]ResponseNetworksGetNetworkSyslogServersServers `json:"servers,omitempty"` // List of the syslog servers for this network
}
type ResponseNetworksGetNetworkSyslogServersServers struct {
	Host  string   `json:"host,omitempty"`  // The IP address of the syslog server
	Port  string   `json:"port,omitempty"`  // The port of the syslog server
	Roles []string `json:"roles,omitempty"` // A list of roles for the syslog server. Options (case-insensitive): 'Wireless event log', 'Appliance event log', 'Switch event log', 'Air Marshal events', 'Flows', 'URLs', 'IDS alerts', 'Security events'
}
type ResponseNetworksUpdateNetworkSyslogServers struct {
	Servers *[]ResponseNetworksUpdateNetworkSyslogServersServers `json:"servers,omitempty"` // List of the syslog servers for this network
}
type ResponseNetworksUpdateNetworkSyslogServersServers struct {
	Host  string   `json:"host,omitempty"`  // The IP address of the syslog server
	Port  string   `json:"port,omitempty"`  // The port of the syslog server
	Roles []string `json:"roles,omitempty"` // A list of roles for the syslog server. Options (case-insensitive): 'Wireless event log', 'Appliance event log', 'Switch event log', 'Air Marshal events', 'Flows', 'URLs', 'IDS alerts', 'Security events'
}
type ResponseNetworksGetNetworkTopologyLinkLayer struct {
	Errors []string                                            `json:"errors,omitempty"` //
	Links  *[]ResponseNetworksGetNetworkTopologyLinkLayerLinks `json:"links,omitempty"`  //
	Nodes  *[]ResponseNetworksGetNetworkTopologyLinkLayerNodes `json:"nodes,omitempty"`  //
}
type ResponseNetworksGetNetworkTopologyLinkLayerLinks struct {
	Ends           *[]ResponseNetworksGetNetworkTopologyLinkLayerLinksEnds `json:"ends,omitempty"`           //
	LastReportedAt string                                                  `json:"lastReportedAt,omitempty"` //
}
type ResponseNetworksGetNetworkTopologyLinkLayerLinksEnds struct {
	Device     *ResponseNetworksGetNetworkTopologyLinkLayerLinksEndsDevice     `json:"device,omitempty"`     //
	Discovered *ResponseNetworksGetNetworkTopologyLinkLayerLinksEndsDiscovered `json:"discovered,omitempty"` //
	Node       *ResponseNetworksGetNetworkTopologyLinkLayerLinksEndsNode       `json:"node,omitempty"`       //
}
type ResponseNetworksGetNetworkTopologyLinkLayerLinksEndsDevice struct {
	Name   string `json:"name,omitempty"`   //
	Serial string `json:"serial,omitempty"` //
}
type ResponseNetworksGetNetworkTopologyLinkLayerLinksEndsDiscovered struct {
	Cdp  *ResponseNetworksGetNetworkTopologyLinkLayerLinksEndsDiscoveredCdp  `json:"cdp,omitempty"`  //
	Lldp *ResponseNetworksGetNetworkTopologyLinkLayerLinksEndsDiscoveredLldp `json:"lldp,omitempty"` //
}
type ResponseNetworksGetNetworkTopologyLinkLayerLinksEndsDiscoveredCdp struct {
	NativeVLAN *int   `json:"nativeVlan,omitempty"` //
	PortID     string `json:"portId,omitempty"`     //
}
type ResponseNetworksGetNetworkTopologyLinkLayerLinksEndsDiscoveredLldp struct {
	PortDescription string `json:"portDescription,omitempty"` //
	PortID          string `json:"portId,omitempty"`          //
}
type ResponseNetworksGetNetworkTopologyLinkLayerLinksEndsNode struct {
	DerivedID string `json:"derivedId,omitempty"` //
	Type      string `json:"type,omitempty"`      //
}
type ResponseNetworksGetNetworkTopologyLinkLayerNodes struct {
	DerivedID  string                                                      `json:"derivedId,omitempty"`  //
	Discovered *ResponseNetworksGetNetworkTopologyLinkLayerNodesDiscovered `json:"discovered,omitempty"` //
	Mac        string                                                      `json:"mac,omitempty"`        //
	Root       *bool                                                       `json:"root,omitempty"`       //
	Type       string                                                      `json:"type,omitempty"`       //
}
type ResponseNetworksGetNetworkTopologyLinkLayerNodesDiscovered struct {
	Cdp  string                                                          `json:"cdp,omitempty"`  //
	Lldp *ResponseNetworksGetNetworkTopologyLinkLayerNodesDiscoveredLldp `json:"lldp,omitempty"` //
}
type ResponseNetworksGetNetworkTopologyLinkLayerNodesDiscoveredLldp struct {
	ChassisID          string   `json:"chassisId,omitempty"`          //
	ManagementAddress  string   `json:"managementAddress,omitempty"`  //
	SystemCapabilities []string `json:"systemCapabilities,omitempty"` //
	SystemDescription  string   `json:"systemDescription,omitempty"`  //
	SystemName         string   `json:"systemName,omitempty"`         //
}
type ResponseNetworksGetNetworkTraffic []ResponseItemNetworksGetNetworkTraffic // Array of ResponseNetworksGetNetworkTraffic
type ResponseItemNetworksGetNetworkTraffic struct {
	ActiveTime  *int   `json:"activeTime,omitempty"`  //
	Application string `json:"application,omitempty"` //
	Flows       *int   `json:"flows,omitempty"`       //
	NumClients  *int   `json:"numClients,omitempty"`  //
	Port        *int   `json:"port,omitempty"`        //
	Protocol    string `json:"protocol,omitempty"`    //
	Recv        *int   `json:"recv,omitempty"`        //
	Sent        *int   `json:"sent,omitempty"`        //
}
type ResponseNetworksGetNetworkTrafficAnalysis struct {
	CustomPieChartItems *[]ResponseNetworksGetNetworkTrafficAnalysisCustomPieChartItems `json:"customPieChartItems,omitempty"` //
	Mode                string                                                          `json:"mode,omitempty"`                //
}
type ResponseNetworksGetNetworkTrafficAnalysisCustomPieChartItems struct {
	Name  string `json:"name,omitempty"`  //
	Type  string `json:"type,omitempty"`  //
	Value string `json:"value,omitempty"` //
}
type ResponseNetworksUpdateNetworkTrafficAnalysis interface{}
type ResponseNetworksGetNetworkTrafficShapingApplicationCategories struct {
	ApplicationCategories *[]ResponseNetworksGetNetworkTrafficShapingApplicationCategoriesApplicationCategories `json:"applicationCategories,omitempty"` //
}
type ResponseNetworksGetNetworkTrafficShapingApplicationCategoriesApplicationCategories struct {
	Applications *[]ResponseNetworksGetNetworkTrafficShapingApplicationCategoriesApplicationCategoriesApplications `json:"applications,omitempty"` //
	ID           string                                                                                            `json:"id,omitempty"`           //
	Name         string                                                                                            `json:"name,omitempty"`         //
}
type ResponseNetworksGetNetworkTrafficShapingApplicationCategoriesApplicationCategoriesApplications struct {
	ID   string `json:"id,omitempty"`   //
	Name string `json:"name,omitempty"` //
}
type ResponseNetworksGetNetworkTrafficShapingDscpTaggingOptions []ResponseItemNetworksGetNetworkTrafficShapingDscpTaggingOptions // Array of ResponseNetworksGetNetworkTrafficShapingDscpTaggingOptions
type ResponseItemNetworksGetNetworkTrafficShapingDscpTaggingOptions struct {
	Description  string `json:"description,omitempty"`  //
	DscpTagValue *int   `json:"dscpTagValue,omitempty"` //
}
type ResponseNetworksUnbindNetwork struct {
	EnrollmentString        string   `json:"enrollmentString,omitempty"`        // Enrollment string for the network
	ID                      string   `json:"id,omitempty"`                      // Network ID
	IsBoundToConfigTemplate *bool    `json:"isBoundToConfigTemplate,omitempty"` // If the network is bound to a config template
	Name                    string   `json:"name,omitempty"`                    // Network name
	Notes                   string   `json:"notes,omitempty"`                   // Notes for the network
	OrganizationID          string   `json:"organizationId,omitempty"`          // Organization ID
	ProductTypes            []string `json:"productTypes,omitempty"`            // List of the product types that the network supports
	Tags                    []string `json:"tags,omitempty"`                    // Network tags
	TimeZone                string   `json:"timeZone,omitempty"`                // Timezone of the network
	URL                     string   `json:"url,omitempty"`                     // URL to the network Dashboard UI
}
type ResponseNetworksGetNetworkWebhooksHTTPServers []ResponseItemNetworksGetNetworkWebhooksHTTPServers // Array of ResponseNetworksGetNetworkWebhooksHttpServers
type ResponseItemNetworksGetNetworkWebhooksHTTPServers struct {
	ID              string                                                            `json:"id,omitempty"`              // A Base64 encoded ID.
	Name            string                                                            `json:"name,omitempty"`            // A name for easy reference to the HTTP server
	NetworkID       string                                                            `json:"networkId,omitempty"`       // A Meraki network ID.
	PayloadTemplate *ResponseItemNetworksGetNetworkWebhooksHTTPServersPayloadTemplate `json:"payloadTemplate,omitempty"` // The payload template to use when posting data to the HTTP server.
	URL             string                                                            `json:"url,omitempty"`             // The URL of the HTTP server.
}
type ResponseItemNetworksGetNetworkWebhooksHTTPServersPayloadTemplate struct {
	Name              string `json:"name,omitempty"`              // The name of the payload template.
	PayloadTemplateID string `json:"payloadTemplateId,omitempty"` // The ID of the payload template.
}
type ResponseNetworksCreateNetworkWebhooksHTTPServer struct {
	ID              string                                                          `json:"id,omitempty"`              // A Base64 encoded ID.
	Name            string                                                          `json:"name,omitempty"`            // A name for easy reference to the HTTP server
	NetworkID       string                                                          `json:"networkId,omitempty"`       // A Meraki network ID.
	PayloadTemplate *ResponseNetworksCreateNetworkWebhooksHTTPServerPayloadTemplate `json:"payloadTemplate,omitempty"` // The payload template to use when posting data to the HTTP server.
	URL             string                                                          `json:"url,omitempty"`             // The URL of the HTTP server.
}
type ResponseNetworksCreateNetworkWebhooksHTTPServerPayloadTemplate struct {
	Name              string `json:"name,omitempty"`              // The name of the payload template.
	PayloadTemplateID string `json:"payloadTemplateId,omitempty"` // The ID of the payload template.
}
type ResponseNetworksGetNetworkWebhooksHTTPServer struct {
	ID              string                                                       `json:"id,omitempty"`              // A Base64 encoded ID.
	Name            string                                                       `json:"name,omitempty"`            // A name for easy reference to the HTTP server
	NetworkID       string                                                       `json:"networkId,omitempty"`       // A Meraki network ID.
	PayloadTemplate *ResponseNetworksGetNetworkWebhooksHTTPServerPayloadTemplate `json:"payloadTemplate,omitempty"` // The payload template to use when posting data to the HTTP server.
	URL             string                                                       `json:"url,omitempty"`             // The URL of the HTTP server.
}
type ResponseNetworksGetNetworkWebhooksHTTPServerPayloadTemplate struct {
	Name              string `json:"name,omitempty"`              // The name of the payload template.
	PayloadTemplateID string `json:"payloadTemplateId,omitempty"` // The ID of the payload template.
}
type ResponseNetworksUpdateNetworkWebhooksHTTPServer struct {
	ID              string                                                          `json:"id,omitempty"`              // A Base64 encoded ID.
	Name            string                                                          `json:"name,omitempty"`            // A name for easy reference to the HTTP server
	NetworkID       string                                                          `json:"networkId,omitempty"`       // A Meraki network ID.
	PayloadTemplate *ResponseNetworksUpdateNetworkWebhooksHTTPServerPayloadTemplate `json:"payloadTemplate,omitempty"` // The payload template to use when posting data to the HTTP server.
	URL             string                                                          `json:"url,omitempty"`             // The URL of the HTTP server.
}
type ResponseNetworksUpdateNetworkWebhooksHTTPServerPayloadTemplate struct {
	Name              string `json:"name,omitempty"`              // The name of the payload template.
	PayloadTemplateID string `json:"payloadTemplateId,omitempty"` // The ID of the payload template.
}
type ResponseNetworksGetNetworkWebhooksPayloadTemplates []ResponseItemNetworksGetNetworkWebhooksPayloadTemplates // Array of ResponseNetworksGetNetworkWebhooksPayloadTemplates
type ResponseItemNetworksGetNetworkWebhooksPayloadTemplates struct {
	Body              string                                                           `json:"body,omitempty"`              // The body of the payload template, in liquid template
	Headers           *[]ResponseItemNetworksGetNetworkWebhooksPayloadTemplatesHeaders `json:"headers,omitempty"`           // The payload template headers, will be rendered as a key-value pair in the webhook.
	Name              string                                                           `json:"name,omitempty"`              // The name of the payload template
	PayloadTemplateID string                                                           `json:"payloadTemplateId,omitempty"` // Webhook payload template Id
	Sharing           *ResponseItemNetworksGetNetworkWebhooksPayloadTemplatesSharing   `json:"sharing,omitempty"`           // Information on which entities have access to the template
	Type              string                                                           `json:"type,omitempty"`              // The type of the payload template
}
type ResponseItemNetworksGetNetworkWebhooksPayloadTemplatesHeaders struct {
	Name     string `json:"name,omitempty"`     // The name of the header attribute
	Template string `json:"template,omitempty"` // The value returned in the header attribute, in liquid template
}
type ResponseItemNetworksGetNetworkWebhooksPayloadTemplatesSharing struct {
	ByNetwork *ResponseItemNetworksGetNetworkWebhooksPayloadTemplatesSharingByNetwork `json:"byNetwork,omitempty"` // Information on network access to the template
}
type ResponseItemNetworksGetNetworkWebhooksPayloadTemplatesSharingByNetwork struct {
	AdminsCanModify *bool `json:"adminsCanModify,omitempty"` // Indicates whether network admins may modify this template
}
type ResponseNetworksCreateNetworkWebhooksPayloadTemplate struct {
	Body              string                                                         `json:"body,omitempty"`              // The body of the payload template, in liquid template
	Headers           *[]ResponseNetworksCreateNetworkWebhooksPayloadTemplateHeaders `json:"headers,omitempty"`           // The payload template headers, will be rendered as a key-value pair in the webhook.
	Name              string                                                         `json:"name,omitempty"`              // The name of the payload template
	PayloadTemplateID string                                                         `json:"payloadTemplateId,omitempty"` // Webhook payload template Id
	Sharing           *ResponseNetworksCreateNetworkWebhooksPayloadTemplateSharing   `json:"sharing,omitempty"`           // Information on which entities have access to the template
	Type              string                                                         `json:"type,omitempty"`              // The type of the payload template
}
type ResponseNetworksCreateNetworkWebhooksPayloadTemplateHeaders struct {
	Name     string `json:"name,omitempty"`     // The name of the header attribute
	Template string `json:"template,omitempty"` // The value returned in the header attribute, in liquid template
}
type ResponseNetworksCreateNetworkWebhooksPayloadTemplateSharing struct {
	ByNetwork *ResponseNetworksCreateNetworkWebhooksPayloadTemplateSharingByNetwork `json:"byNetwork,omitempty"` // Information on network access to the template
}
type ResponseNetworksCreateNetworkWebhooksPayloadTemplateSharingByNetwork struct {
	AdminsCanModify *bool `json:"adminsCanModify,omitempty"` // Indicates whether network admins may modify this template
}
type ResponseNetworksGetNetworkWebhooksPayloadTemplate struct {
	Body              string                                                      `json:"body,omitempty"`              // The body of the payload template, in liquid template
	Headers           *[]ResponseNetworksGetNetworkWebhooksPayloadTemplateHeaders `json:"headers,omitempty"`           // The payload template headers, will be rendered as a key-value pair in the webhook.
	Name              string                                                      `json:"name,omitempty"`              // The name of the payload template
	PayloadTemplateID string                                                      `json:"payloadTemplateId,omitempty"` // Webhook payload template Id
	Sharing           *ResponseNetworksGetNetworkWebhooksPayloadTemplateSharing   `json:"sharing,omitempty"`           // Information on which entities have access to the template
	Type              string                                                      `json:"type,omitempty"`              // The type of the payload template
}
type ResponseNetworksGetNetworkWebhooksPayloadTemplateHeaders struct {
	Name     string `json:"name,omitempty"`     // The name of the header attribute
	Template string `json:"template,omitempty"` // The value returned in the header attribute, in liquid template
}
type ResponseNetworksGetNetworkWebhooksPayloadTemplateSharing struct {
	ByNetwork *ResponseNetworksGetNetworkWebhooksPayloadTemplateSharingByNetwork `json:"byNetwork,omitempty"` // Information on network access to the template
}
type ResponseNetworksGetNetworkWebhooksPayloadTemplateSharingByNetwork struct {
	AdminsCanModify *bool `json:"adminsCanModify,omitempty"` // Indicates whether network admins may modify this template
}
type ResponseNetworksUpdateNetworkWebhooksPayloadTemplate struct {
	Body              string                                                         `json:"body,omitempty"`              // The body of the payload template, in liquid template
	Headers           *[]ResponseNetworksUpdateNetworkWebhooksPayloadTemplateHeaders `json:"headers,omitempty"`           // The payload template headers, will be rendered as a key-value pair in the webhook.
	Name              string                                                         `json:"name,omitempty"`              // The name of the payload template
	PayloadTemplateID string                                                         `json:"payloadTemplateId,omitempty"` // Webhook payload template Id
	Sharing           *ResponseNetworksUpdateNetworkWebhooksPayloadTemplateSharing   `json:"sharing,omitempty"`           // Information on which entities have access to the template
	Type              string                                                         `json:"type,omitempty"`              // The type of the payload template
}
type ResponseNetworksUpdateNetworkWebhooksPayloadTemplateHeaders struct {
	Name     string `json:"name,omitempty"`     // The name of the header attribute
	Template string `json:"template,omitempty"` // The value returned in the header attribute, in liquid template
}
type ResponseNetworksUpdateNetworkWebhooksPayloadTemplateSharing struct {
	ByNetwork *ResponseNetworksUpdateNetworkWebhooksPayloadTemplateSharingByNetwork `json:"byNetwork,omitempty"` // Information on network access to the template
}
type ResponseNetworksUpdateNetworkWebhooksPayloadTemplateSharingByNetwork struct {
	AdminsCanModify *bool `json:"adminsCanModify,omitempty"` // Indicates whether network admins may modify this template
}
type ResponseNetworksCreateNetworkWebhooksWebhookTest struct {
	ID     string `json:"id,omitempty"`     // Webhook delivery identifier
	Status string `json:"status,omitempty"` // Current status of the webhook delivery
	URL    string `json:"url,omitempty"`    // URL where the webhook was delivered
}
type ResponseNetworksGetNetworkWebhooksWebhookTest struct {
	ID     string `json:"id,omitempty"`     // Webhook delivery identifier
	Status string `json:"status,omitempty"` // Current status of the webhook delivery
	URL    string `json:"url,omitempty"`    // URL where the webhook was delivered
}
type ResponseNetworksCombineOrganizationNetworks struct {
	ResultingNetwork *ResponseNetworksCombineOrganizationNetworksResultingNetwork `json:"resultingNetwork,omitempty"` // Network after the combination
}
type ResponseNetworksCombineOrganizationNetworksResultingNetwork struct {
	EnrollmentString        string   `json:"enrollmentString,omitempty"`        // Enrollment string for the network
	ID                      string   `json:"id,omitempty"`                      // Network ID
	IsBoundToConfigTemplate *bool    `json:"isBoundToConfigTemplate,omitempty"` // If the network is bound to a config template
	Name                    string   `json:"name,omitempty"`                    // Network name
	Notes                   string   `json:"notes,omitempty"`                   // Notes for the network
	OrganizationID          string   `json:"organizationId,omitempty"`          // Organization ID
	ProductTypes            []string `json:"productTypes,omitempty"`            // List of the product types that the network supports
	Tags                    []string `json:"tags,omitempty"`                    // Network tags
	TimeZone                string   `json:"timeZone,omitempty"`                // Timezone of the network
	URL                     string   `json:"url,omitempty"`                     // URL to the network Dashboard UI
}
type RequestNetworksUpdateNetwork struct {
	EnrollmentString string   `json:"enrollmentString,omitempty"` // A unique identifier which can be used for device enrollment or easy access through the Meraki SM Registration page or the Self Service Portal. Please note that changing this field may cause existing bookmarks to break.
	Name             string   `json:"name,omitempty"`             // The name of the network
	Notes            string   `json:"notes,omitempty"`            // Add any notes or additional information about this network here.
	Tags             []string `json:"tags,omitempty"`             // A list of tags to be applied to the network
	TimeZone         string   `json:"timeZone,omitempty"`         // The timezone of the network. For a list of allowed timezones, please see the 'TZ' column in the table in <a target='_blank' href='https://en.wikipedia.org/wiki/List_of_tz_database_time_zones'>this article.</a>
}
type RequestNetworksUpdateNetworkAlertsSettings struct {
	Alerts              *[]RequestNetworksUpdateNetworkAlertsSettingsAlerts            `json:"alerts,omitempty"`              // Alert-specific configuration for each type. Only alerts that pertain to the network can be updated.
	DefaultDestinations *RequestNetworksUpdateNetworkAlertsSettingsDefaultDestinations `json:"defaultDestinations,omitempty"` // The network-wide destinations for all alerts on the network.
}
type RequestNetworksUpdateNetworkAlertsSettingsAlerts struct {
	AlertDestinations *RequestNetworksUpdateNetworkAlertsSettingsAlertsAlertDestinations `json:"alertDestinations,omitempty"` // A hash of destinations for this specific alert
	Enabled           *bool                                                              `json:"enabled,omitempty"`           // A boolean depicting if the alert is turned on or off
	Filters           *RequestNetworksUpdateNetworkAlertsSettingsAlertsFilters           `json:"filters,omitempty"`           // A hash of specific configuration data for the alert. Only filters specific to the alert will be updated.
	Type              string                                                             `json:"type,omitempty"`              // The type of alert
}
type RequestNetworksUpdateNetworkAlertsSettingsAlertsAlertDestinations struct {
	AllAdmins     *bool    `json:"allAdmins,omitempty"`     // If true, then all network admins will receive emails for this alert
	Emails        []string `json:"emails,omitempty"`        // A list of emails that will receive information about the alert
	HTTPServerIDs []string `json:"httpServerIds,omitempty"` // A list of HTTP server IDs to send a Webhook to for this alert
	SNMP          *bool    `json:"snmp,omitempty"`          // If true, then an SNMP trap will be sent for this alert if there is an SNMP trap server configured for this network
}
type RequestNetworksUpdateNetworkAlertsSettingsAlertsFilters struct {
	Timeout   *int `json:"timeout,omitempty"`
	Threshold *int `json:"threshold,omitempty"`
	Period    *int `json:"period,omitempty"`
}
type RequestNetworksUpdateNetworkAlertsSettingsDefaultDestinations struct {
	AllAdmins     *bool    `json:"allAdmins,omitempty"`     // If true, then all network admins will receive emails.
	Emails        []string `json:"emails,omitempty"`        // A list of emails that will recieve the alert(s).
	HTTPServerIDs []string `json:"httpServerIds,omitempty"` // A list of HTTP server IDs to send a Webhook to
	SNMP          *bool    `json:"snmp,omitempty"`          // If true, then an SNMP trap will be sent if there is an SNMP trap server configured for this network.
}
type RequestNetworksBindNetwork struct {
	AutoBind         *bool  `json:"autoBind,omitempty"`         // Optional boolean indicating whether the network's switches should automatically bind to profiles of the same model. Defaults to false if left unspecified. This option only affects switch networks and switch templates. Auto-bind is not valid unless the switch template has at least one profile and has at most one profile per switch model.
	ConfigTemplateID string `json:"configTemplateId,omitempty"` // The ID of the template to which the network should be bound.
}
type RequestNetworksProvisionNetworkClients struct {
	Clients                     *[]RequestNetworksProvisionNetworkClientsClients                   `json:"clients,omitempty"`                     // The array of clients to provision
	DevicePolicy                string                                                             `json:"devicePolicy,omitempty"`                // The policy to apply to the specified client. Can be 'Group policy', 'Allowed', 'Blocked', 'Per connection' or 'Normal'. Required.
	GroupPolicyID               string                                                             `json:"groupPolicyId,omitempty"`               // The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.
	PoliciesBySecurityAppliance *RequestNetworksProvisionNetworkClientsPoliciesBySecurityAppliance `json:"policiesBySecurityAppliance,omitempty"` // An object, describing what the policy-connection association is for the security appliance. (Only relevant if the security appliance is actually within the network)
	PoliciesBySSID              *RequestNetworksProvisionNetworkClientsPoliciesBySSID              `json:"policiesBySsid,omitempty"`              // An object, describing the policy-connection associations for each active SSID within the network. Keys should be the number of enabled SSIDs, mapping to an object describing the client's policy
}
type RequestNetworksProvisionNetworkClientsClients struct {
	Mac  string `json:"mac,omitempty"`  // The MAC address of the client. Required.
	Name string `json:"name,omitempty"` // The display name for the client. Optional. Limited to 255 bytes.
}
type RequestNetworksProvisionNetworkClientsPoliciesBySecurityAppliance struct {
	DevicePolicy string `json:"devicePolicy,omitempty"` // The policy to apply to the specified client. Can be 'Allowed', 'Blocked' or 'Normal'. Required.
}
type RequestNetworksProvisionNetworkClientsPoliciesBySSID struct {
	Status0  *RequestNetworksProvisionNetworkClientsPoliciesBySSID0  `json:"0,omitempty"`  // The number for the SSID
	Status1  *RequestNetworksProvisionNetworkClientsPoliciesBySSID1  `json:"1,omitempty"`  // The number for the SSID
	Status10 *RequestNetworksProvisionNetworkClientsPoliciesBySSID10 `json:"10,omitempty"` // The number for the SSID
	Status11 *RequestNetworksProvisionNetworkClientsPoliciesBySSID11 `json:"11,omitempty"` // The number for the SSID
	Status12 *RequestNetworksProvisionNetworkClientsPoliciesBySSID12 `json:"12,omitempty"` // The number for the SSID
	Status13 *RequestNetworksProvisionNetworkClientsPoliciesBySSID13 `json:"13,omitempty"` // The number for the SSID
	Status14 *RequestNetworksProvisionNetworkClientsPoliciesBySSID14 `json:"14,omitempty"` // The number for the SSID
	Status2  *RequestNetworksProvisionNetworkClientsPoliciesBySSID2  `json:"2,omitempty"`  // The number for the SSID
	Status3  *RequestNetworksProvisionNetworkClientsPoliciesBySSID3  `json:"3,omitempty"`  // The number for the SSID
	Status4  *RequestNetworksProvisionNetworkClientsPoliciesBySSID4  `json:"4,omitempty"`  // The number for the SSID
	Status5  *RequestNetworksProvisionNetworkClientsPoliciesBySSID5  `json:"5,omitempty"`  // The number for the SSID
	Status6  *RequestNetworksProvisionNetworkClientsPoliciesBySSID6  `json:"6,omitempty"`  // The number for the SSID
	Status7  *RequestNetworksProvisionNetworkClientsPoliciesBySSID7  `json:"7,omitempty"`  // The number for the SSID
	Status8  *RequestNetworksProvisionNetworkClientsPoliciesBySSID8  `json:"8,omitempty"`  // The number for the SSID
	Status9  *RequestNetworksProvisionNetworkClientsPoliciesBySSID9  `json:"9,omitempty"`  // The number for the SSID
}
type RequestNetworksProvisionNetworkClientsPoliciesBySSID0 struct {
	DevicePolicy  string `json:"devicePolicy,omitempty"`  // The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
	GroupPolicyID string `json:"groupPolicyId,omitempty"` // The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.
}
type RequestNetworksProvisionNetworkClientsPoliciesBySSID1 struct {
	DevicePolicy  string `json:"devicePolicy,omitempty"`  // The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
	GroupPolicyID string `json:"groupPolicyId,omitempty"` // The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.
}
type RequestNetworksProvisionNetworkClientsPoliciesBySSID10 struct {
	DevicePolicy  string `json:"devicePolicy,omitempty"`  // The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
	GroupPolicyID string `json:"groupPolicyId,omitempty"` // The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.
}
type RequestNetworksProvisionNetworkClientsPoliciesBySSID11 struct {
	DevicePolicy  string `json:"devicePolicy,omitempty"`  // The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
	GroupPolicyID string `json:"groupPolicyId,omitempty"` // The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.
}
type RequestNetworksProvisionNetworkClientsPoliciesBySSID12 struct {
	DevicePolicy  string `json:"devicePolicy,omitempty"`  // The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
	GroupPolicyID string `json:"groupPolicyId,omitempty"` // The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.
}
type RequestNetworksProvisionNetworkClientsPoliciesBySSID13 struct {
	DevicePolicy  string `json:"devicePolicy,omitempty"`  // The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
	GroupPolicyID string `json:"groupPolicyId,omitempty"` // The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.
}
type RequestNetworksProvisionNetworkClientsPoliciesBySSID14 struct {
	DevicePolicy  string `json:"devicePolicy,omitempty"`  // The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
	GroupPolicyID string `json:"groupPolicyId,omitempty"` // The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.
}
type RequestNetworksProvisionNetworkClientsPoliciesBySSID2 struct {
	DevicePolicy  string `json:"devicePolicy,omitempty"`  // The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
	GroupPolicyID string `json:"groupPolicyId,omitempty"` // The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.
}
type RequestNetworksProvisionNetworkClientsPoliciesBySSID3 struct {
	DevicePolicy  string `json:"devicePolicy,omitempty"`  // The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
	GroupPolicyID string `json:"groupPolicyId,omitempty"` // The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.
}
type RequestNetworksProvisionNetworkClientsPoliciesBySSID4 struct {
	DevicePolicy  string `json:"devicePolicy,omitempty"`  // The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
	GroupPolicyID string `json:"groupPolicyId,omitempty"` // The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.
}
type RequestNetworksProvisionNetworkClientsPoliciesBySSID5 struct {
	DevicePolicy  string `json:"devicePolicy,omitempty"`  // The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
	GroupPolicyID string `json:"groupPolicyId,omitempty"` // The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.
}
type RequestNetworksProvisionNetworkClientsPoliciesBySSID6 struct {
	DevicePolicy  string `json:"devicePolicy,omitempty"`  // The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
	GroupPolicyID string `json:"groupPolicyId,omitempty"` // The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.
}
type RequestNetworksProvisionNetworkClientsPoliciesBySSID7 struct {
	DevicePolicy  string `json:"devicePolicy,omitempty"`  // The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
	GroupPolicyID string `json:"groupPolicyId,omitempty"` // The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.
}
type RequestNetworksProvisionNetworkClientsPoliciesBySSID8 struct {
	DevicePolicy  string `json:"devicePolicy,omitempty"`  // The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
	GroupPolicyID string `json:"groupPolicyId,omitempty"` // The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.
}
type RequestNetworksProvisionNetworkClientsPoliciesBySSID9 struct {
	DevicePolicy  string `json:"devicePolicy,omitempty"`  // The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
	GroupPolicyID string `json:"groupPolicyId,omitempty"` // The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.
}
type RequestNetworksUpdateNetworkClientPolicy struct {
	DevicePolicy  string `json:"devicePolicy,omitempty"`  // The policy to assign. Can be 'Whitelisted', 'Blocked', 'Normal' or 'Group policy'. Required.
	GroupPolicyID string `json:"groupPolicyId,omitempty"` // [optional] If 'devicePolicy' is set to 'Group policy' this param is used to specify the group policy ID.
}
type RequestNetworksUpdateNetworkClientSplashAuthorizationStatus struct {
	SSIDs *RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs `json:"ssids,omitempty"` // The target SSIDs. Each SSID must be enabled and must have Click-through splash enabled. For each SSID where isAuthorized is true, the expiration time will automatically be set according to the SSID's splash frequency. Not all networks support configuring all SSIDs
}
type RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs struct {
	Status0  *RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs0  `json:"0,omitempty"`  // Splash authorization for SSID 0
	Status1  *RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs1  `json:"1,omitempty"`  // Splash authorization for SSID 1
	Status10 *RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs10 `json:"10,omitempty"` // Splash authorization for SSID 10
	Status11 *RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs11 `json:"11,omitempty"` // Splash authorization for SSID 11
	Status12 *RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs12 `json:"12,omitempty"` // Splash authorization for SSID 12
	Status13 *RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs13 `json:"13,omitempty"` // Splash authorization for SSID 13
	Status14 *RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs14 `json:"14,omitempty"` // Splash authorization for SSID 14
	Status2  *RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs2  `json:"2,omitempty"`  // Splash authorization for SSID 2
	Status3  *RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs3  `json:"3,omitempty"`  // Splash authorization for SSID 3
	Status4  *RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs4  `json:"4,omitempty"`  // Splash authorization for SSID 4
	Status5  *RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs5  `json:"5,omitempty"`  // Splash authorization for SSID 5
	Status6  *RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs6  `json:"6,omitempty"`  // Splash authorization for SSID 6
	Status7  *RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs7  `json:"7,omitempty"`  // Splash authorization for SSID 7
	Status8  *RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs8  `json:"8,omitempty"`  // Splash authorization for SSID 8
	Status9  *RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs9  `json:"9,omitempty"`  // Splash authorization for SSID 9
}
type RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs0 struct {
	IsAuthorized *bool `json:"isAuthorized,omitempty"` // New authorization status for the SSID (true, false).
}
type RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs1 struct {
	IsAuthorized *bool `json:"isAuthorized,omitempty"` // New authorization status for the SSID (true, false).
}
type RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs10 struct {
	IsAuthorized *bool `json:"isAuthorized,omitempty"` // New authorization status for the SSID (true, false).
}
type RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs11 struct {
	IsAuthorized *bool `json:"isAuthorized,omitempty"` // New authorization status for the SSID (true, false).
}
type RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs12 struct {
	IsAuthorized *bool `json:"isAuthorized,omitempty"` // New authorization status for the SSID (true, false).
}
type RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs13 struct {
	IsAuthorized *bool `json:"isAuthorized,omitempty"` // New authorization status for the SSID (true, false).
}
type RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs14 struct {
	IsAuthorized *bool `json:"isAuthorized,omitempty"` // New authorization status for the SSID (true, false).
}
type RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs2 struct {
	IsAuthorized *bool `json:"isAuthorized,omitempty"` // New authorization status for the SSID (true, false).
}
type RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs3 struct {
	IsAuthorized *bool `json:"isAuthorized,omitempty"` // New authorization status for the SSID (true, false).
}
type RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs4 struct {
	IsAuthorized *bool `json:"isAuthorized,omitempty"` // New authorization status for the SSID (true, false).
}
type RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs5 struct {
	IsAuthorized *bool `json:"isAuthorized,omitempty"` // New authorization status for the SSID (true, false).
}
type RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs6 struct {
	IsAuthorized *bool `json:"isAuthorized,omitempty"` // New authorization status for the SSID (true, false).
}
type RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs7 struct {
	IsAuthorized *bool `json:"isAuthorized,omitempty"` // New authorization status for the SSID (true, false).
}
type RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs8 struct {
	IsAuthorized *bool `json:"isAuthorized,omitempty"` // New authorization status for the SSID (true, false).
}
type RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs9 struct {
	IsAuthorized *bool `json:"isAuthorized,omitempty"` // New authorization status for the SSID (true, false).
}
type RequestNetworksClaimNetworkDevices struct {
	Serials []string `json:"serials"` // A list of serials of devices to claim
}
type RequestNetworksVmxNetworkDevicesClaim struct {
	Size string `json:"size,omitempty"` // The size of the vMX you claim. It can be one of: small, medium, large, 100
}
type RequestNetworksRemoveNetworkDevices struct {
	Serial string `json:"serial,omitempty"` // The serial of a device
}
type RequestNetworksUpdateNetworkFirmwareUpgrades struct {
	Products      *RequestNetworksUpdateNetworkFirmwareUpgradesProducts      `json:"products,omitempty"`      // Contains information about the network to update
	Timezone      string                                                     `json:"timezone,omitempty"`      // The timezone for the network
	UpgradeWindow *RequestNetworksUpdateNetworkFirmwareUpgradesUpgradeWindow `json:"upgradeWindow,omitempty"` // Upgrade window for devices in network
}
type RequestNetworksUpdateNetworkFirmwareUpgradesProducts struct {
	Appliance       *RequestNetworksUpdateNetworkFirmwareUpgradesProductsAppliance       `json:"appliance,omitempty"`       // The network device to be updated
	Camera          *RequestNetworksUpdateNetworkFirmwareUpgradesProductsCamera          `json:"camera,omitempty"`          // The network device to be updated
	CellularGateway *RequestNetworksUpdateNetworkFirmwareUpgradesProductsCellularGateway `json:"cellularGateway,omitempty"` // The network device to be updated
	Sensor          *RequestNetworksUpdateNetworkFirmwareUpgradesProductsSensor          `json:"sensor,omitempty"`          // The network device to be updated
	Switch          *RequestNetworksUpdateNetworkFirmwareUpgradesProductsSwitch          `json:"switch,omitempty"`          // The network device to be updated
	Wireless        *RequestNetworksUpdateNetworkFirmwareUpgradesProductsWireless        `json:"wireless,omitempty"`        // The network device to be updated
}
type RequestNetworksUpdateNetworkFirmwareUpgradesProductsAppliance struct {
	NextUpgrade                  *RequestNetworksUpdateNetworkFirmwareUpgradesProductsApplianceNextUpgrade `json:"nextUpgrade,omitempty"`                  // The pending firmware upgrade if it exists
	ParticipateInNextBetaRelease *bool                                                                     `json:"participateInNextBetaRelease,omitempty"` // Whether or not the network wants beta firmware
}
type RequestNetworksUpdateNetworkFirmwareUpgradesProductsApplianceNextUpgrade struct {
	Time      string                                                                             `json:"time,omitempty"`      // The time of the last successful upgrade
	ToVersion *RequestNetworksUpdateNetworkFirmwareUpgradesProductsApplianceNextUpgradeToVersion `json:"toVersion,omitempty"` // The version to be updated to
}
type RequestNetworksUpdateNetworkFirmwareUpgradesProductsApplianceNextUpgradeToVersion struct {
	ID string `json:"id,omitempty"` // The version ID
}
type RequestNetworksUpdateNetworkFirmwareUpgradesProductsCamera struct {
	NextUpgrade                  *RequestNetworksUpdateNetworkFirmwareUpgradesProductsCameraNextUpgrade `json:"nextUpgrade,omitempty"`                  // The pending firmware upgrade if it exists
	ParticipateInNextBetaRelease *bool                                                                  `json:"participateInNextBetaRelease,omitempty"` // Whether or not the network wants beta firmware
}
type RequestNetworksUpdateNetworkFirmwareUpgradesProductsCameraNextUpgrade struct {
	Time      string                                                                          `json:"time,omitempty"`      // The time of the last successful upgrade
	ToVersion *RequestNetworksUpdateNetworkFirmwareUpgradesProductsCameraNextUpgradeToVersion `json:"toVersion,omitempty"` // The version to be updated to
}
type RequestNetworksUpdateNetworkFirmwareUpgradesProductsCameraNextUpgradeToVersion struct {
	ID string `json:"id,omitempty"` // The version ID
}
type RequestNetworksUpdateNetworkFirmwareUpgradesProductsCellularGateway struct {
	NextUpgrade                  *RequestNetworksUpdateNetworkFirmwareUpgradesProductsCellularGatewayNextUpgrade `json:"nextUpgrade,omitempty"`                  // The pending firmware upgrade if it exists
	ParticipateInNextBetaRelease *bool                                                                           `json:"participateInNextBetaRelease,omitempty"` // Whether or not the network wants beta firmware
}
type RequestNetworksUpdateNetworkFirmwareUpgradesProductsCellularGatewayNextUpgrade struct {
	Time      string                                                                                   `json:"time,omitempty"`      // The time of the last successful upgrade
	ToVersion *RequestNetworksUpdateNetworkFirmwareUpgradesProductsCellularGatewayNextUpgradeToVersion `json:"toVersion,omitempty"` // The version to be updated to
}
type RequestNetworksUpdateNetworkFirmwareUpgradesProductsCellularGatewayNextUpgradeToVersion struct {
	ID string `json:"id,omitempty"` // The version ID
}
type RequestNetworksUpdateNetworkFirmwareUpgradesProductsSensor struct {
	NextUpgrade                  *RequestNetworksUpdateNetworkFirmwareUpgradesProductsSensorNextUpgrade `json:"nextUpgrade,omitempty"`                  // The pending firmware upgrade if it exists
	ParticipateInNextBetaRelease *bool                                                                  `json:"participateInNextBetaRelease,omitempty"` // Whether or not the network wants beta firmware
}
type RequestNetworksUpdateNetworkFirmwareUpgradesProductsSensorNextUpgrade struct {
	Time      string                                                                          `json:"time,omitempty"`      // The time of the last successful upgrade
	ToVersion *RequestNetworksUpdateNetworkFirmwareUpgradesProductsSensorNextUpgradeToVersion `json:"toVersion,omitempty"` // The version to be updated to
}
type RequestNetworksUpdateNetworkFirmwareUpgradesProductsSensorNextUpgradeToVersion struct {
	ID string `json:"id,omitempty"` // The version ID
}
type RequestNetworksUpdateNetworkFirmwareUpgradesProductsSwitch struct {
	NextUpgrade                  *RequestNetworksUpdateNetworkFirmwareUpgradesProductsSwitchNextUpgrade `json:"nextUpgrade,omitempty"`                  // The pending firmware upgrade if it exists
	ParticipateInNextBetaRelease *bool                                                                  `json:"participateInNextBetaRelease,omitempty"` // Whether or not the network wants beta firmware
}
type RequestNetworksUpdateNetworkFirmwareUpgradesProductsSwitchNextUpgrade struct {
	Time      string                                                                          `json:"time,omitempty"`      // The time of the last successful upgrade
	ToVersion *RequestNetworksUpdateNetworkFirmwareUpgradesProductsSwitchNextUpgradeToVersion `json:"toVersion,omitempty"` // The version to be updated to
}
type RequestNetworksUpdateNetworkFirmwareUpgradesProductsSwitchNextUpgradeToVersion struct {
	ID string `json:"id,omitempty"` // The version ID
}
type RequestNetworksUpdateNetworkFirmwareUpgradesProductsWireless struct {
	NextUpgrade                  *RequestNetworksUpdateNetworkFirmwareUpgradesProductsWirelessNextUpgrade `json:"nextUpgrade,omitempty"`                  // The pending firmware upgrade if it exists
	ParticipateInNextBetaRelease *bool                                                                    `json:"participateInNextBetaRelease,omitempty"` // Whether or not the network wants beta firmware
}
type RequestNetworksUpdateNetworkFirmwareUpgradesProductsWirelessNextUpgrade struct {
	Time      string                                                                            `json:"time,omitempty"`      // The time of the last successful upgrade
	ToVersion *RequestNetworksUpdateNetworkFirmwareUpgradesProductsWirelessNextUpgradeToVersion `json:"toVersion,omitempty"` // The version to be updated to
}
type RequestNetworksUpdateNetworkFirmwareUpgradesProductsWirelessNextUpgradeToVersion struct {
	ID string `json:"id,omitempty"` // The version ID
}
type RequestNetworksUpdateNetworkFirmwareUpgradesUpgradeWindow struct {
	DayOfWeek string `json:"dayOfWeek,omitempty"` // Day of the week
	HourOfDay string `json:"hourOfDay,omitempty"` // Hour of the day
}
type RequestNetworksCreateNetworkFirmwareUpgradesRollback struct {
	Product   string                                                         `json:"product,omitempty"`   // Product type to rollback (if the network is a combined network)
	Reasons   *[]RequestNetworksCreateNetworkFirmwareUpgradesRollbackReasons `json:"reasons,omitempty"`   // Reasons for the rollback
	Time      string                                                         `json:"time,omitempty"`      // Scheduled time for the rollback
	ToVersion *RequestNetworksCreateNetworkFirmwareUpgradesRollbackToVersion `json:"toVersion,omitempty"` // Version to downgrade to (if the network has firmware flexibility)
}
type RequestNetworksCreateNetworkFirmwareUpgradesRollbackReasons struct {
	Category string `json:"category,omitempty"` // Reason for the rollback
	Comment  string `json:"comment,omitempty"`  // Additional comment about the rollback
}
type RequestNetworksCreateNetworkFirmwareUpgradesRollbackToVersion struct {
	ID string `json:"id,omitempty"` // The version ID
}
type RequestNetworksCreateNetworkFirmwareUpgradesStagedEvent struct {
	Products *RequestNetworksCreateNetworkFirmwareUpgradesStagedEventProducts `json:"products,omitempty"` // Contains firmware upgrade version information
	Stages   *[]RequestNetworksCreateNetworkFirmwareUpgradesStagedEventStages `json:"stages,omitempty"`   // All firmware upgrade stages in the network with their start time.
}
type RequestNetworksCreateNetworkFirmwareUpgradesStagedEventProducts struct {
	Switch *RequestNetworksCreateNetworkFirmwareUpgradesStagedEventProductsSwitch `json:"switch,omitempty"` // Version information for the switch network being upgraded
}
type RequestNetworksCreateNetworkFirmwareUpgradesStagedEventProductsSwitch struct {
	NextUpgrade *RequestNetworksCreateNetworkFirmwareUpgradesStagedEventProductsSwitchNextUpgrade `json:"nextUpgrade,omitempty"` // The next upgrade version for the switch network
}
type RequestNetworksCreateNetworkFirmwareUpgradesStagedEventProductsSwitchNextUpgrade struct {
	ToVersion *RequestNetworksCreateNetworkFirmwareUpgradesStagedEventProductsSwitchNextUpgradeToVersion `json:"toVersion,omitempty"` // The version to be updated to for switch devices
}
type RequestNetworksCreateNetworkFirmwareUpgradesStagedEventProductsSwitchNextUpgradeToVersion struct {
	ID string `json:"id,omitempty"` // The version ID
}
type RequestNetworksCreateNetworkFirmwareUpgradesStagedEventStages struct {
	Group      *RequestNetworksCreateNetworkFirmwareUpgradesStagedEventStagesGroup      `json:"group,omitempty"`      // The Staged Upgrade Group containing the name and ID
	Milestones *RequestNetworksCreateNetworkFirmwareUpgradesStagedEventStagesMilestones `json:"milestones,omitempty"` // The Staged Upgrade Milestones for the specific stage
}
type RequestNetworksCreateNetworkFirmwareUpgradesStagedEventStagesGroup struct {
	ID string `json:"id,omitempty"` // ID of the Staged Upgrade Group
}
type RequestNetworksCreateNetworkFirmwareUpgradesStagedEventStagesMilestones struct {
	ScheduledFor string `json:"scheduledFor,omitempty"` // The start time of the staged upgrade stage. (In ISO-8601 format, in the time zone of the network.)
}
type RequestNetworksUpdateNetworkFirmwareUpgradesStagedEvents struct {
	Stages *[]RequestNetworksUpdateNetworkFirmwareUpgradesStagedEventsStages `json:"stages,omitempty"` // All firmware upgrade stages in the network with their start time.
}
type RequestNetworksUpdateNetworkFirmwareUpgradesStagedEventsStages struct {
	Group      *RequestNetworksUpdateNetworkFirmwareUpgradesStagedEventsStagesGroup      `json:"group,omitempty"`      // The Staged Upgrade Group containing the name and ID
	Milestones *RequestNetworksUpdateNetworkFirmwareUpgradesStagedEventsStagesMilestones `json:"milestones,omitempty"` // The Staged Upgrade Milestones for the specific stage
}
type RequestNetworksUpdateNetworkFirmwareUpgradesStagedEventsStagesGroup struct {
	ID string `json:"id,omitempty"` // ID of the Staged Upgrade Group
}
type RequestNetworksUpdateNetworkFirmwareUpgradesStagedEventsStagesMilestones struct {
	ScheduledFor string `json:"scheduledFor,omitempty"` // The start time of the staged upgrade stage. (In ISO-8601 format, in the time zone of the network.)
}
type RequestNetworksRollbacksNetworkFirmwareUpgradesStagedEvents struct {
	Reasons *[]RequestNetworksRollbacksNetworkFirmwareUpgradesStagedEventsReasons `json:"reasons,omitempty"` // The reason for rolling back the staged upgrade
	Stages  *[]RequestNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStages  `json:"stages,omitempty"`  // All completed or in-progress stages in the network with their new start times. All pending stages will be canceled
}
type RequestNetworksRollbacksNetworkFirmwareUpgradesStagedEventsReasons struct {
	Category string `json:"category,omitempty"` // Reason for the rollback
	Comment  string `json:"comment,omitempty"`  // Additional comment about the rollback
}
type RequestNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStages struct {
	Group      *RequestNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStagesGroup      `json:"group,omitempty"`      // The Staged Upgrade Group containing the name and ID
	Milestones *RequestNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStagesMilestones `json:"milestones,omitempty"` // The Staged Upgrade Milestones for the specific stage
}
type RequestNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStagesGroup struct {
	ID string `json:"id,omitempty"` // ID of the Staged Upgrade Group
}
type RequestNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStagesMilestones struct {
	ScheduledFor string `json:"scheduledFor,omitempty"` // The start time of the staged upgrade stage. (In ISO-8601 format, in the time zone of the network.)
}
type RequestNetworksCreateNetworkFirmwareUpgradesStagedGroup struct {
	AssignedDevices *RequestNetworksCreateNetworkFirmwareUpgradesStagedGroupAssignedDevices `json:"assignedDevices,omitempty"` // The devices and Switch Stacks assigned to the Group
	Description     string                                                                  `json:"description,omitempty"`     // Description of the Staged Upgrade Group. Length must be 1 to 255 characters
	IsDefault       *bool                                                                   `json:"isDefault,omitempty"`       // Boolean indicating the default Group. Any device that does not have a group explicitly assigned will upgrade with this group
	Name            string                                                                  `json:"name,omitempty"`            // Name of the Staged Upgrade Group. Length must be 1 to 255 characters
}
type RequestNetworksCreateNetworkFirmwareUpgradesStagedGroupAssignedDevices struct {
	Devices      *[]RequestNetworksCreateNetworkFirmwareUpgradesStagedGroupAssignedDevicesDevices      `json:"devices,omitempty"`      // Data Array of Devices containing the name and serial
	SwitchStacks *[]RequestNetworksCreateNetworkFirmwareUpgradesStagedGroupAssignedDevicesSwitchStacks `json:"switchStacks,omitempty"` // Data Array of Switch Stacks containing the name and id
}
type RequestNetworksCreateNetworkFirmwareUpgradesStagedGroupAssignedDevicesDevices struct {
	Name   string `json:"name,omitempty"`   // Name of the device
	Serial string `json:"serial,omitempty"` // Serial of the device
}
type RequestNetworksCreateNetworkFirmwareUpgradesStagedGroupAssignedDevicesSwitchStacks struct {
	ID   string `json:"id,omitempty"`   // ID of the Switch Stack
	Name string `json:"name,omitempty"` // Name of the Switch Stack
}
type RequestNetworksUpdateNetworkFirmwareUpgradesStagedGroup struct {
	AssignedDevices *RequestNetworksUpdateNetworkFirmwareUpgradesStagedGroupAssignedDevices `json:"assignedDevices,omitempty"` // The devices and Switch Stacks assigned to the Group
	Description     string                                                                  `json:"description,omitempty"`     // Description of the Staged Upgrade Group. Length must be 1 to 255 characters
	IsDefault       *bool                                                                   `json:"isDefault,omitempty"`       // Boolean indicating the default Group. Any device that does not have a group explicitly assigned will upgrade with this group
	Name            string                                                                  `json:"name,omitempty"`            // Name of the Staged Upgrade Group. Length must be 1 to 255 characters
}
type RequestNetworksUpdateNetworkFirmwareUpgradesStagedGroupAssignedDevices struct {
	Devices      *[]RequestNetworksUpdateNetworkFirmwareUpgradesStagedGroupAssignedDevicesDevices      `json:"devices,omitempty"`      // Data Array of Devices containing the name and serial
	SwitchStacks *[]RequestNetworksUpdateNetworkFirmwareUpgradesStagedGroupAssignedDevicesSwitchStacks `json:"switchStacks,omitempty"` // Data Array of Switch Stacks containing the name and id
}
type RequestNetworksUpdateNetworkFirmwareUpgradesStagedGroupAssignedDevicesDevices struct {
	Name   string `json:"name,omitempty"`   // Name of the device
	Serial string `json:"serial,omitempty"` // Serial of the device
}
type RequestNetworksUpdateNetworkFirmwareUpgradesStagedGroupAssignedDevicesSwitchStacks struct {
	ID   string `json:"id,omitempty"`   // ID of the Switch Stack
	Name string `json:"name,omitempty"` // Name of the Switch Stack
}
type RequestNetworksUpdateNetworkFirmwareUpgradesStagedStages struct {
	JSON *[]RequestNetworksUpdateNetworkFirmwareUpgradesStagedStagesJSON `json:"_json,omitempty"` // Array of Staged Upgrade Groups
}
type RequestNetworksUpdateNetworkFirmwareUpgradesStagedStagesJSON struct {
	Group *RequestNetworksUpdateNetworkFirmwareUpgradesStagedStagesJSONGroup `json:"group,omitempty"` // The Staged Upgrade Group
}
type RequestNetworksUpdateNetworkFirmwareUpgradesStagedStagesJSONGroup struct {
	ID string `json:"id,omitempty"` // ID of the Staged Upgrade Group
}
type RequestNetworksCreateNetworkFloorPlan struct {
	BottomLeftCorner  *RequestNetworksCreateNetworkFloorPlanBottomLeftCorner  `json:"bottomLeftCorner,omitempty"`  // The longitude and latitude of the bottom left corner of your floor plan.
	BottomRightCorner *RequestNetworksCreateNetworkFloorPlanBottomRightCorner `json:"bottomRightCorner,omitempty"` // The longitude and latitude of the bottom right corner of your floor plan.
	Center            *RequestNetworksCreateNetworkFloorPlanCenter            `json:"center,omitempty"`            // The longitude and latitude of the center of your floor plan. The 'center' or two adjacent corners (e.g. 'topLeftCorner' and 'bottomLeftCorner') must be specified. If 'center' is specified, the floor plan is placed over that point with no rotation. If two adjacent corners are specified, the floor plan is rotated to line up with the two specified points. The aspect ratio of the floor plan's image is preserved regardless of which corners/center are specified. (This means if that more than two corners are specified, only two corners may be used to preserve the floor plan's aspect ratio.). No two points can have the same latitude, longitude pair.
	ImageContents     string                                                  `json:"imageContents,omitempty"`     // The file contents (a base 64 encoded string) of your image. Supported formats are PNG, GIF, and JPG. Note that all images are saved as PNG files, regardless of the format they are uploaded in.
	Name              string                                                  `json:"name,omitempty"`              // The name of your floor plan.
	TopLeftCorner     *RequestNetworksCreateNetworkFloorPlanTopLeftCorner     `json:"topLeftCorner,omitempty"`     // The longitude and latitude of the top left corner of your floor plan.
	TopRightCorner    *RequestNetworksCreateNetworkFloorPlanTopRightCorner    `json:"topRightCorner,omitempty"`    // The longitude and latitude of the top right corner of your floor plan.
}
type RequestNetworksCreateNetworkFloorPlanBottomLeftCorner struct {
	Lat *float64 `json:"lat,omitempty"` // Latitude
	Lng *float64 `json:"lng,omitempty"` // Longitude
}
type RequestNetworksCreateNetworkFloorPlanBottomRightCorner struct {
	Lat *float64 `json:"lat,omitempty"` // Latitude
	Lng *float64 `json:"lng,omitempty"` // Longitude
}
type RequestNetworksCreateNetworkFloorPlanCenter struct {
	Lat *float64 `json:"lat,omitempty"` // Latitude
	Lng *float64 `json:"lng,omitempty"` // Longitude
}
type RequestNetworksCreateNetworkFloorPlanTopLeftCorner struct {
	Lat *float64 `json:"lat,omitempty"` // Latitude
	Lng *float64 `json:"lng,omitempty"` // Longitude
}
type RequestNetworksCreateNetworkFloorPlanTopRightCorner struct {
	Lat *float64 `json:"lat,omitempty"` // Latitude
	Lng *float64 `json:"lng,omitempty"` // Longitude
}
type RequestNetworksUpdateNetworkFloorPlan struct {
	BottomLeftCorner  *RequestNetworksUpdateNetworkFloorPlanBottomLeftCorner  `json:"bottomLeftCorner,omitempty"`  // The longitude and latitude of the bottom left corner of your floor plan.
	BottomRightCorner *RequestNetworksUpdateNetworkFloorPlanBottomRightCorner `json:"bottomRightCorner,omitempty"` // The longitude and latitude of the bottom right corner of your floor plan.
	Center            *RequestNetworksUpdateNetworkFloorPlanCenter            `json:"center,omitempty"`            // The longitude and latitude of the center of your floor plan. If you want to change the geolocation data of your floor plan, either the 'center' or two adjacent corners (e.g. 'topLeftCorner' and 'bottomLeftCorner') must be specified. If 'center' is specified, the floor plan is placed over that point with no rotation. If two adjacent corners are specified, the floor plan is rotated to line up with the two specified points. The aspect ratio of the floor plan's image is preserved regardless of which corners/center are specified. (This means if that more than two corners are specified, only two corners may be used to preserve the floor plan's aspect ratio.). No two points can have the same latitude, longitude pair.
	ImageContents     string                                                  `json:"imageContents,omitempty"`     // The file contents (a base 64 encoded string) of your new image. Supported formats are PNG, GIF, and JPG. Note that all images are saved as PNG files, regardless of the format they are uploaded in. If you upload a new image, and you do NOT specify any new geolocation fields ('center, 'topLeftCorner', etc), the floor plan will be recentered with no rotation in order to maintain the aspect ratio of your new image.
	Name              string                                                  `json:"name,omitempty"`              // The name of your floor plan.
	TopLeftCorner     *RequestNetworksUpdateNetworkFloorPlanTopLeftCorner     `json:"topLeftCorner,omitempty"`     // The longitude and latitude of the top left corner of your floor plan.
	TopRightCorner    *RequestNetworksUpdateNetworkFloorPlanTopRightCorner    `json:"topRightCorner,omitempty"`    // The longitude and latitude of the top right corner of your floor plan.
}
type RequestNetworksUpdateNetworkFloorPlanBottomLeftCorner struct {
	Lat *float64 `json:"lat,omitempty"` // Latitude
	Lng *float64 `json:"lng,omitempty"` // Longitude
}
type RequestNetworksUpdateNetworkFloorPlanBottomRightCorner struct {
	Lat *float64 `json:"lat,omitempty"` // Latitude
	Lng *float64 `json:"lng,omitempty"` // Longitude
}
type RequestNetworksUpdateNetworkFloorPlanCenter struct {
	Lat *float64 `json:"lat,omitempty"` // Latitude
	Lng *float64 `json:"lng,omitempty"` // Longitude
}
type RequestNetworksUpdateNetworkFloorPlanTopLeftCorner struct {
	Lat *float64 `json:"lat,omitempty"` // Latitude
	Lng *float64 `json:"lng,omitempty"` // Longitude
}
type RequestNetworksUpdateNetworkFloorPlanTopRightCorner struct {
	Lat *float64 `json:"lat,omitempty"` // Latitude
	Lng *float64 `json:"lng,omitempty"` // Longitude
}
type RequestNetworksCreateNetworkGroupPolicy struct {
	Bandwidth                 *RequestNetworksCreateNetworkGroupPolicyBandwidth                 `json:"bandwidth,omitempty"`                 //     The bandwidth settings for clients bound to your group policy.
	BonjourForwarding         *RequestNetworksCreateNetworkGroupPolicyBonjourForwarding         `json:"bonjourForwarding,omitempty"`         // The Bonjour settings for your group policy. Only valid if your network has a wireless configuration.
	ContentFiltering          *RequestNetworksCreateNetworkGroupPolicyContentFiltering          `json:"contentFiltering,omitempty"`          // The content filtering settings for your group policy
	FirewallAndTrafficShaping *RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShaping `json:"firewallAndTrafficShaping,omitempty"` //     The firewall and traffic shaping rules and settings for your policy.
	Name                      string                                                            `json:"name,omitempty"`                      // The name for your group policy. Required.
	Scheduling                *RequestNetworksCreateNetworkGroupPolicyScheduling                `json:"scheduling,omitempty"`                //     The schedule for the group policy. Schedules are applied to days of the week.
	SplashAuthSettings        string                                                            `json:"splashAuthSettings,omitempty"`        // Whether clients bound to your policy will bypass splash authorization or behave according to the network's rules. Can be one of 'network default' or 'bypass'. Only available if your network has a wireless configuration.
	VLANTagging               *RequestNetworksCreateNetworkGroupPolicyVLANTagging               `json:"vlanTagging,omitempty"`               // The VLAN tagging settings for your group policy. Only available if your network has a wireless configuration.
}
type RequestNetworksCreateNetworkGroupPolicyBandwidth struct {
	BandwidthLimits *RequestNetworksCreateNetworkGroupPolicyBandwidthBandwidthLimits `json:"bandwidthLimits,omitempty"` // The bandwidth limits object, specifying upload and download speed for clients bound to the group policy. These are only enforced if 'settings' is set to 'custom'.
	Settings        string                                                           `json:"settings,omitempty"`        // How bandwidth limits are enforced. Can be 'network default', 'ignore' or 'custom'.
}
type RequestNetworksCreateNetworkGroupPolicyBandwidthBandwidthLimits struct {
	LimitDown *int `json:"limitDown,omitempty"` // The maximum download limit (integer, in Kbps). null indicates no limit
	LimitUp   *int `json:"limitUp,omitempty"`   // The maximum upload limit (integer, in Kbps). null indicates no limit
}
type RequestNetworksCreateNetworkGroupPolicyBonjourForwarding struct {
	Rules    *[]RequestNetworksCreateNetworkGroupPolicyBonjourForwardingRules `json:"rules,omitempty"`    // A list of the Bonjour forwarding rules for your group policy. If 'settings' is set to 'custom', at least one rule must be specified.
	Settings string                                                           `json:"settings,omitempty"` // How Bonjour rules are applied. Can be 'network default', 'ignore' or 'custom'.
}
type RequestNetworksCreateNetworkGroupPolicyBonjourForwardingRules struct {
	Description string   `json:"description,omitempty"` // A description for your Bonjour forwarding rule. Optional.
	Services    []string `json:"services,omitempty"`    // A list of Bonjour services. At least one service must be specified. Available services are 'All Services', 'AirPlay', 'AFP', 'BitTorrent', 'FTP', 'iChat', 'iTunes', 'Printers', 'Samba', 'Scanners' and 'SSH'
	VLANID      string   `json:"vlanId,omitempty"`      // The ID of the service VLAN. Required.
}
type RequestNetworksCreateNetworkGroupPolicyContentFiltering struct {
	AllowedURLPatterns   *RequestNetworksCreateNetworkGroupPolicyContentFilteringAllowedURLPatterns   `json:"allowedUrlPatterns,omitempty"`   // Settings for allowed URL patterns
	BlockedURLCategories *RequestNetworksCreateNetworkGroupPolicyContentFilteringBlockedURLCategories `json:"blockedUrlCategories,omitempty"` // Settings for blocked URL categories
	BlockedURLPatterns   *RequestNetworksCreateNetworkGroupPolicyContentFilteringBlockedURLPatterns   `json:"blockedUrlPatterns,omitempty"`   // Settings for blocked URL patterns
}
type RequestNetworksCreateNetworkGroupPolicyContentFilteringAllowedURLPatterns struct {
	Patterns []string `json:"patterns,omitempty"` // A list of URL patterns that are allowed
	Settings string   `json:"settings,omitempty"` // How URL patterns are applied. Can be 'network default', 'append' or 'override'.
}
type RequestNetworksCreateNetworkGroupPolicyContentFilteringBlockedURLCategories struct {
	Categories []string `json:"categories,omitempty"` // A list of URL categories to block
	Settings   string   `json:"settings,omitempty"`   // How URL categories are applied. Can be 'network default', 'append' or 'override'.
}
type RequestNetworksCreateNetworkGroupPolicyContentFilteringBlockedURLPatterns struct {
	Patterns []string `json:"patterns,omitempty"` // A list of URL patterns that are blocked
	Settings string   `json:"settings,omitempty"` // How URL patterns are applied. Can be 'network default', 'append' or 'override'.
}
type RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShaping struct {
	L3FirewallRules     *[]RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRules     `json:"l3FirewallRules,omitempty"`     // An ordered array of the L3 firewall rules
	L7FirewallRules     *[]RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules     `json:"l7FirewallRules,omitempty"`     // An ordered array of L7 firewall rules
	Settings            string                                                                                 `json:"settings,omitempty"`            // How firewall and traffic shaping rules are enforced. Can be 'network default', 'ignore' or 'custom'.
	TrafficShapingRules *[]RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRules `json:"trafficShapingRules,omitempty"` //     An array of traffic shaping rules. Rules are applied in the order that     they are specified in. An empty list (or null) means no rules. Note that     you are allowed a maximum of 8 rules.
}
type RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRules struct {
	Comment  string `json:"comment,omitempty"`  // Description of the rule (optional)
	DestCidr string `json:"destCidr,omitempty"` // Destination IP address (in IP or CIDR notation), a fully-qualified domain name (FQDN, if your network supports it) or 'any'.
	DestPort string `json:"destPort,omitempty"` // Destination port (integer in the range 1-65535), a port range (e.g. 8080-9090), or 'any'
	Policy   string `json:"policy,omitempty"`   // 'allow' or 'deny' traffic specified by this rule
	Protocol string `json:"protocol,omitempty"` // The type of protocol (must be 'tcp', 'udp', 'icmp', 'icmp6' or 'any')
}
type RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules struct {
	Policy string `json:"policy,omitempty"` // The policy applied to matching traffic. Must be 'deny'.
	Type   string `json:"type,omitempty"`   // Type of the L7 Rule. Must be 'application', 'applicationCategory', 'host', 'port' or 'ipRange'
	Value  string `json:"value,omitempty"`  // The 'value' of what you want to block. If 'type' is 'host', 'port' or 'ipRange', 'value' must be a string matching either a hostname (e.g. somewhere.com), a port (e.g. 8080), or an IP range (e.g. 192.1.0.0/16). If 'type' is 'application' or 'applicationCategory', then 'value' must be an object with an ID for the application.
}
type RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRules struct {
	Definitions              *[]RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions            `json:"definitions,omitempty"`              //     A list of objects describing the definitions of your traffic shaping rule. At least one definition is required.
	DscpTagValue             *int                                                                                                         `json:"dscpTagValue,omitempty"`             //     The DSCP tag applied by your rule. null means 'Do not change DSCP tag'.     For a list of possible tag values, use the trafficShaping/dscpTaggingOptions endpoint.
	PcpTagValue              *int                                                                                                         `json:"pcpTagValue,omitempty"`              //     The PCP tag applied by your rule. Can be 0 (lowest priority) through 7 (highest priority).     null means 'Do not set PCP tag'.
	PerClientBandwidthLimits *RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimits `json:"perClientBandwidthLimits,omitempty"` //     An object describing the bandwidth settings for your rule.
	Priority                 string                                                                                                       `json:"priority,omitempty"`                 //     A string, indicating the priority level for packets bound to your rule.     Can be 'low', 'normal' or 'high'.
}
type RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions struct {
	Type  string `json:"type,omitempty"`  // The type of definition. Can be one of 'application', 'applicationCategory', 'host', 'port', 'ipRange' or 'localNet'.
	Value string `json:"value,omitempty"` //     If "type" is 'host', 'port', 'ipRange' or 'localNet', then "value" must be a string, matching either     a hostname (e.g. "somesite.com"), a port (e.g. 8080), or an IP range ("192.1.0.0",     "192.1.0.0/16", or "10.1.0.0/16:80"). 'localNet' also supports CIDR notation, excluding     custom ports.      If "type" is 'application' or 'applicationCategory', then "value" must be an object     with the structure { "id": "meraki:layer7/..." }, where "id" is the application category or     application ID (for a list of IDs for your network, use the trafficShaping/applicationCategories     endpoint).
}
type RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimits struct {
	BandwidthLimits *RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimitsBandwidthLimits `json:"bandwidthLimits,omitempty"` // The bandwidth limits object, specifying the upload ('limitUp') and download ('limitDown') speed in Kbps. These are only enforced if 'settings' is set to 'custom'.
	Settings        string                                                                                                                      `json:"settings,omitempty"`        // How bandwidth limits are applied by your rule. Can be one of 'network default', 'ignore' or 'custom'.
}
type RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimitsBandwidthLimits struct {
	LimitDown *int `json:"limitDown,omitempty"` // The maximum download limit (integer, in Kbps).
	LimitUp   *int `json:"limitUp,omitempty"`   // The maximum upload limit (integer, in Kbps).
}
type RequestNetworksCreateNetworkGroupPolicyScheduling struct {
	Enabled   *bool                                                       `json:"enabled,omitempty"`   // Whether scheduling is enabled (true) or disabled (false). Defaults to false. If true, the schedule objects for each day of the week (monday - sunday) are parsed.
	Friday    *RequestNetworksCreateNetworkGroupPolicySchedulingFriday    `json:"friday,omitempty"`    // The schedule object for Friday.
	Monday    *RequestNetworksCreateNetworkGroupPolicySchedulingMonday    `json:"monday,omitempty"`    // The schedule object for Monday.
	Saturday  *RequestNetworksCreateNetworkGroupPolicySchedulingSaturday  `json:"saturday,omitempty"`  // The schedule object for Saturday.
	Sunday    *RequestNetworksCreateNetworkGroupPolicySchedulingSunday    `json:"sunday,omitempty"`    // The schedule object for Sunday.
	Thursday  *RequestNetworksCreateNetworkGroupPolicySchedulingThursday  `json:"thursday,omitempty"`  // The schedule object for Thursday.
	Tuesday   *RequestNetworksCreateNetworkGroupPolicySchedulingTuesday   `json:"tuesday,omitempty"`   // The schedule object for Tuesday.
	Wednesday *RequestNetworksCreateNetworkGroupPolicySchedulingWednesday `json:"wednesday,omitempty"` // The schedule object for Wednesday.
}
type RequestNetworksCreateNetworkGroupPolicySchedulingFriday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
}
type RequestNetworksCreateNetworkGroupPolicySchedulingMonday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
}
type RequestNetworksCreateNetworkGroupPolicySchedulingSaturday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
}
type RequestNetworksCreateNetworkGroupPolicySchedulingSunday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
}
type RequestNetworksCreateNetworkGroupPolicySchedulingThursday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
}
type RequestNetworksCreateNetworkGroupPolicySchedulingTuesday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
}
type RequestNetworksCreateNetworkGroupPolicySchedulingWednesday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
}
type RequestNetworksCreateNetworkGroupPolicyVLANTagging struct {
	Settings string `json:"settings,omitempty"` // How VLAN tagging is applied. Can be 'network default', 'ignore' or 'custom'.
	VLANID   string `json:"vlanId,omitempty"`   // The ID of the vlan you want to tag. This only applies if 'settings' is set to 'custom'.
}
type RequestNetworksUpdateNetworkGroupPolicy struct {
	Bandwidth                 *RequestNetworksUpdateNetworkGroupPolicyBandwidth                 `json:"bandwidth,omitempty"`                 //     The bandwidth settings for clients bound to your group policy.
	BonjourForwarding         *RequestNetworksUpdateNetworkGroupPolicyBonjourForwarding         `json:"bonjourForwarding,omitempty"`         // The Bonjour settings for your group policy. Only valid if your network has a wireless configuration.
	ContentFiltering          *RequestNetworksUpdateNetworkGroupPolicyContentFiltering          `json:"contentFiltering,omitempty"`          // The content filtering settings for your group policy
	FirewallAndTrafficShaping *RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShaping `json:"firewallAndTrafficShaping,omitempty"` //     The firewall and traffic shaping rules and settings for your policy.
	Name                      string                                                            `json:"name,omitempty"`                      // The name for your group policy.
	Scheduling                *RequestNetworksUpdateNetworkGroupPolicyScheduling                `json:"scheduling,omitempty"`                //     The schedule for the group policy. Schedules are applied to days of the week.
	SplashAuthSettings        string                                                            `json:"splashAuthSettings,omitempty"`        // Whether clients bound to your policy will bypass splash authorization or behave according to the network's rules. Can be one of 'network default' or 'bypass'. Only available if your network has a wireless configuration.
	VLANTagging               *RequestNetworksUpdateNetworkGroupPolicyVLANTagging               `json:"vlanTagging,omitempty"`               // The VLAN tagging settings for your group policy. Only available if your network has a wireless configuration.
}
type RequestNetworksUpdateNetworkGroupPolicyBandwidth struct {
	BandwidthLimits *RequestNetworksUpdateNetworkGroupPolicyBandwidthBandwidthLimits `json:"bandwidthLimits,omitempty"` // The bandwidth limits object, specifying upload and download speed for clients bound to the group policy. These are only enforced if 'settings' is set to 'custom'.
	Settings        string                                                           `json:"settings,omitempty"`        // How bandwidth limits are enforced. Can be 'network default', 'ignore' or 'custom'.
}
type RequestNetworksUpdateNetworkGroupPolicyBandwidthBandwidthLimits struct {
	LimitDown *int `json:"limitDown,omitempty"` // The maximum download limit (integer, in Kbps). null indicates no limit
	LimitUp   *int `json:"limitUp,omitempty"`   // The maximum upload limit (integer, in Kbps). null indicates no limit
}
type RequestNetworksUpdateNetworkGroupPolicyBonjourForwarding struct {
	Rules    *[]RequestNetworksUpdateNetworkGroupPolicyBonjourForwardingRules `json:"rules,omitempty"`    // A list of the Bonjour forwarding rules for your group policy. If 'settings' is set to 'custom', at least one rule must be specified.
	Settings string                                                           `json:"settings,omitempty"` // How Bonjour rules are applied. Can be 'network default', 'ignore' or 'custom'.
}
type RequestNetworksUpdateNetworkGroupPolicyBonjourForwardingRules struct {
	Description string   `json:"description,omitempty"` // A description for your Bonjour forwarding rule. Optional.
	Services    []string `json:"services,omitempty"`    // A list of Bonjour services. At least one service must be specified. Available services are 'All Services', 'AirPlay', 'AFP', 'BitTorrent', 'FTP', 'iChat', 'iTunes', 'Printers', 'Samba', 'Scanners' and 'SSH'
	VLANID      string   `json:"vlanId,omitempty"`      // The ID of the service VLAN. Required.
}
type RequestNetworksUpdateNetworkGroupPolicyContentFiltering struct {
	AllowedURLPatterns   *RequestNetworksUpdateNetworkGroupPolicyContentFilteringAllowedURLPatterns   `json:"allowedUrlPatterns,omitempty"`   // Settings for allowed URL patterns
	BlockedURLCategories *RequestNetworksUpdateNetworkGroupPolicyContentFilteringBlockedURLCategories `json:"blockedUrlCategories,omitempty"` // Settings for blocked URL categories
	BlockedURLPatterns   *RequestNetworksUpdateNetworkGroupPolicyContentFilteringBlockedURLPatterns   `json:"blockedUrlPatterns,omitempty"`   // Settings for blocked URL patterns
}
type RequestNetworksUpdateNetworkGroupPolicyContentFilteringAllowedURLPatterns struct {
	Patterns []string `json:"patterns,omitempty"` // A list of URL patterns that are allowed
	Settings string   `json:"settings,omitempty"` // How URL patterns are applied. Can be 'network default', 'append' or 'override'.
}
type RequestNetworksUpdateNetworkGroupPolicyContentFilteringBlockedURLCategories struct {
	Categories []string `json:"categories,omitempty"` // A list of URL categories to block
	Settings   string   `json:"settings,omitempty"`   // How URL categories are applied. Can be 'network default', 'append' or 'override'.
}
type RequestNetworksUpdateNetworkGroupPolicyContentFilteringBlockedURLPatterns struct {
	Patterns []string `json:"patterns,omitempty"` // A list of URL patterns that are blocked
	Settings string   `json:"settings,omitempty"` // How URL patterns are applied. Can be 'network default', 'append' or 'override'.
}
type RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShaping struct {
	L3FirewallRules     *[]RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRules     `json:"l3FirewallRules,omitempty"`     // An ordered array of the L3 firewall rules
	L7FirewallRules     *[]RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules     `json:"l7FirewallRules,omitempty"`     // An ordered array of L7 firewall rules
	Settings            string                                                                                 `json:"settings,omitempty"`            // How firewall and traffic shaping rules are enforced. Can be 'network default', 'ignore' or 'custom'.
	TrafficShapingRules *[]RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRules `json:"trafficShapingRules,omitempty"` //     An array of traffic shaping rules. Rules are applied in the order that     they are specified in. An empty list (or null) means no rules. Note that     you are allowed a maximum of 8 rules.
}
type RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRules struct {
	Comment  string `json:"comment,omitempty"`  // Description of the rule (optional)
	DestCidr string `json:"destCidr,omitempty"` // Destination IP address (in IP or CIDR notation), a fully-qualified domain name (FQDN, if your network supports it) or 'any'.
	DestPort string `json:"destPort,omitempty"` // Destination port (integer in the range 1-65535), a port range (e.g. 8080-9090), or 'any'
	Policy   string `json:"policy,omitempty"`   // 'allow' or 'deny' traffic specified by this rule
	Protocol string `json:"protocol,omitempty"` // The type of protocol (must be 'tcp', 'udp', 'icmp', 'icmp6' or 'any')
}
type RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules struct {
	Policy string `json:"policy,omitempty"` // The policy applied to matching traffic. Must be 'deny'.
	Type   string `json:"type,omitempty"`   // Type of the L7 Rule. Must be 'application', 'applicationCategory', 'host', 'port' or 'ipRange'
	Value  string `json:"value,omitempty"`  // The 'value' of what you want to block. If 'type' is 'host', 'port' or 'ipRange', 'value' must be a string matching either a hostname (e.g. somewhere.com), a port (e.g. 8080), or an IP range (e.g. 192.1.0.0/16). If 'type' is 'application' or 'applicationCategory', then 'value' must be an object with an ID for the application.
}
type RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRules struct {
	Definitions              *[]RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions            `json:"definitions,omitempty"`              //     A list of objects describing the definitions of your traffic shaping rule. At least one definition is required.
	DscpTagValue             *int                                                                                                         `json:"dscpTagValue,omitempty"`             //     The DSCP tag applied by your rule. null means 'Do not change DSCP tag'.     For a list of possible tag values, use the trafficShaping/dscpTaggingOptions endpoint.
	PcpTagValue              *int                                                                                                         `json:"pcpTagValue,omitempty"`              //     The PCP tag applied by your rule. Can be 0 (lowest priority) through 7 (highest priority).     null means 'Do not set PCP tag'.
	PerClientBandwidthLimits *RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimits `json:"perClientBandwidthLimits,omitempty"` //     An object describing the bandwidth settings for your rule.
	Priority                 string                                                                                                       `json:"priority,omitempty"`                 //     A string, indicating the priority level for packets bound to your rule.     Can be 'low', 'normal' or 'high'.
}
type RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions struct {
	Type  string `json:"type,omitempty"`  // The type of definition. Can be one of 'application', 'applicationCategory', 'host', 'port', 'ipRange' or 'localNet'.
	Value string `json:"value,omitempty"` //     If "type" is 'host', 'port', 'ipRange' or 'localNet', then "value" must be a string, matching either     a hostname (e.g. "somesite.com"), a port (e.g. 8080), or an IP range ("192.1.0.0",     "192.1.0.0/16", or "10.1.0.0/16:80"). 'localNet' also supports CIDR notation, excluding     custom ports.      If "type" is 'application' or 'applicationCategory', then "value" must be an object     with the structure { "id": "meraki:layer7/..." }, where "id" is the application category or     application ID (for a list of IDs for your network, use the trafficShaping/applicationCategories     endpoint).
}
type RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimits struct {
	BandwidthLimits *RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimitsBandwidthLimits `json:"bandwidthLimits,omitempty"` // The bandwidth limits object, specifying the upload ('limitUp') and download ('limitDown') speed in Kbps. These are only enforced if 'settings' is set to 'custom'.
	Settings        string                                                                                                                      `json:"settings,omitempty"`        // How bandwidth limits are applied by your rule. Can be one of 'network default', 'ignore' or 'custom'.
}
type RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimitsBandwidthLimits struct {
	LimitDown *int `json:"limitDown,omitempty"` // The maximum download limit (integer, in Kbps).
	LimitUp   *int `json:"limitUp,omitempty"`   // The maximum upload limit (integer, in Kbps).
}
type RequestNetworksUpdateNetworkGroupPolicyScheduling struct {
	Enabled   *bool                                                       `json:"enabled,omitempty"`   // Whether scheduling is enabled (true) or disabled (false). Defaults to false. If true, the schedule objects for each day of the week (monday - sunday) are parsed.
	Friday    *RequestNetworksUpdateNetworkGroupPolicySchedulingFriday    `json:"friday,omitempty"`    // The schedule object for Friday.
	Monday    *RequestNetworksUpdateNetworkGroupPolicySchedulingMonday    `json:"monday,omitempty"`    // The schedule object for Monday.
	Saturday  *RequestNetworksUpdateNetworkGroupPolicySchedulingSaturday  `json:"saturday,omitempty"`  // The schedule object for Saturday.
	Sunday    *RequestNetworksUpdateNetworkGroupPolicySchedulingSunday    `json:"sunday,omitempty"`    // The schedule object for Sunday.
	Thursday  *RequestNetworksUpdateNetworkGroupPolicySchedulingThursday  `json:"thursday,omitempty"`  // The schedule object for Thursday.
	Tuesday   *RequestNetworksUpdateNetworkGroupPolicySchedulingTuesday   `json:"tuesday,omitempty"`   // The schedule object for Tuesday.
	Wednesday *RequestNetworksUpdateNetworkGroupPolicySchedulingWednesday `json:"wednesday,omitempty"` // The schedule object for Wednesday.
}
type RequestNetworksUpdateNetworkGroupPolicySchedulingFriday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
}
type RequestNetworksUpdateNetworkGroupPolicySchedulingMonday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
}
type RequestNetworksUpdateNetworkGroupPolicySchedulingSaturday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
}
type RequestNetworksUpdateNetworkGroupPolicySchedulingSunday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
}
type RequestNetworksUpdateNetworkGroupPolicySchedulingThursday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
}
type RequestNetworksUpdateNetworkGroupPolicySchedulingTuesday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
}
type RequestNetworksUpdateNetworkGroupPolicySchedulingWednesday struct {
	Active *bool  `json:"active,omitempty"` // Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	From   string `json:"from,omitempty"`   // The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	To     string `json:"to,omitempty"`     // The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
}
type RequestNetworksUpdateNetworkGroupPolicyVLANTagging struct {
	Settings string `json:"settings,omitempty"` // How VLAN tagging is applied. Can be 'network default', 'ignore' or 'custom'.
	VLANID   string `json:"vlanId,omitempty"`   // The ID of the vlan you want to tag. This only applies if 'settings' is set to 'custom'.
}
type RequestNetworksCreateNetworkMerakiAuthUser struct {
	AccountType         string                                                      `json:"accountType,omitempty"`         // Authorization type for user. Can be 'Guest' or '802.1X' for wireless networks, or 'Client VPN' for wired networks. Defaults to '802.1X'.
	Authorizations      *[]RequestNetworksCreateNetworkMerakiAuthUserAuthorizations `json:"authorizations,omitempty"`      // Authorization zones and expiration dates for the user.
	Email               string                                                      `json:"email,omitempty"`               // Email address of the user
	EmailPasswordToUser *bool                                                       `json:"emailPasswordToUser,omitempty"` // Whether or not Meraki should email the password to user. Default is false.
	IsAdmin             *bool                                                       `json:"isAdmin,omitempty"`             // Whether or not the user is a Dashboard administrator.
	Name                string                                                      `json:"name,omitempty"`                // Name of the user. Only required If the user is not a Dashboard administrator.
	Password            string                                                      `json:"password,omitempty"`            // The password for this user account. Only required If the user is not a Dashboard administrator.
}
type RequestNetworksCreateNetworkMerakiAuthUserAuthorizations struct {
	ExpiresAt  string `json:"expiresAt,omitempty"`  // Date for authorization to expire. Set to 'Never' for the authorization to not expire, which is the default.
	SSIDNumber *int   `json:"ssidNumber,omitempty"` // Required for wireless networks. The SSID for which the user is being authorized, which must be configured for the user's given accountType.
}
type RequestNetworksUpdateNetworkMerakiAuthUser struct {
	Authorizations      *[]RequestNetworksUpdateNetworkMerakiAuthUserAuthorizations `json:"authorizations,omitempty"`      // Authorization zones and expiration dates for the user.
	EmailPasswordToUser *bool                                                       `json:"emailPasswordToUser,omitempty"` // Whether or not Meraki should email the password to user. Default is false.
	Name                string                                                      `json:"name,omitempty"`                // Name of the user. Only allowed If the user is not Dashboard administrator.
	Password            string                                                      `json:"password,omitempty"`            // The password for this user account. Only allowed If the user is not Dashboard administrator.
}
type RequestNetworksUpdateNetworkMerakiAuthUserAuthorizations struct {
	ExpiresAt  string `json:"expiresAt,omitempty"`  // Date for authorization to expire. Default is for authorization to not expire.
	SSIDNumber *int   `json:"ssidNumber,omitempty"` // SSID for which the user is being authorized
}
type RequestNetworksCreateNetworkMqttBroker struct {
	Authentication *RequestNetworksCreateNetworkMqttBrokerAuthentication `json:"authentication,omitempty"` // Authentication settings of the MQTT broker
	Host           string                                                `json:"host,omitempty"`           // Host name/IP address where the MQTT broker runs.
	Name           string                                                `json:"name,omitempty"`           // Name of the MQTT broker.
	Port           *int                                                  `json:"port,omitempty"`           // Host port though which the MQTT broker can be reached.
	Security       *RequestNetworksCreateNetworkMqttBrokerSecurity       `json:"security,omitempty"`       // Security settings of the MQTT broker.
}
type RequestNetworksCreateNetworkMqttBrokerAuthentication interface{}
type RequestNetworksCreateNetworkMqttBrokerSecurity struct {
	Mode     string                                                  `json:"mode,omitempty"`     // Security protocol of the MQTT broker.
	Security *RequestNetworksCreateNetworkMqttBrokerSecuritySecurity `json:"security,omitempty"` // TLS settings of the MQTT broker.
}
type RequestNetworksCreateNetworkMqttBrokerSecuritySecurity struct {
	CaCertificate   string `json:"caCertificate,omitempty"`   // CA Certificate of the MQTT broker.
	VerifyHostnames *bool  `json:"verifyHostnames,omitempty"` // Whether the TLS hostname verification is enabled for the MQTT broker.
}
type RequestNetworksUpdateNetworkMqttBroker struct {
	Authentication *RequestNetworksUpdateNetworkMqttBrokerAuthentication `json:"authentication,omitempty"` // Authentication settings of the MQTT broker
	Host           string                                                `json:"host,omitempty"`           // Host name/IP address where the MQTT broker runs.
	Name           string                                                `json:"name,omitempty"`           // Name of the MQTT broker.
	Port           *int                                                  `json:"port,omitempty"`           // Host port though which the MQTT broker can be reached.
	Security       *RequestNetworksUpdateNetworkMqttBrokerSecurity       `json:"security,omitempty"`       // Security settings of the MQTT broker.
}
type RequestNetworksUpdateNetworkMqttBrokerAuthentication interface{}
type RequestNetworksUpdateNetworkMqttBrokerSecurity struct {
	Mode     string                                                  `json:"mode,omitempty"`     // Security protocol of the MQTT broker.
	Security *RequestNetworksUpdateNetworkMqttBrokerSecuritySecurity `json:"security,omitempty"` // TLS settings of the MQTT broker.
}
type RequestNetworksUpdateNetworkMqttBrokerSecuritySecurity struct {
	CaCertificate   string `json:"caCertificate,omitempty"`   // CA Certificate of the MQTT broker.
	VerifyHostnames *bool  `json:"verifyHostnames,omitempty"` // Whether the TLS hostname verification is enabled for the MQTT broker.
}
type RequestNetworksUpdateNetworkNetflow struct {
	CollectorIP      string `json:"collectorIp,omitempty"`      // The IPv4 address of the NetFlow collector.
	CollectorPort    *int   `json:"collectorPort,omitempty"`    // The port that the NetFlow collector will be listening on.
	EtaDstPort       *int   `json:"etaDstPort,omitempty"`       // The port that the Encrypted Traffic Analytics collector will be listening on.
	EtaEnabled       *bool  `json:"etaEnabled,omitempty"`       // Boolean indicating whether Encrypted Traffic Analytics is enabled (true) or disabled (false).
	ReportingEnabled *bool  `json:"reportingEnabled,omitempty"` // Boolean indicating whether NetFlow traffic reporting is enabled (true) or disabled (false).
}
type RequestNetworksCreateNetworkPiiRequest struct {
	Datasets   []string `json:"datasets,omitempty"`   // The datasets related to the provided key that should be deleted. Only applies to "delete" requests. The value "all" will be expanded to all datasets applicable to this type. The datasets by applicable to each type are: mac (usage, events, traffic), email (users, loginAttempts), username (users, loginAttempts), bluetoothMac (client, connectivity), smDeviceId (device), smUserId (user)
	Email      string   `json:"email,omitempty"`      // The email of a network user account. Only applies to "delete" requests.
	Mac        string   `json:"mac,omitempty"`        // The MAC of a network client device. Applies to both "restrict processing" and "delete" requests.
	SmDeviceID string   `json:"smDeviceId,omitempty"` // The sm_device_id of a Systems Manager device. The only way to "restrict processing" or "delete" a Systems Manager device. Must include "device" in the dataset for a "delete" request to destroy the device.
	SmUserID   string   `json:"smUserId,omitempty"`   // The sm_user_id of a Systems Manager user. The only way to "restrict processing" or "delete" a Systems Manager user. Must include "user" in the dataset for a "delete" request to destroy the user.
	Type       string   `json:"type,omitempty"`       // One of "delete" or "restrict processing"
	Username   string   `json:"username,omitempty"`   // The username of a network log in. Only applies to "delete" requests.
}
type RequestNetworksUpdateNetworkSettings struct {
	LocalStatusPage         *RequestNetworksUpdateNetworkSettingsLocalStatusPage `json:"localStatusPage,omitempty"`         // A hash of Local Status page(s)' authentication options applied to the Network.
	LocalStatusPageEnabled  *bool                                                `json:"localStatusPageEnabled,omitempty"`  // Enables / disables the local device status pages (<a target='_blank' href='http://my.meraki.com/'>my.meraki.com, </a><a target='_blank' href='http://ap.meraki.com/'>ap.meraki.com, </a><a target='_blank' href='http://switch.meraki.com/'>switch.meraki.com, </a><a target='_blank' href='http://wired.meraki.com/'>wired.meraki.com</a>). Optional (defaults to false)
	RemoteStatusPageEnabled *bool                                                `json:"remoteStatusPageEnabled,omitempty"` // Enables / disables access to the device status page (<a target='_blank'>http://[device's LAN IP])</a>. Optional. Can only be set if localStatusPageEnabled is set to true
	SecurePort              *RequestNetworksUpdateNetworkSettingsSecurePort      `json:"securePort,omitempty"`              // A hash of SecureConnect options applied to the Network.
}
type RequestNetworksUpdateNetworkSettingsLocalStatusPage struct {
	Authentication *RequestNetworksUpdateNetworkSettingsLocalStatusPageAuthentication `json:"authentication,omitempty"` // A hash of Local Status page(s)' authentication options applied to the Network.
}
type RequestNetworksUpdateNetworkSettingsLocalStatusPageAuthentication struct {
	Enabled  *bool  `json:"enabled,omitempty"`  // Enables / disables the authentication on Local Status page(s).
	Password string `json:"password,omitempty"` // The password used for Local Status Page(s). Set this to null to clear the password.
}
type RequestNetworksUpdateNetworkSettingsSecurePort struct {
	Enabled *bool `json:"enabled,omitempty"` // Enables / disables SecureConnect on the network. Optional.
}
type RequestNetworksUpdateNetworkSNMP struct {
	Access          string                                   `json:"access,omitempty"`          // The type of SNMP access. Can be one of 'none' (disabled), 'community' (V1/V2c), or 'users' (V3).
	CommunityString string                                   `json:"communityString,omitempty"` // The SNMP community string. Only relevant if 'access' is set to 'community'.
	Users           *[]RequestNetworksUpdateNetworkSNMPUsers `json:"users,omitempty"`           // The list of SNMP users. Only relevant if 'access' is set to 'users'.
}
type RequestNetworksUpdateNetworkSNMPUsers struct {
	Passphrase string `json:"passphrase,omitempty"` // The passphrase for the SNMP user. Required.
	Username   string `json:"username,omitempty"`   // The username for the SNMP user. Required.
}
type RequestNetworksUpdateNetworkSyslogServers struct {
	Servers *[]RequestNetworksUpdateNetworkSyslogServersServers `json:"servers,omitempty"` // A list of the syslog servers for this network
}
type RequestNetworksUpdateNetworkSyslogServersServers struct {
	Host  string   `json:"host,omitempty"`  // The IP address of the syslog server
	Port  string   `json:"port,omitempty"`  // The port of the syslog server
	Roles []string `json:"roles,omitempty"` // A list of roles for the syslog server. Options (case-insensitive): 'Wireless event log', 'Appliance event log', 'Switch event log', 'Air Marshal events', 'Flows', 'URLs', 'IDS alerts', 'Security events'
}
type RequestNetworksUpdateNetworkTrafficAnalysis struct {
	CustomPieChartItems *[]RequestNetworksUpdateNetworkTrafficAnalysisCustomPieChartItems `json:"customPieChartItems,omitempty"` // The list of items that make up the custom pie chart for traffic reporting.
	Mode                string                                                            `json:"mode,omitempty"`                //     The traffic analysis mode for the network. Can be one of 'disabled' (do not collect traffic types),     'basic' (collect generic traffic categories), or 'detailed' (collect destination hostnames).
}
type RequestNetworksUpdateNetworkTrafficAnalysisCustomPieChartItems struct {
	Name  string `json:"name,omitempty"`  // The name of the custom pie chart item.
	Type  string `json:"type,omitempty"`  //     The signature type for the custom pie chart item. Can be one of 'host', 'port' or 'ipRange'.
	Value string `json:"value,omitempty"` //     The value of the custom pie chart item. Valid syntax depends on the signature type of the chart item     (see sample request/response for more details).
}
type RequestNetworksUnbindNetwork struct {
	RetainConfigs *bool `json:"retainConfigs,omitempty"` // Optional boolean to retain all the current configs given by the template.
}
type RequestNetworksCreateNetworkWebhooksHTTPServer struct {
	Name            string                                                         `json:"name,omitempty"`            // A name for easy reference to the HTTP server
	PayloadTemplate *RequestNetworksCreateNetworkWebhooksHTTPServerPayloadTemplate `json:"payloadTemplate,omitempty"` // The payload template to use when posting data to the HTTP server.
	SharedSecret    string                                                         `json:"sharedSecret,omitempty"`    // A shared secret that will be included in POSTs sent to the HTTP server. This secret can be used to verify that the request was sent by Meraki.
	URL             string                                                         `json:"url,omitempty"`             // The URL of the HTTP server. Once set, cannot be updated.
}
type RequestNetworksCreateNetworkWebhooksHTTPServerPayloadTemplate struct {
	Name              string `json:"name,omitempty"`              // The name of the payload template.
	PayloadTemplateID string `json:"payloadTemplateId,omitempty"` // The ID of the payload template. Defaults to 'wpt_00001' for the Meraki template. For Meraki-included templates: for the Webex (included) template use 'wpt_00002'; for the Slack (included) template use 'wpt_00003'; for the Microsoft Teams (included) template use 'wpt_00004'; for the ServiceNow (included) template use 'wpt_00006'
}
type RequestNetworksUpdateNetworkWebhooksHTTPServer struct {
	Name            string                                                         `json:"name,omitempty"`            // A name for easy reference to the HTTP server
	PayloadTemplate *RequestNetworksUpdateNetworkWebhooksHTTPServerPayloadTemplate `json:"payloadTemplate,omitempty"` // The payload template to use when posting data to the HTTP server.
	SharedSecret    string                                                         `json:"sharedSecret,omitempty"`    // A shared secret that will be included in POSTs sent to the HTTP server. This secret can be used to verify that the request was sent by Meraki.
}
type RequestNetworksUpdateNetworkWebhooksHTTPServerPayloadTemplate struct {
	PayloadTemplateID string `json:"payloadTemplateId,omitempty"` // The ID of the payload template. Defaults to 'wpt_00001' for the Meraki template. For Meraki-included templates: for the Webex (included) template use 'wpt_00002'; for the Slack (included) template use 'wpt_00003'; for the Microsoft Teams (included) template use 'wpt_00004'; for the ServiceNow (included) template use 'wpt_00006'
}
type RequestNetworksCreateNetworkWebhooksPayloadTemplate struct {
	Body        string                                                        `json:"body,omitempty"`        // The liquid template used for the body of the webhook message. Either *body* or *bodyFile* must be specified.
	BodyFile    string                                                        `json:"bodyFile,omitempty"`    // A file containing liquid template used for the body of the webhook message. Either *body* or *bodyFile* must be specified.
	Headers     *[]RequestNetworksCreateNetworkWebhooksPayloadTemplateHeaders `json:"headers,omitempty"`     // The liquid template used with the webhook headers.
	HeadersFile string                                                        `json:"headersFile,omitempty"` // A file containing the liquid template used with the webhook headers.
	Name        string                                                        `json:"name,omitempty"`        // The name of the new template
}
type RequestNetworksCreateNetworkWebhooksPayloadTemplateHeaders struct {
	Name     string `json:"name,omitempty"`     // The name of the header template
	Template string `json:"template,omitempty"` // The liquid template for the headers
}
type RequestNetworksUpdateNetworkWebhooksPayloadTemplate struct {
	Body        string                                                        `json:"body,omitempty"`        // The liquid template used for the body of the webhook message.
	BodyFile    string                                                        `json:"bodyFile,omitempty"`    // A file containing liquid template used for the body of the webhook message.
	Headers     *[]RequestNetworksUpdateNetworkWebhooksPayloadTemplateHeaders `json:"headers,omitempty"`     // The liquid template used with the webhook headers.
	HeadersFile string                                                        `json:"headersFile,omitempty"` // A file containing the liquid template used with the webhook headers.
	Name        string                                                        `json:"name,omitempty"`        // The name of the template
}
type RequestNetworksUpdateNetworkWebhooksPayloadTemplateHeaders struct {
	Name     string `json:"name,omitempty"`     // The name of the header template
	Template string `json:"template,omitempty"` // The liquid template for the headers
}
type RequestNetworksCreateNetworkWebhooksWebhookTest struct {
	AlertTypeID         string `json:"alertTypeId,omitempty"`         // The type of alert which the test webhook will send. Optional. Defaults to power_supply_down.
	PayloadTemplateID   string `json:"payloadTemplateId,omitempty"`   // The ID of the payload template of the test webhook. Defaults to the HTTP server's template ID if one exists for the given URL, or Generic template ID otherwise
	PayloadTemplateName string `json:"payloadTemplateName,omitempty"` // The name of the payload template.
	SharedSecret        string `json:"sharedSecret,omitempty"`        // The shared secret the test webhook will send. Optional. Defaults to an empty string.
	URL                 string `json:"url,omitempty"`                 // The URL where the test webhook will be sent
}
type RequestNetworksCombineOrganizationNetworks struct {
	EnrollmentString string   `json:"enrollmentString,omitempty"` // A unique identifier which can be used for device enrollment or easy access through the Meraki SM Registration page or the Self Service Portal. Please note that changing this field may cause existing bookmarks to break. All networks that are part of this combined network will have their enrollment string appended by '-network_type'. If left empty, all exisitng enrollment strings will be deleted.
	Name             string   `json:"name,omitempty"`             // The name of the combined network
	NetworkIDs       []string `json:"networkIds,omitempty"`       // A list of the network IDs that will be combined. If an ID of a combined network is included in this list, the other networks in the list will be grouped into that network
}

//GetNetwork Return a network
/* Return a network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network
*/
func (s *NetworksService) GetNetwork(networkID string) (*ResponseNetworksGetNetwork, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworksGetNetwork{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetwork")
	}

	result := response.Result().(*ResponseNetworksGetNetwork)
	return result, response, err

}

//GetNetworkAlertsHistory Return the alert history for this network
/* Return the alert history for this network

@param networkID networkId path parameter. Network ID
@param getNetworkAlertsHistoryQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-alerts-history
*/
func (s *NetworksService) GetNetworkAlertsHistory(networkID string, getNetworkAlertsHistoryQueryParams *GetNetworkAlertsHistoryQueryParams) (*ResponseNetworksGetNetworkAlertsHistory, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/alerts/history"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(getNetworkAlertsHistoryQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworksGetNetworkAlertsHistory{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkAlertsHistory")
	}

	result := response.Result().(*ResponseNetworksGetNetworkAlertsHistory)
	return result, response, err

}

//GetNetworkAlertsSettings Return the alert configuration for this network
/* Return the alert configuration for this network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-alerts-settings
*/
func (s *NetworksService) GetNetworkAlertsSettings(networkID string) (*ResponseNetworksGetNetworkAlertsSettings, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/alerts/settings"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworksGetNetworkAlertsSettings{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkAlertsSettings")
	}

	result := response.Result().(*ResponseNetworksGetNetworkAlertsSettings)
	return result, response, err

}

//GetNetworkBluetoothClients List the Bluetooth clients seen by APs in this network
/* List the Bluetooth clients seen by APs in this network

@param networkID networkId path parameter. Network ID
@param getNetworkBluetoothClientsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-bluetooth-clients
*/
func (s *NetworksService) GetNetworkBluetoothClients(networkID string, getNetworkBluetoothClientsQueryParams *GetNetworkBluetoothClientsQueryParams) (*ResponseNetworksGetNetworkBluetoothClients, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/bluetoothClients"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(getNetworkBluetoothClientsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworksGetNetworkBluetoothClients{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkBluetoothClients")
	}

	result := response.Result().(*ResponseNetworksGetNetworkBluetoothClients)
	return result, response, err

}

//GetNetworkBluetoothClient Return a Bluetooth client
/* Return a Bluetooth client. Bluetooth clients can be identified by their ID or their MAC.

@param networkID networkId path parameter. Network ID
@param bluetoothClientID bluetoothClientId path parameter. Bluetooth client ID
@param getNetworkBluetoothClientQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-bluetooth-client
*/
func (s *NetworksService) GetNetworkBluetoothClient(networkID string, bluetoothClientID string, getNetworkBluetoothClientQueryParams *GetNetworkBluetoothClientQueryParams) (*ResponseNetworksGetNetworkBluetoothClient, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/bluetoothClients/{bluetoothClientId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{bluetoothClientId}", fmt.Sprintf("%v", bluetoothClientID), -1)

	queryString, _ := query.Values(getNetworkBluetoothClientQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworksGetNetworkBluetoothClient{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkBluetoothClient")
	}

	result := response.Result().(*ResponseNetworksGetNetworkBluetoothClient)
	return result, response, err

}

//GetNetworkClients List the clients that have used this network in the timespan
/* List the clients that have used this network in the timespan

@param networkID networkId path parameter. Network ID
@param getNetworkClientsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-clients
*/
func (s *NetworksService) GetNetworkClients(networkID string, getNetworkClientsQueryParams *GetNetworkClientsQueryParams) (*[]ResponseNetworksGetNetworkClients, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/clients"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(getNetworkClientsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&[]ResponseNetworksGetNetworkClients{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkClients")
	}

	result := response.Result().(*[]ResponseNetworksGetNetworkClients)
	return result, response, err

}

//GetNetworkClientsApplicationUsage Return the application usage data for clients
/* Return the application usage data for clients. Usage data is in kilobytes. Clients can be identified by client keys or either the MACs or IPs depending on whether the network uses Track-by-IP.

@param networkID networkId path parameter. Network ID
@param getNetworkClientsApplicationUsageQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-clients-application-usage
*/
func (s *NetworksService) GetNetworkClientsApplicationUsage(networkID string, getNetworkClientsApplicationUsageQueryParams *GetNetworkClientsApplicationUsageQueryParams) (*ResponseNetworksGetNetworkClientsApplicationUsage, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/clients/applicationUsage"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(getNetworkClientsApplicationUsageQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworksGetNetworkClientsApplicationUsage{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkClientsApplicationUsage")
	}

	result := response.Result().(*ResponseNetworksGetNetworkClientsApplicationUsage)
	return result, response, err

}

//GetNetworkClientsBandwidthUsageHistory Returns a timeseries of total traffic consumption rates for all clients on a network within a given timespan, in megabits per second.
/* Returns a timeseries of total traffic consumption rates for all clients on a network within a given timespan, in megabits per second.

@param networkID networkId path parameter. Network ID
@param getNetworkClientsBandwidthUsageHistoryQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-clients-bandwidth-usage-history
*/
func (s *NetworksService) GetNetworkClientsBandwidthUsageHistory(networkID string, getNetworkClientsBandwidthUsageHistoryQueryParams *GetNetworkClientsBandwidthUsageHistoryQueryParams) (*ResponseNetworksGetNetworkClientsBandwidthUsageHistory, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/clients/bandwidthUsageHistory"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(getNetworkClientsBandwidthUsageHistoryQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworksGetNetworkClientsBandwidthUsageHistory{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkClientsBandwidthUsageHistory")
	}

	result := response.Result().(*ResponseNetworksGetNetworkClientsBandwidthUsageHistory)
	return result, response, err

}

//GetNetworkClientsOverview Return overview statistics for network clients
/* Return overview statistics for network clients

@param networkID networkId path parameter. Network ID
@param getNetworkClientsOverviewQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-clients-overview
*/
func (s *NetworksService) GetNetworkClientsOverview(networkID string, getNetworkClientsOverviewQueryParams *GetNetworkClientsOverviewQueryParams) (*ResponseNetworksGetNetworkClientsOverview, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/clients/overview"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(getNetworkClientsOverviewQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworksGetNetworkClientsOverview{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkClientsOverview")
	}

	result := response.Result().(*ResponseNetworksGetNetworkClientsOverview)
	return result, response, err

}

//GetNetworkClientsUsageHistories Return the usage histories for clients
/* Return the usage histories for clients. Usage data is in kilobytes. Clients can be identified by client keys or either the MACs or IPs depending on whether the network uses Track-by-IP.

@param networkID networkId path parameter. Network ID
@param getNetworkClientsUsageHistoriesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-clients-usage-histories
*/
func (s *NetworksService) GetNetworkClientsUsageHistories(networkID string, getNetworkClientsUsageHistoriesQueryParams *GetNetworkClientsUsageHistoriesQueryParams) (*ResponseNetworksGetNetworkClientsUsageHistories, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/clients/usageHistories"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(getNetworkClientsUsageHistoriesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworksGetNetworkClientsUsageHistories{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkClientsUsageHistories")
	}

	result := response.Result().(*ResponseNetworksGetNetworkClientsUsageHistories)
	return result, response, err

}

//GetNetworkClient Return the client associated with the given identifier
/* Return the client associated with the given identifier. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.

@param networkID networkId path parameter. Network ID
@param clientID clientId path parameter. Client ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-client
*/
func (s *NetworksService) GetNetworkClient(networkID string, clientID string) (*ResponseNetworksGetNetworkClient, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/clients/{clientId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{clientId}", fmt.Sprintf("%v", clientID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworksGetNetworkClient{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkClient")
	}

	result := response.Result().(*ResponseNetworksGetNetworkClient)
	return result, response, err

}

//GetNetworkClientPolicy Return the policy assigned to a client on the network
/* Return the policy assigned to a client on the network. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.

@param networkID networkId path parameter. Network ID
@param clientID clientId path parameter. Client ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-client-policy
*/
func (s *NetworksService) GetNetworkClientPolicy(networkID string, clientID string) (*ResponseNetworksGetNetworkClientPolicy, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/clients/{clientId}/policy"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{clientId}", fmt.Sprintf("%v", clientID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworksGetNetworkClientPolicy{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkClientPolicy")
	}

	result := response.Result().(*ResponseNetworksGetNetworkClientPolicy)
	return result, response, err

}

//GetNetworkClientSplashAuthorizationStatus Return the splash authorization for a client, for each SSID they've associated with through splash
/* Return the splash authorization for a client, for each SSID they've associated with through splash. Only enabled SSIDs with Click-through splash enabled will be included. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.

@param networkID networkId path parameter. Network ID
@param clientID clientId path parameter. Client ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-client-splash-authorization-status
*/
func (s *NetworksService) GetNetworkClientSplashAuthorizationStatus(networkID string, clientID string) (*ResponseNetworksGetNetworkClientSplashAuthorizationStatus, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/clients/{clientId}/splashAuthorizationStatus"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{clientId}", fmt.Sprintf("%v", clientID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworksGetNetworkClientSplashAuthorizationStatus{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkClientSplashAuthorizationStatus")
	}

	result := response.Result().(*ResponseNetworksGetNetworkClientSplashAuthorizationStatus)
	return result, response, err

}

//GetNetworkClientTrafficHistory Return the client's network traffic data over time
/* Return the client's network traffic data over time. Usage data is in kilobytes. This endpoint requires detailed traffic analysis to be enabled on the Network-wide > General page. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.

@param networkID networkId path parameter. Network ID
@param clientID clientId path parameter. Client ID
@param getNetworkClientTrafficHistoryQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-client-traffic-history
*/
func (s *NetworksService) GetNetworkClientTrafficHistory(networkID string, clientID string, getNetworkClientTrafficHistoryQueryParams *GetNetworkClientTrafficHistoryQueryParams) (*ResponseNetworksGetNetworkClientTrafficHistory, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/clients/{clientId}/trafficHistory"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{clientId}", fmt.Sprintf("%v", clientID), -1)

	queryString, _ := query.Values(getNetworkClientTrafficHistoryQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworksGetNetworkClientTrafficHistory{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkClientTrafficHistory")
	}

	result := response.Result().(*ResponseNetworksGetNetworkClientTrafficHistory)
	return result, response, err

}

//GetNetworkClientUsageHistory Return the client's daily usage history
/* Return the client's daily usage history. Usage data is in kilobytes. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.

@param networkID networkId path parameter. Network ID
@param clientID clientId path parameter. Client ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-client-usage-history
*/
func (s *NetworksService) GetNetworkClientUsageHistory(networkID string, clientID string) (*ResponseNetworksGetNetworkClientUsageHistory, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/clients/{clientId}/usageHistory"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{clientId}", fmt.Sprintf("%v", clientID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworksGetNetworkClientUsageHistory{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkClientUsageHistory")
	}

	result := response.Result().(*ResponseNetworksGetNetworkClientUsageHistory)
	return result, response, err

}

//GetNetworkDevices List the devices in a network
/* List the devices in a network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-devices
*/
func (s *NetworksService) GetNetworkDevices(networkID string) (*ResponseNetworksGetNetworkDevices, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/devices"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworksGetNetworkDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkDevices")
	}

	result := response.Result().(*ResponseNetworksGetNetworkDevices)
	return result, response, err

}

//GetNetworkEvents List the events for the network
/* List the events for the network

@param networkID networkId path parameter. Network ID
@param getNetworkEventsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-events
*/
func (s *NetworksService) GetNetworkEvents(networkID string, getNetworkEventsQueryParams *GetNetworkEventsQueryParams) (*ResponseNetworksGetNetworkEvents, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/events"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(getNetworkEventsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworksGetNetworkEvents{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkEvents")
	}

	result := response.Result().(*ResponseNetworksGetNetworkEvents)
	return result, response, err

}

//GetNetworkEventsEventTypes List the event type to human-readable description
/* List the event type to human-readable description

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-events-event-types
*/
func (s *NetworksService) GetNetworkEventsEventTypes(networkID string) (*ResponseNetworksGetNetworkEventsEventTypes, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/events/eventTypes"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworksGetNetworkEventsEventTypes{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkEventsEventTypes")
	}

	result := response.Result().(*ResponseNetworksGetNetworkEventsEventTypes)
	return result, response, err

}

//GetNetworkFirmwareUpgrades Get firmware upgrade information for a network
/* Get firmware upgrade information for a network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-firmware-upgrades
*/
func (s *NetworksService) GetNetworkFirmwareUpgrades(networkID string) (*ResponseNetworksGetNetworkFirmwareUpgrades, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/firmwareUpgrades"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworksGetNetworkFirmwareUpgrades{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkFirmwareUpgrades")
	}

	result := response.Result().(*ResponseNetworksGetNetworkFirmwareUpgrades)
	return result, response, err

}

//GetNetworkFirmwareUpgradesStagedEvents Get the Staged Upgrade Event from a network
/* Get the Staged Upgrade Event from a network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-firmware-upgrades-staged-events
*/
func (s *NetworksService) GetNetworkFirmwareUpgradesStagedEvents(networkID string) (*ResponseNetworksGetNetworkFirmwareUpgradesStagedEvents, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/firmwareUpgrades/staged/events"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworksGetNetworkFirmwareUpgradesStagedEvents{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkFirmwareUpgradesStagedEvents")
	}

	result := response.Result().(*ResponseNetworksGetNetworkFirmwareUpgradesStagedEvents)
	return result, response, err

}

//GetNetworkFirmwareUpgradesStagedGroups List of Staged Upgrade Groups in a network
/* List of Staged Upgrade Groups in a network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-firmware-upgrades-staged-groups
*/
func (s *NetworksService) GetNetworkFirmwareUpgradesStagedGroups(networkID string) (*ResponseNetworksGetNetworkFirmwareUpgradesStagedGroups, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/firmwareUpgrades/staged/groups"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworksGetNetworkFirmwareUpgradesStagedGroups{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkFirmwareUpgradesStagedGroups")
	}

	result := response.Result().(*ResponseNetworksGetNetworkFirmwareUpgradesStagedGroups)
	return result, response, err

}

//GetNetworkFirmwareUpgradesStagedGroup Get a Staged Upgrade Group from a network
/* Get a Staged Upgrade Group from a network

@param networkID networkId path parameter. Network ID
@param groupID groupId path parameter. Group ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-firmware-upgrades-staged-group
*/
func (s *NetworksService) GetNetworkFirmwareUpgradesStagedGroup(networkID string, groupID string) (*ResponseNetworksGetNetworkFirmwareUpgradesStagedGroup, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/firmwareUpgrades/staged/groups/{groupId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{groupId}", fmt.Sprintf("%v", groupID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworksGetNetworkFirmwareUpgradesStagedGroup{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkFirmwareUpgradesStagedGroup")
	}

	result := response.Result().(*ResponseNetworksGetNetworkFirmwareUpgradesStagedGroup)
	return result, response, err

}

//GetNetworkFirmwareUpgradesStagedStages Order of Staged Upgrade Groups in a network
/* Order of Staged Upgrade Groups in a network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-firmware-upgrades-staged-stages
*/
func (s *NetworksService) GetNetworkFirmwareUpgradesStagedStages(networkID string) (*ResponseNetworksGetNetworkFirmwareUpgradesStagedStages, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/firmwareUpgrades/staged/stages"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworksGetNetworkFirmwareUpgradesStagedStages{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkFirmwareUpgradesStagedStages")
	}

	result := response.Result().(*ResponseNetworksGetNetworkFirmwareUpgradesStagedStages)
	return result, response, err

}

//GetNetworkFloorPlans List the floor plans that belong to your network
/* List the floor plans that belong to your network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-floor-plans
*/
func (s *NetworksService) GetNetworkFloorPlans(networkID string) (*ResponseNetworksGetNetworkFloorPlans, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/floorPlans"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworksGetNetworkFloorPlans{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkFloorPlans")
	}

	result := response.Result().(*ResponseNetworksGetNetworkFloorPlans)
	return result, response, err

}

//GetNetworkFloorPlan Find a floor plan by ID
/* Find a floor plan by ID

@param networkID networkId path parameter. Network ID
@param floorPlanID floorPlanId path parameter. Floor plan ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-floor-plan
*/
func (s *NetworksService) GetNetworkFloorPlan(networkID string, floorPlanID string) (*ResponseNetworksGetNetworkFloorPlan, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/floorPlans/{floorPlanId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{floorPlanId}", fmt.Sprintf("%v", floorPlanID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworksGetNetworkFloorPlan{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkFloorPlan")
	}

	result := response.Result().(*ResponseNetworksGetNetworkFloorPlan)
	return result, response, err

}

//GetNetworkGroupPolicies List the group policies in a network
/* List the group policies in a network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-group-policies
*/
func (s *NetworksService) GetNetworkGroupPolicies(networkID string) (*ResponseNetworksGetNetworkGroupPolicies, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/groupPolicies"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworksGetNetworkGroupPolicies{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkGroupPolicies")
	}

	result := response.Result().(*ResponseNetworksGetNetworkGroupPolicies)
	return result, response, err

}

//GetNetworkGroupPolicy Display a group policy
/* Display a group policy

@param networkID networkId path parameter. Network ID
@param groupPolicyID groupPolicyId path parameter. Group policy ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-group-policy
*/
func (s *NetworksService) GetNetworkGroupPolicy(networkID string, groupPolicyID string) (*ResponseNetworksGetNetworkGroupPolicy, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/groupPolicies/{groupPolicyId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{groupPolicyId}", fmt.Sprintf("%v", groupPolicyID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworksGetNetworkGroupPolicy{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkGroupPolicy")
	}

	result := response.Result().(*ResponseNetworksGetNetworkGroupPolicy)
	return result, response, err

}

//GetNetworkHealthAlerts Return all global alerts on this network
/* Return all global alerts on this network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-health-alerts
*/
func (s *NetworksService) GetNetworkHealthAlerts(networkID string) (*ResponseNetworksGetNetworkHealthAlerts, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/health/alerts"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworksGetNetworkHealthAlerts{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkHealthAlerts")
	}

	result := response.Result().(*ResponseNetworksGetNetworkHealthAlerts)
	return result, response, err

}

//GetNetworkMerakiAuthUsers List the users configured under Meraki Authentication for a network (splash guest or RADIUS users for a wireless network, or client VPN users for a wired network)
/* List the users configured under Meraki Authentication for a network (splash guest or RADIUS users for a wireless network, or client VPN users for a wired network)

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-meraki-auth-users
*/
func (s *NetworksService) GetNetworkMerakiAuthUsers(networkID string) (*ResponseNetworksGetNetworkMerakiAuthUsers, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/merakiAuthUsers"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworksGetNetworkMerakiAuthUsers{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkMerakiAuthUsers")
	}

	result := response.Result().(*ResponseNetworksGetNetworkMerakiAuthUsers)
	return result, response, err

}

//GetNetworkMerakiAuthUser Return the Meraki Auth splash guest, RADIUS, or client VPN user
/* Return the Meraki Auth splash guest, RADIUS, or client VPN user

@param networkID networkId path parameter. Network ID
@param merakiAuthUserID merakiAuthUserId path parameter. Meraki auth user ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-meraki-auth-user
*/
func (s *NetworksService) GetNetworkMerakiAuthUser(networkID string, merakiAuthUserID string) (*ResponseNetworksGetNetworkMerakiAuthUser, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/merakiAuthUsers/{merakiAuthUserId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{merakiAuthUserId}", fmt.Sprintf("%v", merakiAuthUserID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworksGetNetworkMerakiAuthUser{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkMerakiAuthUser")
	}

	result := response.Result().(*ResponseNetworksGetNetworkMerakiAuthUser)
	return result, response, err

}

//GetNetworkMqttBrokers List the MQTT brokers for this network
/* List the MQTT brokers for this network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-mqtt-brokers
*/
func (s *NetworksService) GetNetworkMqttBrokers(networkID string) (*ResponseNetworksGetNetworkMqttBrokers, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/mqttBrokers"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworksGetNetworkMqttBrokers{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkMqttBrokers")
	}

	result := response.Result().(*ResponseNetworksGetNetworkMqttBrokers)
	return result, response, err

}

//GetNetworkMqttBroker Return an MQTT broker
/* Return an MQTT broker

@param networkID networkId path parameter. Network ID
@param mqttBrokerID mqttBrokerId path parameter. Mqtt broker ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-mqtt-broker
*/
func (s *NetworksService) GetNetworkMqttBroker(networkID string, mqttBrokerID string) (*ResponseNetworksGetNetworkMqttBroker, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/mqttBrokers/{mqttBrokerId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{mqttBrokerId}", fmt.Sprintf("%v", mqttBrokerID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworksGetNetworkMqttBroker{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkMqttBroker")
	}

	result := response.Result().(*ResponseNetworksGetNetworkMqttBroker)
	return result, response, err

}

//GetNetworkNetflow Return the NetFlow traffic reporting settings for a network
/* Return the NetFlow traffic reporting settings for a network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-netflow
*/
func (s *NetworksService) GetNetworkNetflow(networkID string) (*ResponseNetworksGetNetworkNetflow, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/netflow"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworksGetNetworkNetflow{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkNetflow")
	}

	result := response.Result().(*ResponseNetworksGetNetworkNetflow)
	return result, response, err

}

//GetNetworkNetworkHealthChannelUtilization Get the channel utilization over each radio for all APs in a network.
/* Get the channel utilization over each radio for all APs in a network.

@param networkID networkId path parameter. Network ID
@param getNetworkNetworkHealthChannelUtilizationQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-network-health-channel-utilization
*/
func (s *NetworksService) GetNetworkNetworkHealthChannelUtilization(networkID string, getNetworkNetworkHealthChannelUtilizationQueryParams *GetNetworkNetworkHealthChannelUtilizationQueryParams) (*ResponseNetworksGetNetworkNetworkHealthChannelUtilization, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/networkHealth/channelUtilization"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(getNetworkNetworkHealthChannelUtilizationQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworksGetNetworkNetworkHealthChannelUtilization{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkNetworkHealthChannelUtilization")
	}

	result := response.Result().(*ResponseNetworksGetNetworkNetworkHealthChannelUtilization)
	return result, response, err

}

//GetNetworkPiiPiiKeys List the keys required to access Personally Identifiable Information (PII) for a given identifier
/* List the keys required to access Personally Identifiable Information (PII) for a given identifier. Exactly one identifier will be accepted. If the organization contains org-wide Systems Manager users matching the key provided then there will be an entry with the key "0" containing the applicable keys.

## ALTERNATE PATH

***
/organizations/{organizationId}/pii/piiKeys
***

@param networkID networkId path parameter. Network ID
@param getNetworkPiiPiiKeysQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-pii-pii-keys
*/
func (s *NetworksService) GetNetworkPiiPiiKeys(networkID string, getNetworkPiiPiiKeysQueryParams *GetNetworkPiiPiiKeysQueryParams) (*ResponseNetworksGetNetworkPiiPiiKeys, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/pii/piiKeys"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(getNetworkPiiPiiKeysQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworksGetNetworkPiiPiiKeys{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkPiiPiiKeys")
	}

	result := response.Result().(*ResponseNetworksGetNetworkPiiPiiKeys)
	return result, response, err

}

//GetNetworkPiiRequests List the PII requests for this network or organization
/* List the PII requests for this network or organization

## ALTERNATE PATH

***
/organizations/{organizationId}/pii/requests
***

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-pii-requests
*/
func (s *NetworksService) GetNetworkPiiRequests(networkID string) (*ResponseNetworksGetNetworkPiiRequests, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/pii/requests"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworksGetNetworkPiiRequests{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkPiiRequests")
	}

	result := response.Result().(*ResponseNetworksGetNetworkPiiRequests)
	return result, response, err

}

//GetNetworkPiiRequest Return a PII request
/* Return a PII request

## ALTERNATE PATH

***
/organizations/{organizationId}/pii/requests/{requestId}
***

@param networkID networkId path parameter. Network ID
@param requestID requestId path parameter. Request ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-pii-request
*/
func (s *NetworksService) GetNetworkPiiRequest(networkID string, requestID string) (*ResponseNetworksGetNetworkPiiRequest, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/pii/requests/{requestId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{requestId}", fmt.Sprintf("%v", requestID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworksGetNetworkPiiRequest{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkPiiRequest")
	}

	result := response.Result().(*ResponseNetworksGetNetworkPiiRequest)
	return result, response, err

}

//GetNetworkPiiSmDevicesForKey Given a piece of Personally Identifiable Information (PII), return the Systems Manager device ID(s) associated with that identifier
/* Given a piece of Personally Identifiable Information (PII), return the Systems Manager device ID(s) associated with that identifier. These device IDs can be used with the Systems Manager API endpoints to retrieve device details. Exactly one identifier will be accepted.

## ALTERNATE PATH

***
/organizations/{organizationId}/pii/smDevicesForKey
***

@param networkID networkId path parameter. Network ID
@param getNetworkPiiSmDevicesForKeyQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-pii-sm-devices-for-key
*/
func (s *NetworksService) GetNetworkPiiSmDevicesForKey(networkID string, getNetworkPiiSmDevicesForKeyQueryParams *GetNetworkPiiSmDevicesForKeyQueryParams) (*ResponseNetworksGetNetworkPiiSmDevicesForKey, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/pii/smDevicesForKey"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(getNetworkPiiSmDevicesForKeyQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworksGetNetworkPiiSmDevicesForKey{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkPiiSmDevicesForKey")
	}

	result := response.Result().(*ResponseNetworksGetNetworkPiiSmDevicesForKey)
	return result, response, err

}

//GetNetworkPiiSmOwnersForKey Given a piece of Personally Identifiable Information (PII), return the Systems Manager owner ID(s) associated with that identifier
/* Given a piece of Personally Identifiable Information (PII), return the Systems Manager owner ID(s) associated with that identifier. These owner IDs can be used with the Systems Manager API endpoints to retrieve owner details. Exactly one identifier will be accepted.

## ALTERNATE PATH

***
/organizations/{organizationId}/pii/smOwnersForKey
***

@param networkID networkId path parameter. Network ID
@param getNetworkPiiSmOwnersForKeyQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-pii-sm-owners-for-key
*/
func (s *NetworksService) GetNetworkPiiSmOwnersForKey(networkID string, getNetworkPiiSmOwnersForKeyQueryParams *GetNetworkPiiSmOwnersForKeyQueryParams) (*ResponseNetworksGetNetworkPiiSmOwnersForKey, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/pii/smOwnersForKey"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(getNetworkPiiSmOwnersForKeyQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworksGetNetworkPiiSmOwnersForKey{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkPiiSmOwnersForKey")
	}

	result := response.Result().(*ResponseNetworksGetNetworkPiiSmOwnersForKey)
	return result, response, err

}

//GetNetworkPoliciesByClient Get policies for all clients with policies
/* Get policies for all clients with policies

@param networkID networkId path parameter. Network ID
@param getNetworkPoliciesByClientQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-policies-by-client
*/
func (s *NetworksService) GetNetworkPoliciesByClient(networkID string, getNetworkPoliciesByClientQueryParams *GetNetworkPoliciesByClientQueryParams) (*ResponseNetworksGetNetworkPoliciesByClient, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/policies/byClient"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(getNetworkPoliciesByClientQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworksGetNetworkPoliciesByClient{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkPoliciesByClient")
	}

	result := response.Result().(*ResponseNetworksGetNetworkPoliciesByClient)
	return result, response, err

}

//GetNetworkSettings Return the settings for a network
/* Return the settings for a network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-settings
*/
func (s *NetworksService) GetNetworkSettings(networkID string) (*ResponseNetworksGetNetworkSettings, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/settings"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworksGetNetworkSettings{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSettings")
	}

	result := response.Result().(*ResponseNetworksGetNetworkSettings)
	return result, response, err

}

//GetNetworkSNMP Return the SNMP settings for a network
/* Return the SNMP settings for a network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-snmp
*/
func (s *NetworksService) GetNetworkSNMP(networkID string) (*ResponseNetworksGetNetworkSNMP, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/snmp"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworksGetNetworkSNMP{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSnmp")
	}

	result := response.Result().(*ResponseNetworksGetNetworkSNMP)
	return result, response, err

}

//GetNetworkSplashLoginAttempts List the splash login attempts for a network
/* List the splash login attempts for a network

@param networkID networkId path parameter. Network ID
@param getNetworkSplashLoginAttemptsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-splash-login-attempts
*/
func (s *NetworksService) GetNetworkSplashLoginAttempts(networkID string, getNetworkSplashLoginAttemptsQueryParams *GetNetworkSplashLoginAttemptsQueryParams) (*ResponseNetworksGetNetworkSplashLoginAttempts, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/splashLoginAttempts"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(getNetworkSplashLoginAttemptsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworksGetNetworkSplashLoginAttempts{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSplashLoginAttempts")
	}

	result := response.Result().(*ResponseNetworksGetNetworkSplashLoginAttempts)
	return result, response, err

}

//GetNetworkSyslogServers List the syslog servers for a network
/* List the syslog servers for a network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-syslog-servers
*/
func (s *NetworksService) GetNetworkSyslogServers(networkID string) (*ResponseNetworksGetNetworkSyslogServers, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/syslogServers"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworksGetNetworkSyslogServers{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSyslogServers")
	}

	result := response.Result().(*ResponseNetworksGetNetworkSyslogServers)
	return result, response, err

}

//GetNetworkTopologyLinkLayer List the LLDP and CDP information for all discovered devices and connections in a network.
/* List the LLDP and CDP information for all discovered devices and connections in a network.

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-topology-link-layer
*/
func (s *NetworksService) GetNetworkTopologyLinkLayer(networkID string) (*ResponseNetworksGetNetworkTopologyLinkLayer, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/topology/linkLayer"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworksGetNetworkTopologyLinkLayer{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkTopologyLinkLayer")
	}

	result := response.Result().(*ResponseNetworksGetNetworkTopologyLinkLayer)
	return result, response, err

}

//GetNetworkTraffic Return the traffic analysis data for this network
/* Return the traffic analysis data for this network. Traffic analysis with hostname visibility must be enabled on the network.

@param networkID networkId path parameter. Network ID
@param getNetworkTrafficQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-traffic
*/
func (s *NetworksService) GetNetworkTraffic(networkID string, getNetworkTrafficQueryParams *GetNetworkTrafficQueryParams) (*ResponseNetworksGetNetworkTraffic, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/traffic"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(getNetworkTrafficQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworksGetNetworkTraffic{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkTraffic")
	}

	result := response.Result().(*ResponseNetworksGetNetworkTraffic)
	return result, response, err

}

//GetNetworkTrafficAnalysis Return the traffic analysis settings for a network
/* Return the traffic analysis settings for a network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-traffic-analysis
*/
func (s *NetworksService) GetNetworkTrafficAnalysis(networkID string) (*ResponseNetworksGetNetworkTrafficAnalysis, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/trafficAnalysis"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworksGetNetworkTrafficAnalysis{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkTrafficAnalysis")
	}

	result := response.Result().(*ResponseNetworksGetNetworkTrafficAnalysis)
	return result, response, err

}

//GetNetworkTrafficShapingApplicationCategories Returns the application categories for traffic shaping rules.
/* Returns the application categories for traffic shaping rules.

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-traffic-shaping-application-categories
*/
func (s *NetworksService) GetNetworkTrafficShapingApplicationCategories(networkID string) (*ResponseNetworksGetNetworkTrafficShapingApplicationCategories, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/trafficShaping/applicationCategories"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworksGetNetworkTrafficShapingApplicationCategories{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkTrafficShapingApplicationCategories")
	}

	result := response.Result().(*ResponseNetworksGetNetworkTrafficShapingApplicationCategories)
	return result, response, err

}

//GetNetworkTrafficShapingDscpTaggingOptions Returns the available DSCP tagging options for your traffic shaping rules.
/* Returns the available DSCP tagging options for your traffic shaping rules.

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-traffic-shaping-dscp-tagging-options
*/
func (s *NetworksService) GetNetworkTrafficShapingDscpTaggingOptions(networkID string) (*ResponseNetworksGetNetworkTrafficShapingDscpTaggingOptions, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/trafficShaping/dscpTaggingOptions"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworksGetNetworkTrafficShapingDscpTaggingOptions{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkTrafficShapingDscpTaggingOptions")
	}

	result := response.Result().(*ResponseNetworksGetNetworkTrafficShapingDscpTaggingOptions)
	return result, response, err

}

//GetNetworkWebhooksHTTPServers List the HTTP servers for a network
/* List the HTTP servers for a network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-webhooks-http-servers
*/
func (s *NetworksService) GetNetworkWebhooksHTTPServers(networkID string) (*ResponseNetworksGetNetworkWebhooksHTTPServers, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/webhooks/httpServers"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworksGetNetworkWebhooksHTTPServers{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWebhooksHttpServers")
	}

	result := response.Result().(*ResponseNetworksGetNetworkWebhooksHTTPServers)
	return result, response, err

}

//GetNetworkWebhooksHTTPServer Return an HTTP server for a network
/* Return an HTTP server for a network

@param networkID networkId path parameter. Network ID
@param httpServerID httpServerId path parameter. Http server ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-webhooks-http-server
*/
func (s *NetworksService) GetNetworkWebhooksHTTPServer(networkID string, httpServerID string) (*ResponseNetworksGetNetworkWebhooksHTTPServer, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/webhooks/httpServers/{httpServerId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{httpServerId}", fmt.Sprintf("%v", httpServerID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworksGetNetworkWebhooksHTTPServer{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWebhooksHttpServer")
	}

	result := response.Result().(*ResponseNetworksGetNetworkWebhooksHTTPServer)
	return result, response, err

}

//GetNetworkWebhooksPayloadTemplates List the webhook payload templates for a network
/* List the webhook payload templates for a network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-webhooks-payload-templates
*/
func (s *NetworksService) GetNetworkWebhooksPayloadTemplates(networkID string) (*ResponseNetworksGetNetworkWebhooksPayloadTemplates, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/webhooks/payloadTemplates"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworksGetNetworkWebhooksPayloadTemplates{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWebhooksPayloadTemplates")
	}

	result := response.Result().(*ResponseNetworksGetNetworkWebhooksPayloadTemplates)
	return result, response, err

}

//GetNetworkWebhooksPayloadTemplate Get the webhook payload template for a network
/* Get the webhook payload template for a network

@param networkID networkId path parameter. Network ID
@param payloadTemplateID payloadTemplateId path parameter. Payload template ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-webhooks-payload-template
*/
func (s *NetworksService) GetNetworkWebhooksPayloadTemplate(networkID string, payloadTemplateID string) (*ResponseNetworksGetNetworkWebhooksPayloadTemplate, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/webhooks/payloadTemplates/{payloadTemplateId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{payloadTemplateId}", fmt.Sprintf("%v", payloadTemplateID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworksGetNetworkWebhooksPayloadTemplate{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWebhooksPayloadTemplate")
	}

	result := response.Result().(*ResponseNetworksGetNetworkWebhooksPayloadTemplate)
	return result, response, err

}

//GetNetworkWebhooksWebhookTest Return the status of a webhook test for a network
/* Return the status of a webhook test for a network

@param networkID networkId path parameter. Network ID
@param webhookTestID webhookTestId path parameter. Webhook test ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-webhooks-webhook-test
*/
func (s *NetworksService) GetNetworkWebhooksWebhookTest(networkID string, webhookTestID string) (*ResponseNetworksGetNetworkWebhooksWebhookTest, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/webhooks/webhookTests/{webhookTestId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{webhookTestId}", fmt.Sprintf("%v", webhookTestID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworksGetNetworkWebhooksWebhookTest{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWebhooksWebhookTest")
	}

	result := response.Result().(*ResponseNetworksGetNetworkWebhooksWebhookTest)
	return result, response, err

}

//BindNetwork Bind a network to a template.
/* Bind a network to a template.

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!bind-network
*/

func (s *NetworksService) BindNetwork(networkID string, requestNetworksBindNetwork *RequestNetworksBindNetwork) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/bind"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworksBindNetwork).
		// SetResult(&ResponseNetworksBindNetwork{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation BindNetwork")
	}

	return response, err

}

//ProvisionNetworkClients Provisions a client with a name and policy
/* Provisions a client with a name and policy. Clients can be provisioned before they associate to the network.

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!provision-network-clients
*/

func (s *NetworksService) ProvisionNetworkClients(networkID string, requestNetworksProvisionNetworkClients *RequestNetworksProvisionNetworkClients) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/clients/provision"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworksProvisionNetworkClients).
		// SetResult(&ResponseNetworksProvisionNetworkClients{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation ProvisionNetworkClients")
	}

	return response, err

}

//ClaimNetworkDevices Claim devices into a network. (Note: for recently claimed devices, it may take a few minutes for API requsts against that device to succeed)
/* Claim devices into a network. (Note: for recently claimed devices, it may take a few minutes for API requsts against that device to succeed)

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!claim-network-devices
*/

func (s *NetworksService) ClaimNetworkDevices(networkID string, requestNetworksClaimNetworkDevices *RequestNetworksClaimNetworkDevices) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/devices/claim"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworksClaimNetworkDevices).
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

func (s *NetworksService) VmxNetworkDevicesClaim(networkID string, requestNetworksVmxNetworkDevicesClaim *RequestNetworksVmxNetworkDevicesClaim) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/devices/claim/vmx"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworksVmxNetworkDevicesClaim).
		// SetResult(&ResponseNetworksVmxNetworkDevicesClaim{}).
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

func (s *NetworksService) RemoveNetworkDevices(networkID string, requestNetworksRemoveNetworkDevices *RequestNetworksRemoveNetworkDevices) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/devices/remove"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworksRemoveNetworkDevices).
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

//CreateNetworkFirmwareUpgradesRollback Rollback a Firmware Upgrade For A Network
/* Rollback a Firmware Upgrade For A Network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-network-firmware-upgrades-rollback
*/

func (s *NetworksService) CreateNetworkFirmwareUpgradesRollback(networkID string, requestNetworksCreateNetworkFirmwareUpgradesRollback *RequestNetworksCreateNetworkFirmwareUpgradesRollback) (*ResponseNetworksCreateNetworkFirmwareUpgradesRollback, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/firmwareUpgrades/rollbacks"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworksCreateNetworkFirmwareUpgradesRollback).
		SetResult(&ResponseNetworksCreateNetworkFirmwareUpgradesRollback{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateNetworkFirmwareUpgradesRollback")
	}

	result := response.Result().(*ResponseNetworksCreateNetworkFirmwareUpgradesRollback)
	return result, response, err

}

//CreateNetworkFirmwareUpgradesStagedEvent Create a Staged Upgrade Event for a network
/* Create a Staged Upgrade Event for a network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-network-firmware-upgrades-staged-event
*/

func (s *NetworksService) CreateNetworkFirmwareUpgradesStagedEvent(networkID string, requestNetworksCreateNetworkFirmwareUpgradesStagedEvent *RequestNetworksCreateNetworkFirmwareUpgradesStagedEvent) (*ResponseNetworksCreateNetworkFirmwareUpgradesStagedEvent, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/firmwareUpgrades/staged/events"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworksCreateNetworkFirmwareUpgradesStagedEvent).
		SetResult(&ResponseNetworksCreateNetworkFirmwareUpgradesStagedEvent{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateNetworkFirmwareUpgradesStagedEvent")
	}

	result := response.Result().(*ResponseNetworksCreateNetworkFirmwareUpgradesStagedEvent)
	return result, response, err

}

//DeferNetworkFirmwareUpgradesStagedEvents Postpone by 1 week all pending staged upgrade stages for a network
/* Postpone by 1 week all pending staged upgrade stages for a network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!defer-network-firmware-upgrades-staged-events
*/

func (s *NetworksService) DeferNetworkFirmwareUpgradesStagedEvents(networkID string) (*ResponseNetworksDeferNetworkFirmwareUpgradesStagedEvents, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/firmwareUpgrades/staged/events/defer"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworksDeferNetworkFirmwareUpgradesStagedEvents{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeferNetworkFirmwareUpgradesStagedEvents")
	}

	result := response.Result().(*ResponseNetworksDeferNetworkFirmwareUpgradesStagedEvents)
	return result, response, err

}

//RollbacksNetworkFirmwareUpgradesStagedEvents Rollback a Staged Upgrade Event for a network
/* Rollback a Staged Upgrade Event for a network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!rollbacks-network-firmware-upgrades-staged-events
*/

func (s *NetworksService) RollbacksNetworkFirmwareUpgradesStagedEvents(networkID string, requestNetworksRollbacksNetworkFirmwareUpgradesStagedEvents *RequestNetworksRollbacksNetworkFirmwareUpgradesStagedEvents) (*ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEvents, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/firmwareUpgrades/staged/events/rollbacks"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworksRollbacksNetworkFirmwareUpgradesStagedEvents).
		SetResult(&ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEvents{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation RollbacksNetworkFirmwareUpgradesStagedEvents")
	}

	result := response.Result().(*ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEvents)
	return result, response, err

}

//CreateNetworkFirmwareUpgradesStagedGroup Create a Staged Upgrade Group for a network
/* Create a Staged Upgrade Group for a network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-network-firmware-upgrades-staged-group
*/

func (s *NetworksService) CreateNetworkFirmwareUpgradesStagedGroup(networkID string, requestNetworksCreateNetworkFirmwareUpgradesStagedGroup *RequestNetworksCreateNetworkFirmwareUpgradesStagedGroup) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/firmwareUpgrades/staged/groups"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworksCreateNetworkFirmwareUpgradesStagedGroup).
		// SetResult(&ResponseNetworksCreateNetworkFirmwareUpgradesStagedGroup{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateNetworkFirmwareUpgradesStagedGroup")
	}

	return response, err

}

//CreateNetworkFloorPlan Upload a floor plan
/* Upload a floor plan

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-network-floor-plan
*/

func (s *NetworksService) CreateNetworkFloorPlan(networkID string, requestNetworksCreateNetworkFloorPlan *RequestNetworksCreateNetworkFloorPlan) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/floorPlans"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworksCreateNetworkFloorPlan).
		// SetResult(&ResponseNetworksCreateNetworkFloorPlan{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateNetworkFloorPlan")
	}

	return response, err

}

//CreateNetworkGroupPolicy Create a group policy
/* Create a group policy

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-network-group-policy
*/

func (s *NetworksService) CreateNetworkGroupPolicy(networkID string, requestNetworksCreateNetworkGroupPolicy *RequestNetworksCreateNetworkGroupPolicy) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/groupPolicies"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworksCreateNetworkGroupPolicy).
		// SetResult(&ResponseNetworksCreateNetworkGroupPolicy{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateNetworkGroupPolicy")
	}

	return response, err

}

//CreateNetworkMerakiAuthUser Authorize a user configured with Meraki Authentication for a network (currently supports 802.1X, splash guest, and client VPN users, and currently, organizations have a 50,000 user cap)
/* Authorize a user configured with Meraki Authentication for a network (currently supports 802.1X, splash guest, and client VPN users, and currently, organizations have a 50,000 user cap)

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-network-meraki-auth-user
*/

func (s *NetworksService) CreateNetworkMerakiAuthUser(networkID string, requestNetworksCreateNetworkMerakiAuthUser *RequestNetworksCreateNetworkMerakiAuthUser) (*ResponseNetworksCreateNetworkMerakiAuthUser, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/merakiAuthUsers"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworksCreateNetworkMerakiAuthUser).
		SetResult(&ResponseNetworksCreateNetworkMerakiAuthUser{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateNetworkMerakiAuthUser")
	}

	result := response.Result().(*ResponseNetworksCreateNetworkMerakiAuthUser)
	return result, response, err

}

//CreateNetworkMqttBroker Add an MQTT broker
/* Add an MQTT broker

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-network-mqtt-broker
*/

func (s *NetworksService) CreateNetworkMqttBroker(networkID string, requestNetworksCreateNetworkMqttBroker *RequestNetworksCreateNetworkMqttBroker) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/mqttBrokers"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworksCreateNetworkMqttBroker).
		// SetResult(&ResponseNetworksCreateNetworkMqttBroker{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateNetworkMqttBroker")
	}

	return response, err

}

//CreateNetworkPiiRequest Submit a new delete or restrict processing PII request
/* Submit a new delete or restrict processing PII request

## ALTERNATE PATH

***
/organizations/{organizationId}/pii/requests
***

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-network-pii-request
*/

func (s *NetworksService) CreateNetworkPiiRequest(networkID string, requestNetworksCreateNetworkPiiRequest *RequestNetworksCreateNetworkPiiRequest) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/pii/requests"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworksCreateNetworkPiiRequest).
		// SetResult(&ResponseNetworksCreateNetworkPiiRequest{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateNetworkPiiRequest")
	}

	return response, err

}

//SplitNetwork Split a combined network into individual networks for each type of device
/* Split a combined network into individual networks for each type of device

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!split-network
*/

func (s *NetworksService) SplitNetwork(networkID string) (*ResponseNetworksSplitNetwork, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/split"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworksSplitNetwork{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation SplitNetwork")
	}

	result := response.Result().(*ResponseNetworksSplitNetwork)
	return result, response, err

}

//UnbindNetwork Unbind a network from a template.
/* Unbind a network from a template.

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!unbind-network
*/

func (s *NetworksService) UnbindNetwork(networkID string, requestNetworksUnbindNetwork *RequestNetworksUnbindNetwork) (*ResponseNetworksUnbindNetwork, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/unbind"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworksUnbindNetwork).
		SetResult(&ResponseNetworksUnbindNetwork{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UnbindNetwork")
	}

	result := response.Result().(*ResponseNetworksUnbindNetwork)
	return result, response, err

}

//CreateNetworkWebhooksHTTPServer Add an HTTP server to a network
/* Add an HTTP server to a network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-network-webhooks-http-server
*/

func (s *NetworksService) CreateNetworkWebhooksHTTPServer(networkID string, requestNetworksCreateNetworkWebhooksHttpServer *RequestNetworksCreateNetworkWebhooksHTTPServer) (*ResponseNetworksCreateNetworkWebhooksHTTPServer, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/webhooks/httpServers"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworksCreateNetworkWebhooksHttpServer).
		SetResult(&ResponseNetworksCreateNetworkWebhooksHTTPServer{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateNetworkWebhooksHttpServer")
	}

	result := response.Result().(*ResponseNetworksCreateNetworkWebhooksHTTPServer)
	return result, response, err

}

//CreateNetworkWebhooksPayloadTemplate Create a webhook payload template for a network
/* Create a webhook payload template for a network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-network-webhooks-payload-template
*/

func (s *NetworksService) CreateNetworkWebhooksPayloadTemplate(networkID string, requestNetworksCreateNetworkWebhooksPayloadTemplate *RequestNetworksCreateNetworkWebhooksPayloadTemplate) (*ResponseNetworksCreateNetworkWebhooksPayloadTemplate, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/webhooks/payloadTemplates"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworksCreateNetworkWebhooksPayloadTemplate).
		SetResult(&ResponseNetworksCreateNetworkWebhooksPayloadTemplate{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateNetworkWebhooksPayloadTemplate")
	}

	result := response.Result().(*ResponseNetworksCreateNetworkWebhooksPayloadTemplate)
	return result, response, err

}

//CreateNetworkWebhooksWebhookTest Send a test webhook for a network
/* Send a test webhook for a network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-network-webhooks-webhook-test
*/

func (s *NetworksService) CreateNetworkWebhooksWebhookTest(networkID string, requestNetworksCreateNetworkWebhooksWebhookTest *RequestNetworksCreateNetworkWebhooksWebhookTest) (*ResponseNetworksCreateNetworkWebhooksWebhookTest, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/webhooks/webhookTests"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworksCreateNetworkWebhooksWebhookTest).
		SetResult(&ResponseNetworksCreateNetworkWebhooksWebhookTest{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateNetworkWebhooksWebhookTest")
	}

	result := response.Result().(*ResponseNetworksCreateNetworkWebhooksWebhookTest)
	return result, response, err

}

//CombineOrganizationNetworks Combine multiple networks into a single network
/* Combine multiple networks into a single network

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!combine-organization-networks
*/
func (s *NetworksService) CombineOrganizationNetworks(organizationID string, requestNetworksCombineOrganizationNetworks *RequestNetworksCombineOrganizationNetworks) (*ResponseNetworksCombineOrganizationNetworks, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/networks/combine"
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworksCombineOrganizationNetworks).
		SetResult(&ResponseNetworksCombineOrganizationNetworks{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CombineOrganizationNetworks")
	}

	result := response.Result().(*ResponseNetworksCombineOrganizationNetworks)
	return result, response, err

}

//UpdateNetwork Update a network
/* Update a network

@param networkID networkId path parameter. Network ID
*/
func (s *NetworksService) UpdateNetwork(networkID string, requestNetworksUpdateNetwork *RequestNetworksUpdateNetwork) (*ResponseNetworksUpdateNetwork, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworksUpdateNetwork).
		SetResult(&ResponseNetworksUpdateNetwork{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetwork")
	}

	result := response.Result().(*ResponseNetworksUpdateNetwork)
	return result, response, err

}

//UpdateNetworkAlertsSettings Update the alert configuration for this network
/* Update the alert configuration for this network

@param networkID networkId path parameter. Network ID
*/
func (s *NetworksService) UpdateNetworkAlertsSettings(networkID string, requestNetworksUpdateNetworkAlertsSettings *RequestNetworksUpdateNetworkAlertsSettings) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/alerts/settings"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworksUpdateNetworkAlertsSettings).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkAlertsSettings")
	}

	return response, err

}

//UpdateNetworkClientPolicy Update the policy assigned to a client on the network
/* Update the policy assigned to a client on the network. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.

@param networkID networkId path parameter. Network ID
@param clientID clientId path parameter. Client ID
*/
func (s *NetworksService) UpdateNetworkClientPolicy(networkID string, clientID string, requestNetworksUpdateNetworkClientPolicy *RequestNetworksUpdateNetworkClientPolicy) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/clients/{clientId}/policy"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{clientId}", fmt.Sprintf("%v", clientID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworksUpdateNetworkClientPolicy).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkClientPolicy")
	}

	return response, err

}

//UpdateNetworkClientSplashAuthorizationStatus Update a client's splash authorization
/* Update a client's splash authorization. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.

@param networkID networkId path parameter. Network ID
@param clientID clientId path parameter. Client ID
*/
func (s *NetworksService) UpdateNetworkClientSplashAuthorizationStatus(networkID string, clientID string, requestNetworksUpdateNetworkClientSplashAuthorizationStatus *RequestNetworksUpdateNetworkClientSplashAuthorizationStatus) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/clients/{clientId}/splashAuthorizationStatus"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{clientId}", fmt.Sprintf("%v", clientID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworksUpdateNetworkClientSplashAuthorizationStatus).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkClientSplashAuthorizationStatus")
	}

	return response, err

}

//UpdateNetworkFirmwareUpgrades Update firmware upgrade information for a network
/* Update firmware upgrade information for a network

@param networkID networkId path parameter. Network ID
*/
func (s *NetworksService) UpdateNetworkFirmwareUpgrades(networkID string, requestNetworksUpdateNetworkFirmwareUpgrades *RequestNetworksUpdateNetworkFirmwareUpgrades) (*ResponseNetworksUpdateNetworkFirmwareUpgrades, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/firmwareUpgrades"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworksUpdateNetworkFirmwareUpgrades).
		SetResult(&ResponseNetworksUpdateNetworkFirmwareUpgrades{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkFirmwareUpgrades")
	}

	result := response.Result().(*ResponseNetworksUpdateNetworkFirmwareUpgrades)
	return result, response, err

}

//UpdateNetworkFirmwareUpgradesStagedEvents Update the Staged Upgrade Event for a network
/* Update the Staged Upgrade Event for a network

@param networkID networkId path parameter. Network ID
*/
func (s *NetworksService) UpdateNetworkFirmwareUpgradesStagedEvents(networkID string, requestNetworksUpdateNetworkFirmwareUpgradesStagedEvents *RequestNetworksUpdateNetworkFirmwareUpgradesStagedEvents) (*ResponseNetworksUpdateNetworkFirmwareUpgradesStagedEvents, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/firmwareUpgrades/staged/events"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworksUpdateNetworkFirmwareUpgradesStagedEvents).
		SetResult(&ResponseNetworksUpdateNetworkFirmwareUpgradesStagedEvents{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkFirmwareUpgradesStagedEvents")
	}

	result := response.Result().(*ResponseNetworksUpdateNetworkFirmwareUpgradesStagedEvents)
	return result, response, err

}

//UpdateNetworkFirmwareUpgradesStagedGroup Update a Staged Upgrade Group for a network
/* Update a Staged Upgrade Group for a network

@param networkID networkId path parameter. Network ID
@param groupID groupId path parameter. Group ID
*/
func (s *NetworksService) UpdateNetworkFirmwareUpgradesStagedGroup(networkID string, groupID string, requestNetworksUpdateNetworkFirmwareUpgradesStagedGroup *RequestNetworksUpdateNetworkFirmwareUpgradesStagedGroup) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/firmwareUpgrades/staged/groups/{groupId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{groupId}", fmt.Sprintf("%v", groupID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworksUpdateNetworkFirmwareUpgradesStagedGroup).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkFirmwareUpgradesStagedGroup")
	}

	return response, err

}

//UpdateNetworkFirmwareUpgradesStagedStages Assign Staged Upgrade Group order in the sequence.
/* Assign Staged Upgrade Group order in the sequence.

@param networkID networkId path parameter. Network ID
*/
func (s *NetworksService) UpdateNetworkFirmwareUpgradesStagedStages(networkID string, requestNetworksUpdateNetworkFirmwareUpgradesStagedStages *RequestNetworksUpdateNetworkFirmwareUpgradesStagedStages) (*ResponseNetworksUpdateNetworkFirmwareUpgradesStagedStages, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/firmwareUpgrades/staged/stages"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworksUpdateNetworkFirmwareUpgradesStagedStages).
		SetResult(&ResponseNetworksUpdateNetworkFirmwareUpgradesStagedStages{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkFirmwareUpgradesStagedStages")
	}

	result := response.Result().(*ResponseNetworksUpdateNetworkFirmwareUpgradesStagedStages)
	return result, response, err

}

//UpdateNetworkFloorPlan Update a floor plan's geolocation and other meta data
/* Update a floor plan's geolocation and other meta data

@param networkID networkId path parameter. Network ID
@param floorPlanID floorPlanId path parameter. Floor plan ID
*/
func (s *NetworksService) UpdateNetworkFloorPlan(networkID string, floorPlanID string, requestNetworksUpdateNetworkFloorPlan *RequestNetworksUpdateNetworkFloorPlan) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/floorPlans/{floorPlanId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{floorPlanId}", fmt.Sprintf("%v", floorPlanID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworksUpdateNetworkFloorPlan).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkFloorPlan")
	}

	return response, err

}

//UpdateNetworkGroupPolicy Update a group policy
/* Update a group policy

@param networkID networkId path parameter. Network ID
@param groupPolicyID groupPolicyId path parameter. Group policy ID
*/
func (s *NetworksService) UpdateNetworkGroupPolicy(networkID string, groupPolicyID string, requestNetworksUpdateNetworkGroupPolicy *RequestNetworksUpdateNetworkGroupPolicy) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/groupPolicies/{groupPolicyId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{groupPolicyId}", fmt.Sprintf("%v", groupPolicyID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworksUpdateNetworkGroupPolicy).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkGroupPolicy")
	}

	return response, err

}

//UpdateNetworkMerakiAuthUser Update a user configured with Meraki Authentication (currently, 802.1X RADIUS, splash guest, and client VPN users can be updated)
/* Update a user configured with Meraki Authentication (currently, 802.1X RADIUS, splash guest, and client VPN users can be updated)

@param networkID networkId path parameter. Network ID
@param merakiAuthUserID merakiAuthUserId path parameter. Meraki auth user ID
*/
func (s *NetworksService) UpdateNetworkMerakiAuthUser(networkID string, merakiAuthUserID string, requestNetworksUpdateNetworkMerakiAuthUser *RequestNetworksUpdateNetworkMerakiAuthUser) (*ResponseNetworksUpdateNetworkMerakiAuthUser, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/merakiAuthUsers/{merakiAuthUserId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{merakiAuthUserId}", fmt.Sprintf("%v", merakiAuthUserID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworksUpdateNetworkMerakiAuthUser).
		SetResult(&ResponseNetworksUpdateNetworkMerakiAuthUser{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkMerakiAuthUser")
	}

	result := response.Result().(*ResponseNetworksUpdateNetworkMerakiAuthUser)
	return result, response, err

}

//UpdateNetworkMqttBroker Update an MQTT broker
/* Update an MQTT broker

@param networkID networkId path parameter. Network ID
@param mqttBrokerID mqttBrokerId path parameter. Mqtt broker ID
*/
func (s *NetworksService) UpdateNetworkMqttBroker(networkID string, mqttBrokerID string, requestNetworksUpdateNetworkMqttBroker *RequestNetworksUpdateNetworkMqttBroker) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/mqttBrokers/{mqttBrokerId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{mqttBrokerId}", fmt.Sprintf("%v", mqttBrokerID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworksUpdateNetworkMqttBroker).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkMqttBroker")
	}

	return response, err

}

//UpdateNetworkNetflow Update the NetFlow traffic reporting settings for a network
/* Update the NetFlow traffic reporting settings for a network

@param networkID networkId path parameter. Network ID
*/
func (s *NetworksService) UpdateNetworkNetflow(networkID string, requestNetworksUpdateNetworkNetflow *RequestNetworksUpdateNetworkNetflow) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/netflow"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworksUpdateNetworkNetflow).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkNetflow")
	}

	return response, err

}

//UpdateNetworkSettings Update the settings for a network
/* Update the settings for a network

@param networkID networkId path parameter. Network ID
*/
func (s *NetworksService) UpdateNetworkSettings(networkID string, requestNetworksUpdateNetworkSettings *RequestNetworksUpdateNetworkSettings) (*ResponseNetworksUpdateNetworkSettings, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/settings"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworksUpdateNetworkSettings).
		SetResult(&ResponseNetworksUpdateNetworkSettings{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkSettings")
	}

	result := response.Result().(*ResponseNetworksUpdateNetworkSettings)
	return result, response, err

}

//UpdateNetworkSNMP Update the SNMP settings for a network
/* Update the SNMP settings for a network

@param networkID networkId path parameter. Network ID
*/
func (s *NetworksService) UpdateNetworkSNMP(networkID string, requestNetworksUpdateNetworkSnmp *RequestNetworksUpdateNetworkSNMP) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/snmp"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworksUpdateNetworkSnmp).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkSnmp")
	}

	return response, err

}

//UpdateNetworkSyslogServers Update the syslog servers for a network
/* Update the syslog servers for a network

@param networkID networkId path parameter. Network ID
*/
func (s *NetworksService) UpdateNetworkSyslogServers(networkID string, requestNetworksUpdateNetworkSyslogServers *RequestNetworksUpdateNetworkSyslogServers) (*ResponseNetworksUpdateNetworkSyslogServers, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/syslogServers"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworksUpdateNetworkSyslogServers).
		SetResult(&ResponseNetworksUpdateNetworkSyslogServers{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkSyslogServers")
	}

	result := response.Result().(*ResponseNetworksUpdateNetworkSyslogServers)
	return result, response, err

}

//UpdateNetworkTrafficAnalysis Update the traffic analysis settings for a network
/* Update the traffic analysis settings for a network

@param networkID networkId path parameter. Network ID
*/
func (s *NetworksService) UpdateNetworkTrafficAnalysis(networkID string, requestNetworksUpdateNetworkTrafficAnalysis *RequestNetworksUpdateNetworkTrafficAnalysis) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/trafficAnalysis"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworksUpdateNetworkTrafficAnalysis).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkTrafficAnalysis")
	}

	return response, err

}

//UpdateNetworkWebhooksHTTPServer Update an HTTP server
/* Update an HTTP server. To change a URL, create a new HTTP server.

@param networkID networkId path parameter. Network ID
@param httpServerID httpServerId path parameter. Http server ID
*/
func (s *NetworksService) UpdateNetworkWebhooksHTTPServer(networkID string, httpServerID string, requestNetworksUpdateNetworkWebhooksHttpServer *RequestNetworksUpdateNetworkWebhooksHTTPServer) (*ResponseNetworksUpdateNetworkWebhooksHTTPServer, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/webhooks/httpServers/{httpServerId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{httpServerId}", fmt.Sprintf("%v", httpServerID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworksUpdateNetworkWebhooksHttpServer).
		SetResult(&ResponseNetworksUpdateNetworkWebhooksHTTPServer{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkWebhooksHttpServer")
	}

	result := response.Result().(*ResponseNetworksUpdateNetworkWebhooksHTTPServer)
	return result, response, err

}

//UpdateNetworkWebhooksPayloadTemplate Update a webhook payload template for a network
/* Update a webhook payload template for a network

@param networkID networkId path parameter. Network ID
@param payloadTemplateID payloadTemplateId path parameter. Payload template ID
*/
func (s *NetworksService) UpdateNetworkWebhooksPayloadTemplate(networkID string, payloadTemplateID string, requestNetworksUpdateNetworkWebhooksPayloadTemplate *RequestNetworksUpdateNetworkWebhooksPayloadTemplate) (*ResponseNetworksUpdateNetworkWebhooksPayloadTemplate, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/webhooks/payloadTemplates/{payloadTemplateId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{payloadTemplateId}", fmt.Sprintf("%v", payloadTemplateID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworksUpdateNetworkWebhooksPayloadTemplate).
		SetResult(&ResponseNetworksUpdateNetworkWebhooksPayloadTemplate{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkWebhooksPayloadTemplate")
	}

	result := response.Result().(*ResponseNetworksUpdateNetworkWebhooksPayloadTemplate)
	return result, response, err

}

//DeleteNetwork Delete a network
/* Delete a network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-network
*/
func (s *NetworksService) DeleteNetwork(networkID string) (*resty.Response, error) {
	//networkID string
	path := "/api/v1/networks/{networkId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteNetwork")
	}

	return response, err

}

//DeleteNetworkFirmwareUpgradesStagedGroup Delete a Staged Upgrade Group
/* Delete a Staged Upgrade Group

@param networkID networkId path parameter. Network ID
@param groupID groupId path parameter. Group ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-network-firmware-upgrades-staged-group
*/
func (s *NetworksService) DeleteNetworkFirmwareUpgradesStagedGroup(networkID string, groupID string) (*resty.Response, error) {
	//networkID string,groupID string
	path := "/api/v1/networks/{networkId}/firmwareUpgrades/staged/groups/{groupId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{groupId}", fmt.Sprintf("%v", groupID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteNetworkFirmwareUpgradesStagedGroup")
	}

	return response, err

}

//DeleteNetworkFloorPlan Destroy a floor plan
/* Destroy a floor plan

@param networkID networkId path parameter. Network ID
@param floorPlanID floorPlanId path parameter. Floor plan ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-network-floor-plan
*/
func (s *NetworksService) DeleteNetworkFloorPlan(networkID string, floorPlanID string) (*resty.Response, error) {
	//networkID string,floorPlanID string
	path := "/api/v1/networks/{networkId}/floorPlans/{floorPlanId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{floorPlanId}", fmt.Sprintf("%v", floorPlanID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteNetworkFloorPlan")
	}

	return response, err

}

//DeleteNetworkGroupPolicy Delete a group policy
/* Delete a group policy

@param networkID networkId path parameter. Network ID
@param groupPolicyID groupPolicyId path parameter. Group policy ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-network-group-policy
*/
func (s *NetworksService) DeleteNetworkGroupPolicy(networkID string, groupPolicyID string) (*resty.Response, error) {
	//networkID string,groupPolicyID string
	path := "/api/v1/networks/{networkId}/groupPolicies/{groupPolicyId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{groupPolicyId}", fmt.Sprintf("%v", groupPolicyID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteNetworkGroupPolicy")
	}

	return response, err

}

//DeleteNetworkMerakiAuthUser Deauthorize a user
/* Deauthorize a user. To reauthorize a user after deauthorizing them, POST to this endpoint. (Currently, 802.1X RADIUS, splash guest, and client VPN users can be deauthorized.)

@param networkID networkId path parameter. Network ID
@param merakiAuthUserID merakiAuthUserId path parameter. Meraki auth user ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-network-meraki-auth-user
*/
func (s *NetworksService) DeleteNetworkMerakiAuthUser(networkID string, merakiAuthUserID string) (*resty.Response, error) {
	//networkID string,merakiAuthUserID string
	path := "/api/v1/networks/{networkId}/merakiAuthUsers/{merakiAuthUserId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{merakiAuthUserId}", fmt.Sprintf("%v", merakiAuthUserID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteNetworkMerakiAuthUser")
	}

	return response, err

}

//DeleteNetworkMqttBroker Delete an MQTT broker
/* Delete an MQTT broker

@param networkID networkId path parameter. Network ID
@param mqttBrokerID mqttBrokerId path parameter. Mqtt broker ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-network-mqtt-broker
*/
func (s *NetworksService) DeleteNetworkMqttBroker(networkID string, mqttBrokerID string) (*resty.Response, error) {
	//networkID string,mqttBrokerID string
	path := "/api/v1/networks/{networkId}/mqttBrokers/{mqttBrokerId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{mqttBrokerId}", fmt.Sprintf("%v", mqttBrokerID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteNetworkMqttBroker")
	}

	return response, err

}

//DeleteNetworkPiiRequest Delete a restrict processing PII request
/* Delete a restrict processing PII request

## ALTERNATE PATH

***
/organizations/{organizationId}/pii/requests/{requestId}
***

@param networkID networkId path parameter. Network ID
@param requestID requestId path parameter. Request ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-network-pii-request
*/
func (s *NetworksService) DeleteNetworkPiiRequest(networkID string, requestID string) (*resty.Response, error) {
	//networkID string,requestID string
	path := "/api/v1/networks/{networkId}/pii/requests/{requestId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{requestId}", fmt.Sprintf("%v", requestID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteNetworkPiiRequest")
	}

	return response, err

}

//DeleteNetworkWebhooksHTTPServer Delete an HTTP server from a network
/* Delete an HTTP server from a network

@param networkID networkId path parameter. Network ID
@param httpServerID httpServerId path parameter. Http server ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-network-webhooks-http-server
*/
func (s *NetworksService) DeleteNetworkWebhooksHTTPServer(networkID string, httpServerID string) (*resty.Response, error) {
	//networkID string,httpServerID string
	path := "/api/v1/networks/{networkId}/webhooks/httpServers/{httpServerId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{httpServerId}", fmt.Sprintf("%v", httpServerID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteNetworkWebhooksHttpServer")
	}

	return response, err

}

//DeleteNetworkWebhooksPayloadTemplate Destroy a webhook payload template for a network
/* Destroy a webhook payload template for a network. Does not work for included templates ('wpt_00001', 'wpt_00002', 'wpt_00003', 'wpt_00004', 'wpt_00005' or 'wpt_00006')

@param networkID networkId path parameter. Network ID
@param payloadTemplateID payloadTemplateId path parameter. Payload template ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-network-webhooks-payload-template
*/
func (s *NetworksService) DeleteNetworkWebhooksPayloadTemplate(networkID string, payloadTemplateID string) (*resty.Response, error) {
	//networkID string,payloadTemplateID string
	path := "/api/v1/networks/{networkId}/webhooks/payloadTemplates/{payloadTemplateId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{payloadTemplateId}", fmt.Sprintf("%v", payloadTemplateID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteNetworkWebhooksPayloadTemplate")
	}

	return response, err

}
