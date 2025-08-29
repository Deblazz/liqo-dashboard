import { Component, OnInit } from '@angular/core';
import { ColDef } from "ag-grid-community";

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

  public peerings: ForeignCluster[] = [];

  columnDefs: ColDef[] = [];

  public defaultColDef: ColDef = {
    sortable: true,
    filter: true,
    resizable: true,
  };

  ngOnInit(): void {
    this.peerings = [
      {
        id: 'cluster-1',
        role: 'Consumer',
        networkingStatus: 'Unhealty',
        authenticationStatus: 'Unhealthy',
        offloadingStatus: 'Unhealthy'
      },
      {
        id: 'cluster-2',
        role: 'Provider',
        networkingStatus: 'Healthy',
        authenticationStatus: 'Healthy',
        offloadingStatus: 'Healthy'
      },
      {
        id: 'cluster-3',
        role: 'ConsumerAndProvider',
        networkingStatus: 'Unhealty',
        authenticationStatus: 'Unhealthy',
        offloadingStatus: 'Unhealthy'
      }
    ];

    this.columnDefs = [
      { headerName: 'ID', field: 'id' },
      { headerName: 'Role', field: 'role' },
      { headerName: 'Networking Status', field: 'networkingStatus' },
      { headerName: 'Authentication Status', field: 'authenticationStatus' },
      { headerName: 'Offloading Status', field: 'offloadingStatus' },
    ];
  }
}
