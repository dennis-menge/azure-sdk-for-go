//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armoperationalinsights

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// LinkedStorageAccountsClient contains the methods for the LinkedStorageAccounts group.
// Don't use this type directly, use NewLinkedStorageAccountsClient() instead.
type LinkedStorageAccountsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewLinkedStorageAccountsClient creates a new instance of LinkedStorageAccountsClient with the specified values.
func NewLinkedStorageAccountsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *LinkedStorageAccountsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &LinkedStorageAccountsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CreateOrUpdate - Create or Update a link relation between current workspace and a group of storage accounts of a specific data source type.
// If the operation fails it returns a generic error.
func (client *LinkedStorageAccountsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, dataSourceType DataSourceType, parameters LinkedStorageAccountsResource, options *LinkedStorageAccountsCreateOrUpdateOptions) (LinkedStorageAccountsCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, dataSourceType, parameters, options)
	if err != nil {
		return LinkedStorageAccountsCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LinkedStorageAccountsCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LinkedStorageAccountsCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *LinkedStorageAccountsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, dataSourceType DataSourceType, parameters LinkedStorageAccountsResource, options *LinkedStorageAccountsCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/linkedStorageAccounts/{dataSourceType}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if dataSourceType == "" {
		return nil, errors.New("parameter dataSourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataSourceType}", url.PathEscape(string(dataSourceType)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *LinkedStorageAccountsClient) createOrUpdateHandleResponse(resp *http.Response) (LinkedStorageAccountsCreateOrUpdateResponse, error) {
	result := LinkedStorageAccountsCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.LinkedStorageAccountsResource); err != nil {
		return LinkedStorageAccountsCreateOrUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *LinkedStorageAccountsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Delete - Deletes all linked storage accounts of a specific data source type associated with the specified workspace.
// If the operation fails it returns a generic error.
func (client *LinkedStorageAccountsClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, dataSourceType DataSourceType, options *LinkedStorageAccountsDeleteOptions) (LinkedStorageAccountsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, dataSourceType, options)
	if err != nil {
		return LinkedStorageAccountsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LinkedStorageAccountsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LinkedStorageAccountsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return LinkedStorageAccountsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *LinkedStorageAccountsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, dataSourceType DataSourceType, options *LinkedStorageAccountsDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/linkedStorageAccounts/{dataSourceType}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if dataSourceType == "" {
		return nil, errors.New("parameter dataSourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataSourceType}", url.PathEscape(string(dataSourceType)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *LinkedStorageAccountsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Gets all linked storage account of a specific data source type associated with the specified workspace.
// If the operation fails it returns a generic error.
func (client *LinkedStorageAccountsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, dataSourceType DataSourceType, options *LinkedStorageAccountsGetOptions) (LinkedStorageAccountsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, dataSourceType, options)
	if err != nil {
		return LinkedStorageAccountsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LinkedStorageAccountsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LinkedStorageAccountsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LinkedStorageAccountsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, dataSourceType DataSourceType, options *LinkedStorageAccountsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/linkedStorageAccounts/{dataSourceType}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if dataSourceType == "" {
		return nil, errors.New("parameter dataSourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataSourceType}", url.PathEscape(string(dataSourceType)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LinkedStorageAccountsClient) getHandleResponse(resp *http.Response) (LinkedStorageAccountsGetResponse, error) {
	result := LinkedStorageAccountsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.LinkedStorageAccountsResource); err != nil {
		return LinkedStorageAccountsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *LinkedStorageAccountsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByWorkspace - Gets all linked storage accounts associated with the specified workspace, storage accounts will be sorted by their data source type.
// If the operation fails it returns a generic error.
func (client *LinkedStorageAccountsClient) ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string, options *LinkedStorageAccountsListByWorkspaceOptions) (LinkedStorageAccountsListByWorkspaceResponse, error) {
	req, err := client.listByWorkspaceCreateRequest(ctx, resourceGroupName, workspaceName, options)
	if err != nil {
		return LinkedStorageAccountsListByWorkspaceResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LinkedStorageAccountsListByWorkspaceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LinkedStorageAccountsListByWorkspaceResponse{}, client.listByWorkspaceHandleError(resp)
	}
	return client.listByWorkspaceHandleResponse(resp)
}

// listByWorkspaceCreateRequest creates the ListByWorkspace request.
func (client *LinkedStorageAccountsClient) listByWorkspaceCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *LinkedStorageAccountsListByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/linkedStorageAccounts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByWorkspaceHandleResponse handles the ListByWorkspace response.
func (client *LinkedStorageAccountsClient) listByWorkspaceHandleResponse(resp *http.Response) (LinkedStorageAccountsListByWorkspaceResponse, error) {
	result := LinkedStorageAccountsListByWorkspaceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.LinkedStorageAccountsListResult); err != nil {
		return LinkedStorageAccountsListByWorkspaceResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByWorkspaceHandleError handles the ListByWorkspace error response.
func (client *LinkedStorageAccountsClient) listByWorkspaceHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
