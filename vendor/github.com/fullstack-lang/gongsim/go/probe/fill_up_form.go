// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongsim/go/models"
	"github.com/fullstack-lang/gongsim/go/orm"
)

var __dummy_orm_fillup_form = orm.BackRepoStruct{}

func FillUpForm[T models.Gongstruct](
	instance *T,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.DummyAgent:
		// insertion point
		BasicFieldtoForm("TechName", instanceWithInferedType.TechName, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup, false)

	case *models.Engine:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("EndTime", instanceWithInferedType.EndTime, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("CurrentTime", instanceWithInferedType.CurrentTime, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("SecondsSinceStart", instanceWithInferedType.SecondsSinceStart, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("Fired", instanceWithInferedType.Fired, instanceWithInferedType, probe.formStage, formGroup, false)
		EnumTypeStringToForm("ControlMode", instanceWithInferedType.ControlMode, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("State", instanceWithInferedType.State, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Speed", instanceWithInferedType.Speed, instanceWithInferedType, probe.formStage, formGroup, false)

	case *models.Event:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("Duration", instanceWithInferedType.Duration, instanceWithInferedType, probe.formStage, formGroup, false)

	case *models.GongsimCommand:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup, false)
		EnumTypeStringToForm("Command", instanceWithInferedType.Command, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("CommandDate", instanceWithInferedType.CommandDate, instanceWithInferedType, probe.formStage, formGroup, false)
		EnumTypeStringToForm("SpeedCommandType", instanceWithInferedType.SpeedCommandType, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("DateSpeedCommand", instanceWithInferedType.DateSpeedCommand, instanceWithInferedType, probe.formStage, formGroup, false)
		AssociationFieldToForm("Engine", instanceWithInferedType.Engine, formGroup, probe)

	case *models.GongsimStatus:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup, false)
		EnumTypeStringToForm("CurrentCommand", instanceWithInferedType.CurrentCommand, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("CompletionDate", instanceWithInferedType.CompletionDate, instanceWithInferedType, probe.formStage, formGroup, false)
		EnumTypeStringToForm("CurrentSpeedCommand", instanceWithInferedType.CurrentSpeedCommand, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("SpeedCommandCompletionDate", instanceWithInferedType.SpeedCommandCompletionDate, instanceWithInferedType, probe.formStage, formGroup, false)

	default:
		_ = instanceWithInferedType
	}
}
