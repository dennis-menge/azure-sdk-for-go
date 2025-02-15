//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package azurestackhci

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/azurestackhci/mgmt/2020-10-01/azurestackhci"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ProvisioningState = original.ProvisioningState

const (
	Accepted     ProvisioningState = original.Accepted
	Canceled     ProvisioningState = original.Canceled
	Failed       ProvisioningState = original.Failed
	Provisioning ProvisioningState = original.Provisioning
	Succeeded    ProvisioningState = original.Succeeded
)

type Status = original.Status

const (
	ConnectedRecently    Status = original.ConnectedRecently
	Disconnected         Status = original.Disconnected
	Error                Status = original.Error
	NotConnectedRecently Status = original.NotConnectedRecently
	NotYetRegistered     Status = original.NotYetRegistered
)

type AvailableOperations = original.AvailableOperations
type AzureEntityResource = original.AzureEntityResource
type BaseClient = original.BaseClient
type Cluster = original.Cluster
type ClusterList = original.ClusterList
type ClusterListIterator = original.ClusterListIterator
type ClusterListPage = original.ClusterListPage
type ClusterNode = original.ClusterNode
type ClusterProperties = original.ClusterProperties
type ClusterReportedProperties = original.ClusterReportedProperties
type ClusterUpdate = original.ClusterUpdate
type ClustersClient = original.ClustersClient
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorDetail = original.ErrorDetail
type ErrorResponse = original.ErrorResponse
type OperationDetail = original.OperationDetail
type OperationDisplay = original.OperationDisplay
type OperationsClient = original.OperationsClient
type ProxyResource = original.ProxyResource
type Resource = original.Resource
type TrackedResource = original.TrackedResource

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewClusterListIterator(page ClusterListPage) ClusterListIterator {
	return original.NewClusterListIterator(page)
}
func NewClusterListPage(cur ClusterList, getNextPage func(context.Context, ClusterList) (ClusterList, error)) ClusterListPage {
	return original.NewClusterListPage(cur, getNextPage)
}
func NewClustersClient(subscriptionID string) ClustersClient {
	return original.NewClustersClient(subscriptionID)
}
func NewClustersClientWithBaseURI(baseURI string, subscriptionID string) ClustersClient {
	return original.NewClustersClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleStatusValues() []Status {
	return original.PossibleStatusValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
