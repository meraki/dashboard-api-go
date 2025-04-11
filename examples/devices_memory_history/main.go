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

	nResponse, _, err := client.Devices.GetOrganizationDevicesSystemMemoryUsageHistoryByInterval("orgId", &meraki.GetOrganizationDevicesSystemMemoryUsageHistoryByIntervalQueryParams{
		PerPage:  5,
		Timespan: 300,
	})

	if err != nil {
		fmt.Println(err)
		return
	}
	if nResponse != nil {
		fmt.Println("\n Cantidad: ", len(*nResponse.Items))

		// for _, i := range *nResponse.Items {
		// 	fmt.Println("ID: ", i.Name)
		// 	fmt.Println("Free median memory (kB): ", i.Free.Median)
		// 	fmt.Println("Used median memory (kB): ", i.Used.Median)
		// }

		return
	}

	fmt.Println("There's no data on response")
}
