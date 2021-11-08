import { Component, OnInit } from '@angular/core';
import { VisualCircleService } from 'gongleaflet';

import * as L from 'leaflet';

import * as manageLeafletItems from '../manage-leaflet-items';

@Component({
  selector: 'app-cartoatc-circles',
  templateUrl: './cartoatc-circles.component.html',
  styleUrls: ['./cartoatc-circles.component.scss']
})
export class CartoatcCirclesComponent implements OnInit {
  render: Array<L.Layer> = [];

  constructor(
    private visualCircleService: VisualCircleService,
  ) { }

  ngOnInit(): void {
    this.visualCircleService.getVisualCircles().subscribe(
      (visualCircles) => {
        visualCircles.forEach(visualCircle => {
          this.render.push(manageLeafletItems.newCircle(visualCircle));
        })
      }
    );
  }
}
