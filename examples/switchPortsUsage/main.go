package main

import (
	"fmt"

	meraki "github.com/meraki/dashboard-api-go/v4/sdk"
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

	nResponse, _, err := client.Switch.GetOrganizationSwitchPortsUsageHistoryByDeviceByInterval("orgId", &meraki.GetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalQueryParams{
		PerPage:  -1,
		Interval: 300,
		Timespan: 900,
	})

	if err != nil {
		fmt.Println(err)
		return
	}
	if nResponse != nil {
		fmt.Println("Total: ", len(*nResponse.Items))

		for _, i := range *nResponse.Items {
			for _, port := range *i.Ports {
				fmt.Printf("Switch: %s, Port: %s\n", i.Serial, port.PortID)
				for _, interval := range *port.Intervals {
					fmt.Printf("Interval: %v\n", interval.StartTs)
					fmt.Printf(" usage downstream - %v\n", interval.Data.Usage.Downstream)
					fmt.Printf(" usage upstream - %v\n", interval.Data.Usage.Upstream)
					fmt.Printf(" usage total - %v\n", interval.Data.Usage.Total)
					fmt.Printf(" bandwith downstream - %v\n", interval.Bandwidth.Usage.Downstream)
					fmt.Printf(" bandwith upstream - %v\n", interval.Bandwidth.Usage.Upstream)
					fmt.Printf(" bandwith total - %v\n", interval.Bandwidth.Usage.Total)

					// Check if Energy.Usage exists (since Energy itself is not a pointer)
					if interval.Energy.Usage != nil {
						fmt.Printf(" energy total - %v\n", interval.Energy.Usage.Total)
						fmt.Printf(" energy downstream - %v\n", interval.Energy.Usage.Downstream)
						fmt.Printf(" energy upstream - %v\n", interval.Energy.Usage.Upstream)
					} else {
						fmt.Println(" energy data not available")
					}
				}
			}
		}

		return
	}

	fmt.Println("There's no data on response")
}
