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

func LoadLayerGroups(gongleafletStage *gongleaflet_models.StageStruct, engine *gongsim_models.Engine) {
	AircraftLayerGroup = (&gongleaflet_models.LayerGroup{
		Name: string(models.Aircraft_),
	}).Stage(gongleafletStage)

	SatelliteLayerGroup = (&gongleaflet_models.LayerGroup{
		Name: string(models.Satellite_),
	}).Stage(gongleafletStage)

	NetworkLayerGroup = (&gongleaflet_models.LayerGroup{
		Name: string(models.Network_),
	}).Stage(gongleafletStage)

	CenterLayerGroup = (&gongleaflet_models.LayerGroup{
		Name: string(models.Center_),
	}).Stage(gongleafletStage)

	SystemLayerGroup = (&gongleaflet_models.LayerGroup{
		Name: string(models.System_),
	}).Stage(gongleafletStage)

}
