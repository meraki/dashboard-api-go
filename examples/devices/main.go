package main

import (
	"fmt"

	meraki "github.com/meraki/dashboard-api-go/v5/sdk"
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

	nResponse, _, err := client.Devices.GetDeviceClients("QBSC-ALSL-3GXN", nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	if nResponse != nil {
		// fmt.Println("\n Cantidad: ", len(*nResponse))
		fmt.Printf("%v", *nResponse)
		return
	}

	fmt.Println("There's no data on response")
}
