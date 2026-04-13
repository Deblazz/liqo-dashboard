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

// Package main, contains the entrypoint of the application.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/ArubaKube/liqo-dashboard/pkg/server/api"
	"github.com/ArubaKube/liqo-dashboard/pkg/server/handlers"
	"github.com/ArubaKube/liqo-dashboard/pkg/utils"
)

func main() {
	listeningPort := flag.Uint("port", 8080, "The port where the rest API will be exposed")
	liqoNamespace := flag.String("liqo-namespace", "liqo", "The namespace where liqo is installed")
	flag.Parse()

	// Get the k8s client to provide to the server
	oClient, err := utils.GetClient()
	if err != nil {
		log.Fatalf("Error getting client: %v", err)
	}

	nativeClient, err := utils.GetNativeClient()
	if err != nil {
		log.Fatalf("Error getting native client: %v", err)
	}

	// create a type that satisfies the `api.ServerInterface`, which contains an implementation of every operation from the generated code
	server := handlers.NewServer(oClient, nativeClient, *liqoNamespace)

	r := gin.Default()

	api.RegisterHandlers(r, server)

	// And we serve HTTP until the world ends.

	s := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("0.0.0.0:%d", *listeningPort),
		// Enforce timeouts for server
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Println("Starting server on port 8080")
	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
