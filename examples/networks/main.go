package main

import (
	"fmt"

	meraki "github.com/meraki/dashboard-api-go/v4/sdk"
)

var client *meraki.Client

func main() {
	var err error
	fmt.Println("Authenticating")
	client, err = meraki.NewClientWithOptions("https://api.meraki.com/",
		"12f2eb53588c75e28d89e108a05ea0c2487b08cf",
		"true", "AplicationName VendorName")
	if err != nil {
		fmt.Println(err)
		return
	}
	var a []string
	a = append(a, "appliance")
	a = append(a, "camera")
	re := &meraki.RequestOrganizationsCreateOrganizationNetwork{
		Name:         "Test2",
		ProductTypes: a,
	}
	fmt.Println(re)
	client.Organizations.CreateOrganizationNetwork("828099381482762270", re)

	// nResponse2, _, err := client.Organizations.GetOrganizationNetworks("828099381482762270", nil)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// if nResponse2 != nil {
	// 	fmt.Println(nResponse2)
	// }

	// for _, r := range *nResponse2 {
	// 	nResponse, _, err := client.Networks.GetNetworkClients(r.ID, nil)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	if nResponse != nil {
	// 		fmt.Println(nResponse)
	// 	}
	// 	fmt.Println("There's no data on response")
	// }

}
