# Organization Inventory Report

This code generates reports of organizations, including tabs for associated networks, devices, and licenses. 
The data is fetched concurrently from the Meraki Dashboard API and saved in Excel spreadsheets using the excelize library.


## Download & Run Report

This example is offered to users as a single executable file so no programming experience is required for usage.

This is because Go allows us to compile cross-platform static binaries, so it will work on Linux, MacOS, and Windows. 

Simply download the correct binary for your operating system and run it with your apikey for authentication:

| Operating System | Architecture | intel                                                  | arm                                                    |
|------------------|--------------|--------------------------------------------------------|--------------------------------------------------------|
| Linux            | 64-bit       | [download](build/organizationReport-linux-amd64)       | [download](build/organizationReport-linux-arm64)       |
| MacOS            | 64-bit       | [download](build/organizationReport-macos-amd64)       | [download](build/organizationReport-macos-arm64)       |
| Windows          | 64-bit       | [download](build/organizationReport-windows-amd64.exe) | [download](build/organizationReport-windows-arm64.exe) |



Exporting your Meraki API key to an environmental variable is the recommended workflow:  

```shell

# linux / macos
MERAKI_DASHBOARD_API_KEY=093b24e85df15a3e66f1fc359f4c48493eaa1b73

# windows
set MERAKI_DASHBOARD_API_KEY=093b24e85df15a3e66f1fc359f4c48493eaa1b73

```



Alternatively, you can use the `-apikey` flag as described in the cli flags help message:

```text
-apikey string, Meraki dashboard api key. This defaults to the env var: 'MERAKI_DASHBOARD_API_KEY'

-help bool, Prints usage information

-out string,   output directory to generate report. Defaults to current working directory.

```




Running the report:

```shell

# linux / macos
./OrganizationReport-{OS}-{ARCH} -apikey=093b24e85df15a3e66f1fc359f4c48493eaa1b73 -out=$PWD/out

# windows
.\OrganizationReport-windows-{ARCH}.exe -apikey=093b24e85df15a3e66f1fc359f4c48493eaa1b73

```
For reference: 

* {OS} = "linux", "macos", or "windows" 
* {ARCH} = "arm64" or "amd64"

**For example, the M1/M2 MacBook binary is [organizationReport-macos-arm64](build/organizationReport-macos-arm64)**




## Generating new Binaries (Optional)

For those who wish to modify this code and generate new binaries, a Makefile is included for convenience: 

```shell

make clean
make build

```