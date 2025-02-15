package migrate

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// MachinesClient is the discover your workloads for Azure.
type MachinesClient struct {
	BaseClient
}

// NewMachinesClient creates an instance of the MachinesClient client.
func NewMachinesClient() MachinesClient {
	return NewMachinesClientWithBaseURI(DefaultBaseURI)
}

// NewMachinesClientWithBaseURI creates an instance of the MachinesClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewMachinesClientWithBaseURI(baseURI string) MachinesClient {
	return MachinesClient{NewWithBaseURI(baseURI)}
}

// GetAllMachinesInSite sends the get all machines in site request.
// Parameters:
// subscriptionID - the ID of the target subscription.
// resourceGroupName - the name of the resource group. The name is case insensitive.
// siteName - site name.
// APIVersion - the API version to use for this operation.
// continuationToken - optional parameter for continuation token.
// totalRecordCount - total count of machines in the given site.
func (client MachinesClient) GetAllMachinesInSite(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string, filter string, top *int32, continuationToken string, totalRecordCount *int32) (result VMwareMachineCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MachinesClient.GetAllMachinesInSite")
		defer func() {
			sc := -1
			if result.vmmc.Response.Response != nil {
				sc = result.vmmc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.getAllMachinesInSiteNextResults
	req, err := client.GetAllMachinesInSitePreparer(ctx, subscriptionID, resourceGroupName, siteName, APIVersion, filter, top, continuationToken, totalRecordCount)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.MachinesClient", "GetAllMachinesInSite", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetAllMachinesInSiteSender(req)
	if err != nil {
		result.vmmc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.MachinesClient", "GetAllMachinesInSite", resp, "Failure sending request")
		return
	}

	result.vmmc, err = client.GetAllMachinesInSiteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.MachinesClient", "GetAllMachinesInSite", resp, "Failure responding to request")
		return
	}
	if result.vmmc.hasNextLink() && result.vmmc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// GetAllMachinesInSitePreparer prepares the GetAllMachinesInSite request.
func (client MachinesClient) GetAllMachinesInSitePreparer(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string, filter string, top *int32, continuationToken string, totalRecordCount *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"siteName":          autorest.Encode("path", siteName),
		"subscriptionId":    autorest.Encode("path", subscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(continuationToken) > 0 {
		queryParameters["continuationToken"] = autorest.Encode("query", continuationToken)
	}
	if totalRecordCount != nil {
		queryParameters["totalRecordCount"] = autorest.Encode("query", *totalRecordCount)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/VMwareSites/{siteName}/machines", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetAllMachinesInSiteSender sends the GetAllMachinesInSite request. The method will close the
// http.Response Body if it receives an error.
func (client MachinesClient) GetAllMachinesInSiteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetAllMachinesInSiteResponder handles the response to the GetAllMachinesInSite request. The method always
// closes the http.Response Body.
func (client MachinesClient) GetAllMachinesInSiteResponder(resp *http.Response) (result VMwareMachineCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// getAllMachinesInSiteNextResults retrieves the next set of results, if any.
func (client MachinesClient) getAllMachinesInSiteNextResults(ctx context.Context, lastResults VMwareMachineCollection) (result VMwareMachineCollection, err error) {
	req, err := lastResults.vMwareMachineCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "migrate.MachinesClient", "getAllMachinesInSiteNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.GetAllMachinesInSiteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "migrate.MachinesClient", "getAllMachinesInSiteNextResults", resp, "Failure sending next results request")
	}
	result, err = client.GetAllMachinesInSiteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.MachinesClient", "getAllMachinesInSiteNextResults", resp, "Failure responding to next results request")
	}
	return
}

// GetAllMachinesInSiteComplete enumerates all values, automatically crossing page boundaries as required.
func (client MachinesClient) GetAllMachinesInSiteComplete(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string, filter string, top *int32, continuationToken string, totalRecordCount *int32) (result VMwareMachineCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MachinesClient.GetAllMachinesInSite")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.GetAllMachinesInSite(ctx, subscriptionID, resourceGroupName, siteName, APIVersion, filter, top, continuationToken, totalRecordCount)
	return
}

// GetMachine sends the get machine request.
// Parameters:
// subscriptionID - the ID of the target subscription.
// resourceGroupName - the name of the resource group. The name is case insensitive.
// siteName - site name.
// machineName - machine ARM name.
// APIVersion - the API version to use for this operation.
func (client MachinesClient) GetMachine(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, machineName string, APIVersion string) (result VMwareMachine, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MachinesClient.GetMachine")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetMachinePreparer(ctx, subscriptionID, resourceGroupName, siteName, machineName, APIVersion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.MachinesClient", "GetMachine", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetMachineSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.MachinesClient", "GetMachine", resp, "Failure sending request")
		return
	}

	result, err = client.GetMachineResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.MachinesClient", "GetMachine", resp, "Failure responding to request")
		return
	}

	return
}

// GetMachinePreparer prepares the GetMachine request.
func (client MachinesClient) GetMachinePreparer(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, machineName string, APIVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"machineName":       autorest.Encode("path", machineName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"siteName":          autorest.Encode("path", siteName),
		"subscriptionId":    autorest.Encode("path", subscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/VMwareSites/{siteName}/machines/{machineName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetMachineSender sends the GetMachine request. The method will close the
// http.Response Body if it receives an error.
func (client MachinesClient) GetMachineSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetMachineResponder handles the response to the GetMachine request. The method always
// closes the http.Response Body.
func (client MachinesClient) GetMachineResponder(resp *http.Response) (result VMwareMachine, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// StartMachine sends the start machine request.
// Parameters:
// subscriptionID - the ID of the target subscription.
// resourceGroupName - the name of the resource group. The name is case insensitive.
// siteName - site name.
// machineName - machine ARM name.
// APIVersion - the API version to use for this operation.
func (client MachinesClient) StartMachine(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, machineName string, APIVersion string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MachinesClient.StartMachine")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.StartMachinePreparer(ctx, subscriptionID, resourceGroupName, siteName, machineName, APIVersion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.MachinesClient", "StartMachine", nil, "Failure preparing request")
		return
	}

	resp, err := client.StartMachineSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "migrate.MachinesClient", "StartMachine", resp, "Failure sending request")
		return
	}

	result, err = client.StartMachineResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.MachinesClient", "StartMachine", resp, "Failure responding to request")
		return
	}

	return
}

// StartMachinePreparer prepares the StartMachine request.
func (client MachinesClient) StartMachinePreparer(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, machineName string, APIVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"machineName":       autorest.Encode("path", machineName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"siteName":          autorest.Encode("path", siteName),
		"subscriptionId":    autorest.Encode("path", subscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/VMwareSites/{siteName}/machines/{machineName}/start", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StartMachineSender sends the StartMachine request. The method will close the
// http.Response Body if it receives an error.
func (client MachinesClient) StartMachineSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// StartMachineResponder handles the response to the StartMachine request. The method always
// closes the http.Response Body.
func (client MachinesClient) StartMachineResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// StopMachine sends the stop machine request.
// Parameters:
// subscriptionID - the ID of the target subscription.
// resourceGroupName - the name of the resource group. The name is case insensitive.
// siteName - site name.
// machineName - machine ARM name.
// APIVersion - the API version to use for this operation.
func (client MachinesClient) StopMachine(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, machineName string, APIVersion string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MachinesClient.StopMachine")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.StopMachinePreparer(ctx, subscriptionID, resourceGroupName, siteName, machineName, APIVersion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.MachinesClient", "StopMachine", nil, "Failure preparing request")
		return
	}

	resp, err := client.StopMachineSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "migrate.MachinesClient", "StopMachine", resp, "Failure sending request")
		return
	}

	result, err = client.StopMachineResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.MachinesClient", "StopMachine", resp, "Failure responding to request")
		return
	}

	return
}

// StopMachinePreparer prepares the StopMachine request.
func (client MachinesClient) StopMachinePreparer(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, machineName string, APIVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"machineName":       autorest.Encode("path", machineName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"siteName":          autorest.Encode("path", siteName),
		"subscriptionId":    autorest.Encode("path", subscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/VMwareSites/{siteName}/machines/{machineName}/stop", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StopMachineSender sends the StopMachine request. The method will close the
// http.Response Body if it receives an error.
func (client MachinesClient) StopMachineSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// StopMachineResponder handles the response to the StopMachine request. The method always
// closes the http.Response Body.
func (client MachinesClient) StopMachineResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}
