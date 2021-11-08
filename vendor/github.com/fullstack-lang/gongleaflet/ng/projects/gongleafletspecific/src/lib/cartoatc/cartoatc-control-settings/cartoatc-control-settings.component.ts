import { Component, Input, OnInit } from '@angular/core';
import { MatSlideToggleChange } from '@angular/material/slide-toggle';
import { VisualLayerService } from 'gongleaflet';
import { setVisibilityHTMLElement } from '../manage-leaflet-items';

// https://stackoverflow.com/questions/54734329/ngx-leaflet-how-to-add-a-custom-control
@Component({
  selector: 'app-cartoatc-control-settings',
  templateUrl: './cartoatc-control-settings.component.html',
  styleUrls: ['./cartoatc-control-settings.component.scss'],
})
export class CartoatcControlSettingsComponent implements OnInit {
  list: Array<any> = [];
  open: boolean = false;

  constructor(private visualLayerService: VisualLayerService) {}

  ngOnInit(): void {
    this.visualLayerService.getVisualLayers().subscribe((visualLayers) => {
      visualLayers.forEach((visualLayer) => {
        this.list.push({
          id: visualLayer.ID,
          name: visualLayer.Name,
          display: visualLayer.DisplayName || visualLayer.Name,
          status: true,
        });
      });
    });
  }

  toggleOpen() {
    this.open = !this.open;
  }

  handleChange(change: MatSlideToggleChange, id: number) {
    let layerItems = document.getElementsByClassName('layer-' + id);
    for (let index = 0; index < layerItems.length; index++) {
      let htmlElement: Element | any = layerItems[index];
      setVisibilityHTMLElement(htmlElement, change.checked);
    }
  }
}
