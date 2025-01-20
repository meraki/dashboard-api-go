package main

import (
	"fmt"

	meraki "github.com/meraki/dashboard-api-go/v4/sdk"
)

// Client is DNA Center API client
var client *meraki.Client

func main() {
	var err error
	fmt.Println("Authenticating")
	client, err = meraki.NewClient()
	if err != nil {
		fmt.Println(err)
		return
	}

	nResponse, _, err := client.Administered.GetAdministeredIDentitiesMe()
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
