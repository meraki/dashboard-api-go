package main

import (
	"fmt"
	"os"

	meraki "github.com/meraki/dashboard-api-go/v2/sdk"
)

// Client is Dasboard Meraki API client
var client *meraki.Client

func main() {
	var err error
	fmt.Println("Authenticating")
	requestPerSecond := 5
	custom_user_agent := "customUA"
	client, err = meraki.NewClientWithOptions("https://api.meraki.com/",
		"12f2eb53588c75e28d89e108a05ea0c2487b08cf",
		"true", &custom_user_agent, &requestPerSecond)
	if err != nil {
		fmt.Println(err)
		return
	}

	nResponse, _, err := client.Administered.GetAdministeredIDentitiesMe()
	if err != nil {

		os.Exit(1)
		return
	}
	if nResponse != nil {
		fmt.Println(nResponse)

		return
	}
	nResponse2, _, err := client.Organizations.GetOrganization("828099381482762270")
	if err != nil {

		os.Exit(1)
		return
	}
	if nResponse2 != nil {
		fmt.Println(nResponse2)

		return
	}

}
