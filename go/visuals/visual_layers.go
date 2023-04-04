package visuals

import (
	"github.com/fullstack-lang/gongfly/go/models"

	gongleaflet_models "github.com/fullstack-lang/gongleaflet/go/models"
	gongsim_models "github.com/fullstack-lang/gongsim/go/models"
)

var AircraftLayerGroup *gongleaflet_models.LayerGroup
var SatelliteLayerGroup *gongleaflet_models.LayerGroup
var NetworkLayerGroup *gongleaflet_models.LayerGroup
var CenterLayerGroup *gongleaflet_models.LayerGroup
var SystemLayerGroup *gongleaflet_models.LayerGroup

func LoadLayerGroups(stage *gongleaflet_models.StageStruct, engine *gongsim_models.Engine) {
	AircraftLayerGroup = (&gongleaflet_models.LayerGroup{
		Name: string(models.Aircraft_),
	}).Stage(stage)

	SatelliteLayerGroup = (&gongleaflet_models.LayerGroup{
		Name: string(models.Satellite_),
	}).Stage(stage)

	NetworkLayerGroup = (&gongleaflet_models.LayerGroup{
		Name: string(models.Network_),
	}).Stage(stage)

	CenterLayerGroup = (&gongleaflet_models.LayerGroup{
		Name: string(models.Center_),
	}).Stage(stage)

	SystemLayerGroup = (&gongleaflet_models.LayerGroup{
		Name: string(models.System_),
	}).Stage(stage)

}
