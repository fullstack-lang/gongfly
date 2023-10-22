// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongfly/go/models"
)

func GetReverseFieldOwnerName[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.CivilianAirport:
		tmp := GetInstanceDBFromInstance[models.CivilianAirport, CivilianAirportDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "CivilianAirport":
			switch reverseField.Fieldname {
			}
		case "Liner":
			switch reverseField.Fieldname {
			}
		case "Message":
			switch reverseField.Fieldname {
			}
		case "OpsLine":
			switch reverseField.Fieldname {
			}
		case "Radar":
			switch reverseField.Fieldname {
			}
		case "Satellite":
			switch reverseField.Fieldname {
			}
		case "Scenario":
			switch reverseField.Fieldname {
			}
		}

	case *models.Liner:
		tmp := GetInstanceDBFromInstance[models.Liner, LinerDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "CivilianAirport":
			switch reverseField.Fieldname {
			}
		case "Liner":
			switch reverseField.Fieldname {
			}
		case "Message":
			switch reverseField.Fieldname {
			}
		case "OpsLine":
			switch reverseField.Fieldname {
			}
		case "Radar":
			switch reverseField.Fieldname {
			}
		case "Satellite":
			switch reverseField.Fieldname {
			}
		case "Scenario":
			switch reverseField.Fieldname {
			}
		}

	case *models.Message:
		tmp := GetInstanceDBFromInstance[models.Message, MessageDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "CivilianAirport":
			switch reverseField.Fieldname {
			}
		case "Liner":
			switch reverseField.Fieldname {
			}
		case "Message":
			switch reverseField.Fieldname {
			}
		case "OpsLine":
			switch reverseField.Fieldname {
			}
		case "Radar":
			switch reverseField.Fieldname {
			}
		case "Satellite":
			switch reverseField.Fieldname {
			}
		case "Scenario":
			switch reverseField.Fieldname {
			}
		}

	case *models.OpsLine:
		tmp := GetInstanceDBFromInstance[models.OpsLine, OpsLineDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "CivilianAirport":
			switch reverseField.Fieldname {
			}
		case "Liner":
			switch reverseField.Fieldname {
			}
		case "Message":
			switch reverseField.Fieldname {
			}
		case "OpsLine":
			switch reverseField.Fieldname {
			}
		case "Radar":
			switch reverseField.Fieldname {
			}
		case "Satellite":
			switch reverseField.Fieldname {
			}
		case "Scenario":
			switch reverseField.Fieldname {
			}
		}

	case *models.Radar:
		tmp := GetInstanceDBFromInstance[models.Radar, RadarDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "CivilianAirport":
			switch reverseField.Fieldname {
			}
		case "Liner":
			switch reverseField.Fieldname {
			}
		case "Message":
			switch reverseField.Fieldname {
			}
		case "OpsLine":
			switch reverseField.Fieldname {
			}
		case "Radar":
			switch reverseField.Fieldname {
			}
		case "Satellite":
			switch reverseField.Fieldname {
			}
		case "Scenario":
			switch reverseField.Fieldname {
			}
		}

	case *models.Satellite:
		tmp := GetInstanceDBFromInstance[models.Satellite, SatelliteDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "CivilianAirport":
			switch reverseField.Fieldname {
			}
		case "Liner":
			switch reverseField.Fieldname {
			}
		case "Message":
			switch reverseField.Fieldname {
			}
		case "OpsLine":
			switch reverseField.Fieldname {
			}
		case "Radar":
			switch reverseField.Fieldname {
			}
		case "Satellite":
			switch reverseField.Fieldname {
			}
		case "Scenario":
			switch reverseField.Fieldname {
			}
		}

	case *models.Scenario:
		tmp := GetInstanceDBFromInstance[models.Scenario, ScenarioDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "CivilianAirport":
			switch reverseField.Fieldname {
			}
		case "Liner":
			switch reverseField.Fieldname {
			}
		case "Message":
			switch reverseField.Fieldname {
			}
		case "OpsLine":
			switch reverseField.Fieldname {
			}
		case "Radar":
			switch reverseField.Fieldname {
			}
		case "Satellite":
			switch reverseField.Fieldname {
			}
		case "Scenario":
			switch reverseField.Fieldname {
			}
		}

	default:
		_ = inst
	}
	return
}

func GetReverseFieldOwner[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res any) {

	res = nil
	switch inst := any(instance).(type) {
	// insertion point
	case *models.CivilianAirport:
		tmp := GetInstanceDBFromInstance[models.CivilianAirport, CivilianAirportDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "CivilianAirport":
			switch reverseField.Fieldname {
			}
		case "Liner":
			switch reverseField.Fieldname {
			}
		case "Message":
			switch reverseField.Fieldname {
			}
		case "OpsLine":
			switch reverseField.Fieldname {
			}
		case "Radar":
			switch reverseField.Fieldname {
			}
		case "Satellite":
			switch reverseField.Fieldname {
			}
		case "Scenario":
			switch reverseField.Fieldname {
			}
		}

	case *models.Liner:
		tmp := GetInstanceDBFromInstance[models.Liner, LinerDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "CivilianAirport":
			switch reverseField.Fieldname {
			}
		case "Liner":
			switch reverseField.Fieldname {
			}
		case "Message":
			switch reverseField.Fieldname {
			}
		case "OpsLine":
			switch reverseField.Fieldname {
			}
		case "Radar":
			switch reverseField.Fieldname {
			}
		case "Satellite":
			switch reverseField.Fieldname {
			}
		case "Scenario":
			switch reverseField.Fieldname {
			}
		}

	case *models.Message:
		tmp := GetInstanceDBFromInstance[models.Message, MessageDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "CivilianAirport":
			switch reverseField.Fieldname {
			}
		case "Liner":
			switch reverseField.Fieldname {
			}
		case "Message":
			switch reverseField.Fieldname {
			}
		case "OpsLine":
			switch reverseField.Fieldname {
			}
		case "Radar":
			switch reverseField.Fieldname {
			}
		case "Satellite":
			switch reverseField.Fieldname {
			}
		case "Scenario":
			switch reverseField.Fieldname {
			}
		}

	case *models.OpsLine:
		tmp := GetInstanceDBFromInstance[models.OpsLine, OpsLineDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "CivilianAirport":
			switch reverseField.Fieldname {
			}
		case "Liner":
			switch reverseField.Fieldname {
			}
		case "Message":
			switch reverseField.Fieldname {
			}
		case "OpsLine":
			switch reverseField.Fieldname {
			}
		case "Radar":
			switch reverseField.Fieldname {
			}
		case "Satellite":
			switch reverseField.Fieldname {
			}
		case "Scenario":
			switch reverseField.Fieldname {
			}
		}

	case *models.Radar:
		tmp := GetInstanceDBFromInstance[models.Radar, RadarDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "CivilianAirport":
			switch reverseField.Fieldname {
			}
		case "Liner":
			switch reverseField.Fieldname {
			}
		case "Message":
			switch reverseField.Fieldname {
			}
		case "OpsLine":
			switch reverseField.Fieldname {
			}
		case "Radar":
			switch reverseField.Fieldname {
			}
		case "Satellite":
			switch reverseField.Fieldname {
			}
		case "Scenario":
			switch reverseField.Fieldname {
			}
		}

	case *models.Satellite:
		tmp := GetInstanceDBFromInstance[models.Satellite, SatelliteDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "CivilianAirport":
			switch reverseField.Fieldname {
			}
		case "Liner":
			switch reverseField.Fieldname {
			}
		case "Message":
			switch reverseField.Fieldname {
			}
		case "OpsLine":
			switch reverseField.Fieldname {
			}
		case "Radar":
			switch reverseField.Fieldname {
			}
		case "Satellite":
			switch reverseField.Fieldname {
			}
		case "Scenario":
			switch reverseField.Fieldname {
			}
		}

	case *models.Scenario:
		tmp := GetInstanceDBFromInstance[models.Scenario, ScenarioDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "CivilianAirport":
			switch reverseField.Fieldname {
			}
		case "Liner":
			switch reverseField.Fieldname {
			}
		case "Message":
			switch reverseField.Fieldname {
			}
		case "OpsLine":
			switch reverseField.Fieldname {
			}
		case "Radar":
			switch reverseField.Fieldname {
			}
		case "Satellite":
			switch reverseField.Fieldname {
			}
		case "Scenario":
			switch reverseField.Fieldname {
			}
		}

	default:
		_ = inst
	}
	return res
}
