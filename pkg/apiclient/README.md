# Go API client for apiclient

Daytona Server API

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v0.0.0-dev
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import apiclient "github.com/GIT_USER_ID/GIT_REPO_ID/apiclient"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `apiclient.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), apiclient.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `apiclient.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), apiclient.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `apiclient.ContextOperationServerIndices` and `apiclient.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), apiclient.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), apiclient.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:3986*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ApiKeyAPI* | [**CreateApiKey**](docs/ApiKeyAPI.md#createapikey) | **Post** /apikey/{apiKeyName} | Create an API key
*ApiKeyAPI* | [**DeleteApiKey**](docs/ApiKeyAPI.md#deleteapikey) | **Delete** /apikey/{apiKeyName} | Delete API key
*ApiKeyAPI* | [**ListClientApiKeys**](docs/ApiKeyAPI.md#listclientapikeys) | **Get** /apikey | List API keys
*BuildAPI* | [**CreateBuild**](docs/BuildAPI.md#createbuild) | **Post** /build | Create a build
*BuildAPI* | [**DeleteAllBuilds**](docs/BuildAPI.md#deleteallbuilds) | **Delete** /build | Delete ALL builds
*BuildAPI* | [**DeleteBuild**](docs/BuildAPI.md#deletebuild) | **Delete** /build/{buildId} | Delete build
*BuildAPI* | [**DeleteBuildsFromPrebuild**](docs/BuildAPI.md#deletebuildsfromprebuild) | **Delete** /build/prebuild/{prebuildId} | Delete builds
*BuildAPI* | [**FindBuild**](docs/BuildAPI.md#findbuild) | **Get** /build/{buildId} | Find build
*BuildAPI* | [**ListBuilds**](docs/BuildAPI.md#listbuilds) | **Get** /build | List builds
*BuildAPI* | [**ListSuccessfulBuilds**](docs/BuildAPI.md#listsuccessfulbuilds) | **Get** /build/successful/{repoUrl} | List successful builds for Git repository
*ContainerRegistryAPI* | [**FindContainerRegistry**](docs/ContainerRegistryAPI.md#findcontainerregistry) | **Get** /container-registry/{server} | Find container registry
*DefaultAPI* | [**HealthCheck**](docs/DefaultAPI.md#healthcheck) | **Get** /health | Health check
*EnvVarAPI* | [**DeleteEnvironmentVariable**](docs/EnvVarAPI.md#deleteenvironmentvariable) | **Delete** /env/{key} | Delete environment variable
*EnvVarAPI* | [**ListEnvironmentVariables**](docs/EnvVarAPI.md#listenvironmentvariables) | **Get** /env | List environment variables
*EnvVarAPI* | [**SaveEnvironmentVariable**](docs/EnvVarAPI.md#saveenvironmentvariable) | **Put** /env | Save environment variable
*GitProviderAPI* | [**DeleteGitProvider**](docs/GitProviderAPI.md#deletegitprovider) | **Delete** /gitprovider/{gitProviderId} | Delete Git provider
*GitProviderAPI* | [**FindGitProvider**](docs/GitProviderAPI.md#findgitprovider) | **Get** /gitprovider/{gitProviderId} | Find Git provider
*GitProviderAPI* | [**FindGitProviderIdForUrl**](docs/GitProviderAPI.md#findgitprovideridforurl) | **Get** /gitprovider/id-for-url/{url} | Find Git provider ID
*GitProviderAPI* | [**GetGitContext**](docs/GitProviderAPI.md#getgitcontext) | **Post** /gitprovider/context | Get Git context
*GitProviderAPI* | [**GetGitUser**](docs/GitProviderAPI.md#getgituser) | **Get** /gitprovider/{gitProviderId}/user | Get Git context
*GitProviderAPI* | [**GetNamespaces**](docs/GitProviderAPI.md#getnamespaces) | **Get** /gitprovider/{gitProviderId}/namespaces | Get Git namespaces
*GitProviderAPI* | [**GetRepoBranches**](docs/GitProviderAPI.md#getrepobranches) | **Get** /gitprovider/{gitProviderId}/{namespaceId}/{repositoryId}/branches | Get Git repository branches
*GitProviderAPI* | [**GetRepoPRs**](docs/GitProviderAPI.md#getrepoprs) | **Get** /gitprovider/{gitProviderId}/{namespaceId}/{repositoryId}/pull-requests | Get Git repository PRs
*GitProviderAPI* | [**GetRepositories**](docs/GitProviderAPI.md#getrepositories) | **Get** /gitprovider/{gitProviderId}/{namespaceId}/repositories | Get Git repositories
*GitProviderAPI* | [**GetUrlFromRepository**](docs/GitProviderAPI.md#geturlfromrepository) | **Post** /gitprovider/context/url | Get URL from Git repository
*GitProviderAPI* | [**ListGitProviders**](docs/GitProviderAPI.md#listgitproviders) | **Get** /gitprovider | List Git providers
*GitProviderAPI* | [**ListGitProvidersForUrl**](docs/GitProviderAPI.md#listgitprovidersforurl) | **Get** /gitprovider/for-url/{url} | List Git providers for url
*GitProviderAPI* | [**SaveGitProvider**](docs/GitProviderAPI.md#savegitprovider) | **Put** /gitprovider | Save Git provider
*JobAPI* | [**ListJobs**](docs/JobAPI.md#listjobs) | **Get** /job | List jobs
*PrebuildAPI* | [**DeletePrebuild**](docs/PrebuildAPI.md#deleteprebuild) | **Delete** /workspace-template/{templateName}/prebuild/{prebuildId} | Delete prebuild
*PrebuildAPI* | [**FindPrebuild**](docs/PrebuildAPI.md#findprebuild) | **Get** /workspace-template/{templateName}/prebuild/{prebuildId} | Find prebuild
*PrebuildAPI* | [**ListPrebuilds**](docs/PrebuildAPI.md#listprebuilds) | **Get** /workspace-template/prebuild | List prebuilds
*PrebuildAPI* | [**ListPrebuildsForWorkspaceTemplate**](docs/PrebuildAPI.md#listprebuildsforworkspacetemplate) | **Get** /workspace-template/{templateName}/prebuild | List prebuilds for workspace template
*PrebuildAPI* | [**ProcessGitEvent**](docs/PrebuildAPI.md#processgitevent) | **Post** /workspace-template/prebuild/process-git-event | ProcessGitEvent
*PrebuildAPI* | [**SavePrebuild**](docs/PrebuildAPI.md#saveprebuild) | **Put** /workspace-template/{templateName}/prebuild | Save prebuild
*ProviderAPI* | [**GetRunnerProviders**](docs/ProviderAPI.md#getrunnerproviders) | **Get** /runner/{runnerId}/provider | Get runner providers
*ProviderAPI* | [**InstallProvider**](docs/ProviderAPI.md#installprovider) | **Post** /runner/{runnerId}/provider/{providerName}/install | Install provider
*ProviderAPI* | [**ListProviders**](docs/ProviderAPI.md#listproviders) | **Get** /runner/provider | List providers
*ProviderAPI* | [**ListProvidersForInstall**](docs/ProviderAPI.md#listprovidersforinstall) | **Get** /runner/provider/for-install | List providers available for installation
*ProviderAPI* | [**UninstallProvider**](docs/ProviderAPI.md#uninstallprovider) | **Post** /runner/{runnerId}/provider/{providerName}/uninstall | Uninstall provider
*ProviderAPI* | [**UpdateProvider**](docs/ProviderAPI.md#updateprovider) | **Post** /runner/{runnerId}/provider/{providerName}/update | Update provider
*RunnerAPI* | [**CreateRunner**](docs/RunnerAPI.md#createrunner) | **Post** /runner | Create a runner
*RunnerAPI* | [**DeleteRunner**](docs/RunnerAPI.md#deleterunner) | **Delete** /runner/{runnerId} | Delete runner
*RunnerAPI* | [**FindRunner**](docs/RunnerAPI.md#findrunner) | **Get** /runner/{runnerId} | Find a runner
*RunnerAPI* | [**ListRunnerJobs**](docs/RunnerAPI.md#listrunnerjobs) | **Get** /runner/{runnerId}/jobs | List runner jobs
*RunnerAPI* | [**ListRunners**](docs/RunnerAPI.md#listrunners) | **Get** /runner | List runners
*RunnerAPI* | [**UpdateJobState**](docs/RunnerAPI.md#updatejobstate) | **Post** /runner/{runnerId}/jobs/{jobId}/state | Update job state
*RunnerAPI* | [**UpdateRunnerMetadata**](docs/RunnerAPI.md#updaterunnermetadata) | **Post** /runner/{runnerId}/metadata | Update runner metadata
*SampleAPI* | [**ListSamples**](docs/SampleAPI.md#listsamples) | **Get** /sample | List samples
*ServerAPI* | [**CreateNetworkKey**](docs/ServerAPI.md#createnetworkkey) | **Post** /server/network-key | Create a new authentication key
*ServerAPI* | [**GetConfig**](docs/ServerAPI.md#getconfig) | **Get** /server/config | Get the server configuration
*ServerAPI* | [**GetServerLogFiles**](docs/ServerAPI.md#getserverlogfiles) | **Get** /server/logs | Get server log files
*ServerAPI* | [**SaveConfig**](docs/ServerAPI.md#saveconfig) | **Put** /server/config | Save the server configuration
*TargetAPI* | [**CreateTarget**](docs/TargetAPI.md#createtarget) | **Post** /target | Create a target
*TargetAPI* | [**DeleteTarget**](docs/TargetAPI.md#deletetarget) | **Delete** /target/{targetId} | Delete target
*TargetAPI* | [**FindTarget**](docs/TargetAPI.md#findtarget) | **Get** /target/{targetId} | Find target
*TargetAPI* | [**GetTargetState**](docs/TargetAPI.md#gettargetstate) | **Get** /target/{targetId}/state | Get target state
*TargetAPI* | [**HandleSuccessfulCreation**](docs/TargetAPI.md#handlesuccessfulcreation) | **Post** /target/{targetId}/handle-successful-creation | Handles successful creation of the target
*TargetAPI* | [**ListTargets**](docs/TargetAPI.md#listtargets) | **Get** /target | List targets
*TargetAPI* | [**RestartTarget**](docs/TargetAPI.md#restarttarget) | **Post** /target/{targetId}/restart | Restart target
*TargetAPI* | [**SetDefaultTarget**](docs/TargetAPI.md#setdefaulttarget) | **Patch** /target/{targetId}/set-default | Set target to be used by default
*TargetAPI* | [**StartTarget**](docs/TargetAPI.md#starttarget) | **Post** /target/{targetId}/start | Start target
*TargetAPI* | [**StopTarget**](docs/TargetAPI.md#stoptarget) | **Post** /target/{targetId}/stop | Stop target
*TargetAPI* | [**UpdateTargetMetadata**](docs/TargetAPI.md#updatetargetmetadata) | **Post** /target/{targetId}/metadata | Update target metadata
*TargetAPI* | [**UpdateTargetProviderMetadata**](docs/TargetAPI.md#updatetargetprovidermetadata) | **Post** /target/{targetId}/provider-metadata | Update target provider metadata
*TargetConfigAPI* | [**CreateTargetConfig**](docs/TargetConfigAPI.md#createtargetconfig) | **Post** /target-config | Create a target config
*TargetConfigAPI* | [**DeleteTargetConfig**](docs/TargetConfigAPI.md#deletetargetconfig) | **Delete** /target-config/{configId} | Delete a target config
*TargetConfigAPI* | [**ListTargetConfigs**](docs/TargetConfigAPI.md#listtargetconfigs) | **Get** /target-config | List target configs
*WorkspaceAPI* | [**CreateWorkspace**](docs/WorkspaceAPI.md#createworkspace) | **Post** /workspace | Create a workspace
*WorkspaceAPI* | [**DeleteWorkspace**](docs/WorkspaceAPI.md#deleteworkspace) | **Delete** /workspace/{workspaceId} | Delete workspace
*WorkspaceAPI* | [**FindWorkspace**](docs/WorkspaceAPI.md#findworkspace) | **Get** /workspace/{workspaceId} | Find workspace
*WorkspaceAPI* | [**GetWorkspaceState**](docs/WorkspaceAPI.md#getworkspacestate) | **Get** /workspace/{workspaceId}/state | Get workspace state
*WorkspaceAPI* | [**ListWorkspaces**](docs/WorkspaceAPI.md#listworkspaces) | **Get** /workspace | List workspaces
*WorkspaceAPI* | [**RestartWorkspace**](docs/WorkspaceAPI.md#restartworkspace) | **Post** /workspace/{workspaceId}/restart | Restart workspace
*WorkspaceAPI* | [**StartWorkspace**](docs/WorkspaceAPI.md#startworkspace) | **Post** /workspace/{workspaceId}/start | Start workspace
*WorkspaceAPI* | [**StopWorkspace**](docs/WorkspaceAPI.md#stopworkspace) | **Post** /workspace/{workspaceId}/stop | Stop workspace
*WorkspaceAPI* | [**UpdateWorkspaceLabels**](docs/WorkspaceAPI.md#updateworkspacelabels) | **Post** /workspace/{workspaceId}/labels | Update workspace labels
*WorkspaceAPI* | [**UpdateWorkspaceMetadata**](docs/WorkspaceAPI.md#updateworkspacemetadata) | **Post** /workspace/{workspaceId}/metadata | Update workspace metadata
*WorkspaceAPI* | [**UpdateWorkspaceProviderMetadata**](docs/WorkspaceAPI.md#updateworkspaceprovidermetadata) | **Post** /workspace/{workspaceId}/provider-metadata | Update workspace provider metadata
*WorkspaceTemplateAPI* | [**DeleteWorkspaceTemplate**](docs/WorkspaceTemplateAPI.md#deleteworkspacetemplate) | **Delete** /workspace-template/{templateName} | Delete workspace template data
*WorkspaceTemplateAPI* | [**FindWorkspaceTemplate**](docs/WorkspaceTemplateAPI.md#findworkspacetemplate) | **Get** /workspace-template/{templateName} | Find a workspace template
*WorkspaceTemplateAPI* | [**GetDefaultWorkspaceTemplate**](docs/WorkspaceTemplateAPI.md#getdefaultworkspacetemplate) | **Get** /workspace-template/default/{gitUrl} | Get default workspace templates by git url
*WorkspaceTemplateAPI* | [**ListWorkspaceTemplates**](docs/WorkspaceTemplateAPI.md#listworkspacetemplates) | **Get** /workspace-template | List workspace templates
*WorkspaceTemplateAPI* | [**SaveWorkspaceTemplate**](docs/WorkspaceTemplateAPI.md#saveworkspacetemplate) | **Put** /workspace-template | Set workspace template data
*WorkspaceTemplateAPI* | [**SetDefaultWorkspaceTemplate**](docs/WorkspaceTemplateAPI.md#setdefaultworkspacetemplate) | **Patch** /workspace-template/{templateName}/set-default | Set workspace template to default
*WorkspaceToolboxAPI* | [**CreateSession**](docs/WorkspaceToolboxAPI.md#createsession) | **Post** /workspace/{workspaceId}/toolbox/process/session | Create exec session
*WorkspaceToolboxAPI* | [**DeleteSession**](docs/WorkspaceToolboxAPI.md#deletesession) | **Delete** /workspace/{workspaceId}/toolbox/process/session/{sessionId} | Delete session
*WorkspaceToolboxAPI* | [**FsCreateFolder**](docs/WorkspaceToolboxAPI.md#fscreatefolder) | **Post** /workspace/{workspaceId}/toolbox/files/folder | Create folder
*WorkspaceToolboxAPI* | [**FsDeleteFile**](docs/WorkspaceToolboxAPI.md#fsdeletefile) | **Delete** /workspace/{workspaceId}/toolbox/files | Delete file
*WorkspaceToolboxAPI* | [**FsDownloadFile**](docs/WorkspaceToolboxAPI.md#fsdownloadfile) | **Get** /workspace/{workspaceId}/toolbox/files/download | Download file
*WorkspaceToolboxAPI* | [**FsFindInFiles**](docs/WorkspaceToolboxAPI.md#fsfindinfiles) | **Get** /workspace/{workspaceId}/toolbox/files/find | Search for text/pattern in files
*WorkspaceToolboxAPI* | [**FsGetFileDetails**](docs/WorkspaceToolboxAPI.md#fsgetfiledetails) | **Get** /workspace/{workspaceId}/toolbox/files/info | Get file info
*WorkspaceToolboxAPI* | [**FsListFiles**](docs/WorkspaceToolboxAPI.md#fslistfiles) | **Get** /workspace/{workspaceId}/toolbox/files | List files
*WorkspaceToolboxAPI* | [**FsMoveFile**](docs/WorkspaceToolboxAPI.md#fsmovefile) | **Post** /workspace/{workspaceId}/toolbox/files/move | Create folder
*WorkspaceToolboxAPI* | [**FsReplaceInFiles**](docs/WorkspaceToolboxAPI.md#fsreplaceinfiles) | **Post** /workspace/{workspaceId}/toolbox/files/replace | Repleace text/pattern in files
*WorkspaceToolboxAPI* | [**FsSearchFiles**](docs/WorkspaceToolboxAPI.md#fssearchfiles) | **Get** /workspace/{workspaceId}/toolbox/files/search | Search for files
*WorkspaceToolboxAPI* | [**FsSetFilePermissions**](docs/WorkspaceToolboxAPI.md#fssetfilepermissions) | **Post** /workspace/{workspaceId}/toolbox/files/permissions | Set file owner/group/permissions
*WorkspaceToolboxAPI* | [**FsUploadFile**](docs/WorkspaceToolboxAPI.md#fsuploadfile) | **Post** /workspace/{workspaceId}/toolbox/files/upload | Upload file
*WorkspaceToolboxAPI* | [**GetSessionCommandLogs**](docs/WorkspaceToolboxAPI.md#getsessioncommandlogs) | **Get** /workspace/{workspaceId}/toolbox/process/session/{sessionId}/command/{commandId}/logs | Get session command logs
*WorkspaceToolboxAPI* | [**GetWorkspaceDir**](docs/WorkspaceToolboxAPI.md#getworkspacedir) | **Get** /workspace/{workspaceId}/toolbox/workspace-dir | Get workspace dir
*WorkspaceToolboxAPI* | [**GitAddFiles**](docs/WorkspaceToolboxAPI.md#gitaddfiles) | **Post** /workspace/{workspaceId}/toolbox/git/add | Add files
*WorkspaceToolboxAPI* | [**GitBranchList**](docs/WorkspaceToolboxAPI.md#gitbranchlist) | **Get** /workspace/{workspaceId}/toolbox/git/branches | Get branch list
*WorkspaceToolboxAPI* | [**GitCloneRepository**](docs/WorkspaceToolboxAPI.md#gitclonerepository) | **Post** /workspace/{workspaceId}/toolbox/git/clone | Clone git repository
*WorkspaceToolboxAPI* | [**GitCommitChanges**](docs/WorkspaceToolboxAPI.md#gitcommitchanges) | **Post** /workspace/{workspaceId}/toolbox/git/commit | Commit changes
*WorkspaceToolboxAPI* | [**GitCommitHistory**](docs/WorkspaceToolboxAPI.md#gitcommithistory) | **Get** /workspace/{workspaceId}/toolbox/git/history | Get commit history
*WorkspaceToolboxAPI* | [**GitCreateBranch**](docs/WorkspaceToolboxAPI.md#gitcreatebranch) | **Post** /workspace/{workspaceId}/toolbox/git/branches | Create branch
*WorkspaceToolboxAPI* | [**GitGitStatus**](docs/WorkspaceToolboxAPI.md#gitgitstatus) | **Get** /workspace/{workspaceId}/toolbox/git/status | Get git status
*WorkspaceToolboxAPI* | [**GitPullChanges**](docs/WorkspaceToolboxAPI.md#gitpullchanges) | **Post** /workspace/{workspaceId}/toolbox/git/pull | Pull changes
*WorkspaceToolboxAPI* | [**GitPushChanges**](docs/WorkspaceToolboxAPI.md#gitpushchanges) | **Post** /workspace/{workspaceId}/toolbox/git/push | Push changes
*WorkspaceToolboxAPI* | [**ListSessions**](docs/WorkspaceToolboxAPI.md#listsessions) | **Get** /workspace/{workspaceId}/toolbox/process/session | List sessions
*WorkspaceToolboxAPI* | [**LspCompletions**](docs/WorkspaceToolboxAPI.md#lspcompletions) | **Post** /workspace/{workspaceId}/toolbox/lsp/completions | Get Lsp Completions
*WorkspaceToolboxAPI* | [**LspDidClose**](docs/WorkspaceToolboxAPI.md#lspdidclose) | **Post** /workspace/{workspaceId}/toolbox/lsp/did-close | Call Lsp DidClose
*WorkspaceToolboxAPI* | [**LspDidOpen**](docs/WorkspaceToolboxAPI.md#lspdidopen) | **Post** /workspace/{workspaceId}/toolbox/lsp/did-open | Call Lsp DidOpen
*WorkspaceToolboxAPI* | [**LspDocumentSymbols**](docs/WorkspaceToolboxAPI.md#lspdocumentsymbols) | **Get** /workspace/{workspaceId}/toolbox/lsp/document-symbols | Call Lsp DocumentSymbols
*WorkspaceToolboxAPI* | [**LspStart**](docs/WorkspaceToolboxAPI.md#lspstart) | **Post** /workspace/{workspaceId}/toolbox/lsp/start | Start Lsp server
*WorkspaceToolboxAPI* | [**LspStop**](docs/WorkspaceToolboxAPI.md#lspstop) | **Post** /workspace/{workspaceId}/toolbox/lsp/stop | Stop Lsp server
*WorkspaceToolboxAPI* | [**LspWorkspaceSymbols**](docs/WorkspaceToolboxAPI.md#lspworkspacesymbols) | **Get** /workspace/{workspaceId}/toolbox/lsp/workspace-symbols | Call Lsp WorkspaceSymbols
*WorkspaceToolboxAPI* | [**ProcessExecuteCommand**](docs/WorkspaceToolboxAPI.md#processexecutecommand) | **Post** /workspace/{workspaceId}/toolbox/process/execute | Execute command
*WorkspaceToolboxAPI* | [**SessionExecuteCommand**](docs/WorkspaceToolboxAPI.md#sessionexecutecommand) | **Post** /workspace/{workspaceId}/toolbox/process/session/{sessionId}/exec | Execute command in session


## Documentation For Models

 - [ApiKeyViewDTO](docs/ApiKeyViewDTO.md)
 - [BuildConfig](docs/BuildConfig.md)
 - [BuildDTO](docs/BuildDTO.md)
 - [CachedBuild](docs/CachedBuild.md)
 - [CloneTarget](docs/CloneTarget.md)
 - [Command](docs/Command.md)
 - [CompletionContext](docs/CompletionContext.md)
 - [CompletionItem](docs/CompletionItem.md)
 - [CompletionList](docs/CompletionList.md)
 - [ContainerConfig](docs/ContainerConfig.md)
 - [ContainerRegistry](docs/ContainerRegistry.md)
 - [CreateBuildDTO](docs/CreateBuildDTO.md)
 - [CreatePrebuildDTO](docs/CreatePrebuildDTO.md)
 - [CreateRunnerDTO](docs/CreateRunnerDTO.md)
 - [CreateRunnerResultDTO](docs/CreateRunnerResultDTO.md)
 - [CreateSessionRequest](docs/CreateSessionRequest.md)
 - [CreateTargetConfigDTO](docs/CreateTargetConfigDTO.md)
 - [CreateTargetDTO](docs/CreateTargetDTO.md)
 - [CreateWorkspaceDTO](docs/CreateWorkspaceDTO.md)
 - [CreateWorkspaceSourceDTO](docs/CreateWorkspaceSourceDTO.md)
 - [CreateWorkspaceTemplateDTO](docs/CreateWorkspaceTemplateDTO.md)
 - [DevcontainerConfig](docs/DevcontainerConfig.md)
 - [EnvironmentVariable](docs/EnvironmentVariable.md)
 - [ExecuteRequest](docs/ExecuteRequest.md)
 - [ExecuteResponse](docs/ExecuteResponse.md)
 - [FRPSConfig](docs/FRPSConfig.md)
 - [FileInfo](docs/FileInfo.md)
 - [FileStatus](docs/FileStatus.md)
 - [GetRepositoryContext](docs/GetRepositoryContext.md)
 - [GitAddRequest](docs/GitAddRequest.md)
 - [GitBranch](docs/GitBranch.md)
 - [GitBranchRequest](docs/GitBranchRequest.md)
 - [GitCloneRequest](docs/GitCloneRequest.md)
 - [GitCommitInfo](docs/GitCommitInfo.md)
 - [GitCommitRequest](docs/GitCommitRequest.md)
 - [GitCommitResponse](docs/GitCommitResponse.md)
 - [GitNamespace](docs/GitNamespace.md)
 - [GitProvider](docs/GitProvider.md)
 - [GitPullRequest](docs/GitPullRequest.md)
 - [GitRepoRequest](docs/GitRepoRequest.md)
 - [GitRepository](docs/GitRepository.md)
 - [GitStatus](docs/GitStatus.md)
 - [GitUser](docs/GitUser.md)
 - [Job](docs/Job.md)
 - [JobState](docs/JobState.md)
 - [ListBranchResponse](docs/ListBranchResponse.md)
 - [LogFileConfig](docs/LogFileConfig.md)
 - [LspCompletionParams](docs/LspCompletionParams.md)
 - [LspDocumentRequest](docs/LspDocumentRequest.md)
 - [LspLocation](docs/LspLocation.md)
 - [LspPosition](docs/LspPosition.md)
 - [LspRange](docs/LspRange.md)
 - [LspServerRequest](docs/LspServerRequest.md)
 - [LspSymbol](docs/LspSymbol.md)
 - [Match](docs/Match.md)
 - [ModelsApiKeyType](docs/ModelsApiKeyType.md)
 - [ModelsJobAction](docs/ModelsJobAction.md)
 - [ModelsResourceStateName](docs/ModelsResourceStateName.md)
 - [ModelsTargetConfigPropertyType](docs/ModelsTargetConfigPropertyType.md)
 - [NetworkKey](docs/NetworkKey.md)
 - [Position](docs/Position.md)
 - [PrebuildConfig](docs/PrebuildConfig.md)
 - [PrebuildDTO](docs/PrebuildDTO.md)
 - [ProviderDTO](docs/ProviderDTO.md)
 - [ProviderInfo](docs/ProviderInfo.md)
 - [ReplaceRequest](docs/ReplaceRequest.md)
 - [ReplaceResult](docs/ReplaceResult.md)
 - [RepositoryUrl](docs/RepositoryUrl.md)
 - [ResourceState](docs/ResourceState.md)
 - [ResourceType](docs/ResourceType.md)
 - [RunnerDTO](docs/RunnerDTO.md)
 - [RunnerMetadata](docs/RunnerMetadata.md)
 - [Sample](docs/Sample.md)
 - [SearchFilesResponse](docs/SearchFilesResponse.md)
 - [ServerConfig](docs/ServerConfig.md)
 - [Session](docs/Session.md)
 - [SessionExecuteRequest](docs/SessionExecuteRequest.md)
 - [SessionExecuteResponse](docs/SessionExecuteResponse.md)
 - [SetGitProviderConfig](docs/SetGitProviderConfig.md)
 - [SigningMethod](docs/SigningMethod.md)
 - [Status](docs/Status.md)
 - [Target](docs/Target.md)
 - [TargetConfig](docs/TargetConfig.md)
 - [TargetConfigProperty](docs/TargetConfigProperty.md)
 - [TargetDTO](docs/TargetDTO.md)
 - [TargetMetadata](docs/TargetMetadata.md)
 - [UpdateJobState](docs/UpdateJobState.md)
 - [UpdateRunnerMetadataDTO](docs/UpdateRunnerMetadataDTO.md)
 - [UpdateTargetMetadataDTO](docs/UpdateTargetMetadataDTO.md)
 - [UpdateTargetProviderMetadataDTO](docs/UpdateTargetProviderMetadataDTO.md)
 - [UpdateWorkspaceMetadataDTO](docs/UpdateWorkspaceMetadataDTO.md)
 - [UpdateWorkspaceProviderMetadataDTO](docs/UpdateWorkspaceProviderMetadataDTO.md)
 - [Workspace](docs/Workspace.md)
 - [WorkspaceDTO](docs/WorkspaceDTO.md)
 - [WorkspaceDirResponse](docs/WorkspaceDirResponse.md)
 - [WorkspaceMetadata](docs/WorkspaceMetadata.md)
 - [WorkspaceTemplate](docs/WorkspaceTemplate.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### Bearer

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		apiclient.ContextAPIKeys,
		map[string]apiclient.APIKey{
			"Authorization": {Key: "API_KEY_STRING"},
		},
	)
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



