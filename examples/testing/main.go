package main

import (
	"fmt"
	"sync"

	meraki "github.com/meraki/dashboard-api-go/v5/sdk"
)

func main() {

	var err error
	client, err := meraki.NewClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		// Do work.
		getAdministeredIDentitiesMe(client)
		wg.Done()
	}()
	go func() {
		// Do work.
		getAdministeredIDentitiesMe(client)
		wg.Done()
	}()
	wg.Wait()

	// r, resp, err := client.Organizations.GetOrganizationFloorPlansAutoLocateDevices("828099381482762270", &meraki.GetOrganizationFloorPlansAutoLocateDevicesQueryParams{
	// 	PerPage: 100,
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// getAdministeredIDentitiesMe(client)
	// r, resp, err := client.Administered.GenerateAdministeredIDentitiesMeAPIKeysWithRetries()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// r, resp, err := client.Devices.BulkUpdateOrganizationDevicesDetailsWithRetries("828099381482762270", &meraki.RequestDevicesBulkUpdateOrganizationDevicesDetails{
	// 	Details: &[]meraki.RequestDevicesBulkUpdateOrganizationDevicesDetailsDetails{
	// 		{
	// 			Name:  "name",
	// 			Value: "value",
	// 		},
	// 	},
	// 	Serials: []string{"QBSC-ALSL-3GXN"},
	// })
	// fmt.Println(r)
	// fmt.Println(resp)
	// r, resp, err := client.Administered.GetAdministeredIDentitiesMeWithRetries()
	// getAdministeredIDentitiesMe(client)
}

func getAdministeredIDentitiesMe(client *meraki.Client) {
	for {
		r, resp, err := client.Administered.GetAdministeredIDentitiesMe()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(r)
		fmt.Println(resp)
	}

}
