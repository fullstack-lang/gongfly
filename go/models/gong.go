// generated code - do not edit
package models

import (
	"cmp"
	"errors"
	"fmt"
	"math"
	"slices"
	"time"

	"golang.org/x/exp/maps"
)

func __Gong__Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// errUnkownEnum is returns when a value cannot match enum values
var errUnkownEnum = errors.New("unkown enum")

// needed to avoid when fmt package is not needed by generated code
var __dummy__fmt_variable fmt.Scanner

// idem for math package when not need by generated code
var __dummy_math_variable = math.E

// swagger:ignore
type __void any

// needed for creating set of instances in the stage
var __member __void

// GongStructInterface is the interface met by GongStructs
// It allows runtime reflexion of instances (without the hassle of the "reflect" package)
type GongStructInterface interface {
	GetName() (res string)
	// GetID() (res int)
	// GetFields() (res []string)
	// GetFieldStringValue(fieldName string) (res string)
}

// StageStruct enables storage of staged instances
// swagger:ignore
type StageStruct struct {
	path string

	// insertion point for definition of arrays registering instances
	CivilianAirports           map[*CivilianAirport]any
	CivilianAirports_mapString map[string]*CivilianAirport

	// insertion point for slice of pointers maps
	OnAfterCivilianAirportCreateCallback OnAfterCreateInterface[CivilianAirport]
	OnAfterCivilianAirportUpdateCallback OnAfterUpdateInterface[CivilianAirport]
	OnAfterCivilianAirportDeleteCallback OnAfterDeleteInterface[CivilianAirport]
	OnAfterCivilianAirportReadCallback   OnAfterReadInterface[CivilianAirport]

	Liners           map[*Liner]any
	Liners_mapString map[string]*Liner

	// insertion point for slice of pointers maps
	OnAfterLinerCreateCallback OnAfterCreateInterface[Liner]
	OnAfterLinerUpdateCallback OnAfterUpdateInterface[Liner]
	OnAfterLinerDeleteCallback OnAfterDeleteInterface[Liner]
	OnAfterLinerReadCallback   OnAfterReadInterface[Liner]

	Messages           map[*Message]any
	Messages_mapString map[string]*Message

	// insertion point for slice of pointers maps
	OnAfterMessageCreateCallback OnAfterCreateInterface[Message]
	OnAfterMessageUpdateCallback OnAfterUpdateInterface[Message]
	OnAfterMessageDeleteCallback OnAfterDeleteInterface[Message]
	OnAfterMessageReadCallback   OnAfterReadInterface[Message]

	OpsLines           map[*OpsLine]any
	OpsLines_mapString map[string]*OpsLine

	// insertion point for slice of pointers maps
	OnAfterOpsLineCreateCallback OnAfterCreateInterface[OpsLine]
	OnAfterOpsLineUpdateCallback OnAfterUpdateInterface[OpsLine]
	OnAfterOpsLineDeleteCallback OnAfterDeleteInterface[OpsLine]
	OnAfterOpsLineReadCallback   OnAfterReadInterface[OpsLine]

	Radars           map[*Radar]any
	Radars_mapString map[string]*Radar

	// insertion point for slice of pointers maps
	OnAfterRadarCreateCallback OnAfterCreateInterface[Radar]
	OnAfterRadarUpdateCallback OnAfterUpdateInterface[Radar]
	OnAfterRadarDeleteCallback OnAfterDeleteInterface[Radar]
	OnAfterRadarReadCallback   OnAfterReadInterface[Radar]

	Satellites           map[*Satellite]any
	Satellites_mapString map[string]*Satellite

	// insertion point for slice of pointers maps
	OnAfterSatelliteCreateCallback OnAfterCreateInterface[Satellite]
	OnAfterSatelliteUpdateCallback OnAfterUpdateInterface[Satellite]
	OnAfterSatelliteDeleteCallback OnAfterDeleteInterface[Satellite]
	OnAfterSatelliteReadCallback   OnAfterReadInterface[Satellite]

	Scenarios           map[*Scenario]any
	Scenarios_mapString map[string]*Scenario

	// insertion point for slice of pointers maps
	OnAfterScenarioCreateCallback OnAfterCreateInterface[Scenario]
	OnAfterScenarioUpdateCallback OnAfterUpdateInterface[Scenario]
	OnAfterScenarioDeleteCallback OnAfterDeleteInterface[Scenario]
	OnAfterScenarioReadCallback   OnAfterReadInterface[Scenario]

	AllModelsStructCreateCallback AllModelsStructCreateInterface

	AllModelsStructDeleteCallback AllModelsStructDeleteInterface

	BackRepo BackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback          OnInitCommitInterface
	OnInitCommitFromFrontCallback OnInitCommitInterface
	OnInitCommitFromBackCallback  OnInitCommitInterface

	// store the number of instance per gongstruct
	Map_GongStructName_InstancesNb map[string]int

	// store meta package import
	MetaPackageImportPath  string
	MetaPackageImportAlias string

	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	// map to enable docLink renaming when an identifier is renamed
	Map_DocLink_Renaming map[string]GONG__Identifier
	// the to be removed stops here
}

func (stage *StageStruct) GetType() string {
	return "github.com/fullstack-lang/gongfly/go/models"
}

type GONG__Identifier struct {
	Ident string
	Type  GONG__ExpressionType
}

type OnInitCommitInterface interface {
	BeforeCommit(stage *StageStruct)
}

// OnAfterCreateInterface callback when an instance is updated from the front
type OnAfterCreateInterface[Type Gongstruct] interface {
	OnAfterCreate(stage *StageStruct,
		instance *Type)
}

// OnAfterReadInterface callback when an instance is updated from the front
type OnAfterReadInterface[Type Gongstruct] interface {
	OnAfterRead(stage *StageStruct,
		instance *Type)
}

// OnAfterUpdateInterface callback when an instance is updated from the front
type OnAfterUpdateInterface[Type Gongstruct] interface {
	OnAfterUpdate(stage *StageStruct, old, new *Type)
}

// OnAfterDeleteInterface callback when an instance is updated from the front
type OnAfterDeleteInterface[Type Gongstruct] interface {
	OnAfterDelete(stage *StageStruct,
		staged, front *Type)
}

type BackRepoInterface interface {
	Commit(stage *StageStruct)
	Checkout(stage *StageStruct)
	Backup(stage *StageStruct, dirPath string)
	Restore(stage *StageStruct, dirPath string)
	BackupXL(stage *StageStruct, dirPath string)
	RestoreXL(stage *StageStruct, dirPath string)
	// insertion point for Commit and Checkout signatures
	CommitCivilianAirport(civilianairport *CivilianAirport)
	CheckoutCivilianAirport(civilianairport *CivilianAirport)
	CommitLiner(liner *Liner)
	CheckoutLiner(liner *Liner)
	CommitMessage(message *Message)
	CheckoutMessage(message *Message)
	CommitOpsLine(opsline *OpsLine)
	CheckoutOpsLine(opsline *OpsLine)
	CommitRadar(radar *Radar)
	CheckoutRadar(radar *Radar)
	CommitSatellite(satellite *Satellite)
	CheckoutSatellite(satellite *Satellite)
	CommitScenario(scenario *Scenario)
	CheckoutScenario(scenario *Scenario)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(path string) (stage *StageStruct) {

	stage = &StageStruct{ // insertion point for array initiatialisation
		CivilianAirports:           make(map[*CivilianAirport]any),
		CivilianAirports_mapString: make(map[string]*CivilianAirport),

		Liners:           make(map[*Liner]any),
		Liners_mapString: make(map[string]*Liner),

		Messages:           make(map[*Message]any),
		Messages_mapString: make(map[string]*Message),

		OpsLines:           make(map[*OpsLine]any),
		OpsLines_mapString: make(map[string]*OpsLine),

		Radars:           make(map[*Radar]any),
		Radars_mapString: make(map[string]*Radar),

		Satellites:           make(map[*Satellite]any),
		Satellites_mapString: make(map[string]*Satellite),

		Scenarios:           make(map[*Scenario]any),
		Scenarios_mapString: make(map[string]*Scenario),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		path: path,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here
	}

	return
}

func (stage *StageStruct) GetPath() string {
	return stage.path
}

func (stage *StageStruct) CommitWithSuspendedCallbacks() {

	tmp := stage.OnInitCommitFromBackCallback
	stage.OnInitCommitFromBackCallback = nil
	stage.Commit()
	stage.OnInitCommitFromBackCallback = tmp
}

func (stage *StageStruct) Commit() {
	stage.ComputeReverseMaps()

	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}

	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["CivilianAirport"] = len(stage.CivilianAirports)
	stage.Map_GongStructName_InstancesNb["Liner"] = len(stage.Liners)
	stage.Map_GongStructName_InstancesNb["Message"] = len(stage.Messages)
	stage.Map_GongStructName_InstancesNb["OpsLine"] = len(stage.OpsLines)
	stage.Map_GongStructName_InstancesNb["Radar"] = len(stage.Radars)
	stage.Map_GongStructName_InstancesNb["Satellite"] = len(stage.Satellites)
	stage.Map_GongStructName_InstancesNb["Scenario"] = len(stage.Scenarios)

}

func (stage *StageStruct) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["CivilianAirport"] = len(stage.CivilianAirports)
	stage.Map_GongStructName_InstancesNb["Liner"] = len(stage.Liners)
	stage.Map_GongStructName_InstancesNb["Message"] = len(stage.Messages)
	stage.Map_GongStructName_InstancesNb["OpsLine"] = len(stage.OpsLines)
	stage.Map_GongStructName_InstancesNb["Radar"] = len(stage.Radars)
	stage.Map_GongStructName_InstancesNb["Satellite"] = len(stage.Satellites)
	stage.Map_GongStructName_InstancesNb["Scenario"] = len(stage.Scenarios)

}

// backup generates backup files in the dirPath
func (stage *StageStruct) Backup(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Backup(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) Restore(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Restore(stage, dirPath)
	}
}

// backup generates backup files in the dirPath
func (stage *StageStruct) BackupXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.BackupXL(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) RestoreXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.RestoreXL(stage, dirPath)
	}
}

// insertion point for cumulative sub template with model space calls
// Stage puts civilianairport to the model stage
func (civilianairport *CivilianAirport) Stage(stage *StageStruct) *CivilianAirport {
	stage.CivilianAirports[civilianairport] = __member
	stage.CivilianAirports_mapString[civilianairport.Name] = civilianairport

	return civilianairport
}

// Unstage removes civilianairport off the model stage
func (civilianairport *CivilianAirport) Unstage(stage *StageStruct) *CivilianAirport {
	delete(stage.CivilianAirports, civilianairport)
	delete(stage.CivilianAirports_mapString, civilianairport.Name)
	return civilianairport
}

// UnstageVoid removes civilianairport off the model stage
func (civilianairport *CivilianAirport) UnstageVoid(stage *StageStruct) {
	delete(stage.CivilianAirports, civilianairport)
	delete(stage.CivilianAirports_mapString, civilianairport.Name)
}

// commit civilianairport to the back repo (if it is already staged)
func (civilianairport *CivilianAirport) Commit(stage *StageStruct) *CivilianAirport {
	if _, ok := stage.CivilianAirports[civilianairport]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCivilianAirport(civilianairport)
		}
	}
	return civilianairport
}

func (civilianairport *CivilianAirport) CommitVoid(stage *StageStruct) {
	civilianairport.Commit(stage)
}

// Checkout civilianairport to the back repo (if it is already staged)
func (civilianairport *CivilianAirport) Checkout(stage *StageStruct) *CivilianAirport {
	if _, ok := stage.CivilianAirports[civilianairport]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutCivilianAirport(civilianairport)
		}
	}
	return civilianairport
}

// for satisfaction of GongStruct interface
func (civilianairport *CivilianAirport) GetName() (res string) {
	return civilianairport.Name
}

// Stage puts liner to the model stage
func (liner *Liner) Stage(stage *StageStruct) *Liner {
	stage.Liners[liner] = __member
	stage.Liners_mapString[liner.Name] = liner

	return liner
}

// Unstage removes liner off the model stage
func (liner *Liner) Unstage(stage *StageStruct) *Liner {
	delete(stage.Liners, liner)
	delete(stage.Liners_mapString, liner.Name)
	return liner
}

// UnstageVoid removes liner off the model stage
func (liner *Liner) UnstageVoid(stage *StageStruct) {
	delete(stage.Liners, liner)
	delete(stage.Liners_mapString, liner.Name)
}

// commit liner to the back repo (if it is already staged)
func (liner *Liner) Commit(stage *StageStruct) *Liner {
	if _, ok := stage.Liners[liner]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLiner(liner)
		}
	}
	return liner
}

func (liner *Liner) CommitVoid(stage *StageStruct) {
	liner.Commit(stage)
}

// Checkout liner to the back repo (if it is already staged)
func (liner *Liner) Checkout(stage *StageStruct) *Liner {
	if _, ok := stage.Liners[liner]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutLiner(liner)
		}
	}
	return liner
}

// for satisfaction of GongStruct interface
func (liner *Liner) GetName() (res string) {
	return liner.Name
}

// Stage puts message to the model stage
func (message *Message) Stage(stage *StageStruct) *Message {
	stage.Messages[message] = __member
	stage.Messages_mapString[message.Name] = message

	return message
}

// Unstage removes message off the model stage
func (message *Message) Unstage(stage *StageStruct) *Message {
	delete(stage.Messages, message)
	delete(stage.Messages_mapString, message.Name)
	return message
}

// UnstageVoid removes message off the model stage
func (message *Message) UnstageVoid(stage *StageStruct) {
	delete(stage.Messages, message)
	delete(stage.Messages_mapString, message.Name)
}

// commit message to the back repo (if it is already staged)
func (message *Message) Commit(stage *StageStruct) *Message {
	if _, ok := stage.Messages[message]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitMessage(message)
		}
	}
	return message
}

func (message *Message) CommitVoid(stage *StageStruct) {
	message.Commit(stage)
}

// Checkout message to the back repo (if it is already staged)
func (message *Message) Checkout(stage *StageStruct) *Message {
	if _, ok := stage.Messages[message]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutMessage(message)
		}
	}
	return message
}

// for satisfaction of GongStruct interface
func (message *Message) GetName() (res string) {
	return message.Name
}

// Stage puts opsline to the model stage
func (opsline *OpsLine) Stage(stage *StageStruct) *OpsLine {
	stage.OpsLines[opsline] = __member
	stage.OpsLines_mapString[opsline.Name] = opsline

	return opsline
}

// Unstage removes opsline off the model stage
func (opsline *OpsLine) Unstage(stage *StageStruct) *OpsLine {
	delete(stage.OpsLines, opsline)
	delete(stage.OpsLines_mapString, opsline.Name)
	return opsline
}

// UnstageVoid removes opsline off the model stage
func (opsline *OpsLine) UnstageVoid(stage *StageStruct) {
	delete(stage.OpsLines, opsline)
	delete(stage.OpsLines_mapString, opsline.Name)
}

// commit opsline to the back repo (if it is already staged)
func (opsline *OpsLine) Commit(stage *StageStruct) *OpsLine {
	if _, ok := stage.OpsLines[opsline]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitOpsLine(opsline)
		}
	}
	return opsline
}

func (opsline *OpsLine) CommitVoid(stage *StageStruct) {
	opsline.Commit(stage)
}

// Checkout opsline to the back repo (if it is already staged)
func (opsline *OpsLine) Checkout(stage *StageStruct) *OpsLine {
	if _, ok := stage.OpsLines[opsline]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutOpsLine(opsline)
		}
	}
	return opsline
}

// for satisfaction of GongStruct interface
func (opsline *OpsLine) GetName() (res string) {
	return opsline.Name
}

// Stage puts radar to the model stage
func (radar *Radar) Stage(stage *StageStruct) *Radar {
	stage.Radars[radar] = __member
	stage.Radars_mapString[radar.Name] = radar

	return radar
}

// Unstage removes radar off the model stage
func (radar *Radar) Unstage(stage *StageStruct) *Radar {
	delete(stage.Radars, radar)
	delete(stage.Radars_mapString, radar.Name)
	return radar
}

// UnstageVoid removes radar off the model stage
func (radar *Radar) UnstageVoid(stage *StageStruct) {
	delete(stage.Radars, radar)
	delete(stage.Radars_mapString, radar.Name)
}

// commit radar to the back repo (if it is already staged)
func (radar *Radar) Commit(stage *StageStruct) *Radar {
	if _, ok := stage.Radars[radar]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitRadar(radar)
		}
	}
	return radar
}

func (radar *Radar) CommitVoid(stage *StageStruct) {
	radar.Commit(stage)
}

// Checkout radar to the back repo (if it is already staged)
func (radar *Radar) Checkout(stage *StageStruct) *Radar {
	if _, ok := stage.Radars[radar]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutRadar(radar)
		}
	}
	return radar
}

// for satisfaction of GongStruct interface
func (radar *Radar) GetName() (res string) {
	return radar.Name
}

// Stage puts satellite to the model stage
func (satellite *Satellite) Stage(stage *StageStruct) *Satellite {
	stage.Satellites[satellite] = __member
	stage.Satellites_mapString[satellite.Name] = satellite

	return satellite
}

// Unstage removes satellite off the model stage
func (satellite *Satellite) Unstage(stage *StageStruct) *Satellite {
	delete(stage.Satellites, satellite)
	delete(stage.Satellites_mapString, satellite.Name)
	return satellite
}

// UnstageVoid removes satellite off the model stage
func (satellite *Satellite) UnstageVoid(stage *StageStruct) {
	delete(stage.Satellites, satellite)
	delete(stage.Satellites_mapString, satellite.Name)
}

// commit satellite to the back repo (if it is already staged)
func (satellite *Satellite) Commit(stage *StageStruct) *Satellite {
	if _, ok := stage.Satellites[satellite]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSatellite(satellite)
		}
	}
	return satellite
}

func (satellite *Satellite) CommitVoid(stage *StageStruct) {
	satellite.Commit(stage)
}

// Checkout satellite to the back repo (if it is already staged)
func (satellite *Satellite) Checkout(stage *StageStruct) *Satellite {
	if _, ok := stage.Satellites[satellite]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSatellite(satellite)
		}
	}
	return satellite
}

// for satisfaction of GongStruct interface
func (satellite *Satellite) GetName() (res string) {
	return satellite.Name
}

// Stage puts scenario to the model stage
func (scenario *Scenario) Stage(stage *StageStruct) *Scenario {
	stage.Scenarios[scenario] = __member
	stage.Scenarios_mapString[scenario.Name] = scenario

	return scenario
}

// Unstage removes scenario off the model stage
func (scenario *Scenario) Unstage(stage *StageStruct) *Scenario {
	delete(stage.Scenarios, scenario)
	delete(stage.Scenarios_mapString, scenario.Name)
	return scenario
}

// UnstageVoid removes scenario off the model stage
func (scenario *Scenario) UnstageVoid(stage *StageStruct) {
	delete(stage.Scenarios, scenario)
	delete(stage.Scenarios_mapString, scenario.Name)
}

// commit scenario to the back repo (if it is already staged)
func (scenario *Scenario) Commit(stage *StageStruct) *Scenario {
	if _, ok := stage.Scenarios[scenario]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitScenario(scenario)
		}
	}
	return scenario
}

func (scenario *Scenario) CommitVoid(stage *StageStruct) {
	scenario.Commit(stage)
}

// Checkout scenario to the back repo (if it is already staged)
func (scenario *Scenario) Checkout(stage *StageStruct) *Scenario {
	if _, ok := stage.Scenarios[scenario]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutScenario(scenario)
		}
	}
	return scenario
}

// for satisfaction of GongStruct interface
func (scenario *Scenario) GetName() (res string) {
	return scenario.Name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMCivilianAirport(CivilianAirport *CivilianAirport)
	CreateORMLiner(Liner *Liner)
	CreateORMMessage(Message *Message)
	CreateORMOpsLine(OpsLine *OpsLine)
	CreateORMRadar(Radar *Radar)
	CreateORMSatellite(Satellite *Satellite)
	CreateORMScenario(Scenario *Scenario)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMCivilianAirport(CivilianAirport *CivilianAirport)
	DeleteORMLiner(Liner *Liner)
	DeleteORMMessage(Message *Message)
	DeleteORMOpsLine(OpsLine *OpsLine)
	DeleteORMRadar(Radar *Radar)
	DeleteORMSatellite(Satellite *Satellite)
	DeleteORMScenario(Scenario *Scenario)
}

func (stage *StageStruct) Reset() { // insertion point for array reset
	stage.CivilianAirports = make(map[*CivilianAirport]any)
	stage.CivilianAirports_mapString = make(map[string]*CivilianAirport)

	stage.Liners = make(map[*Liner]any)
	stage.Liners_mapString = make(map[string]*Liner)

	stage.Messages = make(map[*Message]any)
	stage.Messages_mapString = make(map[string]*Message)

	stage.OpsLines = make(map[*OpsLine]any)
	stage.OpsLines_mapString = make(map[string]*OpsLine)

	stage.Radars = make(map[*Radar]any)
	stage.Radars_mapString = make(map[string]*Radar)

	stage.Satellites = make(map[*Satellite]any)
	stage.Satellites_mapString = make(map[string]*Satellite)

	stage.Scenarios = make(map[*Scenario]any)
	stage.Scenarios_mapString = make(map[string]*Scenario)

}

func (stage *StageStruct) Nil() { // insertion point for array nil
	stage.CivilianAirports = nil
	stage.CivilianAirports_mapString = nil

	stage.Liners = nil
	stage.Liners_mapString = nil

	stage.Messages = nil
	stage.Messages_mapString = nil

	stage.OpsLines = nil
	stage.OpsLines_mapString = nil

	stage.Radars = nil
	stage.Radars_mapString = nil

	stage.Satellites = nil
	stage.Satellites_mapString = nil

	stage.Scenarios = nil
	stage.Scenarios_mapString = nil

}

func (stage *StageStruct) Unstage() { // insertion point for array nil
	for civilianairport := range stage.CivilianAirports {
		civilianairport.Unstage(stage)
	}

	for liner := range stage.Liners {
		liner.Unstage(stage)
	}

	for message := range stage.Messages {
		message.Unstage(stage)
	}

	for opsline := range stage.OpsLines {
		opsline.Unstage(stage)
	}

	for radar := range stage.Radars {
		radar.Unstage(stage)
	}

	for satellite := range stage.Satellites {
		satellite.Unstage(stage)
	}

	for scenario := range stage.Scenarios {
		scenario.Unstage(stage)
	}

}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type Gongstruct interface {

}

type GongtructBasicField interface {
	int | float64 | bool | string | time.Time | time.Duration
}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type PointerToGongstruct interface {
	GetName() string
	CommitVoid(*StageStruct)
	UnstageVoid(stage *StageStruct)
	comparable
}

func CompareGongstructByName[T PointerToGongstruct](a, b T) int {
	return cmp.Compare(a.GetName(), b.GetName())
}

func SortGongstructSetByName[T PointerToGongstruct](set map[T]any) (sortedSlice []T) {

	sortedSlice = maps.Keys(set)
	slices.SortFunc(sortedSlice, CompareGongstructByName)

	return
}

func GetGongstrucsSorted[T PointerToGongstruct](stage *StageStruct) (sortedSlice []T) {

	set := GetGongstructInstancesSetFromPointerType[T](stage)
	sortedSlice = SortGongstructSetByName(*set)

	return
}

type GongstructSet interface {
	map[any]any
}

type GongstructMapString interface {
	map[any]any
}

// GongGetSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetSet[Type GongstructSet](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[*CivilianAirport]any:
		return any(&stage.CivilianAirports).(*Type)
	case map[*Liner]any:
		return any(&stage.Liners).(*Type)
	case map[*Message]any:
		return any(&stage.Messages).(*Type)
	case map[*OpsLine]any:
		return any(&stage.OpsLines).(*Type)
	case map[*Radar]any:
		return any(&stage.Radars).(*Type)
	case map[*Satellite]any:
		return any(&stage.Satellites).(*Type)
	case map[*Scenario]any:
		return any(&stage.Scenarios).(*Type)
	default:
		return nil
	}
}

// GongGetMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetMap[Type GongstructMapString](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[string]*CivilianAirport:
		return any(&stage.CivilianAirports_mapString).(*Type)
	case map[string]*Liner:
		return any(&stage.Liners_mapString).(*Type)
	case map[string]*Message:
		return any(&stage.Messages_mapString).(*Type)
	case map[string]*OpsLine:
		return any(&stage.OpsLines_mapString).(*Type)
	case map[string]*Radar:
		return any(&stage.Radars_mapString).(*Type)
	case map[string]*Satellite:
		return any(&stage.Satellites_mapString).(*Type)
	case map[string]*Scenario:
		return any(&stage.Scenarios_mapString).(*Type)
	default:
		return nil
	}
}

// GetGongstructInstancesSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSet[Type Gongstruct](stage *StageStruct) *map[*Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case CivilianAirport:
		return any(&stage.CivilianAirports).(*map[*Type]any)
	case Liner:
		return any(&stage.Liners).(*map[*Type]any)
	case Message:
		return any(&stage.Messages).(*map[*Type]any)
	case OpsLine:
		return any(&stage.OpsLines).(*map[*Type]any)
	case Radar:
		return any(&stage.Radars).(*map[*Type]any)
	case Satellite:
		return any(&stage.Satellites).(*map[*Type]any)
	case Scenario:
		return any(&stage.Scenarios).(*map[*Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesSetFromPointerType returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSetFromPointerType[Type PointerToGongstruct](stage *StageStruct) *map[Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *CivilianAirport:
		return any(&stage.CivilianAirports).(*map[Type]any)
	case *Liner:
		return any(&stage.Liners).(*map[Type]any)
	case *Message:
		return any(&stage.Messages).(*map[Type]any)
	case *OpsLine:
		return any(&stage.OpsLines).(*map[Type]any)
	case *Radar:
		return any(&stage.Radars).(*map[Type]any)
	case *Satellite:
		return any(&stage.Satellites).(*map[Type]any)
	case *Scenario:
		return any(&stage.Scenarios).(*map[Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GetGongstructInstancesMap[Type Gongstruct](stage *StageStruct) *map[string]*Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case CivilianAirport:
		return any(&stage.CivilianAirports_mapString).(*map[string]*Type)
	case Liner:
		return any(&stage.Liners_mapString).(*map[string]*Type)
	case Message:
		return any(&stage.Messages_mapString).(*map[string]*Type)
	case OpsLine:
		return any(&stage.OpsLines_mapString).(*map[string]*Type)
	case Radar:
		return any(&stage.Radars_mapString).(*map[string]*Type)
	case Satellite:
		return any(&stage.Satellites_mapString).(*map[string]*Type)
	case Scenario:
		return any(&stage.Scenarios_mapString).(*map[string]*Type)
	default:
		return nil
	}
}

// GetAssociationName is a generic function that returns an instance of Type
// where each association is filled with an instance whose name is the name of the association
//
// This function can be handy for generating navigation function that are refactorable
func GetAssociationName[Type Gongstruct]() *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for instance with special fields
	case CivilianAirport:
		return any(&CivilianAirport{
			// Initialisation of associations
		}).(*Type)
	case Liner:
		return any(&Liner{
			// Initialisation of associations
			// field is initialized with an instance of OpsLine with the name of the field
			ReporingLine: &OpsLine{Name: "ReporingLine"},
		}).(*Type)
	case Message:
		return any(&Message{
			// Initialisation of associations
		}).(*Type)
	case OpsLine:
		return any(&OpsLine{
			// Initialisation of associations
			// field is initialized with an instance of Scenario with the name of the field
			Scenario: &Scenario{Name: "Scenario"},
		}).(*Type)
	case Radar:
		return any(&Radar{
			// Initialisation of associations
		}).(*Type)
	case Satellite:
		return any(&Satellite{
			// Initialisation of associations
		}).(*Type)
	case Scenario:
		return any(&Scenario{
			// Initialisation of associations
		}).(*Type)
	default:
		return nil
	}
}

// GetPointerReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..1) that is a pointer from one staged Gongstruct (type Start)
// instances to another (type End)
//
// The function provides a map with keys as instances of End and values to arrays of *Start
// the map is construed by iterating over all Start instances and populationg keys with End instances
// and values with slice of Start instances
func GetPointerReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End][]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of CivilianAirport
	case CivilianAirport:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Liner
	case Liner:
		switch fieldname {
		// insertion point for per direct association field
		case "ReporingLine":
			res := make(map[*OpsLine][]*Liner)
			for liner := range stage.Liners {
				if liner.ReporingLine != nil {
					opsline_ := liner.ReporingLine
					var liners []*Liner
					_, ok := res[opsline_]
					if ok {
						liners = res[opsline_]
					} else {
						liners = make([]*Liner, 0)
					}
					liners = append(liners, liner)
					res[opsline_] = liners
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Message
	case Message:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of OpsLine
	case OpsLine:
		switch fieldname {
		// insertion point for per direct association field
		case "Scenario":
			res := make(map[*Scenario][]*OpsLine)
			for opsline := range stage.OpsLines {
				if opsline.Scenario != nil {
					scenario_ := opsline.Scenario
					var opslines []*OpsLine
					_, ok := res[scenario_]
					if ok {
						opslines = res[scenario_]
					} else {
						opslines = make([]*OpsLine, 0)
					}
					opslines = append(opslines, opsline)
					res[scenario_] = opslines
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Radar
	case Radar:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Satellite
	case Satellite:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Scenario
	case Scenario:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GetSliceOfPointersReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..N) between one staged Gongstruct instances and many others
//
// The function provides a map with keys as instances of End and values to *Start instances
// the map is construed by iterating over all Start instances and populating keys with End instances
// and values with the Start instances
func GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of CivilianAirport
	case CivilianAirport:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Liner
	case Liner:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Message
	case Message:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of OpsLine
	case OpsLine:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Radar
	case Radar:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Satellite
	case Satellite:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Scenario
	case Scenario:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GetGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetGongstructName[Type Gongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case CivilianAirport:
		res = "CivilianAirport"
	case Liner:
		res = "Liner"
	case Message:
		res = "Message"
	case OpsLine:
		res = "OpsLine"
	case Radar:
		res = "Radar"
	case Satellite:
		res = "Satellite"
	case Scenario:
		res = "Scenario"
	}
	return res
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type PointerToGongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *CivilianAirport:
		res = "CivilianAirport"
	case *Liner:
		res = "Liner"
	case *Message:
		res = "Message"
	case *OpsLine:
		res = "OpsLine"
	case *Radar:
		res = "Radar"
	case *Satellite:
		res = "Satellite"
	case *Scenario:
		res = "Scenario"
	}
	return res
}

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case CivilianAirport:
		res = []string{"Lat", "Lng", "Name"}
	case Liner:
		res = []string{"Name", "Lat", "Lng", "Heading", "Level", "Speed", "State", "TargetHeading", "TargetLocationLat", "TargetLocationLng", "DistanceToTarget", "MaxRotationalSpeed", "VerticalSpeed", "Timestampstring", "ReporingLine"}
	case Message:
		res = []string{"Lat", "Lng", "Heading", "Level", "Speed", "State", "Name", "TargetLocationLat", "TargetLocationLng", "DistanceToTarget", "Timestampstring", "DurationSinceSimulationStart", "Timestampstartstring", "Source", "Destination", "Content", "About_string", "Display"}
	case OpsLine:
		res = []string{"IsTransmitting", "TransmissionMessage", "IsTransmittingBackward", "TransmissionMessageBackward", "Scenario", "State", "Name"}
	case Radar:
		res = []string{"State", "Name", "Lat", "Lng", "Range"}
	case Satellite:
		res = []string{"Name", "Line1", "Line2", "Lat", "Lng", "Heading", "Level", "Speed", "VerticalSpeed", "Timestampstring"}
	case Scenario:
		res = []string{"Name", "Lat", "Lng", "ZoomLevel", "MessageVisualSpeed"}
	}
	return
}

type ReverseField struct {
	GongstructName string
	Fieldname      string
}

func GetReverseFields[Type Gongstruct]() (res []ReverseField) {

	res = make([]ReverseField, 0)

	var ret Type

	switch any(ret).(type) {

	// insertion point for generic get gongstruct name
	case CivilianAirport:
		var rf ReverseField
		_ = rf
	case Liner:
		var rf ReverseField
		_ = rf
	case Message:
		var rf ReverseField
		_ = rf
	case OpsLine:
		var rf ReverseField
		_ = rf
	case Radar:
		var rf ReverseField
		_ = rf
	case Satellite:
		var rf ReverseField
		_ = rf
	case Scenario:
		var rf ReverseField
		_ = rf
	}
	return
}

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *CivilianAirport:
		res = []string{"Lat", "Lng", "Name"}
	case *Liner:
		res = []string{"Name", "Lat", "Lng", "Heading", "Level", "Speed", "State", "TargetHeading", "TargetLocationLat", "TargetLocationLng", "DistanceToTarget", "MaxRotationalSpeed", "VerticalSpeed", "Timestampstring", "ReporingLine"}
	case *Message:
		res = []string{"Lat", "Lng", "Heading", "Level", "Speed", "State", "Name", "TargetLocationLat", "TargetLocationLng", "DistanceToTarget", "Timestampstring", "DurationSinceSimulationStart", "Timestampstartstring", "Source", "Destination", "Content", "About_string", "Display"}
	case *OpsLine:
		res = []string{"IsTransmitting", "TransmissionMessage", "IsTransmittingBackward", "TransmissionMessageBackward", "Scenario", "State", "Name"}
	case *Radar:
		res = []string{"State", "Name", "Lat", "Lng", "Range"}
	case *Satellite:
		res = []string{"Name", "Line1", "Line2", "Lat", "Lng", "Heading", "Level", "Speed", "VerticalSpeed", "Timestampstring"}
	case *Scenario:
		res = []string{"Name", "Lat", "Lng", "ZoomLevel", "MessageVisualSpeed"}
	}
	return
}

func GetFieldStringValueFromPointer[Type PointerToGongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case *CivilianAirport:
		switch fieldName {
		// string value of fields
		case "Lat":
			res = fmt.Sprintf("%f", inferedInstance.Lat)
		case "Lng":
			res = fmt.Sprintf("%f", inferedInstance.Lng)
		case "Name":
			res = inferedInstance.Name
		}
	case *Liner:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Lat":
			res = fmt.Sprintf("%f", inferedInstance.Lat)
		case "Lng":
			res = fmt.Sprintf("%f", inferedInstance.Lng)
		case "Heading":
			res = fmt.Sprintf("%f", inferedInstance.Heading)
		case "Level":
			res = fmt.Sprintf("%f", inferedInstance.Level)
		case "Speed":
			res = fmt.Sprintf("%f", inferedInstance.Speed)
		case "State":
			enum := inferedInstance.State
			res = enum.ToCodeString()
		case "TargetHeading":
			res = fmt.Sprintf("%f", inferedInstance.TargetHeading)
		case "TargetLocationLat":
			res = fmt.Sprintf("%f", inferedInstance.TargetLocationLat)
		case "TargetLocationLng":
			res = fmt.Sprintf("%f", inferedInstance.TargetLocationLng)
		case "DistanceToTarget":
			res = fmt.Sprintf("%f", inferedInstance.DistanceToTarget)
		case "MaxRotationalSpeed":
			res = fmt.Sprintf("%f", inferedInstance.MaxRotationalSpeed)
		case "VerticalSpeed":
			res = fmt.Sprintf("%f", inferedInstance.VerticalSpeed)
		case "Timestampstring":
			res = inferedInstance.Timestampstring
		case "ReporingLine":
			if inferedInstance.ReporingLine != nil {
				res = inferedInstance.ReporingLine.Name
			}
		}
	case *Message:
		switch fieldName {
		// string value of fields
		case "Lat":
			res = fmt.Sprintf("%f", inferedInstance.Lat)
		case "Lng":
			res = fmt.Sprintf("%f", inferedInstance.Lng)
		case "Heading":
			res = fmt.Sprintf("%f", inferedInstance.Heading)
		case "Level":
			res = fmt.Sprintf("%f", inferedInstance.Level)
		case "Speed":
			res = fmt.Sprintf("%f", inferedInstance.Speed)
		case "State":
			enum := inferedInstance.State
			res = enum.ToCodeString()
		case "Name":
			res = inferedInstance.Name
		case "TargetLocationLat":
			res = fmt.Sprintf("%f", inferedInstance.TargetLocationLat)
		case "TargetLocationLng":
			res = fmt.Sprintf("%f", inferedInstance.TargetLocationLng)
		case "DistanceToTarget":
			res = fmt.Sprintf("%f", inferedInstance.DistanceToTarget)
		case "Timestampstring":
			res = inferedInstance.Timestampstring
		case "DurationSinceSimulationStart":
			if math.Abs(inferedInstance.DurationSinceSimulationStart.Hours()) >= 24 {
				days := __Gong__Abs(int(int(inferedInstance.DurationSinceSimulationStart.Hours()) / 24))
				months := int(days / 31)
				days = days - months*31

				remainingHours := int(inferedInstance.DurationSinceSimulationStart.Hours()) % 24
				remainingMinutes := int(inferedInstance.DurationSinceSimulationStart.Minutes()) % 60
				remainingSeconds := int(inferedInstance.DurationSinceSimulationStart.Seconds()) % 60

				if inferedInstance.DurationSinceSimulationStart.Hours() < 0 {
					res = "- "
				}

				if months > 0 {
					if months > 1 {
						res = res + fmt.Sprintf("%d months", months)
					} else {
						res = res + fmt.Sprintf("%d month", months)
					}
				}
				if days > 0 {
					if months != 0 {
						res = res + ", "
					}
					if days > 1 {
						res = res + fmt.Sprintf("%d days", days)
					} else {
						res = res + fmt.Sprintf("%d day", days)
					}

				}
				if remainingHours != 0 || remainingMinutes != 0 || remainingSeconds != 0 {
					if days != 0 || (days == 0 && months != 0) {
						res = res + ", "
					}
					res = res + fmt.Sprintf("%d hours, %d minutes, %d seconds\n", remainingHours, remainingMinutes, remainingSeconds)
				}
			} else {
				res = fmt.Sprintf("%s\n", inferedInstance.DurationSinceSimulationStart.String())
			}
		case "Timestampstartstring":
			res = inferedInstance.Timestampstartstring
		case "Source":
			res = inferedInstance.Source
		case "Destination":
			res = inferedInstance.Destination
		case "Content":
			res = inferedInstance.Content
		case "About_string":
			res = inferedInstance.About_string
		case "Display":
			res = fmt.Sprintf("%t", inferedInstance.Display)
		}
	case *OpsLine:
		switch fieldName {
		// string value of fields
		case "IsTransmitting":
			res = fmt.Sprintf("%t", inferedInstance.IsTransmitting)
		case "TransmissionMessage":
			res = inferedInstance.TransmissionMessage
		case "IsTransmittingBackward":
			res = fmt.Sprintf("%t", inferedInstance.IsTransmittingBackward)
		case "TransmissionMessageBackward":
			res = inferedInstance.TransmissionMessageBackward
		case "Scenario":
			if inferedInstance.Scenario != nil {
				res = inferedInstance.Scenario.Name
			}
		case "State":
			enum := inferedInstance.State
			res = enum.ToCodeString()
		case "Name":
			res = inferedInstance.Name
		}
	case *Radar:
		switch fieldName {
		// string value of fields
		case "State":
			enum := inferedInstance.State
			res = enum.ToCodeString()
		case "Name":
			res = inferedInstance.Name
		case "Lat":
			res = fmt.Sprintf("%f", inferedInstance.Lat)
		case "Lng":
			res = fmt.Sprintf("%f", inferedInstance.Lng)
		case "Range":
			res = fmt.Sprintf("%f", inferedInstance.Range)
		}
	case *Satellite:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Line1":
			res = inferedInstance.Line1
		case "Line2":
			res = inferedInstance.Line2
		case "Lat":
			res = fmt.Sprintf("%f", inferedInstance.Lat)
		case "Lng":
			res = fmt.Sprintf("%f", inferedInstance.Lng)
		case "Heading":
			res = fmt.Sprintf("%f", inferedInstance.Heading)
		case "Level":
			res = fmt.Sprintf("%f", inferedInstance.Level)
		case "Speed":
			res = fmt.Sprintf("%f", inferedInstance.Speed)
		case "VerticalSpeed":
			res = fmt.Sprintf("%f", inferedInstance.VerticalSpeed)
		case "Timestampstring":
			res = inferedInstance.Timestampstring
		}
	case *Scenario:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Lat":
			res = fmt.Sprintf("%f", inferedInstance.Lat)
		case "Lng":
			res = fmt.Sprintf("%f", inferedInstance.Lng)
		case "ZoomLevel":
			res = fmt.Sprintf("%f", inferedInstance.ZoomLevel)
		case "MessageVisualSpeed":
			res = fmt.Sprintf("%f", inferedInstance.MessageVisualSpeed)
		}
	default:
		_ = inferedInstance
	}
	return
}

func GetFieldStringValue[Type Gongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case CivilianAirport:
		switch fieldName {
		// string value of fields
		case "Lat":
			res = fmt.Sprintf("%f", inferedInstance.Lat)
		case "Lng":
			res = fmt.Sprintf("%f", inferedInstance.Lng)
		case "Name":
			res = inferedInstance.Name
		}
	case Liner:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Lat":
			res = fmt.Sprintf("%f", inferedInstance.Lat)
		case "Lng":
			res = fmt.Sprintf("%f", inferedInstance.Lng)
		case "Heading":
			res = fmt.Sprintf("%f", inferedInstance.Heading)
		case "Level":
			res = fmt.Sprintf("%f", inferedInstance.Level)
		case "Speed":
			res = fmt.Sprintf("%f", inferedInstance.Speed)
		case "State":
			enum := inferedInstance.State
			res = enum.ToCodeString()
		case "TargetHeading":
			res = fmt.Sprintf("%f", inferedInstance.TargetHeading)
		case "TargetLocationLat":
			res = fmt.Sprintf("%f", inferedInstance.TargetLocationLat)
		case "TargetLocationLng":
			res = fmt.Sprintf("%f", inferedInstance.TargetLocationLng)
		case "DistanceToTarget":
			res = fmt.Sprintf("%f", inferedInstance.DistanceToTarget)
		case "MaxRotationalSpeed":
			res = fmt.Sprintf("%f", inferedInstance.MaxRotationalSpeed)
		case "VerticalSpeed":
			res = fmt.Sprintf("%f", inferedInstance.VerticalSpeed)
		case "Timestampstring":
			res = inferedInstance.Timestampstring
		case "ReporingLine":
			if inferedInstance.ReporingLine != nil {
				res = inferedInstance.ReporingLine.Name
			}
		}
	case Message:
		switch fieldName {
		// string value of fields
		case "Lat":
			res = fmt.Sprintf("%f", inferedInstance.Lat)
		case "Lng":
			res = fmt.Sprintf("%f", inferedInstance.Lng)
		case "Heading":
			res = fmt.Sprintf("%f", inferedInstance.Heading)
		case "Level":
			res = fmt.Sprintf("%f", inferedInstance.Level)
		case "Speed":
			res = fmt.Sprintf("%f", inferedInstance.Speed)
		case "State":
			enum := inferedInstance.State
			res = enum.ToCodeString()
		case "Name":
			res = inferedInstance.Name
		case "TargetLocationLat":
			res = fmt.Sprintf("%f", inferedInstance.TargetLocationLat)
		case "TargetLocationLng":
			res = fmt.Sprintf("%f", inferedInstance.TargetLocationLng)
		case "DistanceToTarget":
			res = fmt.Sprintf("%f", inferedInstance.DistanceToTarget)
		case "Timestampstring":
			res = inferedInstance.Timestampstring
		case "DurationSinceSimulationStart":
			if math.Abs(inferedInstance.DurationSinceSimulationStart.Hours()) >= 24 {
				days := __Gong__Abs(int(int(inferedInstance.DurationSinceSimulationStart.Hours()) / 24))
				months := int(days / 31)
				days = days - months*31

				remainingHours := int(inferedInstance.DurationSinceSimulationStart.Hours()) % 24
				remainingMinutes := int(inferedInstance.DurationSinceSimulationStart.Minutes()) % 60
				remainingSeconds := int(inferedInstance.DurationSinceSimulationStart.Seconds()) % 60

				if inferedInstance.DurationSinceSimulationStart.Hours() < 0 {
					res = "- "
				}

				if months > 0 {
					if months > 1 {
						res = res + fmt.Sprintf("%d months", months)
					} else {
						res = res + fmt.Sprintf("%d month", months)
					}
				}
				if days > 0 {
					if months != 0 {
						res = res + ", "
					}
					if days > 1 {
						res = res + fmt.Sprintf("%d days", days)
					} else {
						res = res + fmt.Sprintf("%d day", days)
					}

				}
				if remainingHours != 0 || remainingMinutes != 0 || remainingSeconds != 0 {
					if days != 0 || (days == 0 && months != 0) {
						res = res + ", "
					}
					res = res + fmt.Sprintf("%d hours, %d minutes, %d seconds\n", remainingHours, remainingMinutes, remainingSeconds)
				}
			} else {
				res = fmt.Sprintf("%s\n", inferedInstance.DurationSinceSimulationStart.String())
			}
		case "Timestampstartstring":
			res = inferedInstance.Timestampstartstring
		case "Source":
			res = inferedInstance.Source
		case "Destination":
			res = inferedInstance.Destination
		case "Content":
			res = inferedInstance.Content
		case "About_string":
			res = inferedInstance.About_string
		case "Display":
			res = fmt.Sprintf("%t", inferedInstance.Display)
		}
	case OpsLine:
		switch fieldName {
		// string value of fields
		case "IsTransmitting":
			res = fmt.Sprintf("%t", inferedInstance.IsTransmitting)
		case "TransmissionMessage":
			res = inferedInstance.TransmissionMessage
		case "IsTransmittingBackward":
			res = fmt.Sprintf("%t", inferedInstance.IsTransmittingBackward)
		case "TransmissionMessageBackward":
			res = inferedInstance.TransmissionMessageBackward
		case "Scenario":
			if inferedInstance.Scenario != nil {
				res = inferedInstance.Scenario.Name
			}
		case "State":
			enum := inferedInstance.State
			res = enum.ToCodeString()
		case "Name":
			res = inferedInstance.Name
		}
	case Radar:
		switch fieldName {
		// string value of fields
		case "State":
			enum := inferedInstance.State
			res = enum.ToCodeString()
		case "Name":
			res = inferedInstance.Name
		case "Lat":
			res = fmt.Sprintf("%f", inferedInstance.Lat)
		case "Lng":
			res = fmt.Sprintf("%f", inferedInstance.Lng)
		case "Range":
			res = fmt.Sprintf("%f", inferedInstance.Range)
		}
	case Satellite:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Line1":
			res = inferedInstance.Line1
		case "Line2":
			res = inferedInstance.Line2
		case "Lat":
			res = fmt.Sprintf("%f", inferedInstance.Lat)
		case "Lng":
			res = fmt.Sprintf("%f", inferedInstance.Lng)
		case "Heading":
			res = fmt.Sprintf("%f", inferedInstance.Heading)
		case "Level":
			res = fmt.Sprintf("%f", inferedInstance.Level)
		case "Speed":
			res = fmt.Sprintf("%f", inferedInstance.Speed)
		case "VerticalSpeed":
			res = fmt.Sprintf("%f", inferedInstance.VerticalSpeed)
		case "Timestampstring":
			res = inferedInstance.Timestampstring
		}
	case Scenario:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Lat":
			res = fmt.Sprintf("%f", inferedInstance.Lat)
		case "Lng":
			res = fmt.Sprintf("%f", inferedInstance.Lng)
		case "ZoomLevel":
			res = fmt.Sprintf("%f", inferedInstance.ZoomLevel)
		case "MessageVisualSpeed":
			res = fmt.Sprintf("%f", inferedInstance.MessageVisualSpeed)
		}
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template
