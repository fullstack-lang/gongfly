// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongleaflet/go/models"
)

type GongstructDB interface {
	// insertion point for generic types
	// "int" is present to handle the case when no struct is present
	int | CircleDB | DivIconDB | LayerGroupDB | LayerGroupUseDB | MapOptionsDB | MarkerDB | UserClickDB | VLineDB | VisualTrackDB
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Circle:
		circleInstance := any(concreteInstance).(*models.Circle)
		ret2 := backRepo.BackRepoCircle.GetCircleDBFromCirclePtr(circleInstance)
		ret = any(ret2).(*T2)
	case *models.DivIcon:
		diviconInstance := any(concreteInstance).(*models.DivIcon)
		ret2 := backRepo.BackRepoDivIcon.GetDivIconDBFromDivIconPtr(diviconInstance)
		ret = any(ret2).(*T2)
	case *models.LayerGroup:
		layergroupInstance := any(concreteInstance).(*models.LayerGroup)
		ret2 := backRepo.BackRepoLayerGroup.GetLayerGroupDBFromLayerGroupPtr(layergroupInstance)
		ret = any(ret2).(*T2)
	case *models.LayerGroupUse:
		layergroupuseInstance := any(concreteInstance).(*models.LayerGroupUse)
		ret2 := backRepo.BackRepoLayerGroupUse.GetLayerGroupUseDBFromLayerGroupUsePtr(layergroupuseInstance)
		ret = any(ret2).(*T2)
	case *models.MapOptions:
		mapoptionsInstance := any(concreteInstance).(*models.MapOptions)
		ret2 := backRepo.BackRepoMapOptions.GetMapOptionsDBFromMapOptionsPtr(mapoptionsInstance)
		ret = any(ret2).(*T2)
	case *models.Marker:
		markerInstance := any(concreteInstance).(*models.Marker)
		ret2 := backRepo.BackRepoMarker.GetMarkerDBFromMarkerPtr(markerInstance)
		ret = any(ret2).(*T2)
	case *models.UserClick:
		userclickInstance := any(concreteInstance).(*models.UserClick)
		ret2 := backRepo.BackRepoUserClick.GetUserClickDBFromUserClickPtr(userclickInstance)
		ret = any(ret2).(*T2)
	case *models.VLine:
		vlineInstance := any(concreteInstance).(*models.VLine)
		ret2 := backRepo.BackRepoVLine.GetVLineDBFromVLinePtr(vlineInstance)
		ret = any(ret2).(*T2)
	case *models.VisualTrack:
		visualtrackInstance := any(concreteInstance).(*models.VisualTrack)
		ret2 := backRepo.BackRepoVisualTrack.GetVisualTrackDBFromVisualTrackPtr(visualtrackInstance)
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
	case *models.Circle:
		tmp := GetInstanceDBFromInstance[models.Circle, CircleDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DivIcon:
		tmp := GetInstanceDBFromInstance[models.DivIcon, DivIconDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.LayerGroup:
		tmp := GetInstanceDBFromInstance[models.LayerGroup, LayerGroupDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.LayerGroupUse:
		tmp := GetInstanceDBFromInstance[models.LayerGroupUse, LayerGroupUseDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.MapOptions:
		tmp := GetInstanceDBFromInstance[models.MapOptions, MapOptionsDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Marker:
		tmp := GetInstanceDBFromInstance[models.Marker, MarkerDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.UserClick:
		tmp := GetInstanceDBFromInstance[models.UserClick, UserClickDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.VLine:
		tmp := GetInstanceDBFromInstance[models.VLine, VLineDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.VisualTrack:
		tmp := GetInstanceDBFromInstance[models.VisualTrack, VisualTrackDB](
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
	case *models.Circle:
		tmp := GetInstanceDBFromInstance[models.Circle, CircleDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DivIcon:
		tmp := GetInstanceDBFromInstance[models.DivIcon, DivIconDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.LayerGroup:
		tmp := GetInstanceDBFromInstance[models.LayerGroup, LayerGroupDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.LayerGroupUse:
		tmp := GetInstanceDBFromInstance[models.LayerGroupUse, LayerGroupUseDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.MapOptions:
		tmp := GetInstanceDBFromInstance[models.MapOptions, MapOptionsDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Marker:
		tmp := GetInstanceDBFromInstance[models.Marker, MarkerDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.UserClick:
		tmp := GetInstanceDBFromInstance[models.UserClick, UserClickDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.VLine:
		tmp := GetInstanceDBFromInstance[models.VLine, VLineDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.VisualTrack:
		tmp := GetInstanceDBFromInstance[models.VisualTrack, VisualTrackDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}
