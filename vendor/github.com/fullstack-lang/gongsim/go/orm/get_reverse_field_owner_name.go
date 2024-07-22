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
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Engine:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Event:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.GongsimCommand:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.GongsimStatus:
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
	case *models.DummyAgent:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Engine:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Event:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.GongsimCommand:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.GongsimStatus:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return res
}
