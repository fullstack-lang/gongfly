// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongfly/go/models"
	"github.com/fullstack-lang/gongfly/go/orm"
)

var __dummy_orm_fillup_form = orm.BackRepoStruct{}

func FillUpForm[T models.Gongstruct](
	instance *T,
	formGroup *form.FormGroup,
	playground *Playground,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.CivilianAirport:
		// insertion point
		BasicFieldtoForm("Lat", instanceWithInferedType.Lat, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Lng", instanceWithInferedType.Lng, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)

	case *models.Liner:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Lat", instanceWithInferedType.Lat, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Lng", instanceWithInferedType.Lng, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Heading", instanceWithInferedType.Heading, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Level", instanceWithInferedType.Level, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Speed", instanceWithInferedType.Speed, instanceWithInferedType, playground.formStage, formGroup, false)
		EnumTypeStringToForm("State", instanceWithInferedType.State, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("TargetHeading", instanceWithInferedType.TargetHeading, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("TargetLocationLat", instanceWithInferedType.TargetLocationLat, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("TargetLocationLng", instanceWithInferedType.TargetLocationLng, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("DistanceToTarget", instanceWithInferedType.DistanceToTarget, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("MaxRotationalSpeed", instanceWithInferedType.MaxRotationalSpeed, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("VerticalSpeed", instanceWithInferedType.VerticalSpeed, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Timestampstring", instanceWithInferedType.Timestampstring, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationFieldToForm("ReporingLine", instanceWithInferedType.ReporingLine, formGroup, playground)

	case *models.Message:
		// insertion point
		BasicFieldtoForm("Lat", instanceWithInferedType.Lat, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Lng", instanceWithInferedType.Lng, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Heading", instanceWithInferedType.Heading, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Level", instanceWithInferedType.Level, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Speed", instanceWithInferedType.Speed, instanceWithInferedType, playground.formStage, formGroup, false)
		EnumTypeStringToForm("State", instanceWithInferedType.State, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("TargetLocationLat", instanceWithInferedType.TargetLocationLat, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("TargetLocationLng", instanceWithInferedType.TargetLocationLng, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("DistanceToTarget", instanceWithInferedType.DistanceToTarget, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Timestampstring", instanceWithInferedType.Timestampstring, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("DurationSinceSimulationStart", instanceWithInferedType.DurationSinceSimulationStart, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Timestampstartstring", instanceWithInferedType.Timestampstartstring, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Source", instanceWithInferedType.Source, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Destination", instanceWithInferedType.Destination, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("About_string", instanceWithInferedType.About_string, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Display", instanceWithInferedType.Display, instanceWithInferedType, playground.formStage, formGroup, false)

	case *models.OpsLine:
		// insertion point
		BasicFieldtoForm("IsTransmitting", instanceWithInferedType.IsTransmitting, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("TransmissionMessage", instanceWithInferedType.TransmissionMessage, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("IsTransmittingBackward", instanceWithInferedType.IsTransmittingBackward, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("TransmissionMessageBackward", instanceWithInferedType.TransmissionMessageBackward, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationFieldToForm("Scenario", instanceWithInferedType.Scenario, formGroup, playground)
		EnumTypeStringToForm("State", instanceWithInferedType.State, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)

	case *models.Radar:
		// insertion point
		EnumTypeStringToForm("State", instanceWithInferedType.State, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Lat", instanceWithInferedType.Lat, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Lng", instanceWithInferedType.Lng, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Range", instanceWithInferedType.Range, instanceWithInferedType, playground.formStage, formGroup, false)

	case *models.Satellite:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Line1", instanceWithInferedType.Line1, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Line2", instanceWithInferedType.Line2, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Lat", instanceWithInferedType.Lat, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Lng", instanceWithInferedType.Lng, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Heading", instanceWithInferedType.Heading, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Level", instanceWithInferedType.Level, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Speed", instanceWithInferedType.Speed, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("VerticalSpeed", instanceWithInferedType.VerticalSpeed, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Timestampstring", instanceWithInferedType.Timestampstring, instanceWithInferedType, playground.formStage, formGroup, false)

	case *models.Scenario:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Lat", instanceWithInferedType.Lat, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Lng", instanceWithInferedType.Lng, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("ZoomLevel", instanceWithInferedType.ZoomLevel, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("MessageVisualSpeed", instanceWithInferedType.MessageVisualSpeed, instanceWithInferedType, playground.formStage, formGroup, false)

	default:
		_ = instanceWithInferedType
	}
}
