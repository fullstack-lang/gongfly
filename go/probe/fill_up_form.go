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
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.CivilianAirport:
		// insertion point
		BasicFieldtoForm("Lat", instanceWithInferedType.Lat, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("Lng", instanceWithInferedType.Lng, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup, false)

	case *models.Liner:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("Lat", instanceWithInferedType.Lat, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("Lng", instanceWithInferedType.Lng, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("Heading", instanceWithInferedType.Heading, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("Level", instanceWithInferedType.Level, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("Speed", instanceWithInferedType.Speed, instanceWithInferedType, probe.formStage, formGroup, false)
		EnumTypeStringToForm("State", instanceWithInferedType.State, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("TargetHeading", instanceWithInferedType.TargetHeading, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("TargetLocationLat", instanceWithInferedType.TargetLocationLat, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("TargetLocationLng", instanceWithInferedType.TargetLocationLng, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("DistanceToTarget", instanceWithInferedType.DistanceToTarget, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("MaxRotationalSpeed", instanceWithInferedType.MaxRotationalSpeed, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("VerticalSpeed", instanceWithInferedType.VerticalSpeed, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("Timestampstring", instanceWithInferedType.Timestampstring, instanceWithInferedType, probe.formStage, formGroup, false)
		AssociationFieldToForm("ReporingLine", instanceWithInferedType.ReporingLine, formGroup, probe)

	case *models.Message:
		// insertion point
		BasicFieldtoForm("Lat", instanceWithInferedType.Lat, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("Lng", instanceWithInferedType.Lng, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("Heading", instanceWithInferedType.Heading, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("Level", instanceWithInferedType.Level, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("Speed", instanceWithInferedType.Speed, instanceWithInferedType, probe.formStage, formGroup, false)
		EnumTypeStringToForm("State", instanceWithInferedType.State, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("TargetLocationLat", instanceWithInferedType.TargetLocationLat, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("TargetLocationLng", instanceWithInferedType.TargetLocationLng, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("DistanceToTarget", instanceWithInferedType.DistanceToTarget, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("Timestampstring", instanceWithInferedType.Timestampstring, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("DurationSinceSimulationStart", instanceWithInferedType.DurationSinceSimulationStart, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("Timestampstartstring", instanceWithInferedType.Timestampstartstring, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("Source", instanceWithInferedType.Source, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("Destination", instanceWithInferedType.Destination, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("About_string", instanceWithInferedType.About_string, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("Display", instanceWithInferedType.Display, instanceWithInferedType, probe.formStage, formGroup, false)

	case *models.OpsLine:
		// insertion point
		BasicFieldtoForm("IsTransmitting", instanceWithInferedType.IsTransmitting, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("TransmissionMessage", instanceWithInferedType.TransmissionMessage, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("IsTransmittingBackward", instanceWithInferedType.IsTransmittingBackward, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("TransmissionMessageBackward", instanceWithInferedType.TransmissionMessageBackward, instanceWithInferedType, probe.formStage, formGroup, false)
		AssociationFieldToForm("Scenario", instanceWithInferedType.Scenario, formGroup, probe)
		EnumTypeStringToForm("State", instanceWithInferedType.State, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup, false)

	case *models.Radar:
		// insertion point
		EnumTypeStringToForm("State", instanceWithInferedType.State, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("Lat", instanceWithInferedType.Lat, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("Lng", instanceWithInferedType.Lng, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("Range", instanceWithInferedType.Range, instanceWithInferedType, probe.formStage, formGroup, false)

	case *models.Satellite:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("Line1", instanceWithInferedType.Line1, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("Line2", instanceWithInferedType.Line2, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("Lat", instanceWithInferedType.Lat, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("Lng", instanceWithInferedType.Lng, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("Heading", instanceWithInferedType.Heading, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("Level", instanceWithInferedType.Level, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("Speed", instanceWithInferedType.Speed, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("VerticalSpeed", instanceWithInferedType.VerticalSpeed, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("Timestampstring", instanceWithInferedType.Timestampstring, instanceWithInferedType, probe.formStage, formGroup, false)

	case *models.Scenario:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("Lat", instanceWithInferedType.Lat, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("Lng", instanceWithInferedType.Lng, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("ZoomLevel", instanceWithInferedType.ZoomLevel, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("MessageVisualSpeed", instanceWithInferedType.MessageVisualSpeed, instanceWithInferedType, probe.formStage, formGroup, false)

	default:
		_ = instanceWithInferedType
	}
}
