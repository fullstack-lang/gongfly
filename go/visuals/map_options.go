package visuals

import (
	gongfly_models "github.com/fullstack-lang/gongfly/go/models"
	gongfly_reference "github.com/fullstack-lang/gongfly/go/reference"
	gongleaflet_models "github.com/fullstack-lang/gongleaflet/go/models"
)

var MapOptions *gongleaflet_models.MapOptions

var AircraftLayerGroupUse *gongleaflet_models.LayerGroupUse
var SatelitteLayerGroupUse *gongleaflet_models.LayerGroupUse
var CenterLayerGroupUse *gongleaflet_models.LayerGroupUse

func LoadMapOptions(stage *gongleaflet_models.StageStruct) {

	MapOptions = &(gongleaflet_models.MapOptions{

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
	})

	MapOptions.Stage(stage)

}

func LoadLayerGroupsUse(stage *gongleaflet_models.StageStruct) {

	AircraftLayerGroupUse = (&gongleaflet_models.LayerGroupUse{
		Name:        string(gongfly_models.Aircraft_),
		LayerGroup:  AircraftLayerGroup,
		IsDisplayed: true,
	})

	SatelitteLayerGroupUse = (&gongleaflet_models.LayerGroupUse{
		Name:        string(gongfly_models.Satellite_),
		LayerGroup:  SatelliteLayerGroup,
		IsDisplayed: true,
	})

	CenterLayerGroupUse = (&gongleaflet_models.LayerGroupUse{
		Name:        string(gongfly_models.Center_),
		LayerGroup:  CenterLayerGroup,
		IsDisplayed: true,
	})

	AircraftLayerGroupUse.Stage(stage)
	SatelitteLayerGroupUse.Stage(stage)
	CenterLayerGroupUse.Stage(stage)
}
