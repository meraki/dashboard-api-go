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

	nResponse, _, err := client.Wireless.GetOrganizationWirelessDevicesSystemCpuLoadHistory("orgId", &meraki.GetOrganizationWirelessDevicesSystemCpuLoadHistoryParams{
		PerPage:  10,
		Timespan: 300,
	})

	if err != nil {
		fmt.Println(err)
		return
	}
	if nResponse != nil {
		fmt.Println("\n Cantidad: ", len(*nResponse.Items))

		for _, i := range *nResponse.Items {
			fmt.Println("ID: ", i.Name)
			if len(i.Series) > 0 {
				fmt.Println("Time: ", i.Series[0].TS)
				fmt.Println("CPU: ", i.Series[0].CPULoad5)
			} else {
				fmt.Println("cpu data not available")
			}
		}

		return
	}

	fmt.Println("There's no data on response")
}
