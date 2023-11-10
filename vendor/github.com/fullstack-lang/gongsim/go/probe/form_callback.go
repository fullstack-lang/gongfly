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
) (dummyagentFormCallback *DummyAgentFormCallback) {
	dummyagentFormCallback = new(DummyAgentFormCallback)
	dummyagentFormCallback.probe = probe
	dummyagentFormCallback.dummyagent = dummyagent

	dummyagentFormCallback.CreationMode = (dummyagent == nil)

	return
}

type DummyAgentFormCallback struct {
	dummyagent *models.DummyAgent

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
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

	// get the formGroup
	formGroup := dummyagentFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "TechName":
			FormDivBasicFieldToField(&(dummyagent_.TechName), formDiv)
		case "Name":
			FormDivBasicFieldToField(&(dummyagent_.Name), formDiv)
		}
	}

	dummyagentFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.DummyAgent](
		dummyagentFormCallback.probe,
	)
	dummyagentFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if dummyagentFormCallback.CreationMode {
		dummyagentFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__DummyAgentFormCallback(
				nil,
				dummyagentFormCallback.probe,
			),
		}).Stage(dummyagentFormCallback.probe.formStage)
		dummyagent := new(models.DummyAgent)
		FillUpForm(dummyagent, newFormGroup, dummyagentFormCallback.probe)
		dummyagentFormCallback.probe.formStage.Commit()
	}

	fillUpTree(dummyagentFormCallback.probe)
}
func __gong__New__EngineFormCallback(
	engine *models.Engine,
	probe *Probe,
) (engineFormCallback *EngineFormCallback) {
	engineFormCallback = new(EngineFormCallback)
	engineFormCallback.probe = probe
	engineFormCallback.engine = engine

	engineFormCallback.CreationMode = (engine == nil)

	return
}

type EngineFormCallback struct {
	engine *models.Engine

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
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

	// get the formGroup
	formGroup := engineFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(engine_.Name), formDiv)
		case "EndTime":
			FormDivBasicFieldToField(&(engine_.EndTime), formDiv)
		case "CurrentTime":
			FormDivBasicFieldToField(&(engine_.CurrentTime), formDiv)
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

	engineFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Engine](
		engineFormCallback.probe,
	)
	engineFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if engineFormCallback.CreationMode {
		engineFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__EngineFormCallback(
				nil,
				engineFormCallback.probe,
			),
		}).Stage(engineFormCallback.probe.formStage)
		engine := new(models.Engine)
		FillUpForm(engine, newFormGroup, engineFormCallback.probe)
		engineFormCallback.probe.formStage.Commit()
	}

	fillUpTree(engineFormCallback.probe)
}
func __gong__New__EventFormCallback(
	event *models.Event,
	probe *Probe,
) (eventFormCallback *EventFormCallback) {
	eventFormCallback = new(EventFormCallback)
	eventFormCallback.probe = probe
	eventFormCallback.event = event

	eventFormCallback.CreationMode = (event == nil)

	return
}

type EventFormCallback struct {
	event *models.Event

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
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

	// get the formGroup
	formGroup := eventFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(event_.Name), formDiv)
		case "Duration":
			FormDivBasicFieldToField(&(event_.Duration), formDiv)
		}
	}

	eventFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Event](
		eventFormCallback.probe,
	)
	eventFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if eventFormCallback.CreationMode {
		eventFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__EventFormCallback(
				nil,
				eventFormCallback.probe,
			),
		}).Stage(eventFormCallback.probe.formStage)
		event := new(models.Event)
		FillUpForm(event, newFormGroup, eventFormCallback.probe)
		eventFormCallback.probe.formStage.Commit()
	}

	fillUpTree(eventFormCallback.probe)
}
func __gong__New__GongsimCommandFormCallback(
	gongsimcommand *models.GongsimCommand,
	probe *Probe,
) (gongsimcommandFormCallback *GongsimCommandFormCallback) {
	gongsimcommandFormCallback = new(GongsimCommandFormCallback)
	gongsimcommandFormCallback.probe = probe
	gongsimcommandFormCallback.gongsimcommand = gongsimcommand

	gongsimcommandFormCallback.CreationMode = (gongsimcommand == nil)

	return
}

type GongsimCommandFormCallback struct {
	gongsimcommand *models.GongsimCommand

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
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

	// get the formGroup
	formGroup := gongsimcommandFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
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

	gongsimcommandFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.GongsimCommand](
		gongsimcommandFormCallback.probe,
	)
	gongsimcommandFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if gongsimcommandFormCallback.CreationMode {
		gongsimcommandFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__GongsimCommandFormCallback(
				nil,
				gongsimcommandFormCallback.probe,
			),
		}).Stage(gongsimcommandFormCallback.probe.formStage)
		gongsimcommand := new(models.GongsimCommand)
		FillUpForm(gongsimcommand, newFormGroup, gongsimcommandFormCallback.probe)
		gongsimcommandFormCallback.probe.formStage.Commit()
	}

	fillUpTree(gongsimcommandFormCallback.probe)
}
func __gong__New__GongsimStatusFormCallback(
	gongsimstatus *models.GongsimStatus,
	probe *Probe,
) (gongsimstatusFormCallback *GongsimStatusFormCallback) {
	gongsimstatusFormCallback = new(GongsimStatusFormCallback)
	gongsimstatusFormCallback.probe = probe
	gongsimstatusFormCallback.gongsimstatus = gongsimstatus

	gongsimstatusFormCallback.CreationMode = (gongsimstatus == nil)

	return
}

type GongsimStatusFormCallback struct {
	gongsimstatus *models.GongsimStatus

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
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

	// get the formGroup
	formGroup := gongsimstatusFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
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

	gongsimstatusFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.GongsimStatus](
		gongsimstatusFormCallback.probe,
	)
	gongsimstatusFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if gongsimstatusFormCallback.CreationMode {
		gongsimstatusFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__GongsimStatusFormCallback(
				nil,
				gongsimstatusFormCallback.probe,
			),
		}).Stage(gongsimstatusFormCallback.probe.formStage)
		gongsimstatus := new(models.GongsimStatus)
		FillUpForm(gongsimstatus, newFormGroup, gongsimstatusFormCallback.probe)
		gongsimstatusFormCallback.probe.formStage.Commit()
	}

	fillUpTree(gongsimstatusFormCallback.probe)
}
