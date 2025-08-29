import { MenuItem } from '../models/menu.model';

export class Menu {
  public static pages: MenuItem[] = [
    {
      group: 'menu.dashboard',
      route: '/dashboard',
      items: []
    },
    {
      group: 'menu.offloadedNamespaces',
      route: '/namespaces',
      items: []
    },
    {
      group: 'menu.liqoStatus',
      route: '/status',
      items: []
    }
  ];
}
