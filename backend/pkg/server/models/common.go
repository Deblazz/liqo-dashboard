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
	"k8s.io/apimachinery/pkg/api/resource"
)

// Resources represents the resources shared by a cluster or available in a node.
type Resources struct {
	CPU              resource.Quantity `json:"cpu"`
	Memory           resource.Quantity `json:"memory"`
	GPU              resource.Quantity `json:"gpu"`
	Pods             resource.Quantity `json:"pods"`
	EphemeralStorage resource.Quantity `json:"ephemeralStorage"`
}
