import { Component, Input } from '@angular/core';

@Component({
  selector: 'lib-gongfly-dashboard',
  templateUrl: './gongfly-dashboard.component.html',
  styleUrls: ['./gongfly-dashboard.component.css']
})
export class GongflyDashboardComponent {

  @Input() GONG__StackPath: string = ""

}
