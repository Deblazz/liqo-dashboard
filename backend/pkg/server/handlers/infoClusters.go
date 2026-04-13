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

package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/liqotech/liqo/pkg/liqoctl/factory"

	"github.com/ArubaKube/liqo-dashboard/pkg/utils/getters"
)

// GetV1Info implements the `GET /v1/info` route, returning a simple "ok" status to indicate that the server is running.
func (s Server) GetV1Info(c *gin.Context) {
	ctx := c.Request.Context()

	f := factory.Factory{
		LiqoNamespace: s.liqoNamespace,
		KubeClient:    s.nativeClient,
		CRClient:      s.oClient,
	}

	clusterInfo, err := getters.GetInfoClusters(ctx, &f)
	if err != nil {
		log.Printf("Error collecting cluster info: %v", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, clusterInfo)
}
