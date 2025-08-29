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

import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { ResponsiveHelperComponent } from './components/responsive-helper/responsive-helper.component';
import { ClickOutsideDirective } from './directives/click-outside.directive';

import { FormsModule } from "@angular/forms";
import { RouterModule } from "@angular/router";
import { FontAwesomeModule } from "@fortawesome/angular-fontawesome";
import { TippyDirective } from "@ngneat/helipopper";
import { AgGridModule } from "ag-grid-angular";
import { AngularSvgIconModule } from "angular-svg-icon";
import { NgxEchartsModule } from "ngx-echarts";
import { NgxPaginationModule } from "ngx-pagination";
import { TranslocoRootModule } from "../transloco-root.module";
import { ClusterStatusBadge } from './components/cluster-status-badge/cluster-status-badge.component';
import { ClusterTableComponent } from "./components/cluster-table/cluster-table.component";
import { ClusterActionsRendererComponent } from "./components/cluster-table/renderer/cluster-actions-renderer.component";
import { ClusterStatusRendererComponent } from './components/cluster-table/renderer/cluster-status-renderer.component';
import { ModuleStatusBadge } from './components/module-status-badge/module-status-badge.component';
import { NamespaceTableComponent } from "./components/namespace-table/namespace-table.component";
import {
  NamespaceActionsRendererComponent
} from "./components/namespace-table/renderer/namespace-actions-renderer.component";
import {
  NamespaceStatusRendererComponent
} from "./components/namespace-table/renderer/namespace-status-renderer.component";
import { PeeringTableComponent } from "./components/peering-table/peering-table.component";
import { PodTableComponent } from "./components/pod-table/pod-table.component";
import { PodLabelsRendererComponent } from "./components/pod-table/renderer/pod-labels-renderer.component";
import { PodStatusRendererComponent } from "./components/pod-table/renderer/pod-status-renderer.component";
@NgModule({
  declarations: [
    ResponsiveHelperComponent,
    ClickOutsideDirective,
    ClusterTableComponent,
    NamespaceTableComponent,
    NamespaceActionsRendererComponent,
    NamespaceStatusRendererComponent,
    PeeringTableComponent,
    PodTableComponent,
    PodLabelsRendererComponent,
    PodStatusRendererComponent,
    ClusterActionsRendererComponent,
    ClusterStatusRendererComponent,
    ClusterStatusBadge,
    ModuleStatusBadge
  ],
  imports: [
    FormsModule,
    CommonModule,
    AngularSvgIconModule.forRoot(),
    RouterModule,
    AgGridModule,
    TippyDirective,
    NgxPaginationModule,
    TranslocoRootModule,
    FontAwesomeModule,
    NgxEchartsModule.forRoot({
      /**
       * This will import all modules from echarts.
       * If you only need custom modules,
       * please refer to [Custom Build] section.
       */
      echarts: () => import('echarts'), // or import('./path-to-my-custom-echarts')
    })],
  exports: [
    ResponsiveHelperComponent,
    ClickOutsideDirective,
    ClusterTableComponent,
    ClusterStatusBadge,
    ModuleStatusBadge,
    NamespaceTableComponent,
    PeeringTableComponent,
    ClusterActionsRendererComponent,
    NamespaceActionsRendererComponent,
    NamespaceStatusRendererComponent,
    PodTableComponent,
    PodLabelsRendererComponent,
    PodStatusRendererComponent,
    AgGridModule,
    FormsModule,
    NgxEchartsModule,
    TippyDirective,
    TranslocoRootModule,
    FontAwesomeModule
  ],
})
export class SharedModule { }
