// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongleaflet/go/models"
)

func GetReverseFieldOwnerName[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Circle:
		tmp := GetInstanceDBFromInstance[models.Circle, CircleDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "DivIcon":
			switch reverseField.Fieldname {
			}
		case "LayerGroup":
			switch reverseField.Fieldname {
			}
		case "LayerGroupUse":
			switch reverseField.Fieldname {
			}
		case "MapOptions":
			switch reverseField.Fieldname {
			}
		case "Marker":
			switch reverseField.Fieldname {
			}
		case "UserClick":
			switch reverseField.Fieldname {
			}
		case "VLine":
			switch reverseField.Fieldname {
			}
		case "VisualTrack":
			switch reverseField.Fieldname {
			}
		}

	case *models.DivIcon:
		tmp := GetInstanceDBFromInstance[models.DivIcon, DivIconDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "DivIcon":
			switch reverseField.Fieldname {
			}
		case "LayerGroup":
			switch reverseField.Fieldname {
			}
		case "LayerGroupUse":
			switch reverseField.Fieldname {
			}
		case "MapOptions":
			switch reverseField.Fieldname {
			}
		case "Marker":
			switch reverseField.Fieldname {
			}
		case "UserClick":
			switch reverseField.Fieldname {
			}
		case "VLine":
			switch reverseField.Fieldname {
			}
		case "VisualTrack":
			switch reverseField.Fieldname {
			}
		}

	case *models.LayerGroup:
		tmp := GetInstanceDBFromInstance[models.LayerGroup, LayerGroupDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "DivIcon":
			switch reverseField.Fieldname {
			}
		case "LayerGroup":
			switch reverseField.Fieldname {
			}
		case "LayerGroupUse":
			switch reverseField.Fieldname {
			}
		case "MapOptions":
			switch reverseField.Fieldname {
			}
		case "Marker":
			switch reverseField.Fieldname {
			}
		case "UserClick":
			switch reverseField.Fieldname {
			}
		case "VLine":
			switch reverseField.Fieldname {
			}
		case "VisualTrack":
			switch reverseField.Fieldname {
			}
		}

	case *models.LayerGroupUse:
		tmp := GetInstanceDBFromInstance[models.LayerGroupUse, LayerGroupUseDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "DivIcon":
			switch reverseField.Fieldname {
			}
		case "LayerGroup":
			switch reverseField.Fieldname {
			}
		case "LayerGroupUse":
			switch reverseField.Fieldname {
			}
		case "MapOptions":
			switch reverseField.Fieldname {
			case "LayerGroupUses":
				if _mapoptions, ok := stage.MapOptions_LayerGroupUses_reverseMap[inst]; ok {
					res = _mapoptions.Name
				}
			}
		case "Marker":
			switch reverseField.Fieldname {
			}
		case "UserClick":
			switch reverseField.Fieldname {
			}
		case "VLine":
			switch reverseField.Fieldname {
			}
		case "VisualTrack":
			switch reverseField.Fieldname {
			}
		}

	case *models.MapOptions:
		tmp := GetInstanceDBFromInstance[models.MapOptions, MapOptionsDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "DivIcon":
			switch reverseField.Fieldname {
			}
		case "LayerGroup":
			switch reverseField.Fieldname {
			}
		case "LayerGroupUse":
			switch reverseField.Fieldname {
			}
		case "MapOptions":
			switch reverseField.Fieldname {
			}
		case "Marker":
			switch reverseField.Fieldname {
			}
		case "UserClick":
			switch reverseField.Fieldname {
			}
		case "VLine":
			switch reverseField.Fieldname {
			}
		case "VisualTrack":
			switch reverseField.Fieldname {
			}
		}

	case *models.Marker:
		tmp := GetInstanceDBFromInstance[models.Marker, MarkerDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "DivIcon":
			switch reverseField.Fieldname {
			}
		case "LayerGroup":
			switch reverseField.Fieldname {
			}
		case "LayerGroupUse":
			switch reverseField.Fieldname {
			}
		case "MapOptions":
			switch reverseField.Fieldname {
			}
		case "Marker":
			switch reverseField.Fieldname {
			}
		case "UserClick":
			switch reverseField.Fieldname {
			}
		case "VLine":
			switch reverseField.Fieldname {
			}
		case "VisualTrack":
			switch reverseField.Fieldname {
			}
		}

	case *models.UserClick:
		tmp := GetInstanceDBFromInstance[models.UserClick, UserClickDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "DivIcon":
			switch reverseField.Fieldname {
			}
		case "LayerGroup":
			switch reverseField.Fieldname {
			}
		case "LayerGroupUse":
			switch reverseField.Fieldname {
			}
		case "MapOptions":
			switch reverseField.Fieldname {
			}
		case "Marker":
			switch reverseField.Fieldname {
			}
		case "UserClick":
			switch reverseField.Fieldname {
			}
		case "VLine":
			switch reverseField.Fieldname {
			}
		case "VisualTrack":
			switch reverseField.Fieldname {
			}
		}

	case *models.VLine:
		tmp := GetInstanceDBFromInstance[models.VLine, VLineDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "DivIcon":
			switch reverseField.Fieldname {
			}
		case "LayerGroup":
			switch reverseField.Fieldname {
			}
		case "LayerGroupUse":
			switch reverseField.Fieldname {
			}
		case "MapOptions":
			switch reverseField.Fieldname {
			}
		case "Marker":
			switch reverseField.Fieldname {
			}
		case "UserClick":
			switch reverseField.Fieldname {
			}
		case "VLine":
			switch reverseField.Fieldname {
			}
		case "VisualTrack":
			switch reverseField.Fieldname {
			}
		}

	case *models.VisualTrack:
		tmp := GetInstanceDBFromInstance[models.VisualTrack, VisualTrackDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "DivIcon":
			switch reverseField.Fieldname {
			}
		case "LayerGroup":
			switch reverseField.Fieldname {
			}
		case "LayerGroupUse":
			switch reverseField.Fieldname {
			}
		case "MapOptions":
			switch reverseField.Fieldname {
			}
		case "Marker":
			switch reverseField.Fieldname {
			}
		case "UserClick":
			switch reverseField.Fieldname {
			}
		case "VLine":
			switch reverseField.Fieldname {
			}
		case "VisualTrack":
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
	case *models.Circle:
		tmp := GetInstanceDBFromInstance[models.Circle, CircleDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "DivIcon":
			switch reverseField.Fieldname {
			}
		case "LayerGroup":
			switch reverseField.Fieldname {
			}
		case "LayerGroupUse":
			switch reverseField.Fieldname {
			}
		case "MapOptions":
			switch reverseField.Fieldname {
			}
		case "Marker":
			switch reverseField.Fieldname {
			}
		case "UserClick":
			switch reverseField.Fieldname {
			}
		case "VLine":
			switch reverseField.Fieldname {
			}
		case "VisualTrack":
			switch reverseField.Fieldname {
			}
		}

	case *models.DivIcon:
		tmp := GetInstanceDBFromInstance[models.DivIcon, DivIconDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "DivIcon":
			switch reverseField.Fieldname {
			}
		case "LayerGroup":
			switch reverseField.Fieldname {
			}
		case "LayerGroupUse":
			switch reverseField.Fieldname {
			}
		case "MapOptions":
			switch reverseField.Fieldname {
			}
		case "Marker":
			switch reverseField.Fieldname {
			}
		case "UserClick":
			switch reverseField.Fieldname {
			}
		case "VLine":
			switch reverseField.Fieldname {
			}
		case "VisualTrack":
			switch reverseField.Fieldname {
			}
		}

	case *models.LayerGroup:
		tmp := GetInstanceDBFromInstance[models.LayerGroup, LayerGroupDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "DivIcon":
			switch reverseField.Fieldname {
			}
		case "LayerGroup":
			switch reverseField.Fieldname {
			}
		case "LayerGroupUse":
			switch reverseField.Fieldname {
			}
		case "MapOptions":
			switch reverseField.Fieldname {
			}
		case "Marker":
			switch reverseField.Fieldname {
			}
		case "UserClick":
			switch reverseField.Fieldname {
			}
		case "VLine":
			switch reverseField.Fieldname {
			}
		case "VisualTrack":
			switch reverseField.Fieldname {
			}
		}

	case *models.LayerGroupUse:
		tmp := GetInstanceDBFromInstance[models.LayerGroupUse, LayerGroupUseDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "DivIcon":
			switch reverseField.Fieldname {
			}
		case "LayerGroup":
			switch reverseField.Fieldname {
			}
		case "LayerGroupUse":
			switch reverseField.Fieldname {
			}
		case "MapOptions":
			switch reverseField.Fieldname {
			case "LayerGroupUses":
				res = stage.MapOptions_LayerGroupUses_reverseMap[inst]
			}
		case "Marker":
			switch reverseField.Fieldname {
			}
		case "UserClick":
			switch reverseField.Fieldname {
			}
		case "VLine":
			switch reverseField.Fieldname {
			}
		case "VisualTrack":
			switch reverseField.Fieldname {
			}
		}

	case *models.MapOptions:
		tmp := GetInstanceDBFromInstance[models.MapOptions, MapOptionsDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "DivIcon":
			switch reverseField.Fieldname {
			}
		case "LayerGroup":
			switch reverseField.Fieldname {
			}
		case "LayerGroupUse":
			switch reverseField.Fieldname {
			}
		case "MapOptions":
			switch reverseField.Fieldname {
			}
		case "Marker":
			switch reverseField.Fieldname {
			}
		case "UserClick":
			switch reverseField.Fieldname {
			}
		case "VLine":
			switch reverseField.Fieldname {
			}
		case "VisualTrack":
			switch reverseField.Fieldname {
			}
		}

	case *models.Marker:
		tmp := GetInstanceDBFromInstance[models.Marker, MarkerDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "DivIcon":
			switch reverseField.Fieldname {
			}
		case "LayerGroup":
			switch reverseField.Fieldname {
			}
		case "LayerGroupUse":
			switch reverseField.Fieldname {
			}
		case "MapOptions":
			switch reverseField.Fieldname {
			}
		case "Marker":
			switch reverseField.Fieldname {
			}
		case "UserClick":
			switch reverseField.Fieldname {
			}
		case "VLine":
			switch reverseField.Fieldname {
			}
		case "VisualTrack":
			switch reverseField.Fieldname {
			}
		}

	case *models.UserClick:
		tmp := GetInstanceDBFromInstance[models.UserClick, UserClickDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "DivIcon":
			switch reverseField.Fieldname {
			}
		case "LayerGroup":
			switch reverseField.Fieldname {
			}
		case "LayerGroupUse":
			switch reverseField.Fieldname {
			}
		case "MapOptions":
			switch reverseField.Fieldname {
			}
		case "Marker":
			switch reverseField.Fieldname {
			}
		case "UserClick":
			switch reverseField.Fieldname {
			}
		case "VLine":
			switch reverseField.Fieldname {
			}
		case "VisualTrack":
			switch reverseField.Fieldname {
			}
		}

	case *models.VLine:
		tmp := GetInstanceDBFromInstance[models.VLine, VLineDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "DivIcon":
			switch reverseField.Fieldname {
			}
		case "LayerGroup":
			switch reverseField.Fieldname {
			}
		case "LayerGroupUse":
			switch reverseField.Fieldname {
			}
		case "MapOptions":
			switch reverseField.Fieldname {
			}
		case "Marker":
			switch reverseField.Fieldname {
			}
		case "UserClick":
			switch reverseField.Fieldname {
			}
		case "VLine":
			switch reverseField.Fieldname {
			}
		case "VisualTrack":
			switch reverseField.Fieldname {
			}
		}

	case *models.VisualTrack:
		tmp := GetInstanceDBFromInstance[models.VisualTrack, VisualTrackDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "DivIcon":
			switch reverseField.Fieldname {
			}
		case "LayerGroup":
			switch reverseField.Fieldname {
			}
		case "LayerGroupUse":
			switch reverseField.Fieldname {
			}
		case "MapOptions":
			switch reverseField.Fieldname {
			}
		case "Marker":
			switch reverseField.Fieldname {
			}
		case "UserClick":
			switch reverseField.Fieldname {
			}
		case "VLine":
			switch reverseField.Fieldname {
			}
		case "VisualTrack":
			switch reverseField.Fieldname {
			}
		}

	default:
		_ = inst
	}
	return res
}
