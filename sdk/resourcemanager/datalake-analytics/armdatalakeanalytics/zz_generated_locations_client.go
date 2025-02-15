//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatalakeanalytics

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

// LocationsClient contains the methods for the Locations group.
// Don't use this type directly, use NewLocationsClient() instead.
type LocationsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewLocationsClient creates a new instance of LocationsClient with the specified values.
func NewLocationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *LocationsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &LocationsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// GetCapability - Gets subscription-level properties and limits for Data Lake Analytics specified by resource location.
// If the operation fails it returns the *ErrorResponse error type.
func (client *LocationsClient) GetCapability(ctx context.Context, location string, options *LocationsGetCapabilityOptions) (LocationsGetCapabilityResponse, error) {
	req, err := client.getCapabilityCreateRequest(ctx, location, options)
	if err != nil {
		return LocationsGetCapabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LocationsGetCapabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNotFound) {
		return LocationsGetCapabilityResponse{}, client.getCapabilityHandleError(resp)
	}
	return client.getCapabilityHandleResponse(resp)
}

// getCapabilityCreateRequest creates the GetCapability request.
func (client *LocationsClient) getCapabilityCreateRequest(ctx context.Context, location string, options *LocationsGetCapabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DataLakeAnalytics/locations/{location}/capability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getCapabilityHandleResponse handles the GetCapability response.
func (client *LocationsClient) getCapabilityHandleResponse(resp *http.Response) (LocationsGetCapabilityResponse, error) {
	result := LocationsGetCapabilityResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CapabilityInformation); err != nil {
		return LocationsGetCapabilityResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getCapabilityHandleError handles the GetCapability error response.
func (client *LocationsClient) getCapabilityHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
