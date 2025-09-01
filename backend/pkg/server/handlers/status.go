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
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/liqotech/liqo/pkg/liqoctl/factory"
	"github.com/liqotech/liqo/pkg/liqoctl/info"
	"github.com/liqotech/liqo/pkg/liqoctl/info/localstatus"
)

// GetV1StatusLocalInfo implements the `GET /v1/status` route, returning a simple "ok" status to indicate that the server is running.
func (s Server) GetV1StatusLocalInfo(c *gin.Context) {
	ctx := c.Request.Context()

	f := factory.NewForLocal()
	if err := f.Initialize(); err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("factory init failed: %v", err)})
		return
	}

	o := info.NewOptions(f)

	checkers := []info.Checker{
		&localstatus.InstallationChecker{},
		&localstatus.HealthChecker{},
		&localstatus.NetworkChecker{},
		&localstatus.PeeringChecker{},
	}

	// Start collecting the data via the checkers
	for i := range checkers {
		checkers[i].Collect(ctx, *o)
		for _, err := range checkers[i].GetCollectionErrors() {
			o.Printer.Warning.Println(err)
		}
	}

	var err error
	var output string
	switch {
	// If no format is specified, format and print a user-friendly output
	case o.Format == "" && o.GetQuery == "":
		for i := range checkers {
			o.Printer.BoxSetTitle(checkers[i].GetTitle())
			o.Printer.BoxPrintln(checkers[i].Format(*o))
		}
	// If query specified try to retrieve the field from the output
	case o.GetQuery != "":
		data := collectDataFromCheckers(checkers)
		log.Println(data)
	default:
		data := collectDataFromCheckers(checkers)
		log.Println(data)

	}

	if err != nil {
		o.Printer.Error.Println(err)
	} else {
		fmt.Println(output)
	}

	c.JSON(200, gin.H{"status": "checkers eseguiti, vedi console per dettagli"})

}

func collectDataFromCheckers(checkers []info.Checker) map[string]interface{} {
	data := map[string]interface{}{}

	for i := range checkers {
		data[checkers[i].GetID()] = checkers[i].GetData()
	}

	return data
}
