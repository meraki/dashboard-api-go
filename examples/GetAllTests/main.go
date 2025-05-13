package main

import (
	"fmt"

	meraki "github.com/meraki/dashboard-api-go/v5/sdk"
)

var client *meraki.Client

func main() {
	var err error
	fmt.Println("Authenticating")
	client, err = meraki.NewClient()
	if err != nil {
		fmt.Println(err)
		return
	}

	nResponse, _, err := client.Organizations.GetOrganizationDevices("828099381482762270", &meraki.GetOrganizationDevicesQueryParams{
		PerPage:        -1,
		TagsFilterType: "withAnyTags",
	})
	// nResponse, _, err := client.Organizations.GetOrganizationAssuranceAlertsOverviewByNetwork("828099381482762270", &meraki.GetOrganizationAssuranceAlertsOverviewByNetworkQueryParams{
	// 	PerPage: -1,
	// })
	// nResponse, _, err := client.Devices.GetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice("828099381482762270", &meraki.GetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceQueryParams{
	// 	PerPage: -1,
	// })

	// nResponse, _, err := client.Switch.GetOrganizationSwitchPortsClientsOverviewByDevice("828099381482762270", &meraki.GetOrganizationSwitchPortsClientsOverviewByDeviceQueryParams{
	// 	PerPage: -1,
	// })
	// nResponse, _, err := client.Switch.GetOrganizationSwitchPortsTopologyDiscoveryByDevice("828099381482762270", &meraki.GetOrganizationSwitchPortsTopologyDiscoveryByDeviceQueryParams{
	// 	PerPage: -1,
	// })

	// nResponse, _, err := client.WirelessController.GetOrganizationWirelessControllerAvailabilitiesChangeHistory("828099381482762270", &meraki.GetOrganizationWirelessControllerAvailabilitiesChangeHistoryQueryParams{
	// 	PerPage: -1,
	// })
	// nResponse, _, err := client.WirelessController.GetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice("828099381482762270", &meraki.GetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceQueryParams{
	// 	PerPage: -1,
	// })

	// nResponse, _, err := client.Wireless.GetOrganizationWirelessAirMarshalRules("828099381482762270", &meraki.GetOrganizationWirelessAirMarshalRulesQueryParams{
	// 	PerPage: -1,
	// })

	if err != nil {
		fmt.Println(err)
		return
	}
	if nResponse != nil {
		fmt.Println("\n Cantidad: ", len(*nResponse))
		fmt.Printf("%v", *nResponse)
		return
	}

	fmt.Println("There's no data on response")
}
