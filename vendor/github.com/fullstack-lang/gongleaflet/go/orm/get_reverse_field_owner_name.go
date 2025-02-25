// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongleaflet/go/models"
)

func GetReverseFieldOwnerName(
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance any,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Circle:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.DivIcon:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.LayerGroup:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.LayerGroupUse:
		switch reverseField.GongstructName {
		// insertion point
		case "MapOptions":
			switch reverseField.Fieldname {
			case "LayerGroupUses":
				if _mapoptions, ok := stage.MapOptions_LayerGroupUses_reverseMap[inst]; ok {
					res = _mapoptions.Name
				}
			}
		}

	case *models.MapOptions:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Marker:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.UserClick:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.VLine:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.VisualTrack:
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
	case *models.Circle:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.DivIcon:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.LayerGroup:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.LayerGroupUse:
		switch reverseField.GongstructName {
		// insertion point
		case "MapOptions":
			switch reverseField.Fieldname {
			case "LayerGroupUses":
				res = stage.MapOptions_LayerGroupUses_reverseMap[inst]
			}
		}

	case *models.MapOptions:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Marker:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.UserClick:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.VLine:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.VisualTrack:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return res
}
