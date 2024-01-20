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
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Circle:
		// insertion point
		BasicFieldtoForm("Lat", instanceWithInferedType.Lat, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Lng", instanceWithInferedType.Lng, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Radius", instanceWithInferedType.Radius, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("ColorEnum", instanceWithInferedType.ColorEnum, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("DashStyleEnum", instanceWithInferedType.DashStyleEnum, instanceWithInferedType, probe.formStage, formGroup)
		AssociationFieldToForm("LayerGroup", instanceWithInferedType.LayerGroup, formGroup, probe)

	case *models.DivIcon:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("SVG", instanceWithInferedType.SVG, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.LayerGroup:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DisplayName", instanceWithInferedType.DisplayName, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.LayerGroupUse:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsDisplayed", instanceWithInferedType.IsDisplayed, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("LayerGroup", instanceWithInferedType.LayerGroup, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "MapOptions"
			rf.Fieldname = "LayerGroupUses"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.MapOptions),
					"LayerGroupUses",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.MapOptions, *models.LayerGroupUse](
					nil,
					"LayerGroupUses",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.MapOptions:
		// insertion point
		BasicFieldtoForm("Lat", instanceWithInferedType.Lat, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Lng", instanceWithInferedType.Lng, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ZoomLevel", instanceWithInferedType.ZoomLevel, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("UrlTemplate", instanceWithInferedType.UrlTemplate, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Attribution", instanceWithInferedType.Attribution, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MaxZoom", instanceWithInferedType.MaxZoom, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ZoomControl", instanceWithInferedType.ZoomControl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("AttributionControl", instanceWithInferedType.AttributionControl, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ZoomSnap", instanceWithInferedType.ZoomSnap, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("LayerGroupUses", instanceWithInferedType, &instanceWithInferedType.LayerGroupUses, formGroup, probe)

	case *models.Marker:
		// insertion point
		BasicFieldtoForm("Lat", instanceWithInferedType.Lat, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Lng", instanceWithInferedType.Lng, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("ColorEnum", instanceWithInferedType.ColorEnum, instanceWithInferedType, probe.formStage, formGroup)
		AssociationFieldToForm("LayerGroup", instanceWithInferedType.LayerGroup, formGroup, probe)
		AssociationFieldToForm("DivIcon", instanceWithInferedType.DivIcon, formGroup, probe)

	case *models.UserClick:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Lat", instanceWithInferedType.Lat, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Lng", instanceWithInferedType.Lng, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("TimeOfClick", instanceWithInferedType.TimeOfClick, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.VLine:
		// insertion point
		BasicFieldtoForm("StartLat", instanceWithInferedType.StartLat, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StartLng", instanceWithInferedType.StartLng, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EndLat", instanceWithInferedType.EndLat, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EndLng", instanceWithInferedType.EndLng, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("ColorEnum", instanceWithInferedType.ColorEnum, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("DashStyleEnum", instanceWithInferedType.DashStyleEnum, instanceWithInferedType, probe.formStage, formGroup)
		AssociationFieldToForm("LayerGroup", instanceWithInferedType.LayerGroup, formGroup, probe)
		EnumTypeStringToForm("IsTransmitting", instanceWithInferedType.IsTransmitting, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Message", instanceWithInferedType.Message, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("IsTransmittingBackward", instanceWithInferedType.IsTransmittingBackward, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("MessageBackward", instanceWithInferedType.MessageBackward, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.VisualTrack:
		// insertion point
		BasicFieldtoForm("Lat", instanceWithInferedType.Lat, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Lng", instanceWithInferedType.Lng, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Heading", instanceWithInferedType.Heading, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Level", instanceWithInferedType.Level, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Speed", instanceWithInferedType.Speed, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("VerticalSpeed", instanceWithInferedType.VerticalSpeed, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("ColorEnum", instanceWithInferedType.ColorEnum, instanceWithInferedType, probe.formStage, formGroup)
		AssociationFieldToForm("LayerGroup", instanceWithInferedType.LayerGroup, formGroup, probe)
		AssociationFieldToForm("DivIcon", instanceWithInferedType.DivIcon, formGroup, probe)
		BasicFieldtoForm("DisplayTrackHistory", instanceWithInferedType.DisplayTrackHistory, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DisplayLevelAndSpeed", instanceWithInferedType.DisplayLevelAndSpeed, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	default:
		_ = instanceWithInferedType
	}
}
