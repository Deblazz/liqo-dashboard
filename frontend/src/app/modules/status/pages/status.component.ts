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

import { Component, OnInit } from '@angular/core';
import { SpinnerNgxAdapterService } from "../../../core/services/spinner/spinner-ngx-adapter.service";
import { StatusService } from '../services/status.service';

@Component({
  selector: 'app-status',
  templateUrl: './status.component.html',
})
export class StatusComponent implements OnInit {
  liqoInfo: any;
  constructor(private statusService: StatusService, private spinnerService: SpinnerNgxAdapterService) { }

  ngOnInit(): void {
    this.spinnerService.show();
    this.statusService.getLiqoInfo().subscribe({
      next: (data) => {
        this.liqoInfo = data;
        this.spinnerService.hide();
      },
      error: () => {
        this.spinnerService.hide();
      }
    });
  }
}
