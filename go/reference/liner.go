package reference

import (
	"github.com/fullstack-lang/gongfly/go/models"

	gongsim_models "github.com/fullstack-lang/gongsim/go/models"
)

var Sc1_AF_CDG_HYE_ref = models.Liner{
	Agent: gongsim_models.Agent{
		TechName: "AF CDG->HYE",
	},
	Name: "AF CDG->HYE",
	MovingObject: models.MovingObject{
		Lat:     CDG_LFPG_ref.Lat,
		Lng:     CDG_LFPG_ref.Lng + visualShiftInMinutes,
		Level:   220.0,
		Speed:   900.0,
		Heading: 90.0,
	},
	TargetHeading: 90.0,

	MaxRotationalSpeed: 0.5, // degrees / seconds
	TargetLocationLat:  TLN_LFTH_ref.Lat,
	TargetLocationLng:  TLN_LFTH_ref.Lng,
}

var Sc1_AF_3577_MDM *models.Liner

func LoadLiners(stage *models.StageStruct, engine *gongsim_models.Engine) {
	Sc1_AF_3577_MDM = Sc1_AF_CDG_HYE_ref.Stage(stage).Register(engine)
}
