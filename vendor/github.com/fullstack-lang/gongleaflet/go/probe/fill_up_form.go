// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongleaflet/go/models"
	"github.com/fullstack-lang/gongleaflet/go/orm"
)

var __dummy_orm_fillup_form = orm.BackRepoStruct{}

func FillUpForm[T models.Gongstruct](
	instance *T,
	formGroup *form.FormGroup,
	playground *Playground,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Circle:
		// insertion point
		BasicFieldtoForm("Lat", instanceWithInferedType.Lat, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Lng", instanceWithInferedType.Lng, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Radius", instanceWithInferedType.Radius, instanceWithInferedType, playground.formStage, formGroup, false)
		EnumTypeStringToForm("ColorEnum", instanceWithInferedType.ColorEnum, instanceWithInferedType, playground.formStage, formGroup)
		EnumTypeStringToForm("DashStyleEnum", instanceWithInferedType.DashStyleEnum, instanceWithInferedType, playground.formStage, formGroup)
		AssociationFieldToForm("LayerGroup", instanceWithInferedType.LayerGroup, formGroup, playground)

	case *models.DivIcon:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("SVG", instanceWithInferedType.SVG, instanceWithInferedType, playground.formStage, formGroup, false)

	case *models.LayerGroup:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("DisplayName", instanceWithInferedType.DisplayName, instanceWithInferedType, playground.formStage, formGroup, false)

	case *models.LayerGroupUse:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("IsDisplayed", instanceWithInferedType.IsDisplayed, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationFieldToForm("LayerGroup", instanceWithInferedType.LayerGroup, formGroup, playground)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "MapOptions"
			rf.Fieldname = "LayerGroupUses"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.MapOptions),
					"LayerGroupUses",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.MapOptions, *models.LayerGroupUse](
					nil,
					"LayerGroupUses",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.MapOptions:
		// insertion point
		BasicFieldtoForm("Lat", instanceWithInferedType.Lat, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Lng", instanceWithInferedType.Lng, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("ZoomLevel", instanceWithInferedType.ZoomLevel, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("UrlTemplate", instanceWithInferedType.UrlTemplate, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Attribution", instanceWithInferedType.Attribution, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("MaxZoom", instanceWithInferedType.MaxZoom, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("ZoomControl", instanceWithInferedType.ZoomControl, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("AttributionControl", instanceWithInferedType.AttributionControl, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("ZoomSnap", instanceWithInferedType.ZoomSnap, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("LayerGroupUses", instanceWithInferedType, &instanceWithInferedType.LayerGroupUses, formGroup, playground)

	case *models.Marker:
		// insertion point
		BasicFieldtoForm("Lat", instanceWithInferedType.Lat, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Lng", instanceWithInferedType.Lng, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		EnumTypeStringToForm("ColorEnum", instanceWithInferedType.ColorEnum, instanceWithInferedType, playground.formStage, formGroup)
		AssociationFieldToForm("LayerGroup", instanceWithInferedType.LayerGroup, formGroup, playground)
		AssociationFieldToForm("DivIcon", instanceWithInferedType.DivIcon, formGroup, playground)

	case *models.UserClick:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Lat", instanceWithInferedType.Lat, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Lng", instanceWithInferedType.Lng, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("TimeOfClick", instanceWithInferedType.TimeOfClick, instanceWithInferedType, playground.formStage, formGroup, false)

	case *models.VLine:
		// insertion point
		BasicFieldtoForm("StartLat", instanceWithInferedType.StartLat, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("StartLng", instanceWithInferedType.StartLng, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("EndLat", instanceWithInferedType.EndLat, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("EndLng", instanceWithInferedType.EndLng, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		EnumTypeStringToForm("ColorEnum", instanceWithInferedType.ColorEnum, instanceWithInferedType, playground.formStage, formGroup)
		EnumTypeStringToForm("DashStyleEnum", instanceWithInferedType.DashStyleEnum, instanceWithInferedType, playground.formStage, formGroup)
		AssociationFieldToForm("LayerGroup", instanceWithInferedType.LayerGroup, formGroup, playground)
		EnumTypeStringToForm("IsTransmitting", instanceWithInferedType.IsTransmitting, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Message", instanceWithInferedType.Message, instanceWithInferedType, playground.formStage, formGroup, false)
		EnumTypeStringToForm("IsTransmittingBackward", instanceWithInferedType.IsTransmittingBackward, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("MessageBackward", instanceWithInferedType.MessageBackward, instanceWithInferedType, playground.formStage, formGroup, false)

	case *models.VisualTrack:
		// insertion point
		BasicFieldtoForm("Lat", instanceWithInferedType.Lat, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Lng", instanceWithInferedType.Lng, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Heading", instanceWithInferedType.Heading, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Level", instanceWithInferedType.Level, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Speed", instanceWithInferedType.Speed, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("VerticalSpeed", instanceWithInferedType.VerticalSpeed, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		EnumTypeStringToForm("ColorEnum", instanceWithInferedType.ColorEnum, instanceWithInferedType, playground.formStage, formGroup)
		AssociationFieldToForm("LayerGroup", instanceWithInferedType.LayerGroup, formGroup, playground)
		AssociationFieldToForm("DivIcon", instanceWithInferedType.DivIcon, formGroup, playground)
		BasicFieldtoForm("DisplayTrackHistory", instanceWithInferedType.DisplayTrackHistory, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("DisplayLevelAndSpeed", instanceWithInferedType.DisplayLevelAndSpeed, instanceWithInferedType, playground.formStage, formGroup, false)

	default:
		_ = instanceWithInferedType
	}
}
