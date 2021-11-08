import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { combineLatest, timer, Observable } from 'rxjs';

import * as L from 'leaflet';
import 'leaflet-rotatedmarker';

import * as gongleaflet from 'gongleaflet';

import * as manageLeafletItems from './manage-leaflet-items';
import { VisualTrackDB } from 'gongleaflet';

import { dotBlur } from '../../assets/icons/dot_blur';

export const DEFAULT_ICON_SIZE = 60

@Component({
  selector: 'app-cartoatc',
  templateUrl: './cartoatc.component.html',
  styleUrls: ['./cartoatc.component.scss'],
})
export class CartoatcComponent implements OnInit {
  mapOptions: any = null;
  visualLayers: L.Layer[] = [];
  visualTracksHistory: L.Layer[] = [];

  // map of visualTrack ID to visualTrackMarker in order to perform updates
  mapVisualTrackID_VisualMarker = new Map<number, L.Marker>();

  // map of visualTrackMarker to visualTrack ID in order to delete deleted visualTrack
  mapVisualMarker_VisualTrackID = new Map<L.Marker, number>();

  traceLayerID: number = 0
  icons = new Map<number, string>();

  visualCenters: Array<gongleaflet.VisualCenterDB> = new Array();

  tracksHistories: Map<string, Array<L.LatLng>> = new Map();

  constructor(
    private frontRepo: gongleaflet.FrontRepoService,
    private visualTrackService: gongleaflet.VisualTrackService,
    private router: Router
  ) {
    this.router.routeReuseStrategy.shouldReuseRoute = () => false;
  }

  onMapReady(map: L.Map) {
    setTimeout(() => {
      map.invalidateSize();
    }, 0);
  }

  ngOnInit(): void {
    combineLatest([this.frontRepo.pull()]).subscribe(([frontRepo]) => {
      this.visualCenters = Array.from(frontRepo.VisualCenters.values());
      this.mapOptions = manageLeafletItems.setMapOptions(
        Array.from(frontRepo.VisualMaps.values())[0]
      );
      let visualMapDB = frontRepo.VisualMaps_array[0];
      console.log(
        'map options ' +
        visualMapDB.Lat +
        ' ' +
        visualMapDB.Lng +
        ' ' +
        visualMapDB.UrlTemplate
      );
    });

    this.visualTrackService.VisualTrackServiceChanged.next('update');

    // update visual line when the data has changed
    // observable for changes in structs
    this.visualTrackService.VisualTrackServiceChanged.subscribe((message) => {
      if (message == 'post' || message == 'update' || message == 'delete') {
        // update track position by using the front repo
        this.frontRepo.VisualTrackPull().subscribe((frontRepo) => {
          // update marker from visual track
          frontRepo.VisualTracks.forEach(
            (vTrack: gongleaflet.VisualTrackDB) => {
              let _currentMarker: L.Marker<any> = this.mapVisualTrackID_VisualMarker.get(vTrack.ID)!
              if (!_currentMarker) {
                this.manageNewVisualTrackMarker(vTrack);
              } else {
                this.manageUpdateVisualTrackMarker(vTrack, _currentMarker);
              }
            }
          );

          // remove markers tat have no visual tracks
          this.mapVisualMarker_VisualTrackID.forEach((visualTrackID) => {
            if (frontRepo.VisualTracks.get(visualTrackID) == undefined) {
              var marker = this.mapVisualTrackID_VisualMarker.get(
                visualTrackID
              );

              // remove marker from the visual layer
              marker?.remove();

              this.mapVisualTrackID_VisualMarker.delete(visualTrackID);
              this.mapVisualMarker_VisualTrackID.delete(marker!);
            }
          });
        });
      }
    });
  }

  manageNewVisualTrackMarker(visualTrack: gongleaflet.VisualTrackDB) {
    var color = manageLeafletItems.getColor(visualTrack.VisualColorEnum);

    if (visualTrack.VisualIcon) {
      var icon: L.DivIcon = manageLeafletItems.newIcon(
        visualTrack.ID + '-track',
        'layer-' + visualTrack.VisualLayerID.Int64,
        visualTrack.VisualIcon.SVG,
        DEFAULT_ICON_SIZE,
        color,
        visualTrack.Name
      );
      var marker = manageLeafletItems.newMarkerWithIcon(
        visualTrack.Lat,
        visualTrack.Lng,
        icon
      );

      {
        // @ts-ignore: Unreachable code error
        marker.setRotationOrigin('center center');
      }

      manageLeafletItems.rotateIcon(
        visualTrack.ID + '-track',
        visualTrack.Heading
      );
      this.traceLayerID = visualTrack.VisualLayerID.Int64;

      this.mapVisualTrackID_VisualMarker.set(visualTrack.ID, marker);
      this.mapVisualMarker_VisualTrackID.set(marker, visualTrack.ID);
      this.visualLayers.push(marker);
    }
  }

  manageUpdateVisualTrackMarker(
    visualTrack: gongleaflet.VisualTrackDB,
    marker: L.Marker<any>
  ) {
    if (!visualTrack.Display) {
      marker.remove();
      return;
    }
    marker.setLatLng([visualTrack.Lat, visualTrack.Lng]);
    manageLeafletItems.setIconLabel(
      visualTrack.ID + '-track',
      this.formatTrackLabel(visualTrack)
    );
    manageLeafletItems.rotateIcon(
      visualTrack.ID + '-track',
      visualTrack.Heading
    );
    if (visualTrack.DisplayTrackHistory) {
      this.manageTracksHistory(
        visualTrack.Name,
        L.latLng(visualTrack.Lat, visualTrack.Lng)
      );
      this.visualTracksHistory = this.renderTracksLayer;
    }
  }

  manageTracksHistory(trackName: string, coordinates: L.LatLng) {
    let trackHistory: L.LatLng[] = [];
    const LIMIT_HISTORY_LENGTH = 10;
    trackHistory = this.tracksHistories.get(trackName)!
    const pushManageLimit = (
      list: Array<L.LatLng>,
      newItem: L.LatLng
    ): Array<L.LatLng> => {
      let tmpList = list;
      if (tmpList.length >= LIMIT_HISTORY_LENGTH) {
        tmpList.shift();
      }
      tmpList.push(newItem);
      return tmpList;
    };

    if (!trackHistory) {
      trackHistory = [];
    }
    if (trackHistory.length) {
      if (
        trackHistory[trackHistory.length - 1].lat !== coordinates.lat &&
        trackHistory[trackHistory.length - 1].lng !== coordinates.lng
      ) {
        trackHistory = pushManageLimit(trackHistory, coordinates);
      } else {
        if (trackHistory.length < LIMIT_HISTORY_LENGTH) {
          trackHistory.push(coordinates);
        }
      }
    } else {
      trackHistory.push(coordinates);
    }
    this.tracksHistories.set(trackName, trackHistory);
  }

  get renderTracksLayer(): Array<L.Layer> {
    let render: L.Layer[] = [];
    let icon = manageLeafletItems.newIcon(
      'icon',
      'layer-' + this.traceLayerID,
      dotBlur,
      5,
      '#004E92'
    );
    this.tracksHistories.forEach((trackHistory) => {
      trackHistory.map((coordinates: L.LatLng) => {
        render.push(
          manageLeafletItems.newMarkerWithIcon(
            coordinates.lat,
            coordinates.lng,
            icon
          )
        );
      });
    });
    return render;
  }

  formatTrackLabel = (track: VisualTrackDB): string => {
    let label: string = track.Name;
    if (track.DisplayLevelAndSpeed) {
      label += `</br>
      ${track.Level.toFixed(0)} - ${track.Speed.toFixed(0)}</br>`;
    }
    return label;
  };
}
