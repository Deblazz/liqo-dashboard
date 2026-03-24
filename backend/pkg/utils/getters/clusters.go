// Copyright 2025 ArubaKube S.r.l.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package getters

import (
	"context"
	"strings"

	authv1beta1 "github.com/liqotech/liqo/apis/authentication/v1beta1"
	liqov1beta1 "github.com/liqotech/liqo/apis/core/v1beta1"
	offloadingv1beta1 "github.com/liqotech/liqo/apis/offloading/v1beta1"
	liqoconsts "github.com/liqotech/liqo/pkg/consts"
	fcutils "github.com/liqotech/liqo/pkg/utils/foreigncluster"
	liqogetters "github.com/liqotech/liqo/pkg/utils/getters"
	liqolabels "github.com/liqotech/liqo/pkg/utils/labels"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metricsv1beta1 "k8s.io/metrics/pkg/apis/metrics/v1beta1"
	"k8s.io/utils/ptr"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/ArubaKube/liqo-dashboard/pkg/server/models"
)

// GetForeignClusters returns all the ForeignClusters.
func GetForeignClusters(ctx context.Context, cl client.Client) ([]models.ForeignCluster, error) {
	var foreignClusterList liqov1beta1.ForeignClusterList
	if err := cl.List(ctx, &foreignClusterList); err != nil {
		return nil, err
	}

	foreignClusters := make([]models.ForeignCluster, len(foreignClusterList.Items))
	for i := range foreignClusterList.Items {
		fc, err := parseForeignCluster(ctx, cl, &foreignClusterList.Items[i])
		if err != nil {
			return nil, err
		}
		foreignClusters[i] = fc
	}

	return foreignClusters, nil
}

// GetForeignClusterByID returns the ForeignCluster with the given clusterID.
func GetForeignClusterByID(ctx context.Context, cl client.Client, clusterID string) (*models.ForeignCluster, error) {
	cluster, err := fcutils.GetForeignClusterByID(ctx, cl, liqov1beta1.ClusterID(clusterID))
	if err != nil {
		return nil, err
	}

	fc, err := parseForeignCluster(ctx, cl, cluster)
	if err != nil {
		return nil, err
	}

	return ptr.To(fc), nil
}

// GetVirtualNodesByClusterID returns the VirtualNodes related to the given clusterID.
func GetVirtualNodesByClusterID(ctx context.Context, cl client.Client, clusterID string) ([]models.Node, error) {
	virtualNodes, err := liqogetters.ListVirtualNodesByClusterID(ctx, cl, liqov1beta1.ClusterID(clusterID))
	if err != nil {
		return nil, err
	}

	var nodes []models.Node
	for i := range virtualNodes {
		parsedNode, err := parseVirtualNode(ctx, cl, &virtualNodes[i])
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, parsedNode)
	}

	return nodes, nil
}

func parseVirtualNode(ctx context.Context, cl client.Client, vnode *offloadingv1beta1.VirtualNode) (models.Node, error) {
	nodeName := vnode.Name // the name of the virtual node is the same as the name of the node
	nodeExists := ptr.Deref(vnode.Spec.CreateNode, false)

	var capacity models.Resources
	var capacityUsed models.Resources
	if nodeExists {
		// Get the resources shared with this cluster
		var node corev1.Node
		if err := cl.Get(ctx, client.ObjectKey{Name: nodeName}, &node); err != nil {
			return models.Node{}, err
		}

		var gpuCapacity resource.Quantity
		for key, val := range node.Status.Capacity {
			if strings.Contains(string(key), "gpu") {
				gpuCapacity.Add(val)
			}
		}
		capacity = models.Resources{
			CPU:              *node.Status.Capacity.Cpu(),
			Memory:           *node.Status.Capacity.Memory(),
			GPU:              gpuCapacity,
			Pods:             *node.Status.Capacity.Pods(),
			EphemeralStorage: *node.Status.Capacity.StorageEphemeral(),
		}

		var nodeMetrics metricsv1beta1.NodeMetrics
		if err := cl.Get(ctx, client.ObjectKey{Name: nodeName}, &nodeMetrics); err != nil {
			return models.Node{}, err
		}
		capacityUsed = models.Resources{
			CPU:              *nodeMetrics.Usage.Cpu(),
			Memory:           *nodeMetrics.Usage.Memory(),
			Pods:             *nodeMetrics.Usage.Pods(),
			EphemeralStorage: *nodeMetrics.Usage.StorageEphemeral(),
		}
	}

	return models.Node{
		Name:         vnode.Name,
		ClusterID:    vnode.Spec.ClusterID,
		Capacity:     capacity,
		CapacityUsed: capacityUsed,
	}, nil
}

func parseForeignCluster(ctx context.Context, cl client.Client, fc *liqov1beta1.ForeignCluster) (models.ForeignCluster, error) {
	// Get the resources acquired from this cluster (foreigncluster is a provider)
	resourcesAcquired, err := getResourcesAcquired(ctx, cl, fc)
	if err != nil {
		return models.ForeignCluster{}, err
	}

	// Get the resources offered to this cluster (foreigncluster is a consumer)
	resourcesOffered, err := getResourcesOffered(ctx, cl, fc)
	if err != nil {
		return models.ForeignCluster{}, err
	}

	// Get the network latency
	networkLatency, err := getNetworkLatency(ctx, cl, fc)
	if err != nil {
		return models.ForeignCluster{}, err
	}

	return models.ForeignCluster{
		ID:                   fc.Spec.ClusterID,
		Role:                 fc.Status.Role,
		APIServerURL:         fc.Status.APIServerURL,
		APIServerStatus:      fcutils.GetAPIServerStatus(fc),
		NetworkStatus:        getNetworkStatus(fc),
		AuthenticationStatus: getAuthenticationStatus(fc),
		OffloadingStatus:     getOffloadingStatus(fc),
		NetworkLatency:       networkLatency,
		ResourcesAcquired:    resourcesAcquired,
		ResourcesOffered:     resourcesOffered,
	}, nil
}

func getResourcesAcquired(ctx context.Context, cl client.Client, fc *liqov1beta1.ForeignCluster) (models.Resources, error) {
	localResSlices, err := liqogetters.ListResourceSlicesByLabel(ctx, cl, corev1.NamespaceAll,
		liqolabels.LocalLabelSelectorForCluster(string(fc.Spec.ClusterID)))
	if err != nil {
		return models.Resources{}, err
	}

	return getResourcesFromResourceSlice(localResSlices), nil
}

func getResourcesOffered(ctx context.Context, cl client.Client, fc *liqov1beta1.ForeignCluster) (models.Resources, error) {
	remoteResSlices, err := liqogetters.ListResourceSlicesByLabel(ctx, cl, corev1.NamespaceAll,
		liqolabels.RemoteLabelSelectorForCluster(string(fc.Spec.ClusterID)))
	if err != nil {
		return models.Resources{}, err
	}

	return getResourcesFromResourceSlice(remoteResSlices), nil
}

func getResourcesFromResourceSlice(resSlices []authv1beta1.ResourceSlice) models.Resources {
	// Initialize the total offered resources to 0
	cpuTot := resource.NewQuantity(0, resource.DecimalSI)
	memTot := resource.NewQuantity(0, resource.BinarySI)
	gpuTot := resource.NewQuantity(0, resource.DecimalSI)
	podsTot := resource.NewQuantity(0, resource.DecimalSI)
	storageTot := resource.NewQuantity(0, resource.BinarySI)

	for i := range resSlices {
		if resSlices[i].Status.Resources == nil {
			continue
		}
		cpuTot.Add(*resSlices[i].Status.Resources.Cpu())
		memTot.Add(*resSlices[i].Status.Resources.Memory())
		podsTot.Add(*resSlices[i].Status.Resources.Pods())
		storageTot.Add(*resSlices[i].Status.Resources.StorageEphemeral())

		for key, val := range resSlices[i].Status.Resources {
			if strings.Contains(string(key), "gpu") {
				gpuTot.Add(val)
			}
		}
	}

	return models.Resources{
		CPU:              *cpuTot,
		Memory:           *memTot,
		GPU:              *gpuTot,
		Pods:             *podsTot,
		EphemeralStorage: *storageTot,
	}
}

func getNetworkStatus(fc *liqov1beta1.ForeignCluster) liqov1beta1.ConditionStatusType {
	if !fc.Status.Modules.Networking.Enabled {
		return liqov1beta1.ConditionStatusNone
	}

	return fcutils.GetStatus(fc.Status.Modules.Networking.Conditions, liqov1beta1.NetworkConnectionStatusCondition)
}

func getAuthenticationStatus(fc *liqov1beta1.ForeignCluster) liqov1beta1.ConditionStatusType {
	if !fc.Status.Modules.Authentication.Enabled {
		return liqov1beta1.ConditionStatusNone
	}

	switch fc.Status.Role {
	case liqov1beta1.ConsumerRole:
		return fcutils.GetStatus(fc.Status.Modules.Authentication.Conditions, liqov1beta1.AuthTenantStatusCondition)
	case liqov1beta1.ProviderRole:
		return fcutils.GetStatus(fc.Status.Modules.Authentication.Conditions, liqov1beta1.AuthIdentityControlPlaneStatusCondition)
	case liqov1beta1.ConsumerAndProviderRole:
		tenantCond := fcutils.GetStatus(fc.Status.Modules.Authentication.Conditions, liqov1beta1.AuthTenantStatusCondition)
		identityCond := fcutils.GetStatus(fc.Status.Modules.Authentication.Conditions, liqov1beta1.AuthIdentityControlPlaneStatusCondition)
		if tenantCond == liqov1beta1.ConditionStatusReady && identityCond == liqov1beta1.ConditionStatusReady {
			return liqov1beta1.ConditionStatusReady
		}
		if tenantCond == liqov1beta1.ConditionStatusError || identityCond == liqov1beta1.ConditionStatusError {
			return liqov1beta1.ConditionStatusError
		}
		return liqov1beta1.ConditionStatusNone
	default:
		return liqov1beta1.ConditionStatusNone
	}
}

func getOffloadingStatus(fc *liqov1beta1.ForeignCluster) liqov1beta1.ConditionStatusType {
	if !fc.Status.Modules.Offloading.Enabled {
		return liqov1beta1.ConditionStatusNone
	}

	switch fc.Status.Role {
	case liqov1beta1.ProviderRole, liqov1beta1.ConsumerAndProviderRole:
		vnodeCond := fcutils.GetStatus(fc.Status.Modules.Offloading.Conditions, liqov1beta1.OffloadingVirtualNodeStatusCondition)
		nodeCond := fcutils.GetStatus(fc.Status.Modules.Offloading.Conditions, liqov1beta1.OffloadingNodeStatusCondition)
		if vnodeCond == liqov1beta1.ConditionStatusReady && nodeCond == liqov1beta1.ConditionStatusReady {
			return liqov1beta1.ConditionStatusReady
		}
		if vnodeCond == liqov1beta1.ConditionStatusError || nodeCond == liqov1beta1.ConditionStatusError {
			return liqov1beta1.ConditionStatusError
		}
		return liqov1beta1.ConditionStatusNone
	default:
		return liqov1beta1.ConditionStatusNone
	}
}

func getNetworkLatency(ctx context.Context, cl client.Client, fc *liqov1beta1.ForeignCluster) (string, error) {
	if !fc.Status.Modules.Networking.Enabled {
		return liqoconsts.NotApplicable, nil
	}

	connection, err := liqogetters.GetConnectionByClusterID(ctx, cl, string(fc.Spec.ClusterID))
	switch {
	case client.IgnoreNotFound(err) != nil:
		return "", err
	case apierrors.IsNotFound(err):
		// The connection object does not exist yet. We cannot evaluate the latency.
		return liqoconsts.NotApplicable, nil
	default:
		return connection.Status.Latency.Value, nil
	}
}
