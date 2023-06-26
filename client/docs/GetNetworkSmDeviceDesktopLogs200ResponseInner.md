# GetNetworkSmDeviceDesktopLogs200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MeasuredAt** | Pointer to **string** | The time the data was measured at. | [optional] 
**User** | Pointer to **string** | The user during connection. | [optional] 
**NetworkDevice** | Pointer to **string** | The network device for the device used for connection. | [optional] 
**NetworkDriver** | Pointer to **string** | The network driver for the device. | [optional] 
**WifiChannel** | Pointer to **string** | Channel through which the connection is routing. | [optional] 
**WifiAuth** | Pointer to **string** | The type of authentication used by the SSID. | [optional] 
**WifiBssid** | Pointer to **string** | The MAC of the access point the device is connected to. | [optional] 
**WifiSsid** | Pointer to **string** | The name of the network the device is connected to. | [optional] 
**WifiRssi** | Pointer to **string** | The Received Signal Strength Indicator for the device. | [optional] 
**WifiNoise** | Pointer to **string** | The wireless signal power level received by the device. | [optional] 
**DhcpServer** | Pointer to **string** | The IP address of the DCHP Server. | [optional] 
**Ip** | Pointer to **string** | The IP of the device during connection. | [optional] 
**NetworkMTU** | Pointer to **string** | The network max transmission unit. | [optional] 
**Subnet** | Pointer to **string** | The subnet of the device connection. | [optional] 
**Gateway** | Pointer to **string** | The gateway IP the device was connected to. | [optional] 
**PublicIP** | Pointer to **string** | The public IP address of the device. | [optional] 
**DnsServer** | Pointer to **string** | The DNS Server during the connection. | [optional] 
**Ts** | Pointer to **string** | The time the connection was logged. | [optional] 

## Methods

### NewGetNetworkSmDeviceDesktopLogs200ResponseInner

`func NewGetNetworkSmDeviceDesktopLogs200ResponseInner() *GetNetworkSmDeviceDesktopLogs200ResponseInner`

NewGetNetworkSmDeviceDesktopLogs200ResponseInner instantiates a new GetNetworkSmDeviceDesktopLogs200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSmDeviceDesktopLogs200ResponseInnerWithDefaults

`func NewGetNetworkSmDeviceDesktopLogs200ResponseInnerWithDefaults() *GetNetworkSmDeviceDesktopLogs200ResponseInner`

NewGetNetworkSmDeviceDesktopLogs200ResponseInnerWithDefaults instantiates a new GetNetworkSmDeviceDesktopLogs200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeasuredAt

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) GetMeasuredAt() string`

GetMeasuredAt returns the MeasuredAt field if non-nil, zero value otherwise.

### GetMeasuredAtOk

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) GetMeasuredAtOk() (*string, bool)`

GetMeasuredAtOk returns a tuple with the MeasuredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasuredAt

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) SetMeasuredAt(v string)`

SetMeasuredAt sets MeasuredAt field to given value.

### HasMeasuredAt

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) HasMeasuredAt() bool`

HasMeasuredAt returns a boolean if a field has been set.

### GetUser

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetNetworkDevice

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) GetNetworkDevice() string`

GetNetworkDevice returns the NetworkDevice field if non-nil, zero value otherwise.

### GetNetworkDeviceOk

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) GetNetworkDeviceOk() (*string, bool)`

GetNetworkDeviceOk returns a tuple with the NetworkDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDevice

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) SetNetworkDevice(v string)`

SetNetworkDevice sets NetworkDevice field to given value.

### HasNetworkDevice

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) HasNetworkDevice() bool`

HasNetworkDevice returns a boolean if a field has been set.

### GetNetworkDriver

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) GetNetworkDriver() string`

GetNetworkDriver returns the NetworkDriver field if non-nil, zero value otherwise.

### GetNetworkDriverOk

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) GetNetworkDriverOk() (*string, bool)`

GetNetworkDriverOk returns a tuple with the NetworkDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDriver

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) SetNetworkDriver(v string)`

SetNetworkDriver sets NetworkDriver field to given value.

### HasNetworkDriver

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) HasNetworkDriver() bool`

HasNetworkDriver returns a boolean if a field has been set.

### GetWifiChannel

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) GetWifiChannel() string`

GetWifiChannel returns the WifiChannel field if non-nil, zero value otherwise.

### GetWifiChannelOk

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) GetWifiChannelOk() (*string, bool)`

GetWifiChannelOk returns a tuple with the WifiChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiChannel

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) SetWifiChannel(v string)`

SetWifiChannel sets WifiChannel field to given value.

### HasWifiChannel

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) HasWifiChannel() bool`

HasWifiChannel returns a boolean if a field has been set.

### GetWifiAuth

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) GetWifiAuth() string`

GetWifiAuth returns the WifiAuth field if non-nil, zero value otherwise.

### GetWifiAuthOk

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) GetWifiAuthOk() (*string, bool)`

GetWifiAuthOk returns a tuple with the WifiAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiAuth

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) SetWifiAuth(v string)`

SetWifiAuth sets WifiAuth field to given value.

### HasWifiAuth

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) HasWifiAuth() bool`

HasWifiAuth returns a boolean if a field has been set.

### GetWifiBssid

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) GetWifiBssid() string`

GetWifiBssid returns the WifiBssid field if non-nil, zero value otherwise.

### GetWifiBssidOk

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) GetWifiBssidOk() (*string, bool)`

GetWifiBssidOk returns a tuple with the WifiBssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiBssid

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) SetWifiBssid(v string)`

SetWifiBssid sets WifiBssid field to given value.

### HasWifiBssid

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) HasWifiBssid() bool`

HasWifiBssid returns a boolean if a field has been set.

### GetWifiSsid

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) GetWifiSsid() string`

GetWifiSsid returns the WifiSsid field if non-nil, zero value otherwise.

### GetWifiSsidOk

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) GetWifiSsidOk() (*string, bool)`

GetWifiSsidOk returns a tuple with the WifiSsid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiSsid

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) SetWifiSsid(v string)`

SetWifiSsid sets WifiSsid field to given value.

### HasWifiSsid

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) HasWifiSsid() bool`

HasWifiSsid returns a boolean if a field has been set.

### GetWifiRssi

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) GetWifiRssi() string`

GetWifiRssi returns the WifiRssi field if non-nil, zero value otherwise.

### GetWifiRssiOk

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) GetWifiRssiOk() (*string, bool)`

GetWifiRssiOk returns a tuple with the WifiRssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiRssi

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) SetWifiRssi(v string)`

SetWifiRssi sets WifiRssi field to given value.

### HasWifiRssi

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) HasWifiRssi() bool`

HasWifiRssi returns a boolean if a field has been set.

### GetWifiNoise

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) GetWifiNoise() string`

GetWifiNoise returns the WifiNoise field if non-nil, zero value otherwise.

### GetWifiNoiseOk

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) GetWifiNoiseOk() (*string, bool)`

GetWifiNoiseOk returns a tuple with the WifiNoise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiNoise

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) SetWifiNoise(v string)`

SetWifiNoise sets WifiNoise field to given value.

### HasWifiNoise

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) HasWifiNoise() bool`

HasWifiNoise returns a boolean if a field has been set.

### GetDhcpServer

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) GetDhcpServer() string`

GetDhcpServer returns the DhcpServer field if non-nil, zero value otherwise.

### GetDhcpServerOk

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) GetDhcpServerOk() (*string, bool)`

GetDhcpServerOk returns a tuple with the DhcpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServer

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) SetDhcpServer(v string)`

SetDhcpServer sets DhcpServer field to given value.

### HasDhcpServer

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) HasDhcpServer() bool`

HasDhcpServer returns a boolean if a field has been set.

### GetIp

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetNetworkMTU

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) GetNetworkMTU() string`

GetNetworkMTU returns the NetworkMTU field if non-nil, zero value otherwise.

### GetNetworkMTUOk

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) GetNetworkMTUOk() (*string, bool)`

GetNetworkMTUOk returns a tuple with the NetworkMTU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMTU

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) SetNetworkMTU(v string)`

SetNetworkMTU sets NetworkMTU field to given value.

### HasNetworkMTU

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) HasNetworkMTU() bool`

HasNetworkMTU returns a boolean if a field has been set.

### GetSubnet

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetGateway

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetPublicIP

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) GetPublicIP() string`

GetPublicIP returns the PublicIP field if non-nil, zero value otherwise.

### GetPublicIPOk

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) GetPublicIPOk() (*string, bool)`

GetPublicIPOk returns a tuple with the PublicIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIP

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) SetPublicIP(v string)`

SetPublicIP sets PublicIP field to given value.

### HasPublicIP

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) HasPublicIP() bool`

HasPublicIP returns a boolean if a field has been set.

### GetDnsServer

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) GetDnsServer() string`

GetDnsServer returns the DnsServer field if non-nil, zero value otherwise.

### GetDnsServerOk

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) GetDnsServerOk() (*string, bool)`

GetDnsServerOk returns a tuple with the DnsServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServer

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) SetDnsServer(v string)`

SetDnsServer sets DnsServer field to given value.

### HasDnsServer

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) HasDnsServer() bool`

HasDnsServer returns a boolean if a field has been set.

### GetTs

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetNetworkSmDeviceDesktopLogs200ResponseInner) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


