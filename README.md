# Go API client for openapi

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 4.1.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import openapi "github.com/hibiki31/virty-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `openapi.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `openapi.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `openapi.ContextOperationServerIndices` and `openapi.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://virty-pr.hinagiku.me/api*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AuthAPI* | [**Login**](docs/AuthAPI.md#login) | **Post** /api/auth | Login For Access Token
*AuthAPI* | [**Setup**](docs/AuthAPI.md#setup) | **Post** /api/auth/setup | Api Auth Setup
*AuthAPI* | [**ValidateToken**](docs/AuthAPI.md#validatetoken) | **Get** /api/auth/validate | Read Auth Validate
*FlavorsAPI* | [**CreateFlavor**](docs/FlavorsAPI.md#createflavor) | **Post** /api/flavors | Post Api Flavors
*FlavorsAPI* | [**DeleteFlavor**](docs/FlavorsAPI.md#deleteflavor) | **Delete** /api/flavors/{flavor_id} | Delete Flavors
*FlavorsAPI* | [**GetFlavors**](docs/FlavorsAPI.md#getflavors) | **Get** /api/flavors | Get Api Flavors
*ImagesAPI* | [**GetImages**](docs/ImagesAPI.md#getimages) | **Get** /api/images | Get Api Images
*ImagesAPI* | [**ScpImage**](docs/ImagesAPI.md#scpimage) | **Put** /api/images/scp | Put Api Images Scp
*ImagesAPI* | [**UpdateImageFlavor**](docs/ImagesAPI.md#updateimageflavor) | **Patch** /api/images | Patch Api Images
*ImagesTaskAPI* | [**DownloadImage**](docs/ImagesTaskAPI.md#downloadimage) | **Post** /api/tasks/images/download | Post Image Download
*ImagesTaskAPI* | [**RefreshImages**](docs/ImagesTaskAPI.md#refreshimages) | **Put** /api/tasks/images | Put Api Images
*MetricsAPI* | [**GetMetrics**](docs/MetricsAPI.md#getmetrics) | **Get** /api/metrics | Exporter Get
*MixinAPI* | [**GetVersion**](docs/MixinAPI.md#getversion) | **Get** /api/version | Get Version
*NetworksAPI* | [**CreateNetworkPool**](docs/NetworksAPI.md#createnetworkpool) | **Post** /api/networks/pools | Post Api Networks Pools
*NetworksAPI* | [**DeleteNetworkPool**](docs/NetworksAPI.md#deletenetworkpool) | **Delete** /api/networks/pools/{id} | Delete Pools Uuid
*NetworksAPI* | [**GetNetwork**](docs/NetworksAPI.md#getnetwork) | **Get** /api/networks/{uuid} | Get Api Networks Uuid
*NetworksAPI* | [**GetNetworkPools**](docs/NetworksAPI.md#getnetworkpools) | **Get** /api/networks/pools | Get Api Networks Pools
*NetworksAPI* | [**GetNetworks**](docs/NetworksAPI.md#getnetworks) | **Get** /api/networks | Get Api Networks
*NetworksAPI* | [**UpdateNetworkPool**](docs/NetworksAPI.md#updatenetworkpool) | **Patch** /api/networks/pools | Patch Api Networks Pools
*NetworksTaskAPI* | [**CreateNetwork**](docs/NetworksTaskAPI.md#createnetwork) | **Post** /api/tasks/networks | Post Api Storage
*NetworksTaskAPI* | [**CreateNetworkOvs**](docs/NetworksTaskAPI.md#createnetworkovs) | **Post** /api/tasks/networks/{uuid}/ovs | Post Uuid Ovs
*NetworksTaskAPI* | [**DeleteNetwork**](docs/NetworksTaskAPI.md#deletenetwork) | **Delete** /api/tasks/networks/{uuid} | Delete Api Storage
*NetworksTaskAPI* | [**DeleteNetworkOvs**](docs/NetworksTaskAPI.md#deletenetworkovs) | **Delete** /api/tasks/networks/{uuid}/ovs/{name} | Post Api Networks Uuid Ovs
*NetworksTaskAPI* | [**PostUuidOvsApiTasksNetworksProvidersPost**](docs/NetworksTaskAPI.md#postuuidovsapitasksnetworksproviderspost) | **Post** /api/tasks/networks/providers | Post Uuid Ovs
*NetworksTaskAPI* | [**RefreshNetworks**](docs/NetworksTaskAPI.md#refreshnetworks) | **Put** /api/tasks/networks | Put Api Networks
*NodesAPI* | [**GetNode**](docs/NodesAPI.md#getnode) | **Get** /api/nodes/{name} | Get Api Node
*NodesAPI* | [**GetNodeFacts**](docs/NodesAPI.md#getnodefacts) | **Get** /api/nodes/{name}/facts | Get Node Name Facts
*NodesAPI* | [**GetNodeNameFactsApiNodesNameNetworkGet**](docs/NodesAPI.md#getnodenamefactsapinodesnamenetworkget) | **Get** /api/nodes/{name}/network | Get Node Name Facts
*NodesAPI* | [**GetNodes**](docs/NodesAPI.md#getnodes) | **Get** /api/nodes | Get Api Nodes
*NodesAPI* | [**GetSshKeyPair**](docs/NodesAPI.md#getsshkeypair) | **Get** /api/nodes/key | Get Ssh Key Pair
*NodesAPI* | [**UpdateSshKeyPair**](docs/NodesAPI.md#updatesshkeypair) | **Post** /api/nodes/key | Post Ssh Key Pair
*NodesTaskAPI* | [**CreateNode**](docs/NodesTaskAPI.md#createnode) | **Post** /api/tasks/nodes | Post Tasks Nodes
*NodesTaskAPI* | [**DeleteNode**](docs/NodesTaskAPI.md#deletenode) | **Delete** /api/tasks/nodes/{name} | Delete Tasks Nodes Name
*NodesTaskAPI* | [**UpdateNodeRole**](docs/NodesTaskAPI.md#updatenoderole) | **Patch** /api/tasks/nodes/roles | Patch Api Node Role
*ProjectsAPI* | [**CreateProject**](docs/ProjectsAPI.md#createproject) | **Post** /api/tasks/projects | Post Api Projects
*ProjectsAPI* | [**DeleteProject**](docs/ProjectsAPI.md#deleteproject) | **Delete** /api/tasks/projects/{project_id} | Delete Api Projects
*ProjectsAPI* | [**GetProjects**](docs/ProjectsAPI.md#getprojects) | **Get** /api/projects | Get Api Projects
*ProjectsAPI* | [**UpdateProject**](docs/ProjectsAPI.md#updateproject) | **Put** /api/projects | Put Api Projects
*StoragesAPI* | [**CreateStoragePool**](docs/StoragesAPI.md#createstoragepool) | **Post** /api/storages/pools | Post Api Storages Pools
*StoragesAPI* | [**GetStorage**](docs/StoragesAPI.md#getstorage) | **Get** /api/storages/{uuid} | Get Api Storages Uuid
*StoragesAPI* | [**GetStoragePools**](docs/StoragesAPI.md#getstoragepools) | **Get** /api/storages/pools | Get Api Storages Pools
*StoragesAPI* | [**GetStorages**](docs/StoragesAPI.md#getstorages) | **Get** /api/storages | Get Api Storages
*StoragesAPI* | [**UpdateStorageMetadata**](docs/StoragesAPI.md#updatestoragemetadata) | **Patch** /api/storages | Post Api Storage
*StoragesAPI* | [**UpdateStoragePool**](docs/StoragesAPI.md#updatestoragepool) | **Patch** /api/storages/pools | Post Api Storages Pools
*StoragesTaskAPI* | [**CreateStorage**](docs/StoragesTaskAPI.md#createstorage) | **Post** /api/tasks/storages | Post Api Storage
*StoragesTaskAPI* | [**DeleteStorage**](docs/StoragesTaskAPI.md#deletestorage) | **Delete** /api/tasks/storages/{uuid} | Delete Api Storages
*TasksAPI* | [**DeleteTasks**](docs/TasksAPI.md#deletetasks) | **Delete** /api/tasks/ | Delete Tasks
*TasksAPI* | [**GetIncompleteTasks**](docs/TasksAPI.md#getincompletetasks) | **Get** /api/tasks/incomplete | Get Tasks Incomplete
*TasksAPI* | [**GetTask**](docs/TasksAPI.md#gettask) | **Get** /api/tasks/{uuid} | Get Tasks
*TasksAPI* | [**GetTasks**](docs/TasksAPI.md#gettasks) | **Get** /api/tasks | Get Tasks
*UsersAPI* | [**CreateUser**](docs/UsersAPI.md#createuser) | **Post** /api/users | Post Api Users
*UsersAPI* | [**DeleteUser**](docs/UsersAPI.md#deleteuser) | **Delete** /api/users/{username} | Delete User
*UsersAPI* | [**GetCurrentUser**](docs/UsersAPI.md#getcurrentuser) | **Get** /api/users/me | Read Users Me
*UsersAPI* | [**GetUsers**](docs/UsersAPI.md#getusers) | **Get** /api/users | Get Api Users
*VmsAPI* | [**GetVm**](docs/VmsAPI.md#getvm) | **Get** /api/vms/{uuid} | Get Api Domain Uuid
*VmsAPI* | [**GetVms**](docs/VmsAPI.md#getvms) | **Get** /api/vms | Get Api Domain
*VmsAPI* | [**GetVncAddress**](docs/VmsAPI.md#getvncaddress) | **Get** /api/vms/vnc/{token} | Get Api Domain
*VmsTaskAPI* | [**ControlVmCdrom**](docs/VmsTaskAPI.md#controlvmcdrom) | **Patch** /api/tasks/vms/{uuid}/cdrom | Patch Api Tasks Vms Uuid Cdrom
*VmsTaskAPI* | [**CreateVm**](docs/VmsTaskAPI.md#createvm) | **Post** /api/tasks/vms | Post Api Vms
*VmsTaskAPI* | [**DeleteVm**](docs/VmsTaskAPI.md#deletevm) | **Delete** /api/tasks/vms/{uuid} | Delete Api Domains
*VmsTaskAPI* | [**PathVmsProjectApiTasksVmsProjectPatch**](docs/VmsTaskAPI.md#pathvmsprojectapitasksvmsprojectpatch) | **Patch** /api/tasks/vms/project | Path Vms Project
*VmsTaskAPI* | [**RefreshVms**](docs/VmsTaskAPI.md#refreshvms) | **Put** /api/tasks/vms | Publish Task To Update Vm List
*VmsTaskAPI* | [**UpdateVmNetwork**](docs/VmsTaskAPI.md#updatevmnetwork) | **Patch** /api/tasks/vms/{uuid}/network | Patch Api Vm Network
*VmsTaskAPI* | [**UpdateVmPowerStatus**](docs/VmsTaskAPI.md#updatevmpowerstatus) | **Patch** /api/tasks/vms/{uuid}/power | Patch Api Tasks Vms Uuid Power


## Documentation For Models

 - [AuthValidateResponse](docs/AuthValidateResponse.md)
 - [CdromForUpdateDomain](docs/CdromForUpdateDomain.md)
 - [CloudInitInsert](docs/CloudInitInsert.md)
 - [Domain](docs/Domain.md)
 - [DomainDetail](docs/DomainDetail.md)
 - [DomainDrive](docs/DomainDrive.md)
 - [DomainForCreate](docs/DomainForCreate.md)
 - [DomainForCreateDisk](docs/DomainForCreateDisk.md)
 - [DomainForCreateInterface](docs/DomainForCreateInterface.md)
 - [DomainInterface](docs/DomainInterface.md)
 - [DomainPagenation](docs/DomainPagenation.md)
 - [DomainProject](docs/DomainProject.md)
 - [DomainProjectForUpdate](docs/DomainProjectForUpdate.md)
 - [Flavor](docs/Flavor.md)
 - [FlavorForCreate](docs/FlavorForCreate.md)
 - [FlavorPage](docs/FlavorPage.md)
 - [HTTPValidationError](docs/HTTPValidationError.md)
 - [Image](docs/Image.md)
 - [ImageDomain](docs/ImageDomain.md)
 - [ImageDownloadForCreate](docs/ImageDownloadForCreate.md)
 - [ImageForUpdateImageFlavor](docs/ImageForUpdateImageFlavor.md)
 - [ImagePage](docs/ImagePage.md)
 - [ImageSCP](docs/ImageSCP.md)
 - [Network](docs/Network.md)
 - [NetworkForCreate](docs/NetworkForCreate.md)
 - [NetworkForNetworkPool](docs/NetworkForNetworkPool.md)
 - [NetworkForUpdateDomain](docs/NetworkForUpdateDomain.md)
 - [NetworkOVSForCreate](docs/NetworkOVSForCreate.md)
 - [NetworkPage](docs/NetworkPage.md)
 - [NetworkPool](docs/NetworkPool.md)
 - [NetworkPoolForCreate](docs/NetworkPoolForCreate.md)
 - [NetworkPoolForUpdate](docs/NetworkPoolForUpdate.md)
 - [NetworkPoolPort](docs/NetworkPoolPort.md)
 - [NetworkPortgroup](docs/NetworkPortgroup.md)
 - [NetworkProviderForCreate](docs/NetworkProviderForCreate.md)
 - [Node](docs/Node.md)
 - [NodeForCreate](docs/NodeForCreate.md)
 - [NodeInterface](docs/NodeInterface.md)
 - [NodeInterfaceIpv4Info](docs/NodeInterfaceIpv4Info.md)
 - [NodeInterfaceIpv6Info](docs/NodeInterfaceIpv6Info.md)
 - [NodePage](docs/NodePage.md)
 - [NodeRole](docs/NodeRole.md)
 - [NodeRoleForUpdate](docs/NodeRoleForUpdate.md)
 - [PowerStatusForUpdateDomain](docs/PowerStatusForUpdateDomain.md)
 - [Project](docs/Project.md)
 - [ProjectForCreate](docs/ProjectForCreate.md)
 - [ProjectForUpdate](docs/ProjectForUpdate.md)
 - [ProjectPage](docs/ProjectPage.md)
 - [ProjectUser](docs/ProjectUser.md)
 - [SSHKeyPair](docs/SSHKeyPair.md)
 - [SetupRequest](docs/SetupRequest.md)
 - [Storage](docs/Storage.md)
 - [StorageContainerForStoragePool](docs/StorageContainerForStoragePool.md)
 - [StorageForCreate](docs/StorageForCreate.md)
 - [StorageForStorageContainer](docs/StorageForStorageContainer.md)
 - [StorageMetadata](docs/StorageMetadata.md)
 - [StorageMetadataForUpdate](docs/StorageMetadataForUpdate.md)
 - [StoragePage](docs/StoragePage.md)
 - [StoragePool](docs/StoragePool.md)
 - [StoragePoolForCreate](docs/StoragePoolForCreate.md)
 - [StoragePoolForUpdate](docs/StoragePoolForUpdate.md)
 - [Task](docs/Task.md)
 - [TaskIncomplete](docs/TaskIncomplete.md)
 - [TaskPagesnation](docs/TaskPagesnation.md)
 - [TokenData](docs/TokenData.md)
 - [TokenRFC6749Response](docs/TokenRFC6749Response.md)
 - [User](docs/User.md)
 - [UserForCreate](docs/UserForCreate.md)
 - [UserPage](docs/UserPage.md)
 - [UserProject](docs/UserProject.md)
 - [UserScope](docs/UserScope.md)
 - [ValidationError](docs/ValidationError.md)
 - [Version](docs/Version.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### OAuth2PasswordBearer


- **Type**: OAuth
- **Flow**: password
- **Authorization URL**: 
- **Scopes**: 
 - **admin**: Have all authority
 - **user**: User authority

Example

```go
auth := context.WithValue(context.Background(), openapi.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```go
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, openapi.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



