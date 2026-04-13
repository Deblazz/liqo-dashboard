import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LayoutComponent } from './layout.component';

const routes: Routes = [
  {
    path: 'dashboard',
    component: LayoutComponent,
    loadChildren: () => import('../dashboard/dashboard.module').then((m) => m.DashboardModule),
  },
  {
    path: 'clusters',
    component: LayoutComponent,
    loadChildren: () => import('../cluster/cluster.module').then((m) => m.ClusterModule),
  },
  {
    path: 'namespaces',
    component: LayoutComponent,
    loadChildren: () => import('../namespace/namespace.module').then((m) => m.NamespaceModule),
  },
  {
    path: 'status',
    component: LayoutComponent,
    loadChildren: () => import('../status/status.module').then((m)=> m.StatusModule),
  },
  {
    path: 'pods',
    component: LayoutComponent,
    loadChildren: () => import('../pod/pod.module').then((m) => m.PodModule),
  },
  { path: '', redirectTo: 'dashboard', pathMatch: 'full' },
  { path: '**', redirectTo: 'error/404' },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class LayoutRoutingModule {}
