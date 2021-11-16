package reference

import (
	"github.com/fullstack-lang/gongfly/go/models"

	gongsim_models "github.com/fullstack-lang/gongsim/go/models"
)

var PropagationTestCase1_ref = models.Satellite{
	Agent: gongsim_models.Agent{
		TechName: "PropagationTestCase1",
	},

	// from go-satellite test suite
	Line1: "1 00005U 58002B   00179.78495062  .00000023  00000-0  28098-4 0  4753",
	Line2: "2 00005  34.2682 348.7242 1859667 331.7664  19.3264 10.82419157413667",

	Name: "Sat 1",
	MovingObject: models.MovingObject{
		Lat:     45,
		Lng:     3,
		Level:   220.0,
		Speed:   900.0,
		Heading: 90.0,
	},
	TargetHeading: 90.0,

	MaxRotationalSpeed: 0.5, // degrees / seconds
	TargetLocationLat:  TLN_LFTH_ref.Lat,
	TargetLocationLng:  TLN_LFTH_ref.Lng,
}
var PropagationTestCase1 = PropagationTestCase1_ref.InitFromTLE().StageCopy().Register()
