# dashboard-api-golang
Dashboard API for Go

The Meraki Dashboard API Golang library provides all current Meraki [dashboard API](https://developer.cisco.com/meraki/api-v1/) calls to interface with the Cisco Meraki cloud-managed platform.


## Requirements
[Go Programming Language](https://golang.org/doc/install) >= 1.16


## Installation

Clone the repository, set environmental variables

```shell
go get github.com/meraki/dashboard-api-go/client

export MERAKI_DASHBOARD_API_KEY="0123456789"

```

## Usage

Client API documentation is available in the /client/docs directory

### GetOrganizations Example

```go
package main

import "github.com/meraki/dashboard-api-go/client"
import "context"
import "fmt"
import "os"

func main() {
	configuration := client.NewConfiguration()
	configuration.AddDefaultHeader("X-Cisco-Meraki-API-Key", os.Getenv("MERAKI_DASHBOARD_API_KEY"))
	apiClient := client.NewAPIClient(configuration)

	orgs, _, err := apiClient.ConfigureApi.GetOrganizations(context.Background()).Execute()

	if err != nil {
		fmt.Println("Meraki API call failed. Details: ", err)
	}

	for _, org := range orgs {
		fmt.Printf("%s\n", *org.Name)
	}
}
```