package visuals

import (
	"github.com/fullstack-lang/gongfly/go/models"

	gongleaflet_models "github.com/fullstack-lang/gongleaflet/go/models"
)

var AircraftVisualLayer = (&gongleaflet_models.VisualLayer{
	Name: string(models.Aircraft_),
}).StageCopy()

var NetworkVisualLayer = (&gongleaflet_models.VisualLayer{
	Name: string(models.Network_),
}).StageCopy()

var CenterVisualLayer = (&gongleaflet_models.VisualLayer{
	Name: string(models.Center_),
}).StageCopy()

var SystemVisualLayer = (&gongleaflet_models.VisualLayer{
	Name: string(models.System_),
}).StageCopy()
