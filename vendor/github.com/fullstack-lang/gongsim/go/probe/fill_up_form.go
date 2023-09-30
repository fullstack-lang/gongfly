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
	playground *Playground,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.DummyAgent:
		// insertion point
		BasicFieldtoForm("TechName", instanceWithInferedType.TechName, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)

	case *models.Engine:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("EndTime", instanceWithInferedType.EndTime, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("CurrentTime", instanceWithInferedType.CurrentTime, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("SecondsSinceStart", instanceWithInferedType.SecondsSinceStart, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Fired", instanceWithInferedType.Fired, instanceWithInferedType, playground.formStage, formGroup, false)
		EnumTypeStringToForm("ControlMode", instanceWithInferedType.ControlMode, instanceWithInferedType, playground.formStage, formGroup)
		EnumTypeStringToForm("State", instanceWithInferedType.State, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Speed", instanceWithInferedType.Speed, instanceWithInferedType, playground.formStage, formGroup, false)

	case *models.Event:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Duration", instanceWithInferedType.Duration, instanceWithInferedType, playground.formStage, formGroup, false)

	case *models.GongsimCommand:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		EnumTypeStringToForm("Command", instanceWithInferedType.Command, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("CommandDate", instanceWithInferedType.CommandDate, instanceWithInferedType, playground.formStage, formGroup, false)
		EnumTypeStringToForm("SpeedCommandType", instanceWithInferedType.SpeedCommandType, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("DateSpeedCommand", instanceWithInferedType.DateSpeedCommand, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationFieldToForm("Engine", instanceWithInferedType.Engine, formGroup, playground)

	case *models.GongsimStatus:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		EnumTypeStringToForm("CurrentCommand", instanceWithInferedType.CurrentCommand, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("CompletionDate", instanceWithInferedType.CompletionDate, instanceWithInferedType, playground.formStage, formGroup, false)
		EnumTypeStringToForm("CurrentSpeedCommand", instanceWithInferedType.CurrentSpeedCommand, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("SpeedCommandCompletionDate", instanceWithInferedType.SpeedCommandCompletionDate, instanceWithInferedType, playground.formStage, formGroup, false)

	default:
		_ = instanceWithInferedType
	}
}
