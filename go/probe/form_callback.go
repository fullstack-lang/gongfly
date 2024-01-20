// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongfly/go/models"
	"github.com/fullstack-lang/gongfly/go/orm"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)
var __dummy_orm = orm.BackRepoStruct{}

// insertion point
func __gong__New__CivilianAirportFormCallback(
	civilianairport *models.CivilianAirport,
	probe *Probe,
	formGroup *table.FormGroup,
) (civilianairportFormCallback *CivilianAirportFormCallback) {
	civilianairportFormCallback = new(CivilianAirportFormCallback)
	civilianairportFormCallback.probe = probe
	civilianairportFormCallback.civilianairport = civilianairport
	civilianairportFormCallback.formGroup = formGroup

	civilianairportFormCallback.CreationMode = (civilianairport == nil)

	return
}

type CivilianAirportFormCallback struct {
	civilianairport *models.CivilianAirport

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (civilianairportFormCallback *CivilianAirportFormCallback) OnSave() {

	log.Println("CivilianAirportFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	civilianairportFormCallback.probe.formStage.Checkout()

	if civilianairportFormCallback.civilianairport == nil {
		civilianairportFormCallback.civilianairport = new(models.CivilianAirport).Stage(civilianairportFormCallback.probe.stageOfInterest)
	}
	civilianairport_ := civilianairportFormCallback.civilianairport
	_ = civilianairport_

	for _, formDiv := range civilianairportFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Lat":
			FormDivBasicFieldToField(&(civilianairport_.Lat), formDiv)
		case "Lng":
			FormDivBasicFieldToField(&(civilianairport_.Lng), formDiv)
		case "Name":
			FormDivBasicFieldToField(&(civilianairport_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if civilianairportFormCallback.formGroup.HasSuppressButtonBeenPressed {
		civilianairport_.Unstage(civilianairportFormCallback.probe.stageOfInterest)
	}

	civilianairportFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.CivilianAirport](
		civilianairportFormCallback.probe,
	)
	civilianairportFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if civilianairportFormCallback.CreationMode || civilianairportFormCallback.formGroup.HasSuppressButtonBeenPressed {
		civilianairportFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(civilianairportFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__CivilianAirportFormCallback(
			nil,
			civilianairportFormCallback.probe,
			newFormGroup,
		)
		civilianairport := new(models.CivilianAirport)
		FillUpForm(civilianairport, newFormGroup, civilianairportFormCallback.probe)
		civilianairportFormCallback.probe.formStage.Commit()
	}

	fillUpTree(civilianairportFormCallback.probe)
}
func __gong__New__LinerFormCallback(
	liner *models.Liner,
	probe *Probe,
	formGroup *table.FormGroup,
) (linerFormCallback *LinerFormCallback) {
	linerFormCallback = new(LinerFormCallback)
	linerFormCallback.probe = probe
	linerFormCallback.liner = liner
	linerFormCallback.formGroup = formGroup

	linerFormCallback.CreationMode = (liner == nil)

	return
}

type LinerFormCallback struct {
	liner *models.Liner

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (linerFormCallback *LinerFormCallback) OnSave() {

	log.Println("LinerFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	linerFormCallback.probe.formStage.Checkout()

	if linerFormCallback.liner == nil {
		linerFormCallback.liner = new(models.Liner).Stage(linerFormCallback.probe.stageOfInterest)
	}
	liner_ := linerFormCallback.liner
	_ = liner_

	for _, formDiv := range linerFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(liner_.Name), formDiv)
		case "Lat":
			FormDivBasicFieldToField(&(liner_.Lat), formDiv)
		case "Lng":
			FormDivBasicFieldToField(&(liner_.Lng), formDiv)
		case "Heading":
			FormDivBasicFieldToField(&(liner_.Heading), formDiv)
		case "Level":
			FormDivBasicFieldToField(&(liner_.Level), formDiv)
		case "Speed":
			FormDivBasicFieldToField(&(liner_.Speed), formDiv)
		case "State":
			FormDivEnumStringFieldToField(&(liner_.State), formDiv)
		case "TargetHeading":
			FormDivBasicFieldToField(&(liner_.TargetHeading), formDiv)
		case "TargetLocationLat":
			FormDivBasicFieldToField(&(liner_.TargetLocationLat), formDiv)
		case "TargetLocationLng":
			FormDivBasicFieldToField(&(liner_.TargetLocationLng), formDiv)
		case "DistanceToTarget":
			FormDivBasicFieldToField(&(liner_.DistanceToTarget), formDiv)
		case "MaxRotationalSpeed":
			FormDivBasicFieldToField(&(liner_.MaxRotationalSpeed), formDiv)
		case "VerticalSpeed":
			FormDivBasicFieldToField(&(liner_.VerticalSpeed), formDiv)
		case "Timestampstring":
			FormDivBasicFieldToField(&(liner_.Timestampstring), formDiv)
		case "ReporingLine":
			FormDivSelectFieldToField(&(liner_.ReporingLine), linerFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if linerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		liner_.Unstage(linerFormCallback.probe.stageOfInterest)
	}

	linerFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Liner](
		linerFormCallback.probe,
	)
	linerFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if linerFormCallback.CreationMode || linerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		linerFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(linerFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__LinerFormCallback(
			nil,
			linerFormCallback.probe,
			newFormGroup,
		)
		liner := new(models.Liner)
		FillUpForm(liner, newFormGroup, linerFormCallback.probe)
		linerFormCallback.probe.formStage.Commit()
	}

	fillUpTree(linerFormCallback.probe)
}
func __gong__New__MessageFormCallback(
	message *models.Message,
	probe *Probe,
	formGroup *table.FormGroup,
) (messageFormCallback *MessageFormCallback) {
	messageFormCallback = new(MessageFormCallback)
	messageFormCallback.probe = probe
	messageFormCallback.message = message
	messageFormCallback.formGroup = formGroup

	messageFormCallback.CreationMode = (message == nil)

	return
}

type MessageFormCallback struct {
	message *models.Message

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (messageFormCallback *MessageFormCallback) OnSave() {

	log.Println("MessageFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	messageFormCallback.probe.formStage.Checkout()

	if messageFormCallback.message == nil {
		messageFormCallback.message = new(models.Message).Stage(messageFormCallback.probe.stageOfInterest)
	}
	message_ := messageFormCallback.message
	_ = message_

	for _, formDiv := range messageFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Lat":
			FormDivBasicFieldToField(&(message_.Lat), formDiv)
		case "Lng":
			FormDivBasicFieldToField(&(message_.Lng), formDiv)
		case "Heading":
			FormDivBasicFieldToField(&(message_.Heading), formDiv)
		case "Level":
			FormDivBasicFieldToField(&(message_.Level), formDiv)
		case "Speed":
			FormDivBasicFieldToField(&(message_.Speed), formDiv)
		case "State":
			FormDivEnumStringFieldToField(&(message_.State), formDiv)
		case "Name":
			FormDivBasicFieldToField(&(message_.Name), formDiv)
		case "TargetLocationLat":
			FormDivBasicFieldToField(&(message_.TargetLocationLat), formDiv)
		case "TargetLocationLng":
			FormDivBasicFieldToField(&(message_.TargetLocationLng), formDiv)
		case "DistanceToTarget":
			FormDivBasicFieldToField(&(message_.DistanceToTarget), formDiv)
		case "Timestampstring":
			FormDivBasicFieldToField(&(message_.Timestampstring), formDiv)
		case "DurationSinceSimulationStart":
			FormDivBasicFieldToField(&(message_.DurationSinceSimulationStart), formDiv)
		case "Timestampstartstring":
			FormDivBasicFieldToField(&(message_.Timestampstartstring), formDiv)
		case "Source":
			FormDivBasicFieldToField(&(message_.Source), formDiv)
		case "Destination":
			FormDivBasicFieldToField(&(message_.Destination), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(message_.Content), formDiv)
		case "About_string":
			FormDivBasicFieldToField(&(message_.About_string), formDiv)
		case "Display":
			FormDivBasicFieldToField(&(message_.Display), formDiv)
		}
	}

	// manage the suppress operation
	if messageFormCallback.formGroup.HasSuppressButtonBeenPressed {
		message_.Unstage(messageFormCallback.probe.stageOfInterest)
	}

	messageFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Message](
		messageFormCallback.probe,
	)
	messageFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if messageFormCallback.CreationMode || messageFormCallback.formGroup.HasSuppressButtonBeenPressed {
		messageFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(messageFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__MessageFormCallback(
			nil,
			messageFormCallback.probe,
			newFormGroup,
		)
		message := new(models.Message)
		FillUpForm(message, newFormGroup, messageFormCallback.probe)
		messageFormCallback.probe.formStage.Commit()
	}

	fillUpTree(messageFormCallback.probe)
}
func __gong__New__OpsLineFormCallback(
	opsline *models.OpsLine,
	probe *Probe,
	formGroup *table.FormGroup,
) (opslineFormCallback *OpsLineFormCallback) {
	opslineFormCallback = new(OpsLineFormCallback)
	opslineFormCallback.probe = probe
	opslineFormCallback.opsline = opsline
	opslineFormCallback.formGroup = formGroup

	opslineFormCallback.CreationMode = (opsline == nil)

	return
}

type OpsLineFormCallback struct {
	opsline *models.OpsLine

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (opslineFormCallback *OpsLineFormCallback) OnSave() {

	log.Println("OpsLineFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	opslineFormCallback.probe.formStage.Checkout()

	if opslineFormCallback.opsline == nil {
		opslineFormCallback.opsline = new(models.OpsLine).Stage(opslineFormCallback.probe.stageOfInterest)
	}
	opsline_ := opslineFormCallback.opsline
	_ = opsline_

	for _, formDiv := range opslineFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "IsTransmitting":
			FormDivBasicFieldToField(&(opsline_.IsTransmitting), formDiv)
		case "TransmissionMessage":
			FormDivBasicFieldToField(&(opsline_.TransmissionMessage), formDiv)
		case "IsTransmittingBackward":
			FormDivBasicFieldToField(&(opsline_.IsTransmittingBackward), formDiv)
		case "TransmissionMessageBackward":
			FormDivBasicFieldToField(&(opsline_.TransmissionMessageBackward), formDiv)
		case "Scenario":
			FormDivSelectFieldToField(&(opsline_.Scenario), opslineFormCallback.probe.stageOfInterest, formDiv)
		case "State":
			FormDivEnumStringFieldToField(&(opsline_.State), formDiv)
		case "Name":
			FormDivBasicFieldToField(&(opsline_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if opslineFormCallback.formGroup.HasSuppressButtonBeenPressed {
		opsline_.Unstage(opslineFormCallback.probe.stageOfInterest)
	}

	opslineFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.OpsLine](
		opslineFormCallback.probe,
	)
	opslineFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if opslineFormCallback.CreationMode || opslineFormCallback.formGroup.HasSuppressButtonBeenPressed {
		opslineFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(opslineFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__OpsLineFormCallback(
			nil,
			opslineFormCallback.probe,
			newFormGroup,
		)
		opsline := new(models.OpsLine)
		FillUpForm(opsline, newFormGroup, opslineFormCallback.probe)
		opslineFormCallback.probe.formStage.Commit()
	}

	fillUpTree(opslineFormCallback.probe)
}
func __gong__New__RadarFormCallback(
	radar *models.Radar,
	probe *Probe,
	formGroup *table.FormGroup,
) (radarFormCallback *RadarFormCallback) {
	radarFormCallback = new(RadarFormCallback)
	radarFormCallback.probe = probe
	radarFormCallback.radar = radar
	radarFormCallback.formGroup = formGroup

	radarFormCallback.CreationMode = (radar == nil)

	return
}

type RadarFormCallback struct {
	radar *models.Radar

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (radarFormCallback *RadarFormCallback) OnSave() {

	log.Println("RadarFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	radarFormCallback.probe.formStage.Checkout()

	if radarFormCallback.radar == nil {
		radarFormCallback.radar = new(models.Radar).Stage(radarFormCallback.probe.stageOfInterest)
	}
	radar_ := radarFormCallback.radar
	_ = radar_

	for _, formDiv := range radarFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "State":
			FormDivEnumStringFieldToField(&(radar_.State), formDiv)
		case "Name":
			FormDivBasicFieldToField(&(radar_.Name), formDiv)
		case "Lat":
			FormDivBasicFieldToField(&(radar_.Lat), formDiv)
		case "Lng":
			FormDivBasicFieldToField(&(radar_.Lng), formDiv)
		case "Range":
			FormDivBasicFieldToField(&(radar_.Range), formDiv)
		}
	}

	// manage the suppress operation
	if radarFormCallback.formGroup.HasSuppressButtonBeenPressed {
		radar_.Unstage(radarFormCallback.probe.stageOfInterest)
	}

	radarFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Radar](
		radarFormCallback.probe,
	)
	radarFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if radarFormCallback.CreationMode || radarFormCallback.formGroup.HasSuppressButtonBeenPressed {
		radarFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(radarFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RadarFormCallback(
			nil,
			radarFormCallback.probe,
			newFormGroup,
		)
		radar := new(models.Radar)
		FillUpForm(radar, newFormGroup, radarFormCallback.probe)
		radarFormCallback.probe.formStage.Commit()
	}

	fillUpTree(radarFormCallback.probe)
}
func __gong__New__SatelliteFormCallback(
	satellite *models.Satellite,
	probe *Probe,
	formGroup *table.FormGroup,
) (satelliteFormCallback *SatelliteFormCallback) {
	satelliteFormCallback = new(SatelliteFormCallback)
	satelliteFormCallback.probe = probe
	satelliteFormCallback.satellite = satellite
	satelliteFormCallback.formGroup = formGroup

	satelliteFormCallback.CreationMode = (satellite == nil)

	return
}

type SatelliteFormCallback struct {
	satellite *models.Satellite

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (satelliteFormCallback *SatelliteFormCallback) OnSave() {

	log.Println("SatelliteFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	satelliteFormCallback.probe.formStage.Checkout()

	if satelliteFormCallback.satellite == nil {
		satelliteFormCallback.satellite = new(models.Satellite).Stage(satelliteFormCallback.probe.stageOfInterest)
	}
	satellite_ := satelliteFormCallback.satellite
	_ = satellite_

	for _, formDiv := range satelliteFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(satellite_.Name), formDiv)
		case "Line1":
			FormDivBasicFieldToField(&(satellite_.Line1), formDiv)
		case "Line2":
			FormDivBasicFieldToField(&(satellite_.Line2), formDiv)
		case "Lat":
			FormDivBasicFieldToField(&(satellite_.Lat), formDiv)
		case "Lng":
			FormDivBasicFieldToField(&(satellite_.Lng), formDiv)
		case "Heading":
			FormDivBasicFieldToField(&(satellite_.Heading), formDiv)
		case "Level":
			FormDivBasicFieldToField(&(satellite_.Level), formDiv)
		case "Speed":
			FormDivBasicFieldToField(&(satellite_.Speed), formDiv)
		case "VerticalSpeed":
			FormDivBasicFieldToField(&(satellite_.VerticalSpeed), formDiv)
		case "Timestampstring":
			FormDivBasicFieldToField(&(satellite_.Timestampstring), formDiv)
		}
	}

	// manage the suppress operation
	if satelliteFormCallback.formGroup.HasSuppressButtonBeenPressed {
		satellite_.Unstage(satelliteFormCallback.probe.stageOfInterest)
	}

	satelliteFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Satellite](
		satelliteFormCallback.probe,
	)
	satelliteFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if satelliteFormCallback.CreationMode || satelliteFormCallback.formGroup.HasSuppressButtonBeenPressed {
		satelliteFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(satelliteFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SatelliteFormCallback(
			nil,
			satelliteFormCallback.probe,
			newFormGroup,
		)
		satellite := new(models.Satellite)
		FillUpForm(satellite, newFormGroup, satelliteFormCallback.probe)
		satelliteFormCallback.probe.formStage.Commit()
	}

	fillUpTree(satelliteFormCallback.probe)
}
func __gong__New__ScenarioFormCallback(
	scenario *models.Scenario,
	probe *Probe,
	formGroup *table.FormGroup,
) (scenarioFormCallback *ScenarioFormCallback) {
	scenarioFormCallback = new(ScenarioFormCallback)
	scenarioFormCallback.probe = probe
	scenarioFormCallback.scenario = scenario
	scenarioFormCallback.formGroup = formGroup

	scenarioFormCallback.CreationMode = (scenario == nil)

	return
}

type ScenarioFormCallback struct {
	scenario *models.Scenario

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (scenarioFormCallback *ScenarioFormCallback) OnSave() {

	log.Println("ScenarioFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	scenarioFormCallback.probe.formStage.Checkout()

	if scenarioFormCallback.scenario == nil {
		scenarioFormCallback.scenario = new(models.Scenario).Stage(scenarioFormCallback.probe.stageOfInterest)
	}
	scenario_ := scenarioFormCallback.scenario
	_ = scenario_

	for _, formDiv := range scenarioFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(scenario_.Name), formDiv)
		case "Lat":
			FormDivBasicFieldToField(&(scenario_.Lat), formDiv)
		case "Lng":
			FormDivBasicFieldToField(&(scenario_.Lng), formDiv)
		case "ZoomLevel":
			FormDivBasicFieldToField(&(scenario_.ZoomLevel), formDiv)
		case "MessageVisualSpeed":
			FormDivBasicFieldToField(&(scenario_.MessageVisualSpeed), formDiv)
		}
	}

	// manage the suppress operation
	if scenarioFormCallback.formGroup.HasSuppressButtonBeenPressed {
		scenario_.Unstage(scenarioFormCallback.probe.stageOfInterest)
	}

	scenarioFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Scenario](
		scenarioFormCallback.probe,
	)
	scenarioFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if scenarioFormCallback.CreationMode || scenarioFormCallback.formGroup.HasSuppressButtonBeenPressed {
		scenarioFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(scenarioFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ScenarioFormCallback(
			nil,
			scenarioFormCallback.probe,
			newFormGroup,
		)
		scenario := new(models.Scenario)
		FillUpForm(scenario, newFormGroup, scenarioFormCallback.probe)
		scenarioFormCallback.probe.formStage.Commit()
	}

	fillUpTree(scenarioFormCallback.probe)
}
