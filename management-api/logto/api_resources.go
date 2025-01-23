/*
Logto API references

API references for Logto services.  Note: The documentation is for Logto Cloud. If you are using Logto OSS, please refer to the response of `/api/swagger.json` endpoint on your Logto instance.

API version: Cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package logto

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// ResourcesAPIService ResourcesAPI service
type ResourcesAPIService service

type ApiCreateResourceRequest struct {
	ctx context.Context
	ApiService *ResourcesAPIService
	createResourceRequest *CreateResourceRequest
}

func (r ApiCreateResourceRequest) CreateResourceRequest(createResourceRequest CreateResourceRequest) ApiCreateResourceRequest {
	r.createResourceRequest = &createResourceRequest
	return r
}

func (r ApiCreateResourceRequest) Execute() (*ListResources200ResponseInner, *http.Response, error) {
	return r.ApiService.CreateResourceExecute(r)
}

/*
CreateResource Create an API resource

Create an API resource in the current tenant.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateResourceRequest
*/
func (a *ResourcesAPIService) CreateResource(ctx context.Context) ApiCreateResourceRequest {
	return ApiCreateResourceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListResources200ResponseInner
func (a *ResourcesAPIService) CreateResourceExecute(r ApiCreateResourceRequest) (*ListResources200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListResources200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourcesAPIService.CreateResource")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/resources"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createResourceRequest == nil {
		return localVarReturnValue, nil, reportError("createResourceRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	localVarPostBody = r.createResourceRequest
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

type ApiCreateResourceScopeRequest struct {
	ctx context.Context
	ApiService *ResourcesAPIService
	resourceId string
	createResourceScopeRequest *CreateResourceScopeRequest
}

func (r ApiCreateResourceScopeRequest) CreateResourceScopeRequest(createResourceScopeRequest CreateResourceScopeRequest) ApiCreateResourceScopeRequest {
	r.createResourceScopeRequest = &createResourceScopeRequest
	return r
}

func (r ApiCreateResourceScopeRequest) Execute() (*ListResources200ResponseInnerScopesInner, *http.Response, error) {
	return r.ApiService.CreateResourceScopeExecute(r)
}

/*
CreateResourceScope Create API resource scope

Create a new scope (permission) for an API resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param resourceId The unique identifier of the resource.
 @return ApiCreateResourceScopeRequest
*/
func (a *ResourcesAPIService) CreateResourceScope(ctx context.Context, resourceId string) ApiCreateResourceScopeRequest {
	return ApiCreateResourceScopeRequest{
		ApiService: a,
		ctx: ctx,
		resourceId: resourceId,
	}
}

// Execute executes the request
//  @return ListResources200ResponseInnerScopesInner
func (a *ResourcesAPIService) CreateResourceScopeExecute(r ApiCreateResourceScopeRequest) (*ListResources200ResponseInnerScopesInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListResources200ResponseInnerScopesInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourcesAPIService.CreateResourceScope")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/resources/{resourceId}/scopes"
	localVarPath = strings.Replace(localVarPath, "{"+"resourceId"+"}", url.PathEscape(parameterValueToString(r.resourceId, "resourceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createResourceScopeRequest == nil {
		return localVarReturnValue, nil, reportError("createResourceScopeRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	localVarPostBody = r.createResourceScopeRequest
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

type ApiDeleteResourceRequest struct {
	ctx context.Context
	ApiService *ResourcesAPIService
	id string
}

func (r ApiDeleteResourceRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteResourceExecute(r)
}

/*
DeleteResource Delete API resource

Delete an API resource by ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The unique identifier of the resource.
 @return ApiDeleteResourceRequest
*/
func (a *ResourcesAPIService) DeleteResource(ctx context.Context, id string) ApiDeleteResourceRequest {
	return ApiDeleteResourceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ResourcesAPIService) DeleteResourceExecute(r ApiDeleteResourceRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourcesAPIService.DeleteResource")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/resources/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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

type ApiDeleteResourceScopeRequest struct {
	ctx context.Context
	ApiService *ResourcesAPIService
	resourceId string
	scopeId string
}

func (r ApiDeleteResourceScopeRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteResourceScopeExecute(r)
}

/*
DeleteResourceScope Delete API resource scope

Delete an API resource scope (permission) from the given resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param resourceId The unique identifier of the resource.
 @param scopeId The unique identifier of the scope.
 @return ApiDeleteResourceScopeRequest
*/
func (a *ResourcesAPIService) DeleteResourceScope(ctx context.Context, resourceId string, scopeId string) ApiDeleteResourceScopeRequest {
	return ApiDeleteResourceScopeRequest{
		ApiService: a,
		ctx: ctx,
		resourceId: resourceId,
		scopeId: scopeId,
	}
}

// Execute executes the request
func (a *ResourcesAPIService) DeleteResourceScopeExecute(r ApiDeleteResourceScopeRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourcesAPIService.DeleteResourceScope")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/resources/{resourceId}/scopes/{scopeId}"
	localVarPath = strings.Replace(localVarPath, "{"+"resourceId"+"}", url.PathEscape(parameterValueToString(r.resourceId, "resourceId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"scopeId"+"}", url.PathEscape(parameterValueToString(r.scopeId, "scopeId")), -1)

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

type ApiGetResourceRequest struct {
	ctx context.Context
	ApiService *ResourcesAPIService
	id string
}

func (r ApiGetResourceRequest) Execute() (*GetJwtCustomizer200ResponseOneOfContextSampleUserRolesInnerScopesInnerResource, *http.Response, error) {
	return r.ApiService.GetResourceExecute(r)
}

/*
GetResource Get API resource

Get an API resource details by ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The unique identifier of the resource.
 @return ApiGetResourceRequest
*/
func (a *ResourcesAPIService) GetResource(ctx context.Context, id string) ApiGetResourceRequest {
	return ApiGetResourceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return GetJwtCustomizer200ResponseOneOfContextSampleUserRolesInnerScopesInnerResource
func (a *ResourcesAPIService) GetResourceExecute(r ApiGetResourceRequest) (*GetJwtCustomizer200ResponseOneOfContextSampleUserRolesInnerScopesInnerResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetJwtCustomizer200ResponseOneOfContextSampleUserRolesInnerScopesInnerResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourcesAPIService.GetResource")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/resources/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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

type ApiListResourceScopesRequest struct {
	ctx context.Context
	ApiService *ResourcesAPIService
	resourceId string
	page *int32
	pageSize *int32
	searchParams *map[string]string
}

// Page number (starts from 1).
func (r ApiListResourceScopesRequest) Page(page int32) ApiListResourceScopesRequest {
	r.page = &page
	return r
}

// Entries per page.
func (r ApiListResourceScopesRequest) PageSize(pageSize int32) ApiListResourceScopesRequest {
	r.pageSize = &pageSize
	return r
}

// Search query parameters.
func (r ApiListResourceScopesRequest) SearchParams(searchParams map[string]string) ApiListResourceScopesRequest {
	r.searchParams = &searchParams
	return r
}

func (r ApiListResourceScopesRequest) Execute() ([]ListResources200ResponseInnerScopesInner, *http.Response, error) {
	return r.ApiService.ListResourceScopesExecute(r)
}

/*
ListResourceScopes Get API resource scopes

Get scopes (permissions) defined for an API resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param resourceId The unique identifier of the resource.
 @return ApiListResourceScopesRequest
*/
func (a *ResourcesAPIService) ListResourceScopes(ctx context.Context, resourceId string) ApiListResourceScopesRequest {
	return ApiListResourceScopesRequest{
		ApiService: a,
		ctx: ctx,
		resourceId: resourceId,
	}
}

// Execute executes the request
//  @return []ListResources200ResponseInnerScopesInner
func (a *ResourcesAPIService) ListResourceScopesExecute(r ApiListResourceScopesRequest) ([]ListResources200ResponseInnerScopesInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ListResources200ResponseInnerScopesInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourcesAPIService.ListResourceScopes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/resources/{resourceId}/scopes"
	localVarPath = strings.Replace(localVarPath, "{"+"resourceId"+"}", url.PathEscape(parameterValueToString(r.resourceId, "resourceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	} else {
		var defaultValue int32 = 1
		r.page = &defaultValue
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize, "form", "")
	} else {
		var defaultValue int32 = 20
		r.pageSize = &defaultValue
	}
	if r.searchParams != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search_params", r.searchParams, "form", "")
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

type ApiListResourcesRequest struct {
	ctx context.Context
	ApiService *ResourcesAPIService
	includeScopes *string
	page *int32
	pageSize *int32
}

// If it&#39;s provided with a truthy value (&#x60;true&#x60;, &#x60;1&#x60;, &#x60;yes&#x60;), the scopes of each resource will be included in the response.
func (r ApiListResourcesRequest) IncludeScopes(includeScopes string) ApiListResourcesRequest {
	r.includeScopes = &includeScopes
	return r
}

// Page number (starts from 1).
func (r ApiListResourcesRequest) Page(page int32) ApiListResourcesRequest {
	r.page = &page
	return r
}

// Entries per page.
func (r ApiListResourcesRequest) PageSize(pageSize int32) ApiListResourcesRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiListResourcesRequest) Execute() ([]ListResources200ResponseInner, *http.Response, error) {
	return r.ApiService.ListResourcesExecute(r)
}

/*
ListResources Get API resources

Get API resources in the current tenant with pagination.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListResourcesRequest
*/
func (a *ResourcesAPIService) ListResources(ctx context.Context) ApiListResourcesRequest {
	return ApiListResourcesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ListResources200ResponseInner
func (a *ResourcesAPIService) ListResourcesExecute(r ApiListResourcesRequest) ([]ListResources200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ListResources200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourcesAPIService.ListResources")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/resources"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.includeScopes != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeScopes", r.includeScopes, "form", "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	} else {
		var defaultValue int32 = 1
		r.page = &defaultValue
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize, "form", "")
	} else {
		var defaultValue int32 = 20
		r.pageSize = &defaultValue
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

type ApiUpdateResourceRequest struct {
	ctx context.Context
	ApiService *ResourcesAPIService
	id string
	updateResourceRequest *UpdateResourceRequest
}

func (r ApiUpdateResourceRequest) UpdateResourceRequest(updateResourceRequest UpdateResourceRequest) ApiUpdateResourceRequest {
	r.updateResourceRequest = &updateResourceRequest
	return r
}

func (r ApiUpdateResourceRequest) Execute() (*GetJwtCustomizer200ResponseOneOfContextSampleUserRolesInnerScopesInnerResource, *http.Response, error) {
	return r.ApiService.UpdateResourceExecute(r)
}

/*
UpdateResource Update API resource

Update an API resource details by ID with the given data. This method performs a partial update.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The unique identifier of the resource.
 @return ApiUpdateResourceRequest
*/
func (a *ResourcesAPIService) UpdateResource(ctx context.Context, id string) ApiUpdateResourceRequest {
	return ApiUpdateResourceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return GetJwtCustomizer200ResponseOneOfContextSampleUserRolesInnerScopesInnerResource
func (a *ResourcesAPIService) UpdateResourceExecute(r ApiUpdateResourceRequest) (*GetJwtCustomizer200ResponseOneOfContextSampleUserRolesInnerScopesInnerResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetJwtCustomizer200ResponseOneOfContextSampleUserRolesInnerScopesInnerResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourcesAPIService.UpdateResource")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/resources/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateResourceRequest == nil {
		return localVarReturnValue, nil, reportError("updateResourceRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	localVarPostBody = r.updateResourceRequest
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

type ApiUpdateResourceIsDefaultRequest struct {
	ctx context.Context
	ApiService *ResourcesAPIService
	id string
	updateResourceIsDefaultRequest *UpdateResourceIsDefaultRequest
}

func (r ApiUpdateResourceIsDefaultRequest) UpdateResourceIsDefaultRequest(updateResourceIsDefaultRequest UpdateResourceIsDefaultRequest) ApiUpdateResourceIsDefaultRequest {
	r.updateResourceIsDefaultRequest = &updateResourceIsDefaultRequest
	return r
}

func (r ApiUpdateResourceIsDefaultRequest) Execute() (*GetJwtCustomizer200ResponseOneOfContextSampleUserRolesInnerScopesInnerResource, *http.Response, error) {
	return r.ApiService.UpdateResourceIsDefaultExecute(r)
}

/*
UpdateResourceIsDefault Set API resource as default

Set an API resource as the default resource for the current tenant.

Each tenant can have only one default API resource. If an API resource is set as default, the previously set default API resource will be set as non-default. See [this section](https://docs.logto.io/docs/references/resources/#default-api) for more information.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The unique identifier of the resource.
 @return ApiUpdateResourceIsDefaultRequest
*/
func (a *ResourcesAPIService) UpdateResourceIsDefault(ctx context.Context, id string) ApiUpdateResourceIsDefaultRequest {
	return ApiUpdateResourceIsDefaultRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return GetJwtCustomizer200ResponseOneOfContextSampleUserRolesInnerScopesInnerResource
func (a *ResourcesAPIService) UpdateResourceIsDefaultExecute(r ApiUpdateResourceIsDefaultRequest) (*GetJwtCustomizer200ResponseOneOfContextSampleUserRolesInnerScopesInnerResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetJwtCustomizer200ResponseOneOfContextSampleUserRolesInnerScopesInnerResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourcesAPIService.UpdateResourceIsDefault")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/resources/{id}/is-default"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateResourceIsDefaultRequest == nil {
		return localVarReturnValue, nil, reportError("updateResourceIsDefaultRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	localVarPostBody = r.updateResourceIsDefaultRequest
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

type ApiUpdateResourceScopeRequest struct {
	ctx context.Context
	ApiService *ResourcesAPIService
	resourceId string
	scopeId string
	updateResourceScopeRequest *UpdateResourceScopeRequest
}

func (r ApiUpdateResourceScopeRequest) UpdateResourceScopeRequest(updateResourceScopeRequest UpdateResourceScopeRequest) ApiUpdateResourceScopeRequest {
	r.updateResourceScopeRequest = &updateResourceScopeRequest
	return r
}

func (r ApiUpdateResourceScopeRequest) Execute() (*ListResources200ResponseInnerScopesInner, *http.Response, error) {
	return r.ApiService.UpdateResourceScopeExecute(r)
}

/*
UpdateResourceScope Update API resource scope

Update an API resource scope (permission) for the given resource. This method performs a partial update.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param resourceId The unique identifier of the resource.
 @param scopeId The unique identifier of the scope.
 @return ApiUpdateResourceScopeRequest
*/
func (a *ResourcesAPIService) UpdateResourceScope(ctx context.Context, resourceId string, scopeId string) ApiUpdateResourceScopeRequest {
	return ApiUpdateResourceScopeRequest{
		ApiService: a,
		ctx: ctx,
		resourceId: resourceId,
		scopeId: scopeId,
	}
}

// Execute executes the request
//  @return ListResources200ResponseInnerScopesInner
func (a *ResourcesAPIService) UpdateResourceScopeExecute(r ApiUpdateResourceScopeRequest) (*ListResources200ResponseInnerScopesInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListResources200ResponseInnerScopesInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourcesAPIService.UpdateResourceScope")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/resources/{resourceId}/scopes/{scopeId}"
	localVarPath = strings.Replace(localVarPath, "{"+"resourceId"+"}", url.PathEscape(parameterValueToString(r.resourceId, "resourceId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"scopeId"+"}", url.PathEscape(parameterValueToString(r.scopeId, "scopeId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateResourceScopeRequest == nil {
		return localVarReturnValue, nil, reportError("updateResourceScopeRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	localVarPostBody = r.updateResourceScopeRequest
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
