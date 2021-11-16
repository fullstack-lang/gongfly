package visuals

import (
	"github.com/fullstack-lang/gongfly/go/models"

	gongleaflet_models "github.com/fullstack-lang/gongleaflet/go/models"
)

var AircraftLayerGroup = (&gongleaflet_models.LayerGroup{
	Name: string(models.Aircraft_),
}).StageCopy()

var SatelliteLayerGroup = (&gongleaflet_models.LayerGroup{
	Name: string(models.Satellite_),
}).StageCopy()

var NetworkLayerGroup = (&gongleaflet_models.LayerGroup{
	Name: string(models.Network_),
}).StageCopy()

var CenterLayerGroup = (&gongleaflet_models.LayerGroup{
	Name: string(models.Center_),
}).StageCopy()

var SystemLayerGroup = (&gongleaflet_models.LayerGroup{
	Name: string(models.System_),
}).StageCopy()
