package main

import (
	"fmt"

	meraki "github.com/meraki/dashboard-api-go/v2/sdk"
)

var client *meraki.Client

func main() {
	var err error
	fmt.Println("Authenticating")
	client, err = meraki.NewClientWithOptions("https://api.meraki.com/",
		"12f2eb53588c75e28d89e108a05ea0c2487b08cf",
		"true")
	if err != nil {
		fmt.Println(err)
		return
	}

	nResponse, _, err := client.Organizations.GetOrganizations()
	if err != nil {
		fmt.Println(err)
		return
	}
	if nResponse != nil {
		fmt.Println(nResponse)
		return
	}
	fmt.Println("There's no data on response")
}
