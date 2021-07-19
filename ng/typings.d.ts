import * as L from 'leaflet';

// https://github.com/Asymmetrik/ngx-leaflet/issues/101
// ===== CREATED BY V. OMNÈS ===== //
// This declaration make leaflet-movingMarker
// compatible with typescripe 'typing'
declare module 'leaflet' {
    interface MarkerOptions {
        autostart?: boolean;
        loop?: boolean;
    }

    export interface Marker {
        movingMarker(latlngs: L.LatLng[], durations: Number[], options?: Object): Marker;
        addStation(pointIndex: number, duration: number): this;
        start(): this;
        stop(): this;
        isRunning(): this;
        isEnded(): this;
        isStarted(): this;
        isPaused(): this;
        resume(): this;
    }
    export interface MovingMarker extends Marker {
        movingMarker(latlngs: L.LatLng[], durations: Number[], options?: Object): MovingMarker;
    }
}