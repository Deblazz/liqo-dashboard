import { Component } from '@angular/core';
import { ICellRendererAngularComp } from 'ag-grid-angular';
import { TranslocoService } from '@jsverse/transloco';
import { Router } from '@angular/router';

@Component({
  selector: 'app-details-link-renderer',
  template: `
    <a class="btn btn-primary" [routerLink]="['/clusters/detail', params.data.clusterID]">
      {{ translateService.translate('status.peering.detailsButtonLabel') }}
    </a>
  `
})

export class PeeringInfoRendererComponent implements ICellRendererAngularComp {
  params: any;

  constructor(public translateService: TranslocoService, private router: Router) {}

  agInit(params: any): void { this.params = params; }
  refresh(): boolean { return false; }
}
