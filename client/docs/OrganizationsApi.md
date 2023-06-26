# \OrganizationsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignOrganizationLicensesSeats**](OrganizationsApi.md#AssignOrganizationLicensesSeats) | **Post** /organizations/{organizationId}/licenses/assignSeats | Assign SM seats to a network
[**ClaimIntoOrganization**](OrganizationsApi.md#ClaimIntoOrganization) | **Post** /organizations/{organizationId}/claim | Claim a list of devices, licenses, and/or orders into an organization
[**ClaimIntoOrganizationInventory**](OrganizationsApi.md#ClaimIntoOrganizationInventory) | **Post** /organizations/{organizationId}/inventory/claim | Claim a list of devices, licenses, and/or orders into an organization inventory
[**CloneOrganization**](OrganizationsApi.md#CloneOrganization) | **Post** /organizations/{organizationId}/clone | Create a new organization by cloning the addressed organization
[**CombineOrganizationNetworks**](OrganizationsApi.md#CombineOrganizationNetworks) | **Post** /organizations/{organizationId}/networks/combine | Combine multiple networks into a single network
[**CreateOrganization**](OrganizationsApi.md#CreateOrganization) | **Post** /organizations | Create a new organization
[**CreateOrganizationActionBatch**](OrganizationsApi.md#CreateOrganizationActionBatch) | **Post** /organizations/{organizationId}/actionBatches | Create an action batch
[**CreateOrganizationAdaptivePolicyAcl**](OrganizationsApi.md#CreateOrganizationAdaptivePolicyAcl) | **Post** /organizations/{organizationId}/adaptivePolicy/acls | Creates new adaptive policy ACL
[**CreateOrganizationAdaptivePolicyGroup**](OrganizationsApi.md#CreateOrganizationAdaptivePolicyGroup) | **Post** /organizations/{organizationId}/adaptivePolicy/groups | Creates a new adaptive policy group
[**CreateOrganizationAdaptivePolicyPolicy**](OrganizationsApi.md#CreateOrganizationAdaptivePolicyPolicy) | **Post** /organizations/{organizationId}/adaptivePolicy/policies | Add an Adaptive Policy
[**CreateOrganizationAdmin**](OrganizationsApi.md#CreateOrganizationAdmin) | **Post** /organizations/{organizationId}/admins | Create a new dashboard administrator
[**CreateOrganizationAlertsProfile**](OrganizationsApi.md#CreateOrganizationAlertsProfile) | **Post** /organizations/{organizationId}/alerts/profiles | Create an organization-wide alert configuration
[**CreateOrganizationBrandingPolicy**](OrganizationsApi.md#CreateOrganizationBrandingPolicy) | **Post** /organizations/{organizationId}/brandingPolicies | Add a new branding policy to an organization
[**CreateOrganizationConfigTemplate**](OrganizationsApi.md#CreateOrganizationConfigTemplate) | **Post** /organizations/{organizationId}/configTemplates | Create a new configuration template
[**CreateOrganizationEarlyAccessFeaturesOptIn**](OrganizationsApi.md#CreateOrganizationEarlyAccessFeaturesOptIn) | **Post** /organizations/{organizationId}/earlyAccess/features/optIns | Create a new early access feature opt-in for an organization
[**CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent**](OrganizationsApi.md#CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent) | **Post** /organizations/{organizationId}/inventory/onboarding/cloudMonitoring/exportEvents | Imports event logs related to the onboarding app into elastisearch
[**CreateOrganizationInventoryOnboardingCloudMonitoringImport**](OrganizationsApi.md#CreateOrganizationInventoryOnboardingCloudMonitoringImport) | **Post** /organizations/{organizationId}/inventory/onboarding/cloudMonitoring/imports | Commits the import operation to complete the onboarding of a device into Dashboard for monitoring.
[**CreateOrganizationInventoryOnboardingCloudMonitoringPrepare**](OrganizationsApi.md#CreateOrganizationInventoryOnboardingCloudMonitoringPrepare) | **Post** /organizations/{organizationId}/inventory/onboarding/cloudMonitoring/prepare | Initiates or updates an import session
[**CreateOrganizationNetwork**](OrganizationsApi.md#CreateOrganizationNetwork) | **Post** /organizations/{organizationId}/networks | Create a network
[**CreateOrganizationPolicyObject**](OrganizationsApi.md#CreateOrganizationPolicyObject) | **Post** /organizations/{organizationId}/policyObjects | Creates a new Policy Object.
[**CreateOrganizationPolicyObjectsGroup**](OrganizationsApi.md#CreateOrganizationPolicyObjectsGroup) | **Post** /organizations/{organizationId}/policyObjects/groups | Creates a new Policy Object Group.
[**CreateOrganizationSamlIdp**](OrganizationsApi.md#CreateOrganizationSamlIdp) | **Post** /organizations/{organizationId}/saml/idps | Create a SAML IdP for your organization.
[**CreateOrganizationSamlRole**](OrganizationsApi.md#CreateOrganizationSamlRole) | **Post** /organizations/{organizationId}/samlRoles | Create a SAML role
[**DeleteOrganization**](OrganizationsApi.md#DeleteOrganization) | **Delete** /organizations/{organizationId} | Delete an organization
[**DeleteOrganizationActionBatch**](OrganizationsApi.md#DeleteOrganizationActionBatch) | **Delete** /organizations/{organizationId}/actionBatches/{actionBatchId} | Delete an action batch
[**DeleteOrganizationAdaptivePolicyAcl**](OrganizationsApi.md#DeleteOrganizationAdaptivePolicyAcl) | **Delete** /organizations/{organizationId}/adaptivePolicy/acls/{aclId} | Deletes the specified adaptive policy ACL
[**DeleteOrganizationAdaptivePolicyGroup**](OrganizationsApi.md#DeleteOrganizationAdaptivePolicyGroup) | **Delete** /organizations/{organizationId}/adaptivePolicy/groups/{id} | Deletes the specified adaptive policy group and any associated policies and references
[**DeleteOrganizationAdaptivePolicyPolicy**](OrganizationsApi.md#DeleteOrganizationAdaptivePolicyPolicy) | **Delete** /organizations/{organizationId}/adaptivePolicy/policies/{id} | Delete an Adaptive Policy
[**DeleteOrganizationAdmin**](OrganizationsApi.md#DeleteOrganizationAdmin) | **Delete** /organizations/{organizationId}/admins/{adminId} | Revoke all access for a dashboard administrator within this organization
[**DeleteOrganizationAlertsProfile**](OrganizationsApi.md#DeleteOrganizationAlertsProfile) | **Delete** /organizations/{organizationId}/alerts/profiles/{alertConfigId} | Removes an organization-wide alert config
[**DeleteOrganizationBrandingPolicy**](OrganizationsApi.md#DeleteOrganizationBrandingPolicy) | **Delete** /organizations/{organizationId}/brandingPolicies/{brandingPolicyId} | Delete a branding policy
[**DeleteOrganizationConfigTemplate**](OrganizationsApi.md#DeleteOrganizationConfigTemplate) | **Delete** /organizations/{organizationId}/configTemplates/{configTemplateId} | Remove a configuration template
[**DeleteOrganizationEarlyAccessFeaturesOptIn**](OrganizationsApi.md#DeleteOrganizationEarlyAccessFeaturesOptIn) | **Delete** /organizations/{organizationId}/earlyAccess/features/optIns/{optInId} | Delete an early access feature opt-in
[**DeleteOrganizationPolicyObject**](OrganizationsApi.md#DeleteOrganizationPolicyObject) | **Delete** /organizations/{organizationId}/policyObjects/{policyObjectId} | Deletes a Policy Object.
[**DeleteOrganizationPolicyObjectsGroup**](OrganizationsApi.md#DeleteOrganizationPolicyObjectsGroup) | **Delete** /organizations/{organizationId}/policyObjects/groups/{policyObjectGroupId} | Deletes a Policy Object Group.
[**DeleteOrganizationSamlIdp**](OrganizationsApi.md#DeleteOrganizationSamlIdp) | **Delete** /organizations/{organizationId}/saml/idps/{idpId} | Remove a SAML IdP in your organization.
[**DeleteOrganizationSamlRole**](OrganizationsApi.md#DeleteOrganizationSamlRole) | **Delete** /organizations/{organizationId}/samlRoles/{samlRoleId} | Remove a SAML role
[**DeleteOrganizationUser**](OrganizationsApi.md#DeleteOrganizationUser) | **Delete** /organizations/{organizationId}/users/{userId} | Delete a user and all of its authentication methods.
[**GetOrganization**](OrganizationsApi.md#GetOrganization) | **Get** /organizations/{organizationId} | Return an organization
[**GetOrganizationActionBatch**](OrganizationsApi.md#GetOrganizationActionBatch) | **Get** /organizations/{organizationId}/actionBatches/{actionBatchId} | Return an action batch
[**GetOrganizationActionBatches**](OrganizationsApi.md#GetOrganizationActionBatches) | **Get** /organizations/{organizationId}/actionBatches | Return the list of action batches in the organization
[**GetOrganizationAdaptivePolicyAcl**](OrganizationsApi.md#GetOrganizationAdaptivePolicyAcl) | **Get** /organizations/{organizationId}/adaptivePolicy/acls/{aclId} | Returns the adaptive policy ACL information
[**GetOrganizationAdaptivePolicyAcls**](OrganizationsApi.md#GetOrganizationAdaptivePolicyAcls) | **Get** /organizations/{organizationId}/adaptivePolicy/acls | List adaptive policy ACLs in a organization
[**GetOrganizationAdaptivePolicyGroup**](OrganizationsApi.md#GetOrganizationAdaptivePolicyGroup) | **Get** /organizations/{organizationId}/adaptivePolicy/groups/{id} | Returns an adaptive policy group
[**GetOrganizationAdaptivePolicyGroups**](OrganizationsApi.md#GetOrganizationAdaptivePolicyGroups) | **Get** /organizations/{organizationId}/adaptivePolicy/groups | List adaptive policy groups in a organization
[**GetOrganizationAdaptivePolicyOverview**](OrganizationsApi.md#GetOrganizationAdaptivePolicyOverview) | **Get** /organizations/{organizationId}/adaptivePolicy/overview | Returns adaptive policy aggregate statistics for an organization
[**GetOrganizationAdaptivePolicyPolicies**](OrganizationsApi.md#GetOrganizationAdaptivePolicyPolicies) | **Get** /organizations/{organizationId}/adaptivePolicy/policies | List adaptive policies in an organization
[**GetOrganizationAdaptivePolicyPolicy**](OrganizationsApi.md#GetOrganizationAdaptivePolicyPolicy) | **Get** /organizations/{organizationId}/adaptivePolicy/policies/{id} | Return an adaptive policy
[**GetOrganizationAdaptivePolicySettings**](OrganizationsApi.md#GetOrganizationAdaptivePolicySettings) | **Get** /organizations/{organizationId}/adaptivePolicy/settings | Returns global adaptive policy settings in an organization
[**GetOrganizationAdmins**](OrganizationsApi.md#GetOrganizationAdmins) | **Get** /organizations/{organizationId}/admins | List the dashboard administrators in this organization
[**GetOrganizationAlertsProfiles**](OrganizationsApi.md#GetOrganizationAlertsProfiles) | **Get** /organizations/{organizationId}/alerts/profiles | List all organization-wide alert configurations
[**GetOrganizationApiRequests**](OrganizationsApi.md#GetOrganizationApiRequests) | **Get** /organizations/{organizationId}/apiRequests | List the API requests made by an organization
[**GetOrganizationApiRequestsOverview**](OrganizationsApi.md#GetOrganizationApiRequestsOverview) | **Get** /organizations/{organizationId}/apiRequests/overview | Return an aggregated overview of API requests data
[**GetOrganizationApiRequestsOverviewResponseCodesByInterval**](OrganizationsApi.md#GetOrganizationApiRequestsOverviewResponseCodesByInterval) | **Get** /organizations/{organizationId}/apiRequests/overview/responseCodes/byInterval | Tracks organizations&#39; API requests by response code across a given time period
[**GetOrganizationBrandingPolicies**](OrganizationsApi.md#GetOrganizationBrandingPolicies) | **Get** /organizations/{organizationId}/brandingPolicies | List the branding policies of an organization
[**GetOrganizationBrandingPoliciesPriorities**](OrganizationsApi.md#GetOrganizationBrandingPoliciesPriorities) | **Get** /organizations/{organizationId}/brandingPolicies/priorities | Return the branding policy IDs of an organization in priority order
[**GetOrganizationBrandingPolicy**](OrganizationsApi.md#GetOrganizationBrandingPolicy) | **Get** /organizations/{organizationId}/brandingPolicies/{brandingPolicyId} | Return a branding policy
[**GetOrganizationClientsBandwidthUsageHistory**](OrganizationsApi.md#GetOrganizationClientsBandwidthUsageHistory) | **Get** /organizations/{organizationId}/clients/bandwidthUsageHistory | Return data usage (in megabits per second) over time for all clients in the given organization within a given time range.
[**GetOrganizationClientsOverview**](OrganizationsApi.md#GetOrganizationClientsOverview) | **Get** /organizations/{organizationId}/clients/overview | Return summary information around client data usage (in mb) across the given organization.
[**GetOrganizationClientsSearch**](OrganizationsApi.md#GetOrganizationClientsSearch) | **Get** /organizations/{organizationId}/clients/search | Return the client details in an organization
[**GetOrganizationConfigTemplate**](OrganizationsApi.md#GetOrganizationConfigTemplate) | **Get** /organizations/{organizationId}/configTemplates/{configTemplateId} | Return a single configuration template
[**GetOrganizationConfigTemplates**](OrganizationsApi.md#GetOrganizationConfigTemplates) | **Get** /organizations/{organizationId}/configTemplates | List the configuration templates for this organization
[**GetOrganizationConfigurationChanges**](OrganizationsApi.md#GetOrganizationConfigurationChanges) | **Get** /organizations/{organizationId}/configurationChanges | View the Change Log for your organization
[**GetOrganizationDevices**](OrganizationsApi.md#GetOrganizationDevices) | **Get** /organizations/{organizationId}/devices | List the devices in an organization
[**GetOrganizationDevicesAvailabilities**](OrganizationsApi.md#GetOrganizationDevicesAvailabilities) | **Get** /organizations/{organizationId}/devices/availabilities | List the availability information for devices in an organization
[**GetOrganizationDevicesPowerModulesStatusesByDevice**](OrganizationsApi.md#GetOrganizationDevicesPowerModulesStatusesByDevice) | **Get** /organizations/{organizationId}/devices/powerModules/statuses/byDevice | List the power status information for devices in an organization
[**GetOrganizationDevicesProvisioningStatuses**](OrganizationsApi.md#GetOrganizationDevicesProvisioningStatuses) | **Get** /organizations/{organizationId}/devices/provisioning/statuses | List the provisioning statuses information for devices in an organization.
[**GetOrganizationDevicesStatuses**](OrganizationsApi.md#GetOrganizationDevicesStatuses) | **Get** /organizations/{organizationId}/devices/statuses | List the status of every Meraki device in the organization
[**GetOrganizationDevicesStatusesOverview**](OrganizationsApi.md#GetOrganizationDevicesStatusesOverview) | **Get** /organizations/{organizationId}/devices/statuses/overview | Return an overview of current device statuses
[**GetOrganizationDevicesUplinksAddressesByDevice**](OrganizationsApi.md#GetOrganizationDevicesUplinksAddressesByDevice) | **Get** /organizations/{organizationId}/devices/uplinks/addresses/byDevice | List the current uplink addresses for devices in an organization.
[**GetOrganizationDevicesUplinksLossAndLatency**](OrganizationsApi.md#GetOrganizationDevicesUplinksLossAndLatency) | **Get** /organizations/{organizationId}/devices/uplinksLossAndLatency | Return the uplink loss and latency for every MX in the organization from at latest 2 minutes ago
[**GetOrganizationEarlyAccessFeatures**](OrganizationsApi.md#GetOrganizationEarlyAccessFeatures) | **Get** /organizations/{organizationId}/earlyAccess/features | List the available early access features for organization
[**GetOrganizationEarlyAccessFeaturesOptIn**](OrganizationsApi.md#GetOrganizationEarlyAccessFeaturesOptIn) | **Get** /organizations/{organizationId}/earlyAccess/features/optIns/{optInId} | Show an early access feature opt-in for an organization
[**GetOrganizationEarlyAccessFeaturesOptIns**](OrganizationsApi.md#GetOrganizationEarlyAccessFeaturesOptIns) | **Get** /organizations/{organizationId}/earlyAccess/features/optIns | List the early access feature opt-ins for an organization
[**GetOrganizationFirmwareUpgrades**](OrganizationsApi.md#GetOrganizationFirmwareUpgrades) | **Get** /organizations/{organizationId}/firmware/upgrades | Get firmware upgrade information for an organization
[**GetOrganizationFirmwareUpgradesByDevice**](OrganizationsApi.md#GetOrganizationFirmwareUpgradesByDevice) | **Get** /organizations/{organizationId}/firmware/upgrades/byDevice | Get firmware upgrade status for the filtered devices
[**GetOrganizationInventoryDevice**](OrganizationsApi.md#GetOrganizationInventoryDevice) | **Get** /organizations/{organizationId}/inventory/devices/{serial} | Return a single device from the inventory of an organization
[**GetOrganizationInventoryDevices**](OrganizationsApi.md#GetOrganizationInventoryDevices) | **Get** /organizations/{organizationId}/inventory/devices | Return the device inventory for an organization
[**GetOrganizationInventoryOnboardingCloudMonitoringImports**](OrganizationsApi.md#GetOrganizationInventoryOnboardingCloudMonitoringImports) | **Get** /organizations/{organizationId}/inventory/onboarding/cloudMonitoring/imports | Check the status of a committed Import operation
[**GetOrganizationInventoryOnboardingCloudMonitoringNetworks**](OrganizationsApi.md#GetOrganizationInventoryOnboardingCloudMonitoringNetworks) | **Get** /organizations/{organizationId}/inventory/onboarding/cloudMonitoring/networks | Returns list of networks eligible for adding cloud monitored device
[**GetOrganizationLicense**](OrganizationsApi.md#GetOrganizationLicense) | **Get** /organizations/{organizationId}/licenses/{licenseId} | Display a license
[**GetOrganizationLicenses**](OrganizationsApi.md#GetOrganizationLicenses) | **Get** /organizations/{organizationId}/licenses | List the licenses for an organization
[**GetOrganizationLicensesOverview**](OrganizationsApi.md#GetOrganizationLicensesOverview) | **Get** /organizations/{organizationId}/licenses/overview | Return an overview of the license state for an organization
[**GetOrganizationLoginSecurity**](OrganizationsApi.md#GetOrganizationLoginSecurity) | **Get** /organizations/{organizationId}/loginSecurity | Returns the login security settings for an organization.
[**GetOrganizationNetworks**](OrganizationsApi.md#GetOrganizationNetworks) | **Get** /organizations/{organizationId}/networks | List the networks that the user has privileges on in an organization
[**GetOrganizationOpenapiSpec**](OrganizationsApi.md#GetOrganizationOpenapiSpec) | **Get** /organizations/{organizationId}/openapiSpec | Return the OpenAPI Specification of the organization&#39;s API documentation in JSON
[**GetOrganizationPolicyObject**](OrganizationsApi.md#GetOrganizationPolicyObject) | **Get** /organizations/{organizationId}/policyObjects/{policyObjectId} | Shows details of a Policy Object.
[**GetOrganizationPolicyObjects**](OrganizationsApi.md#GetOrganizationPolicyObjects) | **Get** /organizations/{organizationId}/policyObjects | Lists Policy Objects belonging to the organization.
[**GetOrganizationPolicyObjectsGroup**](OrganizationsApi.md#GetOrganizationPolicyObjectsGroup) | **Get** /organizations/{organizationId}/policyObjects/groups/{policyObjectGroupId} | Shows details of a Policy Object Group.
[**GetOrganizationPolicyObjectsGroups**](OrganizationsApi.md#GetOrganizationPolicyObjectsGroups) | **Get** /organizations/{organizationId}/policyObjects/groups | Lists Policy Object Groups belonging to the organization.
[**GetOrganizationSaml**](OrganizationsApi.md#GetOrganizationSaml) | **Get** /organizations/{organizationId}/saml | Returns the SAML SSO enabled settings for an organization.
[**GetOrganizationSamlIdp**](OrganizationsApi.md#GetOrganizationSamlIdp) | **Get** /organizations/{organizationId}/saml/idps/{idpId} | Get a SAML IdP from your organization.
[**GetOrganizationSamlIdps**](OrganizationsApi.md#GetOrganizationSamlIdps) | **Get** /organizations/{organizationId}/saml/idps | List the SAML IdPs in your organization.
[**GetOrganizationSamlRole**](OrganizationsApi.md#GetOrganizationSamlRole) | **Get** /organizations/{organizationId}/samlRoles/{samlRoleId} | Return a SAML role
[**GetOrganizationSamlRoles**](OrganizationsApi.md#GetOrganizationSamlRoles) | **Get** /organizations/{organizationId}/samlRoles | List the SAML roles for this organization
[**GetOrganizationSnmp**](OrganizationsApi.md#GetOrganizationSnmp) | **Get** /organizations/{organizationId}/snmp | Return the SNMP settings for an organization
[**GetOrganizationSummaryTopAppliancesByUtilization**](OrganizationsApi.md#GetOrganizationSummaryTopAppliancesByUtilization) | **Get** /organizations/{organizationId}/summary/top/appliances/byUtilization | Return the top 10 appliances sorted by utilization over given time range.
[**GetOrganizationSummaryTopClientsByUsage**](OrganizationsApi.md#GetOrganizationSummaryTopClientsByUsage) | **Get** /organizations/{organizationId}/summary/top/clients/byUsage | Return metrics for organization&#39;s top 10 clients by data usage (in mb) over given time range.
[**GetOrganizationSummaryTopClientsManufacturersByUsage**](OrganizationsApi.md#GetOrganizationSummaryTopClientsManufacturersByUsage) | **Get** /organizations/{organizationId}/summary/top/clients/manufacturers/byUsage | Return metrics for organization&#39;s top clients by data usage (in mb) over given time range, grouped by manufacturer.
[**GetOrganizationSummaryTopDevicesByUsage**](OrganizationsApi.md#GetOrganizationSummaryTopDevicesByUsage) | **Get** /organizations/{organizationId}/summary/top/devices/byUsage | Return metrics for organization&#39;s top 10 devices sorted by data usage over given time range
[**GetOrganizationSummaryTopDevicesModelsByUsage**](OrganizationsApi.md#GetOrganizationSummaryTopDevicesModelsByUsage) | **Get** /organizations/{organizationId}/summary/top/devices/models/byUsage | Return metrics for organization&#39;s top 10 device models sorted by data usage over given time range
[**GetOrganizationSummaryTopSsidsByUsage**](OrganizationsApi.md#GetOrganizationSummaryTopSsidsByUsage) | **Get** /organizations/{organizationId}/summary/top/ssids/byUsage | Return metrics for organization&#39;s top 10 ssids by data usage over given time range
[**GetOrganizationSummaryTopSwitchesByEnergyUsage**](OrganizationsApi.md#GetOrganizationSummaryTopSwitchesByEnergyUsage) | **Get** /organizations/{organizationId}/summary/top/switches/byEnergyUsage | Return metrics for organization&#39;s top 10 switches by energy usage over given time range
[**GetOrganizationUplinksStatuses**](OrganizationsApi.md#GetOrganizationUplinksStatuses) | **Get** /organizations/{organizationId}/uplinks/statuses | List the uplink status of every Meraki MX, MG and Z series devices in the organization
[**GetOrganizationWebhooksAlertTypes**](OrganizationsApi.md#GetOrganizationWebhooksAlertTypes) | **Get** /organizations/{organizationId}/webhooks/alertTypes | Return a list of alert types to be used with managing webhook alerts
[**GetOrganizationWebhooksLogs**](OrganizationsApi.md#GetOrganizationWebhooksLogs) | **Get** /organizations/{organizationId}/webhooks/logs | Return the log of webhook POSTs sent
[**GetOrganizations**](OrganizationsApi.md#GetOrganizations) | **Get** /organizations | List the organizations that the user has privileges on
[**MoveOrganizationLicenses**](OrganizationsApi.md#MoveOrganizationLicenses) | **Post** /organizations/{organizationId}/licenses/move | Move licenses to another organization
[**MoveOrganizationLicensesSeats**](OrganizationsApi.md#MoveOrganizationLicensesSeats) | **Post** /organizations/{organizationId}/licenses/moveSeats | Move SM seats to another organization
[**ReleaseFromOrganizationInventory**](OrganizationsApi.md#ReleaseFromOrganizationInventory) | **Post** /organizations/{organizationId}/inventory/release | Release a list of claimed devices from an organization.
[**RenewOrganizationLicensesSeats**](OrganizationsApi.md#RenewOrganizationLicensesSeats) | **Post** /organizations/{organizationId}/licenses/renewSeats | Renew SM seats of a license
[**UpdateOrganization**](OrganizationsApi.md#UpdateOrganization) | **Put** /organizations/{organizationId} | Update an organization
[**UpdateOrganizationActionBatch**](OrganizationsApi.md#UpdateOrganizationActionBatch) | **Put** /organizations/{organizationId}/actionBatches/{actionBatchId} | Update an action batch
[**UpdateOrganizationAdaptivePolicyAcl**](OrganizationsApi.md#UpdateOrganizationAdaptivePolicyAcl) | **Put** /organizations/{organizationId}/adaptivePolicy/acls/{aclId} | Updates an adaptive policy ACL
[**UpdateOrganizationAdaptivePolicyGroup**](OrganizationsApi.md#UpdateOrganizationAdaptivePolicyGroup) | **Put** /organizations/{organizationId}/adaptivePolicy/groups/{id} | Updates an adaptive policy group
[**UpdateOrganizationAdaptivePolicyPolicy**](OrganizationsApi.md#UpdateOrganizationAdaptivePolicyPolicy) | **Put** /organizations/{organizationId}/adaptivePolicy/policies/{id} | Update an Adaptive Policy
[**UpdateOrganizationAdaptivePolicySettings**](OrganizationsApi.md#UpdateOrganizationAdaptivePolicySettings) | **Put** /organizations/{organizationId}/adaptivePolicy/settings | Update global adaptive policy settings
[**UpdateOrganizationAdmin**](OrganizationsApi.md#UpdateOrganizationAdmin) | **Put** /organizations/{organizationId}/admins/{adminId} | Update an administrator
[**UpdateOrganizationAlertsProfile**](OrganizationsApi.md#UpdateOrganizationAlertsProfile) | **Put** /organizations/{organizationId}/alerts/profiles/{alertConfigId} | Update an organization-wide alert config
[**UpdateOrganizationBrandingPoliciesPriorities**](OrganizationsApi.md#UpdateOrganizationBrandingPoliciesPriorities) | **Put** /organizations/{organizationId}/brandingPolicies/priorities | Update the priority ordering of an organization&#39;s branding policies.
[**UpdateOrganizationBrandingPolicy**](OrganizationsApi.md#UpdateOrganizationBrandingPolicy) | **Put** /organizations/{organizationId}/brandingPolicies/{brandingPolicyId} | Update a branding policy
[**UpdateOrganizationConfigTemplate**](OrganizationsApi.md#UpdateOrganizationConfigTemplate) | **Put** /organizations/{organizationId}/configTemplates/{configTemplateId} | Update a configuration template
[**UpdateOrganizationEarlyAccessFeaturesOptIn**](OrganizationsApi.md#UpdateOrganizationEarlyAccessFeaturesOptIn) | **Put** /organizations/{organizationId}/earlyAccess/features/optIns/{optInId} | Update an early access feature opt-in for an organization
[**UpdateOrganizationLicense**](OrganizationsApi.md#UpdateOrganizationLicense) | **Put** /organizations/{organizationId}/licenses/{licenseId} | Update a license
[**UpdateOrganizationLoginSecurity**](OrganizationsApi.md#UpdateOrganizationLoginSecurity) | **Put** /organizations/{organizationId}/loginSecurity | Update the login security settings for an organization
[**UpdateOrganizationPolicyObject**](OrganizationsApi.md#UpdateOrganizationPolicyObject) | **Put** /organizations/{organizationId}/policyObjects/{policyObjectId} | Updates a Policy Object.
[**UpdateOrganizationPolicyObjectsGroup**](OrganizationsApi.md#UpdateOrganizationPolicyObjectsGroup) | **Put** /organizations/{organizationId}/policyObjects/groups/{policyObjectGroupId} | Updates a Policy Object Group.
[**UpdateOrganizationSaml**](OrganizationsApi.md#UpdateOrganizationSaml) | **Put** /organizations/{organizationId}/saml | Updates the SAML SSO enabled settings for an organization.
[**UpdateOrganizationSamlIdp**](OrganizationsApi.md#UpdateOrganizationSamlIdp) | **Put** /organizations/{organizationId}/saml/idps/{idpId} | Update a SAML IdP in your organization
[**UpdateOrganizationSamlRole**](OrganizationsApi.md#UpdateOrganizationSamlRole) | **Put** /organizations/{organizationId}/samlRoles/{samlRoleId} | Update a SAML role
[**UpdateOrganizationSnmp**](OrganizationsApi.md#UpdateOrganizationSnmp) | **Put** /organizations/{organizationId}/snmp | Update the SNMP settings for an organization



## AssignOrganizationLicensesSeats

> InlineResponse200131 AssignOrganizationLicensesSeats(ctx, organizationId).AssignOrganizationLicensesSeats(assignOrganizationLicensesSeats).Execute()

Assign SM seats to a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    assignOrganizationLicensesSeats := *openapiclient.NewInlineObject207("LicenseId_example", "NetworkId_example", int32(123)) // InlineObject207 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.AssignOrganizationLicensesSeats(context.Background(), organizationId).AssignOrganizationLicensesSeats(assignOrganizationLicensesSeats).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.AssignOrganizationLicensesSeats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignOrganizationLicensesSeats`: InlineResponse200131
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.AssignOrganizationLicensesSeats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignOrganizationLicensesSeatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assignOrganizationLicensesSeats** | [**InlineObject207**](InlineObject207.md) |  | 

### Return type

[**InlineResponse200131**](InlineResponse200131.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClaimIntoOrganization

> map[string]interface{} ClaimIntoOrganization(ctx, organizationId).ClaimIntoOrganization(claimIntoOrganization).Execute()

Claim a list of devices, licenses, and/or orders into an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    claimIntoOrganization := *openapiclient.NewInlineObject193() // InlineObject193 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.ClaimIntoOrganization(context.Background(), organizationId).ClaimIntoOrganization(claimIntoOrganization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.ClaimIntoOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClaimIntoOrganization`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.ClaimIntoOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiClaimIntoOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **claimIntoOrganization** | [**InlineObject193**](InlineObject193.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClaimIntoOrganizationInventory

> map[string]interface{} ClaimIntoOrganizationInventory(ctx, organizationId).ClaimIntoOrganizationInventory(claimIntoOrganizationInventory).Execute()

Claim a list of devices, licenses, and/or orders into an organization inventory



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    claimIntoOrganizationInventory := *openapiclient.NewInlineObject202() // InlineObject202 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.ClaimIntoOrganizationInventory(context.Background(), organizationId).ClaimIntoOrganizationInventory(claimIntoOrganizationInventory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.ClaimIntoOrganizationInventory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClaimIntoOrganizationInventory`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.ClaimIntoOrganizationInventory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiClaimIntoOrganizationInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **claimIntoOrganizationInventory** | [**InlineObject202**](InlineObject202.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloneOrganization

> InlineResponse20099 CloneOrganization(ctx, organizationId).CloneOrganization(cloneOrganization).Execute()

Create a new organization by cloning the addressed organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    cloneOrganization := *openapiclient.NewInlineObject194("Name_example") // InlineObject194 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.CloneOrganization(context.Background(), organizationId).CloneOrganization(cloneOrganization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.CloneOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloneOrganization`: InlineResponse20099
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.CloneOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloneOrganization** | [**InlineObject194**](InlineObject194.md) |  | 

### Return type

[**InlineResponse20099**](InlineResponse20099.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CombineOrganizationNetworks

> InlineResponse200137 CombineOrganizationNetworks(ctx, organizationId).CombineOrganizationNetworks(combineOrganizationNetworks).Execute()

Combine multiple networks into a single network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    combineOrganizationNetworks := *openapiclient.NewInlineObject215("Name_example", []string{"NetworkIds_example"}) // InlineObject215 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.CombineOrganizationNetworks(context.Background(), organizationId).CombineOrganizationNetworks(combineOrganizationNetworks).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.CombineOrganizationNetworks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CombineOrganizationNetworks`: InlineResponse200137
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.CombineOrganizationNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCombineOrganizationNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **combineOrganizationNetworks** | [**InlineObject215**](InlineObject215.md) |  | 

### Return type

[**InlineResponse200137**](InlineResponse200137.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganization

> InlineResponse20099 CreateOrganization(ctx).CreateOrganization(createOrganization).Execute()

Create a new organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createOrganization := *openapiclient.NewInlineObject170("Name_example") // InlineObject170 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.CreateOrganization(context.Background()).CreateOrganization(createOrganization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.CreateOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganization`: InlineResponse20099
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.CreateOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createOrganization** | [**InlineObject170**](InlineObject170.md) |  | 

### Return type

[**InlineResponse20099**](InlineResponse20099.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationActionBatch

> InlineResponse2014 CreateOrganizationActionBatch(ctx, organizationId).CreateOrganizationActionBatch(createOrganizationActionBatch).Execute()

Create an action batch



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    createOrganizationActionBatch := *openapiclient.NewInlineObject172([]openapiclient.OrganizationsOrganizationIdActionBatchesActions{*openapiclient.NewOrganizationsOrganizationIdActionBatchesActions("Resource_example", "Operation_example")}) // InlineObject172 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.CreateOrganizationActionBatch(context.Background(), organizationId).CreateOrganizationActionBatch(createOrganizationActionBatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.CreateOrganizationActionBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationActionBatch`: InlineResponse2014
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.CreateOrganizationActionBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationActionBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationActionBatch** | [**InlineObject172**](InlineObject172.md) |  | 

### Return type

[**InlineResponse2014**](InlineResponse2014.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationAdaptivePolicyAcl

> InlineResponse200100 CreateOrganizationAdaptivePolicyAcl(ctx, organizationId).CreateOrganizationAdaptivePolicyAcl(createOrganizationAdaptivePolicyAcl).Execute()

Creates new adaptive policy ACL



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    createOrganizationAdaptivePolicyAcl := *openapiclient.NewInlineObject174("Name_example", []openapiclient.OrganizationsOrganizationIdAdaptivePolicyAclsRules1{*openapiclient.NewOrganizationsOrganizationIdAdaptivePolicyAclsRules1("Policy_example", "Protocol_example")}, "IpVersion_example") // InlineObject174 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.CreateOrganizationAdaptivePolicyAcl(context.Background(), organizationId).CreateOrganizationAdaptivePolicyAcl(createOrganizationAdaptivePolicyAcl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.CreateOrganizationAdaptivePolicyAcl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationAdaptivePolicyAcl`: InlineResponse200100
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.CreateOrganizationAdaptivePolicyAcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationAdaptivePolicyAclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationAdaptivePolicyAcl** | [**InlineObject174**](InlineObject174.md) |  | 

### Return type

[**InlineResponse200100**](InlineResponse200100.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationAdaptivePolicyGroup

> map[string]interface{} CreateOrganizationAdaptivePolicyGroup(ctx, organizationId).CreateOrganizationAdaptivePolicyGroup(createOrganizationAdaptivePolicyGroup).Execute()

Creates a new adaptive policy group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    createOrganizationAdaptivePolicyGroup := *openapiclient.NewInlineObject176("Name_example", int32(123)) // InlineObject176 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.CreateOrganizationAdaptivePolicyGroup(context.Background(), organizationId).CreateOrganizationAdaptivePolicyGroup(createOrganizationAdaptivePolicyGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.CreateOrganizationAdaptivePolicyGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationAdaptivePolicyGroup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.CreateOrganizationAdaptivePolicyGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationAdaptivePolicyGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationAdaptivePolicyGroup** | [**InlineObject176**](InlineObject176.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationAdaptivePolicyPolicy

> map[string]interface{} CreateOrganizationAdaptivePolicyPolicy(ctx, organizationId).CreateOrganizationAdaptivePolicyPolicy(createOrganizationAdaptivePolicyPolicy).Execute()

Add an Adaptive Policy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    createOrganizationAdaptivePolicyPolicy := *openapiclient.NewInlineObject178(*openapiclient.NewOrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup(), *openapiclient.NewOrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup()) // InlineObject178 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.CreateOrganizationAdaptivePolicyPolicy(context.Background(), organizationId).CreateOrganizationAdaptivePolicyPolicy(createOrganizationAdaptivePolicyPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.CreateOrganizationAdaptivePolicyPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationAdaptivePolicyPolicy`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.CreateOrganizationAdaptivePolicyPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationAdaptivePolicyPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationAdaptivePolicyPolicy** | [**InlineObject178**](InlineObject178.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationAdmin

> InlineResponse200102 CreateOrganizationAdmin(ctx, organizationId).CreateOrganizationAdmin(createOrganizationAdmin).Execute()

Create a new dashboard administrator



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    createOrganizationAdmin := *openapiclient.NewInlineObject181("Email_example", "Name_example", "OrgAccess_example") // InlineObject181 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.CreateOrganizationAdmin(context.Background(), organizationId).CreateOrganizationAdmin(createOrganizationAdmin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.CreateOrganizationAdmin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationAdmin`: InlineResponse200102
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.CreateOrganizationAdmin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationAdminRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationAdmin** | [**InlineObject181**](InlineObject181.md) |  | 

### Return type

[**InlineResponse200102**](InlineResponse200102.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationAlertsProfile

> map[string]interface{} CreateOrganizationAlertsProfile(ctx, organizationId).CreateOrganizationAlertsProfile(createOrganizationAlertsProfile).Execute()

Create an organization-wide alert configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    createOrganizationAlertsProfile := *openapiclient.NewInlineObject183("Type_example", *openapiclient.NewOrganizationsOrganizationIdAlertsProfilesAlertCondition(), *openapiclient.NewOrganizationsOrganizationIdAlertsProfilesRecipients(), []string{"NetworkTags_example"}) // InlineObject183 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.CreateOrganizationAlertsProfile(context.Background(), organizationId).CreateOrganizationAlertsProfile(createOrganizationAlertsProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.CreateOrganizationAlertsProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationAlertsProfile`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.CreateOrganizationAlertsProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationAlertsProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationAlertsProfile** | [**InlineObject183**](InlineObject183.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationBrandingPolicy

> InlineResponse2015 CreateOrganizationBrandingPolicy(ctx, organizationId).CreateOrganizationBrandingPolicy(createOrganizationBrandingPolicy).Execute()

Add a new branding policy to an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    createOrganizationBrandingPolicy := *openapiclient.NewInlineObject188() // InlineObject188 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.CreateOrganizationBrandingPolicy(context.Background(), organizationId).CreateOrganizationBrandingPolicy(createOrganizationBrandingPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.CreateOrganizationBrandingPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationBrandingPolicy`: InlineResponse2015
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.CreateOrganizationBrandingPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationBrandingPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationBrandingPolicy** | [**InlineObject188**](InlineObject188.md) |  | 

### Return type

[**InlineResponse2015**](InlineResponse2015.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationConfigTemplate

> map[string]interface{} CreateOrganizationConfigTemplate(ctx, organizationId).CreateOrganizationConfigTemplate(createOrganizationConfigTemplate).Execute()

Create a new configuration template



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    createOrganizationConfigTemplate := *openapiclient.NewInlineObject195("Name_example") // InlineObject195 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.CreateOrganizationConfigTemplate(context.Background(), organizationId).CreateOrganizationConfigTemplate(createOrganizationConfigTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.CreateOrganizationConfigTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationConfigTemplate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.CreateOrganizationConfigTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationConfigTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationConfigTemplate** | [**InlineObject195**](InlineObject195.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationEarlyAccessFeaturesOptIn

> map[string]interface{} CreateOrganizationEarlyAccessFeaturesOptIn(ctx, organizationId).CreateOrganizationEarlyAccessFeaturesOptIn(createOrganizationEarlyAccessFeaturesOptIn).Execute()

Create a new early access feature opt-in for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    createOrganizationEarlyAccessFeaturesOptIn := *openapiclient.NewInlineObject198("ShortName_example") // InlineObject198 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.CreateOrganizationEarlyAccessFeaturesOptIn(context.Background(), organizationId).CreateOrganizationEarlyAccessFeaturesOptIn(createOrganizationEarlyAccessFeaturesOptIn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.CreateOrganizationEarlyAccessFeaturesOptIn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationEarlyAccessFeaturesOptIn`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.CreateOrganizationEarlyAccessFeaturesOptIn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationEarlyAccessFeaturesOptInRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationEarlyAccessFeaturesOptIn** | [**InlineObject198**](InlineObject198.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent

> map[string]interface{} CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent(ctx, organizationId).CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent(createOrganizationInventoryOnboardingCloudMonitoringExportEvent).Execute()

Imports event logs related to the onboarding app into elastisearch



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    createOrganizationInventoryOnboardingCloudMonitoringExportEvent := *openapiclient.NewInlineObject203("LogEvent_example", int32(123)) // InlineObject203 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent(context.Background(), organizationId).CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent(createOrganizationInventoryOnboardingCloudMonitoringExportEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationInventoryOnboardingCloudMonitoringExportEvent** | [**InlineObject203**](InlineObject203.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationInventoryOnboardingCloudMonitoringImport

> []InlineResponse2016 CreateOrganizationInventoryOnboardingCloudMonitoringImport(ctx, organizationId).CreateOrganizationInventoryOnboardingCloudMonitoringImport(createOrganizationInventoryOnboardingCloudMonitoringImport).Execute()

Commits the import operation to complete the onboarding of a device into Dashboard for monitoring.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    createOrganizationInventoryOnboardingCloudMonitoringImport := *openapiclient.NewInlineObject204([]openapiclient.OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevices{*openapiclient.NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevices("DeviceId_example", "Udi_example", "NetworkId_example")}) // InlineObject204 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.CreateOrganizationInventoryOnboardingCloudMonitoringImport(context.Background(), organizationId).CreateOrganizationInventoryOnboardingCloudMonitoringImport(createOrganizationInventoryOnboardingCloudMonitoringImport).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.CreateOrganizationInventoryOnboardingCloudMonitoringImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationInventoryOnboardingCloudMonitoringImport`: []InlineResponse2016
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.CreateOrganizationInventoryOnboardingCloudMonitoringImport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationInventoryOnboardingCloudMonitoringImport** | [**InlineObject204**](InlineObject204.md) |  | 

### Return type

[**[]InlineResponse2016**](InlineResponse2016.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationInventoryOnboardingCloudMonitoringPrepare

> []InlineResponse2017 CreateOrganizationInventoryOnboardingCloudMonitoringPrepare(ctx, organizationId).CreateOrganizationInventoryOnboardingCloudMonitoringPrepare(createOrganizationInventoryOnboardingCloudMonitoringPrepare).Execute()

Initiates or updates an import session



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    createOrganizationInventoryOnboardingCloudMonitoringPrepare := *openapiclient.NewInlineObject205([]openapiclient.OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices{*openapiclient.NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices("Sudi_example")}) // InlineObject205 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.CreateOrganizationInventoryOnboardingCloudMonitoringPrepare(context.Background(), organizationId).CreateOrganizationInventoryOnboardingCloudMonitoringPrepare(createOrganizationInventoryOnboardingCloudMonitoringPrepare).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.CreateOrganizationInventoryOnboardingCloudMonitoringPrepare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationInventoryOnboardingCloudMonitoringPrepare`: []InlineResponse2017
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.CreateOrganizationInventoryOnboardingCloudMonitoringPrepare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationInventoryOnboardingCloudMonitoringPrepare** | [**InlineObject205**](InlineObject205.md) |  | 

### Return type

[**[]InlineResponse2017**](InlineResponse2017.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationNetwork

> InlineResponse20012 CreateOrganizationNetwork(ctx, organizationId).CreateOrganizationNetwork(createOrganizationNetwork).Execute()

Create a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    createOrganizationNetwork := *openapiclient.NewInlineObject214("Name_example", []string{"ProductTypes_example"}) // InlineObject214 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.CreateOrganizationNetwork(context.Background(), organizationId).CreateOrganizationNetwork(createOrganizationNetwork).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.CreateOrganizationNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationNetwork`: InlineResponse20012
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.CreateOrganizationNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationNetwork** | [**InlineObject214**](InlineObject214.md) |  | 

### Return type

[**InlineResponse20012**](InlineResponse20012.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationPolicyObject

> map[string]interface{} CreateOrganizationPolicyObject(ctx, organizationId).CreateOrganizationPolicyObject(createOrganizationPolicyObject).Execute()

Creates a new Policy Object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    createOrganizationPolicyObject := *openapiclient.NewInlineObject216("Name_example", "Category_example", "Type_example") // InlineObject216 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.CreateOrganizationPolicyObject(context.Background(), organizationId).CreateOrganizationPolicyObject(createOrganizationPolicyObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.CreateOrganizationPolicyObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationPolicyObject`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.CreateOrganizationPolicyObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationPolicyObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationPolicyObject** | [**InlineObject216**](InlineObject216.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationPolicyObjectsGroup

> map[string]interface{} CreateOrganizationPolicyObjectsGroup(ctx, organizationId).CreateOrganizationPolicyObjectsGroup(createOrganizationPolicyObjectsGroup).Execute()

Creates a new Policy Object Group.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    createOrganizationPolicyObjectsGroup := *openapiclient.NewInlineObject217("Name_example") // InlineObject217 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.CreateOrganizationPolicyObjectsGroup(context.Background(), organizationId).CreateOrganizationPolicyObjectsGroup(createOrganizationPolicyObjectsGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.CreateOrganizationPolicyObjectsGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationPolicyObjectsGroup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.CreateOrganizationPolicyObjectsGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationPolicyObjectsGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationPolicyObjectsGroup** | [**InlineObject217**](InlineObject217.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationSamlIdp

> []InlineResponse200139 CreateOrganizationSamlIdp(ctx, organizationId).CreateOrganizationSamlIdp(createOrganizationSamlIdp).Execute()

Create a SAML IdP for your organization.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    createOrganizationSamlIdp := *openapiclient.NewInlineObject221("X509certSha1Fingerprint_example") // InlineObject221 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.CreateOrganizationSamlIdp(context.Background(), organizationId).CreateOrganizationSamlIdp(createOrganizationSamlIdp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.CreateOrganizationSamlIdp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationSamlIdp`: []InlineResponse200139
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.CreateOrganizationSamlIdp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationSamlIdpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationSamlIdp** | [**InlineObject221**](InlineObject221.md) |  | 

### Return type

[**[]InlineResponse200139**](InlineResponse200139.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationSamlRole

> map[string]interface{} CreateOrganizationSamlRole(ctx, organizationId).CreateOrganizationSamlRole(createOrganizationSamlRole).Execute()

Create a SAML role



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    createOrganizationSamlRole := *openapiclient.NewInlineObject223("Role_example", "OrgAccess_example") // InlineObject223 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.CreateOrganizationSamlRole(context.Background(), organizationId).CreateOrganizationSamlRole(createOrganizationSamlRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.CreateOrganizationSamlRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationSamlRole`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.CreateOrganizationSamlRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationSamlRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationSamlRole** | [**InlineObject223**](InlineObject223.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganization

> DeleteOrganization(ctx, organizationId).Execute()

Delete an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.DeleteOrganization(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.DeleteOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationActionBatch

> DeleteOrganizationActionBatch(ctx, organizationId, actionBatchId).Execute()

Delete an action batch



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    actionBatchId := "actionBatchId_example" // string | Action batch ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.DeleteOrganizationActionBatch(context.Background(), organizationId, actionBatchId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.DeleteOrganizationActionBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**actionBatchId** | **string** | Action batch ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationActionBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationAdaptivePolicyAcl

> DeleteOrganizationAdaptivePolicyAcl(ctx, organizationId, aclId).Execute()

Deletes the specified adaptive policy ACL



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    aclId := "aclId_example" // string | Acl ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.DeleteOrganizationAdaptivePolicyAcl(context.Background(), organizationId, aclId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.DeleteOrganizationAdaptivePolicyAcl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**aclId** | **string** | Acl ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationAdaptivePolicyAclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationAdaptivePolicyGroup

> DeleteOrganizationAdaptivePolicyGroup(ctx, organizationId, id).Execute()

Deletes the specified adaptive policy group and any associated policies and references



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    id := "id_example" // string | ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.DeleteOrganizationAdaptivePolicyGroup(context.Background(), organizationId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.DeleteOrganizationAdaptivePolicyGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationAdaptivePolicyGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationAdaptivePolicyPolicy

> DeleteOrganizationAdaptivePolicyPolicy(ctx, organizationId, id).Execute()

Delete an Adaptive Policy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    id := "id_example" // string | ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.DeleteOrganizationAdaptivePolicyPolicy(context.Background(), organizationId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.DeleteOrganizationAdaptivePolicyPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationAdaptivePolicyPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationAdmin

> DeleteOrganizationAdmin(ctx, organizationId, adminId).Execute()

Revoke all access for a dashboard administrator within this organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    adminId := "adminId_example" // string | Admin ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.DeleteOrganizationAdmin(context.Background(), organizationId, adminId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.DeleteOrganizationAdmin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**adminId** | **string** | Admin ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationAdminRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationAlertsProfile

> DeleteOrganizationAlertsProfile(ctx, organizationId, alertConfigId).Execute()

Removes an organization-wide alert config



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    alertConfigId := "alertConfigId_example" // string | Alert config ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.DeleteOrganizationAlertsProfile(context.Background(), organizationId, alertConfigId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.DeleteOrganizationAlertsProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**alertConfigId** | **string** | Alert config ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationAlertsProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationBrandingPolicy

> DeleteOrganizationBrandingPolicy(ctx, organizationId, brandingPolicyId).Execute()

Delete a branding policy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    brandingPolicyId := "brandingPolicyId_example" // string | Branding policy ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.DeleteOrganizationBrandingPolicy(context.Background(), organizationId, brandingPolicyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.DeleteOrganizationBrandingPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**brandingPolicyId** | **string** | Branding policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationBrandingPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationConfigTemplate

> DeleteOrganizationConfigTemplate(ctx, organizationId, configTemplateId).Execute()

Remove a configuration template



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    configTemplateId := "configTemplateId_example" // string | Config template ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.DeleteOrganizationConfigTemplate(context.Background(), organizationId, configTemplateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.DeleteOrganizationConfigTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**configTemplateId** | **string** | Config template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationConfigTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationEarlyAccessFeaturesOptIn

> DeleteOrganizationEarlyAccessFeaturesOptIn(ctx, organizationId, optInId).Execute()

Delete an early access feature opt-in



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    optInId := "optInId_example" // string | Opt in ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.DeleteOrganizationEarlyAccessFeaturesOptIn(context.Background(), organizationId, optInId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.DeleteOrganizationEarlyAccessFeaturesOptIn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**optInId** | **string** | Opt in ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationEarlyAccessFeaturesOptInRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationPolicyObject

> DeleteOrganizationPolicyObject(ctx, organizationId, policyObjectId).Execute()

Deletes a Policy Object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    policyObjectId := "policyObjectId_example" // string | Policy object ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.DeleteOrganizationPolicyObject(context.Background(), organizationId, policyObjectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.DeleteOrganizationPolicyObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**policyObjectId** | **string** | Policy object ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationPolicyObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationPolicyObjectsGroup

> DeleteOrganizationPolicyObjectsGroup(ctx, organizationId, policyObjectGroupId).Execute()

Deletes a Policy Object Group.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    policyObjectGroupId := "policyObjectGroupId_example" // string | Policy object group ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.DeleteOrganizationPolicyObjectsGroup(context.Background(), organizationId, policyObjectGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.DeleteOrganizationPolicyObjectsGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**policyObjectGroupId** | **string** | Policy object group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationPolicyObjectsGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationSamlIdp

> DeleteOrganizationSamlIdp(ctx, organizationId, idpId).Execute()

Remove a SAML IdP in your organization.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    idpId := "idpId_example" // string | Idp ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.DeleteOrganizationSamlIdp(context.Background(), organizationId, idpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.DeleteOrganizationSamlIdp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**idpId** | **string** | Idp ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationSamlIdpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationSamlRole

> DeleteOrganizationSamlRole(ctx, organizationId, samlRoleId).Execute()

Remove a SAML role



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    samlRoleId := "samlRoleId_example" // string | Saml role ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.DeleteOrganizationSamlRole(context.Background(), organizationId, samlRoleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.DeleteOrganizationSamlRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**samlRoleId** | **string** | Saml role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationSamlRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationUser

> DeleteOrganizationUser(ctx, organizationId, userId).Execute()

Delete a user and all of its authentication methods.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    userId := "userId_example" // string | User ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.DeleteOrganizationUser(context.Background(), organizationId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.DeleteOrganizationUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**userId** | **string** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganization

> InlineResponse20099 GetOrganization(ctx, organizationId).Execute()

Return an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganization(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganization`: InlineResponse20099
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20099**](InlineResponse20099.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationActionBatch

> InlineResponse2014 GetOrganizationActionBatch(ctx, organizationId, actionBatchId).Execute()

Return an action batch



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    actionBatchId := "actionBatchId_example" // string | Action batch ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationActionBatch(context.Background(), organizationId, actionBatchId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationActionBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationActionBatch`: InlineResponse2014
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationActionBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**actionBatchId** | **string** | Action batch ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationActionBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse2014**](InlineResponse2014.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationActionBatches

> []map[string]interface{} GetOrganizationActionBatches(ctx, organizationId).Status(status).Execute()

Return the list of action batches in the organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    status := "status_example" // string | Filter batches by status. Valid types are pending, completed, and failed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationActionBatches(context.Background(), organizationId).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationActionBatches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationActionBatches`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationActionBatches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationActionBatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | Filter batches by status. Valid types are pending, completed, and failed. | 

### Return type

**[]map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAdaptivePolicyAcl

> InlineResponse200100 GetOrganizationAdaptivePolicyAcl(ctx, organizationId, aclId).Execute()

Returns the adaptive policy ACL information



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    aclId := "aclId_example" // string | Acl ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationAdaptivePolicyAcl(context.Background(), organizationId, aclId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationAdaptivePolicyAcl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAdaptivePolicyAcl`: InlineResponse200100
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationAdaptivePolicyAcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**aclId** | **string** | Acl ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAdaptivePolicyAclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200100**](InlineResponse200100.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAdaptivePolicyAcls

> []InlineResponse200100 GetOrganizationAdaptivePolicyAcls(ctx, organizationId).Execute()

List adaptive policy ACLs in a organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationAdaptivePolicyAcls(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationAdaptivePolicyAcls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAdaptivePolicyAcls`: []InlineResponse200100
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationAdaptivePolicyAcls`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAdaptivePolicyAclsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse200100**](InlineResponse200100.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAdaptivePolicyGroup

> map[string]interface{} GetOrganizationAdaptivePolicyGroup(ctx, organizationId, id).Execute()

Returns an adaptive policy group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    id := "id_example" // string | ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationAdaptivePolicyGroup(context.Background(), organizationId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationAdaptivePolicyGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAdaptivePolicyGroup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationAdaptivePolicyGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAdaptivePolicyGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAdaptivePolicyGroups

> []map[string]interface{} GetOrganizationAdaptivePolicyGroups(ctx, organizationId).Execute()

List adaptive policy groups in a organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationAdaptivePolicyGroups(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationAdaptivePolicyGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAdaptivePolicyGroups`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationAdaptivePolicyGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAdaptivePolicyGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAdaptivePolicyOverview

> InlineResponse200101 GetOrganizationAdaptivePolicyOverview(ctx, organizationId).Execute()

Returns adaptive policy aggregate statistics for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationAdaptivePolicyOverview(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationAdaptivePolicyOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAdaptivePolicyOverview`: InlineResponse200101
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationAdaptivePolicyOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAdaptivePolicyOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200101**](InlineResponse200101.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAdaptivePolicyPolicies

> []map[string]interface{} GetOrganizationAdaptivePolicyPolicies(ctx, organizationId).Execute()

List adaptive policies in an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationAdaptivePolicyPolicies(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationAdaptivePolicyPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAdaptivePolicyPolicies`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationAdaptivePolicyPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAdaptivePolicyPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAdaptivePolicyPolicy

> map[string]interface{} GetOrganizationAdaptivePolicyPolicy(ctx, organizationId, id).Execute()

Return an adaptive policy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    id := "id_example" // string | ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationAdaptivePolicyPolicy(context.Background(), organizationId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationAdaptivePolicyPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAdaptivePolicyPolicy`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationAdaptivePolicyPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAdaptivePolicyPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAdaptivePolicySettings

> map[string]interface{} GetOrganizationAdaptivePolicySettings(ctx, organizationId).Execute()

Returns global adaptive policy settings in an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationAdaptivePolicySettings(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationAdaptivePolicySettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAdaptivePolicySettings`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationAdaptivePolicySettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAdaptivePolicySettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAdmins

> []InlineResponse200102 GetOrganizationAdmins(ctx, organizationId).Execute()

List the dashboard administrators in this organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationAdmins(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationAdmins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAdmins`: []InlineResponse200102
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationAdmins`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAdminsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse200102**](InlineResponse200102.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAlertsProfiles

> []map[string]interface{} GetOrganizationAlertsProfiles(ctx, organizationId).Execute()

List all organization-wide alert configurations



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationAlertsProfiles(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationAlertsProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAlertsProfiles`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationAlertsProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAlertsProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApiRequests

> []InlineResponse200103 GetOrganizationApiRequests(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).AdminId(adminId).Path(path).Method(method).ResponseCode(responseCode).SourceIp(sourceIp).UserAgent(userAgent).Version(version).OperationIds(operationIds).Execute()

List the API requests made by an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 31 days. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    adminId := "adminId_example" // string | Filter the results by the ID of the admin who made the API requests (optional)
    path := "path_example" // string | Filter the results by the path of the API requests (optional)
    method := "method_example" // string | Filter the results by the method of the API requests (must be 'GET', 'PUT', 'POST' or 'DELETE') (optional)
    responseCode := int32(56) // int32 | Filter the results by the response code of the API requests (optional)
    sourceIp := "sourceIp_example" // string | Filter the results by the IP address of the originating API request (optional)
    userAgent := "userAgent_example" // string | Filter the results by the user agent string of the API request (optional)
    version := int32(56) // int32 | Filter the results by the API version of the API request (optional)
    operationIds := []string{"Inner_example"} // []string | Filter the results by one or more operation IDs for the API request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationApiRequests(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).AdminId(adminId).Path(path).Method(method).ResponseCode(responseCode).SourceIp(sourceIp).UserAgent(userAgent).Version(version).OperationIds(operationIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationApiRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApiRequests`: []InlineResponse200103
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationApiRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApiRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 31 days. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **adminId** | **string** | Filter the results by the ID of the admin who made the API requests | 
 **path** | **string** | Filter the results by the path of the API requests | 
 **method** | **string** | Filter the results by the method of the API requests (must be &#39;GET&#39;, &#39;PUT&#39;, &#39;POST&#39; or &#39;DELETE&#39;) | 
 **responseCode** | **int32** | Filter the results by the response code of the API requests | 
 **sourceIp** | **string** | Filter the results by the IP address of the originating API request | 
 **userAgent** | **string** | Filter the results by the user agent string of the API request | 
 **version** | **int32** | Filter the results by the API version of the API request | 
 **operationIds** | **[]string** | Filter the results by one or more operation IDs for the API request | 

### Return type

[**[]InlineResponse200103**](InlineResponse200103.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApiRequestsOverview

> InlineResponse200104 GetOrganizationApiRequestsOverview(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()

Return an aggregated overview of API requests data



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 31 days. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationApiRequestsOverview(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationApiRequestsOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApiRequestsOverview`: InlineResponse200104
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationApiRequestsOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApiRequestsOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 31 days. | 

### Return type

[**InlineResponse200104**](InlineResponse200104.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApiRequestsOverviewResponseCodesByInterval

> []InlineResponse200105 GetOrganizationApiRequestsOverviewResponseCodesByInterval(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Version(version).OperationIds(operationIds).SourceIps(sourceIps).AdminIds(adminIds).UserAgent(userAgent).Execute()

Tracks organizations' API requests by response code across a given time period



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 31 days. If interval is provided, the timespan will be autocalculated. (optional)
    interval := int32(56) // int32 | The time interval in seconds for returned data. The valid intervals are: 120, 3600, 14400, 21600. The default is 21600. Interval is calculated if time params are provided. (optional)
    version := int32(56) // int32 | Filter by API version of the endpoint. Allowable values are: [0, 1] (optional)
    operationIds := []string{"Inner_example"} // []string | Filter by operation ID of the endpoint (optional)
    sourceIps := []string{"Inner_example"} // []string | Filter by source IP that made the API request (optional)
    adminIds := []string{"Inner_example"} // []string | Filter by admin ID of user that made the API request (optional)
    userAgent := "userAgent_example" // string | Filter by user agent string for API request. This will filter by a complete or partial match. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationApiRequestsOverviewResponseCodesByInterval(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Version(version).OperationIds(operationIds).SourceIps(sourceIps).AdminIds(adminIds).UserAgent(userAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationApiRequestsOverviewResponseCodesByInterval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApiRequestsOverviewResponseCodesByInterval`: []InlineResponse200105
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationApiRequestsOverviewResponseCodesByInterval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 31 days. If interval is provided, the timespan will be autocalculated. | 
 **interval** | **int32** | The time interval in seconds for returned data. The valid intervals are: 120, 3600, 14400, 21600. The default is 21600. Interval is calculated if time params are provided. | 
 **version** | **int32** | Filter by API version of the endpoint. Allowable values are: [0, 1] | 
 **operationIds** | **[]string** | Filter by operation ID of the endpoint | 
 **sourceIps** | **[]string** | Filter by source IP that made the API request | 
 **adminIds** | **[]string** | Filter by admin ID of user that made the API request | 
 **userAgent** | **string** | Filter by user agent string for API request. This will filter by a complete or partial match. | 

### Return type

[**[]InlineResponse200105**](InlineResponse200105.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationBrandingPolicies

> []InlineResponse200108 GetOrganizationBrandingPolicies(ctx, organizationId).Execute()

List the branding policies of an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationBrandingPolicies(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationBrandingPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationBrandingPolicies`: []InlineResponse200108
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationBrandingPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationBrandingPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse200108**](InlineResponse200108.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationBrandingPoliciesPriorities

> InlineResponse200109 GetOrganizationBrandingPoliciesPriorities(ctx, organizationId).Execute()

Return the branding policy IDs of an organization in priority order



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationBrandingPoliciesPriorities(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationBrandingPoliciesPriorities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationBrandingPoliciesPriorities`: InlineResponse200109
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationBrandingPoliciesPriorities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationBrandingPoliciesPrioritiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200109**](InlineResponse200109.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationBrandingPolicy

> InlineResponse200108 GetOrganizationBrandingPolicy(ctx, organizationId, brandingPolicyId).Execute()

Return a branding policy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    brandingPolicyId := "brandingPolicyId_example" // string | Branding policy ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationBrandingPolicy(context.Background(), organizationId, brandingPolicyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationBrandingPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationBrandingPolicy`: InlineResponse200108
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationBrandingPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**brandingPolicyId** | **string** | Branding policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationBrandingPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200108**](InlineResponse200108.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationClientsBandwidthUsageHistory

> []InlineResponse200111 GetOrganizationClientsBandwidthUsageHistory(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()

Return data usage (in megabits per second) over time for all clients in the given organization within a given time range.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    t0 := "t0_example" // string | The beginning of the timespan for the data. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationClientsBandwidthUsageHistory(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationClientsBandwidthUsageHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationClientsBandwidthUsageHistory`: []InlineResponse200111
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationClientsBandwidthUsageHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationClientsBandwidthUsageHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 

### Return type

[**[]InlineResponse200111**](InlineResponse200111.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationClientsOverview

> InlineResponse200112 GetOrganizationClientsOverview(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()

Return summary information around client data usage (in mb) across the given organization.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    t0 := "t0_example" // string | The beginning of the timespan for the data. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationClientsOverview(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationClientsOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationClientsOverview`: InlineResponse200112
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationClientsOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationClientsOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 

### Return type

[**InlineResponse200112**](InlineResponse200112.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationClientsSearch

> map[string]interface{} GetOrganizationClientsSearch(ctx, organizationId).Mac(mac).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Return the client details in an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    mac := "mac_example" // string | The MAC address of the client. Required.
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 5. Default is 5. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationClientsSearch(context.Background(), organizationId).Mac(mac).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationClientsSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationClientsSearch`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationClientsSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationClientsSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mac** | **string** | The MAC address of the client. Required. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 5. Default is 5. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationConfigTemplate

> map[string]interface{} GetOrganizationConfigTemplate(ctx, organizationId, configTemplateId).Execute()

Return a single configuration template



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    configTemplateId := "configTemplateId_example" // string | Config template ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationConfigTemplate(context.Background(), organizationId, configTemplateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationConfigTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationConfigTemplate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationConfigTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**configTemplateId** | **string** | Config template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationConfigTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationConfigTemplates

> []map[string]interface{} GetOrganizationConfigTemplates(ctx, organizationId).Execute()

List the configuration templates for this organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationConfigTemplates(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationConfigTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationConfigTemplates`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationConfigTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationConfigTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationConfigurationChanges

> []map[string]interface{} GetOrganizationConfigurationChanges(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkId(networkId).AdminId(adminId).Execute()

View the Change Log for your organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 365 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 365 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 365 days. The default is 365 days. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 5000. Default is 5000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkId := "networkId_example" // string | Filters on the given network (optional)
    adminId := "adminId_example" // string | Filters on the given Admin (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationConfigurationChanges(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkId(networkId).AdminId(adminId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationConfigurationChanges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationConfigurationChanges`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationConfigurationChanges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationConfigurationChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 365 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 365 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 365 days. The default is 365 days. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 5000. Default is 5000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkId** | **string** | Filters on the given network | 
 **adminId** | **string** | Filters on the given Admin | 

### Return type

**[]map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevices

> []InlineResponse200115 GetOrganizationDevices(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).ConfigurationUpdatedAfter(configurationUpdatedAfter).NetworkIds(networkIds).ProductTypes(productTypes).Tags(tags).TagsFilterType(tagsFilterType).Name(name).Mac(mac).Serial(serial).Model(model).Macs(macs).Serials(serials).SensorMetrics(sensorMetrics).SensorAlertProfileIds(sensorAlertProfileIds).Models(models).Execute()

List the devices in an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    configurationUpdatedAfter := "configurationUpdatedAfter_example" // string | Filter results by whether or not the device's configuration has been updated after the given timestamp (optional)
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter devices by network. (optional)
    productTypes := []string{"ProductTypes_example"} // []string | Optional parameter to filter devices by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, and cloudGateway. (optional)
    tags := []string{"Inner_example"} // []string | Optional parameter to filter devices by tags. (optional)
    tagsFilterType := "tagsFilterType_example" // string | Optional parameter of value 'withAnyTags' or 'withAllTags' to indicate whether to return networks which contain ANY or ALL of the included tags. If no type is included, 'withAnyTags' will be selected. (optional)
    name := "name_example" // string | Optional parameter to filter devices by name. All returned devices will have a name that contains the search term or is an exact match. (optional)
    mac := "mac_example" // string | Optional parameter to filter devices by MAC address. All returned devices will have a MAC address that contains the search term or is an exact match. (optional)
    serial := "serial_example" // string | Optional parameter to filter devices by serial number. All returned devices will have a serial number that contains the search term or is an exact match. (optional)
    model := "model_example" // string | Optional parameter to filter devices by model. All returned devices will have a model that contains the search term or is an exact match. (optional)
    macs := []string{"Inner_example"} // []string | Optional parameter to filter devices by one or more MAC addresses. All returned devices will have a MAC address that is an exact match. (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter devices by one or more serial numbers. All returned devices will have a serial number that is an exact match. (optional)
    sensorMetrics := []string{"Inner_example"} // []string | Optional parameter to filter devices by the metrics that they provide. Only applies to sensor devices. (optional)
    sensorAlertProfileIds := []string{"Inner_example"} // []string | Optional parameter to filter devices by the alert profiles that are bound to them. Only applies to sensor devices. (optional)
    models := []string{"Inner_example"} // []string | Optional parameter to filter devices by one or more models. All returned devices will have a model that is an exact match. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationDevices(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).ConfigurationUpdatedAfter(configurationUpdatedAfter).NetworkIds(networkIds).ProductTypes(productTypes).Tags(tags).TagsFilterType(tagsFilterType).Name(name).Mac(mac).Serial(serial).Model(model).Macs(macs).Serials(serials).SensorMetrics(sensorMetrics).SensorAlertProfileIds(sensorAlertProfileIds).Models(models).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevices`: []InlineResponse200115
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **configurationUpdatedAfter** | **string** | Filter results by whether or not the device&#39;s configuration has been updated after the given timestamp | 
 **networkIds** | **[]string** | Optional parameter to filter devices by network. | 
 **productTypes** | **[]string** | Optional parameter to filter devices by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, and cloudGateway. | 
 **tags** | **[]string** | Optional parameter to filter devices by tags. | 
 **tagsFilterType** | **string** | Optional parameter of value &#39;withAnyTags&#39; or &#39;withAllTags&#39; to indicate whether to return networks which contain ANY or ALL of the included tags. If no type is included, &#39;withAnyTags&#39; will be selected. | 
 **name** | **string** | Optional parameter to filter devices by name. All returned devices will have a name that contains the search term or is an exact match. | 
 **mac** | **string** | Optional parameter to filter devices by MAC address. All returned devices will have a MAC address that contains the search term or is an exact match. | 
 **serial** | **string** | Optional parameter to filter devices by serial number. All returned devices will have a serial number that contains the search term or is an exact match. | 
 **model** | **string** | Optional parameter to filter devices by model. All returned devices will have a model that contains the search term or is an exact match. | 
 **macs** | **[]string** | Optional parameter to filter devices by one or more MAC addresses. All returned devices will have a MAC address that is an exact match. | 
 **serials** | **[]string** | Optional parameter to filter devices by one or more serial numbers. All returned devices will have a serial number that is an exact match. | 
 **sensorMetrics** | **[]string** | Optional parameter to filter devices by the metrics that they provide. Only applies to sensor devices. | 
 **sensorAlertProfileIds** | **[]string** | Optional parameter to filter devices by the alert profiles that are bound to them. Only applies to sensor devices. | 
 **models** | **[]string** | Optional parameter to filter devices by one or more models. All returned devices will have a model that is an exact match. | 

### Return type

[**[]InlineResponse200115**](InlineResponse200115.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesAvailabilities

> []InlineResponse200116 GetOrganizationDevicesAvailabilities(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).ProductTypes(productTypes).Serials(serials).Tags(tags).TagsFilterType(tagsFilterType).Execute()

List the availability information for devices in an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter device availabilities by network ID. This filter uses multiple exact matches. (optional)
    productTypes := []string{"Inner_example"} // []string | Optional parameter to filter device availabilities by device product types. This filter uses multiple exact matches. (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter device availabilities by device serial numbers. This filter uses multiple exact matches. (optional)
    tags := []string{"Inner_example"} // []string | An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, 'tagsFilterType' should also be included (see below). This filter uses multiple exact matches. (optional)
    tagsFilterType := "tagsFilterType_example" // string | An optional parameter of value 'withAnyTags' or 'withAllTags' to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, 'withAnyTags' will be selected. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationDevicesAvailabilities(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).ProductTypes(productTypes).Serials(serials).Tags(tags).TagsFilterType(tagsFilterType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationDevicesAvailabilities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesAvailabilities`: []InlineResponse200116
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationDevicesAvailabilities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesAvailabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Optional parameter to filter device availabilities by network ID. This filter uses multiple exact matches. | 
 **productTypes** | **[]string** | Optional parameter to filter device availabilities by device product types. This filter uses multiple exact matches. | 
 **serials** | **[]string** | Optional parameter to filter device availabilities by device serial numbers. This filter uses multiple exact matches. | 
 **tags** | **[]string** | An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, &#39;tagsFilterType&#39; should also be included (see below). This filter uses multiple exact matches. | 
 **tagsFilterType** | **string** | An optional parameter of value &#39;withAnyTags&#39; or &#39;withAllTags&#39; to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, &#39;withAnyTags&#39; will be selected. | 

### Return type

[**[]InlineResponse200116**](InlineResponse200116.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesPowerModulesStatusesByDevice

> []InlineResponse200117 GetOrganizationDevicesPowerModulesStatusesByDevice(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).ProductTypes(productTypes).Serials(serials).Tags(tags).TagsFilterType(tagsFilterType).Execute()

List the power status information for devices in an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter device availabilities by network ID. This filter uses multiple exact matches. (optional)
    productTypes := []string{"Inner_example"} // []string | Optional parameter to filter device availabilities by device product types. This filter uses multiple exact matches. (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter device availabilities by device serial numbers. This filter uses multiple exact matches. (optional)
    tags := []string{"Inner_example"} // []string | An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, 'tagsFilterType' should also be included (see below). This filter uses multiple exact matches. (optional)
    tagsFilterType := "tagsFilterType_example" // string | An optional parameter of value 'withAnyTags' or 'withAllTags' to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, 'withAnyTags' will be selected. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationDevicesPowerModulesStatusesByDevice(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).ProductTypes(productTypes).Serials(serials).Tags(tags).TagsFilterType(tagsFilterType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationDevicesPowerModulesStatusesByDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesPowerModulesStatusesByDevice`: []InlineResponse200117
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationDevicesPowerModulesStatusesByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Optional parameter to filter device availabilities by network ID. This filter uses multiple exact matches. | 
 **productTypes** | **[]string** | Optional parameter to filter device availabilities by device product types. This filter uses multiple exact matches. | 
 **serials** | **[]string** | Optional parameter to filter device availabilities by device serial numbers. This filter uses multiple exact matches. | 
 **tags** | **[]string** | An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, &#39;tagsFilterType&#39; should also be included (see below). This filter uses multiple exact matches. | 
 **tagsFilterType** | **string** | An optional parameter of value &#39;withAnyTags&#39; or &#39;withAllTags&#39; to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, &#39;withAnyTags&#39; will be selected. | 

### Return type

[**[]InlineResponse200117**](InlineResponse200117.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesProvisioningStatuses

> []InlineResponse200118 GetOrganizationDevicesProvisioningStatuses(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).ProductTypes(productTypes).Serials(serials).Status(status).Tags(tags).TagsFilterType(tagsFilterType).Execute()

List the provisioning statuses information for devices in an organization.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter device by network ID. This filter uses multiple exact matches. (optional)
    productTypes := []string{"Inner_example"} // []string | Optional parameter to filter device by device product types. This filter uses multiple exact matches. (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter device by device serial numbers. This filter uses multiple exact matches. (optional)
    status := "status_example" // string | An optional parameter to filter devices by the provisioning status. Accepted statuses: unprovisioned, incomplete, complete. (optional)
    tags := []string{"Inner_example"} // []string | An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, 'tagsFilterType' should also be included (see below). This filter uses multiple exact matches. (optional)
    tagsFilterType := "tagsFilterType_example" // string | An optional parameter of value 'withAnyTags' or 'withAllTags' to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, 'withAnyTags' will be selected. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationDevicesProvisioningStatuses(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).ProductTypes(productTypes).Serials(serials).Status(status).Tags(tags).TagsFilterType(tagsFilterType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationDevicesProvisioningStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesProvisioningStatuses`: []InlineResponse200118
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationDevicesProvisioningStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesProvisioningStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Optional parameter to filter device by network ID. This filter uses multiple exact matches. | 
 **productTypes** | **[]string** | Optional parameter to filter device by device product types. This filter uses multiple exact matches. | 
 **serials** | **[]string** | Optional parameter to filter device by device serial numbers. This filter uses multiple exact matches. | 
 **status** | **string** | An optional parameter to filter devices by the provisioning status. Accepted statuses: unprovisioned, incomplete, complete. | 
 **tags** | **[]string** | An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, &#39;tagsFilterType&#39; should also be included (see below). This filter uses multiple exact matches. | 
 **tagsFilterType** | **string** | An optional parameter of value &#39;withAnyTags&#39; or &#39;withAllTags&#39; to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, &#39;withAnyTags&#39; will be selected. | 

### Return type

[**[]InlineResponse200118**](InlineResponse200118.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesStatuses

> InlineResponse200119 GetOrganizationDevicesStatuses(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Serials(serials).Statuses(statuses).ProductTypes(productTypes).Models(models).Tags(tags).TagsFilterType(tagsFilterType).Execute()

List the status of every Meraki device in the organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter devices by network ids. (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter devices by serials. (optional)
    statuses := []string{"Statuses_example"} // []string | Optional parameter to filter devices by statuses. Valid statuses are [\"online\", \"alerting\", \"offline\", \"dormant\"]. (optional)
    productTypes := []string{"ProductTypes_example"} // []string | An optional parameter to filter device statuses by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, and cloudGateway. (optional)
    models := []string{"Inner_example"} // []string | Optional parameter to filter devices by models. (optional)
    tags := []string{"Inner_example"} // []string | An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, 'tagsFilterType' should also be included (see below). (optional)
    tagsFilterType := "tagsFilterType_example" // string | An optional parameter of value 'withAnyTags' or 'withAllTags' to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, 'withAnyTags' will be selected. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationDevicesStatuses(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Serials(serials).Statuses(statuses).ProductTypes(productTypes).Models(models).Tags(tags).TagsFilterType(tagsFilterType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationDevicesStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesStatuses`: InlineResponse200119
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationDevicesStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Optional parameter to filter devices by network ids. | 
 **serials** | **[]string** | Optional parameter to filter devices by serials. | 
 **statuses** | **[]string** | Optional parameter to filter devices by statuses. Valid statuses are [\&quot;online\&quot;, \&quot;alerting\&quot;, \&quot;offline\&quot;, \&quot;dormant\&quot;]. | 
 **productTypes** | **[]string** | An optional parameter to filter device statuses by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, and cloudGateway. | 
 **models** | **[]string** | Optional parameter to filter devices by models. | 
 **tags** | **[]string** | An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, &#39;tagsFilterType&#39; should also be included (see below). | 
 **tagsFilterType** | **string** | An optional parameter of value &#39;withAnyTags&#39; or &#39;withAllTags&#39; to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, &#39;withAnyTags&#39; will be selected. | 

### Return type

[**InlineResponse200119**](InlineResponse200119.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesStatusesOverview

> InlineResponse200120 GetOrganizationDevicesStatusesOverview(ctx, organizationId).ProductTypes(productTypes).NetworkIds(networkIds).Execute()

Return an overview of current device statuses



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    productTypes := []string{"ProductTypes_example"} // []string | An optional parameter to filter device statuses by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, and cloudGateway. (optional)
    networkIds := []string{"Inner_example"} // []string | An optional parameter to filter device statuses by network. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationDevicesStatusesOverview(context.Background(), organizationId).ProductTypes(productTypes).NetworkIds(networkIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationDevicesStatusesOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesStatusesOverview`: InlineResponse200120
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationDevicesStatusesOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesStatusesOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **productTypes** | **[]string** | An optional parameter to filter device statuses by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, and cloudGateway. | 
 **networkIds** | **[]string** | An optional parameter to filter device statuses by network. | 

### Return type

[**InlineResponse200120**](InlineResponse200120.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesUplinksAddressesByDevice

> []InlineResponse200121 GetOrganizationDevicesUplinksAddressesByDevice(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).ProductTypes(productTypes).Serials(serials).Tags(tags).TagsFilterType(tagsFilterType).Execute()

List the current uplink addresses for devices in an organization.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter device uplinks by network ID. This filter uses multiple exact matches. (optional)
    productTypes := []string{"Inner_example"} // []string | Optional parameter to filter device uplinks by device product types. This filter uses multiple exact matches. (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter device availabilities by device serial numbers. This filter uses multiple exact matches. (optional)
    tags := []string{"Inner_example"} // []string | An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, 'tagsFilterType' should also be included (see below). This filter uses multiple exact matches. (optional)
    tagsFilterType := "tagsFilterType_example" // string | An optional parameter of value 'withAnyTags' or 'withAllTags' to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, 'withAnyTags' will be selected. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationDevicesUplinksAddressesByDevice(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).ProductTypes(productTypes).Serials(serials).Tags(tags).TagsFilterType(tagsFilterType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationDevicesUplinksAddressesByDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesUplinksAddressesByDevice`: []InlineResponse200121
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationDevicesUplinksAddressesByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesUplinksAddressesByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Optional parameter to filter device uplinks by network ID. This filter uses multiple exact matches. | 
 **productTypes** | **[]string** | Optional parameter to filter device uplinks by device product types. This filter uses multiple exact matches. | 
 **serials** | **[]string** | Optional parameter to filter device availabilities by device serial numbers. This filter uses multiple exact matches. | 
 **tags** | **[]string** | An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, &#39;tagsFilterType&#39; should also be included (see below). This filter uses multiple exact matches. | 
 **tagsFilterType** | **string** | An optional parameter of value &#39;withAnyTags&#39; or &#39;withAllTags&#39; to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, &#39;withAnyTags&#39; will be selected. | 

### Return type

[**[]InlineResponse200121**](InlineResponse200121.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesUplinksLossAndLatency

> []InlineResponse200122 GetOrganizationDevicesUplinksLossAndLatency(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).Uplink(uplink).Ip(ip).Execute()

Return the uplink loss and latency for every MX in the organization from at latest 2 minutes ago



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 60 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 5 minutes after t0. The latest possible time that t1 can be is 2 minutes into the past. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 5 minutes. The default is 5 minutes. (optional)
    uplink := "uplink_example" // string | Optional filter for a specific WAN uplink. Valid uplinks are wan1, wan2, wan3, cellular. Default will return all uplinks. (optional)
    ip := "ip_example" // string | Optional filter for a specific destination IP. Default will return all destination IPs. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationDevicesUplinksLossAndLatency(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).Uplink(uplink).Ip(ip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationDevicesUplinksLossAndLatency``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesUplinksLossAndLatency`: []InlineResponse200122
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationDevicesUplinksLossAndLatency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesUplinksLossAndLatencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 60 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 5 minutes after t0. The latest possible time that t1 can be is 2 minutes into the past. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 5 minutes. The default is 5 minutes. | 
 **uplink** | **string** | Optional filter for a specific WAN uplink. Valid uplinks are wan1, wan2, wan3, cellular. Default will return all uplinks. | 
 **ip** | **string** | Optional filter for a specific destination IP. Default will return all destination IPs. | 

### Return type

[**[]InlineResponse200122**](InlineResponse200122.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationEarlyAccessFeatures

> []InlineResponse200123 GetOrganizationEarlyAccessFeatures(ctx, organizationId).Execute()

List the available early access features for organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationEarlyAccessFeatures(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationEarlyAccessFeatures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationEarlyAccessFeatures`: []InlineResponse200123
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationEarlyAccessFeatures`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationEarlyAccessFeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse200123**](InlineResponse200123.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationEarlyAccessFeaturesOptIn

> map[string]interface{} GetOrganizationEarlyAccessFeaturesOptIn(ctx, organizationId, optInId).Execute()

Show an early access feature opt-in for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    optInId := "optInId_example" // string | Opt in ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationEarlyAccessFeaturesOptIn(context.Background(), organizationId, optInId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationEarlyAccessFeaturesOptIn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationEarlyAccessFeaturesOptIn`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationEarlyAccessFeaturesOptIn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**optInId** | **string** | Opt in ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationEarlyAccessFeaturesOptInRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationEarlyAccessFeaturesOptIns

> []map[string]interface{} GetOrganizationEarlyAccessFeaturesOptIns(ctx, organizationId).Execute()

List the early access feature opt-ins for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationEarlyAccessFeaturesOptIns(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationEarlyAccessFeaturesOptIns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationEarlyAccessFeaturesOptIns`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationEarlyAccessFeaturesOptIns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationEarlyAccessFeaturesOptInsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationFirmwareUpgrades

> []InlineResponse200124 GetOrganizationFirmwareUpgrades(ctx, organizationId).Status(status).ProductType(productType).Execute()

Get firmware upgrade information for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    status := []string{"Inner_example"} // []string | The status of an upgrade  (optional)
    productType := []string{"Inner_example"} // []string | The product type in a given upgrade ID (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationFirmwareUpgrades(context.Background(), organizationId).Status(status).ProductType(productType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationFirmwareUpgrades``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationFirmwareUpgrades`: []InlineResponse200124
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationFirmwareUpgrades`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationFirmwareUpgradesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **[]string** | The status of an upgrade  | 
 **productType** | **[]string** | The product type in a given upgrade ID | 

### Return type

[**[]InlineResponse200124**](InlineResponse200124.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationFirmwareUpgradesByDevice

> []InlineResponse200125 GetOrganizationFirmwareUpgradesByDevice(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Serials(serials).Macs(macs).FirmwareUpgradeIds(firmwareUpgradeIds).FirmwareUpgradeBatchIds(firmwareUpgradeBatchIds).Execute()

Get firmware upgrade status for the filtered devices



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 50. Default is 50. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter by network (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter by serial number.  All returned devices will have a serial number that is an exact match. (optional)
    macs := []string{"Inner_example"} // []string | Optional parameter to filter by one or more MAC addresses belonging to devices. All devices returned belong to MAC addresses that are an exact match. (optional)
    firmwareUpgradeIds := []string{"Inner_example"} // []string | Optional parameter to filter by firmware upgrade ids. (optional)
    firmwareUpgradeBatchIds := []string{"Inner_example"} // []string | Optional parameter to filter by firmware upgrade batch ids. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationFirmwareUpgradesByDevice(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Serials(serials).Macs(macs).FirmwareUpgradeIds(firmwareUpgradeIds).FirmwareUpgradeBatchIds(firmwareUpgradeBatchIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationFirmwareUpgradesByDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationFirmwareUpgradesByDevice`: []InlineResponse200125
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationFirmwareUpgradesByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationFirmwareUpgradesByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 50. Default is 50. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Optional parameter to filter by network | 
 **serials** | **[]string** | Optional parameter to filter by serial number.  All returned devices will have a serial number that is an exact match. | 
 **macs** | **[]string** | Optional parameter to filter by one or more MAC addresses belonging to devices. All devices returned belong to MAC addresses that are an exact match. | 
 **firmwareUpgradeIds** | **[]string** | Optional parameter to filter by firmware upgrade ids. | 
 **firmwareUpgradeBatchIds** | **[]string** | Optional parameter to filter by firmware upgrade batch ids. | 

### Return type

[**[]InlineResponse200125**](InlineResponse200125.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationInventoryDevice

> InlineResponse200128 GetOrganizationInventoryDevice(ctx, organizationId, serial).Execute()

Return a single device from the inventory of an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    serial := "serial_example" // string | Serial

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationInventoryDevice(context.Background(), organizationId, serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationInventoryDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationInventoryDevice`: InlineResponse200128
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationInventoryDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInventoryDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200128**](InlineResponse200128.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationInventoryDevices

> []InlineResponse200128 GetOrganizationInventoryDevices(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).UsedState(usedState).Search(search).Macs(macs).NetworkIds(networkIds).Serials(serials).Models(models).OrderNumbers(orderNumbers).Tags(tags).TagsFilterType(tagsFilterType).ProductTypes(productTypes).Execute()

Return the device inventory for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    usedState := "usedState_example" // string | Filter results by used or unused inventory. Accepted values are 'used' or 'unused'. (optional)
    search := "search_example" // string | Search for devices in inventory based on serial number, mac address, or model. (optional)
    macs := []string{"Inner_example"} // []string | Search for devices in inventory based on mac addresses. (optional)
    networkIds := []string{"Inner_example"} // []string | Search for devices in inventory based on network ids. (optional)
    serials := []string{"Inner_example"} // []string | Search for devices in inventory based on serials. (optional)
    models := []string{"Inner_example"} // []string | Search for devices in inventory based on model. (optional)
    orderNumbers := []string{"Inner_example"} // []string | Search for devices in inventory based on order numbers. (optional)
    tags := []string{"Inner_example"} // []string | Filter devices by tags. The filtering is case-sensitive. If tags are included, 'tagsFilterType' should also be included (see below). (optional)
    tagsFilterType := "tagsFilterType_example" // string | To use with 'tags' parameter, to filter devices which contain ANY or ALL given tags. Accepted values are 'withAnyTags' or 'withAllTags', default is 'withAnyTags'. (optional)
    productTypes := []string{"ProductTypes_example"} // []string | Filter devices by product type. Accepted values are appliance, camera, cellularGateway, cloudGateway, sensor, switch, systemsManager, and wireless. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationInventoryDevices(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).UsedState(usedState).Search(search).Macs(macs).NetworkIds(networkIds).Serials(serials).Models(models).OrderNumbers(orderNumbers).Tags(tags).TagsFilterType(tagsFilterType).ProductTypes(productTypes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationInventoryDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationInventoryDevices`: []InlineResponse200128
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationInventoryDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInventoryDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **usedState** | **string** | Filter results by used or unused inventory. Accepted values are &#39;used&#39; or &#39;unused&#39;. | 
 **search** | **string** | Search for devices in inventory based on serial number, mac address, or model. | 
 **macs** | **[]string** | Search for devices in inventory based on mac addresses. | 
 **networkIds** | **[]string** | Search for devices in inventory based on network ids. | 
 **serials** | **[]string** | Search for devices in inventory based on serials. | 
 **models** | **[]string** | Search for devices in inventory based on model. | 
 **orderNumbers** | **[]string** | Search for devices in inventory based on order numbers. | 
 **tags** | **[]string** | Filter devices by tags. The filtering is case-sensitive. If tags are included, &#39;tagsFilterType&#39; should also be included (see below). | 
 **tagsFilterType** | **string** | To use with &#39;tags&#39; parameter, to filter devices which contain ANY or ALL given tags. Accepted values are &#39;withAnyTags&#39; or &#39;withAllTags&#39;, default is &#39;withAnyTags&#39;. | 
 **productTypes** | **[]string** | Filter devices by product type. Accepted values are appliance, camera, cellularGateway, cloudGateway, sensor, switch, systemsManager, and wireless. | 

### Return type

[**[]InlineResponse200128**](InlineResponse200128.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationInventoryOnboardingCloudMonitoringImports

> []InlineResponse200129 GetOrganizationInventoryOnboardingCloudMonitoringImports(ctx, organizationId).ImportIds(importIds).Execute()

Check the status of a committed Import operation



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    importIds := []string{"Inner_example"} // []string | import ids from an imports

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationInventoryOnboardingCloudMonitoringImports(context.Background(), organizationId).ImportIds(importIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationInventoryOnboardingCloudMonitoringImports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationInventoryOnboardingCloudMonitoringImports`: []InlineResponse200129
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationInventoryOnboardingCloudMonitoringImports`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **importIds** | **[]string** | import ids from an imports | 

### Return type

[**[]InlineResponse200129**](InlineResponse200129.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationInventoryOnboardingCloudMonitoringNetworks

> []InlineResponse20012 GetOrganizationInventoryOnboardingCloudMonitoringNetworks(ctx, organizationId).DeviceType(deviceType).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Returns list of networks eligible for adding cloud monitored device



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    deviceType := "deviceType_example" // string | Device Type switch or wireless controller
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 100000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationInventoryOnboardingCloudMonitoringNetworks(context.Background(), organizationId).DeviceType(deviceType).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationInventoryOnboardingCloudMonitoringNetworks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationInventoryOnboardingCloudMonitoringNetworks`: []InlineResponse20012
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationInventoryOnboardingCloudMonitoringNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInventoryOnboardingCloudMonitoringNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceType** | **string** | Device Type switch or wireless controller | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 100000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]InlineResponse20012**](InlineResponse20012.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationLicense

> InlineResponse200130 GetOrganizationLicense(ctx, organizationId, licenseId).Execute()

Display a license



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    licenseId := "licenseId_example" // string | License ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationLicense(context.Background(), organizationId, licenseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationLicense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationLicense`: InlineResponse200130
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationLicense`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**licenseId** | **string** | License ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200130**](InlineResponse200130.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationLicenses

> []InlineResponse200130 GetOrganizationLicenses(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).DeviceSerial(deviceSerial).NetworkId(networkId).State(state).Execute()

List the licenses for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    deviceSerial := "deviceSerial_example" // string | Filter the licenses to those assigned to a particular device. Returned in the same order that they are queued to the device. (optional)
    networkId := "networkId_example" // string | Filter the licenses to those assigned in a particular network (optional)
    state := "state_example" // string | Filter the licenses to those in a particular state. Can be one of 'active', 'expired', 'expiring', 'recentlyQueued', 'unused' or 'unusedActive' (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationLicenses(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).DeviceSerial(deviceSerial).NetworkId(networkId).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationLicenses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationLicenses`: []InlineResponse200130
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationLicenses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationLicensesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **deviceSerial** | **string** | Filter the licenses to those assigned to a particular device. Returned in the same order that they are queued to the device. | 
 **networkId** | **string** | Filter the licenses to those assigned in a particular network | 
 **state** | **string** | Filter the licenses to those in a particular state. Can be one of &#39;active&#39;, &#39;expired&#39;, &#39;expiring&#39;, &#39;recentlyQueued&#39;, &#39;unused&#39; or &#39;unusedActive&#39; | 

### Return type

[**[]InlineResponse200130**](InlineResponse200130.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationLicensesOverview

> map[string]interface{} GetOrganizationLicensesOverview(ctx, organizationId).Execute()

Return an overview of the license state for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationLicensesOverview(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationLicensesOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationLicensesOverview`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationLicensesOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationLicensesOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationLoginSecurity

> InlineResponse200136 GetOrganizationLoginSecurity(ctx, organizationId).Execute()

Returns the login security settings for an organization.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationLoginSecurity(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationLoginSecurity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationLoginSecurity`: InlineResponse200136
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationLoginSecurity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationLoginSecurityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200136**](InlineResponse200136.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationNetworks

> []InlineResponse20012 GetOrganizationNetworks(ctx, organizationId).ConfigTemplateId(configTemplateId).IsBoundToConfigTemplate(isBoundToConfigTemplate).Tags(tags).TagsFilterType(tagsFilterType).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List the networks that the user has privileges on in an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    configTemplateId := "configTemplateId_example" // string | An optional parameter that is the ID of a config template. Will return all networks bound to that template. (optional)
    isBoundToConfigTemplate := true // bool | An optional parameter to filter config template bound networks. If configTemplateId is set, this cannot be false. (optional)
    tags := []string{"Inner_example"} // []string | An optional parameter to filter networks by tags. The filtering is case-sensitive. If tags are included, 'tagsFilterType' should also be included (see below). (optional)
    tagsFilterType := "tagsFilterType_example" // string | An optional parameter of value 'withAnyTags' or 'withAllTags' to indicate whether to return networks which contain ANY or ALL of the included tags. If no type is included, 'withAnyTags' will be selected. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 100000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationNetworks(context.Background(), organizationId).ConfigTemplateId(configTemplateId).IsBoundToConfigTemplate(isBoundToConfigTemplate).Tags(tags).TagsFilterType(tagsFilterType).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationNetworks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationNetworks`: []InlineResponse20012
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **configTemplateId** | **string** | An optional parameter that is the ID of a config template. Will return all networks bound to that template. | 
 **isBoundToConfigTemplate** | **bool** | An optional parameter to filter config template bound networks. If configTemplateId is set, this cannot be false. | 
 **tags** | **[]string** | An optional parameter to filter networks by tags. The filtering is case-sensitive. If tags are included, &#39;tagsFilterType&#39; should also be included (see below). | 
 **tagsFilterType** | **string** | An optional parameter of value &#39;withAnyTags&#39; or &#39;withAllTags&#39; to indicate whether to return networks which contain ANY or ALL of the included tags. If no type is included, &#39;withAnyTags&#39; will be selected. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 100000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]InlineResponse20012**](InlineResponse20012.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationOpenapiSpec

> map[string]interface{} GetOrganizationOpenapiSpec(ctx, organizationId).Version(version).Execute()

Return the OpenAPI Specification of the organization's API documentation in JSON



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    version := int32(56) // int32 | OpenAPI Specification version to return. Default is 2 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationOpenapiSpec(context.Background(), organizationId).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationOpenapiSpec``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationOpenapiSpec`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationOpenapiSpec`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationOpenapiSpecRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **int32** | OpenAPI Specification version to return. Default is 2 | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationPolicyObject

> map[string]interface{} GetOrganizationPolicyObject(ctx, organizationId, policyObjectId).Execute()

Shows details of a Policy Object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    policyObjectId := "policyObjectId_example" // string | Policy object ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationPolicyObject(context.Background(), organizationId, policyObjectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationPolicyObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationPolicyObject`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationPolicyObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**policyObjectId** | **string** | Policy object ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationPolicyObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationPolicyObjects

> []map[string]interface{} GetOrganizationPolicyObjects(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Lists Policy Objects belonging to the organization.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 10 - 5000. Default is 5000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationPolicyObjects(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationPolicyObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationPolicyObjects`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationPolicyObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationPolicyObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 10 - 5000. Default is 5000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

**[]map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationPolicyObjectsGroup

> map[string]interface{} GetOrganizationPolicyObjectsGroup(ctx, organizationId, policyObjectGroupId).Execute()

Shows details of a Policy Object Group.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    policyObjectGroupId := "policyObjectGroupId_example" // string | Policy object group ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationPolicyObjectsGroup(context.Background(), organizationId, policyObjectGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationPolicyObjectsGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationPolicyObjectsGroup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationPolicyObjectsGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**policyObjectGroupId** | **string** | Policy object group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationPolicyObjectsGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationPolicyObjectsGroups

> []map[string]interface{} GetOrganizationPolicyObjectsGroups(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Lists Policy Object Groups belonging to the organization.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 10 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationPolicyObjectsGroups(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationPolicyObjectsGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationPolicyObjectsGroups`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationPolicyObjectsGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationPolicyObjectsGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 10 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

**[]map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSaml

> InlineResponse200138 GetOrganizationSaml(ctx, organizationId).Execute()

Returns the SAML SSO enabled settings for an organization.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationSaml(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationSaml``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSaml`: InlineResponse200138
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationSaml`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSamlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200138**](InlineResponse200138.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSamlIdp

> InlineResponse200139 GetOrganizationSamlIdp(ctx, organizationId, idpId).Execute()

Get a SAML IdP from your organization.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    idpId := "idpId_example" // string | Idp ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationSamlIdp(context.Background(), organizationId, idpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationSamlIdp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSamlIdp`: InlineResponse200139
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationSamlIdp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**idpId** | **string** | Idp ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSamlIdpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200139**](InlineResponse200139.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSamlIdps

> []InlineResponse200139 GetOrganizationSamlIdps(ctx, organizationId).Execute()

List the SAML IdPs in your organization.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationSamlIdps(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationSamlIdps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSamlIdps`: []InlineResponse200139
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationSamlIdps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSamlIdpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse200139**](InlineResponse200139.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSamlRole

> InlineResponse200140 GetOrganizationSamlRole(ctx, organizationId, samlRoleId).Execute()

Return a SAML role



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    samlRoleId := "samlRoleId_example" // string | Saml role ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationSamlRole(context.Background(), organizationId, samlRoleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationSamlRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSamlRole`: InlineResponse200140
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationSamlRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**samlRoleId** | **string** | Saml role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSamlRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200140**](InlineResponse200140.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSamlRoles

> []InlineResponse200140 GetOrganizationSamlRoles(ctx, organizationId).Execute()

List the SAML roles for this organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationSamlRoles(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationSamlRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSamlRoles`: []InlineResponse200140
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationSamlRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSamlRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse200140**](InlineResponse200140.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSnmp

> map[string]interface{} GetOrganizationSnmp(ctx, organizationId).Execute()

Return the SNMP settings for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationSnmp(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationSnmp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSnmp`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationSnmp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSnmpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSummaryTopAppliancesByUtilization

> []InlineResponse200145 GetOrganizationSummaryTopAppliancesByUtilization(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()

Return the top 10 appliances sorted by utilization over given time range.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    t0 := "t0_example" // string | The beginning of the timespan for the data. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationSummaryTopAppliancesByUtilization(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationSummaryTopAppliancesByUtilization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSummaryTopAppliancesByUtilization`: []InlineResponse200145
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationSummaryTopAppliancesByUtilization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSummaryTopAppliancesByUtilizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 

### Return type

[**[]InlineResponse200145**](InlineResponse200145.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSummaryTopClientsByUsage

> []InlineResponse200146 GetOrganizationSummaryTopClientsByUsage(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()

Return metrics for organization's top 10 clients by data usage (in mb) over given time range.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    t0 := "t0_example" // string | The beginning of the timespan for the data. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationSummaryTopClientsByUsage(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationSummaryTopClientsByUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSummaryTopClientsByUsage`: []InlineResponse200146
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationSummaryTopClientsByUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSummaryTopClientsByUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 

### Return type

[**[]InlineResponse200146**](InlineResponse200146.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSummaryTopClientsManufacturersByUsage

> []InlineResponse200147 GetOrganizationSummaryTopClientsManufacturersByUsage(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()

Return metrics for organization's top clients by data usage (in mb) over given time range, grouped by manufacturer.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    t0 := "t0_example" // string | The beginning of the timespan for the data. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationSummaryTopClientsManufacturersByUsage(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationSummaryTopClientsManufacturersByUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSummaryTopClientsManufacturersByUsage`: []InlineResponse200147
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationSummaryTopClientsManufacturersByUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSummaryTopClientsManufacturersByUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 

### Return type

[**[]InlineResponse200147**](InlineResponse200147.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSummaryTopDevicesByUsage

> []InlineResponse200148 GetOrganizationSummaryTopDevicesByUsage(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()

Return metrics for organization's top 10 devices sorted by data usage over given time range



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    t0 := "t0_example" // string | The beginning of the timespan for the data. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationSummaryTopDevicesByUsage(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationSummaryTopDevicesByUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSummaryTopDevicesByUsage`: []InlineResponse200148
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationSummaryTopDevicesByUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSummaryTopDevicesByUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 

### Return type

[**[]InlineResponse200148**](InlineResponse200148.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSummaryTopDevicesModelsByUsage

> []InlineResponse200149 GetOrganizationSummaryTopDevicesModelsByUsage(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()

Return metrics for organization's top 10 device models sorted by data usage over given time range



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    t0 := "t0_example" // string | The beginning of the timespan for the data. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationSummaryTopDevicesModelsByUsage(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationSummaryTopDevicesModelsByUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSummaryTopDevicesModelsByUsage`: []InlineResponse200149
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationSummaryTopDevicesModelsByUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSummaryTopDevicesModelsByUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 

### Return type

[**[]InlineResponse200149**](InlineResponse200149.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSummaryTopSsidsByUsage

> []InlineResponse200150 GetOrganizationSummaryTopSsidsByUsage(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()

Return metrics for organization's top 10 ssids by data usage over given time range



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    t0 := "t0_example" // string | The beginning of the timespan for the data. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationSummaryTopSsidsByUsage(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationSummaryTopSsidsByUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSummaryTopSsidsByUsage`: []InlineResponse200150
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationSummaryTopSsidsByUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSummaryTopSsidsByUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 

### Return type

[**[]InlineResponse200150**](InlineResponse200150.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSummaryTopSwitchesByEnergyUsage

> []InlineResponse200151 GetOrganizationSummaryTopSwitchesByEnergyUsage(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()

Return metrics for organization's top 10 switches by energy usage over given time range



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    t0 := "t0_example" // string | The beginning of the timespan for the data. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationSummaryTopSwitchesByEnergyUsage(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationSummaryTopSwitchesByEnergyUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSummaryTopSwitchesByEnergyUsage`: []InlineResponse200151
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationSummaryTopSwitchesByEnergyUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSummaryTopSwitchesByEnergyUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 

### Return type

[**[]InlineResponse200151**](InlineResponse200151.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationUplinksStatuses

> []InlineResponse200153 GetOrganizationUplinksStatuses(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Serials(serials).Iccids(iccids).Execute()

List the uplink status of every Meraki MX, MG and Z series devices in the organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | A list of network IDs. The returned devices will be filtered to only include these networks. (optional)
    serials := []string{"Inner_example"} // []string | A list of serial numbers. The returned devices will be filtered to only include these serials. (optional)
    iccids := []string{"Inner_example"} // []string | A list of ICCIDs. The returned devices will be filtered to only include these ICCIDs. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationUplinksStatuses(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Serials(serials).Iccids(iccids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationUplinksStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationUplinksStatuses`: []InlineResponse200153
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationUplinksStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationUplinksStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | A list of network IDs. The returned devices will be filtered to only include these networks. | 
 **serials** | **[]string** | A list of serial numbers. The returned devices will be filtered to only include these serials. | 
 **iccids** | **[]string** | A list of ICCIDs. The returned devices will be filtered to only include these ICCIDs. | 

### Return type

[**[]InlineResponse200153**](InlineResponse200153.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWebhooksAlertTypes

> []map[string]interface{} GetOrganizationWebhooksAlertTypes(ctx, organizationId).ProductType(productType).Execute()

Return a list of alert types to be used with managing webhook alerts



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    productType := "productType_example" // string | Filter sample alerts to a specific product type (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationWebhooksAlertTypes(context.Background(), organizationId).ProductType(productType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationWebhooksAlertTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWebhooksAlertTypes`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationWebhooksAlertTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWebhooksAlertTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **productType** | **string** | Filter sample alerts to a specific product type | 

### Return type

**[]map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWebhooksLogs

> []InlineResponse200154 GetOrganizationWebhooksLogs(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Url(url).Execute()

Return the log of webhook POSTs sent



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 90 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    url := "url_example" // string | The URL the webhook was sent to (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationWebhooksLogs(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Url(url).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationWebhooksLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWebhooksLogs`: []InlineResponse200154
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationWebhooksLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWebhooksLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 90 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **url** | **string** | The URL the webhook was sent to | 

### Return type

[**[]InlineResponse200154**](InlineResponse200154.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizations

> []InlineResponse20099 GetOrganizations(ctx).Execute()

List the organizations that the user has privileges on



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizations`: []InlineResponse20099
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationsRequest struct via the builder pattern


### Return type

[**[]InlineResponse20099**](InlineResponse20099.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveOrganizationLicenses

> InlineResponse200132 MoveOrganizationLicenses(ctx, organizationId).MoveOrganizationLicenses(moveOrganizationLicenses).Execute()

Move licenses to another organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    moveOrganizationLicenses := *openapiclient.NewInlineObject208("DestOrganizationId_example", []string{"LicenseIds_example"}) // InlineObject208 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.MoveOrganizationLicenses(context.Background(), organizationId).MoveOrganizationLicenses(moveOrganizationLicenses).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.MoveOrganizationLicenses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoveOrganizationLicenses`: InlineResponse200132
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.MoveOrganizationLicenses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveOrganizationLicensesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **moveOrganizationLicenses** | [**InlineObject208**](InlineObject208.md) |  | 

### Return type

[**InlineResponse200132**](InlineResponse200132.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveOrganizationLicensesSeats

> InlineResponse200133 MoveOrganizationLicensesSeats(ctx, organizationId).MoveOrganizationLicensesSeats(moveOrganizationLicensesSeats).Execute()

Move SM seats to another organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    moveOrganizationLicensesSeats := *openapiclient.NewInlineObject209("DestOrganizationId_example", "LicenseId_example", int32(123)) // InlineObject209 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.MoveOrganizationLicensesSeats(context.Background(), organizationId).MoveOrganizationLicensesSeats(moveOrganizationLicensesSeats).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.MoveOrganizationLicensesSeats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoveOrganizationLicensesSeats`: InlineResponse200133
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.MoveOrganizationLicensesSeats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveOrganizationLicensesSeatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **moveOrganizationLicensesSeats** | [**InlineObject209**](InlineObject209.md) |  | 

### Return type

[**InlineResponse200133**](InlineResponse200133.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReleaseFromOrganizationInventory

> map[string]interface{} ReleaseFromOrganizationInventory(ctx, organizationId).ReleaseFromOrganizationInventory(releaseFromOrganizationInventory).Execute()

Release a list of claimed devices from an organization.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    releaseFromOrganizationInventory := *openapiclient.NewInlineObject206() // InlineObject206 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.ReleaseFromOrganizationInventory(context.Background(), organizationId).ReleaseFromOrganizationInventory(releaseFromOrganizationInventory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.ReleaseFromOrganizationInventory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReleaseFromOrganizationInventory`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.ReleaseFromOrganizationInventory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiReleaseFromOrganizationInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **releaseFromOrganizationInventory** | [**InlineObject206**](InlineObject206.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenewOrganizationLicensesSeats

> InlineResponse200131 RenewOrganizationLicensesSeats(ctx, organizationId).RenewOrganizationLicensesSeats(renewOrganizationLicensesSeats).Execute()

Renew SM seats of a license



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    renewOrganizationLicensesSeats := *openapiclient.NewInlineObject210("LicenseIdToRenew_example", "UnusedLicenseId_example") // InlineObject210 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.RenewOrganizationLicensesSeats(context.Background(), organizationId).RenewOrganizationLicensesSeats(renewOrganizationLicensesSeats).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.RenewOrganizationLicensesSeats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RenewOrganizationLicensesSeats`: InlineResponse200131
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.RenewOrganizationLicensesSeats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenewOrganizationLicensesSeatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **renewOrganizationLicensesSeats** | [**InlineObject210**](InlineObject210.md) |  | 

### Return type

[**InlineResponse200131**](InlineResponse200131.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganization

> InlineResponse20099 UpdateOrganization(ctx, organizationId).UpdateOrganization(updateOrganization).Execute()

Update an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    updateOrganization := *openapiclient.NewInlineObject171() // InlineObject171 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.UpdateOrganization(context.Background(), organizationId).UpdateOrganization(updateOrganization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganization`: InlineResponse20099
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.UpdateOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganization** | [**InlineObject171**](InlineObject171.md) |  | 

### Return type

[**InlineResponse20099**](InlineResponse20099.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationActionBatch

> map[string]interface{} UpdateOrganizationActionBatch(ctx, organizationId, actionBatchId).UpdateOrganizationActionBatch(updateOrganizationActionBatch).Execute()

Update an action batch



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    actionBatchId := "actionBatchId_example" // string | Action batch ID
    updateOrganizationActionBatch := *openapiclient.NewInlineObject173() // InlineObject173 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.UpdateOrganizationActionBatch(context.Background(), organizationId, actionBatchId).UpdateOrganizationActionBatch(updateOrganizationActionBatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateOrganizationActionBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationActionBatch`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.UpdateOrganizationActionBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**actionBatchId** | **string** | Action batch ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationActionBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationActionBatch** | [**InlineObject173**](InlineObject173.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationAdaptivePolicyAcl

> InlineResponse200100 UpdateOrganizationAdaptivePolicyAcl(ctx, organizationId, aclId).UpdateOrganizationAdaptivePolicyAcl(updateOrganizationAdaptivePolicyAcl).Execute()

Updates an adaptive policy ACL



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    aclId := "aclId_example" // string | Acl ID
    updateOrganizationAdaptivePolicyAcl := *openapiclient.NewInlineObject175() // InlineObject175 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.UpdateOrganizationAdaptivePolicyAcl(context.Background(), organizationId, aclId).UpdateOrganizationAdaptivePolicyAcl(updateOrganizationAdaptivePolicyAcl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateOrganizationAdaptivePolicyAcl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationAdaptivePolicyAcl`: InlineResponse200100
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.UpdateOrganizationAdaptivePolicyAcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**aclId** | **string** | Acl ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationAdaptivePolicyAclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationAdaptivePolicyAcl** | [**InlineObject175**](InlineObject175.md) |  | 

### Return type

[**InlineResponse200100**](InlineResponse200100.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationAdaptivePolicyGroup

> map[string]interface{} UpdateOrganizationAdaptivePolicyGroup(ctx, organizationId, id).UpdateOrganizationAdaptivePolicyGroup(updateOrganizationAdaptivePolicyGroup).Execute()

Updates an adaptive policy group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    id := "id_example" // string | ID
    updateOrganizationAdaptivePolicyGroup := *openapiclient.NewInlineObject177() // InlineObject177 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.UpdateOrganizationAdaptivePolicyGroup(context.Background(), organizationId, id).UpdateOrganizationAdaptivePolicyGroup(updateOrganizationAdaptivePolicyGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateOrganizationAdaptivePolicyGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationAdaptivePolicyGroup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.UpdateOrganizationAdaptivePolicyGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationAdaptivePolicyGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationAdaptivePolicyGroup** | [**InlineObject177**](InlineObject177.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationAdaptivePolicyPolicy

> map[string]interface{} UpdateOrganizationAdaptivePolicyPolicy(ctx, organizationId, id).UpdateOrganizationAdaptivePolicyPolicy(updateOrganizationAdaptivePolicyPolicy).Execute()

Update an Adaptive Policy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    id := "id_example" // string | ID
    updateOrganizationAdaptivePolicyPolicy := *openapiclient.NewInlineObject179() // InlineObject179 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.UpdateOrganizationAdaptivePolicyPolicy(context.Background(), organizationId, id).UpdateOrganizationAdaptivePolicyPolicy(updateOrganizationAdaptivePolicyPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateOrganizationAdaptivePolicyPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationAdaptivePolicyPolicy`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.UpdateOrganizationAdaptivePolicyPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationAdaptivePolicyPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationAdaptivePolicyPolicy** | [**InlineObject179**](InlineObject179.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationAdaptivePolicySettings

> map[string]interface{} UpdateOrganizationAdaptivePolicySettings(ctx, organizationId).UpdateOrganizationAdaptivePolicySettings(updateOrganizationAdaptivePolicySettings).Execute()

Update global adaptive policy settings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    updateOrganizationAdaptivePolicySettings := *openapiclient.NewInlineObject180() // InlineObject180 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.UpdateOrganizationAdaptivePolicySettings(context.Background(), organizationId).UpdateOrganizationAdaptivePolicySettings(updateOrganizationAdaptivePolicySettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateOrganizationAdaptivePolicySettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationAdaptivePolicySettings`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.UpdateOrganizationAdaptivePolicySettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationAdaptivePolicySettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationAdaptivePolicySettings** | [**InlineObject180**](InlineObject180.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationAdmin

> InlineResponse200102 UpdateOrganizationAdmin(ctx, organizationId, adminId).UpdateOrganizationAdmin(updateOrganizationAdmin).Execute()

Update an administrator



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    adminId := "adminId_example" // string | Admin ID
    updateOrganizationAdmin := *openapiclient.NewInlineObject182() // InlineObject182 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.UpdateOrganizationAdmin(context.Background(), organizationId, adminId).UpdateOrganizationAdmin(updateOrganizationAdmin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateOrganizationAdmin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationAdmin`: InlineResponse200102
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.UpdateOrganizationAdmin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**adminId** | **string** | Admin ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationAdminRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationAdmin** | [**InlineObject182**](InlineObject182.md) |  | 

### Return type

[**InlineResponse200102**](InlineResponse200102.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationAlertsProfile

> map[string]interface{} UpdateOrganizationAlertsProfile(ctx, organizationId, alertConfigId).UpdateOrganizationAlertsProfile(updateOrganizationAlertsProfile).Execute()

Update an organization-wide alert config



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    alertConfigId := "alertConfigId_example" // string | Alert config ID
    updateOrganizationAlertsProfile := *openapiclient.NewInlineObject184() // InlineObject184 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.UpdateOrganizationAlertsProfile(context.Background(), organizationId, alertConfigId).UpdateOrganizationAlertsProfile(updateOrganizationAlertsProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateOrganizationAlertsProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationAlertsProfile`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.UpdateOrganizationAlertsProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**alertConfigId** | **string** | Alert config ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationAlertsProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationAlertsProfile** | [**InlineObject184**](InlineObject184.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationBrandingPoliciesPriorities

> InlineResponse200109 UpdateOrganizationBrandingPoliciesPriorities(ctx, organizationId).UpdateOrganizationBrandingPoliciesPriorities(updateOrganizationBrandingPoliciesPriorities).Execute()

Update the priority ordering of an organization's branding policies.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    updateOrganizationBrandingPoliciesPriorities := *openapiclient.NewInlineObject189() // InlineObject189 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.UpdateOrganizationBrandingPoliciesPriorities(context.Background(), organizationId).UpdateOrganizationBrandingPoliciesPriorities(updateOrganizationBrandingPoliciesPriorities).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateOrganizationBrandingPoliciesPriorities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationBrandingPoliciesPriorities`: InlineResponse200109
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.UpdateOrganizationBrandingPoliciesPriorities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationBrandingPoliciesPrioritiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationBrandingPoliciesPriorities** | [**InlineObject189**](InlineObject189.md) |  | 

### Return type

[**InlineResponse200109**](InlineResponse200109.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationBrandingPolicy

> InlineResponse200108 UpdateOrganizationBrandingPolicy(ctx, organizationId, brandingPolicyId).UpdateOrganizationBrandingPolicy(updateOrganizationBrandingPolicy).Execute()

Update a branding policy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    brandingPolicyId := "brandingPolicyId_example" // string | Branding policy ID
    updateOrganizationBrandingPolicy := *openapiclient.NewInlineObject190() // InlineObject190 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.UpdateOrganizationBrandingPolicy(context.Background(), organizationId, brandingPolicyId).UpdateOrganizationBrandingPolicy(updateOrganizationBrandingPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateOrganizationBrandingPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationBrandingPolicy`: InlineResponse200108
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.UpdateOrganizationBrandingPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**brandingPolicyId** | **string** | Branding policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationBrandingPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationBrandingPolicy** | [**InlineObject190**](InlineObject190.md) |  | 

### Return type

[**InlineResponse200108**](InlineResponse200108.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationConfigTemplate

> map[string]interface{} UpdateOrganizationConfigTemplate(ctx, organizationId, configTemplateId).UpdateOrganizationConfigTemplate(updateOrganizationConfigTemplate).Execute()

Update a configuration template



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    configTemplateId := "configTemplateId_example" // string | Config template ID
    updateOrganizationConfigTemplate := *openapiclient.NewInlineObject196() // InlineObject196 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.UpdateOrganizationConfigTemplate(context.Background(), organizationId, configTemplateId).UpdateOrganizationConfigTemplate(updateOrganizationConfigTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateOrganizationConfigTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationConfigTemplate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.UpdateOrganizationConfigTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**configTemplateId** | **string** | Config template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationConfigTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationConfigTemplate** | [**InlineObject196**](InlineObject196.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationEarlyAccessFeaturesOptIn

> map[string]interface{} UpdateOrganizationEarlyAccessFeaturesOptIn(ctx, organizationId, optInId).UpdateOrganizationEarlyAccessFeaturesOptIn(updateOrganizationEarlyAccessFeaturesOptIn).Execute()

Update an early access feature opt-in for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    optInId := "optInId_example" // string | Opt in ID
    updateOrganizationEarlyAccessFeaturesOptIn := *openapiclient.NewInlineObject199() // InlineObject199 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.UpdateOrganizationEarlyAccessFeaturesOptIn(context.Background(), organizationId, optInId).UpdateOrganizationEarlyAccessFeaturesOptIn(updateOrganizationEarlyAccessFeaturesOptIn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateOrganizationEarlyAccessFeaturesOptIn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationEarlyAccessFeaturesOptIn`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.UpdateOrganizationEarlyAccessFeaturesOptIn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**optInId** | **string** | Opt in ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationEarlyAccessFeaturesOptInRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationEarlyAccessFeaturesOptIn** | [**InlineObject199**](InlineObject199.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationLicense

> InlineResponse200130 UpdateOrganizationLicense(ctx, organizationId, licenseId).UpdateOrganizationLicense(updateOrganizationLicense).Execute()

Update a license



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    licenseId := "licenseId_example" // string | License ID
    updateOrganizationLicense := *openapiclient.NewInlineObject211() // InlineObject211 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.UpdateOrganizationLicense(context.Background(), organizationId, licenseId).UpdateOrganizationLicense(updateOrganizationLicense).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateOrganizationLicense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationLicense`: InlineResponse200130
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.UpdateOrganizationLicense`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**licenseId** | **string** | License ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationLicense** | [**InlineObject211**](InlineObject211.md) |  | 

### Return type

[**InlineResponse200130**](InlineResponse200130.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationLoginSecurity

> InlineResponse200136 UpdateOrganizationLoginSecurity(ctx, organizationId).UpdateOrganizationLoginSecurity(updateOrganizationLoginSecurity).Execute()

Update the login security settings for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    updateOrganizationLoginSecurity := *openapiclient.NewInlineObject213() // InlineObject213 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.UpdateOrganizationLoginSecurity(context.Background(), organizationId).UpdateOrganizationLoginSecurity(updateOrganizationLoginSecurity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateOrganizationLoginSecurity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationLoginSecurity`: InlineResponse200136
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.UpdateOrganizationLoginSecurity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationLoginSecurityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationLoginSecurity** | [**InlineObject213**](InlineObject213.md) |  | 

### Return type

[**InlineResponse200136**](InlineResponse200136.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationPolicyObject

> map[string]interface{} UpdateOrganizationPolicyObject(ctx, organizationId, policyObjectId).UpdateOrganizationPolicyObject(updateOrganizationPolicyObject).Execute()

Updates a Policy Object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    policyObjectId := "policyObjectId_example" // string | Policy object ID
    updateOrganizationPolicyObject := *openapiclient.NewInlineObject219() // InlineObject219 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.UpdateOrganizationPolicyObject(context.Background(), organizationId, policyObjectId).UpdateOrganizationPolicyObject(updateOrganizationPolicyObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateOrganizationPolicyObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationPolicyObject`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.UpdateOrganizationPolicyObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**policyObjectId** | **string** | Policy object ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationPolicyObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationPolicyObject** | [**InlineObject219**](InlineObject219.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationPolicyObjectsGroup

> map[string]interface{} UpdateOrganizationPolicyObjectsGroup(ctx, organizationId, policyObjectGroupId).UpdateOrganizationPolicyObjectsGroup(updateOrganizationPolicyObjectsGroup).Execute()

Updates a Policy Object Group.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    policyObjectGroupId := "policyObjectGroupId_example" // string | Policy object group ID
    updateOrganizationPolicyObjectsGroup := *openapiclient.NewInlineObject218() // InlineObject218 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.UpdateOrganizationPolicyObjectsGroup(context.Background(), organizationId, policyObjectGroupId).UpdateOrganizationPolicyObjectsGroup(updateOrganizationPolicyObjectsGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateOrganizationPolicyObjectsGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationPolicyObjectsGroup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.UpdateOrganizationPolicyObjectsGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**policyObjectGroupId** | **string** | Policy object group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationPolicyObjectsGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationPolicyObjectsGroup** | [**InlineObject218**](InlineObject218.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationSaml

> InlineResponse200138 UpdateOrganizationSaml(ctx, organizationId).UpdateOrganizationSaml(updateOrganizationSaml).Execute()

Updates the SAML SSO enabled settings for an organization.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    updateOrganizationSaml := *openapiclient.NewInlineObject220() // InlineObject220 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.UpdateOrganizationSaml(context.Background(), organizationId).UpdateOrganizationSaml(updateOrganizationSaml).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateOrganizationSaml``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationSaml`: InlineResponse200138
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.UpdateOrganizationSaml`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationSamlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationSaml** | [**InlineObject220**](InlineObject220.md) |  | 

### Return type

[**InlineResponse200138**](InlineResponse200138.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationSamlIdp

> []InlineResponse200139 UpdateOrganizationSamlIdp(ctx, organizationId, idpId).UpdateOrganizationSamlIdp(updateOrganizationSamlIdp).Execute()

Update a SAML IdP in your organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    idpId := "idpId_example" // string | Idp ID
    updateOrganizationSamlIdp := *openapiclient.NewInlineObject222() // InlineObject222 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.UpdateOrganizationSamlIdp(context.Background(), organizationId, idpId).UpdateOrganizationSamlIdp(updateOrganizationSamlIdp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateOrganizationSamlIdp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationSamlIdp`: []InlineResponse200139
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.UpdateOrganizationSamlIdp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**idpId** | **string** | Idp ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationSamlIdpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationSamlIdp** | [**InlineObject222**](InlineObject222.md) |  | 

### Return type

[**[]InlineResponse200139**](InlineResponse200139.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationSamlRole

> InlineResponse200140 UpdateOrganizationSamlRole(ctx, organizationId, samlRoleId).UpdateOrganizationSamlRole(updateOrganizationSamlRole).Execute()

Update a SAML role



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    samlRoleId := "samlRoleId_example" // string | Saml role ID
    updateOrganizationSamlRole := *openapiclient.NewInlineObject224() // InlineObject224 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.UpdateOrganizationSamlRole(context.Background(), organizationId, samlRoleId).UpdateOrganizationSamlRole(updateOrganizationSamlRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateOrganizationSamlRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationSamlRole`: InlineResponse200140
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.UpdateOrganizationSamlRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**samlRoleId** | **string** | Saml role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationSamlRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationSamlRole** | [**InlineObject224**](InlineObject224.md) |  | 

### Return type

[**InlineResponse200140**](InlineResponse200140.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationSnmp

> map[string]interface{} UpdateOrganizationSnmp(ctx, organizationId).UpdateOrganizationSnmp(updateOrganizationSnmp).Execute()

Update the SNMP settings for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    updateOrganizationSnmp := *openapiclient.NewInlineObject225() // InlineObject225 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.UpdateOrganizationSnmp(context.Background(), organizationId).UpdateOrganizationSnmp(updateOrganizationSnmp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateOrganizationSnmp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationSnmp`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.UpdateOrganizationSnmp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationSnmpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationSnmp** | [**InlineObject225**](InlineObject225.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

