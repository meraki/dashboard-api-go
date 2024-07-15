package meraki

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type WirelessService service

type GetDeviceWirelessConnectionStatsQueryParams struct {
	T0       string  `url:"t0,omitempty"`       //The beginning of the timespan for the data. The maximum lookback period is 180 days from today.
	T1       string  `url:"t1,omitempty"`       //The end of the timespan for the data. t1 can be a maximum of 7 days after t0.
	Timespan float64 `url:"timespan,omitempty"` //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days.
	Band     string  `url:"band,omitempty"`     //Filter results by band (either '2.4', '5' or '6'). Note that data prior to February 2020 will not have band information.
	SSID     int     `url:"ssid,omitempty"`     //Filter results by SSID
	VLAN     int     `url:"vlan,omitempty"`     //Filter results by VLAN
	ApTag    string  `url:"apTag,omitempty"`    //Filter results by AP Tag
}
type GetDeviceWirelessLatencyStatsQueryParams struct {
	T0       string  `url:"t0,omitempty"`       //The beginning of the timespan for the data. The maximum lookback period is 180 days from today.
	T1       string  `url:"t1,omitempty"`       //The end of the timespan for the data. t1 can be a maximum of 7 days after t0.
	Timespan float64 `url:"timespan,omitempty"` //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days.
	Band     string  `url:"band,omitempty"`     //Filter results by band (either '2.4', '5' or '6'). Note that data prior to February 2020 will not have band information.
	SSID     int     `url:"ssid,omitempty"`     //Filter results by SSID
	VLAN     int     `url:"vlan,omitempty"`     //Filter results by VLAN
	ApTag    string  `url:"apTag,omitempty"`    //Filter results by AP Tag
	Fields   string  `url:"fields,omitempty"`   //Partial selection: If present, this call will return only the selected fields of ["rawDistribution", "avg"]. All fields will be returned by default. Selected fields must be entered as a comma separated string.
}
type GetNetworkWirelessAirMarshalQueryParams struct {
	T0       string  `url:"t0,omitempty"`       //The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	Timespan float64 `url:"timespan,omitempty"` //The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.
}
type GetNetworkWirelessChannelUtilizationHistoryQueryParams struct {
	T0             string  `url:"t0,omitempty"`             //The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	T1             string  `url:"t1,omitempty"`             //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan       float64 `url:"timespan,omitempty"`       //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.
	Resolution     int     `url:"resolution,omitempty"`     //The time resolution in seconds for returned data. The valid resolutions are: 600, 1200, 3600, 14400, 86400. The default is 86400.
	AutoResolution bool    `url:"autoResolution,omitempty"` //Automatically select a data resolution based on the given timespan; this overrides the value specified by the 'resolution' parameter. The default setting is false.
	ClientID       string  `url:"clientId,omitempty"`       //Filter results by network client to return per-device, per-band AP channel utilization metrics inner joined by the queried client's connection history.
	DeviceSerial   string  `url:"deviceSerial,omitempty"`   //Filter results by device to return AP channel utilization metrics for the queried device; either :band or :clientId must be jointly specified.
	ApTag          string  `url:"apTag,omitempty"`          //Filter results by AP tag to return AP channel utilization metrics for devices labeled with the given tag; either :clientId or :deviceSerial must be jointly specified.
	Band           string  `url:"band,omitempty"`           //Filter results by band (either '2.4', '5' or '6').
}
type GetNetworkWirelessClientCountHistoryQueryParams struct {
	T0             string  `url:"t0,omitempty"`             //The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	T1             string  `url:"t1,omitempty"`             //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan       float64 `url:"timespan,omitempty"`       //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.
	Resolution     int     `url:"resolution,omitempty"`     //The time resolution in seconds for returned data. The valid resolutions are: 300, 600, 1200, 3600, 14400, 86400. The default is 86400.
	AutoResolution bool    `url:"autoResolution,omitempty"` //Automatically select a data resolution based on the given timespan; this overrides the value specified by the 'resolution' parameter. The default setting is false.
	ClientID       string  `url:"clientId,omitempty"`       //Filter results by network client to return per-device client counts over time inner joined by the queried client's connection history.
	DeviceSerial   string  `url:"deviceSerial,omitempty"`   //Filter results by device.
	ApTag          string  `url:"apTag,omitempty"`          //Filter results by AP tag.
	Band           string  `url:"band,omitempty"`           //Filter results by band (either '2.4', '5' or '6').
	SSID           int     `url:"ssid,omitempty"`           //Filter results by SSID number.
}
type GetNetworkWirelessClientsConnectionStatsQueryParams struct {
	T0       string  `url:"t0,omitempty"`       //The beginning of the timespan for the data. The maximum lookback period is 180 days from today.
	T1       string  `url:"t1,omitempty"`       //The end of the timespan for the data. t1 can be a maximum of 7 days after t0.
	Timespan float64 `url:"timespan,omitempty"` //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days.
	Band     string  `url:"band,omitempty"`     //Filter results by band (either '2.4', '5' or '6'). Note that data prior to February 2020 will not have band information.
	SSID     int     `url:"ssid,omitempty"`     //Filter results by SSID
	VLAN     int     `url:"vlan,omitempty"`     //Filter results by VLAN
	ApTag    string  `url:"apTag,omitempty"`    //Filter results by AP Tag
}
type GetNetworkWirelessClientsLatencyStatsQueryParams struct {
	T0       string  `url:"t0,omitempty"`       //The beginning of the timespan for the data. The maximum lookback period is 180 days from today.
	T1       string  `url:"t1,omitempty"`       //The end of the timespan for the data. t1 can be a maximum of 7 days after t0.
	Timespan float64 `url:"timespan,omitempty"` //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days.
	Band     string  `url:"band,omitempty"`     //Filter results by band (either '2.4', '5' or '6'). Note that data prior to February 2020 will not have band information.
	SSID     int     `url:"ssid,omitempty"`     //Filter results by SSID
	VLAN     int     `url:"vlan,omitempty"`     //Filter results by VLAN
	ApTag    string  `url:"apTag,omitempty"`    //Filter results by AP Tag
	Fields   string  `url:"fields,omitempty"`   //Partial selection: If present, this call will return only the selected fields of ["rawDistribution", "avg"]. All fields will be returned by default. Selected fields must be entered as a comma separated string.
}
type GetNetworkWirelessClientConnectionStatsQueryParams struct {
	T0       string  `url:"t0,omitempty"`       //The beginning of the timespan for the data. The maximum lookback period is 180 days from today.
	T1       string  `url:"t1,omitempty"`       //The end of the timespan for the data. t1 can be a maximum of 7 days after t0.
	Timespan float64 `url:"timespan,omitempty"` //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days.
	Band     string  `url:"band,omitempty"`     //Filter results by band (either '2.4', '5' or '6'). Note that data prior to February 2020 will not have band information.
	SSID     int     `url:"ssid,omitempty"`     //Filter results by SSID
	VLAN     int     `url:"vlan,omitempty"`     //Filter results by VLAN
	ApTag    string  `url:"apTag,omitempty"`    //Filter results by AP Tag
}
type GetNetworkWirelessClientConnectivityEventsQueryParams struct {
	PerPage            int      `url:"perPage,omitempty"`              //The number of entries per page returned. Acceptable range is 3 - 1000.
	StartingAfter      string   `url:"startingAfter,omitempty"`        //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore       string   `url:"endingBefore,omitempty"`         //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	T0                 string   `url:"t0,omitempty"`                   //The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	T1                 string   `url:"t1,omitempty"`                   //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan           float64  `url:"timespan,omitempty"`             //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
	Types              []string `url:"types[],omitempty"`              //A list of event types to include. If not specified, events of all types will be returned. Valid types are 'assoc', 'disassoc', 'auth', 'deauth', 'dns', 'dhcp', 'roam', 'connection' and/or 'sticky'.
	IncludedSeverities []string `url:"includedSeverities[],omitempty"` //A list of severities to include. If not specified, events of all severities will be returned. Valid severities are 'good', 'info', 'warn' and/or 'bad'.
	Band               string   `url:"band,omitempty"`                 //Filter results by band (either '2.4', '5', '6').
	SSIDNumber         int      `url:"ssidNumber,omitempty"`           //An SSID number to include. If not specified, events for all SSIDs will be returned.
	DeviceSerial       string   `url:"deviceSerial,omitempty"`         //Filter results by an AP's serial number.
}
type GetNetworkWirelessClientLatencyHistoryQueryParams struct {
	T0         string  `url:"t0,omitempty"`         //The beginning of the timespan for the data. The maximum lookback period is 791 days from today.
	T1         string  `url:"t1,omitempty"`         //The end of the timespan for the data. t1 can be a maximum of 791 days after t0.
	Timespan   float64 `url:"timespan,omitempty"`   //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 791 days. The default is 1 day.
	Resolution int     `url:"resolution,omitempty"` //The time resolution in seconds for returned data. The valid resolutions are: 86400. The default is 86400.
}
type GetNetworkWirelessClientLatencyStatsQueryParams struct {
	T0       string  `url:"t0,omitempty"`       //The beginning of the timespan for the data. The maximum lookback period is 180 days from today.
	T1       string  `url:"t1,omitempty"`       //The end of the timespan for the data. t1 can be a maximum of 7 days after t0.
	Timespan float64 `url:"timespan,omitempty"` //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days.
	Band     string  `url:"band,omitempty"`     //Filter results by band (either '2.4', '5' or '6'). Note that data prior to February 2020 will not have band information.
	SSID     int     `url:"ssid,omitempty"`     //Filter results by SSID
	VLAN     int     `url:"vlan,omitempty"`     //Filter results by VLAN
	ApTag    string  `url:"apTag,omitempty"`    //Filter results by AP Tag
	Fields   string  `url:"fields,omitempty"`   //Partial selection: If present, this call will return only the selected fields of ["rawDistribution", "avg"]. All fields will be returned by default. Selected fields must be entered as a comma separated string.
}
type GetNetworkWirelessConnectionStatsQueryParams struct {
	T0       string  `url:"t0,omitempty"`       //The beginning of the timespan for the data. The maximum lookback period is 180 days from today.
	T1       string  `url:"t1,omitempty"`       //The end of the timespan for the data. t1 can be a maximum of 7 days after t0.
	Timespan float64 `url:"timespan,omitempty"` //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days.
	Band     string  `url:"band,omitempty"`     //Filter results by band (either '2.4', '5' or '6'). Note that data prior to February 2020 will not have band information.
	SSID     int     `url:"ssid,omitempty"`     //Filter results by SSID
	VLAN     int     `url:"vlan,omitempty"`     //Filter results by VLAN
	ApTag    string  `url:"apTag,omitempty"`    //Filter results by AP Tag
}
type GetNetworkWirelessDataRateHistoryQueryParams struct {
	T0             string  `url:"t0,omitempty"`             //The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	T1             string  `url:"t1,omitempty"`             //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan       float64 `url:"timespan,omitempty"`       //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.
	Resolution     int     `url:"resolution,omitempty"`     //The time resolution in seconds for returned data. The valid resolutions are: 300, 600, 1200, 3600, 14400, 86400. The default is 86400.
	AutoResolution bool    `url:"autoResolution,omitempty"` //Automatically select a data resolution based on the given timespan; this overrides the value specified by the 'resolution' parameter. The default setting is false.
	ClientID       string  `url:"clientId,omitempty"`       //Filter results by network client.
	DeviceSerial   string  `url:"deviceSerial,omitempty"`   //Filter results by device.
	ApTag          string  `url:"apTag,omitempty"`          //Filter results by AP tag.
	Band           string  `url:"band,omitempty"`           //Filter results by band (either '2.4', '5' or '6').
	SSID           int     `url:"ssid,omitempty"`           //Filter results by SSID number.
}
type GetNetworkWirelessDevicesConnectionStatsQueryParams struct {
	T0       string  `url:"t0,omitempty"`       //The beginning of the timespan for the data. The maximum lookback period is 180 days from today.
	T1       string  `url:"t1,omitempty"`       //The end of the timespan for the data. t1 can be a maximum of 7 days after t0.
	Timespan float64 `url:"timespan,omitempty"` //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days.
	Band     string  `url:"band,omitempty"`     //Filter results by band (either '2.4', '5' or '6'). Note that data prior to February 2020 will not have band information.
	SSID     int     `url:"ssid,omitempty"`     //Filter results by SSID
	VLAN     int     `url:"vlan,omitempty"`     //Filter results by VLAN
	ApTag    string  `url:"apTag,omitempty"`    //Filter results by AP Tag
}
type GetNetworkWirelessDevicesLatencyStatsQueryParams struct {
	T0       string  `url:"t0,omitempty"`       //The beginning of the timespan for the data. The maximum lookback period is 180 days from today.
	T1       string  `url:"t1,omitempty"`       //The end of the timespan for the data. t1 can be a maximum of 7 days after t0.
	Timespan float64 `url:"timespan,omitempty"` //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days.
	Band     string  `url:"band,omitempty"`     //Filter results by band (either '2.4', '5' or '6'). Note that data prior to February 2020 will not have band information.
	SSID     int     `url:"ssid,omitempty"`     //Filter results by SSID
	VLAN     int     `url:"vlan,omitempty"`     //Filter results by VLAN
	ApTag    string  `url:"apTag,omitempty"`    //Filter results by AP Tag
	Fields   string  `url:"fields,omitempty"`   //Partial selection: If present, this call will return only the selected fields of ["rawDistribution", "avg"]. All fields will be returned by default. Selected fields must be entered as a comma separated string.
}
type GetNetworkWirelessFailedConnectionsQueryParams struct {
	T0       string  `url:"t0,omitempty"`       //The beginning of the timespan for the data. The maximum lookback period is 180 days from today.
	T1       string  `url:"t1,omitempty"`       //The end of the timespan for the data. t1 can be a maximum of 7 days after t0.
	Timespan float64 `url:"timespan,omitempty"` //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days.
	Band     string  `url:"band,omitempty"`     //Filter results by band (either '2.4', '5' or '6'). Note that data prior to February 2020 will not have band information.
	SSID     int     `url:"ssid,omitempty"`     //Filter results by SSID
	VLAN     int     `url:"vlan,omitempty"`     //Filter results by VLAN
	ApTag    string  `url:"apTag,omitempty"`    //Filter results by AP Tag
	Serial   string  `url:"serial,omitempty"`   //Filter by AP
	ClientID string  `url:"clientId,omitempty"` //Filter by client MAC
}
type GetNetworkWirelessLatencyHistoryQueryParams struct {
	T0             string  `url:"t0,omitempty"`             //The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	T1             string  `url:"t1,omitempty"`             //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan       float64 `url:"timespan,omitempty"`       //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.
	Resolution     int     `url:"resolution,omitempty"`     //The time resolution in seconds for returned data. The valid resolutions are: 300, 600, 1200, 3600, 14400, 86400. The default is 86400.
	AutoResolution bool    `url:"autoResolution,omitempty"` //Automatically select a data resolution based on the given timespan; this overrides the value specified by the 'resolution' parameter. The default setting is false.
	ClientID       string  `url:"clientId,omitempty"`       //Filter results by network client.
	DeviceSerial   string  `url:"deviceSerial,omitempty"`   //Filter results by device.
	ApTag          string  `url:"apTag,omitempty"`          //Filter results by AP tag.
	Band           string  `url:"band,omitempty"`           //Filter results by band (either '2.4', '5' or '6').
	SSID           int     `url:"ssid,omitempty"`           //Filter results by SSID number.
	AccessCategory string  `url:"accessCategory,omitempty"` //Filter by access category.
}
type GetNetworkWirelessLatencyStatsQueryParams struct {
	T0       string  `url:"t0,omitempty"`       //The beginning of the timespan for the data. The maximum lookback period is 180 days from today.
	T1       string  `url:"t1,omitempty"`       //The end of the timespan for the data. t1 can be a maximum of 7 days after t0.
	Timespan float64 `url:"timespan,omitempty"` //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days.
	Band     string  `url:"band,omitempty"`     //Filter results by band (either '2.4', '5' or '6'). Note that data prior to February 2020 will not have band information.
	SSID     int     `url:"ssid,omitempty"`     //Filter results by SSID
	VLAN     int     `url:"vlan,omitempty"`     //Filter results by VLAN
	ApTag    string  `url:"apTag,omitempty"`    //Filter results by AP Tag
	Fields   string  `url:"fields,omitempty"`   //Partial selection: If present, this call will return only the selected fields of ["rawDistribution", "avg"]. All fields will be returned by default. Selected fields must be entered as a comma separated string.
}
type GetNetworkWirelessMeshStatusesQueryParams struct {
	PerPage       int    `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 500. Default is 50.
	StartingAfter string `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
}
type GetNetworkWirelessRfProfilesQueryParams struct {
	IncludeTemplateProfiles bool `url:"includeTemplateProfiles,omitempty"` //If the network is bound to a template, this parameter controls whether or not the non-basic RF profiles defined on the template should be included in the response alongside the non-basic profiles defined on the bound network. Defaults to false.
}
type GetNetworkWirelessSignalQualityHistoryQueryParams struct {
	T0             string  `url:"t0,omitempty"`             //The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	T1             string  `url:"t1,omitempty"`             //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan       float64 `url:"timespan,omitempty"`       //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.
	Resolution     int     `url:"resolution,omitempty"`     //The time resolution in seconds for returned data. The valid resolutions are: 300, 600, 1200, 3600, 14400, 86400. The default is 86400.
	AutoResolution bool    `url:"autoResolution,omitempty"` //Automatically select a data resolution based on the given timespan; this overrides the value specified by the 'resolution' parameter. The default setting is false.
	ClientID       string  `url:"clientId,omitempty"`       //Filter results by network client.
	DeviceSerial   string  `url:"deviceSerial,omitempty"`   //Filter results by device.
	ApTag          string  `url:"apTag,omitempty"`          //Filter results by AP tag; either :clientId or :deviceSerial must be jointly specified.
	Band           string  `url:"band,omitempty"`           //Filter results by band (either '2.4', '5' or '6').
	SSID           int     `url:"ssid,omitempty"`           //Filter results by SSID number.
}
type GetNetworkWirelessUsageHistoryQueryParams struct {
	T0             string  `url:"t0,omitempty"`             //The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	T1             string  `url:"t1,omitempty"`             //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan       float64 `url:"timespan,omitempty"`       //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.
	Resolution     int     `url:"resolution,omitempty"`     //The time resolution in seconds for returned data. The valid resolutions are: 300, 600, 1200, 3600, 14400, 86400. The default is 86400.
	AutoResolution bool    `url:"autoResolution,omitempty"` //Automatically select a data resolution based on the given timespan; this overrides the value specified by the 'resolution' parameter. The default setting is false.
	ClientID       string  `url:"clientId,omitempty"`       //Filter results by network client to return per-device AP usage over time inner joined by the queried client's connection history.
	DeviceSerial   string  `url:"deviceSerial,omitempty"`   //Filter results by device. Requires :band.
	ApTag          string  `url:"apTag,omitempty"`          //Filter results by AP tag; either :clientId or :deviceSerial must be jointly specified.
	Band           string  `url:"band,omitempty"`           //Filter results by band (either '2.4', '5' or '6').
	SSID           int     `url:"ssid,omitempty"`           //Filter results by SSID number.
}

type GetOrganizationWirelessDevicesEthernetStatusesQueryParams struct {
	PerPage       int      `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100.
	StartingAfter string   `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string   `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	NetworkIDs    []string `url:"networkIds[],omitempty"`  //A list of Meraki network IDs to filter results to contain only specified networks. E.g.: networkIds[]=N_12345678&networkIds[]=L_3456
}

type ResponseWirelessUpdateDeviceWirelessAlternateManagementInterfaceIPv6 struct {
	Addresses *[]ResponseWirelessUpdateDeviceWirelessAlternateManagementInterfaceIPv6Addresses `json:"addresses,omitempty"` // configured alternate management interface addresses
}
type ResponseWirelessUpdateDeviceWirelessAlternateManagementInterfaceIPv6Addresses struct {
	Address        string                                                                                    `json:"address,omitempty"`        // The IP address configured for the alternate management interface
	AssignmentMode string                                                                                    `json:"assignmentMode,omitempty"` // The type of address assignment. Either static or dynamic.
	Gateway        string                                                                                    `json:"gateway,omitempty"`        // The gateway address configured for the alternate managment interface
	Nameservers    *ResponseWirelessUpdateDeviceWirelessAlternateManagementInterfaceIPv6AddressesNameservers `json:"nameservers,omitempty"`    // The DNS servers settings for this address.
	Prefix         string                                                                                    `json:"prefix,omitempty"`         // The IPv6 prefix of the interface. Required if IPv6 object is included.
	Protocol       string                                                                                    `json:"protocol,omitempty"`       // The IP protocol used for the address
}
type ResponseWirelessUpdateDeviceWirelessAlternateManagementInterfaceIPv6AddressesNameservers struct {
	Addresses []string `json:"addresses,omitempty"` // Up to 2 nameserver addresses to use, ordered in priority from highest to lowest priority.
}
type ResponseWirelessGetDeviceWirelessBluetoothSettings struct {
	Major *int   `json:"major,omitempty"` // Desired major value of the beacon. If the value is set to null it will reset to Dashboard's automatically generated value.
	Minor *int   `json:"minor,omitempty"` // Desired minor value of the beacon. If the value is set to null it will reset to Dashboard's automatically generated value.
	UUID  string `json:"uuid,omitempty"`  // Desired UUID of the beacon. If the value is set to null it will reset to Dashboard's automatically generated value.
}
type ResponseWirelessUpdateDeviceWirelessBluetoothSettings struct {
	Major *int   `json:"major,omitempty"` // Desired major value of the beacon. If the value is set to null it will reset to Dashboard's automatically generated value.
	Minor *int   `json:"minor,omitempty"` // Desired minor value of the beacon. If the value is set to null it will reset to Dashboard's automatically generated value.
	UUID  string `json:"uuid,omitempty"`  // Desired UUID of the beacon. If the value is set to null it will reset to Dashboard's automatically generated value.
}
type ResponseWirelessGetDeviceWirelessConnectionStats struct {
	ConnectionStats *ResponseWirelessGetDeviceWirelessConnectionStatsConnectionStats `json:"connectionStats,omitempty"` // The connection stats of the device
	Serial          string                                                           `json:"serial,omitempty"`          // The serial number for the device
}
type ResponseWirelessGetDeviceWirelessConnectionStatsConnectionStats struct {
	Assoc   *int `json:"assoc,omitempty"`   // The number of failed association attempts
	Auth    *int `json:"auth,omitempty"`    // The number of failed authentication attempts
	Dhcp    *int `json:"dhcp,omitempty"`    // The number of failed DHCP attempts
	DNS     *int `json:"dns,omitempty"`     // The number of failed DNS attempts
	Success *int `json:"success,omitempty"` // The number of successful connection attempts
}
type ResponseWirelessGetDeviceWirelessLatencyStats struct {
	LatencyStats *ResponseWirelessGetDeviceWirelessLatencyStatsLatencyStats `json:"latencyStats,omitempty"` //
	Serial       string                                                     `json:"serial,omitempty"`       //
}
type ResponseWirelessGetDeviceWirelessLatencyStatsLatencyStats struct {
	BackgroundTraffic *ResponseWirelessGetDeviceWirelessLatencyStatsLatencyStatsBackgroundTraffic `json:"backgroundTraffic,omitempty"` //
	BestEffortTraffic string                                                                      `json:"bestEffortTraffic,omitempty"` //
	VideoTraffic      string                                                                      `json:"videoTraffic,omitempty"`      //
	VoiceTraffic      string                                                                      `json:"voiceTraffic,omitempty"`      //
}
type ResponseWirelessGetDeviceWirelessLatencyStatsLatencyStatsBackgroundTraffic struct {
	Avg             *float64                                                                                   `json:"avg,omitempty"`             //
	RawDistribution *ResponseWirelessGetDeviceWirelessLatencyStatsLatencyStatsBackgroundTrafficRawDistribution `json:"rawDistribution,omitempty"` //
}
type ResponseWirelessGetDeviceWirelessLatencyStatsLatencyStatsBackgroundTrafficRawDistribution struct {
	Status0    *int `json:"0,omitempty"`    //
	Status1    *int `json:"1,omitempty"`    //
	Status1024 *int `json:"1024,omitempty"` //
	Status128  *int `json:"128,omitempty"`  //
	Status16   *int `json:"16,omitempty"`   //
	Status2    *int `json:"2,omitempty"`    //
	Status2048 *int `json:"2048,omitempty"` //
	Status256  *int `json:"256,omitempty"`  //
	Status32   *int `json:"32,omitempty"`   //
	Status4    *int `json:"4,omitempty"`    //
	Status512  *int `json:"512,omitempty"`  //
	Status64   *int `json:"64,omitempty"`   //
	Status8    *int `json:"8,omitempty"`    //
}
type ResponseWirelessGetDeviceWirelessRadioSettings struct {
	FiveGhzSettings    *ResponseWirelessGetDeviceWirelessRadioSettingsFiveGhzSettings    `json:"fiveGhzSettings,omitempty"`    //
	RfProfileID        string                                                            `json:"rfProfileId,omitempty"`        //
	Serial             string                                                            `json:"serial,omitempty"`             //
	TwoFourGhzSettings *ResponseWirelessGetDeviceWirelessRadioSettingsTwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"` //
}
type ResponseWirelessGetDeviceWirelessRadioSettingsFiveGhzSettings struct {
	Channel      *int `json:"channel,omitempty"`      //
	ChannelWidth *int `json:"channelWidth,omitempty"` //
	TargetPower  *int `json:"targetPower,omitempty"`  //
}
type ResponseWirelessGetDeviceWirelessRadioSettingsTwoFourGhzSettings struct {
	Channel     *int `json:"channel,omitempty"`     //
	TargetPower *int `json:"targetPower,omitempty"` //
}
type ResponseWirelessUpdateDeviceWirelessRadioSettings interface{}
type ResponseWirelessGetDeviceWirelessStatus struct {
	BasicServiceSets *[]ResponseWirelessGetDeviceWirelessStatusBasicServiceSets `json:"basicServiceSets,omitempty"` // SSID status list
}
type ResponseWirelessGetDeviceWirelessStatusBasicServiceSets struct {
	Band         string `json:"band,omitempty"`         // Frequency range used by wireless network
	Broadcasting *bool  `json:"broadcasting,omitempty"` // Whether the SSID is broadcasting based on an availability schedule
	Bssid        string `json:"bssid,omitempty"`        // Unique identifier of wireless access point
	Channel      *int   `json:"channel,omitempty"`      // Frequency channel used by wireless network
	ChannelWidth string `json:"channelWidth,omitempty"` // Width of frequency channel used by wireless network
	Enabled      *bool  `json:"enabled,omitempty"`      // Status of wireless network
	Power        string `json:"power,omitempty"`        // Strength of wireless signal
	SSIDName     string `json:"ssidName,omitempty"`     // Name of wireless network
	SSIDNumber   *int   `json:"ssidNumber,omitempty"`   // Unique identifier of wireless network
	Visible      *bool  `json:"visible,omitempty"`      // Whether the SSID is advertised or hidden
}
type ResponseWirelessGetNetworkWirelessAirMarshal []ResponseItemWirelessGetNetworkWirelessAirMarshal // Array of ResponseWirelessGetNetworkWirelessAirMarshal
type ResponseItemWirelessGetNetworkWirelessAirMarshal struct {
	Bssids        *[]ResponseItemWirelessGetNetworkWirelessAirMarshalBssids `json:"bssids,omitempty"`        //
	Channels      []string                                                  `json:"channels,omitempty"`      //
	FirstSeen     *int                                                      `json:"firstSeen,omitempty"`     //
	LastSeen      *int                                                      `json:"lastSeen,omitempty"`      //
	SSID          string                                                    `json:"ssid,omitempty"`          //
	WiredLastSeen *int                                                      `json:"wiredLastSeen,omitempty"` //
	WiredMacs     []string                                                  `json:"wiredMacs,omitempty"`     //
	WiredVLANs    []string                                                  `json:"wiredVlans,omitempty"`    //
}
type ResponseItemWirelessGetNetworkWirelessAirMarshalBssids struct {
	Bssid      string                                                              `json:"bssid,omitempty"`      //
	Contained  *bool                                                               `json:"contained,omitempty"`  //
	DetectedBy *[]ResponseItemWirelessGetNetworkWirelessAirMarshalBssidsDetectedBy `json:"detectedBy,omitempty"` //
}
type ResponseItemWirelessGetNetworkWirelessAirMarshalBssidsDetectedBy struct {
	Device string `json:"device,omitempty"` //
	Rssi   *int   `json:"rssi,omitempty"`   //
}
type ResponseWirelessGetNetworkWirelessAlternateManagementInterface struct {
	AccessPoints *[]ResponseWirelessGetNetworkWirelessAlternateManagementInterfaceAccessPoints `json:"accessPoints,omitempty"` //
	Enabled      *bool                                                                         `json:"enabled,omitempty"`      //
	Protocols    []string                                                                      `json:"protocols,omitempty"`    //
	VLANID       *int                                                                          `json:"vlanId,omitempty"`       //
}
type ResponseWirelessGetNetworkWirelessAlternateManagementInterfaceAccessPoints struct {
	AlternateManagementIP string `json:"alternateManagementIp,omitempty"` //
	DNS1                  string `json:"dns1,omitempty"`                  //
	DNS2                  string `json:"dns2,omitempty"`                  //
	Gateway               string `json:"gateway,omitempty"`               //
	Serial                string `json:"serial,omitempty"`                //
	SubnetMask            string `json:"subnetMask,omitempty"`            //
}
type ResponseWirelessUpdateNetworkWirelessAlternateManagementInterface interface{}
type ResponseWirelessGetNetworkWirelessBilling struct {
	Currency string                                            `json:"currency,omitempty"` // The currency code of this node group's billing plans
	Plans    *[]ResponseWirelessGetNetworkWirelessBillingPlans `json:"plans,omitempty"`    // Array of billing plans in the node group. (Can configure a maximum of 5)
}
type ResponseWirelessGetNetworkWirelessBillingPlans struct {
	BandwidthLimits *ResponseWirelessGetNetworkWirelessBillingPlansBandwidthLimits `json:"bandwidthLimits,omitempty"` // The uplink bandwidth settings for the pricing plan.
	ID              string                                                         `json:"id,omitempty"`              // The id of the pricing plan to update.
	Price           *float64                                                       `json:"price,omitempty"`           // The price of the billing plan.
	TimeLimit       string                                                         `json:"timeLimit,omitempty"`       // The time limit of the pricing plan in minutes.
}
type ResponseWirelessGetNetworkWirelessBillingPlansBandwidthLimits struct {
	LimitDown *int `json:"limitDown,omitempty"` // The maximum download limit (integer, in Kbps).
	LimitUp   *int `json:"limitUp,omitempty"`   // The maximum upload limit (integer, in Kbps).
}
type ResponseWirelessUpdateNetworkWirelessBilling struct {
	Currency string                                               `json:"currency,omitempty"` // The currency code of this node group's billing plans
	Plans    *[]ResponseWirelessUpdateNetworkWirelessBillingPlans `json:"plans,omitempty"`    // Array of billing plans in the node group. (Can configure a maximum of 5)
}
type ResponseWirelessUpdateNetworkWirelessBillingPlans struct {
	BandwidthLimits *ResponseWirelessUpdateNetworkWirelessBillingPlansBandwidthLimits `json:"bandwidthLimits,omitempty"` // The uplink bandwidth settings for the pricing plan.
	ID              string                                                            `json:"id,omitempty"`              // The id of the pricing plan to update.
	Price           *float64                                                          `json:"price,omitempty"`           // The price of the billing plan.
	TimeLimit       string                                                            `json:"timeLimit,omitempty"`       // The time limit of the pricing plan in minutes.
}
type ResponseWirelessUpdateNetworkWirelessBillingPlansBandwidthLimits struct {
	LimitDown *int `json:"limitDown,omitempty"` // The maximum download limit (integer, in Kbps).
	LimitUp   *int `json:"limitUp,omitempty"`   // The maximum upload limit (integer, in Kbps).
}
type ResponseWirelessGetNetworkWirelessBluetoothSettings struct {
	AdvertisingEnabled       *bool  `json:"advertisingEnabled,omitempty"`       // Whether APs will advertise beacons.
	EslEnabled               *bool  `json:"eslEnabled,omitempty"`               // Whether ESL is enabled on this network.
	Major                    *int   `json:"major,omitempty"`                    // The major number to be used in the beacon identifier. Only valid in 'Non-unique' mode.
	MajorMinorAssignmentMode string `json:"majorMinorAssignmentMode,omitempty"` // The way major and minor number should be assigned to nodes in the network. ('Unique', 'Non-unique')
	Minor                    *int   `json:"minor,omitempty"`                    // The minor number to be used in the beacon identifier. Only valid in 'Non-unique' mode.
	ScanningEnabled          *bool  `json:"scanningEnabled,omitempty"`          // Whether APs will scan for Bluetooth enabled clients.
	UUID                     string `json:"uuid,omitempty"`                     // The UUID to be used in the beacon identifier.
}
type ResponseWirelessUpdateNetworkWirelessBluetoothSettings struct {
	AdvertisingEnabled       *bool  `json:"advertisingEnabled,omitempty"`       // Whether APs will advertise beacons.
	EslEnabled               *bool  `json:"eslEnabled,omitempty"`               // Whether ESL is enabled on this network.
	Major                    *int   `json:"major,omitempty"`                    // The major number to be used in the beacon identifier. Only valid in 'Non-unique' mode.
	MajorMinorAssignmentMode string `json:"majorMinorAssignmentMode,omitempty"` // The way major and minor number should be assigned to nodes in the network. ('Unique', 'Non-unique')
	Minor                    *int   `json:"minor,omitempty"`                    // The minor number to be used in the beacon identifier. Only valid in 'Non-unique' mode.
	ScanningEnabled          *bool  `json:"scanningEnabled,omitempty"`          // Whether APs will scan for Bluetooth enabled clients.
	UUID                     string `json:"uuid,omitempty"`                     // The UUID to be used in the beacon identifier.
}
type ResponseWirelessGetNetworkWirelessChannelUtilizationHistory []ResponseItemWirelessGetNetworkWirelessChannelUtilizationHistory // Array of ResponseWirelessGetNetworkWirelessChannelUtilizationHistory
type ResponseItemWirelessGetNetworkWirelessChannelUtilizationHistory struct {
	EndTs               string   `json:"endTs,omitempty"`               // The end time of the query range
	StartTs             string   `json:"startTs,omitempty"`             // The start time of the query range
	Utilization80211    *float64 `json:"utilization80211,omitempty"`    // Average wifi utilization
	UtilizationNon80211 *float64 `json:"utilizationNon80211,omitempty"` // Average signal interference
	UtilizationTotal    *float64 `json:"utilizationTotal,omitempty"`    // Total channel utilization
}
type ResponseWirelessGetNetworkWirelessClientCountHistory []ResponseItemWirelessGetNetworkWirelessClientCountHistory // Array of ResponseWirelessGetNetworkWirelessClientCountHistory
type ResponseItemWirelessGetNetworkWirelessClientCountHistory struct {
	ClientCount *int   `json:"clientCount,omitempty"` // Number of connected clients
	EndTs       string `json:"endTs,omitempty"`       // The end time of the query range
	StartTs     string `json:"startTs,omitempty"`     // The start time of the query range
}
type ResponseWirelessGetNetworkWirelessClientsConnectionStats []ResponseItemWirelessGetNetworkWirelessClientsConnectionStats // Array of ResponseWirelessGetNetworkWirelessClientsConnectionStats
type ResponseItemWirelessGetNetworkWirelessClientsConnectionStats struct {
	ConnectionStats *ResponseItemWirelessGetNetworkWirelessClientsConnectionStatsConnectionStats `json:"connectionStats,omitempty"` //
	Mac             string                                                                       `json:"mac,omitempty"`             //
}
type ResponseItemWirelessGetNetworkWirelessClientsConnectionStatsConnectionStats struct {
	Assoc   *int `json:"assoc,omitempty"`   //
	Auth    *int `json:"auth,omitempty"`    //
	Dhcp    *int `json:"dhcp,omitempty"`    //
	DNS     *int `json:"dns,omitempty"`     //
	Success *int `json:"success,omitempty"` //
}
type ResponseWirelessGetNetworkWirelessClientsLatencyStats []ResponseItemWirelessGetNetworkWirelessClientsLatencyStats // Array of ResponseWirelessGetNetworkWirelessClientsLatencyStats
type ResponseItemWirelessGetNetworkWirelessClientsLatencyStats struct {
	LatencyStats *ResponseItemWirelessGetNetworkWirelessClientsLatencyStatsLatencyStats `json:"latencyStats,omitempty"` //
	Mac          string                                                                 `json:"mac,omitempty"`          //
}
type ResponseItemWirelessGetNetworkWirelessClientsLatencyStatsLatencyStats struct {
	BackgroundTraffic *ResponseItemWirelessGetNetworkWirelessClientsLatencyStatsLatencyStatsBackgroundTraffic `json:"backgroundTraffic,omitempty"` //
	BestEffortTraffic string                                                                                  `json:"bestEffortTraffic,omitempty"` //
	VideoTraffic      string                                                                                  `json:"videoTraffic,omitempty"`      //
	VoiceTraffic      string                                                                                  `json:"voiceTraffic,omitempty"`      //
}
type ResponseItemWirelessGetNetworkWirelessClientsLatencyStatsLatencyStatsBackgroundTraffic struct {
	Avg             *float64                                                                                               `json:"avg,omitempty"`             //
	RawDistribution *ResponseItemWirelessGetNetworkWirelessClientsLatencyStatsLatencyStatsBackgroundTrafficRawDistribution `json:"rawDistribution,omitempty"` //
}
type ResponseItemWirelessGetNetworkWirelessClientsLatencyStatsLatencyStatsBackgroundTrafficRawDistribution struct {
	Status0    *int `json:"0,omitempty"`    //
	Status1    *int `json:"1,omitempty"`    //
	Status1024 *int `json:"1024,omitempty"` //
	Status128  *int `json:"128,omitempty"`  //
	Status16   *int `json:"16,omitempty"`   //
	Status2    *int `json:"2,omitempty"`    //
	Status2048 *int `json:"2048,omitempty"` //
	Status256  *int `json:"256,omitempty"`  //
	Status32   *int `json:"32,omitempty"`   //
	Status4    *int `json:"4,omitempty"`    //
	Status512  *int `json:"512,omitempty"`  //
	Status64   *int `json:"64,omitempty"`   //
	Status8    *int `json:"8,omitempty"`    //
}
type ResponseWirelessGetNetworkWirelessClientConnectionStats struct {
	ConnectionStats *ResponseWirelessGetNetworkWirelessClientConnectionStatsConnectionStats `json:"connectionStats,omitempty"` //
	Mac             string                                                                  `json:"mac,omitempty"`             //
}
type ResponseWirelessGetNetworkWirelessClientConnectionStatsConnectionStats struct {
	Assoc   *int `json:"assoc,omitempty"`   //
	Auth    *int `json:"auth,omitempty"`    //
	Dhcp    *int `json:"dhcp,omitempty"`    //
	DNS     *int `json:"dns,omitempty"`     //
	Success *int `json:"success,omitempty"` //
}
type ResponseWirelessGetNetworkWirelessClientConnectivityEvents []ResponseItemWirelessGetNetworkWirelessClientConnectivityEvents // Array of ResponseWirelessGetNetworkWirelessClientConnectivityEvents
type ResponseItemWirelessGetNetworkWirelessClientConnectivityEvents struct {
	Band         *int                                                                     `json:"band,omitempty"`         //
	Channel      *int                                                                     `json:"channel,omitempty"`      //
	DeviceSerial string                                                                   `json:"deviceSerial,omitempty"` //
	DurationMs   *float64                                                                 `json:"durationMs,omitempty"`   //
	EventData    *ResponseItemWirelessGetNetworkWirelessClientConnectivityEventsEventData `json:"eventData,omitempty"`    //
	OccurredAt   *int                                                                     `json:"occurredAt,omitempty"`   //
	Rssi         *int                                                                     `json:"rssi,omitempty"`         //
	Severity     string                                                                   `json:"severity,omitempty"`     //
	SSIDNumber   *int                                                                     `json:"ssidNumber,omitempty"`   //
	Subtype      string                                                                   `json:"subtype,omitempty"`      //
	Type         string                                                                   `json:"type,omitempty"`         //
}
type ResponseItemWirelessGetNetworkWirelessClientConnectivityEventsEventData struct {
	ClientIP string `json:"clientIp,omitempty"` //
}
type ResponseWirelessGetNetworkWirelessClientLatencyHistory []ResponseItemWirelessGetNetworkWirelessClientLatencyHistory // Array of ResponseWirelessGetNetworkWirelessClientLatencyHistory
type ResponseItemWirelessGetNetworkWirelessClientLatencyHistory struct {
	LatencyBinsByCategory *ResponseItemWirelessGetNetworkWirelessClientLatencyHistoryLatencyBinsByCategory `json:"latencyBinsByCategory,omitempty"` // The latency buckets by category
	T0                    *int                                                                             `json:"t0,omitempty"`                    // The latency history bucket start time in seconds
	T1                    *int                                                                             `json:"t1,omitempty"`                    // The latency history bucket end time in seconds
}
type ResponseItemWirelessGetNetworkWirelessClientLatencyHistoryLatencyBinsByCategory struct {
	BackgroundTraffic *ResponseItemWirelessGetNetworkWirelessClientLatencyHistoryLatencyBinsByCategoryBackgroundTraffic `json:"backgroundTraffic,omitempty"` // The time bucket's background traffic latency history
	BestEffortTraffic *ResponseItemWirelessGetNetworkWirelessClientLatencyHistoryLatencyBinsByCategoryBestEffortTraffic `json:"bestEffortTraffic,omitempty"` // The time bucket's best effort traffic latency history
	VideoTraffic      *ResponseItemWirelessGetNetworkWirelessClientLatencyHistoryLatencyBinsByCategoryVideoTraffic      `json:"videoTraffic,omitempty"`      // The time bucket's video traffic latency history
	VoiceTraffic      *ResponseItemWirelessGetNetworkWirelessClientLatencyHistoryLatencyBinsByCategoryVoiceTraffic      `json:"voiceTraffic,omitempty"`      // The time bucket's voice traffic latency history
}
type ResponseItemWirelessGetNetworkWirelessClientLatencyHistoryLatencyBinsByCategoryBackgroundTraffic struct {
	Status05    *int `json:"0.5,omitempty"`    // The latency bucket for background traffic in 0.5 seconds
	Status10    *int `json:"1.0,omitempty"`    // The latency bucket for background traffic in 1.0 seconds
	Status10240 *int `json:"1024.0,omitempty"` // The latency bucket for background traffic in 1024.0 seconds
	Status1280  *int `json:"128.0,omitempty"`  // The latency bucket for background traffic in 128.0 seconds
	Status160   *int `json:"16.0,omitempty"`   // The latency bucket for background traffic in 16.0 seconds
	Status20    *int `json:"2.0,omitempty"`    // The latency bucket for background traffic in 2.0 seconds
	Status20480 *int `json:"2048.0,omitempty"` // The latency bucket for background traffic in 2048.0 seconds
	Status2560  *int `json:"256.0,omitempty"`  // The latency bucket for background traffic in 256.0 seconds
	Status320   *int `json:"32.0,omitempty"`   // The latency bucket for background traffic in 32.0 seconds
	Status40    *int `json:"4.0,omitempty"`    // The latency bucket for background traffic in 4.0 seconds
	Status5120  *int `json:"512.0,omitempty"`  // The latency bucket for background traffic in 512.0 seconds
	Status640   *int `json:"64.0,omitempty"`   // The latency bucket for background traffic in 64.0 seconds
	Status80    *int `json:"8.0,omitempty"`    // The latency bucket for background traffic in 8.0 seconds
}
type ResponseItemWirelessGetNetworkWirelessClientLatencyHistoryLatencyBinsByCategoryBestEffortTraffic struct {
	Status05    *int `json:"0.5,omitempty"`    // The latency bucket for best effort traffic in 0.5 seconds
	Status10    *int `json:"1.0,omitempty"`    // The latency bucket for best effort traffic in 1.0 seconds
	Status10240 *int `json:"1024.0,omitempty"` // The latency bucket for best effort traffic in 1024.0 seconds
	Status1280  *int `json:"128.0,omitempty"`  // The latency bucket for best effort traffic in 128.0 seconds
	Status160   *int `json:"16.0,omitempty"`   // The latency bucket for best effort traffic in 16.0 seconds
	Status20    *int `json:"2.0,omitempty"`    // The latency bucket for best effort traffic in 2.0 seconds
	Status20480 *int `json:"2048.0,omitempty"` // The latency bucket for best effort traffic in 2048.0 seconds
	Status2560  *int `json:"256.0,omitempty"`  // The latency bucket for best effort traffic in 256.0 seconds
	Status320   *int `json:"32.0,omitempty"`   // The latency bucket for best effort traffic in 32.0 seconds
	Status40    *int `json:"4.0,omitempty"`    // The latency bucket for best effort traffic in 4.0 seconds
	Status5120  *int `json:"512.0,omitempty"`  // The latency bucket for best effort traffic in 512.0 seconds
	Status640   *int `json:"64.0,omitempty"`   // The latency bucket for best effort traffic in 64.0 seconds
	Status80    *int `json:"8.0,omitempty"`    // The latency bucket for best effort traffic in 8.0 seconds
}
type ResponseItemWirelessGetNetworkWirelessClientLatencyHistoryLatencyBinsByCategoryVideoTraffic struct {
	Status05    *int `json:"0.5,omitempty"`    // The latency bucket for video traffic in 0.5 seconds
	Status10    *int `json:"1.0,omitempty"`    // The latency bucket for video traffic in 1.0 seconds
	Status10240 *int `json:"1024.0,omitempty"` // The latency bucket for video traffic in 1024.0 seconds
	Status1280  *int `json:"128.0,omitempty"`  // The latency bucket for video traffic in 128.0 seconds
	Status160   *int `json:"16.0,omitempty"`   // The latency bucket for video traffic in 16.0 seconds
	Status20    *int `json:"2.0,omitempty"`    // The latency bucket for video traffic in 2.0 seconds
	Status20480 *int `json:"2048.0,omitempty"` // The latency bucket for video traffic in 2048.0 seconds
	Status2560  *int `json:"256.0,omitempty"`  // The latency bucket for video traffic in 256.0 seconds
	Status320   *int `json:"32.0,omitempty"`   // The latency bucket for video traffic in 32.0 seconds
	Status40    *int `json:"4.0,omitempty"`    // The latency bucket for video traffic in 4.0 seconds
	Status5120  *int `json:"512.0,omitempty"`  // The latency bucket for video traffic in 512.0 seconds
	Status640   *int `json:"64.0,omitempty"`   // The latency bucket for video traffic in 64.0 seconds
	Status80    *int `json:"8.0,omitempty"`    // The latency bucket for video traffic in 8.0 seconds
}
type ResponseItemWirelessGetNetworkWirelessClientLatencyHistoryLatencyBinsByCategoryVoiceTraffic struct {
	Status05    *int `json:"0.5,omitempty"`    // The latency bucket for voice traffic in 0.5 seconds
	Status10    *int `json:"1.0,omitempty"`    // The latency bucket for voice traffic in 1.0 seconds
	Status10240 *int `json:"1024.0,omitempty"` // The latency bucket for voice traffic in 1024.0 seconds
	Status1280  *int `json:"128.0,omitempty"`  // The latency bucket for voice traffic in 128.0 seconds
	Status160   *int `json:"16.0,omitempty"`   // The latency bucket for voice traffic in 16.0 seconds
	Status20    *int `json:"2.0,omitempty"`    // The latency bucket for voice traffic in 2.0 seconds
	Status20480 *int `json:"2048.0,omitempty"` // The latency bucket for voice traffic in 2048.0 seconds
	Status2560  *int `json:"256.0,omitempty"`  // The latency bucket for voice traffic in 256.0 seconds
	Status320   *int `json:"32.0,omitempty"`   // The latency bucket for voice traffic in 32.0 seconds
	Status40    *int `json:"4.0,omitempty"`    // The latency bucket for voice traffic in 4.0 seconds
	Status5120  *int `json:"512.0,omitempty"`  // The latency bucket for voice traffic in 512.0 seconds
	Status640   *int `json:"64.0,omitempty"`   // The latency bucket for voice traffic in 64.0 seconds
	Status80    *int `json:"8.0,omitempty"`    // The latency bucket for voice traffic in 8.0 seconds
}
type ResponseWirelessGetNetworkWirelessClientLatencyStats struct {
	LatencyStats *ResponseWirelessGetNetworkWirelessClientLatencyStatsLatencyStats `json:"latencyStats,omitempty"` //
	Mac          string                                                            `json:"mac,omitempty"`          //
}
type ResponseWirelessGetNetworkWirelessClientLatencyStatsLatencyStats struct {
	BackgroundTraffic *ResponseWirelessGetNetworkWirelessClientLatencyStatsLatencyStatsBackgroundTraffic `json:"backgroundTraffic,omitempty"` //
	BestEffortTraffic string                                                                             `json:"bestEffortTraffic,omitempty"` //
	VideoTraffic      string                                                                             `json:"videoTraffic,omitempty"`      //
	VoiceTraffic      string                                                                             `json:"voiceTraffic,omitempty"`      //
}
type ResponseWirelessGetNetworkWirelessClientLatencyStatsLatencyStatsBackgroundTraffic struct {
	Avg             *float64                                                                                          `json:"avg,omitempty"`             //
	RawDistribution *ResponseWirelessGetNetworkWirelessClientLatencyStatsLatencyStatsBackgroundTrafficRawDistribution `json:"rawDistribution,omitempty"` //
}
type ResponseWirelessGetNetworkWirelessClientLatencyStatsLatencyStatsBackgroundTrafficRawDistribution struct {
	Status0    *int `json:"0,omitempty"`    //
	Status1    *int `json:"1,omitempty"`    //
	Status1024 *int `json:"1024,omitempty"` //
	Status128  *int `json:"128,omitempty"`  //
	Status16   *int `json:"16,omitempty"`   //
	Status2    *int `json:"2,omitempty"`    //
	Status2048 *int `json:"2048,omitempty"` //
	Status256  *int `json:"256,omitempty"`  //
	Status32   *int `json:"32,omitempty"`   //
	Status4    *int `json:"4,omitempty"`    //
	Status512  *int `json:"512,omitempty"`  //
	Status64   *int `json:"64,omitempty"`   //
	Status8    *int `json:"8,omitempty"`    //
}
type ResponseWirelessGetNetworkWirelessConnectionStats struct {
	Assoc   *int `json:"assoc,omitempty"`   // The number of failed association attempts
	Auth    *int `json:"auth,omitempty"`    // The number of failed authentication attempts
	Dhcp    *int `json:"dhcp,omitempty"`    // The number of failed DHCP attempts
	DNS     *int `json:"dns,omitempty"`     // The number of failed DNS attempts
	Success *int `json:"success,omitempty"` // The number of successful connection attempts
}
type ResponseWirelessGetNetworkWirelessDataRateHistory []ResponseItemWirelessGetNetworkWirelessDataRateHistory // Array of ResponseWirelessGetNetworkWirelessDataRateHistory
type ResponseItemWirelessGetNetworkWirelessDataRateHistory struct {
	AverageKbps  *int   `json:"averageKbps,omitempty"`  // Average data rate in kilobytes-per-second
	DownloadKbps *int   `json:"downloadKbps,omitempty"` // Download rate in kilobytes-per-second
	EndTs        string `json:"endTs,omitempty"`        // The end time of the query range
	StartTs      string `json:"startTs,omitempty"`      // The start time of the query range
	UploadKbps   *int   `json:"uploadKbps,omitempty"`   // Upload rate in kilobytes-per-second
}
type ResponseWirelessGetNetworkWirelessDevicesConnectionStats []ResponseItemWirelessGetNetworkWirelessDevicesConnectionStats // Array of ResponseWirelessGetNetworkWirelessDevicesConnectionStats
type ResponseItemWirelessGetNetworkWirelessDevicesConnectionStats struct {
	ConnectionStats *ResponseItemWirelessGetNetworkWirelessDevicesConnectionStatsConnectionStats `json:"connectionStats,omitempty"` // The connection stats of the device
	Serial          string                                                                       `json:"serial,omitempty"`          // The serial number for the device
}
type ResponseItemWirelessGetNetworkWirelessDevicesConnectionStatsConnectionStats struct {
	Assoc   *int `json:"assoc,omitempty"`   // The number of failed association attempts
	Auth    *int `json:"auth,omitempty"`    // The number of failed authentication attempts
	Dhcp    *int `json:"dhcp,omitempty"`    // The number of failed DHCP attempts
	DNS     *int `json:"dns,omitempty"`     // The number of failed DNS attempts
	Success *int `json:"success,omitempty"` // The number of successful connection attempts
}
type ResponseWirelessGetNetworkWirelessDevicesLatencyStats []ResponseItemWirelessGetNetworkWirelessDevicesLatencyStats // Array of ResponseWirelessGetNetworkWirelessDevicesLatencyStats
type ResponseItemWirelessGetNetworkWirelessDevicesLatencyStats struct {
	LatencyStats *ResponseItemWirelessGetNetworkWirelessDevicesLatencyStatsLatencyStats `json:"latencyStats,omitempty"` //
	Serial       string                                                                 `json:"serial,omitempty"`       //
}
type ResponseItemWirelessGetNetworkWirelessDevicesLatencyStatsLatencyStats struct {
	BackgroundTraffic *ResponseItemWirelessGetNetworkWirelessDevicesLatencyStatsLatencyStatsBackgroundTraffic `json:"backgroundTraffic,omitempty"` //
	BestEffortTraffic string                                                                                  `json:"bestEffortTraffic,omitempty"` //
	VideoTraffic      string                                                                                  `json:"videoTraffic,omitempty"`      //
	VoiceTraffic      string                                                                                  `json:"voiceTraffic,omitempty"`      //
}
type ResponseItemWirelessGetNetworkWirelessDevicesLatencyStatsLatencyStatsBackgroundTraffic struct {
	Avg             *float64                                                                                               `json:"avg,omitempty"`             //
	RawDistribution *ResponseItemWirelessGetNetworkWirelessDevicesLatencyStatsLatencyStatsBackgroundTrafficRawDistribution `json:"rawDistribution,omitempty"` //
}
type ResponseItemWirelessGetNetworkWirelessDevicesLatencyStatsLatencyStatsBackgroundTrafficRawDistribution struct {
	Status0    *int `json:"0,omitempty"`    //
	Status1    *int `json:"1,omitempty"`    //
	Status1024 *int `json:"1024,omitempty"` //
	Status128  *int `json:"128,omitempty"`  //
	Status16   *int `json:"16,omitempty"`   //
	Status2    *int `json:"2,omitempty"`    //
	Status2048 *int `json:"2048,omitempty"` //
	Status256  *int `json:"256,omitempty"`  //
	Status32   *int `json:"32,omitempty"`   //
	Status4    *int `json:"4,omitempty"`    //
	Status512  *int `json:"512,omitempty"`  //
	Status64   *int `json:"64,omitempty"`   //
	Status8    *int `json:"8,omitempty"`    //
}
type ResponseWirelessGetNetworkWirelessEthernetPortsProfiles []ResponseItemWirelessGetNetworkWirelessEthernetPortsProfiles // Array of ResponseWirelessGetNetworkWirelessEthernetPortsProfiles
type ResponseItemWirelessGetNetworkWirelessEthernetPortsProfiles struct {
	IsDefault *bool                                                                  `json:"isDefault,omitempty"` // Is default profile
	Name      string                                                                 `json:"name,omitempty"`      // AP port profile name
	Ports     *[]ResponseItemWirelessGetNetworkWirelessEthernetPortsProfilesPorts    `json:"ports,omitempty"`     // Ports config
	ProfileID string                                                                 `json:"profileId,omitempty"` // AP port profile ID
	UsbPorts  *[]ResponseItemWirelessGetNetworkWirelessEthernetPortsProfilesUsbPorts `json:"usbPorts,omitempty"`  // Usb ports config
}
type ResponseItemWirelessGetNetworkWirelessEthernetPortsProfilesPorts struct {
	Enabled    *bool  `json:"enabled,omitempty"`    // Enabled
	Name       string `json:"name,omitempty"`       // Name
	Number     *int   `json:"number,omitempty"`     // Number
	PskGroupID string `json:"pskGroupId,omitempty"` // PSK Group number
	SSID       *int   `json:"ssid,omitempty"`       // Ssid number
}
type ResponseItemWirelessGetNetworkWirelessEthernetPortsProfilesUsbPorts struct {
	Enabled *bool  `json:"enabled,omitempty"` // Enabled
	Name    string `json:"name,omitempty"`    // Name
	SSID    *int   `json:"ssid,omitempty"`    // Ssid number
}
type ResponseWirelessCreateNetworkWirelessEthernetPortsProfile struct {
	IsDefault *bool                                                                `json:"isDefault,omitempty"` // Is default profile
	Name      string                                                               `json:"name,omitempty"`      // AP port profile name
	Ports     *[]ResponseWirelessCreateNetworkWirelessEthernetPortsProfilePorts    `json:"ports,omitempty"`     // Ports config
	ProfileID string                                                               `json:"profileId,omitempty"` // AP port profile ID
	UsbPorts  *[]ResponseWirelessCreateNetworkWirelessEthernetPortsProfileUsbPorts `json:"usbPorts,omitempty"`  // Usb ports config
}
type ResponseWirelessCreateNetworkWirelessEthernetPortsProfilePorts struct {
	Enabled    *bool  `json:"enabled,omitempty"`    // Enabled
	Name       string `json:"name,omitempty"`       // Name
	Number     *int   `json:"number,omitempty"`     // Number
	PskGroupID string `json:"pskGroupId,omitempty"` // PSK Group number
	SSID       *int   `json:"ssid,omitempty"`       // Ssid number
}
type ResponseWirelessCreateNetworkWirelessEthernetPortsProfileUsbPorts struct {
	Enabled *bool  `json:"enabled,omitempty"` // Enabled
	Name    string `json:"name,omitempty"`    // Name
	SSID    *int   `json:"ssid,omitempty"`    // Ssid number
}
type ResponseWirelessAssignNetworkWirelessEthernetPortsProfiles struct {
	ProfileID string   `json:"profileId,omitempty"` // AP profile ID
	Serials   []string `json:"serials,omitempty"`   // List of updated AP serials
}
type ResponseWirelessSetNetworkWirelessEthernetPortsProfilesDefault struct {
	ProfileID string `json:"profileId,omitempty"` // AP profile ID
}
type ResponseWirelessGetNetworkWirelessEthernetPortsProfile struct {
	IsDefault *bool                                                             `json:"isDefault,omitempty"` // Is default profile
	Name      string                                                            `json:"name,omitempty"`      // AP port profile name
	Ports     *[]ResponseWirelessGetNetworkWirelessEthernetPortsProfilePorts    `json:"ports,omitempty"`     // Ports config
	ProfileID string                                                            `json:"profileId,omitempty"` // AP port profile ID
	UsbPorts  *[]ResponseWirelessGetNetworkWirelessEthernetPortsProfileUsbPorts `json:"usbPorts,omitempty"`  // Usb ports config
}
type ResponseWirelessGetNetworkWirelessEthernetPortsProfilePorts struct {
	Enabled    *bool  `json:"enabled,omitempty"`    // Enabled
	Name       string `json:"name,omitempty"`       // Name
	Number     *int   `json:"number,omitempty"`     // Number
	PskGroupID string `json:"pskGroupId,omitempty"` // PSK Group number
	SSID       *int   `json:"ssid,omitempty"`       // Ssid number
}
type ResponseWirelessGetNetworkWirelessEthernetPortsProfileUsbPorts struct {
	Enabled *bool  `json:"enabled,omitempty"` // Enabled
	Name    string `json:"name,omitempty"`    // Name
	SSID    *int   `json:"ssid,omitempty"`    // Ssid number
}
type ResponseWirelessUpdateNetworkWirelessEthernetPortsProfile struct {
	IsDefault *bool                                                                `json:"isDefault,omitempty"` // Is default profile
	Name      string                                                               `json:"name,omitempty"`      // AP port profile name
	Ports     *[]ResponseWirelessUpdateNetworkWirelessEthernetPortsProfilePorts    `json:"ports,omitempty"`     // Ports config
	ProfileID string                                                               `json:"profileId,omitempty"` // AP port profile ID
	UsbPorts  *[]ResponseWirelessUpdateNetworkWirelessEthernetPortsProfileUsbPorts `json:"usbPorts,omitempty"`  // Usb ports config
}
type ResponseWirelessUpdateNetworkWirelessEthernetPortsProfilePorts struct {
	Enabled    *bool  `json:"enabled,omitempty"`    // Enabled
	Name       string `json:"name,omitempty"`       // Name
	Number     *int   `json:"number,omitempty"`     // Number
	PskGroupID string `json:"pskGroupId,omitempty"` // PSK Group number
	SSID       *int   `json:"ssid,omitempty"`       // Ssid number
}
type ResponseWirelessUpdateNetworkWirelessEthernetPortsProfileUsbPorts struct {
	Enabled *bool  `json:"enabled,omitempty"` // Enabled
	Name    string `json:"name,omitempty"`    // Name
	SSID    *int   `json:"ssid,omitempty"`    // Ssid number
}
type ResponseWirelessGetNetworkWirelessFailedConnections []ResponseItemWirelessGetNetworkWirelessFailedConnections // Array of ResponseWirelessGetNetworkWirelessFailedConnections
type ResponseItemWirelessGetNetworkWirelessFailedConnections struct {
	ClientMac   string `json:"clientMac,omitempty"`   // Client Mac
	FailureStep string `json:"failureStep,omitempty"` // The failed onboarding step. One of: assoc, auth, dhcp, dns.
	Serial      string `json:"serial,omitempty"`      // Serial Number
	SSIDNumber  *int   `json:"ssidNumber,omitempty"`  // SSID Number
	Ts          string `json:"ts,omitempty"`          // The timestamp when the client mac failed
	Type        string `json:"type,omitempty"`        // The failure type in the onboarding step
	VLAN        *int   `json:"vlan,omitempty"`        // LAN
}
type ResponseWirelessGetNetworkWirelessLatencyHistory []ResponseItemWirelessGetNetworkWirelessLatencyHistory // Array of ResponseWirelessGetNetworkWirelessLatencyHistory
type ResponseItemWirelessGetNetworkWirelessLatencyHistory struct {
	AvgLatencyMs *int   `json:"avgLatencyMs,omitempty"` // Average latency in milliseconds
	EndTs        string `json:"endTs,omitempty"`        // The end time of the query range
	StartTs      string `json:"startTs,omitempty"`      // The start time of the query range
}
type ResponseWirelessGetNetworkWirelessLatencyStats struct {
	BackgroundTraffic *ResponseWirelessGetNetworkWirelessLatencyStatsBackgroundTraffic `json:"backgroundTraffic,omitempty"` //
	BestEffortTraffic string                                                           `json:"bestEffortTraffic,omitempty"` //
	VideoTraffic      string                                                           `json:"videoTraffic,omitempty"`      //
	VoiceTraffic      string                                                           `json:"voiceTraffic,omitempty"`      //
}
type ResponseWirelessGetNetworkWirelessLatencyStatsBackgroundTraffic struct {
	Avg             *float64                                                                        `json:"avg,omitempty"`             //
	RawDistribution *ResponseWirelessGetNetworkWirelessLatencyStatsBackgroundTrafficRawDistribution `json:"rawDistribution,omitempty"` //
}
type ResponseWirelessGetNetworkWirelessLatencyStatsBackgroundTrafficRawDistribution struct {
	Status0    *int `json:"0,omitempty"`    //
	Status1    *int `json:"1,omitempty"`    //
	Status1024 *int `json:"1024,omitempty"` //
	Status128  *int `json:"128,omitempty"`  //
	Status16   *int `json:"16,omitempty"`   //
	Status2    *int `json:"2,omitempty"`    //
	Status2048 *int `json:"2048,omitempty"` //
	Status256  *int `json:"256,omitempty"`  //
	Status32   *int `json:"32,omitempty"`   //
	Status4    *int `json:"4,omitempty"`    //
	Status512  *int `json:"512,omitempty"`  //
	Status64   *int `json:"64,omitempty"`   //
	Status8    *int `json:"8,omitempty"`    //
}
type ResponseWirelessGetNetworkWirelessMeshStatuses []ResponseItemWirelessGetNetworkWirelessMeshStatuses // Array of ResponseWirelessGetNetworkWirelessMeshStatuses
type ResponseItemWirelessGetNetworkWirelessMeshStatuses struct {
	LatestMeshPerformance *ResponseItemWirelessGetNetworkWirelessMeshStatusesLatestMeshPerformance `json:"latestMeshPerformance,omitempty"` // Current metrics on how the mesh is performing.
	MeshRoute             []string                                                                 `json:"meshRoute,omitempty"`             // List of device serials that make up the mesh.
	Serial                string                                                                   `json:"serial,omitempty"`                // The serial number for the device.
}
type ResponseItemWirelessGetNetworkWirelessMeshStatusesLatestMeshPerformance struct {
	Mbps            *int   `json:"mbps,omitempty"`            // Average Mbps.
	Metric          *int   `json:"metric,omitempty"`          // Represents the quality of the entire route from the repeater access point to its gateway access point.
	UsagePercentage string `json:"usagePercentage,omitempty"` // Mesh utilization as a percentage.
}

type ResponseWirelessGetNetworkWirelessRfProfiles []ResponseItemWirelessGetNetworkWirelessRfProfiles

type ResponseItemWirelessGetNetworkWirelessRfProfiles struct {
	ApBandSettings         *ResponseWirelessGetNetworkWirelessRfProfilesApBandSettings     `json:"apBandSettings,omitempty"`         // Settings that will be enabled if selectionType is set to 'ap'.
	BandSelectionType      string                                                          `json:"bandSelectionType,omitempty"`      // Band selection can be set to either 'ssid' or 'ap'. This param is required on creation.
	ClientBalancingEnabled *bool                                                           `json:"clientBalancingEnabled,omitempty"` // Steers client to best available access point. Can be either true or false. Defaults to true.
	FiveGhzSettings        *ResponseWirelessGetNetworkWirelessRfProfilesFiveGhzSettings    `json:"fiveGhzSettings,omitempty"`        // Settings related to 5Ghz band
	ID                     string                                                          `json:"id,omitempty"`                     // The name of the new profile. Must be unique.
	MinBitrateType         string                                                          `json:"minBitrateType,omitempty"`         // Minimum bitrate can be set to either 'band' or 'ssid'. Defaults to band.
	Name                   string                                                          `json:"name,omitempty"`                   // The name of the new profile. Must be unique. This param is required on creation.
	NetworkID              string                                                          `json:"networkId,omitempty"`              // The network ID of the RF Profile
	PerSSIDSettings        *ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings    `json:"perSsidSettings,omitempty"`        // Per-SSID radio settings by number.
	SixGhzSettings         *ResponseWirelessGetNetworkWirelessRfProfilesSixGhzSettings     `json:"sixGhzSettings,omitempty"`         // Settings related to 6Ghz band. Only applicable to networks with 6Ghz capable APs
	Transmission           *ResponseWirelessGetNetworkWirelessRfProfilesTransmission       `json:"transmission,omitempty"`           // Settings related to radio transmission.
	TwoFourGhzSettings     *ResponseWirelessGetNetworkWirelessRfProfilesTwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"`     // Settings related to 2.4Ghz band
}
type ResponseWirelessGetNetworkWirelessRfProfilesApBandSettings struct {
	BandOperationMode   string                                                           `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'. Defaults to dual.
	BandSteeringEnabled *bool                                                            `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band. Can be either true or false. Defaults to true.
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilesApBandSettingsBands `json:"bands,omitempty"`               // Settings related to all bands
}
type ResponseWirelessGetNetworkWirelessRfProfilesApBandSettingsBands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessGetNetworkWirelessRfProfilesFiveGhzSettings struct {
	ChannelWidth      string `json:"channelWidth,omitempty"`      // Sets channel width (MHz) for 5Ghz band. Can be one of 'auto', '20', '40' or '80'. Defaults to auto.
	MaxPower          *int   `json:"maxPower,omitempty"`          // Sets max power (dBm) of 5Ghz band. Can be integer between 2 and 30. Defaults to 30.
	MinBitrate        *int   `json:"minBitrate,omitempty"`        // Sets min bitrate (Mbps) of 5Ghz band. Can be one of '6', '9', '12', '18', '24', '36', '48' or '54'. Defaults to 12.
	MinPower          *int   `json:"minPower,omitempty"`          // Sets min power (dBm) of 5Ghz band. Can be integer between 2 and 30. Defaults to 8.
	Rxsop             *int   `json:"rxsop,omitempty"`             // The RX-SOP level controls the sensitivity of the radio. It is strongly recommended to use RX-SOP only after consulting a wireless expert. RX-SOP can be configured in the range of -65 to -95 (dBm). A value of null will reset this to the default.
	ValidAutoChannels *[]int `json:"validAutoChannels,omitempty"` // Sets valid auto channels for 5Ghz band. Can be one of '36', '40', '44', '48', '52', '56', '60', '64', '100', '104', '108', '112', '116', '120', '124', '128', '132', '136', '140', '144', '149', '153', '157', '161' or '165'.Defaults to [36, 40, 44, 48, 52, 56, 60, 64, 100, 104, 108, 112, 116, 120, 124, 128, 132, 136, 140, 144, 149, 153, 157, 161, 165].
}
type ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings struct {
	Status0  *ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings0  `json:"0,omitempty"`  // Settings for SSID 0
	Status1  *ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings1  `json:"1,omitempty"`  // Settings for SSID 1
	Status10 *ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings10 `json:"10,omitempty"` // Settings for SSID 10
	Status11 *ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings11 `json:"11,omitempty"` // Settings for SSID 11
	Status12 *ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings12 `json:"12,omitempty"` // Settings for SSID 12
	Status13 *ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings13 `json:"13,omitempty"` // Settings for SSID 13
	Status14 *ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings14 `json:"14,omitempty"` // Settings for SSID 14
	Status2  *ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings2  `json:"2,omitempty"`  // Settings for SSID 2
	Status3  *ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings3  `json:"3,omitempty"`  // Settings for SSID 3
	Status4  *ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings4  `json:"4,omitempty"`  // Settings for SSID 4
	Status5  *ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings5  `json:"5,omitempty"`  // Settings for SSID 5
	Status6  *ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings6  `json:"6,omitempty"`  // Settings for SSID 6
	Status7  *ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings7  `json:"7,omitempty"`  // Settings for SSID 7
	Status8  *ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings8  `json:"8,omitempty"`  // Settings for SSID 8
	Status9  *ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings9  `json:"9,omitempty"`  // Settings for SSID 9
}
type ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings0 struct {
	BandOperationMode   string                                                             `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                              `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings0Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                               `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                             `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings0Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings1 struct {
	BandOperationMode   string                                                             `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                              `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings1Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                               `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                             `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings1Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings10 struct {
	BandOperationMode   string                                                              `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                               `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings10Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                                `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                              `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings10Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings11 struct {
	BandOperationMode   string                                                              `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                               `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings11Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                                `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                              `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings11Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings12 struct {
	BandOperationMode   string                                                              `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                               `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings12Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                                `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                              `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings12Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings13 struct {
	BandOperationMode   string                                                              `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                               `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings13Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                                `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                              `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings13Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings14 struct {
	BandOperationMode   string                                                              `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                               `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings14Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                                `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                              `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings14Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings2 struct {
	BandOperationMode   string                                                             `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                              `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings2Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                               `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                             `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings2Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings3 struct {
	BandOperationMode   string                                                             `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                              `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings3Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                               `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                             `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings3Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings4 struct {
	BandOperationMode   string                                                             `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                              `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings4Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                               `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                             `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings4Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings5 struct {
	BandOperationMode   string                                                             `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                              `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings5Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                               `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                             `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings5Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings6 struct {
	BandOperationMode   string                                                             `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                              `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings6Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                               `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                             `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings6Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings7 struct {
	BandOperationMode   string                                                             `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                              `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings7Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                               `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                             `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings7Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings8 struct {
	BandOperationMode   string                                                             `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                              `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings8Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                               `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                             `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings8Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings9 struct {
	BandOperationMode   string                                                             `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                              `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings9Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                               `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                             `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessGetNetworkWirelessRfProfilesPerSSIDSettings9Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessGetNetworkWirelessRfProfilesSixGhzSettings struct {
	ChannelWidth      string `json:"channelWidth,omitempty"`      // Sets channel width (MHz) for 6Ghz band. Can be one of '0', '20', '40', '80' or '160'. Defaults to auto.
	MaxPower          *int   `json:"maxPower,omitempty"`          // Sets max power (dBm) of 6Ghz band. Can be integer between 2 and 30. Defaults to 30.
	MinBitrate        *int   `json:"minBitrate,omitempty"`        // Sets min bitrate (Mbps) of 6Ghz band. Can be one of '6', '9', '12', '18', '24', '36', '48' or '54'. Defaults to 12.
	MinPower          *int   `json:"minPower,omitempty"`          // Sets min power (dBm) of 6Ghz band. Can be integer between 2 and 30. Defaults to 8.
	Rxsop             *int   `json:"rxsop,omitempty"`             // The RX-SOP level controls the sensitivity of the radio. It is strongly recommended to use RX-SOP only after consulting a wireless expert. RX-SOP can be configured in the range of -65 to -95 (dBm). A value of null will reset this to the default.
	ValidAutoChannels *[]int `json:"validAutoChannels,omitempty"` // Sets valid auto channels for 6Ghz band. Can be one of '1', '5', '9', '13', '17', '21', '25', '29', '33', '37', '41', '45', '49', '53', '57', '61', '65', '69', '73', '77', '81', '85', '89', '93', '97', '101', '105', '109', '113', '117', '121', '125', '129', '133', '137', '141', '145', '149', '153', '157', '161', '165', '169', '173', '177', '181', '185', '189', '193', '197', '201', '205', '209', '213', '217', '221', '225', '229' or '233'. Defaults to auto.
}
type ResponseWirelessGetNetworkWirelessRfProfilesTransmission struct {
	Enabled *bool `json:"enabled,omitempty"` // Toggle for radio transmission. When false, radios will not transmit at all.
}
type ResponseWirelessGetNetworkWirelessRfProfilesTwoFourGhzSettings struct {
	AxEnabled         *bool    `json:"axEnabled,omitempty"`         // Determines whether ax radio on 2.4Ghz band is on or off. Can be either true or false. If false, we highly recommend disabling band steering. Defaults to true.
	MaxPower          *int     `json:"maxPower,omitempty"`          // Sets max power (dBm) of 2.4Ghz band. Can be integer between 2 and 30. Defaults to 30.
	MinBitrate        *float64 `json:"minBitrate,omitempty"`        // Sets min bitrate (Mbps) of 2.4Ghz band. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'. Defaults to 11.
	MinPower          *int     `json:"minPower,omitempty"`          // Sets min power (dBm) of 2.4Ghz band. Can be integer between 2 and 30. Defaults to 5.
	Rxsop             *int     `json:"rxsop,omitempty"`             // The RX-SOP level controls the sensitivity of the radio. It is strongly recommended to use RX-SOP only after consulting a wireless expert. RX-SOP can be configured in the range of -65 to -95 (dBm). A value of null will reset this to the default.
	ValidAutoChannels *[]int   `json:"validAutoChannels,omitempty"` // Sets valid auto channels for 2.4Ghz band. Can be one of '1', '6' or '11'. Defaults to [1, 6, 11].
}
type ResponseWirelessCreateNetworkWirelessRfProfile struct {
	ApBandSettings         *ResponseWirelessCreateNetworkWirelessRfProfileApBandSettings     `json:"apBandSettings,omitempty"`         // Settings that will be enabled if selectionType is set to 'ap'.
	BandSelectionType      string                                                            `json:"bandSelectionType,omitempty"`      // Band selection can be set to either 'ssid' or 'ap'. This param is required on creation.
	ClientBalancingEnabled *bool                                                             `json:"clientBalancingEnabled,omitempty"` // Steers client to best available access point. Can be either true or false. Defaults to true.
	FiveGhzSettings        *ResponseWirelessCreateNetworkWirelessRfProfileFiveGhzSettings    `json:"fiveGhzSettings,omitempty"`        // Settings related to 5Ghz band
	ID                     string                                                            `json:"id,omitempty"`                     // The name of the new profile. Must be unique.
	MinBitrateType         string                                                            `json:"minBitrateType,omitempty"`         // Minimum bitrate can be set to either 'band' or 'ssid'. Defaults to band.
	Name                   string                                                            `json:"name,omitempty"`                   // The name of the new profile. Must be unique. This param is required on creation.
	NetworkID              string                                                            `json:"networkId,omitempty"`              // The network ID of the RF Profile
	PerSSIDSettings        *ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings    `json:"perSsidSettings,omitempty"`        // Per-SSID radio settings by number.
	SixGhzSettings         *ResponseWirelessCreateNetworkWirelessRfProfileSixGhzSettings     `json:"sixGhzSettings,omitempty"`         // Settings related to 6Ghz band. Only applicable to networks with 6Ghz capable APs
	Transmission           *ResponseWirelessCreateNetworkWirelessRfProfileTransmission       `json:"transmission,omitempty"`           // Settings related to radio transmission.
	TwoFourGhzSettings     *ResponseWirelessCreateNetworkWirelessRfProfileTwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"`     // Settings related to 2.4Ghz band
}
type ResponseWirelessCreateNetworkWirelessRfProfileApBandSettings struct {
	BandOperationMode   string                                                             `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'. Defaults to dual.
	BandSteeringEnabled *bool                                                              `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band. Can be either true or false. Defaults to true.
	Bands               *ResponseWirelessCreateNetworkWirelessRfProfileApBandSettingsBands `json:"bands,omitempty"`               // Settings related to all bands
}
type ResponseWirelessCreateNetworkWirelessRfProfileApBandSettingsBands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4","5", "6", "disabled"
}
type ResponseWirelessCreateNetworkWirelessRfProfileFiveGhzSettings struct {
	ChannelWidth      string `json:"channelWidth,omitempty"`      // Sets channel width (MHz) for 5Ghz band. Can be one of 'auto', '20', '40' or '80'. Defaults to auto.
	MaxPower          *int   `json:"maxPower,omitempty"`          // Sets max power (dBm) of 5Ghz band. Can be integer between 2 and 30. Defaults to 30.
	MinBitrate        *int   `json:"minBitrate,omitempty"`        // Sets min bitrate (Mbps) of 5Ghz band. Can be one of '6', '9', '12', '18', '24', '36', '48' or '54'. Defaults to 12.
	MinPower          *int   `json:"minPower,omitempty"`          // Sets min power (dBm) of 5Ghz band. Can be integer between 2 and 30. Defaults to 8.
	Rxsop             *int   `json:"rxsop,omitempty"`             // The RX-SOP level controls the sensitivity of the radio. It is strongly recommended to use RX-SOP only after consulting a wireless expert. RX-SOP can be configured in the range of -65 to -95 (dBm). A value of null will reset this to the default.
	ValidAutoChannels *[]int `json:"validAutoChannels,omitempty"` // Sets valid auto channels for 5Ghz band. Can be one of '36', '40', '44', '48', '52', '56', '60', '64', '100', '104', '108', '112', '116', '120', '124', '128', '132', '136', '140', '144', '149', '153', '157', '161' or '165'.Defaults to [36, 40, 44, 48, 52, 56, 60, 64, 100, 104, 108, 112, 116, 120, 124, 128, 132, 136, 140, 144, 149, 153, 157, 161, 165].
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings struct {
	Status0  *ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings0  `json:"0,omitempty"`  // Settings for SSID 0
	Status1  *ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings1  `json:"1,omitempty"`  // Settings for SSID 1
	Status10 *ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings10 `json:"10,omitempty"` // Settings for SSID 10
	Status11 *ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings11 `json:"11,omitempty"` // Settings for SSID 11
	Status12 *ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings12 `json:"12,omitempty"` // Settings for SSID 12
	Status13 *ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings13 `json:"13,omitempty"` // Settings for SSID 13
	Status14 *ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings14 `json:"14,omitempty"` // Settings for SSID 14
	Status2  *ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings2  `json:"2,omitempty"`  // Settings for SSID 2
	Status3  *ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings3  `json:"3,omitempty"`  // Settings for SSID 3
	Status4  *ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings4  `json:"4,omitempty"`  // Settings for SSID 4
	Status5  *ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings5  `json:"5,omitempty"`  // Settings for SSID 5
	Status6  *ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings6  `json:"6,omitempty"`  // Settings for SSID 6
	Status7  *ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings7  `json:"7,omitempty"`  // Settings for SSID 7
	Status8  *ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings8  `json:"8,omitempty"`  // Settings for SSID 8
	Status9  *ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings9  `json:"9,omitempty"`  // Settings for SSID 9
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings0 struct {
	BandOperationMode   string                                                               `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                                `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings0Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                                 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                               `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings0Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4","5", "6", "disabled"
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings1 struct {
	BandOperationMode   string                                                               `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                                `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings1Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                                 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                               `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings1Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4","5", "6", "disabled"
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings10 struct {
	BandOperationMode   string                                                                `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                                 `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings10Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                                  `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                                `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings10Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings11 struct {
	BandOperationMode   string                                                                `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                                 `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings11Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                                  `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                                `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings11Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4","5", "6", "disabled"
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings12 struct {
	BandOperationMode   string                                                                `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                                 `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings12Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                                  `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                                `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings12Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4","5", "6", "disabled"
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings13 struct {
	BandOperationMode   string                                                                `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                                 `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings13Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                                  `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                                `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings13Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4","5", "6", "disabled"
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings14 struct {
	BandOperationMode   string                                                                `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                                 `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings14Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                                  `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                                `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings14Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings2 struct {
	BandOperationMode   string                                                               `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                                `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings2Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                                 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                               `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings2Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings3 struct {
	BandOperationMode   string                                                               `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                                `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings3Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                                 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                               `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings3Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings4 struct {
	BandOperationMode   string                                                               `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                                `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings4Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                                 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                               `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings4Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings5 struct {
	BandOperationMode   string                                                               `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                                `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings5Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                                 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                               `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings5Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings6 struct {
	BandOperationMode   string                                                               `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                                `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings6Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                                 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                               `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings6Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings7 struct {
	BandOperationMode   string                                                               `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                                `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings7Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                                 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                               `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings7Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings8 struct {
	BandOperationMode   string                                                               `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                                `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings8Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                                 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                               `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings8Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings9 struct {
	BandOperationMode   string                                                               `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                                `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings9Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                                 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                               `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings9Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessCreateNetworkWirelessRfProfileSixGhzSettings struct {
	ChannelWidth      string `json:"channelWidth,omitempty"`      // Sets channel width (MHz) for 6Ghz band. Can be one of '0', '20', '40', '80' or '160'. Defaults to auto.
	MaxPower          *int   `json:"maxPower,omitempty"`          // Sets max power (dBm) of 6Ghz band. Can be integer between 2 and 30. Defaults to 30.
	MinBitrate        *int   `json:"minBitrate,omitempty"`        // Sets min bitrate (Mbps) of 6Ghz band. Can be one of '6', '9', '12', '18', '24', '36', '48' or '54'. Defaults to 12.
	MinPower          *int   `json:"minPower,omitempty"`          // Sets min power (dBm) of 6Ghz band. Can be integer between 2 and 30. Defaults to 8.
	Rxsop             *int   `json:"rxsop,omitempty"`             // The RX-SOP level controls the sensitivity of the radio. It is strongly recommended to use RX-SOP only after consulting a wireless expert. RX-SOP can be configured in the range of -65 to -95 (dBm). A value of null will reset this to the default.
	ValidAutoChannels *[]int `json:"validAutoChannels,omitempty"` // Sets valid auto channels for 6Ghz band. Can be one of '1', '5', '9', '13', '17', '21', '25', '29', '33', '37', '41', '45', '49', '53', '57', '61', '65', '69', '73', '77', '81', '85', '89', '93', '97', '101', '105', '109', '113', '117', '121', '125', '129', '133', '137', '141', '145', '149', '153', '157', '161', '165', '169', '173', '177', '181', '185', '189', '193', '197', '201', '205', '209', '213', '217', '221', '225', '229' or '233'. Defaults to auto.
}
type ResponseWirelessCreateNetworkWirelessRfProfileTransmission struct {
	Enabled *bool `json:"enabled,omitempty"` // Toggle for radio transmission. When false, radios will not transmit at all.
}
type ResponseWirelessCreateNetworkWirelessRfProfileTwoFourGhzSettings struct {
	AxEnabled         *bool    `json:"axEnabled,omitempty"`         // Determines whether ax radio on 2.4Ghz band is on or off. Can be either true or false. If false, we highly recommend disabling band steering. Defaults to true.
	MaxPower          *int     `json:"maxPower,omitempty"`          // Sets max power (dBm) of 2.4Ghz band. Can be integer between 2 and 30. Defaults to 30.
	MinBitrate        *float64 `json:"minBitrate,omitempty"`        // Sets min bitrate (Mbps) of 2.4Ghz band. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'. Defaults to 11.
	MinPower          *int     `json:"minPower,omitempty"`          // Sets min power (dBm) of 2.4Ghz band. Can be integer between 2 and 30. Defaults to 5.
	Rxsop             *int     `json:"rxsop,omitempty"`             // The RX-SOP level controls the sensitivity of the radio. It is strongly recommended to use RX-SOP only after consulting a wireless expert. RX-SOP can be configured in the range of -65 to -95 (dBm). A value of null will reset this to the default.
	ValidAutoChannels *[]int   `json:"validAutoChannels,omitempty"` // Sets valid auto channels for 2.4Ghz band. Can be one of '1', '6' or '11'. Defaults to [1, 6, 11].
}
type ResponseWirelessGetNetworkWirelessRfProfile struct {
	ApBandSettings         *ResponseWirelessGetNetworkWirelessRfProfileApBandSettings     `json:"apBandSettings,omitempty"`         // Settings that will be enabled if selectionType is set to 'ap'.
	BandSelectionType      string                                                         `json:"bandSelectionType,omitempty"`      // Band selection can be set to either 'ssid' or 'ap'. This param is required on creation.
	ClientBalancingEnabled *bool                                                          `json:"clientBalancingEnabled,omitempty"` // Steers client to best available access point. Can be either true or false. Defaults to true.
	FiveGhzSettings        *ResponseWirelessGetNetworkWirelessRfProfileFiveGhzSettings    `json:"fiveGhzSettings,omitempty"`        // Settings related to 5Ghz band
	ID                     string                                                         `json:"id,omitempty"`                     // The name of the new profile. Must be unique.
	MinBitrateType         string                                                         `json:"minBitrateType,omitempty"`         // Minimum bitrate can be set to either 'band' or 'ssid'. Defaults to band.
	Name                   string                                                         `json:"name,omitempty"`                   // The name of the new profile. Must be unique. This param is required on creation.
	NetworkID              string                                                         `json:"networkId,omitempty"`              // The network ID of the RF Profile
	PerSSIDSettings        *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings    `json:"perSsidSettings,omitempty"`        // Per-SSID radio settings by number.
	SixGhzSettings         *ResponseWirelessGetNetworkWirelessRfProfileSixGhzSettings     `json:"sixGhzSettings,omitempty"`         // Settings related to 6Ghz band. Only applicable to networks with 6Ghz capable APs
	Transmission           *ResponseWirelessGetNetworkWirelessRfProfileTransmission       `json:"transmission,omitempty"`           // Settings related to radio transmission.
	TwoFourGhzSettings     *ResponseWirelessGetNetworkWirelessRfProfileTwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"`     // Settings related to 2.4Ghz band
}
type ResponseWirelessGetNetworkWirelessRfProfileApBandSettings struct {
	BandOperationMode   string                                                          `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'. Defaults to dual.
	BandSteeringEnabled *bool                                                           `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band. Can be either true or false. Defaults to true.
	Bands               *ResponseWirelessGetNetworkWirelessRfProfileApBandSettingsBands `json:"bands,omitempty"`               // Settings related to all bands
}
type ResponseWirelessGetNetworkWirelessRfProfileApBandSettingsBands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessGetNetworkWirelessRfProfileFiveGhzSettings struct {
	ChannelWidth      string `json:"channelWidth,omitempty"`      // Sets channel width (MHz) for 5Ghz band. Can be one of 'auto', '20', '40' or '80'. Defaults to auto.
	MaxPower          *int   `json:"maxPower,omitempty"`          // Sets max power (dBm) of 5Ghz band. Can be integer between 2 and 30. Defaults to 30.
	MinBitrate        *int   `json:"minBitrate,omitempty"`        // Sets min bitrate (Mbps) of 5Ghz band. Can be one of '6', '9', '12', '18', '24', '36', '48' or '54'. Defaults to 12.
	MinPower          *int   `json:"minPower,omitempty"`          // Sets min power (dBm) of 5Ghz band. Can be integer between 2 and 30. Defaults to 8.
	Rxsop             *int   `json:"rxsop,omitempty"`             // The RX-SOP level controls the sensitivity of the radio. It is strongly recommended to use RX-SOP only after consulting a wireless expert. RX-SOP can be configured in the range of -65 to -95 (dBm). A value of null will reset this to the default.
	ValidAutoChannels *[]int `json:"validAutoChannels,omitempty"` // Sets valid auto channels for 5Ghz band. Can be one of '36', '40', '44', '48', '52', '56', '60', '64', '100', '104', '108', '112', '116', '120', '124', '128', '132', '136', '140', '144', '149', '153', '157', '161' or '165'.Defaults to [36, 40, 44, 48, 52, 56, 60, 64, 100, 104, 108, 112, 116, 120, 124, 128, 132, 136, 140, 144, 149, 153, 157, 161, 165].
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings struct {
	Status0  *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings0  `json:"0,omitempty"`  // Settings for SSID 0
	Status1  *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings1  `json:"1,omitempty"`  // Settings for SSID 1
	Status10 *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings10 `json:"10,omitempty"` // Settings for SSID 10
	Status11 *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings11 `json:"11,omitempty"` // Settings for SSID 11
	Status12 *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings12 `json:"12,omitempty"` // Settings for SSID 12
	Status13 *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings13 `json:"13,omitempty"` // Settings for SSID 13
	Status14 *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings14 `json:"14,omitempty"` // Settings for SSID 14
	Status2  *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings2  `json:"2,omitempty"`  // Settings for SSID 2
	Status3  *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings3  `json:"3,omitempty"`  // Settings for SSID 3
	Status4  *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings4  `json:"4,omitempty"`  // Settings for SSID 4
	Status5  *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings5  `json:"5,omitempty"`  // Settings for SSID 5
	Status6  *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings6  `json:"6,omitempty"`  // Settings for SSID 6
	Status7  *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings7  `json:"7,omitempty"`  // Settings for SSID 7
	Status8  *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings8  `json:"8,omitempty"`  // Settings for SSID 8
	Status9  *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings9  `json:"9,omitempty"`  // Settings for SSID 9
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings0 struct {
	BandOperationMode   string                                                            `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                             `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings0Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                              `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                            `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings0Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings1 struct {
	BandOperationMode   string                                                            `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                             `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings1Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                              `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                            `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings1Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings10 struct {
	BandOperationMode   string                                                             `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                              `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings10Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                               `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                             `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings10Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings11 struct {
	BandOperationMode   string                                                             `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                              `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings11Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                               `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                             `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings11Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings12 struct {
	BandOperationMode   string                                                             `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                              `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings12Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                               `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                             `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings12Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings13 struct {
	BandOperationMode   string                                                             `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                              `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings13Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                               `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                             `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings13Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings14 struct {
	BandOperationMode   string                                                             `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                              `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings14Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                               `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                             `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings14Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings2 struct {
	BandOperationMode   string                                                            `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                             `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings2Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                              `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                            `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings2Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings3 struct {
	BandOperationMode   string                                                            `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                             `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings3Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                              `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                            `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings3Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings4 struct {
	BandOperationMode   string                                                            `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                             `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings4Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                              `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                            `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings4Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings5 struct {
	BandOperationMode   string                                                            `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                             `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings5Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                              `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                            `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings5Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings6 struct {
	BandOperationMode   string                                                            `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                             `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings6Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                              `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                            `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings6Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings7 struct {
	BandOperationMode   string                                                            `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                             `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings7Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                              `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                            `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings7Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings8 struct {
	BandOperationMode   string                                                            `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                             `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings8Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                              `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                            `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings8Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings9 struct {
	BandOperationMode   string                                                            `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                             `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings9Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                              `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                            `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings9Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessGetNetworkWirelessRfProfileSixGhzSettings struct {
	ChannelWidth      string `json:"channelWidth,omitempty"`      // Sets channel width (MHz) for 6Ghz band. Can be one of '0', '20', '40', '80' or '160'. Defaults to auto.
	MaxPower          *int   `json:"maxPower,omitempty"`          // Sets max power (dBm) of 6Ghz band. Can be integer between 2 and 30. Defaults to 30.
	MinBitrate        *int   `json:"minBitrate,omitempty"`        // Sets min bitrate (Mbps) of 6Ghz band. Can be one of '6', '9', '12', '18', '24', '36', '48' or '54'. Defaults to 12.
	MinPower          *int   `json:"minPower,omitempty"`          // Sets min power (dBm) of 6Ghz band. Can be integer between 2 and 30. Defaults to 8.
	Rxsop             *int   `json:"rxsop,omitempty"`             // The RX-SOP level controls the sensitivity of the radio. It is strongly recommended to use RX-SOP only after consulting a wireless expert. RX-SOP can be configured in the range of -65 to -95 (dBm). A value of null will reset this to the default.
	ValidAutoChannels *[]int `json:"validAutoChannels,omitempty"` // Sets valid auto channels for 6Ghz band. Can be one of '1', '5', '9', '13', '17', '21', '25', '29', '33', '37', '41', '45', '49', '53', '57', '61', '65', '69', '73', '77', '81', '85', '89', '93', '97', '101', '105', '109', '113', '117', '121', '125', '129', '133', '137', '141', '145', '149', '153', '157', '161', '165', '169', '173', '177', '181', '185', '189', '193', '197', '201', '205', '209', '213', '217', '221', '225', '229' or '233'. Defaults to auto.
}
type ResponseWirelessGetNetworkWirelessRfProfileTransmission struct {
	Enabled *bool `json:"enabled,omitempty"` // Toggle for radio transmission. When false, radios will not transmit at all.
}
type ResponseWirelessGetNetworkWirelessRfProfileTwoFourGhzSettings struct {
	AxEnabled         *bool    `json:"axEnabled,omitempty"`         // Determines whether ax radio on 2.4Ghz band is on or off. Can be either true or false. If false, we highly recommend disabling band steering. Defaults to true.
	MaxPower          *int     `json:"maxPower,omitempty"`          // Sets max power (dBm) of 2.4Ghz band. Can be integer between 2 and 30. Defaults to 30.
	MinBitrate        *float64 `json:"minBitrate,omitempty"`        // Sets min bitrate (Mbps) of 2.4Ghz band. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'. Defaults to 11.
	MinPower          *int     `json:"minPower,omitempty"`          // Sets min power (dBm) of 2.4Ghz band. Can be integer between 2 and 30. Defaults to 5.
	Rxsop             *int     `json:"rxsop,omitempty"`             // The RX-SOP level controls the sensitivity of the radio. It is strongly recommended to use RX-SOP only after consulting a wireless expert. RX-SOP can be configured in the range of -65 to -95 (dBm). A value of null will reset this to the default.
	ValidAutoChannels *[]int   `json:"validAutoChannels,omitempty"` // Sets valid auto channels for 2.4Ghz band. Can be one of '1', '6' or '11'. Defaults to [1, 6, 11].
}
type ResponseWirelessUpdateNetworkWirelessRfProfile struct {
	ApBandSettings         *ResponseWirelessUpdateNetworkWirelessRfProfileApBandSettings     `json:"apBandSettings,omitempty"`         // Settings that will be enabled if selectionType is set to 'ap'.
	BandSelectionType      string                                                            `json:"bandSelectionType,omitempty"`      // Band selection can be set to either 'ssid' or 'ap'. This param is required on creation.
	ClientBalancingEnabled *bool                                                             `json:"clientBalancingEnabled,omitempty"` // Steers client to best available access point. Can be either true or false. Defaults to true.
	FiveGhzSettings        *ResponseWirelessUpdateNetworkWirelessRfProfileFiveGhzSettings    `json:"fiveGhzSettings,omitempty"`        // Settings related to 5Ghz band
	ID                     string                                                            `json:"id,omitempty"`                     // The name of the new profile. Must be unique.
	MinBitrateType         string                                                            `json:"minBitrateType,omitempty"`         // Minimum bitrate can be set to either 'band' or 'ssid'. Defaults to band.
	Name                   string                                                            `json:"name,omitempty"`                   // The name of the new profile. Must be unique. This param is required on creation.
	NetworkID              string                                                            `json:"networkId,omitempty"`              // The network ID of the RF Profile
	PerSSIDSettings        *ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings    `json:"perSsidSettings,omitempty"`        // Per-SSID radio settings by number.
	SixGhzSettings         *ResponseWirelessUpdateNetworkWirelessRfProfileSixGhzSettings     `json:"sixGhzSettings,omitempty"`         // Settings related to 6Ghz band. Only applicable to networks with 6Ghz capable APs
	Transmission           *ResponseWirelessUpdateNetworkWirelessRfProfileTransmission       `json:"transmission,omitempty"`           // Settings related to radio transmission.
	TwoFourGhzSettings     *ResponseWirelessUpdateNetworkWirelessRfProfileTwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"`     // Settings related to 2.4Ghz band
}
type ResponseWirelessUpdateNetworkWirelessRfProfileApBandSettings struct {
	BandOperationMode   string                                                             `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'. Defaults to dual.
	BandSteeringEnabled *bool                                                              `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band. Can be either true or false. Defaults to true.
	Bands               *ResponseWirelessUpdateNetworkWirelessRfProfileApBandSettingsBands `json:"bands,omitempty"`               // Settings related to all bands
}
type ResponseWirelessUpdateNetworkWirelessRfProfileApBandSettingsBands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessUpdateNetworkWirelessRfProfileFiveGhzSettings struct {
	ChannelWidth      string `json:"channelWidth,omitempty"`      // Sets channel width (MHz) for 5Ghz band. Can be one of 'auto', '20', '40' or '80'. Defaults to auto.
	MaxPower          *int   `json:"maxPower,omitempty"`          // Sets max power (dBm) of 5Ghz band. Can be integer between 2 and 30. Defaults to 30.
	MinBitrate        *int   `json:"minBitrate,omitempty"`        // Sets min bitrate (Mbps) of 5Ghz band. Can be one of '6', '9', '12', '18', '24', '36', '48' or '54'. Defaults to 12.
	MinPower          *int   `json:"minPower,omitempty"`          // Sets min power (dBm) of 5Ghz band. Can be integer between 2 and 30. Defaults to 8.
	Rxsop             *int   `json:"rxsop,omitempty"`             // The RX-SOP level controls the sensitivity of the radio. It is strongly recommended to use RX-SOP only after consulting a wireless expert. RX-SOP can be configured in the range of -65 to -95 (dBm). A value of null will reset this to the default.
	ValidAutoChannels *[]int `json:"validAutoChannels,omitempty"` // Sets valid auto channels for 5Ghz band. Can be one of '36', '40', '44', '48', '52', '56', '60', '64', '100', '104', '108', '112', '116', '120', '124', '128', '132', '136', '140', '144', '149', '153', '157', '161' or '165'.Defaults to [36, 40, 44, 48, 52, 56, 60, 64, 100, 104, 108, 112, 116, 120, 124, 128, 132, 136, 140, 144, 149, 153, 157, 161, 165].
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings struct {
	Status0  *ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings0  `json:"0,omitempty"`  // Settings for SSID 0
	Status1  *ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings1  `json:"1,omitempty"`  // Settings for SSID 1
	Status10 *ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings10 `json:"10,omitempty"` // Settings for SSID 10
	Status11 *ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings11 `json:"11,omitempty"` // Settings for SSID 11
	Status12 *ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings12 `json:"12,omitempty"` // Settings for SSID 12
	Status13 *ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings13 `json:"13,omitempty"` // Settings for SSID 13
	Status14 *ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings14 `json:"14,omitempty"` // Settings for SSID 14
	Status2  *ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings2  `json:"2,omitempty"`  // Settings for SSID 2
	Status3  *ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings3  `json:"3,omitempty"`  // Settings for SSID 3
	Status4  *ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings4  `json:"4,omitempty"`  // Settings for SSID 4
	Status5  *ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings5  `json:"5,omitempty"`  // Settings for SSID 5
	Status6  *ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings6  `json:"6,omitempty"`  // Settings for SSID 6
	Status7  *ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings7  `json:"7,omitempty"`  // Settings for SSID 7
	Status8  *ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings8  `json:"8,omitempty"`  // Settings for SSID 8
	Status9  *ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings9  `json:"9,omitempty"`  // Settings for SSID 9
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings0 struct {
	BandOperationMode   string                                                               `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                                `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings0Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                                 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                               `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings0Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings1 struct {
	BandOperationMode   string                                                               `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                                `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings1Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                                 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                               `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings1Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings10 struct {
	BandOperationMode   string                                                                `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                                 `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings10Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                                  `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                                `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings10Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings11 struct {
	BandOperationMode   string                                                                `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                                 `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings11Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                                  `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                                `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings11Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings12 struct {
	BandOperationMode   string                                                                `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                                 `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings12Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                                  `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                                `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings12Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings13 struct {
	BandOperationMode   string                                                                `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                                 `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings13Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                                  `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                                `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings13Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings14 struct {
	BandOperationMode   string                                                                `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                                 `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings14Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                                  `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                                `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings14Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings2 struct {
	BandOperationMode   string                                                               `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                                `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings2Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                                 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                               `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings2Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings3 struct {
	BandOperationMode   string                                                               `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                                `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings3Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                                 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                               `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings3Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings4 struct {
	BandOperationMode   string                                                               `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                                `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings4Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                                 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                               `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings4Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings5 struct {
	BandOperationMode   string                                                               `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                                `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings5Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                                 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                               `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings5Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings6 struct {
	BandOperationMode   string                                                               `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                                `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings6Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                                 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                               `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings6Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings7 struct {
	BandOperationMode   string                                                               `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                                `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings7Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                                 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                               `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings7Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings8 struct {
	BandOperationMode   string                                                               `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                                `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings8Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                                 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                               `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings8Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings9 struct {
	BandOperationMode   string                                                               `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                                `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings9Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *int                                                                 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string                                                               `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings9Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type ResponseWirelessUpdateNetworkWirelessRfProfileSixGhzSettings struct {
	ChannelWidth      string `json:"channelWidth,omitempty"`      // Sets channel width (MHz) for 6Ghz band. Can be one of '0', '20', '40', '80' or '160'. Defaults to auto.
	MaxPower          *int   `json:"maxPower,omitempty"`          // Sets max power (dBm) of 6Ghz band. Can be integer between 2 and 30. Defaults to 30.
	MinBitrate        *int   `json:"minBitrate,omitempty"`        // Sets min bitrate (Mbps) of 6Ghz band. Can be one of '6', '9', '12', '18', '24', '36', '48' or '54'. Defaults to 12.
	MinPower          *int   `json:"minPower,omitempty"`          // Sets min power (dBm) of 6Ghz band. Can be integer between 2 and 30. Defaults to 8.
	Rxsop             *int   `json:"rxsop,omitempty"`             // The RX-SOP level controls the sensitivity of the radio. It is strongly recommended to use RX-SOP only after consulting a wireless expert. RX-SOP can be configured in the range of -65 to -95 (dBm). A value of null will reset this to the default.
	ValidAutoChannels *[]int `json:"validAutoChannels,omitempty"` // Sets valid auto channels for 6Ghz band. Can be one of '1', '5', '9', '13', '17', '21', '25', '29', '33', '37', '41', '45', '49', '53', '57', '61', '65', '69', '73', '77', '81', '85', '89', '93', '97', '101', '105', '109', '113', '117', '121', '125', '129', '133', '137', '141', '145', '149', '153', '157', '161', '165', '169', '173', '177', '181', '185', '189', '193', '197', '201', '205', '209', '213', '217', '221', '225', '229' or '233'. Defaults to auto.
}
type ResponseWirelessUpdateNetworkWirelessRfProfileTransmission struct {
	Enabled *bool `json:"enabled,omitempty"` // Toggle for radio transmission. When false, radios will not transmit at all.
}
type ResponseWirelessUpdateNetworkWirelessRfProfileTwoFourGhzSettings struct {
	AxEnabled         *bool    `json:"axEnabled,omitempty"`         // Determines whether ax radio on 2.4Ghz band is on or off. Can be either true or false. If false, we highly recommend disabling band steering. Defaults to true.
	MaxPower          *int     `json:"maxPower,omitempty"`          // Sets max power (dBm) of 2.4Ghz band. Can be integer between 2 and 30. Defaults to 30.
	MinBitrate        *float64 `json:"minBitrate,omitempty"`        // Sets min bitrate (Mbps) of 2.4Ghz band. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'. Defaults to 11.
	MinPower          *int     `json:"minPower,omitempty"`          // Sets min power (dBm) of 2.4Ghz band. Can be integer between 2 and 30. Defaults to 5.
	Rxsop             *int     `json:"rxsop,omitempty"`             // The RX-SOP level controls the sensitivity of the radio. It is strongly recommended to use RX-SOP only after consulting a wireless expert. RX-SOP can be configured in the range of -65 to -95 (dBm). A value of null will reset this to the default.
	ValidAutoChannels *[]int   `json:"validAutoChannels,omitempty"` // Sets valid auto channels for 2.4Ghz band. Can be one of '1', '6' or '11'. Defaults to [1, 6, 11].
}
type ResponseWirelessGetNetworkWirelessSettings struct {
	IPv6BridgeEnabled        *bool                                                       `json:"ipv6BridgeEnabled,omitempty"`        // Toggle for enabling or disabling IPv6 bridging in a network (Note: if enabled, SSIDs must also be configured to use bridge mode)
	LedLightsOn              *bool                                                       `json:"ledLightsOn,omitempty"`              // Toggle for enabling or disabling LED lights on all APs in the network (making them run dark)
	LocationAnalyticsEnabled *bool                                                       `json:"locationAnalyticsEnabled,omitempty"` // Toggle for enabling or disabling location analytics for your network
	MeshingEnabled           *bool                                                       `json:"meshingEnabled,omitempty"`           // Toggle for enabling or disabling meshing in a network
	NamedVLANs               *ResponseWirelessGetNetworkWirelessSettingsNamedVLANs       `json:"namedVlans,omitempty"`               // Named VLAN settings for wireless networks.
	RegulatoryDomain         *ResponseWirelessGetNetworkWirelessSettingsRegulatoryDomain `json:"regulatoryDomain,omitempty"`         // Regulatory domain information for this network.
	Upgradestrategy          string                                                      `json:"upgradeStrategy,omitempty"`          // The upgrade strategy to apply to the network. Must be one of 'minimizeUpgradeTime' or 'minimizeClientDowntime'. Requires firmware version MR 26.8 or higher'
}
type ResponseWirelessGetNetworkWirelessSettingsNamedVLANs struct {
	PoolDhcpMonitoring *ResponseWirelessGetNetworkWirelessSettingsNamedVLANsPoolDhcpMonitoring `json:"poolDhcpMonitoring,omitempty"` // Named VLAN Pool DHCP Monitoring settings.
}
type ResponseWirelessGetNetworkWirelessSettingsNamedVLANsPoolDhcpMonitoring struct {
	Duration *int  `json:"duration,omitempty"` // The duration in minutes that devices will refrain from using dirty VLANs before adding them back to the pool.
	Enabled  *bool `json:"enabled,omitempty"`  // Whether or not devices using named VLAN pools should remove dirty VLANs from the pool, thereby preventing clients from being assigned to VLANs where they would be unable to obtain an IP address via DHCP
}
type ResponseWirelessGetNetworkWirelessSettingsRegulatoryDomain struct {
	CountryCode string `json:"countryCode,omitempty"` // The country code of the regulatory domain.
	Name        string `json:"name,omitempty"`        // The name of the regulatory domain for this network.
	Permits6E   *bool  `json:"permits6e,omitempty"`   // Whether or not the regulatory domain for this network permits Wifi 6E.
}
type ResponseWirelessUpdateNetworkWirelessSettings struct {
	IPv6BridgeEnabled        *bool                                                          `json:"ipv6BridgeEnabled,omitempty"`        // Toggle for enabling or disabling IPv6 bridging in a network (Note: if enabled, SSIDs must also be configured to use bridge mode)
	LedLightsOn              *bool                                                          `json:"ledLightsOn,omitempty"`              // Toggle for enabling or disabling LED lights on all APs in the network (making them run dark)
	LocationAnalyticsEnabled *bool                                                          `json:"locationAnalyticsEnabled,omitempty"` // Toggle for enabling or disabling location analytics for your network
	MeshingEnabled           *bool                                                          `json:"meshingEnabled,omitempty"`           // Toggle for enabling or disabling meshing in a network
	NamedVLANs               *ResponseWirelessUpdateNetworkWirelessSettingsNamedVLANs       `json:"namedVlans,omitempty"`               // Named VLAN settings for wireless networks.
	RegulatoryDomain         *ResponseWirelessUpdateNetworkWirelessSettingsRegulatoryDomain `json:"regulatoryDomain,omitempty"`         // Regulatory domain information for this network.
	Upgradestrategy          string                                                         `json:"upgradeStrategy,omitempty"`          // The upgrade strategy to apply to the network. Must be one of 'minimizeUpgradeTime' or 'minimizeClientDowntime'. Requires firmware version MR 26.8 or higher'
}
type ResponseWirelessUpdateNetworkWirelessSettingsNamedVLANs struct {
	PoolDhcpMonitoring *ResponseWirelessUpdateNetworkWirelessSettingsNamedVLANsPoolDhcpMonitoring `json:"poolDhcpMonitoring,omitempty"` // Named VLAN Pool DHCP Monitoring settings.
}
type ResponseWirelessUpdateNetworkWirelessSettingsNamedVLANsPoolDhcpMonitoring struct {
	Duration *int  `json:"duration,omitempty"` // The duration in minutes that devices will refrain from using dirty VLANs before adding them back to the pool.
	Enabled  *bool `json:"enabled,omitempty"`  // Whether or not devices using named VLAN pools should remove dirty VLANs from the pool, thereby preventing clients from being assigned to VLANs where they would be unable to obtain an IP address via DHCP
}
type ResponseWirelessUpdateNetworkWirelessSettingsRegulatoryDomain struct {
	CountryCode string `json:"countryCode,omitempty"` // The country code of the regulatory domain.
	Name        string `json:"name,omitempty"`        // The name of the regulatory domain for this network.
	Permits6E   *bool  `json:"permits6e,omitempty"`   // Whether or not the regulatory domain for this network permits Wifi 6E.
}
type ResponseWirelessGetNetworkWirelessSignalQualityHistory []ResponseItemWirelessGetNetworkWirelessSignalQualityHistory // Array of ResponseWirelessGetNetworkWirelessSignalQualityHistory
type ResponseItemWirelessGetNetworkWirelessSignalQualityHistory struct {
	EndTs   string `json:"endTs,omitempty"`   // The end time of the query range
	Rssi    *int   `json:"rssi,omitempty"`    // Received signal strength indicator
	Snr     *int   `json:"snr,omitempty"`     // Signal to noise ratio
	StartTs string `json:"startTs,omitempty"` // The start time of the query range
}
type ResponseWirelessGetNetworkWirelessSSIDs []ResponseItemWirelessGetNetworkWirelessSSIDs // Array of ResponseWirelessGetNetworkWirelessSsids
type ResponseItemWirelessGetNetworkWirelessSSIDs struct {
	AdminSplashURL                  string                                                                `json:"adminSplashUrl,omitempty"`                  // URL for the admin splash page
	AuthMode                        string                                                                `json:"authMode,omitempty"`                        // The association control method for the SSID
	AvailabilityTags                []string                                                              `json:"availabilityTags,omitempty"`                // List of tags for this SSID. If availableOnAllAps is false, then the SSID is only broadcast by APs with tags matching any of the tags in this list
	AvailableOnAllAps               *bool                                                                 `json:"availableOnAllAps,omitempty"`               // Whether all APs broadcast the SSID or if it's restricted to APs matching any availability tags
	BandSelection                   string                                                                `json:"bandSelection,omitempty"`                   // The client-serving radio frequencies of this SSID in the default indoor RF profile
	Enabled                         *bool                                                                 `json:"enabled,omitempty"`                         // Whether or not the SSID is enabled
	EncryptionMode                  string                                                                `json:"encryptionMode,omitempty"`                  // The psk encryption mode for the SSID
	IPAssignmentMode                string                                                                `json:"ipAssignmentMode,omitempty"`                // The client IP assignment mode
	LocalAuth                       *bool                                                                 `json:"localAuth,omitempty"`                       // Extended local auth flag for Enterprise NAC
	MandatoryDhcpEnabled            *bool                                                                 `json:"mandatoryDhcpEnabled,omitempty"`            // Whether clients connecting to this SSID must use the IP address assigned by the DHCP server
	MinBitrate                      *int                                                                  `json:"minBitrate,omitempty"`                      // The minimum bitrate in Mbps of this SSID in the default indoor RF profile
	Name                            string                                                                `json:"name,omitempty"`                            // The name of the SSID
	Number                          *int                                                                  `json:"number,omitempty"`                          // Unique identifier of the SSID
	PerClientBandwidthLimitDown     *int                                                                  `json:"perClientBandwidthLimitDown,omitempty"`     // The download bandwidth limit in Kbps. (0 represents no limit.)
	PerClientBandwidthLimitUp       *int                                                                  `json:"perClientBandwidthLimitUp,omitempty"`       // The upload bandwidth limit in Kbps. (0 represents no limit.)
	PerSSIDBandwidthLimitDown       *int                                                                  `json:"perSsidBandwidthLimitDown,omitempty"`       // The total download bandwidth limit in Kbps (0 represents no limit)
	PerSSIDBandwidthLimitUp         *int                                                                  `json:"perSsidBandwidthLimitUp,omitempty"`         // The total upload bandwidth limit in Kbps (0 represents no limit)
	RadiusAccountingEnabled         *bool                                                                 `json:"radiusAccountingEnabled,omitempty"`         // Whether or not RADIUS accounting is enabled
	RadiusAccountingServers         *[]ResponseItemWirelessGetNetworkWirelessSSIDsRadiusAccountingServers `json:"radiusAccountingServers,omitempty"`         // List of RADIUS accounting 802.1X servers to be used for authentication
	RadiusAttributeForGroupPolicies string                                                                `json:"radiusAttributeForGroupPolicies,omitempty"` // RADIUS attribute used to look up group policies
	RadiusEnabled                   *bool                                                                 `json:"radiusEnabled,omitempty"`                   // Whether RADIUS authentication is enabled
	RadiusFailoverPolicy            string                                                                `json:"radiusFailoverPolicy,omitempty"`            // Policy which determines how authentication requests should be handled in the event that all of the configured RADIUS servers are unreachable
	RadiusLoadBalancingPolicy       string                                                                `json:"radiusLoadBalancingPolicy,omitempty"`       // Policy which determines which RADIUS server will be contacted first in an authentication attempt, and the ordering of any necessary retry attempts
	RadiusServers                   *[]ResponseItemWirelessGetNetworkWirelessSSIDsRadiusServers           `json:"radiusServers,omitempty"`                   // List of RADIUS 802.1X servers to be used for authentication
	SplashPage                      string                                                                `json:"splashPage,omitempty"`                      // The type of splash page for the SSID
	SplashTimeout                   string                                                                `json:"splashTimeout,omitempty"`                   // Splash page timeout
	SSIDAdminAccessible             *bool                                                                 `json:"ssidAdminAccessible,omitempty"`             // SSID Administrator access status
	Visible                         *bool                                                                 `json:"visible,omitempty"`                         // Whether the SSID is advertised or hidden by the AP
	WalledGardenEnabled             *bool                                                                 `json:"walledGardenEnabled,omitempty"`             // Allow users to access a configurable list of IP ranges prior to sign-on
	WalledGardenRanges              []string                                                              `json:"walledGardenRanges,omitempty"`              // Domain names and IP address ranges available in Walled Garden mode
	WpaEncryptionMode               string                                                                `json:"wpaEncryptionMode,omitempty"`               // The types of WPA encryption
}
type ResponseItemWirelessGetNetworkWirelessSSIDsRadiusAccountingServers struct {
	CaCertificate            string `json:"caCertificate,omitempty"`            // Certificate used for authorization for the RADSEC Server
	Host                     string `json:"host,omitempty"`                     // IP address (or FQDN) to which the APs will send RADIUS accounting messages
	OpenRoamingCertificateID *int   `json:"openRoamingCertificateId,omitempty"` // The ID of the Openroaming Certificate attached to radius server
	Port                     *int   `json:"port,omitempty"`                     // Port on the RADIUS server that is listening for accounting messages
}
type ResponseItemWirelessGetNetworkWirelessSSIDsRadiusServers struct {
	CaCertificate            string `json:"caCertificate,omitempty"`            // Certificate used for authorization for the RADSEC Server
	Host                     string `json:"host,omitempty"`                     // IP address (or FQDN) of your RADIUS server
	OpenRoamingCertificateID *int   `json:"openRoamingCertificateId,omitempty"` // The ID of the Openroaming Certificate attached to radius server
	Port                     *int   `json:"port,omitempty"`                     // UDP port the RADIUS server listens on for Access-requests
}
type ResponseWirelessGetNetworkWirelessSSID struct {
	AdminSplashURL                  string                                                           `json:"adminSplashUrl,omitempty"`                  // URL for the admin splash page
	AuthMode                        string                                                           `json:"authMode,omitempty"`                        // The association control method for the SSID
	AvailabilityTags                []string                                                         `json:"availabilityTags,omitempty"`                // List of tags for this SSID. If availableOnAllAps is false, then the SSID is only broadcast by APs with tags matching any of the tags in this list
	AvailableOnAllAps               *bool                                                            `json:"availableOnAllAps,omitempty"`               // Whether all APs broadcast the SSID or if it's restricted to APs matching any availability tags
	BandSelection                   string                                                           `json:"bandSelection,omitempty"`                   // The client-serving radio frequencies of this SSID in the default indoor RF profile
	Enabled                         *bool                                                            `json:"enabled,omitempty"`                         // Whether or not the SSID is enabled
	EncryptionMode                  string                                                           `json:"encryptionMode,omitempty"`                  // The psk encryption mode for the SSID
	IPAssignmentMode                string                                                           `json:"ipAssignmentMode,omitempty"`                // The client IP assignment mode
	LocalAuth                       *bool                                                            `json:"localAuth,omitempty"`                       // Extended local auth flag for Enterprise NAC
	MandatoryDhcpEnabled            *bool                                                            `json:"mandatoryDhcpEnabled,omitempty"`            // Whether clients connecting to this SSID must use the IP address assigned by the DHCP server
	MinBitrate                      *int                                                             `json:"minBitrate,omitempty"`                      // The minimum bitrate in Mbps of this SSID in the default indoor RF profile
	Name                            string                                                           `json:"name,omitempty"`                            // The name of the SSID
	Number                          *int                                                             `json:"number,omitempty"`                          // Unique identifier of the SSID
	PerClientBandwidthLimitDown     *int                                                             `json:"perClientBandwidthLimitDown,omitempty"`     // The download bandwidth limit in Kbps. (0 represents no limit.)
	PerClientBandwidthLimitUp       *int                                                             `json:"perClientBandwidthLimitUp,omitempty"`       // The upload bandwidth limit in Kbps. (0 represents no limit.)
	PerSSIDBandwidthLimitDown       *int                                                             `json:"perSsidBandwidthLimitDown,omitempty"`       // The total download bandwidth limit in Kbps (0 represents no limit)
	PerSSIDBandwidthLimitUp         *int                                                             `json:"perSsidBandwidthLimitUp,omitempty"`         // The total upload bandwidth limit in Kbps (0 represents no limit)
	RadiusAccountingEnabled         *bool                                                            `json:"radiusAccountingEnabled,omitempty"`         // Whether or not RADIUS accounting is enabled
	RadiusAccountingServers         *[]ResponseWirelessGetNetworkWirelessSSIDRadiusAccountingServers `json:"radiusAccountingServers,omitempty"`         // List of RADIUS accounting 802.1X servers to be used for authentication
	RadiusAttributeForGroupPolicies string                                                           `json:"radiusAttributeForGroupPolicies,omitempty"` // RADIUS attribute used to look up group policies
	RadiusEnabled                   *bool                                                            `json:"radiusEnabled,omitempty"`                   // Whether RADIUS authentication is enabled
	RadiusFailoverPolicy            string                                                           `json:"radiusFailoverPolicy,omitempty"`            // Policy which determines how authentication requests should be handled in the event that all of the configured RADIUS servers are unreachable
	RadiusLoadBalancingPolicy       string                                                           `json:"radiusLoadBalancingPolicy,omitempty"`       // Policy which determines which RADIUS server will be contacted first in an authentication attempt, and the ordering of any necessary retry attempts
	RadiusServers                   *[]ResponseWirelessGetNetworkWirelessSSIDRadiusServers           `json:"radiusServers,omitempty"`                   // List of RADIUS 802.1X servers to be used for authentication
	SplashPage                      string                                                           `json:"splashPage,omitempty"`                      // The type of splash page for the SSID
	SplashTimeout                   string                                                           `json:"splashTimeout,omitempty"`                   // Splash page timeout
	SSIDAdminAccessible             *bool                                                            `json:"ssidAdminAccessible,omitempty"`             // SSID Administrator access status
	Visible                         *bool                                                            `json:"visible,omitempty"`                         // Whether the SSID is advertised or hidden by the AP
	WalledGardenEnabled             *bool                                                            `json:"walledGardenEnabled,omitempty"`             // Allow users to access a configurable list of IP ranges prior to sign-on
	WalledGardenRanges              []string                                                         `json:"walledGardenRanges,omitempty"`              // Domain names and IP address ranges available in Walled Garden mode
	WpaEncryptionMode               string                                                           `json:"wpaEncryptionMode,omitempty"`               // The types of WPA encryption
}
type ResponseWirelessGetNetworkWirelessSSIDRadiusAccountingServers struct {
	CaCertificate            string `json:"caCertificate,omitempty"`            // Certificate used for authorization for the RADSEC Server
	Host                     string `json:"host,omitempty"`                     // IP address (or FQDN) to which the APs will send RADIUS accounting messages
	OpenRoamingCertificateID *int   `json:"openRoamingCertificateId,omitempty"` // The ID of the Openroaming Certificate attached to radius server
	Port                     *int   `json:"port,omitempty"`                     // Port on the RADIUS server that is listening for accounting messages
}
type ResponseWirelessGetNetworkWirelessSSIDRadiusServers struct {
	CaCertificate            string `json:"caCertificate,omitempty"`            // Certificate used for authorization for the RADSEC Server
	Host                     string `json:"host,omitempty"`                     // IP address (or FQDN) of your RADIUS server
	OpenRoamingCertificateID *int   `json:"openRoamingCertificateId,omitempty"` // The ID of the Openroaming Certificate attached to radius server
	Port                     *int   `json:"port,omitempty"`                     // UDP port the RADIUS server listens on for Access-requests
}
type ResponseWirelessUpdateNetworkWirelessSSID struct {
	AdminSplashURL                  string                                                              `json:"adminSplashUrl,omitempty"`                  // URL for the admin splash page
	AuthMode                        string                                                              `json:"authMode,omitempty"`                        // The association control method for the SSID
	AvailabilityTags                []string                                                            `json:"availabilityTags,omitempty"`                // List of tags for this SSID. If availableOnAllAps is false, then the SSID is only broadcast by APs with tags matching any of the tags in this list
	AvailableOnAllAps               *bool                                                               `json:"availableOnAllAps,omitempty"`               // Whether all APs broadcast the SSID or if it's restricted to APs matching any availability tags
	BandSelection                   string                                                              `json:"bandSelection,omitempty"`                   // The client-serving radio frequencies of this SSID in the default indoor RF profile
	Enabled                         *bool                                                               `json:"enabled,omitempty"`                         // Whether or not the SSID is enabled
	EncryptionMode                  string                                                              `json:"encryptionMode,omitempty"`                  // The psk encryption mode for the SSID
	IPAssignmentMode                string                                                              `json:"ipAssignmentMode,omitempty"`                // The client IP assignment mode
	LocalAuth                       *bool                                                               `json:"localAuth,omitempty"`                       // Extended local auth flag for Enterprise NAC
	MandatoryDhcpEnabled            *bool                                                               `json:"mandatoryDhcpEnabled,omitempty"`            // Whether clients connecting to this SSID must use the IP address assigned by the DHCP server
	MinBitrate                      *int                                                                `json:"minBitrate,omitempty"`                      // The minimum bitrate in Mbps of this SSID in the default indoor RF profile
	Name                            string                                                              `json:"name,omitempty"`                            // The name of the SSID
	Number                          *int                                                                `json:"number,omitempty"`                          // Unique identifier of the SSID
	PerClientBandwidthLimitDown     *int                                                                `json:"perClientBandwidthLimitDown,omitempty"`     // The download bandwidth limit in Kbps. (0 represents no limit.)
	PerClientBandwidthLimitUp       *int                                                                `json:"perClientBandwidthLimitUp,omitempty"`       // The upload bandwidth limit in Kbps. (0 represents no limit.)
	PerSSIDBandwidthLimitDown       *int                                                                `json:"perSsidBandwidthLimitDown,omitempty"`       // The total download bandwidth limit in Kbps (0 represents no limit)
	PerSSIDBandwidthLimitUp         *int                                                                `json:"perSsidBandwidthLimitUp,omitempty"`         // The total upload bandwidth limit in Kbps (0 represents no limit)
	RadiusAccountingEnabled         *bool                                                               `json:"radiusAccountingEnabled,omitempty"`         // Whether or not RADIUS accounting is enabled
	RadiusAccountingServers         *[]ResponseWirelessUpdateNetworkWirelessSSIDRadiusAccountingServers `json:"radiusAccountingServers,omitempty"`         // List of RADIUS accounting 802.1X servers to be used for authentication
	RadiusAttributeForGroupPolicies string                                                              `json:"radiusAttributeForGroupPolicies,omitempty"` // RADIUS attribute used to look up group policies
	RadiusEnabled                   *bool                                                               `json:"radiusEnabled,omitempty"`                   // Whether RADIUS authentication is enabled
	RadiusFailoverPolicy            string                                                              `json:"radiusFailoverPolicy,omitempty"`            // Policy which determines how authentication requests should be handled in the event that all of the configured RADIUS servers are unreachable
	RadiusLoadBalancingPolicy       string                                                              `json:"radiusLoadBalancingPolicy,omitempty"`       // Policy which determines which RADIUS server will be contacted first in an authentication attempt, and the ordering of any necessary retry attempts
	RadiusServers                   *[]ResponseWirelessUpdateNetworkWirelessSSIDRadiusServers           `json:"radiusServers,omitempty"`                   // List of RADIUS 802.1X servers to be used for authentication
	SplashPage                      string                                                              `json:"splashPage,omitempty"`                      // The type of splash page for the SSID
	SplashTimeout                   string                                                              `json:"splashTimeout,omitempty"`                   // Splash page timeout
	SSIDAdminAccessible             *bool                                                               `json:"ssidAdminAccessible,omitempty"`             // SSID Administrator access status
	Visible                         *bool                                                               `json:"visible,omitempty"`                         // Whether the SSID is advertised or hidden by the AP
	WalledGardenEnabled             *bool                                                               `json:"walledGardenEnabled,omitempty"`             // Allow users to access a configurable list of IP ranges prior to sign-on
	WalledGardenRanges              []string                                                            `json:"walledGardenRanges,omitempty"`              // Domain names and IP address ranges available in Walled Garden mode
	WpaEncryptionMode               string                                                              `json:"wpaEncryptionMode,omitempty"`               // The types of WPA encryption
}
type ResponseWirelessUpdateNetworkWirelessSSIDRadiusAccountingServers struct {
	CaCertificate            string `json:"caCertificate,omitempty"`            // Certificate used for authorization for the RADSEC Server
	Host                     string `json:"host,omitempty"`                     // IP address (or FQDN) to which the APs will send RADIUS accounting messages
	OpenRoamingCertificateID *int   `json:"openRoamingCertificateId,omitempty"` // The ID of the Openroaming Certificate attached to radius server
	Port                     *int   `json:"port,omitempty"`                     // Port on the RADIUS server that is listening for accounting messages
}
type ResponseWirelessUpdateNetworkWirelessSSIDRadiusServers struct {
	CaCertificate            string `json:"caCertificate,omitempty"`            // Certificate used for authorization for the RADSEC Server
	Host                     string `json:"host,omitempty"`                     // IP address (or FQDN) of your RADIUS server
	OpenRoamingCertificateID *int   `json:"openRoamingCertificateId,omitempty"` // The ID of the Openroaming Certificate attached to radius server
	Port                     *int   `json:"port,omitempty"`                     // UDP port the RADIUS server listens on for Access-requests
}
type ResponseWirelessGetNetworkWirelessSSIDBonjourForwarding struct {
	Enabled   *bool                                                             `json:"enabled,omitempty"`   // If true, Bonjour forwarding is enabled on the SSID.
	Exception *ResponseWirelessGetNetworkWirelessSSIDBonjourForwardingException `json:"exception,omitempty"` // Bonjour forwarding exception
	Rules     *[]ResponseWirelessGetNetworkWirelessSSIDBonjourForwardingRules   `json:"rules,omitempty"`     // Bonjour forwarding rules
}
type ResponseWirelessGetNetworkWirelessSSIDBonjourForwardingException struct {
	Enabled *bool `json:"enabled,omitempty"` // If true, Bonjour forwarding exception is enabled on this SSID. Exception is required to enable L2 isolation and Bonjour forwarding to work together.
}
type ResponseWirelessGetNetworkWirelessSSIDBonjourForwardingRules struct {
	Description string   `json:"description,omitempty"` // Desctiption of the bonjour forwarding rule
	Services    []string `json:"services,omitempty"`    // A list of Bonjour services. At least one service must be specified. Available services are 'All Services', 'AirPlay', 'AFP', 'BitTorrent', 'FTP', 'iChat', 'iTunes', 'Printers', 'Samba', 'Scanners' and 'SSH'
	VLANID      string   `json:"vlanId,omitempty"`      // The ID of the service VLAN. Required
}
type ResponseWirelessUpdateNetworkWirelessSSIDBonjourForwarding struct {
	Enabled   *bool                                                                `json:"enabled,omitempty"`   // If true, Bonjour forwarding is enabled on the SSID.
	Exception *ResponseWirelessUpdateNetworkWirelessSSIDBonjourForwardingException `json:"exception,omitempty"` // Bonjour forwarding exception
	Rules     *[]ResponseWirelessUpdateNetworkWirelessSSIDBonjourForwardingRules   `json:"rules,omitempty"`     // Bonjour forwarding rules
}
type ResponseWirelessUpdateNetworkWirelessSSIDBonjourForwardingException struct {
	Enabled *bool `json:"enabled,omitempty"` // If true, Bonjour forwarding exception is enabled on this SSID. Exception is required to enable L2 isolation and Bonjour forwarding to work together.
}
type ResponseWirelessUpdateNetworkWirelessSSIDBonjourForwardingRules struct {
	Description string   `json:"description,omitempty"` // Desctiption of the bonjour forwarding rule
	Services    []string `json:"services,omitempty"`    // A list of Bonjour services. At least one service must be specified. Available services are 'All Services', 'AirPlay', 'AFP', 'BitTorrent', 'FTP', 'iChat', 'iTunes', 'Printers', 'Samba', 'Scanners' and 'SSH'
	VLANID      string   `json:"vlanId,omitempty"`      // The ID of the service VLAN. Required
}
type ResponseWirelessGetNetworkWirelessSSIDDeviceTypeGroupPolicies struct {
	DeviceTypePolicies *[]ResponseWirelessGetNetworkWirelessSSIDDeviceTypeGroupPoliciesDeviceTypePolicies `json:"deviceTypePolicies,omitempty"` //
	Enabled            *bool                                                                              `json:"enabled,omitempty"`            //
}
type ResponseWirelessGetNetworkWirelessSSIDDeviceTypeGroupPoliciesDeviceTypePolicies struct {
	DevicePolicy string `json:"devicePolicy,omitempty"` //
	DeviceType   string `json:"deviceType,omitempty"`   //
}
type ResponseWirelessUpdateNetworkWirelessSSIDDeviceTypeGroupPolicies interface{}
type ResponseWirelessGetNetworkWirelessSSIDEapOverride struct {
	EapolKey   *ResponseWirelessGetNetworkWirelessSSIDEapOverrideEapolKey `json:"eapolKey,omitempty"`   // EAPOL Key settings.
	IDentity   *ResponseWirelessGetNetworkWirelessSSIDEapOverrideIDentity `json:"identity,omitempty"`   // EAP settings for identity requests.
	MaxRetries *int                                                       `json:"maxRetries,omitempty"` // Maximum number of general EAP retries.
	Timeout    *int                                                       `json:"timeout,omitempty"`    // General EAP timeout in seconds.
}
type ResponseWirelessGetNetworkWirelessSSIDEapOverrideEapolKey struct {
	Retries     *int `json:"retries,omitempty"`     // Maximum number of EAPOL key retries.
	TimeoutInMs *int `json:"timeoutInMs,omitempty"` // EAPOL Key timeout in milliseconds.
}
type ResponseWirelessGetNetworkWirelessSSIDEapOverrideIDentity struct {
	Retries *int `json:"retries,omitempty"` // Maximum number of EAP retries.
	Timeout *int `json:"timeout,omitempty"` // EAP timeout in seconds.
}
type ResponseWirelessUpdateNetworkWirelessSSIDEapOverride struct {
	EapolKey   *ResponseWirelessUpdateNetworkWirelessSSIDEapOverrideEapolKey `json:"eapolKey,omitempty"`   // EAPOL Key settings.
	IDentity   *ResponseWirelessUpdateNetworkWirelessSSIDEapOverrideIDentity `json:"identity,omitempty"`   // EAP settings for identity requests.
	MaxRetries *int                                                          `json:"maxRetries,omitempty"` // Maximum number of general EAP retries.
	Timeout    *int                                                          `json:"timeout,omitempty"`    // General EAP timeout in seconds.
}
type ResponseWirelessUpdateNetworkWirelessSSIDEapOverrideEapolKey struct {
	Retries     *int `json:"retries,omitempty"`     // Maximum number of EAPOL key retries.
	TimeoutInMs *int `json:"timeoutInMs,omitempty"` // EAPOL Key timeout in milliseconds.
}
type ResponseWirelessUpdateNetworkWirelessSSIDEapOverrideIDentity struct {
	Retries *int `json:"retries,omitempty"` // Maximum number of EAP retries.
	Timeout *int `json:"timeout,omitempty"` // EAP timeout in seconds.
}
type ResponseWirelessGetNetworkWirelessSSIDFirewallL3FirewallRules struct {
	Rules *[]ResponseWirelessGetNetworkWirelessSSIDFirewallL3FirewallRulesRules `json:"rules,omitempty"` // An ordered array of the firewall rules for this SSID (not including the local LAN access rule or the default rule).
}
type ResponseWirelessGetNetworkWirelessSSIDFirewallL3FirewallRulesRules struct {
	Comment  string `json:"comment,omitempty"`  // Description of the rule (optional)
	DestCidr string `json:"destCidr,omitempty"` // Comma-separated list of destination IP address(es) (in IP or CIDR notation), fully-qualified domain names (FQDN) or 'any'
	DestPort string `json:"destPort,omitempty"` // Comma-separated list of destination port(s) (integer in the range 1-65535), or 'any'
	Policy   string `json:"policy,omitempty"`   // 'allow' or 'deny' traffic specified by this rule
	Protocol string `json:"protocol,omitempty"` // The type of protocol (must be 'tcp', 'udp', 'icmp', 'icmp6' or 'any')
	IpVer    string `json:"ipVer,omitempty"`    //
}
type ResponseWirelessUpdateNetworkWirelessSSIDFirewallL3FirewallRules struct {
	Rules *[]ResponseWirelessUpdateNetworkWirelessSSIDFirewallL3FirewallRulesRules `json:"rules,omitempty"` // An ordered array of the firewall rules for this SSID (not including the local LAN access rule or the default rule).
}
type ResponseWirelessUpdateNetworkWirelessSSIDFirewallL3FirewallRulesRules struct {
	Comment  string `json:"comment,omitempty"`  // Description of the rule (optional)
	DestCidr string `json:"destCidr,omitempty"` // Comma-separated list of destination IP address(es) (in IP or CIDR notation), fully-qualified domain names (FQDN) or 'any'
	DestPort string `json:"destPort,omitempty"` // Comma-separated list of destination port(s) (integer in the range 1-65535), or 'any'
	Policy   string `json:"policy,omitempty"`   // 'allow' or 'deny' traffic specified by this rule
	Protocol string `json:"protocol,omitempty"` // The type of protocol (must be 'tcp', 'udp', 'icmp', 'icmp6' or 'any')
}
type ResponseWirelessGetNetworkWirelessSSIDFirewallL7FirewallRules struct {
	Rules *[]ResponseWirelessGetNetworkWirelessSSIDFirewallL7FirewallRulesRules `json:"rules,omitempty"` // An ordered array of the firewall rules for this SSID (not including the local LAN access rule or the default rule).
}
type ResponseWirelessGetNetworkWirelessSSIDFirewallL7FirewallRulesRules struct {
	Policy    string                                                                    `json:"policy,omitempty"` // 'Deny' traffic specified by this rule
	Type      string                                                                    `json:"type,omitempty"`   // Type of the L7 firewall rule. One of: 'application', 'applicationCategory', 'host', 'port', 'ipRange'
	Value     *string                                                                   //
	ValueObj  *ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesRulesValueObj //
	ValueList *[]string                                                                 //
}
type ResponseWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRules struct {
	Rules *[]ResponseWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRulesRules `json:"rules,omitempty"` // An ordered array of the firewall rules for this SSID (not including the local LAN access rule or the default rule).
}
type ResponseWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRulesRules struct {
	Policy    string                                                                    `json:"policy,omitempty"` // 'Deny' traffic specified by this rule
	Type      string                                                                    `json:"type,omitempty"`   // Type of the L7 firewall rule. One of: 'application', 'applicationCategory', 'host', 'port', 'ipRange'
	Value     *string                                                                   //
	ValueObj  *ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesRulesValueObj //
	ValueList *[]string                                                                 //
}
type ResponseWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRulesRulesValueObj struct {
	ID   string `json:"id,omitempty"`   //
	Name string `json:"name,omitempty"` //
}

func (r *ResponseWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRulesRules) UnmarshalJSON(data []byte) error {
	type Alias ResponseWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRulesRules
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
func (r *ResponseWirelessGetNetworkWirelessSSIDFirewallL7FirewallRulesRules) UnmarshalJSON(data []byte) error {
	type Alias ResponseWirelessGetNetworkWirelessSSIDFirewallL7FirewallRulesRules
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

type ResponseWirelessGetNetworkWirelessSSIDHotspot20 struct {
	Domains           []string                                                    `json:"domains,omitempty"`           //
	Enabled           *bool                                                       `json:"enabled,omitempty"`           //
	MccMncs           *[]ResponseWirelessGetNetworkWirelessSSIDHotspot20MccMncs   `json:"mccMncs,omitempty"`           //
	NaiRealms         *[]ResponseWirelessGetNetworkWirelessSSIDHotspot20NaiRealms `json:"naiRealms,omitempty"`         //
	NetworkAccessType string                                                      `json:"networkAccessType,omitempty"` //
	Operator          *ResponseWirelessGetNetworkWirelessSSIDHotspot20Operator    `json:"operator,omitempty"`          //
	RoamConsortOis    []string                                                    `json:"roamConsortOis,omitempty"`    //
	Venue             *ResponseWirelessGetNetworkWirelessSSIDHotspot20Venue       `json:"venue,omitempty"`             //
}
type ResponseWirelessGetNetworkWirelessSSIDHotspot20MccMncs struct {
	Mcc string `json:"mcc,omitempty"` //
	Mnc string `json:"mnc,omitempty"` //
}
type ResponseWirelessGetNetworkWirelessSSIDHotspot20NaiRealms struct {
	Format  string                                                             `json:"format,omitempty"`  //
	Methods *[]ResponseWirelessGetNetworkWirelessSSIDHotspot20NaiRealmsMethods `json:"methods,omitempty"` //
	Name    string                                                             `json:"name,omitempty"`    //
}
type ResponseWirelessGetNetworkWirelessSSIDHotspot20NaiRealmsMethods struct {
	AuthenticationTypes *ResponseWirelessGetNetworkWirelessSSIDHotspot20NaiRealmsMethodsAuthenticationTypes `json:"authenticationTypes,omitempty"` //
	ID                  string                                                                              `json:"id,omitempty"`                  //
}
type ResponseWirelessGetNetworkWirelessSSIDHotspot20NaiRealmsMethodsAuthenticationTypes struct {
	Credentials                  []string `json:"credentials,omitempty"`                  //
	EapinnerAuthentication       []string `json:"eapInnerAuthentication,omitempty"`       //
	NonEapinnerAuthentication    []string `json:"nonEapInnerAuthentication,omitempty"`    //
	TunneledEapMethodCredentials []string `json:"tunneledEapMethodCredentials,omitempty"` //
}
type ResponseWirelessGetNetworkWirelessSSIDHotspot20Operator struct {
	Name string `json:"name,omitempty"` //
}
type ResponseWirelessGetNetworkWirelessSSIDHotspot20Venue struct {
	Name string `json:"name,omitempty"` //
	Type string `json:"type,omitempty"` //
}
type ResponseWirelessUpdateNetworkWirelessSSIDHotspot20 interface{}
type ResponseWirelessGetNetworkWirelessSSIDIDentityPsks []ResponseItemWirelessGetNetworkWirelessSSIDIDentityPsks // Array of ResponseWirelessGetNetworkWirelessSsidIdentityPsks
type ResponseItemWirelessGetNetworkWirelessSSIDIDentityPsks struct {
	Email                 string `json:"email,omitempty"`                 // The email associated with the System's Manager User
	ExpiresAt             string `json:"expiresAt,omitempty"`             // Timestamp for when the Identity PSK expires, or 'null' to never expire
	GroupPolicyID         string `json:"groupPolicyId,omitempty"`         // The group policy to be applied to clients
	ID                    string `json:"id,omitempty"`                    // The unique identifier of the Identity PSK
	Name                  string `json:"name,omitempty"`                  // The name of the Identity PSK
	Passphrase            string `json:"passphrase,omitempty"`            // The passphrase for client authentication
	WifiPersonalNetworkID string `json:"wifiPersonalNetworkId,omitempty"` // The WiFi Personal Network unique identifier
}
type ResponseWirelessCreateNetworkWirelessSSIDIDentityPsk struct {
	Email                 string `json:"email,omitempty"`                 // The email associated with the System's Manager User
	ExpiresAt             string `json:"expiresAt,omitempty"`             // Timestamp for when the Identity PSK expires, or 'null' to never expire
	GroupPolicyID         string `json:"groupPolicyId,omitempty"`         // The group policy to be applied to clients
	ID                    string `json:"id,omitempty"`                    // The unique identifier of the Identity PSK
	Name                  string `json:"name,omitempty"`                  // The name of the Identity PSK
	Passphrase            string `json:"passphrase,omitempty"`            // The passphrase for client authentication
	WifiPersonalNetworkID string `json:"wifiPersonalNetworkId,omitempty"` // The WiFi Personal Network unique identifier
}
type ResponseWirelessGetNetworkWirelessSSIDIDentityPsk struct {
	Email                 string `json:"email,omitempty"`                 // The email associated with the System's Manager User
	ExpiresAt             string `json:"expiresAt,omitempty"`             // Timestamp for when the Identity PSK expires, or 'null' to never expire
	GroupPolicyID         string `json:"groupPolicyId,omitempty"`         // The group policy to be applied to clients
	ID                    string `json:"id,omitempty"`                    // The unique identifier of the Identity PSK
	Name                  string `json:"name,omitempty"`                  // The name of the Identity PSK
	Passphrase            string `json:"passphrase,omitempty"`            // The passphrase for client authentication
	WifiPersonalNetworkID string `json:"wifiPersonalNetworkId,omitempty"` // The WiFi Personal Network unique identifier
}
type ResponseWirelessUpdateNetworkWirelessSSIDIDentityPsk struct {
	Email                 string `json:"email,omitempty"`                 // The email associated with the System's Manager User
	ExpiresAt             string `json:"expiresAt,omitempty"`             // Timestamp for when the Identity PSK expires, or 'null' to never expire
	GroupPolicyID         string `json:"groupPolicyId,omitempty"`         // The group policy to be applied to clients
	ID                    string `json:"id,omitempty"`                    // The unique identifier of the Identity PSK
	Name                  string `json:"name,omitempty"`                  // The name of the Identity PSK
	Passphrase            string `json:"passphrase,omitempty"`            // The passphrase for client authentication
	WifiPersonalNetworkID string `json:"wifiPersonalNetworkId,omitempty"` // The WiFi Personal Network unique identifier
}
type ResponseWirelessGetNetworkWirelessSSIDSchedules struct {
	Enabled *bool                                                    `json:"enabled,omitempty"` //
	Ranges  *[]ResponseWirelessGetNetworkWirelessSSIDSchedulesRanges `json:"ranges,omitempty"`  //
}
type ResponseWirelessGetNetworkWirelessSSIDSchedulesRanges struct {
	EndDay    string `json:"endDay,omitempty"`    //
	EndTime   string `json:"endTime,omitempty"`   //
	StartDay  string `json:"startDay,omitempty"`  //
	StartTime string `json:"startTime,omitempty"` //
}
type ResponseWirelessUpdateNetworkWirelessSSIDSchedules interface{}
type ResponseWirelessGetNetworkWirelessSSIDSplashSettings struct {
	AllowSimultaneousLogins         *bool                                                                   `json:"allowSimultaneousLogins,omitempty"`         // Whether or not to allow simultaneous logins from different devices.
	Billing                         *ResponseWirelessGetNetworkWirelessSSIDSplashSettingsBilling            `json:"billing,omitempty"`                         // Details associated with billing splash
	BlockAllTrafficBeforeSignOn     *bool                                                                   `json:"blockAllTrafficBeforeSignOn,omitempty"`     // How restricted allowing traffic should be. If true, all traffic types are blocked until the splash page is acknowledged. If false, all non-HTTP traffic is allowed before the splash page is acknowledged.
	ControllerDisconnectionBehavior string                                                                  `json:"controllerDisconnectionBehavior,omitempty"` // How login attempts should be handled when the controller is unreachable.
	GuestSponsorship                *ResponseWirelessGetNetworkWirelessSSIDSplashSettingsGuestSponsorship   `json:"guestSponsorship,omitempty"`                // Details associated with guest sponsored splash
	RedirectURL                     string                                                                  `json:"redirectUrl,omitempty"`                     // The custom redirect URL where the users will go after the splash page.
	SelfRegistration                *ResponseWirelessGetNetworkWirelessSSIDSplashSettingsSelfRegistration   `json:"selfRegistration,omitempty"`                // Self-registration for splash with Meraki authentication.
	SentryEnrollment                *ResponseWirelessGetNetworkWirelessSSIDSplashSettingsSentryEnrollment   `json:"sentryEnrollment,omitempty"`                // Systems Manager sentry enrollment splash settings.
	SplashImage                     *ResponseWirelessGetNetworkWirelessSSIDSplashSettingsSplashImage        `json:"splashImage,omitempty"`                     // The image used in the splash page.
	SplashLogo                      *ResponseWirelessGetNetworkWirelessSSIDSplashSettingsSplashLogo         `json:"splashLogo,omitempty"`                      // The logo used in the splash page.
	SplashPage                      string                                                                  `json:"splashPage,omitempty"`                      // The type of splash page for this SSID
	SplashPrepaidFront              *ResponseWirelessGetNetworkWirelessSSIDSplashSettingsSplashPrepaidFront `json:"splashPrepaidFront,omitempty"`              // The prepaid front image used in the splash page.
	SplashTimeout                   *int                                                                    `json:"splashTimeout,omitempty"`                   // Splash timeout in minutes.
	SplashURL                       string                                                                  `json:"splashUrl,omitempty"`                       // The custom splash URL of the click-through splash page.
	SSIDNumber                      *int                                                                    `json:"ssidNumber,omitempty"`                      // SSID number
	ThemeID                         string                                                                  `json:"themeId,omitempty"`                         // The id of the selected splash theme.
	UseRedirectURL                  *bool                                                                   `json:"useRedirectUrl,omitempty"`                  // The Boolean indicating whether the the user will be redirected to the custom redirect URL after the splash page.
	UseSplashURL                    *bool                                                                   `json:"useSplashUrl,omitempty"`                    // Boolean indicating whether the users will be redirected to the custom splash url
	WelcomeMessage                  string                                                                  `json:"welcomeMessage,omitempty"`                  // The welcome message for the users on the splash page.
}
type ResponseWirelessGetNetworkWirelessSSIDSplashSettingsBilling struct {
	FreeAccess                    *ResponseWirelessGetNetworkWirelessSSIDSplashSettingsBillingFreeAccess `json:"freeAccess,omitempty"`                    // Details associated with a free access plan with limits
	PrepaidAccessFastLoginEnabled *bool                                                                  `json:"prepaidAccessFastLoginEnabled,omitempty"` // Whether or not billing uses the fast login prepaid access option.
	ReplyToEmailAddress           string                                                                 `json:"replyToEmailAddress,omitempty"`           // The email address that reeceives replies from clients
}
type ResponseWirelessGetNetworkWirelessSSIDSplashSettingsBillingFreeAccess struct {
	DurationInMinutes *int  `json:"durationInMinutes,omitempty"` // How long a device can use a network for free.
	Enabled           *bool `json:"enabled,omitempty"`           // Whether or not free access is enabled.
}
type ResponseWirelessGetNetworkWirelessSSIDSplashSettingsGuestSponsorship struct {
	DurationInMinutes        *int  `json:"durationInMinutes,omitempty"`        // Duration in minutes of sponsored guest authorization.
	GuestCanRequestTimeframe *bool `json:"guestCanRequestTimeframe,omitempty"` // Whether or not guests can specify how much time they are requesting.
}
type ResponseWirelessGetNetworkWirelessSSIDSplashSettingsSelfRegistration struct {
	AuthorizationType string `json:"authorizationType,omitempty"` // How created user accounts should be authorized.
	Enabled           *bool  `json:"enabled,omitempty"`           // Whether or not to allow users to create their own account on the network.
}
type ResponseWirelessGetNetworkWirelessSSIDSplashSettingsSentryEnrollment struct {
	EnforcedSystems       []string                                                                                   `json:"enforcedSystems,omitempty"`       // The system types that the Sentry enforces.
	Strength              string                                                                                     `json:"strength,omitempty"`              // The strength of the enforcement of selected system types.
	SystemsManagerNetwork *ResponseWirelessGetNetworkWirelessSSIDSplashSettingsSentryEnrollmentSystemsManagerNetwork `json:"systemsManagerNetwork,omitempty"` // Systems Manager network targeted for sentry enrollment.
}
type ResponseWirelessGetNetworkWirelessSSIDSplashSettingsSentryEnrollmentSystemsManagerNetwork struct {
	ID string `json:"id,omitempty"` // The network ID of the Systems Manager network.
}
type ResponseWirelessGetNetworkWirelessSSIDSplashSettingsSplashImage struct {
	Extension string `json:"extension,omitempty"` // The extension of the image file.
	Md5       string `json:"md5,omitempty"`       // The MD5 value of the image file.
}
type ResponseWirelessGetNetworkWirelessSSIDSplashSettingsSplashLogo struct {
	Extension string `json:"extension,omitempty"` // The extension of the logo file.
	Md5       string `json:"md5,omitempty"`       // The MD5 value of the logo file.
}
type ResponseWirelessGetNetworkWirelessSSIDSplashSettingsSplashPrepaidFront struct {
	Extension string `json:"extension,omitempty"` // The extension of the prepaid front image file.
	Md5       string `json:"md5,omitempty"`       // The MD5 value of the prepaid front image file.
}
type ResponseWirelessUpdateNetworkWirelessSSIDSplashSettings struct {
	AllowSimultaneousLogins         *bool                                                                      `json:"allowSimultaneousLogins,omitempty"`         // Whether or not to allow simultaneous logins from different devices.
	Billing                         *ResponseWirelessUpdateNetworkWirelessSSIDSplashSettingsBilling            `json:"billing,omitempty"`                         // Details associated with billing splash
	BlockAllTrafficBeforeSignOn     *bool                                                                      `json:"blockAllTrafficBeforeSignOn,omitempty"`     // How restricted allowing traffic should be. If true, all traffic types are blocked until the splash page is acknowledged. If false, all non-HTTP traffic is allowed before the splash page is acknowledged.
	ControllerDisconnectionBehavior string                                                                     `json:"controllerDisconnectionBehavior,omitempty"` // How login attempts should be handled when the controller is unreachable.
	GuestSponsorship                *ResponseWirelessUpdateNetworkWirelessSSIDSplashSettingsGuestSponsorship   `json:"guestSponsorship,omitempty"`                // Details associated with guest sponsored splash
	RedirectURL                     string                                                                     `json:"redirectUrl,omitempty"`                     // The custom redirect URL where the users will go after the splash page.
	SelfRegistration                *ResponseWirelessUpdateNetworkWirelessSSIDSplashSettingsSelfRegistration   `json:"selfRegistration,omitempty"`                // Self-registration for splash with Meraki authentication.
	SentryEnrollment                *ResponseWirelessUpdateNetworkWirelessSSIDSplashSettingsSentryEnrollment   `json:"sentryEnrollment,omitempty"`                // Systems Manager sentry enrollment splash settings.
	SplashImage                     *ResponseWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashImage        `json:"splashImage,omitempty"`                     // The image used in the splash page.
	SplashLogo                      *ResponseWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashLogo         `json:"splashLogo,omitempty"`                      // The logo used in the splash page.
	SplashPage                      string                                                                     `json:"splashPage,omitempty"`                      // The type of splash page for this SSID
	SplashPrepaidFront              *ResponseWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashPrepaidFront `json:"splashPrepaidFront,omitempty"`              // The prepaid front image used in the splash page.
	SplashTimeout                   *int                                                                       `json:"splashTimeout,omitempty"`                   // Splash timeout in minutes.
	SplashURL                       string                                                                     `json:"splashUrl,omitempty"`                       // The custom splash URL of the click-through splash page.
	SSIDNumber                      *int                                                                       `json:"ssidNumber,omitempty"`                      // SSID number
	ThemeID                         string                                                                     `json:"themeId,omitempty"`                         // The id of the selected splash theme.
	UseRedirectURL                  *bool                                                                      `json:"useRedirectUrl,omitempty"`                  // The Boolean indicating whether the the user will be redirected to the custom redirect URL after the splash page.
	UseSplashURL                    *bool                                                                      `json:"useSplashUrl,omitempty"`                    // Boolean indicating whether the users will be redirected to the custom splash url
	WelcomeMessage                  string                                                                     `json:"welcomeMessage,omitempty"`                  // The welcome message for the users on the splash page.
}
type ResponseWirelessUpdateNetworkWirelessSSIDSplashSettingsBilling struct {
	FreeAccess                    *ResponseWirelessUpdateNetworkWirelessSSIDSplashSettingsBillingFreeAccess `json:"freeAccess,omitempty"`                    // Details associated with a free access plan with limits
	PrepaidAccessFastLoginEnabled *bool                                                                     `json:"prepaidAccessFastLoginEnabled,omitempty"` // Whether or not billing uses the fast login prepaid access option.
	ReplyToEmailAddress           string                                                                    `json:"replyToEmailAddress,omitempty"`           // The email address that reeceives replies from clients
}
type ResponseWirelessUpdateNetworkWirelessSSIDSplashSettingsBillingFreeAccess struct {
	DurationInMinutes *int  `json:"durationInMinutes,omitempty"` // How long a device can use a network for free.
	Enabled           *bool `json:"enabled,omitempty"`           // Whether or not free access is enabled.
}
type ResponseWirelessUpdateNetworkWirelessSSIDSplashSettingsGuestSponsorship struct {
	DurationInMinutes        *int  `json:"durationInMinutes,omitempty"`        // Duration in minutes of sponsored guest authorization.
	GuestCanRequestTimeframe *bool `json:"guestCanRequestTimeframe,omitempty"` // Whether or not guests can specify how much time they are requesting.
}
type ResponseWirelessUpdateNetworkWirelessSSIDSplashSettingsSelfRegistration struct {
	AuthorizationType string `json:"authorizationType,omitempty"` // How created user accounts should be authorized.
	Enabled           *bool  `json:"enabled,omitempty"`           // Whether or not to allow users to create their own account on the network.
}
type ResponseWirelessUpdateNetworkWirelessSSIDSplashSettingsSentryEnrollment struct {
	EnforcedSystems       []string                                                                                      `json:"enforcedSystems,omitempty"`       // The system types that the Sentry enforces.
	Strength              string                                                                                        `json:"strength,omitempty"`              // The strength of the enforcement of selected system types.
	SystemsManagerNetwork *ResponseWirelessUpdateNetworkWirelessSSIDSplashSettingsSentryEnrollmentSystemsManagerNetwork `json:"systemsManagerNetwork,omitempty"` // Systems Manager network targeted for sentry enrollment.
}
type ResponseWirelessUpdateNetworkWirelessSSIDSplashSettingsSentryEnrollmentSystemsManagerNetwork struct {
	ID string `json:"id,omitempty"` // The network ID of the Systems Manager network.
}
type ResponseWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashImage struct {
	Extension string `json:"extension,omitempty"` // The extension of the image file.
	Md5       string `json:"md5,omitempty"`       // The MD5 value of the image file.
}
type ResponseWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashLogo struct {
	Extension string `json:"extension,omitempty"` // The extension of the logo file.
	Md5       string `json:"md5,omitempty"`       // The MD5 value of the logo file.
}
type ResponseWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashPrepaidFront struct {
	Extension string `json:"extension,omitempty"` // The extension of the prepaid front image file.
	Md5       string `json:"md5,omitempty"`       // The MD5 value of the prepaid front image file.
}
type ResponseWirelessGetNetworkWirelessSSIDTrafficShapingRules struct {
	DefaultRulesEnabled   *bool                                                             `json:"defaultRulesEnabled,omitempty"`   // Whether default traffic shaping rules are enabled (true) or disabled (false). There are 4 default rules, which can be seen on your network's traffic shaping page. Note that default rules count against the rule limit of 8.
	Rules                 *[]ResponseWirelessGetNetworkWirelessSSIDTrafficShapingRulesRules `json:"rules,omitempty"`                 //     An array of traffic shaping rules. Rules are applied in the order that     they are specified in. An empty list (or null) means no rules. Note that     you are allowed a maximum of 8 rules.
	TrafficShapingEnabled *bool                                                             `json:"trafficShapingEnabled,omitempty"` // Whether traffic shaping rules are applied to clients on your SSID.
}
type ResponseWirelessGetNetworkWirelessSSIDTrafficShapingRulesRules struct {
	Definitions              *[]ResponseWirelessGetNetworkWirelessSSIDTrafficShapingRulesRulesDefinitions            `json:"definitions,omitempty"`              //     A list of objects describing the definitions of your traffic shaping rule. At least one definition is required.
	DscpTagValue             *int                                                                                    `json:"dscpTagValue,omitempty"`             //     The DSCP tag applied by your rule. null means 'Do not change DSCP tag'.     For a list of possible tag values, use the trafficShaping/dscpTaggingOptions endpoint.
	PcpTagValue              *int                                                                                    `json:"pcpTagValue,omitempty"`              //     The PCP tag applied by your rule. Can be 0 (lowest priority) through 7 (highest priority).     null means 'Do not set PCP tag'.
	PerClientBandwidthLimits *ResponseWirelessGetNetworkWirelessSSIDTrafficShapingRulesRulesPerClientBandwidthLimits `json:"perClientBandwidthLimits,omitempty"` //     An object describing the bandwidth settings for your rule.
}
type ResponseWirelessGetNetworkWirelessSSIDTrafficShapingRulesRulesDefinitions struct {
	Type  string `json:"type,omitempty"`  // The type of definition. Can be one of 'application', 'applicationCategory', 'host', 'port', 'ipRange' or 'localNet'.
	Value string `json:"value,omitempty"` //     If "type" is 'host', 'port', 'ipRange' or 'localNet', then "value" must be a string, matching either     a hostname (e.g. "somesite.com"), a port (e.g. 8080), or an IP range ("192.1.0.0",     "192.1.0.0/16", or "10.1.0.0/16:80"). 'localNet' also supports CIDR notation, excluding     custom ports.      If "type" is 'application' or 'applicationCategory', then "value" must be an object     with the structure { "id": "meraki:layer7/..." }, where "id" is the application category or     application ID (for a list of IDs for your network, use the trafficShaping/applicationCategories     endpoint).
}
type ResponseWirelessGetNetworkWirelessSSIDTrafficShapingRulesRulesPerClientBandwidthLimits struct {
	BandwidthLimits *ResponseWirelessGetNetworkWirelessSSIDTrafficShapingRulesRulesPerClientBandwidthLimitsBandwidthLimits `json:"bandwidthLimits,omitempty"` // The bandwidth limits object, specifying the upload ('limitUp') and download ('limitDown') speed in Kbps. These are only enforced if 'settings' is set to 'custom'.
	Settings        string                                                                                                 `json:"settings,omitempty"`        // How bandwidth limits are applied by your rule. Can be one of 'network default', 'ignore' or 'custom'.
}
type ResponseWirelessGetNetworkWirelessSSIDTrafficShapingRulesRulesPerClientBandwidthLimitsBandwidthLimits struct {
	LimitDown *int `json:"limitDown,omitempty"` // The maximum download limit (integer, in Kbps).
	LimitUp   *int `json:"limitUp,omitempty"`   // The maximum upload limit (integer, in Kbps).
}
type ResponseWirelessUpdateNetworkWirelessSSIDTrafficShapingRules struct {
	DefaultRulesEnabled   *bool                                                                `json:"defaultRulesEnabled,omitempty"`   // Whether default traffic shaping rules are enabled (true) or disabled (false). There are 4 default rules, which can be seen on your network's traffic shaping page. Note that default rules count against the rule limit of 8.
	Rules                 *[]ResponseWirelessUpdateNetworkWirelessSSIDTrafficShapingRulesRules `json:"rules,omitempty"`                 //     An array of traffic shaping rules. Rules are applied in the order that     they are specified in. An empty list (or null) means no rules. Note that     you are allowed a maximum of 8 rules.
	TrafficShapingEnabled *bool                                                                `json:"trafficShapingEnabled,omitempty"` // Whether traffic shaping rules are applied to clients on your SSID.
}
type ResponseWirelessUpdateNetworkWirelessSSIDTrafficShapingRulesRules struct {
	Definitions              *[]ResponseWirelessUpdateNetworkWirelessSSIDTrafficShapingRulesRulesDefinitions            `json:"definitions,omitempty"`              //     A list of objects describing the definitions of your traffic shaping rule. At least one definition is required.
	DscpTagValue             *int                                                                                       `json:"dscpTagValue,omitempty"`             //     The DSCP tag applied by your rule. null means 'Do not change DSCP tag'.     For a list of possible tag values, use the trafficShaping/dscpTaggingOptions endpoint.
	PcpTagValue              *int                                                                                       `json:"pcpTagValue,omitempty"`              //     The PCP tag applied by your rule. Can be 0 (lowest priority) through 7 (highest priority).     null means 'Do not set PCP tag'.
	PerClientBandwidthLimits *ResponseWirelessUpdateNetworkWirelessSSIDTrafficShapingRulesRulesPerClientBandwidthLimits `json:"perClientBandwidthLimits,omitempty"` //     An object describing the bandwidth settings for your rule.
}
type ResponseWirelessUpdateNetworkWirelessSSIDTrafficShapingRulesRulesDefinitions struct {
	Type  string `json:"type,omitempty"`  // The type of definition. Can be one of 'application', 'applicationCategory', 'host', 'port', 'ipRange' or 'localNet'.
	Value string `json:"value,omitempty"` //     If "type" is 'host', 'port', 'ipRange' or 'localNet', then "value" must be a string, matching either     a hostname (e.g. "somesite.com"), a port (e.g. 8080), or an IP range ("192.1.0.0",     "192.1.0.0/16", or "10.1.0.0/16:80"). 'localNet' also supports CIDR notation, excluding     custom ports.      If "type" is 'application' or 'applicationCategory', then "value" must be an object     with the structure { "id": "meraki:layer7/..." }, where "id" is the application category or     application ID (for a list of IDs for your network, use the trafficShaping/applicationCategories     endpoint).
}
type ResponseWirelessUpdateNetworkWirelessSSIDTrafficShapingRulesRulesPerClientBandwidthLimits struct {
	BandwidthLimits *ResponseWirelessUpdateNetworkWirelessSSIDTrafficShapingRulesRulesPerClientBandwidthLimitsBandwidthLimits `json:"bandwidthLimits,omitempty"` // The bandwidth limits object, specifying the upload ('limitUp') and download ('limitDown') speed in Kbps. These are only enforced if 'settings' is set to 'custom'.
	Settings        string                                                                                                    `json:"settings,omitempty"`        // How bandwidth limits are applied by your rule. Can be one of 'network default', 'ignore' or 'custom'.
}
type ResponseWirelessUpdateNetworkWirelessSSIDTrafficShapingRulesRulesPerClientBandwidthLimitsBandwidthLimits struct {
	LimitDown *int `json:"limitDown,omitempty"` // The maximum download limit (integer, in Kbps).
	LimitUp   *int `json:"limitUp,omitempty"`   // The maximum upload limit (integer, in Kbps).
}
type ResponseWirelessGetNetworkWirelessSSIDVpn struct {
	Concentrator *ResponseWirelessGetNetworkWirelessSSIDVpnConcentrator `json:"concentrator,omitempty"` //
	Failover     *ResponseWirelessGetNetworkWirelessSSIDVpnFailover     `json:"failover,omitempty"`     //
	SplitTunnel  *ResponseWirelessGetNetworkWirelessSSIDVpnSplitTunnel  `json:"splitTunnel,omitempty"`  //
}
type ResponseWirelessGetNetworkWirelessSSIDVpnConcentrator struct {
	Name      string `json:"name,omitempty"`      //
	NetworkID string `json:"networkId,omitempty"` //
	VLANID    *int   `json:"vlanId,omitempty"`    //
}
type ResponseWirelessGetNetworkWirelessSSIDVpnFailover struct {
	HeartbeatInterval *int   `json:"heartbeatInterval,omitempty"` //
	IDleTimeout       *int   `json:"idleTimeout,omitempty"`       //
	RequestIP         string `json:"requestIp,omitempty"`         //
}
type ResponseWirelessGetNetworkWirelessSSIDVpnSplitTunnel struct {
	Enabled *bool                                                        `json:"enabled,omitempty"` //
	Rules   *[]ResponseWirelessGetNetworkWirelessSSIDVpnSplitTunnelRules `json:"rules,omitempty"`   //
}
type ResponseWirelessGetNetworkWirelessSSIDVpnSplitTunnelRules struct {
	Comment  string `json:"comment,omitempty"`  //
	DestCidr string `json:"destCidr,omitempty"` //
	DestPort string `json:"destPort,omitempty"` //
	Policy   string `json:"policy,omitempty"`   //
	Protocol string `json:"protocol,omitempty"` //
}
type ResponseWirelessUpdateNetworkWirelessSSIDVpn interface{}
type ResponseWirelessGetNetworkWirelessUsageHistory []ResponseItemWirelessGetNetworkWirelessUsageHistory // Array of ResponseWirelessGetNetworkWirelessUsageHistory
type ResponseItemWirelessGetNetworkWirelessUsageHistory struct {
	EndTs        string `json:"endTs,omitempty"`        // The end time of the query range
	ReceivedKbps *int   `json:"receivedKbps,omitempty"` // Received kilobytes-per-second
	SentKbps     *int   `json:"sentKbps,omitempty"`     // Sent kilobytes-per-second
	StartTs      string `json:"startTs,omitempty"`      // The start time of the query range
	TotalKbps    *int   `json:"totalKbps,omitempty"`    // Total usage in kilobytes-per-second
}
type ResponseWirelessGetOrganizationWirelessDevicesChannelUtilizationByDevice []ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByDevice // Array of ResponseWirelessGetOrganizationWirelessDevicesChannelUtilizationByDevice
type ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByDevice struct {
	ByBand  *[]ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByDeviceByBand `json:"byBand,omitempty"`  // Channel utilization broken down by band.
	Mac     string                                                                                `json:"mac,omitempty"`     // The MAC address of the device.
	Network *ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByDeviceNetwork  `json:"network,omitempty"` // Network for the given utilization metrics.
	Serial  string                                                                                `json:"serial,omitempty"`  // The serial number for the device.
}
type ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByDeviceByBand struct {
	Band    string                                                                                     `json:"band,omitempty"`    // The band for the given metrics.
	NonWifi *ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByDeviceByBandNonWifi `json:"nonWifi,omitempty"` // An object containing non-wifi utilization.
	Total   *ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByDeviceByBandTotal   `json:"total,omitempty"`   // An object containing total channel utilization.
	Wifi    *ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByDeviceByBandWifi    `json:"wifi,omitempty"`    // An object containing wifi utilization.
}
type ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByDeviceByBandNonWifi struct {
	Percentage *float64 `json:"percentage,omitempty"` // Percentage of non-wifi channel utiliation for the given band.
}
type ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByDeviceByBandTotal struct {
	Percentage *float64 `json:"percentage,omitempty"` // Percentage of total channel utiliation for the given band.
}
type ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByDeviceByBandWifi struct {
	Percentage *float64 `json:"percentage,omitempty"` // Percentage of wifi channel utiliation for the given band.
}
type ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByDeviceNetwork struct {
	ID string `json:"id,omitempty"` // Network ID of the given utilization metrics.
}
type ResponseWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetwork []ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetwork // Array of ResponseWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetwork
type ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetwork struct {
	ByBand  *[]ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetworkByBand `json:"byBand,omitempty"`  // Channel utilization broken down by band.
	Network *ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetworkNetwork  `json:"network,omitempty"` // Network for the given utilization metrics.
}
type ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetworkByBand struct {
	Band    string                                                                                      `json:"band,omitempty"`    // The band for the given metrics.
	NonWifi *ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetworkByBandNonWifi `json:"nonWifi,omitempty"` // An object containing non-wifi utilization.
	Total   *ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetworkByBandTotal   `json:"total,omitempty"`   // An object containing total channel utilization.
	Wifi    *ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetworkByBandWifi    `json:"wifi,omitempty"`    // An object containing wifi utilization.
}
type ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetworkByBandNonWifi struct {
	Percentage *float64 `json:"percentage,omitempty"` // Percentage of non-wifi channel utiliation for the given band.
}
type ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetworkByBandTotal struct {
	Percentage *float64 `json:"percentage,omitempty"` // Percentage of total channel utiliation for the given band.
}
type ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetworkByBandWifi struct {
	Percentage *float64 `json:"percentage,omitempty"` // Percentage of wifi channel utiliation for the given band.
}
type ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetworkNetwork struct {
	ID string `json:"id,omitempty"` // Network ID of the given utilization metrics.
}
type ResponseWirelessGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval []ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval // Array of ResponseWirelessGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval
type ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval struct {
	ByBand  *[]ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalByBand `json:"byBand,omitempty"`  // Channel utilization broken down by band.
	EndTs   string                                                                                                 `json:"endTs,omitempty"`   // The end time of the channel utilization interval.
	Mac     string                                                                                                 `json:"mac,omitempty"`     // The MAC address of the device.
	Network *ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalNetwork  `json:"network,omitempty"` // Network for the given utilization metrics.
	Serial  string                                                                                                 `json:"serial,omitempty"`  // The serial number for the device.
	StartTs string                                                                                                 `json:"startTs,omitempty"` // The start time of the channel utilization interval.
}
type ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalByBand struct {
	Band    string                                                                                                      `json:"band,omitempty"`    // The band for the given metrics.
	NonWifi *ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalByBandNonWifi `json:"nonWifi,omitempty"` // An object containing non-wifi utilization.
	Total   *ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalByBandTotal   `json:"total,omitempty"`   // An object containing total channel utilization.
	Wifi    *ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalByBandWifi    `json:"wifi,omitempty"`    // An object containing wifi utilization.
}
type ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalByBandNonWifi struct {
	Percentage *float64 `json:"percentage,omitempty"` // Percentage of non-wifi channel utiliation for the given band.
}
type ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalByBandTotal struct {
	Percentage *float64 `json:"percentage,omitempty"` // Percentage of total channel utiliation for the given band.
}
type ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalByBandWifi struct {
	Percentage *float64 `json:"percentage,omitempty"` // Percentage of wifi channel utiliation for the given band.
}
type ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalNetwork struct {
	ID string `json:"id,omitempty"` // Network ID of the given utilization metrics.
}
type ResponseWirelessGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval []ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval // Array of ResponseWirelessGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval
type ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval struct {
	ByBand  *[]ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalByBand `json:"byBand,omitempty"`  // Channel utilization broken down by band.
	EndTs   string                                                                                                  `json:"endTs,omitempty"`   // The end time of the channel utilization interval.
	Network *ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalNetwork  `json:"network,omitempty"` // Network for the given utilization metrics.
	StartTs string                                                                                                  `json:"startTs,omitempty"` // The start time of the channel utilization interval.
}
type ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalByBand struct {
	Band    string                                                                                                       `json:"band,omitempty"`    // The band for the given metrics.
	NonWifi *ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalByBandNonWifi `json:"nonWifi,omitempty"` // An object containing non-wifi utilization.
	Total   *ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalByBandTotal   `json:"total,omitempty"`   // An object containing total channel utilization.
	Wifi    *ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalByBandWifi    `json:"wifi,omitempty"`    // An object containing wifi utilization.
}
type ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalByBandNonWifi struct {
	Percentage *float64 `json:"percentage,omitempty"` // Percentage of non-wifi channel utiliation for the given band.
}
type ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalByBandTotal struct {
	Percentage *float64 `json:"percentage,omitempty"` // Percentage of total channel utiliation for the given band.
}
type ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalByBandWifi struct {
	Percentage *float64 `json:"percentage,omitempty"` // Percentage of wifi channel utiliation for the given band.
}
type ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalNetwork struct {
	ID string `json:"id,omitempty"` // Network ID of the given utilization metrics.
}
type ResponseWirelessGetOrganizationWirelessDevicesEthernetStatuses []ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatuses // Array of ResponseWirelessGetOrganizationWirelessDevicesEthernetStatuses
type ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatuses struct {
	Aggregation *ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesAggregation `json:"aggregation,omitempty"` // Aggregation details object
	Name        string                                                                         `json:"name,omitempty"`        // The name of the AP
	Network     *ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesNetwork     `json:"network,omitempty"`     // Network details object
	Ports       *[]ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesPorts     `json:"ports,omitempty"`       // List of port details
	Power       *ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesPower       `json:"power,omitempty"`       // Power details object
	Serial      string                                                                         `json:"serial,omitempty"`      // The serial number of the AP
}
type ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesAggregation struct {
	Enabled *bool `json:"enabled,omitempty"` // Link Aggregation enabled flag
	Speed   *int  `json:"speed,omitempty"`   // Link Aggregation speed
}
type ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesNetwork struct {
	ID string `json:"id,omitempty"` // The network ID the AP is associated to
}
type ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesPorts struct {
	LinkNegotiation *ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesPortsLinkNegotiation `json:"linkNegotiation,omitempty"` // Link negotiation details object for the port
	Name            string                                                                                  `json:"name,omitempty"`            // Label of the port
	Poe             *ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesPortsPoe             `json:"poe,omitempty"`             // PoE details object for the port
}
type ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesPortsLinkNegotiation struct {
	Duplex string `json:"duplex,omitempty"` // The duplex mode of the port. Can be 'full' or 'half'
	Speed  *int   `json:"speed,omitempty"`  // The speed of the port
}
type ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesPortsPoe struct {
	Standard string `json:"standard,omitempty"` // The PoE Standard for the port. Can be '802.3at', '802.3af', '802.3bt', or null
}
type ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesPower struct {
	Ac   *ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesPowerAc  `json:"ac,omitempty"`   // AC power details object
	Mode string                                                                      `json:"mode,omitempty"` // The PoE power mode for the AP. Can be 'full' or 'low'
	Poe  *ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesPowerPoe `json:"poe,omitempty"`  // PoE power details object
}
type ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesPowerAc struct {
	IsConnected *bool `json:"isConnected,omitempty"` // AC power connected
}
type ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesPowerPoe struct {
	IsConnected *bool `json:"isConnected,omitempty"` // PoE power connected
}
type ResponseWirelessGetOrganizationWirelessDevicesPacketLossByClient []ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByClient // Array of ResponseWirelessGetOrganizationWirelessDevicesPacketLossByClient
type ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByClient struct {
	Client     *ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByClientClient     `json:"client,omitempty"`     // Client.
	Downstream *ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByClientDownstream `json:"downstream,omitempty"` // Packets sent from an AP to a client.
	Network    *ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByClientNetwork    `json:"network,omitempty"`    // Network.
	Upstream   *ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByClientUpstream   `json:"upstream,omitempty"`   // Packets sent from a client to an AP.
}
type ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByClientClient struct {
	ID  string `json:"id,omitempty"`  // Client ID.
	Mac string `json:"mac,omitempty"` // MAC address.
}
type ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByClientDownstream struct {
	LossPercentage *float64 `json:"lossPercentage,omitempty"` // Percentage of lost packets.
	Lost           *int     `json:"lost,omitempty"`           // Total packets sent by an AP that did not reach the client.
	Total          *int     `json:"total,omitempty"`          // Total packets received by a client.
}
type ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByClientNetwork struct {
	ID   string `json:"id,omitempty"`   // Network ID.
	Name string `json:"name,omitempty"` // Name of the network.
}
type ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByClientUpstream struct {
	LossPercentage *float64 `json:"lossPercentage,omitempty"` // Percentage of lost packets.
	Lost           *int     `json:"lost,omitempty"`           // Total packets sent by a client and did not reach the AP.
	Total          *int     `json:"total,omitempty"`          // Total packets sent by a client to an AP.
}
type ResponseWirelessGetOrganizationWirelessDevicesPacketLossByDevice []ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByDevice // Array of ResponseWirelessGetOrganizationWirelessDevicesPacketLossByDevice
type ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByDevice struct {
	Device     *ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByDeviceDevice     `json:"device,omitempty"`     // Device.
	Downstream *ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByDeviceDownstream `json:"downstream,omitempty"` // Packets sent from an AP to a client.
	Network    *ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByDeviceNetwork    `json:"network,omitempty"`    // Network.
	Upstream   *ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByDeviceUpstream   `json:"upstream,omitempty"`   // Packets sent from a client to an AP.
}
type ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByDeviceDevice struct {
	Mac    string `json:"mac,omitempty"`    // MAC address
	Name   string `json:"name,omitempty"`   // Name
	Serial string `json:"serial,omitempty"` // Serial Number
}
type ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByDeviceDownstream struct {
	LossPercentage *float64 `json:"lossPercentage,omitempty"` // Percentage of lost packets.
	Lost           *int     `json:"lost,omitempty"`           // Total packets sent by an AP that did not reach the client.
	Total          *int     `json:"total,omitempty"`          // Total packets received by a client.
}
type ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByDeviceNetwork struct {
	ID   string `json:"id,omitempty"`   // Network ID.
	Name string `json:"name,omitempty"` // Name of the network.
}
type ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByDeviceUpstream struct {
	LossPercentage *float64 `json:"lossPercentage,omitempty"` // Percentage of lost packets.
	Lost           *int     `json:"lost,omitempty"`           // Total packets sent by a client and did not reach the AP.
	Total          *int     `json:"total,omitempty"`          // Total packets sent by a client to an AP.
}
type ResponseWirelessGetOrganizationWirelessDevicesPacketLossByNetwork []ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByNetwork // Array of ResponseWirelessGetOrganizationWirelessDevicesPacketLossByNetwork
type ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByNetwork struct {
	Downstream *ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByNetworkDownstream `json:"downstream,omitempty"` // Packets sent from an AP to a client.
	Network    *ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByNetworkNetwork    `json:"network,omitempty"`    // Network.
	Upstream   *ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByNetworkUpstream   `json:"upstream,omitempty"`   // Packets sent from a client to an AP.
}
type ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByNetworkDownstream struct {
	LossPercentage *float64 `json:"lossPercentage,omitempty"` // Percentage of lost packets.
	Lost           *int     `json:"lost,omitempty"`           // Total packets sent by an AP that did not reach the client.
	Total          *int     `json:"total,omitempty"`          // Total packets received by a client.
}
type ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByNetworkNetwork struct {
	ID   string `json:"id,omitempty"`   // Network ID.
	Name string `json:"name,omitempty"` // Name of the network.
}
type ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByNetworkUpstream struct {
	LossPercentage *float64 `json:"lossPercentage,omitempty"` // Percentage of lost packets.
	Lost           *int     `json:"lost,omitempty"`           // Total packets sent by a client and did not reach the AP.
	Total          *int     `json:"total,omitempty"`          // Total packets sent by a client to an AP.
}
type RequestWirelessUpdateDeviceWirelessAlternateManagementInterfaceIPv6 struct {
	Addresses *[]RequestWirelessUpdateDeviceWirelessAlternateManagementInterfaceIPv6Addresses `json:"addresses,omitempty"` // configured alternate management interface addresses
}
type RequestWirelessUpdateDeviceWirelessAlternateManagementInterfaceIPv6Addresses struct {
	Address        string                                                                                   `json:"address,omitempty"`        // The IP address configured for the alternate management interface
	AssignmentMode string                                                                                   `json:"assignmentMode,omitempty"` // The type of address assignment. Either static or dynamic.
	Gateway        string                                                                                   `json:"gateway,omitempty"`        // The gateway address configured for the alternate managment interface
	Nameservers    *RequestWirelessUpdateDeviceWirelessAlternateManagementInterfaceIPv6AddressesNameservers `json:"nameservers,omitempty"`    // The DNS servers settings for this address.
	Prefix         string                                                                                   `json:"prefix,omitempty"`         // The IPv6 prefix length of the IPv6 interface. Required if IPv6 object is included.
	Protocol       string                                                                                   `json:"protocol,omitempty"`       // The IP protocol used for the address
}
type RequestWirelessUpdateDeviceWirelessAlternateManagementInterfaceIPv6AddressesNameservers struct {
	Addresses []string `json:"addresses,omitempty"` // Up to 2 nameserver addresses to use, ordered in priority from highest to lowest priority.
}
type RequestWirelessUpdateDeviceWirelessBluetoothSettings struct {
	Major *int   `json:"major,omitempty"` // Desired major value of the beacon. If the value is set to null it will reset to Dashboard's automatically generated value.
	Minor *int   `json:"minor,omitempty"` // Desired minor value of the beacon. If the value is set to null it will reset to Dashboard's automatically generated value.
	UUID  string `json:"uuid,omitempty"`  // Desired UUID of the beacon. If the value is set to null it will reset to Dashboard's automatically generated value.
}
type RequestWirelessUpdateDeviceWirelessRadioSettings struct {
	FiveGhzSettings    *RequestWirelessUpdateDeviceWirelessRadioSettingsFiveGhzSettings    `json:"fiveGhzSettings,omitempty"`    // Manual radio settings for 5 GHz.
	RfProfileID        string                                                              `json:"rfProfileId,omitempty"`        // The ID of an RF profile to assign to the device. If the value of this parameter is null, the appropriate basic RF profile (indoor or outdoor) will be assigned to the device. Assigning an RF profile will clear ALL manually configured overrides on the device (channel width, channel, power).
	TwoFourGhzSettings *RequestWirelessUpdateDeviceWirelessRadioSettingsTwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"` // Manual radio settings for 2.4 GHz.
}
type RequestWirelessUpdateDeviceWirelessRadioSettingsFiveGhzSettings struct {
	Channel      *int `json:"channel,omitempty"`      // Sets a manual channel for 5 GHz. Can be '36', '40', '44', '48', '52', '56', '60', '64', '100', '104', '108', '112', '116', '120', '124', '128', '132', '136', '140', '144', '149', '153', '157', '161', '165', '169', '173' or '177' or null for using auto channel.
	ChannelWidth *int `json:"channelWidth,omitempty"` // Sets a manual channel for 5 GHz. Can be '0', '20', '40', '80' or '160' or null for using auto channel width.
	TargetPower  *int `json:"targetPower,omitempty"`  // Set a manual target power for 5 GHz. Can be between '8' or '30' or null for using auto power range.
}
type RequestWirelessUpdateDeviceWirelessRadioSettingsTwoFourGhzSettings struct {
	Channel     *int `json:"channel,omitempty"`     // Sets a manual channel for 2.4 GHz. Can be '1', '2', '3', '4', '5', '6', '7', '8', '9', '10', '11', '12', '13' or '14' or null for using auto channel.
	TargetPower *int `json:"targetPower,omitempty"` // Set a manual target power for 2.4 GHz. Can be between '5' or '30' or null for using auto power range.
}
type RequestWirelessUpdateNetworkWirelessAlternateManagementInterface struct {
	AccessPoints *[]RequestWirelessUpdateNetworkWirelessAlternateManagementInterfaceAccessPoints `json:"accessPoints,omitempty"` // Array of access point serial number and IP assignment. Note: accessPoints IP assignment is not applicable for template networks, in other words, do not put 'accessPoints' in the body when updating template networks. Also, an empty 'accessPoints' array will remove all previous static IP assignments
	Enabled      *bool                                                                           `json:"enabled,omitempty"`      // Boolean value to enable or disable alternate management interface
	Protocols    []string                                                                        `json:"protocols,omitempty"`    // Can be one or more of the following values: 'radius', 'snmp', 'syslog' or 'ldap'
	VLANID       *int                                                                            `json:"vlanId,omitempty"`       // Alternate management interface VLAN, must be between 1 and 4094
}
type RequestWirelessUpdateNetworkWirelessAlternateManagementInterfaceAccessPoints struct {
	AlternateManagementIP string `json:"alternateManagementIp,omitempty"` // Wireless alternate management interface device IP. Provide an empty string to remove alternate management IP assignment
	DNS1                  string `json:"dns1,omitempty"`                  // Primary DNS must be in IP format
	DNS2                  string `json:"dns2,omitempty"`                  // Optional secondary DNS must be in IP format
	Gateway               string `json:"gateway,omitempty"`               // Gateway must be in IP format
	Serial                string `json:"serial,omitempty"`                // Serial number of access point to be configured with alternate management IP
	SubnetMask            string `json:"subnetMask,omitempty"`            // Subnet mask must be in IP format
}
type RequestWirelessUpdateNetworkWirelessBilling struct {
	Currency string                                              `json:"currency,omitempty"` // The currency code of this node group's billing plans
	Plans    *[]RequestWirelessUpdateNetworkWirelessBillingPlans `json:"plans,omitempty"`    // Array of billing plans in the node group. (Can configure a maximum of 5)
}
type RequestWirelessUpdateNetworkWirelessBillingPlans struct {
	BandwidthLimits *RequestWirelessUpdateNetworkWirelessBillingPlansBandwidthLimits `json:"bandwidthLimits,omitempty"` // The uplink bandwidth settings for the pricing plan.
	ID              string                                                           `json:"id,omitempty"`              // The id of the pricing plan to update.
	Price           *float64                                                         `json:"price,omitempty"`           // The price of the billing plan.
	TimeLimit       string                                                           `json:"timeLimit,omitempty"`       // The time limit of the pricing plan in minutes. Can be '1 hour', '1 day', '1 week', or '30 days'.
}
type RequestWirelessUpdateNetworkWirelessBillingPlansBandwidthLimits struct {
	LimitDown *int `json:"limitDown,omitempty"` // The maximum download limit (integer, in Kbps). null indicates no limit
	LimitUp   *int `json:"limitUp,omitempty"`   // The maximum upload limit (integer, in Kbps). null indicates no limit
}
type RequestWirelessUpdateNetworkWirelessBluetoothSettings struct {
	AdvertisingEnabled       *bool  `json:"advertisingEnabled,omitempty"`       // Whether APs will advertise beacons.
	Major                    *int   `json:"major,omitempty"`                    // The major number to be used in the beacon identifier. Only valid in 'Non-unique' mode.
	MajorMinorAssignmentMode string `json:"majorMinorAssignmentMode,omitempty"` // The way major and minor number should be assigned to nodes in the network. ('Unique', 'Non-unique')
	Minor                    *int   `json:"minor,omitempty"`                    // The minor number to be used in the beacon identifier. Only valid in 'Non-unique' mode.
	ScanningEnabled          *bool  `json:"scanningEnabled,omitempty"`          // Whether APs will scan for Bluetooth enabled clients.
	UUID                     string `json:"uuid,omitempty"`                     // The UUID to be used in the beacon identifier.
}
type RequestWirelessCreateNetworkWirelessEthernetPortsProfile struct {
	Name     string                                                              `json:"name,omitempty"`     // AP port profile name
	Ports    *[]RequestWirelessCreateNetworkWirelessEthernetPortsProfilePorts    `json:"ports,omitempty"`    // AP ports configuration
	UsbPorts *[]RequestWirelessCreateNetworkWirelessEthernetPortsProfileUsbPorts `json:"usbPorts,omitempty"` // AP usb ports configuration
}
type RequestWirelessCreateNetworkWirelessEthernetPortsProfilePorts struct {
	Enabled    *bool  `json:"enabled,omitempty"`    // AP port enabled
	Name       string `json:"name,omitempty"`       // AP port name
	PskGroupID string `json:"pskGroupId,omitempty"` // AP port PSK Group ID
	SSID       *int   `json:"ssid,omitempty"`       // AP port ssid number
}
type RequestWirelessCreateNetworkWirelessEthernetPortsProfileUsbPorts struct {
	Enabled *bool  `json:"enabled,omitempty"` // AP usb port enabled
	Name    string `json:"name,omitempty"`    // AP usb port name
	SSID    *int   `json:"ssid,omitempty"`    // AP usb port ssid number
}
type RequestWirelessAssignNetworkWirelessEthernetPortsProfiles struct {
	ProfileID string   `json:"profileId,omitempty"` // AP profile ID
	Serials   []string `json:"serials,omitempty"`   // List of AP serials
}
type RequestWirelessSetNetworkWirelessEthernetPortsProfilesDefault struct {
	ProfileID string `json:"profileId,omitempty"` // AP profile ID
}
type RequestWirelessUpdateNetworkWirelessEthernetPortsProfile struct {
	Name     string                                                              `json:"name,omitempty"`     // AP port profile name
	Ports    *[]RequestWirelessUpdateNetworkWirelessEthernetPortsProfilePorts    `json:"ports,omitempty"`    // AP ports configuration
	UsbPorts *[]RequestWirelessUpdateNetworkWirelessEthernetPortsProfileUsbPorts `json:"usbPorts,omitempty"` // AP usb ports configuration
}
type RequestWirelessUpdateNetworkWirelessEthernetPortsProfilePorts struct {
	Enabled    *bool  `json:"enabled,omitempty"`    // AP port enabled
	Name       string `json:"name,omitempty"`       // AP port name
	PskGroupID string `json:"pskGroupId,omitempty"` // AP port PSK Group number
	SSID       *int   `json:"ssid,omitempty"`       // AP port ssid number
}
type RequestWirelessUpdateNetworkWirelessEthernetPortsProfileUsbPorts struct {
	Enabled *bool  `json:"enabled,omitempty"` // AP usb port enabled
	Name    string `json:"name,omitempty"`    // AP usb port name
	SSID    *int   `json:"ssid,omitempty"`    // AP usb port ssid number
}
type RequestWirelessCreateNetworkWirelessRfProfile struct {
	ApBandSettings         *RequestWirelessCreateNetworkWirelessRfProfileApBandSettings     `json:"apBandSettings,omitempty"`         // Settings that will be enabled if selectionType is set to 'ap'.
	BandSelectionType      string                                                           `json:"bandSelectionType,omitempty"`      // Band selection can be set to either 'ssid' or 'ap'. This param is required on creation.
	ClientBalancingEnabled *bool                                                            `json:"clientBalancingEnabled,omitempty"` // Steers client to best available access point. Can be either true or false. Defaults to true.
	FiveGhzSettings        *RequestWirelessCreateNetworkWirelessRfProfileFiveGhzSettings    `json:"fiveGhzSettings,omitempty"`        // Settings related to 5Ghz band
	FlexRadios             *RequestWirelessCreateNetworkWirelessRfProfileFlexRadios         `json:"flexRadios,omitempty"`             // Flex radio settings.
	MinBitrateType         string                                                           `json:"minBitrateType,omitempty"`         // Minimum bitrate can be set to either 'band' or 'ssid'. Defaults to band.
	Name                   string                                                           `json:"name,omitempty"`                   // The name of the new profile. Must be unique. This param is required on creation.
	PerSSIDSettings        *RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings    `json:"perSsidSettings,omitempty"`        // Per-SSID radio settings by number.
	SixGhzSettings         *RequestWirelessCreateNetworkWirelessRfProfileSixGhzSettings     `json:"sixGhzSettings,omitempty"`         // Settings related to 6Ghz band. Only applicable to networks with 6Ghz capable APs
	Transmission           *RequestWirelessCreateNetworkWirelessRfProfileTransmission       `json:"transmission,omitempty"`           // Settings related to radio transmission.
	TwoFourGhzSettings     *RequestWirelessCreateNetworkWirelessRfProfileTwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"`     // Settings related to 2.4Ghz band
}
type RequestWirelessCreateNetworkWirelessRfProfileApBandSettings struct {
	BandOperationMode   string                                                            `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'. Defaults to dual.
	BandSteeringEnabled *bool                                                             `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band. Can be either true or false. Defaults to true.
	Bands               *RequestWirelessCreateNetworkWirelessRfProfileApBandSettingsBands `json:"bands,omitempty"`               // Settings related to all bands
}
type RequestWirelessCreateNetworkWirelessRfProfileApBandSettingsBands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type RequestWirelessCreateNetworkWirelessRfProfileFiveGhzSettings struct {
	ChannelWidth      string `json:"channelWidth,omitempty"`      // Sets channel width (MHz) for 5Ghz band. Can be one of 'auto', '20', '40' or '80'. Defaults to auto.
	MaxPower          *int   `json:"maxPower,omitempty"`          // Sets max power (dBm) of 5Ghz band. Can be integer between 2 and 30. Defaults to 30.
	MinBitrate        *int   `json:"minBitrate,omitempty"`        // Sets min bitrate (Mbps) of 5Ghz band. Can be one of '6', '9', '12', '18', '24', '36', '48' or '54'. Defaults to 12.
	MinPower          *int   `json:"minPower,omitempty"`          // Sets min power (dBm) of 5Ghz band. Can be integer between 2 and 30. Defaults to 8.
	Rxsop             *int   `json:"rxsop,omitempty"`             // The RX-SOP level controls the sensitivity of the radio. It is strongly recommended to use RX-SOP only after consulting a wireless expert. RX-SOP can be configured in the range of -65 to -95 (dBm). A value of null will reset this to the default.
	ValidAutoChannels *[]int `json:"validAutoChannels,omitempty"` // Sets valid auto channels for 5Ghz band. Can be one of '36', '40', '44', '48', '52', '56', '60', '64', '100', '104', '108', '112', '116', '120', '124', '128', '132', '136', '140', '144', '149', '153', '157', '161' or '165'.Defaults to [36, 40, 44, 48, 52, 56, 60, 64, 100, 104, 108, 112, 116, 120, 124, 128, 132, 136, 140, 144, 149, 153, 157, 161, 165].
}
type RequestWirelessCreateNetworkWirelessRfProfileFlexRadios struct {
	ByModel *[]RequestWirelessCreateNetworkWirelessRfProfileFlexRadiosByModel `json:"byModel,omitempty"` // Flex radios by model.
}
type RequestWirelessCreateNetworkWirelessRfProfileFlexRadiosByModel struct {
	Bands []string `json:"bands,omitempty"` // Band to use for each flex radio. For example, ['6'] will set the AP's first flex radio to 6 GHz
	Model string   `json:"model,omitempty"` // Model of the AP
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings struct {
	Status0  *RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings0  `json:"0,omitempty"`  // Settings for SSID 0
	Status1  *RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings1  `json:"1,omitempty"`  // Settings for SSID 1
	Status10 *RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings10 `json:"10,omitempty"` // Settings for SSID 10
	Status11 *RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings11 `json:"11,omitempty"` // Settings for SSID 11
	Status12 *RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings12 `json:"12,omitempty"` // Settings for SSID 12
	Status13 *RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings13 `json:"13,omitempty"` // Settings for SSID 13
	Status14 *RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings14 `json:"14,omitempty"` // Settings for SSID 14
	Status2  *RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings2  `json:"2,omitempty"`  // Settings for SSID 2
	Status3  *RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings3  `json:"3,omitempty"`  // Settings for SSID 3
	Status4  *RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings4  `json:"4,omitempty"`  // Settings for SSID 4
	Status5  *RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings5  `json:"5,omitempty"`  // Settings for SSID 5
	Status6  *RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings6  `json:"6,omitempty"`  // Settings for SSID 6
	Status7  *RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings7  `json:"7,omitempty"`  // Settings for SSID 7
	Status8  *RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings8  `json:"8,omitempty"`  // Settings for SSID 8
	Status9  *RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings9  `json:"9,omitempty"`  // Settings for SSID 9
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings0 struct {
	BandOperationMode   string                                                              `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                               `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings0Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *float64                                                            `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings0Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings1 struct {
	BandOperationMode   string                                                              `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                               `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings1Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *float64                                                            `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings1Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings10 struct {
	BandOperationMode   string                                                               `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                                `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings10Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *float64                                                             `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings10Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings11 struct {
	BandOperationMode   string                                                               `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                                `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings11Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *float64                                                             `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings11Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings12 struct {
	BandOperationMode   string                                                               `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                                `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings12Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *float64                                                             `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings12Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings13 struct {
	BandOperationMode   string                                                               `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                                `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings13Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *float64                                                             `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings13Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4","5", "6", "disabled"
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings14 struct {
	BandOperationMode   string                                                               `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                                `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings14Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *float64                                                             `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings14Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4","5", "6", "disabled"
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings2 struct {
	BandOperationMode   string                                                              `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                               `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings2Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *float64                                                            `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings2Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4","5", "6", "disabled"
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings3 struct {
	BandOperationMode   string                                                              `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                               `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings3Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *float64                                                            `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings3Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4","5", "6", "disabled"
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings4 struct {
	BandOperationMode   string                                                              `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                               `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings4Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *float64                                                            `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings4Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4","5", "6", "disabled"
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings5 struct {
	BandOperationMode   string                                                              `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                               `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings5Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *float64                                                            `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings5Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4","5", "6", "disabled"
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings6 struct {
	BandOperationMode   string                                                              `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                               `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings6Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *float64                                                            `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings6Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4","5", "6", "disabled"
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings7 struct {
	BandOperationMode   string                                                              `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                               `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings7Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *float64                                                            `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings7Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4","5", "6", "disabled"
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings8 struct {
	BandOperationMode   string                                                              `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                               `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings8Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *float64                                                            `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings8Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4","5", "6", "disabled"
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings9 struct {
	BandOperationMode   string                                                              `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                               `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings9Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *float64                                                            `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings9Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4","5", "6", "disabled"
}
type RequestWirelessCreateNetworkWirelessRfProfileSixGhzSettings struct {
	ChannelWidth      string `json:"channelWidth,omitempty"`      // Sets channel width (MHz) for 6Ghz band. Can be one of '0', '20', '40', '80' or '160'. Defaults to 0.
	MaxPower          *int   `json:"maxPower,omitempty"`          // Sets max power (dBm) of 6Ghz band. Can be integer between 2 and 30. Defaults to 30.
	MinBitrate        *int   `json:"minBitrate,omitempty"`        // Sets min bitrate (Mbps) of 6Ghz band. Can be one of '6', '9', '12', '18', '24', '36', '48' or '54'. Defaults to 12.
	MinPower          *int   `json:"minPower,omitempty"`          // Sets min power (dBm) of 6Ghz band. Can be integer between 2 and 30. Defaults to 8.
	Rxsop             *int   `json:"rxsop,omitempty"`             // The RX-SOP level controls the sensitivity of the radio. It is strongly recommended to use RX-SOP only after consulting a wireless expert. RX-SOP can be configured in the range of -65 to -95 (dBm). A value of null will reset this to the default.
	ValidAutoChannels *[]int `json:"validAutoChannels,omitempty"` // Sets valid auto channels for 6Ghz band. Can be one of '1', '5', '9', '13', '17', '21', '25', '29', '33', '37', '41', '45', '49', '53', '57', '61', '65', '69', '73', '77', '81', '85', '89', '93', '97', '101', '105', '109', '113', '117', '121', '125', '129', '133', '137', '141', '145', '149', '153', '157', '161', '165', '169', '173', '177', '181', '185', '189', '193', '197', '201', '205', '209', '213', '217', '221', '225', '229' or '233'.Defaults to [1, 5, 9, 13, 17, 21, 25, 29, 33, 37, 41, 45, 49, 53, 57, 61, 65, 69, 73, 77, 81, 85, 89, 93, 97, 101, 105, 109, 113, 117, 121, 125, 129, 133, 137, 141, 145, 149, 153, 157, 161, 165, 169, 173, 177, 181, 185, 189, 193, 197, 201, 205, 209, 213, 217, 221, 225, 229, 233].
}
type RequestWirelessCreateNetworkWirelessRfProfileTransmission struct {
	Enabled *bool `json:"enabled,omitempty"` // Toggle for radio transmission. When false, radios will not transmit at all.
}
type RequestWirelessCreateNetworkWirelessRfProfileTwoFourGhzSettings struct {
	AxEnabled         *bool    `json:"axEnabled,omitempty"`         // Determines whether ax radio on 2.4Ghz band is on or off. Can be either true or false. If false, we highly recommend disabling band steering. Defaults to true.
	MaxPower          *int     `json:"maxPower,omitempty"`          // Sets max power (dBm) of 2.4Ghz band. Can be integer between 2 and 30. Defaults to 30.
	MinBitrate        *float64 `json:"minBitrate,omitempty"`        // Sets min bitrate (Mbps) of 2.4Ghz band. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'. Defaults to 11.
	MinPower          *int     `json:"minPower,omitempty"`          // Sets min power (dBm) of 2.4Ghz band. Can be integer between 2 and 30. Defaults to 5.
	Rxsop             *int     `json:"rxsop,omitempty"`             // The RX-SOP level controls the sensitivity of the radio. It is strongly recommended to use RX-SOP only after consulting a wireless expert. RX-SOP can be configured in the range of -65 to -95 (dBm). A value of null will reset this to the default.
	ValidAutoChannels *[]int   `json:"validAutoChannels,omitempty"` // Sets valid auto channels for 2.4Ghz band. Can be one of '1', '6' or '11'. Defaults to [1, 6, 11].
}
type RequestWirelessUpdateNetworkWirelessRfProfile struct {
	ApBandSettings         *RequestWirelessUpdateNetworkWirelessRfProfileApBandSettings     `json:"apBandSettings,omitempty"`         // Settings that will be enabled if selectionType is set to 'ap'.
	BandSelectionType      string                                                           `json:"bandSelectionType,omitempty"`      // Band selection can be set to either 'ssid' or 'ap'.
	ClientBalancingEnabled *bool                                                            `json:"clientBalancingEnabled,omitempty"` // Steers client to best available access point. Can be either true or false.
	FiveGhzSettings        *RequestWirelessUpdateNetworkWirelessRfProfileFiveGhzSettings    `json:"fiveGhzSettings,omitempty"`        // Settings related to 5Ghz band
	FlexRadios             *RequestWirelessUpdateNetworkWirelessRfProfileFlexRadios         `json:"flexRadios,omitempty"`             // Flex radio settings.
	MinBitrateType         string                                                           `json:"minBitrateType,omitempty"`         // Minimum bitrate can be set to either 'band' or 'ssid'.
	Name                   string                                                           `json:"name,omitempty"`                   // The name of the new profile. Must be unique.
	PerSSIDSettings        *RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings    `json:"perSsidSettings,omitempty"`        // Per-SSID radio settings by number.
	SixGhzSettings         *RequestWirelessUpdateNetworkWirelessRfProfileSixGhzSettings     `json:"sixGhzSettings,omitempty"`         // Settings related to 6Ghz band. Only applicable to networks with 6Ghz capable APs
	Transmission           *RequestWirelessUpdateNetworkWirelessRfProfileTransmission       `json:"transmission,omitempty"`           // Settings related to radio transmission.
	TwoFourGhzSettings     *RequestWirelessUpdateNetworkWirelessRfProfileTwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"`     // Settings related to 2.4Ghz band
}
type RequestWirelessUpdateNetworkWirelessRfProfileApBandSettings struct {
	BandOperationMode   string                                                            `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                             `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band. Can be either true or false.
	Bands               *RequestWirelessUpdateNetworkWirelessRfProfileApBandSettingsBands `json:"bands,omitempty"`               // Settings related to all bands
}
type RequestWirelessUpdateNetworkWirelessRfProfileApBandSettingsBands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type RequestWirelessUpdateNetworkWirelessRfProfileFiveGhzSettings struct {
	ChannelWidth      string `json:"channelWidth,omitempty"`      // Sets channel width (MHz) for 5Ghz band. Can be one of 'auto', '20', '40' or '80'.
	MaxPower          *int   `json:"maxPower,omitempty"`          // Sets max power (dBm) of 5Ghz band. Can be integer between 2 and 30.
	MinBitrate        *int   `json:"minBitrate,omitempty"`        // Sets min bitrate (Mbps) of 5Ghz band. Can be one of '6', '9', '12', '18', '24', '36', '48' or '54'.
	MinPower          *int   `json:"minPower,omitempty"`          // Sets min power (dBm) of 5Ghz band. Can be integer between 2 and 30.
	Rxsop             *int   `json:"rxsop,omitempty"`             // The RX-SOP level controls the sensitivity of the radio. It is strongly recommended to use RX-SOP only after consulting a wireless expert. RX-SOP can be configured in the range of -65 to -95 (dBm). A value of null will reset this to the default.
	ValidAutoChannels *[]int `json:"validAutoChannels,omitempty"` // Sets valid auto channels for 5Ghz band. Can be one of '36', '40', '44', '48', '52', '56', '60', '64', '100', '104', '108', '112', '116', '120', '124', '128', '132', '136', '140', '144', '149', '153', '157', '161' or '165'.
}
type RequestWirelessUpdateNetworkWirelessRfProfileFlexRadios struct {
	ByModel *[]RequestWirelessUpdateNetworkWirelessRfProfileFlexRadiosByModel `json:"byModel,omitempty"` // Flex radios by model.
}
type RequestWirelessUpdateNetworkWirelessRfProfileFlexRadiosByModel struct {
	Bands []string `json:"bands,omitempty"` // Band to use for each flex radio. For example, ['6'] will set the AP's first flex radio to 6 GHz
	Model string   `json:"model,omitempty"` // Model of the AP
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings struct {
	Status0  *RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings0  `json:"0,omitempty"`  // Settings for SSID 0
	Status1  *RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings1  `json:"1,omitempty"`  // Settings for SSID 1
	Status10 *RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings10 `json:"10,omitempty"` // Settings for SSID 10
	Status11 *RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings11 `json:"11,omitempty"` // Settings for SSID 11
	Status12 *RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings12 `json:"12,omitempty"` // Settings for SSID 12
	Status13 *RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings13 `json:"13,omitempty"` // Settings for SSID 13
	Status14 *RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings14 `json:"14,omitempty"` // Settings for SSID 14
	Status2  *RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings2  `json:"2,omitempty"`  // Settings for SSID 2
	Status3  *RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings3  `json:"3,omitempty"`  // Settings for SSID 3
	Status4  *RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings4  `json:"4,omitempty"`  // Settings for SSID 4
	Status5  *RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings5  `json:"5,omitempty"`  // Settings for SSID 5
	Status6  *RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings6  `json:"6,omitempty"`  // Settings for SSID 6
	Status7  *RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings7  `json:"7,omitempty"`  // Settings for SSID 7
	Status8  *RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings8  `json:"8,omitempty"`  // Settings for SSID 8
	Status9  *RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings9  `json:"9,omitempty"`  // Settings for SSID 9
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings0 struct {
	BandOperationMode   string                                                              `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                               `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings0Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *float64                                                            `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings0Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings1 struct {
	BandOperationMode   string                                                              `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                               `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings1Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *float64                                                            `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings1Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings10 struct {
	BandOperationMode   string                                                               `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                                `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings10Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *float64                                                             `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings10Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings11 struct {
	BandOperationMode   string                                                               `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                                `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings11Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *float64                                                             `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings11Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings12 struct {
	BandOperationMode   string                                                               `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                                `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings12Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *float64                                                             `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings12Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings13 struct {
	BandOperationMode   string                                                               `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                                `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings13Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *float64                                                             `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings13Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings14 struct {
	BandOperationMode   string                                                               `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                                `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings14Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *float64                                                             `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings14Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings2 struct {
	BandOperationMode   string                                                              `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                               `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings2Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *float64                                                            `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings2Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings3 struct {
	BandOperationMode   string                                                              `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                               `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings3Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *float64                                                            `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings3Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings4 struct {
	BandOperationMode   string                                                              `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                               `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings4Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *float64                                                            `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings4Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings5 struct {
	BandOperationMode   string                                                              `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                               `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings5Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *float64                                                            `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings5Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings6 struct {
	BandOperationMode   string                                                              `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                               `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings6Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *float64                                                            `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings6Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings7 struct {
	BandOperationMode   string                                                              `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                               `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings7Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *float64                                                            `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings7Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings8 struct {
	BandOperationMode   string                                                              `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                               `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings8Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *float64                                                            `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings8Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings9 struct {
	BandOperationMode   string                                                              `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandSteeringEnabled *bool                                                               `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	Bands               *RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings9Bands `json:"bands,omitempty"`               // Settings related to all bands
	MinBitrate          *float64                                                            `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings9Bands struct {
	Enabled []string `json:"enabled,omitempty"` // List of enabled bands. Can include ["2.4", "5", "6", "disabled"
}
type RequestWirelessUpdateNetworkWirelessRfProfileSixGhzSettings struct {
	ChannelWidth      string `json:"channelWidth,omitempty"`      // Sets channel width (MHz) for 6Ghz band. Can be one of '0', '20', '40', '80' or '160'.
	MaxPower          *int   `json:"maxPower,omitempty"`          // Sets max power (dBm) of 6Ghz band. Can be integer between 2 and 30.
	MinBitrate        *int   `json:"minBitrate,omitempty"`        // Sets min bitrate (Mbps) of 6Ghz band. Can be one of '6', '9', '12', '18', '24', '36', '48' or '54'.
	MinPower          *int   `json:"minPower,omitempty"`          // Sets min power (dBm) of 6Ghz band. Can be integer between 2 and 30.
	Rxsop             *int   `json:"rxsop,omitempty"`             // The RX-SOP level controls the sensitivity of the radio. It is strongly recommended to use RX-SOP only after consulting a wireless expert. RX-SOP can be configured in the range of -65 to -95 (dBm). A value of null will reset this to the default.
	ValidAutoChannels *[]int `json:"validAutoChannels,omitempty"` // Sets valid auto channels for 6Ghz band. Can be one of '1', '5', '9', '13', '17', '21', '25', '29', '33', '37', '41', '45', '49', '53', '57', '61', '65', '69', '73', '77', '81', '85', '89', '93', '97', '101', '105', '109', '113', '117', '121', '125', '129', '133', '137', '141', '145', '149', '153', '157', '161', '165', '169', '173', '177', '181', '185', '189', '193', '197', '201', '205', '209', '213', '217', '221', '225', '229' or '233'.
}
type RequestWirelessUpdateNetworkWirelessRfProfileTransmission struct {
	Enabled *bool `json:"enabled,omitempty"` // Toggle for radio transmission. When false, radios will not transmit at all.
}
type RequestWirelessUpdateNetworkWirelessRfProfileTwoFourGhzSettings struct {
	AxEnabled         *bool    `json:"axEnabled,omitempty"`         // Determines whether ax radio on 2.4Ghz band is on or off. Can be either true or false. If false, we highly recommend disabling band steering.
	MaxPower          *int     `json:"maxPower,omitempty"`          // Sets max power (dBm) of 2.4Ghz band. Can be integer between 2 and 30.
	MinBitrate        *float64 `json:"minBitrate,omitempty"`        // Sets min bitrate (Mbps) of 2.4Ghz band. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	MinPower          *int     `json:"minPower,omitempty"`          // Sets min power (dBm) of 2.4Ghz band. Can be integer between 2 and 30.
	Rxsop             *int     `json:"rxsop,omitempty"`             // The RX-SOP level controls the sensitivity of the radio. It is strongly recommended to use RX-SOP only after consulting a wireless expert. RX-SOP can be configured in the range of -65 to -95 (dBm). A value of null will reset this to the default.
	ValidAutoChannels *[]int   `json:"validAutoChannels,omitempty"` // Sets valid auto channels for 2.4Ghz band. Can be one of '1', '6' or '11'.
}
type RequestWirelessUpdateNetworkWirelessSettings struct {
	IPv6BridgeEnabled        *bool                                                   `json:"ipv6BridgeEnabled,omitempty"`        // Toggle for enabling or disabling IPv6 bridging in a network (Note: if enabled, SSIDs must also be configured to use bridge mode)
	LedLightsOn              *bool                                                   `json:"ledLightsOn,omitempty"`              // Toggle for enabling or disabling LED lights on all APs in the network (making them run dark)
	LocationAnalyticsEnabled *bool                                                   `json:"locationAnalyticsEnabled,omitempty"` // Toggle for enabling or disabling location analytics for your network
	MeshingEnabled           *bool                                                   `json:"meshingEnabled,omitempty"`           // Toggle for enabling or disabling meshing in a network
	NamedVLANs               *RequestWirelessUpdateNetworkWirelessSettingsNamedVLANs `json:"namedVlans,omitempty"`               // Named VLAN settings for wireless networks.
	Upgradestrategy          string                                                  `json:"upgradeStrategy,omitempty"`          // The upgrade strategy to apply to the network. Must be one of 'minimizeUpgradeTime' or 'minimizeClientDowntime'. Requires firmware version MR 26.8 or higher'
}
type RequestWirelessUpdateNetworkWirelessSettingsNamedVLANs struct {
	PoolDhcpMonitoring *RequestWirelessUpdateNetworkWirelessSettingsNamedVLANsPoolDhcpMonitoring `json:"poolDhcpMonitoring,omitempty"` // Named VLAN Pool DHCP Monitoring settings.
}
type RequestWirelessUpdateNetworkWirelessSettingsNamedVLANsPoolDhcpMonitoring struct {
	Duration *int  `json:"duration,omitempty"` // The duration in minutes that devices will refrain from using dirty VLANs before adding them back to the pool.
	Enabled  *bool `json:"enabled,omitempty"`  // Whether or not devices using named VLAN pools should remove dirty VLANs from the pool, thereby preventing clients from being assigned to VLANs where they would be unable to obtain an IP address via DHCP.
}
type RequestWirelessUpdateNetworkWirelessSSID struct {
	ActiveDirectory                  *RequestWirelessUpdateNetworkWirelessSSIDActiveDirectory           `json:"activeDirectory,omitempty"`                  // The current setting for Active Directory. Only valid if splashPage is 'Password-protected with Active Directory'
	AdultContentFilteringEnabled     *bool                                                              `json:"adultContentFilteringEnabled,omitempty"`     // Boolean indicating whether or not adult content will be blocked
	ApTagsAndVLANIDs                 *[]RequestWirelessUpdateNetworkWirelessSSIDApTagsAndVLANIDs        `json:"apTagsAndVlanIds,omitempty"`                 // The list of tags and VLAN IDs used for VLAN tagging. This param is only valid when the ipAssignmentMode is 'Bridge mode' or 'Layer 3 roaming'
	AuthMode                         string                                                             `json:"authMode,omitempty"`                         // The association control method for the SSID ('open', 'open-enhanced', 'psk', 'open-with-radius', 'open-with-nac', '8021x-meraki', '8021x-nac', '8021x-radius', '8021x-google', '8021x-localradius', 'ipsk-with-radius', 'ipsk-without-radius' or 'ipsk-with-nac')
	AvailabilityTags                 []string                                                           `json:"availabilityTags,omitempty"`                 // Accepts a list of tags for this SSID. If availableOnAllAps is false, then the SSID will only be broadcast by APs with tags matching any of the tags in this list.
	AvailableOnAllAps                *bool                                                              `json:"availableOnAllAps,omitempty"`                // Boolean indicating whether all APs should broadcast the SSID or if it should be restricted to APs matching any availability tags. Can only be false if the SSID has availability tags.
	BandSelection                    string                                                             `json:"bandSelection,omitempty"`                    // The client-serving radio frequencies of this SSID in the default indoor RF profile. ('Dual band operation', '5 GHz band only' or 'Dual band operation with Band Steering')
	ConcentratorNetworkID            string                                                             `json:"concentratorNetworkId,omitempty"`            // The concentrator to use when the ipAssignmentMode is 'Layer 3 roaming with a concentrator' or 'VPN'.
	DefaultVLANID                    *int                                                               `json:"defaultVlanId,omitempty"`                    // The default VLAN ID used for 'all other APs'. This param is only valid when the ipAssignmentMode is 'Bridge mode' or 'Layer 3 roaming'
	DisassociateClientsOnVpnFailover *bool                                                              `json:"disassociateClientsOnVpnFailover,omitempty"` // Disassociate clients when 'VPN' concentrator failover occurs in order to trigger clients to re-associate and generate new DHCP requests. This param is only valid if ipAssignmentMode is 'VPN'.
	DNSRewrite                       *RequestWirelessUpdateNetworkWirelessSSIDDNSRewrite                `json:"dnsRewrite,omitempty"`                       // DNS servers rewrite settings
	Dot11R                           *RequestWirelessUpdateNetworkWirelessSSIDDot11R                    `json:"dot11r,omitempty"`                           // The current setting for 802.11r
	Dot11W                           *RequestWirelessUpdateNetworkWirelessSSIDDot11W                    `json:"dot11w,omitempty"`                           // The current setting for Protected Management Frames (802.11w).
	Enabled                          *bool                                                              `json:"enabled,omitempty"`                          // Whether or not the SSID is enabled
	EncryptionMode                   string                                                             `json:"encryptionMode,omitempty"`                   // The psk encryption mode for the SSID ('wep' or 'wpa'). This param is only valid if the authMode is 'psk'
	EnterpriseAdminAccess            string                                                             `json:"enterpriseAdminAccess,omitempty"`            // Whether or not an SSID is accessible by 'enterprise' administrators ('access disabled' or 'access enabled')
	Gre                              *RequestWirelessUpdateNetworkWirelessSSIDGre                       `json:"gre,omitempty"`                              // Ethernet over GRE settings
	IPAssignmentMode                 string                                                             `json:"ipAssignmentMode,omitempty"`                 // The client IP assignment mode ('NAT mode', 'Bridge mode', 'Layer 3 roaming', 'Ethernet over GRE', 'Layer 3 roaming with a concentrator' or 'VPN')
	LanIsolationEnabled              *bool                                                              `json:"lanIsolationEnabled,omitempty"`              // Boolean indicating whether Layer 2 LAN isolation should be enabled or disabled. Only configurable when ipAssignmentMode is 'Bridge mode'.
	Ldap                             *RequestWirelessUpdateNetworkWirelessSSIDLdap                      `json:"ldap,omitempty"`                             // The current setting for LDAP. Only valid if splashPage is 'Password-protected with LDAP'.
	LocalRadius                      *RequestWirelessUpdateNetworkWirelessSSIDLocalRadius               `json:"localRadius,omitempty"`                      // The current setting for Local Authentication, a built-in RADIUS server on the access point. Only valid if authMode is '8021x-localradius'.
	MandatoryDhcpEnabled             *bool                                                              `json:"mandatoryDhcpEnabled,omitempty"`             // If true, Mandatory DHCP will enforce that clients connecting to this SSID must use the IP address assigned by the DHCP server. Clients who use a static IP address won't be able to associate.
	MinBitrate                       *float64                                                           `json:"minBitrate,omitempty"`                       // The minimum bitrate in Mbps of this SSID in the default indoor RF profile. ('1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54')
	Name                             string                                                             `json:"name,omitempty"`                             // The name of the SSID
	NamedVLANs                       *RequestWirelessUpdateNetworkWirelessSSIDNamedVLANs                `json:"namedVlans,omitempty"`                       // Named VLAN settings.
	Oauth                            *RequestWirelessUpdateNetworkWirelessSSIDOauth                     `json:"oauth,omitempty"`                            // The OAuth settings of this SSID. Only valid if splashPage is 'Google OAuth'.
	PerClientBandwidthLimitDown      *int                                                               `json:"perClientBandwidthLimitDown,omitempty"`      // The download bandwidth limit in Kbps. (0 represents no limit.)
	PerClientBandwidthLimitUp        *int                                                               `json:"perClientBandwidthLimitUp,omitempty"`        // The upload bandwidth limit in Kbps. (0 represents no limit.)
	PerSSIDBandwidthLimitDown        *int                                                               `json:"perSsidBandwidthLimitDown,omitempty"`        // The total download bandwidth limit in Kbps. (0 represents no limit.)
	PerSSIDBandwidthLimitUp          *int                                                               `json:"perSsidBandwidthLimitUp,omitempty"`          // The total upload bandwidth limit in Kbps. (0 represents no limit.)
	Psk                              string                                                             `json:"psk,omitempty"`                              // The passkey for the SSID. This param is only valid if the authMode is 'psk'
	RadiusAccountingEnabled          *bool                                                              `json:"radiusAccountingEnabled,omitempty"`          // Whether or not RADIUS accounting is enabled. This param is only valid if the authMode is 'open-with-radius', '8021x-radius' or 'ipsk-with-radius'
	RadiusAccountingInterimInterval  *int                                                               `json:"radiusAccountingInterimInterval,omitempty"`  // The interval (in seconds) in which accounting information is updated and sent to the RADIUS accounting server.
	RadiusAccountingServers          *[]RequestWirelessUpdateNetworkWirelessSSIDRadiusAccountingServers `json:"radiusAccountingServers,omitempty"`          // The RADIUS accounting 802.1X servers to be used for authentication. This param is only valid if the authMode is 'open-with-radius', '8021x-radius' or 'ipsk-with-radius' and radiusAccountingEnabled is 'true'
	RadiusAttributeForGroupPolicies  string                                                             `json:"radiusAttributeForGroupPolicies,omitempty"`  // Specify the RADIUS attribute used to look up group policies ('Filter-Id', 'Reply-Message', 'Airespace-ACL-Name' or 'Aruba-User-Role'). Access points must receive this attribute in the RADIUS Access-Accept message
	RadiusAuthenticationNasID        string                                                             `json:"radiusAuthenticationNasId,omitempty"`        // The template of the NAS identifier to be used for RADIUS authentication (ex. $NODE_MAC$:$VAP_NUM$).
	RadiusCalledStationID            string                                                             `json:"radiusCalledStationId,omitempty"`            // The template of the called station identifier to be used for RADIUS (ex. $NODE_MAC$:$VAP_NUM$).
	RadiusCoaEnabled                 *bool                                                              `json:"radiusCoaEnabled,omitempty"`                 // If true, Meraki devices will act as a RADIUS Dynamic Authorization Server and will respond to RADIUS Change-of-Authorization and Disconnect messages sent by the RADIUS server.
	RadiusFailoverPolicy             string                                                             `json:"radiusFailoverPolicy,omitempty"`             // This policy determines how authentication requests should be handled in the event that all of the configured RADIUS servers are unreachable ('Deny access' or 'Allow access')
	RadiusFallbackEnabled            *bool                                                              `json:"radiusFallbackEnabled,omitempty"`            // Whether or not higher priority RADIUS servers should be retried after 60 seconds.
	RadiusGuestVLANEnabled           *bool                                                              `json:"radiusGuestVlanEnabled,omitempty"`           // Whether or not RADIUS Guest VLAN is enabled. This param is only valid if the authMode is 'open-with-radius' and addressing mode is not set to 'isolated' or 'nat' mode
	RadiusGuestVLANID                *int                                                               `json:"radiusGuestVlanId,omitempty"`                // VLAN ID of the RADIUS Guest VLAN. This param is only valid if the authMode is 'open-with-radius' and addressing mode is not set to 'isolated' or 'nat' mode
	RadiusLoadBalancingPolicy        string                                                             `json:"radiusLoadBalancingPolicy,omitempty"`        // This policy determines which RADIUS server will be contacted first in an authentication attempt and the ordering of any necessary retry attempts ('Strict priority order' or 'Round robin')
	RadiusOverride                   *bool                                                              `json:"radiusOverride,omitempty"`                   // If true, the RADIUS response can override VLAN tag. This is not valid when ipAssignmentMode is 'NAT mode'.
	RadiusProxyEnabled               *bool                                                              `json:"radiusProxyEnabled,omitempty"`               // If true, Meraki devices will proxy RADIUS messages through the Meraki cloud to the configured RADIUS auth and accounting servers.
	RadiusServerAttemptsLimit        *int                                                               `json:"radiusServerAttemptsLimit,omitempty"`        // The maximum number of transmit attempts after which a RADIUS server is failed over (must be between 1-5).
	RadiusServerTimeout              *int                                                               `json:"radiusServerTimeout,omitempty"`              // The amount of time for which a RADIUS client waits for a reply from the RADIUS server (must be between 1-10 seconds).
	RadiusServers                    *[]RequestWirelessUpdateNetworkWirelessSSIDRadiusServers           `json:"radiusServers,omitempty"`                    // The RADIUS 802.1X servers to be used for authentication. This param is only valid if the authMode is 'open-with-radius', '8021x-radius' or 'ipsk-with-radius'
	RadiusTestingEnabled             *bool                                                              `json:"radiusTestingEnabled,omitempty"`             // If true, Meraki devices will periodically send Access-Request messages to configured RADIUS servers using identity 'meraki_8021x_test' to ensure that the RADIUS servers are reachable.
	SecondaryConcentratorNetworkID   string                                                             `json:"secondaryConcentratorNetworkId,omitempty"`   // The secondary concentrator to use when the ipAssignmentMode is 'VPN'. If configured, the APs will switch to using this concentrator if the primary concentrator is unreachable. This param is optional. ('disabled' represents no secondary concentrator.)
	SpeedBurst                       *RequestWirelessUpdateNetworkWirelessSSIDSpeedBurst                `json:"speedBurst,omitempty"`                       // The SpeedBurst setting for this SSID'
	SplashGuestSponsorDomains        []string                                                           `json:"splashGuestSponsorDomains,omitempty"`        // Array of valid sponsor email domains for sponsored guest splash type.
	SplashPage                       string                                                             `json:"splashPage,omitempty"`                       // The type of splash page for the SSID ('', 'Click-through splash page', 'Billing', 'Password-protected with Meraki RADIUS', 'Password-protected with custom RADIUS', 'Password-protected with Active Directory', 'Password-protected with LDAP', 'SMS authentication', 'Systems Manager Sentry', 'Facebook Wi-Fi', 'Google OAuth', 'Sponsored guest', 'Cisco ISE' or 'Google Apps domain'). This attribute is not supported for template children.
	UseVLANTagging                   *bool                                                              `json:"useVlanTagging,omitempty"`                   // Whether or not traffic should be directed to use specific VLANs. This param is only valid if the ipAssignmentMode is 'Bridge mode' or 'Layer 3 roaming'
	Visible                          *bool                                                              `json:"visible,omitempty"`                          // Boolean indicating whether APs should advertise or hide this SSID. APs will only broadcast this SSID if set to true
	VLANID                           *int                                                               `json:"vlanId,omitempty"`                           // The VLAN ID used for VLAN tagging. This param is only valid when the ipAssignmentMode is 'Layer 3 roaming with a concentrator' or 'VPN'
	WalledGardenEnabled              *bool                                                              `json:"walledGardenEnabled,omitempty"`              // Allow access to a configurable list of IP ranges, which users may access prior to sign-on.
	WalledGardenRanges               []string                                                           `json:"walledGardenRanges,omitempty"`               // Specify your walled garden by entering an array of addresses, ranges using CIDR notation, domain names, and domain wildcards (e.g. '192.168.1.1/24', '192.168.37.10/32', 'www.yahoo.com', '*.google.com']). Meraki's splash page is automatically included in your walled garden.
	WpaEncryptionMode                string                                                             `json:"wpaEncryptionMode,omitempty"`                // The types of WPA encryption. ('WPA1 only', 'WPA1 and WPA2', 'WPA2 only', 'WPA3 Transition Mode', 'WPA3 only' or 'WPA3 192-bit Security')
}
type RequestWirelessUpdateNetworkWirelessSSIDActiveDirectory struct {
	Credentials *RequestWirelessUpdateNetworkWirelessSSIDActiveDirectoryCredentials `json:"credentials,omitempty"` // (Optional) The credentials of the user account to be used by the AP to bind to your Active Directory server. The Active Directory account should have permissions on all your Active Directory servers. Only valid if the splashPage is 'Password-protected with Active Directory'.
	Servers     *[]RequestWirelessUpdateNetworkWirelessSSIDActiveDirectoryServers   `json:"servers,omitempty"`     // The Active Directory servers to be used for authentication.
}
type RequestWirelessUpdateNetworkWirelessSSIDActiveDirectoryCredentials struct {
	LogonName string `json:"logonName,omitempty"` // The logon name of the Active Directory account.
	Password  string `json:"password,omitempty"`  // The password to the Active Directory user account.
}
type RequestWirelessUpdateNetworkWirelessSSIDActiveDirectoryServers struct {
	Host string `json:"host,omitempty"` // IP address (or FQDN) of your Active Directory server.
	Port *int   `json:"port,omitempty"` // (Optional) UDP port the Active Directory server listens on. By default, uses port 3268.
}
type RequestWirelessUpdateNetworkWirelessSSIDApTagsAndVLANIDs struct {
	Tags   []string `json:"tags,omitempty"`   // Array of AP tags
	VLANID *int     `json:"vlanId,omitempty"` // Numerical identifier that is assigned to the VLAN
}
type RequestWirelessUpdateNetworkWirelessSSIDDNSRewrite struct {
	DNSCustomNameservers []string `json:"dnsCustomNameservers,omitempty"` // User specified DNS servers (up to two servers)
	Enabled              *bool    `json:"enabled,omitempty"`              // Boolean indicating whether or not DNS server rewrite is enabled. If disabled, upstream DNS will be used
}
type RequestWirelessUpdateNetworkWirelessSSIDDot11R struct {
	Adaptive *bool `json:"adaptive,omitempty"` // (Optional) Whether 802.11r is adaptive or not.
	Enabled  *bool `json:"enabled,omitempty"`  // Whether 802.11r is enabled or not.
}
type RequestWirelessUpdateNetworkWirelessSSIDDot11W struct {
	Enabled  *bool `json:"enabled,omitempty"`  // Whether 802.11w is enabled or not.
	Required *bool `json:"required,omitempty"` // (Optional) Whether 802.11w is required or not.
}
type RequestWirelessUpdateNetworkWirelessSSIDGre struct {
	Concentrator *RequestWirelessUpdateNetworkWirelessSSIDGreConcentrator `json:"concentrator,omitempty"` // The EoGRE concentrator's settings
	Key          *int                                                     `json:"key,omitempty"`          // Optional numerical identifier that will add the GRE key field to the GRE header. Used to identify an individual traffic flow within a tunnel.
}
type RequestWirelessUpdateNetworkWirelessSSIDGreConcentrator struct {
	Host string `json:"host,omitempty"` // The EoGRE concentrator's IP or FQDN. This param is required when ipAssignmentMode is 'Ethernet over GRE'.
}
type RequestWirelessUpdateNetworkWirelessSSIDLdap struct {
	BaseDistinguishedName string                                                           `json:"baseDistinguishedName,omitempty"` // The base distinguished name of users on the LDAP server.
	Credentials           *RequestWirelessUpdateNetworkWirelessSSIDLdapCredentials         `json:"credentials,omitempty"`           // (Optional) The credentials of the user account to be used by the AP to bind to your LDAP server. The LDAP account should have permissions on all your LDAP servers.
	ServerCaCertificate   *RequestWirelessUpdateNetworkWirelessSSIDLdapServerCaCertificate `json:"serverCaCertificate,omitempty"`   // The CA certificate used to sign the LDAP server's key.
	Servers               *[]RequestWirelessUpdateNetworkWirelessSSIDLdapServers           `json:"servers,omitempty"`               // The LDAP servers to be used for authentication.
}
type RequestWirelessUpdateNetworkWirelessSSIDLdapCredentials struct {
	DistinguishedName string `json:"distinguishedName,omitempty"` // The distinguished name of the LDAP user account (example: cn=user,dc=meraki,dc=com).
	Password          string `json:"password,omitempty"`          // The password of the LDAP user account.
}
type RequestWirelessUpdateNetworkWirelessSSIDLdapServerCaCertificate struct {
	Contents string `json:"contents,omitempty"` // The contents of the CA certificate. Must be in PEM or DER format.
}
type RequestWirelessUpdateNetworkWirelessSSIDLdapServers struct {
	Host string `json:"host,omitempty"` // IP address (or FQDN) of your LDAP server.
	Port *int   `json:"port,omitempty"` // UDP port the LDAP server listens on.
}
type RequestWirelessUpdateNetworkWirelessSSIDLocalRadius struct {
	CacheTimeout              *int                                                                          `json:"cacheTimeout,omitempty"`              // The duration (in seconds) for which LDAP and OCSP lookups are cached.
	CertificateAuthentication *RequestWirelessUpdateNetworkWirelessSSIDLocalRadiusCertificateAuthentication `json:"certificateAuthentication,omitempty"` // The current setting for certificate verification.
	PasswordAuthentication    *RequestWirelessUpdateNetworkWirelessSSIDLocalRadiusPasswordAuthentication    `json:"passwordAuthentication,omitempty"`    // The current setting for password-based authentication.
}
type RequestWirelessUpdateNetworkWirelessSSIDLocalRadiusCertificateAuthentication struct {
	ClientRootCaCertificate *RequestWirelessUpdateNetworkWirelessSSIDLocalRadiusCertificateAuthenticationClientRootCaCertificate `json:"clientRootCaCertificate,omitempty"` // The Client CA Certificate used to sign the client certificate.
	Enabled                 *bool                                                                                                `json:"enabled,omitempty"`                 // Whether or not to use EAP-TLS certificate-based authentication to validate wireless clients.
	OcspResponderURL        string                                                                                               `json:"ocspResponderUrl,omitempty"`        // (Optional) The URL of the OCSP responder to verify client certificate status.
	UseLdap                 *bool                                                                                                `json:"useLdap,omitempty"`                 // Whether or not to verify the certificate with LDAP.
	UseOcsp                 *bool                                                                                                `json:"useOcsp,omitempty"`                 // Whether or not to verify the certificate with OCSP.
}
type RequestWirelessUpdateNetworkWirelessSSIDLocalRadiusCertificateAuthenticationClientRootCaCertificate struct {
	Contents string `json:"contents,omitempty"` // The contents of the Client CA Certificate. Must be in PEM or DER format.
}
type RequestWirelessUpdateNetworkWirelessSSIDLocalRadiusPasswordAuthentication struct {
	Enabled *bool `json:"enabled,omitempty"` // Whether or not to use EAP-TTLS/PAP or PEAP-GTC password-based authentication via LDAP lookup.
}
type RequestWirelessUpdateNetworkWirelessSSIDNamedVLANs struct {
	Radius  *RequestWirelessUpdateNetworkWirelessSSIDNamedVLANsRadius  `json:"radius,omitempty"`  // RADIUS settings. This param is only valid when authMode is 'open-with-radius' and ipAssignmentMode is not 'NAT mode'.
	Tagging *RequestWirelessUpdateNetworkWirelessSSIDNamedVLANsTagging `json:"tagging,omitempty"` // VLAN tagging settings. This param is only valid when ipAssignmentMode is 'Bridge mode' or 'Layer 3 roaming'.
}
type RequestWirelessUpdateNetworkWirelessSSIDNamedVLANsRadius struct {
	GuestVLAN *RequestWirelessUpdateNetworkWirelessSSIDNamedVLANsRadiusGuestVLAN `json:"guestVlan,omitempty"` // Guest VLAN settings. Used to direct traffic to a guest VLAN when none of the RADIUS servers are reachable or a client receives access-reject from the RADIUS server.
}
type RequestWirelessUpdateNetworkWirelessSSIDNamedVLANsRadiusGuestVLAN struct {
	Enabled *bool  `json:"enabled,omitempty"` // Whether or not RADIUS guest named VLAN is enabled.
	Name    string `json:"name,omitempty"`    // RADIUS guest VLAN name.
}
type RequestWirelessUpdateNetworkWirelessSSIDNamedVLANsTagging struct {
	ByApTags        *[]RequestWirelessUpdateNetworkWirelessSSIDNamedVLANsTaggingByApTags `json:"byApTags,omitempty"`        // The list of AP tags and VLAN names used for named VLAN tagging. If an AP has a tag matching one in the list, then traffic on this SSID will be directed to use the VLAN name associated to the tag.
	DefaultVLANName string                                                               `json:"defaultVlanName,omitempty"` // The default VLAN name used to tag traffic in the absence of a matching AP tag.
	Enabled         *bool                                                                `json:"enabled,omitempty"`         // Whether or not traffic should be directed to use specific VLAN names.
}
type RequestWirelessUpdateNetworkWirelessSSIDNamedVLANsTaggingByApTags struct {
	Tags     []string `json:"tags,omitempty"`     // List of AP tags.
	VLANName string   `json:"vlanName,omitempty"` // VLAN name that will be used to tag traffic.
}
type RequestWirelessUpdateNetworkWirelessSSIDOauth struct {
	AllowedDomains []string `json:"allowedDomains,omitempty"` // (Optional) The list of domains allowed access to the network.
}
type RequestWirelessUpdateNetworkWirelessSSIDRadiusAccountingServers struct {
	CaCertificate string `json:"caCertificate,omitempty"` // Certificate used for authorization for the RADSEC Server
	Host          string `json:"host,omitempty"`          // IP address (or FQDN) to which the APs will send RADIUS accounting messages
	Port          *int   `json:"port,omitempty"`          // Port on the RADIUS server that is listening for accounting messages
	RadsecEnabled *bool  `json:"radsecEnabled,omitempty"` // Use RADSEC (TLS over TCP) to connect to this RADIUS accounting server. Requires radiusProxyEnabled.
	Secret        string `json:"secret,omitempty"`        // Shared key used to authenticate messages between the APs and RADIUS server
}
type RequestWirelessUpdateNetworkWirelessSSIDRadiusServers struct {
	CaCertificate            string `json:"caCertificate,omitempty"`            // Certificate used for authorization for the RADSEC Server
	Host                     string `json:"host,omitempty"`                     // IP address (or FQDN) of your RADIUS server
	OpenRoamingCertificateID *int   `json:"openRoamingCertificateId,omitempty"` // The ID of the Openroaming Certificate attached to radius server.
	Port                     *int   `json:"port,omitempty"`                     // UDP port the RADIUS server listens on for Access-requests
	RadsecEnabled            *bool  `json:"radsecEnabled,omitempty"`            // Use RADSEC (TLS over TCP) to connect to this RADIUS server. Requires radiusProxyEnabled.
	Secret                   string `json:"secret,omitempty"`                   // RADIUS client shared secret
}
type RequestWirelessUpdateNetworkWirelessSSIDSpeedBurst struct {
	Enabled *bool `json:"enabled,omitempty"` // Boolean indicating whether or not to allow users to temporarily exceed the bandwidth limit for short periods while still keeping them under the bandwidth limit over time.
}
type RequestWirelessUpdateNetworkWirelessSSIDBonjourForwarding struct {
	Enabled   *bool                                                               `json:"enabled,omitempty"`   // If true, Bonjour forwarding is enabled on this SSID.
	Exception *RequestWirelessUpdateNetworkWirelessSSIDBonjourForwardingException `json:"exception,omitempty"` // Bonjour forwarding exception
	Rules     *[]RequestWirelessUpdateNetworkWirelessSSIDBonjourForwardingRules   `json:"rules,omitempty"`     // List of bonjour forwarding rules.
}
type RequestWirelessUpdateNetworkWirelessSSIDBonjourForwardingException struct {
	Enabled *bool `json:"enabled,omitempty"` // If true, Bonjour forwarding exception is enabled on this SSID. Exception is required to enable L2 isolation and Bonjour forwarding to work together.
}
type RequestWirelessUpdateNetworkWirelessSSIDBonjourForwardingRules struct {
	Description string   `json:"description,omitempty"` // A description for your Bonjour forwarding rule. Optional.
	Services    []string `json:"services,omitempty"`    // A list of Bonjour services. At least one service must be specified. Available services are 'All Services', 'AirPlay', 'AFP', 'BitTorrent', 'FTP', 'iChat', 'iTunes', 'Printers', 'Samba', 'Scanners' and 'SSH'
	VLANID      string   `json:"vlanId,omitempty"`      // The ID of the service VLAN. Required.
}
type RequestWirelessUpdateNetworkWirelessSSIDDeviceTypeGroupPolicies struct {
	DeviceTypePolicies *[]RequestWirelessUpdateNetworkWirelessSSIDDeviceTypeGroupPoliciesDeviceTypePolicies `json:"deviceTypePolicies,omitempty"` // List of device type policies.
	Enabled            *bool                                                                                `json:"enabled,omitempty"`            // If true, the SSID device type group policies are enabled.
}
type RequestWirelessUpdateNetworkWirelessSSIDDeviceTypeGroupPoliciesDeviceTypePolicies struct {
	DevicePolicy  string `json:"devicePolicy,omitempty"`  // The device policy. Can be one of 'Allowed', 'Blocked' or 'Group policy'
	DeviceType    string `json:"deviceType,omitempty"`    // The device type. Can be one of 'Android', 'BlackBerry', 'Chrome OS', 'iPad', 'iPhone', 'iPod', 'Mac OS X', 'Windows', 'Windows Phone', 'B&N Nook' or 'Other OS'
	GroupPolicyID *int   `json:"groupPolicyId,omitempty"` // ID of the group policy object.
}
type RequestWirelessUpdateNetworkWirelessSSIDEapOverride struct {
	EapolKey   *RequestWirelessUpdateNetworkWirelessSSIDEapOverrideEapolKey `json:"eapolKey,omitempty"`   // EAPOL Key settings.
	IDentity   *RequestWirelessUpdateNetworkWirelessSSIDEapOverrideIDentity `json:"identity,omitempty"`   // EAP settings for identity requests.
	MaxRetries *int                                                         `json:"maxRetries,omitempty"` // Maximum number of general EAP retries.
	Timeout    *int                                                         `json:"timeout,omitempty"`    // General EAP timeout in seconds.
}
type RequestWirelessUpdateNetworkWirelessSSIDEapOverrideEapolKey struct {
	Retries     *int `json:"retries,omitempty"`     // Maximum number of EAPOL key retries.
	TimeoutInMs *int `json:"timeoutInMs,omitempty"` // EAPOL Key timeout in milliseconds.
}
type RequestWirelessUpdateNetworkWirelessSSIDEapOverrideIDentity struct {
	Retries *int `json:"retries,omitempty"` // Maximum number of EAP retries.
	Timeout *int `json:"timeout,omitempty"` // EAP timeout in seconds.
}
type RequestWirelessUpdateNetworkWirelessSSIDFirewallL3FirewallRules struct {
	AllowLanAccess *bool                                                                   `json:"allowLanAccess,omitempty"` // Allow wireless client access to local LAN (boolean value - true allows access and false denies access) (optional)
	Rules          *[]RequestWirelessUpdateNetworkWirelessSSIDFirewallL3FirewallRulesRules `json:"rules,omitempty"`          // An ordered array of the firewall rules for this SSID (not including the local LAN access rule or the default rule).
}
type RequestWirelessUpdateNetworkWirelessSSIDFirewallL3FirewallRulesRules struct {
	Comment  string `json:"comment,omitempty"`  // Description of the rule (optional)
	DestCidr string `json:"destCidr,omitempty"` // Comma-separated list of destination IP address(es) (in IP or CIDR notation), fully-qualified domain names (FQDN) or 'any'
	DestPort string `json:"destPort,omitempty"` // Comma-separated list of destination port(s) (integer in the range 1-65535), or 'any'
	Policy   string `json:"policy,omitempty"`   // 'allow' or 'deny' traffic specified by this rule
	Protocol string `json:"protocol,omitempty"` // The type of protocol (must be 'tcp', 'udp', 'icmp', 'icmp6' or 'any')
	IpVer    string `json:"ipVer,omitempty"`    //
}
type RequestWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRules struct {
	Rules *[]RequestWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRulesRules `json:"rules,omitempty"` // An array of L7 firewall rules for this SSID. Rules will get applied in the same order user has specified in request. Empty array will clear the L7 firewall rule configuration.
}
type RequestWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRulesRules struct {
	Policy string      `json:"policy,omitempty"` // 'Deny' traffic specified by this rule
	Type   string      `json:"type,omitempty"`   // Type of the L7 firewall rule. One of: 'application', 'applicationCategory', 'host', 'port', 'ipRange'
	Value  interface{} `json:"value,omitempty"`  // The 'value' of what you want to block. Format of 'value' varies depending on type of the rule. The application categories and application ids can be retrieved from the the 'MX L7 application categories' endpoint. The countries follow the two-letter ISO 3166-1 alpha-2 format.
}

type RequestWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRulesRulesValue struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type RequestWirelessUpdateNetworkWirelessSSIDHotspot20 struct {
	Domains           []string                                                      `json:"domains,omitempty"`           // An array of domain names
	Enabled           *bool                                                         `json:"enabled,omitempty"`           // Whether or not Hotspot 2.0 for this SSID is enabled
	MccMncs           *[]RequestWirelessUpdateNetworkWirelessSSIDHotspot20MccMncs   `json:"mccMncs,omitempty"`           // An array of MCC/MNC pairs
	NaiRealms         *[]RequestWirelessUpdateNetworkWirelessSSIDHotspot20NaiRealms `json:"naiRealms,omitempty"`         // An array of NAI realms
	NetworkAccessType string                                                        `json:"networkAccessType,omitempty"` // The network type of this SSID ('Private network', 'Private network with guest access', 'Chargeable public network', 'Free public network', 'Personal device network', 'Emergency services only network', 'Test or experimental', 'Wildcard')
	Operator          *RequestWirelessUpdateNetworkWirelessSSIDHotspot20Operator    `json:"operator,omitempty"`          // Operator settings for this SSID
	RoamConsortOis    []string                                                      `json:"roamConsortOis,omitempty"`    // An array of roaming consortium OIs (hexadecimal number 3-5 octets in length)
	Venue             *RequestWirelessUpdateNetworkWirelessSSIDHotspot20Venue       `json:"venue,omitempty"`             // Venue settings for this SSID
}
type RequestWirelessUpdateNetworkWirelessSSIDHotspot20MccMncs struct {
	Mcc string `json:"mcc,omitempty"` // MCC value
	Mnc string `json:"mnc,omitempty"` // MNC value
}
type RequestWirelessUpdateNetworkWirelessSSIDHotspot20NaiRealms struct {
	Format  string                                                               `json:"format,omitempty"`  // The format for the realm ('1' or '0')
	Methods *[]RequestWirelessUpdateNetworkWirelessSSIDHotspot20NaiRealmsMethods `json:"methods,omitempty"` // An array of EAP methods for the realm.
	Realm   string                                                               `json:"realm,omitempty"`   // The name of the realm
}
type RequestWirelessUpdateNetworkWirelessSSIDHotspot20NaiRealmsMethods struct {
	AuthenticationTypes *RequestWirelessUpdateNetworkWirelessSSIDHotspot20NaiRealmsMethodsAuthenticationTypes `json:"authenticationTypes,omitempty"` // The authentication types for the method. These should be formatted as an object with the EAP method category in camelcase as the key and the list of types as the value (nonEapInnerAuthentication: Reserved, PAP, CHAP, MSCHAP, MSCHAPV2; eapInnerAuthentication: EAP-TLS, EAP-SIM, EAP-AKA, EAP-TTLS with MSCHAPv2; credentials: SIM, USIM, NFC Secure Element, Hardware Token, Softoken, Certificate, username/password, none, Reserved, Vendor Specific; tunneledEapMethodCredentials: SIM, USIM, NFC Secure Element, Hardware Token, Softoken, Certificate, username/password, Reserved, Anonymous, Vendor Specific)
	ID                  string                                                                                `json:"id,omitempty"`                  // ID of method
}
type RequestWirelessUpdateNetworkWirelessSSIDHotspot20NaiRealmsMethodsAuthenticationTypes interface{}
type RequestWirelessUpdateNetworkWirelessSSIDHotspot20Operator struct {
	Name string `json:"name,omitempty"` // Operator name
}
type RequestWirelessUpdateNetworkWirelessSSIDHotspot20Venue struct {
	Name string `json:"name,omitempty"` // Venue name
	Type string `json:"type,omitempty"` // Venue type ('Unspecified', 'Unspecified Assembly', 'Arena', 'Stadium', 'Passenger Terminal', 'Amphitheater', 'Amusement Park', 'Place of Worship', 'Convention Center', 'Library', 'Museum', 'Restaurant', 'Theater', 'Bar', 'Coffee Shop', 'Zoo or Aquarium', 'Emergency Coordination Center', 'Unspecified Business', 'Doctor or Dentist office', 'Bank', 'Fire Station', 'Police Station', 'Post Office', 'Professional Office', 'Research and Development Facility', 'Attorney Office', 'Unspecified Educational', 'School, Primary', 'School, Secondary', 'University or College', 'Unspecified Factory and Industrial', 'Factory', 'Unspecified Institutional', 'Hospital', 'Long-Term Care Facility', 'Alcohol and Drug Rehabilitation Center', 'Group Home', 'Prison or Jail', 'Unspecified Mercantile', 'Retail Store', 'Grocery Market', 'Automotive Service Station', 'Shopping Mall', 'Gas Station', 'Unspecified Residential', 'Private Residence', 'Hotel or Motel', 'Dormitory', 'Boarding House', 'Unspecified Storage', 'Unspecified Utility and Miscellaneous', 'Unspecified Vehicular', 'Automobile or Truck', 'Airplane', 'Bus', 'Ferry', 'Ship or Boat', 'Train', 'Motor Bike', 'Unspecified Outdoor', 'Muni-mesh Network', 'City Park', 'Rest Area', 'Traffic Control', 'Bus Stop', 'Kiosk')
}
type RequestWirelessCreateNetworkWirelessSSIDIDentityPsk struct {
	ExpiresAt     string `json:"expiresAt,omitempty"`     // Timestamp for when the Identity PSK expires. Will not expire if left blank.
	GroupPolicyID string `json:"groupPolicyId,omitempty"` // The group policy to be applied to clients
	Name          string `json:"name,omitempty"`          // The name of the Identity PSK
	Passphrase    string `json:"passphrase,omitempty"`    // The passphrase for client authentication. If left blank, one will be auto-generated.
}
type RequestWirelessUpdateNetworkWirelessSSIDIDentityPsk struct {
	ExpiresAt     string `json:"expiresAt,omitempty"`     // Timestamp for when the Identity PSK expires, or 'null' to never expire
	GroupPolicyID string `json:"groupPolicyId,omitempty"` // The group policy to be applied to clients
	Name          string `json:"name,omitempty"`          // The name of the Identity PSK
	Passphrase    string `json:"passphrase,omitempty"`    // The passphrase for client authentication
}
type RequestWirelessUpdateNetworkWirelessSSIDSchedules struct {
	Enabled         *bool                                                               `json:"enabled,omitempty"`         // If true, the SSID outage schedule is enabled.
	Ranges          *[]RequestWirelessUpdateNetworkWirelessSSIDSchedulesRanges          `json:"ranges,omitempty"`          // List of outage ranges. Has a start date and time, and end date and time. If this parameter is passed in along with rangesInSeconds parameter, this will take precedence.
	RangesInSeconds *[]RequestWirelessUpdateNetworkWirelessSSIDSchedulesRangesInSeconds `json:"rangesInSeconds,omitempty"` // List of outage ranges in seconds since Sunday at Midnight. Has a start and end. If this parameter is passed in along with the ranges parameter, ranges will take precedence.
}
type RequestWirelessUpdateNetworkWirelessSSIDSchedulesRanges struct {
	EndDay    string `json:"endDay,omitempty"`    // Day of when the outage ends. Can be either full day name, or three letter abbreviation
	EndTime   string `json:"endTime,omitempty"`   // 24 hour time when the outage ends.
	StartDay  string `json:"startDay,omitempty"`  // Day of when the outage starts. Can be either full day name, or three letter abbreviation.
	StartTime string `json:"startTime,omitempty"` // 24 hour time when the outage starts.
}
type RequestWirelessUpdateNetworkWirelessSSIDSchedulesRangesInSeconds struct {
	End   *int `json:"end,omitempty"`   // Seconds since Sunday at midnight when that outage range ends.
	Start *int `json:"start,omitempty"` // Seconds since Sunday at midnight when the outage range starts.
}
type RequestWirelessUpdateNetworkWirelessSSIDSplashSettings struct {
	AllowSimultaneousLogins         *bool                                                                     `json:"allowSimultaneousLogins,omitempty"`         // Whether or not to allow simultaneous logins from different devices.
	Billing                         *RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsBilling            `json:"billing,omitempty"`                         // Details associated with billing splash.
	BlockAllTrafficBeforeSignOn     *bool                                                                     `json:"blockAllTrafficBeforeSignOn,omitempty"`     // How restricted allowing traffic should be. If true, all traffic types are blocked until the splash page is acknowledged. If false, all non-HTTP traffic is allowed before the splash page is acknowledged.
	ControllerDisconnectionBehavior string                                                                    `json:"controllerDisconnectionBehavior,omitempty"` // How login attempts should be handled when the controller is unreachable. Can be either 'open', 'restricted', or 'default'.
	GuestSponsorship                *RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsGuestSponsorship   `json:"guestSponsorship,omitempty"`                // Details associated with guest sponsored splash.
	RedirectURL                     string                                                                    `json:"redirectUrl,omitempty"`                     // The custom redirect URL where the users will go after the splash page.
	SentryEnrollment                *RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsSentryEnrollment   `json:"sentryEnrollment,omitempty"`                // Systems Manager sentry enrollment splash settings.
	SplashImage                     *RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashImage        `json:"splashImage,omitempty"`                     // The image used in the splash page.
	SplashLogo                      *RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashLogo         `json:"splashLogo,omitempty"`                      // The logo used in the splash page.
	SplashPrepaidFront              *RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashPrepaidFront `json:"splashPrepaidFront,omitempty"`              // The prepaid front image used in the splash page.
	SplashTimeout                   *int                                                                      `json:"splashTimeout,omitempty"`                   // Splash timeout in minutes. This will determine how often users will see the splash page.
	SplashURL                       string                                                                    `json:"splashUrl,omitempty"`                       // [optional] The custom splash URL of the click-through splash page. Note that the URL can be configured without necessarily being used. In order to enable the custom URL, see 'useSplashUrl'
	ThemeID                         string                                                                    `json:"themeId,omitempty"`                         // The id of the selected splash theme.
	UseRedirectURL                  *bool                                                                     `json:"useRedirectUrl,omitempty"`                  // The Boolean indicating whether the the user will be redirected to the custom redirect URL after the splash page. A custom redirect URL must be set if this is true.
	UseSplashURL                    *bool                                                                     `json:"useSplashUrl,omitempty"`                    // [optional] Boolean indicating whether the users will be redirected to the custom splash url. A custom splash URL must be set if this is true. Note that depending on your SSID's access control settings, it may not be possible to use the custom splash URL.
	WelcomeMessage                  string                                                                    `json:"welcomeMessage,omitempty"`                  // The welcome message for the users on the splash page.
}
type RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsBilling struct {
	FreeAccess                    *RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsBillingFreeAccess `json:"freeAccess,omitempty"`                    // Details associated with a free access plan with limits.
	PrepaidAccessFastLoginEnabled *bool                                                                    `json:"prepaidAccessFastLoginEnabled,omitempty"` // Whether or not billing uses the fast login prepaid access option.
	ReplyToEmailAddress           string                                                                   `json:"replyToEmailAddress,omitempty"`           // The email address that receives replies from clients.
}
type RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsBillingFreeAccess struct {
	DurationInMinutes *int  `json:"durationInMinutes,omitempty"` // How long a device can use a network for free.
	Enabled           *bool `json:"enabled,omitempty"`           // Whether or not free access is enabled.
}
type RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsGuestSponsorship struct {
	DurationInMinutes        *int  `json:"durationInMinutes,omitempty"`        // Duration in minutes of sponsored guest authorization. Must be between 1 and 60480 (6 weeks)
	GuestCanRequestTimeframe *bool `json:"guestCanRequestTimeframe,omitempty"` // Whether or not guests can specify how much time they are requesting.
}
type RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsSentryEnrollment struct {
	EnforcedSystems       []string                                                                                     `json:"enforcedSystems,omitempty"`       // The system types that the Sentry enforces. Must be included in: 'iOS, 'Android', 'macOS', and 'Windows'.
	Strength              string                                                                                       `json:"strength,omitempty"`              // The strength of the enforcement of selected system types. Must be one of: 'focused', 'click-through', and 'strict'.
	SystemsManagerNetwork *RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsSentryEnrollmentSystemsManagerNetwork `json:"systemsManagerNetwork,omitempty"` // Systems Manager network targeted for sentry enrollment.
}
type RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsSentryEnrollmentSystemsManagerNetwork struct {
	ID string `json:"id,omitempty"` // The network ID of the Systems Manager network.
}
type RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashImage struct {
	Extension string                                                                  `json:"extension,omitempty"` // The extension of the image file.
	Image     *RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashImageImage `json:"image,omitempty"`     // Properties for setting a new image.
	Md5       string                                                                  `json:"md5,omitempty"`       // The MD5 value of the image file. Setting this to null will remove the image from the splash page.
}
type RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashImageImage struct {
	Contents string `json:"contents,omitempty"` // The file contents (a base 64 encoded string) of your new image.
	Format   string `json:"format,omitempty"`   // The format of the encoded contents. Supported formats are 'png', 'gif', and jpg'.
}
type RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashLogo struct {
	Extension string                                                                 `json:"extension,omitempty"` // The extension of the logo file.
	Image     *RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashLogoImage `json:"image,omitempty"`     // Properties for setting a new image.
	Md5       string                                                                 `json:"md5,omitempty"`       // The MD5 value of the logo file. Setting this to null will remove the logo from the splash page.
}
type RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashLogoImage struct {
	Contents string `json:"contents,omitempty"` // The file contents (a base 64 encoded string) of your new logo.
	Format   string `json:"format,omitempty"`   // The format of the encoded contents. Supported formats are 'png', 'gif', and jpg'.
}
type RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashPrepaidFront struct {
	Extension string                                                                         `json:"extension,omitempty"` // The extension of the prepaid front image file.
	Image     *RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashPrepaidFrontImage `json:"image,omitempty"`     // Properties for setting a new image.
	Md5       string                                                                         `json:"md5,omitempty"`       // The MD5 value of the prepaid front image file. Setting this to null will remove the prepaid front from the splash page.
}
type RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashPrepaidFrontImage struct {
	Contents string `json:"contents,omitempty"` // The file contents (a base 64 encoded string) of your new prepaid front.
	Format   string `json:"format,omitempty"`   // The format of the encoded contents. Supported formats are 'png', 'gif', and jpg'.
}
type RequestWirelessUpdateNetworkWirelessSSIDTrafficShapingRules struct {
	DefaultRulesEnabled   *bool                                                               `json:"defaultRulesEnabled,omitempty"`   // Whether default traffic shaping rules are enabled (true) or disabled (false). There are 4 default rules, which can be seen on your network's traffic shaping page. Note that default rules count against the rule limit of 8.
	Rules                 *[]RequestWirelessUpdateNetworkWirelessSSIDTrafficShapingRulesRules `json:"rules,omitempty"`                 //     An array of traffic shaping rules. Rules are applied in the order that     they are specified in. An empty list (or null) means no rules. Note that     you are allowed a maximum of 8 rules.
	TrafficShapingEnabled *bool                                                               `json:"trafficShapingEnabled,omitempty"` // Whether traffic shaping rules are applied to clients on your SSID.
}
type RequestWirelessUpdateNetworkWirelessSSIDTrafficShapingRulesRules struct {
	Definitions              *[]RequestWirelessUpdateNetworkWirelessSSIDTrafficShapingRulesRulesDefinitions            `json:"definitions,omitempty"`              //     A list of objects describing the definitions of your traffic shaping rule. At least one definition is required.
	DscpTagValue             *int                                                                                      `json:"dscpTagValue,omitempty"`             //     The DSCP tag applied by your rule. null means 'Do not change DSCP tag'.     For a list of possible tag values, use the trafficShaping/dscpTaggingOptions endpoint.
	PcpTagValue              *int                                                                                      `json:"pcpTagValue,omitempty"`              //     The PCP tag applied by your rule. Can be 0 (lowest priority) through 7 (highest priority).     null means 'Do not set PCP tag'.
	PerClientBandwidthLimits *RequestWirelessUpdateNetworkWirelessSSIDTrafficShapingRulesRulesPerClientBandwidthLimits `json:"perClientBandwidthLimits,omitempty"` //     An object describing the bandwidth settings for your rule.
}
type RequestWirelessUpdateNetworkWirelessSSIDTrafficShapingRulesRulesDefinitions struct {
	Type  string `json:"type,omitempty"`  // The type of definition. Can be one of 'application', 'applicationCategory', 'host', 'port', 'ipRange' or 'localNet'.
	Value string `json:"value,omitempty"` //     If "type" is 'host', 'port', 'ipRange' or 'localNet', then "value" must be a string, matching either     a hostname (e.g. "somesite.com"), a port (e.g. 8080), or an IP range ("192.1.0.0",     "192.1.0.0/16", or "10.1.0.0/16:80"). 'localNet' also supports CIDR notation, excluding     custom ports.      If "type" is 'application' or 'applicationCategory', then "value" must be an object     with the structure { "id": "meraki:layer7/..." }, where "id" is the application category or     application ID (for a list of IDs for your network, use the trafficShaping/applicationCategories     endpoint).
}
type RequestWirelessUpdateNetworkWirelessSSIDTrafficShapingRulesRulesPerClientBandwidthLimits struct {
	BandwidthLimits *RequestWirelessUpdateNetworkWirelessSSIDTrafficShapingRulesRulesPerClientBandwidthLimitsBandwidthLimits `json:"bandwidthLimits,omitempty"` // The bandwidth limits object, specifying the upload ('limitUp') and download ('limitDown') speed in Kbps. These are only enforced if 'settings' is set to 'custom'.
	Settings        string                                                                                                   `json:"settings,omitempty"`        // How bandwidth limits are applied by your rule. Can be one of 'network default', 'ignore' or 'custom'.
}
type RequestWirelessUpdateNetworkWirelessSSIDTrafficShapingRulesRulesPerClientBandwidthLimitsBandwidthLimits struct {
	LimitDown *int `json:"limitDown,omitempty"` // The maximum download limit (integer, in Kbps).
	LimitUp   *int `json:"limitUp,omitempty"`   // The maximum upload limit (integer, in Kbps).
}
type RequestWirelessUpdateNetworkWirelessSSIDVpn struct {
	Concentrator *RequestWirelessUpdateNetworkWirelessSSIDVpnConcentrator `json:"concentrator,omitempty"` // The VPN concentrator settings for this SSID.
	Failover     *RequestWirelessUpdateNetworkWirelessSSIDVpnFailover     `json:"failover,omitempty"`     // Secondary VPN concentrator settings. This is only used when two VPN concentrators are configured on the SSID.
	SplitTunnel  *RequestWirelessUpdateNetworkWirelessSSIDVpnSplitTunnel  `json:"splitTunnel,omitempty"`  // The VPN split tunnel settings for this SSID.
}
type RequestWirelessUpdateNetworkWirelessSSIDVpnConcentrator struct {
	NetworkID string `json:"networkId,omitempty"` // The NAT ID of the concentrator that should be set.
	VLANID    *int   `json:"vlanId,omitempty"`    // The VLAN that should be tagged for the concentrator.
}
type RequestWirelessUpdateNetworkWirelessSSIDVpnFailover struct {
	HeartbeatInterval *int   `json:"heartbeatInterval,omitempty"` // Idle timer interval in seconds.
	IDleTimeout       *int   `json:"idleTimeout,omitempty"`       // Idle timer timeout in seconds.
	RequestIP         string `json:"requestIp,omitempty"`         // IP addressed reserved on DHCP server where SSID will terminate.
}
type RequestWirelessUpdateNetworkWirelessSSIDVpnSplitTunnel struct {
	Enabled *bool                                                          `json:"enabled,omitempty"` // If true, VPN split tunnel is enabled.
	Rules   *[]RequestWirelessUpdateNetworkWirelessSSIDVpnSplitTunnelRules `json:"rules,omitempty"`   // List of VPN split tunnel rules.
}
type RequestWirelessUpdateNetworkWirelessSSIDVpnSplitTunnelRules struct {
	Comment  string `json:"comment,omitempty"`  // Description for this split tunnel rule (optional).
	DestCidr string `json:"destCidr,omitempty"` // Destination for this split tunnel rule. IP address, fully-qualified domain names (FQDN) or 'any'.
	DestPort string `json:"destPort,omitempty"` // Destination port for this split tunnel rule, (integer in the range 1-65535), or 'any'.
	Policy   string `json:"policy,omitempty"`   // Traffic policy specified for this split tunnel rule, 'allow' or 'deny'.
	Protocol string `json:"protocol,omitempty"` // Protocol for this split tunnel rule.
}

//GetDeviceWirelessBluetoothSettings Return the bluetooth settings for a wireless device
/* Return the bluetooth settings for a wireless device

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-wireless-bluetooth-settings
*/
func (s *WirelessService) GetDeviceWirelessBluetoothSettings(serial string) (*ResponseWirelessGetDeviceWirelessBluetoothSettings, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/wireless/bluetooth/settings"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetDeviceWirelessBluetoothSettings{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceWirelessBluetoothSettings")
	}

	result := response.Result().(*ResponseWirelessGetDeviceWirelessBluetoothSettings)
	return result, response, err

}

//GetDeviceWirelessConnectionStats Aggregated connectivity info for a given AP on this network
/* Aggregated connectivity info for a given AP on this network

@param serial serial path parameter.
@param getDeviceWirelessConnectionStatsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-wireless-connection-stats
*/
func (s *WirelessService) GetDeviceWirelessConnectionStats(serial string, getDeviceWirelessConnectionStatsQueryParams *GetDeviceWirelessConnectionStatsQueryParams) (*ResponseWirelessGetDeviceWirelessConnectionStats, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/wireless/connectionStats"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	queryString, _ := query.Values(getDeviceWirelessConnectionStatsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetDeviceWirelessConnectionStats{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceWirelessConnectionStats")
	}

	result := response.Result().(*ResponseWirelessGetDeviceWirelessConnectionStats)
	return result, response, err

}

//GetDeviceWirelessLatencyStats Aggregated latency info for a given AP on this network
/* Aggregated latency info for a given AP on this network

@param serial serial path parameter.
@param getDeviceWirelessLatencyStatsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-wireless-latency-stats
*/
func (s *WirelessService) GetDeviceWirelessLatencyStats(serial string, getDeviceWirelessLatencyStatsQueryParams *GetDeviceWirelessLatencyStatsQueryParams) (*ResponseWirelessGetDeviceWirelessLatencyStats, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/wireless/latencyStats"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	queryString, _ := query.Values(getDeviceWirelessLatencyStatsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetDeviceWirelessLatencyStats{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceWirelessLatencyStats")
	}

	result := response.Result().(*ResponseWirelessGetDeviceWirelessLatencyStats)
	return result, response, err

}

//GetDeviceWirelessRadioSettings Return the radio settings of a device
/* Return the radio settings of a device

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-wireless-radio-settings
*/
func (s *WirelessService) GetDeviceWirelessRadioSettings(serial string) (*ResponseWirelessGetDeviceWirelessRadioSettings, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/wireless/radio/settings"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetDeviceWirelessRadioSettings{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceWirelessRadioSettings")
	}

	result := response.Result().(*ResponseWirelessGetDeviceWirelessRadioSettings)
	return result, response, err

}

//GetDeviceWirelessStatus Return the SSID statuses of an access point
/* Return the SSID statuses of an access point

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-wireless-status
*/
func (s *WirelessService) GetDeviceWirelessStatus(serial string) (*ResponseWirelessGetDeviceWirelessStatus, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/wireless/status"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetDeviceWirelessStatus{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceWirelessStatus")
	}

	result := response.Result().(*ResponseWirelessGetDeviceWirelessStatus)
	return result, response, err

}

//GetNetworkWirelessAirMarshal List Air Marshal scan results from a network
/* List Air Marshal scan results from a network

@param networkID networkId path parameter. Network ID
@param getNetworkWirelessAirMarshalQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-air-marshal
*/
func (s *WirelessService) GetNetworkWirelessAirMarshal(networkID string, getNetworkWirelessAirMarshalQueryParams *GetNetworkWirelessAirMarshalQueryParams) (*ResponseWirelessGetNetworkWirelessAirMarshal, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/airMarshal"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(getNetworkWirelessAirMarshalQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetNetworkWirelessAirMarshal{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWirelessAirMarshal")
	}

	result := response.Result().(*ResponseWirelessGetNetworkWirelessAirMarshal)
	return result, response, err

}

//GetNetworkWirelessAlternateManagementInterface Return alternate management interface and devices with IP assigned
/* Return alternate management interface and devices with IP assigned

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-alternate-management-interface
*/
func (s *WirelessService) GetNetworkWirelessAlternateManagementInterface(networkID string) (*ResponseWirelessGetNetworkWirelessAlternateManagementInterface, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/alternateManagementInterface"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetNetworkWirelessAlternateManagementInterface{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWirelessAlternateManagementInterface")
	}

	result := response.Result().(*ResponseWirelessGetNetworkWirelessAlternateManagementInterface)
	return result, response, err

}

//GetNetworkWirelessBilling Return the billing settings of this network
/* Return the billing settings of this network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-billing
*/
func (s *WirelessService) GetNetworkWirelessBilling(networkID string) (*ResponseWirelessGetNetworkWirelessBilling, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/billing"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetNetworkWirelessBilling{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWirelessBilling")
	}

	result := response.Result().(*ResponseWirelessGetNetworkWirelessBilling)
	return result, response, err

}

//GetNetworkWirelessBluetoothSettings Return the Bluetooth settings for a network. <a href="https://documentation.meraki.com/MR/Bluetooth/Bluetooth_Low_Energy_(BLE)">Bluetooth settings</a> must be enabled on the network.
/* Return the Bluetooth settings for a network.
Bluetooth settings
 must be enabled on the network.

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-bluetooth-settings
*/
func (s *WirelessService) GetNetworkWirelessBluetoothSettings(networkID string) (*ResponseWirelessGetNetworkWirelessBluetoothSettings, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/bluetooth/settings"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetNetworkWirelessBluetoothSettings{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWirelessBluetoothSettings")
	}

	result := response.Result().(*ResponseWirelessGetNetworkWirelessBluetoothSettings)
	return result, response, err

}

//GetNetworkWirelessChannelUtilizationHistory Return AP channel utilization over time for a device or network client
/* Return AP channel utilization over time for a device or network client

@param networkID networkId path parameter. Network ID
@param getNetworkWirelessChannelUtilizationHistoryQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-channel-utilization-history
*/
func (s *WirelessService) GetNetworkWirelessChannelUtilizationHistory(networkID string, getNetworkWirelessChannelUtilizationHistoryQueryParams *GetNetworkWirelessChannelUtilizationHistoryQueryParams) (*ResponseWirelessGetNetworkWirelessChannelUtilizationHistory, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/channelUtilizationHistory"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(getNetworkWirelessChannelUtilizationHistoryQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetNetworkWirelessChannelUtilizationHistory{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWirelessChannelUtilizationHistory")
	}

	result := response.Result().(*ResponseWirelessGetNetworkWirelessChannelUtilizationHistory)
	return result, response, err

}

//GetNetworkWirelessClientCountHistory Return wireless client counts over time for a network, device, or network client
/* Return wireless client counts over time for a network, device, or network client

@param networkID networkId path parameter. Network ID
@param getNetworkWirelessClientCountHistoryQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-client-count-history
*/
func (s *WirelessService) GetNetworkWirelessClientCountHistory(networkID string, getNetworkWirelessClientCountHistoryQueryParams *GetNetworkWirelessClientCountHistoryQueryParams) (*ResponseWirelessGetNetworkWirelessClientCountHistory, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/clientCountHistory"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(getNetworkWirelessClientCountHistoryQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetNetworkWirelessClientCountHistory{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWirelessClientCountHistory")
	}

	result := response.Result().(*ResponseWirelessGetNetworkWirelessClientCountHistory)
	return result, response, err

}

//GetNetworkWirelessClientsConnectionStats Aggregated connectivity info for this network, grouped by clients
/* Aggregated connectivity info for this network, grouped by clients

@param networkID networkId path parameter. Network ID
@param getNetworkWirelessClientsConnectionStatsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-clients-connection-stats
*/
func (s *WirelessService) GetNetworkWirelessClientsConnectionStats(networkID string, getNetworkWirelessClientsConnectionStatsQueryParams *GetNetworkWirelessClientsConnectionStatsQueryParams) (*ResponseWirelessGetNetworkWirelessClientsConnectionStats, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/clients/connectionStats"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(getNetworkWirelessClientsConnectionStatsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetNetworkWirelessClientsConnectionStats{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWirelessClientsConnectionStats")
	}

	result := response.Result().(*ResponseWirelessGetNetworkWirelessClientsConnectionStats)
	return result, response, err

}

//GetNetworkWirelessClientsLatencyStats Aggregated latency info for this network, grouped by clients
/* Aggregated latency info for this network, grouped by clients

@param networkID networkId path parameter. Network ID
@param getNetworkWirelessClientsLatencyStatsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-clients-latency-stats
*/
func (s *WirelessService) GetNetworkWirelessClientsLatencyStats(networkID string, getNetworkWirelessClientsLatencyStatsQueryParams *GetNetworkWirelessClientsLatencyStatsQueryParams) (*ResponseWirelessGetNetworkWirelessClientsLatencyStats, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/clients/latencyStats"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(getNetworkWirelessClientsLatencyStatsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetNetworkWirelessClientsLatencyStats{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWirelessClientsLatencyStats")
	}

	result := response.Result().(*ResponseWirelessGetNetworkWirelessClientsLatencyStats)
	return result, response, err

}

//GetNetworkWirelessClientConnectionStats Aggregated connectivity info for a given client on this network
/* Aggregated connectivity info for a given client on this network. Clients are identified by their MAC.

@param networkID networkId path parameter. Network ID
@param clientID clientId path parameter. Client ID
@param getNetworkWirelessClientConnectionStatsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-client-connection-stats
*/
func (s *WirelessService) GetNetworkWirelessClientConnectionStats(networkID string, clientID string, getNetworkWirelessClientConnectionStatsQueryParams *GetNetworkWirelessClientConnectionStatsQueryParams) (*ResponseWirelessGetNetworkWirelessClientConnectionStats, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/clients/{clientId}/connectionStats"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{clientId}", fmt.Sprintf("%v", clientID), -1)

	queryString, _ := query.Values(getNetworkWirelessClientConnectionStatsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetNetworkWirelessClientConnectionStats{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWirelessClientConnectionStats")
	}

	result := response.Result().(*ResponseWirelessGetNetworkWirelessClientConnectionStats)
	return result, response, err

}

//GetNetworkWirelessClientConnectivityEvents List the wireless connectivity events for a client within a network in the timespan.
/* List the wireless connectivity events for a client within a network in the timespan.

@param networkID networkId path parameter. Network ID
@param clientID clientId path parameter. Client ID
@param getNetworkWirelessClientConnectivityEventsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-client-connectivity-events
*/
func (s *WirelessService) GetNetworkWirelessClientConnectivityEvents(networkID string, clientID string, getNetworkWirelessClientConnectivityEventsQueryParams *GetNetworkWirelessClientConnectivityEventsQueryParams) (*ResponseWirelessGetNetworkWirelessClientConnectivityEvents, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/clients/{clientId}/connectivityEvents"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{clientId}", fmt.Sprintf("%v", clientID), -1)

	queryString, _ := query.Values(getNetworkWirelessClientConnectivityEventsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetNetworkWirelessClientConnectivityEvents{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWirelessClientConnectivityEvents")
	}

	result := response.Result().(*ResponseWirelessGetNetworkWirelessClientConnectivityEvents)
	return result, response, err

}

//GetNetworkWirelessClientLatencyHistory Return the latency history for a client
/* Return the latency history for a client. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP. The latency data is from a sample of 2% of packets and is grouped into 4 traffic categories: background, best effort, video, voice. Within these categories the sampled packet counters are bucketed by latency in milliseconds.

@param networkID networkId path parameter. Network ID
@param clientID clientId path parameter. Client ID
@param getNetworkWirelessClientLatencyHistoryQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-client-latency-history
*/
func (s *WirelessService) GetNetworkWirelessClientLatencyHistory(networkID string, clientID string, getNetworkWirelessClientLatencyHistoryQueryParams *GetNetworkWirelessClientLatencyHistoryQueryParams) (*ResponseWirelessGetNetworkWirelessClientLatencyHistory, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/clients/{clientId}/latencyHistory"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{clientId}", fmt.Sprintf("%v", clientID), -1)

	queryString, _ := query.Values(getNetworkWirelessClientLatencyHistoryQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetNetworkWirelessClientLatencyHistory{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWirelessClientLatencyHistory")
	}

	result := response.Result().(*ResponseWirelessGetNetworkWirelessClientLatencyHistory)
	return result, response, err

}

//GetNetworkWirelessClientLatencyStats Aggregated latency info for a given client on this network
/* Aggregated latency info for a given client on this network. Clients are identified by their MAC.

@param networkID networkId path parameter. Network ID
@param clientID clientId path parameter. Client ID
@param getNetworkWirelessClientLatencyStatsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-client-latency-stats
*/
func (s *WirelessService) GetNetworkWirelessClientLatencyStats(networkID string, clientID string, getNetworkWirelessClientLatencyStatsQueryParams *GetNetworkWirelessClientLatencyStatsQueryParams) (*ResponseWirelessGetNetworkWirelessClientLatencyStats, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/clients/{clientId}/latencyStats"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{clientId}", fmt.Sprintf("%v", clientID), -1)

	queryString, _ := query.Values(getNetworkWirelessClientLatencyStatsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetNetworkWirelessClientLatencyStats{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWirelessClientLatencyStats")
	}

	result := response.Result().(*ResponseWirelessGetNetworkWirelessClientLatencyStats)
	return result, response, err

}

//GetNetworkWirelessConnectionStats Aggregated connectivity info for this network
/* Aggregated connectivity info for this network

@param networkID networkId path parameter. Network ID
@param getNetworkWirelessConnectionStatsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-connection-stats
*/
func (s *WirelessService) GetNetworkWirelessConnectionStats(networkID string, getNetworkWirelessConnectionStatsQueryParams *GetNetworkWirelessConnectionStatsQueryParams) (*ResponseWirelessGetNetworkWirelessConnectionStats, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/connectionStats"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(getNetworkWirelessConnectionStatsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetNetworkWirelessConnectionStats{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWirelessConnectionStats")
	}

	result := response.Result().(*ResponseWirelessGetNetworkWirelessConnectionStats)
	return result, response, err

}

//GetNetworkWirelessDataRateHistory Return PHY data rates over time for a network, device, or network client
/* Return PHY data rates over time for a network, device, or network client

@param networkID networkId path parameter. Network ID
@param getNetworkWirelessDataRateHistoryQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-data-rate-history
*/
func (s *WirelessService) GetNetworkWirelessDataRateHistory(networkID string, getNetworkWirelessDataRateHistoryQueryParams *GetNetworkWirelessDataRateHistoryQueryParams) (*ResponseWirelessGetNetworkWirelessDataRateHistory, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/dataRateHistory"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(getNetworkWirelessDataRateHistoryQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetNetworkWirelessDataRateHistory{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWirelessDataRateHistory")
	}

	result := response.Result().(*ResponseWirelessGetNetworkWirelessDataRateHistory)
	return result, response, err

}

//GetNetworkWirelessDevicesConnectionStats Aggregated connectivity info for this network, grouped by node
/* Aggregated connectivity info for this network, grouped by node

@param networkID networkId path parameter. Network ID
@param getNetworkWirelessDevicesConnectionStatsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-devices-connection-stats
*/
func (s *WirelessService) GetNetworkWirelessDevicesConnectionStats(networkID string, getNetworkWirelessDevicesConnectionStatsQueryParams *GetNetworkWirelessDevicesConnectionStatsQueryParams) (*ResponseWirelessGetNetworkWirelessDevicesConnectionStats, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/devices/connectionStats"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(getNetworkWirelessDevicesConnectionStatsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetNetworkWirelessDevicesConnectionStats{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWirelessDevicesConnectionStats")
	}

	result := response.Result().(*ResponseWirelessGetNetworkWirelessDevicesConnectionStats)
	return result, response, err

}

//GetNetworkWirelessDevicesLatencyStats Aggregated latency info for this network, grouped by node
/* Aggregated latency info for this network, grouped by node

@param networkID networkId path parameter. Network ID
@param getNetworkWirelessDevicesLatencyStatsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-devices-latency-stats
*/
func (s *WirelessService) GetNetworkWirelessDevicesLatencyStats(networkID string, getNetworkWirelessDevicesLatencyStatsQueryParams *GetNetworkWirelessDevicesLatencyStatsQueryParams) (*ResponseWirelessGetNetworkWirelessDevicesLatencyStats, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/devices/latencyStats"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(getNetworkWirelessDevicesLatencyStatsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetNetworkWirelessDevicesLatencyStats{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWirelessDevicesLatencyStats")
	}

	result := response.Result().(*ResponseWirelessGetNetworkWirelessDevicesLatencyStats)
	return result, response, err

}

//GetNetworkWirelessEthernetPortsProfiles List the AP port profiles for this network
/* List the AP port profiles for this network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-ethernet-ports-profiles
*/
func (s *WirelessService) GetNetworkWirelessEthernetPortsProfiles(networkID string) (*ResponseWirelessGetNetworkWirelessEthernetPortsProfiles, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ethernet/ports/profiles"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetNetworkWirelessEthernetPortsProfiles{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWirelessEthernetPortsProfiles")
	}

	result := response.Result().(*ResponseWirelessGetNetworkWirelessEthernetPortsProfiles)
	return result, response, err

}

//GetNetworkWirelessEthernetPortsProfile Show the AP port profile by ID for this network
/* Show the AP port profile by ID for this network

@param networkID networkId path parameter. Network ID
@param profileID profileId path parameter. Profile ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-ethernet-ports-profile
*/
func (s *WirelessService) GetNetworkWirelessEthernetPortsProfile(networkID string, profileID string) (*ResponseWirelessGetNetworkWirelessEthernetPortsProfile, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ethernet/ports/profiles/{profileId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{profileId}", fmt.Sprintf("%v", profileID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetNetworkWirelessEthernetPortsProfile{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWirelessEthernetPortsProfile")
	}

	result := response.Result().(*ResponseWirelessGetNetworkWirelessEthernetPortsProfile)
	return result, response, err

}

//GetNetworkWirelessFailedConnections List of all failed client connection events on this network in a given time range
/* List of all failed client connection events on this network in a given time range

@param networkID networkId path parameter. Network ID
@param getNetworkWirelessFailedConnectionsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-failed-connections
*/
func (s *WirelessService) GetNetworkWirelessFailedConnections(networkID string, getNetworkWirelessFailedConnectionsQueryParams *GetNetworkWirelessFailedConnectionsQueryParams) (*ResponseWirelessGetNetworkWirelessFailedConnections, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/failedConnections"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(getNetworkWirelessFailedConnectionsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetNetworkWirelessFailedConnections{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWirelessFailedConnections")
	}

	result := response.Result().(*ResponseWirelessGetNetworkWirelessFailedConnections)
	return result, response, err

}

//GetNetworkWirelessLatencyHistory Return average wireless latency over time for a network, device, or network client
/* Return average wireless latency over time for a network, device, or network client

@param networkID networkId path parameter. Network ID
@param getNetworkWirelessLatencyHistoryQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-latency-history
*/
func (s *WirelessService) GetNetworkWirelessLatencyHistory(networkID string, getNetworkWirelessLatencyHistoryQueryParams *GetNetworkWirelessLatencyHistoryQueryParams) (*ResponseWirelessGetNetworkWirelessLatencyHistory, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/latencyHistory"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(getNetworkWirelessLatencyHistoryQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetNetworkWirelessLatencyHistory{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWirelessLatencyHistory")
	}

	result := response.Result().(*ResponseWirelessGetNetworkWirelessLatencyHistory)
	return result, response, err

}

//GetNetworkWirelessLatencyStats Aggregated latency info for this network
/* Aggregated latency info for this network

@param networkID networkId path parameter. Network ID
@param getNetworkWirelessLatencyStatsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-latency-stats
*/
func (s *WirelessService) GetNetworkWirelessLatencyStats(networkID string, getNetworkWirelessLatencyStatsQueryParams *GetNetworkWirelessLatencyStatsQueryParams) (*ResponseWirelessGetNetworkWirelessLatencyStats, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/latencyStats"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(getNetworkWirelessLatencyStatsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetNetworkWirelessLatencyStats{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWirelessLatencyStats")
	}

	result := response.Result().(*ResponseWirelessGetNetworkWirelessLatencyStats)
	return result, response, err

}

//GetNetworkWirelessMeshStatuses List wireless mesh statuses for repeaters
/* List wireless mesh statuses for repeaters

@param networkID networkId path parameter. Network ID
@param getNetworkWirelessMeshStatusesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-mesh-statuses
*/
func (s *WirelessService) GetNetworkWirelessMeshStatuses(networkID string, getNetworkWirelessMeshStatusesQueryParams *GetNetworkWirelessMeshStatusesQueryParams) (*ResponseWirelessGetNetworkWirelessMeshStatuses, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/meshStatuses"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(getNetworkWirelessMeshStatusesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetNetworkWirelessMeshStatuses{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWirelessMeshStatuses")
	}

	result := response.Result().(*ResponseWirelessGetNetworkWirelessMeshStatuses)
	return result, response, err

}

//GetNetworkWirelessRfProfiles List RF profiles for this network
/* List RF profiles for this network

@param networkID networkId path parameter. Network ID
@param getNetworkWirelessRfProfilesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-rf-profiles
*/
func (s *WirelessService) GetNetworkWirelessRfProfiles(networkID string, getNetworkWirelessRfProfilesQueryParams *GetNetworkWirelessRfProfilesQueryParams) (*ResponseWirelessGetNetworkWirelessRfProfiles, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/rfProfiles"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(getNetworkWirelessRfProfilesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetNetworkWirelessRfProfiles{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWirelessRfProfiles")
	}

	result := response.Result().(*ResponseWirelessGetNetworkWirelessRfProfiles)
	return result, response, err

}

//GetNetworkWirelessRfProfile Return a RF profile
/* Return a RF profile

@param networkID networkId path parameter. Network ID
@param rfProfileID rfProfileId path parameter. Rf profile ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-rf-profile
*/
func (s *WirelessService) GetNetworkWirelessRfProfile(networkID string, rfProfileID string) (*ResponseWirelessGetNetworkWirelessRfProfile, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/rfProfiles/{rfProfileId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{rfProfileId}", fmt.Sprintf("%v", rfProfileID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetNetworkWirelessRfProfile{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWirelessRfProfile")
	}

	result := response.Result().(*ResponseWirelessGetNetworkWirelessRfProfile)
	return result, response, err

}

//GetNetworkWirelessSettings Return the wireless settings for a network
/* Return the wireless settings for a network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-settings
*/
func (s *WirelessService) GetNetworkWirelessSettings(networkID string) (*ResponseWirelessGetNetworkWirelessSettings, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/settings"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetNetworkWirelessSettings{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWirelessSettings")
	}

	result := response.Result().(*ResponseWirelessGetNetworkWirelessSettings)
	return result, response, err

}

//GetNetworkWirelessSignalQualityHistory Return signal quality (SNR/RSSI) over time for a device or network client
/* Return signal quality (SNR/RSSI) over time for a device or network client

@param networkID networkId path parameter. Network ID
@param getNetworkWirelessSignalQualityHistoryQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-signal-quality-history
*/
func (s *WirelessService) GetNetworkWirelessSignalQualityHistory(networkID string, getNetworkWirelessSignalQualityHistoryQueryParams *GetNetworkWirelessSignalQualityHistoryQueryParams) (*ResponseWirelessGetNetworkWirelessSignalQualityHistory, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/signalQualityHistory"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(getNetworkWirelessSignalQualityHistoryQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetNetworkWirelessSignalQualityHistory{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWirelessSignalQualityHistory")
	}

	result := response.Result().(*ResponseWirelessGetNetworkWirelessSignalQualityHistory)
	return result, response, err

}

//GetNetworkWirelessSSIDs List the MR SSIDs in a network
/* List the MR SSIDs in a network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-ssids
*/
func (s *WirelessService) GetNetworkWirelessSSIDs(networkID string) (*ResponseWirelessGetNetworkWirelessSSIDs, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ssids"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetNetworkWirelessSSIDs{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWirelessSsids")
	}

	result := response.Result().(*ResponseWirelessGetNetworkWirelessSSIDs)
	return result, response, err

}

//GetNetworkWirelessSSID Return a single MR SSID
/* Return a single MR SSID

@param networkID networkId path parameter. Network ID
@param number number path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-ssid
*/
func (s *WirelessService) GetNetworkWirelessSSID(networkID string, number string) (*ResponseWirelessGetNetworkWirelessSSID, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ssids/{number}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{number}", fmt.Sprintf("%v", number), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetNetworkWirelessSSID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWirelessSsid")
	}

	result := response.Result().(*ResponseWirelessGetNetworkWirelessSSID)
	return result, response, err

}

//GetNetworkWirelessSSIDBonjourForwarding List the Bonjour forwarding setting and rules for the SSID
/* List the Bonjour forwarding setting and rules for the SSID

@param networkID networkId path parameter. Network ID
@param number number path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-ssid-bonjour-forwarding
*/
func (s *WirelessService) GetNetworkWirelessSSIDBonjourForwarding(networkID string, number string) (*ResponseWirelessGetNetworkWirelessSSIDBonjourForwarding, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ssids/{number}/bonjourForwarding"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{number}", fmt.Sprintf("%v", number), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetNetworkWirelessSSIDBonjourForwarding{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWirelessSsidBonjourForwarding")
	}

	result := response.Result().(*ResponseWirelessGetNetworkWirelessSSIDBonjourForwarding)
	return result, response, err

}

//GetNetworkWirelessSSIDDeviceTypeGroupPolicies List the device type group policies for the SSID
/* List the device type group policies for the SSID

@param networkID networkId path parameter. Network ID
@param number number path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-ssid-device-type-group-policies
*/
func (s *WirelessService) GetNetworkWirelessSSIDDeviceTypeGroupPolicies(networkID string, number string) (*ResponseWirelessGetNetworkWirelessSSIDDeviceTypeGroupPolicies, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ssids/{number}/deviceTypeGroupPolicies"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{number}", fmt.Sprintf("%v", number), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetNetworkWirelessSSIDDeviceTypeGroupPolicies{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWirelessSsidDeviceTypeGroupPolicies")
	}

	result := response.Result().(*ResponseWirelessGetNetworkWirelessSSIDDeviceTypeGroupPolicies)
	return result, response, err

}

//GetNetworkWirelessSSIDEapOverride Return the EAP overridden parameters for an SSID
/* Return the EAP overridden parameters for an SSID

@param networkID networkId path parameter. Network ID
@param number number path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-ssid-eap-override
*/
func (s *WirelessService) GetNetworkWirelessSSIDEapOverride(networkID string, number string) (*ResponseWirelessGetNetworkWirelessSSIDEapOverride, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ssids/{number}/eapOverride"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{number}", fmt.Sprintf("%v", number), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetNetworkWirelessSSIDEapOverride{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWirelessSsidEapOverride")
	}

	result := response.Result().(*ResponseWirelessGetNetworkWirelessSSIDEapOverride)
	return result, response, err

}

//GetNetworkWirelessSSIDFirewallL3FirewallRules Return the L3 firewall rules for an SSID on an MR network
/* Return the L3 firewall rules for an SSID on an MR network

@param networkID networkId path parameter. Network ID
@param number number path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-ssid-firewall-l3-firewall-rules
*/
func (s *WirelessService) GetNetworkWirelessSSIDFirewallL3FirewallRules(networkID string, number string) (*ResponseWirelessGetNetworkWirelessSSIDFirewallL3FirewallRules, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ssids/{number}/firewall/l3FirewallRules"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{number}", fmt.Sprintf("%v", number), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetNetworkWirelessSSIDFirewallL3FirewallRules{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWirelessSsidFirewallL3FirewallRules")
	}

	result := response.Result().(*ResponseWirelessGetNetworkWirelessSSIDFirewallL3FirewallRules)
	return result, response, err

}

//GetNetworkWirelessSSIDFirewallL7FirewallRules Return the L7 firewall rules for an SSID on an MR network
/* Return the L7 firewall rules for an SSID on an MR network

@param networkID networkId path parameter. Network ID
@param number number path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-ssid-firewall-l7-firewall-rules
*/
func (s *WirelessService) GetNetworkWirelessSSIDFirewallL7FirewallRules(networkID string, number string) (*ResponseWirelessGetNetworkWirelessSSIDFirewallL7FirewallRules, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ssids/{number}/firewall/l7FirewallRules"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{number}", fmt.Sprintf("%v", number), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetNetworkWirelessSSIDFirewallL7FirewallRules{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWirelessSsidFirewallL7FirewallRules")
	}

	result := response.Result().(*ResponseWirelessGetNetworkWirelessSSIDFirewallL7FirewallRules)
	return result, response, err

}

//GetNetworkWirelessSSIDHotspot20 Return the Hotspot 2.0 settings for an SSID
/* Return the Hotspot 2.0 settings for an SSID

@param networkID networkId path parameter. Network ID
@param number number path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-ssid-hotspot20
*/
func (s *WirelessService) GetNetworkWirelessSSIDHotspot20(networkID string, number string) (*ResponseWirelessGetNetworkWirelessSSIDHotspot20, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ssids/{number}/hotspot20"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{number}", fmt.Sprintf("%v", number), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetNetworkWirelessSSIDHotspot20{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWirelessSsidHotspot20")
	}

	result := response.Result().(*ResponseWirelessGetNetworkWirelessSSIDHotspot20)
	return result, response, err

}

//GetNetworkWirelessSSIDIDentityPsks List all Identity PSKs in a wireless network
/* List all Identity PSKs in a wireless network

@param networkID networkId path parameter. Network ID
@param number number path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-ssid-identity-psks
*/
func (s *WirelessService) GetNetworkWirelessSSIDIDentityPsks(networkID string, number string) (*ResponseWirelessGetNetworkWirelessSSIDIDentityPsks, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ssids/{number}/identityPsks"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{number}", fmt.Sprintf("%v", number), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetNetworkWirelessSSIDIDentityPsks{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWirelessSsidIdentityPsks")
	}

	result := response.Result().(*ResponseWirelessGetNetworkWirelessSSIDIDentityPsks)
	return result, response, err

}

//GetNetworkWirelessSSIDIDentityPsk Return an Identity PSK
/* Return an Identity PSK

@param networkID networkId path parameter. Network ID
@param number number path parameter.
@param identityPskID identityPskId path parameter. Identity psk ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-ssid-identity-psk
*/
func (s *WirelessService) GetNetworkWirelessSSIDIDentityPsk(networkID string, number string, identityPskID string) (*ResponseWirelessGetNetworkWirelessSSIDIDentityPsk, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ssids/{number}/identityPsks/{identityPskId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{number}", fmt.Sprintf("%v", number), -1)
	path = strings.Replace(path, "{identityPskId}", fmt.Sprintf("%v", identityPskID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetNetworkWirelessSSIDIDentityPsk{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWirelessSsidIdentityPsk")
	}

	result := response.Result().(*ResponseWirelessGetNetworkWirelessSSIDIDentityPsk)
	return result, response, err

}

//GetNetworkWirelessSSIDSchedules List the outage schedule for the SSID
/* List the outage schedule for the SSID

@param networkID networkId path parameter. Network ID
@param number number path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-ssid-schedules
*/
func (s *WirelessService) GetNetworkWirelessSSIDSchedules(networkID string, number string) (*ResponseWirelessGetNetworkWirelessSSIDSchedules, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ssids/{number}/schedules"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{number}", fmt.Sprintf("%v", number), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetNetworkWirelessSSIDSchedules{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWirelessSsidSchedules")
	}

	result := response.Result().(*ResponseWirelessGetNetworkWirelessSSIDSchedules)
	return result, response, err

}

//GetNetworkWirelessSSIDSplashSettings Display the splash page settings for the given SSID
/* Display the splash page settings for the given SSID

@param networkID networkId path parameter. Network ID
@param number number path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-ssid-splash-settings
*/
func (s *WirelessService) GetNetworkWirelessSSIDSplashSettings(networkID string, number string) (*ResponseWirelessGetNetworkWirelessSSIDSplashSettings, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ssids/{number}/splash/settings"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{number}", fmt.Sprintf("%v", number), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetNetworkWirelessSSIDSplashSettings{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWirelessSsidSplashSettings")
	}

	result := response.Result().(*ResponseWirelessGetNetworkWirelessSSIDSplashSettings)
	return result, response, err

}

//GetNetworkWirelessSSIDTrafficShapingRules Display the traffic shaping settings for a SSID on an MR network
/* Display the traffic shaping settings for a SSID on an MR network

@param networkID networkId path parameter. Network ID
@param number number path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-ssid-traffic-shaping-rules
*/
func (s *WirelessService) GetNetworkWirelessSSIDTrafficShapingRules(networkID string, number string) (*ResponseWirelessGetNetworkWirelessSSIDTrafficShapingRules, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ssids/{number}/trafficShaping/rules"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{number}", fmt.Sprintf("%v", number), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetNetworkWirelessSSIDTrafficShapingRules{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWirelessSsidTrafficShapingRules")
	}

	result := response.Result().(*ResponseWirelessGetNetworkWirelessSSIDTrafficShapingRules)
	return result, response, err

}

//GetNetworkWirelessSSIDVpn List the VPN settings for the SSID.
/* List the VPN settings for the SSID.

@param networkID networkId path parameter. Network ID
@param number number path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-ssid-vpn
*/
func (s *WirelessService) GetNetworkWirelessSSIDVpn(networkID string, number string) (*ResponseWirelessGetNetworkWirelessSSIDVpn, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ssids/{number}/vpn"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{number}", fmt.Sprintf("%v", number), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetNetworkWirelessSSIDVpn{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWirelessSsidVpn")
	}

	result := response.Result().(*ResponseWirelessGetNetworkWirelessSSIDVpn)
	return result, response, err

}

//GetNetworkWirelessUsageHistory Return AP usage over time for a device or network client
/* Return AP usage over time for a device or network client

@param networkID networkId path parameter. Network ID
@param getNetworkWirelessUsageHistoryQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-usage-history
*/
func (s *WirelessService) GetNetworkWirelessUsageHistory(networkID string, getNetworkWirelessUsageHistoryQueryParams *GetNetworkWirelessUsageHistoryQueryParams) (*ResponseWirelessGetNetworkWirelessUsageHistory, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/usageHistory"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(getNetworkWirelessUsageHistoryQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetNetworkWirelessUsageHistory{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkWirelessUsageHistory")
	}

	result := response.Result().(*ResponseWirelessGetNetworkWirelessUsageHistory)
	return result, response, err

}

//GetOrganizationWirelessDevicesChannelUtilizationByDevice Get average channel utilization for all bands in a network, split by AP
/* Get average channel utilization for all bands in a network, split by AP

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessDevicesChannelUtilizationByDeviceQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-wireless-devices-channel-utilization-by-device
*/
func (s *WirelessService) GetOrganizationWirelessDevicesChannelUtilizationByDevice(organizationID string, getOrganizationWirelessDevicesChannelUtilizationByDeviceQueryParams *GetOrganizationWirelessDevicesChannelUtilizationByDeviceQueryParams) (*ResponseWirelessGetOrganizationWirelessDevicesChannelUtilizationByDevice, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wireless/devices/channelUtilization/byDevice"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessDevicesChannelUtilizationByDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetOrganizationWirelessDevicesChannelUtilizationByDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWirelessDevicesChannelUtilizationByDevice")
	}

	result := response.Result().(*ResponseWirelessGetOrganizationWirelessDevicesChannelUtilizationByDevice)
	return result, response, err

}

//GetOrganizationWirelessDevicesChannelUtilizationByNetwork Get average channel utilization across all bands for all networks in the organization
/* Get average channel utilization across all bands for all networks in the organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessDevicesChannelUtilizationByNetworkQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-wireless-devices-channel-utilization-by-network
*/
func (s *WirelessService) GetOrganizationWirelessDevicesChannelUtilizationByNetwork(organizationID string, getOrganizationWirelessDevicesChannelUtilizationByNetworkQueryParams *GetOrganizationWirelessDevicesChannelUtilizationByNetworkQueryParams) (*ResponseWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetwork, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wireless/devices/channelUtilization/byNetwork"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessDevicesChannelUtilizationByNetworkQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetwork{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWirelessDevicesChannelUtilizationByNetwork")
	}

	result := response.Result().(*ResponseWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetwork)
	return result, response, err

}

//GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval Get a time-series of average channel utilization for all bands, segmented by device.
/* Get a time-series of average channel utilization for all bands, segmented by device.

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-wireless-devices-channel-utilization-history-by-device-by-interval
*/
func (s *WirelessService) GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval(organizationID string, getOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalQueryParams *GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalQueryParams) (*ResponseWirelessGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wireless/devices/channelUtilization/history/byDevice/byInterval"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval")
	}

	result := response.Result().(*ResponseWirelessGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval)
	return result, response, err

}

//GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval Get a time-series of average channel utilization for all bands
/* Get a time-series of average channel utilization for all bands

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-wireless-devices-channel-utilization-history-by-network-by-interval
*/
func (s *WirelessService) GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval(organizationID string, getOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalQueryParams *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalQueryParams) (*ResponseWirelessGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wireless/devices/channelUtilization/history/byNetwork/byInterval"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval")
	}

	result := response.Result().(*ResponseWirelessGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval)
	return result, response, err

}

//GetOrganizationWirelessDevicesEthernetStatuses List the most recent Ethernet link speed, duplex, aggregation and power mode and status information for wireless devices.
/* List the most recent Ethernet link speed, duplex, aggregation and power mode and status information for wireless devices.

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessDevicesEthernetStatusesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-wireless-devices-ethernet-statuses
*/
func (s *WirelessService) GetOrganizationWirelessDevicesEthernetStatuses(organizationID string, getOrganizationWirelessDevicesEthernetStatusesQueryParams *GetOrganizationWirelessDevicesEthernetStatusesQueryParams) (*ResponseWirelessGetOrganizationWirelessDevicesEthernetStatuses, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wireless/devices/ethernet/statuses"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessDevicesEthernetStatusesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetOrganizationWirelessDevicesEthernetStatuses{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWirelessDevicesEthernetStatuses")
	}

	result := response.Result().(*ResponseWirelessGetOrganizationWirelessDevicesEthernetStatuses)
	return result, response, err

}

//GetOrganizationWirelessDevicesPacketLossByClient Get average packet loss for the given timespan for all clients in the organization.
/* Get average packet loss for the given timespan for all clients in the organization.

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessDevicesPacketLossByClientQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-wireless-devices-packet-loss-by-client
*/
func (s *WirelessService) GetOrganizationWirelessDevicesPacketLossByClient(organizationID string, getOrganizationWirelessDevicesPacketLossByClientQueryParams *GetOrganizationWirelessDevicesPacketLossByClientQueryParams) (*ResponseWirelessGetOrganizationWirelessDevicesPacketLossByClient, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wireless/devices/packetLoss/byClient"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessDevicesPacketLossByClientQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetOrganizationWirelessDevicesPacketLossByClient{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWirelessDevicesPacketLossByClient")
	}

	result := response.Result().(*ResponseWirelessGetOrganizationWirelessDevicesPacketLossByClient)
	return result, response, err

}

//GetOrganizationWirelessDevicesPacketLossByDevice Get average packet loss for the given timespan for all devices in the organization
/* Get average packet loss for the given timespan for all devices in the organization. Does not include device's own traffic.

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessDevicesPacketLossByDeviceQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-wireless-devices-packet-loss-by-device
*/
func (s *WirelessService) GetOrganizationWirelessDevicesPacketLossByDevice(organizationID string, getOrganizationWirelessDevicesPacketLossByDeviceQueryParams *GetOrganizationWirelessDevicesPacketLossByDeviceQueryParams) (*ResponseWirelessGetOrganizationWirelessDevicesPacketLossByDevice, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wireless/devices/packetLoss/byDevice"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessDevicesPacketLossByDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetOrganizationWirelessDevicesPacketLossByDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWirelessDevicesPacketLossByDevice")
	}

	result := response.Result().(*ResponseWirelessGetOrganizationWirelessDevicesPacketLossByDevice)
	return result, response, err

}

//GetOrganizationWirelessDevicesPacketLossByNetwork Get average packet loss for the given timespan for all networks in the organization.
/* Get average packet loss for the given timespan for all networks in the organization.

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessDevicesPacketLossByNetworkQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-wireless-devices-packet-loss-by-network
*/
func (s *WirelessService) GetOrganizationWirelessDevicesPacketLossByNetwork(organizationID string, getOrganizationWirelessDevicesPacketLossByNetworkQueryParams *GetOrganizationWirelessDevicesPacketLossByNetworkQueryParams) (*ResponseWirelessGetOrganizationWirelessDevicesPacketLossByNetwork, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wireless/devices/packetLoss/byNetwork"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessDevicesPacketLossByNetworkQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetOrganizationWirelessDevicesPacketLossByNetwork{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWirelessDevicesPacketLossByNetwork")
	}

	result := response.Result().(*ResponseWirelessGetOrganizationWirelessDevicesPacketLossByNetwork)
	return result, response, err

}

//CreateNetworkWirelessEthernetPortsProfile Create an AP port profile
/* Create an AP port profile

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-network-wireless-ethernet-ports-profile
*/

func (s *WirelessService) CreateNetworkWirelessEthernetPortsProfile(networkID string, requestWirelessCreateNetworkWirelessEthernetPortsProfile *RequestWirelessCreateNetworkWirelessEthernetPortsProfile) (*ResponseWirelessCreateNetworkWirelessEthernetPortsProfile, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ethernet/ports/profiles"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessCreateNetworkWirelessEthernetPortsProfile).
		SetResult(&ResponseWirelessCreateNetworkWirelessEthernetPortsProfile{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateNetworkWirelessEthernetPortsProfile")
	}

	result := response.Result().(*ResponseWirelessCreateNetworkWirelessEthernetPortsProfile)
	return result, response, err

}

//AssignNetworkWirelessEthernetPortsProfiles Assign AP port profile to list of APs
/* Assign AP port profile to list of APs

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!assign-network-wireless-ethernet-ports-profiles
*/

func (s *WirelessService) AssignNetworkWirelessEthernetPortsProfiles(networkID string, requestWirelessAssignNetworkWirelessEthernetPortsProfiles *RequestWirelessAssignNetworkWirelessEthernetPortsProfiles) (*ResponseWirelessAssignNetworkWirelessEthernetPortsProfiles, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ethernet/ports/profiles/assign"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessAssignNetworkWirelessEthernetPortsProfiles).
		SetResult(&ResponseWirelessAssignNetworkWirelessEthernetPortsProfiles{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation AssignNetworkWirelessEthernetPortsProfiles")
	}

	result := response.Result().(*ResponseWirelessAssignNetworkWirelessEthernetPortsProfiles)
	return result, response, err

}

//SetNetworkWirelessEthernetPortsProfilesDefault Set the AP port profile to be default for this network
/* Set the AP port profile to be default for this network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!set-network-wireless-ethernet-ports-profiles-default
*/

func (s *WirelessService) SetNetworkWirelessEthernetPortsProfilesDefault(networkID string, requestWirelessSetNetworkWirelessEthernetPortsProfilesDefault *RequestWirelessSetNetworkWirelessEthernetPortsProfilesDefault) (*ResponseWirelessSetNetworkWirelessEthernetPortsProfilesDefault, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ethernet/ports/profiles/setDefault"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessSetNetworkWirelessEthernetPortsProfilesDefault).
		SetResult(&ResponseWirelessSetNetworkWirelessEthernetPortsProfilesDefault{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation SetNetworkWirelessEthernetPortsProfilesDefault")
	}

	result := response.Result().(*ResponseWirelessSetNetworkWirelessEthernetPortsProfilesDefault)
	return result, response, err

}

//CreateNetworkWirelessRfProfile Creates new RF profile for this network
/* Creates new RF profile for this network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-network-wireless-rf-profile
*/

func (s *WirelessService) CreateNetworkWirelessRfProfile(networkID string, requestWirelessCreateNetworkWirelessRfProfile *RequestWirelessCreateNetworkWirelessRfProfile) (*ResponseWirelessCreateNetworkWirelessRfProfile, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/rfProfiles"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessCreateNetworkWirelessRfProfile).
		SetResult(&ResponseWirelessCreateNetworkWirelessRfProfile{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateNetworkWirelessRfProfile")
	}

	result := response.Result().(*ResponseWirelessCreateNetworkWirelessRfProfile)
	return result, response, err

}

//CreateNetworkWirelessSSIDIDentityPsk Create an Identity PSK
/* Create an Identity PSK

@param networkID networkId path parameter. Network ID
@param number number path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-network-wireless-ssid-identity-psk
*/

func (s *WirelessService) CreateNetworkWirelessSSIDIDentityPsk(networkID string, number string, requestWirelessCreateNetworkWirelessSsidIdentityPsk *RequestWirelessCreateNetworkWirelessSSIDIDentityPsk) (*ResponseWirelessCreateNetworkWirelessSSIDIDentityPsk, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ssids/{number}/identityPsks"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{number}", fmt.Sprintf("%v", number), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessCreateNetworkWirelessSsidIdentityPsk).
		SetResult(&ResponseWirelessCreateNetworkWirelessSSIDIDentityPsk{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateNetworkWirelessSsidIdentityPsk")
	}

	result := response.Result().(*ResponseWirelessCreateNetworkWirelessSSIDIDentityPsk)
	return result, response, err

}

//UpdateDeviceWirelessAlternateManagementInterfaceIPv6 Update alternate management interface IPv6 address
/* Update alternate management interface IPv6 address

@param serial serial path parameter.
*/
func (s *WirelessService) UpdateDeviceWirelessAlternateManagementInterfaceIPv6(serial string, requestWirelessUpdateDeviceWirelessAlternateManagementInterfaceIpv6 *RequestWirelessUpdateDeviceWirelessAlternateManagementInterfaceIPv6) (*ResponseWirelessUpdateDeviceWirelessAlternateManagementInterfaceIPv6, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/wireless/alternateManagementInterface/ipv6"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateDeviceWirelessAlternateManagementInterfaceIpv6).
		SetResult(&ResponseWirelessUpdateDeviceWirelessAlternateManagementInterfaceIPv6{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateDeviceWirelessAlternateManagementInterfaceIpv6")
	}

	result := response.Result().(*ResponseWirelessUpdateDeviceWirelessAlternateManagementInterfaceIPv6)
	return result, response, err

}

//UpdateDeviceWirelessBluetoothSettings Update the bluetooth settings for a wireless device
/* Update the bluetooth settings for a wireless device

@param serial serial path parameter.
*/
func (s *WirelessService) UpdateDeviceWirelessBluetoothSettings(serial string, requestWirelessUpdateDeviceWirelessBluetoothSettings *RequestWirelessUpdateDeviceWirelessBluetoothSettings) (*ResponseWirelessUpdateDeviceWirelessBluetoothSettings, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/wireless/bluetooth/settings"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateDeviceWirelessBluetoothSettings).
		SetResult(&ResponseWirelessUpdateDeviceWirelessBluetoothSettings{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateDeviceWirelessBluetoothSettings")
	}

	result := response.Result().(*ResponseWirelessUpdateDeviceWirelessBluetoothSettings)
	return result, response, err

}

//UpdateDeviceWirelessRadioSettings Update the radio settings of a device
/* Update the radio settings of a device

@param serial serial path parameter.
*/
func (s *WirelessService) UpdateDeviceWirelessRadioSettings(serial string, requestWirelessUpdateDeviceWirelessRadioSettings *RequestWirelessUpdateDeviceWirelessRadioSettings) (*resty.Response, error) {
	path := "/api/v1/devices/{serial}/wireless/radio/settings"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateDeviceWirelessRadioSettings).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateDeviceWirelessRadioSettings")
	}

	return response, err

}

//UpdateNetworkWirelessAlternateManagementInterface Update alternate management interface and device static IP
/* Update alternate management interface and device static IP

@param networkID networkId path parameter. Network ID
*/
func (s *WirelessService) UpdateNetworkWirelessAlternateManagementInterface(networkID string, requestWirelessUpdateNetworkWirelessAlternateManagementInterface *RequestWirelessUpdateNetworkWirelessAlternateManagementInterface) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/alternateManagementInterface"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateNetworkWirelessAlternateManagementInterface).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkWirelessAlternateManagementInterface")
	}

	return response, err

}

//UpdateNetworkWirelessBilling Update the billing settings
/* Update the billing settings

@param networkID networkId path parameter. Network ID
*/
func (s *WirelessService) UpdateNetworkWirelessBilling(networkID string, requestWirelessUpdateNetworkWirelessBilling *RequestWirelessUpdateNetworkWirelessBilling) (*ResponseWirelessUpdateNetworkWirelessBilling, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/billing"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateNetworkWirelessBilling).
		SetResult(&ResponseWirelessUpdateNetworkWirelessBilling{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkWirelessBilling")
	}

	result := response.Result().(*ResponseWirelessUpdateNetworkWirelessBilling)
	return result, response, err

}

//UpdateNetworkWirelessBluetoothSettings Update the Bluetooth settings for a network
/* Update the Bluetooth settings for a network. See the docs page for
Bluetooth settings
.

@param networkID networkId path parameter. Network ID
*/
func (s *WirelessService) UpdateNetworkWirelessBluetoothSettings(networkID string, requestWirelessUpdateNetworkWirelessBluetoothSettings *RequestWirelessUpdateNetworkWirelessBluetoothSettings) (*ResponseWirelessUpdateNetworkWirelessBluetoothSettings, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/bluetooth/settings"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateNetworkWirelessBluetoothSettings).
		SetResult(&ResponseWirelessUpdateNetworkWirelessBluetoothSettings{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkWirelessBluetoothSettings")
	}

	result := response.Result().(*ResponseWirelessUpdateNetworkWirelessBluetoothSettings)
	return result, response, err

}

//UpdateNetworkWirelessEthernetPortsProfile Update the AP port profile by ID for this network
/* Update the AP port profile by ID for this network

@param networkID networkId path parameter. Network ID
@param profileID profileId path parameter. Profile ID
*/
func (s *WirelessService) UpdateNetworkWirelessEthernetPortsProfile(networkID string, profileID string, requestWirelessUpdateNetworkWirelessEthernetPortsProfile *RequestWirelessUpdateNetworkWirelessEthernetPortsProfile) (*ResponseWirelessUpdateNetworkWirelessEthernetPortsProfile, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ethernet/ports/profiles/{profileId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{profileId}", fmt.Sprintf("%v", profileID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateNetworkWirelessEthernetPortsProfile).
		SetResult(&ResponseWirelessUpdateNetworkWirelessEthernetPortsProfile{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkWirelessEthernetPortsProfile")
	}

	result := response.Result().(*ResponseWirelessUpdateNetworkWirelessEthernetPortsProfile)
	return result, response, err

}

//UpdateNetworkWirelessRfProfile Updates specified RF profile for this network
/* Updates specified RF profile for this network

@param networkID networkId path parameter. Network ID
@param rfProfileID rfProfileId path parameter. Rf profile ID
*/
func (s *WirelessService) UpdateNetworkWirelessRfProfile(networkID string, rfProfileID string, requestWirelessUpdateNetworkWirelessRfProfile *RequestWirelessUpdateNetworkWirelessRfProfile) (*ResponseWirelessUpdateNetworkWirelessRfProfile, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/rfProfiles/{rfProfileId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{rfProfileId}", fmt.Sprintf("%v", rfProfileID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateNetworkWirelessRfProfile).
		SetResult(&ResponseWirelessUpdateNetworkWirelessRfProfile{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkWirelessRfProfile")
	}

	result := response.Result().(*ResponseWirelessUpdateNetworkWirelessRfProfile)
	return result, response, err

}

//UpdateNetworkWirelessSettings Update the wireless settings for a network
/* Update the wireless settings for a network

@param networkID networkId path parameter. Network ID
*/
func (s *WirelessService) UpdateNetworkWirelessSettings(networkID string, requestWirelessUpdateNetworkWirelessSettings *RequestWirelessUpdateNetworkWirelessSettings) (*ResponseWirelessUpdateNetworkWirelessSettings, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/settings"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateNetworkWirelessSettings).
		SetResult(&ResponseWirelessUpdateNetworkWirelessSettings{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkWirelessSettings")
	}

	result := response.Result().(*ResponseWirelessUpdateNetworkWirelessSettings)
	return result, response, err

}

//UpdateNetworkWirelessSSID Update the attributes of an MR SSID
/* Update the attributes of an MR SSID

@param networkID networkId path parameter. Network ID
@param number number path parameter.
*/
func (s *WirelessService) UpdateNetworkWirelessSSID(networkID string, number string, requestWirelessUpdateNetworkWirelessSsid *RequestWirelessUpdateNetworkWirelessSSID) (*ResponseWirelessUpdateNetworkWirelessSSID, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ssids/{number}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{number}", fmt.Sprintf("%v", number), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateNetworkWirelessSsid).
		SetResult(&ResponseWirelessUpdateNetworkWirelessSSID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkWirelessSsid")
	}

	result := response.Result().(*ResponseWirelessUpdateNetworkWirelessSSID)
	return result, response, err

}

//UpdateNetworkWirelessSSIDBonjourForwarding Update the bonjour forwarding setting and rules for the SSID
/* Update the bonjour forwarding setting and rules for the SSID

@param networkID networkId path parameter. Network ID
@param number number path parameter.
*/
func (s *WirelessService) UpdateNetworkWirelessSSIDBonjourForwarding(networkID string, number string, requestWirelessUpdateNetworkWirelessSsidBonjourForwarding *RequestWirelessUpdateNetworkWirelessSSIDBonjourForwarding) (*ResponseWirelessUpdateNetworkWirelessSSIDBonjourForwarding, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ssids/{number}/bonjourForwarding"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{number}", fmt.Sprintf("%v", number), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateNetworkWirelessSsidBonjourForwarding).
		SetResult(&ResponseWirelessUpdateNetworkWirelessSSIDBonjourForwarding{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkWirelessSsidBonjourForwarding")
	}

	result := response.Result().(*ResponseWirelessUpdateNetworkWirelessSSIDBonjourForwarding)
	return result, response, err

}

//UpdateNetworkWirelessSSIDDeviceTypeGroupPolicies Update the device type group policies for the SSID
/* Update the device type group policies for the SSID

@param networkID networkId path parameter. Network ID
@param number number path parameter.
*/
func (s *WirelessService) UpdateNetworkWirelessSSIDDeviceTypeGroupPolicies(networkID string, number string, requestWirelessUpdateNetworkWirelessSsidDeviceTypeGroupPolicies *RequestWirelessUpdateNetworkWirelessSSIDDeviceTypeGroupPolicies) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ssids/{number}/deviceTypeGroupPolicies"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{number}", fmt.Sprintf("%v", number), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateNetworkWirelessSsidDeviceTypeGroupPolicies).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkWirelessSsidDeviceTypeGroupPolicies")
	}

	return response, err

}

//UpdateNetworkWirelessSSIDEapOverride Update the EAP overridden parameters for an SSID.
/* Update the EAP overridden parameters for an SSID.

@param networkID networkId path parameter. Network ID
@param number number path parameter.
*/
func (s *WirelessService) UpdateNetworkWirelessSSIDEapOverride(networkID string, number string, requestWirelessUpdateNetworkWirelessSsidEapOverride *RequestWirelessUpdateNetworkWirelessSSIDEapOverride) (*ResponseWirelessUpdateNetworkWirelessSSIDEapOverride, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ssids/{number}/eapOverride"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{number}", fmt.Sprintf("%v", number), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateNetworkWirelessSsidEapOverride).
		SetResult(&ResponseWirelessUpdateNetworkWirelessSSIDEapOverride{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkWirelessSsidEapOverride")
	}

	result := response.Result().(*ResponseWirelessUpdateNetworkWirelessSSIDEapOverride)
	return result, response, err

}

//UpdateNetworkWirelessSSIDFirewallL3FirewallRules Update the L3 firewall rules of an SSID on an MR network
/* Update the L3 firewall rules of an SSID on an MR network

@param networkID networkId path parameter. Network ID
@param number number path parameter.
*/
func (s *WirelessService) UpdateNetworkWirelessSSIDFirewallL3FirewallRules(networkID string, number string, requestWirelessUpdateNetworkWirelessSsidFirewallL3FirewallRules *RequestWirelessUpdateNetworkWirelessSSIDFirewallL3FirewallRules) (*ResponseWirelessUpdateNetworkWirelessSSIDFirewallL3FirewallRules, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ssids/{number}/firewall/l3FirewallRules"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{number}", fmt.Sprintf("%v", number), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateNetworkWirelessSsidFirewallL3FirewallRules).
		SetResult(&ResponseWirelessUpdateNetworkWirelessSSIDFirewallL3FirewallRules{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkWirelessSsidFirewallL3FirewallRules")
	}

	result := response.Result().(*ResponseWirelessUpdateNetworkWirelessSSIDFirewallL3FirewallRules)
	return result, response, err

}

//UpdateNetworkWirelessSSIDFirewallL7FirewallRules Update the L7 firewall rules of an SSID on an MR network
/* Update the L7 firewall rules of an SSID on an MR network

@param networkID networkId path parameter. Network ID
@param number number path parameter.
*/
func (s *WirelessService) UpdateNetworkWirelessSSIDFirewallL7FirewallRules(networkID string, number string, requestWirelessUpdateNetworkWirelessSsidFirewallL7FirewallRules *RequestWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRules) (*ResponseWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRules, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ssids/{number}/firewall/l7FirewallRules"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{number}", fmt.Sprintf("%v", number), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateNetworkWirelessSsidFirewallL7FirewallRules).
		SetResult(&ResponseWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRules{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkWirelessSsidFirewallL7FirewallRules")
	}

	result := response.Result().(*ResponseWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRules)
	return result, response, err

}

//UpdateNetworkWirelessSSIDHotspot20 Update the Hotspot 2.0 settings of an SSID
/* Update the Hotspot 2.0 settings of an SSID

@param networkID networkId path parameter. Network ID
@param number number path parameter.
*/
func (s *WirelessService) UpdateNetworkWirelessSSIDHotspot20(networkID string, number string, requestWirelessUpdateNetworkWirelessSsidHotspot20 *RequestWirelessUpdateNetworkWirelessSSIDHotspot20) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ssids/{number}/hotspot20"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{number}", fmt.Sprintf("%v", number), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateNetworkWirelessSsidHotspot20).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkWirelessSsidHotspot20")
	}

	return response, err

}

//UpdateNetworkWirelessSSIDIDentityPsk Update an Identity PSK
/* Update an Identity PSK

@param networkID networkId path parameter. Network ID
@param number number path parameter.
@param identityPskID identityPskId path parameter. Identity psk ID
*/
func (s *WirelessService) UpdateNetworkWirelessSSIDIDentityPsk(networkID string, number string, identityPskID string, requestWirelessUpdateNetworkWirelessSsidIdentityPsk *RequestWirelessUpdateNetworkWirelessSSIDIDentityPsk) (*ResponseWirelessUpdateNetworkWirelessSSIDIDentityPsk, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ssids/{number}/identityPsks/{identityPskId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{number}", fmt.Sprintf("%v", number), -1)
	path = strings.Replace(path, "{identityPskId}", fmt.Sprintf("%v", identityPskID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateNetworkWirelessSsidIdentityPsk).
		SetResult(&ResponseWirelessUpdateNetworkWirelessSSIDIDentityPsk{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkWirelessSsidIdentityPsk")
	}

	result := response.Result().(*ResponseWirelessUpdateNetworkWirelessSSIDIDentityPsk)
	return result, response, err

}

//UpdateNetworkWirelessSSIDSchedules Update the outage schedule for the SSID
/* Update the outage schedule for the SSID

@param networkID networkId path parameter. Network ID
@param number number path parameter.
*/
func (s *WirelessService) UpdateNetworkWirelessSSIDSchedules(networkID string, number string, requestWirelessUpdateNetworkWirelessSsidSchedules *RequestWirelessUpdateNetworkWirelessSSIDSchedules) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ssids/{number}/schedules"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{number}", fmt.Sprintf("%v", number), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateNetworkWirelessSsidSchedules).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkWirelessSsidSchedules")
	}

	return response, err

}

//UpdateNetworkWirelessSSIDSplashSettings Modify the splash page settings for the given SSID
/* Modify the splash page settings for the given SSID

@param networkID networkId path parameter. Network ID
@param number number path parameter.
*/
func (s *WirelessService) UpdateNetworkWirelessSSIDSplashSettings(networkID string, number string, requestWirelessUpdateNetworkWirelessSsidSplashSettings *RequestWirelessUpdateNetworkWirelessSSIDSplashSettings) (*ResponseWirelessUpdateNetworkWirelessSSIDSplashSettings, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ssids/{number}/splash/settings"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{number}", fmt.Sprintf("%v", number), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateNetworkWirelessSsidSplashSettings).
		SetResult(&ResponseWirelessUpdateNetworkWirelessSSIDSplashSettings{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkWirelessSsidSplashSettings")
	}

	result := response.Result().(*ResponseWirelessUpdateNetworkWirelessSSIDSplashSettings)
	return result, response, err

}

//UpdateNetworkWirelessSSIDTrafficShapingRules Update the traffic shaping rules for an SSID on an MR network.
/* Update the traffic shaping rules for an SSID on an MR network.

@param networkID networkId path parameter. Network ID
@param number number path parameter.
*/
func (s *WirelessService) UpdateNetworkWirelessSSIDTrafficShapingRules(networkID string, number string, requestWirelessUpdateNetworkWirelessSsidTrafficShapingRules *RequestWirelessUpdateNetworkWirelessSSIDTrafficShapingRules) (*ResponseWirelessUpdateNetworkWirelessSSIDTrafficShapingRules, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ssids/{number}/trafficShaping/rules"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{number}", fmt.Sprintf("%v", number), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateNetworkWirelessSsidTrafficShapingRules).
		SetResult(&ResponseWirelessUpdateNetworkWirelessSSIDTrafficShapingRules{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkWirelessSsidTrafficShapingRules")
	}

	result := response.Result().(*ResponseWirelessUpdateNetworkWirelessSSIDTrafficShapingRules)
	return result, response, err

}

//UpdateNetworkWirelessSSIDVpn Update the VPN settings for the SSID
/* Update the VPN settings for the SSID

@param networkID networkId path parameter. Network ID
@param number number path parameter.
*/
func (s *WirelessService) UpdateNetworkWirelessSSIDVpn(networkID string, number string, requestWirelessUpdateNetworkWirelessSsidVpn *RequestWirelessUpdateNetworkWirelessSSIDVpn) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ssids/{number}/vpn"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{number}", fmt.Sprintf("%v", number), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateNetworkWirelessSsidVpn).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkWirelessSsidVpn")
	}

	return response, err

}

//DeleteNetworkWirelessEthernetPortsProfile Delete an AP port profile
/* Delete an AP port profile

@param networkID networkId path parameter. Network ID
@param profileID profileId path parameter. Profile ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-network-wireless-ethernet-ports-profile
*/
func (s *WirelessService) DeleteNetworkWirelessEthernetPortsProfile(networkID string, profileID string) (*resty.Response, error) {
	//networkID string,profileID string
	path := "/api/v1/networks/{networkId}/wireless/ethernet/ports/profiles/{profileId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{profileId}", fmt.Sprintf("%v", profileID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteNetworkWirelessEthernetPortsProfile")
	}

	return response, err

}

//DeleteNetworkWirelessRfProfile Delete a RF Profile
/* Delete a RF Profile

@param networkID networkId path parameter. Network ID
@param rfProfileID rfProfileId path parameter. Rf profile ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-network-wireless-rf-profile
*/
func (s *WirelessService) DeleteNetworkWirelessRfProfile(networkID string, rfProfileID string) (*resty.Response, error) {
	//networkID string,rfProfileID string
	path := "/api/v1/networks/{networkId}/wireless/rfProfiles/{rfProfileId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{rfProfileId}", fmt.Sprintf("%v", rfProfileID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteNetworkWirelessRfProfile")
	}

	return response, err

}

//DeleteNetworkWirelessSSIDIDentityPsk Delete an Identity PSK
/* Delete an Identity PSK

@param networkID networkId path parameter. Network ID
@param number number path parameter.
@param identityPskID identityPskId path parameter. Identity psk ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-network-wireless-ssid-identity-psk
*/
func (s *WirelessService) DeleteNetworkWirelessSSIDIDentityPsk(networkID string, number string, identityPskID string) (*resty.Response, error) {
	//networkID string,number string,identityPskID string
	path := "/api/v1/networks/{networkId}/wireless/ssids/{number}/identityPsks/{identityPskId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{number}", fmt.Sprintf("%v", number), -1)
	path = strings.Replace(path, "{identityPskId}", fmt.Sprintf("%v", identityPskID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteNetworkWirelessSsidIdentityPsk")
	}

	return response, err

}
