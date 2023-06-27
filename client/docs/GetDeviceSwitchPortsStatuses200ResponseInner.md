# GetDeviceSwitchPortsStatuses200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortId** | Pointer to **string** | The string identifier of this port on the switch. This is commonly just the port number but may contain additional identifying information such as the slot and module-type if the port is located on a port module. | [optional] 
**Enabled** | Pointer to **bool** | Whether the port is configured to be enabled. | [optional] 
**Status** | Pointer to **string** | The current connection status of the port. | [optional] 
**IsUplink** | Pointer to **bool** | Whether the port is the switch&#39;s uplink. | [optional] 
**Errors** | Pointer to **[]string** | All errors present on the port. | [optional] 
**Warnings** | Pointer to **[]string** | All warnings present on the port. | [optional] 
**Speed** | Pointer to **string** | The current data transfer rate which the port is operating at. | [optional] 
**Duplex** | Pointer to **string** | The current duplex of a connected port. | [optional] 
**UsageInKb** | Pointer to [**GetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb**](GetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb.md) |  | [optional] 
**Cdp** | Pointer to [**GetDeviceSwitchPortsStatuses200ResponseInnerCdp**](GetDeviceSwitchPortsStatuses200ResponseInnerCdp.md) |  | [optional] 
**Lldp** | Pointer to [**GetDeviceSwitchPortsStatuses200ResponseInnerLldp**](GetDeviceSwitchPortsStatuses200ResponseInnerLldp.md) |  | [optional] 
**ClientCount** | Pointer to **int32** | The number of clients connected through this port. | [optional] 
**PowerUsageInWh** | Pointer to **float32** | How much power (in watt-hours) has been delivered by this port during the timespan. | [optional] 
**TrafficInKbps** | Pointer to [**GetDeviceSwitchPortsStatuses200ResponseInnerTrafficInKbps**](GetDeviceSwitchPortsStatuses200ResponseInnerTrafficInKbps.md) |  | [optional] 
**SecurePort** | Pointer to [**GetDeviceSwitchPortsStatuses200ResponseInnerSecurePort**](GetDeviceSwitchPortsStatuses200ResponseInnerSecurePort.md) |  | [optional] 

## Methods

### NewGetDeviceSwitchPortsStatuses200ResponseInner

`func NewGetDeviceSwitchPortsStatuses200ResponseInner() *GetDeviceSwitchPortsStatuses200ResponseInner`

NewGetDeviceSwitchPortsStatuses200ResponseInner instantiates a new GetDeviceSwitchPortsStatuses200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeviceSwitchPortsStatuses200ResponseInnerWithDefaults

`func NewGetDeviceSwitchPortsStatuses200ResponseInnerWithDefaults() *GetDeviceSwitchPortsStatuses200ResponseInner`

NewGetDeviceSwitchPortsStatuses200ResponseInnerWithDefaults instantiates a new GetDeviceSwitchPortsStatuses200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortId

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetEnabled

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetStatus

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIsUplink

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) GetIsUplink() bool`

GetIsUplink returns the IsUplink field if non-nil, zero value otherwise.

### GetIsUplinkOk

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) GetIsUplinkOk() (*bool, bool)`

GetIsUplinkOk returns a tuple with the IsUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUplink

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) SetIsUplink(v bool)`

SetIsUplink sets IsUplink field to given value.

### HasIsUplink

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) HasIsUplink() bool`

HasIsUplink returns a boolean if a field has been set.

### GetErrors

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetSpeed

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetDuplex

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) GetDuplex() string`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) GetDuplexOk() (*string, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) SetDuplex(v string)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetUsageInKb

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) GetUsageInKb() GetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb`

GetUsageInKb returns the UsageInKb field if non-nil, zero value otherwise.

### GetUsageInKbOk

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) GetUsageInKbOk() (*GetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb, bool)`

GetUsageInKbOk returns a tuple with the UsageInKb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageInKb

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) SetUsageInKb(v GetDeviceSwitchPortsStatuses200ResponseInnerUsageInKb)`

SetUsageInKb sets UsageInKb field to given value.

### HasUsageInKb

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) HasUsageInKb() bool`

HasUsageInKb returns a boolean if a field has been set.

### GetCdp

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) GetCdp() GetDeviceSwitchPortsStatuses200ResponseInnerCdp`

GetCdp returns the Cdp field if non-nil, zero value otherwise.

### GetCdpOk

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) GetCdpOk() (*GetDeviceSwitchPortsStatuses200ResponseInnerCdp, bool)`

GetCdpOk returns a tuple with the Cdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdp

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) SetCdp(v GetDeviceSwitchPortsStatuses200ResponseInnerCdp)`

SetCdp sets Cdp field to given value.

### HasCdp

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) HasCdp() bool`

HasCdp returns a boolean if a field has been set.

### GetLldp

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) GetLldp() GetDeviceSwitchPortsStatuses200ResponseInnerLldp`

GetLldp returns the Lldp field if non-nil, zero value otherwise.

### GetLldpOk

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) GetLldpOk() (*GetDeviceSwitchPortsStatuses200ResponseInnerLldp, bool)`

GetLldpOk returns a tuple with the Lldp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldp

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) SetLldp(v GetDeviceSwitchPortsStatuses200ResponseInnerLldp)`

SetLldp sets Lldp field to given value.

### HasLldp

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) HasLldp() bool`

HasLldp returns a boolean if a field has been set.

### GetClientCount

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) GetClientCount() int32`

GetClientCount returns the ClientCount field if non-nil, zero value otherwise.

### GetClientCountOk

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) GetClientCountOk() (*int32, bool)`

GetClientCountOk returns a tuple with the ClientCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCount

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) SetClientCount(v int32)`

SetClientCount sets ClientCount field to given value.

### HasClientCount

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) HasClientCount() bool`

HasClientCount returns a boolean if a field has been set.

### GetPowerUsageInWh

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) GetPowerUsageInWh() float32`

GetPowerUsageInWh returns the PowerUsageInWh field if non-nil, zero value otherwise.

### GetPowerUsageInWhOk

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) GetPowerUsageInWhOk() (*float32, bool)`

GetPowerUsageInWhOk returns a tuple with the PowerUsageInWh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerUsageInWh

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) SetPowerUsageInWh(v float32)`

SetPowerUsageInWh sets PowerUsageInWh field to given value.

### HasPowerUsageInWh

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) HasPowerUsageInWh() bool`

HasPowerUsageInWh returns a boolean if a field has been set.

### GetTrafficInKbps

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) GetTrafficInKbps() GetDeviceSwitchPortsStatuses200ResponseInnerTrafficInKbps`

GetTrafficInKbps returns the TrafficInKbps field if non-nil, zero value otherwise.

### GetTrafficInKbpsOk

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) GetTrafficInKbpsOk() (*GetDeviceSwitchPortsStatuses200ResponseInnerTrafficInKbps, bool)`

GetTrafficInKbpsOk returns a tuple with the TrafficInKbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficInKbps

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) SetTrafficInKbps(v GetDeviceSwitchPortsStatuses200ResponseInnerTrafficInKbps)`

SetTrafficInKbps sets TrafficInKbps field to given value.

### HasTrafficInKbps

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) HasTrafficInKbps() bool`

HasTrafficInKbps returns a boolean if a field has been set.

### GetSecurePort

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) GetSecurePort() GetDeviceSwitchPortsStatuses200ResponseInnerSecurePort`

GetSecurePort returns the SecurePort field if non-nil, zero value otherwise.

### GetSecurePortOk

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) GetSecurePortOk() (*GetDeviceSwitchPortsStatuses200ResponseInnerSecurePort, bool)`

GetSecurePortOk returns a tuple with the SecurePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurePort

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) SetSecurePort(v GetDeviceSwitchPortsStatuses200ResponseInnerSecurePort)`

SetSecurePort sets SecurePort field to given value.

### HasSecurePort

`func (o *GetDeviceSwitchPortsStatuses200ResponseInner) HasSecurePort() bool`

HasSecurePort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


