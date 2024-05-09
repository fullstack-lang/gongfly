import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'lib-gongfly-dashboard',
  templateUrl: './gongfly-dashboard.component.html',
  styleUrls: ['./gongfly-dashboard.component.css']
})
export class GongflyDashboardComponent implements OnInit {

  @Input() GONG__SIM__StackPath: string = ""
  @Input() GONG__LEAFLET__StackPath: string = ""

  loading: boolean = true

  ngOnInit(): void {
    this.loading = false
  }

}
