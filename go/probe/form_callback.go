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
	playground *Playground,
) (civilianairportFormCallback *CivilianAirportFormCallback) {
	civilianairportFormCallback = new(CivilianAirportFormCallback)
	civilianairportFormCallback.playground = playground
	civilianairportFormCallback.civilianairport = civilianairport

	civilianairportFormCallback.CreationMode = (civilianairport == nil)

	return
}

type CivilianAirportFormCallback struct {
	civilianairport *models.CivilianAirport

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (civilianairportFormCallback *CivilianAirportFormCallback) OnSave() {

	log.Println("CivilianAirportFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	civilianairportFormCallback.playground.formStage.Checkout()

	if civilianairportFormCallback.civilianairport == nil {
		civilianairportFormCallback.civilianairport = new(models.CivilianAirport).Stage(civilianairportFormCallback.playground.stageOfInterest)
	}
	civilianairport_ := civilianairportFormCallback.civilianairport
	_ = civilianairport_

	// get the formGroup
	formGroup := civilianairportFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
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

	civilianairportFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.CivilianAirport](
		civilianairportFormCallback.playground,
	)
	civilianairportFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if civilianairportFormCallback.CreationMode {
		civilianairportFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__CivilianAirportFormCallback(
				nil,
				civilianairportFormCallback.playground,
			),
		}).Stage(civilianairportFormCallback.playground.formStage)
		civilianairport := new(models.CivilianAirport)
		FillUpForm(civilianairport, newFormGroup, civilianairportFormCallback.playground)
		civilianairportFormCallback.playground.formStage.Commit()
	}

	fillUpTree(civilianairportFormCallback.playground)
}
func __gong__New__LinerFormCallback(
	liner *models.Liner,
	playground *Playground,
) (linerFormCallback *LinerFormCallback) {
	linerFormCallback = new(LinerFormCallback)
	linerFormCallback.playground = playground
	linerFormCallback.liner = liner

	linerFormCallback.CreationMode = (liner == nil)

	return
}

type LinerFormCallback struct {
	liner *models.Liner

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (linerFormCallback *LinerFormCallback) OnSave() {

	log.Println("LinerFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	linerFormCallback.playground.formStage.Checkout()

	if linerFormCallback.liner == nil {
		linerFormCallback.liner = new(models.Liner).Stage(linerFormCallback.playground.stageOfInterest)
	}
	liner_ := linerFormCallback.liner
	_ = liner_

	// get the formGroup
	formGroup := linerFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
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
			FormDivSelectFieldToField(&(liner_.ReporingLine), linerFormCallback.playground.stageOfInterest, formDiv)
		}
	}

	linerFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Liner](
		linerFormCallback.playground,
	)
	linerFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if linerFormCallback.CreationMode {
		linerFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__LinerFormCallback(
				nil,
				linerFormCallback.playground,
			),
		}).Stage(linerFormCallback.playground.formStage)
		liner := new(models.Liner)
		FillUpForm(liner, newFormGroup, linerFormCallback.playground)
		linerFormCallback.playground.formStage.Commit()
	}

	fillUpTree(linerFormCallback.playground)
}
func __gong__New__MessageFormCallback(
	message *models.Message,
	playground *Playground,
) (messageFormCallback *MessageFormCallback) {
	messageFormCallback = new(MessageFormCallback)
	messageFormCallback.playground = playground
	messageFormCallback.message = message

	messageFormCallback.CreationMode = (message == nil)

	return
}

type MessageFormCallback struct {
	message *models.Message

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (messageFormCallback *MessageFormCallback) OnSave() {

	log.Println("MessageFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	messageFormCallback.playground.formStage.Checkout()

	if messageFormCallback.message == nil {
		messageFormCallback.message = new(models.Message).Stage(messageFormCallback.playground.stageOfInterest)
	}
	message_ := messageFormCallback.message
	_ = message_

	// get the formGroup
	formGroup := messageFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
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

	messageFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Message](
		messageFormCallback.playground,
	)
	messageFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if messageFormCallback.CreationMode {
		messageFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__MessageFormCallback(
				nil,
				messageFormCallback.playground,
			),
		}).Stage(messageFormCallback.playground.formStage)
		message := new(models.Message)
		FillUpForm(message, newFormGroup, messageFormCallback.playground)
		messageFormCallback.playground.formStage.Commit()
	}

	fillUpTree(messageFormCallback.playground)
}
func __gong__New__OpsLineFormCallback(
	opsline *models.OpsLine,
	playground *Playground,
) (opslineFormCallback *OpsLineFormCallback) {
	opslineFormCallback = new(OpsLineFormCallback)
	opslineFormCallback.playground = playground
	opslineFormCallback.opsline = opsline

	opslineFormCallback.CreationMode = (opsline == nil)

	return
}

type OpsLineFormCallback struct {
	opsline *models.OpsLine

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (opslineFormCallback *OpsLineFormCallback) OnSave() {

	log.Println("OpsLineFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	opslineFormCallback.playground.formStage.Checkout()

	if opslineFormCallback.opsline == nil {
		opslineFormCallback.opsline = new(models.OpsLine).Stage(opslineFormCallback.playground.stageOfInterest)
	}
	opsline_ := opslineFormCallback.opsline
	_ = opsline_

	// get the formGroup
	formGroup := opslineFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
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
			FormDivSelectFieldToField(&(opsline_.Scenario), opslineFormCallback.playground.stageOfInterest, formDiv)
		case "State":
			FormDivEnumStringFieldToField(&(opsline_.State), formDiv)
		case "Name":
			FormDivBasicFieldToField(&(opsline_.Name), formDiv)
		}
	}

	opslineFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.OpsLine](
		opslineFormCallback.playground,
	)
	opslineFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if opslineFormCallback.CreationMode {
		opslineFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__OpsLineFormCallback(
				nil,
				opslineFormCallback.playground,
			),
		}).Stage(opslineFormCallback.playground.formStage)
		opsline := new(models.OpsLine)
		FillUpForm(opsline, newFormGroup, opslineFormCallback.playground)
		opslineFormCallback.playground.formStage.Commit()
	}

	fillUpTree(opslineFormCallback.playground)
}
func __gong__New__RadarFormCallback(
	radar *models.Radar,
	playground *Playground,
) (radarFormCallback *RadarFormCallback) {
	radarFormCallback = new(RadarFormCallback)
	radarFormCallback.playground = playground
	radarFormCallback.radar = radar

	radarFormCallback.CreationMode = (radar == nil)

	return
}

type RadarFormCallback struct {
	radar *models.Radar

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (radarFormCallback *RadarFormCallback) OnSave() {

	log.Println("RadarFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	radarFormCallback.playground.formStage.Checkout()

	if radarFormCallback.radar == nil {
		radarFormCallback.radar = new(models.Radar).Stage(radarFormCallback.playground.stageOfInterest)
	}
	radar_ := radarFormCallback.radar
	_ = radar_

	// get the formGroup
	formGroup := radarFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
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

	radarFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Radar](
		radarFormCallback.playground,
	)
	radarFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if radarFormCallback.CreationMode {
		radarFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__RadarFormCallback(
				nil,
				radarFormCallback.playground,
			),
		}).Stage(radarFormCallback.playground.formStage)
		radar := new(models.Radar)
		FillUpForm(radar, newFormGroup, radarFormCallback.playground)
		radarFormCallback.playground.formStage.Commit()
	}

	fillUpTree(radarFormCallback.playground)
}
func __gong__New__SatelliteFormCallback(
	satellite *models.Satellite,
	playground *Playground,
) (satelliteFormCallback *SatelliteFormCallback) {
	satelliteFormCallback = new(SatelliteFormCallback)
	satelliteFormCallback.playground = playground
	satelliteFormCallback.satellite = satellite

	satelliteFormCallback.CreationMode = (satellite == nil)

	return
}

type SatelliteFormCallback struct {
	satellite *models.Satellite

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (satelliteFormCallback *SatelliteFormCallback) OnSave() {

	log.Println("SatelliteFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	satelliteFormCallback.playground.formStage.Checkout()

	if satelliteFormCallback.satellite == nil {
		satelliteFormCallback.satellite = new(models.Satellite).Stage(satelliteFormCallback.playground.stageOfInterest)
	}
	satellite_ := satelliteFormCallback.satellite
	_ = satellite_

	// get the formGroup
	formGroup := satelliteFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
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

	satelliteFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Satellite](
		satelliteFormCallback.playground,
	)
	satelliteFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if satelliteFormCallback.CreationMode {
		satelliteFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__SatelliteFormCallback(
				nil,
				satelliteFormCallback.playground,
			),
		}).Stage(satelliteFormCallback.playground.formStage)
		satellite := new(models.Satellite)
		FillUpForm(satellite, newFormGroup, satelliteFormCallback.playground)
		satelliteFormCallback.playground.formStage.Commit()
	}

	fillUpTree(satelliteFormCallback.playground)
}
func __gong__New__ScenarioFormCallback(
	scenario *models.Scenario,
	playground *Playground,
) (scenarioFormCallback *ScenarioFormCallback) {
	scenarioFormCallback = new(ScenarioFormCallback)
	scenarioFormCallback.playground = playground
	scenarioFormCallback.scenario = scenario

	scenarioFormCallback.CreationMode = (scenario == nil)

	return
}

type ScenarioFormCallback struct {
	scenario *models.Scenario

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (scenarioFormCallback *ScenarioFormCallback) OnSave() {

	log.Println("ScenarioFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	scenarioFormCallback.playground.formStage.Checkout()

	if scenarioFormCallback.scenario == nil {
		scenarioFormCallback.scenario = new(models.Scenario).Stage(scenarioFormCallback.playground.stageOfInterest)
	}
	scenario_ := scenarioFormCallback.scenario
	_ = scenario_

	// get the formGroup
	formGroup := scenarioFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
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

	scenarioFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Scenario](
		scenarioFormCallback.playground,
	)
	scenarioFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if scenarioFormCallback.CreationMode {
		scenarioFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__ScenarioFormCallback(
				nil,
				scenarioFormCallback.playground,
			),
		}).Stage(scenarioFormCallback.playground.formStage)
		scenario := new(models.Scenario)
		FillUpForm(scenario, newFormGroup, scenarioFormCallback.playground)
		scenarioFormCallback.playground.formStage.Commit()
	}

	fillUpTree(scenarioFormCallback.playground)
}
