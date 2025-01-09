/*
Daytona Server API

Daytona Server API

API version: v0.0.0-dev
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// WorkspaceAPIService WorkspaceAPI service
type WorkspaceAPIService service

type ApiCreateWorkspaceRequest struct {
	ctx        context.Context
	ApiService *WorkspaceAPIService
	workspace  *CreateWorkspaceDTO
}

// Create workspace
func (r ApiCreateWorkspaceRequest) Workspace(workspace CreateWorkspaceDTO) ApiCreateWorkspaceRequest {
	r.workspace = &workspace
	return r
}

func (r ApiCreateWorkspaceRequest) Execute() (*WorkspaceDTO, *http.Response, error) {
	return r.ApiService.CreateWorkspaceExecute(r)
}

/*
CreateWorkspace Create a workspace

Create a workspace

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateWorkspaceRequest
*/
func (a *WorkspaceAPIService) CreateWorkspace(ctx context.Context) ApiCreateWorkspaceRequest {
	return ApiCreateWorkspaceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return WorkspaceDTO
func (a *WorkspaceAPIService) CreateWorkspaceExecute(r ApiCreateWorkspaceRequest) (*WorkspaceDTO, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *WorkspaceDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkspaceAPIService.CreateWorkspace")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workspace"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.workspace == nil {
		return localVarReturnValue, nil, reportError("workspace is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.workspace
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Bearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteWorkspaceRequest struct {
	ctx         context.Context
	ApiService  *WorkspaceAPIService
	workspaceId string
	force       *bool
}

// Force
func (r ApiDeleteWorkspaceRequest) Force(force bool) ApiDeleteWorkspaceRequest {
	r.force = &force
	return r
}

func (r ApiDeleteWorkspaceRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteWorkspaceExecute(r)
}

/*
DeleteWorkspace Delete workspace

Delete workspace

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param workspaceId Workspace ID
	@return ApiDeleteWorkspaceRequest
*/
func (a *WorkspaceAPIService) DeleteWorkspace(ctx context.Context, workspaceId string) ApiDeleteWorkspaceRequest {
	return ApiDeleteWorkspaceRequest{
		ApiService:  a,
		ctx:         ctx,
		workspaceId: workspaceId,
	}
}

// Execute executes the request
func (a *WorkspaceAPIService) DeleteWorkspaceExecute(r ApiDeleteWorkspaceRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkspaceAPIService.DeleteWorkspace")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workspace/{workspaceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"workspaceId"+"}", url.PathEscape(parameterValueToString(r.workspaceId, "workspaceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.force != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "force", r.force, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Bearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiFindWorkspaceRequest struct {
	ctx         context.Context
	ApiService  *WorkspaceAPIService
	workspaceId string
}

func (r ApiFindWorkspaceRequest) Execute() (*WorkspaceDTO, *http.Response, error) {
	return r.ApiService.FindWorkspaceExecute(r)
}

/*
FindWorkspace Find workspace

Find workspace

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param workspaceId Workspace ID or Name
	@return ApiFindWorkspaceRequest
*/
func (a *WorkspaceAPIService) FindWorkspace(ctx context.Context, workspaceId string) ApiFindWorkspaceRequest {
	return ApiFindWorkspaceRequest{
		ApiService:  a,
		ctx:         ctx,
		workspaceId: workspaceId,
	}
}

// Execute executes the request
//
//	@return WorkspaceDTO
func (a *WorkspaceAPIService) FindWorkspaceExecute(r ApiFindWorkspaceRequest) (*WorkspaceDTO, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *WorkspaceDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkspaceAPIService.FindWorkspace")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workspace/{workspaceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"workspaceId"+"}", url.PathEscape(parameterValueToString(r.workspaceId, "workspaceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Bearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListWorkspacesRequest struct {
	ctx        context.Context
	ApiService *WorkspaceAPIService
}

func (r ApiListWorkspacesRequest) Execute() ([]WorkspaceDTO, *http.Response, error) {
	return r.ApiService.ListWorkspacesExecute(r)
}

/*
ListWorkspaces List workspaces

List workspaces

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListWorkspacesRequest
*/
func (a *WorkspaceAPIService) ListWorkspaces(ctx context.Context) ApiListWorkspacesRequest {
	return ApiListWorkspacesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []WorkspaceDTO
func (a *WorkspaceAPIService) ListWorkspacesExecute(r ApiListWorkspacesRequest) ([]WorkspaceDTO, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []WorkspaceDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkspaceAPIService.ListWorkspaces")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workspace"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Bearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiRestartWorkspaceRequest struct {
	ctx         context.Context
	ApiService  *WorkspaceAPIService
	workspaceId string
}

func (r ApiRestartWorkspaceRequest) Execute() (*http.Response, error) {
	return r.ApiService.RestartWorkspaceExecute(r)
}

/*
RestartWorkspace Restart workspace

Restart workspace

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param workspaceId Workspace ID or Name
	@return ApiRestartWorkspaceRequest
*/
func (a *WorkspaceAPIService) RestartWorkspace(ctx context.Context, workspaceId string) ApiRestartWorkspaceRequest {
	return ApiRestartWorkspaceRequest{
		ApiService:  a,
		ctx:         ctx,
		workspaceId: workspaceId,
	}
}

// Execute executes the request
func (a *WorkspaceAPIService) RestartWorkspaceExecute(r ApiRestartWorkspaceRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkspaceAPIService.RestartWorkspace")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workspace/{workspaceId}/restart"
	localVarPath = strings.Replace(localVarPath, "{"+"workspaceId"+"}", url.PathEscape(parameterValueToString(r.workspaceId, "workspaceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Bearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiStartWorkspaceRequest struct {
	ctx         context.Context
	ApiService  *WorkspaceAPIService
	workspaceId string
}

func (r ApiStartWorkspaceRequest) Execute() (*http.Response, error) {
	return r.ApiService.StartWorkspaceExecute(r)
}

/*
StartWorkspace Start workspace

Start workspace

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param workspaceId Workspace ID or Name
	@return ApiStartWorkspaceRequest
*/
func (a *WorkspaceAPIService) StartWorkspace(ctx context.Context, workspaceId string) ApiStartWorkspaceRequest {
	return ApiStartWorkspaceRequest{
		ApiService:  a,
		ctx:         ctx,
		workspaceId: workspaceId,
	}
}

// Execute executes the request
func (a *WorkspaceAPIService) StartWorkspaceExecute(r ApiStartWorkspaceRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkspaceAPIService.StartWorkspace")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workspace/{workspaceId}/start"
	localVarPath = strings.Replace(localVarPath, "{"+"workspaceId"+"}", url.PathEscape(parameterValueToString(r.workspaceId, "workspaceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Bearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiStopWorkspaceRequest struct {
	ctx         context.Context
	ApiService  *WorkspaceAPIService
	workspaceId string
}

func (r ApiStopWorkspaceRequest) Execute() (*http.Response, error) {
	return r.ApiService.StopWorkspaceExecute(r)
}

/*
StopWorkspace Stop workspace

Stop workspace

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param workspaceId Workspace ID or Name
	@return ApiStopWorkspaceRequest
*/
func (a *WorkspaceAPIService) StopWorkspace(ctx context.Context, workspaceId string) ApiStopWorkspaceRequest {
	return ApiStopWorkspaceRequest{
		ApiService:  a,
		ctx:         ctx,
		workspaceId: workspaceId,
	}
}

// Execute executes the request
func (a *WorkspaceAPIService) StopWorkspaceExecute(r ApiStopWorkspaceRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkspaceAPIService.StopWorkspace")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workspace/{workspaceId}/stop"
	localVarPath = strings.Replace(localVarPath, "{"+"workspaceId"+"}", url.PathEscape(parameterValueToString(r.workspaceId, "workspaceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Bearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiUpdateWorkspaceMetadataRequest struct {
	ctx               context.Context
	ApiService        *WorkspaceAPIService
	workspaceId       string
	workspaceMetadata *UpdateWorkspaceMetadataDTO
}

// Workspace Metadata
func (r ApiUpdateWorkspaceMetadataRequest) WorkspaceMetadata(workspaceMetadata UpdateWorkspaceMetadataDTO) ApiUpdateWorkspaceMetadataRequest {
	r.workspaceMetadata = &workspaceMetadata
	return r
}

func (r ApiUpdateWorkspaceMetadataRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateWorkspaceMetadataExecute(r)
}

/*
UpdateWorkspaceMetadata Update workspace metadata

Update workspace metadata

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param workspaceId Workspace ID
	@return ApiUpdateWorkspaceMetadataRequest
*/
func (a *WorkspaceAPIService) UpdateWorkspaceMetadata(ctx context.Context, workspaceId string) ApiUpdateWorkspaceMetadataRequest {
	return ApiUpdateWorkspaceMetadataRequest{
		ApiService:  a,
		ctx:         ctx,
		workspaceId: workspaceId,
	}
}

// Execute executes the request
func (a *WorkspaceAPIService) UpdateWorkspaceMetadataExecute(r ApiUpdateWorkspaceMetadataRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkspaceAPIService.UpdateWorkspaceMetadata")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workspace/{workspaceId}/metadata"
	localVarPath = strings.Replace(localVarPath, "{"+"workspaceId"+"}", url.PathEscape(parameterValueToString(r.workspaceId, "workspaceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.workspaceMetadata == nil {
		return nil, reportError("workspaceMetadata is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.workspaceMetadata
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Bearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiUpdateWorkspaceProviderMetadataRequest struct {
	ctx         context.Context
	ApiService  *WorkspaceAPIService
	workspaceId string
	metadata    *UpdateWorkspaceProviderMetadataDTO
}

// Provider metadata
func (r ApiUpdateWorkspaceProviderMetadataRequest) Metadata(metadata UpdateWorkspaceProviderMetadataDTO) ApiUpdateWorkspaceProviderMetadataRequest {
	r.metadata = &metadata
	return r
}

func (r ApiUpdateWorkspaceProviderMetadataRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateWorkspaceProviderMetadataExecute(r)
}

/*
UpdateWorkspaceProviderMetadata Update workspace provider metadata

Update workspace provider metadata

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param workspaceId Workspace ID
	@return ApiUpdateWorkspaceProviderMetadataRequest
*/
func (a *WorkspaceAPIService) UpdateWorkspaceProviderMetadata(ctx context.Context, workspaceId string) ApiUpdateWorkspaceProviderMetadataRequest {
	return ApiUpdateWorkspaceProviderMetadataRequest{
		ApiService:  a,
		ctx:         ctx,
		workspaceId: workspaceId,
	}
}

// Execute executes the request
func (a *WorkspaceAPIService) UpdateWorkspaceProviderMetadataExecute(r ApiUpdateWorkspaceProviderMetadataRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkspaceAPIService.UpdateWorkspaceProviderMetadata")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workspace/{workspaceId}/provider-metadata"
	localVarPath = strings.Replace(localVarPath, "{"+"workspaceId"+"}", url.PathEscape(parameterValueToString(r.workspaceId, "workspaceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.metadata == nil {
		return nil, reportError("metadata is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.metadata
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Bearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
