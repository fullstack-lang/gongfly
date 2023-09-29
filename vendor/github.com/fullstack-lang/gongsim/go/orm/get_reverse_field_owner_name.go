// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongsim/go/models"
)

func GetReverseFieldOwnerName[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.DummyAgent:
		tmp := GetInstanceDBFromInstance[models.DummyAgent, DummyAgentDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "DummyAgent":
			switch reverseField.Fieldname {
			}
		case "Engine":
			switch reverseField.Fieldname {
			}
		case "Event":
			switch reverseField.Fieldname {
			}
		case "GongsimCommand":
			switch reverseField.Fieldname {
			}
		case "GongsimStatus":
			switch reverseField.Fieldname {
			}
		}

	case *models.Engine:
		tmp := GetInstanceDBFromInstance[models.Engine, EngineDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "DummyAgent":
			switch reverseField.Fieldname {
			}
		case "Engine":
			switch reverseField.Fieldname {
			}
		case "Event":
			switch reverseField.Fieldname {
			}
		case "GongsimCommand":
			switch reverseField.Fieldname {
			}
		case "GongsimStatus":
			switch reverseField.Fieldname {
			}
		}

	case *models.Event:
		tmp := GetInstanceDBFromInstance[models.Event, EventDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "DummyAgent":
			switch reverseField.Fieldname {
			}
		case "Engine":
			switch reverseField.Fieldname {
			}
		case "Event":
			switch reverseField.Fieldname {
			}
		case "GongsimCommand":
			switch reverseField.Fieldname {
			}
		case "GongsimStatus":
			switch reverseField.Fieldname {
			}
		}

	case *models.GongsimCommand:
		tmp := GetInstanceDBFromInstance[models.GongsimCommand, GongsimCommandDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "DummyAgent":
			switch reverseField.Fieldname {
			}
		case "Engine":
			switch reverseField.Fieldname {
			}
		case "Event":
			switch reverseField.Fieldname {
			}
		case "GongsimCommand":
			switch reverseField.Fieldname {
			}
		case "GongsimStatus":
			switch reverseField.Fieldname {
			}
		}

	case *models.GongsimStatus:
		tmp := GetInstanceDBFromInstance[models.GongsimStatus, GongsimStatusDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "DummyAgent":
			switch reverseField.Fieldname {
			}
		case "Engine":
			switch reverseField.Fieldname {
			}
		case "Event":
			switch reverseField.Fieldname {
			}
		case "GongsimCommand":
			switch reverseField.Fieldname {
			}
		case "GongsimStatus":
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
	case *models.DummyAgent:
		tmp := GetInstanceDBFromInstance[models.DummyAgent, DummyAgentDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "DummyAgent":
			switch reverseField.Fieldname {
			}
		case "Engine":
			switch reverseField.Fieldname {
			}
		case "Event":
			switch reverseField.Fieldname {
			}
		case "GongsimCommand":
			switch reverseField.Fieldname {
			}
		case "GongsimStatus":
			switch reverseField.Fieldname {
			}
		}
	
	case *models.Engine:
		tmp := GetInstanceDBFromInstance[models.Engine, EngineDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "DummyAgent":
			switch reverseField.Fieldname {
			}
		case "Engine":
			switch reverseField.Fieldname {
			}
		case "Event":
			switch reverseField.Fieldname {
			}
		case "GongsimCommand":
			switch reverseField.Fieldname {
			}
		case "GongsimStatus":
			switch reverseField.Fieldname {
			}
		}
	
	case *models.Event:
		tmp := GetInstanceDBFromInstance[models.Event, EventDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "DummyAgent":
			switch reverseField.Fieldname {
			}
		case "Engine":
			switch reverseField.Fieldname {
			}
		case "Event":
			switch reverseField.Fieldname {
			}
		case "GongsimCommand":
			switch reverseField.Fieldname {
			}
		case "GongsimStatus":
			switch reverseField.Fieldname {
			}
		}
	
	case *models.GongsimCommand:
		tmp := GetInstanceDBFromInstance[models.GongsimCommand, GongsimCommandDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "DummyAgent":
			switch reverseField.Fieldname {
			}
		case "Engine":
			switch reverseField.Fieldname {
			}
		case "Event":
			switch reverseField.Fieldname {
			}
		case "GongsimCommand":
			switch reverseField.Fieldname {
			}
		case "GongsimStatus":
			switch reverseField.Fieldname {
			}
		}
	
	case *models.GongsimStatus:
		tmp := GetInstanceDBFromInstance[models.GongsimStatus, GongsimStatusDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "DummyAgent":
			switch reverseField.Fieldname {
			}
		case "Engine":
			switch reverseField.Fieldname {
			}
		case "Event":
			switch reverseField.Fieldname {
			}
		case "GongsimCommand":
			switch reverseField.Fieldname {
			}
		case "GongsimStatus":
			switch reverseField.Fieldname {
			}
		}
	
	default:
		_ = inst
	}
	return res
}
