package getters

import (
	"context"
	"fmt"

	"github.com/liqotech/liqo/pkg/liqoctl/factory"
	"github.com/liqotech/liqo/pkg/liqoctl/info"
	"github.com/liqotech/liqo/pkg/liqoctl/info/localstatus"

	"github.com/ArubaKube/liqo-dashboard/pkg/server/models"
)

// GetInfoClusters like liqotcl info --verbose, returns info about the local cluster and the peered pods.
func GetInfoClusters(ctx context.Context, f *factory.Factory) (models.ClusterInfo, error) {
	o := info.NewOptions(f)

	checkers := []info.Checker{
		&localstatus.InstallationChecker{},
		&localstatus.HealthChecker{},
		&localstatus.NetworkChecker{},
		&localstatus.PeeringChecker{},
	}

	for i := range checkers {
		checkers[i].Collect(ctx, *o)
		if errs := checkers[i].GetCollectionErrors(); len(errs) > 0 {
			return models.ClusterInfo{}, fmt.Errorf("error collecting data: %v", errs)
		}
	}
	return buildInfoClustersFromCheckers(checkers), nil
}

func buildInfoClustersFromCheckers(checkers []info.Checker) models.ClusterInfo {
	ci := models.ClusterInfo{
		Health:   models.HealthInfo{UnhealthyPods: make(map[string]models.PodHealth)},
		Local:    models.LocalClusterInfo{Labels: make(map[string]string)},
		Network:  models.NetworkInfo{},
		Peerings: models.PeeringsInfo{Peers: []models.PeerInfo{}},
	}

	for _, checker := range checkers {
		switch checker.GetID() {
		case "health":
			if data, ok := checker.GetData().(localstatus.Health); ok {
				ci.Health = parseHealthInfo(data)
			}
		case "local":
			if data, ok := checker.GetData().(localstatus.Installation); ok {
				ci.Local = parseLocalInfo(data)
			}
		case "network":
			if data, ok := checker.GetData().(localstatus.Network); ok {
				ci.Network = parseNetworkInfo(data)
			}
		case "peerings":
			if data, ok := checker.GetData().(localstatus.Peerings); ok {
				ci.Peerings = parsePeeringsInfo(data)
			}
		}
	}

	return ci
}

func parseHealthInfo(data localstatus.Health) models.HealthInfo {
	h := models.HealthInfo{
		Healthy:       data.Healthy,
		UnhealthyPods: make(map[string]models.PodHealth),
	}

	for name, pod := range data.UnhealthyPods {
		h.UnhealthyPods[name] = models.PodHealth{
			Status:          string(pod.Status),
			ReadyContainers: pod.ReadyContainers,
			TotalContainers: pod.TotalContainers,
			Restarts:        int(pod.Restarts),
		}
	}

	return h
}

func parseLocalInfo(data localstatus.Installation) models.LocalClusterInfo {
	l := models.LocalClusterInfo{
		ClusterID:     string(data.ClusterID),
		Version:       data.Version,
		APIServerAddr: data.APIServerAddr,
		Labels:        make(map[string]string),
	}

	for k, v := range data.Labels {
		l.Labels[k] = v
	}

	return l
}

func parseNetworkInfo(data localstatus.Network) models.NetworkInfo {
	return models.NetworkInfo{
		PodCIDR:      data.PodCIDR,
		ServiceCIDR:  data.ServiceCIDR,
		ExternalCIDR: data.ExternalCIDR,
		InternalCIDR: data.InternalCIDR,
	}
}

func parsePeeringsInfo(data localstatus.Peerings) models.PeeringsInfo {
	p := models.PeeringsInfo{Peers: []models.PeerInfo{}}
	for _, peer := range data.Peers {
		p.Peers = append(p.Peers, models.PeerInfo{
			ClusterID:            string(peer.ClusterID),
			Role:                 string(peer.Role),
			NetworkingStatus:     string(peer.NetworkingStatus),
			AuthenticationStatus: string(peer.AuthenticationStatus),
			OffloadingStatus:     string(peer.OffloadingStatus),
		})
	}
	return p
}
