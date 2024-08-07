// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongfly/go/models"
)

type GongstructDB interface {
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup
	case *models.CivilianAirport:
		civilianairportInstance := any(concreteInstance).(*models.CivilianAirport)
		ret2 := backRepo.BackRepoCivilianAirport.GetCivilianAirportDBFromCivilianAirportPtr(civilianairportInstance)
		ret = any(ret2).(*T2)
	case *models.Liner:
		linerInstance := any(concreteInstance).(*models.Liner)
		ret2 := backRepo.BackRepoLiner.GetLinerDBFromLinerPtr(linerInstance)
		ret = any(ret2).(*T2)
	case *models.Message:
		messageInstance := any(concreteInstance).(*models.Message)
		ret2 := backRepo.BackRepoMessage.GetMessageDBFromMessagePtr(messageInstance)
		ret = any(ret2).(*T2)
	case *models.OpsLine:
		opslineInstance := any(concreteInstance).(*models.OpsLine)
		ret2 := backRepo.BackRepoOpsLine.GetOpsLineDBFromOpsLinePtr(opslineInstance)
		ret = any(ret2).(*T2)
	case *models.Radar:
		radarInstance := any(concreteInstance).(*models.Radar)
		ret2 := backRepo.BackRepoRadar.GetRadarDBFromRadarPtr(radarInstance)
		ret = any(ret2).(*T2)
	case *models.Satellite:
		satelliteInstance := any(concreteInstance).(*models.Satellite)
		ret2 := backRepo.BackRepoSatellite.GetSatelliteDBFromSatellitePtr(satelliteInstance)
		ret = any(ret2).(*T2)
	case *models.Scenario:
		scenarioInstance := any(concreteInstance).(*models.Scenario)
		ret2 := backRepo.BackRepoScenario.GetScenarioDBFromScenarioPtr(scenarioInstance)
		ret = any(ret2).(*T2)
	default:
		_ = concreteInstance
	}
	return
}

func GetID[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.CivilianAirport:
		tmp := GetInstanceDBFromInstance[models.CivilianAirport, CivilianAirportDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Liner:
		tmp := GetInstanceDBFromInstance[models.Liner, LinerDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Message:
		tmp := GetInstanceDBFromInstance[models.Message, MessageDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.OpsLine:
		tmp := GetInstanceDBFromInstance[models.OpsLine, OpsLineDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Radar:
		tmp := GetInstanceDBFromInstance[models.Radar, RadarDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Satellite:
		tmp := GetInstanceDBFromInstance[models.Satellite, SatelliteDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Scenario:
		tmp := GetInstanceDBFromInstance[models.Scenario, ScenarioDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}

func GetIDPointer[T models.PointerToGongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.CivilianAirport:
		tmp := GetInstanceDBFromInstance[models.CivilianAirport, CivilianAirportDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Liner:
		tmp := GetInstanceDBFromInstance[models.Liner, LinerDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Message:
		tmp := GetInstanceDBFromInstance[models.Message, MessageDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.OpsLine:
		tmp := GetInstanceDBFromInstance[models.OpsLine, OpsLineDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Radar:
		tmp := GetInstanceDBFromInstance[models.Radar, RadarDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Satellite:
		tmp := GetInstanceDBFromInstance[models.Satellite, SatelliteDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Scenario:
		tmp := GetInstanceDBFromInstance[models.Scenario, ScenarioDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}
