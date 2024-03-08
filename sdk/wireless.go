package meraki

import (
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
	BasicServiceSets *[]ResponseWirelessGetDeviceWirelessStatusBasicServiceSets `json:"basicServiceSets,omitempty"` //
}
type ResponseWirelessGetDeviceWirelessStatusBasicServiceSets struct {
	Band         string `json:"band,omitempty"`         //
	Broadcasting *bool  `json:"broadcasting,omitempty"` //
	Bssid        string `json:"bssid,omitempty"`        //
	Channel      *int   `json:"channel,omitempty"`      //
	ChannelWidth string `json:"channelWidth,omitempty"` //
	Enabled      *bool  `json:"enabled,omitempty"`      //
	Power        string `json:"power,omitempty"`        //
	SSIDName     string `json:"ssidName,omitempty"`     //
	SSIDNumber   *int   `json:"ssidNumber,omitempty"`   //
	Visible      *bool  `json:"visible,omitempty"`      //
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
	Currency string                                            `json:"currency,omitempty"` //
	Plans    *[]ResponseWirelessGetNetworkWirelessBillingPlans `json:"plans,omitempty"`    //
}
type ResponseWirelessGetNetworkWirelessBillingPlans struct {
	BandwidthLimits *ResponseWirelessGetNetworkWirelessBillingPlansBandwidthLimits `json:"bandwidthLimits,omitempty"` //
	ID              string                                                         `json:"id,omitempty"`              //
	Price           *float64                                                       `json:"price,omitempty"`           //
	TimeLimit       string                                                         `json:"timeLimit,omitempty"`       //
}
type ResponseWirelessGetNetworkWirelessBillingPlansBandwidthLimits struct {
	LimitDown *int `json:"limitDown,omitempty"` //
	LimitUp   *int `json:"limitUp,omitempty"`   //
}
type ResponseWirelessUpdateNetworkWirelessBilling interface{}
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
	LatencyBinsByCategory *ResponseItemWirelessGetNetworkWirelessClientLatencyHistoryLatencyBinsByCategory `json:"latencyBinsByCategory,omitempty"` //
	T0                    *int                                                                             `json:"t0,omitempty"`                    //
	T1                    *int                                                                             `json:"t1,omitempty"`                    //
}
type ResponseItemWirelessGetNetworkWirelessClientLatencyHistoryLatencyBinsByCategory struct {
	BackgroundTraffic *ResponseItemWirelessGetNetworkWirelessClientLatencyHistoryLatencyBinsByCategoryBackgroundTraffic `json:"backgroundTraffic,omitempty"` //
	BestEffortTraffic *ResponseItemWirelessGetNetworkWirelessClientLatencyHistoryLatencyBinsByCategoryBestEffortTraffic `json:"bestEffortTraffic,omitempty"` //
	VideoTraffic      *ResponseItemWirelessGetNetworkWirelessClientLatencyHistoryLatencyBinsByCategoryVideoTraffic      `json:"videoTraffic,omitempty"`      //
	VoiceTraffic      *ResponseItemWirelessGetNetworkWirelessClientLatencyHistoryLatencyBinsByCategoryVoiceTraffic      `json:"voiceTraffic,omitempty"`      //
}
type ResponseItemWirelessGetNetworkWirelessClientLatencyHistoryLatencyBinsByCategoryBackgroundTraffic struct {
	Status05    *int `json:"0.5,omitempty"`    //
	Status10    *int `json:"1.0,omitempty"`    //
	Status10240 *int `json:"1024.0,omitempty"` //
	Status1280  *int `json:"128.0,omitempty"`  //
	Status160   *int `json:"16.0,omitempty"`   //
	Status20    *int `json:"2.0,omitempty"`    //
	Status20480 *int `json:"2048.0,omitempty"` //
	Status2560  *int `json:"256.0,omitempty"`  //
	Status320   *int `json:"32.0,omitempty"`   //
	Status40    *int `json:"4.0,omitempty"`    //
	Status5120  *int `json:"512.0,omitempty"`  //
	Status640   *int `json:"64.0,omitempty"`   //
	Status80    *int `json:"8.0,omitempty"`    //
}
type ResponseItemWirelessGetNetworkWirelessClientLatencyHistoryLatencyBinsByCategoryBestEffortTraffic struct {
	Status05    *int `json:"0.5,omitempty"`    //
	Status10    *int `json:"1.0,omitempty"`    //
	Status10240 *int `json:"1024.0,omitempty"` //
	Status1280  *int `json:"128.0,omitempty"`  //
	Status160   *int `json:"16.0,omitempty"`   //
	Status20    *int `json:"2.0,omitempty"`    //
	Status20480 *int `json:"2048.0,omitempty"` //
	Status2560  *int `json:"256.0,omitempty"`  //
	Status320   *int `json:"32.0,omitempty"`   //
	Status40    *int `json:"4.0,omitempty"`    //
	Status5120  *int `json:"512.0,omitempty"`  //
	Status640   *int `json:"64.0,omitempty"`   //
	Status80    *int `json:"8.0,omitempty"`    //
}
type ResponseItemWirelessGetNetworkWirelessClientLatencyHistoryLatencyBinsByCategoryVideoTraffic struct {
	Status05    *int `json:"0.5,omitempty"`    //
	Status10    *int `json:"1.0,omitempty"`    //
	Status10240 *int `json:"1024.0,omitempty"` //
	Status1280  *int `json:"128.0,omitempty"`  //
	Status160   *int `json:"16.0,omitempty"`   //
	Status20    *int `json:"2.0,omitempty"`    //
	Status20480 *int `json:"2048.0,omitempty"` //
	Status2560  *int `json:"256.0,omitempty"`  //
	Status320   *int `json:"32.0,omitempty"`   //
	Status40    *int `json:"4.0,omitempty"`    //
	Status5120  *int `json:"512.0,omitempty"`  //
	Status640   *int `json:"64.0,omitempty"`   //
	Status80    *int `json:"8.0,omitempty"`    //
}
type ResponseItemWirelessGetNetworkWirelessClientLatencyHistoryLatencyBinsByCategoryVoiceTraffic struct {
	Status05    *int `json:"0.5,omitempty"`    //
	Status10    *int `json:"1.0,omitempty"`    //
	Status10240 *int `json:"1024.0,omitempty"` //
	Status1280  *int `json:"128.0,omitempty"`  //
	Status160   *int `json:"16.0,omitempty"`   //
	Status20    *int `json:"2.0,omitempty"`    //
	Status20480 *int `json:"2048.0,omitempty"` //
	Status2560  *int `json:"256.0,omitempty"`  //
	Status320   *int `json:"32.0,omitempty"`   //
	Status40    *int `json:"4.0,omitempty"`    //
	Status5120  *int `json:"512.0,omitempty"`  //
	Status640   *int `json:"64.0,omitempty"`   //
	Status80    *int `json:"8.0,omitempty"`    //
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
type ResponseWirelessGetNetworkWirelessMeshStatuses struct {
	LatestMeshPerformance *ResponseWirelessGetNetworkWirelessMeshStatusesLatestMeshPerformance `json:"latestMeshPerformance,omitempty"` //
	MeshRoute             []string                                                             `json:"meshRoute,omitempty"`             //
	Serial                string                                                               `json:"serial,omitempty"`                //
}
type ResponseWirelessGetNetworkWirelessMeshStatusesLatestMeshPerformance struct {
	Mbps            *int   `json:"mbps,omitempty"`            //
	Metric          *int   `json:"metric,omitempty"`          //
	UsagePercentage string `json:"usagePercentage,omitempty"` //
}
type ResponseWirelessGetNetworkWirelessRfProfiles []ResponseItemWirelessGetNetworkWirelessRfProfiles // Array of ResponseWirelessGetNetworkWirelessRfProfiles
type ResponseItemWirelessGetNetworkWirelessRfProfiles struct {
	AfcEnabled             *bool                                                               `json:"afcEnabled,omitempty"`             //
	ApBandSettings         *ResponseItemWirelessGetNetworkWirelessRfProfilesApBandSettings     `json:"apBandSettings,omitempty"`         //
	BandSelectionType      string                                                              `json:"bandSelectionType,omitempty"`      //
	ClientBalancingEnabled *bool                                                               `json:"clientBalancingEnabled,omitempty"` //
	FiveGhzSettings        *ResponseItemWirelessGetNetworkWirelessRfProfilesFiveGhzSettings    `json:"fiveGhzSettings,omitempty"`        //
	ID                     string                                                              `json:"id,omitempty"`                     //
	MinBitrateType         string                                                              `json:"minBitrateType,omitempty"`         //
	Name                   string                                                              `json:"name,omitempty"`                   //
	NetworkID              string                                                              `json:"networkId,omitempty"`              //
	PerSSIDSettings        *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSSIDSettings    `json:"perSsidSettings,omitempty"`        //
	SixGhzSettings         *ResponseItemWirelessGetNetworkWirelessRfProfilesSixGhzSettings     `json:"sixGhzSettings,omitempty"`         //
	Transmission           *ResponseItemWirelessGetNetworkWirelessRfProfilesTransmission       `json:"transmission,omitempty"`           //
	TwoFourGhzSettings     *ResponseItemWirelessGetNetworkWirelessRfProfilesTwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"`     //
}
type ResponseItemWirelessGetNetworkWirelessRfProfilesApBandSettings struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   //
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` //
}
type ResponseItemWirelessGetNetworkWirelessRfProfilesFiveGhzSettings struct {
	ChannelWidth      string   `json:"channelWidth,omitempty"`      //
	MaxPower          *int     `json:"maxPower,omitempty"`          //
	MinBitrate        *int     `json:"minBitrate,omitempty"`        //
	MinPower          *int     `json:"minPower,omitempty"`          //
	ValidAutoChannels []string `json:"validAutoChannels,omitempty"` //
}
type ResponseItemWirelessGetNetworkWirelessRfProfilesPerSSIDSettings struct {
	Status0  *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSSIDSettings0  `json:"0,omitempty"`  //
	Status1  *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSSIDSettings1  `json:"1,omitempty"`  //
	Status10 *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSSIDSettings10 `json:"10,omitempty"` //
	Status11 *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSSIDSettings11 `json:"11,omitempty"` //
	Status12 *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSSIDSettings12 `json:"12,omitempty"` //
	Status13 *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSSIDSettings13 `json:"13,omitempty"` //
	Status14 *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSSIDSettings14 `json:"14,omitempty"` //
	Status2  *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSSIDSettings2  `json:"2,omitempty"`  //
	Status3  *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSSIDSettings3  `json:"3,omitempty"`  //
	Status4  *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSSIDSettings4  `json:"4,omitempty"`  //
	Status5  *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSSIDSettings5  `json:"5,omitempty"`  //
	Status6  *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSSIDSettings6  `json:"6,omitempty"`  //
	Status7  *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSSIDSettings7  `json:"7,omitempty"`  //
	Status8  *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSSIDSettings8  `json:"8,omitempty"`  //
	Status9  *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSSIDSettings9  `json:"9,omitempty"`  //
}
type ResponseItemWirelessGetNetworkWirelessRfProfilesPerSSIDSettings0 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   //
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` //
	MinBitrate          *int   `json:"minBitrate,omitempty"`          //
	Name                string `json:"name,omitempty"`                //
}
type ResponseItemWirelessGetNetworkWirelessRfProfilesPerSSIDSettings1 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   //
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` //
	MinBitrate          *int   `json:"minBitrate,omitempty"`          //
	Name                string `json:"name,omitempty"`                //
}
type ResponseItemWirelessGetNetworkWirelessRfProfilesPerSSIDSettings10 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   //
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` //
	MinBitrate          *int   `json:"minBitrate,omitempty"`          //
	Name                string `json:"name,omitempty"`                //
}
type ResponseItemWirelessGetNetworkWirelessRfProfilesPerSSIDSettings11 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   //
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` //
	MinBitrate          *int   `json:"minBitrate,omitempty"`          //
	Name                string `json:"name,omitempty"`                //
}
type ResponseItemWirelessGetNetworkWirelessRfProfilesPerSSIDSettings12 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   //
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` //
	MinBitrate          *int   `json:"minBitrate,omitempty"`          //
	Name                string `json:"name,omitempty"`                //
}
type ResponseItemWirelessGetNetworkWirelessRfProfilesPerSSIDSettings13 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   //
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` //
	MinBitrate          *int   `json:"minBitrate,omitempty"`          //
	Name                string `json:"name,omitempty"`                //
}
type ResponseItemWirelessGetNetworkWirelessRfProfilesPerSSIDSettings14 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   //
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` //
	MinBitrate          *int   `json:"minBitrate,omitempty"`          //
	Name                string `json:"name,omitempty"`                //
}
type ResponseItemWirelessGetNetworkWirelessRfProfilesPerSSIDSettings2 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   //
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` //
	MinBitrate          *int   `json:"minBitrate,omitempty"`          //
	Name                string `json:"name,omitempty"`                //
}
type ResponseItemWirelessGetNetworkWirelessRfProfilesPerSSIDSettings3 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   //
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` //
	MinBitrate          *int   `json:"minBitrate,omitempty"`          //
	Name                string `json:"name,omitempty"`                //
}
type ResponseItemWirelessGetNetworkWirelessRfProfilesPerSSIDSettings4 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   //
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` //
	MinBitrate          *int   `json:"minBitrate,omitempty"`          //
	Name                string `json:"name,omitempty"`                //
}
type ResponseItemWirelessGetNetworkWirelessRfProfilesPerSSIDSettings5 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   //
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` //
	MinBitrate          *int   `json:"minBitrate,omitempty"`          //
	Name                string `json:"name,omitempty"`                //
}
type ResponseItemWirelessGetNetworkWirelessRfProfilesPerSSIDSettings6 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   //
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` //
	MinBitrate          *int   `json:"minBitrate,omitempty"`          //
	Name                string `json:"name,omitempty"`                //
}
type ResponseItemWirelessGetNetworkWirelessRfProfilesPerSSIDSettings7 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   //
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` //
	MinBitrate          *int   `json:"minBitrate,omitempty"`          //
	Name                string `json:"name,omitempty"`                //
}
type ResponseItemWirelessGetNetworkWirelessRfProfilesPerSSIDSettings8 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   //
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` //
	MinBitrate          *int   `json:"minBitrate,omitempty"`          //
	Name                string `json:"name,omitempty"`                //
}
type ResponseItemWirelessGetNetworkWirelessRfProfilesPerSSIDSettings9 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   //
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` //
	MinBitrate          *int   `json:"minBitrate,omitempty"`          //
	Name                string `json:"name,omitempty"`                //
}
type ResponseItemWirelessGetNetworkWirelessRfProfilesSixGhzSettings struct {
	AfcEnabled        *bool    `json:"afcEnabled,omitempty"`        //
	ChannelWidth      string   `json:"channelWidth,omitempty"`      //
	MaxPower          *int     `json:"maxPower,omitempty"`          //
	MinBitrate        *int     `json:"minBitrate,omitempty"`        //
	MinPower          *int     `json:"minPower,omitempty"`          //
	ValidAutoChannels []string `json:"validAutoChannels,omitempty"` //
}
type ResponseItemWirelessGetNetworkWirelessRfProfilesTransmission struct {
	Enabled *bool `json:"enabled,omitempty"` //
}
type ResponseItemWirelessGetNetworkWirelessRfProfilesTwoFourGhzSettings struct {
	AxEnabled         *bool    `json:"axEnabled,omitempty"`         //
	MaxPower          *int     `json:"maxPower,omitempty"`          //
	MinBitrate        *int     `json:"minBitrate,omitempty"`        //
	MinPower          *int     `json:"minPower,omitempty"`          //
	ValidAutoChannels []string `json:"validAutoChannels,omitempty"` //
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
	Transmission           *ResponseWirelessCreateNetworkWirelessRfProfileTransmission       `json:"transmission,omitempty"`           // Settings related to radio transmission.
	TwoFourGhzSettings     *ResponseWirelessCreateNetworkWirelessRfProfileTwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"`     // Settings related to 2.4Ghz band
}
type ResponseWirelessCreateNetworkWirelessRfProfileApBandSettings struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'. Defaults to dual.
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band. Can be either true or false. Defaults to true.
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
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *int   `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings1 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *int   `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings10 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *int   `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings11 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *int   `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings12 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *int   `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings13 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *int   `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings14 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *int   `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings2 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *int   `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings3 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *int   `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings4 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *int   `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings5 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *int   `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings6 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *int   `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings7 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *int   `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings8 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *int   `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessCreateNetworkWirelessRfProfilePerSSIDSettings9 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *int   `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string `json:"name,omitempty"`                // Name of SSID
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
	AfcEnabled             *bool                                                          `json:"afcEnabled,omitempty"`             //
	ApBandSettings         *ResponseWirelessGetNetworkWirelessRfProfileApBandSettings     `json:"apBandSettings,omitempty"`         //
	BandSelectionType      string                                                         `json:"bandSelectionType,omitempty"`      //
	ClientBalancingEnabled *bool                                                          `json:"clientBalancingEnabled,omitempty"` //
	FiveGhzSettings        *ResponseWirelessGetNetworkWirelessRfProfileFiveGhzSettings    `json:"fiveGhzSettings,omitempty"`        //
	ID                     string                                                         `json:"id,omitempty"`                     //
	MinBitrateType         string                                                         `json:"minBitrateType,omitempty"`         //
	Name                   string                                                         `json:"name,omitempty"`                   //
	NetworkID              string                                                         `json:"networkId,omitempty"`              //
	PerSSIDSettings        *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings    `json:"perSsidSettings,omitempty"`        //
	SixGhzSettings         *ResponseWirelessGetNetworkWirelessRfProfileSixGhzSettings     `json:"sixGhzSettings,omitempty"`         //
	Transmission           *ResponseWirelessGetNetworkWirelessRfProfileTransmission       `json:"transmission,omitempty"`           //
	TwoFourGhzSettings     *ResponseWirelessGetNetworkWirelessRfProfileTwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"`     //
}
type ResponseWirelessGetNetworkWirelessRfProfileApBandSettings struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   //
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` //
}
type ResponseWirelessGetNetworkWirelessRfProfileFiveGhzSettings struct {
	ChannelWidth      string   `json:"channelWidth,omitempty"`      //
	MaxPower          *int     `json:"maxPower,omitempty"`          //
	MinBitrate        *int     `json:"minBitrate,omitempty"`        //
	MinPower          *int     `json:"minPower,omitempty"`          //
	ValidAutoChannels []string `json:"validAutoChannels,omitempty"` //
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings struct {
	Status0  *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings0  `json:"0,omitempty"`  //
	Status1  *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings1  `json:"1,omitempty"`  //
	Status10 *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings10 `json:"10,omitempty"` //
	Status11 *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings11 `json:"11,omitempty"` //
	Status12 *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings12 `json:"12,omitempty"` //
	Status13 *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings13 `json:"13,omitempty"` //
	Status14 *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings14 `json:"14,omitempty"` //
	Status2  *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings2  `json:"2,omitempty"`  //
	Status3  *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings3  `json:"3,omitempty"`  //
	Status4  *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings4  `json:"4,omitempty"`  //
	Status5  *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings5  `json:"5,omitempty"`  //
	Status6  *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings6  `json:"6,omitempty"`  //
	Status7  *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings7  `json:"7,omitempty"`  //
	Status8  *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings8  `json:"8,omitempty"`  //
	Status9  *ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings9  `json:"9,omitempty"`  //
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings0 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   //
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` //
	MinBitrate          *int   `json:"minBitrate,omitempty"`          //
	Name                string `json:"name,omitempty"`                //
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings1 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   //
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` //
	MinBitrate          *int   `json:"minBitrate,omitempty"`          //
	Name                string `json:"name,omitempty"`                //
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings10 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   //
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` //
	MinBitrate          *int   `json:"minBitrate,omitempty"`          //
	Name                string `json:"name,omitempty"`                //
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings11 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   //
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` //
	MinBitrate          *int   `json:"minBitrate,omitempty"`          //
	Name                string `json:"name,omitempty"`                //
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings12 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   //
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` //
	MinBitrate          *int   `json:"minBitrate,omitempty"`          //
	Name                string `json:"name,omitempty"`                //
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings13 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   //
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` //
	MinBitrate          *int   `json:"minBitrate,omitempty"`          //
	Name                string `json:"name,omitempty"`                //
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings14 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   //
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` //
	MinBitrate          *int   `json:"minBitrate,omitempty"`          //
	Name                string `json:"name,omitempty"`                //
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings2 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   //
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` //
	MinBitrate          *int   `json:"minBitrate,omitempty"`          //
	Name                string `json:"name,omitempty"`                //
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings3 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   //
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` //
	MinBitrate          *int   `json:"minBitrate,omitempty"`          //
	Name                string `json:"name,omitempty"`                //
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings4 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   //
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` //
	MinBitrate          *int   `json:"minBitrate,omitempty"`          //
	Name                string `json:"name,omitempty"`                //
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings5 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   //
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` //
	MinBitrate          *int   `json:"minBitrate,omitempty"`          //
	Name                string `json:"name,omitempty"`                //
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings6 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   //
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` //
	MinBitrate          *int   `json:"minBitrate,omitempty"`          //
	Name                string `json:"name,omitempty"`                //
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings7 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   //
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` //
	MinBitrate          *int   `json:"minBitrate,omitempty"`          //
	Name                string `json:"name,omitempty"`                //
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings8 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   //
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` //
	MinBitrate          *int   `json:"minBitrate,omitempty"`          //
	Name                string `json:"name,omitempty"`                //
}
type ResponseWirelessGetNetworkWirelessRfProfilePerSSIDSettings9 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   //
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` //
	MinBitrate          *int   `json:"minBitrate,omitempty"`          //
	Name                string `json:"name,omitempty"`                //
}
type ResponseWirelessGetNetworkWirelessRfProfileSixGhzSettings struct {
	AfcEnabled        *bool    `json:"afcEnabled,omitempty"`        //
	ChannelWidth      string   `json:"channelWidth,omitempty"`      //
	MaxPower          *int     `json:"maxPower,omitempty"`          //
	MinBitrate        *int     `json:"minBitrate,omitempty"`        //
	MinPower          *int     `json:"minPower,omitempty"`          //
	ValidAutoChannels []string `json:"validAutoChannels,omitempty"` //
}
type ResponseWirelessGetNetworkWirelessRfProfileTransmission struct {
	Enabled *bool `json:"enabled,omitempty"` //
}
type ResponseWirelessGetNetworkWirelessRfProfileTwoFourGhzSettings struct {
	AxEnabled         *bool    `json:"axEnabled,omitempty"`         //
	MaxPower          *int     `json:"maxPower,omitempty"`          //
	MinBitrate        *int     `json:"minBitrate,omitempty"`        //
	MinPower          *int     `json:"minPower,omitempty"`          //
	ValidAutoChannels []string `json:"validAutoChannels,omitempty"` //
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
	Transmission           *ResponseWirelessUpdateNetworkWirelessRfProfileTransmission       `json:"transmission,omitempty"`           // Settings related to radio transmission.
	TwoFourGhzSettings     *ResponseWirelessUpdateNetworkWirelessRfProfileTwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"`     // Settings related to 2.4Ghz band
}
type ResponseWirelessUpdateNetworkWirelessRfProfileApBandSettings struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'. Defaults to dual.
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band. Can be either true or false. Defaults to true.
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
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *int   `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings1 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *int   `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings10 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *int   `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings11 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *int   `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings12 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *int   `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings13 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *int   `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings14 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *int   `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings2 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *int   `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings3 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *int   `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings4 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *int   `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings5 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *int   `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings6 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *int   `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings7 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *int   `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings8 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *int   `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string `json:"name,omitempty"`                // Name of SSID
}
type ResponseWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings9 struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *int   `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	Name                string `json:"name,omitempty"`                // Name of SSID
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
	IPv6BridgeEnabled        *bool                                                 `json:"ipv6BridgeEnabled,omitempty"`        // Toggle for enabling or disabling IPv6 bridging in a network (Note: if enabled, SSIDs must also be configured to use bridge mode)
	LedLightsOn              *bool                                                 `json:"ledLightsOn,omitempty"`              // Toggle for enabling or disabling LED lights on all APs in the network (making them run dark)
	LocationAnalyticsEnabled *bool                                                 `json:"locationAnalyticsEnabled,omitempty"` // Toggle for enabling or disabling location analytics for your network
	MeshingEnabled           *bool                                                 `json:"meshingEnabled,omitempty"`           // Toggle for enabling or disabling meshing in a network
	NamedVLANs               *ResponseWirelessGetNetworkWirelessSettingsNamedVLANs `json:"namedVlans,omitempty"`               // Named VLAN settings for wireless networks.
	Upgradestrategy          string                                                `json:"upgradeStrategy,omitempty"`          // The upgrade strategy to apply to the network. Must be one of 'minimizeUpgradeTime' or 'minimizeClientDowntime'. Requires firmware version MR 26.8 or higher'
}
type ResponseWirelessGetNetworkWirelessSettingsNamedVLANs struct {
	PoolDhcpMonitoring *ResponseWirelessGetNetworkWirelessSettingsNamedVLANsPoolDhcpMonitoring `json:"poolDhcpMonitoring,omitempty"` // Named VLAN Pool DHCP Monitoring settings.
}
type ResponseWirelessGetNetworkWirelessSettingsNamedVLANsPoolDhcpMonitoring struct {
	Duration *int  `json:"duration,omitempty"` // The duration in minutes that devices will refrain from using dirty VLANs before adding them back to the pool.
	Enabled  *bool `json:"enabled,omitempty"`  // Whether or not devices using named VLAN pools should remove dirty VLANs from the pool, thereby preventing clients from being assigned to VLANs where they would be unable to obtain an IP address via DHCP
}
type ResponseWirelessUpdateNetworkWirelessSettings struct {
	IPv6BridgeEnabled        *bool                                                    `json:"ipv6BridgeEnabled,omitempty"`        // Toggle for enabling or disabling IPv6 bridging in a network (Note: if enabled, SSIDs must also be configured to use bridge mode)
	LedLightsOn              *bool                                                    `json:"ledLightsOn,omitempty"`              // Toggle for enabling or disabling LED lights on all APs in the network (making them run dark)
	LocationAnalyticsEnabled *bool                                                    `json:"locationAnalyticsEnabled,omitempty"` // Toggle for enabling or disabling location analytics for your network
	MeshingEnabled           *bool                                                    `json:"meshingEnabled,omitempty"`           // Toggle for enabling or disabling meshing in a network
	NamedVLANs               *ResponseWirelessUpdateNetworkWirelessSettingsNamedVLANs `json:"namedVlans,omitempty"`               // Named VLAN settings for wireless networks.
	Upgradestrategy          string                                                   `json:"upgradeStrategy,omitempty"`          // The upgrade strategy to apply to the network. Must be one of 'minimizeUpgradeTime' or 'minimizeClientDowntime'. Requires firmware version MR 26.8 or higher'
}
type ResponseWirelessUpdateNetworkWirelessSettingsNamedVLANs struct {
	PoolDhcpMonitoring *ResponseWirelessUpdateNetworkWirelessSettingsNamedVLANsPoolDhcpMonitoring `json:"poolDhcpMonitoring,omitempty"` // Named VLAN Pool DHCP Monitoring settings.
}
type ResponseWirelessUpdateNetworkWirelessSettingsNamedVLANsPoolDhcpMonitoring struct {
	Duration *int  `json:"duration,omitempty"` // The duration in minutes that devices will refrain from using dirty VLANs before adding them back to the pool.
	Enabled  *bool `json:"enabled,omitempty"`  // Whether or not devices using named VLAN pools should remove dirty VLANs from the pool, thereby preventing clients from being assigned to VLANs where they would be unable to obtain an IP address via DHCP
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
	AdminSplashURL                  string                                                                `json:"adminSplashUrl,omitempty"`                  //
	AuthMode                        string                                                                `json:"authMode,omitempty"`                        //
	AvailabilityTags                []string                                                              `json:"availabilityTags,omitempty"`                //
	AvailableOnAllAps               *bool                                                                 `json:"availableOnAllAps,omitempty"`               //
	BandSelection                   string                                                                `json:"bandSelection,omitempty"`                   //
	Enabled                         *bool                                                                 `json:"enabled,omitempty"`                         //
	EncryptionMode                  string                                                                `json:"encryptionMode,omitempty"`                  //
	IPAssignmentMode                string                                                                `json:"ipAssignmentMode,omitempty"`                //
	MandatoryDhcpEnabled            *bool                                                                 `json:"mandatoryDhcpEnabled,omitempty"`            //
	MinBitrate                      *int                                                                  `json:"minBitrate,omitempty"`                      //
	Name                            string                                                                `json:"name,omitempty"`                            //
	Number                          *int                                                                  `json:"number,omitempty"`                          //
	PerClientBandwidthLimitDown     *int                                                                  `json:"perClientBandwidthLimitDown,omitempty"`     //
	PerClientBandwidthLimitUp       *int                                                                  `json:"perClientBandwidthLimitUp,omitempty"`       //
	PerSSIDBandwidthLimitDown       *int                                                                  `json:"perSsidBandwidthLimitDown,omitempty"`       //
	PerSSIDBandwidthLimitUp         *int                                                                  `json:"perSsidBandwidthLimitUp,omitempty"`         //
	RadiusAccountingEnabled         *bool                                                                 `json:"radiusAccountingEnabled,omitempty"`         //
	RadiusAccountingServers         *[]ResponseItemWirelessGetNetworkWirelessSSIDsRadiusAccountingServers `json:"radiusAccountingServers,omitempty"`         //
	RadiusAttributeForGroupPolicies string                                                                `json:"radiusAttributeForGroupPolicies,omitempty"` //
	RadiusEnabled                   *bool                                                                 `json:"radiusEnabled,omitempty"`                   //
	RadiusFailoverPolicy            string                                                                `json:"radiusFailoverPolicy,omitempty"`            //
	RadiusLoadBalancingPolicy       string                                                                `json:"radiusLoadBalancingPolicy,omitempty"`       //
	RadiusServers                   *[]ResponseItemWirelessGetNetworkWirelessSSIDsRadiusServers           `json:"radiusServers,omitempty"`                   //
	SplashPage                      string                                                                `json:"splashPage,omitempty"`                      //
	SplashTimeout                   string                                                                `json:"splashTimeout,omitempty"`                   //
	SSIDAdminAccessible             *bool                                                                 `json:"ssidAdminAccessible,omitempty"`             //
	Visible                         *bool                                                                 `json:"visible,omitempty"`                         //
	WalledGardenEnabled             *bool                                                                 `json:"walledGardenEnabled,omitempty"`             //
	WalledGardenRanges              []string                                                              `json:"walledGardenRanges,omitempty"`              //
	WpaEncryptionMode               string                                                                `json:"wpaEncryptionMode,omitempty"`               //
}
type ResponseItemWirelessGetNetworkWirelessSSIDsRadiusAccountingServers struct {
	CaCertificate            string `json:"caCertificate,omitempty"`            //
	Host                     string `json:"host,omitempty"`                     //
	OpenRoamingCertificateID *int   `json:"openRoamingCertificateId,omitempty"` //
	Port                     *int   `json:"port,omitempty"`                     //
}
type ResponseItemWirelessGetNetworkWirelessSSIDsRadiusServers struct {
	CaCertificate            string `json:"caCertificate,omitempty"`            //
	Host                     string `json:"host,omitempty"`                     //
	OpenRoamingCertificateID *int   `json:"openRoamingCertificateId,omitempty"` //
	Port                     *int   `json:"port,omitempty"`                     //
}
type ResponseWirelessGetNetworkWirelessSSID struct {
	AdminSplashURL                  string                                                           `json:"adminSplashUrl,omitempty"`                  //
	AuthMode                        string                                                           `json:"authMode,omitempty"`                        //
	AvailabilityTags                []string                                                         `json:"availabilityTags,omitempty"`                //
	AvailableOnAllAps               *bool                                                            `json:"availableOnAllAps,omitempty"`               //
	BandSelection                   string                                                           `json:"bandSelection,omitempty"`                   //
	Enabled                         *bool                                                            `json:"enabled,omitempty"`                         //
	EncryptionMode                  string                                                           `json:"encryptionMode,omitempty"`                  //
	IPAssignmentMode                string                                                           `json:"ipAssignmentMode,omitempty"`                //
	MandatoryDhcpEnabled            *bool                                                            `json:"mandatoryDhcpEnabled,omitempty"`            //
	MinBitrate                      *int                                                             `json:"minBitrate,omitempty"`                      //
	Name                            string                                                           `json:"name,omitempty"`                            //
	Number                          *int                                                             `json:"number,omitempty"`                          //
	PerClientBandwidthLimitDown     *int                                                             `json:"perClientBandwidthLimitDown,omitempty"`     //
	PerClientBandwidthLimitUp       *int                                                             `json:"perClientBandwidthLimitUp,omitempty"`       //
	PerSSIDBandwidthLimitDown       *int                                                             `json:"perSsidBandwidthLimitDown,omitempty"`       //
	PerSSIDBandwidthLimitUp         *int                                                             `json:"perSsidBandwidthLimitUp,omitempty"`         //
	RadiusAccountingEnabled         *bool                                                            `json:"radiusAccountingEnabled,omitempty"`         //
	RadiusAccountingServers         *[]ResponseWirelessGetNetworkWirelessSSIDRadiusAccountingServers `json:"radiusAccountingServers,omitempty"`         //
	RadiusAttributeForGroupPolicies string                                                           `json:"radiusAttributeForGroupPolicies,omitempty"` //
	RadiusEnabled                   *bool                                                            `json:"radiusEnabled,omitempty"`                   //
	RadiusFailoverPolicy            string                                                           `json:"radiusFailoverPolicy,omitempty"`            //
	RadiusLoadBalancingPolicy       string                                                           `json:"radiusLoadBalancingPolicy,omitempty"`       //
	RadiusServers                   *[]ResponseWirelessGetNetworkWirelessSSIDRadiusServers           `json:"radiusServers,omitempty"`                   //
	SplashPage                      string                                                           `json:"splashPage,omitempty"`                      //
	SplashTimeout                   string                                                           `json:"splashTimeout,omitempty"`                   //
	SSIDAdminAccessible             *bool                                                            `json:"ssidAdminAccessible,omitempty"`             //
	Visible                         *bool                                                            `json:"visible,omitempty"`                         //
	WalledGardenEnabled             *bool                                                            `json:"walledGardenEnabled,omitempty"`             //
	WalledGardenRanges              []string                                                         `json:"walledGardenRanges,omitempty"`              //
	WpaEncryptionMode               string                                                           `json:"wpaEncryptionMode,omitempty"`               //
}
type ResponseWirelessGetNetworkWirelessSSIDRadiusAccountingServers struct {
	CaCertificate            string `json:"caCertificate,omitempty"`            //
	Host                     string `json:"host,omitempty"`                     //
	OpenRoamingCertificateID *int   `json:"openRoamingCertificateId,omitempty"` //
	Port                     *int   `json:"port,omitempty"`                     //
}
type ResponseWirelessGetNetworkWirelessSSIDRadiusServers struct {
	CaCertificate            string `json:"caCertificate,omitempty"`            //
	Host                     string `json:"host,omitempty"`                     //
	OpenRoamingCertificateID *int   `json:"openRoamingCertificateId,omitempty"` //
	Port                     *int   `json:"port,omitempty"`                     //
}
type ResponseWirelessUpdateNetworkWirelessSSID interface{}
type ResponseWirelessGetNetworkWirelessSSIDBonjourForwarding struct {
	Enabled *bool                                                           `json:"enabled,omitempty"` //
	Rules   *[]ResponseWirelessGetNetworkWirelessSSIDBonjourForwardingRules `json:"rules,omitempty"`   //
}
type ResponseWirelessGetNetworkWirelessSSIDBonjourForwardingRules struct {
	Description string   `json:"description,omitempty"` //
	Services    []string `json:"services,omitempty"`    //
	VLANID      string   `json:"vlanId,omitempty"`      //
}
type ResponseWirelessUpdateNetworkWirelessSSIDBonjourForwarding interface{}
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
	Rules *[]ResponseWirelessGetNetworkWirelessSSIDFirewallL3FirewallRulesRules `json:"rules,omitempty"` //
}
type ResponseWirelessGetNetworkWirelessSSIDFirewallL3FirewallRulesRules struct {
	Comment  string `json:"comment,omitempty"`  //
	DestCidr string `json:"destCidr,omitempty"` //
	DestPort string `json:"destPort,omitempty"` //
	Policy   string `json:"policy,omitempty"`   //
	Protocol string `json:"protocol,omitempty"` //
}
type ResponseWirelessUpdateNetworkWirelessSSIDFirewallL3FirewallRules interface{}
type ResponseWirelessGetNetworkWirelessSSIDFirewallL7FirewallRules struct {
	Rules *[]ResponseWirelessGetNetworkWirelessSSIDFirewallL7FirewallRulesRules `json:"rules,omitempty"` //
}
type ResponseWirelessGetNetworkWirelessSSIDFirewallL7FirewallRulesRules struct {
	Policy string `json:"policy,omitempty"` //
	Type   string `json:"type,omitempty"`   //
	Value  string `json:"value,omitempty"`  //
}
type ResponseWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRules interface{}
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
type ResponseWirelessCreateNetworkWirelessSSIDIDentityPsk interface{}
type ResponseWirelessGetNetworkWirelessSSIDIDentityPsk struct {
	Email                 string `json:"email,omitempty"`                 // The email associated with the System's Manager User
	ExpiresAt             string `json:"expiresAt,omitempty"`             // Timestamp for when the Identity PSK expires, or 'null' to never expire
	GroupPolicyID         string `json:"groupPolicyId,omitempty"`         // The group policy to be applied to clients
	ID                    string `json:"id,omitempty"`                    // The unique identifier of the Identity PSK
	Name                  string `json:"name,omitempty"`                  // The name of the Identity PSK
	Passphrase            string `json:"passphrase,omitempty"`            // The passphrase for client authentication
	WifiPersonalNetworkID string `json:"wifiPersonalNetworkId,omitempty"` // The WiFi Personal Network unique identifier
}
type ResponseWirelessUpdateNetworkWirelessSSIDIDentityPsk interface{}
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
	DefaultRulesEnabled   *bool                                                             `json:"defaultRulesEnabled,omitempty"`   //
	Rules                 *[]ResponseWirelessGetNetworkWirelessSSIDTrafficShapingRulesRules `json:"rules,omitempty"`                 //
	TrafficShapingEnabled *bool                                                             `json:"trafficShapingEnabled,omitempty"` //
}
type ResponseWirelessGetNetworkWirelessSSIDTrafficShapingRulesRules struct {
	Definitions              *[]ResponseWirelessGetNetworkWirelessSSIDTrafficShapingRulesRulesDefinitions            `json:"definitions,omitempty"`              //
	DscpTagValue             *int                                                                                    `json:"dscpTagValue,omitempty"`             //
	PcpTagValue              *int                                                                                    `json:"pcpTagValue,omitempty"`              //
	PerClientBandwidthLimits *ResponseWirelessGetNetworkWirelessSSIDTrafficShapingRulesRulesPerClientBandwidthLimits `json:"perClientBandwidthLimits,omitempty"` //
}
type ResponseWirelessGetNetworkWirelessSSIDTrafficShapingRulesRulesDefinitions struct {
	Type  string `json:"type,omitempty"`  //
	Value string `json:"value,omitempty"` //
}
type ResponseWirelessGetNetworkWirelessSSIDTrafficShapingRulesRulesPerClientBandwidthLimits struct {
	BandwidthLimits *ResponseWirelessGetNetworkWirelessSSIDTrafficShapingRulesRulesPerClientBandwidthLimitsBandwidthLimits `json:"bandwidthLimits,omitempty"` //
	Settings        string                                                                                                 `json:"settings,omitempty"`        //
}
type ResponseWirelessGetNetworkWirelessSSIDTrafficShapingRulesRulesPerClientBandwidthLimitsBandwidthLimits struct {
	LimitDown *int `json:"limitDown,omitempty"` //
	LimitUp   *int `json:"limitUp,omitempty"`   //
}
type ResponseWirelessUpdateNetworkWirelessSSIDTrafficShapingRules interface{}
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
type RequestWirelessCreateNetworkWirelessRfProfile struct {
	ApBandSettings         *RequestWirelessCreateNetworkWirelessRfProfileApBandSettings     `json:"apBandSettings,omitempty"`         // Settings that will be enabled if selectionType is set to 'ap'.
	BandSelectionType      string                                                           `json:"bandSelectionType,omitempty"`      // Band selection can be set to either 'ssid' or 'ap'. This param is required on creation.
	ClientBalancingEnabled *bool                                                            `json:"clientBalancingEnabled,omitempty"` // Steers client to best available access point. Can be either true or false. Defaults to true.
	FiveGhzSettings        *RequestWirelessCreateNetworkWirelessRfProfileFiveGhzSettings    `json:"fiveGhzSettings,omitempty"`        // Settings related to 5Ghz band
	MinBitrateType         string                                                           `json:"minBitrateType,omitempty"`         // Minimum bitrate can be set to either 'band' or 'ssid'. Defaults to band.
	Name                   string                                                           `json:"name,omitempty"`                   // The name of the new profile. Must be unique. This param is required on creation.
	PerSSIDSettings        *RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings    `json:"perSsidSettings,omitempty"`        // Per-SSID radio settings by number.
	SixGhzSettings         *RequestWirelessCreateNetworkWirelessRfProfileSixGhzSettings     `json:"sixGhzSettings,omitempty"`         // Settings related to 6Ghz band
	Transmission           *RequestWirelessCreateNetworkWirelessRfProfileTransmission       `json:"transmission,omitempty"`           // Settings related to radio transmission.
	TwoFourGhzSettings     *RequestWirelessCreateNetworkWirelessRfProfileTwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"`     // Settings related to 2.4Ghz band
}
type RequestWirelessCreateNetworkWirelessRfProfileApBandSettings struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'. Defaults to dual.
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band. Can be either true or false. Defaults to true.
}
type RequestWirelessCreateNetworkWirelessRfProfileFiveGhzSettings struct {
	ChannelWidth      string `json:"channelWidth,omitempty"`      // Sets channel width (MHz) for 5Ghz band. Can be one of 'auto', '20', '40' or '80'. Defaults to auto.
	MaxPower          *int   `json:"maxPower,omitempty"`          // Sets max power (dBm) of 5Ghz band. Can be integer between 2 and 30. Defaults to 30.
	MinBitrate        *int   `json:"minBitrate,omitempty"`        // Sets min bitrate (Mbps) of 5Ghz band. Can be one of '6', '9', '12', '18', '24', '36', '48' or '54'. Defaults to 12.
	MinPower          *int   `json:"minPower,omitempty"`          // Sets min power (dBm) of 5Ghz band. Can be integer between 2 and 30. Defaults to 8.
	Rxsop             *int   `json:"rxsop,omitempty"`             // The RX-SOP level controls the sensitivity of the radio. It is strongly recommended to use RX-SOP only after consulting a wireless expert. RX-SOP can be configured in the range of -65 to -95 (dBm). A value of null will reset this to the default.
	ValidAutoChannels *[]int `json:"validAutoChannels,omitempty"` // Sets valid auto channels for 5Ghz band. Can be one of '36', '40', '44', '48', '52', '56', '60', '64', '100', '104', '108', '112', '116', '120', '124', '128', '132', '136', '140', '144', '149', '153', '157', '161' or '165'.Defaults to [36, 40, 44, 48, 52, 56, 60, 64, 100, 104, 108, 112, 116, 120, 124, 128, 132, 136, 140, 144, 149, 153, 157, 161, 165].
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
	BandOperationMode   string   `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool    `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *float64 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings1 struct {
	BandOperationMode   string   `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool    `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *float64 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings10 struct {
	BandOperationMode   string   `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool    `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *float64 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings11 struct {
	BandOperationMode   string   `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool    `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *float64 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings12 struct {
	BandOperationMode   string   `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool    `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *float64 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings13 struct {
	BandOperationMode   string   `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool    `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *float64 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings14 struct {
	BandOperationMode   string   `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool    `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *float64 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings2 struct {
	BandOperationMode   string   `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool    `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *float64 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings3 struct {
	BandOperationMode   string   `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool    `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *float64 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings4 struct {
	BandOperationMode   string   `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool    `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *float64 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings5 struct {
	BandOperationMode   string   `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool    `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *float64 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings6 struct {
	BandOperationMode   string   `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool    `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *float64 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings7 struct {
	BandOperationMode   string   `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool    `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *float64 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings8 struct {
	BandOperationMode   string   `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool    `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *float64 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings9 struct {
	BandOperationMode   string   `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool    `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *float64 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
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
	MinBitrateType         string                                                           `json:"minBitrateType,omitempty"`         // Minimum bitrate can be set to either 'band' or 'ssid'.
	Name                   string                                                           `json:"name,omitempty"`                   // The name of the new profile. Must be unique.
	PerSSIDSettings        *RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings    `json:"perSsidSettings,omitempty"`        // Per-SSID radio settings by number.
	SixGhzSettings         *RequestWirelessUpdateNetworkWirelessRfProfileSixGhzSettings     `json:"sixGhzSettings,omitempty"`         // Settings related to 6Ghz band
	Transmission           *RequestWirelessUpdateNetworkWirelessRfProfileTransmission       `json:"transmission,omitempty"`           // Settings related to radio transmission.
	TwoFourGhzSettings     *RequestWirelessUpdateNetworkWirelessRfProfileTwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"`     // Settings related to 2.4Ghz band
}
type RequestWirelessUpdateNetworkWirelessRfProfileApBandSettings struct {
	BandOperationMode   string `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool  `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band. Can be either true or false.
}
type RequestWirelessUpdateNetworkWirelessRfProfileFiveGhzSettings struct {
	ChannelWidth      string `json:"channelWidth,omitempty"`      // Sets channel width (MHz) for 5Ghz band. Can be one of 'auto', '20', '40' or '80'.
	MaxPower          *int   `json:"maxPower,omitempty"`          // Sets max power (dBm) of 5Ghz band. Can be integer between 2 and 30.
	MinBitrate        *int   `json:"minBitrate,omitempty"`        // Sets min bitrate (Mbps) of 5Ghz band. Can be one of '6', '9', '12', '18', '24', '36', '48' or '54'.
	MinPower          *int   `json:"minPower,omitempty"`          // Sets min power (dBm) of 5Ghz band. Can be integer between 2 and 30.
	Rxsop             *int   `json:"rxsop,omitempty"`             // The RX-SOP level controls the sensitivity of the radio. It is strongly recommended to use RX-SOP only after consulting a wireless expert. RX-SOP can be configured in the range of -65 to -95 (dBm). A value of null will reset this to the default.
	ValidAutoChannels *[]int `json:"validAutoChannels,omitempty"` // Sets valid auto channels for 5Ghz band. Can be one of '36', '40', '44', '48', '52', '56', '60', '64', '100', '104', '108', '112', '116', '120', '124', '128', '132', '136', '140', '144', '149', '153', '157', '161' or '165'.
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
	BandOperationMode   string   `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool    `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *float64 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings1 struct {
	BandOperationMode   string   `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool    `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *float64 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings10 struct {
	BandOperationMode   string   `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool    `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *float64 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings11 struct {
	BandOperationMode   string   `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool    `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *float64 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings12 struct {
	BandOperationMode   string   `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool    `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *float64 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings13 struct {
	BandOperationMode   string   `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool    `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *float64 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings14 struct {
	BandOperationMode   string   `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool    `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *float64 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings2 struct {
	BandOperationMode   string   `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool    `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *float64 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings3 struct {
	BandOperationMode   string   `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool    `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *float64 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings4 struct {
	BandOperationMode   string   `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool    `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *float64 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings5 struct {
	BandOperationMode   string   `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool    `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *float64 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings6 struct {
	BandOperationMode   string   `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool    `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *float64 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings7 struct {
	BandOperationMode   string   `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool    `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *float64 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings8 struct {
	BandOperationMode   string   `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool    `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *float64 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
}
type RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings9 struct {
	BandOperationMode   string   `json:"bandOperationMode,omitempty"`   // Choice between 'dual', '2.4ghz' or '5ghz'.
	BandSteeringEnabled *bool    `json:"bandSteeringEnabled,omitempty"` // Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	MinBitrate          *float64 `json:"minBitrate,omitempty"`          // Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
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
	IPv6BridgeEnabled        *bool  `json:"ipv6BridgeEnabled,omitempty"`        // Toggle for enabling or disabling IPv6 bridging in a network (Note: if enabled, SSIDs must also be configured to use bridge mode)
	LedLightsOn              *bool  `json:"ledLightsOn,omitempty"`              // Toggle for enabling or disabling LED lights on all APs in the network (making them run dark)
	LocationAnalyticsEnabled *bool  `json:"locationAnalyticsEnabled,omitempty"` // Toggle for enabling or disabling location analytics for your network
	MeshingEnabled           *bool  `json:"meshingEnabled,omitempty"`           // Toggle for enabling or disabling meshing in a network
	Upgradestrategy          string `json:"upgradeStrategy,omitempty"`          // The upgrade strategy to apply to the network. Must be one of 'minimizeUpgradeTime' or 'minimizeClientDowntime'. Requires firmware version MR 26.8 or higher'
}
type RequestWirelessUpdateNetworkWirelessSSID struct {
	ActiveDirectory                  *RequestWirelessUpdateNetworkWirelessSSIDActiveDirectory           `json:"activeDirectory,omitempty"`                  // The current setting for Active Directory. Only valid if splashPage is 'Password-protected with Active Directory'
	AdultContentFilteringEnabled     *bool                                                              `json:"adultContentFilteringEnabled,omitempty"`     // Boolean indicating whether or not adult content will be blocked
	ApTagsAndVLANIDs                 *[]RequestWirelessUpdateNetworkWirelessSSIDApTagsAndVLANIDs        `json:"apTagsAndVlanIds,omitempty"`                 // The list of tags and VLAN IDs used for VLAN tagging. This param is only valid when the ipAssignmentMode is 'Bridge mode' or 'Layer 3 roaming'
	AuthMode                         string                                                             `json:"authMode,omitempty"`                         // The association control method for the SSID ('open', 'open-enhanced', 'psk', 'open-with-radius', 'open-with-nac', '8021x-meraki', '8021x-nac', '8021x-radius', '8021x-google', '8021x-localradius', 'ipsk-with-radius' or 'ipsk-without-radius')
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
	Host string `json:"host,omitempty"` // IP address of your Active Directory server.
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
	Host string `json:"host,omitempty"` // IP address of your LDAP server.
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
type RequestWirelessUpdateNetworkWirelessSSIDOauth struct {
	AllowedDomains []string `json:"allowedDomains,omitempty"` // (Optional) The list of domains allowed access to the network.
}
type RequestWirelessUpdateNetworkWirelessSSIDRadiusAccountingServers struct {
	CaCertificate string `json:"caCertificate,omitempty"` // Certificate used for authorization for the RADSEC Server
	Host          string `json:"host,omitempty"`          // IP address to which the APs will send RADIUS accounting messages
	Port          *int   `json:"port,omitempty"`          // Port on the RADIUS server that is listening for accounting messages
	RadsecEnabled *bool  `json:"radsecEnabled,omitempty"` // Use RADSEC (TLS over TCP) to connect to this RADIUS accounting server. Requires radiusProxyEnabled.
	Secret        string `json:"secret,omitempty"`        // Shared key used to authenticate messages between the APs and RADIUS server
}
type RequestWirelessUpdateNetworkWirelessSSIDRadiusServers struct {
	CaCertificate            string `json:"caCertificate,omitempty"`            // Certificate used for authorization for the RADSEC Server
	Host                     string `json:"host,omitempty"`                     // IP address of your RADIUS server
	OpenRoamingCertificateID *int   `json:"openRoamingCertificateId,omitempty"` // The ID of the Openroaming Certificate attached to radius server.
	Port                     *int   `json:"port,omitempty"`                     // UDP port the RADIUS server listens on for Access-requests
	RadsecEnabled            *bool  `json:"radsecEnabled,omitempty"`            // Use RADSEC (TLS over TCP) to connect to this RADIUS server. Requires radiusProxyEnabled.
	Secret                   string `json:"secret,omitempty"`                   // RADIUS client shared secret
}
type RequestWirelessUpdateNetworkWirelessSSIDSpeedBurst struct {
	Enabled *bool `json:"enabled,omitempty"` // Boolean indicating whether or not to allow users to temporarily exceed the bandwidth limit for short periods while still keeping them under the bandwidth limit over time.
}
type RequestWirelessUpdateNetworkWirelessSSIDBonjourForwarding struct {
	Enabled *bool                                                             `json:"enabled,omitempty"` // If true, Bonjour forwarding is enabled on this SSID.
	Rules   *[]RequestWirelessUpdateNetworkWirelessSSIDBonjourForwardingRules `json:"rules,omitempty"`   // List of bonjour forwarding rules.
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
	Rules          *[]RequestWirelessUpdateNetworkWirelessSSIDFirewallL3FirewallRulesRules `json:"rules,omitempty"`          // An ordered array of the firewall rules for this SSID (not including the local LAN access rule or the default rule)
}
type RequestWirelessUpdateNetworkWirelessSSIDFirewallL3FirewallRulesRules struct {
	Comment  string `json:"comment,omitempty"`  // Description of the rule (optional)
	DestCidr string `json:"destCidr,omitempty"` // Comma-separated list of destination IP address(es) (in IP or CIDR notation), fully-qualified domain names (FQDN) or 'any'
	DestPort string `json:"destPort,omitempty"` // Comma-separated list of destination port(s) (integer in the range 1-65535), or 'any'
	Policy   string `json:"policy,omitempty"`   // 'allow' or 'deny' traffic specified by this rule
	Protocol string `json:"protocol,omitempty"` // The type of protocol (must be 'tcp', 'udp', 'icmp', 'icmp6' or 'any')
}
type RequestWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRules struct {
	Rules *[]RequestWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRulesRules `json:"rules,omitempty"` // An array of L7 firewall rules for this SSID. Rules will get applied in the same order user has specified in request. Empty array will clear the L7 firewall rule configuration.
}
type RequestWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRulesRules struct {
	Policy string `json:"policy,omitempty"` // 'Deny' traffic specified by this rule
	Type   string `json:"type,omitempty"`   // Type of the L7 firewall rule. One of: 'application', 'applicationCategory', 'host', 'port', 'ipRange'
	Value  string `json:"value,omitempty"`  // The value of what needs to get blocked. Format of the value varies depending on type of the firewall rule selected.
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
type RequestWirelessUpdateNetworkWirelessSSIDHotspot20NaiRealmsMethodsAuthenticationTypes struct {
	Credentials                  []string `json:"credentials,omitempty"`                  //
	EapinnerAuthentication       []string `json:"eapInnerAuthentication,omitempty"`       //
	NonEapinnerAuthentication    []string `json:"nonEapInnerAuthentication,omitempty"`    //
	TunneledEapMethodCredentials []string `json:"tunneledEapMethodCredentials,omitempty"` //
}
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

//GetNetworkWirelessFailedConnections List of all failed client connection events on this network in a given time range
/* List of all failed client connection events on this network in a given time range

@param networkID networkId path parameter. Network ID
@param getNetworkWirelessFailedConnectionsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-failed-connections
*/
func (s *WirelessService) GetNetworkWirelessFailedConnections(networkID string, getNetworkWirelessFailedConnectionsQueryParams *GetNetworkWirelessFailedConnectionsQueryParams) (*ResponseWirelessGetNetworkWirelessFailedConnections, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/failedConnections"
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

//GetNetworkWirelessRfProfiles List the non-basic RF profiles for this network
/* List the non-basic RF profiles for this network

@param networkID networkId path parameter. Network ID
@param getNetworkWirelessRfProfilesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-wireless-rf-profiles
*/
func (s *WirelessService) GetNetworkWirelessRfProfiles(networkID string, getNetworkWirelessRfProfilesQueryParams *GetNetworkWirelessRfProfilesQueryParams) (*ResponseWirelessGetNetworkWirelessRfProfiles, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/rfProfiles"
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

//GetOrganizationWirelessDevicesEthernetStatuses Endpoint to see power status for wireless devices
/* Endpoint to see power status for wireless devices

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessDevicesEthernetStatusesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-wireless-devices-ethernet-statuses
*/
func (s *WirelessService) GetOrganizationWirelessDevicesEthernetStatuses(organizationID string, getOrganizationWirelessDevicesEthernetStatusesQueryParams *GetOrganizationWirelessDevicesEthernetStatusesQueryParams) (*ResponseWirelessGetOrganizationWirelessDevicesEthernetStatuses, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wireless/devices/ethernet/statuses"
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

//CreateNetworkWirelessRfProfile Creates new RF profile for this network
/* Creates new RF profile for this network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-network-wireless-rf-profile
*/

func (s *WirelessService) CreateNetworkWirelessRfProfile(networkID string, requestWirelessCreateNetworkWirelessRfProfile *RequestWirelessCreateNetworkWirelessRfProfile) (*ResponseWirelessCreateNetworkWirelessRfProfile, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/rfProfiles"
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

func (s *WirelessService) CreateNetworkWirelessSSIDIDentityPsk(networkID string, number string, requestWirelessCreateNetworkWirelessSsidIdentityPsk *RequestWirelessCreateNetworkWirelessSSIDIDentityPsk) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ssids/{number}/identityPsks"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{number}", fmt.Sprintf("%v", number), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessCreateNetworkWirelessSsidIdentityPsk).
		// SetResult(&ResponseWirelessCreateNetworkWirelessSSIDIDentityPsk{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateNetworkWirelessSsidIdentityPsk")
	}

	return response, err

}

//UpdateDeviceWirelessBluetoothSettings Update the bluetooth settings for a wireless device
/* Update the bluetooth settings for a wireless device

@param serial serial path parameter.
*/
func (s *WirelessService) UpdateDeviceWirelessBluetoothSettings(serial string, requestWirelessUpdateDeviceWirelessBluetoothSettings *RequestWirelessUpdateDeviceWirelessBluetoothSettings) (*ResponseWirelessUpdateDeviceWirelessBluetoothSettings, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/wireless/bluetooth/settings"
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
func (s *WirelessService) UpdateNetworkWirelessBilling(networkID string, requestWirelessUpdateNetworkWirelessBilling *RequestWirelessUpdateNetworkWirelessBilling) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/billing"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateNetworkWirelessBilling).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkWirelessBilling")
	}

	return response, err

}

//UpdateNetworkWirelessBluetoothSettings Update the Bluetooth settings for a network
/* Update the Bluetooth settings for a network. See the docs page for
Bluetooth settings
.

@param networkID networkId path parameter. Network ID
*/
func (s *WirelessService) UpdateNetworkWirelessBluetoothSettings(networkID string, requestWirelessUpdateNetworkWirelessBluetoothSettings *RequestWirelessUpdateNetworkWirelessBluetoothSettings) (*ResponseWirelessUpdateNetworkWirelessBluetoothSettings, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/bluetooth/settings"
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

//UpdateNetworkWirelessRfProfile Updates specified RF profile for this network
/* Updates specified RF profile for this network

@param networkID networkId path parameter. Network ID
@param rfProfileID rfProfileId path parameter. Rf profile ID
*/
func (s *WirelessService) UpdateNetworkWirelessRfProfile(networkID string, rfProfileID string, requestWirelessUpdateNetworkWirelessRfProfile *RequestWirelessUpdateNetworkWirelessRfProfile) (*ResponseWirelessUpdateNetworkWirelessRfProfile, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/rfProfiles/{rfProfileId}"
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
func (s *WirelessService) UpdateNetworkWirelessSSID(networkID string, number string, requestWirelessUpdateNetworkWirelessSsid *RequestWirelessUpdateNetworkWirelessSSID) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ssids/{number}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{number}", fmt.Sprintf("%v", number), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateNetworkWirelessSsid).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkWirelessSsid")
	}

	return response, err

}

//UpdateNetworkWirelessSSIDBonjourForwarding Update the bonjour forwarding setting and rules for the SSID
/* Update the bonjour forwarding setting and rules for the SSID

@param networkID networkId path parameter. Network ID
@param number number path parameter.
*/
func (s *WirelessService) UpdateNetworkWirelessSSIDBonjourForwarding(networkID string, number string, requestWirelessUpdateNetworkWirelessSsidBonjourForwarding *RequestWirelessUpdateNetworkWirelessSSIDBonjourForwarding) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ssids/{number}/bonjourForwarding"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{number}", fmt.Sprintf("%v", number), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateNetworkWirelessSsidBonjourForwarding).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkWirelessSsidBonjourForwarding")
	}

	return response, err

}

//UpdateNetworkWirelessSSIDDeviceTypeGroupPolicies Update the device type group policies for the SSID
/* Update the device type group policies for the SSID

@param networkID networkId path parameter. Network ID
@param number number path parameter.
*/
func (s *WirelessService) UpdateNetworkWirelessSSIDDeviceTypeGroupPolicies(networkID string, number string, requestWirelessUpdateNetworkWirelessSsidDeviceTypeGroupPolicies *RequestWirelessUpdateNetworkWirelessSSIDDeviceTypeGroupPolicies) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ssids/{number}/deviceTypeGroupPolicies"
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
func (s *WirelessService) UpdateNetworkWirelessSSIDFirewallL3FirewallRules(networkID string, number string, requestWirelessUpdateNetworkWirelessSsidFirewallL3FirewallRules *RequestWirelessUpdateNetworkWirelessSSIDFirewallL3FirewallRules) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ssids/{number}/firewall/l3FirewallRules"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{number}", fmt.Sprintf("%v", number), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateNetworkWirelessSsidFirewallL3FirewallRules).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkWirelessSsidFirewallL3FirewallRules")
	}

	return response, err

}

//UpdateNetworkWirelessSSIDFirewallL7FirewallRules Update the L7 firewall rules of an SSID on an MR network
/* Update the L7 firewall rules of an SSID on an MR network

@param networkID networkId path parameter. Network ID
@param number number path parameter.
*/
func (s *WirelessService) UpdateNetworkWirelessSSIDFirewallL7FirewallRules(networkID string, number string, requestWirelessUpdateNetworkWirelessSsidFirewallL7FirewallRules *RequestWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRules) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ssids/{number}/firewall/l7FirewallRules"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{number}", fmt.Sprintf("%v", number), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateNetworkWirelessSsidFirewallL7FirewallRules).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkWirelessSsidFirewallL7FirewallRules")
	}

	return response, err

}

//UpdateNetworkWirelessSSIDHotspot20 Update the Hotspot 2.0 settings of an SSID
/* Update the Hotspot 2.0 settings of an SSID

@param networkID networkId path parameter. Network ID
@param number number path parameter.
*/
func (s *WirelessService) UpdateNetworkWirelessSSIDHotspot20(networkID string, number string, requestWirelessUpdateNetworkWirelessSsidHotspot20 *RequestWirelessUpdateNetworkWirelessSSIDHotspot20) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ssids/{number}/hotspot20"
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
func (s *WirelessService) UpdateNetworkWirelessSSIDIDentityPsk(networkID string, number string, identityPskID string, requestWirelessUpdateNetworkWirelessSsidIdentityPsk *RequestWirelessUpdateNetworkWirelessSSIDIDentityPsk) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ssids/{number}/identityPsks/{identityPskId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{number}", fmt.Sprintf("%v", number), -1)
	path = strings.Replace(path, "{identityPskId}", fmt.Sprintf("%v", identityPskID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateNetworkWirelessSsidIdentityPsk).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkWirelessSsidIdentityPsk")
	}

	return response, err

}

//UpdateNetworkWirelessSSIDSchedules Update the outage schedule for the SSID
/* Update the outage schedule for the SSID

@param networkID networkId path parameter. Network ID
@param number number path parameter.
*/
func (s *WirelessService) UpdateNetworkWirelessSSIDSchedules(networkID string, number string, requestWirelessUpdateNetworkWirelessSsidSchedules *RequestWirelessUpdateNetworkWirelessSSIDSchedules) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ssids/{number}/schedules"
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

//UpdateNetworkWirelessSSIDTrafficShapingRules Update the traffic shaping settings for an SSID on an MR network
/* Update the traffic shaping settings for an SSID on an MR network

@param networkID networkId path parameter. Network ID
@param number number path parameter.
*/
func (s *WirelessService) UpdateNetworkWirelessSSIDTrafficShapingRules(networkID string, number string, requestWirelessUpdateNetworkWirelessSsidTrafficShapingRules *RequestWirelessUpdateNetworkWirelessSSIDTrafficShapingRules) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ssids/{number}/trafficShaping/rules"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{number}", fmt.Sprintf("%v", number), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateNetworkWirelessSsidTrafficShapingRules).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkWirelessSsidTrafficShapingRules")
	}

	return response, err

}

//UpdateNetworkWirelessSSIDVpn Update the VPN settings for the SSID
/* Update the VPN settings for the SSID

@param networkID networkId path parameter. Network ID
@param number number path parameter.
*/
func (s *WirelessService) UpdateNetworkWirelessSSIDVpn(networkID string, number string, requestWirelessUpdateNetworkWirelessSsidVpn *RequestWirelessUpdateNetworkWirelessSSIDVpn) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/wireless/ssids/{number}/vpn"
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

//DeleteNetworkWirelessRfProfile Delete a RF Profile
/* Delete a RF Profile

@param networkID networkId path parameter. Network ID
@param rfProfileID rfProfileId path parameter. Rf profile ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-network-wireless-rf-profile
*/
func (s *WirelessService) DeleteNetworkWirelessRfProfile(networkID string, rfProfileID string) (*resty.Response, error) {
	//networkID string,rfProfileID string
	path := "/api/v1/networks/{networkId}/wireless/rfProfiles/{rfProfileId}"
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
