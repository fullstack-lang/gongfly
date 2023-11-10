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
) (civilianairportFormCallback *CivilianAirportFormCallback) {
	civilianairportFormCallback = new(CivilianAirportFormCallback)
	civilianairportFormCallback.probe = probe
	civilianairportFormCallback.civilianairport = civilianairport

	civilianairportFormCallback.CreationMode = (civilianairport == nil)

	return
}

type CivilianAirportFormCallback struct {
	civilianairport *models.CivilianAirport

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
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

	// get the formGroup
	formGroup := civilianairportFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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

	civilianairportFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.CivilianAirport](
		civilianairportFormCallback.probe,
	)
	civilianairportFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if civilianairportFormCallback.CreationMode {
		civilianairportFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__CivilianAirportFormCallback(
				nil,
				civilianairportFormCallback.probe,
			),
		}).Stage(civilianairportFormCallback.probe.formStage)
		civilianairport := new(models.CivilianAirport)
		FillUpForm(civilianairport, newFormGroup, civilianairportFormCallback.probe)
		civilianairportFormCallback.probe.formStage.Commit()
	}

	fillUpTree(civilianairportFormCallback.probe)
}
func __gong__New__LinerFormCallback(
	liner *models.Liner,
	probe *Probe,
) (linerFormCallback *LinerFormCallback) {
	linerFormCallback = new(LinerFormCallback)
	linerFormCallback.probe = probe
	linerFormCallback.liner = liner

	linerFormCallback.CreationMode = (liner == nil)

	return
}

type LinerFormCallback struct {
	liner *models.Liner

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
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

	// get the formGroup
	formGroup := linerFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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
			FormDivSelectFieldToField(&(liner_.ReporingLine), linerFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	linerFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Liner](
		linerFormCallback.probe,
	)
	linerFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if linerFormCallback.CreationMode {
		linerFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__LinerFormCallback(
				nil,
				linerFormCallback.probe,
			),
		}).Stage(linerFormCallback.probe.formStage)
		liner := new(models.Liner)
		FillUpForm(liner, newFormGroup, linerFormCallback.probe)
		linerFormCallback.probe.formStage.Commit()
	}

	fillUpTree(linerFormCallback.probe)
}
func __gong__New__MessageFormCallback(
	message *models.Message,
	probe *Probe,
) (messageFormCallback *MessageFormCallback) {
	messageFormCallback = new(MessageFormCallback)
	messageFormCallback.probe = probe
	messageFormCallback.message = message

	messageFormCallback.CreationMode = (message == nil)

	return
}

type MessageFormCallback struct {
	message *models.Message

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
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

	// get the formGroup
	formGroup := messageFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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

	messageFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Message](
		messageFormCallback.probe,
	)
	messageFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if messageFormCallback.CreationMode {
		messageFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__MessageFormCallback(
				nil,
				messageFormCallback.probe,
			),
		}).Stage(messageFormCallback.probe.formStage)
		message := new(models.Message)
		FillUpForm(message, newFormGroup, messageFormCallback.probe)
		messageFormCallback.probe.formStage.Commit()
	}

	fillUpTree(messageFormCallback.probe)
}
func __gong__New__OpsLineFormCallback(
	opsline *models.OpsLine,
	probe *Probe,
) (opslineFormCallback *OpsLineFormCallback) {
	opslineFormCallback = new(OpsLineFormCallback)
	opslineFormCallback.probe = probe
	opslineFormCallback.opsline = opsline

	opslineFormCallback.CreationMode = (opsline == nil)

	return
}

type OpsLineFormCallback struct {
	opsline *models.OpsLine

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
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

	// get the formGroup
	formGroup := opslineFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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
			FormDivSelectFieldToField(&(opsline_.Scenario), opslineFormCallback.probe.stageOfInterest, formDiv)
		case "State":
			FormDivEnumStringFieldToField(&(opsline_.State), formDiv)
		case "Name":
			FormDivBasicFieldToField(&(opsline_.Name), formDiv)
		}
	}

	opslineFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.OpsLine](
		opslineFormCallback.probe,
	)
	opslineFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if opslineFormCallback.CreationMode {
		opslineFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__OpsLineFormCallback(
				nil,
				opslineFormCallback.probe,
			),
		}).Stage(opslineFormCallback.probe.formStage)
		opsline := new(models.OpsLine)
		FillUpForm(opsline, newFormGroup, opslineFormCallback.probe)
		opslineFormCallback.probe.formStage.Commit()
	}

	fillUpTree(opslineFormCallback.probe)
}
func __gong__New__RadarFormCallback(
	radar *models.Radar,
	probe *Probe,
) (radarFormCallback *RadarFormCallback) {
	radarFormCallback = new(RadarFormCallback)
	radarFormCallback.probe = probe
	radarFormCallback.radar = radar

	radarFormCallback.CreationMode = (radar == nil)

	return
}

type RadarFormCallback struct {
	radar *models.Radar

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
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

	// get the formGroup
	formGroup := radarFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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

	radarFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Radar](
		radarFormCallback.probe,
	)
	radarFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if radarFormCallback.CreationMode {
		radarFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__RadarFormCallback(
				nil,
				radarFormCallback.probe,
			),
		}).Stage(radarFormCallback.probe.formStage)
		radar := new(models.Radar)
		FillUpForm(radar, newFormGroup, radarFormCallback.probe)
		radarFormCallback.probe.formStage.Commit()
	}

	fillUpTree(radarFormCallback.probe)
}
func __gong__New__SatelliteFormCallback(
	satellite *models.Satellite,
	probe *Probe,
) (satelliteFormCallback *SatelliteFormCallback) {
	satelliteFormCallback = new(SatelliteFormCallback)
	satelliteFormCallback.probe = probe
	satelliteFormCallback.satellite = satellite

	satelliteFormCallback.CreationMode = (satellite == nil)

	return
}

type SatelliteFormCallback struct {
	satellite *models.Satellite

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
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

	// get the formGroup
	formGroup := satelliteFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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

	satelliteFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Satellite](
		satelliteFormCallback.probe,
	)
	satelliteFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if satelliteFormCallback.CreationMode {
		satelliteFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__SatelliteFormCallback(
				nil,
				satelliteFormCallback.probe,
			),
		}).Stage(satelliteFormCallback.probe.formStage)
		satellite := new(models.Satellite)
		FillUpForm(satellite, newFormGroup, satelliteFormCallback.probe)
		satelliteFormCallback.probe.formStage.Commit()
	}

	fillUpTree(satelliteFormCallback.probe)
}
func __gong__New__ScenarioFormCallback(
	scenario *models.Scenario,
	probe *Probe,
) (scenarioFormCallback *ScenarioFormCallback) {
	scenarioFormCallback = new(ScenarioFormCallback)
	scenarioFormCallback.probe = probe
	scenarioFormCallback.scenario = scenario

	scenarioFormCallback.CreationMode = (scenario == nil)

	return
}

type ScenarioFormCallback struct {
	scenario *models.Scenario

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
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

	// get the formGroup
	formGroup := scenarioFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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

	scenarioFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Scenario](
		scenarioFormCallback.probe,
	)
	scenarioFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if scenarioFormCallback.CreationMode {
		scenarioFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__ScenarioFormCallback(
				nil,
				scenarioFormCallback.probe,
			),
		}).Stage(scenarioFormCallback.probe.formStage)
		scenario := new(models.Scenario)
		FillUpForm(scenario, newFormGroup, scenarioFormCallback.probe)
		scenarioFormCallback.probe.formStage.Commit()
	}

	fillUpTree(scenarioFormCallback.probe)
}
