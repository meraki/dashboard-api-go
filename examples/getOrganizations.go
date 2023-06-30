package main

import "github.com/meraki/dashboard-api-go/client"
import "context"
import "fmt"

func main() {
	configuration := client.NewConfiguration()

	apiClient := client.NewAPIClient(configuration)

	orgs, _, err := apiClient.ConfigureApi.GetOrganizations(context.Background()).Execute()

	if err != nil {
		fmt.Println("Meraki API call failed. Details: ", err)
	}

	for _, org := range orgs {
		fmt.Printf("%s\n", *org.Name)
	}
}
