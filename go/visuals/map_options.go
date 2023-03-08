package visuals

import (
	gongfly_models "github.com/fullstack-lang/gongfly/go/models"
	gongfly_reference "github.com/fullstack-lang/gongfly/go/reference"
	gongleaflet_models "github.com/fullstack-lang/gongleaflet/go/models"
)

var MapOptions = (&gongleaflet_models.MapOptions{

	Name:               "Whole France",
	Lat:                gongfly_reference.Scenario1.Lat,
	Lng:                gongfly_reference.Scenario1.Lng,
	ZoomLevel:          gongfly_reference.Scenario1.ZoomLevel,
	ZoomControl:        false,
	AttributionControl: true,
	ZoomSnap:           1,
	UrlTemplate:        "https://{s}.basemaps.cartocdn.com/rastertiles/voyager/{z}/{x}/{y}{r}.png",
	Attribution:        "osm",
	MaxZoom:            18,
	LayerGroupUses: []*gongleaflet_models.LayerGroupUse{
		AircraftLayerGroupUse,
		CenterLayerGroupUse,
		SatelitteLayerGroupUse,
	},
}).Stage(&gongleaflet_models.Stage)

var AircraftLayerGroupUse = (&gongleaflet_models.LayerGroupUse{
	Name:       string(gongfly_models.Aircraft_),
	LayerGroup: AircraftLayerGroup,
	Display:    true,
}).Stage(&gongleaflet_models.Stage)

var SatelitteLayerGroupUse = (&gongleaflet_models.LayerGroupUse{
	Name:       string(gongfly_models.Satellite_),
	LayerGroup: SatelliteLayerGroup,
	Display:    true,
}).Stage(&gongleaflet_models.Stage)

var CenterLayerGroupUse = (&gongleaflet_models.LayerGroupUse{
	Name:       string(gongfly_models.Center_),
	LayerGroup: CenterLayerGroup,
	Display:    true,
}).Stage(&gongleaflet_models.Stage)
