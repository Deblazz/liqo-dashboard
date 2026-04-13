package models

// ClusterInfo collects various pieces of info for a cluster. Coming from liqotcl info.
type ClusterInfo struct {
	Health   HealthInfo       `json:"health"`
	Local    LocalClusterInfo `json:"local"`
	Network  NetworkInfo      `json:"network"`
	Peerings PeeringsInfo     `json:"peerings"`
}

// HealthInfo contains info about health for the cluster.
type HealthInfo struct {
	Healthy       bool                 `json:"healthy"`
	UnhealthyPods map[string]PodHealth `json:"unhealthyPods,omitempty"`
}

// PodHealth describes the health status of peered pods.
type PodHealth struct {
	Status          string `json:"status"`
	ReadyContainers int    `json:"readyContainers"`
	TotalContainers int    `json:"totalContainers"`
	Restarts        int    `json:"restarts"`
}

// LocalClusterInfo contains details about the local cluster.
type LocalClusterInfo struct {
	ClusterID     string            `json:"clusterID"`
	Version       string            `json:"version"`
	Labels        map[string]string `json:"labels,omitempty"`
	APIServerAddr string            `json:"APIServerAddr"`
}

// NetworkInfo contains network information of the local cluster.
type NetworkInfo struct {
	PodCIDR      string `json:"podCIDR"`
	ServiceCIDR  string `json:"serviceCIDR"`
	ExternalCIDR string `json:"externalCIDR"`
	InternalCIDR string `json:"internalCIDR"`
}

// PeeringsInfo collects info about the peered pods.
type PeeringsInfo struct {
	Peers []PeerInfo `json:"peers"`
}

// PeerInfo describes the single peered pods.
type PeerInfo struct {
	ClusterID            string `json:"clusterID"`
	Role                 string `json:"role"`
	NetworkingStatus     string `json:"networkingStatus"`
	AuthenticationStatus string `json:"authenticationStatus"`
	OffloadingStatus     string `json:"offloadingStatus"`
}
