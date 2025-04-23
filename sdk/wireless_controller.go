package meraki

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type WirelessControllerService service

type GetOrganizationWirelessControllerAvailabilitiesChangeHistoryQueryParams struct {
	Serials       []string `url:"serials[],omitempty"`     //Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches.
	T0            string   `url:"t0,omitempty"`            //The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	T1            string   `url:"t1,omitempty"`            //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan      float64  `url:"timespan,omitempty"`      //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.
	PerPage       int      `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter string   `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string   `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
}
type GetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalQueryParams struct {
	NetworkIDs    []string `url:"networkIds[],omitempty"`  //Optional parameter to filter wireless LAN controllers by network ID. This filter uses multiple exact matches.
	Serials       []string `url:"serials[],omitempty"`     //Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches.
	T0            string   `url:"t0,omitempty"`            //The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	T1            string   `url:"t1,omitempty"`            //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan      float64  `url:"timespan,omitempty"`      //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.
	PerPage       int      `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter string   `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string   `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	Resolution    int      `url:"resolution,omitempty"`    //The time resolution in seconds for returned data. The valid resolutions are: 300, 600, 1200, 3600, 14400, 86400. The default is 86400.
}
type GetOrganizationWirelessControllerConnectionsQueryParams struct {
	NetworkIDs        []string `url:"networkIds[],omitempty"`        //Optional parameter to filter access points by network ID. This filter uses multiple exact matches.
	ControllerSerials []string `url:"controllerSerials[],omitempty"` //Optional parameter to filter access points by its controller cloud ID. This filter uses multiple exact matches.
	PerPage           int      `url:"perPage,omitempty"`             //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter     string   `url:"startingAfter,omitempty"`       //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore      string   `url:"endingBefore,omitempty"`        //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
}
type GetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceQueryParams struct {
	Serials       []string `url:"serials[],omitempty"`     //Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches.
	T0            string   `url:"t0,omitempty"`            //The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	T1            string   `url:"t1,omitempty"`            //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan      float64  `url:"timespan,omitempty"`      //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.
	PerPage       int      `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter string   `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string   `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
}
type GetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceQueryParams struct {
	Serials                         []string `url:"serials[],omitempty"`                       //Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches.
	IncludeInterfacesWithoutChanges bool     `url:"includeInterfacesWithoutChanges,omitempty"` //By default, interfaces without changes are omitted from the response for brevity. If you want to include the interfaces even if they have no changes, set to true. (default: false)
	T0                              string   `url:"t0,omitempty"`                              //The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	T1                              string   `url:"t1,omitempty"`                              //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan                        float64  `url:"timespan,omitempty"`                        //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.
	PerPage                         int      `url:"perPage,omitempty"`                         //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter                   string   `url:"startingAfter,omitempty"`                   //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore                    string   `url:"endingBefore,omitempty"`                    //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
}
type GetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalQueryParams struct {
	Serials       []string `url:"serials[],omitempty"`     //Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches.
	T0            string   `url:"t0,omitempty"`            //The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	T1            string   `url:"t1,omitempty"`            //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan      float64  `url:"timespan,omitempty"`      //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.
	PerPage       int      `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter string   `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string   `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
}
type GetOrganizationWirelessControllerDevicesInterfacesL3ByDeviceQueryParams struct {
	Serials       []string `url:"serials[],omitempty"`     //Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches.
	T0            string   `url:"t0,omitempty"`            //The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	T1            string   `url:"t1,omitempty"`            //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan      float64  `url:"timespan,omitempty"`      //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.
	PerPage       int      `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter string   `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string   `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
}
type GetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceQueryParams struct {
	Serials                         []string `url:"serials[],omitempty"`                       //Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches.
	IncludeInterfacesWithoutChanges bool     `url:"includeInterfacesWithoutChanges,omitempty"` //By default, interfaces without changes are omitted from the response for brevity. If you want to include the interfaces even if they have no changes, set to true. (default: false)
	T0                              string   `url:"t0,omitempty"`                              //The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	T1                              string   `url:"t1,omitempty"`                              //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan                        float64  `url:"timespan,omitempty"`                        //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.
	PerPage                         int      `url:"perPage,omitempty"`                         //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter                   string   `url:"startingAfter,omitempty"`                   //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore                    string   `url:"endingBefore,omitempty"`                    //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
}
type GetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalQueryParams struct {
	Serials       []string `url:"serials[],omitempty"`     //Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches.
	T0            string   `url:"t0,omitempty"`            //The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	T1            string   `url:"t1,omitempty"`            //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan      float64  `url:"timespan,omitempty"`      //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.
	PerPage       int      `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter string   `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string   `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
}
type GetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceQueryParams struct {
	Serials       []string `url:"serials[],omitempty"`     //Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches.
	Names         []string `url:"names[],omitempty"`       //Optional parameter to filter wireless LAN controller by its interface name. This filter uses multiple exact matches.
	T0            string   `url:"t0,omitempty"`            //The beginning of the timespan for the data. The maximum lookback period is 1 day from today.
	T1            string   `url:"t1,omitempty"`            //The end of the timespan for the data. t1 can be a maximum of 1 day after t0.
	Timespan      float64  `url:"timespan,omitempty"`      //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 1 day. The default is 1 hour.
	PerPage       int      `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter string   `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string   `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
}
type GetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalQueryParams struct {
	Serials       []string `url:"serials[],omitempty"`     //Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches.
	Names         []string `url:"names[],omitempty"`       //Optional parameter to filter wireless LAN controller by its interface name. This filter uses multiple exact matches.
	T0            string   `url:"t0,omitempty"`            //The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	T1            string   `url:"t1,omitempty"`            //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan      float64  `url:"timespan,omitempty"`      //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.
	PerPage       int      `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter string   `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string   `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
}
type GetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryQueryParams struct {
	Serials       []string `url:"serials[],omitempty"`     //Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches.
	T0            string   `url:"t0,omitempty"`            //The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	T1            string   `url:"t1,omitempty"`            //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan      float64  `url:"timespan,omitempty"`      //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.
	PerPage       int      `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter string   `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string   `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
}
type GetOrganizationWirelessControllerDevicesRedundancyStatusesQueryParams struct {
	Serials       []string `url:"serials[],omitempty"`     //Optional parameter to filter wireless LAN controller by its cloud IDs. This filter uses multiple exact matches.
	PerPage       int      `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter string   `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string   `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
}
type GetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalQueryParams struct {
	Serials       []string `url:"serials[],omitempty"`     //Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches.
	T0            string   `url:"t0,omitempty"`            //The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	T1            string   `url:"t1,omitempty"`            //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan      float64  `url:"timespan,omitempty"`      //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.
	PerPage       int      `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter string   `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string   `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
}
type GetOrganizationWirelessControllerOverviewByDeviceQueryParams struct {
	NetworkIDs    []string `url:"networkIds[],omitempty"`  //Optional parameter to filter wireless LAN controllers by network ID. This filter uses multiple exact matches.
	Serials       []string `url:"serials[],omitempty"`     //Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches.
	PerPage       int      `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter string   `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string   `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
}

type ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistory struct {
	Items *[]ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistoryItems `json:"items,omitempty"` // Wireless LAN controller connectivity information
	Meta  *ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistoryMeta    `json:"meta,omitempty"`  // Metadata relevant to the paginated dataset
}
type ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistoryItems struct {
	Changes *[]ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistoryItemsChanges `json:"changes,omitempty"` // Connectivity information of a wireless LAN controller
	Serial  string                                                                                                `json:"serial,omitempty"`  // Wireless LAN controller cloud ID
}
type ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistoryItemsChanges struct {
	EndTs   string `json:"endTs,omitempty"`   // The end time(UTC seconds) of the wireless LAN controller connectivity status change. This attribute is set to be null by default if there's no need to assign.
	StartTs string `json:"startTs,omitempty"` // The start time(UTC seconds) of the wireless LAN controller connectivity status change
	Status  string `json:"status,omitempty"`  // The wireless LAN controller connectivity status
}
type ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistoryMeta struct {
	Counts *ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistoryMetaCounts `json:"counts,omitempty"` // Counts relating to the paginated dataset
}
type ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistoryMetaCounts struct {
	Items *ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistoryMetaCountsItems `json:"items,omitempty"` // Counts relating to the paginated items
}
type ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistoryMetaCountsItems struct {
	Remaining *int `json:"remaining,omitempty"` // The number of items in the dataset that are available on subsequent pages
	Total     *int `json:"total,omitempty"`     // The total number of items in the dataset
}
type ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByInterval struct {
	Items *[]ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalItems `json:"items,omitempty"` // Overview history of wireless LAN controllers
	Meta  *ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalMeta    `json:"meta,omitempty"`  // Metadata relevant to the paginated dataset
}
type ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalItems struct {
	Network  *ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalItemsNetwork    `json:"network,omitempty"`  // Wireless LAN controller network
	Readings *[]ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalItemsReadings `json:"readings,omitempty"` // Overview history of a wireless LAN controller
	Serial   string                                                                                                              `json:"serial,omitempty"`   // Wireless LAN controller cloud ID
}
type ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalItemsNetwork struct {
	ID string `json:"id,omitempty"` // Wireless LAN controller network ID
}
type ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalItemsReadings struct {
	Counts  *ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalItemsReadingsCounts `json:"counts,omitempty"`  // Client counts
	EndTs   string                                                                                                                  `json:"endTs,omitempty"`   // The end time of the query range
	StartTs string                                                                                                                  `json:"startTs,omitempty"` // The start time of the query range
}
type ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalItemsReadingsCounts struct {
	ByStatus *ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalItemsReadingsCountsByStatus `json:"byStatus,omitempty"` // Client counts by its status
}
type ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalItemsReadingsCountsByStatus struct {
	Online *int `json:"online,omitempty"` // Number of connected clients
}
type ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalMeta struct {
	Counts *ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalMetaCounts `json:"counts,omitempty"` // Counts relating to the paginated dataset
}
type ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalMetaCounts struct {
	Items *ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalMetaCountsItems `json:"items,omitempty"` // Counts relating to the paginated items
}
type ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalMetaCountsItems struct {
	Remaining *int `json:"remaining,omitempty"` // The number of items in the dataset that are available on subsequent pages
	Total     *int `json:"total,omitempty"`     // The total number of items in the dataset
}
type ResponseWirelessControllerGetOrganizationWirelessControllerConnections struct {
	Items *[]ResponseWirelessControllerGetOrganizationWirelessControllerConnectionsItems `json:"items,omitempty"` // Access points associated with Wireless LAN controllers
	Meta  *ResponseWirelessControllerGetOrganizationWirelessControllerConnectionsMeta    `json:"meta,omitempty"`  // Metadata relevant to the paginated dataset
}
type ResponseWirelessControllerGetOrganizationWirelessControllerConnectionsItems struct {
	Controller *ResponseWirelessControllerGetOrganizationWirelessControllerConnectionsItemsController `json:"controller,omitempty"` // Associated wireless LAN controller
	Network    *ResponseWirelessControllerGetOrganizationWirelessControllerConnectionsItemsNetwork    `json:"network,omitempty"`    // Access points network
	Serial     string                                                                                 `json:"serial,omitempty"`     // Access points cloud ID
}
type ResponseWirelessControllerGetOrganizationWirelessControllerConnectionsItemsController struct {
	Serial string `json:"serial,omitempty"` // Associated wireless LAN controller cloud ID
}
type ResponseWirelessControllerGetOrganizationWirelessControllerConnectionsItemsNetwork struct {
	ID   string `json:"id,omitempty"`   // Access points network ID
	Name string `json:"name,omitempty"` // Access points network name
	URL  string `json:"url,omitempty"`  // Access points network URL
}
type ResponseWirelessControllerGetOrganizationWirelessControllerConnectionsMeta struct {
	Counts *ResponseWirelessControllerGetOrganizationWirelessControllerConnectionsMetaCounts `json:"counts,omitempty"` // Counts relating to the paginated dataset
}
type ResponseWirelessControllerGetOrganizationWirelessControllerConnectionsMetaCounts struct {
	Items *ResponseWirelessControllerGetOrganizationWirelessControllerConnectionsMetaCountsItems `json:"items,omitempty"` // Counts relating to the paginated items
}
type ResponseWirelessControllerGetOrganizationWirelessControllerConnectionsMetaCountsItems struct {
	Remaining *int `json:"remaining,omitempty"` // The number of items in the dataset that are available on subsequent pages
	Total     *int `json:"total,omitempty"`     // The total number of items in the dataset
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDevice struct {
	Items *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceItems `json:"items,omitempty"` // Wireless LAN controller L2 interfaces
	Meta  *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceMeta    `json:"meta,omitempty"`  // Metadata relevant to the paginated dataset
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceItems struct {
	Interfaces *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceItemsInterfaces `json:"interfaces,omitempty"` // Layer 2 interfaces belongs to the wireless LAN controller
	Serial     string                                                                                                   `json:"serial,omitempty"`     // The cloud ID of the wireless LAN controller
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceItemsInterfaces struct {
	ChannelGroup     *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceItemsInterfacesChannelGroup `json:"channelGroup,omitempty"`     // The channel group of this wireless LAN controller interface
	Description      string                                                                                                             `json:"description,omitempty"`      // The description of the wireless LAN controller interface
	Enabled          *bool                                                                                                              `json:"enabled,omitempty"`          // The status of the wireless LAN controller interface
	IsRedundancyPort *bool                                                                                                              `json:"isRedundancyPort,omitempty"` // Indicate whether the interface is a redundancy port used to perform HA role negotiation
	IsUplink         *bool                                                                                                              `json:"isUplink,omitempty"`         // Indicate whether the interface is uplink
	LinkNegotiation  string                                                                                                             `json:"linkNegotiation,omitempty"`  // The interface negotiation mode
	Mac              string                                                                                                             `json:"mac,omitempty"`              // The MAC address of the wireless LAN controller interface
	Module           *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceItemsInterfacesModule       `json:"module,omitempty"`           // The module of this wireless LAN controller interface
	Name             string                                                                                                             `json:"name,omitempty"`             // The name of the wireless LAN controller interface
	Speed            string                                                                                                             `json:"speed,omitempty"`            // The current data transfer rate which the interface is operating at. enum = [1 Gbps, 2 Gbps, 5 Gbps, 10 Gbps, 20 Gbps, 40 Gbps, 100 Gbps]
	Status           string                                                                                                             `json:"status,omitempty"`           // The status of the wireless LAN controller interface
	VLAN             *int                                                                                                               `json:"vlan,omitempty"`             // The VLAN of the switch port. For a trunk port, this is the native VLAN. A null value will clear the value set for trunk ports.
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceItemsInterfacesChannelGroup struct {
	Number *int `json:"number,omitempty"` // The interface channel group number
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceItemsInterfacesModule struct {
	Model string `json:"model,omitempty"` // The module type of this wireless LAN controller interface
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceMeta struct {
	Counts *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceMetaCounts `json:"counts,omitempty"` // Counts relating to the paginated dataset
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceMetaCounts struct {
	Items *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceMetaCountsItems `json:"items,omitempty"` // Counts relating to the paginated items
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceMetaCountsItems struct {
	Remaining *int `json:"remaining,omitempty"` // The number of items in the dataset that are available on subsequent pages
	Total     *int `json:"total,omitempty"`     // The total number of items in the dataset
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice struct {
	Items *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceItems `json:"items,omitempty"` // Wireless LAN controller layer 2 interfaces historical status
	Meta  *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceMeta    `json:"meta,omitempty"`  // Metadata relevant to the paginated dataset
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceItems struct {
	Interfaces *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceItemsInterfaces `json:"interfaces,omitempty"` // layer 2 interfaces belongs to the wireless LAN controller
	Serial     string                                                                                                                        `json:"serial,omitempty"`     // The cloud ID of the wireless LAN controller
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceItemsInterfaces struct {
	Changes *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceItemsInterfacesChanges `json:"changes,omitempty"` // The statuses of layer 2 interfaces of the wireless LAN controller
	Mac     string                                                                                                                               `json:"mac,omitempty"`     // The MAC address of the wireless LAN controller interface
	Name    string                                                                                                                               `json:"name,omitempty"`    // The name of the wireless LAN controller interface
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceItemsInterfacesChanges struct {
	Errors   []string `json:"errors,omitempty"`   // All errors present on the port
	Status   string   `json:"status,omitempty"`   // The status of the interface
	Ts       string   `json:"ts,omitempty"`       // The timestamp of current status of the interface
	Warnings []string `json:"warnings,omitempty"` // All warnings present on the port
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceMeta struct {
	Counts *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceMetaCounts `json:"counts,omitempty"` // Counts relating to the paginated dataset
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceMetaCounts struct {
	Items *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceMetaCountsItems `json:"items,omitempty"` // Counts relating to the paginated items
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceMetaCountsItems struct {
	Remaining *int `json:"remaining,omitempty"` // The number of items in the dataset that are available on subsequent pages
	Total     *int `json:"total,omitempty"`     // The total number of items in the dataset
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval struct {
	Items *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalItems `json:"items,omitempty"` // Wireless LAN controller layer 2 interfaces usage
	Meta  *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalMeta    `json:"meta,omitempty"`  // Metadata relevant to the paginated dataset
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalItems struct {
	Readings *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalItemsReadings `json:"readings,omitempty"` // The usages of layer 2 interfaces of the wireless LAN controller. Usage is in bytes
	Serial   string                                                                                                               `json:"serial,omitempty"`   // The cloud ID of the wireless LAN controller
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalItemsReadings struct {
	Mac  string `json:"mac,omitempty"`  // The MAC address of the wireless controller interface
	Name string `json:"name,omitempty"` // The name of the wireless LAN controller interface
	Recv *int   `json:"recv,omitempty"` // The volume of data, in bytes/sec, received by wireless controller interface
	Send *int   `json:"send,omitempty"` // The volume of data, in bytes/sec, transmitted by wireless controller interface
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalMeta struct {
	Counts *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalMetaCounts `json:"counts,omitempty"` // Counts relating to the paginated dataset
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalMetaCounts struct {
	Items *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalMetaCountsItems `json:"items,omitempty"` // Counts relating to the paginated items
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalMetaCountsItems struct {
	Remaining *int `json:"remaining,omitempty"` // The number of items in the dataset that are available on subsequent pages
	Total     *int `json:"total,omitempty"`     // The total number of items in the dataset
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3ByDevice struct {
	Items *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3ByDeviceItems `json:"items,omitempty"` // Wireless LAN controller L3 interfaces
	Meta  *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3ByDeviceMeta    `json:"meta,omitempty"`  // Metadata relevant to the paginated dataset
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3ByDeviceItems struct {
	Interfaces *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3ByDeviceItemsInterfaces `json:"interfaces,omitempty"` // Layer 3 interfaces belongs to the wireless LAN controller
	Serial     string                                                                                                   `json:"serial,omitempty"`     // The cloud ID of the wireless LAN controller
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3ByDeviceItemsInterfaces struct {
	Addresses       *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3ByDeviceItemsInterfacesAddresses  `json:"addresses,omitempty"`       // Available addresses for the interface.
	ChannelGroup    *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3ByDeviceItemsInterfacesChannelGroup `json:"channelGroup,omitempty"`    // The channel group of this wireless LAN controller interface
	Description     string                                                                                                             `json:"description,omitempty"`     // The description of the wireless LAN controller interface
	IsUplink        *bool                                                                                                              `json:"isUplink,omitempty"`        // Indicate whether the interface is uplink
	LinkNegotiation string                                                                                                             `json:"linkNegotiation,omitempty"` // The interface negotiation mode
	Mac             string                                                                                                             `json:"mac,omitempty"`             // The MAC address of the wireless LAN controller interface
	Module          *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3ByDeviceItemsInterfacesModule       `json:"module,omitempty"`          // The module of this wireless LAN controller interface
	Name            string                                                                                                             `json:"name,omitempty"`            // The name of the wireless LAN controller interface
	Speed           string                                                                                                             `json:"speed,omitempty"`           // The current data transfer rate which the interface is operating at. enum = [1 Gbps, 2 Gbps, 5 Gbps, 10 Gbps, 20 Gbps, 40 Gbps, 100 Gbps]
	Status          string                                                                                                             `json:"status,omitempty"`          // The status of the wireless LAN controller interface
	VLAN            *int                                                                                                               `json:"vlan,omitempty"`            // The VLAN of the switch port. For a trunk port, this is the native VLAN. A null value will clear the value set for trunk ports.
	Vrf             *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3ByDeviceItemsInterfacesVrf          `json:"vrf,omitempty"`             // The virtual routing and forwarding (VRF) for the wireless LAN controller interface
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3ByDeviceItemsInterfacesAddresses struct {
	Address  string `json:"address,omitempty"`  // The address of the wireless LAN controller interface
	Protocol string `json:"protocol,omitempty"` // Type of address for the device uplink. Available options are: ipv4, ipv6. enum = [ipv4, ipv6]
	Subnet   string `json:"subnet,omitempty"`   // The address of the wireless LAN controller interface
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3ByDeviceItemsInterfacesChannelGroup struct {
	Number *int `json:"number,omitempty"` // The interface channel group number
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3ByDeviceItemsInterfacesModule struct {
	Model string `json:"model,omitempty"` // The module type of this wireless LAN controller interface
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3ByDeviceItemsInterfacesVrf struct {
	Name string `json:"name,omitempty"` // The virtual routing and forwarding (VRF) name
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3ByDeviceMeta struct {
	Counts *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3ByDeviceMetaCounts `json:"counts,omitempty"` // Counts relating to the paginated dataset
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3ByDeviceMetaCounts struct {
	Items *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3ByDeviceMetaCountsItems `json:"items,omitempty"` // Counts relating to the paginated items
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3ByDeviceMetaCountsItems struct {
	Remaining *int `json:"remaining,omitempty"` // The number of items in the dataset that are available on subsequent pages
	Total     *int `json:"total,omitempty"`     // The total number of items in the dataset
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice struct {
	Items *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceItems `json:"items,omitempty"` // Wireless LAN controller layer 3 interfaces historical status
	Meta  *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceMeta    `json:"meta,omitempty"`  // Metadata relevant to the paginated dataset
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceItems struct {
	Interfaces *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceItemsInterfaces `json:"interfaces,omitempty"` // layer 3 interfaces belongs to the wireless LAN controller
	Serial     string                                                                                                                        `json:"serial,omitempty"`     // The cloud ID of the wireless LAN controller
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceItemsInterfaces struct {
	Changes *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceItemsInterfacesChanges `json:"changes,omitempty"` // The statuses of layer 3 interfaces of the wireless LAN controller
	Mac     string                                                                                                                               `json:"mac,omitempty"`     // The MAC address of the wireless LAN controller interface
	Name    string                                                                                                                               `json:"name,omitempty"`    // The name of the wireless LAN controller interface
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceItemsInterfacesChanges struct {
	Errors   []string `json:"errors,omitempty"`   // All errors present on the port
	Status   string   `json:"status,omitempty"`   // The status of the interface
	Ts       string   `json:"ts,omitempty"`       // The timestamp of current status of the interface
	Warnings []string `json:"warnings,omitempty"` // All warnings present on the port
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceMeta struct {
	Counts *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceMetaCounts `json:"counts,omitempty"` // Counts relating to the paginated dataset
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceMetaCounts struct {
	Items *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceMetaCountsItems `json:"items,omitempty"` // Counts relating to the paginated items
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceMetaCountsItems struct {
	Remaining *int `json:"remaining,omitempty"` // The number of items in the dataset that are available on subsequent pages
	Total     *int `json:"total,omitempty"`     // The total number of items in the dataset
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByInterval struct {
	Items *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalItems `json:"items,omitempty"` // Wireless LAN controller layer 3 interfaces usage
	Meta  *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalMeta    `json:"meta,omitempty"`  // Metadata relevant to the paginated dataset
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalItems struct {
	Readings *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalItemsReadings `json:"readings,omitempty"` // The usages of layer 3 interfaces of the wireless LAN controller. Usage is in bytes
	Serial   string                                                                                                               `json:"serial,omitempty"`   // The cloud ID of the wireless LAN controller
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalItemsReadings struct {
	Mac  string `json:"mac,omitempty"`  // The MAC address of the wireless controller interface
	Name string `json:"name,omitempty"` // The name of the wireless LAN controller interface
	Recv *int   `json:"recv,omitempty"` // The volume of data, in bytes/sec, received by wireless controller interface
	Send *int   `json:"send,omitempty"` // The volume of data, in bytes/sec, transmitted by wireless controller interface
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalMeta struct {
	Counts *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalMetaCounts `json:"counts,omitempty"` // Counts relating to the paginated dataset
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalMetaCounts struct {
	Items *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalMetaCountsItems `json:"items,omitempty"` // Counts relating to the paginated items
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalMetaCountsItems struct {
	Remaining *int `json:"remaining,omitempty"` // The number of items in the dataset that are available on subsequent pages
	Total     *int `json:"total,omitempty"`     // The total number of items in the dataset
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice struct {
	Items *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceItems `json:"items,omitempty"` // Wireless LAN controller interfaces packets statuses
	Meta  *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceMeta    `json:"meta,omitempty"`  // Metadata relevant to the paginated dataset
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceItems struct {
	Interfaces *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceItemsInterfaces `json:"interfaces,omitempty"` // Interfaces belongs to the wireless LAN controller
	Serial     string                                                                                                                `json:"serial,omitempty"`     // The cloud ID of the wireless LAN controller
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceItemsInterfaces struct {
	Name     string                                                                                                                        `json:"name,omitempty"`     // The name of the wireless LAN controller interface
	Readings *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceItemsInterfacesReadings `json:"readings,omitempty"` // The status of packets counter on the interfaces of the wireless LAN controller
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceItemsInterfacesReadings struct {
	Name  string                                                                                                                          `json:"name,omitempty"`  // The type of packets being counted
	Rate  *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceItemsInterfacesReadingsRate `json:"rate,omitempty"`  // The interface packet rates measured in packets per second
	Recv  *int                                                                                                                            `json:"recv,omitempty"`  // The total count of packets received by the interface during the timespan
	Send  *int                                                                                                                            `json:"send,omitempty"`  // The total count of packets sent by the interface during the timespan
	Total *int                                                                                                                            `json:"total,omitempty"` // The total count of sent and received packets during the timespan
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceItemsInterfacesReadingsRate struct {
	Recv  *int `json:"recv,omitempty"`  // The rate of packets received during the timespan
	Send  *int `json:"send,omitempty"`  // The rate of packets sent during the timespan
	Total *int `json:"total,omitempty"` // The rate of all packets sent and received during the timespan
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceMeta struct {
	Counts *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceMetaCounts `json:"counts,omitempty"` // Counts relating to the paginated dataset
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceMetaCounts struct {
	Items *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceMetaCountsItems `json:"items,omitempty"` // Counts relating to the paginated items
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceMetaCountsItems struct {
	Remaining *int `json:"remaining,omitempty"` // The number of items in the dataset that are available on subsequent pages
	Total     *int `json:"total,omitempty"`     // The total number of items in the dataset
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval struct {
	Items *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItems `json:"items,omitempty"` // Wireless LAN controller interfaces usage data
	Meta  *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalMeta    `json:"meta,omitempty"`  // Metadata relevant to the paginated dataset
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItems struct {
	Intervals *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItemsIntervals `json:"intervals,omitempty"` // Time interval snapshots of interfaces usage data of the wireless LAN controller
	Serial    string                                                                                                              `json:"serial,omitempty"`    // The cloud ID of the wireless LAN controller
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItemsIntervals struct {
	ByInterface *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItemsIntervalsByInterface `json:"byInterface,omitempty"` // The usage data on the interfaces of the wireless LAN controller
	EndTs       string                                                                                                                         `json:"endTs,omitempty"`       // The end time interval snapshots of the query range with iso8601 format
	Overall     *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItemsIntervalsOverall       `json:"overall,omitempty"`     // The overall usage of all queried interfaces of the wireless LAN controller
	StartTs     string                                                                                                                         `json:"startTs,omitempty"`     // The start time interval snapshots of the query range with iso8601 format
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItemsIntervalsByInterface struct {
	Name  string                                                                                                                            `json:"name,omitempty"`  // The name of the wireless LAN controller interface
	Usage *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItemsIntervalsByInterfaceUsage `json:"usage,omitempty"` // The usage on the interfaces of the wireless LAN controller
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItemsIntervalsByInterfaceUsage struct {
	Recv  *int `json:"recv,omitempty"`  // The received usage on the interface during the interval, unit is bit/sec
	Send  *int `json:"send,omitempty"`  // The sent usage on the interface during the interval, unit is bit/sec
	Total *int `json:"total,omitempty"` // The total usage on the interface during the interval, unit is bit/sec
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItemsIntervalsOverall struct {
	Recv  *int `json:"recv,omitempty"`  // The received usage of all queried interfaces during the interval, unit is bit/sec
	Send  *int `json:"send,omitempty"`  // The sent usage of all queried interfaces during the interval, unit is bit/sec
	Total *int `json:"total,omitempty"` // The total usage of all queried interfaces during the interval, unit is bit/sec
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalMeta struct {
	Counts *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalMetaCounts `json:"counts,omitempty"` // Counts relating to the paginated dataset
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalMetaCounts struct {
	Items *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalMetaCountsItems `json:"items,omitempty"` // Counts relating to the paginated items
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalMetaCountsItems struct {
	Remaining *int `json:"remaining,omitempty"` // The number of items in the dataset that are available on subsequent pages
	Total     *int `json:"total,omitempty"`     // The total number of items in the dataset
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistory []ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistory // Array of ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistory
type ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistory struct {
	Items *[]ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItems `json:"items,omitempty"` // Wireless LAN controller HA failover events
	Meta  *ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryMeta    `json:"meta,omitempty"`  // Metadata relevant to the paginated dataset
}
type ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItems struct {
	Active *ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItemsActive `json:"active,omitempty"` // Details about the active unit
	Failed *ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItemsFailed `json:"failed,omitempty"` // Details about the failed unit
	Reason string                                                                                                      `json:"reason,omitempty"` // Failover reason
	Serial string                                                                                                      `json:"serial,omitempty"` // Wireless LAN controller cloud ID
	Ts     string                                                                                                      `json:"ts,omitempty"`     // Failover time
}
type ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItemsActive struct {
	Chassis *ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItemsActiveChassis `json:"chassis,omitempty"` // Details about the active unit chassis
}
type ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItemsActiveChassis struct {
	Name string `json:"name,omitempty"` // The name of the active chassis unit
}
type ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItemsFailed struct {
	Chassis *ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItemsFailedChassis `json:"chassis,omitempty"` // Details about the failed unit chassis
}
type ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItemsFailedChassis struct {
	Name string `json:"name,omitempty"` // The name of the failed chassis unit
}
type ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryMeta struct {
	Counts *ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryMetaCounts `json:"counts,omitempty"` // Counts relating to the paginated dataset
}
type ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryMetaCounts struct {
	Items *ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryMetaCountsItems `json:"items,omitempty"` // Counts relating to the paginated items
}
type ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryMetaCountsItems struct {
	Remaining *int `json:"remaining,omitempty"` // The number of items in the dataset that are available on subsequent pages
	Total     *int `json:"total,omitempty"`     // The total number of items in the dataset
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatuses struct {
	Items *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesItems `json:"items,omitempty"` // Wireless LAN controller redundancy statuses
	Meta  *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesMeta    `json:"meta,omitempty"`  // Metadata relevant to the paginated dataset
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesItems struct {
	Enabled     *bool                                                                                              `json:"enabled,omitempty"`     // Wireless LAN controller redundancy enablement
	Failover    *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesItemsFailover `json:"failover,omitempty"`    // Wireless LAN controller failover information
	MobilityMac string                                                                                             `json:"mobilityMac,omitempty"` // Wireless LAN controller redundancy mobility mac
	Mode        string                                                                                             `json:"mode,omitempty"`        // Wireless LAN controller redundancy SSO (stateful switchover)
	Serial      string                                                                                             `json:"serial,omitempty"`      // Wireless LAN controller cloud ID
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesItemsFailover struct {
	Counts *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesItemsFailoverCounts `json:"counts,omitempty"` // Wireless LAN controller switchover counts
	Last   *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesItemsFailoverLast   `json:"last,omitempty"`   // Wireless LAN controller last failover information
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesItemsFailoverCounts struct {
	Total *int `json:"total,omitempty"` // Total number of failovers
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesItemsFailoverLast struct {
	Reason string `json:"reason,omitempty"` // Wireless LAN controller last redundancy switchover reason
	Ts     string `json:"ts,omitempty"`     // Wireless LAN controller last redundancy switchover time
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesMeta struct {
	Counts *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesMetaCounts `json:"counts,omitempty"` // Counts relating to the paginated dataset
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesMetaCounts struct {
	Items *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesMetaCountsItems `json:"items,omitempty"` // Counts relating to the paginated items
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesMetaCountsItems struct {
	Remaining *int `json:"remaining,omitempty"` // The number of items in the dataset that are available on subsequent pages
	Total     *int `json:"total,omitempty"`     // The total number of items in the dataset
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByInterval struct {
	Items *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItems `json:"items,omitempty"` // Wireless LAN controller CPU usage data
	Meta  *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalMeta    `json:"meta,omitempty"`  // Metadata relevant to the paginated dataset
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItems struct {
	Intervals *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervals `json:"intervals,omitempty"` // Time interval snapshots of CPU usage data of the wireless LAN controller
	Serial    string                                                                                                                `json:"serial,omitempty"`    // The cloud ID of the wireless LAN controller
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervals struct {
	ByCore  *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsByCore `json:"byCore,omitempty"`  // The CPU usage per core of the wireless LAN controller
	EndTs   string                                                                                                                      `json:"endTs,omitempty"`   // The end time of the query range  with iso8601 format
	Overall *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsOverall  `json:"overall,omitempty"` // The overall CPU usage of the wireless LAN controller
	StartTs string                                                                                                                      `json:"startTs,omitempty"` // The start time of the query range with iso8601 format
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsByCore struct {
	Name  string                                                                                                                         `json:"name,omitempty"`  // The CPU core name
	Usage *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsByCoreUsage `json:"usage,omitempty"` // The specific core CPU usage of the wireless LAN controller
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsByCoreUsage struct {
	Average *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsByCoreUsageAverage `json:"average,omitempty"` // The specific core average CPU usage of the wireless LAN controller
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsByCoreUsageAverage struct {
	Percentage *float64 `json:"percentage,omitempty"` // The specific core CPU usage percentage of the wireless LAN controller
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsOverall struct {
	Usage *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsOverallUsage `json:"usage,omitempty"` // The CPU usage of the wireless LAN controller
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsOverallUsage struct {
	Average *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsOverallUsageAverage `json:"average,omitempty"` // The average CPU usage of the wireless LAN controller
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsOverallUsageAverage struct {
	Percentage *float64 `json:"percentage,omitempty"` // The CPU usage percentage of the wireless LAN controller
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalMeta struct {
	Counts *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalMetaCounts `json:"counts,omitempty"` // Counts relating to the paginated dataset
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalMetaCounts struct {
	Items *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalMetaCountsItems `json:"items,omitempty"` // Counts relating to the paginated items
}
type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalMetaCountsItems struct {
	Remaining *int `json:"remaining,omitempty"` // The number of items in the dataset that are available on subsequent pages
	Total     *int `json:"total,omitempty"`     // The total number of items in the dataset
}
type ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDevice struct {
	Items *[]ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItems `json:"items,omitempty"` // Wireless LAN controller overview
	Meta  *ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceMeta    `json:"meta,omitempty"`  // Metadata relevant to the paginated dataset
}
type ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItems struct {
	Counts     *ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsCounts     `json:"counts,omitempty"`     // Wireless LAN controller client and access point counts
	Firmware   *ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsFirmware   `json:"firmware,omitempty"`   // Wireless LAN controller device firmware information
	Network    *ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsNetwork    `json:"network,omitempty"`    // Wireless LAN controller network
	Redundancy *ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsRedundancy `json:"redundancy,omitempty"` // Wireless LAN controller redundancy information
	Serial     string                                                                                      `json:"serial,omitempty"`     // Wireless LAN controller cloud ID
}
type ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsCounts struct {
	Clients     *ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsCountsClients     `json:"clients,omitempty"`     // Wireless LAN controller client counts
	Connections *ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsCountsConnections `json:"connections,omitempty"` // Wireless LAN controller associated access point counts
}
type ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsCountsClients struct {
	ByStatus *ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsCountsClientsByStatus `json:"byStatus,omitempty"` // Client counts by their status
}
type ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsCountsClientsByStatus struct {
	Online *int `json:"online,omitempty"` // Wireless LAN controller active client count
}
type ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsCountsConnections struct {
	ByStatus *ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsCountsConnectionsByStatus `json:"byStatus,omitempty"` // Access point counts by their status
	Total    *int                                                                                                       `json:"total,omitempty"`    // Wireless LAN controller associated total access point count
}
type ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsCountsConnectionsByStatus struct {
	Offline *int `json:"offline,omitempty"` // Wireless LAN controller associated offline access point count
	Online  *int `json:"online,omitempty"`  // Wireless LAN controller associated online access point count
}
type ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsFirmware struct {
	Version *ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsFirmwareVersion `json:"version,omitempty"` // Wireless LAN controller firmware version
}
type ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsFirmwareVersion struct {
	ShortName string `json:"shortName,omitempty"` // Wireless LAN controller firmware version short name
}
type ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsNetwork struct {
	ID string `json:"id,omitempty"` // Wireless LAN controller network ID
}
type ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsRedundancy struct {
	ChassisName     string                                                                                                `json:"chassisName,omitempty"`     // Wireless LAN controller chassis name
	ID              string                                                                                                `json:"id,omitempty"`              // Wireless LAN controller redundancy ID
	Management      *ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsRedundancyManagement `json:"management,omitempty"`      // Wireless LAN controller redundancy management interface information
	RedundantSerial string                                                                                                `json:"redundantSerial,omitempty"` // Wireless LAN controller redundant device serial
	Role            string                                                                                                `json:"role,omitempty"`            // Wireless LAN controller role(Active, Active recovery, Standby hot, Standby recovery and Offline)
}
type ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsRedundancyManagement struct {
	Addresses *[]ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsRedundancyManagementAddresses `json:"addresses,omitempty"` // Wireless LAN controller redundancy management interface addresses
}
type ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsRedundancyManagementAddresses struct {
	Address string `json:"address,omitempty"` // Wireless LAN controller redundancy management interface ip address
}
type ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceMeta struct {
	Counts *ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceMetaCounts `json:"counts,omitempty"` // Counts relating to the paginated dataset
}
type ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceMetaCounts struct {
	Items *ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceMetaCountsItems `json:"items,omitempty"` // Counts relating to the paginated items
}
type ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceMetaCountsItems struct {
	Remaining *int `json:"remaining,omitempty"` // The number of items in the dataset that are available on subsequent pages
	Total     *int `json:"total,omitempty"`     // The total number of items in the dataset
}

//GetOrganizationWirelessControllerAvailabilitiesChangeHistory List connectivity data of wireless LAN controllers in an organization
/* List connectivity data of wireless LAN controllers in an organization. If it is HA setup, then only returns active WLC data start from switchover

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessControllerAvailabilitiesChangeHistoryQueryParams Filtering parameter


*/

func (s *WirelessControllerService) GetOrganizationWirelessControllerAvailabilitiesChangeHistory(organizationID string, getOrganizationWirelessControllerAvailabilitiesChangeHistoryQueryParams *GetOrganizationWirelessControllerAvailabilitiesChangeHistoryQueryParams) (*ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistory, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wirelessController/availabilities/changeHistory"
	s.rateLimiterBucket.Wait(1)

	if getOrganizationWirelessControllerAvailabilitiesChangeHistoryQueryParams != nil && getOrganizationWirelessControllerAvailabilitiesChangeHistoryQueryParams.PerPage == -1 {
		var result *ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistory
		println("Paginate")
		getOrganizationWirelessControllerAvailabilitiesChangeHistoryQueryParams.PerPage = PAGINATION_PER_PAGE
		result2, response, err := Paginate(s.GetOrganizationWirelessControllerAvailabilitiesChangeHistoryPaginate, organizationID, "", getOrganizationWirelessControllerAvailabilitiesChangeHistoryQueryParams)
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
			var resultTmp *ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistory
			jsonResult2, _ := json.Marshal(paginatedResponse[i])
			err = json.Unmarshal(jsonResult2, &resultTmp)
			// Verificar si result es nil, si lo es inicialiarlo
			if result == nil {
				result = resultTmp
			} else {
				*result.Items = append(*result.Items, *resultTmp.Items...)
			}
		}
		return result, response, err
	}
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessControllerAvailabilitiesChangeHistoryQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistory{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWirelessControllerAvailabilitiesChangeHistory")
	}

	result := response.Result().(*ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistory)
	return result, response, err

}
func (s *WirelessControllerService) GetOrganizationWirelessControllerAvailabilitiesChangeHistoryPaginate(organizationID string, getOrganizationWirelessControllerAvailabilitiesChangeHistoryQueryParams any) (any, *resty.Response, error) {
	getOrganizationWirelessControllerAvailabilitiesChangeHistoryQueryParamsConverted := getOrganizationWirelessControllerAvailabilitiesChangeHistoryQueryParams.(*GetOrganizationWirelessControllerAvailabilitiesChangeHistoryQueryParams)

	return s.GetOrganizationWirelessControllerAvailabilitiesChangeHistory(organizationID, getOrganizationWirelessControllerAvailabilitiesChangeHistoryQueryParamsConverted)
}

//GetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByInterval List wireless client counts of wireless LAN controllers over time in an organization
/* List wireless client counts of wireless LAN controllers over time in an organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalQueryParams Filtering parameter


*/

func (s *WirelessControllerService) GetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByInterval(organizationID string, getOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalQueryParams *GetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalQueryParams) (*ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByInterval, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wirelessController/clients/overview/history/byDevice/byInterval"
	s.rateLimiterBucket.Wait(1)

	if getOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalQueryParams != nil && getOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalQueryParams.PerPage == -1 {
		var result *ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByInterval
		println("Paginate")
		getOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalQueryParams.PerPage = PAGINATION_PER_PAGE
		result2, response, err := Paginate(s.GetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalPaginate, organizationID, "", getOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalQueryParams)
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
			var resultTmp *ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByInterval
			jsonResult2, _ := json.Marshal(paginatedResponse[i])
			err = json.Unmarshal(jsonResult2, &resultTmp)
			// Verificar si result es nil, si lo es inicialiarlo
			if result == nil {
				result = resultTmp
			} else {
				*result.Items = append(*result.Items, *resultTmp.Items...)
			}
		}
		return result, response, err
	}
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByInterval{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByInterval")
	}

	result := response.Result().(*ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByInterval)
	return result, response, err

}
func (s *WirelessControllerService) GetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalPaginate(organizationID string, getOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalQueryParams any) (any, *resty.Response, error) {
	getOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalQueryParamsConverted := getOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalQueryParams.(*GetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalQueryParams)

	return s.GetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByInterval(organizationID, getOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalQueryParamsConverted)
}

//GetOrganizationWirelessControllerConnections List all access points associated with wireless LAN controllers in an organization
/* List all access points associated with wireless LAN controllers in an organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessControllerConnectionsQueryParams Filtering parameter


*/

func (s *WirelessControllerService) GetOrganizationWirelessControllerConnections(organizationID string, getOrganizationWirelessControllerConnectionsQueryParams *GetOrganizationWirelessControllerConnectionsQueryParams) (*ResponseWirelessControllerGetOrganizationWirelessControllerConnections, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wirelessController/connections"
	s.rateLimiterBucket.Wait(1)

	if getOrganizationWirelessControllerConnectionsQueryParams != nil && getOrganizationWirelessControllerConnectionsQueryParams.PerPage == -1 {
		var result *ResponseWirelessControllerGetOrganizationWirelessControllerConnections
		println("Paginate")
		getOrganizationWirelessControllerConnectionsQueryParams.PerPage = PAGINATION_PER_PAGE
		result2, response, err := Paginate(s.GetOrganizationWirelessControllerConnectionsPaginate, organizationID, "", getOrganizationWirelessControllerConnectionsQueryParams)
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
			var resultTmp *ResponseWirelessControllerGetOrganizationWirelessControllerConnections
			jsonResult2, _ := json.Marshal(paginatedResponse[i])
			err = json.Unmarshal(jsonResult2, &resultTmp)
			// Verificar si result es nil, si lo es inicialiarlo
			if result == nil {
				result = resultTmp
			} else {
				*result.Items = append(*result.Items, *resultTmp.Items...)
			}
		}
		return result, response, err
	}
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessControllerConnectionsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessControllerGetOrganizationWirelessControllerConnections{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWirelessControllerConnections")
	}

	result := response.Result().(*ResponseWirelessControllerGetOrganizationWirelessControllerConnections)
	return result, response, err

}
func (s *WirelessControllerService) GetOrganizationWirelessControllerConnectionsPaginate(organizationID string, getOrganizationWirelessControllerConnectionsQueryParams any) (any, *resty.Response, error) {
	getOrganizationWirelessControllerConnectionsQueryParamsConverted := getOrganizationWirelessControllerConnectionsQueryParams.(*GetOrganizationWirelessControllerConnectionsQueryParams)

	return s.GetOrganizationWirelessControllerConnections(organizationID, getOrganizationWirelessControllerConnectionsQueryParamsConverted)
}

//GetOrganizationWirelessControllerDevicesInterfacesL2ByDevice List wireless LAN controller layer 2 interfaces in an organization
/* List wireless LAN controller layer 2 interfaces in an organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessControllerDevicesInterfacesL2ByDeviceQueryParams Filtering parameter


*/

func (s *WirelessControllerService) GetOrganizationWirelessControllerDevicesInterfacesL2ByDevice(organizationID string, getOrganizationWirelessControllerDevicesInterfacesL2ByDeviceQueryParams *GetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceQueryParams) (*ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDevice, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wirelessController/devices/interfaces/l2/byDevice"
	s.rateLimiterBucket.Wait(1)

	if getOrganizationWirelessControllerDevicesInterfacesL2ByDeviceQueryParams != nil && getOrganizationWirelessControllerDevicesInterfacesL2ByDeviceQueryParams.PerPage == -1 {
		var result *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDevice
		println("Paginate")
		getOrganizationWirelessControllerDevicesInterfacesL2ByDeviceQueryParams.PerPage = PAGINATION_PER_PAGE
		result2, response, err := Paginate(s.GetOrganizationWirelessControllerDevicesInterfacesL2ByDevicePaginate, organizationID, "", getOrganizationWirelessControllerDevicesInterfacesL2ByDeviceQueryParams)
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
			var resultTmp *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDevice
			jsonResult2, _ := json.Marshal(paginatedResponse[i])
			err = json.Unmarshal(jsonResult2, &resultTmp)
			// Verificar si result es nil, si lo es inicialiarlo
			if result == nil {
				result = resultTmp
			} else {
				*result.Items = append(*result.Items, *resultTmp.Items...)
			}
		}
		return result, response, err
	}
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessControllerDevicesInterfacesL2ByDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWirelessControllerDevicesInterfacesL2ByDevice")
	}

	result := response.Result().(*ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDevice)
	return result, response, err

}
func (s *WirelessControllerService) GetOrganizationWirelessControllerDevicesInterfacesL2ByDevicePaginate(organizationID string, getOrganizationWirelessControllerDevicesInterfacesL2ByDeviceQueryParams any) (any, *resty.Response, error) {
	getOrganizationWirelessControllerDevicesInterfacesL2ByDeviceQueryParamsConverted := getOrganizationWirelessControllerDevicesInterfacesL2ByDeviceQueryParams.(*GetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceQueryParams)

	return s.GetOrganizationWirelessControllerDevicesInterfacesL2ByDevice(organizationID, getOrganizationWirelessControllerDevicesInterfacesL2ByDeviceQueryParamsConverted)
}

//GetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice List wireless LAN controller layer 2 interfaces history status in an organization
/* List wireless LAN controller layer 2 interfaces history status in an organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceQueryParams Filtering parameter


*/

func (s *WirelessControllerService) GetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice(organizationID string, getOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceQueryParams *GetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceQueryParams) (*ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wirelessController/devices/interfaces/l2/statuses/changeHistory/byDevice"
	s.rateLimiterBucket.Wait(1)

	if getOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceQueryParams != nil && getOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceQueryParams.PerPage == -1 {
		var result *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice
		println("Paginate")
		getOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceQueryParams.PerPage = PAGINATION_PER_PAGE
		result2, response, err := Paginate(s.GetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevicePaginate, organizationID, "", getOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceQueryParams)
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
			var resultTmp *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice
			jsonResult2, _ := json.Marshal(paginatedResponse[i])
			err = json.Unmarshal(jsonResult2, &resultTmp)
			// Verificar si result es nil, si lo es inicialiarlo
			if result == nil {
				result = resultTmp
			} else {
				*result.Items = append(*result.Items, *resultTmp.Items...)
			}
		}
		return result, response, err
	}
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice")
	}

	result := response.Result().(*ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice)
	return result, response, err

}
func (s *WirelessControllerService) GetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevicePaginate(organizationID string, getOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceQueryParams any) (any, *resty.Response, error) {
	getOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceQueryParamsConverted := getOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceQueryParams.(*GetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceQueryParams)

	return s.GetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice(organizationID, getOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceQueryParamsConverted)
}

//GetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval List wireless LAN controller layer 2 interfaces history usage in an organization
/* List wireless LAN controller layer 2 interfaces history usage in an organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalQueryParams Filtering parameter


*/

func (s *WirelessControllerService) GetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval(organizationID string, getOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalQueryParams *GetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalQueryParams) (*ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wirelessController/devices/interfaces/l2/usage/history/byInterval"
	s.rateLimiterBucket.Wait(1)

	if getOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalQueryParams != nil && getOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalQueryParams.PerPage == -1 {
		var result *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval
		println("Paginate")
		getOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalQueryParams.PerPage = PAGINATION_PER_PAGE
		result2, response, err := Paginate(s.GetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalPaginate, organizationID, "", getOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalQueryParams)
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
			var resultTmp *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval
			jsonResult2, _ := json.Marshal(paginatedResponse[i])
			err = json.Unmarshal(jsonResult2, &resultTmp)
			// Verificar si result es nil, si lo es inicialiarlo
			if result == nil {
				result = resultTmp
			} else {
				*result.Items = append(*result.Items, *resultTmp.Items...)
			}
		}
		return result, response, err
	}
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval")
	}

	result := response.Result().(*ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval)
	return result, response, err

}
func (s *WirelessControllerService) GetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalPaginate(organizationID string, getOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalQueryParams any) (any, *resty.Response, error) {
	getOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalQueryParamsConverted := getOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalQueryParams.(*GetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalQueryParams)

	return s.GetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval(organizationID, getOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalQueryParamsConverted)
}

//GetOrganizationWirelessControllerDevicesInterfacesL3ByDevice List wireless LAN controller layer 3 interfaces in an organization
/* List wireless LAN controller layer 3 interfaces in an organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessControllerDevicesInterfacesL3ByDeviceQueryParams Filtering parameter


*/

func (s *WirelessControllerService) GetOrganizationWirelessControllerDevicesInterfacesL3ByDevice(organizationID string, getOrganizationWirelessControllerDevicesInterfacesL3ByDeviceQueryParams *GetOrganizationWirelessControllerDevicesInterfacesL3ByDeviceQueryParams) (*ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3ByDevice, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wirelessController/devices/interfaces/l3/byDevice"
	s.rateLimiterBucket.Wait(1)

	if getOrganizationWirelessControllerDevicesInterfacesL3ByDeviceQueryParams != nil && getOrganizationWirelessControllerDevicesInterfacesL3ByDeviceQueryParams.PerPage == -1 {
		var result *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3ByDevice
		println("Paginate")
		getOrganizationWirelessControllerDevicesInterfacesL3ByDeviceQueryParams.PerPage = PAGINATION_PER_PAGE
		result2, response, err := Paginate(s.GetOrganizationWirelessControllerDevicesInterfacesL3ByDevicePaginate, organizationID, "", getOrganizationWirelessControllerDevicesInterfacesL3ByDeviceQueryParams)
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
			var resultTmp *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3ByDevice
			jsonResult2, _ := json.Marshal(paginatedResponse[i])
			err = json.Unmarshal(jsonResult2, &resultTmp)
			// Verificar si result es nil, si lo es inicialiarlo
			if result == nil {
				result = resultTmp
			} else {
				*result.Items = append(*result.Items, *resultTmp.Items...)
			}
		}
		return result, response, err
	}
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessControllerDevicesInterfacesL3ByDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3ByDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWirelessControllerDevicesInterfacesL3ByDevice")
	}

	result := response.Result().(*ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3ByDevice)
	return result, response, err

}
func (s *WirelessControllerService) GetOrganizationWirelessControllerDevicesInterfacesL3ByDevicePaginate(organizationID string, getOrganizationWirelessControllerDevicesInterfacesL3ByDeviceQueryParams any) (any, *resty.Response, error) {
	getOrganizationWirelessControllerDevicesInterfacesL3ByDeviceQueryParamsConverted := getOrganizationWirelessControllerDevicesInterfacesL3ByDeviceQueryParams.(*GetOrganizationWirelessControllerDevicesInterfacesL3ByDeviceQueryParams)

	return s.GetOrganizationWirelessControllerDevicesInterfacesL3ByDevice(organizationID, getOrganizationWirelessControllerDevicesInterfacesL3ByDeviceQueryParamsConverted)
}

//GetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice List wireless LAN controller layer 3 interfaces history status in an organization
/* List wireless LAN controller layer 3 interfaces history status in an organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceQueryParams Filtering parameter


*/

func (s *WirelessControllerService) GetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice(organizationID string, getOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceQueryParams *GetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceQueryParams) (*ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wirelessController/devices/interfaces/l3/statuses/changeHistory/byDevice"
	s.rateLimiterBucket.Wait(1)

	if getOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceQueryParams != nil && getOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceQueryParams.PerPage == -1 {
		var result *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice
		println("Paginate")
		getOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceQueryParams.PerPage = PAGINATION_PER_PAGE
		result2, response, err := Paginate(s.GetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevicePaginate, organizationID, "", getOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceQueryParams)
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
			var resultTmp *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice
			jsonResult2, _ := json.Marshal(paginatedResponse[i])
			err = json.Unmarshal(jsonResult2, &resultTmp)
			// Verificar si result es nil, si lo es inicialiarlo
			if result == nil {
				result = resultTmp
			} else {
				*result.Items = append(*result.Items, *resultTmp.Items...)
			}
		}
		return result, response, err
	}
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice")
	}

	result := response.Result().(*ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice)
	return result, response, err

}
func (s *WirelessControllerService) GetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevicePaginate(organizationID string, getOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceQueryParams any) (any, *resty.Response, error) {
	getOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceQueryParamsConverted := getOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceQueryParams.(*GetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceQueryParams)

	return s.GetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice(organizationID, getOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceQueryParamsConverted)
}

//GetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByInterval List wireless LAN controller layer 3 interfaces history usage in an organization
/* List wireless LAN controller layer 3 interfaces history usage in an organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalQueryParams Filtering parameter


*/

func (s *WirelessControllerService) GetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByInterval(organizationID string, getOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalQueryParams *GetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalQueryParams) (*ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByInterval, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wirelessController/devices/interfaces/l3/usage/history/byInterval"
	s.rateLimiterBucket.Wait(1)

	if getOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalQueryParams != nil && getOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalQueryParams.PerPage == -1 {
		var result *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByInterval
		println("Paginate")
		getOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalQueryParams.PerPage = PAGINATION_PER_PAGE
		result2, response, err := Paginate(s.GetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalPaginate, organizationID, "", getOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalQueryParams)
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
			var resultTmp *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByInterval
			jsonResult2, _ := json.Marshal(paginatedResponse[i])
			err = json.Unmarshal(jsonResult2, &resultTmp)
			// Verificar si result es nil, si lo es inicialiarlo
			if result == nil {
				result = resultTmp
			} else {
				*result.Items = append(*result.Items, *resultTmp.Items...)
			}
		}
		return result, response, err
	}
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByInterval{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByInterval")
	}

	result := response.Result().(*ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByInterval)
	return result, response, err

}
func (s *WirelessControllerService) GetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalPaginate(organizationID string, getOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalQueryParams any) (any, *resty.Response, error) {
	getOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalQueryParamsConverted := getOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalQueryParams.(*GetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalQueryParams)

	return s.GetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByInterval(organizationID, getOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalQueryParamsConverted)
}

//GetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice Retrieve the packet counters for the interfaces of a Wireless LAN controller
/* Retrieve the packet counters for the interfaces of a Wireless LAN controller

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceQueryParams Filtering parameter


*/

func (s *WirelessControllerService) GetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice(organizationID string, getOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceQueryParams *GetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceQueryParams) (*ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wirelessController/devices/interfaces/packets/overview/byDevice"
	s.rateLimiterBucket.Wait(1)

	if getOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceQueryParams != nil && getOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceQueryParams.PerPage == -1 {
		var result *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice
		println("Paginate")
		getOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceQueryParams.PerPage = PAGINATION_PER_PAGE
		result2, response, err := Paginate(s.GetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevicePaginate, organizationID, "", getOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceQueryParams)
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
			var resultTmp *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice
			jsonResult2, _ := json.Marshal(paginatedResponse[i])
			err = json.Unmarshal(jsonResult2, &resultTmp)
			// Verificar si result es nil, si lo es inicialiarlo
			if result == nil {
				result = resultTmp
			} else {
				*result.Items = append(*result.Items, *resultTmp.Items...)
			}
		}
		return result, response, err
	}
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice")
	}

	result := response.Result().(*ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice)
	return result, response, err

}
func (s *WirelessControllerService) GetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevicePaginate(organizationID string, getOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceQueryParams any) (any, *resty.Response, error) {
	getOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceQueryParamsConverted := getOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceQueryParams.(*GetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceQueryParams)

	return s.GetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice(organizationID, getOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceQueryParamsConverted)
}

//GetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval Retrieve the traffic for the interfaces of a Wireless LAN controller
/* Retrieve the traffic for the interfaces of a Wireless LAN controller

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalQueryParams Filtering parameter


*/

func (s *WirelessControllerService) GetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval(organizationID string, getOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalQueryParams *GetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalQueryParams) (*ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wirelessController/devices/interfaces/usage/history/byInterval"
	s.rateLimiterBucket.Wait(1)

	if getOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalQueryParams != nil && getOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalQueryParams.PerPage == -1 {
		var result *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval
		println("Paginate")
		getOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalQueryParams.PerPage = PAGINATION_PER_PAGE
		result2, response, err := Paginate(s.GetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalPaginate, organizationID, "", getOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalQueryParams)
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
			var resultTmp *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval
			jsonResult2, _ := json.Marshal(paginatedResponse[i])
			err = json.Unmarshal(jsonResult2, &resultTmp)
			// Verificar si result es nil, si lo es inicialiarlo
			if result == nil {
				result = resultTmp
			} else {
				*result.Items = append(*result.Items, *resultTmp.Items...)
			}
		}
		return result, response, err
	}
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval")
	}

	result := response.Result().(*ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval)
	return result, response, err

}
func (s *WirelessControllerService) GetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalPaginate(organizationID string, getOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalQueryParams any) (any, *resty.Response, error) {
	getOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalQueryParamsConverted := getOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalQueryParams.(*GetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalQueryParams)

	return s.GetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval(organizationID, getOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalQueryParamsConverted)
}

//GetOrganizationWirelessControllerDevicesRedundancyFailoverHistory List the failover events of wireless LAN controllers in an organization
/* List the failover events of wireless LAN controllers in an organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessControllerDevicesRedundancyFailoverHistoryQueryParams Filtering parameter


*/

func (s *WirelessControllerService) GetOrganizationWirelessControllerDevicesRedundancyFailoverHistory(organizationID string, getOrganizationWirelessControllerDevicesRedundancyFailoverHistoryQueryParams *GetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryQueryParams) (*ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistory, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wirelessController/devices/redundancy/failover/history"
	s.rateLimiterBucket.Wait(1)

	if getOrganizationWirelessControllerDevicesRedundancyFailoverHistoryQueryParams != nil && getOrganizationWirelessControllerDevicesRedundancyFailoverHistoryQueryParams.PerPage == -1 {
		var result *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistory
		println("Paginate")
		getOrganizationWirelessControllerDevicesRedundancyFailoverHistoryQueryParams.PerPage = PAGINATION_PER_PAGE
		result2, response, err := Paginate(s.GetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryPaginate, organizationID, "", getOrganizationWirelessControllerDevicesRedundancyFailoverHistoryQueryParams)
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
			var resultTmp *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistory
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

	queryString, _ := query.Values(getOrganizationWirelessControllerDevicesRedundancyFailoverHistoryQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistory{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWirelessControllerDevicesRedundancyFailoverHistory")
	}

	result := response.Result().(*ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistory)
	return result, response, err

}
func (s *WirelessControllerService) GetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryPaginate(organizationID string, getOrganizationWirelessControllerDevicesRedundancyFailoverHistoryQueryParams any) (any, *resty.Response, error) {
	getOrganizationWirelessControllerDevicesRedundancyFailoverHistoryQueryParamsConverted := getOrganizationWirelessControllerDevicesRedundancyFailoverHistoryQueryParams.(*GetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryQueryParams)

	return s.GetOrganizationWirelessControllerDevicesRedundancyFailoverHistory(organizationID, getOrganizationWirelessControllerDevicesRedundancyFailoverHistoryQueryParamsConverted)
}

//GetOrganizationWirelessControllerDevicesRedundancyStatuses List redundancy details of wireless LAN controllers in an organization
/* List redundancy details of wireless LAN controllers in an organization. The failover count refers to the total failovers system happens from the moment of this device onboarding to Dashboard

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessControllerDevicesRedundancyStatusesQueryParams Filtering parameter


*/

func (s *WirelessControllerService) GetOrganizationWirelessControllerDevicesRedundancyStatuses(organizationID string, getOrganizationWirelessControllerDevicesRedundancyStatusesQueryParams *GetOrganizationWirelessControllerDevicesRedundancyStatusesQueryParams) (*ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatuses, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wirelessController/devices/redundancy/statuses"
	s.rateLimiterBucket.Wait(1)

	if getOrganizationWirelessControllerDevicesRedundancyStatusesQueryParams != nil && getOrganizationWirelessControllerDevicesRedundancyStatusesQueryParams.PerPage == -1 {
		var result *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatuses
		println("Paginate")
		getOrganizationWirelessControllerDevicesRedundancyStatusesQueryParams.PerPage = PAGINATION_PER_PAGE
		result2, response, err := Paginate(s.GetOrganizationWirelessControllerDevicesRedundancyStatusesPaginate, organizationID, "", getOrganizationWirelessControllerDevicesRedundancyStatusesQueryParams)
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
			var resultTmp *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatuses
			jsonResult2, _ := json.Marshal(paginatedResponse[i])
			err = json.Unmarshal(jsonResult2, &resultTmp)
			// Verificar si result es nil, si lo es inicialiarlo
			if result == nil {
				result = resultTmp
			} else {
				*result.Items = append(*result.Items, *resultTmp.Items...)
			}
		}
		return result, response, err
	}
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessControllerDevicesRedundancyStatusesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatuses{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWirelessControllerDevicesRedundancyStatuses")
	}

	result := response.Result().(*ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatuses)
	return result, response, err

}
func (s *WirelessControllerService) GetOrganizationWirelessControllerDevicesRedundancyStatusesPaginate(organizationID string, getOrganizationWirelessControllerDevicesRedundancyStatusesQueryParams any) (any, *resty.Response, error) {
	getOrganizationWirelessControllerDevicesRedundancyStatusesQueryParamsConverted := getOrganizationWirelessControllerDevicesRedundancyStatusesQueryParams.(*GetOrganizationWirelessControllerDevicesRedundancyStatusesQueryParams)

	return s.GetOrganizationWirelessControllerDevicesRedundancyStatuses(organizationID, getOrganizationWirelessControllerDevicesRedundancyStatusesQueryParamsConverted)
}

//GetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByInterval List cpu utilization data of wireless LAN controllers in an organization
/* List cpu utilization data of wireless LAN controllers in an organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalQueryParams Filtering parameter


*/

func (s *WirelessControllerService) GetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByInterval(organizationID string, getOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalQueryParams *GetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalQueryParams) (*ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByInterval, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wirelessController/devices/system/utilization/history/byInterval"
	s.rateLimiterBucket.Wait(1)

	if getOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalQueryParams != nil && getOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalQueryParams.PerPage == -1 {
		var result *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByInterval
		println("Paginate")
		getOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalQueryParams.PerPage = PAGINATION_PER_PAGE
		result2, response, err := Paginate(s.GetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalPaginate, organizationID, "", getOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalQueryParams)
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
			var resultTmp *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByInterval
			jsonResult2, _ := json.Marshal(paginatedResponse[i])
			err = json.Unmarshal(jsonResult2, &resultTmp)
			// Verificar si result es nil, si lo es inicialiarlo
			if result == nil {
				result = resultTmp
			} else {
				*result.Items = append(*result.Items, *resultTmp.Items...)
			}
		}
		return result, response, err
	}
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByInterval{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByInterval")
	}

	result := response.Result().(*ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByInterval)
	return result, response, err

}
func (s *WirelessControllerService) GetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalPaginate(organizationID string, getOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalQueryParams any) (any, *resty.Response, error) {
	getOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalQueryParamsConverted := getOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalQueryParams.(*GetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalQueryParams)

	return s.GetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByInterval(organizationID, getOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalQueryParamsConverted)
}

//GetOrganizationWirelessControllerOverviewByDevice List the overview information of wireless LAN controllers in an organization and it is updated every minute.
/* List the overview information of wireless LAN controllers in an organization and it is updated every minute.

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessControllerOverviewByDeviceQueryParams Filtering parameter


*/

func (s *WirelessControllerService) GetOrganizationWirelessControllerOverviewByDevice(organizationID string, getOrganizationWirelessControllerOverviewByDeviceQueryParams *GetOrganizationWirelessControllerOverviewByDeviceQueryParams) (*ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDevice, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wirelessController/overview/byDevice"
	s.rateLimiterBucket.Wait(1)

	if getOrganizationWirelessControllerOverviewByDeviceQueryParams != nil && getOrganizationWirelessControllerOverviewByDeviceQueryParams.PerPage == -1 {
		var result *ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDevice
		println("Paginate")
		getOrganizationWirelessControllerOverviewByDeviceQueryParams.PerPage = PAGINATION_PER_PAGE
		result2, response, err := Paginate(s.GetOrganizationWirelessControllerOverviewByDevicePaginate, organizationID, "", getOrganizationWirelessControllerOverviewByDeviceQueryParams)
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
			var resultTmp *ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDevice
			jsonResult2, _ := json.Marshal(paginatedResponse[i])
			err = json.Unmarshal(jsonResult2, &resultTmp)
			// Verificar si result es nil, si lo es inicialiarlo
			if result == nil {
				result = resultTmp
			} else {
				*result.Items = append(*result.Items, *resultTmp.Items...)
			}
		}
		return result, response, err
	}
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessControllerOverviewByDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWirelessControllerOverviewByDevice")
	}

	result := response.Result().(*ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDevice)
	return result, response, err

}
func (s *WirelessControllerService) GetOrganizationWirelessControllerOverviewByDevicePaginate(organizationID string, getOrganizationWirelessControllerOverviewByDeviceQueryParams any) (any, *resty.Response, error) {
	getOrganizationWirelessControllerOverviewByDeviceQueryParamsConverted := getOrganizationWirelessControllerOverviewByDeviceQueryParams.(*GetOrganizationWirelessControllerOverviewByDeviceQueryParams)

	return s.GetOrganizationWirelessControllerOverviewByDevice(organizationID, getOrganizationWirelessControllerOverviewByDeviceQueryParamsConverted)
}
