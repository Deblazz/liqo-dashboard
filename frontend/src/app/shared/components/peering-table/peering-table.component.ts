/**
 * Copyright 2025 ArubaKube S.r.l.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import { Component, OnInit, Input } from '@angular/core';
import { ColDef } from "ag-grid-community";
import { TranslocoService } from '@jsverse/transloco';
import { PeeringInfoRendererComponent } from './renderer/peering-info.renderer.component';

export interface ForeignCluster {
  id: string;
  role: 'Consumer' | 'Provider' | 'ConsumerAndProvider' | 'Unknown';
  networkingStatus: string;
  authenticationStatus: string;
  offloadingStatus: string;
}

@Component({
  selector: '[peering-table]',
  templateUrl: './peering-table.component.html',
})
export class PeeringTableComponent implements OnInit {
  @Input() peerings: Record<string, any>[] | null = []; constructor(private translateService: TranslocoService) { }

  columnDefs: ColDef[] = [];

  public defaultColDef: ColDef = {
    sortable: true,
    filter: true,
    resizable: true,
  };

  ngOnInit(): void {

    this.columnDefs = [
      { headerName: this.translateService.translate('status.peering.clusterIdLabel'), field: 'clusterID' },
      { headerName: this.translateService.translate('status.peering.roleLabel'), field: 'role' },
      {
        headerName: this.translateService.translate('status.peering.networkingStatusLabel'), field: 'networkingStatus', cellClass: params =>
          this.getStatusClass(params.value),
      },
      {
        headerName: this.translateService.translate('status.peering.authStatusLabel'), field: 'authenticationStatus', cellClass: params =>
          this.getStatusClass(params.value),
      },
      {
        headerName: this.translateService.translate('status.peering.offloadingStatusLabel'), field: 'offloadingStatus', cellClass: params =>
          this.getStatusClass(params.value),
      },
      {
        cellRenderer: PeeringInfoRendererComponent,
        sortable: false,
        filter: false,
        cellStyle: { display: 'flex', justifyContent: 'center', alignItems: 'center' }
      }
    ];
    
  }
  getStatusClass(status: string): string {
    switch (status.toLowerCase()) {
      case 'healthy':
        return 'status-healthy';
      case 'unhealthy':
        return 'status-unhealthy';
      default:
        return '';
    }
  }


  reload() {
    this.ngOnInit()
  }
}
