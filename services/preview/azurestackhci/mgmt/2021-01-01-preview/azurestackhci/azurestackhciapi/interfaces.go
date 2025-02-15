package azurestackhciapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/azurestackhci/mgmt/2021-01-01-preview/azurestackhci"
	"github.com/Azure/go-autorest/autorest"
)

// ArcSettingsClientAPI contains the set of methods on the ArcSettingsClient type.
type ArcSettingsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, arcSetting azurestackhci.ArcSetting) (result azurestackhci.ArcSetting, err error)
	Delete(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string) (result azurestackhci.ArcSettingsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string) (result azurestackhci.ArcSetting, err error)
	ListByCluster(ctx context.Context, resourceGroupName string, clusterName string) (result azurestackhci.ArcSettingListPage, err error)
	ListByClusterComplete(ctx context.Context, resourceGroupName string, clusterName string) (result azurestackhci.ArcSettingListIterator, err error)
}

var _ ArcSettingsClientAPI = (*azurestackhci.ArcSettingsClient)(nil)

// ClustersClientAPI contains the set of methods on the ClustersClient type.
type ClustersClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, clusterName string, cluster azurestackhci.Cluster) (result azurestackhci.Cluster, err error)
	Delete(ctx context.Context, resourceGroupName string, clusterName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, clusterName string) (result azurestackhci.Cluster, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result azurestackhci.ClusterListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result azurestackhci.ClusterListIterator, err error)
	ListBySubscription(ctx context.Context) (result azurestackhci.ClusterListPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result azurestackhci.ClusterListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, clusterName string, cluster azurestackhci.ClusterPatch) (result azurestackhci.Cluster, err error)
}

var _ ClustersClientAPI = (*azurestackhci.ClustersClient)(nil)

// ExtensionsClientAPI contains the set of methods on the ExtensionsClient type.
type ExtensionsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, extensionName string, extension azurestackhci.Extension) (result azurestackhci.ExtensionsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, extensionName string) (result azurestackhci.ExtensionsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, extensionName string) (result azurestackhci.Extension, err error)
	ListByArcSetting(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string) (result azurestackhci.ExtensionListPage, err error)
	ListByArcSettingComplete(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string) (result azurestackhci.ExtensionListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, extensionName string, extension azurestackhci.Extension) (result azurestackhci.ExtensionsUpdateFuture, err error)
}

var _ ExtensionsClientAPI = (*azurestackhci.ExtensionsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result azurestackhci.OperationListResult, err error)
}

var _ OperationsClientAPI = (*azurestackhci.OperationsClient)(nil)
