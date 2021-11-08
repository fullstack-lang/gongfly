import { Type } from '@angular/core';
import {
  VisualCircleDB,
  VisualColorEnum,
  VisualLineDB,
  VisualMapDB,
} from 'gongleaflet';

import * as L from 'leaflet';

const setMapOptions = (vMap: VisualMapDB) => {
  if (!vMap) {
    return null;
  }
  return {
    layers: [
      L.tileLayer(vMap.UrlTemplate, {
        maxZoom: vMap.MaxZoom,
        attribution: vMap.Attribution,
      }),
    ],
    zoom: vMap.ZoomLevel,
    center: L.latLng(vMap.Lat, vMap.Lng),
    mapOptions: {
      zoomControl: vMap.ZoomControl,
      attributionControl: vMap.AttributionControl,
      zoomSnap: vMap.ZoomSnap,
    },
  };
};

const newIcon = (
  id: number | string,
  className: string,
  svg: string,
  size: number,
  color: string,
  label?: string,
  opacity?: number
) => {
  let content = '';
  if (!opacity) {
    opacity = 0.8;
  }
  if (label) {
    content = `
          <div class="pin__content" style="color: ${color}; opacity:${opacity}">
            <h3 id="${id}--label">${label}</h3>
          </div>
        `;
  }
  const html = `
  <div id="${id}--icon-marker">
    <div class="pin" id="${id}--icon" style="color: ${color}; opacity:0.8"/>
      ${svg?.length > 30 ? svg : `<img src='assets/icons/${svg}.svg'/>`}
    </div>
    ${content}
  </div>`;
  return L.divIcon({
    html,
    iconSize: [size, size],
    shadowSize: [0, 0], // size of the shadow
    shadowAnchor: [0, 0], // the same for the shadow
    popupAnchor: [-3, -76], // point from which the popup should open relative to the iconAnchor
    className,
  });
};

const rotateIcon = (targetID: string | number, orientation: number) => {
  let target = document.getElementById(`${targetID}--icon`);
  if (target) {
    target.style.transform = `rotate(${orientation}deg)`;
  }
};

const setIconLabel = (targetID: string | number, newLabel: string) => {
  let target = document.getElementById(`${targetID}--label`);
  if (target) {
    target.innerHTML = newLabel;
  }
};

function newMarkerWithIcon(lat: number, lng: number, icon: any): L.Marker {
  return L.marker([lat, lng], { icon });
};

const newCircle = (visualCircle: VisualCircleDB): L.Circle => {
  return L.circle([visualCircle.Lat, visualCircle.Lng], {
    className: 'layer-' + visualCircle.VisualLayerID.Int64,
    radius: visualCircle.Radius * 1000,
    color: getColor(visualCircle.VisualColorEnum),
    opacity: 0.2,
    dashArray: getDashStyle(visualCircle.DashStyleEnum),
    dashOffset: '0',
    fill: false,
  });
};

const setLine = (newVisualLineData: VisualLineDB): L.Polyline => {
  return new L.Polyline([])
    .setLatLngs([
      L.latLng(newVisualLineData.StartLat, newVisualLineData.StartLng),
      L.latLng(newVisualLineData.EndLat, newVisualLineData.EndLng),
    ])
    .setStyle({
      className: 'layer-' + newVisualLineData.VisualLayerID.Int64,
      weight: 2,
      dashArray: getDashStyle(newVisualLineData.DashStyleEnum),
      opacity: 0.5,
      color: getColor(newVisualLineData.VisualColorEnum),
    });
};

const getColor = (visualColorEnum: string): string => {
  var color = 'grey';
  switch (visualColorEnum) {
    case VisualColorEnum.RED:
      color = 'red';
      break;
    case VisualColorEnum.GREY:
      color = 'grey';
      break;
    case VisualColorEnum.GREEN:
      color = 'green';
      break;
    case VisualColorEnum.BLUE:
      color = 'blue';
      break;
    case VisualColorEnum.LIGHT_BROWN_8D6E63:
      color = '#8D6E63';
      break;
  }
  return color;
};


const getDashStyle = (dashStyleEnumValue: string): string => {
  const types = {
    FIVE_TWENTY: '5 20',
    FIVE_TEN: '5 10',
  };
  return types[dashStyleEnumValue as keyof typeof types];
};

const setVisibilityHTMLElement = (
  htmlElement: HTMLElement | any,
  visible: boolean
) => {
  if (!htmlElement) {
    return;
  }
  try {
    htmlElement.style.display = visible ? '' : 'none';
  } catch {
    htmlElement.setAttribute('display', visible ? '' : 'none');
  }
};

export {
  setMapOptions,
  newIcon,
  rotateIcon,
  setIconLabel,
  newMarkerWithIcon,
  newCircle,
  setLine,
  getColor,
  getDashStyle,
  setVisibilityHTMLElement,
};
