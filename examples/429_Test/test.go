package main

import (
	"fmt"

	meraki "github.com/meraki/dashboard-api-go/v5/sdk"
)

var client *meraki.Client

func main() {
	var err error
	fmt.Println("Authenticating")
	// client, err = meraki.NewClient()
	client, err = meraki.NewClientWithOptions("http://localhost:3006/api/v1", "6776676767676767676767676767676767676767", "false", "Test Cloverhound TEST")
	if err != nil {
		fmt.Println(err)
		return
	}
	// maxRetries := 3
	// maxRetryDelay := 5000 * time.Millisecond
	// maxRetryJitter := 5000 * time.Millisecond
	// useRetryHeader := false
	// client.SetBackoff(&maxRetries, &maxRetryDelay, &maxRetryJitter, &useRetryHeader)

	nResponse, _, err := client.Organizations.GetOrganizations(nil)
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
