package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/meraki/dashboard-api-go/client"
	"github.com/xuri/excelize/v2"
	"log"
	"os"
	"path/filepath"
	"sync"
)

type merakiApiClient struct {
	merakiApiClient *client.APIClient
	ApiKey          string
}

func newMerakiApiClient(apiKey string) *merakiApiClient {

	configuration := client.NewConfiguration()
	configuration.AddDefaultHeader("Authorization", "Bearer "+apiKey)
	apiClient := client.NewAPIClient(configuration)

	return &merakiApiClient{merakiApiClient: apiClient}
}

func (m *merakiApiClient) getOrganizations() ([]client.GetOrganizations200ResponseInner, error) {

	organizations, _, err := m.merakiApiClient.ConfigureApi.GetOrganizations(context.Background()).Execute()
	if err != nil {
		log.Println("failed to retrieve list of organization:", err)
		return nil, err
	}

	return organizations, nil
}

func (m *merakiApiClient) GetOrganizationNetworks(organizationId string) []client.GetNetwork200Response {

	inlineResp, _, err := m.merakiApiClient.ConfigureApi.GetOrganizationNetworks(context.Background(), organizationId).Execute()
	if err != nil {
		log.Println("failed to retrieve list of networks for organization:", organizationId)
		log.Println(err)
		return nil
	}

	return inlineResp
}

// OrganizationDevices - custom struct in lieu of schema objects
type OrganizationDevices struct {
	Name        string   `json:"name"`
	NetworkId   string   `json:"networkId"`
	Firmware    string   `json:"firmware"`
	LanIp       string   `json:"lanIp"`
	Mac         string   `json:"mac"`
	Model       string   `json:"model"`
	Serial      string   `json:"serial"`
	Lat         float32  `json:"lat"`
	Lng         float32  `json:"lng"`
	Address     string   `json:"address"`
	Tags        []string `json:"tags"`
	Notes       string   `json:"notes"`
	ProductType string   `json:"productType"`
}

func (m *merakiApiClient) getOrganizationDevices(organizationId string) []OrganizationDevices {
	_, httpResp, err := m.merakiApiClient.ConfigureApi.GetOrganizationDevices(context.Background(), organizationId).Execute()
	if err != nil {
		log.Println("failed to retrieve list of devices for organization:", organizationId)
		log.Println(err)
		return nil
	}

	var inlineResp []OrganizationDevices
	if err = json.NewDecoder(httpResp.Body).Decode(&inlineResp); err != nil {
		return nil
	}

	return inlineResp
}

func (m *merakiApiClient) getOrganizationLicences(organizationId string) []client.GetOrganizationLicenses200ResponseInner {
	inlineResp, _, err := m.merakiApiClient.ConfigureApi.GetOrganizationLicenses(context.Background(), organizationId).Execute()
	if err != nil {
		log.Println("failed to retrieve list of licences for organization:", organizationId)
		log.Println(err)
		return nil
	}

	return inlineResp
}

type tabData struct {
	Name string
	Rows [][]interface{}
}

func newSpreadsheet(tabsData []tabData) *excelize.File {
	f := excelize.NewFile()

	for _, tabData := range tabsData {
		sheetName := tabData.Name
		rows := tabData.Rows

		index, _ := f.NewSheet(sheetName)

		for i, row := range rows {
			for j, cellValue := range row {
				cellIndex, err := excelize.CoordinatesToCellName(j+1, i+1)
				if err != nil {
					log.Println(err)
				}
				err = f.SetCellValue(sheetName, cellIndex, cellValue)
				if err != nil {
					log.Println(err)
					return nil
				}
			}
		}

		f.SetActiveSheet(index)
	}

	// Delete the default Sheet1 tab
	err := f.DeleteSheet("Sheet1")
	if err != nil {
		return nil
	}

	return f
}

func addRowDataToTab(tab *tabData, data []interface{}) *tabData {

	// Append data to tab
	tab.Rows = append(tab.Rows, data)

	return tab
}

func generateOrganizationReport(wg *sync.WaitGroup, organization client.GetOrganizations200ResponseInner,
	m *merakiApiClient, outPath string) {

	// API calls for retrieving Organization data
	networks := m.GetOrganizationNetworks(organization.GetId())

	devices := m.getOrganizationDevices(organization.GetId())

	licences := m.getOrganizationLicences(organization.GetId())

	// Create Organizations Tab Data
	organizationsTab := tabData{
		Name: "Organizations",
		Rows: [][]interface{}{
			{"organizationId", "name", "url", "api enabled", "licensing model", "cloud region", "management details"},
		},
	}

	// Append organization information to Organization tab
	organizationsData := []interface{}{organization.GetId(), organization.GetName(),
		organization.GetUrl(), fmt.Sprintf("%v", organization.Api.GetEnabled()),
		organization.Licensing.GetModel(), organization.Cloud.Region.GetName(),
		fmt.Sprintf("%v", organization.Management.GetDetails())}

	addRowDataToTab(&organizationsTab, organizationsData)

	// Create Networks Tab Data
	networksTab := tabData{
		Name: "Networks",
		Rows: [][]interface{}{
			{"id", "organizationId", "name", "productTypes", "timeZone", "tags", "enrollmentString", "url", "notes", "isBoundToConfigTemplate"},
		},
	}

	// Append network information to Networks tab
	for _, network := range networks {

		networkData := []interface{}{network.GetId(), network.GetOrganizationId(), network.GetName(),
			fmt.Sprintf("%v", network.GetProductTypes()), network.GetTimeZone(),
			fmt.Sprintf("%v", network.GetTags()), network.GetEnrollmentString(),
			network.GetUrl(), network.GetNotes(), fmt.Sprintf("%v", network.GetIsBoundToConfigTemplate())}

		addRowDataToTab(&networksTab, networkData)

	}

	// Create Devices Tab Data
	devicesTab := tabData{
		Name: "Devices",
		Rows: [][]interface{}{
			{"name", "lat", "lng", "address", "notes", "tags", "networkId", "serial", "model", "mac", "lanIp", "firmware", "productType"},
		},
	}

	// Append device information to Devices tab
	for _, device := range devices {

		deviceData := []interface{}{device.Name, fmt.Sprintf("%v", device.Lat),
			fmt.Sprintf("%v", device.Lng), device.Address, device.Notes,
			fmt.Sprintf("%v", device.Tags), device.NetworkId, device.Serial, device.Model, device.Mac,
			device.LanIp, device.Firmware, device.ProductType}

		addRowDataToTab(&devicesTab, deviceData)
	}

	// Create Licenses Tab Data
	licensesTab := tabData{
		Name: "Licenses",
		Rows: [][]interface{}{
			{"id", "licenseType", "licenseKey", "orderNumber", "deviceSerial", "networkId", "state", "seatCount", "totalDurationInDays", "durationInDays", "claimDate", "activationDate", "expirationDate", "headLicenseId", "permanentlyQueuedLicenses"},
		},
	}

	// Append license information to Per-Device Licenses tab
	for _, licence := range licences {

		licenseData := []interface{}{licence.GetId(), licence.GetLicenseType(),
			licence.GetLicenseKey(), licence.GetOrderNumber(), licence.GetDeviceSerial(), licence.GetNetworkId(),
			licence.GetState(), fmt.Sprintf("%v", licence.GetSeatCount()),
			fmt.Sprintf("%v", licence.GetTotalDurationInDays()),
			fmt.Sprintf("%v", licence.GetDurationInDays()), licence.GetClaimDate(),
			licence.GetActivationDate(), licence.GetExpirationDate(), licence.GetHeadLicenseId(),
			fmt.Sprintf("%v", licence.GetPermanentlyQueuedLicenses())}

		addRowDataToTab(&licensesTab, licenseData)

	}

	// license model check
	if len(licensesTab.Rows) < 2 {
		licensesTab.Rows = append(licensesTab.Rows,
			[]interface{}{"Your organization is likely using a co-termination licensing model:"})
		licensesTab.Rows = append(licensesTab.Rows,
			[]interface{}{"https://documentation.meraki.com/General_Administration/Licensing/Meraki_Licensing_FAQs"})
	}

	// Create New spreadsheet
	tabsData := []tabData{organizationsTab, networksTab, devicesTab, licensesTab}
	f := newSpreadsheet(tabsData)

	// Save file
	err := f.SaveAs(filepath.Join(outPath, organization.GetName()+".xlsx"))
	if err != nil {
		log.Println(err)
	}

	// Mark the tab as completed
	log.Printf("Completed sheet for organization: %s\n", organization.GetName())

	wg.Done()

}

func main() {

	// Define the flags
	var apiKey string
	var outDir string
	var help bool

	// read flags
	flag.BoolVar(&help, "help", false, "Prints usage information")
	flag.StringVar(&apiKey, "apikey", os.Getenv("MERAKI_DASHBOARD_API_KEY"),
		"Meraki dashboard api key. This defaults to the env var: 'MERAKI_DASHBOARD_API_KEY'")
	flag.StringVar(&outDir, "out", "",
		"output directory to generate report. Defaults to current working directory.")

	// Define the usage function
	flag.Usage = func() {
		_, err := fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS]\n\nOptions:\n", os.Args[0])
		if err != nil {
			return
		}
		flag.PrintDefaults()
	}

	flag.Parse()

	// Check the help flag
	if help {
		flag.Usage()
		return
	}

	// check for valid api key
	if apiKey == "" {
		log.Printf("Did not find apikey in '-apikey' flag or environmental variable 'MERAKI_DASHBOARD_API_KEY'\n")
		log.Printf("Please ensure that you have enabled API access for your account in the Meraki Portal:\n")
		log.Printf("https://documentation.meraki.com/General_Administration/Other_Topics/Cisco_Meraki_Dashboard_API\n")
		return
	}

	// get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return
	}

	// set empty outDir to current working directory
	if outDir == "" {
		outDir = cwd
	}

	// check dir path exists
	if _, err := os.Stat(outDir); os.IsNotExist(err) {

		// create the directory path
		if err := os.Mkdir(outDir, 0755); err != nil {
			log.Printf("failed to create specified output directory:\n")
			log.Println(err)
			return
		}
		fmt.Printf("Created directory: %s\n", outDir)
	} else {
		fmt.Printf("Output directory exists: %s\n", outDir)
	}

	// initialize client
	m := newMerakiApiClient(apiKey)

	// retrieve list of organizations
	organizations, err := m.getOrganizations()
	if err != nil {
		log.Println(err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(len(organizations))

	// retrieve data
	for _, organization := range organizations {
		go generateOrganizationReport(&wg, organization, m, outDir)
	}

	wg.Wait()
}
