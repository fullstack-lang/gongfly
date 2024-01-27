import { Component, Input, NgZone, OnInit } from '@angular/core'
import { Router } from '@angular/router'
import { combineLatest, timer, Observable } from 'rxjs'

import * as L from 'leaflet'
import 'leaflet-rotatedmarker'

import * as gongleaflet from 'gongleaflet';
import * as manageLeafletItems from './manage-leaflet-items'
import { dotBlur } from '../../assets/icons/dot_blur'
import { UserClick } from 'gongleaflet'

export const DEFAULT_ICON_SIZE = 60

// MapoptionsComponent is an ngx-asymterix angular component (itself based in leaflet) that is
// - the component that displays all layers
// A gong stack can handle multiple Mapoptions and therefore display more than one map.
@Component({
  selector: 'app-mapoptions',
  templateUrl: './mapoptions.component.html',
  styleUrls: ['./mapoptions.component.scss'],
})
export class MapoptionsComponent implements OnInit {

  @Input() GONG__StackPath: string = ""

  // 1. name of the initial map (must match the name in the backend)
  // 2. the corresponding gong MapOptions object
  // 3. the corresponding [leafletOptions]="mapOptions" that is passed to the leaflet map in the html
  @Input() mapName: string = ""
  mapOptionsID: number = 0
  leafletMapOptions?: L.MapOptions // stangely, impossible to type without ?

  // callback function called on user click 
  @Input() userInterfactionCallbackFunction?: (lat: number, lng: number) => void


  // [leafletLayers]="rootOfLayerGroups" that is passed to one div in the html, ngx-asymetrix
  // https://github.com/Asymmetrik/ngx-leaflet#add-custom-layers-base-layers-markers-shapes-etc
  rootOfLayerGroups: L.Layer[] = [];

  //
  // Visual Track stuff
  //

  // map of visualTrack ID to visualTrackMarker in order to perform updates
  mapVisualTrackID_LeafletTrackMarker = new Map<number, L.Marker>();

  // map of visualTrackMarker to visualTrack ID in order to delete deleted visualTrack
  mapLeafletTrackMarker_VisualTrackID = new Map<L.Marker, number>();

  // mapVisualTrackName_positionsHistory stores tracks histories
  mapVisualTrackID_positionsHistory = new Map<number, Array<L.LatLng>>()
  mapVisualTrackID_LeafletHistoryTrackMarker = new Map<number, L.Marker[]>();

  //
  // Other objets
  //
  // map that store leaflet object according to the gong object ID
  mapGongLayerGroupID_LeafletLayerGroup = new Map<number, L.LayerGroup<L.Layer>>()
  mapVLineID_LeafletPolyline = new Map<number, L.Polyline>()
  mapCircleID_LeafletCircle = new Map<number, L.Circle>()
  mapMarkerID_LeafletMarker = new Map<number, L.Marker>()
  mapGongDivIconID_divIconSVG = new Map<number, string>()

  // // map that stores which 
  // mapGongLayerGroupUseID_GongLayerGroupID = new Map<number, number>()

  // map that allows direct access from the Gong LayerGroupID to the LayerGroupUse of the map
  // if a gong layerGroup is present, it means the layer is is to be displayed on this map
  mapGongLayerGroupID_LayerGroupUse = new Map<number, gongleaflet.LayerGroupUse>()

  // the gong front repo
  frontRepo?: gongleaflet.FrontRepo

  // autonmatic refresh of maps
  obsTimer: Observable<number> = timer(1000, 500) // due time 1', period 0.5'

  // for debug purpose
  leafletMap?: L.Map

  currTime: number = 0

  // commitNb stores the number of commit on the backend
  commitNb: number = 0

  // commitNb stores the number of commit on the frontend
  commitFromFrontNb: number = 0

  DotLeafletDivIcon = manageLeafletItems.newIcon(
    'icon',
    'layer-',
    dotBlur,
    5,
    '#004E92'
  );

  clickNumber: number = 0

  constructor(
    public frontRepoService: gongleaflet.FrontRepoService,
    private visualTrackService: gongleaflet.VisualTrackService,
    private lineService: gongleaflet.VLineService,
    private markerService: gongleaflet.MarkerService,
    private layerGroupUseService: gongleaflet.LayerGroupUseService,
    private userClickService: gongleaflet.UserClickService,
    private commitNbFromBackService: gongleaflet.CommitNbFromBackService,
    private pushFromFrontService: gongleaflet.PushFromFrontNbService,
    private router: Router,
    public zone: NgZone
  ) {
    this.router.routeReuseStrategy.shouldReuseRoute = () => false;
  }

  // not yet clear what those lines mean
  // anyway, the is a paramter in ngx-asymetrix
  // https://github.com/Asymmetrik/ngx-leaflet#leafletmapready-map
  onMapReady(leafletMap: L.Map) {

    this.leafletMap = leafletMap

    setTimeout(() => {
      leafletMap.invalidateSize();
    }, 0);
  }


  onMapClick(e: L.LeafletMouseEvent) {

    if (this.leafletMap) {
      // let popup = new L.Popup()
      // popup.setLatLng(e.latlng)
      // popup.setContent("You clicked the map at " + e.latlng.toString())
      // popup.openOn(this.leafletMap)

      if (this.userInterfactionCallbackFunction) {
        this.userInterfactionCallbackFunction(e.latlng.lat, e.latlng.lng)
      }

      let userClick = new UserClick
      userClick.Name = "Click !" + this.clickNumber
      this.clickNumber = this.clickNumber + 1
      userClick.Lat = e.latlng.lat
      userClick.Lng = e.latlng.lng

      this.userClickService.postFront(userClick, this.GONG__StackPath).subscribe(
        (userClick) => {
          console.log("user clicked")
          this.userClickService.UserClickServiceChanged.next("post")
        }
      )

    }
  }

  ngOnInit(): void {

    this.frontRepoService.pull(this.GONG__StackPath).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        let gongMapOptions = Array.from(this.frontRepo.getFrontArray<gongleaflet.MapOptions>(gongleaflet.MapOptions.GONGSTRUCT_NAME).values())[0]

        // if the map name is set, then map options might differ
        if (this.mapName != "") {
          for (let gongleafletMapOptions of this.frontRepo.getFrontArray<gongleaflet.MapOptions>(gongleaflet.MapOptions.GONGSTRUCT_NAME).values()) {
            if (gongleafletMapOptions.Name == this.mapName) {
              this.mapOptionsID = gongleafletMapOptions.ID
              gongMapOptions = gongleafletMapOptions
            }
          }
        }

        // set the map options that will set up the angular component
        this.leafletMapOptions = manageLeafletItems.visualMapToLeafletMapOptions(gongMapOptions)
        this.refreshMapWithMarkers()

        this.commitNbFromBackService.getCommitNbFromBack(500, this.GONG__StackPath).subscribe(
          commitNbFromBack => {
            // console.log("commit nb in the back " + commitNb + " local commit nb " + this.commitNb)
            if (commitNbFromBack > this.commitNb) {
              this.refreshMapWithMarkers()
              this.commitNb = commitNbFromBack
            }
          }
        )

      }
    )
  }

  // generate a new leaflet marker and store it in the mapVisualTrackID_VisualMarker
  newVisualTrackMarker(visualTrack: gongleaflet.VisualTrack) {
    var color = manageLeafletItems.getColor(visualTrack.ColorEnum);

    if (visualTrack.DivIcon) {
      var icon: L.DivIcon = manageLeafletItems.newIcon(
        visualTrack.ID + '-track',
        'layer-',
        visualTrack.DivIcon.SVG,
        DEFAULT_ICON_SIZE,
        color,
        visualTrack.Name
      );
      var leafletTrackMarker = manageLeafletItems.newMarkerWithIcon(
        visualTrack.Lat,
        visualTrack.Lng,
        icon
      );

      {
        // @ts-ignore: Unreachable code error
        leafletTrackMarker.setRotationOrigin('center center');
      }

      manageLeafletItems.rotateIcon(
        visualTrack.ID + '-track',
        visualTrack.Heading
      );

      this.mapVisualTrackID_LeafletTrackMarker.set(visualTrack.ID, leafletTrackMarker);
      this.mapLeafletTrackMarker_VisualTrackID.set(leafletTrackMarker, visualTrack.ID);

      // get layer of visual track
      let gongLayerGroup = visualTrack.LayerGroup
      if (gongLayerGroup) {
        let leafletLayer = this.mapGongLayerGroupID_LeafletLayerGroup.get(gongLayerGroup.ID)

        if (leafletLayer) {
          leafletTrackMarker.addTo(leafletLayer)
        }
      }

      //      this.rootOfLayerGroups.push(leafletTrackMarker);
    }
  }

  // update the leaflet track marker with the gong VisualTrack info
  //
  // - position
  // - information on the label
  // - heading
  // - track history
  updateVisualTrackMarker(
    visualTrack: gongleaflet.VisualTrack,
    marker: L.Marker<any>
  ) {
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
      let trackHistory = this.generateVisualTracksHistory(visualTrack)
      this.renderTrackHistory(visualTrack, trackHistory)
    }
  }

  // generateVisualTracksHistory adds dots to the track
  generateVisualTracksHistory(visualTrack: gongleaflet.VisualTrack): L.LatLng[] {
    let trackHistory: L.LatLng[] = []

    let coordinates: L.LatLng = new L.LatLng(visualTrack.Lat, visualTrack.Lng)

    // get the track pas positions
    trackHistory = this.mapVisualTrackID_positionsHistory.get(visualTrack.ID)!

    if (!trackHistory) {
      trackHistory = []
    }
    if (trackHistory.length) {
      if (
        trackHistory[trackHistory.length - 1].lat !== coordinates.lat &&
        trackHistory[trackHistory.length - 1].lng !== coordinates.lng
      ) {
        trackHistory = addNewCoordToFIFO(trackHistory, coordinates)
      } else {
        if (trackHistory.length < LIMIT_HISTORY_LENGTH) {
          trackHistory.push(coordinates)
        }
      }
    } else {
      trackHistory.push(coordinates);
    }
    this.mapVisualTrackID_positionsHistory.set(visualTrack.ID, trackHistory)

    return trackHistory
  }

  // renderTrackHistory update the layer of the visual track
  // with dots along the track history
  //
  // remove the former dot history
  renderTrackHistory(visualTrack: gongleaflet.VisualTrack, trackHistory: L.LatLng[]) {

    // get the leaflet layer of the visual track
    let leafletGroupLayer: L.LayerGroup<L.Layer> | undefined
    let groupLayerUse: gongleaflet.LayerGroupUse | undefined
    if (visualTrack.LayerGroup) {
      leafletGroupLayer = this.mapGongLayerGroupID_LeafletLayerGroup.get(visualTrack.LayerGroup.ID)
      groupLayerUse = this.mapGongLayerGroupID_LayerGroupUse.get(visualTrack.LayerGroup.ID)
    } else {
      return
    }

    // remove the ancient history
    let arrayOfDotMarkers = this.mapVisualTrackID_LeafletHistoryTrackMarker.get(visualTrack.ID)

    if (arrayOfDotMarkers) {
      for (let dotMarker of arrayOfDotMarkers) {
        dotMarker.removeFrom(this.leafletMap!)

      }
    }

    // workaround to an issue. If the layer is not part of the root of the layers, the dot icon is
    // not removed properly. Therefore, it is important NOT to add dotIcon is the layer is not set
    if (groupLayerUse?.IsDisplayed == false) {
      return
    }

    arrayOfDotMarkers = new Array<L.Marker>()

    for (let coordinates of trackHistory) {
      let dotIcon = manageLeafletItems.newMarkerWithIcon(
        coordinates.lat,
        coordinates.lng,
        this.DotLeafletDivIcon
      )
      dotIcon.addTo(leafletGroupLayer!)
      arrayOfDotMarkers.push(dotIcon)
    }

    // store the array in the map
    this.mapVisualTrackID_LeafletHistoryTrackMarker.set(visualTrack.ID, arrayOfDotMarkers)
  }

  formatTrackLabel = (track: gongleaflet.VisualTrack): string => {
    let label: string = track.Name;
    if (track.DisplayLevelAndSpeed) {
      label += `</br>
      ${track.Level.toFixed(0)} - ${track.Speed.toFixed(0)}</br>`;
    }
    return label;
  };

  //
  // manage layers
  manageLayers() {
    // Management of layers.
    //
    // 1. get all layerGroupUse by the mapOption and store them in the map "mapGongLayerGroupID_LayerGroup"
    // of layers that have to be displayed
    // If the LayerGroupID is in not in the layer group and to be displayed, add it to the root of LayersGroup
    // If not, remove it from the root of LayersGroup if it is present

    // reset the map of layers that have to be displayed on this map
    this.mapGongLayerGroupID_LayerGroupUse.clear()

    // populate the map with information from layerGroupUse of this map
    let gongleafletMapOptions = this.frontRepo!.getFrontMap<gongleaflet.MapOptions>(gongleaflet.MapOptions.GONGSTRUCT_NAME).get(this.mapOptionsID)
    for (let gongLayerGroupUse of gongleafletMapOptions?.LayerGroupUses!) {
      let gongLayerGroup = gongLayerGroupUse.LayerGroup
      if (gongLayerGroup) {
        this.mapGongLayerGroupID_LayerGroupUse.set(gongLayerGroup.ID, gongLayerGroupUse)

        // if not present, create a leaflet layer group and add it to the root
        let leafletLayerGroup = this.mapGongLayerGroupID_LeafletLayerGroup.get(gongLayerGroup.ID)
        if (!leafletLayerGroup) {
          leafletLayerGroup = new L.LayerGroup<L.Marker>()
          this.mapGongLayerGroupID_LeafletLayerGroup.set(gongLayerGroup.ID, leafletLayerGroup)
        }

        //
        // for each leaflet layerGroup, the algo can do three things
        // 1. Nothing
        //   a. because it is not present and it has to be hidden
        //   b. because it is present and it has to be added
        // 2. Add it to the root of layer groups
        //   a. because it is not present and it has to be added
        // 3. Remove it from the root of layer groups
        //   a. because it is present and it has to be removed

        // is the leaflet layerGroup in the root of layer groups ?
        // if it is there, no need to add it if it has to be displayed
        // but there is a need to remove it if it has not to be displayed
        let layerAlreadyDisplayed = this.rootOfLayerGroups.find(present => present == leafletLayerGroup)

        let hasToBeRemoved: boolean = false
        let hasToBeAdded: boolean = false

        // does the LayerGroup has to be displayed ?
        if (gongLayerGroupUse?.IsDisplayed) {
          // The layer group has to be displayed

          // if the leaflet not layer already in the root of all LayerGroup, add it
          if (!layerAlreadyDisplayed) {
            hasToBeAdded = true
          }
        } else {
          // the layer has to be hidden
          // is it present ?
          if (layerAlreadyDisplayed) {
            hasToBeRemoved = true
          }
        }

        // performed computed operation
        if (hasToBeAdded) {
          // console.log("map " + this.mapName + " has to add layer group named " + gongLayerGroup.Name)
          this.rootOfLayerGroups.push(leafletLayerGroup)
        }

        if (hasToBeRemoved) {
          // console.log("map " + this.mapName + " has to remove layer group named " + gongLayerGroup.Name)
          this.rootOfLayerGroups.forEach((element, index) => {
            if (element == layerAlreadyDisplayed) this.rootOfLayerGroups.splice(index, 1);
          });
        }
      }
    }
  }

  // icons
  manageDivIcons() {
    // pair gong divIcon with leaflet divIcon
    for (let divIcon of this.frontRepo!.getFrontArray<gongleaflet.DivIcon>(gongleaflet.DivIcon.GONGSTRUCT_NAME)) {
      if (!this.mapGongDivIconID_divIconSVG.has(divIcon.ID)) {
        this.mapGongDivIconID_divIconSVG.set(divIcon.ID, divIcon.SVG);
      }
    }
  }

  // markers
  manageMakers() {

    for (let _marker of this.frontRepo!.getFrontArray<gongleaflet.Marker>(gongleaflet.Marker.GONGSTRUCT_NAME)) {

      // get the leaflet kin of the gong Marker
      let leafletMarker: L.Marker | undefined
      leafletMarker = this.mapMarkerID_LeafletMarker.get(_marker.ID)

      // if absent, create the kin
      if (!leafletMarker) {
        // console.log("Gong Marker " + gongMarker.Name + " has no leaflet kin")
        var color = manageLeafletItems.getColor(_marker.ColorEnum);

        var icon: L.DivIcon = manageLeafletItems.newIcon(
          _marker.ID,
          'layer-',
          _marker.DivIcon?.SVG!,
          DEFAULT_ICON_SIZE,
          color,
          _marker.Name
        );

        // creation
        leafletMarker = manageLeafletItems.newMarkerWithIcon(
          _marker.Lat,
          _marker.Lng,
          icon
        )

        // get the leallet layerGroup of the marker
        let leafletLayerGroup: L.LayerGroup<L.Layer> | undefined
        let markerLayerGroup = _marker.LayerGroup
        if (markerLayerGroup) {
          leafletLayerGroup = this.mapGongLayerGroupID_LeafletLayerGroup.get(markerLayerGroup.ID)
          if (leafletLayerGroup) {
            leafletMarker.addTo(leafletLayerGroup)
          }
        }

        // add the kin to the map
        this.mapMarkerID_LeafletMarker.set(_marker.ID, leafletMarker)
      } else {
        // console.log("Gong Marker " + gongMarker.Name + " has already a leaflet kin")
      }
    }
  }

  manageCircles() {

    // display circles
    for (let gongCircle of this.frontRepo!.getFrontArray<gongleaflet.Circle>(gongleaflet.Circle.GONGSTRUCT_NAME)) {

      var leafletCircle = this.mapCircleID_LeafletCircle.get(gongCircle.ID)

      if (!leafletCircle) {
        // get layer of the circle
        let gongLayerGroup = gongCircle.LayerGroup
        if (gongLayerGroup) {
          // is this layer present in the current map ?
          let leafletLayerGroup = this.mapGongLayerGroupID_LeafletLayerGroup.get(gongLayerGroup.ID)

          if (leafletLayerGroup) {
            let leafletCircle = manageLeafletItems.newCircle(gongCircle)
            leafletCircle?.addTo(leafletLayerGroup)

            this.mapCircleID_LeafletCircle.set(gongCircle.ID, leafletCircle);
          }
        }
      } else {

      }
    }
  }

  manageVLines() {

    for (let gongVLine of this.frontRepo?.getFrontArray<gongleaflet.VLine>(gongleaflet.VLine.GONGSTRUCT_NAME)!) {

      var leafletPolyline = this.mapVLineID_LeafletPolyline.get(gongVLine.ID)

      //
      // if lealet has no sister element of the VLine, then create one
      //
      if (!leafletPolyline) {
        // get layer of the circle
        let gongLayerGroup = gongVLine.LayerGroup
        if (gongLayerGroup) {
          // is this layer present in the current map ?
          let leafletLayerGroup = this.mapGongLayerGroupID_LeafletLayerGroup.get(gongLayerGroup.ID)

          if (leafletLayerGroup) {

            leafletPolyline = new L.Polyline([]);
            leafletPolyline = manageLeafletItems.newLine(gongVLine);

            leafletPolyline.addTo(leafletLayerGroup)
            this.mapVLineID_LeafletPolyline.set(gongVLine.ID, leafletPolyline);
          }
        }
      }

      if (leafletPolyline) {
        // update position
        leafletPolyline.setLatLngs([
          [gongVLine.StartLat, gongVLine.StartLng],
          [gongVLine.EndLat, gongVLine.EndLng],
        ]);
        leafletPolyline.options.color = manageLeafletItems.getColor(
          gongVLine.ColorEnum
        );
        leafletPolyline.setStyle(leafletPolyline.options);
      }
    }
  }

  manageTracks() {
    // update marker from visual track
    for (let vTrack of this.frontRepo!.getFrontArray<gongleaflet.VisualTrack>(gongleaflet.VisualTrack.GONGSTRUCT_NAME)) {
      let _currentMarker: L.Marker<any> = this.mapVisualTrackID_LeafletTrackMarker.get(vTrack.ID)!
      if (!_currentMarker) {
        this.newVisualTrackMarker(vTrack);
      } else {
        this.updateVisualTrackMarker(vTrack, _currentMarker);
      }
    }

    // remove markers that have no visual tracks
    this.mapLeafletTrackMarker_VisualTrackID.forEach((visualTrackID) => {
      if (this.frontRepo!.getFrontMap<gongleaflet.VisualTrack>(gongleaflet.VisualTrack.GONGSTRUCT_NAME).get(visualTrackID) == undefined) {
        var marker = this.mapVisualTrackID_LeafletTrackMarker.get(
          visualTrackID
        );

        // remove marker from the visual layer
        marker?.remove();

        this.mapVisualTrackID_LeafletTrackMarker.delete(visualTrackID);
        this.mapLeafletTrackMarker_VisualTrackID.delete(marker!);
      }
    })
  }

  refreshMapWithMarkers() {

    this.frontRepoService.pull(this.GONG__StackPath).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.manageLayers()
        this.manageDivIcons()
        this.manageMakers()
        this.manageCircles()
        this.manageVLines()
        this.manageTracks()

        // console.log("Map : " + this.mapName + ", length of root of leaflet layers: " + this.rootOfLayerGroups.length)
      }
    )
  }
}


export const LIMIT_HISTORY_LENGTH = 10;

const addNewCoordToFIFO = (
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

