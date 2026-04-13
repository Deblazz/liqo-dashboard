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
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Server is the implementation of the REST api interfaces.
type Server struct {
	liqoNamespace string
	oClient       client.Client
	nativeClient  kubernetes.Interface
}

// NewServer returns a new REST api server implementation.
func NewServer(oClient client.Client, nativeClient kubernetes.Interface, liqoNamespace string) Server {
	if liqoNamespace == "" {
		liqoNamespace = "liqo" // default namespace
	}
	return Server{
		liqoNamespace: liqoNamespace,
		oClient:       oClient,
		nativeClient:  nativeClient,
	}
}
