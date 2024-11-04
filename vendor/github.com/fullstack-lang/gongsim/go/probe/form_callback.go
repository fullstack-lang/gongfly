// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongsim/go/models"
	"github.com/fullstack-lang/gongsim/go/orm"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)
var __dummy_orm = orm.BackRepoStruct{}

// insertion point
func __gong__New__DummyAgentFormCallback(
	dummyagent *models.DummyAgent,
	probe *Probe,
	formGroup *table.FormGroup,
) (dummyagentFormCallback *DummyAgentFormCallback) {
	dummyagentFormCallback = new(DummyAgentFormCallback)
	dummyagentFormCallback.probe = probe
	dummyagentFormCallback.dummyagent = dummyagent
	dummyagentFormCallback.formGroup = formGroup

	dummyagentFormCallback.CreationMode = (dummyagent == nil)

	return
}

type DummyAgentFormCallback struct {
	dummyagent *models.DummyAgent

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (dummyagentFormCallback *DummyAgentFormCallback) OnSave() {

	log.Println("DummyAgentFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	dummyagentFormCallback.probe.formStage.Checkout()

	if dummyagentFormCallback.dummyagent == nil {
		dummyagentFormCallback.dummyagent = new(models.DummyAgent).Stage(dummyagentFormCallback.probe.stageOfInterest)
	}
	dummyagent_ := dummyagentFormCallback.dummyagent
	_ = dummyagent_

	for _, formDiv := range dummyagentFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "TechName":
			FormDivBasicFieldToField(&(dummyagent_.TechName), formDiv)
		case "Name":
			FormDivBasicFieldToField(&(dummyagent_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if dummyagentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		dummyagent_.Unstage(dummyagentFormCallback.probe.stageOfInterest)
	}

	dummyagentFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.DummyAgent](
		dummyagentFormCallback.probe,
	)
	dummyagentFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if dummyagentFormCallback.CreationMode || dummyagentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		dummyagentFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(dummyagentFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DummyAgentFormCallback(
			nil,
			dummyagentFormCallback.probe,
			newFormGroup,
		)
		dummyagent := new(models.DummyAgent)
		FillUpForm(dummyagent, newFormGroup, dummyagentFormCallback.probe)
		dummyagentFormCallback.probe.formStage.Commit()
	}

	fillUpTree(dummyagentFormCallback.probe)
}
func __gong__New__EngineFormCallback(
	engine *models.Engine,
	probe *Probe,
	formGroup *table.FormGroup,
) (engineFormCallback *EngineFormCallback) {
	engineFormCallback = new(EngineFormCallback)
	engineFormCallback.probe = probe
	engineFormCallback.engine = engine
	engineFormCallback.formGroup = formGroup

	engineFormCallback.CreationMode = (engine == nil)

	return
}

type EngineFormCallback struct {
	engine *models.Engine

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (engineFormCallback *EngineFormCallback) OnSave() {

	log.Println("EngineFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	engineFormCallback.probe.formStage.Checkout()

	if engineFormCallback.engine == nil {
		engineFormCallback.engine = new(models.Engine).Stage(engineFormCallback.probe.stageOfInterest)
	}
	engine_ := engineFormCallback.engine
	_ = engine_

	for _, formDiv := range engineFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(engine_.Name), formDiv)
		case "EndTime":
			FormDivBasicFieldToField(&(engine_.EndTime), formDiv)
		case "CurrentTime":
			FormDivBasicFieldToField(&(engine_.CurrentTime), formDiv)
		case "DisplayFormat":
			FormDivBasicFieldToField(&(engine_.DisplayFormat), formDiv)
		case "SecondsSinceStart":
			FormDivBasicFieldToField(&(engine_.SecondsSinceStart), formDiv)
		case "Fired":
			FormDivBasicFieldToField(&(engine_.Fired), formDiv)
		case "ControlMode":
			FormDivEnumStringFieldToField(&(engine_.ControlMode), formDiv)
		case "State":
			FormDivEnumStringFieldToField(&(engine_.State), formDiv)
		case "Speed":
			FormDivBasicFieldToField(&(engine_.Speed), formDiv)
		}
	}

	// manage the suppress operation
	if engineFormCallback.formGroup.HasSuppressButtonBeenPressed {
		engine_.Unstage(engineFormCallback.probe.stageOfInterest)
	}

	engineFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Engine](
		engineFormCallback.probe,
	)
	engineFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if engineFormCallback.CreationMode || engineFormCallback.formGroup.HasSuppressButtonBeenPressed {
		engineFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(engineFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__EngineFormCallback(
			nil,
			engineFormCallback.probe,
			newFormGroup,
		)
		engine := new(models.Engine)
		FillUpForm(engine, newFormGroup, engineFormCallback.probe)
		engineFormCallback.probe.formStage.Commit()
	}

	fillUpTree(engineFormCallback.probe)
}
func __gong__New__EventFormCallback(
	event *models.Event,
	probe *Probe,
	formGroup *table.FormGroup,
) (eventFormCallback *EventFormCallback) {
	eventFormCallback = new(EventFormCallback)
	eventFormCallback.probe = probe
	eventFormCallback.event = event
	eventFormCallback.formGroup = formGroup

	eventFormCallback.CreationMode = (event == nil)

	return
}

type EventFormCallback struct {
	event *models.Event

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (eventFormCallback *EventFormCallback) OnSave() {

	log.Println("EventFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	eventFormCallback.probe.formStage.Checkout()

	if eventFormCallback.event == nil {
		eventFormCallback.event = new(models.Event).Stage(eventFormCallback.probe.stageOfInterest)
	}
	event_ := eventFormCallback.event
	_ = event_

	for _, formDiv := range eventFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(event_.Name), formDiv)
		case "Duration":
			FormDivBasicFieldToField(&(event_.Duration), formDiv)
		}
	}

	// manage the suppress operation
	if eventFormCallback.formGroup.HasSuppressButtonBeenPressed {
		event_.Unstage(eventFormCallback.probe.stageOfInterest)
	}

	eventFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Event](
		eventFormCallback.probe,
	)
	eventFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if eventFormCallback.CreationMode || eventFormCallback.formGroup.HasSuppressButtonBeenPressed {
		eventFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(eventFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__EventFormCallback(
			nil,
			eventFormCallback.probe,
			newFormGroup,
		)
		event := new(models.Event)
		FillUpForm(event, newFormGroup, eventFormCallback.probe)
		eventFormCallback.probe.formStage.Commit()
	}

	fillUpTree(eventFormCallback.probe)
}
func __gong__New__GongsimCommandFormCallback(
	gongsimcommand *models.GongsimCommand,
	probe *Probe,
	formGroup *table.FormGroup,
) (gongsimcommandFormCallback *GongsimCommandFormCallback) {
	gongsimcommandFormCallback = new(GongsimCommandFormCallback)
	gongsimcommandFormCallback.probe = probe
	gongsimcommandFormCallback.gongsimcommand = gongsimcommand
	gongsimcommandFormCallback.formGroup = formGroup

	gongsimcommandFormCallback.CreationMode = (gongsimcommand == nil)

	return
}

type GongsimCommandFormCallback struct {
	gongsimcommand *models.GongsimCommand

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (gongsimcommandFormCallback *GongsimCommandFormCallback) OnSave() {

	log.Println("GongsimCommandFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongsimcommandFormCallback.probe.formStage.Checkout()

	if gongsimcommandFormCallback.gongsimcommand == nil {
		gongsimcommandFormCallback.gongsimcommand = new(models.GongsimCommand).Stage(gongsimcommandFormCallback.probe.stageOfInterest)
	}
	gongsimcommand_ := gongsimcommandFormCallback.gongsimcommand
	_ = gongsimcommand_

	for _, formDiv := range gongsimcommandFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongsimcommand_.Name), formDiv)
		case "Command":
			FormDivEnumStringFieldToField(&(gongsimcommand_.Command), formDiv)
		case "CommandDate":
			FormDivBasicFieldToField(&(gongsimcommand_.CommandDate), formDiv)
		case "SpeedCommandType":
			FormDivEnumStringFieldToField(&(gongsimcommand_.SpeedCommandType), formDiv)
		case "DateSpeedCommand":
			FormDivBasicFieldToField(&(gongsimcommand_.DateSpeedCommand), formDiv)
		case "Engine":
			FormDivSelectFieldToField(&(gongsimcommand_.Engine), gongsimcommandFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if gongsimcommandFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongsimcommand_.Unstage(gongsimcommandFormCallback.probe.stageOfInterest)
	}

	gongsimcommandFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.GongsimCommand](
		gongsimcommandFormCallback.probe,
	)
	gongsimcommandFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if gongsimcommandFormCallback.CreationMode || gongsimcommandFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongsimcommandFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(gongsimcommandFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GongsimCommandFormCallback(
			nil,
			gongsimcommandFormCallback.probe,
			newFormGroup,
		)
		gongsimcommand := new(models.GongsimCommand)
		FillUpForm(gongsimcommand, newFormGroup, gongsimcommandFormCallback.probe)
		gongsimcommandFormCallback.probe.formStage.Commit()
	}

	fillUpTree(gongsimcommandFormCallback.probe)
}
func __gong__New__GongsimStatusFormCallback(
	gongsimstatus *models.GongsimStatus,
	probe *Probe,
	formGroup *table.FormGroup,
) (gongsimstatusFormCallback *GongsimStatusFormCallback) {
	gongsimstatusFormCallback = new(GongsimStatusFormCallback)
	gongsimstatusFormCallback.probe = probe
	gongsimstatusFormCallback.gongsimstatus = gongsimstatus
	gongsimstatusFormCallback.formGroup = formGroup

	gongsimstatusFormCallback.CreationMode = (gongsimstatus == nil)

	return
}

type GongsimStatusFormCallback struct {
	gongsimstatus *models.GongsimStatus

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (gongsimstatusFormCallback *GongsimStatusFormCallback) OnSave() {

	log.Println("GongsimStatusFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongsimstatusFormCallback.probe.formStage.Checkout()

	if gongsimstatusFormCallback.gongsimstatus == nil {
		gongsimstatusFormCallback.gongsimstatus = new(models.GongsimStatus).Stage(gongsimstatusFormCallback.probe.stageOfInterest)
	}
	gongsimstatus_ := gongsimstatusFormCallback.gongsimstatus
	_ = gongsimstatus_

	for _, formDiv := range gongsimstatusFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongsimstatus_.Name), formDiv)
		case "CurrentCommand":
			FormDivEnumStringFieldToField(&(gongsimstatus_.CurrentCommand), formDiv)
		case "CompletionDate":
			FormDivBasicFieldToField(&(gongsimstatus_.CompletionDate), formDiv)
		case "CurrentSpeedCommand":
			FormDivEnumStringFieldToField(&(gongsimstatus_.CurrentSpeedCommand), formDiv)
		case "SpeedCommandCompletionDate":
			FormDivBasicFieldToField(&(gongsimstatus_.SpeedCommandCompletionDate), formDiv)
		}
	}

	// manage the suppress operation
	if gongsimstatusFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongsimstatus_.Unstage(gongsimstatusFormCallback.probe.stageOfInterest)
	}

	gongsimstatusFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.GongsimStatus](
		gongsimstatusFormCallback.probe,
	)
	gongsimstatusFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if gongsimstatusFormCallback.CreationMode || gongsimstatusFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongsimstatusFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(gongsimstatusFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GongsimStatusFormCallback(
			nil,
			gongsimstatusFormCallback.probe,
			newFormGroup,
		)
		gongsimstatus := new(models.GongsimStatus)
		FillUpForm(gongsimstatus, newFormGroup, gongsimstatusFormCallback.probe)
		gongsimstatusFormCallback.probe.formStage.Commit()
	}

	fillUpTree(gongsimstatusFormCallback.probe)
}
func __gong__New__UpdateStateFormCallback(
	updatestate *models.UpdateState,
	probe *Probe,
	formGroup *table.FormGroup,
) (updatestateFormCallback *UpdateStateFormCallback) {
	updatestateFormCallback = new(UpdateStateFormCallback)
	updatestateFormCallback.probe = probe
	updatestateFormCallback.updatestate = updatestate
	updatestateFormCallback.formGroup = formGroup

	updatestateFormCallback.CreationMode = (updatestate == nil)

	return
}

type UpdateStateFormCallback struct {
	updatestate *models.UpdateState

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (updatestateFormCallback *UpdateStateFormCallback) OnSave() {

	log.Println("UpdateStateFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	updatestateFormCallback.probe.formStage.Checkout()

	if updatestateFormCallback.updatestate == nil {
		updatestateFormCallback.updatestate = new(models.UpdateState).Stage(updatestateFormCallback.probe.stageOfInterest)
	}
	updatestate_ := updatestateFormCallback.updatestate
	_ = updatestate_

	for _, formDiv := range updatestateFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(updatestate_.Name), formDiv)
		case "Duration":
			FormDivBasicFieldToField(&(updatestate_.Duration), formDiv)
		case "Period":
			FormDivBasicFieldToField(&(updatestate_.Period), formDiv)
		}
	}

	// manage the suppress operation
	if updatestateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		updatestate_.Unstage(updatestateFormCallback.probe.stageOfInterest)
	}

	updatestateFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.UpdateState](
		updatestateFormCallback.probe,
	)
	updatestateFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if updatestateFormCallback.CreationMode || updatestateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		updatestateFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(updatestateFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__UpdateStateFormCallback(
			nil,
			updatestateFormCallback.probe,
			newFormGroup,
		)
		updatestate := new(models.UpdateState)
		FillUpForm(updatestate, newFormGroup, updatestateFormCallback.probe)
		updatestateFormCallback.probe.formStage.Commit()
	}

	fillUpTree(updatestateFormCallback.probe)
}
