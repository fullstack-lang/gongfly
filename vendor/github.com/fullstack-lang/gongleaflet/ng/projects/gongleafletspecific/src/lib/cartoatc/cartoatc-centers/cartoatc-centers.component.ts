import { Component, Input, OnInit } from '@angular/core';

import * as gongleaflet from 'gongleaflet';

import { DEFAULT_ICON_SIZE } from '../cartoatc.component'

import * as manageLeafletItems from '../manage-leaflet-items';
import * as L from 'leaflet'

@Component({
  selector: 'app-cartoatc-centers',
  templateUrl: './cartoatc-centers.component.html',
  styleUrls: ['./cartoatc-centers.component.scss'],
})
export class CartoatcCentersComponent implements OnInit {
  @Input() visualCenters?: Array<gongleaflet.VisualCenterDB>;

  // store relation between the VisualCenters & the markers
  mapVisualCenterID_LeafletMarker = new Map<number, L.Marker>();

  gongleafletFrontRepo?: gongleaflet.FrontRepo

  centersLayer: Array<L.Layer> = [];
  map_visualIconID_visualIconSVG = new Map<number, string>();

  constructor(
    private gongleafletFrontRepoService: gongleaflet.FrontRepoService,
    private visualCenterService: gongleaflet.VisualCenterService,
    private visualIcon: gongleaflet.VisualIconService
  ) { }

  ngOnInit(): void {
    this.refreshMapWithVisualCircle()

    this.visualCenterService.VisualCenterServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refreshMapWithVisualCircle()
        }
      })
  }

  refreshMapWithVisualCircle() {
    this.gongleafletFrontRepoService.pull().subscribe(
      gongleafletFrontRepo => {
        this.gongleafletFrontRepo = gongleafletFrontRepo

        this.gongleafletFrontRepo.VisualIcons.forEach((visualIcon) => {
          if (!this.map_visualIconID_visualIconSVG.has(visualIcon.ID)) {
            this.map_visualIconID_visualIconSVG.set(visualIcon.ID, visualIcon.SVG);
          }
        });

        this.gongleafletFrontRepo.VisualCenters.forEach((visualCenter) => {

          if (!this.mapVisualCenterID_LeafletMarker.has(visualCenter.ID)) {
            var color = manageLeafletItems.getColor(visualCenter.VisualColorEnum);

            var icon: L.DivIcon = manageLeafletItems.newIcon(
              visualCenter.ID,
              'layer-' + visualCenter.VisualLayerID.Int64,
              this.map_visualIconID_visualIconSVG.get(visualCenter.VisualIconID.Int64)!,
              DEFAULT_ICON_SIZE,
              color,
              visualCenter.Name
            );
            var marker: L.Marker
            marker = manageLeafletItems.newMarkerWithIcon(
              visualCenter.Lat,
              visualCenter.Lng,
              icon
            );
            this.centersLayer.push(marker);

            this.mapVisualCenterID_LeafletMarker.set(visualCenter.ID, marker)
          }
        })
      }
    )
  }
}
