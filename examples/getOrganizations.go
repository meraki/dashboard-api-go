package main

import "github.com/meraki/dashboard-api-go/client"
import "context"
import "fmt"
import "os"

func main() {
	configuration := client.NewConfiguration()
	configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))

	apiClient := client.NewAPIClient(configuration)

	orgs, _, err := apiClient.ConfigureApi.GetOrganizations(context.Background()).Execute()

	if err != nil {
		fmt.Println("Meraki API call failed. Details: ", err)
	}

	for _, org := range orgs {
		fmt.Printf("%s\n", *org.Name)
	}
}
