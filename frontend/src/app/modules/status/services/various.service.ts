import { Injectable } from '@angular/core';
import { of } from 'rxjs';

@Injectable({ providedIn: 'root' })
export class VariousService {
  getTestStats() {
    return of({
      icon: '📊',
      title: 'Liqo network status',
      items: [
        { label: 'Utenti', value: '120' },
        { label: 'Ordini', value: '45' },
      ]
    });
  }
}
