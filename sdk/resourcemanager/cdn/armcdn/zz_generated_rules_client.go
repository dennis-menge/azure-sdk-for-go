//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcdn

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// RulesClient contains the methods for the Rules group.
// Don't use this type directly, use NewRulesClient() instead.
type RulesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewRulesClient creates a new instance of RulesClient with the specified values.
func NewRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *RulesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &RulesClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCreate - Creates a new delivery rule within the specified rule set.
// If the operation fails it returns the *AfdErrorResponse error type.
func (client *RulesClient) BeginCreate(ctx context.Context, resourceGroupName string, profileName string, ruleSetName string, ruleName string, rule Rule, options *RulesBeginCreateOptions) (RulesCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, profileName, ruleSetName, ruleName, rule, options)
	if err != nil {
		return RulesCreatePollerResponse{}, err
	}
	result := RulesCreatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("RulesClient.Create", "azure-async-operation", resp, client.pl, client.createHandleError)
	if err != nil {
		return RulesCreatePollerResponse{}, err
	}
	result.Poller = &RulesCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Creates a new delivery rule within the specified rule set.
// If the operation fails it returns the *AfdErrorResponse error type.
func (client *RulesClient) create(ctx context.Context, resourceGroupName string, profileName string, ruleSetName string, ruleName string, rule Rule, options *RulesBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, profileName, ruleSetName, ruleName, rule, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.createHandleError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *RulesClient) createCreateRequest(ctx context.Context, resourceGroupName string, profileName string, ruleSetName string, ruleName string, rule Rule, options *RulesBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/ruleSets/{ruleSetName}/rules/{ruleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if ruleSetName == "" {
		return nil, errors.New("parameter ruleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleSetName}", url.PathEscape(ruleSetName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, rule)
}

// createHandleError handles the Create error response.
func (client *RulesClient) createHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := AfdErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Deletes an existing delivery rule within a rule set.
// If the operation fails it returns the *AfdErrorResponse error type.
func (client *RulesClient) BeginDelete(ctx context.Context, resourceGroupName string, profileName string, ruleSetName string, ruleName string, options *RulesBeginDeleteOptions) (RulesDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, profileName, ruleSetName, ruleName, options)
	if err != nil {
		return RulesDeletePollerResponse{}, err
	}
	result := RulesDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("RulesClient.Delete", "azure-async-operation", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return RulesDeletePollerResponse{}, err
	}
	result.Poller = &RulesDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes an existing delivery rule within a rule set.
// If the operation fails it returns the *AfdErrorResponse error type.
func (client *RulesClient) deleteOperation(ctx context.Context, resourceGroupName string, profileName string, ruleSetName string, ruleName string, options *RulesBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, profileName, ruleSetName, ruleName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, profileName string, ruleSetName string, ruleName string, options *RulesBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/ruleSets/{ruleSetName}/rules/{ruleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if ruleSetName == "" {
		return nil, errors.New("parameter ruleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleSetName}", url.PathEscape(ruleSetName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *RulesClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := AfdErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets an existing delivery rule within a rule set.
// If the operation fails it returns the *AfdErrorResponse error type.
func (client *RulesClient) Get(ctx context.Context, resourceGroupName string, profileName string, ruleSetName string, ruleName string, options *RulesGetOptions) (RulesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, profileName, ruleSetName, ruleName, options)
	if err != nil {
		return RulesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RulesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RulesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, profileName string, ruleSetName string, ruleName string, options *RulesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/ruleSets/{ruleSetName}/rules/{ruleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if ruleSetName == "" {
		return nil, errors.New("parameter ruleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleSetName}", url.PathEscape(ruleSetName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RulesClient) getHandleResponse(resp *http.Response) (RulesGetResponse, error) {
	result := RulesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Rule); err != nil {
		return RulesGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *RulesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := AfdErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByRuleSet - Lists all of the existing delivery rules within a rule set.
// If the operation fails it returns the *AfdErrorResponse error type.
func (client *RulesClient) ListByRuleSet(resourceGroupName string, profileName string, ruleSetName string, options *RulesListByRuleSetOptions) *RulesListByRuleSetPager {
	return &RulesListByRuleSetPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByRuleSetCreateRequest(ctx, resourceGroupName, profileName, ruleSetName, options)
		},
		advancer: func(ctx context.Context, resp RulesListByRuleSetResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RuleListResult.NextLink)
		},
	}
}

// listByRuleSetCreateRequest creates the ListByRuleSet request.
func (client *RulesClient) listByRuleSetCreateRequest(ctx context.Context, resourceGroupName string, profileName string, ruleSetName string, options *RulesListByRuleSetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/ruleSets/{ruleSetName}/rules"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if ruleSetName == "" {
		return nil, errors.New("parameter ruleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleSetName}", url.PathEscape(ruleSetName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByRuleSetHandleResponse handles the ListByRuleSet response.
func (client *RulesClient) listByRuleSetHandleResponse(resp *http.Response) (RulesListByRuleSetResponse, error) {
	result := RulesListByRuleSetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RuleListResult); err != nil {
		return RulesListByRuleSetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByRuleSetHandleError handles the ListByRuleSet error response.
func (client *RulesClient) listByRuleSetHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := AfdErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginUpdate - Updates an existing delivery rule within a rule set.
// If the operation fails it returns the *AfdErrorResponse error type.
func (client *RulesClient) BeginUpdate(ctx context.Context, resourceGroupName string, profileName string, ruleSetName string, ruleName string, ruleUpdateProperties RuleUpdateParameters, options *RulesBeginUpdateOptions) (RulesUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, profileName, ruleSetName, ruleName, ruleUpdateProperties, options)
	if err != nil {
		return RulesUpdatePollerResponse{}, err
	}
	result := RulesUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("RulesClient.Update", "azure-async-operation", resp, client.pl, client.updateHandleError)
	if err != nil {
		return RulesUpdatePollerResponse{}, err
	}
	result.Poller = &RulesUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates an existing delivery rule within a rule set.
// If the operation fails it returns the *AfdErrorResponse error type.
func (client *RulesClient) update(ctx context.Context, resourceGroupName string, profileName string, ruleSetName string, ruleName string, ruleUpdateProperties RuleUpdateParameters, options *RulesBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, profileName, ruleSetName, ruleName, ruleUpdateProperties, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *RulesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, profileName string, ruleSetName string, ruleName string, ruleUpdateProperties RuleUpdateParameters, options *RulesBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/ruleSets/{ruleSetName}/rules/{ruleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if ruleSetName == "" {
		return nil, errors.New("parameter ruleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleSetName}", url.PathEscape(ruleSetName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, ruleUpdateProperties)
}

// updateHandleError handles the Update error response.
func (client *RulesClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := AfdErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
