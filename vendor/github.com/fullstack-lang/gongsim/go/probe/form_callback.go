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
	playground *Playground,
) (dummyagentFormCallback *DummyAgentFormCallback) {
	dummyagentFormCallback = new(DummyAgentFormCallback)
	dummyagentFormCallback.playground = playground
	dummyagentFormCallback.dummyagent = dummyagent

	dummyagentFormCallback.CreationMode = (dummyagent == nil)

	return
}

type DummyAgentFormCallback struct {
	dummyagent *models.DummyAgent

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (dummyagentFormCallback *DummyAgentFormCallback) OnSave() {

	log.Println("DummyAgentFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	dummyagentFormCallback.playground.formStage.Checkout()

	if dummyagentFormCallback.dummyagent == nil {
		dummyagentFormCallback.dummyagent = new(models.DummyAgent).Stage(dummyagentFormCallback.playground.stageOfInterest)
	}
	dummyagent_ := dummyagentFormCallback.dummyagent
	_ = dummyagent_

	// get the formGroup
	formGroup := dummyagentFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "TechName":
			FormDivBasicFieldToField(&(dummyagent_.TechName), formDiv)
		case "Name":
			FormDivBasicFieldToField(&(dummyagent_.Name), formDiv)
		}
	}

	dummyagentFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.DummyAgent](
		dummyagentFormCallback.playground,
	)
	dummyagentFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if dummyagentFormCallback.CreationMode {
		dummyagentFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__DummyAgentFormCallback(
				nil,
				dummyagentFormCallback.playground,
			),
		}).Stage(dummyagentFormCallback.playground.formStage)
		dummyagent := new(models.DummyAgent)
		FillUpForm(dummyagent, newFormGroup, dummyagentFormCallback.playground)
		dummyagentFormCallback.playground.formStage.Commit()
	}

	fillUpTree(dummyagentFormCallback.playground)
}
func __gong__New__EngineFormCallback(
	engine *models.Engine,
	playground *Playground,
) (engineFormCallback *EngineFormCallback) {
	engineFormCallback = new(EngineFormCallback)
	engineFormCallback.playground = playground
	engineFormCallback.engine = engine

	engineFormCallback.CreationMode = (engine == nil)

	return
}

type EngineFormCallback struct {
	engine *models.Engine

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (engineFormCallback *EngineFormCallback) OnSave() {

	log.Println("EngineFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	engineFormCallback.playground.formStage.Checkout()

	if engineFormCallback.engine == nil {
		engineFormCallback.engine = new(models.Engine).Stage(engineFormCallback.playground.stageOfInterest)
	}
	engine_ := engineFormCallback.engine
	_ = engine_

	// get the formGroup
	formGroup := engineFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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

	engineFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Engine](
		engineFormCallback.playground,
	)
	engineFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if engineFormCallback.CreationMode {
		engineFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__EngineFormCallback(
				nil,
				engineFormCallback.playground,
			),
		}).Stage(engineFormCallback.playground.formStage)
		engine := new(models.Engine)
		FillUpForm(engine, newFormGroup, engineFormCallback.playground)
		engineFormCallback.playground.formStage.Commit()
	}

	fillUpTree(engineFormCallback.playground)
}
func __gong__New__EventFormCallback(
	event *models.Event,
	playground *Playground,
) (eventFormCallback *EventFormCallback) {
	eventFormCallback = new(EventFormCallback)
	eventFormCallback.playground = playground
	eventFormCallback.event = event

	eventFormCallback.CreationMode = (event == nil)

	return
}

type EventFormCallback struct {
	event *models.Event

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (eventFormCallback *EventFormCallback) OnSave() {

	log.Println("EventFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	eventFormCallback.playground.formStage.Checkout()

	if eventFormCallback.event == nil {
		eventFormCallback.event = new(models.Event).Stage(eventFormCallback.playground.stageOfInterest)
	}
	event_ := eventFormCallback.event
	_ = event_

	// get the formGroup
	formGroup := eventFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(event_.Name), formDiv)
		case "Duration":
			FormDivBasicFieldToField(&(event_.Duration), formDiv)
		}
	}

	eventFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Event](
		eventFormCallback.playground,
	)
	eventFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if eventFormCallback.CreationMode {
		eventFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__EventFormCallback(
				nil,
				eventFormCallback.playground,
			),
		}).Stage(eventFormCallback.playground.formStage)
		event := new(models.Event)
		FillUpForm(event, newFormGroup, eventFormCallback.playground)
		eventFormCallback.playground.formStage.Commit()
	}

	fillUpTree(eventFormCallback.playground)
}
func __gong__New__GongsimCommandFormCallback(
	gongsimcommand *models.GongsimCommand,
	playground *Playground,
) (gongsimcommandFormCallback *GongsimCommandFormCallback) {
	gongsimcommandFormCallback = new(GongsimCommandFormCallback)
	gongsimcommandFormCallback.playground = playground
	gongsimcommandFormCallback.gongsimcommand = gongsimcommand

	gongsimcommandFormCallback.CreationMode = (gongsimcommand == nil)

	return
}

type GongsimCommandFormCallback struct {
	gongsimcommand *models.GongsimCommand

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (gongsimcommandFormCallback *GongsimCommandFormCallback) OnSave() {

	log.Println("GongsimCommandFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongsimcommandFormCallback.playground.formStage.Checkout()

	if gongsimcommandFormCallback.gongsimcommand == nil {
		gongsimcommandFormCallback.gongsimcommand = new(models.GongsimCommand).Stage(gongsimcommandFormCallback.playground.stageOfInterest)
	}
	gongsimcommand_ := gongsimcommandFormCallback.gongsimcommand
	_ = gongsimcommand_

	// get the formGroup
	formGroup := gongsimcommandFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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
			FormDivSelectFieldToField(&(gongsimcommand_.Engine), gongsimcommandFormCallback.playground.stageOfInterest, formDiv)
		}
	}

	gongsimcommandFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.GongsimCommand](
		gongsimcommandFormCallback.playground,
	)
	gongsimcommandFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if gongsimcommandFormCallback.CreationMode {
		gongsimcommandFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__GongsimCommandFormCallback(
				nil,
				gongsimcommandFormCallback.playground,
			),
		}).Stage(gongsimcommandFormCallback.playground.formStage)
		gongsimcommand := new(models.GongsimCommand)
		FillUpForm(gongsimcommand, newFormGroup, gongsimcommandFormCallback.playground)
		gongsimcommandFormCallback.playground.formStage.Commit()
	}

	fillUpTree(gongsimcommandFormCallback.playground)
}
func __gong__New__GongsimStatusFormCallback(
	gongsimstatus *models.GongsimStatus,
	playground *Playground,
) (gongsimstatusFormCallback *GongsimStatusFormCallback) {
	gongsimstatusFormCallback = new(GongsimStatusFormCallback)
	gongsimstatusFormCallback.playground = playground
	gongsimstatusFormCallback.gongsimstatus = gongsimstatus

	gongsimstatusFormCallback.CreationMode = (gongsimstatus == nil)

	return
}

type GongsimStatusFormCallback struct {
	gongsimstatus *models.GongsimStatus

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (gongsimstatusFormCallback *GongsimStatusFormCallback) OnSave() {

	log.Println("GongsimStatusFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongsimstatusFormCallback.playground.formStage.Checkout()

	if gongsimstatusFormCallback.gongsimstatus == nil {
		gongsimstatusFormCallback.gongsimstatus = new(models.GongsimStatus).Stage(gongsimstatusFormCallback.playground.stageOfInterest)
	}
	gongsimstatus_ := gongsimstatusFormCallback.gongsimstatus
	_ = gongsimstatus_

	// get the formGroup
	formGroup := gongsimstatusFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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

	gongsimstatusFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.GongsimStatus](
		gongsimstatusFormCallback.playground,
	)
	gongsimstatusFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if gongsimstatusFormCallback.CreationMode {
		gongsimstatusFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__GongsimStatusFormCallback(
				nil,
				gongsimstatusFormCallback.playground,
			),
		}).Stage(gongsimstatusFormCallback.playground.formStage)
		gongsimstatus := new(models.GongsimStatus)
		FillUpForm(gongsimstatus, newFormGroup, gongsimstatusFormCallback.playground)
		gongsimstatusFormCallback.playground.formStage.Commit()
	}

	fillUpTree(gongsimstatusFormCallback.playground)
}
