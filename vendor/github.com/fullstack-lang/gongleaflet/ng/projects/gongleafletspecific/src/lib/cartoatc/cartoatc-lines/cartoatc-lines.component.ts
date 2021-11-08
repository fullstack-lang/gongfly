import { Component, OnInit } from '@angular/core';
import { combineLatest, Observable, timer } from 'rxjs';

import * as L from 'leaflet';

import * as manageLeafletItems from '../manage-leaflet-items';

import * as gongleaflet from 'gongleaflet';

// import 'leaflet-movingMarker'

@Component({
  selector: 'app-cartoatc-lines',
  templateUrl: './cartoatc-lines.component.html',
  styleUrls: ['./cartoatc-lines.component.scss'],
})
export class CartoatcLinesComponent implements OnInit {
  linesLayer: Array<L.Layer> = [];

  mapVisualLineID_LeafletPolyline = new Map<number, L.Polyline>();

  // store message markers
  mapVisualLineID_LeafletMovingMarker = new Map<number, L.Marker>();

  messageMarkers: Map<number, L.Marker> = new Map<number, L.Marker>();
  messageIsTransmitting: Map<number, string> = new Map<number, string>();
  messageMarkersBackward: Map<number, L.Marker> = new Map<number, L.Marker>();
  messageIsTransmittingBackward: Map<number, string> = new Map<
    number,
    string
  >()

  // timer that is called regularly to refresh the map
  obsTimer: Observable<number> = timer(1000, 500) // due time 1', period
  currTime: number = 0

  constructor(
    private visualLineService: gongleaflet.VisualLineService
  ) {

  }

  ngOnInit(): void {
    // init polylines
    combineLatest([this.visualLineService.getVisualLines()]).subscribe(
      ([visualLines]) => {
        visualLines.forEach((visualLine) => {
          var polyline: L.Polyline = new L.Polyline([]);
          polyline = manageLeafletItems.setLine(visualLine);

          this.linesLayer.push(polyline);
          this.mapVisualLineID_LeafletPolyline.set(visualLine.ID, polyline);
        });
      }
    );

    this.visualLineService.VisualLineServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          // update line positions
          this.visualLineService.getVisualLines().subscribe((visualLines) => {
            visualLines.forEach((visualLine) => {
              var visualLineMarker = this.mapVisualLineID_LeafletPolyline.get(
                visualLine.ID
              );

              if (visualLineMarker) {
                // update position
                visualLineMarker.setLatLngs([
                  [visualLine.StartLat, visualLine.StartLng],
                  [visualLine.EndLat, visualLine.EndLng],
                ]);
                visualLineMarker.options.color = manageLeafletItems.getColor(
                  visualLine.VisualColorEnum
                );
                visualLineMarker.setStyle(visualLineMarker.options);
              }
            }
            );
          }
          );
        }
      }
    )
  }
}

