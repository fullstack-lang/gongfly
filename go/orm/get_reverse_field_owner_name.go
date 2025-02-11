// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongfly/go/models"
)

func GetReverseFieldOwnerName(
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance any,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.CivilianAirport:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Liner:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Message:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.OpsLine:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Radar:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Satellite:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Scenario:
		switch reverseField.GongstructName {
		// insertion point
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
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Liner:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Message:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.OpsLine:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Radar:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Satellite:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Scenario:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return res
}
