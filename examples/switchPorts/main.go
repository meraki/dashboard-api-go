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

	nResponse, _, err := client.Switch.GetOrganizationSwitchPortsStatusesBySwitch("orgId", &meraki.GetOrganizationSwitchPortsStatusesBySwitchQueryParams{
		PerPage: -1,
	})

	if err != nil {
		fmt.Println(err)
		return
	}
	if nResponse != nil {
		fmt.Println("Total: ", len(*nResponse.Items))
		// fmt.Printf("%v", *nResponse)

		for _, i := range *nResponse.Items {
			portsCount := 0
			for range *i.Ports {
				portsCount++
			}
			fmt.Printf("Switch: %s, Ports: %d\n", i.Serial, portsCount)

		}
		return
	}

	fmt.Println("There's no data on response")
}
