package meraki

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type OrganizationsService service

type GetOrganizationActionBatchesQueryParams struct {
	Status string `url:"status,omitempty"` //Filter batches by status. Valid types are pending, completed, and failed.
}
type GetOrganizationAPIRequestsQueryParams struct {
	T0            string   `url:"t0,omitempty"`             //The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	T1            string   `url:"t1,omitempty"`             //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan      float64  `url:"timespan,omitempty"`       //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 31 days.
	PerPage       int      `url:"perPage,omitempty"`        //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50.
	StartingAfter string   `url:"startingAfter,omitempty"`  //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string   `url:"endingBefore,omitempty"`   //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	AdminID       string   `url:"adminId,omitempty"`        //Filter the results by the ID of the admin who made the API requests
	Path          string   `url:"path,omitempty"`           //Filter the results by the path of the API requests
	Method        string   `url:"method,omitempty"`         //Filter the results by the method of the API requests (must be 'GET', 'PUT', 'POST' or 'DELETE')
	ResponseCode  int      `url:"responseCode,omitempty"`   //Filter the results by the response code of the API requests
	SourceIP      string   `url:"sourceIp,omitempty"`       //Filter the results by the IP address of the originating API request
	UserAgent     string   `url:"userAgent,omitempty"`      //Filter the results by the user agent string of the API request
	Version       int      `url:"version,omitempty"`        //Filter the results by the API version of the API request
	OperationIDs  []string `url:"operationIds[],omitempty"` //Filter the results by one or more operation IDs for the API request
}
type GetOrganizationAPIRequestsOverviewQueryParams struct {
	T0       string  `url:"t0,omitempty"`       //The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	T1       string  `url:"t1,omitempty"`       //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan float64 `url:"timespan,omitempty"` //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 31 days.
}
type GetOrganizationAPIRequestsOverviewResponseCodesByIntervalQueryParams struct {
	T0           string   `url:"t0,omitempty"`             //The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	T1           string   `url:"t1,omitempty"`             //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan     float64  `url:"timespan,omitempty"`       //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 31 days. If interval is provided, the timespan will be autocalculated.
	Interval     int      `url:"interval,omitempty"`       //The time interval in seconds for returned data. The valid intervals are: 120, 3600, 14400, 21600. The default is 21600. Interval is calculated if time params are provided.
	Version      int      `url:"version,omitempty"`        //Filter by API version of the endpoint. Allowable values are: [0, 1]
	OperationIDs []string `url:"operationIds[],omitempty"` //Filter by operation ID of the endpoint
	SourceIPs    []string `url:"sourceIps[],omitempty"`    //Filter by source IP that made the API request
	AdminIDs     []string `url:"adminIds[],omitempty"`     //Filter by admin ID of user that made the API request
	UserAgent    string   `url:"userAgent,omitempty"`      //Filter by user agent string for API request. This will filter by a complete or partial match.
}
type GetOrganizationClientsBandwidthUsageHistoryQueryParams struct {
	T0       string  `url:"t0,omitempty"`       //The beginning of the timespan for the data.
	T1       string  `url:"t1,omitempty"`       //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan float64 `url:"timespan,omitempty"` //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
}
type GetOrganizationClientsOverviewQueryParams struct {
	T0       string  `url:"t0,omitempty"`       //The beginning of the timespan for the data.
	T1       string  `url:"t1,omitempty"`       //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan float64 `url:"timespan,omitempty"` //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
}
type GetOrganizationClientsSearchQueryParams struct {
	Mac           string `url:"mac,omitempty"`           //The MAC address of the client. Required.
	PerPage       int    `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 5. Default is 5.
	StartingAfter string `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
}
type GetOrganizationConfigurationChangesQueryParams struct {
	T0            string  `url:"t0,omitempty"`            //The beginning of the timespan for the data. The maximum lookback period is 365 days from today.
	T1            string  `url:"t1,omitempty"`            //The end of the timespan for the data. t1 can be a maximum of 365 days after t0.
	Timespan      float64 `url:"timespan,omitempty"`      //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 365 days. The default is 365 days.
	PerPage       int     `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 5000. Default is 5000.
	StartingAfter string  `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string  `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	NetworkID     string  `url:"networkId,omitempty"`     //Filters on the given network
	AdminID       string  `url:"adminId,omitempty"`       //Filters on the given Admin
}
type GetOrganizationDevicesQueryParams struct {
	PerPage                   int      `url:"perPage,omitempty"`                   //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter             string   `url:"startingAfter,omitempty"`             //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore              string   `url:"endingBefore,omitempty"`              //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	ConfigurationUpdatedAfter string   `url:"configurationUpdatedAfter,omitempty"` //Filter results by whether or not the device's configuration has been updated after the given timestamp
	NetworkIDs                []string `url:"networkIds[],omitempty"`              //Optional parameter to filter devices by network.
	ProductTypes              []string `url:"productTypes[],omitempty"`            //Optional parameter to filter devices by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, and sensor.
	Tags                      []string `url:"tags[],omitempty"`                    //Optional parameter to filter devices by tags.
	TagsFilterType            string   `url:"tagsFilterType,omitempty"`            //Optional parameter of value 'withAnyTags' or 'withAllTags' to indicate whether to return networks which contain ANY or ALL of the included tags. If no type is included, 'withAnyTags' will be selected.
	Name                      string   `url:"name,omitempty"`                      //Optional parameter to filter devices by name. All returned devices will have a name that contains the search term or is an exact match.
	Mac                       string   `url:"mac,omitempty"`                       //Optional parameter to filter devices by MAC address. All returned devices will have a MAC address that contains the search term or is an exact match.
	Serial                    string   `url:"serial,omitempty"`                    //Optional parameter to filter devices by serial number. All returned devices will have a serial number that contains the search term or is an exact match.
	Model                     string   `url:"model,omitempty"`                     //Optional parameter to filter devices by model. All returned devices will have a model that contains the search term or is an exact match.
	Macs                      []string `url:"macs[],omitempty"`                    //Optional parameter to filter devices by one or more MAC addresses. All returned devices will have a MAC address that is an exact match.
	Serials                   []string `url:"serials[],omitempty"`                 //Optional parameter to filter devices by one or more serial numbers. All returned devices will have a serial number that is an exact match.
	SensorMetrics             []string `url:"sensorMetrics[],omitempty"`           //Optional parameter to filter devices by the metrics that they provide. Only applies to sensor devices.
	SensorAlertProfileIDs     []string `url:"sensorAlertProfileIds[],omitempty"`   //Optional parameter to filter devices by the alert profiles that are bound to them. Only applies to sensor devices.
	Models                    []string `url:"models[],omitempty"`                  //Optional parameter to filter devices by one or more models. All returned devices will have a model that is an exact match.
}
type GetOrganizationDevicesAvailabilitiesQueryParams struct {
	PerPage        int      `url:"perPage,omitempty"`        //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter  string   `url:"startingAfter,omitempty"`  //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore   string   `url:"endingBefore,omitempty"`   //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	NetworkIDs     []string `url:"networkIds[],omitempty"`   //Optional parameter to filter device availabilities by network ID. This filter uses multiple exact matches.
	ProductTypes   []string `url:"productTypes[],omitempty"` //Optional parameter to filter device availabilities by device product types. This filter uses multiple exact matches.
	Serials        []string `url:"serials[],omitempty"`      //Optional parameter to filter device availabilities by device serial numbers. This filter uses multiple exact matches.
	Tags           []string `url:"tags[],omitempty"`         //An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, 'tagsFilterType' should also be included (see below). This filter uses multiple exact matches.
	TagsFilterType string   `url:"tagsFilterType,omitempty"` //An optional parameter of value 'withAnyTags' or 'withAllTags' to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, 'withAnyTags' will be selected.
}
type GetOrganizationDevicesPowerModulesStatusesByDeviceQueryParams struct {
	PerPage        int      `url:"perPage,omitempty"`        //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter  string   `url:"startingAfter,omitempty"`  //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore   string   `url:"endingBefore,omitempty"`   //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	NetworkIDs     []string `url:"networkIds[],omitempty"`   //Optional parameter to filter device availabilities by network ID. This filter uses multiple exact matches.
	ProductTypes   []string `url:"productTypes[],omitempty"` //Optional parameter to filter device availabilities by device product types. This filter uses multiple exact matches.
	Serials        []string `url:"serials[],omitempty"`      //Optional parameter to filter device availabilities by device serial numbers. This filter uses multiple exact matches.
	Tags           []string `url:"tags[],omitempty"`         //An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, 'tagsFilterType' should also be included (see below). This filter uses multiple exact matches.
	TagsFilterType string   `url:"tagsFilterType,omitempty"` //An optional parameter of value 'withAnyTags' or 'withAllTags' to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, 'withAnyTags' will be selected.
}
type GetOrganizationDevicesProvisioningStatusesQueryParams struct {
	PerPage        int      `url:"perPage,omitempty"`        //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter  string   `url:"startingAfter,omitempty"`  //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore   string   `url:"endingBefore,omitempty"`   //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	NetworkIDs     []string `url:"networkIds[],omitempty"`   //Optional parameter to filter device by network ID. This filter uses multiple exact matches.
	ProductTypes   []string `url:"productTypes[],omitempty"` //Optional parameter to filter device by device product types. This filter uses multiple exact matches.
	Serials        []string `url:"serials[],omitempty"`      //Optional parameter to filter device by device serial numbers. This filter uses multiple exact matches.
	Status         string   `url:"status,omitempty"`         //An optional parameter to filter devices by the provisioning status. Accepted statuses: unprovisioned, incomplete, complete.
	Tags           []string `url:"tags[],omitempty"`         //An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, 'tagsFilterType' should also be included (see below). This filter uses multiple exact matches.
	TagsFilterType string   `url:"tagsFilterType,omitempty"` //An optional parameter of value 'withAnyTags' or 'withAllTags' to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, 'withAnyTags' will be selected.
}
type GetOrganizationDevicesStatusesQueryParams struct {
	PerPage        int      `url:"perPage,omitempty"`        //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter  string   `url:"startingAfter,omitempty"`  //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore   string   `url:"endingBefore,omitempty"`   //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	NetworkIDs     []string `url:"networkIds[],omitempty"`   //Optional parameter to filter devices by network ids.
	Serials        []string `url:"serials[],omitempty"`      //Optional parameter to filter devices by serials.
	Statuses       []string `url:"statuses[],omitempty"`     //Optional parameter to filter devices by statuses. Valid statuses are ["online", "alerting", "offline", "dormant"].
	ProductTypes   []string `url:"productTypes[],omitempty"` //An optional parameter to filter device statuses by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, and sensor.
	Models         []string `url:"models[],omitempty"`       //Optional parameter to filter devices by models.
	Tags           []string `url:"tags[],omitempty"`         //An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, 'tagsFilterType' should also be included (see below).
	TagsFilterType string   `url:"tagsFilterType,omitempty"` //An optional parameter of value 'withAnyTags' or 'withAllTags' to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, 'withAnyTags' will be selected.
}
type GetOrganizationDevicesStatusesOverviewQueryParams struct {
	ProductTypes []string `url:"productTypes[],omitempty"` //An optional parameter to filter device statuses by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, and sensor.
	NetworkIDs   []string `url:"networkIds[],omitempty"`   //An optional parameter to filter device statuses by network.
}
type GetOrganizationDevicesUplinksAddressesByDeviceQueryParams struct {
	PerPage        int      `url:"perPage,omitempty"`        //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter  string   `url:"startingAfter,omitempty"`  //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore   string   `url:"endingBefore,omitempty"`   //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	NetworkIDs     []string `url:"networkIds[],omitempty"`   //Optional parameter to filter device uplinks by network ID. This filter uses multiple exact matches.
	ProductTypes   []string `url:"productTypes[],omitempty"` //Optional parameter to filter device uplinks by device product types. This filter uses multiple exact matches.
	Serials        []string `url:"serials[],omitempty"`      //Optional parameter to filter device availabilities by device serial numbers. This filter uses multiple exact matches.
	Tags           []string `url:"tags[],omitempty"`         //An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, 'tagsFilterType' should also be included (see below). This filter uses multiple exact matches.
	TagsFilterType string   `url:"tagsFilterType,omitempty"` //An optional parameter of value 'withAnyTags' or 'withAllTags' to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, 'withAnyTags' will be selected.
}
type GetOrganizationDevicesUplinksLossAndLatencyQueryParams struct {
	T0       string  `url:"t0,omitempty"`       //The beginning of the timespan for the data. The maximum lookback period is 60 days from today.
	T1       string  `url:"t1,omitempty"`       //The end of the timespan for the data. t1 can be a maximum of 5 minutes after t0. The latest possible time that t1 can be is 2 minutes into the past.
	Timespan float64 `url:"timespan,omitempty"` //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 5 minutes. The default is 5 minutes.
	Uplink   string  `url:"uplink,omitempty"`   //Optional filter for a specific WAN uplink. Valid uplinks are wan1, wan2, cellular. Default will return all uplinks.
	IP       string  `url:"ip,omitempty"`       //Optional filter for a specific destination IP. Default will return all destination IPs.
}
type GetOrganizationFirmwareUpgradesQueryParams struct {
	Status      []string `url:"status[],omitempty"`      //The status of an upgrade
	ProductType []string `url:"productType[],omitempty"` //The product type in a given upgrade ID
}
type GetOrganizationFirmwareUpgradesByDeviceQueryParams struct {
	PerPage                 int      `url:"perPage,omitempty"`                   //The number of entries per page returned. Acceptable range is 3 - 50. Default is 50.
	StartingAfter           string   `url:"startingAfter,omitempty"`             //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore            string   `url:"endingBefore,omitempty"`              //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	NetworkIDs              []string `url:"networkIds[],omitempty"`              //Optional parameter to filter by network
	Serials                 []string `url:"serials[],omitempty"`                 //Optional parameter to filter by serial number.  All returned devices will have a serial number that is an exact match.
	Macs                    []string `url:"macs[],omitempty"`                    //Optional parameter to filter by one or more MAC addresses belonging to devices. All devices returned belong to MAC addresses that are an exact match.
	FirmwareUpgradeIDs      []string `url:"firmwareUpgradeIds[],omitempty"`      //Optional parameter to filter by firmware upgrade ids.
	FirmwareUpgradeBatchIDs []string `url:"firmwareUpgradeBatchIds[],omitempty"` //Optional parameter to filter by firmware upgrade batch ids.
}
type GetOrganizationInventoryDevicesQueryParams struct {
	PerPage        int      `url:"perPage,omitempty"`        //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter  string   `url:"startingAfter,omitempty"`  //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore   string   `url:"endingBefore,omitempty"`   //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	UsedState      string   `url:"usedState,omitempty"`      //Filter results by used or unused inventory. Accepted values are 'used' or 'unused'.
	Search         string   `url:"search,omitempty"`         //Search for devices in inventory based on serial number, mac address, or model.
	Macs           []string `url:"macs[],omitempty"`         //Search for devices in inventory based on mac addresses.
	NetworkIDs     []string `url:"networkIds[],omitempty"`   //Search for devices in inventory based on network ids.
	Serials        []string `url:"serials[],omitempty"`      //Search for devices in inventory based on serials.
	Models         []string `url:"models[],omitempty"`       //Search for devices in inventory based on model.
	OrderNumbers   []string `url:"orderNumbers[],omitempty"` //Search for devices in inventory based on order numbers.
	Tags           []string `url:"tags[],omitempty"`         //Filter devices by tags. The filtering is case-sensitive. If tags are included, 'tagsFilterType' should also be included (see below).
	TagsFilterType string   `url:"tagsFilterType,omitempty"` //To use with 'tags' parameter, to filter devices which contain ANY or ALL given tags. Accepted values are 'withAnyTags' or 'withAllTags', default is 'withAnyTags'.
	ProductTypes   []string `url:"productTypes[],omitempty"` //Filter devices by product type. Accepted values are appliance, camera, cellularGateway, sensor, switch, systemsManager, and wireless.
}
type GetOrganizationInventoryOnboardingCloudMonitoringImportsQueryParams struct {
	ImportIDs []string `url:"importIds[],omitempty"` //import ids from an imports
}
type GetOrganizationInventoryOnboardingCloudMonitoringNetworksQueryParams struct {
	DeviceType    string `url:"deviceType,omitempty"`    //Device Type switch or wireless controller
	PerPage       int    `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 100000. Default is 1000.
	StartingAfter string `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
}
type GetOrganizationLicensesQueryParams struct {
	PerPage       int    `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter string `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	DeviceSerial  string `url:"deviceSerial,omitempty"`  //Filter the licenses to those assigned to a particular device. Returned in the same order that they are queued to the device.
	NetworkID     string `url:"networkId,omitempty"`     //Filter the licenses to those assigned in a particular network
	State         string `url:"state,omitempty"`         //Filter the licenses to those in a particular state. Can be one of 'active', 'expired', 'expiring', 'recentlyQueued', 'unused' or 'unusedActive'
}
type GetOrganizationNetworksQueryParams struct {
	ConfigTemplateID        string   `url:"configTemplateId,omitempty"`        //An optional parameter that is the ID of a config template. Will return all networks bound to that template.
	IsBoundToConfigTemplate bool     `url:"isBoundToConfigTemplate,omitempty"` //An optional parameter to filter config template bound networks. If configTemplateId is set, this cannot be false.
	Tags                    []string `url:"tags[],omitempty"`                  //An optional parameter to filter networks by tags. The filtering is case-sensitive. If tags are included, 'tagsFilterType' should also be included (see below).
	TagsFilterType          string   `url:"tagsFilterType,omitempty"`          //An optional parameter of value 'withAnyTags' or 'withAllTags' to indicate whether to return networks which contain ANY or ALL of the included tags. If no type is included, 'withAnyTags' will be selected.
	PerPage                 int      `url:"perPage,omitempty"`                 //The number of entries per page returned. Acceptable range is 3 - 100000. Default is 1000.
	StartingAfter           string   `url:"startingAfter,omitempty"`           //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore            string   `url:"endingBefore,omitempty"`            //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
}
type GetOrganizationPolicyObjectsQueryParams struct {
	PerPage       int    `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 10 - 5000. Default is 5000.
	StartingAfter string `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
}
type GetOrganizationPolicyObjectsGroupsQueryParams struct {
	PerPage       int    `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 10 - 1000. Default is 1000.
	StartingAfter string `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
}
type GetOrganizationSummaryTopAppliancesByUtilizationQueryParams struct {
	T0       string  `url:"t0,omitempty"`       //The beginning of the timespan for the data.
	T1       string  `url:"t1,omitempty"`       //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan float64 `url:"timespan,omitempty"` //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
}
type GetOrganizationSummaryTopClientsByUsageQueryParams struct {
	T0       string  `url:"t0,omitempty"`       //The beginning of the timespan for the data.
	T1       string  `url:"t1,omitempty"`       //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan float64 `url:"timespan,omitempty"` //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
}
type GetOrganizationSummaryTopClientsManufacturersByUsageQueryParams struct {
	T0       string  `url:"t0,omitempty"`       //The beginning of the timespan for the data.
	T1       string  `url:"t1,omitempty"`       //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan float64 `url:"timespan,omitempty"` //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
}
type GetOrganizationSummaryTopDevicesByUsageQueryParams struct {
	T0       string  `url:"t0,omitempty"`       //The beginning of the timespan for the data.
	T1       string  `url:"t1,omitempty"`       //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan float64 `url:"timespan,omitempty"` //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
}
type GetOrganizationSummaryTopDevicesModelsByUsageQueryParams struct {
	T0       string  `url:"t0,omitempty"`       //The beginning of the timespan for the data.
	T1       string  `url:"t1,omitempty"`       //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan float64 `url:"timespan,omitempty"` //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
}
type GetOrganizationSummaryTopSSIDsByUsageQueryParams struct {
	T0       string  `url:"t0,omitempty"`       //The beginning of the timespan for the data.
	T1       string  `url:"t1,omitempty"`       //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan float64 `url:"timespan,omitempty"` //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
}
type GetOrganizationSummaryTopSwitchesByEnergyUsageQueryParams struct {
	T0       string  `url:"t0,omitempty"`       //The beginning of the timespan for the data.
	T1       string  `url:"t1,omitempty"`       //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan float64 `url:"timespan,omitempty"` //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
}
type GetOrganizationUplinksStatusesQueryParams struct {
	PerPage       int      `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter string   `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string   `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	NetworkIDs    []string `url:"networkIds[],omitempty"`  //A list of network IDs. The returned devices will be filtered to only include these networks.
	Serials       []string `url:"serials[],omitempty"`     //A list of serial numbers. The returned devices will be filtered to only include these serials.
	Iccids        []string `url:"iccids[],omitempty"`      //A list of ICCIDs. The returned devices will be filtered to only include these ICCIDs.
}
type GetOrganizationWebhooksAlertTypesQueryParams struct {
	ProductType string `url:"productType,omitempty"` //Filter sample alerts to a specific product type
}
type GetOrganizationWebhooksLogsQueryParams struct {
	T0            string  `url:"t0,omitempty"`            //The beginning of the timespan for the data. The maximum lookback period is 90 days from today.
	T1            string  `url:"t1,omitempty"`            //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan      float64 `url:"timespan,omitempty"`      //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
	PerPage       int     `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50.
	StartingAfter string  `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string  `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	URL           string  `url:"url,omitempty"`           //The URL the webhook was sent to
}

type ResponseOrganizationsGetOrganizations []ResponseItemOrganizationsGetOrganizations // Array of ResponseOrganizationsGetOrganizations
type ResponseItemOrganizationsGetOrganizations struct {
	API        *ResponseItemOrganizationsGetOrganizationsAPI        `json:"api,omitempty"`        // API related settings
	Cloud      *ResponseItemOrganizationsGetOrganizationsCloud      `json:"cloud,omitempty"`      // Data for this organization
	ID         string                                               `json:"id,omitempty"`         // Organization ID
	Licensing  *ResponseItemOrganizationsGetOrganizationsLicensing  `json:"licensing,omitempty"`  // Licensing related settings
	Management *ResponseItemOrganizationsGetOrganizationsManagement `json:"management,omitempty"` // Information about the organization's management system
	Name       string                                               `json:"name,omitempty"`       // Organization name
	URL        string                                               `json:"url,omitempty"`        // Organization URL
}
type ResponseItemOrganizationsGetOrganizationsAPI struct {
	Enabled *bool `json:"enabled,omitempty"` // Enable API access
}
type ResponseItemOrganizationsGetOrganizationsCloud struct {
	Region *ResponseItemOrganizationsGetOrganizationsCloudRegion `json:"region,omitempty"` // Region info
}
type ResponseItemOrganizationsGetOrganizationsCloudRegion struct {
	Name string `json:"name,omitempty"` // Name of region
}
type ResponseItemOrganizationsGetOrganizationsLicensing struct {
	Model string `json:"model,omitempty"` // Organization licensing model. Can be 'co-term', 'per-device', or 'subscription'.
}
type ResponseItemOrganizationsGetOrganizationsManagement struct {
	Details *[]ResponseItemOrganizationsGetOrganizationsManagementDetails `json:"details,omitempty"` // Details related to organization management, possibly empty
}
type ResponseItemOrganizationsGetOrganizationsManagementDetails struct {
	Name  string `json:"name,omitempty"`  // Name of management data
	Value string `json:"value,omitempty"` // Value of management data
}
type ResponseOrganizationsCreateOrganization struct {
	API        *ResponseOrganizationsCreateOrganizationAPI        `json:"api,omitempty"`        // API related settings
	Cloud      *ResponseOrganizationsCreateOrganizationCloud      `json:"cloud,omitempty"`      // Data for this organization
	ID         string                                             `json:"id,omitempty"`         // Organization ID
	Licensing  *ResponseOrganizationsCreateOrganizationLicensing  `json:"licensing,omitempty"`  // Licensing related settings
	Management *ResponseOrganizationsCreateOrganizationManagement `json:"management,omitempty"` // Information about the organization's management system
	Name       string                                             `json:"name,omitempty"`       // Organization name
	URL        string                                             `json:"url,omitempty"`        // Organization URL
}
type ResponseOrganizationsCreateOrganizationAPI struct {
	Enabled *bool `json:"enabled,omitempty"` // Enable API access
}
type ResponseOrganizationsCreateOrganizationCloud struct {
	Region *ResponseOrganizationsCreateOrganizationCloudRegion `json:"region,omitempty"` // Region info
}
type ResponseOrganizationsCreateOrganizationCloudRegion struct {
	Name string `json:"name,omitempty"` // Name of region
}
type ResponseOrganizationsCreateOrganizationLicensing struct {
	Model string `json:"model,omitempty"` // Organization licensing model. Can be 'co-term', 'per-device', or 'subscription'.
}
type ResponseOrganizationsCreateOrganizationManagement struct {
	Details *[]ResponseOrganizationsCreateOrganizationManagementDetails `json:"details,omitempty"` // Details related to organization management, possibly empty
}
type ResponseOrganizationsCreateOrganizationManagementDetails struct {
	Name  string `json:"name,omitempty"`  // Name of management data
	Value string `json:"value,omitempty"` // Value of management data
}
type ResponseOrganizationsGetOrganization struct {
	API        *ResponseOrganizationsGetOrganizationAPI        `json:"api,omitempty"`        // API related settings
	Cloud      *ResponseOrganizationsGetOrganizationCloud      `json:"cloud,omitempty"`      // Data for this organization
	ID         string                                          `json:"id,omitempty"`         // Organization ID
	Licensing  *ResponseOrganizationsGetOrganizationLicensing  `json:"licensing,omitempty"`  // Licensing related settings
	Management *ResponseOrganizationsGetOrganizationManagement `json:"management,omitempty"` // Information about the organization's management system
	Name       string                                          `json:"name,omitempty"`       // Organization name
	URL        string                                          `json:"url,omitempty"`        // Organization URL
}
type ResponseOrganizationsGetOrganizationAPI struct {
	Enabled *bool `json:"enabled,omitempty"` // Enable API access
}
type ResponseOrganizationsGetOrganizationCloud struct {
	Region *ResponseOrganizationsGetOrganizationCloudRegion `json:"region,omitempty"` // Region info
}
type ResponseOrganizationsGetOrganizationCloudRegion struct {
	Name string `json:"name,omitempty"` // Name of region
}
type ResponseOrganizationsGetOrganizationLicensing struct {
	Model string `json:"model,omitempty"` // Organization licensing model. Can be 'co-term', 'per-device', or 'subscription'.
}
type ResponseOrganizationsGetOrganizationManagement struct {
	Details *[]ResponseOrganizationsGetOrganizationManagementDetails `json:"details,omitempty"` // Details related to organization management, possibly empty
}
type ResponseOrganizationsGetOrganizationManagementDetails struct {
	Name  string `json:"name,omitempty"`  // Name of management data
	Value string `json:"value,omitempty"` // Value of management data
}
type ResponseOrganizationsUpdateOrganization struct {
	API        *ResponseOrganizationsUpdateOrganizationAPI        `json:"api,omitempty"`        // API related settings
	Cloud      *ResponseOrganizationsUpdateOrganizationCloud      `json:"cloud,omitempty"`      // Data for this organization
	ID         string                                             `json:"id,omitempty"`         // Organization ID
	Licensing  *ResponseOrganizationsUpdateOrganizationLicensing  `json:"licensing,omitempty"`  // Licensing related settings
	Management *ResponseOrganizationsUpdateOrganizationManagement `json:"management,omitempty"` // Information about the organization's management system
	Name       string                                             `json:"name,omitempty"`       // Organization name
	URL        string                                             `json:"url,omitempty"`        // Organization URL
}
type ResponseOrganizationsUpdateOrganizationAPI struct {
	Enabled *bool `json:"enabled,omitempty"` // Enable API access
}
type ResponseOrganizationsUpdateOrganizationCloud struct {
	Region *ResponseOrganizationsUpdateOrganizationCloudRegion `json:"region,omitempty"` // Region info
}
type ResponseOrganizationsUpdateOrganizationCloudRegion struct {
	Name string `json:"name,omitempty"` // Name of region
}
type ResponseOrganizationsUpdateOrganizationLicensing struct {
	Model string `json:"model,omitempty"` // Organization licensing model. Can be 'co-term', 'per-device', or 'subscription'.
}
type ResponseOrganizationsUpdateOrganizationManagement struct {
	Details *[]ResponseOrganizationsUpdateOrganizationManagementDetails `json:"details,omitempty"` // Details related to organization management, possibly empty
}
type ResponseOrganizationsUpdateOrganizationManagementDetails struct {
	Name  string `json:"name,omitempty"`  // Name of management data
	Value string `json:"value,omitempty"` // Value of management data
}
type ResponseOrganizationsGetOrganizationActionBatches []ResponseItemOrganizationsGetOrganizationActionBatches // Array of ResponseOrganizationsGetOrganizationActionBatches
type ResponseItemOrganizationsGetOrganizationActionBatches struct {
	Actions        *[]ResponseItemOrganizationsGetOrganizationActionBatchesActions `json:"actions,omitempty"`        //
	Confirmed      *bool                                                           `json:"confirmed,omitempty"`      //
	ID             string                                                          `json:"id,omitempty"`             //
	OrganizationID string                                                          `json:"organizationId,omitempty"` //
	Status         *ResponseItemOrganizationsGetOrganizationActionBatchesStatus    `json:"status,omitempty"`         //
	Synchronous    *bool                                                           `json:"synchronous,omitempty"`    //
}
type ResponseItemOrganizationsGetOrganizationActionBatchesActions struct {
	Body      *ResponseItemOrganizationsGetOrganizationActionBatchesActionsBody `json:"body,omitempty"`      //
	Operation string                                                            `json:"operation,omitempty"` //
	Resource  string                                                            `json:"resource,omitempty"`  //
}
type ResponseItemOrganizationsGetOrganizationActionBatchesActionsBody struct {
	Enabled *bool `json:"enabled,omitempty"` //
}
type ResponseItemOrganizationsGetOrganizationActionBatchesStatus struct {
	Completed        *bool                                                                          `json:"completed,omitempty"`        //
	CreatedResources *[]ResponseItemOrganizationsGetOrganizationActionBatchesStatusCreatedResources `json:"createdResources,omitempty"` //
	Errors           []string                                                                       `json:"errors,omitempty"`           //
	Failed           *bool                                                                          `json:"failed,omitempty"`           //
}
type ResponseItemOrganizationsGetOrganizationActionBatchesStatusCreatedResources struct {
	ID  *int   `json:"id,omitempty"`  //
	URI string `json:"uri,omitempty"` //
}
type ResponseOrganizationsCreateOrganizationActionBatch struct {
	Actions        *[]ResponseOrganizationsCreateOrganizationActionBatchActions `json:"actions,omitempty"`        // A set of changes made as part of this action (<a href='https://developer.cisco.com/meraki/api/#/rest/guides/action-batches/'>more details</a>)
	Confirmed      *bool                                                        `json:"confirmed,omitempty"`      // Flag describing whether the action should be previewed before executing or not
	ID             string                                                       `json:"id,omitempty"`             // ID of the action batch. Can be used to check the status of the action batch at /organizations/{organizationId}/actionBatches/{actionBatchId}
	OrganizationID string                                                       `json:"organizationId,omitempty"` // ID of the organization this action batch belongs to
	Status         *ResponseOrganizationsCreateOrganizationActionBatchStatus    `json:"status,omitempty"`         // Status of action batch
	Synchronous    *bool                                                        `json:"synchronous,omitempty"`    // Flag describing whether actions should run synchronously or asynchronously
}
type ResponseOrganizationsCreateOrganizationActionBatchActions struct {
	Operation string `json:"operation,omitempty"` // The operation to be used by this action
	Resource  string `json:"resource,omitempty"`  // Unique identifier for the resource to be acted on
}
type ResponseOrganizationsCreateOrganizationActionBatchStatus struct {
	Completed        *bool                                                                       `json:"completed,omitempty"`        // Flag describing whether all actions in the action batch have completed
	CreatedResources *[]ResponseOrganizationsCreateOrganizationActionBatchStatusCreatedResources `json:"createdResources,omitempty"` // Resources created as a result of this action batch
	Errors           []string                                                                    `json:"errors,omitempty"`           // List of errors encountered when running actions in the action batch
	Failed           *bool                                                                       `json:"failed,omitempty"`           // Flag describing whether any actions in the action batch failed
}
type ResponseOrganizationsCreateOrganizationActionBatchStatusCreatedResources struct {
	ID  string `json:"id,omitempty"`  // ID of the created resource
	URI string `json:"uri,omitempty"` // URI, not including base, of the created resource
}
type ResponseOrganizationsGetOrganizationActionBatch struct {
	Actions        *[]ResponseOrganizationsGetOrganizationActionBatchActions `json:"actions,omitempty"`        // A set of changes made as part of this action (<a href='https://developer.cisco.com/meraki/api/#/rest/guides/action-batches/'>more details</a>)
	Confirmed      *bool                                                     `json:"confirmed,omitempty"`      // Flag describing whether the action should be previewed before executing or not
	ID             string                                                    `json:"id,omitempty"`             // ID of the action batch. Can be used to check the status of the action batch at /organizations/{organizationId}/actionBatches/{actionBatchId}
	OrganizationID string                                                    `json:"organizationId,omitempty"` // ID of the organization this action batch belongs to
	Status         *ResponseOrganizationsGetOrganizationActionBatchStatus    `json:"status,omitempty"`         // Status of action batch
	Synchronous    *bool                                                     `json:"synchronous,omitempty"`    // Flag describing whether actions should run synchronously or asynchronously
}
type ResponseOrganizationsGetOrganizationActionBatchActions struct {
	Operation string `json:"operation,omitempty"` // The operation to be used by this action
	Resource  string `json:"resource,omitempty"`  // Unique identifier for the resource to be acted on
}
type ResponseOrganizationsGetOrganizationActionBatchStatus struct {
	Completed        *bool                                                                    `json:"completed,omitempty"`        // Flag describing whether all actions in the action batch have completed
	CreatedResources *[]ResponseOrganizationsGetOrganizationActionBatchStatusCreatedResources `json:"createdResources,omitempty"` // Resources created as a result of this action batch
	Errors           []string                                                                 `json:"errors,omitempty"`           // List of errors encountered when running actions in the action batch
	Failed           *bool                                                                    `json:"failed,omitempty"`           // Flag describing whether any actions in the action batch failed
}
type ResponseOrganizationsGetOrganizationActionBatchStatusCreatedResources struct {
	ID  string `json:"id,omitempty"`  // ID of the created resource
	URI string `json:"uri,omitempty"` // URI, not including base, of the created resource
}
type ResponseOrganizationsUpdateOrganizationActionBatch interface{}
type ResponseOrganizationsGetOrganizationAdaptivePolicyACLs []ResponseItemOrganizationsGetOrganizationAdaptivePolicyACLs // Array of ResponseOrganizationsGetOrganizationAdaptivePolicyAcls
type ResponseItemOrganizationsGetOrganizationAdaptivePolicyACLs struct {
	ACLID       string                                                             `json:"aclId,omitempty"`       // ID of the adaptive policy ACL
	CreatedAt   string                                                             `json:"createdAt,omitempty"`   // When the adaptive policy ACL was created
	Description string                                                             `json:"description,omitempty"` // Description of the adaptive policy ACL
	IPVersion   string                                                             `json:"ipVersion,omitempty"`   // IP version of adpative policy ACL
	Name        string                                                             `json:"name,omitempty"`        // Name of the adaptive policy ACL
	Rules       *[]ResponseItemOrganizationsGetOrganizationAdaptivePolicyACLsRules `json:"rules,omitempty"`       // An ordered array of the adaptive policy ACL rules
	UpdatedAt   string                                                             `json:"updatedAt,omitempty"`   // When the adaptive policy ACL was last updated
}
type ResponseItemOrganizationsGetOrganizationAdaptivePolicyACLsRules struct {
	DstPort  string `json:"dstPort,omitempty"`  // Destination port
	Policy   string `json:"policy,omitempty"`   // 'allow' or 'deny' traffic specified by this rule
	Protocol string `json:"protocol,omitempty"` // The type of protocol
	SrcPort  string `json:"srcPort,omitempty"`  // Source port
}
type ResponseOrganizationsCreateOrganizationAdaptivePolicyACL struct {
	ACLID       string                                                           `json:"aclId,omitempty"`       // ID of the adaptive policy ACL
	CreatedAt   string                                                           `json:"createdAt,omitempty"`   // When the adaptive policy ACL was created
	Description string                                                           `json:"description,omitempty"` // Description of the adaptive policy ACL
	IPVersion   string                                                           `json:"ipVersion,omitempty"`   // IP version of adpative policy ACL
	Name        string                                                           `json:"name,omitempty"`        // Name of the adaptive policy ACL
	Rules       *[]ResponseOrganizationsCreateOrganizationAdaptivePolicyACLRules `json:"rules,omitempty"`       // An ordered array of the adaptive policy ACL rules
	UpdatedAt   string                                                           `json:"updatedAt,omitempty"`   // When the adaptive policy ACL was last updated
}
type ResponseOrganizationsCreateOrganizationAdaptivePolicyACLRules struct {
	DstPort  string `json:"dstPort,omitempty"`  // Destination port
	Policy   string `json:"policy,omitempty"`   // 'allow' or 'deny' traffic specified by this rule
	Protocol string `json:"protocol,omitempty"` // The type of protocol
	SrcPort  string `json:"srcPort,omitempty"`  // Source port
}
type ResponseOrganizationsGetOrganizationAdaptivePolicyACL struct {
	ACLID       string                                                        `json:"aclId,omitempty"`       // ID of the adaptive policy ACL
	CreatedAt   string                                                        `json:"createdAt,omitempty"`   // When the adaptive policy ACL was created
	Description string                                                        `json:"description,omitempty"` // Description of the adaptive policy ACL
	IPVersion   string                                                        `json:"ipVersion,omitempty"`   // IP version of adpative policy ACL
	Name        string                                                        `json:"name,omitempty"`        // Name of the adaptive policy ACL
	Rules       *[]ResponseOrganizationsGetOrganizationAdaptivePolicyACLRules `json:"rules,omitempty"`       // An ordered array of the adaptive policy ACL rules
	UpdatedAt   string                                                        `json:"updatedAt,omitempty"`   // When the adaptive policy ACL was last updated
}
type ResponseOrganizationsGetOrganizationAdaptivePolicyACLRules struct {
	DstPort  string `json:"dstPort,omitempty"`  // Destination port
	Policy   string `json:"policy,omitempty"`   // 'allow' or 'deny' traffic specified by this rule
	Protocol string `json:"protocol,omitempty"` // The type of protocol
	SrcPort  string `json:"srcPort,omitempty"`  // Source port
}
type ResponseOrganizationsUpdateOrganizationAdaptivePolicyACL struct {
	ACLID       string                                                           `json:"aclId,omitempty"`       // ID of the adaptive policy ACL
	CreatedAt   string                                                           `json:"createdAt,omitempty"`   // When the adaptive policy ACL was created
	Description string                                                           `json:"description,omitempty"` // Description of the adaptive policy ACL
	IPVersion   string                                                           `json:"ipVersion,omitempty"`   // IP version of adpative policy ACL
	Name        string                                                           `json:"name,omitempty"`        // Name of the adaptive policy ACL
	Rules       *[]ResponseOrganizationsUpdateOrganizationAdaptivePolicyACLRules `json:"rules,omitempty"`       // An ordered array of the adaptive policy ACL rules
	UpdatedAt   string                                                           `json:"updatedAt,omitempty"`   // When the adaptive policy ACL was last updated
}
type ResponseOrganizationsUpdateOrganizationAdaptivePolicyACLRules struct {
	DstPort  string `json:"dstPort,omitempty"`  // Destination port
	Policy   string `json:"policy,omitempty"`   // 'allow' or 'deny' traffic specified by this rule
	Protocol string `json:"protocol,omitempty"` // The type of protocol
	SrcPort  string `json:"srcPort,omitempty"`  // Source port
}
type ResponseOrganizationsGetOrganizationAdaptivePolicyGroups []ResponseItemOrganizationsGetOrganizationAdaptivePolicyGroups // Array of ResponseOrganizationsGetOrganizationAdaptivePolicyGroups
type ResponseItemOrganizationsGetOrganizationAdaptivePolicyGroups struct {
	CreatedAt          string                                                                       `json:"createdAt,omitempty"`          //
	Description        string                                                                       `json:"description,omitempty"`        //
	GroupID            string                                                                       `json:"groupId,omitempty"`            //
	IsDefaultGroup     *bool                                                                        `json:"isDefaultGroup,omitempty"`     //
	Name               string                                                                       `json:"name,omitempty"`               //
	PolicyObjects      *[]ResponseItemOrganizationsGetOrganizationAdaptivePolicyGroupsPolicyObjects `json:"policyObjects,omitempty"`      //
	RequiredIPMappings []string                                                                     `json:"requiredIpMappings,omitempty"` //
	Sgt                *int                                                                         `json:"sgt,omitempty"`                //
	UpdatedAt          string                                                                       `json:"updatedAt,omitempty"`          //
}
type ResponseItemOrganizationsGetOrganizationAdaptivePolicyGroupsPolicyObjects struct {
	ID   string `json:"id,omitempty"`   //
	Name string `json:"name,omitempty"` //
}
type ResponseOrganizationsCreateOrganizationAdaptivePolicyGroup interface{}
type ResponseOrganizationsGetOrganizationAdaptivePolicyGroup struct {
	CreatedAt          string                                                                  `json:"createdAt,omitempty"`          //
	Description        string                                                                  `json:"description,omitempty"`        //
	GroupID            string                                                                  `json:"groupId,omitempty"`            //
	IsDefaultGroup     *bool                                                                   `json:"isDefaultGroup,omitempty"`     //
	Name               string                                                                  `json:"name,omitempty"`               //
	PolicyObjects      *[]ResponseOrganizationsGetOrganizationAdaptivePolicyGroupPolicyObjects `json:"policyObjects,omitempty"`      //
	RequiredIPMappings []string                                                                `json:"requiredIpMappings,omitempty"` //
	Sgt                *int                                                                    `json:"sgt,omitempty"`                //
	UpdatedAt          string                                                                  `json:"updatedAt,omitempty"`          //
}
type ResponseOrganizationsGetOrganizationAdaptivePolicyGroupPolicyObjects struct {
	ID   string `json:"id,omitempty"`   //
	Name string `json:"name,omitempty"` //
}
type ResponseOrganizationsUpdateOrganizationAdaptivePolicyGroup interface{}
type ResponseOrganizationsGetOrganizationAdaptivePolicyOverview struct {
	Counts *ResponseOrganizationsGetOrganizationAdaptivePolicyOverviewCounts `json:"counts,omitempty"` // The current amount of various adaptive policy objects.
	Limits *ResponseOrganizationsGetOrganizationAdaptivePolicyOverviewLimits `json:"limits,omitempty"` // The current limits of various adaptive policy objects.
}
type ResponseOrganizationsGetOrganizationAdaptivePolicyOverviewCounts struct {
	AllowPolicies *int `json:"allowPolicies,omitempty"` // Number of adaptive policies currently in the organization that allow all traffic.
	CustomACLs    *int `json:"customAcls,omitempty"`    // Number of user-created adaptive policy ACLs currently in the organization.
	CustomGroups  *int `json:"customGroups,omitempty"`  // Number of user-created adaptive policy groups currently in the organization.
	DenyPolicies  *int `json:"denyPolicies,omitempty"`  // Number of adaptive policies currently in the organization that deny all traffic.
	Groups        *int `json:"groups,omitempty"`        // Number of adaptive policy groups currently in the organization.
	Policies      *int `json:"policies,omitempty"`      // Number of adaptive policies currently in the organization.
	PolicyObjects *int `json:"policyObjects,omitempty"` // Number of policy objects (with the adaptive policy type) currently in the organization.
}
type ResponseOrganizationsGetOrganizationAdaptivePolicyOverviewLimits struct {
	ACLsInAPolicy *int `json:"aclsInAPolicy,omitempty"` // Maximum number of adaptive policy ACLs that can be assigned to an adaptive policy in the organization.
	CustomGroups  *int `json:"customGroups,omitempty"`  // Maximum number of user-created adaptive policy groups allowed in the organization.
	PolicyObjects *int `json:"policyObjects,omitempty"` // Maximum number of policy objects (with the adaptive policy type) allowed in the organization.
	RulesInAnACL  *int `json:"rulesInAnAcl,omitempty"`  // Maximum number of rules allowed in an adaptive policy ACL in the organization.
}
type ResponseOrganizationsGetOrganizationAdaptivePolicyPolicies []ResponseItemOrganizationsGetOrganizationAdaptivePolicyPolicies // Array of ResponseOrganizationsGetOrganizationAdaptivePolicyPolicies
type ResponseItemOrganizationsGetOrganizationAdaptivePolicyPolicies struct {
	ACLs             *[]ResponseItemOrganizationsGetOrganizationAdaptivePolicyPoliciesACLs           `json:"acls,omitempty"`             //
	AdaptivePolicyID string                                                                          `json:"adaptivePolicyId,omitempty"` //
	CreatedAt        string                                                                          `json:"createdAt,omitempty"`        //
	DestinationGroup *ResponseItemOrganizationsGetOrganizationAdaptivePolicyPoliciesDestinationGroup `json:"destinationGroup,omitempty"` //
	LastEntryRule    string                                                                          `json:"lastEntryRule,omitempty"`    //
	SourceGroup      *ResponseItemOrganizationsGetOrganizationAdaptivePolicyPoliciesSourceGroup      `json:"sourceGroup,omitempty"`      //
	UpdatedAt        string                                                                          `json:"updatedAt,omitempty"`        //
}
type ResponseItemOrganizationsGetOrganizationAdaptivePolicyPoliciesACLs struct {
	ID   string `json:"id,omitempty"`   //
	Name string `json:"name,omitempty"` //
}
type ResponseItemOrganizationsGetOrganizationAdaptivePolicyPoliciesDestinationGroup struct {
	ID   string `json:"id,omitempty"`   //
	Name string `json:"name,omitempty"` //
	Sgt  *int   `json:"sgt,omitempty"`  //
}
type ResponseItemOrganizationsGetOrganizationAdaptivePolicyPoliciesSourceGroup struct {
	ID   string `json:"id,omitempty"`   //
	Name string `json:"name,omitempty"` //
	Sgt  *int   `json:"sgt,omitempty"`  //
}
type ResponseOrganizationsCreateOrganizationAdaptivePolicyPolicy struct {
	ACLs             *[]ResponseOrganizationsCreateOrganizationAdaptivePolicyPolicyACLs           `json:"acls,omitempty"`             //
	AdaptivePolicyID string                                                                       `json:"adaptivePolicyId,omitempty"` //
	CreatedAt        string                                                                       `json:"createdAt,omitempty"`        //
	DestinationGroup *ResponseOrganizationsCreateOrganizationAdaptivePolicyPolicyDestinationGroup `json:"destinationGroup,omitempty"` //
	LastEntryRule    string                                                                       `json:"lastEntryRule,omitempty"`    //
	SourceGroup      *ResponseOrganizationsCreateOrganizationAdaptivePolicyPolicySourceGroup      `json:"sourceGroup,omitempty"`      //
	UpdatedAt        string                                                                       `json:"updatedAt,omitempty"`        //
}
type ResponseOrganizationsCreateOrganizationAdaptivePolicyPolicyACLs struct {
	ID   string `json:"id,omitempty"`   //
	Name string `json:"name,omitempty"` //
}
type ResponseOrganizationsCreateOrganizationAdaptivePolicyPolicyDestinationGroup struct {
	ID   string `json:"id,omitempty"`   //
	Name string `json:"name,omitempty"` //
	Sgt  *int   `json:"sgt,omitempty"`  //
}
type ResponseOrganizationsCreateOrganizationAdaptivePolicyPolicySourceGroup struct {
	ID   string `json:"id,omitempty"`   //
	Name string `json:"name,omitempty"` //
	Sgt  *int   `json:"sgt,omitempty"`  //
}
type ResponseOrganizationsGetOrganizationAdaptivePolicyPolicy struct {
	ACLs             *[]ResponseOrganizationsGetOrganizationAdaptivePolicyPolicyACLs           `json:"acls,omitempty"`             //
	AdaptivePolicyID string                                                                    `json:"adaptivePolicyId,omitempty"` //
	CreatedAt        string                                                                    `json:"createdAt,omitempty"`        //
	DestinationGroup *ResponseOrganizationsGetOrganizationAdaptivePolicyPolicyDestinationGroup `json:"destinationGroup,omitempty"` //
	LastEntryRule    string                                                                    `json:"lastEntryRule,omitempty"`    //
	SourceGroup      *ResponseOrganizationsGetOrganizationAdaptivePolicyPolicySourceGroup      `json:"sourceGroup,omitempty"`      //
	UpdatedAt        string                                                                    `json:"updatedAt,omitempty"`        //
}
type ResponseOrganizationsGetOrganizationAdaptivePolicyPolicyACLs struct {
	ID   string `json:"id,omitempty"`   //
	Name string `json:"name,omitempty"` //
}
type ResponseOrganizationsGetOrganizationAdaptivePolicyPolicyDestinationGroup struct {
	ID   string `json:"id,omitempty"`   //
	Name string `json:"name,omitempty"` //
	Sgt  *int   `json:"sgt,omitempty"`  //
}
type ResponseOrganizationsGetOrganizationAdaptivePolicyPolicySourceGroup struct {
	ID   string `json:"id,omitempty"`   //
	Name string `json:"name,omitempty"` //
	Sgt  *int   `json:"sgt,omitempty"`  //
}
type ResponseOrganizationsUpdateOrganizationAdaptivePolicyPolicy interface{}
type ResponseOrganizationsGetOrganizationAdaptivePolicySettings struct {
	EnabledNetworks []string `json:"enabledNetworks,omitempty"` //
}
type ResponseOrganizationsUpdateOrganizationAdaptivePolicySettings interface{}
type ResponseOrganizationsGetOrganizationAdmins []ResponseItemOrganizationsGetOrganizationAdmins // Array of ResponseOrganizationsGetOrganizationAdmins
type ResponseItemOrganizationsGetOrganizationAdmins struct {
	AccountStatus        string                                                    `json:"accountStatus,omitempty"`        // Status of the admin's account
	AuthenticationMethod string                                                    `json:"authenticationMethod,omitempty"` // Admin's authentication method
	Email                string                                                    `json:"email,omitempty"`                // Admin's email address
	HasAPIKey            *bool                                                     `json:"hasApiKey,omitempty"`            // Indicates whether the admin has an API key
	ID                   string                                                    `json:"id,omitempty"`                   // Admin's ID
	LastActive           string                                                    `json:"lastActive,omitempty"`           // Time when the admin was last active
	Name                 string                                                    `json:"name,omitempty"`                 // Admin's username
	Networks             *[]ResponseItemOrganizationsGetOrganizationAdminsNetworks `json:"networks,omitempty"`             // Admin network access information
	OrgAccess            string                                                    `json:"orgAccess,omitempty"`            // Admin's level of access to the organization
	Tags                 *[]ResponseItemOrganizationsGetOrganizationAdminsTags     `json:"tags,omitempty"`                 // Admin tag information
	TwoFactorAuthEnabled *bool                                                     `json:"twoFactorAuthEnabled,omitempty"` // Indicates whether two-factor authentication is enabled
}
type ResponseItemOrganizationsGetOrganizationAdminsNetworks struct {
	Access string `json:"access,omitempty"` // Admin's level of access to the network
	ID     string `json:"id,omitempty"`     // Network ID
}
type ResponseItemOrganizationsGetOrganizationAdminsTags struct {
	Access string `json:"access,omitempty"` // Access level for the tag
	Tag    string `json:"tag,omitempty"`    // Tag value
}
type ResponseOrganizationsCreateOrganizationAdmin struct {
	AccountStatus        string                                                  `json:"accountStatus,omitempty"`        // Status of the admin's account
	AuthenticationMethod string                                                  `json:"authenticationMethod,omitempty"` // Admin's authentication method
	Email                string                                                  `json:"email,omitempty"`                // Admin's email address
	HasAPIKey            *bool                                                   `json:"hasApiKey,omitempty"`            // Indicates whether the admin has an API key
	ID                   string                                                  `json:"id,omitempty"`                   // Admin's ID
	LastActive           string                                                  `json:"lastActive,omitempty"`           // Time when the admin was last active
	Name                 string                                                  `json:"name,omitempty"`                 // Admin's username
	Networks             *[]ResponseOrganizationsCreateOrganizationAdminNetworks `json:"networks,omitempty"`             // Admin network access information
	OrgAccess            string                                                  `json:"orgAccess,omitempty"`            // Admin's level of access to the organization
	Tags                 *[]ResponseOrganizationsCreateOrganizationAdminTags     `json:"tags,omitempty"`                 // Admin tag information
	TwoFactorAuthEnabled *bool                                                   `json:"twoFactorAuthEnabled,omitempty"` // Indicates whether two-factor authentication is enabled
}
type ResponseOrganizationsCreateOrganizationAdminNetworks struct {
	Access string `json:"access,omitempty"` // Admin's level of access to the network
	ID     string `json:"id,omitempty"`     // Network ID
}
type ResponseOrganizationsCreateOrganizationAdminTags struct {
	Access string `json:"access,omitempty"` // Access level for the tag
	Tag    string `json:"tag,omitempty"`    // Tag value
}
type ResponseOrganizationsUpdateOrganizationAdmin struct {
	AccountStatus        string                                                  `json:"accountStatus,omitempty"`        // Status of the admin's account
	AuthenticationMethod string                                                  `json:"authenticationMethod,omitempty"` // Admin's authentication method
	Email                string                                                  `json:"email,omitempty"`                // Admin's email address
	HasAPIKey            *bool                                                   `json:"hasApiKey,omitempty"`            // Indicates whether the admin has an API key
	ID                   string                                                  `json:"id,omitempty"`                   // Admin's ID
	LastActive           string                                                  `json:"lastActive,omitempty"`           // Time when the admin was last active
	Name                 string                                                  `json:"name,omitempty"`                 // Admin's username
	Networks             *[]ResponseOrganizationsUpdateOrganizationAdminNetworks `json:"networks,omitempty"`             // Admin network access information
	OrgAccess            string                                                  `json:"orgAccess,omitempty"`            // Admin's level of access to the organization
	Tags                 *[]ResponseOrganizationsUpdateOrganizationAdminTags     `json:"tags,omitempty"`                 // Admin tag information
	TwoFactorAuthEnabled *bool                                                   `json:"twoFactorAuthEnabled,omitempty"` // Indicates whether two-factor authentication is enabled
}
type ResponseOrganizationsUpdateOrganizationAdminNetworks struct {
	Access string `json:"access,omitempty"` // Admin's level of access to the network
	ID     string `json:"id,omitempty"`     // Network ID
}
type ResponseOrganizationsUpdateOrganizationAdminTags struct {
	Access string `json:"access,omitempty"` // Access level for the tag
	Tag    string `json:"tag,omitempty"`    // Tag value
}
type ResponseOrganizationsGetOrganizationAlertsProfiles []ResponseItemOrganizationsGetOrganizationAlertsProfiles // Array of ResponseOrganizationsGetOrganizationAlertsProfiles
type ResponseItemOrganizationsGetOrganizationAlertsProfiles struct {
	AlertCondition *ResponseItemOrganizationsGetOrganizationAlertsProfilesAlertCondition `json:"alertCondition,omitempty"` //
	Description    string                                                                `json:"description,omitempty"`    //
	Enabled        *bool                                                                 `json:"enabled,omitempty"`        //
	ID             string                                                                `json:"id,omitempty"`             //
	NetworkTags    []string                                                              `json:"networkTags,omitempty"`    //
	Recipients     *ResponseItemOrganizationsGetOrganizationAlertsProfilesRecipients     `json:"recipients,omitempty"`     //
	Type           string                                                                `json:"type,omitempty"`           //
}
type ResponseItemOrganizationsGetOrganizationAlertsProfilesAlertCondition struct {
	BitRateBps *int   `json:"bit_rate_bps,omitempty"` //
	Duration   *int   `json:"duration,omitempty"`     //
	Interface  string `json:"interface,omitempty"`    //
	Window     *int   `json:"window,omitempty"`       //
}
type ResponseItemOrganizationsGetOrganizationAlertsProfilesRecipients struct {
	Emails        []string `json:"emails,omitempty"`        //
	HTTPServerIDs []string `json:"httpServerIds,omitempty"` //
}
type ResponseOrganizationsCreateOrganizationAlertsProfile struct {
	AlertCondition *ResponseOrganizationsCreateOrganizationAlertsProfileAlertCondition `json:"alertCondition,omitempty"` //
	Description    string                                                              `json:"description,omitempty"`    //
	Enabled        *bool                                                               `json:"enabled,omitempty"`        //
	ID             string                                                              `json:"id,omitempty"`             //
	NetworkTags    []string                                                            `json:"networkTags,omitempty"`    //
	Recipients     *ResponseOrganizationsCreateOrganizationAlertsProfileRecipients     `json:"recipients,omitempty"`     //
	Type           string                                                              `json:"type,omitempty"`           //
}
type ResponseOrganizationsCreateOrganizationAlertsProfileAlertCondition struct {
	BitRateBps *int   `json:"bit_rate_bps,omitempty"` //
	Duration   *int   `json:"duration,omitempty"`     //
	Interface  string `json:"interface,omitempty"`    //
	Window     *int   `json:"window,omitempty"`       //
}
type ResponseOrganizationsCreateOrganizationAlertsProfileRecipients struct {
	Emails        []string `json:"emails,omitempty"`        //
	HTTPServerIDs []string `json:"httpServerIds,omitempty"` //
}
type ResponseOrganizationsUpdateOrganizationAlertsProfile interface{}
type ResponseOrganizationsGetOrganizationAPIRequests []ResponseItemOrganizationsGetOrganizationAPIRequests // Array of ResponseOrganizationsGetOrganizationApiRequests
type ResponseItemOrganizationsGetOrganizationAPIRequests struct {
	AdminID      string `json:"adminId,omitempty"`      // Database ID for the admin user who made the API request.
	Host         string `json:"host,omitempty"`         // The host which the API request was directed at.
	Method       string `json:"method,omitempty"`       // HTTP method used in the API request.
	OperationID  string `json:"operationId,omitempty"`  // Operation ID for the endpoint.
	Path         string `json:"path,omitempty"`         // The API request path.
	QueryString  string `json:"queryString,omitempty"`  // The query string sent with the API request.
	ResponseCode *int   `json:"responseCode,omitempty"` // API request response code.
	SourceIP     string `json:"sourceIp,omitempty"`     // Public IP address from which the API request was made.
	Ts           string `json:"ts,omitempty"`           // Timestamp, in iso8601 format, indicating when the API request was made.
	UserAgent    string `json:"userAgent,omitempty"`    // The API request user agent.
	Version      *int   `json:"version,omitempty"`      // API version of the endpoint.
}
type ResponseOrganizationsGetOrganizationAPIRequestsOverview struct {
	ResponseCodeCounts *ResponseOrganizationsGetOrganizationAPIRequestsOverviewResponseCodeCounts `json:"responseCodeCounts,omitempty"` // object of all supported HTTP response code
}
type ResponseOrganizationsGetOrganizationAPIRequestsOverviewResponseCodeCounts struct {
	Status200 *int `json:"200,omitempty"` // HTTP 200 response code count.
	Status201 *int `json:"201,omitempty"` // HTTP 201 response code count.
	Status202 *int `json:"202,omitempty"` // HTTP 202 response code count.
	Status203 *int `json:"203,omitempty"` // HTTP 203 response code count.
	Status204 *int `json:"204,omitempty"` // HTTP 204 response code count.
	Status205 *int `json:"205,omitempty"` // HTTP 205 response code count.
	Status206 *int `json:"206,omitempty"` // HTTP 206 response code count.
	Status207 *int `json:"207,omitempty"` // HTTP 207 response code count.
	Status208 *int `json:"208,omitempty"` // HTTP 208 response code count.
	Status226 *int `json:"226,omitempty"` // HTTP 226 response code count.
	Status300 *int `json:"300,omitempty"` // HTTP 300 response code count.
	Status301 *int `json:"301,omitempty"` // HTTP 301 response code count.
	Status302 *int `json:"302,omitempty"` // HTTP 302 response code count.
	Status303 *int `json:"303,omitempty"` // HTTP 303 response code count.
	Status304 *int `json:"304,omitempty"` // HTTP 304 response code count.
	Status305 *int `json:"305,omitempty"` // HTTP 305 response code count.
	Status306 *int `json:"306,omitempty"` // HTTP 306 response code count.
	Status307 *int `json:"307,omitempty"` // HTTP 307 response code count.
	Status308 *int `json:"308,omitempty"` // HTTP 308 response code count.
	Status400 *int `json:"400,omitempty"` // HTTP 400 response code count.
	Status401 *int `json:"401,omitempty"` // HTTP 401 response code count.
	Status402 *int `json:"402,omitempty"` // HTTP 402 response code count.
	Status403 *int `json:"403,omitempty"` // HTTP 403 response code count.
	Status404 *int `json:"404,omitempty"` // HTTP 404 response code count.
	Status405 *int `json:"405,omitempty"` // HTTP 405 response code count.
	Status406 *int `json:"406,omitempty"` // HTTP 406 response code count.
	Status407 *int `json:"407,omitempty"` // HTTP 407 response code count.
	Status408 *int `json:"408,omitempty"` // HTTP 408 response code count.
	Status409 *int `json:"409,omitempty"` // HTTP 409 response code count.
	Status410 *int `json:"410,omitempty"` // HTTP 410 response code count.
	Status411 *int `json:"411,omitempty"` // HTTP 411 response code count.
	Status412 *int `json:"412,omitempty"` // HTTP 412 response code count.
	Status413 *int `json:"413,omitempty"` // HTTP 413 response code count.
	Status414 *int `json:"414,omitempty"` // HTTP 414 response code count.
	Status415 *int `json:"415,omitempty"` // HTTP 415 response code count.
	Status416 *int `json:"416,omitempty"` // HTTP 416 response code count.
	Status417 *int `json:"417,omitempty"` // HTTP 417 response code count.
	Status421 *int `json:"421,omitempty"` // HTTP 421 response code count.
	Status422 *int `json:"422,omitempty"` // HTTP 422 response code count.
	Status423 *int `json:"423,omitempty"` // HTTP 423 response code count.
	Status424 *int `json:"424,omitempty"` // HTTP 424 response code count.
	Status425 *int `json:"425,omitempty"` // HTTP 425 response code count.
	Status426 *int `json:"426,omitempty"` // HTTP 426 response code count.
	Status428 *int `json:"428,omitempty"` // HTTP 428 response code count.
	Status429 *int `json:"429,omitempty"` // HTTP 429 response code count.
	Status431 *int `json:"431,omitempty"` // HTTP 431 response code count.
	Status451 *int `json:"451,omitempty"` // HTTP 451 response code count.
	Status500 *int `json:"500,omitempty"` // HTTP 500 response code count.
	Status501 *int `json:"501,omitempty"` // HTTP 501 response code count.
	Status502 *int `json:"502,omitempty"` // HTTP 502 response code count.
	Status503 *int `json:"503,omitempty"` // HTTP 503 response code count.
	Status504 *int `json:"504,omitempty"` // HTTP 504 response code count.
	Status505 *int `json:"505,omitempty"` // HTTP 505 response code count.
	Status506 *int `json:"506,omitempty"` // HTTP 506 response code count.
	Status507 *int `json:"507,omitempty"` // HTTP 507 response code count.
	Status508 *int `json:"508,omitempty"` // HTTP 508 response code count.
	Status509 *int `json:"509,omitempty"` // HTTP 509 response code count.
	Status510 *int `json:"510,omitempty"` // HTTP 510 response code count.
	Status511 *int `json:"511,omitempty"` // HTTP 511 response code count.
}
type ResponseOrganizationsGetOrganizationAPIRequestsOverviewResponseCodesByInterval []ResponseItemOrganizationsGetOrganizationAPIRequestsOverviewResponseCodesByInterval // Array of ResponseOrganizationsGetOrganizationApiRequestsOverviewResponseCodesByInterval
type ResponseItemOrganizationsGetOrganizationAPIRequestsOverviewResponseCodesByInterval struct {
	Counts  *[]ResponseItemOrganizationsGetOrganizationAPIRequestsOverviewResponseCodesByIntervalCounts `json:"counts,omitempty"`  // list of response codes and a count of how many requests had that code in the given time period
	EndTs   string                                                                                      `json:"endTs,omitempty"`   // The end time of the access period
	StartTs string                                                                                      `json:"startTs,omitempty"` // The start time of the access period
}
type ResponseItemOrganizationsGetOrganizationAPIRequestsOverviewResponseCodesByIntervalCounts struct {
	Code  *int `json:"code,omitempty"`  // Response status code of the API response
	Count *int `json:"count,omitempty"` // Number of records that match the status code
}
type ResponseOrganizationsGetOrganizationBrandingPolicies []ResponseItemOrganizationsGetOrganizationBrandingPolicies // Array of ResponseOrganizationsGetOrganizationBrandingPolicies
type ResponseItemOrganizationsGetOrganizationBrandingPolicies struct {
	AdminSettings *ResponseItemOrganizationsGetOrganizationBrandingPoliciesAdminSettings `json:"adminSettings,omitempty"` // Settings for describing which kinds of admins this policy applies to.
	CustomLogo    *ResponseItemOrganizationsGetOrganizationBrandingPoliciesCustomLogo    `json:"customLogo,omitempty"`    // Properties describing the custom logo attached to the branding policy.
	Enabled       *bool                                                                  `json:"enabled,omitempty"`       // Boolean indicating whether this policy is enabled.
	HelpSettings  *ResponseItemOrganizationsGetOrganizationBrandingPoliciesHelpSettings  `json:"helpSettings,omitempty"`  //       Settings for describing the modifications to various Help page features. Each property in this object accepts one of       'default or inherit' (do not modify functionality), 'hide' (remove the section from Dashboard), or 'show' (always show       the section on Dashboard). Some properties in this object also accept custom HTML used to replace the section on       Dashboard; see the documentation for each property to see the allowed values.
	Name          string                                                                 `json:"name,omitempty"`          // Name of the Dashboard branding policy.
}
type ResponseItemOrganizationsGetOrganizationBrandingPoliciesAdminSettings struct {
	AppliesTo string   `json:"appliesTo,omitempty"` // Which kinds of admins this policy applies to. Can be one of 'All organization admins', 'All enterprise admins', 'All network admins', 'All admins of networks...', 'All admins of networks tagged...', 'Specific admins...', 'All admins' or 'All SAML admins'.
	Values    []string `json:"values,omitempty"`    //       If 'appliesTo' is set to one of 'Specific admins...', 'All admins of networks...' or 'All admins of networks tagged...', then you must specify this 'values' property to provide the set of       entities to apply the branding policy to. For 'Specific admins...', specify an array of admin IDs. For 'All admins of       networks...', specify an array of network IDs and/or configuration template IDs. For 'All admins of networks tagged...',       specify an array of tag names.
}
type ResponseItemOrganizationsGetOrganizationBrandingPoliciesCustomLogo struct {
	Enabled *bool                                                                    `json:"enabled,omitempty"` // Whether or not there is a custom logo enabled.
	Image   *ResponseItemOrganizationsGetOrganizationBrandingPoliciesCustomLogoImage `json:"image,omitempty"`   // Properties of the image.
}
type ResponseItemOrganizationsGetOrganizationBrandingPoliciesCustomLogoImage struct {
	Preview *ResponseItemOrganizationsGetOrganizationBrandingPoliciesCustomLogoImagePreview `json:"preview,omitempty"` // Preview of the image
}
type ResponseItemOrganizationsGetOrganizationBrandingPoliciesCustomLogoImagePreview struct {
	ExpiresAt string `json:"expiresAt,omitempty"` // Timestamp of the preview image
	URL       string `json:"url,omitempty"`       // Url of the preview image
}
type ResponseItemOrganizationsGetOrganizationBrandingPoliciesHelpSettings struct {
	APIDocsSubtab                      string `json:"apiDocsSubtab,omitempty"`                      //       The 'Help -> API docs' subtab where a detailed description of the Dashboard API is listed. Can be one of       'default or inherit', 'hide' or 'show'.
	CasesSubtab                        string `json:"casesSubtab,omitempty"`                        //       The 'Help -> Cases' Dashboard subtab on which Cisco Meraki support cases for this organization can be managed. Can be one       of 'default or inherit', 'hide' or 'show'.
	CiscoMerakiProductDocumentation    string `json:"ciscoMerakiProductDocumentation,omitempty"`    //       The 'Product Manuals' section of the 'Help -> Get Help' subtab. Can be one of 'default or inherit', 'hide', 'show', or a replacement custom HTML string.
	CommunitySubtab                    string `json:"communitySubtab,omitempty"`                    //       The 'Help -> Community' subtab which provides a link to Meraki Community. Can be one of 'default or inherit', 'hide' or 'show'.
	DataProtectionRequestsSubtab       string `json:"dataProtectionRequestsSubtab,omitempty"`       //       The 'Help -> Data protection requests' Dashboard subtab on which requests to delete, restrict, or export end-user data can       be audited. Can be one of 'default or inherit', 'hide' or 'show'.
	FirewallInfoSubtab                 string `json:"firewallInfoSubtab,omitempty"`                 //       The 'Help -> Firewall info' subtab where necessary upstream firewall rules for communication to the Cisco Meraki cloud are       listed. Can be one of 'default or inherit', 'hide' or 'show'.
	GetHelpSubtab                      string `json:"getHelpSubtab,omitempty"`                      //       The 'Help -> Get Help' subtab on which Cisco Meraki KB, Product Manuals, and Support/Case Information are displayed. Note       that if this subtab is hidden, branding customizations for the KB on 'Get help', Cisco Meraki product documentation,       and support contact info will not be visible. Can be one of 'default or inherit', 'hide' or 'show'.
	GetHelpSubtabKnowledgeBaseSearch   string `json:"getHelpSubtabKnowledgeBaseSearch,omitempty"`   //       The KB search box which appears on the Help page. Can be one of 'default or inherit', 'hide', 'show', or a replacement custom HTML string.
	HardwareReplacementsSubtab         string `json:"hardwareReplacementsSubtab,omitempty"`         //       The 'Help -> Replacement info' subtab where important information regarding device replacements is detailed. Can be one of       'default or inherit', 'hide' or 'show'.
	HelpTab                            string `json:"helpTab,omitempty"`                            //       The Help tab, under which all support information resides. If this tab is hidden, no other 'Help' branding       customizations will be visible. Can be one of 'default or inherit', 'hide' or 'show'.
	HelpWidget                         string `json:"helpWidget,omitempty"`                         //       The 'Help Widget' is a support widget which provides access to live chat, documentation links, Sales contact info,       and other contact avenues to reach Meraki Support. Can be one of 'default or inherit', 'hide' or 'show'.
	NewFeaturesSubtab                  string `json:"newFeaturesSubtab,omitempty"`                  //       The 'Help -> New features' subtab where new Dashboard features are detailed. Can be one of 'default or inherit', 'hide' or 'show'.
	SmForums                           string `json:"smForums,omitempty"`                           //       The 'SM Forums' subtab which links to community-based support for Cisco Meraki Systems Manager. Only configurable for       organizations that contain Systems Manager networks. Can be one of 'default or inherit', 'hide' or 'show'.
	SupportContactInfo                 string `json:"supportContactInfo,omitempty"`                 //       The 'Contact Meraki Support' section of the 'Help -> Get Help' subtab. Can be one of 'default or inherit', 'hide', 'show', or a replacement custom HTML string.
	UniversalSearchKnowledgeBaseSearch string `json:"universalSearchKnowledgeBaseSearch,omitempty"` //       The universal search box always visible on Dashboard will, by default, present results from the Meraki KB. This configures       whether these Meraki KB results should be returned. Can be one of 'default or inherit', 'hide' or 'show'.
}
type ResponseOrganizationsCreateOrganizationBrandingPolicy struct {
	AdminSettings *ResponseOrganizationsCreateOrganizationBrandingPolicyAdminSettings `json:"adminSettings,omitempty"` // Settings for describing which kinds of admins this policy applies to.
	CustomLogo    *ResponseOrganizationsCreateOrganizationBrandingPolicyCustomLogo    `json:"customLogo,omitempty"`    // Properties describing the custom logo attached to the branding policy.
	Enabled       *bool                                                               `json:"enabled,omitempty"`       // Boolean indicating whether this policy is enabled.
	HelpSettings  *ResponseOrganizationsCreateOrganizationBrandingPolicyHelpSettings  `json:"helpSettings,omitempty"`  //       Settings for describing the modifications to various Help page features. Each property in this object accepts one of       'default or inherit' (do not modify functionality), 'hide' (remove the section from Dashboard), or 'show' (always show       the section on Dashboard). Some properties in this object also accept custom HTML used to replace the section on       Dashboard; see the documentation for each property to see the allowed values.  Each property defaults to 'default or inherit' when not provided.
	Name          string                                                              `json:"name,omitempty"`          // Name of the Dashboard branding policy.
}
type ResponseOrganizationsCreateOrganizationBrandingPolicyAdminSettings struct {
	AppliesTo string   `json:"appliesTo,omitempty"` // Which kinds of admins this policy applies to. Can be one of 'All organization admins', 'All enterprise admins', 'All network admins', 'All admins of networks...', 'All admins of networks tagged...', 'Specific admins...', 'All admins' or 'All SAML admins'.
	Values    []string `json:"values,omitempty"`    //       If 'appliesTo' is set to one of 'Specific admins...', 'All admins of networks...' or 'All admins of networks tagged...', then you must specify this 'values' property to provide the set of       entities to apply the branding policy to. For 'Specific admins...', specify an array of admin IDs. For 'All admins of       networks...', specify an array of network IDs and/or configuration template IDs. For 'All admins of networks tagged...',       specify an array of tag names.
}
type ResponseOrganizationsCreateOrganizationBrandingPolicyCustomLogo struct {
	Enabled *bool                                                                 `json:"enabled,omitempty"` // Whether or not there is a custom logo enabled.
	Image   *ResponseOrganizationsCreateOrganizationBrandingPolicyCustomLogoImage `json:"image,omitempty"`   // Properties of the image.
}
type ResponseOrganizationsCreateOrganizationBrandingPolicyCustomLogoImage struct {
	Preview *ResponseOrganizationsCreateOrganizationBrandingPolicyCustomLogoImagePreview `json:"preview,omitempty"` // Preview of the image
}
type ResponseOrganizationsCreateOrganizationBrandingPolicyCustomLogoImagePreview struct {
	ExpiresAt string `json:"expiresAt,omitempty"` // Timestamp of the preview image
	URL       string `json:"url,omitempty"`       // Url of the preview image
}
type ResponseOrganizationsCreateOrganizationBrandingPolicyHelpSettings struct {
	APIDocsSubtab                      string `json:"apiDocsSubtab,omitempty"`                      //       The 'Help -> API docs' subtab where a detailed description of the Dashboard API is listed. Can be one of       'default or inherit', 'hide' or 'show'.
	CasesSubtab                        string `json:"casesSubtab,omitempty"`                        //       The 'Help -> Cases' Dashboard subtab on which Cisco Meraki support cases for this organization can be managed. Can be one       of 'default or inherit', 'hide' or 'show'.
	CiscoMerakiProductDocumentation    string `json:"ciscoMerakiProductDocumentation,omitempty"`    //       The 'Product Manuals' section of the 'Help -> Get Help' subtab. Can be one of 'default or inherit', 'hide', 'show', or a replacement custom HTML string.
	CommunitySubtab                    string `json:"communitySubtab,omitempty"`                    //       The 'Help -> Community' subtab which provides a link to Meraki Community. Can be one of 'default or inherit', 'hide' or 'show'.
	DataProtectionRequestsSubtab       string `json:"dataProtectionRequestsSubtab,omitempty"`       //       The 'Help -> Data protection requests' Dashboard subtab on which requests to delete, restrict, or export end-user data can       be audited. Can be one of 'default or inherit', 'hide' or 'show'.
	FirewallInfoSubtab                 string `json:"firewallInfoSubtab,omitempty"`                 //       The 'Help -> Firewall info' subtab where necessary upstream firewall rules for communication to the Cisco Meraki cloud are       listed. Can be one of 'default or inherit', 'hide' or 'show'.
	GetHelpSubtab                      string `json:"getHelpSubtab,omitempty"`                      //       The 'Help -> Get Help' subtab on which Cisco Meraki KB, Product Manuals, and Support/Case Information are displayed. Note       that if this subtab is hidden, branding customizations for the KB on 'Get help', Cisco Meraki product documentation,       and support contact info will not be visible. Can be one of 'default or inherit', 'hide' or 'show'.
	GetHelpSubtabKnowledgeBaseSearch   string `json:"getHelpSubtabKnowledgeBaseSearch,omitempty"`   //       The KB search box which appears on the Help page. Can be one of 'default or inherit', 'hide', 'show', or a replacement custom HTML string.
	HardwareReplacementsSubtab         string `json:"hardwareReplacementsSubtab,omitempty"`         //       The 'Help -> Replacement info' subtab where important information regarding device replacements is detailed. Can be one of       'default or inherit', 'hide' or 'show'.
	HelpTab                            string `json:"helpTab,omitempty"`                            //       The Help tab, under which all support information resides. If this tab is hidden, no other 'Help' branding       customizations will be visible. Can be one of 'default or inherit', 'hide' or 'show'.
	HelpWidget                         string `json:"helpWidget,omitempty"`                         //       The 'Help Widget' is a support widget which provides access to live chat, documentation links, Sales contact info,       and other contact avenues to reach Meraki Support. Can be one of 'default or inherit', 'hide' or 'show'.
	NewFeaturesSubtab                  string `json:"newFeaturesSubtab,omitempty"`                  //       The 'Help -> New features' subtab where new Dashboard features are detailed. Can be one of 'default or inherit', 'hide' or 'show'.
	SmForums                           string `json:"smForums,omitempty"`                           //       The 'SM Forums' subtab which links to community-based support for Cisco Meraki Systems Manager. Only configurable for       organizations that contain Systems Manager networks. Can be one of 'default or inherit', 'hide' or 'show'.
	SupportContactInfo                 string `json:"supportContactInfo,omitempty"`                 //       The 'Contact Meraki Support' section of the 'Help -> Get Help' subtab. Can be one of 'default or inherit', 'hide', 'show', or a replacement custom HTML string.
	UniversalSearchKnowledgeBaseSearch string `json:"universalSearchKnowledgeBaseSearch,omitempty"` //       The universal search box always visible on Dashboard will, by default, present results from the Meraki KB. This configures       whether these Meraki KB results should be returned. Can be one of 'default or inherit', 'hide' or 'show'.
}
type ResponseOrganizationsGetOrganizationBrandingPoliciesPriorities struct {
	BrandingPolicyIDs []string `json:"brandingPolicyIds,omitempty"` //       An ordered list of branding policy IDs that determines the priority order of how to apply the policies
}
type ResponseOrganizationsUpdateOrganizationBrandingPoliciesPriorities struct {
	BrandingPolicyIDs []string `json:"brandingPolicyIds,omitempty"` //       An ordered list of branding policy IDs that determines the priority order of how to apply the policies
}
type ResponseOrganizationsGetOrganizationBrandingPolicy struct {
	AdminSettings *ResponseOrganizationsGetOrganizationBrandingPolicyAdminSettings `json:"adminSettings,omitempty"` // Settings for describing which kinds of admins this policy applies to.
	CustomLogo    *ResponseOrganizationsGetOrganizationBrandingPolicyCustomLogo    `json:"customLogo,omitempty"`    // Properties describing the custom logo attached to the branding policy.
	Enabled       *bool                                                            `json:"enabled,omitempty"`       // Boolean indicating whether this policy is enabled.
	HelpSettings  *ResponseOrganizationsGetOrganizationBrandingPolicyHelpSettings  `json:"helpSettings,omitempty"`  //       Settings for describing the modifications to various Help page features. Each property in this object accepts one of       'default or inherit' (do not modify functionality), 'hide' (remove the section from Dashboard), or 'show' (always show       the section on Dashboard). Some properties in this object also accept custom HTML used to replace the section on       Dashboard; see the documentation for each property to see the allowed values.
	Name          string                                                           `json:"name,omitempty"`          // Name of the Dashboard branding policy.
}
type ResponseOrganizationsGetOrganizationBrandingPolicyAdminSettings struct {
	AppliesTo string   `json:"appliesTo,omitempty"` // Which kinds of admins this policy applies to. Can be one of 'All organization admins', 'All enterprise admins', 'All network admins', 'All admins of networks...', 'All admins of networks tagged...', 'Specific admins...', 'All admins' or 'All SAML admins'.
	Values    []string `json:"values,omitempty"`    //       If 'appliesTo' is set to one of 'Specific admins...', 'All admins of networks...' or 'All admins of networks tagged...', then you must specify this 'values' property to provide the set of       entities to apply the branding policy to. For 'Specific admins...', specify an array of admin IDs. For 'All admins of       networks...', specify an array of network IDs and/or configuration template IDs. For 'All admins of networks tagged...',       specify an array of tag names.
}
type ResponseOrganizationsGetOrganizationBrandingPolicyCustomLogo struct {
	Enabled *bool                                                              `json:"enabled,omitempty"` // Whether or not there is a custom logo enabled.
	Image   *ResponseOrganizationsGetOrganizationBrandingPolicyCustomLogoImage `json:"image,omitempty"`   // Properties of the image.
}
type ResponseOrganizationsGetOrganizationBrandingPolicyCustomLogoImage struct {
	Preview *ResponseOrganizationsGetOrganizationBrandingPolicyCustomLogoImagePreview `json:"preview,omitempty"` // Preview of the image
}
type ResponseOrganizationsGetOrganizationBrandingPolicyCustomLogoImagePreview struct {
	ExpiresAt string `json:"expiresAt,omitempty"` // Timestamp of the preview image
	URL       string `json:"url,omitempty"`       // Url of the preview image
}
type ResponseOrganizationsGetOrganizationBrandingPolicyHelpSettings struct {
	APIDocsSubtab                      string `json:"apiDocsSubtab,omitempty"`                      //       The 'Help -> API docs' subtab where a detailed description of the Dashboard API is listed. Can be one of       'default or inherit', 'hide' or 'show'.
	CasesSubtab                        string `json:"casesSubtab,omitempty"`                        //       The 'Help -> Cases' Dashboard subtab on which Cisco Meraki support cases for this organization can be managed. Can be one       of 'default or inherit', 'hide' or 'show'.
	CiscoMerakiProductDocumentation    string `json:"ciscoMerakiProductDocumentation,omitempty"`    //       The 'Product Manuals' section of the 'Help -> Get Help' subtab. Can be one of 'default or inherit', 'hide', 'show', or a replacement custom HTML string.
	CommunitySubtab                    string `json:"communitySubtab,omitempty"`                    //       The 'Help -> Community' subtab which provides a link to Meraki Community. Can be one of 'default or inherit', 'hide' or 'show'.
	DataProtectionRequestsSubtab       string `json:"dataProtectionRequestsSubtab,omitempty"`       //       The 'Help -> Data protection requests' Dashboard subtab on which requests to delete, restrict, or export end-user data can       be audited. Can be one of 'default or inherit', 'hide' or 'show'.
	FirewallInfoSubtab                 string `json:"firewallInfoSubtab,omitempty"`                 //       The 'Help -> Firewall info' subtab where necessary upstream firewall rules for communication to the Cisco Meraki cloud are       listed. Can be one of 'default or inherit', 'hide' or 'show'.
	GetHelpSubtab                      string `json:"getHelpSubtab,omitempty"`                      //       The 'Help -> Get Help' subtab on which Cisco Meraki KB, Product Manuals, and Support/Case Information are displayed. Note       that if this subtab is hidden, branding customizations for the KB on 'Get help', Cisco Meraki product documentation,       and support contact info will not be visible. Can be one of 'default or inherit', 'hide' or 'show'.
	GetHelpSubtabKnowledgeBaseSearch   string `json:"getHelpSubtabKnowledgeBaseSearch,omitempty"`   //       The KB search box which appears on the Help page. Can be one of 'default or inherit', 'hide', 'show', or a replacement custom HTML string.
	HardwareReplacementsSubtab         string `json:"hardwareReplacementsSubtab,omitempty"`         //       The 'Help -> Replacement info' subtab where important information regarding device replacements is detailed. Can be one of       'default or inherit', 'hide' or 'show'.
	HelpTab                            string `json:"helpTab,omitempty"`                            //       The Help tab, under which all support information resides. If this tab is hidden, no other 'Help' branding       customizations will be visible. Can be one of 'default or inherit', 'hide' or 'show'.
	HelpWidget                         string `json:"helpWidget,omitempty"`                         //       The 'Help Widget' is a support widget which provides access to live chat, documentation links, Sales contact info,       and other contact avenues to reach Meraki Support. Can be one of 'default or inherit', 'hide' or 'show'.
	NewFeaturesSubtab                  string `json:"newFeaturesSubtab,omitempty"`                  //       The 'Help -> New features' subtab where new Dashboard features are detailed. Can be one of 'default or inherit', 'hide' or 'show'.
	SmForums                           string `json:"smForums,omitempty"`                           //       The 'SM Forums' subtab which links to community-based support for Cisco Meraki Systems Manager. Only configurable for       organizations that contain Systems Manager networks. Can be one of 'default or inherit', 'hide' or 'show'.
	SupportContactInfo                 string `json:"supportContactInfo,omitempty"`                 //       The 'Contact Meraki Support' section of the 'Help -> Get Help' subtab. Can be one of 'default or inherit', 'hide', 'show', or a replacement custom HTML string.
	UniversalSearchKnowledgeBaseSearch string `json:"universalSearchKnowledgeBaseSearch,omitempty"` //       The universal search box always visible on Dashboard will, by default, present results from the Meraki KB. This configures       whether these Meraki KB results should be returned. Can be one of 'default or inherit', 'hide' or 'show'.
}
type ResponseOrganizationsUpdateOrganizationBrandingPolicy struct {
	AdminSettings *ResponseOrganizationsUpdateOrganizationBrandingPolicyAdminSettings `json:"adminSettings,omitempty"` // Settings for describing which kinds of admins this policy applies to.
	CustomLogo    *ResponseOrganizationsUpdateOrganizationBrandingPolicyCustomLogo    `json:"customLogo,omitempty"`    // Properties describing the custom logo attached to the branding policy.
	Enabled       *bool                                                               `json:"enabled,omitempty"`       // Boolean indicating whether this policy is enabled.
	HelpSettings  *ResponseOrganizationsUpdateOrganizationBrandingPolicyHelpSettings  `json:"helpSettings,omitempty"`  //       Settings for describing the modifications to various Help page features. Each property in this object accepts one of       'default or inherit' (do not modify functionality), 'hide' (remove the section from Dashboard), or 'show' (always show       the section on Dashboard). Some properties in this object also accept custom HTML used to replace the section on       Dashboard; see the documentation for each property to see the allowed values.
	Name          string                                                              `json:"name,omitempty"`          // Name of the Dashboard branding policy.
}
type ResponseOrganizationsUpdateOrganizationBrandingPolicyAdminSettings struct {
	AppliesTo string   `json:"appliesTo,omitempty"` // Which kinds of admins this policy applies to. Can be one of 'All organization admins', 'All enterprise admins', 'All network admins', 'All admins of networks...', 'All admins of networks tagged...', 'Specific admins...', 'All admins' or 'All SAML admins'.
	Values    []string `json:"values,omitempty"`    //       If 'appliesTo' is set to one of 'Specific admins...', 'All admins of networks...' or 'All admins of networks tagged...', then you must specify this 'values' property to provide the set of       entities to apply the branding policy to. For 'Specific admins...', specify an array of admin IDs. For 'All admins of       networks...', specify an array of network IDs and/or configuration template IDs. For 'All admins of networks tagged...',       specify an array of tag names.
}
type ResponseOrganizationsUpdateOrganizationBrandingPolicyCustomLogo struct {
	Enabled *bool                                                                 `json:"enabled,omitempty"` // Whether or not there is a custom logo enabled.
	Image   *ResponseOrganizationsUpdateOrganizationBrandingPolicyCustomLogoImage `json:"image,omitempty"`   // Properties of the image.
}
type ResponseOrganizationsUpdateOrganizationBrandingPolicyCustomLogoImage struct {
	Preview *ResponseOrganizationsUpdateOrganizationBrandingPolicyCustomLogoImagePreview `json:"preview,omitempty"` // Preview of the image
}
type ResponseOrganizationsUpdateOrganizationBrandingPolicyCustomLogoImagePreview struct {
	ExpiresAt string `json:"expiresAt,omitempty"` // Timestamp of the preview image
	URL       string `json:"url,omitempty"`       // Url of the preview image
}
type ResponseOrganizationsUpdateOrganizationBrandingPolicyHelpSettings struct {
	APIDocsSubtab                      string `json:"apiDocsSubtab,omitempty"`                      //       The 'Help -> API docs' subtab where a detailed description of the Dashboard API is listed. Can be one of       'default or inherit', 'hide' or 'show'.
	CasesSubtab                        string `json:"casesSubtab,omitempty"`                        //       The 'Help -> Cases' Dashboard subtab on which Cisco Meraki support cases for this organization can be managed. Can be one       of 'default or inherit', 'hide' or 'show'.
	CiscoMerakiProductDocumentation    string `json:"ciscoMerakiProductDocumentation,omitempty"`    //       The 'Product Manuals' section of the 'Help -> Get Help' subtab. Can be one of 'default or inherit', 'hide', 'show', or a replacement custom HTML string.
	CommunitySubtab                    string `json:"communitySubtab,omitempty"`                    //       The 'Help -> Community' subtab which provides a link to Meraki Community. Can be one of 'default or inherit', 'hide' or 'show'.
	DataProtectionRequestsSubtab       string `json:"dataProtectionRequestsSubtab,omitempty"`       //       The 'Help -> Data protection requests' Dashboard subtab on which requests to delete, restrict, or export end-user data can       be audited. Can be one of 'default or inherit', 'hide' or 'show'.
	FirewallInfoSubtab                 string `json:"firewallInfoSubtab,omitempty"`                 //       The 'Help -> Firewall info' subtab where necessary upstream firewall rules for communication to the Cisco Meraki cloud are       listed. Can be one of 'default or inherit', 'hide' or 'show'.
	GetHelpSubtab                      string `json:"getHelpSubtab,omitempty"`                      //       The 'Help -> Get Help' subtab on which Cisco Meraki KB, Product Manuals, and Support/Case Information are displayed. Note       that if this subtab is hidden, branding customizations for the KB on 'Get help', Cisco Meraki product documentation,       and support contact info will not be visible. Can be one of 'default or inherit', 'hide' or 'show'.
	GetHelpSubtabKnowledgeBaseSearch   string `json:"getHelpSubtabKnowledgeBaseSearch,omitempty"`   //       The KB search box which appears on the Help page. Can be one of 'default or inherit', 'hide', 'show', or a replacement custom HTML string.
	HardwareReplacementsSubtab         string `json:"hardwareReplacementsSubtab,omitempty"`         //       The 'Help -> Replacement info' subtab where important information regarding device replacements is detailed. Can be one of       'default or inherit', 'hide' or 'show'.
	HelpTab                            string `json:"helpTab,omitempty"`                            //       The Help tab, under which all support information resides. If this tab is hidden, no other 'Help' branding       customizations will be visible. Can be one of 'default or inherit', 'hide' or 'show'.
	HelpWidget                         string `json:"helpWidget,omitempty"`                         //       The 'Help Widget' is a support widget which provides access to live chat, documentation links, Sales contact info,       and other contact avenues to reach Meraki Support. Can be one of 'default or inherit', 'hide' or 'show'.
	NewFeaturesSubtab                  string `json:"newFeaturesSubtab,omitempty"`                  //       The 'Help -> New features' subtab where new Dashboard features are detailed. Can be one of 'default or inherit', 'hide' or 'show'.
	SmForums                           string `json:"smForums,omitempty"`                           //       The 'SM Forums' subtab which links to community-based support for Cisco Meraki Systems Manager. Only configurable for       organizations that contain Systems Manager networks. Can be one of 'default or inherit', 'hide' or 'show'.
	SupportContactInfo                 string `json:"supportContactInfo,omitempty"`                 //       The 'Contact Meraki Support' section of the 'Help -> Get Help' subtab. Can be one of 'default or inherit', 'hide', 'show', or a replacement custom HTML string.
	UniversalSearchKnowledgeBaseSearch string `json:"universalSearchKnowledgeBaseSearch,omitempty"` //       The universal search box always visible on Dashboard will, by default, present results from the Meraki KB. This configures       whether these Meraki KB results should be returned. Can be one of 'default or inherit', 'hide' or 'show'.
}
type ResponseOrganizationsClaimIntoOrganization interface{}
type ResponseOrganizationsGetOrganizationClientsBandwidthUsageHistory []ResponseItemOrganizationsGetOrganizationClientsBandwidthUsageHistory // Array of ResponseOrganizationsGetOrganizationClientsBandwidthUsageHistory
type ResponseItemOrganizationsGetOrganizationClientsBandwidthUsageHistory struct {
	Downstream *int   `json:"downstream,omitempty"` // Downloaded data, in mbps.
	Total      *int   `json:"total,omitempty"`      // Total bandwidth usage, in mbps.
	Ts         string `json:"ts,omitempty"`         // Timestamp for the bandwidth usage snapshot.
	Upstream   *int   `json:"upstream,omitempty"`   // Uploaded data, in mbps.
}
type ResponseOrganizationsGetOrganizationClientsOverview struct {
	Counts *ResponseOrganizationsGetOrganizationClientsOverviewCounts `json:"counts,omitempty"` // Client count information
	Usage  *ResponseOrganizationsGetOrganizationClientsOverviewUsage  `json:"usage,omitempty"`  // Usage information of all clients across organization
}
type ResponseOrganizationsGetOrganizationClientsOverviewCounts struct {
	Total *int `json:"total,omitempty"` // Total number of clients with data usage in organization
}
type ResponseOrganizationsGetOrganizationClientsOverviewUsage struct {
	Average *float64                                                         `json:"average,omitempty"` // Average data usage (in kb) of each client in organization
	Overall *ResponseOrganizationsGetOrganizationClientsOverviewUsageOverall `json:"overall,omitempty"` // Overall data usage of all clients across organization
}
type ResponseOrganizationsGetOrganizationClientsOverviewUsageOverall struct {
	Downstream *float64 `json:"downstream,omitempty"` // Downstream data usage (in kb) of all clients across organization
	Total      *float64 `json:"total,omitempty"`      // Total data usage (in kb) of all clients across organization
	Upstream   *float64 `json:"upstream,omitempty"`   // Upstream data usage (in kb) of all clients across organization
}
type ResponseOrganizationsGetOrganizationClientsSearch struct {
	ClientID     string                                                      `json:"clientId,omitempty"`     //
	Mac          string                                                      `json:"mac,omitempty"`          //
	Manufacturer string                                                      `json:"manufacturer,omitempty"` //
	Records      *[]ResponseOrganizationsGetOrganizationClientsSearchRecords `json:"records,omitempty"`      //
}
type ResponseOrganizationsGetOrganizationClientsSearchRecords struct {
	ClientVpnConnections *[]ResponseOrganizationsGetOrganizationClientsSearchRecordsClientVpnConnections `json:"clientVpnConnections,omitempty"` //
	Description          string                                                                          `json:"description,omitempty"`          //
	FirstSeen            *int                                                                            `json:"firstSeen,omitempty"`            //
	IP                   string                                                                          `json:"ip,omitempty"`                   //
	IP6                  string                                                                          `json:"ip6,omitempty"`                  //
	LastSeen             *int                                                                            `json:"lastSeen,omitempty"`             //
	Lldp                 []string                                                                        `json:"lldp,omitempty"`                 //
	Network              *ResponseOrganizationsGetOrganizationClientsSearchRecordsNetwork                `json:"network,omitempty"`              //
	Os                   string                                                                          `json:"os,omitempty"`                   //
	SmInstalled          *bool                                                                           `json:"smInstalled,omitempty"`          //
	SSID                 string                                                                          `json:"ssid,omitempty"`                 //
	Status               string                                                                          `json:"status,omitempty"`               //
	User                 string                                                                          `json:"user,omitempty"`                 //
	VLAN                 string                                                                          `json:"vlan,omitempty"`                 //
	WirelessCapabilities string                                                                          `json:"wirelessCapabilities,omitempty"` //
}
type ResponseOrganizationsGetOrganizationClientsSearchRecordsClientVpnConnections struct {
	ConnectedAt    *int   `json:"connectedAt,omitempty"`    //
	DisconnectedAt *int   `json:"disconnectedAt,omitempty"` //
	RemoteIP       string `json:"remoteIp,omitempty"`       //
}
type ResponseOrganizationsGetOrganizationClientsSearchRecordsNetwork struct {
	EnrollmentString        string   `json:"enrollmentString,omitempty"`        //
	ID                      string   `json:"id,omitempty"`                      //
	IsBoundToConfigTemplate *bool    `json:"isBoundToConfigTemplate,omitempty"` //
	Name                    string   `json:"name,omitempty"`                    //
	Notes                   string   `json:"notes,omitempty"`                   //
	OrganizationID          string   `json:"organizationId,omitempty"`          //
	ProductTypes            []string `json:"productTypes,omitempty"`            //
	Tags                    []string `json:"tags,omitempty"`                    //
	TimeZone                string   `json:"timeZone,omitempty"`                //
}
type ResponseOrganizationsCloneOrganization struct {
	API        *ResponseOrganizationsCloneOrganizationAPI        `json:"api,omitempty"`        // API related settings
	Cloud      *ResponseOrganizationsCloneOrganizationCloud      `json:"cloud,omitempty"`      // Data for this organization
	ID         string                                            `json:"id,omitempty"`         // Organization ID
	Licensing  *ResponseOrganizationsCloneOrganizationLicensing  `json:"licensing,omitempty"`  // Licensing related settings
	Management *ResponseOrganizationsCloneOrganizationManagement `json:"management,omitempty"` // Information about the organization's management system
	Name       string                                            `json:"name,omitempty"`       // Organization name
	URL        string                                            `json:"url,omitempty"`        // Organization URL
}
type ResponseOrganizationsCloneOrganizationAPI struct {
	Enabled *bool `json:"enabled,omitempty"` // Enable API access
}
type ResponseOrganizationsCloneOrganizationCloud struct {
	Region *ResponseOrganizationsCloneOrganizationCloudRegion `json:"region,omitempty"` // Region info
}
type ResponseOrganizationsCloneOrganizationCloudRegion struct {
	Name string `json:"name,omitempty"` // Name of region
}
type ResponseOrganizationsCloneOrganizationLicensing struct {
	Model string `json:"model,omitempty"` // Organization licensing model. Can be 'co-term', 'per-device', or 'subscription'.
}
type ResponseOrganizationsCloneOrganizationManagement struct {
	Details *[]ResponseOrganizationsCloneOrganizationManagementDetails `json:"details,omitempty"` // Details related to organization management, possibly empty
}
type ResponseOrganizationsCloneOrganizationManagementDetails struct {
	Name  string `json:"name,omitempty"`  // Name of management data
	Value string `json:"value,omitempty"` // Value of management data
}
type ResponseOrganizationsGetOrganizationConfigTemplates []ResponseItemOrganizationsGetOrganizationConfigTemplates // Array of ResponseOrganizationsGetOrganizationConfigTemplates
type ResponseItemOrganizationsGetOrganizationConfigTemplates struct {
	ID           string   `json:"id,omitempty"`           //
	Name         string   `json:"name,omitempty"`         //
	ProductTypes []string `json:"productTypes,omitempty"` //
	TimeZone     string   `json:"timeZone,omitempty"`     //
}
type ResponseOrganizationsCreateOrganizationConfigTemplate interface{}
type ResponseOrganizationsGetOrganizationConfigTemplate struct {
	ID           string   `json:"id,omitempty"`           //
	Name         string   `json:"name,omitempty"`         //
	ProductTypes []string `json:"productTypes,omitempty"` //
	TimeZone     string   `json:"timeZone,omitempty"`     //
}
type ResponseOrganizationsUpdateOrganizationConfigTemplate interface{}
type ResponseOrganizationsGetOrganizationConfigurationChanges []ResponseItemOrganizationsGetOrganizationConfigurationChanges // Array of ResponseOrganizationsGetOrganizationConfigurationChanges
type ResponseItemOrganizationsGetOrganizationConfigurationChanges struct {
	AdminEmail string `json:"adminEmail,omitempty"` //
	AdminID    string `json:"adminId,omitempty"`    //
	AdminName  string `json:"adminName,omitempty"`  //
	Label      string `json:"label,omitempty"`      //
	NewValue   string `json:"newValue,omitempty"`   //
	OldValue   string `json:"oldValue,omitempty"`   //
	Page       string `json:"page,omitempty"`       //
	Ts         string `json:"ts,omitempty"`         //
}
type ResponseOrganizationsGetOrganizationDevices []ResponseItemOrganizationsGetOrganizationDevices // Array of ResponseOrganizationsGetOrganizationDevices
type ResponseItemOrganizationsGetOrganizationDevices struct {
	Address     string   `json:"address,omitempty"`     // Physical address of the device
	Firmware    string   `json:"firmware,omitempty"`    // Firmware version of the device
	LanIP       string   `json:"lanIp,omitempty"`       // LAN IP address of the device
	Lat         *float64 `json:"lat,omitempty"`         // Latitude of the device
	Lng         *float64 `json:"lng,omitempty"`         // Longitude of the device
	Mac         string   `json:"mac,omitempty"`         // MAC address of the device
	Model       string   `json:"model,omitempty"`       // Model of the device
	Name        string   `json:"name,omitempty"`        // Name of the device
	NetworkID   string   `json:"networkId,omitempty"`   // ID of the network the device belongs to
	Notes       string   `json:"notes,omitempty"`       // Notes for the device, limited to 255 characters
	ProductType string   `json:"productType,omitempty"` // Product type of the device
	Serial      string   `json:"serial,omitempty"`      // Serial number of the device
	Tags        []string `json:"tags,omitempty"`        // List of tags assigned to the device
}
type ResponseOrganizationsGetOrganizationDevicesAvailabilities []ResponseItemOrganizationsGetOrganizationDevicesAvailabilities // Array of ResponseOrganizationsGetOrganizationDevicesAvailabilities
type ResponseItemOrganizationsGetOrganizationDevicesAvailabilities struct {
	Mac         string                                                                `json:"mac,omitempty"`         // The device MAC address.
	Name        string                                                                `json:"name,omitempty"`        // The device name.
	Network     *ResponseItemOrganizationsGetOrganizationDevicesAvailabilitiesNetwork `json:"network,omitempty"`     // Network info.
	ProductType string                                                                `json:"productType,omitempty"` // Device product type.
	Serial      string                                                                `json:"serial,omitempty"`      // The device serial number.
	Status      string                                                                `json:"status,omitempty"`      // Status of the device. Possible values are: online, alerting, offline, dormant.
	Tags        []string                                                              `json:"tags,omitempty"`        // List of custom tags for the device.
}
type ResponseItemOrganizationsGetOrganizationDevicesAvailabilitiesNetwork struct {
	ID string `json:"id,omitempty"` // ID for the network containing the device.
}
type ResponseOrganizationsGetOrganizationDevicesPowerModulesStatusesByDevice []ResponseItemOrganizationsGetOrganizationDevicesPowerModulesStatusesByDevice // Array of ResponseOrganizationsGetOrganizationDevicesPowerModulesStatusesByDevice
type ResponseItemOrganizationsGetOrganizationDevicesPowerModulesStatusesByDevice struct {
	Mac         string                                                                              `json:"mac,omitempty"`         // The device MAC address.
	Name        string                                                                              `json:"name,omitempty"`        // The device name.
	Network     *ResponseItemOrganizationsGetOrganizationDevicesPowerModulesStatusesByDeviceNetwork `json:"network,omitempty"`     // Network info.
	ProductType string                                                                              `json:"productType,omitempty"` // Device product type.
	Serial      string                                                                              `json:"serial,omitempty"`      // The device serial number.
	Slots       *[]ResponseItemOrganizationsGetOrganizationDevicesPowerModulesStatusesByDeviceSlots `json:"slots,omitempty"`       // Information for the device's AC power supplies.
	Tags        []string                                                                            `json:"tags,omitempty"`        // List of custom tags for the device.
}
type ResponseItemOrganizationsGetOrganizationDevicesPowerModulesStatusesByDeviceNetwork struct {
	ID string `json:"id,omitempty"` // ID for the network that the device is associated with.
}
type ResponseItemOrganizationsGetOrganizationDevicesPowerModulesStatusesByDeviceSlots struct {
	Model  string `json:"model,omitempty"`  // The power supply unit model.
	Number *int   `json:"number,omitempty"` // Which slot the AC power supply occupies. Possible values are: 0, 1, 2.
	Serial string `json:"serial,omitempty"` // The power supply unit serial number.
	Status string `json:"status,omitempty"` // Status of the power supply unit. Possible values are: connected, not connected, powering.
}
type ResponseOrganizationsGetOrganizationDevicesProvisioningStatuses []ResponseItemOrganizationsGetOrganizationDevicesProvisioningStatuses // Array of ResponseOrganizationsGetOrganizationDevicesProvisioningStatuses
type ResponseItemOrganizationsGetOrganizationDevicesProvisioningStatuses struct {
	Mac         string                                                                      `json:"mac,omitempty"`         // The device MAC address.
	Name        string                                                                      `json:"name,omitempty"`        // The device name.
	Network     *ResponseItemOrganizationsGetOrganizationDevicesProvisioningStatusesNetwork `json:"network,omitempty"`     // Network info.
	ProductType string                                                                      `json:"productType,omitempty"` // Device product type.
	Serial      string                                                                      `json:"serial,omitempty"`      // The device serial number.
	Status      string                                                                      `json:"status,omitempty"`      // The device provisioning status. Possible statuses: unprovisioned, incomplete, complete.
	Tags        []string                                                                    `json:"tags,omitempty"`        // List of custom tags for the device.
}
type ResponseItemOrganizationsGetOrganizationDevicesProvisioningStatusesNetwork struct {
	ID string `json:"id,omitempty"` // ID for the network containing the device.
}
type ResponseOrganizationsGetOrganizationDevicesStatuses struct {
	Components     *ResponseOrganizationsGetOrganizationDevicesStatusesComponents `json:"components,omitempty"`     // Components
	Gateway        string                                                         `json:"gateway,omitempty"`        // IP Gateway
	IPType         string                                                         `json:"ipType,omitempty"`         // IP Type
	LanIP          string                                                         `json:"lanIp,omitempty"`          // LAN IP Address
	LastReportedAt string                                                         `json:"lastReportedAt,omitempty"` // Device Last Reported Location
	Mac            string                                                         `json:"mac,omitempty"`            // MAC Address
	Model          string                                                         `json:"model,omitempty"`          // Model
	Name           string                                                         `json:"name,omitempty"`           // Device Name
	NetworkID      string                                                         `json:"networkId,omitempty"`      // Network ID
	PrimaryDNS     string                                                         `json:"primaryDns,omitempty"`     // Primary DNS
	ProductType    string                                                         `json:"productType,omitempty"`    // Product Type
	PublicIP       string                                                         `json:"publicIp,omitempty"`       // Public IP Address
	SecondaryDNS   string                                                         `json:"secondaryDns,omitempty"`   // Secondary DNS
	Serial         string                                                         `json:"serial,omitempty"`         // Device Serial Number
	Status         string                                                         `json:"status,omitempty"`         // Device Status
	Tags           []string                                                       `json:"tags,omitempty"`           // Tags
}
type ResponseOrganizationsGetOrganizationDevicesStatusesComponents struct {
	PowerSupplies []string `json:"powerSupplies,omitempty"` // Power Supplies
}
type ResponseOrganizationsGetOrganizationDevicesStatusesOverview struct {
	Counts *ResponseOrganizationsGetOrganizationDevicesStatusesOverviewCounts `json:"counts,omitempty"` // counts
}
type ResponseOrganizationsGetOrganizationDevicesStatusesOverviewCounts struct {
	ByStatus *ResponseOrganizationsGetOrganizationDevicesStatusesOverviewCountsByStatus `json:"byStatus,omitempty"` // byStatus
}
type ResponseOrganizationsGetOrganizationDevicesStatusesOverviewCountsByStatus struct {
	Alerting *int `json:"alerting,omitempty"` // alerting count
	Dormant  *int `json:"dormant,omitempty"`  // dormant count
	Offline  *int `json:"offline,omitempty"`  // offline count
	Online   *int `json:"online,omitempty"`   // online count
}
type ResponseOrganizationsGetOrganizationDevicesUplinksAddressesByDevice []ResponseItemOrganizationsGetOrganizationDevicesUplinksAddressesByDevice // Array of ResponseOrganizationsGetOrganizationDevicesUplinksAddressesByDevice
type ResponseItemOrganizationsGetOrganizationDevicesUplinksAddressesByDevice struct {
	Mac         string                                                                            `json:"mac,omitempty"`         // The device MAC address.
	Name        string                                                                            `json:"name,omitempty"`        // The device name.
	Network     *ResponseItemOrganizationsGetOrganizationDevicesUplinksAddressesByDeviceNetwork   `json:"network,omitempty"`     // Network info.
	ProductType string                                                                            `json:"productType,omitempty"` // Device product type.
	Serial      string                                                                            `json:"serial,omitempty"`      // The device serial number.
	Tags        []string                                                                          `json:"tags,omitempty"`        // List of custom tags for the device.
	Uplinks     *[]ResponseItemOrganizationsGetOrganizationDevicesUplinksAddressesByDeviceUplinks `json:"uplinks,omitempty"`     // List of device uplink addresses information.
}
type ResponseItemOrganizationsGetOrganizationDevicesUplinksAddressesByDeviceNetwork struct {
	ID string `json:"id,omitempty"` // ID for the network containing the device.
}
type ResponseItemOrganizationsGetOrganizationDevicesUplinksAddressesByDeviceUplinks struct {
	Addresses *[]ResponseItemOrganizationsGetOrganizationDevicesUplinksAddressesByDeviceUplinksAddresses `json:"addresses,omitempty"` // Available addresses for the interface.
	Interface string                                                                                     `json:"interface,omitempty"` // Interface for the device uplink. Available options are: cellular, man1, man2, wan1, wan2
}
type ResponseItemOrganizationsGetOrganizationDevicesUplinksAddressesByDeviceUplinksAddresses struct {
	Address        string                                                                                         `json:"address,omitempty"`        // Device uplink address.
	AssignmentMode string                                                                                         `json:"assignmentMode,omitempty"` // Indicates how the device uplink address is assigned. Available options are: static, dynamic.
	Gateway        string                                                                                         `json:"gateway,omitempty"`        // Device uplink gateway address.
	Protocol       string                                                                                         `json:"protocol,omitempty"`       // Type of address for the device uplink. Available options are: ipv4, ipv6.
	Public         *ResponseItemOrganizationsGetOrganizationDevicesUplinksAddressesByDeviceUplinksAddressesPublic `json:"public,omitempty"`         // Public interface information.
}
type ResponseItemOrganizationsGetOrganizationDevicesUplinksAddressesByDeviceUplinksAddressesPublic struct {
	Address string `json:"address,omitempty"` // The device uplink public IP address.
}
type ResponseOrganizationsGetOrganizationDevicesUplinksLossAndLatency []ResponseItemOrganizationsGetOrganizationDevicesUplinksLossAndLatency // Array of ResponseOrganizationsGetOrganizationDevicesUplinksLossAndLatency
type ResponseItemOrganizationsGetOrganizationDevicesUplinksLossAndLatency struct {
	IP         string                                                                            `json:"ip,omitempty"`         // IP address of uplink
	NetworkID  string                                                                            `json:"networkId,omitempty"`  // Network ID
	Serial     string                                                                            `json:"serial,omitempty"`     // Serial of MX device
	TimeSeries *[]ResponseItemOrganizationsGetOrganizationDevicesUplinksLossAndLatencyTimeSeries `json:"timeSeries,omitempty"` // Loss and latency timeseries data
	Uplink     string                                                                            `json:"uplink,omitempty"`     // Uplink interface (wan1, wan2, or cellular)
}
type ResponseItemOrganizationsGetOrganizationDevicesUplinksLossAndLatencyTimeSeries struct {
	LatencyMs   *float64 `json:"latencyMs,omitempty"`   // Latency in milliseconds
	LossPercent *float64 `json:"lossPercent,omitempty"` // Loss percentage
	Ts          string   `json:"ts,omitempty"`          // Timestamp for this data point
}
type ResponseOrganizationsGetOrganizationEarlyAccessFeatures []ResponseItemOrganizationsGetOrganizationEarlyAccessFeatures // Array of ResponseOrganizationsGetOrganizationEarlyAccessFeatures
type ResponseItemOrganizationsGetOrganizationEarlyAccessFeatures struct {
	Descriptions      *ResponseItemOrganizationsGetOrganizationEarlyAccessFeaturesDescriptions `json:"descriptions,omitempty"`      // Descriptions of the early access feature
	DocumentationLink string                                                                   `json:"documentationLink,omitempty"` // Link to the documentation of this early access feature
	IsOrgScopedOnly   *bool                                                                    `json:"isOrgScopedOnly,omitempty"`   // If this early access feature can only be opted in for the entire organization
	Name              string                                                                   `json:"name,omitempty"`              // Name of the early access feature
	ShortName         string                                                                   `json:"shortName,omitempty"`         // Short name of the early access feature
	SupportLink       string                                                                   `json:"supportLink,omitempty"`       // Link to get support for this early access feature
	Topic             string                                                                   `json:"topic,omitempty"`             // Topic of the early access feature
}
type ResponseItemOrganizationsGetOrganizationEarlyAccessFeaturesDescriptions struct {
	Long  string `json:"long,omitempty"`  // Long description
	Short string `json:"short,omitempty"` // Short description
}
type ResponseOrganizationsGetOrganizationEarlyAccessFeaturesOptIns []ResponseItemOrganizationsGetOrganizationEarlyAccessFeaturesOptIns // Array of ResponseOrganizationsGetOrganizationEarlyAccessFeaturesOptIns
type ResponseItemOrganizationsGetOrganizationEarlyAccessFeaturesOptIns struct {
	CreatedAt            string                                                                                   `json:"createdAt,omitempty"`            //
	ID                   string                                                                                   `json:"id,omitempty"`                   //
	LimitScopeToNetworks *[]ResponseItemOrganizationsGetOrganizationEarlyAccessFeaturesOptInsLimitScopeToNetworks `json:"limitScopeToNetworks,omitempty"` //
	ShortName            string                                                                                   `json:"shortName,omitempty"`            //
}
type ResponseItemOrganizationsGetOrganizationEarlyAccessFeaturesOptInsLimitScopeToNetworks struct {
	ID   string `json:"id,omitempty"`   //
	Name string `json:"name,omitempty"` //
}
type ResponseOrganizationsCreateOrganizationEarlyAccessFeaturesOptIn struct {
	CreatedAt            string                                                                                 `json:"createdAt,omitempty"`            //
	ID                   string                                                                                 `json:"id,omitempty"`                   //
	LimitScopeToNetworks *[]ResponseOrganizationsCreateOrganizationEarlyAccessFeaturesOptInLimitScopeToNetworks `json:"limitScopeToNetworks,omitempty"` //
	ShortName            string                                                                                 `json:"shortName,omitempty"`            //
}
type ResponseOrganizationsCreateOrganizationEarlyAccessFeaturesOptInLimitScopeToNetworks struct {
	ID   string `json:"id,omitempty"`   //
	Name string `json:"name,omitempty"` //
}
type ResponseOrganizationsGetOrganizationEarlyAccessFeaturesOptIn struct {
	CreatedAt            string                                                                              `json:"createdAt,omitempty"`            //
	ID                   string                                                                              `json:"id,omitempty"`                   //
	LimitScopeToNetworks *[]ResponseOrganizationsGetOrganizationEarlyAccessFeaturesOptInLimitScopeToNetworks `json:"limitScopeToNetworks,omitempty"` //
	ShortName            string                                                                              `json:"shortName,omitempty"`            //
}
type ResponseOrganizationsGetOrganizationEarlyAccessFeaturesOptInLimitScopeToNetworks struct {
	ID   string `json:"id,omitempty"`   //
	Name string `json:"name,omitempty"` //
}
type ResponseOrganizationsUpdateOrganizationEarlyAccessFeaturesOptIn interface{}
type ResponseOrganizationsGetOrganizationFirmwareUpgrades []ResponseItemOrganizationsGetOrganizationFirmwareUpgrades // Array of ResponseOrganizationsGetOrganizationFirmwareUpgrades
type ResponseItemOrganizationsGetOrganizationFirmwareUpgrades struct {
	CompletedAt    string                                                               `json:"completedAt,omitempty"`    // Timestamp when upgrade completed. Null if status pending.
	FromVersion    *ResponseItemOrganizationsGetOrganizationFirmwareUpgradesFromVersion `json:"fromVersion,omitempty"`    // ID of the upgrade's starting version
	Network        *ResponseItemOrganizationsGetOrganizationFirmwareUpgradesNetwork     `json:"network,omitempty"`        // Network of the upgrade
	ProductType    string                                                               `json:"productType,omitempty"`    // product upgraded [wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor]
	Status         string                                                               `json:"status,omitempty"`         // Status of upgrade event: [Cancelled, Completed]
	Time           string                                                               `json:"time,omitempty"`           // Scheduled start time
	ToVersion      *ResponseItemOrganizationsGetOrganizationFirmwareUpgradesToVersion   `json:"toVersion,omitempty"`      // ID of the upgrade's target version
	UpgradeBatchID string                                                               `json:"upgradeBatchId,omitempty"` // The upgrade batch
	UpgradeID      string                                                               `json:"upgradeId,omitempty"`      // The upgrade
}
type ResponseItemOrganizationsGetOrganizationFirmwareUpgradesFromVersion struct {
	ID          string `json:"id,omitempty"`          // Firmware version ID
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseItemOrganizationsGetOrganizationFirmwareUpgradesNetwork struct {
	ID   string `json:"id,omitempty"`   // ID of network
	Name string `json:"name,omitempty"` // The network
}
type ResponseItemOrganizationsGetOrganizationFirmwareUpgradesToVersion struct {
	ID          string `json:"id,omitempty"`          // Firmware version ID
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseOrganizationsGetOrganizationFirmwareUpgradesByDevice []ResponseItemOrganizationsGetOrganizationFirmwareUpgradesByDevice // Array of ResponseOrganizationsGetOrganizationFirmwareUpgradesByDevice
type ResponseItemOrganizationsGetOrganizationFirmwareUpgradesByDevice struct {
	DeviceStatus string                                                                   `json:"deviceStatus,omitempty"` // Status of the device upgrade
	Name         string                                                                   `json:"name,omitempty"`         // Name assigned to the device
	Serial       string                                                                   `json:"serial,omitempty"`       // Serial of the device
	Upgrade      *ResponseItemOrganizationsGetOrganizationFirmwareUpgradesByDeviceUpgrade `json:"upgrade,omitempty"`      // The devices upgrade details and status
}
type ResponseItemOrganizationsGetOrganizationFirmwareUpgradesByDeviceUpgrade struct {
	FromVersion    *ResponseItemOrganizationsGetOrganizationFirmwareUpgradesByDeviceUpgradeFromVersion `json:"fromVersion,omitempty"`    // The initial version of the device
	ID             string                                                                              `json:"id,omitempty"`             // ID of the upgrade
	Staged         *ResponseItemOrganizationsGetOrganizationFirmwareUpgradesByDeviceUpgradestaged      `json:"staged,omitempty"`         // Staged upgrade
	Status         string                                                                              `json:"status,omitempty"`         // Status of the upgrade
	Time           string                                                                              `json:"time,omitempty"`           // Start time of the upgrade
	ToVersion      *ResponseItemOrganizationsGetOrganizationFirmwareUpgradesByDeviceUpgradeToVersion   `json:"toVersion,omitempty"`      // Version the device is upgrading to
	UpgradeBatchID string                                                                              `json:"upgradeBatchId,omitempty"` // ID of the upgrade batch
}
type ResponseItemOrganizationsGetOrganizationFirmwareUpgradesByDeviceUpgradeFromVersion struct {
	ID          string `json:"id,omitempty"`          // ID of the initial firmware version
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseItemOrganizationsGetOrganizationFirmwareUpgradesByDeviceUpgradestaged struct {
	Group *ResponseItemOrganizationsGetOrganizationFirmwareUpgradesByDeviceUpgradestagedGroup `json:"group,omitempty"` // The staged upgrade group
}
type ResponseItemOrganizationsGetOrganizationFirmwareUpgradesByDeviceUpgradestagedGroup struct {
	ID string `json:"id,omitempty"` // Id of the staged upgrade group
}
type ResponseItemOrganizationsGetOrganizationFirmwareUpgradesByDeviceUpgradeToVersion struct {
	ID          string `json:"id,omitempty"`          // ID of the initial firmware version
	ReleaseDate string `json:"releaseDate,omitempty"` // Release date of the firmware version
	ReleaseType string `json:"releaseType,omitempty"` // Release type of the firmware version
	ShortName   string `json:"shortName,omitempty"`   // Firmware version short name
}
type ResponseOrganizationsClaimIntoOrganizationInventory interface{}
type ResponseOrganizationsGetOrganizationInventoryDevices []ResponseItemOrganizationsGetOrganizationInventoryDevices // Array of ResponseOrganizationsGetOrganizationInventoryDevices
type ResponseItemOrganizationsGetOrganizationInventoryDevices struct {
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
type ResponseOrganizationsGetOrganizationInventoryDevice struct {
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
type ResponseOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringExportEvent interface{}
type ResponseOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringImports []ResponseItemOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringImports // Array of ResponseOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringImports
type ResponseItemOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringImports struct {
	Device   *ResponseItemOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringImportsDevice `json:"device,omitempty"`   // Represents the details of an imported device.
	ImportID string                                                                                   `json:"importId,omitempty"` // Database ID for the new entity entry.
}
type ResponseItemOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringImportsDevice struct {
	Created *bool  `json:"created,omitempty"` // Whether or not the device was successfully created in dashboard.
	Status  string `json:"status,omitempty"`  // Represents the current state of importing the device.
	URL     string `json:"url,omitempty"`     // The url to the device details page within dashboard.
}
type ResponseOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringImport []ResponseItemOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringImport // Array of ResponseOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringImport
type ResponseItemOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringImport struct {
	ImportID string `json:"importId,omitempty"` // Unique id associated with the import of the device
	Message  string `json:"message,omitempty"`  // Response method
	Status   string `json:"status,omitempty"`   // Cloud monitor import status
}
type ResponseOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringNetworks []ResponseItemOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringNetworks // Array of ResponseOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringNetworks
type ResponseItemOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringNetworks struct {
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
type ResponseOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepare []ResponseItemOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepare // Array of ResponseOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepare
type ResponseItemOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepare struct {
	ConfigParams *ResponseItemOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareConfigParams `json:"configParams,omitempty"` // Params used in order to connect to the device
	DeviceID     string                                                                                            `json:"deviceId,omitempty"`     // Import ID from the Import operation
	Message      string                                                                                            `json:"message,omitempty"`      // Message related to whether or not the device was found and can be imported.
	Status       string                                                                                            `json:"status,omitempty"`       // The import status of the device
	Udi          string                                                                                            `json:"udi,omitempty"`          // Device UDI certificate
}
type ResponseItemOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareConfigParams struct {
	CloudStaticIP string                                                                                                  `json:"cloudStaticIp,omitempty"` // Static IP Address used to connect to the device
	Tunnel        *ResponseItemOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareConfigParamsTunnel `json:"tunnel,omitempty"`        // Configuration options used to connect to the device
	User          *ResponseItemOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareConfigParamsUser   `json:"user,omitempty"`          // User credentials used to connect to the device
}
type ResponseItemOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareConfigParamsTunnel struct {
	Host            string                                                                                                                 `json:"host,omitempty"`            // SSH tunnel URL used to connect to the device
	Mode            string                                                                                                                 `json:"mode,omitempty"`            //
	Name            string                                                                                                                 `json:"name,omitempty"`            // The name of the tunnel we are attempting to connect to
	Port            string                                                                                                                 `json:"port,omitempty"`            // The port used for the ssh tunnel.
	RootCertificate *ResponseItemOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareConfigParamsTunnelRootCertificate `json:"rootCertificate,omitempty"` // Root certificate information
}
type ResponseItemOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareConfigParamsTunnelRootCertificate struct {
	Content string `json:"content,omitempty"` // Public certificate value
	Name    string `json:"name,omitempty"`    // The name of the server protected by the certificate
}
type ResponseItemOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareConfigParamsUser struct {
	PublicKey string                                                                                                      `json:"publicKey,omitempty"` // The public key for the registered user
	Secret    *ResponseItemOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecret `json:"secret,omitempty"`    // Stores the user secret hash
	Username  string                                                                                                      `json:"username,omitempty"`  // The username added to Catalyst device
}
type ResponseItemOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecret struct {
	Hash string `json:"hash,omitempty"` // The hashed secret
}
type ResponseOrganizationsReleaseFromOrganizationInventory interface{}
type ResponseOrganizationsGetOrganizationLicenses []ResponseItemOrganizationsGetOrganizationLicenses // Array of ResponseOrganizationsGetOrganizationLicenses
type ResponseItemOrganizationsGetOrganizationLicenses struct {
	ActivationDate            string                                                                       `json:"activationDate,omitempty"`            // The date the license started burning
	ClaimDate                 string                                                                       `json:"claimDate,omitempty"`                 // The date the license was claimed into the organization
	DeviceSerial              string                                                                       `json:"deviceSerial,omitempty"`              // Serial number of the device the license is assigned to
	DurationInDays            *int                                                                         `json:"durationInDays,omitempty"`            // The duration of the individual license
	ExpirationDate            string                                                                       `json:"expirationDate,omitempty"`            // The date the license will expire
	HeadLicenseID             string                                                                       `json:"headLicenseId,omitempty"`             // The id of the head license this license is queued behind. If there is no head license, it returns nil.
	ID                        string                                                                       `json:"id,omitempty"`                        // License ID
	LicenseKey                string                                                                       `json:"licenseKey,omitempty"`                // License key
	LicenseType               string                                                                       `json:"licenseType,omitempty"`               // License type
	NetworkID                 string                                                                       `json:"networkId,omitempty"`                 // ID of the network the license is assigned to
	OrderNumber               string                                                                       `json:"orderNumber,omitempty"`               // Order number
	PermanentlyQueuedLicenses *[]ResponseItemOrganizationsGetOrganizationLicensesPermanentlyQueuedLicenses `json:"permanentlyQueuedLicenses,omitempty"` // DEPRECATED List of permanently queued licenses attached to the license. Instead, use /organizations/{organizationId}/licenses?deviceSerial= to retrieved queued licenses for a given device.
	SeatCount                 *int                                                                         `json:"seatCount,omitempty"`                 // The number of seats of the license. Only applicable to SM licenses.
	State                     string                                                                       `json:"state,omitempty"`                     // The state of the license. All queued licenses have a status of *recentlyQueued*.
	TotalDurationInDays       *int                                                                         `json:"totalDurationInDays,omitempty"`       // The duration of the license plus all permanently queued licenses associated with it
}
type ResponseItemOrganizationsGetOrganizationLicensesPermanentlyQueuedLicenses struct {
	DurationInDays *int   `json:"durationInDays,omitempty"` // The duration of the individual license
	ID             string `json:"id,omitempty"`             // Permanently queued license ID
	LicenseKey     string `json:"licenseKey,omitempty"`     // License key
	LicenseType    string `json:"licenseType,omitempty"`    // License type
	OrderNumber    string `json:"orderNumber,omitempty"`    // Order number
}
type ResponseOrganizationsAssignOrganizationLicensesSeats struct {
	ResultingLicenses *[]ResponseOrganizationsAssignOrganizationLicensesSeatsResultingLicenses `json:"resultingLicenses,omitempty"` // Resulting licenses from the move
}
type ResponseOrganizationsAssignOrganizationLicensesSeatsResultingLicenses struct {
	ActivationDate            string                                                                                            `json:"activationDate,omitempty"`            // The date the license started burning
	ClaimDate                 string                                                                                            `json:"claimDate,omitempty"`                 // The date the license was claimed into the organization
	DeviceSerial              string                                                                                            `json:"deviceSerial,omitempty"`              // Serial number of the device the license is assigned to
	DurationInDays            *int                                                                                              `json:"durationInDays,omitempty"`            // The duration of the individual license
	ExpirationDate            string                                                                                            `json:"expirationDate,omitempty"`            // The date the license will expire
	HeadLicenseID             string                                                                                            `json:"headLicenseId,omitempty"`             // The id of the head license this license is queued behind. If there is no head license, it returns nil.
	ID                        string                                                                                            `json:"id,omitempty"`                        // License ID
	LicenseKey                string                                                                                            `json:"licenseKey,omitempty"`                // License key
	LicenseType               string                                                                                            `json:"licenseType,omitempty"`               // License type
	NetworkID                 string                                                                                            `json:"networkId,omitempty"`                 // ID of the network the license is assigned to
	OrderNumber               string                                                                                            `json:"orderNumber,omitempty"`               // Order number
	PermanentlyQueuedLicenses *[]ResponseOrganizationsAssignOrganizationLicensesSeatsResultingLicensesPermanentlyQueuedLicenses `json:"permanentlyQueuedLicenses,omitempty"` // DEPRECATED List of permanently queued licenses attached to the license. Instead, use /organizations/{organizationId}/licenses?deviceSerial= to retrieved queued licenses for a given device.
	SeatCount                 *int                                                                                              `json:"seatCount,omitempty"`                 // The number of seats of the license. Only applicable to SM licenses.
	State                     string                                                                                            `json:"state,omitempty"`                     // The state of the license. All queued licenses have a status of *recentlyQueued*.
	TotalDurationInDays       *int                                                                                              `json:"totalDurationInDays,omitempty"`       // The duration of the license plus all permanently queued licenses associated with it
}
type ResponseOrganizationsAssignOrganizationLicensesSeatsResultingLicensesPermanentlyQueuedLicenses struct {
	DurationInDays *int   `json:"durationInDays,omitempty"` // The duration of the individual license
	ID             string `json:"id,omitempty"`             // Permanently queued license ID
	LicenseKey     string `json:"licenseKey,omitempty"`     // License key
	LicenseType    string `json:"licenseType,omitempty"`    // License type
	OrderNumber    string `json:"orderNumber,omitempty"`    // Order number
}
type ResponseOrganizationsMoveOrganizationLicenses struct {
	DestOrganizationID string   `json:"destOrganizationId,omitempty"` // The ID of the organization to move the licenses to
	LicenseIDs         []string `json:"licenseIds,omitempty"`         // A list of IDs of licenses to move to the new organization
}
type ResponseOrganizationsMoveOrganizationLicensesSeats struct {
	DestOrganizationID string `json:"destOrganizationId,omitempty"` // The ID of the organization to move the SM seats to
	LicenseID          string `json:"licenseId,omitempty"`          // The ID of the SM license to move the seats from
	SeatCount          *int   `json:"seatCount,omitempty"`          // The number of seats to move to the new organization. Must be less than or equal to the total number of seats of the license
}
type ResponseOrganizationsGetOrganizationLicensesOverview struct {
	ExpirationDate       string                                                                    `json:"expirationDate,omitempty"`       //
	LicensedDeviceCounts *ResponseOrganizationsGetOrganizationLicensesOverviewLicensedDeviceCounts `json:"licensedDeviceCounts,omitempty"` //
	Status               string                                                                    `json:"status,omitempty"`               //
}
type ResponseOrganizationsGetOrganizationLicensesOverviewLicensedDeviceCounts struct {
	MS *int `json:"MS,omitempty"` //
}
type ResponseOrganizationsRenewOrganizationLicensesSeats struct {
	ResultingLicenses *[]ResponseOrganizationsRenewOrganizationLicensesSeatsResultingLicenses `json:"resultingLicenses,omitempty"` // Resulting licenses from the move
}
type ResponseOrganizationsRenewOrganizationLicensesSeatsResultingLicenses struct {
	ActivationDate            string                                                                                           `json:"activationDate,omitempty"`            // The date the license started burning
	ClaimDate                 string                                                                                           `json:"claimDate,omitempty"`                 // The date the license was claimed into the organization
	DeviceSerial              string                                                                                           `json:"deviceSerial,omitempty"`              // Serial number of the device the license is assigned to
	DurationInDays            *int                                                                                             `json:"durationInDays,omitempty"`            // The duration of the individual license
	ExpirationDate            string                                                                                           `json:"expirationDate,omitempty"`            // The date the license will expire
	HeadLicenseID             string                                                                                           `json:"headLicenseId,omitempty"`             // The id of the head license this license is queued behind. If there is no head license, it returns nil.
	ID                        string                                                                                           `json:"id,omitempty"`                        // License ID
	LicenseKey                string                                                                                           `json:"licenseKey,omitempty"`                // License key
	LicenseType               string                                                                                           `json:"licenseType,omitempty"`               // License type
	NetworkID                 string                                                                                           `json:"networkId,omitempty"`                 // ID of the network the license is assigned to
	OrderNumber               string                                                                                           `json:"orderNumber,omitempty"`               // Order number
	PermanentlyQueuedLicenses *[]ResponseOrganizationsRenewOrganizationLicensesSeatsResultingLicensesPermanentlyQueuedLicenses `json:"permanentlyQueuedLicenses,omitempty"` // DEPRECATED List of permanently queued licenses attached to the license. Instead, use /organizations/{organizationId}/licenses?deviceSerial= to retrieved queued licenses for a given device.
	SeatCount                 *int                                                                                             `json:"seatCount,omitempty"`                 // The number of seats of the license. Only applicable to SM licenses.
	State                     string                                                                                           `json:"state,omitempty"`                     // The state of the license. All queued licenses have a status of *recentlyQueued*.
	TotalDurationInDays       *int                                                                                             `json:"totalDurationInDays,omitempty"`       // The duration of the license plus all permanently queued licenses associated with it
}
type ResponseOrganizationsRenewOrganizationLicensesSeatsResultingLicensesPermanentlyQueuedLicenses struct {
	DurationInDays *int   `json:"durationInDays,omitempty"` // The duration of the individual license
	ID             string `json:"id,omitempty"`             // Permanently queued license ID
	LicenseKey     string `json:"licenseKey,omitempty"`     // License key
	LicenseType    string `json:"licenseType,omitempty"`    // License type
	OrderNumber    string `json:"orderNumber,omitempty"`    // Order number
}
type ResponseOrganizationsGetOrganizationLicense struct {
	ActivationDate            string                                                                  `json:"activationDate,omitempty"`            // The date the license started burning
	ClaimDate                 string                                                                  `json:"claimDate,omitempty"`                 // The date the license was claimed into the organization
	DeviceSerial              string                                                                  `json:"deviceSerial,omitempty"`              // Serial number of the device the license is assigned to
	DurationInDays            *int                                                                    `json:"durationInDays,omitempty"`            // The duration of the individual license
	ExpirationDate            string                                                                  `json:"expirationDate,omitempty"`            // The date the license will expire
	HeadLicenseID             string                                                                  `json:"headLicenseId,omitempty"`             // The id of the head license this license is queued behind. If there is no head license, it returns nil.
	ID                        string                                                                  `json:"id,omitempty"`                        // License ID
	LicenseKey                string                                                                  `json:"licenseKey,omitempty"`                // License key
	LicenseType               string                                                                  `json:"licenseType,omitempty"`               // License type
	NetworkID                 string                                                                  `json:"networkId,omitempty"`                 // ID of the network the license is assigned to
	OrderNumber               string                                                                  `json:"orderNumber,omitempty"`               // Order number
	PermanentlyQueuedLicenses *[]ResponseOrganizationsGetOrganizationLicensePermanentlyQueuedLicenses `json:"permanentlyQueuedLicenses,omitempty"` // DEPRECATED List of permanently queued licenses attached to the license. Instead, use /organizations/{organizationId}/licenses?deviceSerial= to retrieved queued licenses for a given device.
	SeatCount                 *int                                                                    `json:"seatCount,omitempty"`                 // The number of seats of the license. Only applicable to SM licenses.
	State                     string                                                                  `json:"state,omitempty"`                     // The state of the license. All queued licenses have a status of *recentlyQueued*.
	TotalDurationInDays       *int                                                                    `json:"totalDurationInDays,omitempty"`       // The duration of the license plus all permanently queued licenses associated with it
}
type ResponseOrganizationsGetOrganizationLicensePermanentlyQueuedLicenses struct {
	DurationInDays *int   `json:"durationInDays,omitempty"` // The duration of the individual license
	ID             string `json:"id,omitempty"`             // Permanently queued license ID
	LicenseKey     string `json:"licenseKey,omitempty"`     // License key
	LicenseType    string `json:"licenseType,omitempty"`    // License type
	OrderNumber    string `json:"orderNumber,omitempty"`    // Order number
}
type ResponseOrganizationsUpdateOrganizationLicense struct {
	ActivationDate            string                                                                     `json:"activationDate,omitempty"`            // The date the license started burning
	ClaimDate                 string                                                                     `json:"claimDate,omitempty"`                 // The date the license was claimed into the organization
	DeviceSerial              string                                                                     `json:"deviceSerial,omitempty"`              // Serial number of the device the license is assigned to
	DurationInDays            *int                                                                       `json:"durationInDays,omitempty"`            // The duration of the individual license
	ExpirationDate            string                                                                     `json:"expirationDate,omitempty"`            // The date the license will expire
	HeadLicenseID             string                                                                     `json:"headLicenseId,omitempty"`             // The id of the head license this license is queued behind. If there is no head license, it returns nil.
	ID                        string                                                                     `json:"id,omitempty"`                        // License ID
	LicenseKey                string                                                                     `json:"licenseKey,omitempty"`                // License key
	LicenseType               string                                                                     `json:"licenseType,omitempty"`               // License type
	NetworkID                 string                                                                     `json:"networkId,omitempty"`                 // ID of the network the license is assigned to
	OrderNumber               string                                                                     `json:"orderNumber,omitempty"`               // Order number
	PermanentlyQueuedLicenses *[]ResponseOrganizationsUpdateOrganizationLicensePermanentlyQueuedLicenses `json:"permanentlyQueuedLicenses,omitempty"` // DEPRECATED List of permanently queued licenses attached to the license. Instead, use /organizations/{organizationId}/licenses?deviceSerial= to retrieved queued licenses for a given device.
	SeatCount                 *int                                                                       `json:"seatCount,omitempty"`                 // The number of seats of the license. Only applicable to SM licenses.
	State                     string                                                                     `json:"state,omitempty"`                     // The state of the license. All queued licenses have a status of *recentlyQueued*.
	TotalDurationInDays       *int                                                                       `json:"totalDurationInDays,omitempty"`       // The duration of the license plus all permanently queued licenses associated with it
}
type ResponseOrganizationsUpdateOrganizationLicensePermanentlyQueuedLicenses struct {
	DurationInDays *int   `json:"durationInDays,omitempty"` // The duration of the individual license
	ID             string `json:"id,omitempty"`             // Permanently queued license ID
	LicenseKey     string `json:"licenseKey,omitempty"`     // License key
	LicenseType    string `json:"licenseType,omitempty"`    // License type
	OrderNumber    string `json:"orderNumber,omitempty"`    // Order number
}
type ResponseOrganizationsGetOrganizationLoginSecurity struct {
	AccountLockoutAttempts    *int                                                                `json:"accountLockoutAttempts,omitempty"`    // Number of consecutive failed login attempts after which users' accounts will be locked.
	APIAuthentication         *ResponseOrganizationsGetOrganizationLoginSecurityAPIAuthentication `json:"apiAuthentication,omitempty"`         // Details for indicating whether organization will restrict access to API (but not Dashboard) to certain IP addresses.
	EnforceAccountLockout     *bool                                                               `json:"enforceAccountLockout,omitempty"`     // Boolean indicating whether users' Dashboard accounts will be locked out after a specified number of consecutive failed login attempts.
	EnforceDifferentPasswords *bool                                                               `json:"enforceDifferentPasswords,omitempty"` // Boolean indicating whether users, when setting a new password, are forced to choose a new password that is different from any past passwords.
	EnforceIDleTimeout        *bool                                                               `json:"enforceIdleTimeout,omitempty"`        // Boolean indicating whether users will be logged out after being idle for the specified number of minutes.
	EnforceLoginIPRanges      *bool                                                               `json:"enforceLoginIpRanges,omitempty"`      // Boolean indicating whether organization will restrict access to Dashboard (including the API) from certain IP addresses.
	EnforcePasswordExpiration *bool                                                               `json:"enforcePasswordExpiration,omitempty"` // Boolean indicating whether users are forced to change their password every X number of days.
	EnforceStrongPasswords    *bool                                                               `json:"enforceStrongPasswords,omitempty"`    // Boolean indicating whether users will be forced to choose strong passwords for their accounts. Strong passwords are at least 8 characters that contain 3 of the following: number, uppercase letter, lowercase letter, and symbol
	EnforceTwoFactorAuth      *bool                                                               `json:"enforceTwoFactorAuth,omitempty"`      // Boolean indicating whether users in this organization will be required to use an extra verification code when logging in to Dashboard. This code will be sent to their mobile phone via SMS, or can be generated by the authenticator application.
	IDleTimeoutMinutes        *int                                                                `json:"idleTimeoutMinutes,omitempty"`        // Number of minutes users can remain idle before being logged out of their accounts.
	LoginIPRanges             []string                                                            `json:"loginIpRanges,omitempty"`             // List of acceptable IP ranges. Entries can be single IP addresses, IP address ranges, and CIDR subnets.
	NumDifferentPasswords     *int                                                                `json:"numDifferentPasswords,omitempty"`     // Number of recent passwords that new password must be distinct from.
	PasswordExpirationDays    *int                                                                `json:"passwordExpirationDays,omitempty"`    // Number of days after which users will be forced to change their password.
}
type ResponseOrganizationsGetOrganizationLoginSecurityAPIAuthentication struct {
	IPRestrictionsForKeys *ResponseOrganizationsGetOrganizationLoginSecurityAPIAuthenticationIPRestrictionsForKeys `json:"ipRestrictionsForKeys,omitempty"` // Details for API-only IP restrictions.
}
type ResponseOrganizationsGetOrganizationLoginSecurityAPIAuthenticationIPRestrictionsForKeys struct {
	Enabled *bool    `json:"enabled,omitempty"` // Boolean indicating whether the organization will restrict API key (not Dashboard GUI) usage to a specific list of IP addresses or CIDR ranges.
	Ranges  []string `json:"ranges,omitempty"`  // List of acceptable IP ranges. Entries can be single IP addresses, IP address ranges, and CIDR subnets.
}
type ResponseOrganizationsUpdateOrganizationLoginSecurity struct {
	AccountLockoutAttempts    *int                                                                   `json:"accountLockoutAttempts,omitempty"`    // Number of consecutive failed login attempts after which users' accounts will be locked.
	APIAuthentication         *ResponseOrganizationsUpdateOrganizationLoginSecurityAPIAuthentication `json:"apiAuthentication,omitempty"`         // Details for indicating whether organization will restrict access to API (but not Dashboard) to certain IP addresses.
	EnforceAccountLockout     *bool                                                                  `json:"enforceAccountLockout,omitempty"`     // Boolean indicating whether users' Dashboard accounts will be locked out after a specified number of consecutive failed login attempts.
	EnforceDifferentPasswords *bool                                                                  `json:"enforceDifferentPasswords,omitempty"` // Boolean indicating whether users, when setting a new password, are forced to choose a new password that is different from any past passwords.
	EnforceIDleTimeout        *bool                                                                  `json:"enforceIdleTimeout,omitempty"`        // Boolean indicating whether users will be logged out after being idle for the specified number of minutes.
	EnforceLoginIPRanges      *bool                                                                  `json:"enforceLoginIpRanges,omitempty"`      // Boolean indicating whether organization will restrict access to Dashboard (including the API) from certain IP addresses.
	EnforcePasswordExpiration *bool                                                                  `json:"enforcePasswordExpiration,omitempty"` // Boolean indicating whether users are forced to change their password every X number of days.
	EnforceStrongPasswords    *bool                                                                  `json:"enforceStrongPasswords,omitempty"`    // Boolean indicating whether users will be forced to choose strong passwords for their accounts. Strong passwords are at least 8 characters that contain 3 of the following: number, uppercase letter, lowercase letter, and symbol
	EnforceTwoFactorAuth      *bool                                                                  `json:"enforceTwoFactorAuth,omitempty"`      // Boolean indicating whether users in this organization will be required to use an extra verification code when logging in to Dashboard. This code will be sent to their mobile phone via SMS, or can be generated by the authenticator application.
	IDleTimeoutMinutes        *int                                                                   `json:"idleTimeoutMinutes,omitempty"`        // Number of minutes users can remain idle before being logged out of their accounts.
	LoginIPRanges             []string                                                               `json:"loginIpRanges,omitempty"`             // List of acceptable IP ranges. Entries can be single IP addresses, IP address ranges, and CIDR subnets.
	NumDifferentPasswords     *int                                                                   `json:"numDifferentPasswords,omitempty"`     // Number of recent passwords that new password must be distinct from.
	PasswordExpirationDays    *int                                                                   `json:"passwordExpirationDays,omitempty"`    // Number of days after which users will be forced to change their password.
}
type ResponseOrganizationsUpdateOrganizationLoginSecurityAPIAuthentication struct {
	IPRestrictionsForKeys *ResponseOrganizationsUpdateOrganizationLoginSecurityAPIAuthenticationIPRestrictionsForKeys `json:"ipRestrictionsForKeys,omitempty"` // Details for API-only IP restrictions.
}
type ResponseOrganizationsUpdateOrganizationLoginSecurityAPIAuthenticationIPRestrictionsForKeys struct {
	Enabled *bool    `json:"enabled,omitempty"` // Boolean indicating whether the organization will restrict API key (not Dashboard GUI) usage to a specific list of IP addresses or CIDR ranges.
	Ranges  []string `json:"ranges,omitempty"`  // List of acceptable IP ranges. Entries can be single IP addresses, IP address ranges, and CIDR subnets.
}
type ResponseOrganizationsGetOrganizationNetworks []ResponseItemOrganizationsGetOrganizationNetworks // Array of ResponseOrganizationsGetOrganizationNetworks
type ResponseItemOrganizationsGetOrganizationNetworks struct {
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
type ResponseOrganizationsCreateOrganizationNetwork struct {
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
type ResponseOrganizationsCombineOrganizationNetworks struct {
	ResultingNetwork *ResponseOrganizationsCombineOrganizationNetworksResultingNetwork `json:"resultingNetwork,omitempty"` // Network after the combination
}
type ResponseOrganizationsCombineOrganizationNetworksResultingNetwork struct {
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
type ResponseOrganizationsGetOrganizationOpenapiSpec struct {
	Info    *ResponseOrganizationsGetOrganizationOpenapiSpecInfo  `json:"info,omitempty"`    //
	Openapi string                                                `json:"openapi,omitempty"` //
	Paths   *ResponseOrganizationsGetOrganizationOpenapiSpecPaths `json:"paths,omitempty"`   //
}
type ResponseOrganizationsGetOrganizationOpenapiSpecInfo struct {
	Description string `json:"description,omitempty"` //
	Title       string `json:"title,omitempty"`       //
	Version     string `json:"version,omitempty"`     //
}
type ResponseOrganizationsGetOrganizationOpenapiSpecPaths struct {
	Organizations *ResponseOrganizationsGetOrganizationOpenapiSpecPathsOrganizations `json:"/organizations,omitempty"` //
}
type ResponseOrganizationsGetOrganizationOpenapiSpecPathsOrganizations struct {
	Get *ResponseOrganizationsGetOrganizationOpenapiSpecPathsOrganizationsGet `json:"get,omitempty"` //
}
type ResponseOrganizationsGetOrganizationOpenapiSpecPathsOrganizationsGet struct {
	Description string                                                                         `json:"description,omitempty"` //
	OperationID string                                                                         `json:"operationId,omitempty"` //
	Responses   *ResponseOrganizationsGetOrganizationOpenapiSpecPathsOrganizationsGetResponses `json:"responses,omitempty"`   //
}
type ResponseOrganizationsGetOrganizationOpenapiSpecPathsOrganizationsGetResponses struct {
	Status200 *ResponseOrganizationsGetOrganizationOpenapiSpecPathsOrganizationsGetResponses200 `json:"200,omitempty"` //
}
type ResponseOrganizationsGetOrganizationOpenapiSpecPathsOrganizationsGetResponses200 struct {
	Description string                                                                                    `json:"description,omitempty"` //
	Examples    *ResponseOrganizationsGetOrganizationOpenapiSpecPathsOrganizationsGetResponses200Examples `json:"examples,omitempty"`    //
}
type ResponseOrganizationsGetOrganizationOpenapiSpecPathsOrganizationsGetResponses200Examples struct {
	ApplicationJSON *[]ResponseOrganizationsGetOrganizationOpenapiSpecPathsOrganizationsGetResponses200ExamplesApplicationJSON `json:"application/json,omitempty"` //
}
type ResponseOrganizationsGetOrganizationOpenapiSpecPathsOrganizationsGetResponses200ExamplesApplicationJSON struct {
	ID   string `json:"id,omitempty"`   //
	Name string `json:"name,omitempty"` //
}
type ResponseOrganizationsGetOrganizationPolicyObjects []ResponseItemOrganizationsGetOrganizationPolicyObjects // Array of ResponseOrganizationsGetOrganizationPolicyObjects
type ResponseItemOrganizationsGetOrganizationPolicyObjects struct {
	Category   string   `json:"category,omitempty"`   //
	Cidr       string   `json:"cidr,omitempty"`       //
	CreatedAt  string   `json:"created_at,omitempty"` //
	GroupIDs   []string `json:"groupIds,omitempty"`   //
	ID         string   `json:"id,omitempty"`         //
	Name       string   `json:"name,omitempty"`       //
	NetworkIDs []string `json:"networkIds,omitempty"` //
	Type       string   `json:"type,omitempty"`       //
	UpdatedAt  string   `json:"updated_at,omitempty"` //
}
type ResponseOrganizationsCreateOrganizationPolicyObject interface{}
type ResponseOrganizationsGetOrganizationPolicyObjectsGroups []ResponseItemOrganizationsGetOrganizationPolicyObjectsGroups // Array of ResponseOrganizationsGetOrganizationPolicyObjectsGroups
type ResponseItemOrganizationsGetOrganizationPolicyObjectsGroups struct {
	Category   string   `json:"category,omitempty"`   //
	CreatedAt  string   `json:"created_at,omitempty"` //
	ID         string   `json:"id,omitempty"`         //
	Name       string   `json:"name,omitempty"`       //
	NetworkIDs []string `json:"networkIds,omitempty"` //
	ObjectIDs  []string `json:"objectIds,omitempty"`  //
	UpdatedAt  string   `json:"updated_at,omitempty"` //
}
type ResponseOrganizationsCreateOrganizationPolicyObjectsGroup interface{}
type ResponseOrganizationsGetOrganizationPolicyObjectsGroup struct {
	Category   string   `json:"category,omitempty"`   //
	CreatedAt  string   `json:"created_at,omitempty"` //
	ID         string   `json:"id,omitempty"`         //
	Name       string   `json:"name,omitempty"`       //
	NetworkIDs []string `json:"networkIds,omitempty"` //
	ObjectIDs  []string `json:"objectIds,omitempty"`  //
	UpdatedAt  string   `json:"updated_at,omitempty"` //
}
type ResponseOrganizationsUpdateOrganizationPolicyObjectsGroup interface{}
type ResponseOrganizationsGetOrganizationPolicyObject struct {
	Category   string   `json:"category,omitempty"`   //
	Cidr       string   `json:"cidr,omitempty"`       //
	CreatedAt  string   `json:"created_at,omitempty"` //
	GroupIDs   []string `json:"groupIds,omitempty"`   //
	ID         string   `json:"id,omitempty"`         //
	Name       string   `json:"name,omitempty"`       //
	NetworkIDs []string `json:"networkIds,omitempty"` //
	Type       string   `json:"type,omitempty"`       //
	UpdatedAt  string   `json:"updated_at,omitempty"` //
}
type ResponseOrganizationsUpdateOrganizationPolicyObject interface{}
type ResponseOrganizationsGetOrganizationSaml struct {
	Enabled *bool `json:"enabled,omitempty"` // Toggle depicting if SAML SSO settings are enabled
}
type ResponseOrganizationsUpdateOrganizationSaml struct {
	Enabled *bool `json:"enabled,omitempty"` // Toggle depicting if SAML SSO settings are enabled
}
type ResponseOrganizationsGetOrganizationSamlIDps []ResponseItemOrganizationsGetOrganizationSamlIDps // Array of ResponseOrganizationsGetOrganizationSamlIdps
type ResponseItemOrganizationsGetOrganizationSamlIDps struct {
	ConsumerURL             string `json:"consumerUrl,omitempty"`             // URL that is consuming SAML Identity Provider (IdP)
	IDpID                   string `json:"idpId,omitempty"`                   // ID associated with the SAML Identity Provider (IdP)
	SloLogoutURL            string `json:"sloLogoutUrl,omitempty"`            // Dashboard will redirect users to this URL when they sign out.
	X509CertSha1Fingerprint string `json:"x509certSha1Fingerprint,omitempty"` // Fingerprint (SHA1) of the SAML certificate provided by your Identity Provider (IdP). This will be used for encryption / validation.
}
type ResponseOrganizationsCreateOrganizationSamlIDp []ResponseItemOrganizationsCreateOrganizationSamlIDp // Array of ResponseOrganizationsCreateOrganizationSamlIdp
type ResponseItemOrganizationsCreateOrganizationSamlIDp struct {
	ConsumerURL             string `json:"consumerUrl,omitempty"`             // URL that is consuming SAML Identity Provider (IdP)
	IDpID                   string `json:"idpId,omitempty"`                   // ID associated with the SAML Identity Provider (IdP)
	SloLogoutURL            string `json:"sloLogoutUrl,omitempty"`            // Dashboard will redirect users to this URL when they sign out.
	X509CertSha1Fingerprint string `json:"x509certSha1Fingerprint,omitempty"` // Fingerprint (SHA1) of the SAML certificate provided by your Identity Provider (IdP). This will be used for encryption / validation.
}
type ResponseOrganizationsGetOrganizationSamlIDp struct {
	ConsumerURL             string `json:"consumerUrl,omitempty"`             // URL that is consuming SAML Identity Provider (IdP)
	IDpID                   string `json:"idpId,omitempty"`                   // ID associated with the SAML Identity Provider (IdP)
	SloLogoutURL            string `json:"sloLogoutUrl,omitempty"`            // Dashboard will redirect users to this URL when they sign out.
	X509CertSha1Fingerprint string `json:"x509certSha1Fingerprint,omitempty"` // Fingerprint (SHA1) of the SAML certificate provided by your Identity Provider (IdP). This will be used for encryption / validation.
}
type ResponseOrganizationsUpdateOrganizationSamlIDp []ResponseItemOrganizationsUpdateOrganizationSamlIDp // Array of ResponseOrganizationsUpdateOrganizationSamlIdp
type ResponseItemOrganizationsUpdateOrganizationSamlIDp struct {
	ConsumerURL             string `json:"consumerUrl,omitempty"`             // URL that is consuming SAML Identity Provider (IdP)
	IDpID                   string `json:"idpId,omitempty"`                   // ID associated with the SAML Identity Provider (IdP)
	SloLogoutURL            string `json:"sloLogoutUrl,omitempty"`            // Dashboard will redirect users to this URL when they sign out.
	X509CertSha1Fingerprint string `json:"x509certSha1Fingerprint,omitempty"` // Fingerprint (SHA1) of the SAML certificate provided by your Identity Provider (IdP). This will be used for encryption / validation.
}
type ResponseOrganizationsGetOrganizationSamlRoles []ResponseItemOrganizationsGetOrganizationSamlRoles // Array of ResponseOrganizationsGetOrganizationSamlRoles
type ResponseItemOrganizationsGetOrganizationSamlRoles struct {
	ID        string                                                       `json:"id,omitempty"`        //
	Networks  *[]ResponseItemOrganizationsGetOrganizationSamlRolesNetworks `json:"networks,omitempty"`  //
	OrgAccess string                                                       `json:"orgAccess,omitempty"` //
	Role      string                                                       `json:"role,omitempty"`      //
	Tags      *[]ResponseItemOrganizationsGetOrganizationSamlRolesTags     `json:"tags,omitempty"`      //
}
type ResponseItemOrganizationsGetOrganizationSamlRolesNetworks struct {
	Access string `json:"access,omitempty"` //
	ID     string `json:"id,omitempty"`     //
}
type ResponseItemOrganizationsGetOrganizationSamlRolesTags struct {
	Access string `json:"access,omitempty"` //
	Tag    string `json:"tag,omitempty"`    //
}
type ResponseOrganizationsCreateOrganizationSamlRole interface{}
type ResponseOrganizationsGetOrganizationSamlRole struct {
	ID        string                                                  `json:"id,omitempty"`        //
	Networks  *[]ResponseOrganizationsGetOrganizationSamlRoleNetworks `json:"networks,omitempty"`  //
	OrgAccess string                                                  `json:"orgAccess,omitempty"` //
	Role      string                                                  `json:"role,omitempty"`      //
	Tags      *[]ResponseOrganizationsGetOrganizationSamlRoleTags     `json:"tags,omitempty"`      //
}
type ResponseOrganizationsGetOrganizationSamlRoleNetworks struct {
	Access string `json:"access,omitempty"` //
	ID     string `json:"id,omitempty"`     //
}
type ResponseOrganizationsGetOrganizationSamlRoleTags struct {
	Access string `json:"access,omitempty"` //
	Tag    string `json:"tag,omitempty"`    //
}
type ResponseOrganizationsUpdateOrganizationSamlRole struct {
	ID        string                                                     `json:"id,omitempty"`        // ID associated with the SAML role
	Networks  *[]ResponseOrganizationsUpdateOrganizationSamlRoleNetworks `json:"networks,omitempty"`  // The list of networks that the SAML administrator has privileges on
	OrgAccess string                                                     `json:"orgAccess,omitempty"` // The privilege of the SAML administrator on the organization
	Role      string                                                     `json:"role,omitempty"`      // The role of the SAML administrator
	Tags      *[]ResponseOrganizationsUpdateOrganizationSamlRoleTags     `json:"tags,omitempty"`      // The list of tags that the SAML administrator has privleges on
}
type ResponseOrganizationsUpdateOrganizationSamlRoleNetworks struct {
	Access string `json:"access,omitempty"` // The privilege of the SAML administrator on the network
	ID     string `json:"id,omitempty"`     // The network ID
}
type ResponseOrganizationsUpdateOrganizationSamlRoleTags struct {
	Access string `json:"access,omitempty"` // The privilege of the SAML administrator on the tag
	Tag    string `json:"tag,omitempty"`    // The name of the tag
}
type ResponseOrganizationsGetOrganizationSNMP struct {
	Hostname   string   `json:"hostname,omitempty"`   //
	PeerIPs    []string `json:"peerIps,omitempty"`    //
	Port       *int     `json:"port,omitempty"`       //
	V2CEnabled *bool    `json:"v2cEnabled,omitempty"` //
	V3AuthMode string   `json:"v3AuthMode,omitempty"` //
	V3Enabled  *bool    `json:"v3Enabled,omitempty"`  //
	V3PrivMode string   `json:"v3PrivMode,omitempty"` //
}
type ResponseOrganizationsUpdateOrganizationSNMP interface{}
type ResponseOrganizationsGetOrganizationSummaryTopAppliancesByUtilization []ResponseItemOrganizationsGetOrganizationSummaryTopAppliancesByUtilization // Array of ResponseOrganizationsGetOrganizationSummaryTopAppliancesByUtilization
type ResponseItemOrganizationsGetOrganizationSummaryTopAppliancesByUtilization struct {
	Mac         string                                                                                `json:"mac,omitempty"`         // Mac address of the appliance
	Model       string                                                                                `json:"model,omitempty"`       // Model of the appliance
	Name        string                                                                                `json:"name,omitempty"`        // Name of the appliance
	Network     *ResponseItemOrganizationsGetOrganizationSummaryTopAppliancesByUtilizationNetwork     `json:"network,omitempty"`     // Network info
	Serial      string                                                                                `json:"serial,omitempty"`      // Serial number of the appliance
	Utilization *ResponseItemOrganizationsGetOrganizationSummaryTopAppliancesByUtilizationUtilization `json:"utilization,omitempty"` // Utilization of the appliance
}
type ResponseItemOrganizationsGetOrganizationSummaryTopAppliancesByUtilizationNetwork struct {
	ID   string `json:"id,omitempty"`   // Network id
	Name string `json:"name,omitempty"` // Network name
}
type ResponseItemOrganizationsGetOrganizationSummaryTopAppliancesByUtilizationUtilization struct {
	Average *ResponseItemOrganizationsGetOrganizationSummaryTopAppliancesByUtilizationUtilizationAverage `json:"average,omitempty"` // Average utilization of the appliance
}
type ResponseItemOrganizationsGetOrganizationSummaryTopAppliancesByUtilizationUtilizationAverage struct {
	Percentage *float64 `json:"percentage,omitempty"` // Average percentage utilization of the appliance
}
type ResponseOrganizationsGetOrganizationSummaryTopClientsByUsage []ResponseItemOrganizationsGetOrganizationSummaryTopClientsByUsage // Array of ResponseOrganizationsGetOrganizationSummaryTopClientsByUsage
type ResponseItemOrganizationsGetOrganizationSummaryTopClientsByUsage struct {
	ID      string                                                                   `json:"id,omitempty"`      // ID of client
	Mac     string                                                                   `json:"mac,omitempty"`     // MAC address of client
	Name    string                                                                   `json:"name,omitempty"`    // Name of client
	Network *ResponseItemOrganizationsGetOrganizationSummaryTopClientsByUsageNetwork `json:"network,omitempty"` //
	Usage   *ResponseItemOrganizationsGetOrganizationSummaryTopClientsByUsageUsage   `json:"usage,omitempty"`   // Data usage information
}
type ResponseItemOrganizationsGetOrganizationSummaryTopClientsByUsageNetwork struct {
	ID   string `json:"id,omitempty"`   // ID of network
	Name string `json:"name,omitempty"` // Name of network
}
type ResponseItemOrganizationsGetOrganizationSummaryTopClientsByUsageUsage struct {
	Downstream *float64 `json:"downstream,omitempty"` // Downstream data usage by client
	Percentage *float64 `json:"percentage,omitempty"` // Percentage of total data usage by client
	Total      *float64 `json:"total,omitempty"`      // Total data usage by client
	Upstream   *float64 `json:"upstream,omitempty"`   // Upstream data usage by client
}
type ResponseOrganizationsGetOrganizationSummaryTopClientsManufacturersByUsage []ResponseItemOrganizationsGetOrganizationSummaryTopClientsManufacturersByUsage // Array of ResponseOrganizationsGetOrganizationSummaryTopClientsManufacturersByUsage
type ResponseItemOrganizationsGetOrganizationSummaryTopClientsManufacturersByUsage struct {
	Clients *ResponseItemOrganizationsGetOrganizationSummaryTopClientsManufacturersByUsageClients `json:"clients,omitempty"` // Clients info
	Name    string                                                                                `json:"name,omitempty"`    // Name of the manufacturer
	Usage   *ResponseItemOrganizationsGetOrganizationSummaryTopClientsManufacturersByUsageUsage   `json:"usage,omitempty"`   // Clients usage
}
type ResponseItemOrganizationsGetOrganizationSummaryTopClientsManufacturersByUsageClients struct {
	Counts *ResponseItemOrganizationsGetOrganizationSummaryTopClientsManufacturersByUsageClientsCounts `json:"counts,omitempty"` // Counts of clients
}
type ResponseItemOrganizationsGetOrganizationSummaryTopClientsManufacturersByUsageClientsCounts struct {
	Total *int `json:"total,omitempty"` // Total counts of clients
}
type ResponseItemOrganizationsGetOrganizationSummaryTopClientsManufacturersByUsageUsage struct {
	Downstream *float64 `json:"downstream,omitempty"` // Downstream data usage by client
	Total      *float64 `json:"total,omitempty"`      // Total data usage by client
	Upstream   *float64 `json:"upstream,omitempty"`   // Upstream data usage by client
}
type ResponseOrganizationsGetOrganizationSummaryTopDevicesByUsage []ResponseItemOrganizationsGetOrganizationSummaryTopDevicesByUsage // Array of ResponseOrganizationsGetOrganizationSummaryTopDevicesByUsage
type ResponseItemOrganizationsGetOrganizationSummaryTopDevicesByUsage struct {
	Clients     *ResponseItemOrganizationsGetOrganizationSummaryTopDevicesByUsageClients `json:"clients,omitempty"`     // Clients
	Mac         string                                                                   `json:"mac,omitempty"`         // Mac address of the device
	Model       string                                                                   `json:"model,omitempty"`       // Model of the device
	Name        string                                                                   `json:"name,omitempty"`        // Name of the device
	Network     *ResponseItemOrganizationsGetOrganizationSummaryTopDevicesByUsageNetwork `json:"network,omitempty"`     // Network info
	ProductType string                                                                   `json:"productType,omitempty"` // Product type of the device
	Serial      string                                                                   `json:"serial,omitempty"`      // Serial number of the device
	Usage       *ResponseItemOrganizationsGetOrganizationSummaryTopDevicesByUsageUsage   `json:"usage,omitempty"`       // Data usage of the device
}
type ResponseItemOrganizationsGetOrganizationSummaryTopDevicesByUsageClients struct {
	Counts *ResponseItemOrganizationsGetOrganizationSummaryTopDevicesByUsageClientsCounts `json:"counts,omitempty"` // Counts of clients
}
type ResponseItemOrganizationsGetOrganizationSummaryTopDevicesByUsageClientsCounts struct {
	Total *int `json:"total,omitempty"` // Total counts of clients
}
type ResponseItemOrganizationsGetOrganizationSummaryTopDevicesByUsageNetwork struct {
	ID   string `json:"id,omitempty"`   // Network id
	Name string `json:"name,omitempty"` // Network name
}
type ResponseItemOrganizationsGetOrganizationSummaryTopDevicesByUsageUsage struct {
	Percentage *float64 `json:"percentage,omitempty"` // Data usage of the device by percentage
	Total      *float64 `json:"total,omitempty"`      // Total data usage of the device
}
type ResponseOrganizationsGetOrganizationSummaryTopDevicesModelsByUsage []ResponseItemOrganizationsGetOrganizationSummaryTopDevicesModelsByUsage // Array of ResponseOrganizationsGetOrganizationSummaryTopDevicesModelsByUsage
type ResponseItemOrganizationsGetOrganizationSummaryTopDevicesModelsByUsage struct {
	Count *int                                                                         `json:"count,omitempty"` // Total number of devices per model
	Model string                                                                       `json:"model,omitempty"` // The device model
	Usage *ResponseItemOrganizationsGetOrganizationSummaryTopDevicesModelsByUsageUsage `json:"usage,omitempty"` // Usage info in megabytes
}
type ResponseItemOrganizationsGetOrganizationSummaryTopDevicesModelsByUsageUsage struct {
	Average *float64 `json:"average,omitempty"` // Average usage in megabytes
	Total   *float64 `json:"total,omitempty"`   // Total usage in megabytes
}
type ResponseOrganizationsGetOrganizationSummaryTopSSIDsByUsage []ResponseItemOrganizationsGetOrganizationSummaryTopSSIDsByUsage // Array of ResponseOrganizationsGetOrganizationSummaryTopSsidsByUsage
type ResponseItemOrganizationsGetOrganizationSummaryTopSSIDsByUsage struct {
	Clients *ResponseItemOrganizationsGetOrganizationSummaryTopSSIDsByUsageClients `json:"clients,omitempty"` // Clients info of the SSID
	Name    string                                                                 `json:"name,omitempty"`    // Name of the SSID
	Usage   *ResponseItemOrganizationsGetOrganizationSummaryTopSSIDsByUsageUsage   `json:"usage,omitempty"`   // Date usage of the SSID, in megabytes
}
type ResponseItemOrganizationsGetOrganizationSummaryTopSSIDsByUsageClients struct {
	Counts *ResponseItemOrganizationsGetOrganizationSummaryTopSSIDsByUsageClientsCounts `json:"counts,omitempty"` // Counts of the clients
}
type ResponseItemOrganizationsGetOrganizationSummaryTopSSIDsByUsageClientsCounts struct {
	Total *int `json:"total,omitempty"` // Total counts of the clients
}
type ResponseItemOrganizationsGetOrganizationSummaryTopSSIDsByUsageUsage struct {
	Downstream *float64 `json:"downstream,omitempty"` // Downstream usage of the SSID
	Percentage *float64 `json:"percentage,omitempty"` // Percentage usage of the SSID
	Total      *float64 `json:"total,omitempty"`      // Total usage of the SSID
	Upstream   *float64 `json:"upstream,omitempty"`   // Upstream usage of the SSID
}
type ResponseOrganizationsGetOrganizationSummaryTopSwitchesByEnergyUsage []ResponseItemOrganizationsGetOrganizationSummaryTopSwitchesByEnergyUsage // Array of ResponseOrganizationsGetOrganizationSummaryTopSwitchesByEnergyUsage
type ResponseItemOrganizationsGetOrganizationSummaryTopSwitchesByEnergyUsage struct {
	Mac     string                                                                          `json:"mac,omitempty"`     // Mac address of the switch
	Model   string                                                                          `json:"model,omitempty"`   // Model of the switch
	Name    string                                                                          `json:"name,omitempty"`    // Name of the switch
	Network *ResponseItemOrganizationsGetOrganizationSummaryTopSwitchesByEnergyUsageNetwork `json:"network,omitempty"` // Network info
	Usage   *ResponseItemOrganizationsGetOrganizationSummaryTopSwitchesByEnergyUsageUsage   `json:"usage,omitempty"`   // Energy usage of the switch
}
type ResponseItemOrganizationsGetOrganizationSummaryTopSwitchesByEnergyUsageNetwork struct {
	ID   string `json:"id,omitempty"`   // Network id
	Name string `json:"name,omitempty"` // Network name
}
type ResponseItemOrganizationsGetOrganizationSummaryTopSwitchesByEnergyUsageUsage struct {
	Total *float64 `json:"total,omitempty"` // Total energy usage of the switch
}
type ResponseOrganizationsGetOrganizationUplinksStatuses []ResponseItemOrganizationsGetOrganizationUplinksStatuses // Array of ResponseOrganizationsGetOrganizationUplinksStatuses
type ResponseItemOrganizationsGetOrganizationUplinksStatuses struct {
	LastReportedAt string                                                            `json:"lastReportedAt,omitempty"` // Last reported time for the device
	Model          string                                                            `json:"model,omitempty"`          // The uplink model
	NetworkID      string                                                            `json:"networkId,omitempty"`      // Network identifier
	Serial         string                                                            `json:"serial,omitempty"`         // The uplink serial
	Uplinks        *[]ResponseItemOrganizationsGetOrganizationUplinksStatusesUplinks `json:"uplinks,omitempty"`        // Uplinks
}
type ResponseItemOrganizationsGetOrganizationUplinksStatusesUplinks struct {
	Apn            string                                                                    `json:"apn,omitempty"`            // Access Point Name
	ConnectionType string                                                                    `json:"connectionType,omitempty"` // Connection Type
	DNS1           string                                                                    `json:"dns1,omitempty"`           // Primary DNS IP
	DNS2           string                                                                    `json:"dns2,omitempty"`           // Secondary DNS IP
	Gateway        string                                                                    `json:"gateway,omitempty"`        // Gateway IP
	Iccid          string                                                                    `json:"iccid,omitempty"`          // Integrated Circuit Card Identification Number
	Interface      string                                                                    `json:"interface,omitempty"`      // Uplink interface
	IP             string                                                                    `json:"ip,omitempty"`             // Uplink IP
	IPAssignedBy   string                                                                    `json:"ipAssignedBy,omitempty"`   // The way in which the IP is assigned
	PrimaryDNS     string                                                                    `json:"primaryDns,omitempty"`     // Primary DNS IP
	Provider       string                                                                    `json:"provider,omitempty"`       // Network Provider
	PublicIP       string                                                                    `json:"publicIp,omitempty"`       // Public IP
	SecondaryDNS   string                                                                    `json:"secondaryDns,omitempty"`   // Secondary DNS IP
	SignalStat     *ResponseItemOrganizationsGetOrganizationUplinksStatusesUplinksSignalStat `json:"signalStat,omitempty"`     // Tower Signal Status
	SignalType     string                                                                    `json:"signalType,omitempty"`     // Signal Type
	Status         string                                                                    `json:"status,omitempty"`         // Uplink status
}
type ResponseItemOrganizationsGetOrganizationUplinksStatusesUplinksSignalStat struct {
	Rsrp string `json:"rsrp,omitempty"` // Reference Signal Received Power
	Rsrq string `json:"rsrq,omitempty"` // Reference Signal Received Quality
}
type ResponseOrganizationsGetOrganizationWebhooksAlertTypes []ResponseItemOrganizationsGetOrganizationWebhooksAlertTypes // Array of ResponseOrganizationsGetOrganizationWebhooksAlertTypes
type ResponseItemOrganizationsGetOrganizationWebhooksAlertTypes struct {
	AlertData        *ResponseItemOrganizationsGetOrganizationWebhooksAlertTypesAlertData `json:"alertData,omitempty"`        //
	AlertID          string                                                               `json:"alertId,omitempty"`          //
	AlertLevel       string                                                               `json:"alertLevel,omitempty"`       //
	AlertType        string                                                               `json:"alertType,omitempty"`        //
	AlertTypeID      string                                                               `json:"alertTypeId,omitempty"`      //
	DeviceMac        string                                                               `json:"deviceMac,omitempty"`        //
	DeviceModel      string                                                               `json:"deviceModel,omitempty"`      //
	DeviceName       string                                                               `json:"deviceName,omitempty"`       //
	DeviceSerial     string                                                               `json:"deviceSerial,omitempty"`     //
	DeviceTags       []string                                                             `json:"deviceTags,omitempty"`       //
	DeviceURL        string                                                               `json:"deviceUrl,omitempty"`        //
	EnrollmentString string                                                               `json:"enrollmentString,omitempty"` //
	NetworkID        string                                                               `json:"networkId,omitempty"`        //
	NetworkName      string                                                               `json:"networkName,omitempty"`      //
	NetworkURL       string                                                               `json:"networkUrl,omitempty"`       //
	Notes            string                                                               `json:"notes,omitempty"`            //
	OccurredAt       string                                                               `json:"occurredAt,omitempty"`       //
	OrganizationID   string                                                               `json:"organizationId,omitempty"`   //
	OrganizationName string                                                               `json:"organizationName,omitempty"` //
	OrganizationURL  string                                                               `json:"organizationUrl,omitempty"`  //
	ProductTypes     []string                                                             `json:"productTypes,omitempty"`     //
	SentAt           string                                                               `json:"sentAt,omitempty"`           //
	SharedSecret     string                                                               `json:"sharedSecret,omitempty"`     //
	Version          string                                                               `json:"version,omitempty"`          //
}
type ResponseItemOrganizationsGetOrganizationWebhooksAlertTypesAlertData interface{}
type ResponseOrganizationsGetOrganizationWebhooksLogs []ResponseItemOrganizationsGetOrganizationWebhooksLogs // Array of ResponseOrganizationsGetOrganizationWebhooksLogs
type ResponseItemOrganizationsGetOrganizationWebhooksLogs struct {
	AlertType        string `json:"alertType,omitempty"`        // Type of alert that the webhook is delivering
	LoggedAt         string `json:"loggedAt,omitempty"`         // When the webhook log was created, in ISO8601 format
	NetworkID        string `json:"networkId,omitempty"`        // Network ID for the webhook log
	OrganizationID   string `json:"organizationId,omitempty"`   // ID for the webhook log's organization
	ResponseCode     *int   `json:"responseCode,omitempty"`     // Response code from the webhook
	ResponseDuration *int   `json:"responseDuration,omitempty"` // Duration of the response, in milliseconds
	SentAt           string `json:"sentAt,omitempty"`           // When the webhook was sent, in ISO8601 format
	URL              string `json:"url,omitempty"`              // URL where the webhook was sent
}
type RequestOrganizationsCreateOrganization struct {
	Management *RequestOrganizationsCreateOrganizationManagement `json:"management,omitempty"` // Information about the organization's management system
	Name       string                                            `json:"name,omitempty"`       // The name of the organization
}
type RequestOrganizationsCreateOrganizationManagement struct {
	Details *[]RequestOrganizationsCreateOrganizationManagementDetails `json:"details,omitempty"` // Details related to organization management, possibly empty
}
type RequestOrganizationsCreateOrganizationManagementDetails struct {
	Name  string `json:"name,omitempty"`  // Name of management data
	Value string `json:"value,omitempty"` // Value of management data
}
type RequestOrganizationsUpdateOrganization struct {
	API        *RequestOrganizationsUpdateOrganizationAPI        `json:"api,omitempty"`        // API-specific settings
	Management *RequestOrganizationsUpdateOrganizationManagement `json:"management,omitempty"` // Information about the organization's management system
	Name       string                                            `json:"name,omitempty"`       // The name of the organization
}
type RequestOrganizationsUpdateOrganizationAPI struct {
	Enabled *bool `json:"enabled,omitempty"` // If true, enable the access to the Cisco Meraki Dashboard API
}
type RequestOrganizationsUpdateOrganizationManagement struct {
	Details *[]RequestOrganizationsUpdateOrganizationManagementDetails `json:"details,omitempty"` // Details related to organization management, possibly empty
}
type RequestOrganizationsUpdateOrganizationManagementDetails struct {
	Name  string `json:"name,omitempty"`  // Name of management data
	Value string `json:"value,omitempty"` // Value of management data
}
type RequestOrganizationsCreateOrganizationActionBatch struct {
	Actions     *[]RequestOrganizationsCreateOrganizationActionBatchActions `json:"actions,omitempty"`     // A set of changes to make as part of this action (<a href='https://developer.cisco.com/meraki/api/#/rest/guides/action-batches/'>more details</a>)
	Confirmed   *bool                                                       `json:"confirmed,omitempty"`   // Set to true for immediate execution. Set to false if the action should be previewed before executing. This property cannot be unset once it is true. Defaults to false.
	Synchronous *bool                                                       `json:"synchronous,omitempty"` // Set to true to force the batch to run synchronous. There can be at most 20 actions in synchronous batch. Defaults to false.
}
type RequestOrganizationsCreateOrganizationActionBatchActions struct {
	Body      *RequestOrganizationsCreateOrganizationActionBatchActionsBody `json:"body,omitempty"`      // The body of the action
	Operation string                                                        `json:"operation,omitempty"` // The operation to be used
	Resource  string                                                        `json:"resource,omitempty"`  // Unique identifier for the resource to be acted on
}
type RequestOrganizationsCreateOrganizationActionBatchActionsBody interface{}
type RequestOrganizationsUpdateOrganizationActionBatch struct {
	Confirmed   *bool `json:"confirmed,omitempty"`   // A boolean representing whether or not the batch has been confirmed. This property cannot be unset once it is true.
	Synchronous *bool `json:"synchronous,omitempty"` // Set to true to force the batch to run synchronous. There can be at most 20 actions in synchronous batch.
}
type RequestOrganizationsCreateOrganizationAdaptivePolicyACL struct {
	Description string                                                          `json:"description,omitempty"` // Description of the adaptive policy ACL
	IPVersion   string                                                          `json:"ipVersion,omitempty"`   // IP version of adpative policy ACL. One of: 'any', 'ipv4' or 'ipv6'
	Name        string                                                          `json:"name,omitempty"`        // Name of the adaptive policy ACL
	Rules       *[]RequestOrganizationsCreateOrganizationAdaptivePolicyACLRules `json:"rules,omitempty"`       // An ordered array of the adaptive policy ACL rules.
}
type RequestOrganizationsCreateOrganizationAdaptivePolicyACLRules struct {
	DstPort  string `json:"dstPort,omitempty"`  // Destination port. Must be in the format of single port: '1', port list: '1,2' or port range: '1-10', and in the range of 1-65535, or 'any'. Default is 'any'.
	Policy   string `json:"policy,omitempty"`   // 'allow' or 'deny' traffic specified by this rule.
	Protocol string `json:"protocol,omitempty"` // The type of protocol (must be 'tcp', 'udp', 'icmp' or 'any').
	SrcPort  string `json:"srcPort,omitempty"`  // Source port. Must be in the format of single port: '1', port list: '1,2' or port range: '1-10', and in the range of 1-65535, or 'any'. Default is 'any'.
}
type RequestOrganizationsUpdateOrganizationAdaptivePolicyACL struct {
	Description string                                                          `json:"description,omitempty"` // Description of the adaptive policy ACL
	IPVersion   string                                                          `json:"ipVersion,omitempty"`   // IP version of adpative policy ACL. One of: 'any', 'ipv4' or 'ipv6'
	Name        string                                                          `json:"name,omitempty"`        // Name of the adaptive policy ACL
	Rules       *[]RequestOrganizationsUpdateOrganizationAdaptivePolicyACLRules `json:"rules,omitempty"`       // An ordered array of the adaptive policy ACL rules. An empty array will clear the rules.
}
type RequestOrganizationsUpdateOrganizationAdaptivePolicyACLRules struct {
	DstPort  string `json:"dstPort,omitempty"`  // Destination port. Must be in the format of single port: '1', port list: '1,2' or port range: '1-10', and in the range of 1-65535, or 'any'. Default is 'any'.
	Policy   string `json:"policy,omitempty"`   // 'allow' or 'deny' traffic specified by this rule.
	Protocol string `json:"protocol,omitempty"` // The type of protocol (must be 'tcp', 'udp', 'icmp' or 'any').
	SrcPort  string `json:"srcPort,omitempty"`  // Source port. Must be in the format of single port: '1', port list: '1,2' or port range: '1-10', and in the range of 1-65535, or 'any'. Default is 'any'.
}
type RequestOrganizationsCreateOrganizationAdaptivePolicyGroup struct {
	Description   string                                                                    `json:"description,omitempty"`   // Description of the group (default: "")
	Name          string                                                                    `json:"name,omitempty"`          // Name of the group
	PolicyObjects *[]RequestOrganizationsCreateOrganizationAdaptivePolicyGroupPolicyObjects `json:"policyObjects,omitempty"` // The policy objects that belong to this group; traffic from addresses specified by these policy objects will be tagged with this group's SGT value if no other tagging scheme is being used (each requires one unique attribute) (default: [])
	Sgt           *int                                                                      `json:"sgt,omitempty"`           // SGT value of the group
}
type RequestOrganizationsCreateOrganizationAdaptivePolicyGroupPolicyObjects struct {
	ID   string `json:"id,omitempty"`   // The ID of the policy object
	Name string `json:"name,omitempty"` // The name of the policy object
}
type RequestOrganizationsUpdateOrganizationAdaptivePolicyGroup struct {
	Description   string                                                                    `json:"description,omitempty"`   // Description of the group
	Name          string                                                                    `json:"name,omitempty"`          // Name of the group
	PolicyObjects *[]RequestOrganizationsUpdateOrganizationAdaptivePolicyGroupPolicyObjects `json:"policyObjects,omitempty"` // The policy objects that belong to this group; traffic from addresses specified by these policy objects will be tagged with this group's SGT value if no other tagging scheme is being used (each requires one unique attribute)
	Sgt           *int                                                                      `json:"sgt,omitempty"`           // SGT value of the group
}
type RequestOrganizationsUpdateOrganizationAdaptivePolicyGroupPolicyObjects struct {
	ID   string `json:"id,omitempty"`   // The ID of the policy object
	Name string `json:"name,omitempty"` // The name of the policy object
}
type RequestOrganizationsCreateOrganizationAdaptivePolicyPolicy struct {
	ACLs             *[]RequestOrganizationsCreateOrganizationAdaptivePolicyPolicyACLs           `json:"acls,omitempty"`             // An ordered array of adaptive policy ACLs (each requires one unique attribute) that apply to this policy (default: [])
	DestinationGroup *RequestOrganizationsCreateOrganizationAdaptivePolicyPolicyDestinationGroup `json:"destinationGroup,omitempty"` // The destination adaptive policy group (requires one unique attribute)
	LastEntryRule    string                                                                      `json:"lastEntryRule,omitempty"`    // The rule to apply if there is no matching ACL (default: "default")
	SourceGroup      *RequestOrganizationsCreateOrganizationAdaptivePolicyPolicySourceGroup      `json:"sourceGroup,omitempty"`      // The source adaptive policy group (requires one unique attribute)
}
type RequestOrganizationsCreateOrganizationAdaptivePolicyPolicyACLs struct {
	ID   string `json:"id,omitempty"`   // The ID of the adaptive policy ACL
	Name string `json:"name,omitempty"` // The name of the adaptive policy ACL
}
type RequestOrganizationsCreateOrganizationAdaptivePolicyPolicyDestinationGroup struct {
	ID   string `json:"id,omitempty"`   // The ID of the destination adaptive policy group
	Name string `json:"name,omitempty"` // The name of the destination adaptive policy group
	Sgt  *int   `json:"sgt,omitempty"`  // The SGT of the destination adaptive policy group
}
type RequestOrganizationsCreateOrganizationAdaptivePolicyPolicySourceGroup struct {
	ID   string `json:"id,omitempty"`   // The ID of the source adaptive policy group
	Name string `json:"name,omitempty"` // The name of the source adaptive policy group
	Sgt  *int   `json:"sgt,omitempty"`  // The SGT of the source adaptive policy group
}
type RequestOrganizationsUpdateOrganizationAdaptivePolicyPolicy struct {
	ACLs             *[]RequestOrganizationsUpdateOrganizationAdaptivePolicyPolicyACLs           `json:"acls,omitempty"`             // An ordered array of adaptive policy ACLs (each requires one unique attribute) that apply to this policy
	DestinationGroup *RequestOrganizationsUpdateOrganizationAdaptivePolicyPolicyDestinationGroup `json:"destinationGroup,omitempty"` // The destination adaptive policy group (requires one unique attribute)
	LastEntryRule    string                                                                      `json:"lastEntryRule,omitempty"`    // The rule to apply if there is no matching ACL
	SourceGroup      *RequestOrganizationsUpdateOrganizationAdaptivePolicyPolicySourceGroup      `json:"sourceGroup,omitempty"`      // The source adaptive policy group (requires one unique attribute)
}
type RequestOrganizationsUpdateOrganizationAdaptivePolicyPolicyACLs struct {
	ID   string `json:"id,omitempty"`   // The ID of the adaptive policy ACL
	Name string `json:"name,omitempty"` // The name of the adaptive policy ACL
}
type RequestOrganizationsUpdateOrganizationAdaptivePolicyPolicyDestinationGroup struct {
	ID   string `json:"id,omitempty"`   // The ID of the destination adaptive policy group
	Name string `json:"name,omitempty"` // The name of the destination adaptive policy group
	Sgt  *int   `json:"sgt,omitempty"`  // The SGT of the destination adaptive policy group
}
type RequestOrganizationsUpdateOrganizationAdaptivePolicyPolicySourceGroup struct {
	ID   string `json:"id,omitempty"`   // The ID of the source adaptive policy group
	Name string `json:"name,omitempty"` // The name of the source adaptive policy group
	Sgt  *int   `json:"sgt,omitempty"`  // The SGT of the source adaptive policy group
}
type RequestOrganizationsUpdateOrganizationAdaptivePolicySettings struct {
	EnabledNetworks []string `json:"enabledNetworks,omitempty"` // List of network IDs with adaptive policy enabled
}
type RequestOrganizationsCreateOrganizationAdmin struct {
	AuthenticationMethod string                                                 `json:"authenticationMethod,omitempty"` // The method of authentication the user will use to sign in to the Meraki dashboard. Can be one of 'Email' or 'Cisco SecureX Sign-On'. The default is Email authentication
	Email                string                                                 `json:"email,omitempty"`                // The email of the dashboard administrator. This attribute can not be updated.
	Name                 string                                                 `json:"name,omitempty"`                 // The name of the dashboard administrator
	Networks             *[]RequestOrganizationsCreateOrganizationAdminNetworks `json:"networks,omitempty"`             // The list of networks that the dashboard administrator has privileges on
	OrgAccess            string                                                 `json:"orgAccess,omitempty"`            // The privilege of the dashboard administrator on the organization. Can be one of 'full', 'read-only', 'enterprise' or 'none'
	Tags                 *[]RequestOrganizationsCreateOrganizationAdminTags     `json:"tags,omitempty"`                 // The list of tags that the dashboard administrator has privileges on
}
type RequestOrganizationsCreateOrganizationAdminNetworks struct {
	Access string `json:"access,omitempty"` // The privilege of the dashboard administrator on the network. Can be one of 'full', 'read-only', 'guest-ambassador' or 'monitor-only'
	ID     string `json:"id,omitempty"`     // The network ID
}
type RequestOrganizationsCreateOrganizationAdminTags struct {
	Access string `json:"access,omitempty"` // The privilege of the dashboard administrator on the tag. Can be one of 'full', 'read-only', 'guest-ambassador' or 'monitor-only'
	Tag    string `json:"tag,omitempty"`    // The name of the tag
}
type RequestOrganizationsUpdateOrganizationAdmin struct {
	Name      string                                                 `json:"name,omitempty"`      // The name of the dashboard administrator
	Networks  *[]RequestOrganizationsUpdateOrganizationAdminNetworks `json:"networks,omitempty"`  // The list of networks that the dashboard administrator has privileges on
	OrgAccess string                                                 `json:"orgAccess,omitempty"` // The privilege of the dashboard administrator on the organization. Can be one of 'full', 'read-only', 'enterprise' or 'none'
	Tags      *[]RequestOrganizationsUpdateOrganizationAdminTags     `json:"tags,omitempty"`      // The list of tags that the dashboard administrator has privileges on
}
type RequestOrganizationsUpdateOrganizationAdminNetworks struct {
	Access string `json:"access,omitempty"` // The privilege of the dashboard administrator on the network. Can be one of 'full', 'read-only', 'guest-ambassador' or 'monitor-only'
	ID     string `json:"id,omitempty"`     // The network ID
}
type RequestOrganizationsUpdateOrganizationAdminTags struct {
	Access string `json:"access,omitempty"` // The privilege of the dashboard administrator on the tag. Can be one of 'full', 'read-only', 'guest-ambassador' or 'monitor-only'
	Tag    string `json:"tag,omitempty"`    // The name of the tag
}
type RequestOrganizationsCreateOrganizationAlertsProfile struct {
	AlertCondition *RequestOrganizationsCreateOrganizationAlertsProfileAlertCondition `json:"alertCondition,omitempty"` // The conditions that determine if the alert triggers
	Description    string                                                             `json:"description,omitempty"`    // User supplied description of the alert
	NetworkTags    []string                                                           `json:"networkTags,omitempty"`    // Networks with these tags will be monitored for the alert
	Recipients     *RequestOrganizationsCreateOrganizationAlertsProfileRecipients     `json:"recipients,omitempty"`     // List of recipients that will recieve the alert.
	Type           string                                                             `json:"type,omitempty"`           // The alert type
}
type RequestOrganizationsCreateOrganizationAlertsProfileAlertCondition struct {
	BitRateBps *int     `json:"bit_rate_bps,omitempty"` // The threshold the metric must cross to be valid for alerting. Used only for WAN Utilization alerts.
	Duration   *int     `json:"duration,omitempty"`     // The total duration in seconds that the threshold should be crossed before alerting
	Interface  string   `json:"interface,omitempty"`    // The uplink observed for the alert.  interface must be one of the following: wan1, wan2, cellular
	JitterMs   *int     `json:"jitter_ms,omitempty"`    // The threshold the metric must cross to be valid for alerting. Used only for VoIP Jitter alerts.
	LatencyMs  *int     `json:"latency_ms,omitempty"`   // The threshold the metric must cross to be valid for alerting. Used only for WAN Latency alerts.
	LossRatio  *float64 `json:"loss_ratio,omitempty"`   // The threshold the metric must cross to be valid for alerting. Used only for Packet Loss alerts.
	Mos        *float64 `json:"mos,omitempty"`          // The threshold the metric must drop below to be valid for alerting. Used only for VoIP MOS alerts.
	Window     *int     `json:"window,omitempty"`       // The look back period in seconds for sensing the alert
}
type RequestOrganizationsCreateOrganizationAlertsProfileRecipients struct {
	Emails        []string `json:"emails,omitempty"`        // A list of emails that will receive information about the alert
	HTTPServerIDs []string `json:"httpServerIds,omitempty"` // A list base64 encoded urls of webhook endpoints that will receive information about the alert
}
type RequestOrganizationsUpdateOrganizationAlertsProfile struct {
	AlertCondition *RequestOrganizationsUpdateOrganizationAlertsProfileAlertCondition `json:"alertCondition,omitempty"` // The conditions that determine if the alert triggers
	Description    string                                                             `json:"description,omitempty"`    // User supplied description of the alert
	Enabled        *bool                                                              `json:"enabled,omitempty"`        // Is the alert config enabled
	NetworkTags    []string                                                           `json:"networkTags,omitempty"`    // Networks with these tags will be monitored for the alert
	Recipients     *RequestOrganizationsUpdateOrganizationAlertsProfileRecipients     `json:"recipients,omitempty"`     // List of recipients that will recieve the alert.
	Type           string                                                             `json:"type,omitempty"`           // The alert type
}
type RequestOrganizationsUpdateOrganizationAlertsProfileAlertCondition struct {
	BitRateBps *int     `json:"bit_rate_bps,omitempty"` // The threshold the metric must cross to be valid for alerting. Used only for WAN Utilization alerts.
	Duration   *int     `json:"duration,omitempty"`     // The total duration in seconds that the threshold should be crossed before alerting
	Interface  string   `json:"interface,omitempty"`    // The uplink observed for the alert.  interface must be one of the following: wan1, wan2, cellular
	JitterMs   *int     `json:"jitter_ms,omitempty"`    // The threshold the metric must cross to be valid for alerting. Used only for VoIP Jitter alerts.
	LatencyMs  *int     `json:"latency_ms,omitempty"`   // The threshold the metric must cross to be valid for alerting. Used only for WAN Latency alerts.
	LossRatio  *float64 `json:"loss_ratio,omitempty"`   // The threshold the metric must cross to be valid for alerting. Used only for Packet Loss alerts.
	Mos        *float64 `json:"mos,omitempty"`          // The threshold the metric must drop below to be valid for alerting. Used only for VoIP MOS alerts.
	Window     *int     `json:"window,omitempty"`       // The look back period in seconds for sensing the alert
}
type RequestOrganizationsUpdateOrganizationAlertsProfileRecipients struct {
	Emails        []string `json:"emails,omitempty"`        // A list of emails that will receive information about the alert
	HTTPServerIDs []string `json:"httpServerIds,omitempty"` // A list base64 encoded urls of webhook endpoints that will receive information about the alert
}
type RequestOrganizationsCreateOrganizationBrandingPolicy struct {
	AdminSettings *RequestOrganizationsCreateOrganizationBrandingPolicyAdminSettings `json:"adminSettings,omitempty"` // Settings for describing which kinds of admins this policy applies to.
	CustomLogo    *RequestOrganizationsCreateOrganizationBrandingPolicyCustomLogo    `json:"customLogo,omitempty"`    // Properties describing the custom logo attached to the branding policy.
	Enabled       *bool                                                              `json:"enabled,omitempty"`       // Boolean indicating whether this policy is enabled.
	HelpSettings  *RequestOrganizationsCreateOrganizationBrandingPolicyHelpSettings  `json:"helpSettings,omitempty"`  //       Settings for describing the modifications to various Help page features. Each property in this object accepts one of       'default or inherit' (do not modify functionality), 'hide' (remove the section from Dashboard), or 'show' (always show       the section on Dashboard). Some properties in this object also accept custom HTML used to replace the section on       Dashboard; see the documentation for each property to see the allowed values.  Each property defaults to 'default or inherit' when not provided.
	Name          string                                                             `json:"name,omitempty"`          // Name of the Dashboard branding policy.
}
type RequestOrganizationsCreateOrganizationBrandingPolicyAdminSettings struct {
	AppliesTo string   `json:"appliesTo,omitempty"` // Which kinds of admins this policy applies to. Can be one of 'All organization admins', 'All enterprise admins', 'All network admins', 'All admins of networks...', 'All admins of networks tagged...', 'Specific admins...', 'All admins' or 'All SAML admins'.
	Values    []string `json:"values,omitempty"`    //       If 'appliesTo' is set to one of 'Specific admins...', 'All admins of networks...' or 'All admins of networks tagged...', then you must specify this 'values' property to provide the set of       entities to apply the branding policy to. For 'Specific admins...', specify an array of admin IDs. For 'All admins of       networks...', specify an array of network IDs and/or configuration template IDs. For 'All admins of networks tagged...',       specify an array of tag names.
}
type RequestOrganizationsCreateOrganizationBrandingPolicyCustomLogo struct {
	Enabled *bool                                                                `json:"enabled,omitempty"` // Whether or not there is a custom logo enabled.
	Image   *RequestOrganizationsCreateOrganizationBrandingPolicyCustomLogoImage `json:"image,omitempty"`   // Properties for setting the image.
}
type RequestOrganizationsCreateOrganizationBrandingPolicyCustomLogoImage struct {
	Contents string `json:"contents,omitempty"` // The file contents (a base 64 encoded string) of your new logo.
	Format   string `json:"format,omitempty"`   // The format of the encoded contents.  Supported formats are 'png', 'gif', and jpg'.
}
type RequestOrganizationsCreateOrganizationBrandingPolicyHelpSettings struct {
	APIDocsSubtab                      string `json:"apiDocsSubtab,omitempty"`                      //       The 'Help -> API docs' subtab where a detailed description of the Dashboard API is listed. Can be one of       'default or inherit', 'hide' or 'show'.
	CasesSubtab                        string `json:"casesSubtab,omitempty"`                        //       The 'Help -> Cases' Dashboard subtab on which Cisco Meraki support cases for this organization can be managed. Can be one       of 'default or inherit', 'hide' or 'show'.
	CiscoMerakiProductDocumentation    string `json:"ciscoMerakiProductDocumentation,omitempty"`    //       The 'Product Manuals' section of the 'Help -> Get Help' subtab. Can be one of 'default or inherit', 'hide', 'show', or a replacement custom HTML string.
	CommunitySubtab                    string `json:"communitySubtab,omitempty"`                    //       The 'Help -> Community' subtab which provides a link to Meraki Community. Can be one of 'default or inherit', 'hide' or 'show'.
	DataProtectionRequestsSubtab       string `json:"dataProtectionRequestsSubtab,omitempty"`       //       The 'Help -> Data protection requests' Dashboard subtab on which requests to delete, restrict, or export end-user data can       be audited. Can be one of 'default or inherit', 'hide' or 'show'.
	FirewallInfoSubtab                 string `json:"firewallInfoSubtab,omitempty"`                 //       The 'Help -> Firewall info' subtab where necessary upstream firewall rules for communication to the Cisco Meraki cloud are       listed. Can be one of 'default or inherit', 'hide' or 'show'.
	GetHelpSubtab                      string `json:"getHelpSubtab,omitempty"`                      //       The 'Help -> Get Help' subtab on which Cisco Meraki KB, Product Manuals, and Support/Case Information are displayed. Note       that if this subtab is hidden, branding customizations for the KB on 'Get help', Cisco Meraki product documentation,       and support contact info will not be visible. Can be one of 'default or inherit', 'hide' or 'show'.
	GetHelpSubtabKnowledgeBaseSearch   string `json:"getHelpSubtabKnowledgeBaseSearch,omitempty"`   //       The KB search box which appears on the Help page. Can be one of 'default or inherit', 'hide', 'show', or a replacement custom HTML string.
	HardwareReplacementsSubtab         string `json:"hardwareReplacementsSubtab,omitempty"`         //       The 'Help -> Replacement info' subtab where important information regarding device replacements is detailed. Can be one of       'default or inherit', 'hide' or 'show'.
	HelpTab                            string `json:"helpTab,omitempty"`                            //       The Help tab, under which all support information resides. If this tab is hidden, no other 'Help' branding       customizations will be visible. Can be one of 'default or inherit', 'hide' or 'show'.
	HelpWidget                         string `json:"helpWidget,omitempty"`                         //       The 'Help Widget' is a support widget which provides access to live chat, documentation links, Sales contact info,       and other contact avenues to reach Meraki Support. Can be one of 'default or inherit', 'hide' or 'show'.
	NewFeaturesSubtab                  string `json:"newFeaturesSubtab,omitempty"`                  //       The 'Help -> New features' subtab where new Dashboard features are detailed. Can be one of 'default or inherit', 'hide' or 'show'.
	SmForums                           string `json:"smForums,omitempty"`                           //       The 'SM Forums' subtab which links to community-based support for Cisco Meraki Systems Manager. Only configurable for       organizations that contain Systems Manager networks. Can be one of 'default or inherit', 'hide' or 'show'.
	SupportContactInfo                 string `json:"supportContactInfo,omitempty"`                 //       The 'Contact Meraki Support' section of the 'Help -> Get Help' subtab. Can be one of 'default or inherit', 'hide', 'show', or a replacement custom HTML string.
	UniversalSearchKnowledgeBaseSearch string `json:"universalSearchKnowledgeBaseSearch,omitempty"` //       The universal search box always visible on Dashboard will, by default, present results from the Meraki KB. This configures       whether these Meraki KB results should be returned. Can be one of 'default or inherit', 'hide' or 'show'.
}
type RequestOrganizationsUpdateOrganizationBrandingPoliciesPriorities struct {
	BrandingPolicyIDs []string `json:"brandingPolicyIds,omitempty"` //       An ordered list of branding policy IDs that determines the priority order of how to apply the policies
}
type RequestOrganizationsUpdateOrganizationBrandingPolicy struct {
	AdminSettings *RequestOrganizationsUpdateOrganizationBrandingPolicyAdminSettings `json:"adminSettings,omitempty"` // Settings for describing which kinds of admins this policy applies to.
	CustomLogo    *RequestOrganizationsUpdateOrganizationBrandingPolicyCustomLogo    `json:"customLogo,omitempty"`    // Properties describing the custom logo attached to the branding policy.
	Enabled       *bool                                                              `json:"enabled,omitempty"`       // Boolean indicating whether this policy is enabled.
	HelpSettings  *RequestOrganizationsUpdateOrganizationBrandingPolicyHelpSettings  `json:"helpSettings,omitempty"`  //       Settings for describing the modifications to various Help page features. Each property in this object accepts one of       'default or inherit' (do not modify functionality), 'hide' (remove the section from Dashboard), or 'show' (always show       the section on Dashboard). Some properties in this object also accept custom HTML used to replace the section on       Dashboard; see the documentation for each property to see the allowed values.
	Name          string                                                             `json:"name,omitempty"`          // Name of the Dashboard branding policy.
}
type RequestOrganizationsUpdateOrganizationBrandingPolicyAdminSettings struct {
	AppliesTo string   `json:"appliesTo,omitempty"` // Which kinds of admins this policy applies to. Can be one of 'All organization admins', 'All enterprise admins', 'All network admins', 'All admins of networks...', 'All admins of networks tagged...', 'Specific admins...', 'All admins' or 'All SAML admins'.
	Values    []string `json:"values,omitempty"`    //       If 'appliesTo' is set to one of 'Specific admins...', 'All admins of networks...' or 'All admins of networks tagged...', then you must specify this 'values' property to provide the set of       entities to apply the branding policy to. For 'Specific admins...', specify an array of admin IDs. For 'All admins of       networks...', specify an array of network IDs and/or configuration template IDs. For 'All admins of networks tagged...',       specify an array of tag names.
}
type RequestOrganizationsUpdateOrganizationBrandingPolicyCustomLogo struct {
	Enabled *bool                                                                `json:"enabled,omitempty"` // Whether or not there is a custom logo enabled.
	Image   *RequestOrganizationsUpdateOrganizationBrandingPolicyCustomLogoImage `json:"image,omitempty"`   // Properties for setting the image.
}
type RequestOrganizationsUpdateOrganizationBrandingPolicyCustomLogoImage struct {
	Contents string `json:"contents,omitempty"` // The file contents (a base 64 encoded string) of your new logo.
	Format   string `json:"format,omitempty"`   // The format of the encoded contents.  Supported formats are 'png', 'gif', and jpg'.
}
type RequestOrganizationsUpdateOrganizationBrandingPolicyHelpSettings struct {
	APIDocsSubtab                      string `json:"apiDocsSubtab,omitempty"`                      //       The 'Help -> API docs' subtab where a detailed description of the Dashboard API is listed. Can be one of       'default or inherit', 'hide' or 'show'.
	CasesSubtab                        string `json:"casesSubtab,omitempty"`                        //       The 'Help -> Cases' Dashboard subtab on which Cisco Meraki support cases for this organization can be managed. Can be one       of 'default or inherit', 'hide' or 'show'.
	CiscoMerakiProductDocumentation    string `json:"ciscoMerakiProductDocumentation,omitempty"`    //       The 'Product Manuals' section of the 'Help -> Get Help' subtab. Can be one of 'default or inherit', 'hide', 'show', or a replacement custom HTML string.
	CommunitySubtab                    string `json:"communitySubtab,omitempty"`                    //       The 'Help -> Community' subtab which provides a link to Meraki Community. Can be one of 'default or inherit', 'hide' or 'show'.
	DataProtectionRequestsSubtab       string `json:"dataProtectionRequestsSubtab,omitempty"`       //       The 'Help -> Data protection requests' Dashboard subtab on which requests to delete, restrict, or export end-user data can       be audited. Can be one of 'default or inherit', 'hide' or 'show'.
	FirewallInfoSubtab                 string `json:"firewallInfoSubtab,omitempty"`                 //       The 'Help -> Firewall info' subtab where necessary upstream firewall rules for communication to the Cisco Meraki cloud are       listed. Can be one of 'default or inherit', 'hide' or 'show'.
	GetHelpSubtab                      string `json:"getHelpSubtab,omitempty"`                      //       The 'Help -> Get Help' subtab on which Cisco Meraki KB, Product Manuals, and Support/Case Information are displayed. Note       that if this subtab is hidden, branding customizations for the KB on 'Get help', Cisco Meraki product documentation,       and support contact info will not be visible. Can be one of 'default or inherit', 'hide' or 'show'.
	GetHelpSubtabKnowledgeBaseSearch   string `json:"getHelpSubtabKnowledgeBaseSearch,omitempty"`   //       The KB search box which appears on the Help page. Can be one of 'default or inherit', 'hide', 'show', or a replacement custom HTML string.
	HardwareReplacementsSubtab         string `json:"hardwareReplacementsSubtab,omitempty"`         //       The 'Help -> Replacement info' subtab where important information regarding device replacements is detailed. Can be one of       'default or inherit', 'hide' or 'show'.
	HelpTab                            string `json:"helpTab,omitempty"`                            //       The Help tab, under which all support information resides. If this tab is hidden, no other 'Help' branding       customizations will be visible. Can be one of 'default or inherit', 'hide' or 'show'.
	HelpWidget                         string `json:"helpWidget,omitempty"`                         //       The 'Help Widget' is a support widget which provides access to live chat, documentation links, Sales contact info,       and other contact avenues to reach Meraki Support. Can be one of 'default or inherit', 'hide' or 'show'.
	NewFeaturesSubtab                  string `json:"newFeaturesSubtab,omitempty"`                  //       The 'Help -> New features' subtab where new Dashboard features are detailed. Can be one of 'default or inherit', 'hide' or 'show'.
	SmForums                           string `json:"smForums,omitempty"`                           //       The 'SM Forums' subtab which links to community-based support for Cisco Meraki Systems Manager. Only configurable for       organizations that contain Systems Manager networks. Can be one of 'default or inherit', 'hide' or 'show'.
	SupportContactInfo                 string `json:"supportContactInfo,omitempty"`                 //       The 'Contact Meraki Support' section of the 'Help -> Get Help' subtab. Can be one of 'default or inherit', 'hide', 'show', or a replacement custom HTML string.
	UniversalSearchKnowledgeBaseSearch string `json:"universalSearchKnowledgeBaseSearch,omitempty"` //       The universal search box always visible on Dashboard will, by default, present results from the Meraki KB. This configures       whether these Meraki KB results should be returned. Can be one of 'default or inherit', 'hide' or 'show'.
}
type RequestOrganizationsClaimIntoOrganization struct {
	Licenses *[]RequestOrganizationsClaimIntoOrganizationLicenses `json:"licenses,omitempty"` // The licenses that should be claimed
	Orders   []string                                             `json:"orders,omitempty"`   // The numbers of the orders that should be claimed
	Serials  []string                                             `json:"serials,omitempty"`  // The serials of the devices that should be claimed
}
type RequestOrganizationsClaimIntoOrganizationLicenses struct {
	Key  string `json:"key,omitempty"`  // The key of the license
	Mode string `json:"mode,omitempty"` // Either 'renew' or 'addDevices'. 'addDevices' will increase the license limit, while 'renew' will extend the amount of time until expiration. Defaults to 'addDevices'. All licenses must be claimed with the same mode, and at most one renewal can be claimed at a time. This parameter is legacy and does not apply to organizations with per-device licensing enabled.
}
type RequestOrganizationsCloneOrganization struct {
	Name string `json:"name,omitempty"` // The name of the new organization
}
type RequestOrganizationsCreateOrganizationConfigTemplate struct {
	Name     string `json:"name,omitempty"`     // The name of the configuration template
	TimeZone string `json:"timeZone,omitempty"` // The timezone of the configuration template. For a list of allowed timezones, please see the 'TZ' column in the table in <a target='_blank' href='https://en.wikipedia.org/wiki/List_of_tz_database_time_zones'>this article</a>. Not applicable if copying from existing network or template
}
type RequestOrganizationsUpdateOrganizationConfigTemplate struct {
	Name     string `json:"name,omitempty"`     // The name of the configuration template
	TimeZone string `json:"timeZone,omitempty"` // The timezone of the configuration template. For a list of allowed timezones, please see the 'TZ' column in the table in <a target='_blank' href='https://en.wikipedia.org/wiki/List_of_tz_database_time_zones'>this article.</a>
}
type RequestOrganizationsCreateOrganizationEarlyAccessFeaturesOptIn struct {
	LimitScopeToNetworks []string `json:"limitScopeToNetworks,omitempty"` // A list of network IDs to apply the opt-in to
	ShortName            string   `json:"shortName,omitempty"`            // Short name of the early access feature
}
type RequestOrganizationsUpdateOrganizationEarlyAccessFeaturesOptIn struct {
	LimitScopeToNetworks []string `json:"limitScopeToNetworks,omitempty"` // A list of network IDs to apply the opt-in to
}
type RequestOrganizationsClaimIntoOrganizationInventory struct {
	Licenses *[]RequestOrganizationsClaimIntoOrganizationInventoryLicenses `json:"licenses,omitempty"` // The licenses that should be claimed
	Orders   []string                                                      `json:"orders,omitempty"`   // The numbers of the orders that should be claimed
	Serials  []string                                                      `json:"serials,omitempty"`  // The serials of the devices that should be claimed
}
type RequestOrganizationsClaimIntoOrganizationInventoryLicenses struct {
	Key  string `json:"key,omitempty"`  // The key of the license
	Mode string `json:"mode,omitempty"` // Co-term licensing only: either 'renew' or 'addDevices'. 'addDevices' will increase the license limit, while 'renew' will extend the amount of time until expiration. Defaults to 'addDevices'. All licenses must be claimed with the same mode, and at most one renewal can be claimed at a time. Does not apply to organizations using per-device licensing model.
}
type RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringExportEvent struct {
	LogEvent  string `json:"logEvent,omitempty"`  // The type of log event this is recording, e.g. download or opening a banner
	Request   string `json:"request,omitempty"`   // Used to describe if this event was the result of a redirect. E.g. a query param if an info banner is being used
	TargetOS  string `json:"targetOS,omitempty"`  // The name of the onboarding distro being downloaded
	Timestamp *int   `json:"timestamp,omitempty"` // A JavaScript UTC datetime stamp for when the even occurred
}
type RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringImport struct {
	Devices *[]RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringImportDevices `json:"devices,omitempty"` // A set of device imports to commit
}
type RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringImportDevices struct {
	DeviceID  string `json:"deviceId,omitempty"`  // Import ID from the Import operation
	NetworkID string `json:"networkId,omitempty"` // Network Id
	Udi       string `json:"udi,omitempty"`       // Device UDI certificate
}
type RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepare struct {
	Devices *[]RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevices `json:"devices,omitempty"` // A set of devices to import (or update)
}
type RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevices struct {
	Sudi   string                                                                                        `json:"sudi,omitempty"`   // Device SUDI certificate
	Tunnel *RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesTunnel `json:"tunnel,omitempty"` // TLS Related Parameters
	User   *RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesUser   `json:"user,omitempty"`   // User parameters
	Vty    *RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVty    `json:"vty,omitempty"`    // VTY Related Parameters
}
type RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesTunnel struct {
	CertificateName string `json:"certificateName,omitempty"` // Name of the configured TLS certificate
	LocalInterface  *int   `json:"localInterface,omitempty"`  // Number of the vlan expected to be used to connect to the cloud
	LoopbackNumber  *int   `json:"loopbackNumber,omitempty"`  // Number of the configured Loopback Interface used for TLS overlay
	Name            string `json:"name,omitempty"`            // Name of the configured TLS tunnel
}
type RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesUser struct {
	Username string `json:"username,omitempty"` // The name of the device user for Meraki monitoring
}
type RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVty struct {
	AccessList      *RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAccessList     `json:"accessList,omitempty"`      // AccessList details
	Authentication  *RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAuthentication `json:"authentication,omitempty"`  // VTY AAA authentication
	Authorization   *RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAuthorization  `json:"authorization,omitempty"`   // VTY AAA authorization
	EndLineNumber   *int                                                                                                     `json:"endLineNumber,omitempty"`   // Ending line VTY number
	RotaryNumber    *int                                                                                                     `json:"rotaryNumber,omitempty"`    // SSH rotary number
	StartLineNumber *int                                                                                                     `json:"startLineNumber,omitempty"` // Starting line VTY number
}
type RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAccessList struct {
	VtyIn  *RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAccessListVtyIn  `json:"vtyIn,omitempty"`  // VTY in ACL
	VtyOut *RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAccessListVtyOut `json:"vtyOut,omitempty"` // VTY out ACL
}
type RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAccessListVtyIn struct {
	Name string `json:"name,omitempty"` // Name
}
type RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAccessListVtyOut struct {
	Name string `json:"name,omitempty"` // Name
}
type RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAuthentication struct {
	Group *RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAuthenticationGroup `json:"group,omitempty"` // Group Details
}
type RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAuthenticationGroup struct {
	Name string `json:"name,omitempty"` // Group Name
}
type RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAuthorization struct {
	Group *RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAuthorizationGroup `json:"group,omitempty"` // Group Details
}
type RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAuthorizationGroup struct {
	Name string `json:"name,omitempty"` // Group Name
}
type RequestOrganizationsReleaseFromOrganizationInventory struct {
	Serials []string `json:"serials,omitempty"` // Serials of the devices that should be released
}
type RequestOrganizationsAssignOrganizationLicensesSeats struct {
	LicenseID string `json:"licenseId,omitempty"` // The ID of the SM license to assign seats from
	NetworkID string `json:"networkId,omitempty"` // The ID of the SM network to assign the seats to
	SeatCount *int   `json:"seatCount,omitempty"` // The number of seats to assign to the SM network. Must be less than or equal to the total number of seats of the license
}
type RequestOrganizationsMoveOrganizationLicenses struct {
	DestOrganizationID string   `json:"destOrganizationId,omitempty"` // The ID of the organization to move the licenses to
	LicenseIDs         []string `json:"licenseIds,omitempty"`         // A list of IDs of licenses to move to the new organization
}
type RequestOrganizationsMoveOrganizationLicensesSeats struct {
	DestOrganizationID string `json:"destOrganizationId,omitempty"` // The ID of the organization to move the SM seats to
	LicenseID          string `json:"licenseId,omitempty"`          // The ID of the SM license to move the seats from
	SeatCount          *int   `json:"seatCount,omitempty"`          // The number of seats to move to the new organization. Must be less than or equal to the total number of seats of the license
}
type RequestOrganizationsRenewOrganizationLicensesSeats struct {
	LicenseIDToRenew string `json:"licenseIdToRenew,omitempty"` // The ID of the SM license to renew. This license must already be assigned to an SM network
	UnusedLicenseID  string `json:"unusedLicenseId,omitempty"`  // The SM license to use to renew the seats on 'licenseIdToRenew'. This license must have at least as many seats available as there are seats on 'licenseIdToRenew'
}
type RequestOrganizationsUpdateOrganizationLicense struct {
	DeviceSerial string `json:"deviceSerial,omitempty"` // The serial number of the device to assign this license to. Set this to  null to unassign the license. If a different license is already active on the device, this parameter will control queueing/dequeuing this license.
}
type RequestOrganizationsUpdateOrganizationLoginSecurity struct {
	AccountLockoutAttempts    *int                                                                  `json:"accountLockoutAttempts,omitempty"`    // Number of consecutive failed login attempts after which users' accounts will be locked.
	APIAuthentication         *RequestOrganizationsUpdateOrganizationLoginSecurityAPIAuthentication `json:"apiAuthentication,omitempty"`         // Details for indicating whether organization will restrict access to API (but not Dashboard) to certain IP addresses.
	EnforceAccountLockout     *bool                                                                 `json:"enforceAccountLockout,omitempty"`     // Boolean indicating whether users' Dashboard accounts will be locked out after a specified number of consecutive failed login attempts.
	EnforceDifferentPasswords *bool                                                                 `json:"enforceDifferentPasswords,omitempty"` // Boolean indicating whether users, when setting a new password, are forced to choose a new password that is different from any past passwords.
	EnforceIDleTimeout        *bool                                                                 `json:"enforceIdleTimeout,omitempty"`        // Boolean indicating whether users will be logged out after being idle for the specified number of minutes.
	EnforceLoginIPRanges      *bool                                                                 `json:"enforceLoginIpRanges,omitempty"`      // Boolean indicating whether organization will restrict access to Dashboard (including the API) from certain IP addresses.
	EnforcePasswordExpiration *bool                                                                 `json:"enforcePasswordExpiration,omitempty"` // Boolean indicating whether users are forced to change their password every X number of days.
	EnforceStrongPasswords    *bool                                                                 `json:"enforceStrongPasswords,omitempty"`    // Boolean indicating whether users will be forced to choose strong passwords for their accounts. Strong passwords are at least 8 characters that contain 3 of the following: number, uppercase letter, lowercase letter, and symbol
	EnforceTwoFactorAuth      *bool                                                                 `json:"enforceTwoFactorAuth,omitempty"`      // Boolean indicating whether users in this organization will be required to use an extra verification code when logging in to Dashboard. This code will be sent to their mobile phone via SMS, or can be generated by the authenticator application.
	IDleTimeoutMinutes        *int                                                                  `json:"idleTimeoutMinutes,omitempty"`        // Number of minutes users can remain idle before being logged out of their accounts.
	LoginIPRanges             []string                                                              `json:"loginIpRanges,omitempty"`             // List of acceptable IP ranges. Entries can be single IP addresses, IP address ranges, and CIDR subnets.
	NumDifferentPasswords     *int                                                                  `json:"numDifferentPasswords,omitempty"`     // Number of recent passwords that new password must be distinct from.
	PasswordExpirationDays    *int                                                                  `json:"passwordExpirationDays,omitempty"`    // Number of days after which users will be forced to change their password.
}
type RequestOrganizationsUpdateOrganizationLoginSecurityAPIAuthentication struct {
	IPRestrictionsForKeys *RequestOrganizationsUpdateOrganizationLoginSecurityAPIAuthenticationIPRestrictionsForKeys `json:"ipRestrictionsForKeys,omitempty"` // Details for API-only IP restrictions.
}
type RequestOrganizationsUpdateOrganizationLoginSecurityAPIAuthenticationIPRestrictionsForKeys struct {
	Enabled *bool    `json:"enabled,omitempty"` // Boolean indicating whether the organization will restrict API key (not Dashboard GUI) usage to a specific list of IP addresses or CIDR ranges.
	Ranges  []string `json:"ranges,omitempty"`  // List of acceptable IP ranges. Entries can be single IP addresses, IP address ranges, and CIDR subnets.
}
type RequestOrganizationsCreateOrganizationNetwork struct {
	CopyFromNetworkID string   `json:"copyFromNetworkId,omitempty"` // The ID of the network to copy configuration from. Other provided parameters will override the copied configuration, except type which must match this network's type exactly.
	Name              string   `json:"name,omitempty"`              // The name of the new network
	Notes             string   `json:"notes,omitempty"`             // Add any notes or additional information about this network here.
	ProductTypes      []string `json:"productTypes,omitempty"`      // The product type(s) of the new network. If more than one type is included, the network will be a combined network.
	Tags              []string `json:"tags,omitempty"`              // A list of tags to be applied to the network
	TimeZone          string   `json:"timeZone,omitempty"`          // The timezone of the network. For a list of allowed timezones, please see the 'TZ' column in the table in <a target='_blank' href='https://en.wikipedia.org/wiki/List_of_tz_database_time_zones'>this article.</a>
}
type RequestOrganizationsCombineOrganizationNetworks struct {
	EnrollmentString string   `json:"enrollmentString,omitempty"` // A unique identifier which can be used for device enrollment or easy access through the Meraki SM Registration page or the Self Service Portal. Please note that changing this field may cause existing bookmarks to break. All networks that are part of this combined network will have their enrollment string appended by '-network_type'. If left empty, all exisitng enrollment strings will be deleted.
	Name             string   `json:"name,omitempty"`             // The name of the combined network
	NetworkIDs       []string `json:"networkIds,omitempty"`       // A list of the network IDs that will be combined. If an ID of a combined network is included in this list, the other networks in the list will be grouped into that network
}
type RequestOrganizationsCreateOrganizationPolicyObject struct {
	Category string `json:"category,omitempty"` // Category of a policy object (one of: adaptivePolicy, network)
	Cidr     string `json:"cidr,omitempty"`     // CIDR Value of a policy object (e.g. 10.11.12.1/24")
	Fqdn     string `json:"fqdn,omitempty"`     // Fully qualified domain name of policy object (e.g. "example.com")
	GroupIDs *[]int `json:"groupIds,omitempty"` // The IDs of policy object groups the policy object belongs to
	IP       string `json:"ip,omitempty"`       // IP Address of a policy object (e.g. "1.2.3.4")
	Mask     string `json:"mask,omitempty"`     // Mask of a policy object (e.g. "255.255.0.0")
	Name     string `json:"name,omitempty"`     // Name of a policy object, unique within the organization (alphanumeric, space, dash, or underscore characters only)
	Type     string `json:"type,omitempty"`     // Type of a policy object (one of: adaptivePolicyIpv4Cidr, cidr, fqdn, ipAndMask)
}
type RequestOrganizationsCreateOrganizationPolicyObjectsGroup struct {
	Category  string `json:"category,omitempty"`  // Category of a policy object group (one of: NetworkObjectGroup, GeoLocationGroup, PortObjectGroup, ApplicationGroup)
	Name      string `json:"name,omitempty"`      // A name for the group of network addresses, unique within the organization (alphanumeric, space, dash, or underscore characters only)
	ObjectIDs *[]int `json:"objectIds,omitempty"` // A list of Policy Object ID's that this NetworkObjectGroup should be associated to (note: these ID's will replace the existing associated Policy Objects)
}
type RequestOrganizationsUpdateOrganizationPolicyObjectsGroup struct {
	Name      string `json:"name,omitempty"`      // A name for the group of network addresses, unique within the organization (alphanumeric, space, dash, or underscore characters only)
	ObjectIDs *[]int `json:"objectIds,omitempty"` // A list of Policy Object ID's that this NetworkObjectGroup should be associated to (note: these ID's will replace the existing associated Policy Objects)
}
type RequestOrganizationsUpdateOrganizationPolicyObject struct {
	Cidr     string `json:"cidr,omitempty"`     // CIDR Value of a policy object (e.g. 10.11.12.1/24")
	Fqdn     string `json:"fqdn,omitempty"`     // Fully qualified domain name of policy object (e.g. "example.com")
	GroupIDs *[]int `json:"groupIds,omitempty"` // The IDs of policy object groups the policy object belongs to
	IP       string `json:"ip,omitempty"`       // IP Address of a policy object (e.g. "1.2.3.4")
	Mask     string `json:"mask,omitempty"`     // Mask of a policy object (e.g. "255.255.0.0")
	Name     string `json:"name,omitempty"`     // Name of a policy object, unique within the organization (alphanumeric, space, dash, or underscore characters only)
}
type RequestOrganizationsUpdateOrganizationSaml struct {
	Enabled *bool `json:"enabled,omitempty"` // Boolean for updating SAML SSO enabled settings.
}
type RequestOrganizationsCreateOrganizationSamlIDp struct {
	SloLogoutURL            string `json:"sloLogoutUrl,omitempty"`            // Dashboard will redirect users to this URL when they sign out.
	X509CertSha1Fingerprint string `json:"x509certSha1Fingerprint,omitempty"` // Fingerprint (SHA1) of the SAML certificate provided by your Identity Provider (IdP). This will be used for encryption / validation.
}
type RequestOrganizationsUpdateOrganizationSamlIDp struct {
	SloLogoutURL            string `json:"sloLogoutUrl,omitempty"`            // Dashboard will redirect users to this URL when they sign out.
	X509CertSha1Fingerprint string `json:"x509certSha1Fingerprint,omitempty"` // Fingerprint (SHA1) of the SAML certificate provided by your Identity Provider (IdP). This will be used for encryption / validation.
}
type RequestOrganizationsCreateOrganizationSamlRole struct {
	Networks  *[]RequestOrganizationsCreateOrganizationSamlRoleNetworks `json:"networks,omitempty"`  // The list of networks that the SAML administrator has privileges on
	OrgAccess string                                                    `json:"orgAccess,omitempty"` // The privilege of the SAML administrator on the organization. Can be one of 'none', 'read-only', 'full' or 'enterprise'
	Role      string                                                    `json:"role,omitempty"`      // The role of the SAML administrator
	Tags      *[]RequestOrganizationsCreateOrganizationSamlRoleTags     `json:"tags,omitempty"`      // The list of tags that the SAML administrator has privleges on
}
type RequestOrganizationsCreateOrganizationSamlRoleNetworks struct {
	Access string `json:"access,omitempty"` // The privilege of the SAML administrator on the network. Can be one of 'full', 'read-only', 'guest-ambassador', 'monitor-only' or 'ssid-admin'
	ID     string `json:"id,omitempty"`     // The network ID
}
type RequestOrganizationsCreateOrganizationSamlRoleTags struct {
	Access string `json:"access,omitempty"` // The privilege of the SAML administrator on the tag. Can be one of 'full', 'read-only', 'guest-ambassador' or 'monitor-only'
	Tag    string `json:"tag,omitempty"`    // The name of the tag
}
type RequestOrganizationsUpdateOrganizationSamlRole struct {
	Networks  *[]RequestOrganizationsUpdateOrganizationSamlRoleNetworks `json:"networks,omitempty"`  // The list of networks that the SAML administrator has privileges on
	OrgAccess string                                                    `json:"orgAccess,omitempty"` // The privilege of the SAML administrator on the organization. Can be one of 'none', 'read-only', 'full' or 'enterprise'
	Role      string                                                    `json:"role,omitempty"`      // The role of the SAML administrator
	Tags      *[]RequestOrganizationsUpdateOrganizationSamlRoleTags     `json:"tags,omitempty"`      // The list of tags that the SAML administrator has privleges on
}
type RequestOrganizationsUpdateOrganizationSamlRoleNetworks struct {
	Access string `json:"access,omitempty"` // The privilege of the SAML administrator on the network. Can be one of 'full', 'read-only', 'guest-ambassador', 'monitor-only' or 'ssid-admin'
	ID     string `json:"id,omitempty"`     // The network ID
}
type RequestOrganizationsUpdateOrganizationSamlRoleTags struct {
	Access string `json:"access,omitempty"` // The privilege of the SAML administrator on the tag. Can be one of 'full', 'read-only', 'guest-ambassador' or 'monitor-only'
	Tag    string `json:"tag,omitempty"`    // The name of the tag
}
type RequestOrganizationsUpdateOrganizationSNMP struct {
	PeerIPs    []string `json:"peerIps,omitempty"`    // The list of IPv4 addresses that are allowed to access the SNMP server.
	V2CEnabled *bool    `json:"v2cEnabled,omitempty"` // Boolean indicating whether SNMP version 2c is enabled for the organization.
	V3AuthMode string   `json:"v3AuthMode,omitempty"` // The SNMP version 3 authentication mode. Can be either 'MD5' or 'SHA'.
	V3AuthPass string   `json:"v3AuthPass,omitempty"` // The SNMP version 3 authentication password. Must be at least 8 characters if specified.
	V3Enabled  *bool    `json:"v3Enabled,omitempty"`  // Boolean indicating whether SNMP version 3 is enabled for the organization.
	V3PrivMode string   `json:"v3PrivMode,omitempty"` // The SNMP version 3 privacy mode. Can be either 'DES' or 'AES128'.
	V3PrivPass string   `json:"v3PrivPass,omitempty"` // The SNMP version 3 privacy password. Must be at least 8 characters if specified.
}

//GetOrganizations List the organizations that the user has privileges on
/* List the organizations that the user has privileges on


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organizations
*/
func (s *OrganizationsService) GetOrganizations() (*ResponseOrganizationsGetOrganizations, *resty.Response, error) {
	path := "/api/v1/organizations"
	s.rateLimiterBucket.Wait(1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseOrganizationsGetOrganizations{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizations")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizations)
	return result, response, err

}

//GetOrganization Return an organization
/* Return an organization

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization
*/
func (s *OrganizationsService) GetOrganization(organizationID string) (*ResponseOrganizationsGetOrganization, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseOrganizationsGetOrganization{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganization")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganization)
	return result, response, err

}

//GetOrganizationActionBatches Return the list of action batches in the organization
/* Return the list of action batches in the organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationActionBatchesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-action-batches
*/
func (s *OrganizationsService) GetOrganizationActionBatches(organizationID string, getOrganizationActionBatchesQueryParams *GetOrganizationActionBatchesQueryParams) (*ResponseOrganizationsGetOrganizationActionBatches, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/actionBatches"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationActionBatchesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseOrganizationsGetOrganizationActionBatches{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationActionBatches")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationActionBatches)
	return result, response, err

}

//GetOrganizationActionBatch Return an action batch
/* Return an action batch

@param organizationID organizationId path parameter. Organization ID
@param actionBatchID actionBatchId path parameter. Action batch ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-action-batch
*/
func (s *OrganizationsService) GetOrganizationActionBatch(organizationID string, actionBatchID string) (*ResponseOrganizationsGetOrganizationActionBatch, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/actionBatches/{actionBatchId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{actionBatchId}", fmt.Sprintf("%v", actionBatchID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseOrganizationsGetOrganizationActionBatch{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationActionBatch")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationActionBatch)
	return result, response, err

}

//GetOrganizationAdaptivePolicyACLs List adaptive policy ACLs in a organization
/* List adaptive policy ACLs in a organization

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-adaptive-policy-acls
*/
func (s *OrganizationsService) GetOrganizationAdaptivePolicyACLs(organizationID string) (*ResponseOrganizationsGetOrganizationAdaptivePolicyACLs, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/adaptivePolicy/acls"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseOrganizationsGetOrganizationAdaptivePolicyACLs{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationAdaptivePolicyAcls")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationAdaptivePolicyACLs)
	return result, response, err

}

//GetOrganizationAdaptivePolicyACL Returns the adaptive policy ACL information
/* Returns the adaptive policy ACL information

@param organizationID organizationId path parameter. Organization ID
@param aclID aclId path parameter. Acl ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-adaptive-policy-acl
*/
func (s *OrganizationsService) GetOrganizationAdaptivePolicyACL(organizationID string, aclID string) (*ResponseOrganizationsGetOrganizationAdaptivePolicyACL, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/adaptivePolicy/acls/{aclId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{aclId}", fmt.Sprintf("%v", aclID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseOrganizationsGetOrganizationAdaptivePolicyACL{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationAdaptivePolicyAcl")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationAdaptivePolicyACL)
	return result, response, err

}

//GetOrganizationAdaptivePolicyGroups List adaptive policy groups in a organization
/* List adaptive policy groups in a organization

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-adaptive-policy-groups
*/
func (s *OrganizationsService) GetOrganizationAdaptivePolicyGroups(organizationID string) (*ResponseOrganizationsGetOrganizationAdaptivePolicyGroups, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/adaptivePolicy/groups"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseOrganizationsGetOrganizationAdaptivePolicyGroups{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationAdaptivePolicyGroups")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationAdaptivePolicyGroups)
	return result, response, err

}

//GetOrganizationAdaptivePolicyGroup Returns an adaptive policy group
/* Returns an adaptive policy group

@param organizationID organizationId path parameter. Organization ID
@param id id path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-adaptive-policy-group
*/
func (s *OrganizationsService) GetOrganizationAdaptivePolicyGroup(organizationID string, id string) (*ResponseOrganizationsGetOrganizationAdaptivePolicyGroup, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/adaptivePolicy/groups/{id}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseOrganizationsGetOrganizationAdaptivePolicyGroup{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationAdaptivePolicyGroup")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationAdaptivePolicyGroup)
	return result, response, err

}

//GetOrganizationAdaptivePolicyOverview Returns adaptive policy aggregate statistics for an organization
/* Returns adaptive policy aggregate statistics for an organization

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-adaptive-policy-overview
*/
func (s *OrganizationsService) GetOrganizationAdaptivePolicyOverview(organizationID string) (*ResponseOrganizationsGetOrganizationAdaptivePolicyOverview, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/adaptivePolicy/overview"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseOrganizationsGetOrganizationAdaptivePolicyOverview{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationAdaptivePolicyOverview")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationAdaptivePolicyOverview)
	return result, response, err

}

//GetOrganizationAdaptivePolicyPolicies List adaptive policies in an organization
/* List adaptive policies in an organization

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-adaptive-policy-policies
*/
func (s *OrganizationsService) GetOrganizationAdaptivePolicyPolicies(organizationID string) (*ResponseOrganizationsGetOrganizationAdaptivePolicyPolicies, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/adaptivePolicy/policies"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseOrganizationsGetOrganizationAdaptivePolicyPolicies{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationAdaptivePolicyPolicies")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationAdaptivePolicyPolicies)
	return result, response, err

}

//GetOrganizationAdaptivePolicyPolicy Return an adaptive policy
/* Return an adaptive policy

@param organizationID organizationId path parameter. Organization ID
@param id id path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-adaptive-policy-policy
*/
func (s *OrganizationsService) GetOrganizationAdaptivePolicyPolicy(organizationID string, id string) (*ResponseOrganizationsGetOrganizationAdaptivePolicyPolicy, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/adaptivePolicy/policies/{id}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseOrganizationsGetOrganizationAdaptivePolicyPolicy{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationAdaptivePolicyPolicy")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationAdaptivePolicyPolicy)
	return result, response, err

}

//GetOrganizationAdaptivePolicySettings Returns global adaptive policy settings in an organization
/* Returns global adaptive policy settings in an organization

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-adaptive-policy-settings
*/
func (s *OrganizationsService) GetOrganizationAdaptivePolicySettings(organizationID string) (*ResponseOrganizationsGetOrganizationAdaptivePolicySettings, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/adaptivePolicy/settings"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseOrganizationsGetOrganizationAdaptivePolicySettings{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationAdaptivePolicySettings")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationAdaptivePolicySettings)
	return result, response, err

}

//GetOrganizationAdmins List the dashboard administrators in this organization
/* List the dashboard administrators in this organization

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-admins
*/
func (s *OrganizationsService) GetOrganizationAdmins(organizationID string) (*ResponseOrganizationsGetOrganizationAdmins, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/admins"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseOrganizationsGetOrganizationAdmins{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationAdmins")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationAdmins)
	return result, response, err

}

//GetOrganizationAlertsProfiles List all organization-wide alert configurations
/* List all organization-wide alert configurations

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-alerts-profiles
*/
func (s *OrganizationsService) GetOrganizationAlertsProfiles(organizationID string) (*ResponseOrganizationsGetOrganizationAlertsProfiles, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/alerts/profiles"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseOrganizationsGetOrganizationAlertsProfiles{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationAlertsProfiles")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationAlertsProfiles)
	return result, response, err

}

//GetOrganizationAPIRequests List the API requests made by an organization
/* List the API requests made by an organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationApiRequestsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-api-requests
*/
func (s *OrganizationsService) GetOrganizationAPIRequests(organizationID string, getOrganizationApiRequestsQueryParams *GetOrganizationAPIRequestsQueryParams) (*ResponseOrganizationsGetOrganizationAPIRequests, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/apiRequests"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationApiRequestsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseOrganizationsGetOrganizationAPIRequests{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationApiRequests")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationAPIRequests)
	return result, response, err

}

//GetOrganizationAPIRequestsOverview Return an aggregated overview of API requests data
/* Return an aggregated overview of API requests data

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationApiRequestsOverviewQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-api-requests-overview
*/
func (s *OrganizationsService) GetOrganizationAPIRequestsOverview(organizationID string, getOrganizationApiRequestsOverviewQueryParams *GetOrganizationAPIRequestsOverviewQueryParams) (*ResponseOrganizationsGetOrganizationAPIRequestsOverview, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/apiRequests/overview"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationApiRequestsOverviewQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseOrganizationsGetOrganizationAPIRequestsOverview{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationApiRequestsOverview")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationAPIRequestsOverview)
	return result, response, err

}

//GetOrganizationAPIRequestsOverviewResponseCodesByInterval Tracks organizations' API requests by response code across a given time period
/* Tracks organizations' API requests by response code across a given time period

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationApiRequestsOverviewResponseCodesByIntervalQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-api-requests-overview-response-codes-by-interval
*/
func (s *OrganizationsService) GetOrganizationAPIRequestsOverviewResponseCodesByInterval(organizationID string, getOrganizationApiRequestsOverviewResponseCodesByIntervalQueryParams *GetOrganizationAPIRequestsOverviewResponseCodesByIntervalQueryParams) (*ResponseOrganizationsGetOrganizationAPIRequestsOverviewResponseCodesByInterval, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/apiRequests/overview/responseCodes/byInterval"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationApiRequestsOverviewResponseCodesByIntervalQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseOrganizationsGetOrganizationAPIRequestsOverviewResponseCodesByInterval{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationApiRequestsOverviewResponseCodesByInterval")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationAPIRequestsOverviewResponseCodesByInterval)
	return result, response, err

}

//GetOrganizationBrandingPolicies List the branding policies of an organization
/* List the branding policies of an organization

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-branding-policies
*/
func (s *OrganizationsService) GetOrganizationBrandingPolicies(organizationID string) (*ResponseOrganizationsGetOrganizationBrandingPolicies, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/brandingPolicies"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseOrganizationsGetOrganizationBrandingPolicies{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationBrandingPolicies")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationBrandingPolicies)
	return result, response, err

}

//GetOrganizationBrandingPoliciesPriorities Return the branding policy IDs of an organization in priority order
/* Return the branding policy IDs of an organization in priority order. IDs are ordered in ascending order of priority (IDs later in the array have higher priority).

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-branding-policies-priorities
*/
func (s *OrganizationsService) GetOrganizationBrandingPoliciesPriorities(organizationID string) (*ResponseOrganizationsGetOrganizationBrandingPoliciesPriorities, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/brandingPolicies/priorities"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseOrganizationsGetOrganizationBrandingPoliciesPriorities{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationBrandingPoliciesPriorities")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationBrandingPoliciesPriorities)
	return result, response, err

}

//GetOrganizationBrandingPolicy Return a branding policy
/* Return a branding policy

@param organizationID organizationId path parameter. Organization ID
@param brandingPolicyID brandingPolicyId path parameter. Branding policy ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-branding-policy
*/
func (s *OrganizationsService) GetOrganizationBrandingPolicy(organizationID string, brandingPolicyID string) (*ResponseOrganizationsGetOrganizationBrandingPolicy, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/brandingPolicies/{brandingPolicyId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{brandingPolicyId}", fmt.Sprintf("%v", brandingPolicyID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseOrganizationsGetOrganizationBrandingPolicy{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationBrandingPolicy")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationBrandingPolicy)
	return result, response, err

}

//GetOrganizationClientsBandwidthUsageHistory Return data usage (in megabits per second) over time for all clients in the given organization within a given time range.
/* Return data usage (in megabits per second) over time for all clients in the given organization within a given time range.

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationClientsBandwidthUsageHistoryQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-clients-bandwidth-usage-history
*/
func (s *OrganizationsService) GetOrganizationClientsBandwidthUsageHistory(organizationID string, getOrganizationClientsBandwidthUsageHistoryQueryParams *GetOrganizationClientsBandwidthUsageHistoryQueryParams) (*ResponseOrganizationsGetOrganizationClientsBandwidthUsageHistory, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/clients/bandwidthUsageHistory"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationClientsBandwidthUsageHistoryQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseOrganizationsGetOrganizationClientsBandwidthUsageHistory{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationClientsBandwidthUsageHistory")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationClientsBandwidthUsageHistory)
	return result, response, err

}

//GetOrganizationClientsOverview Return summary information around client data usage (in mb) across the given organization.
/* Return summary information around client data usage (in mb) across the given organization.

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationClientsOverviewQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-clients-overview
*/
func (s *OrganizationsService) GetOrganizationClientsOverview(organizationID string, getOrganizationClientsOverviewQueryParams *GetOrganizationClientsOverviewQueryParams) (*ResponseOrganizationsGetOrganizationClientsOverview, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/clients/overview"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationClientsOverviewQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseOrganizationsGetOrganizationClientsOverview{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationClientsOverview")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationClientsOverview)
	return result, response, err

}

//GetOrganizationClientsSearch Return the client details in an organization
/* Return the client details in an organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationClientsSearchQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-clients-search
*/
func (s *OrganizationsService) GetOrganizationClientsSearch(organizationID string, getOrganizationClientsSearchQueryParams *GetOrganizationClientsSearchQueryParams) (*ResponseOrganizationsGetOrganizationClientsSearch, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/clients/search"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationClientsSearchQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseOrganizationsGetOrganizationClientsSearch{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationClientsSearch")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationClientsSearch)
	return result, response, err

}

//GetOrganizationConfigTemplates List the configuration templates for this organization
/* List the configuration templates for this organization

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-config-templates
*/
func (s *OrganizationsService) GetOrganizationConfigTemplates(organizationID string) (*ResponseOrganizationsGetOrganizationConfigTemplates, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/configTemplates"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseOrganizationsGetOrganizationConfigTemplates{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationConfigTemplates")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationConfigTemplates)
	return result, response, err

}

//GetOrganizationConfigTemplate Return a single configuration template
/* Return a single configuration template

@param organizationID organizationId path parameter. Organization ID
@param configTemplateID configTemplateId path parameter. Config template ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-config-template
*/
func (s *OrganizationsService) GetOrganizationConfigTemplate(organizationID string, configTemplateID string) (*ResponseOrganizationsGetOrganizationConfigTemplate, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/configTemplates/{configTemplateId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{configTemplateId}", fmt.Sprintf("%v", configTemplateID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseOrganizationsGetOrganizationConfigTemplate{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationConfigTemplate")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationConfigTemplate)
	return result, response, err

}

//GetOrganizationConfigurationChanges View the Change Log for your organization
/* View the Change Log for your organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationConfigurationChangesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-configuration-changes
*/
func (s *OrganizationsService) GetOrganizationConfigurationChanges(organizationID string, getOrganizationConfigurationChangesQueryParams *GetOrganizationConfigurationChangesQueryParams) (*ResponseOrganizationsGetOrganizationConfigurationChanges, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/configurationChanges"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationConfigurationChangesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseOrganizationsGetOrganizationConfigurationChanges{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationConfigurationChanges")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationConfigurationChanges)
	return result, response, err

}

//GetOrganizationDevices List the devices in an organization
/* List the devices in an organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationDevicesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-devices
*/
func (s *OrganizationsService) GetOrganizationDevices(organizationID string, getOrganizationDevicesQueryParams *GetOrganizationDevicesQueryParams) (*ResponseOrganizationsGetOrganizationDevices, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/devices"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationDevicesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseOrganizationsGetOrganizationDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationDevices")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationDevices)
	return result, response, err

}

//GetOrganizationDevicesAvailabilities List the availability information for devices in an organization
/* List the availability information for devices in an organization. The data returned by this endpoint is updated every 5 minutes.

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationDevicesAvailabilitiesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-devices-availabilities
*/
func (s *OrganizationsService) GetOrganizationDevicesAvailabilities(organizationID string, getOrganizationDevicesAvailabilitiesQueryParams *GetOrganizationDevicesAvailabilitiesQueryParams) (*ResponseOrganizationsGetOrganizationDevicesAvailabilities, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/devices/availabilities"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationDevicesAvailabilitiesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseOrganizationsGetOrganizationDevicesAvailabilities{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationDevicesAvailabilities")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationDevicesAvailabilities)
	return result, response, err

}

//GetOrganizationDevicesPowerModulesStatusesByDevice List the power status information for devices in an organization
/* List the power status information for devices in an organization. The data returned by this endpoint is updated every 5 minutes.

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationDevicesPowerModulesStatusesByDeviceQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-devices-power-modules-statuses-by-device
*/
func (s *OrganizationsService) GetOrganizationDevicesPowerModulesStatusesByDevice(organizationID string, getOrganizationDevicesPowerModulesStatusesByDeviceQueryParams *GetOrganizationDevicesPowerModulesStatusesByDeviceQueryParams) (*ResponseOrganizationsGetOrganizationDevicesPowerModulesStatusesByDevice, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/devices/powerModules/statuses/byDevice"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationDevicesPowerModulesStatusesByDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseOrganizationsGetOrganizationDevicesPowerModulesStatusesByDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationDevicesPowerModulesStatusesByDevice")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationDevicesPowerModulesStatusesByDevice)
	return result, response, err

}

//GetOrganizationDevicesProvisioningStatuses List the provisioning statuses information for devices in an organization.
/* List the provisioning statuses information for devices in an organization.

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationDevicesProvisioningStatusesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-devices-provisioning-statuses
*/
func (s *OrganizationsService) GetOrganizationDevicesProvisioningStatuses(organizationID string, getOrganizationDevicesProvisioningStatusesQueryParams *GetOrganizationDevicesProvisioningStatusesQueryParams) (*ResponseOrganizationsGetOrganizationDevicesProvisioningStatuses, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/devices/provisioning/statuses"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationDevicesProvisioningStatusesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseOrganizationsGetOrganizationDevicesProvisioningStatuses{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationDevicesProvisioningStatuses")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationDevicesProvisioningStatuses)
	return result, response, err

}

//GetOrganizationDevicesStatuses List the status of every Meraki device in the organization
/* List the status of every Meraki device in the organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationDevicesStatusesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-devices-statuses
*/
func (s *OrganizationsService) GetOrganizationDevicesStatuses(organizationID string, getOrganizationDevicesStatusesQueryParams *GetOrganizationDevicesStatusesQueryParams) (*ResponseOrganizationsGetOrganizationDevicesStatuses, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/devices/statuses"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationDevicesStatusesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseOrganizationsGetOrganizationDevicesStatuses{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationDevicesStatuses")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationDevicesStatuses)
	return result, response, err

}

//GetOrganizationDevicesStatusesOverview Return an overview of current device statuses
/* Return an overview of current device statuses

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationDevicesStatusesOverviewQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-devices-statuses-overview
*/
func (s *OrganizationsService) GetOrganizationDevicesStatusesOverview(organizationID string, getOrganizationDevicesStatusesOverviewQueryParams *GetOrganizationDevicesStatusesOverviewQueryParams) (*ResponseOrganizationsGetOrganizationDevicesStatusesOverview, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/devices/statuses/overview"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationDevicesStatusesOverviewQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseOrganizationsGetOrganizationDevicesStatusesOverview{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationDevicesStatusesOverview")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationDevicesStatusesOverview)
	return result, response, err

}

//GetOrganizationDevicesUplinksAddressesByDevice List the current uplink addresses for devices in an organization.
/* List the current uplink addresses for devices in an organization.

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationDevicesUplinksAddressesByDeviceQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-devices-uplinks-addresses-by-device
*/
func (s *OrganizationsService) GetOrganizationDevicesUplinksAddressesByDevice(organizationID string, getOrganizationDevicesUplinksAddressesByDeviceQueryParams *GetOrganizationDevicesUplinksAddressesByDeviceQueryParams) (*ResponseOrganizationsGetOrganizationDevicesUplinksAddressesByDevice, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/devices/uplinks/addresses/byDevice"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationDevicesUplinksAddressesByDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseOrganizationsGetOrganizationDevicesUplinksAddressesByDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationDevicesUplinksAddressesByDevice")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationDevicesUplinksAddressesByDevice)
	return result, response, err

}

//GetOrganizationDevicesUplinksLossAndLatency Return the uplink loss and latency for every MX in the organization from at latest 2 minutes ago
/* Return the uplink loss and latency for every MX in the organization from at latest 2 minutes ago

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationDevicesUplinksLossAndLatencyQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-devices-uplinks-loss-and-latency
*/
func (s *OrganizationsService) GetOrganizationDevicesUplinksLossAndLatency(organizationID string, getOrganizationDevicesUplinksLossAndLatencyQueryParams *GetOrganizationDevicesUplinksLossAndLatencyQueryParams) (*ResponseOrganizationsGetOrganizationDevicesUplinksLossAndLatency, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/devices/uplinksLossAndLatency"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationDevicesUplinksLossAndLatencyQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseOrganizationsGetOrganizationDevicesUplinksLossAndLatency{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationDevicesUplinksLossAndLatency")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationDevicesUplinksLossAndLatency)
	return result, response, err

}

//GetOrganizationEarlyAccessFeatures List the available early access features for organization
/* List the available early access features for organization

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-early-access-features
*/
func (s *OrganizationsService) GetOrganizationEarlyAccessFeatures(organizationID string) (*ResponseOrganizationsGetOrganizationEarlyAccessFeatures, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/earlyAccess/features"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseOrganizationsGetOrganizationEarlyAccessFeatures{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationEarlyAccessFeatures")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationEarlyAccessFeatures)
	return result, response, err

}

//GetOrganizationEarlyAccessFeaturesOptIns List the early access feature opt-ins for an organization
/* List the early access feature opt-ins for an organization

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-early-access-features-opt-ins
*/
func (s *OrganizationsService) GetOrganizationEarlyAccessFeaturesOptIns(organizationID string) (*ResponseOrganizationsGetOrganizationEarlyAccessFeaturesOptIns, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/earlyAccess/features/optIns"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseOrganizationsGetOrganizationEarlyAccessFeaturesOptIns{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationEarlyAccessFeaturesOptIns")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationEarlyAccessFeaturesOptIns)
	return result, response, err

}

//GetOrganizationEarlyAccessFeaturesOptIn Show an early access feature opt-in for an organization
/* Show an early access feature opt-in for an organization

@param organizationID organizationId path parameter. Organization ID
@param optInID optInId path parameter. Opt in ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-early-access-features-opt-in
*/
func (s *OrganizationsService) GetOrganizationEarlyAccessFeaturesOptIn(organizationID string, optInID string) (*ResponseOrganizationsGetOrganizationEarlyAccessFeaturesOptIn, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/earlyAccess/features/optIns/{optInId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{optInId}", fmt.Sprintf("%v", optInID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseOrganizationsGetOrganizationEarlyAccessFeaturesOptIn{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationEarlyAccessFeaturesOptIn")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationEarlyAccessFeaturesOptIn)
	return result, response, err

}

//GetOrganizationFirmwareUpgrades Get firmware upgrade information for an organization
/* Get firmware upgrade information for an organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationFirmwareUpgradesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-firmware-upgrades
*/
func (s *OrganizationsService) GetOrganizationFirmwareUpgrades(organizationID string, getOrganizationFirmwareUpgradesQueryParams *GetOrganizationFirmwareUpgradesQueryParams) (*ResponseOrganizationsGetOrganizationFirmwareUpgrades, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/firmware/upgrades"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationFirmwareUpgradesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseOrganizationsGetOrganizationFirmwareUpgrades{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationFirmwareUpgrades")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationFirmwareUpgrades)
	return result, response, err

}

//GetOrganizationFirmwareUpgradesByDevice Get firmware upgrade status for the filtered devices
/* Get firmware upgrade status for the filtered devices

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationFirmwareUpgradesByDeviceQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-firmware-upgrades-by-device
*/
func (s *OrganizationsService) GetOrganizationFirmwareUpgradesByDevice(organizationID string, getOrganizationFirmwareUpgradesByDeviceQueryParams *GetOrganizationFirmwareUpgradesByDeviceQueryParams) (*ResponseOrganizationsGetOrganizationFirmwareUpgradesByDevice, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/firmware/upgrades/byDevice"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationFirmwareUpgradesByDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseOrganizationsGetOrganizationFirmwareUpgradesByDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationFirmwareUpgradesByDevice")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationFirmwareUpgradesByDevice)
	return result, response, err

}

//GetOrganizationInventoryDevices Return the device inventory for an organization
/* Return the device inventory for an organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationInventoryDevicesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-inventory-devices
*/
func (s *OrganizationsService) GetOrganizationInventoryDevices(organizationID string, getOrganizationInventoryDevicesQueryParams *GetOrganizationInventoryDevicesQueryParams) (*ResponseOrganizationsGetOrganizationInventoryDevices, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/inventory/devices"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationInventoryDevicesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseOrganizationsGetOrganizationInventoryDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationInventoryDevices")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationInventoryDevices)
	return result, response, err

}

//GetOrganizationInventoryDevice Return a single device from the inventory of an organization
/* Return a single device from the inventory of an organization

@param organizationID organizationId path parameter. Organization ID
@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-inventory-device
*/
func (s *OrganizationsService) GetOrganizationInventoryDevice(organizationID string, serial string) (*ResponseOrganizationsGetOrganizationInventoryDevice, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/inventory/devices/{serial}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseOrganizationsGetOrganizationInventoryDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationInventoryDevice")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationInventoryDevice)
	return result, response, err

}

//GetOrganizationInventoryOnboardingCloudMonitoringImports Check the status of a committed Import operation
/* Check the status of a committed Import operation

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationInventoryOnboardingCloudMonitoringImportsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-inventory-onboarding-cloud-monitoring-imports
*/
func (s *OrganizationsService) GetOrganizationInventoryOnboardingCloudMonitoringImports(organizationID string, getOrganizationInventoryOnboardingCloudMonitoringImportsQueryParams *GetOrganizationInventoryOnboardingCloudMonitoringImportsQueryParams) (*ResponseOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringImports, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/inventory/onboarding/cloudMonitoring/imports"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationInventoryOnboardingCloudMonitoringImportsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringImports{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationInventoryOnboardingCloudMonitoringImports")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringImports)
	return result, response, err

}

//GetOrganizationInventoryOnboardingCloudMonitoringNetworks Returns list of networks eligible for adding cloud monitored device
/* Returns list of networks eligible for adding cloud monitored device

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationInventoryOnboardingCloudMonitoringNetworksQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-inventory-onboarding-cloud-monitoring-networks
*/
func (s *OrganizationsService) GetOrganizationInventoryOnboardingCloudMonitoringNetworks(organizationID string, getOrganizationInventoryOnboardingCloudMonitoringNetworksQueryParams *GetOrganizationInventoryOnboardingCloudMonitoringNetworksQueryParams) (*ResponseOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringNetworks, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/inventory/onboarding/cloudMonitoring/networks"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationInventoryOnboardingCloudMonitoringNetworksQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringNetworks{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationInventoryOnboardingCloudMonitoringNetworks")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringNetworks)
	return result, response, err

}

//GetOrganizationLicenses List the licenses for an organization
/* List the licenses for an organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationLicensesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-licenses
*/
func (s *OrganizationsService) GetOrganizationLicenses(organizationID string, getOrganizationLicensesQueryParams *GetOrganizationLicensesQueryParams) (*ResponseOrganizationsGetOrganizationLicenses, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/licenses"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationLicensesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseOrganizationsGetOrganizationLicenses{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationLicenses")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationLicenses)
	return result, response, err

}

//GetOrganizationLicensesOverview Return an overview of the license state for an organization
/* Return an overview of the license state for an organization

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-licenses-overview
*/
func (s *OrganizationsService) GetOrganizationLicensesOverview(organizationID string) (*ResponseOrganizationsGetOrganizationLicensesOverview, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/licenses/overview"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseOrganizationsGetOrganizationLicensesOverview{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationLicensesOverview")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationLicensesOverview)
	return result, response, err

}

//GetOrganizationLicense Display a license
/* Display a license

@param organizationID organizationId path parameter. Organization ID
@param licenseID licenseId path parameter. License ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-license
*/
func (s *OrganizationsService) GetOrganizationLicense(organizationID string, licenseID string) (*ResponseOrganizationsGetOrganizationLicense, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/licenses/{licenseId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{licenseId}", fmt.Sprintf("%v", licenseID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseOrganizationsGetOrganizationLicense{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationLicense")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationLicense)
	return result, response, err

}

//GetOrganizationLoginSecurity Returns the login security settings for an organization.
/* Returns the login security settings for an organization.

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-login-security
*/
func (s *OrganizationsService) GetOrganizationLoginSecurity(organizationID string) (*ResponseOrganizationsGetOrganizationLoginSecurity, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/loginSecurity"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseOrganizationsGetOrganizationLoginSecurity{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationLoginSecurity")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationLoginSecurity)
	return result, response, err

}

//GetOrganizationNetworks List the networks that the user has privileges on in an organization
/* List the networks that the user has privileges on in an organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationNetworksQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-networks
*/
func (s *OrganizationsService) GetOrganizationNetworks(organizationID string, getOrganizationNetworksQueryParams *GetOrganizationNetworksQueryParams) (*ResponseOrganizationsGetOrganizationNetworks, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/networks"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationNetworksQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseOrganizationsGetOrganizationNetworks{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationNetworks")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationNetworks)
	return result, response, err

}

//GetOrganizationOpenapiSpec Return the OpenAPI 2.0 Specification of the organization's API documentation in JSON
/* Return the OpenAPI 2.0 Specification of the organization's API documentation in JSON

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-openapi-spec
*/
func (s *OrganizationsService) GetOrganizationOpenapiSpec(organizationID string) (*ResponseOrganizationsGetOrganizationOpenapiSpec, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/openapiSpec"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseOrganizationsGetOrganizationOpenapiSpec{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationOpenapiSpec")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationOpenapiSpec)
	return result, response, err

}

//GetOrganizationPolicyObjects Lists Policy Objects belonging to the organization.
/* Lists Policy Objects belonging to the organization.

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationPolicyObjectsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-policy-objects
*/
func (s *OrganizationsService) GetOrganizationPolicyObjects(organizationID string, getOrganizationPolicyObjectsQueryParams *GetOrganizationPolicyObjectsQueryParams) (*ResponseOrganizationsGetOrganizationPolicyObjects, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/policyObjects"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationPolicyObjectsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseOrganizationsGetOrganizationPolicyObjects{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationPolicyObjects")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationPolicyObjects)
	return result, response, err

}

//GetOrganizationPolicyObjectsGroups Lists Policy Object Groups belonging to the organization.
/* Lists Policy Object Groups belonging to the organization.

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationPolicyObjectsGroupsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-policy-objects-groups
*/
func (s *OrganizationsService) GetOrganizationPolicyObjectsGroups(organizationID string, getOrganizationPolicyObjectsGroupsQueryParams *GetOrganizationPolicyObjectsGroupsQueryParams) (*ResponseOrganizationsGetOrganizationPolicyObjectsGroups, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/policyObjects/groups"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationPolicyObjectsGroupsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseOrganizationsGetOrganizationPolicyObjectsGroups{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationPolicyObjectsGroups")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationPolicyObjectsGroups)
	return result, response, err

}

//GetOrganizationPolicyObjectsGroup Shows details of a Policy Object Group.
/* Shows details of a Policy Object Group.

@param organizationID organizationId path parameter. Organization ID
@param policyObjectGroupID policyObjectGroupId path parameter. Policy object group ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-policy-objects-group
*/
func (s *OrganizationsService) GetOrganizationPolicyObjectsGroup(organizationID string, policyObjectGroupID string) (*ResponseOrganizationsGetOrganizationPolicyObjectsGroup, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/policyObjects/groups/{policyObjectGroupId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{policyObjectGroupId}", fmt.Sprintf("%v", policyObjectGroupID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseOrganizationsGetOrganizationPolicyObjectsGroup{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationPolicyObjectsGroup")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationPolicyObjectsGroup)
	return result, response, err

}

//GetOrganizationPolicyObject Shows details of a Policy Object.
/* Shows details of a Policy Object.

@param organizationID organizationId path parameter. Organization ID
@param policyObjectID policyObjectId path parameter. Policy object ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-policy-object
*/
func (s *OrganizationsService) GetOrganizationPolicyObject(organizationID string, policyObjectID string) (*ResponseOrganizationsGetOrganizationPolicyObject, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/policyObjects/{policyObjectId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{policyObjectId}", fmt.Sprintf("%v", policyObjectID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseOrganizationsGetOrganizationPolicyObject{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationPolicyObject")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationPolicyObject)
	return result, response, err

}

//GetOrganizationSaml Returns the SAML SSO enabled settings for an organization.
/* Returns the SAML SSO enabled settings for an organization.

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-saml
*/
func (s *OrganizationsService) GetOrganizationSaml(organizationID string) (*ResponseOrganizationsGetOrganizationSaml, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/saml"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseOrganizationsGetOrganizationSaml{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationSaml")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationSaml)
	return result, response, err

}

//GetOrganizationSamlIDps List the SAML IdPs in your organization.
/* List the SAML IdPs in your organization.

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-saml-idps
*/
func (s *OrganizationsService) GetOrganizationSamlIDps(organizationID string) (*ResponseOrganizationsGetOrganizationSamlIDps, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/saml/idps"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseOrganizationsGetOrganizationSamlIDps{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationSamlIdps")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationSamlIDps)
	return result, response, err

}

//GetOrganizationSamlIDp Get a SAML IdP from your organization.
/* Get a SAML IdP from your organization.

@param organizationID organizationId path parameter. Organization ID
@param idpID idpId path parameter. Idp ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-saml-idp
*/
func (s *OrganizationsService) GetOrganizationSamlIDp(organizationID string, idpID string) (*ResponseOrganizationsGetOrganizationSamlIDp, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/saml/idps/{idpId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{idpId}", fmt.Sprintf("%v", idpID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseOrganizationsGetOrganizationSamlIDp{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationSamlIdp")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationSamlIDp)
	return result, response, err

}

//GetOrganizationSamlRoles List the SAML roles for this organization
/* List the SAML roles for this organization

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-saml-roles
*/
func (s *OrganizationsService) GetOrganizationSamlRoles(organizationID string) (*ResponseOrganizationsGetOrganizationSamlRoles, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/samlRoles"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseOrganizationsGetOrganizationSamlRoles{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationSamlRoles")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationSamlRoles)
	return result, response, err

}

//GetOrganizationSamlRole Return a SAML role
/* Return a SAML role

@param organizationID organizationId path parameter. Organization ID
@param samlRoleID samlRoleId path parameter. Saml role ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-saml-role
*/
func (s *OrganizationsService) GetOrganizationSamlRole(organizationID string, samlRoleID string) (*ResponseOrganizationsGetOrganizationSamlRole, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/samlRoles/{samlRoleId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{samlRoleId}", fmt.Sprintf("%v", samlRoleID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseOrganizationsGetOrganizationSamlRole{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationSamlRole")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationSamlRole)
	return result, response, err

}

//GetOrganizationSNMP Return the SNMP settings for an organization
/* Return the SNMP settings for an organization

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-snmp
*/
func (s *OrganizationsService) GetOrganizationSNMP(organizationID string) (*ResponseOrganizationsGetOrganizationSNMP, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/snmp"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseOrganizationsGetOrganizationSNMP{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationSnmp")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationSNMP)
	return result, response, err

}

//GetOrganizationSummaryTopAppliancesByUtilization Return the top 10 appliances sorted by utilization over given time range.
/* Return the top 10 appliances sorted by utilization over given time range.

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationSummaryTopAppliancesByUtilizationQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-summary-top-appliances-by-utilization
*/
func (s *OrganizationsService) GetOrganizationSummaryTopAppliancesByUtilization(organizationID string, getOrganizationSummaryTopAppliancesByUtilizationQueryParams *GetOrganizationSummaryTopAppliancesByUtilizationQueryParams) (*ResponseOrganizationsGetOrganizationSummaryTopAppliancesByUtilization, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/summary/top/appliances/byUtilization"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationSummaryTopAppliancesByUtilizationQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseOrganizationsGetOrganizationSummaryTopAppliancesByUtilization{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationSummaryTopAppliancesByUtilization")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationSummaryTopAppliancesByUtilization)
	return result, response, err

}

//GetOrganizationSummaryTopClientsByUsage Return metrics for organization's top 10 clients by data usage (in mb) over given time range.
/* Return metrics for organization's top 10 clients by data usage (in mb) over given time range.

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationSummaryTopClientsByUsageQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-summary-top-clients-by-usage
*/
func (s *OrganizationsService) GetOrganizationSummaryTopClientsByUsage(organizationID string, getOrganizationSummaryTopClientsByUsageQueryParams *GetOrganizationSummaryTopClientsByUsageQueryParams) (*ResponseOrganizationsGetOrganizationSummaryTopClientsByUsage, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/summary/top/clients/byUsage"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationSummaryTopClientsByUsageQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseOrganizationsGetOrganizationSummaryTopClientsByUsage{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationSummaryTopClientsByUsage")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationSummaryTopClientsByUsage)
	return result, response, err

}

//GetOrganizationSummaryTopClientsManufacturersByUsage Return metrics for organization's top clients by data usage (in mb) over given time range, grouped by manufacturer.
/* Return metrics for organization's top clients by data usage (in mb) over given time range, grouped by manufacturer.

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationSummaryTopClientsManufacturersByUsageQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-summary-top-clients-manufacturers-by-usage
*/
func (s *OrganizationsService) GetOrganizationSummaryTopClientsManufacturersByUsage(organizationID string, getOrganizationSummaryTopClientsManufacturersByUsageQueryParams *GetOrganizationSummaryTopClientsManufacturersByUsageQueryParams) (*ResponseOrganizationsGetOrganizationSummaryTopClientsManufacturersByUsage, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/summary/top/clients/manufacturers/byUsage"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationSummaryTopClientsManufacturersByUsageQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseOrganizationsGetOrganizationSummaryTopClientsManufacturersByUsage{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationSummaryTopClientsManufacturersByUsage")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationSummaryTopClientsManufacturersByUsage)
	return result, response, err

}

//GetOrganizationSummaryTopDevicesByUsage Return metrics for organization's top 10 devices sorted by data usage over given time range
/* Return metrics for organization's top 10 devices sorted by data usage over given time range. Default unit is megabytes.

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationSummaryTopDevicesByUsageQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-summary-top-devices-by-usage
*/
func (s *OrganizationsService) GetOrganizationSummaryTopDevicesByUsage(organizationID string, getOrganizationSummaryTopDevicesByUsageQueryParams *GetOrganizationSummaryTopDevicesByUsageQueryParams) (*ResponseOrganizationsGetOrganizationSummaryTopDevicesByUsage, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/summary/top/devices/byUsage"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationSummaryTopDevicesByUsageQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseOrganizationsGetOrganizationSummaryTopDevicesByUsage{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationSummaryTopDevicesByUsage")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationSummaryTopDevicesByUsage)
	return result, response, err

}

//GetOrganizationSummaryTopDevicesModelsByUsage Return metrics for organization's top 10 device models sorted by data usage over given time range
/* Return metrics for organization's top 10 device models sorted by data usage over given time range. Default unit is megabytes.

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationSummaryTopDevicesModelsByUsageQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-summary-top-devices-models-by-usage
*/
func (s *OrganizationsService) GetOrganizationSummaryTopDevicesModelsByUsage(organizationID string, getOrganizationSummaryTopDevicesModelsByUsageQueryParams *GetOrganizationSummaryTopDevicesModelsByUsageQueryParams) (*ResponseOrganizationsGetOrganizationSummaryTopDevicesModelsByUsage, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/summary/top/devices/models/byUsage"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationSummaryTopDevicesModelsByUsageQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseOrganizationsGetOrganizationSummaryTopDevicesModelsByUsage{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationSummaryTopDevicesModelsByUsage")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationSummaryTopDevicesModelsByUsage)
	return result, response, err

}

//GetOrganizationSummaryTopSSIDsByUsage Return metrics for organization's top 10 ssids by data usage over given time range
/* Return metrics for organization's top 10 ssids by data usage over given time range. Default unit is megabytes.

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationSummaryTopSsidsByUsageQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-summary-top-ssids-by-usage
*/
func (s *OrganizationsService) GetOrganizationSummaryTopSSIDsByUsage(organizationID string, getOrganizationSummaryTopSsidsByUsageQueryParams *GetOrganizationSummaryTopSSIDsByUsageQueryParams) (*ResponseOrganizationsGetOrganizationSummaryTopSSIDsByUsage, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/summary/top/ssids/byUsage"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationSummaryTopSsidsByUsageQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseOrganizationsGetOrganizationSummaryTopSSIDsByUsage{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationSummaryTopSsidsByUsage")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationSummaryTopSSIDsByUsage)
	return result, response, err

}

//GetOrganizationSummaryTopSwitchesByEnergyUsage Return metrics for organization's top 10 switches by energy usage over given time range
/* Return metrics for organization's top 10 switches by energy usage over given time range. Default unit is joules.

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationSummaryTopSwitchesByEnergyUsageQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-summary-top-switches-by-energy-usage
*/
func (s *OrganizationsService) GetOrganizationSummaryTopSwitchesByEnergyUsage(organizationID string, getOrganizationSummaryTopSwitchesByEnergyUsageQueryParams *GetOrganizationSummaryTopSwitchesByEnergyUsageQueryParams) (*ResponseOrganizationsGetOrganizationSummaryTopSwitchesByEnergyUsage, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/summary/top/switches/byEnergyUsage"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationSummaryTopSwitchesByEnergyUsageQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseOrganizationsGetOrganizationSummaryTopSwitchesByEnergyUsage{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationSummaryTopSwitchesByEnergyUsage")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationSummaryTopSwitchesByEnergyUsage)
	return result, response, err

}

//GetOrganizationUplinksStatuses List the uplink status of every Meraki MX, MG and Z series devices in the organization
/* List the uplink status of every Meraki MX, MG and Z series devices in the organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationUplinksStatusesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-uplinks-statuses
*/
func (s *OrganizationsService) GetOrganizationUplinksStatuses(organizationID string, getOrganizationUplinksStatusesQueryParams *GetOrganizationUplinksStatusesQueryParams) (*ResponseOrganizationsGetOrganizationUplinksStatuses, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/uplinks/statuses"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationUplinksStatusesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseOrganizationsGetOrganizationUplinksStatuses{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationUplinksStatuses")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationUplinksStatuses)
	return result, response, err

}

//GetOrganizationWebhooksAlertTypes Return a list of alert types to be used with managing webhook alerts
/* Return a list of alert types to be used with managing webhook alerts

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWebhooksAlertTypesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-webhooks-alert-types
*/
func (s *OrganizationsService) GetOrganizationWebhooksAlertTypes(organizationID string, getOrganizationWebhooksAlertTypesQueryParams *GetOrganizationWebhooksAlertTypesQueryParams) (*ResponseOrganizationsGetOrganizationWebhooksAlertTypes, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/webhooks/alertTypes"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWebhooksAlertTypesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseOrganizationsGetOrganizationWebhooksAlertTypes{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWebhooksAlertTypes")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationWebhooksAlertTypes)
	return result, response, err

}

//GetOrganizationWebhooksLogs Return the log of webhook POSTs sent
/* Return the log of webhook POSTs sent

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWebhooksLogsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-webhooks-logs
*/
func (s *OrganizationsService) GetOrganizationWebhooksLogs(organizationID string, getOrganizationWebhooksLogsQueryParams *GetOrganizationWebhooksLogsQueryParams) (*ResponseOrganizationsGetOrganizationWebhooksLogs, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/webhooks/logs"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWebhooksLogsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseOrganizationsGetOrganizationWebhooksLogs{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWebhooksLogs")
	}

	result := response.Result().(*ResponseOrganizationsGetOrganizationWebhooksLogs)
	return result, response, err

}

//CreateOrganization Create a new organization
/* Create a new organization


Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-organization
*/

func (s *OrganizationsService) CreateOrganization(requestOrganizationsCreateOrganization *RequestOrganizationsCreateOrganization) (*ResponseOrganizationsCreateOrganization, *resty.Response, error) {
	path := "/api/v1/organizations"
	s.rateLimiterBucket.Wait(1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsCreateOrganization).
		SetResult(&ResponseOrganizationsCreateOrganization{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateOrganization")
	}

	result := response.Result().(*ResponseOrganizationsCreateOrganization)
	return result, response, err

}

//CreateOrganizationActionBatch Create an action batch
/* Create an action batch

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-organization-action-batch
*/

func (s *OrganizationsService) CreateOrganizationActionBatch(organizationID string, requestOrganizationsCreateOrganizationActionBatch *RequestOrganizationsCreateOrganizationActionBatch) (*ResponseOrganizationsCreateOrganizationActionBatch, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/actionBatches"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsCreateOrganizationActionBatch).
		SetResult(&ResponseOrganizationsCreateOrganizationActionBatch{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateOrganizationActionBatch")
	}

	result := response.Result().(*ResponseOrganizationsCreateOrganizationActionBatch)
	return result, response, err

}

//CreateOrganizationAdaptivePolicyACL Creates new adaptive policy ACL
/* Creates new adaptive policy ACL

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-organization-adaptive-policy-acl
*/

func (s *OrganizationsService) CreateOrganizationAdaptivePolicyACL(organizationID string, requestOrganizationsCreateOrganizationAdaptivePolicyAcl *RequestOrganizationsCreateOrganizationAdaptivePolicyACL) (*ResponseOrganizationsCreateOrganizationAdaptivePolicyACL, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/adaptivePolicy/acls"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsCreateOrganizationAdaptivePolicyAcl).
		SetResult(&ResponseOrganizationsCreateOrganizationAdaptivePolicyACL{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateOrganizationAdaptivePolicyAcl")
	}

	result := response.Result().(*ResponseOrganizationsCreateOrganizationAdaptivePolicyACL)
	return result, response, err

}

//CreateOrganizationAdaptivePolicyGroup Creates a new adaptive policy group
/* Creates a new adaptive policy group

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-organization-adaptive-policy-group
*/

func (s *OrganizationsService) CreateOrganizationAdaptivePolicyGroup(organizationID string, requestOrganizationsCreateOrganizationAdaptivePolicyGroup *RequestOrganizationsCreateOrganizationAdaptivePolicyGroup) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/adaptivePolicy/groups"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsCreateOrganizationAdaptivePolicyGroup).
		// SetResult(&ResponseOrganizationsCreateOrganizationAdaptivePolicyGroup{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateOrganizationAdaptivePolicyGroup")
	}

	return response, err

}

//CreateOrganizationAdaptivePolicyPolicy Add an Adaptive Policy
/* Add an Adaptive Policy

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-organization-adaptive-policy-policy
*/

func (s *OrganizationsService) CreateOrganizationAdaptivePolicyPolicy(organizationID string, requestOrganizationsCreateOrganizationAdaptivePolicyPolicy *RequestOrganizationsCreateOrganizationAdaptivePolicyPolicy) (*ResponseOrganizationsCreateOrganizationAdaptivePolicyPolicy, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/adaptivePolicy/policies"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsCreateOrganizationAdaptivePolicyPolicy).
		SetResult(&ResponseOrganizationsCreateOrganizationAdaptivePolicyPolicy{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateOrganizationAdaptivePolicyPolicy")
	}

	result := response.Result().(*ResponseOrganizationsCreateOrganizationAdaptivePolicyPolicy)
	return result, response, err

}

//CreateOrganizationAdmin Create a new dashboard administrator
/* Create a new dashboard administrator

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-organization-admin
*/

func (s *OrganizationsService) CreateOrganizationAdmin(organizationID string, requestOrganizationsCreateOrganizationAdmin *RequestOrganizationsCreateOrganizationAdmin) (*ResponseOrganizationsCreateOrganizationAdmin, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/admins"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsCreateOrganizationAdmin).
		SetResult(&ResponseOrganizationsCreateOrganizationAdmin{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateOrganizationAdmin")
	}

	result := response.Result().(*ResponseOrganizationsCreateOrganizationAdmin)
	return result, response, err

}

//CreateOrganizationAlertsProfile Create an organization-wide alert configuration
/* Create an organization-wide alert configuration

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-organization-alerts-profile
*/

func (s *OrganizationsService) CreateOrganizationAlertsProfile(organizationID string, requestOrganizationsCreateOrganizationAlertsProfile *RequestOrganizationsCreateOrganizationAlertsProfile) (*ResponseOrganizationsCreateOrganizationAlertsProfile, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/alerts/profiles"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsCreateOrganizationAlertsProfile).
		SetResult(&ResponseOrganizationsCreateOrganizationAlertsProfile{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateOrganizationAlertsProfile")
	}

	result := response.Result().(*ResponseOrganizationsCreateOrganizationAlertsProfile)
	return result, response, err

}

//CreateOrganizationBrandingPolicy Add a new branding policy to an organization
/* Add a new branding policy to an organization

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-organization-branding-policy
*/

func (s *OrganizationsService) CreateOrganizationBrandingPolicy(organizationID string, requestOrganizationsCreateOrganizationBrandingPolicy *RequestOrganizationsCreateOrganizationBrandingPolicy) (*ResponseOrganizationsCreateOrganizationBrandingPolicy, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/brandingPolicies"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsCreateOrganizationBrandingPolicy).
		SetResult(&ResponseOrganizationsCreateOrganizationBrandingPolicy{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateOrganizationBrandingPolicy")
	}

	result := response.Result().(*ResponseOrganizationsCreateOrganizationBrandingPolicy)
	return result, response, err

}

//ClaimIntoOrganization Claim a list of devices, licenses, and/or orders into an organization
/* Claim a list of devices, licenses, and/or orders into an organization. When claiming by order, all devices and licenses in the order will be claimed; licenses will be added to the organization and devices will be placed in the organization's inventory.

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!claim-into-organization
*/

func (s *OrganizationsService) ClaimIntoOrganization(organizationID string, requestOrganizationsClaimIntoOrganization *RequestOrganizationsClaimIntoOrganization) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/claim"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsClaimIntoOrganization).
		// SetResult(&ResponseOrganizationsClaimIntoOrganization{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation ClaimIntoOrganization")
	}

	return response, err

}

//CloneOrganization Create a new organization by cloning the addressed organization
/* Create a new organization by cloning the addressed organization

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!clone-organization
*/

func (s *OrganizationsService) CloneOrganization(organizationID string, requestOrganizationsCloneOrganization *RequestOrganizationsCloneOrganization) (*ResponseOrganizationsCloneOrganization, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/clone"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsCloneOrganization).
		SetResult(&ResponseOrganizationsCloneOrganization{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CloneOrganization")
	}

	result := response.Result().(*ResponseOrganizationsCloneOrganization)
	return result, response, err

}

//CreateOrganizationConfigTemplate Create a new configuration template
/* Create a new configuration template

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-organization-config-template
*/

func (s *OrganizationsService) CreateOrganizationConfigTemplate(organizationID string, requestOrganizationsCreateOrganizationConfigTemplate *RequestOrganizationsCreateOrganizationConfigTemplate) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/configTemplates"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsCreateOrganizationConfigTemplate).
		// SetResult(&ResponseOrganizationsCreateOrganizationConfigTemplate{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateOrganizationConfigTemplate")
	}

	return response, err

}

//CreateOrganizationEarlyAccessFeaturesOptIn Create a new early access feature opt-in for an organization
/* Create a new early access feature opt-in for an organization

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-organization-early-access-features-opt-in
*/

func (s *OrganizationsService) CreateOrganizationEarlyAccessFeaturesOptIn(organizationID string, requestOrganizationsCreateOrganizationEarlyAccessFeaturesOptIn *RequestOrganizationsCreateOrganizationEarlyAccessFeaturesOptIn) (*ResponseOrganizationsCreateOrganizationEarlyAccessFeaturesOptIn, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/earlyAccess/features/optIns"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsCreateOrganizationEarlyAccessFeaturesOptIn).
		SetResult(&ResponseOrganizationsCreateOrganizationEarlyAccessFeaturesOptIn{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateOrganizationEarlyAccessFeaturesOptIn")
	}

	result := response.Result().(*ResponseOrganizationsCreateOrganizationEarlyAccessFeaturesOptIn)
	return result, response, err

}

//ClaimIntoOrganizationInventory Claim a list of devices, licenses, and/or orders into an organization inventory
/* Claim a list of devices, licenses, and/or orders into an organization inventory. When claiming by order, all devices and licenses in the order will be claimed; licenses will be added to the organization and devices will be placed in the organization's inventory. Use /organizations/{organizationId}/inventory/release to release devices from an organization.

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!claim-into-organization-inventory
*/

func (s *OrganizationsService) ClaimIntoOrganizationInventory(organizationID string, requestOrganizationsClaimIntoOrganizationInventory *RequestOrganizationsClaimIntoOrganizationInventory) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/inventory/claim"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsClaimIntoOrganizationInventory).
		// SetResult(&ResponseOrganizationsClaimIntoOrganizationInventory{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation ClaimIntoOrganizationInventory")
	}

	return response, err

}

//CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent Imports event logs related to the onboarding app into elastisearch
/* Imports event logs related to the onboarding app into elastisearch

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-organization-inventory-onboarding-cloud-monitoring-export-event
*/

func (s *OrganizationsService) CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent(organizationID string, requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringExportEvent *RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringExportEvent) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/inventory/onboarding/cloudMonitoring/exportEvents"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringExportEvent).
		// SetResult(&ResponseOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringExportEvent{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent")
	}

	return response, err

}

//CreateOrganizationInventoryOnboardingCloudMonitoringImport Commits the import operation to complete the onboarding of a device into Dashboard for monitoring.
/* Commits the import operation to complete the onboarding of a device into Dashboard for monitoring.

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-organization-inventory-onboarding-cloud-monitoring-import
*/

func (s *OrganizationsService) CreateOrganizationInventoryOnboardingCloudMonitoringImport(organizationID string, requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringImport *RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringImport) (*ResponseOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringImport, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/inventory/onboarding/cloudMonitoring/imports"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringImport).
		SetResult(&ResponseOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringImport{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateOrganizationInventoryOnboardingCloudMonitoringImport")
	}

	result := response.Result().(*ResponseOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringImport)
	return result, response, err

}

//CreateOrganizationInventoryOnboardingCloudMonitoringPrepare Initiates or updates an import session
/* Initiates or updates an import session. An import ID will be generated and used when you are ready to commit the import.

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-organization-inventory-onboarding-cloud-monitoring-prepare
*/

func (s *OrganizationsService) CreateOrganizationInventoryOnboardingCloudMonitoringPrepare(organizationID string, requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepare *RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepare) (*ResponseOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepare, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/inventory/onboarding/cloudMonitoring/prepare"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepare).
		SetResult(&ResponseOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepare{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateOrganizationInventoryOnboardingCloudMonitoringPrepare")
	}

	result := response.Result().(*ResponseOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepare)
	return result, response, err

}

//ReleaseFromOrganizationInventory Release a list of claimed devices from an organization.
/* Release a list of claimed devices from an organization.

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!release-from-organization-inventory
*/

func (s *OrganizationsService) ReleaseFromOrganizationInventory(organizationID string, requestOrganizationsReleaseFromOrganizationInventory *RequestOrganizationsReleaseFromOrganizationInventory) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/inventory/release"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsReleaseFromOrganizationInventory).
		// SetResult(&ResponseOrganizationsReleaseFromOrganizationInventory{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation ReleaseFromOrganizationInventory")
	}

	return response, err

}

//AssignOrganizationLicensesSeats Assign SM seats to a network
/* Assign SM seats to a network. This will increase the managed SM device limit of the network

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!assign-organization-licenses-seats
*/

func (s *OrganizationsService) AssignOrganizationLicensesSeats(organizationID string, requestOrganizationsAssignOrganizationLicensesSeats *RequestOrganizationsAssignOrganizationLicensesSeats) (*ResponseOrganizationsAssignOrganizationLicensesSeats, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/licenses/assignSeats"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsAssignOrganizationLicensesSeats).
		SetResult(&ResponseOrganizationsAssignOrganizationLicensesSeats{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation AssignOrganizationLicensesSeats")
	}

	result := response.Result().(*ResponseOrganizationsAssignOrganizationLicensesSeats)
	return result, response, err

}

//MoveOrganizationLicenses Move licenses to another organization
/* Move licenses to another organization. This will also move any devices that the licenses are assigned to

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!move-organization-licenses
*/

func (s *OrganizationsService) MoveOrganizationLicenses(organizationID string, requestOrganizationsMoveOrganizationLicenses *RequestOrganizationsMoveOrganizationLicenses) (*ResponseOrganizationsMoveOrganizationLicenses, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/licenses/move"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsMoveOrganizationLicenses).
		SetResult(&ResponseOrganizationsMoveOrganizationLicenses{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation MoveOrganizationLicenses")
	}

	result := response.Result().(*ResponseOrganizationsMoveOrganizationLicenses)
	return result, response, err

}

//MoveOrganizationLicensesSeats Move SM seats to another organization
/* Move SM seats to another organization

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!move-organization-licenses-seats
*/

func (s *OrganizationsService) MoveOrganizationLicensesSeats(organizationID string, requestOrganizationsMoveOrganizationLicensesSeats *RequestOrganizationsMoveOrganizationLicensesSeats) (*ResponseOrganizationsMoveOrganizationLicensesSeats, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/licenses/moveSeats"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsMoveOrganizationLicensesSeats).
		SetResult(&ResponseOrganizationsMoveOrganizationLicensesSeats{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation MoveOrganizationLicensesSeats")
	}

	result := response.Result().(*ResponseOrganizationsMoveOrganizationLicensesSeats)
	return result, response, err

}

//RenewOrganizationLicensesSeats Renew SM seats of a license
/* Renew SM seats of a license. This will extend the license expiration date of managed SM devices covered by this license

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!renew-organization-licenses-seats
*/

func (s *OrganizationsService) RenewOrganizationLicensesSeats(organizationID string, requestOrganizationsRenewOrganizationLicensesSeats *RequestOrganizationsRenewOrganizationLicensesSeats) (*ResponseOrganizationsRenewOrganizationLicensesSeats, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/licenses/renewSeats"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsRenewOrganizationLicensesSeats).
		SetResult(&ResponseOrganizationsRenewOrganizationLicensesSeats{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation RenewOrganizationLicensesSeats")
	}

	result := response.Result().(*ResponseOrganizationsRenewOrganizationLicensesSeats)
	return result, response, err

}

//CreateOrganizationNetwork Create a network
/* Create a network

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-organization-network
*/

func (s *OrganizationsService) CreateOrganizationNetwork(organizationID string, requestOrganizationsCreateOrganizationNetwork *RequestOrganizationsCreateOrganizationNetwork) (*ResponseOrganizationsCreateOrganizationNetwork, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/networks"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsCreateOrganizationNetwork).
		SetResult(&ResponseOrganizationsCreateOrganizationNetwork{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateOrganizationNetwork")
	}

	result := response.Result().(*ResponseOrganizationsCreateOrganizationNetwork)
	return result, response, err

}

//CombineOrganizationNetworks Combine multiple networks into a single network
/* Combine multiple networks into a single network

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!combine-organization-networks
*/

func (s *OrganizationsService) CombineOrganizationNetworks(organizationID string, requestOrganizationsCombineOrganizationNetworks *RequestOrganizationsCombineOrganizationNetworks) (*ResponseOrganizationsCombineOrganizationNetworks, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/networks/combine"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsCombineOrganizationNetworks).
		SetResult(&ResponseOrganizationsCombineOrganizationNetworks{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CombineOrganizationNetworks")
	}

	result := response.Result().(*ResponseOrganizationsCombineOrganizationNetworks)
	return result, response, err

}

//CreateOrganizationPolicyObject Creates a new Policy Object.
/* Creates a new Policy Object.

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-organization-policy-object
*/

func (s *OrganizationsService) CreateOrganizationPolicyObject(organizationID string, requestOrganizationsCreateOrganizationPolicyObject *RequestOrganizationsCreateOrganizationPolicyObject) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/policyObjects"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsCreateOrganizationPolicyObject).
		// SetResult(&ResponseOrganizationsCreateOrganizationPolicyObject{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateOrganizationPolicyObject")
	}

	return response, err

}

//CreateOrganizationPolicyObjectsGroup Creates a new Policy Object Group.
/* Creates a new Policy Object Group.

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-organization-policy-objects-group
*/

func (s *OrganizationsService) CreateOrganizationPolicyObjectsGroup(organizationID string, requestOrganizationsCreateOrganizationPolicyObjectsGroup *RequestOrganizationsCreateOrganizationPolicyObjectsGroup) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/policyObjects/groups"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsCreateOrganizationPolicyObjectsGroup).
		// SetResult(&ResponseOrganizationsCreateOrganizationPolicyObjectsGroup{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateOrganizationPolicyObjectsGroup")
	}

	return response, err

}

//CreateOrganizationSamlIDp Create a SAML IdP for your organization.
/* Create a SAML IdP for your organization.

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-organization-saml-idp
*/

func (s *OrganizationsService) CreateOrganizationSamlIDp(organizationID string, requestOrganizationsCreateOrganizationSamlIdp *RequestOrganizationsCreateOrganizationSamlIDp) (*ResponseOrganizationsCreateOrganizationSamlIDp, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/saml/idps"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsCreateOrganizationSamlIdp).
		SetResult(&ResponseOrganizationsCreateOrganizationSamlIDp{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateOrganizationSamlIdp")
	}

	result := response.Result().(*ResponseOrganizationsCreateOrganizationSamlIDp)
	return result, response, err

}

//CreateOrganizationSamlRole Create a SAML role
/* Create a SAML role

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-organization-saml-role
*/

func (s *OrganizationsService) CreateOrganizationSamlRole(organizationID string, requestOrganizationsCreateOrganizationSamlRole *RequestOrganizationsCreateOrganizationSamlRole) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/samlRoles"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsCreateOrganizationSamlRole).
		// SetResult(&ResponseOrganizationsCreateOrganizationSamlRole{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateOrganizationSamlRole")
	}

	return response, err

}

//UpdateOrganization Update an organization
/* Update an organization

@param organizationID organizationId path parameter. Organization ID
*/
func (s *OrganizationsService) UpdateOrganization(organizationID string, requestOrganizationsUpdateOrganization *RequestOrganizationsUpdateOrganization) (*ResponseOrganizationsUpdateOrganization, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsUpdateOrganization).
		SetResult(&ResponseOrganizationsUpdateOrganization{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateOrganization")
	}

	result := response.Result().(*ResponseOrganizationsUpdateOrganization)
	return result, response, err

}

//UpdateOrganizationActionBatch Update an action batch
/* Update an action batch

@param organizationID organizationId path parameter. Organization ID
@param actionBatchID actionBatchId path parameter. Action batch ID
*/
func (s *OrganizationsService) UpdateOrganizationActionBatch(organizationID string, actionBatchID string, requestOrganizationsUpdateOrganizationActionBatch *RequestOrganizationsUpdateOrganizationActionBatch) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/actionBatches/{actionBatchId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{actionBatchId}", fmt.Sprintf("%v", actionBatchID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsUpdateOrganizationActionBatch).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateOrganizationActionBatch")
	}

	return response, err

}

//UpdateOrganizationAdaptivePolicyACL Updates an adaptive policy ACL
/* Updates an adaptive policy ACL

@param organizationID organizationId path parameter. Organization ID
@param aclID aclId path parameter. Acl ID
*/
func (s *OrganizationsService) UpdateOrganizationAdaptivePolicyACL(organizationID string, aclID string, requestOrganizationsUpdateOrganizationAdaptivePolicyAcl *RequestOrganizationsUpdateOrganizationAdaptivePolicyACL) (*ResponseOrganizationsUpdateOrganizationAdaptivePolicyACL, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/adaptivePolicy/acls/{aclId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{aclId}", fmt.Sprintf("%v", aclID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsUpdateOrganizationAdaptivePolicyAcl).
		SetResult(&ResponseOrganizationsUpdateOrganizationAdaptivePolicyACL{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateOrganizationAdaptivePolicyAcl")
	}

	result := response.Result().(*ResponseOrganizationsUpdateOrganizationAdaptivePolicyACL)
	return result, response, err

}

//UpdateOrganizationAdaptivePolicyGroup Updates an adaptive policy group
/* Updates an adaptive policy group. If updating "Infrastructure", only the SGT is allowed. Cannot update "Unknown".

@param organizationID organizationId path parameter. Organization ID
@param id id path parameter.
*/
func (s *OrganizationsService) UpdateOrganizationAdaptivePolicyGroup(organizationID string, id string, requestOrganizationsUpdateOrganizationAdaptivePolicyGroup *RequestOrganizationsUpdateOrganizationAdaptivePolicyGroup) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/adaptivePolicy/groups/{id}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsUpdateOrganizationAdaptivePolicyGroup).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateOrganizationAdaptivePolicyGroup")
	}

	return response, err

}

//UpdateOrganizationAdaptivePolicyPolicy Update an Adaptive Policy
/* Update an Adaptive Policy

@param organizationID organizationId path parameter. Organization ID
@param id id path parameter.
*/
func (s *OrganizationsService) UpdateOrganizationAdaptivePolicyPolicy(organizationID string, id string, requestOrganizationsUpdateOrganizationAdaptivePolicyPolicy *RequestOrganizationsUpdateOrganizationAdaptivePolicyPolicy) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/adaptivePolicy/policies/{id}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsUpdateOrganizationAdaptivePolicyPolicy).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateOrganizationAdaptivePolicyPolicy")
	}

	return response, err

}

//UpdateOrganizationAdaptivePolicySettings Update global adaptive policy settings
/* Update global adaptive policy settings

@param organizationID organizationId path parameter. Organization ID
*/
func (s *OrganizationsService) UpdateOrganizationAdaptivePolicySettings(organizationID string, requestOrganizationsUpdateOrganizationAdaptivePolicySettings *RequestOrganizationsUpdateOrganizationAdaptivePolicySettings) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/adaptivePolicy/settings"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsUpdateOrganizationAdaptivePolicySettings).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateOrganizationAdaptivePolicySettings")
	}

	return response, err

}

//UpdateOrganizationAdmin Update an administrator
/* Update an administrator

@param organizationID organizationId path parameter. Organization ID
@param adminID adminId path parameter. Admin ID
*/
func (s *OrganizationsService) UpdateOrganizationAdmin(organizationID string, adminID string, requestOrganizationsUpdateOrganizationAdmin *RequestOrganizationsUpdateOrganizationAdmin) (*ResponseOrganizationsUpdateOrganizationAdmin, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/admins/{adminId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{adminId}", fmt.Sprintf("%v", adminID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsUpdateOrganizationAdmin).
		SetResult(&ResponseOrganizationsUpdateOrganizationAdmin{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateOrganizationAdmin")
	}

	result := response.Result().(*ResponseOrganizationsUpdateOrganizationAdmin)
	return result, response, err

}

//UpdateOrganizationAlertsProfile Update an organization-wide alert config
/* Update an organization-wide alert config

@param organizationID organizationId path parameter. Organization ID
@param alertConfigID alertConfigId path parameter. Alert config ID
*/
func (s *OrganizationsService) UpdateOrganizationAlertsProfile(organizationID string, alertConfigID string, requestOrganizationsUpdateOrganizationAlertsProfile *RequestOrganizationsUpdateOrganizationAlertsProfile) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/alerts/profiles/{alertConfigId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{alertConfigId}", fmt.Sprintf("%v", alertConfigID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsUpdateOrganizationAlertsProfile).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateOrganizationAlertsProfile")
	}

	return response, err

}

//UpdateOrganizationBrandingPoliciesPriorities Update the priority ordering of an organization's branding policies.
/* Update the priority ordering of an organization's branding policies.

@param organizationID organizationId path parameter. Organization ID
*/
func (s *OrganizationsService) UpdateOrganizationBrandingPoliciesPriorities(organizationID string, requestOrganizationsUpdateOrganizationBrandingPoliciesPriorities *RequestOrganizationsUpdateOrganizationBrandingPoliciesPriorities) (*ResponseOrganizationsUpdateOrganizationBrandingPoliciesPriorities, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/brandingPolicies/priorities"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsUpdateOrganizationBrandingPoliciesPriorities).
		SetResult(&ResponseOrganizationsUpdateOrganizationBrandingPoliciesPriorities{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateOrganizationBrandingPoliciesPriorities")
	}

	result := response.Result().(*ResponseOrganizationsUpdateOrganizationBrandingPoliciesPriorities)
	return result, response, err

}

//UpdateOrganizationBrandingPolicy Update a branding policy
/* Update a branding policy

@param organizationID organizationId path parameter. Organization ID
@param brandingPolicyID brandingPolicyId path parameter. Branding policy ID
*/
func (s *OrganizationsService) UpdateOrganizationBrandingPolicy(organizationID string, brandingPolicyID string, requestOrganizationsUpdateOrganizationBrandingPolicy *RequestOrganizationsUpdateOrganizationBrandingPolicy) (*ResponseOrganizationsUpdateOrganizationBrandingPolicy, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/brandingPolicies/{brandingPolicyId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{brandingPolicyId}", fmt.Sprintf("%v", brandingPolicyID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsUpdateOrganizationBrandingPolicy).
		SetResult(&ResponseOrganizationsUpdateOrganizationBrandingPolicy{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateOrganizationBrandingPolicy")
	}

	result := response.Result().(*ResponseOrganizationsUpdateOrganizationBrandingPolicy)
	return result, response, err

}

//UpdateOrganizationConfigTemplate Update a configuration template
/* Update a configuration template

@param organizationID organizationId path parameter. Organization ID
@param configTemplateID configTemplateId path parameter. Config template ID
*/
func (s *OrganizationsService) UpdateOrganizationConfigTemplate(organizationID string, configTemplateID string, requestOrganizationsUpdateOrganizationConfigTemplate *RequestOrganizationsUpdateOrganizationConfigTemplate) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/configTemplates/{configTemplateId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{configTemplateId}", fmt.Sprintf("%v", configTemplateID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsUpdateOrganizationConfigTemplate).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateOrganizationConfigTemplate")
	}

	return response, err

}

//UpdateOrganizationEarlyAccessFeaturesOptIn Update an early access feature opt-in for an organization
/* Update an early access feature opt-in for an organization

@param organizationID organizationId path parameter. Organization ID
@param optInID optInId path parameter. Opt in ID
*/
func (s *OrganizationsService) UpdateOrganizationEarlyAccessFeaturesOptIn(organizationID string, optInID string, requestOrganizationsUpdateOrganizationEarlyAccessFeaturesOptIn *RequestOrganizationsUpdateOrganizationEarlyAccessFeaturesOptIn) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/earlyAccess/features/optIns/{optInId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{optInId}", fmt.Sprintf("%v", optInID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsUpdateOrganizationEarlyAccessFeaturesOptIn).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateOrganizationEarlyAccessFeaturesOptIn")
	}

	return response, err

}

//UpdateOrganizationLicense Update a license
/* Update a license

@param organizationID organizationId path parameter. Organization ID
@param licenseID licenseId path parameter. License ID
*/
func (s *OrganizationsService) UpdateOrganizationLicense(organizationID string, licenseID string, requestOrganizationsUpdateOrganizationLicense *RequestOrganizationsUpdateOrganizationLicense) (*ResponseOrganizationsUpdateOrganizationLicense, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/licenses/{licenseId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{licenseId}", fmt.Sprintf("%v", licenseID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsUpdateOrganizationLicense).
		SetResult(&ResponseOrganizationsUpdateOrganizationLicense{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateOrganizationLicense")
	}

	result := response.Result().(*ResponseOrganizationsUpdateOrganizationLicense)
	return result, response, err

}

//UpdateOrganizationLoginSecurity Update the login security settings for an organization
/* Update the login security settings for an organization

@param organizationID organizationId path parameter. Organization ID
*/
func (s *OrganizationsService) UpdateOrganizationLoginSecurity(organizationID string, requestOrganizationsUpdateOrganizationLoginSecurity *RequestOrganizationsUpdateOrganizationLoginSecurity) (*ResponseOrganizationsUpdateOrganizationLoginSecurity, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/loginSecurity"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsUpdateOrganizationLoginSecurity).
		SetResult(&ResponseOrganizationsUpdateOrganizationLoginSecurity{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateOrganizationLoginSecurity")
	}

	result := response.Result().(*ResponseOrganizationsUpdateOrganizationLoginSecurity)
	return result, response, err

}

//UpdateOrganizationPolicyObjectsGroup Updates a Policy Object Group.
/* Updates a Policy Object Group.

@param organizationID organizationId path parameter. Organization ID
@param policyObjectGroupID policyObjectGroupId path parameter. Policy object group ID
*/
func (s *OrganizationsService) UpdateOrganizationPolicyObjectsGroup(organizationID string, policyObjectGroupID string, requestOrganizationsUpdateOrganizationPolicyObjectsGroup *RequestOrganizationsUpdateOrganizationPolicyObjectsGroup) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/policyObjects/groups/{policyObjectGroupId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{policyObjectGroupId}", fmt.Sprintf("%v", policyObjectGroupID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsUpdateOrganizationPolicyObjectsGroup).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateOrganizationPolicyObjectsGroup")
	}

	return response, err

}

//UpdateOrganizationPolicyObject Updates a Policy Object.
/* Updates a Policy Object.

@param organizationID organizationId path parameter. Organization ID
@param policyObjectID policyObjectId path parameter. Policy object ID
*/
func (s *OrganizationsService) UpdateOrganizationPolicyObject(organizationID string, policyObjectID string, requestOrganizationsUpdateOrganizationPolicyObject *RequestOrganizationsUpdateOrganizationPolicyObject) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/policyObjects/{policyObjectId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{policyObjectId}", fmt.Sprintf("%v", policyObjectID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsUpdateOrganizationPolicyObject).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateOrganizationPolicyObject")
	}

	return response, err

}

//UpdateOrganizationSaml Updates the SAML SSO enabled settings for an organization.
/* Updates the SAML SSO enabled settings for an organization.

@param organizationID organizationId path parameter. Organization ID
*/
func (s *OrganizationsService) UpdateOrganizationSaml(organizationID string, requestOrganizationsUpdateOrganizationSaml *RequestOrganizationsUpdateOrganizationSaml) (*ResponseOrganizationsUpdateOrganizationSaml, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/saml"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsUpdateOrganizationSaml).
		SetResult(&ResponseOrganizationsUpdateOrganizationSaml{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateOrganizationSaml")
	}

	result := response.Result().(*ResponseOrganizationsUpdateOrganizationSaml)
	return result, response, err

}

//UpdateOrganizationSamlIDp Update a SAML IdP in your organization
/* Update a SAML IdP in your organization

@param organizationID organizationId path parameter. Organization ID
@param idpID idpId path parameter. Idp ID
*/
func (s *OrganizationsService) UpdateOrganizationSamlIDp(organizationID string, idpID string, requestOrganizationsUpdateOrganizationSamlIdp *RequestOrganizationsUpdateOrganizationSamlIDp) (*ResponseOrganizationsUpdateOrganizationSamlIDp, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/saml/idps/{idpId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{idpId}", fmt.Sprintf("%v", idpID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsUpdateOrganizationSamlIdp).
		SetResult(&ResponseOrganizationsUpdateOrganizationSamlIDp{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateOrganizationSamlIdp")
	}

	result := response.Result().(*ResponseOrganizationsUpdateOrganizationSamlIDp)
	return result, response, err

}

//UpdateOrganizationSamlRole Update a SAML role
/* Update a SAML role

@param organizationID organizationId path parameter. Organization ID
@param samlRoleID samlRoleId path parameter. Saml role ID
*/
func (s *OrganizationsService) UpdateOrganizationSamlRole(organizationID string, samlRoleID string, requestOrganizationsUpdateOrganizationSamlRole *RequestOrganizationsUpdateOrganizationSamlRole) (*ResponseOrganizationsUpdateOrganizationSamlRole, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/samlRoles/{samlRoleId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{samlRoleId}", fmt.Sprintf("%v", samlRoleID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsUpdateOrganizationSamlRole).
		SetResult(&ResponseOrganizationsUpdateOrganizationSamlRole{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateOrganizationSamlRole")
	}

	result := response.Result().(*ResponseOrganizationsUpdateOrganizationSamlRole)
	return result, response, err

}

//UpdateOrganizationSNMP Update the SNMP settings for an organization
/* Update the SNMP settings for an organization

@param organizationID organizationId path parameter. Organization ID
*/
func (s *OrganizationsService) UpdateOrganizationSNMP(organizationID string, requestOrganizationsUpdateOrganizationSnmp *RequestOrganizationsUpdateOrganizationSNMP) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/snmp"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestOrganizationsUpdateOrganizationSnmp).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateOrganizationSnmp")
	}

	return response, err

}

//DeleteOrganization Delete an organization
/* Delete an organization

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-organization
*/
func (s *OrganizationsService) DeleteOrganization(organizationID string) (*resty.Response, error) {
	//organizationID string
	path := "/api/v1/organizations/{organizationId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteOrganization")
	}

	return response, err

}

//DeleteOrganizationActionBatch Delete an action batch
/* Delete an action batch

@param organizationID organizationId path parameter. Organization ID
@param actionBatchID actionBatchId path parameter. Action batch ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-organization-action-batch
*/
func (s *OrganizationsService) DeleteOrganizationActionBatch(organizationID string, actionBatchID string) (*resty.Response, error) {
	//organizationID string,actionBatchID string
	path := "/api/v1/organizations/{organizationId}/actionBatches/{actionBatchId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{actionBatchId}", fmt.Sprintf("%v", actionBatchID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteOrganizationActionBatch")
	}

	return response, err

}

//DeleteOrganizationAdaptivePolicyACL Deletes the specified adaptive policy ACL
/* Deletes the specified adaptive policy ACL. Note this adaptive policy ACL will also be removed from policies using it.

@param organizationID organizationId path parameter. Organization ID
@param aclID aclId path parameter. Acl ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-organization-adaptive-policy-acl
*/
func (s *OrganizationsService) DeleteOrganizationAdaptivePolicyACL(organizationID string, aclID string) (*resty.Response, error) {
	//organizationID string,aclID string
	path := "/api/v1/organizations/{organizationId}/adaptivePolicy/acls/{aclId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{aclId}", fmt.Sprintf("%v", aclID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteOrganizationAdaptivePolicyAcl")
	}

	return response, err

}

//DeleteOrganizationAdaptivePolicyGroup Deletes the specified adaptive policy group and any associated policies and references
/* Deletes the specified adaptive policy group and any associated policies and references

@param organizationID organizationId path parameter. Organization ID
@param id id path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-organization-adaptive-policy-group
*/
func (s *OrganizationsService) DeleteOrganizationAdaptivePolicyGroup(organizationID string, id string) (*resty.Response, error) {
	//organizationID string,id string
	path := "/api/v1/organizations/{organizationId}/adaptivePolicy/groups/{id}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteOrganizationAdaptivePolicyGroup")
	}

	return response, err

}

//DeleteOrganizationAdaptivePolicyPolicy Delete an Adaptive Policy
/* Delete an Adaptive Policy

@param organizationID organizationId path parameter. Organization ID
@param id id path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-organization-adaptive-policy-policy
*/
func (s *OrganizationsService) DeleteOrganizationAdaptivePolicyPolicy(organizationID string, id string) (*resty.Response, error) {
	//organizationID string,id string
	path := "/api/v1/organizations/{organizationId}/adaptivePolicy/policies/{id}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteOrganizationAdaptivePolicyPolicy")
	}

	return response, err

}

//DeleteOrganizationAdmin Revoke all access for a dashboard administrator within this organization
/* Revoke all access for a dashboard administrator within this organization

@param organizationID organizationId path parameter. Organization ID
@param adminID adminId path parameter. Admin ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-organization-admin
*/
func (s *OrganizationsService) DeleteOrganizationAdmin(organizationID string, adminID string) (*resty.Response, error) {
	//organizationID string,adminID string
	path := "/api/v1/organizations/{organizationId}/admins/{adminId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{adminId}", fmt.Sprintf("%v", adminID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteOrganizationAdmin")
	}

	return response, err

}

//DeleteOrganizationAlertsProfile Removes an organization-wide alert config
/* Removes an organization-wide alert config

@param organizationID organizationId path parameter. Organization ID
@param alertConfigID alertConfigId path parameter. Alert config ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-organization-alerts-profile
*/
func (s *OrganizationsService) DeleteOrganizationAlertsProfile(organizationID string, alertConfigID string) (*resty.Response, error) {
	//organizationID string,alertConfigID string
	path := "/api/v1/organizations/{organizationId}/alerts/profiles/{alertConfigId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{alertConfigId}", fmt.Sprintf("%v", alertConfigID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteOrganizationAlertsProfile")
	}

	return response, err

}

//DeleteOrganizationBrandingPolicy Delete a branding policy
/* Delete a branding policy

@param organizationID organizationId path parameter. Organization ID
@param brandingPolicyID brandingPolicyId path parameter. Branding policy ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-organization-branding-policy
*/
func (s *OrganizationsService) DeleteOrganizationBrandingPolicy(organizationID string, brandingPolicyID string) (*resty.Response, error) {
	//organizationID string,brandingPolicyID string
	path := "/api/v1/organizations/{organizationId}/brandingPolicies/{brandingPolicyId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{brandingPolicyId}", fmt.Sprintf("%v", brandingPolicyID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteOrganizationBrandingPolicy")
	}

	return response, err

}

//DeleteOrganizationConfigTemplate Remove a configuration template
/* Remove a configuration template

@param organizationID organizationId path parameter. Organization ID
@param configTemplateID configTemplateId path parameter. Config template ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-organization-config-template
*/
func (s *OrganizationsService) DeleteOrganizationConfigTemplate(organizationID string, configTemplateID string) (*resty.Response, error) {
	//organizationID string,configTemplateID string
	path := "/api/v1/organizations/{organizationId}/configTemplates/{configTemplateId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{configTemplateId}", fmt.Sprintf("%v", configTemplateID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteOrganizationConfigTemplate")
	}

	return response, err

}

//DeleteOrganizationEarlyAccessFeaturesOptIn Delete an early access feature opt-in
/* Delete an early access feature opt-in

@param organizationID organizationId path parameter. Organization ID
@param optInID optInId path parameter. Opt in ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-organization-early-access-features-opt-in
*/
func (s *OrganizationsService) DeleteOrganizationEarlyAccessFeaturesOptIn(organizationID string, optInID string) (*resty.Response, error) {
	//organizationID string,optInID string
	path := "/api/v1/organizations/{organizationId}/earlyAccess/features/optIns/{optInId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{optInId}", fmt.Sprintf("%v", optInID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteOrganizationEarlyAccessFeaturesOptIn")
	}

	return response, err

}

//DeleteOrganizationPolicyObjectsGroup Deletes a Policy Object Group.
/* Deletes a Policy Object Group.

@param organizationID organizationId path parameter. Organization ID
@param policyObjectGroupID policyObjectGroupId path parameter. Policy object group ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-organization-policy-objects-group
*/
func (s *OrganizationsService) DeleteOrganizationPolicyObjectsGroup(organizationID string, policyObjectGroupID string) (*resty.Response, error) {
	//organizationID string,policyObjectGroupID string
	path := "/api/v1/organizations/{organizationId}/policyObjects/groups/{policyObjectGroupId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{policyObjectGroupId}", fmt.Sprintf("%v", policyObjectGroupID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteOrganizationPolicyObjectsGroup")
	}

	return response, err

}

//DeleteOrganizationPolicyObject Deletes a Policy Object.
/* Deletes a Policy Object.

@param organizationID organizationId path parameter. Organization ID
@param policyObjectID policyObjectId path parameter. Policy object ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-organization-policy-object
*/
func (s *OrganizationsService) DeleteOrganizationPolicyObject(organizationID string, policyObjectID string) (*resty.Response, error) {
	//organizationID string,policyObjectID string
	path := "/api/v1/organizations/{organizationId}/policyObjects/{policyObjectId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{policyObjectId}", fmt.Sprintf("%v", policyObjectID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteOrganizationPolicyObject")
	}

	return response, err

}

//DeleteOrganizationSamlIDp Remove a SAML IdP in your organization.
/* Remove a SAML IdP in your organization.

@param organizationID organizationId path parameter. Organization ID
@param idpID idpId path parameter. Idp ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-organization-saml-idp
*/
func (s *OrganizationsService) DeleteOrganizationSamlIDp(organizationID string, idpID string) (*resty.Response, error) {
	//organizationID string,idpID string
	path := "/api/v1/organizations/{organizationId}/saml/idps/{idpId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{idpId}", fmt.Sprintf("%v", idpID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteOrganizationSamlIdp")
	}

	return response, err

}

//DeleteOrganizationSamlRole Remove a SAML role
/* Remove a SAML role

@param organizationID organizationId path parameter. Organization ID
@param samlRoleID samlRoleId path parameter. Saml role ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-organization-saml-role
*/
func (s *OrganizationsService) DeleteOrganizationSamlRole(organizationID string, samlRoleID string) (*resty.Response, error) {
	//organizationID string,samlRoleID string
	path := "/api/v1/organizations/{organizationId}/samlRoles/{samlRoleId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{samlRoleId}", fmt.Sprintf("%v", samlRoleID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteOrganizationSamlRole")
	}

	return response, err

}

//DeleteOrganizationUser Delete a user and all of its authentication methods.
/* Delete a user and all of its authentication methods.

@param organizationID organizationId path parameter. Organization ID
@param userID userId path parameter. User ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-organization-user
*/
func (s *OrganizationsService) DeleteOrganizationUser(organizationID string, userID string) (*resty.Response, error) {
	//organizationID string,userID string
	path := "/api/v1/organizations/{organizationId}/users/{userId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{userId}", fmt.Sprintf("%v", userID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteOrganizationUser")
	}

	return response, err

}
