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

package models

import (
	liqov1beta1 "github.com/liqotech/liqo/apis/core/v1beta1"
)

// ForeignCluster is a struct that represents a foreign cluster.
type ForeignCluster struct {
	ID                   liqov1beta1.ClusterID           `json:"id"`
	Role                 liqov1beta1.RoleType            `json:"role"`
	APIServerURL         string                          `json:"apiServerUrl"`
	APIServerStatus      liqov1beta1.ConditionStatusType `json:"apiServerStatus"`
	NetworkStatus        liqov1beta1.ConditionStatusType `json:"networkStatus"`
	AuthenticationStatus liqov1beta1.ConditionStatusType `json:"authenticationStatus"`
	OffloadingStatus     liqov1beta1.ConditionStatusType `json:"offloadingStatus"`
	NetworkLatency       string                          `json:"networkLatency"`
	ResourcesOffered     Resources                       `json:"resourcesOffered"`
	ResourcesAcquired    Resources                       `json:"resourcesAcquired"`
}
