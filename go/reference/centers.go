package reference

import (
	"github.com/fullstack-lang/gongfly/go/models"
	gongsim_models "github.com/fullstack-lang/gongsim/go/models"
)

const visualShiftInMinutes = 15.0 / 60.0

var TLN_LFTH_ref = models.CivilianAirport{
	Name: "Hyères",
	CenterConcept: models.CenterConcept{
		Lat: 43 + 05/60.0,
		Lng: 6 + 8.0/60.0,
	},
}

var CDG_LFPG_ref = &models.CivilianAirport{
	Name: "CDG",
	CenterConcept: models.CenterConcept{
		Lat: 49 + 0.3/60.0,
		Lng: 2 + 32.0/60.0,
	},
}
var CDG_LFPG *models.CivilianAirport

func LoadCenters(stage *models.StageStruct, engine *gongsim_models.Engine) {
	CDG_LFPG = CDG_LFPG_ref.Stage(stage).Register(engine)
}
